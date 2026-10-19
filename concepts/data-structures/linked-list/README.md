# Linked Lists

## Mental model

A singly linked list is a chain of nodes. Each node stores one value and a link
to the next node. The first node is the **head**; a `nil` next link marks the
end.

Unlike an array, a linked list cannot jump directly to position `i`. Reaching a
position requires following links from the head. In exchange, code can reconnect
known nodes without shifting every later value.

## Operations in this package

- `FromSlice` builds a list in the slice's order.
- `ToSlice` follows links from the head and collects the values.
- `Reverse` changes every next link so the chain points in the opposite
  direction.

## Worked example: reversing a list

Reverse `1 -> 2 -> 3 -> nil` while keeping three positions:

1. `previous` is `nil` and `current` is node `1`. Save node `2`, point `1` to
   `nil`, then advance.
2. `previous` is node `1` and `current` is node `2`. Save node `3`, point `2` to
   `1`, then advance.
3. `previous` is node `2` and `current` is node `3`. Save `nil`, point `3` to
   `2`, then advance.
4. `current` is `nil`, so `previous` is the new head of
   `3 -> 2 -> 1 -> nil`.

Saving the next node before changing the link is essential; otherwise the
unvisited remainder of the list is lost.

## Invariants (facts that remain true)

- Following `Next` from the head visits nodes in list order and eventually
  reaches `nil` when the list contains no loop.
- While building a list with a `tail` position, `tail` always names the final
  node. A new node can therefore be attached without walking from the head.
- During reversal, `previous` is the already reversed portion and `current` is
  the first node not yet reversed.

## Common techniques

- A **dummy head** is a temporary node placed before the real first node. It
  avoids separate code for inserting or replacing the head.
- Two positions can advance independently when input lists have different
  lengths.
- Fast and slow positions can detect a loop or locate the middle.
- For numbers stored with the lowest digit first, add matching nodes and carry
  `sum / 10` to the next position; the current output digit is `sum % 10`.

## Work and memory

| Operation | Work | Extra memory |
| --- | --- | --- |
| Visit or search the entire list | `O(n)` | `O(1)` |
| Insert after a node already in hand | `O(1)` | `O(1)` |
| Reverse with a loop | `O(n)` | `O(1)` |
| Copy all values to or from a slice | `O(n)` | `O(n)` |

These memory amounts do not count nodes required for a returned result. A
version that uses nested function calls can also keep as many as `n` unfinished
calls in memory.

## Common mistakes

- Save `node.Next` before replacing it during reversal.
- Make loop conditions handle an empty list and a one-node list.
- Keep the original head when it is still needed; advancing the only reference
  can make earlier nodes unreachable.
- Following links one by one never ends if the list links back to an earlier
  node.
- Prefer a loop when a list may be too long for deeply nested function calls.

## Related problems

- [LeetCode 2: Add Two Numbers](../../../problems/solved_easy_leetcode_0002_add_two_numbers.go)
- [LeetCode 21: Merge Two Sorted Lists](../../../problems/solved_easy_leetcode_0021_merge_two_sorted_lists.go)
