# Problem-Solving Patterns

Pattern pages focus on recognition signals and the invariant that turns a broad
class of problems into a repeatable approach.

Available:

- [Common-prefix scans](common-prefix-scan/)
- [Directional scans](directional-scan/)
- [Two pointers](two-pointers/)

## High-leverage pattern guide

These patterns cover many common interview problems, but they are a recognition
guide rather than a guarantee about any fixed percentage of LeetCode. Choose a
pattern because its invariant matches the problem, not merely because the input
contains an array, tree, or graph.

| Pattern | Recognition signals | Core invariant or tool |
| --- | --- | --- |
| Sliding window | A contiguous subarray or substring must satisfy, maximize, or minimize a condition. | Expand one boundary, update the window state, and shrink the other boundary while the state is invalid or while a better valid window is possible. |
| Subset generation and backtracking | The result must enumerate combinations, permutations, partitions, or choices. | Each state records the choices made so far; explore every valid next choice and undo mutable state when returning. Iterative breadth-first construction is an alternative when it makes each partial result explicit. |
| [Modified binary search](../algorithms/binary-search/) | The input is sorted, rotated, or searchable through a monotonic predicate. | Determine which half or answer range can still contain the target, then discard the other half without losing a valid answer. |
| Top K elements | Only the largest, smallest, closest, or most frequent `k` items matter. | Maintain a heap containing only the best `k` candidates seen so far; its root is the current candidate to evict. |
| Binary-tree DFS | A result depends on complete root-to-leaf paths, subtree results, or ancestor state. | Recursively or iteratively finish one branch before backtracking; define exactly what state enters a node and what result returns from it. |
| Topological sort | Directed edges express prerequisites or dependency order. | Repeatedly remove zero-indegree vertices and decrement their dependents, or use DFS states to order vertices and detect cycles. |
| Binary-tree BFS | The result is grouped by depth, needs the nearest match, or must process siblings together. | A queue holds the current frontier; process a captured level size before moving to the next depth. |
| [Two pointers](two-pointers/) | Sorted data, paired values, in-place partitioning, or two sequences allow monotonic progress. | Each pointer move must eliminate impossible candidates or finalize part of the result, so neither pointer needs to move backward. |

## Choosing among overlapping patterns

- Use a sliding window only when adding and removing boundary elements lets you
  maintain the condition incrementally. Negative values, non-monotonic rules,
  or arbitrary subsequences often break the usual movement rule.
- Use two pointers when a pointer move has a proof that skipped candidates
  cannot help. A window is a specialized two-pointer range with maintained
  aggregate state.
- Use binary-tree DFS for path or subtree dependencies; use binary-tree BFS
  when depth order or shortest unweighted distance matters.
- Use topological sort for a valid dependency ordering. Use ordinary graph BFS
  or DFS when the goal is reachability rather than prerequisite order.
- Account for the output size. Generating all subsets already requires
  `O(2^n)` results, and all permutations require `O(n!)`, regardless of whether
  traversal is depth-first or breadth-first.

See the complete [concept curriculum](../).
