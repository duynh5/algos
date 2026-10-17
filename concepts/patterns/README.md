# Problem-Solving Patterns

Pattern pages focus on recognition signals and the invariant that turns a broad
class of problems into a repeatable approach.

Available:

- [Common-prefix scans](common-prefix-scan/)
- [Directional scans](directional-scan/)
- [Two pointers](two-pointers/)

## Problem-solving method selection guide

Treat recognition signals as hypotheses, not keyword matches. An array does not
automatically imply two pointers, and the word "minimum" does not automatically
imply greedy. A method fits only when you can state the invariant that makes its
next step safe.

### Start with these questions

1. What must the output describe: one value, an index, a path, an ordering, or
   every valid construction?
2. Is the chosen data required to be contiguous, or may it be an arbitrary
   subsequence or subset?
3. Does the input already have useful order, or may it be sorted without losing
   information such as original indices or arrival order?
4. Can a candidate choice or answer be proved impossible and permanently
   discarded?
5. Does the answer for a state depend on answers to smaller states? If so,
   define exactly what information identifies a state.
6. Can the problem be modeled as vertices and edges, even when no graph is
   explicitly provided?
7. What time complexity do the constraints permit? Use the input limit to
   reject methods before implementing them.

### Recognition table

| Problem signal | Likely method | What must be true | Common warning |
| --- | --- | --- | --- |
| Exact lookup, membership, frequencies, complements, or grouping by a key | [Hash map or set](../data-structures/hash-map/) | Equality or a stable derived key identifies the information to reuse. | A map improves lookup time but does not preserve useful ordering unless that behavior is built separately. |
| Sorted values, two ends, pair selection, in-place compaction, or merging two ordered inputs | [Two pointers](two-pointers/) | Moving a pointer permanently eliminates candidates or finalizes an output position. | Sorting may destroy required original positions, and an unjustified pointer move can skip the answer. |
| A target in sorted data, a rotated sorted region, or the first/last value satisfying a monotonic condition | [Binary search](../algorithms/binary-search/) | The search space has an ordered boundary: once a predicate becomes true or false, it does not switch back. | Do not binary-search an answer merely because it is numeric; prove monotonicity and define the boundary precisely. |
| A contiguous subarray or substring must satisfy, maximize, or minimize a maintainable condition | Sliding window | Adding and removing boundary elements updates validity or cost incrementally, and moving a boundary cannot require reconsidering discarded starts. | Negative values and other non-monotonic effects often invalidate the usual shrinking rule. Arbitrary subsequences are not windows. |
| Many range sums or counts, or a subarray property expressible as the difference between two prefixes | Prefix sums or difference arrays | A prefix summary can be combined or subtracted to recover each queried range. | Prefix sums do not by themselves solve properties that cannot be composed from prefix state. Watch integer overflow. |
| Best count, cost, or feasibility is built from overlapping smaller decisions | [Dynamic programming](../algorithms/dynamic-programming/) | A complete state captures every fact from the past that can affect future choices, and transitions cover all legal final choices. | A recurrence without memoization or ordered tabulation may still be exponential. Too much state can make the table infeasible. |
| A locally best choice appears to lead to a global optimum | Greedy | An exchange argument, cut property, or staying-ahead invariant proves that committing now cannot make the optimum unreachable. | Examples passing is not proof. If an early choice may need to be undone, use dynamic programming, backtracking, or search instead. |
| Every combination, permutation, partition, placement, or valid construction must be generated | Backtracking | A partial state lists the remaining legal choices, and undoing a choice restores the exact prior state. | The output may already require exponential or factorial time. Add pruning only when it cannot remove a valid result. |
| Reachability, connected components, subtree aggregation, complete paths, or dependencies nested deeply | Depth-first search | A visited rule prevents unintended revisits, and the recursive return value or explicit stack state has a precise meaning. | Include recursion-stack space. Deep inputs may require an explicit stack in Go. |
| Minimum number of edges in an unweighted graph, level-order tree work, or nearest state reached by uniform-cost moves | Breadth-first search | Every edge has equal cost, so the first visit occurs at minimum distance. | Mark states visited when enqueuing, not later, to avoid duplicate frontier growth. Weighted edges require another method. |
| Prerequisites, dependency ordering, or cycle detection in directed dependencies | Topological sort | Directed edges mean "must occur before," and processing removes satisfied dependencies. | Use ordinary DFS or BFS for reachability. A topological order exists only for a directed acyclic graph. |
| Cheapest paths with nonnegative edge weights | Dijkstra's algorithm | Relaxing the smallest unsettled distance is safe because no later nonnegative path can improve it through a farther vertex. | Negative edges break the finalization rule; use an algorithm designed for them. |
| Repeated connectivity queries while edges are added, or grouping items into components | Disjoint-set union | The required history is component membership, not the actual path or traversal order. | It does not recover shortest paths and is not a general replacement for graph traversal. |
| Repeatedly need the smallest/largest item, the best `k` items, or the next scheduled event | Heap or priority queue | Only the current extreme must be removed efficiently; arbitrary ordered access is unnecessary. | A heap is not fully sorted. Removing or searching for a non-root item is not generally efficient. |
| Nested delimiters, undo/history behavior, expression evaluation, or the next greater/smaller element | Stack or monotonic stack | Only the most recent unresolved item matters; for a monotonic stack, each new value resolves a suffix of pending candidates. | If the oldest pending item must be handled first, use a queue. State clearly what remains unresolved in the stack. |
| Overlapping ranges must be combined, gaps found, or simultaneous events counted | Sort and merge intervals or sweep line | After sorting by a relevant endpoint or event position, earlier ranges can be finalized without later reconsideration. | Define whether touching endpoints overlap and process equal-position events in the correct order. |
| A linked structure may contain a cycle, or two sequences advance at related rates | Fast and slow pointers | Relative pointer speed creates a detectable meeting or establishes a midpoint while using constant space. | Prove that pointer dereferences are safe and that a meeting means the property you claim. |
| Repeated prefix queries, dictionary lookup by characters, or pruning words by shared prefixes | Trie | Prefix sharing is central and the alphabet/branching cost is acceptable. | For exact membership only, a hash set is usually simpler and often smaller. |

