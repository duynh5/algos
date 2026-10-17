//go:build ignore

// Difficulty: Easy
// Source: leetcode
// Problem ID: 69

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/duyserveacc/algos/internal/testrunner"
)

/*
===============================================================================
Sqrt(x)

Source: LeetCode
Problem ID: 69
Source URL: https://leetcode.com/problems/sqrtx/description/

Given a non-negative integer x, return the square root of x rounded down to the nearest integer. The returned integer should be non-negative as well.

You must not use any built-in exponent function or operator.

- For example, do not use pow(x, 0.5) in c++ or x ** 0.5 in python.

Example 1:

Input: x = 4
Output: 2
Explanation: The square root of 4 is 2, so we return 2.

Example 2:

Input: x = 8
Output: 2
Explanation: The square root of 8 is 2.82842..., and since we round it down to the nearest integer, 2 is returned.

Constraints:

- 0 <= x <= 2^31 - 1
===============================================================================
*/

// Problem-specific support.
//
// Keep shared judge types and parsing or formatting helpers here. Name adapters
// after the problem. Put setup used only by the current attempt at the top of
// solveCurrentAttempt below.
//
// This intentionally unsolved practice archive follows the workbench template.

func parseSqrtxInput(input string) int {
	var args []json.RawMessage
	if err := json.Unmarshal([]byte(input), &args); err != nil {
		panic(fmt.Errorf("parse sqrtx input: %w", err))
	}
	if len(args) != 1 {
		panic(fmt.Sprintf("parse sqrtx input: got %d arguments, want 1", len(args)))
	}
	var x int
	argumentError0 := json.Unmarshal(args[0], &x)
	if argumentError0 != nil {
		panic(fmt.Errorf("parse sqrtx argument x: %w", argumentError0))
	}
	return x
}

func formatSqrtxOutput(result int) string {
	return testrunner.FormatJSONOutput(result)
}

