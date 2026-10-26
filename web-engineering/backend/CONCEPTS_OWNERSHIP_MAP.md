# Backend Concepts Ownership Map

This document is the Phase 2 deliverable for the
[incorporation plan](CONCEPTS_INCORPORATION_PLAN.md). Inventory ranges
are inclusive. Every member retains its source-specific nuance in the
[source inventory](CONCEPTS_SOURCE_INVENTORY.md); grouping never makes
that nuance optional.

## Merge rules

- A canonical group owns one underlying decision or mechanism, not
  merely a shared word.
- Overview, detailed lesson, applied lesson, and refresher statements
  may merge only when they require the same invariant and failure
  behavior.
- A group's `Preserve` line names differences that its final owner
must
  retain.
- Guided work, failure tests, and completion evidence merge with the
  concept whose claim they prove.
- Cross-topic checklists and source lists remain supporting material;
  they do not become partial teaching entries.

## Canonical idea list and merge map

### C01. Learning, evidence, and curriculum use

Primary source owner: BE-00-001

Members:

- BE-00-001, BE-00-012, BE-00-040, BE-00-042 through BE-00-046
- BE-01-001, BE-16-001, BE-16-002, BE-16-030

Preserve: Evidence ranges from executable checks to external review
and production responsibility. Active recall and cross-cutting review
support the lessons but do not replace them.

Proposed guide owner: supporting introduction and review checklist.

### C02. Target Order API and initial technical choices

Primary source owner: BE-00-002

Members:

- BE-00-002, BE-00-004, BE-00-041

Preserve: The Order API is a learning context, its initial non-goals
are temporary scope controls, and frontend analogies do not transfer
backend authority or durability.

Proposed guide owner: supporting target-system section.

### C03. Go program organization and value mechanics

Primary source owner: BE-01-002

Members:

- BE-00-005
- BE-01-002 through BE-01-011, BE-01-015
- BE-02-002 through BE-02-005

Preserve: The refresher owns complete language behavior, while the
backend lesson adds package-boundary, API, copying, and testing
consequences. Slice aliasing, Unicode units, typed nil, and method
sets
remain independent failure cases inside one language owner.

Proposed guide owner: active Go values and ownership entry.

### C04. Go errors, resource lifetime, and cancellation

Primary source owner: BE-02-006

Members:

- BE-01-012
- BE-02-001, BE-02-006 through BE-02-008

Preserve: Error identity, panic boundaries, deterministic resource
release, context propagation, child deadlines, and detached-work
ownership solve different failure paths and all need concrete cases.

Proposed guide owner: active Go boundary-lifetime entry.

### C05. Go concurrency and executable verification

Primary source owner: BE-02-009

Members:

- BE-01-013, BE-01-014, BE-01-016, BE-01-017
- BE-02-009 through BE-02-012

Preserve: Memory-race freedom differs from business atomicity;
goroutine lifetime differs from channel closure; table, contention,
race, fuzz, and tool checks provide different evidence.

Proposed guide owner: active Go concurrency and testing entry.

### C06. Domain discovery and modeling

Primary source owner: BE-09-001

Members:

- BE-00-013
- BE-09-001 through BE-09-010
- BE-16-003

Preserve: Facts, assumptions, open questions, and deferred scope stay
separate. Examples, vocabulary, entities, values, commands, events,
state machines, invariants, and contexts keep their distinct roles.

Proposed guide owner: active domain discovery entry.

### C07. Architecture, representations, and service construction

Primary source owner: BE-05-001

Members:

- BE-00-003, BE-00-008, BE-00-014, BE-00-017
- BE-05-001 through BE-05-004, BE-05-006, BE-05-007, BE-05-009
- BE-05-011 through BE-05-013, BE-05-015
- BE-10-001, BE-10-004
- BE-16-004, BE-16-007

Preserve: Architecture localizes reasons to change without ceremonial
layers. Separate representations, vertical packages, narrow
consumer-owned interfaces, constructors, ADRs, and test boundaries
retain their own costs and failure signals.

