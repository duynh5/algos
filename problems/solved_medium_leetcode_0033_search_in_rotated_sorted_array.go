//go:build ignore

// Difficulty: Medium
// Source: leetcode
// Problem ID: 33

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/duyserveacc/algos/internal/testrunner"
)

/*
===============================================================================
Search in Rotated Sorted Array

Source: LeetCode
Problem ID: 33
Source URL: https://leetcode.com/problems/search-in-rotated-sorted-array/description/

There is an integer array nums sorted in ascending order (with distinct values).

Prior to being passed to your function, nums is possibly left rotated at an unknown index k (1 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For example, [0,1,2,4,5,6,7] might be left rotated by 3 indices and become [4,5,6,7,0,1,2].

Given the array nums after the possible rotation and an integer target, return the index of target if it is in nums, or -1 if it is not in nums.

You must write an algorithm with O(log n) runtime complexity.

Example 1:

Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4

Example 2:

Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1

Example 3:

Input: nums = [1], target = 0
Output: -1

Constraints:

- 1 <= nums.length <= 5000

- -10^4 <= nums[i] <= 10^4

- All values of nums are unique.

- nums is an ascending array that is possibly rotated.

- -10^4 <= target <= 10^4
===============================================================================
*/

// Problem-specific support.
//
// Keep shared judge types and parsing or formatting helpers here. Name adapters
// after the problem. Put setup used only by the current attempt at the top of
// solveCurrentAttempt below.
//
// This reviewed practice archive follows the standalone workbench structure.

func parseSearchInRotatedSortedArrayInput(input string) ([]int, int) {
	var args []json.RawMessage
	if err := json.Unmarshal([]byte(input), &args); err != nil {
		panic(fmt.Errorf("parse search-in-rotated-sorted-array input: %w", err))
	}
	if len(args) != 2 {
		panic(fmt.Sprintf("parse search-in-rotated-sorted-array input: got %d arguments, want 2", len(args)))
	}
	var nums []int
	argumentError0 := json.Unmarshal(args[0], &nums)
	if argumentError0 != nil {
		panic(fmt.Errorf("parse search-in-rotated-sorted-array argument nums: %w", argumentError0))
	}
	var target int
	argumentError1 := json.Unmarshal(args[1], &target)
	if argumentError1 != nil {
		panic(fmt.Errorf("parse search-in-rotated-sorted-array argument target: %w", argumentError1))
	}
	return nums, target
}

func maximumSearchInRotatedSortedArrayCase() testrunner.Case {
	const length = 5000
	const rotation = 1234
	nums := make([]int, length)
	for index := range nums {
		nums[index] = -2500 + (index+rotation)%length
	}
	input, err := json.Marshal([]any{nums, -2500})
	if err != nil {
		panic(fmt.Errorf("build maximum search-in-rotated-sorted-array case: %w", err))
	}
	return testrunner.Case{
		Name:  "Boundary: Maximum length finds a target across the rotation",
		Input: string(input),
		Want:  `3766`,
	}
}

func formatSearchInRotatedSortedArrayOutput(result int) string {
	return testrunner.FormatJSONOutput(result)
}

// sampleCases contains the examples that every registered solution must pass.
// Inputs encode the judge arguments as a positional JSON array.
var sampleCases = []testrunner.Case{
	{
		Name:  "Official example 1",
		Input: `[[4,5,6,7,0,1,2],0]`,
		Want:  `4`,
	},
	{
		Name:  "Official example 2",
		Input: `[[4,5,6,7,0,1,2],3]`,
		Want:  `-1`,
	},
	{
		Name:  "Official example 3",
		Input: `[[1],0]`,
		Want:  `-1`,
	},
	{
		Name:  "Additional: Single value found",
		Input: `[[1],1]`,
		Want:  `0`,
	},
	{
		Name:  "Additional: Unrotated array",
		Input: `[[1,2,3,4],3]`,
		Want:  `2`,
	},
	{
		Name:  "Additional: Target is first value before the pivot",
		Input: `[[4,5,6,7,0,1,2],4]`,
		Want:  `0`,
	},
	{
		Name:  "Additional: Target is final value after the pivot",
		Input: `[[4,5,6,7,0,1,2],2]`,
		Want:  `6`,
	},
	{
		Name:  "Boundary: Two-element rotation finds the second value",
		Input: `[[3,1],1]`,
		Want:  `1`,
	},
	{
		Name:  "Additional: Rotation by one places minimum last",
		Input: `[[2,3,4,5,1],1]`,
		Want:  `4`,
	},
	{
		Name:  "Additional: Near-full rotation places maximum first",
		Input: `[[5,1,2,3,4],5]`,
		Want:  `0`,
	},
	{
		Name:  "Adversarial: Target equals right boundary of sorted half",
		Input: `[[5,1,2,3,4],4]`,
		Want:  `4`,
	},
	{
		Name:  "Adversarial: Missing target leaves an interior insertion index",
		Input: `[[5,1,2,3,4],0]`,
		Want:  `-1`,
	},
	{
		Name:  "Boundary: Signed value limits",
		Input: `[[0,10000,-10000],-10000]`,
		Want:  `2`,
	},
	maximumSearchInRotatedSortedArrayCase(),
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
		Name: "improved pivot then binary search",
		Run:  solveByPivotThenBinarySearch,
	},
	{
		Name: "preferred one-pass binary search",
		Run:  solveByOnePassBinarySearch,
	},
}

