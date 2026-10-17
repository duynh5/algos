# Repository Agent Instructions

## Purpose

This repository teaches algorithms and data structures in Go. Keep educational
concept implementations under `concepts/`, the active exercise under
`workbench/`, and completed judge-ready solutions under `problems/`.

## Problem Retrieval

When the user asks to "retrieve" a problem from a URL, treat the authoritative
page at that URL as the source of truth and create an unsolved practice file
directly under `problems/`. Use `workbench/main.go.template` as the canonical
structural skeleton, then adapt it into a standalone practice file by adding
the `//go:build ignore` directive and file-level difficulty, source, and problem
ID metadata used under `problems/`. Do not put the retrieved problem in
`workbench/` and do not add the `solved_` prefix.

1. Open the authoritative problem page and verify the source, problem ID,
   title, difficulty, full statement, constraints, examples, and canonical URL.
   If required metadata is unavailable or ambiguous, stop and ask the user
   rather than inferring it.
2. Name the file
   `<difficulty>_<source>_<problem_id>_<descriptive_slug>.go`, using lowercase
   difficulty and source names and zero-padding numeric problem IDs to at least
   four digits.
3. Fill in the complete prompt, source metadata, problem-specific judge-owned
   types, strongly typed input parser, problem-specific output formatter,
   and all official examples as executable `sampleCases`.
4. Add useful valid coverage beyond the official examples, including minimum,
   boundary, duplicate, adversarial, transition, and constraint-extreme cases
   as applicable. Prefer cases that expose common incorrect assumptions while
   keeping the practice file practical to run locally.
5. Keep `archiveReady` set to `false`, register only
   `currentAttemptSolution`, and leave `agentSolutions` empty. In
   `solveCurrentAttempt`, parse the input into strongly typed arguments and
   leave only the user's algorithm body and result construction unimplemented.
6. Leave the `CURRENT ATTEMPT` and `AGENT SOLUTIONS` sections unsolved. Do not
   add an algorithm, judge solution, hints, completed review analysis, or
   concept-based solution material during retrieval. Preserve the blank review
   analysis template for a later "review" request.
7. Format the generated Go file and verify that its retrieval scaffolding is
   structurally consistent with `workbench/main.go.template` plus the required
   standalone practice-file metadata. Do not treat expected test failures from
   the deliberately blank current attempt as a retrieval failure.
8. When the user later asks to "review" a retrieved practice file, treat that
   file as the single source of truth and as an implementation request. Apply
   the Active Problem Workflow below to that file: preserve and assess the
   user's current attempt, add three distinct agent-authored solutions, run all
   implementations against the same cases, and complete the worked analysis
   and complexity discussion. Substitute the retrieved file for
   `workbench/main.go`; do not move it through the workbench.
9. Validate a reviewed practice file with
   `./scripts/run-problem.sh <filename>`, `./scripts/test-problems.sh`, and
   `go test ./...`. Set `archiveReady` to `true` only when the current attempt,
   all three agent solutions, regression coverage, and analysis are complete
   and the required checks pass.

## Active Problem Workflow

When the user asks for help with, an answer to, or a review of "the current
problem" without naming a file, inspect both `workbench/main.go` and unsolved
practice files directly under `problems/`. Treat a non-template workbench as a
candidate, and use the modified or untracked unsolved practice file as a
candidate when exactly one contains current-attempt work. If exactly one
candidate exists, use it as the single source of truth for the statement,
samples, current attempt, solutions, and analysis. If both locations or
multiple practice files are plausible, ask the user which file to review.

When the user explicitly asks about "the problem in the workbench," treat
`workbench/main.go` as that single source of truth. When the resolved current
problem is under `problems/`, apply this workflow to that file in place and do
not move it through the workbench.

A request to review the resolved current problem is an implementation request
unless the user explicitly asks for findings only or says not to edit files.
Complete the full workflow below, including the preserved current attempt,
three agent-authored solutions, and their analysis; do not stop after reporting
review findings.

When the user says "fill in," treat it as an input-setup request, not as a
request to solve the problem. Fill in the input from the problem statement in
the resolved problem file as executable code, then add inputs and examples
covering all valid edge cases, boundary conditions, adversarial cases, and
constraint extremes. Complete the parsing and formatting setup described below
so the user's current implementation receives strongly typed arguments, while
preserving the user's algorithm body.

1. Read the resolved problem file, `workbench/main_test.go` when the problem is
   in the workbench, and the relevant concept pages before changing code.
2. Keep shared judge-owned types and problem-specific parsing and formatting
   adapters in a compact support block before `sampleCases`. Preserve readable
   statement serialization in the cases, convert structural values inside
   strongly typed problem-named adapters, and keep `internal/testrunner`
   problem-agnostic.
   When the user asks the agent to add inputs or examples, complete the initial
   input setup rather than only populating `sampleCases`: create or update the
   problem-specific parser and call it at the start of the user's current
   implementation so the parsed, strongly typed arguments are ready for their
   algorithm. Keep setup used only by the current attempt at the top of that
   function. Add and wire a problem-specific output formatter at the same
   boundary for every problem. Ordinary scalar and collection results should
   use `testrunner.FormatJSONOutput`; structural and custom-judge results must
   convert to the statement's representation first. Preserve the user's
   algorithm body while adding this setup.
