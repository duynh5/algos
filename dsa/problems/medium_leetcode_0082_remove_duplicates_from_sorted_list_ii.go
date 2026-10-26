//go:build ignore

// Difficulty: Medium
// Source: leetcode
// Problem ID: 82

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/duyserveacc/algos/dsa/internal/testrunner"
)

/*
======================================================================
Remove Duplicates from Sorted List II

Source: LeetCode
Problem ID: 82
Source URL:
../../REFERENCE_LINKS.md#leetcode-source-0082

You are given the head of a sorted linked list.

Delete all nodes that have duplicate numbers, leaving only distinct
numbers from the original list.

Return the linked list sorted as well.

Example 1:

Input: head = [1,2,3,3,4,4,5]
Output: [1,2,5]

Example 2:

Input: head = [1,1,1,2,3]
Output: [2,3]

Constraints:

- The number of nodes in the list is in the range [0, 300].

- -100 <= Node.val <= 100

- The list is guaranteed to be sorted in ascending order.
======================================================================
*/

// Problem-specific support.
//
// Keep shared judge types and parsing or formatting helpers here.
// Name adapters after the problem. Put setup used only by the current
// attempt at the top of solveCurrentAttempt below.
//
// This intentionally unsolved practice archive follows the workbench
// template.

type ListNode struct {
	Val  int
	Next *ListNode
}

func parseRemoveDuplicatesFromSortedListIiInput(
	input string,
) *ListNode {
	var args []json.RawMessage
	if err := json.Unmarshal([]byte(input), &args); err != nil {
		panic(
			fmt.Errorf("parse remove-duplicates-from-sorted-list-ii "+
				"input: %w", err),
		)
	}
	if len(args) != 1 {
		panic(
			fmt.Sprintf(
				"parse remove-duplicates-from-sorted-list-ii "+
					"input: got %d arguments, want 1",
				len(args),
			),
		)
	}
	head, argumentError0 := decodeListNode(args[0])
	if argumentError0 != nil {
		panic(
			fmt.Errorf("parse remove-duplicates-from-sorted-list-ii "+
				"argument head: %w", argumentError0),
		)
	}
	return head
}

func decodeListNode(raw json.RawMessage) (*ListNode, error) {
	var values []int
	if err := json.Unmarshal(raw, &values); err != nil {
		return nil, err
	}
	dummy := &ListNode{}
	tail := dummy
	for _, value := range values {
		tail.Next = &ListNode{Val: value}
		tail = tail.Next
	}
	return dummy.Next, nil
}

func formatRemoveDuplicatesFromSortedListIiOutput(
	head *ListNode,
) string {
	values := make([]int, 0)
	for node := head; node != nil; node = node.Next {
		if len(values) == 300 {
			panic("format remove-duplicates-from-sorted-list-ii " +
				"output: list exceeds 300 nodes or contains a " +
				"cycle")
		}
		values = append(values, node.Val)
	}
	output, err := json.Marshal(values)
	if err != nil {
		panic(
			fmt.Errorf(
				"format remove-duplicates-from-sorted-list-ii "+
					"output: %w",
				err,
			),
		)
	}
	return string(output)
}

func maximumRemoveDuplicatesFromSortedListIiCase() testrunner.Case {
	values := make([]int, 0, 300)
	for value := -100; value < 50; value++ {
		values = append(values, value, value)
	}
	input, err := json.Marshal([]any{values})
	if err != nil {
		panic(
			fmt.Errorf("build maximum remove-duplicates-from-sorted-"+
				"list-ii input: %w", err),
		)
	}
	return testrunner.Case{
		Name:  "Boundary: Maximum list removes every duplicated pair",
		Input: string(input),
		Want:  formatRemoveDuplicatesFromSortedListIiOutput(nil),
	}
}

// sampleCases contains the examples that every registered solution
// must pass. Inputs encode the judge arguments as a positional JSON
// array.
var sampleCases = []testrunner.Case{
	{
		Name:  "Official example 1",
		Input: `[[1,2,3,3,4,4,5]]`,
		Want:  `[1,2,5]`,
	},
	{
		Name:  "Official example 2",
		Input: `[[1,1,1,2,3]]`,
		Want:  `[2,3]`,
	},
	{
		Name:  "Additional: Empty list",
		Input: `[[]]`,
		Want:  `[]`,
	},
	{
		Name:  "Additional: Every node duplicated",
		Input: `[[1,1]]`,
		Want:  `[]`,
	},
	{
		Name:  "Additional: Duplicate suffix",
		Input: `[[1,2,3,3]]`,
		Want:  `[1,2]`,
	},
	{
		Name:  "Additional: Single node remains",
		Input: `[[1]]`,
		Want:  `[1]`,
	},
	{
		Name:  "Additional: Duplicate prefix is removed",
		Input: `[[1,1,2]]`,
		Want:  `[2]`,
	},
	{
		Name:  "Additional: A run longer than two is removed entirely",
		Input: `[[1,1,1,2]]`,
		Want:  `[2]`,
	},
	{
		Name: "Adversarial: Separate duplicate runs " +
			"preserve intervening uniques",
		Input: `[[-100,-100,-99,0,1,1,100]]`,
		Want:  `[-99,0,100]`,
	},
	{
		Name:  "Boundary: Distinct signed limits remain",
		Input: `[[-100,0,100]]`,
		Want:  `[-100,0,100]`,
	},
	maximumRemoveDuplicatesFromSortedListIiCase(),
}

// archiveReady remains false until the current attempt, three agent
// solutions, analysis, and regression coverage are complete and
// verified.
const archiveReady = false

// Registered implementations run against the same cases.
var currentAttemptSolution = testrunner.Solution{
	Name: "current attempt",
	Run:  solveCurrentAttempt,
}

// agentSolutions is intentionally empty until an AI review adds
// direct, improved, and preferred implementations.
var agentSolutions = []testrunner.Solution{}

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
	head := parseRemoveDuplicatesFromSortedListIiInput(input)
	_ = head

	// Implement your algorithm with the typed values above.
	return ""
}

// ===================================================================
// AGENT SOLUTIONS
//
// Intentionally blank until you ask an AI agent to review this
// problem.
// ===================================================================

/*
======================================================================
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
