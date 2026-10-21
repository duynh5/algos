//go:build ignore

// Difficulty: Easy
// Source: leetcode
// Problem ID: 21

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/duyserveacc/algos/internal/testrunner"
)

/*
======================================================================
Merge Two Sorted Lists

Source: LeetCode
Problem ID: 21
Source URL:
https://leetcode.com/problems/merge-two-sorted-lists/description/

You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists into one sorted list. The list should be made by
splicing together the nodes of the first two lists.

Return the head of the merged linked list.

Example 1:

Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]

Example 2:

Input: list1 = [], list2 = []
Output: []

Example 3:

Input: list1 = [], list2 = [0]
Output: [0]

Constraints:

- The number of nodes in both lists is in the range [0, 50].

- -100 <= Node.val <= 100

- Both list1 and list2 are sorted in non-decreasing order.
======================================================================
*/

// Problem-specific support.
//
// Keep shared judge types and parsing or formatting helpers here.
// Name adapters after the problem. Put setup used only by the current
// attempt at the top of solveCurrentAttempt below.
//
// This standalone archive follows the workbench structure.

type ListNode struct {
	Val  int
	Next *ListNode
}

func parseMergeTwoSortedListsInput(
	input string,
) (*ListNode, *ListNode) {
	var args []json.RawMessage
	if err := json.Unmarshal([]byte(input), &args); err != nil {
		panic(
			fmt.Errorf("parse merge-two-sorted-lists input: %w", err),
		)
	}
	if len(args) != 2 {
		panic(
			fmt.Sprintf("parse merge-two-sorted-lists input: got %d "+
				"arguments, want 2", len(args)),
		)
	}
	list1, argumentError0 := decodeListNode(args[0])
	if argumentError0 != nil {
		panic(
			fmt.Errorf(
				"parse merge-two-sorted-lists argument list1: "+
					"%w",
				argumentError0,
			),
		)
	}
	list2, argumentError1 := decodeListNode(args[1])
	if argumentError1 != nil {
		panic(
			fmt.Errorf(
				"parse merge-two-sorted-lists argument list2: "+
					"%w",
				argumentError1,
			),
		)
	}
	return list1, list2
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

func formatMergeTwoSortedListsOutput(head *ListNode) string {
	values := make([]int, 0)
	for node := head; node != nil; node = node.Next {
		values = append(values, node.Val)
	}
	output, err := json.Marshal(values)
	if err != nil {
		panic(
			fmt.Errorf(
				"format merge-two-sorted-lists output: %w",
				err,
			),
		)
	}
	return string(output)
}

func maximumMergeTwoSortedListsCase() testrunner.Case {
	list1 := make([]int, 25)
	list2 := make([]int, 25)
	want := make([]int, 50)
	for index := range 25 {
		list1[index] = index * 2
		list2[index] = index*2 + 1
	}
	for index := range 50 {
		want[index] = index
	}

	input, err := json.Marshal([]any{list1, list2})
	if err != nil {
		panic(
			fmt.Errorf("build maximum merge-two-sorted-lists input: "+
				"%w", err),
		)
	}
	output, err := json.Marshal(want)
	if err != nil {
		panic(
			fmt.Errorf(
				"build maximum merge-two-sorted-lists output: "+
					"%w",
				err,
			),
		)
	}
	return testrunner.Case{
		Name:  "Boundary: Maximum total node count",
		Input: string(input),
		Want:  string(output),
	}
}

// sampleCases contains the examples that every registered solution
// must pass. Inputs encode the judge arguments as a positional JSON
// array.
var sampleCases = []testrunner.Case{
	{
		Name:  "Official example 1",
		Input: `[[1,2,4],[1,3,4]]`,
		Want:  `[1,1,2,3,4,4]`,
	},
	{
		Name:  "Official example 2",
		Input: `[[],[]]`,
		Want:  `[]`,
	},
	{
		Name:  "Official example 3",
		Input: `[[],[0]]`,
		Want:  `[0]`,
	},
	{
		Name:  "Additional: Duplicate and negative values",
		Input: `[[-3,-1,2],[-3,0,2]]`,
		Want:  `[-3,-3,-1,0,2,2]`,
	},
	{
		Name:  "Additional: Only right list populated",
		Input: `[[],[-100,100]]`,
		Want:  `[-100,100]`,
	},
	{
		Name:  "Additional: Only left list populated",
		Input: `[[-100,0,100],[]]`,
		Want:  `[-100,0,100]`,
	},
	{
		Name:  "Additional: Left list entirely precedes right",
		Input: `[[1,2,3],[4,5,6]]`,
		Want:  `[1,2,3,4,5,6]`,
	},
	{
		Name:  "Additional: Right list entirely precedes left",
		Input: `[[4,5,6],[1,2,3]]`,
		Want:  `[1,2,3,4,5,6]`,
	},
	{
		Name:  "Additional: Values alternate between lists",
		Input: `[[1,3,5],[2,4,6]]`,
		Want:  `[1,2,3,4,5,6]`,
	},
	{
		Name:  "Additional: Equal values span both lists",
		Input: `[[1,1,3],[1,2,2]]`,
		Want:  `[1,1,1,2,2,3]`,
	},
	{
		Name:  "Additional: Unequal list lengths",
		Input: `[[1],[2,3,4,5]]`,
		Want:  `[1,2,3,4,5]`,
	},
	maximumMergeTwoSortedListsCase(),
}

// archiveReady records that the completed review passed validation.
const archiveReady = true

// Registered implementations run against the same cases.
var currentAttemptSolution = testrunner.Solution{
	Name: "current attempt",
	Run:  solveCurrentAttempt,
}

// agentSolutions contains distinct direct, improved, and preferred
// approaches.
var agentSolutions = []testrunner.Solution{
	{
		Name: "insertion sort by node splicing",
		Run:  solveByInsertionSort,
	},
	{Name: "recursive two-way merge", Run: solveByRecursiveMerge},
	{Name: "iterative dummy-tail merge", Run: solveByIterativeMerge},
}

var solutions = append(
	[]testrunner.Solution{currentAttemptSolution},
	agentSolutions...,
)

// ===================================================================
// CURRENT ATTEMPT
//
// This preserves the implementation completed before the agent
// review.
// ===================================================================

func solveCurrentAttempt(input string) string {
	list1, list2 := parseMergeTwoSortedListsInput(input)
	dummy := &ListNode{}
	tail := dummy

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}

	if list1 != nil {
		tail.Next = list1
	} else {
		tail.Next = list2
	}

	return formatMergeTwoSortedListsOutput(dummy.Next)
}

