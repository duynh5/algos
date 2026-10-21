# Repository Agent Instructions

## Purpose

This repository teaches algorithms and data structures in Go. Keep the
educational guide in `CONCEPTS.md`, the active exercise under `workbench/`, and
standalone practice and completed judge-ready solutions under `problems/`.

## Code Generation Standard

Before creating or modifying code, read and follow
[`CODING_RULES.md`](CODING_RULES.md). It applies to all agent-generated source
code, tests, scripts, scaffolding, and code examples in documentation. The
workflows below add task-specific rules.

## Non-Negotiable Safety Rules

- Preserve user work. Do not replace a non-template workbench problem until it
  has been archived or the user explicitly says to discard it.
- Never use `reset-workbench.sh --force` unless the user explicitly authorizes
  discarding the unfinished workbench content.
- Treat the resolved problem file as the single source of truth for its
  statement, samples, current attempt, solutions, and analysis.
- `workbench/main.go.template` is the canonical blank workbench. When changing
  its problem-facing API, keep it synchronized with `workbench/main.go` while
  preserving any active problem, samples, solutions, and analysis.
- Keep universal runner behavior and tests in `internal/testrunner`.

## Request Routing

The quoted requests below are natural-language requests to the agent, not shell
commands.

| Request | Required route |
| --- | --- |
| `retrieve <problem-url>` | Create a new unsolved practice file directly under `problems/`; do not use the workbench. |
| `fill in`, `add inputs`, or `add examples` | Resolve the current problem, then complete only its cases, parser, and formatter setup. |
| `review the problem in the workbench` | Use `workbench/main.go` and perform the full review workflow. |
| `review problems/<practice-file>.go` | Review that file in place; do not move it through the workbench. |
| Review of "the current problem" without a path | Resolve the current problem first, then perform the full review workflow. |
| Help with or an answer to "the current problem" | Resolve the current problem first, then follow the user's requested scope. |
| `archive the problem` | Use the archive path for the resolved file's current location. |

## Resolve the Current Problem

Use this resolution process for an unnamed current-problem request:

1. Inspect both `workbench/main.go` and unsolved practice files directly under
   `problems/`.
2. Treat a non-template workbench as a candidate. Treat a modified or untracked
   unsolved practice file as a candidate when exactly one contains
   current-attempt work.
3. If exactly one candidate exists, use it. If both locations or multiple
   practice files are plausible, ask the user which file to use.

When the user explicitly names the problem in the workbench, use
`workbench/main.go`. When the user explicitly names a file under `problems/`,
use that file in place.

Before changing a resolved problem, read the file, `workbench/main_test.go`
when the problem is in the workbench, and the relevant concept pages.

## Shared Problem-File Boundaries

- Keep shared judge-owned types and problem-specific parsing and formatting
  adapters in a compact support block before `sampleCases`.
- Preserve readable statement serialization in the cases. Convert structural
  values inside strongly typed, problem-named adapters.
- Add and wire a problem-specific output formatter for every problem. Ordinary
  scalar and collection results should use `testrunner.FormatJSONOutput`.
  Structural and custom-judge results must convert to the statement's
  representation first.
- Keep `internal/testrunner` problem-agnostic.
- Keep only shared adapters, shared types or helpers, and judge-required
  platform entry points at file scope. Keep judge-required platform functions
  as thin wrappers around the preferred agent approach when appropriate.
- Reserve full-width section banners for `CURRENT ATTEMPT` and
  `AGENT SOLUTIONS`. Use compact one-line labels for adapters, registries, and
  other supporting functions.

## Retrieve Workflow

When the user asks to retrieve a problem from a URL, treat the authoritative
page at that URL as the source of truth. Create an unsolved practice file
directly under `problems/` from `workbench/main.go.template`, adding the
`//go:build ignore` directive and the file-level difficulty, source, and
problem ID metadata used under `problems/`. Do not add the `solved_` prefix.

