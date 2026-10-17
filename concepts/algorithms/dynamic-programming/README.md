# Dynamic Programming

Dynamic programming applies when a problem can be expressed as overlapping
subproblems whose answers combine through a recurrence. Define each state
precisely, establish its base cases, and choose an evaluation order in which
every dependency is already available.

Top-down memoization starts from the requested state and caches recursive
results. Bottom-up tabulation computes states in dependency order. When each
new state depends on only a fixed-size window of earlier states, rolling those
values can reduce auxiliary space without changing the recurrence.

## Invariant

For ordered one-or-two-step sequences, `ways(s)` is the number of sequences
whose moves sum to `s`. Every non-empty sequence ending at `s` has exactly one
last move:

- removing a final one-step move leaves a sequence counted by `ways(s-1)`;
- removing a final two-step move leaves a sequence counted by `ways(s-2)`.

The groups are disjoint and exhaustive, so
`ways(s) = ways(s-1) + ways(s-2)`. In the rolling implementation,
`twoBack` and `oneBack` hold those two dependencies before every iteration.

## Complexity

For this two-branch recurrence with `n` states and O(1) transition work:

- uncached recursion takes exponential time and O(n) call-stack space;
- memoization takes O(n) time and O(n) auxiliary space;
- full tabulation takes O(n) time and O(n) auxiliary space;
- a fixed rolling dependency window takes O(n) time and O(1) auxiliary space.

## Pitfalls

- A recurrence alone is not dynamic programming; repeated subproblems must be
  cached or evaluated once in a deliberate order.
- Define the state before writing the transition. Ambiguous state meaning often
  causes double-counting or missing cases.
- Use base cases that agree with the recurrence. Counting problems commonly
  use one way to construct an empty result, even when empty input is excluded
  by a particular problem.
- Do not compress storage until every dependency has been identified. Updating
  rolling variables in the wrong order can overwrite a value still needed.
- Include recursion-stack space when analyzing a memoized recursive solution.
- Check whether state values can overflow the chosen numeric type.

## Related problems

- [LeetCode 70: Climbing Stairs](../../../problems/solved_easy_leetcode_0070_climbing_stairs.go)
  uses a two-state recurrence that compresses to constant auxiliary space.
