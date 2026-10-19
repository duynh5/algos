# Algorithms and Data Structures in Go

This is the repository's single book-like reference for algorithmic reasoning,
data structures, problem-solving patterns, and the Go details used to implement
them. Problem archives remain the source of complete judge solutions; this book
focuses on recognition, invariants, correctness, complexity, and reusable code.

Use the book in two directions:

- When starting a problem, begin with [method selection](#problem-solving-method-selection-guide).
- When studying, read foundations first, then data structures, patterns, and
  the preserved Go reference implementations.

## Contents

- [Foundations](#foundations)
- [Data structures](#data-structures)
- [Problem-solving patterns](#problem-solving-patterns)
- [Go reference implementations](#go-reference-implementations)
- [Go techniques](#go-techniques)
- [Further-study checklist](#further-study-checklist)

## Foundations

### Complexity and resource accounting

Let `n` describe the meaningful input size. Time complexity counts how work
grows as `n` grows; auxiliary-space complexity counts memory used beyond the
input and required output. Include preprocessing, nested function calls, Go
allocations, and the cost of producing the output. State best, average, and
worst cases when control flow makes them differ.

Useful growth rates are `O(1)`, `O(log n)`, `O(n)`, `O(n log n)`, `O(n^2)`,
`O(2^n)`, and `O(n!)`. Constants still matter, but input constraints usually
eliminate the wrong growth rate before implementation details matter.

### Correctness and invariants

An invariant is a fact that is true before and after every step. A complete
argument identifies:

1. **Initialization:** why the fact is true before the first step.
2. **Preservation:** why one step keeps it true.
3. **Progress:** why the algorithm moves toward termination.
4. **Coverage:** why discarded or skipped candidates cannot contain a missing
   answer.
5. **Termination:** why the invariant implies the returned result is correct.

### Recursion and explicit state

Recursive calls retain parameters, local variables, and a return position.
A recursion depth of `d` therefore uses `O(d)` call-stack space even when no
collection is allocated. Prefer an explicit stack when the depth may exceed
the runtime's safe call depth or when the traversal state must be inspected
or controlled directly.

### Integer limits and bits

Check arithmetic before performing an operation that may overflow. Rearranged
comparisons such as `value > (limit-digit)/10` can prove that multiplication
and addition would be unsafe. For bit work, define whether values are signed
or unsigned and remember that shifts, masks, and bit counts operate on the
chosen fixed-width representation.

### Integer Digit Processing

#### Mental model

An integer's decimal digits can be removed from right to left without first
turning the number into text:

- `value % 10` returns the final digit;
- `value / 10` removes the final digit because integer division drops the
  remainder.

For example, `123 % 10` is `3`, and `123 / 10` is `12`.

#### Worked example: reversing digits

To move each extracted digit into a new integer:

```go
reversed = reversed*10 + value%10
value /= 10
```

Starting with `value = 120` and `reversed = 0`:

| Step | Extracted digit | New `reversed` | Remaining `value` |
| --- | --- | --- | --- |
| 1 | `0` | `0` | `12` |
| 2 | `2` | `2` | `1` |
| 3 | `1` | `21` | `0` |

The result is `21`, not `021`, because an integer does not store leading zeroes.

#### Invariant (fact that remains true)

After `k` passes, `value` is the original number with its final `k` digits
removed. `reversed` contains those removed digits in the opposite order. When
`value` reaches zero, every digit has moved.

#### Preventing integer overflow

Multiplying by `10` may exceed the largest value the integer type can hold.
Before calculating `reversed*10 + digit`, check:

```go
reversed > (limit-digit)/10
```

If the comparison is true, the next result would exceed `limit`. The reference
implementation returns `false` instead of performing that unsafe calculation.

Negative inputs need an explicit rule because the minus sign is not a digit.
`ReverseNonNegative` rejects them and accepts zero.

#### Work and memory

For a number with `d` digits, the loop performs `d` passes, so it uses `O(d)`
work. It stores only the current digit and two integers, so it uses `O(1)` extra
memory.

#### Related problems

- [LeetCode 9: Palindrome Number](../problems/solved_easy_leetcode_0009_palindrome_number.go)
  compares a full reversal with the original number and also shows how reversing
  only half of the digits can prove the same result.

## Data structures

### Arrays, slices, and strings

An array owns a fixed number of values. A slice is a view over an underlying
array and carries a length and capacity; appending may reuse that array or
allocate a new one. A substring or subslice can retain a much larger backing
array, so copy a small result when long-term retention would keep unwanted
data alive. Go strings are immutable byte sequences; indexing returns a byte,
not necessarily a complete Unicode character.

### Linked Lists

#### Mental model

A singly linked list is a chain of nodes. Each node stores one value and a link
to the next node. The first node is the **head**; a `nil` next link marks the
end.

Unlike an array, a linked list cannot jump directly to position `i`. Reaching a
position requires following links from the head. In exchange, code can reconnect
known nodes without shifting every later value.

#### Core operations

- `FromSlice` builds a list in the slice's order.
- `ToSlice` follows links from the head and collects the values.
- `Reverse` changes every next link so the chain points in the opposite
  direction.

#### Worked example: reversing a list

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

#### Invariants (facts that remain true)

- Following `Next` from the head visits nodes in list order and eventually
  reaches `nil` when the list contains no loop.
- While building a list with a `tail` position, `tail` always names the final
  node. A new node can therefore be attached without walking from the head.
- During reversal, `previous` is the already reversed portion and `current` is
  the first node not yet reversed.

#### Common techniques

- A **dummy head** is a temporary node placed before the real first node. It
  avoids separate code for inserting or replacing the head.
- Two positions can advance independently when input lists have different
  lengths.
- Fast and slow positions can detect a loop or locate the middle.
- For numbers stored with the lowest digit first, add matching nodes and carry
  `sum / 10` to the next position; the current output digit is `sum % 10`.

#### Work and memory

| Operation | Work | Extra memory |
| --- | --- | --- |
| Visit or search the entire list | `O(n)` | `O(1)` |
| Insert after a node already in hand | `O(1)` | `O(1)` |
| Reverse with a loop | `O(n)` | `O(1)` |
| Copy all values to or from a slice | `O(n)` | `O(n)` |

These memory amounts do not count nodes required for a returned result. A
version that uses nested function calls can also keep as many as `n` unfinished
calls in memory.

#### Common mistakes

- Save `node.Next` before replacing it during reversal.
- Make loop conditions handle an empty list and a one-node list.
- Keep the original head when it is still needed; advancing the only reference
  can make earlier nodes unreachable.
- Following links one by one never ends if the list links back to an earlier
  node.
- Prefer a loop when a list may be too long for deeply nested function calls.

#### Related problems

- [LeetCode 2: Add Two Numbers](../problems/solved_easy_leetcode_0002_add_two_numbers.go)
- [LeetCode 21: Merge Two Sorted Lists](../problems/solved_easy_leetcode_0021_merge_two_sorted_lists.go)

### Choosing a structure

- Use a [hash map or set](#hash-map-or-set) for exact-key lookup.
- Use a [stack](#stack) for most-recent-first processing and a queue for
  oldest-first processing.
- Use a [heap](#heap-or-priority-queue) when only the next smallest, largest,
  or highest-priority item must be exposed.
- Use a [trie](#trie) for repeated prefix lookup.
- Use [disjoint-set union](#disjoint-set-union) for connectivity under
  incremental unions.
- Use linked nodes when reconnecting known positions matters more than random
  access.

## Problem-Solving Patterns

Pattern cards explain how to recognize a useful method and why that method is
safe. That explanation usually centers on an **invariant**: a fact that remains
true while the algorithm runs.

### Problem-solving method selection guide

Treat recognition signals as possible clues, not keyword matches. An array does
not automatically imply two pointers, and the word "minimum" does not
automatically imply greedy. A method fits only when you can state the invariant
that makes its next step safe.

Use this guide as a sequence of progressively stronger hints. Start with the
decision tree. Open the linked pattern card only when the tree leaves you
unsure how to begin, then try the card's first steps before reading a solution
explanation. This file owns the complete pattern guide; no separate concept
page is required.

#### Quick decision tree

Before choosing a branch, answer three basic questions:

1. **What must you return?** A value, position, path, ordering, one valid
   answer, or every valid answer?
2. **May you rearrange the input?** Check whether items must stay consecutive
   or retain their original order or indices.
3. **How large can the input be?** Eliminate methods that would use too much
   time or memory at the maximum size.

Now choose the branch that best describes the hardest part of the problem. A
leaf suggests a method; confirm its invariant in the linked card before using
it. More than one branch may apply. For example, Dijkstra's algorithm usually
uses a heap.

- **Are you working with connections between items?**
  * Need to arrange tasks so every prerequisite comes first? Use
    **[topological sort](#topological-sort)**.
    **Example:** If course `C` requires both `A` and `B`, return an order with
    `A` and `B` before `C`.
  * Need the shortest or cheapest route from one starting point?
    **Example:** Find the cheapest route from town `A` to every other town on a
    road map.
    + Every move costs the same? Use
      **[breadth-first search](#breadth-first-search)**.
      **Example:** In a maze, each step to an adjacent open cell costs one move.
    + Costs vary, but none are negative? Use
      **[Dijkstra's algorithm](#dijkstra-shortest-paths)**.
      **Example:** Roads have travel times such as `2`, `5`, and `9` minutes.
    + Some costs are negative? Use
      **[Bellman-Ford](#bellman-ford-shortest-paths)**.
      **Example:** A route can include a `-3` rebate, so a later path may make
      an earlier distance cheaper.
  * Are connections added one at a time while you repeatedly ask whether two
    items are connected? Use **[disjoint-set union](#disjoint-set-union)**.
    **Example:** Add roads one by one and, after each road, ask whether towns
    `1` and `5` are now connected.
  * Need to completely explore one path, region, or subtree before the next?
    Use **[depth-first search](#depth-first-search)**.
    **Example:** Calculate a directory's total size by finishing every child
    directory before returning its size to the parent.

- **Are you building an answer by choosing among alternatives?**
  * Can you prove that one choice now will not worsen the best final result?
    Use **[greedy](#greedy)**.
    **Example:** To attend the most meetings, always choose the compatible
    meeting that finishes earliest.
  * Do many choice sequences reach the same remaining situation, and you need
    only its best value, number of solutions, or yes/no result? Use
    **[dynamic programming](#dynamic-programming)**.
    **Example:** Different coin choices can leave the same remaining amount;
    calculate the minimum coins for that amount once.
  * Must you try a choice, explore it, undo it, and try another? Use
    **[backtracking](#backtracking)**.
    **Example:** Place a queen in one column, remove it if a later row cannot be
    completed, and try the next column.

- **Do you need fast repeated lookups?**
  * Looking up an exact value, frequency, matching partner, or group? Use a
    **[hash map or set](#hash-map-or-set)**.
    **Example:** While looking for a pair totaling `9`, ask whether `9 - x` has
    already appeared for each number `x`.
  * Looking up words by a shared beginning? Use a **[trie](#trie)**.
    **Example:** For autocomplete after `ca`, keep `car`, `cart`, and `cat`
    available without scanning every word.

- **Does the problem revolve around ranges?**
  * Need to merge overlaps or find gaps? Use **[intervals](#intervals)**.
    **Example:** Merge reservations `[1, 4]` and `[3, 6]` into `[1, 6]`.
  * Need to know how many ranges are active as positions or times change? Use a
    **[sweep line](#sweep-line)**.
    **Example:** Process booking start and end times to find the maximum number
    of rooms occupied together.
  * Need to answer many sum or count queries over ranges? Use
    **[prefix sums](#prefix-sums)**.
    **Example:** Precompute totals once, then answer many questions such as the
    sum from index `20` through index `80`.
  * Need to apply many range additions and calculate the final values only
    once? Use a **[difference array](#difference-arrays)**.
    **Example:** Apply hundreds of passenger-count changes across stop ranges,
    then reconstruct the count at every stop once.

- **Does the problem revolve around positions in a sequence?**
  * Comparing several sequences to find the beginning shared by all of them?
    Use a **[common-prefix scan](#common-prefix-scan)**.
    **Example:** Compare `"flower"`, `"flow"`, and `"flight"` one column at a
    time until the first mismatch.
  * Does an item's meaning depend on an adjacent item? Use a
    **[directional scan](#directional-scan)** that processes the needed neighbor
    first.
    **Example:** Scan a Roman numeral from right to left so the symbol to the
    right is already known.
  * Need the best consecutive section, and can update its state when either end
    moves? Use a **[sliding window](#sliding-window)**.
    **Example:** Find the longest consecutive part of `"abca"` without a
    repeated character by growing and shrinking one window.
  * Searching sorted values, or the first point where a yes/no test changes?
    Use **[binary search](#binary-search)**.
    **Example:** In `[1, 3, 3, 5]`, find the first position whose value is at
    least `3`.
  * Can two positions move through the sequence while each move safely rules
    out possibilities? Use **[two pointers](#two-pointers)**.
    **Example:** In a sorted array, move the left or right position according
    to whether their sum is too small or too large.
  * Working with a linked structure and need a cycle, midpoint, or fixed gap?
    Use **[fast and slow pointers](#fast-and-slow-pointers)**.
    **Example:** Move one linked-list pointer one step and another two steps;
    they meet if the list contains a cycle.
  * Need the next or nearest greater or smaller value? Use a
    **[monotonic stack](#monotonic-stack)**.
    **Example:** For each daily temperature, find the next later day with a
    higher temperature.

- **Does the order of waiting items determine what happens next?**
  * Must the most recently opened or unfinished item be handled first? Use a
    **[stack](#stack)**.
    **Example:** When checking `([])`, the closing bracket must match the most
    recent unmatched opening bracket.
  * Must you repeatedly remove the smallest or largest item while the stored
    items keep changing? Use a
    **[heap or priority queue](#heap-or-priority-queue)**.
    **Example:** As jobs arrive, repeatedly run the waiting job with the highest
    priority.

If nothing fits, return to the simplest correct solution and ask what work it
repeats. The problem may need a method outside this guide; do not force a
match.

#### Derive a direction before naming a pattern

Do not begin by guessing an algorithm name. First work through a tiny example
by hand and describe the simplest correct method, even if it tries every
possibility. Then ask what that method repeatedly recalculates or revisits.

For example, "for every starting position, I add the same later values again"
suggests storing totals that can be updated or reused. That observation may
lead to a sliding window or prefix sums. It is more useful than matching one
word in the prompt to a memorized technique.

Use this hint ladder during practice:

1. **First 5 minutes:** write or explain the simplest correct method, even if it
   tries every possibility.
2. **At 5 minutes:** use the decision tree to identify at most two possible
   methods.
3. **At 10 to 15 minutes:** read the relevant pattern card and check
   whether its reason for discarding work applies to the problem.
4. **At 20 minutes:** trace the method on one small input before writing code.
5. **At 25 to 30 minutes:** read a full explanation if still blocked, close it,
   and implement the idea without copying its code.

#### Pattern cards

Each card is the complete conceptual reference for its method: what to picture,
when it fits, how to begin, why it is safe, and when to reject it. The cards are
deliberately method-level rather than problem-specific, so they help produce a
first attempt without revealing a judge solution.

Jump directly to a card:

- Sequence and range methods: [common-prefix scan](#common-prefix-scan),
  [directional scan](#directional-scan), [two pointers](#two-pointers),
  [binary search](#binary-search), [sliding window](#sliding-window),
  [prefix sums](#prefix-sums), [difference arrays](#difference-arrays),
  [fast and slow pointers](#fast-and-slow-pointers),
  [monotonic stack](#monotonic-stack), [intervals](#intervals), and
  [sweep line](#sweep-line).
- Lookup and processing-order methods: [hash map or set](#hash-map-or-set),
  [trie](#trie), [stack](#stack), and
  [heap or priority queue](#heap-or-priority-queue).
- Choice methods: [dynamic programming](#dynamic-programming),
  [greedy](#greedy), and [backtracking](#backtracking).
- Relationship methods: [depth-first search](#depth-first-search),
  [breadth-first search](#breadth-first-search),
  [topological sort](#topological-sort),
  [Dijkstra](#dijkstra-shortest-paths),
  [Bellman-Ford](#bellman-ford-shortest-paths), and
  [disjoint-set union](#disjoint-set-union).

##### Common-prefix scan

**Mental model:** Compare several strings one character position at a time,
beginning at the left. The answer ends as soon as one string ends or one
character differs, because a shared beginning cannot skip that position.

**Recognize it:** The answer must start at position `0` in every sequence. If
the answer may begin later or skip positions, this is not a common-prefix scan.

**Start here:** Use the first string as the proposed prefix. For each position,
check that every other string is long enough and contains the same character.
Return the part before the first failed check.

**Worked example:** For `"flower"`, `"flow"`, and `"flight"`, position `0` is
`f` everywhere and position `1` is `l` everywhere. Position `2` contains `o`,
`o`, and `i`, so the longest common prefix is `"fl"`.

**Invariant:** Before checking position `c`, every string shares the prefix
`strs[0][:c]`. A mismatch or ended string proves that prefix is maximal;
otherwise the proven prefix grows by one character.

**Work and memory:** The scan uses `O(S)` time in the worst case, where `S` is
the total input size, and `O(1)` auxiliary space when it returns a slice of the
first string.

**Common traps:** Check a string's length before reading its current position.
Define the result for empty input. Also decide whether a position means a byte,
a Unicode code point, or a user-visible character before applying this method
to general text.

**Reject it when:** The shared result may start after position `0`, skip
positions, or compare reordered items.

**Representative problem:** [Longest Common Prefix](../problems/solved_easy_leetcode_0014_longest_common_prefix.go).

##### Directional scan

**Mental model:** When an item's meaning depends on a neighboring item, scan in
the direction that makes that neighbor already known instead of repeatedly
looking ahead.

**Recognize it:** A local rule refers consistently to the item immediately to
the right or left. A right-neighbor rule suggests scanning right to left; a
left-neighbor rule usually suggests scanning left to right.

**Start here:** State which neighbor determines the current item's treatment,
choose the direction that processes that neighbor first, and retain exactly the
neighboring state needed for the next step.

**Worked example:** Roman numeral `IV` represents `4`. Scanning numeric values
`[1, 5]` from right to left first adds `5`, then sees that `1` is smaller than
the remembered right value and subtracts it, producing `4`.

**Invariant:** Before processing a position, the already-scanned side has been
handled correctly and the retained value represents the adjacent item needed
by the rule. Applying the rule once extends the correct result by one position.

**Work and memory:** A single scan takes `O(n)` time and `O(1)` auxiliary space
when only one neighboring value is retained.

**Common traps:** The neighboring-value rule must come from the problem; an
adjacent comparison is not automatically correct. Handle the first processed
item explicitly when a sentinel would be invalid, and update the remembered
neighbor after every item, including one that contributes negatively.

**Reject it when:** The current result depends on nonlocal state that cannot be
made available by choosing one scan direction.

**Representative problem:** [Roman to Integer](../problems/solved_easy_leetcode_0013_roman_to_integer.go).

##### Hash map or set

**Mental model:** Store a fact under a key that can be calculated again. A map
returns associated information; a set records presence only.

**Recognize it:** You repeatedly ask whether a value appeared, how often it
appeared, which earlier value completes a pair, or which group an item belongs
to. Use a set when no value beyond presence is needed.

**Start here:** Define the exact key and the meaning of its stored value. Decide
whether the current item must query the map before or after it is inserted, and
whether repeated keys keep the first value, latest value, count, or collection.

**Worked example:** For `[2, 7, 11, 15]` and target `9`, the scan first asks for
`7`, finds nothing, and stores `2 -> 0`. At value `7`, it asks for `9 - 7 = 2`,
finds position `0`, and returns positions `0` and `1`. Querying before insertion
prevents an item from pairing with itself while still allowing `[3, 3]`.

**Invariant:** Before processing position `i`, the map contains exactly the
facts retained from positions before `i`, under their defined keys. It contains
nothing from the current or later positions.

**Work and memory:** A scan over `n` items usually takes expected `O(n)` time
with expected constant-time map operations and up to `O(n)` auxiliary space.
An individual hash operation is not guaranteed to be constant time.

**Common traps:** A map does not preserve sorted or insertion order by itself.
In Go, reading a missing key returns the value type's zero value, so use the
`value, found := entries[key]` form when zero is meaningful.

**Reject it when:** The essential operation is ordered traversal, nearest-key
search, or another relationship a hash key does not preserve.

**Representative problem:** [Two Sum](../problems/solved_easy_leetcode_0001_two_sum.go).

##### Binary search

**Mental model:** Test the middle possible answer and discard one half. The
method is safe only when the possibilities form two consecutive groups, such
as all failures followed by all successes.

**Recognize it:** Data is sorted, the problem asks for the first or last
position satisfying a condition, or an ordered numeric answer space has a
single boundary where a yes/no test changes. A numeric answer alone is not
enough.

**Start here:** Define precisely what `left` and `right` include. For a
half-open range `[left, right)`, test `middle = left + (right-left)/2` and keep
the half that can still contain the boundary. Decide which side retains an
equal value based on whether the first or last valid position is required.

**Worked example:** To find the first value at least `3` in `[1, 3, 3, 5]`,
start with `[0, 4)`. Testing value `3` at position `2` keeps `[0, 2)`; testing
value `3` at position `1` keeps `[0, 1)`; testing value `1` discards position
`0`. The empty range meets at position `1`, the first valid position.

**Invariant:** Before every test, the first acceptable position lies in
`[left, right)`. Each update discards only positions proved too early or too
late; when the range is empty, `left` is the boundary.

**Wrapped sorted data:** In data such as `[4, 5, 6, 1, 2, 3]`, the whole range
is not sorted, but at least one side of the middle is. Identify that side and
use its outer endpoint to determine whether the target belongs there.

**Work and memory:** Each test removes about half of the remaining positions,
so an iterative search takes `O(log n)` time and `O(1)` auxiliary space.

**Common traps:** Keeping the tested middle can prevent progress. A valid
insertion position may equal `len(values)`. Ordinary bounds require fully
sorted input; wrapped input needs the separate sorted-side rule.

**Reject it when:** Testing a later candidate can change from yes back to no,
or the middle test cannot prove that one entire half is impossible.

**Representative problems:** [Search Insert Position](../problems/solved_easy_leetcode_0035_search_insert_position.go)
and [Search in Rotated Sorted Array](../problems/solved_medium_leetcode_0033_search_in_rotated_sorted_array.go).

##### Two pointers

**Mental model:** Track two meaningful positions and move them in a coordinated
way so the scan never needs to restart. Common forms move inward from both
ends, merge two sorted inputs, use separate read and write positions, or expand
outward from a possible center.

**Recognize it:** Look for sorted pair selection, comparisons between two ends,
merging sorted inputs, modifying a sequence in place, or testing symmetric
positions. For every movement, you must explain why the position left behind
cannot be needed again.

**Sorted pair example:** Find a pair totaling `9` in `[1, 2, 4, 7, 11]`.
`1 + 11` is too large, so move the right position left; moving the left one
would only increase the sum. Then `1 + 7` is too small, so move the left
position right and find `2 + 7 = 9`.

**Sorted-pair invariant:** No pair using a position already passed by `left` or
`right` can meet the target. Sorted order makes each discard safe.

**Merging sorted inputs:** Track the smallest unread value in each input, copy
the smaller one, and advance only its position. When merging into unused space
at the end of the first input, work backward from the largest values so a write
cannot overwrite unread input.

**Read and write positions:** Let `read` inspect every value and `write` mark
where the next retained value belongs. The invariant is that the part before
`write` contains exactly the retained values seen so far, in their required
order. Because `write` never passes `read`, writing cannot destroy unread data.

**Expanding around centers:** Every palindrome is centered on either one
character or one gap. Test all character and gap centers, then move left and
right while the characters match. The already-accepted interior must be a
palindrome before a matching outer pair is added.

**Work and memory:** Pair search, merging, and in-place compaction take linear
time. They use `O(1)` auxiliary space unless a merged output is allocated.
Expanding around all palindrome centers takes between `O(n)` and `O(n^2)` time
and `O(1)` auxiliary space.

**Common traps:** Sorting may destroy required original positions. The inward
pair-sum form requires sorted values, though negative values are fine. Use the
provided populated lengths when an array contains spare capacity, advance a
write position only after retaining a value, and test both character and gap
centers for palindromes.

**Reject it when:** No pointer movement can be proved to preserve every
remaining answer, or the method would require sorting data whose original
positions or order must remain intact.

**Representative problems:** [Palindrome Number](../problems/solved_easy_leetcode_0009_palindrome_number.go),
[Longest Palindromic Substring](../problems/solved_medium_leetcode_0005_longest_palindromic_substring.go),
[Remove Duplicates from Sorted Array](../problems/solved_easy_leetcode_0026_remove_duplicates_from_sorted_array.go),
and [Remove Element](../problems/solved_easy_leetcode_0027_remove_element.go).

##### Dynamic programming

**Mental model:** Start with a search that tries choices. When different choice
sequences keep reaching the same smaller problem, solve that state once and
reuse its answer.

**Recognize it:** All three conditions should hold:

1. A large answer can be built from answers to smaller versions.
2. A direct search calculates the same smaller answer more than once.
3. A limited set of facts completely identifies each smaller problem.

That identifying set of facts is the **state**. Define it in a sentence before
choosing an array or map.

**Start here:** Write the recurrence, its base cases, and the order in which
dependencies become available. Use top-down recursion with caching when it is
natural to begin from the requested answer, or bottom-up iteration when smaller
states have an obvious order.

**Worked example:** If each move climbs one or two stairs, every route to stair
`4` ends with one step from stair `3` or two steps from stair `2`. Therefore
`ways[4] = ways[3] + ways[2]`. Starting with `ways[0] = 1` and `ways[1] = 1`
produces `1, 1, 2, 3, 5` without recalculating earlier stairs.

**Why the recurrence is complete:** Every non-empty route has exactly one final
move. Removing it places the route in exactly one predecessor group, and the
allowed moves list every such group, so their stored counts include every route
once.

**Work and memory:** Without caching, the stair search grows roughly like
`2^n`. Storing each state reduces it to `O(n)` time and `O(n)` space. When a
state needs only a fixed number of earlier answers, a rolling window can reduce
auxiliary space to `O(1)`.

**Common traps:** A recurrence alone does not prevent repeated work. Do not add
state facts that cannot affect future choices. Choose base cases that agree
with the recurrence, update rolling values without destroying a still-needed
value, count recursive-call memory, and check whether result counts can exceed
the chosen integer type.

**Reject it when:** Subproblems do not repeat, the proposed state omits facts
that affect future choices, or the number of distinct states is itself too
large for the constraints.

**Representative problem:** [Climbing Stairs](../problems/solved_easy_leetcode_0070_climbing_stairs.go).

##### Stack

**Mental model:** Add and remove only at the top, so the most recently added
item is the first removed. This is last in, first out.

**Recognize it:** The most recently opened or unfinished item must be handled
first, as with nested brackets, undo history, expression evaluation, or a
depth-first search implemented with a loop.

**Start here:** Define what each stored item represents and what makes it
finished. Push newly opened work, inspect or pop only the top, and decide how an
empty stack is reported.

In the Go reference implementation, `Push` adds a value, `Pop` removes the
top, `Peek` inspects it, `Len` reports the count, and `IsEmpty` reports whether
the stack contains no values. `Pop` and `Peek` return a zero value plus `false`
when the stack is empty.

**Worked example:** To validate `([])`, push `(` and then `[`. When `]` arrives,
pop `[` and verify the match. When `)` arrives, pop `(`. The input is valid
because every closer matched and the stack is empty at the end.

**Invariant:** The stack contains exactly the opened items not yet completed.
The top is the most recent unfinished item and must be completed first.

**Work and memory:** Push, pop, and top inspection take `O(1)` time, aside from
occasional slice growth when Go uses a slice as the stack. The stack may retain
`O(n)` items. Clear a removed slice position when it may hold pointers that
should become collectable.

**Common traps:** Use a queue when the oldest waiting item must be handled
first. Removing from the front of a Go slice does not model a stack and may
retain a large backing array. Recursive calls already use an implicit stack;
use an explicit one when depth may be unsafe or direct control is needed.

**Reject it when:** The oldest item must be processed first or the algorithm
needs efficient access to arbitrary stored items.

**Representative problem:** [Valid Parentheses](../problems/solved_easy_leetcode_0020_valid_parentheses.go).

##### Sliding window

**Mental model:** Put a movable fence around one consecutive portion of the
input. The fence always contains neighboring items with no gaps.
Extend the right side to consider new values and move the left side when earlier
values must leave.

**Recognize it:** The prompt asks for a longest, shortest, largest, or smallest
consecutive portion of an array or string. You can update the tracked sum,
character counts, or other information when one item enters or leaves.

**Start here:** Define exactly what condition the current range must meet. Move
`right` once per step and add that item to the tracked information. Move `left`
and remove items until the range meets the condition again. Record the answer
at the moment required by the problem.

**Why moving the window is safe:** Before advancing `left`, explain why no
future value could make the current starting position part of a better answer.
For example, when all values are positive and the sum is already too large,
adding more values can only make it larger.

**Reject it when:** The answer may skip items, or adding an item can make the
condition switch back and forth unpredictably. Mixed positive and negative
values often cause that problem for sum-based windows.

##### Prefix sums

**Mental model:** Store a running total at every position so later range totals
can be calculated without adding the same values again.

**Recognize it:** Many questions ask for sums or counts over overlapping ranges.
The simplest solution repeatedly adds values it has already added for an
earlier question.

**Start here:** Let `prefix[i]` mean "the total of the first `i` values." Include
`prefix[0] = 0`. The sum from `left` through `right` is then
`prefix[right+1] - prefix[left]`.

**Example:** For `[2, 4, 1, 3]`, store `[0, 2, 6, 7, 10]`. The sum from index
`1` through index `3` is `prefix[4] - prefix[1] = 10 - 2 = 8`.

**Why subtraction works:** Both stored totals begin at index `0`. Subtracting
the total before `left` removes exactly the values that are outside the
requested range.

**Reject it when:** The requested property cannot be recovered by combining two
running summaries. For example, you cannot find a range's smallest value by
subtracting two previously stored smallest values.

##### Difference arrays

**Mental model:** When many updates affect whole ranges, record where each
change starts and stops instead of modifying every covered position
immediately.

**Recognize it:** The problem applies many additions or subtractions to ranges
and asks for the final array after all updates.

**Start here:** Create a bookkeeping array with one more position than the
result array. For an update that adds `value` from `left` through `right`, add
`value` at `left` and subtract `value` at `right + 1`. After recording every
update, scan the positions of the result array from left to right while keeping
a running total of the changes.

**Example:** Begin with `[0, 0, 0, 0]` and add `3` to indices `1` through `3`.
In the five-position bookkeeping array, record `+3` at position `1` and `-3`
at position `4`. Running totals across the original four positions become
`[0, 3, 3, 3]`; position `4` merely marks where the update stops.

**Why it works:** The running change increases when an update begins and
decreases immediately after that update ends. Every position between those
points receives the update exactly once.

**Reject it when:** You must answer detailed questions after every individual
update rather than after the updates have been collected.

##### Greedy

**Mental model:** Make one choice now and never undo it because you can prove
that it does not reduce the best result still possible.

**Recognize it:** Look for a problem that asks for a minimum or maximum and
builds the answer through repeated choices, such as selecting items, scheduling
jobs, or connecting nodes. Greedy is a candidate when one available choice may
be taken permanently without lowering the best result achievable afterward.

The deciding question is: **after making this choice permanently, can you prove
that the best result still possible has not become worse?** If the answer is
yes at every step, use greedy. If a discarded alternative could produce a
better result because of later choices, keep the alternatives with dynamic
programming or search instead.

Before committing to a greedy approach, verify all three claims:

1. **The objective and state are clear:** You know exactly what is being
   minimized or maximized and what information can affect the remaining
   choices.
2. **The next choice is safe:** You can prove that taking it does not reduce the
   best result still possible, even though other choices are discarded.
3. **The safety argument repeats:** After committing, the remaining work is a
   smaller problem on which the same rule and proof still apply.

**Safe to discard** has a precise meaning. Suppose the greedy rule chooses `A`
and rejects the alternatives. If the best achievable result before the choice
is `X`, it must still be possible to achieve `X` after choosing `A`. You do not
need to know the value of `X`; the proof establishes that it is unchanged.

One strong way to prove this is **dominance**. A choice `A` dominates an
alternative `B` if every feasible final answer beginning with `B` can be
replaced by one beginning with `A` that is at least as good. This comparison
must include both the value gained now and every option left for later.

Common ways to establish that comparison are:

- **Exchange:** Start with any optimal answer using `B`, replace `B` with `A`,
  and prove that the answer remains feasible and no worse.
- **Stays ahead:** Identify the state that controls all future choices and
  prove that, after every step, greedy's state is at least as favorable as the
  state produced by any competing strategy.
- **Cut property:** Divide the problem into completed and uncompleted parts and
  prove that `A` is safe among all choices crossing that boundary.

For interval scheduling, all first choices contribute one meeting. If `A` ends
no later than `B`, every meeting that could follow `B` can also follow `A`.
Therefore replacing `B` with `A` preserves every later choice, so `B` is safe
to discard for that position.

For coin values `[1, 3, 4]` and target `6`, choosing the largest coin does not
dominate choosing `3`. Choosing `4` leaves target `2`, which needs two more
coins, while choosing `3` leaves target `3`, which needs only one. The future
after `4` is worse, so the alternative `3` cannot be discarded.

Choose another method when any claim is missing:

- Use **dynamic programming** when different choices create genuinely
  different future states, several of those states may still lead to the best
  answer, and the same states recur.
- Use **backtracking or exhaustive search** when alternatives cannot be safely
  discarded and you must enumerate them, especially when there is no compact
  repeated state to cache.

**Start here:** State the objective and the exact local rule; "take the best"
is not precise enough. Now imagine a best possible answer that starts
differently. Replace its first choice with yours and check two facts: the
replacement remains valid, and its final value is no worse. Then verify that
the unchosen candidates still form the same kind of problem so the argument can
be repeated.

There is no universal greedy rule such as "always take the smallest value."
Each problem needs its own rule and proof. A practical process is:

1. Define the objective, the choices that are currently legal, and what state
   remains after one choice.
2. Propose a precise rule that makes the remaining state no worse than a
   competing choice would. Sorting, a priority queue, or a one-pass comparison
   often exposes that choice.
3. Prove the rule with an exchange argument, a stays-ahead argument, or a cut
   property. If the proof fails, do not commit greedily.
4. Repeatedly take the proven-safe choice, discard choices it makes impossible,
   and solve the remaining smaller problem with the same rule.

**Execution on one input:** Suppose one meeting may start when another ends. To
attend the maximum number of meetings from `(1, 9)`, `(2, 4)`, `(5, 7)`, and
`(8, 10)`, the best possible result before making a choice is three: the last
three meetings are compatible, while `(1, 9)` overlaps all of them.

1. Sort the meetings by end time, producing `(2, 4)`, `(5, 7)`, `(1, 9)`,
   `(8, 10)`.
2. Set `lastEnd` to negative infinity and start with no selected meetings.
3. Scan the sorted meetings. Select a meeting exactly when its start is at
   least `lastEnd`; after selecting it, set `lastEnd` to its end.

| Meeting | Test | Action | `lastEnd` afterward |
| --- | --- | --- | --- |
| `(2, 4)` | `2 >= -infinity` | Select | `4` |
| `(5, 7)` | `5 >= 4` | Select | `7` |
| `(1, 9)` | `1 < 7` | Skip | `7` |
| `(8, 10)` | `8 >= 7` | Select | `10` |

Choosing `(2, 4)` first leaves `(5, 7)` and `(8, 10)` available, so the best
result still possible remains three. The choice has not made the answer worse.
Choosing `(1, 9)` first would leave no compatible meeting, reducing the best
result still possible from three to one. That choice is therefore unsafe.
This comparison explains this input, but it does not prove the rule for every
input.

**Proof for every input:** At any step, let `R` contain every meeting compatible
with the meetings already selected, and let `E` be the meeting in `R` with the
earliest end. Consider an optimal schedule from `R`, and call its first meeting
`F`.

Because `E` has the earliest end, `E.end <= F.end`. Replace `F` with `E`. If a
later meeting `L` followed `F`, then `L.start >= F.end`. Combining the two
inequalities gives `L.start >= F.end >= E.end`, so `L` can also follow `E`.
The replacement is therefore valid and preserves the schedule's number of
meetings. This proves that at least one optimal schedule from `R` begins with
the greedy choice `E`.

After committing to `E`, remove the meetings that overlap it. The remaining
meetings form another interval-scheduling problem, so the same argument applies
again. Repeating this exchange argument proves that every commitment made by
the algorithm is safe, not merely that it works for the example above.

Sorting costs `O(n log n)` time, and the scan costs `O(n)` time. The scan uses
`O(1)` auxiliary space when the meetings are sorted in place, excluding the
space used to return the selected meetings.

**Invariant:** Each commitment preserves the best achievable final objective
value; it removes only alternatives that are unnecessary to reach that value.

**Reject it when:** A choice that looks best now can block a better combination
later, and you cannot prove that the replacement described above is harmless.
Keep multiple possibilities with dynamic programming or search instead.

##### Backtracking

**Mental model:** Walk through a tree of decisions while carrying one partial
answer. At each branch, make a choice, explore what it can lead to, and undo the
choice before trying its sibling branch.

**Recognize it:** The answer is assembled one decision at a time, and each
partial answer has a clear set of legal next choices. You may need to find one
valid construction, generate every construction, or compare a manageable
number of constructions. Constraints often let you stop exploring a branch as
soon as its partial answer cannot become valid or cannot beat the best answer
already found.

Backtracking is a candidate when all of these are true:

1. A complete answer can be represented as a sequence of choices.
2. From a partial answer, you can enumerate the choices allowed next.
3. No single next choice can be proved safe enough to discard all alternatives.
4. The search space is small enough, or invalid branches can be pruned early
   enough, to make exploration practical.

**Procedure:** Define the partial answer, any state needed to check legality,
and the condition that marks a complete answer. Then:

1. If the partial answer is complete, record a copy or report success.
2. Otherwise, list every legal next choice.
3. Apply one choice to the partial answer and its supporting state.
4. Recursively explore what can follow that choice.
5. Undo every change made in step 3.
6. Repeat steps 3 through 5 for the other choices.

If only one answer is needed, return as soon as a recursive call succeeds. If
every answer is needed, finish every branch. If a partial answer is already
invalid or cannot improve the current best result, return immediately without
exploring its descendants; this is **pruning**.

**Worked example:** Generate every subset of `[a, b]`. At each position, branch
into including the current value or excluding it:

1. Start with `[]`; include `a`, producing `[a]`.
2. Include `b`, producing the complete subset `[a, b]`; save a copy.
3. Undo the inclusion of `b`, returning to `[a]`; exclude `b` and save `[a]`.
4. Undo the inclusion of `a`, returning to `[]`; exclude `a`.
5. Include `b` and save `[b]`; undo it, exclude it, and save `[]`.

The crucial detail is that undoing `b` restores exactly `[a]`, and undoing `a`
restores exactly `[]`. Each sibling branch therefore begins with only the
choices made on its own path.

Generating subsets visits `2^n` complete branches. Saving a copy of every
subset takes `O(n * 2^n)` total time and `O(n * 2^n)` output space. The current
path and recursive calls use `O(n)` auxiliary space, excluding the output.

**Invariant:** At the start of each recursive call, the partial answer and all
supporting state describe exactly the choices on the path to that call. After
the call returns, undoing its choice restores the caller's state exactly.

**Choose another method when:**

- Use **greedy** when one next choice can be proved not to reduce the best
  result still possible, so the other branches may be discarded.
- Use **dynamic programming** when many choice sequences reach the same
  remaining state and only its best value, count, or feasibility matters. Cache
  that state instead of rebuilding the same search tree.
- Use ordinary **DFS** when traversing an existing graph or tree without
  constructing and undoing a candidate answer.
- Avoid exhaustive backtracking when the search space is too large and no
  effective pruning rule or exploitable repeated state exists.

##### Depth-first search

**Mental model:** Follow one route as far as possible. When that route ends,
return to the most recent place that had another route available.

**Recognize it:** You need to visit an entire connected region, inspect every
complete path, or calculate a tree result from the results below each node. You
do not need the first discovered path to use the fewest moves.

**Start here:** Define one position and how to list every position reachable in
one move. Mark a position reached before exploring from it so a loop cannot send
the search around forever. Use nested function calls or your own stack to
remember where to return.

**Example:** To calculate each directory's total size, descend through its files
and subdirectories first. Once a subdirectory has been fully explored, return
its total and add it to its parent directory's size.

**Invariant:** Every position marked reached has either been fully explored or
is waiting on the current return path. No position needs to be scheduled twice.

**Common trap:** Nested function calls consume memory. A path with too many
steps may exceed the program's call limit, so use your own stack collection for
very deep inputs.

##### Breadth-first search

**Mental model:** Explore outward in layers: first everything one move away,
then everything two moves away, and so on.

**Recognize it:** You need the fewest moves through a network where every move
costs the same, the nearest reachable position, or tree nodes grouped by depth.

**Start here:** Use a first-in, first-out queue. Add the starting position and
mark it reached immediately. Repeatedly remove the oldest position, add each
unseen neighbor, and record its distance as one more than the current distance.

**Example:** In a maze, process the start at distance `0`, its open neighbors at
distance `1`, and their unseen neighbors at distance `2`. The first time the
exit is added, you have found a route with the fewest moves.

**Invariant:** The queue never processes a position with a greater distance
before a position with a smaller distance. Each position is added at most once.

**Reject it when:** Moves have different costs. The route with the fewest moves
may then cost more than a longer route; use Dijkstra's algorithm when all costs
are zero or greater.

##### Topological sort

**Mental model:** Repeatedly perform tasks whose prerequisites are already
satisfied.

**Recognize it:** The prompt describes courses, builds, jobs, or other tasks
with rules such as "A must finish before B."

**Start here:** For every prerequisite rule, record which later task it unlocks.
Count each task's unfinished prerequisites. Add every task with a count of zero
to a first-in, first-out queue. When one task finishes, subtract one from the
count of every task that depended on it.

**Invariant:** The queue contains exactly the tasks whose prerequisites are all
finished. A task is added to the output only after its remaining prerequisite
count reaches zero.

**Failure signal:** If the queue becomes empty before every task is processed,
the remaining tasks form a dependency loop, such as `A` needing `B` while `B`
also needs `A`. No valid order exists.

##### Dijkstra shortest paths

**Mental model:** On a map with travel costs, always process the unvisited
location with the cheapest known route from the starting location.

**Recognize it:** You need the cheapest routes through a network, and every
connection has a cost of zero or greater.

**Start here:** Record cost `0` for the starting location and "unknown" for all
others. Use a priority queue that removes the route with the smallest cost
first. When traveling through the current location gives a neighbor a cheaper
route, store that new cost and add the route to the queue.

**Why a processed cost is final:** Suppose the queue removes a route costing
`5`, and `5` is still the cheapest known cost for that location. Every untried
route already costs at least `5` before taking another connection. Because a
connection cannot subtract cost, none can produce a route below `5`.

**Reject it when:** Any connection may have a negative cost. If every connection
has the same cost, breadth-first search is simpler.

##### Bellman-Ford shortest paths

**Mental model:** Let cheaper route costs ripple through every connection,
repeating until an improvement can cross even the longest useful route.

**Recognize it:** You need shortest routes from one starting location, some
connections may have negative costs, and you may need to detect a reachable
negative-cost cycle.

**Start here:** Record cost `0` for the starting location and "unknown" for all
others. Repeat a full pass over every directed connection up to `V - 1` times,
where `V` is the number of locations. For an edge `from -> to` with cost `w`,
replace `distance[to]` with `distance[from] + w` when `from` is reachable and
that total is smaller. Stop early if a full pass changes nothing.

**Example:** Suppose `A -> B` costs `4`, `A -> C` costs `7`, and `C -> B` costs
`-5`. Starting from `A`, the first useful relaxations set `B = 4` and `C = 7`.
Relaxing `C -> B` then improves `B` to `7 + (-5) = 2`. Dijkstra cannot safely
finalize `B = 4` before that later improvement, but Bellman-Ford keeps checking
every edge.

**Why `V - 1` passes are enough:** A finite shortest route never needs to
repeat a location; removing any repeated non-negative cycle cannot make it more
expensive, while a repeated negative cycle would mean there is no finite
shortest route. A route without repeated locations uses at most `V - 1` edges.
After pass `k`, every shortest route using at most `k` edges has had a chance to
propagate its cost. Therefore `V - 1` passes cover every finite shortest route.

**Negative-cycle check:** Make one additional pass. If a reachable edge can
still lower a cost, a negative-cost cycle is reachable from the start. Costs on
that cycle, and on locations reachable from it, have no finite minimum.

**Complexity:** Up to `V - 1` passes over `E` connections take `O(VE)` time.
The distance collection uses `O(V)` auxiliary space, excluding the graph
representation and optional predecessor data used to reconstruct routes.

**Reject it when:** Every cost is non-negative and the graph is large enough
that Dijkstra's better running time matters. If a reachable negative cycle may
exist, do not report ordinary finite shortest costs for locations affected by
that cycle.

##### Disjoint-set union

**Mental model:** Imagine that every connected group has one shared group
leader. To check whether two items are connected, find each item's leader and
compare them. To connect two different groups, make one leader report to the
other.

**Recognize it:** Connections are added over time, and queries ask whether two
items are already connected, whether a new connection creates a loop, or how
many separate groups remain.

**Start here:** Store a `parent` for every item. Initially, each item is its own
parent because every item is in a separate group. `find(item)` follows parent
links until it reaches the group leader. `union(a, b)` finds both leaders and,
if they differ, makes one leader the parent of the other.

**Walk-through:** Start with `parent[1] = 1`, `parent[2] = 2`, and
`parent[3] = 3`. After connecting `1` and `2`, make `2` point to `1`. After
connecting `2` and `3`, follow `2 -> 1` and make `3` point to `1`. Now both
`find(1)` and `find(3)` reach leader `1`, so the items are connected.

**Invariant:** Two items are connected exactly when following their parent
links reaches the same group leader.

**Speed improvements:** Path compression makes visited items point closer to
the leader. Union by size attaches the smaller group below the larger group so
parent chains stay short. These change the speed, not the meaning of the
structure.

**Reject it when:** You need the actual route, the order in which items were
visited, or the shortest path. Knowing only each item's group does not preserve
that information.

##### Heap or priority queue

**Mental model:** Keep the next item you need—the smallest, largest, or most
urgent—easy to remove without fully sorting everything after each change.

**Recognize it:** You repeatedly need the smallest or largest item, a requested
number of the best items, or the next event in priority order.

**Start here:** Decide which item must come out next. A min-heap is a priority
queue that exposes the smallest stored item. To retain the three largest values
seen so far, keep three values in a min-heap; when a larger value arrives,
remove the smallest retained value and insert the new one.

**Invariant:** The heap contains exactly the items still under consideration.
The item exposed at the top is always the next smallest or largest item the
algorithm needs.

**Reject it when:** You need the complete final order of an unchanging set of
items; sorting is usually simpler. A heap also does not make it fast to find any
item other than the smallest, largest, or highest-priority one it was designed
to expose.

##### Monotonic stack

**Mental model:** Keep earlier values whose answer is still unknown, ordered in
only one direction—always increasing or always decreasing. That one-direction
order is what "monotonic" means here.

**Recognize it:** The prompt asks for the next greater or smaller value, the
nearest earlier or later value that becomes larger or smaller, or how long a
value remains the minimum or maximum.

**Start here:** Decide whether to store positions or values and whether the
stored values increase or decrease from bottom to top. While the new value
answers the question for the top item, remove that item and record its answer.
Then add the new item, whose own answer is still unknown.

**Invariant:** Every stored item is still waiting for an answer, and the stored
values remain sorted in one direction. An item is removed only when the current
value is the first later value that answers its question.

**Reject it when:** The answer depends on any earlier value, not specifically
the nearest one that meets the condition, or when removing an item would lose
information that a later value still needs.

##### Intervals

**Mental model:** Each interval covers everything between a start and an end.
Sort intervals by start so overlaps appear next to one another.

**Recognize it:** The input contains ranges, and the output must combine
overlapping ranges or report the gaps between covered regions.

**Start here:** Sort by start and keep one current interval. Compare the next
start with the current end using the problem's rule for touching endpoints. If
they overlap, extend the current end when necessary. If a gap separates them,
save the current interval and begin a new one.

**Why saving an interval is safe:** Once a gap appears, every remaining
interval starts even later because the input is sorted. None can overlap the
interval you just saved.

**Common trap:** The problem must define whether `[1, 2]` and `[2, 3]` overlap.
If both ending points are included, the ranges share `2`. If an end marks the
first excluded value, the first range stops immediately before `2`, so they do
not overlap.

##### Sweep line

**Mental model:** Move from left to right across a timeline or number line.
Record only the positions where something starts or ends, because nothing
changes between those positions.

**Recognize it:** You need the maximum number of simultaneous bookings, the
number of active ranges at each position, or another result that changes only
when a range starts or ends.

**Start here:** Turn each range into two events. A start event adds one to the
active count, and an end event subtracts one. Sort the events by position and
update the count in that order.

**Example:** One booking starts at `1` and ends just before `4`; another starts
at `3` and ends just before `5`. Process `start at 1`, `start at 3`, `end at 4`,
and `end at 5`. The active counts are `1`, `2`, `1`, and `0`, so the maximum is
`2`.

**Invariant:** After processing all events at a position, the running count
equals the number of ranges active immediately after that position.

**Common trap:** Events at the same position need an explicit rule. If a booking
ending at `4` frees its room before another starts at `4`, process the ending
before the new start.

##### Fast and slow pointers

**Mental model:** Move two positions through the same linked list. Their speeds
or the distance between them reveals a loop, the middle, or a node near the end.

**Recognize it:** A linked list may loop back to an earlier node, you need its
middle without first counting every node, or one position must remain a fixed
number of nodes behind another.

**Start here:** Decide how fast each position moves or how far apart they must
remain. Before following the link to the next node, check that the current node
exists. For a fixed distance, move the leading position ahead first and then
move both one node at a time.

**Invariant:** After setup, the pointers keep the same speed relationship or
the same fixed gap after every move.

**Reject it when:** You need to jump directly to any chosen position, remember
the full history, or return the actual visited path. Two moving positions do not
retain those details.

##### Trie

**Mental model:** Store words one character at a time. Words with the same
beginning reuse the same initial steps.

**Recognize it:** You repeatedly ask for words beginning with the same letters,
check dictionary words one character at a time, or want to stop as soon as a
beginning does not exist.

**Start here:** Each stored step records possible next characters and whether a
complete word ends there. To insert a word, follow or create one step for each
character. To look up a word, follow the same steps and check the end marker
after the last character.

**Invariant:** The characters followed from the starting point always spell the
beginning represented by the current step. A word exists only if the last step
is marked as the end of a complete word.

**Reject it when:** You only need to know whether a complete word exists. A hash
set is usually simpler and may use less memory.

#### Choosing among overlapping methods

- Sliding window versus prefix sums: use a window when moving its left or right
  boundary changes the tracked condition predictably, so an earlier starting
  position is never needed again. Use prefix sums when a range total can be
  calculated from two stored totals but boundary movement is not predictable.
- Two pointers versus binary search: use two pointers when both positions can
  move through the data together and each move permanently rules out some
  possibilities. Use binary search when testing the middle independently rules
  out half of the remaining possibilities.
- Greedy versus dynamic programming: use greedy when one choice can be proved
  not to reduce the best result still possible, and the same proof repeats
  after committing. Use dynamic programming when competing choices lead to
  different future states that must remain under consideration and those
  states recur.
- Dynamic programming versus backtracking: use dynamic programming when many
  choice sequences lead to the same remaining problem and you only need a
  total, minimum, maximum, or yes/no answer. Use backtracking when you must
  produce each individual result or when two paths with similar progress can
  still have different legal next choices.
- Backtracking versus DFS: DFS describes the order of exploration—follow one
  route fully before returning. Backtracking uses that order to build an answer
  through choices, then removes each choice before trying the next one. A
  backtracking solution therefore often performs a DFS, but not every DFS
  builds and undoes choices.
- DFS versus BFS: use DFS when you must finish one path, tree branch, or
  connected region before returning to alternatives. Use BFS when you need the
  nearest result or fewest moves and every move has the same cost.
- BFS versus Dijkstra: both explore outward from a starting location. BFS
  handles every waiting location in arrival order and gives cheapest routes
  only when every move costs the same. Dijkstra instead handles the location
  with the cheapest known route next, so it supports differing costs as long as
  none are negative. Dijkstra is therefore related to BFS, but it is not
  ordinary BFS.
- Dijkstra versus Bellman-Ford: both find cheapest routes from one starting
  location. Prefer Dijkstra when every connection cost is non-negative because
  it avoids repeatedly scanning every connection. Use Bellman-Ford when costs
  may be negative or when reachable negative-cycle detection is required.
- Two pointers versus fast and slow pointers: both coordinate two moving
  positions. The fast/slow version is a specialized two-pointer method for
  linked lists and similar structures, where different speeds or a fixed gap
  reveal a loop, a middle, or a position near the end.
- Heap versus sorting: sort when you need every item in order or all data is
  available at once. Use a heap when you repeatedly need only the next smallest
  or largest item, or when you must retain a fixed number of best items as new
  values arrive.

#### Let constraints eliminate bad choices

Let `n` mean the number of input items. The notation describes how quickly the
work grows: `O(n)` is roughly one pass, `O(n^2)` examines pairs, and `O(n^3)`
examines triples. These are rough interview-scale checks, not guarantees;
memory, programming language, and time limits still matter.

| Input scale | Methods worth considering |
| --- | --- |
| `n <= 10` | Trying every ordering (`n!` possibilities) may be feasible if impossible partial answers stop early. |
| `n <= 20` | Trying every subset (`2^n` possibilities) may be feasible. |
| `n <= 100` | Three nested passes (`O(n^3)`) may be possible; examining every pair (`O(n^2)`) is commonly reasonable. |
| `n <= 1,000` | Examining every pair (`O(n^2)`) may be possible but needs scrutiny. |
| `n <= 100,000` | Usually target sorting or a similar amount of work (`O(n log n)`), or a small number of full passes (`O(n)`). |
| `n >= 1,000,000` | Usually target a few full passes and avoid creating large extra arrays, maps, or sets. |

Also count the work required to produce the output. If the problem asks for all
`2^n` subsets or all `n!` orderings, merely writing those results already takes
that much work; generating them in a different order cannot remove it.

#### Before committing to a method

Write one sentence for each item:

- **Stored meaning:** What facts does each variable, table entry, queue item, or
  stack entry represent?
- **Invariant:** What fact remains true before and after every step?
- **Progress:** Why does each step reduce the work left to do or move closer to
  finishing?
- **Coverage:** Why can none of the choices you skipped or discarded contain an
  answer?
- **Work and memory:** How many smaller problems, positions, or elements are
  processed, how often, and what extra information is retained?

If one sentence cannot be made precise, the method is still only a guess.

## Go reference implementations

These snippets preserve the former concept packages in readable form. They are
reference material inside the book and are no longer compiled as repository
packages. The coverage notes record the important cases formerly enforced by
their unit tests.

### Binary search bounds

```go

package binarysearch

// LowerBound returns the first index whose value is greater than or equal to
// target. It returns len(values) when no such index exists.
func LowerBound(values []int, target int) int {
	left, right := 0, len(values)
	for left < right {
		middle := left + (right-left)/2
		if values[middle] < target {
			left = middle + 1
		} else {
			right = middle
		}
	}
	return left
}

// UpperBound returns the first index whose value is greater than target. It
// returns len(values) when no such index exists.
func UpperBound(values []int, target int) int {
	left, right := 0, len(values)
	for left < right {
		middle := left + (right-left)/2
		if values[middle] <= target {
			left = middle + 1
		} else {
			right = middle
		}
	}
	return left
}

// Search returns the index of target in a sorted slice, or -1 when target is
// absent. When duplicates exist, it returns the first matching index.
func Search(values []int, target int) int {
	index := LowerBound(values, target)
	if index == len(values) || values[index] != target {
		return -1
	}
	return index
}

```

**Former executable coverage:** Empty input, targets before and after the data, duplicate ranges, values between entries, first-match lookup, and absent targets.

### Rolling dynamic programming

```go

package dynamicprogramming

// CountOrderedOneOrTwoStepWays returns the number of ordered sequences of
// one-step and two-step moves that sum to steps. Negative totals have no valid
// sequence; the empty sequence is the single way to make a total of zero.
func CountOrderedOneOrTwoStepWays(steps int) int {
	if steps < 0 {
		return 0
	}
	if steps < 2 {
		return 1
	}

	twoBack, oneBack := 1, 1
	for current := 2; current <= steps; current++ {
		twoBack, oneBack = oneBack, twoBack+oneBack
	}
	return oneBack
}

```

**Former executable coverage:** Negative totals, the empty sequence, one and two steps, and several recurrence transitions.

### Hash-map complement lookup

```go

// Package hashmap demonstrates lookup-based uses of Go maps.
package hashmap

// PairSumIndices returns indices of two distinct values whose sum is target.
// The boolean is false when no pair exists.
func PairSumIndices(values []int, target int) (int, int, bool) {
	// Map every previously visited value to one of its earlier indices.
	indexByValue := make(map[int]int)
	// Visit each value once from left to right.
	for index, value := range values {
		// Calculate the value needed to complete the target.
		complement := target - value
		// Look up the complement and whether it exists as two explicit results.
		earlierIndex, found := indexByValue[complement]
		// Continue into this branch only when an earlier complement was found.
		if found {
			// Return both distinct indices and report that a pair was found.
			return earlierIndex, index, true
		}
		// Insert after lookup so the current element cannot match itself.
		indexByValue[value] = index
	}

	// Return the zero indices with false when the input contains no valid pair.
	return 0, 0, false
}

```

**Former executable coverage:** A previously seen complement, equal values at distinct indices, and inputs with no pair.

### Generic singly linked list

```go

package linkedlist

// Node stores one value and a link to the next node.
type Node[T any] struct {
	Value T
	Next  *Node[T]
}

// FromSlice returns a list containing values in their slice order.
func FromSlice[T any](values []T) *Node[T] {
	dummy := &Node[T]{}
	tail := dummy
	for _, value := range values {
		tail.Next = &Node[T]{Value: value}
		tail = tail.Next
	}
	return dummy.Next
}

// ToSlice returns the values encountered from head to nil.
func ToSlice[T any](head *Node[T]) []T {
	values := make([]T, 0)
	for node := head; node != nil; node = node.Next {
		values = append(values, node.Value)
	}
	return values
}

// Reverse reverses a list in place and returns its new head.
func Reverse[T any](head *Node[T]) *Node[T] {
	var previous *Node[T]
	for current := head; current != nil; {
		next := current.Next
		current.Next = previous
		previous = current
		current = next
	}
	return previous
}

```

**Former executable coverage:** Slice round trips, reversal, and empty-list construction, conversion, and reversal.

### Generic stack

```go

package stack

// Stack stores values in last-in, first-out order.
// The zero value is ready to use.
type Stack[T any] struct {
	items []T
}

// Push places value on top of the stack.
func (s *Stack[T]) Push(value T) {
	s.items = append(s.items, value)
}

// Pop removes and returns the top value. The boolean is false when the stack
// is empty.
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}

	last := len(s.items) - 1
	value := s.items[last]
	var zero T
	s.items[last] = zero
	s.items = s.items[:last]
	return value, true
}

// Peek returns the top value without removing it. The boolean is false when
// the stack is empty.
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}

	return s.items[len(s.items)-1], true
}

// Len returns the number of values in the stack.
func (s *Stack[T]) Len() int {
	return len(s.items)
}

// IsEmpty reports whether the stack contains no values.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

```

**Former executable coverage:** Last-in-first-out ordering, top inspection, length tracking, complete removal, and empty operations.

### Overflow-safe integer reversal

```go

package integerdigits

import "math"

// ReverseNonNegative reverses a nonnegative decimal integer. ok is false when
// value is negative or when its reversal would exceed int64.
func ReverseNonNegative(value int64) (reversed int64, ok bool) {
	if value < 0 {
		return 0, false
	}

	for value > 0 {
		digit := value % 10
		if reversed > (math.MaxInt64-digit)/10 {
			return 0, false
		}
		reversed = reversed*10 + digit
		value /= 10
	}
	return reversed, true
}

```

**Former executable coverage:** Zero, palindromes, trailing zeroes, negative rejection, and a reversal that would overflow int64.

### Common-prefix scan

```go

package commonprefixscan

// LongestCommonPrefix returns the longest byte prefix shared by all strings.
// It returns an empty string when strs is empty.
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	for column := 0; column < len(strs[0]); column++ {
		expected := strs[0][column]
		for row := 1; row < len(strs); row++ {
			if column >= len(strs[row]) || strs[row][column] != expected {
				return strs[0][:column]
			}
		}
	}
	return strs[0]
}

```

**Former executable coverage:** Shared and missing prefixes, an early shortest string, identical strings, a single string, and empty inputs.

### Directional scan

```go

package directionalscan

// SumWithSmallerLeftValuesSubtracted applies a local signed-sum rule: a value
// is subtracted when it is smaller than its immediate right neighbor and added
// otherwise.
func SumWithSmallerLeftValuesSubtracted(values []int) int {
	total := 0
	rightValue := 0

	for index := len(values) - 1; index >= 0; index-- {
		value := values[index]
		if value < rightValue {
			total -= value
		} else {
			total += value
		}
		rightValue = value
	}

	return total
}

```

**Former executable coverage:** Empty and single inputs, descending additive values, one subtractive pair, and mixed contributions.

### Sorted two-pointer pair search

```go

package twopointers

// PairSumSorted returns indices of two values whose sum equals target. Values
// must be sorted in ascending order, and relevant sums must fit in an int.
func PairSumSorted(values []int, target int) (left int, right int, ok bool) {
	left, right = 0, len(values)-1
	for left < right {
		sum := values[left] + values[right]
		switch {
		case sum < target:
			left++
		case sum > target:
			right--
		default:
			return left, right, true
		}
	}

	return 0, 0, false
}

```

**Former executable coverage:** A matching pair, negative values, distinct-position enforcement, a missing pair, and empty input.

## Go techniques

Keep algorithm code strongly typed and isolate serialization at the judge
boundary. Prefer table-driven tests for examples and edge cases, use
`sort.Slice` or a typed comparison for problem-specific order, implement
`container/heap` only when repeated priority updates justify it, and use type
parameters when a data structure's behavior is independent of its value type.

When replacing recursion with a loop, store exactly the state each unfinished
call would have retained. When working with slices, account for capacity,
occasional append reallocations, aliasing, and backing arrays kept alive by
small subslices.

## Further-study checklist

The book has complete cards for the methods linked above. Useful future
chapters include sorting and selection, binary-search trees, deques, advanced
range-query structures, divide and conquer, string matching, integer
combinatorics, custom fast input, and detailed `container/heap` templates.
