//go:build ignore

// Difficulty: Medium
// Source: leetcode
// Problem ID: 300

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/duyserveacc/algos/internal/testrunner"
)

/*
===============================================================================
Longest Increasing Subsequence

Source: LeetCode
Problem ID: 300
Source URL: https://leetcode.com/problems/longest-increasing-subsequence/description/

Given an integer array nums, return the length of the longest strictly increasing subsequence.

Example 1:

Input: nums = [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.

Example 2:

Input: nums = [0,1,0,3,2,3]
Output: 4

Example 3:

Input: nums = [7,7,7,7,7,7,7]
Output: 1

Constraints:

- 1 <= nums.length <= 2500

- -10^4 <= nums[i] <= 10^4

Follow up: Can you come up with an algorithm that runs in O(n log(n)) time complexity?
===============================================================================
*/

// Problem-specific support.
//
// Keep shared judge types and parsing or formatting helpers here. Name adapters
// after the problem. Put setup used only by the current attempt at the top of
// solveCurrentAttempt below.
//
// This intentionally unsolved practice archive follows the workbench template.

func parseLongestIncreasingSubsequenceInput(input string) []int {
	var args []json.RawMessage
	if err := json.Unmarshal([]byte(input), &args); err != nil {
		panic(fmt.Errorf("parse longest-increasing-subsequence input: %w", err))
	}
	if len(args) != 1 {
		panic(fmt.Sprintf("parse longest-increasing-subsequence input: got %d arguments, want 1", len(args)))
	}
	var nums []int
	argumentError0 := json.Unmarshal(args[0], &nums)
	if argumentError0 != nil {
		panic(fmt.Errorf("parse longest-increasing-subsequence argument nums: %w", argumentError0))
	}
	return nums
}

func maximumLongestIncreasingSubsequenceCase() testrunner.Case {
	nums := make([]int, 2500)
	for index := range nums {
		nums[index] = index - 1250
	}
	input, err := json.Marshal([]any{nums})
	if err != nil {
		panic(fmt.Errorf("build maximum longest-increasing-subsequence input: %w", err))
	}
	return testrunner.Case{Name: "Boundary: Maximum strictly increasing array", Input: string(input), Want: `2500`}
}

func formatLongestIncreasingSubsequenceOutput(result int) string {
	return testrunner.FormatJSONOutput(result)
}

// sampleCases contains the examples that every registered solution must pass.
// Inputs encode the judge arguments as a positional JSON array.
var sampleCases = []testrunner.Case{
	{
		Name:  "Official example 1",
		Input: `[[10,9,2,5,3,7,101,18]]`,
		Want:  `4`,
	},
	{
		Name:  "Official example 2",
		Input: `[[0,1,0,3,2,3]]`,
		Want:  `4`,
	},
	{
		Name:  "Official example 3",
		Input: `[[7,7,7,7,7,7,7]]`,
		Want:  `1`,
	},
	{
		Name:  "Additional: Single value",
		Input: `[[5]]`,
		Want:  `1`,
	},
	{
		Name:  "Additional: Strictly decreasing",
		Input: `[[5,4,3,2,1]]`,
		Want:  `1`,
	},
	{
		Name:  "Additional: Duplicates do not increase length",
		Input: `[[2,2,2]]`,
		Want:  `1`,
	},
	{
		Name:  "Additional: Strictly increasing",
		Input: `[[-3,-2,-1,0,1]]`,
		Want:  `5`,
	},
	{
		Name:  "Adversarial: Best tails replace earlier larger values",
		Input: `[[3,5,6,2,4,7]]`,
		Want:  `4`,
	},
	{
		Name:  "Adversarial: Alternating highs and lows",
		Input: `[[1,10,2,9,3,8,4,7]]`,
		Want:  `5`,
	},
	{
		Name:  "Boundary: Signed value limits",
		Input: `[[-10000,10000]]`,
		Want:  `2`,
	},
	maximumLongestIncreasingSubsequenceCase(),
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
	nums := parseLongestIncreasingSubsequenceInput(input)
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
