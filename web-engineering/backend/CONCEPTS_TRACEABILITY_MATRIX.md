# Backend Concepts Traceability Matrix

This document is the Phase 6 deliverable for the
[incorporation plan](CONCEPTS_INCORPORATION_PLAN.md). It connects the
237 source records to 35 canonical groups and then to one active or
supporting destination.

## Reading the matrix

The [source inventory](CONCEPTS_SOURCE_INVENTORY.md) records exact
source files and sections. Each record now names a canonical group and
has a terminal status. The
[ownership map](CONCEPTS_OWNERSHIP_MAP.md) lists every member and the
primary source owner for that group.

The mappings below complete the chain from each group to an exact
guide
destination. Therefore a record resolves as:

1. inventory identifier to canonical group;
2. canonical group to primary source in the ownership map; and
3. canonical group to the guide anchor below.

## Supporting destinations

- C01 maps to [How to use this guide][support-use] and the competency
  map's cross-cutting review questions.
- C02 maps to the [backend competency map][support-map].
- C35 maps to [Provenance and further study][support-sources].

These groups orient, review, or source the curriculum. They do not
claim to be active teaching entries.

## Active destinations

- C03 maps to active entry 1, [Go values and ownership][entry-01].
- C04 maps to active entry 2, [errors and cancellation][entry-02].
-  C05 maps to active entry 3, [concurrency and
  verification][entry-03].
- C06 maps to active entry 4, [domain discovery][entry-04].
- C07 maps to active entry 5, [architecture][entry-05].
- C08 maps to active entry 6, [HTTP and process lifecycle][entry-06].
- C09 maps to active entry 7, [relational data][entry-07].
- C10 maps to active entry 8, [transactions][entry-08].
- C11 maps to active entry 9, [identity and security][entry-09].
- C12 maps to active entry 10, [deterministic testing][entry-10].
- C13 maps to active entry 11, [observability][entry-11].
- C14 maps to active entry 12, [delivery and recovery][entry-12].
- C15 maps to active entry 13, [idempotency][entry-13].
- C16 maps to active entry 14, [capacity and overload][entry-14].
- C17 maps to active entry 15, [PostgreSQL operations][entry-15].
- C18 maps to active entry 16, [durable jobs][entry-16].
- C19 maps to active entry 17, [reliable events][entry-17].
- C20 maps to active entry 18, [distributed mechanics][entry-18].
- C21 maps to active entry 19, [service extraction][entry-19].
- C22 maps to active entry 20, [service communication][entry-20].
- C23 maps to active entry 21, [distributed workflows][entry-21].
- C24 maps to active entry 22, [contract evolution][entry-22].
- C25 maps to active entry 23, [resilience experiments][entry-23].
- C26 maps to active entry 24, [scaling and tenancy][entry-24].
- C27 maps to active entry 25, [privacy and supply chain][entry-25].
-  C28 maps to active entry 26, [cloud infrastructure and
  IAM][entry-26].
- C29 maps to active entry 27, [platform operations][entry-27].
- C30 maps to active entry 28, [reliability and cost][entry-28].
- C31 maps to active entry 29, [integration boundaries][entry-29].
- C32 maps to active entry 30, [legacy evolution][entry-30].
- C33 maps to active entry 31, [engineering leadership][entry-31].
- C34 maps to active entry 32, [language portability][entry-32].

## Destination-to-source audit

Each active entry has one canonical source group and no group shares
an active owner. The content in its eleven teaching sections is
supported by that group's member records and preserved nuances. The
concrete trace and Go fragment or code-owner link demonstrate the same
group invariant; they do not introduce an unrelated concept.

The eleven exact section labels, explicit Before, Decision, After, and
Known correct trace, correctness argument, work and memory analysis,
and Related material links come from
[the shared standard](../../CONCEPTS_STANDARDS.md). The backend
process,
integration, operations, and boundary adaptations come from
[the lesson standard](../LESSON_STANDARDS.md). Those requirements are
repository standards rather than untracked source claims.

Cross-links in Related material point to another canonical owner. They
do not transfer teaching ownership. Repository-owned example links
provide executable evidence without replacing the active entry.

## Status audit

