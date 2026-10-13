//go:build ignore

// Difficulty: Medium
// Source: leetcode
// Problem ID: 167

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/duyserveacc/algos/internal/testrunner"
)

/*
===============================================================================
Two Sum II - Input Array Is Sorted

Source: LeetCode
Problem ID: 167
Source URL: https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/

You are given a 1-indexed array of integers numbers that is already sorted in non-decreasing order.

Find two numbers such that they add up to a specific target number. Let these two numbers be numbers[index_1] and numbers[index_2] where 1 <= index_1 < index_2 <= numbers.length.

Return the indices of the two numbers index_1 and index_2 as an integer array [index_1, index_2] of length 2.

The tests are generated such that there is exactly one solution. You may not use the same element twice.

Your solution must use only constant extra space.

Example 1:

Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Explanation: The sum of 2 and 7 is 9. Therefore, index_1 = 1, index_2 = 2. We return [1, 2].

Example 2:

Input: numbers = [2,3,4], target = 6
Output: [1,3]
Explanation: The sum of 2 and 4 is 6. Therefore index_1 = 1, index_2 = 3. We return [1, 3].

Example 3:

Input: numbers = [-1,0], target = -1
Output: [1,2]
Explanation: The sum of -1 and 0 is -1. Therefore index_1 = 1, index_2 = 2. We return [1, 2].

Constraints:

- 2 <= numbers.length <= 3 * 10^4

- -1000 <= numbers[i] <= 1000

- numbers is sorted in non-decreasing order.

- -1000 <= target <= 1000

- The tests are generated such that there is exactly one solution.
===============================================================================
*/

// Problem-specific support.
//
// Keep shared judge types and parsing or formatting helpers here. Name adapters
// after the problem. Put setup used only by the current attempt at the top of
// solveCurrentAttempt below.
//
// This intentionally unsolved practice archive follows the workbench template.

func parseTwoSumIiInputArrayIsSortedInput(input string) ([]int, int) {
	var args []json.RawMessage
	if err := json.Unmarshal([]byte(input), &args); err != nil {
		panic(fmt.Errorf("parse two-sum-ii-input-array-is-sorted input: %w", err))
	}
	if len(args) != 2 {
		panic(fmt.Sprintf("parse two-sum-ii-input-array-is-sorted input: got %d arguments, want 2", len(args)))
	}
	var numbers []int
	argumentError0 := json.Unmarshal(args[0], &numbers)
	if argumentError0 != nil {
		panic(fmt.Errorf("parse two-sum-ii-input-array-is-sorted argument numbers: %w", argumentError0))
	}
	var target int
	argumentError1 := json.Unmarshal(args[1], &target)
	if argumentError1 != nil {
		panic(fmt.Errorf("parse two-sum-ii-input-array-is-sorted argument target: %w", argumentError1))
	}
	return numbers, target
}

// sampleCases contains the examples that every registered solution must pass.
// Inputs encode the judge arguments as a positional JSON array.
var sampleCases = []testrunner.Case{
	{
		Name:  "Official example 1",
		Input: `[[2,7,11,15],9]`,
		Want:  `[1,2]`,
	},
	{
		Name:  "Official example 2",
		Input: `[[2,3,4],6]`,
		Want:  `[1,3]`,
	},
	{
		Name:  "Official example 3",
		Input: `[[-1,0],-1]`,
		Want:  `[1,2]`,
	},
	{
		Name:  "Additional: Minimum pair",
		Input: `[[1,2],3]`,
		Want:  `[1,2]`,
	},
	{
		Name:  "Additional: Duplicate values",
		Input: `[[3,3],6]`,
		Want:  `[1,2]`,
	},
	{
		Name:  "Additional: Negative and positive",
		Input: `[[-5,-1,4,8],3]`,
		Want:  `[2,3]`,
	},
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
	numbers, target := parseTwoSumIiInputArrayIsSortedInput(input)
	_ = numbers
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