Proposed guide owner: active architecture and construction entry.

### C08. HTTP contracts and process lifecycle

Primary source owner: BE-03-001

Members:

- BE-00-006, BE-00-015
- BE-03-001 through BE-03-013
- BE-10-002, BE-16-005

Preserve: Protocol semantics, network path, response contracts,
bounded server, strict decoding, validation ownership, middleware,
health, cancellation, shutdown, tests, and cursor pagination are
separate obligations within one public-boundary owner.

Proposed guide owner: active HTTP and lifecycle entry.

### C09. Relational modeling and PostgreSQL access

Primary source owner: BE-04-001

Members:

- BE-00-007, BE-00-016
- BE-04-001 through BE-04-008, BE-04-011 through BE-04-014
- BE-10-003, BE-16-006

Preserve: Durable modeling, constraint types, index cost, migration
compatibility, disposable environments, pool budgets, parameter
binding, integration boundaries, pagination, and plans remain
distinct decisions.

Proposed guide owner: active relational data entry.

### C10. Business transactions and contested state

Primary source owner: BE-04-009

Members:

- BE-00-018
- BE-04-009, BE-04-010
- BE-05-005, BE-05-008, BE-05-010
- BE-10-005, BE-16-008

Preserve: Input validation and current-state authority occur at
different times. Atomic conditional writes, deterministic lock order,
commit uncertainty, rollback, overflow, captured history, and
owner-scoped reads keep separate tests.

Proposed guide owner: active transactions and concurrency entry.

### C11. Identity, authorization, and API security

Primary source owner: BE-06-001

Members:

- BE-00-009
- BE-06-001 through BE-06-015

Preserve: Token purpose, validation conjunction, key rotation,
principal reduction, role checks, object predicates, token tradeoffs,
browser controls, data exposure, negative tests, and threat modeling
must not collapse into one middleware check.

Proposed guide owner: active identity and API security entry.

### C12. Deterministic tests and risk-owned boundaries

Primary source owner: BE-07-001

Members:

- BE-00-010
- BE-07-001 through BE-07-004, BE-07-013, BE-07-014

Preserve: Unit, handler, integration, and acceptance scopes differ;
tables, small fakes, injected time and identity, barriers, and bounded
waits make their evidence repeatable.

Proposed guide owner: active backend testing entry.

### C13. Observability and operational controls

Primary source owner: BE-07-005

Members:

- BE-00-020
- BE-07-005 through BE-07-008, BE-07-011, BE-07-012
- BE-11-002, BE-16-010

Preserve: Logs, metrics, traces, audits, SLIs, SLOs, health, failure
drills, alerts, and runbooks answer different questions. Cardinality,
privacy, denominator policy, and recovery proof remain explicit.

Proposed guide owner: active observability and operations entry.

### C14. Delivery, deployment, and recovery readiness

Primary source owner: BE-08-001

Members:

- BE-00-011, BE-00-021
- BE-08-001 through BE-08-008, BE-08-010 through BE-08-014
- BE-11-003, BE-16-011

Preserve: Artifact identity, container contents, startup config,
shared CI, schema compatibility, rollback direction, restore truth,
capacity, lifecycle rehearsal, incident response, and clean-clone
handoff remain separate gates.

Proposed guide owner: active delivery and recovery entry.

### C15. Idempotency and ambiguous write outcomes

Primary source owner: BE-11-001

Members:

- BE-00-019, BE-11-001, BE-16-009

Preserve: Key scope, canonical request identity, serialization,
atomic result storage, in-progress response, expiry, cleanup, and
disconnect/crash outcomes are all public correctness behavior.

Proposed guide owner: active idempotency entry.

### C16. Performance, capacity, retry, and overload

Primary source owner: BE-11-004

Members:

- BE-00-022
- BE-07-009, BE-07-010
- BE-08-009
- BE-11-004, BE-16-012

Preserve: Duration bounds, concurrency bounds, retry preconditions,
fleet capacity, representative load, coordinated omission, profiles,
one-change experiments, caching, overload, and recovery solve
different resource failures.