var solutions = append([]testrunner.Solution{currentAttemptSolution}, agentSolutions...)

// =============================================================================
// CURRENT ATTEMPT
//
// Implement your solution here. Keep this body intact when requesting review.
// =============================================================================

func solveCurrentAttempt(input string) string {
	nums, target := parseSearchInRotatedSortedArrayInput(input)
	left, right := 0, len(nums)-1

	for left <= right {
		middle := left + (right-left)/2
		if nums[middle] == target {
			return formatSearchInRotatedSortedArrayOutput(middle)
		}
		if nums[left] <= nums[middle] {
			if target >= nums[left] && target < nums[middle] {
				right = middle - 1
			} else {
				left = middle + 1
			}
		} else {
			if target > nums[middle] && target <= nums[right] {
				left = middle + 1
			} else {
				right = middle - 1
			}
		}
	}

	// Implement your algorithm with the typed values above.
	return formatSearchInRotatedSortedArrayOutput(-1)
}

// =============================================================================
// AGENT SOLUTIONS
//
// These functions remain separate from the preserved current attempt so each
// approach can be compared against the same cases.
// =============================================================================

func solveByLinearScan(input string) string {
	nums, target := parseSearchInRotatedSortedArrayInput(input)

	for index, value := range nums {
		if value == target {
			return formatSearchInRotatedSortedArrayOutput(index)
		}
	}

	return formatSearchInRotatedSortedArrayOutput(-1)
}

func solveByPivotThenBinarySearch(input string) string {
	nums, target := parseSearchInRotatedSortedArrayInput(input)
	left, right := 0, len(nums)-1

	for left < right {
		middle := left + (right-left)/2
		if nums[middle] > nums[right] {
			left = middle + 1
		} else {
			right = middle
		}
	}
	pivot := left

	left, right = 0, len(nums)-1
	if target >= nums[pivot] && target <= nums[len(nums)-1] {
		left = pivot
	} else {
		right = pivot - 1
	}

	for left <= right {
		middle := left + (right-left)/2
		switch {
		case nums[middle] < target:
			left = middle + 1
		case nums[middle] > target:
			right = middle - 1
		default:
			return formatSearchInRotatedSortedArrayOutput(middle)
		}
	}

	return formatSearchInRotatedSortedArrayOutput(-1)
}