1. Open the authoritative page and verify the source, problem ID, title,
   difficulty, full statement, constraints, examples, and canonical URL. If
   required metadata is unavailable or ambiguous, stop and ask the user rather
   than inferring it.
2. Name the file
   `<difficulty>_<source>_<problem_id>_<descriptive_slug>.go`, using lowercase
   difficulty and source names and zero-padding numeric problem IDs to at least
   four digits.
3. Fill in the complete prompt, source metadata, problem-specific judge-owned
   types, strongly typed input parser, problem-specific output formatter, and
   all official examples as executable `sampleCases`.
4. Add useful valid coverage beyond the official examples, including minimum,
   boundary, duplicate, adversarial, transition, and constraint-extreme cases
   as applicable. Prefer cases that expose common incorrect assumptions while
   keeping the practice file practical to run locally.
5. Keep `archiveReady` set to `false`, register only
   `currentAttemptSolution`, and leave `agentSolutions` empty. In
   `solveCurrentAttempt`, parse the input into strongly typed arguments and
   leave only the user's algorithm body and result construction unimplemented.
6. Leave the `CURRENT ATTEMPT` and `AGENT SOLUTIONS` sections unsolved. Do
   not add an algorithm, judge solution, hints, completed review analysis, or
   concept-based solution material. Preserve the blank review analysis template
   for a later review request.
7. Format the generated file and verify that its scaffolding is structurally
   consistent with `workbench/main.go.template` plus the required standalone
   metadata. Do not treat expected test failures from the deliberately blank
   current attempt as a retrieval failure.

## Fill-In Workflow

Treat "fill in," "add inputs," and "add examples" as input-setup requests, not
requests to solve or replace the current algorithm.

1. Fill in the input from the problem statement as executable code.
2. Add official examples, then add inputs and examples covering all valid edge
   cases, boundary conditions, adversarial cases, and constraint extremes.
3. Create or update the problem-specific parser and call it at the start of
   `solveCurrentAttempt` so the user's algorithm receives strongly typed
   arguments. Keep setup used only by the current attempt at the top of that
   function.
4. Create or update the output formatter according to the shared problem-file
   boundaries.
5. Preserve the user's algorithm body. Do not solve the problem.

## Review Workflow

A review of the resolved current problem is an implementation request unless
the user explicitly asks for findings only or says not to edit files. Complete
the full workflow; do not stop after reporting findings. A retrieved practice
file remains in place under `problems/`.

1. Preserve and assess `solveCurrentAttempt`. Keep its algorithm body intact
   except for the problem-specific parsing or formatting setup allowed above.
   Do not refactor it into a helper, rename its variables, or use it as a
   substitute for an agent-authored solution.
2. Score the current attempt out of 100 and briefly explain the score and any
   deductions in both the problem analysis and final response. Explain what is
   correct, the first failing assumption or bottleneck, and how the relevant
   invariant guides the correction.
3. Link the problem to existing material in `CONCEPTS.md`. Add or improve a
   concept entry when the necessary reference does not exist.
4. Implement three distinct correct approaches: a direct or brute-force
   approach, an improved approach, and the preferred approach. An educational
   approach may exceed judge limits, but it must still be correct for inputs it
   can finish. Do not manufacture meaningless variants.
5. Give each agent-authored approach a descriptive function name and register
   it in `agentSolutions`. Register the preserved attempt separately as
   `currentAttemptSolution`, then combine it with the three agent solutions in
   `solutions`. Even when the preferred approach uses the same method as the
   current attempt, implement it in a separate descriptive function.
6. Define helpers used exclusively by one agent-authored solution as local
   function values inside that solution. Do not scatter one approach across
   file-scope helper declarations.
7. Explicitly compare the current attempt with each agent-authored solution.
   When they use the same core algorithm, data structure, or invariant, state
   that relationship in the analysis and describe only material implementation
   differences. Do not imply that a cleaned-up implementation is a distinct
   algorithm. The three agent-authored approaches must remain distinct from one
   another.
