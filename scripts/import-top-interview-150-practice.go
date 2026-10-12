//go:build ignore

// Command import-top-interview-150-practice generates unsolved practice
// archives from the official LeetCode GraphQL metadata downloaded separately.
package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"html"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type catalogEntry struct {
	ID         string
	Difficulty string
	Slug       string
	Title      string
}

type graphQLResponse struct {
	Data struct {
		Question question `json:"question"`
	} `json:"data"`
	Errors []json.RawMessage `json:"errors"`
}

type question struct {
	QuestionFrontendID string `json:"questionFrontendId"`
	Title              string `json:"title"`
	TitleSlug          string `json:"titleSlug"`
	Difficulty         string `json:"difficulty"`
	Content            string `json:"content"`
	ExampleTestcases   string `json:"exampleTestcases"`
	MetaData           string `json:"metaData"`
}

type metadata struct {
	Name      string      `json:"name"`
	ClassName string      `json:"classname"`
	Params    []parameter `json:"params"`
	Return    parameter   `json:"return"`
	Output    *struct {
		ParamIndex int `json:"paramindex"`
	} `json:"output"`
	SystemDesign bool `json:"systemdesign"`
}

type parameter struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type practiceCase struct {
	Name  string
	Input string
	Want  string
}

var (
	prePattern             = regexp.MustCompile(`(?is)<pre[^>]*>(.*?)</pre>`)
	supPattern             = regexp.MustCompile(`(?is)<sup[^>]*>(.*?)</sup>`)
	subPattern             = regexp.MustCompile(`(?is)<sub[^>]*>(.*?)</sub>`)
	lineBreakPattern       = regexp.MustCompile(`(?is)<(?:br\s*/?|/p|/li|/ul|/ol|/h[1-6]|/pre)>`)
	listItemPattern        = regexp.MustCompile(`(?is)<li[^>]*>`)
	tagPattern             = regexp.MustCompile(`(?is)<[^>]+>`)
	spacePattern           = regexp.MustCompile(`[ \t]+`)
	blankPattern           = regexp.MustCompile(`\n[ \t]*\n(?:[ \t]*\n)+`)
	paragraphOutputPattern = regexp.MustCompile(`(?is)<p[^>]*>\s*<strong[^>]*>\s*Output:?\s*</strong>\s*(.*?)</p>`)
	outputPattern          = regexp.MustCompile(`(?m)^Output:?[ \t]*(?:\n)?`)
	outputStopPattern      = regexp.MustCompile(`(?m)^(?:Explanation|Note):?[ \t]*`)
)

func main() {
	manifestPath := flag.String("manifest", "scripts/top-interview-150.tsv", "official study-plan catalog")
	metadataDirectory := flag.String("metadata", "/tmp/leetcode-top-interview-150", "downloaded GraphQL metadata")
	problemDirectory := flag.String("problems", "problems", "archive destination")
	flag.Parse()

	entries, err := readCatalog(*manifestPath)
	if err != nil {
		fatal(err)
	}
	if len(entries) != 150 {
		fatal(fmt.Errorf("catalog contains %d problems; want 150", len(entries)))
	}

	generated := make([]string, 0, len(entries))
	for _, entry := range entries {
		question, meta, err := readQuestion(*metadataDirectory, entry)
		if err != nil {
			fatal(err)
		}
		filename, err := destinationFilename(*problemDirectory, entry)
		if err != nil {
			fatal(err)
		}
		contents, err := renderArchive(entry, question, meta)
		if err != nil {
			fatal(fmt.Errorf("render %s: %w", entry.Slug, err))
		}
		if err := os.WriteFile(filepath.Join(*problemDirectory, filename), contents, 0o644); err != nil {
			fatal(fmt.Errorf("write %s: %w", filename, err))
		}
		generated = append(generated, filename)
	}

	if err := updateProblemIndex(filepath.Join(*problemDirectory, "README.md"), generated); err != nil {
		fatal(err)
	}
	fmt.Printf("generated %d unsolved practice archives\n", len(generated))
}

