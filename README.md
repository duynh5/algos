# Algorithms and Data Structures in Go

This repository is a learning reference for algorithms, data structures, and
repeatable problem-solving methods in Go. Reusable teaching code is kept
separate from complete problem solutions that can be copied to an online judge.

## Repository map

- [`concepts/`](concepts/) contains explanations, reusable Go examples, common
  mistakes, and focused tests.
- [`workbench/`](workbench/) contains the one active problem. It keeps your
  attempt intact while additional solutions are developed.
- [`problems/`](problems/) contains standalone Go files for completed solutions
  and unfinished practice problems.
- [`internal/testrunner/`](internal/testrunner/) contains the shared code that
  runs each solution against its test cases.

Concept code is written for learning and reuse. Each problem's algorithms stay
inside that problem file so the chosen function remains easy to copy into an
online judge.

## Agent workflow requests

These phrases invoke repository-specific Codex workflows. They are not shell
commands.

- **`retrieve <problem-url>`:** Verify the original problem page and create an
  unfinished practice file under `problems/` from `workbench/main.go.template`.
  The file includes source details, the statement, official examples, additional
  edge cases, and helpers that convert between sample text and Go values. The
  current attempt and agent solutions remain blank.
- **`fill in`:** Complete the active problem's test cases, input parser, and
  output formatter without solving or replacing the current algorithm.
- **`review the problem in the workbench`:** Preserve the current attempt, add
  direct, improved, and preferred solutions, test all four implementations with
  the same cases, and explain why each works and how much work and memory it
  uses.
- **`review problems/<practice-file>.go`:** Perform the same review directly in
  a retrieved practice file without moving it into the workbench.
- **`archive the problem`:** Save a completed workbench problem with a `solved_`
  filename. For a completed practice file already under `problems/`, add the
  `solved_` prefix there. Archiving requires `archiveReady = true` and passing
  checks.

## Command reference

Run these commands from the repository root. The scripts require Bash, Go,
Git, `gofmt`, and `rg` (ripgrep). The module currently declares Go 1.26.3.

### Daily problem workflow

- **`./scripts/workbench.sh`:** Run the active workbench program, then run its
  tests. This is equivalent to `go run ./workbench` followed by
  `go test ./workbench`.
- **`./scripts/run-problem.sh [--case <number>] [<filename>]`:** Run one Go file
  under `problems/`. With no `--case`, it runs every test case configured in the
  file. `--case 1` runs only the first case. With no filename, the script uses
  the only modified or new problem file; it stops with an error when there are
  zero or multiple choices. An explicit filename must follow
  `[solved_]<difficulty>_<source>_<problem-id>_<slug>.go`.
- **`./scripts/test-problems.sh`:** Run every completed problem file. Report and
  skip files containing `archiveReady = false`.
- **Archive a completed workbench:**
  `./scripts/archive-workbench.sh <easy|medium|hard> <problem-id> <slug> [source]`
  checks the workbench, creates its standalone `solved_` file under `problems/`,
  adds source details, runs the new file, and appends it to
  `problems/README.md`. The source defaults to `leetcode`. The script refuses an
  incomplete workbench or an existing destination file.
- **`./scripts/reset-workbench.sh`:** Replace `workbench/main.go` and
  `main_test.go` with their standard blank templates after checking the
  replacement. The command refuses to discard an unfinished workbench.
- **`./scripts/reset-workbench.sh --force`:** Perform the same reset while
  allowing unfinished work to be discarded. This is destructive and should be
  used only when that loss is intentional.

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

- **`go test ./...`:** Run every Go package test included in the normal build.
  Standalone problem files begin with `//go:build ignore`, so this command skips
  them; the problem scripts run them separately.
- **`./scripts/test-problem-workflow.sh`:** Test archiving, resetting, filename
  checks, running one problem, and running all completed problems inside a
  temporary copy of the repository. It does not modify the active workbench.

The usual full validation sequence is:

```bash
./scripts/workbench.sh
./scripts/test-problems.sh
./scripts/test-problem-workflow.sh
go test ./...
```

### Maintainer bulk import

The Top Interview 150 import command creates many practice files at once. It is
for repository maintenance, not the normal one-problem workflow:

