//go:build ignore

// Difficulty: Easy
// Source: leetcode
// Problem ID: 9

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/duyserveacc/algos/internal/testrunner"
)

/*
===============================================================================
Palindrome Number

Source: LeetCode
Problem ID: 9
Source URL: https://leetcode.com/problems/palindrome-number/description/

Given an integer x, return true if x is a palindrome, and false otherwise.

Example 1:

Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.

Example 2:

Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.

Example 3:

Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.

Constraints:

- -2^31 <= x <= 2^31 - 1

Follow up: Could you solve it without converting the integer to a string?
===============================================================================
*/

// Problem-specific support.
//
// Keep shared judge types and parsing or formatting helpers here. Name adapters
// after the problem. Put setup used only by the current attempt at the top of
// solveCurrentAttempt below.
//
// This standalone practice archive follows the workbench template.

func parsePalindromeNumberInput(input string) int {
	var args []json.RawMessage
	if err := json.Unmarshal([]byte(input), &args); err != nil {
		panic(fmt.Errorf("parse palindrome-number input: %w", err))
	}
	if len(args) != 1 {
		panic(fmt.Sprintf("parse palindrome-number input: got %d arguments, want 1", len(args)))
	}
	var x int
	argumentError0 := json.Unmarshal(args[0], &x)
	if argumentError0 != nil {
		panic(fmt.Errorf("parse palindrome-number argument x: %w", argumentError0))
	}
	return x
}

// sampleCases contains the examples that every registered solution must pass.
// Inputs encode the judge arguments as a positional JSON array.
var sampleCases = []testrunner.Case{
	{
		Name:  "Official example 1",
		Input: `[121]`,
		Want:  `true`,
	},
	{
		Name:  "Official example 2",
		Input: `[-121]`,
		Want:  `false`,
	},
	{
		Name:  "Official example 3",
		Input: `[10]`,
		Want:  `false`,
	},
	{
		Name:  "Additional: Zero",
		Input: `[0]`,
		Want:  `true`,
	},
	{
		Name:  "Additional: Two equal digits",
		Input: `[11]`,
		Want:  `true`,
	},
	{
		Name:  "Additional: Large odd-length palindrome",
		Input: `[2147447412]`,
		Want:  `true`,
	},
	{
		Name:  "Additional: Single positive digit",
		Input: `[7]`,
		Want:  `true`,
	},
	{
		Name:  "Additional: Two different digits",
		Input: `[12]`,
		Want:  `false`,
	},
	{
		Name:  "Additional: Even-length palindrome",
		Input: `[1221]`,
		Want:  `true`,
	},
	{
		Name:  "Additional: Repeated zeroes inside palindrome",
		Input: `[1001]`,
		Want:  `true`,
	},
	{
		Name:  "Regression: Repeated digits before midpoint",
		Input: `[11111]`,
		Want:  `true`,
	},
	{
		Name:  "Additional: Near-palindrome",
		Input: `[12331]`,
		Want:  `false`,
	},
	{
		Name:  "Additional: Positive number ending in zero",
		Input: `[1000000000]`,
		Want:  `false`,
	},
	{
		Name:  "Boundary: Minimum 32-bit integer",
		Input: `[-2147483648]`,
		Want:  `false`,
	},
	{
		Name:  "Boundary: Maximum 32-bit integer",
		Input: `[2147483647]`,
		Want:  `false`,
	},
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
	{Name: "compare string ends", Run: solveByComparingStringEnds},
	{Name: "reverse every digit", Run: solveByReversingEveryDigit},
	{Name: "reverse half the digits", Run: solveByReversingHalfDigits},
}

var solutions = append([]testrunner.Solution{currentAttemptSolution}, agentSolutions...)

// =============================================================================
// CURRENT ATTEMPT
//
// Implement your solution here. Keep this body intact when requesting review.
// =============================================================================

func solveCurrentAttempt(input string) string {
	x := parsePalindromeNumberInput(input)
	remaining := x
	reversed := 0

	for remaining > 0 {
		reversed = reversed*10 + remaining%10
		remaining = remaining / 10
	}
	return strconv.FormatBool(reversed == x)
}

// =============================================================================
// AGENT SOLUTIONS
//
// Direct, improved, and preferred implementations authored during review.
// =============================================================================

// solveByComparingStringEnds converts the number to decimal and compares each
// symmetric pair of digits.
func solveByComparingStringEnds(input string) string {
	x := parsePalindromeNumberInput(input)
	return strconv.FormatBool(palindromeByComparingStringEnds(x))
}

func palindromeByComparingStringEnds(x int) bool {
	if x < 0 {
		return false
	}
	digits := strconv.Itoa(x)
	for left, right := 0, len(digits)-1; left < right; left, right = left+1, right-1 {
		if digits[left] != digits[right] {
			return false
		}
	}
	return true
}

// solveByReversingEveryDigit reconstructs the entire decimal number backward.
func solveByReversingEveryDigit(input string) string {
	x := parsePalindromeNumberInput(input)
	return strconv.FormatBool(palindromeByReversingEveryDigit(x))
}

func palindromeByReversingEveryDigit(x int) bool {
	if x < 0 {
		return false
	}
	original := int64(x)
	remaining := original
	var reversed int64
	for remaining > 0 {
		reversed = reversed*10 + remaining%10
		remaining /= 10
	}
	return reversed == original
}

// solveByReversingHalfDigits avoids strings and stops once the reversed suffix
// contains at least as many digits as the unread prefix.
func solveByReversingHalfDigits(input string) string {
	x := parsePalindromeNumberInput(input)
	return strconv.FormatBool(isPalindrome(x))
}

