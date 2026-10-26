package delivery

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type fixedVerifier struct {
	err error
}

func (verifier fixedVerifier) Verify(
	_ []byte,
	_ string,
) error {
	return verifier.err
}

func TestWebhookReceiver(t *testing.T) {
	now := time.Unix(1_000, 0)
	receiver := NewWebhookReceiver(
		fixedVerifier{},
		32,
		5*time.Minute,
	)
	accepted, err := receiver.Accept(
		"E-91",
		now,
		now,
		[]byte(`{"ok":true}`),
		"signature",
	)
	if err != nil || !accepted {
		t.Fatalf("Accept() = %v, %v; want true, nil", accepted, err)
	}
	accepted, err = receiver.Accept(
		"E-91",
		now,
		now,
		[]byte(`{"ok":true}`),
		"signature",
	)
	if err != nil || accepted {
		t.Fatalf("duplicate Accept() = %v, %v; want false, nil",
			accepted, err)
	}
	if receiver.PendingWork() != 1 {
		t.Fatalf("pending work = %d; want 1", receiver.PendingWork())
	}

	rejected := NewWebhookReceiver(
		fixedVerifier{err: errors.New("bad signature")},
		32,
		5*time.Minute,
	)
	if _, err := rejected.Accept(
		"E-92",
		now,
		now,
		[]byte(`{"ok":true}`),
		"bad",
	); err == nil {
		t.Fatal("Accept() accepted a bad signature")
	}
}

func TestWebhookReceiverDeduplicatesConcurrentDelivery(t *testing.T) {
	now := time.Unix(1_000, 0)
	receiver := NewWebhookReceiver(
		fixedVerifier{},
		32,
		5*time.Minute,
	)
	var acceptedCount atomic.Int64
	var workers sync.WaitGroup
	workers.Add(20)
	for range 20 {
		go func() {
			defer workers.Done()
			accepted, err := receiver.Accept(
				"E-91",
				now,
				now,
				[]byte(`{"ok":true}`),
				"signature",
			)
			if err != nil {
				t.Errorf("Accept() error = %v", err)
			}
			if accepted {
				acceptedCount.Add(1)
			}
		}()
	}
	workers.Wait()
	if acceptedCount.Load() != 1 || receiver.PendingWork() != 1 {
		t.Fatalf("accepted, pending = %d, %d; want 1, 1",
			acceptedCount.Load(), receiver.PendingWork())
	}
}

func TestRepairPlan(t *testing.T) {
	record := Record{
		TenantID: "tenant-a",
		ID:       "order-1",
		Version:  4,
		Value:    "old",
	}
	plan, err := PlanRepair(record, RepairRequest{
		TenantID:        "tenant-a",
		ID:              "order-1",
		ExpectedVersion: 4,
		NewValue:        "new",
		DryRun:          true,
	})
	if err != nil || plan.Checkpoint == "" {
		t.Fatalf("PlanRepair() = %#v, %v", plan, err)
	}
	if err := ExecuteRepair(&record, plan); err != nil {
		t.Fatalf("ExecuteRepair() error = %v; want nil", err)
	}
	if record.Value != "old" {
		t.Fatalf("dry-run record value = %q; want old", record.Value)
	}

	_, err = PlanRepair(record, RepairRequest{
		TenantID:        "tenant-b",
		ID:              "order-1",
		ExpectedVersion: 4,
		NewValue:        "new",
	})
	if !errors.Is(err, ErrRepairScope) {
		t.Fatalf("PlanRepair() error = %v; want scope error", err)
	}
}

func TestValidateEvidence(t *testing.T) {
	evidence := RequirementEvidence{
		Contract: "cancel is idempotent",
		Test:     "cancel_duplicate_test",
		Signal:   "cancel_outcomes_total",
		Runbook:  "cancel-recovery",
		Status:   "verified",
		Owner:    "orders-team",
		Risk:     "inventory unavailable",
	}
	if err := ValidateEvidence(evidence); err != nil {
		t.Fatalf("ValidateEvidence() error = %v; want nil", err)
	}
	evidence.Runbook = ""
	if err := ValidateEvidence(evidence); err == nil {
		t.Fatal("ValidateEvidence() accepted a missing runbook")
	}
}

type successfulOperation struct{}

func (successfulOperation) Run(context.Context) (string, error) {
	return "placed", nil
}

type waitingOperation struct{}

func (waitingOperation) Run(ctx context.Context) (string, error) {
	<-ctx.Done()
	return "", ctx.Err()
}

func TestEvaluateOperation(t *testing.T) {
	outcome := EvaluateOperation(
		context.Background(),
		successfulOperation{},
	)
	if outcome.Value != "placed" || outcome.Code != OutcomeOK {
		t.Fatalf("success outcome = %#v", outcome)
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	outcome = EvaluateOperation(ctx, waitingOperation{})
	if outcome.Code != OutcomeCanceled {
		t.Fatalf("canceled outcome = %#v", outcome)
	}
}
