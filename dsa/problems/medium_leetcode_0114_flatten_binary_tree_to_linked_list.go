//go:build ignore

// Difficulty: Medium
// Source: leetcode
// Problem ID: 114

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/duyserveacc/algos/dsa/internal/testrunner"
)

/*
======================================================================
Flatten Binary Tree to Linked List

Source: LeetCode
Problem ID: 114
Source URL:
../../REFERENCE_LINKS.md#leetcode-source-0114

Given the root of a binary tree, flatten the tree into a "linked
list":

- The "linked list" should use the same TreeNode class where the right
  child pointer points to the next node in the list and the left child
  pointer is always null.

- The "linked list" should be in the same order as a pre-order
  traversal of the binary tree.

Example 1:

Input: root = [1,2,5,3,4,null,6]
Output: [1,null,2,null,3,null,4,null,5,null,6]

Example 2:

Input: root = []
Output: []

Example 3:

Input: root = [0]
Output: [0]

Constraints:

- The number of nodes in the tree is in the range [0, 2000].

- -100 <= Node.val <= 100

Follow up: Can you flatten the tree in-place (with O(1) extra space)?
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

type TreeNode struct {
	Val               int
	Left, Right, Next *TreeNode
}

func parseFlattenBinaryTreeToLinkedListInput(input string) *TreeNode {
	var args []json.RawMessage
	if err := json.Unmarshal([]byte(input), &args); err != nil {
		panic(fmt.Errorf("parse flatten-binary-tree-to-linked-list "+
			"input: %w", err))
	}
	if len(args) != 1 {
		panic(fmt.Sprintf("parse flatten-binary-tree-to-linked-list "+
			"input: got %d arguments, want 1", len(args)))
	}
	root, argumentError0 := decodeTreeNode(args[0])
	if argumentError0 != nil {
		panic(fmt.Errorf("parse flatten-binary-tree-to-linked-list "+
			"argument root: %w", argumentError0))
	}
	return root
}

func decodeTreeNode(raw json.RawMessage) (*TreeNode, error) {
	var values []*int
	if err := json.Unmarshal(raw, &values); err != nil ||
		len(values) == 0 {
		return nil, err
	}
	if values[0] == nil {
		return nil, nil
	}
	root := &TreeNode{Val: *values[0]}
	queue := []*TreeNode{root}
	next := 1
	for len(queue) > 0 && next < len(values) {
		node := queue[0]
		queue = queue[1:]
		if values[next] != nil {
			node.Left = &TreeNode{Val: *values[next]}
			queue = append(queue, node.Left)
		}
		next++
		if next < len(values) {
			if values[next] != nil {
				node.Right = &TreeNode{Val: *values[next]}
				queue = append(queue, node.Right)
			}
			next++
		}
	}
	return root, nil
}

func formatFlattenBinaryTreeToLinkedListOutput(
	root *TreeNode,
) string {
	values := make([]any, 0)
	for node := root; node != nil; node = node.Right {
		if len(values) == 4000 {
			panic("format flatten-binary-tree-to-linked-list " +
				"output: right chain exceeds 2000 nodes or " +
				"contains a cycle")
		}
		if node.Left != nil {
			panic("format flatten-binary-tree-to-linked-list " +
				"output: flattened node has a left child")
		}
		if len(values) > 0 {
			values = append(values, nil)
		}
		values = append(values, node.Val)
	}
	output, err := json.Marshal(values)
	if err != nil {
		panic(fmt.Errorf("format flatten-binary-tree-to-linked-list "+
			"output: %w", err))
	}
	return string(output)
}

func maximumFlattenBinaryTreeToLinkedListCase() testrunner.Case {
	inputTree := make([]any, 0, 3999)
	want := make([]any, 0, 3999)
	for index := 0; index < 2000; index++ {
		value := index%201 - 100
		inputTree = append(inputTree, value)
		want = append(want, value)
		if index+1 < 2000 {
			if index > 0 {
				inputTree = append(inputTree, nil)
			}
			want = append(want, nil)
		}
	}
	input, err := json.Marshal([]any{inputTree})
	if err != nil {
		panic(
			fmt.Errorf(
				"build maximum flatten-binary-tree input: %w",
				err,
			),
		)
	}
	output, err := json.Marshal(want)
	if err != nil {
		panic(
			fmt.Errorf(
				"build maximum flatten-binary-tree output: %w",
				err,
			),
		)
	}
	return testrunner.Case{
		Name: "Boundary: Maximum left chain becomes a right " +
			"chain",
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
		Input: `[[1,2,5,3,4,null,6]]`,
		Want:  `[1,null,2,null,3,null,4,null,5,null,6]`,
	},
	{
		Name:  "Official example 2",
		Input: `[[]]`,
		Want:  `[]`,
	},
	{
		Name:  "Official example 3",
		Input: `[[0]]`,
		Want:  `[0]`,
	},
	{
		Name:  "Additional: Only left children",
		Input: `[[1,2,null,3]]`,
		Want:  `[1,null,2,null,3]`,
	},
	{
		Name:  "Additional: Only right children",
		Input: `[[1,null,2,null,3]]`,
		Want:  `[1,null,2,null,3]`,
	},
	{
		Name:  "Adversarial: Right subtree follows the entire left subtree",
		Input: `[[1,2,3,null,4,5,6]]`,
		Want:  `[1,null,2,null,4,null,3,null,5,null,6]`,
	},
	{
		Name:  "Additional: Sparse mixed tree",
		Input: `[[1,null,2,3,4]]`,
		Want:  `[1,null,2,null,3,null,4]`,
	},
	{
		Name:  "Boundary: Duplicate extreme values",
		Input: `[[100,-100,100]]`,
		Want:  `[100,null,-100,null,100]`,
	},
	maximumFlattenBinaryTreeToLinkedListCase(),
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
	root := parseFlattenBinaryTreeToLinkedListInput(input)
	_ = root

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
