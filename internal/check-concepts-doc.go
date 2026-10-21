//go:build ignore

// Command check-concepts-doc validates the structure and local links of the
// concepts guide. Run it from the repository root with:
//
//	go run internal/check-concepts-doc.go
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/fs"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

const (
	guidePath       = "CONCEPTS.md"
	linkTargetsPath = "CONCEPTS_LINKS.md"
	guideLineWidth  = 70
)

var (
	headingPattern = regexp.MustCompile(`^(#{1,6})[ \t]+(.+?)[ \t]*#*[ \t]*$`)
	anchorPattern  = regexp.MustCompile(`(?s)<a\s+(?:id|name)\s*=\s*"?([A-Za-z0-9_-]+)"?[^>]*>`)
	linkPattern    = regexp.MustCompile(`(?s)!?\[[^]]*\]\([ \t\r\n]*([^\s)]+)(?:[ \t\r\n]+["'][^)]*)?\)`)
	inboundPattern = regexp.MustCompile(`CONCEPTS\.md#([A-Za-z0-9_-]+)`)
	tablePattern   = regexp.MustCompile(`^[ \t]*\|.*\|[ \t]*$`)
	tableRule      = regexp.MustCompile(`^[ \t]*\|?[ \t]*:?-{3,}:?[ \t]*(?:\|[ \t]*:?-{3,}:?[ \t]*)+\|?[ \t]*$`)
)

var requiredFields = []string{
	"The problem it solves",
	"The idea in simple terms",
	"How it works",
	"Step-by-step examples",
	"Why it works",
	"Go example",
	"Work and memory",
	"When to use it",
	"When not to use it",
	"Common mistakes and edge cases",
	"Related material",
}

var baselineExternalLinks = []string{
	"https://careers.microsoft.com/v2/global/en/hiring-tips/technical-interviewing.html",
	"https://d3no4ktch0fdq4.cloudfront.net/public/course/files/Meta_SWE_full_loop_guide.pdf",
	"https://leetcode.com/discuss/post/6460774/Meta-or-E4-SWE-or-USA-or-Feb-2025-or-Passed-Hiring-Committee/",
	"https://leetcode.com/discuss/post/6522571/",
	"https://leetcode.com/discuss/post/7227003/",
	"https://leetcode.com/discuss/post/7327020/",
	"https://leetcode.com/discuss/post/8420230/amazon-sde-2-interview-experience-2026-e-t1g8/",
	"https://leetcode.com/studyplan/top-interview-150/",
	"https://neetcode.io/roadmap",
	"https://www.amazon.jobs/content/en/how-we-hire/interview-prep/software-development-topics",
	"https://www.techinterviewhandbook.org/grind75/faq",
}

type document struct {
	path            string
	contents        []byte
	prose           []byte
	targets         map[string]struct{}
	headingCount    int
	fenceCount      int
	goBlockCount    int
	textBlockCount  int
	explicitAnchors int
}

type linkCounts struct {
	internal int
	local    int
	external int
	inbound  int
}

func main() {
	repositoryRoot, err := os.Getwd()
	if err != nil {
		fatal(err)
	}

	guide, issues := readDocument(filepath.Join(repositoryRoot, guidePath))
	linkTargets, linkTargetIssues := readDocument(
		filepath.Join(repositoryRoot, linkTargetsPath),
	)
	issues = append(issues, linkTargetIssues...)
	if guide != nil {
		issues = append(issues, validateLineWidth(guidePath, guide.contents)...)
		issues = append(issues, validateTeachingEntries(guide.prose)...)
		issues = append(issues, validateNoTables(guide.prose)...)
	}

	counts := linkCounts{}
	externalLinks := make(map[string]struct{})
	if guide != nil {
		linkIssues, linkTotals, guideExternalLinks := validateDocumentLinks(guide)
		issues = append(issues, linkIssues...)
		counts = linkTotals
		for destination := range guideExternalLinks {
			externalLinks[destination] = struct{}{}
		}

		inboundIssues, inboundCount := validateInboundLinks(repositoryRoot, guide.targets)
		issues = append(issues, inboundIssues...)
		counts.inbound = inboundCount
	}
	if linkTargets != nil {
		linkIssues, linkTotals, targetExternalLinks := validateDocumentLinks(linkTargets)
		issues = append(issues, linkIssues...)
		counts.internal += linkTotals.internal
		counts.local += linkTotals.local
		counts.external += linkTotals.external
		for destination := range targetExternalLinks {
			externalLinks[destination] = struct{}{}
		}
	}
	for _, baselineLink := range baselineExternalLinks {
		if _, exists := externalLinks[baselineLink]; !exists {
			issues = append(issues, fmt.Sprintf("missing baseline external link %s", baselineLink))
		}
	}

	sort.Strings(issues)
	for _, issue := range issues {
		fmt.Fprintln(os.Stderr, issue)
	}
	if len(issues) > 0 {
		os.Exit(1)
	}

	fmt.Printf("headings: %d\n", guide.headingCount)
	fmt.Printf("explicit anchors: %d\n", guide.explicitAnchors)
	fmt.Printf("fenced blocks: %d (%d go, %d text)\n", guide.fenceCount, guide.goBlockCount, guide.textBlockCount)
	fmt.Printf("links: %d internal, %d local files, %d external, %d repository inbound\n", counts.internal, counts.local, counts.external, counts.inbound)
	fmt.Println("teaching entries: 93 with 11 ordered fields each")
}

