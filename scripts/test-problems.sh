#!/usr/bin/env bash

set -euo pipefail

script_directory=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
default_root=$(cd "$script_directory/.." && pwd)
repository_root=${ALGOS_REPO_ROOT:-$default_root}

shopt -s nullglob
problem_files=(
  "$repository_root"/problems/easy_*.go
  "$repository_root"/problems/medium_*.go
  "$repository_root"/problems/hard_*.go
  "$repository_root"/problems/solved_easy_*.go
  "$repository_root"/problems/solved_medium_*.go
  "$repository_root"/problems/solved_hard_*.go
)

if [[ ${#problem_files[@]} -eq 0 ]]; then
  echo "No archived problem files found."
  exit 0
fi

for problem_file in "${problem_files[@]}"; do
  if rg -q '^const archiveReady = false$' "$problem_file"; then
    echo "Skipping $(basename "$problem_file") (unfinished practice archive)"
    continue
  fi
  echo "Running $(basename "$problem_file")"
  (cd "$repository_root" && go run "$problem_file")
done

echo "All completed archived problems passed."
