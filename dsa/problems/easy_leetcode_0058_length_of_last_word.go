//go:build ignore

// Difficulty: Easy
// Source: leetcode
// Problem ID: 58

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
Length of Last Word

Source: LeetCode
Problem ID: 58
Source URL:
https://leetcode.com/problems/length-of-last-word/description/

Given a string s consisting of words and spaces, return the length of
the last word in the string.

A word is a maximal substring consisting of non-space characters only.

Example 1:

Input: s = "Hello World"
Output: 5
Explanation: The last word is "World" with length 5.

Example 2:

Input: s = " fly me to the moon "
Output: 4
Explanation: The last word is "moon" with length 4.

Example 3:

Input: s = "luffy is still joyboy"
Output: 6
Explanation: The last word is "joyboy" with length 6.

Constraints:

- 1 <= s.length <= 10^4

- s consists of only English letters and spaces ' '.

- There will be at least one word in s.
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

func parseLengthOfLastWordInput(input string) string {
	var args []json.RawMessage
	if err := json.Unmarshal([]byte(input), &args); err != nil {
		panic(fmt.Errorf("parse length-of-last-word input: %w", err))
	}
	if len(args) != 1 {
		panic(fmt.Sprintf("parse length-of-last-word input: got %d "+
			"arguments, want 1", len(args)))
	}
	var s string
	argumentError0 := json.Unmarshal(args[0], &s)
	if argumentError0 != nil {
		panic(
			fmt.Errorf(
				"parse length-of-last-word argument s: %w",
				argumentError0,
			),
		)
	}
	return s
}

func maximumLengthOfLastWordCase() testrunner.Case {
	value := strings.Repeat("a", 5000) + strings.Repeat(" ", 5000)
	input, err := json.Marshal([]any{value})
	if err != nil {
		panic(
			fmt.Errorf(
				"build maximum length-of-last-word case: %w",
				err,
			),
		)
	}
	return testrunner.Case{
		Name:  "Boundary: Maximum length with trailing spaces",
		Input: string(input),
		Want:  `5000`,
	}
}

func formatLengthOfLastWordOutput(result int) string {
	return testrunner.FormatJSONOutput(result)
}

// sampleCases contains the examples that every registered solution
// must pass. Inputs encode the judge arguments as a positional JSON
// array.
var sampleCases = []testrunner.Case{
	{
		Name:  "Official example 1",
		Input: `["Hello World"]`,
		Want:  `5`,
	},
	{
		Name:  "Official example 2",
		Input: `["   fly me   to   the moon  "]`,
		Want:  `4`,
	},
	{
		Name:  "Official example 3",
		Input: `["luffy is still joyboy"]`,
		Want:  `6`,
	},
	{
		Name:  "Additional: Single letter",
		Input: `["a"]`,
		Want:  `1`,
	},
	{
		Name:  "Additional: Trailing spaces",
		Input: `["word   "]`,
		Want:  `4`,
	},
	{
		Name:  "Additional: One-letter final word",
		Input: `["long phrase x"]`,
		Want:  `1`,
	},
	{
		Name:  "Additional: Leading spaces before single word",
		Input: `["   alphabet"]`,
		Want:  `8`,
	},
	{
		Name:  "Additional: Repeated spaces between words",
		Input: `["one  two   three"]`,
		Want:  `5`,
	},
	{
		Name:  "Additional: Several one-letter words",
		Input: `["a b c"]`,
		Want:  `1`,
	},
	{
		Name:  "Additional: Long single word",
		Input: `["supercalifragilisticexpialidocious"]`,
		Want:  `34`,
	},
	maximumLengthOfLastWordCase(),
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
	s := parseLengthOfLastWordInput(input)
	_ = s

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
