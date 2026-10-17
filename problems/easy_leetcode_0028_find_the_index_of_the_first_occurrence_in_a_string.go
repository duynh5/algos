//go:build ignore

// Difficulty: Easy
// Source: leetcode
// Problem ID: 28

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/duyserveacc/algos/internal/testrunner"
)

/*
===============================================================================
Find the Index of the First Occurrence in a String

Source: LeetCode
Problem ID: 28
Source URL: https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/description/

Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Example 1:

Input: haystack = "sadbutsad", needle = "sad"
Output: 0
Explanation: "sad" occurs at index 0 and 6.
The first occurrence is at index 0, so we return 0.

Example 2:

Input: haystack = "leetcode", needle = "leeto"
Output: -1
Explanation: "leeto" did not occur in "leetcode", so we return -1.

Constraints:

- 1 <= haystack.length, needle.length <= 10^4

- haystack and needle consist of only lowercase English characters.
===============================================================================
*/

// Problem-specific support.
//
// Keep shared judge types and parsing or formatting helpers here. Name adapters
// after the problem. Put setup used only by the current attempt at the top of
// solveCurrentAttempt below.
//
// This intentionally unsolved practice archive follows the workbench template.

func parseFindTheIndexOfTheFirstOccurrenceInAStringInput(input string) (string, string) {
	var args []json.RawMessage
	if err := json.Unmarshal([]byte(input), &args); err != nil {
		panic(fmt.Errorf("parse find-the-index-of-the-first-occurrence-in-a-string input: %w", err))
	}
	if len(args) != 2 {
		panic(fmt.Sprintf("parse find-the-index-of-the-first-occurrence-in-a-string input: got %d arguments, want 2", len(args)))
	}
	var haystack string
	argumentError0 := json.Unmarshal(args[0], &haystack)
	if argumentError0 != nil {
		panic(fmt.Errorf("parse find-the-index-of-the-first-occurrence-in-a-string argument haystack: %w", argumentError0))
	}
	var needle string
	argumentError1 := json.Unmarshal(args[1], &needle)
	if argumentError1 != nil {
		panic(fmt.Errorf("parse find-the-index-of-the-first-occurrence-in-a-string argument needle: %w", argumentError1))
	}
	return haystack, needle
}

func maximumFirstOccurrenceCase() testrunner.Case {
	haystack := strings.Repeat("a", 10000)
	needle := strings.Repeat("a", 9999) + "b"
	input, err := json.Marshal([]any{haystack, needle})
	if err != nil {
		panic(fmt.Errorf("build maximum first-occurrence case: %w", err))
	}
	return testrunner.Case{
		Name:  "Boundary: Maximum-length repeated-prefix miss",
		Input: string(input),
		Want:  `-1`,
	}
}

func formatFindTheIndexOfTheFirstOccurrenceInAStringOutput(result int) string {
	return testrunner.FormatJSONOutput(result)
}

// sampleCases contains the examples that every registered solution must pass.
// Inputs encode the judge arguments as a positional JSON array.
var sampleCases = []testrunner.Case{
	{
		Name:  "Official example 1",
		Input: `["sadbutsad","sad"]`,
		Want:  `0`,
	},
	{
		Name:  "Official example 2",
		Input: `["leetcode","leeto"]`,
		Want:  `-1`,
	},
	{
		Name:  "Additional: Exact one-character match",
		Input: `["a","a"]`,
		Want:  `0`,
	},
	{
		Name:  "Additional: Match at final position",
		Input: `["abc","c"]`,
		Want:  `2`,
	},
	{
		Name:  "Additional: Needle longer than haystack",
		Input: `["a","aa"]`,
		Want:  `-1`,
	},
	{
		Name:  "Additional: Exact multi-character match",
		Input: `["algorithm","algorithm"]`,
		Want:  `0`,
	},
	{
		Name:  "Additional: Equal lengths without match",
		Input: `["abc","abd"]`,
		Want:  `-1`,
	},
	{
		Name:  "Adversarial: Overlapping false start",
		Input: `["aaaab","aaab"]`,
		Want:  `1`,
	},
	{
		Name:  "Additional: Multi-character match at final position",
		Input: `["abcxabcd","abcd"]`,
		Want:  `4`,
	},
	{
		Name:  "Adversarial: Repeated pattern returns first match",
		Input: `["mississippi","issi"]`,
		Want:  `1`,
	},
	maximumFirstOccurrenceCase(),
}

// archiveReady remains false until the current attempt, three agent solutions,
// analysis, and regression coverage are complete and verified.
const archiveReady = false

// Registered implementations run against the same cases.
var currentAttemptSolution = testrunner.Solution{
	Name: "current attempt",
	Run:  solveCurrentAttempt,
}

// agentSolutions is intentionally empty until an AI review adds direct,
// improved, and preferred implementations.
var agentSolutions = []testrunner.Solution{}

var solutions = append([]testrunner.Solution{currentAttemptSolution}, agentSolutions...)

// =============================================================================
// CURRENT ATTEMPT
//
// Implement your solution here. Keep this body intact when requesting review.
// =============================================================================

func solveCurrentAttempt(input string) string {
	haystack, needle := parseFindTheIndexOfTheFirstOccurrenceInAStringInput(input)
	_ = haystack
	_ = needle

	// Implement your algorithm with the typed values above.
	return ""
}

// =============================================================================
// AGENT SOLUTIONS
//
// Intentionally blank until you ask an AI agent to review this problem.
// =============================================================================

/*
===============================================================================
AGENT REVIEW AND COMPLEXITY ANALYSIS

Current attempt:
- What is correct:
- First failing assumption or bottleneck:
- Relevant concept and invariant:

Solution 1 - Direct:
- Correctness:
- Concrete example:
- Time complexity:
- Auxiliary-space complexity:

Solution 2 - Improved:
- Correctness:
- Concrete example:
- Time complexity:
- Auxiliary-space complexity:

Solution 3 - Preferred:
- Correctness:
- Concrete example:
- Time complexity:
- Auxiliary-space complexity:

AGENT REVIEW AND COMPLEXITY ANALYSIS END
===============================================================================
*/

func main() {
	if err := testrunner.Run(os.Stdout, sampleCases, solutions); err != nil {
		fmt.Fprintln(os.Stderr, "workbench:", err)
		os.Exit(1)
	}
}
