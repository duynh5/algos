#!/usr/bin/env bash

set -euo pipefail

script_directory=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
repository_root=$(cd "$script_directory/.." && pwd)
temporary_root=$(mktemp -d)
trap 'rm -rf "$temporary_root"' EXIT

mkdir -p "$temporary_root/workbench" "$temporary_root/problems"
cp "$repository_root/go.mod" "$temporary_root/go.mod"
mkdir -p "$temporary_root/internal"
cp -R "$repository_root/internal/testrunner" "$temporary_root/internal/testrunner"
cp "$repository_root/workbench/main.go.template" "$temporary_root/workbench/main.go.template"
cp "$repository_root/workbench/main.go.template" "$temporary_root/workbench/main.go"
cp "$repository_root/workbench/main_test.go.template" "$temporary_root/workbench/main_test.go.template"
cp "$repository_root/workbench/main_test.go.template" "$temporary_root/workbench/main_test.go"
cp "$repository_root/problems/README.md" "$temporary_root/problems/README.md"

ALGOS_REPO_ROOT=$temporary_root "$script_directory/reset-workbench.sh"

sed 's/^Source URL:$/Source URL: unarchived-test/' \
  "$repository_root/workbench/main.go.template" > "$temporary_root/workbench/main.go"
if ALGOS_REPO_ROOT=$temporary_root "$script_directory/reset-workbench.sh" >/dev/null 2>&1; then
  echo "error: unfinished workbench reset unexpectedly succeeded" >&2
  exit 1
fi
ALGOS_REPO_ROOT=$temporary_root "$script_directory/reset-workbench.sh" --force
cmp -s "$temporary_root/workbench/main.go.template" "$temporary_root/workbench/main.go"

if ALGOS_REPO_ROOT=$temporary_root "$script_directory/archive-workbench.sh" easy 1 two-sum >/dev/null 2>&1; then
  echo "error: incomplete workbench archive unexpectedly succeeded" >&2
  exit 1
fi

sed 's/^const archiveReady = false$/const archiveReady = true/' \
  "$repository_root/workbench/main.go.template" > "$temporary_root/workbench/main.go"
if ALGOS_REPO_ROOT=$temporary_root "$script_directory/archive-workbench.sh" easy 1 two-sum >/dev/null 2>&1; then
  echo "error: archive without three agent solutions unexpectedly succeeded" >&2
  exit 1
fi

# Build the archive fixture from the canonical template so this workflow test
# does not depend on whether the repository's active problem is complete.
awk '
  /^const archiveReady = false$/ {
    print "const archiveReady = true"
    next
  }
  /^var agentSolutions = \[\]testrunner.Solution\{\}$/ {
    print "var agentSolutions = []testrunner.Solution{"
    print "\t{Name: \"direct fixture\", Run: solveCurrentAttempt},"
    print "\t{Name: \"improved fixture\", Run: solveCurrentAttempt},"
    print "\t{Name: \"preferred fixture\", Run: solveCurrentAttempt},"
    print "}"
    next
  }
  { print }
' "$repository_root/workbench/main.go.template" > "$temporary_root/workbench/main.go"

cp "$temporary_root/workbench/main.go" "$temporary_root/main.go.before-failed-reset"
cp "$temporary_root/workbench/main_test.go.template" "$temporary_root/main_test.go.template.before-failed-reset"
printf '\nfunc TestResetCandidateFailure(t *testing.T) { t.Fatal("intentional reset validation failure") }\n' \
  >> "$temporary_root/workbench/main_test.go.template"
if ALGOS_REPO_ROOT=$temporary_root "$script_directory/reset-workbench.sh" >/dev/null 2>&1; then
  echo "error: reset with failing candidate tests unexpectedly succeeded" >&2
  exit 1
fi
if ! cmp -s "$temporary_root/main.go.before-failed-reset" "$temporary_root/workbench/main.go"; then
  echo "error: failed reset changed the workbench" >&2
  exit 1
fi
mv "$temporary_root/main_test.go.template.before-failed-reset" "$temporary_root/workbench/main_test.go.template"

ALGOS_REPO_ROOT=$temporary_root "$script_directory/archive-workbench.sh" easy 1 two-sum

archived_file="$temporary_root/problems/solved_easy_leetcode_0001_two_sum.go"
test -f "$archived_file"
test "$(sed -n '1p' "$archived_file")" = "//go:build ignore"
rg -q '^// Difficulty: Easy$' "$archived_file"
rg -q '^// Source: leetcode$' "$archived_file"
rg -q '^// Problem ID: 1$' "$archived_file"
rg -q 'solved_easy_leetcode_0001_two_sum.go' "$temporary_root/problems/README.md"

