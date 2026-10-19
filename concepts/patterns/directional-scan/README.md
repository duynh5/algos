# Directional Scans

## Mental model

Sometimes an item's meaning depends on the item beside it. Scan in the direction
that makes that neighbor already known instead of repeatedly looking ahead.

For a rule based on the right neighbor, moving from right to left means the
neighbor was processed in the previous step. For a rule based on the left
neighbor, moving from left to right is usually simpler.

## Worked example

Roman numeral `IV` means `4`: `I` is smaller than the `V` immediately to its
right, so subtract `1` instead of adding it.

Using numeric values `[1, 5]` from right to left:

1. Add `5`; the total is `5`, and remember `5` as the right neighbor.
2. `1` is smaller than `5`, so subtract it; the total is `4`.

`SumWithSmallerLeftValuesSubtracted` implements this local rule for a slice of
nonnegative integers. It starts the remembered right value at zero so the
rightmost nonnegative value is added.

## Invariant (fact that remains true)

Before processing a position, every value to its right has already been added or
subtracted according to the rule, and `rightValue` is the value immediately to
its right. Applying the rule once extends the correct result one position to the
left.

## Work and memory

- Work: `O(n)`, because every value is visited once.
- Extra memory: `O(1)`, because only the total and the neighboring value are
  retained.

## Common mistakes

- A neighboring-value rule must come from the problem. Comparing adjacent
  values does not make the rule correct for every sequence.
- This implementation assumes nonnegative values. When values may be negative,
  add the rightmost value separately instead of comparing it with an initial
  zero.
- Choose the direction from the neighbor you need to know.
- Decide what empty input means. This implementation returns zero.
- Update the remembered neighbor after every value, including a value that was
  subtracted.

## Related problems

- [LeetCode 13: Roman to Integer](../../../problems/solved_easy_leetcode_0013_roman_to_integer.go)
  subtracts a smaller symbol only when it appears immediately before the larger
  symbol that changes its meaning.
