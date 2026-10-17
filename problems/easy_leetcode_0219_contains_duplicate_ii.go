//go:build ignore

// Difficulty: Easy
// Source: leetcode
// Problem ID: 219

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/duyserveacc/algos/internal/testrunner"
)

/*
===============================================================================
Contains Duplicate II

Source: LeetCode
Problem ID: 219
Source URL: https://leetcode.com/problems/contains-duplicate-ii/description/

Given an integer array nums and an integer k, return true if there are two distinct indices i and j in the array such that nums[i] == nums[j] and abs(i - j) <= k.

Example 1:

Input: nums = [1,2,3,1], k = 3
Output: true

Example 2:

Input: nums = [1,0,1,1], k = 1
Output: true

Example 3:

Input: nums = [1,2,3,1,2,3], k = 2
Output: false

Constraints:

- 1 <= nums.length <= 10^5

- -10^9 <= nums[i] <= 10^9

- 0 <= k <= 10^5
===============================================================================
*/

// Problem-specific support.
//
// Keep shared judge types and parsing or formatting helpers here. Name adapters
// after the problem. Put setup used only by the current attempt at the top of
// solveCurrentAttempt below.
//
// This intentionally unsolved practice archive follows the workbench template.

func parseContainsDuplicateIiInput(input string) ([]int, int) {
	var args []json.RawMessage
	if err := json.Unmarshal([]byte(input), &args); err != nil {
		panic(fmt.Errorf("parse contains-duplicate-ii input: %w", err))
	}
	if len(args) != 2 {
		panic(fmt.Sprintf("parse contains-duplicate-ii input: got %d arguments, want 2", len(args)))
	}
	var nums []int
	argumentError0 := json.Unmarshal(args[0], &nums)
	if argumentError0 != nil {
		panic(fmt.Errorf("parse contains-duplicate-ii argument nums: %w", argumentError0))
	}
	var k int
	argumentError1 := json.Unmarshal(args[1], &k)
	if argumentError1 != nil {
		panic(fmt.Errorf("parse contains-duplicate-ii argument k: %w", argumentError1))
	}
	return nums, k
}

func maximumContainsDuplicateIiCase() testrunner.Case {
	nums := make([]int, 100000)
	for index := range nums {
		nums[index] = index
	}
	nums[len(nums)-1] = nums[0]
	input, err := json.Marshal([]any{nums, len(nums) - 1})
	if err != nil {
		panic(fmt.Errorf("build maximum contains-duplicate-ii case: %w", err))
	}
	return testrunner.Case{
		Name:  "Boundary: Maximum length with farthest possible duplicate",
		Input: string(input),
		Want:  `true`,
	}
}

func formatContainsDuplicateIiOutput(result bool) string {
	return testrunner.FormatJSONOutput(result)
}

// sampleCases contains the examples that every registered solution must pass.
// Inputs encode the judge arguments as a positional JSON array.
var sampleCases = []testrunner.Case{
	{
		Name:  "Official example 1",
		Input: `[[1,2,3,1],3]`,
		Want:  `true`,
	},
	{
		Name:  "Official example 2",
		Input: `[[1,0,1,1],1]`,
		Want:  `true`,
	},
	{
		Name:  "Official example 3",
		Input: `[[1,2,3,1,2,3],2]`,
		Want:  `false`,
	},
	{
		Name:  "Additional: Single value and zero distance",
		Input: `[[1],0]`,
		Want:  `false`,
	},
	{
		Name:  "Additional: Duplicate exists but k is zero",
		Input: `[[1,1],0]`,
		Want:  `false`,
	},
	{
		Name:  "Additional: Duplicate exactly k apart",
		Input: `[[1,2,1],2]`,
		Want:  `true`,
	},
	{
		Name:  "Boundary: Duplicate one position beyond k",
		Input: `[[1,2,3,1],2]`,
		Want:  `false`,
	},
	{
		Name:  "Additional: Later occurrence creates a qualifying pair",
		Input: `[[1,2,1,3,1],2]`,
		Want:  `true`,
	},
	{
		Name:  "Additional: Negative and extreme values compare normally",
		Input: `[[-1000000000,0,1000000000,-1000000000],3]`,
		Want:  `true`,
	},
	{
		Name:  "Boundary: Maximum k with no duplicates",
		Input: `[[1,2,3],100000]`,
		Want:  `false`,
	},
	maximumContainsDuplicateIiCase(),
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
	nums, k := parseContainsDuplicateIiInput(input)
	_ = nums
	_ = k

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