Proposed guide owner: active performance and resilience entry.

### C17. Advanced PostgreSQL operation and recovery

Primary source owner: BE-11-005

Members:

- BE-00-023, BE-11-005, BE-16-013

Preserve: MVCC and isolation govern anomalies; locks and deadlocks
govern waiting and abort; vacuum, statistics, bloat, checkpoints, and
WAL govern operation; backfill, PITR, replication, promotion, RPO, and
RTO govern different recovery claims.

Proposed guide owner: active PostgreSQL operations entry.

### C18. Durable background jobs and scheduling

Primary source owner: BE-12-001

Members:

- BE-00-024, BE-12-001, BE-16-014

Preserve: Durable intent, leases, fencing claims, handler idempotency,
retry classes, dead letters, replay, scheduler duplication,
backpressure, queue age, and shutdown remain independently testable.

Proposed guide owner: active durable work entry.

### C19. Reliable event publication and consumption

Primary source owner: BE-12-002

Members:

- BE-00-025, BE-12-002, BE-16-015

Preserve: Outbox atomicity differs from relay acknowledgement; inbox
atomicity differs from provider effects; ordering scope, contract
versioning, replay, retention, reconciliation, and exactly-once scope
remain explicit.

Proposed guide owner: active event delivery entry.

### C20. Distributed-systems mechanics

Primary source owner: BE-12-003

Members:

- BE-00-026, BE-12-003, BE-16-016

Preserve: Clock limits, ambiguous timeout, message faults,
partitions, leases, fencing, quorum assumptions, reconciliation,
hashing, and consensus each solve a different distributed problem.

Proposed guide owner: active distributed mechanics entry.

### C21. Service-boundary discovery and extraction

Primary source owner: BE-13-001

Members:

- BE-00-027, BE-05-014, BE-13-001, BE-16-017

Preserve: A bounded context is not automatically a service. Coupling,
ownership, scale, isolation, release, data autonomy, migration,
failure independence, and ongoing operational cost decide extraction.

Proposed guide owner: active service extraction entry.

### C22. Synchronous service communication

Primary source owner: BE-13-002

Members:

- BE-00-028, BE-13-002, BE-16-018

Preserve: Protocol choice, total deadline, remote operation identity,
pending local intent, reconciliation, workload authorization,
discovery, retry, breaker, bulkhead, concurrency, shedding, fallback,
and gateway roles remain distinct.

Proposed guide owner: active remote communication entry.

### C23. Distributed workflows and eventual consistency

Primary source owner: BE-13-003

Members:

- BE-00-029, BE-13-003, BE-16-019

Preserve: Saga state, versions, intermediate-state bounds, durable
timeouts, late and duplicate events, orchestration versus
choreography,
fallible compensation, reconciliation, operator view, and guarded
repair all remain represented.

Proposed guide owner: active distributed workflows entry.

### C24. Contract evolution and independent delivery

Primary source owner: BE-13-004

Members:

- BE-00-030, BE-13-004, BE-16-020

Preserve: Structural and semantic compatibility differ. Schema checks,
consumer expectations, both deployment orders, delayed events,
deprecation inventory, support window, telemetry, communication, and
adapter cleanup govern independent delivery.

Proposed guide owner: active contract evolution entry.

### C25. Distributed observability and resilience experiments

Primary source owner: BE-13-005

Members:

- BE-00-031, BE-13-005, BE-16-021

Preserve: Context propagation, safe baggage, stable metrics, queue
age, sampling, fan-out, retry storms, blast radius, abort condition,
recovery, and convergence form one end-to-end diagnostic contract.

Proposed guide owner: active distributed observability entry.

### C26. Data scaling, caching, and tenant isolation

Primary source owner: BE-14-001

Members:

- BE-00-032, BE-14-001, BE-16-022

Preserve: Cache freshness, replica lag, partition management,
archival, derived views, tenant-model tradeoffs, row-security pool
state, compound constraints, asynchronous scope, quotas, backups, and
operator paths remain separate risks.

