# Standards for Maintaining Concept Study Guides

Follow these standards whenever you change an active concept study guide:

- `dsa/CONCEPTS.md`
- `frontend/CONCEPTS.md`
- `backend/CONCEPTS.md`

Treat the version currently in the working tree as the source of truth. Keep its useful explanations, examples, code, sources, links, and legacy anchors unless you have verified that something needs to change.

## Line length

- Limit every line in an active concept guide to 70 Unicode characters.
- This limit applies to everything in the guide, including prose, headings, lists, links, comments, diagrams, and code examples.
- Format long URLs, paths, identifiers, and serialized values so that no line exceeds 70 characters.
- After changing a concept guide, run `go run ./internal/linewidth`.

## 1. Purpose and audience

- Write each guide consistently for its intended readers.

  Assume the following background in each guide:

  - DSA readers know basic Go syntax.
  - Frontend readers work with TypeScript and browser concepts.
  - Backend readers work with Go and web-service concepts.

- Teach ways of thinking that readers can reuse instead of collecting solutions to individual problems.
- Keep every entry technically accurate, but do not assume that readers already know its formal terms.
- Add a topic only when it meets a clear learning need, not simply to make the guide cover more material.

## 2. Guide organization

- Keep each domain's `CONCEPTS.md` as the main guide for that domain unless another change explicitly creates a different structure.
- Keep the DSA guide in this learning order:
  1. how to use the guide and current evidence about technical interviews
  2. foundations
  3. ways to store and organize information
  4. ways to solve problems
  5. Go-specific techniques
  6. Go reference implementations
  7. further study
- Keep the table of contents in sync with section names, order, and heading anchors.
- Place the DSA method-selection guidance before the pattern cards so readers learn how to choose among them first.
- Place DSA Go techniques before the reference implementations so readers learn the relevant language behavior before the code relies on it.
- Keep the frontend and backend sections in the order defined by their curriculum indexes.
- Treat evidence sections, competency maps, navigation groups, source references, and further-study lists as supporting material rather than teaching entries. They must follow the teaching-entry contract only when they become standalone active teaching entries.

## 3. Plain language and terminology

- Explain the main idea in everyday language before giving its standard technical name.
- Define each necessary technical term in the paragraph where it first appears. Do not explain one unfamiliar term with several other undefined terms.
- Use concrete verbs such as "visit," "remember," "compare," and "move." Avoid abstract phrases such as "perform traversal" and "leverage state."
- Explain notation in words before using it repeatedly.
- Before showing how an algorithm builds or changes state, explain what one representative value, table entry, lookup, pointer, or boundary means.
- Before explaining updates to indexed or mapped state, state its data contract as `lookup input -> stored fact`. Say what selects an entry, what unit its value measures, and how to read one concrete entry as a sentence.
- Keep logical quantities, storage locations, and stored values separate. When code converts a count, key, index, offset, or boundary, show the conversion with exact values before giving the general expression.
- If the implementation hides the natural mental model, explain the conceptual representation first. Then map it to the actual storage representation. Storage details should refine the mental model, not create a different one.
- When the same data has different roles or locations, label each use by its role and location. Do not let words such as "current," "copy," "position," or "value" quietly refer to different things.
- Give state a plain-language mental model when it makes the idea easier to remember. Follow it with the exact term and rule so the analogy does not replace a correct explanation.
- Use one main mental model throughout an entry's introduction, procedure, traces, correctness argument, and code. If another model is necessary, explain how it maps to the first one.
- Show before telling. Demonstrate a non-obvious rule with exact values, positions, or state changes before describing it in abstract terms.
- Keep standard names such as _binary search_, _dynamic programming_, and _Big-O_ visible. Readers need to recognize these names in problems, documentation, and interviews.
- Keep a precise technical term when removing it would reduce accuracy, but pair it with a plain-English explanation.
- Spell out each abbreviation the first time it appears in a learning context.

## 4. Required teaching-entry contract

Every active concept, data structure, problem-solving method, language technique, and reference implementation must use the following headings in this order. A section may be short, but none may be omitted.

1. **The problem it solves**
   - State the problem, the shape of the input, and the expected result.
   - Do not define the technique by repeating its name.
