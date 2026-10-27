package examples

import (
	"errors"
	"testing"
	"time"
)

func TestDecideFingerprint(t *testing.T) {
	command := Command{CustomerID: "C-17", Quantity: 2}
	decision, fingerprint := DecideFingerprint(nil, command)
	if decision != FingerprintNew {
		t.Fatalf("decision = %v; want new", decision)
	}

	decision, _ = DecideFingerprint(&fingerprint, command)
	if decision != FingerprintReplay {
		t.Fatalf("decision = %v; want replay", decision)
	}
	changed := Command{CustomerID: "C-17", Quantity: 3}
	decision, _ = DecideFingerprint(&fingerprint, changed)
	if decision != FingerprintConflict {
		t.Fatalf("decision = %v; want conflict", decision)
	}
}

func TestRetryPolicy(t *testing.T) {
	policy := RetryPolicy{
		MaximumAttempts: 3,
		MaximumDelay:    200 * time.Millisecond,
		CallBudget:      100 * time.Millisecond,
	}
	delay, allowed := policy.NextDelay(
		1,
		100*time.Millisecond,
		20*time.Millisecond,
		600*time.Millisecond,
	)
	if !allowed || delay != 120*time.Millisecond {
		t.Fatalf("NextDelay() = %v, %v; want 120ms, true",
			delay, allowed)
	}
	_, allowed = policy.NextDelay(
		2,
		100*time.Millisecond,
		0,
		10*time.Millisecond,
	)
	if allowed {
		t.Fatal("NextDelay() allowed work past the remaining budget")
	}
	overflowPolicy := RetryPolicy{
		MaximumAttempts: 100,
		MaximumDelay:    time.Second,
		CallBudget:      time.Millisecond,
	}
	delay, allowed = overflowPolicy.NextDelay(
		80,
		time.Second,
		0,
		2*time.Second,
	)
	if !allowed || delay != time.Second {
		t.Fatalf("overflow-safe delay = %v, %v; want 1s, true",
			delay, allowed)
	}

	// Adding a maximum call budget to any positive delay overflows a
	// duration. The policy must reject the work before doing that sum.
	maximumDuration := time.Duration(1<<63 - 1)
	overflowPolicy = RetryPolicy{
		MaximumAttempts: 2,
		MaximumDelay:    time.Nanosecond,
		CallBudget:      maximumDuration,
	}
	_, allowed = overflowPolicy.NextDelay(
		1,
		time.Nanosecond,
		0,
		maximumDuration,
	)
	if allowed {
		t.Fatal("NextDelay() allowed an overflowing total budget")
	}
}

type databaseError struct {
	code string
}

func (err databaseError) Error() string {
	return "database error"
}

func (err databaseError) SQLState() string {
	return err.code
}

func TestIsRetryableTransactionError(t *testing.T) {
	retryableCodes := map[string]struct{}{
		"40001": {},
		"40P01": {},
	}
	if !IsRetryableTransactionError(
		databaseError{code: "40001"},
		retryableCodes,
	) {
		t.Fatal("serialization failure was not retryable")
	}
	if IsRetryableTransactionError(
		databaseError{code: "23505"},
		retryableCodes,
	) {
		t.Fatal("constraint failure was retryable")
	}
}

func TestJobClaimToken(t *testing.T) {
	job := Job{Token: 8, Attempts: 1, State: JobLeased}
	if err := job.Complete(7); !errors.Is(err, ErrStaleClaim) {
		t.Fatalf("Complete() error = %v; want stale claim", err)
	}
	if job.State != JobLeased {
		t.Fatalf("job state = %v; want leased", job.State)
	}
	if err := job.Complete(8); err != nil {
		t.Fatalf("Complete() error = %v; want nil", err)
	}
	if job.State != JobCompleted {
		t.Fatalf("job state = %v; want completed", job.State)
	}

	failed := Job{Token: 2, Attempts: 3, State: JobLeased}
	if err := failed.Fail(2, true, 3); err != nil {
		t.Fatalf("Fail() error = %v; want nil", err)
	}
	if failed.State != JobDead {
		t.Fatalf("job state = %v; want dead", failed.State)
	}
}

func TestConsumerDeduplicates(t *testing.T) {
	consumer := NewConsumer()
	event := Event{ID: "E9", Delta: 3}
	if applied := consumer.Apply(event); !applied {
		t.Fatal("first event was not applied")
	}
	if applied := consumer.Apply(event); applied {
		t.Fatal("duplicate event was applied")
	}
	if consumer.Total != 3 {
		t.Fatalf("consumer total = %d; want 3", consumer.Total)
	}
}

func TestConsumerZeroValue(t *testing.T) {
	consumer := &Consumer{}
	if applied := consumer.Apply(Event{ID: "E1", Delta: 1}); !applied {
		t.Fatal("zero-value consumer did not apply its first event")
	}
}

func TestProtectedResourceFencesStaleMessages(t *testing.T) {
	resource := ProtectedResource{}
	err := resource.Apply(Message{Token: 8, Value: "new"})
	if err != nil {
		t.Fatalf("Apply() error = %v; want nil", err)
	}
	err = resource.Apply(Message{Token: 7, Value: "stale"})
	if !errors.Is(err, ErrStaleToken) {
		t.Fatalf("Apply() error = %v; want stale token", err)
	}
	if resource.Value != "new" || resource.LatestToken != 8 {
		t.Fatalf("resource = %#v; want token 8 and new value", resource)
	}
}
