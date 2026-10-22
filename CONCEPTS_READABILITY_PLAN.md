# `CONCEPTS.md` Readability Plan

## Goal

Review and rewrite `CONCEPTS.md` so that a reader who knows basic Go
but is new to algorithms can understand every teaching entry without
already knowing its formal vocabulary.

Correctness and preservation remain hard requirements. Within those
limits, use this priority order:

1. Beginner comprehension.
2. Completeness.
3. Brevity and compactness.

Do not shorten an explanation when doing so removes a bridge that a
beginner needs.

## Baseline

As of 2026-10-22, `CONCEPTS.md` contains:

- 13,629 lines;
- about 74,600 words; and
- 93 teaching entries.

Reviewing that much material in one pass makes shallow compliance too
likely. Complete the work in small, independently reviewed batches.

## Progress

The five-entry calibration batch is implemented and validated:

- A.1 Complexity and resource accounting;
- B.1 Arrays, slices, and strings;
- C.1.3.31 Process change points (sweep line);
- D.1 Keep input conversion outside the algorithm; and
- E.1 Binary search bounds.

The user approved the pilot style. The remaining entries are being
audited and rewritten in the batches defined below.

The first foundation batch is also implemented and validated:

- A.2 Correctness and invariants;
- A.3 Recursion and saved work;
- A.4 Integer limits and bits;
- A.5 Bit manipulation;
- A.6 Keeping arithmetic values exact and consistent;
- A.6.1 Shared divisors and multiples; and
- A.6.2 Normalized fractions and ratios.

The second foundation batch is implemented and its revised examples
have passed focused tests:

- A.6.3 Exact slopes without rounding;
- A.6.4 Modular arithmetic and negative division;
- A.6.5 Trailing zeroes without calculating a factorial;
- A.6.6 Factors and prime sieves;
- A.6.7 Numbers too large for an integer type; and
- A.7 Combinatorics and output size.

The final foundation batch is implemented and its revised examples
have passed focused tests:

- A.7.1 Counting combinations without large factorials;
- A.7.2 Repeated input values; and
- A.8 Integer Digit Processing.

The first storage batch is implemented and its revised examples have
passed focused tests:

- B.1.1 Moving values in place;
- B.2 Linked Lists;
- B.3 Queues and deques;
- B.4.1 Terminology and search;
- B.4.2 Traversal order; and
- B.4.3 Recursive return contracts.

The second storage batch is implemented and its revised examples have
passed focused tests:

- B.4.4 BST validation and costs;
- B.4.5 Reconstruction from traversals;
- B.4.6 Traversing through parent links;
- B.5 Graphs and grids;
- B.6 Matrix traversal and in-place state;
- B.6.1 Rows, columns, and shrinking boundaries; and
- B.6.2 Layer transformations.

The final storage batch is implemented and its revised examples have
passed focused tests:

- B.6.3 Reusing the first row and column as markers;
- B.6.4 Keeping current and next state together;
- B.7 Composing data structures;
- B.8 Streaming and online processing; and
- B.9 Choosing a structure.

The method-selection guide batch is implemented and validated:

- the C.1 overview and pattern-card navigation;
- C.1.1 Ask what the input and output allow;
- C.1.2 Derive a direction before naming a pattern;
- C.1.4 Choosing among overlapping methods;
- C.1.5 Let constraints eliminate bad choices; and
- C.1.6 Before committing to a method.

The first pattern-card batch is implemented, and its code paths have
passed focused tests:

- C.1.3.1 Common-prefix scan;
- C.1.3.2 Directional scan;
- C.1.3.3 Simulation;
- C.1.3.4 Hash map or set;
- C.1.3.5 Sorting and selection; and
- C.1.3.6 Binary search.

The second pattern-card batch is implemented, and its code paths have
passed focused tests:

- C.1.3.7 String matching;
- C.1.3.8 Parsing and state machines;
- C.1.3.9 Two pointers;
- C.1.3.10 Dynamic programming;
- C.1.3.10.1 Coordinate compression;
- C.1.3.11 Divide and conquer; and
- C.1.3.12 Kadane's algorithm.