2. **How it works**
   - Use plain language and as much detail as the concept needs.
   - Do not remove a necessary rule, condition, state meaning, or distinction just to shorten the section.
   - Split long explanations into focused paragraphs when that makes them easier to follow.
   - Begin with a familiar action or practical situation when it makes the idea easier to picture.
   - Give the formal name only after the action is clear.
   - Explain the real workflow when an analogy would add complexity.
   - Remove filler, long introductions, and repeated points. Every sentence must add a rule, reason, condition, distinction, or useful mental model.
   - Put the full procedure in the worked trace unless the reader must understand a rule first.
3. **Step-by-step examples**
   - Show the complete method with at least one small, ordinary example.
   - Make the reason for choosing the concept visible in the example's workload, not only in the surrounding explanation.
   - Add a boundary, duplicate, failure, or adversarial example when it changes the behavior or reveals a common mistake.
4. **Why it works**
   - Use plain language and as much detail as the reasoning needs.
   - Begin with the concrete fact that stays true. Use the word _invariant_ only after that fact is clear.
   - For an algorithm, explain why the fact is true at the start, remains true after one step, and leads to the answer when the process ends. Also explain why the process ends.
   - For language, browser, data, or systems behavior, explain the relevant contract or lifecycle instead of forcing a loop proof.
   - State every assumption required by the reasoning.
   - Remove filler and do not replay transitions already clear in the trace. Every sentence must add a reason, condition, or link between the state and the result.
5. **Language-specific example**
   - Use the exact heading `Go example` in the DSA and backend guides.
   - Use the exact heading `TypeScript example` in the frontend guide. Use TSX when browser or framework behavior requires it.
   - Show the smallest complete code fragment that matches the trace, or link to the reference implementation that owns the full example.
6. **Work and memory**
   - Explain each cost in words before using standard notation.
   - Include preprocessing, recursion, output construction, runtime allocations, rendering, network, storage, and server costs when they make a meaningful difference.
7. **When to use it**
   - List clues a reader can see in the problem, input, or required output.
8. **When not to use it**
   - Explain which assumption fails or which cost is too high.
   - Name the closest alternative when it helps the reader choose.
9. **Common mistakes and edge cases**
   - Tie each warning to a specific failing input, invalid state change, or relevant language, browser, network, or runtime behavior.
10. **Related material**
    - Link to prerequisites, contrasting methods, the primary implementation, and representative repository problems. Do not repeat their content.

Use each required label as a Markdown heading exactly one level below the entry heading. For example, an `###` entry uses `#### The problem it solves`. A `#####` pattern card uses `###### The problem it solves`.

Give each required heading a separate teaching purpose. Do not repeat the same explanation under several headings. Let `How it works` build the mental model, the trace teach the procedure, and the correctness section justify it. Refer to the main explanation instead of restating it in every section.

## 5. Example and trace standards

- Give the exact input and expected output for every worked example.
- For an algorithm entry, make the ordinary worked trace and the language-specific example one continuous demonstration. Name a concrete problem, apply the code to the trace's exact input, and follow the code's actual representation, conversions, state changes, branches, and returned result. Do not trace a simplified mock example while the code solves a different problem or contract.
- Write each trace so a reader can reproduce it from the prose alone. For one continuous process, show `Initial state:` once. Then give every meaningful transition:
  - `Decision:` the action and why it is valid
  - `After:` the state produced by the action
  - `Known correct:` the part of the result already established