### Choosing among overlapping methods

- Sliding window versus prefix sums: use a window when boundary movement
  preserves a monotonic validity rule. Use prefix summaries when each range can
  be derived from two endpoints but window validity is not monotonic.
- Two pointers versus binary search: use two pointers when candidates can be
  consumed in one coordinated pass. Use binary search when each query can
  discard half of an ordered search space independently.
- Greedy versus dynamic programming: use greedy only when a proof permits an
  irreversible local choice. Use dynamic programming when several choices must
  remain represented because their future effects differ.
- Dynamic programming versus backtracking: use dynamic programming when many
  paths reach the same state and only an aggregate answer is needed. Use
  backtracking when individual constructions must be emitted or path-specific
  state cannot be merged.
- DFS versus BFS: use DFS for subtree results, exhaustive paths, and component
  traversal. Use BFS for levels and shortest paths when every move has equal
  cost.
- Heap versus sorting: sort when the complete order is needed or the data is a
  fixed batch. Use a heap when only repeated extremes or a bounded top `k` are
  needed, especially for streaming input.

### Let constraints eliminate bad choices

These are rough interview-scale checks, not guarantees; constants, memory,
language, and time limits still matter.

| Input scale | Methods worth considering |
| --- | --- |
| `n <= 10` | Factorial search may be feasible if pruning and per-result work are small. |
| `n <= 20` | `O(2^n)` subset search or bitmask dynamic programming may be feasible. |
| `n <= 100` | Cubic work may be possible; quadratic work is commonly reasonable. |
| `n <= 1,000` | Quadratic work may be possible but needs scrutiny. |
| `n <= 100,000` | Usually target `O(n log n)` or `O(n)`. |
| `n >= 1,000,000` | Usually target a small number of linear passes with controlled allocation. |

Also include unavoidable output work. Generating all `2^n` subsets or all
`n!` permutations cannot be made polynomial merely by changing the traversal.

### Before committing to a method

Write one sentence for each item:

- **State:** What does each variable, table entry, queue item, or stack entry
  mean?
- **Invariant:** What remains true before and after every step?
- **Progress:** Why does each step shrink the unresolved work or advance toward
  termination?
- **Completeness:** Why can no discarded candidate or omitted transition contain
  a valid answer?
- **Cost:** How many states or elements are processed, how often, and what extra
  memory is retained?

If one sentence cannot be made precise, the method is still only a guess.

See the complete [concept curriculum](../).
