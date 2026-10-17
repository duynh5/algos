//go:build ignore

// Difficulty: Medium
// Source: leetcode
// Problem ID: 162

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/duyserveacc/algos/internal/testrunner"
)

/*
===============================================================================
Find Peak Element

Source: LeetCode
Problem ID: 162
Source URL: https://leetcode.com/problems/find-peak-element/description/

A peak element is an element that is strictly greater than its neighbors.

Given a 0-indexed integer array nums, find a peak element, and return its index. If the array contains multiple peaks, return the index to any of the peaks.

You may imagine that nums[-1] = nums[n] = -infinity. In other words, an element is always considered to be strictly greater than a neighbor that is outside the array.

You must write an algorithm that runs in O(log n) time.

Example 1:

Input: nums = [1,2,3,1]
Output: 2
Explanation: 3 is a peak element and your function should return the index number 2.

Example 2:

Input: nums = [1,2,1,3,5,6,4]
Output: 5
Explanation: Your function can return either index number 1 where the peak element is 2, or index number 5 where the peak element is 6.

Constraints:

- 1 <= nums.length <= 1000

- -2^31 <= nums[i] <= 2^31 - 1

- nums[i] != nums[i + 1] for all valid i.
===============================================================================
*/

// Problem-specific support.
//
// Keep shared judge types and parsing or formatting helpers here. Name adapters
// after the problem. Put setup used only by the current attempt at the top of
// solveCurrentAttempt below.
//
// This intentionally unsolved practice archive follows the workbench template.

func parseFindPeakElementInput(input string) []int {
	var args []json.RawMessage
	if err := json.Unmarshal([]byte(input), &args); err != nil {
		panic(fmt.Errorf("parse find-peak-element input: %w", err))
	}
	if len(args) != 1 {
		panic(fmt.Sprintf("parse find-peak-element input: got %d arguments, want 1", len(args)))
	}
	var nums []int
	argumentError0 := json.Unmarshal(args[0], &nums)
	if argumentError0 != nil {
		panic(fmt.Errorf("parse find-peak-element argument nums: %w", argumentError0))
	}
	return nums
}

func maximumFindPeakElementCase() testrunner.Case {
	nums := make([]int, 1000)
	for index := range nums {
		nums[index] = index - 500
	}
	input, err := json.Marshal([]any{nums})
	if err != nil {
		panic(fmt.Errorf("build maximum find-peak-element input: %w", err))
	}
	return testrunner.Case{Name: "Boundary: Maximum strictly increasing array", Input: string(input), Want: `999`}
}

func formatFindPeakElementOutput(result int) string {
	return testrunner.FormatJSONOutput(result)
}

// sampleCases contains the examples that every registered solution must pass.
// Inputs encode the judge arguments as a positional JSON array.
var sampleCases = []testrunner.Case{
	{
		Name:  "Official example 1",
		Input: `[[1,2,3,1]]`,
		Want:  `2`,
	},
	{
		Name:  "Official example 2",
		Input: `[[1,2,1,3,5,6,4]]`,
		Want:  `5`,
	},
	{
		Name:  "Additional: Single value",
		Input: `[[5]]`,
		Want:  `0`,
	},
	{
		Name:  "Additional: Strictly decreasing",
		Input: `[[5,4,3,2]]`,
		Want:  `0`,
	},
	{
		Name:  "Additional: Strictly increasing",
		Input: `[[1,2,3,4]]`,
		Want:  `3`,
	},
	{
		Name:  "Additional: Unique interior peak",
		Input: `[[-3,-1,4,2,0]]`,
		Want:  `2`,
	},
	{
		Name:  "Boundary: Two values peak on the left",
		Input: `[[2147483647,-2147483648]]`,
		Want:  `0`,
	},
	{
		Name:  "Boundary: Two values peak on the right",
		Input: `[[-2147483648,2147483647]]`,
		Want:  `1`,
	},
	{
		Name:  "Adversarial: Peak immediately left of the midpoint",
		Input: `[[1,5,4,3,2]]`,
		Want:  `1`,
	},
	maximumFindPeakElementCase(),
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
	nums := parseFindPeakElementInput(input)
	_ = nums

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
