// Service examples cover backend concept entries 9 through 12.
package examples

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"
)

var (
	ErrForbidden            = errors.New("forbidden")
	ErrInvalidQuantity      = errors.New("quantity must be positive")
	ErrUnsafeTelemetryField = errors.New("unsafe telemetry field")
)

type Principal struct {
	Subject string
	Roles   map[string]struct{}
}

type principalContextKey struct{}

func WithPrincipal(
	ctx context.Context,
	principal Principal,
) context.Context {
	return context.WithValue(ctx, principalContextKey{}, principal)
}

func PrincipalFrom(ctx context.Context) (Principal, bool) {
	principal, found := ctx.Value(principalContextKey{}).(Principal)
	return principal, found
}

func RequireRole(ctx context.Context, requiredRole string) error {
	principal, found := PrincipalFrom(ctx)
	if !found {
		return ErrForbidden
	}
	if _, allowed := principal.Roles[requiredRole]; !allowed {
		return ErrForbidden
	}
	return nil
}

type Order struct {
	ID       string
	Quantity int
	PlacedAt time.Time
}

type PlacementRepository interface {
	Save(context.Context, Order) error
}

type Clock interface {
	Now() time.Time
}

type IDSource interface {
	Next() string
}

type PlacementService struct {
	repository PlacementRepository
	clock      Clock
	idSource   IDSource
}

func NewPlacementService(
	repository PlacementRepository,
	clock Clock,
	idSource IDSource,
) PlacementService {
	return PlacementService{
		repository: repository,
		clock:      clock,
		idSource:   idSource,
	}
}

func (service PlacementService) Place(
	ctx context.Context,
	quantity int,
) (Order, error) {
	// Reject invalid input before constructing or persisting an order.
	if quantity <= 0 {
		return Order{}, ErrInvalidQuantity
	}

	// Injected time and identity keep the resulting command replayable.
	order := Order{
		ID:       service.idSource.Next(),
		Quantity: quantity,
		PlacedAt: service.clock.Now(),
	}
	if err := service.repository.Save(ctx, order); err != nil {
		return Order{}, fmt.Errorf("save order: %w", err)
	}
	return order, nil
}

type AccessEvent struct {
	Route  string
	Code   string
	Fields map[string]string
}

func BuildAccessEvent(
	routeTemplate string,
	publicCode string,
	fields map[string]string,
) (AccessEvent, error) {
	allowedFields := map[string]struct{}{
		"method":     {},
		"request_id": {},
	}
	safeFields := make(map[string]string, len(fields))
	for name, value := range fields {
		// An allowlist prevents credentials and raw input from becoming
		// telemetry merely because a caller supplied another field.
		if _, allowed := allowedFields[name]; !allowed {
			return AccessEvent{}, fmt.Errorf(
				"%w: %s",
				ErrUnsafeTelemetryField,
				name,
			)
		}
		safeFields[name] = value
	}
	return AccessEvent{
		Route:  routeTemplate,
		Code:   publicCode,
		Fields: safeFields,
	}, nil
}

type Config struct {
	RequestTimeout time.Duration
	WorkerCount    int
}

func ParseConfig(values map[string]string) (Config, error) {
	timeoutText, found := values["REQUEST_TIMEOUT"]
	if !found {
		return Config{}, errors.New("REQUEST_TIMEOUT is required")
	}
	requestTimeout, err := time.ParseDuration(timeoutText)
	if err != nil || requestTimeout <= 0 {
		return Config{}, errors.New("REQUEST_TIMEOUT must be positive")
	}

	workerText, found := values["WORKER_COUNT"]
	if !found {
		return Config{}, errors.New("WORKER_COUNT is required")
	}
	workerCount, err := strconv.Atoi(workerText)
	if err != nil || workerCount <= 0 {
		return Config{}, errors.New("WORKER_COUNT must be positive")
	}
	return Config{
		RequestTimeout: requestTimeout,
		WorkerCount:    workerCount,
	}, nil
}
