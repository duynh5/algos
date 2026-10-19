# Stack

## Mental model

A stack behaves like a stack of plates: add and remove only at the top. The most
recently added value is therefore the first one removed. This order is called
**last in, first out**.

Use a stack when the most recently opened or unfinished item must be handled
first, such as nested brackets, undo history, expression evaluation, or a
depth-first search written with a loop.

## Operations in this package

- `Push` adds a value to the top.
- `Pop` removes and returns the top value.
- `Peek` returns the top value without removing it.
- `Len` reports the number of stored values.
- `IsEmpty` reports whether no values are stored.

The implementation uses the final element of a Go slice as the top. Both
`Push` and `Pop` therefore work at the same end.

## Worked example: matching brackets

For `([])`:

1. Read `(` and push it.
2. Read `[` and push it. The stack now contains `(` followed by `[`.
3. Read `]`, pop `[`, and confirm that the pair matches.
4. Read `)`, pop `(`, and confirm that the pair matches.
5. The input is valid because every closing bracket matched and the stack is
   empty.

## Invariant (fact that remains true)

The stack contains exactly the opened items that have not yet been closed. The
top is the most recent unfinished item and must be completed first.

## Work and memory

| Operation | Work | Extra memory added by the operation |
| --- | --- | --- |
| `Push` | Usually `O(1)` | `O(1)`, aside from occasional slice growth |
| `Pop` | `O(1)` | `O(1)` |
| `Peek` | `O(1)` | `O(1)` |

Growing a full slice occasionally copies its stored values. Across many pushes,
those occasional copies still average to constant work per push.

The implementation clears a removed slice position before shortening the slice.
If the values contain pointers, this allows an otherwise unused object to be
freed instead of remaining referenced by hidden slice capacity.

## Common mistakes

- Decide how `Pop` and `Peek` report an empty stack. This implementation returns
  a zero value and `false`.
- Do not remove from the front of the slice. That no longer matches stack order
  and can keep a large underlying array in memory.
- Nested function calls already use a hidden stack that remembers where each
  unfinished call must return. An explicit stack is useful when that nesting
  may become too deep or when code needs direct control over the stored return
  positions.

## Related problems

- [LeetCode 20: Valid Parentheses](../../../problems/solved_easy_leetcode_0020_valid_parentheses.go)
