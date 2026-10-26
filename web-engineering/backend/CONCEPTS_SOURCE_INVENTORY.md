# Backend Concepts Source Inventory

This document is the Phase 1 deliverable for the
[incorporation plan](CONCEPTS_INCORPORATION_PLAN.md). It records
source
ideas before Phase 2 merges repetition or assigns canonical owners.
Line locations preserve the imported files as they existed during the
inventory; the files were removed after incorporation completed.

## Chapter checklist

- [x] `foundations/00-overview-and-competency-map.md`
- [x] `foundations/01-go-language-refresher.md`
- [x] `foundations/02-go-backend-foundations.md`
- [x] `foundations/03-http-contracts-and-service-lifecycle.md`
- [x] `foundations/04-postgresql-and-relational-data.md`
- [x] `foundations/05-business-logic-and-architecture.md`
- [x] `foundations/06-identity-and-api-security.md`
- [x] `foundations/07-testing-observability-and-operations.md`
- [x] `foundations/08-delivery-and-production-readiness.md`
- [x] `advanced/09-domain-discovery.md`
- [x] `advanced/10-service-core.md`
- [x] `advanced/11-production-service.md`
- [x] `advanced/12-distributed-foundations.md`
- [x] `advanced/13-service-extraction.md`
- [x] `advanced/14-scale-security-cloud-operations.md`
- [x] `advanced/15-integrations-maintenance-leadership.md`
- [x] `advanced/16-key-concepts-review.md`
- [x] `advanced/17-language-portability.md`

## Ambiguity list

No source-scope ambiguity remains. The following classifications are
deliberate:

- Repeated overview, refresher, and detailed-lesson statements stay
  separate here. Phase 2 decides which item owns the explanation.
- Guided work and knowledge checks are recorded when they prescribe a
  reusable experiment or completion proof. Pure navigation is not an
  idea and is omitted.
- Tool and library names are implementation examples unless the
  source attaches a distinct contract or failure behavior to them.
- Optional specializations remain source ideas; Phase 2 may mark them
  reference-only only with an explicit reason.

## Inventory

### BE-00-001
Source: `foundations/00-overview-and-competency-map.md`
Source section: `Purpose` (lines 3-28)
Category: principle
Idea: Backend competence is demonstrated through safe, correct,
operable, simple, and then performant production behavior.
Required nuance: Working code is insufficient without review, failure
drills, evidence, and change safety; production readiness is not a
framework or directory shape.
Related IDs: BE-00-030, BE-15-007
Destination: C01, supporting introduction and review checklist
Status: incorporated

### BE-00-002
Source: `foundations/00-overview-and-competency-map.md`
Source section: `Target learning project` (lines 30-90)
Category: procedure
Idea: One modular-monolith Order API provides the shared learning
context for contracts, data, concurrency, security, and operations.
Required nuance: The initial endpoints each expose a named risk;
microservices, Kubernetes, custom identity, generic repositories, and
caching are initial non-goals rather than forbidden future choices.
Related IDs: BE-00-003, BE-00-009
Destination: C02, supporting competency map
Status: incorporated

### BE-00-003
Source: `foundations/00-overview-and-competency-map.md`
Source section: `Recommended architecture` (lines 91-178)
Category: principle
Idea: A vertical modular monolith keeps dependencies pointing inward
while matching structure to actual feature complexity.
Required nuance: Handlers translate, use cases own policy and business
transactions, adapters own mechanisms, and `main` owns construction;
boundaries must avoid both global layers and ceremonial wrappers.
Related IDs: BE-05-001, BE-10-001
Destination: C07, active entry 5
Status: merged into BE-05-001

### BE-00-004
Source: `foundations/00-overview-and-competency-map.md`
Source section: `Technology choices` (lines 179-255)
Category: tradeoff
Idea: Begin with an explicit, minimal Go and PostgreSQL toolset and
add
dependencies only for demonstrated correctness or operating value.
Required nuance: Keep SQL visible, parse configuration once, use
external identity, pin and review dependencies, wrap only real
boundaries, and avoid generic dumping-ground packages.
Related IDs: BE-00-003, BE-08-004
Destination: C02, supporting competency map
Status: merged into BE-00-002

### BE-00-005
Source: `foundations/00-overview-and-competency-map.md`
Source section: `Go as a backend language` (lines 264-303)
Category: language-specific behavior
Idea: Go backend fluency joins value ownership, errors, cleanup,
cancellation, and concurrency rather than treating syntax separately.
Required nuance: Race freedom does not prove business atomicity;
completion requires a tested store, bounded goroutines, toolchain
checks, and an explanation of values, interfaces, and wrapped errors.
Related IDs: BE-01-001, BE-02-001
Destination: C03, active entry 1
Status: merged into BE-01-002

### BE-00-006
Source: `foundations/00-overview-and-competency-map.md`
Source section: `HTTP contracts and request lifecycle` (lines 304-340)
Category: mechanism
Idea: HTTP is an end-to-end public and resource-lifecycle contract.
Required nuance: Method semantics, bounded decoding, middleware order,
stable errors, timeouts, cancellation, and shutdown must be traced
from DNS and sockets through the handler and response.
Related IDs: BE-03-001, BE-10-002
Destination: C08, active entry 6
Status: merged into BE-03-001

### BE-00-007
Source: `foundations/00-overview-and-competency-map.md`
Source section: `Relational data and PostgreSQL` (lines 341-374)
Category: invariant
Idea: PostgreSQL is the final shared authority for relational
integrity and concurrent business operations.
Required nuance: Schema constraints, bounded pools, parameterized SQL,
transactions, lock behavior, migrations, plans, and real-database
tests are one correctness and operations system.
Related IDs: BE-04-001, BE-10-003
Destination: C09, active entry 7
Status: merged into BE-04-001

### BE-00-008
Source: `foundations/00-overview-and-competency-map.md`
Source section: `Business logic and architecture` (lines 375-409)
Category: principle
Idea: Business invariants and business-sized transactions own the
system's policy independently of transport and storage details.
Required nuance: Representations may differ by boundary; interfaces
must be narrow and justified; last-unit contention is the decisive
proof rather than folder symmetry.
Related IDs: BE-05-002, BE-05-004
Destination: C07, active entry 5
Status: merged into BE-05-001

### BE-00-009
Source: `foundations/00-overview-and-competency-map.md`
Source section: `Identity and API security` (lines 410-460)
Category: operational control
Idea: Resource-server security verifies identity, then separately
authorizes actions and object access.
Required nuance: Token purpose, issuer, audience, algorithm, time,
rotation, session/browser concerns, least privilege, abuse limits,
redaction, and threat evidence remain separate controls.
Related IDs: BE-06-001, BE-06-002
Destination: C11, active entry 9
Status: merged into BE-06-001

### BE-00-010
Source: `foundations/00-overview-and-competency-map.md`
Source section: `Tests, observability, and operations` (lines 461-501)
Category: completion evidence
Idea: Each service claim needs evidence at the narrowest test or
operational boundary capable of proving it.
Required nuance: Deterministic tests, logs, metrics, traces, service
objectives, limits, retries, alerts, and runbooks answer different
questions and must distinguish concrete failure classes.
Related IDs: BE-07-001, BE-11-002
Destination: C12, active entry 10
Status: merged into BE-07-001

### BE-00-011
Source: `foundations/00-overview-and-competency-map.md`
Source section: `Delivery and production readiness` (lines 502-542)
Category: operational control
Idea: Delivery promotes one verified artifact through compatible,
observable, and recoverable runtime changes.
Required nuance: Runtime configuration, mixed schema versions,
rollback versus roll-forward, restore evidence, capacity, clean-clone
operation, and an incident drill all belong to readiness.
Related IDs: BE-08-001, BE-11-003
Destination: C14, active entry 12
Status: merged into BE-08-001

### BE-00-012
Source: `foundations/00-overview-and-competency-map.md`
Source section: `Applied engineering progression` (lines 543-649)
Category: procedure
Idea: Deliver the curriculum as sequential, reviewable vertical
increments whose acceptance evidence gates the next increment.
Required nuance: Each increment needs critical feedback, recorded
resolution, and the cross-cutting service-boundary lens even before
any network extraction is justified.
Related IDs: BE-00-001, BE-15-007
Destination: C01, supporting introduction and review checklist
Status: merged into BE-00-001

### BE-00-013
Source: `foundations/00-overview-and-competency-map.md`
Source section: `Domain discovery and modeling` (lines 650-694)
Category: procedure
Idea: Discover domain language, examples, states, invariants, and
ownership before choosing endpoints, tables, or services.
Required nuance: Separate facts from assumptions and unresolved
questions; review competing interpretations and treat context maps as
reasoning boundaries rather than deployment plans.
Related IDs: BE-09-001, BE-16-002
Destination: C06, active entry 4
Status: merged into BE-09-001

### BE-00-014
Source: `foundations/00-overview-and-competency-map.md`
Source section: `Decisions and executable skeleton` (lines 695-715)
Category: procedure
Idea: The first executable slice proves startup, configuration,
composition, shutdown, local checks, CI, and major decisions.
Required nuance: A clean clone and invalid configuration must produce
deterministic outcomes before feature complexity is added.
Related IDs: BE-10-001, BE-16-003
Destination: C07, active entry 5
Status: merged into BE-05-001

### BE-00-015
Source: `foundations/00-overview-and-competency-map.md`
Source section: `HTTP platform` (lines 716-735)
Category: procedure
Idea: Establish one bounded, documented HTTP platform before feature
handlers multiply.
Required nuance: Negative decoding paths, problem details,
correlation, health semantics, timeouts, and OpenAPI agreement are
completion evidence.
Related IDs: BE-03-004, BE-10-002
Destination: C08, active entry 6
Status: merged into BE-03-001

### BE-00-016
Source: `foundations/00-overview-and-competency-map.md`
Source section: `Persistence foundation` (lines 736-752)
Category: procedure
Idea: Add PostgreSQL as a bounded adapter with reproducible schema and
real integration evidence.
Required nuance: Constraints, context-aware queries, cleanup,
cancellation, and migrations from empty state are required before
feature persistence is considered complete.
Related IDs: BE-04-004, BE-10-003
Destination: C09, active entry 7
Status: merged into BE-04-001

### BE-00-017
Source: `foundations/00-overview-and-competency-map.md`
Source section: `Product slice` (lines 753-771)
Category: procedure
Idea: A complete product slice joins use case, handler, repository,
stable pagination, authorization, contract, and tests.
Required nuance: Duplicate identity, deterministic order, and the
distinction between unauthenticated and forbidden outcomes are part
of the public behavior.
Related IDs: BE-10-004, BE-16-006
Destination: C07, active entry 5
Status: merged into BE-05-001

### BE-00-018
Source: `foundations/00-overview-and-competency-map.md`
Source section: `Order transaction` (lines 772-789)
Category: invariant
Idea: Order, items, and inventory form one atomic local business
operation protected against concurrent overselling.
Required nuance: Deterministic contention tests, owner-scoped reads,
stable public errors, and no leaked database details prove the
boundary.
Related IDs: BE-04-009, BE-10-005
Destination: C10, active entry 8
Status: merged into BE-04-009

### BE-00-019
Source: `foundations/00-overview-and-competency-map.md`
Source section: `Idempotency and failure semantics` (lines 790-809)
Category: invariant
Idea: A client retry of one logical write must produce at most one
business effect despite disconnects and concurrent attempts.
Required nuance: Scope, canonical fingerprint, serialization, stored
outcome, conflict behavior, retention, and cleanup are one public
contract.
Related IDs: BE-11-001, BE-16-008
Destination: C15, active entry 13
Status: merged into BE-11-001

### BE-00-020
Source: `foundations/00-overview-and-competency-map.md`
Source section: `Observability and operational controls` (lines
810-829)
Category: operational control
Idea: Telemetry and runbooks must reconstruct user impact and its
first failing boundary with bounded, non-sensitive dimensions.
Required nuance: Version identity, audit evidence, pool pressure,
SLIs, actionable alerts, and a successful failure drill are required.
Related IDs: BE-11-002, BE-16-009
Destination: C13, active entry 11
Status: merged into BE-07-005

### BE-00-021
Source: `foundations/00-overview-and-competency-map.md`
Source section: `Deployment hardening` (lines 830-850)
Category: operational control
Idea: A minimal artifact must run with least privilege and survive
mixed-version schema rollout, termination, and data recovery.
Required nuance: Runtime-only secrets, resource budgets, source and
image scanning, and a verified restore distinguish proof from config.
Related IDs: BE-11-003, BE-16-010
Destination: C14, active entry 12
Status: merged into BE-08-001

### BE-00-022
Source: `foundations/00-overview-and-competency-map.md`
Source section: `Performance and resilience based on evidence`
(lines 851-885)
Category: test or experiment
Idea: Optimize one measured bottleneck under a reproducible,
representative workload and preserve correctness under overload.
Required nuance: Latency distributions, coordinated omission,
profiles, query plans, dependency time, bounded rejection, recovery,
and cache consistency all require separate evidence.
Related IDs: BE-11-004, BE-16-011
Destination: C16, active entry 14
Status: merged into BE-11-004

### BE-00-023
Source: `foundations/00-overview-and-competency-map.md`
Source section: `Advanced PostgreSQL operations and recovery`
(lines 886-936)
Category: operational control
Idea: Isolation, locks, maintenance, WAL, backfills, replicas, and
recovery jointly determine database correctness and availability.
Required nuance: Whole-operation retry, resumable verified backfills,
PITR, standby promotion, and measured RPO/RTO must be exercised rather
than inferred.
Related IDs: BE-11-005, BE-16-012
Destination: C17, active entry 15
Status: merged into BE-11-005

### BE-00-024
Source: `foundations/00-overview-and-competency-map.md`
Source section: `Durable background jobs and scheduling` (lines
937-975)
Category: mechanism
Idea: Durable jobs persist intent, lease bounded work, tolerate
at-least-once attempts, and make scheduler duplication safe.
Required nuance: Claim fencing, idempotent effects, retry classes,
dead letters, backpressure, queue age, replay, and shutdown all remain
observable and bounded.
Related IDs: BE-12-001, BE-16-013
Destination: C18, active entry 16
Status: merged into BE-12-001

### BE-00-025
Source: `foundations/00-overview-and-competency-map.md`
Source section: `Reliable event publication and consumption` (lines
976-1015)
Category: mechanism
Idea: An outbox and inbox bridge local transactions across a broker
without promising global exactly-once effects.
Required nuance: Immutable event identity, scoped ordering, duplicate
delivery, consumer atomicity, replay, retention, compatibility, and
lag/dead-letter evidence remain explicit.
Related IDs: BE-12-002, BE-16-014
Destination: C19, active entry 17
Status: merged into BE-12-002

### BE-00-026
Source: `foundations/00-overview-and-competency-map.md`
Source section: `Distributed-systems mechanics lab` (lines 1016-1064)
Category: test or experiment
Idea: Deterministic simulations expose delay, loss, duplication,
partitions, stale owners, quorum tradeoffs, hashing, and consensus.
Required nuance: Predict outcomes before execution, use fencing for
stale leases, define reconciliation, and delegate production
consensus to maintained components.
Related IDs: BE-12-003, BE-16-015
Destination: C20, active entry 18
Status: merged into BE-12-003

### BE-00-027
Source: `foundations/00-overview-and-competency-map.md`
Source section: `Service-boundary discovery and first extraction`
(lines 1065-1103)
Category: tradeoff
Idea: Extract only a capability with evidence for independent
ownership, scaling, isolation, or release value.
Required nuance: Data and contract ownership, branch-by-abstraction,
incremental migration, independent failure, and added on-call cost
must be measured; notification is the low-risk first case.
Related IDs: BE-13-001, BE-16-016
Destination: C21, active entry 19
Status: merged into BE-13-001

