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

Use a numbered list for traces. When several steps update the same fields, give
each step labeled `Before`, `Decision`, `After`, and `Known correct` details.
Use concise bullets for comparisons or mappings. Do not create Markdown tables,
and do not call an outcome-only paragraph a worked example.

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
| 1 | 0 | Entry contract, guide introduction, and three pilot entries | Complete |
| 2 | 1 | Foundations A.1-A.8 | Complete |
| 3 | 2 | Data structures B.1-B.9 | Complete |
| 4 | 3 | Method selection and all 33 pattern cards in C | Complete |
| 5 | 4 | All nine Go techniques in E | Complete |
| 6 | 5 | All nine Go reference implementations in D | Complete |
| 7 | 6 | Final organization, further study, navigation, and link preservation | Complete |
| 8 | 7 | Beginner-language and technical review of the complete guide | Complete |
| 9 | 8 | Repository-wide documentation and test validation | Complete |

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

### Phase 1: Entry contract, guide introduction, and pilot entries

**Status:** Complete

**Starting headings:** `Algorithms and Data Structures in Go`, `Contents`,
`Contemporary interview evidence`, `A.2 Correctness and invariants`,
`C.1.3.1 Common-prefix scan`, and `E.4 Map presence and sets`. The pilot's
inbound `common-prefix-scan` anchor is preserved.

**Entry standard:** Each teaching entry uses the eleven required headings in
the plan's order, one Markdown level below its entry heading. Repeated state
changes use a numbered list with labeled before state, decision and reason,
after state, and the fact now known to be correct. Comparisons use concise
bullets; teaching entries do not use Markdown tables.
Warnings use `> **Warning:**` and name a concrete failure. Complete Go blocks
reuse trace names; when another entry owns the implementation, the Go section
links to that owner. Related material is a short purpose-labeled link list.

**Changed headings:** Added `How to use this guide`, `Plain English before
technical names`, `Entry contract`, and `Forms used in every entry`.
`C.1.3.1 Common-prefix scan` became `C.1.3.1 Compare from the left
(common-prefix scan)`. `E.4 Map presence and sets` became `E.4 Distinguish
missing keys from stored values (map presence and sets)`. `A.2 Correctness and
invariants` kept its heading. The contents entry for E.4 now matches its new
heading.

**Preserved anchors:** The inbound `common-prefix-scan` anchor remains in
place. Added explicit `c131-common-prefix-scan` and
`e4-map-presence-and-sets` anchors to preserve the fragments formerly generated
by the two renamed headings. All changed local links and both representative
problem paths resolve.

**Ownership decisions:** A.2 owns invariant and correctness-proof structure.
C.1.3.1 owns the common-prefix method and links to D.7 as the sole owner of its
complete Go implementation. E.4 owns Go map presence, set representation, nil
map behavior, and unspecified map iteration order. The three phase-1 matrix
rows now mark all eleven content fields complete; phase-8 editorial and
technical review fields remain unapproved as required.

**Validation results:** All phase-1 checks passed on 2026-10-20:

- `check-non-ascii.py REFACTOR_CONCEPTS.md concepts/README.md`: passed after
  replacing the initial literal `U+00E9` with the Go escape `\u00e9`.
- `git diff --check HEAD -- REFACTOR_CONCEPTS.md concepts/README.md`: passed.
- `rg -n '^#{1,6} ' concepts/README.md`: confirmed the changed slices use the
  intended hierarchy and each pilot has all eleven headings in order.
- Fence scan: 100 markers forming 50 balanced blocks.
- Changed-link and local-target scan: all added or changed fragments and both
  representative problem paths resolve; renamed generated fragments are
  protected by explicit anchors.
- Matrix shape check: A.2, C.1.3.1, and E.4 each retain all 18 columns.
- Temporary example package: `gofmt -w` followed by
  `go test ./.tmp_phase1_examples` passed for the A.2 and E.4 code and the D.7
  implementation used by the C.1.3.1 trace. Tests covered ordinary, empty,
  all-zero, duplicate, missing-key, and early-ended-string cases. The temporary
  package was removed after validation.

**Diff audit:** The introduction's study paths and terminology policy were
retained in the new guide contract. All three pilots retain their original
ideas, examples, warnings, implementation link, source links, and
representative problem links while adding the missing trace, proof, usage,
rejection, and cost detail. Contemporary interview evidence and later-phase
entries were not changed.

### Phase 2: Foundations

**Status:** Complete

**Starting headings:** `A.1 Complexity and resource accounting`; `A.3
Recursion and saved work`; `A.4 Integer limits and bits`; `A.5 Bit
manipulation`; `A.6 Keeping arithmetic values exact and consistent`; `A.6.1
Shared divisors and multiples (GCD and LCM)`; `A.6.2 Normalized fractions and
ratios`; `A.6.3 Exact slopes without rounding`; `A.6.4 Modular arithmetic and
negative division`; `A.6.5 Trailing zeroes without calculating a factorial`;
`A.6.6 Factors and prime sieves`; `A.6.7 Numbers too large for an integer
type`; `A.7 Combinatorics and output size`; `A.7.1 Counting combinations
without large factorials`; `A.7.2 Repeated input values`; and `A.8 Integer
Digit Processing`, including its six field headings. The inbound
`integer-digit-processing` anchor must remain stable.

**Changed headings:** The A.1-A.8 entry headings remain unchanged. Every
foundation entry now has the eleven standard field headings in the required
order. The six former A.8 field headings were replaced by those standard
fields; explicit anchors preserve their generated fragments:
`a81-mental-model`, `a82-worked-example-reversing-digits`,
`a83-invariant-fact-that-remains-true`,
`a84-preventing-integer-overflow`, `a85-work-and-memory`, and
`a86-related-problems`.