- The `After` state of one transition is the next transition's input. Do not repeat it as `Before`. Use `Before:` only for a standalone transition that does not continue from the preceding state.
- Keep trace labels and numbered step titles plain. Do not wrap them in bold emphasis; the raw Markdown must remain easy to scan.
- Use numbered lists for traces that change state. Use concise bullets for comparisons and mappings that do not change state.
- Teach the main path with one small, ordinary example. Put unusual branches, boundaries, failures, and adversarial behavior in focused follow-up examples when combining them would hide the main idea. Several examples may share the coverage; one input does not need to exercise every branch.
- For a table or other derived state, explain how to read one completed representative value before tracing how it was built.
- Before a trace starts with a derived collection or serialized state, derive every displayed entry from the original input. Do not present a completed array, table, tree, map, or serialized value without showing where its displayed values came from.
- Use the actual state and collection names in a trace, not vague labels such as "value" or "table." Include everything needed to replay the transition: the complete relevant collection after the step, the lookup key or index, the returned value, any unchanged retry state, and whether each displayed default is only initialized or already confirmed.
- Split a multi-phase method into named phases. When state changes roles between phases, give each role a distinct name. Use the same phase and state names in the prose, traces, and code.
- Use a compact, aligned text diagram when position, nesting, overlap, ownership, or another spatial relationship is much clearer visually than in prose.
- Show a stateful process as a sequence of aligned snapshots when possible. Keep the input, current positions, derived state, and output in stable locations. Change only the values or markers that actually change. Keep an important unchanged value visible when the step is meant to highlight it.
- Put a derived table, index, pointer, or status beside the data it describes. Put the active rule beside the transition it controls so the reader does not have to combine separate diagrams mentally.
- Show one active rule or branch at a time. Introduce another branch when the example reaches it, after the reader has concrete state to apply it to.
- When an optimized method removes visible wasted work, first show the baseline method doing that work once. The optimization should solve a problem the reader has already seen.
- Make the example workload support the stated reason for choosing the concept. If the concept helps with many, repeated, overlapping, interleaved, concurrent, or failure-prone operations, include enough of that condition to make the benefit visible. One operation that a simpler method handles equally well shows mechanics, not why the concept should be chosen.
- After the examples, the reader should be able to explain why the simpler alternative repeats work, loses required state, breaks a stated rule, or fails on the demonstrated input. If the ordinary trace does not show that difference, add a focused comparison or a second example. Do not rely only on an abstract complexity claim.
- Let the concrete example or diagram carry the explanation. Use prose to point out what matters instead of repeating the same relationship in abstract terms.
- Make the worked trace the main demonstration. Keep the prose before it focused, but include every fact needed to understand the trace.
- Do not narrate every movement already visible in a snapshot. Use prose to explain why the transition happens, which rule applies, or what the visual cannot show by itself.
- A visualization must show real input values and the state or transition being taught. Decorative diagrams and outcome-only images do not demonstrate the rule.
- Do not rely on color alone. Also use labels, position, brackets, or symbols so every state remains understandable without color.
- If a fallback, retry, propagation, or other branch can repeat, include a focused trace in which it repeats. Do not stop after its first pass.
- Build each focused example around the misunderstanding it should correct, not only the branch it should cover. Name the distinction, such as count versus position, key versus value, or initialized versus confirmed state.
- Keep each numbered step focused on one meaningful state change. Show a repeated rule once before compressing identical later transitions.
- An outcome-only paragraph is not a worked example.
- Do not use Markdown tables in teaching entries. They are hard to read on narrow screens and do not handle long state descriptions well.
- Start each warning with `> **Warning:**`, then describe the concrete failure and its correction.
- Use the same names in examples, traces, code, and analysis.

## 6. Correctness standards

- Describe the fact that stays true before calling it an _invariant_.
- For an algorithm, explain:
  1. why the invariant is true at the start
  2. why one iteration or recursive step keeps it true
  3. why the process ends
  4. why the final state contains the required answer
- For a data structure or language technique, explain why its operations or language rules produce the demonstrated result. For browser, framework, data, and operational topics, connect the behavior to the relevant contract or lifecycle. Do not force an algorithm-style proof when a precise behavior explanation is clearer.
- State every assumption required for correctness, such as sorted input, non-negative weights, valid encoding, acyclic relationships, or safe numeric bounds.
- Replay every changed trace against the changed code. The prose and code must make the same decisions in the same order.

## 7. Code example standards

