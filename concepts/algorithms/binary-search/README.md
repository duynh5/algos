# Binary Search and Bounds

## Mental model

Binary search repeatedly tests the middle possible answer and discards one half.
It is safe only when all "no" answers form one consecutive group and all "yes"
answers form another, so the result changes at most once.

The most familiar example is finding a value in sorted data. The same idea can
find the smallest numeric answer for which a yes/no test succeeds.

## Recognize it

Use binary search when:

- the values are sorted;
- the problem asks for the first or last position meeting a condition; or
- possible numeric answers have a single boundary between failure and success.

Do not use it merely because the answer is a number. You must be able to explain
why testing the middle rules out an entire half.

## Finding lower and upper bounds

`LowerBound` finds the first position whose value is at least the target.
`UpperBound` finds the first position whose value is greater than the target.
`Search` uses `LowerBound`, then returns the first matching position or `-1` when
the target is absent.

Both functions keep a range written as `[left, right)`. This notation includes
`left` but excludes `right`. The range therefore contains `right - left`
possible positions and is empty when `left == right`.

**Example:** Find the first value at least `3` in `[1, 3, 3, 5]`.

1. Start with possible positions `[0, 4)`.
2. Test position `2`, whose value is `3`. The answer may be position `2` or
   earlier, so keep `[0, 2)`.
3. Test position `1`, whose value is `3`. Again keep that position and anything
   earlier, producing `[0, 1)`.
4. Test position `0`, whose value is `1`. It is too small, so discard it.
5. The range is empty at position `1`, which is the first valid position.

**Invariant (fact that remains true):** Before every test, the first acceptable
position is somewhere in `[left, right)`. When the range becomes empty, `left`
is where the target could be inserted without breaking the sorted order.

## Sorted arrays that wrap around

An array such as `[4, 5, 6, 1, 2, 3]` was sorted and then shifted. The whole
array is no longer sorted, but at least one side of any tested middle position
is still sorted.

**Example:** To find `2`, first test `6`. The left portion `[4, 5, 6]` is sorted,
but `2` is not between `4` and `6`, so discard it and search `[1, 2, 3]`. The
next middle value is `2`.

When using a range that includes both endpoints, discard a tested middle with
`middle + 1` or `middle - 1`. Keeping the same middle can prevent the search
from finishing.

## Work and memory

- Time: `O(log n)`, because every test removes about half of the remaining
  positions.
- Extra space: `O(1)`, because the loop stores only a few positions.

## Common mistakes

- Ordinary lower and upper bounds require fully sorted input. Wrapped sorted
  input needs the separate rule described above.
- Decide whether a value equal to the target belongs in the kept left or right
  half; that choice determines whether you find the first or last boundary.
- Calculate the middle as `left + (right-left)/2`.
- A valid insertion position may equal `len(values)`, meaning the new value
  belongs after every existing value.
- In a wrapped-array search, include the sorted side's outer value when checking
  whether the target lies on that side.

## Related problems

- [LeetCode 35: Search Insert Position](../../../problems/solved_easy_leetcode_0035_search_insert_position.go)
  finds the first position whose value is at least the target.
- [LeetCode 33: Search in Rotated Sorted Array](../../../problems/solved_medium_leetcode_0033_search_in_rotated_sorted_array.go)
  identifies the sorted side after each middle test.