### BE-00-028
Source: `foundations/00-overview-and-competency-map.md`
Source section: `Synchronous service communication and the system
edge` (lines 1104-1149)
Category: mechanism
Idea: Remote calls require explicit contracts, deadline budgets,
idempotency, workload identity, discovery, and bounded failure
controls.
Required nuance: Protocol choice does not remove uncertainty; gateway,
retry, breaker, bulkhead, concurrency limit, shedding, and fallback
have distinct responsibilities.
Related IDs: BE-13-002, BE-16-017
Destination: C22, active entry 20
Status: merged into BE-13-002

### BE-00-029
Source: `foundations/00-overview-and-competency-map.md`
Source section: `Distributed workflows and eventual consistency`
(lines 1150-1190)
Category: mechanism
Idea: A saga models durable local transitions, expected intermediate
states, fallible compensation, timeouts, and repair.
Required nuance: Orchestration and choreography expose different
ownership; every boundary needs outbox/inbox, version guards,
idempotency, operator visibility, and bounded inconsistency.
Related IDs: BE-13-003, BE-16-018
Destination: C23, active entry 21
Status: merged into BE-13-003

### BE-00-030
Source: `foundations/00-overview-and-competency-map.md`
Source section: `Contract evolution and independent delivery` (lines
1191-1229)
Category: tradeoff
Idea: Independent deployment exists only inside a tested structural
and semantic compatibility window.
Required nuance: Consumer expectations, both deployment orders,
delayed events, deprecation ownership, usage evidence, and adapter
removal govern safe evolution.
Related IDs: BE-13-004, BE-16-019
Destination: C24, active entry 22
Status: merged into BE-13-004

### BE-00-031
Source: `foundations/00-overview-and-competency-map.md`
Source section: `Distributed observability and resilience validation`
(lines 1230-1269)
Category: test or experiment
Idea: Cross-service telemetry and controlled failure experiments must
identify the first failing boundary and prove convergence.
Required nuance: Trace and baggage propagation, sampling, retry
storms,
queue age, blast radius, abort conditions, and recovery checks must be
bounded and free of sensitive data.
Related IDs: BE-13-005, BE-16-020
Destination: C25, active entry 23
Status: merged into BE-13-005

### BE-00-032
Source: `foundations/00-overview-and-competency-map.md`
Source section: `Data scaling, caching, and multi-tenancy` (lines
1270-1312)
Category: tradeoff
Idea: Cache, replica, partition, archive, and tenant choices address
different measured constraints while adding distinct consistency risk.
Required nuance: Tenant identity must fail closed through synchronous,
asynchronous, cache, telemetry, quota, backup, and operator paths;
derived views never silently become truth.
Related IDs: BE-14-001, BE-16-021
Destination: C26, active entry 24
Status: merged into BE-14-001

### BE-00-033
Source: `foundations/00-overview-and-competency-map.md`
Source section: `Security, privacy, and software supply chain` (lines
1313-1353)
Category: operational control
Idea: Every architectural boundary expands identity, privacy, secret,
and artifact-trust obligations.
Required nuance: User and workload identity, least privilege,
rotation, data inventory and deletion, SBOM, provenance, signing, and
vulnerability response need end-to-end tests.
Related IDs: BE-14-002, BE-16-022
Destination: C27, active entry 25
Status: merged into BE-14-002

### BE-00-034
Source: `foundations/00-overview-and-competency-map.md`
Source section: `Cloud infrastructure and IAM` (lines 1354-1400)
Category: operational control
Idea: Reproducible cloud infrastructure separates connectivity,
human identity, workload identity, resource permissions, and failure
domains.
Required nuance: Private data services, controlled ingress and egress,
protected IaC state, policy, drift, budgets, audit, and safe teardown
must be independently verified.
Related IDs: BE-14-003, BE-16-023
Destination: C28, active entry 26
Status: merged into BE-14-003

### BE-00-035
Source: `foundations/00-overview-and-competency-map.md`
Source section: `Orchestration and platform operations` (lines
1401-1442)
Category: operational control
Idea: An orchestrator reconciles process state while the application
still owns valid startup, readiness, shutdown, and compatible data.
Required nuance: Requests, limits, probes, autoscaling, disruption,
rollout, network policy, and diagnostics must respect downstream
capacity and expose rather than hide failure behavior.
Related IDs: BE-14-004, BE-16-024
Destination: C29, active entry 27
Status: merged into BE-14-004

### BE-00-036
Source: `foundations/00-overview-and-competency-map.md`
Source section: `Reliability engineering, incidents, and cost` (lines
1443-1480)
Category: operational control
Idea: Reliability is a measured user outcome balanced against delivery
and total system cost through explicit objectives and error budgets.
Required nuance: Capacity, paging, incident roles, communication,
blameless review, corrective verification, engineering time, and
opportunity cost belong in the decision.
Related IDs: BE-14-005, BE-16-025
Destination: C30, active entry 28
Status: merged into BE-14-005

### BE-00-037
Source: `foundations/00-overview-and-competency-map.md`
Source section: `Common production integration boundaries` (lines
1481-1520)
Category: mechanism
Idea: Third-party calls, webhooks, objects, notifications, search, and
analytics require narrow, bounded, reconcilable adapters.
Required nuance: Ambiguous writes, signature and replay protection,
SSRF, streaming, provider callbacks, quotas, derived-view rebuild,
and sandbox evidence remain explicit.
Related IDs: BE-15-001, BE-16-026
Destination: C31, active entry 29
Status: merged into BE-15-001

### BE-00-038
Source: `foundations/00-overview-and-competency-map.md`
Source section: `Maintenance and legacy-system evolution` (lines
1521-1569)
Category: procedure
Idea: Discover and pin relied-on behavior before refactoring,
upgrading, migrating, repairing, or removing temporary paths.
Required nuance: Evidence comes from code, users, data, and
operations;
canaries, flags, repair scope, dry run, resume, audit, rollback, and
cleanup make change controlled.
Related IDs: BE-15-002, BE-16-027
Destination: C32, active entry 30
Status: merged into BE-15-002

### BE-00-039
Source: `foundations/00-overview-and-competency-map.md`
Source section: `Capstone delivery and engineering leadership` (lines
1570-1607)
Category: completion evidence
Idea: Engineering leadership makes requirements, tradeoffs, review,
implementation, operations, cost, and simplification traceable.
Required nuance: Small compatible increments, finding ownership,
relevant failure proof, handoff, peer review, teaching, and removal of
temporary complexity are required outcomes.
Related IDs: BE-15-003, BE-16-028
Destination: C33, active entry 31
Status: merged into BE-15-003

### BE-00-040
Source: `foundations/00-overview-and-competency-map.md`
Source section: `Concepts a backend contributor must understand`
(lines 1608-1934)
Category: principle
Idea: Contributors must connect domain, process, network, time, data,
transaction, API, identity, concurrency, error, observability,
delivery, cloud, distribution, scale, professional, and test concerns.
Required nuance: Each concern is a concrete ownership and failure
question; none is satisfied by vocabulary recognition alone.
Related IDs: BE-00-005 through BE-00-039
Destination: C01, supporting introduction and review checklist
Status: merged into BE-00-001

### BE-00-041
Source: `foundations/00-overview-and-competency-map.md`
Source section: `Frontend-to-backend mental model` (lines 1935-2001)
Category: principle
Idea: Frontend concepts map to backend concerns but change ownership,
trust, concurrency, durability, and operating scale.
Required nuance: A request crosses machines and failure domains;
server validation, shared-state integrity, resource bounds, and
recovery remain authoritative even when the UI has analogous state.
Related IDs: BE-03-001, BE-17-001
Destination: C02, supporting competency map
Status: merged into BE-00-002

### BE-00-042
Source: `foundations/00-overview-and-competency-map.md`
Source section: `Standard engineering workflow` (lines 2002-2045)
Category: procedure
Idea: Risky backend changes move from discovery through decision,
vertical implementation, evidence, review, rollout, and cleanup.
Required nuance: Each step records uncertainty, preserves
compatibility, watches runtime evidence, rehearses recovery, and
removes temporary mechanisms.
Related IDs: BE-00-012, BE-15-003
Destination: C01, supporting introduction and review checklist
Status: merged into BE-00-001

### BE-00-043
Source: `foundations/00-overview-and-competency-map.md`
Source section: `Production-readiness checklist` (lines 2046-2144)
Category: completion evidence
Idea: Readiness evidence spans contract, correctness, security,
reliability, delivery, cloud, distribution, data, privacy, and supply
chain.
Required nuance: The checklist verifies bounded and tested behavior;
it does not replace the owning lessons or their concrete evidence.
Related IDs: BE-00-001, BE-16-029
Destination: C01, supporting introduction and review checklist
Status: merged into BE-00-001

### BE-00-044
Source: `foundations/00-overview-and-competency-map.md`
Source section: `Foundation competency checkpoint` (lines 2145-2167)
Category: completion evidence
Idea: Safe contribution readiness means performing core contract,
data, identity, testing, telemetry, and delivery tasks without
guessing.
Required nuance: Weakness becomes an explicit pairing goal rather than
hidden confidence; the checkpoint is readiness to contribute, not
total mastery.
Related IDs: BE-00-045, BE-16-029
Destination: C01, supporting introduction and review checklist
Status: merged into BE-00-001

### BE-00-045
Source: `foundations/00-overview-and-competency-map.md`
Source section: `Definition of competent completion` (lines 2168-2237)
Category: completion evidence
Idea: Competence is repeatable, defensible evidence that another
engineer can review, deploy, operate, debug, and change the system.
Required nuance: Language, API, data, security, distribution,
operability, cloud, delivery, judgment, maintenance, and teamwork each
have distinct proof, while real production feedback remains necessary.
Related IDs: BE-00-001, BE-00-044
Destination: C01, supporting introduction and review checklist
Status: merged into BE-00-001

### BE-00-046
Source: `foundations/00-overview-and-competency-map.md`
Source section: `Questions to ask on an unfamiliar backend team`
(lines 2238-2264)
Category: procedure
Idea: Initial team discovery should surface invariants, ownership,
failure, messaging, security, retry, delivery, SLO, recovery, IAM, and
verification contracts before risky changes.
Required nuance: Questions seek measured and exercised behavior, not
only written policies.
Related IDs: BE-00-042, BE-15-002
Destination: C01, supporting introduction and review checklist
Status: merged into BE-00-001

### BE-00-047
Source: `foundations/00-overview-and-competency-map.md`
Source section: `Primary references by topic` (lines 2265-2348)
Category: completion evidence
Idea: Primary specifications and official documentation support the
guide's language, protocol, database, security, operations, cloud, and
distributed-systems claims.
Required nuance: Sources provide deeper authority and must remain
attached to the concepts they support rather than acting as substitute
lessons.
Related IDs: BE-00-001 through BE-00-046
Destination: C35, supporting sources and further study
Status: incorporated

### BE-00-048
Source: `foundations/00-overview-and-competency-map.md`
Source section: `Optional specialization paths` (lines 2349-2383)
Category: tradeoff
Idea: Specialization follows role needs and must preserve the core
standard of real use, failure semantics, security, evidence, and fit.
Required nuance: Payments, event sourcing, real-time, data-intensive,
global, platform, serverless, and protocol paths are optional rather
than universal architecture requirements.
Related IDs: BE-00-001, BE-17-001
Destination: C35, supporting sources and further study
Status: merged into BE-00-047

### BE-01-001
Source: `foundations/01-go-language-refresher.md`
Source section: introduction and `Learning outcomes` (lines 3-34)
Category: procedure
Idea: Rebuild Go fluency around value contents, mutation, sharing, and
failure flow by predicting, typing, and running one catalog example.
Required nuance: Recognition from reading is weaker than executable
recall; the lesson prepares language mechanics before backend layers.
Related IDs: BE-00-005, BE-02-001
Destination: C01, supporting introduction and review checklist
Status: merged into BE-00-001

### BE-01-002
Source: `foundations/01-go-language-refresher.md`
Source section: `Modules, packages, and the shape of a Go program`
(lines 35-118)
Category: language-specific behavior
Idea: Modules own dependency/version boundaries, packages own
directory compilation and visibility, and `main` starts executables.
Required nuance: Capitalization exports names, imports still control
access, one module commonly contains many packages, and `gofmt` plus
the Go commands are normal development tools.
Related IDs: BE-02-002, BE-17-003
Destination: C03, active entry 1
Status: incorporated

### BE-01-003
Source: `foundations/01-go-language-refresher.md`
Source section: `Variables, constants, and basic types` (lines
119-203)
Category: language-specific behavior
Idea: Go remains statically typed under inference and requires
explicit
numeric conversions at representation boundaries.
Required nuance: Use types by meaning, store exact money rather than
binary floats, distinguish byte and rune, and give enum zero values an
intentional meaning.
Related IDs: BE-04-002, BE-17-002
Destination: C03, active entry 1
Status: merged into BE-01-002

### BE-01-004
Source: `foundations/01-go-language-refresher.md`
Source section: `Control flow` (lines 204-303)
Category: language-specific behavior
Idea: Go uses explicit `if`, non-fallthrough `switch`, and four clear
forms of `for`, including `range` over collections.
Required nuance: Scoped initializers support guarded errors; blank
identifiers express deliberate omission; Go 1.22 creates distinct
variables for names declared by a range clause with `:=`.
Related IDs: BE-01-015, BE-17-003
Destination: C03, active entry 1
Status: merged into BE-01-002

### BE-01-005
Source: `foundations/01-go-language-refresher.md`
Source section: `Functions, multiple results, and closures`
(lines 304-380)
Category: language-specific behavior
Idea: Function signatures make inputs, multiple results, and errors
explicit while functions and closures remain ordinary values.
Required nuance: Variadic inputs become slices; named results can
obscure long functions; captured mutable state needs visible lifetime
and ownership, especially across goroutines.
Related IDs: BE-01-013, BE-02-006
Destination: C03, active entry 1
Status: merged into BE-01-002

### BE-01-006
Source: `foundations/01-go-language-refresher.md`
Source section: `Arrays, slices, and ownership` (lines 381-473)
Category: language-specific behavior
Idea: Arrays copy fixed storage while slices copy descriptors that may
still alias one backing array.
Required nuance: `append` may reuse or replace storage; ownership
boundaries require deliberate copies, nested reference-like fields
may need deeper copies, and small subslices can retain large arrays.
Related IDs: BE-02-004, BE-17-003
Destination: C03, active entry 1
Status: merged into BE-01-002

### BE-01-007
Source: `foundations/01-go-language-refresher.md`
Source section: `Maps` (lines 474-517)
Category: language-specific behavior
Idea: Maps distinguish missing keys from stored zero values only with
comma-`ok` and share mutable state across assignments.
Required nuance: Nil maps permit reads but panic on writes; iteration
order is unspecified; stable output needs sorting; concurrent writes
or read/write pairs need synchronization.
Related IDs: BE-01-013, BE-02-004
Destination: C03, active entry 1
Status: merged into BE-01-002

### BE-01-008
Source: `foundations/01-go-language-refresher.md`
Source section: `Strings, bytes, and runes` (lines 518-559)
Category: language-specific behavior
Idea: Strings are immutable bytes, while ranging decodes Unicode code
points at byte offsets rather than human-perceived characters.
Required nuance: String, byte-slice, and rune-slice conversions
usually allocate; requirements must choose bytes, code points, or
grapheme-like characters before truncation or indexing.
Related IDs: BE-01-003, BE-17-003
Destination: C03, active entry 1
Status: merged into BE-01-002

### BE-01-009
Source: `foundations/01-go-language-refresher.md`
Source section: `Structs, methods, pointers, and embedding`
(lines 560-647)
Category: language-specific behavior
Idea: Structs are values; receiver and pointer choices state whether
an operation observes a copy or mutates a shared instance.
Required nuance: Prefer keyed literals, use pointers for mutation,
identity, costly or unsafe copies, or meaningful absence, understand
method sets, and treat embedding as promotion rather than inheritance.
Related IDs: BE-02-003, BE-17-003
Destination: C03, active entry 1
Status: merged into BE-01-002

