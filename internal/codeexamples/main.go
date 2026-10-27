// Command codeexamples validates code fences embedded in Markdown.
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

const typeScriptXPrelude = `declare namespace React {
  function useState<Value>(initialValue: Value): [
    Value,
    (value: Value) => void,
  ];
  function useEffect(
    effect: () => void | (() => void),
    dependencies: readonly unknown[],
  ): void;
}

declare namespace JSX {
  interface IntrinsicElements {
    h2: Record<string, unknown>;
    p: Record<string, unknown>;
  }
}

declare function loadProduct(
  productID: string,
  signal: AbortSignal,
): Promise<Product>;
`

type codeBlock struct {
	path      string
	language  string
	startLine int
	source    string
}

func main() {
	repositoryRoot, err := os.Getwd()
	if err != nil {
		fatal(err)
	}

	markdownPaths, err := repositoryMarkdownPaths(repositoryRoot)
	if err != nil {
		fatal(err)
	}
	temporaryDirectory, err := os.MkdirTemp(
		"",
		"algos-code-examples-*",
	)
	if err != nil {
		fatal(fmt.Errorf("create temporary directory: %w", err))
	}

	// Parse every Markdown file before invoking language-specific tools
	// so malformed fences and invalid code share one sorted report.
	blocks := make([]codeBlock, 0)
	issues := make([]string, 0)
	for _, path := range markdownPaths {
		contents, readErr := os.ReadFile(
			filepath.Join(repositoryRoot, path),
		)
		if readErr != nil {
			issues = append(issues, readErr.Error())
			continue
		}
		fileBlocks, fileIssues := readCodeBlocks(path, contents)
		blocks = append(blocks, fileBlocks...)
		issues = append(issues, fileIssues...)
	}

	validatedCounts := make(map[string]int)
	for blockIndex, block := range blocks {
		blockIssues := validateCodeBlock(
			block,
			blockIndex,
			temporaryDirectory,
		)
		issues = append(issues, blockIssues...)
		validatedCounts[block.language]++
	}
	if err := os.RemoveAll(temporaryDirectory); err != nil {
		issues = append(
			issues,
			fmt.Sprintf("remove temporary examples: %v", err),
		)
	}

	sort.Strings(issues)
	for _, issue := range issues {
		fmt.Fprintln(os.Stderr, issue)
	}
	if len(issues) > 0 {
		os.Exit(1)
	}

	fmt.Printf(
		"code examples: %d Go, %d TypeScript, %d TSX\n",
		validatedCounts["go"],
		validatedCounts["typescript"],
		validatedCounts["tsx"],
	)
}

func repositoryMarkdownPaths(
	repositoryRoot string,
) ([]string, error) {
	paths := make([]string, 0)
	err := filepath.WalkDir(
		repositoryRoot,
		func(path string, entry fs.DirEntry, walkErr error) error {
			if walkErr != nil {
				return walkErr
			}
			if entry.IsDir() {
				switch entry.Name() {
				case ".git", "node_modules":
					return filepath.SkipDir
				default:
					return nil
				}
			}
			if !strings.EqualFold(filepath.Ext(path), ".md") {
				return nil
			}
			relativePath, err := filepath.Rel(repositoryRoot, path)
			if err != nil {
				return err
			}
			paths = append(paths, relativePath)
			return nil
		},
	)
	if err != nil {
		return nil, fmt.Errorf("list Markdown files: %w", err)
	}
	return paths, nil
}

func readCodeBlocks(
	path string,
	contents []byte,
) ([]codeBlock, []string) {
	blocks := make([]codeBlock, 0)
	issues := make([]string, 0)
	scanner := bufio.NewScanner(bytes.NewReader(contents))
	scanner.Buffer(make([]byte, 64*1024), 1024*1024)
	lineNumber := 0
	fenceStart := 0
	language := ""
	inFence := false
	var source strings.Builder

	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "```") {
			if !inFence {
				inFence = true
				fenceStart = lineNumber
				language = fenceLanguage(trimmed)
				source.Reset()
				continue
			}
			if validatesLanguage(language) {
				blocks = append(blocks, codeBlock{
					path:      path,
					language:  language,
					startLine: fenceStart + 1,
					source:    source.String(),
				})
			}
			inFence = false
			language = ""
			continue
		}
		if inFence && validatesLanguage(language) {
			source.WriteString(line)
			source.WriteByte('\n')
		}
	}

	if err := scanner.Err(); err != nil {
		issues = append(
			issues,
			fmt.Sprintf("%s: scan code blocks: %v", path, err),
		)
	}
	if inFence {
		issues = append(
			issues,
			fmt.Sprintf("%s:%d: unclosed code fence", path, fenceStart),
		)
	}
	return blocks, issues
}