func readDocument(path string) (*document, []string) {
	contents, err := os.ReadFile(path)
	if err != nil {
		return nil, []string{err.Error()}
	}

	prose, fences, goBlocks, textBlocks, fenceIssues := proseOutsideFences(path, contents)
	targets, headings, anchors, targetIssues := collectTargets(path, prose)
	return &document{
		path:            path,
		contents:        contents,
		prose:           prose,
		targets:         targets,
		headingCount:    headings,
		fenceCount:      fences,
		goBlockCount:    goBlocks,
		textBlockCount:  textBlocks,
		explicitAnchors: anchors,
	}, append(fenceIssues, targetIssues...)
}

func validateLineWidth(path string, contents []byte) []string {
	issues := make([]string, 0)
	scanner := bufio.NewScanner(bytes.NewReader(contents))
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		width := utf8.RuneCountInString(scanner.Text())
		if width > guideLineWidth {
			issues = append(
				issues,
				fmt.Sprintf(
					"%s:%d: line has %d characters; maximum is %d",
					path,
					lineNumber,
					width,
					guideLineWidth,
				),
			)
		}
	}
	if err := scanner.Err(); err != nil {
		issues = append(issues, fmt.Sprintf("%s: line-width scan: %v", path, err))
	}
	return issues
}

func proseOutsideFences(path string, contents []byte) ([]byte, int, int, int, []string) {
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
		issues = append(issues, fmt.Sprintf("%s: scan: %v", path, err))
	}
	if inFence {
		issues = append(issues, fmt.Sprintf("%s:%d: unclosed fenced code block", path, fenceStart))
	}
	return prose.Bytes(), blocks, goBlocks, textBlocks, issues
}

func collectTargets(path string, prose []byte) (map[string]struct{}, int, int, []string) {
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
				issues = append(issues, fmt.Sprintf("%s:%d: heading jumps from level %d to %d", path, lineNumber, previousLevel, level))
			}
			previousLevel = level
			base := githubSlug(match[2])
			slug := base
			if count := slugUses[base]; count > 0 {
				slug = fmt.Sprintf("%s-%d", base, count)
			}
			slugUses[base]++
			if origin, exists := targetOrigins[slug]; exists {
				issues = append(issues, fmt.Sprintf("%s:%d: generated target %q duplicates %s", path, lineNumber, slug, origin))
			}
			targets[slug] = struct{}{}
			targetOrigins[slug] = fmt.Sprintf("heading at line %d", lineNumber)
			headings++
		}
	}
	// Scan the complete prose so valid HTML anchors may span physical lines.
	for _, match := range anchorPattern.FindAllSubmatchIndex(prose, -1) {
		anchor := string(prose[match[2]:match[3]])
		lineNumber := bytes.Count(prose[:match[0]], []byte{'\n'}) + 1
		explicitUses[anchor]++
		if explicitUses[anchor] > 1 {
			issues = append(issues, fmt.Sprintf("%s:%d: duplicate explicit anchor %q", path, lineNumber, anchor))
		}
		if origin, exists := targetOrigins[anchor]; exists {
			issues = append(issues, fmt.Sprintf("%s:%d: explicit anchor %q duplicates %s", path, lineNumber, anchor, origin))
		}
		targets[anchor] = struct{}{}
		targetOrigins[anchor] = fmt.Sprintf("explicit anchor at line %d", lineNumber)
		anchors++
	}
	return targets, headings, anchors, issues
}

func githubSlug(heading string) string {
	var slug strings.Builder
	for _, current := range strings.ToLower(heading) {
		switch {
		case unicode.IsLetter(current), unicode.IsDigit(current), current == '-', current == '_':
			slug.WriteRune(current)
		case unicode.IsSpace(current):
			slug.WriteByte('-')
		}
	}
	return strings.Trim(slug.String(), "-")
}

func validateTeachingEntries(prose []byte) []string {
	fieldIndex := make(map[string]int, len(requiredFields))
	for index, field := range requiredFields {
		fieldIndex[field] = index
	}

	issues := make([]string, 0)
	expected := 0
	entryCount := 0
	entryLevel := 0
	scanner := bufio.NewScanner(bytes.NewReader(prose))
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		match := headingPattern.FindStringSubmatch(scanner.Text())
		if match == nil {
			continue
		}
		index, isField := fieldIndex[match[2]]
		if !isField {
			continue
		}
		if index != expected {
			issues = append(issues, fmt.Sprintf("%s:%d: teaching field %q appears where %q is required", guidePath, lineNumber, match[2], requiredFields[expected]))
			expected = (index + 1) % len(requiredFields)
			continue
		}
		level := len(match[1])
		if expected == 0 {
			entryLevel = level
		} else if level != entryLevel {
			issues = append(issues, fmt.Sprintf("%s:%d: teaching field %q uses level %d; entry started at level %d", guidePath, lineNumber, match[2], level, entryLevel))
		}
		expected++
		if expected == len(requiredFields) {
			entryCount++
			expected = 0
		}
	}
	if expected != 0 {
		issues = append(issues, fmt.Sprintf("%s: final teaching entry stops before %q", guidePath, requiredFields[expected]))
	}
	if entryCount != 93 {
		issues = append(issues, fmt.Sprintf("%s: teaching entries: got %d, want 93", guidePath, entryCount))
	}
	return issues
}