### BE-01-010
Source: `foundations/01-go-language-refresher.md`
Source section: `Interfaces, assertions, type switches, and generics`
(lines 648-720)
Category: language-specific behavior
Idea: Small consumer-owned interfaces are satisfied implicitly and
separate required behavior from concrete providers.
Required nuance: Return concrete values by default, check uncertain
type assertions, use type switches for several cases, and use generics
only when identical algorithms justify type abstraction.
Related IDs: BE-02-005, BE-05-005
Destination: C03, active entry 1
Status: merged into BE-01-002

### BE-01-011
Source: `foundations/01-go-language-refresher.md`
Source section: `Zero values, nil, make, and new` (lines 721-797)
Category: language-specific behavior
Idea: Every type has a zero value, while `make` initializes slice,
map, and channel runtime state and `new` returns a zero-value pointer.
Required nuance: Useful zero values reduce construction; an interface
with a dynamic pointer type and nil dynamic value is not itself nil;
successful interface returns must use literal nil.
Related IDs: BE-02-003, BE-05-006
Destination: C03, active entry 1
Status: merged into BE-01-002

### BE-01-012
Source: `foundations/01-go-language-refresher.md`
Source section: `Errors, defer, panic, and recover` (lines 798-891)
Category: language-specific behavior
Idea: Returned errors carry expected failure, wrapping preserves
identity, and `defer` attaches cleanup to the acquiring function.
Required nuance: Use `errors.Is` and `errors.As`, check meaningful
close errors, release loop resources promptly, reserve panic for
violated internal assumptions, and recover only at deliberate bounds.
Related IDs: BE-02-006, BE-02-007
Destination: C04, active entry 2
Status: merged into BE-02-006

### BE-01-013
Source: `foundations/01-go-language-refresher.md`
Source section: `Concurrency syntax and the ownership behind it`
(lines 892-972)
Category: language-specific behavior
Idea: Goroutines require owners and termination; channels transfer or
signal, while mutexes protect complete shared invariants.
Required nuance: Buffered channels are bounded queues, producers
normally close channels, cancellation is distinct from closure, and
atomic check-and-change must stay under one synchronization boundary.
Related IDs: BE-02-009, BE-12-001
Destination: C05, active entry 3
Status: merged into BE-02-009

### BE-01-014
Source: `foundations/01-go-language-refresher.md`
Source section: `Testing and everyday tools` (lines 973-1029)
Category: test or experiment
Idea: Table-driven tests make behavioral cases explicit and Go tools
provide progressively broader repeatable evidence.
Required nuance: Assert observable contracts, include input and wanted
output in failures, and remember that the race detector covers only
executed memory accesses, not all paths or business atomicity.
Related IDs: BE-02-010, BE-07-002
Destination: C05, active entry 3
Status: merged into BE-02-009

### BE-01-015
Source: `foundations/01-go-language-refresher.md`
Source section: `Practical traps and habits` (lines 1030-1095)
Category: failure mode
Idea: Shadowing, shallow copies, hidden setup, vague ownership, and
unstable output make otherwise valid Go programs surprising.
Required nuance: Use useful zero values and constructors only for real
invariants, keep context and setup explicit, avoid duplicate logging,
sort output, document closure, and let compiler tools expose defects.
Related IDs: BE-01-004, BE-01-006
Destination: C03, active entry 1
Status: merged into BE-01-002

### BE-01-016
Source: `foundations/01-go-language-refresher.md`
Source section: `Guided recall exercise` (lines 1096-1134)
Category: test or experiment
Idea: A catalog summary exercise integrates types, methods, ownership,
sorting, stable errors, and table tests through deliberate failures.
Required nuance: Missing keys, aliasing, invalid reservation, and map
order must be observed, corrected, and explained; passing tools are
evidence rather than the explanation.
Related IDs: BE-01-002 through BE-01-015
Destination: C05, active entry 3
Status: merged into BE-02-009

### BE-01-017
Source: `foundations/01-go-language-refresher.md`
Source section: `Knowledge check` (lines 1135-1181)
Category: completion evidence
Idea: Fluency requires explaining slice aliases, missing map keys,
pointers, typed nil, errors, cleanup, channels, Unicode, and race
limits.
Required nuance: Each answer must connect syntax to the underlying
value, ownership, or tool guarantee rather than repeat a rule.
Related IDs: BE-01-006 through BE-01-014
Destination: C05, active entry 3
Status: merged into BE-02-009

### BE-02-001
Source: `foundations/02-go-backend-foundations.md`
Source section: introduction and `How to use this lesson` (lines 3-55)
Category: principle
Idea: Backend-oriented Go choices repeatedly ask who owns data, who
may mutate it, how failures travel, and what overlapping work
observes.
Required nuance: Predict and test each boundary using one store rather
than memorizing isolated language features.
Related IDs: BE-01-001, BE-00-005
Destination: C04, active entry 2
Status: merged into BE-02-006

### BE-02-002
Source: `foundations/02-go-backend-foundations.md`
Source section: `Modules, packages, files, and names` (lines 56-205)
Category: language-specific behavior
Idea: Module, package, file, export, `internal`, and test-package
boundaries expose different parts of a Go API.
Required nuance: Unexported is package-wide rather than file-local;
constructors can protect private invariants; external tests exercise
public contracts; every export creates compatibility cost.
Related IDs: BE-01-002, BE-10-001
Destination: C03, active entry 1
Status: merged into BE-01-002

### BE-02-003
Source: `foundations/02-go-backend-foundations.md`
Source section: `Zero values, structs, and methods` (lines 206-278)
Category: language-specific behavior
Idea: Useful zero values and receiver semantics make initialization
and mutation contracts explicit.
Required nuance: Nil maps need construction before writes; value
receivers mutate copies; pointer receivers mutate a specific instance;
meaning matters before assumed performance.
Related IDs: BE-01-009, BE-01-011
Destination: C03, active entry 1
Status: merged into BE-01-002

### BE-02-004
Source: `foundations/02-go-backend-foundations.md`
Source section: `Values, pointers, slices, and maps` (lines 279-380)
Category: language-specific behavior
Idea: Choose values by default for clear small data and introduce
pointers or shared collections only with explicit aliasing semantics.
Required nuance: Slice capacity makes `append` aliasing unstable, map
assignment shares entries, callers must not receive internal mutable
state, and nested reference fields may require deeper copies.
Related IDs: BE-01-006, BE-01-007
Destination: C03, active entry 1
Status: merged into BE-01-002

### BE-02-005
Source: `foundations/02-go-backend-foundations.md`
Source section: `Interfaces describe required behavior` (lines
381-515)
Category: language-specific behavior
Idea: Consumers declare the smallest behavior they call, and concrete
providers satisfy it implicitly.
Required nuance: Broad provider interfaces burden fakes and unrelated
implementations; compiler assertions check method sets without runtime
allocation; pointer receiver method sets affect satisfaction.
Related IDs: BE-01-010, BE-05-005
Destination: C03, active entry 1
Status: merged into BE-01-002

### BE-02-006
Source: `foundations/02-go-backend-foundations.md`
Source section: `Errors are values with identity` (lines 516-633)
Category: language-specific behavior
Idea: Stable sentinel values or types survive contextual wrapping so
each layer can decide behavior without comparing text.
Required nuance: `errors.Is` finds identity, `errors.As` or `AsType`
extracts type, `%w` preserves the chain, and each layer adds only the
context it uniquely owns.
Related IDs: BE-01-012, BE-05-008
Destination: C04, active entry 2
Status: incorporated

### BE-02-007
Source: `foundations/02-go-backend-foundations.md`
Source section: `Resource ownership and defer` (lines 634-803)
Category: language-specific behavior
Idea: Successful acquisition transfers exactly-once cleanup
responsibility to a known function boundary.
Required nuance: Garbage collection does not promptly release scarce
resources; defer follows success, nested resources close in reverse
order, row iteration errors matter, and loop resources need short
scope.
Related IDs: BE-01-012, BE-04-007
Destination: C04, active entry 2
Status: merged into BE-02-006

### BE-02-008
Source: `foundations/02-go-backend-foundations.md`
Source section: `Context carries cancellation across boundaries`
(lines 804-999)
Category: mechanism
Idea: Context propagates one operation's cancellation and deadline
through every blocking boundary.
Required nuance: Cancellation is cooperative, child limits release
their timers, roots alone use `Background`, context values are narrow,
and storing request context or detaching work creates lifetime bugs.
Related IDs: BE-03-009, BE-13-002
Destination: C04, active entry 2
Status: merged into BE-02-006

### BE-02-009
Source: `foundations/02-go-backend-foundations.md`
Source section: `Concurrency` (lines 1000-1132)
Category: invariant
Idea: Synchronization must protect the whole decision whose invariant
would fail under interleaving.
Required nuance: Every goroutine is observable and bounded; channels
and mutexes solve different ownership problems; happens-before makes
writes visible; atomics do not compose multi-step rules; race freedom
and business atomicity require different tests.
Related IDs: BE-01-013, BE-04-009
Destination: C05, active entry 3
Status: incorporated

### BE-02-010
Source: `foundations/02-go-backend-foundations.md`
Source section: `Guided exercise: in-memory product store` (lines
1133-1350)
Category: test or experiment
Idea: A test-first store proves uniqueness, missing outcomes, safe
snapshots, deterministic order, and concurrent access.
Required nuance: Observe red before green, bound and join workers, run
the race detector, then add same-SKU contention proving exactly one
success rather than only absence of a memory race.
Related IDs: BE-01-014, BE-02-009
Destination: C05, active entry 3
Status: merged into BE-02-009

### BE-02-011
Source: `foundations/02-go-backend-foundations.md`
Source section: `Optional stretch: fuzz a validator` (lines 1351-1381)
Category: test or experiment
Idea: Fuzz generated SKU inputs to find panics and invariant
violations while asserting only rules that must hold for all cases.
Required nuance: Seed both valid and invalid cases and keep fuzzing
secondary to the required store and verification evidence.
Related IDs: BE-02-010, BE-07-001
Destination: C05, active entry 3
Status: merged into BE-02-009

### BE-02-012
Source: `foundations/02-go-backend-foundations.md`
Source section: `Knowledge check` and `Completion evidence` (lines
1382-1444)
Category: completion evidence
Idea: Completion joins tool success with explanations and contention
proof for ownership, interfaces, errors, cleanup, cancellation, and
concurrency.
Required nuance: Every goroutine has a bounded path, sentinel identity
survives wrapping, exactly one duplicate create succeeds, and one
mistake is recorded with its prevention rule.
Related IDs: BE-02-002 through BE-02-011
Destination: C05, active entry 3
Status: merged into BE-02-009

### BE-03-001
Source: `foundations/03-http-contracts-and-service-lifecycle.md`
Source section: introduction through `Before you start` (lines 3-79)
Category: principle
Idea: An HTTP endpoint is a public conversation whose
caller-controlled
input and resource use must be bounded at every transition.
Required nuance: Client, server, request, response, handler, and
middleware have distinct roles; bytes must be limited and translated
before application behavior runs.
Related IDs: BE-00-006, BE-10-002
Destination: C08, active entry 6
Status: incorporated

### BE-03-002
Source: `foundations/03-http-contracts-and-service-lifecycle.md`
Source section: `Follow one request through the system` (lines 80-137)
Category: mechanism
Idea: Requests cross DNS, TCP, TLS, proxy, server, middleware,
handler, downstream, and response boundaries with separate owners.
Required nuance: Each stage consumes time, connections, descriptors,
memory, or goroutines; keep-alive saves setup only when client
response
bodies are closed and reusable.
Related IDs: BE-00-041, BE-02-007
Destination: C08, active entry 6
Status: merged into BE-03-001

### BE-03-003
Source: `foundations/03-http-contracts-and-service-lifecycle.md`
Source section: `Methods express semantics` (lines 138-205)
Category: principle
Idea: HTTP methods describe safety and idempotency independently of
routing convenience, status repetition, or cacheability.
Required nuance: Safe operations request no business change;
idempotent repetition has one intended effect; a lost write response
makes non-idempotent POST outcomes ambiguous.
Related IDs: BE-11-001, BE-13-002
Destination: C08, active entry 6
Status: merged into BE-03-001

### BE-03-004
Source: `foundations/03-http-contracts-and-service-lifecycle.md`
Source section: `Status codes, headers, and representations`
(lines 206-279)
Category: mechanism
Idea: Status, headers, representation, and stable problem code form
one response contract.
Required nuance: Headers precede status, creation returns location,
codes distinguish application conditions sharing a status, and public
responses never expose Go, SQL, paths, stacks, or secrets.
Related IDs: BE-03-007, BE-05-008
Destination: C08, active entry 6
Status: merged into BE-03-001

### BE-03-005
Source: `foundations/03-http-contracts-and-service-lifecycle.md`
Source section: `Start a bounded Go HTTP server` (lines 280-337)
Category: operational control
Idea: An explicit server bounds header read, body read, response
write, idle connection, and header-memory lifetimes.
Required nuance: Timeouts are failure bounds rather than performance
targets; values follow measured needs and surrounding deadlines;
`main` constructs the server and dependencies.
Related IDs: BE-02-008, BE-08-008
Destination: C08, active entry 6
Status: merged into BE-03-001

### BE-03-006
Source: `foundations/03-http-contracts-and-service-lifecycle.md`
Source section: `Decode JSON strictly and within a limit`
(lines 338-427)
Category: operational control
Idea: Untrusted JSON must pass media-type, byte, shape, unknown-field,
single-value, and field-rule checks before reaching business logic.
Required nuance: Limit during reading rather than after allocation;
parse media types with parameters; classify oversize separately; log
diagnostics privately while returning stable public problems.
Related IDs: BE-06-010, BE-10-002
Destination: C08, active entry 6
Status: merged into BE-03-001

### BE-03-007
Source: `foundations/03-http-contracts-and-service-lifecycle.md`
Source section: `Separate validation responsibilities` and `Return
consistent problem details` (lines 428-514)
Category: principle
Idea: Handlers validate transport representation while use cases own
business validity, and the edge maps stable outcomes to problems.
Required nuance: A syntactically valid request may fail current-state
rules; business packages remain independent of JSON and HTTP; clients
branch on stable codes rather than human detail text.
Related IDs: BE-05-002, BE-05-008
Destination: C08, active entry 6
Status: merged into BE-03-001

### BE-03-008
Source: `foundations/03-http-contracts-and-service-lifecycle.md`
Source section: `Middleware order is observable behavior`
(lines 515-567)
Category: mechanism
Idea: Nested middleware changes which context and response state each
concern can observe on the way in and out.
Required nuance: Request identity precedes logging and recovery;
access logging records the actual downstream status; recovery is a
correlated boundary, not a bug mask; logs use route templates and
exclude raw bodies and query strings.
Related IDs: BE-06-005, BE-07-005
Destination: C08, active entry 6
Status: merged into BE-03-001

### BE-03-009
Source: `foundations/03-http-contracts-and-service-lifecycle.md`
Source section: `Liveness, readiness, and shutdown` (lines 568-621)
Category: operational control
Idea: Liveness, readiness, request cancellation, and graceful
shutdown answer different lifecycle questions.
Required nuance: Shutdown withdraws traffic, stops intake, drains,
then closes dependencies within the platform window; `Shutdown` does
not itself cancel active handler contexts or force completion.
Related IDs: BE-07-008, BE-08-009
Destination: C08, active entry 6
Status: merged into BE-03-001

### BE-03-010
Source: `foundations/03-http-contracts-and-service-lifecycle.md`
Source section: `Guided build` (lines 622-649)
Category: procedure
Idea: Add health, product read/write, strict decoding, stable
creation,
middleware, bounds, and shutdown in an order that keeps failures
local.
Required nuance: Use `net/http` directly so protocol behavior remains
visible before introducing framework abstraction.
Related IDs: BE-03-001 through BE-03-009
Destination: C08, active entry 6
Status: merged into BE-03-001

