# Repository Validation

Run repository commands from the repository root. The scripts require Bash,
Go, Git, `gofmt`, and `rg` (ripgrep). The module currently declares Go 1.26.3.

The validation commands have separate responsibilities:

- `go test ./...` runs every package test included in the normal build.
  Standalone files under `problems/` use `//go:build ignore`, so validate them
  separately with the problem scripts documented in
  [`problems/README.md`](problems/README.md).
- `go run internal/check-concepts-doc.go` validates the concepts guide's heading
  hierarchy, ordered teaching-entry fields, fenced blocks, baseline source
  links, local paths, internal fragments, and repository links into the guide.
- `./internal/test-problem-workflow.sh` tests archiving, resetting, filename
  checks, focused problem runs, and bulk completed-problem runs in a temporary
  repository copy. It does not modify the active workbench.

The usual full validation sequence is:

```bash
./internal/workbench.sh
./internal/test-problems.sh
./internal/test-problem-workflow.sh
go run internal/check-concepts-doc.go
go test ./...
```