```bash
go run scripts/import-top-interview-150-practice.go \
  scripts/import-top-interview-150-practice-helpers.go \
  [-manifest <catalog.tsv>] \
  [-metadata <metadata-directory>] \
  [-problems <destination-directory>]
```

The three default locations are `scripts/top-interview-150.tsv` for the catalog,
`/tmp/leetcode-top-interview-150` for downloaded problem information, and
`problems` for generated files. Before running the command, separately download
the JSON responses returned by LeetCode's GraphQL web API into the information
directory. The importer requires exactly 150 catalog entries, checks each
downloaded response against the catalog, creates unfinished practice files, and
updates `problems/README.md`. It refuses to overwrite a practice file it did not
generate.

## Starting a problem

All content for the active problem lives in one editable file:
[`workbench/main.go`](workbench/main.go). Its labeled sections hold the problem
statement, test cases, your preserved attempt, three agent-authored solutions,
and explanations of why they work and how much work and memory they use.

1. Paste the statement into the `PROBLEM` block comment.
2. Add any types the online judge supplies, such as `ListNode`. Add a parser
   that converts sample text into the Go values the algorithm needs and a
   formatter that converts the result back into the form shown in the problem.
   Numbers, strings, arrays, and maps can usually use
   `testrunner.FormatJSONOutput`; linked lists and trees need an explicit
   conversion.
3. Paste each example into the `sampleCases` slice using the included template.
4. Implement `solveCurrentAttempt`.
5. Ask Codex to review or solve "the problem in the workbench."

If you ask Codex to add the inputs or examples, it will also create or update
the problem-specific parser and call it at the start of your current
implementation. Your algorithm will receive ready-to-use Go values such as
integers, slices, or nodes rather than being left with raw sample text.

Run the active file at any point with:

```bash
./scripts/workbench.sh
```

The script prints `PASS` or `FAIL` for every solution/sample pair and then runs
the workbench tests.

The default review flow is:

1. Restate the required inputs and outputs, input limits, and important edge
   cases.
2. Review your current attempt without rewriting or discarding it.
3. Identify and link the relevant pages under `concepts/`.
4. Implement direct, improved, and preferred solutions in the workbench.
5. Run the same meaningful tests against your attempt and all three
   agent-authored solutions.
6. Explain why each approach is correct and how its work and extra memory grow
   in the analysis section of `workbench/main.go`.
7. Archive the completed solution with `archive-workbench.sh`, then run the
   archived problem and the repository checks.

Retrieved problem files retain source details and the original page URL for
attribution. Check the source's licensing terms before redistributing imported
problem statements.

## Archiving and running problems

Once the workbench is complete, choose a difficulty and save it as a standalone
problem file:

```bash
./scripts/archive-workbench.sh easy 1 two-sum
```

This creates `problems/solved_easy_leetcode_0001_two_sum.go`. Numeric problem
IDs receive leading zeroes until they contain at least four digits, which keeps
filenames in numeric order inside the same completion status, difficulty, and
source group. The `solved_` prefix marks files that preserve your attempt and a
completed agent review.

LeetCode is the default source; pass a fourth argument only for another source.
The original URL may remain in the problem comment for attribution. Before
archiving, Codex checks the supplied ID against the official problem page and
confirms that its title or statement matches the workbench. The archive command
then runs the workbench tests, including checks that the preserved attempt and
three agent solutions are registered.

Run one problem file or all completed files. An unfinished practice file uses
the same section layout as the workbench: official examples; helpers that
convert between sample text and Go values; a blank current-attempt function; an
empty list of agent solutions; and blank analysis headings. After implementing
the current attempt, run that file directly to check its test cases. The bulk
script skips unfinished files until a later review completes the solutions and
analysis and marks the file ready:

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
copy of the blank template. If any check fails, the current workbench remains
unchanged.

Problem files use `//go:build ignore`, so package-wide Go commands do not try to
combine their independent programs. Passing one filename directly to `go run`
still runs it with the shared `internal/testrunner` package.

## Validation

See the complete [command reference](#command-reference), including the
recommended full validation sequence and the workflow test that runs in a
temporary repository copy.
