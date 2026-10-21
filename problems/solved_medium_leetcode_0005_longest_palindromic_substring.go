//go:build ignore

// Difficulty: Medium
// Source: leetcode
// Problem ID: 5

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
Longest Palindromic Substring

Source: LeetCode
Problem ID: 5
Source URL: https://leetcode.com/problems/longest-palindromic-substring/description/

Given a string s, return the longest palindromic substring in s.

Example 1:

Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.

Example 2:

Input: s = "cbbd"
Output: "bb"

Constraints:

- 1 <= s.length <= 1000

- s consist of only digits and English letters.
===============================================================================
*/

// Problem-specific support.
//
// Keep shared judge types and parsing or formatting helpers here. Name adapters
// after the problem. Put setup used only by the current attempt at the top of
// solveCurrentAttempt below.
//
// This solved archive follows the standalone workbench structure.

func parseLongestPalindromicSubstringInput(input string) string {
	var args []json.RawMessage
	if err := json.Unmarshal([]byte(input), &args); err != nil {
		panic(fmt.Errorf("parse longest-palindromic-substring input: %w", err))
	}
	if len(args) != 1 {
		panic(fmt.Sprintf("parse longest-palindromic-substring input: got %d arguments, want 1", len(args)))
	}
	var s string
	argumentError0 := json.Unmarshal(args[0], &s)
	if argumentError0 != nil {
		panic(fmt.Errorf("parse longest-palindromic-substring argument s: %w", argumentError0))
	}
	return s
}

func formatLongestPalindromicSubstringOutput(result string) string {
	output, err := json.Marshal(result)
	if err != nil {
		panic(fmt.Errorf("format longest-palindromic-substring output: %w", err))
	}
	return string(output)
}

func maximumLongestPalindromicSubstringCase() testrunner.Case {
	s := strings.Repeat("a", 1000)
	input, err := json.Marshal([]any{s})
	if err != nil {
		panic(fmt.Errorf("build maximum longest-palindromic-substring case: %w", err))
	}
	want, err := json.Marshal(s)
	if err != nil {
		panic(fmt.Errorf("build maximum longest-palindromic-substring output: %w", err))
	}
	return testrunner.Case{
		Name:  "Boundary: Maximum length is one palindrome",
		Input: string(input),
		Want:  string(want),
	}
}

// sampleCases contains the examples that every registered solution must pass.
// Inputs encode the judge arguments as a positional JSON array.
var sampleCases = []testrunner.Case{
	{
		Name:  "Official example 1",
		Input: `["babad"]`,
		Want:  `"bab"`,
	},
	{
		Name:  "Official example 2",
		Input: `["cbbd"]`,
		Want:  `"bb"`,
	},
	{
		Name:  "Additional: Minimum length",
		Input: `["a"]`,
		Want:  `"a"`,
	},
	{
		Name:  "Additional: Entire string is a palindrome",
		Input: `["aaaa"]`,
		Want:  `"aaaa"`,
	},
	{
		Name:  "Additional: Long interior palindrome",
		Input: `["forgeeksskeegfor"]`,
		Want:  `"geeksskeeg"`,
	},
	{
		Name:  "Additional: Odd palindrome begins at first character",
		Input: `["racecarx"]`,
		Want:  `"racecar"`,
	},
	{
		Name:  "Additional: Even palindrome is surrounded",
		Input: `["xabccbay"]`,
		Want:  `"abccba"`,
	},
	{
		Name:  "Additional: Even palindrome reaches final character",
		Input: `["xyzzy"]`,
		Want:  `"yzzy"`,
	},
	{
		Name:  "Additional: Digits form the unique longest palindrome",
		Input: `["ab12321cd"]`,
		Want:  `"12321"`,
	},
	{
		Name:  "Regression: No palindrome longer than one character",
		Input: `["abcde"]`,
		Want:  `"a"`,
	},
	{
		Name:  "Regression: Alternating boundary removal skips the suffix",
		Input: `["xyabba"]`,
		Want:  `"abba"`,
	},
	{
		Name:  "Regression: Even palindrome precedes an unrelated suffix",
		Input: `["abbaxyzfset"]`,
		Want:  `"abba"`,
	},
	maximumLongestPalindromicSubstringCase(),
}

// archiveReady is true because the current attempt, three agent solutions,
// analysis, and regression coverage are complete and verified.
const archiveReady = true

// Registered implementations run against the same cases.
var currentAttemptSolution = testrunner.Solution{
	Name: "current attempt",
	Run:  solveCurrentAttempt,
}