**Preserved anchors:** The inbound `integer-digit-processing` anchor remains in
place. All pre-phase A.8 field fragments resolve through the explicit anchors
listed above. No other phase-2 entry heading changed, so its generated fragment
remains stable.

**Ownership decisions:** A.1 owns resource accounting; A.3 owns recursion and
call-stack costs; A.4 owns fixed-width integer bounds; A.5 owns bit operations;
A.6 owns exact representation and delegates concrete rules to A.6.1-A.6.7.
A.6.1 owns GCD and LCM, A.6.2 reduced fractions, A.6.3 direction keys, A.6.4
negative remainder and floor division, A.6.5 factorial factor layers, A.6.6
trial division and sieves, and A.6.7 streaming decimal remainder. A.7 owns
output-size accounting, A.7.1 exact combination counts, A.7.2 duplicate-value
choice rules, and A.8 arithmetic digit extraction. Complete implementations
owned by D.1 and D.6 remain linked rather than duplicated.

**Validation results:** All phase-2 checks passed on 2026-10-20:

- The common checks passed after each of the three ordered batches.
  `check-non-ascii.py` and `git diff --check HEAD` are clean, and the final
  heading scan confirms all 17 foundation matrix entries have exactly the 11
  required fields in the required order.
- The final fence scan reports 96 markers forming 48 balanced blocks. The two
  net block removals from the phase-1 checkpoint are accounted for by replacing
  small trace and formula fences with structured step lists or inline formulas;
  their content remains in the corresponding entries.
- All local file targets and all internal fragments resolve. The final matrix
  shape remains 93 teaching rows with 18 data columns, and every phase-2 field
  is marked `C` while phase-8 review columns remain `-`.
- Temporary Go packages for batches 1, 2, and 3 passed `go test` after `gofmt`.
  They covered ordinary and boundary behavior for every new complete Go
  fragment. A focused QA package also passed after adding regressions for a
  digit larger than its limit and non-nil empty subset and combination output.
  All temporary packages were removed.
- `go test ./...` passed for `internal/testrunner` and `workbench`.

**QA findings resolved:** Review found and corrected three concrete mismatches:
the decimal-append helper now rejects `digit > limit`, the recursion warning
now matches the shown broad base case, and the subset generators create a
non-nil empty row for the documented `[[]]` boundary result. No remaining QA
findings were identified; phase 8 still owns the independent beginner-language
and technical approval columns.

**Diff audit:** Every pre-phase foundation topic remains present: complexity
classes, recursion depth, overflow, bit identities, exact arithmetic, GCD and
LCM, fractions, slopes, signed remainder and division, factorial zeroes,
factoring and sieves, oversized integers, subsets, permutations, combinations,
Pascal's recurrence, duplicate values, and digit reversal. Original examples
were retained and expanded into replayable traces. The palindrome problem link
and all relevant cross-section links remain valid; new representative links
point only to existing repository files. No later-phase prose was rewritten.

### Phase 3: Data structures

**Status:** Complete

**Starting headings:** `B.1 Arrays, slices, and strings`; `B.1.1 Moving values
in place`; `B.2 Linked Lists`, including `B.2.1 Mental model` through `B.2.8
Related problems`; `B.3 Queues and deques`; `B.4 Trees and binary-search
trees`, including `B.4.1 Terminology and search` through `B.4.6 Traversing
through parent links`; `B.5 Graphs and grids`; `B.6 Matrix traversal and
in-place state`, including `B.6.1 Rows, columns, and shrinking boundaries`
through `B.6.4 Keeping current and next state together`; `B.7 Composing data
structures`; `B.8 Streaming and online processing`; and `B.9 Choosing a
structure`. The inbound `linked-lists` anchor must remain stable.

**Changed headings:** The B.1-B.9 entry headings remain unchanged. Every one of
the 19 data-structure matrix entries now has the eleven standard field headings
in the required order. The eight former B.2 field headings were replaced by
the standard fields; explicit anchors preserve their generated fragments:
`b21-mental-model`, `b22-core-operations`,
`b23-worked-example-reversing-a-list`,
`b24-invariants-facts-that-remain-true`, `b25-common-techniques`,
`b26-work-and-memory`, `b27-common-mistakes`, and `b28-related-problems`.

**Preserved anchors:** The inbound `linked-lists` anchor remains immediately
before B.2. The eight former B.2 field fragments listed above also resolve.
No other phase-3 entry heading changed. All baseline links from the B section
remain present, and all repository links into the guide still resolve.

**Ownership decisions:** B.1 owns sequence storage and storage sharing; B.1.1
owns in-place rearrangement; B.2 owns linked-node operations; and B.3 owns
FIFO, deque, and amortized queue behavior. B.4.1-B.4.6 separately own BST
search, traversal order, recursive return contracts, validation,
reconstruction, and parent-link traversal. B.5 owns graph representations and
implicit grids. B.6 owns coordinates and shape, while B.6.1-B.6.4 own
shrinking boundaries, square rotation, matrix markers, and simultaneous-state
encoding. B.7 owns cross-structure invariants, B.8 online state, and B.9 only
the comparison among structures; B.9 links to owners instead of repeating
their definitions.

**Validation results:** All phase-3 checks passed on 2026-10-20:

- The common checks passed after each ordered batch. The final
  `check-non-ascii.py` and `git diff --check HEAD` runs are clean, and the
  heading scan confirms all 19 entries have the required 11 fields at the
  correct level.
- The final fence scan reports 114 markers forming 57 balanced blocks. All 240
  internal links and all local file targets resolve. The 19 B matrix rows each
  retain 18 columns, mark all 11 content fields `C`, and leave the phase-8
  review columns `-`.
- Temporary Go packages were formatted and compiled after each batch. Tests
  replayed ordinary and boundary behavior for slice sharing, rotation, queue
  traversal, every tree example, graph construction, matrix operations,
  composed state, and streaming state. A focused endpoint test confirmed the
  BST validator accepts the platform's minimum and maximum `int` values while
  rejecting an ancestor-bound violation. Temporary packages were removed.
