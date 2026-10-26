//go:build ignore

// Difficulty: Medium
// Source: leetcode
// Problem ID: 34

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/duyserveacc/algos/dsa/internal/testrunner"
)

/*
======================================================================
Find First and Last Position of Element in Sorted Array

Source: LeetCode
Problem ID: 34
Source URL:
../../REFERENCE_LINKS.md#leetcode-source-0034

Given an array of integers nums sorted in non-decreasing order, find
the starting and ending position of a given target value.

If target is not found in the array, return [-1, -1].

You must write an algorithm with O(log n) runtime complexity.

Example 1:

Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]

Example 2:

Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]

Example 3:

Input: nums = [], target = 0
Output: [-1,-1]

Constraints:

- 0 <= nums.length <= 10^5

- -10^9 <= nums[i] <= 10^9

- nums is a non-decreasing array.

- -10^9 <= target <= 10^9
======================================================================
*/

// Problem-specific support.
//
// Keep shared judge types and parsing or formatting helpers here.
// Name adapters after the problem. Put setup used only by the current
// attempt at the top of solveCurrentAttempt below.
//
// This reviewed practice archive follows the standalone workbench
// structure.

func parseFindFirstAndLastPositionOfElementInSortedArrayInput(
	input string,
) ([]int, int) {
	var args []json.RawMessage
	if err := json.Unmarshal([]byte(input), &args); err != nil {
		panic(fmt.Errorf("parse find-first-and-last-position-of-"+
			"element-in-sorted-array input: %w", err))
	}
	if len(args) != 2 {
		panic(fmt.Sprintf("parse find-first-and-last-position-of-"+
			"element-in-sorted-array input: got %d "+
			"arguments, want 2", len(args)))
	}
	var nums []int
	argumentError0 := json.Unmarshal(args[0], &nums)
	if argumentError0 != nil {
		panic(fmt.Errorf("parse find-first-and-last-position-of-"+
			"element-in-sorted-array argument nums: %w", argumentError0))
	}
	var target int
	argumentError1 := json.Unmarshal(args[1], &target)
	if argumentError1 != nil {
		panic(fmt.Errorf("parse find-first-and-last-position-of-"+
			"element-in-sorted-array argument target: %w", argumentError1))
	}
	return nums, target
}

func formatFindFirstAndLastPositionOfElementInSortedArrayOutput(
	bounds []int,
) string {
	output, err := json.Marshal(bounds)
	if err != nil {
		panic(fmt.Errorf("format find-first-and-last-position-of-"+
			"element-in-sorted-array output: %w", err))
	}
	return string(output)
}

func maximumSearchRangeCase() testrunner.Case {
	nums := make([]int, 100000)
	for index := range nums {
		switch {
		case index < 25000:
			nums[index] = -1
		case index < 75000:
			nums[index] = 0
		default:
			nums[index] = 1
		}
	}
	input, err := json.Marshal([]any{nums, 0})
	if err != nil {
		panic(
			fmt.Errorf("build maximum find-first-and-last-position-"+
				"of-element-in-sorted-array case: %w", err),
		)
	}
	return testrunner.Case{
		Name: "Boundary: Maximum length contains a large " +
			"middle target run",
		Input: string(input),
		Want:  `[25000,74999]`,
	}
}