var agentSolutions = []testrunner.Solution{
	{Name: "direct: exhaustive substring search", Run: solveExhaustiveSubstringSearch},
	{Name: "improved: dynamic programming", Run: solvePalindromeDynamicProgramming},
	{Name: "preferred: expand around centers", Run: solveExpandAroundCenters},
}

var solutions = append([]testrunner.Solution{currentAttemptSolution}, agentSolutions...)

// =============================================================================
// CURRENT ATTEMPT
//
// The completed attempt uses center expansion with separate odd and even
// centers.
// =============================================================================

func solveCurrentAttempt(input string) string {
	s := parseLongestPalindromicSubstringInput(input)
	bestStart := 0
	bestLength := 1

	for center := 0; center < len(s); center++ {
		oddStart, oddLength := expandCurrentAttemptPalindrome(s, center, center)
		if oddLength > bestLength {
			bestStart = oddStart
			bestLength = oddLength
		}

		evenStart, evenLength := expandCurrentAttemptPalindrome(s, center, center+1)
		if evenLength > bestLength {
			bestStart = evenStart
			bestLength = evenLength
		}
	}

	return formatLongestPalindromicSubstringOutput(s[bestStart : bestStart+bestLength])
}

func expandCurrentAttemptPalindrome(s string, left, right int) (start, length int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right - left - 1
}

// =============================================================================
// AGENT SOLUTIONS
//
// See ../CONCEPTS.md#two-pointers for the center-expansion
// invariant used by the preferred solution.
// =============================================================================

func solveExhaustiveSubstringSearch(input string) string {
	s := parseLongestPalindromicSubstringInput(input)
	return formatLongestPalindromicSubstringOutput(longestPalindromeExhaustive(s))
}

// longestPalindromeExhaustive checks candidates from longest to shortest and
// returns the first palindrome, which also makes ties resolve leftmost.
func longestPalindromeExhaustive(s string) string {
	for length := len(s); length >= 1; length-- {
		for start := 0; start+length <= len(s); start++ {
			if isPalindromeRange(s, start, start+length-1) {
				return s[start : start+length]
			}
		}
	}
	return ""
}

