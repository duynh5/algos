package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func readDocument(path string) (*document, []string) {
	contents, err := os.ReadFile(path)
	if err != nil {
		return nil, []string{err.Error()}
	}

	prose, fences, goBlocks, textBlocks, fenceIssues :=
		proseOutsideFences(path, contents)
	targets, headings, anchors, targetIssues := collectTargets(
		path,
		prose,
	)
	return &document{
		path:            path,
		prose:           prose,
		targets:         targets,
		headingCount:    headings,
		fenceCount:      fences,
		goBlockCount:    goBlocks,
		textBlockCount:  textBlocks,
		explicitAnchors: anchors,
	}, append(fenceIssues, targetIssues...)
}

func proseOutsideFences(
	path string,
	contents []byte,
) ([]byte, int, int, int, []string) {
	var prose bytes.Buffer
	scanner := bufio.NewScanner(bytes.NewReader(contents))
	scanner.Buffer(make([]byte, 64*1024), 1024*1024)
	inFence := false
	fenceStart := 0
	blocks := 0
	goBlocks := 0
	textBlocks := 0
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "```") {
			if !inFence {
				inFence = true
				fenceStart = lineNumber
				blocks++
				switch strings.TrimSpace(strings.TrimPrefix(trimmed, "```")) {
				case "go":
					goBlocks++
				case "text":
					textBlocks++
				}
			} else {
				inFence = false
			}
			prose.WriteByte('\n')
			continue
		}
		if inFence {
			prose.WriteByte('\n')
			continue
		}
		prose.WriteString(line)
		prose.WriteByte('\n')
	}

	issues := make([]string, 0, 2)
	if err := scanner.Err(); err != nil {
		issues = append(
			issues,
			fmt.Sprintf("%s: scan: %v", path, err),
		)
	}
	if inFence {
		issues = append(
			issues,
			fmt.Sprintf(
				"%s:%d: unclosed fenced code block",
				path,
				fenceStart,
			),
		)
	}
	return prose.Bytes(), blocks, goBlocks, textBlocks, issues
}

func collectTargets(
	path string,
	prose []byte,
) (map[string]struct{}, int, int, []string) {
	targets := make(map[string]struct{})
	targetOrigins := make(map[string]string)
	slugUses := make(map[string]int)
	explicitUses := make(map[string]int)
	issues := make([]string, 0)
	headings := 0
	anchors := 0
	previousLevel := 0

	scanner := bufio.NewScanner(bytes.NewReader(prose))
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		if match := headingPattern.FindStringSubmatch(line); match != nil {
			level := len(match[1])
			if previousLevel > 0 && level > previousLevel+1 {
				issues = append(
					issues,
					fmt.Sprintf(
						"%s:%d: heading jumps from level %d to %d",
						path,
						lineNumber,
						previousLevel,
						level,
					),
				)
			}
			previousLevel = level
			base := githubSlug(match[2])
			slug := base
			if count := slugUses[base]; count > 0 {
				slug = fmt.Sprintf("%s-%d", base, count)
			}
			slugUses[base]++
			if origin, exists := targetOrigins[slug]; exists {
				issues = append(
					issues,
					fmt.Sprintf(
						"%s:%d: generated target %q duplicates %s",
						path,
						lineNumber,
						slug,
						origin,
					),
				)
			}
			targets[slug] = struct{}{}
			targetOrigins[slug] = fmt.Sprintf(
				"heading at line %d",
				lineNumber,
			)
			headings++
		}
	}
	// Scan the complete prose so valid HTML anchors may span physical
	// lines.
	for _, match := range anchorPattern.FindAllSubmatchIndex(prose, -1) {
		anchor := string(prose[match[2]:match[3]])
		lineNumber := bytes.Count(prose[:match[0]], []byte{'\n'}) + 1
		explicitUses[anchor]++
		if explicitUses[anchor] > 1 {
			issues = append(
				issues,
				fmt.Sprintf(
					"%s:%d: duplicate explicit anchor %q",
					path,
					lineNumber,
					anchor,
				),
			)
		}
		if origin, exists := targetOrigins[anchor]; exists {
			issues = append(
				issues,
				fmt.Sprintf(
					"%s:%d: explicit anchor %q duplicates %s",
					path,
					lineNumber,
					anchor,
					origin,
				),
			)
		}
		targets[anchor] = struct{}{}
		targetOrigins[anchor] = fmt.Sprintf(
			"explicit anchor at line %d",
			lineNumber,
		)
		anchors++
	}
	return targets, headings, anchors, issues
}

func githubSlug(heading string) string {
	var slug strings.Builder
	for _, current := range strings.ToLower(heading) {
		switch {
		case unicode.IsLetter(current),
			unicode.IsDigit(current),
			current == '-',
			current == '_':
			slug.WriteRune(current)
		case unicode.IsSpace(current):
			slug.WriteByte('-')
		}
	}
	return strings.Trim(slug.String(), "-")
}

func validateNoTables(path string, prose []byte) []string {
	lines := strings.Split(string(prose), "\n")
	for index := 1; index < len(lines); index++ {
		if tablePattern.MatchString(lines[index-1]) &&
			tableRule.MatchString(lines[index]) {
			return []string{
				fmt.Sprintf(
					"%s:%d: Markdown tables are not allowed in "+
						"teaching entries",
					path,
					index,
				),
			}
		}
	}
	return nil
}

func validateTraceItems(path string, prose []byte) []string {
	standaloneLabels := []string{
		"Before",
		"Decision",
		"After",
		"Known correct",
	}
	chainedLabels := []string{
		"Decision",
		"After",
		"Known correct",
	}
	issues := make([]string, 0)
	inExamples := false
	chainedTrace := false
	itemLine := 0
	var item strings.Builder

	flushItem := func() {
		if itemLine == 0 {
			return
		}
		itemText := item.String()
		requiredLabels := standaloneLabels
		if chainedTrace {
			requiredLabels = chainedLabels
		}
		foundLabels := 0
		missingLabels := make([]string, 0, len(requiredLabels))
		for _, label := range requiredLabels {
			marker := label + ":"
			if strings.Contains(itemText, marker) {
				foundLabels++
				continue
			}
			missingLabels = append(missingLabels, label)
		}
		// A numbered explanation without state labels may accompany a
		// concrete diagram. Once an item adopts the trace vocabulary,
		// require a complete standalone or chained transition.
		if foundLabels > 0 && len(missingLabels) > 0 {
			issues = append(
				issues,
				fmt.Sprintf(
					"%s:%d: trace item is missing %s",
					path,
					itemLine,
					strings.Join(missingLabels, ", "),
				),
			)
		}
		item.Reset()
		itemLine = 0
	}

	scanner := bufio.NewScanner(bytes.NewReader(prose))
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		if match := headingPattern.FindStringSubmatch(line); match != nil {
			switch match[2] {
			case "Step-by-step examples", "Worked walkthrough":
				inExamples = true
				chainedTrace = false
			case "Why it works", "Implementation":
				if inExamples {
					flushItem()
					inExamples = false
					chainedTrace = false
				}
			}
			continue
		}
		if !inExamples {
			continue
		}
		if strings.Contains(line, "Initial state:") {
			flushItem()
			chainedTrace = true
			continue
		}
		if traceItemPattern.MatchString(line) {
			flushItem()
			itemLine = lineNumber
		}
		if itemLine > 0 {
			item.WriteString(line)
			item.WriteByte('\n')
		}
	}
	flushItem()
	return issues
}
