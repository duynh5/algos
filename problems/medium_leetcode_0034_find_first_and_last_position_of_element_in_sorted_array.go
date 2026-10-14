//go:build ignore

// Difficulty: Medium
// Source: leetcode
// Problem ID: 34

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/duyserveacc/algos/internal/testrunner"
)

/*
===============================================================================
Find First and Last Position of Element in Sorted Array

Source: LeetCode
Problem ID: 34
Source URL: https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/

Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.

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
===============================================================================
*/

// Problem-specific support.
//
// Keep shared judge types and parsing or formatting helpers here. Name adapters
// after the problem. Put setup used only by the current attempt at the top of
// solveCurrentAttempt below.
//
// This intentionally unsolved practice archive follows the workbench template.

func parseFindFirstAndLastPositionOfElementInSortedArrayInput(input string) ([]int, int) {
	var args []json.RawMessage
	if err := json.Unmarshal([]byte(input), &args); err != nil {
		panic(fmt.Errorf("parse find-first-and-last-position-of-element-in-sorted-array input: %w", err))
	}
	if len(args) != 2 {
		panic(fmt.Sprintf("parse find-first-and-last-position-of-element-in-sorted-array input: got %d arguments, want 2", len(args)))
	}
	var nums []int
	argumentError0 := json.Unmarshal(args[0], &nums)
	if argumentError0 != nil {
		panic(fmt.Errorf("parse find-first-and-last-position-of-element-in-sorted-array argument nums: %w", argumentError0))
	}
	var target int
	argumentError1 := json.Unmarshal(args[1], &target)
	if argumentError1 != nil {
		panic(fmt.Errorf("parse find-first-and-last-position-of-element-in-sorted-array argument target: %w", argumentError1))
	}
	return nums, target
}

func formatFindFirstAndLastPositionOfElementInSortedArrayOutput(bounds []int) string {
	output, err := json.Marshal(bounds)
	if err != nil {
		panic(fmt.Errorf("format find-first-and-last-position-of-element-in-sorted-array output: %w", err))
	}
	return string(output)
}

func maximumFindFirstAndLastPositionOfElementInSortedArrayCase() testrunner.Case {
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
		panic(fmt.Errorf("build maximum find-first-and-last-position-of-element-in-sorted-array case: %w", err))
	}
	return testrunner.Case{
		Name:  "Boundary: Maximum length contains a large middle target run",
		Input: string(input),
		Want:  `[25000,74999]`,
	}
}

// sampleCases contains the examples that every registered solution must pass.
// Inputs encode the judge arguments as a positional JSON array.
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
		Name:  "Boundary: Minimum and maximum values",
		Input: `[[-1000000000,0,1000000000],1000000000]`,
		Want:  `[2,2]`,
	},
	maximumFindFirstAndLastPositionOfElementInSortedArrayCase(),
}

// archiveReady remains false until the current attempt, three agent solutions,
// analysis, and regression coverage are complete and verified.
const archiveReady = false

// Registered implementations run against the same cases.
var currentAttemptSolution = testrunner.Solution{
	Name: "current attempt",
	Run:  solveCurrentAttempt,
}

// agentSolutions is intentionally empty until an AI review adds direct,
// improved, and preferred implementations.
var agentSolutions = []testrunner.Solution{}

var solutions = append([]testrunner.Solution{currentAttemptSolution}, agentSolutions...)

// =============================================================================
// CURRENT ATTEMPT
//
// Implement your solution here. Keep this body intact when requesting review.
// =============================================================================

func solveCurrentAttempt(input string) string {
	nums, target := parseFindFirstAndLastPositionOfElementInSortedArrayInput(input)
	_ = nums
	_ = target

	// Implement your algorithm with the typed values above.
	return ""
}

// =============================================================================
// AGENT SOLUTIONS
//
// Intentionally blank until you ask an AI agent to review this problem.
// =============================================================================

/*
===============================================================================
AGENT REVIEW AND COMPLEXITY ANALYSIS

Current attempt:
- What is correct:
- First failing assumption or bottleneck:
- Relevant concept and invariant:

Solution 1 - Direct:
- Correctness:
- Concrete example:
- Time complexity:
- Auxiliary-space complexity:

Solution 2 - Improved:
- Correctness:
- Concrete example:
- Time complexity:
- Auxiliary-space complexity:

Solution 3 - Preferred:
- Correctness:
- Concrete example:
- Time complexity:
- Auxiliary-space complexity:

AGENT REVIEW AND COMPLEXITY ANALYSIS END
===============================================================================
*/

func main() {
	if err := testrunner.Run(os.Stdout, sampleCases, solutions); err != nil {
		fmt.Fprintln(os.Stderr, "workbench:", err)
		os.Exit(1)
	}
}