func solveByOnePassBinarySearch(input string) string {
	nums, target := parseSearchInRotatedSortedArrayInput(input)
	left, right := 0, len(nums)-1

	for left <= right {
		middle := left + (right-left)/2
		if nums[middle] == target {
			return formatSearchInRotatedSortedArrayOutput(middle)
		}

		if nums[left] <= nums[middle] {
			if nums[left] <= target && target < nums[middle] {
				right = middle - 1
			} else {
				left = middle + 1
			}
		} else if nums[middle] < target && target <= nums[right] {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return formatSearchInRotatedSortedArrayOutput(-1)
}

/*
===============================================================================
AGENT REVIEW AND COMPLEXITY ANALYSIS

Current attempt:
- What is correct: The attempt uses a closed [left, right] interval, checks the
  midpoint before discarding it, guarantees progress with middle+1 and
  middle-1, includes both outer endpoints in sorted-half membership tests, and
  returns -1 only after the candidate interval is empty. It passes all 14
  configured cases.
- First failing assumption or bottleneck: There is no remaining correctness
  failure in the stated non-empty, distinct-value input domain. The main
  difficulty is the asymmetric boundary logic: [left,middle) uses an inclusive
  left endpoint, while (middle,right] uses an inclusive right endpoint.
- Relevant concept and invariant: As described in
  concepts/README.md#binary-search, at least one half of a rotated
  distinct-value interval is sorted. After testing the midpoint, endpoint
  comparisons must determine whether target is inside that sorted half; only
  then is it safe to discard the other half. Exhaustion means absent, not an
  insertion result.
- Concrete example: For nums=[5,1,2,3,4], target=4, [0,4] checks index 2 with
  value 2. The right half is sorted and 2 < 4 <= 4, so left becomes 3. Interval
  [3,4] checks index 3 with value 3; its left half is sorted but excludes 4, so
  left becomes 4. Index 4 then matches and is returned.
- Comparison with solution 1: Both compare values directly with target. The
  scan never reasons about rotated order, making it simpler but giving up the
  attempt's logarithmic search time.
- Comparison with solution 2: Both use closed binary-search intervals. The
  two-phase solution first restores two globally sorted ranges by locating the
  rotation pivot, then searches one range with ordinary binary search; the
  attempt instead chooses a sorted half during every iteration.
- Comparison with solution 3: Both use the same one-pass modified binary-search
  algorithm, comparisons, and sorted-half invariant. Their material difference
  is only control-flow shape: the preferred version combines the right-sorted
  membership test into an else-if branch. It is not a distinct algorithm from
  the current attempt.
- Time complexity: The search is Theta(1) when the first midpoint matches and
  Theta(log n) on average and in the worst case. JSON parsing makes the complete
  local wrapper Theta(n).
- Auxiliary-space complexity: O(1) for the search. The harness allocates O(n)
  for parsed input and O(log n) bytes for the integer result.

Solution 1 - Direct linear scan:
- Correctness: The scan visits every index exactly once and returns only when
  its value equals target. If no equality is seen after the final index, target
  is absent and -1 is correct. Rotation does not affect this argument.
- Concrete example: For nums=[4,5,6,7,0,1,2], target=0, indices 0 through 3
  contain 4, 5, 6, and 7, so the scan continues. Index 4 contains 0, so it
  returns 4 without inspecting indices 5 and 6.
- Comparison with current attempt: It avoids all sorted-half and endpoint
  reasoning but has linear rather than logarithmic worst-case search time.
- Time complexity: Theta(1) when index 0 matches, Theta(n) on average and in
  the worst case. Parsing makes the complete local wrapper Theta(n).
- Auxiliary-space complexity: O(1) for the scan. Harness parsing allocates
  O(n), and formatting the integer result allocates O(log n) bytes.

Solution 2 - Improved pivot then binary search:
- Correctness: Comparing nums[middle] with nums[right] preserves the index of
  the minimum in [left, right]: a larger midpoint lies before the pivot, while
  a midpoint no larger than the right endpoint lies at or after it. Convergence
  therefore finds the pivot. The pivot divides the array into two sorted
  ranges; endpoint checks select the only range that can contain target, and
  ordinary binary search is correct within that range.
- Concrete example: For nums=[4,5,6,7,0,1,2], target=0, pivot search checks
  middle 3 (7>2) and moves left to 4; middle 5 (1<=2) moves right to 5; middle
  4 (0<=1) moves right to 4. Pivot 4 starts range [0,1,2]. Target 0 selects
  indices [4,6]; binary search checks index 5, moves right to 4, then matches
  index 4.
- Comparison with current attempt: Both are logarithmic closed-interval
  searches. This version makes the rotation boundary explicit in one phase;
  the attempt repeatedly identifies a sorted half without locating the pivot.
- Time complexity: Theta(log n) for n>1 in the best, average, and worst cases:
  pivot search is followed by at most one binary search. Parsing makes the
  complete local wrapper Theta(n).
- Auxiliary-space complexity: O(1) for both searches. The harness additionally
  allocates O(n) for parsing and O(log n) for output.

Solution 3 - Preferred one-pass binary search:
- Correctness: The active interval contains every possible target index. If
  the left half is sorted, target belongs there exactly when
  nums[left] <= target < nums[middle]; otherwise it can only be in the other
  half. When the left half is not sorted, the right half is sorted and target
  belongs there exactly when nums[middle] < target <= nums[right]. Equality is
  checked first, and each branch excludes middle, so the invariant holds while
  the interval strictly shrinks. An empty interval proves absence.
- Concrete example: For nums=[4,5,6,7,0,1,2], target=0, [0,6] checks index 3
  containing 7. The left half [4,5,6,7] is sorted but excludes 0, so left=4.
  Interval [4,6] checks index 5 containing 1; its sorted left half [0,1]
  contains target, so right=4. Index 4 then matches and is returned.
- Comparison with current attempt: This uses the attempt's same core algorithm
  and O(1)-space invariant. It only flattens the attempt's right-sorted branch
  into an else-if; the comparisons and interval updates are equivalent.
- Time complexity: Theta(1) when the initial midpoint matches, and Theta(log n)
  on average and in the worst case. Parsing makes the complete local wrapper
  Theta(n).
- Auxiliary-space complexity: O(1) for the search. The harness additionally
  allocates O(n) for parsed input and O(log n) for output.

AGENT REVIEW AND COMPLEXITY ANALYSIS END
===============================================================================
*/

func main() {
	if err := testrunner.Run(os.Stdout, sampleCases, solutions); err != nil {
		fmt.Fprintln(os.Stderr, "workbench:", err)
		os.Exit(1)
	}
}
