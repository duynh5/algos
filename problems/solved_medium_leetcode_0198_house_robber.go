//go:build ignore

// Difficulty: Medium
// Source: leetcode
// Problem ID: 198

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/duyserveacc/algos/internal/testrunner"
)

/*
===============================================================================
House Robber

Source: LeetCode
Problem ID: 198
Source URL: https://leetcode.com/problems/house-robber/description/

You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security systems connected and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.

Example 1:

Input: nums = [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
Total amount you can rob = 1 + 3 = 4.

Example 2:

Input: nums = [2,7,9,3,1]
Output: 12
Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
Total amount you can rob = 2 + 9 + 1 = 12.

Constraints:

- 1 <= nums.length <= 100

- 0 <= nums[i] <= 400
===============================================================================
*/

// Problem-specific support.
//
// Keep shared judge types and parsing or formatting helpers here. Name adapters
// after the problem. Put setup used only by the current attempt at the top of
// solveCurrentAttempt below.
//
// This reviewed practice archive follows the standalone workbench structure.

func parseHouseRobberInput(input string) []int {
	var args []json.RawMessage
	if err := json.Unmarshal([]byte(input), &args); err != nil {
		panic(fmt.Errorf("parse house-robber input: %w", err))
	}
	if len(args) != 1 {
		panic(fmt.Sprintf("parse house-robber input: got %d arguments, want 1", len(args)))
	}
	var nums []int
	argumentError0 := json.Unmarshal(args[0], &nums)
	if argumentError0 != nil {
		panic(fmt.Errorf("parse house-robber argument nums: %w", argumentError0))
	}
	return nums
}

func maximumHouseRobberCase() testrunner.Case {
	nums := make([]int, 100)
	for index := range nums {
		nums[index] = 400
	}
	input, err := json.Marshal([]any{nums})
	if err != nil {
		panic(fmt.Errorf("build maximum house-robber input: %w", err))
	}
	return testrunner.Case{Name: "Boundary: Maximum houses and values", Input: string(input), Want: `20000`}
}

func formatHouseRobberOutput(result int) string {
	return testrunner.FormatJSONOutput(result)
}

// sampleCases contains the examples that every registered solution must pass.
// Inputs encode the judge arguments as a positional JSON array.
var sampleCases = []testrunner.Case{
	{
		Name:  "Official example 1",
		Input: `[[1,2,3,1]]`,
		Want:  `4`,
	},
	{
		Name:  "Official example 2",
		Input: `[[2,7,9,3,1]]`,
		Want:  `12`,
	},
	{
		Name:  "Additional: Single house",
		Input: `[[5]]`,
		Want:  `5`,
	},
	{
		Name:  "Additional: All zero values",
		Input: `[[0,0,0]]`,
		Want:  `0`,
	},
	{
		Name:  "Additional: Skip two adjacent large houses",
		Input: `[[5,10,10,5]]`,
		Want:  `15`,
	},
	{
		Name:  "Minimum: Single maximum-value house",
		Input: `[[400]]`,
		Want:  `400`,
	},
	{
		Name:  "Additional: Two houses choose the larger",
		Input: `[[1,400]]`,
		Want:  `400`,
	},
	{
		Name:  "Adversarial: Greedy largest-first choice is suboptimal",
		Input: `[[2,3,2]]`,
		Want:  `4`,
	},
	{
		Name:  "Adversarial: Equal middle values favor both ends",
		Input: `[[2,1,1,2]]`,
		Want:  `4`,
	},
	{
		Name:  "Adversarial: Optimal choice excludes both endpoints",
		Input: `[[1,100,1]]`,
		Want:  `100`,
	},
	{
		Name:  "Additional: Zero-value gaps allow every positive house",
		Input: `[[5,0,6,0,7]]`,
		Want:  `18`,
	},
	{
		Name:  "Custom val",
		Input: `[[5,0,6,0,7]]`,
		Want:  `18`,
	},
	maximumHouseRobberCase(),
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
		Name: "direct compatible-predecessor dynamic programming",
		Run:  solveByCompatiblePredecessors,
	},
	{
		Name: "improved top-down memoization",
		Run:  solveByTopDownMemoization,
	},
	{
		Name: "preferred rolling dynamic programming",
		Run:  solveByRollingDynamicProgramming,
	},
}

var solutions = append([]testrunner.Solution{currentAttemptSolution}, agentSolutions...)

// =============================================================================
// CURRENT ATTEMPT
//
// Implement your solution here. Keep this body intact when requesting review.
// =============================================================================

func solveCurrentAttempt(input string) string {
	nums := parseHouseRobberInput(input)
	_ = nums

	if len(nums) == 0 {
		return formatHouseRobberOutput(0)
	}
	if len(nums) == 1 {
		return formatHouseRobberOutput(nums[0])
	}

	best_t1, best_t2 := max(nums[1], nums[0]), nums[0]
	for visit := 2; visit < len(nums); visit++ {
		new_best := max(best_t1, nums[visit]+best_t2)
		best_t2 = best_t1
		best_t1 = new_best
	}
	// Implement your algorithm with the typed values above.
	return formatHouseRobberOutput(max(best_t1, best_t2))
}

