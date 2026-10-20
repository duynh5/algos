# Plan to Refactor `concepts/README.md`

## Goal

Reorganize `concepts/README.md` into a consistent learning guide in which every
algorithm, data structure, problem-solving method, and Go technique:

1. is explained first in ordinary language;
2. introduces a technical name only after explaining the underlying idea;
3. defines every necessary technical term where it first appears;
4. includes one or more concrete examples;
5. traces those examples step by step; and
6. explains why each step is valid, what can go wrong, and how the work and
   memory grow.

The refactor must preserve the guide's technical accuracy, existing subject
coverage, useful Go code, source links, and links from elsewhere in the
repository.

## Interpretation of "no jargon"

Algorithm names such as *binary search*, *dynamic programming*, and *Big-O*
cannot disappear because readers need those names to recognize problems,
documentation, and interview questions. The standard will therefore be **no
unexplained jargon**:

- Lead with the idea in plain English, then give its conventional name.
- Define a new term in the same paragraph in which it first appears.
- Do not define one unfamiliar term with several other unfamiliar terms.
- Prefer concrete verbs such as "visit," "remember," "compare," and "move"
  over abstract phrases such as "perform traversal" or "leverage state."
- Explain notation in words before using it repeatedly.
- Treat the reader as someone who knows basic Go syntax but is new to
  algorithms and data structures.
- Keep a technical term when replacing it would make the explanation less
  precise. Pair it with a plain-English restatement instead of hiding it.

## Current-state findings

The current file is a 3,901-line guide with several different presentation
styles:

- Foundation and data-structure entries alternate between free-form prose,
  bold labels, and numbered subsections.
- The 33 problem-solving pattern cards are the most consistent entries, but
  their field names, ordering, and completeness still differ.
- Go reference implementations follow a code-plus-test-cases format rather
  than the teaching sequence used elsewhere.
- Go techniques use free-form explanations and do not consistently include a
  detailed trace or analysis.
- Some subjects have a simple analogy but no complete worked example. Others
  have an example but do not explain every state change or decision.
- Several explanations repeat material owned by another section, which makes
  the file harder to keep consistent.
- Heading capitalization and labels such as `Example`, `Worked example`,
  `Walk-through`, `Work and memory`, and `Complexity` are inconsistent.

These differences make it difficult for a reader to predict where to find an
example, correctness explanation, cost analysis, or warning.

## Scope

### Included

- The introductory reading guidance and table of contents.
- Foundations A.1 through A.8 and their nested topics.
- Data structures B.1 through B.9 and their nested topics.
- The method-selection guide and all 33 pattern cards in section C.
- All nine Go reference implementations in section D.
- All nine Go techniques in section E.
- Cross-references, heading names, anchors, terminology, and duplicated
  explanations affected by the reorganization.
- The further-study checklist where its wording depends on terms introduced by
  the guide.

### Not included

- Adding new algorithms solely to increase subject coverage.
- Changing completed problem solutions under `problems/`.
- Changing the workbench or test runner.
- Replacing correct Go implementations merely to impose a different coding
  style.
- Splitting the guide into multiple files. The first refactor should keep
  `concepts/README.md` as the single owning document; splitting can be assessed
  separately after the content is consistent.
- Rewriting or removing the interview-evidence section unless a change is
  necessary for navigation or terminology.

## Required format for every concept or technique

Every teaching entry will use the following headings in this order. A short
entry may keep a section concise, but it may not omit a required heading.

### 1. The problem it solves

State the reader's problem without using the technique's name as the
explanation. Describe the shape of the input and the desired outcome.

### 2. The idea in simple terms

Give a short mental picture or everyday analogy, then introduce and define the
formal name. This replaces inconsistent labels such as `ELI5`, `Mental model`,
and `Plain-language picture`.

### 3. How it works

Describe the procedure as a small sequence of explicit actions. Name the
information the method keeps and what each piece means. State required input
conditions before the procedure uses them.

### 4. Step-by-step examples

Include at least:

- one small, ordinary example that demonstrates the complete method; and
- one boundary, duplicate, failure, or adversarial example whenever such an
  input changes behavior or exposes a common mistake.

Each example must identify the exact input and expected output. A detailed
trace must show, for every meaningful step:

- the current value, position, node, or choice;
- the state before the step;
- the decision made and why;
- the state after the step; and
- what part of the result is already known to be correct.

Use a table when several steps update the same fields. Use a numbered list when
the state is small or the process branches. Do not call an outcome-only
paragraph a worked example.

### 5. Why it works

Explain correctness in plain English. Identify the fact that remains true
while the method runs, explain why it is true initially, why one step preserves
it, why the method finishes, and why the final state answers the question.
Introduce the word *invariant* only after describing it as "a fact that remains
true."

For a data structure or language technique, explain why its operations or Go
behavior produce the demonstrated result. Do not force an algorithm-style
proof where a precise behavior explanation is more appropriate.

### 6. Go example

Provide the smallest complete Go fragment that maps directly to the preceding
trace. Explain any Go-specific behavior needed to read it. Code names should
match names used in the example and analysis.

If a full implementation already exists in the reference section, establish
one owner for it and link to that owner rather than maintaining two copies.

### 7. Work and memory

Explain costs in words first, then give conventional notation. Account for
preprocessing, recursion, output construction, and meaningful Go allocations.
State best, usual, and worst cases only when they differ.

### 8. When to use it

List observable clues in a problem statement or input that make the method a
good fit. Avoid advice that merely repeats the method's name.

### 9. When not to use it

State the condition that breaks the method's reasoning or makes its cost
unacceptable. Point to the closest alternative where useful.

### 10. Common mistakes and edge cases

Connect every warning to a concrete failing input, incorrect state change, or
Go behavior. Include empty or minimum input, duplicates, limits, and malformed
or unsupported input when relevant.

### 11. Related material

Link to prerequisite concepts, contrasting techniques, the owning reference
implementation, and representative repository problems. Do not repeat their
full explanations here.

## Proposed organization

The file will retain its broad learning progression but make section ownership
and entry type explicit.

### 1. How to use this guide

- State the assumed reader knowledge.
- Explain the common entry format once.
- Explain that plain-English descriptions come before formal names.
- Provide paths for first-time study, problem selection, and quick reference.
- Keep contemporary interview evidence separate from the instructional
  template because it provides context rather than teaching a concept.

