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

func problemOutputType(entry catalogEntry, meta metadata, parameters []parameter) string {
	if meta.SystemDesign {
		return "operationResults"
	}
	if entry.ID == "138" {
		return "randomList"
	}
	if entry.ID == "133" {
		return "graphNode"
	}
	if entry.ID == "427" {
		return "quadNode"
	}
	if meta.Output != nil {
		if meta.Output.ParamIndex < 0 || meta.Output.ParamIndex >= len(parameters) {
			panic(fmt.Sprintf("problem %s output parameter index %d is out of range", entry.ID, meta.Output.ParamIndex))
		}
		return parameters[meta.Output.ParamIndex].Type
	}
	return meta.Return.Type
}

func writeOutputFormatter(output *bytes.Buffer, entry catalogEntry, functionName string, meta metadata, parameters []parameter) {
	if entry.ID == "26" || entry.ID == "27" || entry.ID == "80" {
		writeRetainedPrefixOutputFormatter(output, entry, functionName)
		return
	}
	problemType := problemOutputType(entry, meta, parameters)
	resultType := goType(problemType)
	fmt.Fprintf(output, "func format%sOutput(result %s) string {\n", functionName, resultType)
	switch problemType {
	case "TreeNode":
		writeTreeOutputFormatter(output, entry)
	case "ListNode":
		fmt.Fprintln(output, "\tvalues := make([]int, 0)")
		fmt.Fprintln(output, "\tfor node := result; node != nil; node = node.Next {")
		fmt.Fprintln(output, "\t\tif len(values) == 100000 { panic(\"format output: linked list exceeds 100000 nodes or contains a cycle\") }")
		fmt.Fprintln(output, "\t\tvalues = append(values, node.Val)")
		fmt.Fprintln(output, "\t}")
		fmt.Fprintln(output, "\treturn testrunner.FormatJSONOutput(values)")
	case "character[][]":
		fmt.Fprintln(output, "\trows := make([][]string, len(result))")
		fmt.Fprintln(output, "\tfor row := range result { rows[row] = make([]string, len(result[row])); for column, value := range result[row] { rows[row][column] = string(value) } }")
		fmt.Fprintln(output, "\treturn testrunner.FormatJSONOutput(rows)")
	case "randomList":
		fmt.Fprintln(output, "\tpairs := make([][2]any, len(result))")
		fmt.Fprintln(output, "\tfor index, entry := range result { pairs[index] = [2]any{entry.Value, entry.RandomIndex} }")
		fmt.Fprintln(output, "\treturn testrunner.FormatJSONOutput(pairs)")
	case "graphNode":
		writeGraphOutputFormatter(output)
	case "quadNode":
		writeQuadTreeOutputFormatter(output)
	case "double":
		fmt.Fprintln(output, "\treturn fmt.Sprintf(\"%.5f\", result)")
	case "double[]", "list<double>":
		fmt.Fprintln(output, "\tformatted := \"[\"")
		fmt.Fprintln(output, "\tfor index, value := range result { if index > 0 { formatted += \",\" }; formatted += fmt.Sprintf(\"%.5f\", value) }")
		fmt.Fprintln(output, "\treturn formatted + \"]\"")
	default:
		fmt.Fprintln(output, "\treturn testrunner.FormatJSONOutput(result)")
	}
	fmt.Fprintln(output, "}")
	fmt.Fprintln(output)
}

func writeGraphOutputFormatter(output *bytes.Buffer) {
	fmt.Fprintln(output, "\tif result == nil { return `[]` }")
	fmt.Fprintln(output, "\tadjacency := make([][]int, 0)")
	fmt.Fprintln(output, "\tseen := make(map[*GraphNode]bool); queue := []*GraphNode{result}")
	fmt.Fprintln(output, "\tfor len(queue) > 0 {")
	fmt.Fprintln(output, "\t\tnode := queue[0]; queue = queue[1:]; if seen[node] { continue }; seen[node] = true")
	fmt.Fprintln(output, "\t\tif len(seen) > 100 { panic(\"format output: graph exceeds 100 nodes\") }")
	fmt.Fprintln(output, "\t\tfor len(adjacency) < node.Val { adjacency = append(adjacency, nil) }")
	fmt.Fprintln(output, "\t\tfor _, neighbor := range node.Neighbors { adjacency[node.Val-1] = append(adjacency[node.Val-1], neighbor.Val); if !seen[neighbor] { queue = append(queue, neighbor) } }")
	fmt.Fprintln(output, "\t}")
	fmt.Fprintln(output, "\treturn testrunner.FormatJSONOutput(adjacency)")
}

