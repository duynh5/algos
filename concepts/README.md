# Concept Curriculum

This index is the learning roadmap and coverage tracker. Linked entries already
have an explanation. Dedicated leaf pages also include tested Go code when an
implementation is useful. Unlinked entries are planned and should be added when
studying the topic or when a solved problem needs them.

When the difficult part is choosing an approach, start with the
[problem-solving method selection guide](patterns/#problem-solving-method-selection-guide).
It starts with a quick recognition guide, then expands unfamiliar methods into
plain-language pattern cards. Each card explains what to picture, how to begin,
the fact that remains true and makes the method safe, and the condition that
should make you reject it.

Every concept page should explain:

- a plain-language mental model;
- what signals suggest the concept and what signals rule it out;
- the first implementation steps to try;
- the fact that must remain true or the rule that determines each decision;
- how work and extra memory grow with the input;
- a Go template or implementation that follows the repository's usual style;
- common mistakes and cases at the smallest or largest allowed limits;
- links to representative solved problems.

This curriculum is a living reference. When a reviewed problem introduces a
new method, recognition clue, correctness rule, warning, or example,
update the owning concept page. Update the method selection guide when that
evidence changes how a solver should choose between approaches. Link to the
owning page instead of duplicating its detailed explanation.

## Foundations

- Complexity analysis: how work and extra memory grow with the input, including
  the average cost of an operation across a longer sequence of operations
- Correctness: facts that remain true, proofs built from smaller cases, and
  reasoning about each pass through a loop
- Recursion and the memory used by nested function calls
- [Integer digit processing](foundations/integer-digits/)
- Integer limits, overflow, and arithmetic that keeps only a remainder
- Operations on the individual binary digits of an integer

## Data structures

- Arrays, slices, and strings
- [Linked lists](data-structures/linked-list/)
- [Stack](data-structures/stack/)
- Queues that remove the oldest item first, including queues that can add or
  remove items at both ends
- [Hash maps and sets](data-structures/hash-map/)
- [Heaps and priority queues](patterns/#heap-or-priority-queue)
- Trees and binary search trees
- [Tries](patterns/#trie)
- [Disjoint-set union](patterns/#disjoint-set-union)
- Ways to represent networks of connected items
- Structures that answer and update information about numeric ranges quickly

## Algorithms

- Sorting values and selecting the item at a requested sorted position
- [Binary search and bounds](algorithms/binary-search/)
- [Breadth-first search](patterns/#breadth-first-search) and
  [depth-first search](patterns/#depth-first-search)
- [Ordering tasks from their prerequisites](patterns/#topological-sort)
- [Cheapest routes through a network](patterns/#dijkstra-shortest-paths)
- Connecting every location with the smallest possible total connection cost
- Solving a problem by splitting it into independent smaller parts
- [Greedy algorithms](patterns/#greedy)
- [Backtracking](patterns/#backtracking)
- [Dynamic programming](algorithms/dynamic-programming/)
- Finding text patterns inside larger text
- Properties of integers and methods for counting possible arrangements

## Problem-solving patterns

- [Common-prefix scans](patterns/common-prefix-scan/)
- [Directional scans](patterns/directional-scan/)
- [Two pointers](patterns/two-pointers/)
- [Sliding window](patterns/#sliding-window)
- [Prefix sums](patterns/#prefix-sums) and
  [difference arrays](patterns/#difference-arrays)
- [Fast and slow pointers](patterns/#fast-and-slow-pointers)
- [Stacks kept in increasing or decreasing order](patterns/#monotonic-stack)
  and queues kept in one of those orders
- [Interval merging](patterns/#intervals)
- Putting each value into the array position associated with that value
- [Sweep line](patterns/#sweep-line)
- Searching possible choices and undoing them when necessary

## Go techniques

- Input and output formats accepted by online judges
- Running many examples through the same test code
- Sorting with a problem-specific order
- Go's `container/heap` support for priority queues
- Reusable structures that work with multiple Go types
- Limits on deeply nested function calls and loop-based alternatives
- When Go requests or retains memory, including slice reuse and avoiding
  unintended retention