- `go test ./...` passed for `internal/testrunner` and `workbench`.

**QA findings resolved:** Technical review replaced finite BST sentinels that
rejected valid endpoint values with explicit optional bounds, added missing
input requirements to compact Go examples, restored the concrete adjacency
matrix and edge-list example, fixed two incorrect pattern-card fragments, and
restored both baseline navigation links. No remaining QA findings were
identified; phase 8 still owns the independent beginner-language and technical
approval columns.

**Diff audit:** Every pre-phase data-structure topic remains present: slice
length, capacity, shared storage, append allocation, retained memory, strings,
frequency slices, in-place rotation and partitioning, linked-list operations,
digit-list carry, queues, deques, two-stack amortization, tree terminology and
orders, recursive contracts, BST validation, reconstruction, parent links,
graph representations and properties, grid neighbors, all four matrix-state
methods, composed structures, online processing, and structure selection.
Their original examples, representative problem links, and cross-section links
were retained and expanded into replayable traces. No phase-4 or later entry
was rewritten.

**Presentation follow-up:** At the user's direction, all 40 Markdown tables in
`concepts/README.md` were reformatted without creating replacement tables.
Thirty-nine trace tables became 135 numbered steps with labeled `Before`,
`Decision`, `After`, and applicable correctness or work details. The B.9
structure comparison became 10 definition-style bullets. Every source cell was
preserved, including inline code containing vertical bars. The guide and plan
contracts now require numbered trace lists and concise comparison bullets and
forbid new Markdown tables in teaching entries. A full scan confirms the guide
contains zero Markdown table blocks.

### Phase 4: Method selection and pattern cards

**Status:** Complete

**Starting headings:** `C. Problem-Solving Patterns`; `C.1 Problem-solving
method selection guide`; `C.1.1 Quick decision tree`; `C.1.2 Derive a direction
before naming a pattern`; `C.1.3 Pattern cards`; the 33 card headings from
`C.1.3.1 Compare from the left (common-prefix scan)` through `C.1.3.33 Trie`;
the nested `C.1.3.10.1 Coordinate compression`; `C.1.4 Choosing among
overlapping methods`; `C.1.5 Let constraints eliminate bad choices`; and
`C.1.6 Before committing to a method`. The completed C.1.3.1 pilot is reviewed
for consistency but not reopened without a concrete defect. Inbound anchors
`common-prefix-scan`, `directional-scan`, `hash-map-or-set`, `binary-search`,
`two-pointers`, `dynamic-programming`, `stack`, and `sliding-window` must remain
stable.

**Changed headings:** C.1.1 is now `Ask what the input and output allow (quick
decision tree)`. C.1.3.2-C.1.3.33 now lead with their plain-English action and
retain the formal searchable name in parentheses where needed: process the
needed neighbor, follow state changes, remember facts by key, put related
values together, find a changing boundary, reuse a partial match, track legal
input modes, move coordinated positions, reuse smaller answers, solve
independent parts, keep the best ending section, finish the newest item, move a
consecutive range, subtract saved totals, mark range changes, repair range
summaries, make a permanent choice, try and undo choices, finish one route,
visit nearby positions, release tasks, finalize the cheapest route, ripple
cheaper routes, track groups by leader, connect everything cheaply, remove the
most urgent item, resolve nearest extremes, keep window extremes, combine
overlapping ranges, process change positions, use speed or a gap, and share
word beginnings. C.1.3.10.1 is now a peer card named `Preserve order with
compact ranks (coordinate compression)` so it can own all eleven required
fields. C.1.2 and C.1.4-C.1.6 retain their entry headings but now use the full
field structure.

**Preserved anchors:** All eight inbound legacy anchors named in the starting
notes remain. Explicit anchors also preserve the former generated fragments
for every renamed card, the former `c11-quick-decision-tree` fragment, and the
coordinate-compression fragment. A repository-wide fragment check found no
unresolved link or duplicate explicit anchor.

**Ownership decisions:** C.1.1 and C.1.2 own selection questions and the
direct-method hint ladder, not algorithm implementations. Each C.1.3 card owns
its method's recognition, procedure, correctness rule, costs, rejection rule,
and edge cases. C.1.3.10.1 owns rank-preserving coordinate compression. C.1.4
owns cross-method comparisons, C.1.5 owns constraint filtering, and C.1.6 owns
the pre-implementation proof checklist. Existing D and E owners remain the
single source for their full reference implementations or Go-specific behavior;
cards link to them instead of duplicating those owners.

**Validation results:** All phase-4 checks passed on 2026-10-20:

- The common `check-non-ascii.py`, `git diff --check HEAD`, and heading scan
  pass. The final C-section scan finds 39 teaching entries, each with the 11
  required fields in order; all 38 phase-4 matrix rows have 18 columns and all
  content fields marked `C`. The phase-1 common-prefix pilot remains complete.
- The guide has 168 balanced fence markers. All current local paths and
  fragments resolve, explicit anchors are unique, and every external or local
  link present in the staged pre-phase guide remains present.
- Twenty-seven Go blocks in section C were extracted into one temporary
  package, formatted, and compiled. A focused temporary test replayed the
  ordinary and boundary behavior for every new executable example; it passed.
- `go test ./...` passed for `internal/testrunner` and `workbench`.

**QA finding resolved:** The first behavioral replay exposed an overflow bug in
the Bellman-Ford example's signed addition guard: subtracting a negative edge
cost from `MaxInt` overflowed and blocked a valid relaxation. The example now
uses `int64`, states its representable-route-cost condition, and passes both the
negative-edge trace and negative-cycle check.

