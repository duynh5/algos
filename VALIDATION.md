# Repository Validation

Run repository commands from the repository root. The scripts require
Bash, Go, Git, `gofmt`, `rg` (ripgrep), the TypeScript compiler
(`tsc`), and Prettier. The module currently declares Go 1.26.3.

The validation commands have separate responsibilities:

- `go test ./...` runs every package test included in the normal build. Standalone files under `dsa/problems/` use `//go:build ignore`, so validate them separately with the problem scripts documented in [`dsa/problems/README.md`](dsa/problems/README.md).
- `go run ./dsa/internal/conceptsdoc` validates the DSA,
  frontend, and backend concept guides' heading hierarchies, ordered
  teaching-entry fields, fenced blocks, local paths, and internal
  fragments. It preserves the DSA baseline source contract and checks
  repository links into every referenced `CONCEPTS.md` guide.
- `go run ./internal/codeexamples` checks every embedded Go block with
  `gofmt`, checks TypeScript and TSX formatting with Prettier, and
  type-checks each TypeScript example in isolation.
- `go run ./internal/linewidth` validates the 70-character physical
  line limit for code, embedded documentation code, and the three
  active concept guides. Ordinary repository prose is outside its
  scope.
- `./dsa/internal/test-problem-workflow.sh` tests archiving, resetting, filename checks, focused problem runs, and bulk completed-problem runs in a temporary repository copy. It does not modify the active workbench.

The usual full validation sequence is:

```bash
./dsa/internal/workbench.sh
./dsa/internal/test-problems.sh
./dsa/internal/test-problem-workflow.sh
go run ./dsa/internal/conceptsdoc
go run ./internal/codeexamples
go run ./internal/linewidth
go test ./...
```
