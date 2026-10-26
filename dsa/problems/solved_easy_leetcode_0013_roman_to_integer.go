//go:build ignore

// Difficulty: Easy
// Source: leetcode
// Problem ID: 13

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/duyserveacc/algos/dsa/internal/testrunner"
)

/*
======================================================================
Roman to Integer

Source: LeetCode
Problem ID: 13
Source URL:
https://leetcode.com/problems/roman-to-integer/description/

Roman numerals are represented by seven different symbols: I, V, X, L,
C, D and M.

Symbol Value
I 1
V 5
X 10
L 50
C 100
D 500
M 1000

For example, 2 is written as II in Roman numeral, just two ones added
together. 12 is written as XII, which is simply X + II. The number 27
is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to
right. However, the numeral for four is not IIII. Instead, the number
four is written as IV. Because the one is before the five we subtract
it making four. The same principle applies to the number nine, which
is written as IX. There are six instances where subtraction is used:

- I can be placed before V (5) and X (10) to make 4 and 9.

- X can be placed before L (50) and C (100) to make 40 and 90.

- C can be placed before D (500) and M (1000) to make 400 and 900.

Given a roman numeral, convert it to an integer.

Example 1:

Input: s = "III"
Output: 3
Explanation: III = 3.

Example 2:

Input: s = "LVIII"
Output: 58
Explanation: L = 50, V= 5, III = 3.

Example 3:

Input: s = "MCMXCIV"
Output: 1994
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.

Constraints:

- 1 <= s.length <= 15

- s contains only the characters ('I', 'V', 'X', 'L', 'C', 'D', 'M').

- It is guaranteed that s is a valid roman numeral in the range [1,
  3999].
======================================================================
*/

// Problem-specific support.
//
// Keep shared judge types and parsing or formatting helpers here.
// Name adapters after the problem. Put setup used only by the current
// attempt at the top of solveCurrentAttempt below.
//
// This standalone practice archive follows the workbench template.

func parseRomanToIntegerInput(input string) string {
	var args []json.RawMessage
	if err := json.Unmarshal([]byte(input), &args); err != nil {
		panic(fmt.Errorf("parse roman-to-integer input: %w", err))
	}
	if len(args) != 1 {
		panic(fmt.Sprintf("parse roman-to-integer input: got %d "+
			"arguments, want 1", len(args)))
	}
	var s string
	argumentError0 := json.Unmarshal(args[0], &s)
	if argumentError0 != nil {
		panic(
			fmt.Errorf(
				"parse roman-to-integer argument s: %w",
				argumentError0,
			),
		)
	}
	return s
}

var subtractiveRomanExpander = strings.NewReplacer(
	"IV", "IIII",
	"IX", "VIIII",
	"XL", "XXXX",
	"XC", "LXXXX",
	"CD", "CCCC",
	"CM", "DCCCC",
)

type romanToken struct {
	symbol string
	value  int
}

var romanTokens = [...]romanToken{
	{symbol: "M", value: 1000},
	{symbol: "CM", value: 900},
	{symbol: "D", value: 500},
	{symbol: "CD", value: 400},
	{symbol: "C", value: 100},
	{symbol: "XC", value: 90},
	{symbol: "L", value: 50},
	{symbol: "XL", value: 40},
	{symbol: "X", value: 10},
	{symbol: "IX", value: 9},
	{symbol: "V", value: 5},
	{symbol: "IV", value: 4},
	{symbol: "I", value: 1},
}

func formatRomanToIntegerOutput(result int) string {
	return testrunner.FormatJSONOutput(result)
}