The third pattern-card batch is implemented, and its code paths have
passed focused tests:

- C.1.3.13 Stack;
- C.1.3.14 Sliding window;
- C.1.3.15 Prefix sums;
- C.1.3.16 Difference arrays;
- C.1.3.17 Fenwick and segment trees;
- C.1.3.18 Greedy; and
- C.1.3.19 Backtracking.

The graph-method pattern-card batch is implemented, and its code paths
have passed focused tests:

- C.1.3.20 Depth-first search;
- C.1.3.21 Breadth-first search;
- C.1.3.22 Topological sort;
- C.1.3.23 Dijkstra;
- C.1.3.24 Bellman-Ford;
- C.1.3.25 Disjoint-set union; and
- C.1.3.26 Minimum spanning tree.

The final pattern-card batch is implemented, and its code paths have
passed focused tests:

- C.1.3.27 Heap or priority queue;
- C.1.3.28 Monotonic stack;
- C.1.3.29 Monotonic deque;
- C.1.3.30 Intervals;
- C.1.3.31 Sweep line, rechecked after its approved pilot rewrite;
- C.1.3.32 Fast and slow pointers; and
- C.1.3.33 Trie.

The Go-technique batch is implemented, and all example code paths have
passed focused tests:

- D.1 Input conversion, rechecked after its approved pilot rewrite;
- D.2 Table-driven tests;
- D.3 Sorting without unexpected mutation;
- D.4 Map presence and sets;
- D.5 Priority queues;
- D.6 Bytes, runes, and visible characters;
- D.7 Type parameters;
- D.8 Replacing recursion with a loop; and
- D.9 Slice storage and retained memory.

The reference-implementation batch is implemented, and all code paths
have passed focused tests:

- E.1 Binary search bounds, rechecked after its approved pilot
  rewrite;
- E.2 Rolling dynamic programming;
- E.3 Hash-map complement lookup;
- E.4 Generic singly linked list;
- E.5 Generic stack;
- E.6 Overflow-safe integer reversal;
- E.7 Common-prefix scan;
- E.8 Directional scan; and
- E.9 Sorted two-pointer pair search.

The final guide-level prose batch is implemented and structurally
validated:

- the opening description and reading paths;
- the table-of-contents navigation;
- the explanation of the guide's teaching format;
- the interview-evidence summary; and
- F. Further study.

The final whole-file consistency and readability pass is complete. It
included cross-section cold reads, terminology searches, transition
checks, and an audit of all numbered traces. All 398 traced steps now
contain separate `Before`, `Decision`, `After`, and `Known correct`
details. The complete validation matrix passes.

## Target reader

Write every entry for someone who:

- understands basic Go syntax;
- has not studied algorithms formally;
- does not already know the concept's vocabulary; and
- may open the entry without reading the preceding entries.

Links to prerequisites are useful, but they must not replace the
minimum explanation needed to understand the current entry.

## Phase 1: Calibrate the writing standard

Rewrite five representative entries first:

1. A foundational idea, such as complexity.
2. A data structure, such as arrays or trees.
3. `C.1.3.31 Sweep line`.
4. A Go technique.
5. A reference implementation.

For each pilot entry:

- begin with a familiar, concrete problem;
- explain the idea without its technical name;
- introduce the formal name only afterward;
- explain notation in words before displaying it;
- trace every meaningful state change separately;
- make the Go code accept the same input used by the prose;
- explain correctness without relying on words such as "obvious";
  and
- explain where time and memory are spent before using Big-O
  notation.

Pause after these five entries for user review. Apply the accepted
style to the rest of the guide only after that review.

## Phase 2: Audit each entry in isolation

Before rewriting an entry, treat it as if it were the only available
page. Record every instance of:

- undefined terminology;
- unexplained notation;
- abstract nouns without concrete examples;
- compressed or skipped state changes;
- vague references such as "it," "state," or "position";
- examples without exact inputs or outputs;
- prose, traces, and code that use different names;
- correctness claims without explanations;
- Big-O symbols whose variables are undefined; and
- rejection guidance that does not name a useful alternative.

