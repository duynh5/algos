# Maintenance Standards for Concept Study Guides

These standards apply whenever any active concept study guide is
changed:

- `dsa/CONCEPTS.md`;
- `frontend/CONCEPTS.md`; and
- `backend/CONCEPTS.md`.

The current working-tree version of each guide is the source of truth.
Preserve its useful teaching material, examples, code, sources, links,
and compatibility anchors unless a verified correction requires a
change.

## Physical line width

- Keep every physical line in each active concept guide at no more
  than 70 Unicode characters.
- Apply the limit to all guide content, including prose, headings,
  lists, links, comments, diagrams, and embedded code examples.
- Represent long URLs, paths, identifiers, and serialized values
  without creating a physical guide line over 70 characters.
- Run `go run ./internal/linewidth` after modifying a concept guide.

## 1. Purpose and audience

- Keep each guide a consistent learning resource for its stated
  audience. The DSA guide assumes basic Go syntax. The frontend guide
  uses TypeScript and browser concepts. The backend guide uses Go and
  web-service concepts.
- Teach reusable reasoning, not a catalog of answers to individual problems.
- Preserve technical precision while making every entry understandable without prior knowledge of its formal terminology.
- Add a subject because it fills a demonstrated learning need, not merely to increase coverage.

## 2. Guide organization

- Keep each domain's `CONCEPTS.md` as its single owning guide unless a
  separate change explicitly establishes a new structure.
- In the DSA guide, preserve the current learning progression:
  1. how to use the guide and contemporary interview evidence;
  2. foundations;
  3. ways to store and organize information;
  4. ways to solve problems;
  5. Go-specific techniques;
  6. Go reference implementations; and
  7. further study.
- Keep the table of contents synchronized with section names, order, and fragments.
- Keep the DSA method-selection material before the pattern cards it
  helps readers choose among.
- Keep DSA Go techniques before the reference-implementation section
  so language behavior is available before implementations rely on it.
- Preserve the frontend and backend progression defined by their
  curriculum indexes.
- Treat evidence sections, competency maps, navigational groups,
  source references, and further-study lists as supporting material.
  They do not need the teaching-entry contract unless promoted into
  standalone active teaching entries.

## 3. Plain language and terminology

- Lead with the underlying idea in ordinary language, then introduce its conventional technical name.
- Define every necessary technical term in the paragraph where it first appears. Do not define one unfamiliar term using several other undefined terms.
- Prefer concrete verbs such as "visit," "remember," "compare," and "move" over phrases such as "perform traversal" or "leverage state."
- Explain notation in words before relying on it repeatedly.
- Explain what one representative state value, table entry, lookup,
  pointer, or boundary means before explaining how the algorithm
  constructs or changes it.
- Before teaching updates to indexed or mapped state, state its data
  contract as `lookup input -> stored fact`. Name what selects an
  entry, what unit the stored value measures, and how to read one
  concrete entry as a sentence.
- Keep logical quantities, storage locations, and stored values
  distinct. When code converts between a count, key, index, offset,
  or boundary, show the conversion with exact values before using its
  general expression.
- If the implementation's representation hides the natural mental
  model, show the conceptual representation first, then map it to the
  actual storage representation. The implementation detail must
  refine the mental model rather than create it.
- When equal data appears in multiple roles or locations, label each
  occurrence by both role and location. Do not let terms such as
  "current," "copy," "position," or "value" silently refer to
  different things.
- Give state a plain-language mental model when that model reduces the
  amount of machinery a learner must remember. Follow it with the
  precise term and rule so the analogy does not replace correctness.
- Choose one primary mental model for an entry and carry it through
  the introduction, procedure, traces, correctness argument, and
  code. Do not switch explanatory models without explicitly mapping
  between them.
- Show before telling. Demonstrate a non-obvious rule with exact
  values, positions, or state changes before summarizing it in
  abstract language.
- Keep standard names such as *binary search*, *dynamic programming*, and *Big-O* visible because readers need to recognize them in problems, documentation, and interviews.
- Pair a precise technical term with a plain-English restatement when removing the term would make the explanation less accurate.
- Expand abbreviations on first use within the relevant learning context.

## 4. Required teaching-entry contract