func fenceLanguage(openingFence string) string {
	description := strings.TrimSpace(
		strings.TrimPrefix(openingFence, "```"),
	)
	fields := strings.Fields(description)
	if len(fields) == 0 {
		return ""
	}
	return strings.ToLower(fields[0])
}

func validatesLanguage(language string) bool {
	switch language {
	case "go", "typescript", "tsx":
		return true
	default:
		return false
	}
}

func validateCodeBlock(
	block codeBlock,
	blockIndex int,
	temporaryDirectory string,
) []string {
	location := fmt.Sprintf("%s:%d", block.path, block.startLine)
	switch block.language {
	case "go":
		return validateFormattedBlock(
			location,
			"gofmt",
			nil,
			block.source,
		)
	case "typescript", "tsx":
		issues := validateFormattedBlock(
			location,
			"prettier",
			[]string{
				"--stdin-filepath",
				"example." + typeScriptExtension(block.language),
			},
			block.source,
		)
		return append(
			issues,
			validateTypeScriptBlock(
				block,
				blockIndex,
				temporaryDirectory,
			)...,
		)
	default:
		return nil
	}
}

func validateFormattedBlock(
	location string,
	tool string,
	arguments []string,
	source string,
) []string {
	command := exec.Command(tool, arguments...)
	command.Stdin = strings.NewReader(source)
	var standardOutput bytes.Buffer
	var standardError bytes.Buffer
	command.Stdout = &standardOutput
	command.Stderr = &standardError
	if err := command.Run(); err != nil {
		return []string{fmt.Sprintf(
			"%s: %s failed: %v: %s",
			location,
			tool,
			err,
			strings.TrimSpace(standardError.String()),
		)}
	}
	if standardOutput.String() != source {
		return []string{fmt.Sprintf(
			"%s: code block is not formatted by %s",
			location,
			tool,
		)}
	}
	return nil
}

func validateTypeScriptBlock(
	block codeBlock,
	blockIndex int,
	temporaryDirectory string,
) []string {
	extension := typeScriptExtension(block.language)
	sourcePath := filepath.Join(
		temporaryDirectory,
		fmt.Sprintf("example-%03d.%s", blockIndex, extension),
	)
	if err := os.WriteFile(
		sourcePath,
		[]byte(block.source),
		0o600,
	); err != nil {
		return []string{fmt.Sprintf(
			"%s:%d: write temporary example: %v",
			block.path,
			block.startLine,
			err,
		)}
	}

	arguments := []string{
		"--noEmit",
		"--strict",
		"--target", "ES2022",
		"--lib", "ES2022,DOM",
		"--jsx", "preserve",
		"--pretty", "false",
	}
	if block.language == "tsx" {
		preludePath := filepath.Join(
			temporaryDirectory,
			"tsx-environment.d.ts",
		)
		if err := os.WriteFile(
			preludePath,
			[]byte(typeScriptXPrelude),
			0o600,
		); err != nil {
			return []string{fmt.Sprintf(
				"%s:%d: write TSX environment: %v",
				block.path,
				block.startLine,
				err,
			)}
		}
		arguments = append(arguments, preludePath)
	}
	arguments = append(arguments, sourcePath)

	command := exec.Command("tsc", arguments...)
	output, err := command.CombinedOutput()
	if err == nil {
		return nil
	}
	cleanOutput := strings.ReplaceAll(
		strings.TrimSpace(string(output)),
		temporaryDirectory,
		"<temporary>",
	)
	return []string{fmt.Sprintf(
		"%s:%d: TypeScript check failed: %v: %s",
		block.path,
		block.startLine,
		err,
		cleanOutput,
	)}
}

func typeScriptExtension(language string) string {
	if language == "tsx" {
		return "tsx"
	}
	return "ts"
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, strings.TrimSpace(err.Error()))
	os.Exit(1)
}
