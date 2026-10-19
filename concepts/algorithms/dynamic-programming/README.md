# Dynamic Programming

## Mental model

Start with a solution that tries several choices. If different choice sequences
keep reaching the same smaller problem, solve that smaller problem once and
reuse its answer.

For example, when climbing stairs in steps of one or two, every route to stair
`4` ends with either:

- one step from stair `3`; or
- two steps from stair `2`.

The number of routes to stair `4` is therefore the number reaching stair `3`
plus the number reaching stair `2`.

## Recognize it

Dynamic programming is useful when all three statements are true:

1. A large answer can be built from answers to smaller versions of the problem.
2. The simplest search calculates the same smaller answer more than once.
3. A limited set of facts completely identifies each smaller problem.

That identifying set of facts is commonly called the **state**. Define it in a
sentence before choosing an array or map to store answers.

## Two implementation directions

- **Top-down with caching:** Begin with the requested answer. Use nested function
  calls to ask for smaller answers, and save each answer the first time it is
  calculated.
- **Bottom-up:** Calculate the smallest answers first, then continue in an order
  that makes every required earlier answer available.

If a new answer needs only a fixed number of earlier answers, keep just those
values instead of an entire table. This is often called a rolling window.

## Worked example

`CountOrderedOneOrTwoStepWays` counts ordered sequences of one-step and two-step
moves:

| Stair | Explanation | Ways |
| --- | --- | --- |
| `0` | Do nothing; the empty sequence is one valid way. | `1` |
| `1` | Take one one-step move. | `1` |
| `2` | Use `1+1` or `2`. | `2` |
| `3` | Extend the ways for stairs `2` and `1`. | `2 + 1 = 3` |
| `4` | Extend the ways for stairs `3` and `2`. | `3 + 2 = 5` |

Before calculating a stair, `twoBack` stores the answer two stairs earlier and
`oneBack` stores the answer one stair earlier. Updating them together moves that
meaning forward by one stair.

## Why the formula is complete

Every non-empty route has exactly one final move. Removing a final one-step move
leaves a route to the previous stair; removing a final two-step move leaves a
route to the stair before that. A route cannot belong to both groups, and no
other final move is allowed, so adding the two counts includes every route once.

## Work and memory

For `n` stairs:

- Trying both choices without saving answers repeats work and grows roughly like
  `2^n`; the nested calls can also be `n` levels deep.
- Saving every answer uses `O(n)` work and `O(n)` extra memory.
- Calculating every answer in an array also uses `O(n)` work and `O(n)` extra
  memory.
- Keeping only the two required earlier answers uses `O(n)` work and `O(1)`
  extra memory.

## Common mistakes

- Writing a formula for smaller answers is not enough; repeated answers must be
  saved or calculated once in a deliberate order.
- Do not store facts that cannot affect later choices. Unnecessary facts can
  make the number of stored combinations too large.
- Choose starting answers that agree with the formula. Counting problems often
  treat an empty construction as one valid way to build nothing.
- Identify every earlier value before reducing memory. Updating rolling values
  in the wrong order can destroy a value that is still needed.
- Include the memory used by nested function calls when analyzing a top-down
  implementation.
- Check whether the counts can exceed the chosen integer type.

## Related problems

- [LeetCode 70: Climbing Stairs](../../../problems/solved_easy_leetcode_0070_climbing_stairs.go)
  needs only the two previous counts at each step.
