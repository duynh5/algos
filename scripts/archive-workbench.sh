#!/usr/bin/env bash

set -euo pipefail

usage() {
  echo "usage: $0 <easy|medium|hard> <problem-id> <descriptive-slug> [source]" >&2
}

fail() {
  echo "error: $1" >&2
  exit 1
}

if [[ $# -lt 3 || $# -gt 4 ]]; then
  usage
  exit 2
fi

difficulty=$1
problem_id=$2
slug=$3
source_name=${4:-leetcode}

case $difficulty in
  easy) difficulty_label=Easy ;;
  medium) difficulty_label=Medium ;;
  hard) difficulty_label=Hard ;;
  *) fail "difficulty must be easy, medium, or hard" ;;
esac

[[ $source_name =~ ^[a-z0-9]+$ ]] || fail "source must contain lowercase letters and digits"
[[ $problem_id =~ ^[a-z0-9]+$ ]] || fail "problem ID must contain lowercase letters and digits"
[[ $slug =~ ^[a-z0-9]+(-[a-z0-9]+)*$ ]] || fail "slug must contain lowercase letters, digits, and single hyphen separators"

filename_problem_id=$problem_id
if [[ $problem_id =~ ^[0-9]+$ ]]; then
  problem_id_number=$((10#$problem_id))
  printf -v filename_problem_id '%04d' "$problem_id_number"
fi

script_directory=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
default_root=$(cd "$script_directory/.." && pwd)
repository_root=${ALGOS_REPO_ROOT:-$default_root}
source_file="$repository_root/workbench/main.go"
problem_root="$repository_root/problems"
filename="solved_${difficulty}_${source_name}_${filename_problem_id}_${slug//-/_}.go"
destination="$problem_root/$filename"

[[ -f $source_file ]] || fail "$source_file was not found"
[[ -f "$problem_root/README.md" ]] || fail "$problem_root/README.md was not found"
[[ ! -e $destination ]] || fail "$destination already exists"
rg -q '^const archiveReady = true$' "$source_file" || fail "workbench is not marked ready to archive"

(cd "$repository_root" && go test ./workbench)

temporary_directory=$(mktemp -d "$repository_root/.archive-workbench.XXXXXX")
temporary_file="$temporary_directory/$filename"
cleanup() {
  rm -f "$temporary_file"
  rmdir "$temporary_directory" 2>/dev/null || true
}
trap cleanup EXIT

{
  printf '%s\n\n' '//go:build ignore'
  printf '// Difficulty: %s\n' "$difficulty_label"
  printf '// Source: %s\n' "$source_name"
  printf '// Problem ID: %s\n\n' "$problem_id"
  sed -n '1,$p' "$source_file"
} > "$temporary_file"

gofmt -w "$temporary_file"
(cd "$repository_root" && go run "$temporary_file")

mv "$temporary_file" "$destination"
printf -- '- [%s](%s)\n' "$filename" "$filename" >> "$problem_root/README.md"
rmdir "$temporary_directory"
trap - EXIT

echo "archived $destination"
echo "after verification, reset with ./scripts/reset-workbench.sh"
