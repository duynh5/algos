#!/usr/bin/env bash

set -euo pipefail

usage() {
  echo "usage: $0 [--case <number>] "\
"[<[solved_]difficulty_source_problem_id_name.go>]" >&2
}

fail() {
  echo "error: $1" >&2
  exit 1
}

is_valid_filename() {
  local filename_pattern
  filename_pattern='^(solved_)?(easy|medium|hard)_[a-z0-9]+_'
  filename_pattern+='([0-9]{4,}|[a-z0-9]*[a-z][a-z0-9]*)_'
  filename_pattern+='[a-z0-9]+(_[a-z0-9]+)*\.go$'
  [[ $1 =~ $filename_pattern ]]
}

script_directory=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
default_root=$(cd "$script_directory/../.." && pwd)
repository_root=${ALGOS_REPO_ROOT:-$default_root}
dsa_root="$repository_root/dsa"

detect_active_problem() {
  local candidate_path candidate_filename
  local active_file_count=0
  local -a active_files

  git -C "$repository_root" rev-parse \
    --is-inside-work-tree >/dev/null 2>&1 ||
    fail "$repository_root is not a Git worktree; provide a filename"

  while IFS= read -r candidate_path; do
    candidate_filename=${candidate_path#dsa/problems/}
    [[ $candidate_filename != "$candidate_path" ]] || continue
    [[ $candidate_filename != */* ]] || continue
    is_valid_filename "$candidate_filename" || continue
    [[ -f $dsa_root/problems/$candidate_filename ]] || continue
    active_files[$active_file_count]=$candidate_filename
    active_file_count=$((active_file_count + 1))
  done < <(
    {
      git -C "$repository_root" diff --name-only \
        --diff-filter=ACMR -- dsa/problems
      git -C "$repository_root" diff --cached --name-only \
        --diff-filter=ACMR -- dsa/problems
      git -C "$repository_root" ls-files --others \
        --exclude-standard -- dsa/problems
    } | LC_ALL=C sort -u
  )

  case $active_file_count in
    0)
      fail "no modified or untracked problem file found; "\
"provide a filename"
      ;;
    1)
      printf '%s\n' "${active_files[0]}"
      ;;
    *)
      echo "error: multiple modified or untracked problem files "\
"found; provide a filename:" >&2
      printf '  %s\n' "${active_files[@]}" >&2
      return 1
      ;;
  esac
}

case_number=
requested_file=
case_option_seen=false

while [[ $# -gt 0 ]]; do
  case $1 in
    --case)
      if [[ $case_option_seen == true || $# -lt 2 ]]; then
        usage
        exit 2
      fi
      case_option_seen=true
      case_number=$2
      shift 2
      ;;
    --case=*)
      if [[ $case_option_seen == true ]]; then
        usage
        exit 2
      fi
      case_option_seen=true
      case_number=${1#--case=}
      shift
      ;;
    --*)
      fail "unknown option: $1"
      ;;
    *)
      if [[ -n $requested_file ]]; then
        usage
        exit 2
      fi
      requested_file=$1
      shift
      ;;
  esac
done

if [[ $case_option_seen == true &&
  ! $case_number =~ ^[1-9][0-9]*$ ]]; then
  fail "--case must be a positive integer"
fi

if [[ -z $requested_file ]]; then
  filename=$(detect_active_problem)
else
  case $requested_file in
    dsa/problems/*) filename=${requested_file#dsa/problems/} ;;
    problems/*) filename=${requested_file#problems/} ;;
    */*) fail "provide a filename or dsa/problems/<filename>" ;;
    *) filename=$requested_file ;;
  esac
fi

is_valid_filename "$filename" ||
  fail "filename must match [solved_]<difficulty>_<source>_"\
"<problem_id>_<descriptive_slug>.go with a zero-padded numeric ID"
problem_file="$dsa_root/problems/$filename"

[[ -f $problem_file ]] || fail "$problem_file was not found"

(
  cd "$repository_root"
  ALGOS_RUN_PROBLEM_CASE_NUMBER=$case_number go run "$problem_file"
)