// sampleCases contains the examples that every registered solution
// must pass. Inputs encode the judge arguments as a positional JSON
// array.
var sampleCases = []testrunner.Case{
	{
		Name:  "Official example 1",
		Input: `[[5,7,7,8,8,10],8]`,
		Want:  `[3,4]`,
	},
	{
		Name:  "Official example 2",
		Input: `[[5,7,7,8,8,10],6]`,
		Want:  `[-1,-1]`,
	},
	{
		Name:  "Official example 3",
		Input: `[[],0]`,
		Want:  `[-1,-1]`,
	},
	{
		Name:  "Additional: Single matching value",
		Input: `[[1],1]`,
		Want:  `[0,0]`,
	},
	{
		Name:  "Additional: All values match",
		Input: `[[2,2,2],2]`,
		Want:  `[0,2]`,
	},
	{
		Name:  "Additional: Target between values",
		Input: `[[1,3],2]`,
		Want:  `[-1,-1]`,
	},
	{
		Name:  "Additional: Single value absent",
		Input: `[[1],2]`,
		Want:  `[-1,-1]`,
	},
	{
		Name:  "Additional: One match in the middle",
		Input: `[[1,2,3],2]`,
		Want:  `[1,1]`,
	},
	{
		Name:  "Boundary: Matching run starts at index zero",
		Input: `[[1,1,2,3],1]`,
		Want:  `[0,1]`,
	},
	{
		Name:  "Boundary: Matching run reaches final index",
		Input: `[[1,2,3,3],3]`,
		Want:  `[2,3]`,
	},
	{
		Name:  "Boundary: Target below the array range",
		Input: `[[1,2,3],0]`,
		Want:  `[-1,-1]`,
	},
	{
		Name:  "Boundary: Target above the array range",
		Input: `[[1,2,3],4]`,
		Want:  `[-1,-1]`,
	},
	{
		Name:  "Boundary: Minimum and maximum values",
		Input: `[[-1000000000,0,1000000000],1000000000]`,
		Want:  `[2,2]`,
	},
	maximumSearchRangeCase(),
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

var agentSolutions = []testrunner.Solution{
	{
		Name: "direct sorted linear scan",
		Run:  solveBySortedLinearScan,
	},
	{
		Name: "improved binary search then expand",
		Run:  solveByBinarySearchAndExpansion,
	},
	{
		Name: "preferred half-open binary bounds",
		Run:  solveByHalfOpenBinaryBounds,
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
	nums, target :=
		parseFindFirstAndLastPositionOfElementInSortedArrayInput(
			input,
		)
	_ = nums
	_ = target

	left, right := 0, len(nums)-1
	first, last := -1, -1
	ret := make([]int, 2)

	if len(nums) == 0 {
		ret[0] = first
		ret[1] = last
		return formatFindFirstAndLastPositionOfElementInSortedArrayOutput(
			ret,
		)
	}

	for left < right {
		middle := left + (right-left)/2
		if target == nums[middle] {
			right = middle
		} else if target < nums[middle] {
			right = middle
		} else if target > nums[middle] {
			left = middle + 1
		}
	}
	if nums[left] == target {
		first = left
	}
	left, right = 0, len(nums)-1

	for left < right {
		middle := left + (right-left+1)/2
		if target >= nums[middle] {
			left = middle
		} else {
			right = middle - 1
		}
	}
	if nums[left] == target {
		last = left
	}

	ret[0] = first
	ret[1] = last

	// Implement your algorithm with the typed values above.
	return formatFindFirstAndLastPositionOfElementInSortedArrayOutput(
		ret,
	)
}

// ===================================================================
// AGENT SOLUTIONS
//
// These functions remain separate from the preserved current attempt
// so each approach can be compared against the same cases.
// ===================================================================

func solveBySortedLinearScan(input string) string {
	nums, target :=
		parseFindFirstAndLastPositionOfElementInSortedArrayInput(
			input,
		)
	first, last := -1, -1

	for index, value := range nums {
		if value < target {
			continue
		}
		if value > target {
			break
		}
		if first == -1 {
			first = index
		}
		last = index
	}

	return formatFindFirstAndLastPositionOfElementInSortedArrayOutput(
		[]int{first, last},
	)
}

func solveByBinarySearchAndExpansion(input string) string {
	nums, target :=
		parseFindFirstAndLastPositionOfElementInSortedArrayInput(
			input,
		)
	left, right := 0, len(nums)-1
	found := -1

	for left <= right {
		middle := left + (right-left)/2
		switch {
		case nums[middle] < target:
			left = middle + 1
		case nums[middle] > target:
			right = middle - 1
		default:
			found = middle
		}
		if found != -1 {
			break
		}
	}

	if found == -1 {
		return formatFindFirstAndLastPositionOfElementInSortedArrayOutput(
			[]int{-1, -1},
		)
	}

	first, last := found, found
	for first > 0 && nums[first-1] == target {
		first--
	}
	for last+1 < len(nums) && nums[last+1] == target {
		last++
	}

	return formatFindFirstAndLastPositionOfElementInSortedArrayOutput(
		[]int{first, last},
	)
}

func solveByHalfOpenBinaryBounds(input string) string {
	nums, target :=
		parseFindFirstAndLastPositionOfElementInSortedArrayInput(
			input,
		)

	firstIndexAtLeast := func() int {
		left, right := 0, len(nums)
		for left < right {
			middle := left + (right-left)/2
			if nums[middle] < target {
				left = middle + 1
			} else {
				right = middle
			}
		}
		return left
	}

	firstIndexGreaterThan := func() int {
		left, right := 0, len(nums)
		for left < right {
			middle := left + (right-left)/2
			if nums[middle] <= target {
				left = middle + 1
			} else {
				right = middle
			}
		}
		return left
	}

	first := firstIndexAtLeast()
	if first == len(nums) || nums[first] != target {
		return formatFindFirstAndLastPositionOfElementInSortedArrayOutput(
			[]int{-1, -1},
		)
	}

	return formatFindFirstAndLastPositionOfElementInSortedArrayOutput(
		[]int{first, firstIndexGreaterThan() - 1},
	)
}

/*
======================================================================
AGENT REVIEW AND COMPLEXITY ANALYSIS

Current attempt:
- What is correct: The corrected attempt passes every supplied case.
  Its first
closed-interval search converges on the earliest value at least
target, and
its second converges on the latest value at most target. The equality
checks
after convergence distinguish an actual match from the closest
boundary.
- First failing assumption or bottleneck: No correctness assumption
  remains
broken under the sorted-input contract. The earlier second search used
a
floor midpoint with left = middle, which could retain [left,right]
forever
when right = left+1. The corrected ceiling midpoint guarantees
progress.
The remaining duplication is that both searches repeat nearly the same
binary-search structure and both search the full range even when
target is
  absent.
- Relevant concept and invariant: See ../CONCEPTS.md#binary-search.
  During the first search, the
first occurrence remains in [left,right] whenever it exists; equality
keeps
middle by assigning right = middle. During the second, the last
occurrence
remains in [left,right]; equality keeps middle by assigning left =
middle.
Floor and ceiling midpoints make the retained interval shrink
respectively.
- Concrete example: For nums=[5,7,7,8,8,10], target=8, the first
  search moves
  [0,5] --mid 2/value 7--> [3,5] --mid 4/value 8--> [3,4]
--mid 3/value 8--> [3,3], so first=3. The second uses ceiling
midpoints:
  [0,5] --mid 3/value 8--> [3,5] --mid 4/value 8--> [4,5]
  --mid 5/value 10--> [4,4], so last=4.
- Comparison with solution 1: Both ultimately identify the first and
  last
equal values. The scan visits the target run directly and can be
simpler,
while the attempt exploits sorted order for logarithmic worst-case
time.
- Comparison with solution 2: That solution uses binary search only to
  find
one match, then expands through adjacent duplicates. The attempt
performs
  logarithmic boundary searches and does not visit every duplicate.
- Comparison with solution 3: Both use two binary searches and the
  same
lower-bound/upper-bound idea, so they are the same core algorithm. The
attempt uses closed intervals, opposite midpoint biases, and final
equality
checks; the preferred version uses uniform half-open intervals and
derives
  last as upperBound-1.
- Time complexity: Empty input takes Theta(1); otherwise both searches
  take
Theta(log n) in the best, average, and worst cases. JSON parsing makes
the
  complete local wrapper Theta(n) end to end.
- Auxiliary-space complexity: O(1) for the searches and two-result
  slice. The
harness additionally allocates Theta(n) for decoded input and O(1)
bounded
  output under the stated integer constraints.

Solution 1 - Direct sorted linear scan:
- Correctness: Sorted order places all target values in one contiguous
  run.
Values below target can be skipped; the first equality records first,
every
equality advances last, and the first greater value proves the run is
over.
  If equality never occurs, both indices correctly remain -1.
- Concrete example: For nums=[1,2,2,2,3], target=2, index 0/value 1 is
  skipped.
Index 1 records first=1 and last=1; indices 2 and 3 update last to 2
and 3.
  Index 4/value 3 ends the scan, producing [1,3].
- Comparison with current attempt: Both rely on the contiguous
  duplicate-run
property. This version expresses it with a direct traversal; the
attempt
  discards ranges and retains logarithmic worst-case search time.
- Time complexity: Theta(1) in the best case when the first value is
  greater
than target, and Theta(n) on average and in the worst case. Parsing
already
  makes the complete local wrapper Theta(n).
- Auxiliary-space complexity: O(1) algorithmic space, including the
  fixed
two-element result. The harness additionally allocates Theta(n) for
parsing
  and O(1) bounded output.

Solution 2 - Improved binary search then expand:
- Correctness: Ordinary binary search either proves target absent or
  returns
one occurrence. Because equal values are contiguous, moving left while
the
predecessor equals target must stop at the first occurrence, and
moving
  right while the successor equals target must stop at the last.
- Concrete example: For nums=[1,2,2,2,3], target=2, binary search
  first checks
middle=2 and finds target. Expansion moves first from 2 to 1, then
stops at
value 1; it moves last from 2 to 3, then stops at value 3. The result
is
  [1,3].
- Comparison with current attempt: Both begin from binary-search
  reasoning,
but this solution trades simpler boundary logic for a possible linear
walk
across duplicates. The attempt locates both ends without traversing
the run.
- Time complexity: Best case is Theta(1) when the first midpoint is a
  unique
target. More generally it is Theta(log n+k), where k is the matched
run
length; average behavior depends on k, and the worst case is Theta(n)
when
all values equal target. Parsing makes the full local wrapper
Theta(n).
- Auxiliary-space complexity: O(1) for search and expansion. The
  harness
  additionally allocates Theta(n) for input and O(1) bounded output.

Solution 3 - Preferred half-open binary bounds:
- Correctness: firstIndexAtLeast maintains [left,right) around the
  first value
greater than or equal to target; firstIndexGreaterThan maintains it
around
the first value strictly greater than target. Each branch removes
middle or
an entire side while preserving its boundary. If the lower bound is
not an
  equal value, target is absent. Otherwise the matching run is exactly
  [lowerBound,upperBound), so [lowerBound,upperBound-1] is the answer.
- Concrete example: For nums=[5,7,7,8,8,10], target=8, lower-bound
  states are
  [0,6) --mid 3/value 8--> [0,3) --mid 1/value 7--> [2,3)
  --mid 2/value 7--> [3,3), giving 3. Upper-bound states are [0,6)
  --mid 3/value 8--> [4,6) --mid 5/value 10--> [4,5)
  --mid 4/value 8--> [5,5), giving 5; the last match is 5-1=4.
- Comparison with current attempt: This is the same two-boundary
  algorithm and
O(1)-space invariant as the attempt, not a distinct algorithmic
method. Its
half-open interval permits one floor-midpoint rule for both bounds,
naturally
represents len(nums), and avoids separate left/right midpoint biases.
- Time complexity: Empty input is Theta(1); otherwise the two bounds
  are
Theta(log n) in the best, average, and worst cases. JSON parsing makes
the
  complete local wrapper Theta(n).
- Auxiliary-space complexity: O(1) for the bounds plus the returned
  two-value
slice. Parsing allocates Theta(n), while formatted output is O(1)
under the
  fixed numeric constraints.

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
