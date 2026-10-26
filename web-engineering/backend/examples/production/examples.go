// Package production contains executable examples for backend concept
// entries 13 through 18.
package production

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"time"
)

type Command struct {
	CustomerID string `json:"customer_id"`
	Quantity   int    `json:"quantity"`
}

type FingerprintDecision int

const (
	FingerprintNew FingerprintDecision = iota
	FingerprintReplay
	FingerprintConflict
)

func DecideFingerprint(
	storedFingerprint *[sha256.Size]byte,
	command Command,
) (FingerprintDecision, [sha256.Size]byte) {
	// A typed value gives semantically equal commands one stable field
	// order before hashing.
	canonicalCommand, err := json.Marshal(command)
	if err != nil {
		panic("fixed command shape cannot fail JSON encoding")
	}
	fingerprint := sha256.Sum256(canonicalCommand)
	if storedFingerprint == nil {
		return FingerprintNew, fingerprint
	}
	if *storedFingerprint == fingerprint {
		return FingerprintReplay, fingerprint
	}
	return FingerprintConflict, fingerprint
}

type RetryPolicy struct {
	MaximumAttempts int
	MaximumDelay    time.Duration
	CallBudget      time.Duration
}

func (policy RetryPolicy) NextDelay(
	attempt int,
	baseDelay time.Duration,
	jitter time.Duration,
	remaining time.Duration,
) (time.Duration, bool) {
	if attempt < 1 || attempt >= policy.MaximumAttempts {
		return 0, false
	}
	if baseDelay <= 0 || policy.MaximumDelay <= 0 ||
		jitter < 0 || policy.CallBudget < 0 {
		return 0, false
	}

	// Saturating multiplication prevents duration overflow even when
	// a configuration permits an unexpectedly large attempt count.
	delay := baseDelay
	completedAttempts := 1
	for completedAttempts < attempt {
		if delay >= policy.MaximumDelay ||
			delay > policy.MaximumDelay-delay {
			delay = policy.MaximumDelay
			break
		}
		delay *= 2
		completedAttempts++
	}
	if delay > policy.MaximumDelay {
		delay = policy.MaximumDelay
	}
	delay += jitter
	if delay < 0 || delay+policy.CallBudget > remaining {
		return 0, false
	}
	return delay, true
}

type sqlStateError interface {
	SQLState() string
}

func IsRetryableTransactionError(
	err error,
	retryableCodes map[string]struct{},
) bool {
	var databaseError sqlStateError
	if !errors.As(err, &databaseError) {
		return false
	}
	_, retryable := retryableCodes[databaseError.SQLState()]
	return retryable
}

var ErrStaleClaim = errors.New("stale job claim")

type JobState string

const (
	JobLeased    JobState = "leased"
	JobReady     JobState = "ready"
	JobCompleted JobState = "completed"
	JobDead      JobState = "dead"
)

type Job struct {
	Token    uint64
	Attempts int
	State    JobState
}

func (job *Job) Complete(token uint64) error {
	if job.State != JobLeased || job.Token != token {
		return ErrStaleClaim
	}
	job.State = JobCompleted
	return nil
}

func (job *Job) Fail(
	token uint64,
	retryable bool,
	maximumAttempts int,
) error {
	if job.State != JobLeased || job.Token != token {
		return ErrStaleClaim
	}
	if retryable && job.Attempts < maximumAttempts {
		job.State = JobReady
		return nil
	}
	job.State = JobDead
	return nil
}

type Event struct {
	ID    string
	Delta int
}

type Consumer struct {
	processedEventIDs map[string]struct{}
	Total             int
}

func NewConsumer() *Consumer {
	return &Consumer{
		processedEventIDs: make(map[string]struct{}),
	}
}

func (consumer *Consumer) Apply(event Event) bool {
	if consumer.processedEventIDs == nil {
		consumer.processedEventIDs = make(map[string]struct{})
	}
	if _, processed := consumer.processedEventIDs[event.ID]; processed {
		return false
	}
	// Recording identity with the effect models one consumer
	// transaction; production storage must commit both atomically.
	consumer.processedEventIDs[event.ID] = struct{}{}
	consumer.Total += event.Delta
	return true
}

var ErrStaleToken = errors.New("stale fencing token")

type Message struct {
	Token uint64
	Value string
}

type ProtectedResource struct {
	LatestToken uint64
	Value       string
}

func (resource *ProtectedResource) Apply(message Message) error {
	if message.Token < resource.LatestToken {
		return ErrStaleToken
	}
	resource.LatestToken = message.Token
	resource.Value = message.Value
	return nil
}
