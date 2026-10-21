//go:build ignore

// Difficulty: Easy
// Source: leetcode
// Problem ID: 14

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/duyserveacc/algos/internal/testrunner"
)

/*
======================================================================
Longest Common Prefix

Source: LeetCode
Problem ID: 14
Source URL:
https://leetcode.com/problems/longest-common-prefix/description/

Write a function to find the longest common prefix string amongst an
array of strings.

If there is no common prefix, return an empty string "".

Example 1:

Input: strs = ["flower","flow","flight"]
Output: "fl"

Example 2:

Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.

Constraints:

- 1 <= strs.length <= 200

- 0 <= strs[i].length <= 200

- strs[i] consists of only lowercase English letters if it is
  non-empty.
======================================================================
*/

// Problem-specific support.
//
// Keep shared judge types and parsing or formatting helpers here.
// Name adapters after the problem. Put setup used only by the current
// attempt at the top of solveCurrentAttempt below.
//
// This standalone practice archive follows the workbench template.

func parseLongestCommonPrefixInput(input string) []string {
	var args []json.RawMessage
	if err := json.Unmarshal([]byte(input), &args); err != nil {
		panic(
			fmt.Errorf("parse longest-common-prefix input: %w", err),
		)
	}
	if len(args) != 1 {
		panic(
			fmt.Sprintf("parse longest-common-prefix input: got %d "+
				"arguments, want 1", len(args)),
		)
	}
	var strs []string
	argumentError0 := json.Unmarshal(args[0], &strs)
	if argumentError0 != nil {
		panic(
			fmt.Errorf("parse longest-common-prefix argument strs: "+
				"%w", argumentError0),
		)
	}
	return strs
}

func formatLongestCommonPrefixOutput(prefix string) string {
	output, err := json.Marshal(prefix)
	if err != nil {
		panic(
			fmt.Errorf(
				"format longest-common-prefix output: %w",
				err,
			),
		)
	}
	return string(output)
}

func maximumLongestCommonPrefixCase() testrunner.Case {
	prefix := strings.Repeat("a", 199)
	strs := make([]string, 200)
	for index := range strs {
		strs[index] = prefix + "a"
	}
	strs[len(strs)-1] = prefix + "b"

	input, err := json.Marshal([]any{strs})
	if err != nil {
		panic(
			fmt.Errorf("build maximum longest-common-prefix input: "+
				"%w", err),
		)
	}
	want, err := json.Marshal(prefix)
	if err != nil {
		panic(
			fmt.Errorf("build maximum longest-common-prefix output: "+
				"%w", err),
		)
	}
	return testrunner.Case{
		Name:  "Boundary: Maximum string count and length",
		Input: string(input),
		Want:  string(want),
	}
}

// sampleCases contains the examples that every registered solution
// must pass. Inputs encode the judge arguments as a positional JSON
// array.
var sampleCases = []testrunner.Case{
	{
		Name:  "Official example 1",
		Input: `[["flower","flow","flight"]]`,
		Want:  `"fl"`,
	},
	{
		Name:  "Official example 2",
		Input: `[["dog","racecar","car"]]`,
		Want:  `""`,
	},
	{
		Name:  "Additional: Single string",
		Input: `[["alone"]]`,
		Want:  `"alone"`,
	},
	{
		Name:  "Boundary: Single empty string",
		Input: `[[""]]`,
		Want:  `""`,
	},
	{
		Name:  "Additional: Empty member",
		Input: `[["","abc"]]`,
		Want:  `""`,
	},
	{
		Name:  "Additional: Prefix ends inside words",
		Input: `[["interstellar","internet","internal"]]`,
		Want:  `"inter"`,
	},
	{
		Name:  "Additional: Every string is identical",
		Input: `[["same","same","same"]]`,
		Want:  `"same"`,
	},
	{
		Name:  "Additional: First string is the complete prefix",
		Input: `[["a","ab","abc"]]`,
		Want:  `"a"`,
	},
	{
		Name:  "Additional: Shortest string appears in the middle",
		Input: `[["prefix","pre","prevent"]]`,
		Want:  `"pre"`,
	},
	{
		Name:  "Additional: Empty member appears last",
		Input: `[["prefix","pre",""]]`,
		Want:  `""`,
	},
	{
		Name:  "Adversarial: Divergence at final character",
		Input: `[["abcdefgh","abcdefgi","abcdefgj"]]`,
		Want:  `"abcdefg"`,
	},
	maximumLongestCommonPrefixCase(),
}