3. Preserve and review the user's attempt. Keep `solveCurrentAttempt` and its
   algorithm body intact except for the problem-specific parsing or formatting
   setup explicitly described above. Do not refactor it into a helper, rename
   its variables, or use it as a substitute for one of the three agent-authored
   solutions. Every review must score the current attempt out of 100 and briefly
   explain both the score and any deductions in the problem analysis and final
   response. Explain what is correct, the first failing assumption or bottleneck,
   and how the relevant invariant guides the correction. Explicitly compare the
   current attempt with each agent-authored solution. When they use the same core
   algorithm, data structure, or invariant, state that relationship in the
   analysis and describe only the material implementation differences; do not
   imply that a cleaned-up implementation is a distinct algorithm. The three
   agent-authored approaches must still be distinct from one another.
4. Link the problem to existing pages under `concepts/`. Add or improve a
   concept page when the necessary reference does not exist.
5. Produce three distinct correct approaches: a direct or brute force
   approach, an improved approach, and the preferred approach. An
   educational approach may exceed judge limits, but it must still be correct
   for inputs it can finish. Do not manufacture meaningless variants.
6. Give each agent-authored approach a descriptive function name and register
   it in `agentSolutions` in the resolved problem file. Register the preserved
   attempt separately as `currentAttemptSolution`, then combine it with the
   three agent solutions in `solutions`. Even when the preferred approach uses
   the same method as the current attempt, implement it in a separate
   descriptive function. Define helpers used exclusively by one agent-authored
   solution as local function values inside that solution function; do not
   scatter one approach across file-scope helper declarations. Keep only shared
   adapters, shared types or helpers, and judge-required platform entry points
   at file scope. Keep judge-required platform functions as thin wrappers around
   the preferred agent approach when appropriate. Reserve full-width section
   banners for `CURRENT ATTEMPT` and `AGENT SOLUTIONS`. Use compact one-line
   labels for adapters, registries, and other supporting functions.
7. Run the same `sampleCases` and table-driven regression cases against the
   current attempt and all three agent approaches.
   Include supplied examples, minimum or empty inputs when valid, boundary
   conditions, duplicate or adversarial values, and at least one case that
   distinguishes a common incorrect solution.
8. In the resolved problem file's analysis section, explain each approach, its
   correctness argument, and its time and auxiliary-space complexity. State
   best, average, and worst cases when they differ, and account for recursion,
   preprocessing, output, and meaningful Go allocations. For each solution,
   include a concrete worked example using an actual input. Trace the relevant
   variables, pointers, data structures, and state changes step by step so the
   reader can map the example directly to the code; an outcome-only example is
   not sufficient.
9. For a workbench problem, run `./scripts/workbench.sh` and `go test ./...`.
   For a problem under `problems/`, run `./scripts/run-problem.sh <filename>`,
   `./scripts/test-problems.sh`, and `go test ./...`. Set `archiveReady` to
   `true` only after the location-appropriate checks pass and the current
   attempt, all three agent solutions, and their analysis are complete.

Do not replace a non-template workbench problem with a new one until its work
has been archived or the user explicitly says to discard it.

## Archiving

After a problem is complete:

When the user asks to archive a completed retrieved practice file that is
already under `problems/`, verify its source metadata and `archiveReady` state,
then rename it in place by adding the `solved_` prefix to its existing basename.
Do not run `archive-workbench.sh` or reset the workbench for this path. After the
rename, run `./scripts/run-problem.sh <solved-filename>`,
`./scripts/test-problems.sh`, and `go test ./...` before reporting success.

1. Assign `easy`, `medium`, or `hard` based on the required insight,
   implementation complexity, edge cases, and constraints. Record the source
   and source problem ID; default the source to `leetcode` unless the user says
   otherwise. Before using any user-provided or inferred problem ID, open the
   authoritative problem page for that source and verify that both the ID and
   problem title or statement match the workbench. Record the canonical source
   URL when available. Do not rely on memory or the supplied ID alone; if the
   source disagrees or remains ambiguous, stop and ask the user before naming
   or archiving the file.
2. Run
   `./scripts/archive-workbench.sh <difficulty> <problem-id> <descriptive-slug> [source]`.
3. Confirm the new archived file under `problems/` runs with
   `./scripts/run-problem.sh <filename>`.
4. Run `./scripts/test-problems.sh` and `go test ./...`.
5. Reset the active file with `./scripts/reset-workbench.sh` only after the
   archive and checks succeed.

Use flat filenames such as `solved_easy_leetcode_0001_two_sum.go`; do not group by
platform or create per-problem directories. Completed filenames must follow
`solved_<difficulty>_<source>_<problem_id>_<descriptive_slug>.go`; blank
practice archives omit the `solved_` prefix. Zero-pad numeric problem IDs to at
least four digits in filenames while keeping canonical IDs unpadded inside the
file. Each completed archive contains all problem-specific material: its
statement, samples, preserved current attempt, three agent-authored solutions,
and analysis. It imports only the shared `internal/testrunner` infrastructure;
reference concept pages in comments rather than importing their code.

Never use `reset-workbench.sh --force` unless the user explicitly authorizes
discarding the unfinished workbench content.

`workbench/main.go.template` is the canonical blank workbench. When changing
its problem-facing API, keep it synchronized with `workbench/main.go` while
preserving any active problem, samples, solutions, and analysis. Universal
runner behavior and tests belong in `internal/testrunner`.
