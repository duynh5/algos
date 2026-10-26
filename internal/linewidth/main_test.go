package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestExistingRepositoryPathsSkipsDeletedFiles(t *testing.T) {
	repositoryRoot := t.TempDir()
	existingPath := "existing.md"
	if err := os.WriteFile(
		filepath.Join(repositoryRoot, existingPath),
		[]byte("content"),
		0o600,
	); err != nil {
		t.Fatalf("write existing file: %v", err)
	}

	paths, err := existingRepositoryPaths(
		repositoryRoot,
		[][]byte{
			[]byte(existingPath),
			[]byte("deleted.md"),
		},
	)
	if err != nil {
		t.Fatalf("existingRepositoryPaths() error = %v", err)
	}
	if len(paths) != 1 || paths[0] != existingPath {
		t.Fatalf("existing paths = %v; want [%s]", paths, existingPath)
	}
}

func TestValidateContents(t *testing.T) {
	testCases := []struct {
		name       string
		contents   string
		wantIssues int
	}{
		{
			name:       "accepts exact limit",
			contents:   strings.Repeat("a", maximumLineWidth),
			wantIssues: 0,
		},
		{
			name:       "rejects one character over limit",
			contents:   strings.Repeat("a", maximumLineWidth+1),
			wantIssues: 1,
		},
		{
			name:       "counts Unicode characters instead of bytes",
			contents:   strings.Repeat("\u00e9", maximumLineWidth),
			wantIssues: 0,
		},
		{
			name: "rejects indivisible URL",
			contents: "https://example.com/" +
				strings.Repeat("a", maximumLineWidth),
			wantIssues: 1,
		},
		{
			name:       "rejects prose with a short URL",
			contents:   strings.Repeat("word ", 15) + "https://x.test",
			wantIssues: 1,
		},
		{
			name: "rejects prose beside a long URL",
			contents: "source: https://example.com/" +
				strings.Repeat("a", maximumLineWidth),
			wantIssues: 1,
		},
		{
			name: "rejects indivisible repository path",
			contents: strings.Repeat("long_name_", 8) +
				"problem.go",
			wantIssues: 1,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			issues := validateContents(
				"example.txt",
				[]byte(testCase.contents),
			)
			if len(issues) != testCase.wantIssues {
				t.Fatalf(
					"validateContents() returned %d issues; want %d: %v",
					len(issues),
					testCase.wantIssues,
					issues,
				)
			}
		})
	}
}

func TestIsExemptPath(t *testing.T) {
	for path := range exemptPaths {
		if !isExemptPath(path) {
			t.Errorf("isExemptPath(%q) = false; want true", path)
		}
	}
	if !isExemptPath("CONCEPTS_STANDARDS.md") {
		t.Error("root concept standards must be exempt")
	}
	if isExemptPath("dsa/CONCEPTS_STANDARDS.md") {
		t.Error("old DSA concept standards path must not be exempt")
	}
	if isExemptPath("docs/AGENTS.md") {
		t.Error("nested AGENTS.md must not be exempt")
	}
}