func isPalindromeRange(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func solvePalindromeDynamicProgramming(input string) string {
	s := parseLongestPalindromicSubstringInput(input)
	return formatLongestPalindromicSubstringOutput(longestPalindromeDynamicProgramming(s))
}

// longestPalindromeDynamicProgramming records whether each closed interval is
// a palindrome. Increasing lengths ensure the inner interval is known first.
func longestPalindromeDynamicProgramming(s string) string {
	palindrome := make([][]bool, len(s))
	for start := range palindrome {
		palindrome[start] = make([]bool, len(s))
		palindrome[start][start] = true
	}

	bestStart := 0
	bestLength := 1
	for length := 2; length <= len(s); length++ {
		for start := 0; start+length <= len(s); start++ {
			end := start + length - 1
			palindrome[start][end] = s[start] == s[end] &&
				(length == 2 || palindrome[start+1][end-1])
			if palindrome[start][end] && length > bestLength {
				bestStart = start
				bestLength = length
			}
		}
	}
	return s[bestStart : bestStart+bestLength]
}

func solveExpandAroundCenters(input string) string {
	s := parseLongestPalindromicSubstringInput(input)
	return formatLongestPalindromicSubstringOutput(longestPalindromeByCenterExpansion(s))
}

// longestPalindromeByCenterExpansion tries every character center and every
// gap center. Strictly longer updates preserve the leftmost answer on ties.
func longestPalindromeByCenterExpansion(s string) string {
	bestStart := 0
	bestLength := 1

	for center := 0; center < len(s); center++ {
		oddStart, oddLength := expandPalindrome(s, center, center)
		if oddLength > bestLength {
			bestStart = oddStart
			bestLength = oddLength
		}

		evenStart, evenLength := expandPalindrome(s, center, center+1)
		if evenLength > bestLength {
			bestStart = evenStart
			bestLength = evenLength
		}
	}

	return s[bestStart : bestStart+bestLength]
}

func expandPalindrome(s string, left, right int) (start, length int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right - left - 1
}

/*
===============================================================================
AGENT REVIEW AND COMPLEXITY ANALYSIS

Current attempt:
- What is correct: The attempt parses the input, tries both the character and
  gap center at every index, expands only through matching mirrored bytes, and
  updates the result only for a strictly longer range. It is correct for the
  stated non-empty ASCII input and preserves the leftmost answer on a tie.
- First failing assumption or bottleneck: No correctness failure remains. Its
  worst-case bottleneck is repeated expansion across overlapping palindromes,
  which takes Theta(n^2) comparisons for inputs such as 1,000 equal bytes.
- Relevant concept and invariant: See ../CONCEPTS.md#two-pointers.
  Before each outward step, the range inside left and right is palindromic; a
  matching pair extends it without breaking that property. Unlike the direct
  solution, the attempt constructs palindromes instead of rechecking complete
  substrings. It matches the preferred agent solution's core algorithm and
  invariant, but remains separately implemented. It has the same Theta(n^2)
  time as dynamic programming while avoiding that solution's table.
- Concrete example: For "cbbd", the gap center left=1,right=2 matches b == b,
  so the best range becomes [1,2] = "bb". Expanding to left=0,right=3 compares
  c != d and stops. Every other center produces only a one-byte palindrome.
- Time complexity: Best case Theta(n), average input-dependent, and worst case
  Theta(n^2).
- Auxiliary-space complexity: O(1) for the algorithm. Input decoding and JSON
  output formatting allocate O(n) bytes outside the core search.

Solution 1 - Exhaustive substring search (direct):
- Correctness: Candidates are visited by decreasing length and, within one
  length, increasing start index. isPalindromeRange returns true exactly when
  every mirrored pair in a candidate matches. Therefore the first accepted
  candidate is palindromic, no longer candidate exists, and it is the leftmost
  among equal-length answers.
- Concrete example: For "cbbd", length=4 checks [0,3]: c != d. At length=3,
  [0,2] fails because c != b and [1,3] fails because b != d. At length=2,
  [0,1] fails, then [1,2] compares b == b and succeeds, so the function returns
  "bb" without examining shorter candidates.
- Time complexity: Best case Theta(n) when the whole string is a palindrome.
  The average depends on the input distribution because mismatches stop checks
  early. Worst case O(n^3): O(n^2) candidates may each compare O(n) mirrored
  pairs.
- Auxiliary-space complexity: O(1) for the search; returned substrings share
  the input's backing bytes. The common JSON parser and formatter separately
  allocate O(n) bytes for input and output serialization.

Solution 2 - Dynamic programming (improved):
- Correctness: A range [start,end] is palindromic exactly when its endpoint
  bytes match and either its length is two or [start+1,end-1] is palindromic.
  Processing lengths from small to large makes that inner result available.
  Every range is considered, and updating only for a strictly greater length
  keeps the leftmost maximum.
- Concrete example: For "babad", every length-one diagonal entry starts true.
  No length-two entry is true. At length=3, [0,2] has b == b and [1,1] true, so
  dp[0][2] becomes true and the best range becomes "bab". Range [1,3] similarly
  proves "aba", but its equal length does not replace "bab". No length-four or
  length-five range is true, so "bab" remains the result.
- Time complexity: Theta(n^2) in the best, average, and worst cases because all
  interval endpoints are filled once.
- Auxiliary-space complexity: Theta(n^2) booleans plus O(n) Go slice-row
  headers and allocations. Parsing and output serialization each use O(n)
  additional bytes and do not change the quadratic bound.

Solution 3 - Expand around centers (preferred):
- Correctness: Every palindrome has exactly one center: a character for odd
  length or a gap for even length. The outer loop visits all n character
  centers and all n-1 in-range gap centers. During expansion, the accepted
  range is a palindrome because a matching pair is added around an already
  palindromic interior. Taking the longest accepted range therefore finds the
  global optimum, while strict updates retain the leftmost tied answer.
- Concrete example: For "abbaxyzfset", center=1 first checks the odd range
  [1,1] = "b". Its gap range starts with left=1,right=2: b == b gives "bb".
  Expansion moves to left=0,right=3: a == a gives "abba". The next move reaches
  left=-1 and stops. All other character and gap centers are also tested, but
  none expands past length four, so the result stays "abba".
- Time complexity: Best case Theta(n) when every center stops immediately or
  after constant work. The average is input-distribution dependent. Worst case
  Theta(n^2), as in a string of identical bytes where most centers expand a
  linear distance.
- Auxiliary-space complexity: O(1) for indices and best-range state. The
  result slice shares the input backing bytes; JSON parsing and formatting add
  O(n) allocation outside the core algorithm.

AGENT REVIEW AND COMPLEXITY ANALYSIS END
===============================================================================
*/

func main() {
	if err := testrunner.Run(os.Stdout, sampleCases, solutions); err != nil {
		fmt.Fprintln(os.Stderr, "workbench:", err)
		os.Exit(1)
	}
}
