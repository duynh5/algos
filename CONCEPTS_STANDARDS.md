# Maintenance Standards for `CONCEPTS.md`

These standards apply whenever `CONCEPTS.md` is changed. The current
working-tree version of the guide is the source of truth. Preserve its useful
teaching material, examples, Go code, sources, links, and compatibility anchors
unless a verified correction requires a change.

## 1. Purpose and audience

- Keep the guide a consistent learning resource for readers who know basic Go
  syntax but are new to algorithms and data structures.
- Teach reusable reasoning, not a catalog of answers to individual problems.
- Preserve technical precision while making every entry understandable without
  prior knowledge of its formal terminology.
- Add a subject because it fills a demonstrated learning need, not merely to
  increase coverage.

## 2. Guide organization

- Keep `CONCEPTS.md` as the single owning guide unless a separate change
  explicitly establishes a new documentation structure.
- Preserve the current learning progression:
  1. how to use the guide and contemporary interview evidence;
  2. foundations;
  3. ways to store and organize information;
  4. ways to solve problems;
  5. Go-specific techniques;
  6. Go reference implementations; and
  7. further study.
- Keep the table of contents synchronized with section names, order, and
  fragments.
- Keep the method-selection material before the pattern cards it helps readers
  choose among.
- Keep Go techniques before the reference-implementation section so language
  behavior is available before implementation details rely on it.
- Treat the interview-evidence section, navigational group sections, and the
  further-study list as supporting material. They do not need the teaching-entry
  contract below unless they are changed into standalone teaching entries.

## 3. Plain language and terminology

- Lead with the underlying idea in ordinary language, then introduce its
  conventional technical name.
- Define every necessary technical term in the paragraph where it first
  appears. Do not define one unfamiliar term using several other undefined
  terms.
- Prefer concrete verbs such as "visit," "remember," "compare," and "move"
  over phrases such as "perform traversal" or "leverage state."
- Explain notation in words before relying on it repeatedly.
- Keep standard names such as *binary search*, *dynamic programming*, and
  *Big-O* visible because readers need to recognize them in problems,
  documentation, and interviews.
- Pair a precise technical term with a plain-English restatement when removing
  the term would make the explanation less accurate.
- Expand abbreviations on first use within the relevant learning context.

## 4. Required teaching-entry contract

Every concept, data structure, problem-solving method, Go technique, and Go
reference implementation must use these headings in this order. A concise
entry may keep a section short, but it must not omit a required heading.

1. **The problem it solves**
   - State the reader's problem, the input shape, and the desired result.
   - Do not use the technique's name as its own explanation.
2. **The idea in simple terms**
   - Give a short mental picture or concrete analogy.
   - Introduce and define the formal name only after the idea is clear.
3. **How it works**
   - Describe the procedure as explicit actions.
   - Name the state the method keeps and what each part means.
   - State required input conditions before the procedure relies on them.
4. **Step-by-step examples**
   - Give at least one small, ordinary example of the complete method.
   - Add a boundary, duplicate, failure, or adversarial example whenever it
     changes behavior or exposes a common mistake.
5. **Why it works**
   - Explain why each step is safe, why the process finishes, and why the final
     state answers the question.
6. **Go example**
   - Provide the smallest complete Go fragment that matches the trace, or link
     to the one owning reference implementation.
7. **Work and memory**
   - Explain costs in words before conventional notation.
   - Include preprocessing, recursion, output construction, and meaningful Go
     allocations.
8. **When to use it**
   - List observable clues in the problem statement, input, or required output.
9. **When not to use it**
   - State which assumption fails or which cost becomes unacceptable.
   - Point to the closest alternative when that helps the reader choose.
10. **Common mistakes and edge cases**
    - Connect each warning to a concrete failing input, invalid state change, or
      relevant Go behavior.
11. **Related material**
    - Link to prerequisites, contrasting methods, the owning implementation,
      and representative repository problems without repeating their content.

Use each required label as a Markdown heading exactly one level below the entry
heading. For example, an `###` entry uses `#### The problem it solves`; a
`#####` pattern card uses `###### The problem it solves`.

## 5. Example and trace standards

- Identify the exact input and expected output for every worked example.
- Make a trace replayable from the prose alone. At each meaningful state
  change, show:
  - **Before:** the current value, position, node, choice, and relevant state;
  - **Decision:** the action and the reason it is valid;
  - **After:** the resulting state; and
  - **Known correct:** the portion of the result already established.
- Use numbered lists for state-changing traces and concise bullets for static
  comparisons or mappings.
- Do not call an outcome-only paragraph a worked example.
- Do not use Markdown tables in teaching entries; narrow screens and long state
  descriptions make them harder to follow.
- Use `> **Warning:**` for a warning, followed by the concrete failure and its
  correction.
- Keep names in examples, traces, code, and analysis consistent.

## 6. Correctness standards

- Describe the fact that remains true before introducing the word *invariant*.
- For an algorithm, explain:
  1. why the invariant is true initially;
  2. why one iteration or recursive step preserves it;
  3. why the process terminates; and
  4. why the final state gives the required answer.
- For a data structure or Go technique, explain why its operations or language
  behavior produce the demonstrated result. Do not force an algorithm-style
  proof where a precise behavior explanation is clearer.
- State assumptions that correctness depends on, such as sorted input,
  non-negative weights, valid encoding, acyclic relationships, or safe numeric
  bounds.
- Replay changed traces against the changed Go code. The prose and code must
  make the same decisions in the same order.

## 7. Go example standards

- Use fenced `go` blocks for Go code.
- Follow `CODING_RULES.md`, including its learner-focused comment standard.
- Keep examples complete enough to understand and compile in their stated
  context.
