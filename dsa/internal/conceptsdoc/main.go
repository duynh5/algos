// Command conceptsdoc validates the structure and local links
// of every active concept study guide. Run it from the repository
// root with:
//
// go run ./dsa/internal/conceptsdoc
package main

import (
	"bytes"
	"fmt"
	"io/fs"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

const (
	linkTargetsPath = "REFERENCE_LINKS.md"
)

var (
	headingPattern = regexp.MustCompile(
		`^(#{1,6})[ \t]+(.+?)[ \t]*#*[ \t]*$`,
	)
	anchorPattern = regexp.MustCompile(
		`(?s)<a\s+(?:id|name)\s*=\s*"?([A-Za-z0-9_-` +
			`]+)"?[^>]*>`,
	)
	linkPattern = regexp.MustCompile(
		`(?s)!?\[[^]]*\]\([ \t\r\n]*([^\s)]+)(?:[ ` +
			`\t\r\n]+["'][^)]*)?[ \t\r\n]*\)`,
	)
	inboundPattern = regexp.MustCompile(
		`([A-Za-z0-9_./-]*CONCEPTS\.md)#([A-Za-z0-9_-]+)`,
	)
	tablePattern = regexp.MustCompile(`^[ \t]*\|.*\|[ \t]*$`)
	tableRule    = regexp.MustCompile(
		`^[ \t]*\|?[ \t]*:?-{3,}:?[ \t]*(?:\|[ \t]*:?-` +
			`{3,}:?[ \t]*)+\|?[ \t]*$`,
	)
	traceItemPattern = regexp.MustCompile(
		`^[ \t]*[0-9]+\.[ \t]+`,
	)
)

var conceptGuideSpecs = []conceptGuideSpec{
	{
		path:              "dsa/CONCEPTS.md",
		requiredFields:    teachingFields("Go example"),
		expectedEntries:   93,
		ownsBaselineLinks: true,
	},
	{
		path: "frontend/CONCEPTS.md",
		requiredFields: teachingFields(
			"TypeScript example",
		),
		expectedEntries: 13,
	},
	{
		path:            "backend/CONCEPTS.md",
		requiredFields:  teachingFields("Go example"),
		expectedEntries: 32,
	},
}

var baselineExternalLinks = []string{
	"https://careers.microsoft.com/v2/global/en/" +
		"hiring-tips/technical-interviewing.html",
	"https://d3no4ktch0fdq4.cloudfront.net/public/" +
		"course/files/Meta_SWE_full_loop_guide.pdf",
	"https://leetcode.com/discuss/post/6460774/" +
		"Meta-or-E4-SWE-or-USA-or-Feb-2025-or-Passed-" +
		"Hiring-Committee/",
	"https://leetcode.com/discuss/post/6522571/",
	"https://leetcode.com/discuss/post/7227003/",
	"https://leetcode.com/discuss/post/7327020/",
	"https://leetcode.com/discuss/post/8420230/" +
		"amazon-sde-2-interview-experience-2026-e-" +
		"t1g8/",
	"https://leetcode.com/studyplan/top-interview-150/",
	"https://neetcode.io/roadmap",
	"https://www.amazon.jobs/content/en/how-we-" +
		"hire/interview-prep/software-development-" +
		"topics",
	"https://www.techinterviewhandbook.org/grind75/faq",
}

type document struct {
	path            string
	prose           []byte
	targets         map[string]struct{}
	headingCount    int
	fenceCount      int
	goBlockCount    int
	textBlockCount  int
	explicitAnchors int
}

type conceptGuideSpec struct {
	path              string
	requiredFields    []string
	expectedEntries   int
	ownsBaselineLinks bool
}

type conceptGuideReport struct {
	spec       conceptGuideSpec
	document   *document
	entryCount int
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

	issues := make([]string, 0)
	counts := linkCounts{}
	externalLinks := make(map[string]struct{})
	reports := make([]conceptGuideReport, 0, len(conceptGuideSpecs))
	for _, spec := range conceptGuideSpecs {
		guide, guideIssues := readDocument(
			filepath.Join(repositoryRoot, spec.path),
		)
		issues = append(issues, guideIssues...)
		if guide == nil {
			continue
		}

		entryIssues, entryCount := validateTeachingEntries(
			spec,
			guide.prose,
		)
		issues = append(issues, entryIssues...)
		issues = append(
			issues,
			validateNoTables(spec.path, guide.prose)...,
		)
		issues = append(
			issues,
			validateTraceItems(spec.path, guide.prose)...,
		)
		linkIssues, linkTotals, guideExternalLinks :=
			validateDocumentLinks(guide)
		issues = append(issues, linkIssues...)
		counts.internal += linkTotals.internal
		counts.local += linkTotals.local
		counts.external += linkTotals.external
		if spec.ownsBaselineLinks {
			for destination := range guideExternalLinks {
				externalLinks[destination] = struct{}{}
			}
		}
		reports = append(reports, conceptGuideReport{
			spec:       spec,
			document:   guide,
			entryCount: entryCount,
		})
	}

	linkTargets, linkTargetIssues := readDocument(
		filepath.Join(repositoryRoot, linkTargetsPath),
	)
	issues = append(issues, linkTargetIssues...)
	if linkTargets != nil {
		linkIssues, linkTotals, targetExternalLinks :=
			validateDocumentLinks(linkTargets)
		issues = append(issues, linkIssues...)
		counts.internal += linkTotals.internal
		counts.local += linkTotals.local
		counts.external += linkTotals.external
		for destination := range targetExternalLinks {
			externalLinks[destination] = struct{}{}
		}
	}
	inboundIssues, inboundCount := validateInboundLinks(repositoryRoot)
	issues = append(issues, inboundIssues...)
	counts.inbound = inboundCount
	for _, baselineLink := range baselineExternalLinks {
		if _, exists := externalLinks[baselineLink]; !exists {
			issues = append(
				issues,
				fmt.Sprintf(
					"missing baseline external link %s",
					baselineLink,
				),
			)
		}
	}

	sort.Strings(issues)
	for _, issue := range issues {
		fmt.Fprintln(os.Stderr, issue)
	}
	if len(issues) > 0 {
		os.Exit(1)
	}

	for _, report := range reports {
		fmt.Printf(
			"%s: %d headings, %d anchors, %d fenced blocks, "+
				"%d teaching entries\n",
			report.spec.path,
			report.document.headingCount,
			report.document.explicitAnchors,
			report.document.fenceCount,
			report.entryCount,
		)
	}
	fmt.Printf(
		"links: %d internal, %d local files, %d "+
			"external, %d repository inbound\n",
		counts.internal,
		counts.local,
		counts.external,
		counts.inbound,
	)
}

func teachingFields(exampleHeading string) []string {
	return []string{
		"The problem it solves",
		"How it works",
		"Step-by-step examples",
		"Why it works",
		exampleHeading,
		"Work and memory",
		"When to use it",
		"When not to use it",
		"Common mistakes and edge cases",
		"Related material",
	}
}

func validateDocumentLinks(
	document *document,
) ([]string, linkCounts, map[string]struct{}) {
	issues := make([]string, 0)
	counts := linkCounts{}
	externalLinks := make(map[string]struct{})
	displayPath := filepath.Base(document.path)
	proseWithoutCode := withoutInlineCode(document.prose)
	for _, match := range linkPattern.FindAllSubmatch(
		proseWithoutCode,
		-1,
	) {
		destination := strings.Trim(string(match[1]), "<>")
		switch {
		case strings.HasPrefix(destination, "http://"),
			strings.HasPrefix(destination, "https://"):
			counts.external++
			externalLinks[destination] = struct{}{}
			continue
		case strings.HasPrefix(destination, "mailto:"):
			continue
		}

		pathPart, fragment, _ := strings.Cut(destination, "#")
		decodedFragment, err := url.PathUnescape(fragment)
		if err != nil {
			issues = append(
				issues,
				fmt.Sprintf(
					"%s: invalid escaped fragment %q: %v",
					displayPath,
					fragment,
					err,
				),
			)
			continue
		}
		if pathPart == "" {
			counts.internal++
			_, fragmentExists := document.targets[decodedFragment]
			if fragment != "" && !fragmentExists {
				issues = append(
					issues,
					fmt.Sprintf(
						"%s: unresolved fragment #%s",
						displayPath,
						decodedFragment,
					),
				)
			}
			continue
		}

		counts.local++
		decodedPath, err := url.PathUnescape(pathPart)
		if err != nil {
			issues = append(
				issues,
				fmt.Sprintf(
					"%s: invalid escaped path %q: %v",
					displayPath,
					pathPart,
					err,
				),
			)
			continue
		}
		targetPath := filepath.Clean(
			filepath.Join(
				filepath.Dir(document.path),
				filepath.FromSlash(decodedPath),
			),
		)
		if _, err := os.Stat(targetPath); err != nil {
			issues = append(
				issues,
				fmt.Sprintf(
					"%s: unresolved local path %s: %v",
					displayPath,
					decodedPath,
					err,
				),
			)
			continue
		}
		if fragment != "" &&
			strings.EqualFold(filepath.Ext(targetPath), ".md") {
			targetDocument, targetIssues := readDocument(targetPath)
			issues = append(issues, targetIssues...)
			if targetDocument != nil {
				if _, exists := targetDocument.targets[decodedFragment]; !exists {
					issues = append(
						issues,
						fmt.Sprintf(
							"%s: unresolved fragment %s#%s",
							displayPath,
							decodedPath,
							decodedFragment,
						),
					)
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

func validateInboundLinks(
	repositoryRoot string,
) ([]string, int) {
	issues := make([]string, 0)
	count := 0
	targetsByGuidePath := make(map[string]map[string]struct{})
	err := filepath.WalkDir(
		repositoryRoot,
		func(path string, entry fs.DirEntry, walkErr error) error {
			if walkErr != nil {
				return walkErr
			}
			if entry.IsDir() {
				if entry.Name() == ".git" {
					return filepath.SkipDir
				}
				return nil
			}
			if isActiveConceptGuide(repositoryRoot, path) {
				return nil
			}
			contents, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			if bytes.IndexByte(contents, 0) >= 0 {
				return nil
			}
			matches := inboundPattern.FindAllSubmatch(contents, -1)
			for _, match := range matches {
				count++
				destination := string(match[1])
				fragment := string(match[2])
				guideFilePath := filepath.Clean(
					filepath.Join(
						filepath.Dir(path),
						filepath.FromSlash(destination),
					),
				)

				// Cache each guide's generated and explicit targets so
				// repeated inbound links do not reread the same file.
				guideTargets, exists := targetsByGuidePath[guideFilePath]
				if !exists {
					guideDocument, guideIssues := readDocument(
						guideFilePath,
					)
					issues = append(issues, guideIssues...)
					if guideDocument == nil {
						continue
					}
					guideTargets = guideDocument.targets
					targetsByGuidePath[guideFilePath] = guideTargets
				}

				if _, exists := guideTargets[fragment]; !exists {
					relativePath, relErr := filepath.Rel(
						repositoryRoot,
						path,
					)
					if relErr != nil {
						relativePath = path
					}
					issues = append(
						issues,
						fmt.Sprintf(
							"%s: unresolved inbound concepts fragment "+
								"%s#%s",
							relativePath,
							destination,
							fragment,
						),
					)
				}
			}
			return nil
		},
	)
	if err != nil {
		issues = append(
			issues,
			fmt.Sprintf("scan repository links: %v", err),
		)
	}
	return issues, count
}

func isActiveConceptGuide(repositoryRoot string, path string) bool {
	for _, spec := range conceptGuideSpecs {
		if path == filepath.Join(repositoryRoot, spec.path) {
			return true
		}
	}
	return false
}

func fatal(err error) {
	if err == nil {
		return
	}
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