### BE-03-011
Source: `foundations/03-http-contracts-and-service-lifecycle.md`
Source section: `Handler tests with httptest` (lines 650-697)
Category: test or experiment
Idea: In-process handler tests cover public HTTP mapping without a
real listening socket.
Required nuance: Exercise valid, media, syntax, unknown, trailing,
oversize, required-field, duplicate, panic, and request-ID cases;
assert stable fields rather than timestamps or whole log lines.
Related IDs: BE-07-001, BE-10-002
Destination: C08, active entry 6
Status: merged into BE-03-001

### BE-03-012
Source: `foundations/03-http-contracts-and-service-lifecycle.md`
Source section: `Opaque cursor pagination` (lines 698-725)
Category: mechanism
Idea: A cursor encodes the last position in a deterministic total
order rather than a mutable row offset.
Required nuance: Validate opaque shape, cap page size, use a unique
tie-breaker such as ID, and return values strictly after the full
ordered pair to prevent gaps or repeats.
Related IDs: BE-04-012, BE-10-004
Destination: C08, active entry 6
Status: merged into BE-03-001

### BE-03-013
Source: `foundations/03-http-contracts-and-service-lifecycle.md`
Source section: `Knowledge check` and `Completion evidence` (lines
726-775)
Category: completion evidence
Idea: HTTP completion requires deterministic negative tests, explicit
limits, safe errors, a full request trace, middleware explanation,
bounded shutdown evidence, and Go checks.
Required nuance: One observed failure must be tied to the boundary
that now prevents or contains it.
Related IDs: BE-03-001 through BE-03-012
Destination: C08, active entry 6
Status: merged into BE-03-001

### BE-04-001
Source: `foundations/04-postgresql-and-relational-data.md`
Source section: introduction through `Before you start` (lines 3-92)
Category: invariant
Idea: Begin with a business rule, encode its last shared authority in
PostgreSQL, and prove it under concurrency and failure.
Required nuance: Tables, rows, columns, constraints, queries,
connections, transactions, migrations, and indexes are distinct
responsibilities serving the no-oversell promise.
Related IDs: BE-00-007, BE-04-009
Destination: C09, active entry 7
Status: incorporated

### BE-04-002
Source: `foundations/04-postgresql-and-relational-data.md`
Source section: `Relational modeling begins with invariants`
(lines 93-130)
Category: principle
Idea: Tables represent durable business facts rather than serialized
Go structs.
Required nuance: Normalize independent identity and lifecycle while
duplicating historical facts deliberately, such as captured order
price, so later source changes do not rewrite history.
Related IDs: BE-05-002, BE-09-006
Destination: C09, active entry 7
Status: merged into BE-04-001

### BE-04-003
Source: `foundations/04-postgresql-and-relational-data.md`
Source section: `Keys, constraints, and nullability` (lines 131-222)
Category: invariant
Idea: Primary, foreign, unique, check, and not-null constraints answer
different integrity questions at the final write boundary.
Required nuance: Null adds a meaningful third state, application
checks improve feedback but cannot resolve races, and tenant ownership
may need compound keys and foreign constraints.
Related IDs: BE-14-001, BE-10-003
Destination: C09, active entry 7
Status: merged into BE-04-001

### BE-04-004
Source: `foundations/04-postgresql-and-relational-data.md`
Source section: `Indexes are not free` (lines 223-255)
Category: tradeoff
Idea: Each index trades storage and write maintenance for a specific
filtering or ordering path.
Required nuance: Composite column order follows query shape, unique
constraints commonly already own an index, and adding every possible
index increases cost without ensuring useful plans.
Related IDs: BE-04-012, BE-11-004
Destination: C09, active entry 7
Status: merged into BE-04-001

### BE-04-005
Source: `foundations/04-postgresql-and-relational-data.md`
Source section: `Migrations are ordered production changes`
(lines 256-311)
Category: procedure
Idea: Versioned migrations reproduce schema while preserving data and
mixed-version application compatibility.
Required nuance: Destructive down migrations are not universal
rollback, additive nullable expansion precedes backfill and required
constraints, and migrations run once as an owned deployment step.
Related IDs: BE-08-005, BE-11-003
Destination: C09, active entry 7
Status: merged into BE-04-001

### BE-04-006
Source: `foundations/04-postgresql-and-relational-data.md`
Source section: `Run a disposable local PostgreSQL` (lines 312-339)
Category: operational control
Idea: Local and integration PostgreSQL must be pinned, health-checked,
documented, and disposable.
Required nuance: Development credentials are non-production, named
volumes may aid manual work, and tests start clean so old state cannot
hide migration errors.
Related IDs: BE-08-011, BE-10-003
Destination: C09, active entry 7
Status: merged into BE-04-001

### BE-04-007
Source: `foundations/04-postgresql-and-relational-data.md`
Source section: `A connection pool is a bounded shared resource`
(lines 340-384)
Category: operational control
Idea: A pool is a fleet-wide concurrency budget and queue in front of
a limited database.
Required nuance: Bound size, lifetime, startup ping, acquisition, and
query with context; multiply per-instance limits; measure wait and
database saturation before increasing capacity.
Related IDs: BE-07-009, BE-11-004
Destination: C09, active entry 7
Status: merged into BE-04-001

### BE-04-008
Source: `foundations/04-postgresql-and-relational-data.md`
Source section: `Parameterized SQL keeps code separate from data`
(lines 385-457)
Category: operational control
Idea: Parameters prevent values from changing SQL syntax, while
dynamic identifiers require fixed allowlists.
Required nuance: Translate only known database identities, keep
unexpected errors visible, close rows, and check iteration errors;
token claims remain data rather than trusted syntax.
Related IDs: BE-06-010, BE-17-005
Destination: C09, active entry 7
Status: merged into BE-04-001

### BE-04-009
Source: `foundations/04-postgresql-and-relational-data.md`
Source section: `Transactions preserve a complete business operation`
and `Prevent overselling atomically` (lines 458-572)
Category: invariant
Idea: One order transaction commits reservation, order, and items
together, with the availability check and decrement in one protected
database decision.
Required nuance: Remote effects stay outside, commit errors matter,
transactions remain short, deterministic lock order reduces but does
not eliminate deadlocks, and retries repeat only a whole safe
operation.
Related IDs: BE-05-007, BE-11-005
Destination: C10, active entry 8
Status: incorporated

### BE-04-010
Source: `foundations/04-postgresql-and-relational-data.md`
Source section: `Prove rollback behavior` (lines 573-593)
Category: test or experiment
Idea: Force failure after inventory change and verify every tentative
write disappears, then race two buyers for the last unit.
Required nuance: Inspect order, items, and original inventory; expect
one commit, one insufficient result, and final stock zero.
Related IDs: BE-04-009, BE-05-010
Destination: C10, active entry 8
Status: merged into BE-04-009

### BE-04-011
Source: `foundations/04-postgresql-and-relational-data.md`
Source section: `Integration-test boundaries` (lines 594-610)
Category: test or experiment
Idea: Only real PostgreSQL proves its syntax, constraints, isolation,
locks, transaction behavior, plans, and driver interaction.
Required nuance: Tests own deterministic isolated data and bounded
dependency readiness; mocks remain suitable only for risks they
actually reproduce.
Related IDs: BE-05-010, BE-07-001
Destination: C09, active entry 7
Status: merged into BE-04-001

### BE-04-012
Source: `foundations/04-postgresql-and-relational-data.md`
Source section: `Pagination and query plans` (lines 611-659)
Category: test or experiment
Idea: Stable tuple cursors and representative execution plans prove
ordered data access and index fit.
Required nuance: Unique tie-breakers, one-extra-row pagination,
estimated versus actual rows, loops, filters, sorts, buffers, skew,
and N+1 queries all affect evidence; sequential scans can be correct.
Related IDs: BE-03-012, BE-11-004
Destination: C09, active entry 7
Status: merged into BE-04-001

### BE-04-013
Source: `foundations/04-postgresql-and-relational-data.md`
Source section: `Guided build` (lines 660-681)
Category: procedure
Idea: Build pinned PostgreSQL, all migrations, parameterized product
operations, readiness, atomic order placement, integration tests, and
plan evidence sequentially.
Required nuance: Map only known database conditions and keep readiness
fast, generic, and free of connection details.
Related IDs: BE-04-001 through BE-04-012
Destination: C09, active entry 7
Status: merged into BE-04-001

### BE-04-014
Source: `foundations/04-postgresql-and-relational-data.md`
Source section: `Knowledge check` and `Completion evidence` (lines
682-740)
Category: completion evidence
Idea: Database completion requires schema reproduction, direct-write
rejection, cancellation, cleanup, rollback, contention, and explained
plan evidence.
Required nuance: Record the observed failure and database rule that
prevents it; formatting, unit, integration, race, and vet checks all
remain required.
Related IDs: BE-04-001 through BE-04-013
Destination: C09, active entry 7
Status: merged into BE-04-001

### BE-05-001
Source: `foundations/05-business-logic-and-architecture.md`
Source section: introduction through `Architecture is assignment of
responsibility` (lines 3-63)
Category: principle
Idea: Architecture assigns each decision to the code that owns its
representation, policy, durable mechanism, or lifecycle.
Required nuance: A boundary earns its place when a concrete HTTP,
business, SQL, or wiring change stays local; symmetric folders and
unexplained types are not evidence.
Related IDs: BE-00-003, BE-05-003
Destination: C07, active entry 5
Status: incorporated

### BE-05-002
Source: `foundations/05-business-logic-and-architecture.md`
Source section: `Architecture manages reasons to change`
(lines 64-101)
Category: principle
Idea: Keep code together when it changes for one reason and separate
different owners, vocabularies, invariants, or dependencies.
Required nuance: One deployable and database can retain strong module
boundaries; both cross-cutting edits and empty wrapper proliferation
signal the wrong structure.
Related IDs: BE-05-001, BE-13-001
Destination: C07, active entry 5
Status: merged into BE-05-001

### BE-05-003
Source: `foundations/05-business-logic-and-architecture.md`
Source section: `One concept can need several representations`
(lines 102-180)
Category: mechanism
Idea: Transport requests, application commands, domain results, and
persistence rows may represent one concept differently at trust and
ownership transitions.
Required nuance: Client-controlled fields are allowlisted, identity
comes from the trusted principal, adapters contain driver shapes, and
types are not duplicated mechanically when ownership is truly shared.
Related IDs: BE-03-007, BE-06-010
Destination: C07, active entry 5
Status: merged into BE-05-001

### BE-05-004
Source: `foundations/05-business-logic-and-architecture.md`
Source section: `Organize vertical feature packages` (lines 181-226)
Category: principle
Idea: Vertical packages keep a capability's transport, use case,
adapter, and tests near each other while dependencies point inward.
Required nuance: Platform packages own only stable shared mechanisms;
domain code avoids HTTP, drivers, configuration, and telemetry; vague
utility packages hide ownership.
Related IDs: BE-00-003, BE-10-004
Destination: C07, active entry 5
Status: merged into BE-05-001

### BE-05-005
Source: `foundations/05-business-logic-and-architecture.md`
Source section: `Write invariants before orchestration` (lines
227-299)
Category: invariant
Idea: State business rules before control flow and place each rule at
its earliest useful and final authoritative boundary.
Required nuance: Input-only rules precede connection acquisition;
current-state rules stay inside the transaction; arithmetic is
checked;
states use defined transitions rather than arbitrary strings.
Related IDs: BE-09-005, BE-09-006
Destination: C10, active entry 8
Status: merged into BE-04-009

### BE-05-006
Source: `foundations/05-business-logic-and-architecture.md`
Source section: `Keep interfaces narrow and consumer-owned`
(lines 300-350)
Category: principle
Idea: A business-sized consumer interface exposes the one atomic
capability a use case needs rather than generic transaction or CRUD
mechanics.
Required nuance: The use case chooses the transaction boundary while
the adapter implements it; separate use cases may own different small
interfaces on one concrete adapter; fakes prove coordination, not SQL.
Related IDs: BE-02-005, BE-05-010
Destination: C07, active entry 5
Status: merged into BE-05-001

### BE-05-007
Source: `foundations/05-business-logic-and-architecture.md`
Source section: `Dependency injection is ordinary construction`
(lines 351-413)
Category: language-specific behavior
Idea: Constructors expose mandatory dependencies and `main` supplies
concrete implementations without a service locator or DI framework.
Required nuance: Deterministic clocks and IDs improve tests; typed-nil
interfaces bypass simple nil checks; programmer configuration errors
fail at construction while runtime failures return errors.
Related IDs: BE-01-011, BE-10-001
Destination: C07, active entry 5
Status: merged into BE-05-001

### BE-05-008
Source: `foundations/05-business-logic-and-architecture.md`
Source section: `Design the PlaceOrder use case` (lines 414-481)
Category: procedure
Idea: The application validates the command and delegates exactly one
atomic placement operation to the persistence adapter.
Required nuance: Copy mutable inputs, inject identity and time, sort
locks, reserve and capture prices, check overflow, insert all records,
and commit once; rollback removes earlier reservations on failure.
Related IDs: BE-04-009, BE-10-005
Destination: C10, active entry 8
Status: merged into BE-04-009

### BE-05-009
Source: `foundations/05-business-logic-and-architecture.md`
Source section: `Stable error identities cross boundaries`
(lines 482-539)
Category: mechanism
Idea: Adapters translate known mechanisms to application identities,
services preserve them while adding context, and handlers map them to
public problems.
Required nuance: Missing-product policy is explicit, unexpected errors
stay unexpected, internal text aids operators, clients see only stable
status and code, and error strings never drive branching.
Related IDs: BE-02-006, BE-03-004
Destination: C07, active entry 5
Status: merged into BE-05-001

### BE-05-010
Source: `foundations/05-business-logic-and-architecture.md`
Source section: `Object ownership belongs in the query`
(lines 540-569)
Category: operational control
Idea: Query by both object and trusted owner so unauthorized data
never
enters application memory.
Required nuance: Wrong-owner and nonexistent rows usually share one
not-found outcome, preventing existence leaks and forgotten later
checks.
Related IDs: BE-06-007, BE-10-005
Destination: C10, active entry 8
Status: merged into BE-04-009

### BE-05-011
Source: `foundations/05-business-logic-and-architecture.md`
Source section: `Test at the boundary that owns the risk`
(lines 570-624)
Category: test or experiment
Idea: Unit, PostgreSQL integration, and handler tests prove
application
coordination, durable behavior, and public mapping respectively.
Required nuance: Each layer deliberately omits risks owned elsewhere;
duplicating the whole workflow in every test slows feedback and makes
failures harder to locate.
Related IDs: BE-04-011, BE-07-001
Destination: C07, active entry 5
Status: merged into BE-05-001

### BE-05-012
Source: `foundations/05-business-logic-and-architecture.md`
Source section: `Guided refactor and build` (lines 625-644)
Category: procedure
Idea: Refactor product and order features sequentially, preserve one
atomic order operation, then remove all stale paths and duplicates.
Required nuance: Search constructors, tests, mocks, imports, and old
package strings after moving ownership; run all boundary-specific
verification.
Related IDs: BE-05-004 through BE-05-011
Destination: C07, active entry 5
Status: merged into BE-05-001

### BE-05-013
Source: `foundations/05-business-logic-and-architecture.md`
Source section: `Optional architecture decision record`
(lines 645-668)
Category: completion evidence
Idea: An ADR records context, alternatives, decision, consequences,
evidence, and the conditions that would make a locking choice stop
fit.
Required nuance: It preserves reasoning rather than declaring victory
or copying implementation details.
Related IDs: BE-10-001, BE-15-003
Destination: C07, active entry 5
Status: merged into BE-05-001

### BE-05-014
Source: `foundations/05-business-logic-and-architecture.md`
Source section: `Microservice counterfactual` (lines 669-696)
Category: tradeoff
Idea: Replacing one local transaction with services creates ambiguous
partial failure, durable identity, compensation, delivery, and
observability obligations.
Required nuance: Record the future boundary but keep modules
in-process
until ownership, scale, isolation, or release evidence pays its cost.
Related IDs: BE-13-001, BE-13-002
Destination: C21, active entry 19
Status: merged into BE-13-001

