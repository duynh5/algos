//go:build ignore

package main

import (
	"strings"
	"testing"
)

func TestValidateLineWidth(t *testing.T) {
	testCases := []struct {
		name       string
		contents   string
		wantIssues int
	}{
		{
			name:       "accepts exact limit",
			contents:   strings.Repeat("a", guideLineWidth),
			wantIssues: 0,
		},
		{
			name:       "rejects one character over limit",
			contents:   strings.Repeat("a", guideLineWidth+1),
			wantIssues: 1,
		},
		{
			name:       "counts Unicode characters instead of bytes",
			contents:   strings.Repeat("\u00e9", guideLineWidth),
			wantIssues: 0,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			issues := validateLineWidth("guide.md", []byte(testCase.contents))
			if len(issues) != testCase.wantIssues {
				t.Fatalf(
					"validateLineWidth() returned %d issues; want %d: %v",
					len(issues),
					testCase.wantIssues,
					issues,
				)
			}
		})
	}
}

func TestCollectTargetsRecognizesMultilineAnchor(t *testing.T) {
	const legacyAnchor = "e4-distinguish-missing-keys-from-stored-values-" +
		"map-presence-and-sets"
	prose := []byte("<a\nid=\n" + legacyAnchor + ">\n</a>\n")

	targets, _, anchors, issues := collectTargets("guide.md", prose)
	if len(issues) != 0 {
		t.Fatalf("collectTargets() returned issues: %v", issues)
	}
	if anchors != 1 {
		t.Fatalf("collectTargets() found %d anchors; want 1", anchors)
	}
	if _, exists := targets[legacyAnchor]; !exists {
		t.Fatalf("collectTargets() did not find %q", legacyAnchor)
	}
}
