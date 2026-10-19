# Two Pointers

## Mental model

Here, a **pointer** simply means a position being tracked. Two pointers avoid
restarting a scan by moving two meaningful positions in a coordinated way.

Common forms:

- start at both ends and move inward;
- move through two sorted inputs together;
- use one position to read and another to write;
- move outward from the center of a possible palindrome; or
- move through a linked list at different speeds.

For every move, explain why the position being left behind cannot be needed
again.

## Finding a pair in sorted values

`PairSumSorted` starts at the smallest and largest values.

**Example:** Find two values totaling `9` in `[1, 2, 4, 7, 11]`.

1. `1 + 11 = 12`, which is too large. Moving the left position would only make
   the sum larger, so move the right position left.
2. `1 + 7 = 8`, which is too small. Moving the right position would only make
   the sum smaller, so move the left position right.
3. `2 + 7 = 9`, so return those positions.

**Invariant (fact that remains true):** No pair using a position already passed
by `left` or `right` can meet the target. Sorted order is what makes each
discarded position safe.

## Merging sorted arrays

For two sorted inputs, track the smallest unread value in each one. Copy the
smaller value to the output and advance only its position.

**Example:** Merge `[1, 4]` and `[2, 3]`. Compare `1` with `2` and write `1`;
then compare `4` with `2` and write `2`; then write `3`; finally copy the
remaining `4`. The result is `[1, 2, 3, 4]`.

If the first array has unused positions at its end, fill it backward. Start at
the largest real value in each input and write the larger one into the final
unused position. Filling forward could overwrite a value in the first array
before it has been read.

## Removing values inside the same array

Use a **read** position to inspect every value and a **write** position to mark
where the next retained value belongs.

For sorted values, duplicates are next to one another. Keep the first value,
then copy a later value only when it differs from the last retained value.

**Example:** In `[1, 1, 2, 2, 3]`, keep the first `1`, skip the second `1`, keep
the first `2`, skip the second `2`, and keep `3`. The first three array
positions become `[1, 2, 3]`, and the write position ends at `3`.

The same method removes values from unsorted input: copy a value only when it
passes the problem's keep/remove test.

**Invariant (fact that remains true):** `values[:write]`, the part before the
write position, contains exactly the retained values seen so far in their
original order. Because the write position never passes the read position,
writing cannot destroy unread input.

## Expanding around palindrome centers

A palindrome reads the same forward and backward. Its center is either one
character, for an odd length, or the gap between two characters, for an even
length. Try every character and every gap as a center, then move left and right
while the characters match.

**Example:** In `"babad"`, start at the center character `a`. The neighboring
characters are both `b`, so the matching range grows to `"bab"`. The next pair
would go beyond the left side and stop.

The characters between the two positions must already form a palindrome before
another matching pair is added. Testing only the center of the complete string
can miss a palindrome centered elsewhere.

## Work and memory

| Use | Work | Extra memory |
| --- | --- | --- |
| Inward pair search | `O(n)` | `O(1)` |
| Merge into a new output array | `O(m+n)` | `O(m+n)` |
| Merge backward into existing unused space | `O(m+n)` | `O(1)` |
| Remove values or repeated copies in the same array | `O(n)` | `O(1)` |
| Expand around every palindrome center | Between `O(n)` and `O(n^2)` | `O(1)` |

Palindrome expansion stops quickly in the best case. In the worst case, many
centers expand across much of the string.

## Common mistakes

- The inward pair-sum method requires sorted input.
- Confirm whether the problem asks for positions or values.
- Negative values are fine when the data is sorted; sorted order, not whether
  values are positive, justifies the pointer movement.
- Check that adding two values cannot exceed Go's `int` limit for the problem.
- When merging into unused space at the end of the first array, use the supplied
  populated lengths. A zero may be real input rather than empty padding.
- Advance the write position only after retaining a value; otherwise gaps appear
  in the retained portion.
- For palindromes, test both character centers and gap centers so even lengths
  are included.

## Related problems

- [LeetCode 9: Palindrome Number](../../../problems/solved_easy_leetcode_0009_palindrome_number.go)
  compares digits from both ends.
- [LeetCode 5: Longest Palindromic Substring](../../../problems/solved_medium_leetcode_0005_longest_palindromic_substring.go)
  expands from every possible center.
- [LeetCode 26: Remove Duplicates from Sorted Array](../../../problems/solved_easy_leetcode_0026_remove_duplicates_from_sorted_array.go)
  keeps one copy of each value with read and write positions.
- [LeetCode 27: Remove Element](../../../problems/solved_easy_leetcode_0027_remove_element.go)
  keeps every value that differs from the requested removal value.