### BE-05-015
Source: `foundations/05-business-logic-and-architecture.md`
Source section: `Knowledge check` and `Completion evidence` (lines
697-748)
Category: completion evidence
Idea: Architecture completion is proven by dependency direction,
atomic delegation, justified interfaces, boundary tests, stable
errors,
owner-scoped reads, and an honest service counterfactual.
Required nuance: The last-unit and rollback tests remain the decisive
data-safety evidence; all configured checks must pass.
Related IDs: BE-05-001 through BE-05-014
Destination: C07, active entry 5
Status: merged into BE-05-001

### BE-06-001
Source: `foundations/06-identity-and-api-security.md`
Source section: introduction through `Before you start` (lines 3-100)
Category: principle
Idea: Security is a trust pipeline from untrusted credential to
verified principal, allowed operation, and scoped object access.
Required nuance: OAuth, OIDC, JWT, JWKS, issuer, audience, subject,
CORS, and CSRF name distinct roles; one person's request still needs
three separate identity, action, and object decisions.
Related IDs: BE-00-009, BE-06-002
Destination: C11, active entry 9
Status: incorporated

### BE-06-002
Source: `foundations/06-identity-and-api-security.md`
Source section: `Authentication and authorization are separate`
(lines 101-129)
Category: principle
Idea: Authentication establishes a principal; authorization decides
whether that principal may perform one action on one object.
Required nuance: A valid token and hidden UI control prove neither
role
nor ownership permission; wrong-owner access commonly conceals
existence with the same 404 as absence.
Related IDs: BE-06-006, BE-06-007
Destination: C11, active entry 9
Status: merged into BE-06-001

### BE-06-003
Source: `foundations/06-identity-and-api-security.md`
Source section: `Participants and token purpose` (lines 130-164)
Category: mechanism
Idea: Resource owner, client, authorization server, and resource
server use access and ID tokens for different audiences and purposes.
Required nuance: This API consumes identity rather than implementing
passwords or issuance; token format resemblance never overrides
purpose and audience.
Related IDs: BE-06-004, BE-06-008
Destination: C11, active entry 9
Status: merged into BE-06-001

### BE-06-004
Source: `foundations/06-identity-and-api-security.md`
Source section: `What access-token validation must prove`
(lines 165-223)
Category: invariant
Idea: JWT acceptance is the conjunction of bounded parsing, allowed
algorithm, trusted key, signature, issuer, audience, time, subject,
and application claim checks.
Required nuance: Protected headers remain attacker-controlled until
verification, payload claims cannot grant authority before it, and a
maintained library plus pinned dependency owns cryptographic details.
Related IDs: BE-06-003, BE-06-005
Destination: C11, active entry 9
Status: merged into BE-06-001

### BE-06-005
Source: `foundations/06-identity-and-api-security.md`
Source section: `Discovery and JWKS rotation` (lines 224-272)
Category: operational control
Idea: A verifier discovers and caches bounded trusted keys while
failing closed through rotation and temporary issuer failure.
Required nuance: Exact issuer and TLS, expiry refresh, one
rate-limited
unknown-key refresh, valid stale cache during brief outage, response
and concurrency limits, and overlap testing prevent per-request fetch
and verification bypass.
Related IDs: BE-06-004, BE-06-011
Destination: C11, active entry 9
Status: merged into BE-06-001

### BE-06-006
Source: `foundations/06-identity-and-api-security.md`
Source section: `Build a trusted principal` (lines 273-347)
Category: mechanism
Idea: Reduce verified provider claims once into a small application
principal carried with a typed request-context key.
Required nuance: Missing or invalid credentials are 401, inability to
decide safely is a logged generic 503, raw claims do not spread, and
tokens or authorization headers are never logged.
Related IDs: BE-03-008, BE-06-002
Destination: C11, active entry 9
Status: merged into BE-06-001

### BE-06-007
Source: `foundations/06-identity-and-api-security.md`
Source section: `Role authorization belongs near the action`
(lines 348-377)
Category: operational control
Idea: Application operations enforce the capability they require so a
future non-HTTP adapter cannot bypass handler-only checks.
Required nuance: Distinguish absent identity from known-but-forbidden;
centralize narrow permissions and avoid scattered universal admin
bypasses.
Related IDs: BE-06-002, BE-10-004
Destination: C11, active entry 9
Status: merged into BE-06-001

### BE-06-008
Source: `foundations/06-identity-and-api-security.md`
Source section: `Object-level authorization belongs in data access`
(lines 378-422)
Category: operational control
Idea: Use the verified subject in tenant/customer-scoped SQL rather
than accepting owner identity from request input or filtering later.
Required nuance: Query and list predicates prevent insecure direct
object reference and avoid leaking data to logs, caches, metrics, or
early returns before a later check.
Related IDs: BE-05-010, BE-14-001
Destination: C11, active entry 9
Status: merged into BE-06-001

### BE-06-009
Source: `foundations/06-identity-and-api-security.md`
Source section: `JWTs and opaque tokens have different tradeoffs`
(lines 423-458)
Category: tradeoff
Idea: JWTs exchange local validation and issuer-outage tolerance for
revocation delay, while opaque tokens centralize current state with a
latency and availability dependency.
Required nuance: Payloads are encoded rather than encrypted; bound
lifetimes and avoid sensitive claims; choose from threat and provider
needs instead of punctuation in the token.
Related IDs: BE-06-003, BE-06-005
Destination: C11, active entry 9
Status: merged into BE-06-001

### BE-06-010
Source: `foundations/06-identity-and-api-security.md`
Source section: `Browser sessions, CORS, and CSRF` (lines 459-501)
Category: operational control
Idea: CORS governs browser cross-origin access, CSRF abuses ambient
credentials, and XSS executes within the application's authority.
Required nuance: Exact origin policy, secure cookie attributes,
SameSite, CSRF tokens or origin checks, bounded sessions, and explicit
logout/revocation solve different problems and need separate tests.
Related IDs: BE-00-041, BE-17-006
Destination: C11, active entry 9
Status: merged into BE-06-001

### BE-06-011
Source: `foundations/06-identity-and-api-security.md`
Source section: `Prevent common API data exposures` (lines 502-581)
Category: operational control
Idea: Allowlisted DTOs, parameterized queries, redacted telemetry,
least privilege, secret rotation, and resource-specific abuse limits
reduce common API exposure.
Required nuance: Caller fields never set ownership, role, price, or
audit state; runtime and migrator database roles differ; rate,
concurrency, body, and quota limits protect distinct resources.
Related IDs: BE-03-006, BE-04-008
Destination: C11, active entry 9
Status: merged into BE-06-001

### BE-06-012
Source: `foundations/06-identity-and-api-security.md`
Source section: `Security test matrix` (lines 582-635)
Category: test or experiment
Idea: Deterministic negative tests cover every token validation claim,
role decision, object decision, key rotation, and verifier outage.
Required nuance: Wrong owner equals nonexistent externally, cached
keys survive bounded refresh failure, unknown keys fail closed, and
captured logs contain no token or authorization header.
Related IDs: BE-06-004 through BE-06-011
Destination: C11, active entry 9
Status: merged into BE-06-001

### BE-06-013
Source: `foundations/06-identity-and-api-security.md`
Source section: `Compact threat model` (lines 636-679)
Category: procedure
Idea: Threat modeling connects assets, attacker capability, entry and
trust boundaries, impact, prevention, detection, recovery, and proof.
Required nuance: Owner-scoped data, role checks, log redaction,
inventory atomicity, key failure, and database privilege each produce
a concrete control and evidence case.
Related IDs: BE-14-002, BE-15-001
Destination: C11, active entry 9
Status: merged into BE-06-001

### BE-06-014
Source: `foundations/06-identity-and-api-security.md`
Source section: `Guided build` (lines 680-702)
Category: procedure
Idea: Add resource-server security from trusted issuer configuration
through verification, principal propagation, use-case authorization,
owner-scoped SQL, negative tests, redaction, and threat modeling.
Required nuance: Do not expand into password login or token issuance;
those are separate products and boundaries.
Related IDs: BE-06-001 through BE-06-013
Destination: C11, active entry 9
Status: merged into BE-06-001

### BE-06-015
Source: `foundations/06-identity-and-api-security.md`
Source section: `Knowledge check` and `Completion evidence` (lines
703-759)
Category: completion evidence
Idea: Security completion requires full claim validation, safe key
rotation, action and object enforcement, a negative matrix, clean
telemetry, threat-to-evidence traceability, and security tool checks.
Required nuance: Every accepted identity and denied path has a named
reason; no unverifiable token or another user's order is exposed.
Related IDs: BE-06-001 through BE-06-014
Destination: C11, active entry 9
Status: merged into BE-06-001

### BE-07-001
Source: `foundations/07-testing-observability-and-operations.md`
Source section: introduction through `Evidence for one failed order`
(lines 3-70)
Category: principle
Idea: Tests, telemetry, alerts, and runbooks are different evidence
for claims before and during service failure.
Required nuance: One pool-saturation incident demonstrates how a test,
log, metric, trace, alert, runbook, SLI, and SLO answer separate but
connected questions.
Related IDs: BE-00-010, BE-07-005
Destination: C12, active entry 10
Status: incorporated

### BE-07-002
Source: `foundations/07-testing-observability-and-operations.md`
Source section: `Tests answer different questions` (lines 71-113)
Category: test or experiment
Idea: Choose the cheapest boundary that reproduces the risk, then use
a small number of broad workflow tests for confidence.
Required nuance: Unit, handler, integration, and acceptance tests own
different mechanisms; mocks cannot prove PostgreSQL and broad tests
alone are slow and hard to diagnose.
Related IDs: BE-05-011, BE-04-011
Destination: C12, active entry 10
Status: merged into BE-07-001

### BE-07-003
Source: `foundations/07-testing-observability-and-operations.md`
Source section: `Table-driven tests make cases visible` and `Prefer
small fakes` (lines 114-199)
Category: test or experiment
Idea: Named case tables expose one behavior across inputs, while small
fakes record only the consumer interaction under test.
Required nuance: Assert stable results and identities; parallelize
only
isolated fixtures; do not invent database semantics or production
interfaces merely to satisfy a mocking framework.
Related IDs: BE-01-014, BE-05-006
Destination: C12, active entry 10
Status: merged into BE-07-001

### BE-07-004
Source: `foundations/07-testing-observability-and-operations.md`
Source section: `Control time, identity, and randomness`
(lines 200-255)
Category: test or experiment
Idea: Inject deterministic clocks, IDs, fixtures, barriers, and
cancellation so assertions and interleavings do not depend on sleeps.
Required nuance: Concurrency tests can be deterministic by signaling
exact transitions; every wait has a deadline and fixtures contain only
the data needed for the case.
Related IDs: BE-05-007, BE-12-003
Destination: C12, active entry 10
Status: merged into BE-07-001

### BE-07-005
Source: `foundations/07-testing-observability-and-operations.md`
Source section: `Logs, metrics, and traces answer different questions`
(lines 256-351)
Category: mechanism
Idea: Logs record discrete context, metrics aggregate bounded trends,
traces connect one causal path, and audits record sensitive actions.
Required nuance: Route templates and stable dimensions avoid private
or high-cardinality labels; unexpected errors are logged once;
request and trace identity have related but distinct protocols.
Related IDs: BE-03-008, BE-11-002
Destination: C13, active entry 11
Status: incorporated

### BE-07-006
Source: `foundations/07-testing-observability-and-operations.md`
Source section: `Instrument at stable boundaries` (lines 352-372)
Category: operational control
Idea: Instrument HTTP, adapter, pool, and durable business boundaries
that can answer named operator questions.
Required nuance: Pure domain rules avoid SDK coupling; every field has
a bounded set, privacy review, user, retention purpose, and justified
cost or it is omitted.
Related IDs: BE-07-005, BE-13-005
Destination: C13, active entry 11
Status: merged into BE-07-005

### BE-07-007
Source: `foundations/07-testing-observability-and-operations.md`
Source section: `Service indicators and objectives` (lines 373-418)
Category: operational control
Idea: Define a user-visible SLI, its target SLO, eligibility rules,
and
allowed error budget before choosing alert thresholds.
Required nuance: Latency uses distributions, invalid requests can be
excluded but stay visible, ambiguous disconnects and missing telemetry
need denominator policy, and burn alerts reflect traffic volume.
Related IDs: BE-14-005, BE-16-009
Destination: C13, active entry 11
Status: merged into BE-07-005

### BE-07-008
Source: `foundations/07-testing-observability-and-operations.md`
Source section: `Health and graceful shutdown are protocols`
(lines 419-457)
Category: operational control
Idea: Health and shutdown coordinate traffic withdrawal, draining,
worker cancellation, dependency closure, and bounded forced exit.
Required nuance: Liveness remains healthy during dependency outage;
readiness fails; shutdown tests both completed drain and deadline
exhaustion because HTTP shutdown alone does not cancel active work.
Related IDs: BE-03-009, BE-08-009
Destination: C13, active entry 11
Status: merged into BE-07-005

### BE-07-009
Source: `foundations/07-testing-observability-and-operations.md`
Source section: `Bound concurrency and dependencies` (lines 458-491)
Category: operational control
Idea: Timeouts bound duration while concurrency and queue limits bound
simultaneous resource consumers.
Required nuance: Explicit rejection and backpressure break the
overload
feedback loop; size limits from dependency capacity and observed
latency rather than runtime ability to spawn work.
Related IDs: BE-04-007, BE-11-004
Destination: C16, active entry 14
Status: merged into BE-11-004

### BE-07-010
Source: `foundations/07-testing-observability-and-operations.md`
Source section: `Retry only when all conditions are satisfied`
(lines 492-548)
Category: operational control
Idea: Retry only transient, safe or idempotency-protected operations
inside one total deadline with bounded attempts, jitter, and one
owner.
Required nuance: Ambiguous writes need identity, business and auth
outcomes never retry, remaining time can stop an allowed numeric
attempt, and layered three-attempt policies can create 27 calls.
Related IDs: BE-11-001, BE-13-002
Destination: C16, active entry 14
Status: merged into BE-11-004

### BE-07-011
Source: `foundations/07-testing-observability-and-operations.md`
Source section: `Failure drills` (lines 549-604)
Category: test or experiment
Idea: Deliberately inject validation, internal, database, and canceled
request failures and trace each through public result to recovery.
Required nuance: Each class has different expected database work,
severity, health, pool, trace, and log behavior; evidence ends with
durable-state verification rather than signal appearance.
Related IDs: BE-08-010, BE-13-005
Destination: C13, active entry 11
Status: merged into BE-07-005

### BE-07-012
Source: `foundations/07-testing-observability-and-operations.md`
Source section: `Alerts and runbooks` (lines 605-648)
Category: operational control
Idea: An actionable alert represents sustained impact or imminent
exhaustion and leads to tested diagnosis, mitigation, escalation, and
recovery criteria.
Required nuance: Internal saturation alone may be context rather than
a page; destructive commands stay out; every action states safety
conditions, permissions, rollback, and how success is verified.
Related IDs: BE-14-005, BE-08-010
Destination: C13, active entry 11
Status: merged into BE-07-005

### BE-07-013
Source: `foundations/07-testing-observability-and-operations.md`
Source section: `Guided build` (lines 649-667)
Category: procedure
Idea: Align test boundaries, deterministic dependencies, telemetry,
shutdown, drills, objectives, alerts, and runbook around one service.
Required nuance: OpenTelemetry stays at process and external
boundaries while vendor/exporter configuration stays out of business
packages.
Related IDs: BE-07-001 through BE-07-012
Destination: C12, active entry 10
Status: merged into BE-07-001

### BE-07-014
Source: `foundations/07-testing-observability-and-operations.md`
Source section: `Knowledge check` and `Completion evidence` (lines
668-725)
Category: completion evidence
Idea: Completion requires deterministic risk-owned tests, distinct
safe telemetry, correct lifecycle behavior, bounded retries, one
actionable objective alert, and all code and security checks.
Required nuance: The evidence must distinguish invalid input, internal
defect, slow database, cancellation, and saturation without arbitrary
time or sensitive fields.
Related IDs: BE-07-001 through BE-07-013
Destination: C12, active entry 10
Status: merged into BE-07-001

