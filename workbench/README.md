# Active Problem Workbench

This directory holds one active problem at a time. Most
problem-specific work belongs in [`main.go`](main.go), whose labeled
sections contain the statement, test cases, the user's attempt, three
agent solutions, and explanations of their correctness, work, and
memory use. Shared code that runs every test case lives in
[`internal/testrunner`](../internal/testrunner/).

## Quick start

1. Paste the problem into the `PROBLEM` block comment in `main.go`.
2. Immediately below the statement, define any types supplied by the
   online judge, such as `ListNode`. Add a parser that converts each
   sample input into the Go values the algorithm needs and a formatter
   that converts the result back to the problem's expected text.
3. Add every supplied example to the `sampleCases` slice.
4. Implement `solveCurrentAttempt` using the parsed Go values.
5. Run `./internal/workbench.sh` from the repository root.
6. Ask Codex to review or solve "the problem in the workbench."

`./internal/workbench.sh` is equivalent to running `go run
./workbench` followed by `go test ./workbench`. The program prints
`PASS` or `FAIL` for every solution/sample pair before the script runs
the workbench tests.

When you ask Codex to add inputs or examples, it also creates or
updates the parser and calls it at the beginning of
`solveCurrentAttempt`. Your algorithm therefore begins with values
such as `[]int`, `string`, or `*ListNode` instead of unparsed sample
text. Codex also connects the matching output formatter without
replacing your algorithm body.

The agent sets `archiveReady` to `true` only after the completed
problem, the preserved current attempt, and all three agent-authored
solutions pass their checks. The archive script refuses incomplete
workbench content.

The shared runner gives each solution one input string and expects one
output string. For a site such as LeetCode, where the submitted
function receives several arguments, write those arguments in call
order as a JSON array--the bracketed text format shown below:

```go
Input: `[[2,7,11,15],9]`,
Want:  `[0,1]`,
```

For ordinary numbers, strings, arrays, and maps, the problem-specific
parser can decode each argument directly into its Go type:

```go
func parseTwoSumInput(input string) ([]int, int) {
	var nums []int
	var target int
	if err := testrunner.DecodeJSONArgs(
		input,
		&nums,
		&target,
	); err != nil {
		panic(fmt.Errorf("parse two sum input: %w", err))
	}
	return nums, target
}
```

Some problems use linked nodes, trees, graphs, or another
representation that cannot be written directly in a short sample
string. Keep the sample readable, then perform the conversion in the
parser. For example, a linked-list parser reads two integer arrays and
builds two `*ListNode` chains.

Every problem also has its own output formatter. Numbers, strings,
arrays, and maps can usually call `testrunner.FormatJSONOutput`. A
linked list, tree, or result checked by a special judging rule must
first be converted into the form shown in the problem statement.
Parsing and formatting failures should be reported at these boundaries
instead of being hidden inside the algorithm.

If copied problem text contains the uncommon sequence `*/`, change it
to `* /` so it does not close the Go block comment early.

During review, Codex preserves the user's attempt and evaluates it
separately from its three solutions. It first restates the required
inputs, outputs, limits, and important edge cases, then links relevant
material in [`CONCEPTS.md`](../CONCEPTS.md). It runs the same official
examples and additional bug-catching cases against all four
implementations. Detailed explanations of why each solution works and
how its work and memory grow remain in the analysis block in
`main.go`.

Once verified, the complete problem is copied to its permanent file
under [`problems/`](../problems/) with
`./internal/archive-workbench.sh <easy|medium|hard> <problem-id>
<descriptive-slug> [source]`. Completed files receive a `solved_`
filename prefix. The source defaults to `leetcode`, and numeric
problem IDs receive leading zeroes until they contain at least four
digits. Provide the optional final argument only for another source.
The archive command adds source metadata, runs the new standalone
file, and adds it to the problem index. It refuses an incomplete
workbench or an existing destination. Do not overwrite an unarchived
workbench problem.

After verifying the archived file, restore the standard blank template
with:

```bash
./internal/reset-workbench.sh
```

The reset command replaces both `main.go` and `main_test.go` and
refuses unfinished changes. Before replacement, it formats, runs, and
tests temporary copies of the blank templates; a failed check leaves
the active workbench unchanged. Use `--force` only when you explicitly
want to discard the current workbench without archiving it.
