// Package delivery contains executable examples for backend concept
// entries 29 through 32.
package delivery

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

type WebhookVerifier interface {
	Verify(body []byte, signature string) error
}

type WebhookReceiver struct {
	mu           sync.Mutex
	verifier     WebhookVerifier
	maximumBytes int
	replayWindow time.Duration
	seenEventIDs map[string]struct{}
	pendingWork  int
}

func NewWebhookReceiver(
	verifier WebhookVerifier,
	maximumBytes int,
	replayWindow time.Duration,
) *WebhookReceiver {
	return &WebhookReceiver{
		verifier:     verifier,
		maximumBytes: maximumBytes,
		replayWindow: replayWindow,
		seenEventIDs: make(map[string]struct{}),
	}
}

func (receiver *WebhookReceiver) Accept(
	eventID string,
	eventTime time.Time,
	now time.Time,
	body []byte,
	signature string,
) (bool, error) {
	if len(body) > receiver.maximumBytes {
		return false, errors.New("webhook body exceeds limit")
	}
	// Verification uses the unchanged bytes before any parser can
	// normalize or discard signed information.
	if err := receiver.verifier.Verify(body, signature); err != nil {
		return false, fmt.Errorf("verify webhook: %w", err)
	}
	age := now.Sub(eventTime)
	if age < -receiver.replayWindow || age > receiver.replayWindow {
		return false, errors.New("webhook timestamp outside replay window")
	}
	receiver.mu.Lock()
	defer receiver.mu.Unlock()
	if _, duplicate := receiver.seenEventIDs[eventID]; duplicate {
		return false, nil
	}

	// Durable storage would commit identity and queued work together.
	receiver.seenEventIDs[eventID] = struct{}{}
	receiver.pendingWork++
	return true, nil
}

func (receiver *WebhookReceiver) PendingWork() int {
	receiver.mu.Lock()
	defer receiver.mu.Unlock()
	return receiver.pendingWork
}

var ErrRepairScope = errors.New("repair target is outside scope")

type Record struct {
	TenantID string
	ID       string
	Version  int
	Value    string
}

type RepairRequest struct {
	TenantID        string
	ID              string
	ExpectedVersion int
	NewValue        string
	DryRun          bool
}

type RepairPlan struct {
	TenantID        string
	ID              string
	ExpectedVersion int
	NewValue        string
	DryRun          bool
	Checkpoint      string
}

func PlanRepair(
	record Record,
	request RepairRequest,
) (RepairPlan, error) {
	if request.TenantID == "" || request.ID == "" ||
		record.TenantID != request.TenantID ||
		record.ID != request.ID {
		return RepairPlan{}, ErrRepairScope
	}
	if record.Version != request.ExpectedVersion {
		return RepairPlan{}, errors.New("record version changed")
	}
	return RepairPlan{
		TenantID:        request.TenantID,
		ID:              request.ID,
		ExpectedVersion: request.ExpectedVersion,
		NewValue:        request.NewValue,
		DryRun:          request.DryRun,
		Checkpoint: fmt.Sprintf(
			"%s/%s/v%d",
			request.TenantID,
			request.ID,
			request.ExpectedVersion,
		),
	}, nil
}

func ExecuteRepair(record *Record, plan RepairPlan) error {
	if record.TenantID != plan.TenantID || record.ID != plan.ID {
		return ErrRepairScope
	}
	if record.Version != plan.ExpectedVersion {
		return errors.New("record version changed")
	}
	if plan.DryRun {
		return nil
	}
	record.Value = plan.NewValue
	record.Version++
	return nil
}

type RequirementEvidence struct {
	Contract string
	Test     string
	Signal   string
	Runbook  string
	Status   string
	Owner    string
	Risk     string
}

func ValidateEvidence(evidence RequirementEvidence) error {
	requiredValues := []struct {
		name  string
		value string
	}{
		{name: "contract", value: evidence.Contract},
		{name: "test", value: evidence.Test},
		{name: "signal", value: evidence.Signal},
		{name: "runbook", value: evidence.Runbook},
		{name: "status", value: evidence.Status},
		{name: "owner", value: evidence.Owner},
		{name: "risk", value: evidence.Risk},
	}
	for _, required := range requiredValues {
		if required.value == "" {
			return fmt.Errorf("%s evidence is required", required.name)
		}
	}
	return nil
}

type Operation interface {
	Run(context.Context) (string, error)
}

type OutcomeCode string

const (
	OutcomeOK       OutcomeCode = "ok"
	OutcomeCanceled OutcomeCode = "canceled"
	OutcomeFailed   OutcomeCode = "failed"
)

type Outcome struct {
	Value string
	Code  OutcomeCode
}

func EvaluateOperation(
	ctx context.Context,
	operation Operation,
) Outcome {
	value, err := operation.Run(ctx)
	switch {
	case err == nil:
		return Outcome{Value: value, Code: OutcomeOK}
	case errors.Is(err, context.Canceled),
		errors.Is(err, context.DeadlineExceeded):
		return Outcome{Code: OutcomeCanceled}
	default:
		return Outcome{Code: OutcomeFailed}
	}
}