// archiveReady is true after the current attempt, three agent
// solutions, analysis, and regression coverage are complete and
// verified.
const archiveReady = true

// Registered implementations run against the same cases.
var currentAttemptSolution = testrunner.Solution{
	Name: "current attempt",
	Run:  solveCurrentAttempt,
}

// agentSolutions contains direct, improved, and preferred
// implementations.
var agentSolutions = []testrunner.Solution{
	{
		Name: "test candidate prefixes",
		Run:  solveByTestingCandidatePrefixes,
	},
	{
		Name: "reduce pairwise prefixes",
		Run:  solveByReducingPairwisePrefixes,
	},
	{
		Name: "scan character columns",
		Run:  solveByScanningCharacterColumns,
	},
}

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
	strs := parseLongestCommonPrefixInput(input)
	longest := ""

	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) {
				return formatLongestCommonPrefixOutput(longest)
			}
			if strs[0][i] != strs[j][i] {
				return formatLongestCommonPrefixOutput(longest)
			}

		}
		longest += string(strs[0][i])
	}

	// Implement your algorithm with the typed values above.
	return formatLongestCommonPrefixOutput(longest)
}

// ===================================================================
// AGENT SOLUTIONS
//
// Direct, improved, and preferred implementations authored during
// review.
// ===================================================================

// solveByTestingCandidatePrefixes tries every prefix of the first
// string from longest to shortest and returns the first candidate
// shared by every string.
func solveByTestingCandidatePrefixes(input string) string {
	strs := parseLongestCommonPrefixInput(input)

	for prefixLength := len(strs[0]); prefixLength >= 0; prefixLength-- {
		candidate := strs[0][:prefixLength]
		shared := true
		for _, str := range strs[1:] {
			if !strings.HasPrefix(str, candidate) {
				shared = false
				break
			}
		}
		if shared {
			return formatLongestCommonPrefixOutput(candidate)
		}
	}

	panic(
		"longest common prefix: empty prefix must match every string",
	)
}

// solveByReducingPairwisePrefixes intersects the accumulated prefix
// with each remaining string. The prefix can only stay the same or
// shrink.
func solveByReducingPairwisePrefixes(input string) string {
	strs := parseLongestCommonPrefixInput(input)
	prefix := strs[0]

	for _, str := range strs[1:] {
		prefix = prefix[:commonPrefixLength(prefix, str)]
		if prefix == "" {
			break
		}
	}

	return formatLongestCommonPrefixOutput(prefix)
}

func commonPrefixLength(left string, right string) int {
	limit := min(len(left), len(right))
	for index := 0; index < limit; index++ {
		if left[index] != right[index] {
			return index
		}
	}
	return limit
}

// solveByScanningCharacterColumns checks one character position
// across every string before admitting that position to the shared
// prefix.
func solveByScanningCharacterColumns(input string) string {
	strs := parseLongestCommonPrefixInput(input)
	return formatLongestCommonPrefixOutput(longestCommonPrefix(strs))
}

// longestCommonPrefix is the judge-facing wrapper for the preferred
// approach.
func longestCommonPrefix(strs []string) string {
	for column := 0; column < len(strs[0]); column++ {
		expected := strs[0][column]
		for row := 1; row < len(strs); row++ {
			if column >= len(strs[row]) ||
				strs[row][column] != expected {
				return strs[0][:column]
			}
		}
	}
	return strs[0]
}

