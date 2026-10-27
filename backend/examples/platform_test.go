package examples

import (
	"errors"
	"math"
	"testing"
)

func TestNewTenantCacheKey(t *testing.T) {
	keyA, err := NewTenantCacheKey("tenant-a", "product-1", 2)
	if err != nil {
		t.Fatalf("NewTenantCacheKey() error = %v; want nil", err)
	}
	keyB, err := NewTenantCacheKey("tenant-b", "product-1", 2)
	if err != nil {
		t.Fatalf("NewTenantCacheKey() error = %v; want nil", err)
	}
	if keyA == keyB {
		t.Fatalf("tenant keys alias: %q", keyA)
	}
	_, err = NewTenantCacheKey("", "product-1", 2)
	if !errors.Is(err, ErrMissingTenant) {
		t.Fatalf("NewTenantCacheKey() error = %v; want missing tenant", err)
	}

	// Delimiters inside an identity component must not let different
	// tenants and items produce the same cache key.
	firstAmbiguousKey, err := NewTenantCacheKey("a:v1", "b", 2)
	if err != nil {
		t.Fatalf("first ambiguous key error = %v; want nil", err)
	}
	secondAmbiguousKey, err := NewTenantCacheKey("a", "v2:b", 1)
	if err != nil {
		t.Fatalf("second ambiguous key error = %v; want nil", err)
	}
	if firstAmbiguousKey == secondAmbiguousKey {
		t.Fatalf(
			"distinct identities produced cache key %q",
			firstAmbiguousKey,
		)
	}
}

func TestAuthorizationAndRotation(t *testing.T) {
	authorizer := NewAuthorizer([]Permission{{
		Workload: "order-worker",
		Resource: "queue/orders",
		Action:   "receive",
	}})
	if !authorizer.Allowed(
		"order-worker",
		"queue/orders",
		"receive",
	) {
		t.Fatal("expected queue permission was denied")
	}
	if authorizer.Allowed("order-worker", "bucket/private", "read") {
		t.Fatal("unlisted bucket permission was allowed")
	}

	rotation := NewRotationState("v4", []string{"v3", "v4"})
	if !rotation.Accepts("v4") || !rotation.Accepts("v3") {
		t.Fatal("rotation did not accept current overlap keys")
	}
	rotation.Retire("v3")
	if rotation.Accepts("v3") {
		t.Fatal("rotation accepted a retired key")
	}
}

func TestClassifyDiagnostic(t *testing.T) {
	if got := ClassifyDiagnostic(ErrNetwork); got != DiagnosticNetwork {
		t.Fatalf("network category = %v", got)
	}
	got := ClassifyDiagnostic(ErrAuthorization)
	if got != DiagnosticAuthorization {
		t.Fatalf("authorization category = %v", got)
	}
	got = ClassifyDiagnostic(errors.New("invalid order"))
	if got != DiagnosticApplication {
		t.Fatalf("application category = %v", got)
	}
}

func TestLifecycleHealth(t *testing.T) {
	lifecycle := Lifecycle{
		Initialized:     true,
		DependencyReady: true,
	}
	health := lifecycle.Health()
	if !health.Startup || !health.Readiness || !health.Liveness {
		t.Fatalf("healthy lifecycle = %#v", health)
	}

	lifecycle.DependencyReady = false
	health = lifecycle.Health()
	if !health.Startup || health.Readiness || !health.Liveness {
		t.Fatalf("dependency failure health = %#v", health)
	}
	lifecycle.Draining = true
	if lifecycle.Health().Readiness {
		t.Fatal("draining lifecycle remained ready")
	}
}

func TestCalculateErrorBudget(t *testing.T) {
	budget, err := CalculateErrorBudget(1_000_000, 999_500, 0.999)
	if err != nil {
		t.Fatalf("CalculateErrorBudget() error = %v; want nil", err)
	}
	if budget.AllowedMisses != 1_000 ||
		budget.ConsumedMisses != 500 ||
		budget.RemainingMisses != 500 {
		t.Fatalf("budget = %#v; want 1000, 500, 500", budget)
	}
	_, err = CalculateErrorBudget(0, 0, 0.999)
	if !errors.Is(err, ErrNoEligibleEvents) {
		t.Fatalf("CalculateErrorBudget() error = %v; want no events", err)
	}
	_, err = CalculateErrorBudget(1, 1, math.NaN())
	if err == nil {
		t.Fatal("CalculateErrorBudget() accepted a NaN target")
	}
}
