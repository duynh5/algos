# Binary Search and Bounds

Binary search applies when a search space has a monotonic decision: after some
boundary, a predicate remains true. Searching a sorted slice is the most common
example, but the same reasoning supports binary search over an answer range.

## Invariant

`LowerBound` maintains a half-open interval `[left, right)` containing the first
index whose value could be at least the target. When the interval is empty,
`left` is the insertion point.

`UpperBound` uses the same interval but finds the first value strictly greater
than the target.

## Rotated arrays

A distinct-value sorted array that has been rotated is not globally sorted,
but every non-empty search interval has at least one sorted half. After checking
the midpoint for equality, compare the interval endpoints to determine which
half is sorted. Keep that half only when its inclusive outer endpoint and
exclusive midpoint bounds contain the target; otherwise keep the other half.

With a closed `[left, right]` interval, every nonmatching step must exclude the
midpoint with `middle+1` or `middle-1`. If the interval becomes empty, the
target is absent.

## Complexity

- Time: O(log n)
- Extra space: O(1)

## Pitfalls

- Ordinary bounds require globally sorted input; rotated input needs the
  sorted-half invariant described above.
- Be deliberate about whether equality belongs on the left or right side.
- Prefer `left + (right-left)/2` to make the midpoint rule explicit and safe.
- A valid insertion point can equal `len(values)`.
- In a rotated search, include the sorted half's outer endpoint when testing
  target membership; excluding it can discard a valid boundary match.

## Related problems

- [LeetCode 35: Search Insert Position](../../../problems/solved_easy_leetcode_0035_search_insert_position.go)
  applies the lower-bound invariant directly.
- [LeetCode 33: Search in Rotated Sorted Array](../../../problems/solved_medium_leetcode_0033_search_in_rotated_sorted_array.go)
  identifies the sorted half of each active interval.