### BE-08-001
Source: `foundations/08-delivery-and-production-readiness.md`
Source section: introduction through `Delivery continues after
compilation` (lines 3-70)
Category: principle
Idea: Delivery is a continuous evidence chain from source to artifact,
compatible rollout, health observation, and recovery rehearsal.
Required nuance: Artifact, CI, rolling deployment, rollback,
roll-forward, RPO, and RTO name distinct responsibilities; a local
success leaves configuration, stop, compatibility, and restore open.
Related IDs: BE-00-011, BE-08-006
Destination: C14, active entry 12
Status: incorporated

### BE-08-002
Source: `foundations/08-delivery-and-production-readiness.md`
Source section: `One immutable artifact, many environments`
(lines 71-91)
Category: operational control
Idea: Build once and promote identical bytes while runtime
configuration supplies environment-specific behavior.
Required nuance: Record revision, toolchain, build time, and digest;
rebuilding a commit can still change bytes through time, dependencies,
base images, or settings.
Related IDs: BE-11-003, BE-14-002
Destination: C14, active entry 12
Status: merged into BE-08-001

### BE-08-003
Source: `foundations/08-delivery-and-production-readiness.md`
Source section: `Build a small non-root container` (lines 92-168)
Category: operational control
Idea: A pinned multi-stage image contains only the runtime binary and
required data, runs non-root, and writes only to explicit paths.
Required nuance: CA, time-zone, and CGO needs are tested; source, Git,
toolchain, and credentials are absent; read-only operation and actual
image contents are verified rather than inferred from Dockerfile.
Related IDs: BE-08-002, BE-14-004
Destination: C14, active entry 12
Status: merged into BE-08-001

### BE-08-004
Source: `foundations/08-delivery-and-production-readiness.md`
Source section: `Configuration is a startup contract` (lines 169-217)
Category: operational control
Idea: Parse environment input once into typed configuration and reject
missing, malformed, contradictory, or unsafe state before listening.
Required nuance: Separate build input, runtime non-secret, and runtime
secret delivery; errors name fields without values; production never
falls back silently to development credentials.
Related IDs: BE-10-001, BE-14-003
Destination: C14, active entry 12
Status: merged into BE-08-001

### BE-08-005
Source: `foundations/08-delivery-and-production-readiness.md`
Source section: `CI reproduces the local quality contract`
(lines 218-264)
Category: completion evidence
Idea: Repository-owned commands define one local and CI pipeline for
format, analysis, tests, migrations, security, build, and smoke proof.
Required nuance: Pin automation and tools, grant least privilege,
protect untrusted PR secrets and caches, publish only passing
artifacts, and treat warnings as failures.
Related IDs: BE-00-014, BE-17-008
Destination: C14, active entry 12
Status: merged into BE-08-001

### BE-08-006
Source: `foundations/08-delivery-and-production-readiness.md`
Source section: `Application and schema deployment are not atomic`
(lines 265-307)
Category: procedure
Idea: Expand-and-contract lets old and new application versions share
one schema through additive shape, compatible code, bounded backfill,
use switch, stopped old writes, and delayed removal.
Required nuance: Evaluate version-specific lock and rewrite behavior;
migrations are one controlled step; contraction waits past rollback
and drain windows.
Related IDs: BE-04-005, BE-13-004
Destination: C14, active entry 12
Status: merged into BE-08-001

### BE-08-007
Source: `foundations/08-delivery-and-production-readiness.md`
Source section: `Rollback and roll-forward protect different failures`
(lines 308-356)
Category: tradeoff
Idea: Roll back only when old code understands current schema and
data; otherwise preserve evidence and repair forward or restore.
Required nuance: Decide artifact, compatibility window, stop signals,
owners, queries, and user verification before rollout; configuration,
binary, schema, and data failures have different recovery directions.
Related IDs: BE-08-006, BE-15-002
Destination: C14, active entry 12
Status: merged into BE-08-001

### BE-08-008
Source: `foundations/08-delivery-and-production-readiness.md`
Source section: `Backups become trustworthy through restore drills`
(lines 357-387)
Category: test or experiment
Idea: An isolated restore must recreate usable, invariant-preserving
application data within measured RPO and RTO.
Required nuance: Backup success proves only stored bytes; drills check
schema, counts, samples, constraints, application startup, encryption,
retention, and never overwrite the only production database.
Related IDs: BE-11-005, BE-14-005
Destination: C14, active entry 12
Status: merged into BE-08-001

### BE-08-009
Source: `foundations/08-delivery-and-production-readiness.md`
Source section: `Capacity is a chain of bounded resources`
(lines 388-430)
Category: operational control
Idea: Safe throughput is limited by the first saturated resource
across edge, process, memory, CPU, pools, locks, telemetry, identity,
and downstream quotas.
Required nuance: Bound queues and shed impossible work; calculate
fleet
budgets and safety reserve; representative load tests report tail,
errors, throughput, saturation, bottleneck, and beyond-capacity
behavior.
Related IDs: BE-07-009, BE-11-004
Destination: C16, active entry 14
Status: merged into BE-11-004

### BE-08-010
Source: `foundations/08-delivery-and-production-readiness.md`
Source section: `Rehearse startup and termination failures`
(lines 431-479)
Category: test or experiment
Idea: Exercise missing configuration, clean startup, readiness loss
and recovery, graceful drain and exhaustion, and clean disposable
database reproduction with the built artifact.
Required nuance: Recovery includes a user workflow and durable-data
check; liveness stays healthy on database loss; destructive recreation
targets only explicit disposable data.
Related IDs: BE-07-008, BE-11-003
Destination: C14, active entry 12
Status: merged into BE-08-001

### BE-08-011
Source: `foundations/08-delivery-and-production-readiness.md`
Source section: `Conduct a 30-minute incident drill` (lines 480-551)
Category: test or experiment
Idea: A time-boxed database incident separates coordination,
investigation, operation, and recording from detection through
recovery.
Required nuance: Detect from signals rather than sabotage, trace one
request, verify health semantics and integrity, stop at the time box,
and turn specific evidence into owned, verifiable corrective action.
Related IDs: BE-07-011, BE-14-005
Destination: C14, active entry 12
Status: merged into BE-08-001

### BE-08-012
Source: `foundations/08-delivery-and-production-readiness.md`
Source section: `Make the repository sufficient for another developer`
(lines 552-577)
Category: completion evidence
Idea: A clean-clone test turns every hidden tool, config, database,
migration, run, check, health, failure, and cleanup step into owned
documentation or repository automation.
Required nuance: Another developer follows only the docs; each hidden
assumption updates prerequisites or setup, then the test restarts.
Related IDs: BE-08-005, BE-15-003
Destination: C14, active entry 12
Status: merged into BE-08-001

### BE-08-013
Source: `foundations/08-delivery-and-production-readiness.md`
Source section: `Guided delivery build` and optional deployment
(lines 578-610)
Category: procedure
Idea: Build container, typed config, shared commands, CI, empty-schema
proof, runtime checks, compatibility decisions, restore, incident, and
external handoff sequentially.
Required nuance: A non-production deployment uses the simplest fitting
platform and records artifact, schema, timings, and exercised
recovery;
Kubernetes is not a completion prop.
Related IDs: BE-08-001 through BE-08-012
Destination: C14, active entry 12
Status: merged into BE-08-001

### BE-08-014
Source: `foundations/08-delivery-and-production-readiness.md`
Source section: `Knowledge check` and `Completion evidence` (lines
611-673)
Category: completion evidence
Idea: Production readiness requires minimal artifact proof, startup
validation, shared local/CI checks, clean migrations, lifecycle tests,
compatible recovery, restore truth, an incident timeline, and handoff.
Required nuance: Claims without exercised paths are marked not yet
established rather than promoted as confidence.
Related IDs: BE-08-001 through BE-08-013
Destination: C14, active entry 12
Status: merged into BE-08-001

### BE-09-001
Source: `advanced/09-domain-discovery.md`
Source section: introduction and `Start with one sentence` (lines
3-41)
Category: procedure
Idea: Turn vague behavior into one concrete case and ask what is true,
changes, becomes impossible, or races before designing technology.
Required nuance: Different coherent interpretations are product
ambiguity; unanswered questions become explicit assumptions or open
decisions rather than silent contracts.
Related IDs: BE-00-013, BE-09-002
Destination: C06, active entry 4
Status: incorporated

### BE-09-002
Source: `advanced/09-domain-discovery.md`
Source section: `Treat requirements as hypotheses` (lines 42-74)
Category: principle
Idea: Classify requirement information as confirmed fact, assumption,
open question, or deferred scope until examples and owners resolve it.
Required nuance: Cancellation exposes state, inventory, actor, repeat,
and clock decisions; working beliefs never silently harden into API or
schema contracts.
Related IDs: BE-09-001, BE-15-002
Destination: C06, active entry 4
Status: merged into BE-09-001

### BE-09-003
Source: `advanced/09-domain-discovery.md`
Source section: `Build ubiquitous language` (lines 75-102)
Category: principle
Idea: A ubiquitous language gives important domain terms one reviewed
meaning across product, code, contracts, events, dashboards, and
runbooks.
Required nuance: Disputed meanings stay recorded; different bounded
contexts may intentionally define a term such as Product differently.
Related IDs: BE-09-007, BE-16-002
Destination: C06, active entry 4
Status: merged into BE-09-001

### BE-09-004
Source: `advanced/09-domain-discovery.md`
Source section: `Use examples to discover rules` (lines 103-131)
Category: test or experiment
Idea: Concrete success, rejection, concurrency, retry, ownership,
cancellation, and dependency examples reveal decisions and invariants.
Required nuance: State exact values and outcomes, including captured
history and exactly-one contention, rather than describing only the
happy path.
Related IDs: BE-05-005, BE-09-006
Destination: C06, active entry 4
Status: merged into BE-09-001

### BE-09-005
Source: `advanced/09-domain-discovery.md`
Source section: `Commands, events, entities, and value objects`
(lines 132-148)
Category: principle
Idea: Commands request fallible change, past-tense events record
completed facts, entities retain identity, and value objects are
defined by equal values.
Required nuance: Identity does not justify arbitrary mutation; expose
operations that preserve invariants instead of public setters.
Related IDs: BE-09-006, BE-12-002
Destination: C06, active entry 4
Status: merged into BE-09-001

### BE-09-006
Source: `advanced/09-domain-discovery.md`
Source section: `State machines expose missing decisions`
(lines 149-176)
Category: mechanism
Idea: A state machine names allowed, terminal, invalid, repeated, and
failure transitions with their actor, preconditions, effects, and
event.
Required nuance: A technically possible transition needs business
meaning and an owner; repeat, timeout, inventory, and terminal
behavior
are part of the transition contract.
Related IDs: BE-05-005, BE-13-003
Destination: C06, active entry 4
Status: merged into BE-09-001

### BE-09-007
Source: `advanced/09-domain-discovery.md`
Source section: `Write invariants independently of technology`
(lines 177-195)
Category: invariant
Idea: State durable business truths without HTTP, Go, SQL, or broker
terminology so every implementation boundary can preserve them.
Required nuance: Inventory, ownership, visibility, positive quantity,
captured totals, atomic placement, logical retry, and terminal-state
rules are the initial Order API invariants.
Related IDs: BE-04-001, BE-13-003
Destination: C06, active entry 4
Status: merged into BE-09-001

### BE-09-008
Source: `advanced/09-domain-discovery.md`
Source section: `Map bounded contexts and information flow`
(lines 196-225)
Category: principle
Idea: Catalog, inventory, ordering, identity, and notification own
coherent models and information flows without automatically becoming
services.
Required nuance: State both what each context owns and does not own;
deploy separately only for evidenced ownership, scale, isolation, or
release value.
Related IDs: BE-09-003, BE-13-001
Destination: C06, active entry 4
Status: merged into BE-09-001

### BE-09-009
Source: `advanced/09-domain-discovery.md`
Source section: `Guided work` and `Failure checks` (lines 226-256)
Category: test or experiment
Idea: A reviewed discovery document joins actors, glossary, examples,
invariants, state machines, context flow, decision classes, and a
service counterfactual.
Required nuance: Give the same phrase to independent reviewers, remove
technical nouns, repeat commands, and race valid commands to expose
remaining ambiguity.
Related IDs: BE-09-001 through BE-09-008
Destination: C06, active entry 4
Status: merged into BE-09-001

### BE-09-010
Source: `advanced/09-domain-discovery.md`
Source section: `Knowledge check` and `Completion evidence` (lines
257-289)
Category: completion evidence
Idea: Discovery is complete when every use case has reviewed examples
and invariants, conflicts are resolved, and another reviewer explains
states without implementation vocabulary.
Required nuance: Facts and assumptions stay separate and bounded
contexts remain model ownership rather than assumed deployments.
Related IDs: BE-09-001 through BE-09-009
Destination: C06, active entry 4
Status: merged into BE-09-001

### BE-10-001
Source: `advanced/10-service-core.md`
Source section: `Decisions and Executable Skeleton` (lines 6-114)
Category: procedure
Idea: An executable skeleton proves typed startup configuration,
composition root, bounded server lifecycle, shared checks, and ADRs
before feature scaffolding grows.
Required nuance: Missing config fails before listening without
secrets;
`main` owns construction and exit; SIGTERM drains; local and CI
failures match; business packages do not discover globals.
Related IDs: BE-00-014, BE-08-004
Destination: C07, active entry 5
Status: merged into BE-05-001

### BE-10-002
Source: `advanced/10-service-core.md`
Source section: `HTTP Platform` (lines 115-212)
Category: mechanism
Idea: One bounded HTTP boundary owns contracts, strict decoding,
problem responses, middleware, lifecycle health, and explicit
timeouts.
Required nuance: Build the negative unknown-field path first, keep
handlers thin, log template routes, document implemented behavior in
OpenAPI, and distinguish draining readiness from process liveness.
Related IDs: BE-03-001 through BE-03-013
Destination: C08, active entry 6
Status: merged into BE-03-001

### BE-10-003
Source: `advanced/10-service-core.md`
Source section: `Persistence Foundation` (lines 213-305)
Category: invariant
Idea: PostgreSQL adapters combine bounded pool capacity, ordered
migrations, constraints, contextual parameterized queries, cleanup,
and real-database tests.
Required nuance: Two concurrent duplicate inserts prove why the
database is final authority; errors translate only when known;
constraints and empty-schema migration tests remain reproducible.
Related IDs: BE-04-001 through BE-04-014
Destination: C09, active entry 7
Status: merged into BE-04-001

### BE-10-004
Source: `advanced/10-service-core.md`
Source section: `Product Slice` (lines 306-394)
Category: procedure
Idea: A vertical product capability joins one use case, strict HTTP,
PostgreSQL, stable pagination, role authorization, contract, and
tests.
Required nuance: Duplicate identity is stable conflict, owner names
the narrow interface, cursor ordering is total, and admin authority
must be enforced beyond a hidden frontend control.
Related IDs: BE-00-017, BE-05-004
Destination: C07, active entry 5
Status: merged into BE-05-001

### BE-10-005
Source: `advanced/10-service-core.md`
Source section: `Order Transaction` (lines 395-496)
Category: invariant
Idea: PlaceOrder validates input, then one PostgreSQL transaction
protects inventory, captured prices, order, items, and owner-scoped
retrieval.
Required nuance: Atomic conditional updates and sorted lock order
handle contention; commit once; errors remain stable and private;
rollback, last-unit, and wrong-owner integration tests prove behavior.
Related IDs: BE-04-009, BE-05-008
Destination: C10, active entry 8
Status: merged into BE-04-009