8. Run the same `sampleCases` and table-driven regression cases against the
   current attempt and all three agent approaches. Include supplied examples,
   minimum or empty inputs when valid, boundary conditions, duplicate or
   adversarial values, and at least one case that distinguishes a common
   incorrect solution.
9. In the problem file's analysis section, explain each approach, its
   correctness argument, and its time and auxiliary-space complexity. State
   best, average, and worst cases when they differ, accounting for recursion,
   preprocessing, output, and meaningful Go allocations.
10. Give every solution a concrete worked example using an actual input. Trace
    the relevant variables, pointers, data structures, and state changes step by
    step so the reader can map the example directly to the code; an outcome-only
    example is insufficient.
11. Run the location-appropriate review checks in the validation matrix. Set
    `archiveReady` to `true` only after the current attempt, all three agent
    solutions, regression coverage, analysis, and required checks are complete.

## Archive Workflow

Archive only a completed problem after the user requests it.

### Archive Format

- Use flat filenames; do not group problems by platform or create per-problem
  directories. For example: `solved_easy_leetcode_0001_two_sum.go`.
- Completed filenames must follow
  `solved_<difficulty>_<source>_<problem_id>_<descriptive_slug>.go`. Blank
  practice files omit the `solved_` prefix.
- Zero-pad numeric problem IDs to at least four digits in filenames while
  keeping canonical IDs unpadded inside the file.
- Each completed archive must contain its statement, samples, preserved current
  attempt, three agent-authored solutions, and analysis.
- Import only the shared `internal/testrunner` infrastructure. Reference
  concept pages in comments rather than importing their code.

### Retrieved Practice File

When the completed file is already under `problems/`:

1. Verify its source metadata and `archiveReady` state.
2. Rename it in place by adding `solved_` to its existing basename.
3. Do not run `archive-workbench.sh` and do not reset the workbench.
4. Run the location-appropriate archive checks in the validation matrix.

### Workbench Problem

When the completed problem is in the workbench:

1. Assign `easy`, `medium`, or `hard` based on the required insight,
   implementation complexity, edge cases, and constraints.
2. Record the source and source problem ID; default the source to `leetcode`
   unless the user says otherwise.
3. Before using any user-provided or inferred problem ID, open the authoritative
   source page and verify that both the ID and problem title or statement match
   the workbench. Record the canonical source URL when available. Do not rely on
   memory or the supplied ID alone. If the source disagrees or remains
   ambiguous, stop and ask the user before naming or archiving the file.
4. Run
   `./internal/archive-workbench.sh <difficulty> <problem-id> <descriptive-slug> [source]`.
5. Run the location-appropriate archive checks in the validation matrix.
6. Reset the active file with `./internal/reset-workbench.sh` only after the
   archive and all checks succeed.

## Validation Matrix

| State and location | Required checks |
| --- | --- |
| Retrieved unsolved practice file | Format it and verify its scaffolding against `workbench/main.go.template`; expected failures from the blank attempt are not retrieval failures. |
| Reviewed workbench problem | `./internal/workbench.sh`, then `go test ./...`. |
| Reviewed or archived file under `problems/`, including an archive created from the workbench | `./internal/run-problem.sh <filename>`, `./internal/test-problems.sh`, then `go test ./...`. |

## Documentation Map

- [`CONCEPTS.md`](CONCEPTS.md) is the learning guide for reusable
  algorithms, data structures, Go techniques, and problem-solving methods.
- [`CONCEPTS_STANDARDS.md`](CONCEPTS_STANDARDS.md) defines how to maintain that
  guide.
- [`workbench/README.md`](workbench/README.md) explains how to prepare and run
  the active problem.
- [`problems/README.md`](problems/README.md) documents standalone problem files
  and provides their index.
- [`VALIDATION.md`](VALIDATION.md) owns repository prerequisites and validation
  commands.
- [`INTERVIEW_PREP_PLAN.md`](INTERVIEW_PREP_PLAN.md) contains the current study
  schedule and readiness criteria.
