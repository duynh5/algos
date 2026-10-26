# Web Engineering Lesson Standards

All active frontend and backend concept-study lessons follow
[`CONCEPTS_STANDARDS.md`](../CONCEPTS_STANDARDS.md). That document
owns the teaching method, example quality, correctness, cost analysis,
content ownership, navigation, and maintenance rules.

This profile defines only the web-curriculum adaptations. The shared
concepts-document checker enforces both frontend and backend guides.

## Required lesson contract

Every active lesson has one primary teaching entry. It uses these
headings once and in this order:

1. The problem it solves.
2. The idea in simple terms.
3. How it works.
4. Step-by-step examples.
5. Why it works.
6. A language-specific example.
7. Work and memory.
8. When to use it.
9. When not to use it.
10. Common mistakes and edge cases.
11. Related material.

The literal sixth heading is:

- `TypeScript example` for frontend lessons, including TSX where the
  browser or framework is relevant.
- `Go example` for backend lessons.

All other headings use the exact labels from the shared teaching
standards.

## Examples and evidence

- State exact inputs and expected outputs.
- Use Before, Decision, After, and Known correct in changing traces.
- Include one boundary, failure, duplicate, or adversarial example
  when it changes behavior.
- Keep examples small enough to replay without hidden state.
- Make code complete in its stated context and comment meaningful
  logical blocks.
- Name important browser, network, runtime, and production state.

## Correctness and cost

For algorithms, explain initialization, preservation, termination, and
the final result. For browser, framework, data, and operational
topics, explain why the demonstrated behavior follows from the
relevant contract or lifecycle.

Explain work, latency, memory, network transfer, browser rendering,
allocations, storage, and server load when they materially apply. Say
when a conventional Big-O expression would hide the more useful cost.

## Content ownership

Each subject has one active lesson owner. Indexes and career guides
link to that owner instead of repeating its full explanation. Imported
backend chapters remain reference material and are not active lessons.

Frontend teaching entries are grouped in
[`frontend/CONCEPTS.md`](frontend/CONCEPTS.md), which is their single
owning guide. Backend teaching entries use the same structure in
[`backend/CONCEPTS.md`](backend/CONCEPTS.md).

The algorithm guide continues to own algorithms and Go techniques.
Frontend lessons link to it for general methods instead of copying
those entries.