**Diff audit:** The selection guide still links every original method. Every
pre-phase method variant remains present, including wrapped binary search,
KMP and rolling hash, expression parsing, two-pointer merge and palindrome
forms, DP state families and compression, fast exponentiation, circular
Kadane, run compression, multi-source and bidirectional BFS, three-state DFS,
both topological-sort forms, stale Dijkstra entries, Bellman-Ford cycle
detection, path compression, Kruskal and Prim, multiway heap merging, weighted
interval scheduling, and trie wildcard behavior. Representative problem links
and all staged external and local links were retained.

### Phase 5: Go techniques

**Status:** Complete

**Starting headings:** `E. Go techniques`; `E.1 Keep input conversion outside
the algorithm`; `E.2 Table-driven tests`; `E.3 Sorting without unexpected
mutation`; `E.4 Distinguish missing keys from stored values (map presence and
sets)`; `E.5 Priority queues`; `E.6 Bytes, runes, and visible characters`;
`E.7 Type parameters`; `E.8 Replacing recursion with a loop`; and `E.9 Slice
storage and retained memory`. The completed E.4 pilot is reviewed only for
cross-link consistency. Its explicit `e4-map-presence-and-sets` anchor must
remain stable.

**Changed headings and anchors:** No entry heading changed. All generated
fragments therefore remain stable, and the explicit
`e4-map-presence-and-sets` pilot anchor remains in place. E.4's links to the
hash-map card, D.3 reference implementation, and invariant owner still resolve.

**Ownership decisions:** E.1 owns decoding, boundary validation, and result
encoding; E.2 owns named table-test behavior; E.3 owns copy-before-sort and Go
sorting mutation; E.4 remains the map-presence, set, and iteration-order owner;
E.5 owns `heap.Interface` mechanics and stale entries; E.6 owns byte, rune, and
visible-symbol distinctions; E.7 owns type parameters and constraints; E.8
owns explicit-stack conversion; and E.9 owns backing-array lifetime and
pointer-slot clearing. Algorithm entries continue to own selection and
correctness reasoning and link to these Go-specific owners.

**Validation results:** All phase-5 checks passed on 2026-10-20:

- The common `check-non-ascii.py`, `git diff --check HEAD`, and full heading
  scan pass. Each of the nine E entries has all 11 required fields in order;
  every E matrix row has 18 columns and all content fields marked `C`.
- The guide has 178 balanced fence markers and no Markdown table lines in
  section E. All current local paths and fragments resolve.
- The eight changed Go blocks were copied into a temporary package, formatted,
  compiled, and exercised against their ordinary and boundary examples with
  `go test ./.tmp_phase5_examples`; both ordered checkpoint runs passed. The
  temporary package was removed after validation.
- `go test ./...` passed for `internal/testrunner` and `workbench`.

**QA findings resolved:** The first phase-wide ASCII check found literal
`U+00E9` characters in E.6; the examples now use Go's `\u00e9` escape while
retaining the same UTF-8 behavior. The per-entry language review also defined
code points, stable sorting, and interface values at first use and replaced an
unnecessary binary-size abbreviation.

**Diff audit:** All pre-phase Go-technique material remains: typed boundary
conversion and JSON context; shared table-test logic and parsing failures;
copy-before-sort, mutation, and tie rules; the completed map pilot; heap `Less`
ordering and stale Dijkstra entries; byte, rune, and visible-symbol choices;
generic stacks and `any`; explicit recursion frames and right-before-left
preorder; shared slice storage, conditional append allocation, retained
backing arrays, and cleared pointer slots. Each former concise example now
appears in a replayable trace or cross-link, and no section outside E was
rewritten.

### Phase 6: Go reference implementations

**Status:** Complete

**Starting headings:** `D. Go reference implementations`; `D.1 Binary search
bounds`; `D.2 Rolling dynamic programming`; `D.3 Hash-map complement lookup`;
`D.4 Generic singly linked list`; `D.5 Generic stack`; `D.6 Overflow-safe
integer reversal`; `D.7 Common-prefix scan`; `D.8 Directional scan`; and `D.9
Sorted two-pointer pair search`.

**Changed headings and anchors:** No entry heading changed, so every generated
D-section fragment remains stable. The guide's 10 explicit legacy anchors are
unchanged, and the reference entries link to the stable pattern-card anchors
for common-prefix scan, directional scan, hash maps, binary search, two
pointers, dynamic programming, and stacks.

**Ownership decisions:** D.1-D.9 own their complete Go implementations,
contracts, required input conditions, variable-level traces, implementation
costs, and concrete cases to test. Their linked A, B, C, and E entries remain
the owners for foundation rules, data-structure behavior, method selection,
and Go-specific behavior. The D introduction now states this boundary instead
of repeating the teaching entries.

**Validation results:** All phase-6 checks passed on 2026-10-20:

- The common `check-non-ascii.py`, `git diff --check HEAD`, and heading scan
  pass. D.1-D.9 contain exactly 99 required field headings, 11 per entry; all
  nine matrix rows have 18 columns and all content fields marked `C`.
- The guide has 178 balanced fence markers, section D contains no Markdown
  tables, and every current local path and fragment resolves.
- A comparison with the pre-phase staged document confirms that all nine Go
  implementations remain exact, excluding blank fence padding. The copied
  implementations and their ordinary and boundary examples were formatted and
  exercised together in a temporary package with
  `GOCACHE=/tmp/algos-phase6-go-build go test ./.tmp_phase6_examples`.
- `GOCACHE=/tmp/algos-phase6-go-build go test ./...` passed for the temporary
  examples, `internal/testrunner`, and `workbench`. The temporary package was
  removed after validation.

**QA findings resolved:** The technical pass corrected the lower-bound proof
so it treats `len(values)` as a sentinel rather than an in-range value. It also
corrected the Palindrome Number cross-link to claim the shared digit-extraction
invariant, not overflow guarding that the linked narrower-range solution does
not perform. An initial test attempt could not write the default Go build cache
under sandboxing; redirecting `GOCACHE` to `/tmp` resolved the environment
failure without changing code.