Rewrite the entry from this audit. Do not merely polish individual
sentences while preserving a confusing teaching order.

## Phase 3: Rewrite in small batches

Use batches of about six to eight entries, adjusted for entry length:

1. Foundations: 17 entries across about three batches.
2. Data structures: 19 entries across about three batches.
3. Method-selection guide: 6 entries in one batch.
4. Pattern cards: 33 entries across five or six batches.
5. Go techniques: 9 entries across two batches.
6. Reference implementations: 9 entries across two batches.
7. Introduction, navigation, interview evidence, and further study in
   one final prose batch.

Complete and validate each batch before beginning the next one.
Preserve any existing uncommitted work in `CONCEPTS.md`.

## Per-entry readability gate

An entry is complete only when a beginner can answer all of these
questions from that entry:

1. What real problem are we trying to solve?
2. What does the input look like?
3. What exact result do we want?
4. What is the central idea in one ordinary sentence?
5. What does every introduced term or symbol mean?
6. Can I replay the example without filling in missing steps?
7. Can I match every traced value to the Go code?
8. Why does each action preserve the correct answer?
9. Why does the process finish?
10. Where do the time and memory costs come from?
11. What clue tells me to use or reject this method?

A heading being present does not satisfy its requirement.

## Example standard

Every main example should use this teaching order:

1. Describe the situation in ordinary words.
2. State the exact input.
3. State the expected output.
4. Show the original representation.
5. Explain any new representation.
6. Trace one state change at a time.
7. Show the final answer.
8. Connect the trace to the code.

Boundary examples must be equally explicit. Notation such as
`[1,4)`, `+1@1`, `i < n`, or `O(n)` must not appear before its meaning
is explained.

## Language standard

Keep necessary technical names because readers need to recognize them,
but introduce each name after its plain-English explanation.

Avoid or immediately explain words such as:

- invariant;
- state;
- event;
- boundary;
- online;
- auxiliary;
- amortized;
- stable;
- monotonic;
- relaxation;
- predecessor;
- partition; and
- optimal substructure.

The goal is not to delete terminology. The goal is to prevent
terminology from carrying the explanation.

## Validation after each batch

Perform two separate reviews.

### Readability review

Check:

- undefined words and notation;
- replayability of examples;
- agreement among prose, traces, and code;
- sentence density;
- hidden prerequisites; and
- whether each explanation works in isolation.

### Technical review

Check:

- algorithm correctness;
- boundary behavior;
- complexity;
- Go semantics;
- links and anchors; and
- compiled and exercised Go examples.

Then run the required commands in `CONCEPTS_STANDARDS.md`:

```bash
~/.codex/scripts/check-non-ascii.py \
  CONCEPTS_STANDARDS.md CONCEPTS.md
git diff --check HEAD -- CONCEPTS_STANDARDS.md CONCEPTS.md
go run internal/check-concepts-doc.go
go run ./internal/linewidth
go vet internal/check-concepts-doc.go
go test internal/check-concepts-doc.go
go test ./...
```

Also format, compile, and exercise every changed complete Go example
in a temporary package. Remove temporary validation artifacts after
the checks.

Automated checks provide structural evidence only. Never treat them as
proof that a beginner can understand the writing.

## Final whole-file pass

After all entries pass individually:

- verify that concepts appear in a sensible learning order;
- remove needless duplication without removing teaching bridges;
- check that terminology is introduced consistently;
- ensure cross-links offer optional depth instead of replacing missing
  explanations;
- read transitions between adjacent entries;
- cold-read sample entries from every section; and
- run the complete validation matrix.

## Definition of done

The readability review is complete only when:

- every teaching entry passes the per-entry readability gate;
- the pilot style has been approved by the user;
- every changed trace agrees with its code;
- every changed Go example compiles and passes its exercised cases;
- all repository validation commands pass;
- links, anchors, and useful existing material remain intact; and
- the final cross-section cold read finds no unexplained prerequisite
  or notation.

The structural contract remains owned by `CONCEPTS_STANDARDS.md`.
This document owns the execution plan for the readability review.
