//go:build ignore

// Difficulty: Medium
// Source: leetcode
// Problem ID: 3

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/duyserveacc/algos/internal/testrunner"
)

/*
======================================================================
PASTE THE COMPLETE PROBLEM BELOW THIS LINE

Given a string s, find the length of the longest substring without
duplicate characters.

Example 1:
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3. Note that
"bca" and "cab" are also correct answers.

Example 2:
Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Example 3:
Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence
and not a substring.


Constraints:
0 <= s.length <= 105
s consists of English letters, digits, symbols and spaces.

PASTE THE COMPLETE PROBLEM ABOVE THIS LINE
======================================================================
*/

// Problem-specific support.
//
// This parser keeps serialized sample input out of the algorithm
// functions.

// parseLongestSubstringInput decodes the judge's single string
// argument so the implementation below can work with the problem's
// strongly typed input.
func parseLongestSubstringInput(input string) string {
	var s string
	if err := testrunner.DecodeJSONArgs(input, &s); err != nil {
		panic(fmt.Errorf("parse longest substring input: %w", err))
	}
	return s
}

func formatLongestSubstringWithoutRepeatingCharactersOutput(
	result int,
) string {
	return testrunner.FormatJSONOutput(result)
}

// sampleCases contains the examples that every registered solution
// must pass. Add one entry per example. Encode arguments as a
// positional JSON array.
var sampleCases = []testrunner.Case{
	{
		Name:  "Example 1",
		Input: `["abcabcbb"]`,
		Want:  `3`,
	},
	{
		Name:  "Example 2",
		Input: `["bbbbb"]`,
		Want:  `1`,
	},
	{
		Name:  "Example 3",
		Input: `["pwwkew"]`,
		Want:  `3`,
	},
	{
		Name:  "Empty string",
		Input: `[""]`,
		Want:  `0`,
	},
	{
		Name:  "Single character",
		Input: `["x"]`,
		Want:  `1`,
	},
	{
		Name:  "All characters distinct",
		Input: `["abcdefghijklmnopqrstuvwxyz"]`,
		Want:  `26`,
	},
	{
		Name:  "Window start must not move backward",
		Input: `["abba"]`,
		Want:  `2`,
	},
	{
		Name:  "Repeated character can start a longer suffix",
		Input: `["dvdf"]`,
		Want:  `3`,
	},
	{
		Name:  "Separated repeated characters",
		Input: `["tmmzuxt"]`,
		Want:  `5`,
	},
	{
		Name:  "Spaces are characters",
		Input: `["a b a"]`,
		Want:  `3`,
	},
	{
		Name:  "Letters are case sensitive",
		Input: `["aA1! a"]`,
		Want:  `5`,
	},
	{
		Name:  "Symbols participate in duplicate detection",
		Input: `["!@#$%^&*()!"]`,
		Want:  `10`,
	},
	{
		Name:  "Suffix grows after an earlier duplicate",
		Input: `["anviaj"]`,
		Want:  `5`,
	},
	{
		Name:  "Escaped quote and backslash are characters",
		Input: `["a\"b\\ca"]`,
		Want:  `5`,
	},
	{
		Name:  "Two spaces form a duplicate",
		Input: `["  "]`,
		Want:  `1`,
	},
	{
		Name:  "Maximum length with one repeated character",
		Input: `["` + strings.Repeat("a", 100_000) + `"]`,
		Want:  `1`,
	},
}

// archiveReady is set to true only after the problem, samples,
// preserved current attempt, three agent solutions, analysis, and
// regression coverage are complete and verified.
const archiveReady = true

// Registered implementations run against the same cases.
var currentAttemptSolution = testrunner.Solution{
	Name: "current attempt",
	Run:  solveCurrentAttempt,
}

var agentSolutions = []testrunner.Solution{
	{Name: "direct: restart scan", Run: solveDirectScan},
	{
		Name: "improved: shrinking frequency window",
		Run:  solveShrinkingWindow,
	},
	{
		Name: "preferred: last-seen jump window",
		Run:  solveLastSeenJumpWindow,
	},
}

var solutions = append(
	[]testrunner.Solution{currentAttemptSolution},
	agentSolutions...,
)

// ===================================================================
// CURRENT ATTEMPT
// ===================================================================

// solveCurrentAttempt is preserved as originally submitted.
func solveCurrentAttempt(input string) string {
	s := parseLongestSubstringInput(input)
	i := 0
	m := make(map[byte]int)
	longest := 0

	for j := 0; j < len(s); j++ {
		character := s[j]
		previousIndex, exists := m[character]
		if exists && previousIndex >= i {
			i = previousIndex + 1
		} else {
			longest = max(longest, j-i+1)
		}
		m[character] = j
	}
	// Parse input, implement your algorithm, and return the formatted
	// answer.
	return fmt.Sprintf("%d", longest)
}

// ===================================================================
// AGENT SOLUTIONS
// ===================================================================

func solveDirectScan(input string) string {
	s := parseLongestSubstringInput(input)
	return fmt.Sprintf("%d", longestSubstringDirectScan(s))
}

// longestSubstringDirectScan tries every starting position and
// extends it until the first duplicate. The fixed table is valid
// because the stated character set is byte-sized ASCII.
func longestSubstringDirectScan(s string) int {
	longest := 0
	for start := 0; start < len(s); start++ {
		var seen [256]bool
		for end := start; end < len(s); end++ {
			character := s[end]
			if seen[character] {
				break
			}
			seen[character] = true
			longest = max(longest, end-start+1)
		}
	}
	return longest
}

func solveShrinkingWindow(input string) string {
	s := parseLongestSubstringInput(input)
	return fmt.Sprintf("%d", longestSubstringShrinkingWindow(s))
}

