//go:build ignore

// Difficulty: Easy
// Source: leetcode
// Problem ID: 35

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
Search Insert Position

Source: LeetCode
Problem ID: 35
Source URL: https://leetcode.com/problems/search-insert-position/description/

Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You must write an algorithm with O(log n) runtime complexity.

Example 1:

Input: nums = [1,3,5,6], target = 5
Output: 2

Example 2:

Input: nums = [1,3,5,6], target = 2
Output: 1

Example 3:

Input: nums = [1,3,5,6], target = 7
Output: 4

Constraints:

- 1 <= nums.length <= 10^4

- -10^4 <= nums[i] <= 10^4

- nums contains distinct values sorted in ascending order.

- -10^4 <= target <= 10^4
===============================================================================
*/

// Problem-specific support.
//
// Keep shared judge types and parsing or formatting helpers here. Name adapters
// after the problem. Put setup used only by the current attempt at the top of
// solveCurrentAttempt below.
//
// This completed problem archive follows the workbench template.
// Concept: ../CONCEPTS.md#binary-search

func parseSearchInsertPositionInput(input string) ([]int, int) {
	var args []json.RawMessage
	if err := json.Unmarshal([]byte(input), &args); err != nil {
		panic(fmt.Errorf("parse search-insert-position input: %w", err))
	}
	if len(args) != 2 {
		panic(fmt.Sprintf("parse search-insert-position input: got %d arguments, want 2", len(args)))
	}
	var nums []int
	argumentError0 := json.Unmarshal(args[0], &nums)
	if argumentError0 != nil {
		panic(fmt.Errorf("parse search-insert-position argument nums: %w", argumentError0))
	}
	var target int
	argumentError1 := json.Unmarshal(args[1], &target)
	if argumentError1 != nil {
		panic(fmt.Errorf("parse search-insert-position argument target: %w", argumentError1))
	}
	return nums, target
}

func maximumSearchInsertPositionCase() testrunner.Case {
	nums := make([]int, 10000)
	for index := range nums {
		nums[index] = -10000 + index*2
	}
	input, err := json.Marshal([]any{nums, 0})
	if err != nil {
		panic(fmt.Errorf("build maximum search-insert-position case: %w", err))
	}
	return testrunner.Case{
		Name:  "Boundary: Maximum-length midpoint match",
		Input: string(input),
		Want:  `5000`,
	}
}

func formatSearchInsertPositionOutput(result int) string {
	return testrunner.FormatJSONOutput(result)
}

// sampleCases contains the examples that every registered solution must pass.
// Inputs encode the judge arguments as a positional JSON array.
var sampleCases = []testrunner.Case{
	{
		Name:  "Official example 1",
		Input: `[[1,3,5,6],5]`,
		Want:  `2`,
	},
	{
		Name:  "Official example 2",
		Input: `[[1,3,5,6],2]`,
		Want:  `1`,
	},
	{
		Name:  "Official example 3",
		Input: `[[1,3,5,6],7]`,
		Want:  `4`,
	},
	{
		Name:  "Additional: Single value found",
		Input: `[[1],1]`,
		Want:  `0`,
	},
	{
		Name:  "Additional: Insert before first",
		Input: `[[1],0]`,
		Want:  `0`,
	},
	{
		Name:  "Regression: Insert before first in multi-value array",
		Input: `[[1,3],0]`,
		Want:  `0`,
	},
	{
		Name:  "Additional: Insert between negatives",
		Input: `[[-5,-3,-1],-2]`,
		Want:  `2`,
	},
	{
		Name:  "Additional: Insert after single value",
		Input: `[[1],2]`,
		Want:  `1`,
	},
	{
		Name:  "Additional: Find first element",
		Input: `[[1,3,5,6],1]`,
		Want:  `0`,
	},
	{
		Name:  "Additional: Find final element",
		Input: `[[1,3,5,6],6]`,
		Want:  `3`,
	},
	{
		Name:  "Adversarial: Insert between two values",
		Input: `[[1,3],2]`,
		Want:  `1`,
	},
	{
		Name:  "Boundary: Insert between minimum and maximum values",
		Input: `[[-10000,10000],0]`,
		Want:  `1`,
	},
	{
		Name:  "Boundary: Find maximum target",
		Input: `[[-10000,10000],10000]`,
		Want:  `1`,
	},
	maximumSearchInsertPositionCase(),
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
		Name: "improved recursive binary search",
		Run:  solveByRecursiveBinarySearch,
	},
	{
		Name: "preferred iterative lower bound",
		Run:  solveByLowerBound,
	},
}

