# Repository Agent Instructions

## Purpose

This repository teaches algorithms and data structures in Go. Keep educational
concept implementations under `concepts/`, the active exercise under
`workbench/`, and completed judge-ready solutions under `problems/`.

## Active Problem Workflow

When the user asks for help with, an answer to, or a review of "the problem in
the workbench," treat `workbench/main.go` as the single source of truth for the
statement, samples, current attempt, solutions, and analysis.

A request to review the workbench problem is an implementation request unless
the user explicitly asks for findings only or says not to edit files. Complete
the full workflow below, including the preserved current attempt, three
agent-authored solutions, and their analysis; do not stop after reporting
review findings.

When the user says "fill in," treat it as an input-setup request, not as a
request to solve the problem. Fill in the input from the problem statement in
`workbench/main.go` as executable code, then add inputs and examples covering
all valid edge cases, boundary conditions, adversarial cases, and constraint
extremes. Complete the parsing and formatting setup described below so the
user's current implementation receives strongly typed arguments, while
preserving the user's algorithm body.

1. Read `workbench/main.go`, `workbench/main_test.go`, and the relevant concept
   pages before changing code.
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
   function. Add and wire an output formatter at the same boundary when the
   judge result needs structural serialization. Preserve the user's algorithm
   body while adding this setup.
3. Preserve and review the user's attempt. Keep `solveCurrentAttempt` and its
   algorithm body intact except for the problem-specific parsing or formatting
   setup explicitly described above. Do not refactor it into a helper, rename
   its variables, or use it as a substitute for one of the three agent-authored
   solutions. Explain what is correct, the first failing assumption or
   bottleneck, and how the relevant invariant guides the correction. Explicitly
   compare the current attempt with each agent-authored solution. When they use
   the same core algorithm, data structure, or invariant, state that relationship
   in the analysis and describe only the material implementation differences;
   do not imply that a cleaned-up implementation is a distinct algorithm. The
   three agent-authored approaches must still be distinct from one another.
4. Link the problem to existing pages under `concepts/`. Add or improve a
   concept page when the necessary reference does not exist.
5. Produce three distinct correct approaches: a direct or brute force
   approach, an improved approach, and the preferred approach. An
   educational approach may exceed judge limits, but it must still be correct
   for inputs it can finish. Do not manufacture meaningless variants.
6. Give each agent-authored approach a descriptive function name and register
   it in `agentSolutions` in `workbench/main.go`. Register the preserved attempt
   separately as `currentAttemptSolution`, then combine it with the three agent
   solutions in `solutions`. Even when the preferred approach uses the same
   method as the current attempt, implement it in a separate descriptive
   function. Keep judge-required platform functions as thin wrappers around the
   preferred agent approach when appropriate. Reserve full-width section
   banners for `CURRENT ATTEMPT` and `AGENT SOLUTIONS`. Use compact one-line
   labels for adapters, registries, and other supporting functions.
7. Run the same `sampleCases` and table-driven regression cases against the
   current attempt and all three agent approaches.
   Include supplied examples, minimum or empty inputs when valid, boundary
   conditions, duplicate or adversarial values, and at least one case that
   distinguishes a common incorrect solution.
8. In the analysis section of `workbench/main.go`, explain each approach, its
   correctness argument, and its time and auxiliary-space complexity. State
   best, average, and worst cases when they differ, and account for recursion,
   preprocessing, output, and meaningful Go allocations. For each solution,
   include a concrete worked example using an actual input. Trace the relevant
   variables, pointers, data structures, and state changes step by step so the
   reader can map the example directly to the code; an outcome-only example is
   not sufficient.
9. Run `./scripts/workbench.sh` and `go test ./...` before claiming that the
   solutions pass. Set `archiveReady` to `true` only after these checks pass and
   the current attempt, all three agent solutions, and their analysis are
   complete.

Do not replace a non-template workbench problem with a new one until its work
has been archived or the user explicitly says to discard it.

## Archiving

After a problem is complete:

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

Use flat filenames such as `solved_easy_leetcode_1_two_sum.go`; do not group by
platform or create per-problem directories. Completed filenames must follow
`solved_<difficulty>_<source>_<problem_id>_<descriptive_slug>.go`; blank
practice archives omit the `solved_` prefix. Each completed archive contains
all problem-specific material: its statement, samples, preserved current
attempt, three agent-authored solutions, and analysis. It imports only the
shared `internal/testrunner` infrastructure; reference concept pages in
comments rather than importing their code.

Never use `reset-workbench.sh --force` unless the user explicitly authorizes
discarding the unfinished workbench content.

`workbench/main.go.template` is the canonical blank workbench. When changing
its problem-facing API, keep it synchronized with `workbench/main.go` while
preserving any active problem, samples, solutions, and analysis. Universal
runner behavior and tests belong in `internal/testrunner`.
