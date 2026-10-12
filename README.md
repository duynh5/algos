# Algorithms and Data Structures in Go

This repository is a learning-focused reference for algorithms, data structures,
and repeatable problem-solving patterns in Go. It keeps reusable teaching
implementations separate from judge-ready problem solutions.

## Repository map

- [`concepts/`](concepts/) contains the curriculum, explanations, reusable Go
  examples, pitfalls, and focused tests.
- [`workbench/`](workbench/) is the runnable active problem and preserves your
  current attempt while solutions are developed.
- [`problems/`](problems/) contains individually runnable Go files for completed
  solutions and unsolved practice archives.
- [`internal/testrunner/`](internal/testrunner/) contains the shared sample
  runner used by the workbench and archived problems.

Concept code is written for learning and reuse. Problem algorithms stay local
to their problem file so the selected judge function remains easy to copy.

## Starting a problem

All active-problem content lives in one editable file:
[`workbench/main.go`](workbench/main.go). Its labeled sections hold the problem
statement, sample cases, your preserved attempt, the three completed
agent-authored solutions, and their analysis.

1. Paste the statement into the `PROBLEM` block comment.
2. Add problem-specific types and input/output adapters when the judge uses
   structural values such as linked lists or trees.
3. Paste each example into the `sampleCases` slice using the included template.
4. Implement `solveCurrentAttempt`.
5. Ask Codex to review or solve "the problem in the workbench."

If you ask Codex to add the inputs or examples, it will also create or update
the problem-specific parser and call it at the start of your current
implementation. Your algorithm will receive ready-to-use, strongly typed
arguments rather than being left with raw sample text.

Run the active file at any point with:

```bash
./scripts/workbench.sh
```

The script prints `PASS` or `FAIL` for every solution/sample pair and then runs
the workbench tests.

The default collaboration flow is:

1. Restate the contract, constraints, and important edge cases.
2. Review your current attempt without rewriting or discarding it.
3. Identify and link the relevant pages under `concepts/`.
4. Implement direct, improved, and preferred solutions in the workbench.
5. Run one shared table of meaningful tests against your attempt and all three
   agent-authored solutions.
6. Document a correctness argument and comprehensive time and space analysis
   for each approach in the analysis section of `workbench/main.go`.
7. Archive the completed solution with `archive-workbench.sh`, then run the
   archived problem and the repository checks.

Because this is a public repository, problem notes should contain an original
summary and a link to the source rather than a verbatim copy of a third-party
problem statement.

## Archiving and running problems

Once the workbench is complete, assign a difficulty and archive its single
problem file:

```bash
./scripts/archive-workbench.sh easy 1 two-sum
```

This creates `problems/solved_easy_leetcode_0001_two_sum.go`. Numeric problem
IDs use at least four digits so filenames sort numerically within their existing
status, difficulty, and source groups. The `solved_` prefix
marks archives that preserve your attempt and the completed agent review.
LeetCode is the default
source; pass a fourth argument only when using another source. The source URL
may remain in the problem comment for attribution. Before archiving, Codex
checks the supplied problem ID against the official source page and confirms
that its title or statement matches the workbench. The archive command then runs
the workbench tests, including the preserved-attempt and three-agent-solution
registration checks.

Run one archived problem or all completed archives. Unsolved practice archives
use the same section layout as the workbench: they include official fixtures
and typed input adapters, a blank current-attempt function, an empty agent
solution registry, and blank analysis headings. After implementing the current
attempt, run that file directly to exercise it against its sample cases. The
bulk runner skips practice archives until a later review fills the solution and
analysis sections and marks the archive ready:

```bash
./scripts/run-problem.sh solved_easy_leetcode_0001_two_sum.go
./scripts/test-problems.sh
```

After confirming the archive, restore the blank workbench template:

```bash
./scripts/reset-workbench.sh
```

The reset refuses to discard unfinished work. `--force` is available only when
you explicitly intend to abandon the current workbench contents. Before
replacing `workbench/main.go`, the script formats, runs, and tests a temporary
candidate; a validation failure leaves the current workbench unchanged.

Archived files use `//go:build ignore`, so package-wide Go commands do not
combine independent programs. Named-file execution still runs them directly,
using the repository's shared `internal/testrunner` package.

## Validation

```bash
go test ./...
./scripts/test-problem-workflow.sh
```
