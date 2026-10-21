# Algorithms and Data Structures in Go

This is the repository's guide to understanding algorithms and data structures
in Go. It is written for readers who know basic Go syntax--variables, loops,
functions, slices, and maps--but are new to choosing algorithms and explaining
why they work. The problem archives contain complete judge solutions; this
guide teaches the reusable ideas behind those solutions.

## Contents

- [How to use this guide](#how-to-use-this-guide)
- [Contemporary interview evidence](#contemporary-interview-evidence)
- [A. Foundations](#a-foundations)
  - [A.1 Complexity and resource accounting](#a1-complexity-and-resource-accounting)
  - [A.2 Correctness and invariants](#a2-correctness-and-invariants)
  - [A.3 Recursion and saved work](#a3-recursion-and-saved-work)
  - [A.4 Integer limits and bits](#a4-integer-limits-and-bits)
  - [A.5 Bit manipulation](#a5-bit-manipulation)
  - [A.6 Keeping arithmetic values exact and
    consistent](#a6-keeping-arithmetic-values-exact-and-consistent)
  - [A.7 Combinatorics and output size](#a7-combinatorics-and-output-size)
  - [A.8 Integer digit processing](#a8-integer-digit-processing)
- [B. Ways to store and organize information](#b-ways-to-store-and-organize-information)
  - [B.1 Arrays, slices, and strings](#b1-arrays-slices-and-strings)
  - [B.2 Linked lists](#b2-linked-lists)
  - [B.3 Queues and deques](#b3-queues-and-deques)
  - [B.4 Trees and binary-search trees](#b4-trees-and-binary-search-trees)
  - [B.5 Graphs and grids](#b5-graphs-and-grids)
  - [B.6 Matrix traversal and in-place state](#b6-matrix-traversal-and-in-place-state)
  - [B.7 Composing data structures](#b7-composing-data-structures)
  - [B.8 Streaming and online processing](#b8-streaming-and-online-processing)
  - [B.9 Choosing a structure](#b9-choosing-a-structure)
- [C. Ways to solve problems](#c-ways-to-solve-problems)
  - [C.1 Problem-solving method selection guide](#c1-problem-solving-method-selection-guide)
    - [C.1.1 Ask what the input and output allow (quick decision
      tree)](#c11-ask-what-the-input-and-output-allow-quick-decision-tree)
    - [C.1.2 Derive a direction before naming a
      pattern](#c12-derive-a-direction-before-naming-a-pattern)
    - [C.1.3 Pattern cards](#c13-pattern-cards)
    - [C.1.4 Choosing among overlapping methods](#c14-choosing-among-overlapping-methods)
    - [C.1.5 Let constraints eliminate bad choices](#c15-let-constraints-eliminate-bad-choices)
    - [C.1.6 Before committing to a method](#c16-before-committing-to-a-method)
- [D. Go techniques](#d-go-techniques)
  - [D.1 Keep input conversion outside the
    algorithm](#d1-keep-input-conversion-outside-the-algorithm)
  - [D.2 Table-driven tests](#d2-table-driven-tests)
  - [D.3 Sorting without unexpected mutation](#d3-sorting-without-unexpected-mutation)
  - [D.4 Distinguish missing keys from stored values (map presence and
    sets)](#d4-map-presence-and-sets)
  - [D.5 Priority queues](#d5-priority-queues)
  - [D.6 Bytes, runes, and visible characters](#d6-bytes-runes-and-visible-characters)
  - [D.7 Type parameters](#d7-type-parameters)
  - [D.8 Replacing recursion with a loop](#d8-replacing-recursion-with-a-loop)
  - [D.9 Slice storage and retained memory](#d9-slice-storage-and-retained-memory)
- [E. Go reference implementations](#e-go-reference-implementations)
  - [E.1 Binary search bounds](#e1-binary-search-bounds)
  - [E.2 Rolling dynamic programming](#e2-rolling-dynamic-programming)
  - [E.3 Hash-map complement lookup](#e3-hash-map-complement-lookup)
  - [E.4 Generic singly linked list](#e4-generic-singly-linked-list)
  - [E.5 Generic stack](#e5-generic-stack)
  - [E.6 Overflow-safe integer reversal](#e6-overflow-safe-integer-reversal)
  - [E.7 Common-prefix scan](#e7-common-prefix-scan)
  - [E.8 Directional scan](#e8-directional-scan)
  - [E.9 Sorted two-pointer pair search](#e9-sorted-two-pointer-pair-search)
- [F. Further study](#f-further-study)

## How to use this guide

Choose the path that matches the work you are doing:

- **First-time study:** Read foundations, then data structures, then
  problem-solving methods. Read a Go technique when an entry links to it.
- **Choosing a solution:** Begin with the
  [method-selection guide](#c1-problem-solving-method-selection-guide), name
  the shape of the input and output, and follow its links to at most two likely
  methods.
- **Quick reference:** Jump to a named entry, replay its examples, and check
  its "When not to use it" section before adapting its code.

Main parts use letters, and nested sections add numbers. For example, `C`
identifies problem-solving methods, `C.1` identifies the method-selection
guide, and `C.1.3` identifies the pattern-card collection within that guide.

### Plain English before technical names

The guide keeps names such as *binary search*, *dynamic programming*, and
*Big-O* because those names appear in problems, documentation, and interviews.
It first explains the action or behavior in ordinary language, then introduces
and defines the conventional name. A technical term should never be required
to understand its own definition.

### Entry contract

Every teaching entry uses these headings in this order. A heading can be brief,
but it is not omitted merely because the topic is small:

1. **The problem it solves** states the input and desired result before naming
   the technique as the answer.
2. **The idea in simple terms** gives a mental picture, then introduces and
   defines the formal name.
3. **How it works** states required input conditions and the exact state the
   procedure keeps.
4. **Step-by-step examples** traces an ordinary input and, when behavior
   changes, a boundary or adversarial input.
5. **Why it works** identifies what remains true, why each step is safe, why
   the process finishes, and why the final state is the requested result.
6. **Go example** shows the smallest complete fragment or links to the one
   owning reference implementation.
7. **Work and memory** explains the costs in words before using notation such
   as `O(n)`.
8. **When to use it** lists clues visible in the problem or input.
9. **When not to use it** names the assumption or cost that rules it out.
10. **Common mistakes and edge cases** ties every warning to a concrete input,
    state change, or Go behavior.
11. **Related material** links to the owner of prerequisite, contrasting, and
    implementation details instead of repeating them.

The labels are Markdown headings one level below the entry heading. For
example, an `###` entry uses `#### The problem it solves`; a deeply nested
`#####` pattern card uses `###### The problem it solves`.

### Forms used in every entry

- **Trace:** Use a numbered list. For repeated state changes, give each step
  labeled `Before`, `Decision`, `After`, and `Known correct` details. For
  comparisons or mappings, use concise bullets. Do not create Markdown tables.
- **Warning:** Write `> **Warning:**` followed by the failing input or state and
  the correction. A warning without a concrete failure is incomplete.
- **Code:** Use a fenced `go` block. Names in the code match the trace, and any
  input requirement appears immediately before the block. When the complete
  implementation has another owner, link to it rather than copying it.
- **Related links:** End with a short bulleted list. Each item says why the
  linked entry matters; it does not duplicate that entry's explanation.

The contemporary interview evidence below gives context for what to practice;
it is not a teaching entry and does not use the eleven-heading contract.

The coverage is cross-checked against LeetCode's
[Top Interview 150](https://leetcode.com/studyplan/top-interview-150/), the
[Grind 75 study guidance](https://www.techinterviewhandbook.org/grind75/faq),
and the [NeetCode roadmap](https://neetcode.io/roadmap). These lists determine
breadth, but the material below is organized by reusable invariants rather than
by problem title.

## Contemporary interview evidence

Current official preparation material from
[Amazon](https://www.amazon.jobs/content/en/how-we-hire/interview-prep/software-development-topics),
[Meta](https://d3no4ktch0fdq4.cloudfront.net/public/course/files/Meta_SWE_full_loop_guide.pdf),
and [Microsoft](https://careers.microsoft.com/v2/global/en/hiring-tips/technical-interviewing.html)
emphasizes applying fundamentals, comparing approaches, explaining complexity,
writing clear code, and testing it. That is stronger evidence about the skills
being assessed than any remembered list of exact questions.

Recent public candidate reports are also sampled for changes in question style,
including reports from [Amazon in 2026](https://leetcode.com/discuss/post/8420230/amazon-sde-2-interview-experience-2026-e-t1g8/),
[Google in 2025](https://leetcode.com/discuss/post/6522571/),
[Meta in 2025](https://leetcode.com/discuss/post/6460774/Meta-or-E4-SWE-or-USA-or-Feb-2025-or-Passed-Hiring-Committee/),
[Microsoft in 2025](https://leetcode.com/discuss/post/7227003/), and
[Uber in 2025](https://leetcode.com/discuss/post/7327020/). Candidate reports
are anecdotal and cannot prove that one exact question is common. This guide
therefore incorporates only repeated, transferable signals:

- Expect a familiar structure with a changed constraint: duplicates, a limited
  number of edits, streaming input, a value-distance restriction, or a demand
  for lower memory.
- Be able to move from a direct solution to a better one and explain what
  repeated work, one-direction ordering, or data-structure operation made the
  improvement possible.
- Practice constructing a graph from relationships, converting a rooted tree
  into an undirected graph when movement may go upward, and recognizing when a
  range query accelerates a dynamic-programming transition.
- Know online combinations such as counts plus a queue, two heaps, a heap over
  sorted streams, and a map plus linked list. The invariant between the two
  structures matters more than memorizing a class name.
- Treat manual verification as part of the solution: trace one ordinary case,
  one boundary case, and one adversarial case before declaring the code done.

Exact company question banks change and are often confidential. Use company
tags to choose practice inputs, not as a substitute for the concepts below.

## A. Foundations

### A.1 Complexity and resource accounting

#### The problem it solves

Two programs can return the same answer but require very different amounts of
work or memory as the input grows. We need a machine-independent way to reject
an approach that cannot finish within the input limits and to account for
memory that is easy to overlook.

#### The idea in simple terms

Imagine finding a name in a sorted 1,024-page phone book. Reading every page
can require 1,024 checks; repeatedly discarding half needs at most about 10.
The important question is how the work changes when the book doubles, not the
exact number of nanoseconds on one computer.

Let `n` mean an input amount, such as the number of slice elements. **Time
complexity** describes how work grows with `n`. **Big-O notation** gives an
upper growth bound, usually written with the tightest simple bound that is
useful. **Auxiliary-space complexity** describes memory created by the method,
normally excluding the given input and required returned output.

#### How it works

1. Define what `n` measures and which operation represents the main work.
2. Count how often that operation can occur. Sequential loops usually add;
   fully nested loops usually multiply.
3. Keep the fastest-growing term and omit fixed multipliers. Thus
   `3n + 8` grows linearly and is written `O(n)`.
4. Analyze different input cases only when control flow changes their cost.
5. Separately count live allocations, unfinished calls, preprocessing, and
   output construction.

Common growth classes are `O(1)` for bounded work, `O(log n)` for repeatedly
shrinking by a fixed fraction, `O(n)` for a full pass, `O(n log n)` for many
efficient sorts, `O(n^2)` for all pairs, `O(2^n)` for all subsets, and `O(n!)`
for all orderings.

#### Step-by-step examples

Search for `9` in the sorted values `[1,3,5,7,9,11,13,15]`.

1. **Method and step: Linear check `0`**
   - **Before:** Eight possible positions
   - **Decision:** `1 != 9`; advance one position.
   - **After:** Seven unchecked positions
   - **Work known so far:** 1 comparison
2. **Method and step: Linear checks `1..3`**
   - **Before:** Positions `1..7` remain
   - **Decision:** `3`, `5`, and `7` differ.
   - **After:** Positions `4..7` remain
   - **Work known so far:** 4 comparisons
3. **Method and step: Linear check `4`**
   - **Before:** Value `9` is current
   - **Decision:** Match and stop.
   - **After:** Return `true`.
   - **Work known so far:** 5 comparisons; a miss could require all 8
4. **Method and step: Halving check `3`**
   - **Before:** Search interval is `0..7`
   - **Decision:** `7 < 9`; positions `0..3` cannot contain `9`.
   - **After:** Interval becomes `4..7`.
   - **Work known so far:** 1 comparison
5. **Method and step: Halving check `5`**
   - **Before:** Interval is `4..7`
   - **Decision:** `11 > 9`; positions `5..7` cannot contain `9`.
   - **After:** Interval becomes `4..4`.
   - **Work known so far:** 2 comparisons
6. **Method and step: Halving check `4`**
   - **Before:** Interval is `4..4`
   - **Decision:** `9 == 9`; stop.
   - **After:** Return `true`.
   - **Work known so far:** 3 comparisons

For empty input, a scan performs zero element comparisons and a halving search
also stops immediately. That best case does not change the scan's worst-case
`O(n)` bound or the halving search's worst-case `O(log n)` bound.

#### Why it works

The count follows the program's control flow. A scan advances past at most one
of `n` positions per comparison, so it performs no more than `n` comparisons.
A halving search leaves at most half of the previous candidates after every
failed comparison, so after `k` failures at most `n/2^k` candidates remain;
the search stops once that quantity falls below one. Solving for `k` gives a
logarithmic bound.

Dropping constants and slower-growing terms preserves the comparison for large
inputs: multiplying every input size by the same fixed amount does not change
whether doubling `n` doubles, quadruples, or barely changes the work.

#### Go example

This scan makes the counted operation--one equality check per visited
element--visible:

```go
func contains(values []int, target int) bool {
	// Visit values from left to right. Each iteration performs the equality
	// check whose count determines this linear scan's running time.
	for _, value := range values {
		if value == target {
			// A match is enough to answer the question, so stop immediately.
			return true
		}
	}
	// Reaching the end means every element was checked without a match.
	return false
}
```

#### Work and memory

If the target is first, the function performs one comparison: best-case
`O(1)` time. A typical successful position still grows with the slice, and a
missing or last target requires `n` comparisons: average and worst-case
`O(n)` time. It stores only the current value, so auxiliary space is `O(1)`.
The input slice is borrowed and no output collection is allocated.

#### When to use it

Use growth analysis whenever choosing among solutions, checking a constraint,
or reviewing a hidden allocation. Define the input measure explicitly when
there are several dimensions, such as rows `r`, columns `c`, and output count
`k`.

#### When not to use it

Do not use a growth class as an exact speed claim for small inputs. Constants,
cache behavior, allocation, and input distribution can make two methods with
the same class perform differently. Measure real performance after the growth
rate is acceptable.

#### Common mistakes and edge cases

> **Warning:** Two consecutive loops over `n` values cost `n + n`, or `O(n)`,
> not `O(n^2)`. The quadratic result applies when one full loop runs inside
> every step of another.

> **Warning:** A recursive function with no explicit slice can still use
> memory. `n` unfinished calls require `O(n)` call-stack space.

Count sorting or map construction before a later fast query, and count copied
output elements when output is part of the required work. For rectangular data,
do not call an `r * c` traversal `O(n)` unless `n` has been defined as the total
cell count.

#### Related material

- [A.2 Correctness and invariants](#a2-correctness-and-invariants) separates
  proving the answer from measuring its cost.
- [A.3 Recursion and saved work](#a3-recursion-and-saved-work) owns call-stack
  accounting.
- [A.7 Combinatorics and output size](#a7-combinatorics-and-output-size) shows
  why required output can dominate every other cost.
- [E.1 Binary search bounds](#e1-binary-search-bounds) provides the complete
  halving-search implementation used in the comparison.

### A.2 Correctness and invariants

#### The problem it solves

A program can return the expected result for a few examples and still be wrong.
To establish correctness, we need to connect every state-changing step to the
required output and show that the process cannot skip an answer or run forever.

#### The idea in simple terms

Imagine tidying toys from left to right while keeping this promise:
"everything left of my hand is already in the right place." The promise is a
**fact that remains true** while the work proceeds. The formal name for such a
fact is an **invariant**. A useful invariant describes exactly which part of the
result is already correct.

#### How it works

Build a correctness argument around four questions:

1. **Start:** Why is the invariant true before the first step?
2. **Preservation:** If it is true before one step, why is it still true after
   that step?
3. **Progress and finish:** What moves closer to a fixed limit, so the process
   must stop?
4. **Result:** When it stops, why does the invariant cover the complete answer?

Also justify every discarded choice. Saying only "the loop has checked the
left side" is too vague; name what was checked and what result that state
proves.

#### Step-by-step examples

Move the non-zero values in `values = [0, 5, 0, 2]` to the front without
changing their order. The expected returned prefix is `[5, 2]`. `read` is the
position being inspected, and `write` is the length of the correct prefix.

1. **Step: Start**
   - **Before:** `values=[0,5,0,2]`, `write=0`
   - **Decision:** No values have been read.
   - **After:** Unchanged.
   - **Known correct:** `values[:0]=[]` contains all non-zero values read so far.
2. **Step: `read=0`, value `0`**
   - **Before:** `values=[0,5,0,2]`, `write=0`
   - **Decision:** Skip `0` because it is not part of the result.
   - **After:** `values=[0,5,0,2]`, `write=0`
   - **Known correct:** The correct prefix is still `[]`.
3. **Step: `read=1`, value `5`**
   - **Before:** `values=[0,5,0,2]`, `write=0`
   - **Decision:** Write `5` at position `0`, then advance `write`.
   - **After:** `values=[5,5,0,2]`, `write=1`
   - **Known correct:** `values[:1]=[5]` is the non-zero input seen so far, in order.
4. **Step: `read=2`, value `0`**
   - **Before:** `values=[5,5,0,2]`, `write=1`
   - **Decision:** Skip `0`.
   - **After:** `values=[5,5,0,2]`, `write=1`
   - **Known correct:** The correct prefix remains `[5]`.
5. **Step: `read=3`, value `2`**
   - **Before:** `values=[5,5,0,2]`, `write=1`
   - **Decision:** Write `2` at position `1`, then advance `write`.
   - **After:** `values=[5,2,0,2]`, `write=2`
   - **Known correct:** `values[:2]=[5,2]` is the complete result.

For the boundary input `[0, 0]`, both steps skip their value. `write` remains
`0`, so the returned prefix is `values[:0]`, or `[]`, as required. The
invariant describes the empty result without a special case.

#### Why it works

Before the loop, no values have been read, so the empty prefix contains exactly
the non-zero values seen so far. During a step, a zero is safely skipped because
it does not belong in the result. A non-zero value is written immediately after
the already-correct prefix, so the prefix gains that value without changing the
order of earlier values. The fact therefore remains true after either action.

`read` advances once per input value and eventually reaches the fixed input
length, so the loop finishes. At that point every input value has been read;
the invariant says `values[:write]` contains every non-zero input value in its
original order, which is the requested output.

#### Go example

The function changes the supplied slice and returns a shorter view of the same
underlying storage:

```go
func compactNonZero(values []int) []int {
	// values[:write] is the compacted prefix. It starts empty because no input
	// elements have been inspected yet.
	write := 0
	for read := range values {
		value := values[read]
		if value == 0 {
			// Zeros do not belong in the result, so only the read position moves.
			continue
		}
		// Copy the next retained value directly after the compacted prefix, then
		// grow that prefix by one. Reading never falls behind writing.
		values[write] = value
		write++
	}
	// Return a slice view of only the compacted prefix; no new array is made.
	return values[:write]
}
```

#### Work and memory

The loop reads each of the `n` input values once in every case, so its work is
linear: `O(n)`. It keeps one index and one current value, so it uses `O(1)`
extra memory. The returned slice is a new slice header over the existing array;
the function does not allocate an output array.

#### When to use it

Use an invariant to reason about a loop, recursion, or data-structure update
when the code builds an answer gradually, mutates state, or discards possible
choices. It is especially useful when a problem asks you to prove that order,
membership, a boundary, or a running total remains valid.

#### When not to use it

Do not force a loop-style invariant onto a direct calculation whose result
follows immediately from a defined operation. For probabilistic or approximate
algorithms, an absolute fact may also be the wrong claim; their argument may
need a probability or error bound instead.

#### Common mistakes and edge cases

> **Warning:** "The processed part is correct" does not say what *correct*
> means. For `[0, 5]`, state that `values[:write]` equals the non-zero values
> read so far, in order.

> **Warning:** Check the invariant after both branches. If the zero branch
> increments `write`, input `[0]` would return a one-value prefix even though
> the expected result is empty.

Empty input is safe: the loop performs no steps, the initial invariant already
describes the whole input, and `values[:0]` is empty. A proof must also name a
progress measure; an invariant alone does not show that a loop terminates.

#### Related material

- [A.1 Complexity and resource accounting](#a1-complexity-and-resource-accounting)
  explains how to account for the work after correctness is established.
- [B.1.1 Moving values in place](#b11-moving-values-in-place) owns the slice
  operation demonstrated here.
- [Remove Element](problems/solved_easy_leetcode_0027_remove_element.go)
  applies the same prefix invariant to a judge problem.

### A.3 Recursion and saved work

#### The problem it solves

Some problems contain smaller instances of the same problem: a directory
contains directories, a tree contains smaller trees, or a calculation for `n`
depends on `n-1`. We need a way to pause the outer work while solving the
smaller part and then resume with its result.

#### The idea in simple terms

Think of opening a smaller box inside a box. The outer box stays open with a
note saying what to do after the inner box closes. A function using the same
function on a smaller input is called **recursion**. Each paused call is a
**stack frame**, which remembers arguments, local values, and the return point.
The number of frames alive at once is the **recursion depth**.

#### How it works

1. Define a **base case**, an input whose answer is returned directly.
2. Change every other input into a strictly smaller or simpler input.
3. Call the function for that smaller input.
4. Combine the returned answer with the current call's work.

The input must move toward the base case. For factorial, require
`0 <= n <= 20` when using `int64`; `20!` fits and `21!` does not.

#### Step-by-step examples

For `factorial(4)`, the expected output is `24`.

1. **Step: Call `4`**
   - **Before:** `n=4`
   - **Decision:** Not a base case; save `4 *` and solve `3`.
   - **After:** Frames wait for `4`, then call `3`.
   - **Known correct:** The answer is `4 * factorial(3)`.
2. **Step: Call `3`**
   - **Before:** `n=3`
   - **Decision:** Save `3 *` and solve `2`.
   - **After:** Frames wait for `4,3`, then call `2`.
   - **Known correct:** The remaining input is smaller.
3. **Step: Call `2`**
   - **Before:** `n=2`
   - **Decision:** Save `2 *` and solve `1`.
   - **After:** Frames wait for `4,3,2`, then call `1`.
   - **Known correct:** The remaining input is smaller.
4. **Step: Call `1`**
   - **Before:** `n=1`
   - **Decision:** Base case; return `1`.
   - **After:** The `2` frame resumes.
   - **Known correct:** `factorial(1)=1`.
5. **Step: Return**
   - **Before:** Saved multipliers are `2,3,4`
   - **Decision:** Apply them in reverse frame order.
   - **After:** Results are `2`, `6`, then `24`.
   - **Known correct:** Each resumed call returns its factorial.

For boundary input `0`, the function returns `1` immediately because there is
one way to arrange no factors in a product. No recursive call is created.

#### Why it works

For `n <= 1`, the direct answer is correct. For `n > 1`, suppose the smaller
call correctly returns `(n-1)!`; multiplying it by `n` gives exactly `n!`.
This argument applies repeatedly until the base case. Because `n` decreases by
one and cannot decrease forever while remaining positive, the process
finishes.

Every unfinished call must remain available until its smaller call returns,
which explains both the reverse return order and the stack memory.

#### Go example

```go
func factorial(n int64) int64 {
	// Zero and one terminate the recursion and both have factorial value one.
	if n <= 1 {
		return 1
	}
	// Leave this call waiting with multiplier n while the smaller subproblem
	// computes (n-1)!, then combine the two results as n * (n-1)!.
	return n * factorial(n-1)
}
```

The stated contract requires `0 <= n <= 20`; validation belongs at the caller
or can be added as a separate checked boundary.

#### Work and memory

The function creates one call for each value from `n` down to `1`, so it does
`O(n)` work in every valid case except the constant-time base cases. At most
`n` calls are unfinished together, so the call stack uses `O(n)` auxiliary
space. Go does not promise to replace this shape with a constant-space loop.

#### When to use it

Use recursion when the input is naturally nested, when each subproblem has the
same contract as the whole problem, and when the maximum depth is safely
bounded. Trees, divide-and-conquer ranges, and choice exploration often have
this shape.

#### When not to use it

Do not use unbounded recursion when a long chain can exhaust the call stack or
when a simple loop expresses the same state more directly. Use an explicit
stack when pending work must be inspected, limited, or processed in a chosen
order.

#### Common mistakes and edge cases

> **Warning:** A broad base case can hide invalid input. Calling the example
> with `-1` returns `1` even though negative factorial is outside its contract;
> enforce the nonnegative precondition at the boundary.

> **Warning:** Count paused calls as memory. `factorial(20)` creates about 20
> live frames at its deepest point even though it allocates no slice.

Repeated recursive calls can also redo the same subproblem. Save repeated
answers as described by [dynamic programming](#c1310-dynamic-programming), or
replace a one-branch recurrence with a loop.

#### Related material

- [A.1 Complexity and resource accounting](#a1-complexity-and-resource-accounting)
  defines auxiliary-space accounting.
- [C.1.3.11 Divide and conquer](#c1311-divide-and-conquer) uses several
  smaller independent calls.
- [D.8 Replacing recursion with a loop](#d8-replacing-recursion-with-a-loop)
  shows how to store pending work explicitly.

### A.4 Integer limits and bits

#### The problem it solves

An arithmetic expression can be mathematically valid but impossible to store
in its Go integer type. We need to reject an operation before it crosses the
allowed range, because checking a value after overflow is too late.

#### The idea in simple terms

An integer type is a fixed-size odometer. A signed 8-bit odometer represents
only `-128` through `127`; the mathematical value `128` has no valid display.
Trying to produce a value outside the type's range is **overflow**. The binary
digits stored in that fixed width are called **bits**.

#### How it works

1. Identify the exact type and its smallest and largest permitted values.
2. Rewrite the safe-result condition so it can be tested before the risky
   multiplication, addition, subtraction, or negation.
3. Reject the input if the condition fails; otherwise perform the operation.
4. Use an unsigned type only when the problem truly excludes negative values
   or treats the bits as a fixed-width pattern.

For nonnegative decimal append `next = value*10 + digit`, where
`0 <= digit <= 9` and `digit <= limit`, the operation is safe exactly when
`value <= (limit-digit)/10`.

#### Step-by-step examples

Use a teaching limit of `127` and append digit `8`.

1. **Input: `value=11`**
   - **Before:** Candidate is `11*10+8`
   - **Decision:** `11 <= (127-8)/10`, or `11 <= 11`.
   - **After:** Produce `118`.
   - **Known correct:** The result is at most `127`.
2. **Input: `value=12`**
   - **Before:** Candidate would be `12*10+8`
   - **Decision:** `12 > 11`, so multiplying would be unsafe.
   - **After:** Return failure without multiplying.
   - **Known correct:** No unrepresentable result was evaluated.

At the boundary `value=12`, digit `7` is accepted because
`12 <= (127-7)/10` and the result is exactly `127`.

#### Why it works

Starting from the required inequality `value*10 + digit <= limit`, subtract
the nonnegative digit and divide by positive `10`. Integer division gives the
largest whole `value` that can be multiplied safely, so the equivalent check
is `value <= (limit-digit)/10`. Performing the comparison first avoids the very
overflow it is meant to detect.

#### Go example

```go
func appendDecimalDigit(value, digit, limit int64) (int64, bool) {
	// Reject values outside the nonnegative decimal-append contract before
	// using them in arithmetic.
	if value < 0 || digit < 0 || digit > 9 || limit < digit {
		return 0, false
	}
	// This rearranges value*10+digit <= limit so the potentially overflowing
	// multiplication is never evaluated when it would be unsafe.
	if value > (limit-digit)/10 {
		return 0, false
	}
	// The preceding inequality proves both the multiplication and addition fit.
	return value*10 + digit, true
}
```

#### Work and memory

The function performs a fixed number of comparisons and arithmetic operations,
so its work is `O(1)`. It stores no collection or recursive state, so auxiliary
space is `O(1)`.

#### When to use it

Use a precondition check whenever input constraints can approach an integer
limit, especially while parsing digits, accumulating products, calculating
midpoints, or converting between integer widths.

#### When not to use it

Do not merely switch to a wider type when the mathematical result is allowed
to grow without a fixed bound; use an arbitrary-precision representation such
as `math/big`. Do not apply this nonnegative formula unchanged to signed
subtraction or negative digit accumulation.

#### Common mistakes and edge cases

> **Warning:** Calculating `next := value*10 + digit` and then checking
> `next > limit` can already have overflowed. For `(12,8,127)`, the check must
> happen before constructing `128` in the bounded type.

> **Warning:** Negating the smallest signed value can overflow because its
> positive magnitude may be one larger than the type's maximum. Handle that
> value explicitly instead of applying `-value` blindly.

The width of `int` depends on the target architecture. Use `int32`, `int64`,
`uint32`, or `uint64` when the contract requires an exact width.

#### Related material

- [A.5 Bit manipulation](#a5-bit-manipulation) explains operations on a
  fixed-width bit pattern.
- [A.6.7 Numbers too large for an integer type](#a67-numbers-too-large-for-an-integer-type)
  processes large decimal input without converting it wholesale.
- [A.8 Integer Digit Processing](#integer-digit-processing) applies the
  precheck while reversing digits.
- [E.6 Overflow-safe integer reversal](#e6-overflow-safe-integer-reversal)
  owns the complete `int64` reversal implementation.

### A.5 Bit manipulation

#### The problem it solves

A program may need to store many independent yes/no choices compactly, inspect
a fixed-width binary value, or use a cancellation identity promised by the
input. We need operations that change or test one binary position without
disturbing the others.

#### The idea in simple terms

Picture a row of 64 light switches. Switch `k` records one choice. An integer
used this way is a **bitmask**. The operators AND (`&`), OR (`|`), exclusive OR
(`^`), and Go's bit clear (`&^`) combine switches according to simple yes/no
rules.

#### How it works

For `0 <= k < 64`, `uint64(1) << k` creates a mask with only switch `k` on.

- `x & mask != 0` tests bit `k`.
- `x | mask` sets bit `k`.
- `x &^ mask` clears bit `k` in Go.
- `x ^ mask` toggles bit `k`.
- `x & (x-1)` clears the lowest set bit.
- `a ^ a` is zero and `a ^ 0` is `a`, so equal pairs cancel under XOR.

Use an unsigned fixed-width type when right shifts must add zeroes on the left
and when the number of positions is part of the contract.

#### Step-by-step examples

Let bits `0`, `1`, and `2` mean read, write, and execute permission. Start with
read and execute enabled: binary `101`, decimal `5`.

1. **Step: Test write**
   - **Before:** `flags=101`, mask `010`
   - **Decision:** `101 & 010` is zero.
   - **After:** Flags stay `101`.
   - **Known correct:** Write is off.
2. **Step: Set write**
   - **Before:** `flags=101`
   - **Decision:** OR with `010` turns that bit on.
   - **After:** `flags=111`
   - **Known correct:** All three permissions are on.
3. **Step: Clear read**
   - **Before:** `flags=111`, mask `001`
   - **Decision:** `&^ 001` clears only bit `0`.
   - **After:** `flags=110`
   - **Known correct:** Write and execute remain on; read is off.

For the cancellation input `[4,7,4]`, the running XOR values are `4`, `3`, and
`7`; the two `4` values cancel. The same method fails for `[4,4,7,7]` if the
task expects one unique value, because every value cancels and produces `0`.

#### Why it works

The one-bit mask is zero at every position except `k`, so each operator's truth
table can affect only that position. OR with one makes it one; bit clear with
one makes it zero; AND reveals whether both operands contain one there.

XOR is associative and commutative, so input order does not matter. Pairing
equal values yields zero, and zero leaves the unpaired value unchanged. That
argument is valid only when the occurrence promise supplies exactly those
pairs.

#### Go example

```go
func hasBit(flags uint64, k uint) bool {
	// A uint64 has positions 0 through 63. Shift a single 1 into position k,
	// AND it with flags, and test whether that position survives.
	return k < 64 && flags&(uint64(1)<<k) != 0
}

func setBit(flags uint64, k uint) (uint64, bool) {
	// Refuse an unrepresentable position instead of pretending it exists.
	if k >= 64 {
		return flags, false
	}
	// OR turns the selected bit on and leaves every other bit unchanged.
	return flags | (uint64(1) << k), true
}

func clearBit(flags uint64, k uint) (uint64, bool) {
	if k >= 64 {
		return flags, false
	}
	// Go's bit-clear operator &^ turns off mask bits that are 1, so only bit k
	// is cleared.
	return flags &^ (uint64(1) << k), true
}
```

#### Work and memory

Each operation on one `uint64` performs bounded machine-word work and uses no
growing storage, so it takes `O(1)` time and `O(1)` auxiliary space. Repeating
`x &= x-1` takes one step per set bit. Enumerating all `n`-bit masks still
visits `2^n` subsets, and copying each chosen subset adds output work.

#### When to use it

Use bit operations when the problem defines a fixed-width bit pattern, when a
small bounded set of flags needs compact membership, or when an identity such
as pair cancellation exactly matches the input promise.

#### When not to use it

Do not use a mask when the number of members can exceed the available width,
when readable set operations matter more than compact storage, or merely
because the inputs happen to be integers. A map or boolean slice is often
clearer for general membership.

#### Common mistakes and edge cases

> **Warning:** Validate the bit position. Shifting a `uint64` mask by `64` or
> more cannot represent a bit in the value; the example returns failure.

> **Warning:** `^x` flips every bit, while `x ^ y` computes XOR between two
> values. They are different uses of the same token in Go.

Right-shifting a signed negative integer preserves sign bits; convert only when
the problem's width and signedness make that interpretation valid. A subset
mask is compact storage, not a shortcut around exponential enumeration.

#### Related material

- [A.4 Integer limits and bits](#a4-integer-limits-and-bits) owns integer width
  and overflow rules.
- [A.7 Combinatorics and output size](#a7-combinatorics-and-output-size)
  accounts for all `2^n` subset outputs.
- [Single Number](problems/easy_leetcode_0136_single_number.go) applies XOR
  cancellation under an exact occurrence promise.
- [Reverse Bits](problems/easy_leetcode_0190_reverse_bits.go) treats the
  input as a fixed-width binary value.

### A.6 Keeping arithmetic values exact and consistent

#### The problem it solves

Values that are mathematically equal can look different in a program, and a
rounded approximation can turn equality into a guess. We need one exact
representation so comparisons, arithmetic, and map keys agree.

#### The idea in simple terms

Choose one kind of label before filing values. Store `$1.20` as `120` cents,
not sometimes as dollars and sometimes as cents. Rewrite equivalent labels
such as `2/4` and `1/2` into one agreed form. That rewrite is called
**normalization**, and its unique result is a **canonical form**.

#### How it works

1. Decide what values count as equal in the problem.
2. Choose a representation that preserves exactly the distinctions the
   problem needs: scaled integers, reduced pairs, remainders, or digit text.
3. Normalize every value at the boundary, before using it as a key or comparing
   it.
4. Check that every intermediate operation fits its integer type.

The child entries below own the rules for divisors, fractions, slopes,
remainders, factor counts, and oversized decimal input.

#### Step-by-step examples

Add prices `$1.20`, `$2.35`, and `$0.45`. The expected exact total is `$4.00`.

1. **Price: `$1.20`**
   - **Before:** `total=0` cents
   - **Decision:** Convert at the boundary to `120` cents.
   - **After:** `total=120`
   - **Known correct:** The first price is exact.
2. **Price: `$2.35`**
   - **Before:** `total=120`
   - **Decision:** Add integer `235`, not a rounded binary decimal.
   - **After:** `total=355`
   - **Known correct:** The first two prices total `$3.55`.
3. **Price: `$0.45`**
   - **Before:** `total=355`
   - **Decision:** Add integer `45`.
   - **After:** `total=400`
   - **Known correct:** All prices total exactly `$4.00`.

The boundary value `$0.00` becomes zero cents and changes nothing. A price
with more precision than the chosen unit, such as `$0.005`, must be rejected or
rounded by an explicit business rule; silently dropping the half-cent would
break the representation contract.

#### Why it works

Every accepted amount is an integer count of the same unit. Integer addition
therefore combines exact counts, and equal amounts have equal stored values.
Normalizing once at the boundary means later code cannot accidentally compare
two different encodings of the same accepted amount.

#### Go example

```go
// Cents gives every stored price the same exact integer unit.
type Cents int64

func addPrices(prices []Cents) Cents {
	// The zero value is the additive identity, so it is the correct empty sum.
	var total Cents
	// Add each already-normalized amount; no floating-point rounding occurs.
	for _, price := range prices {
		total += price
	}
	return total
}
```

The caller must validate precision and check that the sum fits `int64`.

#### Work and memory

Adding `n` already-normalized prices reads each once, so it takes `O(n)` time
in every case and `O(1)` auxiliary space. Parsing, reduction, or output
formatting must be counted separately when the boundary representation is not
already available.

#### When to use it

Use an exact representation when values become equality keys, when rounding
could change a branch, or when the problem defines a precise unit, ratio,
direction, or remainder.

#### When not to use it

Do not invent a scaled integer when the input has unbounded precision or when
the required rounding rule is unknown. Use an arbitrary-precision decimal or
rational type, or clarify the contract first.

#### Common mistakes and edge cases

> **Warning:** Mixing units produces plausible but wrong totals. Adding `120`
> cents directly to `2.35` dollars does not represent `$3.55`; normalize both
> operands before arithmetic.

> **Warning:** Canonical form is part of equality. If `1/2` is reduced but
> `2/4` is not, they remain different Go struct keys despite representing the
> same quantity.

Zero, negative values, invalid denominators, and overflow need explicit rules
for the chosen representation.

#### Related material

- [A.4 Integer limits and bits](#a4-integer-limits-and-bits) explains how to
  guard fixed-width intermediate calculations.
- [A.6.1 Shared divisors and multiples](#a61-shared-divisors-and-multiples-gcd-and-lcm)
  supplies the reduction operation used by ratio-based child entries.
- [A.6.7 Numbers too large for an integer type](#a67-numbers-too-large-for-an-integer-type)
  keeps oversized decimal input as text.

#### A.6.1 Shared divisors and multiples (GCD and LCM)

##### The problem it solves

We need the largest equal group size that divides two counts, or the first
positive cycle length shared by both counts, without listing every divisor or
multiple.

##### The idea in simple terms

If `18` and `12` are split into equally sized groups, the largest possible
group has `6` items. That value is the **greatest common divisor (GCD)**. The
first positive number reached by both counting in steps of `18` and `12` is
`36`; it is the **least common multiple (LCM)**.

Repeatedly replacing the larger question with the leftover part is **Euclid's
algorithm** for the GCD.

##### How it works

For nonnegative `a` and `b`:

1. While `b` is not zero, replace `(a,b)` with `(b,a%b)`.
2. When `b` is zero, `a` is the GCD.
3. If either original value is zero, define the LCM as zero.
4. Otherwise calculate `a/gcd(a,b)*b`, dividing before multiplying and checking
   that the remaining product fits.

##### Step-by-step examples

For `a=18` and `b=12`, the expected GCD is `6` and LCM is `36`.

1. **Step: 1**
   - **Before:** `(a,b)=(18,12)`
   - **Decision:** `18%12=6`; any common divisor of `18` and `12` also divides `6`.
   - **After:** `(12,6)`
   - **Known correct:** The set of common divisors is unchanged.
2. **Step: 2**
   - **Before:** `(12,6)`
   - **Decision:** `12%6=0`.
   - **After:** `(6,0)`
   - **Known correct:** `6` divides both original values.
3. **Step: Stop**
   - **Before:** `(6,0)`
   - **Decision:** No remainder remains.
   - **After:** GCD `6`
   - **Known correct:** No larger common divisor was lost.
4. **Step: LCM**
   - **Before:** GCD `6`
   - **Decision:** Compute `18/6*12`.
   - **After:** `36`
   - **Known correct:** `36` is divisible by both values.

For boundary input `(0,12)`, the GCD loop returns `12`, while the LCM rule
returns `0`. For `(0,0)`, this guide uses the programming convention GCD `0`.

##### Why it works

Write `a = q*b + r`, where `r = a%b`. A number divides both `a` and `b`
exactly when it divides both `b` and `r`, so replacing `(a,b)` with `(b,r)`
preserves all common divisors. Each nonzero remainder is smaller than `b`, so
the process finishes. At `(g,0)`, every positive divisor of both values divides
`g`, and `g` itself divides both, making it greatest.

For positive inputs, `a*b` contains the shared factors twice. Dividing one
operand by the GCD removes one shared copy, producing the smallest value that
contains the factors needed by both inputs.

##### Go example

```go
import "math"

func gcd(a, b int64) int64 {
	// Replacing (a,b) with (b,a%b) preserves their common divisors while making
	// the second value smaller, so the process must eventually reach zero.
	for b != 0 {
		a, b = b, a%b
	}
	// When b is zero, a is the greatest common divisor.
	return a
}

func lcm(a, b int64) (int64, bool) {
	// Under this convention, any least common multiple involving zero is zero.
	if a == 0 || b == 0 {
		return 0, true
	}
	// Divide out the shared factors before multiplying. This reduces the size
	// of the multiplication without changing the mathematical LCM.
	left := a / gcd(a, b)
	// Prove left*b fits in int64 before evaluating the product.
	if left > math.MaxInt64/b {
		return 0, false
	}
	return left * b, true
}
```

The functions require nonnegative inputs.

##### Work and memory

Euclid's algorithm takes `O(log(min(a,b)))` remainder steps in the worst case
and `O(1)` auxiliary space. LCM adds constant work and no growing storage after
the GCD. These bounds treat fixed-width integer operations as constant time.

##### When to use it

Use GCD to reduce ratios, align repeating periods, group counts evenly, or
normalize integer directions. Use LCM when independent positive cycles must
meet again.

##### When not to use it

Do not use LCM as a product shortcut without an overflow check. For a list of
all factors or all primes, use factor enumeration or a sieve instead.

##### Common mistakes and edge cases

> **Warning:** `a*b/gcd(a,b)` can overflow during `a*b` even when the final
> value fits. Divide first, as in `a/gcd(a,b)*b`.

> **Warning:** The example code requires nonnegative values. Passing `-18`
> changes remainder signs and violates its result contract; normalize signs at
> the boundary, taking care with the smallest signed integer.

Zero inputs need declared conventions, and the checked multiplication can
still reject an LCM too large for `int64`.

##### Related material

- [A.6 Keeping arithmetic values exact and consistent](#a6-keeping-arithmetic-values-exact-and-consistent)
  owns the normalization rule.
- [A.6.2 Normalized fractions and ratios](#a62-normalized-fractions-and-ratios)
  uses GCD to create unique ratio keys.
- [A.6.3 Exact slopes without rounding](#a63-exact-slopes-without-rounding)
  uses the same reduction for directions.

#### A.6.2 Normalized fractions and ratios

##### The problem it solves

Pairs such as `(2,4)` and `(3,6)` are different Go map keys even though both
represent one half. We need equal fractions and ratios to produce the same
integer pair without floating-point rounding.

##### The idea in simple terms

Cancel shared pieces and keep the minus sign in one agreed place. The result is
a **normalized fraction**: its numerator and denominator share no positive
factor greater than one, and its denominator is positive.

##### How it works

For `(numerator,denominator)`, require a nonzero denominator and inputs whose
magnitudes can be represented by `int64`.

1. If the numerator is zero, return `(0,1)`.
2. If the denominator is negative, reverse both signs.
3. Find the GCD of the two magnitudes.
4. Divide both values by that GCD.

##### Step-by-step examples

Normalize `(6,-8)`; the expected result is `(-3,4)`.

1. **Step: Sign**
   - **Before:** `(6,-8)`
   - **Decision:** Denominator must be positive, so reverse both signs.
   - **After:** `(-6,8)`
   - **Known correct:** The value is unchanged and the sign has one location.
2. **Step: GCD**
   - **Before:** Magnitudes `(6,8)`
   - **Decision:** Euclid's algorithm returns `2`.
   - **After:** Divisor `2`
   - **Known correct:** `2` is the largest shared factor.
3. **Step: Reduce**
   - **Before:** `(-6,8)`
   - **Decision:** Divide both parts by `2`.
   - **After:** `(-3,4)`
   - **Known correct:** The pair has no remaining shared factor.

`(0,-8)` is the boundary form `(0,1)`: every nonzero denominator represents
the same zero. `(1,0)` is invalid and returns failure.

##### Why it works

Reversing both signs multiplies numerator and denominator by `-1`, so the
ratio is unchanged. Dividing both by the same nonzero GCD also preserves the
ratio. After division the parts are relatively prime, meaning they have no
shared factor greater than one. A positive denominator fixes the only remaining
sign choice, so equal accepted fractions reach the same pair.

##### Go example

```go
func normalizeFraction(numerator, denominator int64) (int64, int64, bool) {
	// Division by zero is not a fraction, while every zero numerator has the
	// single canonical representation 0/1.
	if denominator == 0 {
		return 0, 0, false
	}
	if numerator == 0 {
		return 0, 1, true
	}
	if denominator < 0 {
		// Keep the denominator positive by moving any sign to the numerator.
		numerator = -numerator
		denominator = -denominator
	}

	// Euclid's algorithm needs magnitudes. Only the numerator can still be
	// negative because the denominator was normalized above.
	a, b := numerator, denominator
	if a < 0 {
		// GCD operates on magnitude; the original signed numerator remains intact
		// for the final normalized result.
		a = -a
	}
	// Each Euclidean step preserves the shared divisors and reduces the second
	// operand until a contains the GCD.
	for b != 0 {
		a, b = b, a%b
	}
	// Divide both parts by their GCD to obtain the unique reduced pair.
	return numerator / a, denominator / a, true
}
```

The function requires neither input to equal `math.MinInt64`, whose magnitude
does not fit in `int64`.

##### Work and memory

Sign handling is constant work. GCD takes `O(log(min(|numerator|,
|denominator|)))` remainder steps in the worst case. The function uses `O(1)`
auxiliary space and returns two integers without allocating a collection.

##### When to use it

Use a normalized pair when exact ratios become map keys, equality values, or
directions and the integer parts remain within a known bound.

##### When not to use it

Do not use it for a zero denominator, for approximate measured data whose
tolerance is part of equality, or when intermediate subtraction has already
overflowed before the pair reaches this function.

##### Common mistakes and edge cases

> **Warning:** Reducing magnitude without fixing the sign leaves two keys for
> one value. `(3,-4)` and `(-3,4)` must both normalize to `(-3,4)`.

> **Warning:** Zero needs one canonical key. Leaving `(0,8)` and `(0,3)` as-is
> makes equal zero ratios compare unequal.

Validate `denominator != 0` before remainder or division, and handle the
smallest signed integer under an explicit wider or unsigned-magnitude policy.

##### Related material

- [A.6.1 Shared divisors and multiples](#a61-shared-divisors-and-multiples-gcd-and-lcm)
  proves the reduction step.
- [A.6.3 Exact slopes without rounding](#a63-exact-slopes-without-rounding)
  adds direction-specific duplicate and axis cases.
- [D.4 Map presence and sets](#d4-map-presence-and-sets) explains how normalized
  comparable pairs behave as Go map keys.

#### A.6.3 Exact slopes without rounding

##### The problem it solves

Several points can lie on the same line through a starting point even though
their horizontal and vertical changes have different sizes or opposite signs.
We need one exact map key per undirected line orientation without dividing
into a rounded decimal slope.

##### The idea in simple terms

Treat movement as an integer arrow `(changeX,changeY)`. Shorten the arrow by
its greatest shared factor and choose one of the two opposite signs as the
agreed representation. The shortened pair is an exact **direction key** for
an undirected line through the starting point.

##### How it works

1. Subtract coordinates in a type wide enough for the differences.
2. If both changes are zero, report a duplicate point; it has no direction.
3. Map every vertical direction to `(0,1)` and every horizontal direction to
   `(1,0)`.
4. Otherwise divide both changes by their GCD.
5. If the reduced horizontal change is negative, reverse both signs.

The example code requires differences that fit `int64` and are not
`math.MinInt64`.

##### Step-by-step examples

Use start `(1,1)` and points `(2,3)`, `(3,5)`, and `(0,-1)`. Every expected
direction key is `(1,2)`.

1. **Point: `(2,3)`**
   - **Before:** Change `(1,2)`
   - **Decision:** GCD is `1`; sign is already canonical.
   - **After:** Key `(1,2)`
   - **Known correct:** The key represents this direction.
2. **Point: `(3,5)`**
   - **Before:** Change `(2,4)`
   - **Decision:** Divide both parts by GCD `2`.
   - **After:** Key `(1,2)`
   - **Known correct:** The longer arrow shares the same key.
3. **Point: `(0,-1)`**
   - **Before:** Change `(-1,-2)`
   - **Decision:** GCD is `1`, then reverse both signs because `changeX < 0`.
   - **After:** Key `(1,2)`
   - **Known correct:** The opposite arrow has the same undirected line orientation through the start.

For point `(1,1)`, both changes are zero, so the result is marked duplicate
instead of inventing a slope. For `(1,5)`, the fixed vertical key is `(0,1)`.

##### Why it works

Points on the same line through the start have movement pairs that are integer
multiples of the same reduced pair. Dividing by the GCD removes that scale.
The axis rules remove arbitrary magnitude, and the sign rule chooses one of the
two equivalent orientations for an undirected line. Duplicate points are kept
separate because `(0,0)` is a multiple of every direction and therefore cannot
identify one.

##### Go example

```go
// direction is comparable, so a normalized value can be used as a map key.
type direction struct {
	dx int64
	dy int64
}

func normalizeDirection(dx, dy int64) (direction, bool) {
	// Two identical points have no direction. The boolean reports this duplicate
	// separately from valid vertical and horizontal directions.
	if dx == 0 && dy == 0 {
		return direction{}, true
	}
	if dx == 0 {
		// Upward and downward vectors share the canonical vertical key (0,1)
		// because this function represents an undirected line orientation.
		return direction{dx: 0, dy: 1}, false
	}
	if dy == 0 {
		// Leftward and rightward vectors likewise share the canonical horizontal
		// key (1,0).
		return direction{dx: 1, dy: 0}, false
	}

	// Compute the GCD of the absolute components so equivalent scaled vectors
	// reduce to the same pair.
	a, b := dx, dy
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	// Euclid's algorithm leaves the positive GCD in a.
	for b != 0 {
		a, b = b, a%b
	}
	dx /= a
	dy /= a
	// Choose one orientation convention: dx is always positive for nonvertical
	// directions, so opposite sign encodings cannot create duplicate keys.
	if dx < 0 {
		dx = -dx
		dy = -dy
	}
	return direction{dx: dx, dy: dy}, false
}
```

The boolean result reports a duplicate point.

##### Work and memory

Coordinate subtraction is constant work. GCD takes logarithmic remainder
steps in the smaller change magnitude, so normalization is
`O(log(min(|dx|,|dy|)))` time in the worst case and `O(1)` auxiliary space.
A caller grouping `n` points also pays for its map storage and lookups.

##### When to use it

Use reduced direction keys when points must be grouped by exact collinearity
on undirected lines through one chosen start and coordinates are integers
within safe bounds.

##### When not to use it

Do not use a rounded floating-point slope as an equality key. Do not use this
undirected sign rule when opposite rays must remain different; keep their signs
distinct for that contract.

##### Common mistakes and edge cases

> **Warning:** Dividing `changeY/changeX` loses exactness and divides by zero
> on vertical lines. Normalize the integer pair instead.

> **Warning:** Treating duplicate `(0,0)` as a direction lets it match every
> slope incorrectly. Count duplicates separately.

Calculate coordinate differences before narrowing their type. Horizontal,
vertical, negative, and duplicate cases all need explicit tests.

##### Related material

- [A.6.1 Shared divisors and multiples](#a61-shared-divisors-and-multiples-gcd-and-lcm)
  owns Euclid's algorithm.
- [A.6.2 Normalized fractions and ratios](#a62-normalized-fractions-and-ratios)
  explains canonical reduced pairs.
- [Max Points on a Line](problems/hard_leetcode_0149_max_points_on_a_line.go)
  uses direction keys with duplicate counting.

#### A.6.4 Modular arithmetic and negative division

##### The problem it solves

Go's results for negative remainder and integer division may differ from the
nonnegative bucket or downward-rounded result required by a mathematical
problem. We need explicit conversion rules instead of assuming the operators
already use the desired convention.

##### The idea in simple terms

Wrapping `-1` around a five-position clock lands at position `4`, even though
Go evaluates `-1 % 5` as `-1`. Moving a negative Go remainder back into the
clock range is **normalized modular arithmetic**. Rounding a quotient down
toward negative infinity is **floor division**; Go instead truncates toward
zero.

##### How it works

For positive modulus or divisor `m`:

1. Calculate Go's quotient `x/m` and remainder `x%m`.
2. For a remainder in `0..m-1`, add `m` once when the remainder is negative.
3. For floor division, subtract one from the truncated quotient when the
   remainder is negative.

The relation `x = quotient*m + remainder` connects both corrections.

##### Step-by-step examples

For `x=-7` and `m=3`, the normalized remainder should be `2` and the floor
quotient should be `-3`.

1. **Result: Go operators**
   - **Before:** `x=-7`, `m=3`
   - **Decision:** Go truncates toward zero.
   - **After:** `quotient=-2`, `remainder=-1`
   - **Known correct:** `-7 = -2*3 + -1`.
2. **Result: Normalize remainder**
   - **Before:** `remainder=-1`
   - **Decision:** Add positive modulus `3`.
   - **After:** `2`
   - **Known correct:** The result is in `0..2` and differs from `-7` by a multiple of `3`.
3. **Result: Floor quotient**
   - **Before:** `quotient=-2`, remainder negative
   - **Decision:** Move one whole step downward.
   - **After:** `-3`
   - **Known correct:** `-3 <= -7/3 < -2`.

Boundary input `x=-6,m=3` has remainder zero. Its exact quotient is `-2`, so
neither correction changes the result.

##### Why it works

Go's remainder has magnitude smaller than positive `m`. Therefore a negative
remainder lies in `-(m-1)..-1`, and adding `m` moves it exactly into `1..m-1`
without changing its equivalence class. For division, truncation and floor
differ only when a negative result has a fractional part; a negative remainder
detects exactly that case for positive `m`, so subtracting one reaches the next
lower integer.

##### Go example

```go
func normalizedRemainder(x, modulus int64) (int64, bool) {
	// A positive modulus defines canonical representatives 0 through modulus-1.
	if modulus <= 0 {
		return 0, false
	}
	remainder := x % modulus
	if remainder < 0 {
		// Go keeps the dividend's sign; adding one modulus moves a negative
		// remainder into the canonical nonnegative interval.
		remainder += modulus
	}
	return remainder, true
}

func floorDivide(x, divisor int64) (int64, bool) {
	if divisor <= 0 {
		return 0, false
	}
	// Go truncates toward zero. For a negative nonmultiple, floor is one less
	// than that truncated quotient.
	quotient := x / divisor
	if x%divisor < 0 {
		quotient--
	}
	return quotient, true
}
```

##### Work and memory

Each function performs a fixed number of integer operations, so it takes
`O(1)` time and `O(1)` auxiliary space.

##### When to use it

Use normalized remainders for circular indexes, equivalence classes, and
problem statements that require `0 <= result < modulus`. Use floor division
when intervals or coordinates are defined by rounding downward.

##### When not to use it

Do not add the correction when the required contract is Go's truncation toward
zero. The formulas above also deliberately require a positive divisor; derive
and test a different rule before accepting negative divisors.

##### Common mistakes and edge cases

> **Warning:** In Go, `-1 % 5` is `-1`, so using it directly as a slice index
> is invalid. Normalize it to `4` first.

> **Warning:** Always subtracting one for negative `x` breaks exact divisions.
> `-6/3` is already `-2`; adjust only when the remainder is negative.

Reject zero modulus or divisor before the operator executes. Keep the smallest
signed value divided by `-1` out of generalized variants because that quotient
does not fit the signed type.

##### Related material

- [A.4 Integer limits and bits](#a4-integer-limits-and-bits) explains why
  operator inputs must also fit their type.
- [A.6 Keeping arithmetic values exact and consistent](#a6-keeping-arithmetic-values-exact-and-consistent)
  owns the rule for choosing one representation.
- [A.6.7 Numbers too large for an integer type](#a67-numbers-too-large-for-an-integer-type)
  updates a remainder while reading decimal text.

#### A.6.5 Trailing zeroes without calculating a factorial

##### The problem it solves

We need the number of zero digits at the end of `n!` without constructing the
factorial, whose value becomes enormous long before its factor count is hard to
compute.

##### The idea in simple terms

A trailing decimal zero comes from multiplying `2*5`. The product
`1*2*...*n`, called **`n` factorial** and written `n!`, contains more factors
of `2` than factors of `5`. Count the scarce `5`s instead of building the
product.

##### How it works

For nonnegative `n`:

1. Count multiples of `5`; each contributes at least one factor of `5`.
2. Count multiples of `25`; each contributes a second factor.
3. Continue with `125`, `625`, and so on until the count is zero.

Dividing a running value by `5` each time produces those layer counts without
letting the powers themselves overflow.

##### Step-by-step examples

For `n=25`, the expected result is `6`.

1. **Step: First layer**
   - **Before:** `remaining=25`, `zeroes=0`
   - **Decision:** `25/5=5` counts one `5` from `5,10,15,20,25`.
   - **After:** `remaining=5`, `zeroes=5`
   - **Known correct:** Every number with at least one `5` has contributed once.
2. **Step: Second layer**
   - **Before:** `remaining=5`, `zeroes=5`
   - **Decision:** `5/5=1` counts the second `5` inside `25`.
   - **After:** `remaining=1`, `zeroes=6`
   - **Known correct:** Every factor of `25` has contributed its extra `5`.
3. **Step: Stop**
   - **Before:** `remaining=1`
   - **Decision:** It is smaller than `5`; no higher layer exists.
   - **After:** Return `6`.
   - **Known correct:** All factors of `5` in `25!` are counted.

For boundary inputs `0` through `4`, no factor of `5` occurs, so the loop does
no work and returns `0`.

##### Why it works

Every trailing zero uses one factor of `2` and one factor of `5`. In a
factorial, each multiple of `5` supplies one `5`, each multiple of `25`
supplies one more, and higher powers supply further layers. The sum
`floor(n/5)+floor(n/25)+...` therefore counts every factor of `5` exactly once
per occurrence. Factors of `2` are never fewer and are more frequent once a
factor of `5` appears, so the number of pairs is exactly the number of counted
`5`s.

##### Go example

```go
func trailingZeroes(n int64) int64 {
	var zeroes int64
	// Each division counts multiples of 5, then the extra factors contributed
	// by multiples of 25, 125, and so on.
	for n >= 5 {
		n /= 5
		zeroes += n
	}
	return zeroes
}
```

The function requires `n >= 0`.

##### Work and memory

Each pass divides `n` by `5`, so there are `O(log_5 n)` passes, conventionally
written `O(log n)`. The function stores two integers and uses `O(1)` auxiliary
space. It never allocates or constructs `n!`.

##### When to use it

Use factor-layer counting when a question asks about factors inside a factorial
product, especially trailing zeroes in a known base.

##### When not to use it

Do not use this base-10 shortcut for trailing zeroes in another base without
factoring that base and counting all required prime factors. Do not construct
the factorial merely to inspect its decimal suffix.

##### Common mistakes and edge cases

> **Warning:** Counting only `n/5` misses repeated factors. For `25!`, it gives
> `5` instead of `6` because `25=5*5` contributes twice.

> **Warning:** Multiplying a power variable by `5` can overflow at the loop
> boundary. Repeatedly divide `n`, as the example does.

Reject negative `n`; factorial is not defined by this entry for negative
integers.

##### Related material

- [A.6.6 Factors and prime sieves](#a66-factors-and-prime-sieves) defines prime
  factors and repeated divisibility.
- [Factorial Trailing Zeroes](problems/medium_leetcode_0172_factorial_trailing_zeroes.go)
  applies the factor-layer count.

#### A.6.6 Factors and prime sieves

##### The problem it solves

We need either to test one integer for a nontrivial factor or to identify every
prime up to a shared limit without trying every possible divisor for every
number.

##### The idea in simple terms

Factors come in pairs: if `91=7*13`, one member is no larger than the square
root. For one number, try only possible small partners. For a whole range,
start with every candidate and repeatedly cross out multiples. That range
method is the **Sieve of Eratosthenes**. A **prime** is an integer greater than
one whose only positive divisors are `1` and itself.

##### How it works

For one `n >= 2`, test divisors `2,3,...` while `divisor <= n/divisor`; the
division form avoids overflow from `divisor*divisor`. The first divisor found
is the smallest factor. If none is found, `n` is prime.

For all primes through `limit`:

1. Create a boolean marker for `0..limit`.
2. Starting at `2`, keep an unmarked value as prime.
3. Cross out its multiples beginning at its square; smaller multiples were
   crossed out by earlier primes.
4. Return every unmarked value at least `2`.

##### Step-by-step examples

Trial division for `91` checks `2` through `7`. Values `2..6` leave nonzero
remainders; `91%7=0`, so the smallest factor is `7` and the paired factor is
`13`.

For a sieve through `10`:

1. **Step: Keep `2`**
   - **Before:** Nothing is marked.
   - **Decision:** `2` is unmarked; cross out `4,6,8,10`.
   - **After:** Even composites are marked.
   - **Known correct:** `2` is prime; its composite multiples cannot be prime.
2. **Step: Keep `3`**
   - **Before:** `3` is unmarked.
   - **Decision:** Start at `3*3=9`; mark `9`.
   - **After:** `9` is marked.
   - **Known correct:** `3` is prime; smaller multiples were already handled.
3. **Step: Stop crossing**
   - **Before:** Next candidate square exceeds `10`.
   - **Decision:** Every composite has a factor at most its square root.
   - **After:** Collect unmarked `2,3,5,7`.
   - **Known correct:** These and only these values are prime through `10`.

For `limit < 2`, the result is empty. For `n=2`, trial division performs no
candidate checks and reports `2` itself as prime.

##### Why it works

If composite `n=a*b` had both factors larger than its square root, their
product would exceed `n`. Therefore checking through the square root cannot
miss a composite. In the sieve, every composite has a smallest prime factor
`p`; when `p` is processed, the composite is marked. Starting at `p*p` is safe
because any smaller multiple of `p` has another factor smaller than `p` and
was already marked.

##### Go example

```go
func smallestFactor(n int) int {
	// Values below two have no prime factor under this function's contract.
	if n < 2 {
		return 0
	}
	// If n is composite, at least one factor is no larger than sqrt(n). Writing
	// divisor <= n/divisor avoids overflowing divisor*divisor.
	for divisor := 2; divisor <= n/divisor; divisor++ {
		if n%divisor == 0 {
			// Ascending search makes the first divisor found the smallest one.
			return divisor
		}
	}
	// No smaller divisor exists, so n itself is prime.
	return n
}

func primesThrough(limit int) []int {
	// No prime exists below two, so there is nothing to allocate or scan.
	if limit < 2 {
		return nil
	}
	// composite[value] records that a previously discovered prime divides value.
	composite := make([]bool, limit+1)
	// Only bases through sqrt(limit) can have a first unmarked multiple within
	// the requested range.
	for prime := 2; prime <= limit/prime; prime++ {
		if composite[prime] {
			// Composite bases were already handled by their prime factors.
			continue
		}
		// Smaller multiples were marked by smaller primes, so begin at prime^2.
		for multiple := prime * prime; multiple <= limit; multiple += prime {
			composite[multiple] = true
		}
	}
	// Every unmarked value at least two is prime.
	primes := make([]int, 0)
	for value := 2; value <= limit; value++ {
		if !composite[value] {
			// An unmarked value has no smaller prime divisor and is therefore prime.
			primes = append(primes, value)
		}
	}
	return primes
}
```

##### Work and memory

Trial division uses `O(sqrt(n))` remainder checks in the worst case and `O(1)`
auxiliary space. The sieve takes `O(limit log log limit)` marking work and
`O(limit)` marker memory; collecting primes adds `O(limit)` scan work and an
output slice proportional to the number of primes.

##### When to use it

Use trial division for one or a few moderate integers. Use a sieve when many
queries share the same upper limit or when every prime in a range is needed.

##### When not to use it

Do not allocate a sieve up to an enormous sparse limit for a handful of
queries. Do not repeatedly trial-divide every value in a dense range when one
sieve can share the work.

##### Common mistakes and edge cases

> **Warning:** Values `0` and `1` are not prime. Treating every number with no
> found divisor as prime would classify both incorrectly.

> **Warning:** `divisor*divisor <= n` can overflow for a large fixed-width
> integer. Use `divisor <= n/divisor` when the types approach their limit.

Begin sieve marking at `prime*prime`, but only after the outer condition proves
that square is within the limit.

##### Related material

- [A.6.1 Shared divisors and multiples](#a61-shared-divisors-and-multiples-gcd-and-lcm)
  compares shared-factor queries with factor enumeration.
- [A.6.5 Trailing zeroes](#a65-trailing-zeroes-without-calculating-a-factorial)
  counts one prime factor across an implicit product.

#### A.6.7 Numbers too large for an integer type

##### The problem it solves

A decimal integer can contain thousands of digits and therefore cannot be
converted to `int64`, while the task may need only a bounded property such as
its remainder after division.

##### The idea in simple terms

Read the number the way long division does: carry only the unfinished
remainder into the next digit. The full number stays as text, and the state
never needs to grow with its magnitude.

##### How it works

For a nonempty string containing only the one-byte digit characters `'0'`
through `'9'` from the American Standard Code for Information Interchange
(ASCII), and positive modulus `m`:

1. Start `remainder=0`.
2. Convert the next byte from `'0'..'9'` to its digit value.
3. Update `remainder=(remainder*10+digit)%m`.
4. Reject any nondigit and return the final remainder after all digits.

The example code also requires `m <= 922337203685477579`, which ensures the
intermediate `10*(m-1)+9` fits in `int64`.

##### Step-by-step examples

For digits `"1234"` and modulus `7`, the expected remainder is `2`.

1. **Digit: `'1'`**
   - **Before:** `remainder=0`
   - **Decision:** `(0*10+1)%7=1`
   - **After:** `1`
   - **Known correct:** Remainder of prefix `"1"` is `1`.
2. **Digit: `'2'`**
   - **Before:** `1`
   - **Decision:** `(1*10+2)%7=5`
   - **After:** `5`
   - **Known correct:** Remainder of `"12"` is `5`.
3. **Digit: `'3'`**
   - **Before:** `5`
   - **Decision:** `(5*10+3)%7=4`
   - **After:** `4`
   - **Known correct:** Remainder of `"123"` is `4`.
4. **Digit: `'4'`**
   - **Before:** `4`
   - **Decision:** `(4*10+4)%7=2`
   - **After:** `2`
   - **Known correct:** Remainder of the complete number is `2`.

Leading zeroes are harmless: `"0012"` follows extra zero states and still
returns `5` modulo `7`. Empty text and `"12x"` are invalid.

##### Why it works

Before each digit, `remainder` equals the numeric prefix modulo `m`. Appending
digit `d` changes that prefix from `p` to `10p+d`. Replacing `p` by an equal
remainder does not change the final remainder, so the update preserves the
fact for the longer prefix. After the last digit, the prefix is the complete
input.

##### Go example

```go
const maxSafeDecimalModulus int64 = 922337203685477579

func decimalRemainder(digits string, modulus int64) (int64, bool) {
	// The modulus bound ensures remainder*10+digit fits before the next modulo.
	if digits == "" || modulus <= 0 || modulus > maxSafeDecimalModulus {
		return 0, false
	}
	var remainder int64
	for i := 0; i < len(digits); i++ {
		// Validate and convert one ASCII decimal digit without parsing the whole,
		// potentially enormous, number.
		if digits[i] < '0' || digits[i] > '9' {
			return 0, false
		}
		digit := int64(digits[i] - '0')
		// Appending a digit changes r to (10*r+digit) mod modulus. Keeping only
		// r is valid because congruent prefixes stay congruent after appending.
		remainder = (remainder*10 + digit) % modulus
	}
	return remainder, true
}
```

##### Work and memory

For `d` input digits, the function performs one update per digit, so it takes
`O(d)` time in every valid case. It stores one remainder and indexes the input
without copying it, so auxiliary space is `O(1)`.

##### When to use it

Use streaming remainder updates when the input is an exact decimal string and
only divisibility, a remainder, or another bounded left-to-right state is
needed.

##### When not to use it

Do not use this method when later work needs the complete large value or
arbitrary arithmetic on it; use `math/big.Int`. Do not accept a modulus that
can overflow the update unless the implementation uses a wider or explicitly
overflow-safe modular operation.

##### Common mistakes and edge cases

> **Warning:** Converting the whole text before taking `%m` fails before the
> useful bounded operation runs. A 1,000-digit input cannot fit in `int64`.

> **Warning:** Validate every byte. Treating `'x'-'0'` as a digit silently
> corrupts the state for `"12x"`.

Choose and document rules for a leading sign, whitespace, and empty input. The
example accepts leading zeroes but digits only.

##### Related material

- [A.4 Integer limits and bits](#a4-integer-limits-and-bits) owns the
  intermediate-range check.
- [A.6.4 Modular arithmetic and negative division](#a64-modular-arithmetic-and-negative-division)
  defines normalized remainders for signed values.
- [A.8 Integer Digit Processing](#integer-digit-processing) extracts digits
  from an integer that already fits in memory.

### A.7 Combinatorics and output size

#### The problem it solves

A request for every possible group or ordering can have an enormous answer
even when the input is small. We need to count possible outputs before choosing
an algorithm and to include output construction in its cost.

#### The idea in simple terms

If every input value independently has an "in" or "out" switch, `n` values
create `2^n` switch settings. No implementation can return all those groups
without visiting and storing the requested answers. **Combinatorics** is the
study of counting such choices.

#### How it works

First identify what makes two outputs different:

- A **subset** chooses any number of distinct input positions; order does not
  matter. There are `2^n` subsets.
- A **permutation** arranges all `n` positions; order matters. There are `n!`,
  where `n! = n*(n-1)*...*1`.
- A **size-`k` combination** chooses exactly `k` positions without order. Its
  count is `C(n,k)`, read "`n` choose `k`."

Then count both the number of outputs and their sizes. Generating a result
slice usually copies its elements, so output work can exceed the number of
result rows.

#### Step-by-step examples

Generate subsets of `[A,B,C]`. At each position, first exclude and then include
the current value.

1. **Position and choice: Start**
   - **Before:** `current=[]`, index `0`
   - **Decision:** Explore both choices for `A`.
   - **After:** Exclude branch, then include branch
   - **Known correct:** Every subset belongs to exactly one branch.
2. **Position and choice: Exclude `A`, branch on `B,C`**
   - **Before:** `[]`
   - **Decision:** Repeat the two choices.
   - **After:** `[]`, `[C]`, `[B]`, `[B,C]`
   - **Known correct:** These are exactly the subsets without `A`.
3. **Position and choice: Include `A`, branch on `B,C`**
   - **Before:** `[A]`
   - **Decision:** Repeat the same choices.
   - **After:** `[A]`, `[A,C]`, `[A,B]`, `[A,B,C]`
   - **Known correct:** These are exactly the subsets with `A`.
4. **Position and choice: Finish**
   - **Before:** Eight outputs
   - **Decision:** Both top-level groups are disjoint and complete.
   - **After:** Return all eight.
   - **Known correct:** Every switch setting appears once.

For empty input, there is one subset--the empty subset--so the output is
`[[]]`, not an empty list of answers.

#### Why it works

Before processing position `i`, `current` contains exactly the values selected
from positions before `i`. Every subset either excludes position `i` or
includes it, never both, so the two recursive branches are disjoint and cover
all possibilities. When `i` reaches `n`, the choices for every position are
fixed and one complete subset is copied. Because each level advances `i`, the
process finishes.

#### Go example

```go
func subsets(values []string) [][]string {
	// current holds the choices on the active recursion path; result owns copies
	// of complete paths so later backtracking cannot mutate saved subsets.
	result := make([][]string, 0)
	current := make([]string, 0, len(values))
	var build func(int)
	build = func(index int) {
		if index == len(values) {
			// Every input position has made its include/exclude decision.
			result = append(result, append([]string{}, current...))
			return
		}
		// First enumerate subsets that exclude values[index].
		build(index + 1)
		// Then include it, enumerate that branch, and undo the choice before
		// returning to the caller.
		current = append(current, values[index])
		build(index + 1)
		current = current[:len(current)-1]
	}
	build(0)
	return result
}
```

#### Work and memory

There are `2^n` returned subsets. Copying their elements costs
`n*2^(n-1)` total element copies for `n>0`, so time and required output space
are `O(n*2^n)`. The recursion and `current` slice use `O(n)` auxiliary space,
excluding returned output. A permutation generator similarly cannot beat its
`n!` output count, and storing all size-`k` combinations requires
`O(k*C(n,k))` output space.

#### When to use it

Use combinatorial counting before implementing any "return all choices"
request, and when constraints are meant to tell you whether exhaustive output
is feasible.

#### When not to use it

Do not generate every object when the task asks only for a count, an optimum,
or one valid object. A counting formula, dynamic program, or search that stops
after one result may avoid the full output cost.

#### Common mistakes and edge cases

> **Warning:** Calling subset generation merely `O(2^n)` ignores copying. For
> `[A,B,C]`, eight rows are returned, but their contents hold 12 values in
> total.

> **Warning:** The empty input has one subset and one permutation--the empty
> choice. Returning no answers breaks recursive composition.

Repeated values require a definition of distinct output: positions may be
distinct even when their values compare equal.

#### Related material

- [A.1 Complexity and resource accounting](#a1-complexity-and-resource-accounting)
  defines output and auxiliary-space accounting.
- [A.7.1 Counting combinations without large factorials](#a71-counting-combinations-without-large-factorials)
  computes `C(n,k)` exactly.
- [A.7.2 Repeated input values](#a72-repeated-input-values) distinguishes
  position choices from value choices.
- [C.1.3.19 Backtracking](#c1319-backtracking) owns the general choose-and-undo
  method.

#### A.7.1 Counting combinations without large factorials

##### The problem it solves

We need the exact number of size-`k` groups chosen from `n` distinct positions,
but calculating three full factorials can overflow before their quotient is
small enough to use.

##### The idea in simple terms

Build only the factors that survive cancellation. The count `C(n,k)`, read
"`n` choose `k`," is the same as choosing the `k` kept positions or the `n-k`
omitted positions, so use the smaller side.

##### How it works

For `0 <= k <= n`, replace `k` by `min(k,n-k)`, start `result=1`, and for
`i=1..k` update:

```text
result = result * (n-k+i) / i
```

Each complete step is an integer. When the result can exceed fixed-width
integers, use `big.Int` so both the intermediate product and answer remain
exact.

An alternative for many related counts is **Pascal's recurrence**:
`C(n,k)=C(n-1,k-1)+C(n-1,k)`, with boundary values `C(n,0)=C(n,n)=1`.

##### Step-by-step examples

For `C(5,2)`, the expected result is `10`.

1. **Step: Symmetry**
   - **Before:** `n=5`, `k=2`
   - **Decision:** `min(2,3)=2`; two loop steps are enough.
   - **After:** `result=1`
   - **Known correct:** Counting kept and omitted positions is symmetric.
2. **Step: `i=1`**
   - **Before:** `result=1`
   - **Decision:** Multiply by `4`, divide by `1`.
   - **After:** `result=4`
   - **Known correct:** This is `C(4,1)`.
3. **Step: `i=2`**
   - **Before:** `result=4`
   - **Decision:** Multiply by `5`, divide by `2`.
   - **After:** `result=10`
   - **Known correct:** This is `C(5,2)`.

Pascal's recurrence reaches the same answer by marking `E` in
`[A,B,C,D,E]`: four pairs include `E`, while six exclude it, so
`C(5,2)=C(4,1)+C(4,2)=4+6=10`.

For boundary `k=0` or `k=n`, the loop performs zero iterations after symmetry
and returns `1`. Invalid `k<0` or `k>n` returns failure.

##### Why it works

After iteration `i`, `result` equals `C(n-k+i,i)`. The claim begins at
`C(n-k,0)=1`. The update multiplies by the next available numerator and divides
by the new chosen count, which is exactly the algebraic relationship between
consecutive combination counts. Therefore after `k` steps it equals `C(n,k)`.

Pascal's recurrence is correct because every size-`k` group either includes a
marked value or excludes it. Those cases cannot overlap and cover every group.

##### Go example

```go
import "math/big"

func combinationCount(n, k int64) (*big.Int, bool) {
	// A combination count is defined here only for 0 <= k <= n.
	if n < 0 || k < 0 || k > n {
		return nil, false
	}
	// C(n,k) equals C(n,n-k); using the smaller side shortens the product.
	if k > n-k {
		k = n - k
	}
	// Reuse big.Int temporaries and apply one numerator and denominator factor
	// per step. Every intermediate quotient is exact for this recurrence.
	result := big.NewInt(1)
	factor := new(big.Int)
	divisor := new(big.Int)
	for i := int64(1); i <= k; i++ {
		// After this step, result equals C(n-k+i,i). Multiplying the previous
		// count by n-k+i and dividing by i is exact; at i=k this becomes C(n,k).
		result.Mul(result, factor.SetInt64(n-k+i))
		result.Quo(result, divisor.SetInt64(i))
	}
	return result, true
}
```

##### Work and memory

The loop performs `min(k,n-k)` big-integer multiplications and exact divisions.
That is `O(min(k,n-k))` arithmetic operations; their bit-level cost grows with
the number of bits in `C(n,k)`. The result and temporaries use
`O(log C(n,k))` bits of magnitude storage. A Pascal table for many counts uses
`O(nk)` arithmetic operations and can use `O(k)` big integers when only the
previous row is kept.

##### When to use it

Use the multiplicative formula for one exact combination count. Use Pascal's
recurrence when many nearby counts or a constraint-modulo version are needed.

##### When not to use it

Do not generate all groups merely to count them. Do not use fixed-width
factorials when intermediate values can overflow, even if the final count
would fit.

##### Common mistakes and edge cases

> **Warning:** `n!/(k!*(n-k)!)` can overflow at `n!` before division. For
> `C(30,2)=435`, `30!` is already far outside `int64`.

> **Warning:** In fixed-width code, `result*factor` can overflow before the
> following exact division. Check or use arbitrary precision for the
> intermediate, not only the final answer.

Define invalid inputs and preserve `C(n,0)=C(n,n)=1`.

##### Related material

- [A.4 Integer limits and bits](#a4-integer-limits-and-bits) explains
  pre-operation overflow checks.
- [A.7 Combinatorics and output size](#a7-combinatorics-and-output-size)
  distinguishes counting from generation.
- [C.1.3.10 Dynamic programming](#c1310-dynamic-programming) generalizes
  Pascal's recurrence and rolling rows.

#### A.7.2 Repeated input values

##### The problem it solves

When input values repeat, different position choices can produce identical
value outputs. We need to state which notion of distinctness the result uses
and, when required, generate each value combination once.

##### The idea in simple terms

Label equal positions temporarily: `[1a,1b,2]` has three position pairs, but
`[1a,2]` and `[1b,2]` both display as `[1,2]`. Sort equal values together, then
skip a repeated value only when it would make the same choice at the same
output depth.

##### How it works

For distinct size-`k` value combinations:

1. Copy and sort the input so equal values are adjacent.
2. Build `current` from left to right using increasing indexes.
3. At one recursive depth, skip `sorted[i]` when `i>start` and it equals the
   immediately previous value.
4. After choosing a value, recurse with `i+1`; this still permits another equal
   value at the next depth.
5. Copy `current` when its length reaches `k`.

##### Step-by-step examples

For values `[1a,1b,2]` and `k=2`, the expected value outputs are
`[[1,1],[1,2]]`.

1. **Choice: First `1`**
   - **Before:** `start=0`, `current=[]`
   - **Decision:** Choose index `0`.
   - **After:** `current=[1]`, next start `1`
   - **Known correct:** Every result in this branch begins with `1`.
2. **Choice: Second `1`**
   - **Before:** `start=1`, `current=[1]`
   - **Decision:** At a new depth, equal value is allowed.
   - **After:** Emit `[1,1]`
   - **Known correct:** The two positions can both be used.
3. **Choice: `2`**
   - **Before:** Same depth after undo
   - **Decision:** Choose index `2`.
   - **After:** Emit `[1,2]`
   - **Known correct:** The other distinct suffix is included.
4. **Choice: First-depth index `1`**
   - **Before:** `start=0`, prior value was `1`
   - **Decision:** Skip equal `1`; it would repeat both outputs above.
   - **After:** Continue to `2`
   - **Known correct:** No distinct value output is lost.

For `k=0`, the one valid combination is empty. For `k>len(values)`, no branch
can reach length `k`, so the output is empty.

##### Why it works

Increasing indexes ensure one input position is used at most once. At any
depth, adjacent equal candidates would create the same next value and leave
equivalent remaining suffixes, so keeping only the first removes duplicates
without losing a value sequence. The skip is depth-local: after one `1` is in
`current`, the next depth may choose the other position, which preserves
`[1,1]`.

##### Go example

```go
import "sort"

func uniqueCombinations(values []int, k int) [][]int {
	if k < 0 || k > len(values) {
		return nil
	}
	// Sort a copy so equal values become adjacent without changing the caller's
	// slice. Adjacency makes duplicate branches easy to skip.
	sorted := append([]int(nil), values...)
	sort.Ints(sorted)
	result := make([][]int, 0)
	current := make([]int, 0, k)
	var build func(int)
	build = func(start int) {
		if len(current) == k {
			// Save an independent copy because current is reused during backtracking.
			result = append(result, append([]int{}, current...))
			return
		}
		for i := start; i < len(sorted); i++ {
			if i > start && sorted[i] == sorted[i-1] {
				// At one recursion depth, choosing the next equal value would create
				// the same combination as choosing the previous one.
				continue
			}
			// Choose this value, recurse using only later positions, then undo the
			// choice so the next candidate starts from the same prefix.
			current = append(current, sorted[i])
			build(i + 1)
			current = current[:len(current)-1]
		}
	}
	build(0)
	return result
}
```

##### Work and memory

Sorting takes `O(n log n)` time and the copied slice uses `O(n)` auxiliary
space. If `U` distinct outputs are returned, copying them costs `O(kU)` time
and output space. The search can visit up to
`sum(C(n,d), d=0..k)` states before duplicate pruning; recursion and
`current` use `O(k)` additional space.

##### When to use it

Use depth-local duplicate skipping when outputs are value sequences, input
positions may be used once, and equal values should not create repeated rows.

##### When not to use it

Do not deduplicate when labeled positions are supposed to produce distinct
answers. Do not use this exact rule for unlimited reuse, ordered sequences, or
frequency-only counting without deriving the matching state rules.

##### Common mistakes and edge cases

> **Warning:** Skipping every equal value globally loses `[1,1]` from
> `[1,1,2]`. Skip only equal candidates competing for the same output depth.

> **Warning:** Sorting the caller's slice in place changes observable input.
> The example copies before sorting.

Test all-equal input, `k=0`, `k=n`, and invalid `k`.

##### Related material

- [A.7 Combinatorics and output size](#a7-combinatorics-and-output-size) owns
  output-size accounting.
- [C.1.3.19 Backtracking](#c1319-backtracking) owns choose, recurse, and undo.
- [D.3 Sorting without unexpected mutation](#d3-sorting-without-unexpected-mutation)
  explains the copy-before-sort boundary.

<a id="integer-digit-processing"></a>

### A.8 Integer Digit Processing

#### The problem it solves

We need to inspect or rearrange the decimal digits of an integer that already
fits in memory, without converting it to text and without overflowing while
building a result.

<a id="a81-mental-model"></a>

#### The idea in simple terms

Peel digits from the right like cards from the top of a pile. For a nonnegative
integer, `value%10` extracts the last decimal digit and `value/10` removes it.
Repeated extraction is **integer digit processing**.

<a id="a84-preventing-integer-overflow"></a>

#### How it works

For reversal with a fixed upper `limit`:

1. Require a nonnegative input; a minus sign is not a digit.
2. Extract `digit := value%10`.
3. Before appending, reject when `reversed > (limit-digit)/10`.
4. Set `reversed = reversed*10+digit` and remove the source digit with
   `value /= 10`.
5. Stop when `value` is zero. Input zero returns zero without entering the loop.

<a id="a82-worked-example-reversing-digits"></a>

#### Step-by-step examples

Reverse `120`; the expected integer result is `21`.

1. **Step: Extract `0`**
   - **Before:** `value=120`, `reversed=0`
   - **Decision:** Append last digit `0`, then remove it.
   - **After:** `value=12`, `reversed=0`
   - **Known correct:** One source digit has moved.
2. **Step: Extract `2`**
   - **Before:** `value=12`, `reversed=0`
   - **Decision:** Append `2`.
   - **After:** `value=1`, `reversed=2`
   - **Known correct:** Removed suffix `20` is reversed as integer `2`.
3. **Step: Extract `1`**
   - **Before:** `value=1`, `reversed=2`
   - **Decision:** Append `1`.
   - **After:** `value=0`, `reversed=21`
   - **Known correct:** Every digit has moved in reverse order.

An integer cannot retain a leading zero, so the result is `21`, not `021`.
For an overflow boundary with limit `127`, appending digit `8` to
`reversed=12` is rejected because `12>(127-8)/10`, or `12>11`; the unsafe
value `128` is never evaluated.

<a id="a83-invariant-fact-that-remains-true"></a>

#### Why it works

After `k` passes, `value` is the original number with its last `k` digits
removed, while `reversed` contains those removed digits in opposite order.
Extraction and append move exactly one digit and preserve that fact. Each pass
divides nonzero `value` by `10`, so the process finishes. At zero, all digits
have moved and the invariant describes the complete reversal.

The precheck is equivalent to `reversed*10+digit <= limit`, so every performed
append remains representable.

#### Go example

The complete implementation is owned by
[E.6 Overflow-safe integer reversal](#e6-overflow-safe-integer-reversal). Its
variables `value`, `digit`, and `reversed` match the trace, it accepts zero, and
it rejects negative or overflowing inputs.

<a id="a85-work-and-memory"></a>

#### Work and memory

For a value with `d` decimal digits, the loop performs `d` extraction and
append steps, so it takes `O(d)` time. It stores the remaining value, one digit,
and the result, so auxiliary space is `O(1)`. No string or digit slice is
allocated.

#### When to use it

Use arithmetic digit processing when the value already fits an integer type,
the task works on base-10 digits, and right-to-left order is convenient.

#### When not to use it

Do not use it for a decimal input too large for the integer type; keep that
input as text. Do not use base-10 extraction for a bit-level task, and do not
assume the same sign behavior for negative values without defining it.

#### Common mistakes and edge cases

> **Warning:** Check overflow before `reversed*10+digit`. For limit `127`,
> constructing `128` first defeats the guard.

> **Warning:** A loop written only as `for value>0` still needs an explicit
> contract for zero and negative input. The reference returns `(0,true)` for
> zero and failure for negatives.

Trailing zeroes disappear from the displayed reversal, palindromes remain the
same, and the smallest signed integer cannot be made positive in the same type.

<a id="a86-related-problems"></a>

#### Related material

- [A.4 Integer limits and bits](#a4-integer-limits-and-bits) derives the safe
  decimal-append check.
- [A.6.7 Numbers too large for an integer type](#a67-numbers-too-large-for-an-integer-type)
  handles digit text that cannot first become an integer.
- [E.6 Overflow-safe integer reversal](#e6-overflow-safe-integer-reversal)
  owns the complete Go implementation.
- [Palindrome Number](problems/solved_easy_leetcode_0009_palindrome_number.go)
  compares full reversal with reversing only half the digits.

<a id="b-data-structures"></a>

## B. Ways to store and organize information

### B.1 Arrays, slices, and strings

#### The problem it solves

Programs often need numbered storage: read or replace a value by position,
keep values in order, or expose only a consecutive part of a larger sequence.
The important Go-specific question is whether two sequence values share the
same storage, because a write through one can then change the other.

#### The idea in simple terms

Picture a row of numbered lockers. An **array** owns a fixed number of lockers.
A **slice** is a movable window onto consecutive lockers; it records where the
window begins, how many lockers are visible, and how far the storage continues.
Two slices may therefore look separate while sharing lockers underneath. A Go
**string** is a read-only sequence of bytes.

#### How it works

- An array's length is part of its type and does not change.
- A slice's **length** is the number of visible values. Its **capacity** is the
  number of positions from its start to the end of its underlying array.
- Indexing reads or writes one visible position. Slicing creates another view;
  it does not copy the visible values.
- `append` reuses the underlying array when capacity is available. Otherwise it
  allocates another array, copies the visible values, and returns a slice that
  refers to the new storage.
- A **frequency slice** uses each valid input value as an index and stores that
  value's count there. This requires a known, small, nonnegative value range.

#### Step-by-step examples

Start with `base := []int{10, 20, 30, 40}` and
`middle := base[1:3]`. The expected final values are
`base == [10, 99, 30, 50]` and `middle == [7, 30, 50, 60]`.

1. **Step: Create view**
   - **Before:** `base=[10,20,30,40]`
   - **Decision:** `[1:3]` exposes positions 1 and 2
   - **After:** `middle=[20,30]`, length 2, capacity 3
   - **Known correct:** Both slices refer to `base[1]` and `base[2]`
2. **Step: Replace**
   - **Before:** `middle[0]=20`
   - **Decision:** Write through the shared position
   - **After:** `base=[10,99,30,40]`
   - **Known correct:** `middle[0]` and `base[1]` are the same locker
3. **Step: Append `50`**
   - **Before:** Length 2, capacity 3
   - **Decision:** One unused shared position remains
   - **After:** `base=[10,99,30,50]`, `middle=[99,30,50]`
   - **Known correct:** The first append reused `base`
4. **Step: Append `60`**
   - **Before:** Length equals capacity
   - **Decision:** Growth requires new storage
   - **After:** `middle=[99,30,50,60]`; `base` unchanged
   - **Known correct:** Later `middle` writes cannot reach `base`
5. **Step: Replace again**
   - **Before:** Separate storage
   - **Decision:** Set `middle[0]=7`
   - **After:** `middle=[7,30,50,60]`; `base=[10,99,30,50]`
   - **Known correct:** The two slices no longer share storage

Boundary example: input values `[2, 1, 2]` with promised range `0..2` produce
frequency slice `[0, 1, 2]`. Value `-1` or `3` breaks that promise and cannot be
used as an index into a length-3 slice.

#### Why it works

A slice operation is predictable when one fact remains true: every visible
position maps to its start position plus the index in one underlying array.
Creating a subslice changes the view, not that mapping. A write follows the
mapping to shared storage. Once `append` allocates, the returned slice maps to a
different array, so the old view cannot observe later writes. A frequency slice
works because each permitted value has exactly one reserved index.

#### Go example

```go
func sliceSharing() (base []int, middle []int) {
	// middle is a two-element view whose capacity still reaches base's final
	// element, so both slices initially share the same backing array.
	base = []int{10, 20, 30, 40}
	middle = base[1:3]
	// This indexed write is visible through base because storage is shared.
	middle[0] = 99
	// The first append fits the remaining capacity and overwrites base[3].
	middle = append(middle, 50)
	// The second append exceeds capacity, so append allocates new storage and
	// copies middle. Later writes through middle no longer affect base.
	middle = append(middle, 60)
	// This write changes only the newly allocated middle array, demonstrating
	// that the second append ended storage sharing.
	middle[0] = 7
	return base, middle
}
```

Indexing a string with `text[i]` similarly returns one byte, not necessarily a
complete visible character.

#### Work and memory

Indexing and taking a subslice do constant work and allocate no element copy:
`O(1)` time and `O(1)` auxiliary space. Copying `n` values takes `O(n)` time and
`O(n)` result space. One `append` is usually `O(1)` but can copy all current
values, so a sequence of appends is `O(1)` amortized time per value. A frequency
slice for range size `r` uses `O(r)` memory and counting `n` values takes
`O(n + r)` time when initialization or a full readout is included.

#### When to use it

- Positions are meaningful and direct indexed access is needed.
- Values must remain in order or occupy contiguous storage.
- The input range is small enough for a frequency slice.
- A cheap view of consecutive existing values is safe and useful.

#### When not to use it

- Use a map when keys are sparse, negative, or too large for practical indexed
  storage.
- Use linked nodes when frequent reconnection of already-known positions
  matters more than indexed access.
- Copy a subslice when it must outlive a large backing array or must not share
  future mutations.

#### Common mistakes and edge cases

> **Warning:** With `base := []int{1, 2, 3}` and `part := base[:2]`, appending
> `9` to `part` can overwrite `base[2]`; available capacity means `append` need
> not isolate the result.

- Indexing an empty slice or using an index equal to its length panics.
- A small long-lived slice can retain a much larger underlying array. Copy the
  needed values when the large storage should become reclaimable.
- `text[0]` reads one byte. For `text := "\u00e9"`, that byte is not the whole
  character; choose bytes, runes, or user-visible characters deliberately.
- A nil slice and a non-nil empty slice both have length zero but can differ in
  JavaScript Object Notation (JSON) output and in an application's public
  input/output contract, often called its API.

#### Related material

- [Moving values in place](#b11-moving-values-in-place) changes existing slice
  storage deliberately.
- [Slice storage and retained memory](#d9-slice-storage-and-retained-memory)
  owns the detailed Go allocation discussion.
- [Bytes, runes, and visible characters](#d6-bytes-runes-and-visible-characters)
  explains text indexing units.
- [Map presence and sets](#d4-distinguish-missing-keys-from-stored-values-map-presence-and-sets)
  contrasts dense indexed counts with sparse keys.

#### B.1.1 Moving values in place

##### The problem it solves

Sometimes a sequence must be rearranged under a memory limit that rules out a
second sequence of similar size. The input storage must become the output.

##### The idea in simple terms

Move furniture within one room instead of renting another room. An algorithm
works **in place** when it reuses the input's storage and needs only a fixed
amount of separate state. Reversing ranges can place a rotated sequence without
building another sequence.

##### How it works

To rotate a nonempty slice of length `n` right by a nonnegative `k` positions:

1. Replace `k` by `k % n`, which names the same rotation within one cycle.
2. Reverse the whole slice.
3. Reverse the first `k` values.
4. Reverse the remaining values.

Other in-place methods use one read and one write position, or maintain
classified regions around a comparison value called a **pivot**.

##### Step-by-step examples

Rotate `[1, 2, 3, 4, 5]` right by `7`; the expected output is
`[4, 5, 1, 2, 3]`.

1. **Step: Normalize**
   - **Before:** `k=7`, `n=5`
   - **Decision:** Rotations repeat every five moves
   - **After:** `k=2`
   - **Known correct:** The requested result is unchanged
2. **Step: Reverse all**
   - **Before:** `[1,2,3,4,5]`
   - **Decision:** Bring the final two values to the front in reverse order
   - **After:** `[5,4,3,2,1]`
   - **Known correct:** The two result groups occupy the right regions
3. **Step: Reverse prefix**
   - **Before:** `[5,4 | 3,2,1]`
   - **Decision:** Restore the moved group's order
   - **After:** `[4,5 | 3,2,1]`
   - **Known correct:** Output positions 0 and 1 are final
4. **Step: Reverse suffix**
   - **Before:** `[4,5 | 3,2,1]`
   - **Decision:** Restore the other group's order
   - **After:** `[4,5,1,2,3]`
   - **Known correct:** Every output position is final

Boundary example: rotating `[]` by any `k` leaves it empty. The function must
return before `k %= len(values)`, because remainder by zero is invalid.

For a three-way partition of `[3, 1, 4, 3, 2]` around pivot `3`, the regions
eventually become `[1, 2 | 3, 3 | 4]`: smaller, equal, and greater.

##### Why it works

Before reversal, the sequence is `A B`, where `B` is the final `k` values.
Reversing everything gives `reverse(B) reverse(A)`; reversing each group gives
`B A`, exactly the right rotation. Each reversal moves its two positions inward,
so it finishes. For partitioning, the fact that remains true is that every
classified region satisfies its label and the unclassified region shrinks after
each step.

##### Go example

```go
func rotateRight(values []int, k int) {
	// Avoid remainder by zero; an empty slice is already fully rotated.
	if len(values) == 0 {
		return
	}
	// Rotating by a whole number of slice lengths has no effect.
	k %= len(values)
	// reverse swaps the inclusive range from both ends toward the center.
	reverse := func(left, right int) {
		for left < right {
			values[left], values[right] = values[right], values[left]
			left++
			right--
		}
	}
	// If the original slice is A B and len(B)==k, these three reversals change
	// A B -> reverse(B) reverse(A) -> B A.
	reverse(0, len(values)-1)
	reverse(0, k-1)
	reverse(k, len(values)-1)
}
```

##### Work and memory

Each value participates in a constant number of swaps, so rotating `n` values
takes `O(n)` time in best, usual, and worst cases. The indices and temporary
swap value use `O(1)` auxiliary space. The input itself holds the result.

##### When to use it

- The caller permits input mutation.
- Extra memory must stay constant while values are rearranged.
- The output contains the same values in a new order, called a permutation, or
  is a filtered prefix of the input.

##### When not to use it

- Do not mutate when callers still need the original order; allocate a result.
- Do not claim an operation is in place when recursion can retain `O(n)` calls.
- Prefer a stable copy-based method if equal values must keep their relative
  order and the in-place method does not preserve it.

##### Common mistakes and edge cases

> **Warning:** On an empty slice, `k %= len(values)` divides by zero. Test the
> length first.

- Normalize a negative `k` separately if the input contract permits it; Go's
  remainder keeps the dividend's sign.
- Reversing `[left, right]` requires inclusive endpoints; mixing that with a
  half-open range can skip or repeat an element.
- During three-way partitioning, every branch must grow a classified region or
  shrink the unclassified region, or `[3]` can loop forever.

##### Related material

- [Arrays, slices, and strings](#b1-arrays-slices-and-strings) owns the storage
  and storage-sharing rules used here.
- [Two pointers](#two-pointers) owns read/write positions and inward
  movement as a problem-solving method.
- [Sorting without unexpected mutation](#d3-sorting-without-unexpected-mutation)
  explains when to copy before an in-place library operation.

<a id="linked-lists"></a>

### B.2 Linked Lists

<a id="b21-mental-model"></a>
<a id="b22-core-operations"></a>
<a id="b23-worked-example-reversing-a-list"></a>
<a id="b24-invariants-facts-that-remain-true"></a>
<a id="b25-common-techniques"></a>
<a id="b26-work-and-memory"></a>
<a id="b27-common-mistakes"></a>
<a id="b28-related-problems"></a>

#### The problem it solves

Some ordered data must grow or be reconnected without moving every later value.
The tradeoff is that reaching a position requires following links from the
beginning instead of jumping directly to an index.

#### The idea in simple terms

Think of a treasure hunt in which each note points to the next note. A
**singly linked list** is a chain of nodes; each node stores one value and a
link named `Next`. The first node is the **head**, and `nil` marks the end. A
**doubly linked list** also links backward.

#### How it works

- Follow `Next` to visit nodes in order.
- Keep a `tail` link while building so a new node can be attached without
  walking from the head.
- Reconnect known nodes by changing links; no later values need to shift.
- A **dummy head** is a temporary node before the real head. It gives head
  insertion and middle insertion the same shape.
- To reverse, keep `previous`, `current`, and a saved `next` node. Save before
  changing the only link to the unvisited remainder.

#### Step-by-step examples

Reverse `1 -> 2 -> 3 -> nil`; the expected output is
`3 -> 2 -> 1 -> nil`.

1. **Step: 1**
   - **Before:** `previous=nil`, `current=1 -> 2 -> 3`
   - **Decision:** Save `2`, then point `1.Next` to `nil`
   - **After:** `previous=1`, `current=2 -> 3`
   - **Known correct:** `1` is a complete reversed prefix
2. **Step: 2**
   - **Before:** `previous=1`, `current=2 -> 3`
   - **Decision:** Save `3`, then point `2.Next` to `1`
   - **After:** `previous=2 -> 1`, `current=3`
   - **Known correct:** Nodes `1,2` are reversed without losing `3`
3. **Step: 3**
   - **Before:** `previous=2 -> 1`, `current=3`
   - **Decision:** Save `nil`, then point `3.Next` to `2`
   - **After:** `previous=3 -> 2 -> 1`, `current=nil`
   - **Known correct:** Every original node appears once in reverse order

Boundary example: reversing `nil` performs no loop iteration and returns
`nil`; reversing one node saves `nil`, leaves its link `nil`, and returns the
same node.

For digit lists with the ones digit first, `8 -> nil` plus `7 -> nil` writes
`5` and carries `1`, then appends the remaining carry to produce `5 -> 1`.

#### Why it works

During reversal, `previous` is exactly the already reversed part and `current`
is the first unvisited node. Saving `current.Next` preserves access to every
unvisited node; redirecting `current.Next` adds exactly one node to the reversed
part. The unvisited part shrinks, so the loop finishes. When `current` is nil,
no nodes remain outside `previous`, making `previous` the new head.

#### Go example

The [generic singly linked list reference](#e4-generic-singly-linked-list)
owns complete `FromSlice`, `ToSlice`, and `Reverse` implementations. Its
`Reverse` method uses the same `previous`, `current`, and `next` state as the
trace above.

#### Work and memory

Visiting or searching `n` nodes takes `O(n)` time and `O(1)` auxiliary space.
Inserting after a node already in hand takes `O(1)` time. Loop-based reversal
takes `O(n)` time and `O(1)` auxiliary space. Copying all values to or from a
slice uses `O(n)` time and `O(n)` result space. A recursive walk can also retain
`O(n)` unfinished calls.

#### When to use it

- Nodes at known positions must be inserted, removed, or reconnected often.
- Stable node identity matters even while order changes.
- Sequential access is sufficient and direct indexing is unnecessary.

#### When not to use it

- Use a slice when direct position lookup, cache-friendly scanning, or compact
  storage matters.
- Do not use a plain linear walk on an input that may contain a cycle unless
  the walk tracks or detects repeated nodes.
- Prefer a loop over recursion when the list can be longer than safe call-stack
  depth.

#### Common mistakes and edge cases

> **Warning:** In `1 -> 2 -> 3`, setting `1.Next = nil` before saving the old
> link loses the only route to nodes `2` and `3`.

- Conditions that read `node.Next` must first prove `node != nil`; empty and
  one-node inputs expose this error.
- Advancing the only head variable can make earlier nodes unreachable when the
  original head is still needed.
- Fast and slow positions need a loop condition that proves both the fast node
  and its next link exist before moving two steps.
- For two digit lists of different lengths, treat a missing digit as zero and
  append a final nonzero carry.

#### Related material

- [Generic singly linked list](#e4-generic-singly-linked-list) owns the complete
  reusable Go implementation.
- [Fast and slow pointers](#c1332-fast-and-slow-pointers) uses different speeds
  to detect a cycle or locate a middle node.
- [Add Two Numbers](problems/solved_easy_leetcode_0002_add_two_numbers.go)
  combines unequal digit lists and carry state.
- [Merge Two Sorted Lists](problems/solved_easy_leetcode_0021_merge_two_sorted_lists.go)
  uses a dummy head and a maintained tail.

### B.3 Queues and deques

#### The problem it solves

Some tasks must process values in arrival order, while others need efficient
changes at both ends. Repeatedly shifting a slice after removing its first value
does unnecessary work.

#### The idea in simple terms

A **queue** is a ticket line: the oldest arrival leaves first. This rule is
**first in, first out (FIFO)**. A **deque**, short for double-ended queue, is a
line with a door at each end, so values can enter or leave at either end.

#### How it works

- Add new queue values at the back and remove old values from the front.
- With a Go slice, keep a `head` index naming the next unread value instead of
  shifting all unread values.
- A circular buffer wraps positions back to the start of fixed storage.
- Two stacks can implement a queue: push into an input stack and pop from an
  output stack; transfer only when the output stack is empty.
- A deque additionally supports front insertion and back removal.

#### Step-by-step examples

Suppose `S` connects to `A` and `B`, `A` to `C`, and `B` to `D`. A queue-based
search begins with `(S,0)`, where the number is distance from `S`.

1. **Step: Visit `S`**
   - **Before:** `[(S,0)]`
   - **Decision:** Remove oldest; append its unseen neighbors
   - **After:** `[(A,1),(B,1)]`
   - **Known correct:** No distance-0 item remains
2. **Step: Visit `A`**
   - **Before:** `[(A,1),(B,1)]`
   - **Decision:** Append `(C,2)` behind existing distance-1 work
   - **After:** `[(B,1),(C,2)]`
   - **Known correct:** `A` has the fewest-edge distance 1
3. **Step: Visit `B`**
   - **Before:** `[(B,1),(C,2)]`
   - **Decision:** Append `(D,2)`
   - **After:** `[(C,2),(D,2)]`
   - **Known correct:** Every distance-1 item leaves before distance 2

Two-stack boundary trace, with the stack top at the right: add `A`, `B`, `C` to
get input `[A,B,C]`. The first removal transfers to output `[C,B,A]` and
returns `A`. Add `D` to input `[D]`; later removals return `B`, `C`, then
transfer and return `D`. An empty removal must report failure rather than index
outside either stack.

#### Why it works

The queue's defining fact is that unread items remain in arrival order. Adding
only at the back and moving the front only forward preserves it. In the
two-stack form, transfer reverses the input order once, putting its oldest item
on top of the output stack. Delaying another transfer until that stack empties
prevents newer items from passing older ones.

#### Go example

```go
func breadthFirstOrder(start int, neighbors [][]int) []int {
	// seen prevents duplicate queue entries. Marking start before traversal
	// establishes that every queued vertex is already marked.
	seen := make([]bool, len(neighbors))
	queue := []int{start}
	seen[start] = true
	order := make([]int, 0, len(neighbors))
	// head advances through the slice without shifting queued values. Appends
	// add newly discovered vertices at the back in FIFO order.
	for head := 0; head < len(queue); head++ {
		current := queue[head]
		order = append(order, current)
		for _, next := range neighbors[current] {
			if !seen[next] {
				// Mark when enqueuing so another edge cannot enqueue next again.
				seen[next] = true
				queue = append(queue, next)
			}
		}
	}
	return order
}
```

This example requires `start` and every stored neighbor to be valid indices in
`neighbors`. Marking on insertion ensures a node enters the queue at most once.

#### Work and memory

For a slice queue, advancing `head` is `O(1)` and appending is `O(1)` amortized.
The backing array uses `O(n)` memory for `n` queued values and may retain read
positions. In a two-stack queue, each value is pushed, transferred at most once,
and popped once: `m` operations take `O(m)` total time, or `O(1)` amortized per
operation, although one removal can take `O(n)`. A graph traversal also spends
`O(V + E)` time on `V` vertices and `E` stored neighbor links.

#### When to use it

- Work must happen oldest first.
- A search must process all items at one distance before the next distance.
- A sliding candidate set needs expiration from the front; use a deque when the
  back must also discard dominated candidates.

#### When not to use it

- Use a stack for most-recent-first work.
- Use a heap when priority, not arrival time, decides the next value.
- Do not use ordinary FIFO breadth-first search for minimum weighted cost when
  edge costs differ.

#### Common mistakes and edge cases

> **Warning:** Repeatedly doing `queue = queue[1:]` on a long-lived queue can
> keep the old backing array and removed pointer values alive. Clear removed
> pointers and periodically compact, or use a circular buffer.

- Check for emptiness before front or back removal.
- In graph search, marking only when removing allows two parents to enqueue the
  same node; mark when adding it.
- A circular buffer must distinguish full from empty even when head and tail
  positions match.
- Worst-case operation time and amortized time are different promises.

#### Related material

- [Breadth-first search](#c1321-breadth-first-search) owns the full search
  method and shortest-edge correctness argument.
- [Monotonic deque](#c1329-monotonic-deque) removes expired and dominated
  candidates from opposite ends.
- [Stack](#c1313-stack) provides the most-recent-first component of the
  two-stack queue.
- [Slice storage and retained memory](#d9-slice-storage-and-retained-memory)
  explains why advancing a head does not necessarily release storage.

### B.4 Trees and binary-search trees

Trees organize nodes by parent-child links. The entries below separate search,
visit order, recursive results, validation, reconstruction, and added parent
links because each operation maintains a different fact.

#### B.4.1 Terminology and search

##### The problem it solves

Hierarchical data needs names for its parts and, when values obey an ordering
rule, a way to discard irrelevant regions during lookup.

##### The idea in simple terms

Picture an upside-down family tree:

```text
          8
        /   \
       4     12
      / \    / \
     2   6  10  14
```

`8` is the **root**. `4` and `12` are its children; everything below `4` is
the subtree rooted at `4`. A **tree** has no cycles and only one route between
two nodes. A **binary tree** allows at most two children per node.

This example is a **binary-search tree (BST)**: every value in a node's left
subtree is smaller and every value in its right subtree is larger.

##### How it works

Start at the root. Return success when the current value equals the target. Go
left for a smaller target and right for a larger target. Stop with failure at a
`nil` link. This procedure requires the ordering rule to hold for whole
subtrees, not merely for direct children.

##### Step-by-step examples

Search for `10` in the pictured tree; the expected result is `true`.

1. **Step: 1**
   - **Before:** Current `8`
   - **Decision:** `10 > 8`, so no left-subtree value can match
   - **After:** Current `12`
   - **Known correct:** The entire subtree at `4` is excluded
2. **Step: 2**
   - **Before:** Current `12`
   - **Decision:** `10 < 12`, so no right-subtree value can match
   - **After:** Current `10`
   - **Known correct:** Only `12`'s left subtree remains possible
3. **Step: 3**
   - **Before:** Current `10`
   - **Decision:** Values match
   - **After:** Return `true`
   - **Known correct:** The target exists

Boundary example: searching an empty tree starts at `nil` and returns `false`.
Searching the pictured tree for `9` reaches `8`, `12`, `10`, then a nil left
link, proving `9` is absent under the ordering promise.

##### Why it works

Before each comparison, every possible match lies in the current subtree.
The BST rule makes every value on the rejected side too small or too large, so
choosing one child preserves that fact. Each step moves down one level; reaching
the target succeeds and reaching nil leaves no possible node.

##### Go example

```go
func containsBST(root *TreeNode, target int) bool {
	// The current subtree contains every remaining possible match under the
	// binary-search-tree ordering invariant.
	for root != nil {
		if target == root.Value {
			return true
		}
		if target < root.Value {
			// All right-subtree values are too large; keep only the left side.
			root = root.Left
		} else {
			// All left-subtree values are too small; keep only the right side.
			root = root.Right
		}
	}
	// A nil link means the only possible search path ended without a match.
	return false
}
```

Here `TreeNode` has `Value int` and `Left, Right *TreeNode` fields.

##### Work and memory

Search follows at most the tree height `h`, so it takes `O(h)` time and `O(1)`
auxiliary space. A balanced `n`-node BST has `h = O(log n)`; a one-sided chain
has `h = O(n)`. The best case finds the target at the root in `O(1)` time.

##### When to use it

- Values remain ordered under the stated whole-subtree rule.
- The environment supplies a balanced ordered structure or tree shape matters.
- A lookup can reject an entire side after one comparison.

##### When not to use it

- An arbitrary binary tree is not searchable this way.
- Go has no standard balanced ordered map; sorting, a heap, or binary search on
  a slice may be a better fit.
- Use node identity rather than value when duplicate values represent distinct
  objects.

##### Common mistakes and edge cases

> **Warning:** A tree with root `8`, right child `10`, and `10`'s left child
> `4` passes direct-child comparisons but is not a BST because `4` lies in
> `8`'s right subtree.

- State whether duplicates are forbidden or assigned consistently to one side.
- Do not dereference the root before checking an empty tree.
- Height-based bounds must include the possible one-sided `O(n)` case.

##### Related material

- [BST validation and costs](#b44-bst-validation-and-costs) checks the full
  ordering promise used by search.
- [Binary search](#c136-binary-search) applies the same discard-one-side idea
  to indexed sorted data.
- [Validate Binary Search Tree](problems/medium_leetcode_0098_validate_binary_search_tree.go)
  exercises ancestor bounds.

#### B.4.2 Traversal order

##### The problem it solves

A tree has no single left-to-right storage order. An operation must choose when
to handle a node relative to its children or its distance from the root.

##### The idea in simple terms

Walking through a house can record a room before entering its branches,
between branches, after both branches, or one floor at a time. These choices are
called **preorder**, **inorder**, **postorder**, and **level order** traversal.

##### How it works

- Preorder handles node, left subtree, then right subtree.
- Inorder handles left subtree, node, then right subtree.
- Postorder handles left subtree, right subtree, then node.
- Level order uses a queue to handle nodes by increasing distance from root.

Recursive depth-first orders return immediately on a nil node. Level order
adds a child only after proving it exists.

##### Step-by-step examples

For the tree in B.4.1, preorder is expected to be
`[8, 4, 2, 6, 12, 10, 14]`.

1. **Step: Root**
   - **Before:** At `8`
   - **Decision:** Preorder handles node before children
   - **After:** Output `[8]`; enter `4`
   - **Known correct:** `8` precedes both subtrees
2. **Step: Left**
   - **Before:** At subtree `4,2,6`
   - **Decision:** Apply the same node-left-right rule
   - **After:** Output `[8,4,2,6]`
   - **Known correct:** Entire left subtree is complete
3. **Step: Right**
   - **Before:** At subtree `12,10,14`
   - **Decision:** Apply the same rule
   - **After:** Output `[8,4,2,6,12,10,14]`
   - **Known correct:** Every node appears once

The same tree gives inorder `[2,4,6,8,10,12,14]`, postorder
`[2,6,4,10,14,12,8]`, and level order `[8,4,12,2,6,10,14]`. An empty tree
produces an empty result in every order.

##### Why it works

Each recursive call owns exactly one subtree. It visits that subtree's root and
delegates each child subtree once in the order promised by the traversal name.
The nil case stops at missing children. Thus every reachable node appears once
and relative order follows the placement of the root action. FIFO queue order
similarly keeps every shallower level ahead of nodes it discovers.

##### Go example

```go
func preorder(root *TreeNode) []int {
	order := []int{}
	var visit func(*TreeNode)
	visit = func(node *TreeNode) {
		// Nil children end their branch without contributing a value.
		if node == nil {
			return
		}
		// Preorder records the node before recursively visiting left and right.
		// The shared order slice accumulates values across every call.
		order = append(order, node.Value)
		visit(node.Left)
		visit(node.Right)
	}
	visit(root)
	return order
}
```

##### Work and memory

Every node is handled once, so all four traversals take `O(n)` time. A
depth-first recursion uses `O(h)` call-stack space for height `h`; the returned
order uses `O(n)` result space. Level order uses a queue holding up to the tree's
maximum width, `O(w)`, plus `O(n)` output.

##### When to use it

- Use preorder to emit a parent before descendants or pass state downward.
- Use inorder when BST values must appear sorted.
- Use postorder when a parent result depends on both child results.
- Use level order when depth or nearest distance controls processing.

##### When not to use it

- Inorder is not sorted for an arbitrary binary tree.
- Recursive traversal is unsafe for an extremely deep chain; use an explicit
  stack or loop.
- Level order does not by itself minimize weighted path cost.

##### Common mistakes and edge cases

> **Warning:** Appending the node between its recursive calls implements
> inorder, not preorder; on root `2` with children `1` and `3`, the results are
> `[1,2,3]` and `[2,1,3]` respectively.

- Decide whether nil roots return nil or a non-nil empty slice when output
  representation matters.
- Mark graph-like nodes when queued if parent links can create cycles.
- Account for output storage separately from auxiliary traversal storage.

##### Related material

- [Queues and deques](#b3-queues-and-deques) owns FIFO mechanics for level
  order.
- [Recursive return contracts](#b43-recursive-return-contracts) uses postorder
  to combine child answers.
- [Binary Tree Level Order Traversal](problems/medium_leetcode_0102_binary_tree_level_order_traversal.go)
  groups the queue by depth.

#### B.4.3 Recursive return contracts

##### The problem it solves

A tree answer often depends on answers from smaller child trees. Without a
precise meaning for each return value, code can confuse a path continuing to a
parent with an answer that joins both children.

##### The idea in simple terms

Give every recursive call the same small job and promise. A **recursive return
contract** states exactly what a call returns for the subtree rooted at its
argument. For height, the promise is "number of nodes on the longest downward
path from this node," with a nil subtree returning zero.

##### How it works

1. Define the nil result.
2. Ask the left and right child calls for values with the same meaning.
3. Combine them using only information valid at the current node.
4. Return a value that satisfies the promise for the current subtree.

For height, return `1 + max(leftHeight, rightHeight)`.

##### Step-by-step examples

On the B.4.1 tree, the expected height is `3`.

1. **Step: Leaves**
   - **Before:** Children of `2,6,10,14` are nil
   - **Decision:** Nil height is 0
   - **After:** Each leaf returns 1
   - **Known correct:** Each leaf contract is satisfied
2. **Step: Nodes `4`,`12`**
   - **Before:** Child heights are 1 and 1
   - **Decision:** Take longer child and add current node
   - **After:** Each returns 2
   - **Known correct:** Both height-2 subtrees are correct
3. **Step: Root `8`**
   - **Before:** Child heights are 2 and 2
   - **Decision:** Add root to a longest child path
   - **After:** Return 3
   - **Known correct:** The full-tree height is correct

Boundary example: a nil root returns `0`. A one-sided `1 -> 2 -> 3` chain
returns `3`, not `2`; every node on the path counts.

##### Why it works

The nil result satisfies the contract for an empty subtree. Assuming both child
calls satisfy it, every downward path from the current node continues through
exactly one child, so the longer child height plus one is the longest current
path. Each call moves to a smaller subtree and ends at nil.

##### Go example

```go
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

Every node contributes constant work, so time is `O(n)`. The call stack holds
one unfinished call per level: `O(h)`, which is `O(log n)` for a balanced tree
and `O(n)` for a chain. No collection result is allocated.

##### When to use it

- A parent answer can be built from fixed-size child answers.
- The input is naturally split into smaller subtrees.
- A precise nil result and subtree promise can be stated first.

##### When not to use it

- Use separate global or returned state when the final answer may join both
  children but the value passed upward must follow only one.
- Use an explicit stack when depth may exhaust the call stack.
- Do not force a tree contract onto a graph with cycles without reached-state
  tracking.

##### Common mistakes and edge cases

> **Warning:** Returning `leftHeight + rightHeight + 1` for height incorrectly
> gives `7` on the seven-node balanced example; one downward path cannot take
> both children.

- Define whether height counts nodes or edges before choosing the nil value.
- Do not recompute the same child result in multiple branches.
- Include call-stack space in recursive complexity.

##### Related material

- [Recursion and saved work](#a3-recursion-and-saved-work) owns call structure
  and memoization.
- [Traversal order](#b42-traversal-order) explains why child-first work is
  postorder.
- [Maximum Depth of Binary Tree](problems/easy_leetcode_0104_maximum_depth_of_binary_tree.go)
  uses this exact contract.

#### B.4.4 BST validation and costs

##### The problem it solves

Checking only a node and its children cannot detect an ordering violation
against a more distant ancestor. Validation must carry every ancestor
restriction to each descendant.

##### The idea in simple terms

Each turn down a search tree narrows an allowed number interval. Carry that
interval with the node. A value outside it breaks the **binary-search-tree
invariant**, the fact that every left-subtree value is smaller and every
right-subtree value is larger than every relevant ancestor boundary.

##### How it works

Start with bounds wider than every `int` value. At each node, reject a value not
strictly between its bounds. The left child inherits the lower bound and uses
the current value as its upper bound; the right child does the opposite.

##### Step-by-step examples

This tree is invalid:

```text
    8
     \
     10
     /
    4
```


1. **Step: Root `8`**
   - **Before:** Bounds `(-infinity,+infinity)`
   - **Decision:** `8` is inside
   - **After:** Right bound state becomes `(8,+infinity)`
   - **Known correct:** Root satisfies all restrictions
2. **Step: Node `10`**
   - **Before:** Bounds `(8,+infinity)`
   - **Decision:** `10` is inside
   - **After:** Its left child gets `(8,10)`
   - **Known correct:** Path through `10` remains valid
3. **Step: Node `4`**
   - **Before:** Bounds `(8,10)`
   - **Decision:** `4 <= 8`
   - **After:** Return `false`
   - **Known correct:** A descendant violates ancestor `8`

An empty tree is valid because it contains no violation. If duplicates are
allowed, the input contract must say which comparisons are inclusive.

##### Why it works

The bounds describe exactly the restrictions imposed by every ancestor. The
root begins unrestricted. Replacing one bound with the current value adds the
new parent restriction without losing older ones. If all nodes fit, every
ancestor relation satisfies the BST rule; one rejection supplies a concrete
counterexample. The recursion finishes at nil children.

##### Go example

```go
func validBST(root *TreeNode) bool {
	// Each call carries optional exclusive bounds imposed by all ancestors, not
	// merely by the node's direct parent.
	var valid func(*TreeNode, int, bool, int, bool) bool
	valid = func(node *TreeNode, lower int, hasLower bool, upper int, hasUpper bool) bool {
		// An empty subtree cannot violate either inherited bound.
		if node == nil {
			return true
		}
		if (hasLower && node.Value <= lower) ||
			(hasUpper && node.Value >= upper) {
			// Equality also fails because this contract requires unique values.
			return false
		}
		// The current value becomes the left subtree's upper bound and the right
		// subtree's lower bound. Both subtrees must satisfy their ranges.
		return valid(node.Left, lower, hasLower, node.Value, true) &&
			valid(node.Right, node.Value, true, upper, hasUpper)
	}
	// The root starts without either bound; the booleans make zero a usable value.
	return valid(root, 0, false, 0, false)
}
```

##### Work and memory

Validation visits each node at most once: `O(n)` time. Its recursion uses
`O(h)` stack space. It can stop near the root at the first violation, giving
best-case `O(1)` time as `n` grows; a valid tree or a late violation requires
`O(n)` time.

##### When to use it

- The entire tree must be checked against a strict ordering rule.
- Descendants can violate restrictions established several levels earlier.
- Node values fit safely in the chosen bound representation.

##### When not to use it

- An inorder monotonicity check may be simpler when only ordering matters.
- Adjust the method when duplicates have an allowed side or keys use a custom
  comparison.
- Do not use fixed sentinels equal to valid endpoint values.

##### Common mistakes and edge cases

> **Warning:** Checking only `node.Left.Value < node.Value` and
> `node.Right.Value > node.Value` accepts the pictured invalid node `4`.

- Bounds based on `int` endpoints can reject `math.MinInt` or `math.MaxInt` as
  valid values; use wider bounds or optional bounds.
- Strict and inclusive comparisons must match the duplicate policy.
- A valid empty tree should not dereference a nil root.

##### Related material

- [Terminology and search](#b41-terminology-and-search) states the ordering
  promise being validated.
- [Correctness and invariants](#a2-correctness-and-invariants) gives the general
  proof pattern.
- [Validate Binary Search Tree](problems/medium_leetcode_0098_validate_binary_search_tree.go)
  is the representative exercise.

#### B.4.5 Reconstruction from traversals

##### The problem it solves

Two visit orders may describe a tree without storing its links. The task is to
rebuild the parent-child structure without repeatedly scanning for each root.

##### The idea in simple terms

Preorder names a subtree's root first. Inorder places every left-subtree value
before that root and every right-subtree value after it. One order chooses the
root; the other cuts the remaining values into the two child regions.

##### How it works

The input must contain the same unique values in both traversals. Build a map
from each inorder value to its index. Keep `nextPreorder` and recursively own a
half-open inorder range `[left,right)`. An empty range returns nil; otherwise
consume one root, split at its inorder index, then build left before right.

##### Step-by-step examples

Input:

```text
preorder: [8, 4, 2, 6, 12, 10, 14]
inorder:  [2, 4, 6, 8, 10, 12, 14]
```

1. **Call: `build(0,7)`**
   - **Before:** Next preorder is `8`
   - **Decision:** Inorder index 3 splits root regions
   - **After:** Left `[0,3)`, right `[4,7)`
   - **Known correct:** Root and subtree sizes are fixed
2. **Call: `build(0,3)`**
   - **Before:** Next preorder is `4`
   - **Decision:** Index 1 splits its region
   - **After:** `[0,1)` and `[2,3)`
   - **Known correct:** Root `4` has children `2`,`6`
3. **Call: `build(4,7)`**
   - **Before:** After left side, next is `12`
   - **Decision:** Index 5 splits its region
   - **After:** `[4,5)` and `[6,7)`
   - **Known correct:** Root `12` has children `10`,`14`

Boundary example: two empty traversals produce nil. Duplicate traversals
`preorder=[1,1]`, `inorder=[1,1]` can describe either a left or right child and
therefore do not identify one tree without unique IDs or another rule.

##### Why it works

The recursive fact is that `[left,right)` contains exactly the inorder values
for the next preorder subtree. Preorder supplies that subtree's root; its known
inorder position uniquely divides the remaining nodes. Building the entire left
range consumes exactly the preorder values before the right subtree. Ranges
strictly shrink and empty ranges stop.

##### Go example

```go
func buildTree(preorderValues, inorderValues []int) *TreeNode {
	// Map each unique value to its inorder position so subtree splits are O(1).
	index := make(map[int]int, len(inorderValues))
	for i, value := range inorderValues {
		// The input contract promises unique values, so each position has one key.
		index[value] = i
	}
	// Preorder supplies roots in the order recursive calls need them.
	nextPreorder := 0
	var build func(int, int) *TreeNode
	build = func(left, right int) *TreeNode {
		// [left,right) is empty, so this child link is nil.
		if left == right {
			return nil
		}
		// The next preorder value is this subtree's root. Its inorder position
		// divides the remaining values into left and right subtrees.
		value := preorderValues[nextPreorder]
		nextPreorder++
		split := index[value]
		root := &TreeNode{Value: value}
		// Values before split belong to the left subtree; values after it belong
		// to the right. Recursive preorder consumption supplies each child root.
		root.Left = build(left, split)
		root.Right = build(split+1, right)
		return root
	}
	// The initial half-open range covers the entire inorder traversal.
	return build(0, len(inorderValues))
}
```

##### Work and memory

Building the index and nodes each takes `O(n)` time. The map and returned nodes
use `O(n)` space; recursion uses another `O(h)` stack space. Repeatedly scanning
inorder instead makes worst-case time `O(n^2)`.

##### When to use it

- Preorder and inorder contain the same unique node identifiers.
- The task requires the explicit tree, not only a property of the traversals.
- Linear expected-time map lookup is acceptable.

##### When not to use it

- Duplicate values make reconstruction ambiguous without extra identity rules.
- Invalid or mismatched traversals need explicit validation before the compact
  routine shown here.
- If only one traversal is given, many different trees may match it.

##### Common mistakes and edge cases

> **Warning:** Building the right subtree before the left consumes preorder in
> the wrong order; for `[8,4,2,6,12,10,14]`, value `4` belongs to the left side.

- Keep half-open boundaries consistent; the right range starts at `split+1`.
- Verify a looked-up split lies inside the current range for untrusted input.
- The index map silently overwrites duplicate keys.

##### Related material

- [Traversal order](#b42-traversal-order) defines preorder and inorder.
- [Map presence and sets](#d4-distinguish-missing-keys-from-stored-values-map-presence-and-sets)
  explains checked lookups for untrusted values.
- [Construct Binary Tree from Preorder and Inorder Traversal](problems/medium_leetcode_0105_construct_binary_tree_from_preorder_and_inorder_traversal.go)
  is the representative exercise.

#### B.4.6 Traversing through parent links

##### The problem it solves

Tree nodes usually point only downward, but a query from an internal target may
need to move to children and ancestors. The missing upward links must be made
available without revisiting nodes forever.

##### The idea in simple terms

Add a temporary "came from" address for every child. The tree can then be
viewed as a graph where a node's neighbors are its left child, right child, and
parent. A reached set prevents walking back and forth over the same link.

##### How it works

Require the target to belong to the tree rooted at `root`. First traverse from
the root and map each child pointer to its parent pointer. Then run
breadth-first search from the target. Mark node identities when they enter the
queue. One completed queue layer adds one link of distance.

##### Step-by-step examples

Find nodes two links from target `4` in the B.4.1 tree; expected output is
`[12]`.

1. **Distance: 0**
   - **Before:** Queue `[4]`, seen `{4}`
   - **Decision:** Add children `2`,`6` and parent `8`
   - **After:** Queue `[2,6,8]`
   - **Known correct:** These and only these nodes are one link away
2. **Distance: 1**
   - **Before:** Queue `[2,6,8]`
   - **Decision:** `2`,`6` add nothing; `8` adds unseen `12`
   - **After:** Queue `[12]`
   - **Known correct:** All nodes up to distance 1 are complete
3. **Distance: 2**
   - **Before:** Queue `[12]`
   - **Decision:** Requested layer reached
   - **After:** Return `[12]`
   - **Known correct:** Every returned node is exactly two links away

Boundary example: `k=0` returns the target itself. A nil target returns an
empty result rather than trying to use a nil map key as a real node.

##### Why it works

After the parent map is built, its edges exactly restore the downward links in
reverse. Breadth-first layers contain precisely the unseen nodes at that many
links from the target. Marking at insertion prevents the parent-child pair from
adding each other again. Finite nodes and one insertion per node guarantee
termination.

##### Go example

```go
func nodesAtDistance(root, target *TreeNode, k int) []int {
	// A missing target has no meaningful distance, and negative distances cannot
	// be reached by traversing edges.
	if target == nil || k < 0 {
		return []int{}
	}
	// Record upward links because ordinary tree nodes expose only downward
	// children, while distance may require moving toward an ancestor.
	parent := map[*TreeNode]*TreeNode{}
	var mapParents func(*TreeNode)
	mapParents = func(node *TreeNode) {
		// Missing children terminate their recursive branch.
		if node == nil {
			return
		}
		if node.Left != nil {
			// Store the reverse of the left-child edge so BFS can later move upward.
			parent[node.Left] = node
		}
		if node.Right != nil {
			// Store the matching reverse edge for the right child.
			parent[node.Right] = node
		}
		// Visit both child subtrees to build parent links for the entire tree.
		mapParents(node.Left)
		mapParents(node.Right)
	}
	mapParents(root)
	// Breadth-first search starts at target. Pointer identity in seen prevents
	// revisiting a node through a different one of its three possible neighbors.
	queue := []*TreeNode{target}
	seen := map[*TreeNode]bool{target: true}
	for distance := 0; distance < k; distance++ {
		// Process exactly the current layer so one outer iteration equals one edge
		// of distance from target.
		levelSize := len(queue)
		if levelSize == 0 {
			// No frontier remains, so no node exists at this or any greater distance.
			break
		}
		// Only remove the levelSize nodes that were present when this layer began.
		// Newly appended nodes belong to the next distance and wait in the queue.
		for i := 0; i < levelSize; i++ {
			// Remove the oldest current-layer node from the queue front.
			node := queue[0]
			queue = queue[1:]
			// Treat left child, right child, and parent as the node's undirected
			// neighbors so the search can move away from target in every direction.
			for _, next := range []*TreeNode{node.Left, node.Right, parent[node]} {
				if next != nil && !seen[next] {
					// Mark before enqueueing: another neighbor in this layer may also
					// point to next, but it must enter the next frontier only once.
					seen[next] = true
					queue = append(queue, next)
				}
			}
		}
	}
	// After k complete layers, the queue contains exactly the nodes at distance k.
	values := make([]int, len(queue))
	// Convert the node frontier to the value representation promised by the API.
	for i, node := range queue {
		values[i] = node.Value
	}
	return values
}
```

##### Work and memory

Parent mapping and search each visit at most `n` nodes. The search also stops
when its frontier becomes empty, so total time is `O(n)`, even when `k` is
larger than the tree's height. The parent map, reached set, queue, and output
together use `O(n)` memory; the recursive mapping also uses `O(h)` call-stack
space within that bound.

##### When to use it

- A query starts away from the root and may move toward ancestors.
- Tree nodes do not already store parent links.
- Distance is measured in equal-cost links.

##### When not to use it

- Follow existing parent links instead of rebuilding them when they are valid.
- A root-to-node path alone may use less memory for a single ancestor query.
- Weighted links require a weighted shortest-path method.

##### Common mistakes and edge cases

> **Warning:** If `4` queues parent `8` and nodes are marked only when removed,
> `8` can queue `4` again, turning the temporary graph into repeated work.

- Map pointers, not values; two distinct nodes may both store `7`.
- Reject or separately define a target outside the root's tree; this example
  requires membership.
- Decide whether negative `k` is invalid or produces an empty result.

##### Related material

- [Graphs and grids](#b5-graphs-and-grids) explains the temporary graph view.
- [Breadth-first search](#c1321-breadth-first-search) owns layer correctness.
- [Queues and deques](#b3-queues-and-deques) owns FIFO queue mechanics.

### B.5 Graphs and grids

#### The problem it solves

Many inputs describe objects connected by arbitrary allowed moves rather than
one line or one hierarchy. The representation must preserve direction, cost,
identity, and disconnected groups needed by the algorithm.

#### The idea in simple terms

A graph is a map of things and allowed moves. A thing is a **vertex** or
**node**; a move is an **edge**. A grid is often an **implicit graph**: its
neighbors are calculated from coordinates instead of stored as edge records.

#### How it works

- An **adjacency list** stores each vertex beside its neighbors.
- An **adjacency matrix** stores one cell for every possible ordered pair.
- An **edge list** stores the edges directly.
- For an undirected edge, add both directions; for a directed edge, add only
  its stated direction. A vertex's **indegree** counts incoming directed edges.
- Mark a node when scheduling it. Use an outer loop when every disconnected
  **component**, a group mutually reachable through edges, must be visited.

#### Step-by-step examples

For two-way roads `A-B`, `A-C`, and `C-D`, the adjacency list is:

```text
A: [B, C]
B: [A]
C: [A, D]
D: [C]
```

1. **Step: Add `A-B`**
   - **Before:** Empty lists
   - **Decision:** Road is two-way
   - **After:** `A:[B]`, `B:[A]`
   - **Known correct:** Both allowed moves exist
2. **Step: Add `A-C`**
   - **Before:** Above lists
   - **Decision:** Add both directions
   - **After:** `A:[B,C]`, `C:[A]`
   - **Known correct:** `A` has both neighbors
3. **Step: Add `C-D`**
   - **Before:** `C:[A]`, no `D`
   - **Decision:** Add both directions
   - **After:** `C:[A,D]`, `D:[C]`
   - **Known correct:** Every road is represented twice

A search from `A` schedules `B` and `C`, then reaches `D` through `C`.
Boundary example: adding isolated vertex `E` creates another component. A
search started only at `A` must not claim to have visited the whole graph.

For vertex order `A,B,C,D`, the same graph's adjacency matrix is:

```text
    A B C D
A   0 1 1 0
B   1 0 0 0
C   1 0 0 1
D   0 0 1 0
```

Reading row `C`, column `D` returns `1`, so edge `C-D` exists in constant time;
row `A`, column `D` returns `0`. The edge-list form is
`[(A,B),(A,C),(C,D)]`.

#### Why it works

Each stored neighbor entry corresponds to one allowed directed move; adding
both entries makes an undirected edge traversable both ways. A traversal that
marks on scheduling processes every reachable node at most once. An outer loop
starts the same argument in each still-unseen component. Grid neighbor formulas
are equivalent when they generate exactly the allowed in-bounds moves.

#### Go example

```go
func undirectedGraph(vertexCount int, edges [][2]int) [][]int {
	// Allocate one adjacency list per vertex, initially with no neighbors.
	neighbors := make([][]int, vertexCount)
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		// An undirected edge can be traversed in either direction, so store both
		// directed adjacency entries.
		neighbors[from] = append(neighbors[from], to)
		neighbors[to] = append(neighbors[to], from)
	}
	return neighbors
}
```

This compact builder requires `vertexCount >= 0` and every endpoint to be in
`0..vertexCount-1`.

#### Work and memory

For `V` vertices and `E` directed neighbor entries, an adjacency list takes
`O(V + E)` build time and space; scanning every neighbor has the same time.
An adjacency matrix uses `O(V^2)` space and initialization time but checks one
edge in `O(1)`. An edge list uses `O(E)` space. For an undirected graph, the
adjacency list stores two entries per edge, still `O(V + E)`.

#### When to use it

- Objects have many-to-many relationships or allowed transitions.
- A grid move can be modeled as an edge between cells.
- Reachability, connected groups, dependencies, or path costs matter.

#### When not to use it

- A tree is simpler when the input guarantees no cycles and one parent route.
- A dense adjacency matrix is wasteful for a **sparse graph**, one in which few
  of all possible vertex pairs have edges.
- Unweighted breadth-first search does not minimize differing edge weights.

#### Common mistakes and edge cases

> **Warning:** For an undirected road `A-B`, adding only `A -> B` makes a search
> from `B` unable to reach `A`.

- Key clone maps by original node identity, not stored value; two nodes can both
  contain `7`.
- Check row and column bounds separately for grid neighbors.
- For path-specific [backtracking](#c1319-backtracking), restore a cell after
  the path; for permanent reachability, do not restore it.
- Account for graph construction and duplicate edges in total cost.

#### Related material

- [Breadth-first search](#c1321-breadth-first-search) and
  [depth-first search](#c1320-depth-first-search) own traversal choices.
- [Bellman-Ford](#c1324-bellman-ford-shortest-paths) naturally consumes an
  edge list.
- [Clone Graph](problems/medium_leetcode_0133_clone_graph.go) requires an
  identity-keyed clone map.
- [Number of Islands](problems/medium_leetcode_0200_number_of_islands.go)
  treats a grid as an implicit graph.

### B.6 Matrix traversal and in-place state

#### The problem it solves

Grid algorithms must name cells correctly and reject invalid neighbors. They
often fail by assuming equal dimensions, a square shape, or equal row lengths.

#### The idea in simple terms

A matrix is a set of rows addressed as `matrix[row][column]`. Vertical movement
changes the row; horizontal movement changes the column. A **rectangular**
matrix has equal row lengths. Go also permits **jagged slices**, whose rows have
different lengths.

#### How it works

Check `0 <= row < len(matrix)` first, then check
`0 <= column < len(matrix[row])`. State separately whether an operation requires
a rectangle or square. For four-direction movement, add `(-1,0)`, `(1,0)`,
`(0,-1)`, and `(0,1)` and validate each candidate.

#### Step-by-step examples

For `[[1,2,3],[4,5,6]]`, query neighbors of `(0,2)`.

1. **Candidate: `(-1,2)`**
   - **Before:** Row `-1`
   - **Decision:** Row is outside `0..1`
   - **After:** Reject
   - **Known correct:** No read occurred outside storage
2. **Candidate: `(1,2)`**
   - **Before:** Row 1 exists, length 3
   - **Decision:** Column 2 exists
   - **After:** Accept value `6`
   - **Known correct:** Down neighbor is valid
3. **Candidate: `(0,1)`**
   - **Before:** Row 0 exists
   - **Decision:** Column 1 exists
   - **After:** Accept value `2`
   - **Known correct:** Left neighbor is valid
4. **Candidate: `(0,3)`**
   - **Before:** Row 0 length 3
   - **Decision:** Column equals length
   - **After:** Reject
   - **Known correct:** Right edge is enforced

Boundary example: in jagged `[[1,2],[]]`, position `(1,0)` has a valid row but
an invalid column. Using the first row's length for all rows would panic.

#### Why it works

Every valid cell belongs to an existing row and an existing position in that
specific row. Testing in that order proves both conditions before indexing.
Direction offsets generate each intended adjacent coordinate exactly once;
bounds checks remove only coordinates outside storage.

#### Go example

```go
func inBounds(matrix [][]int, row, column int) bool {
	// Prove the row exists before evaluating len(matrix[row]); && short-circuits
	// left to right. Then validate the column against that row's actual length.
	return row >= 0 && row < len(matrix) &&
		column >= 0 && column < len(matrix[row])
}
```

Go's left-to-right `&&` evaluation stops after a false row check, so the code
does not evaluate `matrix[row]` for an invalid row.

#### Work and memory

One bounds check takes `O(1)` time and space. Visiting every cell takes time
proportional to the total number of cells, `O(cells)`, which is `O(rows*cols)`
for a rectangle. A returned copy needs `O(cells)` result memory; in-place
visits may use `O(1)` auxiliary state unless traversal adds a stack or queue.

#### When to use it

- Input is indexed by row and column.
- Neighbor movement or layer boundaries define the work.
- Shape promises can be stated before the algorithm.

#### When not to use it

- Flattened storage may be simpler when row structure has no meaning.
- Do not apply square-only transformations to a rectangle.
- A jagged input cannot use one global column count unless first validated.

#### Common mistakes and edge cases

> **Warning:** On a 2-by-3 matrix, checking `column < len(matrix)` rejects valid
> column 2 because the row count is 2.

- Empty `matrix` has no `matrix[0]`.
- An existing row may itself be empty.
- Keep `(row,column)` order consistent; swapping them often survives square
  tests and fails only on rectangles.

#### Related material

- [Graphs and grids](#b5-graphs-and-grids) owns implicit neighbor edges.
- [Rows, columns, and shrinking boundaries](#b61-rows-columns-and-shrinking-boundaries)
  applies explicit rectangle boundaries.
- [Matrix traversal problems](problems/medium_leetcode_0054_spiral_matrix.go)
  expose row-column mistakes.

#### B.6.1 Rows, columns, and shrinking boundaries

##### The problem it solves

A rectangular grid can be read layer by layer, but the remaining region changes
shape. Without explicit boundaries and crossing checks, one-row or one-column
remainders are emitted twice.

##### The idea in simple terms

Peel the outer frame from a picture, then repeat on the smaller picture. Four
inclusive positions named `top`, `bottom`, `left`, and `right` describe the
unread rectangle.

##### How it works

Read the top edge left-to-right and increment `top`; read the right edge
top-to-bottom and decrement `right`. If rows remain, read the bottom edge in
reverse and decrement `bottom`. If columns remain, read the left edge upward
and increment `left`.

##### Step-by-step examples

For `[[1,2,3],[4,5,6]]`, expected spiral order is `[1,2,3,6,5,4]`.

1. **Step: Top**
   - **Before:** `top=0,bottom=1,left=0,right=2`
   - **Decision:** Emit row 0
   - **After:** Output `[1,2,3]`, `top=1`
   - **Known correct:** Old top is complete
2. **Step: Right**
   - **Before:** Remaining row 1
   - **Decision:** Emit column 2
   - **After:** Output `[1,2,3,6]`, `right=1`
   - **Known correct:** Old right is complete
3. **Step: Bottom**
   - **Before:** `top <= bottom`
   - **Decision:** Emit row 1 right-to-left
   - **After:** Output `[1,2,3,6,5,4]`, `bottom=0`
   - **Known correct:** Every cell is complete

Boundary example: `[[1,2,3]]` emits its top edge and stops because `top` then
passes `bottom`; it does not emit `3,2,1` again.

##### Why it works

Before each round, everything outside the four boundaries has been emitted
exactly once and everything inside remains unread. Each edge pass emits one
boundary and moves that boundary inward. Crossing checks prevent overlap. The
remaining rectangle shrinks, so all cells are eventually emitted once.

##### Go example

```go
func spiralOrder(matrix [][]int) []int {
	// The second condition may read matrix[0], so first prove a row exists.
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
		// Consume the right edge from top to bottom, then remove that column.
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
			// A left column remains; traverse upward without repeating corners.
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

Every cell is appended once, so time is `O(rows*cols)`. Boundary variables use
`O(1)` auxiliary space, and the returned order uses `O(rows*cols)` result space.

##### When to use it

- Input is a validated rectangle.
- Output follows successive outer layers.
- Four monotonically shrinking boundaries describe unread work.

##### When not to use it

- Validate or adapt the method for jagged rows.
- Use ordinary row-by-row loops, also called row-major order, when layer order
  is irrelevant.
- A square-only rotation has a different destination mapping.

##### Common mistakes and edge cases

> **Warning:** Omitting the post-increment `top <= bottom` check emits a
> one-row matrix once forward and once backward.

- Handle an empty matrix before reading `matrix[0]`.
- Use the column count from the validated rectangular row, not the row count.
- Output allocation is not auxiliary space, but it still consumes memory.

##### Related material

- [Matrix traversal and in-place state](#b6-matrix-traversal-and-in-place-state)
  owns shape and coordinate rules.
- [Spiral Matrix](problems/medium_leetcode_0054_spiral_matrix.go) is the
  representative boundary exercise.

#### B.6.2 Layer transformations

##### The problem it solves

A square matrix must sometimes be rotated while reusing its storage. Writing
destinations immediately can overwrite values that have not moved yet.

##### The idea in simple terms

First reflect the square across its main diagonal, then mirror each row. The
first operation is a **transpose**, which moves `(row,column)` to
`(column,row)`; the second completes a clockwise quarter turn.

##### How it works

Require an `n`-by-`n` matrix. Swap only pairs above the diagonal with their
partners below it. Then reverse every row. Each swap keeps both old values
available, and no pair is transposed twice.

##### Step-by-step examples

Rotate `[[1,2],[3,4]]`; expected output is `[[3,1],[4,2]]`.

1. **Step: Transpose**
   - **Before:** `[[1,2],[3,4]]`
   - **Decision:** Swap `(0,1)` and `(1,0)`
   - **After:** `[[1,3],[2,4]]`
   - **Known correct:** Every value has exchanged row and column
2. **Step: Reverse rows**
   - **Before:** `[[1,3],[2,4]]`
   - **Decision:** Mirror each row
   - **After:** `[[3,1],[4,2]]`
   - **Known correct:** `(r,c)` reached clockwise destination `(c,n-1-r)`

Boundary example: a 1-by-1 matrix remains unchanged. A 2-by-3 rectangle cannot
be rotated in place into the same row layout because its output shape is 3-by-2.

##### Why it works

Transpose sends `(r,c)` to `(c,r)`. Reversing the destination row sends its
column `r` to `n-1-r`. The composition is `(r,c) -> (c,n-1-r)`, exactly a
clockwise rotation. Swapping only `column > row` covers each off-diagonal pair
once; row reversal covers every destination once.

##### Go example

```go
func rotateSquare(matrix [][]int) {
	// Transpose across the main diagonal. Starting above the diagonal ensures
	// each off-diagonal pair is swapped exactly once.
	for row := 0; row < len(matrix); row++ {
		for column := row + 1; column < len(matrix); column++ {
			matrix[row][column], matrix[column][row] =
				matrix[column][row], matrix[row][column]
		}
	}
	// Mirroring every transposed row completes the clockwise coordinate mapping
	// (row,column) -> (column,n-1-row).
	for row := range matrix {
		for left, right := 0, len(matrix[row])-1; left < right; left, right = left+1, right-1 {
			matrix[row][left], matrix[row][right] = matrix[row][right], matrix[row][left]
		}
	}
}
```

##### Work and memory

Transpose and row reversal each touch `O(n^2)` cells, so total time is
`O(n^2)`. Loop indices and swap temporaries use `O(1)` auxiliary space.

##### When to use it

- The input is square and may be mutated.
- A clockwise rotation must reuse the existing storage.
- Coordinate mapping can be proved before coding.

##### When not to use it

- Allocate a differently shaped result for a non-square rectangle.
- Do not mutate when callers retain the original matrix.
- Use a direct copy when clarity matters more than constant auxiliary space.

##### Common mistakes and edge cases

> **Warning:** Transposing every `(row,column)` pair swaps each pair twice and
> restores the original matrix. Visit only one side of the diagonal.

- Validate square shape for untrusted input before indexing `matrix[column]`.
- Reversing columns after transpose rotates counterclockwise, not clockwise.
- Square tests can hide row-column naming mistakes; trace coordinates.

##### Related material

- [Moving values in place](#b11-moving-values-in-place) owns the mutation
  tradeoff.
- [Rotate Image](problems/medium_leetcode_0048_rotate_image.go) exercises
  the transpose-and-reverse mapping.

#### B.6.3 Reusing the first row and column as markers

##### The problem it solves

When one original cell determines changes to its whole row and column, applying
the change immediately destroys evidence about which cells were originally
special. Separate marker arrays solve this but use extra memory.

##### The idea in simple terms

Use the matrix's first row and first column as a notepad. An interior zero marks
its row at column zero and its column at row zero. Two booleans remember whether
the notepad row and column themselves originally contained zero.

##### How it works

Scan the first row and column for their original flags. Scan interior cells and
write markers. In a second interior pass, zero a cell when either marker is
zero. Finally zero the first row and first column according to their saved
flags.

##### Step-by-step examples

Input and expected output:

```text
original           markers recorded       final
[1 2 3]            [1 2 0]                [1 2 0]
[4 5 0]     ->     [0 5 0]         ->     [0 0 0]
[7 8 9]            [7 8 9]                [7 8 0]
```

1. **Step: Inspect `(1,2)`**
   - **Before:** Original zero
   - **Decision:** Record row 1 at `(1,0)` and column 2 at `(0,2)`
   - **After:** Marker cells become zero
   - **Known correct:** The affected row and column are remembered
2. **Step: Apply interior**
   - **Before:** All originals inspected
   - **Decision:** Read markers, not new zeros
   - **After:** Row 1 and column 2 interior cells become zero
   - **Known correct:** No generated zero creates a new marker
3. **Step: Finish borders**
   - **Before:** Original flags are false
   - **Decision:** Apply saved border decisions
   - **After:** Final matrix shown
   - **Known correct:** Every required cell and only it is zero

Boundary example: if `(0,1)` is originally zero, the separate first-row flag is
required; cell `(0,0)` cannot simultaneously encode both the first row and
first column facts.

##### Why it works

After the first pass, an interior row or column has a zero marker exactly when
an original interior zero requires it. No output changes occur before every
original is inspected. The second pass follows those complete markers. Saved
booleans restore the two border decisions that share `(0,0)`.

##### Go example

```go
func setZeroes(matrix [][]int) {
	// Guard matrix[0] and treat a matrix with no cells as already complete.
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}
	// The first row and column will store markers, so save their original zero
	// status separately before changing either border.
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
				// One original zero marks both its row and its column on the borders.
				matrix[row][0], matrix[0][column] = 0, 0
			}
		}
	}
	// Apply the completed markers only after all original interior cells have
	// been inspected, preventing generated zeros from becoming new evidence.
	for row := 1; row < rows; row++ {
		for column := 1; column < columns; column++ {
			if matrix[row][0] == 0 || matrix[0][column] == 0 {
				// Either marker proves this cell shares a row or column with an
				// original zero.
				matrix[row][column] = 0
			}
		}
	}
	// Restore the two border decisions from their saved booleans; matrix[0][0]
	// could not represent both facts by itself.
	if firstRowZero {
		// Apply the saved row decision only after its marker values are no longer needed.
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

The passes inspect `rows*columns` cells a constant number of times, so time is
`O(rows*columns)`. Two booleans and loop positions use `O(1)` auxiliary space;
the matrix stores its own markers.

##### When to use it

- The input is a mutable nonempty rectangle.
- Marker values can safely occupy reserved existing cells.
- Constant auxiliary space is required.

##### When not to use it

- Separate marker slices are clearer when `O(rows+columns)` memory is allowed.
- Do not overwrite border cells before saving their original meaning.
- Jagged rows need validation or a different representation.

##### Common mistakes and edge cases

> **Warning:** Zeroing row 1 immediately after seeing `(1,2)` creates a zero at
> `(1,0)` that a later naive scan may mistake for an original zero in column 0.

- Handle an empty matrix before `matrix[0]`.
- One flag cannot represent both the first row and first column.
- Apply interior cells before erasing marker borders.

##### Related material

- [Matrix traversal and in-place state](#b6-matrix-traversal-and-in-place-state)
  owns rectangular-shape checks.
- [Set Matrix Zeroes](problems/medium_leetcode_0073_set_matrix_zeroes.go)
  is the representative marker exercise.

#### B.6.4 Keeping current and next state together

##### The problem it solves

In a simultaneous grid update, every next value must be decided from the same
old grid. Replacing cells one by one can make later decisions read a mixture of
old and new rounds.

##### The idea in simple terms

Write the old and new status on the same sticky note. If inputs use `0` and
`1`, reserve `2` for "was 0, becomes 1" and `3` for "was 1, becomes 0." Reads
decode the old status; a final pass keeps only the new status.

##### How it works

Define every temporary state's old and new meanings first. During the decision
pass, count a neighbor as formerly live when its value is `1` or `3`. Encode
changes but leave unchanged cells as `0` or `1`. During the final pass, convert
`2` to `1` and `3` to `0`.

##### Step-by-step examples

Suppose a one-row rule says a cell becomes live when exactly one old neighbor
is live. Input `[1,1,0]` should simultaneously become `[1,1,1]`.

1. **Step: Cell 0**
   - **Before:** `[1,1,0]`
   - **Decision:** One old live neighbor; stays 1
   - **After:** `[1,1,0]`
   - **Known correct:** Cell 0 next state is encoded
2. **Step: Cell 1**
   - **Before:** `[1,1,0]`
   - **Decision:** One old live neighbor; stays 1
   - **After:** `[1,1,0]`
   - **Known correct:** Cells 0 and 1 are encoded
3. **Step: Cell 2**
   - **Before:** `[1,1,0]`
   - **Decision:** Old neighbor 1 is live
   - **After:** `[1,1,2]`
   - **Known correct:** `2` preserves old 0 and new 1
4. **Step: Finalize**
   - **Before:** `[1,1,2]`
   - **Decision:** Keep only new meanings
   - **After:** `[1,1,1]`
   - **Known correct:** All cells show the same next round

Boundary example: an empty grid performs neither pass. If input may already
contain `2` or `3`, these temporary codes are ambiguous and cannot be used.

##### Why it works

During the first pass, every stored value can still be decoded to its original
state, even after its next state is added. Therefore every decision observes
one consistent old round. Each cell receives one next-state encoding. The final
pass is safe because no further old-state decisions remain.

##### Go example

```go
func updateRow(values []int) {
	// Values 1 and 3 were live in the old generation; 3 additionally records
	// that the cell will be dead in the new generation.
	oldLive := func(value int) bool {
		return value == 1 || value == 3
	}
	// First pass: decide every next state while preserving a decodable old state.
	for i, value := range values {
		liveNeighbors := 0
		// Count the left neighbor only when its index exists, decoding temporary
		// state 3 as formerly live.
		if i > 0 && oldLive(values[i-1]) {
			liveNeighbors++
		}
		// Apply the identical check to the right neighbor without passing the end.
		if i+1 < len(values) && oldLive(values[i+1]) {
			liveNeighbors++
		}
		nextLive := liveNeighbors == 1
		if value == 0 && nextLive {
			// 2 means old dead, new live.
			values[i] = 2
		} else if value == 1 && !nextLive {
			// 3 means old live, new dead.
			values[i] = 3
		}
	}
	// Second pass: every decision is complete, so discard the old-state encoding.
	for i, value := range values {
		if value == 2 {
			// A dead-to-live transition becomes its final live value.
			values[i] = 1
		} else if value == 3 {
			// A live-to-dead transition becomes its final dead value.
			values[i] = 0
		}
	}
}
```

##### Work and memory

With a fixed number of neighbors, each of `n` cells takes constant work, so
time is `O(n)`; for a matrix, `n = rows*columns`. Temporary codes and counters
use `O(1)` auxiliary space. A clearer separate next-state grid would use
`O(n)` additional space.

##### When to use it

- Updates are logically simultaneous.
- The input value range leaves unambiguous temporary codes.
- Mutation is allowed and auxiliary memory must stay constant.

##### When not to use it

- Use a separate next-state grid when clarity or unrestricted input values
  matter more than memory.
- Do not encode before documenting how every temporary value decodes.
- Concurrent writers need synchronization beyond this single-threaded method.

##### Common mistakes and edge cases

> **Warning:** Treating temporary `3` as dead during the decision pass makes a
> later cell read the new state instead of the required old live state.

- Finalize only after every decision is recorded.
- Ensure temporary values cannot appear in valid input.
- Empty and single-cell inputs still need well-defined neighbor rules.

##### Related material

- [Matrix traversal and in-place state](#b6-matrix-traversal-and-in-place-state)
  owns coordinate and shape checks.
- [Moving values in place](#b11-moving-values-in-place) discusses mutation
  contracts.
- [Game of Life](problems/medium_leetcode_0289_game_of_life.go) applies this
  encoding to eight neighbors.

### B.7 Composing data structures

#### The problem it solves

One structure may make lookup fast but make ordering slow, while another has
the opposite strengths. Some operation sets need two synchronized views of the
same logical values.

#### The idea in simple terms

Pair an address book with a waiting line. The book finds a person quickly; the
line records arrival order. **Composing data structures** means combining such
views and defining the rule that keeps them in agreement.

#### How it works

Choose one structure per required operation, then state a cross-structure fact:

- A slice plus map supports random indexing and value-to-index lookup.
- A map plus doubly linked list supports direct lookup and recency order for a
  **least-recently-used (LRU) cache**, which evicts the item unused for the
  longest time. A permanent dummy end node, called a **sentinel**, removes end
  special cases.
- A max-heap, which exposes the largest value, and a min-heap, which exposes
  the smallest, keep lower and upper halves for a streaming median.
- A map can select one key's sorted history before binary search.
- A value stack plus minimum stack exposes the current minimum.

Update every affected view before the public operation returns.

#### Step-by-step examples

A random-access set stores `values=[10,20,30]` and
`indexByValue={10:0,20:1,30:2}`. Delete `20`; expected retained values are
`[10,30]` with map `{10:0,30:1}`.

1. **Step: Locate**
   - **Before:** Map says `20 -> 1`
   - **Decision:** Direct lookup finds the removal position
   - **After:** `removeIndex=1`
   - **Known correct:** No scan is needed
2. **Step: Move last**
   - **Before:** Last is `30` at 2
   - **Decision:** Fill position 1 and repair `30`'s index
   - **After:** Slice `[10,30,30]`, map has `30 -> 1`
   - **Known correct:** Every retained position already agrees
3. **Step: Shrink**
   - **Before:** Last slot is spare
   - **Decision:** Drop it and delete key `20`
   - **After:** `[10,30]`, `{10:0,30:1}`
   - **Known correct:** Both views agree completely

Boundary example: deleting final value `30` from `[10,20,30]` moves it onto
itself, then shrinks and deletes its key. The same steps remain valid.

#### Why it works

The fact that remains true between operations is
`indexByValue[values[i]] == i` for every retained position. Lookup identifies
the only hole. Moving the last value fills it without shifting other values,
and repairing that one map entry restores agreement for the moved value.
Shrinking removes the duplicate slot and deleting the removed key completes the
invariant. Each operation changes only a fixed number of positions.

#### Go example

```go
// RandomSet keeps the invariant indexByValue[values[i]] == i.
type RandomSet struct {
	values       []int
	indexByValue map[int]int
}

func (set *RandomSet) Remove(value int) bool {
	// The map finds the removal slot without scanning the slice.
	removeIndex, exists := set.indexByValue[value]
	if !exists {
		// Neither representation changes when the requested value is absent.
		return false
	}
	// Move the final value into the hole and repair only its map entry. This also
	// works when removing the final value because it is moved onto itself.
	lastIndex := len(set.values) - 1
	lastValue := set.values[lastIndex]
	set.values[removeIndex] = lastValue
	set.indexByValue[lastValue] = removeIndex
	// Shrink away the duplicate final slot, then remove the requested key.
	set.values = set.values[:lastIndex]
	delete(set.indexByValue, value)
	return true
}
```

#### Work and memory

Map lookup, one slice replacement, and map updates take expected `O(1)` time;
"expected" relies on ordinary hash-map behavior and is not a worst-case
per-operation guarantee. The two views store `O(n)` values. An LRU cache also
uses `O(n)` map and list storage. Two heaps use `O(n)` total storage and
`O(log n)` time per insertion.

#### When to use it

- Required operations have conflicting access-order needs.
- A clear invariant can relate the structures after every operation.
- The memory for multiple views is justified by repeated faster operations.

#### When not to use it

- One structure already meets every required bound.
- Updates cannot repair every view before callers can observe the operation.
- Duplicates are required but the value-to-one-index map does not represent
  them; use sets of indices or another design.

#### Common mistakes and edge cases

> **Warning:** Moving `30` into index 1 without changing map entry `30 -> 2`
> makes the next removal index outside the shortened slice.

- Initialize maps before assignment.
- An LRU access must update both map membership and list recency.
- State whether random choice must be uniform and use an unbiased valid index.
- For two heaps, maintain both size balance and ordering between halves.

#### Related material

- [Map presence and sets](#d4-distinguish-missing-keys-from-stored-values-map-presence-and-sets)
  owns Go map lookup behavior.
- [Heap or priority queue](#c1327-heap-or-priority-queue) owns heap operations.
- [LRU Cache](problems/medium_leetcode_0146_lru_cache.go) combines a map and
  doubly linked list.
- [Insert Delete GetRandom O(1)](problems/medium_leetcode_0380_insert_delete_getrandom_o1.go)
  uses the traced slice-map invariant.

### B.8 Streaming and online processing

#### The problem it solves

Values may arrive over time while an answer is required after each arrival.
The full input may not exist yet, and retaining or rescanning all earlier values
may exceed time or memory limits.

#### The idea in simple terms

Keep a small scoreboard and update it when each event arrives. An **online
algorithm** maintains a valid answer without assuming future input. A
**prefix** is everything received so far. A **batch algorithm** waits for the
complete input before beginning.

#### How it works

Define what the state means after every prefix and when old information can be
discarded. For the first non-repeating item, keep a count map and a FIFO queue
of first occurrences. Increment the new item's count, enqueue it only on its
first occurrence, then remove repeated candidates from the front.

Other summaries pair a recent-window queue with counts, two heaps for a running
median, a size-`k` heap for top items, or one next candidate per sorted stream.

#### Step-by-step examples

For stream `a, b, a`, report the first non-repeating letter after each arrival;
expected reports are `a, a, b`.

1. **Arrival: `a`**
   - **Before:** Counts `{}`, queue `[]`
   - **Decision:** First occurrence; count and enqueue
   - **After:** `{a:1}`, `[a]`, report `a`
   - **Known correct:** Front is earliest unique item
2. **Arrival: `b`**
   - **Before:** `{a:1}`, `[a]`
   - **Decision:** First occurrence; enqueue behind `a`
   - **After:** `{a:1,b:1}`, `[a,b]`, report `a`
   - **Known correct:** Older unique `a` remains first
3. **Arrival: `a`**
   - **Before:** `{a:1,b:1}`, `[a,b]`
   - **Decision:** Count becomes 2; discard repeated front
   - **After:** `{a:2,b:1}`, `[b]`, report `b`
   - **Known correct:** Queue front is again earliest unique item

Boundary example: stream `a, a` reports `a`, then no value. The queue becomes
empty after the second arrival; code must represent "none" separately from a
valid zero-value item.

#### Why it works

Counts describe every received item. Queue order records first-arrival order,
and only first occurrences enter it. Removing front items while their count is
not one discards candidates proved invalid. The first remaining item is
therefore the earliest received value with count one. Each arrival extends the
prefix, and every queued item can be removed only once.

#### Go example

```go
func firstUniqueReports(stream []byte) []int {
	// counts determines uniqueness; queue preserves first-arrival order among
	// values that have appeared at least once.
	counts := map[byte]int{}
	queue := []byte{}
	reports := make([]int, 0, len(stream))
	for _, value := range stream {
		counts[value]++
		if counts[value] == 1 {
			// Enqueue only the first occurrence, so each value enters once.
			queue = append(queue, value)
		}
		// Repeated values can become invalid but never unique again, so stale
		// candidates are permanently removed from the front.
		for len(queue) > 0 && counts[queue[0]] != 1 {
			queue = queue[1:]
		}
		// Report the earliest still-unique value, or -1 when none remains.
		if len(queue) == 0 {
			reports = append(reports, -1)
		} else {
			reports = append(reports, int(queue[0]))
		}
	}
	return reports
}
```

The `-1` result is safe because a byte value cannot equal `-1` after conversion
to `int`.

#### Work and memory

Each of `n` arrivals is enqueued at most once and removed at most once, so total
time is `O(n)` and amortized time per arrival is `O(1)`. The count map and queue
use `O(u)` and `O(u)` memory for up to `u` distinct candidates, bounded by
`O(n)`. The report slice uses `O(n)` result memory. A bounded recent window can
instead cap retained state by its window size.

#### When to use it

- An answer is required after every arrival or input cannot be replayed.
- State has a precise meaning for every received prefix.
- Old items can be summarized or safely expired.

#### When not to use it

- A batch sort may be simpler when the full input is available and intermediate
  answers are unnecessary.
- Exact answers may be impossible with bounded memory if old data can affect
  every future result.
- Define lateness rules before processing events that can arrive out of order.

#### Common mistakes and edge cases

> **Warning:** Returning byte `0` to mean "no unique item" is ambiguous when
> the stream can contain the zero byte; use a separate boolean or an out-of-range
> result such as `-1` after conversion.

- One arrival may remove many candidates, but amortized analysis counts each
  candidate removal only once.
- Advancing a slice queue can retain old backing storage in a long-lived stream.
- A fixed recent window must undo the contribution of the item that expires.

#### Related material

- [Queues and deques](#b3-queues-and-deques) owns FIFO and retained-storage
  details.
- [Composing data structures](#b7-composing-data-structures) explains paired
  state such as counts plus order or two heaps.
- [Heap or priority queue](#c1327-heap-or-priority-queue) supports running
  median, top-`k`, and merging sorted streams.
- [Find Median from Data Stream](problems/hard_leetcode_0295_find_median_from_data_stream.go)
  maintains two balanced halves.

### B.9 Choosing a structure

#### The problem it solves

A problem may require lookup, ordering, removal, or connectivity with different
costs. Choosing by a familiar structure name instead of the required operations
can make the important operation too slow or the invariant impossible to keep.

#### The idea in simple terms

Choose a toolbox drawer by the action you repeat most, not by the shape of the
input alone. A **structure-selection guide** compares observable operation and
ordering needs, then links to the entry that owns each structure's behavior.

#### How it works

Write down required operations, their frequency, ordering rule, mutation rules,
and size limits. Then match the strongest need:

- **Exact-key membership or lookup:** Hash map or set.
- **Direct indexed position:** Slice.
- **Reconnect a known node:** Linked list.
- **Most-recent-first work:** Stack.
- **Oldest-first work:** Queue.
- **Remove at both ends:** Deque.
- **Repeated next minimum or maximum:** Heap.
- **Ordered keys under updates:** [Balanced search tree](#b4-trees-and-binary-search-trees).
- **Repeated string-prefix lookup:** Trie.
- **Incremental group connectivity:** Disjoint-set union.

#### Step-by-step examples

Requirement: accept numbers one at a time and report the largest three seen so
far, without retaining every number.

1. **Step: Name output**
   - **Before:** Need top three, not full sorted order
   - **Decision:** Only the smallest retained candidate must be easy to remove
   - **After:** Choose a min-heap capped at 3
   - **Known correct:** Root is the cutoff candidate
2. **Step: Process `8,2,5`**
   - **Before:** Heap has room
   - **Decision:** Insert all three
   - **After:** Heap contains `{2,5,8}`
   - **Known correct:** These are all values so far
3. **Step: Process `9`**
   - **Before:** Root is `2`
   - **Decision:** `9` is better; replace cutoff
   - **After:** Heap contains `{5,8,9}`
   - **Known correct:** Heap is exactly current top three
4. **Step: Process `1`**
   - **Before:** Root is `5`
   - **Decision:** `1` cannot enter top three
   - **After:** Heap unchanged
   - **Known correct:** Discarding `1` is safe

Boundary example: when `k=0`, retain nothing and avoid reading a heap root. If
the requirement changes to report all values in sorted order, a size-`k` heap
no longer owns enough information.

#### Why it works

Selection is justified by an operation contract. In the example, keeping the
best three and exposing their smallest member is sufficient: a new value no
larger than that cutoff cannot belong to the top three, while a larger value
replaces exactly the current weakest candidate. The chosen structure directly
supports that needed comparison and replacement.

#### Go example

This comparison entry deliberately has no second copy of each implementation.
Use the Go examples owned by [arrays and slices](#b1-arrays-slices-and-strings),
[linked lists](#b2-linked-lists), [queues and deques](#b3-queues-and-deques),
[map presence and sets](#d4-distinguish-missing-keys-from-stored-values-map-presence-and-sets),
and [heap or priority queue](#c1327-heap-or-priority-queue).

#### Work and memory

Selection itself is constant planning work once requirements are known, but the
chosen operation costs control the algorithm. Exact-key map lookup is expected
`O(1)`; slice indexing is `O(1)`; linked-list search is `O(n)`; stack, queue,
and deque end operations are typically `O(1)`; balanced-tree and heap updates
are `O(log n)`. A size-`k` top-items heap processes `n` values in
`O(n log k)` time and uses `O(k)` memory.

#### When to use it

- Before implementing a problem whose operations suggest several structures.
- When constraints require comparing time and memory explicitly.
- When a composed design might be needed for conflicting operations.

#### When not to use it

- Do not treat the comparison list as a keyword matcher; prove the chosen
  invariant.
- Do not choose a BST merely because input is a tree, or two pointers merely
  because input is an array.
- Revisit the choice when mutation, duplicates, ordering, or output needs differ
  from the candidate's assumptions.

#### Common mistakes and edge cases

> **Warning:** A heap exposes only its next-priority value. Using it for repeated
> exact lookup of arbitrary keys can require an `O(n)` scan; choose a map or
> compose the structures when both operations matter.

- Include preprocessing and output construction when comparing total cost.
- Go has a standard heap interface but no standard balanced ordered map.
- Duplicate handling differs across sets, collections that preserve duplicate
  counts, heaps, and search trees.
- Empty removal and `k=0` require explicit guards.

#### Related material

- [Problem-solving method selection guide](#c1-problem-solving-method-selection-guide)
  continues from storage choice to algorithm choice.
- [Composing data structures](#b7-composing-data-structures) handles operation
  sets that one structure cannot satisfy.
- [Hash map or set](#c134-hash-map-or-set), [stack](#c1313-stack),
  [heap](#c1327-heap-or-priority-queue), [trie](#c1333-trie), and
  [disjoint-set union](#c1325-disjoint-set-union) own their problem-solving
  patterns.

<a id="c-problem-solving-patterns"></a>

## C. Ways to solve problems

This section turns observations about an input into a justified solution plan.
A pattern name is only a label; the deciding evidence is the condition that
makes each skipped possibility safe to ignore.

### C.1 Problem-solving method selection guide

Before choosing a method, answer five questions in plain language:

1. What shape does the input have: sequence, range, graph, tree, stream, or
   choices?
2. What must the output contain: a value, position, path, order, one answer, or
   every answer?
3. Must order, adjacency, or original positions be preserved?
4. May the input be changed, and does all of it exist before processing starts?
5. What time and memory can the largest valid input afford?

Treat the links below as hypotheses. Confirm a card's required condition and
correctness argument before implementing it.

<a id="c11-quick-decision-tree"></a>

#### C.1.1 Ask what the input and output allow (quick decision tree)

##### The problem it solves

Many problems resemble several familiar methods. This guide narrows the
choices without assuming that one keyword, such as "minimum," proves a method.

##### The idea in simple terms

Classify the input, result, mutation rules, and limits before naming an
algorithm. A **decision tree** is a sequence of questions whose answers remove
methods that cannot meet those facts.

##### How it works

1. Write the required output and whether order or original positions matter.
2. State whether the input may be reordered or changed.
3. Identify the relationship that drives the work.
4. Reject methods whose worst-case work or memory exceeds the constraints.
5. Open at most two matching cards and verify their required invariant.

Use these observable clues:

- **Connections:** dependency order suggests [topological sort](#c1322-topological-sort);
  equal-cost routes suggest
  [breadth-first search (BFS)](#c1321-breadth-first-search); non-negative
  varying costs suggest [Dijkstra](#c1323-dijkstra-shortest-paths); negative
  costs suggest [Bellman-Ford](#c1324-bellman-ford-shortest-paths); added
  connections suggest [disjoint-set union](#c1325-disjoint-set-union); cheapest
  total network suggests a [minimum spanning tree](#c1326-minimum-spanning-tree);
  and complete branch exploration suggests
  [depth-first search (DFS)](#c1320-depth-first-search).
- **Choices:** a provably permanent choice suggests [greedy](#c1318-greedy);
  repeated smaller questions suggest [dynamic programming](#c1310-dynamic-programming);
  choices that must be tried and undone suggest [backtracking](#c1319-backtracking);
  independent halves suggest [divide and conquer](#c1311-divide-and-conquer).
- **Lookup:** exact presence, counts, partners, or groups suggest a
  [hash map or set](#c134-hash-map-or-set); shared word beginnings suggest a
  [trie](#c1333-trie); several coupled operations may require
  [composed structures](#b7-composing-data-structures).
- **Ranges:** overlap suggests [intervals](#c1330-intervals); active event
  counts suggest a [sweep line](#c1331-sweep-line); fixed-data range totals
  suggest [prefix sums](#c1315-prefix-sums); collected range additions suggest
  [difference arrays](#c1316-difference-arrays); interleaved updates and queries
  suggest [Fenwick or segment trees](#c1317-fenwick-and-segment-trees).
- **Sequences:** aligned beginnings suggest a
  [common-prefix scan](#c131-common-prefix-scan); neighbor-dependent meaning a
  [directional scan](#c132-directional-scan); prescribed transitions
  [simulation](#c133-simulation); reorderable relationships
  [sorting or selection](#c135-sorting-and-selection); a changing yes/no
  boundary [binary search](#c136-binary-search); safe coordinated movement
  [two pointers](#c139-two-pointers); one consecutive range a
  [sliding window](#c1314-sliding-window); exact repeated text
  [string matching](#c137-string-matching); phased input
  [parsing and state machines](#c138-parsing-and-state-machines); a best
  consecutive sum [Kadane's algorithm](#c1312-kadanes-algorithm-best-subarray-scan);
  and linked cycles, middles, or gaps
  [fast and slow pointers](#c1332-fast-and-slow-pointers).
- **Waiting order:** newest first suggests a [stack](#c1313-stack); smallest,
  largest, or most urgent first a
  [heap](#c1327-heap-or-priority-queue); nearest greater or smaller a
  [monotonic stack](#c1328-monotonic-stack); and a moving-window extreme a
  [monotonic deque](#c1329-monotonic-deque).

##### Step-by-step examples

For sorted values `[1, 2, 4, 7, 11]`, target `9`, and a request for any pair of
values, the expected output may be `[2, 7]`.

1. **Check: Output.** **Before:** no method is chosen. **Decision:** only values,
   not original positions, are required. **After:** reordering would be allowed,
   though the input is already sorted. **Known correct:** index preservation is
   not a constraint.
2. **Check: Shape.** **Before:** the input is a sorted sequence. **Decision:** a
   sum that is too small can be increased by moving the left position; a sum
   that is too large can be decreased by moving the right one. **After:** two
   pointers is a candidate. **Known correct:** sorted order can justify each
   discard.
3. **Check: Limits.** **Before:** a pair-by-pair scan would take quadratic work.
   **Decision:** for `n=100,000`, reject that scan and retain the linear
   two-pointer method. **After:** open its card and verify the invariant.
   **Known correct:** the candidate fits the input limit.

For unsorted `[2, 7, 11, 15]` when original indices are required, sorting would
hide the requested positions unless pairs are copied. The lookup branch instead
suggests a [hash map](#c134-hash-map-or-set), which preserves scan positions.

##### Why it works

The guide never proves an algorithm merely by naming it. It preserves the fact
that every surviving candidate still matches the stated input, output,
mutation, and resource conditions. The chosen card must then supply the
algorithm-specific proof. Each answered question removes at least one unknown,
so the selection process finishes; if no card's assumptions hold, returning to
the direct method is the correct outcome.

##### Go example

Selection produces a plan rather than reusable code. After the example above,
the complete Go implementation belongs to
[E.9 Sorted two-pointer pair search](#e9-sorted-two-pointer-pair-search); an
index-preserving version would instead use the map behavior in
[D.4 Map presence and sets](#d4-map-presence-and-sets).

##### Work and memory

Answering the questions is constant-sized planning work. It does not change
the selected algorithm's cost. Count preprocessing such as sorting, retained
maps or queues, recursion, and output separately when comparing candidates.

##### When to use it

Use the guide before implementation when two or more methods seem plausible or
when a familiar keyword is the only reason for a choice.

##### When not to use it

Do not force a listed method when its required ordering, cost, or state rule is
absent. A direct solution or a method outside this guide may be the right fit.

##### Common mistakes and edge cases

> **Warning:** "Minimum" does not prove greedy. Coin values `[1, 3, 4]` and
> target `6` defeat "always take the largest coin": `4+1+1` loses to `3+3`.

Empty input, streaming input, immutable input, duplicate values, negative
values, and output that may itself be exponential can each change which branch
is feasible.

##### Related material

- [C.1.2 Derive a direction before naming a pattern](#c12-derive-a-direction-before-naming-a-pattern)
  starts from a direct solution when no branch is obvious.
- [C.1.4 Choosing among overlapping methods](#c14-choosing-among-overlapping-methods)
  compares methods that survive the first questions.
- [A.1 Complexity and resource accounting](#a1-complexity-and-resource-accounting)
  explains how to reject a method from the constraints.

#### C.1.2 Derive a direction before naming a pattern

##### The problem it solves

When no method is obvious, guessing names encourages an unjustified solution.
The useful starting point is a correct direct process and the work it repeats.

##### The idea in simple terms

Solve a tiny input by hand, describe the simplest correct algorithm, and mark
repeated calculations or discarded information. Those observations give a
direction before a formal pattern name does.

##### How it works

1. State the direct method and why it covers every answer.
2. Trace it on a tiny input.
3. Circle work performed more than once or state rebuilt from scratch.
4. Ask whether that work can be updated, saved, ordered, or safely skipped.
5. Select at most two cards and test their assumptions on the same input.

During timed practice: write the direct method in the first five minutes, use
the guide at five minutes, inspect a relevant card at ten to fifteen minutes,
trace before coding at twenty minutes, and only then read a full explanation if
still blocked.

##### Step-by-step examples

For `[2, 4, 1, 3]`, suppose many queries ask for the sum from `left` through
`right`; one query asks for indices `1..3`, whose expected result is `8`.

1. **Direct attempt.** **Before:** no saved totals exist. **Decision:** add
   `4+1+3`. **After:** return `8`. **Known correct:** every value in the range
   was counted once.
2. **Repeated work.** **Before:** a later query asks for `0..2`. **Decision:**
   notice that `2+4+1` repeats additions used by the first query. **After:** the
   repeated prefix is identified. **Known correct:** the direct method works but
   rebuilds overlapping totals.
3. **Direction.** **Before:** totals always begin at index `0`. **Decision:**
   save `[0,2,6,7,10]` and recover `1..3` as `10-2`. **After:** prefix sums is a
   justified candidate. **Known correct:** subtraction removes exactly the
   values before `left`.

For one query only, preprocessing may cost as much as the direct scan. The
observation then does not justify extra storage.

##### Why it works

The direct method is the coverage baseline: every proposed optimization must
produce the same answer while removing only proven repetition. The fact that
remains true is that each transformation preserves the direct method's set of
considered answers. The process ends after a bounded list of observations; if
none yields a safe improvement, retain the direct solution.

##### Go example

The derived prefix representation and its Go boundary conventions are owned by
[C.1.3.15 Prefix sums](#c1315-prefix-sums). This selection procedure itself
does not add runtime code.

##### Work and memory

Tracing a small example is constant-sized planning work. The selected method's
analysis must include any setup introduced to avoid repetition; in the example,
building prefixes costs linear time and memory before constant-time queries.

##### When to use it

Use it when pattern recall stalls, when a proposed optimization has no proof,
or when the direct solution is clear but too slow.

##### When not to use it

Do not optimize a direct method that already meets the limits unless a simpler
representation improves correctness. Extra machinery can add failure modes
without useful savings.

##### Common mistakes and edge cases

> **Warning:** Repeated syntax is not necessarily repeated work. Two loops may
> process disjoint items, while one nested loop may repeatedly rebuild the same
> range sum.

Check zero queries, one query, empty input, output-size lower bounds, and
whether saved state becomes invalid after updates.

##### Related material

- [C.1.1 Ask what the input and output allow](#c11-ask-what-the-input-and-output-allow-quick-decision-tree)
  narrows the resulting candidates.
- [C.1.5 Let constraints eliminate bad choices](#c15-let-constraints-eliminate-bad-choices)
  checks whether the direct method is already affordable.
- [A.2 Correctness and invariants](#a2-correctness-and-invariants) explains how
  to prove that removing repeated work preserves coverage.

#### C.1.3 Pattern cards

Each card is the complete conceptual reference for its method: what to picture,
when it fits, how to begin, why it is safe, and when to reject it. The cards are
deliberately method-level rather than problem-specific, so they help produce a
first attempt without revealing a judge solution.

Jump directly to a card:

- Sequence and range methods: [common-prefix scan](#c131-common-prefix-scan),
  [directional scan](#c132-directional-scan), [simulation](#c133-simulation),
  [two pointers](#c139-two-pointers),
  [sorting and selection](#c135-sorting-and-selection),
  [binary search](#c136-binary-search), [string matching](#c137-string-matching),
  [parsing and state machines](#c138-parsing-and-state-machines),
  [sliding window](#c1314-sliding-window),
  [prefix sums](#c1315-prefix-sums), [difference arrays](#c1316-difference-arrays),
  [Fenwick and segment trees](#c1317-fenwick-and-segment-trees),
  [fast and slow pointers](#c1332-fast-and-slow-pointers),
  [Kadane's algorithm](#c1312-kadanes-algorithm-best-subarray-scan),
  [monotonic stack](#c1328-monotonic-stack), [monotonic deque](#c1329-monotonic-deque),
  [intervals](#c1330-intervals), and [sweep line](#c1331-sweep-line).
- Lookup and processing-order methods: [hash map or set](#c134-hash-map-or-set),
  [trie](#c1333-trie), [stack](#c1313-stack), and
  [heap or priority queue](#c1327-heap-or-priority-queue).
- Choice methods: [dynamic programming](#c1310-dynamic-programming),
  [greedy](#c1318-greedy), [backtracking](#c1319-backtracking), and
  [divide and conquer](#c1311-divide-and-conquer).
- Relationship methods: [depth-first search](#c1320-depth-first-search),
  [breadth-first search](#c1321-breadth-first-search),
  [topological sort](#c1322-topological-sort),
  [Dijkstra](#c1323-dijkstra-shortest-paths),
  [Bellman-Ford](#c1324-bellman-ford-shortest-paths),
  [disjoint-set union](#c1325-disjoint-set-union), and
  [minimum spanning tree](#c1326-minimum-spanning-tree).

<a id="common-prefix-scan"></a>
<a id="c131-common-prefix-scan"></a>

##### C.1.3.1 Compare from the left (common-prefix scan)

###### The problem it solves

Given several sequences, find the longest beginning that every sequence shares.
The result must start at position `0` and keep the original order; it cannot
skip a position or begin later.

###### The idea in simple terms

Line up the sequences at their left edges and compare one vertical column at a
time. Stop at the first column where a sequence ends or a value differs. The
shared beginning before that point is a **common prefix**; checking it from left
to right is a **common-prefix scan**.

###### How it works

For strings treated as bytes:

1. If there are no strings, return the empty string.
2. Use the first string as the candidate. For each `column` in that string,
   remember its byte as `expected`.
3. For every later `row`, first check that the string reaches `column`, then
   compare its byte with `expected`.
4. On an ended row or mismatch, return `strs[0][:column]`. If every checked
   byte matches, advance to the next column.
5. If the first string ends, return all of it.

The input condition is important: a position means a byte here. Text that must
compare Unicode code points or user-visible characters needs a different
representation before this procedure is used.

###### Step-by-step examples

For `strs = ["flower", "flow", "flight"]`, the expected output is `"fl"`.

1. **Check: Start**
   - **Before:** `column=0`, proven prefix `""`
   - **Decision:** No columns have been checked.
   - **After:** `expected='f'` from row `0`
   - **Known correct:** Every string shares the empty prefix.
2. **Check: Row `1`, column `0`**
   - **Before:** `expected='f'`
   - **Decision:** `"flow"[0]` is `'f'`, so continue.
   - **After:** Row advances to `2`.
   - **Known correct:** Rows `0..1` match at column `0`.
3. **Check: Row `2`, column `0`**
   - **Before:** `expected='f'`
   - **Decision:** `"flight"[0]` is `'f'`; the whole column matches.
   - **After:** `column=1`
   - **Known correct:** Every string shares `"f"`.
4. **Check: Row `1`, column `1`**
   - **Before:** `expected='l'`
   - **Decision:** `"flow"[1]` is `'l'`, so continue.
   - **After:** Row advances to `2`.
   - **Known correct:** Rows `0..1` match at column `1`.
5. **Check: Row `2`, column `1`**
   - **Before:** `expected='l'`
   - **Decision:** `"flight"[1]` is `'l'`; the whole column matches.
   - **After:** `column=2`
   - **Known correct:** Every string shares `"fl"`.
6. **Check: Row `1`, column `2`**
   - **Before:** `expected='o'`
   - **Decision:** `"flow"[2]` is `'o'`, so continue.
   - **After:** Row advances to `2`.
   - **Known correct:** Rows `0..1` match at column `2`.
7. **Check: Row `2`, column `2`**
   - **Before:** `expected='o'`
   - **Decision:** `"flight"[2]` is `'i'`, a mismatch.
   - **After:** Return `strs[0][:2]`
   - **Known correct:** No shared prefix can include column `2`; `"fl"` is longest.

For the boundary input `["dog", ""]`, row `1` ends before column `0`.
The scan returns `strs[0][:0]`, or `""`, without reading past the empty string.
For no input strings, it also returns `""` before choosing a candidate.

###### Why it works

Before each `column`, every string shares exactly the already-proven candidate
`strs[0][:column]`. This fact is true initially because every string has the
empty prefix. If every row reaches the column and matches `expected`, appending
that byte preserves the fact. If a row ends or differs, a prefix containing
that column cannot be shared, and a prefix cannot skip the failed column, so
the already-proven prefix is the longest possible answer.

Each successful outer step advances `column`, and `column` is bounded by the
first string's length, so the scan finishes. If no check fails, every byte in
the first string is shared and no longer prefix can exist because the first
string itself ends there.

###### Go example

The complete implementation is owned by
[E.7 Common-prefix scan](#e7-common-prefix-scan). Its variables `column`,
`row`, and `expected` match the trace above, and its contract states that it
compares byte prefixes and returns `""` for empty input.

###### Work and memory

Let `S` be the total number of input bytes. A mismatch in the first comparison
does constant work, `O(1)`. When long prefixes match, each examined byte belongs
to an input string, so the worst-case work is `O(S)`. The exact work stops at
the first failed column. The scan keeps only indexes and one expected byte, so
it uses `O(1)` extra memory. The returned substring is a view of the first
string's bytes rather than a separately built result, so it allocates no output
byte buffer but can keep the first string's storage reachable while the result
lives.

###### When to use it

Use this method when the requested result is explicitly a shared prefix of
strings or sequences aligned at position `0`, and comparison can stop forever
at the first ended or mismatching position.

###### When not to use it

Do not use it when the shared result may start later, skip positions, tolerate
edits, or compare reordered items. Those conditions break the rule that one
failed column permanently ends every possible answer.

###### Common mistakes and edge cases

> **Warning:** Check length before indexing. With `["dog", ""]`, reading
> `strs[1][0]` before testing its length causes a run-time panic.

> **Warning:** Byte positions are not always character positions. For text such
> as `"\u00e9"`, byte-by-byte comparison answers a byte-prefix question, not
> necessarily a user-visible-character question.

Empty input and a single string need defined results; this implementation
returns `""` and the single string, respectively. Identical strings return the
whole first string. Comparing only the first and last input strings is valid
only after an ordering argument, not for arbitrary input order.

###### Related material

- [A.2 Correctness and invariants](#a2-correctness-and-invariants) explains the
  proof structure used by the proven prefix.
- [E.7 Common-prefix scan](#e7-common-prefix-scan) owns the complete Go
  implementation.
- [D.6 Bytes, runes, and visible characters](#d6-bytes-runes-and-visible-characters)
  explains how to choose the unit of text comparison.
- [Longest Common Prefix](problems/solved_easy_leetcode_0014_longest_common_prefix.go)
  applies this method to a judge problem.

<a id="directional-scan"></a>
<a id="c132-directional-scan"></a>

##### C.1.3.2 Process the needed neighbor first (directional scan)

###### The problem it solves

An item's contribution depends on the item immediately to one side, so a scan
must know that neighbor before deciding what to do.

###### The idea in simple terms

Walk from the side that reveals the needed neighbor first. This is a
**directional scan**: the direction is chosen from the dependency, not habit.

###### How it works

State which neighbor controls the rule, start at that edge, process the first
item explicitly, and retain only the neighbor information required by the next
position.

###### Step-by-step examples

For Roman values `[1, 5]` (`IV`), the expected total is `4`.

1. **Position `1`.** **Before:** `total=0`, no right value. **Decision:** add
   the rightmost `5`. **After:** `total=5`, `right=5`. **Known correct:** the
   processed suffix has value `5`.
2. **Position `0`.** **Before:** current `1`, `right=5`. **Decision:** subtract
   `1` because it is smaller than the value to its right. **After:** `total=4`,
   `right=1`. **Known correct:** the whole input is valued correctly.

For one value `[5]`, add it directly; a nonexistent neighbor must not be read.

###### Why it works

Before each step, the processed side is correct and `right` is exactly the
adjacent value required by the rule. Applying the rule extends the correct
suffix by one position. The index moves toward the opposite edge each time, so
the scan finishes.

###### Go example

The complete implementation is owned by
[E.8 Directional scan](#e8-directional-scan), whose right-to-left variables
match this trace.

###### Work and memory

Each of `n` values is visited once: `O(n)` time and `O(1)` extra memory.

###### When to use it

Use it when one consistent adjacent side determines each local decision.

###### When not to use it

Do not use it when the decision needs an arbitrary distant value or information
from both unprocessed sides.

###### Common mistakes and edge cases

> **Warning:** Update the remembered neighbor even after subtraction. Failing
> to do so gives the next position the wrong adjacent value.

Handle empty and one-item input explicitly, and derive the adjacent rule from
the problem rather than assuming every smaller-left value should subtract.

###### Related material

- [E.8 Directional scan](#e8-directional-scan) owns the Go implementation.
- [Roman to Integer](problems/solved_easy_leetcode_0013_roman_to_integer.go)
  is the representative problem.
- [C.1.3.3 Simulation](#c133-simulation) handles rules needing more state.

<a id="c133-simulation"></a>

##### C.1.3.3 Follow the rules as state changes (simulation)

###### The problem it solves

The statement gives an ordered process: movement, rounds, formatting, or board
updates, and asks for the state or output after those rules run.

###### The idea in simple terms

Write down only what the process currently needs, then apply one rule at a
time. This direct execution is called **simulation**.

###### How it works

Name every state field, define one transition, preserve old state when updates
are simultaneous, and choose a progress measure such as unread commands. The
robot example accepts only `move` and `right`; unsupported commands must be
rejected at the input boundary rather than treated as movement.

###### Step-by-step examples

A robot begins at `(0,0)` facing north. For `move, right, move`, expect
position `(1,1)` facing east.

1. **First `move`.** **Before:** `(0,0,north)`. **Decision:** north adds one to
   `y`. **After:** `(0,1,north)`. **Known correct:** one command is simulated.
2. **`right`.** **Before:** `(0,1,north)`. **Decision:** rotate without moving.
   **After:** `(0,1,east)`. **Known correct:** two commands are simulated.
3. **Final `move`.** **Before:** `(0,1,east)`. **Decision:** east adds one to
   `x`. **After:** `(1,1,east)`. **Known correct:** all commands are simulated.

With no commands, the starting state is also the final state.

###### Why it works

Before every transition, the stored fields equal the specified process after
the commands already consumed. Applying exactly the next rule preserves that
fact. Each step consumes a command, so the process terminates.

###### Go example

```go
type robot struct{ x, y, direction int }

func runRobot(commands []string) (robot, bool) {
	// The zero value starts at (0,0) facing direction 0, which the movement
	// tables define as north.
	r := robot{}
	for _, command := range commands {
		switch command {
		case "right":
			// Direction cycles north, east, south, west; modulo wraps west to north.
			r.direction = (r.direction + 1) % 4
		case "move":
			// Matching table indices turn the current direction into one coordinate
			// step without a separate branch for every compass direction.
			dx := [...]int{0, 1, 0, -1}
			dy := [...]int{1, 0, -1, 0}
			r.x += dx[r.direction]
			r.y += dy[r.direction]
		default:
			// Reject the whole command sequence when its language is invalid.
			return robot{}, false
		}
	}
	return r, true
}
```

Direction values `0..3` mean north, east, south, and west. The boolean is
false when a command is neither `move` nor `right`.

###### Work and memory

For `e` commands, the loop uses `O(e)` time and `O(1)` extra memory. A process
that must retain a board, history, or output must also count that storage.

###### When to use it

Use it when the written transitions are affordable and fully determine the
next state.

###### When not to use it

Do not execute billions of individual events when the rules allow a safe jump
to a later state.

###### Common mistakes and edge cases

> **Warning:** A simultaneous update must read the old snapshot. Updating one
> cell early can make a later cell observe the wrong round.

Test empty input, minimum dimensions, phase boundaries, exact fits, a final
partial group, and one step past an exact fit.

###### Related material

- [B.6 Matrix traversal and in-place state](#b6-matrix-traversal-and-in-place-state)
  explains matrix state rules.
- [C.1.3.8 Parsing and state machines](#c138-parsing-and-state-machines)
  handles inputs whose legal next action depends on a mode.

<a id="hash-map-or-set"></a>
<a id="c134-hash-map-or-set"></a>

##### C.1.3.4 Remember facts by key (hash map or set)

###### The problem it solves

Repeated scans ask whether a value appeared, how often it appeared, which
earlier item completes a pair, or which group owns an item.

###### The idea in simple terms

Store a fact under a reproducible key. A **hash map** associates keys with
values; a **set** stores only whether each key exists.

###### How it works

Define the key and stored meaning, decide whether to query before insertion,
and specify whether duplicates keep the first position, last position, count,
or a collection.

###### Step-by-step examples

For `[2,7,11,15]` and target `9`, expect indices `[0,1]`.

1. **Index `0`.** **Before:** `seen={}`. **Decision:** complement `7` is absent;
   store `2 -> 0`. **After:** `seen={2:0}`. **Known correct:** no earlier pair
   ends at index `0`.
2. **Index `1`.** **Before:** value `7`, `seen={2:0}`. **Decision:** complement
   `2` exists at `0`. **After:** return `[0,1]`. **Known correct:** those values
   sum to `9` and use distinct positions.

For `[3,3]`, querying before insertion lets the second `3` find the first while
preventing index `0` from pairing with itself.

###### Why it works

Before index `i`, the map contains exactly the defined facts from earlier
positions. A successful complement lookup therefore proves a valid earlier
partner; an unsuccessful lookup proves no retained earlier partner exists.
The index advances each step, so the scan finishes.

###### Go example

```go
func pairIndices(values []int, target int) (int, int, bool) {
	// seen maps each earlier value to an index where it occurred.
	seen := make(map[int]int, len(values))
	for i, value := range values {
		// If target-value appeared earlier, the two distinct indices form a pair.
		if j, found := seen[target-value]; found {
			return j, i, true
		}
		// Insert only after lookup so a value cannot pair with its own index.
		seen[value] = i
	}
	return 0, 0, false
}
```

###### Work and memory

The scan takes expected `O(n)` time and up to `O(n)` extra memory. Individual
hash operations are not promised worst-case constant time.

###### When to use it

Use it for exact lookup, frequency counting, partner search, grouping, or
presence tests where order between keys is irrelevant.

###### When not to use it

Do not use it for sorted traversal, nearest-key queries, or relationships the
chosen key does not preserve.

###### Common mistakes and edge cases

> **Warning:** Go returns a type's zero value for a missing key. Use
> `value, found := entries[key]` when zero is also a valid stored value.

Maps do not preserve insertion order. Grouping needs a **canonical key**, one
representation shared by logically equal objects. Anagrams can use fixed
character-count arrays for a fixed alphabet. Isomorphic strings need mappings
in both directions so two source characters cannot collapse into one target.

###### Related material

- [D.4 Map presence and sets](#d4-map-presence-and-sets) owns Go map behavior.
- [Two Sum](problems/solved_easy_leetcode_0001_two_sum.go) is the
  representative problem.
- [C.1.3.33 Trie](#c1333-trie) supports shared-prefix lookup.

<a id="c135-sorting-and-selection"></a>

##### C.1.3.5 Put related values together (sorting and selection)

###### The problem it solves

Useful relationships are scattered through an order that does not matter, or
only the item at one rank is needed.

###### The idea in simple terms

**Sorting** places every item in order so related values become neighbors.
**Selection** places only enough items to identify a requested rank such as the
third smallest.

###### How it works

Define a consistent comparison and tie rule, preserve original positions when
needed, then either sort everything or partition around a pivot and continue
only on the side containing the desired rank.

###### Step-by-step examples

To detect a duplicate in `[4,1,4,2]`, expect `true`.

1. **Sort.** **Before:** equal `4`s are separated. **Decision:** order values.
   **After:** `[1,2,4,4]`. **Known correct:** equal values are adjacent.
2. **Scan.** **Before:** positions `0..2` have no equal neighbors. **Decision:**
   compare the final `4,4`. **After:** return `true`. **Known correct:** a
   duplicate exists.

For selection in `[4,1,3,2]`, partitioning around `3` may produce
`[1,2 | 3 | 4]`; `3` is already the third-smallest value. Equal keys require a
defined rank convention and careful partition progress.

###### Why it works

Sorting creates a monotonic relationship: after a scan proves a candidate too
small, early, or disjoint, later ordered items cannot reverse that fact.
Quickselect's pivot reaches its final rank, so the side not containing the
requested rank can be discarded.

###### Go example

```go
func hasDuplicate(values []int) bool {
	// Sort a copy so the caller's order remains unchanged.
	ordered := append([]int(nil), values...)
	sort.Ints(ordered)
	// Sorting makes equal values adjacent, reducing duplicate detection to one
	// comparison per neighboring pair.
	for i := 1; i < len(ordered); i++ {
		if ordered[i] == ordered[i-1] {
			return true
		}
	}
	return false
}
```

This complete fragment requires `import "sort"` and copies the input before
sorting.

###### Work and memory

Comparison sorting uses `O(n log n)` time; the copy above uses `O(n)` memory.
Quickselect has expected `O(n)` time but `O(n^2)` worst-case time with poor
pivots. Counting over a small key range `k` costs `O(n+k)` time and `O(k)`
memory. Returning an ordered copy also needs `O(n)` output space.

###### When to use it

Use sorting when input order is irrelevant and order makes duplicates,
intervals, pairs, or greedy choices easy to process. Use selection for one rank.

###### When not to use it

Avoid it when original order cannot be carried along, linear time is required,
or a set, counting array, heap, or direct selection is simpler.

###### Common mistakes and edge cases

> **Warning:** `sort.Ints` and `sort.Slice` mutate their slices. Copy when the
> caller's order must survive; use stable sorting when equal records must keep
> their relative order.

A comparator must be consistent and include required tie-breaking. A huge key
range makes counting storage impractical.

###### Related material

- [D.3 Sorting without unexpected mutation](#d3-sorting-without-unexpected-mutation)
  owns Go mutation details.
- [C.1.3.27 Heap or priority queue](#c1327-heap-or-priority-queue) compares
  streaming top-`k` work.
- [C.1.3.30 Intervals](#c1330-intervals) uses sorted starts.

<a id="binary-search"></a>

<a id="c136-binary-search"></a>

##### C.1.3.6 Find a changing boundary (binary search)

###### The problem it solves

Find an exact value or the first or last position where an ordered yes/no condition changes.

###### The idea in simple terms

Test the middle possible answer and discard one half. The
method is safe only when the possibilities form two consecutive groups, such
as all failures followed by all successes.

###### How it works

Define precisely what `left` and `right` include. A range
written as `[left, right)` includes `left` but excludes `right`. For that kind
of range, test `middle = left + (right-left)/2` and keep
the half that can still contain the boundary. Decide which side retains an
equal value based on whether the first or last valid position is required.

###### Step-by-step examples

For `[1,3,3,5]`, find the first value at least `3`; expect index `1`.

1. **Range `[0,4)`.** **Before:** every possible boundary is retained. **Decision:** middle `2` has value `3`, so keep `[0,2)`. **After:** `right=2`. **Known correct:** positions `2..3` cannot be the first valid position.
2. **Range `[0,2)`.** **Before:** positions `0..1` remain. **Decision:** middle `1` is valid, so keep `[0,1)`. **After:** `right=1`. **Known correct:** position `1` is still possible.
3. **Range `[0,1)`.** **Before:** only position `0` is unclassified. **Decision:** value `1` is too small, so set `left=1`. **After:** the range is empty at `1`. **Known correct:** index `1` is the first valid position.

For an empty slice, the same half-open bounds start and finish at `0`, the insertion position.

###### Why it works

Before every test, the first acceptable position lies
in `[left, right)`. Each update discards only positions proved too early or too
late; when the range is empty, `left` is the boundary.

In data such as `[4, 5, 6, 1, 2, 3]`, the whole range
is not sorted, but at least one side of the middle is. Identify that side and
use its outer endpoint to determine whether the target belongs there. When
searching for `2`, middle value `6` shows that `[4, 5, 6]` is the sorted side.
Because `2` is not between `4` and `6`, discard that side and continue in
`[1, 2, 3]`.

Duplicates can make both sides appear equally sorted and remove the evidence
needed to discard half. When equality at both boundaries is allowed, shrinking
a boundary may be necessary and the worst case can fall to `O(n)`. Likewise,
binary search for a peak or local minimum needs the exact strictness rule;
plateaus invalidate proofs that rely on one adjacent comparison selecting a
unique downhill direction.

###### Go example

The complete lower-bound implementation is owned by [E.1 Binary search bounds](#e1-binary-search-bounds); its `left`, `right`, and `middle` variables match this trace.

###### Work and memory

Each test removes about half of the remaining positions,
so a loop-based search takes `O(log n)` time and `O(1)` extra memory.

###### When to use it

Data is sorted, the problem asks for the first or last
position satisfying a condition, or an ordered numeric answer space has a
single boundary where a yes/no test changes. A numeric answer alone is not
enough.

###### When not to use it

Testing a later candidate can change from yes back to no,
or the middle test cannot prove that one entire half is impossible.

###### Common mistakes and edge cases

Keeping the tested middle can prevent progress. A valid
insertion position may equal `len(values)`. Ordinary bounds require fully
sorted input; wrapped input needs the separate sorted-side rule.

###### Related material

- [E.1 Binary search bounds](#e1-binary-search-bounds) owns the Go implementation.
- [Search Insert Position](problems/solved_easy_leetcode_0035_search_insert_position.go) and [Search in Rotated Sorted Array](problems/solved_medium_leetcode_0033_search_in_rotated_sorted_array.go) are representative problems.

<a id="c137-string-matching"></a>

##### C.1.3.7 Reuse a partial match (string matching)

###### The problem it solves

Find the same exact pattern inside a longer text without repeatedly comparing text already understood.

###### The idea in simple terms

Search a text for an exact pattern while remembering how much
of the pattern is still useful after a mismatch.

###### How it works

First decide that positions mean bytes, Unicode code points, or visible
symbols; the example uses bytes. For a direct search, try every possible start
and compare left to right until the pattern matches or one byte differs.

For Knuth-Morris-Pratt (KMP), first build a table for the pattern. At each
pattern position, the table records the length of the longest proper prefix
that is also a suffix of the pattern ending there. While scanning the text,
advance `matched` on equality. On a mismatch, use the table to shorten
`matched` to the next possible reusable prefix without moving the text
position backward. Define the empty-pattern result before the scan.

For a rolling hash, maintain a numeric fingerprint for each pattern-sized text
window. Equal fingerprints select candidates, but compare the actual bytes
before returning because different windows can collide.

###### Step-by-step examples

Search for `abab` in `ababab`; expect index `0`.

1. **Start `0`.** **Before:** `matched=0`. **Decision:** compare `a`, `b`, `a`, and `b`; all agree. **After:** `matched=4`. **Known correct:** the pattern occurs at index `0`.

For pattern `abac` at start `0`, the first three bytes match and the fourth fails. A direct scan restarts at the next start; KMP instead keeps the longest matched suffix that is also a pattern prefix. An empty pattern needs an explicitly defined result.

The Knuth-Morris-Pratt algorithm, usually
called **KMP**, avoids restarting the pattern from zero after every mismatch.
A **prefix** begins at the pattern's first character; a **suffix** ends at its
last character. A **proper prefix** is shorter than the complete string. A
piece that is both a prefix and suffix is sometimes called a **border**.

For pattern `ababac` and text beginning `ababab`, the first five characters
`ababa` match before pattern character `c` disagrees with text character `b`.
The matched text ends with `aba`, which is also the pattern's first three
characters. KMP keeps those three useful matches and continues instead of
restarting at zero. A table prepared from the pattern records how many
characters remain useful after each mismatch.

KMP's invariant is that the first `matched` pattern bytes equal the ending of
the text processed so far. On mismatch, the table shortens `matched` to the
next possible border without moving the text position backward. Each text
position advances once, so a text of length `n` and pattern of length `m` takes
`O(n + m)` time and the table uses `O(m)` memory.

A hash is a compact numeric fingerprint. Maintain the
fingerprint for the current length-`m` text window by removing the outgoing
character's contribution and adding the incoming character's contribution. A
decimal toy version moves from window `123` to `234` with
`(123 - 1*100)*10 + 4 = 234`; a real string hash applies the same remove,
shift, and add steps under its chosen base and modulus. Two different windows
can occasionally have the same fingerprint; this is a **collision**. Equal
hashes therefore identify candidates, not proof. Compare the actual bytes
before reporting a match unless the problem explicitly accepts a collision
risk.

###### Why it works

A direct search checks every legal start, so it cannot miss an occurrence.
For KMP, the first `matched` pattern bytes always equal the suffix of the text
processed so far. Its fallback table keeps the longest shorter suffix that
could still be a prefix, so it discards only alignments already disproved by
the mismatch. Each step either advances the text position or shortens
`matched`; total shortening is linear because `matched` can grow only when the
text advances. The scan therefore finishes in linear time. Reaching the full
pattern length proves an occurrence. A rolling hash remains correct only when
an equal fingerprint is followed by the byte comparison that rules out a
collision.

###### Go example

```go
func firstMatch(text, pattern string) int {
	// strings.Index searches for the first exact byte substring and returns -1
	// when pattern does not occur; an empty pattern matches at index zero.
	return strings.Index(text, pattern)
}
```

This complete one-off byte search requires `import "strings"`. Use KMP when
the exercise or worst-case guarantee requires its saved-border table.

###### Work and memory

The direct method can take `O(nm)` time and `O(1)` extra
memory, such as when many repeated characters match before the final character
fails. KMP takes `O(n + m)` time and `O(m)` memory. A rolling hash can scan in
`O(n)` expected time after `O(m)` setup, but correctness still requires the
collision check described above.

###### When to use it

The same pattern is searched inside a larger string, and a
straightforward restart would compare the same text characters many times. For
a short pattern or one search, that straightforward method may still be best.

###### When not to use it

One short or one-off search is clearer with a direct scan
or a standard library operation, or when the match permits edits, reordered
characters, or another relationship that exact string matching does not model.

###### Common mistakes and edge cases

Define the empty-pattern result. Byte-based matching is
correct for inputs defined as bytes or a one-byte alphabet, but it can split a
Unicode character encoded with several bytes. KMP is valuable when its
worst-case guarantee matters; it is not a requirement to replace a clear
linear-library search in every interview.

###### Related material

- [D.6 Bytes, runes, and visible characters](#d6-bytes-runes-and-visible-characters) explains byte versus text positions.
- [C.1.3.1 Common-prefix scan](#c131-common-prefix-scan) compares aligned starts rather than searching every start.

<a id="c138-parsing-and-state-machines"></a>

##### C.1.3.8 Track legal input modes (parsing and state machines)

###### The problem it solves

Read input whose next legal symbol depends on which parts have already appeared.

###### The idea in simple terms

The parser is always in one of a small number of modes, and
the current input symbol determines the next mode and any action. The mode
summarizes exactly what prefixes are legal so far.

Here, a **grammar** is the set of arrangements the input permits. A **token**
is one meaningful piece, such as a number, name, or operator. **Tokenization**
splits characters into those pieces; **evaluation** applies their meaning.

###### How it works

Separate tokenization from evaluation when that makes the rules
clearer. List the states, the symbols accepted in each state, the transition,
and which ending states are valid. Reject an invalid transition immediately
rather than letting a later fallback reinterpret malformed input.

For integer parsing, a useful sequence is `leading space -> optional sign ->
digits -> stop`. Before `value = value*10 + digit`, check the bound using
`value > (limit-digit)/10`. Apply the sign according to the required range;
the negative limit may have a larger magnitude than the positive limit.

###### Step-by-step examples

Parse `"  -42x"` with the rule `space -> optional sign -> digits -> stop`; expect `-42`.

1. **Spaces.** **Before:** state is `leading space`. **Decision:** consume two spaces. **After:** the next byte is `-`. **Known correct:** the consumed prefix is legal.
2. **Sign.** **Before:** no digits were read. **Decision:** record a negative sign and enter `digits`. **After:** value is `0`. **Known correct:** another sign is no longer legal.
3. **Digits.** **Before:** value `0`. **Decision:** consume `4`, then `2`. **After:** value `42`. **Known correct:** the numeric token so far is `42`.
4. **Stop.** **Before:** next byte is `x`. **Decision:** stop under this grammar and apply the sign. **After:** return `-42`. **Known correct:** no invalid byte was consumed.

For `"4-2"`, the `-` is not legal after digits; reject it or stop there according to the stated grammar.

For arithmetic expressions, a stack handles nested parentheses, while operator
precedence can be handled by operator/value stacks or by recursive-descent
functions: one function handles additions, another multiplications, and
another individual values or parenthesized expressions. This organization is
called **recursive descent** because the functions follow the grammar's nested
levels. For `2+3*4`, the value parser reads `3` and `4`, the multiplication
level produces `12`, and only then does the addition level produce `14`.
Decide whether **unary** `+` and `-`, which apply to one following value, are
operators or part of numeric tokenization. Do not rely on splitting by spaces
when the grammar permits no spaces or arbitrary spacing.

###### Why it works

After consuming a prefix, the state describes every
legal way that prefix may continue and the accumulated values represent exactly
the completed tokens. Each transition consumes input or terminates, so the
parser cannot silently loop on an invalid character.

###### Go example

```go
func parseSignedPrefix(input string) (int64, bool) {
	i, negative := 0, false
	// Skip only the leading ASCII spaces allowed by this parser's contract.
	for i < len(input) && input[i] == ' ' {
		i++
	}
	// Consume at most one optional sign before any digits.
	if i < len(input) && (input[i] == '+' || input[i] == '-') {
		if input[i] == '-' {
			negative = true
		}
		i++
	}
	// A negative int64 may have magnitude MaxInt64+1, so select the allowable
	// unsigned magnitude before accumulating digits.
	limit := uint64(math.MaxInt64)
	if negative {
		limit++
	}
	start := i
	var magnitude uint64
	// Parse the maximal digit prefix and rearrange the limit test so overflow is
	// rejected before multiplying the current magnitude by ten.
	for i < len(input) && input[i] >= '0' && input[i] <= '9' {
		digit := uint64(input[i] - '0')
		if magnitude > (limit-digit)/10 {
			return 0, false
		}
		magnitude = magnitude*10 + digit
		i++
	}
	if i == start {
		// Whitespace and a sign alone do not constitute a parsed integer.
		return 0, false
	}
	if negative {
		// MinInt64 cannot be formed by negating a positive int64, so return the
		// named constant for its one larger unsigned magnitude.
		if magnitude == uint64(math.MaxInt64)+1 {
			return math.MinInt64, true
		}
		return -int64(magnitude), true
	}
	return int64(magnitude), true
}
```

The fragment requires `import "math"`, matches the stated
stop-at-first-nondigit grammar, and rejects a magnitude outside the `int64`
range before multiplying.

###### Work and memory

Reading `n` input characters once takes `O(n)` time. A
simple number parser can use `O(1)` extra memory. Nested expressions may need a
stack whose size is proportional to the nesting depth, and retaining every
token can require `O(n)` memory.

###### When to use it

Input has a grammar or phased format: optional whitespace,
an optional sign, digits, perhaps a decimal or exponent, nested delimiters, or
operators with precedence. A collection of unrelated boolean flags is often an
implicit state machine that should be made explicit.

###### When not to use it

The input has no meaningful grammar or phases, or a
well-tested standard parser already implements exactly the required language.

###### Common mistakes and edge cases

> **Warning:** Without the pre-operation bound check, parsing
> `"9223372036854775808"` overflows `int64` before the function can reject it.

An empty string, spaces with no digits, or a sign with no following digit must
return failure. State whether trailing nondigits stop a valid prefix, as here,
or make the whole input invalid.

###### Related material

- [A.4 Integer limits and bits](#a4-integer-limits-and-bits) owns overflow bounds.
- [C.1.3.13 Stack](#c1313-stack) supports nested delimiters and expression evaluation.

<a id="c139-two-pointers"></a>
<a id="two-pointers"></a>

##### C.1.3.9 Move coordinated positions (two pointers)

###### The problem it solves

Process a sequence without restarting by moving two positions whose next move can safely rule out possibilities.

###### The idea in simple terms

Track two meaningful positions and move them in a coordinated
way so the scan never needs to restart. Common forms move inward from both
ends, merge two sorted inputs, use separate read and write positions, or expand
outward from a possible center.

###### How it works

For the sorted pair-sum form, start one position at each end. If the sum is too
small, move the left position right; sorted order proves that its old value
cannot succeed with any remaining partner. If the sum is too large, move the
right position left by the symmetric argument. Stop on equality or when the
positions meet. Other two-pointer forms must state their own retained region
and safe movement rule rather than copying this one.

###### Step-by-step examples

Find a pair totaling `9` in `[1,2,4,7,11]`; expect `[2,7]`.

1. **Ends `1` and `11`.** **Before:** every pair between the positions remains. **Decision:** sum `12` is too large, so move the right position left. **After:** ends are `1` and `7`. **Known correct:** no pair using `11` can sum to `9` because `1` was the smallest partner.
2. **Ends `1` and `7`.** **Before:** sum is `8`. **Decision:** move the left position right. **After:** ends are `2` and `7`. **Known correct:** no pair using `1` can reach `9` with an available right value.
3. **Ends `2` and `7`.** **Before:** both positions remain possible. **Decision:** sum is `9`. **After:** return the pair. **Known correct:** the requested pair was found.

An empty or one-item input has no pair and must finish before indexing both ends.

###### Why it works

No pair using a position already
passed by `left` or `right` can meet the target. Sorted order makes each discard
safe.

Track the smallest unread value in each input, copy
the smaller one, and advance only its position. For `[1, 4]` and `[2, 3]`, emit
`1`, then `2`, then `3`, and finally the remaining `4`. When merging into
unused space at the end of the first input, work backward from the largest
values so a write cannot overwrite unread input.

Let `read` inspect every value and `write` mark
where the next retained value belongs. The invariant is that the part before
`write` contains exactly the retained values seen so far, in their required
order. Because `write` never passes `read`, writing cannot destroy unread data.
The [non-zero compaction trace](#a2-correctness-and-invariants) shows both
positions changing on actual values.

Every palindrome is centered on either one
character or one gap. Test all character and gap centers, then move left and
right while the characters match. The already-accepted interior must be a
palindrome before a matching outer pair is added. In `abba`, the center gap
first accepts `bb`, then expands to matching outer `a` values.

When at most one deletion is allowed, scan inward until the first mismatch.
Only two possibilities remain: skip the left value or skip the right value,
then verify the rest of each range with an ordinary palindrome scan. This is a
bounded branch, not general backtracking, so the total work remains `O(n)`. In
`abca`, the first mismatch is `b` versus `c`; skipping either one leaves a
palindrome (`aca` or `aba`).

###### Go example

The complete implementation is owned by [E.9 Sorted two-pointer pair search](#e9-sorted-two-pointer-pair-search). Its boundary checks and movements match this trace.

###### Work and memory

Pair search, merging, and in-place compaction take linear
time. They use `O(1)` extra memory unless a merged output is created.
Expanding around all palindrome centers takes between `O(n)` and `O(n^2)` time
and `O(1)` extra memory.

###### When to use it

Look for sorted pair selection, comparisons between two ends,
merging sorted inputs, modifying a sequence in place, or testing symmetric
positions. For every movement, you must explain why the position left behind
cannot be needed again.

###### When not to use it

No pointer movement can be proved to preserve every
remaining answer, or the method would require sorting data whose original
positions or order must remain intact.

###### Common mistakes and edge cases

Sorting may destroy required original positions. The inward
pair-sum form requires sorted values, though negative values are fine. Use the
provided populated lengths when an array contains spare capacity, advance a
write position only after retaining a value, and test both character and gap
centers for palindromes.

###### Related material

- [E.9 Sorted two-pointer pair search](#e9-sorted-two-pointer-pair-search) owns the Go implementation.
- [A.2 Correctness and invariants](#a2-correctness-and-invariants) traces read/write compaction.
- [Palindrome Number](problems/solved_easy_leetcode_0009_palindrome_number.go), [Longest Palindromic Substring](problems/solved_medium_leetcode_0005_longest_palindromic_substring.go), [Remove Duplicates from Sorted Array](problems/solved_easy_leetcode_0026_remove_duplicates_from_sorted_array.go), and [Remove Element](problems/solved_easy_leetcode_0027_remove_element.go) are representative problems.

<a id="dynamic-programming"></a>

<a id="c1310-dynamic-programming"></a>

##### C.1.3.10 Reuse answers to repeated smaller problems (dynamic programming)

###### The problem it solves

Choose among alternatives when different choice sequences repeatedly reach the same smaller question.

###### The idea in simple terms

Dynamic programming (**DP**) starts with a search that tries
choices. If different paths keep reaching the same smaller question, answer
that question once and reuse the result.

###### How it works

First write a rule that builds one answer from smaller answers;
this rule is called a **recurrence**. Then write the smallest answers known
without using that rule; these are the **base cases**. Finally, choose an order
that calculates every smaller answer before something needs it. You can start
from the requested answer and save results during recursive calls, or start
from the smallest answers and build upward with a loop.

###### Step-by-step examples

Count ways to climb `4` stairs using steps of one or two; expect `5`.

1. **Base states.** **Before:** no larger answer exists. **Decision:** set `ways[0]=1` and `ways[1]=1`. **After:** the two smallest states are known. **Known correct:** there is one empty route and one route to stair `1`.
2. **Stairs `2` and `3`.** **Before:** their required smaller states are known. **Decision:** add the previous two answers. **After:** `ways[2]=2`, `ways[3]=3`. **Known correct:** routes are grouped by a final one-step or two-step move.
3. **Stair `4`.** **Before:** `ways[2]=2`, `ways[3]=3`. **Decision:** compute `ways[4]=3+2`. **After:** return `5`. **Known correct:** every route has been counted once by its final move.

For `n=0`, the base value `1` represents the empty way; a different problem contract may define that boundary differently.

###### Why it works

Every non-empty route has one
final move. Removing that move puts the route into exactly one group: routes
whose final move was one step, or routes whose final move was two steps. Those
groups cover every allowed final move and do not overlap.

###### Go example

The rolling Go implementation is owned by [E.2 Rolling dynamic programming](#e2-rolling-dynamic-programming). It retains only the two earlier states used by the recurrence.

###### Work and memory

Without caching, the stair search grows roughly like
`2^n`. Saving the answer for each stair reduces it to `O(n)` time and `O(n)`
memory. When the next answer needs only a fixed number of earlier answers, keep
just those recent answers and reduce extra memory to `O(1)`.

Common state families recur across interview problems:

- **Prefix or position:** `dp[i]` answers for the first `i` items. Process
  smaller prefixes before larger ones.
- **Take or skip:** `dp[i]` is the best result ending at or through item `i`.
  Retain exactly the earlier states each choice can follow.
- **Capacity or amount:** `dp[c]` answers for capacity or total `c`. When each
  item may be used once, scan capacity backward so it is not reused. Scan
  forward only when reuse is allowed.
- **Grid:** `dp[row][col]` answers for a cell or rectangle boundary. Visit
  every cell that can lead into this cell first.
- **Two sequences:** `dp[i][j]` answers for prefixes of lengths `i` and `j`.
  Fill by increasing prefix length or save each pair during recursive calls.
- **Interval:** `dp[left][right]` answers inside one range. Process shorter
  ranges before longer ranges.
- **Finite state machine:** `dp[i][state]` records a small mode such as holding
  stock or transaction count. Every transition must represent one legal
  action.

State compression means retaining only the table
entries future calculations still need. Suppose a grid calculation uses the
cell above and the cell to the left. In one row-sized slice scanned left to
right, `dp[col]` still holds the result from above before it is replaced, while
`dp[col-1]` already holds the new result from the left. For grid path counts,
suppose the saved previous row is `[1, 1]`. On the next row, column `0` remains
`1`; column `1` then combines the old `1` from above with the new `1` from the
left and becomes `2`, leaving `[1, 2]`. Scanning in the opposite direction
would change those meanings. Explain this dependency before reducing
`O(rows * cols)` storage to `O(cols)`; otherwise overwritten values may be
reused or lost incorrectly.

For longest increasing subsequence, let
`tails[L]` store the smallest ending value found for any increasing subsequence
of length `L`, where `L` is that length. For `[3, 5, 2, 4, 1]`, the stored
endings change as follows:

```text
3 -> [3]
5 -> [3, 5]
2 -> [2, 5]
4 -> [2, 4]
1 -> [1, 4]
```

The two stored entries prove that a length-two subsequence exists, but `[1, 4]`
is not itself a subsequence of this input because `1` arrived after `4`.
Reconstructing the actual values requires each chosen value to link to the
earlier chosen value before it. For each new value,
[binary search](#c136-binary-search) finds the first stored ending greater than or
equal to it; replacing that ending produces an `O(n log n)` length calculation
for strictly increasing subsequences. A non-decreasing variant must use a
different equality rule.

Derive complexity from the number of reachable states multiplied by the work
per transition. A two-dimensional table is not automatically `O(n^2)` if each
cell loops over many choices. **Memoization**, meaning saving the result of a
recursive state for reuse, does not help states that are never repeated.

###### When to use it

All three conditions should hold:

1. The full answer can be built from answers to smaller versions of the
   problem.
2. A direct search answers the same smaller question more than once.
3. A small set of facts completely describes each smaller question.

That small set of identifying facts is called the **state**. For climbing
stairs, the current stair number is enough state because the choices available
next do not depend on how you reached that stair. Describe the state in one
sentence before choosing an array or map to store answers.

###### When not to use it

The smaller questions do not repeat, the saved facts leave
out something that affects future choices, or there are still too many
different smaller questions for the input limits.

###### Common mistakes and edge cases

Writing the rule alone does not prevent repeated work; the
answers must also be saved. Do not store facts that cannot affect future
choices. Make the smallest known answers agree with the rule, update recent
values without destroying one that is still needed, count memory used by
unfinished recursive calls, and check whether the count can exceed the chosen
integer type.

###### Related material

- [E.2 Rolling dynamic programming](#e2-rolling-dynamic-programming) owns the Go implementation.
- [C.1.3.10.1 Preserve order with compact ranks](#c13101-coordinate-compression) owns coordinate compression.
- [Climbing Stairs](problems/solved_easy_leetcode_0070_climbing_stairs.go) is the representative problem.

<a id="c1311-divide-and-conquer"></a>

<a id="c13101-coordinate-compression"></a>

##### C.1.3.10.1 Preserve order with compact ranks (coordinate compression)

###### The problem it solves

Values may be negative or enormous, but an algorithm needs a small index for
each distinct value while preserving their comparison order.

###### The idea in simple terms

Sort the distinct values and replace each value by its position in that order.
This replacement is called **coordinate compression**; the new position is the
value's **rank**.

###### How it works

Copy the values, sort the copy, remove duplicates, map each remaining value to
its index, and translate the original sequence without changing its order.

###### Step-by-step examples

For `[100,-5,100,20]`, expect ranks `[2,0,2,1]`.

1. **Sort distinct values.** **Before:** values are in input order. **Decision:**
   sort and deduplicate. **After:** `[-5,20,100]`. **Known correct:** every
   distinct value appears once in increasing order.
2. **Assign ranks.** **Before:** no mapping exists. **Decision:** map `-5 -> 0`,
   `20 -> 1`, and `100 -> 2`. **After:** every value has a compact index.
   **Known correct:** smaller original values have smaller ranks.
3. **Translate.** **Before:** original order is `[100,-5,100,20]`. **Decision:**
   replace each value through the map. **After:** `[2,0,2,1]`. **Known correct:**
   duplicates share a rank and input positions remain in order.

For an empty input, both the distinct list and output are empty. Equal values
must receive one rank rather than separate positions.

###### Why it works

The sorted distinct list contains every value once. If `a < b`, `a` appears
before `b`, so `rank[a] < rank[b]`; equal values use the same map entry. Thus
comparisons are preserved. Translation advances through the finite input, so
it finishes.

###### Go example

```go
func compress(values []int) []int {
	// Sorting a copy groups duplicates and preserves the input order for output.
	ordered := append([]int(nil), values...)
	sort.Ints(ordered)
	// Assign the next rank only on a value's first sorted occurrence.
	rank := make(map[int]int, len(ordered))
	for _, value := range ordered {
		if _, found := rank[value]; !found {
			rank[value] = len(rank)
		}
	}
	// Translate the original sequence through the completed value-to-rank map.
	result := make([]int, len(values))
	for i, value := range values {
		result[i] = rank[value]
	}
	return result
}
```

This complete fragment requires `import "sort"`.

###### Work and memory

Sorting `n` copied values takes `O(n log n)` time; mapping and translation take
expected `O(n)` time. The copy, map, and output use `O(n)` memory.

###### When to use it

Use it when only value order matters and a Fenwick tree, segment tree, or DP
transition needs compact indices for large or negative values.

###### When not to use it

Do not use ranks when actual numeric gaps matter, such as physical distance
`100-20`; ranks preserve order, not magnitude.

###### Common mistakes and edge cases

> **Warning:** Compression does not reorder the input. Processing the sorted
> copy instead of the original sequence breaks order-dependent DP.

Include every query endpoint that must receive a rank, deduplicate equal values,
and do not confuse a value rank with an original input position.

###### Related material

- [C.1.3.10 Dynamic programming](#c1310-dynamic-programming) uses compressed
  ranks to speed value-range transitions.
- [C.1.3.17 Fenwick and segment trees](#c1317-fenwick-and-segment-trees) query
  and update those ranks.
- [C.1.3.5 Sorting and selection](#c135-sorting-and-selection) owns the sort.

##### C.1.3.11 Solve independent parts and combine them (divide and conquer)

###### The problem it solves

Split an input into smaller independent parts, solve each part by the same rule, and combine their results.

###### The idea in simple terms

Split one problem into smaller independent parts, solve each
part by the same rule, and combine their answers.

###### How it works

Define the smallest input that can be answered directly, how
the input is divided, and exactly what information the combine step needs. Then
write a recurrence that includes both recursive work and the combine cost.

###### Step-by-step examples

Sort `[4,1,3,2]`; expect `[1,2,3,4]`.

1. **Split.** **Before:** one four-value problem. **Decision:** divide it into `[4,1]` and `[3,2]`. **After:** two independent halves remain. **Known correct:** together they contain every original value.
2. **Solve halves.** **Before:** each half has two values. **Decision:** split to one-value lists and merge. **After:** halves are `[1,4]` and `[2,3]`. **Known correct:** each half is sorted.
3. **Combine.** **Before:** two sorted halves. **Decision:** repeatedly take the smaller front value. **After:** `[1,2,3,4]`. **Known correct:** the output is sorted and contains every value.

An empty or one-value input is already solved and prevents endless splitting.

For merge sort, each call sorts two halves, then merges them by repeatedly
taking the smaller unread value. The invariant during merging is that the
output prefix is sorted and contains exactly the smallest values consumed from
both halves. The recurrence `T(n) = 2T(n/2) + O(n)` gives `O(n log n)` time.
The reason is concrete: repeatedly halving `n` creates about `log n` levels,
and all merges on one level process `n` values in total. Multiplying the levels
by the work per level gives `O(n log n)`. Ordinary slice-based merging needs
`O(n)` extra storage in addition to `O(log n)` recursive stack depth for
balanced splits.

Exponentiation by squaring uses the same idea without two recursive calls. If
`n` is even, `x^n = (x^(n/2))^2`; if odd, multiply by one extra `x`. Repeatedly
halving the exponent takes `O(log |n|)` multiplications. For example,
`x^10 = (x^5)^2`, and `x^5 = x * (x^2)^2`, so the calculation reuses squared
results instead of multiplying `x` ten times. For negative exponents, use
`x^(-n) = 1/x^n`. Take care when converting `-n` to a positive magnitude:
negating the most negative signed integer exceeds that type's positive limit.

###### Why it works

For merge sort, assume each smaller call returns its half in sorted order.
During the merge, the output prefix is sorted and contains exactly the values
already removed from the two halves. Their smaller front value must be the
smallest value not yet emitted, so appending it preserves that fact. When one
half empties, the other half is already sorted and every one of its values is
at least the output's last value. The base case handles fewer than two values,
and each split makes the ranges smaller, so recursion finishes with a sorted
permutation of the input.

Divide and conquer is most natural when subproblems do not overlap. Dynamic
programming is needed when different branches repeatedly ask for the same
state. Memoizing a divide-and-conquer recursion because its parts overlap has
effectively turned it into DP.

###### Go example

```go
func mergeSort(values []int) []int {
	// A zero- or one-element slice is already sorted; return a copy so every
	// call owns its result independently of the input.
	if len(values) < 2 {
		return append([]int(nil), values...)
	}
	// Recursively sort two disjoint halves.
	middle := len(values) / 2
	left, right := mergeSort(values[:middle]), mergeSort(values[middle:])
	// Repeatedly take the smaller front value. Choosing left on equality keeps
	// equal input values in their original relative order.
	result := make([]int, 0, len(values))
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			// The left front is the smallest remaining value; append and consume it.
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

The copied base case prevents the result from aliasing a caller-owned input.

###### Work and memory

The cost depends on the number and sizes of the smaller
parts plus the work needed to combine them. Write that relationship before
claiming a bound: merge sort is `O(n log n)` time with `O(n)` merge storage,
while exponentiation by squaring is `O(log |n|)` time and can use `O(1)` extra
memory as a loop.

###### When to use it

A balanced split makes the input substantially smaller, the
parts can be solved without sharing evolving state, and their results have a
clear merge step. Merge sort, tree reconstruction, fast exponentiation, and
some geometric algorithms use this shape.

###### When not to use it

Splits can remain badly unbalanced, the parts repeatedly
solve the same smaller questions, or combining their answers costs as much as
solving the original problem directly.

###### Common mistakes and edge cases

> **Warning:** A split that does not make both recursive ranges smaller loops
> forever. If `[left,right)` is split at `middle`, recurse on `[left,middle)`
> and `[middle,right)`, not the original range again.

An empty or one-value input is the base case. During merge, copy the remainder
of either half after the other empties, and preserve a declared tie rule for
equal values when stable order matters.

###### Related material

- [C.1.3.10 Dynamic programming](#c1310-dynamic-programming) handles overlapping rather than independent subproblems.
- [C.1.3.9 Two pointers](#c139-two-pointers) owns sorted merging.
- [A.3 Recursion and saved work](#a3-recursion-and-saved-work) counts call-stack memory.

<a id="c1312-kadanes-algorithm-best-subarray-scan"></a>

##### C.1.3.12 Keep the best section ending here (Kadane's algorithm)

###### The problem it solves

Find the largest sum of one non-empty consecutive section without trying every pair of boundaries.

###### The idea in simple terms

At each position, keep the best consecutive section that must
end here. It either extends the best section ending one position earlier or
starts fresh at the current value.

For value `x`, update:

```text
endingHere = max(x, endingHere + x)
bestOverall = max(bestOverall, endingHere)
```

###### How it works

Require a nonempty input. Initialize `endingHere` and `bestOverall` from the
first value. For every later value, set `endingHere` to the larger of the value
alone and the previous `endingHere` plus that value. Then update
`bestOverall` with the larger of its old value and the new `endingHere`.

###### Step-by-step examples

For `[-2,3,-1,4,-5]`, expect sum `6` from `[3,-1,4]`.

1. **Value `-2`.** **Before:** no section exists. **Decision:** initialize both tracked sums to `-2`. **After:** `endingHere=-2`, `bestOverall=-2`. **Known correct:** both describe the only section seen.
2. **Value `3`.** **Before:** extending gives `1`. **Decision:** start at `3` instead. **After:** both sums are `3`. **Known correct:** `3` is the best section ending here and overall.
3. **Values `-1,4,-5`.** **Before:** best overall is `3`. **Decision:** extend in order, producing ending sums `2,6,1`. **After:** `bestOverall=6`. **Known correct:** every possible ending position was represented.

For `[-4,-2]`, initializing from the first value returns `-2`; initializing to zero would wrongly allow an empty section.

###### Why it works

Every non-empty subarray ending at the current
position is either the current value alone or a subarray ending at the previous
position followed by the current value. Among extensions, only the best
previous ending can matter. `bestOverall` records the best ending among every
position processed so far.

Initialize from the first value when the answer must be non-empty. Initializing
both values to zero incorrectly returns zero for an all-negative input. To
return indices, remember the candidate start whenever beginning at the current
value is better, and save it when the global best improves.

For a circular array, a wrapping maximum equals the total sum minus the minimum
non-wrapping subarray because the wrapping answer keeps everything outside one
consecutive middle section. For `[5, -3, 5]`, the total is `7`; removing the
minimum section `[-3]` gives wrapping sum `7 - (-3) = 10`, representing the two
end values joined around the circle. Compare that result with ordinary Kadane.
When all values are negative, removing the minimum would remove the whole array
and produce an invalid empty answer, so use the ordinary non-empty result.

###### Go example

```go
func maxSubarray(values []int) (int, bool) {
	// An empty slice has no nonempty subarray to score.
	if len(values) == 0 {
		return 0, false
	}
	// endingHere is the best sum ending at the current position; bestOverall is
	// the best sum ending anywhere in the processed prefix.
	endingHere, bestOverall := values[0], values[0]
	for _, value := range values[1:] {
		// Either extend the previous ending range or start fresh at this value.
		if endingHere+value > value {
			endingHere += value
		} else {
			endingHere = value
		}
		if endingHere > bestOverall {
			// A newly computed ending range improves the best answer seen anywhere.
			bestOverall = endingHere
		}
	}
	return bestOverall, true
}
```

The boolean gives empty input an explicit result instead of inventing an empty
subarray sum.

###### Work and memory

The scan takes `O(n)` time and `O(1)` extra memory.

###### When to use it

Find the largest sum of one non-empty consecutive section without trying every pair of boundaries.

###### When not to use it

The answer may skip values, is not additive, or deciding
whether to extend needs more information than the best section ending at the
previous position.

###### Common mistakes and edge cases

> **Warning:** Initializing both sums to zero permits an empty section. For
> `[-4,-2]`, that mistake returns `0` instead of the required nonempty answer
> `-2`.

Reject or separately define empty input, and check addition overflow when an
accepted value range can make `endingHere+value` exceed `int`.

###### Related material

- [C.1.3.14 Sliding window](#c1314-sliding-window) handles ranges whose validity can be repaired by moving boundaries.
- [A.2 Correctness and invariants](#a2-correctness-and-invariants) explains the ending-here invariant.

<a id="c1313-stack"></a>
<a id="stack"></a>

##### C.1.3.13 Finish the newest open item first (stack)

###### The problem it solves

Track nested or unfinished work when the most recently opened item must be completed first.

###### The idea in simple terms

Add and remove only at the top, so the most recently added
item is the first removed. This is last in, first out.

###### How it works

Define what each stored item represents and what makes it
finished. Push newly opened work, inspect or pop only the top, and decide how an
empty stack is reported.

In the Go reference implementation, `Push` adds a value, `Pop` removes the
top, `Peek` inspects it, `Len` reports the count, and `IsEmpty` reports whether
the stack contains no values. `Pop` and `Peek` return a zero value plus `false`
when the stack is empty; the boolean distinguishes that default value from a
successful removal.

###### Step-by-step examples

Validate `([])`; expect `true`.

1. **Open `(` and `[`.** **Before:** stack is empty. **Decision:** push each opener. **After:** stack is `['(', '[']`. **Known correct:** it contains exactly the unmatched openers.
2. **Close `]`.** **Before:** top is `[`. **Decision:** pop the matching opener. **After:** stack is `['(']`. **Known correct:** the inner pair is valid.
3. **Close `)`.** **Before:** top is `(`. **Decision:** pop it. **After:** stack is empty. **Known correct:** every closer matched the newest opener.

For `([)]`, `)` does not match top `[`, so reject immediately. Empty input is valid only if the problem defines an empty sequence as balanced.

###### Why it works

The stack contains exactly the opened items not yet
completed. The top is the most recent unfinished item and must be completed
first.

A stack can also compress adjacent runs. Store `(value, count)` for each run;
merge a new value with the top when equal, and pop when a rule eliminates the
completed run. After a pop, recheck the new top because two previously separated
runs may now touch and combine. For example, if three equal letters disappear,
processing `aabbbaa` first removes the run `bbb`; the two stored `aa` runs then
touch and combine into `aaaa`. Rechecking the top handles that new adjacency.
This avoids repeated full rescans when each item is pushed and popped only a
constant number of times.

###### Go example

The complete generic implementation is owned by [E.5 Generic stack](#e5-generic-stack). A bracket validator may use a byte slice directly and must check length before reading or removing the top.

###### Work and memory

Push, pop, and top inspection take `O(1)` time, aside from
occasional slice growth when Go uses a slice as the stack. The stack may retain
`O(n)` items. Clear a removed slice position when it may hold pointers so Go's
garbage collector can reclaim objects no longer in use.

###### When to use it

The most recently opened or unfinished item must be handled
first, as with nested brackets, undo history, expression evaluation, or a
depth-first search implemented with a loop.

###### When not to use it

The oldest item must be processed first or the algorithm
needs efficient access to arbitrary stored items.

###### Common mistakes and edge cases

Use a queue when the oldest waiting item must be handled
first. Removing from the front of a Go slice does not model a stack and may
retain a large underlying array. Recursive calls already use memory like a
stack behind the scenes; use your own stack when nesting may be unsafe or
direct control is needed.

###### Related material

- [E.5 Generic stack](#e5-generic-stack) owns the Go implementation.
- [C.1.3.20 Depth-first search](#c1320-depth-first-search) can replace recursive calls with an explicit stack.
- [Valid Parentheses](problems/solved_easy_leetcode_0020_valid_parentheses.go) is the representative problem.

<a id="sliding-window"></a>

<a id="c1314-sliding-window"></a>

##### C.1.3.14 Move one consecutive range (sliding window)

###### The problem it solves

Find a best consecutive range when its tracked state can be updated as either boundary moves.

###### The idea in simple terms

Put a movable fence around one consecutive portion of the
input. The fence always contains neighboring items with no gaps.
Extend the right side to consider new values and move the left side when earlier
values must leave.

###### How it works

Define exactly what condition the current range must meet. Move
`right` once per step and add that item to the tracked information. Move `left`
and remove items until the range meets the condition again. Record the answer
at the moment required by the problem.

A common form gives the window a budget of at most `k` violations, such as
zeros that may be flipped or characters that may be replaced. Track the number
of violations or the amount required to repair the window. Shrink while the
budget is exceeded; afterward the window is the longest valid one ending at
the current right position under that invariant.

###### Step-by-step examples

Find the shortest positive-value range totaling at least `7` in `[2,3,1,2,4,3]`; expect length `2`.

1. **Grow.** **Before:** the window is empty. **Decision:** extend through the first `2`, giving `[2,3,1,2]` with sum `8`. **After:** best length is `4`. **Known correct:** it is the first valid window ending there.
2. **Shrink and continue.** **Before:** positive values make removal decrease the sum. **Decision:** remove left values while the sum remains at least `7`, then keep growing. **After:** later window `[4,3]` has length `2`. **Known correct:** every shorter window ending at each right boundary was tested.

Mixed negative values break this shrinking rule because removing a negative value can increase the sum.

###### Why it works

Before each new `right`, every earlier ending position has already been handled.
After adding `values[right]`, positive values make the sum grow. While the sum
meets the target, recording the length and removing the left value tests every
valid start for this `right` from longest to shortest; once the sum falls below
the target, removing more positive values cannot restore validity. Both
boundaries move only forward, so the scan finishes after considering the
shortest valid range ending at every position.

###### Go example

```go
func shortestPositiveWindow(values []int, target int) int {
	// best starts as an impossible length so it also records whether any valid
	// window has been found.
	left, sum, best := 0, 0, len(values)+1
	for right, value := range values {
		// Expanding right can only increase sum because all values are positive.
		sum += value
		// While valid, record the window and remove its leftmost value to find the
		// shortest valid window ending at right.
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

This form requires positive values and a positive target; those conditions make
left-boundary removal safe.

###### Work and memory

Each end moves forward at most `n` times, so the usual
window scan takes `O(n)` time. A running sum uses `O(1)` extra memory; character
or value counts use memory proportional to the number of distinct values kept.

###### When to use it

The prompt asks for a longest, shortest, largest, or smallest
consecutive portion of an array or string. You can update the tracked sum,
character counts, or other information when one item enters or leaves.

###### When not to use it

The answer may skip items, or adding an item can make the
condition switch back and forth unpredictably. Mixed positive and negative
values often cause that problem for sum-based windows.

###### Common mistakes and edge cases

> **Warning:** The inner loop is unsafe for `target <= 0`: with
> `values=[1]`, it remains true after `left` passes the slice and then indexes
> outside it. This implementation requires a positive target.

Empty input returns `0`. A single value at least as large as the target gives
length `1`; mixed negative values require another method because shrinking no
longer changes the sum predictably.

###### Related material

- [C.1.3.15 Prefix sums](#c1315-prefix-sums) handles fixed range queries without a safe moving boundary.
- [C.1.3.29 Monotonic deque](#c1329-monotonic-deque) maintains a window extreme.

<a id="c1315-prefix-sums"></a>

##### C.1.3.15 Subtract saved earlier totals (prefix sums)

###### The problem it solves

Answer many overlapping range-sum or range-count questions without adding the same prefix repeatedly.

###### The idea in simple terms

Store a summary of everything before each position, and when
useful a second summary of everything after it, so later answers do not repeat
the same work.

###### How it works

Let `prefix[i]` mean "the total of the first `i` values." Include
`prefix[0] = 0`. The sum from `left` through `right` is then
`prefix[right+1] - prefix[left]`.

###### Step-by-step examples

For `[2,4,1,3]`, query indices `1..3`; expect `8`.

1. **Build.** **Before:** `prefix=[0]`. **Decision:** append each running total. **After:** `[0,2,6,7,10]`. **Known correct:** `prefix[i]` is the sum of the first `i` values.
2. **Query.** **Before:** `prefix[4]=10` includes indices `0..3`. **Decision:** subtract `prefix[1]=2`. **After:** `10-2=8`. **Known correct:** exactly index `0` was removed.

An empty range needs a defined convention; using a leading zero makes a range beginning at index `0` require no special subtraction case.

###### Why it works

Both stored totals begin at index `0`. Subtracting
the total before `left` removes exactly the values that are outside the
requested range.

The same preprocessing idea works beyond sums when the combination rule
supports the required query. For `Product of Array Except Self`, a left-to-right
pass stores the product before each position, then a right-to-left running
product multiplies in everything after it. Zeros need no division special case,
and the output slice can hold the prefix values while only one suffix product
uses extra memory.

Prefix frequencies can answer how many times a character appears in a range.
Prefix XOR works for range XOR because equal prefixes cancel. A prefix minimum
can answer "smallest value seen so far" but cannot generally be subtracted to
recover an arbitrary range minimum; use a monotonic deque for moving ranges or
a range-query structure for arbitrary updates and queries.

###### Go example

```go
func prefixSums(values []int) []int {
	// prefix has an extra leading zero so prefix[i] equals the sum before index i.
	prefix := make([]int, len(values)+1)
	for i, value := range values {
		// Extend the total through values[i] and store it one position to the right.
		prefix[i+1] = prefix[i] + value
	}
	return prefix
}

func rangeSum(prefix []int, left, right int) int {
	// Subtract the sum before left from the sum through right; everything outside
	// the inclusive range cancels.
	return prefix[right+1] - prefix[left]
}
```

The caller must provide `0 <= left <= right` and `right+1 < len(prefix)`.

###### Work and memory

Building summaries for `n` values takes `O(n)` time and
usually `O(n)` memory. Each sum, count, or XOR range query then takes `O(1)`
time when its answer can be recovered from two stored prefixes.

###### When to use it

Many questions ask for sums or counts over overlapping ranges.
The simplest solution repeatedly adds values it has already added for an
earlier question.

###### When not to use it

The requested property cannot be recovered by combining two
running summaries. For example, you cannot find a range's smallest value by
subtracting two previously stored smallest values.

###### Common mistakes and edge cases

> **Warning:** Omitting `prefix[0]=0` shifts every boundary. For
> `values=[2,4]`, the range `0..0` must use `prefix[1]-prefix[0]=2`, not read a
> position before the slice.

Reject `left > right`, negative indexes, and `right+1 >= len(prefix)`. Use a
wider sum type or checked addition when valid inputs can overflow `int`.

###### Related material

- [C.1.3.16 Difference arrays](#c1316-difference-arrays) reverses this idea for collected range updates.
- [C.1.3.17 Fenwick and segment trees](#c1317-fenwick-and-segment-trees) supports updates between queries.

<a id="c1316-difference-arrays"></a>

##### C.1.3.16 Mark where range changes start and stop (difference arrays)

###### The problem it solves

Apply many additions to ranges and materialize all final values only after every update is known.

###### The idea in simple terms

When many updates affect whole ranges, record where each
change starts and stops instead of modifying every covered position
immediately.

###### How it works

Create a bookkeeping array with one more position than the
result array. For an update that adds `value` from `left` through `right`, add
`value` at `left` and subtract `value` at `right + 1`. After recording every
update, scan the positions of the result array from left to right while keeping
a running total of the changes.

###### Step-by-step examples

Start with four zeroes and add `3` to indices `1..3`; expect `[0,3,3,3]`.

1. **Mark boundaries.** **Before:** `diff=[0,0,0,0,0]`. **Decision:** add `3` at `1` and subtract `3` at `4`. **After:** `diff=[0,3,0,0,-3]`. **Known correct:** the update's lifetime is encoded.
2. **Rebuild.** **Before:** running change is `0`. **Decision:** accumulate positions `0..3`. **After:** `[0,3,3,3]`. **Known correct:** every covered position receives the update once.

The extra sentinel position permits an update ending at the last result index without an out-of-range write.

###### Why it works

The running change increases when an update begins and
decreases immediately after that update ends. Every position between those
points receives the update exactly once.

###### Go example

```go
type rangeAdd struct{ left, right, delta int }

func applyRangeAdds(length int, updates []rangeAdd) []int {
	// difference records where each delta starts and where it stops affecting
	// later positions. The extra slot safely holds a stop at index length.
	difference := make([]int, length+1)
	for _, update := range updates {
		difference[update.left] += update.delta
		difference[update.right+1] -= update.delta
	}
	// A prefix sum reconstructs the total active delta at every output position.
	result, running := make([]int, length), 0
	for i := range result {
		running += difference[i]
		result[i] = running
	}
	return result
}
```

Every update must satisfy `0 <= left <= right < length`.

###### Work and memory

Recording `u` updates and rebuilding `n` final values
takes `O(u + n)` time and `O(n)` extra memory.

###### When to use it

The problem applies many additions or subtractions to ranges
and asks for the final array after all updates.

###### When not to use it

You must answer detailed questions after every individual
update rather than after the updates have been collected.

###### Common mistakes and edge cases

> **Warning:** Allocating only `length` difference positions makes an update
> ending at `length-1` write `difference[length]` out of bounds. Reserve the
> extra sentinel position.

Require `length >= 0` and validate every update before writing either marker.
An empty result accepts no nonempty update, and accumulated deltas must fit the
chosen integer type.

###### Related material

- [C.1.3.15 Prefix sums](#c1315-prefix-sums) owns the accumulating rebuild.
- [C.1.3.17 Fenwick and segment trees](#c1317-fenwick-and-segment-trees) handles interleaved updates and queries.

<a id="c1317-fenwick-and-segment-trees"></a>

##### C.1.3.17 Repair only affected range summaries (Fenwick and segment trees)

###### The problem it solves

Alternate point or range updates with range queries without rebuilding every prefix after each change.

###### The idea in simple terms

Imagine marble boxes with trusted totals for useful groups.
A query combines a few group totals; changing one box repairs only groups that
contain it.

###### How it works

A segment tree stores one summary per represented range. A point update changes
the leaf and recomputes only its ancestors; a query splits into disjoint stored
ranges whose summaries combine to the answer. A Fenwick tree stores selected
suffixes of prefixes in a one-based array. Adding `i & -i` visits every stored
range containing a changed position, while subtracting it decomposes a prefix
into disjoint stored ranges. Public zero-based indexes must be validated before
converting them to one-based positions.

###### Step-by-step examples

For `[2,1,5,3]`, query indices `1..2`, then change index `2` to `6`; expect query `6` and total `12`.

1. **Query.** **Before:** tree ranges summarize the original values. **Decision:** combine singleton ranges `[1]=1` and `[2]=5`. **After:** result is `6`. **Known correct:** the requested range is covered once.
2. **Update.** **Before:** index `2` contributes `5`. **Decision:** add `1` to every stored range containing it. **After:** `[2..3]` becomes `9` and `[0..3]` becomes `12`; `[0..1]` stays `3`. **Known correct:** every summary again equals its range.

For no values, queries must be rejected or return the operation's identity according to the API contract.

Detailed segment-tree representation for `[2, 1, 5, 3]`: its single values,
half sums, and whole sum form this hierarchy:

```text
                 [0..3] = 11
                /           \
          [0..1] = 3      [2..3] = 8
           /     \          /     \
       [0] = 2 [1] = 1  [2] = 5 [3] = 3
```

This is a **segment tree**: each node stores one range's sum. Query `1..2` by
combining `[1] = 1` and `[2] = 5`, giving `6`. Changing index `2` from `5` to
`6` repairs only ranges containing it:

```text
[2] = 6 -> [2..3] = 9 -> [0..3] = 12
```

`[0..1] = 3` stays unchanged.

The combine operation must be **associative**: regrouping values cannot change
the result, as `(2 + 1) + 5 = 2 + (1 + 5)`. It also needs an **identity**, a
neutral value for unused ranges. Sum uses `0`; minimum uses positive infinity.

A **Fenwick tree**, or binary indexed tree, stores overlapping prefix sums more
compactly. For the same values and one-based positions:

- **Position `1`:** `1 & -1` is `1`, so it summarizes range `1..1` and
  stores sum `2`.
- **Position `2`:** `2 & -2` is `2`, so it summarizes range `1..2` and
  stores sum `3`.
- **Position `3`:** `3 & -3` is `1`, so it summarizes range `3..3` and
  stores sum `5`.
- **Position `4`:** `4 & -4` is `4`, so it summarizes range `1..4` and
  stores sum `11`.

`i & -i` gives the stored range size. Prefix `1..3` uses entries `3` and `2`:
`5 + 3 = 8`. Changing position `3` by `+1` updates entries `3` and `4`; entries
`1` and `2` stay unchanged. Subtract two prefix sums to get any range sum.

###### Why it works

Every stored summary equals its range after all
updates. Fenwick trees follow low bits; segment trees follow parent and child
ranges. Both touch `O(log n)` summaries.

For range updates, **lazy propagation** delays child updates. Adding `3` to
`[5, 3]` changes that node from `8` to `14` and the root from `11` to `17`, then
stores a pending `+3`. Push it to children only when a later query enters them.
The parent sum must already include every pending update.

Use a prefix array for fixed data, a difference array when updates precede one
final rebuild, and these trees when updates and queries alternate.

[Coordinate compression](#c13101-coordinate-compression) maps large or negative
values to sorted ranks. Include every query endpoint and do not confuse a rank
with an input position.

###### Go example

```go
type fenwick struct{ sums []int }

// Internal indices are one-based. For a positive index i, i & -i isolates its
// lowest 1 bit; that power of two is the size of the range stored at i.
func newFenwick(length int) *fenwick { return &fenwick{sums: make([]int, length+1)} }
func (f *fenwick) add(index, delta int) {
	// Start at the public index's one-based position. Adding i & -i moves to the
	// next larger stored range that also contains the updated value.
	for i := index + 1; i < len(f.sums); i += i & -i {
		f.sums[i] += delta
	}
}
func (f *fenwick) prefix(end int) int {
	total := 0
	// Subtracting i & -i removes the range just counted. The visited disjoint
	// ranges exactly cover the public half-open prefix [0,end).
	for i := end; i > 0; i -= i & -i {
		total += f.sums[i]
	}
	return total
}
// Two prefix sums isolate the inclusive range [left,right].
func (f *fenwick) rangeSum(left, right int) int { return f.prefix(right+1) - f.prefix(left) }
```

The public indices are zero-based; the stored Fenwick positions are one-based.

###### Work and memory

Both use `O(n)` memory. Segment-tree construction is
`O(n)`; compatible queries and updates on either tree are `O(log n)`.

###### When to use it

Alternate point or range updates with range queries without rebuilding every prefix after each change.

###### When not to use it

The data never changes, all updates can be collected before
one final pass, or the number of operations is small enough that a direct scan
is clearer and fast enough.

###### Common mistakes and edge cases

> **Warning:** Calling `add(-1, delta)` converts the public index to stored
> position `0`; because `0 & -0` is zero, the loop never advances. Require
> `0 <= index < length` before calling the compact method.

Likewise require `0 <= left <= right < length` for `rangeSum`. Choose the
operation's correct identity, preserve association order for noncommutative
operations, and account for sum overflow.

###### Related material

- [C.1.3.10.1 Coordinate compression](#c13101-coordinate-compression) maps large values to tree indices.
- [C.1.3.15 Prefix sums](#c1315-prefix-sums) is simpler for immutable data.
- [C.1.3.16 Difference arrays](#c1316-difference-arrays) is simpler when all updates precede one rebuild.

<a id="c1318-greedy"></a>

##### C.1.3.18 Make a provably safe permanent choice (greedy)

###### The problem it solves

Build an optimum result through choices that never need to be undone.

###### The idea in simple terms

Make one choice now and never undo it because you can prove
that it does not reduce the best result still possible.

###### How it works

State the objective and a precise local rule. Imagine an optimal
answer that begins with another choice, replace that choice with yours, and
prove two facts: the result remains valid, and its value is no worse. Then show
that the remaining input is a smaller problem to which the same rule applies.

That replacement proof is called an **exchange argument**. Two related proof
styles are **stays ahead**, which shows the greedy partial answer remains at
least as favorable as any competitor, and the **cut property**, which proves a
choice is safe across the boundary between finished and unfinished work. All
three must account for what the choice gains now and what it leaves possible
later.

###### Step-by-step examples

Choose the most non-overlapping meetings from `(1,9),(2,4),(5,7),(8,10)`; expect three meetings.

1. **Choose `(2,4)`.** **Before:** no meeting is selected. **Decision:** it finishes earliest. **After:** `lastEnd=4`. **Known correct:** replacing any optimal first meeting with this one cannot remove a later option.
2. **Choose `(5,7)`.** **Before:** it starts after `4`. **Decision:** accept it. **After:** `lastEnd=7`. **Known correct:** the same replacement argument applies to the remaining meetings.
3. **Finish.** **Before:** `(1,9)` conflicts and `(8,10)` fits. **Decision:** skip the former and accept the latter. **After:** three meetings are selected. **Known correct:** no optimal schedule needs an earlier rejected finish.

Coins `[1,3,4]` and target `6` defeat "take the largest": `4+1+1` loses to `3+3`.

Full exchange argument: suppose one meeting may start when another ends. To attend
the maximum number of meetings from `(1, 9)`, `(2, 4)`, `(5, 7)`, and `(8, 10)`,
sort by end time, set `lastEnd` to negative infinity, and accept the next
compatible meeting:

1. **Meeting `(2, 4)`:** `2 >= -infinity`, so select it and set `lastEnd` to
   `4`.
2. **Meeting `(5, 7)`:** `5 >= 4`, so select it and set `lastEnd` to `7`.
3. **Meeting `(1, 9)`:** `1 < 7`, so skip it and leave `lastEnd` at `7`.
4. **Meeting `(8, 10)`:** `8 >= 7`, so select it and set `lastEnd` to `10`.

The trace selects three meetings. To prove the rule for every input, let `E` be
the compatible meeting with the earliest end and let `F` be the first meeting
in any optimal remaining schedule. Because `E.end <= F.end`, every meeting
that can follow `F` can also follow `E`. Replacing `F` with `E` preserves both
validity and the number of meetings. After choosing `E`, the remaining meetings
form the same kind of scheduling problem, so the argument repeats.

For a counterexample, coin values `[1, 3, 4]` and target `6` do not support
"take the largest coin." Choosing `4` requires two more coins, while choosing
`3` requires only one. The local choice worsens the future, so it cannot be
made permanent.

Boyer-Moore majority vote keeps a candidate and a
balance. Equal values add to the balance; different values cancel. A value
occupying more than half the input cannot be completely canceled and must be
the final candidate. Without a guaranteed majority, make a second pass to
verify it; the first pass produces only a candidate.

###### Why it works

After every choice, some optimal complete answer is
still reachable, and the remaining work obeys the same rule. A successful
trace on one input is not enough; the exchange, stays-ahead, or cut argument
must justify every permanent discard.

###### Go example

```go
type meeting struct{ start, end int }

func selectMostMeetings(meetings []meeting) []meeting {
	// Sort the supplied slice in place by end time. Finishing earlier leaves at
	// least as much room for every later meeting, but the caller's order changes.
	sort.Slice(meetings, func(i, j int) bool { return meetings[i].end < meetings[j].end })
	selected, lastEnd := make([]meeting, 0, len(meetings)), math.MinInt
	for _, current := range meetings {
		if current.start >= lastEnd {
			// Accept the earliest-finishing compatible meeting and move the boundary
			// to its end; incompatible meetings are skipped.
			selected, lastEnd = append(selected, current), current.end
		}
	}
	return selected
}
```

This complete fragment requires `import "math"` and `import "sort"`; copy the
slice first when mutation is forbidden.

###### Work and memory

Sorting costs `O(n log n)` time, and the scan costs `O(n)`
time. The scan uses `O(1)` extra memory when the meetings are sorted in place,
excluding the space used to return the selected meetings.

###### When to use it

The problem asks for a minimum or maximum built through
permanent choices, and one available choice appears to leave every useful
future option intact. Scheduling, selecting items, and connecting components
often have this form. The optimization goal alone is not evidence that greedy
works.

###### When not to use it

A choice that looks best now can block a better combination
later and no safe replacement proof exists. Use dynamic programming when
different choices lead to repeated future states, or backtracking when the
alternatives must be explored and cannot be summarized that way.

###### Common mistakes and edge cases

There is no universal rule such as "take the smallest" or
"take the largest." The rule and proof depend on the problem. Verify every
extra promise too, such as whether a majority is guaranteed to exist.

###### Related material

- [C.1.3.10 Dynamic programming](#c1310-dynamic-programming) keeps competing future states.
- [C.1.3.19 Backtracking](#c1319-backtracking) explores choices that cannot be made permanent.
- [C.1.3.26 Minimum spanning tree](#c1326-minimum-spanning-tree) uses the cut property for greedy edges.

<a id="c1319-backtracking"></a>

##### C.1.3.19 Try a choice and undo it (backtracking)

###### The problem it solves

Construct one or every valid answer when no next choice can safely eliminate all alternatives.

###### The idea in simple terms

Walk through a tree of decisions while carrying one partial
answer. At each branch, make a choice, explore what it can lead to, and undo the
choice before trying its sibling branch.

###### How it works

Define the partial answer, the facts needed to check whether the
next choice is allowed, and the condition that marks a complete answer. Then:

1. If the partial answer is complete, record a copy or report success.
2. Otherwise, list every legal next choice.
3. Apply one choice to the partial answer and its supporting facts.
4. Recursively explore what can follow that choice.
5. Undo every change made in step 3.
6. Repeat steps 3 through 5 for the other choices.

If only one answer is needed, return as soon as a recursive call succeeds. If
every answer is needed, finish every branch. If a partial answer is already
invalid or cannot improve the current best result, stop that branch immediately
instead of exploring the choices below it. This early stopping is called
**pruning**.

###### Step-by-step examples

Generate all subsets of `[a,b]`; expect `[a,b]`, `[a]`, `[b]`, and `[]`.

1. **Include path.** **Before:** path is `[]`. **Decision:** include `a`, then `b`. **After:** save `[a,b]`. **Known correct:** that branch represents both inclusions.
2. **Undo `b`.** **Before:** path is `[a,b]`. **Decision:** restore `[a]`, exclude `b`, and save `[a]`. **After:** the `a` branch is complete. **Known correct:** both choices for `b` were covered.
3. **Undo `a`.** **Before:** path is `[a]`. **Decision:** restore `[]` and explore `b` included and excluded. **After:** save `[b]` and `[]`. **Known correct:** every two-choice path was visited once.

Saving the path without copying is a boundary error: later undo steps would mutate all saved results.

Undo sequence in full: generate every subset of `[a, b]`. At each position, branch
into including the current value or excluding it:

1. Start with `[]`; include `a`, producing `[a]`.
2. Include `b`, producing the complete subset `[a, b]`; save a copy.
3. Undo the inclusion of `b`, returning to `[a]`; exclude `b` and save `[a]`.
4. Undo the inclusion of `a`, returning to `[]`; exclude `a`.
5. Include `b` and save `[b]`; undo it, exclude it, and save `[]`.

The crucial detail is that undoing `b` restores exactly `[a]`, and undoing `a`
restores exactly `[]`. Each sibling branch therefore begins with only the
choices made on its own path.

###### Why it works

At the start of each recursive call, the partial
answer and its supporting facts describe exactly the choices on the path to
that call. After the call returns, undoing its choice restores exactly what the
caller had before trying that branch.

###### Go example

```go
func subsets[T any](values []T) [][]T {
	// path is the current choice sequence. result preallocates for 2^n subsets;
	// callers must keep n small enough that the shift and allocation are valid.
	result, path := make([][]T, 0, 1<<len(values)), make([]T, 0, len(values))
	var visit func(int)
	visit = func(index int) {
		if index == len(values) {
			// Copy the completed path because backtracking will reuse its storage.
			result = append(result, append([]T(nil), path...))
			return
		}
		// Explore the include branch, undo that choice, then explore exclusion.
		path = append(path, values[index])
		visit(index + 1)
		path = path[:len(path)-1]
		visit(index + 1)
	}
	visit(0)
	return result
}
```

The shift capacity hint is suitable only when `len(values)` is small enough for
enumeration and the integer shift.

###### Work and memory

Generating subsets visits `2^n` complete branches. Saving
a copy of every subset takes `O(n * 2^n)` total time and `O(n * 2^n)` output
space. The current path and unfinished function calls use `O(n)` extra memory,
excluding the output.

###### When to use it

The answer is assembled one decision at a time, and each
partial answer has a clear set of legal next choices. You may need to find one
valid construction, generate every construction, or compare a manageable
number of constructions. Input rules often let you stop exploring a branch as
soon as its partial answer cannot become valid or cannot beat the best answer
already found.

Backtracking is a candidate when all of these are true:

1. A complete answer can be represented as a sequence of choices.
2. From a partial answer, you can enumerate the choices allowed next.
3. No single next choice can be proved safe enough to discard all alternatives.
4. There are few enough possibilities, or invalid branches can be stopped
   early enough, to make exploration practical.

###### When not to use it

- Use **greedy** when one next choice can be proved not to reduce the best
  result still possible, so the other branches may be discarded.
- Use **dynamic programming** when many choice sequences reach the same
  remaining situation and only its best value, count, or yes/no result matters.
  Save that result instead of rebuilding the same search tree.
- Use ordinary **[depth-first search (DFS)](#c1320-depth-first-search)** when
  traversing an existing graph or tree without constructing and undoing a
  candidate answer.
- Avoid exhaustive backtracking when there are too many possibilities and no
  useful early-stopping rule or repeated smaller question exists.

###### Common mistakes and edge cases

> **Warning:** Saving `path` itself instead of a copy makes all subset rows
> share one backing array; later undo operations can turn previously saved
> answers into the final path state. Copy at every completed answer.

Empty input has one completed choice path, the empty subset. Repeated values
need a stated duplicate-output rule, and the `1<<len(values)` capacity hint
must be omitted when the shift or allocation is not safely bounded.

###### Related material

- [C.1.3.20 Depth-first search](#c1320-depth-first-search) describes the exploration order.
- [C.1.3.10 Dynamic programming](#c1310-dynamic-programming) saves repeated states when individual constructions are unnecessary.
- [A.7 Combinatorics and output size](#a7-combinatorics-and-output-size) counts unavoidable output.

<a id="c1320-depth-first-search"></a>

##### C.1.3.20 Finish one route before the next (depth-first search)

###### The problem it solves

Explore every reachable position, complete path, or subtree result without requiring the fewest-move path first.

###### The idea in simple terms

Depth-first search (**DFS**) follows one route as far as
possible. When that route ends, return to the most recent place that had
another route available.

###### How it works

Define one position and how to list every position reachable in
one move. Mark a position reached before exploring from it so a loop cannot send
the search around forever. Use nested function calls or your own stack to
remember where to return.

###### Step-by-step examples

Traverse edges `A-B`, `A-C`, and `B-D` from `A`; one valid DFS order is `A,B,D,C`.

1. **Visit `A`.** **Before:** all vertices are unseen. **Decision:** mark `A`, then follow neighbor `B`. **After:** active route is `A,B`. **Known correct:** `A` cannot be scheduled again through a cycle.
2. **Finish `B`.** **Before:** `B` has unseen child `D`. **Decision:** visit and finish `D`, then return. **After:** `B` is complete. **Known correct:** its reachable branch was fully explored.
3. **Return to `A`.** **Before:** `C` remains unseen. **Decision:** visit `C`. **After:** order is `A,B,D,C`. **Known correct:** every reachable vertex was visited once.

In a directed cycle `A->B->C->A`, an edge to an active vertex proves a cycle; an edge to a finished vertex does not.

A tree-shaped use is calculating each directory's total size: descend through its files
and subdirectories first. Once a subdirectory has been fully explored, return
its total and add it to its parent directory's size.

###### Why it works

When `visit(vertex)` begins, the vertex is marked and will never be scheduled
again. The call examines every outgoing edge; each unseen neighbor is explored
under the same rule before the call returns. By induction over the return path,
returning from a vertex means every vertex reachable through its outgoing
edges has been visited. Marking prevents cycles from creating repeated work,
and the finite graph lets every call eventually finish.

###### Go example

```go
func depthFirst(graph [][]int, start int) []int {
	// seen prevents cycles from causing repeated visits; order records the first
	// time each reachable vertex is entered.
	seen, order := make([]bool, len(graph)), make([]int, 0, len(graph))
	var visit func(int)
	visit = func(vertex int) {
		// Mark before following edges so a back edge cannot re-enter this vertex.
		seen[vertex], order = true, append(order, vertex)
		// Explore every outgoing edge; recurse only when its destination has not
		// already been reached through another path.
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

Validate `start` before calling this fragment; use an explicit stack when a
valid path may be too deep for recursion.

###### Work and memory

With an adjacency list, a complete traversal takes
`O(V + E)` time because it visits each vertex and examines each edge a constant
number of times. Reached state and the explicit or recursive stack use `O(V)`
extra memory in the worst case.

For cycle detection, one permanent `visited` flag is not always enough. In a
directed graph, use three states: unseen, currently on the active DFS path, and
fully finished. An edge to an active vertex is a directed cycle; an edge to a
finished vertex is not. For edges `A -> B`, `B -> C`, and `C -> A`, exploring
from `A` marks `A`, `B`, and `C` active in order; `C -> A` reaches an active
vertex and proves a cycle. In an undirected graph, ignore the edge back to the
current vertex's parent and treat another visited neighbor as a cycle.

###### When to use it

You need to visit an entire connected region, inspect every
complete path, or calculate a tree result from the results below each node. You
do not need the first discovered path to use the fewest moves.

###### When not to use it

You need a route with the fewest equal-cost moves; use BFS
instead. Also avoid recursive DFS when a valid path can be too deep for the
call stack unless the environment guarantees enough stack space.

###### Common mistakes and edge cases

Nested function calls consume memory. A path with too many
steps may exceed the program's call limit, so use your own stack collection for
very deep inputs.

###### Related material

- [B.5 Graphs and grids](#b5-graphs-and-grids) owns graph representation.
- [C.1.3.21 Breadth-first search](#c1321-breadth-first-search) finds fewest equal-cost moves.
- [C.1.3.19 Backtracking](#c1319-backtracking) adds and undoes candidate choices during DFS order.

<a id="c1321-breadth-first-search"></a>

##### C.1.3.21 Visit nearby positions before distant ones (breadth-first search)

###### The problem it solves

Find the nearest position or fewest moves when every allowed move has the same cost.

###### The idea in simple terms

Breadth-first search (**BFS**) explores outward in layers:
first everything one move away, then everything two moves away, and so on.

###### How it works

Use a first-in, first-out queue. Add the starting position and
mark it reached immediately. Repeatedly remove the oldest position, add each
unseen neighbor, and record its distance as one more than the current distance.

###### Step-by-step examples

On path `A-B-C-D`, find the distance from `A` to `D`; expect `3`.

1. **Layer `0`.** **Before:** queue contains `A`. **Decision:** remove `A`, mark and enqueue `B` at distance `1`. **After:** queue is `[B]`. **Known correct:** no unseen vertex is closer than `B`.
2. **Layers `1` and `2`.** **Before:** queue order is by discovery distance. **Decision:** process `B`, then `C`, enqueueing the next unseen vertex. **After:** `D` is enqueued at distance `3`. **Known correct:** all shorter layers were exhausted first.
3. **Reach `D`.** **Before:** no route shorter than `3` remains unprocessed. **Decision:** return `3`. **After:** search ends. **Known correct:** this is the fewest-move route.

With several sources, enqueue all at distance zero before starting. With no source, the queue stays empty and no position is reachable.

A grid use is a maze: process the start at distance `0`, its open neighbors at
distance `1`, and their unseen neighbors at distance `2`. The first time the
exit is added, you have found a route with the fewest moves.

###### Why it works

The queue contains discovered positions in nondecreasing distance order. When
a vertex at distance `d` is removed, every newly discovered neighbor receives
distance `d+1`; any undiscovered route to that neighbor must pass through the
current or a later layer and therefore cannot be shorter. Marking on insertion
makes that first distance permanent. Each vertex is added at most once, so the
finite search terminates.

When the question asks for the nearest source for every
position, add all sources to the queue at distance zero before the search
starts. The layers then grow from every source simultaneously, and the first
time a position is reached is its minimum distance to any source. Running a
separate BFS from every source repeats work and is usually much slower. On the
line `[S, ., ., ., S]`, enqueue both `S` positions at distance `0`; the next
layer gives their neighbors distance `1`, and the center receives distance `2`
from either side.

When both one start and one target are known and moves
can be reversed, grow a search from both ends until they meet. A **frontier** is
the outermost layer reached so far. Expanding the smaller frontier can reduce
the number of states explored dramatically, though the worst-case growth-rate
bound may be unchanged. On path `A-B-C-D-E`, searches from `A` and `E` meet at
`C`; their stored distances `2` and `2` combine to route length `4`. Keep the
two visited-distance maps separate so those distances can be combined
correctly.

###### Go example

```go
func breadthFirstDistances(graph [][]int, start int) []int {
	// -1 means undiscovered, while nonnegative values are finalized edge counts
	// from start in this unweighted graph.
	distance := make([]int, len(graph))
	for i := range distance {
		distance[i] = -1
	}
	queue := []int{start}
	distance[start] = 0
	// FIFO processing visits all distance-d vertices before distance d+1.
	for head := 0; head < len(queue); head++ {
		vertex := queue[head]
		for _, next := range graph[vertex] {
			if distance[next] == -1 {
				// The first discovery is a shortest path: extend the current shortest
				// path by this one edge, then enqueue the vertex once.
				distance[next] = distance[vertex] + 1
				queue = append(queue, next)
			}
		}
	}
	return distance
}
```

Validate `start`; `-1` means unreachable.

###### Work and memory

With an adjacency list, ordinary BFS takes `O(V + E)` time
and `O(V)` extra memory for its queue, distances, and reached state.

###### When to use it

You need the fewest moves through a network where every move
costs the same, the nearest reachable position, or tree nodes grouped by depth.

###### When not to use it

Moves have different costs. The route with the fewest moves
may then cost more than a longer route; use Dijkstra's algorithm when all costs
are zero or greater.

###### Common mistakes and edge cases

> **Warning:** Marking only when a vertex is removed lets two vertices in the
> same layer enqueue it twice. Mark at insertion so the first, shortest
> discovery owns the queue entry.

Validate the start and every neighbor index. Empty graphs have no valid start,
and unreachable vertices retain the declared `-1` distance.

###### Related material

- [B.3 Queues and deques](#b3-queues-and-deques) owns FIFO behavior.
- [C.1.3.20 Depth-first search](#c1320-depth-first-search) explores one route first.
- [C.1.3.23 Dijkstra shortest paths](#c1323-dijkstra-shortest-paths) handles unequal non-negative costs.

<a id="c1322-topological-sort"></a>

##### C.1.3.22 Release tasks after prerequisites finish (topological sort)

###### The problem it solves

Order directed tasks so every prerequisite appears before each task that depends on it.

###### The idea in simple terms

Repeatedly perform tasks whose prerequisites are already
satisfied.

###### How it works

For every prerequisite rule, record which later task it unlocks.
Count each task's unfinished prerequisites. Add every task with a count of zero
to a first-in, first-out queue. When one task finishes, subtract one from the
count of every task that depended on it.

###### Step-by-step examples

For `Shop -> Cook -> Eat`, expect `[Shop,Cook,Eat]`.

1. **Initialize.** **Before:** unfinished prerequisite counts are `0,1,1`. **Decision:** enqueue `Shop`, the only zero-count task. **After:** queue is `[Shop]`. **Known correct:** every queued task is ready.
2. **Release `Cook`.** **Before:** removing `Shop` completes one prerequisite. **Decision:** decrement `Cook` to zero and enqueue it. **After:** output is `[Shop]`. **Known correct:** `Cook` now follows its prerequisite.
3. **Release `Eat`.** **Before:** `Cook` is ready. **Decision:** output it and decrement `Eat`. **After:** output becomes `[Shop,Cook,Eat]`. **Known correct:** every dependency is respected.

Adding `Eat -> Shop` leaves no zero-count task; processing fewer than all tasks proves a cycle blocks any valid order.

###### Why it works

The queue contains exactly the tasks whose prerequisites are all
finished. A task is added to the output only after its remaining prerequisite
count reaches zero.

If the queue becomes empty before every task is processed,
the graph of unprocessed tasks contains a dependency cycle, such as `A`
needing `B` while `B` also needs `A`. Some other unprocessed tasks may not be
part of that cycle; they can simply be waiting on a task that is. No valid
order exists.

An alternative uses DFS with the three states described above. After all
outgoing dependencies of a vertex are finished, append that vertex to a
postorder list; reverse the list for a topological order. The queue-based method
above is often simpler when the problem also needs layer-by-layer processing or
an explicit count of unfinished prerequisites.

###### Go example

```go
func topologicalOrder(graph [][]int) ([]int, bool) {
	// indegree counts each vertex's unmet prerequisites.
	indegree := make([]int, len(graph))
	// Count every incoming edge once, including parallel prerequisites if present.
	for _, edges := range graph {
		for _, to := range edges {
			indegree[to]++
		}
	}
	// Vertices with no prerequisites can appear first.
	queue := make([]int, 0, len(graph))
	for vertex, count := range indegree {
		if count == 0 {
			// This vertex is ready immediately because it has no incoming edge.
			queue = append(queue, vertex)
		}
	}
	// Removing a vertex satisfies each outgoing prerequisite. A neighbor becomes
	// ready exactly when its count reaches zero.
	order := make([]int, 0, len(graph))
	for head := 0; head < len(queue); head++ {
		// Emit the oldest ready vertex without shifting the queue slice.
		vertex := queue[head]
		order = append(order, vertex)
		for _, to := range graph[vertex] {
			// Emitting vertex satisfies one prerequisite of this destination.
			indegree[to]--
			if indegree[to] == 0 {
				// The destination is ready after all its prerequisites reach the output.
				queue = append(queue, to)
			}
		}
	}
	// A cycle leaves at least one positive indegree, so not every vertex is emitted.
	return order, len(order) == len(graph)
}
```

The boolean is false when a dependency cycle prevents a complete order.

###### Work and memory

Both forms take `O(V + E)` time and `O(V)` extra memory
with adjacency lists, not counting the stored graph itself.

###### When to use it

The prompt describes courses, builds, jobs, or other tasks
with rules such as "A must finish before B."

###### When not to use it

The relationships do not represent prerequisites, the
graph may be undirected, or the problem asks for shortest routes rather than a
valid dependency order.

###### Common mistakes and edge cases

> **Warning:** Reversing an edge changes the prerequisite meaning. If `A` must
> finish before `B`, store `A -> B` and increment `indegree[B]`, not
> `indegree[A]`.

Validate every destination index. Duplicate stored edges count as duplicate
prerequisites unless they are deduplicated consistently in both the adjacency
list and indegree counts. An empty graph has the valid empty order.

###### Related material

- [C.1.3.21 Breadth-first search](#c1321-breadth-first-search) owns queue-layer traversal.
- [C.1.3.20 Depth-first search](#c1320-depth-first-search) provides the postorder alternative.
- [B.5 Graphs and grids](#b5-graphs-and-grids) defines directed edges.

<a id="c1323-dijkstra-shortest-paths"></a>

##### C.1.3.23 Finalize the cheapest reachable route next (Dijkstra shortest paths)

###### The problem it solves

Find cheapest routes from one start through edges whose costs are all zero or greater.

###### The idea in simple terms

On a map with travel costs, always process the unvisited
location with the cheapest known route from the starting location.

###### How it works

Record cost `0` for the starting location and "unknown" for all
others. Use a priority queue that removes the route with the smallest cost
first. When traveling through the current location gives a neighbor a cheaper
route, store that new cost and add the route to the queue.

###### Step-by-step examples

Let `A->B=4`, `A->C=1`, and `C->B=2`; from `A`, expect `B=3`.

1. **Expand `A`.** **Before:** `A=0`, others unknown. **Decision:** offer `B=4` and `C=1`. **After:** the heap exposes `C=1`. **Known correct:** these are the cheapest discovered routes.
2. **Expand `C`.** **Before:** `C=1` is cheapest. **Decision:** improve `B` to `1+2=3`. **After:** both `B=3` and stale `B=4` entries may remain queued. **Known correct:** stored `B=3` is the best discovered route.
3. **Finalize `B`.** **Before:** heap removes `B=3`. **Decision:** accept it and later ignore `B=4`. **After:** `B` is final at `3`. **Known correct:** non-negative edges cannot produce a cheaper later route.

A negative edge invalidates that finalization argument even when this particular trace happens to work.

###### Why it works

Suppose the queue removes a route costing
`5`, and `5` is still the cheapest known cost for that location. Every untried
route already costs at least `5` before taking another connection. Because a
connection cannot subtract cost, none can produce a route below `5`.

###### Go example

```go
type weightedEdge struct{ to, cost int }

func dijkstraDense(graph [][]weightedEdge, start int) []int {
	// ^uint(0) has every unsigned bit set. Shifting right clears the top bit, so
	// converting the result to int produces its largest positive value. This
	// infinity sentinel marks vertices not yet reached; done marks final distances.
	const infinity = int(^uint(0) >> 1)
	distance, done := make([]int, len(graph)), make([]bool, len(graph))
	// Begin with every vertex unreachable, then give start its zero-length path.
	for i := range distance {
		distance[i] = infinity
	}
	distance[start] = 0
	for range graph {
		// Among unfinished vertices, choose the smallest tentative distance.
		vertex := -1
		for i := range graph {
			if !done[i] && (vertex < 0 || distance[i] < distance[vertex]) {
				vertex = i
			}
		}
		if vertex < 0 || distance[vertex] == infinity {
			// Every reachable vertex is done; the rest cannot be reached from start.
			break
		}
		done[vertex] = true
		for _, edge := range graph[vertex] {
			// Relax the edge only when addition is safe and the path improves the
			// destination. Correctness requires every edge cost to be nonnegative.
			if distance[vertex] <= infinity-edge.cost && distance[vertex]+edge.cost < distance[edge.to] {
				distance[edge.to] = distance[vertex] + edge.cost
			}
		}
	}
	return distance
}
```

This compact scan is `O(V^2+E)` and requires non-negative costs. A heap changes
selection to the `O((V+E) log V)` form analyzed above.

###### Work and memory

With adjacency lists and a **binary heap**, a priority
queue organized as a tree with at most two children per position, the usual
implementation takes `O((V + E) log V)` time and `O(V + E)` storage including
the graph. Some priority queues support **decrease-key**, which edits the
priority of an entry already stored. When that operation is unavailable, push
a new improved entry and later discard any popped entry whose cost no longer
equals the stored best cost.

###### When to use it

You need the cheapest routes through a network, and every
connection has a cost of zero or greater.

###### When not to use it

Any connection may have a negative cost. If every connection
has the same cost, breadth-first search is simpler.

###### Common mistakes and edge cases

> **Warning:** Negative cost invalidates finalization. On edges `A->B=4`,
> `A->C=7`, and `C->B=-5`, finalizing `B=4` before processing `C` misses the
> cheaper route of cost `2`.

Validate the start and endpoints, reject negative costs, ignore stale heap
entries, and guard route-cost addition before evaluating it.

###### Related material

- [C.1.3.21 Breadth-first search](#c1321-breadth-first-search) is simpler for equal costs.
- [C.1.3.24 Bellman-Ford](#c1324-bellman-ford-shortest-paths) permits negative edges.
- [D.5 Priority queues](#d5-priority-queues) owns Go heap behavior.

<a id="c1324-bellman-ford-shortest-paths"></a>

##### C.1.3.24 Let cheaper routes ripple through every edge (Bellman-Ford)

###### The problem it solves

Find shortest routes when edges may be negative and detect a reachable negative-cost cycle.

###### The idea in simple terms

Let cheaper route costs ripple through every connection,
repeating until an improvement can cross even the longest useful route.

###### How it works

Record cost `0` for the starting location and "unknown" for all
others. Repeat a full pass over every directed connection up to `V - 1` times,
where `V` is the number of locations. For an edge `from -> to` with cost `w`,
replace `distance[to]` with `distance[from] + w` when `from` is reachable and
that total is smaller. This attempt to lower a stored cost is often called
**relaxing an edge**. Stop early if a full pass changes nothing.

###### Step-by-step examples

Let `A->B=4`, `A->C=7`, and `C->B=-5`; from `A`, expect `B=2`.

1. **First edges.** **Before:** only `A=0` is known. **Decision:** set `B=4` and `C=7`. **After:** both destinations are reachable. **Known correct:** these costs use one edge.
2. **Negative edge.** **Before:** `C=7`, `B=4`. **Decision:** lower `B` through `C` to `7-5=2`. **After:** `B=2`. **Known correct:** the two-edge route has propagated.
3. **Extra pass.** **Before:** up to `V-1` passes are complete. **Decision:** scan once more. **After:** no cost falls, so no reachable negative cycle is detected. **Known correct:** every finite simple route was covered.

For `B->C=1` and `C->B=-2`, an extra pass still lowers a cost and proves no finite minimum exists for affected vertices.

###### Why it works

A finite shortest route never needs to
repeat a location; removing any repeated non-negative cycle cannot make it more
expensive, while a repeated negative cycle would mean there is no finite
shortest route. A route without repeated locations uses at most `V - 1` edges.
After pass `k`, every shortest route using at most `k` edges has had a chance to
propagate its cost. Therefore `V - 1` passes cover every finite shortest route.

Make one additional pass. If a reachable
connection can still lower a cost, the route can loop and become cheaper every
time around. There is then no single lowest cost for locations on that loop or
locations reachable from it. For `B -> C = 1` and `C -> B = -2`, one trip
around the loop changes `B`'s cost by `-1`; the extra pass can therefore lower
it again. This is a **negative cycle**.

###### Go example

```go
type directedEdge struct {
	from, to int
	cost     int64
}

func bellmanFord(vertices, start int, edges []directedEdge) ([]int64, bool) {
	// MaxInt64 represents unreachable vertices; start reaches itself at cost zero.
	distance := make([]int64, vertices)
	for i := range distance {
		distance[i] = math.MaxInt64
	}
	distance[start] = 0
	// relax applies the one-edge recurrence and reports whether it improved a
	// reachable destination.
	relax := func(edge directedEdge) bool {
		if distance[edge.from] == math.MaxInt64 {
			// No known path reaches the edge's source, so it cannot form a candidate.
			return false
		}
		// Extend the best known source path by this edge. The function's contract
		// requires every relevant path sum to fit in int64.
		candidate := distance[edge.from] + edge.cost
		if candidate >= distance[edge.to] {
			// The destination already has an equal or cheaper known path.
			return false
		}
		// Record the cheaper path and report a change to the caller.
		distance[edge.to] = candidate
		return true
	}
	// Any simple shortest path has at most vertices-1 edges. Stop early if a full
	// pass makes no improvement.
	for pass := 1; pass < vertices; pass++ {
		changed := false
		for _, edge := range edges {
			// Keep changed true once any edge improves during this pass; evaluating
			// relax first ensures later edges are still processed.
			changed = relax(edge) || changed
		}
		if !changed {
			// With no new distance, another identical pass cannot improve anything.
			break
		}
	}
	// A further improvement proves a reachable negative-cost cycle, because all
	// simple paths have already been considered.
	for _, edge := range edges {
		if relax(edge) {
			return distance, false
		}
	}
	return distance, true
}
```

This complete fragment requires `import "math"` and the input condition that
every evaluated sum `distance[edge.from]+edge.cost` fits `int64`. The boolean
is false when a reachable negative cycle can still lower a cost.

###### Work and memory

If `V` is the number of locations and `E` is the number of
connections, up to `V - 1` full passes take `O(VE)` time. The stored distances
use `O(V)` extra memory, not counting the input map or optional links used to
rebuild the chosen routes.

###### When to use it

You need shortest routes from one starting location, some
connections may have negative costs, and you may need to detect a reachable
negative-cost cycle.

###### When not to use it

Every cost is non-negative and the graph is large enough
that Dijkstra's better running time matters. If a reachable negative cycle may
exist, do not report ordinary finite shortest costs for locations affected by
that cycle.

###### Common mistakes and edge cases

> **Warning:** The sentinel check does not prevent ordinary addition overflow.
> If `distance[from]=math.MinInt64` and `cost=-1`, constructing the candidate
> is invalid; enforce the stated route-sum bound or use checked arithmetic.

Validate `vertices`, `start`, and every endpoint before indexing. An empty
graph has no valid start. Parallel edges are allowed, but every one is scanned;
unreachable negative cycles do not invalidate distances from this start.

###### Related material

- [C.1.3.23 Dijkstra shortest paths](#c1323-dijkstra-shortest-paths) is faster for non-negative edges.
- [B.5 Graphs and grids](#b5-graphs-and-grids) owns graph storage.
- [A.4 Integer limits and bits](#a4-integer-limits-and-bits) covers safe bounds.

<a id="c1325-disjoint-set-union"></a>

##### C.1.3.25 Track connected groups by leader (disjoint-set union)

###### The problem it solves

Add connections and repeatedly ask whether two items now belong to the same connected group.

###### The idea in simple terms

Imagine that every connected group has one shared group
leader. To check whether two items are connected, find each item's leader and
compare them. To connect two different groups, make one leader report to the
other.

###### How it works

Store a `parent` for every item. Initially, each item is its own
parent because every item is in a separate group. `find(item)` follows parent
links until it reaches the group leader. `union(a, b)` finds both leaders and,
if they differ, makes one leader the parent of the other.

###### Step-by-step examples

Connect `1-2`, then `2-3`; expect `1` and `3` to be connected.

1. **Initialize.** **Before:** parents are `[1:1,2:2,3:3]`. **Decision:** each item begins as its own group. **After:** three leaders exist. **Known correct:** no connection has been added.
2. **Union `1,2`.** **Before:** leaders differ. **Decision:** point leader `2` to `1`. **After:** `2->1`. **Known correct:** exactly `1` and `2` share a group.
3. **Union `2,3`.** **Before:** `find(2)` reaches `1`, while `3` leads itself. **Decision:** attach `3` to `1`. **After:** both `find(1)` and `find(3)` return `1`. **Known correct:** all added connections are represented.

Unioning already-connected items changes nothing and can signal that a new undirected edge would form a cycle.

###### Why it works

Every parent link stays inside one represented connected group, and every group
has exactly one leader that points to itself. Initially this is true because
each item is alone. A union changes only one leader's parent and does so only
after an added edge connects the two groups, so it merges exactly those groups
without splitting either one. Therefore two items are connected exactly when
following their parent links reaches the same leader. Path compression changes
only intermediate links to that same leader, and union by size changes only
which leader becomes the parent, so neither optimization changes membership.

While finding a leader, make each visited item point
closer to that leader so later searches take fewer steps. This is called
**path compression**. When joining groups, attach the smaller group below the
larger group so the chains stay short; this is called **union by size**. Both
changes improve speed without changing which items belong together. If parent
links are `4 -> 3 -> 1`, `find(4)` reaches leader `1` and rewrites `4 -> 1`;
the next lookup skips node `3`.

###### Go example

```go
type disjointSet struct{ parent, size []int }

func newDisjointSet(n int) *disjointSet {
	// Initially every element is a one-member tree whose root is itself.
	d := &disjointSet{make([]int, n), make([]int, n)}
	for i := range d.parent {
		// Self-parenting makes i its root; size one records its single member.
		d.parent[i], d.size[i] = i, 1
	}
	return d
}
func (d *disjointSet) find(x int) int {
	for x != d.parent[x] {
		// Point x to its grandparent while walking upward, shortening future finds.
		d.parent[x] = d.parent[d.parent[x]]
		x = d.parent[x]
	}
	return x
}
func (d *disjointSet) union(a, b int) bool {
	// Only roots represent groups; equal roots mean the elements are connected.
	a, b = d.find(a), d.find(b)
	if a == b {
		// They already share a component, so no structural change is needed.
		return false
	}
	if d.size[a] < d.size[b] {
		// Attach the smaller tree below the larger to keep paths shallow.
		a, b = b, a
	}
	d.parent[b] = a
	// Only root a keeps a meaningful component size after the attachment.
	d.size[a] += d.size[b]
	return true
}
```

`union` reports whether two previously separate groups were joined.

###### Work and memory

With both improvements, a long sequence of operations is
almost constant time per operation in practice. Spread across the whole
sequence, which is what **amortized** means here, the formal bound is
`O(alpha(n))`, where `alpha` is a mathematical function called the inverse
Ackermann function. You do not need to calculate it for implementation work:
it grows so slowly that it remains below a tiny constant for any realistic
input. The parent and size arrays use `O(n)` memory.

###### When to use it

Connections are added over time, and queries ask whether two
items are already connected, whether a new connection creates a loop, or how
many separate groups remain.

###### When not to use it

You need the actual route, the order in which items were
visited, or the shortest path. Knowing only each item's group does not preserve
that information.

###### Common mistakes and edge cases

> **Warning:** `find(-1)` or `find(n)` indexes outside the parent slice. Validate
> every item before calling the compact implementation.

An empty set has no valid item. Unioning an item with itself or two already
connected items returns `false` and must not increase the recorded group size.

###### Related material

- [C.1.3.26 Minimum spanning tree](#c1326-minimum-spanning-tree) uses group leaders in Kruskal's algorithm.
- [B.5 Graphs and grids](#b5-graphs-and-grids) explains what connectivity preserves and what it loses.

<a id="c1326-minimum-spanning-tree"></a>

##### C.1.3.26 Connect everything for the least total cost (minimum spanning tree)

###### The problem it solves

Choose undirected weighted edges that connect every vertex with minimum total edge cost and no cycle.

###### The idea in simple terms

Choose enough weighted edges to connect every vertex without
forming a cycle, while minimizing their total cost. The result spans all
vertices and has exactly `V - 1` edges when the graph is connected.

Do not confuse a minimum spanning tree with shortest paths. A spanning tree
minimizes the total cost of the selected network; paths from a particular start
inside that tree need not be the cheapest possible paths in the original graph.

###### How it works

For Kruskal's algorithm, sort all undirected edges by increasing cost. Start
with each vertex in a separate disjoint-set group. Scan the edges in order,
accepting one only when its endpoints have different leaders, and stop after
accepting `V-1` edges. If the scan ends sooner, report that no spanning tree
exists. For Prim's algorithm, instead grow one connected region and repeatedly
remove its cheapest edge to an unseen vertex from a min-heap.

###### Step-by-step examples

For `A-B=1`, `B-C=2`, and `A-C=5`, expect total cost `3`.

1. **Edge `A-B`.** **Before:** three separate components. **Decision:** accept the cheapest edge joining two components. **After:** `{A,B}` and `{C}` remain. **Known correct:** the accepted forest can extend to a minimum tree.
2. **Edge `B-C`.** **Before:** its endpoints are separate. **Decision:** accept cost `2`. **After:** all vertices are connected with two edges. **Known correct:** the tree has `V-1` edges and total `3`.
3. **Edge `A-C`.** **Before:** its endpoints already share a component. **Decision:** reject it. **After:** the tree stays unchanged. **Known correct:** accepting it would create a cycle and add cost.

If fewer than `V-1` edges can be accepted, the graph is disconnected and has only a minimum spanning forest.

Sort edges by weight. Scan from cheapest to most
expensive and accept an edge exactly when its endpoints are currently in
different disjoint-set components. The invariant is that the accepted edges
form a **forest**, meaning a collection of separate trees with no cycles, that
can still be extended to some minimum spanning tree. The **cut property** says
that if the vertices are split into two groups, a cheapest edge crossing
between those groups can be included in some minimum spanning tree. That is
why Kruskal can safely accept the next cheapest edge that joins two current
components. Sorting dominates the work at `O(E log E)`; disjoint-set operations
are nearly constant amortized time with path compression and union by size.

Begin from any vertex and keep the cheapest edge leaving
the connected portion in a min-heap. Repeatedly take the cheapest edge reaching
an unseen vertex. In the same office example, starting at `A` adds `A-B = 1`
and `A-C = 5` to the heap. Taking `A-B` adds `B-C = 2`; taking `B-C` next
connects `C` for the same total cost `3`, so the unused `A-C` entry is skipped.
With adjacency lists and a binary heap, this takes `O(E log V)` time. Ignore
stale heap entries whose destination has already been added.

###### Why it works

Kruskal's accepted edges always form a forest because an edge whose endpoints
already share a component is rejected. Consider the cut separating one current
component from the others. The scanned acceptable edge is a cheapest remaining
edge across such a cut; the cut property says it can replace the crossing edge
in some minimum spanning tree without increasing total cost. Thus after every
acceptance, the forest can still be extended to a minimum spanning tree. Each
accepted edge reduces the number of components, so `V-1` accepted edges form a
minimum-cost spanning tree. If the scan cannot reach that count, the input is
disconnected.

###### Go example

```go
type mstEdge struct{ from, to, cost int }

func kruskal(vertices int, edges []mstEdge) (int, bool) {
	// Sort the supplied edge slice in place, changing the caller's order.
	// Considering edges from cheapest to most expensive makes each accepted edge
	// the cheapest way to connect two currently separate components.
	sort.Slice(edges, func(i, j int) bool { return edges[i].cost < edges[j].cost })
	groups, total, used := newDisjointSet(vertices), 0, 0
	for _, edge := range edges {
		if groups.union(edge.from, edge.to) {
			// union rejects cycle-forming edges; accepted edges grow the forest.
			total += edge.cost
			used++
			if used == vertices-1 {
				break
			}
		}
	}
	// A nonempty spanning tree must contain exactly vertices-1 accepted edges.
	return total, vertices == 0 || used == vertices-1
}
```

This complete fragment uses the disjoint-set code above and `import "sort"`;
copy `edges` first when its order must survive.

###### Work and memory

Kruskal takes `O(E log E)` time and `O(V + E)` storage
including its edge list and disjoint-set data. Prim with adjacency lists and a
binary heap takes `O(E log V)` time and `O(V + E)` storage including the graph.

If fewer than `V - 1` edges can connect all vertices, the
graph is disconnected and no single spanning tree exists. Kruskal naturally
produces a **minimum spanning forest**: one minimum-cost tree for each separate
connected group. Prim as described above covers only the group containing its
starting vertex; restart it from every still-unseen vertex to produce the same
kind of forest for the whole graph.

###### When to use it

Choose undirected weighted edges that connect every vertex with minimum total edge cost and no cycle.

###### When not to use it

The goal is the cheapest route from one starting vertex,
not the cheapest total network, or when the directed edges do not represent an
undirected connection problem.

###### Common mistakes and edge cases

> **Warning:** Accepting an edge whose endpoints already share a leader creates
> a cycle. In the triangle `A-B=1`, `B-C=2`, `A-C=5`, the last edge must be
> rejected even though it appears in the input.

Validate nonnegative vertex count and every endpoint. Copy the edge slice when
the caller's order must survive, and use checked or wider arithmetic when the
selected edge costs can overflow `int`. This guide treats an empty graph as
successfully spanned with cost zero.

###### Related material

- [C.1.3.25 Disjoint-set union](#c1325-disjoint-set-union) supports Kruskal.
- [C.1.3.27 Heap or priority queue](#c1327-heap-or-priority-queue) supports Prim.
- [C.1.3.23 Dijkstra shortest paths](#c1323-dijkstra-shortest-paths) minimizes routes from one start, not total network cost.

<a id="c1327-heap-or-priority-queue"></a>

##### C.1.3.27 Remove the next most urgent item quickly (heap or priority queue)

###### The problem it solves

Repeatedly remove the smallest, largest, or most urgent stored item while items continue to change.

###### The idea in simple terms

Keep the next item you need (the smallest, largest, or most
urgent) easy to remove without fully sorting everything after each change.

###### How it works

Decide which item must come out next. A min-heap is a priority
queue that exposes the smallest stored item. To retain the three largest values
seen so far, keep three values in a min-heap; when a larger value arrives,
remove the smallest retained value and insert the new one.

###### Step-by-step examples

Insert job priorities `5`, `2`, and `8` into a min-heap; expect removals `2,5,8`.

1. **Insert.** **Before:** heap is empty. **Decision:** add all three priorities while repairing parent-child order. **After:** `2` is at the top. **Known correct:** the top is the smallest stored item; the rest need not be globally sorted.
2. **Remove `2`.** **Before:** three jobs wait. **Decision:** replace the top with the last item and repair downward. **After:** `5` is exposed. **Known correct:** `2` was the required next job.
3. **Finish.** **Before:** `5` and `8` remain. **Decision:** remove the top twice. **After:** output is `2,5,8`. **Known correct:** every removal returned the current minimum.

Peeking or removing an empty heap requires an explicit failure contract.

###### Why it works

The heap contains exactly the items still under consideration, and every
parent has at least as much priority as its children. Any path from the root
therefore cannot reach an item that should precede the root. Push and pop repair
this parent-child rule along one root-to-leaf path, so the root is again the
required next item after each update. Each repair crosses at most the heap's
height and therefore terminates.

###### Go example

[D.5 Priority queues](#d5-priority-queues) owns Go's `heap.Interface` mechanics. For top `k`, keep at most `k` items and replace the root only when a new candidate belongs in the retained set.

###### Work and memory

Building a heap from `n` existing items can be done in
`O(n)` time and uses `O(n)` storage. Adding or removing the top takes
`O(log n)`, and inspecting the top takes `O(1)`. A heap does not provide
constant-time search or removal for an arbitrary value unless an additional
index map is maintained and repaired after every swap.

For `k` sorted inputs containing `N` total items, put the first item from each
non-empty input into a min-heap. Whenever one is removed, add the next item from
that same input. The heap contains at most `k` candidates and its top is the
smallest unread item overall, giving `O(N log k)` time and `O(k)` extra
space. For streams `[1, 4]` and `[2, 3]`, start with candidates `1` and `2`;
removing `1` exposes `4`, then removing `2` exposes `3`, so the output becomes
`1, 2, 3, 4`. If output must be unique, compare against the last emitted value;
do not drop duplicates before they have advanced their source input.

###### When to use it

You repeatedly need the smallest or largest item, a requested
number of the best items, or the next event in priority order.

###### When not to use it

You need the complete final order of an unchanging set of
items; sorting is usually simpler. A heap also does not make it fast to find any
item other than the smallest, largest, or highest-priority one it was designed
to expose.

###### Common mistakes and edge cases

> **Warning:** The heap's backing slice is not globally sorted. After inserting
> `1,4,3`, only the root `1` is guaranteed; reading the remaining slice as the
> next output order can incorrectly report `4` before `3`.

Check emptiness before peek or pop. Define tie-breaking when equal priorities
must be stable, and for a top-`k` heap handle `k=0` before reading its root.

###### Related material

- [D.5 Priority queues](#d5-priority-queues) owns the Go implementation pattern.
- [C.1.3.5 Sorting and selection](#c135-sorting-and-selection) is simpler when every item must be ordered once.
- [C.1.3.29 Monotonic deque](#c1329-monotonic-deque) is linear when candidates expire in sequence order.

<a id="c1328-monotonic-stack"></a>

##### C.1.3.28 Resolve nearest greater or smaller items (monotonic stack)

###### The problem it solves

Find the next or nearest later value that becomes greater or smaller for every position.

###### The idea in simple terms

Keep earlier values whose answer is still unknown, ordered in
only one direction: always increasing or always decreasing. That one-direction
order is what "monotonic" means here.

###### How it works

Decide whether to store positions or values and whether the
stored values increase or decrease from bottom to top. While the new value
answers the question for the top item, remove that item and record its answer.
Then add the new item, whose own answer is still unknown.

###### Step-by-step examples

For `[2,1,4]`, find each next greater value; expect `[4,4,none]`.

1. **Values `2,1`.** **Before:** stack is empty. **Decision:** push positions `0` and `1` because neither answer is known. **After:** stored values decrease as `[2,1]`. **Known correct:** both still await a greater later value.
2. **Value `4`.** **Before:** stack top is `1`. **Decision:** pop `1`, then `2`, and assign `4` to both. **After:** push position `2`. **Known correct:** `4` is the first later greater value for both popped positions.
3. **End.** **Before:** position `2` remains. **Decision:** assign no answer. **After:** results are `[4,4,none]`. **Known correct:** no later value exists for it.

With equal values, choose `<` or `<=` according to whether "greater" is strict; the wrong rule changes duplicate answers.

###### Why it works

Every stored item is still waiting for an answer, and
the stored values remain sorted in one direction. An item is removed only when
the current value is the first later value that answers its question.

###### Go example

```go
func nextGreater(values []int) []int {
	// -1 is the answer for indices that never encounter a greater value.
	result := make([]int, len(values))
	for i := range result {
		result[i] = -1
	}
	// stack holds unresolved indices whose values are non-increasing from bottom
	// to top; equal values can coexist because neither is greater than the other.
	stack := make([]int, 0, len(values))
	for i, value := range values {
		for len(stack) > 0 && values[stack[len(stack)-1]] < value {
			// value is the first greater value to the right of every popped index.
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[top] = value
		}
		stack = append(stack, i)
	}
	return result
}
```

This strict comparison leaves equal values waiting; change the rule only when
the problem asks for greater-or-equal.

###### Work and memory

Each position is added once and removed at most once, so
the whole scan takes `O(n)` time even though one step may remove several items.
The stack can hold `O(n)` positions.

###### When to use it

The prompt asks for the next greater or smaller value, the
nearest earlier or later value that becomes larger or smaller, or how long a
value remains the minimum or maximum.

###### When not to use it

The answer depends on any earlier value, not specifically
the nearest one that meets the condition, or when removing an item would lose
information that a later value still needs.

###### Common mistakes and edge cases

> **Warning:** Using `<=` instead of `<` for a strictly greater query pops an
> equal value too early. For `[2,2,3]`, the first `2` must wait for `3`, not use
> the second `2` as its answer.

The example uses `-1` for no answer; if `-1` is a valid input value, return
indexes, booleans, or another unambiguous representation. Empty input returns
an empty result.

###### Related material

- [C.1.3.13 Stack](#c1313-stack) owns stack behavior.
- [C.1.3.29 Monotonic deque](#c1329-monotonic-deque) also removes dominated values but supports expiration at the front.

<a id="c1329-monotonic-deque"></a>

##### C.1.3.29 Keep only possible moving-window extremes (monotonic deque)

###### The problem it solves

Report the minimum or maximum of every moving window without rescanning each window.

###### The idea in simple terms

Keep only candidates that could still become the minimum or
maximum of a current or future window. The front is the answer for the current
window; the back is where dominated candidates are removed.

For a moving-window maximum, store positions whose values decrease from front
to back:

1. Remove front positions that have left the window.
2. While the new value is at least the value at the back, remove the back; the
   new value is both newer and no smaller, so the old one can never win again.
3. Add the new position. The front now identifies the maximum.

###### How it works

Require `1 <= width <= len(values)` and store positions, not just values. For
each new position, remove expired positions from the front. Then remove back
positions whose values are no larger than the new value, because the new value
expires later and is at least as useful. Append the new position. Once a full
window exists, the front position holds its maximum.

###### Step-by-step examples

For `[1,3,2,5]` with window size `3`, expect maxima `[3,5]`.

1. **Read `1,3,2`.** **Before:** deque begins empty. **Decision:** `3` removes older smaller `1`; `2` remains behind `3`. **After:** stored values are `[3,2]`. **Known correct:** front `3` is the first window maximum.
2. **Read `5`.** **Before:** deque holds candidates `3,2`. **Decision:** remove both because newer `5` is no smaller and expires later. **After:** deque is `[5]`. **Known correct:** `5` is the next window maximum.

Store positions, not only values, so duplicates and expiration are distinguishable. Window size zero is invalid and size one returns every input value.

###### Why it works

Positions remain in increasing time order, values
remain monotonic, and every removed position is either expired or dominated by
a later candidate that will remain available for at least as long.

###### Go example

```go
func windowMaximum(values []int, width int) []int {
	if width <= 0 || width > len(values) {
		return nil
	}
	// deque stores candidate indices in decreasing value order; the front is the
	// maximum for the current window.
	deque, result := make([]int, 0, width), make([]int, 0, len(values)-width+1)
	for i, value := range values {
		if len(deque) > 0 && deque[0] <= i-width {
			// The front index has left the window and can no longer be an answer.
			deque = deque[1:]
		}
		for len(deque) > 0 && values[deque[len(deque)-1]] <= value {
			// Newer value dominates smaller or equal candidates for every future
			// window in which both would appear.
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
		if i+1 >= width {
			// A complete window exists, and its maximum is the front candidate.
			result = append(result, values[deque[0]])
		}
	}
	return result
}
```

Indices preserve both expiration time and duplicate identity.

###### Work and memory

Each position enters and leaves the deque at most once, so
all windows are processed in `O(n)` time with `O(window size)` memory. Store
positions rather than values when expiration depends on location or duplicate
values exist.

Use a heap when priorities change independently or arbitrary items remain
eligible; use a monotonic deque when eligibility expires in sequence order and
new values permanently dominate older ones.

###### When to use it

Report the minimum or maximum of every moving window without rescanning each window.

###### When not to use it

Candidates do not expire in input order, or a newer value
cannot be proved to make an older value permanently irrelevant.

###### Common mistakes and edge cases

> **Warning:** Storing only values cannot identify which duplicate expired.
> With `[5,5,1]` and width `2`, positions distinguish the old `5` leaving from
> the newer `5` that still belongs to the window.

Reject width zero, a negative width, or a width larger than the input. Using
`<=` when pruning equal values keeps the newer duplicate, which is safe because
it expires later.

###### Related material

- [B.3 Queues and deques](#b3-queues-and-deques) owns deque operations.
- [C.1.3.14 Sliding window](#c1314-sliding-window) owns boundary movement.
- [C.1.3.27 Heap or priority queue](#c1327-heap-or-priority-queue) handles priorities that do not expire in sequence order.

<a id="c1330-intervals"></a>

##### C.1.3.30 Put overlapping ranges together (intervals)

###### The problem it solves

Combine overlapping ranges, report gaps, or choose compatible ranges under an explicit endpoint rule.

###### The idea in simple terms

Each interval covers everything between a start and an end.
Sort intervals by start so overlaps appear next to one another.

###### How it works

Sort by start and keep one current interval. Compare the next
start with the current end using the problem's rule for touching endpoints. If
they overlap, extend the current end when necessary. If a gap separates them,
save the current interval and begin a new one.

###### Step-by-step examples

Merge `[5,7]`, `[1,3]`, and `[2,6]`; expect `[1,7]`.

1. **Sort.** **Before:** possible overlaps are separated. **Decision:** order by start. **After:** `[1,3],[2,6],[5,7]`. **Known correct:** any interval that can extend the current one appears before a later gap.
2. **Merge `[2,6]`.** **Before:** current is `[1,3]`. **Decision:** start `2` is within the current end, so extend to `6`. **After:** current is `[1,6]`. **Known correct:** both ranges are covered.
3. **Merge `[5,7]`.** **Before:** current ends at `6`. **Decision:** `5` overlaps, so extend to `7`. **After:** save `[1,7]`. **Known correct:** every input point is represented once in the merged coverage.

Whether `[1,2]` and `[2,3]` overlap depends on whether endpoints are inclusive or half-open; decide before comparing.

###### Why it works

Once a gap appears, every remaining
interval starts even later because the input is sorted. None can overlap the
interval you just saved.

###### Go example

```go
type interval struct{ start, end int }

func mergeIntervals(intervals []interval) []interval {
	if len(intervals) == 0 {
		return nil
	}
	// Sort the supplied slice in place by start, changing the caller's order.
	// Every later interval then begins no earlier than the current merged interval.
	sort.Slice(intervals, func(i, j int) bool { return intervals[i].start < intervals[j].start })
	result := []interval{intervals[0]}
	for _, next := range intervals[1:] {
		last := &result[len(result)-1]
		if next.start <= last.end {
			// Overlap means only the farther end can enlarge the merged interval.
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

This version uses inclusive endpoints and requires `import "sort"`.

###### Work and memory

Sorting `n` intervals takes `O(n log n)` time, and the
following scan takes `O(n)`. Sorting in place uses no separate interval copy,
but the merged or selected result can contain `O(n)` intervals.

###### When to use it

The input contains ranges, and the output must combine
overlapping ranges or report the gaps between covered regions.

###### When not to use it

The ranges change online and each change must be queried
immediately, or when the problem is really about individual points rather than
relationships between ranges.

###### Common mistakes and edge cases

The problem must define whether `[1, 2]` and `[2, 3]` overlap.
If both ending points are included, the ranges share `2`. If an end marks the
first excluded value, the first range stops immediately before `2`, so they do
not overlap.

Different interval objectives require different rules:

- To select the largest number of non-overlapping intervals, sort by end and
  greedily accept the next compatible interval. Finishing earliest leaves every
  later option available.
- To maximize total interval weight or profit, earliest finish is not enough.
  Sort by end, binary-search the last compatible interval for each choice, and
  use DP to compare taking it plus its compatible prefix against skipping it.
- To count simultaneous intervals or required rooms, sort start/end events or
  keep current ending times in a min-heap. The endpoint convention decides
  whether an end at time `t` frees a resource before a start at `t`.

###### Related material

- [C.1.3.5 Sorting and selection](#c135-sorting-and-selection) owns ordering.
- [C.1.3.31 Sweep line](#c1331-sweep-line) counts simultaneous ranges.
- [C.1.3.18 Greedy](#c1318-greedy) owns earliest-finish selection.

<a id="c1331-sweep-line"></a>

##### C.1.3.31 Process only positions where activity changes (sweep line)

###### The problem it solves

Count active ranges or another quantity that changes only at range starts and ends.

###### The idea in simple terms

Move from left to right across a timeline or number line.
Record only the positions where something starts or ends, because nothing
changes between those positions.

###### How it works

Turn each range into two events. A start event adds one to the
active count, and an end event subtracts one. Sort the events by position and
update the count in that order.

###### Step-by-step examples

Bookings `[1,4)` and `[3,5)` need at most two rooms.

1. **Create events.** **Before:** ranges are stored whole. **Decision:** emit `+1@1`, `+1@3`, `-1@4`, and `-1@5`. **After:** only change positions remain. **Known correct:** activity is constant between events.
2. **Sweep.** **Before:** active count is `0`. **Decision:** apply events in position order. **After:** counts are `1,2,1,0`; maximum is `2`. **Known correct:** each count equals the ranges active after its event.

For `[1,4)` and `[4,5)`, process the end at `4` before the start at `4`; otherwise the sweep wrongly reports two simultaneous rooms.

###### Why it works

After processing all events at a position, the running count
equals the number of ranges active immediately after that position.

###### Go example

```go
type sweepEvent struct{ position, delta int }

func maximumOverlap(events []sweepEvent) int {
	// Sort the supplied event slice in place, changing the caller's order. At a
	// tie, process negative end deltas before positive start deltas to model
	// half-open intervals [start,end).
	sort.Slice(events, func(i, j int) bool {
		if events[i].position == events[j].position {
			return events[i].delta < events[j].delta
		}
		return events[i].position < events[j].position
	})
	// active is the overlap immediately after applying each ordered boundary;
	// best records the largest such count.
	active, best := 0, 0
	for _, event := range events {
		// Start events increase the active count and end events decrease it.
		active += event.delta
		if active > best {
			// Preserve the largest overlap reached after any boundary update.
			best = active
		}
	}
	return best
}
```

Negative end events sort before positive start events at the same position,
matching non-empty half-open ranges; this requires `import "sort"`. Ignore or
reject an empty range before creating its start and end events.

###### Work and memory

Creating and sorting two events per range takes
`O(n log n)` time and `O(n)` event storage. The sweep itself is linear after
sorting.

###### When to use it

You need the maximum number of simultaneous bookings, the
number of active ranges at each position, or another result that changes only
when a range starts or ends.

###### When not to use it

Values change continuously between event positions, or
events arrive online and must be answered before they can be collected and
sorted.

###### Common mistakes and edge cases

Events at the same position need an explicit rule. If a booking
ending at `4` frees its room before another starts at `4`, process the ending
before the new start.

###### Related material

- [C.1.3.30 Intervals](#c1330-intervals) owns overlap and endpoint conventions.
- [C.1.3.16 Difference arrays](#c1316-difference-arrays) is a discrete sweep over a small fixed coordinate range.

<a id="c1332-fast-and-slow-pointers"></a>

##### C.1.3.32 Use speed or a fixed gap to reveal structure (fast and slow pointers)

###### The problem it solves

Detect a linked cycle, find a middle, or retain a fixed distance from the end without storing the full path.

###### The idea in simple terms

Move two positions through the same linked list. Their speeds
or the distance between them reveals a loop, the middle, or a node near the end.

###### How it works

Decide how fast each position moves or how far apart they must
remain. Before following the link to the next node, check that the current node
exists. For a fixed distance, move the leading position ahead first and then
move both one node at a time.

###### Step-by-step examples

For `1->2->3->4->nil`, return the second middle `3`.

1. **First move.** **Before:** both pointers are at `1`. **Decision:** move `slow` once and `fast` twice. **After:** they are at `2` and `3`. **Known correct:** `fast` has traveled twice as far.
2. **Second move.** **Before:** `fast` can advance toward the end. **Decision:** move again. **After:** `slow` is at `3` and `fast` reaches `nil`. **Known correct:** under this stopping rule, `slow` is the second middle.

For even lengths, a different loop condition can return the first middle. Check both `fast` and the next link before a two-step move; empty input returns no node.

###### Why it works

For the middle-node loop, after `t` iterations `slow` has moved `t` links and
`fast` has moved `2t`. When `fast` reaches the end, `slow` has covered half the
list; the loop condition chooses the later middle for even lengths. For a fixed
gap, moving both pointers together preserves the established distance, so when
the leader reaches the end the follower has the requested distance from it.
For cycle detection, once both pointers are inside a finite cycle, their
relative offset changes by one position per iteration and must eventually
become zero. Every form moves toward an end or around a finite cycle, so it
terminates under its stated stopping rule.

###### Go example

```go
type listNode struct {
	value int
	next  *listNode
}

func secondMiddle(head *listNode) *listNode {
	// slow advances one node while fast advances two. When fast reaches the end,
	// slow has covered half the list.
	slow, fast := head, head
	for fast != nil && fast.next != nil {
		slow, fast = slow.next, fast.next.next
	}
	// Starting both at head selects the second middle for an even-length list.
	return slow
}
```

The loop returns the later middle for even lengths and `nil` for an empty list.

###### Work and memory

Each pointer makes at most a linear number of moves, so the
method takes `O(n)` time and `O(1)` extra memory.

###### When to use it

A linked list may loop back to an earlier node, you need its
middle without first counting every node, or one position must remain a fixed
number of nodes behind another.

The structure need not be stored as nodes. Repeatedly applying a deterministic
function `next(state)` creates a **functional graph** in which every state has
one outgoing edge. Fast and slow states can therefore detect repetition in
digit transforms or other generated sequences without storing every seen state.

###### When not to use it

You need to jump directly to any chosen position, remember
the full history, or return the actual visited path. Two moving positions do not
retain those details.

###### Common mistakes and edge cases

> **Warning:** Testing only `fast != nil` before `fast.next.next` panics on a
> one-node suffix. Require both `fast != nil` and `fast.next != nil`.

Empty input returns no node. State whether an even-length list needs the first
or second middle, and establish a requested fixed gap before moving both
positions.

###### Related material

- [B.2 Linked Lists](#b2-linked-lists) owns node and cycle behavior.
- [C.1.3.9 Two pointers](#c139-two-pointers) owns the broader coordinated-position method.
- [C.1.3.4 Hash map or set](#c134-hash-map-or-set) retains full visited membership when needed.

<a id="c1333-trie"></a>

##### C.1.3.33 Share stored word beginnings (trie)

###### The problem it solves

Insert and query many words by exact value or shared beginning without scanning every stored word.

###### The idea in simple terms

Store words one character at a time. Words with the same
beginning reuse the same initial steps.

###### How it works

Each stored step records possible next characters and whether a
complete word ends there. To insert a word, follow or create one step for each
character. To look up a word, follow the same steps and check the end marker
after the last character.

###### Step-by-step examples

Insert `car` and `cat`; expect prefix `ca` to exist but full word `ca` not to exist.

1. **Insert `car`.** **Before:** root has no children. **Decision:** create `c->a->r` and mark `r` as a word end. **After:** one word is stored. **Known correct:** its complete path spells `car`.
2. **Insert `cat`.** **Before:** `c->a` already exists. **Decision:** reuse it, create child `t`, and mark that end. **After:** the words share their prefix. **Known correct:** both full paths exist.
3. **Query `ca`.** **Before:** both edges exist. **Decision:** prefix lookup succeeds, but inspect the end marker for full-word lookup. **After:** prefix is true and word is false. **Known correct:** a stored beginning is not confused with a stored complete word.

Empty-string insertion is represented by the root end marker if the contract allows it. A wildcard may branch to many children and lose linear lookup time.

Use a fixed child array for a small fixed alphabet and a map for a sparse or
large alphabet. A wildcard search branches over every existing child at that
position, so its worst case can be exponential in the number of wildcards even
though ordinary exact lookup is `O(word length)`. During a board search, carry
the trie node reached by the current path so a prefix missing from the trie can
stop that branch immediately.

###### Why it works

The characters followed from the starting point always spell the
beginning represented by the current step. A word exists only if the last step
is marked as the end of a complete word.

###### Go example

```go
type trieNode struct {
	children map[byte]*trieNode
	word     bool
}

func (root *trieNode) insert(word string) {
	node := root
	for i := range len(word) {
		// Allocate each sparse child map only when this prefix first needs one.
		if node.children == nil {
			node.children = make(map[byte]*trieNode)
		}
		if node.children[word[i]] == nil {
			// Missing byte edges create new prefix nodes; existing edges are shared
			// with words that have the same prefix.
			node.children[word[i]] = &trieNode{}
		}
		node = node.children[word[i]]
	}
	// The terminal flag distinguishes a complete word from a prefix alone.
	node.word = true
}

func (root *trieNode) find(text string, wholeWord bool) bool {
	node := root
	// Follow one byte edge per input position; a missing edge proves absence.
	for i := range len(text) {
		node = node.children[text[i]]
		if node == nil {
			return false
		}
	}
	// Prefix lookup needs only the path; whole-word lookup also needs the terminal.
	return !wholeWord || node.word
}
```

This byte trie requires one-byte alphabet semantics; use runes when the problem
defines positions as Unicode code points.

###### Work and memory

Inserting or looking up a word of length `L` takes `O(L)`
time. An insertion may create `O(L)` nodes; across all words, storage is
proportional to the number of distinct prefix nodes. Wildcards can require
exploring many branches and may take exponential time in the number of
wildcards.

###### When to use it

You repeatedly ask for words beginning with the same letters,
check dictionary words one character at a time, or want to stop as soon as a
beginning does not exist.

###### When not to use it

You only need to know whether a complete word exists. A hash
set is usually simpler and may use less memory.

###### Common mistakes and edge cases

> **Warning:** Treating a reached prefix as a complete word makes `find("ca",
> true)` succeed after inserting only `"car"`; require the end marker for a
> whole-word query.

The methods require a non-nil root. Decide whether the empty string is allowed,
and use runes or visible-symbol segmentation instead of bytes when the input
contract requires those text units.

###### Related material

- [C.1.3.4 Hash map or set](#c134-hash-map-or-set) is simpler for exact whole-word membership.
- [D.6 Bytes, runes, and visible characters](#d6-bytes-runes-and-visible-characters) determines what one trie step represents.
- [B.4 Trees and binary-search trees](#b4-trees-and-binary-search-trees) defines tree paths.

#### C.1.4 Choosing among overlapping methods

##### The problem it solves

Several methods may fit the same input shape, but their discard rules, update
support, or required output differ.

##### The idea in simple terms

Compare the fact each candidate needs to remain true, not just their names or
surface syntax. The first differing requirement usually decides the method.

##### How it works

- Use a [sliding window](#c1314-sliding-window) when moving a boundary updates a
  consecutive range predictably; use [prefix sums](#c1315-prefix-sums) when a
  fixed range answer comes from two saved totals.
- Use prefix sums for immutable data; use [Fenwick or segment trees](#c1317-fenwick-and-segment-trees)
  when updates and range queries alternate.
- Use [two pointers](#c139-two-pointers) when coordinated moves discard
  possibilities; use [binary search](#c136-binary-search) when one middle test
  independently discards half.
- Use [greedy](#c1318-greedy) after a safe permanent-choice proof;
  [dynamic programming](#c1310-dynamic-programming) when competing repeated
  states must remain; and [backtracking](#c1319-backtracking) when individual
  constructions must be explored and undone.
- Use [divide and conquer](#c1311-divide-and-conquer) for independent parts and
  dynamic programming for overlapping parts.
- DFS is an exploration order; backtracking adds and removes candidate choices
  during that order. Use [DFS](#c1320-depth-first-search) for complete branch
  exploration and [BFS](#c1321-breadth-first-search) for nearest equal-cost
  results.
- Use BFS for equal edge costs, [Dijkstra](#c1323-dijkstra-shortest-paths) for
  unequal non-negative costs, and [Bellman-Ford](#c1324-bellman-ford-shortest-paths)
  for negative edges or reachable negative-cycle detection.
- Fast/slow pointers specialize two pointers for linked cycles, middles, and
  fixed gaps.
- Sort when every item is needed in order, use quickselect for one rank in
  mutable batch data, and use a heap for streaming top-`k` or repeated removal.
- A monotonic stack resolves nearest relationships; a monotonic deque maintains
  an extreme while window positions expire.

##### Step-by-step examples

A road problem asks for cheapest routes from `A`.

1. **Equal costs.** **Before:** every edge costs one move. **Decision:** choose
   BFS because FIFO layers are cost layers. **After:** no heap is needed.
   **Known correct:** first discovery has the fewest moves and lowest cost.
2. **Varying non-negative costs.** **Before:** edges cost `1`, `4`, or `9`.
   **Decision:** choose Dijkstra and remove the cheapest known route next.
   **After:** FIFO order is rejected. **Known correct:** non-negative edges make
   finalization safe.
3. **Negative cost.** **Before:** one edge costs `-3`. **Decision:** reject
   Dijkstra and choose Bellman-Ford. **After:** repeated relaxation remains.
   **Known correct:** a later negative edge may improve an earlier route.

##### Why it works

The comparison preserves every method whose assumptions still match the
problem and removes one only when a required fact fails. Once one candidate's
discard rule and resource bound both hold, its own card supplies the complete
correctness proof.

##### Go example

This comparison produces an implementation choice rather than shared runtime
code. Follow the selected card's Go boundary and its linked implementation
owner.

##### Work and memory

Comparison is constant-sized planning work. The selected algorithm still owns
all preprocessing, runtime, retained state, recursion, and output costs.

##### When to use it

Use it when two named methods appear plausible after the quick decision tree.

##### When not to use it

Do not compare labels before stating the input and output requirements; the
real method may be outside the apparent pair.

##### Common mistakes and edge cases

> **Warning:** Dijkstra is not "BFS with a heap" when negative edges exist.
> The negative edge removes the proof that a popped distance is final.

Also distinguish batch from streaming input, immutable from changing data,
one answer from every answer, and actual output cost from auxiliary memory.

##### Related material

- [C.1.1 Ask what the input and output allow](#c11-ask-what-the-input-and-output-allow-quick-decision-tree)
  produces the initial candidates.
- [C.1.6 Before committing to a method](#c16-before-committing-to-a-method)
  verifies the final choice.

#### C.1.5 Let constraints eliminate bad choices

##### The problem it solves

A correct method may still be unusable because the largest valid input makes
its work, memory, or output too large.

##### The idea in simple terms

Estimate growth before coding and use the constraints to cross out methods.
Big-O describes that growth; [A.1](#a1-complexity-and-resource-accounting)
defines the notation and its limits.

##### How it works

- Around `n <= 10`, trying every ordering (`n!`) may be practical with pruning.
- Around `n <= 20`, trying every subset (`2^n`) may be practical.
- Around `n <= 100`, cubic work may fit and pairwise work is often reasonable.
- Around `n <= 1,000`, quadratic work needs scrutiny.
- Around `n <= 100,000`, target sorting (`O(n log n)`) or a few linear passes.
- At a million or more items, favor few passes and closely budget large maps or
  arrays.

These are rough interview checks, not machine-independent guarantees.

##### Step-by-step examples

For `n=100,000`, compare pairwise search with sorting.

1. **Pairwise candidate.** **Before:** roughly `n^2` pairs are possible.
   **Decision:** estimate about ten billion checks. **After:** reject it.
   **Known correct:** its growth is incompatible with the input scale.
2. **Sorting candidate.** **Before:** `n log n` comparisons are expected.
   **Decision:** retain it for further proof and memory review. **After:** one
   plausible candidate remains. **Known correct:** constraints have not ruled
   it out.

If the output requires all `2^n` subsets, writing the result already takes
exponential work; no generator can make total runtime linear in `n`.

##### Why it works

An upper bound that already exceeds the budget cannot become acceptable through
implementation detail alone. Constraint filtering only rejects; it does not
prove the remaining algorithm correct.

##### Go example

There is no runtime helper for this planning step. In Go, include meaningful
allocations, integer sizes, recursion depth, and garbage-retained storage in the
estimate before choosing a representation.

##### Work and memory

The estimate itself is constant planning work. Always count preprocessing,
per-state transitions, output construction, recursion, and retained buffers in
the candidate's actual bound.

##### When to use it

Use it before implementation and whenever a direct solution seems correct but
possibly too slow or memory-heavy.

##### When not to use it

Do not treat the rough thresholds as guarantees; benchmark or use stated limits
when constants, language costs, or strict time caps dominate.

##### Common mistakes and edge cases

> **Warning:** An `O(n)` algorithm that stores a million large objects may be
> less feasible than an in-place `O(n log n)` algorithm.

Use the correct dimension: vertices plus edges for graphs, rows times columns
for grids, reachable states times transitions for DP, and output size for
generation.

##### Related material

- [A.1 Complexity and resource accounting](#a1-complexity-and-resource-accounting)
  owns growth notation and Go allocation accounting.
- [A.7 Combinatorics and output size](#a7-combinatorics-and-output-size)
  explains factorial and exponential output lower bounds.

#### C.1.6 Before committing to a method

##### The problem it solves

A plausible method still needs a precise state meaning, safety argument,
termination argument, coverage claim, and resource bound before code begins.

##### The idea in simple terms

Write the proof obligations as short sentences. If one cannot be made precise,
the method is still a guess.

##### How it works

1. **Stored meaning:** define every variable, table entry, queue item, or stack
   item.
2. **Invariant:** state the fact that remains true before and after each step.
3. **Progress:** name the bounded measure that moves toward completion.
4. **Coverage:** explain why discarded choices cannot contain an answer.
5. **Work and memory:** count processing frequency and retained information.

Then challenge the plan with empty input, equal or negative values, immutable
or streaming input, path reconstruction, repeated queries, and removal of the
assumption that enabled the optimization.

##### Step-by-step examples

For lower-bound binary search on `[1,3,3,5]`:

1. **Meaning.** **Before:** `left` and `right` are undefined. **Decision:** let
   `[left,right)` contain every possible first valid position. **After:** start
   with `[0,4)`. **Known correct:** the answer has not been discarded.
2. **Progress and coverage.** **Before:** a middle position is tested.
   **Decision:** keep only the half that can contain the first valid position.
   **After:** the interval strictly shrinks. **Known correct:** discarded
   positions were proved too early or too late.
3. **Cost and boundaries.** **Before:** the plan handles the ordinary case.
   **Decision:** test empty input, duplicates, and insertion at `len(values)`.
   **After:** the half-open contract covers all three. **Known correct:** the
   method and boundary behavior agree.

##### Why it works

Together, the five sentences show that the stored state is meaningful, every
step is safe, the process finishes, all possible answers are covered, and the
implementation fits the limits. That is exactly what a correctness and
feasibility argument must establish.

##### Go example

Translate each sentence into a nearby name, condition, or assertion. For the
trace, [E.1 Binary search bounds](#e1-binary-search-bounds) uses the same
`left`, `right`, and half-open range contract.

##### Work and memory

Writing the checklist is constant planning work. Its final item records the
selected method's complete cost, including setup, recursion, output, and
meaningful Go allocations.

##### When to use it

Use it after selecting a candidate and before implementation, and repeat it
when requirements or constraints change.

##### When not to use it

Do not replace an actual trace with slogans such as "it is greedy" or "binary
search halves the work." The sentences must name this problem's state and rule.

##### Common mistakes and edge cases

> **Warning:** A loop that usually moves is not enough. If one equality branch
> keeps the same middle position, the progress argument fails and the loop may
> never terminate.

Also verify mutation permission, integer overflow, empty output, duplicates,
unreachable graph states, and whether an optimization prevents reconstructing
the required path or indices.

##### Related material

- [A.2 Correctness and invariants](#a2-correctness-and-invariants) owns the
  proof structure.
- [C.1.5 Let constraints eliminate bad choices](#c15-let-constraints-eliminate-bad-choices)
  supplies the resource screen.
- [C.1.3 Pattern cards](#c13-pattern-cards) provides method-specific checks.

<a id="e-go-techniques"></a>

## D. Go techniques

These are Go-specific details that commonly affect otherwise correct
algorithms. Each rule below states the behavior that motivates it.

<a id="e1-keep-input-conversion-outside-the-algorithm"></a>

### D.1 Keep input conversion outside the algorithm

#### The problem it solves

An automated judge, command-line program, or web handler often receives text,
while an algorithm needs values such as `[]int`, `string`, or `*Node`. Mixing
text decoding, input checks, algorithm steps, and output encoding in one
function makes failures harder to locate and the algorithm harder to reuse.

#### The idea in simple terms

Treat the program boundary like a reception desk. The desk translates an
outside request into the program's own types, rejects invalid requests, and
translates the answer back. This small translation layer is called a
**boundary adapter**. The algorithm behind it receives only problem values and
does not need to know whether they came from JSON, a command line, or a test.

#### How it works

1. Decode the outside representation into a problem-specific Go type.
2. Validate rules that belong to the input contract, such as required fields
   or legal ranges. Return a useful error before calling the algorithm when the
   request is invalid.
3. Pass the typed fields to an algorithm whose parameters and return values
   describe the problem itself.
4. Convert the algorithm's result into the outside representation.

JSON means JavaScript Object Notation, a text format built from objects,
arrays, strings, numbers, booleans, and `null`. Go's `encoding/json` package
performs the JSON-specific work in the example below.

#### Step-by-step examples

Decode `raw = {"values":[2,7,11,15],"target":9}`. The expected result is
`pairInput{Values: []int{2, 7, 11, 15}, Target: 9}`.

1. **Step: Decode the JSON object**
   - **Before:** `input` has zero values: `Values=nil`, `Target=0`.
   - **Decision:** The JSON has the expected field names and value types, so
     `json.Unmarshal` can assign both fields.
   - **After:** `input.Values=[2,7,11,15]`, `input.Target=9`, and `err=nil`.
   - **Known correct:** The adapter now holds the same four values and target
     represented by `raw`.
2. **Step: Validate the typed request**
   - **Before:** `len(input.Values)=4`.
   - **Decision:** Four values satisfy this adapter's requirement of at least
     two values.
   - **After:** The adapter returns `input` unchanged.
   - **Known correct:** A called pair-search algorithm receives a valid typed
     request and no JSON concerns.

For a failure case, use `raw = {"values":[2,7],"target":"nine"}`. The expected
result is a decoding error and no usable `pairInput`.

1. **Step: Decode the malformed target**
   - **Before:** `input` contains only zero values.
   - **Decision:** A JSON string cannot be assigned to the Go `int` field
     `Target`, so `json.Unmarshal` returns an error.
   - **After:** `parsePairInput` returns a wrapped error instead of continuing
     to validation or an algorithm call.
   - **Known correct:** No partly decoded request crosses the boundary.

#### Why it works

After a successful parse, every field has the Go type promised to the
algorithm. Validation then establishes any additional input rules before the
algorithm runs. If either step fails, returning immediately prevents partial
state from being mistaken for valid input. Encoding happens only after a typed
result exists, so representation failures remain separate from algorithm
failures.

#### Go example

```go
type pairInput struct {
	// JSON tags bind the external field names to strongly typed Go fields.
	Values []int `json:"values"`
	Target int   `json:"target"`
}

func parsePairInput(raw []byte) (pairInput, error) {
	// Decode at the boundary so the algorithm itself can accept pairInput rather
	// than inspect untyped bytes.
	var input pairInput
	if err := json.Unmarshal(raw, &input); err != nil {
		// Wrap the decoder error with operation context while preserving it for
		// errors.Is or errors.As.
		return pairInput{}, fmt.Errorf("decode pair input: %w", err)
	}
	// Structural decoding cannot enforce this problem-specific requirement.
	if len(input.Values) < 2 {
		return pairInput{}, fmt.Errorf("pair input needs at least two values")
	}
	return input, nil
}

func formatPairOutput(left, right int) ([]byte, error) {
	// A fixed-size array guarantees that successful output has exactly two indices.
	return json.Marshal([2]int{left, right})
}
```

The struct tags such as ``json:"values"`` connect JSON field names to Go
fields. Returning `%w` in the error keeps the original decoding error available
to callers while adding boundary context.

#### Work and memory

For `b` input bytes containing `n` numbers, decoding examines the input and
builds the `Values` slice, so it takes `O(b)` work and `O(n)` additional memory.
Validation here takes constant work. Formatting two indices takes constant
work and a small output allocation. These boundary costs are separate from the
called algorithm's own work and memory.

#### When to use it

Use a boundary adapter when stored or transmitted data differs from the types
that make the algorithm easiest to state, when several callers should share
one algorithm, or when malformed input needs a clear boundary error.

#### When not to use it

Do not add serialization to a function whose caller already supplies the
correct typed values. Do not put problem reasoning in a generic decoder; keep
problem-specific validation beside the problem adapter.

#### Common mistakes and edge cases

> **Warning:** Ignoring a decode error can pass zero values into the algorithm.
> For `{"target":"nine"}`, using `input.Target` after a failed unmarshal may
> incorrectly search for target `0`; return the error immediately.

An absent JSON field also leaves its Go field at the zero value. If zero is
valid but omission is not, use a pointer field or a custom presence check to
distinguish them. Decide whether `null`, an empty array, and unknown fields are
allowed instead of accepting them accidentally.

#### Related material

- [E.3 Hash-map complement lookup](#e3-hash-map-complement-lookup) owns the
  complete pair-search implementation.
- [D.2 Table-driven tests](#d2-table-driven-tests) keeps valid and invalid
  boundary cases under the same test logic.
- [D.4 Map presence and sets](#d4-map-presence-and-sets) explains the related
  problem of distinguishing absence from a stored zero value.

<a id="e2-table-driven-tests"></a>

### D.2 Table-driven tests

#### The problem it solves

One function usually needs several examples, boundary cases, and failure
cases. Hand-writing separate test logic for each case invites small differences
that can hide bugs or make failures difficult to identify.

#### The idea in simple terms

Write the checking recipe once and feed it a list of named examples. In Go,
this is called a **table-driven test**: the table is a slice of test-case
structs, and each row contains the input and expected result for one run.

#### How it works

1. Define a struct whose fields hold one case's name, inputs, and expected
   outputs.
2. Build a slice containing ordinary, boundary, and adversarial rows.
3. Loop over the slice and call `t.Run` so each row is reported by name.
4. Inside every subtest, call the same behavior and compare every promised
   result with that row's expectation.
5. Stop a row immediately with `t.Fatal` or `t.Fatalf` when setup fails and
   later checks would be meaningless.

#### Step-by-step examples

Run rows `{name:"positive even", value:4, want:true}` and
`{name:"negative odd", value:-3, want:false}`. The expected outcome is two
passing subtests.

1. **Row: `positive even`**
   - **Before:** `test.value=4`, `test.want=true`.
   - **Decision:** `isEven(4)` returns `true`, which equals `test.want`.
   - **After:** The row records no error.
   - **Known correct:** The ordinary even case used the shared assertion.
2. **Row: `negative odd`**
   - **Before:** `test.value=-3`, `test.want=false`.
   - **Decision:** In Go, `-3%2` is `-1`, so the result is `false` and matches
     `test.want`.
   - **After:** The row records no error.
   - **Known correct:** A negative value used exactly the same call and check.

For a boundary row, use `{name:"zero", value:0, want:true}`. The expected
outcome is a passing subtest because `0%2 == 0`.

1. **Row: `zero`**
   - **Before:** `test.value=0`, `test.want=true`.
   - **Decision:** `isEven(0)` returns `true`.
   - **After:** The comparison passes without special-case test logic.
   - **Known correct:** The boundary value is covered by the same contract as
     every other row.

#### Why it works

Every row reaches the same call and comparison, so differences in results come
from the case data rather than from different test code. Naming each subtest
ties a failure to its input. The loop is finite because the table is finite;
after all rows pass, the tested contract holds for every listed case. This is
evidence for those cases, not proof for every possible input.

#### Go example

```go
func isEven(value int) bool {
	// Multiples of two, including zero and negative evens, have remainder zero.
	return value%2 == 0
}

func TestIsEven(t *testing.T) {
	// Each row names one behavior and keeps its input beside the expected result.
	tests := []struct {
		name  string
		value int
		want  bool
	}{
		{name: "positive even", value: 4, want: true},
		{name: "negative odd", value: -3, want: false},
		{name: "zero", value: 0, want: true},
	}

	// Run every row as a named subtest so a failure identifies the exact case.
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Compare the actual result with the row's expectation and report both.
			if got := isEven(test.value); got != test.want {
				t.Errorf("isEven(%d) = %v, want %v", test.value, got, test.want)
			}
		})
	}
}
```

The outer test reports the group; `t.Run` creates a named subtest for each row.
Use `t.Helper` in extracted assertion helpers so failures point to the calling
row rather than the helper body.

#### Work and memory

With `r` rows and per-row behavior costing `T`, the loop does `O(r * T)` work.
The table itself stores `O(r)` case data. Sequential subtests add only bounded
runner state beyond the current row, although individual inputs and outputs
may allocate their own memory.

#### When to use it

Use table-driven tests when many inputs share one contract and assertion shape,
especially for boundary values, malformed input, duplicates, and regression
cases that should remain named.

#### When not to use it

Do not force unrelated behaviors with different setup or assertions into one
wide table. Separate focused tests are clearer for long stateful workflows;
randomized or generated-input testing is a better companion when a property
must be checked across many automatically produced inputs.

#### Common mistakes and edge cases

> **Warning:** An empty table executes zero assertions and still passes. A test
> with `tests := []caseType{}` proves nothing; keep at least one ordinary row
> and the relevant boundary rows.

When a row owns a slice, map, or pointer that the tested function mutates,
freshly construct or clone that value per row so another row cannot inherit
changed state. If setup returns an error, call `t.Fatalf` before reading the
partly initialized result.

#### Related material

- [A.2 Correctness and invariants](#a2-correctness-and-invariants) separates
  test evidence from a general correctness argument.
- [D.1 Keep input conversion outside the algorithm](#d1-keep-input-conversion-outside-the-algorithm)
  identifies parsing failures that should become explicit rows.
- [E.9 Sorted two-pointer pair search](#e9-sorted-two-pointer-pair-search) lists
  representative cases for a reference implementation.

<a id="e3-sorting-without-unexpected-mutation"></a>

### D.3 Sorting without unexpected mutation

#### The problem it solves

Go's sorting functions rearrange the slice they receive. A caller that still
needs the original order can therefore observe surprising changes after a
helper sorts its input.

#### The idea in simple terms

Make a separate row of cards before arranging them. Changing a value that the
caller can observe is called **mutation**. A shallow slice copy creates a new
backing array for the slice elements, so sorting the copy does not reorder the
caller's slice.

#### How it works

1. Decide whether the function's contract allows it to reorder the input.
2. If not, clone the slice before sorting. A **clone** here means a new slice
   with the same element values.
3. Sort the clone with `sort.Ints`, `sort.Strings`, or an explicit comparison
   for structured values.
4. Return or use the sorted clone, leaving the input order unchanged.

The copy is shallow: pointer, slice, or map elements still refer to the same
nested values. It protects the outer ordering, not mutations inside an element.

#### Step-by-step examples

Call `sortedCopy(values)` with `values = [3,1,2]`. The expected result is
`sorted = [1,2,3]` while `values` remains `[3,1,2]`.

1. **Step: Clone the slice**
   - **Before:** `values` refers to one backing array containing `[3,1,2]`.
   - **Decision:** The contract promises not to reorder `values`, so
     `slices.Clone` creates a separate backing array for `sorted`.
   - **After:** Both slices contain `[3,1,2]`, but changing one array's elements
     cannot change the other array's elements.
   - **Known correct:** The caller's order is isolated before sorting starts.
2. **Step: Sort the clone**
   - **Before:** `sorted=[3,1,2]`, `values=[3,1,2]`.
   - **Decision:** `sort.Ints` rearranges only `sorted` into ascending order.
   - **After:** `sorted=[1,2,3]`, `values=[3,1,2]`.
   - **Known correct:** The result is ordered and the promised input state is
     preserved.

For the adversarial version, assign `sorted := values` instead of cloning. The
expected observation after `sort.Ints(sorted)` is that both variables show
`[1,2,3]`.

1. **Step: Copy only the slice descriptor**
   - **Before:** `values=[3,1,2]`.
   - **Decision:** Assignment copies the pointer, length, and capacity, not the
     backing array.
   - **After:** `sorted` and `values` view the same elements.
2. **Step: Sort through the alias**
   - **Before:** Both slices view `[3,1,2]`.
   - **Decision:** Sorting writes into the shared backing array.
   - **After:** Both show `[1,2,3]`; the no-mutation contract is broken.

#### Why it works

Before sorting begins, the clone holds the same element values in an
independent outer array. Sorting permutes only that array, so the returned
values are ordered while every position in the caller's array keeps its
original value. The sorting call finishes after a bounded number of
comparisons and swaps, and the independent arrays preserve the two promised
observations.

#### Go example

```go
func sortedCopy(values []int) []int {
	// Clone allocates independent storage; sort.Ints may now mutate sorted without
	// changing the caller's slice.
	sorted := slices.Clone(values)
	sort.Ints(sorted)
	return sorted
}
```

`slices.Clone(nil)` returns `nil`; a non-nil input returns a slice with copied
elements. For records sorted with `sort.Slice` or `slices.SortFunc`, state the
primary order and every required tie-break rule in the comparison.

#### Work and memory

For `n` integers, cloning copies `n` values in `O(n)` time and allocates `O(n)`
additional memory. Comparison sorting takes `O(n log n)` time; the standard
integer sort does not allocate another result slice. The combined bounds are
`O(n log n)` time and `O(n)` additional memory because the clone is retained as
the output.

#### When to use it

Clone before sorting when the function promises to preserve input order, when
the caller may reuse the input, or when tests compare the input before and
after the call.

#### When not to use it

Do not pay for a clone when the contract explicitly allows in-place reordering
and the caller no longer needs the old order. A shallow clone is also
insufficient when sorting code mutates fields through pointer-like elements;
copy those nested values when isolation requires it.

#### Common mistakes and edge cases

> **Warning:** `sorted := values` does not copy slice elements. With
> `values=[3,1,2]`, sorting `sorted` also changes `values` to `[1,2,3]`; use
> `slices.Clone` when the old order must survive.

Duplicates require a tie rule only when the relative order of equal primary
keys matters; choose a **stable sort**, which preserves the earlier order of
equal keys, for that contract. Nil and empty slices are already sorted. Tests
should cover both allowed in-place behavior and promised non-mutation
explicitly.

#### Related material

- [B.1 Arrays, slices, and strings](#b1-arrays-slices-and-strings) owns slice
  descriptors, backing arrays, capacity, and append behavior.
- [C.1.3.5 Put related values together (sorting and selection)](#c135-sorting-and-selection)
  owns the algorithm-selection reasons for sorting.
- [D.9 Slice storage and retained memory](#d9-slice-storage-and-retained-memory)
  explains how sharing an array also affects memory lifetime.

<a id="e4-map-presence-and-sets"></a>

<a id="e4-distinguish-missing-keys-from-stored-values-map-presence-and-sets"></a>
<a id="d4-map-presence-and-sets"></a>

### D.4 Distinguish missing keys from stored values (map presence and sets)

#### The problem it solves

A Go map lookup must sometimes distinguish "this key is absent" from "this key
is present and stores its type's default value." Programs also need to record
membership when no separate value is meaningful.

#### The idea in simple terms

Think of labeled mailboxes. Opening an unknown label gives you an empty-looking
result, but that alone cannot tell you whether the mailbox exists and is empty.
Go can return both the stored value and a presence flag. A collection that
records only presence is called a **set**; Go commonly represents one with a
map whose keys are the set members.

#### How it works

A map read in the form `value := entries[key]` returns the value type's **zero
value** when `key` is missing: for example, `0` for `int`, `false` for `bool`,
and `nil` for a pointer or slice. The two-result form adds the missing fact:
`value, found := entries[key]`.

`found` is `true` only when the key exists. For a set, use
`map[T]struct{}` and store `struct{}{}` for each member; the empty structure
has no value payload because membership is carried by the key. `T` must be a
**comparable** type, meaning Go permits two values of that type to be compared
with `==` and used as map keys. Map iteration does not promise an order, so
preserve insertion order separately or collect and sort keys when output order
matters.

#### Step-by-step examples

Start with `entries = {"ready": 0}` and query `"ready"` and `"missing"`.
The expected lookup results are `(0, true)` and `(0, false)`.

1. **Step: Query `"ready"`**
   - **Before:** `entries={"ready":0}`
   - **Decision:** The key exists and its stored value is `0`.
   - **After:** `value=0`, `found=true`; map unchanged
   - **Known correct:** The zero came from a present entry.
2. **Step: Query `"missing"`**
   - **Before:** `entries={"ready":0}`
   - **Decision:** The key does not exist, so Go supplies the integer zero value.
   - **After:** `value=0`, `found=false`; map unchanged
   - **Known correct:** The same value is now identified as a missing-key result.

For a set example, process `values = ["go", "go", "rust"]` and keep first
appearances. The expected result is `["go", "rust"]`.

1. **Value: `"go"`**
   - **Before:** `seen={}`, `result=[]`
   - **Decision:** `"go"` is absent; add and append it.
   - **After:** `seen={"go"}`, `result=["go"]`
2. **Value: `"go"`**
   - **Before:** `seen={"go"}`, `result=["go"]`
   - **Decision:** `"go"` is present; skip the duplicate.
   - **After:** State unchanged.
3. **Value: `"rust"`**
   - **Before:** `seen={"go"}`, `result=["go"]`
   - **Decision:** `"rust"` is absent; add and append it.
   - **After:** `seen={"go","rust"}`, `result=["go","rust"]`

The result order comes from scanning `values`, not from iterating over `seen`.

#### Why it works

The two-result lookup separates two facts that the value alone cannot encode:
what value a key maps to and whether the mapping exists. A set uses exactly the
second fact. Before each value in the set trace, `seen` contains every distinct
earlier value and `result` lists those values in first-appearance order. A
present value is safely skipped; an absent value is added to both structures,
so the fact remains true. After the finite input is exhausted, `result`
contains each input value exactly once in the promised order.

#### Go example

```go
func lookup(entries map[string]int, key string) (int, bool) {
	// The comma-ok form distinguishes a missing key from a present key storing 0.
	value, found := entries[key]
	return value, found
}

func uniqueInOrder(values []string) []string {
	// struct{} is a zero-size membership marker; result preserves encounter order.
	seen := make(map[string]struct{})
	result := make([]string, 0, len(values))
	for _, value := range values {
		if _, found := seen[value]; found {
			// A prior occurrence already occupies this value's output position.
			continue
		}
		seen[value] = struct{}{}
		result = append(result, value)
	}
	return result
}
```

#### Work and memory

Go's language rules do not promise a fixed time bound for map operations. In
the usual hash-map cost model--where a computed key fingerprint selects a small
storage area--a lookup or insertion takes expected constant time, written
`O(1)`. Under that model, processing `n` values takes expected `O(n)` time. If
there are `u` distinct values, `seen` stores `u` keys. The result has length
`u`, but `make(..., 0, len(values))` reserves room for `n` strings to avoid
growth allocations, so the function's total additional storage is `O(n)`.
`struct{}` avoids a per-entry value payload, but map keys and bucket storage
still use memory. Collecting and sorting `u` keys for ordered output adds
`O(u log u)` work and `O(u)` key storage.

#### When to use it

Use the two-result lookup when a zero value is valid data, when membership
changes control flow, or when a missing key must produce a different result.
Use a set when the question is only whether a comparable key has appeared or
belongs to a collection.

#### When not to use it

Do not rely on a map when iteration order is part of the result. For a small,
dense integer range, a boolean slice may express membership with less overhead.
A plain set also cannot answer ordering or range questions without another
structure.

#### Common mistakes and edge cases

> **Warning:** `entries[key] == 0` is not an absence test. With
> `entries={"ready":0}`, it incorrectly treats the present key as missing; use
> the `found` result.

> **Warning:** Reading a nil map is allowed, but writing to it panics. Declare
> `var seen map[string]struct{}` and then execute
> `seen["go"] = struct{}{}` to reproduce the failure; call `make` first.

A `map[T]bool` can represent a set, but a stored `false` is indistinguishable
from absence unless the code also reads `found`. Duplicate input must not be
appended twice, and tests must not expect `range` over a map to return a stable
order.

#### Related material

- [C.1.3.4 Hash map or set](#c134-hash-map-or-set) owns the algorithm-selection
  clues and invariants for lookup-based methods.
- [E.3 Hash-map complement lookup](#e3-hash-map-complement-lookup) shows map
  presence in a complete reference algorithm.
- [A.2 Correctness and invariants](#a2-correctness-and-invariants) explains how
  the `seen` and `result` relationship proves duplicate removal correct.

<a id="e5-priority-queues"></a>

### D.5 Priority queues

#### The problem it solves

A program may keep receiving items but repeatedly need only the smallest,
largest, or most urgent remaining one. Sorting the whole collection after each
insertion repeats work, while scanning every item for each removal is slow.

#### The idea in simple terms

Keep the next item near the front while allowing the rest to be only partly
ordered. A **priority queue** removes items by priority rather than arrival
time. Go's `container/heap` package implements one with a **heap**, an array
whose parent-child ordering makes the highest-priority item available at index
zero.

#### How it works

1. Define `Len`, `Less`, and `Swap` on a slice type. `Less(i, j)` must report
   whether item `i` should come out before item `j`.
2. Define pointer-receiver `Push` and `Pop` methods so the package can grow and
   shrink the slice.
3. Call `heap.Init` when starting from existing values.
4. Use `heap.Push` and `heap.Pop`--not the methods directly--so the package
   repairs the parent-child ordering after each change.
5. Check `Len` before reading index zero or removing an item.

If `Less` uses `<`, the smallest item comes out first; reversing the comparison
creates a largest-first queue.

#### Step-by-step examples

Push `4`, `1`, and `3` into an empty smallest-first queue, then remove one
item. The expected removed value is `1` and the remaining heap array is
`[3,4]`.

1. **Operation: Push `4`**
   - **Before:** `h=[]`.
   - **Decision:** Append `4`; it has no parent to compare with.
   - **After:** `h=[4]`.
   - **Known correct:** Index zero contains the only and therefore smallest
     item.
2. **Operation: Push `1`**
   - **Before:** `h=[4]`.
   - **Decision:** Append `1`, then swap it with parent `4` because `1 < 4`.
   - **After:** `h=[1,4]`.
   - **Known correct:** Every child is at least its parent, so `1` is at the
     root.
3. **Operation: Push `3`**
   - **Before:** `h=[1,4]`.
   - **Decision:** Append `3`. Its parent is `1`, so no swap is needed.
   - **After:** `h=[1,4,3]`.
   - **Known correct:** The root is still the smallest item.
4. **Operation: Pop**
   - **Before:** `h=[1,4,3]`.
   - **Decision:** The package moves root `1` to the end, repairs the remaining
     heap, and asks the type's `Pop` method to shorten the slice.
   - **After:** Returned value `1`; `h=[3,4]`.
   - **Known correct:** `1` was no larger than any queued item, and `3` is now
     the smallest remaining item.

Some algorithms push a newer record instead of editing an older queued record.
For the boundary example, Dijkstra's algorithm first queues route `B=4`, later
finds `B=3`, and queues that too. The expected behavior is to use `B=3` and
eventually skip `B=4`.

1. **Record: `B=3` reaches the root**
   - **Before:** Stored best distance for `B` is `3`.
   - **Decision:** The queued distance matches the stored best distance.
   - **After:** Process `B=3`.
   - **Known correct:** The record represents the current best known route.
2. **Record: `B=4` reaches the root**
   - **Before:** Stored best distance for `B` remains `3`.
   - **Decision:** `4 != 3`, so the old record is **stale**, meaning a newer
     record has replaced it.
   - **After:** Discard `B=4` without processing its outgoing edges.
   - **Known correct:** Skipping an outdated record cannot discard the stored
     shorter route.

#### Why it works

The fact that remains true is that every parent has at least as much priority
as its children. Push repairs that fact while moving the new item toward the
root; pop repairs it while moving the replacement root downward. Consequently,
no descendant can outrank index zero, so removing the root returns the required
item. Each repair moves through a finite number of array levels.

#### Go example

```go
type intMinHeap []int

// Len, Less, and Swap let heap.Interface maintain this slice as a min-heap.
func (h intMinHeap) Len() int           { return len(h) }
func (h intMinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *intMinHeap) Push(value any) {
	// heap.Interface passes values as any. The .(int) assertion recovers the
	// required element type and would panic if a caller supplied a non-int value.
	// The pointer receiver then grows the heap's backing slice.
	*h = append(*h, value.(int))
}

func (h *intMinHeap) Pop() any {
	// container/heap first swaps the root with the final slot; Pop removes and
	// returns that final slot rather than searching for the minimum itself.
	old := *h
	last := len(old) - 1
	value := old[last]
	*h = old[:last]
	return value
}

func popAscending(values []int) []int {
	// Clone protects the caller, and Init rearranges the copy into heap order.
	h := intMinHeap(slices.Clone(values))
	heap.Init(&h)

	// Every heap pop removes the current minimum, producing ascending output.
	result := make([]int, 0, len(values))
	for h.Len() > 0 {
		result = append(result, heap.Pop(&h).(int))
	}
	return result
}
```

The package moves the root to the final slot before it calls the type's `Pop`
method, so that method removes the last element. The type assertion `value.(int)`
is safe here because this queue's callers push only integers.

#### Work and memory

For `n` initial items, `heap.Init` takes `O(n)` time. Reading the root takes
`O(1)` time. Each push or pop repairs at most the heap height, so it takes
`O(log n)` time. The heap stores `O(n)` items. `popAscending` also clones the
input and constructs an `n`-element output, so its total additional memory is
`O(n)` and its total work is `O(n log n)`.

#### When to use it

Use a priority queue when items arrive or change over time and each step needs
the current smallest, largest, earliest, cheapest, or otherwise most urgent
item without fully ordering everything.

#### When not to use it

Do not use a heap when one final sorted order is the only goal; sorting once is
simpler. A heap is also the wrong tool when arbitrary items must be found or
removed quickly unless another structure tracks their positions.

#### Common mistakes and edge cases

> **Warning:** Calling `h.Pop()` directly skips the package's root swap and
> repair. With `h=[1,4,3]`, the method removes `3` rather than minimum `1`; call
> `heap.Pop(&h)`.

Popping or reading `h[0]` when `Len()==0` panics. Reversing `Less` changes
smallest-first to largest-first. For queues of pointers, clear the removed final
slot before shortening the slice so the backing array does not retain it; see
[D.9](#d9-slice-storage-and-retained-memory).

#### Related material

- [C.1.3.27 Remove the next most urgent item quickly](#c1327-heap-or-priority-queue)
  owns heap selection, correctness, and algorithmic variants.
- [C.1.3.23 Dijkstra shortest paths](#c1323-dijkstra-shortest-paths) explains
  why skipping stale shortest-path entries is safe.
- [D.9 Slice storage and retained memory](#d9-slice-storage-and-retained-memory)
  owns clearing pointer elements removed from long-lived heaps.

<a id="e6-bytes-runes-and-visible-characters"></a>

### D.6 Bytes, runes, and visible characters

#### The problem it solves

An algorithm that says "character" can mean a stored byte, a Unicode code
point, or one symbol a reader sees. Treating those units as interchangeable
causes wrong lengths, broken slices, and incorrect indexes.

#### The idea in simple terms

Text has layers. A Go string stores bytes. Unicode assigns a number called a
**code point** to each textual unit, and UTF-8 is the encoding that maps that
number to one or more bytes. Go calls one code point a **rune**. A user-visible
symbol can contain several runes, such as a letter followed by a combining
accent, so even a rune is not always a complete visible character.

#### How it works

1. Read the problem's definition of one unit before choosing an operation.
2. Use `len(text)` and byte indexes for bytes or a stated one-byte alphabet.
3. Use `for index, value := range text` to decode runes while keeping each
   rune's starting byte index.
4. Convert with `[]rune(text)` when the algorithm needs rune positions or
   reverse indexing.
5. Use a Unicode text-segmentation library when the contract requires
   user-visible symbols rather than code points.

#### Step-by-step examples

Use `text = "A\u00e9"`. The expected byte length is `3`, and the expected rune
sequence is `['A','\u00e9']`.

1. **Step: Inspect stored bytes**
   - **Before:** UTF-8 encodes `A` as byte `0x41` and `\u00e9` as bytes
     `0xC3 0xA9`.
   - **Decision:** `len` reports stored bytes, not runes.
   - **After:** `byteCount=3`.
   - **Known correct:** All three encoded bytes are counted once.
2. **Step: Range over runes**
   - **Before:** Decoding begins at byte index `0`.
   - **Decision:** The first byte is the complete rune `A`.
   - **After:** The first pair is `index=0, value='A'`; decoding continues at
     byte index `1`.
3. **Step: Decode the multibyte rune**
   - **Before:** Bytes `0xC3 0xA9` begin at index `1`.
   - **Decision:** UTF-8 combines them into rune `\u00e9`.
   - **After:** The second pair is `index=1, value='\u00e9'`; the rune sequence
     has length `2`.
   - **Known correct:** Byte positions and rune values both refer to the same
     original string without pretending that every rune is one byte.

For the boundary case, use `text = "e\u0301"`: `e` followed by a combining
acute accent. The expected rune count is `2` even though many displays show one
accented symbol.

1. **Step: Decode code points**
   - **Before:** The string contains the code points `U+0065` and `U+0301`.
   - **Decision:** `[]rune` preserves both code points separately.
   - **After:** `runes=['e','\u0301']` and `len(runes)=2`.
   - **Known correct:** Rune processing answers a code-point question, not a
     visible-symbol question.

#### Why it works

UTF-8 defines exactly how the stored bytes decode into code points. `range`
advances by the width of each decoded rune, so each valid code point is visited
once and its reported index is its true starting byte position. Converting to
`[]rune` records those decoded values in order. Visible-symbol grouping follows
additional Unicode rules, so a code-point loop cannot promise that grouping by
itself.

#### Go example

```go
type textUnits struct {
	// ByteCount describes storage length; Runes describes decoded Unicode code
	// points. Neither necessarily equals the number of visible grapheme clusters.
	ByteCount int
	Runes     []rune
}

func describeText(text string) textUnits {
	// len counts UTF-8 bytes, while []rune decodes the string and allocates a
	// separate rune slice.
	return textUnits{
		ByteCount: len(text),
		Runes:     []rune(text),
	}
}
```

For `"A\u00e9"` this returns byte count `3` and two runes. A `range` loop is
often better when the algorithm can stream through runes without allocating a
slice.

#### Work and memory

`len(text)` reads the stored byte length in `O(1)` time. Ranging over `b` bytes
takes `O(b)` time and constant auxiliary memory. Building `[]rune(text)` also
takes `O(b)` time and allocates `O(r)` rune storage, where `r` is the decoded
rune count. Visible-symbol segmentation has its own library-specific state and
costs.

#### When to use it

Use bytes for binary data, ASCII-only contracts, or protocols defined in
bytes. Use runes when comparisons, positions, or counts are defined by Unicode
code points. Use visible-symbol segmentation for cursor movement or limits that
must match what a reader sees.

#### When not to use it

Do not convert every string to runes when the contract guarantees a one-byte
alphabet or the algorithm only copies bytes. Do not use rune indexes as byte
slice boundaries without converting them back to byte positions.

#### Common mistakes and edge cases

> **Warning:** `text[1:]` can begin inside a multibyte rune. With
> `text="\u00e9"`, byte index `1` starts at `0xA9` and the result is not valid
> UTF-8; slice only at known byte boundaries.

Invalid UTF-8 causes `range` to produce `utf8.RuneError` for invalid encoding
units. An empty string has zero bytes and zero runes. Normalized-looking text
may have different code-point sequences, as `"\u00e9"` and `"e\u0301"`
demonstrate, so equality rules must match the problem contract.

#### Related material

- [B.1 Arrays, slices, and strings](#b1-arrays-slices-and-strings) owns string
  immutability and basic sequence storage.
- [C.1.3.7 Reuse a partial match](#c137-string-matching) explains how the
  chosen text unit affects string-matching positions.
- [C.1.3.33 Share stored word beginnings](#c1333-trie) shows a byte-based trie
  whose alphabet contract must be explicit.

<a id="e7-type-parameters"></a>

### D.7 Type parameters

#### The problem it solves

Several Go types may need the same operation, yet copying the implementation
for `int`, `string`, and every other type creates repeated code that can drift.
Using `any` everywhere removes useful compile-time information and requires
callers to recover concrete types.

#### The idea in simple terms

Write one recipe with a labeled blank for the value type, then let each caller
fill in that blank. The label is a **type parameter**. Code using type
parameters is often called **generic** because one checked implementation works
with a group of concrete types.

#### How it works

1. Declare a type parameter in brackets, such as `T any`.
2. Use `T` in parameters, results, fields, or local values.
3. Give `T` a **constraint**, which states which types and operations are
   allowed. `any` permits every type; `comparable` permits `==` and `!=`.
4. Let the compiler infer `T` from arguments when possible, or write it
   explicitly, as in `first[string](values)`.
5. Return typed values so callers do not need type assertions.

#### Step-by-step examples

Call `first([]int{8,13})` and `first([]string{"go","rust"})`. The expected
results are `(8,true)` and `("go",true)`.

1. **Call: integer slice**
   - **Before:** The argument type is `[]int`.
   - **Decision:** The compiler infers `T=int`; `len(values)>0`.
   - **After:** The function returns the typed integer `8` and `true`.
   - **Known correct:** No string or other type can appear in this call's
     result.
2. **Call: string slice**
   - **Before:** The argument type is `[]string`.
   - **Decision:** The compiler infers `T=string`; `len(values)>0`.
   - **After:** The function returns the typed string `"go"` and `true`.
   - **Known correct:** The same control flow is checked for a different value
     type.

For the boundary case, call `first([]int{})`. The expected result is
`(0,false)`.

1. **Call: empty integer slice**
   - **Before:** `T=int` and `len(values)=0`.
   - **Decision:** There is no first element, so declare `var zero T`.
   - **After:** Go gives `zero` the integer zero value `0` and the function
     returns `0,false`.
   - **Known correct:** The boolean distinguishes "no value" from a stored
     first value that happens to be zero.

#### Why it works

Each call chooses one concrete type that satisfies the declared constraint.
The compiler checks the generic body using only operations promised by that
constraint, then checks callers with their chosen type. Therefore the returned
`T` matches the input element type, and the absence boolean safely separates an
empty input from the type's zero value.

#### Go example

```go
func first[T any](values []T) (T, bool) {
	if len(values) == 0 {
		// A generic type has no universal literal, so declare its zero value and
		// pair it with false to report absence.
		var zero T
		return zero, false
	}
	// Indexing is safe after the length check; true distinguishes a stored zero.
	return values[0], true
}
```

`any` is an alias for `interface{}`, but `T` remains one specific type for each
call. Use a narrower constraint when the implementation needs comparison,
ordering, arithmetic, or methods.

#### Work and memory

`first` checks the length and reads at most one element, so every call takes
`O(1)` time and `O(1)` additional memory. Type parameters do not make this
operation scan or copy the slice.

#### When to use it

Use type parameters when the same data-structure or function behavior applies
to multiple types and its required operations can be expressed by a small,
clear constraint.

#### When not to use it

Do not add a type parameter when only one concrete type is meaningful or when
different types require different behavior. Use an **interface value**, which
can hold any concrete type that provides the required methods, when a runtime
collection must mix such types.

#### Common mistakes and edge cases

> **Warning:** `left == right` does not compile when `T` is constrained by
> `any` because slices, maps, and functions are not comparable. Use
> `T comparable` only when equality is required.

An empty input needs a second result or another explicit absence convention
because every `T` has a zero value. Avoid constraints that expose operations
the function never uses; they reject valid callers without improving safety.

#### Related material

- [E.4 Generic singly linked list](#e4-generic-singly-linked-list) applies a
  type parameter to nodes.
- [E.5 Generic stack](#e5-generic-stack) applies one to a reusable container.
- [D.4 Map presence and sets](#d4-map-presence-and-sets) explains why generic
  map keys need the `comparable` constraint.

<a id="e8-replacing-recursion-with-a-loop"></a>

### D.8 Replacing recursion with a loop

#### The problem it solves

A recursive function relies on Go's call stack to remember unfinished work.
Very deep input can exhaust that stack, and some workflows need to pause,
inspect, or control the pending work directly.

#### The idea in simple terms

Replace the hidden pile of unfinished calls with a slice the program manages.
That slice is an **explicit stack**. Each stored item must contain everything a
recursive call would need when its turn resumes: the current value, the next
choice, and any partial result used after a child finishes.

#### How it works

For preorder tree traversal, where a node is visited before its children:

1. Return an empty result for a nil root; otherwise push the root.
2. Pop the final stack item and append its value to the result.
3. Push its right child first, then its left child.
4. Repeat until the stack is empty.

Right is pushed first because a stack removes the newest item first, so the
left child will be processed next. More complicated recursion needs a frame
struct that also records which child or choice comes next.

#### Step-by-step examples

Traverse a tree with root `8`, left child `4`, and right child `12`. Write the
stack with its top on the right. The expected preorder result is `[8,4,12]`.

1. **Step: Start**
   - **Before:** `stack=[]`, `result=[]`.
   - **Decision:** The root is non-nil, so push `8`.
   - **After:** `stack=[8]`, `result=[]`.
   - **Known correct:** Every unvisited reachable node is either the root on the
     stack or below it.
2. **Step: Visit `8`**
   - **Before:** `stack=[8]`, `result=[]`.
   - **Decision:** Pop and append `8`; push right `12` before left `4`.
   - **After:** `stack=[12,4]`, `result=[8]`.
   - **Known correct:** `8` is the required first value, and `4` is next.
3. **Step: Visit `4`**
   - **Before:** `stack=[12,4]`, `result=[8]`.
   - **Decision:** Pop and append leaf `4`; it has no children to push.
   - **After:** `stack=[12]`, `result=[8,4]`.
   - **Known correct:** The root and complete left subtree are in preorder.
4. **Step: Visit `12`**
   - **Before:** `stack=[12]`, `result=[8,4]`.
   - **Decision:** Pop and append leaf `12`.
   - **After:** `stack=[]`, `result=[8,4,12]`.
   - **Known correct:** No pending node remains, and the full preorder is
     complete.

For the boundary case, use `root=nil`. The expected result is an empty slice.

1. **Step: Check the root**
   - **Before:** `root=nil`, `stack=[]`, `result=[]`.
   - **Decision:** There is no call frame to simulate, so return immediately.
   - **After:** `result=[]`.
   - **Known correct:** An empty tree contains no node to visit.

#### Why it works

Before each loop, the stack contains exactly the roots of subtrees that the
recursive version has not yet visited, in reverse visit order. Popping visits
the same node the next recursive call would visit. Pushing right before left
preserves the left-before-right order. Each node is pushed and popped once, so
the finite stack eventually empties and the result matches recursive preorder.

#### Go example

```go
type traversalNode struct {
	Value int
	Left  *traversalNode
	Right *traversalNode
}

func preorderIterative(root *traversalNode) []int {
	// Avoid placing a nil pointer on the stack; an empty tree has empty preorder.
	if root == nil {
		return []int{}
	}

	// The explicit stack stores nodes whose subtrees still need processing.
	stack := []*traversalNode{root}
	result := make([]int, 0)
	for len(stack) > 0 {
		// Pop the final element to obtain last-in, first-out behavior.
		last := len(stack) - 1
		node := stack[last]
		stack = stack[:last]

		// Preorder visits the node before its children. Push right first so the
		// left child, placed on top last, is processed next.
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

This frame needs only a node pointer because preorder performs all work for a
node before scheduling its children. Postorder or backtracking may need a frame
with a next-step field and partial results.

#### Work and memory

For `n` nodes, each node is pushed, popped, and appended once, so traversal
takes `O(n)` time. The result uses `O(n)` output memory. The explicit stack uses
`O(h)` auxiliary memory for this binary-tree traversal, where `h` is the tree
height; in the worst case `h=n`. The recursive form uses a call stack with the
same asymptotic height bound.

#### When to use it

Use an explicit stack when depth may be too large for safe recursion, when the
pending work must be paused or inspected, or when an iterative interface is
required.

#### When not to use it

Do not replace short, naturally bounded recursion when the manual frames make
the invariant harder to see. A simple node-only stack is not enough when work
must happen after a child returns; encode that return state explicitly.

#### Common mistakes and edge cases

> **Warning:** Pushing left before right reverses preorder. For root `8` with
> left child `4` and right child `12`, it produces `[8,12,4]`; push right first
> so left is popped first.

Handle a nil root before pushing it, or dereferencing the popped nil pointer
will panic. For graph traversal, track visited nodes separately or a cycle can
keep adding work forever.

#### Related material

- [A.3 Recursion and saved work](#a3-recursion-and-saved-work) owns recursive
  call frames and stack-memory accounting.
- [B.4.2 Traversal order](#b42-traversal-order) owns preorder, inorder, and
  postorder meaning.
- [C.1.3.20 Depth-first search](#c1320-depth-first-search) explains the broader
  search method that can use either call stack.

<a id="e9-slice-storage-and-retained-memory"></a>

### D.9 Slice storage and retained memory

#### The problem it solves

A short slice can keep a much larger backing array reachable, and a shortened
pointer slice can leave removed objects referenced in unused slots. The visible
length looks small even while memory associated with old data remains live.

#### The idea in simple terms

A slice is a window onto an array, not the array itself. Narrowing the window
does not cut away the room behind it. When a small window must live much longer
than its source, copy it into a right-sized array. When removing a pointer from
a long-lived array, erase the old slot before hiding it beyond the slice length.

#### How it works

1. Remember that a slice stores a pointer to an array, a length, and a capacity.
2. Reslicing changes length and perhaps capacity but can keep the same array.
3. To detach a small surviving portion, allocate exactly its length and use
   `copy`.
4. To remove the last pointer from a reusable slice, save the value, assign
   `nil` to that slot, and then shorten the slice.
5. Treat `append` as conditional: it may reuse shared capacity or allocate a
   new array.

#### Step-by-step examples

Suppose `large` refers to a 1,048,576-byte array and `header := large[:16]`
must be kept in a cache after `large` is no longer needed. The expected compact
result contains the same 16 bytes without retaining the large array through
`header`.

1. **Step: Take a small view**
   - **Before:** `large` reaches the 1,048,576-byte backing array.
   - **Decision:** `large[:16]` changes the visible length but shares the array.
   - **After:** `header` shows 16 bytes and still points into the large array.
   - **Known correct:** The byte values are right, but storage is not detached.
2. **Step: Copy the surviving bytes**
   - **Before:** `header` shares the 1,048,576-byte array.
   - **Decision:** Allocate `len(header)=16` bytes and copy the visible data.
   - **After:** `compact` has the same 16 bytes in an independent array with
     length and capacity `16`.
   - **Known correct:** Keeping `compact` no longer requires the old large array
     through this slice.

For a pointer-removal example, let `items=[A,B,C]` and pop `C`. The expected
visible result is `[A,B]`, and the old final array slot should be `nil`.

1. **Step: Save the removed pointer**
   - **Before:** `last=2` and the backing-array slot holds `C`.
   - **Decision:** Save `items[2]` so it can be returned.
   - **After:** `value=C`; the slice and backing array are otherwise unchanged.
2. **Step: Clear and shorten**
   - **Before:** The final slot still references `C`.
   - **Decision:** Assign `nil` to `items[2]`, then reslice to `items[:2]`.
   - **After:** Visible `items=[A,B]`; the reusable final slot is nil.
   - **Known correct:** The returned value is preserved, while the backing array
     no longer retains `C` through the removed slot.

#### Why it works

`copy` writes each visible element into a separately allocated array, so the
new slice cannot keep the source array reachable. Clearing a removed pointer
replaces the backing array's reference before reslicing hides the slot. The
visible sequence remains correct because copying preserves order and popping
returns the saved final value.

#### Go example

```go
func compactBytes(view []byte) []byte {
	// Allocate exactly enough independent storage, then copy the visible bytes.
	// The returned slice cannot retain or mutate view's larger backing array.
	compact := make([]byte, len(view))
	copy(compact, view)
	return compact
}

func popPointer[T any](items []*T) ([]*T, *T, bool) {
	// Empty input has no final slot to read or clear.
	if len(items) == 0 {
		return items, nil, false
	}

	// Save the value before shrinking, then clear the pointer slot so the backing
	// array does not unnecessarily keep the referenced object reachable.
	last := len(items) - 1
	value := items[last]
	items[last] = nil
	// Return the shortened descriptor separately; reslicing this local variable
	// cannot change the caller's slice length by itself.
	return items[:last], value, true
}
```

`make` gives `compact` exactly the requested length and capacity here.
`popPointer` returns the shortened slice because changing a local slice
descriptor does not change the caller's descriptor.

#### Work and memory

Copying a `k`-element view takes `O(k)` time and allocates `O(k)` memory. Popping
and clearing the final pointer takes `O(1)` time and no new slice storage. The
copy deliberately trades a small allocation now for allowing a much larger
array to become unreachable sooner.

#### When to use it

Compact a slice when a small portion will outlive a much larger input. Clear
removed pointer-bearing slots in long-lived stacks, queues, heaps, buffers, or
pools whose backing arrays remain reachable.

#### When not to use it

Do not copy every short-lived subslice; the allocation can cost more than brief
retention. Clearing scalar slots such as integers usually does not affect
object reachability, though it may still be required by a data-structure
invariant.

#### Common mistakes and edge cases

> **Warning:** `items = items[:last]` hides but does not overwrite the removed
> pointer. With backing storage `[A,B,C]`, the unused slot can still retain `C`;
> assign `items[last] = nil` first.

A copied outer slice still shares objects reached through pointer-like
elements. Appending to one of two shared slices may mutate the common array
when capacity remains and may detach when it does not, so correctness must not
depend on guessing which outcome occurs. Empty and nil slices require no
compaction; `popPointer` reports `false` instead of indexing them.

#### Related material

- [B.1 Arrays, slices, and strings](#b1-arrays-slices-and-strings) owns slice
  length, capacity, shared arrays, and append behavior.
- [D.3 Sorting without unexpected mutation](#d3-sorting-without-unexpected-mutation)
  uses cloning to protect observable order.
- [D.5 Priority queues](#d5-priority-queues) is one long-lived slice-backed
  structure that may hold pointers.

<a id="d-go-reference-implementations"></a>

## E. Go reference implementations

These snippets are complete, small implementations of ideas taught elsewhere
in the guide. Each entry states its contract and required input conditions,
traces the variables used by the code, and gives concrete cases to test. The
linked owner explains how to recognize and choose the method; this section
keeps the implementation details beside runnable Go code.

Some snippets use a type parameter written as `T`. Treat `T` as a placeholder
for the value type chosen by the caller, such as `int` or `string`; the
[type-parameter note](#d7-type-parameters) explains why this is useful.

<a id="d1-binary-search-bounds"></a>

### E.1 Binary search bounds

#### The problem it solves

Given values sorted in ascending order, find where a target begins, where it
ends, or whether it occurs at all. Equal values make an ordinary "find any
match" search insufficient when the caller needs exact range boundaries.

#### The idea in simple terms

Keep a half-open candidate range: `left` is included and `right` is excluded.
Inspect its middle value and discard the half that cannot contain the first
valid boundary. This is **binary search**; a **lower bound** is the first value
greater than or equal to the target, while an **upper bound** is the first
value strictly greater than it.

#### How it works

The input must be sorted in ascending order.

1. Start with `[left,right) = [0,len(values))`.
2. For a lower bound, move `left` past `middle` when
   `values[middle] < target`; otherwise keep `middle` by moving `right` to it.
3. For an upper bound, use `values[middle] <= target` as the condition for
   moving `left` past `middle`.
4. When `left == right`, return that insertion position. For exact search,
   confirm that the lower-bound position exists and contains the target.

#### Step-by-step examples

For `values=[1,2,2,2,5]` and `target=2`, `LowerBound` should return `1`.

1. **Step: Inspect index 2**
   - **Before:** `left=0`, `right=5`, `middle=2`, value `2`.
   - **Decision:** `2` is not smaller than the target, so index 2 might be the
     first match; set `right=2`.
   - **After:** The candidate range is `[0,2)`.
   - **Known correct:** No index from 2 onward can be earlier than a valid
     boundary still in the range.
2. **Step: Inspect index 1**
   - **Before:** `left=0`, `right=2`, `middle=1`, value `2`.
   - **Decision:** Keep this possible first match by setting `right=1`.
   - **After:** The candidate range is `[0,1)`.
   - **Known correct:** Any lower bound must be at index 0 or 1.
3. **Step: Inspect index 0**
   - **Before:** `left=0`, `right=1`, `middle=0`, value `1`.
   - **Decision:** `1 < 2`, so index 0 cannot qualify; set `left=1`.
   - **After:** `left=right=1`; return `1`.
   - **Known correct:** Index 0 is too small and index 1 is the first value at
     least 2.

For the boundary case `values=[1,2,2,2,5]`, `target=6`, `LowerBound` returns
`5`. Every inspected value is smaller than 6, so `left` advances to
`len(values)`; `Search` checks that position before indexing and returns `-1`.

#### Why it works

For a lower bound, every index before `left` is known to contain a value
smaller than the target, and every in-range index at or after `right` is known
to contain a value at least as large as the target. Those facts remain true
after either update because the values are sorted. Each step makes the
candidate range shorter, so the loop finishes. At `left == right`, no earlier
index qualifies; that index is the first qualifying position, or it equals
`len(values)` when none exists. The upper-bound proof is identical with
"smaller than or equal to" as the discarded-left condition.

#### Go example

```go
package binarysearch

// LowerBound returns the first index whose value is greater than or equal to
// target. It returns len(values) when no such index exists.
func LowerBound(values []int, target int) int {
	// Only values in the half-open range [left,right) remain unclassified.
	// The boundaries converge on the insertion position, which may be len(values).
	left, right := 0, len(values)
	for left < right {
		// This form avoids adding the two potentially large indexes directly.
		middle := left + (right-left)/2
		if values[middle] < target {
			// middle and everything before it are too small.
			left = middle + 1
		} else {
			// middle may be the first qualifying index, so keep it in the range.
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
			// Upper bound requires a strictly greater value, so discard equality.
			left = middle + 1
		} else {
			// middle is greater and remains a possible first qualifying position.
			right = middle
		}
	}
	return left
}

// Search returns the index of target in a sorted slice, or -1 when target is
// absent. When duplicates exist, it returns the first matching index.
func Search(values []int, target int) int {
	// LowerBound gives the only position that could be target's first occurrence.
	index := LowerBound(values, target)
	// Check the len sentinel before indexing, then confirm exact equality.
	if index == len(values) || values[index] != target {
		return -1
	}
	return index
}
```

Concrete cases to test include empty input; targets before, within, between,
and after the data; `values=[1,2,2,2,5], target=2` producing lower bound `1`,
upper bound `4`, and search result `1`; and an absent target producing `-1`.

#### Work and memory

Each update removes at least half of the remaining candidate positions, so a
nonempty input takes `O(log n)` time and an empty input takes `O(1)`. `Search`
adds one constant-time check after `LowerBound`. The functions keep only index
variables and use `O(1)` auxiliary memory without allocations.

#### When to use it

Use these functions when values are already sorted and the caller needs an
insertion point, an exact first match, a duplicate range, or a count such as
`UpperBound(values,x)-LowerBound(values,x)`.

#### When not to use it

Do not use the comparisons on unsorted input: for `values=[3,1,2]`, discarding
one side says nothing about values on that side. A map or linear scan is often
better for repeated exact lookups when sorting is unavailable or forbidden.

#### Common mistakes and edge cases

> **Warning:** Changing the range to inclusive bounds without changing the
> updates can skip an index or loop forever. With this half-open contract,
> initialize `right=len(values)` and never read `values[right]`.

Check `index == len(values)` before reading `values[index]`. Keep the lower-
and upper-bound equality rules distinct, especially for duplicates. Computing
`middle` as `left+(right-left)/2` also avoids overflowing `left+right` in
integer types where very large indexes are possible.

#### Related material

- [C.1.3.6 Find a changing boundary (binary search)](#binary-search)
  owns method selection and the boundary invariant.
- [A.2 Correctness and invariants](#a2-correctness-and-invariants) explains
  how the maintained facts prove the result.
- [Find First and Last Position of Element in Sorted Array](problems/solved_medium_leetcode_0034_find_first_and_last_position_of_element_in_sorted_array.go)
  uses both bounds around a duplicate range.

<a id="d2-rolling-dynamic-programming"></a>

### E.2 Rolling dynamic programming

#### The problem it solves

Count ordered sequences of one-step and two-step moves that reach an exact
nonnegative total. Recomputing every smaller total recursively repeats the
same work, while storing every answer uses more memory than this recurrence
needs.

#### The idea in simple terms

Every final sequence reaches `current` from either `current-1` or
`current-2`. Add those two already known counts, then keep only the newest two
counts for the next total. Reusing answers to smaller problems is **dynamic
programming**; retaining only the recent states is a **rolling** form of it.

#### How it works

The result must fit in `int`.

1. Return zero for a negative total because no sequence of positive moves can
   reach it.
2. Return one for totals zero and one: the empty sequence reaches zero, and
   one one-step move reaches one.
3. Let `twoBack` and `oneBack` hold the counts for totals `current-2` and
   `current-1`.
4. Set the new count to their sum and roll both variables forward until
   `current == steps`.

#### Step-by-step examples

For `steps=4`, the expected answer is `5`.

1. **Step: Establish totals 0 and 1**
   - **Before:** No loop iteration has run.
   - **Decision:** There is one way to make 0 and one way to make 1.
   - **After:** `twoBack=1`, `oneBack=1`.
   - **Known correct:** The two base totals have exact counts.
2. **Step: Build total 2**
   - **Before:** `current=2`, counts are `1` and `1`.
   - **Decision:** Append a two-step move to the total-0 sequences or a
     one-step move to the total-1 sequences.
   - **After:** `twoBack=1`, `oneBack=2`.
   - **Known correct:** The two sequences are `[2]` and `[1,1]`.
3. **Step: Build total 3**
   - **Before:** `current=3`, counts for totals 1 and 2 are `1` and `2`.
   - **Decision:** Add them.
   - **After:** `twoBack=2`, `oneBack=3`.
   - **Known correct:** All sequences reaching 3 end in exactly one of the two
     allowed move sizes.
4. **Step: Build total 4**
   - **Before:** `current=4`, counts are `2` and `3`.
   - **Decision:** Add them and roll forward.
   - **After:** `twoBack=3`, `oneBack=5`; return `5`.
   - **Known correct:** All five ordered sequences are counted once.

For boundary inputs, `steps=-1` returns `0` and `steps=0` returns `1` without
entering the loop. The second result counts the one empty sequence, not one
physical move.

#### Why it works

Before each iteration, `twoBack` equals the answer for `current-2` and
`oneBack` equals the answer for `current-1`. Every sequence reaching
`current` has a unique final move of size one or two; removing that move puts
it in exactly one of those two groups. Their sum is therefore the new exact
count. The assignment restores the same fact for the next total, and
`current` increases until it reaches `steps`.

#### Go example

```go
package dynamicprogramming

// CountOrderedOneOrTwoStepWays returns the number of ordered sequences of
// one-step and two-step moves that sum to steps. Negative totals have no valid
// sequence; the empty sequence is the single way to make a total of zero. The
// result must fit in an int.
func CountOrderedOneOrTwoStepWays(steps int) int {
	// Positive moves cannot form a negative total.
	if steps < 0 {
		return 0
	}
	if steps < 2 {
		// There is one empty sequence for zero and one [1] sequence for one.
		return 1
	}

	// Before each iteration these hold the counts for current-2 and current-1.
	twoBack, oneBack := 1, 1
	for current := 2; current <= steps; current++ {
		// Every sequence ends in either a 1-step or 2-step move. Multiple
		// assignment evaluates both old values before rolling them forward.
		twoBack, oneBack = oneBack, twoBack+oneBack
	}
	return oneBack
}
```

Concrete cases to test are `-1 -> 0`, `0 -> 1`, `1 -> 1`, `2 -> 2`, and
`4 -> 5`, plus the largest accepted `steps` whose answer still fits in `int`.

#### Work and memory

Negative, zero, and one return in `O(1)` time. For larger `steps`, the loop
runs `steps-1` times, so time is `O(steps)`. Two counts and one loop index use
`O(1)` auxiliary memory, and the function makes no allocations. This removes
the `O(steps)` table allocation but not the time needed to build each answer.

#### When to use it

Use rolling state when a recurrence proceeds in one direction, the output is
only the final value, and each new answer depends on a fixed number of recent
answers.

#### When not to use it

Keep a full table when reconstruction or later queries need older states. Do
not use this two-state recurrence when allowed move sizes or other state make
the next answer depend on more information.

#### Common mistakes and edge cases

> **Warning:** Treating total zero as having zero ways breaks later answers:
> `steps=2` would miss the sequence `[2]`. The empty sequence is the base way
> to make a zero total.

State the overflow limit; Go integer addition does not report overflow. In a
multiple assignment, right-hand expressions use the old values, so
`twoBack, oneBack = oneBack, twoBack+oneBack` performs the intended roll.

#### Related material

- [C.1.3.10 Reuse answers to repeated smaller problems (dynamic programming)](#dynamic-programming)
  owns recurrence design and state selection.
- [A.3 Recursion and saved work](#a3-recursion-and-saved-work) compares saved
  answers with repeated recursive calls.
- [Climbing Stairs](problems/solved_easy_leetcode_0070_climbing_stairs.go)
  uses this recurrence.

<a id="d3-hash-map-complement-lookup"></a>

### E.3 Hash-map complement lookup

#### The problem it solves

Find two distinct positions whose integer values add to a target without
checking every pair. The returned positions must refer to values already in
the input, and one position cannot be used twice.

#### The idea in simple terms

For each value, calculate the partner it still needs. Ask a map whether that
partner appeared earlier, then remember the current value for later items. The
needed partner is the **complement**, `target-value`; a Go map provides the
value-to-earlier-index lookup.

#### How it works

Every subtraction `target-value` must fit in `int`.

1. Start with an empty `indexByValue` map.
2. For each `index,value`, calculate `complement`.
3. Look up `complement`. If present, return its earlier index and the current
   index.
4. Otherwise store the current value and index. Insert only after the lookup
   so the current position cannot match itself.
5. Return `false` after the scan when no pair exists.

#### Step-by-step examples

For `values=[2,7,11,15]` and `target=9`, the expected result is `(0,1,true)`.

1. **Step: Visit index 0**
   - **Before:** `indexByValue={}`, `value=2`.
   - **Decision:** `complement=7` is absent, so no earlier partner exists.
   - **After:** `indexByValue={2:0}`.
   - **Known correct:** No pair ending at index 0 exists.
2. **Step: Visit index 1**
   - **Before:** `indexByValue={2:0}`, `value=7`.
   - **Decision:** `complement=2` is present at index 0.
   - **After:** Return `(0,1,true)` without further mutation.
   - **Known correct:** The positions are distinct and `2+7=9`.

For the duplicate case `values=[3,3]`, `target=6`, index 0 is stored before
index 1 looks up its complement. The result is `(0,1,true)`; looking up before
inserting is what prevents index 0 from matching itself on its first visit.

#### Why it works

Before visiting `index`, the map contains indexes from earlier positions only.
If the complement is present, its value plus the current value equals the
target and its index is necessarily distinct. If it is absent, no earlier
position can pair with the current value, so remembering the current value is
safe. The finite scan either returns a valid pair or proves that no pair has a
later endpoint in the input.

#### Go example

```go
// Package hashmap demonstrates lookup-based uses of Go maps.
package hashmap

// PairSumIndices returns indices of two distinct values whose sum is target.
// The boolean is false when no pair exists. Subtracting each value from target
// must fit in an int.
func PairSumIndices(values []int, target int) (int, int, bool) {
	// The map contains only values from earlier indices.
	indexByValue := make(map[int]int)
	for index, value := range values {
		// complement is the exact value needed to complete target.
		complement := target - value
		earlierIndex, found := indexByValue[complement]
		if found {
			// earlierIndex and index are necessarily distinct because insertion of
			// the current value happens below.
			return earlierIndex, index, true
		}
		// Insert after lookup so an index cannot match itself.
		indexByValue[value] = index
	}

	return 0, 0, false
}
```

Concrete cases to test include `[2,7,11,15],9 -> (0,1,true)`,
`[3,3],6 -> (0,1,true)`, and `[1,2,4],8 -> false`. For the false result, the
two returned integers are placeholders and must not be used.

#### Work and memory

With ordinary Go map behavior, each lookup and insertion takes constant time
on average, so the scan takes `O(n)` average time and can take `O(n^2)` in the
map's pathological worst case. The map stores at most `n` entries and uses
`O(n)` additional memory. A successful early pair can reduce actual work but
does not change the worst-case bound.

#### When to use it

Use complement lookup when original indices matter, one pass is useful, and
`O(n)` extra map memory is acceptable.

#### When not to use it

Use the sorted two-pointer implementation in [E.9](#e9-sorted-two-pointer-pair-search)
when the input is already sorted or may be reordered and constant auxiliary
memory matters. A frequency map is needed instead when the task asks for all
pairs or counts rather than one pair.

#### Common mistakes and edge cases

> **Warning:** Inserting before lookup lets `values=[3]`, `target=6` return the
> same index twice. Perform the lookup first.

Use the map's boolean presence result rather than treating index zero as
absence. Decide whether any valid pair is acceptable because repeated values
can produce several answers. Widen or validate arithmetic when
`target-value` may overflow `int`.

#### Related material

- [C.1.3.4 Remember facts by key (hash map or set)](#hash-map-or-set)
  owns the lookup method.
- [D.4 Distinguish missing keys from stored values (map presence and sets)](#d4-distinguish-missing-keys-from-stored-values-map-presence-and-sets)
  explains the two-result lookup used here.
- [Two Sum](problems/solved_easy_leetcode_0001_two_sum.go) is the
  representative problem.

<a id="d4-generic-singly-linked-list"></a>

### E.4 Generic singly linked list

#### The problem it solves

Build a one-way linked list from ordinary values, convert it back for
inspection, and reverse its links without allocating replacement nodes. The
same operations should work for different stored value types.

#### The idea in simple terms

Each node is a card holding a value and a pointer to the next card. A
**singly linked list** follows only those forward links. A type parameter `T`
lets the node store a caller-chosen type, and reversing the list turns each
link around while preserving the remaining path before it is overwritten.

#### How it works

1. `FromSlice` starts with a temporary `dummy` node and appends one new node
   after `tail` for every slice value; `dummy.Next` is the real head.
2. `ToSlice` follows `node.Next` until `nil`, appending each `node.Value`.
3. `Reverse` keeps `previous`, `current`, and `next`. Save `current.Next` in
   `next`, point `current.Next` backward, then advance both moving pointers.
4. A `nil` head represents an empty list and is valid for conversion and
   reversal.

#### Step-by-step examples

Build and reverse `values=[10,20,30]`; the expected final slice is
`[30,20,10]`.

1. **Step: Build the list**
   - **Before:** `dummy.Next=nil`, `tail=dummy`.
   - **Decision:** Append nodes 10, 20, and 30 after the current `tail`, moving
     `tail` to each new node.
   - **After:** `dummy -> 10 -> 20 -> 30 -> nil`; return `dummy.Next`.
   - **Known correct:** The links preserve the input slice order.
2. **Step: Reverse node 10**
   - **Before:** `previous=nil`, `current=10 -> 20`, `next` not yet saved.
   - **Decision:** Save node 20, then set `10.Next=nil`.
   - **After:** `previous=10 -> nil`, `current=20 -> 30`.
   - **Known correct:** Node 10 forms the complete reversed prefix, and node
     20 still reaches every unprocessed node.
3. **Step: Reverse node 20**
   - **Before:** `previous=10`, `current=20`, `next=30` after saving.
   - **Decision:** Set `20.Next=10` and advance.
   - **After:** `previous=20 -> 10`, `current=30`.
   - **Known correct:** The first two original nodes are reversed without
     losing node 30.
4. **Step: Reverse node 30 and convert**
   - **Before:** `previous=20 -> 10`, `current=30`.
   - **Decision:** Save `nil`, point 30 to 20, and advance to `nil`.
   - **After:** Return `previous=30 -> 20 -> 10`; `ToSlice` returns
     `[30,20,10]`.
   - **Known correct:** Every original node appears once in reverse order.

For an empty `values=[]`, `FromSlice` returns `nil`, `Reverse(nil)` returns
`nil`, and `ToSlice(nil)` returns an empty slice without entering either loop.

#### Why it works

During construction, the list after `dummy` contains exactly the values
already visited, in their slice order, and `tail` is its last node. During
reversal, `previous` heads the fully reversed processed prefix while `current`
heads the untouched suffix. Saving `next` before changing the link preserves
the suffix; redirecting one link extends the reversed prefix. Each iteration
processes one node, so the loop ends at `nil` with all nodes reversed.

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
	// A dummy node gives both empty and nonempty construction one append path;
	// tail always names the node after which the next value is attached.
	dummy := &Node[T]{}
	tail := dummy
	for _, value := range values {
		// Allocate one node, link it at the end, and advance tail to preserve order.
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
	// previous is the reversed prefix; current begins the unreversed remainder.
	var previous *Node[T]
	for current := head; current != nil; {
		// Save the only link to the remainder before redirecting current.Next.
		next := current.Next
		current.Next = previous
		previous = current
		current = next
	}
	// When current reaches nil, previous contains every node in reverse order.
	return previous
}
```

Concrete cases to test are the `[10,20,30]` round trip and reversal, one node
remaining unchanged, repeated values retaining their separate nodes, and all
three operations on an empty list.

#### Work and memory

For `n` values, `FromSlice` takes `O(n)` time and allocates `n` result nodes;
the temporary dummy may remain on the stack or be optimized by the compiler.
`ToSlice` takes `O(n)` time and `O(n)` result memory for the returned slice.
`Reverse` takes `O(n)` time and `O(1)` auxiliary memory and allocates no nodes.

#### When to use it

Use this reference when a problem supplies singly linked nodes, needs constant-
time link changes after a known node, or teaches pointer rewiring. The slice
adapters are useful at test and input boundaries.

#### When not to use it

Prefer a slice when indexed access, cache-friendly traversal, or compact
storage matters. This node type has no backward link, so do not use it when
frequent movement toward previous nodes is required.

#### Common mistakes and edge cases

> **Warning:** Setting `current.Next=previous` before saving the old next link
> loses the unprocessed suffix. For `10 -> 20 -> 30`, node 20 becomes
> unreachable after the first rewrite.

Return `dummy.Next`, not `dummy`. Treat `nil` as a valid empty list. A cycle
violates `ToSlice` and `Reverse` contracts because their loops would not reach
`nil`; detect cycles separately when input is not guaranteed acyclic.

#### Related material

- [B.2 Linked lists](#linked-lists) owns node operations and reversal
  reasoning.
- [D.7 Type parameters](#d7-type-parameters) explains `T` and `any`.
- [D.9 Slice storage and retained memory](#d9-slice-storage-and-retained-memory)
  contrasts contiguous slice storage with separately allocated nodes.

<a id="d5-generic-stack"></a>

### E.5 Generic stack

#### The problem it solves

Store values so that the most recently added value is the first one removed,
while allowing callers to inspect the top, count items, and handle an empty
stack without a panic.

#### The idea in simple terms

Use the end of a slice as the top of a stack of trays. Adding and removing at
that one end implements **last in, first out**, often shortened to LIFO. A
generic `Stack[T]` gives the same operations to any chosen value type.

#### How it works

1. `Push` appends the value to `items`.
2. `Peek` reads `items[len(items)-1]` only after checking that the slice is not
   empty.
3. `Pop` saves the last value, clears that slot to the zero value of `T`,
   shortens the slice by one, and returns the saved value with `true`.
4. Empty `Pop` and `Peek` return the zero value with `false`. `Len` and
   `IsEmpty` derive their answers from `len(items)`.

#### Step-by-step examples

Starting with the zero value `var stack Stack[string]`, push `"a"`, push
`"b"`, peek, then pop. The expected peek and pop value is `"b"`.

1. **Step: Push `"a"`**
   - **Before:** `items=nil`, length 0.
   - **Decision:** Append at the slice end.
   - **After:** `items=["a"]`, length 1.
   - **Known correct:** `"a"` is the newest and top value.
2. **Step: Push `"b"`**
   - **Before:** `items=["a"]`.
   - **Decision:** Append without disturbing the earlier item.
   - **After:** `items=["a","b"]`, length 2.
   - **Known correct:** `"b"` is now the top.
3. **Step: Peek**
   - **Before:** Last index is 1.
   - **Decision:** Read that slot without slicing it away.
   - **After:** Return `"b",true`; `items` remains `["a","b"]`.
   - **Known correct:** Inspection did not change stack order or length.
4. **Step: Pop**
   - **Before:** `last=1`, `value="b"`.
   - **Decision:** Clear slot 1, then shorten the visible slice.
   - **After:** Return `"b",true`; `items=["a"]`.
   - **Known correct:** The newest value was removed first.

For an empty stack, both `Pop` and `Peek` return `"",false` for
`Stack[string]`; the boolean distinguishes failure from a successfully stored
empty string.

#### Why it works

The slice order is the insertion order and the final element is always the
most recent unremoved value. `Push` adds exactly there, while `Peek` reads and
`Pop` removes exactly there, so no older value can leave first. Clearing a
removed slot does not change the logical order; it only stops that hidden
array position from retaining a reference. All operations finish after a
bounded number of slice and index operations, except for occasional append
growth.

#### Go example

```go
package stack

// Stack stores values in last-in, first-out order.
// The zero value is ready to use.
type Stack[T any] struct {
	items []T
}

// Push places value on top of the stack.
func (s *Stack[T]) Push(value T) {
	// The slice's final position is the stack top.
	s.items = append(s.items, value)
}

// Pop removes and returns the top value. The boolean is false when the stack
// is empty.
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		// Return T's zero value with false because there is no top element.
		var zero T
		return zero, false
	}

	last := len(s.items) - 1
	value := s.items[last]
	// Clear the hidden slot before shrinking so pointer-bearing T values do not
	// remain reachable through the backing array.
	var zero T
	s.items[last] = zero
	s.items = s.items[:last]
	return value, true
}

// Peek returns the top value without removing it. The boolean is false when
// the stack is empty.
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.items) == 0 {
		// As with Pop, false distinguishes absence from a stored zero value.
		var zero T
		return zero, false
	}

	// Read the final item without changing the slice length.
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

Concrete cases to test include the string sequence above, integer zero stored
successfully versus an empty result, length changes after every push and pop,
complete removal, and repeated empty operations on the zero-value stack.

#### Work and memory

`Peek`, `Pop`, `Len`, and `IsEmpty` take `O(1)` time. `Push` takes `O(1)`
amortized time: most appends are constant work, while an occasional capacity
growth copies `O(n)` items. Storing `n` values uses `O(n)` memory. The methods
use `O(1)` auxiliary memory beyond the slice storage.

#### When to use it

Use a stack when work must resume in reverse discovery order, such as nested
delimiter parsing, undo behavior, depth-first traversal, or replacing
recursion with explicit frames.

#### When not to use it

Use a queue when the oldest item must leave first. Use a priority queue when a
comparison, rather than insertion time, determines the next item. Do not expose
the backing slice when callers must not bypass stack ordering.

#### Common mistakes and edge cases

> **Warning:** Returning only the zero value for an empty pop is ambiguous.
> An integer stack can validly contain `0`; check the boolean result.

Check for emptiness before using `len(items)-1`. Clear removed pointer-bearing
values before reslicing so the backing array does not retain them. Do not
assume every `Push` is worst-case constant time because slice growth can copy
the stored items.

#### Related material

- [C.1.3.13 Finish the newest open item first (stack)](#stack)
  owns stack-method selection.
- [D.7 Type parameters](#d7-type-parameters) explains the generic type.
- [D.9 Slice storage and retained memory](#d9-slice-storage-and-retained-memory)
  explains why `Pop` clears the removed slot.

<a id="d6-overflow-safe-integer-reversal"></a>

### E.6 Overflow-safe integer reversal

#### The problem it solves

Reverse the decimal digits of a nonnegative `int64` without first producing a
value outside the type's range. Checking only after multiplication is too late
because overflowing signed integer arithmetic has already lost the intended
value.

#### The idea in simple terms

Take the last digit from the input and append it to the result, one digit at a
time. Before shifting the result left by one decimal place, ask whether the
shift and addition would still fit. This is a **precondition overflow check**:
it proves the next arithmetic operation is safe before executing it.

#### How it works

The function accepts only nonnegative input and reports failure when the
reversed value cannot fit in `int64`.

1. Reject a negative `value`.
2. Extract `digit=value%10`.
3. Before calculating `reversed*10+digit`, require
   `reversed <= (math.MaxInt64-digit)/10`.
4. Append the digit and remove it from `value` with `value/=10`.
5. Return the result when no digits remain. Zero naturally returns zero.

#### Step-by-step examples

For `value=120`, the expected result is `(21,true)`.

1. **Step: Move digit 0**
   - **Before:** `value=120`, `reversed=0`, `digit=0`.
   - **Decision:** The guard passes, so append 0.
   - **After:** `value=12`, `reversed=0`.
   - **Known correct:** The processed suffix `0` is represented as integer 0;
     its leading zero need not be retained.
2. **Step: Move digit 2**
   - **Before:** `value=12`, `reversed=0`, `digit=2`.
   - **Decision:** The guarded expression fits.
   - **After:** `value=1`, `reversed=2`.
   - **Known correct:** The processed suffix `20` is reversed to `2`.
3. **Step: Move digit 1**
   - **Before:** `value=1`, `reversed=2`, `digit=1`.
   - **Decision:** Append the final digit.
   - **After:** `value=0`, `reversed=21`; return `(21,true)`.
   - **Known correct:** Every input digit has moved to its reversed decimal
     position.

For the overflow case `value=8085774586302733229`, the digits processed
before the final leading `8` produce `reversed=922337203685477580`. The guard
compares this with `(math.MaxInt64-8)/10 = 922337203685477579`, rejects the
append, and returns `(0,false)` without overflowing.

#### Why it works

Before each iteration, `reversed` is exactly the reverse of the digits already
removed from the original value, with insignificant leading zeroes omitted.
The remainder operator selects the next digit and division removes that same
digit. Because `digit` is between 0 and 9, subtracting it from `MaxInt64` is
safe; the guard is algebraically equivalent to
`reversed*10+digit <= MaxInt64` but does not perform the risky multiplication.
Each division removes one decimal digit, so the loop finishes.

#### Go example

```go
package integerdigits

import "math"

// ReverseNonNegative reverses a nonnegative decimal integer. ok is false when
// value is negative or when its reversal would exceed int64.
func ReverseNonNegative(value int64) (reversed int64, ok bool) {
	// This implementation deliberately excludes a sign from digit processing.
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
		// Append the extracted digit and remove it from the remaining input.
		reversed = reversed*10 + digit
		value /= 10
	}
	return reversed, true
}
```

Concrete cases to test are `0 -> (0,true)`, `121 -> (121,true)`,
`120 -> (21,true)`, `-1 -> false`, and
`8085774586302733229 -> false` because its reversal is one greater than
`math.MaxInt64`.

#### Work and memory

If `d` is the number of decimal digits, the loop takes `O(d)` time, which is
`O(log value)` for positive input. Zero and rejected negative input take
`O(1)` time. The function stores a fixed number of integers and uses `O(1)`
auxiliary memory without allocations.

#### When to use it

Use this implementation when decimal digits must be processed numerically,
negative input is outside the contract, and overflow must be reported rather
than wrapped or clamped.

#### When not to use it

Use text reversal or arbitrary-precision arithmetic when every input digit
must be preserved beyond `int64`, including leading zeroes from textual input.
Derive a separate sign rule when negative values should be reversed rather
than rejected.

#### Common mistakes and edge cases

> **Warning:** Calculating `candidate := reversed*10+digit` and then comparing
> `candidate` with `MaxInt64` cannot detect the intended mathematical overflow:
> the `int64` operation has already wrapped before the comparison.

Zero skips the loop and is a successful result. Trailing input zeroes become
leading zeroes in the reversed digit sequence and disappear from its integer
representation. Keep the guard's subtraction before division; comparing only
`reversed > MaxInt64/10` misses the last-digit tie case.

#### Related material

- [A.8 Integer digit processing](#integer-digit-processing) owns digit
  extraction and the reversal invariant.
- [A.4 Integer limits and bits](#a4-integer-limits-and-bits) explains why
  range checks must precede unsafe arithmetic.
- [Palindrome Number](problems/solved_easy_leetcode_0009_palindrome_number.go)
  uses the same digit-extraction invariant under a narrower input range.

<a id="d7-common-prefix-scan"></a>

### E.7 Common-prefix scan

#### The problem it solves

Find the longest starting byte sequence shared by every string. The scan must
stop safely when a shorter string ends and must define what empty input means.

#### The idea in simple terms

Read the strings as aligned rows and compare one column at a time. The first
unequal or missing byte ends the answer because a prefix cannot skip a column.
This is a **common-prefix scan**; this implementation deliberately compares
bytes rather than Unicode code points or user-visible symbols.

#### How it works

1. Return `""` when `strs` is empty.
2. Use each byte position in `strs[0]` as a candidate `column`.
3. Compare `strs[row][column]` with the first string's `expected` byte, first
   checking that the row is long enough.
4. On a missing or unequal byte, return `strs[0][:column]`.
5. If every byte in the first string matches, return that whole string.

#### Step-by-step examples

For `strs=["flower","flow","flight"]`, the expected result is `"fl"`.

1. **Step: Compare column 0**
   - **Before:** `column=0`, `expected='f'`.
   - **Decision:** Rows 1 and 2 both contain `'f'` at byte 0.
   - **After:** Continue to column 1.
   - **Known correct:** Every string shares prefix `"f"`.
2. **Step: Compare column 1**
   - **Before:** `column=1`, `expected='l'`.
   - **Decision:** Both other rows also contain `'l'`.
   - **After:** Continue to column 2.
   - **Known correct:** Every string shares prefix `"fl"`.
3. **Step: Compare column 2**
   - **Before:** `column=2`, `expected='o'`; row 1 contains `'o'`, row 2
     contains `'i'`.
   - **Decision:** The mismatch means no common prefix can include column 2.
   - **After:** Return `strs[0][:2]`, which is `"fl"`.
   - **Known correct:** The first two bytes match everywhere and the third
     does not.

For `strs=["","abc"]`, the outer loop has no columns and returns `""`.
For `strs=[]`, the explicit guard returns `""` before reading `strs[0]`.

#### Why it works

Before each new column, every earlier byte has matched in every string, so the
slice before `column` is known to be common. A short row or mismatch proves
that no longer common prefix is possible. Otherwise the newly checked byte can
be added to the known prefix. The scan ends at the first failure or after all
bytes of the first string, which is itself the maximum possible answer length.

#### Go example

```go
package commonprefixscan

// LongestCommonPrefix returns the longest byte prefix shared by all strings.
// It returns an empty string when strs is empty.
func LongestCommonPrefix(strs []string) string {
	// Empty input has no first row to use as the comparison baseline.
	if len(strs) == 0 {
		return ""
	}

	// Each byte in the first string is the expected value for one column.
	for column := 0; column < len(strs[0]); column++ {
		expected := strs[0][column]
		for row := 1; row < len(strs); row++ {
			// Check length before indexing. A missing or unequal byte proves that the
			// previously verified columns form the longest common prefix.
			if column >= len(strs[row]) || strs[row][column] != expected {
				return strs[0][:column]
			}
		}
	}
	// Every byte of the first string matched, so no longer prefix is possible.
	return strs[0]
}
```

Concrete cases to test include the `"flower"` example; `["dog","racecar"]`
producing `""`; `["ab","a"]` producing `"a"`; identical strings; one
string; an empty first or later string; and no strings.

#### Work and memory

If the first string has `m` bytes and there are `n` strings, the worst case
does `O(n*m)` byte comparisons. A mismatch can return earlier, including in
constant time at the first column. The function uses `O(1)` auxiliary memory.
The returned substring reuses the first string's bytes rather than allocating
a copied result.

#### When to use it

Use this implementation when byte-prefix equality is the contract, all strings
must participate, and returning a substring of the first input is acceptable.

#### When not to use it

Do not use byte columns when the task defines positions as Unicode code points
or visible symbols; use the representation described in [D.6](#d6-bytes-runes-and-visible-characters).
A trie may be more useful when many changing queries share a large stored
dictionary.

#### Common mistakes and edge cases

> **Warning:** Reading `strs[row][column]` before checking
> `column < len(strs[row])` panics for `["ab","a"]` at column 1.

Guard empty input before reading `strs[0]`. The first string may be the shortest
without being specially selected. Byte slicing can split a multi-byte UTF-8
encoding if a different string disagrees inside that encoding, which is why
the byte contract must be explicit.

#### Related material

- [C.1.3.1 Compare from the left (common-prefix scan)](#common-prefix-scan)
  owns the method and its invariant.
- [D.6 Bytes, runes, and visible characters](#d6-bytes-runes-and-visible-characters)
  explains how to choose the unit of comparison.
- [Longest Common Prefix](problems/solved_easy_leetcode_0014_longest_common_prefix.go)
  is the representative problem.

<a id="d8-directional-scan"></a>

### E.8 Directional scan

#### The problem it solves

Apply a local signed-sum rule in which each value is subtracted when it is
smaller than its immediate right neighbor and added otherwise. Each decision
needs only one neighbor, but that neighbor must remain the original adjacent
value rather than the running total.

#### The idea in simple terms

Start at the right edge, whose value is always added, and carry that original
value one position left at a time. A pass whose direction makes needed context
available is a **directional scan**. Here `rightValue` holds exactly the
neighbor needed for the next comparison.

#### How it works

The running total must fit in `int`.

1. Return zero for empty input.
2. Initialize `total` and `rightValue` from the final value.
3. Move `index` from the next-to-last position toward zero.
4. Subtract `value` when `value < rightValue`; otherwise add it.
5. Set `rightValue=value` so the next iteration sees its immediate original
   neighbor.

#### Step-by-step examples

For `values=[3,4,2]`, the expected result is `3` because the signed expression
is `-3+4+2`.

1. **Step: Establish the right edge**
   - **Before:** `last=2`, no total yet.
   - **Decision:** The last value has no right neighbor, so add it.
   - **After:** `total=2`, `rightValue=2`.
   - **Known correct:** Index 2 has its final contribution.
2. **Step: Visit index 1**
   - **Before:** `value=4`, `rightValue=2`, `total=2`.
   - **Decision:** `4` is not smaller than 2, so add it.
   - **After:** `total=6`, `rightValue=4`.
   - **Known correct:** Indices 1 and 2 have their final contributions, and
     `rightValue` is the original value at index 1.
3. **Step: Visit index 0**
   - **Before:** `value=3`, `rightValue=4`, `total=6`.
   - **Decision:** `3 < 4`, so subtract it.
   - **After:** `total=3`, `rightValue=3`; return `3`.
   - **Known correct:** Every index was classified against its immediate
     original right neighbor.

For empty input, the result is zero. For `values=[7]`, initialization produces
`total=7`, the loop performs no iterations, and the result is 7.

#### Why it works

Before each loop iteration, `total` contains the final signed contributions of
all positions to the right of `index`, and `rightValue` is exactly the original
value at `index+1`. The comparison therefore applies the required local rule
to the current position. Assigning `rightValue=value` establishes the same
fact one position farther left. The index decreases to zero, so all positions
are processed once and the final sum follows the rule.

#### Go example

```go
package directionalscan

// SumWithSmallerLeftValuesSubtracted applies a local signed-sum rule: a value
// is subtracted when it is smaller than its immediate right neighbor and added
// otherwise. The running total must fit in an int.
func SumWithSmallerLeftValuesSubtracted(values []int) int {
	if len(values) == 0 {
		return 0
	}

	// The final value has no right neighbor, so it is always added. rightValue
	// carries that original neighbor value into the next iteration.
	last := len(values) - 1
	total := values[last]
	rightValue := values[last]

	for index := last - 1; index >= 0; index-- {
		value := values[index]
		// Apply the local rule against the immediate original right neighbor, not
		// against the accumulated total.
		if value < rightValue {
			total -= value
		} else {
			total += value
		}
		// Shift the neighbor state one position left for the next comparison.
		rightValue = value
	}

	return total
}
```

Concrete cases to test are `[] -> 0`, `[7] -> 7`, `[5,4,3] -> 12`,
`[3,4] -> 1`, and `[3,4,2] -> 3`, plus values near the accepted integer
limits to enforce the no-overflow contract.

#### Work and memory

The function visits each of `n` values once, so it takes `O(n)` time; empty and
single inputs are still within that bound. It stores a fixed number of
integers, uses `O(1)` auxiliary memory, and makes no allocations.

#### When to use it

Use a directional scan when each decision depends on a fixed nearby value that
can be carried from one position to the next and the output can be accumulated
without retaining all intermediate states.

#### When not to use it

Do not use this exact local rule for encodings whose subtraction decision
depends on a larger token or on a reduced value to the right. Use explicit
parsing state when several prior symbols determine one unit.

#### Common mistakes and edge cases

> **Warning:** Assigning `rightValue=total` compares the next value with an
> accumulated sum instead of its neighbor. For `[3,4,2]`, that mistake can
> classify index 0 against 6 rather than 4.

Handle empty input before calculating `len(values)-1`. State what happens at
the last item. Check equality explicitly through the contract: equal adjacent
values are added because only a strictly smaller left value is subtracted.

#### Related material

- [C.1.3.2 Process the needed neighbor first (directional scan)](#directional-scan)
  owns method selection and direction choice.
- [A.2 Correctness and invariants](#a2-correctness-and-invariants) explains
  the maintained `rightValue` fact.
- [Roman to Integer](problems/solved_easy_leetcode_0013_roman_to_integer.go)
  uses a related right-neighbor decision under Roman-numeral input rules.

<a id="d9-sorted-two-pointer-pair-search"></a>

### E.9 Sorted two-pointer pair search

#### The problem it solves

Find two distinct positions in an ascending sorted integer slice whose values
sum to a target, using no map. A full pair-by-pair search does quadratic work
even though the ordering can rule out many pairs at once.

#### The idea in simple terms

Place one pointer at each end. If their sum is too small, only a larger left
value can help; if it is too large, only a smaller right value can help. Moving
the two boundary positions inward is the **two-pointer** method for this sorted
pair search.

#### How it works

The values must be sorted in ascending order, and every evaluated sum must fit
in `int`.

1. Set `left=0` and `right=len(values)-1`.
2. While `left < right`, compare `values[left]+values[right]` with `target`.
3. Increment `left` when the sum is too small; decrement `right` when it is too
   large.
4. Return both indexes on equality. If the pointers meet without equality,
   return `false`.

#### Step-by-step examples

For `values=[-3,1,2,5,8]` and `target=7`, the expected result is
`(2,3,true)`.

1. **Step: Check indexes 0 and 4**
   - **Before:** `left=0`, `right=4`, sum `-3+8=5`.
   - **Decision:** The sum is too small, so increment `left`.
   - **After:** Candidate range is indexes 1 through 4.
   - **Known correct:** Index 0 cannot pair with any remaining value because 8
     was already its largest possible partner.
2. **Step: Check indexes 1 and 4**
   - **Before:** Sum `1+8=9`.
   - **Decision:** The sum is too large, so decrement `right`.
   - **After:** Candidate range is indexes 1 through 3.
   - **Known correct:** Index 4 cannot pair with any remaining value because 1
     was already its smallest possible partner.
3. **Step: Check indexes 1 and 3**
   - **Before:** Sum `1+5=6`.
   - **Decision:** The sum is too small; increment `left`.
   - **After:** `left=2`, `right=3`.
   - **Known correct:** Index 1 has no valid remaining partner.
4. **Step: Check indexes 2 and 3**
   - **Before:** Sum `2+5=7`.
   - **Decision:** The sum equals the target.
   - **After:** Return `(2,3,true)`.
   - **Known correct:** The positions are distinct and their values sum to 7.

For `values=[4]`, `target=8`, `left=0` and `right=0`, so the loop never runs
and returns `false`; one position cannot be used twice. Empty input also
returns `false` without indexing the slice.

#### Why it works

If the current sum is too small, pairing `values[left]` with any index below
`right` can only make the sum no larger, so discarding `left` cannot remove a
solution. The symmetric argument safely discards `right` when the sum is too
large. Thus every move preserves any possible solution inside the remaining
range. One pointer moves on every unsuccessful iteration, so they eventually
meet or a valid pair is returned.

#### Go example

```go
package twopointers

// PairSumSorted returns indices of two values whose sum equals target. Values
// must be sorted in ascending order, and relevant sums must fit in an int.
func PairSumSorted(values []int, target int) (left int, right int, ok bool) {
	// Begin with the smallest and largest available values.
	left, right = 0, len(values)-1
	for left < right {
		sum := values[left] + values[right]
		switch {
		case sum < target:
			// Even the largest partner was too small for this left value, so no pair
			// using left can succeed.
			left++
		case sum > target:
			// Even the smallest partner was too large for this right value, so no pair
			// using right can succeed.
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
`[1,2,4,6],8 -> (1,3,true)`; `[4],8 -> false`; `[1,2,3],7 -> false`; duplicate
values such as `[3,3],6`; and empty input.

#### Work and memory

Each unsuccessful iteration moves one pointer inward, so there are at most
`n-1` iterations and `O(n)` time. An early match may take `O(1)` actual time.
The function uses `O(1)` auxiliary memory and makes no allocations. Sorting an
unsorted input first would add the sorting time and any copy required to
preserve caller order.

#### When to use it

Use this implementation when the values are already sorted, one matching pair
is enough, and constant auxiliary memory or ordered elimination is preferable
to a map.

#### When not to use it

Do not apply it directly to unsorted values. Sorting may also be wrong when
the original indices must be returned and are not carried through the sort;
the map implementation in [E.3](#e3-hash-map-complement-lookup) preserves
those indexes in one pass.

#### Common mistakes and edge cases

> **Warning:** Using `left <= right` allows one position to pair with itself.
> For `[4]`, target 8 would incorrectly succeed; require `left < right`.

The ordering proof fails if the slice is descending or unsorted. Include
duplicates, negative values, empty input, and missing pairs in tests. Widen or
validate addition when two accepted `int` values might overflow.

#### Related material

- [C.1.3.9 Move coordinated positions (two pointers)](#two-pointers)
  owns the elimination proof and related forms.
- [E.3 Hash-map complement lookup](#e3-hash-map-complement-lookup) is the
  unsorted one-pass alternative.
- [D.3 Sorting without unexpected mutation](#d3-sorting-without-unexpected-mutation)
  explains the cost of preparing an unsorted input without reordering it.

<a id="f-further-study-checklist"></a>

## F. Further study

The material above covers the reusable concepts emphasized by the repository's
problem set and contemporary general interview-preparation lists. The topics
below are later extensions, not prerequisites for the core methods in this
guide. Study one when a target role, contest, or interview loop calls for it:

- **Strongly connected components** divide a directed graph into largest
  groups in which every node can reach every other node. Start with
  [depth-first search](#c1320-depth-first-search) and graph direction in
  [graphs and grids](#b5-graphs-and-grids).
- **All-pairs shortest paths** find the cheapest route for every start-and-end
  pair, instead of routes from one start. Compare repeated
  [Dijkstra searches](#c1323-dijkstra-shortest-paths) with reusing smaller
  answers through [dynamic programming](#c1310-dynamic-programming).
- **Maximum flow** asks how much material can travel from a source to a
  destination when every directed connection has a capacity. Its starting
  point is the adjacency representation in
  [graphs and grids](#b5-graphs-and-grids).
- **Suffix arrays and suffix automata** organize all endings of a string, or
  remembered text states, so repeated substring questions can be answered
  quickly. Begin with [string matching](#c137-string-matching) and
  [parsing state](#c138-parsing-and-state-machines).
- **Computational geometry beyond normalized slopes** uses exact orientation,
  intersection, and boundary tests to reason about points and shapes. The
  prerequisite is [exact slopes without rounding](#a63-exact-slopes-without-rounding).
- **Randomized balanced search trees** keep ordered values while random
  priorities make long one-sided shapes unlikely, giving expected logarithmic
  search and update work. Compare [binary-search-tree ordering](#b41-terminology-and-search)
  with [heap priorities](#c1327-heap-or-priority-queue).