// ===================================================================
// AGENT SOLUTIONS
//
// These implementations remain separate from the preserved current
// attempt.
// ===================================================================

// solveByInsertionSort repeatedly inserts detached input nodes into a
// sorted result. It is intentionally direct and slower than a two-way
// merge.
func solveByInsertionSort(input string) string {
	list1, list2 := parseMergeTwoSortedListsInput(input)
	sorted := &ListNode{}

	for _, head := range []*ListNode{list1, list2} {
		for head != nil {
			node := head
			head = head.Next

			previous := sorted
			for previous.Next != nil && previous.Next.Val <= node.Val {
				previous = previous.Next
			}
			node.Next = previous.Next
			previous.Next = node
		}
	}

	return formatMergeTwoSortedListsOutput(sorted.Next)
}

// solveByRecursiveMerge selects the smaller head recursively.
func solveByRecursiveMerge(input string) string {
	list1, list2 := parseMergeTwoSortedListsInput(input)
	return formatMergeTwoSortedListsOutput(
		mergeSortedListsRecursively(list1, list2),
	)
}

func mergeSortedListsRecursively(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val <= list2.Val {
		list1.Next = mergeSortedListsRecursively(list1.Next, list2)
		return list1
	}
	list2.Next = mergeSortedListsRecursively(list1, list2.Next)
	return list2
}

// solveByIterativeMerge delegates to the preferred constant-space
// helper.
func solveByIterativeMerge(input string) string {
	list1, list2 := parseMergeTwoSortedListsInput(input)
	return formatMergeTwoSortedListsOutput(
		mergeTwoLists(list1, list2),
	)
}

