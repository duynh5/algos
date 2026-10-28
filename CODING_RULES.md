# AI Code Generation Rules

These rules are mandatory whenever an AI agent creates or modifies
content expressed as code syntax in this repository. This includes
source code, tests, scripts, configuration, queries, command examples,
generated scaffolding, and code examples embedded in documentation. It
does not include ordinary prose, raw data, lockfiles, or verbatim
third-party content. Workflow-specific instructions in `AGENTS.md` add
to these rules and take precedence when they are more specific.

## Physical Line Width

- Keep every physical line in code files at no more than 70 Unicode
  characters.
- Apply the limit to source code, tests, scripts, templates,
  configuration, queries, generated scaffolding, comments, and code
  examples embedded in documentation.
- Ordinary prose outside the active concept guides has no physical
  line-width limit.
- Represent long URLs, paths, identifiers, and serialized values
  without creating a physical code line over 70 characters.
- Run `go run ./internal/linewidth` after modifying code.

## Preserve Intent and Existing Behavior

- Make the smallest correct change that satisfies the request.
- Preserve existing APIs, commands, data shapes, shortcuts, output, and user experience unless the task explicitly changes them.
- Preserve user-authored work. Follow any workflow-specific restrictions on which functions or sections an agent may edit.
- Before changing shared behavior, renaming a concept, or removing code, trace the affected callers, types, tests, documentation, exports, mocks, fixtures, string references, and dynamic loading paths.
- Follow the repository's established patterns before adding dependencies, abstractions, files, or directories.

## Structure and Maintainability

- Prefer clear, direct code over clever or compressed implementations.
- Write code so a reader can follow each non-obvious operation from
  local names, control flow, and nearby comments without reconstructing
  hidden intermediate steps.
- Make non-obvious state transitions explicit. Prefer named intermediate
  values or separate statements when a compact expression hides what is
  read, changed, preserved, or returned.
- Keep a concise standard idiom when it is clearer than expanded code.
  When index arithmetic, a sentinel, a bit operation, a fallback table,
  mutation, or a compound condition depends on an invariant, explain the
  invariant and why the operation is safe immediately before it.
- Do not add temporary values or comments to self-evident operations.
  Explicitness means exposing necessary reasoning, not narrating syntax.
- Prefer explicit, role-based identifiers even when they are longer. Name a count, index, boundary, or state by what it measures and where it applies; for example, use `numNodesInCurrentLayer` instead of `levelSize` and `currentIndex` instead of `i`.
- Avoid abbreviations such as `cur`, `prev`, and `idx`, vague names whose role is not evident, and single-letter identifiers. Use conventional mathematical symbols only when the surrounding explanation defines the same notation, or when an external interface requires the name.
- Name mappings and indexed collections by both the stored fact and its lookup dimension when that relationship is not otherwise obvious; for example, use `distanceByVertex` or `earlierIndexByValue` instead of `distance` or `seen`.
- When renaming an identifier, update its comments, worked traces, documentation, tests, and other explanatory references in the same change.
- Do not shorten a precise identifier merely to fit a long statement. Split declarations, conditions, or control flow into readable steps instead.
- Put new code in the smallest project-appropriate location and connect it through existing entry points.
- Extract a helper, type, or module only when it removes meaningful duplication or isolates real complexity with stable inputs and low coupling.
- Keep functions focused and control flow easy to follow. Avoid hidden flow, unnecessary indirection, and recursion when a simpler bounded approach is appropriate.
- Give loops, retries, queues, caches, concurrency, and other resource use clear bounds derived from the problem or validated inputs.
- Keep state as local as practical. Minimize shared mutation, aliasing, callback indirection, reflection, dynamic loading, and conditional build behavior.
- When an authored code file exceeds 500 lines, treat splitting its responsibilities into focused project files as part of the change. Exclude generated files, snapshots, lockfiles, and data or configuration blobs.
- In a non-root folder with more than 10 files, organize new work into an appropriate subfolder or split it by responsibility before adding another file.

## Correctness and Boundaries

- Encode important invariants with validation, types, side-effect-free assertions, or tests.
- Use each variable for one clear purpose and the narrowest practical scope.
- Validate boundary inputs, check meaningful return values and errors, and fail fast when internal state violates an invariant.
- At system boundaries, catch errors narrowly and either recover meaningfully or return them with enough context to diagnose the failure.
- Surface defects instead of hiding them behind broad fallbacks, swallowed errors, or unrelated default values.
- Treat compiler, formatter, linter, static-analysis, and test warnings as failures unless the repository documents an exception.

## Code Comments

Comment agent-generated code so a learner can follow every meaningful logical block without relying on separate prose. Comments must remain concise enough that the code itself stays readable.

- Place comments immediately before the code they explain. Cover meaningful phases such as input or boundary handling, state initialization, the main loop or recursion, important branches and mutations, state restoration or backtracking, and result construction.
- Explain why a block exists, what invariant or state transition it maintains, and how that advances the algorithm. Do not merely translate individual statements into English or repeat self-explanatory syntax.
- Explain non-obvious language or algorithm details at the point of use. Common examples include bit operations, sentinel values, type assertions, half-open ranges, overflow guards, queue or stack ordering, visited-mark timing, and pruning conditions.
- State externally visible side effects and important assumptions, including sorting or otherwise mutating an input slice, sharing backing storage, or relying on constraints that make an index or conversion safe.
- Give nested loops, recursive cases, and substantial branch bodies enough context that their role and state changes are clear.
- Avoid comments that duplicate full correctness proofs, worked traces, or complexity analysis. Keep those explanations in the appropriate surrounding documentation.
- Re-read every comment against the final code. Correct stale, misleading, or ambiguous explanations before considering the work complete.

## Tests and Validation

- Use the repository's configured formatter, tests, static analysis, and validation scripts; do not introduce a parallel toolchain for one change.
- Develop new behavior test-first when the affected area has tests. For a bug fix, add a focused regression test that fails before the fix when practical.
- Pin risky untested behavior with a characterization test before changing it when practical.
- Add or update only the tests related to the requested change.
- Cover the main successful path, relevant boundaries and adversarial cases, and at least one meaningful failure or exit path when applicable.
- Run the narrowest useful validation first, then broaden it when the change crosses package, workflow, or repository boundaries.
- Treat clarity as a review requirement, not a formatter result. During
  review, trace each non-obvious expression and state change from its
  local names and comments; formatting, compilation, and tests alone do
  not establish compliance.
- Format Go code with `gofmt`. Run the workflow-specific commands required by `AGENTS.md`, followed by `go test ./...` when the change affects buildable Go packages or shared behavior.
- Review the final diff for accidental changes and verify that implementation, tests, comments, and documentation agree.
- Turn repeated or correctness-sensitive manual validation into a deterministic script, fixture, test, or documented command when practical.
