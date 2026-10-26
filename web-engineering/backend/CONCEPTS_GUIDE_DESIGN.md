# Backend Concepts Guide Design

This document records the Phase 3 structure for the
[incorporation plan](CONCEPTS_INCORPORATION_PLAN.md). Canonical group
ownership comes from the
[ownership map](CONCEPTS_OWNERSHIP_MAP.md).

## Progression and outline

The final guide keeps three supporting sections first:

1. `Contents` is navigation only.
2. `How to use this guide` owns the learning and evidence workflow.
3. `Backend competency map` becomes a short selection guide whose
   links point to the active owners below.

The active entries then appear in this dependency order:

1. `Go values, program structure, and ownership` owns C03.
2. `Go errors, resources, and cancellation` owns C04.
3. `Go concurrency and executable verification` owns C05.
4. `Domain discovery and modeling` owns C06.
5. `Architecture and service construction` owns C07.
6. `HTTP contracts and process lifecycle` owns C08.
7. `Relational modeling and PostgreSQL access` owns C09.
8. `Business transactions and contested state` owns C10.
9. `Identity, authorization, and API security` owns C11.
10. `Deterministic backend testing` owns C12.
11. `Observability and operational controls` owns C13.
12. `Delivery, deployment, and recovery readiness` owns C14.
13. `Idempotency and ambiguous write outcomes` owns C15.
14. `Performance, capacity, retry, and overload` owns C16.
15. `Advanced PostgreSQL operation and recovery` owns C17.
16. `Durable background jobs and scheduling` owns C18.
17. `Reliable event publication and consumption` owns C19.
18. `Distributed-systems mechanics` owns C20.
19. `Service-boundary discovery and extraction` owns C21.
20. `Synchronous service communication` owns C22.
21. `Distributed workflows and eventual consistency` owns C23.
22. `Contract evolution and independent delivery` owns C24.
23. `Distributed observability and resilience experiments` owns C25.
24. `Data scaling, caching, and tenant isolation` owns C26.
25. `Security, privacy, and software supply chain` owns C27.
26. `Cloud infrastructure and IAM` owns C28.
27. `Orchestration and platform operations` owns C29.
28. `Reliability engineering, incidents, and cost` owns C30.
29. `Production integration boundaries` owns C31.
30. `Maintenance and legacy-system evolution` owns C32.
31. `Capstone delivery and engineering leadership` owns C33.
32. `Language portability` owns C34.

The final supporting section, `Provenance and further study`, owns
C35. It links to the inventory, traceability matrix, and optional
specializations without restating their lessons.

## Classification

Groups C03 through C34 are active teaching entries. Each must contain
the eleven exact headings in the shared-standard order and use the
literal `Go example` sixth heading.

Groups C01, C02, and C35 are supporting material:

- C01 explains learning, evidence, and review.
- C02 defines the target Order API and initial scope.
- C35 preserves sources and optional study paths.

The competency map remains supporting only. After all active owners
are complete, it will keep selection clues and one-sentence
relationships, not full teaching prose.

## Dependency checks

- Go value behavior precedes Go errors, resource lifetime,
  cancellation, concurrency, and all Go examples.
- Domain language and invariants precede architecture, contracts, and
  persistence decisions.
- Architecture precedes HTTP adapters and database adapters.
- HTTP precedes identity middleware and public error mapping.
- Relational modeling precedes transactions and advanced PostgreSQL.
- Local transactions and idempotency precede jobs, events, and remote
  workflows.
- Testing precedes claims based on drills; observability precedes
  distributed diagnosis and incident response.
- Delivery precedes independent deployment, cloud, and orchestration.
- Message, time, lease, and partition mechanics precede service
  extraction and sagas.
- Service contracts precede cross-service compatibility and
  distributed observability.
- Security and tenant ownership foundations precede advanced privacy,
  cloud IAM, and platform policy.
- Maintenance precedes the capstone so characterization and cleanup
  are part of leadership evidence.
- Language portability comes last because it maps established
  guarantees into another ecosystem.

## Anchor and inbound-link migration

Repository search found three inbound backend lesson fragments, all
from `web-engineering/backend/README.md`. The current guide also links
to those fragments internally.

- Preserve `<a id="1-http-api-contracts"></a>` immediately before
  `HTTP contracts and process lifecycle`.
- Preserve
  `<a id="2-data-transactions-and-architecture"></a>` immediately
  before a short migration note linking to architecture, relational
  data, and transactions. The note is supporting material, not a
  teaching entry.
- Preserve
  `<a id="3-production-and-distributed-systems"></a>` immediately
  before `Observability and operational controls`, the first
  production-facing entry after testing. Its migration note links to
  delivery, idempotency, durable work, events, and distribution.

The new table of contents will link directly to every active heading.
The backend README will be updated in Phase 7 to expose the new groups
while the three legacy guide fragments continue resolving.

No current active-entry heading is renamed or removed until its
replacement entries pass their complete teaching contracts. The
explicit legacy anchors are added in the same change that removes the
old headings.

## Code ownership design

Complete implementations stay in the guide, the Order API lab, or the
repository-owned example packages. Active entries use the smallest
compiling Go fragment that makes one trace replayable, or link to the
tested package that owns it.

Examples that teach one reusable primitive may own a small package:

- strict JSON decoding;
- safe conditional in-memory reservation;
- bounded retry delay;
- canonical state transitions;
- fencing-token comparison; and
- tenant-scoped cache-key construction.

Operational entries whose mechanism cannot be represented honestly
in one small Go fragment link to the source implementation plan and
use a typed state-transition or validation fragment instead of
pretending to implement PostgreSQL, a broker, cloud IAM, or
Kubernetes.

## Table-of-contents design

The contents follows the exact 32-entry order above. Long labels use
reference-style links where necessary to preserve the 70-character
line limit. Supporting sections appear once and do not duplicate the
active-entry list in prose.
