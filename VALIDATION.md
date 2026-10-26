# Repository Validation

Run repository commands from the repository root. The scripts require Bash, Go, Git, `gofmt`, and `rg` (ripgrep). The module currently declares Go 1.26.3.

The validation commands have separate responsibilities:

- `go test ./...` runs every package test included in the normal build. Standalone files under `dsa/problems/` use `//go:build ignore`, so validate them separately with the problem scripts documented in [`dsa/problems/README.md`](dsa/problems/README.md).
- `go run dsa/internal/check-concepts-doc.go` validates the DSA guide's
  heading hierarchy, ordered teaching-entry fields, fenced blocks,
  baseline source links, local paths, and internal fragments. It also
  checks repository links into every referenced `CONCEPTS.md` guide.
- `go run ./internal/linewidth` validates the project-wide 70-character physical line limit. Only the five files named in `AGENTS.md` are exempt.
- `./dsa/internal/test-problem-workflow.sh` tests archiving, resetting, filename checks, focused problem runs, and bulk completed-problem runs in a temporary repository copy. It does not modify the active workbench.

The usual full validation sequence is:

```bash
./dsa/internal/workbench.sh
./dsa/internal/test-problems.sh
./dsa/internal/test-problem-workflow.sh
go run dsa/internal/check-concepts-doc.go
go run ./internal/linewidth
go test ./...
```