// =============================================================================
// AGENT SOLUTIONS
//
// These functions remain separate from the preserved current attempt so each
// approach can be compared against the same cases.
// =============================================================================

func solveByCompatiblePredecessors(input string) string {
	nums := parseHouseRobberInput(input)
	bestEndingAt := make([]int, len(nums))
	bestOverall := 0

	for currentHouse, amount := range nums {
		bestCompatible := 0
		for previousHouse := 0; previousHouse <= currentHouse-2; previousHouse++ {
			bestCompatible = max(bestCompatible, bestEndingAt[previousHouse])
		}
		bestEndingAt[currentHouse] = amount + bestCompatible
		bestOverall = max(bestOverall, bestEndingAt[currentHouse])
	}

	return formatHouseRobberOutput(bestOverall)
}

func solveByTopDownMemoization(input string) string {
	nums := parseHouseRobberInput(input)
	memo := make([]int, len(nums))
	for index := range memo {
		memo[index] = -1
	}

	var bestFrom func(int) int
	bestFrom = func(house int) int {
		if house >= len(nums) {
			return 0
		}
		if memo[house] != -1 {
			return memo[house]
		}

		robCurrent := nums[house] + bestFrom(house+2)
		skipCurrent := bestFrom(house + 1)
		memo[house] = max(robCurrent, skipCurrent)
		return memo[house]
	}

	return formatHouseRobberOutput(bestFrom(0))
}

func solveByRollingDynamicProgramming(input string) string {
	nums := parseHouseRobberInput(input)
	bestTwoHousesBack, bestOneHouseBack := 0, 0

	for _, amount := range nums {
		bestThroughCurrent := max(bestOneHouseBack, bestTwoHousesBack+amount)
		bestTwoHousesBack, bestOneHouseBack = bestOneHouseBack, bestThroughCurrent
	}

	return formatHouseRobberOutput(bestOneHouseBack)
}

