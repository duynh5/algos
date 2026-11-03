# Concept Study Guide Standards

## Purpose and scope

These standards govern `dsa/CONCEPTS.md` and `backend/CONCEPTS.md`. Each guide must be consistent, understandable, and complete enough for a reader to apply its concepts correctly.

A lesson succeeds when the reader can explain the problem in everyday language, reproduce an actual example step by step, implement the method, explain why it works, recognize its limits, and handle the cases that commonly break it. Definitions, code, and links alone do not meet this standard.

Comprehensive means covering the concept's contract, mechanism, reasoning, implementation, costs, alternatives, and relevant boundaries. It does not mean mentioning every related topic or making every lesson equally long. Give difficult ideas the space they need; remove repetition rather than necessary reasoning.

## 1. Organize for learning and lookup

- Begin each guide with its audience, assumed knowledge, how to study it, and a table of contents. DSA assumes basic Go syntax; backend assumes basic Go and web-service concepts. Teach or link additional prerequisites before relying on them.
- Order lessons from foundations to applications. Teach method selection before catalogs of methods and language behavior before implementations depend on it. Follow domain curriculum indexes where present.
- Give each concept one primary teaching entry. Other entries explain their application and link back instead of maintaining competing explanations or code copies.
- Keep standard technical names searchable in headings and contents. Pair a familiar description with the established name when helpful.
- Distinguish complete lessons from outlines, roadmaps, evidence summaries, and further-study lists. Label planned material honestly; a heading alone is not coverage.
- Link reusable DSA explanations from other guides instead of copying them. Keep domain-specific behavior in the guide that owns it.
- Preserve useful explanations, examples, sources, attribution, and incoming links. Before renaming or moving headings, search repository references and preserve old fragments with explicit anchors where needed.

## 2. Use a consistent lesson sequence

Use these eight section headings, in order, for new or fully revised teaching entries. Place them one level below the concept heading. Add focused subheadings when needed without exceeding Markdown's six heading levels. Section length follows the concept's difficulty, not a fixed word count.

### Problem and contract

Explain what the reader needs to accomplish before naming the mechanism. State the input, expected output or observable behavior, and preconditions. Specify ordering, duplicate handling, mutation, ownership, valid ranges, failure behavior, and timing guarantees whenever they affect correctness.

Identify prerequisites and give a concrete reason to learn the concept. Establish a simple baseline when it helps reveal the repeated work, ambiguity, or failure being addressed.

### Mental model

Explain the idea in plain English. Identify the information remembered, what each part means, how the next action is chosen, and what ends the process. Introduce technical names alongside those meanings.

Define state before updating it. For a map, array, table, queue, or other structured representation, explain what selects an element and what fact its value stores. Read one concrete lookup as a sentence. Distinguish values from indices, indices from counts, and tentative results from confirmed ones.

For example, explain that a table entry stores the number of ways to reach a particular total before showing its update formula. If logical positions and storage positions differ, show one exact conversion before giving the general rule.

### Worked walkthrough

Demonstrate the complete process on actual values using the trace requirements below. Include an ordinary example and focused examples for relevant boundaries and failure behavior.

The main trace and implementation must describe the same problem, input, representation, decisions, and result. A simplified trace of a different algorithm is insufficient.

### Implementation

Provide readable Go code implementing the demonstrated contract, or link to the repository implementation that owns it. Identify the execution context, dependencies, assumptions, returned result, and side effects.

A link may avoid duplicate full code, but it cannot replace the local explanation or walkthrough. Follow `CODING_RULES.md` for code, comments, formatting, and validation.

### Why it works

Explain the fact that makes each decision safe. For algorithms, show why it holds initially, survives every kind of update, and implies the required result at termination. Explain why the process terminates and where each precondition matters.

For language, browser, framework, or service behavior, explain the applicable contract, lifecycle, or guarantee. State what it guarantees and what it leaves unspecified. Use the reasoning appropriate to the subject rather than forcing every lesson into a loop proof.

### Costs and tradeoffs

Explain the work performed and information stored before giving complexity notation. Define each size symbol and separate independent input dimensions. Include meaningful preprocessing, allocation, copying, recursion, output construction, retained storage, and external work. Separate auxiliary memory from output storage.

Distinguish best, average, amortized, and worst-case costs when meaningful, including assumptions behind average or expected costs. For systems topics, explain relevant latency, network requests, contention, capacity, resource lifetime, and operational costs rather than forcing everything into Big O notation.

### Edge cases and mistakes

For each relevant case, give a concrete input or event sequence, expected behavior, what a mistaken implementation does, and the rule that prevents the failure. Link to its focused trace or test when already shown.

Avoid generic warnings such as "watch out for boundaries." Identify the boundary, the branch it changes, and the exact mistake it exposes.