### BE-11-001
Source: `advanced/11-production-service.md`
Source section: `Idempotency and Failure Semantics` (lines 6-118)
Category: invariant
Idea: Order creation binds a principal-scoped key to one canonical
command, serializes duplicates, and persists its compatible outcome
with the effect.
Required nuance: Raw JSON is not canonical; changed payload conflicts;
in-progress and expired behavior are public; crash points, disconnect,
retention, cleanup, and record telemetry need explicit tests.
Related IDs: BE-00-019, BE-07-010
Destination: C15, active entry 13
Status: incorporated

### BE-11-002
Source: `advanced/11-production-service.md`
Source section: `Observability and Operational Controls`
(lines 119-223)
Category: operational control
Idea: Boundary telemetry, durable audit evidence, user-first
dashboards, objectives, alerts, and runbooks explain service outcomes.
Required nuance: Logs, metrics, traces, and audits differ; labels and
attributes are bounded and redacted; one success and failure connect
public result through resource cause without debug code.
Related IDs: BE-07-005 through BE-07-012
Destination: C13, active entry 11
Status: merged into BE-07-005

### BE-11-003
Source: `advanced/11-production-service.md`
Source section: `Deployment Hardening` (lines 224-317)
Category: operational control
Idea: Promote a traceable least-privilege artifact through resource,
health, secret, schema, scanning, and recovery contracts.
Required nuance: Run non-root and read-only, separate migration
privilege, align limits, rehearse old/new schema states, scan source
and image, and restore data with application checks.
Related IDs: BE-08-002 through BE-08-010
Destination: C14, active entry 12
Status: merged into BE-08-001

### BE-11-004
Source: `advanced/11-production-service.md`
Source section: `Performance and Resilience Based on Evidence`
(lines 318-420)
Category: test or experiment
Idea: Reproduce a representative workload, profile all plausible
resources, change one measured bottleneck, and repeat the same test.
Required nuance: Avoid coordinated omission, report percentiles,
errors and saturation, test overload and recovery, and add a cache
only
with truth, key, TTL, invalidation, stampede, outage, and removal
rules.
Related IDs: BE-00-022, BE-08-009
Destination: C16, active entry 14
Status: incorporated

### BE-11-005
Source: `advanced/11-production-service.md`
Source section: `Advanced PostgreSQL Operations and Recovery`
(lines 421-534)
Category: operational control
Idea: MVCC, isolation, locks, vacuum, statistics, bloat, checkpoints,
WAL, replicas, backfills, backups, and recovery affect one live
system.
Required nuance: Simulate anomalies and deadlocks, retry whole safe
operations, resume idempotent batches, diagnose with database views,
perform PITR and standby promotion, and measure RPO, RTO, and lag.
Related IDs: BE-00-023, BE-04-009
Destination: C17, active entry 15
Status: incorporated

### BE-12-001
Source: `advanced/12-distributed-foundations.md`
Source section: `Durable Background Jobs and Scheduling`
(lines 7-126)
Category: mechanism
Idea: Persist job intent, claim bounded batches with expiring leases
and new claim tokens, perform idempotent effects, then conditionally
acknowledge the current claim.
Required nuance: Lease expiry does not stop old workers; classify
success, transient, permanent, and exhausted results; bound payload,
attempt, concurrency, storage, replay, dead letter, scheduler, age,
backpressure, and shutdown.
Related IDs: BE-00-024, BE-13-003
Destination: C18, active entry 16
Status: incorporated

### BE-12-002
Source: `advanced/12-distributed-foundations.md`
Source section: `Reliable Event Publication and Consumption`
(lines 127-245)
Category: mechanism
Idea: Commit business state and outbox intent locally, relay with
possible duplicates, and commit inbox dedupe with each consumer
effect.
Required nuance: Reconciliation repairs weaker derived intent;
ordering is scoped, contracts version semantics, replay respects
dedupe, retention is explicit, and exactly once names a boundary
rather
than every external effect.
Related IDs: BE-00-025, BE-13-003
Destination: C19, active entry 17
Status: incorporated

### BE-12-003
Source: `advanced/12-distributed-foundations.md`
Source section: `Distributed-Systems Mechanics Lab` (lines 246-378)
Category: test or experiment
Idea: A logical simulator lets tests choose message delay, loss,
duplication, reorder, partition, lease, quorum, hashing, and leader
events without sleeps.
Required nuance: Timeouts prove only stopped waiting; monotonic clocks
measure local duration; fencing rejects stale authority; quorum
overlap depends on assumptions; hashing is not replication; production
consensus belongs to maintained systems.
Related IDs: BE-00-026, BE-07-004
Destination: C20, active entry 18
Status: incorporated

### BE-13-001
Source: `advanced/13-service-extraction.md`
Source section: `Service-Boundary Discovery and First Extraction`
(lines 7-110)
Category: tradeoff
Idea: Extract a coherent business capability only when change,
ownership, scaling, isolation, or release evidence exceeds distributed
operating cost.
Required nuance: Services own contract, data, runtime, deployment, and
on-call load; no cross-schema access; branch by abstraction and staged
traffic preserve rollback; notification failure must not block orders.
Related IDs: BE-05-014, BE-09-008
Destination: C21, active entry 19
Status: incorporated

### BE-13-002
Source: `advanced/13-service-extraction.md`
Source section: `Synchronous Service Communication and the System
Edge` (lines 111-252)
Category: mechanism
Idea: A remote inventory call replaces local atomicity with a
versioned
contract, operation identity, pending intent, total deadline, and
reconciliation.
Required nuance: HTTP versus gRPC is a tooling tradeoff; never hold a
local transaction across the call; query or release ambiguous
reservations; authenticate workload identity; separate gateway,
retry, concurrency, bulkhead, breaker, shedding, and fallback roles.
Related IDs: BE-00-028, BE-07-010
Destination: C22, active entry 20
Status: incorporated

### BE-13-003
Source: `advanced/13-service-extraction.md`
Source section: `Distributed Workflows and Eventual Consistency`
(lines 253-365)
Category: mechanism
Idea: A saga persists versioned workflow state and next action across
local commits, expected intermediate states, timeouts, and
compensation.
Required nuance: Compensation is a new fallible idempotent command;
orchestration and choreography have different visibility; every async
edge uses outbox/inbox; late and out-of-order events obey guarded
state;
repair is scoped, previewed, audited, and verified.
Related IDs: BE-09-006, BE-12-002
Destination: C23, active entry 21
Status: incorporated

### BE-13-004
Source: `advanced/13-service-extraction.md`
Source section: `Contract Evolution and Independent Delivery`
(lines 366-460)
Category: tradeoff
Idea: Structural and semantic compatibility must support old/new
provider and consumer deployment in both orders.
Required nuance: Additive enum or validation changes can break;
OpenAPI,
Protobuf, and event schemas need examples and owners; consumer
contracts complement other tests; deprecation requires inventory,
telemetry, window, communication, criteria, and adapter cleanup.
Related IDs: BE-00-030, BE-08-006
Destination: C24, active entry 22
Status: incorporated

### BE-13-005
Source: `advanced/13-service-extraction.md`
Source section: `Distributed Observability and Resilience Validation`
(lines 461-566)
Category: test or experiment
Idea: Propagate safe correlation across synchronous and asynchronous
work, then run bounded hypothesis-driven failure experiments.
Required nuance: Allowlist baggage, bound sampling infrastructure,
retain slow and error traces, observe RED signals and queue age,
define
blast radius and abort, expose retry amplification, and verify durable
convergence after recovery.
Related IDs: BE-00-031, BE-07-011
Destination: C25, active entry 23
Status: incorporated

### BE-14-001
Source: `advanced/14-scale-security-cloud-operations.md`
Source section: `Data Scaling, Caching, and Multi-Tenancy`
(lines 7-135)
Category: tradeoff
Idea: Scale a measured bottleneck with a cache, replica, partition,
archive, or tenant model whose consistency and operating costs are
explicit.
Required nuance: Cache-aside needs versioned tenant keys, TTL,
invalidation, outage, bounds, and stampede control; replicas expose
lag; partitions need plan and lifecycle proof; tenant identity comes
from trust and scopes every key, constraint, connection, async path,
operator path, quota, backup, and repair.
Related IDs: BE-00-032, BE-06-008
Destination: C26, active entry 24
Status: incorporated

### BE-14-002
Source: `advanced/14-scale-security-cloud-operations.md`
Source section: `Security, Privacy, and Software Supply Chain`
(lines 136-256)
Category: operational control
Idea: New boundaries require distinct user and workload identity,
least privilege, rotatable trust, data-lifecycle ownership, and
verifiable artifact lineage.
Required nuance: Encryption is not authorization; rotation overlaps
new and old credentials; deletion/export reconcile every primary,
cached, queued, derived, provider, audit, and backup policy; SBOM,
provenance, signatures, and advisory drills connect component to
patch.
Related IDs: BE-06-013, BE-08-002
Destination: C27, active entry 25
Status: incorporated

### BE-14-003
Source: `advanced/14-scale-security-cloud-operations.md`
Source section: `Cloud Infrastructure and IAM` (lines 257-364)
Category: operational control
Idea: Infrastructure as code reproduces cloud failure, network, human,
workload, resource, encryption, logging, cost, and teardown
boundaries.
Required nuance: Connectivity and authorization are independent;
only the edge is public; workloads use short-lived identities; IaC
state is protected and locked; plan, policy, drift, tags, budgets,
audit, and scoped teardown are reviewed and failure-tested.
Related IDs: BE-00-034, BE-14-002
Destination: C28, active entry 26
Status: incorporated

### BE-14-004
Source: `advanced/14-scale-security-cloud-operations.md`
Source section: `Orchestration and Platform Operations`
(lines 365-472)
Category: operational control
Idea: An orchestrator reconciles declared workload state but cannot
supply truthful application startup, readiness, liveness, drain, or
schema compatibility.
Required nuance: Requests schedule and limits cap; probes model three
different states; rollouts, disruption, network policy, termination,
and autoscaling need downstream budgets; platform abstractions expose
generated resources and failures to non-admin operators.
Related IDs: BE-03-009, BE-08-009
Destination: C29, active entry 27
Status: incorporated

### BE-14-005
Source: `advanced/14-scale-security-cloud-operations.md`
Source section: `Reliability Engineering, Incidents, and Cost`
(lines 473-574)
Category: operational control
Idea: User-centered SLIs, SLOs, error budgets, capacity, incident
roles, evidence, corrective action, and total cost guide reliability
tradeoffs.
Required nuance: Alerts begin from impact, capacity includes failure
domains and dependency ceilings, game days measure the full response,
reviews stay blameless, and cost includes people, operational
complexity, opportunity, headroom, and SLO risk.
Related IDs: BE-07-007, BE-08-011
Destination: C30, active entry 28
Status: incorporated

### BE-15-001
Source: `advanced/15-integrations-maintenance-leadership.md`
Source section: `Common Production Integration Boundaries`
(lines 7-136)
Category: mechanism
Idea: Each external provider sits behind a bounded adapter that turns
ambiguous, quota-bound protocol outcomes into application-owned
states.
Required nuance: Outbound effects use idempotency or lookup;
webhooks verify bounded raw bytes, time, signature, rotation, and
dedupe before durable acknowledgement; subscriber URLs need complete
SSRF defense; objects stream with scoped grants and scanning;
notification callbacks update later state; derived views rebuild and
reconcile from authority.
Related IDs: BE-00-037, BE-06-013
Destination: C31, active entry 29
Status: incorporated

### BE-15-002
Source: `advanced/15-integrations-maintenance-leadership.md`
Source section: `Maintenance and Legacy-System Evolution`
(lines 137-256)
Category: procedure
Idea: Discover actual relied-on behavior, pin it at the narrowest
boundary, and evolve through small reversible changes with evidence.
Required nuance: Code, tests, telemetry, data, docs, flags, incidents,
support, and users may disagree; upgrades have layer-specific risks;
flags and canaries require removal; repair is scoped, dry-run,
idempotent, resumable, rate-limited, authorized, audited, guarded,
killable, and independently verified; regression work fixes cause.
Related IDs: BE-00-038, BE-08-007
Destination: C32, active entry 30
Status: incorporated

### BE-15-003
Source: `advanced/15-integrations-maintenance-leadership.md`
Source section: `Capstone Delivery and Engineering Leadership`
(lines 257-399)
Category: completion evidence
Idea: Lead a valuable, risk-bearing feature from discovery and
alternatives through review, compatible increments, evidence,
operation, handoff, and simplification.
Required nuance: The design covers domain, contracts, data, failures,
security, telemetry, capacity, cost, rollout, cleanup, and residual
risk; trace every requirement to implementation, test, signal, and
operator action; rehearse relevant concurrency, dependency, duplicate,
restart, mixed-version, migration, overload, and recovery cases;
enable
another engineer and remove temporary complexity.
Related IDs: BE-00-039, BE-00-045
Destination: C33, active entry 31
Status: incorporated

### BE-16-001
Source: `advanced/16-key-concepts-review.md`
Source section: introduction and `How to use this refresher actively`
(lines 3-37)
Category: procedure
Idea: Use active recall by explaining one concrete success and
failure,
answering checks, and reopening the detailed lesson wherever reasoning
becomes vague.
Required nuance: Recognition is not understanding; replay ambiguous
timeout, operation identity, and reconciliation rather than repeating
a slogan.
Related IDs: BE-00-001, BE-16-029
Destination: C01, supporting introduction and review checklist
Status: merged into BE-00-001

### BE-16-002
Source: `advanced/16-key-concepts-review.md`
Source section: `How the topics fit together` (lines 38-82)
Category: principle
Idea: Domain, service, distributed, extraction, scale, integration,
and leadership topics form a prerequisite chain.
Required nuance: Distribution magnifies weak foundations; use the
service-boundary lens throughout as a counterfactual before evidence
justifies independently deployed infrastructure.
Related IDs: BE-00-012, BE-09-008
Destination: C01, supporting introduction and review checklist
Status: merged into BE-00-001

### BE-16-003
Source: `advanced/16-key-concepts-review.md`
Source section: `Domain discovery and modeling` (lines 83-116)
Category: principle
Idea: Examples, shared language, entities and values, commands and
events, invariants, states, decision classes, and contexts precede
implementation choices.
Required nuance: Explain valid and invalid behavior in domain language
and keep overloaded meanings and assumptions visible.
Related IDs: BE-09-001 through BE-09-010
Destination: C06, active entry 4
Status: merged into BE-09-001

### BE-16-004
Source: `advanced/16-key-concepts-review.md`
Source section: `Decisions and executable skeleton` (lines 117-148)
Category: principle
Idea: A composition root, typed startup, ADRs, bounded shutdown, and
one local/CI path establish an executable service foundation.
Required nuance: Startup establishes preconditions or fails loudly;
decisions remain revisitable through recorded alternatives and
triggers.
Related IDs: BE-10-001
Destination: C07, active entry 5
Status: merged into BE-05-001

### BE-16-005
Source: `advanced/16-key-concepts-review.md`
Source section: `HTTP platform` (lines 149-183)
Category: principle
Idea: Handlers translate within strict body, protocol, error,
middleware, health, and resource contracts.
Required nuance: Envelope and business validation differ; panic has a
generic public and correlated private outcome; liveness and readiness
remain distinct.
Related IDs: BE-10-002
Destination: C08, active entry 6
Status: merged into BE-03-001

### BE-16-006
Source: `advanced/16-key-concepts-review.md`
Source section: `Persistence foundation` (lines 184-216)
Category: principle
Idea: Bounded pools, constraints, parameters, cancellation, cleanup,
migrations, and real PostgreSQL tests form the persistence foundation.
Required nuance: Application rules explain while constraints arbitrate
races; acquire late and release early; schema changes deploy like
code.
Related IDs: BE-10-003
Destination: C09, active entry 7
Status: merged into BE-04-001

### BE-16-007
Source: `advanced/16-key-concepts-review.md`
Source section: `Product slice` (lines 217-248)
Category: principle
Idea: A product vertical slice owns capability-local code, total-order
cursor pagination, stable domain errors, and role authorization.
Required nuance: Cursor is a position rather than page number and
unauthenticated differs from forbidden.
Related IDs: BE-10-004
Destination: C07, active entry 5
Status: merged into BE-05-001

