# Problem-Solving Patterns

Pattern pages explain how to recognize a useful method and why that method is
safe. That explanation usually centers on an **invariant**: a fact that remains
true while the algorithm runs.

Available:

- [Common-prefix scans](common-prefix-scan/)
- [Directional scans](directional-scan/)
- [Two pointers](two-pointers/)

## Problem-solving method selection guide

Treat recognition signals as possible clues, not keyword matches. An array does
not automatically imply two pointers, and the word "minimum" does not
automatically imply greedy. A method fits only when you can state the invariant
that makes its next step safe.

Use this guide as a sequence of progressively stronger hints. Start with the
questions and recognition guide. Open a practical pattern card only when the
guide leaves you unsure how to begin. Read a complete concept page or solution
explanation only after trying the card's first steps yourself.

### Start with these questions

1. What must the output describe: one value, an index, a path, an ordering, or
   every valid construction?
2. Must the chosen items be next to one another, or may the answer skip items
   or select them in any order?
3. Does the input already have useful order, or may it be sorted without losing
   information such as original indices or arrival order?
4. Can a possible choice or answer be proved impossible and permanently
   discarded?
5. Can a large version of the problem be built from smaller versions of the
   same problem? If so, what facts uniquely describe each smaller problem?
6. Can you represent the problem as items connected to other items, even when
   the prompt does not explicitly call it a graph?
7. How much work can the input size permit? Use the largest possible input to
   reject methods that would perform too many steps.

### Derive a direction before naming a pattern

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
2. **At 5 minutes:** use the recognition guide to identify at most two possible
   methods.
3. **At 10 to 15 minutes:** read the relevant practical pattern card and check
   whether its reason for discarding work applies to the problem.
4. **At 20 minutes:** trace the method on one small input before writing code.
5. **At 25 to 30 minutes:** read a full explanation if still blocked, close it,
   and implement the idea without copying its code.

### Recognition guide

Scan the method names, then read all four lines before choosing one. A signal
is only a clue; the method fits only when the explanation under **Why it fits**
matches the problem.

These groups identify the main question each method answers; they are not rigid
categories. Methods often work together. For example, Dijkstra's algorithm
uses a heap, and backtracking usually explores one path fully before returning
to try another, which is depth-first search (DFS) order.

#### Retrieve by key or prefix

Use these when the main operation is repeatedly finding information using an
exact value, something calculated from an item, or the beginning of a word.

- **[Hash map or set](../data-structures/hash-map/)**
  - **Signal:** You repeatedly ask whether a value has appeared, how often it
    appears, which earlier value would complete a pair, or which group an item
    belongs to.
  - **Why it fits:** When processing an item, store the fact you will need later
    under a value you can calculate again. A later item can calculate that same
    value and retrieve the fact without scanning all earlier items.
  - **Example:** For `[2, 7, 11, 15]` and target `9`, store `2` as you scan.
    When you reach `7`, ask the map whether its needed partner was stored:
    `9 - 7 = 2`.
  - **Common trap:** A map can find a stored entry quickly, but it does not keep
    its entries sorted or preserve a useful processing order by itself.