Every active concept, data structure, problem-solving method,
language technique, and reference implementation must use these
headings in this order. A concise entry may keep a section short, but
it must not omit a required heading.

1. **The problem it solves**
   - State the reader's problem, the input shape, and the desired result.
   - Do not use the technique's name as its own explanation.
2. **How it works**
   - Use simple language and as much space as the concept needs.
   - Do not omit a required rule, condition, state meaning, or
     distinction merely to make the section shorter.
   - Split a longer explanation into focused paragraphs when that
     makes its ideas easier to follow.
   - Start with a familiar action or practical situation when it makes
     the idea easier to picture.
   - Name the formal idea only after the reader can picture the action.
   - Prefer the real workflow when an analogy would add more machinery
     than it removes.
   - Remove filler, throat-clearing, and restatements. Every sentence
     must add a rule, reason, condition, distinction, or useful mental
     model.
   - Leave the full procedure to the worked trace unless a rule must be
     understood first.
3. **Step-by-step examples**
   - Give at least one small, ordinary example of the complete method.
   - Add a boundary, duplicate, failure, or adversarial example whenever it changes behavior or exposes a common mistake.
4. **Why it works**
   - Use simple language and as much space as the reasoning needs.
   - Start with the concrete fact that remains true; introduce
     *invariant* only after that fact is understandable.
   - For an algorithm, explain why the fact starts true, survives one
     step, the process finishes, and the final state answers the
     question.
   - For language, browser, data, or systems behavior, explain the
     relevant contract or lifecycle instead of forcing a loop proof.
   - State every assumption the reasoning needs.
   - Remove filler and avoid replaying transitions already clear in
     the trace. Every sentence must add a reason, condition, or link
     from state to result.
5. **Language-specific example**
   - Use the literal `Go example` heading in the DSA and backend
     guides.
   - Use the literal `TypeScript example` heading in the frontend
     guide, including TSX when browser or framework behavior matters.
   - Provide the smallest complete fragment that matches the trace,
     or link to the one owning reference implementation.
6. **Work and memory**
   - Explain costs in words before conventional notation.
   - Include preprocessing, recursion, output construction, runtime
     allocations, rendering, network, storage, and server costs when
     they materially apply.
7. **When to use it**
   - List observable clues in the problem statement, input, or required output.
8. **When not to use it**
   - State which assumption fails or which cost becomes unacceptable.
   - Point to the closest alternative when that helps the reader choose.
9. **Common mistakes and edge cases**
    - Connect each warning to a concrete failing input, invalid state
      change, or relevant language, browser, network, or runtime
      behavior.
10. **Related material**
    - Link to prerequisites, contrasting methods, the owning implementation, and representative repository problems without repeating their content.

Use each required label as a Markdown heading exactly one level below the entry heading. For example, an `###` entry uses `#### The problem it solves`; a `#####` pattern card uses `###### The problem it solves`.

Treat the required headings as distinct teaching responsibilities,
not prompts to repeat the same explanation. Let `How it works` build
the complete mental model, the trace teach the procedure, and the
correctness section justify it. Refer back to the owning explanation
instead of rephrasing it in every section.

## 5. Example and trace standards

- Identify the exact input and expected output for every worked example.
- Make a trace replayable from the prose alone. At each meaningful state change, show:
  - **Before:** the current value, position, node, choice, and relevant state;
  - **Decision:** the action and the reason it is valid;
  - **After:** the resulting state; and
  - **Known correct:** the portion of the result already established.
- Use numbered lists for state-changing traces and concise bullets for static comparisons or mappings.
- Use one small ordinary example to teach the main path. Put unusual
  branches, boundaries, failures, and adversarial behavior in focused
  follow-up examples when combining them would obscure the main idea.
  Coverage may span several examples; one input need not exercise
  every branch.
- For a table or other derived state, show how to read one completed
  representative value before tracing its construction.
- In a state trace, use the real state and collection names rather
  than generic labels such as "value" or "table." Show every item
  needed to replay the transition, including the complete relevant
  collection after the step, the lookup key or index, the returned
  value, unchanged retry state, and whether displayed defaults are
  initialized or confirmed.
- Separate a multi-phase method into named phases. Give state distinct
  role-based names when its meaning changes between phases, and keep
  those phase and state names consistent in prose, traces, and code.