- State required input conditions immediately before the code.
- Explain non-obvious Go behavior at the point where the reader needs it,
  including slice aliasing, map presence, rune decoding, overflow, heap order,
  and allocation behavior.
- If a complete implementation already exists in the reference section, keep
  that section as the implementation owner and link to it instead of
  maintaining a second copy.
- Format and compile every complete Go example changed by an update. Exercise
  its ordinary case and at least one relevant boundary or failure case.

## 8. Work and memory standards

- Define what each input-size symbol measures before using asymptotic notation.
- Explain the source of the work or memory before writing `O(...)`.
- State best, usual, and worst cases only when they materially differ.
- Count preprocessing such as sorting or index construction.
- Count recursion depth, explicit stacks or queues, maps, copied slices,
  retained backing arrays, and other meaningful live allocations.
- Distinguish auxiliary space from required output space when output size is
  significant.
- Account for every input dimension when a single `n` would hide important
  behavior, such as rows, columns, vertices, edges, or output count.

## 9. Content ownership and duplication

- Give every idea one primary teaching entry.
- Let the method-selection guide help readers choose an entry; do not repeat an
  entry's full explanation there.
- Let a data-structure entry own its basic representation and operations.
  Pattern cards should explain how a method uses those operations.
- Let a Go-technique entry own Go language or standard-library behavior.
  Algorithm entries should link to it when that behavior matters.
- Let a reference implementation own its full code, contract, required input
  conditions, variable-level trace, implementation costs, and concrete test
  cases.
- Summarize cross-entry relationships in one or two sentences and link to the
  owner instead of copying its explanation.
- When moving ownership, confirm the new owner is complete before removing
  material from the old location.

## 10. Navigation and compatibility

- Keep formal algorithm and data-structure names searchable in headings and in
  the table of contents, even when plain-English wording comes first.
- Preserve every explicit legacy anchor and every fragment referenced elsewhere
  in the repository.
- Before renaming or moving a heading, search the whole repository for inbound
  links. Add an explicit HTML anchor for the old fragment when existing links
  must remain valid.
- Update reciprocal links only when they materially improve navigation.
- Ensure every table-of-contents link, local path, in-guide fragment, and
  repository link into the guide resolves after the edit.
- Keep representative problem links attached to the concept they demonstrate.
- Do not remove a useful source, attribution, example, or link without a
  specific, documented reason.

## 11. Evidence and source maintenance

- Keep claims about current interview practices separate from timeless concept
  explanations.
- Verify time-sensitive claims against their linked sources before changing
  them.
- Prefer official preparation material for claims about a company's stated
  process. Label candidate reports as anecdotal evidence.
- Do not turn company tags or remembered question lists into claims about what
  is currently asked.
- When intentionally replacing or removing a baseline external source, update
  `scripts/check-concepts-doc.go` in the same change and explain why the new
  source supports the claim.

## 12. Update workflow

Before editing:

1. Read this standards file, the complete entry being changed, its linked
   owners, and `CODING_RULES.md` when code or code examples are involved.
2. Inspect `git status` and the relevant diff. Preserve unrelated and
   user-authored changes; never restore the guide from `HEAD` over working-tree
   work.
3. Identify the primary owner of the subject and search for headings, anchors,
   links, code names, examples, and related problem references affected by the
   change.

While editing:

1. Make the smallest complete change that keeps the entry accurate.
2. Finish all eleven required sections of an added or materially reorganized
   entry. Do not leave a partially converted entry.
3. Keep cross-section changes limited to corrections required for consistency,
   ownership, or navigation.
4. Preserve an old fragment with an explicit anchor as soon as its heading is
   renamed.
5. Update `scripts/check-concepts-doc.go` when an intentional structural change
   alters the number of teaching entries or the maintained source contract.

After editing:

1. Re-read each changed entry from a beginner's perspective for undefined
   terms, vague pronouns, unexplained notation, and circular explanations.
2. Perform a separate technical pass: replay each trace, compare it with the Go
   code, verify the correctness argument, and recalculate work and memory.
3. Review the complete diff for lost prose, examples, links, code, sources, and
   accidental changes outside the requested scope.
4. Run the validation commands below and resolve every failure before treating
   the update as complete.

## 13. Validation

Run these commands from the repository root after every update to
`CONCEPTS.md`:

```bash
/Users/nolan/.codex/scripts/check-non-ascii.py CONCEPTS_STANDARDS.md CONCEPTS.md
git diff --check HEAD -- CONCEPTS_STANDARDS.md CONCEPTS.md
go run scripts/check-concepts-doc.go
go vet scripts/check-concepts-doc.go
go test scripts/check-concepts-doc.go
go test ./...
```

Also format, compile, and exercise each changed complete Go example in a
temporary package. Remove temporary validation artifacts after the checks.

The concepts-document checker is the executable structural contract. It must
continue to verify:

- a valid heading hierarchy with no duplicate targets;
- balanced fenced blocks;
- all teaching entries and their eleven ordered fields;
- the no-Markdown-table rule for teaching content;
- local paths and fragments;
- repository links into the guide; and
- the maintained external-source set.

## 14. Definition of done

An update to `CONCEPTS.md` is complete only when:

- every changed teaching entry follows the eleven-part contract;
- new terminology is explained in plain language before it is relied on;
- ordinary and relevant boundary behavior are demonstrated with replayable
  traces;
- the correctness explanation, Go code, and cost analysis agree;
- content ownership is clear and duplicated teaching prose has been replaced
  with links;
- existing navigation remains compatible;
- no useful concept, example, attribution, source, or link was lost without an
  explicit reason;
- changed Go examples compile and pass their exercised cases; and
- all applicable validation commands pass.