func validateNoTables(prose []byte) []string {
	lines := strings.Split(string(prose), "\n")
	for index := 1; index < len(lines); index++ {
		if tablePattern.MatchString(lines[index-1]) && tableRule.MatchString(lines[index]) {
			return []string{fmt.Sprintf("%s:%d: Markdown tables are not allowed in teaching entries", guidePath, index)}
		}
	}
	return nil
}

func validateDocumentLinks(document *document) ([]string, linkCounts, map[string]struct{}) {
	issues := make([]string, 0)
	counts := linkCounts{}
	externalLinks := make(map[string]struct{})
	displayPath := filepath.Base(document.path)
	for _, match := range linkPattern.FindAllSubmatch(withoutInlineCode(document.prose), -1) {
		destination := strings.Trim(string(match[1]), "<>")
		switch {
		case strings.HasPrefix(destination, "http://"), strings.HasPrefix(destination, "https://"):
			counts.external++
			externalLinks[destination] = struct{}{}
			continue
		case strings.HasPrefix(destination, "mailto:"):
			continue
		}

		pathPart, fragment, _ := strings.Cut(destination, "#")
		decodedFragment, err := url.PathUnescape(fragment)
		if err != nil {
			issues = append(issues, fmt.Sprintf("%s: invalid escaped fragment %q: %v", displayPath, fragment, err))
			continue
		}
		if pathPart == "" {
			counts.internal++
			if _, exists := document.targets[decodedFragment]; fragment != "" && !exists {
				issues = append(issues, fmt.Sprintf("%s: unresolved fragment #%s", displayPath, decodedFragment))
			}
			continue
		}

		counts.local++
		decodedPath, err := url.PathUnescape(pathPart)
		if err != nil {
			issues = append(issues, fmt.Sprintf("%s: invalid escaped path %q: %v", displayPath, pathPart, err))
			continue
		}
		targetPath := filepath.Clean(filepath.Join(filepath.Dir(document.path), filepath.FromSlash(decodedPath)))
		if _, err := os.Stat(targetPath); err != nil {
			issues = append(issues, fmt.Sprintf("%s: unresolved local path %s: %v", displayPath, decodedPath, err))
			continue
		}
		if fragment != "" && strings.EqualFold(filepath.Ext(targetPath), ".md") {
			targetDocument, targetIssues := readDocument(targetPath)
			issues = append(issues, targetIssues...)
			if targetDocument != nil {
				if _, exists := targetDocument.targets[decodedFragment]; !exists {
					issues = append(issues, fmt.Sprintf("%s: unresolved fragment %s#%s", displayPath, decodedPath, decodedFragment))
				}
			}
		}
	}
	return issues, counts, externalLinks
}

func withoutInlineCode(prose []byte) []byte {
	result := append([]byte(nil), prose...)
	delimiterLength := 0
	for index := 0; index < len(result); {
		if result[index] == '`' {
			runStart := index
			for index < len(result) && result[index] == '`' {
				index++
			}
			runLength := index - runStart
			if delimiterLength == 0 {
				delimiterLength = runLength
			} else if runLength == delimiterLength {
				delimiterLength = 0
			}
			for position := runStart; position < index; position++ {
				result[position] = ' '
			}
			continue
		}
		if delimiterLength > 0 && result[index] != '\n' {
			result[index] = ' '
		}
		index++
	}
	return result
}

func validateInboundLinks(repositoryRoot string, guideTargets map[string]struct{}) ([]string, int) {
	issues := make([]string, 0)
	count := 0
	err := filepath.WalkDir(repositoryRoot, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() {
			if entry.Name() == ".git" {
				return filepath.SkipDir
			}
			return nil
		}
		if path == filepath.Join(repositoryRoot, guidePath) {
			return nil
		}
		contents, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		if bytes.IndexByte(contents, 0) >= 0 {
			return nil
		}
		for _, match := range inboundPattern.FindAllSubmatch(contents, -1) {
			count++
			fragment := string(match[1])
			if _, exists := guideTargets[fragment]; !exists {
				relativePath, relErr := filepath.Rel(repositoryRoot, path)
				if relErr != nil {
					relativePath = path
				}
				issues = append(issues, fmt.Sprintf("%s: unresolved inbound concepts fragment #%s", relativePath, fragment))
			}
		}
		return nil
	})
	if err != nil {
		issues = append(issues, fmt.Sprintf("scan repository links: %v", err))
	}
	return issues, count
}

func fatal(err error) {
	if err == nil {
		return
	}
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
