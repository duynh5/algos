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

func TestValidateCodeBlockContents(t *testing.T) {
	longLine := strings.Repeat("a", maximumLineWidth+1)
	contents := strings.Join(
		[]string{
			longLine,
			"```go",
			longLine,
			"```",
		},
		"\n",
	)

	issues := validateCodeBlockContents(
		"README.md",
		[]byte(contents),
	)
	if len(issues) != 1 {
		t.Fatalf(
			"validateCodeBlockContents() returned %d issues; want 1: %v",
			len(issues),
			issues,
		)
	}
}

func TestScopeForPath(t *testing.T) {
	testCases := []struct {
		path      string
		wantScope validationScope
	}{
		{path: "dsa/CONCEPTS.md", wantScope: validateWholeFile},
		{path: "main.go", wantScope: validateWholeFile},
		{path: "main.go.template", wantScope: validateWholeFile},
		{path: "script.sh", wantScope: validateWholeFile},
		{path: "README.md", wantScope: validateCodeBlocks},
		{path: "data.csv", wantScope: skipFile},
	}

	for _, testCase := range testCases {
		t.Run(testCase.path, func(t *testing.T) {
			gotScope := scopeForPath(testCase.path)
			if gotScope != testCase.wantScope {
				t.Errorf(
					"scopeForPath(%q) = %d; want %d",
					testCase.path,
					gotScope,
					testCase.wantScope,
				)
			}
		})
	}
}
