//go:build ignore

// Difficulty: Easy
// Source: leetcode
// Problem ID: 392

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/duyserveacc/algos/dsa/internal/testrunner"
)

/*
======================================================================
Is Subsequence

Source: LeetCode
Problem ID: 392
Source URL: https://leetcode.com/problems/is-subsequence/description/

Given two strings s and t, return true if s is a subsequence of t, or
false otherwise.

A subsequence of a string is a new string that is formed from the
original string by deleting some (can be none) of the characters
without disturbing the relative positions of the remaining characters.
(i.e., "ace" is a subsequence of "abcde" while "aec" is not).

Example 1:

Input: s = "abc", t = "ahbgdc"
Output: true

Example 2:

Input: s = "axc", t = "ahbgdc"
Output: false

Constraints:

- 0 <= s.length <= 100

- 0 <= t.length <= 10^4

- s and t consist only of lowercase English letters.

Follow up: Suppose there are lots of incoming s, say s_1, s_2, ...,
s_k where k >= 10^9, and you want to check one by one to see if t has
its subsequence. In this scenario, how would you change your code?
======================================================================
*/

// Problem-specific support.
//
// Keep shared judge types and parsing or formatting helpers here.
// Name adapters after the problem. Put setup used only by the current
// attempt at the top of solveCurrentAttempt below.
//
// This intentionally unsolved practice archive follows the workbench
// template.

func parseIsSubsequenceInput(input string) (string, string) {
	var args []json.RawMessage
	if err := json.Unmarshal([]byte(input), &args); err != nil {
		panic(fmt.Errorf("parse is-subsequence input: %w", err))
	}
	if len(args) != 2 {
		panic(
			fmt.Sprintf(
				"parse is-subsequence input: got %d arguments,"+
					" want 2",
				len(args),
			),
		)
	}
	var s string
	argumentError0 := json.Unmarshal(args[0], &s)
	if argumentError0 != nil {
		panic(
			fmt.Errorf(
				"parse is-subsequence argument s: %w",
				argumentError0,
			),
		)
	}
	var t string
	argumentError1 := json.Unmarshal(args[1], &t)
	if argumentError1 != nil {
		panic(
			fmt.Errorf(
				"parse is-subsequence argument t: %w",
				argumentError1,
			),
		)
	}
	return s, t
}

func maximumIsSubsequenceCase() testrunner.Case {
	s := strings.Repeat("z", 100)
	t := strings.Repeat("a", 9900) + s
	input, err := json.Marshal([]any{s, t})
	if err != nil {
		panic(
			fmt.Errorf("build maximum is-subsequence case: %w", err),
		)
	}
	return testrunner.Case{
		Name: "Boundary: Maximum target length matches at " +
			"the final positions",
		Input: string(input),
		Want:  `true`,
	}
}

func formatIsSubsequenceOutput(result bool) string {
	return testrunner.FormatJSONOutput(result)
}

// sampleCases contains the examples that every registered solution
// must pass. Inputs encode the judge arguments as a positional JSON
// array.
var sampleCases = []testrunner.Case{
	{
		Name:  "Official example 1",
		Input: `["abc","ahbgdc"]`,
		Want:  `true`,
	},
	{
		Name:  "Official example 2",
		Input: `["axc","ahbgdc"]`,
		Want:  `false`,
	},
	{
		Name:  "Additional: Both strings empty",
		Input: `["",""]`,
		Want:  `true`,
	},
	{
		Name:  "Additional: Empty subsequence",
		Input: `["","abc"]`,
		Want:  `true`,
	},
	{
		Name:  "Additional: Nonempty cannot fit in empty",
		Input: `["a",""]`,
		Want:  `false`,
	},
	{
		Name:  "Additional: Equal strings are subsequences",
		Input: `["abc","abc"]`,
		Want:  `true`,
	},
	{
		Name:  "Additional: Longer candidate cannot be a subsequence",
		Input: `["abcd","abc"]`,
		Want:  `false`,
	},
	{
		Name:  "Adversarial: Enough characters in the wrong order",
		Input: `["aaaaab","baaaaa"]`,
		Want:  `false`,
	},
	{
		Name:  "Additional: Repeated matches require distinct positions",
		Input: `["aaa","aa"]`,
		Want:  `false`,
	},
	{
		Name:  "Additional: Noncontiguous characters preserve order",
		Input: `["ace","abcde"]`,
		Want:  `true`,
	},
	maximumIsSubsequenceCase(),
}

// archiveReady remains false until the current attempt, three agent
// solutions, analysis, and regression coverage are complete and
// verified.
const archiveReady = false

// Registered implementations run against the same cases.
var currentAttemptSolution = testrunner.Solution{
	Name: "current attempt",
	Run:  solveCurrentAttempt,
}

// agentSolutions is intentionally empty until an AI review adds
// direct, improved, and preferred implementations.
var agentSolutions = []testrunner.Solution{}

var solutions = append(
	[]testrunner.Solution{currentAttemptSolution},
	agentSolutions...,
)

// ===================================================================
// CURRENT ATTEMPT
//
// Implement your solution here. Keep this body intact when requesting
// review.
// ===================================================================

func solveCurrentAttempt(input string) string {
	s, t := parseIsSubsequenceInput(input)
	_ = s
	_ = t

	// Implement your algorithm with the typed values above.
	return ""
}

// ===================================================================
// AGENT SOLUTIONS
//
// Intentionally blank until you ask an AI agent to review this
// problem.
// ===================================================================

/*
======================================================================
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
======================================================================
*/

func main() {
	if err := testrunner.Run(
		os.Stdout,
		sampleCases,
		solutions,
	); err != nil {
		fmt.Fprintln(os.Stderr, "workbench:", err)
		os.Exit(1)
	}
}