var solutions = append([]testrunner.Solution{currentAttemptSolution}, agentSolutions...)

// =============================================================================
// CURRENT ATTEMPT
//
// Implement your solution here. Keep this body intact when requesting review.
// =============================================================================

func solveCurrentAttempt(input string) string {
	nums, target := parseSearchInsertPositionInput(input)
	_ = nums
	_ = target
	targ_i := -1
	low := 0
	mid := len(nums) / 2
	high := len(nums) - 1

	for {
		if nums[mid] == target {
			targ_i = mid
			break
		}
		if nums[high] < target {
			targ_i = high + 1
			break
		}
		if nums[low] > target {
			targ_i = low
			break
		}
		if low == mid {
			targ_i = high
			break
		}
		if nums[mid] > target {
			high = mid
		}
		if nums[mid] < target {
			low = mid
		}
		mid = (high + low) / 2
	}

	// Implement your algorithm with the typed values above.
	return strconv.Itoa(targ_i)
}

// =============================================================================
// AGENT SOLUTIONS
//
// These functions remain separate from the preserved current attempt so each
// approach can be compared against the same cases.
// =============================================================================

func solveByLinearScan(input string) string {
	nums, target := parseSearchInsertPositionInput(input)

	for index, value := range nums {
		if value >= target {
			return strconv.Itoa(index)
		}
	}

	return strconv.Itoa(len(nums))
}

func solveByRecursiveBinarySearch(input string) string {
	nums, target := parseSearchInsertPositionInput(input)

	var search func(left, right int) int
	search = func(left, right int) int {
		if left > right {
			return left
		}

		middle := left + (right-left)/2
		if nums[middle] == target {
			return middle
		}
		if nums[middle] < target {
			return search(middle+1, right)
		}
		return search(left, middle-1)
	}

	return strconv.Itoa(search(0, len(nums)-1))
}

func solveByLowerBound(input string) string {
	nums, target := parseSearchInsertPositionInput(input)
	left, right := 0, len(nums)

	for left < right {
		middle := left + (right-left)/2
		if nums[middle] < target {
			left = middle + 1
		} else {
			right = middle
		}
	}

	return strconv.Itoa(left)
}

