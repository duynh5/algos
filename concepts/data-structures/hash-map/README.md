# Hash Maps

## Mental model

A hash map is like a labeled collection of boxes. Give the map a label, called
a **key**, and it quickly returns the value stored under that key. Go provides
hash maps with its built-in `map` type.

Use a map when you repeatedly need to know whether something appeared earlier,
how often it appeared, which group it belongs to, or what information was saved
for it. Use a set—a map that stores only presence—when no additional value is
needed.

## Worked example: pair sum

`PairSumIndices` finds two different values whose sum equals a target.

For `[2, 7, 11, 15]` and target `9`:

1. At `2`, the needed partner is `9 - 2 = 7`. It has not appeared, so store
   value `2` with position `0`.
2. At `7`, the needed partner is `9 - 7 = 2`. The map reports that `2` appeared
   at position `0`.
3. Return positions `0` and `1`.

The function checks before storing the current value. This prevents one item
from being paired with itself while still allowing equal values at different
positions, such as `[3, 3]` for target `6`.

## Invariant (fact that remains true)

Before processing position `i`, the map contains a stored position for every
distinct value already visited in `values[0:i]`. It contains nothing from
position `i` or later.

## Work and memory

- Expected time: `O(n)`. The function visits each value once, and Go maps are
  designed to make each search and insertion take roughly constant time on
  average.
- Extra space: `O(n)` in the worst case, because the map may retain every
  visited value.

Go does not promise that an individual map operation always takes constant
time. Expected bounds are the practical model for ordinary algorithm problems.

## Common mistakes

- Decide whether repeated keys should keep the first position, the latest
  position, a count, or a collection of positions.
- Search before inserting when the current item must not match itself.
- Reading a missing key returns the value type's zero value. Use Go's
  `value, found := entries[key]` form when zero is also a valid stored value.
- Iterating over a Go map does not produce a guaranteed order. Sort the keys or
  maintain a separate ordered collection when output order matters.

## Related problems

- [LeetCode 1: Two Sum](../../../problems/solved_easy_leetcode_0001_two_sum.go)