/*
======================================================================
AGENT REVIEW AND COMPLEXITY ANALYSIS

Current attempt:
- What is correct: The attempt performs a vertical scan. It admits a
  character
from strs[0] only after every other string has a character at that
index and
that character matches. It returns immediately on a short string or
mismatch,
so it is correct for every input allowed by the problem. The submitted
  solveCurrentAttempt algorithm body is preserved verbatim.
- First failing assumption or bottleneck: No correctness failure
  remains under
the stated non-empty-array and lowercase-English-letter constraints.
The
  avoidable cost is longest += string(strs[0][i]): immutable string
concatenation copies the accumulated prefix on every successful
column. The
trailing implementation reminder is also stale, but remains untouched
as
  part of the preserved body.
- Relevant concept and invariant: See
  ../CONCEPTS.md#common-prefix-scan. Before each outer iteration,
longest is exactly strs[0][:i], and every processed string shares it.
The
inner loop either proves column i is shared or returns the
already-proven
  prefix, so the invariant is preserved.
- Concrete example: For ["flower","flow","flight"], i=0 compares f
  with f
and f, then longest becomes "f". At i=1, both comparisons match l, so
longest
becomes "fl". At i=2, strs[1] matches o but strs[2] contains i, so the
  function returns "fl".
- Time complexity: Let S be the encoded input size, including string
  count and
characters, and P the returned prefix length. JSON parsing and
comparisons
  take Theta(S) in the worst case; repeated prefix construction adds
Theta(P^2) copied bytes. Best time is Theta(S) when P is constant.
With no
input distribution, no tighter average bound is justified; worst time
is
  Theta(S+P^2).
- Auxiliary-space complexity: The algorithm retains a prefix string
  whose
successive allocations total Theta(P^2) bytes over the call, with
Theta(P)
live at once. Parsing allocates Theta(S), and output formatting
allocates
  Theta(P).
- Relationship to agent solutions: Solution 1 searches whole candidate
prefixes rather than proving one column at a time. Solution 2
repeatedly
intersects a running prefix with one complete string. Solution 3 uses
the
same vertical-scan invariant as the current attempt, but returns a
slice of
  strs[0] and avoids incremental prefix allocations; it is a separate
  implementation, not a different underlying algorithm.

Solution 1 - Test candidate prefixes (direct):
- Correctness: Every common prefix must be a prefix of strs[0].
  Candidates are
examined from len(strs[0]) down to zero, and a candidate is returned
only if
  every remaining string starts with it. Therefore the first returned
candidate is common and no longer common candidate exists. The empty
prefix
  guarantees termination.
- Concrete example: For ["prefix","pre","prevent"], candidate "prefix"
fails against "pre"; then "prefi", "pref", and "pre" are tested in
order.
  "pre" matches both remaining strings and is returned.
- Time complexity: Parsing is Theta(S). Accepting the first candidate
  is
  Theta(S); repeatedly rejecting near-matching candidates can compare
  1+2+...+L characters for each of n strings, so worst time is
O(S+n*L^2), where L=len(strs[0]). No tighter average follows without
an
  input distribution.
- Auxiliary-space complexity: Candidate slicing and scan state use
  O(1)
algorithmic space. Parsing uses Theta(S), and formatting the
P-character
  answer uses Theta(P).

Solution 2 - Reduce pairwise prefixes (improved):
- Correctness: Initially prefix equals the complete first string,
  which is the
longest prefix shared by the one processed string. commonPrefixLength
then
replaces it with the exact intersection of that prefix and the next
string.
By induction, after every iteration prefix is the longest prefix
shared by
all processed strings; after the last string it is the required
answer.
- Concrete example: For ["interstellar","internet","internal"], prefix
starts as "interstellar". Intersecting it with "internet" yields
"inter".
Intersecting "inter" with "internal" leaves "inter", which is
returned.
- Time complexity: Best, average, and worst time are Theta(S),
  including JSON
parsing. Each pairwise comparison advances once through at most the
current
  prefix and the next string, and the prefix never grows.
- Auxiliary-space complexity: The algorithm uses O(1) scan state and
  string
  slices. Parsing uses Theta(S), and output formatting uses Theta(P).

Solution 3 - Scan character columns (preferred):
- Correctness: Before column c, all strings are known to share
  strs[0][:c]. If
any string ends at c or has a different byte, no common prefix can
include
column c, so strs[0][:c] is maximal. Otherwise every string shares the
next
byte and the invariant extends to c+1. Completing the first string
proves it
  is itself the longest common prefix.
- Concrete example: For ["abcdefgh","abcdefgi","abcdefgj"], columns 0
through 6 match across all three strings. At column 7 the bytes are h,
i,
  and j; the scan returns strs[0][:7], which is "abcdefg".
- Time complexity: Best, average, and worst time are Theta(S),
  including JSON
parsing. The scan compares only columns that exist up to the first
decisive
  mismatch, while parsing necessarily reads all input characters.
- Auxiliary-space complexity: The algorithm uses O(1) state and
  returns a
substring view. Parsing uses Theta(S), and formatting the P-character
result
  uses Theta(P).

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
