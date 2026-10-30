# Algorithms and Data Structures in Go

This guide explains algorithms and ways to organize data in Go. It is
for readers who know basic Go, including variables, loops, functions,
slices, and maps, but are still learning how to choose a solution and
explain why it works.

The files under `problems/` contain complete solutions to individual
problems. This guide explains the ideas that can be reused across many
problems. You do not need to read it from beginning to end before
solving anything.

## Contents

- [How to use this guide](#how-to-use-this-guide)
- [Contemporary interview evidence](#contemporary-interview-evidence)
- [A. Foundations](#a-foundations)
  - [A.1 Complexity and resource
    accounting](#a1-complexity-and-resource-accounting)
  - [A.2 Correctness and invariants](#a2-correctness-and-invariants)
  - [A.3 Recursion and saved work](#a3-recursion-and-saved-work)
  - [A.4 Integer limits and bits](#a4-integer-limits-and-bits)
  - [A.5 Bit manipulation](#a5-bit-manipulation)
  - [A.6 Keeping arithmetic values exact and
    consistent](#a6-keeping-arithmetic-values-exact-and-consistent)
  - [A.7 Combinatorics and output
    size](#a7-combinatorics-and-output-size)
  - [A.8 Integer digit processing](#a8-integer-digit-processing)
- [B. Ways to store and organize
  information](#b-ways-to-store-and-organize-information)
  - [B.1 Arrays, slices, and strings](#b1-arrays-slices-and-strings)
  - [B.2 Linked lists](#b2-linked-lists)
  - [B.3 Queues and deques](#b3-queues-and-deques)
  - [B.4 Trees and binary-search
    trees](#b4-trees-and-binary-search-trees)
  - [B.5 Graphs and grids](#b5-graphs-and-grids)
  - [B.6 Matrix traversal and in-place
    state](#b6-matrix-traversal-and-in-place-state)
  - [B.7 Composing data structures](#b7-composing-data-structures)
  - [B.8 Streaming and online
    processing](#b8-streaming-and-online-processing)
  - [B.9 Choosing a structure](#b9-choosing-a-structure)
- [C. Ways to solve problems](#c-ways-to-solve-problems)
  - [C.1 Problem-solving method selection
    guide](#c1-problem-solving-method-selection-guide)
    - [C.1.1 Ask what the input and output allow (quick decision
      tree)](#c11-quick-decision-tree)
    - [C.1.2 Derive a direction before naming a
      pattern](#c12-derive-a-direction-before-naming-a-pattern)
    - [C.1.3 Pattern cards](#c13-pattern-cards)
    - [C.1.4 Choosing among overlapping
      methods](#c14-choosing-among-overlapping-methods)
    - [C.1.5 Let constraints eliminate bad
      choices](#c15-let-constraints-eliminate-bad-choices)
    - [C.1.6 Before committing to a
      method](#c16-before-committing-to-a-method)
- [D. Go techniques](#d-go-techniques)
  - [D.1 Keep input conversion outside the
    algorithm](#d1-keep-input-conversion-outside-the-algorithm)
  - [D.2 Table-driven tests](#d2-table-driven-tests)
  - [D.3 Sorting without unexpected
    mutation](#d3-sorting-without-unexpected-mutation)
  - [D.4 Map presence and sets](#d4-map-presence-and-sets)
  - [D.5 Priority queues](#d5-priority-queues)
  - [D.6 Bytes, runes, and visible
    characters](#d6-bytes-runes-and-visible-characters)
  - [D.7 Type parameters](#d7-type-parameters)
  - [D.8 Replacing recursion with a
    loop](#d8-replacing-recursion-with-a-loop)
  - [D.9 Slice storage and retained
    memory](#d9-slice-storage-and-retained-memory)
- [E. Go reference implementations](#e-go-reference-implementations)
  - [E.1 Binary search bounds](#e1-binary-search-bounds)
  - [E.2 Rolling dynamic programming](#e2-rolling-dynamic-programming)
  - [E.3 Hash-map complement lookup](#e3-hash-map-complement-lookup)
  - [E.4 Generic singly linked list](#e4-generic-singly-linked-list)
  - [E.5 Generic stack](#e5-generic-stack)
  - [E.6 Overflow-safe integer
    reversal](#e6-overflow-safe-integer-reversal)
  - [E.7 Common-prefix scan](#e7-common-prefix-scan)
  - [E.8 Directional scan](#e8-directional-scan)
  - [E.9 Sorted two-pointer pair
    search](#e9-sorted-two-pointer-pair-search)
- [F. Further study](#f-further-study)

## How to use this guide

Choose the path that matches what you need now:

- **Learning from the beginning:** Read foundations, then ways to
  store information, then ways to solve problems. Follow a link to a
  Go technique when you need the language detail.
- **Solving a problem:** Begin with the [method-selection
  guide](#c1-problem-solving-method-selection-guide). Describe the
  input and required answer. Then compare the one or two methods that
  best fit those requirements.
- **Looking something up:** Jump to its entry, work through the
  examples, and read "When not to use it" before using the code.

Main parts use letters, and nested sections add numbers. For example,
`C` identifies problem-solving methods, `C.1` identifies the
method-selection guide, and `C.1.3` identifies the pattern-card
collection within that guide.

### Plain English before technical names

The guide keeps names such as *binary search*, *dynamic programming*,
and *Big-O* because those names appear in problems, documentation, and
interviews. It first explains the action or behavior in ordinary
language, then introduces and defines the conventional name. A
technical term should never be required to understand its own
definition.

### What every teaching entry contains

Every teaching entry uses these headings in this order. A heading can
be brief, but it is not omitted merely because the topic is small:

1.  **The problem it solves** states the input and desired result
    before naming the technique as the answer.
2.  **How it works** builds one clear mental model in ordinary
    language. It uses a familiar action when that makes the idea
    easier to picture, then explains every rule or distinction needed
    to understand the example without filler.
3.  **Step-by-step examples** teach the procedure with an ordinary
    input and, when behavior changes, a boundary or deliberately
    tricky input.
4.  **Why it works** uses plain language to identify what remains
    true, why each step is safe, why the process finishes, and why the
    final state gives the requested result, without filler.
5.  **Go example** shows the smallest complete fragment or links to
    the complete reference implementation elsewhere in the guide.
6.  **Work and memory** explains the costs in words before using
    notation such as `O(n)`.
7.  **When to use it** lists clues visible in the problem or input.
8.  **When not to use it** names the assumption or cost that rules it
    out.
9.  **Common mistakes and edge cases** ties every warning to a
    concrete input, state change, or Go behavior.
10. **Related material** links to prerequisite, contrasting, and
    implementation details instead of repeating them.

The labels are Markdown headings one level below the entry heading.
For example, an `###` entry uses `#### The problem it solves`; a
deeply nested `#####` pattern card uses
`###### The problem it solves`.

### Forms used in every entry

- Trace: Show one `Initial state`, then use a numbered list of
  `Decision`, `After`, and `Known correct` transitions. Each `After`
  is the next step's input; do not repeat it as `Before`. Use concise
  bullets for comparisons or mappings. Do not create Markdown tables.
- **Warning:** Write `> **Warning:**` followed by the failing input or
  state and the correction. A warning without a concrete failure is
  incomplete.
- **Code:** Use a fenced `go` block. Names in the code match the
  trace, and any input requirement appears immediately before the
  block. When the complete implementation has another owner, link to
  it rather than copying it.
- **Related links:** End with a short bulleted list. Each item says
  why the linked entry matters; it does not duplicate that entry's
  explanation.

The contemporary interview evidence below gives context for what to
practice; it is not a teaching entry, so it does not use those ten
headings.

The list of topics is checked against LeetCode's [Top Interview
150](https://leetcode.com/studyplan/top-interview-150/), the [Grind 75
study guidance](https://www.techinterviewhandbook.org/grind75/faq),
and the [NeetCode roadmap](https://neetcode.io/roadmap). These lists
help reveal missing common topics. The guide is organized by reusable
ideas rather than by the title of a practice problem.

## Contemporary interview evidence

Official preparation guides from
[Amazon](../REFERENCE_LINKS.md#amazon-interview-topics),
[Meta](../REFERENCE_LINKS.md#meta-interview-guide),
and
[Microsoft](../REFERENCE_LINKS.md#microsoft-technical-interviews)
all describe skills such as understanding a problem, choosing and
explaining an approach, writing clear code, and checking the result.
They are better guides to preparation than a remembered list of exact
questions.

Public candidate reports can show how one person's interview differed
from the official outline. Examples include reports from [Amazon in
2026](../REFERENCE_LINKS.md#amazon-2026-report),
[Google in 2025](https://leetcode.com/discuss/post/6522571/),
[Meta in 2025](../REFERENCE_LINKS.md#meta-2025-report),
[Microsoft in 2025](https://leetcode.com/discuss/post/7227003/), and
[Uber in 2025](https://leetcode.com/discuss/post/7327020/). Candidate
reports describe individual experiences. They do not prove that a
question is common or that another candidate will receive it. The
useful lessons are therefore general:

- A familiar problem may change one input rule. It may allow duplicate
  values, limit the number of edits, deliver values one at a time, or
  require less memory.
- Start with a direct solution that is correct. Then find work it
  repeats and explain why the improved solution can avoid that work.
- Sometimes the input does not look like a graph at first. People,
  roads, dependencies, or cells may still form nodes and connections.
- A solution may use two data structures together. State the rule that
  keeps them in agreement after every change.
- Check the code by tracing an ordinary input, a smallest or boundary
  input, and an input designed to expose a likely mistake.

Exact company questions change and may be confidential. Use company
tags to choose extra practice, not as a substitute for understanding
the ideas below.

## A. Foundations

### A.1 Complexity and resource accounting

#### The problem it solves

Given two correct approaches and the largest allowed input, determine
whether each approach can finish within the time and memory limits.
The comparison must describe how resource use grows with the input,
not how fast one run happens to be on one computer. It must also count
hidden costs such as recursion, copied output, and preprocessing.

#### How it works

Finding a name by reading every page of a sorted phone book gets much
slower as the book grows. Opening near the middle and discarding the
half that cannot contain the name needs only one more decision each
time the book's size doubles. **Big-O notation** describes how work or
memory grows with an input amount such as `n` values. **Auxiliary
space** is the extra memory the method needs beyond its input and
required output.

#### Step-by-step examples

Search for `9` in the sorted values `[1,3,5,7,9,11,13,15]`.
The indexes run from `0` through `7`. Both methods should return
`true`; the scan returns `(true,5)` because it checks five values.

Initial state: scan has `index=0`, `comparisons=0`,
`result=unknown`.

1.  Scan through the failed values
    - Decision: Compare `1`, `3`, `5`, and `7` with `9` in that
      order. Each differs, so advance after each comparison.
    - After: `index=4`, `comparisons=4`, `result=unknown`.
    - Known correct: Indexes `0` through `3` are not the answer.
2.  Finish the scan
    - Decision: `values[4]==9`; stop.
    - After: `index=4`, `comparisons=5`, `result=true`.
    - Known correct: Index `4` contains the target; this method
      made 5 comparisons. A miss could require all 8.

Initial state: halving has `left=0`, `right=7`, `comparisons=0`,
`result=unknown`.

1.  Check midpoint `3`
    - Decision: `values[3]=7 < 9`; discard indexes `0` through `3`.
    - After: `left=4`, `right=7`, `comparisons=1`,
      `result=unknown`.
    - Known correct: No discarded index can contain `9`; this
      method has made 1 comparison.
2.  Check midpoint `5`
    - Decision: `values[5]=11 > 9`; discard indexes `5` through `7`.
    - After: `left=4`, `right=4`, `comparisons=2`,
      `result=unknown`.
    - Known correct: No discarded index can contain `9`; this
      method has made 2 comparisons.
3.  Check midpoint `4`
    - Decision: `values[4]=9`; match and stop.
    - After: `left=4`, `right=4`, `comparisons=3`, `result=true`.
    - Known correct: Index `4` contains the target; this method
      made 3 comparisons.

The worst-case misses make the different growth visible:

- With `values=1` through `8` and `target=9`, a scan checks all
  `8` values. The same halving rule needs at most `4` checks.
- With `values=1` through `16` and `target=17`, the scan doubles to
  `16` checks. The halving search needs at most `5`, only one more.

For empty input, a scan performs zero element comparisons and a
halving search also stops immediately. That best case does not change
the scan's worst-case `O(n)` bound or the halving search's worst-case
`O(log n)` bound.

#### Why it works

The count follows the program's control flow. A scan advances past at
most one of `n` positions per comparison, so it performs no more than
`n` comparisons. A halving search leaves at most half of the previous
candidates after every failed comparison. The number of times `n` can
be halved before nothing remains is written `log n`, so the bound is
`O(log n)`.

Big-O keeps the part that controls growth for large inputs. If a
method performs `3n + 20` operations, doubling `n` roughly doubles the
work once `n` is large. The fixed multiplier `3` and fixed extra `20`
do not change that linear growth, so the method is `O(n)`.

#### Go example

Calling `containsWithComparisonCount` with the traced values and
`target=9` returns `(true,5)`. The function makes the counted
operation--one equality check per visited element--visible:

```go
func containsWithComparisonCount(
	values []int,
	target int,
) (bool, int) {
	comparisons := 0
	// Visit values from left to right. Each iteration performs the
	// equality check whose count determines this linear scan's running
	// time.
	for index := range values {
		comparisons++
		if values[index] == target {
			// A match is enough to answer the question, so return both the
			// result and the work performed.
			return true, comparisons
		}
	}
	// Reaching the end means every element was checked without a match.
	return false, comparisons
}

func containsByHalvingWithComparisonCount(
	values []int,
	target int,
) (bool, int) {
	// Every possible match remains inside the inclusive boundaries.
	left, right := 0, len(values)-1
	comparisons := 0
	for left <= right {
		middle := left + (right-left)/2
		comparisons++
		switch {
		case values[middle] < target:
			// Sorted values through middle are too small.
			left = middle + 1
		case values[middle] > target:
			// Sorted values from middle onward are too large.
			right = middle - 1
		default:
			// Equality proves the target exists at middle.
			return true, comparisons
		}
	}
	// Crossing boundaries prove that no candidate index remains.
	return false, comparisons
}
```

Calling the halving function with the same values and target returns
`(true,3)`, matching the second trace.

#### Work and memory

Here `n` is the number of values. If the target is first, the function
performs one comparison: best-case `O(1)` time. A missing or last
target requires `n` comparisons, so the worst-case time is `O(n)`.
A successful search can stop anywhere between those cases. Each
function stores only indexes and a comparison count, so extra memory
is `O(1)`. The input slice is borrowed and no output collection is
allocated.

#### When to use it

Use growth analysis whenever choosing among solutions, checking a
constraint, or reviewing a hidden allocation. Define the input measure
explicitly when there are several dimensions, such as rows `r`,
columns `c`, and output count `k`.

#### When not to use it

Do not use a growth class as an exact speed claim for small inputs.
Constants, cache behavior, allocation, and input distribution can make
two methods with the same class perform differently. Measure real
performance after the growth rate is acceptable.

#### Common mistakes and edge cases

> **Warning:** Two consecutive loops over `n` values cost `n + n`, or
> `O(n)`, not `O(n^2)`. The quadratic result applies when one full
> loop runs inside every step of another.

> **Warning:** A recursive function with no explicit slice can still
> use memory. `n` unfinished calls require `O(n)` call-stack space.

Count sorting or map construction before a later fast query, and count
copied output elements when output is part of the required work. For
rectangular data, do not call an `r * c` traversal `O(n)` unless `n`
has been defined as the total cell count.

#### Related material

- [A.2 Correctness and invariants](#a2-correctness-and-invariants)
  separates proving the answer from measuring its cost.
- [A.3 Recursion and saved work](#a3-recursion-and-saved-work)
  explains how recursive calls use stack memory.
- [A.7 Combinatorics and output
  size](#a7-combinatorics-and-output-size) shows why required output
  can dominate every other cost.
- [E.1 Binary search bounds](#e1-binary-search-bounds) provides the
  complete halving-search implementation used in the comparison.

### A.2 Correctness and invariants

#### The problem it solves

Given an algorithm, its valid inputs, and its required output, decide
whether it returns the right answer for every valid input rather than
only for the samples. The reasoning must explain what each remembered
value means, why every state-changing step preserves that meaning,
and why the process finishes without skipping a possible answer.

#### How it works

Imagine tidying toys from left to right while keeping this promise:
"everything left of my hand is already in the right place." The
promise is a **fact that remains true** while the work proceeds. The
formal name for such a fact is an **invariant**. A useful invariant
describes exactly which part of the result is already correct.

#### Step-by-step examples

Move the non-zero values in `values = [0, 5, 0, 2]` to the front
without changing their order. The expected returned prefix is
`[5, 2]`. `read` is the position being inspected, and `write` is the
length of the correct prefix. The slice expression `values[:write]`
starts at index `0` and stops before index `write`, so it contains
exactly that finished prefix.

Initial state: `values=[0,5,0,2]`, `read=0`, `write=0`,
`values[:write]=[]`.

1.  Read value `0`
    - Decision: Skip `0` because it is not part of the result.
    - After: `values=[0,5,0,2]`, `read=1`, `write=0`,
      `values[:write]=[]`.
    - Known correct: The correct prefix is still empty.
2.  Read value `5`
    - Decision: Write `5` at position `0`, then advance `write`.
    - After: `values=[5,5,0,2]`, `read=2`, `write=1`,
      `values[:write]=[5]`.
    - Known correct: `[5]` is the non-zero input seen
      so far, in order.
3.  Read value `0`
    - Decision: Skip `0`.
    - After: `values=[5,5,0,2]`, `read=3`, `write=1`,
      `values[:write]=[5]`.
    - Known correct: The correct prefix remains `[5]`.
4.  Read value `2`
    - Decision: Write `2` at position `1`, then advance `write`.
    - After: `values=[5,2,0,2]`, `read=4`, `write=2`,
      `values[:write]=[5,2]`.
    - Known correct: `[5,2]` is the complete result.

For the boundary input `[0, 0]`, both steps skip their value. `write`
remains `0`, so the returned prefix is `values[:0]`, or `[]`, as
required. The invariant describes the empty result without a special
case.

#### Why it works

Before the loop, no values have been read, so the empty prefix
contains exactly the non-zero values seen so far. During a step, a
zero is safely skipped because it does not belong in the result. A
non-zero value is written immediately after the already-correct
prefix, so the prefix gains that value without changing the order of
earlier values. The fact therefore remains true after either action.

`read` advances once per input value and eventually reaches the fixed
input length, so the loop finishes. At that point every input value
has been read; the invariant says `values[:write]` contains every
non-zero input value in its original order, which is the requested
output.

#### Go example

The function changes the supplied slice and returns a shorter view of
the same underlying storage:

```go
func compactNonZero(values []int) []int {
	// values[:write] is the compacted prefix. It starts empty because no
	// input elements have been inspected yet.
	write := 0
	for read := range values {
		value := values[read]
		if value == 0 {
			// Zeros do not belong in the result, so only the read position
			// moves.
			continue
		}
		// Copy the next retained value directly after the compacted prefix,
		// then grow that prefix by one. Reading never falls behind writing.
		values[write] = value
		write++
	}
	// Return a slice view of only the compacted prefix; no new array is
	// made.
	return values[:write]
}
```

#### Work and memory

The loop reads each of the `n` input values once in every case, so its
work is linear: `O(n)`. It keeps one index and one current value, so
it uses `O(1)` extra memory. The returned slice is a new slice header
over the existing array; the function does not allocate an output
array.

#### When to use it

Use an invariant to reason about a loop, recursion, or data-structure
update when the code builds an answer gradually, changes state, or
discards possible choices. It is especially useful when a problem asks
you to prove that order, membership, a boundary, or a running total
remains valid.

#### When not to use it

Do not force a loop-style invariant onto a direct calculation whose
result follows immediately from a defined operation. A method that
allows uncertainty or approximation may need a probability or error
limit instead of an absolute claim.

#### Common mistakes and edge cases

> **Warning:** "The processed part is correct" does not say what
> *correct* means. For `[0, 5]`, state that `values[:write]` equals
> the non-zero values read so far, in order.

> **Warning:** Check the invariant after both branches. If the zero
> branch increments `write`, input `[0]` would return a one-value
> prefix even though the expected result is empty.

Empty input is safe: the loop performs no steps, the initial invariant
already describes the whole input, and `values[:0]` is empty. A proof
must also name a progress measure; an invariant alone does not show
that a loop terminates.

#### Related material

- [A.1 Complexity and resource
  accounting](#a1-complexity-and-resource-accounting) explains how to
  account for the work after correctness is established.
- [B.1.1 Moving values in place](#b11-moving-values-in-place)
  explains the slice operation demonstrated here.
- [Remove
  Element](problems/solved_easy_leetcode_0027_remove_element.go)
  applies the same prefix invariant to a judge problem.

### A.3 Recursion and saved work

#### The problem it solves

Given nested input such as directories inside directories, compute a
whole answer from answers for the smaller nested parts. The method
must remember unfinished outer work while it solves an inner part,
then resume that work when the smaller result is available.

#### How it works

Open a box inside another box and leave a note in the outer one about
what to do when you return. **Recursion** works the same way: each
call solves a strictly smaller problem, while a **stack frame**
remembers the paused work. A **base case** answers the smallest
problem without another call, so the chain can stop. The number of
waiting frames is the **recursion depth**.

#### Step-by-step examples

Count files in this directory tree; expect `7`:

```text
root: 1 direct file
|-- photos: 2 direct files
`-- work: 1 direct file
    `-- drafts: 3 direct files
```

Each call starts `totalFiles` at its directory's direct-file count.
A waiting frame records that total and which child comes next.

Initial state: `active=root(start)`, `waiting=[]`, `returned=none`.

1.  Enter `photos`
    - Decision: Set root's `totalFiles=1`; save root's unfinished
      `work` child, then call `photos`.
    - After: `active=photos(start)`,
      `waiting=[root(totalFiles=1,next=work)]`, `returned=none`.
    - Known correct: Root's one direct file has been counted once.
2.  Return from `photos`
    - Decision: Set `totalFiles=2`. It has no children, so return
      `2`.
    - After: `active=root(totalFiles=1,next=work)`, `waiting=[]`,
      `returned=2`.
    - Known correct: The complete `photos` subtree contains `2`
      files.
3.  Enter `work`
    - Decision: Add the returned `2`, making root's
      `totalFiles=3`; save root, then call `work`.
    - After: `active=work(start)`,
      `waiting=[root(totalFiles=3,next=done)]`, `returned=none`.
    - Known correct: Root's direct file and `photos` subtree account
      for `3` files.
4.  Enter `drafts`
    - Decision: Set work's `totalFiles=1`; save work, then call its
      `drafts` child.
    - After: `active=drafts(start)`,
      `waiting=[root(totalFiles=3,next=done),`
      `work(totalFiles=1,next=done)]`, `returned=none`.
    - Known correct: Work's one direct file has been counted once.
5.  Return from `drafts`
    - Decision: Set `totalFiles=3`. It has no children, so return
      `3`.
    - After: `active=work(totalFiles=1,next=done)`,
      `waiting=[root(totalFiles=3,next=done)]`, `returned=3`.
    - Known correct: The complete `drafts` subtree contains `3`
      files.
6.  Return from `work`
    - Decision: Add the returned `3`, making work's
      `totalFiles=4`; no children remain, so return `4`.
    - After: `active=root(totalFiles=3,next=done)`, `waiting=[]`,
      `returned=4`.
    - Known correct: The complete `work` subtree contains `4`
      files.
7.  Return from `root`
    - Decision: Add the returned `4`, making root's
      `totalFiles=7`; no children remain, so return `7`.
    - After: `active=none`, `waiting=[]`, `returned=7`.
    - Known correct: All four directories have contributed their
      files exactly once, so the full tree contains `7` files.

A nil directory returns `0` without creating another call. A
directory with no files and no children also returns `0` after its
child loop performs no steps.

#### Why it works

A nil directory contributes no files. For a non-nil directory,
`totalFiles` starts with exactly its direct files. Assuming a child
call returns the complete count for that child subtree, adding the
returned count preserves the meaning: `totalFiles` covers the direct
files and every finished child. When no children remain, it is the
count for the entire current subtree.

Every call moves to a child in a finite, acyclic directory tree, so
the process reaches a directory with no children and eventually
finishes. Each unfinished parent must remain available until its child
returns, which explains both the return order and the stack memory.

#### Go example

```go
type directoryNode struct {
	NumDirectFiles   int
	ChildDirectories []*directoryNode
}

func countFiles(directory *directoryNode) int {
	// A missing directory contributes nothing and ends this branch.
	if directory == nil {
		return 0
	}
	// Begin with the files stored directly here. Each child call pauses
	// this frame, returns its complete subtree count, and is added once.
	totalFiles := directory.NumDirectFiles
	for _, childDirectory := range directory.ChildDirectories {
		totalFiles += countFiles(childDirectory)
	}
	return totalFiles
}
```

#### Work and memory

Let `d` be the number of directories and `h` the greatest number of
directories on a root-to-leaf path. The function visits each directory
and each parent-child link once, so it does `O(d)` work for a tree. At
most `h` calls are unfinished together, so the call stack uses
`O(h)` extra memory: `O(d)` for a one-child chain. It creates no
result collection.

#### When to use it

Use recursion when the input is naturally nested, each child asks for
the same kind of answer as its parent, and the maximum depth is safely
bounded. Directory trees, syntax trees, split ranges, and choice
exploration often have this shape.

#### When not to use it

Do not use unbounded recursion when a long chain can exhaust the call
stack or when a simple loop expresses the same state more directly.
Use an explicit stack when pending work must be inspected, limited, or
processed in a chosen order.

#### Common mistakes and edge cases

> **Warning:** Recursion needs input that eventually reaches a base
> case. If `work` contains `root` as a child in the example, calls
> alternate forever until the stack fails. Validate that links form a
> tree or record visited directories when cycles are allowed.

> **Warning:** Count paused calls as memory. A chain of `100,000`
> one-child directories may exhaust the call stack even though the
> function allocates no explicit stack slice.

Repeated recursive calls can also redo the same smaller problem. Save
repeated answers as described by [dynamic
programming](#c1310-dynamic-programming), or replace a one-branch
rule with a loop.

#### Related material

- [A.1 Complexity and resource
  accounting](#a1-complexity-and-resource-accounting) explains how
  to count extra memory.
- [B.4.3 Recursive return contracts](#b43-recursive-return-contracts)
  explains how a parent combines returned child answers.
- [C.1.3.11 Divide and conquer](#c1311-divide-and-conquer) uses
  several smaller independent calls.
- [D.8 Replacing recursion with a
  loop](#d8-replacing-recursion-with-a-loop) shows how to store
  pending work explicitly.

### A.4 Integer limits and bits

#### The problem it solves

Given Go integers whose types have fixed minimum and maximum values,
either produce an arithmetic result or report that it cannot fit in
the chosen type. The check must happen before the overflowing add,
subtract, or multiply, because the intended value is already lost
after overflow occurs.

#### How it works

An integer type is a fixed-size odometer. A signed 8-bit odometer
represents only `-128` through `127`; the mathematical value `128` has
no valid display. Trying to produce a value outside the type's range
is **overflow**. The binary digits stored in that fixed width are
called **bits**.

#### Step-by-step examples

Use a teaching limit of `127` and append digit `8`.

Initial state: `value=11`, `digit=8`, `limit=127`.

1.  Accept the safe value
    - Decision: `11 <= (127-8)/10`, or `11 <= 11`.
    - After: `result=118`, `valid=true`.
    - Known correct: The result is at most `127`.

Initial state: `value=12`, `digit=8`, `limit=127`.

1.  Reject the unsafe value
    - Decision: `12 > (127-8)/10`, or `12 > 11`.
    - After: `result=0`, `valid=false`; multiplication did not run.
    - Known correct: No unrepresentable result was evaluated.

At the boundary `value=12`, digit `7` is accepted because
`12 <= (127-7)/10` and the result is exactly `127`.

#### Why it works

Starting from the required inequality `value*10 + digit <= limit`,
subtract the nonnegative digit and divide by positive `10`. Integer
division gives the largest whole `value` that can be multiplied
safely, so the equivalent check is `value <= (limit-digit)/10`.
Performing the comparison first avoids the very overflow it is meant
to detect.

#### Go example

```go
func appendDecimalDigit(value, digit, limit int64) (int64, bool) {
	// Reject values outside the nonnegative decimal-append contract
	// before using them in arithmetic.
	if value < 0 || digit < 0 || digit > 9 || limit < digit {
		return 0, false
	}
	// This rearranges value*10+digit <= limit so the potentially
	// overflowing multiplication is never evaluated when it would be
	// unsafe.
	if value > (limit-digit)/10 {
		return 0, false
	}
	// The preceding inequality proves both the multiplication and
	// addition fit.
	return value*10 + digit, true
}
```

#### Work and memory

The function performs a fixed number of comparisons and arithmetic
operations, so its work is `O(1)`. It stores no collection or
recursive state, so extra memory is `O(1)`.

#### When to use it

Check before the operation whenever allowed inputs can approach an
integer limit. This matters especially while reading digits, building
a product, calculating a middle index, or converting between integer
widths.

#### When not to use it

Do not merely switch to a wider type when the mathematical result may
grow without a fixed bound. Use an integer representation that grows
as needed, such as `math/big`. Do not apply this nonnegative formula
unchanged to signed subtraction or to building a negative number from
digits.

#### Common mistakes and edge cases

> **Warning:** Calculating `next := value*10 + digit` and then
> checking `next > limit` can already have overflowed. For
> `(12,8,127)`, the check must happen before constructing `128` in the
> bounded type.

> **Warning:** Negating the smallest signed value can overflow because
> its positive magnitude may be one larger than the type's maximum.
> Handle that value explicitly instead of applying `-value` blindly.

The width of `int` depends on the target architecture. Use `int32`,
`int64`, `uint32`, or `uint64` when the problem requires an exact
width.

#### Related material

- [A.5 Bit manipulation](#a5-bit-manipulation) explains operations on
  a fixed-width bit pattern.
- [A.6.7 Numbers too large for an integer
  type](#a67-numbers-too-large-for-an-integer-type) processes large
  decimal input without converting it wholesale.
- [A.8 Integer digit processing](#integer-digit-processing) applies
  the precheck while reversing digits.
- [E.6 Overflow-safe integer
  reversal](#e6-overflow-safe-integer-reversal) provides the complete
  `int64` reversal implementation.

### A.5 Bit manipulation

#### The problem it solves

Given an integer's stored binary digits, test, set, clear, or combine
selected positions while leaving the other positions unchanged. This
is useful when each position represents an independent yes/no fact or
when the input promises that paired values cancel as their binary
digits are combined.

#### How it works

A **bit** is a switch that is either `0` or `1`. An integer used as a
row of these switches is a **bitmask**. AND (`&`) keeps shared `1`
bits, OR (`|`) keeps either input's `1` bits, and exclusive OR
(**XOR**, `^`) keeps bits that differ. Bit clear (`&^`) turns off
selected bits.

#### Step-by-step examples

Let bits `0`, `1`, and `2` mean read, write, and execute permission.
Start with read and execute enabled: binary `101`, which is decimal
`5`. The three displayed digits correspond to bits `2`, `1`, and `0`
from left to right. Leading zeroes in masks keep all three positions
visible.

Initial state: `flags=101`.

1.  Test write with mask `010`
    - Decision: `101 & 010` is zero.
    - After: `flags=101`, `writeEnabled=false`.
    - Known correct: Write is off; no flag changed.
2.  Set write with mask `010`
    - Decision: `101 | 010` turns bit `1` on.
    - After: `flags=111`.
    - Known correct: All three permissions are on.
3.  Clear read with mask `001`
    - Decision: `111 &^ 001` clears only bit `0`.
    - After: `flags=110`.
    - Known correct: Write and execute remain on; read is off.

For the cancellation input `[4,7,4]`, the expected unpaired value is
`7`. Here, `unpaired` stores the XOR of `values[:index]`, the input
prefix already processed. For example, after two values,
`unpaired=3` means `4 ^ 7 = 3`.

Initial state: `index=0`, `unpaired=0`.

1.  Combine the first `4`
    - Decision: Compute `0 ^ 4`.
    - After: `index=1`, `unpaired=4`.
    - Known correct: The processed prefix contains one unpaired
      `4`.
2.  Combine `7`
    - Decision: Compute `4 ^ 7`.
    - After: `index=2`, `unpaired=3`.
    - Known correct: `unpaired=3` is the XOR of processed prefix
      `[4,7]`.
3.  Combine the second `4`
    - Decision: Compute `3 ^ 4`.
    - After: `index=3`, `unpaired=7`.
    - Known correct: The equal `4` values cancel, leaving the
      expected unpaired value `7`.

For `[4,4,7,7]`, every value cancels and produces `0`. That input
violates the promise that exactly one value appears once.

#### Why it works

The one-bit mask is zero at every position except `k`, so an operation
with that mask can change or test only that position. OR with one
makes it one; bit clear with one makes it zero; AND reveals whether
both inputs contain one there.

XOR gives the same result when values are regrouped or reordered, so
input order does not matter. Two equal values produce zero when XORed,
and XOR with zero leaves the other value unchanged. That argument is
valid only when the input promise supplies exactly those pairs.

#### Go example

For the permission trace, `flags=0b101` and `bitPosition=1` produce
the same `0b010` mask shown above. Calls to `hasBit(0b101, 1)`,
`setBit(0b101, 1)`, and `clearBit(0b111, 0)` perform the three traced
operations. The final function uses the exact cancellation input
`values=[]uint64{4, 7, 4}` and returns `7`.

```go
func hasBit(flags uint64, bitPosition uint) bool {
	// A uint64 has positions 0 through 63. Shift a single 1 into
	// bitPosition, AND it with flags, and test whether that position
	// survives.
	return bitPosition < 64 && flags&(uint64(1)<<bitPosition) != 0
}

func setBit(flags uint64, bitPosition uint) (uint64, bool) {
	// Refuse an unrepresentable position instead of pretending it
	// exists.
	if bitPosition >= 64 {
		return flags, false
	}
	// OR turns the selected bit on and leaves every other bit unchanged.
	return flags | (uint64(1) << bitPosition), true
}

func clearBit(flags uint64, bitPosition uint) (uint64, bool) {
	if bitPosition >= 64 {
		return flags, false
	}
	// Go's bit-clear operator &^ turns off mask bits that are 1, so only
	// the selected bit is cleared.
	return flags &^ (uint64(1) << bitPosition), true
}

func unpairedValue(values []uint64) uint64 {
	// unpaired stores the XOR of the input prefix already processed.
	// Equal paired values cancel regardless of their input positions.
	var unpaired uint64
	for _, value := range values {
		unpaired ^= value
	}
	return unpaired
}
```

#### Work and memory

Each permission operation on one `uint64` performs bounded
machine-word work and uses no growing storage, so it takes `O(1)`
time and `O(1)` extra space. For `n` input values,
`unpairedValue` combines each value once, so it takes `O(n)` time and
`O(1)` extra space. A loop that enumerates every `n`-bit mask still
visits `2^n` subsets, and copying each chosen subset adds output work.

#### When to use it

Use bit operations when the problem defines a fixed-width bit pattern,
when a small bounded set of flags needs compact membership, or when an
input rule such as pair cancellation exactly matches the promise.

#### When not to use it

Do not use a mask when the number of members can exceed the available
width, when readable set operations matter more than compact storage,
or merely because the inputs happen to be integers. A map or boolean
slice is often clearer for general membership.

#### Common mistakes and edge cases

> **Warning:** Validate the bit position. Shifting a `uint64` mask by
> `64` or more cannot represent a bit in the value; the example
> returns failure.

> **Warning:** `^x` flips every bit, while `x ^ y` computes XOR
> between two values. They are different uses of the same token in Go.

Right-shifting a signed negative integer preserves sign bits; convert
only when the problem's width and signedness make that interpretation
valid. A subset mask is compact storage, not a shortcut around
exponential enumeration.

#### Related material

- [A.4 Integer limits and bits](#a4-integer-limits-and-bits) explains
  integer width and overflow rules.
- [A.7 Combinatorics and output
  size](#a7-combinatorics-and-output-size) accounts for all `2^n`
  subset outputs.
- [Single Number](problems/easy_leetcode_0136_single_number.go)
  applies XOR cancellation under an exact occurrence promise.
- [Reverse Bits](problems/easy_leetcode_0190_reverse_bits.go) treats
  the input as a fixed-width binary value.

### A.6 Keeping arithmetic values exact and consistent

#### The problem it solves

Given values such as fractions, directions, remainders, or large
decimal integers, represent mathematically equal values in exactly the
same program form. The representation must make comparisons,
arithmetic, and map lookups agree without relying on rounded
floating-point values.

#### How it works

Choose one kind of label before filing values. Store `$1.20` as `120`
cents, not sometimes as dollars and sometimes as cents. Rewrite
equivalent labels such as `2/4` and `1/2` into one agreed form. That
rewrite is called **normalization**, and its unique result is a
**canonical form**.

#### Step-by-step examples

Add the normalized prices `[]Cents{120,235,45}`, representing `$1.20`,
`$2.35`, and `$0.45`. The expected exact total is `400` cents, or
`$4.00`.

Initial state: `index=0`, `total=0` cents.

1.  Add `$1.20`
    - Decision: Read the normalized value `120` cents and add it.
    - After: `index=1`, `total=120` cents.
    - Known correct: The first price is exact.
2.  Add `$2.35`
    - Decision: Read and add the next normalized value, `235` cents.
    - After: `index=2`, `total=355` cents.
    - Known correct: The first two prices total `$3.55`.
3.  Add `$0.45`
    - Decision: Add integer `45`.
    - After: `index=3`, `total=400` cents.
    - Known correct: All prices total exactly `$4.00`.

The representation changes a real equality decision. If `first=0.10`
and `second=0.20` are stored in `float64` variables, then
`first+second == 0.30` is false because the stored binary
approximations do not add to the stored approximation of `0.30`.
With cents, `10+20 == 30` is exactly true.

The boundary value `$0.00` becomes zero cents and changes nothing. A
price with more precision than the chosen unit, such as `$0.005`, must
be rejected or rounded by an explicit business rule; silently dropping
the half-cent would break the chosen representation.

#### Why it works

Every accepted amount is an integer count of the same unit. Integer
addition therefore combines exact counts, and equal amounts have equal
stored values. Normalizing once at the boundary means later code
cannot accidentally compare two different encodings of the same
accepted amount.

#### Go example

Calling `addPrices([]Cents{120, 235, 45})` uses the exact normalized
input from the trace and returns `400`. Boundary code must convert
accepted prices to cents before this function runs, and their total
must fit in `int64`.

```go
// Cents gives every stored price the same exact integer unit.
type Cents int64

func addPrices(prices []Cents) Cents {
	// The zero value is the additive identity, so it is the correct
	// empty sum.
	var total Cents
	// Add each already-normalized amount; no floating-point rounding
	// occurs.
	for index := range prices {
		total += prices[index]
	}
	return total
}
```

#### Work and memory

Adding `n` already-normalized prices reads each once, so it takes
`O(n)` time in every case and `O(1)` extra memory. Parsing,
reduction, or output formatting must be counted separately when the
boundary representation is not already available.

#### When to use it

Use an exact representation when values become equality keys, when
rounding could change a branch, or when the problem defines a precise
unit, ratio, direction, or remainder.

#### When not to use it

Do not invent a scaled integer when the input can contain any number
of decimal places or when the required rounding rule is unknown. Use
a decimal or fraction type that grows as needed, or clarify the rule
first.

#### Common mistakes and edge cases

> **Warning:** Mixing units produces plausible but wrong totals.
> Adding `120` cents directly to `2.35` dollars does not represent
> `$3.55`; normalize both operands before arithmetic.

> **Warning:** Canonical form is part of equality. If `1/2` is reduced
> but `2/4` is not, they remain different Go struct keys despite
> representing the same quantity.

Zero, negative values, invalid denominators, and overflow need
explicit rules for the chosen representation.

#### Related material

- [A.4 Integer limits and bits](#a4-integer-limits-and-bits) explains
  how to guard fixed-width intermediate calculations.
- [A.6.1 Shared divisors and
  multiples](#a61-shared-divisors-and-multiples-gcd-and-lcm) supplies
  the reduction operation used by ratio-based child entries.
- [A.6.7 Numbers too large for an integer
  type](#a67-numbers-too-large-for-an-integer-type) keeps oversized
  decimal input as text.

#### A.6.1 Shared divisors and multiples (GCD and LCM)

##### The problem it solves

Given two nonnegative counts, find either the largest group size that
divides both counts exactly or the smallest positive length at which
both repeating cycles meet. Listing all divisors or successive
multiples does unnecessary work when the counts are large.

##### How it works

If `48` and `18` are split into equally sized groups, the largest
possible group has `6` items. That value is the **greatest common
divisor (GCD)**. The first positive number reached by both counting in
steps of `48` and `18` is `144`; it is the **least common multiple
(LCM)**. **Euclid's algorithm** finds the GCD by repeatedly replacing
the larger question with its leftover part.

##### Step-by-step examples

For `firstValue=48` and `secondValue=18`, the expected GCD is `6`
and LCM is `144`.

Initial state: `(firstValue,secondValue)=(48,18)`, `gcd=unknown`,
`lcm=unknown`.

1.  Take the first remainder
    - Decision: `48%18=12`; any common divisor of `48` and `18`
      also divides `12`.
    - After: `(firstValue,secondValue)=(18,12)`, `gcd=unknown`,
      `lcm=unknown`.
    - Known correct: The set of common divisors is unchanged.
2.  Take the next remainder
    - Decision: `18%12=6`; replace the pair again.
    - After: `(firstValue,secondValue)=(12,6)`, `gcd=unknown`,
      `lcm=unknown`.
    - Known correct: The smaller pair still has exactly the original
      common divisors.
3.  Reach a zero remainder
    - Decision: `12%6=0`; stop the GCD loop.
    - After: `(firstValue,secondValue)=(6,0)`, `gcd=6`,
      `lcm=unknown`.
    - Known correct: `6` is the greatest common divisor.
4.  Compute the LCM
    - Decision: Compute `48/6*18`.
    - After: `gcd=6`, `lcm=144`.
    - Known correct: `144` is the smallest positive multiple shared
      by both values.

For boundary input `(0,12)`, the GCD loop returns `12`, while the LCM
rule returns `0`. For `(0,0)`, this guide uses the programming
convention GCD `0`.

##### Why it works

Dividing `firstValue` by `secondValue` gives some whole-number
quotient and a remainder. A number divides both values exactly when
it divides `secondValue` and that remainder. Replacing the pair with
`(secondValue, firstValue%secondValue)` therefore keeps every shared
divisor. Each nonzero remainder is smaller than `secondValue`, so the
process finishes. When `secondValue=0`, `firstValue` is the largest
divisor that the original values share.

For positive inputs, `firstValue*secondValue` contains the shared
factors twice. Dividing one operand by the GCD removes one shared
copy, producing the smallest value that contains the factors needed
by both inputs.

##### Go example

The functions below require nonnegative inputs. For the traced input,
`gcd(48, 18)` returns `6`, and `lcm(48, 18)` returns `(144, true)`.

```go
import "math"

func gcd(firstValue, secondValue int64) int64 {
	// Replacing the pair with (secondValue, firstValue%secondValue)
	// preserves their common divisors while making the second value
	// smaller.
	for secondValue != 0 {
		firstValue, secondValue = secondValue, firstValue%secondValue
	}
	// When secondValue is zero, firstValue is the greatest common
	// divisor.
	return firstValue
}

func lcm(firstValue, secondValue int64) (int64, bool) {
	// Under this convention, any least common multiple involving zero is
	// zero.
	if firstValue == 0 || secondValue == 0 {
		return 0, true
	}
	// Divide out the shared factors before multiplying. This reduces the
	// size of the multiplication without changing the mathematical LCM.
	reducedFirstValue := firstValue / gcd(firstValue, secondValue)
	// Prove the product fits in int64 before evaluating it.
	if reducedFirstValue > math.MaxInt64/secondValue {
		return 0, false
	}
	return reducedFirstValue * secondValue, true
}
```

##### Work and memory

Euclid's algorithm takes a number of remainder steps proportional to
the logarithm of the smaller input. This is written
`O(log(min(firstValue,secondValue)))`. It uses `O(1)` extra memory.
LCM adds a fixed amount of work and no growing storage after the GCD.
These bounds treat fixed-width integer operations as constant time.

##### When to use it

Use GCD to reduce ratios, align repeating periods, group counts
evenly, or normalize integer directions. Use LCM when independent
positive cycles must meet again.

##### When not to use it

Do not use LCM as a product shortcut without an overflow check. To
list factors or prime numbers instead, use the methods in
[A.6.6 Factors and prime
sieves](#a66-factors-and-prime-sieves).

##### Common mistakes and edge cases

> **Warning:** `a*b/gcd(a,b)` can overflow during `a*b` even when the
> final value fits. Divide first, as in `a/gcd(a,b)*b`.

> **Warning:** The example code requires nonnegative values. Passing
> `-18` changes remainder signs and breaks the stated input rule;
> normalize signs at the boundary, taking care with the smallest
> signed integer.

Zero inputs need declared conventions, and the checked multiplication
can still reject an LCM too large for `int64`.

##### Related material

- [A.6 Keeping arithmetic values exact and
  consistent](#a6-keeping-arithmetic-values-exact-and-consistent)
  explains the normalization rule.
- [A.6.2 Normalized fractions and
  ratios](#a62-normalized-fractions-and-ratios) uses GCD to create
  unique ratio keys.
- [A.6.3 Exact slopes without
  rounding](#a63-exact-slopes-without-rounding) uses the same
  reduction for directions.

#### A.6.2 Normalized fractions and ratios

##### The problem it solves

Given an integer numerator and denominator, produce one standard pair
that can be compared or used as a Go map key. Equivalent inputs such
as `(2,4)` and `(3,6)` must both become `(1,2)`, without converting
the ratio to a rounded floating-point number.

##### How it works

In a fraction such as `3/4`, the top number `3` is the **numerator**,
and the bottom number `4` is the **denominator**. Cancel factors
shared by both numbers and keep the minus sign in one agreed place.
The result is a **normalized fraction**: its two numbers share no
positive factor greater than one, and its denominator is positive.

##### Step-by-step examples

Normalize `(6,-8)`; the expected result is `(-3,4)`.

Initial state: `numerator=6`, `denominator=-8`,
`greatestCommonDivisor=unknown`.

1.  Normalize the sign
    - Decision: The denominator must be positive, so reverse both
      signs.
    - After: `numerator=-6`, `denominator=8`,
      `greatestCommonDivisor=unknown`.
    - Known correct: The value is unchanged and the sign has one
      location.
2.  Find the shared factor
    - Decision: Run Euclid's algorithm on magnitudes `(6,8)`.
    - After: `numerator=-6`, `denominator=8`,
      `greatestCommonDivisor=2`.
    - Known correct: `2` is the largest shared factor.
3.  Reduce both parts
    - Decision: Divide numerator and denominator by `2`.
    - After: `numerator=-3`, `denominator=4`,
      `greatestCommonDivisor=2`.
    - Known correct: `(-3,4)` has the canonical sign and no remaining
      shared factor.

The equivalent input `(-9,12)` also reduces to `(-3,4)`. A Go map
therefore uses one key for both ratios instead of storing two entries
for the same mathematical value.

`(0,-8)` is the boundary form `(0,1)`: every nonzero denominator
represents the same zero. `(1,0)` is invalid and returns failure.

##### Why it works

Reversing both signs multiplies numerator and denominator by `-1`, so
the ratio is unchanged. Dividing both by the same nonzero GCD also
preserves the ratio. After division the parts are relatively prime,
meaning they have no shared factor greater than one. A positive
denominator fixes the only remaining sign choice, so equal accepted
fractions reach the same pair.

##### Go example

The function requires neither input to equal `math.MinInt64`. The
positive magnitude of that one value cannot be stored in `int64`.
For the traced input, `normalizeFraction(6, -8)` returns
`(-3, 4, true)`.

```go
func normalizeFraction(
	numerator, denominator int64,
) (int64, int64, bool) {
	// Division by zero is not a fraction, while every zero numerator has
	// the single canonical representation 0/1.
	if denominator == 0 {
		return 0, 0, false
	}
	if numerator == 0 {
		return 0, 1, true
	}
	if denominator < 0 {
		// Keep the denominator positive by moving any sign to the
		// numerator.
		numerator = -numerator
		denominator = -denominator
	}

	// Euclid's algorithm needs magnitudes. Only the numerator can still
	// be negative because the denominator was normalized above.
	numeratorMagnitude := numerator
	denominatorMagnitude := denominator
	if numeratorMagnitude < 0 {
		// GCD operates on magnitude; the original signed numerator remains
		// intact for the final normalized result.
		numeratorMagnitude = -numeratorMagnitude
	}
	// Each Euclidean step preserves the shared divisors and reduces the
	// second operand until numeratorMagnitude contains the GCD.
	for denominatorMagnitude != 0 {
		numeratorMagnitude, denominatorMagnitude =
			denominatorMagnitude,
			numeratorMagnitude%denominatorMagnitude
	}
	// Divide both parts by their GCD to obtain the unique reduced pair.
	greatestCommonDivisor := numeratorMagnitude
	return numerator / greatestCommonDivisor,
		denominator / greatestCommonDivisor,
		true
}
```

##### Work and memory

Sign handling takes a fixed amount of work. GCD takes a number of
remainder steps proportional to the logarithm of the smaller input
magnitude. The function uses `O(1)` extra memory and returns two
integers without allocating a collection.

##### When to use it

Use a normalized pair when exact ratios become map keys, equality
values, or directions and the integer parts remain within a known
bound.

##### When not to use it

Do not use it for a zero denominator or when measured values may count
as equal within an allowed difference. Do not use it when earlier
subtraction has already overflowed before the pair reaches this
function.

##### Common mistakes and edge cases

> **Warning:** Reducing magnitude without fixing the sign leaves two
> keys for one value. `(3,-4)` and `(-3,4)` must both normalize to
> `(-3,4)`.

> **Warning:** Zero needs one canonical key. Leaving `(0,8)` and
> `(0,3)` as-is makes equal zero ratios compare unequal.

Validate `denominator != 0` before remainder or division, and handle
the smallest signed integer under an explicit wider or
unsigned-magnitude policy.

##### Related material

- [A.6.1 Shared divisors and
  multiples](#a61-shared-divisors-and-multiples-gcd-and-lcm) proves
  the reduction step.
- [A.6.3 Exact slopes without
  rounding](#a63-exact-slopes-without-rounding) adds
  direction-specific duplicate and axis cases.
- [D.4 Map presence and sets](#d4-map-presence-and-sets) explains how
  normalized comparable pairs behave as Go map keys.

#### A.6.3 Exact slopes without rounding

##### The problem it solves

Given a starting point and other integer-coordinate points, group the
other points by the exact line direction they share with the start.
Changes such as `(2,4)` and `(-1,-2)` must receive the same direction
key, while distinct directions must not collide because a decimal
slope was rounded.

##### How it works

Two hiking directions such as `2` steps right and `4` up, or `1`
right and `2` up, point the same way. Store horizontal change first
and vertical change second, then divide both changes by their shared
factor to make one exact **direction key**. When opposite travel
counts as the same undirected line, choose one agreed sign. The
integer pair avoids rounded decimal slopes.

##### Step-by-step examples

Use start `(1,1)` and points `(2,3)`, `(3,5)`, and `(0,-1)`. Every
expected direction key is `(1,2)`.

Initial state: `pointIndex=0`, `keys=[]`.

1.  Normalize point `(2,3)`
    - Decision: Its change from `(1,1)` is `(1,2)`; GCD is `1` and
      the sign is canonical.
    - After: `pointIndex=1`, `keys=[(1,2)]`.
    - Known correct: The first key represents its exact direction.
2.  Normalize point `(3,5)`
    - Decision: Its change is `(2,4)`; divide both parts by GCD `2`.
    - After: `pointIndex=2`, `keys=[(1,2),(1,2)]`.
    - Known correct: The longer arrow shares the first key.
3.  Normalize point `(0,-1)`
    - Decision: Its change is `(-1,-2)`; GCD is `1`, then reverse
      both signs because
      `changeX < 0`.
    - After: `pointIndex=3`, `keys=[(1,2),(1,2),(1,2)]`.
    - Known correct: The opposite arrow has the same undirected
      line orientation through the start.

For point `(1,1)`, both changes are zero, so the result is marked
duplicate instead of inventing a slope. For `(1,5)`, the fixed
vertical key is `(0,1)`. Direct integer division would try to divide
the vertical change by horizontal change `0`; the pair representation
handles that input without division or a special floating-point key.

##### Why it works

Points on the same line through the start have movement pairs that are
integer multiples of the same reduced pair. Dividing by the GCD
removes that scale. The axis rules remove arbitrary magnitude, and the
sign rule chooses one of the two equivalent orientations for an
undirected line. Duplicate points are kept separate because `(0,0)` is
a multiple of every direction and therefore cannot identify one.

##### Go example

The coordinate differences must fit in `int64` and must not equal
`math.MinInt64`; that value has no positive `int64` magnitude.
Subtract the start before calling the function. In the trace, the
three points produce calls `normalizeDirection(1, 2)`,
`normalizeDirection(2, 4)`, and `normalizeDirection(-1, -2)`; each
returns the key `(1,2)` with duplicate set to `false`.

```go
// direction is comparable, so a normalized value can be used as a map
// key.
type direction struct {
	horizontalChange int64
	verticalChange   int64
}

func normalizeDirection(
	horizontalChange, verticalChange int64,
) (direction, bool) {
	// Two identical points have no direction. The boolean reports this
	// duplicate separately from valid vertical and horizontal
	// directions.
	if horizontalChange == 0 && verticalChange == 0 {
		return direction{}, true
	}
	if horizontalChange == 0 {
		// Upward and downward vectors share the canonical vertical key
		// (0,1) because this function represents an undirected line
		// orientation.
		return direction{horizontalChange: 0, verticalChange: 1}, false
	}
	if verticalChange == 0 {
		// Leftward and rightward vectors likewise share the canonical
		// horizontal key (1,0).
		return direction{horizontalChange: 1, verticalChange: 0}, false
	}

	// Compute the GCD of the absolute components so equivalent scaled
	// vectors reduce to the same pair.
	horizontalMagnitude := horizontalChange
	verticalMagnitude := verticalChange
	if horizontalMagnitude < 0 {
		horizontalMagnitude = -horizontalMagnitude
	}
	if verticalMagnitude < 0 {
		verticalMagnitude = -verticalMagnitude
	}
	// Euclid's algorithm leaves the positive GCD in horizontalMagnitude.
	for verticalMagnitude != 0 {
		horizontalMagnitude, verticalMagnitude =
			verticalMagnitude,
			horizontalMagnitude%verticalMagnitude
	}
	horizontalChange /= horizontalMagnitude
	verticalChange /= horizontalMagnitude
	// Choose one orientation convention: horizontalChange is always
	// positive for nonvertical directions, so opposite sign encodings
	// cannot create duplicate keys.
	if horizontalChange < 0 {
		horizontalChange = -horizontalChange
		verticalChange = -verticalChange
	}
	return direction{
		horizontalChange: horizontalChange,
		verticalChange:   verticalChange,
	}, false
}
```

The boolean result reports a duplicate point.

##### Work and memory

A caller takes a fixed amount of work to subtract the start point from
one other point. The GCD takes a number of remainder steps
proportional to the logarithm of the smaller horizontal or vertical
change. Normalizing one direction uses `O(1)` extra memory. A caller
grouping `n` points must also count its map storage and lookups.

##### When to use it

Use reduced direction keys when integer-coordinate points must be
grouped by the exact straight line they share with one chosen start.
Use this sign rule only when opposite directions count as the same
line.

##### When not to use it

Do not use a rounded floating-point slope as an equality key. Do not
use this undirected sign rule when opposite rays must remain
different; keep their signs distinct for that requirement.

##### Common mistakes and edge cases

> **Warning:** Dividing `changeY/changeX` loses exactness and divides
> by zero on vertical lines. Normalize the integer pair instead.

> **Warning:** Treating duplicate `(0,0)` as a direction lets it match
> every slope incorrectly. Count duplicates separately.

Calculate coordinate differences before narrowing their type.
Horizontal, vertical, negative, and duplicate cases all need explicit
tests.

##### Related material

- [A.6.1 Shared divisors and
  multiples](#a61-shared-divisors-and-multiples-gcd-and-lcm)
  explains Euclid's algorithm.
- [A.6.2 Normalized fractions and
  ratios](#a62-normalized-fractions-and-ratios) explains canonical
  reduced pairs.
- [Max Points on a
  Line](problems/hard_leetcode_0149_max_points_on_a_line.go) uses
  direction keys with duplicate counting.

#### A.6.4 Modular arithmetic and negative division

##### The problem it solves

Given possibly negative integers, compute the quotient or remainder
under the convention required by the problem. Go truncates division
toward zero and can return a negative remainder, so code that needs
a floor quotient or a bucket in `0..m-1` must convert the operator
result explicitly.

##### How it works

Wrapping `-1` around a five-position clock lands at position `4`, even
though Go evaluates `-1 % 5` as `-1`. Moving a negative Go remainder
back into the clock range is **normalized modular arithmetic**.
Rounding a quotient down toward negative infinity is **floor
division**; Go instead truncates toward zero.

##### Step-by-step examples

For `x=-7` and `m=3`, the normalized remainder should be `2` and the
floor quotient should be `-3`.

Initial state: `x=-7`, `m=3`, `quotient=unknown`,
`remainder=unknown`.

1.  Apply Go's operators
    - Decision: Divide and take the remainder using truncation toward
      zero.
    - After: `quotient=-2`, `remainder=-1`.
    - Known correct: `-7 = -2*3 + -1`.
2.  Normalize the remainder
    - Decision: Add the positive modulus `3` to `-1`.
    - After: `quotient=-2`, `remainder=-1`,
      `normalizedRemainder=2`.
    - Known correct: The normalized result is from `0` through `2` and
      differs from `-7` by a multiple of `3`.
3.  Convert to the floor quotient
    - Decision: The original remainder is negative, so move the
      truncated quotient one whole step downward.
    - After: `normalizedRemainder=2`, `floorQuotient=-3`.
    - Known correct: `-3 <= -7/3 < -2`.

Boundary input `x=-6,m=3` has remainder zero. Its exact quotient is
`-2`, so neither correction changes the result.

##### Why it works

Go's remainder is always closer to zero than positive `m`. A negative
remainder is therefore between `-(m-1)` and `-1`. Adding `m` moves it
between `1` and `m-1` without changing its clock position. For
division, truncation and floor differ only when a negative result has
a fractional part. With positive `m`, a negative remainder identifies
exactly that case, so subtracting one reaches the next lower integer.

##### Go example

For the traced input, `normalizedRemainder(-7, 3)` returns
`(2, true)`, and `floorDivide(-7, 3)` returns `(-3, true)`. The trace
uses `m` for the positive second argument; the functions name that
same role `modulus` or `divisor` according to the requested result.

```go
func normalizedRemainder(value, modulus int64) (int64, bool) {
	// A positive modulus defines canonical representatives 0 through
	// modulus-1.
	if modulus <= 0 {
		return 0, false
	}
	remainder := value % modulus
	if remainder < 0 {
		// Go keeps the dividend's sign; adding one modulus moves a negative
		// remainder into the canonical nonnegative interval.
		remainder += modulus
	}
	return remainder, true
}

func floorDivide(value, divisor int64) (int64, bool) {
	if divisor <= 0 {
		return 0, false
	}
	// Go truncates toward zero. For a negative nonmultiple, floor is one
	// less than that truncated quotient.
	quotient := value / divisor
	if value%divisor < 0 {
		quotient--
	}
	return quotient, true
}
```

##### Work and memory

Each function performs a fixed number of integer operations, so it
takes `O(1)` time and `O(1)` extra memory.

##### When to use it

Use normalized remainders for circular indexes, groups of values that
share the same remainder, and problems that require
`0 <= result < modulus`. Use floor division when ranges or coordinates
are defined by rounding downward.

##### When not to use it

Do not add the correction when the required behavior is Go's
truncation toward zero. The formulas above also deliberately require a
positive divisor; derive and test a different rule before accepting
negative divisors.

##### Common mistakes and edge cases

> **Warning:** In Go, `-1 % 5` is `-1`, so using it directly as a
> slice index is invalid. Normalize it to `4` first.

> **Warning:** Always subtracting one for negative `x` breaks exact
> divisions. `-6/3` is already `-2`; adjust only when the remainder is
> negative.

Reject zero modulus or divisor before the operator executes. If a
different version permits divisor `-1`, it must reject the smallest
signed value because its positive quotient does not fit the type.

##### Related material

- [A.4 Integer limits and bits](#a4-integer-limits-and-bits) explains
  why operator inputs must also fit their type.
- [A.6 Keeping arithmetic values exact and
  consistent](#a6-keeping-arithmetic-values-exact-and-consistent)
  explains how to choose one representation.
- [A.6.7 Numbers too large for an integer
  type](#a67-numbers-too-large-for-an-integer-type) updates a
  remainder while reading decimal text.

#### A.6.5 Trailing zeroes without calculating a factorial

##### The problem it solves

Given a nonnegative integer `n`, return how many decimal zeroes appear
at the end of `n!`. Constructing the factorial is unnecessary and soon
exceeds ordinary integer types; the answer depends only on how many
matched factors of `2` and `5` its product contains.

##### How it works

A trailing decimal zero comes from multiplying `2*5`. The product
`1*2*...*n`, called **`n` factorial** and written `n!`, contains more
factors of `2` than factors of `5`. Count the scarce `5`s instead of
building the product.

##### Step-by-step examples

For `n=25`, the expected result is `6`.

Initial state: `remaining=25`, `zeroes=0`.

1.  Count the first factor layer
    - Decision: `25/5=5` counts one `5` from `5,10,15,20,25`.
    - After: `remaining=5`, `zeroes=5`.
    - Known correct: Every number with at least one `5` has
      contributed once.
2.  Count the second factor layer
    - Decision: `5/5=1` counts the second `5` inside `25`.
    - After: `remaining=1`, `zeroes=6`.
    - Known correct: Every factor of `25` has contributed its
      extra `5`.
3.  Stop
    - Decision: `remaining=1 < 5`; no higher layer exists.
    - After: `result=6`.
    - Known correct: All factors of `5` in `25!` are counted.

For boundary inputs `0` through `4`, no factor of `5` occurs, so the
loop does no work and returns `0`.

##### Why it works

Every trailing zero uses one factor of `2` and one factor of `5`. In a
factorial, each multiple of `5` supplies one `5`, each multiple of
`25` supplies one more, and higher powers supply further layers. Add
the whole-number results of dividing `n` by `5`, `25`, `125`, and so
on. This counts every factor of `5` once per occurrence. Factors of
`2` are more common, so the number of pairs is exactly the number of
counted `5`s.

##### Go example

The function requires `n >= 0`. For the traced input,
`trailingZeroes(25)` returns `6`.

```go
func trailingZeroes(n int64) int64 {
	remaining := n
	var zeroes int64
	// Each division counts multiples of 5, then the extra factors
	// contributed by multiples of 25, 125, and so on.
	for remaining >= 5 {
		remaining /= 5
		zeroes += remaining
	}
	return zeroes
}
```

##### Work and memory

Each pass divides `n` by `5`, so there are `O(log_5 n)` passes,
conventionally written `O(log n)`. The function stores two integers
and uses `O(1)` extra memory. It never allocates or constructs
`n!`.

##### When to use it

Use repeated factor counting when a question asks about factors inside
a factorial product, especially trailing zeroes in a known base.

##### When not to use it

Do not use this base-10 shortcut for trailing zeroes in another base
without factoring that base and counting all required prime factors.
Do not construct the factorial merely to inspect its decimal suffix.

##### Common mistakes and edge cases

> **Warning:** Counting only `n/5` misses repeated factors. For `25!`,
> it gives `5` instead of `6` because `25=5*5` contributes twice.

> **Warning:** Multiplying a power variable by `5` can overflow at the
> loop boundary. Repeatedly divide `n`, as the example does.

Reject negative `n`; factorial is not defined by this entry for
negative integers.

##### Related material

- [A.6.6 Factors and prime sieves](#a66-factors-and-prime-sieves)
  defines prime factors and repeated divisibility.
- [Factorial Trailing
  Zeroes](problems/medium_leetcode_0172_factorial_trailing_zeroes.go)
  applies the factor-layer count.

#### A.6.6 Factors and prime sieves

##### The problem it solves

Given either one integer or an upper limit, determine whether that one
integer is prime or list every prime up to the limit. Testing every
possible divisor for every candidate repeats work; factor pairs and
shared composite marking provide tighter bounds.

##### How it works

Factors come in pairs: if `91=7*13`, one member is no larger than the
square root. For one number, try only possible small partners. For a
whole range, start with every candidate and repeatedly cross out
multiples. That range method is the **Sieve of Eratosthenes**. A
**prime** is an integer greater than one whose only positive divisors
are `1` and itself. In the sieve, `composite[value] -> whether a
discovered prime divides value`. An initial `false` means only "not
marked yet"; it means "confirmed prime" only after all possible
smaller prime factors have been processed.

##### Step-by-step examples

For trial division with `value=91`, the expected smallest factor is
`7`.

Initial state: `divisor=2`, `smallestFactor=unknown`.

1.  Reject the smaller divisors
    - Decision: Test `2`, `3`, `4`, `5`, and `6` in order. Their
      remainders are `1`, `1`, `3`, `1`, and `1`, so advance after
      each test.
    - After: `divisor=7`, `smallestFactor=unknown`.
    - Known correct: No integer from `2` through `6` divides
      `91`.
2.  Accept divisor `7`
    - Decision: `91%7=0`, so `7` divides `91`.
    - After: `smallestFactor=7`, `pairedFactor=13`.
    - Known correct: Ascending order makes `7` the smallest
      factor.

For a sieve through `10`:

Initial state: `prime=2`; for indexes `2..10`,
`composite=[F,F,F,F,F,F,F,F,F]`.

1.  Keep `2`
    - Decision: `2` is unmarked; cross out `4,6,8,10`.
    - After: `prime=3`; for indexes `2..10`,
      `composite=[F,F,T,F,T,F,T,F,T]`.
    - Known correct: `2` is prime; its composite multiples cannot
      be prime.
2.  Keep `3`
    - Decision: `composite[3]=F`; start at `3*3=9` and mark `9`.
    - After: `prime=4`; for indexes `2..10`,
      `composite=[F,F,T,F,T,F,T,T,T]`.
    - Known correct: `3` is prime; smaller multiples were already
      handled.
3.  Stop crossing out values
    - Decision: `prime=4` and `4*4 > 10`; every composite has a
      factor at most its square
      root.
    - After: Read indexes `2..10`; `primes=[2,3,5,7]`.
    - Known correct: These and only these values remain unmarked after
      every possible smaller prime factor was processed.

For `limit < 2`, the result is empty. For `n=2`, trial division
performs no candidate checks and reports `2` itself as prime.

##### Why it works

If a non-prime value `n=a*b` had both factors larger than its square
root, their product would exceed `n`. Therefore checking through the
square root cannot miss a factor. In the sieve, every non-prime value
has a smallest prime factor `p`; when `p` is processed, that value is
marked. Starting at `p*p` is safe because any smaller multiple of `p`
has another factor smaller than `p` and was already marked.

##### Go example

```go
func smallestFactor(value int) int {
	// Values below two have no prime factor under this function's
	// contract.
	if value < 2 {
		return 0
	}
	// If value is composite, at least one factor is no larger than its
	// square root. Writing divisor <= value/divisor avoids overflowing
	// divisor*divisor.
	for divisor := 2; divisor <= value/divisor; divisor++ {
		if value%divisor == 0 {
			// Ascending search makes the first divisor found the smallest one.
			return divisor
		}
	}
	// No smaller divisor exists, so value itself is prime.
	return value
}

func primesThrough(limit int) []int {
	// No prime exists below two, so there is nothing to allocate or
	// scan.
	if limit < 2 {
		return nil
	}
	// composite[value] records that a previously discovered prime
	// divides value.
	composite := make([]bool, limit+1)
	// Only bases through sqrt(limit) can have a first unmarked multiple
	// within the requested range.
	for prime := 2; prime <= limit/prime; prime++ {
		if composite[prime] {
			// Composite bases were already handled by their prime factors.
			continue
		}
		// Smaller multiples were marked by smaller primes, so begin at
		// prime^2.
		multiple := prime * prime
		for multiple <= limit {
			composite[multiple] = true
			multiple += prime
		}
	}
	// Every unmarked value at least two is prime.
	primes := make([]int, 0)
	for value := 2; value <= limit; value++ {
		if !composite[value] {
			// An unmarked value has no smaller prime divisor and is therefore
			// prime.
			primes = append(primes, value)
		}
	}
	return primes
}
```

##### Work and memory

Trial division checks at most the integers through the square root of
`n`. This is written `O(sqrt(n))`; it uses `O(1)` extra memory. The
sieve crosses out multiples of each discovered prime. Its standard
bound is `O(limit log log limit)` marking work and `O(limit)` marker
memory. Collecting the answers adds one `O(limit)` scan and an output
slice with one element per returned prime.

##### When to use it

Use trial division for one or a few moderate integers. Use a sieve
when many queries share the same upper limit or when every prime in a
range is needed.

##### When not to use it

Do not allocate a marker for every value up to an enormous limit when
only a handful of values will be queried. Likewise, do not test every
number separately when one sieve can handle the whole range together.

##### Common mistakes and edge cases

> **Warning:** Values `0` and `1` are not prime. Treating every number
> with no found divisor as prime would classify both incorrectly.

> **Warning:** `divisor*divisor <= n` can overflow for a large
> fixed-width integer. Use `divisor <= n/divisor` when the types
> approach their limit.

Begin sieve marking at `prime*prime`, but only after the outer
condition proves that square is within the limit.

##### Related material

- [A.6.1 Shared divisors and
  multiples](#a61-shared-divisors-and-multiples-gcd-and-lcm) compares
  shared-factor queries with factor enumeration.
- [A.6.5 Trailing
  zeroes](#a65-trailing-zeroes-without-calculating-a-factorial) counts
  one prime factor across an implicit product.

#### A.6.7 Numbers too large for an integer type

##### The problem it solves

Given a decimal integer as text, compute a bounded property such as
its remainder after division even when the full number has thousands
of digits and cannot fit in `int64`. Each digit must update only the
small state needed for that property rather than constructing the
entire numeric value.

##### How it works

Read the number the way long division does: carry only the unfinished
remainder into the next digit. The full number stays as text, and the
remembered remainder never needs to grow with the full number.

##### Step-by-step examples

For digits `"1234"` and modulus `7`, the expected remainder is `2`.

Initial state: `digitIndex=0`, `remainder=0`.

1.  Read `'1'`
    - Decision: Compute `(0*10+1)%7=1`.
    - After: `digitIndex=1`, `remainder=1`.
    - Known correct: The remainder of prefix `"1"` is `1`.
2.  Read `'2'`
    - Decision: Compute `(1*10+2)%7=5`.
    - After: `digitIndex=2`, `remainder=5`.
    - Known correct: The remainder of prefix `"12"` is `5`.
3.  Read `'3'`
    - Decision: Compute `(5*10+3)%7=4`.
    - After: `digitIndex=3`, `remainder=4`.
    - Known correct: The remainder of prefix `"123"` is `4`.
4.  Read `'4'`
    - Decision: Compute `(4*10+4)%7=2`.
    - After: `digitIndex=4`, `remainder=2`.
    - Known correct: The remainder of the complete input is `2`.

The same updates process
`"123456789012345678901234567891"`, which cannot fit in `int64`.
With modulus `7`, they return remainder `1`, and the remembered
remainder stays between `0` and `6` after every digit.

Leading zeroes are harmless: `"0012"` follows extra zero states and
still returns `5` modulo `7`. Empty text and `"12x"` are invalid.

##### Why it works

Before each digit, `remainder` equals the numeric prefix divided by
`modulus`, keeping only the remainder. Appending `digit` multiplies
the prefix by `10` and adds `digit`. Replacing the old prefix with its
remainder does not change the new remainder, so the update preserves
the fact for the longer prefix. After the last digit, that prefix is
the complete input.

##### Go example

```go
const maxSafeDecimalModulus int64 = 922337203685477580

func decimalRemainder(digits string, modulus int64) (int64, bool) {
	// The modulus bound ensures remainder*10+digit fits before the next
	// modulo.
	if digits == "" || modulus <= 0 || modulus > maxSafeDecimalModulus {
		return 0, false
	}
	var remainder int64
	for digitIndex := 0; digitIndex < len(digits); digitIndex++ {
		// Validate and convert one decimal digit byte without parsing the
		// whole, potentially enormous, number.
		if digits[digitIndex] < '0' || digits[digitIndex] > '9' {
			return 0, false
		}
		digit := int64(digits[digitIndex] - '0')
		// Appending a digit changes the remainder to (10*remainder+digit)
		// mod modulus. Congruent prefixes stay congruent after appending,
		// so retaining only the remainder is sufficient.
		remainder = (remainder*10 + digit) % modulus
	}
	return remainder, true
}
```

##### Work and memory

For `d` input digits, the function performs one update per digit, so
it takes `O(d)` time in every valid case. It stores one remainder and
indexes the input without copying it, so extra memory is `O(1)`.

##### When to use it

Use left-to-right remainder updates when the input is an exact decimal
string and only divisibility, a remainder, or another small remembered
value is needed.

##### When not to use it

Do not use this method when later work needs the complete large value
or arithmetic on it. Use Go's `big.Int`, an integer type that grows as
needed. Do not accept a modulus that can overflow the update unless
the implementation uses a wider or explicitly checked operation.

##### Common mistakes and edge cases

> **Warning:** Converting the whole text before taking `% modulus`
> fails before the useful bounded operation runs. A 1,000-digit input
> cannot fit in `int64`.

> **Warning:** Validate every byte. Treating `'x'-'0'` as a digit
> silently corrupts the remembered remainder for `"12x"`.

Choose and document rules for a leading sign, whitespace, and empty
input. The example accepts leading zeroes but digits only.

##### Related material

- [A.4 Integer limits and bits](#a4-integer-limits-and-bits) explains
  the check that keeps intermediate values in range.
- [A.6.4 Modular arithmetic and negative
  division](#a64-modular-arithmetic-and-negative-division) defines
  normalized remainders for signed values.
- [A.8 Integer digit processing](#integer-digit-processing) extracts
  digits from an integer that already fits in memory.

### A.7 Combinatorics and output size

#### The problem it solves

Given a request to count or generate groups, orderings, or subsets,
estimate how many distinct results can exist before choosing an
algorithm. Even a short input can require an enormous output, and a
method that must return every result cannot run faster than the work
needed to construct that output.

#### How it works

If every input value independently has an "in" or "out" switch, `n`
values create `2^n` switch settings: multiply by `2` once for each
value. No implementation can return all those groups without visiting
and storing the requested answers. **Combinatorics** is the study of
counting such choices.

#### Step-by-step examples

Generate subsets of `[A,B]`. At each position, first exclude and then
include the current value. The expected output order is
`[[],[B],[A],[A,B]]`.

Initial state: `currentIndex=0`, `currentSubset=[]`, `result=[]`.

1.  Exclude `A`
    - Decision: Continue to index `1` without appending `A`.
    - After: `currentIndex=1`, `currentSubset=[]`, `result=[]`.
    - Known correct: Subsets without `A` will be completed before
      subsets with `A`.
2.  Exclude `B`
    - Decision: Reach index `2` without appending `B`; save a copy.
    - After: `currentIndex=1`, `currentSubset=[]`, `result=[[]]`.
    - Known correct: The subset containing neither value is saved.
3.  Include `B`
    - Decision: Append `B`, save `[B]` at index `2`, then undo `B`.
    - After: `currentIndex=0`, `currentSubset=[]`,
      `result=[[],[B]]`.
    - Known correct: Both subsets without `A` are saved.
4.  Include `A`
    - Decision: Append `A` and continue to index `1`.
    - After: `currentIndex=1`, `currentSubset=[A]`,
      `result=[[],[B]]`.
    - Known correct: Every later result in this branch contains
      `A`.
5.  Exclude `B` after `A`
    - Decision: Reach index `2` without appending `B`; save a copy.
    - After: `currentIndex=1`, `currentSubset=[A]`,
      `result=[[],[B],[A]]`.
    - Known correct: The subset containing only `A` is saved.
6.  Include `B` after `A`
    - Decision: Append `B`, save `[A,B]`, then undo both active
      choices while returning.
    - After: `currentIndex=done`, `currentSubset=[]`,
      `result=[[],[B],[A],[A,B]]`.
    - Known correct: Every include-or-exclude choice appears once.

The workload grows before the input looks large. Exactly `20`
distinct values require `2^20 = 1,048,576` returned subsets. Their
contents hold `20*2^19 = 10,485,760` values in total, so even a
perfect generator must construct and store that much requested
output.

For empty input, there is one subset--the empty subset--so the output
is `[[]]`, not an empty list of answers.

#### Why it works

Before processing `currentIndex`, `currentSubset` contains exactly
the values selected from earlier positions. Every subset either
excludes the current position or includes it, but never both. The two
calls therefore do not repeat an answer, and together they cover every
choice. When `currentIndex` reaches the input length, the choices for
every position are fixed and one complete subset is copied. Because
each call advances `currentIndex`, the process finishes.

#### Go example

```go
func subsets(values []string) [][]string {
	// currentSubset holds the choices on the active recursion path;
	// result owns copies so later choices cannot mutate saved
	// subsets.
	result := make([][]string, 0)
	currentSubset := make([]string, 0, len(values))
	var build func(int)
	build = func(currentIndex int) {
		if currentIndex == len(values) {
			// Every input position has made its include/exclude decision.
			result = append(result, append([]string{}, currentSubset...))
			return
		}
		// First enumerate subsets that exclude values[currentIndex].
		build(currentIndex + 1)
		// Then include it, enumerate that branch, and undo the choice
		// before returning to the caller.
		currentSubset = append(currentSubset, values[currentIndex])
		build(currentIndex + 1)
		currentSubset = currentSubset[:len(currentSubset)-1]
	}
	build(0)
	return result
}
```

#### Work and memory

There are `2^n` returned subsets. Each of the `n` input values appears
in half of them, so copying the contents requires `n*2^(n-1)` element
copies when `n>0`. Time and required output space are therefore
`O(n*2^n)`. The unfinished calls and `currentSubset` use `O(n)` extra
memory, excluding returned output. A permutation generator likewise
cannot beat its `n!` output count. Storing all size-`k` combinations
requires `O(k*C(n,k))` output space, where `C(n,k)` is the number of
ways to choose `k` positions from `n` positions.

#### When to use it

Use combinatorial counting before implementing any "return all
choices" request, and when constraints are meant to tell you whether
exhaustive output is feasible.

#### When not to use it

Do not generate every object when the task asks only for a count, an
optimum, or one valid object. A counting formula, dynamic program, or
search that stops after one result may avoid the full output cost.

#### Common mistakes and edge cases

> **Warning:** Calling subset generation merely `O(2^n)` ignores
> copying. For `[A,B,C]`, eight rows are returned, but their contents
> hold 12 values in total.

> **Warning:** The empty input has one subset and one permutation--the
> empty choice. Returning no answers breaks the include-or-exclude
> rule used by a larger call.

Repeated values require a definition of distinct output: positions may
be distinct even when their values compare equal.

#### Related material

- [A.1 Complexity and resource
  accounting](#a1-complexity-and-resource-accounting) defines output
  and extra-memory accounting.
- [A.7.1 Counting combinations without large
  factorials](#a71-counting-combinations-without-large-factorials)
  computes `C(n,k)` exactly.
- [A.7.2 Repeated input values](#a72-repeated-input-values)
  distinguishes position choices from value choices.
- [C.1.3.19 Backtracking](#c1319-backtracking) explains the general
  choose-and-undo method.

#### A.7.1 Counting combinations without large factorials

##### The problem it solves

Given `n` distinct items and a group size `k`, return the number of
different unordered groups of exactly `k` items without generating
them. Computing `n!`, `k!`, and `(n-k)!` separately can overflow even
when their final quotient fits in the chosen integer type.

##### How it works

Choosing `2` people for a team of `5` is the same as choosing the `3`
people left out. The count is written `C(5,2)`, read "5 choose 2."
Build that count a little at a time instead of calculating large
factorials.

##### Step-by-step examples

For `C(5,2)`, the expected result is `10`.

Initial state: `n=5`, `requestedK=2`, `effectiveK=unknown`,
`numChosenSoFar=0`, `result=1`.

1.  Choose the shorter calculation
    - Decision: Keeping `2` items is shorter than naming the `3`
      omitted items, so two loop steps are enough.
    - After: `effectiveK=2`, `numChosenSoFar=1`, `result=1`.
    - Known correct: Every kept pair identifies one omitted group
      of three, and the reverse is also true.
2.  Choose the first item
    - Decision: Multiply by `4`, then divide by `1`.
    - After: `numChosenSoFar=2`, `result=4`.
    - Known correct: The stored count is `C(4,1)`.
3.  Choose the second item
    - Decision: Multiply by `5`, then divide by `2`.
    - After: `numChosenSoFar=3`, `result=10`.
    - Known correct: The stored count is `C(5,2)`.

For `C(30,2)`, the same loop stores `29` after its first step and
`435` after its second. The factorial formula tries to construct
`30!` before canceling; that intermediate is larger than `int64` even
though the final answer `435` is small.

Pascal's recurrence reaches the same answer by marking `E` in
`[A,B,C,D,E]`. A pair containing `E` chooses one partner from the
other four, giving `4` pairs. A pair without `E` chooses both members
from those four, giving `6` pairs. The two nonoverlapping cases total
`4+6=10`.

For boundary `k=0` or `k=n`, choosing the smaller side gives zero loop
iterations and the function returns `1`. Invalid `k<0` or `k>n`
returns failure.

##### Why it works

After step `i`, `result` counts ways to choose `i` items from
`n-k+i` items. The numerator factors count choices where the order of
the chosen items matters. The divisor factors `1` through `i` remove
the different orders of each same group. The quotient is therefore a
whole-number group count. Before any step, the one size-zero group is
the empty group. After step `k`, `n-k+k` is `n`, so the result counts
ways to choose the requested `k` items from all `n` items.

Pascal's recurrence is correct because every size-`k` group either
includes a marked value or excludes it. Those cases cannot overlap and
cover every group.

##### Go example

The mathematical input `n=5,k=2` maps to
`combinationCount(numItems=5, numChosen=2)`, which returns a
`big.Int` containing `10` and `true`. After the symmetry check,
the code's `numChosen` is the trace's `effectiveK`; its
`numChosenSoFar` and `result` have the same meanings as in the trace.

```go
import "math/big"

func combinationCount(numItems, numChosen int64) (*big.Int, bool) {
	// A combination count is defined only when the chosen count is
	// within the total item count.
	if numItems < 0 || numChosen < 0 || numChosen > numItems {
		return nil, false
	}
	// C(n,k) equals C(n,n-k); using the smaller side shortens the
	// product.
	if numChosen > numItems-numChosen {
		numChosen = numItems - numChosen
	}
	// Reuse big.Int temporaries and apply one numerator and denominator
	// factor per step. Every intermediate quotient is exact for this
	// recurrence.
	result := big.NewInt(1)
	factor := new(big.Int)
	divisor := new(big.Int)
	numChosenSoFar := int64(1)
	for numChosenSoFar <= numChosen {
		// Extend the count by one chosen item. This recurrence's division
		// is exact at every step, including the final requested count.
		result.Mul(
			result,
			factor.SetInt64(numItems-numChosen+numChosenSoFar),
		)
		result.Quo(result, divisor.SetInt64(numChosenSoFar))
		numChosenSoFar++
	}
	return result, true
}
```

##### Work and memory

The loop runs once for the smaller of `k` and `n-k`. This is
`O(min(k,n-k))` large-integer operations. An operation becomes more
expensive as the stored number gains digits. The result and temporary
values need enough memory to hold the digits of `C(n,k)`. A table
built with Pascal's rule calculates up to `n*k` entries. Keeping only
the previous row reduces its working storage to about `k` large
integers.

##### When to use it

Use the multiply-then-divide loop for one exact combination count. Use
Pascal's rule when many nearby counts are needed or when the problem
asks only for each count's remainder after division by a fixed value.

##### When not to use it

Do not generate all groups merely to count them. Do not use
fixed-width factorials when intermediate values can overflow, even if
the final count would fit.

##### Common mistakes and edge cases

> **Warning:** In fixed-width code, `result*factor` can overflow
> before the following exact division. Check or use arbitrary
> precision for the intermediate, not only the final answer.

Define invalid inputs and preserve `C(n,0)=C(n,n)=1`.

##### Related material

- [A.4 Integer limits and bits](#a4-integer-limits-and-bits) explains
  pre-operation overflow checks.
- [A.7 Combinatorics and output
  size](#a7-combinatorics-and-output-size) distinguishes counting from
  generation.
- [C.1.3.10 Dynamic programming](#c1310-dynamic-programming)
  generalizes Pascal's recurrence and rolling rows.

#### A.7.2 Repeated input values

##### The problem it solves

Given input values that may repeat, generate each distinct visible
group exactly once. Choosing equal values from different input
positions can create duplicate construction paths, so the algorithm
must distinguish position identity from the displayed value that
defines the output.

##### How it works

Temporarily label the two `1` values by their positions. Input
`[1a,1b,2]` has three pairs of positions, but pairs `[1a,2]` and
`[1b,2]` both display as `[1,2]`. Put equal values next to each other,
then allow only the first equal value to begin a particular part of
the search. A later equal value may still be chosen after the first;
otherwise `[1,1]` would disappear.

##### Step-by-step examples

For the actual input `values=[1,1,2]` and `k=2`, the expected value
outputs are `[[1,1],[1,2]]`. The labels `1a` and `1b` below refer only
to the first and second input positions; both stored values are `1`.

Initial state: `startIndex=0`, `candidateIndex=0`,
`currentCombination=[]`, `result=[]`.

1.  Choose the first `1`
    - Decision: Choose index `0` for the first output position.
    - After: `startIndex=1`, `candidateIndex=1`,
      `currentCombination=[1]`, `result=[]`.
    - Known correct: Every result in this branch begins with `1`.
2.  Choose the second `1`
    - Decision: This call fills a new output position, so choose
      the second `1`.
    - After: `startIndex=1`, `candidateIndex=2`,
      `currentCombination=[1]`, `result=[[1,1]]`.
    - Known correct: The two equal-valued positions can both be used.
3.  Choose `2`
    - Decision: In the same call, choose index `2`.
    - After: `startIndex=0`, `candidateIndex=1`,
      `currentCombination=[]`, `result=[[1,1],[1,2]]`.
    - Known correct: The other displayed pair beginning with `1`
      is included.
4.  Skip the second `1` at the first position
    - Decision: The first call already tried an equal `1` for this
      output position, so skip index `1`. Trying index `2` cannot fill
      the second position, so return.
    - After: `startIndex=0`, `candidateIndex=done`,
      `currentCombination=[]`, `result=[[1,1],[1,2]]`.
    - Known correct: Starting with the second `1` would repeat
      `[1,2]`, while `[1,1]` was already created in step 2.

For `k=0`, the one valid combination is empty. For `k>len(values)`, no
sequence of choices can reach length `k`, so the output is empty.

##### Why it works

Increasing indexes ensure that one input position is never used
twice. Within one call, two adjacent equal candidates would place the
same value in the same output position. Choosing the later one cannot
create a new displayed group, because it leaves fewer remaining
positions than choosing the earlier one. Keeping the first therefore
loses nothing. The next call fills a different output position, so it
may choose another copy of the value and preserve `[1,1]`. Every call
starts farther right, so the search finishes.

##### Go example

The trace's `k` is the code's `combinationLength`. Calling
`uniqueCombinations([]int{1, 1, 2}, 2)` returns
`[][]int{{1, 1}, {1, 2}}`. The code first copies and sorts `values`;
the traced input is already in that sorted order.

```go
import "sort"

func uniqueCombinations(values []int, combinationLength int) [][]int {
	if combinationLength < 0 || combinationLength > len(values) {
		return nil
	}
	// Sort a copy so equal values become adjacent without changing the
	// caller's slice. Adjacency makes duplicate choices easy to skip.
	sorted := append([]int(nil), values...)
	sort.Ints(sorted)
	result := make([][]int, 0)
	currentCombination := make([]int, 0, combinationLength)
	var build func(int)
	build = func(startIndex int) {
		if len(currentCombination) == combinationLength {
			// Save an independent copy because later search steps reuse the
			// active combination.
			result = append(result, append([]int{}, currentCombination...))
			return
		}
		candidateIndex := startIndex
		for candidateIndex < len(sorted) {
			if candidateIndex > startIndex &&
				sorted[candidateIndex] == sorted[candidateIndex-1] {
				// This call already tried the same value for this output
				// position, so trying it again would repeat answers.
				candidateIndex++
				continue
			}
			// Choose this value, fill later positions, then remove it so
			// the next candidate starts with the same earlier choices.
			currentCombination = append(
				currentCombination,
				sorted[candidateIndex],
			)
			build(candidateIndex + 1)
			currentCombination = currentCombination[:len(currentCombination)-1]
			candidateIndex++
		}
	}
	build(0)
	return result
}
```

##### Work and memory

Sorting `n` values takes `O(n log n)` time, and the sorted copy uses
`O(n)` extra memory. If `U` unique groups are returned, each has `k`
items, so copying the answers costs `O(k*U)` time and output space.
Before repeated choices are skipped, the search may visit every
partial group of size zero through `k`. The unfinished calls and
`currentCombination` hold at most `k` choices and use `O(k)` more
memory.

##### When to use it

Use same-call duplicate skipping when answers are groups of displayed
values, input positions may be used once, and equal input values
should not create repeated rows.

##### When not to use it

Do not remove duplicates when labeled positions are supposed to
produce distinct answers. Do not use this exact rule when a value may
be reused without limit, when order matters, or when only a count is
needed; those problems need different choice rules.

##### Common mistakes and edge cases

> **Warning:** Skipping every equal value everywhere loses `[1,1]`
> from `[1,1,2]`. Skip equal candidates only when the same call is
> choosing the same output position.

> **Warning:** Sorting the caller's slice in place changes observable
> input. The example copies before sorting.

Test all-equal input, `k=0`, `k=n`, and invalid `k`.

##### Related material

- [A.7 Combinatorics and output
  size](#a7-combinatorics-and-output-size) explains output-size
  accounting.
- [C.1.3.19 Backtracking](#c1319-backtracking) explains how to choose,
  recurse, and undo.
- [D.3 Sorting without unexpected
  mutation](#d3-sorting-without-unexpected-mutation) explains the
  copy-before-sort boundary.

<a id="integer-digit-processing"></a>

### A.8 Integer digit processing

#### The problem it solves

Given a nonnegative integer that fits in memory, inspect or rearrange
its decimal digits using arithmetic rather than text conversion. If
the operation builds a new integer, it must reject the next digit
before that digit would push the result above the allowed limit.

<a id="a81-mental-model"></a>

#### How it works

Peel digits from the right like cards from the top of a pile. Dividing
by `10` leaves a remainder from `0` through `9`; that remainder is the
last decimal digit. Whole-number division by `10` then removes the
last digit. Repeating those operations is **integer digit
processing**.

<a id="a84-preventing-integer-overflow"></a>

#### Step-by-step examples

For `value=120` and `limit=127`, the expected result is `(21,true)`.

Initial state: `value=120`, `reversed=0`, `valid=true`.

1.  Extract `0`
    - Decision: Append last digit `0`, then remove it from `value`.
    - After: `value=12`, `reversed=0`, `valid=true`.
    - Known correct: One source digit has moved.
2.  Extract `2`
    - Decision: Append `2`, then remove it from `value`.
    - After: `value=1`, `reversed=2`, `valid=true`.
    - Known correct: Removed suffix `20` is reversed as integer
      `2`.
3.  Extract `1`
    - Decision: Append `1`, then remove it from `value`.
    - After: `value=0`, `reversed=21`, `valid=true`.
    - Known correct: Every digit has moved in reverse order.

An integer cannot retain a leading zero, so the result is `21`, not
`021`. For the separate input `821` with limit `127`, the first two
steps produce `reversed=12`. Appending the final digit `8` is rejected
because `(127-8)/10` is `11` and `12` is larger. The unsafe value
`128` is never evaluated.

<a id="a83-invariant-fact-that-remains-true"></a>

#### Why it works

After every pass, `value` is the original number with the processed
ending digits removed. At the same time, `reversed` contains those
removed digits in opposite order. Extraction and append move exactly
one digit and preserve both facts. Each pass divides nonzero `value`
by `10`, so the process finishes. At zero, all digits have moved and
`reversed` is the complete answer.

When `digit` is no larger than `limit`, subtracting it is safe. The
second check is the rearranged form of requiring
`reversed*10+digit` to be no larger than `limit`. It checks the same
condition without first calculating the possibly excessive result.
Therefore every performed append stays within the allowed range.

#### Go example

Calling `reverseWithinLimit(120, 127)` follows the ordinary trace and
returns `(21,true)`. Calling it with the focused overflow input
`reverseWithinLimit(821, 127)` returns `(0,false)` before constructing
`128`.

```go
func reverseWithinLimit(value, limit int64) (int64, bool) {
	// This entry accepts only the nonnegative values used by the trace.
	if value < 0 || limit < 0 {
		return 0, false
	}
	var reversed int64
	for value > 0 {
		digit := value % 10
		// Rearrange reversed*10+digit <= limit so the unsafe
		// multiplication is never evaluated when it would overflow.
		if digit > limit || reversed > (limit-digit)/10 {
			return 0, false
		}
		reversed = reversed*10 + digit
		value /= 10
	}
	return reversed, true
}
```

The full fixed-`int64` implementation is owned by [E.6 Overflow-safe
integer reversal](#e6-overflow-safe-integer-reversal).

<a id="a85-work-and-memory"></a>

#### Work and memory

If `value` has `d` decimal digits, the loop performs `d` extraction
and append steps, so it takes `O(d)` time. It stores the remaining
value, one digit, the limit, and the result. That fixed amount of
storage is `O(1)` extra memory. No string or digit slice is allocated.

#### When to use it

Use arithmetic digit processing when the value already fits an integer
type, the task works on base-10 digits, and right-to-left order is
convenient.

#### When not to use it

Do not use it for a decimal input too large for the integer type; keep
that input as text. Do not use base-10 extraction for a bit-level
task, and do not assume the same sign behavior for negative values
without defining it.

#### Common mistakes and edge cases

> **Warning:** Check overflow before `reversed*10+digit`. For limit
> `127`, constructing `128` first defeats the guard.

> **Warning:** A loop written only as `for value>0` still needs an
> explicit rule for zero and negative input. The example returns
> `(0,true)` for zero and failure for negatives.

Trailing zeroes disappear from the displayed reversal, palindromes
remain the same, and the smallest signed integer cannot be made
positive in the same type.

<a id="a86-related-problems"></a>

#### Related material

- [A.4 Integer limits and bits](#a4-integer-limits-and-bits) derives
  the safe decimal-append check.
- [A.6.7 Numbers too large for an integer
  type](#a67-numbers-too-large-for-an-integer-type) handles digit text
  that cannot first become an integer.
- [E.6 Overflow-safe integer
  reversal](#e6-overflow-safe-integer-reversal) provides the complete
  Go implementation.
- [Palindrome
  Number](problems/solved_easy_leetcode_0009_palindrome_number.go)
  compares full reversal with reversing only half the digits.

<a id="b-data-structures"></a>

## B. Ways to store and organize information

### B.1 Arrays, slices, and strings

#### The problem it solves

Given an ordered sequence, store it so code can read or replace a
value by position and can pass around a consecutive subrange. In Go,
the choice among arrays, slices, and strings also determines whether
the size can change, whether writes are allowed, and whether two views
can modify the same underlying storage.

#### How it works

Picture a row of numbered lockers. An **array** owns the whole fixed
row. A **slice** is a movable window onto neighboring lockers; it
records a starting point, a visible length, and a capacity that limits
how far it may grow before `append` needs new storage. Two windows may
still share the same **underlying array**. A Go **string** is a
read-only row of bytes.

#### Step-by-step examples

Start with `base := []int{10, 20, 30, 40}`. The expression
`base[1:3]` includes index `1` and stops before index `3`, so
`middle := base[1:3]` initially shows `[20, 30]`. After the writes and
appends below, the expected values are
`base == [10, 99, 30, 50]` and `middle == [7, 30, 50, 60]`.

Initial state: `base=[10,20,30,40]`, `middle=unset`,
`storage=baseArray`.

1.  Create the view
    - Decision: Slice `base[1:3]`, exposing positions `1` and `2`.
    - After: `base=[10,20,30,40]`, `middle=[20,30]`, `len=2`,
      `cap=3`, `storage=baseArray`.
    - Known correct: Both slices reach `base[1]` and `base[2]`.
2.  Replace through the view
    - Decision: Set `middle[0]=99` in shared storage.
    - After: `base=[10,99,30,40]`, `middle=[99,30]`, `len=2`,
      `cap=3`, `storage=baseArray`.
    - Known correct: `middle[0]` and `base[1]` are one location.
3.  Append `50`
    - Decision: Reuse the one unused position within capacity.
    - After: `base=[10,99,30,50]`, `middle=[99,30,50]`,
      `len=3`, `cap=3`, `storage=baseArray`.
    - Known correct: The append overwrote `base[3]`.
4.  Append `60`
    - Decision: Length equals capacity, so allocate new storage.
    - After: `base=[10,99,30,50]`, `middle=[99,30,50,60]`,
      `len=4`, `storage=middleArray`.
    - Known correct: Later `middle` writes cannot reach `base`.
5.  Replace through the new storage
    - Decision: Set `middle[0]=7`.
    - After: `base=[10,99,30,50]`, `middle=[7,30,50,60]`,
      `storage=middleArray`.
    - Known correct: The two slices no longer share storage.

Boundary example: input values `[2, 1, 2]`, where every value must be
between `0` and `2`, produce frequency slice `[0, 1, 2]`. Value `-1`
or `3` breaks that promise and cannot be used as an index into a
length-3 slice.

#### Why it works

A slice operation is predictable when one fact remains true: every
visible position maps to its start position plus the index in one
underlying array. Creating a subslice changes the view, not that
mapping. A write follows the mapping to shared storage. Once `append`
allocates, the returned slice maps to a different array, so the old
view cannot observe later writes. A frequency slice works because each
permitted value has exactly one reserved index.

#### Go example

```go
func sliceSharing() (base []int, middle []int) {
	// middle is a two-element view whose capacity still reaches base's
	// final element, so both slices initially share the same backing
	// array.
	base = []int{10, 20, 30, 40}
	middle = base[1:3]
	// This indexed write is visible through base because storage is
	// shared.
	middle[0] = 99
	// The first append fits the remaining capacity and overwrites
	// base[3].
	middle = append(middle, 50)
	// The second append exceeds capacity, so append allocates new
	// storage and copies middle. Later writes through middle no longer
	// affect base.
	middle = append(middle, 60)
	// This write changes only the newly allocated middle array,
	// demonstrating that the second append ended storage sharing.
	middle[0] = 7
	return base, middle
}
```

Indexing a string with `text[i]` similarly returns one byte, not
necessarily a complete visible character.

#### Work and memory

Indexing and taking a subslice do a fixed amount of work and copy no
elements. They take `O(1)` time and `O(1)` extra memory. If `n` is
the number of copied values, copying takes `O(n)` time and `O(n)`
result space. One `append` usually takes constant time, but a growth
step can copy every current value. Across many appends, those
occasional copies average to `O(1)` time per appended value. This
average across a sequence of operations is called **amortized time**.

If `r` is the number of possible values, a frequency slice uses
`O(r)` memory. Counting `n` input values and preparing or reading all
`r` count positions takes `O(n + r)` time.

#### When to use it

- Positions are meaningful and direct indexed access is needed.
- Values must remain in order or sit next to one another in memory.
- The input range is small enough for a frequency slice.
- A cheap view of consecutive existing values is safe and useful.

#### When not to use it

- Use a map when keys are sparse, negative, or too large for practical
  indexed storage.
- Use linked nodes when frequent reconnection of already-known
  positions matters more than indexed access.
- Copy a subslice when it must outlive a large backing array or must
  not share future mutations.

#### Common mistakes and edge cases

> **Warning:** With `base := []int{1, 2, 3}` and `part := base[:2]`,
> appending `9` to `part` can overwrite `base[2]`; available capacity
> means `append` need not isolate the result.

- Indexing an empty slice or using an index equal to its length
  panics.
- A small long-lived slice can retain a much larger underlying array.
  Copy the needed values when the large storage should become
  reclaimable.
- `text[0]` reads one byte. For `text := "\u00e9"`, that byte is not
  the whole character; choose bytes, runes, or user-visible characters
  deliberately.
- A nil slice and a non-nil empty slice both have length zero but can
  differ in JavaScript Object Notation (JSON) output and in an
  application's public input/output contract, or **application
  programming interface (API)**.

#### Related material

- [Moving values in place](#b11-moving-values-in-place) changes
  existing slice storage deliberately.
- [Slice storage and retained
  memory](#d9-slice-storage-and-retained-memory) explains the Go
  allocation discussion.
- [Bytes, runes, and visible
  characters](#d6-bytes-runes-and-visible-characters) explains text
  indexing units.
- [Map presence and
  sets](#d4-map-presence-and-sets)
  contrasts dense indexed counts with sparse keys.

#### B.1.1 Moving values in place

##### The problem it solves

Given a mutable sequence and a required rearrangement, produce the
result inside the input's existing storage because a second sequence
of similar size is not allowed. Each unread value must remain
available until it has been copied, swapped, or deliberately removed.

##### How it works

Move furniture within one room instead of renting another room. An
algorithm works **in place** when it reuses the input's storage and
needs only a fixed number of separate variables. Reversing parts of a
slice can rotate its values without building another slice.

##### Step-by-step examples

Rotate `[1, 2, 3, 4, 5]` right by `7`; the expected output is
`[4, 5, 1, 2, 3]`.

Initial state: `values=[1,2,3,4,5]`, `numPositions=7`, `n=5`.

1.  Normalize the rotation
    - Decision: Rotations repeat every five moves; compute `7%5`.
    - After: `values=[1,2,3,4,5]`, `numPositions=2`.
    - Known correct: The requested result is unchanged.
2.  Reverse all values
    - Decision: Bring the final two values to the front in reverse
      order.
    - After: `values=[5,4,3,2,1]`, `numPositions=2`.
    - Known correct: The two result groups are now on the correct
      sides, although their internal orders are reversed.
3.  Reverse the prefix
    - Decision: Restore the moved group's internal order.
    - After: `values=[4,5 | 3,2,1]`, `numPositions=2`.
    - Known correct: Output positions `0` and `1` are final.
4.  Reverse the suffix
    - Decision: Restore the other group's internal order.
    - After: `values=[4,5,1,2,3]`, `numPositions=2`.
    - Known correct: Every output position is final.

Boundary example: rotating `[]` by any number of positions leaves it
empty. The function must return before `numPositions %= len(values)`,
because remainder by zero is invalid.

##### Why it works

Call the first `n-k` values group `A` and the final `k` values group
`B`. The slice starts as `A B`. Reversing everything produces the
values of `B` backward followed by the values of `A` backward.
Reversing those two parts separately restores each part's original
order and leaves `B A`, exactly the requested rotation. Each swap
moves the two active positions inward, so every reversal finishes.

##### Go example

The function requires `numPositions >= 0`.

```go
func rotateRight(values []int, numPositions int) {
	// Avoid remainder by zero; an empty slice is already fully rotated.
	if len(values) == 0 {
		return
	}
	// Rotating by a whole number of slice lengths has no effect.
	numPositions %= len(values)
	// reverse swaps the inclusive range from both ends toward the
	// center.
	reverse := func(left, right int) {
		for left < right {
			values[left], values[right] = values[right], values[left]
			left++
			right--
		}
	}
	// If the original slice is A B and len(B)==numPositions, these
	// reversals change A B -> reverse(B) reverse(A) -> B A.
	reverse(0, len(values)-1)
	reverse(0, numPositions-1)
	reverse(numPositions, len(values)-1)
}
```

##### Work and memory

Each value participates in a constant number of swaps, so rotating `n`
values takes `O(n)` time in best, usual, and worst cases. The indices
and temporary swap value use `O(1)` extra memory. The input itself
holds the result.

##### When to use it

- The caller permits input mutation.
- Extra memory must stay constant while values are rearranged.
- The output contains the same values in a new order, called a
  **permutation**, or is a filtered prefix of the input.

##### When not to use it

- Do not mutate when callers still need the original order; allocate a
  result.
- Do not call an operation constant-memory when recursive calls can
  leave `O(n)` unfinished calls on the call stack.
- Prefer a copy-based method when equal values must remain in their
  original relative order and the in-place method does not preserve
  that order. Preserving it is called **stability**.

##### Common mistakes and edge cases

> **Warning:** On an empty slice, `numPositions %= len(values)`
> divides by zero. Test the length first.

- Normalize a negative `numPositions` separately if the input rules
  permit it; Go's remainder keeps the dividend's sign.
- Reversing `[left, right]` requires inclusive endpoints; mixing that
  with a half-open range can skip or repeat an element.

##### Related material

- [Arrays, slices, and strings](#b1-arrays-slices-and-strings)
  explains the storage-sharing rules used here.
- [Two pointers](#two-pointers) explains read/write positions and
  inward movement as a problem-solving method.
- [Sorting without unexpected
  mutation](#d3-sorting-without-unexpected-mutation) explains when to
  copy before an in-place library operation.

<a id="linked-lists"></a>

### B.2 Linked lists

<a id="b21-mental-model"></a> <a id="b22-core-operations"></a>
<a id="b23-worked-example-reversing-a-list"></a>
<a id="b24-invariants-facts-that-remain-true"></a>
<a id="b25-common-techniques"></a> <a id="b26-work-and-memory"></a>
<a id="b27-common-mistakes"></a> <a id="b28-related-problems"></a>

#### The problem it solves

Given ordered values that must frequently gain, lose, or reconnect
items, preserve the order without shifting every later value in
contiguous storage. Each item instead stores a link to the next item,
which makes local reconnection cheap but requires following links to
reach a distant position.

#### How it works

Think of a treasure hunt in which each note points to the next note. A
**singly linked list** is a chain of nodes. Each node stores one value
and a link named `Next`; the link holds the location of the following
node. The first node is the **head**, and `nil` means there is no next
node. A **doubly linked list** also has a link to the previous node.

#### Step-by-step examples

Reverse `1 -> 2 -> 3 -> nil`; the expected output is
`3 -> 2 -> 1 -> nil`.

Initial state: `previous=nil`, `current=1 -> 2 -> 3 -> nil`.

1.  Reverse node `1`
    - Decision: Save `2`, point `1.Next` to `nil`, and advance.
    - After: `previous=1 -> nil`, `current=2 -> 3 -> nil`.
    - Known correct: Node `1` is a complete reversed prefix.
2.  Reverse node `2`
    - Decision: Save `3`, point `2.Next` to `1`, and advance.
    - After: `previous=2 -> 1 -> nil`, `current=3 -> nil`.
    - Known correct: Nodes `1,2` are reversed without losing `3`.
3.  Reverse node `3`
    - Decision: Save `nil`, point `3.Next` to `2`, and advance.
    - After: `previous=3 -> 2 -> 1 -> nil`, `current=nil`.
    - Known correct: Every original node appears once in reverse
      order.

For local insertion, start with a held link to node `2` in
`1 -> 2 -> 4 -> nil`; insert node `3` after it.

Initial state: `current=2`, `newNode=3`,
`list=1 -> 2 -> 4 -> nil`.

1.  Reconnect two links
    - Decision: Point `newNode.Next` to `4`, then point
      `current.Next` to `newNode`.
    - After: `current=2`, `newNode=3`,
      `list=1 -> 2 -> 3 -> 4 -> nil`.
    - Known correct: Insertion changed two links and moved no existing
      node. A slice insertion at the same position must shift the
      later values to open a slot.

Boundary example: reversing `nil` performs no loop iteration and
returns `nil`; reversing one node saves `nil`, leaves its link `nil`,
and returns the same node.

#### Why it works

During reversal, `previous` is exactly the already reversed part and
`current` is the first unvisited node. Saving `current.Next` preserves
access to every unvisited node; redirecting `current.Next` adds
exactly one node to the reversed part. The unvisited part shrinks, so
the loop finishes. When `current` is nil, no nodes remain outside
`previous`, making `previous` the new head.

#### Go example

The complete reusable implementation is owned by [E.4 Generic singly
linked list](#e4-generic-singly-linked-list). Its `Reverse` function
uses the same `previous`, `current`, and `next` roles as this trace.
Calling `FromSlice([]int{1, 2, 3})` and passing its result to
`Reverse` returns the traced `3 -> 2 -> 1 -> nil` list.

#### Work and memory

Visiting or searching `n` nodes takes `O(n)` time and `O(1)` extra
space. Inserting after a node already in hand takes `O(1)` time.
Loop-based reversal takes `O(n)` time and `O(1)` extra memory.
Copying all values to or from a slice uses `O(n)` time and `O(n)`
result space. A recursive walk can also retain `O(n)` unfinished
calls.

#### When to use it

- Nodes at known positions must be inserted, removed, or reconnected
  often.
- Other code holds links to particular nodes and those same node
  objects must survive while their order changes.
- Sequential access is sufficient and direct indexing is unnecessary.

#### When not to use it

- Use a slice when code must jump directly to a numbered position or
  when tightly packed storage and fast sequential reading matter.
- Do not use a plain linear walk on an input that may contain a cycle
  unless the walk tracks or detects repeated nodes.
- Prefer a loop when the list could create too many unfinished
  recursive calls for the call stack.

#### Common mistakes and edge cases

> **Warning:** In `1 -> 2 -> 3`, setting `1.Next = nil` before saving
> the old link loses the only route to nodes `2` and `3`.

- Conditions that read `node.Next` must first prove `node != nil`;
  empty and one-node inputs expose this error.
- Advancing the only head variable can make earlier nodes unreachable
  when the original head is still needed.
- Fast and slow positions need a loop condition that proves both the
  fast node and its next link exist before moving two steps.

#### Related material

- [Generic singly linked list](#e4-generic-singly-linked-list)
  provides the complete reusable Go implementation.
- [Fast and slow pointers](#c1332-fast-and-slow-pointers) uses
  different speeds to detect a cycle or locate a middle node.
- [Add Two
  Numbers](problems/solved_easy_leetcode_0002_add_two_numbers.go)
  combines unequal digit lists and carry state.
- [Merge Two Sorted
  Lists](problems/solved_easy_leetcode_0021_merge_two_sorted_lists.go)
  uses a dummy head and a maintained tail.

### B.3 Queues and deques

#### The problem it solves

Given values that arrive over time, either remove them in the same
order they arrived or support additions and removals at both ends.
The representation must avoid shifting every remaining value whenever
the front item is removed.

#### How it works

A **queue** is a ticket line: the oldest arrival leaves first. This
rule is **first in, first out (FIFO)**. A **deque**, short for
double-ended queue, is a line with a door at each end, so values can
enter or leave at either end.

A queue stored in a slice can remember the next unread position
instead of deleting the first element. The lookup contract is
`nextQueueIndex -> oldest unread value`. For example, if
`values=["", "B"]` and `nextQueueIndex=1`, the empty string marks a
cleared slot and `"B"` is the next value to remove.

#### Step-by-step examples

Add `A` and `B` to an empty queue, remove one value, add `C`, and
remove the rest. The expected removal order is `A`, `B`, `C`.

Initial state: `values=[]`, `nextQueueIndex=0`, `removed=[]`.

1.  Add `A`
    - Decision: Append the new arrival at the back.
    - After: `values=[A]`, `nextQueueIndex=0`, `removed=[]`.
    - Known correct: `A` is the oldest unread value.
2.  Add `B`
    - Decision: Append `B` behind `A`.
    - After: `values=[A,B]`, `nextQueueIndex=0`, `removed=[]`.
    - Known correct: The unread queue is `[A,B]`.
3.  Remove the front
    - Decision: Return `values[0]`, clear it, and advance the
      unread position.
    - After: `values=["",B]`, `nextQueueIndex=1`, `removed=[A]`.
    - Known correct: The unread queue is `[B]`.
4.  Add `C`
    - Decision: Append `C` behind every unread value.
    - After: `values=["",B,C]`, `nextQueueIndex=1`,
      `removed=[A]`.
    - Known correct: The unread queue is `[B,C]`.
5.  Remove `B`, then `C`
    - Decision: Apply the front-removal rule twice; reset when empty.
    - After: `values=[]`, `nextQueueIndex=0`, `removed=[A,B,C]`.
    - Known correct: Removal order equals arrival order.

For a deque, start with `[B]`. Add at both ends, then remove at both
ends. The expected removed values are `C`, then `A`; `[B]` remains.

Initial state: `deque=[B]`, `removed=[]`.

1.  Add `A` at the front
    - Decision: Place `A` before the current front.
    - After: `deque=[A,B]`, `removed=[]`.
    - Known correct: `A` is the front and `B` is the back.
2.  Add `C` at the back
    - Decision: Place `C` after the current back.
    - After: `deque=[A,B,C]`, `removed=[]`.
    - Known correct: The earlier values keep their order.
3.  Remove from the back
    - Decision: Remove `C` at the back boundary.
    - After: `deque=[A,B]`, `removed=[C]`.
    - Known correct: The front side is unchanged.
4.  Remove from the front
    - Decision: Remove `A` at the front boundary.
    - After: `deque=[B]`, `removed=[C,A]`.
    - Known correct: The expected removals and remainder are
      final.

An empty queue or deque removal must report failure.

#### Why it works

The unread slice begins in arrival order. Appending puts a new value
after every unread value, and advancing `nextQueueIndex` removes only
the oldest one. Both operations therefore preserve arrival order. A
deque applies the same rule independently at its two boundaries:
changing one end does not reorder the values between the ends.

#### Go example

The zero value of `stringQueue` is empty and ready to use.

```go
type stringQueue struct {
	values         []string
	nextQueueIndex int
}

func (queue *stringQueue) add(value string) {
	// New arrivals go behind every unread value.
	queue.values = append(queue.values, value)
}

func (queue *stringQueue) remove() (string, bool) {
	if queue.nextQueueIndex == len(queue.values) {
		return "", false
	}
	// Read and clear the oldest unread slot before moving the boundary.
	value := queue.values[queue.nextQueueIndex]
	queue.values[queue.nextQueueIndex] = ""
	queue.nextQueueIndex++
	if queue.nextQueueIndex == len(queue.values) {
		// Reuse the slice after the queue becomes empty.
		queue.values = queue.values[:0]
		queue.nextQueueIndex = 0
	}
	return value, true
}
```

#### Work and memory

Advancing `nextQueueIndex` takes `O(1)` time. Across many additions,
slice growth averages `O(1)` time per value, although one growth can
copy all stored values. If `n` values have been stored since the last
reset, the backing slice can use `O(n)` memory. A circular-buffer or
linked-node deque can add or remove at either end in `O(1)` time and
uses `O(n)` memory for `n` stored values. Prepending to an ordinary Go
slice instead copies its existing values and takes `O(n)` time.

#### When to use it

- Work must happen oldest first.
- A search must process all items at one distance before the next
  distance.
- A changing range of candidates must remove expired values from the
  front. Use a deque when the back must also remove values that can no
  longer become an answer.

#### When not to use it

- Use a stack for most-recent-first work.
- Use a heap when priority, not arrival time, decides the next value.
- Do not use ordinary FIFO breadth-first search for minimum weighted
  cost when edge costs differ.

#### Common mistakes and edge cases

> **Warning:** Repeatedly doing `queue = queue[1:]` on a long-lived
> queue can keep the old backing array and removed pointer values
> alive. Clear removed pointers and periodically compact, or use a
> circular buffer.

> **Warning:** `append([]T{value}, deque...)` copies every existing
> value. It does not provide constant-time front insertion.

- Check for emptiness before front or back removal.
- A circular buffer must distinguish full from empty even when head
  and tail positions match.
- One operation can be slow even when the average cost across many
  operations is constant. This sequence-wide average is called an
  **amortized** cost.

#### Related material

- [Breadth-first search](#c1321-breadth-first-search) explains the
  full search method and shortest-edge correctness argument.
- [Monotonic deque](#c1329-monotonic-deque) removes expired and
  dominated candidates from opposite ends.
- [Stack](#c1313-stack) contrasts most-recent-first processing with
  a queue's oldest-first processing.
- [Slice storage and retained
  memory](#d9-slice-storage-and-retained-memory) explains why
  advancing a head does not necessarily release storage.

### B.4 Trees and binary-search trees

Trees organize nodes by parent-child links. The entries below separate
search, visit order, recursive results, validation, reconstruction,
and added parent links because each operation maintains a different
fact.

#### B.4.1 Terminology and search

##### The problem it solves

Given values connected as parent and child nodes, describe positions
such as the root, leaves, ancestors, and subtrees unambiguously. If
the tree also obeys the binary-search ordering rule, use comparisons
to choose one child subtree and rule out the other during lookup.

##### How it works

A **tree** is one connected group of nodes with no cycle: there is
exactly one path from the starting node to every other node. The
starting node is the **root**; nodes directly below another node are
its **children**; and the node directly above a child is its
**parent**. An **ancestor** is any node reached by repeatedly moving
to a parent. A node with no children is a **leaf**, and a node plus
everything below it is a **subtree**. A **binary-search tree (BST)**
has at most two children per node, with smaller values on the left and
larger values on the right.

##### Step-by-step examples

Search for `10` in the pictured tree; the expected result is `true`.

```text
    8
   / \
  4  12
     /
    10
```

Initial state: `target=10`, `current=8`, `result=pending`.

1.  Compare with `8`
    - Decision: `10 > 8`, so follow the right link.
    - After: `target=10`, `current=12`, `result=pending`.
    - Known correct: The subtree rooted at `4` cannot match.
2.  Compare with `12`
    - Decision: `10 < 12`, so follow the left link.
    - After: `target=10`, `current=10`, `result=pending`.
    - Known correct: Only `12`'s left subtree can still match.
3.  Compare with `10`
    - Decision: The values match, so return success.
    - After: `target=10`, `current=10`, `result=true`.
    - Known correct: The target exists.

Boundary example: searching an empty tree starts at `nil` and returns
`false`. Searching the pictured tree for `9` reaches `8`, `12`, `10`,
then a nil left link, proving `9` is absent under the ordering
promise.

##### Why it works

Before each comparison, every possible match lies in the current
subtree. The BST rule makes every value on the rejected side too small
or too large, so choosing one child preserves that fact. Each step
moves down one level; reaching the target succeeds and reaching nil
leaves no possible node.

##### Go example

The example uses this node shape and forbids duplicate values:

```go
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func containsBST(root *TreeNode, target int) bool {
	// The current subtree contains every value that could still match.
	current := root
	for current != nil {
		if target == current.Value {
			return true
		}
		if target < current.Value {
			// All right-subtree values are too large; keep only the left side.
			current = current.Left
		} else {
			// All left-subtree values are too small; keep only the right side.
			current = current.Right
		}
	}
	// A nil link means the only possible search path ended without a
	// match.
	return false
}
```

##### Work and memory

The **height** `h` is the number of nodes on the longest route from
the root down to a leaf. Search visits at most one node on
each level, so it takes `O(h)` time and `O(1)` extra memory. When an
`n`-node tree stays roughly even on both sides, its height is
`O(log n)`. A one-sided chain has height `n`, so search takes `O(n)`
time. The best case finds the target at the root in `O(1)` time.

##### When to use it

- Values remain ordered under the stated whole-subtree rule.
- The environment keeps the tree roughly even on both sides, or the
  exact parent-child shape matters.
- A lookup can reject an entire side after one comparison.

##### When not to use it

- An arbitrary binary tree is not searchable this way.
- Go has no standard balanced ordered map; sorting, a heap, or binary
  search on a slice may be a better fit.
- Use node identity rather than value when duplicate values represent
  distinct objects.

##### Common mistakes and edge cases

> **Warning:** A tree with root `8`, right child `10`, and `10`'s left
> child `4` passes direct-child comparisons but is not a BST because
> `4` lies in `8`'s right subtree.

- State whether duplicates are forbidden or assigned consistently to
  one side.
- Do not dereference the root before checking an empty tree.
- Height-based bounds must include the possible one-sided `O(n)` case.

##### Related material

- [BST validation and costs](#b44-bst-validation-and-costs) checks the
  full ordering promise used by search.
- [Binary search](#c136-binary-search) applies the same
  discard-one-side idea to indexed sorted data.
- [Validate Binary Search
  Tree](problems/medium_leetcode_0098_validate_binary_search_tree.go)
  exercises ancestor bounds.

#### B.4.2 Traversal order

##### The problem it solves

Given a tree and an operation to perform on every node, choose the
order in which nodes should be visited. The required result may need a
parent before its children, children before their parent, left values
in sorted order, or all nodes at one distance before deeper nodes.

##### How it works

Walking through a house can record a room before entering its two
wings, between the wings, after both wings, or one floor at a time.
Visiting every node this way is a **traversal**. The four orders are
called **preorder** (node, left, right), **inorder** (left, node,
right), **postorder** (left, right, node), and **level order**
(shallowest level before the next level).

##### Step-by-step examples

Use this tree; preorder should return `[1,2,4,3]`:

```text
    1
   / \
  2   3
 /
4
```

Initial state: `current=1`, `output=[]`, `pending=[1]`.

1.  Record the root
    - Decision: Append `1`, then visit its left child first.
    - After: `current=2`, `output=[1]`, `pending=[2,3]`.
    - Known correct: Root `1` precedes both child subtrees.
2.  Record node `2`
    - Decision: Append `2`, then visit its left child `4`.
    - After: `current=4`, `output=[1,2]`, `pending=[4,3]`.
    - Known correct: Node `2` precedes its descendant `4`.
3.  Record node `4`
    - Decision: Append `4`; its two `nil` child calls add nothing.
    - After: `current=3`, `output=[1,2,4]`, `pending=[3]`.
    - Known correct: The complete left subtree follows root `1`.
4.  Record node `3`
    - Decision: Append `3`; its two `nil` child calls add nothing.
    - After: `current=nil`, `output=[1,2,4,3]`, `pending=[]`.
    - Known correct: Every node appears once in node-left-right
      order.

The same tree gives inorder `[4,2,1,3]`, postorder `[4,2,3,1]`, and
level order `[1,2,3,4]`. Each required order is now visibly
different. An empty tree produces an empty result in every order.

##### Why it works

Each recursive call handles exactly one subtree. It visits that
subtree's root and handles each child subtree once in the order named
by the traversal. The nil case stops at missing children. Thus every
reachable node appears once and relative order follows the placement
of the root action. A level-order queue keeps first-added nodes first,
so every shallower level stays ahead of the deeper nodes it discovers.

##### Go example

The trace's `pending` list makes the recursive call stack visible:
its first node is the active call, followed by subtrees whose parent
calls are waiting. The Go runtime stores that list in stack frames
rather than in an explicit slice. The pictured nodes are the
`TreeNode` input to `preorder`, and the function returns the traced
`output=[1,2,4,3]`.

```go
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func preorder(root *TreeNode) []int {
	output := []int{}
	var visit func(*TreeNode)
	visit = func(node *TreeNode) {
		// Nil children end their branch without contributing a value.
		if node == nil {
			return
		}
		// Preorder records the node before recursively visiting left and
		// right. The shared order slice accumulates values across every
		// call.
		output = append(output, node.Value)
		visit(node.Left)
		visit(node.Right)
	}
	visit(root)
	return output
}
```

##### Work and memory

For `n` nodes, every traversal handles each node once and takes
`O(n)` time. A depth-first recursive traversal can leave one
unfinished call per tree level. For height `h`, that is `O(h)`
call-stack space. The returned order contains `n` values and uses
`O(n)` result space.
Level order uses a queue holding up to `w` nodes, where `w` is the
largest number of nodes on one level, plus the `O(n)` output.

##### When to use it

- Use preorder to emit a parent before descendants or pass state
  downward.
- Use inorder when BST values must appear sorted.
- Use postorder when a parent result depends on both child results.
- Use level order when depth or nearest distance controls processing.

##### When not to use it

- Inorder is not sorted for an arbitrary binary tree.
- Recursive traversal is unsafe for an extremely deep chain; use an
  explicit stack or loop.
- Level order does not by itself minimize weighted path cost.

##### Common mistakes and edge cases

> **Warning:** Appending the node between its recursive calls
> implements inorder, not preorder. On the example tree, the results
> are `[4,2,1,3]` and `[1,2,4,3]` respectively.

- Decide whether nil roots return nil or a non-nil empty slice when
  output representation matters.
- Mark graph-like nodes when queued if parent links can create cycles.
- Account for output storage separately from extra traversal
  storage.

##### Related material

- [Queues and deques](#b3-queues-and-deques) explains FIFO mechanics
  for level order.
- [Recursive return contracts](#b43-recursive-return-contracts) uses
  postorder to combine child answers.
- [Binary Tree Level Order Traversal](
  problems/medium_leetcode_0102_binary_tree_level_order_traversal.go
  )
  groups the queue by depth.

#### B.4.3 Recursive return contracts

##### The problem it solves

Given a tree question whose answer depends on both child subtrees,
define exactly what one recursive call returns to its parent. The
returned value must contain only information the parent may legally
combine; otherwise, values from separate branches can be joined into
a path or result that does not exist in the tree.

##### How it works

Give every recursive call the same small job and promise. Write down
exactly what one call returns for the subtree starting at its input
node. This promise is called a **recursive return contract**. For
height, the promise is "return the number of nodes on the longest path
from this node down to a leaf." A `nil` subtree returns zero.

##### Step-by-step examples

Use a tree whose root is `2`, whose left child is `1`, and whose right
child is `3`; node `3` has right child `4`. The expected height is
`3`.

Initial state: `active=treeHeight(2)`, `waiting=[]`,
`returned=pending`.

1.  Enter the left child
    - Decision: Root `2` is non-nil, so call `treeHeight(1)` before
      calculating `leftHeight`.
    - After: `active=treeHeight(1)`,
      `waiting=[node 2(leftHeight=pending)]`, `returned=pending`.
    - Known correct: The root frame is saved until its left result is
      available.
2.  Return from leaf `1`
    - Decision: Its two nil child calls return `0`; set
      `leftHeight=0`, `rightHeight=0`, and return `1`.
    - After: `active=treeHeight(2)`,
      `waiting=[]`, `returned=1`, `node 2.leftHeight=1`.
    - Known correct: Node `1`'s longest downward path has one node.
3.  Enter node `3`, then node `4`
    - Decision: Call `treeHeight(3)` for root `2`'s right child.
      Node `3`'s nil left child returns `0`, so call its right child
      `treeHeight(4)`.
    - After: `active=treeHeight(4)`,
      `waiting=[node 2(leftHeight=1),node 3(leftHeight=0)]`,
      `returned=pending`.
    - Known correct: Each waiting frame has saved every child result
      obtained so far.
4.  Return from leaf `4`
    - Decision: Its two nil child calls return `0`; return `1`.
    - After: `active=treeHeight(3)`,
      `waiting=[node 2(leftHeight=1)]`, `returned=1`,
      `node 3.leftHeight=0`, `node 3.rightHeight=1`.
    - Known correct: Node `4`'s longest downward path has one node.
5.  Return from node `3`
    - Decision: Its right height `1` exceeds its left height `0`, so
      return `1+1=2`.
    - After: `active=treeHeight(2)`, `waiting=[]`, `returned=2`,
      `node 2.leftHeight=1`, `node 2.rightHeight=2`.
    - Known correct: Node `3`'s longest path is `3 -> 4`.
6.  Return from root `2`
    - Decision: Its right height `2` exceeds its left height `1`, so
      return `1+2=3`.
    - After: `active=none`, `waiting=[]`, `returned=3`.
    - Known correct: The longest full-tree path is `2 -> 3 -> 4`.

Boundary example: a nil root returns `0`. A one-sided `1 -> 2 -> 3`
chain returns `3`, not `2`; every node on the path counts.

##### Why it works

Zero is the promised answer for an empty subtree. Assuming both child
calls keep the same promise, every downward path from the current node
continues through exactly one child, so the longer child height plus
one is the longest current path. Each call moves to a smaller subtree
and ends at nil.

##### Go example

```go
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func treeHeight(root *TreeNode) int {
	// An empty subtree contributes zero levels.
	if root == nil {
		return 0
	}
	// Recursively obtain the complete height for each independent child.
	leftHeight := treeHeight(root.Left)
	rightHeight := treeHeight(root.Right)
	// The current node adds one level above the taller child subtree.
	if leftHeight > rightHeight {
		return 1 + leftHeight
	}
	return 1 + rightHeight
}
```

##### Work and memory

For `n` nodes, every node contributes a fixed amount of work, so time
is `O(n)`. The call stack holds one unfinished call per level. For
height `h`, that is `O(h)` memory: `O(log n)` for a roughly even tree
and `O(n)` for a one-sided chain. The function allocates no result
collection.

##### When to use it

- A parent answer can be built from fixed-size child answers.
- The input is naturally split into smaller subtrees.
- A precise nil result and subtree promise can be stated first.

##### When not to use it

- Return two clearly named values when the final answer may join paths
  from both children but the path offered to the parent may follow
  only one child.
- Use an explicit stack when depth may exhaust the call stack.
- Do not use this tree rule on links that can form a cycle unless the
  algorithm records which nodes it has already reached.

##### Common mistakes and edge cases

> **Warning:** Returning `leftHeight + rightHeight + 1` for height
> incorrectly gives `4` at root `2` in the example; one downward path
> cannot take both children.

- Define whether height counts nodes or edges before choosing the nil
  value.
- Do not recompute the same child result in multiple branches.
- Include call-stack space in recursive complexity.

##### Related material

- [Recursion and saved work](#a3-recursion-and-saved-work) explains
  call structure and how to save repeated answers.
- [Traversal order](#b42-traversal-order) explains why child-first
  work is postorder.
- [Maximum Depth of Binary
  Tree](problems/easy_leetcode_0104_maximum_depth_of_binary_tree.go)
  uses this exact return promise.

#### B.4.4 BST validation and costs

##### The problem it solves

Given a binary tree, determine whether every node satisfies the full
binary-search-tree ordering rule. Comparing a node only with its
children misses violations against distant ancestors, so each child
must inherit the complete lower and upper bounds imposed by the path
from the root.

##### How it works

Think of filing numbers into nested drawers. Entering a left drawer
lowers the largest allowed number; entering a right drawer raises the
smallest. A **binary-search-tree invariant** says every number must
stay inside all limits set by the drawers above it.

##### Step-by-step examples

This tree is invalid:

```text
    8
     \
     10
     /
    4
```

Initial state: `node=8`, `lower=none`, `upper=none`,
`result=pending`.

1.  Validate root `8`
    - Decision: `8` fits its bounds; inspect its right subtree.
    - After: `node=10`, `lower=8`, `upper=none`,
      `result=pending`.
    - Known correct: Root `8` satisfies every restriction.
2.  Validate node `10`
    - Decision: `10 > 8`; inspect its left subtree with upper `10`.
    - After: `node=4`, `lower=8`, `upper=10`, `result=pending`.
    - Known correct: The path through `10` remains valid.
3.  Reject node `4`
    - Decision: `4` is not greater than lower bound `8`.
    - After: `node=4`, `lower=8`, `upper=10`, `result=false`.
    - Known correct: A descendant violates the bound from ancestor
      `8`.

An empty tree is valid because it contains no violation. If
duplicates are allowed, the input rules must say which comparisons
allow equality.

##### Why it works

The bounds describe exactly the restrictions imposed by every
ancestor. The root begins unrestricted. Replacing one bound with the
current value adds the new parent restriction without losing older
ones. If all nodes fit, every ancestor relation satisfies the BST
rule; one rejection supplies a concrete counterexample. The recursion
finishes at nil children.

##### Go example

```go
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func validBST(root *TreeNode) bool {
	// Each call carries optional exclusive bounds imposed by all
	// ancestors, not merely by the node's direct parent.
	var valid func(*TreeNode, int, bool, int, bool) bool
	valid = func(
		node *TreeNode,
		lower int,
		hasLower bool,
		upper int,
		hasUpper bool,
	) bool {
		// An empty subtree cannot violate either inherited bound.
		if node == nil {
			return true
		}
		if (hasLower && node.Value <= lower) ||
			(hasUpper && node.Value >= upper) {
			// Equality also fails because this contract requires unique
			// values.
			return false
		}
		// The current value becomes the left subtree's upper bound and the
		// right subtree's lower bound. Both subtrees must satisfy their
		// ranges.
		return valid(node.Left, lower, hasLower, node.Value, true) &&
			valid(node.Right, node.Value, true, upper, hasUpper)
	}
	// The root starts without either bound; the booleans make zero a
	// usable value.
	return valid(root, 0, false, 0, false)
}
```

##### Work and memory

Validation visits each node at most once: `O(n)` time. Its recursion
uses `O(h)` stack space. It can stop near the root at the first
violation, giving best-case `O(1)` time as `n` grows; a valid tree or
a late violation requires `O(n)` time.

##### When to use it

- The entire tree must be checked against a strict ordering rule.
- Descendants can violate restrictions established several levels
  earlier.
- Node values fit safely in the chosen bound representation.

##### When not to use it

- Reading values in inorder and checking that each value is larger
  than the preceding one may be simpler when only ordering matters.
- Adjust the method when duplicates have an allowed side or keys use a
  custom comparison.
- Do not use a fixed minimum or maximum marker that is also a valid
  node value. Optional bounds, as in the example, avoid that conflict.

##### Common mistakes and edge cases

> **Warning:** Checking only `node.Left.Value < node.Value` and
> `node.Right.Value > node.Value` accepts the pictured invalid node
> `4`.

- Bounds based on `int` endpoints can reject `math.MinInt` or
  `math.MaxInt` as valid values; use wider bounds or optional bounds.
- Strict and inclusive comparisons must match the duplicate policy.
- A valid empty tree should not dereference a nil root.

##### Related material

- [Terminology and search](#b41-terminology-and-search) states the
  ordering promise being validated.
- [Correctness and invariants](#a2-correctness-and-invariants) gives
  the general proof pattern.
- [Validate Binary Search
  Tree](problems/medium_leetcode_0098_validate_binary_search_tree.go)
  is the representative exercise.

#### B.4.5 Reconstruction from traversals

##### The problem it solves

Given two traversal sequences that uniquely describe a tree, rebuild
its parent-child links. Each recursive range needs to locate its root
in the other traversal, and repeated linear searches for those roots
can turn an otherwise linear reconstruction into quadratic work.

##### How it works

Preorder names a subtree's root first. Inorder places every
left-subtree value before that root and every right-subtree value
after it. One order chooses the root; the other cuts the remaining
values into the two child regions. Build the lookup
`inorderIndexByValue[value] -> position in inorder`. For example,
inorder `[1,2,3]` produces `{1:0,2:1,3:2}`, so looking up value `2`
returns its split position `1`.

##### Step-by-step examples

Exact input:

```text
preorder: [2, 1, 3]
inorder:  [1, 2, 3]
```

The expected tree has root `2`, left child `1`, and right child `3`.
Scanning the supplied inorder input creates the complete lookup
`inorderIndexByValue={1:0,2:1,3:2}` before reconstruction begins.

Initial state: `nextPreorder=0`, `pending=[build(0,3)]`,
`tree=empty`.

1.  Build range `[0,3)`
    - Decision: Take preorder value `2`; inorder index `1` splits it.
    - After: `nextPreorder=1`,
      `pending=[build(0,1),build(2,3)]`, `tree=2`.
    - Known correct: Values left and right of index `1` belong to the
      corresponding child subtrees.
2.  Build range `[0,1)`
    - Decision: Take preorder value `1`; inorder index `0` splits it.
    - After: `nextPreorder=2`,
      `pending=[build(0,0),build(1,1),build(2,3)]`,
      `tree=2(left=1)`.
    - Known correct: Node `1` occupies the root's left range.
3.  Finish node `1`'s empty ranges
    - Decision: Both ranges have equal endpoints, so return `nil`.
    - After: `nextPreorder=2`, `pending=[build(2,3)]`,
      `tree=2(left=1)`.
    - Known correct: Node `1` is a leaf; no preorder value was used.
4.  Build range `[2,3)`
    - Decision: Take preorder value `3`; inorder index `2` splits it.
    - After: `nextPreorder=3`,
      `pending=[build(2,2),build(3,3)]`,
      `tree=2(left=1,right=3)`.
    - Known correct: Node `3` occupies the root's right range.
5.  Finish node `3`'s empty ranges
    - Decision: Both ranges have equal endpoints, so return `nil`.
    - After: `nextPreorder=3`, `pending=[]`,
      `tree=2(left=1,right=3)`.
    - Known correct: Every traversal value has one final tree node.

The lookup removes visible repeated work on a one-sided tree. For
`preorder=[5,4,3,2,1]` and `inorder=[1,2,3,4,5]`, scanning each
remaining inorder range from the left needs `5+4+3+2+1=15`
comparisons. Building `inorderIndexByValue` scans the five values
once, then each node obtains its split with one map lookup.

Boundary example: two empty traversals produce nil. Duplicate
traversals `preorder=[1,1]`, `inorder=[1,1]` can describe either a
left or right child and therefore do not identify one tree without
unique IDs or another rule.

##### Why it works

The recursive fact is that `[left,right)` contains exactly the inorder
values for the next preorder subtree. Preorder supplies that subtree's
root; its known inorder position uniquely divides the remaining nodes.
Building the entire left range consumes exactly the preorder values
before the right subtree. Ranges strictly shrink and empty ranges
stop.

##### Go example

The two slices must have equal lengths, contain the same unique
values, and describe a valid tree. The trace's `preorder` and
`inorder` are the code's `preorderValues` and `inorderValues`.
Calling `buildTree([]int{2, 1, 3}, []int{1, 2, 3})` returns the
pictured tree. The trace's `pending` ranges are recursive calls held
by Go's call stack rather than an explicit collection in the code.

```go
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorderValues, inorderValues []int) *TreeNode {
	// Map each unique value to its inorder position so subtree splits
	// are O(1).
	inorderIndexByValue := make(map[int]int, len(inorderValues))
	for inorderIndex, value := range inorderValues {
		// The input contract promises unique values, so each position has
		// one key.
		inorderIndexByValue[value] = inorderIndex
	}
	// Preorder supplies roots in the order recursive calls need them.
	nextPreorder := 0
	var build func(int, int) *TreeNode
	build = func(left, right int) *TreeNode {
		// [left,right) is empty, so this child link is nil.
		if left == right {
			return nil
		}
		// The next preorder value is this subtree's root. Its inorder
		// position divides the remaining values into left and right
		// subtrees.
		value := preorderValues[nextPreorder]
		nextPreorder++
		rootInorderIndex := inorderIndexByValue[value]
		root := &TreeNode{Value: value}
		// Values before rootInorderIndex belong to the left subtree; values
		// after it belong to the right. Recursive preorder consumption
		// supplies each child root.
		root.Left = build(left, rootInorderIndex)
		root.Right = build(rootInorderIndex+1, right)
		return root
	}
	// The initial half-open range covers the entire inorder traversal.
	return build(0, len(inorderValues))
}
```

##### Work and memory

Building the index and nodes each takes `O(n)` time. The index map
uses `O(n)` auxiliary space, and recursion uses another `O(h)` stack
space. The returned `n` nodes require `O(n)` output space. Repeatedly
scanning inorder instead makes worst-case time `O(n^2)`.

##### When to use it

- Preorder and inorder contain the same unique node identifiers.
- The task requires the explicit tree, not only a property of the
  traversals.
- A map from values to inorder positions fits the input constraints.

##### When not to use it

- Duplicate values make reconstruction ambiguous without extra
  identity rules.
- Invalid or mismatched traversals need explicit validation before the
  compact routine shown here.
- If only one traversal is given, many different trees may match it.

##### Common mistakes and edge cases

> **Warning:** Building the right subtree before the left consumes
> preorder in the wrong order. For preorder `[2,1,3]`, value `1`
> belongs to the left side.

- Keep half-open boundaries consistent; the right range starts at
  `split+1`.
- Verify a looked-up split lies inside the current range for untrusted
  input.
- The index map silently overwrites duplicate keys.

##### Related material

- [Traversal order](#b42-traversal-order) defines preorder and
  inorder.
- [Map presence and
  sets](#d4-map-presence-and-sets)
  explains checked lookups for untrusted values.
- [Construct Binary Tree from Preorder and Inorder Traversal](
  ../REFERENCE_LINKS.md#leetcode-0105
  )
  is the representative exercise.

#### B.4.6 Traversing through parent links

##### The problem it solves

Given a tree whose nodes point only to children and a query that
starts at an internal target, visit nodes reachable by moving both
down to children and up to ancestors. The algorithm must recover the
missing upward connections and remember visited nodes so the new
two-way movement does not bounce along the same edge forever.

##### How it works

Record a temporary parent link for every child. Then each node can
move in three possible directions: to its left child, right child, or
parent. This turns the tree into an undirected graph for the query. A
set of already reached node locations prevents walking back and forth
over the same link. The contracts are
`parent[child] -> that child's parent node` and
`seen[node] -> whether that node has entered the queue`.

##### Step-by-step examples

Use a tree with root `8`; children `4` and `12`; children `2` and `6`
under `4`; and children `10` and `14` under `12`. Find nodes two links
from target node `4`. The expected output is `[12]`. Walking the tree
first creates `parent={4:8,12:8,2:4,6:4,10:12,14:12}`. The search
then derives `queue=[4]` and `seen={4}` directly from the target.

Initial state: `distance=0`, `queue=[4]`, `seen={4}`,
`result=pending`.

1.  Expand distance `0`
    - Decision: Add unseen children `2`,`6` and parent `8`.
    - After: `distance=1`, `queue=[2,6,8]`, `seen={2,4,6,8}`,
      `result=pending`.
    - Known correct: The queue contains exactly the nodes one link
      away.
2.  Expand distance `1`
    - Decision: Nodes `2`,`6` add nothing; node `8` adds unseen `12`.
    - After: `distance=2`, `queue=[12]`,
      `seen={2,4,6,8,12}`, `result=pending`.
    - Known correct: The queue contains exactly the nodes two links
      away.
3.  Return distance `2`
    - Decision: The requested layer is ready, so return its values.
    - After: `distance=2`, `queue=[12]`,
      `seen={2,4,6,8,12}`, `result=[12]`.
    - Known correct: Every returned node is exactly two links away.

Boundary example: `targetDistance=0` returns the target itself. A nil
target returns an empty result rather than trying to use a nil map key
as a real node.

##### Why it works

After the parent map is built, every child-to-parent entry supplies
the reverse of an existing parent-to-child link. The queue therefore
has all three possible moves. After one complete queue layer, the
queued nodes are exactly one link farther away. Marking on insertion
prevents a parent and child from adding each other again. The tree
has finitely many nodes and each enters once, so the search finishes.

##### Go example

This compact function requires every non-nil `target` to be a node in
the tree rooted at `root`.

```go
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func nodesAtDistance(
	root, target *TreeNode,
	targetDistance int,
) []int {
	// A missing target has no meaningful distance, and negative
	// distances cannot be reached by traversing edges.
	if target == nil || targetDistance < 0 {
		return []int{}
	}
	// Record upward links because ordinary tree nodes expose only
	// downward children, while distance may require moving toward an
	// ancestor.
	parent := map[*TreeNode]*TreeNode{}
	var mapParents func(*TreeNode)
	mapParents = func(node *TreeNode) {
		// Missing children terminate their recursive branch.
		if node == nil {
			return
		}
		if node.Left != nil {
			// Store the reverse of the left-child edge so breadth-first
			// search can later move upward.
			parent[node.Left] = node
		}
		if node.Right != nil {
			// Store the matching reverse edge for the right child.
			parent[node.Right] = node
		}
		// Visit both child subtrees to build parent links for the entire
		// tree.
		mapParents(node.Left)
		mapParents(node.Right)
	}
	mapParents(root)
	// Breadth-first search starts at target. Pointer identity in seen
	// prevents revisiting a node through a different one of its three
	// possible neighbors.
	queue := []*TreeNode{target}
	seen := map[*TreeNode]bool{target: true}
	for distance := 0; distance < targetDistance; distance++ {
		// Process exactly the current layer so one outer iteration equals
		// one edge of distance from target.
		numNodesInCurrentLayer := len(queue)
		if numNodesInCurrentLayer == 0 {
			// No queued node remains, so no node exists at this or any
			// greater distance.
			break
		}
		// Only remove the nodes that were present when this layer began.
		// Newly appended nodes belong to the next distance and wait in the
		// queue.
		numNodesProcessedInCurrentLayer := 0
		for numNodesProcessedInCurrentLayer < numNodesInCurrentLayer {
			// Remove the oldest current-layer node from the queue front.
			node := queue[0]
			queue = queue[1:]
			// Treat left child, right child, and parent as the node's
			// neighbors so the search can move away from target in every
			// direction.
			neighbors := []*TreeNode{
				node.Left,
				node.Right,
				parent[node],
			}
			for _, neighboringNode := range neighbors {
				if neighboringNode != nil && !seen[neighboringNode] {
					// Mark before enqueueing: another neighbor in this layer may
					// also point to this node, but it must enter the next queue
					// layer only once.
					seen[neighboringNode] = true
					queue = append(queue, neighboringNode)
				}
			}
			numNodesProcessedInCurrentLayer++
		}
	}
	// After the requested number of complete layers, the queue contains
	// exactly the nodes at targetDistance.
	values := make([]int, len(queue))
	// Convert the final queue layer to the promised value slice.
	for nodeIndex, node := range queue {
		values[nodeIndex] = node.Value
	}
	return values
}
```

##### Work and memory

Parent mapping and search each visit at most `n` nodes. The search
also stops when its queue becomes empty, so total time is `O(n)`,
even when `targetDistance` is larger than the tree's height. The
parent map, reached set, and queue use `O(n)` auxiliary memory. The
recursive mapping also uses `O(h)` call-stack space within that bound.
The returned layer can contain up to `n` values and therefore uses
`O(n)` output space.

##### When to use it

- A query starts away from the root and may move toward ancestors.
- Tree nodes do not already store parent links.
- Distance is measured in equal-cost links.

##### When not to use it

- Follow existing parent links instead of rebuilding them when they
  are valid.
- A root-to-node path alone may use less memory for a single ancestor
  query.
- Weighted links require a weighted shortest-path method.

##### Common mistakes and edge cases

> **Warning:** If `4` queues parent `8` and nodes are marked only when
> removed, `8` can queue `4` again. The two nodes can keep adding each
> other instead of finishing.

- Map pointers, not values; two distinct nodes may both store `7`.
- Reject or separately define a target outside the root's tree; this
  example requires membership.
- Decide whether a negative `targetDistance` is invalid or produces an
  empty result.

##### Related material

- [Graphs and grids](#b5-graphs-and-grids) explains the temporary
  graph view.
- [Breadth-first search](#c1321-breadth-first-search) explains why
  layer-by-layer search is correct.
- [Queues and deques](#b3-queues-and-deques) explains FIFO queue
  mechanics.

### B.5 Graphs and grids

#### The problem it solves

Given objects and pairwise connections between them, represent which
objects can be reached from which others. Unlike a sequence or tree,
an object may have many neighbors, several routes may reach the same
object, and some groups may be disconnected; direction and edge cost
must also be preserved when the input defines them.

#### How it works

A graph is a map of things and allowed moves. A thing is a **vertex**
or **node**; a move is an **edge**. A grid is often an **implicit
graph**: its neighbors are calculated from coordinates instead of
stored as edge records. In an adjacency list,
`neighbors[vertex] -> every vertex reachable by one stored edge`.
For example, `neighbors[2]=[0,3]` means vertex `2` has one-step moves
to vertices `0` and `3`.

#### Step-by-step examples

Use vertices `0` through `3` and two-way roads `0-1`, `0-2`, and
`2-3`. The adjacency list is:

```text
0: [1, 2]
1: [0]
2: [0, 3]
3: [2]
```

Initial state: `neighbors=[[],[],[],[]]`, `nextRoad=0-1`.

1.  Add road `0-1`
    - Decision: Append each endpoint to the other's list.
    - After: `neighbors=[[1],[0],[],[]]`, `nextRoad=0-2`.
    - Known correct: Both allowed moves for road `0-1` exist.
2.  Add road `0-2`
    - Decision: Append each endpoint to the other's list.
    - After: `neighbors=[[1,2],[0],[0],[]]`, `nextRoad=2-3`.
    - Known correct: Vertex `0` has both of its neighbors.
3.  Add road `2-3`
    - Decision: Append each endpoint to the other's list.
    - After: `neighbors=[[1,2],[0],[0,3],[2]]`, `nextRoad=none`.
    - Known correct: Every road is represented in both directions.

The lists show the allowed route `0 -> 2 -> 3`. Adding isolated vertex
`4` with an empty neighbor list creates another component, a group
with no route to or from the first four vertices.

For vertex order `0,1,2,3`, the same graph's adjacency matrix is:

```text
    0 1 2 3
0   0 1 1 0
1   1 0 0 0
2   1 0 0 1
3   0 0 1 0
```

Reading row `2`, column `3` returns `1`, so edge `2-3` exists in a
fixed amount of time. Row `0`, column `3` returns `0`. The edge-list
form is `[(0,1),(0,2),(2,3)]`.

#### Why it works

Each stored neighbor entry corresponds to one allowed directed move;
adding both entries makes an undirected edge traversable both ways. A
stored path is a sequence in which every next vertex appears in the
current vertex's neighbor list. An empty neighbor list for vertex `4`
therefore creates no path between it and the first component. Grid
neighbor formulas are equivalent when they generate exactly the
allowed in-bounds moves.

#### Go example

The builder requires a nonnegative `vertexCount`. Every endpoint must
be from `0` through `vertexCount-1`.

```go
func undirectedGraph(vertexCount int, edges [][2]int) [][]int {
	// Allocate one adjacency list per vertex, initially with no
	// neighbors.
	neighbors := make([][]int, vertexCount)
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		// An undirected edge can be traversed in either direction, so store
		// both directed adjacency entries.
		neighbors[from] = append(neighbors[from], to)
		neighbors[to] = append(neighbors[to], from)
	}
	return neighbors
}
```

The traced call is
`undirectedGraph(4,[][2]int{{0,1},{0,2},{2,3}})`.

#### Work and memory

Let `V` be the number of vertices and `E` the number of input edges.
An adjacency list takes `O(V + E)` build time and space; scanning
every neighbor has the same time. An adjacency matrix uses `O(V^2)`
space and initialization time but checks one edge in `O(1)`. An edge
list uses `O(E)` space. An undirected adjacency list stores two
neighbor entries per edge, which remains `O(V + E)` space.

#### When to use it

- Objects have many-to-many relationships or allowed transitions.
- A grid move can be modeled as an edge between cells.
- Reachability, connected groups, dependencies, or path costs matter.

#### When not to use it

- A tree is simpler when the input guarantees no cycles and one parent
  route.
- An adjacency matrix stores every possible pair, so it wastes space
  for a **sparse graph**, where only a few pairs have edges.
- Unweighted breadth-first search does not minimize differing edge
  weights.

#### Common mistakes and edge cases

> **Warning:** For an undirected road `0-1`, adding only `0 -> 1`
> makes a search from `1` unable to reach `0`.

- Key clone maps by original node identity, not stored value; two
  nodes can both contain `7`.
- Check row and column bounds separately for grid neighbors.
- Account for graph construction and duplicate edges in total cost.

#### Related material

- [Breadth-first search](#c1321-breadth-first-search) and [depth-first
  search](#c1320-depth-first-search) explain the main visit orders.
- [Bellman-Ford](#c1324-bellman-ford-shortest-paths) naturally
  consumes an edge list.
- [Clone Graph](problems/medium_leetcode_0133_clone_graph.go) requires
  an identity-keyed clone map.
- [Number of
  Islands](problems/medium_leetcode_0200_number_of_islands.go) treats
  a grid as an implicit graph.

### B.6 Matrix traversal and in-place state

#### The problem it solves

Given a grid and rules for visiting or updating its cells, identify
each cell by a row and a column and generate only neighbors that
actually exist. The implementation must handle rectangular or ragged
rows when allowed instead of silently assuming a square grid with
equal dimensions.

#### How it works

A matrix is a set of rows addressed as `matrix[row][column]`. Vertical
movement changes the row; horizontal movement changes the column. A
**rectangular** matrix has equal row lengths. Go also permits **jagged
slices**, whose rows have different lengths.

#### Step-by-step examples

For `[[1,2,3],[4,5,6]]`, query neighbors of `(0,2)`.

Initial state: `candidate=(-1,2)`, `validNeighbors=[]`.

1.  Check `(-1,2)`
    - Decision: Valid row indexes run from `0` through `1`, so
      row `-1` is outside them.
    - After: `candidate=(1,2)`, `validNeighbors=[]`.
    - Known correct: No read occurred outside storage.
2.  Check `(1,2)`
    - Decision: Row `1` and column `2` exist; accept value `6`.
    - After: `candidate=(0,1)`, `validNeighbors=[6]`.
    - Known correct: The down neighbor is valid.
3.  Check `(0,1)`
    - Decision: Row `0` and column `1` exist; accept value `2`.
    - After: `candidate=(0,3)`, `validNeighbors=[6,2]`.
    - Known correct: The left neighbor is valid.
4.  Check `(0,3)`
    - Decision: Column `3` equals row `0`'s length, so reject it.
    - After: `candidate=done`, `validNeighbors=[6,2]`.
    - Known correct: Every accepted position is inside its own row.

Boundary example: in jagged `[[1,2],[]]`, position `(1,0)` has a valid
row but an invalid column. Using the first row's length for all rows
would panic.

#### Why it works

Every valid cell belongs to an existing row and an existing position
in that specific row. Testing in that order proves both conditions
before indexing. Adding the four direction changes creates each
intended adjacent position once; the checks reject only positions
outside storage.

#### Go example

For the traced matrix, calls for candidates `(-1,2)`, `(1,2)`,
`(0,1)`, and `(0,3)` return `false`, `true`, `true`, and `false` in
that order. The caller appends a value to `validNeighbors` only after
`inBounds` returns true.

```go
func inBounds(matrix [][]int, row, column int) bool {
	// Prove the row exists before evaluating len(matrix[row]); &&
	// short-circuits left to right. Then validate the column against
	// that row's actual length.
	return row >= 0 && row < len(matrix) &&
		column >= 0 && column < len(matrix[row])
}
```

Go's left-to-right `&&` evaluation stops after a false row check, so
the code does not evaluate `matrix[row]` for an invalid row.

#### Work and memory

One bounds check takes `O(1)` time and space. Visiting every cell
takes time proportional to the total number of cells, `O(cells)`,
which is `O(rows*cols)` for a rectangle. A returned copy needs
`O(cells)` result memory; in-place visits may use `O(1)` extra
state unless traversal adds a stack or queue.

#### When to use it

- Input is indexed by row and column.
- Neighbor movement or layer boundaries define the work.
- Shape promises can be stated before the algorithm.

#### When not to use it

- Flattened storage may be simpler when row structure has no meaning.
- Do not apply square-only transformations to a rectangle.
- A jagged input cannot use one global column count unless first
  validated.

#### Common mistakes and edge cases

> **Warning:** On a 2-by-3 matrix, checking `column < len(matrix)`
> rejects valid column 2 because the row count is 2.

- Empty `matrix` has no `matrix[0]`.
- An existing row may itself be empty.
- Keep `(row,column)` order consistent; swapping them often survives
  square tests and fails only on rectangles.

#### Related material

- [Graphs and grids](#b5-graphs-and-grids) explains implicit neighbor
  edges.
- [Rows, columns, and shrinking
  boundaries](#b61-rows-columns-and-shrinking-boundaries) applies
  explicit rectangle boundaries.
- [Matrix traversal
  problems](problems/medium_leetcode_0054_spiral_matrix.go) expose
  row-column mistakes.

#### B.6.1 Rows, columns, and shrinking boundaries

##### The problem it solves

Given a rectangular grid, return its values one outer layer at a time
in the required order, such as a clockwise spiral. After each side is
consumed, the remaining top, bottom, left, and right boundaries move;
they may cross and leave only one row or column, which must be emitted
once rather than twice.

##### How it works

Peel the outer frame from a picture, then repeat on the smaller
picture. Four positions named `top`, `bottom`, `left`, and `right`
describe the unread rectangle. The cells at those positions are still
included, so the boundaries are called **inclusive**.

##### Step-by-step examples

For `[[1,2,3],[4,5,6],[7,8,9]]`, expected spiral order is
`[1,2,3,6,9,8,7,4,5]`.

Initial state: `top=0`, `bottom=2`, `left=0`, `right=2`,
`output=[]`.

1.  Emit the top edge
    - Decision: Append row `0` from columns `0` through `2`.
    - After: `top=1`, `bottom=2`, `left=0`, `right=2`,
      `output=[1,2,3]`.
    - Known correct: The old top edge is complete.
2.  Emit the right edge
    - Decision: Append column `2` from rows `1` through `2`.
    - After: `top=1`, `bottom=2`, `left=0`, `right=1`,
      `output=[1,2,3,6,9]`.
    - Known correct: The old right edge is complete.
3.  Emit the bottom edge
    - Decision: Append row `2` from columns `1` down through `0`.
    - After: `top=1`, `bottom=1`, `left=0`, `right=1`,
      `output=[1,2,3,6,9,8,7]`.
    - Known correct: The old bottom edge is complete.
4.  Emit the left edge
    - Decision: Append column `0` from rows `1` through `1`.
    - After: `top=1`, `bottom=1`, `left=1`, `right=1`,
      `output=[1,2,3,6,9,8,7,4]`.
    - Known correct: The outer layer is complete; only center `5`
      remains inside the boundaries.
5.  Emit the inner layer
    - Decision: Append the one-cell top edge at `(1,1)`. The right
      edge is empty, so move `right` inward without appending.
    - After: `top=2`, `bottom=1`, `left=1`, `right=0`,
      `output=[1,2,3,6,9,8,7,4,5]`.
    - Known correct: The boundaries crossed after every cell was
      emitted exactly once.

Boundary example: `[[1,2,3]]` emits its top edge and stops because
`top` then passes `bottom`; it does not emit `3,2,1` again.

##### Why it works

Before each round, everything outside the four boundaries has been
emitted exactly once and everything inside remains unread. Each edge
pass emits one boundary and moves that boundary inward. Crossing
checks prevent overlap. The remaining rectangle shrinks, so all cells
are eventually emitted once.

##### Go example

The function requires every row to have the same length.

```go
func spiralOrder(matrix [][]int) []int {
	// The second condition may read matrix[0], so first prove a row
	// exists.
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	// These inclusive boundaries enclose every cell not yet emitted.
	top, bottom := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1
	order := make([]int, 0, len(matrix)*len(matrix[0]))
	for top <= bottom && left <= right {
		// Consume the top edge from left to right, then remove that row.
		for column := left; column <= right; column++ {
			order = append(order, matrix[top][column])
		}
		top++
		// Consume the right edge from top to bottom, then remove that
		// column.
		for row := top; row <= bottom; row++ {
			order = append(order, matrix[row][right])
		}
		right--
		if top <= bottom {
			// A bottom row remains; traverse it in the opposite direction.
			for column := right; column >= left; column-- {
				order = append(order, matrix[bottom][column])
			}
			bottom--
		}
		if left <= right {
			// A left column remains; traverse upward without repeating
			// corners.
			for row := bottom; row >= top; row-- {
				order = append(order, matrix[row][left])
			}
			left++
		}
	}
	return order
}
```

##### Work and memory

For `r` rows and `c` columns, every cell is appended once, so time is
`O(r*c)`. Boundary variables use `O(1)` extra memory, and the returned
order uses `O(r*c)` result space.

##### When to use it

- Input is a validated rectangle.
- Output follows successive outer layers.
- Four boundaries that move only inward describe unread work.

##### When not to use it

- Validate or adapt the method for jagged rows.
- Use ordinary row-by-row loops, also called row-major order, when
  layer order is irrelevant.
- A square-only rotation has a different destination mapping.

##### Common mistakes and edge cases

> **Warning:** Omitting the post-increment `top <= bottom` check emits
> a one-row matrix once forward and once backward.

- Handle an empty matrix before reading `matrix[0]`.
- Use the column count from the validated rectangular row, not the row
  count.
- Output allocation is not counted as extra working memory, but it
  still consumes memory.

##### Related material

- [Matrix traversal and in-place
  state](#b6-matrix-traversal-and-in-place-state) explains shape and
  coordinate rules.
- [Spiral Matrix](problems/medium_leetcode_0054_spiral_matrix.go) is
  the representative boundary exercise.

#### B.6.2 Layer transformations

##### The problem it solves

Given a square matrix, rotate its values using the matrix's existing
storage. A value cannot be written directly to its destination if
that write would destroy another original value that still needs to
move, so the transformation must preserve displaced data.

##### How it works

First reflect the square across the diagonal running from its top-left
corner to its bottom-right corner, then reverse every row. The first
operation is a **transpose**: a value at a given row and column moves
to that column and row. The row reversal completes a clockwise quarter
turn.

##### Step-by-step examples

Rotate `[[1,2],[3,4]]`; expected output is `[[3,1],[4,2]]`.

Initial state: `matrix=[[1,2],[3,4]]`, `phase=transpose`.

1.  Transpose
    - Decision: Swap cells `(0,1)` and `(1,0)`.
    - After: `matrix=[[1,3],[2,4]]`, `phase=reverseRows`.
    - Known correct: Every value has exchanged row and column.
2.  Reverse rows
    - Decision: Mirror each row around its center.
    - After: `matrix=[[3,1],[4,2]]`, `phase=done`.
    - Known correct: A value from row `r`, column `c` reaches row
      `c`, column `n-1-r`, its clockwise destination.

Boundary example: a 1-by-1 matrix remains unchanged. A 2-by-3
rectangle cannot be rotated in place into the same row layout because
its output shape is 3-by-2.

##### Why it works

Call an original row `r` and column `c`. Transpose moves that value to
row `c`, column `r`. Reversing its new row changes column `r` to
column `n-1-r`, where `n` is the side length. The two operations
together move `(r,c)` to `(c,n-1-r)`, exactly its clockwise
destination. Swapping only when `column > row` handles each pair on
opposite sides of the diagonal once. Row reversal handles every final
position once.

##### Go example

The function requires a square matrix whose rows all have length
`len(matrix)`.

```go
func rotateSquare(matrix [][]int) {
	// Transpose across the main diagonal. Starting above the diagonal
	// ensures each off-diagonal pair is swapped exactly once.
	for row := 0; row < len(matrix); row++ {
		for column := row + 1; column < len(matrix); column++ {
			matrix[row][column], matrix[column][row] =
				matrix[column][row], matrix[row][column]
		}
	}
	// Mirroring every transposed row completes the clockwise coordinate
	// mapping (row,column) -> (column,n-1-row).
	for row := range matrix {
		left := 0
		right := len(matrix[row]) - 1
		for left < right {
			matrix[row][left], matrix[row][right] =
				matrix[row][right], matrix[row][left]
			left++
			right--
		}
	}
}
```

##### Work and memory

For side length `n`, the matrix contains `n*n` cells. Transpose and
row reversal each touch that many cells, so total time is `O(n^2)`.
Loop indexes and temporary swap values use `O(1)` extra memory.

##### When to use it

- The input is square and may be mutated.
- A clockwise rotation must reuse the existing storage.
- Coordinate mapping can be proved before coding.

##### When not to use it

- Allocate a differently shaped result for a non-square rectangle.
- Do not mutate when callers retain the original matrix.
- Use a direct copy when clarity matters more than constant extra
  space.

##### Common mistakes and edge cases

> **Warning:** Transposing every `(row,column)` pair swaps each pair
> twice and restores the original matrix. Visit only one side of the
> diagonal.

- Validate square shape for untrusted input before indexing
  `matrix[column]`.
- Reversing columns after transpose rotates counterclockwise, not
  clockwise.
- Square tests can hide row-column naming mistakes; trace coordinates.

##### Related material

- [Moving values in place](#b11-moving-values-in-place) explains the
  tradeoff of changing the input.
- [Rotate Image](problems/medium_leetcode_0048_rotate_image.go)
  exercises the transpose-and-reverse mapping.

#### B.6.3 Reusing the first row and column as markers

##### The problem it solves

Given a matrix in which an original special cell, such as zero, forces
changes to its entire row and column, produce the final matrix in
place. Changing cells during discovery creates new special values that
look original, while separate marker arrays violate the constant-space
goal; existing boundary cells must safely store the markers instead.

##### How it works

Use the matrix's first row and first column as a notepad. An interior
zero marks its row at column zero and its column at row zero. Two
booleans remember whether the notepad row and column themselves
originally contained zero. After discovery,
`matrix[row][0] -> whether that interior row contained an original
zero`, and `matrix[0][column] -> whether that interior column
contained an original zero`. Thus `matrix[1][0]=0` means row `1`
must become zero; it is a marker, not a newly discovered input zero.

##### Step-by-step examples

Input and expected output:

```text
original           markers recorded       final
[1 2 3]            [1 2 0]                [1 2 0]
[4 5 0]     ->     [0 5 0]         ->     [0 0 0]
[7 8 9]            [7 8 9]                [7 8 0]
```

Initial state: `matrix=[[1,2,3],[4,5,0],[7,8,9]]`,
`firstRowZero=false`, `firstColumnZero=false`, `phase=mark`.

1.  Record the zero at `(1,2)`
    - Decision: Set row marker `(1,0)` and column marker `(0,2)`.
    - After: `matrix=[[1,2,0],[0,5,0],[7,8,9]]`,
      `firstRowZero=false`, `firstColumnZero=false`, `phase=apply`.
    - Known correct: The affected row and column are remembered.
2.  Apply the row marker
    - Decision: Marker `(1,0)` is zero, so set `(1,1)` to zero.
    - After: `matrix=[[1,2,0],[0,0,0],[7,8,9]]`,
      `firstRowZero=false`, `firstColumnZero=false`, `phase=apply`.
    - Known correct: Every interior cell in row `1` is final.
3.  Apply the column marker
    - Decision: Keep `(2,1)=8`; marker `(0,2)` is zero, so
      change `(2,2)` from `9` to `0`.
    - After: `matrix=[[1,2,0],[0,0,0],[7,8,0]]`,
      `firstRowZero=false`, `firstColumnZero=false`, `phase=borders`.
    - Known correct: Every interior cell in row `2` is final.
4.  Finish the borders
    - Decision: Neither saved flag requires another change.
    - After: `matrix=[[1,2,0],[0,0,0],[7,8,0]]`,
      `firstRowZero=false`, `firstColumnZero=false`, `phase=done`.
    - Known correct: Every required cell, and no other cell, is
      zero.

For boundary matrix `[[1,0],[3,4]]`, expect `[[0,0],[3,0]]`.

Initial state: `matrix=[[1,0],[3,4]]`, `firstRowZero=false`,
`firstColumnZero=false`, `phase=saveBorders`.

1.  Save the border facts
    - Decision: The original zero at `(0,1)` sets
      `firstRowZero=true`; the first column has no zero.
    - After: `matrix=[[1,0],[3,4]]`, `firstRowZero=true`,
      `firstColumnZero=false`, `phase=apply`.
    - Known correct: The first-row decision is saved before marker
      cells change roles.
2.  Apply the column marker
    - Decision: Marker `(0,1)` is zero, so set `(1,1)` to zero.
    - After: `matrix=[[1,0],[3,0]]`, `firstRowZero=true`,
      `firstColumnZero=false`, `phase=borders`.
    - Known correct: Column `1` is final without treating its new
      zero as fresh input.
3.  Restore the first row
    - Decision: The saved flag is true, so zero row `0`.
    - After: `matrix=[[0,0],[3,0]]`, `firstRowZero=true`,
      `firstColumnZero=false`, `phase=done`.
    - Known correct: The original border zero affected both its row
      and column exactly once.

An original zero in the first column needs the symmetric saved flag;
cell `(0,0)` cannot encode both border decisions by itself.

##### Why it works

After the first pass, an interior row or column has a zero marker
exactly when an original interior zero requires it. No output changes
occur before every original is inspected. The second pass follows
those complete markers. Saved booleans restore the two border
decisions that share `(0,0)`.

##### Go example

The function requires a rectangular matrix.

```go
func setZeroes(matrix [][]int) {
	// Guard matrix[0] and treat a matrix with no cells as already
	// complete.
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}
	// The first row and column will store markers, so save their
	// original zero status separately before changing either border.
	rows, columns := len(matrix), len(matrix[0])
	firstRowZero, firstColumnZero := false, false
	// Capture whether the marker row itself must eventually be zeroed.
	for column := 0; column < columns; column++ {
		firstRowZero = firstRowZero || matrix[0][column] == 0
	}
	// Capture the corresponding fact for the marker column.
	for row := 0; row < rows; row++ {
		firstColumnZero = firstColumnZero || matrix[row][0] == 0
	}
	// Record which interior rows and columns contain an original zero.
	for row := 1; row < rows; row++ {
		for column := 1; column < columns; column++ {
			if matrix[row][column] == 0 {
				// One original zero marks both its row and its column on the
				// borders.
				matrix[row][0], matrix[0][column] = 0, 0
			}
		}
	}
	// Apply the completed markers only after all original interior cells
	// have been inspected, preventing generated zeros from becoming new
	// evidence.
	for row := 1; row < rows; row++ {
		for column := 1; column < columns; column++ {
			if matrix[row][0] == 0 || matrix[0][column] == 0 {
				// Either marker proves this cell shares a row or column with an
				// original zero.
				matrix[row][column] = 0
			}
		}
	}
	// Restore the two border decisions from their saved booleans;
	// matrix[0][0] could not represent both facts by itself.
	if firstRowZero {
		// Apply the saved row decision only after its marker values are no
		// longer needed.
		for column := 0; column < columns; column++ {
			matrix[0][column] = 0
		}
	}
	if firstColumnZero {
		// Finish with the saved column decision for the same reason.
		for row := 0; row < rows; row++ {
			matrix[row][0] = 0
		}
	}
}
```

##### Work and memory

For `r` rows and `c` columns, the passes inspect all `r*c` cells a
constant number of times, so time is `O(r*c)`. Two booleans and loop
positions use `O(1)` extra memory; the matrix stores its own markers.

##### When to use it

- The input is a mutable nonempty rectangle.
- Marker values can safely occupy reserved existing cells.
- Extra memory must stay constant.

##### When not to use it

- Separate marker slices are clearer when `O(rows+columns)` memory is
  allowed.
- Do not overwrite border cells before saving their original meaning.
- Jagged rows need validation or a different representation.

##### Common mistakes and edge cases

> **Warning:** Zeroing row 1 immediately after seeing `(1,2)` creates
> a zero at `(1,0)` that a later naive scan may mistake for an
> original zero in column 0.

- Handle an empty matrix before `matrix[0]`.
- One flag cannot represent both the first row and first column.
- Apply interior cells before erasing marker borders.

##### Related material

- [Matrix traversal and in-place
  state](#b6-matrix-traversal-and-in-place-state) explains
  rectangular-shape checks.
- [Set Matrix
  Zeroes](problems/medium_leetcode_0073_set_matrix_zeroes.go) is the
  representative marker exercise.

#### B.6.4 Keeping current and next state together

##### The problem it solves

Given a grid update whose cells all change at the same logical moment,
compute every next value from the same original grid. Updating cells
one by one is incorrect because later cells would read a mixture of
old and new values; each cell must temporarily retain both states or
the next state must be stored elsewhere.

##### How it works

Write the old and new status on the same sticky note. If inputs use
`0` and `1`, reserve `2` for "was 0, becomes 1" and `3` for "was 1,
becomes 0." Reads recover the old status; a final pass keeps only the
new status. Storing both meanings temporarily in each input cell is an
**in-place state encoding**.

##### Step-by-step examples

Suppose a one-row rule says a cell becomes live when exactly one old
neighbor is live. Input `[1,1,1,0]` should simultaneously become
`[1,0,1,1]`.

Initial state: `encoded=[1,1,1,0]`, `cell=0`, `phase=encode`.

1.  Encode cell `0`
    - Decision: It has one old live neighbor, so it stays
      live.
    - After: `encoded=[1,1,1,0]`, `cell=1`, `phase=encode`.
    - Known correct: Cell `0`'s next status is recorded.
2.  Encode cell `1`
    - Decision: It has two old live neighbors, so encode `3` for
      old live and new dead.
    - After: `encoded=[1,3,1,0]`, `cell=2`, `phase=encode`.
    - Known correct: Cell `1`'s old live state remains decodable.
3.  Encode cell `2`
    - Decision: Decode left neighbor `3` as old live and right
      neighbor `0` as old dead; exactly one was live, so stay live.
    - After: `encoded=[1,3,1,0]`, `cell=3`, `phase=encode`.
    - Known correct: Cell `2` used one consistent old generation,
      not cell `1`'s new dead state.
4.  Encode cell `3`
    - Decision: Its old live left neighbor makes it become live, so
      encode `2` for old dead and new live.
    - After: `encoded=[1,3,1,2]`, `cell=done`, `phase=finalize`.
    - Known correct: Both kinds of state change are recorded without
      losing an old value.
5.  Finalize
    - Decision: Replace every temporary code with its new value.
    - After: `encoded=[1,0,1,1]`, `cell=done`, `phase=done`.
    - Known correct: Every cell now shows the result for the same
      next round.

Boundary example: an empty grid performs neither pass. If input may
already contain `2` or `3`, these temporary codes are ambiguous and
cannot be used.

##### Why it works

During the first pass, every stored value can still be decoded to its
original state, even after its next state is added. Therefore every
decision observes one consistent old round. Each cell receives one
next-state encoding. The final pass is safe because no further
old-state decisions remain.

##### Go example

The function requires every input value to be `0` or `1`.
The trace's `encoded` slice is the code's `values`, and its `cell`
is `cellIndex`. Starting with `values := []int{1, 1, 1, 0}` and
calling `updateRow(values)` mutates it to `[]int{1, 0, 1, 1}`.

```go
func updateRow(values []int) {
	// Values 1 and 3 were live in the old generation; 3 additionally
	// records that the cell will be dead in the new generation.
	wasAlive := func(value int) bool {
		return value == 1 || value == 3
	}
	// First pass: decide every next state while preserving a decodable
	// old state.
	for cellIndex, value := range values {
		numLiveNeighbors := 0
		// Count the left neighbor only when its index exists, decoding
		// temporary state 3 as formerly live.
		if cellIndex > 0 && wasAlive(values[cellIndex-1]) {
			numLiveNeighbors++
		}
		// Apply the identical check to the right neighbor without passing
		// the end.
		if cellIndex+1 < len(values) && wasAlive(values[cellIndex+1]) {
			numLiveNeighbors++
		}
		willBeAlive := numLiveNeighbors == 1
		if value == 0 && willBeAlive {
			// 2 means old dead, new live.
			values[cellIndex] = 2
		} else if value == 1 && !willBeAlive {
			// 3 means old live, new dead.
			values[cellIndex] = 3
		}
	}
	// Second pass: every decision is complete, so discard the old-state
	// encoding.
	for cellIndex, value := range values {
		if value == 2 {
			// A dead-to-live transition becomes its final live value.
			values[cellIndex] = 1
		} else if value == 3 {
			// A live-to-dead transition becomes its final dead value.
			values[cellIndex] = 0
		}
	}
}
```

##### Work and memory

With a fixed number of neighbors, each of `n` cells takes constant
work, so time is `O(n)`; for a matrix, `n = rows*columns`. Temporary
codes and counters use `O(1)` extra memory. A clearer separate
next-state grid would use `O(n)` additional space.

##### When to use it

- Updates are logically simultaneous.
- The input value range leaves unambiguous temporary codes.
- Changing the input is allowed and extra memory must stay constant.

##### When not to use it

- Use a separate next-state grid when clarity or unrestricted input
  values matter more than memory.
- Do not encode before documenting how every temporary value decodes.
- If several parts of a program can write the grid at the same time,
  they need a separate coordination rule; this example uses one
  execution thread.

##### Common mistakes and edge cases

> **Warning:** Treating temporary `3` as dead during the decision pass
> makes a later cell read the new state instead of the required old
> live state.

- Finalize only after every decision is recorded.
- Ensure temporary values cannot appear in valid input.
- Empty and single-cell inputs still need well-defined neighbor rules.

##### Related material

- [Matrix traversal and in-place
  state](#b6-matrix-traversal-and-in-place-state) explains coordinate
  and shape checks.
- [Moving values in place](#b11-moving-values-in-place) discusses when
  changing the input is allowed.
- [Game of Life](problems/medium_leetcode_0289_game_of_life.go)
  applies this encoding to eight neighbors.

### B.7 Composing data structures

#### The problem it solves

Given operations that need incompatible strengths, such as finding
and removing a value quickly while also selecting by numeric position,
store the same logical items in two cooperating data structures. Every
insertion, update, and removal must repair both views before the
operation finishes.

#### How it works

Store the values densely in a slice so a numeric position is direct.
Also store `value -> slice position` in a map so a value is direct.
**Composing data structures** means combining such views and defining
the rule that keeps them in agreement.

#### Step-by-step examples

A structure for fast removal by value and direct selection by position
stores
`values=[10,20,30]` and
`indexByValue={10:0,20:1,30:2}`. Delete `20`; expected retained values
are `[10,30]` with map `{10:0,30:1}`. This set requires its stored
values to be unique.

Initial state: `values=[10,20,30]`,
`indexByValue={10:0,20:1,30:2}`, `removeIndex=unknown`.

1.  Locate `20`
    - Decision: Read its stored position from `indexByValue`.
    - After: `values=[10,20,30]`,
      `indexByValue={10:0,20:1,30:2}`, `removeIndex=1`.
    - Known correct: No scan is needed.
2.  Move the final value
    - Decision: Copy `30` to index `1` and repair its map entry.
    - After: `values=[10,30,30]`,
      `indexByValue={10:0,20:1,30:1}`, `removeIndex=1`.
    - Known correct: Retained positions `0` and `1` agree; the
      spare final slot and removed key still need deletion.
3.  Shrink
    - Decision: Drop the duplicate final slice slot and delete map
      key `20`.
    - After: `values=[10,30]`, `indexByValue={10:0,30:1}`,
      `removeIndex=1`.
    - Known correct: Both views agree completely.

Boundary example: deleting final value `30` from `[10,20,30]` moves it
onto itself, then shrinks and deletes its key. The same steps remain
valid.

#### Why it works

Between operations, every slice value's map entry must name that
value's current slice position. In symbols,
`indexByValue[values[i]] == i` for every retained position `i`.
Lookup identifies the only hole. Moving the last value fills it
without shifting other values, and repairing that one map entry
restores agreement for the moved value. Shrinking removes the
duplicate slot, and deleting the removed key restores the complete
rule. Each operation changes only a fixed number of positions.

#### Go example

`initialValues` must contain no duplicates.

```go
// indexedSet keeps the invariant indexByValue[values[i]] == i.
type indexedSet struct {
	values       []int
	indexByValue map[int]int
}

func newIndexedSet(initialValues []int) indexedSet {
	// The input promise requires unique values. Copy them so later
	// changes do not modify the caller's slice.
	set := indexedSet{
		values:       append([]int(nil), initialValues...),
		indexByValue: make(map[int]int, len(initialValues)),
	}
	for valueIndex, value := range set.values {
		set.indexByValue[value] = valueIndex
	}
	return set
}

func (set *indexedSet) remove(value int) bool {
	// The map finds the removal slot without scanning the slice.
	removeIndex, exists := set.indexByValue[value]
	if !exists {
		// Neither representation changes when the requested value is
		// absent.
		return false
	}
	// Move the final value into the hole and repair only its map entry.
	// This also works when removing the final value because it is moved
	// onto itself.
	lastIndex := len(set.values) - 1
	lastValue := set.values[lastIndex]
	set.values[removeIndex] = lastValue
	set.indexByValue[lastValue] = removeIndex
	// Shrink away the duplicate final slot, then remove the requested
	// key.
	set.values = set.values[:lastIndex]
	delete(set.indexByValue, value)
	return true
}

func (set *indexedSet) valueAt(position int) (int, bool) {
	if position < 0 || position >= len(set.values) {
		return 0, false
	}
	// Dense slice storage provides direct selection by valid position.
	return set.values[position], true
}
```

The traced operation is
`set := newIndexedSet([]int{10,20,30}); set.remove(20)`. It returns
`true` and leaves the two internal views exactly as shown above.
Calling `set.valueAt(1)` then returns `(30,true)`.

#### Work and memory

Map lookup, one slice replacement, and map updates take expected
`O(1)` time; "expected" relies on ordinary hash-map behavior and is
not a worst-case per-operation guarantee. The two views store `O(n)`
values. Selecting a valid slice position also takes `O(1)` time.

#### When to use it

- Required operations have conflicting lookup needs.
- A clear rule can relate the structures after every operation.
- The memory for multiple views is justified by repeated faster
  operations.

#### When not to use it

- One structure already meets every required bound.
- Updates cannot repair every view before callers can observe the
  operation.
- Duplicates are required but the value-to-one-index map does not
  represent them; use sets of indices or another design.

#### Common mistakes and edge cases

> **Warning:** Moving `30` into index 1 without changing map entry
> `30 -> 2` makes the next removal index outside the shortened slice.

- Initialize maps before assignment.

#### Related material

- [Map presence and
  sets](#d4-map-presence-and-sets)
  explains Go map lookup behavior.
- [Heap or priority queue](#c1327-heap-or-priority-queue) explains
  heap operations.
- [LRU Cache](problems/medium_leetcode_0146_lru_cache.go) combines a
  map and doubly linked list.
- [Insert Delete GetRandom
  O(1)](problems/medium_leetcode_0380_insert_delete_getrandom_o1.go)
  uses the traced slice-map invariant.

### B.8 Streaming and online processing

#### The problem it solves

Given a sequence whose values arrive one at a time, return the
required answer after each arrival without assuming future values are
already available. The algorithm must keep a bounded summary of the
useful past when storing or rescanning the entire history would exceed
the time or memory limit.

#### How it works

Keep a small scoreboard and update it when each value arrives. An
**online algorithm** maintains a valid answer without assuming future
input. A **prefix** is everything received so far. A **batch
algorithm** waits for the complete input before beginning. Here,
`counts[value] -> number of received copies of value`, while the queue
stores first occurrences in arrival order. For example,
`counts['a']=2` means `a` is no longer unique.

#### Step-by-step examples

For stream `a, b, a, a, b`, report the first non-repeating letter
after each arrival; expected reports are `a, a, b, b, no value`.

Initial state: `counts={}`, `queue=[]`, `reports=[]`.

1.  Receive `a`
    - Decision: Increment its count and enqueue its first occurrence.
    - After: `counts={a:1}`, `queue=[a]`, `reports=[a]`.
    - Known correct: The front is the earliest unique item.
2.  Receive `b`
    - Decision: Increment its count and enqueue its first occurrence.
    - After: `counts={a:1,b:1}`, `queue=[a,b]`, `reports=[a,a]`.
    - Known correct: Older unique `a` remains first.
3.  Receive `a`
    - Decision: Its count becomes `2`; discard the repeated front.
    - After: `counts={a:2,b:1}`, `queue=[b]`,
      `reports=[a,a,b]`.
    - Known correct: The front is again the earliest unique item.
4.  Receive `a`
    - Decision: Its count becomes `3`; do not enqueue it again.
    - After: `counts={a:3,b:1}`, `queue=[b]`,
      `reports=[a,a,b,b]`.
    - Known correct: Removed `a` stays out; unique `b` stays first.
5.  Receive `b`
    - Decision: Its count becomes `2`; discard the repeated front.
    - After: `counts={a:3,b:2}`, `queue=[]`,
      `reports=[a,a,b,b,none]`.
    - Known correct: No received letter occurs exactly once.

Boundary example: stream `a, a` reports `a`, then no value. The queue
becomes empty after the second arrival; code must represent "none"
separately from a valid zero-value item.

#### Why it works

Counts describe every received item. Queue order records first-arrival
order, and only first occurrences enter it. Removing front items while
their count is not one discards candidates proved invalid. The first
remaining item is therefore the earliest received value with count
one. Each arrival extends the prefix, and every queued item can be
removed only once.

#### Go example

```go
func firstUniqueReports(stream []byte) []string {
	// counts determines uniqueness; queue preserves first-arrival order
	// among values that have appeared at least once.
	counts := map[byte]int{}
	queue := []byte{}
	reports := make([]string, 0, len(stream))
	for _, value := range stream {
		counts[value]++
		if counts[value] == 1 {
			// Enqueue only the first occurrence, so each value enters once.
			queue = append(queue, value)
		}
		// Repeated values can become invalid but never unique again, so
		// stale candidates are permanently removed from the front.
		for len(queue) > 0 && counts[queue[0]] != 1 {
			queue = queue[1:]
		}
		// Report the earliest still-unique byte as a one-byte string. An
		// empty string means no unique byte remains.
		if len(queue) == 0 {
			reports = append(reports, "")
		} else {
			reports = append(reports, string(queue[0]))
		}
	}
	return reports
}
```

For input `[]byte{'a','b','a','a','b'}`, the returned strings are
`["a","a","b","b",""]`. The empty string means "no value" and
cannot be confused with the one-byte string made from any input byte.

#### Work and memory

Each of `n` arrivals is added at most once and removed at most once,
so total time is `O(n)`. Even if one arrival removes several items,
the cost averages `O(1)` per arrival across the full stream. This is
the amortized cost. The count map and queue each use `O(u)` memory for
up to `u` distinct candidates, bounded by `O(n)`. The report slice
uses `O(n)` result memory. A bounded recent window can instead cap
retained state by its window size.

#### When to use it

- An answer is required after every arrival or input cannot be
  replayed.
- The remembered data has a precise meaning after every received
  prefix.
- Old items can be summarized or safely expired.

#### When not to use it

- A batch sort may be simpler when the full input is available and
  intermediate answers are unnecessary.
- Exact answers may be impossible with bounded memory if old data can
  affect every future result.
- Define lateness rules before processing events that can arrive out
  of order.

#### Common mistakes and edge cases

> **Warning:** Returning byte `0` to mean "no unique item" is
> ambiguous when the stream can contain the zero byte. Use a separate
> boolean or a result type with an unambiguous "none" value.

- One arrival may remove many candidates, but the average-over-time
  analysis counts each candidate removal only once.
- Advancing a slice queue can retain old backing storage in a
  long-lived stream.
- A fixed recent window must undo the contribution of the item that
  expires.

#### Related material

- [Queues and deques](#b3-queues-and-deques) explains FIFO and
  retained-storage details.
- [Composing data structures](#b7-composing-data-structures) explains
  paired state such as counts plus order or two heaps.
- [Heap or priority queue](#c1327-heap-or-priority-queue) supports
  running median, top-`k`, and merging sorted streams.
- [Find Median from Data
  Stream](problems/hard_leetcode_0295_find_median_from_data_stream.go)
  maintains two balanced halves.

### B.9 Choosing a structure

#### The problem it solves

Given the operations a problem performs and how often it performs
each one, choose storage that makes the dominant operations affordable
while preserving the required order, membership, or connectivity
rule. Choosing a familiar structure by name alone can make the most
frequent operation too slow or the structure's invariant impossible
to maintain.

#### How it works

Choose a toolbox drawer by the action you repeat most, not by the
shape of the input alone. A **structure-selection guide** compares
observable operation and ordering needs, then links to the entry that
explains each structure's behavior.

#### Step-by-step examples

Suppose numbers arrive one at a time. Keep only the largest three seen
so far. After `8, 2, 5, 9, 1`, the retained values should be
`{5, 8, 9}`.

Initial state: `k=3`, `structure=unselected`, `retained=unset`,
`nextValues=[8,2,5,9,1]`.

1.  Choose what to store
    - Decision: Only three values are needed, and the smallest of
      those three must be easy to replace.
    - After: `k=3`, `structure=minHeap`, `retained={}`,
      `nextValues=[8,2,5,9,1]`.
    - Known correct: The structure supports the two needed jobs:
      find the smallest retained value and replace it.
2.  Process `8,2,5`
    - Decision: The heap has room, so keep all three values.
    - After: `k=3`, `structure=minHeap`, `retained={2,5,8}`,
      `nextValues=[9,1]`.
    - Known correct: Fewer than four values have arrived, so none
      should be discarded.
3.  Process `9`
    - Decision: `9` exceeds heap minimum `2`, so replace `2`.
    - After: `k=3`, `structure=minHeap`, `retained={5,8,9}`,
      `nextValues=[1]`.
    - Known correct: These are the largest three values seen.
4.  Process `1`
    - Decision: `1` is smaller than heap minimum `5`, so
      discard it.
    - After: `k=3`, `structure=minHeap`, `retained={5,8,9}`,
      `nextValues=[]`.
    - Known correct: `1` cannot be one of the largest three.

Here, `k` means the number of values to keep. When `k=0`, the answer
is empty, so do not read the heap's top. If the program must return
every value in sorted order, a heap that keeps only `k` values has
discarded information the answer needs. Sorting all values is a
better fit.

#### Why it works

After each arrival, the heap contains the largest three values seen so
far, or every value if fewer than three have arrived. Once the heap is
full, its top is the smallest retained value. A new value no larger
than the top cannot enter the largest three. A larger value replaces
that one smaller value. This keeps the same fact true after every
arrival and gives the required answer when the input ends.

#### Go example

```go
import (
	"container/heap"
	"sort"
)

type smallestHeap []int

func (values smallestHeap) Len() int { return len(values) }
func (values smallestHeap) Less(left, right int) bool {
	return values[left] < values[right]
}
func (values smallestHeap) Swap(left, right int) {
	values[left], values[right] = values[right], values[left]
}
func (values *smallestHeap) Push(value any) {
	*values = append(*values, value.(int))
}
func (values *smallestHeap) Pop() any {
	oldValues := *values
	lastIndex := len(oldValues) - 1
	value := oldValues[lastIndex]
	*values = oldValues[:lastIndex]
	return value
}

func largestValues(values []int, numToKeep int) []int {
	if numToKeep <= 0 {
		return []int{}
	}
	kept := &smallestHeap{}
	heap.Init(kept)
	for _, value := range values {
		if kept.Len() < numToKeep {
			heap.Push(kept, value)
			continue
		}
		if value <= (*kept)[0] {
			continue
		}
		// Replace the smallest retained value, then restore heap order.
		(*kept)[0] = value
		heap.Fix(kept, 0)
	}
	// Heap storage is not fully sorted, so sort a copy for the caller.
	result := append([]int(nil), (*kept)...)
	sort.Ints(result)
	return result
}
```

The trace's `k` is the code's `numToKeep`, and its `retained` values
are stored in the `kept` min-heap. Calling
`largestValues([]int{8,2,5,9,1},3)` returns `[5,8,9]`. The function
returns an empty slice when `numToKeep` is zero or negative.

#### Work and memory

Choosing a structure is a planning step. The running cost comes from
the structure ultimately used. A map usually finds a key in constant
time, written expected `O(1)`, and a slice reads a known position in
`O(1)`. Searching a linked list of `n` nodes may inspect all of them,
so it takes `O(n)`. Adding or removing at the intended end of a stack,
queue, or deque usually takes `O(1)`. Adding to or removing from a
balanced tree or heap takes `O(log n)`.

For the example, let `n` be the number of arriving values and `k` be
the number to keep. Updating a heap of at most `k` values takes up to
`O(log k)` time per arrival. The full process therefore takes
`O(n log k)` time. Sorting the returned copy adds `O(k log k)` time.
The heap and returned copy each store at most `k` values, so working
memory and result memory are both `O(k)`.

#### When to use it

- Before writing code when several structures seem possible.
- When the largest allowed input makes speed or memory important.
- When no single structure makes every required job fast; two
  structures may need to work together.

#### When not to use it

- Do not choose from one matching word alone. Check that the structure
  supports every operation the problem requires.
- The input's appearance does not decide the answer: tree-shaped input
  does not automatically require a binary search tree, and an array
  does not automatically require two pointers.
- Reconsider the choice when values may repeat, the input cannot be
  changed, order matters, or the output needs information the
  structure does not keep.

#### Common mistakes and edge cases

> **Warning:** A heap quickly provides only the smallest, largest, or
> highest-priority value it was built to expose. Finding some other
> exact value may require checking all `n` stored values. If both jobs
> must be fast, pair the heap with a map and keep them in agreement.

- Count setup work, such as sorting, and the work needed to build the
  output when comparing total cost.
- Go's standard library supports heaps, but it does not provide a map
  that automatically keeps its keys sorted.
- Decide what repeated values mean. A set removes repeats, while a
  heap or search tree may keep them.
- Check for an empty structure before removing a value. For top-`k`
  problems, also handle `k=0` before reading the heap's top.

#### Related material

- [Problem-solving method selection
  guide](#c1-problem-solving-method-selection-guide) helps choose the
  algorithm after choosing how to store the data.
- [Composing data structures](#b7-composing-data-structures) explains
  how two structures can support jobs that one cannot handle alone.
- [Hash map or set](#c134-hash-map-or-set), [stack](#c1313-stack),
  [heap](#c1327-heap-or-priority-queue), [trie](#c1333-trie), and
  [disjoint-set union](#c1325-disjoint-set-union) explain when each
  structure fits a problem.

<a id="c-problem-solving-patterns"></a>

## C. Ways to solve problems

This section turns what you notice about an input into a solution plan
you can explain and trust. A method's name is only a label. What
matters is why that method may safely ignore work that cannot affect
the answer.

### C.1 Problem-solving method selection guide

Before choosing a method, answer five questions in plain language:

1.  What shape does the input have: sequence, range, graph, tree,
    stream, or choices?
2.  What must the output contain: a value, position, path, order, one
    answer, or every answer?
3.  Must the original order, neighboring relationships, or positions
    stay unchanged?
4.  May the input be changed, and does all of it exist before
    processing starts?
5.  What time and memory can the largest valid input afford?

Treat the links below as possible starting points. Before writing
code, confirm that the method's requirements match the problem and
that its explanation of correctness makes sense for this input.
You do not need to memorize every method name. Start with a concrete
question, follow one promising link, and reject the method if its
required conditions are absent.

<a id="c11-quick-decision-tree"></a>

#### C.1.1 Ask what the input and output allow (quick decision tree)

##### The problem it solves

Given a new problem statement, its input limits, and its required
output, narrow the possible solution methods before writing code. The
decision must come from observable facts such as ordering, repeated
queries, allowed moves, and edge costs; a word such as "minimum" does
not by itself prove that one particular algorithm applies.

##### How it works

Choosing an algorithm is like choosing transport for a delivery: ask
about the load, destination, allowed changes, and deadline before
picking a vehicle. A **decision tree** is that ordered list of
questions; each answer rules out methods that do not fit.

##### Step-by-step examples

For a problem allowing up to `100000` values, consider sorted sample
`[1,2,4,7,11]`, target `9`, and a request for any pair of values. The
expected sample output may be `[2,7]`.

Initial state: `sorted=true`, `maxN=100000`, `output=values`,
`method=unknown`.

1.  Check the output contract
    - Decision: Only values are required, so reordering is allowed.
    - After: `sorted=true`, `maxN=100000`, `output=values`,
      `reorderingAllowed=true`, `method=unknown`.
    - Known correct: Original positions are not part of the
      requested answer.
2.  Check the input shape
    - Decision: A small sum can increase by moving the left position;
      a large sum can decrease by moving the right position.
    - After: `sorted=true`, `maxN=100000`, `output=values`,
      `reorderingAllowed=true`, `method=twoPointersCandidate`.
    - Known correct: Sorted order can justify discarding pairs
      after each move.
3.  Check the limits
    - Decision: About five billion pair checks are too many;
      inward-only pointers take at most one pass.
    - After: `sorted=true`, `maxN=100000`, `output=values`,
      `reorderingAllowed=true`, `method=twoPointers`.
    - Known correct: The selected method can fit the input size.

For unsorted `[2, 7, 11, 15]` when original indices are required,
sorting would hide the requested positions unless pairs are copied.
The lookup branch instead suggests a [hash
map](#c134-hash-map-or-set), which preserves scan positions.

##### Why it works

Naming an algorithm does not prove that it works. Each question keeps
only methods that still match the input, required output, permission
to change data, and available time and memory. The chosen card must
then explain why that specific method gives the right answer. If none
of the cards fit, use the simple direct method instead.

##### Go example

Selection produces a plan rather than reusable code. For the example
above, see [E.9 Sorted two-pointer pair
search](#e9-sorted-two-pointer-pair-search). On the traced input, it
returns indexes `(1,3,true)`; reading those positions produces the
requested values `[2,7]`. A version needing original positions from
unsorted input would use the map behavior explained in
[D.4 Map presence and sets](#d4-map-presence-and-sets).

##### Work and memory

Answering these few questions is a planning step, not part of the
algorithm's running time. When comparing methods, count setup work
such as sorting, extra storage such as maps or queues, unfinished
function calls from recursion, and the work needed to build the
output.

##### When to use it

Use the guide before implementation when two or more methods seem
plausible or when a familiar keyword is the only reason for a choice.

##### When not to use it

Do not force a listed method when its required ordering, cost, or
state rule is absent. A direct solution or a method outside this guide
may be the right fit.

##### Common mistakes and edge cases

> **Warning:** "Minimum" does not prove greedy. Coin values
> `[1, 3, 4]` and target `6` defeat "always take the largest coin":
> `4+1+1` loses to `3+3`.

An empty input, values that arrive over time, an input that cannot be
changed, repeated or negative values, and an extremely large output
can each change which method is practical.

##### Related material

- [C.1.2 Derive a direction before naming a
  pattern](#c12-derive-a-direction-before-naming-a-pattern) starts
  from a direct solution when the first questions do not select a
  clear method.
- [C.1.4 Choosing among overlapping
  methods](#c14-choosing-among-overlapping-methods) compares methods
  that survive the first questions.
- [A.1 Complexity and resource
  accounting](#a1-complexity-and-resource-accounting) explains how to
  reject a method from the constraints.

#### C.1.2 Derive a direction before naming a pattern

##### The problem it solves

Given a problem that does not clearly match a known pattern, derive a
correct direction instead of guessing algorithm names. First describe
the simplest process that must find the answer, then identify exactly
which comparisons, states, or subproblems it repeats and what fact
would allow that repeated work to be saved or skipped.

##### How it works

Solve a tiny input by hand. Describe the simplest correct algorithm,
then mark calculations it repeats or information it throws away. Those
details point toward an improvement before you need to recall a formal
method name. For the example below,
`prefixSum[endExclusive] -> sum of values before endExclusive`.
An inclusive right boundary becomes `endExclusive=right+1`; for
`right=3`, lookup position `4` includes indexes `0` through `3`.

##### Step-by-step examples

For `[2,4,1,3]`, suppose many queries ask for the sum from index
`left` through index `right`, including both ends. One query uses
`left=1` and `right=3`, so it asks for `4+1+3` and should return `8`.

Initial state: `method=direct`, `savedTotals=[]`, `answers=[]`.

1.  Answer `[1,3]` directly
    - Decision: Add values at indexes `1`, `2`, and `3`:
      `4+1+3`.
    - After: `method=direct`, `savedTotals=[]`, `answers=[8]`.
    - Known correct: Every value in the requested range was
      counted once.
2.  Answer `[0,2]` directly
    - Decision: Add `2+4+1`; this repeats additions from the first
      query already used.
    - After: `method=direct`, `savedTotals=[]`, `answers=[8,7]`.
    - Known correct: The direct method works, but every
      query rebuilds totals from individual values.
3.  Save totals from the beginning
    - Decision: Store `[0,2,6,7,10]`; entry `i` sums values before
      index `i`.
    - After: `method=prefixSum`, `savedTotals=[0,2,6,7,10]`,
      `answers=[8,7]`.
    - Known correct: `savedTotals[4]-savedTotals[1]=8` removes values
      before index `1` and leaves indexes `1` through `3`.

For one query only, the setup may cost as much as the direct scan.
The observation then does not justify extra storage.

##### Why it works

The direct method is the correctness baseline: it explicitly performs
all work needed to produce the answer. A shortcut is valid only when
its saved fact accounts for that same work or proves that some work
cannot affect the result. In the example, subtracting two saved totals
includes every requested value once and excludes every other value.
If no such argument is clear, keep the direct solution.

##### Go example

The [prefix sums](#c1315-prefix-sums) entry explains the saved totals
and their Go index rules. This planning process itself does not add
runtime code.

##### Work and memory

Tracing a small example is planning work. The final cost must include
any setup added to avoid repetition. In this example, building the
saved totals takes work and memory proportional to the input size;
each later range query then takes constant time.

##### When to use it

Use it when pattern recall stalls, when a proposed optimization has no
proof, or when the direct solution is clear but too slow.

##### When not to use it

Do not optimize a direct method that already meets the limits unless a
simpler representation improves correctness. Extra machinery can add
failure modes without useful savings.

##### Common mistakes and edge cases

> **Warning:** Repeated syntax is not necessarily repeated work. Two
> loops may process disjoint items, while one nested loop may
> repeatedly rebuild the same range sum.

Check zero queries, one query, empty input, the unavoidable size of
the output, and whether saved information becomes invalid after
updates.

##### Related material

- [C.1.1 Ask what the input and output
  allow](#c11-ask-what-the-input-and-output-allow-quick-decision-tree)
  narrows the resulting candidates.
- [C.1.5 Let constraints eliminate bad
  choices](#c15-let-constraints-eliminate-bad-choices) checks whether
  the direct method is already affordable.
- [A.2 Correctness and invariants](#a2-correctness-and-invariants)
  explains how to prove that removing repeated work preserves
  coverage.

#### C.1.3 Pattern cards

Each card explains one method: what to picture, when to use it, how to
start, why it works, and when to choose something else. The cards
describe reusable methods rather than answers to individual problems,
so they can guide a first attempt without giving away a judge
solution.

Jump directly to a card:

- Sequence and range methods: [common-prefix
  scan](#c131-common-prefix-scan), [directional
  scan](#c132-directional-scan), [simulation](#c133-simulation), [two
  pointers](#c139-two-pointers), [sorting and
  selection](#c135-sorting-and-selection), [binary
  search](#c136-binary-search), [string
  matching](#c137-string-matching), [parsing and state
  machines](#c138-parsing-and-state-machines), [sliding
  window](#c1314-sliding-window), [prefix sums](#c1315-prefix-sums),
  [difference arrays](#c1316-difference-arrays), [Fenwick and segment
  trees](#c1317-fenwick-and-segment-trees), [fast and slow
  pointers](#c1332-fast-and-slow-pointers), [Kadane's
  algorithm](#c1312-kadanes-algorithm-best-subarray-scan), [monotonic
  stack](#c1328-monotonic-stack), [monotonic
  deque](#c1329-monotonic-deque), [intervals](#c1330-intervals), and
  [sweep line](#c1331-sweep-line).
- Lookup and processing-order methods: [hash map or
  set](#c134-hash-map-or-set), [trie](#c1333-trie),
  [stack](#c1313-stack), and [heap or priority
  queue](#c1327-heap-or-priority-queue).
- Choice methods: [dynamic programming](#c1310-dynamic-programming),
  [greedy](#c1318-greedy), [backtracking](#c1319-backtracking), and
  [divide and conquer](#c1311-divide-and-conquer).
- Relationship methods: [depth-first
  search](#c1320-depth-first-search), [breadth-first
  search](#c1321-breadth-first-search), [topological
  sort](#c1322-topological-sort),
  [Dijkstra](#c1323-dijkstra-shortest-paths),
  [Bellman-Ford](#c1324-bellman-ford-shortest-paths), [disjoint-set
  union](#c1325-disjoint-set-union), and [minimum spanning
  tree](#c1326-minimum-spanning-tree).

<a id="common-prefix-scan"></a> <a id="c131-common-prefix-scan"></a>

##### C.1.3.1 Compare from the left (common-prefix scan)

###### The problem it solves

Given several sequences, return the longest consecutive beginning
shared by every sequence. The result must start at position `0`, keep
the original order, and stop permanently at the first position where
one sequence ends or differs; matching values later cannot repair that
failed prefix.

###### How it works

Line up the sequences at their left edges and compare one vertical
column at a time. Stop at the first column where a sequence ends or a
value differs. The shared beginning before that point is a **common
prefix**; checking it from left to right is a **common-prefix scan**.

###### Step-by-step examples

For `strs = ["flower", "flow", "flight"]`, the expected output is
`"fl"`:

```text
column:   0 1 2 3 4 5
strs[0]:  f l o w e r
strs[1]:  f l o w
strs[2]:  f l i g h t
```

Each state shows the current `column`, the `row` being checked, the
`expectedByte` from `strs[0]`, and the known common prefix. `prefix`
means `strs[0][:column]`, already matched by every row.

Initial state: `column=0`, `row=1`, `expectedByte='f'`,
`prefix=""`.

1.  Match row `1`, column `0`
    - Decision: `strs[1][0]=='f'`; advance `row`.
    - After: `column=0`, `row=2`, `expectedByte='f'`,
      `prefix=""`.
    - Known correct: Rows `0` and `1` match at column `0`.
2.  Match row `2`, column `0`
    - Decision: `strs[2][0]=='f'`; finish the column.
    - After: `column=1`, `row=1`, `expectedByte='l'`,
      `prefix="f"`.
    - Known correct: Every string shares prefix `"f"`.
3.  Match row `1`, column `1`
    - Decision: `strs[1][1]=='l'`; advance `row`.
    - After: `column=1`, `row=2`, `expectedByte='l'`,
      `prefix="f"`.
    - Known correct: Rows `0` and `1` match through column `1`;
      every row still shares prefix `"f"`.
4.  Match row `2`, column `1`
    - Decision: `strs[2][1]=='l'`; finish the column.
    - After: `column=2`, `row=1`, `expectedByte='o'`,
      `prefix="fl"`.
    - Known correct: Every string shares prefix `"fl"`.
5.  Match row `1`, column `2`
    - Decision: `strs[1][2]=='o'`; advance `row`.
    - After: `column=2`, `row=2`, `expectedByte='o'`,
      `prefix="fl"`.
    - Known correct: Rows `0` and `1` match through column `2`;
      every row still shares prefix `"fl"`.
6.  Reject row `2`, column `2`
    - Decision: `strs[2][2]=='i'`; `'i' != 'o'`, so stop.
    - After: `result=strs[0][:2]="fl"`.
    - Known correct: Columns `0` and `1` match everywhere; column
      `2` does not. Therefore `"fl"` is longest.

For the boundary input `strs=["dog",""]`, the expected output is
`""`.

Initial state: `column=0`, `row=1`, `len(strs[1])=0`,
`prefix=""`.

1.  Stop before indexing the empty row
    - Decision: `column >= len(strs[1])`; do not index row `1`.
    - After: `result=strs[0][:0]=""`.
    - Known correct: No byte exists to match at column `0`.

For `strs=[]`, the expected output is also `""`.

Initial state: `len(strs)=0`.

1.  Handle no input rows
    - Decision: The empty-input guard runs before `strs[0]`.
    - After: `result=""`; neither loop starts.
    - Known correct: No missing row is indexed.

###### Why it works

Before each `column`, every string shares exactly the already-proven
part `strs[0][:column]`. This fact is true initially because
every string has the empty prefix. If every row reaches the column and
matches `expectedByte`, advancing `column` adds that byte to the known
prefix and preserves the fact. If a row ends or differs, a prefix
containing that column cannot be shared, and a prefix cannot skip the
failed column, so the already-proven prefix is the longest possible
answer.

Each successful outer step advances `column`, and `column` is bounded
by the first string's length, so the scan finishes. If no check fails,
every byte in the first string is shared and no longer prefix can
exist because the first string itself ends there.

###### Go example

The complete implementation is in [E.7 Common-prefix
scan](#e7-common-prefix-scan). Its variables `column`, `row`, and
`expectedByte` match the trace above. Its documented rules say that it
compares byte prefixes and returns `""` for empty input.

###### Work and memory

Let `S` be the total number of input bytes. A mismatch in the first
comparison does constant work, `O(1)`. When long prefixes match, each
examined byte belongs to an input string, so the worst-case work is
`O(S)`. The exact work stops at the first failed column. The scan
keeps only indexes and one expected byte, so it uses `O(1)` extra
memory. The returned substring is a view of the first string's bytes
rather than a separately built result, so it allocates no output byte
buffer but can keep the first string's storage reachable while the
result lives.

###### When to use it

Use this method when the requested result is explicitly a shared
prefix of strings or sequences aligned at position `0`, and comparison
can stop forever at the first ended or mismatching position.

###### When not to use it

Do not use it when the shared result may start later, skip positions,
tolerate edits, or compare reordered items. Those conditions break the
rule that one failed column permanently ends every possible answer.

###### Common mistakes and edge cases

> **Warning:** Check length before indexing. With `["dog", ""]`,
> reading `strs[1][0]` before testing its length causes a run-time
> panic.

> **Warning:** Byte positions are not always character positions. For
> text such as `"\u00e9"`, byte-by-byte comparison answers a
> byte-prefix question, not necessarily a user-visible-character
> question.

Empty input and a single string need defined results; this
implementation returns `""` and the single string, respectively.
Identical strings return the whole first string. Comparing only the
first and last input strings is valid only after an ordering argument,
not for arbitrary input order.

###### Related material

- [A.2 Correctness and invariants](#a2-correctness-and-invariants)
  explains the proof structure used by the proven prefix.
- [E.7 Common-prefix scan](#e7-common-prefix-scan) provides a complete
  Go implementation.
- [D.6 Bytes, runes, and visible
  characters](#d6-bytes-runes-and-visible-characters) explains how to
  choose the unit of text comparison.
- [Longest Common
  Prefix](problems/solved_easy_leetcode_0014_longest_common_prefix.go)
  applies this method to a judge problem.

<a id="directional-scan"></a> <a id="c132-directional-scan"></a>

##### C.1.3.2 Process the needed neighbor first (directional scan)

###### The problem it solves

Given a sequence in which each item's result depends on its immediate
left or right neighbor, compute every contribution in one pass. The
scan must begin at the side that makes the required neighbor available
before the current item is processed, rather than saving or revisiting
the whole sequence.

###### How it works

If each shelf label needs information from the shelf to its right,
start at the right and walk left. A **directional scan** chooses its
travel direction so the needed neighbor is already known.

###### Step-by-step examples

In the Roman numeral `IV`, `I` has value `1` and `V` has value `5`.
A smaller value immediately before a larger value is subtracted, so
the input values `[1,5]` should produce `4`. Scanning from right to
left makes the right neighbor available.

Initial state: `index=1`, `total=0`, `right=none`.

1.  Process position `1`
    - Decision: Add rightmost value `5`; it has no right neighbor.
    - After: `index=0`, `total=5`, `right=5`.
    - Known correct: The processed suffix has value `5`.
2.  Process position `0`
    - Decision: Subtract `1` because it is smaller than the value
      immediately to its right.
    - After: `index=done`, `total=4`, `right=1`.
    - Known correct: Every input value has been processed under
      the Roman numeral rule.

For one value `[5]`, add it directly; a nonexistent neighbor must not
be read.

###### Why it works

Before each step, the processed side is correct and `right` is exactly
the adjacent value required by the rule. Applying the rule extends the
correct suffix by one position. The index moves toward the opposite
edge each time, so the scan finishes.

###### Go example

The complete implementation is owned by [E.8 Directional
scan](#e8-directional-scan), whose right-to-left variables match this
trace. Calling `SumWithSmallerLeftValuesSubtracted([]int{1, 5})`
returns the traced result `4`.

###### Work and memory

Each of `n` values is visited once: `O(n)` time and `O(1)` extra
memory.

###### When to use it

Use it when one consistent adjacent side determines each local
decision.

###### When not to use it

Do not use it when the decision needs an arbitrary distant value or
information from both unprocessed sides.

###### Common mistakes and edge cases

> **Warning:** Update the remembered neighbor even after subtraction.
> Failing to do so gives the next position the wrong adjacent value.

Handle empty and one-item input explicitly, and derive the adjacent
rule from the problem rather than assuming every smaller-left value
should subtract.

###### Related material

- [E.8 Directional scan](#e8-directional-scan) explains the Go
  implementation.
- [Roman to
  Integer](problems/solved_easy_leetcode_0013_roman_to_integer.go) is
  the representative problem.
- [C.1.3.3 Simulation](#c133-simulation) handles rules needing more
  state.

<a id="c133-simulation"></a>

##### C.1.3.3 Follow the rules as state changes (simulation)

###### The problem it solves

Given a starting state and an ordered list of rules or commands,
return the state or emitted output after applying them in order. Each
step may depend on the state produced by earlier steps, so the
implementation must store every fact needed by the next transition
and respect whether updates are sequential or simultaneous.

###### How it works

Move a board-game piece one rule at a time while recording only its
position, direction, and other facts needed for the next move. Those
facts are the **state**. Following the rules directly is
**simulation**.

###### Step-by-step examples

A robot begins at `(0,0)` facing north. For `move, right, move`,
expect position `(1,1)` facing east.

Initial state: `position=(0,0)`, `direction=north`, `next=move`.

1.  Apply the first `move`
    - Decision: Moving north adds one to `y`.
    - After: `position=(0,1)`, `direction=north`, `next=right`.
    - Known correct: The state matches the process after one
      command.
2.  Apply `right`
    - Decision: Turn east without changing either coordinate.
    - After: `position=(0,1)`, `direction=east`, `next=move`.
    - Known correct: The state matches the process after two
      commands.
3.  Apply the final `move`
    - Decision: Moving east adds one to `x`.
    - After: `position=(1,1)`, `direction=east`, `next=none`.
    - Known correct: All three commands have been applied in
      order.

With no commands, the starting state is also the final state.

###### Why it works

The stored state equals the result of applying exactly the commands
already consumed. It starts true because the stored state matches the
robot's starting position and direction. Applying the next command by
its stated rule preserves the fact. Each step consumes one command, so
the finite command list ends. The returned state therefore includes
every command in order and is the requested final state. A command
with no rule is outside the accepted process, so rejecting the whole
sequence is correct.

###### Go example

The code stores north, east, south, and west as direction values `0`,
`1`, `2`, and `3`. The lookup contracts are
`horizontalStepByDirection[direction] -> change in x` and
`verticalStepByDirection[direction] -> change in y`. Thus direction
`1` selects horizontal step `1` and vertical step `0`, which means
east. The trace's position `(x,y)` is stored as `xCoordinate` and
`yCoordinate`. Calling
`runRobot([]string{"move", "right", "move"})` returns coordinates
`(1,1)`, direction `1` (east), and `true`.

```go
type robot struct{ xCoordinate, yCoordinate, direction int }

func runRobot(commands []string) (robot, bool) {
	// The zero value starts at (0,0) facing direction 0, which the
	// movement tables define as north.
	robotState := robot{}
	for _, command := range commands {
		switch command {
		case "right":
			// Direction cycles north, east, south, west; modulo wraps west to
			// north.
			robotState.direction = (robotState.direction + 1) % 4
		case "move":
			// Matching table indices turn the current direction into one
			// coordinate step without a separate branch for every compass
			// direction.
			horizontalStepByDirection := [...]int{0, 1, 0, -1}
			verticalStepByDirection := [...]int{1, 0, -1, 0}
			robotState.xCoordinate +=
				horizontalStepByDirection[robotState.direction]
			robotState.yCoordinate +=
				verticalStepByDirection[robotState.direction]
		default:
			// Reject the whole command sequence when its language is invalid.
			return robot{}, false
		}
	}
	return robotState, true
}
```

The boolean is false when a command is neither `move` nor `right`.

###### Work and memory

For `e` commands, the loop uses `O(e)` time and `O(1)` extra memory. A
process that must retain a board, history, or output must also count
that storage.

###### When to use it

Use it when the written transitions are affordable and fully determine
the next state.

###### When not to use it

Do not execute billions of individual events when the rules allow a
safe jump to a later state.

###### Common mistakes and edge cases

> **Warning:** A simultaneous update must read the old values.
> Updating one cell early can make a later cell observe the wrong
> round.

Test no commands, four right turns that wrap back to north, an invalid
command, and enough moves to approach the coordinate type's limits.

###### Related material

- [B.6 Matrix traversal and in-place
  state](#b6-matrix-traversal-and-in-place-state) explains matrix
  state rules.
- [C.1.3.8 Parsing and state
  machines](#c138-parsing-and-state-machines) handles inputs whose
  legal next action depends on a mode.

<a id="hash-map-or-set"></a> <a id="c134-hash-map-or-set"></a>

##### C.1.3.4 Remember facts by key (hash map or set)

###### The problem it solves

Given a sequence of values, answer repeated exact questions such as
whether a value appeared, how often it appeared, which earlier index
contains a needed partner, or which group owns a key. Rescanning all
earlier values for every question repeats work; store each fact under
the key used to ask for it.

###### How it works

A coat-check ticket lets you find one coat without searching every
rack. A **hash map** similarly stores a value under a key. A **set**
is like a guest list: it records only whether each name is present.
For the pair search below,
`indexByValue[value] -> an earlier index containing that value`. For
example, `indexByValue[2]=0` means input index `0` held `2`.

###### Step-by-step examples

For `[2,7,11,15]` and target `9`, expect `(0,1,true)`.

Initial state: `index=0`, `indexByValue={}`, `result=pending`.

1.  Process index `0`
    - Decision: Needed value `7` is absent, so store `2 -> 0`.
    - After: `index=1`, `indexByValue={2:0}`, `result=pending`.
    - Known correct: No earlier position exists, so no pair can
      end at index `0`.
2.  Process index `1`
    - Decision: Needed value `2` exists at index `0`.
    - After: `index=1`, `indexByValue={2:0}`,
      `result=(0,1,true)`.
    - Known correct: Values `2` and `7` total `9`, and their
      positions are distinct.

For `[3,3]`, querying before insertion lets the second `3` find the
first while preventing index `0` from pairing with itself.

The missing-target input `[1,2,3,4,5]` with target `20` makes the
saved lookup useful: a direct pair search checks all `10` pairs. The
map scan performs one complement lookup for each of the `5` values.

###### Why it works

Before `index`, the map contains exactly the defined facts from
earlier positions. A successful complement lookup therefore proves a
valid earlier partner; an unsuccessful lookup proves no retained
earlier partner exists. Storing the current value then preserves the
fact for the next index. The index advances each step, so the scan
finishes. If it finishes without a match, every value was checked
against every possible earlier partner, so no valid pair exists.

###### Go example

The complete implementation is owned by [E.3 Hash-map complement
lookup](#e3-hash-map-complement-lookup). Its `indexByValue` map has
the same stored-fact contract as the trace, and it checks that map
before inserting the current value. Calling
`PairSumIndices([]int{2, 7, 11, 15}, 9)` returns `(0,1,true)`.
Its mathematical difference `target-value` must fit in `int`.

###### Work and memory

The scan performs one lookup and at most one insertion per input
value. For ordinary, non-adversarial keys, Go maps usually make each
operation take a fixed amount of time. This gives expected `O(n)`
total time, but it is not a worst-case promise for every operation.
The map uses up to `O(n)` extra memory.

###### When to use it

Use it for exact lookup, frequency counting, partner search, grouping,
or presence tests where order between keys is irrelevant.

###### When not to use it

Do not use it for sorted traversal, nearest-key queries, or
relationships the chosen key does not preserve.

###### Common mistakes and edge cases

> **Warning:** Go returns a type's zero value for a missing key. Use
> `value, found := entries[key]` when zero is also a valid stored
> value.

Maps do not preserve insertion order. Grouping needs a **canonical
key**, one representation shared by logically equal objects. Anagrams
can use fixed character-count arrays for a fixed alphabet. Isomorphic
strings need mappings in both directions so two source characters
cannot collapse into one target.

###### Related material

- [D.4 Map presence and
  sets](#d4-map-presence-and-sets) explains Go map behavior.
- [Two Sum](problems/solved_easy_leetcode_0001_two_sum.go) is the
  representative problem.
- [C.1.3.33 Trie](#c1333-trie) supports shared-prefix lookup.

<a id="c135-sorting-and-selection"></a>

##### C.1.3.5 Put related values together (sorting and selection)

###### The problem it solves

Given values whose original order is irrelevant, either rearrange them
so related values become neighbors or return only the value at a
requested rank, such as the third smallest. Sorting exposes all
ordered relationships; selection avoids fully ordering values when
only one rank is required.

###### How it works

Arrange books by height and similar sizes become neighbors. That is
**sorting**. If you need only the third-shortest book, **selection**
arranges only enough items to find that rank.

###### Step-by-step examples

To detect a duplicate in `[4,1,4,2]`, expect `true`.

Initial state: `values=[4,1,4,2]`, `scanIndex=unset`,
`result=pending`.

1.  Sort
    - Decision: Put all values in increasing order.
    - After: `values=[1,2,4,4]`, `scanIndex=1`, `result=pending`.
    - Known correct: Sorting places equal values next to each
      other.
2.  Compare neighbors
    - Decision: Skip differing pairs `(1,2)` and `(2,4)`; match
      final pair `(4,4)`.
    - After: `values=[1,2,4,4]`, `scanIndex=3`, `result=true`.
    - Known correct: Two separate input positions contain the same
      value.

To select the fourth-smallest value from `[7,2,5,1,4,3,6]`, use
zero-based rank `3`. The trace below uses the last value of each
active range as its pivot, the value used to split that range.

Initial state: `ordered=[7,2,5,1,4,3,6]`, `rank=3`,
`range=[0,7)`, `result=pending`.

1.  Partition around pivot `6`
    - Decision: Move `2,5,1,4,3` before pivot `6`; leave `7`
      after it. The pivot reaches index `5`.
    - After: `ordered=[2,5,1,4,3,6,7]`, `rank=3`,
      `range=[0,5)`,
      `result=pending`.
    - Known correct: Pivot index `5` is final; requested rank `3`
      can occur only to its left. That remaining side is not sorted.
2.  Partition around pivot `3`
    - Decision: Keep `2` and `1` before pivot `3`; its index becomes
      `2` while `5` and `4` remain on the other side.
    - After: `ordered=[2,1,3,4,5,6,7]`, `rank=3`,
      `range=[3,5)`,
      `result=pending`.
    - Known correct: Pivot index `2` is final; requested rank `3`
      can occur only to its right.
3.  Partition around pivot `5`
    - Decision: Keep `4` before pivot `5`; its index becomes `4`.
    - After: `ordered=[2,1,3,4,5,6,7]`, `rank=3`,
      `range=[3,4)`, `result=pending`.
    - Known correct: Requested rank `3` can occur only to the
      pivot's left.
4.  Finish the one-value range
    - Decision: Only `ordered[3]=4` remains; no partition is needed.
    - After: `ordered=[2,1,3,4,5,6,7]`, `rank=3`,
      `range=[3,4)`, `result=4`.
    - Known correct: `4` is the value a full sort would place at
      zero-based rank `3`, even though positions `0` and `1` remain
      unsorted.

Repeated values occupy separate ranks. For `[2,2,3]`, ranks `0` and
`1` both contain `2`.

###### Why it works

After sorting, all copies of one value are next to one another. The
scan checks every neighboring pair, so finding an equal pair proves a
duplicate exists. If the finite scan ends without one, no value occurs
twice.

Quickselect partitions the active range so one pivot has a rank it
could occupy after a full sort. Values before it are no greater, and
values after it are no smaller. The requested rank can lie on only one
side, so discarding the other side cannot discard the answer. A
partition must shrink the active range; when the pivot reaches the
requested rank, it is the value a full sort would place there. Equal
values require the partition's chosen rank rule to make progress.

###### Go example

For the duplicate trace, `hasDuplicate([]int{4, 1, 4, 2})` returns
`true`; the trace's sorted `values` are the code's copied `ordered`
slice. For the selection trace,
`kthSmallest([]int{7, 2, 5, 1, 4, 3, 6}, 3)` returns `(4, true)`;
its `ordered` slice and zero-based `rank` use the trace's names.

```go
import "sort"

func hasDuplicate(values []int) bool {
	// Sort a copy so the caller's order remains unchanged.
	ordered := append([]int(nil), values...)
	sort.Ints(ordered)
	// Sorting makes equal values adjacent, reducing duplicate detection
	// to one comparison per neighboring pair.
	for currentIndex := 1; currentIndex < len(ordered); currentIndex++ {
		if ordered[currentIndex] == ordered[currentIndex-1] {
			return true
		}
	}
	return false
}

func kthSmallest(values []int, rank int) (int, bool) {
	// A rank is a zero-based position in the fully sorted order.
	if rank < 0 || rank >= len(values) {
		return 0, false
	}
	// Partition a copy so selection does not change the caller's slice.
	ordered := append([]int(nil), values...)
	left, right := 0, len(ordered)
	for left < right {
		if right-left == 1 {
			return ordered[left], true
		}
		// Use the last active value as the pivot. Values no larger than
		// it are gathered at the active range's left side.
		pivotIndex := right - 1
		pivotValue := ordered[pivotIndex]
		nextSmallerOrEqual := left
		for scanIndex := left; scanIndex < pivotIndex; scanIndex++ {
			if ordered[scanIndex] <= pivotValue {
				ordered[nextSmallerOrEqual], ordered[scanIndex] =
					ordered[scanIndex], ordered[nextSmallerOrEqual]
				nextSmallerOrEqual++
			}
		}
		// Put the pivot between the two partitions. Its index is final,
		// although values inside each side remain unsorted.
		ordered[nextSmallerOrEqual], ordered[pivotIndex] =
			ordered[pivotIndex], ordered[nextSmallerOrEqual]
		if nextSmallerOrEqual == rank {
			return ordered[nextSmallerOrEqual], true
		}
		if rank < nextSmallerOrEqual {
			right = nextSmallerOrEqual
		} else {
			left = nextSmallerOrEqual + 1
		}
	}
	return 0, false
}
```

###### Work and memory

Comparison sorting uses `O(n log n)` time. Each function above copies
`n` values and therefore uses `O(n)` extra memory. Quickselect does
linear work on average when input order makes its pivots split the
active range reasonably well. Repeatedly choosing its smallest or
largest value makes it rescan nearly the whole remaining range and
takes `O(n^2)` time. Counting over a small range of `k` possible keys
takes `O(n+k)` time and `O(k)` memory. Returning an ordered copy also
needs `O(n)` output space.

###### When to use it

Use sorting when input order is irrelevant and order makes duplicates,
intervals, pairs, or greedy choices easy to process. Use selection for
one rank.

###### When not to use it

Avoid it when original order cannot be carried along, linear time is
required, or a set, counting array, heap, or direct selection is
simpler.

###### Common mistakes and edge cases

> **Warning:** `sort.Ints` and `sort.Slice` mutate their slices. Copy
> when the caller's order must survive. Use **stable sorting** when
> equal records must also remain in their original relative order.

A comparator must be consistent and include required tie-breaking. A
huge key range makes counting storage impractical.

###### Related material

- [D.3 Sorting without unexpected
  mutation](#d3-sorting-without-unexpected-mutation) explains when Go
  sorting changes the caller's slice.
- [C.1.3.27 Heap or priority queue](#c1327-heap-or-priority-queue)
  compares streaming top-`k` work.
- [C.1.3.30 Intervals](#c1330-intervals) uses sorted starts.

<a id="binary-search"></a>

<a id="c136-binary-search"></a>

##### C.1.3.6 Find a changing boundary (binary search)

###### The problem it solves

Given an ordered search space, return an exact target or the first or
last position where a yes/no condition changes. A comparison at the
middle must prove that one entire half cannot contain the requested
boundary; if the answers can switch back and forth, discarding half
can lose the result.

###### How it works

To find a word in a dictionary, open near the middle and discard the
half that cannot contain it. **Binary search** is safe only when one
test splits the possibilities into two solid groups, such as failures
then successes. Each step must remove at least one remaining
possibility.

###### Step-by-step examples

For `[1,3,3,5]`, find the first value at least `3`; expect index `1`.

Initial state: `left=0`, `right=4`, `result=pending`.

1.  Check range `[0,4)`
    - Decision: Middle index `2` has value `3`, so a first
      valid position cannot be later than `2`. Set `right=2`.
    - After: `left=0`, `right=2`, `result=pending`.
    - Known correct: Positions `0` and `1` remain unchecked;
      position `2` is the earliest known success.
2.  Check range `[0,2)`
    - Decision: Middle index `1` has value `3`, so set
      `right=1`.
    - After: `left=0`, `right=1`, `result=pending`.
    - Known correct: Position `0` remains unchecked; position `1`
      is now the earliest known success.
3.  Check range `[0,1)`
    - Decision: Middle index `0` has value `1`; set `left=1`.
    - After: `left=1`, `right=1`, `result=1`.
    - Known correct: Position `0` failed and position `1` is the
      first valid position.

For an empty slice, the same half-open bounds start and finish at `0`,
the insertion position.

###### Why it works

Before every test, positions before `left` are known failures, while
`right` is the earliest known success or the position after the
slice. Only positions in `[left,right)` remain unchecked. A failed
middle proves that it and every earlier unchecked position also fail,
so the next unchecked position is `middle+1`. A successful middle
becomes the new earliest known success. Each update makes the
unchecked range smaller. When it is empty, the failure and success
boundaries meet, so `left` is the first acceptable position.

###### Go example

The complete lower-bound implementation is owned by [E.1 Binary search
bounds](#e1-binary-search-bounds); its `left`, `right`, and `middle`
variables match this trace. Calling `LowerBound([]int{1, 3, 3, 5}, 3)`
returns `1`.

###### Work and memory

Each test removes about half of the remaining positions. After `k`
tests, roughly `n/2^k` positions remain, so the number of tests grows
as `log n`. A loop-based search therefore takes `O(log n)` time and
uses `O(1)` extra memory.

###### When to use it

Data is sorted, the problem asks for the first or last position
satisfying a condition, or an ordered numeric answer space has a
single boundary where a yes/no test changes. A numeric answer alone is
not enough.

###### When not to use it

Testing a later value can change from yes back to no, or the middle
test cannot prove that one entire half is impossible.

###### Common mistakes and edge cases

> **Warning:** Keeping a tested middle in the unchecked range can
> prevent progress. With two positions left, an update that changes
> neither boundary repeats the same test forever.

A valid insertion position may equal `len(values)`. Ordinary bounds
require fully sorted input. Rotated arrays and peak searches need
different facts that prove which side is safe to discard; do not copy
this lower-bound update rule into those problems without deriving the
new boundary proof.

###### Related material

- [E.1 Binary search bounds](#e1-binary-search-bounds) explains the Go
  implementation.
- [Search Insert Position](
  problems/solved_easy_leetcode_0035_search_insert_position.go
  ) and [Search in Rotated Sorted Array](
  ../REFERENCE_LINKS.md#leetcode-0033
  )
  are representative problems.

<a id="c137-string-matching"></a>

##### C.1.3.7 Reuse a partial match (string matching)

###### The problem it solves

Given a pattern and a longer text, return the first position where the
exact pattern occurs, or report that it does not occur. After a
mismatch, a naive search restarts the pattern and can compare the same
text characters many times; saved information about matching pattern
prefixes must identify the next viable partial match.

###### How it works

When a pattern match fails, keep the longest matched ending that is
also the pattern's beginning. The **Knuth-Morris-Pratt algorithm
(KMP)** prepares these backup lengths so it can reuse known matches
instead of restarting from zero. For an end position `i`,
`prefixLengths[i]` stores how many bytes at the start of the pattern
also appear at the end of `pattern[:i+1]`, excluding that entire
piece. The position selects the table entry; the stored number is the
partial-match length to keep after a failure.

###### Step-by-step examples

Example 1 builds the table for pattern `"ababac"`. The expected
`prefixLengths` is `[0,0,1,2,3,0]`.

Initial state: `patternIndex=1`, `overlapLength=0`,
`prefixLengths=[0,0,0,0,0,0]`.

1.  Compare `b` with `a`
    - Decision: They differ with no backup; store `0` and advance.
    - After: `patternIndex=2`, `overlapLength=0`,
      `prefixLengths=[0,0,0,0,0,0]`.
    - Known correct: Positions `0` and `1` have final answers.
2.  Compare `a` with `a`
    - Decision: Grow the overlap to `1`, store it, and advance.
    - After: `patternIndex=3`, `overlapLength=1`,
      `prefixLengths=[0,0,1,0,0,0]`.
    - Known correct: Prefix and suffix `"a"` match through index `2`.
3.  Compare `b` with `b`
    - Decision: Grow the overlap to `2`, store it, and advance.
    - After: `patternIndex=4`, `overlapLength=2`,
      `prefixLengths=[0,0,1,2,0,0]`.
    - Known correct: Prefix and suffix `"ab"` match through index `3`.
4.  Compare `a` with `a`
    - Decision: Grow the overlap to `3`, store it, and advance.
    - After: `patternIndex=5`, `overlapLength=3`,
      `prefixLengths=[0,0,1,2,3,0]`.
    - Known correct: Prefix and suffix `"aba"` match through index
      `4`.
5.  Reject the three-byte backup
    - Decision: `c != b`; set overlap to `prefixLengths[2]=1`.
    - After: `patternIndex=5`, `overlapLength=1`,
      `prefixLengths=[0,0,1,2,3,0]`.
    - Known correct: The same `c` can still extend a smaller backup.
6.  Reject the one-byte backup
    - Decision: `c != b`; set overlap to `prefixLengths[0]=0`.
    - After: `patternIndex=5`, `overlapLength=0`,
      `prefixLengths=[0,0,1,2,3,0]`.
    - Known correct: No nonempty backup remains.
7.  Confirm zero at index `5`
    - Decision: `c != a`; store `0` and advance.
    - After: `patternIndex=6`, `overlapLength=0`,
      `prefixLengths=[0,0,1,2,3,0]`.
    - Known correct: Every table position has its longest backup.

Example 2 shows a smaller backup succeeding for pattern `"aabaaab"`.

Initial state: `patternIndex=5`, `overlapLength=2`,
`prefixLengths=[0,1,0,1,2,0,0]`.

1.  Fall back once
    - Decision: `pattern[5]='a'` differs from `pattern[2]='b'`; use
      `prefixLengths[1]=1`.
    - After: `patternIndex=5`, `overlapLength=1`,
      `prefixLengths=[0,1,0,1,2,0,0]`.
    - Known correct: The same current byte remains unconsumed.
2.  Grow the smaller backup
    - Decision: `pattern[5]=='a'` matches `pattern[1]=='a'`; store
      `2`.
    - After: `patternIndex=6`, `overlapLength=2`,
      `prefixLengths=[0,1,0,1,2,2,0]`.
    - Known correct: The longest prefix-suffix match at index `5`
      has length `2`.

Example 3 uses the table while searching. Search for pattern
`"ababac"` in text `"abababac"`. The answer should be position `2`.

The pattern's table is `[0,0,1,2,3,0]`.

Initial state: `textIndex=0`, `matchedLength=0`, `result=pending`.

1.  Match the first five bytes
    - Decision: Compare equal byte pairs while advancing both
      positions.
    - After: `textIndex=5`, `matchedLength=5`, `result=pending`.
    - Known correct: Those five text bytes equal pattern positions
      `0` through `4`.
2.  Keep the longest valid backup
    - Decision: Text `'b'` differs from pattern `'c'`. Keep the text
      position, read `prefixLengths[4]=3`, and retain the final
      `"aba"` of the matched text.
    - After: `textIndex=5`, `matchedLength=3`, `result=pending`.
    - Known correct: The current `'b'` remains unconsumed, and the
      kept `"aba"` equals the first three pattern bytes.

    ```text
    text:    a b a b a | b a c
    pattern:     a b a | b a c
    ```

3.  Finish from the smaller backup
    - Decision: Match the remaining bytes `"bac"`.
    - After: `textIndex=8`, `matchedLength=6`, `result=2`.
    - Known correct: Every pattern byte matches text positions `2`
      through `7`, so the first complete match starts at position `2`.

Small boundary cases: searching text `"x"` for pattern `"a"`
returns `-1`: no bytes matched, so the search simply advances through
the text. An empty pattern returns `0`. A pattern longer than the text
cannot complete and returns `-1`.

A direct search is usually easier to write, but it may compare the
same text bytes many times. A rolling hash compares fingerprints of
pattern-sized windows; equal fingerprints still require a byte check
because different windows can share a fingerprint.

###### Why it works

Every backup answer is proved before it is stored. For example,
storing `3` for the matched piece `"ababa"` means the builder directly
confirmed this relationship:

```text
pattern beginning:  a b a
matched ending:         a b a
```

Therefore, if the next comparison fails, keeping those three bytes is
safe. They already equal the first three pattern bytes.

When a backup count such as `3` also fails, the builder needs the
backup answer for a three-byte matched piece. The count-to-position
conversion selects that previously proved answer:

```text
matched count:     3
end position:      3 - 1 = 2
backup answer:     prefixLengths[2] = 1
```

The builder tries these proved backup counts from largest to smallest.
It therefore discards only choices that have failed and does not skip
a smaller choice that could work.

The search uses the same answers. It returns only after every pattern
byte has matched. Each loop pass either moves to a new byte or changes
a positive backup count to a smaller one, so both phases finish.

###### Go example

```go
func buildPrefixLengths(pattern string) []int {
	prefixLengths := make([]int, len(pattern))
	patternIndex := 1
	overlapLength := 0

	// overlapLength is the backup count being tested for the piece
	// immediately before patternIndex.
	for patternIndex < len(pattern) {
		if pattern[patternIndex] == pattern[overlapLength] {
			overlapLength++
			prefixLengths[patternIndex] = overlapLength
			patternIndex++
			continue
		}
		if overlapLength == 0 {
			patternIndex++
			continue
		}

		// Convert the failed backup count to its end position. That
		// entry contains the next backup count to try.
		backupPieceEndIndex := overlapLength - 1
		overlapLength = prefixLengths[backupPieceEndIndex]
	}
	return prefixLengths
}

func firstMatchKMP(text, pattern string) int {
	if len(pattern) == 0 {
		return 0
	}

	prefixLengths := buildPrefixLengths(pattern)
	textIndex := 0
	matchedLength := 0

	// The processed text ends with the first matchedLength bytes of
	// the pattern. Compare the next text and pattern bytes.
	for textIndex < len(text) {
		if text[textIndex] == pattern[matchedLength] {
			textIndex++
			matchedLength++
			if matchedLength == len(pattern) {
				return textIndex - matchedLength
			}
			continue
		}
		if matchedLength == 0 {
			textIndex++
			continue
		}

		// Convert the matched count to its end position, read its
		// backup count, and retry the same text byte.
		matchedPieceEndIndex := matchedLength - 1
		matchedLength = prefixLengths[matchedPieceEndIndex]
	}
	return -1
}
```

The code compares bytes. It needs no imports.

###### Work and memory

Let `n` be the number of text bytes and `m` the number of pattern
bytes. A direct search may compare as many as `m` bytes at each text
start, taking `O(nm)` time and `O(1)` extra memory.

KMP builds the table in `O(m)` time and scans the text in `O(n)`
time. It therefore takes `O(n + m)` total time and stores `m` table
entries, using `O(m)` auxiliary memory.

A rolling hash takes `O(m)` setup and `O(n)` hash updates. Candidate
byte comparisons can make its worst case `O(nm)`, although expected
time is `O(n + m)` when collisions and repeated candidates are rare.

###### When to use it

Use KMP when the task requires an exact substring, inputs can be
large, and a worst-case linear guarantee matters. It is especially
useful when repeated pattern pieces would make a direct search compare
the same text bytes again.

###### When not to use it

A short or one-off search is often clearer with a direct scan or a
standard library operation. KMP also does not model approximate
matches, reordered characters, or other non-exact relationships.

###### Common mistakes and edge cases

> **Warning:** The slice stores backup lengths, but its indexes are
> end positions. Keep those two roles separate:

```text
matched count:        5
lookup position:      5 - 1 = 4
stored backup length: prefixLengths[4] = 3
```

> **Warning:** Do not move past the current pattern or text byte
> during a fallback. That byte failed against one pattern position,
> but it may match the next position after the pattern is shortened.

Define the empty-pattern result before scanning. This implementation
returns position `0`.

Byte-based matching is correct for byte-defined input or a one-byte
alphabet. It can report a position inside the multi-byte encoding of
a Unicode character when the task instead defines character or
visible-symbol positions.

###### Related material

- [D.6 Bytes, runes, and visible
  characters](#d6-bytes-runes-and-visible-characters) explains byte
  versus text positions.
- [C.1.3.1 Common-prefix scan](#c131-common-prefix-scan) compares
  aligned starts rather than searching every start.

<a id="c138-parsing-and-state-machines"></a>

##### C.1.3.8 Track legal input modes (parsing and state machines)

###### The problem it solves

Given text whose next legal symbol depends on what has already been
read, decide whether the whole input is valid or convert it into a
structured value. The parser must remember its current mode, such as
"inside digits" or "after a decimal point," because the same symbol
can be legal in one mode and illegal in another.

###### How it works

A vending machine accepts different buttons after money, a drink
choice, or a refund. Each situation is a **state**, and the next input
chooses the next state and action. A parser uses the same **state
machine** to enforce a **grammar**, the allowed input shapes. A
**token** is one meaningful piece such as a number or operator.

###### Step-by-step examples

Parse `"  -42x"` with the rule
`space -> optional sign -> digits -> stop`; expect `-42`.

Initial state: `currentIndex=0`, `mode=spaces`, `negative=false`,
`magnitude=0`, `result=pending`.

1.  Consume leading spaces
    - Decision: Skip the two spaces allowed by `mode=spaces`.
    - After: `currentIndex=2`, `mode=sign`, `negative=false`,
      `magnitude=0`, `result=pending`.
    - Known correct: The consumed beginning is legal.
2.  Consume the optional sign
    - Decision: Record `'-'` and enter digit mode.
    - After: `currentIndex=3`, `mode=digits`, `negative=true`,
      `magnitude=0`, `result=pending`.
    - Known correct: A second sign is no longer legal.
3.  Consume digit `4`
    - Decision: Append digit `4` after checking the limit.
    - After: `currentIndex=4`, `mode=digits`, `negative=true`,
      `magnitude=4`, `result=pending`.
    - Known correct: The stored magnitude equals parsed text
      `"4"`.
4.  Consume digit `2`
    - Decision: Append digit `2` after checking the limit.
    - After: `currentIndex=5`, `mode=digits`, `negative=true`,
      `magnitude=42`, `result=pending`.
    - Known correct: The stored magnitude equals parsed text
      `"42"`.
5.  Stop at `x`
    - Decision: This grammar accepts the longest numeric prefix,
      so stop and apply the negative sign.
    - After: `currentIndex=5`, `mode=done`, `negative=true`,
      `magnitude=42`, `result=-42`.
    - Known correct: Every accepted byte was interpreted once;
      `'x'` was not consumed.

For `"4-2"`, the `-` is not legal after digits. This entry's
stop-at-first-nondigit grammar returns `4`. A whole-input grammar
would instead reject it.

###### Why it works

After consuming a prefix, the state describes every legal way that
prefix may continue and the accumulated values represent exactly the
completed tokens. The empty starting state makes this true before any
input is read, and each legal transition preserves it. Every
transition consumes input or stops, so the parser finishes. At that
point it either rejects a missing required part or returns the value
represented by the accepted prefix.

###### Go example

Calling `parseSignedPrefix("  -42x")` returns `(-42, true)`. The
trace's `mode` is represented by the code's ordered phases: the first
loop consumes spaces, the sign branch runs once, the digit loop
builds `magnitude`, and returning applies the sign. `currentIndex`,
`negative`, and `magnitude` keep the trace's names.

```go
import "math"

func parseSignedPrefix(input string) (int64, bool) {
	currentIndex, negative := 0, false
	// Skip only the leading space bytes allowed by this parser's
	// contract.
	for currentIndex < len(input) && input[currentIndex] == ' ' {
		currentIndex++
	}
	// Consume at most one optional sign before any digits.
	if currentIndex < len(input) &&
		(input[currentIndex] == '+' || input[currentIndex] == '-') {
		if input[currentIndex] == '-' {
			negative = true
		}
		currentIndex++
	}
	// A negative int64 may have magnitude MaxInt64+1, so select the
	// allowable unsigned magnitude before accumulating digits.
	limit := uint64(math.MaxInt64)
	if negative {
		limit++
	}
	firstDigitIndex := currentIndex
	var magnitude uint64
	// Parse the maximal digit prefix and rearrange the limit test so
	// overflow is rejected before multiplying the current magnitude by
	// ten.
	for currentIndex < len(input) &&
		input[currentIndex] >= '0' &&
		input[currentIndex] <= '9' {
		digit := uint64(input[currentIndex] - '0')
		if magnitude > (limit-digit)/10 {
			return 0, false
		}
		magnitude = magnitude*10 + digit
		currentIndex++
	}
	if currentIndex == firstDigitIndex {
		// Whitespace and a sign alone do not constitute a parsed integer.
		return 0, false
	}
	if negative {
		// MinInt64 cannot be formed by negating a positive int64, so return
		// the named constant for its one larger unsigned magnitude.
		if magnitude == uint64(math.MaxInt64)+1 {
			return math.MinInt64, true
		}
		return -int64(magnitude), true
	}
	return int64(magnitude), true
}
```

The fragment matches the stated stop-at-first-nondigit grammar and
rejects a magnitude outside the `int64` range before multiplying.

###### Work and memory

Reading `n` input bytes once takes `O(n)` time. A simple number
parser can use `O(1)` extra memory. Nested expressions may need a
stack whose size is proportional to the nesting depth, and retaining
every token can require `O(n)` memory.

###### When to use it

Input has a grammar or phased format: optional whitespace, an optional
sign, digits, perhaps a decimal or exponent, nested delimiters, or
operators with precedence. A collection of unrelated boolean flags is
often an implicit state machine that should be made explicit.

###### When not to use it

The input has no meaningful grammar or phases, or a well-tested
standard parser already implements exactly the required language.

###### Common mistakes and edge cases

> **Warning:** Without the pre-operation bound check, parsing
> `"9223372036854775808"` overflows `int64` before the function can
> reject it.

An empty string, spaces with no digits, or a sign with no following
digit must return failure. State whether trailing nondigits stop a
valid prefix, as here, or make the whole input invalid.

###### Related material

- [A.4 Integer limits and bits](#a4-integer-limits-and-bits) explains
  overflow bounds.
- [C.1.3.13 Stack](#c1313-stack) supports nested delimiters and
  expression evaluation.

<a id="c139-two-pointers"></a> <a id="two-pointers"></a>

##### C.1.3.9 Move coordinated positions (two pointers)

###### The problem it solves

Given one or two sequences whose position order matters, find a pair,
range, or compacted result by moving two positions through the data.
After each comparison, the chosen movement must prove that every
skipped candidate is impossible or already handled; otherwise, moving
a pointer could miss a valid answer.

###### How it works

Imagine two people checking a shelf from opposite ends and moving only
the person whose side cannot give the answer. The two changing
positions are **pointers**. Coordinating them avoids restarting the
scan. Here a pointer means an index or position, not necessarily a Go
memory pointer.

The movement depends on the task. In a sorted pair-sum search, start
at both ends: move the left position when the sum is too small and the
right position when it is too large. Other tasks may coordinate the
next unread positions of two inputs, a read and write position in one
slice, or two symmetric positions. Each version needs its own reason
that moving a position cannot discard required work.

###### Step-by-step examples

Find a pair totaling `9` in `[1,2,4,7,11]`; expect index result
`(1,3,true)`.

Initial state: `left=0`, `right=4`, `result=pending`.

1.  Check values `1` and `11`
    - Decision: Sum `12` is too large; decrement `right`.
    - After: `left=0`, `right=3`, `result=pending`.
    - Known correct: No pair using `11` can total `9`, because `1`
      was its smallest available partner.
2.  Check values `1` and `7`
    - Decision: Sum `8` is too small; increment `left`.
    - After: `left=1`, `right=3`, `result=pending`.
    - Known correct: No pair using `1` can total `9`, because `7`
      was its largest available partner.
3.  Check values `2` and `7`
    - Decision: Their sum is `9`, so return them.
    - After: `left=1`, `right=3`, `result=(1,3,true)`.
    - Known correct: The returned values are at distinct positions
      and total the target.

An empty or one-item input has no pair and must finish before indexing
both ends.

###### Why it works

No pair using a position already passed by `left` or `right` can meet
the target. If the sum is too small, `right` already supplies the
largest available partner for `values[left]`, so that left value
cannot work with any remaining position. The too-large case is
symmetric: `left` is the smallest available partner for
`values[right]`. Equality returns a valid pair. Otherwise one boundary
moves inward, so the finite range eventually becomes empty.

###### Go example

The complete implementation is owned by [E.9 Sorted two-pointer pair
search](#e9-sorted-two-pointer-pair-search). Its boundary checks and
movements match this trace. Calling
`PairSumSorted([]int{1, 2, 4, 7, 11}, 9)` returns `(1,3,true)`.

###### Work and memory

For `n` sorted values, each boundary moves inward at most `n` times,
so pair search takes `O(n)` time and `O(1)` extra memory. A two-input
merge or a read/write compaction can also be linear, but a returned
merged sequence needs space for its output.

###### When to use it

Look for sorted pair selection, comparisons between two ends, merging
sorted inputs, modifying a sequence in place, or testing symmetric
positions. For every movement, you must explain why the position left
behind cannot be needed again.

###### When not to use it

No pointer movement can be proved to preserve every remaining answer,
or the method would require sorting data whose original positions or
order must remain intact.

###### Common mistakes and edge cases

> **Warning:** The inward pair-sum movement requires sorted values.
> On `[4,1,7,2]`, an end sum gives no safe reason to discard either
> endpoint. Sorting may also destroy required original positions, so
> preserve positions separately when the output needs them.

###### Related material

- [E.9 Sorted two-pointer pair
  search](#e9-sorted-two-pointer-pair-search) explains the Go
  implementation.
- [A.2 Correctness and invariants](#a2-correctness-and-invariants)
  traces read/write compaction.
- [Palindrome
  Number](problems/solved_easy_leetcode_0009_palindrome_number.go),
  [Longest Palindromic Substring](
  ../REFERENCE_LINKS.md#leetcode-0005
  ), [Remove Duplicates from Sorted Array](
  ../REFERENCE_LINKS.md#leetcode-0026
  ),
  and [Remove
  Element](problems/solved_easy_leetcode_0027_remove_element.go) are
  representative problems.

<a id="dynamic-programming"></a>

<a id="c1310-dynamic-programming"></a>

##### C.1.3.10 Reuse smaller answers (dynamic programming)

###### The problem it solves

Given a problem with several choices at each step, compute the best,
counted, or feasible result when different choice paths repeatedly
reach the same smaller question. Define what one saved answer means,
then compute each distinct smaller state once instead of rebuilding it
for every path that reaches it.

###### How it works

If several delivery routes return to the same intersection, solve the
remaining trip from that intersection once and reuse it. **Dynamic
programming (DP)** stores answers to repeated smaller questions. The
few facts naming one such question are its **state**. Define precisely
what one state answers, write how it depends on smaller states, set
the smallest known answers, then evaluate states only after the states
they need. The dependency rule is the **recurrence**. **Memoization**
computes and saves states on demand; **tabulation** fills a table in a
planned order. In the stair example,
`ways[stair] -> number of routes that end exactly at stair`.
For example, `ways[2]=2` records routes `[1,1]` and `[2]`.

###### Step-by-step examples

Count ways to climb `4` stairs using steps of one or two; expect
`(5,true)`.

A direct recursive search repeats the same smaller question:

```text
ways(4)
|-- ways(3)
|   |-- ways(2)
|   `-- ways(1)
`-- ways(2)  <- computed again
```

The table below stores `ways(2)` once and reuses it from both places.

Initial state: `stair=0`, `ways=[]`, `result=pending`.

1.  Store the base answers
    - Decision: Set `ways[0]=1` and `ways[1]=1`.
    - After: `stair=2`, `ways=[1,1]`, `result=pending`.
    - Known correct: There is one empty route to stair `0` and one
      one-step route to stair `1`.
2.  Compute stair `2`
    - Decision: A final one-step move comes from stair `1`; a
      final two-step move comes from stair `0`. Add both counts.
    - After: `stair=3`, `ways=[1,1,2]`, `result=pending`.
    - Known correct: Routes `[1,1]` and `[2]` are counted once.
3.  Compute stair `3`
    - Decision: Add routes ending with a one-step move and routes
      ending with a two-step move.
    - After: `stair=4`, `ways=[1,1,2,3]`, `result=pending`.
    - Known correct: The two final-move groups do not overlap.
4.  Compute stair `4`
    - Decision: Set `ways[4]=ways[3]+ways[2]=5`.
    - After: `stair=done`, `ways=[1,1,2,3,5]`,
      `result=(5,true)`.
    - Known correct: Every route reaches stair `4` with exactly
      one of the two allowed final moves.

For `n=0`, the base value `1` represents the empty way; another
problem may define that boundary differently.

###### Why it works

The base values are correct: stair `0` has one empty route, and stair
`1` has one one-step route. From stair `2` onward, every route has one
final move. Removing it leaves a route to either the previous
stair or the stair before that. These two groups cover every allowed
route and do not overlap, so their saved counts add to the exact new
count. Processing stairs in increasing order makes both earlier counts
available. The finite loop ends with the exact count for stair `n`.

###### Go example

Calling `countStairWays(4)` builds the traced table
`[]int{1, 1, 2, 3, 5}` and returns `(5,true)`.

```go
func countStairWays(numStairs int) (int, bool) {
	if numStairs < 0 {
		return 0, false
	}
	// ways[stair] counts routes ending exactly at that stair. The
	// empty route is the one way to end at stair zero.
	ways := make([]int, numStairs+1)
	ways[0] = 1
	if numStairs >= 1 {
		ways[1] = 1
	}
	// Every later route ends with either a one-step or two-step move,
	// so the two earlier nonoverlapping groups add to the new count.
	for stair := 2; stair <= numStairs; stair++ {
		ways[stair] = ways[stair-1] + ways[stair-2]
	}
	return ways[numStairs], true
}
```

[E.2 Rolling dynamic programming](#e2-rolling-dynamic-programming)
keeps the same recurrence but retains only its two earlier states.

###### Work and memory

Without caching, the stair search grows roughly like `2^n`. Saving the
answer for each stair reduces it to `O(n)` time and `O(n)` memory.
When the next answer needs only a fixed number of earlier answers,
keep just those recent answers and reduce extra memory to `O(1)`.

Derive complexity from the number of reachable states multiplied by
the work per transition. A two-dimensional table is not automatically
`O(n^2)` if each cell loops over many choices. Memoization does not
help states that are never repeated.

###### When to use it

All three conditions should hold:

1.  The full answer can be built from answers to smaller versions of
    the problem.
2.  A direct search answers the same smaller question more than once.
3.  A small set of facts completely describes each smaller question.

For climbing stairs, the current stair number is enough state because
the choices available next do not depend on how that stair was
reached. Describe the state in one sentence before choosing an array
or map to store answers.

###### When not to use it

The smaller questions do not repeat, the saved facts leave out
something that affects future choices, or there are still too many
different smaller questions for the input limits.

###### Common mistakes and edge cases

Writing the rule alone does not prevent repeated work; the answers
must also be saved. Do not store facts that cannot affect future
choices. Make the smallest known answers agree with the rule, update
recent values without destroying one that is still needed, count
memory used by unfinished recursive calls, and check whether the count
can exceed the chosen integer type.

###### Related material

- [E.2 Rolling dynamic programming](#e2-rolling-dynamic-programming)
  provides the Go implementation.
- [C.1.3.10.1 Preserve order (coordinate
  compression)](#c13101-coordinate-compression) explains
  coordinate compression.
- [Climbing
  Stairs](problems/solved_easy_leetcode_0070_climbing_stairs.go) is
  the representative problem.

<a id="c13101-coordinate-compression"></a>

##### C.1.3.10.1 Preserve order (coordinate compression)

###### The problem it solves

Given possibly negative or very large values, replace each distinct
value with one of the consecutive indexes starting at zero while
preserving equality and sorted order. If `a < b`, the assigned index
for `a` must be smaller than the index for `b`; equal values must
receive the same index.

###### How it works

Race times such as `120`, `450`, and `900` may be far apart, but their
useful order is first, second, third. **Coordinate compression**
replaces each distinct value with this ordered rank. Equal original
values receive the same rank. The lookup contract is
`original value -> position in the sorted distinct values`. For
`[-5,20,100]`, `rankByValue[20]=1` means `20` is the second distinct
value, at zero-based position `1`.

###### Step-by-step examples

For `[100,-5,100,20]`, expect ranks `[2,0,2,1]`.

Initial state: `values=[100,-5,100,20]`, `distinct=[]`,
`rankByValue={}`, `result=[]`.

1.  Sort distinct values
    - Decision: Sort a copy and keep one copy of each value.
    - After: `values=[100,-5,100,20]`, `distinct=[-5,20,100]`,
      `rankByValue={}`, `result=[]`.
    - Known correct: Every different input value appears once in
      increasing order.
2.  Assign ranks
    - Decision: Map `-5 -> 0`, `20 -> 1`, and `100 -> 2`.
    - After: `values=[100,-5,100,20]`, `distinct=[-5,20,100]`,
      `rankByValue={-5:0,20:1,100:2}`, `result=[]`.
    - Known correct: Smaller original values have smaller ranks.
3.  Translate the original order
    - Decision: Replace each value with its mapped rank.
    - After: `values=[100,-5,100,20]`, `distinct=[-5,20,100]`,
      `rankByValue={-5:0,20:1,100:2}`, `result=[2,0,2,1]`.
    - Known correct: Equal values share rank `2`, and output
      positions still correspond to the original input positions.

For `[-1000000000,1000000000]`, compression returns `[0,1]`. Direct
indexing across the raw inclusive range would reserve `2000000001`
positions to store facts about only two values.

For an empty input, both the distinct list and output are empty. Equal
values must receive one rank rather than separate positions.

###### Why it works

The sorted distinct list contains every value once. If `a < b`, `a`
appears before `b`, so `rank[a] < rank[b]`; equal values use the same
map entry. Thus comparisons are preserved. Translation advances
through the finite input, so it finishes.

###### Go example

```go
import "sort"

func compress(values []int) []int {
	// Sorting a copy groups duplicates and preserves the input order for
	// output.
	ordered := append([]int(nil), values...)
	sort.Ints(ordered)
	// Assign the next rank only on a value's first sorted occurrence.
	rankByValue := make(map[int]int, len(ordered))
	for _, value := range ordered {
		if _, found := rankByValue[value]; !found {
			rankByValue[value] = len(rankByValue)
		}
	}
	// Translate the original sequence through the completed
	// value-to-rank map.
	result := make([]int, len(values))
	for valueIndex, value := range values {
		result[valueIndex] = rankByValue[value]
	}
	return result
}
```

###### Work and memory

Sorting `n` copied values takes `O(n log n)` time; mapping and
translation take expected `O(n)` time. The sorted copy and map use
`O(n)` auxiliary memory. The returned ranks use `O(n)` output space.

###### When to use it

Use it when only value order matters and a Fenwick tree, segment tree,
or DP transition needs compact indices for large or negative values.

###### When not to use it

Do not use ranks when actual numeric gaps matter, such as physical
distance `100-20`; ranks preserve order, not magnitude.

###### Common mistakes and edge cases

> **Warning:** Compression does not reorder the input. Processing the
> sorted copy instead of the original sequence breaks order-dependent
> DP.

Include every query endpoint that must receive a rank, deduplicate
equal values, and do not confuse a value rank with an original input
position.

###### Related material

- [C.1.3.10 Dynamic programming](#c1310-dynamic-programming) uses
  compressed ranks to speed value-range transitions.
- [C.1.3.17 Fenwick and segment
  trees](#c1317-fenwick-and-segment-trees) query and update those
  ranks.
- [C.1.3.5 Sorting and
  selection](#c135-sorting-and-selection) explains the sort.

<a id="c1311-divide-and-conquer"></a>

##### C.1.3.11 Split and combine (divide and conquer)

###### The problem it solves

Given a large input that can be split into smaller independent inputs
of the same kind, compute each part and combine the part results into
the whole answer. Splitting is useful only when the combination step
is correct and does not redo so much work that it erases the benefit
of solving smaller parts.

###### How it works

Give two people separate halves of a stack of papers, let each sort a
half, then merge the results. **Divide and conquer** splits
independent work, solves each part by the same rule, and combines the
answers.

###### Step-by-step examples

Sort `[4,1,3,2]`; expect `[1,2,3,4]`.

Initial state: `pending=[[4,1,3,2]]`, `solved=[]`,
`result=pending`.

1.  Split the full input
    - Decision: Divide it into `[4,1]` and `[3,2]`.
    - After: `pending=[[4,1],[3,2]]`, `solved=[]`,
      `result=pending`.
    - Known correct: The halves do not overlap. Together, they
      contain every original value.
2.  Solve the left half
    - Decision: Merge sorted one-value lists `[4]`
      and `[1]` by taking the smaller front value first.
    - After: `pending=[[3,2]]`, `solved=[[1,4]]`,
      `result=pending`.
    - Known correct: Both left-half values appear in order.
3.  Solve the right half
    - Decision: Merge sorted one-value lists `[3]` and `[2]`.
    - After: `pending=[]`, `solved=[[1,4],[2,3]]`,
      `result=pending`.
    - Known correct: Both right-half values appear in order.
4.  Combine the two halves
    - Decision: Repeatedly take the smaller unread front value:
      `1`, then `2`, then `3`, then `4`.
    - After: `pending=[]`, `solved=[[1,4],[2,3]]`,
      `result=[1,2,3,4]`.
    - Known correct: The result is sorted and contains every input
      value exactly once.

###### Why it works

For merge sort, assume each smaller call returns its half in sorted
order. During the merge, the output prefix is sorted and contains
exactly the values already removed from the two halves. Their smaller
front value must be the smallest value not yet emitted, so appending
it preserves that fact. When one half empties, the other half is
already sorted and every one of its values is at least the output's
last value. The base case handles fewer than two values, and each
split makes the ranges smaller, so recursion finishes with a sorted
permutation of the input.

Divide and conquer is most natural when subproblems do not overlap.
Dynamic programming is needed when different branches repeatedly ask
for the same state. Memoizing a divide-and-conquer recursion because
its parts overlap has effectively turned it into DP.

###### Go example

Calling `mergeSort([]int{4, 1, 3, 2})` returns
`[]int{1, 2, 3, 4}`. The trace's `pending` work is represented by
recursive calls on the two halves, while its `solved` halves are the
code's returned `left` and `right` slices. Both representations apply
the same left-half, right-half, then merge phases.

```go
func mergeSort(values []int) []int {
	// A zero- or one-element slice is already sorted; return a copy so
	// every call owns its result independently of the input.
	if len(values) < 2 {
		return append([]int(nil), values...)
	}
	// Recursively sort two disjoint halves.
	middle := len(values) / 2
	left, right := mergeSort(values[:middle]), mergeSort(values[middle:])
	// Repeatedly take the smaller front value. Choosing left on equality
	// keeps equal input values in their original relative order.
	result := make([]int, 0, len(values))
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			// The left front is the smallest remaining value; append and
			// consume it.
			result, left = append(result, left[0]), left[1:]
		} else {
			// Otherwise the right front is smaller and must come next.
			result, right = append(result, right[0]), right[1:]
		}
	}
	// At most one side remains; its existing order is already sorted.
	result = append(result, left...)
	return append(result, right...)
}
```

The copied base case prevents the result from aliasing a caller-owned
input.

###### Work and memory

The cost depends on the number and sizes of the smaller parts plus the
work needed to combine them. For `n` values, merge sort has `log n`
split levels, and merging all ranges on one level handles `n` values.
It therefore takes `O(n log n)` time. The active child results use
`O(n)` auxiliary memory, the call stack uses `O(log n)` memory, and
the returned sorted slice uses `O(n)` output memory.

###### When to use it

A balanced split makes the input substantially smaller, the parts can
be solved without sharing evolving state, and their results have a
clear merge step. Merge sort, tree reconstruction, fast
exponentiation, and some geometric algorithms use this shape.

###### When not to use it

Splits can remain badly unbalanced, the parts repeatedly solve the
same smaller questions, or combining their answers costs as much as
solving the original problem directly.

###### Common mistakes and edge cases

> **Warning:** A split that does not make both recursive ranges
> smaller loops forever. If `[left,right)` is split at `middle`,
> recurse on `[left,middle)` and `[middle,right)`, not the original
> range again.

An empty or one-value input is the base case. During merge, copy the
remainder of either half after the other empties. When equal values
must remain in their original relative order, called stable order,
declare and preserve a tie rule.

###### Related material

- [C.1.3.10 Dynamic programming](#c1310-dynamic-programming) handles
  overlapping rather than independent subproblems.
- [C.1.3.9 Two pointers](#c139-two-pointers) explains sorted merging.
- [A.3 Recursion and saved work](#a3-recursion-and-saved-work) counts
  call-stack memory.

<a id="c1312-kadanes-algorithm-best-subarray-scan"></a>

##### C.1.3.12 Keep the best section ending here (Kadane's algorithm)

###### The problem it solves

Given positive and negative numbers, return the largest sum of any
non-empty consecutive section, also called a **subarray**. Trying
every pair of start and end boundaries repeats sums; a one-pass method
must decide whether the best section ending at the current value
should extend the previous section or start fresh.

###### How it works

Track the best daily-profit streak that must end today. Either extend
yesterday's streak or start again today, then remember the best streak
ever seen. This is **Kadane's algorithm**.

###### Step-by-step examples

For `[-2,3,-1,4,-5]`, expect sum `6` from `[3,-1,4]`.

Initial state: `index=0`, `endingHere=unset`, `bestOverall=unset`.

1.  Process `-2`
    - Decision: Initialize both sums from the first value.
    - After: `index=1`, `endingHere=-2`, `bestOverall=-2`.
    - Known correct: `[-2]` is the only available section.
2.  Process `3`
    - Decision: Starting fresh at `3` beats extending to `1`.
    - After: `index=2`, `endingHere=3`, `bestOverall=3`.
    - Known correct: `[3]` is best at this ending position and
      among all positions processed so far.
3.  Process `-1`
    - Decision: Extending gives `2`, which beats starting at `-1`.
    - After: `index=3`, `endingHere=2`, `bestOverall=3`.
    - Known correct: `[3,-1]` is best ending here, while `[3]`
      remains best overall.
4.  Process `4`
    - Decision: Extending gives `6`; update the overall best.
    - After: `index=4`, `endingHere=6`, `bestOverall=6`.
    - Known correct: `[3,-1,4]` is best ending here and best
      overall.
5.  Process `-5`
    - Decision: Extending gives `1`; keep the previous overall
      best.
    - After: `index=done`, `endingHere=1`, `bestOverall=6`.
    - Known correct: Every possible ending position has been
      represented, and the largest sum is `6`.

For `[-4,-2]`, initializing from the first value returns `-2`;
initializing to zero would wrongly allow an empty section.

###### Why it works

Every non-empty subarray ending at the current position is either the
current value alone or a subarray ending at the previous position
followed by the current value. Among extensions, only the best
previous ending can matter. `bestOverall` records the best ending
among every position processed so far. Both facts are true at the
first value. Choosing the better of starting and extending preserves
the first fact, and taking the larger overall sum preserves the
second. The loop advances once per remaining value, so it finishes.
At the end, every possible ending position has contributed its best
sum, making `bestOverall` the requested maximum.

###### Go example

```go
func maxSubarray(values []int) (int, bool) {
	// An empty slice has no nonempty subarray to score.
	if len(values) == 0 {
		return 0, false
	}
	// endingHere is the best sum ending at the current position;
	// bestOverall is the best sum ending anywhere in the processed
	// prefix.
	endingHere, bestOverall := values[0], values[0]
	for _, value := range values[1:] {
		// Either extend the previous ending range or start fresh at this
		// value.
		if endingHere+value > value {
			endingHere += value
		} else {
			endingHere = value
		}
		if endingHere > bestOverall {
			// A newly computed ending range improves the best answer seen
			// anywhere.
			bestOverall = endingHere
		}
	}
	return bestOverall, true
}
```

The boolean gives empty input an explicit result instead of inventing
an empty subarray sum.

###### Work and memory

For `n` values, the scan takes `O(n)` time and `O(1)` extra memory.

###### When to use it

Find the largest sum of one non-empty consecutive section without
trying every pair of boundaries.

###### When not to use it

The answer may skip values, is not additive, or deciding whether to
extend needs more information than the best section ending at the
previous position.

###### Common mistakes and edge cases

> **Warning:** Initializing both sums to zero permits an empty
> section. For `[-4,-2]`, that mistake returns `0` instead of the
> required nonempty answer `-2`.

To return indices, remember the candidate start whenever beginning at
the current value is better, and save it when `bestOverall` improves.
Reject or separately define empty input, and check addition overflow
when `endingHere+value` can exceed `int`.

###### Related material

- [C.1.3.14 Sliding window](#c1314-sliding-window) handles ranges
  whose validity can be repaired by moving boundaries.
- [A.2 Correctness and invariants](#a2-correctness-and-invariants)
  explains why the best section ending at each position stays correct.

<a id="c1313-stack"></a> <a id="stack"></a>

##### C.1.3.13 Finish the newest open item first (stack)

###### The problem it solves

Given nested delimiters, function calls, or other unfinished items,
process each closing or completion event against the most recently
opened item that is still pending. Older pending items must remain
available underneath it, because they cannot finish until all work
opened after them has finished.

###### How it works

Stack trays in a cafeteria: the last tray placed on top is the first
one removed. A **stack** uses this same rule, called **last in, first
out (LIFO)**.

###### Step-by-step examples

Validate the four-character input `([])`; expect `true`. In the stack
shown below, the rightmost item is the top.

Initial state: `index=0`, `stack=[]`, `result=pending`.

1.  Read `(`
    - Decision: It is an opener, so push it.
    - After: `index=1`, `stack=['(']`, `result=pending`.
    - Known correct: The stack contains the only opener not yet
      matched.
2.  Read `[`
    - Decision: It is an opener, so push it.
    - After: `index=2`, `stack=['(', '[']`, `result=pending`.
    - Known correct: The newest unmatched opener, `[`, is on top.
3.  Read `]`
    - Decision: It matches top `[`, so pop the top.
    - After: `index=3`, `stack=['(']`, `result=pending`.
    - Known correct: The inner pair is complete; `(` is still
      unmatched.
4.  Read `)`
    - Decision: It matches top `(`, so pop the top.
    - After: `index=4`, `stack=[]`, `result=true`.
    - Known correct: Every closer matched the newest unmatched
      opener, and no opener remains.

For `([)]`, `)` does not match top `[`, so reject immediately. Empty
input is valid only if the problem defines an empty sequence as
balanced.

###### Why it works

The stack contains exactly the opened items not yet completed. It
starts empty because no input has been read. Pushing an opener adds
the newest unfinished item. A valid closer must match the top because
an older opener cannot close through a still-open inner item. Popping
the match preserves the fact. Each input item is read once, so the
scan finishes. An empty stack at the end proves that no opener
remains.

###### Go example

Calling `balancedBrackets("([])")` follows the trace and returns
`true`. The byte slice's final element is the stack top.

```go
func balancedBrackets(input string) bool {
	stack := make([]byte, 0, len(input))
	for index := range input {
		current := input[index]
		switch current {
		case '(', '[', '{':
			// Openers become the newest unfinished item.
			stack = append(stack, current)
		case ')', ']', '}':
			// A closer requires a pending opener before the stack top can
			// be inspected.
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			matches := top == '(' && current == ')' ||
				top == '[' && current == ']' ||
				top == '{' && current == '}'
			if !matches {
				return false
			}
			stack = stack[:len(stack)-1]
		default:
			// This validator's language contains bracket bytes only.
			return false
		}
	}
	return len(stack) == 0
}
```

[E.5 Generic stack](#e5-generic-stack) provides a reusable container
for other last-in, first-out tasks.

###### Work and memory

Push, pop, and top inspection take `O(1)` time, aside from occasional
slice growth when Go uses a slice as the stack. For `n` input items,
the stack may retain `O(n)` items. Clear a removed slice position when
it may hold pointers so Go's garbage collector can reclaim objects no
longer in use.

###### When to use it

The most recently opened or unfinished item must be handled first, as
with nested brackets, undo history, expression evaluation, or a
depth-first search implemented with a loop.

###### When not to use it

The oldest item must be processed first or the algorithm needs
efficient access to arbitrary stored items.

###### Common mistakes and edge cases

Use a queue when the oldest waiting item must be handled first.
Removing from the front of a Go slice does not model a stack and may
retain a large underlying array. Recursive calls already use memory
like a stack behind the scenes; use your own stack when nesting may be
unsafe or direct control is needed.

###### Related material

- [E.5 Generic stack](#e5-generic-stack) provides the Go
  implementation.
- [C.1.3.20 Depth-first search](#c1320-depth-first-search) can replace
  recursive calls with an explicit stack.
- [Valid Parentheses](
  problems/solved_easy_leetcode_0020_valid_parentheses.go
  )
  is the representative problem.

<a id="sliding-window"></a>

<a id="c1314-sliding-window"></a>

##### C.1.3.14 Move one consecutive range (sliding window)

###### The problem it solves

Given a sequence, return a longest, shortest, or otherwise best
consecutive range whose validity can be updated when either boundary
moves. Instead of recomputing every candidate range, add the new right
value, remove values from the left when needed, and maintain only the
summary required to test the current range.

###### How it works

Put a movable fence around one consecutive portion of the input. The
fence always contains neighboring items with no gaps. Extend the right
side to consider new values and move the left side when earlier values
must leave. The portion inside the fence is called a **window**;
moving both ends forward is the **sliding-window** method.

###### Step-by-step examples

Find the shortest positive-value range totaling at least `7` in
`[2,3,1,2,4,3]`; expect length `2`.

Initial state: `left=0`, `right=-1`, `sum=0`, `best=none`.

1.  Grow through index `3`
    - Decision: Add `2`, `3`, `1`, then `2`; running sums are `2`,
      `5`, `6`, and `8`. Record length `4` at the target.
    - After: `left=0`, `right=3`, `sum=8`, `best=4`.
    - Known correct: No window ending before index `3` reaches
      `7`, and a valid window ending at index `3` has been found.
2.  Shrink at right index `3`
    - Decision: Remove the leftmost `2`; sum `6` is invalid, so
      stop shrinking.
    - After: `left=1`, `right=3`, `sum=6`, `best=4`.
    - Known correct: Every valid start for right index `3` was
      tested; removing another positive value cannot restore `7`.
3.  Add index `4` and shrink
    - Decision: Add `4`, giving sum `10`. Record length `4`,
      remove `3`, record the shorter length `3` at sum `7`, then
      remove `1`. Sum `6` is invalid, so stop shrinking.
    - After: `left=3`, `right=4`, `sum=6`, `best=3`.
    - Known correct: Every valid start for right index `4` was
      tested, and length `3` is the shortest seen so far.
4.  Add index `5` and finish
    - Decision: Add `3`, giving sum `9`. Record length `3`, remove
      `2`, and record length `2` at sum `7`. Remove `4`; sum `3` is
      invalid, so stop.
    - After: `left=5`, `right=5`, `sum=3`, `best=2`.
    - Known correct: Every ending position has been handled, so
      return `2` for window `[4,3]`.

Focused failure case: for values `[1,-1,5]` and target `5`, the
expected shortest length is `1`, but this positive-value method would
record length `3`, remove `1`, see sum `4`, and stop. It would never
remove `-1`, even though doing so raises the sum to `5`. Negative
values therefore require another method.

###### Why it works

Before each new `right`, every earlier ending position has already
been handled. After adding `values[right]`, positive values make the
sum grow. While the sum meets the target, recording the length and
removing the left value tests every valid start for this `right` from
longest to shortest; once the sum falls below the target, removing
more positive values cannot restore validity. Both boundaries move
only forward, so the scan finishes after considering the shortest
valid range ending at every position.

###### Go example

This form requires positive values, a positive target, and running
sums that fit in `int`; the positivity makes left-boundary removal
safe.

```go
func shortestPositiveWindow(values []int, target int) int {
	// best starts as an impossible length so it also records whether any
	// valid window has been found.
	left, sum, best := 0, 0, len(values)+1
	for right, value := range values {
		// Expanding right can only increase sum because all values are
		// positive.
		sum += value
		// While valid, record the window and remove its leftmost value to
		// find the shortest valid window ending at right.
		for sum >= target {
			if right-left+1 < best {
				best = right - left + 1
			}
			sum -= values[left]
			left++
		}
	}
	if best == len(values)+1 {
		return 0
	}
	return best
}
```

###### Work and memory

For `n` input values, each end moves forward at most `n` times, so the
usual window scan takes `O(n)` time. A running sum uses `O(1)` extra
memory; character or value counts use memory proportional to the
number of distinct values kept.

###### When to use it

The prompt asks for a longest, shortest, largest, or smallest
consecutive portion of an array or string. You can update the tracked
sum, character counts, or other information when one item enters or
leaves.

###### When not to use it

The answer may skip items, or adding an item can make the condition
switch back and forth unpredictably. Mixed positive and negative
values often cause that problem for sum-based windows.

###### Common mistakes and edge cases

> **Warning:** The inner loop is unsafe for `target <= 0`: with
> `values=[1]`, it remains true after `left` passes the slice and then
> indexes outside it. This implementation requires a positive target.

Empty input returns `0`. A single value at least as large as the
target gives length `1`; mixed negative values require another method
because shrinking no longer changes the sum predictably.

###### Related material

- [C.1.3.15 Prefix sums](#c1315-prefix-sums) handles fixed range
  queries without a safe moving boundary.
- [C.1.3.29 Monotonic deque](#c1329-monotonic-deque) maintains a
  window extreme.

<a id="c1315-prefix-sums"></a>

##### C.1.3.15 Subtract saved earlier totals (prefix sums)

###### The problem it solves

Given an unchanged sequence and many range-sum or range-count queries,
return each range result without visiting every value in that range.
Precompute the total before every boundary so a query can subtract the
total before its left edge from the total through its right edge.

###### How it works

A bank statement can show the running balance after every transaction.
Subtracting two saved balances gives the change between them without
adding each transaction again. The list of saved running totals is a
**prefix-sum array**. Its contract is
`prefix[endExclusive] -> sum of values before endExclusive`. For
example, `prefix[3]=7` for `[2,4,1,3]` means indexes `0` through `2`
sum to `7`.

###### Step-by-step examples

For values `[2,4,1,3]`, answer two overlapping inclusive queries:
indexes `1..3` should total `8`, and indexes `2..3` should total `4`.
Adding each range directly would add the values at indexes `2` and `3`
twice. The saved prefix totals should answer both without revisiting
either range.

Initial state: `valueIndex=0`, `prefix=[0]`, `answers=[]`.

1.  Read value `2`
    - Decision: Add it to the latest total `0`.
    - After: `valueIndex=1`, `prefix=[0,2]`, `answers=[]`.
    - Known correct: `prefix[1]` totals the first value.
2.  Read value `4`
    - Decision: Add it to the latest total `2`.
    - After: `valueIndex=2`, `prefix=[0,2,6]`, `answers=[]`.
    - Known correct: `prefix[2]` totals the first two values.
3.  Read value `1`
    - Decision: Add it to the latest total `6`.
    - After: `valueIndex=3`, `prefix=[0,2,6,7]`, `answers=[]`.
    - Known correct: `prefix[3]` totals the first three values.
4.  Read value `3`
    - Decision: Add it to the latest total `7`.
    - After: `valueIndex=done`, `prefix=[0,2,6,7,10]`,
      `answers=[]`.
    - Known correct: `prefix[4]` totals all four values.
5.  Answer query `1..3`
    - Decision: Subtract `prefix[1]=2` from `prefix[4]=10`.
    - After: `valueIndex=done`, `prefix=[0,2,6,7,10]`,
      `answers=[8]`.
    - Known correct: Only index `0` was removed, leaving indexes
      `1`, `2`, and `3`.
6.  Answer query `2..3`
    - Decision: Subtract `prefix[2]=6` from `prefix[4]=10`.
    - After: `valueIndex=done`, `prefix=[0,2,6,7,10]`,
      `answers=[8,4]`.
    - Known correct: Indexes `2` and `3` were summed from saved
      totals; their values were not visited again.

The leading zero represents the total before any values and makes a
range beginning at index `0` use the same subtraction rule. The
function below treats `left > right` as invalid rather than defining
an empty inclusive range.

###### Why it works

`prefix[0]=0` is the sum before the first value. If `prefix[i]` is the
sum before input index `i`, adding `values[i]` makes
`prefix[i+1]` the sum before the next index, so the meaning is
preserved through the finite build loop. For a query, both saved
totals start at index `0`. Subtracting the total before `left` removes
exactly the values before the requested range and leaves each value
from `left` through `right` once.

###### Go example

The caller must provide `0 <= left <= right` and
`right+1 < len(prefixTotals)`.
The trace's `prefix` is the code's `prefixTotals`. Calling
`prefixSums([]int{2, 4, 1, 3})` returns `[0,2,6,7,10]`; passing that
result to `rangeSum` with `(left,right)=(1,3)` and `(2,3)` returns
`8` and `4`.

```go
func prefixSums(values []int) []int {
	// prefixTotals has an extra leading zero so each position stores the
	// sum of the values before the corresponding input index.
	prefixTotals := make([]int, len(values)+1)
	for valueIndex, value := range values {
		// Extend the total through the current value and store it one
		// position to the right.
		prefixTotals[valueIndex+1] = prefixTotals[valueIndex] + value
	}
	return prefixTotals
}

func rangeSum(prefixTotals []int, left, right int) int {
	// Subtract the sum before left from the sum through right;
	// everything outside the inclusive range cancels.
	return prefixTotals[right+1] - prefixTotals[left]
}
```

###### Work and memory

Building `n+1` totals for `n` values takes `O(n)` time and `O(n)`
stored-summary memory. Each compatible sum or count query then takes
`O(1)` time.

###### When to use it

Many questions ask for sums or counts over overlapping ranges. The
simplest solution repeatedly adds values it has already added for an
earlier question. To count values with one property, store `1` when
an item has that property and `0` otherwise; the same prefix sums then
answer range-count queries.

###### When not to use it

The requested property cannot be recovered by combining two running
summaries. For example, you cannot find a range's smallest value by
subtracting two previously stored smallest values; use a monotonic
deque for a moving range or a range-query structure for updates and
arbitrary queries.

###### Common mistakes and edge cases

> **Warning:** Omitting `prefix[0]=0` shifts every boundary. For
> `values=[2,4]`, the range `0..0` must use `prefix[1]-prefix[0]=2`,
> not read a position before the slice.

Reject `left > right`, negative indexes, and
`right+1 >= len(prefixTotals)`. Use a wider sum type or checked
addition when valid inputs can overflow `int`.

###### Related material

- [C.1.3.16 Difference arrays](#c1316-difference-arrays) reverses this
  idea for collected range updates.
- [C.1.3.17 Fenwick and segment
  trees](#c1317-fenwick-and-segment-trees) supports updates between
  queries.

<a id="c1316-difference-arrays"></a>

##### C.1.3.16 Mark range changes (difference arrays)

###### The problem it solves

Given an array of initially zero values and many inclusive range
additions, return the final value at every position. When only that
final result is needed, record where each addition starts and stops,
then rebuild all values once instead of visiting every covered
position for every update.

###### How it works

Iterating through every covered index for every update is correct. If
many updates cover long ranges, however, the same positions change
again and again. On a staffing calendar, mark `+2` when two workers
start and `-2` when they finish instead of rewriting every hour they
cover. A **difference array** stores those start and stop changes; one
later running total rebuilds every hour's staffing. Its contract is
`difference[index] -> net change that begins at index`. For example,
`difference[2]=4` means the running total gains `4` before result
index `2` is stored.

###### Step-by-step examples

Start with eight zeroes. Apply `+3` to indexes `0..7`, `+2` to indexes
`1..7`, `-1` to indexes `0..6`, and `+4` to indexes `2..7`, including
both endpoints. Expect `[2,4,8,8,8,8,8,9]`.

Baseline: iterate through every covered index.

Initial state: `result=[0,0,0,0,0,0,0,0]`, `writes=0`.

1.  Apply all four updates directly
    - Decision: Visit `8+7+7+6=28` covered positions and add each
      update's delta at every position it covers.
    - After: `result=[2,4,8,8,8,8,8,9]`, `writes=28`.
    - Known correct: Direct iteration gives the required result,
      but overlapping positions were rewritten up to four times.

The difference-array method performs eight boundary additions and
eight rebuilding steps instead of 28 range additions on this input.

Difference-array method:

Initial state: `nextUpdate=+3@0..7`,
`difference=[0,0,0,0,0,0,0,0,0]`, `result=[]`,
`runningTotal=0`.

1.  Record `+3` on `0..7`
    - Decision: Add `3` at index `0` and subtract `3` at index
      `8`.
    - After: `nextUpdate=+2@1..7`,
      `difference=[3,0,0,0,0,0,0,0,-3]`, `result=[]`,
      `runningTotal=0`.
    - Known correct: Every result position will receive `+3`.
2.  Record `+2` on `1..7`
    - Decision: Add `2` at index `1` and subtract `2` at index
      `8`.
    - After: `nextUpdate=-1@0..6`,
      `difference=[3,2,0,0,0,0,0,0,-5]`, `result=[]`,
      `runningTotal=0`.
    - Known correct: Index `0` receives only `+3`; later indexes
      also receive `+2`.
3.  Record `-1` on `0..6`
    - Decision: Add `-1` at index `0` and add `1` at index `7` to
      stop that negative delta.
    - After: `nextUpdate=+4@2..7`,
      `difference=[2,2,0,0,0,0,0,1,-5]`, `result=[]`,
      `runningTotal=0`.
    - Known correct: The first seven positions receive `-1`, but
      index `7` does not.
4.  Record `+4` on `2..7`
    - Decision: Add `4` at index `2` and subtract `4` at index
      `8`.
    - After: `nextUpdate=none`,
      `difference=[2,2,4,0,0,0,0,1,-9]`, `result=[]`,
      `runningTotal=0`.
    - Known correct: Indexes `2` through `7` also receive `+4`.
5.  Rebuild indexes `0` through `2`
    - Decision: Add markers `2`, `2`, and `4` in order, storing
      each new total.
    - After: `nextUpdate=none`,
      `difference=[2,2,4,0,0,0,0,1,-9]`, `result=[2,4,8]`,
      `runningTotal=8`.
    - Known correct: Index `2` contains `3+2-1+4=8`, so all four
      overlapping updates contribute exactly once.
6.  Finish indexes `3` through `7`
    - Decision: Store `8` through index `6`, then add marker `1`
      and store `9` at index `7`.
    - After: `nextUpdate=none`,
      `difference=[2,2,4,0,0,0,0,1,-9]`,
      `result=[2,4,8,8,8,8,8,9]`, `runningTotal=9`.
    - Known correct: Every result position contains the sum of all
      updates whose ranges include it.

The extra final position allows an update to end at the last result
index without writing outside the array.

###### Why it works

The running change increases when an update begins and decreases
immediately after that update ends. Every position between those
points receives the update exactly once. Recording consumes one
finite update at a time, and rebuilding advances through each result
position once, so both phases finish.

###### Go example

The function requires `length >= 0` and every update to satisfy
`0 <= left <= right < length`.
For the trace, call `applyRangeAdds` with `length=8` and updates
`{0,7,3}`, `{1,7,2}`, `{0,6,-1}`, and `{2,7,4}`. The returned slice
is `[2,4,8,8,8,8,8,9]`; each triple supplies `left`, `right`, and
`delta` to one `rangeAdd`.

```go
type rangeAdd struct{ left, right, delta int }

func applyRangeAdds(length int, updates []rangeAdd) []int {
	// difference records where each delta starts and where it stops
	// affecting later positions. The extra slot safely holds a stop at
	// index length.
	difference := make([]int, length+1)
	for _, update := range updates {
		difference[update.left] += update.delta
		difference[update.right+1] -= update.delta
	}
	// A prefix sum reconstructs the total active delta at every output
	// position.
	result, runningTotal := make([]int, length), 0
	for resultIndex := range result {
		runningTotal += difference[resultIndex]
		result[resultIndex] = runningTotal
	}
	return result
}
```

###### Work and memory

Let `k` be the total number of positions covered across all updates.
Direct iteration takes `O(k)` time, which can become `O(u*n)` when `u`
updates cover most of an `n`-position result. It uses no range-marker
array. Recording boundaries and rebuilding takes `O(u+n)` time and
`O(n)` extra memory. For a few short updates, direct iteration is
often simpler and may be faster.

###### When to use it

The problem applies many additions or subtractions to ranges and asks
for the final array after all updates.

###### When not to use it

There are only a few short updates, or you must inspect the complete
array after every individual update rather than after all updates have
been collected. Direct iteration is simpler in those cases.

###### Common mistakes and edge cases

> **Warning:** Allocating only `length` difference positions makes an
> update ending at `length-1` write `difference[length]` out of
> bounds. Reserve one extra position at the end.

Require `length >= 0` and validate every update before writing either
marker. An empty result accepts no nonempty update, and accumulated
deltas must fit the chosen integer type.

###### Related material

- [C.1.3.15 Prefix sums](#c1315-prefix-sums) explains the accumulating
  rebuild.
- [C.1.3.17 Fenwick and segment
  trees](#c1317-fenwick-and-segment-trees) handles interleaved updates
  and queries.

<a id="c1317-fenwick-and-segment-trees"></a>

##### C.1.3.17 Update range sums (Fenwick and segment trees)

###### The problem it solves

An array changes between questions. Two operations may alternate:

1. set one array value; and
2. ask for the sum from index `left` through index `right`, including
   both endpoints.

Scanning that range answers one question but repeats work on later
questions. A prefix-sum array answers quickly, but changing one value
makes every later prefix sum stale. Segment trees and Fenwick trees
save smaller block sums, so one change repairs only a few saved sums.

###### How it works

Both structures follow one rule:

> Every saved number is the sum of one named block of array values.

A query combines saved blocks, or subtracts two combined prefixes, to
cover the requested range exactly. An update repairs only blocks
containing the changed index.

**Segment tree: show every block.**

A segment tree makes the blocks visible. For `[2,1,5,3]`, its leaves
hold individual values. Each parent holds the sum of its two children:

```text
                 [0..3] = 11
                /           \
          [0..1] = 3      [2..3] = 8
           /     \          /     \
       [0] = 2 [1] = 1  [2] = 5 [3] = 3
```

To set index `2` to `6`, change its leaf and rebuild the ancestors on
the path to the root. To sum `0..1`, read the saved block `[0..1]`.
To sum `1..2`, combine the two leaves `[1]` and `[2]`.

**Fenwick tree: keep only selected sum blocks.**

A Fenwick tree is a compact choice for sums. Public array indexes are
zero-based. Fenwick positions are one-based, and position `0` is
unused. Public index `0` maps to Fenwick position `1`, index `1` maps
to position `2`, and so on. For `[2,1,5,3]`, the complete stored
meaning is:

```text
Fenwick   Binary   Public indexes in block   Saved sum
position
1         001      0                          2
2         010      0..1                       3
3         011      2                          5
4         100      0..3                      11
```

For a positive Fenwick position `p`, `p & -p` returns its block size.
The expression isolates the lowest-value `1` bit. For example, using
eight displayed bits for position `12`:

```text
p          00001100
-p         11110100
p & -p     00000100 = 4
```

Go represents a negative integer so that `-p` is formed by flipping
the bits of `p` and adding one. The addition leaves the lowest-value
`1` in the same position, while `&` clears the other positions. Thus
position `3` (`011`) has block size `1`, while position `4` (`100`)
has block size `4`. Position `p` stores the block ending at public
index `p-1` with that size.

A prefix sum is the sum from public index `0` up to, but not
including, an `end` index. Start at Fenwick position `end`. After
reading one block, subtract its size to move just before that block.
For example, `prefix(3)` reads positions `3` and `2`. Their blocks
are public indexes `2` and `0..1`, which together cover `0..2`.

The inclusive sum `left..right` is:

```text
prefix(right + 1) - prefix(left)
```

Setting a value first computes the change:

```text
delta = new value - old value
```

Starting at Fenwick position `index+1`, add `delta` to that block.
Then add the block size to reach the next stored block containing the
same public index. Stop after passing the final Fenwick position.

For this four-value example, every possible update path is:

```text
public index   start position   repaired Fenwick positions
0              1                1, 2, 4
1              2                2, 4
2              3                3, 4
3              4                4
```

Each listed position's named block contains the public index in that
row. No unlisted stored block contains it.

###### Step-by-step examples

Use `[2,1,5,3]`. First ask for the sum of public indexes `1..2`.
Then set index `2` from `5` to `6`.

Baseline: scan or rebuild.

Initial state: `values=[2,1,5,3]`, `request=1..2`,
`answer=pending`.

1.  Answer by scanning
    - Decision: Read both requested values and add `1+5`.
    - After: `values=[2,1,5,3]`, `request=1..2`, `answer=6`.
    - Known correct: Every value in `1..2` was included once, but
      a long range would require many reads for every question.

Initial state: `values=[2,1,5,3]`, `prefix=[0,2,3,8,11]`,
`update=set(2,6)`.

1.  Repair saved prefix sums
    - Decision: Setting index `2` from `5` to `6` changes every
      prefix sum after that value.
    - After: `values=[2,1,6,3]`, `prefix=[0,2,3,9,12]`,
      `update=done`.
    - Known correct: Prefix sums make the query cheap but can make
      one update repair many positions.

Segment tree: visible blocks.

Initial state: `blocks={[0..3]:11,[0..1]:3,[2..3]:8,
[0]:2,[1]:1,[2]:5,[3]:3}`, `total=0`, `update=pending`.

1.  Query indexes `1..2`
    - Decision: The smallest exact cover is `[1]` and `[2]`; add
      their saved sums: `1+5`.
    - After: `blocks` unchanged, `total=6`, `update=pending`.
    - Known correct: The chosen blocks do not overlap, and their
      indexes together are exactly `1..2`.
2.  Set index `2` to `6`
    - Decision: Set leaf `[2]` to `6`, then recompute only its
      parent and the root.
    - After: `blocks={[0..3]:12,[0..1]:3,[2..3]:9,
      [0]:2,[1]:1,[2]:6,[3]:3}`, `total=6`, `update=done`.
    - Known correct: Those are exactly the saved blocks containing
      index `2`.

Fenwick tree: selected blocks.

Build the initial Fenwick state from `values=[2,1,5,3]` by storing
the sum of each position's block:

```text
sums[0] = 0                   unused position
sums[1] = values[0]          = 2
sums[2] = values[0] + values[1]
        = 2 + 1              = 3
sums[3] = values[2]          = 5
sums[4] = values[0] + values[1] + values[2] + values[3]
        = 2 + 1 + 5 + 3      = 11
```

Therefore the initial stored state is `sums=[0,2,3,5,11]`. It is not
the prefix-sum array: position `3` stores only its one-value block,
not the sum of public indexes `0..2`.

Initial state: `values=[2,1,5,3]`, `sums=[0,2,3,5,11]`,
`position=3`, `total=0`, `rangeAnswer=pending`.

1.  Begin `prefix(3)`
    - Decision: Position `3` has block size `1` and stores public
      index `2`. Add `5`, then move to `3-1=2`.
    - After: `values=[2,1,5,3]`, `sums=[0,2,3,5,11]`,
      `position=2`, `total=5`, `rangeAnswer=pending`.
    - Known correct: Public index `2` has been counted once.
2.  Finish `prefix(3)`
    - Decision: Position `2` has block size `2` and stores public
      indexes `0..1`. Add `3`, then move to `2-2=0`.
    - After: `values=[2,1,5,3]`, `sums=[0,2,3,5,11]`,
      `position=0`, `total=8`, `rangeAnswer=pending`.
    - Known correct: The disjoint blocks `0..1` and `2` exactly
      cover the prefix `0..2`.
3.  Compute the requested range
    - Decision: Subtract `prefix(1)=2` from `prefix(3)=8`.
    - After: `values=[2,1,5,3]`, `sums=[0,2,3,5,11]`,
      `position=0`, `total=8`, `rangeAnswer=6`.
    - Known correct: Subtraction removes index `0`, leaving only
      values `1+5`.
4.  Begin setting index `2` to `6`
    - Decision: Compute `delta=1`; add it to position `3`,
      whose block is public index `2`. Move to `3+1=4`.
    - After: `values=[2,1,6,3]`, `sums=[0,2,3,6,11]`,
      `position=4`, `total=8`, `rangeAnswer=6`.
    - Known correct: The saved one-value block now matches the
      changed array value.
5.  Finish setting index `2`
    - Decision: Position `4` has block size `4` and contains
      index `2`. Add `1`, then move to `4+4=8`, past the tree.
    - After: `values=[2,1,6,3]`, `sums=[0,2,3,6,12]`,
      `position=8`, `total=8`, `rangeAnswer=6`.
    - Known correct: Every stored block containing index `2` was
      repaired, and every other block stayed unchanged.

###### Why it works

The shared invariant is that every saved number equals the sum of its
named block.

In a segment tree, leaves satisfy the invariant because each contains
one array value. A parent satisfies it after adding its two child
blocks. An update rebuilds every ancestor of the changed leaf, so the
invariant remains true. A query combines non-overlapping blocks whose
union is the requested range, so it includes every requested value
once and no value outside the range.

Fenwick position `p` stores the block of `p & -p` values ending at
public index `p-1`. During a prefix query, subtracting that block size
moves to the position immediately before the block just counted. The
visited blocks therefore do not overlap and together cover the whole
prefix. Subtracting the prefix before `left` from the prefix through
`right` leaves exactly the inclusive range `left..right`.

During a Fenwick update, starting at `index+1` and adding each block
size visits exactly the stored blocks containing that public index.
Adding the same `delta` repairs every affected sum. Query positions
strictly decrease, and update positions strictly increase, so both
loops finish.

###### Go example

This implementation stores the original values so `set` can accept a
replacement value rather than requiring the caller to calculate a
delta. Public indexes are zero-based. Internal Fenwick positions are
one-based.

Require `0 <= index < len(values)` for `set`. Require
`0 <= left <= right < len(values)` for `rangeSum`. Sums must fit in
`int`. The subtraction `newValue-oldValue` must also fit in `int`.

```go
type fenwick struct {
	values []int
	sums   []int
}

func newFenwick(values []int) *fenwick {
	tree := &fenwick{
		values: append([]int(nil), values...),
		sums:   make([]int, len(values)+1),
	}
	for index, value := range values {
		tree.addDelta(index, value)
	}
	return tree
}

// blockSize returns the number of array values summarized at a
// positive, one-based Fenwick position.
func blockSize(oneBasedPosition int) int {
	// Two's-complement negation and AND isolate the lowest-value 1 bit;
	// its numeric value is this position's block length.
	return oneBasedPosition & -oneBasedPosition
}

func (tree *fenwick) addDelta(index, delta int) {
	// Convert public index to its one-based Fenwick position. Every
	// visited block contains that public index.
	position := index + 1
	for position < len(tree.sums) {
		tree.sums[position] += delta
		// Adding this block's size reaches the next larger stored block
		// that also contains index.
		position += blockSize(position)
	}
}

func (tree *fenwick) set(index, newValue int) {
	// Convert replacement into the change shared by every block that
	// contains index, then keep the original-value copy synchronized.
	delta := newValue - tree.values[index]
	tree.values[index] = newValue
	tree.addDelta(index, delta)
}

// prefix returns the sum of public indexes in [0,end).
func (tree *fenwick) prefix(end int) int {
	total := 0
	for position := end; position > 0; {
		// This block ends immediately before the uncounted prefix boundary.
		total += tree.sums[position]
		// Subtract its size to move before the block just counted.
		position -= blockSize(position)
	}
	return total
}

// rangeSum includes both left and right.
func (tree *fenwick) rangeSum(left, right int) int {
	return tree.prefix(right+1) - tree.prefix(left)
}
```

###### Work and memory

Let `n` be the number of array values. A direct range scan takes up to
`O(n)` time per query. Repairing a prefix-sum array takes up to `O(n)`
time per update.

A segment tree can be built in `O(n)` time. The shown Fenwick
constructor adds `n` values in `O(n log n)` time. After construction,
either structure handles one point update or compatible range query
in `O(log n)` time and uses `O(n)` memory.

###### When to use it

Use one of these structures when point updates and range questions
alternate often. Choose a Fenwick tree for the smallest implementation
when the questions are prefix sums or range sums. Choose a segment
tree when visible ranges help clarity or when nodes must store another
summary that can be combined from neighboring blocks, such as a
minimum or maximum.

###### When not to use it

Use direct scans when there are only a few operations. Use prefix sums
when values never change. Use a difference array when all range
updates happen before one final rebuild. These simpler choices have
less indexing machinery.

###### Common mistakes and edge cases

> **Warning:** Fenwick position `0` has block size `0`. Starting an
> update there would never advance. Convert a valid public index with
> `index+1`, and reject negative or out-of-range indexes first.

Do not mix the inclusive public range `left..right` with the half-open
prefix `[0,end)`. An empty tree has no valid nonempty range query.
Check whether sums can overflow `int` for the problem's limits.

For a segment tree, define what a query returns when a node does not
overlap the requested range. For sums that neutral value is `0`.

###### Related material

- [C.1.3.10.1 Coordinate compression](#c13101-coordinate-compression)
  maps large values to tree indices.
- [C.1.3.15 Prefix sums](#c1315-prefix-sums) is simpler for fixed
  data.
- [C.1.3.16 Difference arrays](#c1316-difference-arrays) is simpler
  when all updates precede one rebuild.

<a id="c1318-greedy"></a>

##### C.1.3.18 Make a provably safe permanent choice (greedy)

###### The problem it solves

Given a problem that asks for the best possible, or **optimal**,
result, build the answer through permanent local choices. Each choice
needs a proof that some optimal complete answer includes it, so the
algorithm never needs to undo that choice or explore the rejected
alternatives.

###### How it works

Choose a meeting now only after proving it cannot block a better final
schedule. A **greedy** method makes one such safe choice at a time and
never needs to undo it.

###### Step-by-step examples

Choose the most non-overlapping meetings from
`(1,9),(2,4),(5,7),(8,10)`; expect three meetings. One meeting may
start exactly when the previous meeting ends.

Initial state: `meetings=[(1,9),(2,4),(5,7),(8,10)]`,
`candidate=unset`, `selected=[]`, `lastEnd=-infinity`.

1.  Sort by ending time
    - Decision: Put earlier ending times first.
    - After: `meetings=[(2,4),(5,7),(1,9),(8,10)]`,
      `candidate=(2,4)`, `selected=[]`, `lastEnd=-infinity`.
    - Known correct: The scan will always see the compatible
      meeting that leaves the most time first.
2.  Inspect `(2,4)`
    - Decision: Its start is allowed, so select it.
    - After: `meetings=[(2,4),(5,7),(1,9),(8,10)]`,
      `candidate=(5,7)`, `selected=[(2,4)]`, `lastEnd=4`.
    - Known correct: Finishing at `4` leaves at least as much room
      as any other possible first meeting.
3.  Inspect `(5,7)`
    - Decision: Its start `5` is at least `lastEnd=4`; select it.
    - After: `meetings=[(2,4),(5,7),(1,9),(8,10)]`,
      `candidate=(1,9)`, `selected=[(2,4),(5,7)]`, `lastEnd=7`.
    - Known correct: It does not overlap either selected meeting.
4.  Inspect `(1,9)`
    - Decision: Its start `1` is before `lastEnd=7`; skip it.
    - After: `meetings=[(2,4),(5,7),(1,9),(8,10)]`,
      `candidate=(8,10)`, `selected=[(2,4),(5,7)]`, `lastEnd=7`.
    - Known correct: A meeting ending at `9` would leave no more
      future room than the selected meeting ending at `7`.
5.  Inspect `(8,10)`
    - Decision: Its start `8` is at least `lastEnd=7`; select it.
    - After: `meetings=[(2,4),(5,7),(1,9),(8,10)]`,
      `candidate=done`, `selected=[(2,4),(5,7),(8,10)]`,
      `lastEnd=10`.
    - Known correct: The new meeting starts after the previous one
      ends.

Focused failure case: for coin values `[1,3,4]` and target `6`,
"take the largest coin" chooses `4+1+1`, using three coins. The
expected best result is `3+3`, using two. The apparently largest
choice blocks the better combination, so that rule is unsafe.

###### Why it works

Let `E` be the compatible meeting with the earliest end and `F` be
the first meeting in any optimal remaining schedule. Because
`E.end <= F.end`, every meeting that can follow `F` can also follow
`E`. Replacing `F` with `E` preserves validity and the number of
meetings. The remaining meetings form the same smaller problem, so
the replacement argument repeats. The scan examines one finite input
meeting per iteration and therefore finishes.

###### Go example

Each meeting must satisfy `start <= end`. The function sorts the
supplied slice in place; copy it first when mutation is forbidden.
For the trace, the input is `[]meeting{{1,9},{2,4},{5,7},{8,10}}`.
The returned slice is `[]meeting{{2,4},{5,7},{8,10}}`. The trace's
`lastEnd` is the code's `endOfLastSelectedMeeting`.

```go
import (
	"math"
	"sort"
)

type meeting struct{ start, end int }

func selectMostMeetings(meetings []meeting) []meeting {
	// Sort the supplied slice in place by end time. Finishing earlier
	// leaves at least as much room for every later meeting, but the
	// caller's order changes.
	sort.Slice(meetings, func(leftIndex, rightIndex int) bool {
		return meetings[leftIndex].end < meetings[rightIndex].end
	})
	selected := make([]meeting, 0, len(meetings))
	endOfLastSelectedMeeting := math.MinInt
	for _, candidateMeeting := range meetings {
		if candidateMeeting.start >= endOfLastSelectedMeeting {
			// Accept the earliest-finishing compatible meeting and move the
			// boundary to its end; incompatible meetings are skipped.
			selected = append(selected, candidateMeeting)
			endOfLastSelectedMeeting = candidateMeeting.end
		}
	}
	return selected
}
```

###### Work and memory

Sorting costs `O(n log n)` time, and the scan costs `O(n)` time. The
scan uses `O(1)` extra memory when the meetings are sorted in place,
excluding the space used to return the selected meetings.

###### When to use it

The problem asks for a minimum or maximum built through permanent
choices, and one available choice appears to leave every useful future
option intact. Scheduling, selecting items, and connecting components
often have this form. The optimization goal alone is not evidence that
greedy works.

###### When not to use it

A choice that looks best now can block a better combination later and
no safe replacement proof exists. Use dynamic programming when
different choices lead to repeated future states, or backtracking when
the alternatives must be explored and cannot be summarized that way.

###### Common mistakes and edge cases

There is no universal rule such as "take the smallest" or "take the
largest." The rule and proof depend on the problem. In interval
scheduling, sorting by start time or by shortest duration does not
have the same safe-choice proof as sorting by end time.

###### Related material

- [C.1.3.10 Dynamic programming](#c1310-dynamic-programming) keeps
  competing future states.
- [C.1.3.19 Backtracking](#c1319-backtracking) explores choices that
  cannot be made permanent.
- [C.1.3.26 Minimum spanning tree](#c1326-minimum-spanning-tree) uses
  the cut property for greedy edges.

<a id="c1319-backtracking"></a>

##### C.1.3.19 Try a choice and undo it (backtracking)

###### The problem it solves

Given a set of choices and validity rules, construct one or every
valid complete answer when no single next choice can safely replace
all alternatives. Try one choice, continue while the partial answer
can still succeed, then undo that choice so the next branch begins
from the same state.

###### How it works

Try a key in a lock, return it to the ring if it fails, then try the
next key. **Backtracking** builds an answer the same way: make a
choice, explore it, undo it, and try another choice from the same
point. Each choice starts a **branch**; choices from the same point
are **sibling branches**.

###### Step-by-step examples

Generate all subsets of `[a,b]`; expect `[a,b]`, `[a]`, `[b]`, and
`[]`, in that order. Here `path` holds included values, and `index`
names the value whose include/exclude choice comes next. `result`
holds copies of completed paths.

Initial state: `index=0`, `path=[]`, `result=[]`.

1.  Include `a`
    - Decision: Append `a` and visit index `1`.
    - After: `index=1`, `path=[a]`, `result=[]`.
    - Known correct: This branch represents subsets containing
      `a`.
2.  Include `b`
    - Decision: Append `b`, visit index `2`, and save a copy.
    - After: `index=1`, `path=[a,b]`, `result=[[a,b]]`.
    - Known correct: Both input values were included once.
3.  Undo and exclude `b`
    - Decision: Remove `b`, then visit index `2` without it.
    - After: `index=0`, `path=[a]`, `result=[[a,b],[a]]`.
    - Known correct: Both choices for `b` under the include-`a`
      branch are complete.
4.  Undo and exclude `a`
    - Decision: Remove `a`, then visit index `1`
      without appending `a`.
    - After: `index=1`, `path=[]`, `result=[[a,b],[a]]`.
    - Known correct: The new branch contains no choice left over
      from the include-`a` branch.
5.  Include `b`
    - Decision: Append `b`, visit index `2`, and save a copy.
    - After: `index=1`, `path=[b]`, `result=[[a,b],[a],[b]]`.
    - Known correct: This is the branch excluding `a` and
      including `b`.
6.  Undo and exclude `b`
    - Decision: Remove `b`, visit index `2` without it, and save a
      copy.
    - After: `index=done`, `path=[]`,
      `result=[[a,b],[a],[b],[]]`.
    - Known correct: Every value had exactly one include or
      exclude decision on each completed path.

Saving `path` without copying would be wrong because later undo steps
would mutate the storage used by earlier results.

###### Why it works

At the start of each recursive call, the partial answer and its
supporting facts describe exactly the choices on the path to that
call. After the call returns, undoing its choice restores exactly what
the caller had before trying that branch. At each level, the include
and exclude calls cover both possibilities without overlap. The index
increases on every call, so it eventually reaches the finite input
length and recursion finishes.

###### Go example

```go
func subsets[T any](values []T) [][]T {
	// path is the current choice sequence. result owns one copied row
	// for every completed choice path.
	result := make([][]T, 0)
	path := make([]T, 0, len(values))
	var visit func(int)
	visit = func(index int) {
		if index == len(values) {
			// Copy the completed path because backtracking will reuse its
			// storage.
			result = append(result, append([]T(nil), path...))
			return
		}
		// Explore the include branch, undo that choice, then explore
		// exclusion.
		path = append(path, values[index])
		visit(index + 1)
		path = path[:len(path)-1]
		visit(index + 1)
	}
	visit(0)
	return result
}
```

###### Work and memory

Generating subsets visits `2^n` complete branches. Saving a copy of
every subset takes `O(n * 2^n)` total time and `O(n * 2^n)` output
space. The current path and unfinished function calls use `O(n)` extra
memory, excluding the output.

###### When to use it

The answer is assembled one decision at a time, and each partial
answer has a clear set of legal next choices. You may need to find one
valid construction, generate every construction, or compare a
manageable number of constructions. Input rules often let you stop
exploring a branch as soon as its partial answer cannot become valid
or cannot beat the best answer already found.

Backtracking is a good fit when all of these are true:

1.  A complete answer can be represented as a sequence of choices.
2.  From a partial answer, you can enumerate the choices allowed next.
3.  No single next choice can be proved safe enough to discard all
    alternatives.
4.  There are few enough possibilities, or invalid branches can be
    stopped early enough, to make exploration practical.

###### When not to use it

- Use **greedy** when one next choice can be proved not to reduce the
  best result still possible, so the other branches may be discarded.
- Use **dynamic programming** when many choice sequences reach the
  same remaining situation and only its best value, count, or yes/no
  result matters. Save that result instead of rebuilding the same
  search tree.
- Use ordinary **[depth-first search
  (DFS)](#c1320-depth-first-search)** when traversing an existing
  graph or tree without building and undoing a possible answer.
- Avoid exhaustive backtracking when there are too many possibilities
  and no useful early-stopping rule or repeated smaller question
  exists.

###### Common mistakes and edge cases

> **Warning:** Saving `path` itself instead of a copy makes all subset
> rows share one backing array; later undo operations can turn
> previously saved answers into the final path state. Copy at every
> completed answer.

Empty input has one completed choice path, the empty subset. Repeated
values need a stated duplicate-output rule. Check the exponential
output size before accepting a large input.

###### Related material

- [C.1.3.20 Depth-first search](#c1320-depth-first-search) describes
  the exploration order.
- [C.1.3.10 Dynamic programming](#c1310-dynamic-programming) saves
  repeated states when individual constructions are unnecessary.
- [A.7 Combinatorics and output
  size](#a7-combinatorics-and-output-size) counts unavoidable output.

<a id="c1320-depth-first-search"></a>

##### C.1.3.20 Finish one route before the next (depth-first search)

###### The problem it solves

Given a graph, grid, tree, or choice space and a starting point,
explore a route as far as possible before returning to the newest
unfinished branch. Use this order to mark everything reachable,
enumerate complete paths, or combine subtree results when the answer
does not require visiting positions by fewest moves first.

###### How it works

In a maze, follow one corridor until it ends, then return to the
latest junction with an untried path. **Depth-first search (DFS)**
explores a graph in that order. Each location is a **vertex**, and
each direct connection is an **edge**.

For ordinary reachability, one `seen` flag per vertex prevents repeat
visits. Mark a vertex before following its edges. Otherwise an edge
back to an earlier vertex can enter it again before it is marked. The
contract is `seen[vertex] -> whether that vertex has been entered`.
For example, `seen[B]=true` means DFS must not enter `B` again.

###### Step-by-step examples

Traverse edges `A-B`, `A-C`, and `B-D` from `A`; one valid DFS order
is `A,B,D,C`.

Initial state: `seen={}`, `order=[]`, `activeRoute=[]`.

1.  Enter `A`
    - Decision: Mark `A` and append it to the order.
    - After: `seen={A}`, `order=[A]`, `activeRoute=[A]`.
    - Known correct: A later edge back to `A` cannot schedule it
      again.
2.  Follow edge `A-B`
    - Decision: Enter unseen `B`, mark it, and append it.
    - After: `seen={A,B}`, `order=[A,B]`, `activeRoute=[A,B]`.
    - Known correct: Both entered vertices appear exactly once.
3.  Follow edge `B-D`
    - Decision: Enter unseen `D`, mark it, and append it.
    - After: `seen={A,B,D}`, `order=[A,B,D]`,
      `activeRoute=[A,B,D]`.
    - Known correct: Every explored edge from `B` has led to its
      unseen destination.
4.  Return from `D`, then `B`
    - Decision: Neither has another unseen neighbor, so return to `A`.
    - After: `seen={A,B,D}`, `order=[A,B,D]`, `activeRoute=[A]`.
    - Known correct: Every route beginning through `B` is fully
      explored.
5.  Follow edge `A-C`
    - Decision: Enter and mark `C`, then return because it has no
      unseen neighbor.
    - After: `seen={A,B,C,D}`, `order=[A,B,D,C]`, `activeRoute=[]`.
    - Known correct: Every vertex reachable from `A` was entered
      once.

###### Why it works

When `visit(vertex)` begins, the vertex is marked and will never be
scheduled again. The call examines every outgoing edge; each unseen
neighbor is explored under the same rule before the call returns. By
induction over the return path, returning from a vertex means every
vertex reachable through its outgoing edges has been visited. Marking
prevents cycles from creating repeated work, and the finite graph lets
every call eventually finish.

###### Go example

Validate `start` and every stored neighbor before calling this
fragment. Use an explicit stack when a valid path may be too deep for
recursion.
For the trace, map `A=0`, `B=1`, `C=2`, and `D=3`, and pass
`graph=[[1,2],[0,3],[0],[1]]` with `start=0`. The returned integer
order `[0,1,3,2]` maps back to `[A,B,D,C]`.

```go
func depthFirst(graph [][]int, start int) []int {
	// seen prevents cycles from causing repeated visits; order records
	// the first time each reachable vertex is entered.
	seen, order := make([]bool, len(graph)), make([]int, 0, len(graph))
	var visit func(int)
	visit = func(vertex int) {
		// Mark before following edges so a back edge cannot re-enter this
		// vertex.
		seen[vertex], order = true, append(order, vertex)
		// Explore every outgoing edge; recurse only when its destination
		// has not already been reached through another path.
		for _, next := range graph[vertex] {
			if !seen[next] {
				visit(next)
			}
		}
	}
	visit(start)
	return order
}
```

###### Work and memory

With an adjacency list, which stores each location's neighbors, a
complete search takes `O(V + E)` time because it visits each of the
`V` locations and examines each of the `E` connections a constant
number of times. Reached state and the explicit or recursive stack use
`O(V)` extra memory in the worst case.

###### When to use it

You need to visit an entire connected region, inspect every complete
path, or calculate a tree result from the results below each node. You
do not need the first discovered path to use the fewest moves.

###### When not to use it

You need a route with the fewest equal-cost moves; use breadth-first
search instead.
Also avoid recursive DFS when a valid path can be too deep for the
call stack unless the environment guarantees enough stack space.

###### Common mistakes and edge cases

Nested function calls consume memory. A path with too many steps may
exceed the program's call limit, so use your own stack collection for
very deep inputs. Mark each vertex before exploring its neighbors;
marking after recursion can revisit a cycle forever.

One call visits only the vertices reachable from `start`. To traverse
a disconnected graph, scan all vertices and start another DFS at each
still-unseen vertex. Neighbor order can change the traversal order,
but not which reachable vertices are visited.

###### Related material

- [B.5 Graphs and grids](#b5-graphs-and-grids) explains graph
  representation.
- [C.1.3.21 Breadth-first search](#c1321-breadth-first-search) finds
  fewest equal-cost moves.
- [C.1.3.19 Backtracking](#c1319-backtracking) adds and undoes choices
  while following depth-first order.

<a id="c1321-breadth-first-search"></a>

##### C.1.3.21 Visit by distance (breadth-first search)

###### The problem it solves

Given a starting point and connections whose moves all have the same
cost, return the nearest matching position or the fewest moves to
every reachable position. All positions at distance `d` must be
processed before any position at distance `d+1`, so the first visit
records a shortest route.

###### How it works

Ripples reach every point one step from a splash before points two
steps away. **Breadth-first search (BFS)** explores a graph in the
same layers, so each layer is one more move from the start. It uses a
first-in, first-out queue: remove the oldest discovered vertex, add
each unseen neighbor at the next distance, and mark that neighbor as
soon as it enters the queue. Here,
`distanceByVertex[vertex] -> fewest edges from the start`, with `-1`
meaning undiscovered. For example, `distanceByVertex[E]=2` means the
shortest route from `A` to `E` uses two edges.

###### Step-by-step examples

For this undirected graph, find the fewest edges from `A` to every
vertex:

```text
A --- B
|     |
C --- D
 \   /
   E
```

The neighbor order is `A:[B,C]`, `B:[A,D]`, `C:[A,D,E]`,
`D:[B,C,E]`, and `E:[C,D]`. Expect
`distanceByVertex={A:0,B:1,C:1,D:2,E:2}`. The queue below lists only
vertices still waiting to be removed.

Initial state: `queue=[]`,
`distanceByVertex={A:-1,B:-1,C:-1,D:-1,E:-1}`.

1.  Initialize `A`
    - Decision: Set its distance to `0` and enqueue it.
    - After: `queue=[A]`,
      `distanceByVertex={A:0,B:-1,C:-1,D:-1,E:-1}`.
    - Known correct: The start is zero moves from itself.
2.  Remove `A`
    - Decision: Discover `B` and `C` at distance `1`; enqueue them
      in that order.
    - After: `queue=[B,C]`,
      `distanceByVertex={A:0,B:1,C:1,D:-1,E:-1}`.
    - Known correct: `B` and `C` form the complete distance-1
      layer.
3.  Remove `B`
    - Decision: Skip discovered `A`; discover `D` at distance `2`
      and mark it when it enters the queue.
    - After: `queue=[C,D]`,
      `distanceByVertex={A:0,B:1,C:1,D:2,E:-1}`.
    - Known correct: `C` still waits at distance `1`, ahead of `D`
      at distance `2`.
4.  Remove `C`
    - Decision: Skip `A` and the already-enqueued `D`; discover
      `E` at distance `2` and enqueue it once.
    - After: `queue=[D,E]`,
      `distanceByVertex={A:0,B:1,C:1,D:2,E:2}`.
    - Known correct: Both distance-1 vertices were removed before
      either distance-2 vertex. Marking `D` on insertion prevented
      `C` from enqueueing it again.
5.  Remove `D`
    - Decision: Skip its discovered neighbors, including `E`.
      Reaching `E` through `A-B-D-E` would take `3` edges.
    - After: `queue=[E]`,
      `distanceByVertex={A:0,B:1,C:1,D:2,E:2}`.
    - Known correct: The earlier route `A-C-E` keeps the shorter
      distance `2`.
6.  Remove `E`
    - Decision: Both neighbors are already discovered, so add
      nothing.
    - After: `queue=[]`,
      `distanceByVertex={A:0,B:1,C:1,D:2,E:2}`.
    - Known correct: The queue is empty, and every reachable
      vertex has its fewest-edge distance from `A`.

###### Why it works

The queue contains discovered positions in nondecreasing distance
order. When a vertex at distance `d` is removed, every newly
discovered neighbor receives distance `d+1`; any undiscovered route to
that neighbor must pass through the current or a later layer and
therefore cannot be shorter. Marking on insertion makes that first
distance permanent. Each vertex is added at most once, so the finite
search terminates.

###### Go example

Validate `start` and every stored neighbor before calling this
fragment. Returned distance `-1` means unreachable.
For the trace, map `A=0`, `B=1`, `C=2`, `D=3`, and `E=4`. Pass
`graph=[[1,2],[0,3],[0,3,4],[1,2,4],[2,3]]` with `start=0`.
The returned slice `[0,1,1,2,2]` stores the distances shown for
`A` through `E` in that index order.

```go
func breadthFirstDistances(graph [][]int, start int) []int {
	// -1 means undiscovered, while nonnegative values are finalized edge
	// counts from start in this unweighted graph.
	distanceByVertex := make([]int, len(graph))
	for vertex := range distanceByVertex {
		distanceByVertex[vertex] = -1
	}
	queue := []int{start}
	distanceByVertex[start] = 0
	// FIFO processing visits all distance-d vertices before distance
	// d+1.
	nextQueueIndex := 0
	for nextQueueIndex < len(queue) {
		vertex := queue[nextQueueIndex]
		for _, neighboringVertex := range graph[vertex] {
			if distanceByVertex[neighboringVertex] == -1 {
				// The first discovery is a shortest path: extend the current
				// shortest path by this one edge, then enqueue the vertex once.
				distanceByVertex[neighboringVertex] = distanceByVertex[vertex] + 1
				queue = append(queue, neighboringVertex)
			}
		}
		nextQueueIndex++
	}
	return distanceByVertex
}
```

###### Work and memory

With an adjacency list, which stores each location's neighbors,
ordinary BFS takes `O(V + E)` time and `O(V)` extra memory for its
queue, distances, and reached state. Here, `V` counts locations and
`E` counts connections.

###### When to use it

You need the fewest moves through a network where every move costs the
same, the nearest reachable position, or tree nodes grouped by depth.

###### When not to use it

Moves have different costs. The route with the fewest moves may then
cost more than a longer route; use Dijkstra's algorithm when all costs
are zero or greater.

###### Common mistakes and edge cases

> **Warning:** Marking only when a vertex is removed lets two vertices
> in the same layer enqueue it twice. Mark at insertion so the first,
> shortest discovery is the only one that adds it to the queue.

Validate the start and every neighbor index. Empty graphs have no
valid start, and unreachable vertices retain the declared `-1`
distance.

###### Related material

- [B.3 Queues and deques](#b3-queues-and-deques) explains FIFO
  behavior.
- [C.1.3.20 Depth-first search](#c1320-depth-first-search) explores
  one route first.
- [C.1.3.23 Dijkstra shortest paths](#c1323-dijkstra-shortest-paths)
  handles unequal non-negative costs.

<a id="c1322-topological-sort"></a>

##### C.1.3.22 Release ready tasks (topological sort)

###### The problem it solves

Given directed tasks and prerequisite relationships, return an order
in which every prerequisite appears before every task that depends on
it, or report that no such order exists. A dependency cycle makes the
order impossible because every task in the cycle waits for another
unfinished task in that same cycle.

###### How it works

While cooking, chop ingredients before heating them and heat them
before serving. A **topological order** lists every task after all of
its prerequisites, repeatedly taking tasks that are ready. The lookup
`unmetPrerequisites[task] -> unfinished prerequisites for task`
controls readiness. A stored zero means the task may enter the ready
queue.

###### Step-by-step examples

Use rules `Shop -> Cook`, `Prep -> Cook`, and `Cook -> Eat`. If ready
tasks enter the queue in the listed task order, expect one valid order
`[Shop,Prep,Cook,Eat]`.

The Go representation uses task order `[Shop,Prep,Cook,Eat]`, so the
task indexes are `Shop=0`, `Prep=1`, `Cook=2`, and `Eat=3`.
`graph[prerequisite]` stores the tasks it unlocks, giving
`graph=[[2],[2],[3],[]]`. The integer output `[0,1,2,3]` maps back to
the named order above.

Initial state: `unmetPrerequisites=unset`, `queue=[]`, `output=[]`.

1.  Count unfinished prerequisites
    - Decision: Count incoming rules and enqueue each zero count.
    - After: `unmetPrerequisites={Shop:0,Prep:0,Cook:2,Eat:1}`,
      `queue=[Shop,Prep]`, `output=[]`.
    - Known correct: `Shop` and `Prep` are the only tasks ready
      immediately.
2.  Complete `Shop`
    - Decision: Output it and reduce `Cook`'s unfinished count from
      `2` to `1`.
    - After: `unmetPrerequisites={Shop:0,Prep:0,Cook:1,Eat:1}`,
      `queue=[Prep]`, `output=[Shop]`.
    - Known correct: `Cook` is not ready while `Prep` remains
      unfinished.
3.  Complete `Prep`
    - Decision: Output it and reduce `Cook`'s count from `1` to `0`.
    - After: `unmetPrerequisites={Shop:0,Prep:0,Cook:0,Eat:1}`,
      `queue=[Cook]`, `output=[Shop,Prep]`.
    - Known correct: `Cook` enters only after both prerequisites
      appear in the output.
4.  Complete `Cook`
    - Decision: Output it and reduce `Eat`'s count to zero.
    - After: `unmetPrerequisites={Shop:0,Prep:0,Cook:0,Eat:0}`,
      `queue=[Eat]`, `output=[Shop,Prep,Cook]`.
    - Known correct: Every output task follows all its prerequisites.
5.  Complete `Eat`
    - Decision: Output it; it unlocks nothing else.
    - After: `unmetPrerequisites={Shop:0,Prep:0,Cook:0,Eat:0}`,
      `queue=[]`,
      `output=[Shop,Prep,Cook,Eat]`.
    - Known correct: Every arrow points from an earlier output
      position to a later one.

Focused failure case: adding `Eat -> Shop` leaves only `Prep` ready.
After outputting `Prep`, `Cook` still waits for `Shop`, while
`Shop -> Cook -> Eat -> Shop` forms a cycle. The queue empties with
only one of four tasks processed, so no complete order exists.

###### Why it works

The queue contains exactly the tasks not yet output whose
prerequisites are all finished. Removing one task and lowering the
counts of its dependents preserves that fact. Each task enters the
output at most once, so the finite graph makes the process end.

If every task is output, each one appears after every prerequisite, so
the output is a valid order. If the queue empties early, every
remaining task has a remaining prerequisite. Following those
prerequisites in a finite graph eventually repeats a task, proving a
cycle. A cycle cannot have a valid prerequisite order.

###### Go example

Every stored destination must be a valid index in `graph`. For the
traced input, `topologicalOrder([][]int{{2}, {2}, {3}, {}})` returns
`([0,1,2,3], true)`, which maps to
`([Shop,Prep,Cook,Eat], true)` using the declared task order.

```go
func topologicalOrder(graph [][]int) ([]int, bool) {
	// unmetPrerequisites stores each vertex's unfinished incoming-edge
	// count.
	unmetPrerequisites := make([]int, len(graph))
	// Count every incoming edge once, including parallel prerequisites
	// if present.
	for _, destinations := range graph {
		for _, destinationVertex := range destinations {
			unmetPrerequisites[destinationVertex]++
		}
	}
	// Vertices with no prerequisites can appear first.
	queue := make([]int, 0, len(graph))
	for vertex := range unmetPrerequisites {
		if unmetPrerequisites[vertex] == 0 {
			// This vertex is ready immediately because it has no incoming
			// edge.
			queue = append(queue, vertex)
		}
	}
	// Removing a vertex satisfies each outgoing prerequisite. A neighbor
	// becomes ready exactly when its count reaches zero.
	order := make([]int, 0, len(graph))
	nextQueueIndex := 0
	for nextQueueIndex < len(queue) {
		// Emit the oldest ready vertex without shifting the queue slice.
		vertex := queue[nextQueueIndex]
		order = append(order, vertex)
		for _, destinationVertex := range graph[vertex] {
			// Emitting vertex satisfies one prerequisite of this destination.
			unmetPrerequisites[destinationVertex]--
			if unmetPrerequisites[destinationVertex] == 0 {
				// The destination is ready after all its prerequisites reach the
				// output.
				queue = append(queue, destinationVertex)
			}
		}
		nextQueueIndex++
	}
	// A cycle leaves a positive unfinished-prerequisite count, so not
	// every vertex is emitted.
	return order, len(order) == len(graph)
}
```

The boolean is false when a dependency cycle prevents a complete
order.

###### Work and memory

This queue form takes `O(V + E)` time and `O(V)` extra memory. Here,
`V` counts tasks and `E` counts prerequisite rules. An adjacency list
stores the tasks unlocked by each task; the stored graph itself is not
included in the extra-memory count.

###### When to use it

The prompt describes courses, builds, jobs, or other tasks with rules
such as "A must finish before B."

###### When not to use it

The relationships do not represent prerequisites, the graph may be
undirected, or the problem asks for shortest routes rather than a
valid dependency order.

###### Common mistakes and edge cases

> **Warning:** Reversing an edge changes the prerequisite meaning. If
> `A` must finish before `B`, store `A -> B` and increment
> `B`'s unfinished-prerequisite count, not `A`'s.

Validate every destination index. Repeated stored connections count as
repeated prerequisites unless they are removed from both the neighbor
lists and the unfinished-prerequisite counts. An empty task list has
the valid empty order.

###### Related material

- [B.3 Queues and deques](#b3-queues-and-deques) explains the ready
  queue's first-in, first-out behavior.
- [C.1.3.20 Depth-first search](#c1320-depth-first-search) provides
  the postorder alternative.
- [B.5 Graphs and grids](#b5-graphs-and-grids) defines directed edges.

<a id="c1323-dijkstra-shortest-paths"></a>

##### C.1.3.23 Finalize cheapest routes (Dijkstra)

###### The problem it solves

Given a weighted graph, one starting vertex, and edge costs that are
all zero or greater, return the cheapest route cost from the start to
every reachable vertex. The method must decide when a tentative route
is final so later work cannot reveal a cheaper route to that vertex.

###### How it works

On a map with travel costs, always process the unvisited location with
the cheapest known route from the starting location. This is
**Dijkstra's algorithm**. A known cost that may still improve is
**tentative**; a cost proved minimal is **finalized**. The contracts
are `distanceByVertex[vertex] -> cheapest route found so far` and
`finalized[vertex] -> whether that stored cost is proved minimal`.

###### Step-by-step examples

Apply the Go function below to the Network Delay Time problem. Each
entry `[source,destination,time]` is a directed signal route. For
`times=[[1,2,4],[1,3,1],[3,2,2],[2,4,1],[3,4,5]]`,
`numVertices=4`, and `startVertex=1`, the expected answer is `4`.

The input labels vertices from `1` through `4`, while Go slices use
indexes `0` through `3`. The graph-building loop converts every edge:
`[1,2,4]` becomes `(0->1,4)`, `[1,3,1]` becomes `(0->2,1)`,
`[3,2,2]` becomes `(2->1,2)`, `[2,4,1]` becomes `(1->3,1)`,
and `[3,4,5]` becomes `(2->3,5)`. Each stored pair means
`{destination,travelTime}`.

Initial state: `graph=[[],[],[],[]]`,
`distanceByVertex=unset`, `finalized=unset`,
`latestArrival=unset`.

1.  Build the adjacency list
    - Decision: Convert each one-based endpoint to a zero-based index
      and append its route to the source's list.
    - After: `graph=[[{1,4},{2,1}],[{3,1}],`
      `[{1,2},{3,5}],[]]`, `distanceByVertex=unset`,
      `finalized=unset`, `latestArrival=unset`.
    - Known correct: Every input route appears once under its source.
2.  Initialize the route state
    - Decision: Set every distance to `infinity`, then set start index
      `0` to `0`; no vertex is finalized yet.
    - After: `distanceByVertex=[0,infinity,infinity,infinity]`,
      `finalized=[false,false,false,false]`,
      `latestArrival=unset`.
    - Known correct: Vertex `1` receives the signal at time `0`; no
      route to another vertex has been examined.
3.  Finalize index `0`, input vertex `1`
    - Decision: The dense scan selects distance `0`. Its two routes
      offer arrival times `0+4=4` at index `1` and `0+1=1` at
      index `2`.
    - After: `distanceByVertex=[0,4,1,infinity]`,
      `finalized=[true,false,false,false]`,
      `latestArrival=unset`.
    - Known correct: Index `0` has final distance `0`; the other
      stored values are the cheapest routes found so far.
4.  Finalize index `2`, input vertex `3`
    - Decision: Its distance `1` is the smallest unfinished value.
      Route `2->1` offers `1+2=3`, replacing `4`; route `2->3`
      offers `1+5=6`.
    - After: `distanceByVertex=[0,3,1,6]`,
      `finalized=[true,false,true,false]`,
      `latestArrival=unset`.
    - Known correct: Index `2` has final distance `1`; index `1`
      now records the cheaper route `0->2->1`.
5.  Finalize index `1`, input vertex `2`
    - Decision: Its distance `3` is now smallest. Route `1->3`
      offers `3+1=4`, replacing `6`.
    - After: `distanceByVertex=[0,3,1,4]`,
      `finalized=[true,true,true,false]`,
      `latestArrival=unset`.
    - Known correct: Index `1` has final distance `3`; index `3`
      keeps the cheaper complete route found so far.
6.  Finalize index `3`, input vertex `4`
    - Decision: Select its distance `4`; it has no outgoing route.
    - After: `distanceByVertex=[0,3,1,4]`,
      `finalized=[true,true,true,true]`,
      `latestArrival=unset`.
    - Known correct: Every vertex now has its minimum arrival time.
7.  Produce the problem's answer
    - Decision: Every distance is finite, and their maximum is `4`.
    - After: `distanceByVertex=[0,3,1,4]`,
      `finalized=[true,true,true,true]`, `latestArrival=4`,
      `result=4`.
    - Known correct: All vertices have received the signal by time
      `4`, and vertex `4` cannot receive it earlier.

For the focused failure input `times=[[1,2,1]]`,
`numVertices=3`, and `startVertex=1`, index `2` remains at
`infinity`. The final scan returns `-1` because not every vertex can
receive the signal.

###### Why it works

Let `d` be the cheapest unfinished route cost. Every route through
another unfinished vertex already costs at least `d` before taking
another edge. Because an edge cannot subtract cost, none can produce
a route below `d`. The dense scan makes this same cheapest-unfinished
choice. Each outer iteration finalizes one vertex or stops when none
is reachable, so at most `V` iterations run.

After those iterations, each finite entry is that vertex's earliest
signal-arrival time. All vertices have received the signal exactly
when the latest of those arrivals occurs, so their maximum is the
required answer. An `infinity` entry proves that no route from the
start reaches that vertex, which requires the problem's `-1` result.

###### Go example

This function solves the problem traced above. It requires a positive
`numVertices`, a `startVertex` from `1` through `numVertices`, and
entries containing valid one-based endpoints and non-negative travel
times. Every relevant route sum must fit in `int`.

```go
import "math"

type networkEdge struct {
	destination int
	travelTime  int
}

func networkDelayTime(
	times [][]int,
	numVertices int,
	startVertex int,
) int {
	// Convert one-based input labels to a zero-based adjacency list.
	graph := make([][]networkEdge, numVertices)
	for _, time := range times {
		source := time[0] - 1
		destination := time[1] - 1
		edge := networkEdge{
			destination: destination,
			travelTime:  time[2],
		}
		graph[source] = append(graph[source], edge)
	}

	// MaxInt marks vertices that no examined route has reached.
	const infinity = math.MaxInt
	distanceByVertex := make([]int, numVertices)
	finalized := make([]bool, numVertices)
	for vertex := range distanceByVertex {
		distanceByVertex[vertex] = infinity
	}
	startIndex := startVertex - 1
	distanceByVertex[startIndex] = 0

	// Finalize one cheapest unfinished vertex per iteration.
	for range graph {
		selectedVertex := -1
		for candidateVertex := range graph {
			if finalized[candidateVertex] {
				continue
			}
			// A negative index means the scan has no candidate yet.
			if selectedVertex < 0 {
				selectedVertex = candidateVertex
				continue
			}
			candidateDistance :=
				distanceByVertex[candidateVertex]
			selectedDistance :=
				distanceByVertex[selectedVertex]
			if candidateDistance < selectedDistance {
				selectedVertex = candidateVertex
			}
		}
		if selectedVertex < 0 {
			// Every vertex is already final.
			break
		}
		if distanceByVertex[selectedVertex] == infinity {
			// No route can reach any unfinished vertex.
			break
		}
		finalized[selectedVertex] = true
		for _, edge := range graph[selectedVertex] {
			// Guard the addition before extending the selected route.
			if distanceByVertex[selectedVertex] >
				infinity-edge.travelTime {
				continue
			}
			candidateDistance :=
				distanceByVertex[selectedVertex] + edge.travelTime
			if candidateDistance <
				distanceByVertex[edge.destination] {
				// Keep the cheapest route found to this destination.
				distanceByVertex[edge.destination] = candidateDistance
			}
		}
	}

	// The final receiver determines when the whole network is reached.
	latestArrival := 0
	for _, arrivalTime := range distanceByVertex {
		if arrivalTime == infinity {
			return -1
		}
		if arrivalTime > latestArrival {
			latestArrival = arrivalTime
		}
	}
	return latestArrival
}
```

###### Work and memory

The shown dense version performs up to `V` selections, and each one
scans `V` locations. Together with examining `E` connections, that is
`O(V^2 + E)` time. Building the adjacency list stores `E` edges, and
the graph, distance slice, and finalized slice use `O(V + E)` memory.
Here, `V` counts network vertices and `E` counts directed signal
routes.

###### When to use it

You need the cheapest routes through a network, and every connection
has a cost of zero or greater.

###### When not to use it

Any connection may have a negative cost. If every connection has the
same cost, breadth-first search is simpler.

###### Common mistakes and edge cases

> **Warning:** Negative cost invalidates finalization. On edges
> `A->B=4`, `A->C=7`, and `C->B=-5`, finalizing `B=4` before
> processing `C` misses the cheaper route of cost `2`.

Validate the start and endpoints, reject negative costs, and guard
route-cost addition before evaluating it. For Network Delay Time,
also return the greatest finalized distance rather than one chosen
destination's distance, and return `-1` if any distance is infinite.

###### Related material

- [C.1.3.21 Breadth-first search](#c1321-breadth-first-search) is
  simpler for equal costs.
- [C.1.3.24 Bellman-Ford](#c1324-bellman-ford-shortest-paths) permits
  negative edges.
- [LeetCode 743: Network Delay
  Time](https://leetcode.com/problems/network-delay-time/) is the
  concrete problem traced above.

<a id="c1324-bellman-ford-shortest-paths"></a>

##### C.1.3.24 Ripple cheaper routes (Bellman-Ford)

###### The problem it solves

Given a weighted graph and one starting vertex, return cheapest route
costs even when some edges have negative cost, and report whether a
reachable negative-cost cycle makes those costs unbounded. Cheaper
routes may need to ripple through many edges, so no vertex can be
finalized after its first apparently cheapest route.

###### How it works

Imagine updating a road map whenever one route sign reveals a cheaper
trip to the next town. Repeat across every road until cheaper news has
had time to cross the whole map. This is the **Bellman-Ford** method.
For an edge `from -> to`, add its cost to the known cost of `from`.
If that route is cheaper than the stored cost of `to`, replace the
stored cost. This comparison is called **relaxing** the edge. Here,
`distanceByVertex[vertex] -> cheapest route cost found from start`,
with `infinity` meaning no route has reached that vertex yet.

###### Step-by-step examples

Let `C->B=-5`, `A->C=7`, and `A->B=4`; from `A`, expect `B=2`.
Scan the edges in that order so cheaper news must cross a later pass.

Initial state: `pass=1`, `nextEdge=C->B`,
`distanceByVertex={A:0,B:infinity,C:infinity}`, `changed=false`.

1.  Skip `C->B=-5`
    - Decision: `C` is still unreachable, so this edge cannot yet
      offer a route to `B`.
    - After: `pass=1`, `nextEdge=A->C`,
      `distanceByVertex={A:0,B:infinity,C:infinity}`,
      `changed=false`.
    - Known correct: No route to `C` has reached this earlier edge.
2.  Relax `A->C=7`
    - Decision: Candidate `0+7=7` is cheaper, so store `C=7`.
    - After: `pass=1`, `nextEdge=A->B`,
      `distanceByVertex={A:0,B:infinity,C:7}`, `changed=true`.
    - Known correct: The stored route is exactly `A->C`.
3.  Relax `A->B=4`
    - Decision: Candidate `0+4=4` is cheaper, so store `B=4`.
    - After: `pass=2`, `nextEdge=C->B`,
      `distanceByVertex={A:0,B:4,C:7}`, `changed=false`.
    - Known correct: The first pass found both one-edge routes, but
      the earlier `C->B` edge has not reused `C=7` yet.
4.  Ripple the cheaper route on pass `2`
    - Decision: `C->B` now offers `7-5=2`, replacing `B=4`.
      Neither edge from `A` improves afterward.
    - After: `pass=cycleCheck`, `nextEdge=C->B`,
      `distanceByVertex={A:0,B:2,C:7}`, `changed=true`.
    - Known correct: The two-edge route `A->C->B` became available
      only after its first edge was known from the prior pass.
5.  Run the cycle check
    - Decision: Scan every edge once more; no cost improves.
    - After: `pass=done`, `nextEdge=none`,
      `distanceByVertex={A:0,B:2,C:7}`, `changed=false`.
    - Known correct: No reachable negative cycle can keep reducing
      a cost.

Focused failure case: use edges `A->B=1`, `B->C=1`, and
`C->B=-2`. After the required two main passes, costs are `A=0`,
`B=-1`, and `C=1`. On the cycle-check pass, `B->C` offers
`-1+1=0`, which is below `C=1`. The cost can still fall because the
cycle `B->C->B` costs `-1` each time around, so no finite minimum
exists for `B` or `C`; return `false`.

###### Why it works

A finite shortest route never needs to repeat a location. When no
reachable negative cycle exists, any repeated cycle has cost zero or
greater, so removing it cannot make the route more expensive. A
repeated negative cycle would instead mean there is no finite shortest
route.

A route without repeated locations uses at most `V - 1` edges. After
pass `k`, every shortest route using at most `k` edges has had a
chance to propagate its cost. Therefore `V - 1` passes cover every
finite shortest route.

There are at most `V-1` main passes and one cycle-check pass. Each
scans the finite edge list, so the algorithm always stops.

One additional pass detects a **negative cycle**. If a reachable edge
can still lower a cost after every simple route was covered, some
repeated cycle makes the route cheaper each time around. No finite
minimum then exists for that cycle or anything reachable from it.

###### Go example

This fragment requires `numVertices > 0`, a valid `start`, valid
endpoints, and every evaluated sum
`distance[edge.from]+edge.cost` to fit in `int64`.
For the trace, map `A=0`, `B=1`, and `C=2`, then pass the edges
`{2,1,-5}`, `{0,2,7}`, and `{0,1,4}` in that order with
`numVertices=3` and `start=0`. The function returns
`([0,2,7], true)` for vertices `[A,B,C]`.

```go
import "math"

type directedEdge struct {
	from, to int
	cost     int64
}

func bellmanFord(
	numVertices, start int,
	edges []directedEdge,
) ([]int64, bool) {
	// MaxInt64 represents unreachable vertices; start reaches itself at
	// cost zero.
	distanceByVertex := make([]int64, numVertices)
	for vertex := range distanceByVertex {
		distanceByVertex[vertex] = math.MaxInt64
	}
	distanceByVertex[start] = 0
	// relax applies the one-edge recurrence and reports whether it
	// improved a reachable destination.
	relax := func(edge directedEdge) bool {
		if distanceByVertex[edge.from] == math.MaxInt64 {
			// No known path reaches the edge's source, so it cannot form a
			// candidate.
			return false
		}
		// Extend the best known source path by this edge. The function's
		// contract requires every relevant path sum to fit in int64.
		candidateDistance := distanceByVertex[edge.from] + edge.cost
		if candidateDistance >= distanceByVertex[edge.to] {
			// The destination already has an equal or cheaper known path.
			return false
		}
		// Record the cheaper path and report a change to the caller.
		distanceByVertex[edge.to] = candidateDistance
		return true
	}
	// Any simple shortest path has at most numVertices-1 edges. Stop
	// early if a full pass makes no improvement.
	for pass := 1; pass < numVertices; pass++ {
		changed := false
		for _, edge := range edges {
			// Keep changed true once any edge improves during this pass;
			// evaluating relax first ensures later edges are still processed.
			changed = relax(edge) || changed
		}
		if !changed {
			// With no new distance, another identical pass cannot improve
			// anything.
			break
		}
	}
	// A further improvement proves a reachable negative-cost cycle,
	// because all simple paths have already been considered.
	for _, edge := range edges {
		if relax(edge) {
			return distanceByVertex, false
		}
	}
	return distanceByVertex, true
}
```

The boolean is false when a reachable negative cycle can still lower
a cost.

###### Work and memory

If `V` is the number of locations and `E` is the number of
connections, up to `V - 1` full passes take `O(VE)` time. The stored
distances use `O(V)` extra memory, not counting the input map or
optional links used to rebuild the chosen routes.

###### When to use it

You need shortest routes from one starting location, some connections
may have negative costs, and you may need to detect a reachable
negative-cost cycle.

###### When not to use it

Every cost is non-negative and the graph is large enough that
Dijkstra's better running time matters. If a reachable negative cycle
may exist, do not report ordinary finite shortest costs for locations
affected by that cycle.

###### Common mistakes and edge cases

> **Warning:** Checking the special value alone does not stop
> ordinary addition from overflowing. If
> `distance[from]=math.MinInt64` and
> `cost=-1`,
> constructing the possible distance is invalid; enforce the route-sum
> bound or use checked arithmetic.

Validate `vertices`, `start`, and every endpoint before indexing. An
empty graph has no valid start. Parallel edges are allowed, but every
one is scanned; unreachable negative cycles do not invalidate
distances from this start.

###### Related material

- [C.1.3.23 Dijkstra shortest paths](#c1323-dijkstra-shortest-paths)
  is faster for non-negative edges.
- [B.5 Graphs and grids](#b5-graphs-and-grids) explains graph storage.
- [A.4 Integer limits and bits](#a4-integer-limits-and-bits) covers
  safe bounds.

<a id="c1325-disjoint-set-union"></a>

##### C.1.3.25 Track connected groups by leader (disjoint-set union)

###### The problem it solves

Given initially separate items, process a sequence of added undirected
connections and connectivity queries. Each query asks whether two
items now belong to the same connected group; the structure must merge
whole groups without searching every item after each new connection.

###### How it works

Imagine that every connected group has one shared group leader. To
check whether two items are connected, find each item's leader and
compare them. To connect two different groups, make one leader report
to the other. This structure is a **disjoint-set union (DSU)**. It is
also called **union-find**, after its two main operations. Its lookup
contracts are `parentByItem[item] -> next item toward the leader` and
`numItemsByRoot[root] -> number of items in that leader's group`.
Only self-parenting leaders have a meaningful stored size.

###### Step-by-step examples

Create three items, connect `0-1`, then `1-2`; expect `0` and `2` to
be connected. In each shown pair, `item:parent` names an item and its
stored parent.

Initial state: `parentByItem=unset`, `numItemsByRoot=unset`,
`nextUnion=0-1`.

1.  Initialize three groups
    - Decision: Make every item its own parent with size `1`.
    - After: `parentByItem=[0:0,1:1,2:2]`,
      `numItemsByRoot=[1,1,1]`,
      `nextUnion=0-1`.
    - Known correct: No connection exists, so each item is its own
      one-item group.
2.  Union `0` and `1`
    - Decision: Attach root `1` below root `0` and add its size.
    - After: `parentByItem=[0:0,1:0,2:2]`,
      `numItemsByRoot=[2,1,1]`,
      `nextUnion=1-2`.
    - Known correct: Exactly items `0` and `1` share a leader.
3.  Find leaders for union `1` and `2`
    - Decision: `find(1)` follows its parent and returns `0`;
      `find(2)` returns `2`.
    - After: `parentByItem=[0:0,1:0,2:2]`,
      `numItemsByRoot=[2,1,1]`,
      `nextUnion=roots(0,2)`.
    - Known correct: The roots differ, so the new edge joins two
      separate groups.
4.  Join the two groups
    - Decision: Attach smaller root `2` below root `0` and add its
      size.
    - After: `parentByItem=[0:0,1:0,2:0]`,
      `numItemsByRoot=[3,1,1]`,
      `nextUnion=none`.
    - Known correct: `find(0)` and `find(2)` both return `0`, so
      all added connections are represented.

Unioning already-connected items changes nothing and can signal that a
new undirected edge would form a cycle.

Path compression becomes visible with four items. After unions
`0-1`, `2-3`, and `0-2`, a valid parent state is
`parentByItem=[0:0,1:0,2:0,3:2]`.

Initial state: `item=3`,
`parentByItem=[0:0,1:0,2:0,3:2]`, `root=unknown`.

1.  Find and compress item `3`
    - Decision: Follow `3 -> 2 -> 0`; redirect `3` to root `0`.
    - After: `item=0`, `parentByItem=[0:0,1:0,2:0,3:0]`,
      `root=0`.
    - Known correct: Group membership is unchanged, while the next
      `find(3)` reaches its leader in one link.

###### Why it works

Every parent link stays inside one connected group, and every group
has one self-parenting leader. This starts true because each item is
its own group. A union first finds two leaders. Equal leaders require
no change; different leaders belong to groups joined by the new edge,
so pointing one leader to the other merges exactly those groups. It
cannot create a cycle because a leader has no parent inside the other
tree.

A `find` follows a finite parent tree and therefore reaches its
leader. Two items are connected exactly when they reach the same
leader. Path compression redirects an item only to that same leader,
and union by size changes only which leader represents the merged
group. Both keep membership unchanged while shortening or limiting
later searches.

###### Go example

`newDisjointSet` requires `numItems >= 0`. Validate every item against
the set's length before calling `find` or `union`.
For the trace, create `set := newDisjointSet(3)`, then call
`set.union(0, 1)` and `set.union(1, 2)`. Both return `true`, and
`set.find(0) == set.find(2)` is true.

```go
type disjointSet struct {
	parentByItem   []int
	numItemsByRoot []int
}

func newDisjointSet(numItems int) *disjointSet {
	// Initially every element is a one-member tree whose root is itself.
	set := &disjointSet{
		parentByItem:   make([]int, numItems),
		numItemsByRoot: make([]int, numItems),
	}
	for item := range set.parentByItem {
		// Self-parenting makes item its root; one records its single
		// member.
		set.parentByItem[item], set.numItemsByRoot[item] = item, 1
	}
	return set
}
func (set *disjointSet) find(item int) int {
	for item != set.parentByItem[item] {
		// Point the item to its grandparent while walking upward,
		// shortening future searches for the same component root.
		set.parentByItem[item] = set.parentByItem[set.parentByItem[item]]
		item = set.parentByItem[item]
	}
	return item
}
func (set *disjointSet) union(firstItem, secondItem int) bool {
	// Only roots represent groups; equal roots mean the elements are
	// connected.
	firstRoot := set.find(firstItem)
	secondRoot := set.find(secondItem)
	if firstRoot == secondRoot {
		// They already share a component, so no structural change is
		// needed.
		return false
	}
	if set.numItemsByRoot[firstRoot] < set.numItemsByRoot[secondRoot] {
		// Attach the smaller tree below the larger to keep paths shallow.
		firstRoot, secondRoot = secondRoot, firstRoot
	}
	set.parentByItem[secondRoot] = firstRoot
	// Only firstRoot keeps a meaningful component size after the
	// attachment.
	set.numItemsByRoot[firstRoot] += set.numItemsByRoot[secondRoot]
	return true
}
```

`union` reports whether two previously separate groups were joined.

###### Work and memory

With both improvements, a long sequence of operations is almost
constant time per operation in practice. Spread across the whole
sequence, which is what **amortized** means here, the formal bound is
`O(alpha(n))`, where `alpha` is a mathematical function called the
inverse Ackermann function. You do not need to calculate it for
implementation work: it grows so slowly that it remains below a tiny
constant for any realistic input. The parent and size arrays use
`O(n)` memory.

###### When to use it

Connections are added over time, and queries ask whether two items are
already connected, whether a new connection creates a loop, or how
many separate groups remain.

###### When not to use it

You need the actual route, the order in which items were visited, or
the shortest path. Knowing only each item's group does not preserve
that information.

###### Common mistakes and edge cases

> **Warning:** `find(-1)` or `find(n)` indexes outside the parent
> slice. Validate every item before calling the compact
> implementation.

An empty set has no valid item. Unioning an item with itself or two
already connected items returns `false` and must not increase the
recorded group size.

###### Related material

- [C.1.3.26 Minimum spanning tree](#c1326-minimum-spanning-tree) uses
  group leaders in Kruskal's algorithm.
- [B.5 Graphs and grids](#b5-graphs-and-grids) explains what
  connectivity preserves and what it loses.

<a id="c1326-minimum-spanning-tree"></a>

##### C.1.3.26 Connect at lowest cost (minimum spanning tree)

###### The problem it solves

Given locations and optional two-way connections with costs, choose a
set of connections that forms a tree across every location. Among all
such **spanning trees**, return one with the smallest total cost. A
tree is connected and contains no cycle.

###### How it works

Suppose villages need roads connecting everyone, but the total
building bill should be as small as possible. A **minimum spanning
tree** chooses that cheapest connected network without a cycle. For
`V` connected villages it uses `V-1` roads. It does not minimize the
route from one chosen village.

###### Step-by-step examples

For `A-B=1`, `B-C=2`, `A-C=3`, `C-D=4`, and `A-D=10`, expect
total cost `7`.

Initial state: `edges=[A-D:10,A-C:3,C-D:4,A-B:1,B-C:2]`,
`next=unsorted`, `groups={{A},{B},{C},{D}}`, `accepted=[]`,
`total=0`.

1.  Sort the edges
    - Decision: Sort from smallest to largest cost.
    - After: `edges=[A-B:1,B-C:2,A-C:3,C-D:4,A-D:10]`,
      `next=A-B`, `groups={{A},{B},{C},{D}}`, `accepted=[]`,
      `total=0`.
    - Known correct: Every later unexamined edge costs at least as
      much as the current edge.
2.  Inspect `A-B=1`
    - Decision: Its ends are in different groups, so accept it and
      join their groups.
    - After: `next=B-C`, `groups={{A,B},{C},{D}}`,
      `accepted=[A-B]`, `total=1`.
    - Known correct: The accepted forest has no loop.
3.  Inspect `B-C=2`
    - Decision: Its ends are in different groups, so accept it.
    - After: `next=A-C`, `groups={{A,B,C},{D}}`,
      `accepted=[A-B,B-C]`, `total=3`.
    - Known correct: The accepted edges connect three vertices
      without a cycle; `D` still needs a connection.
4.  Inspect `A-C=3`
    - Decision: Both ends already belong to `{A,B,C}`, so reject
      this cycle-forming edge.
    - After: `next=C-D`, `groups={{A,B,C},{D}}`,
      `accepted=[A-B,B-C]`, `total=3`.
    - Known correct: Rejecting the edge preserves the same connected
      groups and keeps the accepted forest acyclic.
5.  Inspect `C-D=4`
    - Decision: Its ends are in different groups, so accept it.
    - After: `next=A-D`, `groups={{A,B,C,D}}`,
      `accepted=[A-B,B-C,C-D]`, `total=7`.
    - Known correct: Four vertices are connected by `V-1=3` edges,
      so the result is a tree.
6.  Stop
    - Decision: `V-1` edges are accepted; skip `A-D=10`.
    - After: `next=done`, `groups={{A,B,C,D}}`,
      `accepted=[A-B,B-C,C-D]`, `total=7`.
    - Known correct: The accepted edges already form a spanning tree;
      any additional edge would create a cycle.

Focused failure case: with only `A-B=1` and isolated vertex `C`,
the scan accepts one edge, fewer than `V-1=2`. The expected result is
failure plus the forest edge `A-B`, not a spanning tree.

###### Why it works

Kruskal never creates a loop because it rejects a connection whose
ends are already in the same group. For an accepted edge, separate
one endpoint's current group from every vertex outside it. Any final
tree needs an edge across that split. Every cheaper edge was examined
already; a usable one would have joined the groups. The accepted edge
is therefore a cheapest crossing edge and can replace another tree's
crossing edge without raising cost. This is the **cut property**.

Each acceptance joins two groups, so `V - 1` accepted edges leave one
connected tree. If the scan cannot reach that count, some vertices
cannot be connected to the others.

On a disconnected graph, no single spanning tree exists. Kruskal
instead produces a **minimum spanning forest**, one minimum-cost tree
for each separate connected group.

###### Go example

This fragment uses the disjoint-set code above and requires
`numVertices >= 0` and valid endpoints. It sorts `edges` in place;
copy the slice first when its order must survive.
For the trace, map `A=0`, `B=1`, `C=2`, and `D=3`. Pass
`numVertices=4` and edges `{0,3,10}`, `{0,2,3}`, `{2,3,4}`,
`{0,1,1}`, and `{1,2,2}`. `kruskal` returns `(7, true)`.

```go
import "sort"

type mstEdge struct{ from, to, cost int }

func kruskal(numVertices int, edges []mstEdge) (int, bool) {
	// Sort the supplied edge slice in place, changing the caller's
	// order. Considering edges from cheapest to most expensive makes
	// each accepted edge the cheapest way to connect two currently
	// separate components.
	sort.Slice(edges, func(leftIndex, rightIndex int) bool {
		return edges[leftIndex].cost < edges[rightIndex].cost
	})
	groups := newDisjointSet(numVertices)
	totalCost := 0
	numEdgesUsed := 0
	for _, edge := range edges {
		if groups.union(edge.from, edge.to) {
			// union rejects cycle-forming edges; accepted edges grow the
			// forest.
			totalCost += edge.cost
			numEdgesUsed++
			if numEdgesUsed == numVertices-1 {
				break
			}
		}
	}
	// A nonempty spanning tree must contain exactly numVertices-1
	// accepted edges.
	return totalCost, numVertices == 0 || numEdgesUsed == numVertices-1
}
```

###### Work and memory

Kruskal takes `O(E log E)` time to sort and scan `E` edges. The
disjoint set uses `O(V)` auxiliary memory for `V` vertices. The input
edge list itself uses `O(E)` memory; sorting it in place avoids an
additional edge-list copy.

###### When to use it

Choose undirected weighted edges that connect every vertex with
minimum total edge cost and no cycle.

###### When not to use it

Do not use it when the goal is the cheapest route from one starting
vertex rather than the cheapest total tree. It also does not apply
when directed edges cannot be treated as two-way connections.

###### Common mistakes and edge cases

> **Warning:** Accepting an edge whose endpoints already share a
> leader creates a cycle. In the example, `A-C=3` must be rejected
> even though vertex `D` is not connected yet. Only an edge crossing
> from `{A,B,C}` to `{D}` can help complete the tree.

Validate nonnegative vertex count and every endpoint. Copy the edge
slice when the caller's order must survive, and use checked or wider
arithmetic when the selected edge costs can overflow `int`. This guide
treats an empty graph as successfully spanned with cost zero.

###### Related material

- [C.1.3.25 Disjoint-set union](#c1325-disjoint-set-union) supports
  Kruskal.
- [C.1.3.23 Dijkstra shortest paths](#c1323-dijkstra-shortest-paths)
  minimizes routes from one start, not total network cost.

<a id="c1327-heap-or-priority-queue"></a>

##### C.1.3.27 Remove urgent items (heap or priority queue)

###### The problem it solves

Given items that continue to arrive, repeatedly inspect or remove the
item with the smallest, largest, or most urgent key. Fully sorting the
collection after every insertion repeats work, while scanning all
items for every removal makes each priority operation too expensive.

###### How it works

An emergency room serves the most urgent waiting patient, not always
the earliest arrival. A **priority queue** follows that rule. A
**heap** keeps the next patient easy to find without fully sorting the
whole waiting room after every change.

A min-heap stores a tree in a slice. The root is at index `0`. For an
item at index `i`, its parent is at `(i-1)/2`; its children are at
`2*i+1` and `2*i+2` when those indexes exist. Every parent is no
greater than its children, so the root is always the minimum. Other
slice positions are not globally sorted.

###### Step-by-step examples

Insert job priorities `5`, `2`, and `8` into a min-heap, where a
smaller number means a more urgent job. Expect removals `2,5,8`. The
slice displays the tree one level at a time, with its root first.

Initial state: `heap=[]`, `next=insert(5)`, `removed=[]`.

1.  Insert `5`
    - Decision: Append `5`; it has no parent to compare with.
    - After: `heap=[5]`, `next=insert(2)`, `removed=[]`.
    - Known correct: The only value is the minimum.
2.  Insert `2`
    - Decision: Append `2`, then swap it with parent `5` because
      `2` is smaller.
    - After: `heap=[2,5]`, `next=insert(8)`, `removed=[]`.
    - Known correct: The parent is no greater than its child.
3.  Insert `8`
    - Decision: Append `8`; its parent `2` is already smaller.
    - After: `heap=[2,5,8]`, `next=remove`, `removed=[]`.
    - Known correct: Root `2` is smaller than both children.
4.  Remove `2`
    - Decision: Save `2`, move `8` to the root, then swap it with
      smaller child `5`.
    - After: `heap=[5,8]`, `next=remove`, `removed=[2]`.
    - Known correct: The removed value was the minimum, and root
      `5` is the new minimum.
5.  Remove `5`
    - Decision: Save `5` and move final value `8` to the root.
    - After: `heap=[8]`, `next=remove`, `removed=[2,5]`.
    - Known correct: The only remaining value is the new minimum.
6.  Remove `8`
    - Decision: Remove the sole value.
    - After: `heap=[]`, `next=done`, `removed=[2,5,8]`.
    - Known correct: Every removal returned the smallest value
      stored at that time.

Before reading or removing from an empty heap, decide how the function
will report that no value exists.

###### Why it works

The heap contains exactly the items still under consideration, and
in a min-heap every parent value is no greater than either child. Any
path from the root therefore cannot reach a value smaller than the
root. Push and pop repair this parent-child rule along a single path
between a leaf and the root. The root is then the required next item
again. Each repair crosses at most the heap's height and therefore
terminates.

###### Go example

These functions return the possibly changed slice, so callers must
keep that returned value. `popMinHeap` reports `false` for an empty
heap.
For the trace, start with `heap := []int{}` and replace it with each
returned slice after pushes of `5`, `2`, and `8`. Three calls to
`popMinHeap` then return minimum values `2`, `5`, and `8`; its slice
result is the trace's updated `heap`.

```go
func pushMinHeap(values []int, value int) []int {
	// Append as the next leaf, then move the new value toward the root
	// while it is smaller than its parent.
	values = append(values, value)
	child := len(values) - 1
	for child > 0 {
		parent := (child - 1) / 2
		if values[parent] <= values[child] {
			break
		}
		values[parent], values[child] = values[child], values[parent]
		child = parent
	}
	return values
}

func popMinHeap(values []int) (int, []int, bool) {
	if len(values) == 0 {
		return 0, values, false
	}
	// Save the root, move the last leaf into its position, and remove
	// the now-duplicate final slot.
	minimum := values[0]
	lastIndex := len(values) - 1
	values[0] = values[lastIndex]
	values = values[:lastIndex]

	// Move the replacement down through its smaller child until the
	// parent-child rule is restored along the whole repair path.
	parent := 0
	for {
		leftChild := 2*parent + 1
		if leftChild >= len(values) {
			break
		}
		smallerChild := leftChild
		rightChild := leftChild + 1
		if rightChild < len(values) &&
			values[rightChild] < values[leftChild] {
			smallerChild = rightChild
		}
		if values[parent] <= values[smallerChild] {
			break
		}
		values[parent], values[smallerChild] =
			values[smallerChild], values[parent]
		parent = smallerChild
	}
	return minimum, values, true
}
```

###### Work and memory

A bottom-up build can turn `n` existing items into a heap in `O(n)`
time. Inserting all `n` items with the shown `pushMinHeap` instead
takes `O(n log n)` time. One push or pop takes `O(log n)` time;
reading a nonempty root takes `O(1)`. The stored heap uses `O(n)`
memory, and one operation uses `O(1)` auxiliary memory apart from any
slice growth.

A heap does not provide constant-time search or removal for an
arbitrary value. Only the root has the heap's priority guarantee.

###### When to use it

You repeatedly need the smallest or largest item, a requested number
of the best items, or the next event in priority order.

###### When not to use it

You need the complete final order of an unchanging set of items;
sorting is usually simpler. A heap also does not make it fast to find
any item other than the smallest, largest, or highest-priority one it
was designed to expose.

###### Common mistakes and edge cases

> **Warning:** The heap's backing slice is not globally sorted. After
> inserting `1,4,3`, only the root `1` is guaranteed; reading the
> remaining slice as the next output order can incorrectly report `4`
> before `3`.

Check emptiness before peek or pop. Define tie-breaking when equal
priorities must be stable.

###### Related material

- [D.5 Priority queues](#d5-priority-queues) explains the Go
  implementation pattern.
- [C.1.3.5 Sorting and selection](#c135-sorting-and-selection) is
  simpler when every item must be ordered once.
- [C.1.3.29 Monotonic deque](#c1329-monotonic-deque) is linear when
  candidates expire in sequence order.

<a id="c1328-monotonic-stack"></a>

##### C.1.3.28 Resolve nearest values (monotonic stack)

###### The problem it solves

Given a sequence, return for every position the nearest later or
earlier value that is greater or smaller under a stated strictness
rule. A direct search restarts from every position; the structure must
retain only unresolved candidates that the current value has not made
irrelevant.

###### How it works

Keep a pile of earlier days still waiting for a warmer day. When a new
temperature answers days on top, remove them; unresolved temperatures
remain ordered. For a next-greater query, their values stay
non-increasing from the stack's bottom to its top. This **monotonic
stack** rule also finds the nearest larger or smaller value in other
sequences. The result contract is
`result[index] -> nearest later greater value`, or `none` while that
answer is unresolved or absent.

###### Step-by-step examples

For `[2,1,4]`, find each next greater value; expect `[4,4,none]`.

Initial state: `index=0`, `stack=[]`, `result=[none,none,none]`.

1.  Read value `2`
    - Decision: No earlier position can be resolved, so push index
      `0`.
    - After: `index=1`, `stack=[0]`, `result=[none,none,none]`.
    - Known correct: Index `0` still awaits a greater later value.
2.  Read value `1`
    - Decision: `1` is not greater than `2`, so push index `1`.
    - After: `index=2`, `stack=[0,1]`, `result=[none,none,none]`.
    - Known correct: Neither stored position has seen a greater
      later value.
3.  Resolve index `1` with `4`
    - Decision: Pop index `1` and set its result to `4`.
    - After: `index=2`, `stack=[0]`, `result=[none,4,none]`.
    - Known correct: `4` is the first greater value after index
      `1`.
4.  Resolve index `0` with `4`
    - Decision: Pop index `0` and set its result to `4`.
    - After: `index=2`, `stack=[]`, `result=[4,4,none]`.
    - Known correct: Index `1` was not greater than `2`, so `4` is
      also the first valid answer for index `0`.
5.  Store index `2` and finish
    - Decision: Push index `2`; after the input ends, leave its
      result as `none`.
    - After: `index=done`, `stack=[2]`, `result=[4,4,none]`.
    - Known correct: No later value exists for the last position.

With equal values, choose `<` or `<=` according to whether "greater"
is strict; the wrong rule changes duplicate answers.

###### Why it works

Every stored item is still waiting for an answer, and the stored
values remain sorted in one direction. An item is removed only when
the current value is the first later value that answers its question.
Each input position is pushed once and popped at most once, so the
finite scan and every inner removal loop finish.

###### Go example

The code uses `-1` for the trace's `none`. Therefore
`nextGreater([]int{2, 1, 4})` returns `[]int{4, 4, -1}`. The trace's
`index` is the code's `currentIndex`; `stack` and `result` keep the
same roles.

```go
func nextGreater(values []int) []int {
	// -1 is the answer for indices that never encounter a greater value.
	result := make([]int, len(values))
	for resultIndex := range result {
		result[resultIndex] = -1
	}
	// stack holds unresolved indices whose values are non-increasing
	// from bottom to top; equal values can coexist because neither is
	// greater than the other.
	stack := make([]int, 0, len(values))
	for currentIndex, value := range values {
		for len(stack) > 0 && values[stack[len(stack)-1]] < value {
			// value is the first greater value to the right of every popped
			// index.
			unresolvedIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[unresolvedIndex] = value
		}
		stack = append(stack, currentIndex)
	}
	return result
}
```

This strict comparison leaves equal values waiting; change the rule
only when the problem asks for greater-or-equal.

###### Work and memory

Each position is added once and removed at most once, so the whole
scan takes `O(n)` time even though one step may remove several items.
The stack can hold `O(n)` positions.

###### When to use it

The prompt asks for the next greater or smaller value, the nearest
earlier or later value that becomes larger or smaller, or how long a
value remains the minimum or maximum.

###### When not to use it

The answer depends on any earlier value, not specifically the nearest
one that meets the condition, or when removing an item would lose
information that a later value still needs.

###### Common mistakes and edge cases

> **Warning:** Using `<=` instead of `<` for a strictly greater query
> pops an equal value too early. For `[2,2,3]`, the first `2` must
> wait for `3`, not use the second `2` as its answer.

The example uses `-1` for no answer; if `-1` is a valid input value,
return indexes, booleans, or another unambiguous representation. Empty
input returns an empty result.

###### Related material

- [C.1.3.13 Stack](#c1313-stack) explains stack behavior.
- [C.1.3.29 Monotonic deque](#c1329-monotonic-deque) also removes
  older values made useless by newer ones, and it can remove expired
  positions from the front.

<a id="c1329-monotonic-deque"></a>

##### C.1.3.29 Keep window extremes (monotonic deque)

###### The problem it solves

Given a sequence and a fixed window length, return the minimum or
maximum for every consecutive window. Neighboring windows overlap, so
rescanning all their values repeats work; expired positions and values
that can never become the extreme must be removed as the window moves.

###### How it works

Picture candidates waiting to be the highest score in a moving time
window. Remove candidates whose time expired and those beaten by a
newer score. A **monotonic deque** keeps the remaining candidates
in decreasing value order, with the current maximum at the front.

###### Step-by-step examples

For `[4,2,1,3]` with window size `3`, expect maxima `[4,3]`.

Initial state: `index=0`, `deque=[]`, `output=[]`.

1.  Read value `4`
    - Decision: Append index `0` at the back.
    - After: `index=1`, `deque=[0]`, `output=[]`.
    - Known correct: Index `0` is the only maximum candidate.
2.  Read value `2`
    - Decision: Keep index `0` because value `4` is larger, then
      append index `1`.
    - After: `index=2`, `deque=[0,1]`, `output=[]`.
    - Known correct: Candidate values decrease from front to back.
3.  Read value `1`
    - Decision: Keep both earlier candidates, append index `2`, and
      report front value `4`.
    - After: `index=3`, `deque=[0,1,2]`, `output=[4]`.
    - Known correct: The first full window is indexes `0..2`, and
      its maximum is front value `4`.
4.  Check expiration at index `3`
    - Decision: Remove front index `0` because it is outside the new
      window `1..3`.
    - After: `index=3`, `deque=[1,2]`, `output=[4]`.
    - Known correct: Every remaining candidate belongs to the new
      window; expired value `4` cannot be reported again.
5.  Prune with value `3`
    - Decision: Remove index `2`, then index `1`, because both values
      are no larger than newer value `3`.
    - After: `index=3`, `deque=[]`, `output=[4]`.
    - Known correct: Both removed positions expire earlier and can
      never beat `3` in a later shared window.
6.  Append index `3` and report
    - Decision: Append index `3`; report its front value `3`.
    - After: `index=done`, `deque=[3]`, `output=[4,3]`.
    - Known correct: Front value `3` is the maximum of indexes
      `1..3`.

Store positions, not only values, so duplicates and expiration are
distinguishable. Window size zero is invalid and size one returns
every input value.

###### Why it works

Positions remain in arrival order, and their values stay ordered in
one direction. Every removed position has either left the window or
has been replaced by a later value that is at least as useful and will
remain available longer. Each position enters once and can leave only
once, so all removal loops and the finite outer scan terminate.

###### Go example

Calling `windowMaximum([]int{4, 2, 1, 3}, 3)` returns
`[]int{4, 3}`. The trace calls the second input the window size; the
code names it `width`. Its `deque` is the same index collection, and
its `result` is the trace's `output`.

```go
func windowMaximum(values []int, width int) []int {
	if width <= 0 || width > len(values) {
		return nil
	}
	// deque stores candidate indices in decreasing value order; the
	// front is the maximum for the current window.
	deque := make([]int, 0, width)
	result := make(
		[]int,
		0,
		len(values)-width+1,
	)
	for currentIndex, value := range values {
		if len(deque) > 0 && deque[0] <= currentIndex-width {
			// The front index has left the window and can no longer be an
			// answer.
			deque = deque[1:]
		}
		for len(deque) > 0 && values[deque[len(deque)-1]] <= value {
			// Newer value dominates smaller or equal candidates for every
			// future window in which both would appear.
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, currentIndex)
		if currentIndex+1 >= width {
			// A complete window exists, and its maximum is the front
			// candidate.
			result = append(result, values[deque[0]])
		}
	}
	return result
}
```

Indices preserve both expiration time and duplicate identity.

###### Work and memory

Each position enters and leaves the deque at most once, so all windows
are processed in `O(n)` time with `O(window size)` memory. Store
positions rather than values when expiration depends on location or
duplicate values exist.

###### When to use it

Report the minimum or maximum of every moving window without
rescanning each window. Positions must leave in sequence order, and a
newer value must be able to make an older one permanently unnecessary.

###### When not to use it

Use a heap instead when priorities change independently or items stop
being valid in no predictable order.

###### Common mistakes and edge cases

> **Warning:** Storing only values cannot identify which duplicate
> expired. With `[5,5,1]` and width `2`, positions distinguish the old
> `5` leaving from the newer `5` that still belongs to the window.

Reject width zero, a negative width, or a width larger than the input.
Using `<=` when pruning equal values keeps the newer duplicate, which
is safe because it expires later.

###### Related material

- [B.3 Queues and deques](#b3-queues-and-deques) explains deque
  operations.
- [C.1.3.14 Sliding window](#c1314-sliding-window) explains boundary
  movement.
- [C.1.3.27 Heap or priority queue](#c1327-heap-or-priority-queue)
  handles priorities that do not expire in sequence order.

<a id="c1330-intervals"></a>

##### C.1.3.30 Put overlapping ranges together (intervals)

###### The problem it solves

Given ranges with inclusive endpoints, combine every overlap and
return the smallest list of ranges with the same coverage. Ranges
that overlap may be far apart in the original input, so sort them by
start before deciding which ones interact.

###### How it works

Calendar bookings have a start and end. Sort them by start time, and
overlaps appear beside each other so they can be joined. Each covered
start-to-end range is an **interval**. This entry includes both ends,
so `[1,3]` includes its endpoints `1` and `3`.

###### Step-by-step examples

Merge `[8,10]`, `[1,4]`, `[2,3]`, and `[3,6]`; expect
`[[1,6],[8,10]]`.

Initial state: `intervals=[[8,10],[1,4],[2,3],[3,6]]`,
`next=unsorted`, `result=[]`.

1.  Sort by start
    - Decision: Put the smallest start first; seed the result.
    - After: `intervals=[[1,4],[2,3],[3,6],[8,10]]`,
      `next=[2,3]`, `result=[[1,4]]`.
    - Known correct: A possible overlap with the current interval
      now appears before any later gap.
2.  Read `[2,3]`
    - Decision: It lies inside `[1,4]`, so keep the existing end.
    - After: `next=[3,6]`, `result=[[1,4]]`.
    - Known correct: The contained interval adds no new covered
      point.
3.  Read `[3,6]`
    - Decision: Inclusive start `3` is no greater than end `4`, so
      extend the current end to `6`.
    - After: `next=[8,10]`, `result=[[1,6]]`.
    - Known correct: The result covers exactly the first three
      sorted intervals as one maximal range.
4.  Read `[8,10]`
    - Decision: Start `8` is greater than end `6`, so append a new
      range.
    - After: `next=done`, `result=[[1,6],[8,10]]`.
    - Known correct: The gap `6..8` proves the two covered ranges
      cannot be merged.

Whether `[1,2]` and `[2,3]` overlap depends on whether endpoints are
inclusive or half-open; decide before comparing.

###### Why it works

After each step, the result is exactly the union of the intervals
already processed, stored as separate maximal ranges. The first
interval makes this true. If the next interval overlaps the last
range, extending that range's end preserves the same union. If there
is a gap, sorting proves that no later interval can overlap the last
range, so it is safe to start a new one. Each step consumes one
interval. When the finite scan ends, the result is the complete merged
union.

###### Go example

This version requires inclusive endpoints and `start <= end` for
every interval. It sorts the input slice in place.
For the trace, call `mergeIntervals` with intervals `{8,10}`,
`{1,4}`, `{2,3}`, and `{3,6}`. It returns `{1,6}` and `{8,10}`.

```go
import "sort"

type interval struct{ start, end int }

func mergeIntervals(intervals []interval) []interval {
	if len(intervals) == 0 {
		return nil
	}
	// Sort the supplied slice in place by start, changing the caller's
	// order. Every later interval then begins no earlier than the
	// current merged interval.
	sort.Slice(intervals, func(leftIndex, rightIndex int) bool {
		return intervals[leftIndex].start < intervals[rightIndex].start
	})
	result := []interval{intervals[0]}
	for _, next := range intervals[1:] {
		last := &result[len(result)-1]
		if next.start <= last.end {
			// Overlap means only the farther end can enlarge the merged
			// interval.
			if next.end > last.end {
				last.end = next.end
			}
		} else {
			// A gap proves next cannot overlap last or any earlier interval.
			result = append(result, next)
		}
	}
	return result
}
```

###### Work and memory

Sorting `n` intervals takes `O(n log n)` time, and the following scan
takes `O(n)`. Sorting in place uses no separate interval copy, but the
merged result can contain `O(n)` intervals.

###### When to use it

The input contains ranges, and the output must combine overlapping
ranges into their union.

###### When not to use it

Do not use this offline sort-and-scan method when ranges change and
each change must be queried immediately. It is also unnecessary when
the problem is about unrelated individual points rather than ranges.

###### Common mistakes and edge cases

Validate that every start is no greater than its end. The problem must
define whether `[1, 2]` and `[2, 3]` overlap. If both ending points
are included, the ranges share `2`. Under the half-open rule
`[start,end)`, the first range excludes `2`, so they do not overlap.

###### Related material

- [C.1.3.5 Sorting and
  selection](#c135-sorting-and-selection) explains
  ordering.
- [C.1.3.31 Sweep line](#c1331-sweep-line) counts simultaneous ranges.

<a id="c1331-sweep-line"></a>

##### C.1.3.31 Process change points (sweep line)

###### The problem it solves

Given bookings, jobs, or occupied ranges, return the largest number
active at once. Time may be continuous or span a huge coordinate
range, so inspect only start and end events where the active count can
change rather than checking every possible moment.

###### How it works

A room becomes busy only when a booking starts and becomes freer when
one ends. Sort those **events** and update the busy count only at
those times. Moving through ordered change points is a **sweep line**.
The lookup `changeByTime[time] -> net change in active bookings at
time` combines simultaneous starts and ends before the sweep.

###### Step-by-step examples

One booking uses a room from time `1` until time `4`. Another uses a
room from time `3` until time `5`. The input is `[1,4)` and `[3,5)`,
and the expected maximum is `2` busy rooms.

Initial state: `changes={}`, `nextBooking=[1,4)`, `time=unset`,
`activeRooms=0`, `maximumRooms=0`.

1.  Record booking `[1,4)`
    - Decision: Add `+1` at `1` and `-1` at `4`.
    - After: `changes={1:+1,4:-1}`, `nextBooking=[3,5)`,
      `time=unset`, `activeRooms=0`, `maximumRooms=0`.
    - Known correct: The first booking contributes one busy room
      from `1` until, but not including, `4`.
2.  Record booking `[3,5)`
    - Decision: Add `+1` at `3` and `-1` at `5`.
    - After: `changes={1:+1,3:+1,4:-1,5:-1}`, `nextBooking=none`,
      `time=1`, `activeRooms=0`, `maximumRooms=0`.
    - Known correct: Every start and end is now represented.
3.  Visit time `1`
    - Decision: Apply change `+1` and update the maximum.
    - After: `changes={1:+1,3:+1,4:-1,5:-1}`, `nextBooking=none`,
      `time=3`, `activeRooms=1`, `maximumRooms=1`.
    - Known correct: One booking is active after time `1`.
4.  Visit time `3`
    - Decision: Apply change `+1` and update the maximum.
    - After: `changes={1:+1,3:+1,4:-1,5:-1}`, `nextBooking=none`,
      `time=4`, `activeRooms=2`, `maximumRooms=2`.
    - Known correct: Both bookings are active after time `3`.
5.  Visit time `4`
    - Decision: Apply change `-1` because the first booking
      has ended.
    - After: `changes={1:+1,3:+1,4:-1,5:-1}`, `nextBooking=none`,
      `time=5`, `activeRooms=1`, `maximumRooms=2`.
    - Known correct: Only the second booking remains active.
6.  Visit time `5`
    - Decision: Apply change `-1`.
    - After: `changes={1:+1,3:+1,4:-1,5:-1}`, `nextBooking=none`,
      `time=done`, `activeRooms=0`, `maximumRooms=2`.
    - Known correct: No booking remains, and the largest count
      seen was two.

For a boundary example, use bookings `[1,4)` and `[4,5)`. The first
ends exactly when the second starts, so the expected maximum is `1`.
At time `4`, the changes `-1` and `+1` add to zero. The running count
therefore stays at one instead of briefly and incorrectly becoming
two.

###### Why it works

Before the first change, no booking is active. At each change time,
every booking that starts contributes one and every booking that ends
removes one. After their combined change is applied, `activeRooms`
therefore equals the number of bookings active immediately after that
time. Nothing starts or ends between two recorded times, so the count
cannot change there.

The loop visits every distinct change time once and then finishes.
Because `maximumRooms` remembers the largest correct active count seen
at any time, it is the largest number of simultaneous bookings.

###### Go example

For the trace, call `maximumSimultaneousBookings` with bookings
`{start:1,end:4}` and `{start:3,end:5}`. It returns `(2, true)`.
The trace's `changes` map is the code's `changeByTime`; the remaining
state names match directly.

```go
import "sort"

type booking struct {
	start int
	end   int
}

func maximumSimultaneousBookings(bookings []booking) (int, bool) {
	// Combine all room-count changes that happen at the same time. This
	// makes an ending and a start at one time cancel each other.
	changeByTime := make(map[int]int, len(bookings)*2)
	for _, currentBooking := range bookings {
		// A half-open booking must contain at least one moment. Report an
		// invalid range instead of silently changing its meaning.
		if currentBooking.start >= currentBooking.end {
			return 0, false
		}
		changeByTime[currentBooking.start]++
		changeByTime[currentBooking.end]--
	}

	// Map iteration has no promised order, so collect and sort the
	// distinct times before sweeping through them.
	times := make([]int, 0, len(changeByTime))
	for changeTime := range changeByTime {
		times = append(times, changeTime)
	}
	sort.Ints(times)

	// activeRooms is correct immediately after each visited time.
	// maximumRooms remembers the largest such count.
	activeRooms, maximumRooms := 0, 0
	for _, changeTime := range times {
		activeRooms += changeByTime[changeTime]
		if activeRooms > maximumRooms {
			maximumRooms = activeRooms
		}
	}
	return maximumRooms, true
}
```

The function returns `false` when any booking has `start >= end`.
Empty input is valid and returns `(0, true)`. The function does not
reorder the caller's bookings.

###### Work and memory

Let `n` be the number of bookings. Recording their changes takes
`O(n)` time. There are at most `2n` distinct change times, so sorting
them takes `O(n log n)` time. Visiting the sorted times takes
`O(n)` more time. The map and time slice together use `O(n)` extra
memory.

###### When to use it

Use this method when the answer changes only at known starts and ends.
Typical requests ask for the maximum number of simultaneous bookings,
people in a building, running jobs, or overlapping ranges.

###### When not to use it

If every possible position belongs to a small fixed integer range, a
difference array may be simpler. If bookings arrive one at a time and
each one needs an immediate answer before all bookings are known, use
a structure that maintains ordered changes instead of sorting once at
the end.

###### Common mistakes and edge cases

> **Warning:** Applying a start before an end at the same time can
> invent an overlap. For `[1,4)` and `[4,5)`, combine `-1` and `+1`
> at time `4`, or otherwise make the end take effect first.

Reject a booking whose start is not earlier than its end. Decide
whether end times are included or excluded before creating changes;
the rule changes how touching bookings behave.

###### Related material

- [C.1.3.30 Intervals](#c1330-intervals) explains start and end rules.
- [C.1.3.16 Difference arrays](#c1316-difference-arrays) records
  changes when every possible position belongs to a small range.

<a id="c1332-fast-and-slow-pointers"></a>

##### C.1.3.32 Find a linked-list middle (fast and slow pointers)

###### The problem it solves

Given a one-way linked list, return its middle node without first
counting all nodes or storing them elsewhere. For an even number of
nodes, this entry returns the second of the two middle nodes.

###### How it works

Put two runners at the list's first node. Move `slow` one link and
`fast` two links per iteration. When `fast` reaches the end, `slow`
has moved half as far and points at the middle. This is the
**fast-and-slow pointer** method.

###### Step-by-step examples

For `1->2->3->4->nil`, return the second middle `3`.

Initial state: `slow=1`, `fast=1`, `result=pending`.

1.  Move both pointers
    - Decision: Move `slow` one link to `2`; move `fast` two links
      to `3`.
    - After: `slow=2`, `fast=3`, `result=pending`.
    - Known correct: `fast` has moved twice as many links.
2.  Move both pointers again
    - Decision: Move `slow` one link to `3`; move `fast` through
      `4` to `nil`.
    - After: `slow=3`, `fast=nil`, `result=pending`.
    - Known correct: `slow` has moved two of the four links from
      the head position through the end.
3.  Stop
    - Decision: `fast` is `nil`, so return `slow`.
    - After: `slow=3`, `fast=nil`, `result=3`.
    - Known correct: Starting both pointers at the head selects
      the later of the two middle nodes for an even-length list.

For even lengths, a different loop condition can return the first
middle. Check both `fast` and the next link before a two-step move;
empty input returns no node.

###### Why it works

For the middle-node loop, after `t` iterations `slow` has moved `t`
links and `fast` has moved `2t`. When `fast` reaches the end, `slow`
has covered half the list; the loop condition chooses the later middle
for even lengths. Each iteration moves `fast` closer to the end by two
links, so a finite list makes the loop terminate.

###### Go example

```go
type listNode struct {
	value int
	next  *listNode
}

func secondMiddle(head *listNode) *listNode {
	// slow advances one node while fast advances two. When fast reaches
	// the end, slow has covered half the list.
	slow, fast := head, head
	for fast != nil && fast.next != nil {
		slow, fast = slow.next, fast.next.next
	}
	// Starting both at head selects the second middle for an even-length
	// list.
	return slow
}
```

The loop returns the later middle for even lengths and `nil` for an
empty list.

###### Work and memory

For `n` list nodes, each pointer makes at most `n` moves. The method
takes `O(n)` time and `O(1)` extra memory.

###### When to use it

You need the middle of a finite linked list without making one pass to
count nodes and a second pass to reach the middle.

###### When not to use it

The list may contain a cycle, or you need an arbitrary position rather
than the middle. On an indexed slice, `len(values)/2` directly selects
the second middle.

###### Common mistakes and edge cases

> **Warning:** Testing only `fast != nil` before `fast.next.next`
> panics on a one-node suffix. Require both `fast != nil` and
> `fast.next != nil`.

Empty input returns no node. State whether an even-length list needs
the first or second middle before choosing the loop condition.

###### Related material

- [B.2 Linked lists](#b2-linked-lists) explains linked nodes.
- [C.1.3.9 Two pointers](#c139-two-pointers) explains the broader
  coordinated-position method.

<a id="c1333-trie"></a>

##### C.1.3.33 Share stored word beginnings (trie)

###### The problem it solves

Given many strings, support insertions and queries for an exact string
or a shared beginning, also called a prefix. Repeatedly scanning every
stored string repeats prefix comparisons; share one path for common
leading symbols so a query follows only the symbols it contains.

###### How it works

An autocomplete menu can store `car`, `card`, and `care` along one
shared `car` path, then branch at the next letter. A **trie** is this
tree of shared word beginnings. Its root comes before any character,
and an end marker distinguishes a stored word from a mere prefix. At
each node, `children[byte] -> next node for that byte`; a missing key
means no stored word has that extended prefix.

###### Step-by-step examples

Insert `car` and `cat`; expect prefix `ca` to exist but full word `ca`
not to exist.

Initial state: `trie=root`, `nextOperation=insert(car)`,
`answers=[]`.

1.  Insert `car`
    - Decision: Create and follow children `c`, `a`, and `r` in
      order, then mark the `r` node as a word end.
    - After: `trie={ca->{r:word}}`, `nextOperation=insert(cat)`,
      `answers=[]`.
    - Known correct: Exact lookup for `car` can follow every byte
      and find its end marker.
2.  Insert `cat`
    - Decision: Reuse `c`,`a`; create `t` beside `r` and mark it.
    - After: `trie={ca->{r:word,t:word}}`,
      `nextOperation=prefix(ca)`, `answers=[]`.
    - Known correct: Both words share `ca`; their final bytes use
      separate marked nodes.

3.  Query prefix `ca`
    - Decision: Follow `c`, then `a`; do not require a
      word-end marker for a prefix query.
    - After: `trie={ca->{r:word,t:word}}`,
      `nextOperation=word(ca)`, `answers=[true]`.
    - Known correct: At least one stored word begins with `ca`.
4.  Query whole word `ca`
    - Decision: Follow the same path and require a marker on `a`.
    - After: `trie={ca->{r:word,t:word}}`, `nextOperation=done`,
      `answers=[true,false]`.
    - Known correct: `ca` is only a beginning, not an inserted
      complete word.

###### Why it works

The characters followed from the starting point always spell the
beginning represented by the current step. A word exists only if the
last step is marked as the end of a complete word. Insert and lookup
advance one input byte per loop iteration, so both finish after the
finite text ends or lookup encounters a missing child.

###### Go example

These methods require a non-nil root and one-byte alphabet semantics.
Use runes when the problem defines positions as Unicode code points.
For the trace, create `root := &trieNode{}`, insert `"car"` and
`"cat"`, then call `root.find("ca", false)` for the prefix query and
`root.find("ca", true)` for the whole-word query. They return `true`
and `false` in that order.

```go
type trieNode struct {
	children map[byte]*trieNode
	word     bool
}

func (root *trieNode) insert(word string) {
	node := root
	for byteIndex := range len(word) {
		// Allocate each sparse child map only when this prefix first needs
		// one.
		if node.children == nil {
			node.children = make(map[byte]*trieNode)
		}
		if node.children[word[byteIndex]] == nil {
			// Missing byte edges create new prefix nodes; existing edges are
			// shared with words that have the same prefix.
			node.children[word[byteIndex]] = &trieNode{}
		}
		node = node.children[word[byteIndex]]
	}
	// The terminal flag distinguishes a complete word from a prefix
	// alone.
	node.word = true
}

func (root *trieNode) find(text string, wholeWord bool) bool {
	node := root
	// Follow one byte edge per input position; a missing edge proves
	// absence.
	for byteIndex := range len(text) {
		node = node.children[text[byteIndex]]
		if node == nil {
			return false
		}
	}
	// Prefix lookup needs only the path; whole-word lookup also needs
	// the terminal.
	return !wholeWord || node.word
}
```

###### Work and memory

For a word containing `L` bytes, insertion or lookup takes `O(L)`
time. One insertion may create `O(L)` nodes. Across all words, memory
is proportional to the number of distinct stored prefix bytes.

###### When to use it

You repeatedly ask for words beginning with the same letters, check
dictionary words one character at a time, or want to stop as soon as a
beginning does not exist.

###### When not to use it

You only need to know whether a complete word exists. A hash set is
usually simpler and may use less memory.

###### Common mistakes and edge cases

> **Warning:** Treating a reached prefix as a complete word makes
> `find("ca", true)` succeed after inserting only `"car"`; require the
> end marker for a whole-word query.

The methods require a non-nil root. Decide whether the empty string is
allowed, and use runes or visible-symbol segmentation instead of bytes
when the problem requires those text units.

###### Related material

- [C.1.3.4 Hash map or set](#c134-hash-map-or-set) is simpler for
  exact whole-word membership.
- [D.6 Bytes, runes, and visible
  characters](#d6-bytes-runes-and-visible-characters) determines what
  one trie step represents.
- [B.4 Trees and binary-search
  trees](#b4-trees-and-binary-search-trees) defines tree paths.

#### C.1.4 Choosing among overlapping methods

##### The problem it solves

Given several methods that appear compatible with the input, choose
the one whose assumptions match the required output and update model.
Compare the exact work each method may skip, the state it must retain,
whether the data changes between questions, and whether it returns the
needed result rather than a related one.

##### How it works

Choose a tool by the job it must keep doing correctly, not because two
tools look alike. Compare each method's required promise; the first
promise that differs usually decides which method fits.

##### Step-by-step examples

A road problem asks for cheapest routes from `A`.

Initial state: `edgeCosts=allOne`, `method=unknown`.

1.  Select for equal costs
    - Decision: Choose BFS because first-in, first-out queue
      layers are also cost layers.
    - After: `edgeCosts=allOne`, `method=BFS`.
    - Known correct: First discovery has the fewest moves and
      therefore the lowest cost.

Initial state: `edgeCosts={1,4,9}`, `method=unknown`.

1.  Select for varying nonnegative costs
    - Decision: Choose Dijkstra and select the cheapest known
      route next.
    - After: `edgeCosts={1,4,9}`, `method=Dijkstra`.
    - Known correct: Costs never decrease after taking another
      edge, which lets Dijkstra finalize the cheapest selected route.

Initial state: `edgeCosts=contains(-3)`, `method=unknown`.

1.  Select for a negative cost
    - Decision: Reject Dijkstra and choose Bellman-Ford.
    - After: `edgeCosts=contains(-3)`, `method=BellmanFord`.
    - Known correct: A later negative edge may improve a route
      that looked finished earlier.

##### Why it works

The comparison keeps every method whose requirements still match the
problem. It removes a method only when one of those requirements
fails. Once a method can safely skip work and also fits the time and
memory limits, its own card explains why it is correct.

##### Go example

This comparison chooses an implementation rather than providing shared
runtime code. Follow the selected card's Go input rules and its link
to the complete implementation.

##### Work and memory

This short comparison is planning work. The final cost must still
include setup, the main processing, saved data, unfinished recursive
calls, and construction of the output.

##### When to use it

Use it when two named methods appear plausible after the quick
decision tree.

##### When not to use it

Do not compare labels before stating the input and output
requirements; the real method may be outside the apparent pair.

##### Common mistakes and edge cases

> **Warning:** Dijkstra is not "weighted BFS" when negative edges
> exist. A negative edge removes the proof that the cheapest selected
> distance is final.

Also ask whether all input exists at the start or arrives over time,
whether it may change, whether one or every answer is needed, and how
much work is required just to produce the output.

##### Related material

- [C.1.1 Ask what the input and output
  allow](#c11-ask-what-the-input-and-output-allow-quick-decision-tree)
  produces the initial candidates.
- [C.1.6 Before committing to a
  method](#c16-before-committing-to-a-method) verifies the final
  choice.

#### C.1.5 Let constraints eliminate bad choices

##### The problem it solves

Given a correct method and the problem's maximum input sizes, decide
whether its worst-case work, memory, and required output can fit the
limits. Substitute the actual limits into the growth expression;
calling a method "fast" or naming its Big-O class without estimating
the resulting scale is not enough.

##### How it works

A delivery plan that takes ten hours cannot meet a one-hour deadline.
Estimate how work grows before coding, then use input limits to cross
out methods that cannot finish. The [Big-O
guide](#a1-complexity-and-resource-accounting) names that growth.

##### Step-by-step examples

For at most `n=100,000` integers, decide whether any value appears
twice. Compare every pair with sorting followed by one adjacent scan.

Initial state: `n=100000`, `candidates={pairwise,sorting}`,
`selected=none`.

1.  Evaluate pairwise comparison
    - Decision: It requires about five billion
      checks.
    - After: `n=100000`, `candidates={sorting}`, `selected=none`.
    - Known correct: Its required number of checks is incompatible
      with the input size.
2.  Evaluate sorting
    - Decision: About `1.7` million comparisons fit the limit,
      so retain it for correctness and memory review.
    - After: `n=100000`, `candidates={sorting}`, `selected=sorting`.
    - Known correct: The input limit has not ruled it out.

##### Why it works

If a valid worst-case input forces more operations or memory than the
budget permits, implementation details cannot make that method fit. A
loose upper bound alone is not enough to reject it; count the method's
actual work as tightly as practical. Constraint filtering only
rejects methods; it does not prove a remaining method correct.

##### Go example

There is no runtime helper for this planning step. For Go code, count
large memory allocations, integer sizes, unfinished recursive calls,
and memory kept alive by references before choosing how to store data.

##### Work and memory

The estimate itself is planning work. For the chosen method, count
setup, the choices checked from each saved state, output construction,
unfinished recursive calls, and memory that remains in use.

##### When to use it

Use it before implementation and whenever a direct solution seems
correct but possibly too slow or memory-heavy.

##### When not to use it

Do not treat the rough thresholds as guarantees; benchmark or use
stated limits when constants, language costs, or strict time caps
dominate.

##### Common mistakes and edge cases

> **Warning:** An `O(n)` algorithm that stores a million large objects
> may be less feasible than an in-place `O(n log n)` algorithm.

Count the quantities the method actually processes: locations plus
connections for a graph, rows times columns for a grid, saved smaller
results times the choices from each result, and the number of values
produced by a generator.

##### Related material

- [A.1 Complexity and resource
  accounting](#a1-complexity-and-resource-accounting) explains growth
  notation and Go allocation accounting.
- [A.7 Combinatorics and output
  size](#a7-combinatorics-and-output-size) explains factorial and
  exponential output lower bounds.

#### C.1.6 Before committing to a method

##### The problem it solves

Given a promising solution idea, turn it into a precise algorithm that
can be implemented and reviewed. Define every saved value, state why
each transition is legal, show that progress must finish, explain why
no valid answer is missed, and estimate the complete time and memory
cost before committing to code.

##### How it works

Use a preflight checklist before takeoff: state what is already known,
why one step is safe, why progress is guaranteed, what answer is
covered, and what it costs. If one check is unclear, the method is
still a guess.

##### Step-by-step examples

Use binary search to find the first position in `[1,3,3,5]` whose
value is at least `3`. This is called a lower-bound search, and
the expected position is `1`.

Initial state: `left=unset`, `right=unset`, `result=pending`,
`verification=pending`.

1.  Define the saved boundaries
    - Decision: Treat `0..4` as possible insertion positions.
      Positions before `left` are proved too small. The first valid
      position is somewhere from `left` through `right`, including
      both boundaries.
    - After: `left=0`, `right=4`, `result=pending`,
      `verification=pending`.
    - Known correct: Expected position `1` is still possible, and
      position `4` can represent "no array value is large enough."
2.  Test middle index `2`
    - Decision: The value is large enough, so the first
      valid position is at or before `2`; set `right=2`.
    - After: `left=0`, `right=2`, `result=pending`,
      `verification=pending`.
    - Known correct: Positions after `2` cannot be the first
      valid one because position `2` already works.
3.  Test middle index `1`
    - Decision: Position `1` works, so set `right=1`.
    - After: `left=0`, `right=1`, `result=pending`,
      `verification=pending`.
    - Known correct: Position `1` remains possible; later
      positions cannot be first.
4.  Test middle index `0`
    - Decision: Value `1` is too small, so set `left=1`.
    - After: `left=1`, `right=1`, `result=1`,
      `verification=pending`.
    - Known correct: Position `0` was proved too small, and the
      only remaining boundary is the expected answer.
5.  Check progress, cost, and boundaries
    - Decision: Verify that the possible boundary range shrinks,
      then test empty input, duplicates, and a result of
      `len(values)`.
    - After: `left=1`, `right=1`, `result=1`,
      `verification=passed`.
    - Known correct: The invariant includes every insertion
      boundary, and every iteration makes the range smaller.

##### Why it works

Together, these checks show that the stored information is
meaningful, every step is safe, the process finishes, all possible
answers are covered, and the implementation fits the limits. That is
what a correctness and practicality argument must establish.

##### Go example

Translate each checklist item into a nearby name, condition, or
explicit check. For the trace, [E.1 Binary search
bounds](#e1-binary-search-bounds) provides a complete lower-bound
implementation with its boundary meanings stated beside the code.
Calling `LowerBound([]int{1, 3, 3, 5}, 3)` returns `1`.

##### Work and memory

The checklist is design work rather than runtime work. Its final item
records the selected method's complete cost, including setup,
recursion, output, and meaningful Go allocations.

##### When to use it

Use it after selecting a method and before implementation, and
repeat it when requirements or constraints change.

##### When not to use it

Do not replace an actual trace with slogans such as "it is greedy" or
"binary search halves the work." The sentences must name this
problem's stored information and rule.

##### Common mistakes and edge cases

> **Warning:** A loop that usually moves is not enough. If one
> equality branch keeps the same middle position, the progress
> argument fails and the loop may never terminate.

Also verify whether the input may change, whether arithmetic can
overflow, what empty output means, how repeated values behave, whether
some graph locations are unreachable, and whether skipped information
is later needed to rebuild a path or return original positions.

##### Related material

- [A.2 Correctness and invariants](#a2-correctness-and-invariants)
  explains this proof structure in more detail.
- [C.1.5 Let constraints eliminate bad
  choices](#c15-let-constraints-eliminate-bad-choices) supplies the
  resource screen.
- [C.1.3 Pattern cards](#c13-pattern-cards) provides method-specific
  checks.

<a id="e-go-techniques"></a>

## D. Go techniques

These are Go-specific details that commonly affect otherwise correct
algorithms. Each rule below states the behavior that motivates it.

<a id="e1-keep-input-conversion-outside-the-algorithm"></a>

### D.1 Keep input conversion outside the algorithm

#### The problem it solves

Given external text from an automated judge, command line, or web
request, validate and convert it into the typed values an algorithm
expects, then convert the algorithm's result back to the required
output form. Keeping those boundary steps outside the algorithm makes
parsing failures distinct from logic failures and lets the same
algorithm work with inputs from other sources.

#### How it works

Treat the program boundary like a reception desk. The desk translates
an outside request into the program's own types, rejects invalid
requests, and translates the answer back. This small translation layer
is called a **boundary adapter**. The algorithm behind it receives
only problem values and does not need to know whether they came from
JSON, a command line, or a test.

#### Step-by-step examples

Decode `raw = {"values":[2,7,11,15],"target":9}`. The expected result
is `pairInput{Values: []int{2, 7, 11, 15}, Target: 9}`.

Initial state: `stage=decode`, `input={Values:nil,Target:0}`,
`algorithmResult=unset`, `output=unset`, `err=nil`.

1.  Decode the JSON object
    - Decision: The JSON has the expected field names and value
      types, so `json.Unmarshal` can assign both fields.
    - After: `stage=validate`, `input={Values:[2,7,11,15],Target:9}`,
      `algorithmResult=unset`, `output=unset`, `err=nil`.
    - Known correct: The adapter now holds the same four values
      and target represented by `raw`.
2.  Validate the typed request
    - Decision: Four values satisfy this adapter's requirement of
      at least two values.
    - After: `stage=algorithm`,
      `input={Values:[2,7,11,15],Target:9}`,
      `algorithmResult=[0,1]`, `output=unset`, `err=nil`.
    - Known correct: The pair-search algorithm receives a
      valid typed request and no JSON concerns.
3.  Format the successful answer
    - Decision: Encode the returned indexes as a JSON array.
    - After: `stage=done`, `input={Values:[2,7,11,15],Target:9}`,
      `algorithmResult=[0,1]`, `output=[0,1]`, `err=nil`.
    - Known correct: The outside caller receives the same two
      indexes in its required text format.

For a failure case, use `raw = {"values":[2,7],"target":"nine"}`. The
expected result is a decoding error and no usable `pairInput`.

Initial state: `stage=decode`, `input={Values:nil,Target:0}`,
`result=unset`, `err=nil`.

1.  Decode the malformed target
    - Decision: A JSON string cannot be assigned to the Go `int`
      field `Target`, so `json.Unmarshal` returns an error.
    - After: `stage=done`, `input=unusable`, `result=unset`,
      `err=wrappedDecodeError`.
    - Known correct: No partly decoded request crosses the
      boundary.

#### Why it works

After a successful parse, every field has the Go type promised to the
algorithm. Validation then establishes any additional input rules
before the algorithm runs. If either step fails, returning immediately
prevents partial state from being mistaken for valid input. Encoding
happens only after a typed result exists, so representation failures
remain separate from algorithm failures.

#### Go example

`adaptPairRequest` performs the trace's decode, algorithm, and format
stages. Its `findPair` argument is the typed algorithm boundary; the
hash-map implementation in E.3 satisfies that function shape.

```go
package boundary

import (
	"encoding/json"
	"fmt"
)

type pairInput struct {
	// JSON tags bind the external field names to strongly typed Go
	// fields.
	Values []int `json:"values"`
	Target int   `json:"target"`
}

func parsePairInput(raw []byte) (pairInput, error) {
	// Decode at the boundary so the algorithm itself can accept
	// pairInput rather than inspect untyped bytes.
	var input pairInput
	if err := json.Unmarshal(raw, &input); err != nil {
		// Wrap the decoder error with operation context while preserving it
		// for errors.Is or errors.As.
		return pairInput{}, fmt.Errorf("decode pair input: %w", err)
	}
	// Structural decoding cannot enforce this problem-specific
	// requirement.
	if len(input.Values) < 2 {
		return pairInput{}, fmt.Errorf(
			"pair input needs at least two values",
		)
	}
	return input, nil
}

func formatPairOutput(left, right int) ([]byte, error) {
	// A fixed-size array guarantees that successful output has exactly
	// two indices.
	return json.Marshal([2]int{left, right})
}

type pairFinder func(values []int, target int) (int, int, bool)

func adaptPairRequest(
	raw []byte,
	findPair pairFinder,
) ([]byte, error) {
	// Decode and validate before the algorithm can observe the request.
	input, err := parsePairInput(raw)
	if err != nil {
		return nil, err
	}
	// The algorithm receives only typed problem values. Report absence
	// separately instead of formatting invented indexes.
	left, right, found := findPair(input.Values, input.Target)
	if !found {
		return nil, fmt.Errorf("pair input has no matching pair")
	}
	// Convert the typed result to the external JSON representation only
	// after the algorithm succeeds.
	return formatPairOutput(left, right)
}
```

The struct tags such as `json:"values"` connect JSON field names to Go
fields. Returning `%w` in the error keeps the original decoding error
available to callers while adding boundary context.

#### Work and memory

For `b` input bytes containing `n` numbers, decoding examines the
input and builds the `Values` slice, so it takes `O(b)` work and
`O(n)` extra memory. Validation here takes constant work.
Formatting two indices takes constant work and a small output
allocation. These boundary costs are separate from the called
algorithm's own work and memory.

#### When to use it

Use a boundary adapter when stored or transmitted data differs from
the types that make the algorithm easiest to state, when several
callers should share one algorithm, or when malformed input needs a
clear boundary error.

#### When not to use it

Do not add text encoding or decoding to a function whose caller
already supplies the correct Go values. Do not put problem reasoning
in a generic decoder; keep problem-specific validation beside the
problem adapter.

#### Common mistakes and edge cases

> **Warning:** Ignoring a decode error can pass zero values into the
> algorithm. For `{"target":"nine"}`, using `input.Target` after a
> failed unmarshal may incorrectly search for target `0`; return the
> error immediately.

An absent JSON field also leaves its Go field at the zero value. If
zero is valid but omission is not, use a pointer field or a custom
presence check to distinguish them. Decide whether `null`, an empty
array, and unknown fields are allowed instead of accepting them
accidentally.

#### Related material

- [E.3 Hash-map complement lookup](#e3-hash-map-complement-lookup)
  provides the complete pair-search implementation.
- [D.2 Table-driven tests](#d2-table-driven-tests) keeps valid and
  invalid boundary cases under the same test logic.
- [D.4 Map presence and sets](#d4-map-presence-and-sets) explains the
  related problem of distinguishing absence from a stored zero value.

<a id="e2-table-driven-tests"></a>

### D.2 Table-driven tests

#### The problem it solves

Given one function and many cases containing inputs, expected outputs,
or expected errors, run every case through the same test procedure and
identify failures by name. Repeating the test logic by hand for each
case introduces irrelevant differences and makes it easier for one
boundary or failure case to skip an assertion.

#### How it works

Write the checking recipe once and feed it a list of named examples.
In Go, this is called a **table-driven test**: the table is a slice of
test-case structs, and each row contains the input and expected result
for one run.

#### Step-by-step examples

Run rows `{name:"positive even", value:4, want:true}` and
`{name:"negative odd", value:-3, want:false}`. The expected outcome is
two passing subtests.

Initial state: `row=positiveEven`, `passed=[]`, `failed=[]`.

1.  Run `positive even`
    - Decision: `isEven(4)` returns `true`, which equals
      `test.want`.
    - After: `row=negativeOdd`, `passed=[positiveEven]`, `failed=[]`.
    - Known correct: The ordinary even case used the shared
      assertion.
2.  Run `negative odd`
    - Decision: In Go, `-3%2` is `-1`, so the result is `false`
      and matches `test.want`.
    - After: `row=done`, `passed=[positiveEven,negativeOdd]`,
      `failed=[]`.
    - Known correct: A negative value used exactly the same call
      and check.

For a boundary row, use `{name:"zero", value:0, want:true}`. The
expected outcome is a passing subtest because `0%2 == 0`.

Initial state: `row=zero`, `passed=[]`, `failed=[]`.

1.  Run `zero`
    - Decision: `isEven(0)` returns `true`, matching `test.want`.
    - After: `row=done`, `passed=[zero]`, `failed=[]`.
    - Known correct: The boundary value is checked by the same
      rules as every other row.

#### Why it works

Every row reaches the same call and comparison, so differences in
results come from the case data rather than from different test code.
Naming each subtest ties a failure to its input. The loop is finite
because the table is finite; after all rows pass, the promised
behavior holds for every listed case. This is evidence for those
cases, not proof for every possible input.

#### Go example

```go
import "testing"

func isEven(value int) bool {
	// Multiples of two, including zero and negative evens, have
	// remainder zero.
	return value%2 == 0
}

func TestIsEven(testContext *testing.T) {
	// Each row names one behavior and keeps its input beside the
	// expected result.
	tests := []struct {
		name  string
		value int
		want  bool
	}{
		{name: "positive even", value: 4, want: true},
		{name: "negative odd", value: -3, want: false},
		{name: "zero", value: 0, want: true},
	}

	// Run every row as a named subtest so a failure identifies the exact
	// case.
	for _, testCase := range tests {
		testContext.Run(testCase.name, func(subtestContext *testing.T) {
			// Compare the actual result with the row's expectation and report
			// both.
			if actual := isEven(testCase.value); actual != testCase.want {
				subtestContext.Errorf(
					"isEven(%d) = %v, want %v",
					testCase.value,
					actual,
					testCase.want,
				)
			}
		})
	}
}
```

The outer test reports the group; `t.Run` creates a named subtest for
each row. Use `t.Helper` in extracted assertion helpers so failures
point to the calling row rather than the helper body.

#### Work and memory

With `r` rows and per-row behavior costing `T`, the loop does
`O(r * T)` work. The table itself stores `O(r)` case data. Sequential
subtests add only bounded runner state beyond the current row,
although individual inputs and outputs may allocate their own memory.

#### When to use it

Use table-driven tests when many inputs follow the same rules and use
the same checks, especially for boundary values, malformed input,
duplicates, and regression cases that should remain named.

#### When not to use it

Do not force unrelated behaviors with different setup or assertions
into one wide table. Separate focused tests are clearer for long
stateful workflows; randomized or generated-input testing is a better
companion when a property must be checked across many automatically
produced inputs.

#### Common mistakes and edge cases

> **Warning:** An empty table executes zero assertions and still
> passes. A test with `tests := []caseType{}` proves nothing; keep at
> least one ordinary row and the relevant boundary rows.

When a row owns a slice, map, or pointer that the tested function
mutates, freshly construct or clone that value per row so another row
cannot inherit changed state. If setup returns an error, call
`t.Fatalf` before reading the partly initialized result.

#### Related material

- [A.2 Correctness and invariants](#a2-correctness-and-invariants)
  separates test evidence from a general correctness argument.
- [D.1 Keep input conversion outside the
  algorithm](#d1-keep-input-conversion-outside-the-algorithm)
  identifies parsing failures that should become explicit rows.
- [E.9 Sorted two-pointer pair
  search](#e9-sorted-two-pointer-pair-search) lists representative
  cases for a reference implementation.

<a id="e3-sorting-without-unexpected-mutation"></a>

### D.3 Sorting without unexpected mutation

#### The problem it solves

Given a slice that must be sorted for an algorithm, decide whether the
caller permits its order to change. Go's sorting functions rearrange
the supplied backing array, so a helper that promises to preserve the
input must copy the values before sorting rather than copying only the
slice header.

#### How it works

Make a separate row of cards before arranging them. Changing a value
that the caller can observe is called **mutation**. A shallow slice
clone copies the outer elements into a new backing array, so sorting
the clone does not reorder the caller's slice. Pointer-like elements
inside those copied values may still refer to shared nested data.

#### Step-by-step examples

Call `sortedCopy(values)` with `values = [3,1,2]`. The expected result
is `sorted = [1,2,3]` while `values` remains `[3,1,2]`.

Initial state: `values=[3,1,2]@inputArray`, `sorted=unset`.

1.  Clone the slice
    - Decision: The function promises not to reorder `values`, so
      `slices.Clone` creates a separate backing array for `sorted`.
    - After: `values=[3,1,2]@inputArray`,
      `sorted=[3,1,2]@resultArray`.
    - Known correct: The caller's order is isolated before sorting
      starts.
2.  Sort the clone
    - Decision: `sort.Ints` rearranges only `sorted` into
      ascending order.
    - After: `values=[3,1,2]@inputArray`,
      `sorted=[1,2,3]@resultArray`.
    - Known correct: The result is ordered and the promised input
      state is preserved.

For a deliberately broken version, assign `sorted := values` instead
of cloning. After `sort.Ints(sorted)`, both variables unexpectedly
show `[1,2,3]`.

Initial state: `values=[3,1,2]@sharedArray`, `sorted=unset`.

1.  Copy only the slice descriptor
    - Decision: Assignment copies the pointer, length, and
      capacity, not the backing array.
    - After: `values=[3,1,2]@sharedArray`,
      `sorted=[3,1,2]@sharedArray`.
    - Known correct: No independent element storage was created.
2.  Sort through the alias
    - Decision: Sorting writes into the shared backing array.
    - After: `values=[1,2,3]@sharedArray`,
      `sorted=[1,2,3]@sharedArray`.
    - Known correct: The promise to preserve the input is broken.

#### Why it works

Before sorting begins, the clone holds the same element values in an
independent outer array. Sorting permutes only that array, so the
returned values are ordered while every position in the caller's array
keeps its original value. The sorting call finishes after a bounded
number of comparisons and swaps, and the independent arrays preserve
the two promised observations.

#### Go example

```go
import (
	"slices"
	"sort"
)

func sortedCopy(values []int) []int {
	// Clone allocates independent storage; sort.Ints may now mutate
	// sorted without changing the caller's slice.
	sorted := slices.Clone(values)
	sort.Ints(sorted)
	return sorted
}
```

`slices.Clone(nil)` returns `nil`; a non-nil input returns a slice
with copied elements. For records sorted with `sort.Slice` or
`slices.SortFunc`, state the primary order and every required
tie-break rule in the comparison.

#### Work and memory

For `n` integers, cloning copies `n` values in `O(n)` time and
allocates `O(n)` output space. Comparison sorting takes `O(n log n)`
time and rearranges that output slice without allocating another
`O(n)` result. Its internal control stack uses `O(log n)` auxiliary
space. The combined time is `O(n log n)`; the required returned copy
is the dominant memory cost.

#### When to use it

Clone before sorting when the function promises to preserve input
order, when the caller may reuse the input, or when tests compare the
input before and after the call.

#### When not to use it

Do not pay for a clone when the function explicitly allows in-place
reordering and the caller no longer needs the old order. A shallow
clone is also insufficient when sorting code mutates fields through
pointer-like elements; copy those nested values when isolation
requires it.

#### Common mistakes and edge cases

> **Warning:** `sorted := values` does not copy slice elements. With
> `values=[3,1,2]`, sorting `sorted` also changes `values` to
> `[1,2,3]`; use `slices.Clone` when the old order must survive.

Duplicates require a tie rule only when the relative order of equal
primary keys matters; choose a **stable sort**, which preserves the
earlier order of equal keys, when the function promises that behavior.
Nil and empty slices are already sorted. Tests should cover both
allowed in-place changes and a promise to preserve the input.

#### Related material

- [B.1 Arrays, slices, and strings](#b1-arrays-slices-and-strings)
  explains slice descriptors, backing arrays, capacity, and append
  behavior.
- [C.1.3.5 Put related values together (sorting and
  selection)](#c135-sorting-and-selection) explains the
  algorithm-selection reasons for sorting.
- [D.9 Slice storage and retained
  memory](#d9-slice-storage-and-retained-memory) explains how sharing
  an array also affects memory lifetime.

<a id="e4-map-presence-and-sets"></a>

<a
id=
e4-distinguish-missing-keys-from-stored-values-map-presence-and-sets>
</a>

### D.4 Map presence and sets

#### The problem it solves

Given a Go map query, distinguish a missing key from a present key
whose stored value equals the type's zero value. When the only
required fact is membership, represent a set without inventing
unrelated value data and still use the presence result when absence
matters.

#### How it works

Think of labeled mailboxes. Opening an unknown label gives you an
empty-looking result, but that alone cannot tell you whether the
mailbox exists and is empty. Go can return both the stored value and a
presence flag. A collection that records only presence is called a
**set**; Go commonly represents one with a map whose keys are the set
members. The contracts below are `entries[key] -> stored integer` and
`seen[value] -> whether value has appeared`. For example,
`entries["ready"]=0` is present data, while a missing `seen["go"]`
entry means `"go"` has not appeared yet.

#### Step-by-step examples

Start with `entries = {"ready": 0}` and query `"ready"` and
`"missing"`. The expected lookup results are `(0, true)` and
`(0, false)`.

Initial state: `entries={"ready":0}`, `nextKey="ready"`,
`answers=[]`.

1.  Query `"ready"`
    - Decision: The key exists and its stored value is `0`.
    - After: `entries={"ready":0}`, `nextKey="missing"`,
      `answers=[(0,true)]`.
    - Known correct: The zero came from a present entry.
2.  Query `"missing"`
    - Decision: The key does not exist, so Go supplies the integer
      zero value.
    - After: `entries={"ready":0}`, `nextKey=done`,
      `answers=[(0,true),(0,false)]`.
    - Known correct: The same value is now identified as a
      missing-key result.

For a set example, process `values = ["go", "go", "rust"]` and keep
first appearances. The expected result is `["go", "rust"]`.

Initial state: `index=0`, `seen={}`, `result=[]`.

1.  Read `"go"`
    - Decision: It is absent; add and append it.
    - After: `index=1`, `seen={"go"}`, `result=["go"]`.
    - Known correct: The first distinct input appears once.
2.  Read the second `"go"`
    - Decision: It is present; skip the duplicate.
    - After: `index=2`, `seen={"go"}`, `result=["go"]`.
    - Known correct: The duplicate creates no second output.
3.  Read `"rust"`
    - Decision: It is absent; add and append it.
    - After: `index=done`, `seen={"go","rust"}`,
      `result=["go","rust"]`.
    - Known correct: Each distinct input now appears once, in
      first-appearance order.

The result order comes from scanning `values`, not from iterating over
`seen`.

#### Why it works

The two-result lookup separates two facts that the value alone cannot
encode: what value a key maps to and whether the mapping exists. A set
uses exactly the second fact. Before each value in the set trace,
`seen` contains every distinct earlier value and `result` lists those
values in first-appearance order. A present value is safely skipped;
an absent value is added to both structures, so the fact remains true.
After the finite input is exhausted, `result` contains each input
value exactly once in the promised order.

#### Go example

```go
func lookup(entries map[string]int, key string) (int, bool) {
	// The comma-ok form distinguishes a missing key from a present key
	// storing 0.
	value, found := entries[key]
	return value, found
}

func uniqueInOrder(values []string) []string {
	// struct{} is a zero-size membership marker; result preserves
	// encounter order.
	seen := make(map[string]struct{})
	result := make([]string, 0, len(values))
	for _, value := range values {
		if _, found := seen[value]; found {
			// A prior occurrence already occupies this value's output
			// position.
			continue
		}
		seen[value] = struct{}{}
		result = append(result, value)
	}
	return result
}
```

#### Work and memory

Go's language rules do not promise a fixed time bound for map
operations. In the usual hash-map cost model--where a computed key
fingerprint selects a small storage area--a lookup or insertion takes
expected constant time, written `O(1)`. Under that model, processing
`n` values takes expected `O(n)` time. If there are `u` distinct
values, `seen` stores `u` keys. The result has length `u`, but
`make(..., 0, len(values))` reserves room for `n` strings to avoid
growth allocations, so the function's total additional storage is
`O(n)`. `struct{}` avoids a per-entry value payload, but map keys and
bucket storage still use memory.

#### When to use it

Use the two-result lookup when a zero value is valid data, when
membership changes control flow, or when a missing key must produce a
different result. Use a set when the question is only whether a
comparable key has appeared or belongs to a collection.

#### When not to use it

Do not rely on a map when iteration order is part of the result. For a
small, dense integer range, a boolean slice may express membership
with less overhead. A plain set also cannot answer ordering or range
questions without another structure.

#### Common mistakes and edge cases

> **Warning:** `entries[key] == 0` is not an absence test. With
> `entries={"ready":0}`, it incorrectly treats the present key as
> missing; use the `found` result.

> **Warning:** Reading a nil map is allowed, but writing to it panics.
> Declare `var seen map[string]struct{}` and then execute
> `seen["go"] = struct{}{}` to reproduce the failure; call `make`
> first.

A `map[T]bool` can represent a set, but a stored `false` is
indistinguishable from absence unless the code also reads `found`.
Duplicate input must not be appended twice, and tests must not expect
`range` over a map to return a stable order.

#### Related material

- [C.1.3.4 Hash map or set](#c134-hash-map-or-set) explains the
  algorithm-selection clues and invariants for lookup-based methods.
- [E.3 Hash-map complement lookup](#e3-hash-map-complement-lookup)
  shows map presence in a complete reference algorithm.
- [A.2 Correctness and invariants](#a2-correctness-and-invariants)
  explains how the `seen` and `result` relationship proves duplicate
  removal correct.

<a id="e5-priority-queues"></a>

### D.5 Priority queues

#### The problem it solves

Given items that arrive over time, repeatedly add an item and remove
the remaining item with the smallest, largest, or most urgent key.
Sorting the whole collection after every insertion repeats work, while
scanning every item for each removal makes each priority operation
linear.

#### How it works

Keep the next item near the front while allowing the rest to be only
partly ordered. A **priority queue** removes items by priority rather
than arrival time. Go's `container/heap` package implements one with a
**heap**, an array whose parent-child ordering makes the
highest-priority item available at index zero.

#### Step-by-step examples

Push `4`, `1`, and `3` into an empty smallest-first queue, then remove
one item. The expected removed value is `1` and the remaining heap
array is `[3,4]`.

Initial state: `h=[]`, `next=push(4)`, `removed=[]`.

1.  Push `4`
    - Decision: Append `4`; it has no parent to compare with.
    - After: `h=[4]`, `next=push(1)`, `removed=[]`.
    - Known correct: Index zero contains the only and therefore
      smallest item.
2.  Push `1`
    - Decision: Append `1`, then swap it with parent `4` because
      `1 < 4`.
    - After: `h=[1,4]`, `next=push(3)`, `removed=[]`.
    - Known correct: Every child is at least its parent, so `1` is
      at the root.
3.  Push `3`
    - Decision: Append `3`. Its parent is `1`, so no swap is
      needed.
    - After: `h=[1,4,3]`, `next=pop`, `removed=[]`.
    - Known correct: The root is still the smallest item.
4.  Pop
    - Decision: The package moves root `1` to the end, repairs the
      remaining heap, and asks the type's `Pop` method to shorten the
      slice.
    - After: `h=[3,4]`, `next=done`, `removed=[1]`.
    - Known correct: `1` was no larger than any queued item, and
      `3` is now the smallest remaining item.

#### Why it works

The fact that remains true is that every parent has at least as much
priority as its children. Push repairs that fact while moving the new
item toward the root; pop repairs it while moving the replacement root
downward. Consequently, no descendant can outrank index zero, so
removing the root returns the required item. Each repair moves through
a finite number of array levels.

#### Go example

Callers must push only integers into `intMinHeap`.

```go
import (
	"container/heap"
	"slices"
)

type intMinHeap []int

// Len, Less, and Swap let heap.Interface maintain this slice as a
// min-heap.
func (heapValues intMinHeap) Len() int { return len(heapValues) }
func (heapValues intMinHeap) Less(leftIndex, rightIndex int) bool {
	return heapValues[leftIndex] < heapValues[rightIndex]
}
func (heapValues intMinHeap) Swap(leftIndex, rightIndex int) {
	heapValues[leftIndex], heapValues[rightIndex] =
		heapValues[rightIndex], heapValues[leftIndex]
}

func (heapValues *intMinHeap) Push(value any) {
	// heap.Interface passes values as any. The .(int) assertion recovers
	// the required element type and would panic if a caller supplied a
	// non-int value. The pointer receiver then grows the heap's backing
	// slice.
	*heapValues = append(*heapValues, value.(int))
}

func (heapValues *intMinHeap) Pop() any {
	// container/heap first swaps the root with the final slot; Pop
	// removes and returns that final slot rather than searching for the
	// minimum itself.
	valuesBeforeRemoval := *heapValues
	lastIndex := len(valuesBeforeRemoval) - 1
	value := valuesBeforeRemoval[lastIndex]
	*heapValues = valuesBeforeRemoval[:lastIndex]
	return value
}

func pushThenPop(values []int) (int, []int) {
	// Start empty, push each traced arrival through container/heap, then
	// remove exactly one minimum as the trace does.
	heapValues := &intMinHeap{}
	heap.Init(heapValues)
	for _, value := range values {
		heap.Push(heapValues, value)
	}
	minimum := heap.Pop(heapValues).(int)
	return minimum, slices.Clone(*heapValues)
}

func popAscending(values []int) []int {
	// Clone protects the caller, and Init rearranges the copy into heap
	// order.
	heapValues := intMinHeap(slices.Clone(values))
	heap.Init(&heapValues)

	// Every heap pop removes the current minimum, producing ascending
	// output.
	result := make([]int, 0, len(values))
	for heapValues.Len() > 0 {
		result = append(result, heap.Pop(&heapValues).(int))
	}
	return result
}
```

For the trace, `pushThenPop([]int{4, 1, 3})` returns removed value
`1` and remaining heap slice `[3,4]`. `popAscending` demonstrates the
separate operation of removing every value in priority order.

The package moves the root to the final slot before it calls the
type's `Pop` method, so that method removes the last element. The type
assertion `value.(int)` is safe here because this queue's callers push
only integers.

#### Work and memory

For `n` initial items, `heap.Init` takes `O(n)` time. Reading the root
takes `O(1)` time. Each push or pop repairs at most the heap height,
so it takes `O(log n)` time. The heap stores `O(n)` items.
`popAscending` uses the `O(n)` clone as working heap storage and
constructs an `O(n)` returned output. Those slices coexist but still
use `O(n)` total memory. Its total work is `O(n log n)`.

#### When to use it

Use a priority queue when items arrive or change over time and each
step needs the current smallest, largest, earliest, cheapest, or
otherwise most urgent item without fully ordering everything.

#### When not to use it

Do not use a heap when one final sorted order is the only goal;
sorting once is simpler. A heap is also the wrong tool when arbitrary
items must be found or removed quickly unless another structure tracks
their positions.

#### Common mistakes and edge cases

> **Warning:** Calling `h.Pop()` directly skips the package's root
> swap and repair. With `h=[1,4,3]`, the method removes `3` rather
> than minimum `1`; call `heap.Pop(&h)`.

Popping or reading `h[0]` when `Len()==0` panics. Reversing `Less`
changes smallest-first to largest-first. For queues of pointers, clear
the removed final slot before shortening the slice so the backing
array does not retain it; see
[D.9](#d9-slice-storage-and-retained-memory).

#### Related material

- [C.1.3.27 Remove urgent
  items](#c1327-heap-or-priority-queue) explains heap selection,
  correctness, and basic operations.
- [D.9 Slice storage and retained
  memory](#d9-slice-storage-and-retained-memory) explains clearing
  pointer elements removed from long-lived heaps.

<a id="e6-bytes-runes-and-visible-characters"></a>

### D.6 Bytes, runes, and visible characters

#### The problem it solves

Given a Go string and a requirement stated in terms of "characters,"
choose whether the algorithm should process bytes, Unicode code points
called runes, or user-visible symbols. These units have different
counts and boundaries; treating them as interchangeable can split an
encoding or return indexes in the wrong unit.

#### How it works

Think of visible text as symbols packed into bytes. **Unicode
Transformation Format 8 (UTF-8)** says how one or more bytes form a
Unicode **code point**, called a **rune** in Go. One visible symbol,
such as an accented letter or emoji, may still need several runes.
Unicode calls that full displayed group a **grapheme cluster**.

#### Step-by-step examples

Use `text = "A\u00e9"`. The expected byte length is `3`, and the
expected rune sequence is `['A','\u00e9']`.

Initial state: `byteIndex=unset`, `byteCount=unset`, `runes=[]`.

1.  Inspect stored bytes
    - Decision: `len` reports stored bytes, not runes.
    - After: `byteIndex=0`, `byteCount=3`, `runes=[]`.
    - Known correct: All three encoded bytes are counted once.
2.  Decode `A`
    - Decision: Byte `0x41` is the complete rune `A`.
    - After: `byteIndex=1`, `byteCount=3`, `runes=['A']`.
    - Known correct: The first byte and first rune have both been
      processed once.
3.  Decode `\u00e9`
    - Decision: UTF-8 combines bytes `0xC3 0xA9` into one rune.
    - After: `byteIndex=done`, `byteCount=3`,
      `runes=['A','\u00e9']`.
    - Known correct: Byte positions and rune values both refer to
      the same original string without pretending that every rune is
      one byte.

For the boundary case, use `text = "e\u0301"`: `e` followed by a
combining acute accent. The expected rune count is `2` even though
many displays show one accented symbol.

Initial state: `text="e\u0301"`, `runes=[]`.

1.  Decode code points
    - Decision: `[]rune` preserves `U+0065` and `U+0301` separately.
    - After: `text="e\u0301"`, `runes=['e','\u0301']`.
    - Known correct: Rune processing answers a code-point
      question, not a visible-symbol question.

#### Why it works

UTF-8 defines exactly how the stored bytes decode into code points.
`range` advances by the width of each decoded rune, so each valid code
point is visited once and its reported index is its true starting byte
position. Converting to `[]rune` records those decoded values in
order. Visible-symbol grouping follows additional Unicode rules, so a
code-point loop cannot promise that grouping by itself.

#### Go example

```go
type textUnits struct {
	// ByteCount describes storage length; Runes describes decoded
	// Unicode code points. Neither necessarily equals the number of
	// visible grapheme clusters.
	ByteCount int
	Runes     []rune
}

func describeText(text string) textUnits {
	// len counts UTF-8 bytes, while []rune decodes the string and
	// allocates a separate rune slice.
	return textUnits{
		ByteCount: len(text),
		Runes:     []rune(text),
	}
}
```

For `"A\u00e9"` this returns byte count `3` and two runes. A `range`
loop is often better when the algorithm can stream through runes
without allocating a slice.

#### Work and memory

`len(text)` reads the stored byte length in `O(1)` time. Ranging over
`b` bytes takes `O(b)` time and constant extra memory. Building
`[]rune(text)` also takes `O(b)` time and allocates `O(r)` rune
storage, where `r` is the decoded rune count. Visible-symbol
segmentation has its own library-specific state and costs.

#### When to use it

Use bytes for binary data, American Standard Code for Information
Interchange (ASCII) contracts, or protocols defined in bytes. Use
runes when comparisons, positions, or counts are defined by Unicode
code points. Use visible-symbol segmentation for cursor movement or
limits that must match what a reader sees.

#### When not to use it

Do not convert every string to runes when the input uses a guaranteed
one-byte alphabet or the algorithm only copies bytes. Do not use rune
indexes as byte slice boundaries without converting them back to byte
positions.

#### Common mistakes and edge cases

> **Warning:** `text[1:]` can begin inside a multibyte rune. With
> `text="\u00e9"`, byte index `1` starts at `0xA9` and the result is
> not valid UTF-8; slice only at known byte boundaries.

Invalid UTF-8 causes `range` to produce `utf8.RuneError` for invalid
encoding units. An empty string has zero bytes and zero runes.
Text that looks identical may have different code-point sequences, as
`"\u00e9"` and `"e\u0301"` demonstrate. Equality must therefore follow
the problem's definition.

#### Related material

- [B.1 Arrays, slices, and strings](#b1-arrays-slices-and-strings)
  explains string immutability and basic sequence storage.
- [C.1.3.7 Reuse a partial match](#c137-string-matching) explains how
  the chosen text unit affects string-matching positions.
- [C.1.3.33 Share stored word beginnings](#c1333-trie) shows a
  byte-based trie that must clearly state its allowed alphabet.

<a id="e7-type-parameters"></a>

### D.7 Type parameters

#### The problem it solves

Given one operation that should work for several Go types, implement
it once while preserving the operations and return type each caller
may use safely. Copying the implementation for every concrete type can
drift, while accepting `any` discards useful compile-time checks and
forces callers to recover concrete values.

#### How it works

Write one recipe with a labeled blank for the value type, then let
each caller fill in that blank. The label is a **type parameter**.
Code using type parameters is often called **generic** because one
checked implementation works with a group of concrete types.

#### Step-by-step examples

Call `first([]int{8,13})` and `first([]string{"go","rust"})`. The
expected results are `(8,true)` and `("go",true)`.

Initial state: `nextCall=first([]int{8,13})`, `results=[]`.

1.  Call with integers
    - Decision: The compiler infers `T=int`; the slice is nonempty.
    - After: `nextCall=first([]string{"go","rust"})`,
      `results=[(8,true)]`.
    - Known correct: No string or other type can appear in this
      call's result.
2.  Call with strings
    - Decision: The compiler infers `T=string`; the slice is nonempty.
    - After: `nextCall=done`, `results=[(8,true),("go",true)]`.
    - Known correct: The same control flow is checked for a
      different value type.

For the boundary case, call `first([]int{})`. The expected result is
`(0,false)`.

Initial state: `T=int`, `values=[]`, `result=pending`.

1.  Call with an empty slice
    - Decision: There is no first element, so declare
      `var zero T`.
    - After: `T=int`, `values=[]`, `result=(0,false)`.
    - Known correct: The boolean distinguishes "no value" from a
      stored first value that happens to be zero.

#### Why it works

Each call chooses one concrete type that satisfies the declared
constraint. The compiler checks the generic body using only operations
promised by that constraint, then checks callers with their chosen
type. Therefore the returned `T` matches the input element type, and
the absence boolean safely separates an empty input from the type's
zero value.

#### Go example

```go
func first[T any](values []T) (T, bool) {
	if len(values) == 0 {
		// A generic type has no universal literal, so declare its zero
		// value and pair it with false to report absence.
		var zero T
		return zero, false
	}
	// Indexing is safe after the length check; true distinguishes a
	// stored zero.
	return values[0], true
}
```

`any` is an alias for `interface{}`, but `T` remains one specific type
for each call. Use a narrower constraint when the implementation needs
comparison, ordering, arithmetic, or methods.

#### Work and memory

`first` checks the length and reads at most one element, so every call
takes `O(1)` time and `O(1)` additional memory. Type parameters do not
make this operation scan or copy the slice.

#### When to use it

Use type parameters when the same data-structure or function behavior
applies to multiple types and its required operations can be expressed
by a small, clear constraint.

#### When not to use it

Do not add a type parameter when only one concrete type is meaningful
or when different types require different behavior. Use an **interface
value**, which can hold any concrete type that provides the required
methods, when a runtime collection must mix such types.

#### Common mistakes and edge cases

> **Warning:** `left == right` does not compile when `T` is
> constrained by `any` because slices, maps, and functions are not
> comparable. Use `T comparable` only when equality is required.

An empty input needs a second result or another explicit absence
convention because every `T` has a zero value. Avoid constraints that
expose operations the function never uses; they reject valid callers
without improving safety.

#### Related material

- [E.4 Generic singly linked list](#e4-generic-singly-linked-list)
  applies a type parameter to nodes.
- [E.5 Generic stack](#e5-generic-stack) applies one to a reusable
  container.
- [D.4 Map presence and sets](#d4-map-presence-and-sets) explains why
  generic map keys need the `comparable` constraint.

<a id="e8-replacing-recursion-with-a-loop"></a>

### D.8 Replacing recursion with a loop

#### The problem it solves

Given a recursive traversal whose input depth may be too large for the
Go call stack, perform the same visits with an explicit bounded data
structure and a loop. The replacement must store every piece of
unfinished work that a recursive stack frame previously remembered,
including the point at which processing should resume.

#### How it works

Replace the hidden pile of unfinished calls with a slice the program
manages. That slice is an **explicit stack**. Each stored item must
contain everything a recursive call would need when its turn resumes:
the current value, the next choice, and any partial result used after
a child finishes.

#### Step-by-step examples

Traverse a tree with root `8`, left child `4`, and right child `12`.
Write the stack with its top on the right. The expected preorder
result is `[8,4,12]`.

Initial state: `stack=[]`, `result=[]`, `next=root(8)`.

1.  Push the root
    - Decision: The root is non-nil, so push `8`.
    - After: `stack=[8]`, `result=[]`, `next=pop`.
    - Known correct: Every unvisited reachable node is either the
      root on the stack or below it.
2.  Visit `8`
    - Decision: Pop and append `8`; push right `12` before left
      `4`.
    - After: `stack=[12,4]`, `result=[8]`, `next=pop`.
    - Known correct: `8` is the required first value, and `4` is
      next.
3.  Visit `4`
    - Decision: Pop and append leaf `4`; it has no children to
      push.
    - After: `stack=[12]`, `result=[8,4]`, `next=pop`.
    - Known correct: The root and complete left subtree are in
      preorder.
4.  Visit `12`
    - Decision: Pop and append leaf `12`.
    - After: `stack=[]`, `result=[8,4,12]`, `next=done`.
    - Known correct: No pending node remains, and the full
      preorder is complete.

For the boundary case, use `root=nil`. The expected result is an empty
slice.

Initial state: `root=nil`, `stack=[]`, `result=[]`.

1.  Check the root
    - Decision: There is no call frame to simulate, so return
      immediately.
    - After: `root=nil`, `stack=[]`, `result=[]`.
    - Known correct: An empty tree contains no node to visit.

#### Why it works

Before each loop, the stack contains exactly the roots of subtrees
that the recursive version has not yet visited, in reverse visit
order. Popping visits the same node the next recursive call would
visit. Pushing right before left preserves the left-before-right
order. Each node is pushed and popped once, so the finite stack
eventually empties and the result matches recursive preorder.

#### Go example

```go
type traversalNode struct {
	Value int
	Left  *traversalNode
	Right *traversalNode
}

func preorderIterative(root *traversalNode) []int {
	// Avoid placing a nil pointer on the stack; an empty tree has empty
	// preorder.
	if root == nil {
		return []int{}
	}

	// The explicit stack stores nodes whose subtrees still need
	// processing.
	stack := []*traversalNode{root}
	result := make([]int, 0)
	for len(stack) > 0 {
		// Pop the final element to obtain last-in, first-out behavior.
		last := len(stack) - 1
		node := stack[last]
		stack = stack[:last]

		// Preorder visits the node before its children. Push right first so
		// the left child, placed on top last, is processed next.
		result = append(result, node.Value)
		if node.Right != nil {
			// Right waits below left on the LIFO stack.
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			// Left is pushed last and is therefore the next node popped.
			stack = append(stack, node.Left)
		}
	}
	return result
}
```

This frame needs only a node pointer because preorder performs all
work for a node before scheduling its children. Postorder or
backtracking may need a frame with a next-step field and partial
results.

#### Work and memory

For `n` nodes, each node is pushed, popped, and appended once, so
traversal takes `O(n)` time. The result uses `O(n)` output memory. The
explicit stack uses `O(h)` extra memory for this binary-tree
traversal, where `h` is the tree height; in the worst case `h=n`. The
recursive form uses a call stack with the same `O(h)` height bound.

#### When to use it

Use an explicit stack when depth may be too large for safe recursion,
when the pending work must be paused or inspected, or when an
iterative interface is required.

#### When not to use it

Do not replace short, safely bounded recursion when manual stack items
make the logic harder to understand. A stack containing only nodes is
not enough when work must happen after a child returns; each item must
also record what should happen next.

#### Common mistakes and edge cases

> **Warning:** Pushing left before right reverses preorder. For root
> `8` with left child `4` and right child `12`, it produces
> `[8,12,4]`; push right first so left is popped first.

Handle a nil root before pushing it, or dereferencing the popped nil
pointer will panic. For graph traversal, track visited nodes
separately or a cycle can keep adding work forever.

#### Related material

- [A.3 Recursion and saved
  work](#a3-recursion-and-saved-work) explains
  recursive call frames and stack-memory accounting.
- [B.4.2 Traversal order](#b42-traversal-order) explains preorder,
  inorder, and postorder meaning.
- [C.1.3.20 Depth-first search](#c1320-depth-first-search) explains
  the broader search method that can use either call stack.

<a id="e9-slice-storage-and-retained-memory"></a>

### D.9 Slice storage and retained memory

#### The problem it solves

Given a slice that is shortened or reduced to a small subrange, ensure
that data no longer needed can become unreachable and be reclaimed.
The visible slice length is not enough: its backing array may still be
large, and unused slots in a pointer-containing array may still refer
to removed objects.

#### How it works

A slice is a window onto an array, not the array itself. Narrowing the
window does not cut away the room behind it. When a small window must
live much longer than its source, copy it into a right-sized array.
When removing a pointer from a long-lived array, erase the old slot
before hiding it beyond the slice length.

#### Step-by-step examples

Suppose `large` refers to a 1,048,576-byte array and
`header := large[:16]` must be kept in a cache after `large` is no
longer needed. The expected compact result contains the same 16 bytes
without retaining the large array through `header`.

Initial state: `large=1048576Bytes@largeArray`, `header=unset`,
`compact=unset`.

1.  Take a small view
    - Decision: `large[:16]` changes the visible length but shares
      the array.
    - After: `large=1048576Bytes@largeArray`,
      `header=16Bytes@largeArray`, `compact=unset`.
    - Known correct: The byte values are right, but storage is not
      detached.
2.  Copy the surviving bytes
    - Decision: Allocate `len(header)=16` bytes and copy the
      visible data.
    - After: `large=1048576Bytes@largeArray`,
      `header=16Bytes@largeArray`, `compact=16Bytes@compactArray`.
    - Known correct: Keeping `compact` no longer requires the old
      large array through this slice.

For a pointer-removal example, let `items=[A,B,C]` and pop `C`. The
expected visible result is `[A,B]`, and the old final array slot
should be `nil`.

Initial state: `items=[A,B,C]`, `last=2`, `value=unset`.

1.  Save the removed pointer
    - Decision: Save `items[2]` so it can be returned.
    - After: `items=[A,B,C]`, `last=2`, `value=C`.
    - Known correct: The value remains available to return after
      its old array slot is cleared.
2.  Clear the old slot
    - Decision: Assign `nil` to `items[2]`.
    - After: `items=[A,B,nil]`, `last=2`, `value=C`.
    - Known correct: The result is preserved separately, and the
      reusable array slot no longer retains `C`.
3.  Shorten the visible slice
    - Decision: Reslice to `items[:2]`.
    - After: `items=[A,B]`, `last=2`, `value=C`, `hiddenSlot=nil`.
    - Known correct: The returned slice contains exactly the
      unremoved pointers in their original order.

#### Why it works

`copy` writes each visible element into a separately allocated array,
so the new slice cannot keep the source array reachable. Clearing a
removed pointer replaces the backing array's reference before
reslicing hides the slot. The visible sequence remains correct because
copying preserves order and popping returns the saved final value.

#### Go example

```go
func compactBytes(view []byte) []byte {
	// Allocate exactly enough independent storage, then copy the visible
	// bytes. The returned slice cannot retain or mutate view's larger
	// backing array.
	compact := make([]byte, len(view))
	copy(compact, view)
	return compact
}

func popPointer[T any](items []*T) ([]*T, *T, bool) {
	// Empty input has no final slot to read or clear.
	if len(items) == 0 {
		return items, nil, false
	}

	// Save the value before shrinking, then clear the pointer slot so
	// the backing array does not unnecessarily keep the referenced
	// object reachable.
	last := len(items) - 1
	value := items[last]
	items[last] = nil
	// Return the shortened descriptor separately; reslicing this local
	// variable cannot change the caller's slice length by itself.
	return items[:last], value, true
}
```

`make` gives `compact` exactly the requested length and capacity here.
`popPointer` returns the shortened slice because changing a local
slice descriptor does not change the caller's descriptor.

#### Work and memory

Copying a `k`-element view takes `O(k)` time and allocates `O(k)`
memory. Popping and clearing the final pointer takes `O(1)` time and
no new slice storage. The copy deliberately trades a small allocation
now for allowing a much larger array to become unreachable sooner.

#### When to use it

Compact a slice when a small portion will outlive a much larger input.
Clear removed pointer-bearing slots in long-lived stacks, queues,
heaps, buffers, or pools whose backing arrays remain reachable.

#### When not to use it

Do not copy every short-lived subslice; the allocation can cost more
than brief retention. Clearing scalar slots such as integers usually
does not affect object reachability, though the data structure's rules
may still require it.

#### Common mistakes and edge cases

> **Warning:** `items = items[:last]` hides but does not overwrite the
> removed pointer. With backing storage `[A,B,C]`, the unused slot can
> still retain `C`; assign `items[last] = nil` first.

A copied outer slice still shares objects reached through pointer-like
elements. Appending to one of two shared slices may mutate the common
array when capacity remains and may detach when it does not, so
correctness must not depend on guessing which outcome occurs. Empty
and nil slices require no compaction; `popPointer` reports `false`
instead of indexing them.

#### Related material

- [B.1 Arrays, slices, and strings](#b1-arrays-slices-and-strings)
  explains slice length, capacity, shared arrays, and append behavior.
- [D.3 Sorting without unexpected
  mutation](#d3-sorting-without-unexpected-mutation) uses cloning to
  protect observable order.
- [D.5 Priority queues](#d5-priority-queues) is one long-lived
  slice-backed structure that may hold pointers.

<a id="d-go-reference-implementations"></a>

## E. Go reference implementations

These snippets are complete, small implementations of ideas taught
elsewhere in the guide. Each entry states its promised behavior and
required input conditions, traces the variables used by the code, and
gives concrete cases to test. The linked teaching entry explains how
to recognize and choose the method. This section keeps implementation
details beside runnable Go code.

Some snippets use a type parameter written as `T`. Treat `T` as a
placeholder for the value type chosen by the caller, such as `int` or
`string`; the [type-parameter note](#d7-type-parameters) explains why
this is useful.

<a id="d1-binary-search-bounds"></a>

### E.1 Binary search bounds

#### The problem it solves

Given values sorted in ascending order, find where a target begins,
where it ends, or whether it occurs at all. Equal values make an
ordinary "find any match" search insufficient when the caller needs
exact range boundaries.

#### How it works

Open a dictionary near the middle, then keep only the half that can
contain the first matching word. **Binary search** repeats that move.
The code maintains three regions. Indexes before `left` are proved too
small. Indexes from `right` onward are proved large enough. The
half-open range `[left,right)`, which includes `left` but not `right`,
is still unclassified. A lower bound finds the first value at least as
large as the target; an upper bound finds the first larger value.
Either boundary may be `len(values)`.

#### Step-by-step examples

For `values=[1,2,2,2,5]` and `target=2`, `LowerBound` should return
`1`.

Initial state: `left=0`, `right=5`, `result=pending`.

1.  Inspect index `2`
    - Decision: `2` is not smaller than the target, so index `2`
      might be the first match; set `right=2`.
    - After: `left=0`, `right=2`, `result=pending`.
    - Known correct: No index after `2` can be the first success.
2.  Inspect index `1`
    - Decision: Keep this possible first match by setting
      `right=1`.
    - After: `left=0`, `right=1`, `result=pending`.
    - Known correct: The lower bound is either unchecked index `0`
      or known-success index `1`.
3.  Inspect index `0`
    - Decision: `1 < 2`, so index `0` cannot qualify; set `left=1`.
    - After: `left=1`, `right=1`, `result=1`.
    - Known correct: Index `0` is too small and index `1` is the first
      value at least `2`.

For the same input and target, `UpperBound` should return `4`, the
first index whose value is greater than `2`.

Initial state: `left=0`, `right=5`, `result=pending`.

1.  Inspect index `2`
    - Decision: An upper bound must be strictly greater than `2`,
      so set `left=3`.
    - After: `left=3`, `right=5`, `result=pending`.
    - Known correct: Indexes `0` through `2` cannot be the answer;
      index `5` remains the allowed after-the-slice boundary.
2.  Inspect index `4`
    - Decision: Index `4` qualifies and may be first, so set
      `right=4`.
    - After: `left=3`, `right=4`, `result=pending`.
    - Known correct: No index after `4` can be the first answer.
3.  Inspect index `3`
    - Decision: Equality does not qualify, so set `left=4`.
    - After: `left=4`, `right=4`, `result=4`.
    - Known correct: Index `4` is the first value greater than `2`.

For the boundary case `values=[1,2,2,2,5]`, `target=6`, `LowerBound`
returns `5`. Every inspected value is smaller than 6, so `left`
advances to `len(values)`; `Search` checks that position before
indexing and returns `-1`.

#### Why it works

For a lower bound, every index before `left` is known to contain a
value smaller than the target. Index `right` and every later slice
index are known to contain a value at least as large as the target;
`right=len(values)` also represents no qualifying slice value. Those
facts remain true after either update because the values are sorted.
Each step makes the unchecked range shorter, so the loop finishes. At
`left == right`, the known-failure and known-success boundaries meet.
No earlier index qualifies, so that index is the first qualifying
position, or it equals `len(values)` when none exists. The upper-bound
proof is identical with "smaller than or equal to" as the failed-left
condition.

#### Go example

All three functions require `values` to be sorted in ascending order.

```go
package binarysearch

// LowerBound returns the first index whose value is greater than or
// equal to target. It returns len(values) when no such index exists.
func LowerBound(values []int, target int) int {
	// Only values in the half-open range [left,right) remain
	// unclassified. The boundaries converge on the insertion position,
	// which may be len(values).
	left, right := 0, len(values)
	for left < right {
		// This form avoids adding the two potentially large indexes
		// directly.
		middle := left + (right-left)/2
		if values[middle] < target {
			// middle and everything before it are too small.
			left = middle + 1
		} else {
			// middle may be the first qualifying index, so save it as the
			// right boundary.
			right = middle
		}
	}
	return left
}

// UpperBound returns the first index whose value is greater than
// target. It returns len(values) when no such index exists.
func UpperBound(values []int, target int) int {
	left, right := 0, len(values)
	for left < right {
		middle := left + (right-left)/2
		if values[middle] <= target {
			// Upper bound requires a strictly greater value, so discard
			// equality.
			left = middle + 1
		} else {
			// middle is greater and remains a possible first qualifying
			// position.
			right = middle
		}
	}
	return left
}

// Search returns the index of target in a sorted slice, or -1 when
// target is absent. When duplicates exist, it returns the first
// matching index.
func Search(values []int, target int) int {
	// LowerBound gives the only position that could be target's first
	// occurrence.
	index := LowerBound(values, target)
	// The lower bound may be one position past the final element. Check
	// for that result before indexing, then confirm exact equality.
	if index == len(values) || values[index] != target {
		return -1
	}
	return index
}
```

Concrete cases to test include empty input; targets before, within,
between, and after the data; `values=[1,2,2,2,5], target=2` producing
lower bound `1`, upper bound `4`, and search result `1`; and an absent
target producing `-1`.

#### Work and memory

Let `n` be the number of values. Each update removes at least half of
the remaining possible positions, so nonempty input takes
`O(log n)` time and empty input takes `O(1)`. `Search` adds one
constant-time check after `LowerBound`. The functions keep only index
variables and use `O(1)` extra memory without allocations.

#### When to use it

Use these functions when values are already sorted and the caller
needs an insertion point, an exact first match, a duplicate range, or
a count such as `UpperBound(values,x)-LowerBound(values,x)`.

#### When not to use it

Do not use the comparisons on unsorted input: for `values=[3,1,2]`,
discarding one side says nothing about values on that side. A map or
linear scan is often better for repeated exact lookups when sorting is
unavailable or forbidden.

#### Common mistakes and edge cases

> **Warning:** Changing the range to inclusive bounds without changing
> the updates can skip an index or loop forever. With this half-open
> range rule, initialize `right=len(values)` and never read
> `values[right]`.

Check `index == len(values)` before reading `values[index]`. Keep the
lower- and upper-bound equality rules distinct, especially for
duplicates. Computing `middle` as `left+(right-left)/2` also avoids
overflowing `left+right` in integer types where very large indexes are
possible.

#### Related material

- [C.1.3.6 Find a changing boundary (binary search)](#binary-search)
  explains when binary search applies and why its boundaries are safe.
- [A.2 Correctness and invariants](#a2-correctness-and-invariants)
  explains how the maintained facts prove the result.
- [Find First and Last Position of Element in Sorted Array](
  ../REFERENCE_LINKS.md#leetcode-0034
  )
  uses both bounds around a duplicate range.

<a id="d2-rolling-dynamic-programming"></a>

### E.2 Rolling dynamic programming

#### The problem it solves

Count ordered sequences of one-step and two-step moves that reach an
exact nonnegative total. Recomputing every smaller total recursively
repeats the same work, while storing every answer uses more memory
than this answer-building rule needs.

#### How it works

To count ways up a staircase, the final step came from one or two
stairs below. Add those two saved counts, then keep only the newest
two for the next stair. This is **rolling dynamic programming**. Each
saved count is a **state** because a later answer still depends on it.

#### Step-by-step examples

For `steps=4`, the expected answer is `5`.

Initial state: `current=0`, `twoBack=unset`, `oneBack=unset`,
`result=pending`.

1.  Establish totals `0` and `1`
    - Decision: There is one way to make `0` and one way to make `1`.
    - After: `current=2`, `twoBack=1`, `oneBack=1`,
      `result=pending`.
    - Known correct: The two base totals have exact counts.
2.  Build total `2`
    - Decision: Append a two-step move to the total-0 sequences or
      a one-step move to the total-1 sequences.
    - After: `current=3`, `twoBack=1`, `oneBack=2`,
      `result=pending`.
    - Known correct: The two sequences are `[2]` and `[1,1]`.
3.  Build total `3`
    - Decision: Add counts `1` and `2`.
    - After: `current=4`, `twoBack=2`, `oneBack=3`,
      `result=pending`.
    - Known correct: All sequences reaching `3` end in exactly one
      of the two allowed move sizes.
4.  Build total `4`
    - Decision: Add counts `2` and `3`, then roll forward.
    - After: `current=done`, `twoBack=3`, `oneBack=5`, `result=5`.
    - Known correct: All five ordered sequences are counted once.

For boundary inputs, `steps=-1` returns `0` and `steps=0` returns `1`
without entering the loop. The second result counts the one empty
sequence, not one physical move.

#### Why it works

Before each iteration, `twoBack` equals the answer for `current-2` and
`oneBack` equals the answer for `current-1`. Every sequence reaching
`current` has a unique final move of size one or two; removing that
move puts it in exactly one of those two groups. Their sum is
therefore the new exact count. The assignment restores the same fact
for the next total, and `current` increases until it reaches `steps`.

#### Go example

```go
package dynamicprogramming

// CountOrderedOneOrTwoStepWays returns the number of ordered
// sequences of one-step and two-step moves that sum to steps.
// Negative totals have no valid sequence; the empty sequence is the
// single way to make a total of zero. The result must fit in an int.
func CountOrderedOneOrTwoStepWays(steps int) int {
	// Positive moves cannot form a negative total.
	if steps < 0 {
		return 0
	}
	if steps < 2 {
		// There is one empty sequence for zero and one [1] sequence for
		// one.
		return 1
	}

	// Before each iteration these hold the counts for current-2 and
	// current-1.
	twoBack, oneBack := 1, 1
	for current := 2; current <= steps; current++ {
		// Every sequence ends in either a 1-step or 2-step move. Multiple
		// assignment evaluates both old values before rolling them forward.
		twoBack, oneBack = oneBack, twoBack+oneBack
	}
	return oneBack
}
```

Concrete cases to test are `-1 -> 0`, `0 -> 1`, `1 -> 1`, `2 -> 2`,
and `4 -> 5`, plus the largest accepted `steps` whose answer still
fits in `int`.

#### Work and memory

Negative, zero, and one return in `O(1)` time. For larger `steps`, the
loop runs `steps-1` times, so time is `O(steps)`. Two counts and one
loop index use `O(1)` extra memory, and the function makes no
allocations. This removes the `O(steps)` table allocation but not the
time needed to build each answer.

#### When to use it

Use rolling state when answers are built in one direction, only the
final answer is needed, and each new answer depends on a fixed number
of recent answers.

#### When not to use it

Keep a full table when reconstruction or later queries need older
states. Do not use this two-state recurrence when allowed move sizes
or other state make the next answer depend on more information.

#### Common mistakes and edge cases

> **Warning:** Treating total zero as having zero ways breaks later
> answers: `steps=2` would miss the sequence `[2]`. The empty sequence
> is the base way to make a zero total.

State the overflow limit; Go integer addition does not report
overflow. In a multiple assignment, right-hand expressions use the old
values, so `twoBack, oneBack = oneBack, twoBack+oneBack` performs the
intended roll.

#### Related material

- [C.1.3.10 Reuse smaller answers (dynamic
  programming)](#dynamic-programming) explains recurrence design and
  state selection.
- [A.3 Recursion and saved work](#a3-recursion-and-saved-work)
  compares saved answers with repeated recursive calls.
- [Climbing
  Stairs](problems/solved_easy_leetcode_0070_climbing_stairs.go) uses
  this recurrence.

<a id="d3-hash-map-complement-lookup"></a>

### E.3 Hash-map complement lookup

#### The problem it solves

Find two distinct positions whose integer values add to a target
without checking every pair. The returned positions must refer to
values already in the input, and one position cannot be used twice.

#### How it works

If two items must fit a `$20` budget, each price tells you the partner
price still needed. That **complement** is `target-currentPrice`. Ask
a map whether it appeared earlier, then store the current price and
its index for later items. The contract is
`indexByValue[value] -> earlier input index containing value`. Thus
`indexByValue[2]=0` means index `0` is an available earlier partner
whose value is `2`.

#### Step-by-step examples

For `values=[2,7,11,15]` and `target=9`, the expected result is
`(0,1,true)`.

Initial state: `index=0`, `indexByValue={}`, `result=pending`.

1.  Visit index `0`
    - Decision: Complement `7` is absent, so no earlier partner
      exists.
    - After: `index=1`, `indexByValue={2:0}`, `result=pending`.
    - Known correct: No pair ending at index `0` exists.
2.  Visit index `1`
    - Decision: Complement `2` is present at index `0`.
    - After: `index=1`, `indexByValue={2:0}`,
      `result=(0,1,true)`.
    - Known correct: The positions are distinct and `2+7=9`.

For the duplicate case `values=[3,3]`, `target=6`, index 0 is stored
before index 1 looks up its complement. The result is `(0,1,true)`;
looking up before inserting is what prevents index 0 from matching
itself on its first visit.

#### Why it works

Before visiting `index`, the map contains indexes from earlier
positions only. If the complement is present, its value plus the
current value equals the target and its index is necessarily distinct.
If it is absent, no earlier position can pair with the current value,
so remembering the current value is safe. The finite scan either
returns a valid pair or proves that no pair has a later endpoint in
the input.

#### Go example

```go
// Package hashmap demonstrates lookup-based uses of Go maps.
package hashmap

// PairSumIndices returns indices of two distinct values whose sum is
// target. The boolean is false when no pair exists. Subtracting each
// value from target must fit in an int.
func PairSumIndices(values []int, target int) (int, int, bool) {
	// The map contains only values from earlier indices.
	indexByValue := make(map[int]int)
	for index, value := range values {
		// complement is the exact value needed to complete target.
		complement := target - value
		earlierIndex, found := indexByValue[complement]
		if found {
			// earlierIndex and index are necessarily distinct because
			// insertion of the current value happens below.
			return earlierIndex, index, true
		}
		// Insert after lookup so an index cannot match itself.
		indexByValue[value] = index
	}

	return 0, 0, false
}
```

Concrete cases to test include `[2,7,11,15],9 -> (0,1,true)`,
`[3,3],6 -> (0,1,true)`, and `[1,2,4],8 -> false`. For the false
result, the two returned integers are placeholders and must not be
used.

#### Work and memory

In the usual hash-map cost model, each lookup and insertion takes
expected `O(1)` time, so the scan takes expected `O(n)` time. Go's
language rules do not promise a fixed worst-case time for map
operations. The map stores at most `n` entries and uses `O(n)`
additional memory. A successful early pair can reduce actual work.

#### When to use it

Use complement lookup when original indices matter, one pass is
useful, and `O(n)` extra map memory is acceptable.

#### When not to use it

Use the sorted two-pointer implementation in
[E.9](#e9-sorted-two-pointer-pair-search) when the input is already
sorted or may be reordered and constant extra memory matters.

#### Common mistakes and edge cases

> **Warning:** Inserting before lookup lets `values=[3]`, `target=6`
> return the same index twice. Perform the lookup first.

Use the map's boolean presence result rather than treating index zero
as absence. Decide whether any valid pair is acceptable because
repeated values can produce several answers. Widen or validate
arithmetic when `target-value` may overflow `int`.

#### Related material

- [C.1.3.4 Remember facts by key (hash map or set)](#hash-map-or-set)
  explains the lookup method.
- [D.4 Map presence and sets](#d4-map-presence-and-sets)
  explains the two-result lookup used here.
- [Two Sum](problems/solved_easy_leetcode_0001_two_sum.go) is the
  representative problem.

<a id="d4-generic-singly-linked-list"></a>

### E.4 Generic singly linked list

#### The problem it solves

Build a one-way linked list from ordinary values, convert it back for
inspection, and reverse its links without allocating replacement
nodes. The same operations should work for different stored value
types.

#### How it works

Each node is a card holding a value and a pointer to the next card. A
**singly linked list** follows only those forward links. A type
parameter `T` lets the node store a caller-chosen type, and reversing
the list turns each link around while preserving the remaining path
before it is overwritten.

#### Step-by-step examples

Build and reverse `values=[10,20,30]`; the expected final slice is
`[30,20,10]`.

Phase 1: build the list.

Initial state: `nextValue=10`, `list=empty`, `tail=dummy`.

1.  Append `10`
    - Decision: Create node `10` after `tail`; move `tail` to it.
    - After: `nextValue=20`, `list=10 -> nil`, `tail=10`.
    - Known correct: The built list equals the processed prefix
      `[10]`.
2.  Append `20`
    - Decision: Create node `20` after `tail`; move `tail` to it.
    - After: `nextValue=30`, `list=10 -> 20 -> nil`, `tail=20`.
    - Known correct: The built list equals `[10,20]`.
3.  Append `30` and return the real head
    - Decision: Append node `30`, then return `dummy.Next`.
    - After: `nextValue=done`, `list=10 -> 20 -> 30 -> nil`,
      `tail=30`.
    - Known correct: The dummy is excluded, and every input value
      appears once in order.

Phase 2: reverse the links.

Initial state: `previous=nil`, `current=10 -> 20 -> 30 -> nil`.

1.  Reverse node `10`
    - Decision: Save node `20`, set `10.Next=nil`, and advance.
    - After: `previous=10 -> nil`, `current=20 -> 30 -> nil`.
    - Known correct: Node `10` forms the complete reversed prefix,
      and node 20 still reaches every unprocessed node.
2.  Reverse node `20`
    - Decision: Save node `30`, set `20.Next=10`, and advance.
    - After: `previous=20 -> 10 -> nil`, `current=30 -> nil`.
    - Known correct: The first two original nodes are reversed
      without losing node 30.
3.  Reverse node `30`
    - Decision: Save `nil`, point `30` to `20`, and advance.
    - After: `previous=30 -> 20 -> 10 -> nil`, `current=nil`.
    - Known correct: Every original node appears once in reverse
      order.

Phase 3: convert the result for inspection.

Initial state: `node=30 -> 20 -> 10 -> nil`, `values=[]`.

1.  Read node `30`
    - Decision: Append `30`; advance to node `20`.
    - After: `node=20 -> 10 -> nil`, `values=[30]`.
    - Known correct: The slice equals the visited list prefix.
2.  Read node `20`
    - Decision: Append `20`; advance to node `10`.
    - After: `node=10 -> nil`, `values=[30,20]`.
    - Known correct: Both visited nodes appear in encounter order.
3.  Read node `10` and finish
    - Decision: Append `10`; advance to `nil` and stop.
    - After: `node=nil`, `values=[30,20,10]`.
    - Known correct: Every reversed-list node was visited once.

For an empty `values=[]`, `FromSlice` returns `nil`, `Reverse(nil)`
returns `nil`, and `ToSlice(nil)` returns an empty slice without
entering either loop.

#### Why it works

During construction, the list after `dummy` contains exactly the
values already visited, in their slice order, and `tail` is its last
node. During reversal, `previous` heads the fully reversed processed
prefix while `current` heads the untouched suffix. Saving `next`
before changing the link preserves the suffix; redirecting one link
extends the reversed prefix. Each iteration processes one node, so the
loop ends at `nil` with all nodes reversed.

#### Go example

```go
package linkedlist

// Node stores one value and a link to the next node.
type Node[T any] struct {
	Value T
	Next  *Node[T]
}

// FromSlice returns a list containing values in their slice order.
func FromSlice[T any](values []T) *Node[T] {
	// A dummy node gives both empty and nonempty construction one append
	// path; tail always names the node after which the next value is
	// attached.
	dummy := &Node[T]{}
	tail := dummy
	for _, value := range values {
		// Allocate one node, link it at the end, and advance tail to
		// preserve order.
		tail.Next = &Node[T]{Value: value}
		tail = tail.Next
	}
	// Exclude the construction-only dummy node from the returned list.
	return dummy.Next
}

// ToSlice returns the values encountered from head to nil.
func ToSlice[T any](head *Node[T]) []T {
	// Follow each Next link once and append values in encounter order.
	values := make([]T, 0)
	for node := head; node != nil; node = node.Next {
		values = append(values, node.Value)
	}
	return values
}

// Reverse reverses a list in place and returns its new head.
func Reverse[T any](head *Node[T]) *Node[T] {
	// previous is the reversed prefix; current begins the unreversed
	// remainder.
	var previous *Node[T]
	for current := head; current != nil; {
		// Save the only link to the remainder before redirecting
		// current.Next.
		next := current.Next
		current.Next = previous
		previous = current
		current = next
	}
	// When current reaches nil, previous contains every node in reverse
	// order.
	return previous
}
```

Concrete cases to test are the `[10,20,30]` round trip and reversal,
one node remaining unchanged, repeated values retaining their separate
nodes, and all three operations on an empty list.

#### Work and memory

For `n` values, `FromSlice` takes `O(n)` time and allocates `n` result
nodes; the temporary dummy may remain on the stack or be optimized by
the compiler. `ToSlice` takes `O(n)` time and `O(n)` result memory for
the returned slice. `Reverse` takes `O(n)` time, uses `O(1)` extra
memory, and allocates no nodes.

#### When to use it

Use this reference when a problem supplies singly linked nodes, needs
constant-time link changes after a known node, or teaches pointer
rewiring. The slice adapters are useful at test and input boundaries.

#### When not to use it

Prefer a slice when indexed access, cache-friendly traversal, or
compact storage matters. This node type has no backward link, so do
not use it when frequent movement toward previous nodes is required.

#### Common mistakes and edge cases

> **Warning:** Setting `current.Next=previous` before saving the old
> next link loses the unprocessed suffix. For `10 -> 20 -> 30`, node
> 20 becomes unreachable after the first rewrite.

Return `dummy.Next`, not `dummy`. Treat `nil` as a valid empty list. A
cycle violates `ToSlice` and `Reverse` contracts because their loops
would not reach `nil`; detect cycles separately when input is not
guaranteed acyclic.

#### Related material

- [B.2 Linked lists](#linked-lists) explains node operations and
  reversal reasoning.
- [D.7 Type parameters](#d7-type-parameters) explains `T` and `any`.
- [D.9 Slice storage and retained
  memory](#d9-slice-storage-and-retained-memory) contrasts contiguous
  slice storage with separately allocated nodes.

<a id="d5-generic-stack"></a>

### E.5 Generic stack

#### The problem it solves

Store values so that the most recently added value is the first one
removed, while allowing callers to inspect the top, count items, and
handle an empty stack without a panic.

#### How it works

Use the end of a slice as the top of a stack of trays. Adding and
removing at that one end implements **last in, first out (LIFO)**. A
generic `Stack[T]` gives the same operations to any chosen value type.

#### Step-by-step examples

Starting with the zero value `var stack Stack[string]`, push `"a"`,
push `"b"`, peek, then pop. The expected peek and pop value is `"b"`.

Initial state: `items=nil`, `next=push("a")`, `observed=[]`,
`savedValue=unset`.

1.  Push `"a"`
    - Decision: Append at the slice end.
    - After: `items=["a"]`, `next=push("b")`, `observed=[]`,
      `savedValue=unset`.
    - Known correct: `"a"` is the newest and top value.
2.  Push `"b"`
    - Decision: Append without disturbing the earlier item.
    - After: `items=["a","b"]`, `next=peek`, `observed=[]`,
      `savedValue=unset`.
    - Known correct: `"b"` is now the top.
3.  Peek
    - Decision: Read the final slot without slicing it away.
    - After: `items=["a","b"]`, `next=pop`,
      `observed=[("b",true)]`, `savedValue=unset`.
    - Known correct: Inspection did not change stack order or
      length.
4.  Save the top value
    - Decision: Save `items[1]` in `savedValue`.
    - After: `items=["a","b"]`, `next=clear`,
      `observed=[("b",true)]`, `savedValue="b"`.
    - Known correct: The value to return is preserved before its
      slot changes.
5.  Clear the old slot
    - Decision: Assign the string zero value `""` to slot `1`.
    - After: `items=["a",""]`, `next=shorten`,
      `observed=[("b",true)]`, `savedValue="b"`.
    - Known correct: The saved result remains `"b"`, and the old
      slot no longer retains it.
6.  Shorten and return
    - Decision: Reslice before index `1` and return the saved value.
    - After: `items=["a"]`, `next=done`,
      `observed=[("b",true),("b",true)]`, `savedValue="b"`.
    - Known correct: The newest value was removed first.

For an empty stack, both `Pop` and `Peek` return `"",false` for
`Stack[string]`; the boolean distinguishes failure from a successfully
stored empty string.

#### Why it works

The slice order is the insertion order and the final element is always
the most recent unremoved value. `Push` adds exactly there, while
`Peek` reads and `Pop` removes exactly there, so no older value can
leave first. Clearing a removed slot does not change the logical
order; it only stops that hidden array position from retaining a
reference. All operations finish after a bounded number of slice and
index operations, except for occasional append growth.

#### Go example

```go
package stack

// Stack stores values in last-in, first-out order. The zero value is
// ready to use.
type Stack[T any] struct {
	items []T
}

// Push places value on top of the stack.
func (stack *Stack[T]) Push(value T) {
	// The slice's final position is the stack top.
	stack.items = append(stack.items, value)
}

// Pop removes and returns the top value. The boolean is false when
// the stack is empty.
func (stack *Stack[T]) Pop() (T, bool) {
	if len(stack.items) == 0 {
		// Return T's zero value with false because there is no top element.
		var zeroValue T
		return zeroValue, false
	}

	lastItemIndex := len(stack.items) - 1
	value := stack.items[lastItemIndex]
	// Clear the hidden slot before shrinking so pointer-bearing T values
	// do not remain reachable through the backing array.
	var zeroValue T
	stack.items[lastItemIndex] = zeroValue
	stack.items = stack.items[:lastItemIndex]
	return value, true
}

// Peek returns the top value without removing it. The boolean is
// false when the stack is empty.
func (stack *Stack[T]) Peek() (T, bool) {
	if len(stack.items) == 0 {
		// As with Pop, false distinguishes absence from a stored zero
		// value.
		var zeroValue T
		return zeroValue, false
	}

	// Read the final item without changing the slice length.
	return stack.items[len(stack.items)-1], true
}

// Len returns the number of values in the stack.
func (stack *Stack[T]) Len() int {
	return len(stack.items)
}

// IsEmpty reports whether the stack contains no values.
func (stack *Stack[T]) IsEmpty() bool {
	return len(stack.items) == 0
}
```

Concrete cases to test include the string sequence above, integer zero
stored successfully versus an empty result, length changes after every
push and pop, complete removal, and repeated empty operations on the
zero-value stack.

#### Work and memory

`Peek`, `Pop`, `Len`, and `IsEmpty` take `O(1)` time. `Push` takes
`O(1)` amortized time. **Amortized** means averaged across a sequence:
most appends are constant work, while an occasional capacity growth
copies `O(n)` items. Storing `n` values uses `O(n)` memory. The
methods use `O(1)` extra memory beyond the slice storage.

#### When to use it

Use a stack when work must resume in reverse discovery order, such as
nested delimiter parsing, undo behavior, depth-first traversal, or
replacing recursion with explicit frames.

#### When not to use it

Use a queue when the oldest item must leave first. Use a priority
queue when a comparison, rather than insertion time, determines the
next item. Do not expose the backing slice when callers must not
bypass stack ordering.

#### Common mistakes and edge cases

> **Warning:** Returning only the zero value for an empty pop is
> ambiguous. An integer stack can validly contain `0`; check the
> boolean result.

Check for emptiness before using `len(items)-1`. Clear removed
pointer-bearing values before reslicing so the backing array does not
retain them. Do not assume every `Push` is worst-case constant time
because slice growth can copy the stored items.

#### Related material

- [C.1.3.13 Finish the newest open item first
  (stack)](#stack) explains stack-method selection.
- [D.7 Type parameters](#d7-type-parameters) explains the generic
  type.
- [D.9 Slice storage and retained
  memory](#d9-slice-storage-and-retained-memory) explains why `Pop`
  clears the removed slot.

<a id="d6-overflow-safe-integer-reversal"></a>

### E.6 Overflow-safe integer reversal

#### The problem it solves

Reverse the decimal digits of a nonnegative `int64` without first
producing a value outside the type's range. Checking only after
multiplication is too late because overflowing signed integer
arithmetic has already lost the intended value.

#### How it works

Take the last digit from the input and append it to the result, one
digit at a time. Before shifting the result left by one decimal place,
ask whether the shift and addition would still fit. This is a
**precondition overflow check**: it proves the next arithmetic
operation is safe before executing it.

#### Step-by-step examples

For `value=120`, the expected result is `(21,true)`.

Initial state: `value=120`, `reversed=0`, `result=pending`.

1.  Move digit `0`
    - Decision: The guard passes, so append `0`.
    - After: `value=12`, `reversed=0`, `result=pending`.
    - Known correct: The processed suffix `0` is represented as
      integer 0; its leading zero need not be retained.
2.  Move digit `2`
    - Decision: The guarded expression fits, so append `2`.
    - After: `value=1`, `reversed=2`, `result=pending`.
    - Known correct: The processed suffix `20` is reversed to `2`.
3.  Move digit `1`
    - Decision: Append the final digit.
    - After: `value=0`, `reversed=21`, `result=(21,true)`.
    - Known correct: Every input digit has moved to its reversed
      decimal position.

For the overflow case `value=8085774586302733229`, the digits
processed before the final leading `8` produce
`reversed=922337203685477580`. The guard compares this with
`(math.MaxInt64-8)/10 = 922337203685477579`, rejects the append, and
returns `(0,false)` without overflowing.

#### Why it works

Before each iteration, `reversed` is exactly the reverse of the digits
already removed from the original value, with insignificant leading
zeroes omitted. The remainder operator selects the next digit and
division removes that same digit. Because `digit` is between 0 and 9,
subtracting it from `MaxInt64` is safe; the guard is algebraically
equivalent to `reversed*10+digit <= MaxInt64` but does not perform the
risky multiplication. Each division removes one decimal digit, so the
loop finishes.

#### Go example

```go
package integerdigits

import "math"

// ReverseNonNegative reverses a nonnegative decimal integer. ok is
// false when value is negative or when its reversal would exceed
// int64.
func ReverseNonNegative(value int64) (reversed int64, ok bool) {
	// This implementation deliberately excludes a sign from digit
	// processing.
	if value < 0 {
		return 0, false
	}

	for value > 0 {
		// Remainder extracts the rightmost decimal digit.
		digit := value % 10
		// Rearranging reversed*10+digit <= MaxInt64 checks safety before
		// evaluating the multiplication that could overflow.
		if reversed > (math.MaxInt64-digit)/10 {
			return 0, false
		}
		// Append the extracted digit and remove it from the remaining
		// input.
		reversed = reversed*10 + digit
		value /= 10
	}
	return reversed, true
}
```

Concrete cases to test are `0 -> (0,true)`, `121 -> (121,true)`,
`120 -> (21,true)`, `-1 -> false`, and `8085774586302733229 -> false`
because its reversal is one greater than `math.MaxInt64`.

#### Work and memory

If `d` is the number of decimal digits, the loop takes `O(d)` time,
which is `O(log value)` for positive input. Zero and rejected negative
input take `O(1)` time. The function stores a fixed number of integers
and uses `O(1)` extra memory without allocations.

#### When to use it

Use this implementation when decimal digits must be processed as
numbers, negative input is not allowed, and overflow must be reported
rather than wrapped or clamped.

#### When not to use it

Use text reversal or an integer library that grows beyond `int64` when
every input digit must be preserved, including leading zeroes from
textual input. Derive a separate sign rule when negative values should
be reversed rather than rejected.

#### Common mistakes and edge cases

> **Warning:** Calculating `candidate := reversed*10+digit` and then
> comparing `candidate` with `MaxInt64` cannot detect the intended
> mathematical overflow: the `int64` operation has already wrapped
> before the comparison.

Zero skips the loop and is a successful result. Trailing input zeroes
become leading zeroes in the reversed digit sequence and disappear
from its integer representation. Keep the guard's subtraction before
division; comparing only `reversed > MaxInt64/10` misses the
last-digit tie case.

#### Related material

- [A.8 Integer digit processing](#integer-digit-processing) explains
  digit extraction and why the partially reversed value stays correct.
- [A.4 Integer limits and bits](#a4-integer-limits-and-bits) explains
  why range checks must precede unsafe arithmetic.
- [Palindrome
  Number](problems/solved_easy_leetcode_0009_palindrome_number.go)
  uses the same digit-extraction rule with a narrower input range.

<a id="d7-common-prefix-scan"></a>

### E.7 Common-prefix scan

#### The problem it solves

Find the longest starting byte sequence shared by every string. The
scan must stop safely when a shorter string ends and must define what
empty input means.

#### How it works

Read the strings as aligned rows and compare one column at a time. The
first unequal or missing byte ends the answer because a prefix cannot
skip a column. This is a **common-prefix scan**; this implementation
deliberately compares bytes rather than Unicode code points or
user-visible symbols.

#### Step-by-step examples

For `strs=["flower","flow","flight"]`, the expected result is `"fl"`.

Initial state: `column=0`, `expected='f'`, `prefix=""`,
`result=pending`.

1.  Compare column `0`
    - Decision: Rows `1` and `2` both contain `'f'`.
    - After: `column=1`, `expected='l'`, `prefix="f"`,
      `result=pending`.
    - Known correct: Every string shares prefix `"f"`.
2.  Compare column `1`
    - Decision: Rows `1` and `2` both contain `'l'`.
    - After: `column=2`, `expected='o'`, `prefix="fl"`,
      `result=pending`.
    - Known correct: Every string shares prefix `"fl"`.
3.  Compare column `2`
    - Decision: Row `2` contains `'i'`, so no prefix can include
      column 2.
    - After: `column=2`, `expected='o'`, `prefix="fl"`,
      `result="fl"`.
    - Known correct: The first two bytes match everywhere and the
      third does not.

For `strs=["","abc"]`, the outer loop has no columns and returns `""`.
For `strs=[]`, the explicit guard returns `""` before reading
`strs[0]`.

#### Why it works

Before each new column, every earlier byte has matched in every
string, so the slice before `column` is known to be common. A short
row or mismatch proves that no longer common prefix is possible.
Otherwise the newly checked byte can be added to the known prefix. The
scan ends at the first failure or after all bytes of the first string,
which is itself the maximum possible answer length.

#### Go example

```go
package commonprefixscan

// LongestCommonPrefix returns the longest byte prefix shared by all
// strings. It returns an empty string when strs is empty.
func LongestCommonPrefix(strs []string) string {
	// Empty input has no first row to use as the comparison baseline.
	if len(strs) == 0 {
		return ""
	}

	// Each byte in the first string is the expected value for one
	// column.
	for column := 0; column < len(strs[0]); column++ {
		expectedByte := strs[0][column]
		for row := 1; row < len(strs); row++ {
			// Check length before indexing. A missing or unequal byte proves
			// that the previously verified columns form the longest common
			// prefix.
			if column >= len(strs[row]) ||
				strs[row][column] != expectedByte {
				return strs[0][:column]
			}
		}
	}
	// Every byte of the first string matched, so no longer prefix is
	// possible.
	return strs[0]
}
```

Concrete cases to test include the `"flower"` example;
`["dog","racecar"]` producing `""`; `["ab","a"]` producing `"a"`;
identical strings; one string; an empty first or later string; and no
strings.

#### Work and memory

If the first string has `m` bytes and there are `n` strings, the worst
case does `O(n*m)` byte comparisons. A mismatch can return earlier,
including in constant time at the first column. The function uses
`O(1)` extra memory. The returned substring reuses the first
string's bytes rather than allocating a copied result.

#### When to use it

Use this implementation when the problem compares byte prefixes, all
strings must participate, and returning a substring of the first input
is acceptable.

#### When not to use it

Do not use byte columns when the task defines positions as Unicode
code points or visible symbols; use the representation described in
[D.6](#d6-bytes-runes-and-visible-characters). A trie may be more
useful when many changing queries share a large stored dictionary.

#### Common mistakes and edge cases

> **Warning:** Reading `strs[row][column]` before checking
> `column < len(strs[row])` panics for `["ab","a"]` at column 1.

Guard empty input before reading `strs[0]`. The first string may be
the shortest without being specially selected. Byte slicing can split
a multi-byte UTF-8 encoding if a different string disagrees inside
that encoding, which is why the byte-based comparison rule must be
explicit.

#### Related material

- [C.1.3.1 Compare from the left (common-prefix
  scan)](#common-prefix-scan) explains the method and why the known
  prefix stays correct.
- [D.6 Bytes, runes, and visible
  characters](#d6-bytes-runes-and-visible-characters) explains how to
  choose the unit of comparison.
- [Longest Common
  Prefix](problems/solved_easy_leetcode_0014_longest_common_prefix.go)
  is the representative problem.

<a id="d8-directional-scan"></a>

### E.8 Directional scan

#### The problem it solves

Apply a local signed-sum rule in which each value is subtracted when
it is smaller than its immediate right neighbor and added otherwise.
Each decision needs only one neighbor, but that neighbor must remain
the original adjacent value rather than the running total.

#### How it works

If each train car needs the original value from the car on its right,
inspect cars from right to left and carry that value with you. A
**directional scan** starts where the needed context is already
available. Here `rightValue` always means the untouched neighbor
needed for the next comparison.

#### Step-by-step examples

For `values=[3,4,2]`, the expected result is `3` because the signed
expression is `-3+4+2`.

Initial state: `index=2`, `total=0`, `rightValue=none`,
`result=pending`.

1.  Establish the right edge
    - Decision: Value `2` has no right neighbor, so add it.
    - After: `index=1`, `total=2`, `rightValue=2`,
      `result=pending`.
    - Known correct: Index `2` has its final contribution.
2.  Visit index `1`
    - Decision: `4` is not smaller than `2`, so add it.
    - After: `index=0`, `total=6`, `rightValue=4`,
      `result=pending`.
    - Known correct: Indices `1` and `2` have their final
      contributions, and `rightValue` is the original value at index
      1.
3.  Visit index `0`
    - Decision: `3 < 4`, so subtract it.
    - After: `index=done`, `total=3`, `rightValue=3`, `result=3`.
    - Known correct: Every index was classified against its
      immediate original right neighbor.

For empty input, the result is zero. For `values=[7]`, initialization
produces `total=7`, the loop performs no iterations, and the result is
7.

#### Why it works

Before each loop iteration, `total` contains the final signed
contributions of all positions to the right of `index`, and
`rightValue` is exactly the original value at `index+1`. The
comparison therefore applies the required local rule to the current
position. Assigning `rightValue=value` establishes the same fact one
position farther left. The index decreases to zero, so all positions
are processed once and the final sum follows the rule.

#### Go example

```go
package directionalscan

// SumWithSmallerLeftValuesSubtracted applies a local signed-sum rule:
// a value is subtracted when it is smaller than its immediate right
// neighbor and added otherwise. The running total must fit in an int.
func SumWithSmallerLeftValuesSubtracted(values []int) int {
	if len(values) == 0 {
		return 0
	}

	// The final value has no right neighbor, so it is always added.
	// rightValue carries that original neighbor value into the next
	// iteration.
	last := len(values) - 1
	total := values[last]
	rightValue := values[last]

	for index := last - 1; index >= 0; index-- {
		value := values[index]
		// Apply the local rule against the immediate original right
		// neighbor, not against the accumulated total.
		if value < rightValue {
			total -= value
		} else {
			total += value
		}
		// Shift the neighbor state one position left for the next
		// comparison.
		rightValue = value
	}

	return total
}
```

Concrete cases to test are `[] -> 0`, `[7] -> 7`, `[5,4,3] -> 12`,
`[3,4] -> 1`, and `[3,4,2] -> 3`. Also test values near the accepted
integer limits while ensuring every running total still fits in
`int`; this function does not report overflow outside that contract.

#### Work and memory

The function visits each of `n` values once, so it takes `O(n)` time;
empty and single inputs are still within that bound. It stores a fixed
number of integers, uses `O(1)` extra memory, and makes no
allocations.

#### When to use it

Use a directional scan when each decision depends on a fixed nearby
value that can be carried from one position to the next and the output
can be accumulated without retaining all intermediate states.

#### When not to use it

Do not use this exact local rule for encodings whose subtraction
decision depends on a larger token or on a reduced value to the right.
Use explicit parsing state when several prior symbols determine one
unit.

#### Common mistakes and edge cases

> **Warning:** Assigning `rightValue=total` compares the next value
> with an accumulated sum instead of its neighbor. For `[3,4,2]`, that
> mistake can classify index 0 against 6 rather than 4.

Handle empty input before calculating `len(values)-1`. State what
happens at the last item. Handle equality explicitly: equal adjacent
values are added because only a strictly smaller left value is
subtracted.

#### Related material

- [C.1.3.2 Process the needed neighbor first (directional
  scan)](#directional-scan) explains method selection and direction
  choice.
- [A.2 Correctness and invariants](#a2-correctness-and-invariants)
  explains the maintained `rightValue` fact.
- [Roman to
  Integer](problems/solved_easy_leetcode_0013_roman_to_integer.go)
  uses a related right-neighbor decision under Roman-numeral input
  rules.

<a id="d9-sorted-two-pointer-pair-search"></a>

### E.9 Sorted two-pointer pair search

#### The problem it solves

Find two distinct positions in an ascending sorted integer slice whose
values sum to a target, using no map. A full pair-by-pair search does
quadratic work even though the ordering can rule out many pairs at
once.

#### How it works

Two people check prices from opposite ends of a sorted shelf. If their
sum is too small, move the low-price side up; if too large, move the
high-price side down. These coordinated integer indexes form the
**two-pointer method**; they are not Go memory pointers.

#### Step-by-step examples

For `values=[-3,1,2,5,8]` and `target=7`, the expected result is
`(2,3,true)`.

Initial state: `left=0`, `right=4`, `result=pending`.

1.  Check indexes `0` and `4`
    - Decision: Sum `5` is too small, so increment `left`.
    - After: `left=1`, `right=4`, `result=pending`.
    - Known correct: Index `0` cannot pair with any remaining value
      because 8 was already its largest possible partner.
2.  Check indexes `1` and `4`
    - Decision: Sum `9` is too large, so decrement `right`.
    - After: `left=1`, `right=3`, `result=pending`.
    - Known correct: Index `4` cannot pair with any remaining value
      because 1 was already its smallest possible partner.
3.  Check indexes `1` and `3`
    - Decision: Sum `6` is too small, so increment `left`.
    - After: `left=2`, `right=3`, `result=pending`.
    - Known correct: Index `1` has no valid remaining partner.
4.  Check indexes `2` and `3`
    - Decision: Sum `7` equals the target.
    - After: `left=2`, `right=3`, `result=(2,3,true)`.
    - Known correct: The positions are distinct and their values
      sum to 7.

For `values=[4]`, `target=8`, `left=0` and `right=0`, so the loop
never runs and returns `false`; one position cannot be used twice.
Empty input also returns `false` without indexing the slice.

#### Why it works

If the current sum is too small, pairing `values[left]` with any index
below `right` can only make the sum no larger, so discarding `left`
cannot remove a solution. The symmetric argument safely discards
`right` when the sum is too large. Thus every move preserves any
possible solution inside the remaining range. One pointer moves on
every unsuccessful iteration, so they eventually meet or a valid pair
is returned.

#### Go example

```go
package twopointers

// PairSumSorted returns indices of two values whose sum equals
// target. Values must be sorted in ascending order, and relevant sums
// must fit in an int.
func PairSumSorted(
	values []int,
	target int,
) (left int, right int, ok bool) {
	// Begin with the smallest and largest available values.
	left, right = 0, len(values)-1
	for left < right {
		sum := values[left] + values[right]
		switch {
		case sum < target:
			// Even the largest partner was too small for this left value, so
			// no pair using left can succeed.
			left++
		case sum > target:
			// Even the smallest partner was too large for this right value, so
			// no pair using right can succeed.
			right--
		default:
			// left < right guarantees two distinct positions.
			return left, right, true
		}
	}

	return 0, 0, false
}
```

Concrete cases to test include the negative-value trace above;
`[1,2,4,6],8 -> (1,3,true)`; `[4],8 -> false`; `[1,2,3],7 -> false`;
duplicate values such as `[3,3],6`; and empty input.

#### Work and memory

Each unsuccessful iteration moves one pointer inward, so there are at
most `n-1` iterations and `O(n)` time. An early match may take `O(1)`
actual time. The function uses `O(1)` extra memory and makes no
allocations. Sorting an unsorted input first would add the sorting
time and any copy required to preserve caller order.

#### When to use it

Use this implementation when the values are already sorted, one
matching pair is enough, and constant extra memory or ordered
elimination is preferable to a map.

#### When not to use it

Do not apply it directly to unsorted values. Sorting may also be wrong
when the original indices must be returned and are not carried through
the sort; the map implementation in
[E.3](#e3-hash-map-complement-lookup) preserves those indexes in one
pass.

#### Common mistakes and edge cases

> **Warning:** Using `left <= right` allows one position to pair with
> itself. For `[4]`, target 8 would incorrectly succeed; require
> `left < right`.

The ordering proof fails if the slice is descending or unsorted.
Include duplicates, negative values, empty input, and missing pairs in
tests. Widen or validate addition when two accepted `int` values might
overflow.

#### Related material

- [C.1.3.9 Move coordinated positions (two pointers)](#two-pointers)
  explains the same ordered discard proof.
- [E.3 Hash-map complement lookup](#e3-hash-map-complement-lookup) is
  the unsorted one-pass alternative.
- [D.3 Sorting without unexpected
  mutation](#d3-sorting-without-unexpected-mutation) explains the cost
  of preparing an unsorted input without reordering it.

<a id="f-further-study-checklist"></a>

## F. Further study

The topics below are optional. The rest of this guide does not require
them. Study one when a problem, contest, or job calls for it. Follow
the linked starting point first if the new topic is unfamiliar.

- **Strongly connected components** find groups in a graph whose
  connections have directions. Within one group, every node can reach
  every other node by following those directions. Start with
  [depth-first search](#c1320-depth-first-search) and graph direction
  in [graphs and grids](#b5-graphs-and-grids).
- **All-pairs shortest paths** find the cheapest route between every
  pair of nodes, rather than routes from only one starting node.
  For nonnegative costs, compare repeated [Dijkstra
  searches](#c1323-dijkstra-shortest-paths) with [dynamic
  programming](#c1310-dynamic-programming) that gradually allows more
  intermediate nodes on a route.
- **Maximum flow** asks how much material can travel from a starting
  point to a destination. Each one-way connection has a **capacity**,
  meaning the largest amount it can carry. Start with the way
  [graphs and grids](#b5-graphs-and-grids) store each node's
  neighbors.
- **Suffix arrays** organize every suffix of a string. A suffix is a
  piece that starts at some position and continues to the end.
  **Suffix automata** compactly represent the string's substrings.
  Both can speed up repeated questions about text inside one string.
  Begin with [string matching](#c137-string-matching) and [parsing
  state](#c138-parsing-and-state-machines).
- **Computational geometry beyond normalized slopes** uses exact
  arithmetic instead of rounded decimal values. It can decide which
  side of a line contains a point, whether shapes cross, and whether a
  point lies on a boundary. Start with [exact slopes without
  rounding](#a63-exact-slopes-without-rounding).
- **Randomized balanced search trees** keep values in order. They use
  random priorities to make a long, one-sided shape unlikely. Search
  and update usually need only a number of steps proportional to the
  tree's height. Compare
  [binary-search-tree ordering](#b41-terminology-and-search) with
  [heap priorities](#c1327-heap-or-priority-queue).