func writeQuadTreeOutputFormatter(output *bytes.Buffer) {
	fmt.Fprintln(output, "\tif result == nil { return `[]` }")
	fmt.Fprintln(output, "\tvalues := make([]any, 0); queue := []*QuadNode{result}")
	fmt.Fprintln(output, "\tfor len(queue) > 0 {")
	fmt.Fprintln(output, "\t\tnode := queue[0]; queue = queue[1:]")
	fmt.Fprintln(output, "\t\tif node == nil { values = append(values, nil); continue }")
	fmt.Fprintln(output, "\t\tleaf, value := 0, 0; if node.IsLeaf { leaf = 1 }; if node.Val || !node.IsLeaf { value = 1 }")
	fmt.Fprintln(output, "\t\tvalues = append(values, [2]int{leaf, value})")
	fmt.Fprintln(output, "\t\tif node.IsLeaf { queue = append(queue, nil, nil, nil, nil) } else { queue = append(queue, node.TopLeft, node.TopRight, node.BottomLeft, node.BottomRight) }")
	fmt.Fprintln(output, "\t\tif len(values) > 5500 { panic(\"format output: quad tree exceeds its maximum serialized size or contains a cycle\") }")
	fmt.Fprintln(output, "\t}")
	fmt.Fprintln(output, "\tfor len(values) > 0 && values[len(values)-1] == nil { values = values[:len(values)-1] }")
	fmt.Fprintln(output, "\treturn testrunner.FormatJSONOutput(values)")
}

func writeRetainedPrefixOutputFormatter(output *bytes.Buffer, entry catalogEntry, functionName string) {
	fmt.Fprintf(output, "func format%sOutput(k int, nums []int) string {\n", functionName)
	fmt.Fprintln(output, "\tif k < 0 || k > len(nums) { panic(fmt.Sprintf(\"format output: k=%d outside [0,%d]\", k, len(nums))) }")
	if entry.ID == "27" {
		fmt.Fprintln(output, "\tretained := append([]int(nil), nums[:k]...)")
		fmt.Fprintln(output, "\tfor index := 1; index < len(retained); index++ { for current := index; current > 0 && retained[current] < retained[current-1]; current-- { retained[current], retained[current-1] = retained[current-1], retained[current] } }")
		fmt.Fprintln(output, "\tnums = append(retained, nums[k:]...)")
	}
	fmt.Fprintln(output, "\tformatted := fmt.Sprintf(\"%d, nums = [\", k)")
	fmt.Fprintln(output, "\tfor index := range nums { if index > 0 { formatted += \",\" }; if index < k { formatted += fmt.Sprint(nums[index]) } else { formatted += \"_\" } }")
	fmt.Fprintln(output, "\treturn formatted + \"]\"")
	fmt.Fprintln(output, "}")
	fmt.Fprintln(output)
}

func writeTreeOutputFormatter(output *bytes.Buffer, entry catalogEntry) {
	if entry.ID == "236" {
		fmt.Fprintln(output, "\tif result == nil { return `null` }")
		fmt.Fprintln(output, "\treturn testrunner.FormatJSONOutput(result.Val)")
		return
	}
	if entry.ID == "117" {
		fmt.Fprintln(output, "\tif result == nil { return `[]` }")
		fmt.Fprintln(output, "\tparts := make([]string, 0)")
		fmt.Fprintln(output, "\tfor levelStart := result; levelStart != nil; {")
		fmt.Fprintln(output, "\t\tvar nextLevelStart *TreeNode")
		fmt.Fprintln(output, "\t\tfor node := levelStart; node != nil; node = node.Next {")
		fmt.Fprintln(output, "\t\t\tparts = append(parts, fmt.Sprint(node.Val))")
		fmt.Fprintln(output, "\t\t\tif nextLevelStart == nil { if node.Left != nil { nextLevelStart = node.Left } else if node.Right != nil { nextLevelStart = node.Right } }")
		fmt.Fprintln(output, "\t\t}")
		fmt.Fprintln(output, "\t\tparts = append(parts, \"#\")")
		fmt.Fprintln(output, "\t\tlevelStart = nextLevelStart")
		fmt.Fprintln(output, "\t}")
		fmt.Fprintln(output, "\tformatted := \"[\"")
		fmt.Fprintln(output, "\tfor index, part := range parts { if index > 0 { formatted += \",\" }; formatted += part }")
		fmt.Fprintln(output, "\treturn formatted + \"]\"")
		return
	}
	fmt.Fprintln(output, "\tif result == nil { return `[]` }")
	fmt.Fprintln(output, "\tvalues := make([]any, 0)")
	fmt.Fprintln(output, "\tqueue := []*TreeNode{result}")
	fmt.Fprintln(output, "\tnodeCount := 0")
	fmt.Fprintln(output, "\tfor len(queue) > 0 {")
	fmt.Fprintln(output, "\t\tnode := queue[0]; queue = queue[1:]")
	fmt.Fprintln(output, "\t\tif node == nil { values = append(values, nil); continue }")
	fmt.Fprintln(output, "\t\tnodeCount++; if nodeCount > 100000 { panic(\"format output: tree exceeds 100000 nodes or contains a cycle\") }")
	fmt.Fprintln(output, "\t\tvalues = append(values, node.Val); queue = append(queue, node.Left, node.Right)")
	fmt.Fprintln(output, "\t}")
	fmt.Fprintln(output, "\tfor len(values) > 0 && values[len(values)-1] == nil { values = values[:len(values)-1] }")
	fmt.Fprintln(output, "\treturn testrunner.FormatJSONOutput(values)")
}

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
	fmt.Fprintln(output, "\t// Implement your algorithm with the typed values above, then pass its")
	fmt.Fprintf(output, "\t// typed result to format%sOutput.\n", functionName)
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
	case "operationResults":
		return "[]any"
	case "graphNode":
		return "*GraphNode"
	case "quadNode":
		return "*QuadNode"
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
