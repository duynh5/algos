package examples

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"
)

func TestRequireRole(t *testing.T) {
	principal := Principal{
		Subject: "operator-1",
		Roles:   map[string]struct{}{"admin": {}},
	}
	ctx := WithPrincipal(context.Background(), principal)

	if err := RequireRole(ctx, "admin"); err != nil {
		t.Fatalf("RequireRole() error = %v; want nil", err)
	}
	err := RequireRole(ctx, "customer")
	if !errors.Is(err, ErrForbidden) {
		t.Fatalf("RequireRole() error = %v; want forbidden", err)
	}
	if _, found := PrincipalFrom(context.Background()); found {
		t.Fatal("PrincipalFrom() found an absent principal")
	}
}

type recordingRepository struct {
	orders []Order
}

func (repository *recordingRepository) Save(
	_ context.Context,
	order Order,
) error {
	repository.orders = append(repository.orders, order)
	return nil
}

type fixedClock struct {
	now time.Time
}

func (clock fixedClock) Now() time.Time {
	return clock.now
}

type fixedIDSource struct {
	id string
}

func (source fixedIDSource) Next() string {
	return source.id
}

func TestPlacementService(t *testing.T) {
	placedAt := time.Date(2026, time.January, 2, 3, 4, 5, 0, time.UTC)
	repository := &recordingRepository{}
	service := NewPlacementService(
		repository,
		fixedClock{now: placedAt},
		fixedIDSource{id: "order-123"},
	)

	order, err := service.Place(context.Background(), 2)
	if err != nil {
		t.Fatalf("Place() error = %v; want nil", err)
	}
	if order.ID != "order-123" || order.PlacedAt != placedAt {
		t.Fatalf("Place() order = %#v; want fixed values", order)
	}
	if len(repository.orders) != 1 {
		t.Fatalf("saved orders = %d; want 1", len(repository.orders))
	}

	_, err = service.Place(context.Background(), 0)
	if !errors.Is(err, ErrInvalidQuantity) {
		t.Fatalf("Place() error = %v; want invalid quantity", err)
	}
	if len(repository.orders) != 1 {
		t.Fatalf("saved orders = %d after rejection; want 1",
			len(repository.orders))
	}
}

func TestBuildAccessEvent(t *testing.T) {
	event, err := BuildAccessEvent(
		"/v1/orders/{id}",
		"order_denied",
		map[string]string{
			"method":     "GET",
			"request_id": "request-7",
		},
	)
	if err != nil {
		t.Fatalf("BuildAccessEvent() error = %v; want nil", err)
	}
	if event.Route != "/v1/orders/{id}" {
		t.Fatalf("event route = %q; want route template", event.Route)
	}

	_, err = BuildAccessEvent(
		"/v1/orders/{id}",
		"order_denied",
		map[string]string{"authorization": "Bearer secret"},
	)
	if !errors.Is(err, ErrUnsafeTelemetryField) {
		t.Fatalf("BuildAccessEvent() error = %v; want unsafe field", err)
	}
}

func TestParseConfig(t *testing.T) {
	config, err := ParseConfig(map[string]string{
		"REQUEST_TIMEOUT": "750ms",
		"WORKER_COUNT":    "4",
	})
	if err != nil {
		t.Fatalf("ParseConfig() error = %v; want nil", err)
	}
	if config.RequestTimeout != 750*time.Millisecond ||
		config.WorkerCount != 4 {
		t.Fatalf("ParseConfig() = %#v; want 750ms and 4", config)
	}

	_, err = ParseConfig(map[string]string{
		"DATABASE_PASSWORD": "highly-secret",
		"WORKER_COUNT":      "4",
	})
	if err == nil {
		t.Fatal("ParseConfig() accepted a missing timeout")
	}
	if strings.Contains(err.Error(), "highly-secret") {
		t.Fatalf("ParseConfig() leaked secret in error: %v", err)
	}
}