func readCatalog(path string) ([]catalogEntry, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open catalog: %w", err)
	}
	defer file.Close()

	seen := make(map[string]bool)
	var entries []catalogEntry
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		fields := strings.SplitN(line, "\t", 4)
		if len(fields) != 4 {
			return nil, fmt.Errorf("catalog line %q does not have four tab-separated fields", line)
		}
		entry := catalogEntry{ID: fields[0], Difficulty: fields[1], Slug: fields[2], Title: fields[3]}
		if seen[entry.ID] {
			return nil, fmt.Errorf("catalog repeats problem ID %s", entry.ID)
		}
		seen[entry.ID] = true
		entries = append(entries, entry)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scan catalog: %w", err)
	}
	return entries, nil
}

func readQuestion(directory string, entry catalogEntry) (question, metadata, error) {
	path := filepath.Join(directory, entry.ID+".json")
	contents, err := os.ReadFile(path)
	if err != nil {
		return question{}, metadata{}, fmt.Errorf("read metadata for %s: %w", entry.ID, err)
	}
	var response graphQLResponse
	if err := json.Unmarshal(contents, &response); err != nil {
		return question{}, metadata{}, fmt.Errorf("decode metadata for %s: %w", entry.ID, err)
	}
	if len(response.Errors) > 0 {
		return question{}, metadata{}, fmt.Errorf("GraphQL returned errors for %s: %s", entry.ID, response.Errors)
	}
	q := response.Data.Question
	if q.QuestionFrontendID != entry.ID || q.Title != entry.Title || q.TitleSlug != entry.Slug || !strings.EqualFold(q.Difficulty, entry.Difficulty) {
		return question{}, metadata{}, fmt.Errorf("official metadata disagrees with catalog for %s", entry.ID)
	}
	var meta metadata
	if err := json.Unmarshal([]byte(q.MetaData), &meta); err != nil {
		return question{}, metadata{}, fmt.Errorf("decode signature metadata for %s: %w", entry.ID, err)
	}
	return q, meta, nil
}

func destinationFilename(directory string, entry catalogEntry) (string, error) {
	problemID, err := strconv.Atoi(entry.ID)
	if err != nil {
		return "", fmt.Errorf("problem ID %q is not numeric: %w", entry.ID, err)
	}
	base := fmt.Sprintf("%s_leetcode_%04d_%s.go", entry.Difficulty, problemID, strings.ReplaceAll(entry.Slug, "-", "_"))
	practice := strings.TrimSuffix(base, ".go") + "_practice.go"
	practicePath := filepath.Join(directory, practice)
	if _, err := os.Stat(practicePath); err == nil {
		if isGeneratedPracticeArchive(practicePath) {
			return practice, nil
		}
		return "", fmt.Errorf("refusing to overwrite existing archive %s", practicePath)
	} else if !errors.Is(err, os.ErrNotExist) {
		return "", fmt.Errorf("inspect destination %s: %w", practice, err)
	}

	path := filepath.Join(directory, base)
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return base, nil
	} else if err != nil {
		return "", fmt.Errorf("inspect destination %s: %w", base, err)
	}
	if isGeneratedPracticeArchive(path) {
		return base, nil
	}
	return practice, nil
}

func isGeneratedPracticeArchive(path string) bool {
	contents, err := os.ReadFile(path)
	return err == nil && bytes.Contains(contents, []byte("intentionally unsolved practice archive"))
}

