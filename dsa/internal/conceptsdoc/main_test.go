package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestConceptGuideSpecsCoverEveryActiveGuide(t *testing.T) {
	wantedPaths := []string{
		"dsa/CONCEPTS.md",
		"frontend/CONCEPTS.md",
		"backend/CONCEPTS.md",
	}
	wantedEntries := []int{93, 8, 32}
	if len(conceptGuideSpecs) != len(wantedPaths) {
		t.Fatalf(
			"concept guide count = %d; want %d",
			len(conceptGuideSpecs),
			len(wantedPaths),
		)
	}
	for index, wantedPath := range wantedPaths {
		if conceptGuideSpecs[index].path != wantedPath {
			t.Fatalf(
				"concept guide %d = %q; want %q",
				index,
				conceptGuideSpecs[index].path,
				wantedPath,
			)
		}
		if conceptGuideSpecs[index].expectedEntries !=
			wantedEntries[index] {
			t.Fatalf(
				"concept guide %d entries = %d; want %d",
				index,
				conceptGuideSpecs[index].expectedEntries,
				wantedEntries[index],
			)
		}
	}
}

func TestValidateTeachingEntriesUsesGuideExampleHeading(
	t *testing.T,
) {
	spec := conceptGuideSpec{
		path: "frontend/CONCEPTS.md",
		requiredFields: teachingFields(
			"TypeScript example",
		),
		expectedEntries: 1,
	}
	prose := []byte(strings.Join([]string{
		"## Frontend lesson",
		"### The problem it solves",
		"### The idea in simple terms",
		"### How it works",
		"### Step-by-step examples",
		"### Why it works",
		"### TypeScript example",
		"### Work and memory",
		"### When to use it",
		"### When not to use it",
		"### Common mistakes and edge cases",
		"### Related material",
	}, "\n"))

	issues, entryCount := validateTeachingEntries(spec, prose)
	if len(issues) != 0 {
		t.Fatalf("validateTeachingEntries() issues = %v", issues)
	}
	if entryCount != 1 {
		t.Fatalf("teaching entry count = %d; want 1", entryCount)
	}
}

func TestValidateTeachingEntriesRejectsWrongExampleHeading(
	t *testing.T,
) {
	spec := conceptGuideSpec{
		path: "frontend/CONCEPTS.md",
		requiredFields: teachingFields(
			"TypeScript example",
		),
		expectedEntries: 1,
	}
	prose := []byte(strings.Join([]string{
		"## Frontend lesson",
		"### The problem it solves",
		"### The idea in simple terms",
		"### How it works",
		"### Step-by-step examples",
		"### Why it works",
		"### Go example",
		"### Work and memory",
		"### When to use it",
		"### When not to use it",
		"### Common mistakes and edge cases",
		"### Related material",
	}, "\n"))

	issues, _ := validateTeachingEntries(spec, prose)
	if len(issues) == 0 {
		t.Fatal("wrong example heading returned no issue")
	}
	if !strings.Contains(issues[0], spec.path) ||
		!strings.Contains(issues[0], "TypeScript example") {
		t.Fatalf(
			"wrong example heading issue = %q; want path and field",
			issues[0],
		)
	}
}

func TestCollectTargetsRecognizesMultilineAnchor(t *testing.T) {
	const legacyAnchor = "e4-distinguish-missing-keys-from-stored-" +
		"values-" +
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

func TestValidateDocumentLinksRecognizesMultilineDestination(
	t *testing.T,
) {
	const destination = "https://example.com/a-very-long-destination"
	document := &document{
		path: "links.md",
		prose: []byte(
			"- [Example](\n" + destination + "\n)\n",
		),
		targets: make(map[string]struct{}),
	}

	issues, counts, externalLinks := validateDocumentLinks(document)
	if len(issues) != 0 {
		t.Fatalf("validateDocumentLinks() returned issues: %v", issues)
	}
	if counts.external != 1 {
		t.Fatalf("external link count = %d; want 1", counts.external)
	}
	if _, exists := externalLinks[destination]; !exists {
		t.Fatalf("external links do not contain %q", destination)
	}
}

func TestValidateInboundLinksUsesDestinationGuideTargets(
	t *testing.T,
) {
	repositoryRoot := t.TempDir()
	frontendDirectory := filepath.Join(
		repositoryRoot,
		"frontend",
	)
	if err := os.MkdirAll(frontendDirectory, 0o755); err != nil {
		t.Fatalf("create frontend directory: %v", err)
	}

	// Give the frontend guide a target that is different from every
	// target in the DSA guide.
	frontendGuide := []byte(
		"# Frontend\n\n## Frontend entry\n",
	)
	frontendGuidePath := filepath.Join(
		frontendDirectory,
		"CONCEPTS.md",
	)
	if err := os.WriteFile(
		frontendGuidePath,
		frontendGuide,
		0o600,
	); err != nil {
		t.Fatalf("write frontend guide: %v", err)
	}

	// The inbound link must resolve against the file named in the
	// destination rather than against a single global target set.
	readme := []byte(
		"[Entry](frontend/" +
			"CONCEPTS.md" + "#frontend-entry)\n",
	)
	if err := os.WriteFile(
		filepath.Join(repositoryRoot, "README.md"),
		readme,
		0o600,
	); err != nil {
		t.Fatalf("write repository README: %v", err)
	}

	issues, count := validateInboundLinks(repositoryRoot)
	if len(issues) != 0 {
		t.Fatalf("validateInboundLinks() returned issues: %v", issues)
	}
	if count != 1 {
		t.Fatalf("inbound link count = %d; want 1", count)
	}

	// A fragment absent from the destination guide must still fail.
	invalidReadme := []byte(
		"[Entry](frontend/" +
			"CONCEPTS.md" + "#missing-entry)\n",
	)
	if err := os.WriteFile(
		filepath.Join(repositoryRoot, "README.md"),
		invalidReadme,
		0o600,
	); err != nil {
		t.Fatalf("write invalid repository README: %v", err)
	}

	issues, count = validateInboundLinks(repositoryRoot)
	if len(issues) != 1 ||
		!strings.Contains(issues[0], "#missing-entry") {
		t.Fatalf("missing fragment issues = %v; want one issue", issues)
	}
	if count != 1 {
		t.Fatalf("invalid inbound link count = %d; want 1", count)
	}
}