### When to choose it

Give recognizable clues favoring this concept, conditions ruling it out, and the closest useful alternatives. Explain the decisive difference through a workload or counterexample rather than declaring a method simply faster or better.

Finish with relevant links to prerequisites, contrasting concepts, the primary implementation, repository problems or labs, and authoritative sources. Links support a self-contained explanation; they do not substitute for teaching it.

## 3. Explain in plain English

- Start with concrete actions: compare values, move a boundary, remember a result, wait for a response. Introduce formal names after their meaning is clear.
- Define unfamiliar terms, abbreviations, symbols, units, and interval conventions at first use. Do not explain one unfamiliar term with several undefined terms.
- Explain why an action happens. "Move the left boundary because every earlier value is too small" teaches more than "update left."
- Use one consistent mental model. If an analogy helps, explain its correspondence and limits. Prefer the real workflow when an analogy adds complexity.
- Keep names consistent across prose, diagrams, traces, code, correctness reasoning, and cost analysis. Give different roles different names even when their values match.
- Demonstrate non-obvious expressions with exact values before generalizing. Expose index conversions, sentinels, fallback rules, and compound decisions.
- Use short, connected paragraphs. Give each section a separate teaching purpose. Avoid repeating the trace in the mental model or narrating every value already visible in a diagram.
- Distinguish guarantees, assumptions, heuristics, observations, and behavior that depends on a version or environment.
- Never shorten a lesson by removing a prerequisite, state meaning, decision rule, or failure case needed to reproduce it.

## 4. Make walkthroughs replayable

### Choose inputs that reveal the mechanism

Use small, explicit inputs with exact expected results. An ordinary example must show meaningful decisions, state changes, and the stopping condition. If the benefit depends on repeated queries, overlapping work, interleaving, or failures, include that condition.

Show the relevant wasted work or failure of a simpler approach before explaining an optimization. A complexity formula alone does not demonstrate why the concept helps.

Keep the main example readable. Use separate examples for unusual branches and boundaries when combining them would obscure the idea. Across the examples, cover every materially different decision rule. Compress identical repeated steps only after demonstrating the rule.

### Show each meaningful transition

Begin a continuous trace with the original input, expected result, and `Initial state:`. Derive any displayed preprocessing or initialized state from the input. Use numbered steps containing:

- `Decision:` the values inspected, condition evaluated, action chosen, and reason it is valid.
- `After:` the resulting relevant state, including changed boundaries, stored values, pending work, or output.
- `Known correct:` what has been established and what remains unresolved.

Each step starts from the preceding step's `After` state. Do not repeat it as `Before:`. Reserve `Before:` for independent transitions that also include their own decision, resulting state, and established fact. Show the final stopping check and returned value or observable outcome.

Include enough information to replay the process without running code: indices and values, active ranges, branch conditions, queue order, recursion arguments, restored state, or event order as appropriate. Distinguish initialized placeholders from computed answers. Name phases in multiphase methods and show their handoff.

When shrinking, retry, fallback, propagation, or backtracking can repeat, include a case where it repeats. Show what changes, what remains valid, and what bounds the repetition.

### Adapt traces to the subject

- Algorithms: follow reads, decisions, updates, and return values in implementation order.
- Data structures: trace interacting operations that reveal representation, ordering, and boundaries.
- Language behavior: show actual expressions, values, ownership or aliasing, and observable results.
- Backend: show requests or events, durable and temporary state, failure points, and externally visible outcomes. Identify scheduling and failure assumptions rather than implying every execution has one order.

Use aligned text snapshots when position, nesting, ownership, or boundaries are clearer visually. Keep related values together and positions stable. Label meaning explicitly rather than relying on color. Avoid Markdown tables in teaching entries so explanations remain readable on narrow screens.

An outcome-only paragraph, a decorative diagram, or a list of expected outputs is not a walkthrough.

## 5. Cover edge cases deliberately

Identify applicable categories for each entry. Explain why an apparently relevant case is outside the contract instead of silently ignoring it. Coverage follows behavior, not a fixed example quota.

- Input size: empty when valid, smallest valid input, one item, and large or maximum supported input.
- Values: duplicates, all equal, zero, negatives, missing targets, extreme values, overflow, and precision limits.
- Position and shape: first and last positions, exact boundaries, sorted or reversed order, skewed structures, disconnected components, cycles, and uneven dimensions.
- Transitions: no update, repeated update, ties, exhausted work, early exit, and conditions changing from valid to invalid or back again.
- Representation: nil versus empty, byte versus character, shared storage, mutation, ownership, invalid indices, and sentinel collisions.
- External behavior: malformed input, unavailable dependencies, partial success, duplicate delivery, reordered completion, timeout, cancellation, cleanup, and recovery.

