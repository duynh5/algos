package examples

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
	"time"
)

type recordingNotifier struct {
	messages []string
	err      error
}

func (notifier *recordingNotifier) Notify(
	_ context.Context,
	message string,
) error {
	if notifier.err != nil {
		return notifier.err
	}
	notifier.messages = append(notifier.messages, message)
	return nil
}

func TestNotificationServiceSelectsAdapter(t *testing.T) {
	local := &recordingNotifier{}
	service := NewNotificationService(local)
	err := service.Notify(context.Background(), "placed")
	if err != nil {
		t.Fatalf("Notify() error = %v; want nil", err)
	}
	if len(local.messages) != 1 {
		t.Fatalf("local messages = %d; want 1", len(local.messages))
	}

	unavailable := errors.New("publisher unavailable")
	remote := &recordingNotifier{err: unavailable}
	service = NewNotificationService(remote)
	err = service.Notify(context.Background(), "placed")
	if !errors.Is(err, unavailable) {
		t.Fatalf("Notify() error = %v; want unavailable", err)
	}
}

func TestReservationCallState(t *testing.T) {
	state := NewReservationCall("reserve-O-42")
	deadline, err := ChildDeadline(
		time.Unix(100, 0),
		time.Unix(101, 0),
		200*time.Millisecond,
	)
	if err != nil || !deadline.Equal(time.Unix(100, 800_000_000)) {
		t.Fatalf("ChildDeadline() = %v, %v", deadline, err)
	}
	if err := state.Apply(ReservationReserved); err != nil {
		t.Fatalf("Apply() error = %v; want nil", err)
	}
	if state.State != ReservationConfirmed {
		t.Fatalf("state = %v; want confirmed", state.State)
	}

	ambiguous := NewReservationCall("reserve-O-43")
	if err := ambiguous.Apply(ReservationTimedOut); err != nil {
		t.Fatalf("Apply() error = %v; want nil", err)
	}
	if ambiguous.State != ReservationReconcile {
		t.Fatalf("state = %v; want reconcile", ambiguous.State)
	}
	if _, err := ChildDeadline(
		time.Unix(100, 0),
		time.Unix(101, 0),
		-1,
	); !errors.Is(err, ErrInsufficientDeadline) {
		t.Fatalf("ChildDeadline() error = %v; want insufficient", err)
	}
}

func TestSagaReducer(t *testing.T) {
	saga := NewOrderSaga("O-42")
	if err := saga.Apply(SagaPaymentAccepted, 1); err != nil {
		t.Fatalf("Apply() error = %v; want nil", err)
	}
	if saga.State != SagaReserving || saga.Version != 1 {
		t.Fatalf("saga = %#v; want reserving version 1", saga)
	}
	if err := saga.Apply(SagaPaymentAccepted, 1); err != nil {
		t.Fatalf("duplicate Apply() error = %v; want nil", err)
	}
	err := saga.Apply(SagaInventoryReserved, 1)
	if !errors.Is(err, ErrOutOfOrder) {
		t.Fatalf("conflicting duplicate error = %v; want out of order", err)
	}
	err = saga.Apply(SagaInventoryReserved, 3)
	if !errors.Is(err, ErrOutOfOrder) {
		t.Fatalf("Apply() error = %v; want out of order", err)
	}
}

func TestDecodeOrderResponse(t *testing.T) {
	response, err := DecodeOrderResponse([]byte(
		`{"id":"O-42","status":"placed"}`,
	))
	if err != nil || response.Status != OrderPlaced {
		t.Fatalf("DecodeOrderResponse() = %#v, %v", response, err)
	}

	response, err = DecodeOrderResponse([]byte(
		`{"id":"O-43","status":"future","note":"new"}`,
	))
	if err != nil || response.Status != OrderUnknown {
		t.Fatalf("DecodeOrderResponse() = %#v, %v", response, err)
	}
	if response.Note != "new" {
		t.Fatalf("response note = %q; want new", response.Note)
	}
}

func TestBuildEventEnvelope(t *testing.T) {
	envelope, err := BuildEventEnvelope(
		"E9",
		map[string]string{"traceparent": "trace-1"},
		json.RawMessage(`{"order_id":"O-42"}`),
	)
	if err != nil || envelope.Metadata["traceparent"] != "trace-1" {
		t.Fatalf("BuildEventEnvelope() = %#v, %v", envelope, err)
	}

	_, err = BuildEventEnvelope(
		"E10",
		map[string]string{"authorization": "secret"},
		nil,
	)
	if !errors.Is(err, ErrUnsafeMetadata) {
		t.Fatalf("BuildEventEnvelope() error = %v; want unsafe", err)
	}
}