**Diff audit:** Every original contract, implementation, and listed case is
retained. Each former case list now appears as executable test data, a full
trace, or a concrete edge-case warning. No D heading, source, or pre-existing
link was removed, no section outside D was rewritten, and the pre-phase staged
state was restored after validation tooling attempted to stage the phase edits.

### Phase 7: Final organization and navigation

**Status:** Complete

**Starting headings:** `How to use this guide`; `A. Foundations`; `B. Data
structures`; `C. Problem-Solving Patterns`; `D. Go reference implementations`
with entries D.1-D.9; `E. Go techniques` with entries E.1-E.9; and `F.
Further-study checklist`. The table of contents lists reference implementations
before Go techniques. All generated D and E fragments, the explicit
`e4-map-presence-and-sets` anchor, the 10 existing explicit legacy anchors,
and every repository link into `concepts/README.md` must remain valid.

**Changed headings:** B is now `B. Ways to store and organize information`, C
is now `C. Ways to solve problems`, and F is now `F. Further study`. The former
E Go-technique section and entries are now D and appear before the former D
reference section and entries, which are now E. The contents use the new order,
names, and canonical fragments; its C.1.1 label now matches the completed entry
heading.

**Preserved anchors:** Explicit anchors preserve the five renamed top-level
section fragments and all 18 former D/E entry fragments. The existing
`e4-map-presence-and-sets` anchor remains, and the new short
`d4-map-presence-and-sets` anchor keeps guide links readable. All earlier
explicit anchors remain unchanged. Guide-owned links use current fragments,
while repository links using the old fragments continue to resolve.

**Ownership decisions:** D.1-D.9 retain ownership of the Go-specific behavior
established when those entries were in section E. E.1-E.9 retain ownership of
the complete reference implementations established when those entries were in
section D. The coverage matrix now uses current D/E entry numbers while its
original-heading column records the pre-move names. Section F remains a
next-step list rather than an eleven-field teaching entry; each advanced topic
now has a plain-language description and links to its closest prerequisite.

**Validation results:** All phase-7 checks passed on 2026-10-20:

- `check-non-ascii.py`, `git diff --check HEAD`, and the complete heading scan
  pass. The guide retains 1,132 headings and 178 fence markers forming 89
  balanced blocks.
- A temporary link audit verified every local file, every current fragment,
  every repository link into the guide, and all phase-7 legacy fragments. The
  audit script was removed after validation.
- The pre-phase and current documents both contain 11 external links. A
  unique-line diff accounts for every removed line as a renamed heading,
  renumbered in-guide D/E link, or prose replaced by the expanded further-study
  list. No code block, source, teaching entry, or useful example was lost.
- `GOCACHE=/tmp/algos-phase7-go-build go test ./...` passed for
  `internal/testrunner` and `workbench`.

**QA findings resolved:** The first navigation audit found that the mechanical
D/E link swap had changed the stable pattern fragments `dynamic-programming`
and `directional-scan`; both were restored before completion. The audit also
confirmed that the initial remaining reports were parser false positives from
Go syntax and trailing prose punctuation, not broken links. Editing tooling
temporarily staged the phase changes, so the two Git index entries were restored
to their exact pre-phase blob IDs; phase-7 work remains unstaged and the user's
existing staged work is preserved.

**Diff audit:** Go techniques now precede reference implementations without
changing their teaching prose or code. All D/E cross-references were renumbered
to their current owners, and legacy anchors cover the old public fragments.
The former one-paragraph advanced-topic list is now six concise descriptions
with prerequisite links. No foundation, data-structure, pattern-card,
Go-technique, or reference-implementation content was otherwise rewritten.

### Phase 8: Beginner-language and technical review

**Status:** Complete

**Starting headings:** `A. Foundations`; `B. Ways to store and organize
information`; `C. Ways to solve problems`; `D. Go-specific techniques`; `E.
Go reference implementations`; and `F. Further study`. The review covers all
93 teaching entries in the coverage matrix, in the three ordered batches named
by the phase definition. Existing heading names and explicit anchors are the
starting navigation contract.

**Review results:** All three ordered batches received separate
beginner-language and technical passes. All 93 matrix rows are approved in both
columns. The review defined first-use terms such as JSON, API, LRU, min-heap,
max-heap, BFS, and DFS; replaced repeated placeholder explanations with
method-specific procedures, correctness arguments, and failing examples; and
removed duplicate or outcome-only prose that repeated a complete trace or
related-material list.

**Technical findings resolved:** Exact-slope wording now distinguishes an
undirected line orientation from a directed ray. Parent-link traversal stops
when its frontier is empty, preserving the stated `O(n)` bound for very large
distances. Simulation rejects unsupported commands. Signed-prefix parsing
checks the asymmetric `int64` limits before multiplication. Kruskal stops after
accepting `V-1` edges. String matching, divide and conquer, prefix structures,
graph methods, heaps, monotonic structures, pointer methods, and tries now give
method-specific preservation and termination arguments instead of generic
claims. Bellman-Ford states its checked-sum precondition explicitly.

**Changed headings and anchors:** None. All existing headings and explicit
anchors remain the navigation contract. The three batches changed explanatory
prose and four complete Go examples only; ownership remains with the entries
already recorded in the matrix.

**Validation results:** `check-non-ascii.py` and `git diff --check HEAD` pass.
The complete heading scan reports 1,132 headings; each of the eleven required
entry headings appears 93 times. All 178 fence markers form 89 balanced code
blocks. A prose-only local-link audit reports 467 links and no unresolved file
or fragment target. A temporary Go test compiled and exercised the changed
simulation, parsing, parent-link traversal, and Kruskal examples, including
failure and boundary cases; the temporary artifact was removed. The external
interview-evidence links and their nearby claims were rechecked on 2026-10-20;
no correction was required. `GOCACHE=/tmp/algos-phase8-go-build go test ./...`
passes for `internal/testrunner` and `workbench`.

