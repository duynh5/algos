// Package distributed contains executable examples for backend
// concept entries 19 through 23.
package distributed

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type NotificationAdapter interface {
	Notify(context.Context, string) error
}

type NotificationService struct {
	adapter NotificationAdapter
}

func NewNotificationService(
	adapter NotificationAdapter,
) NotificationService {
	return NotificationService{adapter: adapter}
}

func (service NotificationService) Notify(
	ctx context.Context,
	message string,
) error {
	// The composition root selects local or remote delivery. The use
	// case depends only on this narrow capability.
	if err := service.adapter.Notify(ctx, message); err != nil {
		return fmt.Errorf("notify: %w", err)
	}
	return nil
}

var ErrInsufficientDeadline = errors.New("insufficient deadline")

func ChildDeadline(
	now time.Time,
	parentDeadline time.Time,
	reservedTime time.Duration,
) (time.Time, error) {
	if reservedTime < 0 {
		return time.Time{}, ErrInsufficientDeadline
	}
	childDeadline := parentDeadline.Add(-reservedTime)
	if !childDeadline.After(now) {
		return time.Time{}, ErrInsufficientDeadline
	}
	return childDeadline, nil
}

type ReservationState string

const (
	ReservationPending   ReservationState = "pending"
	ReservationConfirmed ReservationState = "confirmed"
	ReservationRejected  ReservationState = "rejected"
	ReservationReconcile ReservationState = "reconcile"
)

type ReservationOutcome string

const (
	ReservationReserved ReservationOutcome = "reserved"
	ReservationDenied   ReservationOutcome = "denied"
	ReservationTimedOut ReservationOutcome = "timed_out"
)

type ReservationCall struct {
	OperationID string
	State       ReservationState
}

func NewReservationCall(operationID string) *ReservationCall {
	return &ReservationCall{
		OperationID: operationID,
		State:       ReservationPending,
	}
}

func (call *ReservationCall) Apply(
	outcome ReservationOutcome,
) error {
	if call.State != ReservationPending {
		return errors.New("reservation call is already resolved")
	}
	switch outcome {
	case ReservationReserved:
		call.State = ReservationConfirmed
	case ReservationDenied:
		call.State = ReservationRejected
	case ReservationTimedOut:
		// A timeout is ambiguous, so it schedules reconciliation rather
		// than claiming that the remote reservation failed.
		call.State = ReservationReconcile
	default:
		return errors.New("unknown reservation outcome")
	}
	return nil
}

var ErrOutOfOrder = errors.New("out-of-order saga event")

type SagaState string

const (
	SagaAwaitingPayment SagaState = "awaiting_payment"
	SagaReserving       SagaState = "reserving"
	SagaCompleted       SagaState = "completed"
)

type SagaEvent string

const (
	SagaPaymentAccepted   SagaEvent = "payment_accepted"
	SagaInventoryReserved SagaEvent = "inventory_reserved"
)

type OrderSaga struct {
	OrderID   string
	State     SagaState
	Version   int
	LastEvent SagaEvent
}

func NewOrderSaga(orderID string) *OrderSaga {
	return &OrderSaga{
		OrderID: orderID,
		State:   SagaAwaitingPayment,
	}
}

func (saga *OrderSaga) Apply(event SagaEvent, version int) error {
	if version == saga.Version && event == saga.LastEvent {
		return nil
	}
	if version <= saga.Version {
		return ErrOutOfOrder
	}
	if version != saga.Version+1 {
		return ErrOutOfOrder
	}

	switch {
	case saga.State == SagaAwaitingPayment &&
		event == SagaPaymentAccepted:
		saga.State = SagaReserving
	case saga.State == SagaReserving &&
		event == SagaInventoryReserved:
		saga.State = SagaCompleted
	default:
		return ErrOutOfOrder
	}
	saga.Version = version
	saga.LastEvent = event
	return nil
}

type OrderStatus string

const (
	OrderPlaced   OrderStatus = "placed"
	OrderCanceled OrderStatus = "canceled"
	OrderUnknown  OrderStatus = "unknown"
)

type OrderResponse struct {
	ID     string      `json:"id"`
	Status OrderStatus `json:"-"`
	Note   string      `json:"note,omitempty"`
}

func DecodeOrderResponse(data []byte) (OrderResponse, error) {
	var wireResponse struct {
		ID     string `json:"id"`
		Status string `json:"status"`
		Note   string `json:"note,omitempty"`
	}
	if err := json.Unmarshal(data, &wireResponse); err != nil {
		return OrderResponse{}, fmt.Errorf("decode order: %w", err)
	}

	status := OrderStatus(wireResponse.Status)
	switch status {
	case OrderPlaced, OrderCanceled:
	default:
		status = OrderUnknown
	}
	return OrderResponse{
		ID:     wireResponse.ID,
		Status: status,
		Note:   wireResponse.Note,
	}, nil
}

var ErrUnsafeMetadata = errors.New("unsafe event metadata")

type EventEnvelope struct {
	ID       string
	Metadata map[string]string
	Payload  json.RawMessage
}

func BuildEventEnvelope(
	eventID string,
	metadata map[string]string,
	payload json.RawMessage,
) (EventEnvelope, error) {
	allowedMetadata := map[string]struct{}{
		"correlation_id": {},
		"traceparent":    {},
		"tracestate":     {},
	}
	safeMetadata := make(map[string]string, len(metadata))
	for name, value := range metadata {
		if _, allowed := allowedMetadata[name]; !allowed {
			return EventEnvelope{}, fmt.Errorf(
				"%w: %s",
				ErrUnsafeMetadata,
				name,
			)
		}
		safeMetadata[name] = value
	}
	return EventEnvelope{
		ID:       eventID,
		Metadata: safeMetadata,
		Payload:  append(json.RawMessage(nil), payload...),
	}, nil
}
