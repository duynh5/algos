//go:build ignore

// Difficulty: Easy
// Source: leetcode
// Problem ID: 121

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/duyserveacc/algos/dsa/internal/testrunner"
)

/*
======================================================================
Best Time to Buy and Sell Stock

Source: LeetCode
Problem ID: 121
Source URL:
../../REFERENCE_LINKS.md#leetcode-source-0121

You are given an array prices where prices[i] is the price of a given
stock on the i^th day.

You want to maximize your profit by choosing a single day to buy one
stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If
you cannot achieve any profit, return 0.

Example 1:

Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6),
profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because
you must buy before you sell.

Example 2:

Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transactions are done and the max profit
= 0.

Constraints:

- 1 <= prices.length <= 10^5

- 0 <= prices[i] <= 10^4
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

func parseBestTimeToBuyAndSellStockInput(input string) []int {
	var args []json.RawMessage
	if err := json.Unmarshal([]byte(input), &args); err != nil {
		panic(
			fmt.Errorf(
				"parse best-time-to-buy-and-sell-stock input: "+
					"%w",
				err,
			),
		)
	}
	if len(args) != 1 {
		panic(
			fmt.Sprintf(
				"parse best-time-to-buy-and-sell-stock input: "+
					"got %d arguments, want 1",
				len(args),
			),
		)
	}
	var prices []int
	argumentError0 := json.Unmarshal(args[0], &prices)
	if argumentError0 != nil {
		panic(fmt.Errorf("parse best-time-to-buy-and-sell-stock "+
			"argument prices: %w", argumentError0))
	}
	return prices
}

func maximumBestTimeToBuyAndSellStockCase() testrunner.Case {
	prices := make([]int, 100000)
	for index := range prices {
		prices[index] = 5000
	}
	prices[0] = 0
	prices[len(prices)-1] = 10000
	input, err := json.Marshal([]any{prices})
	if err != nil {
		panic(fmt.Errorf("build maximum stock case: %w", err))
	}
	return testrunner.Case{
		Name:  "Boundary: Maximum day count and profit",
		Input: string(input),
		Want:  `10000`,
	}
}

func formatBestTimeToBuyAndSellStockOutput(result int) string {
	return testrunner.FormatJSONOutput(result)
}

// sampleCases contains the examples that every registered solution
// must pass. Inputs encode the judge arguments as a positional JSON
// array.
var sampleCases = []testrunner.Case{
	{
		Name:  "Official example 1",
		Input: `[[7,1,5,3,6,4]]`,
		Want:  `5`,
	},
	{
		Name:  "Official example 2",
		Input: `[[7,6,4,3,1]]`,
		Want:  `0`,
	},
	{
		Name:  "Additional: Single price",
		Input: `[[5]]`,
		Want:  `0`,
	},
	{
		Name:  "Additional: Two rising prices",
		Input: `[[1,2]]`,
		Want:  `1`,
	},
	{
		Name:  "Additional: Equal prices",
		Input: `[[3,3,3]]`,
		Want:  `0`,
	},
	{
		Name:  "Additional: Better opportunity follows early profit",
		Input: `[[2,4,1,7]]`,
		Want:  `6`,
	},
	{
		Name:  "Additional: Repeated lowest price",
		Input: `[[5,1,1,5]]`,
		Want:  `4`,
	},
	{
		Name:  "Adversarial: Best buy occurs near end",
		Input: `[[8,7,6,1,10]]`,
		Want:  `9`,
	},
	{
		Name:  "Boundary: Minimum then maximum price",
		Input: `[[0,10000]]`,
		Want:  `10000`,
	},
	{
		Name:  "Boundary: Maximum precedes minimum",
		Input: `[[10000,0]]`,
		Want:  `0`,
	},
	maximumBestTimeToBuyAndSellStockCase(),
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
	prices := parseBestTimeToBuyAndSellStockInput(input)
	_ = prices

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
