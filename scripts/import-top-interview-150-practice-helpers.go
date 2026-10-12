//go:build ignore

package main

import (
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func writeStructuralDecoders(output *bytes.Buffer, parameters []parameter) {
	types := make(map[string]bool)
	for _, param := range parameters {
		types[param.Type] = true
	}
	if types["TreeNode"] {
		fmt.Fprintln(output, "func decodeTreeNode(raw json.RawMessage) (*TreeNode, error) {")
		fmt.Fprintln(output, "\tvar values []*int")
		fmt.Fprintln(output, "\tif err := json.Unmarshal(raw, &values); err != nil || len(values) == 0 { return nil, err }")
		fmt.Fprintln(output, "\tif values[0] == nil { return nil, nil }")
		fmt.Fprintln(output, "\troot := &TreeNode{Val: *values[0]}; queue := []*TreeNode{root}; next := 1")
		fmt.Fprintln(output, "\tfor len(queue) > 0 && next < len(values) {")
		fmt.Fprintln(output, "\t\tnode := queue[0]; queue = queue[1:]")
		fmt.Fprintln(output, "\t\tif values[next] != nil { node.Left = &TreeNode{Val: *values[next]}; queue = append(queue, node.Left) }; next++")
		fmt.Fprintln(output, "\t\tif next < len(values) { if values[next] != nil { node.Right = &TreeNode{Val: *values[next]}; queue = append(queue, node.Right) }; next++ }")
		fmt.Fprintln(output, "\t}")
		fmt.Fprintln(output, "\treturn root, nil")
		fmt.Fprintln(output, "}")
		fmt.Fprintln(output)
	}
	if types["ListNode"] || types["ListNode[]"] {
		fmt.Fprintln(output, "func decodeListNode(raw json.RawMessage) (*ListNode, error) {")
		fmt.Fprintln(output, "\tvar values []int; if err := json.Unmarshal(raw, &values); err != nil { return nil, err }")
		fmt.Fprintln(output, "\tdummy := &ListNode{}; tail := dummy")
		fmt.Fprintln(output, "\tfor _, value := range values { tail.Next = &ListNode{Val: value}; tail = tail.Next }")
		fmt.Fprintln(output, "\treturn dummy.Next, nil")
		fmt.Fprintln(output, "}")
		fmt.Fprintln(output)
	}
	if types["ListNode[]"] {
		fmt.Fprintln(output, "func decodeListNodes(raw json.RawMessage) ([]*ListNode, error) {")
		fmt.Fprintln(output, "\tvar lists []json.RawMessage; if err := json.Unmarshal(raw, &lists); err != nil { return nil, err }")
		fmt.Fprintln(output, "\tresult := make([]*ListNode, len(lists)); for index := range lists { head, err := decodeListNode(lists[index]); if err != nil { return nil, err }; result[index] = head }")
		fmt.Fprintln(output, "\treturn result, nil")
		fmt.Fprintln(output, "}")
		fmt.Fprintln(output)
	}
	if types["character[][]"] {
		fmt.Fprintln(output, "func decodeCharacterMatrix(raw json.RawMessage) ([][]byte, error) {")
		fmt.Fprintln(output, "\tvar rows [][]string; if err := json.Unmarshal(raw, &rows); err != nil { return nil, err }")
		fmt.Fprintln(output, "\tresult := make([][]byte, len(rows)); for row := range rows { result[row] = make([]byte, len(rows[row])); for column, value := range rows[row] { if len(value) != 1 { return nil, fmt.Errorf(\"cell (%d,%d) is not one character\", row, column) }; result[row][column] = value[0] } }")
		fmt.Fprintln(output, "\treturn result, nil")
		fmt.Fprintln(output, "}")
		fmt.Fprintln(output)
	}
	if types["randomList"] {
		fmt.Fprintln(output, "func decodeRandomList(raw json.RawMessage) ([]randomListEntry, error) {")
		fmt.Fprintln(output, "\tvar pairs [][]*int; if err := json.Unmarshal(raw, &pairs); err != nil { return nil, err }")
		fmt.Fprintln(output, "\tresult := make([]randomListEntry, len(pairs)); for index, pair := range pairs { if len(pair) != 2 || pair[0] == nil { return nil, fmt.Errorf(\"entry %d must contain [value, randomIndex]\", index) }; result[index] = randomListEntry{Value: *pair[0], RandomIndex: pair[1]} }")
		fmt.Fprintln(output, "\treturn result, nil")
		fmt.Fprintln(output, "}")
		fmt.Fprintln(output)
	}
}

func writeWorkbenchScaffold(output *bytes.Buffer, functionName string, parameters []parameter) {
	fmt.Fprintln(output, "// archiveReady remains false until the current attempt, three agent solutions,")
	fmt.Fprintln(output, "// analysis, and regression coverage are complete and verified.")
	fmt.Fprintln(output, "const archiveReady = false")
	fmt.Fprintln(output)
	fmt.Fprintln(output, "// Registered implementations run against the same cases.")
	fmt.Fprintln(output, "var currentAttemptSolution = testrunner.Solution{")
	fmt.Fprintln(output, "\tName: \"current attempt\",")
	fmt.Fprintln(output, "\tRun:  solveCurrentAttempt,")
	fmt.Fprintln(output, "}")
	fmt.Fprintln(output)
	fmt.Fprintln(output, "// agentSolutions is intentionally empty until an AI review adds direct,")
	fmt.Fprintln(output, "// improved, and preferred implementations.")
	fmt.Fprintln(output, "var agentSolutions = []testrunner.Solution{}")
	fmt.Fprintln(output)
	fmt.Fprintln(output, "var solutions = append([]testrunner.Solution{currentAttemptSolution}, agentSolutions...)")
	fmt.Fprintln(output)
	fmt.Fprintln(output, "// =============================================================================")
	fmt.Fprintln(output, "// CURRENT ATTEMPT")
	fmt.Fprintln(output, "//")
	fmt.Fprintln(output, "// Implement your solution here. Keep this body intact when requesting review.")
	fmt.Fprintln(output, "// =============================================================================")
	fmt.Fprintln(output)
	fmt.Fprintln(output, "func solveCurrentAttempt(input string) string {")
	names := parameterNames(parameters)
	fmt.Fprintf(output, "\t%s := parse%sInput(input)\n", strings.Join(names, ", "), functionName)
	for _, name := range names {
		fmt.Fprintf(output, "\t_ = %s\n", name)
	}
	fmt.Fprintln(output)
	fmt.Fprintln(output, "\t// Implement your algorithm with the typed values above.")
	fmt.Fprintln(output, "\treturn \"\"")
	fmt.Fprintln(output, "}")
	fmt.Fprintln(output)
	fmt.Fprintln(output, "// =============================================================================")
	fmt.Fprintln(output, "// AGENT SOLUTIONS")
	fmt.Fprintln(output, "//")
	fmt.Fprintln(output, "// Intentionally blank until you ask an AI agent to review this problem.")
	fmt.Fprintln(output, "// =============================================================================")
	fmt.Fprintln(output)
	fmt.Fprintln(output, "/*")
	fmt.Fprintln(output, "===============================================================================")
	fmt.Fprintln(output, "AGENT REVIEW AND COMPLEXITY ANALYSIS")
	fmt.Fprintln(output)
	fmt.Fprintln(output, "Current attempt:")
	fmt.Fprintln(output, "- What is correct:")
	fmt.Fprintln(output, "- First failing assumption or bottleneck:")
	fmt.Fprintln(output, "- Relevant concept and invariant:")
	fmt.Fprintln(output)
	fmt.Fprintln(output, "Solution 1 - Direct:")
	fmt.Fprintln(output, "- Correctness:")
	fmt.Fprintln(output, "- Concrete example:")
	fmt.Fprintln(output, "- Time complexity:")
	fmt.Fprintln(output, "- Auxiliary-space complexity:")
	fmt.Fprintln(output)
	fmt.Fprintln(output, "Solution 2 - Improved:")
	fmt.Fprintln(output, "- Correctness:")
	fmt.Fprintln(output, "- Concrete example:")
	fmt.Fprintln(output, "- Time complexity:")
	fmt.Fprintln(output, "- Auxiliary-space complexity:")
	fmt.Fprintln(output)
	fmt.Fprintln(output, "Solution 3 - Preferred:")
	fmt.Fprintln(output, "- Correctness:")
	fmt.Fprintln(output, "- Concrete example:")
	fmt.Fprintln(output, "- Time complexity:")
	fmt.Fprintln(output, "- Auxiliary-space complexity:")
	fmt.Fprintln(output)
	fmt.Fprintln(output, "AGENT REVIEW AND COMPLEXITY ANALYSIS END")
	fmt.Fprintln(output, "===============================================================================")
	fmt.Fprintln(output, "*/")
	fmt.Fprintln(output)
	fmt.Fprintln(output, "func main() {")
	fmt.Fprintln(output, "\tif err := testrunner.Run(os.Stdout, sampleCases, solutions); err != nil {")
	fmt.Fprintln(output, "\t\tfmt.Fprintln(os.Stderr, \"workbench:\", err)")
	fmt.Fprintln(output, "\t\tos.Exit(1)")
	fmt.Fprintln(output, "\t}")
	fmt.Fprintln(output, "}")
}

func parameterNames(parameters []parameter) []string {
	names := make([]string, len(parameters))
	for index, param := range parameters {
		names[index] = uniqueName(goIdentifier(param.Name), names[:index])
	}
	return names
}

func goType(problemType string) string {
	switch problemType {
	case "integer":
		return "int"
	case "double":
		return "float64"
	case "boolean":
		return "bool"
	case "string":
		return "string"
	case "integer[]", "list<integer>":
		return "[]int"
	case "integer[][]", "list<list<integer>>":
		return "[][]int"
	case "double[]", "list<double>":
		return "[]float64"
	case "string[]", "list<string>", "list<String>":
		return "[]string"
	case "string[][]", "list<list<string>>":
		return "[][]string"
	case "character[][]":
		return "[][]byte"
	case "TreeNode":
		return "*TreeNode"
	case "ListNode":
		return "*ListNode"
	case "ListNode[]":
		return "[]*ListNode"
	case "operationArguments":
		return "[][]json.RawMessage"
	case "randomList":
		return "[]randomListEntry"
	default:
		panic("unsupported LeetCode type: " + problemType)
	}
}

func exportedIdentifier(slug string) string {
	var result strings.Builder
	for _, part := range strings.Split(slug, "-") {
		if part == "" {
			continue
		}
		runes := []rune(part)
		result.WriteRune(unicode.ToUpper(runes[0]))
		result.WriteString(string(runes[1:]))
	}
	return result.String()
}

func goIdentifier(name string) string {
	if name == "" {
		return "value"
	}
	var result strings.Builder
	for index, char := range name {
		if unicode.IsLetter(char) || char == '_' || (index > 0 && unicode.IsDigit(char)) {
			result.WriteRune(char)
		}
	}
	identifier := result.String()
	if identifier == "" {
		return "value"
	}
	reserved := map[string]bool{"break": true, "case": true, "chan": true, "const": true, "continue": true, "default": true, "defer": true, "else": true, "fallthrough": true, "for": true, "func": true, "go": true, "goto": true, "if": true, "import": true, "interface": true, "map": true, "package": true, "range": true, "return": true, "select": true, "struct": true, "switch": true, "type": true, "var": true}
	if reserved[identifier] {
		return identifier + "Value"
	}
	return identifier
}

func uniqueName(name string, prior []string) string {
	candidate := name
	for suffix := 2; ; suffix++ {
		used := false
		for _, existing := range prior {
			if candidate == existing {
				used = true
				break
			}
		}
		if !used {
			return candidate
		}
		candidate = name + strconv.Itoa(suffix)
	}
}

func rawLiteral(value string) string {
	if !strings.Contains(value, "`") {
		return "`" + value + "`"
	}
	return strconv.Quote(value)
}

func updateProblemIndex(path string, filenames []string) error {
	contents, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read problem index: %w", err)
	}
	text := string(contents)
	sort.Strings(filenames)
	var additions strings.Builder
	for _, filename := range filenames {
		link := fmt.Sprintf("- [%s](%s)", filename, filename)
		if !strings.Contains(text, link) {
			additions.WriteString(link + "\n")
		}
	}
	if additions.Len() == 0 {
		return nil
	}
	if !strings.HasSuffix(text, "\n") {
		text += "\n"
	}
	return os.WriteFile(path, []byte(text+additions.String()), 0o644)
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, "import top interview 150:", err)
	os.Exit(1)
}