// sampleCases contains the examples that every registered solution must pass.
// Inputs encode the judge arguments as a positional JSON array.
var sampleCases = []testrunner.Case{
	{
		Name:  "Official example 1",
		Input: `[4]`,
		Want:  `2`,
	},
	{
		Name:  "Official example 2",
		Input: `[8]`,
		Want:  `2`,
	},
	{
		Name:  "Additional: Zero",
		Input: `[0]`,
		Want:  `0`,
	},
	{
		Name:  "Additional: One",
		Input: `[1]`,
		Want:  `1`,
	},
	{
		Name:  "Additional: Maximum input",
		Input: `[2147483647]`,
		Want:  `46340`,
	},
	{
		Name:  "Additional: Small non-square two",
		Input: `[2]`,
		Want:  `1`,
	},
	{
		Name:  "Additional: Small non-square three",
		Input: `[3]`,
		Want:  `1`,
	},
	{
		Name:  "Additional: Immediately below perfect square",
		Input: `[15]`,
		Want:  `3`,
	},
	{
		Name:  "Additional: Larger perfect square",
		Input: `[100000000]`,
		Want:  `10000`,
	},
	{
		Name:  "Regression: User walkthrough perfect square",
		Input: `[81]`,
		Want:  `9`,
	},
	{
		Name:  "Adversarial: Immediately below user walkthrough square",
		Input: `[80]`,
		Want:  `8`,
	},
	{
		Name:  "Boundary: Immediately below largest square in range",
		Input: `[2147395599]`,
		Want:  `46339`,
	},
	{
		Name:  "Boundary: Largest square in range",
		Input: `[2147395600]`,
		Want:  `46340`,
	},
	{
		Name:  "Boundary: Immediately above largest square in range",
		Input: `[2147395601]`,
		Want:  `46340`,
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

var agentSolutions = []testrunner.Solution{
	{
		Name: "direct linear scan",
		Run:  solveByLinearScan,
	},
	{
		Name: "improved binary search",
		Run:  solveByBinarySearch,
	},
	{
		Name: "preferred integer Newton iteration",
		Run:  solveByIntegerNewtonIteration,
	},
}

var solutions = append([]testrunner.Solution{currentAttemptSolution}, agentSolutions...)

// =============================================================================
// CURRENT ATTEMPT
//
// Implement your solution here. Keep this body intact when requesting review.
// =============================================================================

func solveCurrentAttempt(input string) string {
	x := parseSqrtxInput(input)
	low, high := 1, x

	if x == 0 {
		return formatSqrtxOutput(0)
	}
	if x == 1 {
		return formatSqrtxOutput(1)
	}

	for low <= high {
		mid := low + (high-low)/2
		if mid == x/mid {
			return formatSqrtxOutput(mid)
		} else if mid > x/mid {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	// Implement your algorithm with the typed values above.
	return formatSqrtxOutput(low - 1)
}

// =============================================================================
// AGENT SOLUTIONS
//
// These functions remain separate from the preserved current attempt so each
// approach can be compared against the same cases.
// =============================================================================

func solveByLinearScan(input string) string {
	x := parseSqrtxInput(input)
	root := 0

	for root+1 <= x/(root+1) {
		root++
	}

	return formatSqrtxOutput(root)
}

func solveByBinarySearch(input string) string {
	x := parseSqrtxInput(input)
	if x < 2 {
		return formatSqrtxOutput(x)
	}

	low, high := 1, x/2
	root := 1
	for low <= high {
		mid := low + (high-low)/2
		if mid <= x/mid {
			root = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return formatSqrtxOutput(root)
}

func solveByIntegerNewtonIteration(input string) string {
	x := int64(parseSqrtxInput(input))
	if x < 2 {
		return formatSqrtxOutput(int(x))
	}

	estimate := x
	for {
		next := (estimate + x/estimate) / 2
		if next >= estimate {
			return formatSqrtxOutput(int(estimate))
		}
		estimate = next
	}
}

/*
===============================================================================
AGENT REVIEW AND COMPLEXITY ANALYSIS

Current attempt:
- Score: 94/100. The implementation is correct, overflow-safe, O(log x) time,
  and O(1) auxiliary space. The deductions are for starting with the looser
  high = x bound and for the correct but less immediately obvious
  mid == x/mid early-return condition.
- What is correct: The attempt handles 0 and 1 explicitly, uses division
  instead of squaring to avoid multiplication overflow, and performs a closed-
  interval binary search. Its mid == x/mid branch does not mean only that mid
  is a perfect square root: with integer division, it correctly recognizes any
  mid for which mid^2 <= x < mid*(mid+1). The attempt passes all supplied and
  regression cases.
- First failing assumption or bottleneck: There is no correctness failure for
  0 <= x <= 2^31-1. The revised low <= high guard now directly describes a
  non-empty closed search interval. The remaining minor tradeoff is that high
  starts at x even though floor(sqrt(x)) <= x/2 for every x >= 2; the wider
  initial interval costs at most a small constant number of iterations and
  does not change the O(log x) complexity.
- Relevant concept and invariant: This is binary search over an answer range;
  see ../concepts/algorithms/binary-search/README.md. Before each iteration,
  every integer below low has been proved feasible and every integer above
  high has been proved infeasible. If mid < x/mid, mid is feasible, so low can
  move to mid+1. If mid > x/mid, mid is infeasible, so high can move to mid-1.
  If the interval empties without an early return, high = low-1 is therefore
  the largest feasible integer and the required rounded-down square root.
- Concrete example: For x = 81, the attempt starts with [low, high] = [1,81].
  It checks mid = 41, 20, and 10; each is too large, so high becomes 40, 19,
  and 9. It then checks 5 and 7; both are too small, so low becomes 6 and then
  8. Mid = 8 is also too small, moving low to 9. Finally, 81/9 = 9, so the
  equality branch returns 9.
- Comparison with solution 1: Both avoid multiplication overflow and seek the
  greatest feasible integer. The scan tests candidates one at a time, whereas
  the attempt uses the monotonic predicate candidate <= x/candidate to discard
  roughly half of the remaining range on every iteration.
- Comparison with solution 2: Both use the same core binary-search algorithm,
  division predicate, conventional low <= high guard, and O(log x) bound.
  Solution 2 starts with the tighter high = x/2 and records the best feasible
  midpoint, while the attempt starts with high = x and returns the final
  boundary low-1. Those are implementation differences, not distinct methods.
- Comparison with solution 3: Both generate guesses and avoid guess*guess.
  The attempt chooses the midpoint of explicit bounds, while Newton iteration
  derives each estimate from estimate and x/estimate and converges much faster
  near the root. Newton iteration does not maintain low/high search bounds.
- Time complexity: Theta(1) for x in {0,1} and when the first midpoint is the
  answer. Otherwise, best case is Theta(1), while average and worst cases are
  O(log x). JSON parsing and output formatting are Theta(1) for this bounded
  scalar input.
- Auxiliary-space complexity: O(1); the algorithm keeps only integer bounds
  and a midpoint, with no input-sized allocation.

Solution 1 - Direct linear scan:
- Correctness: root starts at 0, which is feasible for every allowed x. The
  loop advances exactly while root+1 is feasible. Therefore, when it stops,
  root^2 <= x and (root+1)^2 > x, which uniquely defines floor(sqrt(x)). The
  division comparison is safe because root+1 is never zero.
- Concrete example: For x = 15, root advances through 1 because 1 <= 15/1,
  through 2 because 2 <= 15/2, and through 3 because 3 <= 15/3. Candidate 4
  fails because 4 > 15/4 = 3, so the function returns root = 3.
- Comparison with current attempt: Both use the same feasibility predicate,
  but the scan checks every integer through the answer instead of discarding
  ranges. It is simpler but much slower near the maximum constraint.
- Time complexity: Theta(sqrt(x)) in the average and worst cases because it
  visits every candidate through the answer. The best case is Theta(1) for
  x = 0. Parsing and formatting remain Theta(1) for the bounded scalar.
- Auxiliary-space complexity: O(1); only x and root are stored.

Solution 2 - Improved binary search:
- Correctness: root stores the largest feasible midpoint seen so far. If mid
  is feasible, every smaller value is feasible, so recording mid and searching
  [mid+1, high] cannot lose a better answer. If mid is infeasible, every larger
  value is infeasible, so searching [low, mid-1] is safe. When low > high, no
  unchecked candidate can improve root, making root floor(sqrt(x)).
- Concrete example: For x = 80, [low, high] starts as [1,40]. Mid 20 and 10
  are infeasible, reducing high to 19 and then 9. Mid 5 is feasible, so root
  becomes 5 and low becomes 6. Mid 7 and 8 are feasible, successively making
  root 7 and 8. Mid 9 is infeasible; the interval then empties and 8 returns.
- Comparison with current attempt: This is the same binary-search method as the
  attempt. It differs by starting from the tighter upper bound x/2, tracking
  the best feasible value explicitly, and treating every feasible midpoint
  uniformly instead of returning early when mid == x/mid.
- Time complexity: Theta(1) for x in {0,1}. For x >= 2, the best case is
  Theta(1) when the initial midpoint is the answer; average and worst cases are
  O(log x), because each iteration halves the candidate interval.
- Auxiliary-space complexity: O(1); the search uses four integer variables.

Solution 3 - Preferred integer Newton iteration:
- Correctness: For positive estimate r, (r + x/r)/2 is the integer form of
  averaging r with the complementary factor x/r. Starting from r = x keeps
  estimates at or above floor(sqrt(x)) and strictly decreases them until integer
  rounding prevents another decrease. At that crossing point, r is feasible;
  because no estimate went below floor(sqrt(x)), r is exactly floor(sqrt(x)).
  Using int64 keeps the addition in the update safe for the full input range.
- Concrete example: For x = 81, estimates change as 81 -> 41 -> 21 -> 12 -> 9.
  From 9, the next value is (9 + 81/9)/2 = 9, so the non-decreasing check stops
  and returns 9. For x = 80, the sequence reaches 8; its next value is 9, so 8
  is returned as the last decreasing estimate.
- Comparison with current attempt: The attempt maintains and halves an explicit
  answer interval. Newton iteration instead uses the current estimate and its
  complementary quotient; it usually takes fewer steps but its invariant is
  less immediately visible to readers learning binary search.
- Time complexity: Theta(1) for x in {0,1}. Starting from x, the estimate first
  decreases by roughly half, so the average and worst cases take O(log x)
  iterations overall. Convergence becomes quadratic near the root. Counting
  arbitrary-precision division would also require each update's bit complexity.
- Auxiliary-space complexity: O(1); the iteration stores x, estimate, and next.

AGENT REVIEW AND COMPLEXITY ANALYSIS END
===============================================================================
*/

func main() {
	if err := testrunner.Run(os.Stdout, sampleCases, solutions); err != nil {
		fmt.Fprintln(os.Stderr, "workbench:", err)
		os.Exit(1)
	}
}