- Put Go code in fenced `go` blocks. Put frontend code in fenced `typescript` or `tsx` blocks.
- Follow `CODING_RULES.md`, including its standard for comments written for learners.
- Make each example understandable on its own. The names, control flow, and nearby comments must explain every non-obvious state change. A reader should not need to search earlier prose for a missing step.
- Do not hide several reasoning steps in one unexplained expression. Explain index offsets, sentinel values, fallback lookups, bit operations, boundary updates, mutations, and compound conditions with named intermediate steps or an immediate invariant-based comment. Keep obvious language idioms concise.
- When code converts a logical quantity to an index, key, offset, or boundary, use role-based intermediate names if they clarify the conversion. The prose and trace must show the same input, conversion, lookup, and returned value.
- Match the code's phase boundaries and role-based state names to the explanation. Do not use one vague state name for different roles just because they share a type or update rule.
- Make each example complete enough to understand and compile in its stated context.
- State required input conditions immediately before the code.
- Explain non-obvious language and runtime behavior where the reader first needs it. Examples include Go aliasing and allocation, JavaScript coercion and task ordering, browser lifecycle behavior, and framework ownership rules.
- If the reference section already has a complete implementation, keep the full code there and link to it. Do not maintain a second copy.
- Format and compile or type-check every complete code example changed by an update. Use the repository's configured tools to run its ordinary case and at least one relevant boundary or failure case.

## 8. Work and memory standards

- Before writing a growth rate such as `O(n)`, define what each input-size symbol measures.
- Explain what work or stored data creates each cost before writing `O(...)`.
- List best, usual, and worst cases only when they differ in a meaningful way.
- Include preprocessing costs such as sorting and building an index.
- Account for recursion depth, explicit stacks and queues, maps, copied collections, retained backing storage, rendering, network transfer, storage, server load, and every other meaningful cost.
- When output size matters, separate extra working memory from the memory required to hold the output.
- Use separate symbols for important input dimensions when one `n` would hide their effects. Examples include rows, columns, vertices, edges, and output count.

## 9. Content ownership and duplication

- Give each idea one primary teaching entry.
- Use selection guides and competency maps to help readers choose an entry. Do not repeat the entry's full explanation in them.
- Put a data structure's basic representation and operations in its own entry. Pattern and application entries should explain how they use those operations.
- Put language, runtime, browser, and standard-library behavior in the relevant language-technique entry. Other entries should link to it when they depend on that behavior.
- Keep reusable algorithms and DSA-specific Go techniques in the DSA guide. Web entries should link to them instead of copying them.
- Keep a reference implementation's full code, contract, input requirements, variable-level trace, implementation costs, and concrete test cases together in its primary entry.
- Describe relationships between entries in one or two sentences, then link to the primary entry instead of copying its explanation.
- Before moving material, make sure its new primary entry is complete. Only then remove it from the old location.

## 10. Navigation and compatibility

- Keep standard algorithm, data-structure, browser, framework, and system names searchable in headings and the table of contents, even when a plain-English phrase comes first.
- Preserve every explicit legacy anchor and every heading anchor used elsewhere in the repository.
- Before renaming or moving a heading, search the entire repository for links to it. If old links must keep working, add an explicit HTML anchor for the old fragment.
- Add or update links in both directions only when they make navigation meaningfully better.
- After an edit, verify every table-of-contents link, local path, in-guide anchor, and repository link into the guide.
- Keep links to representative problems, projects, labs, and sources with the concept they demonstrate.
- Do not remove a useful source, attribution, example, or link without a specific, documented reason.

## 11. Evidence and source maintenance

- Keep time-sensitive claims about interviews, browsers, frameworks, services, and tools separate from explanations of lasting concepts.
- Check a time-sensitive claim against its linked sources before changing it.
- Use official preparation material for claims about a company's stated process whenever possible. Label candidate reports as anecdotal evidence.
- Do not use company tags or remembered question lists as proof of what a company currently asks.
- When replacing or removing a baseline external source for the DSA guide, update the checker under `dsa/internal/conceptsdoc/` in the same change. Explain why the new source supports the claim.

## 12. Update workflow

Before editing:

1. Read this standards file, the relevant curriculum index, the entire entry being changed, and the primary entries it links to. If the change involves code or code examples, also read `CODING_RULES.md`.
2. Check `git status` and the relevant diff. Preserve unrelated and user-authored changes. Never overwrite working-tree changes by restoring the guide from `HEAD`.
3. Find the subject's primary entry. Search for every related heading, anchor, link, code name, example, and problem reference that the change could affect.

While editing:

