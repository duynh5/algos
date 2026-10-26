// Package platform contains executable examples for backend concept
// entries 24 through 28.
package platform

import (
	"errors"
	"fmt"
	"math"
)

var ErrMissingTenant = errors.New("tenant identity is required")

func NewTenantCacheKey(
	tenantID string,
	itemID string,
	version int,
) (string, error) {
	if tenantID == "" {
		return "", ErrMissingTenant
	}
	if itemID == "" || version < 1 {
		return "", errors.New("item and positive version are required")
	}
	// Tenant and version are part of identity, so two authorities or
	// representations cannot alias one cached value.
	return fmt.Sprintf("%s:v%d:%s", tenantID, version, itemID), nil
}

type Permission struct {
	Workload string
	Resource string
	Action   string
}

type Authorizer struct {
	allowedPermissions map[Permission]struct{}
}

func NewAuthorizer(permissions []Permission) Authorizer {
	allowed := make(map[Permission]struct{}, len(permissions))
	for _, permission := range permissions {
		allowed[permission] = struct{}{}
	}
	return Authorizer{allowedPermissions: allowed}
}

func (authorizer Authorizer) Allowed(
	workload string,
	resource string,
	action string,
) bool {
	_, allowed := authorizer.allowedPermissions[Permission{
		Workload: workload,
		Resource: resource,
		Action:   action,
	}]
	return allowed
}

type RotationState struct {
	CurrentVersion  string
	acceptedVersion map[string]struct{}
}

func NewRotationState(
	currentVersion string,
	acceptedVersions []string,
) *RotationState {
	accepted := make(map[string]struct{}, len(acceptedVersions))
	for _, version := range acceptedVersions {
		accepted[version] = struct{}{}
	}
	return &RotationState{
		CurrentVersion:  currentVersion,
		acceptedVersion: accepted,
	}
}

func (state *RotationState) Accepts(version string) bool {
	_, accepted := state.acceptedVersion[version]
	return accepted
}

func (state *RotationState) Retire(version string) {
	if version == state.CurrentVersion {
		return
	}
	delete(state.acceptedVersion, version)
}

var (
	ErrNetwork        = errors.New("network failure")
	ErrAuthentication = errors.New("authentication failure")
	ErrAuthorization  = errors.New("authorization failure")
)

type DiagnosticCategory string

const (
	DiagnosticNetwork        DiagnosticCategory = "network"
	DiagnosticAuthentication DiagnosticCategory = "authentication"
	DiagnosticAuthorization  DiagnosticCategory = "authorization"
	DiagnosticApplication    DiagnosticCategory = "application"
)

func ClassifyDiagnostic(err error) DiagnosticCategory {
	switch {
	case errors.Is(err, ErrNetwork):
		return DiagnosticNetwork
	case errors.Is(err, ErrAuthentication):
		return DiagnosticAuthentication
	case errors.Is(err, ErrAuthorization):
		return DiagnosticAuthorization
	default:
		return DiagnosticApplication
	}
}

type Lifecycle struct {
	Initialized     bool
	DependencyReady bool
	Draining        bool
}

type Health struct {
	Startup   bool
	Readiness bool
	Liveness  bool
}

func (lifecycle Lifecycle) Health() Health {
	return Health{
		Startup: lifecycle.Initialized,
		Readiness: lifecycle.Initialized &&
			lifecycle.DependencyReady &&
			!lifecycle.Draining,
		// Dependency loss should remove traffic, not restart a process
		// whose event loop can still report and recover from the outage.
		Liveness: true,
	}
}

var ErrNoEligibleEvents = errors.New(
	"eligible events must be positive",
)

type ErrorBudget struct {
	AllowedMisses   int64
	ConsumedMisses  int64
	RemainingMisses int64
}

func CalculateErrorBudget(
	eligibleEvents int64,
	goodEvents int64,
	target float64,
) (ErrorBudget, error) {
	if eligibleEvents <= 0 {
		return ErrorBudget{}, ErrNoEligibleEvents
	}
	if goodEvents < 0 || goodEvents > eligibleEvents {
		return ErrorBudget{}, errors.New("good events are out of range")
	}
	if target <= 0 || target > 1 {
		return ErrorBudget{}, errors.New("target must be in (0, 1]")
	}

	allowedMisses := int64(math.Round(
		float64(eligibleEvents) * (1 - target),
	))
	consumedMisses := eligibleEvents - goodEvents
	return ErrorBudget{
		AllowedMisses:   allowedMisses,
		ConsumedMisses:  consumedMisses,
		RemainingMisses: allowedMisses - consumedMisses,
	}, nil
}