ALGOS_REPO_ROOT=$temporary_root "$script_directory/run-problem.sh" solved_easy_leetcode_0001_two_sum.go
ALGOS_REPO_ROOT=$temporary_root "$script_directory/run-problem.sh" problems/solved_easy_leetcode_0001_two_sum.go

git -C "$temporary_root" init --quiet
git -C "$temporary_root" config user.email "workflow-test@example.com"
git -C "$temporary_root" config user.name "Problem Workflow Test"
git -C "$temporary_root" add problems
git -C "$temporary_root" commit --quiet -m "create problem runner fixture"

if no_active_output=$(ALGOS_REPO_ROOT=$temporary_root "$script_directory/run-problem.sh" 2>&1); then
  echo "error: problem auto-detection without an active file unexpectedly succeeded" >&2
  exit 1
fi
if [[ $no_active_output != *"error: no modified or untracked problem file found"* ]]; then
  echo "error: problem auto-detection returned an unexpected no-file error:" >&2
  echo "$no_active_output" >&2
  exit 1
fi

second_problem="$temporary_root/problems/easy_leetcode_0002_add_two_numbers.go"
cp "$archived_file" "$second_problem"
ALGOS_REPO_ROOT=$temporary_root "$script_directory/run-problem.sh"
rm "$second_problem"

printf '\n' >> "$archived_file"
ALGOS_REPO_ROOT=$temporary_root "$script_directory/run-problem.sh"

cp "$archived_file" "$second_problem"
if ambiguous_output=$(ALGOS_REPO_ROOT=$temporary_root "$script_directory/run-problem.sh" 2>&1); then
  echo "error: ambiguous problem auto-detection unexpectedly succeeded" >&2
  exit 1
fi
if [[ $ambiguous_output != *"error: multiple modified or untracked problem files found"* ]]; then
  echo "error: problem auto-detection returned an unexpected ambiguity error:" >&2
  echo "$ambiguous_output" >&2
  exit 1
fi
rm "$second_problem"

if ALGOS_REPO_ROOT=$temporary_root "$script_directory/run-problem.sh" solved_easy_leetcode_1_two_sum.go >/dev/null 2>&1; then
	echo "error: unpadded numeric problem ID unexpectedly succeeded" >&2
	exit 1
fi
if ALGOS_REPO_ROOT=$temporary_root "$script_directory/run-problem.sh" easy_two_sum.go >/dev/null 2>&1; then
  echo "error: legacy archive filename unexpectedly succeeded" >&2
  exit 1
fi
ALGOS_REPO_ROOT=$temporary_root "$script_directory/test-problems.sh"

if ALGOS_REPO_ROOT=$temporary_root "$script_directory/archive-workbench.sh" easy 0001 two-sum >/dev/null 2>&1; then
  echo "error: duplicate archive unexpectedly succeeded" >&2
  exit 1
fi

if ALGOS_REPO_ROOT=$temporary_root "$script_directory/archive-workbench.sh" expert 1 invalid >/dev/null 2>&1; then
  echo "error: invalid difficulty unexpectedly succeeded" >&2
  exit 1
fi

if ALGOS_REPO_ROOT=$temporary_root "$script_directory/archive-workbench.sh" hard problem-1 invalid >/dev/null 2>&1; then
  echo "error: invalid problem ID unexpectedly succeeded" >&2
  exit 1
fi

if ALGOS_REPO_ROOT=$temporary_root "$script_directory/archive-workbench.sh" hard 1 invalid--slug >/dev/null 2>&1; then
  echo "error: invalid slug unexpectedly succeeded" >&2
  exit 1
fi

if ALGOS_REPO_ROOT=$temporary_root "$script_directory/archive-workbench.sh" hard 1 invalid CodeForces >/dev/null 2>&1; then
  echo "error: invalid source unexpectedly succeeded" >&2
  exit 1
fi

ALGOS_REPO_ROOT=$temporary_root "$script_directory/reset-workbench.sh"
cmp -s "$temporary_root/workbench/main.go.template" "$temporary_root/workbench/main.go"
cmp -s "$temporary_root/workbench/main_test.go.template" "$temporary_root/workbench/main_test.go"

echo "problem workflow tests passed"
