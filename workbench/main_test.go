package main

import (
	"strings"
	"testing"

	"github.com/duyserveacc/algos/internal/testrunner"
)

func TestArchiveReadyWorkbenchHasCurrentAttemptAndThreeAgentSolutions(
	t *testing.T,
) {
	if !archiveReady {
		t.Skip("workbench is not ready to archive")
	}

	if got, want := len(agentSolutions), 3; got != want {
		t.Fatalf("len(agentSolutions) = %d, want %d", got, want)
	}
	if got, want := len(solutions), 4; got != want {
		t.Fatalf("len(solutions) = %d, want %d", got, want)
	}
	if solutions[0].Name != "current attempt" {
		t.Fatalf(
			"solutions[0].Name = %q, want current attempt",
			solutions[0].Name,
		)
	}

	seen := make(map[string]struct{}, len(solutions))
	for _, solution := range solutions {
		if strings.TrimSpace(solution.Name) == "" {
			t.Fatal("registered solution has an empty name")
		}
		if _, exists := seen[solution.Name]; exists {
			t.Fatalf("duplicate solution name %q", solution.Name)
		}
		seen[solution.Name] = struct{}{}
	}
}

func TestConfiguredSamplesAgainstEverySolution(t *testing.T) {
	if len(sampleCases) == 0 {
		t.Skip(
			"add workbench samples to enable active-problem checks",
		)
	}

	for _, result := range testrunner.Evaluate(sampleCases, solutions) {
		if !result.Passed {
			t.Errorf(
				"%s / %s: got %q, want %q",
				result.Solution,
				result.Sample,
				result.Got,
				result.Want,
			)
		}
	}
}
