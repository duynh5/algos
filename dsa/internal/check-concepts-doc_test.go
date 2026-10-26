//go:build ignore

package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

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
		"web-engineering",
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
		"[Entry](web-engineering/frontend/" +
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
		"[Entry](web-engineering/frontend/" +
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
