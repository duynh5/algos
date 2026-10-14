//go:build ignore

// Difficulty: Medium
// Source: leetcode
// Problem ID: 128

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/duyserveacc/algos/internal/testrunner"
)

/*
===============================================================================
Longest Consecutive Sequence

Source: LeetCode
Problem ID: 128
Source URL: https://leetcode.com/problems/longest-consecutive-sequence/description/

Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

You must write an algorithm that runs in O(n) time.

Example 1:

Input: nums = [100,4,200,1,3,2]
Output: 4
Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.

Example 2:

Input: nums = [0,3,7,2,5,8,4,6,0,1]
Output: 9

Example 3:

Input: nums = [1,0,1,2]
Output: 3

Constraints:

- 0 <= nums.length <= 10^5

- -10^9 <= nums[i] <= 10^9
===============================================================================
*/

// Problem-specific support.
//
// Keep shared judge types and parsing or formatting helpers here. Name adapters
// after the problem. Put setup used only by the current attempt at the top of
// solveCurrentAttempt below.
//
// This intentionally unsolved practice archive follows the workbench template.

func parseLongestConsecutiveSequenceInput(input string) []int {
	var args []json.RawMessage
	if err := json.Unmarshal([]byte(input), &args); err != nil {
		panic(fmt.Errorf("parse longest-consecutive-sequence input: %w", err))
	}
	if len(args) != 1 {
		panic(fmt.Sprintf("parse longest-consecutive-sequence input: got %d arguments, want 1", len(args)))
	}
	var nums []int
	argumentError0 := json.Unmarshal(args[0], &nums)
	if argumentError0 != nil {
		panic(fmt.Errorf("parse longest-consecutive-sequence argument nums: %w", argumentError0))
	}
	return nums
}

func maximumLongestConsecutiveSequenceCase() testrunner.Case {
	nums := make([]int, 100000)
	for index := range nums {
		nums[index] = 49999 - index
	}
	input, err := json.Marshal([]any{nums})
	if err != nil {
		panic(fmt.Errorf("build maximum longest-consecutive-sequence input: %w", err))
	}
	return testrunner.Case{Name: "Boundary: Maximum descending consecutive run", Input: string(input), Want: `100000`}
}

// sampleCases contains the examples that every registered solution must pass.
// Inputs encode the judge arguments as a positional JSON array.
var sampleCases = []testrunner.Case{
	{
		Name:  "Official example 1",
		Input: `[[100,4,200,1,3,2]]`,
		Want:  `4`,
	},
	{
		Name:  "Official example 2",
		Input: `[[0,3,7,2,5,8,4,6,0,1]]`,
		Want:  `9`,
	},
	{
		Name:  "Official example 3",
		Input: `[[1,0,1,2]]`,
		Want:  `3`,
	},
	{
		Name:  "Additional: Empty array",
		Input: `[[]]`,
		Want:  `0`,
	},
	{
		Name:  "Additional: Duplicates do not extend a run",
		Input: `[[1,2,0,1]]`,
		Want:  `3`,
	},
	{
		Name:  "Additional: Negative run",
		Input: `[[-2,-1,0,2]]`,
		Want:  `3`,
	},
	{
		Name:  "Minimum: Single value",
		Input: `[[42]]`,
		Want:  `1`,
	},
	{
		Name:  "Boundary: Separated integer limits",
		Input: `[[-1000000000,1000000000]]`,
		Want:  `1`,
	},
	{
		Name:  "Adversarial: Unsorted duplicates surround the longest run",
		Input: `[[5,2,99,3,4,3,1,100]]`,
		Want:  `5`,
	},
	{
		Name:  "Additional: Competing runs choose the longer one",
		Input: `[[-3,-2,10,11,12,13,-1]]`,
		Want:  `4`,
	},
	maximumLongestConsecutiveSequenceCase(),
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
	nums := parseLongestConsecutiveSequenceInput(input)
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
