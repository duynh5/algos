# Directional Scans

A directional scan is useful when the contribution of the current item depends
on a nearby item. Choose the direction that makes that dependency part of the
state already processed instead of repeatedly looking ahead or rescanning.

`SumWithSmallerLeftValuesSubtracted` demonstrates the form used by ordered
symbol systems: scan from right to left, subtract a value when it is smaller
than its right neighbor, and add it otherwise.

## Invariant

Before each iteration, `total` is the signed sum of the suffix to the right of
the current index, and `rightValue` is the first value in that suffix. Applying
the local rule to the current value extends the invariant by one position.

## Complexity

- Time: Theta(n), because every value is visited once.
- Extra space: O(1), because only the total and neighboring value are retained.

## Pitfalls

- The local rule must be justified by the input's grammar or ordering contract.
  It is not valid for arbitrary sequences just because two adjacent values can
  be compared.
- Pick the scan direction from the dependency. A right-neighbor rule is usually
  simpler from right to left; a left-neighbor rule is usually simpler forward.
- Define the empty-input identity deliberately. This example returns zero.
- Update the remembered neighbor on every iteration, including subtraction.

## Related problems

- [LeetCode 13: Roman to Integer](../../../problems/solved_easy_leetcode_0013_roman_to_integer.go)
  uses this pattern because canonical subtractive notation places a smaller
  symbol immediately before the larger symbol that changes its sign.