### 2. Foundations

Keep complexity, correctness, saved work, integer limits, bit operations,
exact arithmetic, combinations, and digit processing together. Convert every
top-level and nested foundation topic into the required format. Avoid treating
a long parent section as a substitute for examples in its distinct child
topics.

### 3. Ways to store and organize information

Keep arrays and strings, linked lists, queues, trees, graphs, matrices,
composed structures, streaming state, and structure selection together. Give
each data structure a concrete operation trace and make costs specific to the
operation being demonstrated.

### 4. Ways to solve problems

Keep the selection guide followed by the pattern cards. Rename pattern-card
headings so the plain-English action comes first where the formal name is not
self-explanatory, for example:

- "Remember answers to smaller problems (dynamic programming)"
- "Try a choice and undo it (backtracking)"
- "Visit nearby items before distant items (breadth-first search)"
- "Keep the next most urgent item easy to remove (priority queue)"

The conventional name must remain visible in each heading and in the table of
contents so readers can still search for it.

### 5. Go-specific techniques

Move the current section E before the reference-code index. Treat each Go
technique as a full teaching entry, including a traced example and an
explanation of the relevant language behavior. Cross-link general concepts
instead of re-explaining them.

### 6. Go reference implementations

Turn section D into an implementation index and detailed code companion:

- Each implementation names the concept or technique it demonstrates.
- The teaching explanation has one clear owner.
- Each snippet includes its contract, required input conditions, a trace that
  uses the actual variable names, cost analysis, and concrete cases to test.
- Duplicate prose is replaced with links to the owning concept entry.

### 7. Further study

Keep this as a clearly labeled next-step list. Give each unfamiliar name a
one-sentence plain-English description or link instead of presenting a wall of
undefined terminology.

## Content ownership and duplication rules

- Every idea has one primary teaching entry.
- The method-selection guide helps choose an entry; it does not repeat the
  entry's full explanation.
- A reference implementation demonstrates an idea; it links back to the
  primary explanation.
- A Go technique explains Go behavior; an algorithm entry links to it when
  that behavior matters.
- A data-structure entry owns the structure's basic operations. Pattern cards
  explain how a problem-solving method uses those operations.
- Related sections summarize a relationship in one or two sentences and link
  to the owner.

## Phase-by-phase execution contract

Each row in the phase ledger below is one independently implementable change
set. A request such as "implement phase 3" means only that row, plus the
shared checks and tracker updates required by this contract. Do not begin the
next row merely because time remains.

At the start of every phase:

1. Read this plan, the phase's exact slice of `concepts/README.md`, and the
   current Git status and diff. The working-tree version is the source of truth;
   never restore the file from `HEAD` because it may contain user changes.
2. Confirm that every dependency in the ledger is `Complete`. If a dependency
   is incomplete, stop and report it instead of implementing phases out of
   order.
3. Change the phase status to `In progress` and record the starting heading
   names in the implementation notes created by phase 0.

While implementing a phase:

- Edit only the named scope. A necessary cross-section correction may be made
  only when the current entry would otherwise be wrong; record it explicitly in
  the implementation notes. Put unrelated discoveries in the deferred-work
  list instead of expanding the phase.
- Complete each entry in the phase. Do not leave an entry half-converted to the
  eleven-part format.
- Preserve an old heading's fragment identifier with an explicit HTML anchor
  whenever the heading changes. Do not postpone that protection to the final
  navigation phase.
- Update the coverage matrix and content-ownership decisions for the entries
  touched by the phase.
- Do not stage or commit changes unless the user separately asks for that Git
  action.

At the end of every phase:

1. Apply the per-entry review checklist to every changed entry.
2. Run the common phase checks below and any extra checks named by the phase.
3. Inspect the complete phase diff for lost prose, examples, links, code, or
   sources.
4. Change the ledger status to `Complete` only when the exit condition passes.
   Otherwise leave it `In progress` and record the raw failure and the next
   executable step.
5. Record changed headings, preserved anchors, ownership decisions, validation
   commands, and results in the implementation notes. This is the handoff for
   the next phase.

### Common phase checks

Run these checks after every phase:

```bash
/Users/nolan/.codex/scripts/check-non-ascii.py REFACTOR_CONCEPTS.md concepts/README.md
git diff --check HEAD -- REFACTOR_CONCEPTS.md concepts/README.md
rg -n '^#{1,6} ' concepts/README.md
```

Use the heading output to inspect hierarchy and numbering for the changed
slice. Also verify that fenced code blocks remain balanced and that every local
path or fragment added or changed by the phase resolves. When a phase changes a
complete Go example, format and compile that example in a temporary file. The
temporary file is validation material, not a new repository artifact.

### Phase ledger

Phase 0 adds and thereafter maintains two sections at the end of this file:
`Implementation notes` and `Coverage matrix`. The notes include a deferred-work
list. Valid statuses are `Not started`, `In progress`, `Blocked`, and
`Complete`.

| Phase | Depends on | Exact scope | Status |
| --- | --- | --- | --- |
| 0 | None | Baseline inventory and tracking artifacts | Complete |
| 1 | 0 | Entry contract, guide introduction, and three pilot entries | Not started |
| 2 | 1 | Foundations A.1-A.8 | Not started |
| 3 | 2 | Data structures B.1-B.9 | Not started |
| 4 | 3 | Method selection and all 33 pattern cards in C | Not started |
| 5 | 4 | All nine Go techniques in E | Not started |
| 6 | 5 | All nine Go reference implementations in D | Not started |
| 7 | 6 | Final organization, further study, navigation, and link preservation | Not started |
| 8 | 7 | Beginner-language and technical review of the complete guide | Not started |
| 9 | 8 | Repository-wide documentation and test validation | Not started |

There are exactly ten phases, numbered 0 through 9. The ordered batches inside
a phase are checkpoints, not additional phases: they do not receive separate
ledger rows, and the phase remains `In progress` until every batch passes its
checks.

## Phase definitions

### Phase 0: Preserve the baseline and build the trackers

Inventory every existing heading, explicit anchor, inbound repository link,
outbound link, fenced code block, example, source, and representative problem
link. Add the coverage matrix and implementation-notes sections to this file.
The matrix must have one row per teaching entry and one column for each of the
eleven required fields, plus columns for owner, original heading or anchor,
phase, editorial review, technical review, and notes.