- Use a compact aligned text diagram when position, nesting, overlap,
  ownership, or another spatial relationship is materially clearer
  visually than in prose.
- For a stateful process, prefer a sequence of aligned snapshots that
  keeps the input, current positions, derived state, and output in
  stable locations. Change only the values or markers that actually
  change, and keep an important unchanged value visible when that is
  the point of the step.
- Place a derived table, index, pointer, or status beside the data it
  describes. Place the active rule beside the state transition it
  controls so the learner does not have to join separated diagrams
  mentally.
- Show one active rule or branch at a time. Introduce another branch
  when the example reaches it instead of displaying the entire control
  flow before the learner has a concrete state for it.
- When an optimized method exists to avoid visible wasted work, first
  show the baseline method performing that waste once. Let the
  optimization answer a problem the learner has already seen.
- Let the concrete example or diagram carry the explanation. Use
  prose to label what the learner should notice instead of restating
  the same relationship abstractly.
- Treat the worked trace as the main demonstration. Keep prose before
  it focused, but do not remove a fact the learner needs to understand
  the trace merely to reach the example sooner.
- Do not narrate every movement already visible in a snapshot. Prose
  should add the reason for the transition, the rule being applied, or
  a distinction the visual cannot express by itself.
- A visualization must encode real input values and the state or
  transition being taught. Decorative diagrams and outcome-only
  pictures do not count as evidence that the rule was demonstrated.
- Use labels, position, brackets, or symbols in addition to color so
  every state distinction remains understandable without color.
- If a fallback, retry, propagation, or other branch can repeat, add a
  focused trace in which it repeats. Do not stop the demonstration
  after its first iteration.
- Design focused examples around the misconception they must resolve,
  not only around branch coverage. State the distinction the example
  demonstrates, such as count versus position, key versus value, or
  initialized versus confirmed state.
- Keep each numbered step focused on one meaningful state transition.
  Compress repeated identical transitions only after showing the rule
  they repeat.
- Do not call an outcome-only paragraph a worked example.
- Do not use Markdown tables in teaching entries; narrow screens and long state descriptions make them harder to follow.
- Use `> **Warning:**` for a warning, followed by the concrete failure and its correction.
- Keep names in examples, traces, code, and analysis consistent.

## 6. Correctness standards

- Describe the fact that remains true before introducing the word *invariant*.
- For an algorithm, explain:
  1. why the invariant is true initially;
  2. why one iteration or recursive step preserves it;
  3. why the process terminates; and
  4. why the final state gives the required answer.
- For a data structure or language technique, explain why its
  operations or language behavior produce the demonstrated result.
  For browser, framework, data, and operational topics, explain why
  the behavior follows from the relevant contract or lifecycle. Do
  not force an algorithm-style proof where a precise behavior
  explanation is clearer.
- State assumptions that correctness depends on, such as sorted input, non-negative weights, valid encoding, acyclic relationships, or safe numeric bounds.
- Replay changed traces against the changed code. The prose and code
  must make the same decisions in the same order.

## 7. Code example standards

- Use fenced `go` blocks for Go code and fenced `typescript` or `tsx`
  blocks for frontend code.
- Follow `CODING_RULES.md`, including its learner-focused comment
  standard.
- Make every example locally understandable. A learner must be able to
  follow its non-obvious state changes from the names, control flow,
  and nearby comments without searching earlier prose for missing
  steps.
- Do not compress several reasoning steps into one unexplained
  expression. Explain index offsets, sentinel values, fallback
  lookups, bit operations, boundary updates, mutations, and compound
  conditions with named intermediate steps or an immediate
  invariant-based comment. Keep self-evident language idioms concise.
- When code converts a logical quantity into an index, key, offset, or
  boundary, use role-based intermediate names when they clarify the
  conversion. Make the prose and trace show the same input,
  conversion, lookup, and returned value.
- Make the code's phase boundaries and role-based state names match
  the explanation. Do not reuse one vague state name for different
  roles merely because the underlying type or update rule is similar.
- Keep examples complete enough to understand and compile in their stated context.
- State required input conditions immediately before the code.
- Explain non-obvious language and runtime behavior at the point where
  the reader needs it. This includes Go aliasing and allocation,
  JavaScript coercion and task ordering, browser lifecycle behavior,
  and framework ownership rules.
