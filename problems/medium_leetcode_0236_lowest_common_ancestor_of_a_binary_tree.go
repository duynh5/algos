//go:build ignore

// Difficulty: Medium
// Source: leetcode
// Problem ID: 236

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/duyserveacc/algos/internal/testrunner"
)

/*
===============================================================================
Lowest Common Ancestor of a Binary Tree

Source: LeetCode
Problem ID: 236
Source URL: https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/description/

Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.

According to the definition of LCA on Wikipedia: "The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself)."

Example 1:

Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
Output: 3
Explanation: The LCA of nodes 5 and 1 is 3.

Example 2:

Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
Output: 5
Explanation: The LCA of nodes 5 and 4 is 5, since a node can be a descendant of itself according to the LCA definition.

Example 3:

Input: root = [1,2], p = 1, q = 2
Output: 1

Constraints:

- The number of nodes in the tree is in the range [2, 10^5].

- -10^9 <= Node.val <= 10^9

- All Node.val are unique.

- p != q

- p and q will exist in the tree.
===============================================================================
*/

// Problem-specific support.
//
// Keep shared judge types and parsing or formatting helpers here. Name adapters
// after the problem. Put setup used only by the current attempt at the top of
// solveCurrentAttempt below.
//
// This intentionally unsolved practice archive follows the workbench template.

type TreeNode struct {
	Val               int
	Left, Right, Next *TreeNode
}

func parseLowestCommonAncestorOfABinaryTreeInput(input string) (*TreeNode, int, int) {
	var args []json.RawMessage
	if err := json.Unmarshal([]byte(input), &args); err != nil {
		panic(fmt.Errorf("parse lowest-common-ancestor-of-a-binary-tree input: %w", err))
	}
	if len(args) != 3 {
		panic(fmt.Sprintf("parse lowest-common-ancestor-of-a-binary-tree input: got %d arguments, want 3", len(args)))
	}
	root, argumentError0 := decodeTreeNode(args[0])
	if argumentError0 != nil {
		panic(fmt.Errorf("parse lowest-common-ancestor-of-a-binary-tree argument root: %w", argumentError0))
	}
	var p int
	argumentError1 := json.Unmarshal(args[1], &p)
	if argumentError1 != nil {
		panic(fmt.Errorf("parse lowest-common-ancestor-of-a-binary-tree argument p: %w", argumentError1))
	}
	var q int
	argumentError2 := json.Unmarshal(args[2], &q)
	if argumentError2 != nil {
		panic(fmt.Errorf("parse lowest-common-ancestor-of-a-binary-tree argument q: %w", argumentError2))
	}
	return root, p, q
}

func decodeTreeNode(raw json.RawMessage) (*TreeNode, error) {
	var values []*int
	if err := json.Unmarshal(raw, &values); err != nil || len(values) == 0 {
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

// sampleCases contains the examples that every registered solution must pass.
// Inputs encode the judge arguments as a positional JSON array.
var sampleCases = []testrunner.Case{
	{
		Name:  "Official example 1",
		Input: `[[3,5,1,6,2,0,8,null,null,7,4],5,1]`,
		Want:  `3`,
	},
	{
		Name:  "Official example 2",
		Input: `[[3,5,1,6,2,0,8,null,null,7,4],5,4]`,
		Want:  `5`,
	},
	{
		Name:  "Official example 3",
		Input: `[[1,2],1,2]`,
		Want:  `1`,
	},
	{
		Name:  "Additional: Root is the ancestor",
		Input: `[[1,2,3],2,3]`,
		Want:  `1`,
	},
	{
		Name:  "Additional: One target is an ancestor",
		Input: `[[1,2,null,3],2,3]`,
		Want:  `2`,
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
	root, p, q := parseLowestCommonAncestorOfABinaryTreeInput(input)
	_ = root
	_ = p
	_ = q

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