Assign a destination and primary owner to every topic. Record uncertain claims
without rewriting them. Capture enough baseline counts to detect accidental
loss later, but do not edit `concepts/README.md` in this phase.

**Exit condition:** Every current teaching topic has one destination and phase,
all inbound links and explicit anchors are recorded, and the baseline counts
can be reproduced.

### Phase 1: Establish the contract with real pilots

Rewrite the guide introduction so it states the reader assumptions, the common
entry format, the no-unexplained-jargon rule, and the available reading paths.
Define the exact Markdown forms for traces, warnings, code, and related links.
Then fully refactor these pilot entries:

- A.2, correctness and invariants;
- C.1.3.1, common-prefix scan; and
- E.4, map presence and sets.

The pilots are completed content, not disposable examples. Later phases review
them but do not rewrite them without a concrete defect.

**Exit condition:** All three entry types satisfy the checklist without empty
or forced sections, and the agreed structure is recorded in the implementation
notes as the standard for later phases.

### Phase 2: Refactor foundations

Complete these batches in order, running the common checks and updating the
coverage matrix after each one:

1. A.1 and A.3-A.5. Reconcile complexity terms, recursion memory, integer
   limits, and bit terminology with completed pilot A.2.
2. A.6-A.6.4. Establish exact-arithmetic ownership before removing duplicated
   explanations.
3. A.6.5-A.8. Give each factoring, large-number, combinatorics,
   repeated-value, and digit-processing topic its own complete example rather
   than relying on a parent section.

**Exit condition:** Every foundation matrix row is complete and has a
replayable trace, verified cost analysis, and preserved source and
related-problem links.

### Phase 3: Refactor data structures

Complete these batches in order:

1. B.1-B.3: arrays, slices, strings, linked lists, queues, and deques.
2. B.4-B.6: trees, binary-search trees, graphs, grids, and matrix state.
3. B.7-B.9: composed structures, streaming state, and structure selection.

For every structure, identify what it stores, trace at least one concrete
mutation or query, state what remains true, and give operation-specific costs.
B.9 compares structures and links to their owners; it must not repeat their
full definitions.

**Exit condition:** A reader can reconstruct the before and after state of
every demonstrated operation from the prose alone, and all data-structure
matrix rows are complete.

### Phase 4: Refactor method selection and pattern cards

First rewrite the C introduction, C.1, C.1.1, and C.1.2 as questions about
input shape, required output, allowed mutation, constraints, and correctness
requirements. Then refactor the pattern cards in these ordered batches:

1. C.1.3.2-C.1.3.5; C.1.3.1 was completed as a pilot.
2. C.1.3.6-C.1.3.9.
3. C.1.3.10-C.1.3.13, including C.1.3.10.1.
4. C.1.3.14-C.1.3.17.
5. C.1.3.18-C.1.3.21.
6. C.1.3.22-C.1.3.25.
7. C.1.3.26-C.1.3.29.
8. C.1.3.30-C.1.3.33.

Run the common checks and update the coverage matrix after every batch. For
each card:

- retain useful material while replacing inconsistent field labels;
- add an ordinary trace and an adversarial or boundary trace when relevant;
- introduce specialized variants only after the base method;
- derive rejection conditions from the correctness argument; and
- keep the formal name searchable in the heading.

Finally reconcile C.1.4-C.1.6 with the completed cards. This material owns
comparisons such as depth-first versus breadth-first search and sliding windows
versus prefix sums.

**Exit condition:** All 33 cards and the selection guide are independently
understandable, pass the checklist, and agree with the comparison guidance.

### Phase 5: Refactor Go techniques

Complete E.1-E.3, review pilot E.4 for cross-link consistency without otherwise
reopening it, then complete E.5-E.9. Run the common checks after each of those
three checkpoints.

Give each technique a complete example, a trace, and an explanation of the Go
behavior that produces the result. Move a Go-specific warning out of an
algorithm entry only when the warning is not part of the algorithm's reasoning;
leave a cross-link at the old location.

**Extra check:** Format and compile every complete Go example changed in the
phase.

**Exit condition:** All nine Go techniques have one clear owner, and their
prose, traces, code, and stated allocations agree.

### Phase 6: Refactor Go reference implementations

Complete D.1-D.5, run the common and extra checks, then complete D.6-D.9 and
repeat the checks. Match each implementation to its owning concept or Go
technique. Keep the contract, required input conditions, reference code,
variable-level trace, cost analysis, and concrete test cases beside the
implementation. Replace duplicated teaching prose with an owner link only
after confirming the owner is complete.

**Extra check:** Format and compile each complete implementation and its
examples in a temporary package.

**Exit condition:** All nine reference implementations compile, and their
code, traces, tests, contracts, and costs agree.

### Phase 7: Finalize organization and navigation

Move Go techniques before Go reference implementations, then update section
letters, headings, and the table of contents once. Refactor the further-study
checklist so every unfamiliar term has a plain-language description or owner
link. Search the entire repository for links into `concepts/README.md`, preserve
used legacy fragments with explicit anchors, verify all local paths and
fragments, and add reciprocal related-material links only where they improve
navigation.

**Exit condition:** The proposed organization is present, existing repository
links still resolve, every table-of-contents link resolves, and no teaching
entry is orphaned.

### Phase 8: Review the completed guide

Review the guide in three ordered batches: foundations and data structures;
method selection and pattern cards; then Go techniques, reference
implementations, and further study. Run the common checks and update the matrix
after each batch.

For every batch, perform two distinct passes:

1. **Beginner-language pass:** Find undefined terms, unexplained symbols,
   noun-heavy sentences, vague pronouns, and reasoning that assumes its own
   conclusion.
2. **Technical pass:** Replay each trace, compare it with the Go code, verify
   the correctness argument, and recalculate work and memory costs including
   preprocessing, recursion, output, and meaningful allocations.

Mark the two review columns separately; one pass cannot stand in for the other.

**Exit condition:** Every matrix row is approved in both review columns, every
example is reproducible from the text, and all findings are fixed or recorded
as a blocker with exact evidence.

### Phase 9: Run final validation

Run the common checks against the whole document, plus:

```bash
go test ./...
```

Verify final heading hierarchy and numbering, table-of-contents fragments,
local file links, balanced fences, compiled Go examples, required-entry heading
counts, and baseline inventory counts. Search for unexplained abbreviations and
vague placeholder phrases. Inspect the complete diff from the recorded
baseline and account for every moved or removed concept, example, source, and
link.