**QA and diff audit:** The QA pass found no unresolved correctness, navigation,
or regression issue after the fixes above. The phase preserves every teaching
entry, source, heading, explicit anchor, and complete example. It removes only
redundant trace/link prose and replaces vague text with concrete teaching
material. At that checkpoint, phase 9 remained not started.

### Phase 9: Repository-wide documentation and test validation

**Status:** Complete

**Starting headings:** The complete phase-8 navigation contract: 1,132
headings from `Algorithms and Data Structures in Go` through `F. Further
study`, including all 93 teaching entries and their eleven required field
headings. The ordered heading inventory is reproducible with
`rg -n '^#{1,6} ' concepts/README.md`; no heading or anchor change is planned
unless final validation finds a concrete defect.

**Changed headings and anchors:** None. The final guide retains 1,132 headings
and all 84 explicit anchors. The only guide-content correction expands ASCII
at first use; it does not change an entry heading, fragment, example, or
algorithm.

**Validation tooling:** Added `scripts/check-concepts-doc.go` because earlier
one-off link scans produced parser false positives. The command ignores fenced
and inline code, reproduces GitHub-style heading fragments, rejects heading
level jumps and duplicate targets, verifies the 93 ordered teaching-entry
contracts, checks balanced fences and the no-table rule, resolves local and
internal links, verifies repository links into the guide, and requires every
baseline external source URL. `README.md` now lists the command in the
repository validation sequence.

**Validation results:** All phase-9 checks passed on 2026-10-20:

- `/Users/nolan/.codex/scripts/check-non-ascii.py REFACTOR_CONCEPTS.md
  concepts/README.md README.md scripts/check-concepts-doc.go`: passed.
- `git diff --check HEAD -- REFACTOR_CONCEPTS.md concepts/README.md README.md`
  and `git diff --no-index --check /dev/null scripts/check-concepts-doc.go`:
  passed.
- `go run scripts/check-concepts-doc.go`: reports 1,132 headings, 84 explicit
  anchors, 89 balanced fenced blocks (78 Go and 11 text), 415 internal links,
  41 local-file links, all 11 baseline external links, all 20 repository
  inbound links, and 93 entries with all 11 fields in order.
- A temporary extractor formatted and compiled all 78 Go examples in isolated
  packages; the temporary source and generated packages were removed.
- `go vet scripts/check-concepts-doc.go` and
  `go test scripts/check-concepts-doc.go` pass.
- `GOCACHE=/tmp/algos-phase9-go-build go test ./...` passes for
  `internal/testrunner` and `workbench`.

**Baseline and language audit:** The final inventories meet or exceed the
recorded baseline while preserving its exact externally owned material: 20
inbound repository links remain 20, all 11 source URLs remain present, local
links grew from 15 to 41, internal links grew from 146 to 415, fenced blocks
grew from 49 to 89, and lines naming examples, traces, or cases grew from 124
to 275. The coverage matrix still contains all 93 owners. The placeholder and
abbreviation scan found only contextual uses after expanding ASCII at first
use; no deferred item affects correctness or navigation.

**QA and diff audit:** The QA pass corrected multi-backtick inline-code
handling and replaced an insufficient external-link count check with exact
baseline URL verification. After those fixes, no unresolved correctness,
navigation, data-loss, or test issue remains. The complete phase notes account
for every moved section, removed duplicate explanation, retained source,
example, owner, and link from the recorded baseline.

### Deferred work

- None affecting correctness or navigation after phase 9.

## Coverage matrix

`U` means the baseline material has not yet been judged against that required
field; `C` means the field or review is complete. Group headings, the evidence
section, and the further-study checklist are in the heading inventory but are
not eleven-field teaching entries.

