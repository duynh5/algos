//go:build ignore

// Difficulty: Easy
// Source: leetcode
// Problem ID: 20

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/duyserveacc/algos/internal/testrunner"
)

/*
===============================================================================
Valid Parentheses

Source: LeetCode
Problem ID: 20
Source URL: https://leetcode.com/problems/valid-parentheses/description/

Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

- Open brackets must be closed by the same type of brackets.

- Open brackets must be closed in the correct order.

- Every close bracket has a corresponding open bracket of the same type.

Example 1:

Input: s = "()"

Output: true

Example 2:

Input: s = "()[]{}"

Output: true

Example 3:

Input: s = "(]"

Output: false

Example 4:

Input: s = "([])"

Output: true

Example 5:

Input: s = "([)]"

Output: false

Constraints:

- 1 <= s.length <= 10^4

- s consists of parentheses only '()[]{}'.
===============================================================================
*/

// Problem-specific support.
//
// Keep shared judge types and parsing or formatting helpers here. Name adapters
// after the problem. Put setup used only by the current attempt at the top of
// solveCurrentAttempt below.
//
// This intentionally unsolved practice archive follows the workbench template.

func parseValidParenthesesInput(input string) string {
	var args []json.RawMessage
	if err := json.Unmarshal([]byte(input), &args); err != nil {
		panic(fmt.Errorf("parse valid-parentheses input: %w", err))
	}
	if len(args) != 1 {
		panic(fmt.Sprintf("parse valid-parentheses input: got %d arguments, want 1", len(args)))
	}
	var s string
	argumentError0 := json.Unmarshal(args[0], &s)
	if argumentError0 != nil {
		panic(fmt.Errorf("parse valid-parentheses argument s: %w", argumentError0))
	}
	return s
}

func maximumValidParenthesesCase() testrunner.Case {
	value := strings.Repeat("(", 5000) + strings.Repeat(")", 5000)
	input, err := json.Marshal([]any{value})
	if err != nil {
		panic(fmt.Errorf("build maximum valid-parentheses case: %w", err))
	}
	return testrunner.Case{
		Name:  "Boundary: Maximum-length valid nesting",
		Input: string(input),
		Want:  `true`,
	}
}

func maximumInvalidParenthesesCase() testrunner.Case {
	value := strings.Repeat("(", 5000) + strings.Repeat(")", 4999) + "]"
	input, err := json.Marshal([]any{value})
	if err != nil {
		panic(fmt.Errorf("build maximum invalid-parentheses case: %w", err))
	}
	return testrunner.Case{
		Name:  "Boundary: Maximum-length late mismatch",
		Input: string(input),
		Want:  `false`,
	}
}

// sampleCases contains the examples that every registered solution must pass.
// Inputs encode the judge arguments as a positional JSON array.
var sampleCases = []testrunner.Case{
	{
		Name:  "Official example 1",
		Input: `["()"]`,
		Want:  `true`,
	},
	{
		Name:  "Official example 2",
		Input: `["()[]{}"]`,
		Want:  `true`,
	},
	{
		Name:  "Official example 3",
		Input: `["(]"]`,
		Want:  `false`,
	},
	{
		Name:  "Official example 4",
		Input: `["([])"]`,
		Want:  `true`,
	},
	{
		Name:  "Official example 5",
		Input: `["([)]"]`,
		Want:  `false`,
	},
	{
		Name:  "Additional: Minimum unmatched opener",
		Input: `["("]`,
		Want:  `false`,
	},
	{
		Name:  "Additional: Deep mixed nesting",
		Input: `["{[()]}"]`,
		Want:  `true`,
	},
	{
		Name:  "Additional: Minimum unmatched closer",
		Input: `["]"]`,
		Want:  `false`,
	},
	{
		Name:  "Additional: Deep same-type nesting",
		Input: `["[[[]]]"]`,
		Want:  `true`,
	},
	{
		Name:  "Additional: Concatenated nested groups",
		Input: `["([]){}[()]"]`,
		Want:  `true`,
	},
	{
		Name:  "Adversarial: Extra closer after valid prefix",
		Input: `["()]"]`,
		Want:  `false`,
	},
	{
		Name:  "Adversarial: Unclosed opener after valid prefix",
		Input: `["()[]{"]`,
		Want:  `false`,
	},
	{
		Name:  "Adversarial: Correct totals but wrong nested type",
		Input: `["{[}]"]`,
		Want:  `false`,
	},
	{
		Name:  "Additional: Odd length cannot close",
		Input: `["(()"]`,
		Want:  `false`,
	},
	maximumValidParenthesesCase(),
	maximumInvalidParenthesesCase(),
}

// archiveReady is true after the current attempt, three agent solutions,
// analysis, and regression coverage are complete and verified.
const archiveReady = true