- If a complete implementation already exists in the reference section, keep that section as the implementation owner and link to it instead of maintaining a second copy.
- Format and compile or type-check every complete code example changed
  by an update. Exercise its ordinary case and at least one relevant
  boundary or failure case with the repository's configured tools.

## 8. Work and memory standards

- Define what each input-size symbol measures before using asymptotic notation.
- Explain the source of the work or memory before writing `O(...)`.
- State best, usual, and worst cases only when they materially differ.
- Count preprocessing such as sorting or index construction.
- Count recursion depth, explicit stacks or queues, maps, copied
  collections, retained backing storage, rendering work, network
  transfer, storage, server load, and other meaningful costs.
- Distinguish auxiliary space from required output space when output size is significant.
- Account for every input dimension when a single `n` would hide important behavior, such as rows, columns, vertices, edges, or output count.

## 9. Content ownership and duplication

- Give every idea one primary teaching entry.
- Let a selection guide or competency map help readers choose an
  entry; do not repeat an entry's full explanation there.
- Let a data-structure entry own its basic representation and
  operations. Pattern or application entries should explain how they
  use those operations.
- Let a language-technique entry own language, runtime, browser, or
  standard-library behavior. Other entries should link to it when
  that behavior matters.
- Let the DSA guide own reusable algorithms and DSA-specific Go
  techniques. Web entries should link to those owners instead of
  copying their explanations.
- Let a reference implementation own its full code, contract, required input conditions, variable-level trace, implementation costs, and concrete test cases.
- Summarize cross-entry relationships in one or two sentences and link to the owner instead of copying its explanation.
- When moving ownership, confirm the new owner is complete before removing material from the old location.

## 10. Navigation and compatibility

- Keep formal algorithm, data-structure, browser, framework, and
  systems names searchable in headings and in the table of contents,
  even when plain-English wording comes first.
- Preserve every explicit legacy anchor and every fragment referenced elsewhere in the repository.
- Before renaming or moving a heading, search the whole repository for inbound links. Add an explicit HTML anchor for the old fragment when existing links must remain valid.
- Update reciprocal links only when they materially improve navigation.
- Ensure every table-of-contents link, local path, in-guide fragment, and repository link into the guide resolves after the edit.
- Keep representative problem, project, lab, and source-reference
  links attached to the concept they demonstrate.
- Do not remove a useful source, attribution, example, or link without a specific, documented reason.

## 11. Evidence and source maintenance

- Keep time-sensitive claims about interviews, browsers, frameworks,
  services, and tools separate from timeless concept explanations.
- Verify time-sensitive claims against their linked sources before changing them.
- Prefer official preparation material for claims about a company's stated process. Label candidate reports as anecdotal evidence.
- Do not turn company tags or remembered question lists into claims about what is currently asked.
- When intentionally replacing or removing a DSA baseline external
  source,
  update the checker under `dsa/internal/conceptsdoc/` in the same
  change and
  explain why the new source supports the claim.

## 12. Update workflow

Before editing:

1. Read this standards file, the relevant curriculum index, the
   complete entry being changed, its linked owners, and
   `CODING_RULES.md` when code or code examples are involved.
2. Inspect `git status` and the relevant diff. Preserve unrelated and user-authored changes; never restore the guide from `HEAD` over working-tree work.
3. Identify the primary owner of the subject and search for headings, anchors, links, code names, examples, and related problem references affected by the change.

While editing:

1. Make the smallest complete change that keeps the entry accurate.
2. When learner feedback exposes a failed explanation, find the
   earliest misunderstood prerequisite and rebuild from there. Replace
   the failed model and remove superseded prose instead of layering on
   more terminology, exceptions, and supplemental examples.
3. Finish all ten required sections of an added or materially
   reorganized entry. Do not leave a partially converted entry.
4. Keep cross-section changes limited to corrections required for consistency, ownership, or navigation.
5. Preserve an old fragment with an explicit anchor as soon as its heading is renamed.
6. Update `dsa/internal/conceptsdoc/` when an intentional structural
   change alters the number of teaching entries or the maintained
   source contract.

After editing:

1. Re-read each changed entry from a beginner's perspective for undefined terms, vague pronouns, unexplained notation, and circular explanations.
   Confirm that the reader can explain one representative state value
   before following how it is constructed.