func renderArchive(entry catalogEntry, q question, meta metadata) ([]byte, error) {
	cases, err := buildCases(q.ExampleTestcases, q.Content, inputParameters(entry, meta))
	if err != nil {
		return nil, err
	}
	functionName := exportedIdentifier(entry.Slug)
	parameters := inputParameters(entry, meta)
	var output bytes.Buffer
	fmt.Fprintln(&output, "//go:build ignore")
	fmt.Fprintln(&output)
	fmt.Fprintf(&output, "// Difficulty: %s\n", strings.ToUpper(entry.Difficulty[:1])+entry.Difficulty[1:])
	fmt.Fprintln(&output, "// Source: leetcode")
	fmt.Fprintf(&output, "// Problem ID: %s\n\n", entry.ID)
	fmt.Fprintln(&output, "package main")
	fmt.Fprintln(&output)
	fmt.Fprintln(&output, "import (")
	fmt.Fprintln(&output, "\t\"encoding/json\"")
	fmt.Fprintln(&output, "\t\"fmt\"")
	fmt.Fprintln(&output, "\t\"os\"")
	fmt.Fprintln(&output)
	fmt.Fprintln(&output, "\t\"github.com/duyserveacc/algos/internal/testrunner\"")
	fmt.Fprintln(&output, ")")
	fmt.Fprintln(&output)
	fmt.Fprintln(&output, "/*")
	fmt.Fprintln(&output, "===============================================================================")
	fmt.Fprintf(&output, "%s\n\n", entry.Title)
	fmt.Fprintln(&output, "Source: LeetCode")
	fmt.Fprintf(&output, "Problem ID: %s\n", entry.ID)
	fmt.Fprintf(&output, "Source URL: https://leetcode.com/problems/%s/description/\n\n", entry.Slug)
	fmt.Fprintln(&output, cleanHTML(q.Content))
	fmt.Fprintln(&output, "===============================================================================")
	fmt.Fprintln(&output, "*/")
	fmt.Fprintln(&output)
	fmt.Fprintln(&output, "// Problem-specific support.")
	fmt.Fprintln(&output, "//")
	fmt.Fprintln(&output, "// Keep shared judge types and parsing or formatting helpers here. Name adapters")
	fmt.Fprintln(&output, "// after the problem. Put setup used only by the current attempt at the top of")
	fmt.Fprintln(&output, "// solveCurrentAttempt below.")
	fmt.Fprintln(&output, "//")
	fmt.Fprintln(&output, "// This intentionally unsolved practice archive follows the workbench template.")
	fmt.Fprintln(&output)
	writeSupportTypes(&output, entry, parameters)
	writeParser(&output, entry, functionName, parameters)
	writeCases(&output, cases)
	writeWorkbenchScaffold(&output, functionName, parameters)
	return output.Bytes(), nil
}

func inputParameters(entry catalogEntry, meta metadata) []parameter {
	if meta.SystemDesign {
		return []parameter{{Name: "operations", Type: "string[]"}, {Name: "arguments", Type: "operationArguments"}}
	}
	if entry.ID == "138" {
		return []parameter{{Name: "head", Type: "randomList"}}
	}
	return meta.Params
}

func buildCases(exampleStream, content string, parameters []parameter) ([]practiceCase, error) {
	if len(parameters) == 0 {
		return nil, errors.New("problem metadata has no input parameters")
	}
	decoder := json.NewDecoder(strings.NewReader(exampleStream))
	var values []json.RawMessage
	for {
		var value json.RawMessage
		if err := decoder.Decode(&value); errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return nil, fmt.Errorf("decode example input stream: %w", err)
		}
		values = append(values, append(json.RawMessage(nil), value...))
	}
	if len(values)%len(parameters) != 0 {
		return nil, fmt.Errorf("%d example values do not divide into %d parameters", len(values), len(parameters))
	}
	outputs := exampleOutputs(content)
	count := len(values) / len(parameters)
	if len(outputs) != count {
		return nil, fmt.Errorf("found %d example outputs for %d example inputs", len(outputs), count)
	}
	cases := make([]practiceCase, 0, count)
	for index := 0; index < count; index++ {
		group := values[index*len(parameters) : (index+1)*len(parameters)]
		input, err := json.Marshal(group)
		if err != nil {
			return nil, fmt.Errorf("encode example %d: %w", index+1, err)
		}
		cases = append(cases, practiceCase{Name: fmt.Sprintf("Official example %d", index+1), Input: string(input), Want: outputs[index]})
	}
	return cases, nil
}