If repeated manual checks proved error-prone during earlier phases, add a
focused repository script for document structure and local links in this phase,
then run and document it. Do not add tooling merely because it might be useful.

**Exit condition:** All applicable checks pass, every matrix row and phase is
complete, no deferred item affects correctness or navigation, and the final
diff accounts for all baseline content.

## Review checklist for each entry

- [ ] The entry starts with the reader's problem, not a formal definition.
- [ ] The simple explanation makes sense without knowing the formal name.
- [ ] Every necessary technical term is defined on first use.
- [ ] The procedure states required input conditions.
- [ ] At least one complete ordinary example is present.
- [ ] A boundary or adversarial example is present when relevant.
- [ ] The trace shows state before and after every meaningful step.
- [ ] The correctness analysis explains why steps are safe and why the method
      finishes.
- [ ] The Go example uses the same names and state as the trace.
- [ ] Work and memory are explained in words before notation.
- [ ] Setup, recursion, output, and meaningful allocations are counted.
- [ ] Usage clues are concrete and observable.
- [ ] Rejection conditions identify what assumption would fail.
- [ ] Common mistakes include a concrete failing input or state.
- [ ] Related sections are linked instead of duplicated.
- [ ] Existing anchors and inbound links remain valid.

## Definition of done

The refactor is complete when:

1. every concept and technique in the current guide appears in the final
   coverage matrix;
2. every teaching entry follows the eleven-part format;
3. no technical term is used before it is explained in ordinary language;
4. every entry has a complete, replayable example and detailed analysis;
5. entries with meaningful boundary behavior include a second example that
   exposes it;
6. every Go snippet agrees with its example, analysis, and stated costs;
7. duplicated explanations have one clear owner and cross-links;
8. old repository links and intentional legacy anchors still work;
9. no existing concept, useful example, attribution, or source is lost without
   an explicit documented reason; and
10. all applicable validation checks pass.

## Recommended execution strategy

Make the refactor as a series of reviewable batches rather than one full-file
rewrite. Complete and validate one phase before starting the next. Keep the
coverage matrix current throughout the work so progress is measured by
teaching completeness, not by the number of rewritten lines.

## Implementation notes

### Phase 0: Baseline inventory and tracking artifacts

**Status:** Complete

**Baseline source:** The working-tree `concepts/README.md`, not `HEAD`. The
file already had staged edits when phase 0 began, so those edits are part of
the preserved baseline and were not changed by this phase.

**Starting headings:** There are 119 headings: one level-1 heading, eight
level-2 headings, 36 level-3 headings, 40 level-4 headings, 33 level-5
headings, and one level-6 heading. The ordered names are:

- Document and context: `Algorithms and Data Structures in Go`, `Contents`,
  and `Contemporary interview evidence`.
- Foundations: `A. Foundations`; `A.1 Complexity and resource accounting`;
  `A.2 Correctness and invariants`; `A.3 Recursion and saved work`;
  `A.4 Integer limits and bits`; `A.5 Bit manipulation`; `A.6 Keeping
  arithmetic values exact and consistent`; `A.6.1 Shared divisors and
  multiples (GCD and LCM)`; `A.6.2 Normalized fractions and ratios`; `A.6.3
  Exact slopes without rounding`; `A.6.4 Modular arithmetic and negative
  division`; `A.6.5 Trailing zeroes without calculating a factorial`; `A.6.6
  Factors and prime sieves`; `A.6.7 Numbers too large for an integer type`;
  `A.7 Combinatorics and output size`; `A.7.1 Counting combinations without
  large factorials`; `A.7.2 Repeated input values`; `A.8 Integer Digit
  Processing`; and its six field headings `A.8.1 Mental model`, `A.8.2 Worked
  example: reversing digits`, `A.8.3 Invariant (fact that remains true)`,
  `A.8.4 Preventing integer overflow`, `A.8.5 Work and memory`, and `A.8.6
  Related problems`.
- Data structures: `B. Data structures`; `B.1 Arrays, slices, and strings`;
  `B.1.1 Moving values in place`; `B.2 Linked Lists`; `B.2.1 Mental model`;
  `B.2.2 Core operations`; `B.2.3 Worked example: reversing a list`; `B.2.4
  Invariants (facts that remain true)`; `B.2.5 Common techniques`; `B.2.6 Work
  and memory`; `B.2.7 Common mistakes`; `B.2.8 Related problems`; `B.3 Queues
  and deques`; `B.4 Trees and binary-search trees`; `B.4.1 Terminology and search`;
  `B.4.2 Traversal order`; `B.4.3 Recursive return contracts`; `B.4.4 BST
  validation and costs`; `B.4.5 Reconstruction from traversals`; `B.4.6
  Traversing through parent links`; `B.5 Graphs and grids`; `B.6 Matrix
  traversal and in-place state`; `B.6.1 Rows, columns, and shrinking
  boundaries`; `B.6.2 Layer transformations`; `B.6.3 Reusing the first row and
  column as markers`; `B.6.4 Keeping current and next state together`; `B.7
  Composing data structures`; `B.8 Streaming and online processing`; and `B.9
  Choosing a structure`.
- Method selection and pattern cards: `C. Problem-Solving Patterns`; `C.1
  Problem-solving method selection guide`; `C.1.1 Quick decision tree`; `C.1.2
  Derive a direction before naming a pattern`; `C.1.3 Pattern cards`; the 33
  headings `C.1.3.1 Common-prefix scan` through `C.1.3.33 Trie`; the nested
  heading `C.1.3.10.1 Coordinate compression`; `C.1.4 Choosing among
  overlapping methods`; `C.1.5 Let constraints eliminate bad choices`; and
  `C.1.6 Before committing to a method`. The exact card names are repeated in
  the coverage matrix below.
- Reference implementations: `D. Go reference implementations`; `D.1 Binary
  search bounds`; `D.2 Rolling dynamic programming`; `D.3 Hash-map complement
  lookup`; `D.4 Generic singly linked list`; `D.5 Generic stack`; `D.6
  Overflow-safe integer reversal`; `D.7 Common-prefix scan`; `D.8 Directional
  scan`; and `D.9 Sorted two-pointer pair search`.
