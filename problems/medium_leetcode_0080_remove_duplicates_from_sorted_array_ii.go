//go:build ignore

// Difficulty: Medium
// Source: leetcode
// Problem ID: 80

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/duyserveacc/algos/internal/testrunner"
)

/*
===============================================================================
Remove Duplicates from Sorted Array II

Source: LeetCode
Problem ID: 80
Source URL: https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/description/

Given an integer array nums sorted in non-decreasing order, remove some duplicates in-place such that each unique element appears at most twice. The relative order of the elements should be kept the same.

Since it is impossible to change the length of the array in some languages, you must instead have the result be placed in the first part of the array nums. More formally, if there are k elements after removing the duplicates, then the first k elements of nums should hold the final result. It does not matter what you leave beyond the first k elements.

Return k after placing the final result in the first k slots of nums.

Do not allocate extra space for another array. You must do this by modifying the input array in-place with O(1) extra memory.

Custom Judge:

The judge will test your solution with the following code:

int[] nums = [...]; // Input array
int[] expectedNums = [...]; // The expected answer with correct length

int k = removeDuplicates(nums); // Calls your implementation

assert k == expectedNums.length;
for (int i = 0; i < k; i++) {
assert nums[i] == expectedNums[i];
}

If all assertions pass, then your solution will be accepted.

Example 1:

Input: nums = [1,1,1,2,2,3]
Output: 5, nums = [1,1,2,2,3,_]
Explanation: Your function should return k = 5, with the first five elements of nums being 1, 1, 2, 2 and 3 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).

Example 2:

Input: nums = [0,0,1,1,1,1,2,3,3]
Output: 7, nums = [0,0,1,1,2,3,3,_,_]
Explanation: Your function should return k = 7, with the first seven elements of nums being 0, 0, 1, 1, 2, 3 and 3 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).

Constraints:

- 1 <= nums.length <= 3 * 10^4

- -10^4 <= nums[i] <= 10^4

- nums is sorted in non-decreasing order.
===============================================================================
*/

// Problem-specific support.
//
// Keep shared judge types and parsing or formatting helpers here. Name adapters
// after the problem. Put setup used only by the current attempt at the top of
// solveCurrentAttempt below.
//
// This intentionally unsolved practice archive follows the workbench template.

func parseRemoveDuplicatesFromSortedArrayIiInput(input string) []int {
	var args []json.RawMessage
	if err := json.Unmarshal([]byte(input), &args); err != nil {
		panic(fmt.Errorf("parse remove-duplicates-from-sorted-array-ii input: %w", err))
	}
	if len(args) != 1 {
		panic(fmt.Sprintf("parse remove-duplicates-from-sorted-array-ii input: got %d arguments, want 1", len(args)))
	}
	var nums []int
	argumentError0 := json.Unmarshal(args[0], &nums)
	if argumentError0 != nil {
		panic(fmt.Errorf("parse remove-duplicates-from-sorted-array-ii argument nums: %w", argumentError0))
	}
	return nums
}

func formatRemoveDuplicatesFromSortedArrayIiOutput(nums []int, k int) string {
	if k < 0 || k > len(nums) {
		panic(fmt.Sprintf("format remove-duplicates-from-sorted-array-ii output: k=%d outside [0,%d]", k, len(nums)))
	}
	var output strings.Builder
	fmt.Fprintf(&output, "%d, nums = [", k)
	for index := range nums {
		if index > 0 {
			output.WriteByte(',')
		}
		if index < k {
			fmt.Fprint(&output, nums[index])
		} else {
			output.WriteByte('_')
		}
	}
	output.WriteByte(']')
	return output.String()
}

func maximumRemoveDuplicatesFromSortedArrayIiCase() testrunner.Case {
	nums := make([]int, 30000)
	for index := range nums {
		nums[index] = -10000
	}
	input, err := json.Marshal([]any{nums})
	if err != nil {
		panic(fmt.Errorf("build maximum remove-duplicates-from-sorted-array-ii input: %w", err))
	}
	return testrunner.Case{
		Name:  "Boundary: Maximum equal run keeps exactly two values",
		Input: string(input),
		Want:  formatRemoveDuplicatesFromSortedArrayIiOutput(nums, 2),
	}
}

// sampleCases contains the examples that every registered solution must pass.
// Inputs encode the judge arguments as a positional JSON array.
var sampleCases = []testrunner.Case{
	{
		Name:  "Official example 1",
		Input: `[[1,1,1,2,2,3]]`,
		Want:  `5, nums = [1,1,2,2,3,_]`,
	},
	{
		Name:  "Official example 2",
		Input: `[[0,0,1,1,1,1,2,3,3]]`,
		Want:  `7, nums = [0,0,1,1,2,3,3,_,_]`,
	},
	{
		Name:  "Additional: Minimum array",
		Input: `[[1]]`,
		Want:  `1, nums = [1]`,
	},
	{
		Name:  "Additional: Three equal values",
		Input: `[[1,1,1]]`,
		Want:  `2, nums = [1,1,_]`,
	},
	{
		Name:  "Additional: No removals",
		Input: `[[1,1,2,2]]`,
		Want:  `4, nums = [1,1,2,2]`,
	},
	{
		Name:  "Additional: Longer equal run keeps two values",
		Input: `[[1,1,1,1,1]]`,
		Want:  `2, nums = [1,1,_,_,_]`,
	},
	{
		Name:  "Additional: Mixed duplicate runs preserve sorted order",
		Input: `[[-2,-2,-2,-1,-1,0,0,0,0]]`,
		Want:  `6, nums = [-2,-2,-1,-1,0,0,_,_,_]`,
	},
	{
		Name:  "Boundary: Distinct signed limits need no removals",
		Input: `[[-10000,0,10000]]`,
		Want:  `3, nums = [-10000,0,10000]`,
	},
	{
		Name:  "Adversarial: Excess duplicates occur at the end",
		Input: `[[1,2,3,3,3]]`,
		Want:  `4, nums = [1,2,3,3,_]`,
	},
	maximumRemoveDuplicatesFromSortedArrayIiCase(),
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
	nums := parseRemoveDuplicatesFromSortedArrayIiInput(input)
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
