// Command linewidth checks code and concept-guide line widths.
package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"unicode/utf8"
)

const maximumLineWidth = 70

var exemptPaths = map[string]struct{}{
	"AGENTS.md":             {},
	"CODING_RULES.md":       {},
	"CONCEPTS_STANDARDS.md": {},
	"README.md":             {},
	"REFERENCE_LINKS.md":    {},
	"VALIDATION.md":         {},
}

type validationScope int

const (
	skipFile validationScope = iota
	validateWholeFile
	validateCodeBlocks
)

var conceptGuidePaths = map[string]struct{}{
	"backend/CONCEPTS.md":  {},
	"dsa/CONCEPTS.md":      {},
	"frontend/CONCEPTS.md": {},
}

var codeExtensions = map[string]struct{}{
	".bash":  {},
	".c":     {},
	".cc":    {},
	".conf":  {},
	".cpp":   {},
	".cs":    {},
	".css":   {},
	".go":    {},
	".h":     {},
	".html":  {},
	".java":  {},
	".js":    {},
	".jsx":   {},
	".kt":    {},
	".kts":   {},
	".mod":   {},
	".php":   {},
	".proto": {},
	".py":    {},
	".rb":    {},
	".rs":    {},
	".scss":  {},
	".sh":    {},
	".sql":   {},
	".swift": {},
	".toml":  {},
	".ts":    {},
	".tsx":   {},
	".vue":   {},
	".xml":   {},
	".yaml":  {},
	".yml":   {},
}

var extensionlessCodeFiles = map[string]struct{}{
	"Dockerfile": {},
	"Makefile":   {},
}

func main() {
	repositoryRoot, err := os.Getwd()
	if err != nil {
		fatal(err)
	}

	paths, err := repositoryPaths(repositoryRoot)
	if err != nil {
		fatal(err)
	}

	issues := make([]string, 0)
	wholeFileCount := 0
	codeBlockFileCount := 0
	skippedFileCount := 0
	for _, path := range paths {
		scope := scopeForPath(path)
		if scope == skipFile {
			skippedFileCount++
			continue
		}
		contents, readErr := os.ReadFile(
			filepath.Join(repositoryRoot, path),
		)
		if readErr != nil {
			issues = append(issues, readErr.Error())
			continue
		}
		if bytes.IndexByte(contents, 0) >= 0 {
			continue
		}
		if scope == validateWholeFile {
			wholeFileCount++
			issues = append(
				issues,
				validateContents(path, contents)...,
			)
		} else {
			codeBlockFileCount++
			issues = append(
				issues,
				validateCodeBlockContents(path, contents)...,
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

	fmt.Printf(
		"line width: %d code or concept files checked; maximum %d\n",
		wholeFileCount,
		maximumLineWidth,
	)
	fmt.Printf(
		"line width: %d documentation files checked for code blocks\n",
		codeBlockFileCount,
	)
	fmt.Printf("line width: %d files outside scope\n", skippedFileCount)
}

func repositoryPaths(repositoryRoot string) ([]string, error) {
	command := exec.Command(
		"git",
		"ls-files",
		"-z",
		"--cached",
		"--others",
		"--exclude-standard",
	)
	command.Dir = repositoryRoot
	output, err := command.Output()
	if err != nil {
		return nil, fmt.Errorf("list tracked files: %w", err)
	}

	trimmedOutput := bytes.TrimSuffix(output, []byte{0})
	if len(trimmedOutput) == 0 {
		return nil, nil
	}

	encodedPaths := bytes.Split(trimmedOutput, []byte{0})
	return existingRepositoryPaths(repositoryRoot, encodedPaths)
}

func existingRepositoryPaths(
	repositoryRoot string,
	encodedPaths [][]byte,
) ([]string, error) {
	paths := make([]string, 0, len(encodedPaths))
	for _, encodedPath := range encodedPaths {
		path := string(encodedPath)
		_, err := os.Stat(filepath.Join(repositoryRoot, path))
		if errors.Is(err, os.ErrNotExist) {
			continue
		}
		if err != nil {
			return nil, fmt.Errorf("inspect repository path %s: %w", path, err)
		}
		paths = append(paths, path)
	}
	return paths, nil
}

func validateContents(path string, contents []byte) []string {
	return validateSelectedLines(path, contents, func(string) bool {
		return true
	})
}

func validateCodeBlockContents(
	path string,
	contents []byte,
) []string {
	inCodeBlock := false
	fenceMarker := ""
	return validateSelectedLines(
		path,
		contents,
		func(line string) bool {
			trimmedLine := strings.TrimSpace(line)
			if !inCodeBlock {
				fenceMarker = openingFenceMarker(trimmedLine)
				inCodeBlock = fenceMarker != ""
				return false
			}
			if strings.HasPrefix(trimmedLine, fenceMarker) {
				inCodeBlock = false
				fenceMarker = ""
				return false
			}
			return true
		},
	)
}

func validateSelectedLines(
	path string,
	contents []byte,
	shouldValidate func(string) bool,
) []string {
	issues := make([]string, 0)
	scanner := bufio.NewScanner(bytes.NewReader(contents))
	scanner.Buffer(make([]byte, 64*1024), 1024*1024)
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		if !shouldValidate(line) {
			continue
		}
		width := utf8.RuneCountInString(line)
		if width <= maximumLineWidth {
			continue
		}
		issues = append(
			issues,
			fmt.Sprintf(
				"%s:%d: line has %d characters; maximum is %d",
				path,
				lineNumber,
				width,
				maximumLineWidth,
			),
		)
	}
	if err := scanner.Err(); err != nil {
		issues = append(
			issues,
			fmt.Sprintf("%s: line-width scan: %v", path, err),
		)
	}
	return issues
}

func openingFenceMarker(line string) string {
	if strings.HasPrefix(line, "```") {
		return "```"
	}
	if strings.HasPrefix(line, "~~~") {
		return "~~~"
	}
	return ""
}

func scopeForPath(path string) validationScope {
	normalizedPath := filepath.ToSlash(path)
	_, isExempt := exemptPaths[normalizedPath]
	if isExempt {
		return skipFile
	}
	_, isConceptGuide := conceptGuidePaths[normalizedPath]
	if isConceptGuide {
		return validateWholeFile
	}
	if filepath.Ext(normalizedPath) == ".md" {
		return validateCodeBlocks
	}
	if isCodePath(normalizedPath) {
		return validateWholeFile
	}
	return skipFile
}

func isCodePath(path string) bool {
	base := filepath.Base(path)
	_, isExtensionlessCode := extensionlessCodeFiles[base]
	if isExtensionlessCode {
		return true
	}
	pathWithoutTemplate := strings.TrimSuffix(path, ".template")
	codeExtension := strings.ToLower(
		filepath.Ext(pathWithoutTemplate),
	)
	_, isCodeExtension := codeExtensions[codeExtension]
	return isCodeExtension
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, strings.TrimSpace(err.Error()))
	os.Exit(1)
}