- Go techniques and further study: `E. Go techniques`; `E.1 Keep input
  conversion outside the algorithm`; `E.2 Table-driven tests`; `E.3 Sorting
  without unexpected mutation`; `E.4 Map presence and sets`; `E.5 Priority
  queues`; `E.6 Bytes, runes, and visible characters`; `E.7 Type parameters`;
  `E.8 Replacing recursion with a loop`; `E.9 Slice storage and retained
  memory`; and `F. Further-study checklist`.

**Legacy anchors and inbound links:** The document has 10 explicit anchors:
`integer-digit-processing`, `linked-lists`, `common-prefix-scan`,
`directional-scan`, `hash-map-or-set`, `binary-search`, `two-pointers`,
`dynamic-programming`, `stack`, and `sliding-window`. All are used from other
repository files and must remain stable. The baseline has 20 inbound fragment
references distributed as follows:

- `binary-search` (4): `solved_easy_leetcode_0069_sqrtx.go`,
  `solved_medium_leetcode_0034_find_first_and_last_position_of_element_in_sorted_array.go`,
  `solved_medium_leetcode_0033_search_in_rotated_sorted_array.go`, and
  `solved_easy_leetcode_0035_search_insert_position.go`.
- `common-prefix-scan` (1):
  `solved_easy_leetcode_0014_longest_common_prefix.go`.
- `directional-scan` (1): `solved_easy_leetcode_0013_roman_to_integer.go`.
- `dynamic-programming` (2): `solved_medium_leetcode_0198_house_robber.go` and
  `solved_easy_leetcode_0070_climbing_stairs.go`.
- `hash-map-or-set` (2):
  `solved_medium_leetcode_0003_longest_substring_without_repeating_characters.go`
  and `solved_easy_leetcode_0001_two_sum.go`.
- `integer-digit-processing` (1):
  `solved_easy_leetcode_0009_palindrome_number.go`.
- `linked-lists` (1): `solved_easy_leetcode_0002_add_two_numbers.go`.
- `sliding-window` (1):
  `solved_medium_leetcode_0003_longest_substring_without_repeating_characters.go`.
- `stack` (1): `solved_easy_leetcode_0020_valid_parentheses.go`.
- `two-pointers` (6): `solved_easy_leetcode_0088_merge_sorted_array.go`,
  `solved_easy_leetcode_0001_two_sum.go`,
  `solved_easy_leetcode_0027_remove_element.go`, two references in
  `solved_medium_leetcode_0005_longest_palindromic_substring.go`, and
  `solved_easy_leetcode_0026_remove_duplicates_from_sorted_array.go`.

**Outbound-link inventory:** There are 172 Markdown links: 146 links to 84
unique fragments in this document, 11 external source links, and 15 local-file
links to 14 unique files. The local targets are
`../scripts/top-interview-150.tsv` and 13 representative problem files:
LeetCode 1, 2, 5, 9, 13, 14, 20, 21, 26, 27, 33, 35, and 70. The palindrome
problem is linked twice; every other local target is linked once. The external
sources are:

- LeetCode Top Interview 150, Grind 75, and the NeetCode roadmap;
- official Amazon, Meta, and Microsoft interview-preparation material; and
- the five linked Amazon, Google, Meta, Microsoft, and Uber candidate reports.

The exact URLs remain in `concepts/README.md`; this inventory deliberately
records both their role and count so a changed or removed source is visible.

**Examples, sources, and code blocks:** The baseline has 124 lines containing
`example`, `trace`, or `cases to test`. Of those, 49 begin an explicitly
labeled worked example, map, trace, example, or reference-implementation case
list. Their starting line numbers are `173, 209, 645, 719, 800, 874, 1118,
1316, 1629, 1667, 1706, 1745, 1799, 1855, 1904, 1980, 2024, 2106, 2225, 2277,
2328, 2385, 2419, 2462, 2483, 2570, 2654, 2705, 2746, 2792, 2834, 2872, 2956,
3008, 3052, 3083, 3117, 3165, 3204, 3234, 3438, 3469, 3499, 3548, 3607, 3638,
3667, 3703, 3733`. Unlabeled inline examples are retained by the broader
124-line inventory rather than being silently excluded.

There are 49 balanced fenced code blocks: 26 `text` blocks and 23 `go` blocks.
Their starting lines are `212, 311, 322, 335, 367, 397, 421, 440, 478, 528,
534, 542, 552, 565, 581, 593, 649, 673, 723, 760, 893, 938, 976, 987, 1020,
1040, 1051, 1084, 1121, 1147, 1221, 1237, 2157, 2272, 2486, 2498, 3392,
3444, 3474, 3504, 3553, 3612, 3643, 3672, 3708, 3766, 3795, 3808, 3883`.

**Reproduction commands:** Run these from the repository root against the
working-tree file:

