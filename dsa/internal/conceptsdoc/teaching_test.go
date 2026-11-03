package main

import (
	"strings"
	"testing"
)

func TestTeachingLayoutsDuringMigration(t *testing.T) {
	modern := []string{
		"Problem and contract",
		"Mental model",
		"Worked walkthrough",
		"Implementation",
		"Why it works",
		"Costs and tradeoffs",
		"Edge cases and mistakes",
		"When to choose it",
	}
	legacy := teachingFields("Go example")
	// Each lesson owns its fields; fields cannot cross a boundary.
	lesson := func(fields []string) string {
		return "## Lesson\n### " +
			strings.Join(fields, "\n### ") + "\n"
	}
	tests := []struct {
		name    string
		prose   string
		entries int
		valid   bool
	}{
		{"modern", lesson(modern), 1, true},
		{"mixed", lesson(modern) + lesson(legacy), 2, true},
		{"reordered", lesson(append(
			[]string{modern[0], modern[2], modern[1]},
			modern[3:]...,
		)), 1, false},
		{"interrupted", lesson(modern[:3]) +
			lesson(modern[3:]), 1, false},
		{"wrong level", strings.Replace(
			lesson(modern), "### Mental", "#### Mental", 1,
		), 1, false},
		{"missing final", lesson(modern[:7]), 1, false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			spec := conceptGuideSpec{
				path:            "test.md",
				requiredFields:  legacy,
				expectedEntries: test.entries,
			}
			issues, _ := validateTeachingEntries(
				spec, []byte(test.prose),
			)
			if (len(issues) == 0) != test.valid {
				t.Fatalf("issues = %v; valid = %v", issues, test.valid)
			}
		})
	}
}

func TestModernWalkthroughChecksTraceLabels(t *testing.T) {
	prose := []byte("##### Worked walkthrough\n" +
		"Initial state: empty.\n1. Decision: insert.\n" +
		"Known correct: one item.\n##### Implementation\n")
	issues := validateTraceItems("test.md", prose)
	if len(issues) != 1 || !strings.Contains(issues[0], "After") {
		t.Fatalf("issues = %v; want missing After", issues)
	}
}
