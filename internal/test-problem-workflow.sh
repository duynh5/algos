#!/usr/bin/env bash

set -euo pipefail

script_directory=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
repository_root=$(cd "$script_directory/.." && pwd)
temporary_root=$(mktemp -d)
trap 'rm -rf "$temporary_root"' EXIT

temporary_workbench="$temporary_root/workbench"
temporary_problems="$temporary_root/problems"
repository_workbench="$repository_root/workbench"
problem_readme="$temporary_problems/README.md"
link_targets="$temporary_root/CONCEPTS_LINKS.md"

run_reset_workbench() {
  ALGOS_REPO_ROOT=$temporary_root \
    "$script_directory/reset-workbench.sh" "$@"
}

run_archive_workbench() {
  ALGOS_REPO_ROOT=$temporary_root \
    "$script_directory/archive-workbench.sh" "$@"
}

run_problem() {
  ALGOS_REPO_ROOT=$temporary_root \
    "$script_directory/run-problem.sh" "$@"
}

mkdir -p "$temporary_workbench" "$temporary_problems"
cp "$repository_root/go.mod" "$temporary_root/go.mod"
mkdir -p "$temporary_root/internal"
cp -R "$repository_root/internal/testrunner" \
  "$temporary_root/internal/testrunner"
cp "$repository_workbench/main.go.template" \
  "$temporary_workbench/main.go.template"
cp "$repository_workbench/main.go.template" \
  "$temporary_workbench/main.go"
cp "$repository_workbench/main_test.go.template" \
  "$temporary_workbench/main_test.go.template"
cp "$repository_workbench/main_test.go.template" \
  "$temporary_workbench/main_test.go"
cp "$repository_root/problems/README.md" "$problem_readme"
cp "$repository_root/CONCEPTS_LINKS.md" "$link_targets"

run_reset_workbench

sed 's/^Source URL:$/Source URL: unarchived-test/' \
  "$repository_workbench/main.go.template" \
  > "$temporary_workbench/main.go"
if run_reset_workbench >/dev/null 2>&1; then
  echo "error: unfinished workbench reset unexpectedly succeeded" >&2
  exit 1
fi
run_reset_workbench --force
cmp -s "$temporary_workbench/main.go.template" \
  "$temporary_workbench/main.go"

if run_archive_workbench easy 1 two-sum >/dev/null 2>&1; then
  echo "error: incomplete workbench archive unexpectedly "\
"succeeded" >&2
  exit 1
fi

sed 's/^const archiveReady = false$/const archiveReady = true/' \
  "$repository_workbench/main.go.template" \
  > "$temporary_workbench/main.go"
if run_archive_workbench easy 1 two-sum >/dev/null 2>&1; then
  echo "error: archive without three agent solutions unexpectedly "\
"succeeded" >&2
  exit 1
fi

# Build the archive fixture from the canonical template. This keeps
# the workflow test independent of the active problem's state.
awk '
  /^var sampleCases = \[\]testrunner.Case\{$/ {
    print
    print "\t{Name: \"First fixture\", Input: \"first\", Want: \"\"},"
    print "\t{Name: \"Second fixture\", " \
      "Input: \"second\", Want: \"\"},"
    replacing_cases = 1
    next
  }
  replacing_cases && /^}$/ {
    print
    replacing_cases = 0
    next
  }
  replacing_cases { next }
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
' "$repository_workbench/main.go.template" \
  > "$temporary_workbench/main.go"

before_reset_main="$temporary_root/main.go.before-failed-reset"
before_reset_test="$temporary_root/test.go.before-failed-reset"
cp "$temporary_workbench/main.go" "$before_reset_main"
cp "$temporary_workbench/main_test.go.template" "$before_reset_test"
failing_test='func TestResetCandidateFailure(t *testing.T) {'
failing_test+=' t.Fatal("intentional reset validation failure") }'
printf '\n%s\n' "$failing_test" \
  >> "$temporary_workbench/main_test.go.template"
if run_reset_workbench >/dev/null 2>&1; then
  echo "error: reset with failing candidate tests unexpectedly "\
"succeeded" >&2
  exit 1
fi
if ! cmp -s "$before_reset_main" \
  "$temporary_workbench/main.go"; then
  echo "error: failed reset changed the workbench" >&2
  exit 1
fi
mv "$before_reset_test" \
  "$temporary_workbench/main_test.go.template"

run_archive_workbench easy 1 two-sum

archived_file="$temporary_problems/"\
"solved_easy_leetcode_0001_two_sum.go"
test -f "$archived_file"
test "$(sed -n '1p' "$archived_file")" = "//go:build ignore"
rg -q '^// Difficulty: Easy$' "$archived_file"
rg -q '^// Source: leetcode$' "$archived_file"
rg -q '^// Problem ID: 1$' "$archived_file"
rg -q 'solved_easy_leetcode_0001_two_sum.go' "$problem_readme"