```bash
awk '/^#{1,6} / {count++; level=match($0,/[^#]/)-1; by[level]++} END {print count; for (i=1;i<=6;i++) print i, by[i]}' concepts/README.md
rg -n '^#{1,6} |^<a (id|name)=' concepts/README.md
awk '/^```/ {fences++; if (!open) {open=1; lang=substr($0,4); openings[lang]++; blocks++} else {open=0}} END {print fences, blocks, open; for (lang in openings) print lang, openings[lang]}' concepts/README.md
rg -i -n 'example|trace|cases to test' concepts/README.md
rg -o --no-filename 'concepts/README\.md#[a-z0-9-]+' --glob '!concepts/README.md' --glob '!REFACTOR_CONCEPTS.md' --glob '!.git/**' . | sort | uniq -c
perl -0777 -ne 's/^```.*?^```\s*//msg; while (/\[[^\]]+\]\(([^)]+)\)/g) { print "$1\n" }' concepts/README.md | sort | uniq -c
```

**Ownership decisions:** Foundations own arithmetic, correctness, and resource
terms; data-structure entries own their operations; pattern cards own method
reasoning; Go-technique entries own language behavior; reference entries own
only implementation contracts, code, traces, costs, and cases. The method and
structure selection guides compare and link to owners without duplicating
their definitions. The matrix records each current entry's destination and,
where an implementation is not the teaching owner, names the owner explicitly.

**Changed headings:** None. Phase 0 did not edit `concepts/README.md`.

**Validation results:** All phase-0 checks passed on 2026-10-20:

- `check-non-ascii.py REFACTOR_CONCEPTS.md concepts/README.md`: passed.
- `git diff --check HEAD -- REFACTOR_CONCEPTS.md concepts/README.md`: passed.
- `rg -n '^#{1,6} ' concepts/README.md`: confirmed the recorded 119-heading
  hierarchy.
- Fence scan: 98 markers forming 49 balanced blocks.
- Local-target scan: all 15 local links resolve.
- Fragment scan: all 146 internal links resolve to a generated heading
  fragment or an explicit preserved anchor.
- Matrix shape check: 93 teaching-entry rows, each with all 18 columns.

### Deferred work

- The external interview-evidence claims and URLs are time-sensitive. Preserve
  them through content phases, then reverify them during the final review and
  validation phases rather than changing them in phase 0.
- The 23 Go blocks are inventoried but intentionally not compiled in phase 0;
  phases 5 and 6 own compilation after their examples are made complete.

## Coverage matrix

`U` means the baseline material has not yet been judged against that required
field; `C` will mean the field is complete. Editorial and technical review stay
`-` until their distinct phase-8 passes. Group headings, the evidence section,
and the further-study checklist are in the heading inventory but are not
eleven-field teaching entries.

| Entry | Destination / primary owner | Original heading or anchor | Phase | Problem | Idea | How | Examples | Why | Go | Work | Use | Avoid | Mistakes | Related | Editorial | Technical | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| A.1 | Foundations / complexity | A.1 Complexity and resource accounting | 2 | U | U | U | U | U | U | U | U | U | U | U | - | - | Resource-accounting owner |
| A.2 | Foundations / correctness | A.2 Correctness and invariants | 1 | U | U | U | U | U | U | U | U | U | U | U | - | - | Pilot; invariant owner |
| A.3 | Foundations / recursion | A.3 Recursion and saved work | 2 | U | U | U | U | U | U | U | U | U | U | U | - | - | Owns call-stack costs |
| A.4 | Foundations / integer bounds | A.4 Integer limits and bits | 2 | U | U | U | U | U | U | U | U | U | U | U | - | - | Integer-representation owner |
| A.5 | Foundations / bit operations | A.5 Bit manipulation | 2 | U | U | U | U | U | U | U | U | U | U | U | - | - | Bit-operation owner |
| A.6 | Foundations / exact representation | A.6 Keeping arithmetic values exact and consistent | 2 | U | U | U | U | U | U | U | U | U | U | U | - | - | Normalization owner |
| A.6.1 | Foundations / GCD and LCM | A.6.1 Shared divisors and multiples (GCD and LCM) | 2 | U | U | U | U | U | U | U | U | U | U | U | - | - | Exact-arithmetic child |
| A.6.2 | Foundations / fractions | A.6.2 Normalized fractions and ratios | 2 | U | U | U | U | U | U | U | U | U | U | U | - | - | Exact-arithmetic child |
| A.6.3 | Foundations / slopes | A.6.3 Exact slopes without rounding | 2 | U | U | U | U | U | U | U | U | U | U | U | - | - | Exact-arithmetic child |
| A.6.4 | Foundations / modular arithmetic | A.6.4 Modular arithmetic and negative division | 2 | U | U | U | U | U | U | U | U | U | U | U | - | - | Exact-arithmetic child |
| A.6.5 | Foundations / factorial zeroes | A.6.5 Trailing zeroes without calculating a factorial | 2 | U | U | U | U | U | U | U | U | U | U | U | - | - | Factor-counting owner |
| A.6.6 | Foundations / factors and sieves | A.6.6 Factors and prime sieves | 2 | U | U | U | U | U | U | U | U | U | U | U | - | - | Factoring owner |
| A.6.7 | Foundations / large integers | A.6.7 Numbers too large for an integer type | 2 | U | U | U | U | U | U | U | U | U | U | U | - | - | Large-number owner |
| A.7 | Foundations / combinatorics | A.7 Combinatorics and output size | 2 | U | U | U | U | U | U | U | U | U | U | U | - | - | Output-size owner |
| A.7.1 | Foundations / combinations | A.7.1 Counting combinations without large factorials | 2 | U | U | U | U | U | U | U | U | U | U | U | - | - | Combination-count owner |
| A.7.2 | Foundations / repeated values | A.7.2 Repeated input values | 2 | U | U | U | U | U | U | U | U | U | U | U | - | - | Duplicate-choice owner |
| A.8 | Foundations / digit processing | A.8 Integer Digit Processing; `integer-digit-processing` | 2 | U | U | U | U | U | U | U | U | U | U | U | - | - | Legacy anchor is inbound |
| B.1 | Store information / arrays, slices, strings | B.1 Arrays, slices, and strings | 3 | U | U | U | U | U | U | U | U | U | U | U | - | - | Owns basic sequence storage |
| B.1.1 | Store information / in-place sequence mutation | B.1.1 Moving values in place | 3 | U | U | U | U | U | U | U | U | U | U | U | - | - | Sequence-operation child |
| B.2 | Store information / linked lists | B.2 Linked Lists; `linked-lists` | 3 | U | U | U | U | U | U | U | U | U | U | U | - | - | Legacy anchor is inbound |
| B.3 | Store information / queues and deques | B.3 Queues and deques | 3 | U | U | U | U | U | U | U | U | U | U | U | - | - | Owns FIFO and double-ended operations |
| B.4.1 | Store information / tree terms and BST search | B.4.1 Terminology and search | 3 | U | U | U | U | U | U | U | U | U | U | U | - | - | Tree/BST owner |
| B.4.2 | Store information / tree traversal | B.4.2 Traversal order | 3 | U | U | U | U | U | U | U | U | U | U | U | - | - | Traversal-order owner |
| B.4.3 | Store information / recursive tree contracts | B.4.3 Recursive return contracts | 3 | U | U | U | U | U | U | U | U | U | U | U | - | - | Tree-return-contract owner |
| B.4.4 | Store information / BST validation | B.4.4 BST validation and costs | 3 | U | U | U | U | U | U | U | U | U | U | U | - | - | BST-validation owner |
| B.4.5 | Store information / tree reconstruction | B.4.5 Reconstruction from traversals | 3 | U | U | U | U | U | U | U | U | U | U | U | - | - | Reconstruction owner |
| B.4.6 | Store information / parent-link traversal | B.4.6 Traversing through parent links | 3 | U | U | U | U | U | U | U | U | U | U | U | - | - | Parent-link owner |
| B.5 | Store information / graphs and grids | B.5 Graphs and grids | 3 | U | U | U | U | U | U | U | U | U | U | U | - | - | Graph representation owner |
| B.6 | Store information / matrix state | B.6 Matrix traversal and in-place state | 3 | U | U | U | U | U | U | U | U | U | U | U | - | - | Matrix-coordinate owner |
| B.6.1 | Store information / shrinking boundaries | B.6.1 Rows, columns, and shrinking boundaries | 3 | U | U | U | U | U | U | U | U | U | U | U | - | - | Matrix child |
| B.6.2 | Store information / layer transforms | B.6.2 Layer transformations | 3 | U | U | U | U | U | U | U | U | U | U | U | - | - | Matrix child |
| B.6.3 | Store information / in-place markers | B.6.3 Reusing the first row and column as markers | 3 | U | U | U | U | U | U | U | U | U | U | U | - | - | Matrix child |
| B.6.4 | Store information / encoded matrix state | B.6.4 Keeping current and next state together | 3 | U | U | U | U | U | U | U | U | U | U | U | - | - | Matrix child |
| B.7 | Store information / composed structures | B.7 Composing data structures | 3 | U | U | U | U | U | U | U | U | U | U | U | - | - | Cross-structure invariant owner |
| B.8 | Store information / streaming state | B.8 Streaming and online processing | 3 | U | U | U | U | U | U | U | U | U | U | U | - | - | Online-processing owner |
| B.9 | Store information / selection guide | B.9 Choosing a structure | 3 | U | U | U | U | U | U | U | U | U | U | U | - | - | Comparison only; link to owners |
| C.1.1 | Solve problems / quick selection | C.1.1 Quick decision tree | 4 | U | U | U | U | U | U | U | U | U | U | U | - | - | Selection guide, not method owner |
| C.1.2 | Solve problems / derive a direction | C.1.2 Derive a direction before naming a pattern | 4 | U | U | U | U | U | U | U | U | U | U | U | - | - | Selection guide, not method owner |
| C.1.3.1 | Solve problems / common-prefix scan | C.1.3.1 Common-prefix scan; `common-prefix-scan` | 1 | U | U | U | U | U | U | U | U | U | U | U | - | - | Pilot; legacy anchor is inbound |
| C.1.3.2 | Solve problems / directional scan | C.1.3.2 Directional scan; `directional-scan` | 4 | U | U | U | U | U | U | U | U | U | U | U | - | - | Legacy anchor is inbound |
| C.1.3.3 | Solve problems / simulation | C.1.3.3 Simulation | 4 | U | U | U | U | U | U | U | U | U | U | U | - | - | Method owner |
| C.1.3.4 | Solve problems / hash map or set | C.1.3.4 Hash map or set; `hash-map-or-set` | 4 | U | U | U | U | U | U | U | U | U | U | U | - | - | Legacy anchor is inbound |
| C.1.3.5 | Solve problems / sorting and selection | C.1.3.5 Sorting and selection | 4 | U | U | U | U | U | U | U | U | U | U | U | - | - | Method owner |
| C.1.3.6 | Solve problems / binary search | C.1.3.6 Binary search; `binary-search` | 4 | U | U | U | U | U | U | U | U | U | U | U | - | - | Legacy anchor is inbound |
| C.1.3.7 | Solve problems / string matching | C.1.3.7 String matching | 4 | U | U | U | U | U | U | U | U | U | U | U | - | - | Method owner |
| C.1.3.8 | Solve problems / parsing state | C.1.3.8 Parsing and state machines | 4 | U | U | U | U | U | U | U | U | U | U | U | - | - | Method owner |
| C.1.3.9 | Solve problems / two pointers | C.1.3.9 Two pointers; `two-pointers` | 4 | U | U | U | U | U | U | U | U | U | U | U | - | - | Legacy anchor is inbound |
| C.1.3.10 | Solve problems / dynamic programming | C.1.3.10 Dynamic programming; `dynamic-programming` | 4 | U | U | U | U | U | U | U | U | U | U | U | - | - | Legacy anchor is inbound |
| C.1.3.10.1 | Solve problems / coordinate compression | C.1.3.10.1 Coordinate compression | 4 | U | U | U | U | U | U | U | U | U | U | U | - | - | Specialized preprocessing owner |
| C.1.3.11 | Solve problems / divide and conquer | C.1.3.11 Divide and conquer | 4 | U | U | U | U | U | U | U | U | U | U | U | - | - | Method owner |
| C.1.3.12 | Solve problems / best subarray scan | C.1.3.12 Kadane's algorithm (best subarray scan) | 4 | U | U | U | U | U | U | U | U | U | U | U | - | - | Method owner |
| C.1.3.13 | Solve problems / stack method | C.1.3.13 Stack; `stack` | 4 | U | U | U | U | U | U | U | U | U | U | U | - | - | Legacy anchor is inbound |
| C.1.3.14 | Solve problems / sliding window | C.1.3.14 Sliding window; `sliding-window` | 4 | U | U | U | U | U | U | U | U | U | U | U | - | - | Legacy anchor is inbound |
| C.1.3.15 | Solve problems / prefix sums | C.1.3.15 Prefix sums | 4 | U | U | U | U | U | U | U | U | U | U | U | - | - | Method owner |
| C.1.3.16 | Solve problems / difference arrays | C.1.3.16 Difference arrays | 4 | U | U | U | U | U | U | U | U | U | U | U | - | - | Method owner |
| C.1.3.17 | Solve problems / Fenwick and segment trees | C.1.3.17 Fenwick and segment trees | 4 | U | U | U | U | U | U | U | U | U | U | U | - | - | Method owner |
| C.1.3.18 | Solve problems / greedy | C.1.3.18 Greedy | 4 | U | U | U | U | U | U | U | U | U | U | U | - | - | Method owner |
| C.1.3.19 | Solve problems / backtracking | C.1.3.19 Backtracking | 4 | U | U | U | U | U | U | U | U | U | U | U | - | - | Method owner |
| C.1.3.20 | Solve problems / depth-first search | C.1.3.20 Depth-first search | 4 | U | U | U | U | U | U | U | U | U | U | U | - | - | Method owner |
| C.1.3.21 | Solve problems / breadth-first search | C.1.3.21 Breadth-first search | 4 | U | U | U | U | U | U | U | U | U | U | U | - | - | Method owner |
| C.1.3.22 | Solve problems / topological sort | C.1.3.22 Topological sort | 4 | U | U | U | U | U | U | U | U | U | U | U | - | - | Method owner |
| C.1.3.23 | Solve problems / Dijkstra | C.1.3.23 Dijkstra shortest paths | 4 | U | U | U | U | U | U | U | U | U | U | U | - | - | Method owner |
| C.1.3.24 | Solve problems / Bellman-Ford | C.1.3.24 Bellman-Ford shortest paths | 4 | U | U | U | U | U | U | U | U | U | U | U | - | - | Method owner |
| C.1.3.25 | Solve problems / disjoint-set union | C.1.3.25 Disjoint-set union | 4 | U | U | U | U | U | U | U | U | U | U | U | - | - | Method owner |
| C.1.3.26 | Solve problems / minimum spanning tree | C.1.3.26 Minimum spanning tree | 4 | U | U | U | U | U | U | U | U | U | U | U | - | - | Method owner |
| C.1.3.27 | Solve problems / heap or priority queue | C.1.3.27 Heap or priority queue | 4 | U | U | U | U | U | U | U | U | U | U | U | - | - | Method owner |
| C.1.3.28 | Solve problems / monotonic stack | C.1.3.28 Monotonic stack | 4 | U | U | U | U | U | U | U | U | U | U | U | - | - | Method owner |
| C.1.3.29 | Solve problems / monotonic deque | C.1.3.29 Monotonic deque | 4 | U | U | U | U | U | U | U | U | U | U | U | - | - | Method owner |
| C.1.3.30 | Solve problems / intervals | C.1.3.30 Intervals | 4 | U | U | U | U | U | U | U | U | U | U | U | - | - | Method owner |
| C.1.3.31 | Solve problems / sweep line | C.1.3.31 Sweep line | 4 | U | U | U | U | U | U | U | U | U | U | U | - | - | Method owner |
| C.1.3.32 | Solve problems / fast and slow pointers | C.1.3.32 Fast and slow pointers | 4 | U | U | U | U | U | U | U | U | U | U | U | - | - | Method owner |
| C.1.3.33 | Solve problems / trie | C.1.3.33 Trie | 4 | U | U | U | U | U | U | U | U | U | U | U | - | - | Method owner |
| C.1.4 | Solve problems / method comparisons | C.1.4 Choosing among overlapping methods | 4 | U | U | U | U | U | U | U | U | U | U | U | - | - | Owns cross-method comparisons |
| C.1.5 | Solve problems / constraint filtering | C.1.5 Let constraints eliminate bad choices | 4 | U | U | U | U | U | U | U | U | U | U | U | - | - | Selection guide |
| C.1.6 | Solve problems / commitment checklist | C.1.6 Before committing to a method | 4 | U | U | U | U | U | U | U | U | U | U | U | - | - | Selection guide |
| D.1 | Reference implementations / binary search bounds | D.1 Binary search bounds | 6 | U | U | U | U | U | U | U | U | U | U | U | - | - | Teaching owner C.1.3.6 |
| D.2 | Reference implementations / rolling DP | D.2 Rolling dynamic programming | 6 | U | U | U | U | U | U | U | U | U | U | U | - | - | Teaching owner C.1.3.10 |
| D.3 | Reference implementations / complement lookup | D.3 Hash-map complement lookup | 6 | U | U | U | U | U | U | U | U | U | U | U | - | - | Teaching owner C.1.3.4 |
| D.4 | Reference implementations / linked list | D.4 Generic singly linked list | 6 | U | U | U | U | U | U | U | U | U | U | U | - | - | Teaching owner B.2; generics owner E.7 |
| D.5 | Reference implementations / stack | D.5 Generic stack | 6 | U | U | U | U | U | U | U | U | U | U | U | - | - | Teaching owner C.1.3.13; generics owner E.7 |
| D.6 | Reference implementations / integer reversal | D.6 Overflow-safe integer reversal | 6 | U | U | U | U | U | U | U | U | U | U | U | - | - | Teaching owners A.4 and A.8 |
| D.7 | Reference implementations / prefix scan | D.7 Common-prefix scan | 6 | U | U | U | U | U | U | U | U | U | U | U | - | - | Teaching owner C.1.3.1 |
| D.8 | Reference implementations / directional scan | D.8 Directional scan | 6 | U | U | U | U | U | U | U | U | U | U | U | - | - | Teaching owner C.1.3.2 |
| D.9 | Reference implementations / sorted pair scan | D.9 Sorted two-pointer pair search | 6 | U | U | U | U | U | U | U | U | U | U | U | - | - | Teaching owner C.1.3.9 |
| E.1 | Go-specific techniques / boundary conversion | E.1 Keep input conversion outside the algorithm | 5 | U | U | U | U | U | U | U | U | U | U | U | - | - | Owns boundary-conversion behavior |
| E.2 | Go-specific techniques / table tests | E.2 Table-driven tests | 5 | U | U | U | U | U | U | U | U | U | U | U | - | - | Owns table-test behavior |
| E.3 | Go-specific techniques / sorting mutation | E.3 Sorting without unexpected mutation | 5 | U | U | U | U | U | U | U | U | U | U | U | - | - | Owns Go sorting mutation behavior |
| E.4 | Go-specific techniques / maps and sets | E.4 Map presence and sets | 1 | U | U | U | U | U | U | U | U | U | U | U | - | - | Pilot; owns Go map behavior |
| E.5 | Go-specific techniques / priority queues | E.5 Priority queues | 5 | U | U | U | U | U | U | U | U | U | U | U | - | - | Owns `heap.Interface` behavior |
| E.6 | Go-specific techniques / string encoding | E.6 Bytes, runes, and visible characters | 5 | U | U | U | U | U | U | U | U | U | U | U | - | - | Owns Go text representation |
| E.7 | Go-specific techniques / type parameters | E.7 Type parameters | 5 | U | U | U | U | U | U | U | U | U | U | U | - | - | Owns generic type behavior |
| E.8 | Go-specific techniques / iterative recursion | E.8 Replacing recursion with a loop | 5 | U | U | U | U | U | U | U | U | U | U | U | - | - | Owns explicit-stack Go technique |
| E.9 | Go-specific techniques / slice retention | E.9 Slice storage and retained memory | 5 | U | U | U | U | U | U | U | U | U | U | U | - | - | Owns slice backing-array behavior |