// sampleCases contains the examples that every registered solution
// must pass. Inputs encode the judge arguments as a positional JSON
// array.
var sampleCases = []testrunner.Case{
	{
		Name:  "Official example 1",
		Input: `["III"]`,
		Want:  `3`,
	},
	{
		Name:  "Official example 2",
		Input: `["LVIII"]`,
		Want:  `58`,
	},
	{
		Name:  "Official example 3",
		Input: `["MCMXCIV"]`,
		Want:  `1994`,
	},
	{
		Name:  "Additional: Minimum numeral",
		Input: `["I"]`,
		Want:  `1`,
	},
	{
		Name:  "Additional: Subtractive pairs CD, XL, and IV",
		Input: `["CDXLIV"]`,
		Want:  `444`,
	},
	{
		Name:  "Additional: Maximum numeral",
		Input: `["MMMCMXCIX"]`,
		Want:  `3999`,
	},
	{
		Name:  "Additional: Subtractive pair IX",
		Input: `["IX"]`,
		Want:  `9`,
	},
	{
		Name:  "Additional: Adjacent subtractive groups",
		Input: `["XLIX"]`,
		Want:  `49`,
	},
	{
		Name:  "Additional: Subtractive pairs CM, XC, and IX",
		Input: `["CMXCIX"]`,
		Want:  `999`,
	},
	{
		Name:  "Additional: Every symbol in descending order",
		Input: `["MDCLXVI"]`,
		Want:  `1666`,
	},
	{
		Name:  "Boundary: Maximum input length",
		Input: `["MMMDCCCLXXXVIII"]`,
		Want:  `3888`,
	},
	{
		Name:  "Regression: Additive prefix before subtractive suffix",
		Input: `["XXIX"]`,
		Want:  `29`,
	},
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
		Name: "expand subtractive pairs",
		Run:  solveByExpandingSubtractivePairs,
	},
	{
		Name: "consume descending tokens",
		Run:  solveByConsumingDescendingTokens,
	},
	{
		Name: "scan from right to left",
		Run:  solveByScanningRightToLeft,
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
	s := parseRomanToIntegerInput(input)
	l := strings.Split(s, "")
	num_mapping := make(map[string]int)
	num_mapping["I"] = 1
	num_mapping["V"] = 5
	num_mapping["X"] = 10
	num_mapping["L"] = 50
	num_mapping["C"] = 100
	num_mapping["D"] = 500
	num_mapping["M"] = 1000
	num_mapping["IV"] = 4
	num_mapping["IX"] = 9
	num_mapping["XL"] = 40
	num_mapping["XC"] = 90
	num_mapping["CD"] = 400
	num_mapping["CM"] = 900

	sum := 0

	for i := 0; i < len(l); i++ {
		if i == len(l)-1 {
			sum += num_mapping[l[i]]
			break
		}

		if (l[i] == "I" && l[i+1] == "V") ||
			(l[i] == "I" && l[i+1] == "X") ||
			(l[i] == "X" && l[i+1] == "L") ||
			(l[i] == "X" && l[i+1] == "C") ||
			(l[i] == "C" && l[i+1] == "D") ||
			(l[i] == "C" && l[i+1] == "M") {
			sum += num_mapping[l[i]+l[i+1]]
			i += 1
			continue
		}

		sum += num_mapping[l[i]]
	}

	// Implement your algorithm with the typed values above.
	return strconv.Itoa(sum)
}

// ===================================================================
// AGENT SOLUTIONS
//
// Direct, improved, and preferred implementations authored during
// review.
// ===================================================================

// solveByExpandingSubtractivePairs rewrites the six special pairs
// into equivalent additive notation, then sums every remaining
// symbol.
func solveByExpandingSubtractivePairs(input string) string {
	s := parseRomanToIntegerInput(input)
	expanded := subtractiveRomanExpander.Replace(s)

	total := 0
	for index := 0; index < len(expanded); index++ {
		total += romanSymbolValue(expanded[index])
	}
	return strconv.Itoa(total)
}

// solveByConsumingDescendingTokens matches the largest legal numeral
// token at the unread prefix. Canonical Roman notation makes this
// choice unambiguous.
func solveByConsumingDescendingTokens(input string) string {
	s := parseRomanToIntegerInput(input)
	total := 0

	for index := 0; index < len(s); {
		matched := false
		for _, token := range romanTokens {
			if strings.HasPrefix(s[index:], token.symbol) {
				total += token.value
				index += len(token.symbol)
				matched = true
				break
			}
		}
		if !matched {
			panic(
				fmt.Sprintf(
					"roman to integer: no token at byte %d",
					index,
				),
			)
		}
	}

	return strconv.Itoa(total)
}

// solveByScanningRightToLeft makes each symbol's sign depend on the
// symbol to its right, which has already been processed.
func solveByScanningRightToLeft(input string) string {
	s := parseRomanToIntegerInput(input)
	return strconv.Itoa(romanToInt(s))
}

// romanToInt is the judge-facing wrapper around the preferred
// approach.
func romanToInt(s string) int {
	total := 0
	rightValue := 0

	for index := len(s) - 1; index >= 0; index-- {
		value := romanSymbolValue(s[index])
		if value < rightValue {
			total -= value
		} else {
			total += value
		}
		rightValue = value
	}

	return total
}

func romanSymbolValue(symbol byte) int {
	switch symbol {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		panic(
			fmt.Sprintf(
				"roman to integer: invalid symbol %q",
				symbol,
			),
		)
	}
}