- **[Trie](#trie)**
  - **Signal:** You repeatedly search for words that begin with the same letters,
    check dictionary words one character at a time, or stop searching as soon
    as a beginning does not exist.
  - **Why it fits:** Words with the same beginning share the same path through
    their characters, so following `ca` keeps `car`, `cart`, and `cat` available
    without checking each whole word separately.
  - **Example:** Inserting `car`, `cart`, and `cat` stores the steps for `c` and
    `a` only once. Following the letters `ca` confirms that all three words are
    still possible.
  - **Common trap:** A trie may create many small records for possible
    next characters. If you only need to check complete words, a hash set is
    usually simpler and often smaller.

#### Scan a structure, narrow the possible answers, or summarize a range

Use these when you can avoid restarting from every position: move through the
input in a coordinated way, discard an impossible part of the search, or
calculate reusable totals in advance.

- **[Two pointers](two-pointers/)**
  - **Signal:** Sorted values, comparisons between two ends, pair selection,
    removing unwanted values without creating another array, or merging two
    sorted inputs.
  - **Why it fits:** After comparing the current values, you can prove that
    moving one pointer cannot discard a valid answer. For example, if a sorted
    pair sum is too small, moving the left pointer right is the only move that
    can increase it.
  - **Example:** In `[1, 2, 4, 7, 11]`, look for a pair totaling `9`. Since
    `1 + 11` is too large, move the right pointer left. Since `1 + 7` is too
    small, move the left pointer right and find `2 + 7 = 9`.
  - **Common trap:** Sorting may lose required original positions. Moving a
    pointer without a similar proof can skip the answer.

- **[Binary search](../algorithms/binary-search/)**
  - **Signal:** A target in sorted data; wrapped sorted data such as
    `[4, 5, 1, 2, 3]`; or the first or last position where a yes/no test changes
    its answer.
  - **Why it fits:** The possible positions split into two consecutive groups,
    such as "too small" followed by "large enough." Testing the middle tells
    you which half cannot contain the first acceptable position.
  - **Example:** In `[1, 3, 3, 5]`, find the first value at least `3`. Testing
    an index with value `3` proves the answer is there or farther left; testing
    value `1` proves the answer must be to its right.
  - **Common trap:** A numeric answer is not enough. Across the ordered list of
    possible answers, you need a test whose result changes only once.

- **[Sliding window](#sliding-window)**
  - **Signal:** A consecutive portion of an array or string must meet a
    condition or produce the largest or smallest result.
  - **Why it fits:** You can update the condition when one item enters or leaves
    the range. Once the left edge moves past a starting position, future items
    cannot make that discarded start useful again.
  - **Example:** For the longest consecutive sequence of characters without a
    repeat in `"abca"`, grow the window to `"abc"`. The next `a` repeats, so
    move the left edge past the earlier `a` and continue with `"bca"`.
  - **Common trap:** Negative values can make a sum rise or fall unexpectedly,
    so repeatedly shrinking the range while its condition fails may not work. A
    window also requires neighboring items; it cannot skip over items to form
    the answer.

- **[Prefix sums](#prefix-sums)**
  - **Signal:** Many queries ask for totals or counts over overlapping ranges,
    causing a direct solution to scan the same values repeatedly.
  - **Why it fits:** Store the total up to every position. Then the result for
    `[left, right]` is the total through `right` minus the total before `left`.
  - **Example:** For `[2, 4, 1, 3]`, prefix sums are `[0, 2, 6, 7, 10]`.
    The sum from indices `1` through `3` is `10 - 2 = 8`.
  - **Common trap:** Subtraction works for sums and counts, but not for every
    property. For example, you cannot find a range's smallest value by
    subtracting two previously stored smallest values. Also use a number type
    large enough to hold the accumulated total.

- **[Difference arrays](#difference-arrays)**
  - **Signal:** Many updates add or subtract a value across an entire range, and
    you need the final values after all updates are known.
  - **Why it fits:** Record only where an update begins and where it stops. One
    final left-to-right pass carries the change across every position in
    between.
  - **Example:** Start with four zeros and add `3` to indices `1` through `3`.
    Use one extra bookkeeping position: record `+3` at index `1` and `-3` at
    position `4`, just beyond the original array. Adding the recorded changes
    across original indices `0` through `3` produces `[0, 3, 3, 3]`.
  - **Common trap:** This is most useful when updates can be collected before
    producing the final array. It is less useful when every update must be
    followed immediately by queries about the changed values.

- **[Fast and slow pointers](#fast-and-slow-pointers)**
  - **Signal:** A linked list may loop back to an earlier node, you need its
    middle without first counting every node, or two positions must remain a
    fixed distance apart.
  - **Why it fits:** Moving pointers at different speeds makes them meet inside
    a loop, or lets one pointer reach the end when the other reaches the
    midpoint. Keeping a fixed gap can similarly identify a node near the end.
  - **Example:** In `1 -> 2 -> 3 -> 2 -> ...`, a slow pointer moves one node per
    step and a fast pointer moves two. Both eventually meet inside the loop. In
    a list without a loop, the fast pointer reaches the end instead.
  - **Common trap:** Before following the link to the next node, check that the
    current node exists. Also explain why a meeting or fixed distance proves
    the property you need.

- **[Monotonic stack](#monotonic-stack)**
  - **Signal:** The problem asks for the next greater or smaller value, the
    nearest value that becomes larger or smaller than the current value, or how
    long a value remains the minimum or maximum.
  - **Why it fits:** Earlier values whose answers are still unknown can be kept
    in increasing or decreasing order; this one-direction order is what
    "monotonic" means here. One new value may answer the question for several
    recent values, which are then removed together.
  - **Example:** In `[2, 1, 3]`, ask for the next value to the right that is
    greater. Store the positions of `2` and `1`; their values decrease from
    bottom to top. When `3` arrives, it answers the question for both, so remove
    both positions and record `3` as their answer.
  - **Common trap:** Be able to state exactly what question each stored item is
    waiting to answer. Do not remove an item if a later value may still need
    information from it.

#### Explore choices

Use these when solving the problem means choosing among alternatives and the
key question is which possibilities must remain available.

- **[Dynamic programming](../algorithms/dynamic-programming/)**
  - **Signal:** A solution that tries every choice keeps solving the same
    smaller problem again and again.
  - **Why it fits:** You can describe each smaller problem with a limited set of
    facts, store its answer once, and build larger answers from already solved
    ones.
  - **Example:** If you may climb one or two stairs, the number of ways to reach
    stair `4` is `ways[3] + ways[2]`. Store each earlier result instead of
    calculating the same earlier result again every time it is needed.
  - **Common trap:** Writing a formula that refers to smaller problems does not
    stop repeated work unless you store and reuse their answers.
    Conversely, storing facts that do not affect future choices can require too
    much memory.

- **[Greedy](#greedy)**
  - **Signal:** Choosing what looks best right now appears to produce the best
    possible final answer.
  - **Why it fits:** You can prove that if the best possible final answer begins
    with a different choice, replacing that first choice with yours does not
    make the result worse. Your choice therefore does not block the best result.
  - **Example:** To attend the most non-overlapping meetings from `(1, 3)`,
    `(2, 5)`, and `(4, 6)`, choose `(1, 3)` because it finishes first, then
    choose `(4, 6)`.
  - **Common trap:** Passing examples is not proof. If a choice that looks best
    now may need to be undone later, keep multiple possibilities with dynamic
    programming or search.

- **[Backtracking](#backtracking)**
  - **Signal:** You must generate every valid selection, ordering, grouping,
    placement, or path.
  - **Why it fits:** From any partial answer, you can list every legal next
    choice. After exploring one choice, you can undo it completely and explore
    the next one.
  - **Example:** To generate subsets of `[a, b]`, first choose `a`, then choose
    or skip `b`; undo `a`, then repeat the same choices for `b`. The resulting
    subsets are `[a, b]`, `[a]`, `[b]`, and `[]`.
  - **Common trap:** The number of answers may grow extremely quickly. Stop
    exploring a partial answer only when you can prove it cannot lead to a
    valid result.

#### Traverse relationships

Use these when the input is naturally a network of items and connections: for
example, a road map, tree, dependency list, grid, or collection of positions
connected by legal moves.

- **[Depth-first search](#depth-first-search)**
  - **Signal:** You need to know what can be reached, identify connected groups,
    explore complete paths, or combine results from the nodes below a tree node.
  - **Why it fits:** Following one path as far as possible lets you finish a
    complete path, everything below one tree node, or one connected region
    before returning to alternatives. Recording reached positions prevents a
    loop from sending the search around forever.
  - **Example:** To count islands in a grid, start DFS at one unvisited land
    cell and visit every connected land neighbor. When that search finishes,
    one entire island has been counted.
  - **Common trap:** Nested function calls use memory too. A very deep search
    can exceed the program's call limit and may need a stack collection managed
    by your code instead.

- **[Breadth-first search](#breadth-first-search)**
  - **Signal:** You need the fewest moves through a network where every move has
    the same cost, tree nodes grouped by depth, or the nearest reachable
    position.
  - **Why it fits:** A first-in, first-out queue processes every position one
    move away before any position two moves away. When every move has the same
    cost, the first visit therefore uses the fewest moves.
  - **Example:** In a maze, add the start at distance `0`, all reachable
    neighbors at distance `1`, and then their neighbors at distance `2`. The
    first time you reach the exit gives the fewest moves.
  - **Common trap:** Mark a position reached when adding it to the queue;
    waiting until removal may add the same position many times. If moves have
    different costs, use a method such as Dijkstra's algorithm instead.

- **[Topological sort](#topological-sort)**
  - **Signal:** Tasks have prerequisites, you need a valid work order, or you
    need to detect dependencies that form an impossible loop.
  - **Why it fits:** Each one-way connection says one task must happen before
    another. Repeatedly choosing a task with no unmet prerequisites produces a
    valid order.
  - **Example:** If `B` and `C` must finish before `D`, begin with `B` and `C` in
    either order. Only after both are processed does `D` have zero remaining
    prerequisites.
  - **Common trap:** Use ordinary DFS or BFS if you only need to know whether
    one item can reach another. If task dependencies form a loop, no valid task
    order exists.

- **[Dijkstra's algorithm](#dijkstra-shortest-paths)**
  - **Signal:** You need the cheapest route through a network, and every road or
    connection has a cost of zero or greater.
  - **Why it fits:** Always process the location with the cheapest known route
    from the start. Reaching it later through a route that already costs more,
    followed by a connection that cannot reduce cost, cannot improve it.
  - **Example:** Direct travel from `A` to `B` costs `4`. Travel from `A` to `C`
    costs `1`, and `C` to `B` costs `2`. Going through `C` lowers the known cost
    to `B` from `4` to `1 + 2 = 3`.
  - **Common trap:** A negative connection cost could later create a cheaper
    route to a location already processed, so this method is not valid then.

- **[Disjoint-set union](#disjoint-set-union)**
  - **Signal:** Connections are added one at a time, and you repeatedly need to
    know whether two items are already connected or how many separate groups
    remain.
  - **Why it fits:** Give every connected group one shared label. Comparing two
    labels answers whether the items are connected. When a connection joins two
    different groups, give the combined group one shared label.
  - **Example:** Towns `1`, `2`, and `3` initially form three separate groups.
    Adding roads `1-2` and `2-3` combines them into one group. Before adding
    road `1-3`, the structure reports that `1` and `3` are already connected,
    so the new road creates a loop.
  - **Common trap:** The structure remembers only which group each item belongs
    to. It cannot show the route between two items or find a shortest route.

#### Process items or events in a useful order

Use these when the answer becomes easier after choosing a specific processing
order: the most urgent item, the most recently added item, or the next range or
event by time or position.

- **[Heap or priority queue](#heap-or-priority-queue)**
  - **Signal:** You repeatedly need the smallest or largest item, a requested
    number of the best items, or the next scheduled event as new items arrive.
  - **Why it fits:** This structure keeps the next smallest or largest item easy
    to remove without spending time sorting every stored item after each change.
  - **Example:** To keep the three largest values seen so far, suppose you have
    `[4, 7, 9]` and then receive `10`. Remove the smallest retained value, `4`,
    and keep `[7, 9, 10]`.
  - **Common trap:** The structure guarantees only which item comes next; the
    remaining items are not necessarily in sorted order.

- **[Stack](../data-structures/stack/)**
  - **Signal:** Nested parentheses or brackets, undo/history behavior, or a
    calculation where the most recently opened or added item must be handled
    first.
  - **Why it fits:** The most recently added unfinished item is always available
    at the top, so it can be completed or undone before older items.
  - **Example:** To check the brackets in `([])`, place each opening bracket on
    top of the stack. When a closing bracket appears, remove the top opening
    bracket and verify that the pair matches.
  - **Common trap:** Use a queue instead when the oldest waiting item must be
    handled first.

- **[Intervals](#intervals)**
  - **Signal:** The input contains ranges such as meeting times, reservations,
    or numeric spans, and you need to combine overlaps or find gaps.
  - **Why it fits:** Sort the ranges by where they start. If the next range
    begins before the current one ends—or at the same point when touching counts
    as overlap—combine them. Once a gap appears, no remaining range can overlap
    the current one.
  - **Example:** After sorting `[1, 4]`, `[3, 6]`, and `[8, 10]`, compare the
    first two. Because `3 <= 4`, combine them into `[1, 6]`. Because `8 > 6`,
    save `[1, 6]` and begin a new range at `[8, 10]`.
  - **Common trap:** Decide whether touching ranges such as `[1, 2]` and
    `[2, 3]` count as overlapping in the problem.

- **[Sweep line](#sweep-line)**
  - **Signal:** Many ranges start and end at different positions, and you need
    to know how many are active at once or where that number changes.
  - **Why it fits:** Replace every range with a start event and an end event,
    then process those events in order. A running count increases at a start
    and decreases at an end.
  - **Example:** Suppose one booking starts at `1` and ends just before `4`, and
    another starts at `3` and ends just before `5`. The running count becomes
    `2` from time `3` until time `4`, so two rooms are needed then.
  - **Common trap:** Define what an end time means. If a booking ending at `4`
    frees its room before another booking starting at `4`, process that ending
    before the new start.

### Practical pattern cards

The cards below expand methods that do not yet have a complete concept page.
They are deliberately hints rather than finished solutions. A card should help
you produce and justify a first attempt without revealing problem-specific
code.

#### Sliding window

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

#### Prefix sums

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

#### Difference arrays

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

#### Greedy

**Mental model:** Make one choice now and never undo it because you can prove
that this choice does not prevent the best possible final result.

**Recognize it:** One available choice leaves at least as many useful options
for later as every competing choice. For example, a meeting that finishes
earlier leaves at least as much time for later meetings.

**Start here:** State the immediate choice in one sentence. Now imagine a best
possible answer that starts differently. Replace its first choice with yours.
If the replacement remains valid and does not reduce the result, your choice is
safe.

**Invariant:** After each commitment, at least one best possible final answer
still begins with all choices made so far.

**Reject it when:** A choice that looks best now can block a better combination
later, and you cannot prove that the replacement described above is harmless.
Keep multiple possibilities with dynamic programming or search instead.

#### Backtracking

**Mental model:** Build one possible answer one choice at a time. Try everything
that can follow the latest choice, remove that choice, and try the next option.

**Recognize it:** You must produce every valid selection, ordering, grouping,
path, or placement, or the total number of possibilities is small enough to try
them all.

**Start here:** Define the answer built so far, every legal next choice, and the
condition that makes an answer complete. Then repeat three actions: add one
choice, explore what can follow it, and remove that choice before trying the
next option.

**Invariant:** While exploring one possibility, the working data contains
exactly the choices made for that possibility. Removing the latest choice must
restore the data to what it was before that choice.

**Reject it when:** The input is too large to explore every possibility and many
different paths lead to the same remaining problem. Store and reuse the answer
to that repeated problem with dynamic programming instead.

#### Depth-first search

**Mental model:** Follow one route as far as possible. When that route ends,
return to the most recent place that had another route available.

**Recognize it:** You need to visit an entire connected region, inspect every
complete path, or calculate a tree result from the results below each node. You
do not need the first discovered path to use the fewest moves.

**Start here:** Define one position and how to list every position reachable in
one move. Mark a position reached before exploring from it so a loop cannot send
the search around forever. Use nested function calls or your own stack to
remember where to return.

**Example:** To count islands in a grid, begin at one unvisited land cell and
continue through neighboring land until no more can be reached. That completed
search accounts for one island.

**Invariant:** Every position marked reached has either been fully explored or
is waiting on the current return path. No position needs to be scheduled twice.

**Common trap:** Nested function calls consume memory. A path with too many
steps may exceed the program's call limit, so use your own stack collection for
very deep inputs.

#### Breadth-first search

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

#### Topological sort

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

#### Dijkstra shortest paths

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

#### Disjoint-set union

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

#### Heap or priority queue

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

#### Monotonic stack

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

#### Intervals

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

#### Sweep line

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

#### Fast and slow pointers

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

#### Trie

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

### Choosing among overlapping methods

- Sliding window versus prefix sums: use a window when moving its left or right
  boundary changes the tracked condition predictably, so an earlier starting
  position is never needed again. Use prefix sums when a range total can be
  calculated from two stored totals but boundary movement is not predictable.
- Two pointers versus binary search: use two pointers when both positions can
  move through the data together and each move permanently rules out some
  possibilities. Use binary search when testing the middle independently rules
  out half of the remaining possibilities.
- Greedy versus dynamic programming: use greedy only when you can prove that
  committing to the current choice cannot prevent the best possible answer.
  Use dynamic programming when several current choices must be kept because
  they can lead to different future results.
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
- Two pointers versus fast and slow pointers: both coordinate two moving
  positions. The fast/slow version is a specialized two-pointer method for
  linked lists and similar structures, where different speeds or a fixed gap
  reveal a loop, a middle, or a position near the end.
- Heap versus sorting: sort when you need every item in order or all data is
  available at once. Use a heap when you repeatedly need only the next smallest
  or largest item, or when you must retain a fixed number of best items as new
  values arrive.

### Let constraints eliminate bad choices

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

### Before committing to a method

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

See the complete [concept curriculum](../).