2. Test the entry with a short learner-paraphrase prompt for each
   non-obvious representation, such as identifying the lookup input
   and stored fact in one concrete expression. If the expected answer
   requires terminology not established by the example, revise it.
3. Perform a separate technical pass: replay each trace, compare it
   with the code, verify the correctness or lifecycle argument, and
   recalculate the applicable costs.
4. Review the complete diff for lost prose, examples, links, code, sources, and accidental changes outside the requested scope.
5. Inspect every fenced code example in each changed guide for local
   clarity. Formatting, compilation, and tests do not replace this
   review.
6. Check that the ordinary example teaches the main path without
   depending on an edge-case trace, and that added coverage uses
   focused examples instead of one branch-heavy input.
7. Perform a density pass. Remove filler and prose that merely narrates
   an adjacent visual or repeats another required section. Preserve
   every rule, reason, distinction, boundary, and verification the
   reader needs, even when doing so requires a longer explanation.
8. Run the validation commands below and resolve every failure before treating the update as complete.

## 13. Validation

Run these commands from the repository root after every update to an
active concept study guide:

```bash
~/.codex/scripts/check-non-ascii.py \
  CONCEPTS_STANDARDS.md \
  dsa/CONCEPTS.md \
  frontend/CONCEPTS.md \
  backend/CONCEPTS.md
git diff --check HEAD -- \
  CONCEPTS_STANDARDS.md \
  dsa/CONCEPTS.md \
  frontend/CONCEPTS.md \
  backend/CONCEPTS.md
go run ./dsa/internal/conceptsdoc
go run ./internal/codeexamples
go run ./internal/linewidth
go vet ./dsa/internal/conceptsdoc
go test ./dsa/internal/conceptsdoc
go test ./...
```

Also format, compile or type-check, and exercise each changed complete
code example in its stated environment or a temporary package. Remove
temporary validation artifacts after the checks.

The concepts-document checker is the executable structural contract
for guide-specific rules. The repository line-width checker enforces
the physical line limits defined here and in `CODING_RULES.md`. The
concepts checker must continue to verify:

- a valid heading hierarchy with no duplicate targets;
- balanced fenced blocks;
- all teaching entries and their ten ordered fields;
- the no-Markdown-table rule for teaching content;
- complete before, decision, after, and known-correct state in every
  numbered item that uses the trace vocabulary;
- local paths and fragments;
- repository links into the guide; and
- the maintained external-source set.

These automated checks establish structure, links, formatting, and
executable consistency. They do not establish that a learner can form
the intended mental model. The beginner and technical reading passes
remain required.

## 14. Definition of done

An update to an active concept study guide is complete only when:

- every line in the changed guide is at most 70 characters;
- every changed teaching entry follows the ten-part contract;
- new terminology is explained in plain language before it is relied on;
- representative state values are given concrete meaning before their
  construction mechanics are traced;
- indexed and mapped state declares its lookup input, stored fact, and
  any conversion between the learner's logical quantity and the
  implementation's storage location;
- non-obvious rules are demonstrated on concrete state before they
  are summarized, with actual-value diagrams for relationships that
  are materially easier to see than describe;
- traces use real identifiers, show complete relevant state, and
  distinguish default or unconfirmed values from confirmed results;
- stateful explanations use stable, aligned snapshots, keep related
  state together, and present the active rule at the step where it
  executes;
- every code example exposes its non-obvious state transitions locally
  through explicit names, control flow, or invariant-based comments;
- ordinary and relevant boundary behavior are demonstrated with
  replayable traces, using focused examples when one combined trace
  would hide the main idea;
- prose, traces, and code use matching phase boundaries and role-based
  state names;
- one primary mental model remains consistent across the entry, with
  explicit mappings wherever the implementation requires a different
  representation;
- each required section contributes a distinct teaching function, and
  prose does not duplicate information already clear from an adjacent
  visual or owning section;
- the correctness or lifecycle explanation, code, and cost analysis
  agree;
- content ownership is clear and duplicated teaching prose has been replaced with links;
- existing navigation remains compatible;
- no useful concept, example, attribution, source, or link was lost without an explicit reason;
- changed code examples compile or type-check and pass their exercised
  cases; and
- all applicable validation commands pass.