func exampleOutputs(content string) []string {
	paragraphMatches := paragraphOutputPattern.FindAllStringSubmatch(content, -1)
	if len(paragraphMatches) > 0 {
		outputs := make([]string, 0, len(paragraphMatches))
		for _, match := range paragraphMatches {
			outputs = append(outputs, cleanHTML(match[1]))
		}
		return outputs
	}
	matches := prePattern.FindAllStringSubmatch(content, -1)
	outputs := make([]string, 0, len(matches))
	for _, match := range matches {
		text := cleanHTML(match[1])
		location := outputPattern.FindStringIndex(text)
		if location == nil {
			continue
		}
		value := strings.TrimSpace(text[location[1]:])
		if stop := outputStopPattern.FindStringIndex(value); stop != nil {
			value = strings.TrimSpace(value[:stop[0]])
		}
		outputs = append(outputs, value)
	}
	return outputs
}

func cleanHTML(source string) string {
	text := supPattern.ReplaceAllString(source, "^$1")
	text = subPattern.ReplaceAllString(text, "_$1")
	text = listItemPattern.ReplaceAllString(text, "\n- ")
	text = lineBreakPattern.ReplaceAllString(text, "\n")
	text = tagPattern.ReplaceAllString(text, "")
	text = html.UnescapeString(text)
	text = strings.NewReplacer(
		"\u00a0", " ",
		"\u200b", "",
		"\u2192", "->",
		"\u221e", "infinity",
		"\u230a", "floor(",
		"\u230b", ")",
		"\u201c", "\"",
		"\u201d", "\"",
	).Replace(text)
	lines := strings.Split(strings.ReplaceAll(text, "\r", ""), "\n")
	for index := range lines {
		lines[index] = strings.TrimSpace(spacePattern.ReplaceAllString(lines[index], " "))
	}
	text = strings.Join(lines, "\n")
	text = blankPattern.ReplaceAllString(text, "\n\n")
	return strings.TrimSpace(text)
}

func writeSupportTypes(output *bytes.Buffer, entry catalogEntry, parameters []parameter) {
	types := make(map[string]bool)
	for _, param := range parameters {
		types[param.Type] = true
	}
	if types["TreeNode"] {
		fmt.Fprintln(output, "type TreeNode struct {")
		fmt.Fprintln(output, "\tVal int")
		fmt.Fprintln(output, "\tLeft, Right, Next *TreeNode")
		fmt.Fprintln(output, "}")
		fmt.Fprintln(output)
	}
	if types["ListNode"] || types["ListNode[]"] {
		fmt.Fprintln(output, "type ListNode struct {")
		fmt.Fprintln(output, "\tVal int")
		fmt.Fprintln(output, "\tNext *ListNode")
		fmt.Fprintln(output, "}")
		fmt.Fprintln(output)
	}
	if types["randomList"] {
		fmt.Fprintln(output, "type randomListEntry struct {")
		fmt.Fprintln(output, "\tValue int")
		fmt.Fprintln(output, "\tRandomIndex *int")
		fmt.Fprintln(output, "}")
		fmt.Fprintln(output)
	}
	_ = entry
}

func writeCases(output *bytes.Buffer, cases []practiceCase) {
	fmt.Fprintln(output, "// sampleCases contains the examples that every registered solution must pass.")
	fmt.Fprintln(output, "// Inputs encode the judge arguments as a positional JSON array.")
	fmt.Fprintln(output, "var sampleCases = []testrunner.Case{")
	for _, practice := range cases {
		fmt.Fprintln(output, "\t{")
		fmt.Fprintf(output, "\t\tName: %s,\n", strconv.Quote(practice.Name))
		fmt.Fprintf(output, "\t\tInput: %s,\n", rawLiteral(practice.Input))
		fmt.Fprintf(output, "\t\tWant: %s,\n", rawLiteral(practice.Want))
		fmt.Fprintln(output, "\t},")
	}
	fmt.Fprintln(output, "}")
	fmt.Fprintln(output)
}