long_slug="archive-link-that-exceeds-the-physical-line-width-limit"
run_archive_workbench easy 2 "$long_slug"
long_filename="solved_easy_leetcode_0002_"\
"archive_link_that_exceeds_the_physical_line_width_limit.go"
test -f "$temporary_problems/$long_filename"
rg -q 'problem-file-leetcode-0002' "$problem_readme"
rg -q "$long_filename" "$link_targets"
if awk 'length($0) > 70 { found = 1 } END { exit !found }' \
  "$problem_readme"; then
  echo "error: archive index contains an over-limit line" >&2
  exit 1
fi

run_problem solved_easy_leetcode_0001_two_sum.go
run_problem problems/solved_easy_leetcode_0001_two_sum.go

case_output=$(run_problem \
  --case 1 solved_easy_leetcode_0001_two_sum.go)
if [[ $case_output != *"PASS  current attempt / First fixture"* ||
  $case_output == *"Second fixture"* ]]; then
  echo "error: focused problem run returned unexpected output:" >&2
  echo "$case_output" >&2
  exit 1
fi
if out_of_range_output=$(run_problem \
  --case 3 solved_easy_leetcode_0001_two_sum.go 2>&1); then
  echo "error: out-of-range focused problem run unexpectedly "\
"succeeded" >&2
  exit 1
fi
expected_range_error="case 3 is out of range; 2 cases configured"
if [[ $out_of_range_output != *"$expected_range_error"* ]]; then
  echo "error: focused problem run returned an unexpected "\
"range error:" >&2
  echo "$out_of_range_output" >&2
  exit 1
fi
if run_problem --case 0 solved_easy_leetcode_0001_two_sum.go \
  >/dev/null 2>&1; then
  echo "error: non-positive focused case unexpectedly succeeded" >&2
  exit 1
fi

git -C "$temporary_root" init --quiet
git -C "$temporary_root" config user.email "workflow-test@example.com"
git -C "$temporary_root" config user.name "Problem Workflow Test"
git -C "$temporary_root" add problems
git -C "$temporary_root" commit --quiet \
  -m "create problem runner fixture"

if no_active_output=$(run_problem 2>&1); then
  echo "error: problem auto-detection without an active file "\
"unexpectedly succeeded" >&2
  exit 1
fi
no_active_error="error: no modified or untracked problem file found"
if [[ $no_active_output != *"$no_active_error"* ]]; then
  echo "error: problem auto-detection returned an unexpected "\
"no-file error:" >&2
  echo "$no_active_output" >&2
  exit 1
fi

second_problem="$temporary_problems/"\
"easy_leetcode_0002_add_two_numbers.go"
cp "$archived_file" "$second_problem"
run_problem
rm "$second_problem"

printf '\n' >> "$archived_file"
run_problem

cp "$archived_file" "$second_problem"
if ambiguous_output=$(run_problem 2>&1); then
  echo "error: ambiguous problem auto-detection unexpectedly "\
"succeeded" >&2
  exit 1
fi
ambiguous_error="error: multiple modified or untracked problem files"
if [[ $ambiguous_output != *"$ambiguous_error"* ]]; then
  echo "error: problem auto-detection returned an unexpected "\
"ambiguity error:" >&2
  echo "$ambiguous_output" >&2
  exit 1
fi
rm "$second_problem"

if run_problem solved_easy_leetcode_1_two_sum.go \
  >/dev/null 2>&1; then
	echo "error: unpadded numeric problem ID unexpectedly succeeded" >&2
	exit 1
fi
if run_problem easy_two_sum.go >/dev/null 2>&1; then
  echo "error: legacy archive filename unexpectedly succeeded" >&2
  exit 1
fi
ALGOS_REPO_ROOT=$temporary_root \
  "$script_directory/test-problems.sh"

if run_archive_workbench easy 0001 two-sum \
  >/dev/null 2>&1; then
  echo "error: duplicate archive unexpectedly succeeded" >&2
  exit 1
fi

if run_archive_workbench expert 1 invalid \
  >/dev/null 2>&1; then
  echo "error: invalid difficulty unexpectedly succeeded" >&2
  exit 1
fi

if run_archive_workbench hard problem-1 invalid \
  >/dev/null 2>&1; then
  echo "error: invalid problem ID unexpectedly succeeded" >&2
  exit 1
fi

if run_archive_workbench hard 1 invalid--slug \
  >/dev/null 2>&1; then
  echo "error: invalid slug unexpectedly succeeded" >&2
  exit 1
fi

if run_archive_workbench hard 1 invalid CodeForces \
  >/dev/null 2>&1; then
  echo "error: invalid source unexpectedly succeeded" >&2
  exit 1
fi

run_reset_workbench
cmp -s "$temporary_workbench/main.go.template" \
  "$temporary_workbench/main.go"
cmp -s "$temporary_workbench/main_test.go.template" \
  "$temporary_workbench/main_test.go"

echo "problem workflow tests passed"