// Registered implementations run against the same cases.
var currentAttemptSolution = testrunner.Solution{
	Name: "current attempt",
	Run:  solveCurrentAttempt,
}

// agentSolutions contains direct, improved, and preferred implementations.
var agentSolutions = []testrunner.Solution{
	{Name: "remove adjacent pairs", Run: solveByRemovingAdjacentPairs},
	{Name: "parse bracket grammar recursively", Run: solveByRecursiveDescent},
	{Name: "stack expected closing brackets", Run: solveByExpectedClosers},
}

var solutions = append([]testrunner.Solution{currentAttemptSolution}, agentSolutions...)

// =============================================================================
// CURRENT ATTEMPT
//
// Implement your solution here. Keep this body intact when requesting review.
// =============================================================================

func solveCurrentAttempt(input string) string {
	s := parseValidParenthesesInput(input)
	paren_mapping := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := make([]byte, 0, len(s))

	for index := 0; index < len(s); index++ {
		bracket := s[index]
		expectedOpening, isClosing := paren_mapping[bracket]
		if !isClosing {
			stack = append(stack, bracket)
			continue
		}

		if len(stack) == 0 || stack[len(stack)-1] != expectedOpening {
			return strconv.FormatBool(false)
		}
		stack = stack[:len(stack)-1]
	}

	return strconv.FormatBool(len(stack) == 0)
}

// =============================================================================
// AGENT SOLUTIONS
//
// Direct, improved, and preferred implementations authored during review.
// =============================================================================

// solveByRemovingAdjacentPairs repeatedly removes innermost matching pairs.
// Any valid non-empty bracket sequence must contain at least one such pair.
func solveByRemovingAdjacentPairs(input string) string {
	s := parseValidParenthesesInput(input)
	removePairs := strings.NewReplacer("()", "", "[]", "", "{}", "")

	for {
		reduced := removePairs.Replace(s)
		if len(reduced) == len(s) {
			return strconv.FormatBool(reduced == "")
		}
		s = reduced
	}
}

// solveByRecursiveDescent recognizes the bracket grammar while tracking the
// closing bracket required by each active recursive call.
func solveByRecursiveDescent(input string) string {
	s := parseValidParenthesesInput(input)
	index := 0
	valid := consumeBracketSequence(s, &index, 0) && index == len(s)
	return strconv.FormatBool(valid)
}

func consumeBracketSequence(s string, index *int, expectedClosing byte) bool {
	for *index < len(s) {
		bracket := s[*index]
		switch bracket {
		case '(', '[', '{':
			(*index)++
			closing := mapOpeningToClosing(bracket)
			if !consumeBracketSequence(s, index, closing) {
				return false
			}
		case ')', ']', '}':
			if expectedClosing == 0 || bracket != expectedClosing {
				return false
			}
			(*index)++
			return true
		}
	}

	return expectedClosing == 0
}

func mapOpeningToClosing(opening byte) byte {
	switch opening {
	case '(':
		return ')'
	case '[':
		return ']'
	case '{':
		return '}'
	default:
		panic(fmt.Sprintf("valid parentheses: unexpected opening bracket %q", opening))
	}
}

// solveByExpectedClosers parses the input and delegates to the preferred
// judge-facing implementation.
func solveByExpectedClosers(input string) string {
	s := parseValidParenthesesInput(input)
	return strconv.FormatBool(isValid(s))
}

// isValid pushes the closing bracket each opener requires, so every closer can
// be compared directly with the top of the stack.
func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	expectedClosers := make([]byte, 0, len(s)/2)
	for index := 0; index < len(s); index++ {
		switch bracket := s[index]; bracket {
		case '(':
			expectedClosers = append(expectedClosers, ')')
		case '[':
			expectedClosers = append(expectedClosers, ']')
		case '{':
			expectedClosers = append(expectedClosers, '}')
		default:
			if len(expectedClosers) == 0 || expectedClosers[len(expectedClosers)-1] != bracket {
				return false
			}
			expectedClosers = expectedClosers[:len(expectedClosers)-1]
		}
	}

	return len(expectedClosers) == 0
}

