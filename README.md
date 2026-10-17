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

## Agent workflow requests

These phrases invoke repository-specific Codex workflows. They are not shell
commands.

| Request | Result |
| --- | --- |
| `retrieve <problem-url>` | Verifies the authoritative page, then creates an unsolved practice file under `problems/` from `workbench/main.go.template`. It includes metadata, the prompt, typed input and output adapters, official examples, and additional edge cases, while leaving the current attempt and agent solutions unsolved. |
| `fill in` | Completes input cases, parsing, and output formatting for the active workbench problem without solving or replacing the current algorithm. |
| `review the problem in the workbench` | Preserves the current attempt, adds direct, improved, and preferred solutions, runs shared regression cases, and completes the correctness and complexity analysis. |
| `review problems/<practice-file>.go` | Applies the same three-solution review workflow directly to a retrieved practice file without moving it through the workbench. |
| `archive the problem` | Archives a completed workbench problem with a `solved_` filename. For a completed retrieved practice file, adds the `solved_` prefix in place. Archiving requires `archiveReady = true` and successful validation. |

## Command reference

Run these commands from the repository root. The scripts require Bash, Go,
Git, `gofmt`, and `rg` (ripgrep). The module currently declares Go 1.26.3.

### Daily problem workflow

| Command | Purpose and behavior |
| --- | --- |
| `./scripts/workbench.sh` | Runs the active workbench program and then its package tests. Equivalent to `go run ./workbench` followed by `go test ./workbench`. |
| `./scripts/run-problem.sh [--case <number>] [<filename>]` | Runs one file under `problems/`. By default it runs every configured case; `--case` runs one case using a one-based number. With no filename, it runs the sole modified or untracked problem file and fails if that choice is missing or ambiguous. An explicit basename or `problems/<filename>` must use the `[solved_]<difficulty>_<source>_<problem-id>_<slug>.go` naming format. |
| `./scripts/test-problems.sh` | Runs every completed problem file. Files containing `archiveReady = false` are reported and skipped. |
| `./scripts/archive-workbench.sh <easy\|medium\|hard> <problem-id> <descriptive-slug> [source]` | Validates the completed workbench, creates a standalone `solved_` file under `problems/`, adds its metadata, runs it, and appends it to `problems/README.md`. The source defaults to `leetcode`; the command refuses an incomplete workbench or an existing destination. |
| `./scripts/reset-workbench.sh` | Restores `workbench/main.go` and `main_test.go` from their canonical templates after validating the replacement. It refuses to discard an unfinished workbench. |
| `./scripts/reset-workbench.sh --force` | Performs the same reset while explicitly allowing unfinished workbench content to be discarded. This is destructive and should only be used intentionally. |

Examples:

```bash
./scripts/workbench.sh
./scripts/run-problem.sh
./scripts/run-problem.sh --case 1 easy_leetcode_0066_plus_one.go
./scripts/run-problem.sh easy_leetcode_0066_plus_one.go
./scripts/archive-workbench.sh easy 1 two-sum
./scripts/reset-workbench.sh
```

### Repository validation

| Command | Purpose and behavior |
| --- | --- |
| `go test ./...` | Runs every Go package test that is not excluded by build constraints. Standalone problem files use `//go:build ignore`, so they are exercised by the problem scripts instead. |
| `./scripts/test-problem-workflow.sh` | Runs the archive, reset, filename-validation, single-problem, and bulk-problem workflow tests inside a temporary repository fixture. It does not modify the active workbench. |

The usual full validation sequence is:

```bash
./scripts/workbench.sh
./scripts/test-problems.sh
./scripts/test-problem-workflow.sh
go test ./...
```

### Maintainer bulk import

The Top Interview 150 importer is a repository maintenance command, not part of
the normal single-problem workflow:

```bash
go run scripts/import-top-interview-150-practice.go \
  scripts/import-top-interview-150-practice-helpers.go \
  [-manifest <catalog.tsv>] \
  [-metadata <metadata-directory>] \
  [-problems <destination-directory>]
```

Defaults are `scripts/top-interview-150.tsv`,
`/tmp/leetcode-top-interview-150`, and `problems`. The metadata directory must
already contain the separately downloaded LeetCode GraphQL JSON responses. The
importer requires a 150-entry catalog, verifies each response against it,
generates unsolved practice files, and updates `problems/README.md`. It refuses
to overwrite a non-generated practice file.

## Starting a problem

All active-problem content lives in one editable file:
[`workbench/main.go`](workbench/main.go). Its labeled sections hold the problem
statement, sample cases, your preserved attempt, the three completed
agent-authored solutions, and their analysis.

1. Paste the statement into the `PROBLEM` block comment.
2. Add problem-specific types, a typed input parser, and an output formatter.
   Ordinary results can delegate to `testrunner.FormatJSONOutput`; structural
   judges such as linked lists or trees need their statement representation.
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

Retrieved problem files retain source metadata and the canonical URL for
attribution. Check the source's licensing terms before redistributing imported
problem statements.

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
and typed input/output adapters, a blank current-attempt function, an empty agent
solution registry, and blank analysis headings. After implementing the current
attempt, run that file directly to exercise it against its sample cases. The
bulk runner skips practice archives until a later review fills the solution and
analysis sections and marks the archive ready:

```bash
./scripts/run-problem.sh
./scripts/run-problem.sh --case 1 solved_easy_leetcode_0001_two_sum.go
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

See the complete [command reference](#command-reference), including the
recommended full validation sequence and the isolated workflow regression
test.