### BE-16-008
Source: `advanced/16-key-concepts-review.md`
Source section: `Order transaction` (lines 249-279)
Category: invariant
Idea: One protected database operation commits orders, items, and
inventory while sorted locking and owner predicates preserve races and
privacy.
Required nuance: Public failure hides tables, SQL, and other users;
repositories do not commit pieces independently.
Related IDs: BE-10-005
Destination: C10, active entry 8
Status: merged into BE-04-009

### BE-16-009
Source: `advanced/16-key-concepts-review.md`
Source section: `Idempotency and failure semantics` (lines 280-311)
Category: invariant
Idea: Same scoped key and canonical payload identify one effect and
stored response; changed payload conflicts.
Required nuance: Concurrent attempts serialize, disconnect is safe,
and expiry/retention changes old-retry behavior.
Related IDs: BE-11-001
Destination: C15, active entry 13
Status: merged into BE-11-001

### BE-16-010
Source: `advanced/16-key-concepts-review.md`
Source section: `Observability and operational controls`
(lines 312-347)
Category: operational control
Idea: Bounded logs, metrics, traces, audits, objectives, alerts, and
runbooks answer incident questions and lead to action.
Required nuance: Reconstruct one failure, keep identifiers out of
metric labels, and make every page actionable.
Related IDs: BE-11-002
Destination: C13, active entry 11
Status: merged into BE-07-005

### BE-16-011
Source: `advanced/16-key-concepts-review.md`
Source section: `Deployment hardening` (lines 348-381)
Category: operational control
Idea: Immutable least-privilege artifacts, runtime config, truthful
health, compatible migration, limits, and restore proof enable
release.
Required nuance: Old and new code coexist and recovery requires tested
restoration rather than a successful backup indicator.
Related IDs: BE-11-003
Destination: C14, active entry 12
Status: merged into BE-08-001

### BE-16-012
Source: `advanced/16-key-concepts-review.md`
Source section: `Performance and resilience based on evidence`
(lines 382-413)
Category: test or experiment
Idea: Measure representative tail latency and saturation, profile the
actual bottleneck, change one cause, and measure again.
Required nuance: Avoid coordinated omission, preserve overload safety,
and accept cache consistency cost only with evidence.
Related IDs: BE-11-004
Destination: C16, active entry 14
Status: merged into BE-11-004

### BE-16-013
Source: `advanced/16-key-concepts-review.md`
Source section: `Advanced PostgreSQL operations and recovery`
(lines 414-452)
Category: operational control
Idea: MVCC, isolation, deadlocks, maintenance, resumable migrations,
PITR, replicas, and drills define database operation and recovery.
Required nuance: Retry only whole safe work and distinguish backup,
restore, replication, failover, RPO, and RTO.
Related IDs: BE-11-005
Destination: C17, active entry 15
Status: merged into BE-11-005

### BE-16-014
Source: `advanced/16-key-concepts-review.md`
Source section: `Durable background jobs and scheduling`
(lines 453-487)
Category: mechanism
Idea: Persist intent, lease bounded work, make repeated effects safe,
classify retries, quarantine poison work, and schedule idempotently.
Required nuance: Queue age represents user delay; backpressure and
shutdown preserve dependencies and leased work.
Related IDs: BE-12-001
Destination: C18, active entry 16
Status: merged into BE-12-001

### BE-16-015
Source: `advanced/16-key-concepts-review.md`
Source section: `Reliable event publication and consumption`
(lines 488-523)
Category: mechanism
Idea: Outbox prevents lost publication intent and inbox prevents
unsafe
repeated consumption within stated ordering and replay boundaries.
Required nuance: Relays can duplicate after broker success; exactly
once is scoped; event ownership, semantics, versions, retention, and
external effects remain explicit.
Related IDs: BE-12-002
Destination: C19, active entry 17
Status: merged into BE-12-002

### BE-16-016
Source: `advanced/16-key-concepts-review.md`
Source section: `Distributed-systems mechanics lab` (lines 524-559)
Category: test or experiment
Idea: Predict and simulate partial failure, clocks, leases, fencing,
quorums, hashing, and consensus without production reinvention.
Required nuance: Timeout is ambiguous, fencing protects resources from
old owners, and reconciliation closes partition conflicts.
Related IDs: BE-12-003
Destination: C20, active entry 18
Status: merged into BE-12-003

### BE-16-017
Source: `advanced/16-key-concepts-review.md`
Source section: `Service-boundary discovery and first extraction`
(lines 560-591)
Category: tradeoff
Idea: Business ownership and measured independence justify extraction;
table or layer splits produce distributed monoliths.
Required nuance: Services own data and operations, migration is
incremental, and benefits must exceed latency, failure, deployment,
telemetry, and on-call cost.
Related IDs: BE-13-001
Destination: C21, active entry 19
Status: merged into BE-13-001

### BE-16-018
Source: `advanced/16-key-concepts-review.md`
Source section: `Synchronous service communication and the system
edge` (lines 592-626)
Category: mechanism
Idea: Remote calls consume one deadline and need stable identity,
idempotency, discovery, authorization, bounded retry, and pressure
controls.
Required nuance: Gateways own edge policy rather than shared business
rules; outage remains bounded; ambiguous success needs query or
repair.
Related IDs: BE-13-002
Destination: C22, active entry 20
Status: merged into BE-13-002

### BE-16-019
Source: `advanced/16-key-concepts-review.md`
Source section: `Distributed workflows and eventual consistency`
(lines 627-659)
Category: mechanism
Idea: Durable sagas model intermediate state, correlation, timeout,
compensation, observability, and manual repair across local commits.
Required nuance: Compensation is fallible rather than rewind and stuck
states are expected operator-visible conditions with time bounds.
Related IDs: BE-13-003
Destination: C23, active entry 21
Status: merged into BE-13-003

### BE-16-020
Source: `advanced/16-key-concepts-review.md`
Source section: `Contract evolution and independent delivery`
(lines 660-690)
Category: tradeoff
Idea: Compatibility spans schema and semantics, and deprecation needs
real consumer evidence before removal.
Required nuance: Additive changes can break strict consumers; both
deployment orders and delayed older events need tests.
Related IDs: BE-13-004
Destination: C24, active entry 22
Status: merged into BE-13-004

### BE-16-021
Source: `advanced/16-key-concepts-review.md`
Source section: `Distributed observability and resilience validation`
(lines 691-724)
Category: test or experiment
Idea: Safe trace propagation, bounded metrics and sampling, and
hypothesis-driven chaos reveal emergent cross-service failures.
Required nuance: Bound baggage, retry storms, fan-out, backlog, blast
radius, abort, and recovery evidence.
Related IDs: BE-13-005
Destination: C25, active entry 23
Status: merged into BE-13-005

### BE-16-022
Source: `advanced/16-key-concepts-review.md`
Source section: `Data scaling, caching, and multi-tenancy`
(lines 725-761)
Category: tradeoff
Idea: Cache, replica, partition, archive, derived view, and tenant
choices solve distinct measured constraints and consistency problems.
Required nuance: Isolation fails closed across every path and cache or
replica state never silently becomes authority.
Related IDs: BE-14-001
Destination: C26, active entry 24
Status: merged into BE-14-001

### BE-16-023
Source: `advanced/16-key-concepts-review.md`
Source section: `Security, privacy, and software supply chain`
(lines 762-798)
Category: operational control
Idea: Trust boundaries require workload identity, least privilege,
rotation, data lifecycle, and verifiable build artifacts.
Required nuance: Deletion spans all copies and a signed
SBOM/provenance
chain supports but does not itself perform vulnerability response.
Related IDs: BE-14-002
Destination: C27, active entry 25
Status: merged into BE-14-002

### BE-16-024
Source: `advanced/16-key-concepts-review.md`
Source section: `Cloud infrastructure and IAM` (lines 799-836)
Category: operational control
Idea: Versioned IaC owns network, identity, private resources, state,
audit, cost, drift, and scoped teardown.
Required nuance: Reachability and authorization are independent and
must be denied and diagnosed independently.
Related IDs: BE-14-003
Destination: C28, active entry 26
Status: merged into BE-14-003

### BE-16-025
Source: `advanced/16-key-concepts-review.md`
Source section: `Orchestration and platform operations`
(lines 837-870)
Category: operational control
Idea: Reconciliation, resources, probes, discovery, policy, rollout,
and scaling remain platform contracts around application correctness.
Required nuance: Autoscale the bottleneck without exceeding dependency
budgets and preserve failure visibility through abstractions.
Related IDs: BE-14-004
Destination: C29, active entry 27
Status: merged into BE-14-004

### BE-16-026
Source: `advanced/16-key-concepts-review.md`
Source section: `Reliability engineering, incidents, and cost`
(lines 871-906)
Category: operational control
Idea: User SLIs, objectives, budgets, capacity, response roles,
review,
and total cost connect reliability to delivery decisions.
Required nuance: Pages are actionable and corrective work has owner,
deadline, verification, and completion evidence.
Related IDs: BE-14-005
Destination: C30, active entry 28
Status: merged into BE-14-005

### BE-16-027
Source: `advanced/16-key-concepts-review.md`
Source section: `Common production integration boundaries`
(lines 907-943)
Category: mechanism
Idea: External calls and messages need bounded adapters, ambiguity
handling, verification, streaming, asynchronous outcomes, and rebuild.
Required nuance: Idempotency or reconciliation handles timed-out
writes; webhook raw bytes and replay identity precede effects; derived
state is disposable.
Related IDs: BE-15-001
Destination: C31, active entry 29
Status: merged into BE-15-001

### BE-16-028
Source: `advanced/16-key-concepts-review.md`
Source section: `Maintenance and legacy-system evolution`
(lines 944-977)
Category: procedure
Idea: Observe, characterize, make reversible change, verify, and
remove
temporary flags, branches, dependencies, and repair machinery.
Required nuance: Refactoring preserves behavior while migration
changes
it compatibly; repair remains scoped and auditable.
Related IDs: BE-15-002
Destination: C32, active entry 30
Status: merged into BE-15-002

### BE-16-029
Source: `advanced/16-key-concepts-review.md`
Source section: `Capstone delivery and engineering leadership`
(lines 978-1012)
Category: completion evidence
Idea: Leadership traces user needs through alternatives, review,
delivery, evidence, operation, handoff, teaching, and simplification.
Required nuance: Critical feedback changes the design visibly and
ownership includes naming complexity removable under simpler needs.
Related IDs: BE-15-003
Destination: C33, active entry 31
Status: merged into BE-15-003

### BE-16-030
Source: `advanced/16-key-concepts-review.md`
Source section: `Cross-topic recall checklist` (lines 1013-1041)
Category: completion evidence
Idea: Every capability review asks the same contract, data, time,
resource, failure, distribution, security, observability, delivery,
evidence, and simplicity questions.
Required nuance: The checklist links topics but does not replace any
owning explanation or proof.
Related IDs: BE-00-043, BE-16-001
Destination: C01, supporting introduction and review checklist
Status: merged into BE-00-001

### BE-17-001
Source: `advanced/17-language-portability.md`
Source section: introduction and `The portability contract`
(lines 3-36)
Category: principle
Idea: A backend port preserves contracts, invariants, failure
behavior,
resource bounds, security, evidence, operations, and delivery using
native ecosystem idioms rather than Go-shaped syntax.
Required nuance: Framework defaults are not proof; effective limits,
retry, serialization, middleware, and shutdown need repository-owned
tests or configuration.
Related IDs: BE-00-048, BE-17-002
Destination: C34, active entry 32
Status: incorporated

### BE-17-002
Source: `advanced/17-language-portability.md`
Source section: `Concept mapping` (lines 37-107)
Category: language-specific behavior
Idea: HTTP, cancellation, error, concurrency, PostgreSQL, wiring,
logging, and test concepts map to different native mechanisms in Java,
.NET, Node.js, Python, and Rust.
Required nuance: Listed frameworks are examples rather than mandates;
prefer the established stack and make implicit boundary behavior
explicit.
Related IDs: BE-17-001, BE-17-003
Destination: C34, active entry 32
Status: merged into BE-17-001

### BE-17-003
Source: `advanced/17-language-portability.md`
Source section: `Universal and Go-specific material` (lines 108-121)
Category: language-specific behavior
Idea: Backend contracts, data, security, operations, and distribution
are universal while packages, slices, errors, defer, goroutines,
context, and Go tools need idiomatic replacement.
Required nuance: Replace syntax exercises but retain the business
exercise and its completion evidence.
Related IDs: BE-01-002 through BE-02-012
Destination: C34, active entry 32
Status: merged into BE-17-001

### BE-17-004
Source: `advanced/17-language-portability.md`
Source section: `Adapting the core chapters` (lines 122-191)
Category: procedure
Idea: Rebuild language ownership, bounded HTTP, real PostgreSQL,
application boundaries, identity, deterministic evidence, immutable
delivery, and incident proof in the target ecosystem.
Required nuance: Preserve strict negative cases, object-scoped access,
cancellation, driver semantics, mixed schema rollout, restore,
capacity, and lifecycle; do not reproduce Go abstractions for
symmetry.
Related IDs: BE-17-001, BE-17-002
Destination: C34, active entry 32
Status: merged into BE-17-001

### BE-17-005
Source: `advanced/17-language-portability.md`
Source section: `Java and Kotlin` and `C# and .NET` (lines 192-217)
Category: language-specific behavior
Idea: Managed-runtime ports must prove effective transaction proxies,
DI lifetimes, cancellation, executor and client bounds, and generated
data behavior.
Required nuance: Self-invocation can bypass transaction annotations;
interrupts must survive; scoped dependencies cannot enter singletons;
ORM tracking, N+1 queries, locks, retries, and SQL remain visible.
Related IDs: BE-04-008, BE-17-004
Destination: C34, active entry 32
Status: merged into BE-17-001

### BE-17-006
Source: `advanced/17-language-portability.md`
Source section: `TypeScript and Node.js` and `Python` (lines 218-237)
Category: language-specific behavior
Idea: Static hints and event loops do not validate runtime input or
make multi-step state changes atomic.
Required nuance: Propagate abort or task cancellation, keep blocking
work off async loops, observe detached failures, and bound promises,
tasks, workers, sockets, queues, bodies, and pools.
Related IDs: BE-06-010, BE-17-004
Destination: C34, active entry 32
Status: merged into BE-17-001

### BE-17-007
Source: `advanced/17-language-portability.md`
Source section: `Rust` (lines 238-248)
Category: language-specific behavior
Idea: Rust memory safety does not establish transaction, distributed,
or cancellation correctness.
Required nuance: Dropped futures need cancellation-safe design,
blocking work stays off async executors, tasks and channels remain
bounded, and domain error enums translate infrastructure at the edge.
Related IDs: BE-13-003, BE-17-004
Destination: C34, active entry 32
Status: merged into BE-17-001

### BE-17-008
Source: `advanced/17-language-portability.md`
Source section: `Equivalent verification pipeline` (lines 249-267)
Category: completion evidence
Idea: Every stack owns exact commands for format, warnings-as-errors,
unit and PostgreSQL tests, concurrency, fuzzing, dependencies,
contracts, container smoke, and current failure drills.
Required nuance: A clean clone reproduces the pipeline without IDE or
hidden machine state.
Related IDs: BE-08-005, BE-17-001
Destination: C34, active entry 32
Status: merged into BE-17-001

### BE-17-009
Source: `advanced/17-language-portability.md`
Source section: `Completion standard in another language`
(lines 268-285)
Category: completion evidence
Idea: A port is complete when an ecosystem-native engineer can
explain,
build, test, contend, diagnose, deploy, restore, defend, and change it
without Go-specific assumptions.
Required nuance: Equivalent guarantees and operational evidence, not
surface API similarity, define completion.
Related IDs: BE-00-045, BE-17-008
Destination: C34, active entry 32
Status: merged into BE-17-001