/*
======================================================================
AGENT REVIEW AND COMPLEXITY ANALYSIS

Current attempt:
- What is correct: The attempt recognizes all six legal subtractive
  pairs and
otherwise adds the current single-symbol value. Advancing i inside a
matched
pair ensures its second symbol is consumed exactly once. The
final-symbol
branch prevents an out-of-range i+1 access. Under the stated guarantee
that
the input is a valid canonical Roman numeral, it returns the correct
value.
  The submitted solveCurrentAttempt body is preserved verbatim.
- First failing assumption or bottleneck: No correctness failure
  remains under
the problem contract. The avoidable cost is allocation: strings.Split
builds
a []string, and the 13-entry map is allocated and populated for every
call.
A long condition also duplicates facts already encoded by the pair
entries.
- Relevant concept and invariant: See
  ../CONCEPTS.md#directional-scan. Before each loop iteration, sum
equals the value of the fully consumed prefix. The next iteration
consumes
either one additive symbol or one complete subtractive pair,
preserving that
  invariant and guaranteeing progress.
- Concrete example: For s="MCMXCIV", i=0 consumes M and sum becomes
  1000;
i=1 consumes CM and advances i past M, making sum 1900; i=3 consumes
XC and
  makes sum 1990; i=5 consumes IV and finishes at 1994.
- Time complexity: Best, average, and worst time are Theta(n): the
  split, map
construction, and scan are all linear or constant-sized work. JSON
parsing
  adds Theta(n) at the wrapper boundary.
- Auxiliary-space complexity: Theta(n) for the []string produced by
  Split,
plus a constant-sized map. JSON parsing separately allocates Theta(n).
- Relationship to agent solutions: Solution 1 also handles subtractive
  pairs
explicitly, but rewrites all six forms before a uniform symbol sum.
Solution
2 consumes the same legal one- or two-symbol tokens from a descending
table.
Solution 3 does not enumerate pairs; it derives each sign from scan
order.

Solution 1 - Expand subtractive pairs (direct):
- Correctness: Each replacement preserves value: for example, IV and
  IIII are
both 4, while CM and DCCCC are both 900. Canonical input contains no
other
subtractive forms. After all replacements, every remaining symbol is
purely
additive, so summing symbol values yields the original numeral's
value.
- Concrete example: MCMXCIV expands to MDCCCCLXXXXIIII. The scan adds
  M=1000,
  D+CCCC=900, L+XXXX=90, and IIII=4, producing 1994.
- Time complexity: Best, average, and worst time are Theta(n). The
  replacement
pass and additive scan each process output proportional to n because
every
original pair expands by at most a constant factor. Parsing is
Theta(n).
- Auxiliary-space complexity: Theta(n) for the expanded string, in
  addition to
Theta(n) workbench JSON parsing. The summation itself uses O(1) space.

Solution 2 - Consume descending tokens (improved):
- Correctness: romanTokens lists every legal canonical token, with
  each
subtractive pair before its one-symbol prefix. At each unread
position, the
first matching token is therefore exactly the next numeral token.
Adding it
and advancing by its length partitions the whole valid input without
overlap
  or omission, so total equals the numeral's value when the scan ends.
- Concrete example: For CDXLIV, index 0 matches CD and total becomes
  400;
index 2 matches XL and total becomes 440; index 4 matches IV and total
  becomes 444; index advances to 6, which equals len(s).
- Time complexity: Best, average, and worst time are Theta(n): at most
  13
constant-length token comparisons occur per consumed token. Parsing is
also
  Theta(n).
- Auxiliary-space complexity: O(1) for the fixed token table and scan
  state.
  The workbench JSON adapter separately allocates Theta(n).

Solution 3 - Scan right to left (preferred):
- Correctness: In a valid Roman numeral, a symbol smaller than the
  symbol
immediately to its right is the leading half of a subtractive pair;
every
other symbol contributes positively. Scanning right to left makes that
right
value available in O(1) state. Each iteration applies the required
sign to
exactly one symbol, so total is the value of the processed suffix;
when the
  scan finishes, it is the value of the complete numeral.
- Concrete example: For MCMXCIV, scanning from the end changes
(symbol,rightValue,total) through (V,5,5), (I,1,4), (C,100,104),
(X,10,94),
  (M,1000,1094), (C,100,994), and (M,1000,1994).
- Time complexity: Best, average, and worst time are Theta(n),
  including the
  workbench parser. Every byte is visited once.
- Auxiliary-space complexity: O(1) for the algorithm. The workbench
  JSON
  adapter allocates Theta(n) outside the judge-facing function.

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