Cover valid boundaries separately from invalid inputs. State whether invalid input is rejected, excluded by a precondition, or handled explicitly. Do not invent behavior the problem or platform does not promise.

Give a fully worked focused trace when a case introduces new reasoning. A compact input, expected result, and explanation is sufficient when only the data changes. For large constraint cases, use deterministic generated tests and analytical reasoning rather than printing thousands of transitions.

Include at least one counterexample to a plausible wrong approach. Identify its first incorrect decision and connect the correction to the contract or correctness rule. Preserve useful failures as regression cases when the repository has an executable home for them.

## 6. Keep code and sources trustworthy

- Make code complete in its stated context. Identify deliberate fragments and their missing context. Never present pseudocode as compiling code.
- Use language-tagged fences. Expose meaningful state changes through clear names, control flow, and nearby comments explaining the reasoning. Do not hide the lesson inside helpers or unrelated setup.
- Keep ordering, mutation, errors, and tie-breaking consistent with the contract and trace. State the acceptance rule when multiple results are valid.
- Keep full reference implementations in one primary location with their contract, trace, costs, and cases. Link applications to that location.
- Format and compile or type-check changed complete examples. Run ordinary cases and relevant boundaries or failures in the stated environment. Formatting alone does not establish correctness.
- Prefer authoritative documentation and original sources. Separate version-dependent and time-sensitive evidence from durable explanations. Verify changing claims before updating them and record the relevant version or verification date.
- Label interview reports as anecdotal. Remembered questions and platform tags do not establish a company's current process.
- Preserve attribution. Verify replacement baseline sources and update the checker contract in the same change.

## 7. Formatting and adoption

Every line in an active concept guide must be at most 70 Unicode characters, including headings, prose, links, diagrams, and code. Wrap text and use supported link layouts without breaking targets or shortening meaningful names. This root standards document follows the repository's unwrapped-prose convention.

Maintain balanced fences, a valid heading hierarchy, working paths and fragments, and a synchronized table of contents. Keep trace labels plain and use emphasis only for meaningful distinctions.

These standards replace the former ten-section teaching contract. Existing entries may retain their old headings until deliberately migrated; they are not automatically compliant with the new teaching requirements. Apply the substantive explanation, walkthrough, and edge-case requirements to every changed explanation, even when retaining legacy headings for a focused edit.

The current `dsa/internal/conceptsdoc/` checker enforces the legacy headings. When adopting the eight-section sequence in a guide, update the checker and its tests in the same change to recognize both layouts during migration and validate each complete entry's order. Retain link, fence, trace, source, and hierarchy checks. Update expected entry counts only for intentional additions, removals, or splits. Remove legacy-layout support after all active guides migrate. Do not disable checks to make a conversion pass.

Rewriting these standards alone does not authorize rewriting all guides or their validator. Keep migrations scoped, preserve useful content and anchors, and report which material was actually reviewed.

## 8. Review and definition of done

Before editing a guide, read its current working-tree entry, linked prerequisites, relevant curriculum, these standards, and `CODING_RULES.md` when code is involved. Inspect the diff, identify the primary owner, and trace references before moving or renaming material. Preserve user changes.

After editing, perform two separate reading passes:

1. Beginner pass: can the reader explain the problem, interpret each state value, and replay every decision without guessing? Repair the earliest missing explanation. Check that the example demonstrates the reason for choosing the concept.
2. Technical pass: replay traces against code, check boundaries and termination, verify guarantees, recalculate costs, and check alternatives and counterexamples. Exercise the main path and a relevant failure, cancellation, or exit path where applicable.

Review the diff for lost content, unintended changes, duplication, broken navigation, and stale sources. A section being present does not establish that it teaches its subject.

Use the commands and prerequisites in [VALIDATION.md](VALIDATION.md). For guide changes, run the concept-document, embedded-code, and line-width checkers, followed by repository tests. Execute changed complete examples as described above. For a standards-only edit, check Markdown formatting, whitespace, and consistency with the existing checker; do not claim unchanged guides meet new requirements.

An entry is ready only when:

- Its contract and prerequisites are explicit, and unfamiliar language is explained locally or through a clear prerequisite link.
- Its mental model defines the remembered state and rule for choosing the next action.
- Its ordinary walkthrough reaches the actual result through reproducible decisions and state changes matching the implementation.
- Its focused examples cover materially different branches, relevant boundaries, and a plausible incorrect approach.
- Its reasoning establishes correctness under stated assumptions, including termination or lifecycle behavior.
- Its cost analysis accounts for meaningful work and storage, and its selection guidance explains useful alternatives.
- Its implementation is readable and verified, navigation works, and applicable checks pass.

Report exact validation failures and unresolved limitations. Automated checks provide structural and executable evidence; the two reading passes establish whether the lesson actually teaches the concept.
