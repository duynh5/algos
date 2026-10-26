// Command linewidth checks line widths in repository text files.
package main

import (
	"bufio"
	"bytes"
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
	"AGENTS.md":                 {},
	"CODING_RULES.md":           {},
	"REFERENCE_LINKS.md":        {},
	"VALIDATION.md":             {},
	"dsa/CONCEPTS_STANDARDS.md": {},
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
	textFileCount := 0
	exemptFileCount := 0
	for _, path := range paths {
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
		if isExemptPath(path) {
			exemptFileCount++
			continue
		}
		textFileCount++
		issues = append(issues, validateContents(path, contents)...)
	}

	sort.Strings(issues)
	for _, issue := range issues {
		fmt.Fprintln(os.Stderr, issue)
	}
	if len(issues) > 0 {
		os.Exit(1)
	}

	fmt.Printf(
		"line width: %d repository text files checked; maximum %d\n",
		textFileCount,
		maximumLineWidth,
	)
	fmt.Printf("line width exemptions: %d files\n", exemptFileCount)
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
	paths := make([]string, 0, len(encodedPaths))
	for _, encodedPath := range encodedPaths {
		paths = append(paths, string(encodedPath))
	}
	return paths, nil
}

func validateContents(path string, contents []byte) []string {
	issues := make([]string, 0)
	scanner := bufio.NewScanner(bytes.NewReader(contents))
	scanner.Buffer(make([]byte, 64*1024), 1024*1024)
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
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

func isExemptPath(path string) bool {
	_, exempt := exemptPaths[filepath.ToSlash(path)]
	return exempt
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, strings.TrimSpace(err.Error()))
	os.Exit(1)
}