1. Make the smallest complete change that keeps the entry accurate.
2. When learner feedback reveals a failed explanation, find the earliest prerequisite the learner misunderstood and rebuild from there. Replace the failed mental model and remove outdated prose. Do not pile on more terminology, exceptions, and extra examples.
3. Complete all ten required sections when adding or substantially reorganizing an entry. Do not leave the conversion unfinished.
4. Change other sections only when consistency, ownership, or navigation requires it.
5. As soon as a heading is renamed, preserve its old link target with an explicit HTML anchor.
6. Update `dsa/internal/conceptsdoc/` when an intentional structural change affects the number of teaching entries or the required source set.

After editing:

1. Re-read each changed entry as a beginner. Look for undefined terms, vague pronouns, unexplained notation, and circular explanations. Make sure a reader can explain one representative state value before learning how it is built.
2. For each non-obvious representation, use a short prompt that asks a learner to paraphrase a concrete expression. For example, ask them to identify its lookup input and stored fact. If the answer requires terms the example has not introduced, revise the entry.
3. Make a separate technical pass. Replay every trace, compare it with the code, verify the correctness or lifecycle explanation, and recalculate each relevant cost.
4. Review the full diff for missing prose, examples, links, code, and sources. Also check for accidental changes outside the requested scope.
5. Read every fenced code example in each changed guide for local clarity. Formatting, compilation, and tests do not replace this review.
6. Make sure the ordinary example teaches the main path without relying on an edge-case trace. Use focused examples for added coverage instead of one input with too many branches. Compare the workload with `The problem it solves` and `When to use it`. The input must actually show the condition that justifies choosing the concept.
7. Make a density pass. Remove filler, narration of an adjacent visual, and content repeated from another required section. Keep every rule, reason, distinction, boundary, and check the reader needs, even when the explanation must remain long.
8. Run the validation commands below. Fix every failure before calling the update complete.

## 13. Validation

After every update to an active concept study guide, run these commands from the repository root:

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

Also format, compile or type-check, and run each changed complete code example in its stated environment or a temporary package. Remove all temporary files and other validation artifacts after the checks.

The concepts-document checker enforces the guide-specific structure. The repository line-width checker enforces the line limits defined here and in `CODING_RULES.md`. The concepts checker must continue to verify:

- a valid heading hierarchy with no duplicate targets
- balanced fenced blocks
- every teaching entry and its ten sections in the required order
- the ban on Markdown tables in teaching content
- complete initial-state chains or complete standalone transitions in every numbered item that uses the trace vocabulary
- local paths and fragments
- repository links into the guide
- the maintained set of external sources

These automated checks verify structure, links, formatting, and executable consistency. They cannot prove that a learner will form the intended mental model. The beginner and technical reading passes are still required.

## 14. Definition of done

An update to an active concept study guide is complete only when:

- every line in the changed guide is no more than 70 characters
- every changed teaching entry has all ten required sections in order
- each new term is explained in plain language before it is used
- representative state values have a concrete meaning before the trace shows how they are built
- every indexed or mapped state names its lookup input and stored fact and explains any conversion between a logical quantity and its storage location
- every non-obvious rule is shown with concrete state before it is summarized; diagrams use actual values when a relationship is much easier to see than describe
- traces use real identifiers, show all relevant state, and separate default or unconfirmed values from confirmed results
- chained traces show the initial state once and do not repeat one transition's `After` state as the next transition's `Before` state
- stateful explanations use stable, aligned snapshots, keep related state together, and show each active rule beside the step it controls
- every code example makes its non-obvious state changes clear through explicit names, control flow, or invariant-based comments
- ordinary and relevant boundary behavior appear in replayable traces; focused examples are used when a combined trace would hide the main idea
- each example workload shows the repeated work, overlap, interleaving, failure, state requirement, or other condition that makes the concept better than its simpler alternative
- the prose, traces, and code use the same phase boundaries and role-based state names
- one main mental model stays consistent throughout the entry, with an explicit mapping wherever the implementation uses a different representation
- each required section has a separate teaching purpose, and the prose does not repeat information already clear from a nearby visual or the primary entry
- the correctness or lifecycle explanation, code, and cost analysis agree
- content ownership is clear, and links replace duplicated teaching prose
- existing navigation still works
- no useful concept, example, attribution, source, or link was removed without a stated reason
- changed code examples compile or type-check and pass the cases used to exercise them
- every applicable validation command passes
