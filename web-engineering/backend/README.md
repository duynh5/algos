# Backend Engineering Extension

This extension teaches backend engineering with Go through one
evolving Order API. Active lessons follow the shared
[lesson standards](../LESSON_STANDARDS.md).

Frontend candidates should first follow the frontend path from the
[web engineering guide](../README.md). Use this extension to develop
full-stack range, understand the systems behind browser applications,
or prepare for backend-heavy roles.

Reading alone is not completion. Each chapter asks for implementation,
failure testing, explanation, and production evidence.

## Active concept guide

The single [backend concept guide](CONCEPTS.md) owns all active
lessons. Its 32 entries form four dependency-ordered stages:

1. [Go language and execution foundations][go-stage], entries 1-3.
2. [Service and data foundations][service-stage], entries 4-12.
3. [Production and distributed systems][production-stage],
   entries 13-23.
4. [Scale, operations, and leadership][leadership-stage],
   entries 24-32.

The [source inventory](CONCEPTS_SOURCE_INVENTORY.md),
[ownership map](CONCEPTS_OWNERSHIP_MAP.md), and
[traceability matrix](CONCEPTS_TRACEABILITY_MATRIX.md) show how every
imported idea reaches one active or supporting destination. They are
audit artifacts rather than competing lessons.

Complete examples for entries 9 through 32 live under
[`examples/`](examples/). Each package has focused ordinary and
failure-path tests and is linked from its owning guide entry.

## Lab

The imported [Order API exercise](labs/order-api/) is intentionally
incomplete. Building it is part of the curriculum. The destination
repository's Go module owns compilation and testing.

See [migration provenance](SOURCE.md) for the precise source mapping
and excluded repository metadata.

[go-stage]: CONCEPTS.md#1-go-values-program-structure-and-ownership
[service-stage]: CONCEPTS.md#4-domain-discovery-and-modeling
[production-stage]:
  CONCEPTS.md#13-idempotency-and-ambiguous-write-outcomes
[leadership-stage]:
  CONCEPTS.md#24-data-scaling-caching-and-tenant-isolation