// longestSubstringShrinkingWindow keeps exactly one duplicate-free
// window. A duplicate makes the left boundary advance until that byte
// occurs once again.
func longestSubstringShrinkingWindow(s string) int {
	left := 0
	longest := 0
	frequency := make(map[byte]int)

	for right := 0; right < len(s); right++ {
		character := s[right]
		frequency[character]++
		for frequency[character] > 1 {
			frequency[s[left]]--
			left++
		}
		longest = max(longest, right-left+1)
	}
	return longest
}

func solveLastSeenJumpWindow(input string) string {
	s := parseLongestSubstringInput(input)
	return fmt.Sprintf("%d", longestSubstringLastSeenWindow(s))
}

// longestSubstringLastSeenWindow jumps past an in-window duplicate
// instead of removing the invalid prefix one byte at a time.
func longestSubstringLastSeenWindow(s string) int {
	left := 0
	lastIndex := make(map[byte]int)
	longest := 0

	for right := 0; right < len(s); right++ {
		character := s[right]
		if previousIndex, exists := lastIndex[character]; exists &&
			previousIndex >= left {
			left = previousIndex + 1
		}
		lastIndex[character] = right
		longest = max(longest, right-left+1)
	}
	return longest
}

/*
======================================================================
AGENT REVIEW AND COMPLEXITY ANALYSIS

Current attempt:
- What is correct: The attempt uses a sliding window and remembers
  each byte's
latest index. It checks previousIndex >= i before jumping, so a stale
duplicate cannot move the window backward. Updating a byte's latest
index on
  every visit is also correct. Its choice to skip the maximum update
on an in-window duplicate was safe because the shortened window cannot
be
longer than the valid window immediately before that duplicate. The
attempt
  is preserved unchanged in solveCurrentAttempt.
- First failing assumption or bottleneck: No correctness failure was
  found for
the stated English-letter, digit, symbol, and space inputs. The
implementation
deliberately compares bytes, so it would count UTF-8 bytes rather than
Unicode code points if the input contract were expanded beyond ASCII.
- Relevant concept and invariant: See
  ../CONCEPTS.md#sliding-window and ../CONCEPTS.md#hash-map-or-set.
  After each right index is
processed, s[i:j+1] contains no duplicate byte, i never decreases, m
holds
  the latest processed index for each byte, and longest is the
  greatest valid-window length observed so far.

Solution 1 - Direct restart scan:
- Correctness: For every start index, the inner loop examines
  successive end
indices until the first duplicate. Every window it records is valid.
Any
duplicate-free substring begins at some start, and this scan reaches
its end
  before it can stop, so the maximum recorded length is optimal.
- Concrete example: For "abca", start=0 clears seen, then visits a, b,
  and c,
recording lengths 1, 2, and 3. At end=3, a is already marked, so that
scan
stops. At start=1, it visits b, c, a and again reaches length 3. Later
starts
  produce only 2 and 1, so the answer remains 3.
- Time complexity: In the usual input-size model with an alphabet
  large enough
to contain n distinct characters, the best case is Theta(n), the
average is
distribution-dependent, and the worst case is Theta(n^2). In this
concrete
byte implementation, a scan stops after at most 256 distinct bytes and
each
fixed table also has 256 cells, so its tighter bound is O(256*n) =
O(n).
- Auxiliary-space complexity: Theta(1): each start uses a fresh
  256-entry bool
table on the stack. The shared JSON adapter separately allocates
Theta(n) for
the decoded string, and formatting the integer result allocates a
small output
  string of Theta(log n) bytes.

Solution 2 - Shrinking frequency window:
- Correctness: Adding s[right] can invalidate only that byte's
  uniqueness. The
inner loop removes bytes from the left until its frequency returns to
one;
therefore the window is valid before its length is recorded. Advancing
right
considers every possible ending index, while removing only the minimum
  invalid prefix leaves the longest valid window for that end.
- Concrete example: For "pwwkew", right=0 and right=1 build "pw" and
  longest=2.
Adding the second w at right=2 makes frequency[w]=2. The loop removes
p
(left=1) and then the first w (left=2), restoring window "w". Rights 3
and 4
grow it to "wke", longest=3. At right=5, removing the old w leaves
"kew", so
  the final answer is 3.
- Time complexity: Theta(n) in the best, average, and worst cases.
  Right advances
n times and left advances at most n times across all inner-loop
iterations.
- Auxiliary-space complexity: O(min(n, A)) map entries, where A is the
  byte
alphabet; under this problem's byte input A <= 256, so it is O(1). The
Go map
is the meaningful algorithm allocation. JSON decoding and output
formatting
  have the same wrapper costs described above.

Solution 3 - Last-seen jump window (preferred agent solution):
- Correctness: Before each iteration, the current window is
  duplicate-free.
When s[right] last occurred inside it, setting left to previousIndex+1
removes
that occurrence and the smallest necessary prefix. A stale occurrence
before
left is ignored, so left never retreats. Thus the new window is the
longest
valid window ending at right, and taking the maximum over all right
endpoints
  returns the global optimum.
- Concrete example: For "abba", right=0 stores lastIndex[a]=0 and
  records
length 1. right=1 stores lastIndex[b]=1 and records 2. At right=2, b's
previous index 1 is at or after left=0, so left jumps to 2 and the
window
becomes "b". At right=3, lastIndex[a]=0 is before left, so left stays
2;
  window "ba" has length 2 and the answer is 2.
- Time complexity: Theta(n) in the best, average, and worst cases for
  this
bounded byte keyspace: each index is processed once with constant-size
map
  lookup and update work.
- Auxiliary-space complexity: O(min(n, A)) last-index entries, bounded
  by 256
and therefore O(1) for the stated byte alphabet. The Go map, shared
JSON
  decoding, and formatted output are the meaningful allocations.

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
