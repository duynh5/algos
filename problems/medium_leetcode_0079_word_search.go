//go:build ignore

// Difficulty: Medium
// Source: leetcode
// Problem ID: 79

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/duyserveacc/algos/internal/testrunner"
)

/*
===============================================================================
Word Search

Source: LeetCode
Problem ID: 79
Source URL: https://leetcode.com/problems/word-search/description/

Given an m x n grid of characters board and a string word, return true if word exists in the grid.

The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once.

Example 1:

Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
Output: true

Example 2:

Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
Output: true

Example 3:

Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
Output: false

Constraints:

- m == board.length

- n = board[i].length

- 1 <= m, n <= 6

- 1 <= word.length <= 15

- board and word consists of only lowercase and uppercase English letters.

Follow up: Could you use search pruning to make your solution faster with a larger board?
===============================================================================
*/

// Problem-specific support.
//
// Keep shared judge types and parsing or formatting helpers here. Name adapters
// after the problem. Put setup used only by the current attempt at the top of
// solveCurrentAttempt below.
//
// This intentionally unsolved practice archive follows the workbench template.

func parseWordSearchInput(input string) ([][]byte, string) {
	var args []json.RawMessage
	if err := json.Unmarshal([]byte(input), &args); err != nil {
		panic(fmt.Errorf("parse word-search input: %w", err))
	}
	if len(args) != 2 {
		panic(fmt.Sprintf("parse word-search input: got %d arguments, want 2", len(args)))
	}
	board, argumentError0 := decodeCharacterMatrix(args[0])
	if argumentError0 != nil {
		panic(fmt.Errorf("parse word-search argument board: %w", argumentError0))
	}
	var word string
	argumentError1 := json.Unmarshal(args[1], &word)
	if argumentError1 != nil {
		panic(fmt.Errorf("parse word-search argument word: %w", argumentError1))
	}
	return board, word
}

func decodeCharacterMatrix(raw json.RawMessage) ([][]byte, error) {
	var rows [][]string
	if err := json.Unmarshal(raw, &rows); err != nil {
		return nil, err
	}
	result := make([][]byte, len(rows))
	for row := range rows {
		result[row] = make([]byte, len(rows[row]))
		for column, value := range rows[row] {
			if len(value) != 1 {
				return nil, fmt.Errorf("cell (%d,%d) is not one character", row, column)
			}
			result[row][column] = value[0]
		}
	}
	return result, nil
}

func maximumWordSearchCase() testrunner.Case {
	board := make([][]string, 6)
	for row := range board {
		board[row] = make([]string, 6)
		for column := range board[row] {
			board[row][column] = "Z"
		}
	}
	board[0] = []string{"a", "b", "c", "d", "e", "f"}
	board[1] = []string{"l", "k", "j", "i", "h", "g"}
	board[2][0] = "m"
	board[2][1] = "n"
	board[2][2] = "o"
	input, err := json.Marshal([]any{board, "abcdefghijklmno"})
	if err != nil {
		panic(fmt.Errorf("build maximum word-search input: %w", err))
	}
	return testrunner.Case{
		Name:  "Boundary: Maximum board and word follow a turning path",
		Input: string(input),
		Want:  `true`,
	}
}

// sampleCases contains the examples that every registered solution must pass.
// Inputs encode the judge arguments as a positional JSON array.
var sampleCases = []testrunner.Case{
	{
		Name:  "Official example 1",
		Input: `[[["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]],"ABCCED"]`,
		Want:  `true`,
	},
	{
		Name:  "Official example 2",
		Input: `[[["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]],"SEE"]`,
		Want:  `true`,
	},
	{
		Name:  "Official example 3",
		Input: `[[["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]],"ABCB"]`,
		Want:  `false`,
	},
	{
		Name:  "Additional: One cell found",
		Input: `[[["A"]],"A"]`,
		Want:  `true`,
	},
	{
		Name:  "Additional: One cell absent",
		Input: `[[["A"]],"B"]`,
		Want:  `false`,
	},
	{
		Name:  "Additional: A cell cannot be reused",
		Input: `[[["A","A"]],"AAA"]`,
		Want:  `false`,
	},
	{
		Name:  "Additional: A word can traverse a row in reverse",
		Input: `[[["A","B","C","D"]],"DCBA"]`,
		Want:  `true`,
	},
	{
		Name:  "Adversarial: Diagonal cells are not adjacent",
		Input: `[[["A","B"],["C","D"]],"AD"]`,
		Want:  `false`,
	},
	{
		Name:  "Adversarial: Letter matching is case-sensitive",
		Input: `[[["a","A"]],"aa"]`,
		Want:  `false`,
	},
	{
		Name:  "Adversarial: Failed branch releases cells for another direction",
		Input: `[[["A","A","B"],["A","C","C"]],"AABC"]`,
		Want:  `true`,
	},
	maximumWordSearchCase(),
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
	board, word := parseWordSearchInput(input)
	_ = board
	_ = word

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