Proposed guide owner: active scale and tenancy entry.

### C27. Security, privacy, and software supply chain

Primary source owner: BE-14-002

Members:

- BE-00-033, BE-14-002, BE-16-023

Preserve: User and workload identity, transport and action authority,
least privilege, overlap rotation, sensitive-data inventory,
minimization, export, deletion, backup policy, SBOM, provenance,
signing, and vulnerability response need different evidence.

Proposed guide owner: active privacy and supply-chain entry.

### C28. Cloud infrastructure and IAM

Primary source owner: BE-14-003

Members:

- BE-00-034, BE-14-003, BE-16-024

Preserve: Administrative and failure domains, packet paths, resource
authorization, private endpoints, human and workload identity,
encryption roles, IaC state, locking, policy, drift, audit, budgets,
cost, and teardown remain explicit.

Proposed guide owner: active cloud infrastructure entry.

### C29. Orchestration and platform operations

Primary source owner: BE-14-004

Members:

- BE-00-035, BE-14-004, BE-16-025

Preserve: Desired-state reconciliation does not create business
correctness. Requests, limits, probes, rollout, disruption, network
policy, termination, autoscaling, dependency budgets, and limited-
access diagnostics keep distinct behaviors.

Proposed guide owner: active platform operations entry.

### C30. Reliability engineering, incidents, and cost

Primary source owner: BE-14-005

Members:

- BE-00-036, BE-14-005, BE-16-026

Preserve: User outcomes, objective windows, error budgets, capacity,
failure domains, paging, response roles, communication, mitigation,
convergence, blameless review, corrective evidence, and total cost all
affect the reliability decision.

Proposed guide owner: active reliability and incidents entry.

### C31. Production integration boundaries

Primary source owner: BE-15-001

Members:

- BE-00-037, BE-15-001, BE-16-027

Preserve: Outbound ambiguity, quotas, sandbox proof, inbound signature
and replay, outbound webhook SSRF and retry, object streaming and
scanning, asynchronous notification state, and derived-view rebuild
each retain their own threat and failure behavior.

Proposed guide owner: active integrations entry.

### C32. Maintenance and legacy-system evolution

Primary source owner: BE-15-002

Members:

- BE-00-038, BE-15-002, BE-16-028

Preserve: Actual behavior evidence, characterization, refactor versus
migration, layer-specific upgrades, flag and canary lifecycle, shadow
side effects, repair scope and resume, regression diagnosis, recovery,
and temporary-code removal remain explicit.

Proposed guide owner: active maintenance entry.

### C33. Capstone delivery and engineering leadership

Primary source owner: BE-15-003

Members:

- BE-00-039, BE-15-003, BE-16-029

Preserve: Problem selection, alternatives, design review, finding
decisions, small increments, requirement traceability, relevant tests,
rollout, failure and recovery, operational handoff, peer review,
teaching, cost, residual risk, and removable complexity define the
leadership evidence.

Proposed guide owner: active engineering leadership entry.

### C34. Language portability

Primary source owner: BE-17-001

Members:

- BE-17-001 through BE-17-009

Preserve: Universal guarantees remain constant while module, value,
error, cancellation, concurrency, framework, ORM, runtime, and
toolchain behavior use ecosystem-native mechanisms. Java/Kotlin,
.NET, Node.js, Python, and Rust retain distinct traps.

Proposed guide owner: active language portability entry.

### C35. Sources and optional specialization

Primary source owner: BE-00-047

Members:

- BE-00-047, BE-00-048

Preserve: Provenance records support rather than replace teaching.
Optional specializations remain role-driven and must meet the same
contract, failure, security, evidence, and fit standard.

Proposed guide owner: supporting provenance and further-study section.

## Independent teaching ownership

Canonical groups C03 through C34 require active teaching ownership.
C01, C02, and C35 remain supporting material because they organize,
orient, or source the curriculum rather than teach one reusable
mechanism through a single trace.

No canonical group shares a primary owner. Cross-links may summarize
relationships, but the final guide must not reproduce another group's
full explanation.
