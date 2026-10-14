# Two Pointers

The two-pointers pattern tracks two meaningful positions instead of repeatedly
scanning a range. Common forms move inward from both ends, advance through two
sorted inputs, move at different speeds, or maintain the boundary of a valid
window. Pointers can also move symmetrically outward from a candidate center.

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

## Expanding around palindrome centers

Every palindrome has one center: a character for odd lengths or a gap between
characters for even lengths. Try all `n` character centers and `n-1` gap
centers. For each center, move `left` and `right` outward together while their
characters match.

The invariant is that the range between the pointers is already palindromic
before the next matching pair is included. Equal endpoints alone are not
enough: the enclosed range must also be a palindrome. Testing every center is
what provides coverage; expanding only from the whole string's midpoint can
miss a palindrome elsewhere.

Center expansion uses O(1) auxiliary space. Its best case is Theta(n) when
each center stops after constant work, and its worst case is Theta(n^2) when
many centers expand across much of the input.

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
- Alternating which endpoint moves after a mismatch does not prove that the
  discarded boundary is safe and can skip the longest palindrome.
- Check both character and gap centers; checking only character centers misses
  even-length palindromes.

## Related problems

- [LeetCode 9: Palindrome Number](../../../problems/solved_easy_leetcode_0009_palindrome_number.go)
  compares mirrored digits with inward-moving endpoints in its direct solution.
- [LeetCode 5: Longest Palindromic Substring](../../../problems/solved_medium_leetcode_0005_longest_palindromic_substring.go)
  expands mirrored byte pointers outward from every possible center.
