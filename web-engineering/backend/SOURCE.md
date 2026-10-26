# Backend Curriculum Provenance

The backend curriculum was migrated from:

```text
/Users/nolan/Development/be-standard
```

The source repository was clean at migration time. Its substantive
tracked content consisted of:

- One course README.
- Eighteen numbered study chapters.
- A small, intentionally incomplete Go product-store exercise.
- A standalone module declaration named `be-standard` using Go
  `1.26.3`.

The numbered chapters are preserved under `reference/foundations/` and
`reference/advanced/`. The Go exercise is preserved under
`labs/order-api/`. The README was adapted into the backend index.

The standalone `go.mod` was not copied because this repository already
has a module owner. Its module name and Go version are recorded above,
so no substantive source metadata is lost.

Git internals were not copied. The source repository remains unchanged
and retains its original history.

Formatting changes include prose reflow, table-to-list conversion,
link relocation, and code-line wrapping required by this repository's
70-character physical line limit. These changes do not intentionally
alter the curriculum's technical meaning.

The imported chapters are source reference material. The active
backend lessons are grouped in
[`CONCEPTS.md`](CONCEPTS.md) and follow the
[DSA teaching standards](../../dsa/CONCEPTS_STANDARDS.md) through the
shared web lesson profile.
