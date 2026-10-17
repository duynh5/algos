//go:build ignore

// Difficulty: Easy
// Source: leetcode
// Problem ID: 136

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/duyserveacc/algos/internal/testrunner"
)

/*
===============================================================================
Single Number

Source: LeetCode
Problem ID: 136
Source URL: https://leetcode.com/problems/single-number/description/

Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

You must implement a solution with a linear runtime complexity and use only constant extra space.

Example 1:

Input: nums = [2,2,1]

Output: 1

Example 2:

Input: nums = [4,1,2,1,2]

Output: 4

Example 3:

Input: nums = [1]

Output: 1

Constraints:

- 1 <= nums.length <= 3 * 10^4

- -3 * 10^4 <= nums[i] <= 3 * 10^4

- Each element in the array appears twice except for one element which appears only once.
===============================================================================
*/

// Problem-specific support.
//
// Keep shared judge types and parsing or formatting helpers here. Name adapters
// after the problem. Put setup used only by the current attempt at the top of
// solveCurrentAttempt below.
//
// This intentionally unsolved practice archive follows the workbench template.

func parseSingleNumberInput(input string) []int {
	var args []json.RawMessage
	if err := json.Unmarshal([]byte(input), &args); err != nil {
		panic(fmt.Errorf("parse single-number input: %w", err))
	}
	if len(args) != 1 {
		panic(fmt.Sprintf("parse single-number input: got %d arguments, want 1", len(args)))
	}
	var nums []int
	argumentError0 := json.Unmarshal(args[0], &nums)
	if argumentError0 != nil {
		panic(fmt.Errorf("parse single-number argument nums: %w", argumentError0))
	}
	return nums
}

func maximumSingleNumberCase() testrunner.Case {
	nums := make([]int, 0, 29999)
	for value := -14999; value <= -1; value++ {
		nums = append(nums, value)
	}
	nums = append(nums, 30000)
	for value := -1; value >= -14999; value-- {
		nums = append(nums, value)
	}
	input, err := json.Marshal([]any{nums})
	if err != nil {
		panic(fmt.Errorf("build maximum single-number case: %w", err))
	}
	return testrunner.Case{
		Name:  "Boundary: Maximum valid odd length",
		Input: string(input),
		Want:  `30000`,
	}
}

func formatSingleNumberOutput(result int) string {
	return testrunner.FormatJSONOutput(result)
}

// sampleCases contains the examples that every registered solution must pass.
// Inputs encode the judge arguments as a positional JSON array.
var sampleCases = []testrunner.Case{
	{
		Name:  "Official example 1",
		Input: `[[2,2,1]]`,
		Want:  `1`,
	},
	{
		Name:  "Official example 2",
		Input: `[[4,1,2,1,2]]`,
		Want:  `4`,
	},
	{
		Name:  "Official example 3",
		Input: `[[1]]`,
		Want:  `1`,
	},
	{
		Name:  "Additional: Single value",
		Input: `[[7]]`,
		Want:  `7`,
	},
	{
		Name:  "Additional: Negative unique value",
		Input: `[[-1,2,2]]`,
		Want:  `-1`,
	},
	{
		Name:  "Additional: Zero paired",
		Input: `[[0,1,0]]`,
		Want:  `1`,
	},
	{
		Name:  "Additional: Zero is unique",
		Input: `[[-1,-1,0,2,2]]`,
		Want:  `0`,
	},
	{
		Name:  "Additional: Negative values are paired",
		Input: `[[-2,3,-2]]`,
		Want:  `3`,
	},
	{
		Name:  "Boundary: Minimum value is unique",
		Input: `[[-30000,1,1,30000,30000]]`,
		Want:  `-30000`,
	},
	{
		Name:  "Boundary: Maximum value is unique",
		Input: `[[-30000,-30000,30000]]`,
		Want:  `30000`,
	},
	{
		Name:  "Adversarial: Duplicate pairs are separated",
		Input: `[[4,1,2,4,2]]`,
		Want:  `1`,
	},
	maximumSingleNumberCase(),
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
	nums := parseSingleNumberInput(input)
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
