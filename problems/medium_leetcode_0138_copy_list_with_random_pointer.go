//go:build ignore

// Difficulty: Medium
// Source: leetcode
// Problem ID: 138

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/duyserveacc/algos/internal/testrunner"
)

/*
===============================================================================
Copy List with Random Pointer

Source: LeetCode
Problem ID: 138
Source URL: https://leetcode.com/problems/copy-list-with-random-pointer/description/

A linked list of length n is given such that each node contains an additional random pointer, which could point to any node in the list, or null.

Construct a deep copy of the list. The deep copy should consist of exactly n brand new nodes, where each new node has its value set to the value of its corresponding original node. Both the next and random pointer of the new nodes should point to new nodes in the copied list such that the pointers in the original list and copied list represent the same list state. None of the pointers in the new list should point to nodes in the original list.

For example, if there are two nodes X and Y in the original list, where X.random --> Y, then for the corresponding two nodes x and y in the copied list, x.random --> y.

Return the head of the copied linked list.

The linked list is represented in the input/output as a list of n nodes. Each node is represented as a pair of [val, random_index] where:

- val: an integer representing Node.val

- random_index: the index of the node (range from 0 to n-1) that the random pointer points to, or null if it does not point to any node.

Your code will only be given the head of the original linked list.

Example 1:

Input: head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
Output: [[7,null],[13,0],[11,4],[10,2],[1,0]]

Example 2:

Input: head = [[1,1],[2,1]]
Output: [[1,1],[2,1]]

Example 3:

Input: head = [[3,null],[3,0],[3,null]]
Output: [[3,null],[3,0],[3,null]]

Constraints:

- 0 <= n <= 1000

- -10^4 <= Node.val <= 10^4

- Node.random is null or is pointing to some node in the linked list.
===============================================================================
*/

// Problem-specific support.
//
// Keep shared judge types and parsing or formatting helpers here. Name adapters
// after the problem. Put setup used only by the current attempt at the top of
// solveCurrentAttempt below.
//
// This intentionally unsolved practice archive follows the workbench template.

type RandomNode struct {
	Val          int
	Next, Random *RandomNode
}

func parseCopyListWithRandomPointerInput(input string) *RandomNode {
	var args []json.RawMessage
	if err := json.Unmarshal([]byte(input), &args); err != nil {
		panic(fmt.Errorf("parse copy-list-with-random-pointer input: %w", err))
	}
	if len(args) != 1 {
		panic(fmt.Sprintf("parse copy-list-with-random-pointer input: got %d arguments, want 1", len(args)))
	}
	head, argumentError0 := decodeRandomList(args[0])
	if argumentError0 != nil {
		panic(fmt.Errorf("parse copy-list-with-random-pointer argument head: %w", argumentError0))
	}
	return head
}

func decodeRandomList(raw json.RawMessage) (*RandomNode, error) {
	var pairs [][]*int
	if err := json.Unmarshal(raw, &pairs); err != nil {
		return nil, err
	}
	if len(pairs) == 0 {
		return nil, nil
	}
	nodes := make([]*RandomNode, len(pairs))
	for index, pair := range pairs {
		if len(pair) != 2 || pair[0] == nil {
			return nil, fmt.Errorf("entry %d must contain [value, randomIndex]", index)
		}
		nodes[index] = &RandomNode{Val: *pair[0]}
		if index > 0 {
			nodes[index-1].Next = nodes[index]
		}
	}
	for index, pair := range pairs {
		if pair[1] == nil {
			continue
		}
		if *pair[1] < 0 || *pair[1] >= len(nodes) {
			return nil, fmt.Errorf("entry %d has out-of-range random index %d", index, *pair[1])
		}
		nodes[index].Random = nodes[*pair[1]]
	}
	return nodes[0], nil
}

func formatCopyListWithRandomPointerOutput(head *RandomNode) string {
	nodes := make([]*RandomNode, 0)
	indexByNode := make(map[*RandomNode]int)
	for node := head; node != nil; node = node.Next {
		if len(nodes) == 1000 {
			panic("format copy-list-with-random-pointer output: list exceeds 1000 nodes or contains a cycle")
		}
		indexByNode[node] = len(nodes)
		nodes = append(nodes, node)
	}
	pairs := make([][2]any, len(nodes))
	for index, node := range nodes {
		pairs[index][0] = node.Val
		if node.Random == nil {
			pairs[index][1] = nil
		} else if randomIndex, ok := indexByNode[node.Random]; ok {
			pairs[index][1] = randomIndex
		} else {
			panic("format copy-list-with-random-pointer output: random pointer leaves copied list")
		}
	}
	output, err := json.Marshal(pairs)
	if err != nil {
		panic(fmt.Errorf("format copy-list-with-random-pointer output: %w", err))
	}
	return string(output)
}

func maximumCopyListWithRandomPointerCase() testrunner.Case {
	pairs := make([][2]int, 1000)
	for index := range pairs {
		if index%2 == 0 {
			pairs[index][0] = -10000
		} else {
			pairs[index][0] = 10000
		}
		pairs[index][1] = len(pairs) - 1 - index
	}
	input, err := json.Marshal([]any{pairs})
	if err != nil {
		panic(fmt.Errorf("build maximum copy-random-list input: %w", err))
	}
	want, err := json.Marshal(pairs)
	if err != nil {
		panic(fmt.Errorf("build maximum copy-random-list output: %w", err))
	}
	return testrunner.Case{Name: "Boundary: Maximum list with mirrored random pointers", Input: string(input), Want: string(want)}
}

// sampleCases contains the examples that every registered solution must pass.
// Inputs encode the judge arguments as a positional JSON array.
var sampleCases = []testrunner.Case{
	{
		Name:  "Official example 1",
		Input: `[[[7,null],[13,0],[11,4],[10,2],[1,0]]]`,
		Want:  `[[7,null],[13,0],[11,4],[10,2],[1,0]]`,
	},
	{
		Name:  "Official example 2",
		Input: `[[[1,1],[2,1]]]`,
		Want:  `[[1,1],[2,1]]`,
	},
	{
		Name:  "Official example 3",
		Input: `[[[3,null],[3,0],[3,null]]]`,
		Want:  `[[3,null],[3,0],[3,null]]`,
	},
	{
		Name:  "Additional: Single self-reference",
		Input: `[[[1,0]]]`,
		Want:  `[[1,0]]`,
	},
	{
		Name:  "Additional: Single null reference",
		Input: `[[[-1,null]]]`,
		Want:  `[[-1,null]]`,
	},
	{
		Name:  "Minimum: Empty list",
		Input: `[[]]`,
		Want:  `[]`,
	},
	{
		Name:  "Adversarial: Forward and backward random references",
		Input: `[[[-10000,2],[0,0],[10000,1]]]`,
		Want:  `[[-10000,2],[0,0],[10000,1]]`,
	},
	{
		Name:  "Additional: Every random pointer is null",
		Input: `[[[1,null],[1,null],[1,null]]]`,
		Want:  `[[1,null],[1,null],[1,null]]`,
	},
	maximumCopyListWithRandomPointerCase(),
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
	head := parseCopyListWithRandomPointerInput(input)
	_ = head

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