// isPalindrome is the judge-facing wrapper around the preferred approach.
func isPalindrome(x int) bool {
	if x < 0 || x != 0 && x%10 == 0 {
		return false
	}

	reversedHalf := 0
	for x > reversedHalf {
		reversedHalf = reversedHalf*10 + x%10
		x /= 10
	}
	return x == reversedHalf || x == reversedHalf/10
}

/*
===============================================================================
AGENT REVIEW AND COMPLEXITY ANALYSIS

Current attempt:
- What is correct: The attempt extracts the final digit with modulo 10, appends
  it to reversed, and removes it from remaining with division by 10. When the
  loop ends, reversed is the complete decimal reversal of x. Comparing it with
  x correctly recognizes every nonnegative palindrome. Zero returns true, and
  negative inputs return false because the loop is skipped and zero cannot
  equal a negative x. The submitted solveCurrentAttempt body is preserved.
- First failing assumption or bottleneck: No correctness failure remains on
  the repository's 64-bit Go target under the stated 32-bit input constraint.
  The portability limit is that reversed uses int: on a 32-bit Go target, the
  reverse of a valid input such as 2147483647 exceeds int. Using int64 as in
  Solution 2 or reversing only half as in Solution 3 removes that dependency.
- Relevant concept and invariant: See
  concepts/foundations/integer-digits/README.md. After k iterations, remaining
  is x with its final k decimal digits removed, while reversed contains those k
  digits in reverse order.
- Concrete example: For x=1221, (remaining,reversed) changes from (1221,0) to
  (122,1), (12,12), (1,122), and (0,1221). Since reversed equals the unchanged
  x, the attempt returns true. For x=10, the final reversed value is 1, so it
  returns false.
- Time complexity: For a nonnegative d-digit number, best, average, and worst
  time are Theta(d). Negative inputs return in Theta(1). JSON parsing adds
  Theta(d) work at the wrapper boundary.
- Auxiliary-space complexity: O(1) for the arithmetic algorithm. The JSON
  adapter separately allocates Theta(d) for its serialized input.
- Relationship to agent solutions: Solution 2 uses the same full arithmetic
  reversal algorithm. Its material difference is int64 intermediates, which
  make the 32-bit input contract safe on every Go architecture. Solution 1
  compares string endpoints, while Solution 3 reverses only half the digits.

Solution 1 - Compare string ends (direct):
- Correctness: Each iteration compares digits at equally distant positions
  from the two ends. A mismatch proves the number is not a palindrome. If all
  pairs match, every digit equals its mirrored partner, which is exactly the
  palindrome condition. Negative values are rejected before comparison.
- Concrete example: For x=12321, digits="12321". (left,right)=(0,4) compares
  1 with 1, then (1,3) compares 2 with 2. The pointers meet at index 2, whose
  digit needs no partner comparison, so the function returns true.
- Time complexity: Theta(d) in the best, average, and worst cases when decimal
  conversion is included; although comparison can stop at its first mismatch,
  strconv.Itoa first emits all d digits. Parsing also takes Theta(d).
- Auxiliary-space complexity: Theta(d) for the decimal string, plus Theta(d)
  for the workbench JSON input. The pointer loop itself uses O(1) space.

Solution 2 - Reverse every digit (improved):
- Correctness: Each iteration removes remaining's last digit and appends it to
  reversed, so after k iterations reversed contains exactly the last k digits
  of original in reverse order. Once remaining is zero, equality with original
  holds exactly when the decimal digits are palindromic. int64 safely contains
  the reverse of every value allowed by the 32-bit input constraint.
- Concrete example: For x=1221, (remaining,reversed) changes from (1221,0) to
  (122,1), (12,12), (1,122), and (0,1221). The rebuilt value equals original,
  so the result is true. For x=10 it becomes 1, not 10, so the result is false.
- Time complexity: Theta(d) in the best, average, and worst nonnegative cases;
  negative values return in Theta(1). Workbench input parsing is Theta(d).
- Auxiliary-space complexity: O(1) for the arithmetic algorithm. The workbench
  JSON adapter separately allocates Theta(d) for its serialized input.

Solution 3 - Reverse half the digits (preferred):
- Correctness: reversedHalf is the reverse of the suffix already removed from
  x. The loop stops when that suffix has at least as many digits as the prefix.
  An even-length palindrome has x==reversedHalf; an odd-length palindrome has
  x==reversedHalf/10 because the middle digit is irrelevant. A nonzero number
  ending in zero cannot start with zero and is rejected before the loop.
- Concrete example: For x=12321, (x,reversedHalf) changes from (12321,0) to
  (1232,1), (123,12), and (12,123). The suffix has crossed the prefix, and
  12==123/10, so the number is palindromic. For x=10, the trailing-zero guard
  immediately returns false.
- Time complexity: Theta(d) in the best, average, and worst cases up to a
  constant factor of one half for accepted nonnegative inputs; negative and
  nonzero trailing-zero inputs return in Theta(1). Parsing remains Theta(d).
- Auxiliary-space complexity: O(1) for the algorithm. The workbench JSON
  adapter allocates Theta(d) outside the judge-facing function.

AGENT REVIEW AND COMPLEXITY ANALYSIS END
===============================================================================
*/

func main() {
	if err := testrunner.Run(os.Stdout, sampleCases, solutions); err != nil {
		fmt.Fprintln(os.Stderr, "workbench:", err)
		os.Exit(1)
	}
}