- Inventory records: 237.
- Incorporated primary owners: 35.
- Merged records naming their primary owner: 202.
- Pending records: 0.
- Reference-only records: 0.
- Not-an-idea records: 0.

No reference-only review is required. Repeated material remains
traceable as `merged into BE-..-...` in the inventory rather than
being
silently discarded.

## Bidirectional spot checks

- BE-01-006 maps to C03, whose primary is BE-01-002, then to entry 1.
  The source inventory retains slice-aliasing nuance; entry 1 traces
  an
  explicit clone and explains its `O(t)` copy cost.
- BE-03-003 maps to C08, whose primary is BE-03-001, then to entry 6.
  The source request-lifecycle rule appears in the bounded decoder and
  process shutdown trace.
- BE-11-001 maps to C15, whose primary is BE-11-001, then to entry 13.
  The active entry preserves request hashing, unique key ownership,
  replay, conflict, retention, and ambiguous commit behavior.
- Entry 17 maps backward to C19, its three inventory members, and
  primary BE-12-002. Its outbox and inbox trace preserves atomic
  publication intent, duplicate delivery, and local effect ownership.
- Entry 32 maps backward to C34 and BE-17-001 through BE-17-009. It
  preserves the universal invariant while requiring target-runtime
  evidence for ownership, errors, cancellation, and concurrency.

[support-use]: CONCEPTS.md#how-to-use-this-guide
[support-map]: CONCEPTS.md#backend-competency-map
[support-sources]: CONCEPTS.md#provenance-and-further-study
[entry-01]: CONCEPTS.md#1-go-values-program-structure-and-ownership
[entry-02]: CONCEPTS.md#2-go-errors-resources-and-cancellation
[entry-03]: CONCEPTS.md#3-go-concurrency-and-executable-verification
[entry-04]: CONCEPTS.md#4-domain-discovery-and-modeling
[entry-05]: CONCEPTS.md#5-architecture-and-service-construction
[entry-06]: CONCEPTS.md#6-http-contracts-and-process-lifecycle
[entry-07]: CONCEPTS.md#7-relational-modeling-and-postgresql-access
[entry-08]: CONCEPTS.md#8-business-transactions-and-contested-state
[entry-09]: CONCEPTS.md#9-identity-authorization-and-api-security
[entry-10]: CONCEPTS.md#10-deterministic-backend-testing
[entry-11]: CONCEPTS.md#11-observability-and-operational-controls
[entry-12]: CONCEPTS.md#12-delivery-deployment-and-recovery-readiness
[entry-13]: CONCEPTS.md#13-idempotency-and-ambiguous-write-outcomes
[entry-14]: CONCEPTS.md#14-performance-capacity-retry-and-overload
[entry-15]: CONCEPTS.md#15-advanced-postgresql-operation-and-recovery
[entry-16]: CONCEPTS.md#16-durable-background-jobs-and-scheduling
[entry-17]: CONCEPTS.md#17-reliable-event-publication-and-consumption
[entry-18]: CONCEPTS.md#18-distributed-systems-mechanics
[entry-19]: CONCEPTS.md#19-service-boundary-discovery-and-extraction
[entry-20]: CONCEPTS.md#20-synchronous-service-communication
[entry-21]:
  CONCEPTS.md#21-distributed-workflows-and-eventual-consistency
[entry-22]: CONCEPTS.md#22-contract-evolution-and-independent-delivery
[entry-23]:
  CONCEPTS.md#23-distributed-observability-and-resilience-experiments
[entry-24]: CONCEPTS.md#24-data-scaling-caching-and-tenant-isolation
[entry-25]:
  CONCEPTS.md#25-security-privacy-and-software-supply-chain
[entry-26]: CONCEPTS.md#26-cloud-infrastructure-and-iam
[entry-27]: CONCEPTS.md#27-orchestration-and-platform-operations
[entry-28]:
  CONCEPTS.md#28-reliability-engineering-incidents-and-cost
[entry-29]: CONCEPTS.md#29-production-integration-boundaries
[entry-30]: CONCEPTS.md#30-maintenance-and-legacy-system-evolution
[entry-31]:
  CONCEPTS.md#31-capstone-delivery-and-engineering-leadership
[entry-32]: CONCEPTS.md#32-language-portability
