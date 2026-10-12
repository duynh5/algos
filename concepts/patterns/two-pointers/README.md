# Two Pointers

The two-pointers pattern tracks two meaningful positions instead of repeatedly
scanning a range. Common forms move inward from both ends, advance through two
sorted inputs, move at different speeds, or maintain the boundary of a valid
window.

`PairSumSorted` demonstrates inward-moving pointers on ascending data.

## Merging sorted arrays

For two sorted inputs, each pointer identifies the smallest unread value. Move
the pointer whose value is emitted; when either input is exhausted, the other
input's unread suffix can be copied without further comparisons.

When the first array has spare capacity at its end, merge backward instead.
Pointers begin at both inputs' largest populated values, and a third pointer
begins at the last output slot. Writing the larger unread value preserves the
invariant that the suffix after the output pointer is complete and sorted. It
also avoids overwriting unread values in the first array.

## Invariant

For indices outside `[left, right]`, no pair can meet the target. When the sum
is too small, advancing `left` is the only move that can increase it. When the
sum is too large, decreasing `right` is the only move that can reduce it.

## Complexity

- Inward pair search: O(n) time and O(1) extra space.
- Forward merge with an output buffer: O(m+n) time and O(m+n) extra space.
- Backward merge into existing capacity: O(m+n) time and O(1) extra space.

## Pitfalls

- The inward-moving pair-sum version requires sorted input.
- Define whether indices or values must be returned.
- Negative values are valid when the data remains sorted. Do not reuse this
  movement rule when the problem's predicate breaks monotonicity.
- Ensure adding two input values cannot overflow `int` for the problem limits.
- A forward in-place merge can overwrite unread values in the destination.
  Merge backward when its unused capacity is at the end.
- Padding values are storage, not input data; use the populated lengths rather
  than testing whether a value is zero.

## Related problems

No solved problems are linked yet.
