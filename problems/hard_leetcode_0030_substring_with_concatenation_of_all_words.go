//go:build ignore

// Difficulty: Hard
// Source: leetcode
// Problem ID: 30

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/duyserveacc/algos/internal/testrunner"
)

/*
===============================================================================
Substring with Concatenation of All Words

Source: LeetCode
Problem ID: 30
Source URL: https://leetcode.com/problems/substring-with-concatenation-of-all-words/description/

You are given a string s and an array of strings words. All the strings of words are of the same length.

A concatenated string is a string that exactly contains all the strings of any permutation of words concatenated.

- For example, if words = ["ab","cd","ef"], then "abcdef", "abefcd", "cdabef", "cdefab", "efabcd", and "efcdab" are all concatenated strings. "acdbef" is not a concatenated string because it is not the concatenation of any permutation of words.

Return an array of the starting indices of all the concatenated substrings in s. You can return the answer in any order.

Example 1:

Input: s = "barfoothefoobarman", words = ["foo","bar"]

Output: [0,9]

Explanation:

The substring starting at 0 is "barfoo". It is the concatenation of ["bar","foo"] which is a permutation of words.

The substring starting at 9 is "foobar". It is the concatenation of ["foo","bar"] which is a permutation of words.

Example 2:

Input: s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]

Output: []

Explanation:

There is no concatenated substring.

Example 3:

Input: s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]

Output: [6,9,12]

Explanation:

The substring starting at 6 is "foobarthe". It is the concatenation of ["foo","bar","the"].

The substring starting at 9 is "barthefoo". It is the concatenation of ["bar","the","foo"].

The substring starting at 12 is "thefoobar". It is the concatenation of ["the","foo","bar"].

Constraints:

- 1 <= s.length <= 10^4

- 1 <= words.length <= 5000

- 1 <= words[i].length <= 30

- s and words[i] consist of lowercase English letters.
===============================================================================
*/

// Problem-specific support.
//
// Keep shared judge types and parsing or formatting helpers here. Name adapters
// after the problem. Put setup used only by the current attempt at the top of
// solveCurrentAttempt below.
//
// This intentionally unsolved practice archive follows the workbench template.

func parseSubstringWithConcatenationOfAllWordsInput(input string) (string, []string) {
	var args []json.RawMessage
	if err := json.Unmarshal([]byte(input), &args); err != nil {
		panic(fmt.Errorf("parse substring-with-concatenation-of-all-words input: %w", err))
	}
	if len(args) != 2 {
		panic(fmt.Sprintf("parse substring-with-concatenation-of-all-words input: got %d arguments, want 2", len(args)))
	}
	var s string
	argumentError0 := json.Unmarshal(args[0], &s)
	if argumentError0 != nil {
		panic(fmt.Errorf("parse substring-with-concatenation-of-all-words argument s: %w", argumentError0))
	}
	var words []string
	argumentError1 := json.Unmarshal(args[1], &words)
	if argumentError1 != nil {
		panic(fmt.Errorf("parse substring-with-concatenation-of-all-words argument words: %w", argumentError1))
	}
	return s, words
}

// sampleCases contains the examples that every registered solution must pass.
// Inputs encode the judge arguments as a positional JSON array.
var sampleCases = []testrunner.Case{
	{
		Name:  "Official example 1",
		Input: `["barfoothefoobarman",["foo","bar"]]`,
		Want:  `[0,9]`,
	},
	{
		Name:  "Official example 2",
		Input: `["wordgoodgoodgoodbestword",["word","good","best","word"]]`,
		Want:  `[]`,
	},
	{
		Name:  "Official example 3",
		Input: `["barfoofoobarthefoobarman",["bar","foo","the"]]`,
		Want:  `[6,9,12]`,
	},
	{
		Name:  "Additional: Repeated required word",
		Input: `["wordgoodgoodgoodbestword",["word","good","best","good"]]`,
		Want:  `[8]`,
	},
	{
		Name:  "Additional: Overlapping starts",
		Input: `["aaaaaa",["aa","aa"]]`,
		Want:  `[0,1,2]`,
	},
	{
		Name:  "Additional: Combined words longer than source",
		Input: `["a",["a","a"]]`,
		Want:  `[]`,
	},
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
	s, words := parseSubstringWithConcatenationOfAllWordsInput(input)
	_ = s
	_ = words

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