// mergeTwoLists is the judge-facing preferred implementation.
func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}

	if list1 != nil {
		tail.Next = list1
	} else {
		tail.Next = list2
	}

	return dummy.Next
}

/*
======================================================================
AGENT REVIEW AND COMPLEXITY ANALYSIS

Current attempt:
- What is correct: It performs a standard two-way merge, reuses every
  input
node, handles either empty list, and attaches the remaining suffix
without
another traversal. For [1,2,4] and [1,3,4], tail receives 1 from
list1, 1
from list2, 2 from list1, and 3 from list2; list1 then points at [4]
and
  list2 at [4], so the final two selections produce [1,1,2,3,4,4].
- First failing assumption or bottleneck: None for the stated
  constraints. Its
O(m+n) worst case is optimal when the values interleave through both
lists.
- Relevant concept and invariant: dummy removes the head special case.
  Before
each comparison, dummy.Next through tail is sorted and final, while
list1
and list2 point to the smallest unmerged nodes. This is the same
algorithm
and invariant as the preferred solution; only the preferred version
extracts
it into the judge-facing mergeTwoLists helper. Compared with Solution
1 it
avoids rescanning the result, and compared with Solution 2 it avoids
stack
  growth.
- Time complexity: The merge loop is Theta(min(m,n)) in the best case
  and
O(m+n) in the average and worst cases. If either list is empty,
attaching the
other list is O(1). Parsing and formatting make this runner
Theta(m+n).
- Auxiliary-space complexity: O(1): one dummy node and a fixed set of
  pointers.
Parsing and output formatting in this local harness use O(m+n) space.

Solution 1 - Direct insertion sort by node splicing:
- Correctness: Each input node is detached and inserted after the last
  result
node whose value is no greater. The result is sorted after every
insertion,
  and each original node is inserted exactly once.
- Concrete example: With [1,2,4] and [1,3,4], processing list1 changes
  sorted
through [1], [1,2], [1,2,4]. The next 1 is inserted after the first 1,
producing [1,1,2,4]; then 3 goes between 2 and 4, and the final 4 goes
after
  the existing 4, producing [1,1,2,3,4,4].
- Time complexity: Theta((m+n)^2) in all cases because nodes from each
  already
  sorted input pass their previously inserted peers.
- Auxiliary-space complexity: O(1): nodes are reused and only a dummy
  node and
pointers are allocated. Harness parsing and formatting use O(m+n)
space.

Solution 2 - Improved recursive two-way merge:
- Correctness: The smaller head is the smallest remaining value, so it
  is a
valid next result node. Recursively merging the remaining suffixes
completes
  the sorted list; a nil list returns the other suffix unchanged.
- Concrete example: For [1,2,4] and [1,3,4], calls select list1's 1,
  list2's 1,
list1's 2, list2's 3, and list1's 4. The base call returns list2's
final [4],
  and the stack reconnects those chosen nodes as [1,1,2,3,4,4].
- Time complexity: Recursion makes Theta(min(m,n)) calls in the best
  case and
O(m+n) in the average and worst cases. The runner remains Theta(m+n).
- Auxiliary-space complexity: Theta(min(m,n)) best-case and O(m+n)
  worst-case
call stack, with no new result nodes. Harness parsing and formatting
also use
  O(m+n) space.

Solution 3 - Preferred iterative dummy-tail merge:
- Correctness: At each iteration the smaller head is the smallest
  unmerged
value. Linking it after tail preserves sorted order and advancing only
its
source preserves every other node. Once either list is empty, its
counterpart
  is already a sorted suffix and can be attached whole.
- Concrete example: For [1,2,4] and [1,3,4], (list1,list2,tail)
  advances through
(1,1,dummy), (2,1,1), (2,3,1), (4,3,2), and (4,4,3). It selects
list1's 4,
  then attaches list2's [4], giving [1,1,2,3,4,4].
- Time complexity: The merge loop is Theta(min(m,n)) in the best case
  and
  O(m+n) in the average and worst cases; the runner is Theta(m+n).
- Auxiliary-space complexity: O(1): one dummy node and fixed pointer
  state.
  Harness parsing and formatting use O(m+n) space.

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
