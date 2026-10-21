//go:build ignore

// Difficulty: Easy
// Source: leetcode
// Problem ID: 70

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/duyserveacc/algos/internal/testrunner"
)

/*
======================================================================
Climbing Stairs

Source: LeetCode
Problem ID: 70
Source URL: https://leetcode.com/problems/climbing-stairs/description/

You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways
can you climb to the top?

Example 1:

Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps

Example 2:

Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step

Constraints:

- 1 <= n <= 45
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

func parseClimbingStairsInput(input string) int {
	var args []json.RawMessage
	if err := json.Unmarshal([]byte(input), &args); err != nil {
		panic(fmt.Errorf("parse climbing-stairs input: %w", err))
	}
	if len(args) != 1 {
		panic(fmt.Sprintf("parse climbing-stairs input: got %d "+
			"arguments, want 1", len(args)))
	}
	var n int
	argumentError0 := json.Unmarshal(args[0], &n)
	if argumentError0 != nil {
		panic(
			fmt.Errorf(
				"parse climbing-stairs argument n: %w",
				argumentError0,
			),
		)
	}
	return n
}

func formatClimbingStairsOutput(result int) string {
	return testrunner.FormatJSONOutput(result)
}

// sampleCases contains the examples that every registered solution
// must pass. Inputs encode the judge arguments as a positional JSON
// array.
var sampleCases = []testrunner.Case{
	{
		Name:  "Official example 1",
		Input: `[2]`,
		Want:  `2`,
	},
	{
		Name:  "Official example 2",
		Input: `[3]`,
		Want:  `3`,
	},
	{
		Name:  "Additional: Minimum staircase",
		Input: `[1]`,
		Want:  `1`,
	},
	{
		Name:  "Additional: Small transition",
		Input: `[4]`,
		Want:  `5`,
	},
	{
		Name:  "Additional: Constraint maximum",
		Input: `[45]`,
		Want:  `1836311903`,
	},
	{
		Name:  "Additional: Next recurrence transition",
		Input: `[5]`,
		Want:  `8`,
	},
	{
		Name:  "Additional: Two-digit staircase",
		Input: `[10]`,
		Want:  `89`,
	},
	{
		Name:  "Additional: Mid-range staircase",
		Input: `[20]`,
		Want:  `10946`,
	},
	{
		Name:  "Additional: Large staircase",
		Input: `[35]`,
		Want:  `14930352`,
	},
	{
		Name:  "Boundary: Immediately below maximum",
		Input: `[44]`,
		Want:  `1134903170`,
	},
}

// archiveReady is true after the current attempt, three agent
// solutions, analysis, and regression coverage are complete and
// verified.
const archiveReady = true

// Registered implementations run against the same cases.
var currentAttemptSolution = testrunner.Solution{
	Name: "current attempt",
	Run:  solveCurrentAttempt,
}

var agentSolutions = []testrunner.Solution{
	{
		Name: "direct combinatorial counting",
		Run:  solveByCombinatorialCounting,
	},
	{
		Name: "improved memoized recurrence",
		Run:  solveByMemoizedRecurrence,
	},
	{
		Name: "preferred rolling dynamic programming",
		Run:  solveByRollingDynamicProgramming,
	},
}

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
	n := parseClimbingStairsInput(input)
	_ = n
	var computeWays func(int) int
	memo := make(map[int]int, n)
	memo[0] = 0
	memo[1] = 1
	memo[2] = 2

	computeWays = func(w int) int {
		if val, exists := memo[w]; exists {
			return val
		}
		memo[w] = computeWays(w-1) + computeWays(w-2)
		return memo[w]
	}

	// Implement your algorithm with the typed values above.
	return formatClimbingStairsOutput(computeWays(n))
}

// ===================================================================
// AGENT SOLUTIONS
//
// These functions remain separate from the preserved current attempt
// so each approach can be compared against the same cases.
// ===================================================================

func solveByCombinatorialCounting(input string) string {
	n := parseClimbingStairsInput(input)
	totalWays := int64(0)

	for twoStepMoves := 0; twoStepMoves <= n/2; twoStepMoves++ {
		oneStepMoves := n - 2*twoStepMoves
		totalMoves := oneStepMoves + twoStepMoves
		chosenMoves := twoStepMoves
		if oneStepMoves < chosenMoves {
			chosenMoves = oneStepMoves
		}

		arrangements := int64(1)
		for choice := 1; choice <= chosenMoves; choice++ {
			arrangements = arrangements * int64(
				totalMoves-chosenMoves+choice,
			) / int64(
				choice,
			)
		}
		totalWays += arrangements
	}

	return formatClimbingStairsOutput(int(totalWays))
}

func solveByMemoizedRecurrence(input string) string {
	n := parseClimbingStairsInput(input)
	memo := make([]int, n+1)
	var countWays func(int) int

	countWays = func(steps int) int {
		if steps <= 2 {
			return steps
		}
		if memo[steps] != 0 {
			return memo[steps]
		}
		memo[steps] = countWays(steps-1) + countWays(steps-2)
		return memo[steps]
	}

	return formatClimbingStairsOutput(countWays(n))
}

func solveByRollingDynamicProgramming(input string) string {
	n := parseClimbingStairsInput(input)
	if n <= 2 {
		return formatClimbingStairsOutput(n)
	}

	waysTwoStepsBack, waysOneStepBack := 1, 2
	for step := 3; step <= n; step++ {
		waysTwoStepsBack, waysOneStepBack =
			waysOneStepBack, waysTwoStepsBack+waysOneStepBack
	}

	return formatClimbingStairsOutput(waysOneStepBack)
}

/*
======================================================================
AGENT REVIEW AND COMPLEXITY ANALYSIS

Current attempt:
- Score: 93/100. The recurrence, memoization, parsing, and formatting
  are
correct, and the attempt returns the right answer for every allowed
input in
linear expected time. The seven-point deduction is for avoidable map
and
recursion overhead, linear auxiliary space when constant space is
available,
and a memo[0] value that would be misleading outside the stated
constraints.
- What is correct: Every valid climb to w ends with either a one-step
  move
appended to a climb of w-1 or a two-step move appended to a climb of
w-2.
  These groups are disjoint and exhaustive, so computeWays(w-1) +
computeWays(w-2) is the correct recurrence. The map preloads ways(1)=1
and
ways(2)=2, and every computed larger state is stored before returning.
A
repeated state therefore becomes a map lookup rather than another
recursive
  subtree. The attempt passes all supplied and regression cases.
- First failing assumption or bottleneck: There is no correctness
  failure for
1<=n<=45. The remaining bottleneck is representation: a map lookup and
a
recursive call frame are used for each state even though the states
are dense
integers and each transition needs only the preceding two results.
Also,
memo[0]=0 is never reached for valid input, but if ways(w) means the
number of
sequences summing to w, the empty sequence convention would make
ways(0)=1.
- Relevant concept and invariant: This is dynamic programming over a
  one-
  dimensional state; see ../CONCEPTS.md#dynamic-programming.
The state ways(w) means exactly the number of ordered move sequences
whose
  sum is w. Partitioning those sequences by their final move proves
ways(w) = ways(w-1) + ways(w-2). After computeWays(w) returns, memo[w]
  stores that value, so the same state is never expanded twice.
- Concrete example: For n=5, memo begins with 1:1 and 2:2.
  computeWays(5)
misses, then computeWays(4) misses, then computeWays(3) misses. State
3 reads
cached values 2 and 1 and stores 3. State 4 then reads cached 3 and 2
and
stores 5. Back at state 5, the second request for state 3 is a cache
hit, so
  state 5 stores and returns 5+3=8.
- Comparison with solution 1: Both count the same ordered move
  sequences. The
attempt partitions them by their final move and caches each
remaining-step
state, while combinatorial counting partitions them by the number of
two-step
moves and counts each ordering with a binomial coefficient. The
attempt is
linear expected time but uses linear space; solution 1 uses quadratic
time
  and constant auxiliary space.
- Comparison with solution 2: Both use the same recurrence and base
  cases, so
they are the same top-down memoized algorithm rather than distinct
methods.
The attempt uses a map with explicit cached base values; solution 2
uses a
dense slice, handles the bases in the function, and uses zero as the
sentinel
  for uncached larger states. The slice has lower constant overhead.
- Comparison with solution 3: Both use the same recurrence and compute
  each
state once. The attempt discovers states top-down and retains every
result in
a map plus recursion stack; rolling dynamic programming evaluates
bottom-up
  and retains only the two states required for the next step.
- Time complexity: Theta(1) for n<=2. For n>2, expected Theta(n): each
  state
from 3 through n is computed once and each Go map operation is
expected O(1).
There is no distinct input-dependent average case for a fixed n.
Parsing and
  formatting a bounded scalar are Theta(1).
- Auxiliary-space complexity: Theta(n) for the memo map and Theta(n)
  for the
  deepest recursion path. Together they remain Theta(n).

Solution 1 - Direct combinatorial counting:
- Correctness: Fix twoStepMoves=k. The remaining n-2k distance must
  use
oneStepMoves=j, so every sequence in this group has j+k total moves.
Choosing
which k of those positions contain a two-step move gives C(j+k,k)
distinct
sequences. Every valid sequence has exactly one value of k, so summing
these
disjoint groups for 0<=k<=floor(n/2) counts every climb exactly once.
- Concrete example: For n=4, k=0 gives j=4 and C(4,0)=1 sequence:
  1111. k=1
gives j=2 and C(3,1)=3 sequences: 211, 121, and 112. k=2 gives j=0 and
  C(2,2)=1 sequence: 22. The running total is 1, then 4, then 5.
- Comparison with current attempt: The attempt splits sequences by
  their last
move and memoizes each remaining distance in expected linear time and
linear
space. This solution groups complete sequences by their number of
two-step
  moves and counts each group in quadratic time and constant space.
- Time complexity: O(n^2) worst-case arithmetic operations: there are
  O(n)
values of k, and computing a binomial coefficient uses O(n)
multiplicative
steps. The best case is Theta(1) at n=1; across growing valid n, the
work is
quadratic. Parsing and formatting are Theta(1) for the bounded scalar.
- Auxiliary-space complexity: O(1); only counters and int64
  accumulators are
  retained. int64 keeps intermediate products safe for n<=45.

Solution 2 - Improved memoized recurrence:
- Correctness: The base states return ways(1)=1 and ways(2)=2. For
  every larger
state, the last move is either one step following a valid climb of
steps-1 or
two steps following a valid climb of steps-2. The recursive sum is
therefore
correct, and caching only reuses an already proved result without
changing
  it.
- Concrete example: For n=5, countWays(5) first descends through
  states 4 and
3. State 3 stores 2+1=3; state 4 stores 3+2=5. When state 5 next
   requests
  state 3, memo[3]=3 returns immediately, and memo[5] becomes 5+3=8.
- Comparison with current attempt: This is the same top-down memoized
  algorithm
as the attempt, not a distinct underlying method. The material
difference is
storage: the attempt uses a map with explicit base entries, while this
solution uses a dense slice and function base cases. Both compute each
state
  once, but the slice avoids hashing and per-entry map overhead.
- Time complexity: Theta(n) for n>2 because each state from 3 through
  n is
computed once and every cache access is O(1); n<=2 takes Theta(1).
There is no
  input-dependent average-case variation for a fixed n.
- Auxiliary-space complexity: Theta(n) for the memo slice plus
  Theta(n) for
  the recursion stack. These add to Theta(n), not Theta(2n).

Solution 3 - Preferred rolling dynamic programming:
- Correctness: Before each iteration for step s, waysTwoStepsBack
  stores
ways(s-2) and waysOneStepBack stores ways(s-1). Their sum is ways(s)
by the
final-move partition. The simultaneous assignment advances the
invariant to
  the next step, so after processing n, waysOneStepBack is ways(n).
- Concrete example: For n=5, the pair starts as
  (ways(1),ways(2))=(1,2).
Processing step 3 changes it to (2,3); step 4 changes it to (3,5); and
step 5
  changes it to (5,8). The function returns the second value, 8.
- Comparison with current attempt: Both implement the same recurrence
  and avoid
repeated state computation. The attempt uses top-down recursion, a
map, and
Theta(n) auxiliary space. This solution evaluates states in increasing
order
and discards any state older than the preceding two, reducing
auxiliary space
  to O(1).
- Time complexity: Theta(1) for n<=2 and Theta(n) for n>2. Each
  additional
stair performs one constant-time iteration, with no distinct average
case.
- Auxiliary-space complexity: O(1); the loop retains two prior states
  and one
  loop index regardless of n.

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