/*
===============================================================================
AGENT REVIEW AND COMPLEXITY ANALYSIS

Current attempt:
- What is correct: The attempt implements the proposed last-in, first-out
  matching process. Opening brackets are pushed. For each closer, the map finds
  the only acceptable opener, and the function rejects an empty stack or a
  different top before popping. Requiring an empty stack after the complete
  scan also rejects leftover openers.
- First failing assumption or bottleneck: No correctness failure remains under
  the stated constraint that every byte is one of ()[]{}. The map is allocated
  for each call and every closer performs a hash lookup; those are small
  constant costs that the preferred solution avoids. An odd input length could
  be rejected before scanning, but omitting that shortcut does not change the
  asymptotic bound.
- Relevant concept and invariant: See
  concepts/data-structures/stack/README.md. Before processing s[index], stack
  contains exactly the unmatched opening brackets from s[:index], ordered from
  oldest to newest. A matching closer removes only the newest opener, which is
  precisely the required nesting order.
- Concrete example: For s="([{}])", the stack contents change from "" to
  "(", "([", and "([{". The '}' maps to '{', matches the top, and pops it;
  ']' then pops '[', and ')' pops '('. The final empty stack produces true.
- Time complexity: JSON parsing is Theta(n). The scan is O(1) in the best case
  after parsing when the first bracket is an unmatched closer, and Theta(n) in
  the average and worst scanning cases. End to end, best, average, and worst
  time are Theta(n) because parsing must read the complete encoded input.
- Auxiliary-space complexity: The stack holds at most n opening brackets, so
  algorithmic space is O(n); the fixed three-entry map is O(1). Parsing also
  allocates Theta(n), while boolean output is constant-size.
- Relationship to agent solutions: Solution 1 uses fixed-point string
  reduction rather than explicit nesting state. Solution 2 stores nesting in
  recursive calls. Solution 3 uses the same stack invariant as this attempt,
  but pushes expected closers and compares them directly instead of mapping a
  closer back to its opener; it is a separate implementation of the same core
  algorithm.

Solution 1 - Remove adjacent pairs (direct):
- Correctness: Removing (), [], or {} preserves validity because each is an
  already-complete innermost group. Conversely, every non-empty valid sequence
  contains an innermost matching pair with no brackets between its endpoints.
  Therefore repeated removal empties every valid string, while a non-empty
  fixed point cannot be valid.
- Concrete example: For s="([{}])", the first pass removes {} and produces
  "([])". The second removes [] and produces "()". The third removes () and
  produces ""; the next pass makes no change, so the function returns true.
- Time complexity: Parsing and each replacement pass take Theta(n) relative to
  the current string. A flat input such as "()[]{}" finishes in Theta(n), while
  n/2 nested pairs require Theta(n) passes over progressively shorter strings,
  for Theta(n^2) worst-case time. No input distribution supports a tighter
  average-case claim than O(n^2).
- Auxiliary-space complexity: Each pass creates a replacement string, with
  O(n) live auxiliary space and O(n^2) total allocation volume in the worst
  case. Parsing retains Theta(n) input data; output is constant-size.

Solution 2 - Recursive bracket grammar (improved):
- Correctness: Each recursive call owns one expected closing bracket. It
  recursively consumes every nested opener and returns only after consuming
  the exact closer belonging to that call. A mismatched closer, a closer at the
  top level, or end-of-input before an expected closer returns false. Reaching
  the end at the top level therefore means every bracket was matched in order.
- Concrete example: For s="([]){}", the top-level call sees '(' at index 0
  and starts a call expecting ')'. That call sees '[' at index 1 and starts a
  call expecting ']'; it consumes ']' at index 2, then its parent consumes ')'
  at index 3. The top level similarly consumes the {} group and ends at index
  6 with no expected closer, so it returns true.
- Time complexity: Each input byte is consumed once, so the scan is Theta(n).
  Including JSON parsing, best, average, and worst time are Theta(n).
- Auxiliary-space complexity: Recursion depth equals maximum nesting depth d,
  giving O(d) algorithmic space and O(n) in the worst case. Parsing allocates
  Theta(n); boolean output is constant-size.

Solution 3 - Stack expected closing brackets (preferred):
- Correctness: For each opener, the function pushes the one closer that would
  validly finish that group. Thus the stack top is always the only legal next
  closer. A different closer or an empty stack proves invalidity immediately;
  popping a match preserves the invariant. After the scan, an empty stack is
  equivalent to every opened group having been closed.
- Concrete example: For s="{[()]}", the expected-closer stack changes from
  "" to "}", "}]", and "}])". Reading ')' pops ')', reading ']' pops ']',
  and reading '}' pops '}'. The stack is empty, so the result is true.
- Time complexity: After the odd-length O(1) check, each byte is pushed or
  compared and popped once, giving Theta(n) scanning time. Including JSON
  parsing, best, average, and worst end-to-end time are Theta(n).
- Auxiliary-space complexity: At most n/2 expected closers can be live for a
  valid even-length input, so algorithmic space is O(n). Parsing also allocates
  Theta(n); boolean output is constant-size.

AGENT REVIEW AND COMPLEXITY ANALYSIS END
===============================================================================
*/

func main() {
	if err := testrunner.Run(os.Stdout, sampleCases, solutions); err != nil {
		fmt.Fprintln(os.Stderr, "workbench:", err)
		os.Exit(1)
	}
}
