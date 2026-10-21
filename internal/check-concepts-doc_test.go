//go:build ignore

package main

import (
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