func writeParser(output *bytes.Buffer, entry catalogEntry, functionName string, parameters []parameter) {
	returnTypes := make([]string, len(parameters))
	for index, param := range parameters {
		returnTypes[index] = goType(param.Type)
	}
	fmt.Fprintf(output, "func parse%sInput(input string) (%s) {\n", functionName, strings.Join(returnTypes, ", "))
	fmt.Fprintln(output, "\tvar args []json.RawMessage")
	fmt.Fprintln(output, "\tif err := json.Unmarshal([]byte(input), &args); err != nil {")
	fmt.Fprintf(output, "\t\tpanic(fmt.Errorf(\"parse %s input: %%w\", err))\n", entry.Slug)
	fmt.Fprintln(output, "\t}")
	fmt.Fprintf(output, "\tif len(args) != %d {\n", len(parameters))
	fmt.Fprintf(output, "\t\tpanic(fmt.Sprintf(\"parse %s input: got %%d arguments, want %d\", len(args)))\n", entry.Slug, len(parameters))
	fmt.Fprintln(output, "\t}")
	names := parameterNames(parameters)
	for index, param := range parameters {
		name := names[index]
		writeDecodeParameter(output, entry, name, param.Type, index)
	}
	if len(parameters) == 2 && parameters[0].Type == "string[]" && parameters[1].Type == "operationArguments" {
		fmt.Fprintln(output, "\tif len(operations) != len(arguments) {")
		fmt.Fprintf(output, "\t\tpanic(fmt.Sprintf(\"parse %s input: got %%d operations and %%d argument lists\", len(operations), len(arguments)))\n", entry.Slug)
		fmt.Fprintln(output, "\t}")
	}
	if entry.ID == "141" {
		fmt.Fprintln(output, "\tif pos >= 0 {")
		fmt.Fprintln(output, "\t\ttail, target := head, head")
		fmt.Fprintln(output, "\t\tfor index := 0; index < pos; index++ { target = target.Next }")
		fmt.Fprintln(output, "\t\tfor tail.Next != nil { tail = tail.Next }")
		fmt.Fprintln(output, "\t\ttail.Next = target")
		fmt.Fprintln(output, "\t}")
	}
	fmt.Fprintf(output, "\treturn %s\n", strings.Join(names, ", "))
	fmt.Fprintln(output, "}")
	fmt.Fprintln(output)
	writeStructuralDecoders(output, parameters)
}

func writeDecodeParameter(output *bytes.Buffer, entry catalogEntry, name, problemType string, index int) {
	errorName := fmt.Sprintf("argumentError%d", index)
	switch problemType {
	case "TreeNode":
		fmt.Fprintf(output, "\t%s, %s := decodeTreeNode(args[%d])\n", name, errorName, index)
	case "ListNode":
		fmt.Fprintf(output, "\t%s, %s := decodeListNode(args[%d])\n", name, errorName, index)
	case "ListNode[]":
		fmt.Fprintf(output, "\t%s, %s := decodeListNodes(args[%d])\n", name, errorName, index)
	case "character[][]":
		fmt.Fprintf(output, "\t%s, %s := decodeCharacterMatrix(args[%d])\n", name, errorName, index)
	case "randomList":
		fmt.Fprintf(output, "\t%s, %s := decodeRandomList(args[%d])\n", name, errorName, index)
	default:
		fmt.Fprintf(output, "\tvar %s %s\n", name, goType(problemType))
		fmt.Fprintf(output, "\t%s := json.Unmarshal(args[%d], &%s)\n", errorName, index, name)
	}
	fmt.Fprintf(output, "\tif %s != nil {\n", errorName)
	fmt.Fprintf(output, "\t\tpanic(fmt.Errorf(\"parse %s argument %s: %%w\", %s))\n", entry.Slug, name, errorName)
	fmt.Fprintln(output, "\t}")
}