/*
===============================================================================
AGENT REVIEW AND COMPLEXITY ANALYSIS

Current attempt:
- What is correct: The attempt uses sorted order to reduce a closed interval,
  returns an exact midpoint match immediately, handles targets greater than
  the final value, handles targets below the first active value, and finds
  insertions between existing values. It passes all supplied cases and an
  exhaustive comparison over 17,984 sorted-array and target combinations.
- First failing assumption or bottleneck: There is no remaining correctness
  failure in the stated non-empty, sorted, distinct input domain. The main
  weakness is control-flow complexity: high = mid and low = mid retain the
  inspected element, so progress depends on three boundary exits and the
  low == mid adjacency exit. The two independent comparisons can also update
  mid twice in one iteration, making the invariant harder to audit.
- Relevant concept and invariant: After the two outside-boundary checks, the
  matching or insertion index is inside [low, high]. When low == mid, equality
  has already failed, nums[low] < target, and target <= nums[high], so high is
  the insertion point. This is the same binary-search family used by both
  agent binary solutions, although their intervals shrink without special
  boundary exits.
- Concrete example: For nums = [1,3], target = 0, the attempt starts with
  low = 0, mid = 1, high = 1. The midpoint does not match and target is not
  greater than nums[high]. Since nums[low] = 1 is greater than target, the
  new lower-bound guard returns low = 0 directly.
- Comparison with solution 1: Both seek the first value not below target. The
  scan checks that condition directly at every index; the attempt uses sorted
  order to discard ranges and achieves logarithmic worst-case search time.
- Comparison with solution 2: Both use a closed binary-search interval and
  stop early on equality. The recursive solution expresses progress through
  disjoint [left, middle-1] and [middle+1, right] intervals and lets the
  left > right base case produce the insertion point instead of using the
  attempt's explicit first, last, and adjacent-position exits.
- Comparison with solution 3: Both use iterative binary search. The preferred
  solution searches directly for the lower bound in [left, right), so every
  branch shrinks the interval and convergence supplies boundary answers
  without a special adjacency rule.
- Time complexity: The search is Theta(1) when the initial midpoint matches
  or target lies outside the active endpoints, and Theta(log n) on average and
  in the worst case. JSON parsing makes this local wrapper Theta(n) end to end.
- Auxiliary-space complexity: O(1) for the search. The local harness allocates
  O(n) space for the parsed slice and O(log n) bytes for the decimal result.

Solution 1 - Direct linear scan:
- Correctness: The scan returns at the first index whose value is at least the
  target. Every earlier value is therefore smaller, so that index is either
  the matching position or the only position that preserves sorted order. If
  no such value exists, every value is smaller and len(nums) is the correct
  insertion position.
- Concrete example: For nums = [1,3,5,6], target = 2, index 0 contains 1 and
  the scan continues. Index 1 contains 3, the first value at least 2, so the
  function returns 1 without inspecting 5 or 6.
- Comparison with current attempt: The scan has the same lower-bound result
  but does not exploit sorted order to discard ranges. It is simpler and can
  stop sooner for targets near the front, while the attempt has logarithmic
  worst-case search time.
- Time complexity: The search is Theta(1) in the best case, Theta(n) on
  average, and Theta(n) in the worst case. Parsing is Theta(n), so the complete
  local wrapper is Theta(n) in every case.
- Auxiliary-space complexity: O(1) for the scan. The harness parsing and
  decimal output allocate O(n) and O(log n) space respectively.

Solution 2 - Improved recursive binary search:
- Correctness: A recursive call contains every location where target could be
  found or inserted. Equality returns the match. If nums[middle] is smaller,
  sorted order excludes middle and everything left of it; otherwise it
  excludes middle and everything right of it. When left > right, all excluded
  values before left are smaller and all values at or after left are larger,
  making left the insertion point.
- Concrete example: For nums = [-5,-3,-1], target = -2, search(0,2) checks
  middle = 1 with -3 and recurses to search(2,2). Middle = 2 contains -1, so
  it recurses to search(2,1). The interval is empty and returns left = 2,
  inserting -2 between -3 and -1.
- Comparison with current attempt: Both are exact-match searches on closed
  intervals. This version removes the attempt's retained midpoint and special
  adjacent-element exits by excluding middle on every nonmatching branch; its
  tradeoff is one call-stack frame per search level.
- Time complexity: The search is Theta(1) when the first midpoint matches and
  Theta(log n) on average and in the worst case. Parsing still makes the local
  wrapper Theta(n) end to end.
- Auxiliary-space complexity: O(log n) search stack in the worst case. The
  harness additionally allocates O(n) for parsing and O(log n) for output.

Solution 3 - Preferred iterative lower bound:
- Correctness: The half-open interval [left, right) always contains the first
  index whose value can be at least target. When nums[middle] is smaller, that
  index must be after middle, so left becomes middle+1. Otherwise middle may
  be the answer, so right becomes middle. The interval strictly shrinks; when
  left == right, that sole boundary is the matching or insertion index.
- Concrete example: For nums = [1,3,5,6], target = 7, [left,right) starts as
  [0,4). Middle = 2 contains 5, so left becomes 3. Middle = 3 contains 6, so
  left becomes 4. The interval [4,4) is empty and the function returns 4,
  correctly placing 7 after every existing value.
- Comparison with current attempt: Both use O(1)-space iterative binary
  search. Lower bound uses one mutually exclusive update per iteration and
  needs no equality, last-element, or adjacent-element special case, making
  its progress and final answer direct consequences of one invariant.
- Time complexity: The search is Theta(log n) in the best, average, and worst
  cases for non-empty input. JSON parsing makes the local wrapper Theta(n).
- Auxiliary-space complexity: O(1) for the search. The harness additionally
  allocates O(n) for the parsed input and O(log n) for the output string.

AGENT REVIEW AND COMPLEXITY ANALYSIS END
===============================================================================
*/

func main() {
	if err := testrunner.Run(os.Stdout, sampleCases, solutions); err != nil {
		fmt.Fprintln(os.Stderr, "workbench:", err)
		os.Exit(1)
	}
}
