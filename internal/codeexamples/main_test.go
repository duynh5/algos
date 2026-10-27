package main

import (
	"strings"
	"testing"
)

func TestReadCodeBlocksSelectsValidatedLanguages(t *testing.T) {
	contents := []byte(strings.Join([]string{
		"# Examples",
		"",
		"```go",
		"func answer() int { return 42 }",
		"```",
		"",
		"```text",
		"not code",
		"```",
		"",
		"```typescript",
		"const answer: number = 42;",
		"```",
		"",
	}, "\n"))

	blocks, issues := readCodeBlocks("guide.md", contents)
	if len(issues) != 0 {
		t.Fatalf("readCodeBlocks() issues = %v", issues)
	}
	if len(blocks) != 2 {
		t.Fatalf("readCodeBlocks() found %d blocks; want 2", len(blocks))
	}
	if blocks[0].language != "go" || blocks[0].startLine != 4 {
		t.Fatalf("first block = %#v; want Go block at line 4", blocks[0])
	}
	if blocks[1].language != "typescript" ||
		blocks[1].startLine != 12 {
		t.Fatalf(
			"second block = %#v; want TypeScript block at line 12",
			blocks[1],
		)
	}
}

func TestReadCodeBlocksRejectsUnclosedFence(t *testing.T) {
	blocks, issues := readCodeBlocks(
		"guide.md",
		[]byte("```go\nfunc answer() int { return 42 }\n"),
	)
	if len(blocks) != 0 {
		t.Fatalf("readCodeBlocks() blocks = %v; want none", blocks)
	}
	if len(issues) != 1 ||
		!strings.Contains(issues[0], "unclosed code fence") {
		t.Fatalf(
			"readCodeBlocks() issues = %v; want unclosed fence",
			issues,
		)
	}
}

func TestValidateFormattedBlockRejectsUnformattedGo(t *testing.T) {
	issues := validateFormattedBlock(
		"guide.md:2",
		"gofmt",
		nil,
		"func answer() int {\nreturn 42\n}\n",
	)
	if len(issues) != 1 ||
		!strings.Contains(issues[0], "not formatted") {
		t.Fatalf("format issues = %v; want one formatting issue", issues)
	}
}