| Entry | Destination / primary owner | Original heading or anchor | Phase | Problem | Idea | How | Examples | Why | Go | Work | Use | Avoid | Mistakes | Related | Editorial | Technical | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| A.1 | Foundations / complexity | A.1 Complexity and resource accounting | 2 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; resource-accounting owner |
| A.2 | Foundations / correctness | A.2 Correctness and invariants | 1 | C | C | C | C | C | C | C | C | C | C | C | C | C | Pilot complete; invariant and correctness-argument owner |
| A.3 | Foundations / recursion | A.3 Recursion and saved work | 2 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; owns call-stack costs |
| A.4 | Foundations / integer bounds | A.4 Integer limits and bits | 2 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; integer-representation owner |
| A.5 | Foundations / bit operations | A.5 Bit manipulation | 2 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; bit-operation owner |
| A.6 | Foundations / exact representation | A.6 Keeping arithmetic values exact and consistent | 2 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; normalization owner |
| A.6.1 | Foundations / GCD and LCM | A.6.1 Shared divisors and multiples (GCD and LCM) | 2 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; GCD and LCM owner |
| A.6.2 | Foundations / fractions | A.6.2 Normalized fractions and ratios | 2 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; normalized-fraction owner |
| A.6.3 | Foundations / slopes | A.6.3 Exact slopes without rounding | 2 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; exact-direction owner |
| A.6.4 | Foundations / modular arithmetic | A.6.4 Modular arithmetic and negative division | 2 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; negative remainder and floor-division owner |
| A.6.5 | Foundations / factorial zeroes | A.6.5 Trailing zeroes without calculating a factorial | 2 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; factor-layer counting owner |
| A.6.6 | Foundations / factors and sieves | A.6.6 Factors and prime sieves | 2 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; trial-division and sieve owner |
| A.6.7 | Foundations / large integers | A.6.7 Numbers too large for an integer type | 2 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; streaming decimal-remainder owner |
| A.7 | Foundations / combinatorics | A.7 Combinatorics and output size | 2 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; output-size owner |
| A.7.1 | Foundations / combinations | A.7.1 Counting combinations without large factorials | 2 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; exact combination-count owner |
| A.7.2 | Foundations / repeated values | A.7.2 Repeated input values | 2 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; duplicate-choice owner |
| A.8 | Foundations / digit processing | A.8 Integer Digit Processing; `integer-digit-processing` | 2 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; inbound and former field anchors preserved |
| B.1 | Store information / arrays, slices, strings | B.1 Arrays, slices, and strings | 3 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; basic sequence-storage owner |
| B.1.1 | Store information / in-place sequence mutation | B.1.1 Moving values in place | 3 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; in-place sequence-operation owner |
| B.2 | Store information / linked lists | B.2 Linked Lists; `linked-lists` | 3 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; inbound and former field anchors preserved |
| B.3 | Store information / queues and deques | B.3 Queues and deques | 3 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; FIFO, deque, and amortized-queue owner |
| B.4.1 | Store information / tree terms and BST search | B.4.1 Terminology and search | 3 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; tree-term and BST-search owner |
| B.4.2 | Store information / tree traversal | B.4.2 Traversal order | 3 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; traversal-order owner |
| B.4.3 | Store information / recursive tree contracts | B.4.3 Recursive return contracts | 3 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; tree-return-contract owner |
| B.4.4 | Store information / BST validation | B.4.4 BST validation and costs | 3 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; ancestor-bound validation owner |
| B.4.5 | Store information / tree reconstruction | B.4.5 Reconstruction from traversals | 3 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; traversal-reconstruction owner |
| B.4.6 | Store information / parent-link traversal | B.4.6 Traversing through parent links | 3 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; parent-link traversal owner |
| B.5 | Store information / graphs and grids | B.5 Graphs and grids | 3 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; graph-representation owner |
| B.6 | Store information / matrix state | B.6 Matrix traversal and in-place state | 3 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; matrix-coordinate and shape owner |
| B.6.1 | Store information / shrinking boundaries | B.6.1 Rows, columns, and shrinking boundaries | 3 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; shrinking-boundary owner |
| B.6.2 | Store information / layer transforms | B.6.2 Layer transformations | 3 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; square-rotation mapping owner |
| B.6.3 | Store information / in-place markers | B.6.3 Reusing the first row and column as markers | 3 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; in-place matrix-marker owner |
| B.6.4 | Store information / encoded matrix state | B.6.4 Keeping current and next state together | 3 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; simultaneous-state encoding owner |
| B.7 | Store information / composed structures | B.7 Composing data structures | 3 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; cross-structure invariant owner |
| B.8 | Store information / streaming state | B.8 Streaming and online processing | 3 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; online-processing owner |
| B.9 | Store information / selection guide | B.9 Choosing a structure | 3 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; comparison only, with owner links |
| C.1.1 | Solve problems / quick selection | C.1.1 Quick decision tree | 4 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; Selection guide, not method owner |
| C.1.2 | Solve problems / derive a direction | C.1.2 Derive a direction before naming a pattern | 4 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; Selection guide, not method owner |
| C.1.3.1 | Solve problems / common-prefix scan | C.1.3.1 Common-prefix scan; `common-prefix-scan` | 1 | C | C | C | C | C | C | C | C | C | C | C | C | C | Pilot complete; E.7 owns full code; inbound and former generated anchors preserved |
| C.1.3.2 | Solve problems / directional scan | C.1.3.2 Directional scan; `directional-scan` | 4 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; Legacy anchor is inbound |
| C.1.3.3 | Solve problems / simulation | C.1.3.3 Simulation | 4 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; Method owner |
| C.1.3.4 | Solve problems / hash map or set | C.1.3.4 Hash map or set; `hash-map-or-set` | 4 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; Legacy anchor is inbound |
| C.1.3.5 | Solve problems / sorting and selection | C.1.3.5 Sorting and selection | 4 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; Method owner |
| C.1.3.6 | Solve problems / binary search | C.1.3.6 Binary search; `binary-search` | 4 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; Legacy anchor is inbound |
| C.1.3.7 | Solve problems / string matching | C.1.3.7 String matching | 4 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; Method owner |
| C.1.3.8 | Solve problems / parsing state | C.1.3.8 Parsing and state machines | 4 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; Method owner |
| C.1.3.9 | Solve problems / two pointers | C.1.3.9 Two pointers; `two-pointers` | 4 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; Legacy anchor is inbound |
| C.1.3.10 | Solve problems / dynamic programming | C.1.3.10 Dynamic programming; `dynamic-programming` | 4 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; Legacy anchor is inbound |
| C.1.3.10.1 | Solve problems / coordinate compression | C.1.3.10.1 Coordinate compression | 4 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; Specialized preprocessing owner |
| C.1.3.11 | Solve problems / divide and conquer | C.1.3.11 Divide and conquer | 4 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; Method owner |
| C.1.3.12 | Solve problems / best subarray scan | C.1.3.12 Kadane's algorithm (best subarray scan) | 4 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; Method owner |
| C.1.3.13 | Solve problems / stack method | C.1.3.13 Stack; `stack` | 4 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; Legacy anchor is inbound |
| C.1.3.14 | Solve problems / sliding window | C.1.3.14 Sliding window; `sliding-window` | 4 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; Legacy anchor is inbound |
| C.1.3.15 | Solve problems / prefix sums | C.1.3.15 Prefix sums | 4 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; Method owner |
| C.1.3.16 | Solve problems / difference arrays | C.1.3.16 Difference arrays | 4 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; Method owner |
| C.1.3.17 | Solve problems / Fenwick and segment trees | C.1.3.17 Fenwick and segment trees | 4 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; Method owner |
| C.1.3.18 | Solve problems / greedy | C.1.3.18 Greedy | 4 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; Method owner |
| C.1.3.19 | Solve problems / backtracking | C.1.3.19 Backtracking | 4 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; Method owner |
| C.1.3.20 | Solve problems / depth-first search | C.1.3.20 Depth-first search | 4 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; Method owner |
| C.1.3.21 | Solve problems / breadth-first search | C.1.3.21 Breadth-first search | 4 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; Method owner |
| C.1.3.22 | Solve problems / topological sort | C.1.3.22 Topological sort | 4 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; Method owner |
| C.1.3.23 | Solve problems / Dijkstra | C.1.3.23 Dijkstra shortest paths | 4 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; Method owner |
| C.1.3.24 | Solve problems / Bellman-Ford | C.1.3.24 Bellman-Ford shortest paths | 4 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; Method owner |
| C.1.3.25 | Solve problems / disjoint-set union | C.1.3.25 Disjoint-set union | 4 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; Method owner |
| C.1.3.26 | Solve problems / minimum spanning tree | C.1.3.26 Minimum spanning tree | 4 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; Method owner |
| C.1.3.27 | Solve problems / heap or priority queue | C.1.3.27 Heap or priority queue | 4 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; Method owner |
| C.1.3.28 | Solve problems / monotonic stack | C.1.3.28 Monotonic stack | 4 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; Method owner |
| C.1.3.29 | Solve problems / monotonic deque | C.1.3.29 Monotonic deque | 4 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; Method owner |
| C.1.3.30 | Solve problems / intervals | C.1.3.30 Intervals | 4 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; Method owner |
| C.1.3.31 | Solve problems / sweep line | C.1.3.31 Sweep line | 4 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; Method owner |
| C.1.3.32 | Solve problems / fast and slow pointers | C.1.3.32 Fast and slow pointers | 4 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; Method owner |
| C.1.3.33 | Solve problems / trie | C.1.3.33 Trie | 4 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; Method owner |
| C.1.4 | Solve problems / method comparisons | C.1.4 Choosing among overlapping methods | 4 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; Owns cross-method comparisons |
| C.1.5 | Solve problems / constraint filtering | C.1.5 Let constraints eliminate bad choices | 4 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; Selection guide |
| C.1.6 | Solve problems / commitment checklist | C.1.6 Before committing to a method | 4 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; Selection guide |
| D.1 | Go-specific techniques / boundary conversion | E.1 Keep input conversion outside the algorithm | 5 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; moved to D.1 in phase 7; owns decoding, boundary validation, and result encoding |
| D.2 | Go-specific techniques / table tests | E.2 Table-driven tests | 5 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; moved to D.2 in phase 7; owns table-test and named-subtest behavior |
| D.3 | Go-specific techniques / sorting mutation | E.3 Sorting without unexpected mutation | 5 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; moved to D.3 in phase 7; owns Go sorting mutation and copy-before-sort behavior |
| D.4 | Go-specific techniques / maps and sets | E.4 Map presence and sets | 1 | C | C | C | C | C | C | C | C | C | C | C | C | C | Pilot complete; moved to D.4 in phase 7; owns Go map presence, set representation, and iteration-order behavior |
| D.5 | Go-specific techniques / priority queues | E.5 Priority queues | 5 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; moved to D.5 in phase 7; owns `heap.Interface`, root access, and stale-entry behavior |
| D.6 | Go-specific techniques / string encoding | E.6 Bytes, runes, and visible characters | 5 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; moved to D.6 in phase 7; owns byte, rune, and visible-symbol distinctions |
| D.7 | Go-specific techniques / type parameters | E.7 Type parameters | 5 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; moved to D.7 in phase 7; owns generic type parameters and constraints |
| D.8 | Go-specific techniques / iterative recursion | E.8 Replacing recursion with a loop | 5 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; moved to D.8 in phase 7; owns explicit-stack Go technique |
| D.9 | Go-specific techniques / slice retention | E.9 Slice storage and retained memory | 5 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; moved to D.9 in phase 7; owns slice backing-array lifetime and pointer-slot clearing |
| E.1 | Reference implementations / binary search bounds | D.1 Binary search bounds | 6 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; moved to E.1 in phase 7; implementation owner, teaching owner C.1.3.6 |
| E.2 | Reference implementations / rolling DP | D.2 Rolling dynamic programming | 6 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; moved to E.2 in phase 7; implementation owner, teaching owner C.1.3.10 |
| E.3 | Reference implementations / complement lookup | D.3 Hash-map complement lookup | 6 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; moved to E.3 in phase 7; implementation owner, teaching owner C.1.3.4 |
| E.4 | Reference implementations / linked list | D.4 Generic singly linked list | 6 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; moved to E.4 in phase 7; implementation owner, teaching owner B.2; generics owner D.7 |
| E.5 | Reference implementations / stack | D.5 Generic stack | 6 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; moved to E.5 in phase 7; implementation owner, teaching owner C.1.3.13; generics owner D.7 |
| E.6 | Reference implementations / integer reversal | D.6 Overflow-safe integer reversal | 6 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; moved to E.6 in phase 7; implementation owner, teaching owners A.4 and A.8 |
| E.7 | Reference implementations / prefix scan | D.7 Common-prefix scan | 6 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; moved to E.7 in phase 7; implementation owner, teaching owner C.1.3.1 |
| E.8 | Reference implementations / directional scan | D.8 Directional scan | 6 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; moved to E.8 in phase 7; implementation owner, teaching owner C.1.3.2 |
| E.9 | Reference implementations / sorted pair scan | D.9 Sorted two-pointer pair search | 6 | C | C | C | C | C | C | C | C | C | C | C | C | C | Complete; moved to E.9 in phase 7; implementation owner, teaching owner C.1.3.9 |