/*
===============================================================================
AGENT REVIEW AND COMPLEXITY ANALYSIS

Current attempt:
- Score: 96/100. The attempt is correct and asymptotically optimal for every
  allowed input. The four-point deduction is for clarity rather than behavior:
  best_t1, best_t2, new_best, and visit do not follow Go naming conventions;
  max(best_t1, best_t2) is redundant because best_t1 never decreases; and the
  generated _ = nums and implementation comment remain after completion.
- What is correct: The length checks establish the first two prefix states.
  Before iteration visit, best_t1 is the maximum obtainable from houses
  [0,visit), while best_t2 is the maximum from houses [0,visit-1). Robbing
  nums[visit] must extend best_t2, while skipping it preserves best_t1. Taking
  their maximum establishes the invariant for the next iteration. Because all
  valid choices either rob or skip that house, the returned value is optimal.
  The attempt passes all configured official, adversarial, and boundary cases.
- First failing assumption or bottleneck: There is no correctness failure under
  1<=len(nums)<=100 and 0<=nums[i]<=400. The only bottleneck is readability:
  the state names obscure which prefix each value represents, making the update
  order harder to audit than descriptive rolling-state names would be.
- Relevant concept and invariant: This is bottom-up dynamic programming; see
  ../CONCEPTS.md#dynamic-programming. If best(i) is the
  optimum through house i, then best(i)=max(best(i-1), best(i-2)+nums[i]).
  Only the preceding two states are needed, so retaining the whole table is
  unnecessary.
- Concrete example: For nums=[2,7,9,3,1], initialization gives best_t1=7 and
  best_t2=2. At visit=2, new_best=max(7,9+2)=11, so the pair becomes (11,7).
  At visit=3 it becomes (max(11,3+7),11)=(11,11). At visit=4 it becomes
  (max(11,1+11),11)=(12,11), and the function returns 12.
- Comparison with solution 1: Both use dynamic programming, but the attempt
  makes one constant-time skip-or-rob transition per prefix. Solution 1 treats
  each house as the final robbed house and scans every compatible predecessor,
  making its direct formulation quadratic and retaining one value per house.
- Comparison with solution 2: Both implement the same skip-or-rob recurrence.
  The attempt evaluates prefixes iteratively and retains two states; solution 2
  evaluates suffixes recursively, memoizes every suffix, and uses linear stack
  and memo space. They differ in evaluation direction and storage, not in the
  underlying optimal-substructure argument.
- Comparison with solution 3: This is the same rolling dynamic-programming
  algorithm and invariant as the current attempt, not a distinct algorithm.
  The preferred version starts from the empty-prefix base state, uses
  descriptive Go-style names, and uses parallel assignment to make the state
  transition explicit.
- Time complexity: Theta(n) in the best, average, and worst cases because every
  house is visited once. JSON parsing also takes Theta(n); scalar formatting is
  Theta(1).
- Auxiliary-space complexity: Theta(1) for the algorithm itself. The boundary
  parser allocates the decoded nums slice, so the complete wrapper uses
  Theta(n) auxiliary space including input conversion.

Solution 1 - Direct compatible-predecessor dynamic programming:
- Correctness: bestEndingAt[i] is the greatest amount among valid robberies
  whose final robbed house is i. Any earlier robbed house must have index at
  most i-2, so adding nums[i] to the maximum bestEndingAt value over exactly
  those predecessors considers every valid plan ending at i and no invalid
  adjacent pair. bestOverall takes the maximum over every possible final house,
  plus the empty plan represented by zero, so it is globally optimal.
- Concrete example: For nums=[2,1,4,9], house 0 has no predecessor, so
  bestEndingAt[0]=2 and bestOverall=2. House 1 also has none, producing 1.
  House 2 can follow house 0, producing 4+2=6. House 3 scans houses 0 and 1;
  the better predecessor is house 0 with value 2, so it produces 9+2=11.
  The final bestEndingAt array is [2,1,6,11], and bestOverall is 11.
- Comparison with current attempt: The attempt folds each prefix into a
  constant-time skip-or-rob transition. This solution explicitly considers
  every possible compatible predecessor and therefore does quadratic work.
- Time complexity: Theta(n^2) in the best, average, and worst cases for growing
  n because the nested-loop bounds do not depend on the amounts. Parsing adds
  Theta(n), and scalar formatting is Theta(1).
- Auxiliary-space complexity: Theta(n) for bestEndingAt. The parsed nums slice
  is another Theta(n) allocation, leaving the complete wrapper at Theta(n).

Solution 2 - Improved top-down memoization:
- Correctness: bestFrom(i) partitions every valid suffix plan into two disjoint
  cases: rob house i and continue at i+2, or skip it and continue at i+1. The
  maximum is therefore optimal. Indices beyond the final house correctly yield
  zero, and memoization reuses an already proved suffix result without changing
  it. By induction from the empty suffix back to index 0, bestFrom(0) is the
  optimum for the whole street.
- Concrete example: For nums=[2,7,9,3,1], the solved suffix states from right
  to left are bestFrom(4)=max(1+0,0)=1, bestFrom(3)=max(3+0,1)=3,
  bestFrom(2)=max(9+1,3)=10, bestFrom(1)=max(7+3,10)=10, and
  bestFrom(0)=max(2+10,10)=12. Later requests for those suffixes return their
  memo entries instead of expanding another recursion tree.
- Comparison with current attempt: Both use the same recurrence. The attempt
  evaluates prefix states bottom-up with two integers; this solution evaluates
  suffix states on demand and retains both a memo slice and recursion frames.
- Time complexity: Theta(n) in the best, average, and worst cases: memoization
  computes each valid suffix once, and each transition is constant time. JSON
  parsing is also Theta(n); scalar formatting is Theta(1).
- Auxiliary-space complexity: Theta(n) for the memo slice and Theta(n) for the
  worst-case recursion stack. The parsed slice is Theta(n), so total wrapper
  auxiliary space remains Theta(n).

Solution 3 - Preferred rolling dynamic programming:
- Correctness: Before processing an amount, bestOneHouseBack is the optimum for
  the processed prefix and bestTwoHousesBack is the optimum excluding its last
  house. Skipping the current house preserves bestOneHouseBack; robbing it adds
  its amount to bestTwoHousesBack. Their maximum is exactly the optimum for the
  extended prefix, and parallel assignment advances both invariant states.
- Concrete example: For nums=[2,7,9,3,1], the rolling pair starts at (0,0).
  Processing 2 changes it to (0,2); processing 7 to (2,7); processing 9 to
  (7,11); processing 3 to (11,11); and processing 1 to (11,12). The final
  bestOneHouseBack is 12.
- Comparison with current attempt: Both are the same constant-space bottom-up
  algorithm. The current attempt special-cases lengths zero and one and seeds
  its pair from the first two houses; this version uses the empty-prefix state,
  so one loop uniformly handles every length, including the valid domain.
- Time complexity: Theta(n) in the best, average, and worst cases. The loop
  visits every house once, JSON parsing is Theta(n), and scalar formatting is
  Theta(1).
- Auxiliary-space complexity: Theta(1) for the algorithm. Including the nums
  slice allocated by JSON decoding, the complete wrapper uses Theta(n) space.

AGENT REVIEW AND COMPLEXITY ANALYSIS END
===============================================================================
*/

func main() {
	if err := testrunner.Run(os.Stdout, sampleCases, solutions); err != nil {
		fmt.Fprintln(os.Stderr, "workbench:", err)
		os.Exit(1)
	}
}
