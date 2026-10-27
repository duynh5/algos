# Backend Engineering in Go

This guide owns the active backend teaching material. The lessons use
one evolving Order API to connect transport contracts, data safety,
architecture, and production operations.

## Contents

- [Backend competency map](#backend-competency-map)
- [Further study](#further-study)
- [Go values and
  ownership](#1-go-values-program-structure-and-ownership)
- [Errors and cancellation](#2-go-errors-resources-and-cancellation)
- [Concurrency and
  verification](#3-go-concurrency-and-executable-verification)
- [Domain discovery](#4-domain-discovery-and-modeling)
- [Architecture](#5-architecture-and-service-construction)
- [HTTP and process
  lifecycle](#6-http-contracts-and-process-lifecycle)
- [Relational data](#7-relational-modeling-and-postgresql-access)
- [Transactions](#8-business-transactions-and-contested-state)
- [Identity and security](#9-identity-authorization-and-api-security)
- [Deterministic testing](#10-deterministic-backend-testing)
- [Observability](#11-observability-and-operational-controls)
- [Delivery](#12-delivery-deployment-and-recovery-readiness)
- [Idempotency](#13-idempotency-and-ambiguous-write-outcomes)
- [Capacity and overload](#14-performance-capacity-retry-and-overload)
- [PostgreSQL
  operations](#15-advanced-postgresql-operation-and-recovery)
- [Durable jobs](#16-durable-background-jobs-and-scheduling)
- [Reliable events](#17-reliable-event-publication-and-consumption)
- [Distributed mechanics](#18-distributed-systems-mechanics)
- [Service extraction](#19-service-boundary-discovery-and-extraction)
- [Service communication](#20-synchronous-service-communication)
- [Distributed
  workflows](#21-distributed-workflows-and-eventual-consistency)
- [Contract
  evolution](#22-contract-evolution-and-independent-delivery)
- [Resilience experiments][resilience-experiments]
- [Scaling and tenancy](#24-data-scaling-caching-and-tenant-isolation)
- [Privacy and supply
  chain](#25-security-privacy-and-software-supply-chain)
- [Cloud and IAM](#26-cloud-infrastructure-and-iam)
- [Platform operations](#27-orchestration-and-platform-operations)
- [Reliability and
  cost](#28-reliability-engineering-incidents-and-cost)
- [Integrations](#29-production-integration-boundaries)
- [Legacy evolution](#30-maintenance-and-legacy-system-evolution)
- [Leadership](#31-capstone-delivery-and-engineering-leadership)
- [Language portability](#32-language-portability)

## How to use this guide

Read the lessons in order on a first pass. For each lesson, implement
the example in the Order API lab, exercise its failure path, and keep
the resulting test or operational evidence. Return through the
contents when using the guide as a reference.

The competency map is a supporting roadmap for the active entries.
The 32 entries that follow own the teaching contract and progress from
Go mechanics through production leadership. Repository-owned examples
and tests provide executable evidence without a second set of lesson
owners.

## Backend competency map

Backend engineering is the work of preserving useful behavior across
untrusted input, shared data, concurrency, process restarts, network
failure, deployment, and change. Language and framework skill matter,
but the final evidence is a service another engineer can build, test,
operate, recover, and modify safely.

The map below incorporates the substantive ideas from every imported
backend chapter. Reusable exercises and evidence requirements appear
in the active entries or the cross-cutting review questions below.

### Learning model and target system

- Learn through one evolving Order API. Start with a modular monolith
  whose packages separate transport, application, domain, and data
  concerns. Add a network boundary only when ownership or isolation
  evidence justifies its cost.
- Treat a deployed service as an operating-system process. Account for
  its arguments, environment, identity, permissions, signals, exit
  codes, filesystem, memory, CPU, threads, file descriptors, and
  sockets. Containers isolate resources; they do not remove them.
- Treat reading as orientation. Competence requires implementation,
  negative tests, concurrency tests, operational evidence, and an
  explanation of why the result stays correct during failure.
- Separate confirmed facts, assumptions, open questions, and deferred
  scope. A plausible guess must not silently become a public contract.
- Prefer one reproducible path from a clean clone to a running
  service. Local and continuous-integration commands should exercise
  the same formatter, analysis, tests, race checks, vulnerability
  checks, build, and smoke tests.
- Use an architecture decision record for choices with lasting
  tradeoffs. Record the context, alternatives, decision,
  consequences, and evidence that would cause reconsideration.
- Organize delivery as small vertical changes. Trace each requirement
  to code, tests, telemetry, rollout behavior, and a runbook.

### Go language and backend mechanics

- A module is the dependency and versioning boundary named by
  `go.mod`. A package is the directory-level compilation and
  visibility boundary. Uppercase names are exported, while
  `internal` packages restrict importers by directory ancestry.
- Go is statically typed even when `:=` infers a local type. Use exact
  integer or decimal representations for money, explicit conversions
  at boundaries, and useful zero values when a type can be safe
  without construction.
- `if`, `switch`, and the forms of `for` keep control flow explicit.
  `range` over a string yields byte offsets and Unicode code points;
  it does not yield human-perceived characters.
- Functions may return several values and errors. Closures capture
  variables, so their lifetime and mutation must remain clear.
  Variadic calls and generics are useful only when they remove real
  repetition without hiding domain meaning.
- Arrays own fixed-size storage. A slice is a descriptor over its
  backing storage, so copies can alias and `append` can reuse or
  replace that storage. Copy mutable slices when crossing an
  ownership boundary.
- Maps return zero values for absent keys, so use the comma-`ok` form
  when absence differs from a stored zero. Initialize before writes,
  never depend on iteration order, and synchronize concurrent access.
- Strings are immutable bytes, normally UTF-8. Choose deliberately
  among byte counts, code points, and grapheme-like visible
  characters; arbitrary byte truncation can split encoded text.
- Structs group values. Use a pointer when mutation, identity, or
  meaningful absence requires sharing; do not use one by default.
  Embedding promotes members for composition but is not inheritance.
- Interfaces are satisfied implicitly. Define the smallest useful
  interface near its consumer and usually return concrete types. Use
  checked assertions or type switches when a dynamic type is not an
  internal certainty.
- `nil` slices can be read and appended to, while `nil` maps cannot be
  written. `make` initializes slices, maps, and channels; `new`
  returns a pointer to a zero value. An interface containing a typed
  nil pointer is not itself nil.
- Errors are values. Wrap them with operation context while preserving
  identity for `errors.Is` or type information for `errors.As`. Use
  `panic` for broken internal assumptions, and recover only at a
  deliberate boundary that reports a generic failure and records the
  defect.
- Acquire a resource, check success, and immediately arrange cleanup
  with `defer` when function scope is the right ownership boundary.
  Do not retain scarce resources until the end of a large loop.
- Goroutines are owned resources. Every goroutine needs a termination
  path and observable failure. Channels communicate and coordinate;
  mutexes protect shared invariants. The producer that knows no more
  values will arrive normally owns channel closure.
- A buffered channel is a bounded queue, not infinite capacity.
  Protect complete read-check-write rules, not just individual map
  operations, and use the race detector without confusing memory-race
  freedom with business atomicity.
- Pass `context.Context` as the first argument to work that can wait.
  Propagate cancellation and deadlines through database and network
  calls; do not store request context in a long-lived struct.
- Avoid accidental `:=` shadowing, hidden package state, and `init`
  ordering. Copy values according to what their nested fields share,
  sort map-derived output when stability matters, and make resource
  ownership explicit.

### Domain discovery and architecture

- Treat requirements as hypotheses. Concrete examples should expose
  preconditions, successful transitions, rejected actions, failure
  outcomes, and unresolved policy before endpoint or table design.
- Build a ubiquitous language: one agreed meaning for each important
  term across product discussion, code, contracts, events, and
  operations. Different bounded contexts may intentionally use
  different models of a word such as `Product`.
- An entity keeps identity while changing. A value object is defined
  by its values. A command asks for change; an event names a fact that
  already occurred.
- State machines make allowed transitions, terminal states, and late
  or duplicate actions explicit. Invariants state what must always be
  true independently of HTTP, Go, or SQL.
- A bounded context owns a coherent model and vocabulary. It can be a
  package in a modular monolith and is not automatically a service.
- Architecture assigns reasons to change. Transport values,
  application commands, domain values, and persistence rows may
  represent one concept differently because they answer different
  concerns.
- Prefer vertical feature packages to global controller, service, and
  repository buckets. Dependencies point inward toward policy, while
  outer adapters translate protocols and infrastructure.
- The composition root parses and validates typed configuration,
  constructs concrete dependencies, and owns process lifecycle.
  Business packages should not discover global infrastructure.
- Dependency injection is ordinary construction. Add a narrow,
  consumer-owned interface only at a real substitution or boundary;
  avoid speculative abstractions.
- One use case owns one business operation and its transaction.
  Stable domain error identities cross inward boundaries, then the
  transport edge maps them without parsing error text.

### HTTP contracts and service lifecycle

- Trace a request through DNS, TCP, Transport Layer Security (TLS), a
  proxy or load balancer, the Go server, application code, database,
  and downstream clients. Know which layer owns each timeout, pool,
  identity check, and error.
- Treat HTTP method, target, headers, representation, status, caching,
  pagination, limits, ordering, and retry semantics as one public
  contract. Do not expose persistence rows as that contract.
- Handlers translate rather than decide business policy: bound and
  decode input, validate the transport envelope, call one use case,
  and map its known outcomes to stable responses.
- Enforce media type, body size, unknown-field rejection, and exactly
  one JSON value. Client-side validation never makes server-side
  validation optional.
- Separate syntax checks, business validation, authentication,
  authorization, conflict, not-found, dependency failure, and
  internal failure. Use stable machine-readable problem details and a
  correlation identity without exposing internals.
- Middleware order is observable. Establish request identity before
  logging and panic recovery, authenticate before authorization, and
  keep business authorization near the action that owns the rule.
- Give the server explicit header, body, idle, request, and shutdown
  bounds. A timeout ends waiting; it does not prove remote work
  stopped or that a write failed to commit.
- Scope each idempotency key to a principal and operation. Bind it to
  a canonical request fingerprint, serialize concurrent duplicates,
  and store the stable completed response with the business effect.
  Different input under one key is a conflict, not a retry.
- Define idempotency retention from the client's retry window. Bounded
  cleanup, expiry behavior, and the response for in-progress work are
  part of the public contract.
- Liveness asks whether restart may help. Readiness asks whether the
  instance should receive traffic. Startup indicates initialization
  is complete. Conflating them causes premature traffic or restart
  loops.
- Graceful shutdown is a bounded protocol: stop intake, mark the
  instance unready, drain in-flight work, release resources, then
  terminate at the deadline.
- Stable cursor pagination records a position in a deterministic
  total order, usually with a unique tie-breaker. Offset pagination
  can skip or repeat records during concurrent changes.

### Relational data and business transactions

- Model relations from business invariants. Primary and foreign keys,
  unique and check constraints, nullability, and exact data types are
  the final integrity barrier when application checks race.
- Missing, empty, and zero have different meanings. Store money as
  exact decimal values or integer minor units, event time as a UTC
  instant, and civil scheduling with its named time zone.
- Historical facts keep the values used at the time. For example, an
  order item stores the charged price rather than depending on the
  product's current price.
- An index trades write work and storage for a particular lookup or
  ordering. Verify the actual query plan and data shape; redundant or
  unused indexes still impose costs.
- Parameter binding separates data from SQL. Object-relational tools
  do not remove the need to inspect generated queries, row counts,
  lock behavior, or the N+1 query pattern.
- A connection pool is bounded shared concurrency. Configure its
  size, acquisition deadline, lifetime, health, and relationship to
  database capacity. Acquire late and release early.
- Context cancellation must reach pool acquisition, query execution,
  row iteration, and commit. Close rows, roll back unfinished
  transactions, check iteration errors, and handle commit failure.
- A transaction follows one complete business operation. Keep it
  short, avoid external calls while holding it, and publish success
  only after commit.
- Protect contested state with an atomic conditional update, ordered
  row locks, or an explicitly chosen isolation level. An earlier read
  does not reserve inventory.
- Object ownership belongs in the data predicate, such as querying by
  both object and owner identity. This both enforces access and avoids
  revealing another user's object.
- Multi-version concurrency control allows consistent reads while
  writes continue. Isolation chooses which anomalies are prevented
  and which operations may wait or abort; it is not one generic
  safety switch.
- Deterministic lock order prevents many deadlocks. When PostgreSQL
  aborts a participant, retry only the entire operation, with a bound,
  and only when repetition is safe.
- Vacuum, statistics, bloat, checkpoints, and the write-ahead log
  affect reuse, planning, write latency, replication, and recovery.
  These are application-operability concerns, not only database
  administration.
- Migrations are ordered production changes. They must work from an
  empty database and with existing data, mixed application versions,
  interruption, verification, and recovery.
- Large backfills should be lock-aware, batched, resumable,
  observable, and independently verified. One huge transaction is
  not automatically safer.
- Base backups plus archived write-ahead logs can support point-in-
  time recovery. Replication improves availability, but it can copy
  accidental deletion. Restore and promotion drills must prove the
  recovery point and recovery time objectives.

### Identity, authorization, and API security

- Authentication establishes who or what the principal is.
  Authorization decides whether that principal may perform this
  action on this object. Default to deny and test both role and
  horizontal ownership boundaries.
- Distinguish the resource owner, client, authorization server, and
  resource server. Validate that every token is intended for this
  issuer, audience, algorithm, key source, time window, and use.
- Discovery metadata and JSON Web Key Set caching need refresh and
  safe key rotation. Bounded clock skew handles small clock
  differences without accepting indefinitely expired tokens.
- Build a small trusted principal after validation. Do not pass raw
  claims throughout the application or treat possession of an opaque
  object identifier as permission.
- JSON Web Tokens enable local validation but complicate immediate
  revocation. Opaque tokens or sessions centralize current state but
  require a lookup. Choose from the actual revocation, latency, and
  availability needs.
- Browser cookies need secure attributes, safe login and logout, and
  cross-site request forgery protection. Cross-origin resource
  sharing controls which browser origins may read a response; it is
  not authentication.
- Decode into allowlisted request values to prevent over-posting and
  mass assignment. Parameterize queries, escape output for its sink,
  and avoid logging credentials, bodies, or personal data.
- Store secrets outside source, scope access narrowly, rotate without
  outage, and make abuse controls explicit. Rate limits, quotas, body
  limits, and bounded concurrency protect different resources.
- Threat modeling follows assets, actors, entry points, trust
  boundaries, abuse cases, mitigations, and remaining risk. Repeat it
  when a new service, queue, datastore, admin path, or integration
  changes the boundary.

### Testing, observability, and operations

- Test at the boundary that owns the risk. Unit tests cover pure
  rules; handler tests cover HTTP; real-PostgreSQL integration tests
  cover SQL, constraints, locks, and transactions; contract tests
  cover consumer expectations; acceptance tests cover critical
  workflows.
- Add concurrency tests for races in business effects, fuzz or
  property tests for parsers and invariants, migration tests for empty
  and populated states, load tests for stated workloads, and bounded
  failure experiments for recovery claims.
- Table-driven tests make cases visible. Prefer small fakes with
  controlled behavior to mocks that duplicate implementation. Control
  clocks, identifiers, randomness, and scheduling where assertions
  depend on them.
- A regression test should fail for the intended reason before the
  fix when practical. Characterization tests pin important existing
  behavior before risky refactoring.
- Logs describe discrete events, metrics show aggregate trends,
  traces connect causal work, and profiles locate CPU, allocation,
  goroutine, lock, or blocking costs. None replaces the others.
- Instrument stable boundaries with correlation and trace context.
  Keep secrets and personal data out of telemetry, and keep
  high-cardinality values such as user IDs and raw URLs out of metric
  labels.
- Useful service signals cover traffic, error, latency distributions,
  and saturation. A service-level indicator measures user-visible
  behavior; a service-level objective sets its target over a window;
  the error budget states allowed unreliability.
- Alerts should represent sustained user impact or imminent resource
  exhaustion and lead to a concrete action. Runbooks connect symptoms
  to investigation, mitigation, rollback, escalation, and recovery
  verification.
- Bound request bodies, goroutines, queues, workers, connections,
  retries, caches, telemetry cardinality, and shutdown. Backpressure
  and predictable rejection preserve useful work during overload.
- Retry only a transient failure, only when repetition is safe, and
  only within attempt, delay, and total-time bounds. Add exponential
  backoff and jitter, honor server guidance, and avoid retries at
  several layers.

### Delivery, performance, and production readiness

- Build one immutable, traceable artifact and promote it through
  environments. Parse runtime configuration into typed values and
  fail before serving when required preconditions are invalid.
- Keep containers small, non-root, and read-only except for explicit
  ephemeral paths. Declare resource requests and limits so scheduling
  and exhaustion behavior are visible.
- Continuous integration reproduces the local quality contract.
  Warnings fail, dependencies and images are checked, and the built
  artifact receives a smoke test.
- Application and schema deployment are not atomic. Use expand and
  contract: add a compatible shape, deploy compatible code, backfill
  and verify, switch use, then remove the old shape after the rollback
  window closes.
- Rollback restores an older compatible artifact. Roll-forward repairs
  a state that cannot safely accept the old version. Decide and
  rehearse both paths before release.
- A backup is only a recovery possibility until a restore drill proves
  data and application invariants. Rehearse startup failure,
  readiness loss, termination, clean database creation, and incident
  response.
- Throughput measures completed work per time. Latency is a
  distribution. Utilization shows resource use, while saturation
  exposes waiting near a bound; tail latency often reveals the user
  problem before averages do.
- Representative data, traffic mix, concurrency, warmup, and
  environment are part of every performance result. Load tools must
  avoid coordinated omission, which hides delays by issuing less work
  while the system is slow.
- Profile the limiting workload before optimizing. Measure, locate one
  bottleneck, change one thing, and measure again. Add a cache only
  after evidence shows repeated expensive reads and the consistency
  cost is acceptable.
- Capacity planning links expected demand, dependency ceilings,
  headroom, growth, recovery, and the first saturated resource.
  Cost includes infrastructure, engineering time, operational burden,
  and lost opportunities.

### Durable work, messages, and distributed mechanics

- A goroutine cannot preserve important work through process exit.
  Durable jobs persist intent, lease work to bounded workers, classify
  retryable failures, quarantine poison work, and expose replay or
  cancellation to operators.
- At-least-once execution means a handler can repeat after completing
  an effect but before acknowledgement. Make each effect idempotent,
  and stop intake before draining or releasing leases at shutdown.
- Claim jobs in bounded batches and give each lease a unique attempt
  token. Completion, extension, or retry must match that token so a
  paused stale worker cannot overwrite the current claim.
- Schedulers can run on several replicas. Coordinate the trigger or
  make the scheduled operation idempotent, and observe queue age as
  well as queue length.
- A local database transaction cannot normally commit atomically with
  a broker publication. A transactional outbox stores the business
  change and publication intent together; a relay tolerates repeated
  publish attempts.
- Consumers expect duplicates. An inbox or deduplication record should
  commit with the local state change when a repeated effect is unsafe.
  Ordering is promised only within a stated key or partition.
- Event contracts need owners, past-tense facts, stable identities,
  semantics, compatibility rules, version strategy, retention,
  ordering, and replay behavior. "Exactly once" always ends at some
  boundary.
- Distributed failure is partial. A timeout creates uncertainty, wall
  clocks can jump or disagree, and timestamps do not create a total
  order. Use monotonic time for elapsed duration.
- Quorum choices trade consistency, availability, and latency under a
  stated failure model. Consistent hashing limits key movement but
  does not itself provide consistency.
- A lease can expire while an old owner is paused. A monotonically
  increasing fencing token lets the protected resource reject that
  stale owner. Use maintained consensus systems instead of inventing
  a protocol.
- Deterministic simulations of delay, loss, duplication, reordering,
  and partitions produce better evidence than tests built on sleeps.
  Define the convergence or reconciliation rule after healing.

### Service boundaries and distributed delivery

- A service boundary aligns business capability, data ownership, team
  responsibility, deployment, telemetry, and operations. Splitting by
  table or layer creates coordination without autonomy.
- The owning service is the only writer and contract authority for its
  data. Shared tables, cross-service queries, and lockstep releases
  are signs of a distributed monolith.
- Extraction adds latency, partial failure, contract evolution,
  security, pipelines, telemetry, and on-call cost. Require measured
  autonomy, scaling, or isolation value before paying that cost.
- Branch by abstraction places a stable seam around old behavior,
  runs old and new implementations during migration, and shifts
  traffic incrementally. Prefer a low-risk asynchronous capability
  for a first extraction.
- Choose HTTP with JSON or a generated binary remote-call protocol
  from client compatibility, tooling, streaming, contract, and
  latency needs rather than fashion.
- The caller owns one end-to-end deadline budget. Bound every
  dependency pool and concurrency path, propagate cancellation, and
  prevent nested retry policies from multiplying work.
- Circuit breakers limit calls to a failing dependency. Bulkheads
  contain shared-resource damage. Load shedding rejects excess work.
  A fallback is valid only when its answer remains semantically
  honest.
- Gateways can own TLS, routing, authentication, and coarse rate
  limits. The service that owns the business rule still owns
  authorization.
- Cross-service workflows cannot share an ordinary local transaction.
  A saga records a durable sequence of local steps. Orchestration
  centralizes progress; choreography reacts to events and can obscure
  the workflow as interactions grow.
- Compensation is a new, fallible business action rather than time
  travel. Persist current step, attempts, deadlines, correlation, late
  messages, compensation, and manual-repair decisions.
- Eventual consistency makes intermediate states part of the model.
  Bound how long they may remain, guard transitions against delayed or
  duplicate messages, and provide audited operator repair commands.
- Independent deployment requires provider and consumer versions to
  coexist within a compatibility window. Compatibility includes
  defaults, validation, ordering, timing, and error meaning, not only
  schema shape.
- Consumer-driven contract tests record behavior a consumer depends
  on. They complement, rather than replace, integration and end-to-end
  tests. Deprecation requires an owner, notice, telemetry, migration
  path, deadline, and measured absence of callers.
- Propagate trace context through remote calls, events, and jobs.
  Keep trace baggage small and nonsensitive. Sampling should retain a
  useful share of slow and failed operations.
- Chaos experiments need a hypothesis, bounded blast radius, abort
  condition, observation plan, and recovery proof. Test claims, not
  random destruction.

### Scale, tenancy, privacy, cloud, and reliability

- Cache-aside loads from the source of truth after a miss. Define
  entry lifetime, invalidation, stampede protection, bypass, and
  failure behavior. The system must remain correct when the cache is
  empty or stale.
- Read replicas reduce primary read work but add lag. Use them only
  for operations whose staleness and read-after-write behavior are
  acceptable.
- Indexing, archival, retention, partitioning, replicas, caches, and
  sharding solve different growth problems. Introduce the smallest
  measured intervention with rollback and reconciliation.
- Tenant identity is an invariant across requests, queries, jobs,
  events, cache keys, quotas, encryption, telemetry, and operator
  tools. Shared tables, per-tenant schemas, and per-tenant databases
  trade density, isolation, customization, restore, and cost.
- Derive tenant identity from trusted authentication, not a body
  field. Include it in owned keys and constraints, fail closed when it
  is absent, and prevent pooled connections from retaining another
  tenant's session scope. Test privileged bypass separately.
- Search and analytics are derived views. Keep an authoritative source
  and complete replay, rebuild, and reconciliation paths.
- Privacy engineering records each sensitive field's purpose, owner,
  location, access, retention, export, and deletion. Backups and
  immutable logs may need expiry or cryptographic erasure rather than
  immediate physical deletion.
- Distinguish end-user identity from workload identity. Encrypt in
  transit, use mutual authentication when warranted, prefer
  short-lived credentials, grant least privilege, audit elevation,
  and rehearse rotation.
- A software bill of materials identifies components. Build
  provenance and signatures connect source, build, and artifact. A
  vulnerability process must identify, rebuild, verify, deploy, and
  confirm replacement of affected artifacts.
- Cloud accounts or projects are administrative boundaries; regions
  and availability zones are failure and latency boundaries. Match
  the design to the provider's actual guarantees.
- Routes determine possible network paths. Security controls restrict
  them. Network address translation permits controlled outbound
  access, and private endpoints avoid unnecessary public transit.
- Domain Name System records, certificates, TLS termination, load
  balancers, and health checks form layers at the edge. Telemetry must
  distinguish failure at each layer.
- Managed services transfer infrastructure work, not application
  ownership of schema, permissions, capacity, failure behavior, or
  recovery.
- Infrastructure as code needs protected state, review, policy checks,
  drift detection, ownership labels, budgets, and precise teardown.
  Required console-only steps defeat reproducibility.
- An orchestrator reconciles desired and observed process state. It
  cannot repair corrupt business state or unsafe retry behavior.
  Requests, limits, disruption budgets, rollout settings, probes,
  networking, secrets, and storage remain explicit contracts.
- Autoscale from a signal tied to the bottleneck and respect database
  or broker capacity. Scaling only the frontmost workload can amplify
  an overload.
- Incident response separates coordination, investigation,
  communication, and mitigation. A blameless review explains system
  conditions and evidence, then assigns corrective work with owners,
  deadlines, and verification.

### Integrations, maintenance, and leadership

- Put each third-party system behind a narrow adapter with deadlines,
  quotas, bounded retries, stable internal errors, fakes, sandbox
  tests, telemetry, and reconciliation. A timed-out provider write
  may already have succeeded.
- Use provider idempotency keys or status lookup before repeating an
  ambiguous effect. Treat external systems as slow, fallible,
  quota-bound, and semantically different from internal models.
- Verify webhook signatures over the exact raw payload, enforce a
  timestamp or replay window, deduplicate deliveries, and acknowledge
  according to the provider's retry contract. Support verification
  key rotation with an overlap window.
- Subscriber-controlled URLs create server-side request forgery risk.
  Require HTTPS, reject private and link-local destinations after
  resolution, revalidate redirects, defend against DNS rebinding, and
  restrict runtime egress.
- Stream large objects directly to object storage with short-lived,
  narrow upload grants, random keys, authorization, size and type
  limits, checksums, encryption, lifecycle policy, and asynchronous
  scanning. Treat filenames and client media types as untrusted.
  Never buffer an unbounded transfer in process memory.
- Notification delivery is asynchronous. Bounces, complaints,
  suppression lists, preferences, and provider callbacks update later
  state.
- Discover legacy behavior from code, tests, telemetry, stored data,
  documentation, and user dependence. Protect important behavior with
  characterization tests before changing it.
- Feature flags, canaries, shadow traffic, and compatibility windows
  make change observable and reversible. Give every temporary path an
  owner and removal condition.
- Upgrades need release-note review, compatibility tests, staged
  rollout, and a layer-specific rollback or roll-forward plan. Data
  repair tools need exact scope, dry run, resume, idempotency, audit,
  and independent verification.
- Refactoring preserves observable behavior. Migration intentionally
  changes it through compatible steps. Remove obsolete flags,
  branches, dependencies, and scaffolding after evidence proves the
  transition complete.
- Engineering leadership connects a user problem to requirements,
  alternatives, tradeoffs, implementation, review, tests, telemetry,
  rollout, operations, and cost.
- Design documents expose assumptions, rejected options, consequences,
  and unresolved risk before implementation hardens a choice. Record
  how critical review changed the design.
- Operational ownership includes dashboards, migration, staged
  release, failure and recovery drills, incident learning, and a
  handoff another engineer can use. Simplicity includes naming what
  can be removed when requirements shrink.

### Language portability

- Go is the reference implementation, not the definition of backend
  quality. A port must preserve public contracts, invariants,
  transaction boundaries, cancellation, bounded resources, stable
  errors, authorization, tests, telemetry, rollout, and failure
  evidence using the target ecosystem's idioms.
- Translate mechanisms rather than syntax: `context.Context` may
  become a cancellation token or abort signal; returned errors may
  become results or typed exceptions; goroutines may become tasks,
  coroutines, threads, or actors.
- Framework defaults are not evidence. Verify the effective decoder,
  middleware order, transaction behavior, timeouts, pool limits,
  retries, serialization, and shutdown in repository-owned tests and
  configuration.
- Preserve real PostgreSQL tests. An in-memory substitute cannot prove
  PostgreSQL constraints, multi-version behavior, locks, plans, error
  codes, or migrations.
- Java and Kotlin ports must verify proxy-based transactions, executor
  bounds, cancellation, and generated SQL. .NET ports must propagate
  cancellation, choose dependency lifetimes, and inspect client and
  data-layer behavior.
- TypeScript types and Python hints do not validate runtime input.
  Node.js and Python services must keep blocking work off event loops
  and explicitly bound promises, tasks, workers, sockets, and pools.
- Rust memory safety does not prove transaction or distributed
  correctness. Verify cancellation safety, keep blocking work off
  async executors, and bound tasks and channels.
- Every stack needs one documented pipeline for formatting, warnings-
  as-errors compilation, unit and PostgreSQL integration tests,
  concurrency checks, fuzzing, dependency checks, contract checks,
  container smoke tests, and the required failure drills.

### Cross-cutting review questions

Before treating a capability as complete, answer all of these:

- **Contract:** What inputs, outputs, errors, ordering, compatibility,
  and ownership promises are public?
- **Data:** Which rules are authoritative, where are they enforced,
  and what happens during concurrent access?
- **Time:** Which deadlines, leases, retry windows, retention periods,
  and clock assumptions exist?
- **Resources:** Which payloads, pools, queues, workers, retries,
  caches, and telemetry dimensions have explicit bounds?
- **Failure:** What can fail partially or become ambiguous, and how
  does the system converge, compensate, or receive repair?
- **Security:** Which user or workload is the principal, what can it
  reach, and how small is credential and data exposure?
- **Operations:** Can an operator detect, diagnose, mitigate, and
  verify recovery without exposing private information?
- **Delivery:** Can old and new versions coexist, and are rollback,
  roll-forward, restore, and termination behavior proven?
- **Evidence:** Which deterministic test, measurement, experiment,
  drill, or review supports each important claim?
- **Simplicity:** Which complexity is necessary now, and what can be
  deferred or removed?

## Further study

Optional specialization is role-driven. Streaming, machine learning,
ledger, real-time media, and other specialist systems belong here only
when the work requires them. Each specialization must still define its
contract, failure model, security boundary, operating evidence, and
fit with the owning service rather than bypassing the common standard.

## 1. Go values, program structure, and ownership

### The problem it solves

Backend boundaries fail when callers and callees disagree about who
owns mutable data. A copied slice can still share an array, a map is
always a reference to shared state, and an interface can hold a typed
nil while the interface itself is not nil.

### The idea in simple terms

A Go value has a representation and an owner. Copying an array copies
its elements. A slice copy duplicates only its small descriptor. A map
or pointer copy preserves access to the same storage. Make ownership
visible at package boundaries and clone mutable data when ownership
must change.

### How it works

Modules group versioned packages; packages define the visibility
boundary. Capitalized names are exported. Structs group fields,
methods attach behavior, and interfaces describe required methods.
Zero values should be useful when practical. Use `make` for slices,
maps, and channels; use `new` only when a pointer to a zero value is
actually useful.

Strings contain bytes, so range over runes when the rule is about
Unicode code points. Use the comma-ok map form when stored zero and
missing have different meanings. Prefer value receivers for small
immutable values and pointer receivers for mutation or costly copies.

### Step-by-step examples

1. **Before:** `Product{Tags: ["tea", "sale"]}` owns one slice array.
2.  **Decision:** `Clone` allocates a new array before copying the
   tags.
3. **After:** a caller tag change does not alter the clone.
4. **Known correct:** both products retain the intended two-tag shape,
  and their first elements can differ without a hidden write.

### Why it works

The invariant is that no mutable backing storage crosses an ownership
boundary unintentionally. A deep-enough copy establishes independent
storage. Every later mutation reaches only the selected owner's data,
so the invariant is preserved for the lifetime of both values.

### Go example

```go
package catalog

type Product struct {
	SKU  string
	Tags []string
}

func Clone(product Product) Product {
	clone := product
	clone.Tags = append([]string(nil), product.Tags...)
	return clone
}

func LookupPrice(prices map[string]int, sku string) (int, bool) {
	price, found := prices[sku]
	return price, found
}
```

`Clone` covers the ownership transfer. `LookupPrice` keeps a missing
SKU distinct from a deliberately stored zero price.

### Work and memory

Cloning takes `O(t)` time and space for `t` tags. Copying the outer
struct is constant work. Map lookup is `O(1)` average time and uses no
new storage. Converting a string to `[]rune` costs `O(n)` time and
space; simple byte indexing does not.

### When to use it

Apply these rules at package, goroutine, cache, persistence, and API
boundaries. They are especially useful when collections outlive the
call that supplied them.

### When not to use it

Do not add pointers, interfaces, generics, or deep copies without an
ownership, substitution, or mutation need. Plain values are easier to
reason about when their size and lifetime are small.

### Common mistakes and edge cases

- Treating slice assignment as an element copy.
- Writing to a nil map or depending on randomized map iteration order.
- Truncating Unicode by treating a byte index as a character index.
- Returning a typed nil through an interface and comparing it to nil.
- Shadowing an outer error with `:=` and checking the wrong variable.

### Related material

See [Go language refresher][go-refresher],
[Go backend foundations][go-foundations],
[errors and cancellation](#2-go-errors-resources-and-cancellation),
and [language portability](#32-language-portability).

## 2. Go errors, resources, and cancellation

### The problem it solves

A request can fail after acquiring rows, files, timers, or remote
work. Cleanup must still happen, and callers need both useful context
and the original error identity to choose a response.

### The idea in simple terms

The layer that acquires a resource owns its release. Pass one context
down the call tree, let errors travel upward, and wrap errors without
erasing the cause that a boundary must classify.

### How it works

Check acquisition before deferring cleanup. Put `defer` immediately
after a successful acquire so later returns cannot bypass it. Accept a
`context.Context` as the first parameter of blocking work. A layer may
shorten a deadline it owns, but it must call the returned cancel
function to release timer resources.

Wrap with `%w`, add operation context, and classify with `errors.Is`,
`errors.As`, or `errors.AsType` at a policy boundary. Sentinels suit a
small stable condition set; typed errors suit structured details.

### Step-by-step examples

1. **Before:** neither a product result nor cancellation is ready.
2. **Decision:** the request context becomes ready first in `select`.
3. **After:** it returns `context.Canceled` and stops waiting.
4. **Known correct:** no result was consumed, no child timer remains,
  and the caller can still identify cancellation with `errors.Is`.

### Why it works

The invariant is that each acquired resource has exactly one owner and
one release path. Last-in-first-out defers mirror nested acquisition.
Cancellation is cooperative: each blocking layer observes the same
signal and returns. Error wrapping preserves the causal chain while
adding the operation that failed.

### Go example

```go
package catalog

import (
	"context"
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("product not found")

func Await(
	ctx context.Context,
	results <-chan string,
) (string, error) {
	select {
	case <-ctx.Done():
		return "", fmt.Errorf("await product: %w", ctx.Err())
	case product := <-results:
		if product == "" {
			return "", fmt.Errorf("await product: %w", ErrNotFound)
		}
		return product, nil
	}
}
```

The ordinary path returns a product. The failure paths retain either
absence or cancellation identity through the wrapper.

### Work and memory

The selection itself is constant work. Waiting time is bounded by the
result or context lifetime. Each wrapper and child context adds small
constant memory; held resources can be far more expensive, so release
latency belongs in the complexity discussion.

### When to use it

Use this pattern for I/O, locks, queues, timers, database operations,
and remote calls whose caller may stop waiting or classify failure.

### When not to use it

Do not add context to instant pure functions, store it in a struct, or
use it as a general parameter bag. Do not recover ordinary operational
errors as panics.

### Common mistakes and edge cases

- Wrapping with `%v`, which loses the traversable error chain.
- Deferring cleanup before checking whether acquisition succeeded.
- Ignoring row iteration, body-close, or buffered-write errors.
- Creating loop-scoped defers that retain every resource until return.
- Replacing a request context with `context.Background` mid-call.
- Returning a typed-nil error through the `error` interface.

### Related material

See [Go backend foundations][go-foundations],
[HTTP lifecycle](#6-http-contracts-and-process-lifecycle), and
[synchronous communication](#20-synchronous-service-communication).

## 3. Go concurrency and executable verification

### The problem it solves

Code can be free of data races and still violate a business rule.
Separately locked reads and writes let two goroutines observe absence
and then create the same SKU.

### The idea in simple terms

Protect the whole decision, not individual memory accesses. Give every
goroutine an owner and a bounded ending, then verify both the memory
safety and the business outcome.

### How it works

Write the invariant before choosing a primitive. Use a mutex when
several fields form one state transition. Use channels when ownership
transfer or stream coordination is the clearer model. Keep the
check-and-change inside one critical section, bound worker counts, and
join all started work.

Functional tests prove outcomes. Contention tests arrange overlap.
`go test -race` detects unsynchronized memory access. It does not
prove
uniqueness, ordering, or exactly-once business behavior. `go vet` and
targeted fuzzing cover different defect classes.

### Step-by-step examples

1. **Before:** `TEA-001` is absent; goroutines A and B are ready.
2. **Decision:** A locks, checks, inserts, and unlocks while B waits.
3. **After:** B locks, observes the stored SKU, and returns duplicate.
4.  **Known correct:** one call succeeds, one fails, and the map has
   one
  product regardless of which goroutine acquired the lock first.

### Why it works

The invariant is that a SKU is inserted at most once. It holds for the
empty store. Mutual exclusion makes check-and-insert one indivisible
transition, so a successful insert preserves uniqueness. Joining every
goroutine makes the test terminate only after all outcomes are known.

### Go example

```go
package catalog

import (
	"errors"
	"sync"
)

var ErrDuplicate = errors.New("duplicate product")

type Store struct {
	mu       sync.Mutex
	products map[string]string
}

func NewStore() *Store {
	return &Store{products: make(map[string]string)}
}

func (store *Store) Create(sku string) error {
	store.mu.Lock()
	defer store.mu.Unlock()
	if _, found := store.products[sku]; found {
		return ErrDuplicate
	}
	store.products[sku] = sku
	return nil
}
```

A contention test should release two workers from one start channel,
join them with `sync.WaitGroup`, and assert one nil and one duplicate.

### Work and memory

Map access is `O(1)` average time. Lock wait depends on contention and
critical-section duration. The store uses `O(n)` memory for `n` SKUs;
each live goroutine also consumes stack and scheduler resources, which
is why concurrency must be bounded.

### When to use it

Use it for shared in-process state, bounded worker coordination, and
characterizing an invariant before moving enforcement to PostgreSQL.

### When not to use it

An in-memory mutex cannot coordinate several processes and cannot
replace a durable transaction. Do not add goroutines to work that is
already serial, tiny, or governed by a stricter external limit.

### Common mistakes and edge cases

- Locking the lookup and insert separately.
- Starting unbounded goroutines or failing to join an error path.
- Letting receivers close a channel owned by the sender.
- Treating a buffered channel as an infinite queue.
- Using atomics for an invariant that spans several values.
- Using sleeps as synchronization or overclaiming race-detector proof.

### Related material

See [Go language refresher][go-refresher],
[deterministic testing](#10-deterministic-backend-testing), and
[contested state](#8-business-transactions-and-contested-state).

## 4. Domain discovery and modeling

### The problem it solves

Teams can implement the same product sentence consistently and still
produce incompatible behavior. Words such as placed, paid, canceled,
and refunded need agreed meaning before routes or tables freeze an
accidental interpretation.

### The idea in simple terms

Use concrete stories to discover business language, commands, events,
entities, value objects, aggregates, policies, and external systems.
Model only distinctions that change behavior, and make every allowed
state transition explicit.

### How it works

Start with a successful story and then add rejection, concurrency, and
external-failure stories. Name commands as requested actions and
events
as facts that already happened. Put identity and lifecycle on
entities;
make descriptive values immutable. An aggregate is the smallest unit
that must change atomically to preserve one invariant.

Record assumptions and unresolved terms. A model is useful when a
product reviewer can decide whether each transition is legal without
reading HTTP, SQL, or framework code.

### Step-by-step examples

1. **Before:** order `O-42` is placed for customer `C-17`.
2. **Decision:** a cancel command checks ownership and current state.
3. **After:** the order becomes canceled and emits `OrderCanceled`.
4.  **Known correct:** another customer is rejected, and a shipped
   order
  cannot move backward to canceled.

### Why it works

The invariant is that only the owning customer may perform a legal
order transition. Every command names its actor, target, and expected
state. The transition table covers every current state and reaches a
terminal accepted or rejected result, so there is no implicit path.

### Go example

```go
package order

import "errors"

var ErrCannotCancel = errors.New("order cannot be canceled")

type Status string

const (
	Placed   Status = "placed"
	Canceled Status = "canceled"
)

type Order struct {
	CustomerID string
	Status     Status
}

func (order *Order) Cancel(customerID string) error {
	if order.CustomerID != customerID || order.Status != Placed {
		return ErrCannotCancel
	}
	order.Status = Canceled
	return nil
}
```

The ordinary case performs one legal transition. Wrong ownership or
state returns a stable domain error and leaves the order unchanged.

### Work and memory

A direct state transition is constant time and space. Discovery work
grows with the number of distinct stories and stakeholders, not code
size. Event histories and large rule tables require explicit bounds
and indexes when they become runtime data.

### When to use it

Use this process when requirements contain lifecycle words, contested
resources, permissions, money, or coordination with another system.

### When not to use it

Do not invent aggregates, events, or repositories for a mechanical
transformation with no durable business rule. A plain function is
often the honest model.

### Common mistakes and edge cases

- Turning every noun into an entity or every method into a service.
- Naming database writes as events before the business action commits.
- Letting an aggregate span data that cannot change atomically.
- Modeling only the happy path and leaving rejection semantics hidden.
- Treating a diagram as final evidence instead of testing real
  stories.

### Related material

See [domain discovery][domain-reference],
[service construction](#5-architecture-and-service-construction),
and [distributed
workflows](#21-distributed-workflows-and-eventual-consistency).

<a id="2-data-transactions-and-architecture"></a>

> The former combined data and architecture lesson now begins here
>  and continues through [relational
  data](#7-relational-modeling-and-postgresql-access)
> and [transactions](#8-business-transactions-and-contested-state).

## 5. Architecture and service construction

### The problem it solves

Transport, business policy, and storage change for different reasons.
When they are mixed in handlers or repository-shaped services, tests
require infrastructure and a small contract change spreads everywhere.

### The idea in simple terms

Dependencies point inward toward stable policy. Handlers adapt HTTP to
commands, use cases own business orchestration, repositories adapt
persistence, and the composition root constructs concrete
dependencies.

### How it works

Define a narrow use-case input around business intent. Keep domain
types unaware of JSON and SQL. Declare an interface only at the caller
that needs substitution. Make each transaction boundary match one use
case, and wire concrete implementations in `main` or another explicit
composition root.

A modular monolith is the default: one deployment, clear internal
packages, and database-enforced invariants. Extract a service only
when
an independently owned boundary has evidence that repays distributed
cost.

### Step-by-step examples

1. **Before:** an HTTP body requests placement of order `O-42`.
2. **Decision:** the handler validates shape and calls `PlaceOrder`.
3. **After:** the use case applies policy and asks storage to persist.
4. **Known correct:** invalid input never calls storage, and the same
  use
  case can be invoked by a job without importing HTTP types.

### Why it works

The invariant is that business decisions depend only on inward-facing
contracts. Adapters translate at the edge and the composition root is
the only place that knows the complete graph. Each dependency can
therefore change or be replaced without reversing the policy boundary.

### Go example

```go
package orders

import (
	"context"
	"errors"
)

var ErrInvalidQuantity = errors.New("quantity must be positive")

type Store interface {
	Save(context.Context, Order) error
}

type Order struct {
	ID       string
	Quantity int
}

type Service struct {
	store Store
}

func NewService(store Store) Service {
	return Service{store: store}
}

func (service Service) Place(
	ctx context.Context,
	order Order,
) error {
	if order.Quantity <= 0 {
		return ErrInvalidQuantity
	}
	return service.store.Save(ctx, order)
}
```

The interface belongs to the use case. A rejected order does not reach
the store; a valid order delegates exactly once.

### Work and memory

Layering adds constant call and translation overhead. Total work still
comes from validation, database access, and external calls. Excessive
mapping and broad interfaces add allocations and change surface
without
improving isolation.

### When to use it

Use these boundaries for services with durable policy, multiple entry
points, or infrastructure that changes independently of business
rules.

### When not to use it

Do not add an interface for every struct or split one deployable into
network services merely to imitate an architecture diagram.

### Common mistakes and edge cases

- Defining repository interfaces in the implementation package.
- Allowing handlers to begin transactions or assemble domain policy.
- Hiding dependency construction in globals or package initialization.
- Returning database records directly as public transport contracts.
- Calling a package modular while it mutates another package's tables.

### Related material

See [service core][service-core],
[HTTP lifecycle](#6-http-contracts-and-process-lifecycle),
[relational data](#7-relational-modeling-and-postgresql-access), and
[service extraction](#19-service-boundary-discovery-and-extraction).

<a id="1-http-api-contracts"></a>

## 6. HTTP contracts and process lifecycle

### The problem it solves

A browser and server need one observable agreement about requests,
responses, validation, identity, caching, cancellation, and failure.
Without it, clients guess and servers expose accidental behavior.

### The idea in simple terms

Treat an API as a public promise, not a route that happens to decode
JSON. HTTP supplies method, target, headers, representation, status,
and caching semantics. The application contract narrows those choices
for one operation.

### How it works

1. Define the operation's intent and whether it is safe or idempotent.
2. Bound the request body and processing time.
3. Verify media type, identity, authorization, and syntax.
4. Translate transport data into an application command.
5. Let the use case decide business outcomes.
6. Map known outcomes to stable status and problem codes.
7. Return one documented representation and correlation identity.
8. Propagate cancellation and release every owned resource.
9. Start dependencies before advertising readiness.
10.  On shutdown, stop admission, drain bounded work, then close
    owners.

### Step-by-step examples

Input: `POST /v1/orders` with valid JSON and operation key `K1`.
Expected output: one order response or one stable documented error.

1. **Before:** The server has untrusted bytes and headers.
   **Decision:** Enforce the body limit and expected media type.
   **After:** Oversized or incompatible input is rejected early.
   **Known correct:** Later decoding cannot exceed the stated bound.
2. **Before:** JSON syntax is valid. **Decision:** Validate required
   fields and caller authorization. **After:** One typed command
   reaches the use case. **Known correct:** Transport and access
   failures cannot enter logic.
3. **Before:** The use case returns a known outcome. **Decision:** Map
   it to one status and stable machine code. **After:** The client can
   choose a documented next action. **Known correct:** Internal error
   details remain private.

Failure example: a second JSON value after the first is rejected
rather than silently ignored.

Process trace: before a termination signal, readiness is true and
requests may enter. The signal makes readiness false and begins a
bounded server shutdown. After handlers drain or the deadline expires,
owned dependencies close. The process is known correct when no new
request was admitted during the drain and listener failure reached the
supervising goroutine.

### Why it works

Each boundary narrows possible state before handing data inward.
Stable error identity lets one boundary translate an outcome without
parsing text. Bounded input, time, and resources guarantee that one
request cannot consume unlimited process capacity.

### Go example

This complete package fragment owns strict decoding for one request.

```go
package requestjson

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func Decode(
	w http.ResponseWriter,
	r *http.Request,
	destination any,
) error {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(destination); err != nil {
		return err
	}
	if err := decoder.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		return errors.New("body must contain one JSON value")
	}
	return nil
}
```

### Work and memory

Decoding visits the input bytes, so work grows with body size. The
decoded value and decoder buffers consume memory based on that bounded
input. Network time, authentication lookup, database work, and output
serialization are additional dimensions and need separate limits.

### When to use it

Use an explicit contract for every browser-facing or service-facing
operation, especially authenticated writes, pagination, uploads, and
work that clients may retry.

### When not to use it

Do not expose a generic persistence interface as an API contract. Do
not use JSON request-response semantics for a sustained bidirectional
stream when a streaming protocol is the actual requirement.

### Common mistakes and edge cases

> **Warning:** CORS allows selected browser origins to read a
> response. It does not authenticate or authorize the caller.

- A `200` status with an error-shaped body creates two contracts.
- Retrying writes without idempotency can duplicate state.
- Liveness checks must not depend on external services.
- Offset pagination becomes unstable under concurrent insertion.
- Logging raw bodies or credentials can leak private data.
- Starting `ListenAndServe` in an unobserved goroutine hides failure.
- Unlimited graceful shutdown can prevent deployment completion.

### Related material

- [Architecture and service construction](
  #5-architecture-and-service-construction)
- [Frontend network and security][frontend-network]
- [Data, transactions, and architecture][data-lesson]

[frontend-network]:
  ../frontend/CONCEPTS.md#a3-networking-storage-and-browser-security
[data-lesson]: #8-business-transactions-and-contested-state
[production-lesson]: #3-production-and-distributed-systems

## 7. Relational modeling and PostgreSQL access

### The problem it solves

Application values must survive restarts, concurrent readers, schema
change, and operational recovery. A convenient object shape is not
automatically a relational model with enforceable meaning.

### The idea in simple terms

Model durable facts with tables, keys, constraints, and indexes. Keep
queries explicit, bind values as parameters, propagate context, and
own every row iterator and transaction resource you acquire.

### How it works

Choose primary keys for stable identity and foreign keys for valid
relationships. Use `NOT NULL`, checks, and unique constraints for
facts
the database can enforce. Normalize update-sensitive facts, then
denormalize only with measured read evidence and a repair strategy.

Select only required columns and scan them in one reviewed order. Use
`QueryRowContext` for one row and `QueryContext` for many. Close rows,
check iteration errors, distinguish no row from failure, and design
indexes from actual predicates and ordering. Migrations must be safe
for the old and new application during rollout.

### Step-by-step examples

1. **Before:** SKU `TEA-001` is absent and an insert is parameterized.
2. **Decision:** PostgreSQL checks `NOT NULL` and unique SKU
  constraints.
3. **After:** one row exists, or the duplicate maps to a stable
  conflict.
4.  **Known correct:** malformed data cannot bypass the same
   constraints
  through another process, and values never become SQL syntax.

### Why it works

The invariant is that every stored product has required fields and one
unique SKU. Constraints run in the same atomic statement as the write,
so every writer preserves the rule. Parameter binding separates data
from query structure, and explicit cleanup returns scarce connections.

### Go example

```go
package catalog

import (
	"context"
	"database/sql"
)

type Product struct {
	SKU  string
	Name string
}

type RowScanner interface {
	Scan(...any) error
}

func ScanProduct(row RowScanner) (Product, error) {
	var product Product
	err := row.Scan(&product.SKU, &product.Name)
	return product, err
}

func Find(
	ctx context.Context,
	database *sql.DB,
	sku string,
) (Product, error) {
	const query = `SELECT sku, name FROM products WHERE sku = $1`
	row := database.QueryRowContext(ctx, query, sku)
	return ScanProduct(row)
}
```

The bound parameter is ordinary data. The caller must classify
`sql.ErrNoRows` separately from an unavailable or canceled database.

### Work and memory

With a selective B-tree index, lookup is commonly `O(log n)` plus row
access. Without it, a scan is `O(n)`. Result memory grows with rows
read;
streaming still holds a connection, so result size and time need
bounds.
Every index consumes write work, storage, cache, and vacuum effort.

### When to use it

Use relational storage for durable, constrained facts with joins,
transactions, and operational recovery requirements.

### When not to use it

Do not add an index for every column, open a transaction for one
atomic
statement, or store an opaque document when individual fields govern
integrity and access paths.

### Common mistakes and edge cases

- Treating application validation as a substitute for constraints.
- Using `SELECT *` and silently coupling scans to schema order.
- Forgetting `rows.Close` or `rows.Err` on streamed results.
- Concatenating values into SQL or logging sensitive parameters.
- Shipping a destructive migration before all readers stop using data.
- Ignoring pool saturation, lock waits, vacuum, and restore testing.

### Related material

See [PostgreSQL and relational data][postgres-reference],
[contested state](#8-business-transactions-and-contested-state), and
[advanced PostgreSQL](#15-advanced-postgresql-operation-and-recovery).

## 8. Business transactions and contested state

### The problem it solves

Business actions often change several records under concurrent access.
The system must preserve rules such as non-negative inventory even
when requests overlap or one step fails.

### The idea in simple terms

Put one business action behind one use-case boundary. A transaction
makes its required data changes succeed or fail together. Locking or
an atomic condition prevents concurrent actions from spending the same
resource.

### How it works

1. Write the business rule independently of transport and storage.
2. Validate values that do not require database state.
3. Begin a transaction at the use-case boundary.
4. Read or update contested records through one consistent strategy.
5. Check affected rows and translate expected conflicts.
6. Write every dependent record in the same transaction.
7. Commit once the invariant holds or roll back every partial change.
8. Keep HTTP and database representations outside the domain rule.

### Step-by-step examples

Input: inventory `1`; requests A and B each reserve `1`. Expected
output: one success, one insufficient-inventory result, final stock
`0`.

1. **Before:** Stock is `1` and neither request owns it. **Decision:**
   A performs an atomic decrement conditioned on stock at least `1`.
   **After:** A changes stock to `0` and owns the reservation. **Known
   correct:** Stock has never been negative.
2. **Before:** B evaluates the same condition against current stock
   `0`. **Decision:** Its conditional update affects no row.
   **After:** B returns the expected business conflict. **Known
   correct:** No second reservation exists.
3. **Before:** A writes its order and commits. **Decision:** Publish
   success only after commit. **After:** Order and inventory become
   visible together. **Known correct:** The database contains no
   partial successful order.

Failure example: if order insertion fails after the decrement,
rollback restores inventory to `1`.

### Why it works

The conditional write makes the inventory rule part of the same atomic
database decision as the change. The transaction hides partial state
and either publishes both records or neither. The use case owns this
boundary because repositories cannot coordinate a business action by
committing independently.

### Go example

This in-memory fragment demonstrates the same invariant. The database
lesson replaces the mutex with a transaction and conditional SQL.

```go
package inventory

import (
	"errors"
	"sync"
)

var ErrInsufficient = errors.New("insufficient inventory")

type Store struct {
	mu       sync.Mutex
	quantity int
}

func NewStore(quantity int) *Store {
	return &Store{quantity: quantity}
}

func (store *Store) Reserve(quantity int) error {
	if quantity <= 0 {
		return errors.New("quantity must be positive")
	}
	store.mu.Lock()
	defer store.mu.Unlock()
	if store.quantity < quantity {
		return ErrInsufficient
	}
	store.quantity -= quantity
	return nil
}
```

### Work and memory

The in-memory operation does constant work and retains constant extra
memory. A database transaction also pays connection, lock-wait, log,
index, and network costs. Sorting `n` lock targets costs `O(n log n)`
but can give concurrent transactions one consistent lock order.

### When to use it

Use a transaction when one business invariant spans multiple reads or
writes. Use conditional writes or locks for inventory, balances,
quotas, unique allocation, and state-machine transitions.

### When not to use it

Do not hold a database transaction open across a slow external network
call. Do not add repository interfaces without a real boundary. For
cross-service actions, use durable workflow and compensation instead
of pretending one local transaction covers every service.

### Common mistakes and edge cases

> **Warning:** Reading stock and decrementing it in separate unlocked
> statements lets two requests both approve the last unit.

- Application validation cannot replace database constraints.
- Retrying a whole transaction needs bounded policy and idempotency.
- Locking resources in different orders can deadlock.
- A successful commit can still have an uncertain client outcome.
- Long transactions retain locks and old row versions.

### Related material

- [Relational modeling](#7-relational-modeling-and-postgresql-access)
- [Architecture and service construction](
  #5-architecture-and-service-construction)
- [Production and distributed systems][production-lesson]

## 9. Identity, authorization, and API security

### The problem it solves

An API must know who is acting, whether that actor may perform this
operation on this object, and whether untrusted input can escape its
intended data boundary.

### The idea in simple terms

Authentication establishes an identity. Authorization evaluates that
identity, action, resource, and context for every protected operation.
Validation, output encoding, secret handling, and browser policy form
additional independent controls.

### How it works

Verify credentials with an established library, normalize one
principal,
and place only that identity in request context. Enforce object-level
authorization in the use case, close to the loaded resource. Prefer
short-lived credentials, explicit revocation strategy, least
privilege,
parameterized SQL, bounded input, and redacted logs. Cookies carrying
authority need `Secure`, `HttpOnly`, suitable `SameSite`, and CSRF
protection where browser ambient authority applies.

### Step-by-step examples

1. **Before:** user `U-7` requests order `O-42`, owned by `U-9`.
2. **Decision:** the use case compares the principal with resource
  owner.
3. **After:** it returns a stable forbidden result without mutation.
4. **Known correct:** changing the URL identifier never grants access,
  and logs contain identifiers rather than credentials or raw tokens.

### Why it works

The invariant is that every protected state change has an
authenticated
principal and a positive resource-specific authorization decision.
Enforcement after resource lookup prevents identifier substitution,
and
independent input and output controls limit injection and disclosure.

### Go example

See the typed principal and role guard in
[the service examples][service-examples]. Its
[tests][service-tests] exercise an allowed administrator, a forbidden
role, and a missing principal.


### Work and memory

Token verification is usually linear in credential size plus constant
cryptographic work. Resource authorization adds its lookup or join.
Password hashing is intentionally expensive and must have bounded
concurrency so it cannot exhaust CPU or memory.

### When to use it

Use layered controls for every non-public operation and every boundary
that accepts browser, partner, file, webhook, or database input.

### When not to use it

Do not invent cryptography, treat CORS as access control, or authorize
only at route level when ownership depends on the loaded object.

### Common mistakes and edge cases

- Accepting an algorithm or key choice directly from an untrusted
  token.
- Omitting audience, issuer, expiry, clock-skew, or revocation checks.
- Logging secrets, reset links, session cookies, or full request
  bodies.
- Returning different details that reveal whether a private object
  exists.
- Trusting a proxy identity header without a verified trust boundary.

### Related material

See [identity and API security][security-reference],
[privacy and supply
chain](#25-security-privacy-and-software-supply-chain),
and [production integrations](#29-production-integration-boundaries).

## 10. Deterministic backend testing

### The problem it solves

Tests become slow or flaky when time, randomness, networks, global
state,
and goroutine ordering are uncontrolled. Passing examples then provide
little evidence about failures or concurrent invariants.

### The idea in simple terms

Test through the smallest public boundary that proves the behavior.
Inject nondeterministic inputs, arrange concurrency with signals, and
separate fast policy tests from realistic integration and contract
tests.

### How it works

Use table-driven tests for input classes and explicit expected
outcomes.
Replace wall-clock reads, random identifiers, and remote calls through
narrow caller-owned seams. Use `httptest` for HTTP, real PostgreSQL
for
query and transaction semantics, and fakes only when their behavior is
smaller than the real dependency. Gate concurrent work with channels
or
barriers, join it, and run the race detector separately from
assertions.

### Step-by-step examples

1.  **Before:** two reservation workers are blocked on one start
   signal.
2. **Decision:** the test releases both, waits for both, and sorts
  results.
3. **After:** exactly one succeeds and final inventory is zero.
4. **Known correct:** no sleep or scheduler guess affects the
  assertion;
  a deliberately injected storage error also leaves inventory
  unchanged.

### Why it works

The invariant is that test inputs fully determine observable results.
Explicit signals define the relevant happens-before edges, while
public
outcome assertions remain independent of internal scheduling. Each
test
terminates through a join or bounded context.

### Go example

See the deterministic placement service in
[the service examples][service-examples]. Its
[tests][service-tests] inject time and identity, then prove invalid
input makes no repository call.


### Work and memory

Table tests cost the sum of their cases. Integration setup often
dominates
runtime. Fuzzing and race detection multiply execution and memory, so
CI
needs explicit time, worker, corpus, and package bounds.

### When to use it

Use deterministic seams for policy, error mapping, deadlines, retries,
state transitions, and any concurrency regression that must be
replayed.

### When not to use it

Do not mock a database query whose correctness depends on PostgreSQL,
or
assert private call order when the public outcome is the real
contract.

### Common mistakes and edge cases

- Sleeping until work is probably complete.
- Sharing mutable fixtures across parallel tests.
- Treating race-detector silence as a business-invariant proof.
- Building a fake more complex than the production dependency
  contract.
- Omitting cleanup, failure injection, cancellation, and retry cases.

### Related material

See [testing and operations][testing-reference],
[Go concurrency](#3-go-concurrency-and-executable-verification), and
[resilience experiments][resilience-experiments].

<a id="3-production-and-distributed-systems"></a>

> The former production lesson begins here and continues through the
> reliability, delivery, messaging, and distributed-system entries.

## 11. Observability and operational controls

### The problem it solves

Operators need to distinguish invalid requests, application defects,
dependency failure, saturation, and slow work without reading private
payloads or reproducing production locally.

### The idea in simple terms

Logs describe events, metrics summarize behavior, and traces connect
one
request across boundaries. Health endpoints, timeouts, rate limits,
and
profiles turn that evidence into safe operating controls.

### How it works

Define user-visible service indicators first. Emit structured logs
with
stable event names and correlation fields. Use bounded-cardinality
metric
labels, propagate trace context, and record spans around meaningful
remote
or queued work. Liveness answers whether the process should restart;
readiness answers whether it should receive new traffic. Expose
diagnostic
profiles only on authenticated, restricted interfaces.

### Step-by-step examples

1. **Before:** request `R-8` enters while database latency rises.
2.  **Decision:** the handler span records the query child and
   deadline.
3. **After:** the request returns unavailable and increments one
  counter.
4. **Known correct:** logs, metric, and trace share `R-8`; no body,
  token,
  or unbounded customer value became a metric label.

### Why it works

The invariant is that each important operation produces enough stable,
non-secret evidence to classify its outcome. Shared correlation
preserves
causal context, and bounded dimensions keep the telemetry system
operable
under the same failure it is meant to explain.

### Go example

See the bounded access-event builder in
[the service examples][service-examples]. Its
[tests][service-tests] accept stable public fields and reject an
authorization field before it reaches telemetry.


### Work and memory

Each signal adds serialization, allocation, and export work.
Histograms,
span volume, and label combinations can grow rapidly. Sample
intentionally,
bound queues, and ensure telemetry backpressure cannot stop request
work.

### When to use it

Use these controls for every production service, worker, scheduler,
and
integration boundary with an operator or reliability objective.

### When not to use it

Do not log every payload, create per-user metric labels, or make
liveness
depend on a remote service whose outage a restart cannot repair.

### Common mistakes and edge cases

- Logging an error without operation, outcome, or correlation fields.
- Reporting ready before migrations or required local initialization.
- Hiding listener or exporter failures in abandoned goroutines.
- Allowing telemetry queues or profiles to consume unbounded
  resources.
- Alerting on implementation noise instead of user-visible symptoms.

### Related material

See [testing and operations][testing-reference],
[delivery and
recovery](#12-delivery-deployment-and-recovery-readiness),
and [reliability
engineering](#28-reliability-engineering-incidents-and-cost).

## 12. Delivery, deployment, and recovery readiness

### The problem it solves

A locally correct binary can fail because its artifact is
unrepeatable,
configuration is unsafe, a migration is incompatible, rollout lacks a
stop signal, or backups cannot actually restore service.

### The idea in simple terms

Build once, promote the same immutable artifact, keep secrets outside
it,
deploy through compatibility stages, observe explicit health criteria,
and rehearse rollback and data recovery before an incident.

### How it works

Use reproducible builds, pinned dependencies, least-privilege runtime
images, and generated provenance. Validate configuration at startup
and
redact secrets. Apply expand-migrate-contract schema changes. Roll out
to
a small population, compare service indicators, pause on threshold
breach,
and keep rollback compatible. Define recovery point and recovery time
objectives, then prove them with restore exercises.

### Step-by-step examples

1.  **Before:** version B serves 5% of traffic beside healthy version
   A.
2. **Decision:** B exceeds the error threshold during the observation
  window.
3. **After:** admission to B stops and traffic returns to A.
4. **Known correct:** the schema remains backward compatible, the same
  artifact was tested, and the rollback condition is recorded.

### Why it works

The invariant is that every deployment step has an observable
acceptance
condition and a compatible reversal. Immutable artifacts remove
rebuild
variance, staged schemas preserve mixed-version operation, and restore
tests replace assumptions about recoverability with measured evidence.

### Go example

See the typed configuration parser in
[the service examples][service-examples]. Its
[tests][service-tests] cover valid startup, missing required values,
and secret-safe failure text.


### Work and memory

Build and scan cost scales with dependency and artifact size. Canary
time
depends on traffic needed for a useful signal. Backups cost storage
and I/O;
restore duration must fit the recovery objective under realistic
volume.

### When to use it

Use this discipline for every deployable artifact and every persistent
schema, queue, secret, or stateful dependency it changes.

### When not to use it

Do not rebuild per environment, hide configuration defaults, or claim
a
rollback that cannot read the newly written schema or data
representation.

### Common mistakes and edge cases

- Deploying destructive schema changes with the first application
  version.
- Using floating image tags or unverified build inputs.
- Making readiness true before dependencies and listeners are usable.
- Rotating a secret without overlap for in-flight credentials.
- Calling a backup successful without a timed restore and integrity
  check.

### Related material

See [delivery readiness][delivery-reference],
[orchestration](#27-orchestration-and-platform-operations), and
[incidents and cost](#28-reliability-engineering-incidents-and-cost).

## 13. Idempotency and ambiguous write outcomes

### The problem it solves

A client can time out after a server commits. Retrying may duplicate a
charge or order, while refusing to retry may hide a successful result.

### The idea in simple terms

Give one logical write a stable key. Atomically bind that key to the
request fingerprint and durable outcome, then replay the same outcome
or
reject conflicting reuse.

### How it works

Scope keys by tenant and operation. Canonicalize and hash the relevant
request. In one transaction, claim a unique key, store processing
state,
perform the write, and persist the terminal response reference. A
duplicate
with the same hash reads or waits for that record; a different hash is
a
conflict. Retention must exceed the documented retry window.

### Step-by-step examples

1. **Before:** key `K1` has no record and create-order begins.
2. **Decision:** the transaction claims `K1`, creates `O-42`, and
  commits.
3.  **After:** the response is lost; a retry finds the completed
   record.
4. **Known correct:** the retry returns `O-42`, while changed input
  with
  `K1` is rejected and no second order exists.

### Why it works

The invariant is one durable semantic outcome per scoped key and
request
hash. A uniqueness constraint serializes competing claims, and
recording
the result in the same transaction prevents a committed action with an
unclaimed key.

### Go example

See canonical command fingerprinting in
[the production examples][production-examples]. Its
[tests][production-tests] distinguish a new operation, a compatible
replay, and conflicting key reuse.


### Work and memory

Each logical write adds a key lookup and durable record. Indexed
access is
commonly `O(log n)`. Storage grows with request volume times
retention;
cleanup needs bounded batches and must not race the supported retry
window.

### When to use it

Use it for retried non-idempotent writes, payment-like effects, and
any API
where a committed response can be lost.

### When not to use it

Do not treat in-memory locks, request IDs, or a read-before-write
check as
durable idempotency across processes and restarts.

### Common mistakes and edge cases

- Reusing keys across tenants or operations.
- Ignoring changed request bodies under an existing key.
- Recording completion in a transaction separate from the business
  write.
- Expiring records before clients stop retrying.
- Claiming exactly-once delivery rather than one semantic effect.

### Related material

See [transactions](#8-business-transactions-and-contested-state),
[reliable events](#17-reliable-event-publication-and-consumption), and
[distributed mechanics](#18-distributed-systems-mechanics).

## 14. Performance, capacity, retry, and overload

### The problem it solves

A service that is fast at average load can collapse when queues grow,
dependencies slow, retries multiply, or one request consumes unbounded
work.

### The idea in simple terms

Measure before optimizing, budget end-to-end latency, bound every
resource,
and reject excess work early enough that admitted work can still
finish.

### How it works

Profile CPU, allocation, blocking, queries, and network calls with
realistic
traffic. Derive capacity from arrival rate, service time, concurrency,
and
tail latency. Give downstream calls shorter deadlines than the caller,
retry only transient idempotent operations with exponential backoff
and
jitter, cap attempts, and share a retry budget. Use semaphores, queue
bounds,
rate limits, and circuit breakers as admission controls with explicit
scope.

### Step-by-step examples

1. **Before:** ten worker slots are full and the bounded queue holds
  twenty.
2. **Decision:** a new request is rejected with a retryable overload
  result.
3. **After:** no extra goroutine or queue node is created.
4. **Known correct:** admitted work keeps its latency budget and
  retries cannot
  create more offered load than the configured retry budget.

### Why it works

The invariant is that in-flight and queued work never exceed proved
capacity.
Bounds preserve memory and scheduling headroom. Deadlines stop
obsolete work,
and randomized bounded retries prevent synchronized amplification.

### Go example

See the bounded retry-delay policy in
[the production examples][production-examples]. Its
[tests][production-tests] cover an allowed retry and an exit when the
remaining deadline cannot fund another call.


### Work and memory

Queue memory is `O(q)` for configured bound `q`; concurrency is `O(w)`
for
worker bound `w`. A retry policy multiplies worst-case calls by its
attempt
cap, so both latency and downstream load must include that factor.

### When to use it

Use it for every remote dependency, worker pool, expensive endpoint,
and
service with a latency or availability objective.

### When not to use it

Do not optimize from intuition, retry permanent errors, or use an
unbounded
queue to hide insufficient capacity.

### Common mistakes and edge cases

- Measuring only averages and missing tail latency or saturation.
- Setting independent deadlines whose sum exceeds the caller budget.
- Retrying at several layers and creating multiplicative attempts.
- Opening a circuit globally when failures are tenant or dependency
  scoped.
- Rejecting after expensive parsing or allocation has already
  occurred.

### Related material

See [delivery readiness][delivery-reference],
[synchronous communication](#20-synchronous-service-communication),
and
[reliability
engineering](#28-reliability-engineering-incidents-and-cost).

## 15. Advanced PostgreSQL operation and recovery

### The problem it solves

Correct queries still fail operationally through pool exhaustion,
deadlocks,
table bloat, replication lag, unsafe migrations, or untested recovery.

### The idea in simple terms

Treat connections, locks, row versions, replicas, migrations, backups,
and
restore time as finite operational resources with measured owners.

### How it works

Size the pool against database capacity across all instances. Observe
wait
time separately from query time. Keep transactions short, lock
resources in
one order, retry whole deadlocked transactions only when safe, and
monitor
vacuum and long snapshots. Route replica reads only when stale data is
valid.
Use expand-migrate-contract changes, checksum backups, and rehearse
point-in-
time recovery to a stated objective.

### Step-by-step examples

1. **Before:** transaction A locks account 1; B locks account 2.
2. **Decision:** both operations instead sort identifiers before
  locking.
3.  **After:** one waits for account 1 rather than forming a lock
   cycle.
4. **Known correct:** balances remain atomic; an injected deadlock
  retries the
  complete idempotent transaction within a bounded attempt count.

### Why it works

The invariant is a consistent lock acquisition order and one atomic
retry
unit. A total order removes cyclic wait. Short lifetimes release locks
and old
versions, while verified restore procedures bound data loss and
recovery time.

### Go example

See the configured PostgreSQL retry classifier in
[the production examples][production-examples]. Its
[tests][production-tests] accept serialization failure and reject a
permanent constraint error.


### Work and memory

Lock sorting costs `O(n log n)`. Pool memory and database sessions
scale with
the configured maximum across replicas. Indexes and retained row
versions add
storage, write amplification, cache pressure, and vacuum work.

### When to use it

Use it once PostgreSQL is a production dependency with concurrency,
schema
evolution, replicas, or recovery objectives.

### When not to use it

Do not increase pool size to mask slow queries or route
consistency-sensitive
reads to a lagging replica.

### Common mistakes and edge cases

- Multiplying a safe per-process pool by too many application
  instances.
- Retrying only the final statement after a transaction aborts.
- Holding transactions open during remote calls or user think time.
- Running blocking migrations without lock and duration analysis.
- Assuming a backup is usable without restoring and validating it.

### Related material

See [relational data](#7-relational-modeling-and-postgresql-access),
[transactions](#8-business-transactions-and-contested-state), and
[delivery and
recovery](#12-delivery-deployment-and-recovery-readiness).

## 16. Durable background jobs and scheduling

### The problem it solves

In-process goroutines disappear on restart and cannot safely
coordinate work
across replicas. Jobs may be leased twice, abandoned, or retried
forever.

### The idea in simple terms

Persist work before acknowledging it. Workers lease bounded batches,
renew
ownership, make effects idempotent, record attempts, and move poison
work to
a reviewed terminal state.

### How it works

Define job identity, payload version, availability time, attempt
limit, and
lease expiry. Claim with an atomic conditional update or skip-locked
query.
Heartbeat long work, complete only while holding the lease, and use
backoff
with jitter for transient failure. Schedulers create durable jobs from
a
stored schedule and stable occurrence key so failover does not
duplicate the
semantic run.

### Step-by-step examples

1. **Before:** job `J7` is ready with no valid lease.
2. **Decision:** worker A atomically leases it until time `T`.
3. **After:** A crashes; after `T`, worker B claims and completes it.
4.  **Known correct:** the effect key prevents duplication, attempts
   are
  bounded, and a permanent failure reaches the dead-letter state.

### Why it works

The invariant is that only the current unexpired lease may finalize a
job,
while the semantic effect is idempotent across leases. Expiry
guarantees
abandoned work becomes eligible, and an attempt limit guarantees
termination.

### Go example

See the fenced job transition in
[the production examples][production-examples]. Its
[tests][production-tests] reject a stale claimant, accept the current
claim, and terminate exhausted retryable work.


### Work and memory

Claiming an indexed ready job is commonly `O(log n)`. Live work is
bounded by
worker count and batch size. History grows with retention; heartbeat
frequency
trades recovery delay against write load.

### When to use it

Use durable jobs for slow, retryable, scheduled, or fan-out work that
must
survive process and deployment failure.

### When not to use it

Do not enqueue work whose required state has not committed, or use a
durable
queue for trivial synchronous work whose result the caller requires
now.

### Common mistakes and edge cases

- Acknowledging before the effect is durable.
- Using a lease without checking ownership during completion.
- Retrying validation failures forever.
- Omitting payload versioning and deploy-time compatibility.
- Running unbounded workers against a slower dependency.

### Related material

See [reliable events](#17-reliable-event-publication-and-consumption),
[idempotency](#13-idempotency-and-ambiguous-write-outcomes), and
[orchestration](#27-orchestration-and-platform-operations).

## 17. Reliable event publication and consumption

### The problem it solves

Writing business state and publishing a message as separate operations
can
lose the event or publish an event for a transaction that later rolls
back.
Consumers also receive duplicates and out-of-order messages.

### The idea in simple terms

Write the business change and an outbox record in one local
transaction.
Publish the outbox asynchronously, and make consumers deduplicate
effects by
stable event identity while tolerating reordering.

### How it works

Store event ID, aggregate ID, sequence, type, version, payload, and
creation
time with the business write. A relay claims unpublished rows,
publishes, and
marks progress; a crash may publish twice. Consumers insert an inbox
or effect
key in the same transaction as their state change. Partition by
aggregate
when per-aggregate order matters, and evolve payloads additively.

### Step-by-step examples

1. **Before:** order `O-42` and outbox event `E9` are uncommitted.
2. **Decision:** one transaction commits both; the relay publishes
  `E9` twice.
3. **After:** the consumer's unique inbox key applies the effect once.
4. **Known correct:** no committed order lacks an event, rollback
  publishes
  nothing, and duplicate delivery cannot duplicate consumer state.

### Why it works

The producer invariant is that business state and publication intent
commit
together. The consumer invariant is one local effect per event ID.
Local
transactions and unique constraints enforce both despite relay
crashes.

### Go example

See the idempotent consumer transition in
[the production examples][production-examples]. Its
[tests][production-tests] apply the first event and ignore a duplicate
without repeating the business effect.


### Work and memory

Each event adds producer and consumer durable writes plus broker
traffic.
Indexed claims are commonly `O(log n)`. Outbox, inbox, and broker
retention
need bounded cleanup that preserves replay and audit windows.

### When to use it

Use it when another component must eventually observe a committed
local fact
without a distributed transaction.

### When not to use it

Do not publish speculative commands as committed facts or promise
global
ordering when only per-key partition order exists.

### Common mistakes and edge cases

- Marking an outbox row sent before broker acknowledgement.
- Performing consumer effects outside the inbox transaction.
- Deduplicating only in process memory.
- Removing or reinterpreting fields in an existing event version.
- Assuming broker acknowledgement means every consumer applied the
  event.

### Related material

See [distributed foundations][distributed-reference],
[durable jobs](#16-durable-background-jobs-and-scheduling), and
[contract evolution](#22-contract-evolution-and-independent-delivery).

## 18. Distributed-systems mechanics

### The problem it solves

Across a network, delay is indistinguishable from failure, clocks
disagree,
messages repeat, and replicas observe different histories. Local
reasoning no
longer proves a global outcome.

### The idea in simple terms

State the consistency and availability requirement per operation. Use
stable
identity, monotonic versions, quorum or leader rules, and
reconciliation
instead of relying on timing or a mythical exactly-once network.

### How it works

Model failure domains and the result of lost requests, replies, and
leader
changes. A leader can serialize one log but needs fencing after lease
loss.
Quorums intersect only under their stated replica assumptions. Logical
or
hybrid clocks express causal order better than wall time.
At-least-once
delivery needs idempotent effects; eventually consistent state needs a
merge
rule, version check, or explicit conflict workflow.

### Step-by-step examples

1. **Before:** worker A holds fencing token 7 and pauses past lease
  expiry.
2. **Decision:** worker B obtains token 8 and writes; A later resumes.
3. **After:** storage rejects A's stale token 7 write.
4. **Known correct:** the newest lease owns the resource even if old
  work is
  still running, and a timeout never proves that A stopped.

### Why it works

The invariant is that accepted mutations carry a token no lower than
the
latest accepted token. Durable monotonic allocation and storage-side
checking
preserve that order; timeouts affect liveness but cannot grant stale
safety.

### Go example

See the storage-side fencing check in
[the production examples][production-examples]. Its
[tests][production-tests] accept a newer owner and reject a paused
stale holder after reassignment.


### Work and memory

Replication multiplies write, storage, and network work by replica
count.
Quorum latency follows the required response order statistic. Metadata
for
versions, deduplication, and reconciliation grows with keys and
retention.

### When to use it

Use these mechanics when state or work crosses process, machine,
region, or
broker boundaries and failures can be partial.

### When not to use it

Do not introduce consensus, distributed locks, or multi-region writes
when a
single transactional owner meets the real availability requirement.

### Common mistakes and edge cases

- Treating a timeout as proof that an operation did not commit.
- Using wall-clock order as causality or lease safety.
- Calling at-least-once transport exactly-once processing.
- Omitting fencing from a distributed lock or lease.
- Quoting CAP without naming the operation and partition behavior.

### Related material

See [distributed foundations][distributed-reference],
[service communication](#20-synchronous-service-communication), and
[distributed
workflows](#21-distributed-workflows-and-eventual-consistency).

## 19. Service-boundary discovery and extraction

### The problem it solves

Splitting code along technical layers creates chatty services and
distributed
transactions without establishing independent business ownership.

### The idea in simple terms

Discover boundaries from business capabilities, invariants, data
ownership,
team ownership, and change patterns. Prove them inside a modular
monolith,
then extract one measured seam at a time.

### How it works

Map commands, events, consistency needs, and coupled changes. Give
each
candidate one owner for writes and a published contract for access.
Measure
latency, deployment coupling, and team contention. Extract with a
facade,
shadow reads, backfill and verify data, switch traffic gradually, then
remove
the old path only after rollback and reconciliation are proved.

### Step-by-step examples

1. **Before:** billing code owns invoice writes behind an internal
  interface.
2. **Decision:** a facade records old and candidate results during
  shadowing.
3. **After:** verified traffic moves to the billing service by
  percentage.
4. **Known correct:** one writer owns each invoice, mismatches are
  measured,
  and the route can return to the monolith without losing writes.

### Why it works

The invariant is single ownership of each durable mutation throughout
the
migration. The facade makes routing explicit, comparison supplies
evidence,
and staged cutover preserves a tested reversal at every transition.

### Go example

See the composition-root-selected notification adapter in
[the distributed examples][distributed-examples]. Its
[tests][distributed-tests] cover local success and unavailable remote
delivery without hiding the failure.


### Work and memory

Shadowing can nearly double read or compute work. Backfill is `O(n)`
in moved
records and needs bounded batches. Network extraction adds
serialization,
tail latency, telemetry, and failure-state costs to every call.

### When to use it

Use extraction when ownership and independent delivery benefits are
measured
and exceed permanent distributed operational cost.

### When not to use it

Do not split by table, controller, or repository, or extract merely
because a
codebase is large; improve internal boundaries first.

### Common mistakes and edge cases

- Allowing both old and new services to write the same records.
- Migrating data without checksums or reconciliation.
- Ignoring cross-boundary transactions until after cutover.
- Comparing only happy-path responses during shadowing.
- Removing the rollback path before a full operational window.

### Related material

See [service extraction][extraction-reference],
[architecture](#5-architecture-and-service-construction), and
[contract evolution](#22-contract-evolution-and-independent-delivery).

## 20. Synchronous service communication

### The problem it solves

A remote call consumes connection, goroutine, deadline, and downstream
capacity while adding partial failure that an ordinary function call
lacks.

### The idea in simple terms

Make every call explicit about contract, time budget, retry safety,
admission,
error classification, and observability. Treat fallback as product
behavior.

### How it works

Derive the child timeout from the remaining caller budget. Reuse
bounded
clients and connections. Validate protocol and schema at both ends.
Retry
only classified transient failures for idempotent operations, with
jitter and
an attempt cap. Scope bulkheads and breakers to the failing
dependency, and
propagate identity plus trace context through a trusted boundary.

### Step-by-step examples

1. **Before:** the caller has 800 ms; local work reserves 200 ms.
2. **Decision:** the dependency receives a 500 ms deadline and one
  safe retry.
3. **After:** the first transient failure backs off; the second call
  succeeds.
4. **Known correct:** total work remains within 800 ms, and a
  permanent error
  or exhausted budget is returned without another attempt.

### Why it works

The invariant is that downstream attempts cannot outlive or exceed the
caller's budget. Derived deadlines, bounded attempts, and cancellation
preserve that limit, while classification prevents unsafe duplicate
effects.

### Go example

See the remaining-deadline calculator and reservation reducer in
[the distributed examples][distributed-examples]. Their
[tests][distributed-tests] cover a reserved response and an ambiguous
timeout that requires reconciliation.


### Work and memory

Worst-case calls equal the attempt cap. Latency includes each attempt
plus
backoff but remains under the parent deadline. Pools, queues, and
concurrency
limits bound sockets, goroutines, and downstream offered load.

### When to use it

Use synchronous calls when the caller needs an immediate answer and
the
combined availability and latency budget is acceptable.

### When not to use it

Do not hold a database transaction across the call or retry ambiguous
non-idempotent writes without a durable key.

### Common mistakes and edge cases

- Giving every hop the full original timeout.
- Retrying at client, proxy, and server simultaneously.
- Creating a new transport client per request.
- Falling back to stale or partial data without a product contract.
- Treating all 5xx, resets, and timeouts as the same retry condition.

### Related material

See [service extraction][extraction-reference],
[performance and
overload](#14-performance-capacity-retry-and-overload),
and [distributed mechanics](#18-distributed-systems-mechanics).

## 21. Distributed workflows and eventual consistency

### The problem it solves

One business action may span services that cannot share a database
transaction. Partial success must be visible, retryable, and
repairable.

### The idea in simple terms

Represent the workflow as durable states and messages. Advance each
local
step atomically, compensate only effects that are semantically
reversible,
and expose pending or failed states as honest product behavior.

### How it works

Choose orchestration when one owner must observe and direct the
sequence;
choose choreography only when event relationships remain
understandable.
Persist workflow ID, version, state, deadlines, attempts, and step
results.
Use outbox and inbox guarantees for transitions. Define compensation,
manual repair, and reconciliation before release, including late
replies and
duplicate messages.

### Step-by-step examples

1. **Before:** order `O-42` is pending payment and workflow state is
  `pay`.
2. **Decision:** payment succeeds, but inventory reservation is
  rejected.
3. **After:** the workflow requests a refund and ends `canceled` after
  reply.
4. **Known correct:** each step effect is deduplicated, pending is
  observable,
  and an exhausted compensation becomes an operator-owned repair item.

### Why it works

The invariant is that every external effect has a durable
corresponding
workflow state and idempotency identity. A transition and its next
message
commit locally, so crashes can replay without losing progress. Finite
states
and attempt limits guarantee a terminal success, compensated, or
repair state.

### Go example

See the versioned saga reducer in
[the distributed examples][distributed-examples]. Its
[tests][distributed-tests] cover a valid transition, duplicate no-op,
and out-of-order rejection.


### Work and memory

Work is linear in attempted steps and retries. Durable state and audit
history
grow per workflow; indexes must support ready, timed-out, and repair
states.
Completion latency includes every asynchronous wait and backoff.

### When to use it

Use it for cross-service actions whose partial states are acceptable
and can
be compensated or repaired.

### When not to use it

Do not distribute an invariant that truly requires immediate atomicity
when
one transactional owner can satisfy it.

### Common mistakes and edge cases

- Calling compensation a rollback when an external fact already
  escaped.
- Keeping workflow state only in process memory.
- Omitting late, duplicate, and out-of-order replies.
- Hiding pending state from clients and operators.
- Allowing choreography to form cycles with no accountable owner.

### Related material

See [service extraction][extraction-reference],
[reliable events](#17-reliable-event-publication-and-consumption), and
[idempotency](#13-idempotency-and-ambiguous-write-outcomes).

## 22. Contract evolution and independent delivery

### The problem it solves

Services deployed at different times must communicate while producers
and
consumers temporarily run different contract versions.

### The idea in simple terms

Make changes additive first, measure old usage, test consumer
expectations,
and remove behavior only after every known consumer has migrated.

### How it works

Specify syntax and semantics, not just field types. Add optional
fields with
safe defaults, keep old producers readable, and tolerate unknown
fields where
the protocol permits. Use consumer-driven or provider contract tests
for
important interactions, schema compatibility checks for events, and a
deprecation inventory with owner, deadline, traffic evidence, and
rollback.

### Step-by-step examples

1. **Before:** consumers read event v1 without `currency`.
2. **Decision:** the producer adds optional `currency` and still emits
  v1 data.
3. **After:** old consumers ignore it; new consumers default missing
  values.
4. **Known correct:** both version combinations pass contracts, and
  removal is
  blocked until telemetry shows no old consumer traffic.

### Why it works

The invariant is that every supported producer-consumer version pair
shares
one valid interpretation. Additive changes preserve the previous
subset, and
compatibility tests plus usage evidence make the safe removal point
explicit.

### Go example

See the compatible response decoder in
[the distributed examples][distributed-examples]. Its
[tests][distributed-tests] accept an absent optional field and map an
unknown enum to an explicit fallback.


### Work and memory

Compatibility cost grows with the supported version matrix. Keep the
window
bounded with explicit migration deadlines. Dual read or write paths
add
temporary CPU, storage, and reconciliation work.

### When to use it

Use it for public APIs, events, persisted payloads, and independently
deployed
services or clients.

### When not to use it

Do not create a new version for an additive compatible field, or
retain every
historic behavior without an owned support policy.

### Common mistakes and edge cases

- Changing meaning while preserving the same field name and type.
- Testing schemas but not defaults, errors, ordering, or retry
  semantics.
- Removing a field based on deployment completion rather than
  observed use.
- Dual-writing without comparison and repair.
- Letting undocumented consumers define permanent accidental
  contracts.

### Related material

See [service extraction][extraction-reference],
[reliable events](#17-reliable-event-publication-and-consumption), and
[delivery readiness](#12-delivery-deployment-and-recovery-readiness).

## 23. Distributed observability and resilience experiments

### The problem it solves

A distributed request can fail across several services, queues, and
retries.
Healthy local dashboards do not reveal the causal chain or prove
recovery.

### The idea in simple terms

Propagate correlation and trace context across every boundary, measure
each
service against user-visible objectives, and test one bounded failure
hypothesis with an abort condition and recovery proof.

### How it works

Create spans around remote attempts and message processing, linking
rather
than pretending asynchronous work is one call stack. Record stable
outcome,
dependency, attempt, and queue-delay fields. Build service-level
objectives
from end-user success and latency. For resilience experiments, name
the
steady state, inject one approved failure in a limited blast radius,
stop on
guardrail breach, restore, and record evidence plus follow-up work.

### Step-by-step examples

1. **Before:** checkout success is 99.95% and one canary worker is
  selected.
2. **Decision:** add 200 ms payment latency with a 99.5% abort
  guardrail.
3. **After:** deadlines shed work; success stays above the guardrail.
4. **Known correct:** linked traces show the delayed attempt, rollback
  removes
  injection, and the service returns to the measured steady state.

### Why it works

The invariant is that an experiment cannot exceed its approved scope
or
continue after a guardrail breach. Explicit targeting and automated
abort
preserve that boundary; before-and-after indicators prove restoration
rather
than merely ending the injection.

### Go example

See the allowlisted event envelope in
[the distributed examples][distributed-examples]. Its
[tests][distributed-tests] preserve trace metadata and reject
credential baggage.


### Work and memory

Tracing and experiments add bounded signal volume and test traffic.
Sampling,
label cardinality, retention, and injection concurrency need hard
limits.
Analysis cost grows with services and attempts on the traced path.

### When to use it

Use it after baseline observability and recovery controls exist for a
critical
distributed path.

### When not to use it

Do not inject unbounded production failure, experiment without an
owner, or
use chaos to discover basic missing timeouts that static review
already shows.

### Common mistakes and edge cases

- Starting a new trace instead of propagating the parent or message
  link.
- Using high-cardinality identifiers as metric labels.
- Defining steady state from internal health instead of user outcomes.
- Running without a tested abort and restoration procedure.
- Declaring success because the injection ended rather than service
  recovered.

### Related material

See [service extraction][extraction-reference],
[observability](#11-observability-and-operational-controls), and
[reliability
engineering](#28-reliability-engineering-incidents-and-cost).

## 24. Data scaling, caching, and tenant isolation

### The problem it solves

Growing data and traffic pressure one database, while caches,
partitions, and
shared tenancy can introduce stale reads, hot keys, leakage, and
migrations
that are harder than the original bottleneck.

### The idea in simple terms

Scale from measured access patterns. Preserve one source of truth,
define
freshness and invalidation, choose a partition key that matches
ownership, and
enforce tenant identity in every storage and cache key.

### How it works

Optimize queries and indexes before adding replicas or shards. Cache
only
when hit rate and tolerated staleness justify invalidation complexity;
use
bounded TTL, size, request coalescing, and negative-cache policy.
Partition by
a stable high-cardinality key, route through one owner, detect hot
partitions,
and design resharding before launch. Carry tenant ID from
authenticated
principal through authorization, rows, indexes, cache keys, jobs, and
logs.

### Step-by-step examples

1. **Before:** tenant A and B both request product key `P1`.
2. **Decision:** cache lookup uses `(tenant, product)` and misses only
  for A.
3. **After:** A loads and caches its row; B cannot observe A's value.
4. **Known correct:** database policy repeats the tenant predicate,
  invalidation
  names the same compound key, and a hot key cannot create unbounded
  fills.

### Why it works

The invariant is that every derived location retains the authoritative
tenant
and object identity. Compound keys and storage policy preserve
isolation. Length-prefixed components keep delimiters inside tenant or
object
identities from creating the same serialized key. Storage policy and
bounded coalescing allow one refill while waiters share its result.

### Go example

See the typed tenant cache key in
[the platform examples][platform-examples]. Its
[tests][platform-tests] reject empty tenant identity and prove that
delimiter-containing identities cannot alias.


### Work and memory

Cache lookup is `O(1)` average and memory is bounded by configured
entries or
bytes. Partition routing is constant after metadata lookup. Resharding
is
`O(n)` in moved data and consumes source, destination, and network
capacity.

### When to use it

Use these tools after measurements identify data volume, read load,
locality,
or isolation as the limiting dimension.

### When not to use it

Do not cache correctness-critical mutable state without a freshness
contract,
or shard before a single well-operated database is insufficient.

### Common mistakes and edge cases

- Omitting tenant ID from a cache, queue, or deduplication key.
- Treating TTL as precise invalidation.
- Allowing cache stampedes or unbounded negative caching.
- Choosing a time or geography key that concentrates current writes.
- Using replicas for read-after-write paths without a consistency
  policy.

### Related material

See [scale and cloud operations][scale-reference],
[advanced PostgreSQL](#15-advanced-postgresql-operation-and-recovery),
and
[privacy and supply
chain](#25-security-privacy-and-software-supply-chain).

## 25. Security, privacy, and software supply chain

### The problem it solves

Production risk includes vulnerable dependencies, compromised builds,
excess
privilege, retained personal data, unsafe logs, and secrets that
outlive their
intended scope.

### The idea in simple terms

Threat-model assets and trust boundaries, minimize collected data and
privilege, secure the build inputs and artifact, and make deletion,
rotation,
and incident response executable rather than aspirational.

### How it works

List assets, actors, boundaries, abuse cases, and mitigations.
Classify data,
collect only required fields, define purpose and retention, encrypt in
transit
and at rest, and keep searchable secrets out of logs. Pin and scan
dependencies,
generate an SBOM and provenance, sign artifacts, protect CI
credentials, and
patch by severity plus exploitability. Rotate keys with overlap and
audit every
privileged data or secret action.

### Step-by-step examples

1. **Before:** signing key version 3 serves traffic and version 4 is
  published.
2. **Decision:** new artifacts use v4 while verification accepts v3
  and v4.
3. **After:** old artifacts age out; v3 is revoked and removed.
4. **Known correct:** no unsigned gap occurred, access was audited,
  and a
  rollback window existed before destructive retirement.

### Why it works

The invariant is that every accepted artifact or credential chains to
a
currently trusted, auditable authority. Overlap preserves availability
during
rotation, while explicit retirement bounds exposure. Data minimization
reduces
the impact and compliance surface when another control fails.

### Go example

See least-privilege authorization and overlap rotation in
[the platform examples][platform-examples]. Their
[tests][platform-tests] cover one allowed tuple, denied excess access,
accepted overlap keys, and retirement.


### Work and memory

Verification adds bounded cryptographic and metadata work. Scanning
grows with
dependency and artifact count. Encryption, audit logs, and retention
copies
consume storage; deletion must cover indexes, caches, backups, and
replicas
under their documented lifecycle.

### When to use it

Use it for every build and production system, with depth proportional
to data
sensitivity, privilege, exposure, and recovery cost.

### When not to use it

Do not collect data for hypothetical future use or grant permanent
broad
credentials because rotation and policy design were postponed.

### Common mistakes and edge cases

- Scanning dependencies without an owner or remediation deadline.
- Signing an artifact after an untrusted step can still modify it.
- Treating encryption as a substitute for authorization and
  minimization.
- Deleting the primary row while copies remain in caches or exports.
- Rotating secrets without dual-read overlap or rollback.

### Related material

See [scale and cloud operations][scale-reference],
[identity and security](#9-identity-authorization-and-api-security),
and
[cloud IAM](#26-cloud-infrastructure-and-iam).

## 26. Cloud infrastructure and IAM

### The problem it solves

Cloud APIs make infrastructure easy to create and equally easy to
drift,
overprivilege, expose publicly, or recreate differently during
recovery.

### The idea in simple terms

Describe infrastructure as reviewed state, separate environments and
failure
domains deliberately, and grant each workload only short-lived
permissions to
the exact resources and actions it needs.

### How it works

Use versioned infrastructure definitions, remote locked state, plan
review,
policy checks, and drift detection. Build private network paths by
default,
restrict ingress and egress, and expose only intended load balancers.
Bind
workload identity to deployment identity, avoid static cloud keys,
constrain
resource and action scope, and audit assumptions behind cross-account
access.
Test replacement and regional recovery from definitions plus durable
data.

### Step-by-step examples

1.  **Before:** order worker identity may read queue Q but not bucket
   B.
2. **Decision:** it requests one message from Q with a short-lived
  credential.
3. **After:** queue access succeeds; an attempted bucket read is
  denied.
4. **Known correct:** policy names the workload, action, and resource;
  no static
  secret exists in the image, environment file, or repository.

### Why it works

The invariant is that deployed resources and permissions are
reproducible from
reviewed declarations, and every allowed action matches an explicit
workload
need. Short credential lifetime and narrow scope bound misuse after
compromise.

### Go example

See the boundary-aware diagnostic classifier in
[the platform examples][platform-examples]. Its
[tests][platform-tests] preserve network, authorization, and
application failure categories.


### Work and memory

Plans scale with resource graph size and provider calls. Policy
evaluation is
small per request but audit volume grows with activity. Multi-zone or
regional
redundancy multiplies running cost, data transfer, and recovery
coordination.

### When to use it

Use reviewed infrastructure and workload identity for every non-local
environment, especially shared accounts and automated deployment
systems.

### When not to use it

Do not put changing application rollout logic into low-level
infrastructure
state, or use one administrator credential for all workloads.

### Common mistakes and edge cases

- Applying an unreviewed plan or editing managed resources manually.
- Sharing state files or broad roles across unrelated environments.
- Assuming private addressing alone authenticates a caller.
- Allowing unrestricted egress that bypasses dependency policy.
- Designing redundancy without testing data restoration and traffic
  failover.

### Related material

See [scale and cloud operations][scale-reference],
[delivery readiness](#12-delivery-deployment-and-recovery-readiness),
and
[platform operations](#27-orchestration-and-platform-operations).

## 27. Orchestration and platform operations

### The problem it solves

Schedulers can restart and scale processes, but poor requests, probes,
shutdown, disruption policy, or queue signals turn automation into
churn.

### The idea in simple terms

Declare resource needs and lifecycle truth, keep stateless compute
replaceable,
and let autoscaling and disruption act only on signals the workload
can honor.

### How it works

Package one immutable process, run as a non-root identity, and declare
CPU and
memory requests plus limits from measurement. Separate startup,
readiness, and
liveness probes. On termination, fail readiness, stop admission, drain
within
the grace period, and persist or return leased work. Use disruption
budgets,
topology spread, and rollouts that preserve capacity. Scale request
services
from demand and workers from queue age plus dependency capacity, not
queue
length alone.

### Step-by-step examples

1. **Before:** a worker pod holds job `J7` and receives termination.
2. **Decision:** readiness fails, no new lease is claimed, and `J7` is
  finished
  or released before grace expiry.
3. **After:** the scheduler replaces the pod and the job remains
  recoverable.
4. **Known correct:** no acknowledgement precedes durable completion,
  and the
  rollout maintains the minimum available capacity.

### Why it works

The invariant is that only ready instances receive new work and every
owned
unit is completed or made eligible before process death. Probe
separation and
bounded drain preserve that lifecycle under restart and deployment.

### Go example

See the startup, readiness, and liveness state object in
[the platform examples][platform-examples]. Its
[tests][platform-tests] cover healthy service, dependency loss, and
draining without creating a restart loop.


### Work and memory

Requests reserve cluster capacity; limits bound but can throttle or
kill.
Autoscaling lag requires headroom. Replicas and topology rules
multiply
baseline cost, while queue-based scaling must remain below downstream
capacity.

### When to use it

Use orchestration once multiple deployable instances need automated
placement,
replacement, rollout, and capacity management.

### When not to use it

Do not add a cluster when a simpler supervised deployment meets
reliability
needs, or use restarts to mask deterministic application defects.

### Common mistakes and edge cases

- Using the same endpoint for startup, readiness, and liveness.
- Marking ready before listeners and required local state are usable.
- Setting limits without measuring runtime and burst behavior.
- Scaling workers beyond database or partner capacity.
- Ignoring termination grace while holding leases or serving requests.

### Related material

See [scale and cloud operations][scale-reference],
[HTTP lifecycle](#6-http-contracts-and-process-lifecycle), and
[durable jobs](#16-durable-background-jobs-and-scheduling).

## 28. Reliability engineering, incidents, and cost

### The problem it solves

Without explicit objectives and ownership, teams optimize noisy
metrics,
repeat incidents, and spend unpredictably without knowing which cost
buys a
user-visible reliability improvement.

### The idea in simple terms

Define service indicators and objectives, spend an error budget
consciously,
design on-call and runbooks around actionable symptoms, learn through
blameless
incident analysis, and attribute cost to the same demand and
reliability model.

### How it works

Choose availability and latency indicators at the user boundary. Set
targets
from product impact and dependency limits, then alert on multi-window
burn
rather than isolated spikes. Run incidents with one commander,
communication
cadence, timestamped decisions, and explicit mitigation. Follow with
factual
timeline, contributing conditions, systemic actions, owners, and
verification.
Track unit cost per request, tenant, job, or stored byte and include
redundancy,
egress, telemetry, and idle headroom.

### Step-by-step examples

1. **Before:** checkout has consumed 40% of its monthly error budget
  in one hour.
2. **Decision:** the fast burn alert pages; rollout stops and traffic
  reverts.
3. **After:** success recovers, stakeholders receive status, and
  evidence is kept.
4. **Known correct:** the alert maps to the SLO, mitigation has an
  owner, and
  follow-up tests the condition rather than assigning personal blame.

### Why it works

The invariant is that reliability decisions use a user-visible
objective and
remaining budget. Burn rates connect short evidence to long targets,
while
owned actions and verification convert incident learning into changed
system
behavior.

### Go example

See the integer event-budget calculation in
[the platform examples][platform-examples]. Its
[tests][platform-tests] cover a 99.9 percent objective and reject a
zero-event denominator.


### Work and memory

Metric aggregation grows with signals and label cardinality.
Redundancy and
headroom add direct cost; incidents add human interruption cost. Unit
economics
must include shared infrastructure allocation and peak rather than
only average
demand.

### When to use it

Use it for every production capability with a user, an on-call owner,
or a
material infrastructure and incident cost.

### When not to use it

Do not create objectives no team can observe or act on, or reduce
reliability
to an availability percentage without latency and correctness.

### Common mistakes and edge cases

- Paging on CPU alone without a user symptom or exhaustion risk.
- Setting every target to 100% and eliminating room for change.
- Writing postmortems with no verified systemic action.
- Ignoring third-party, egress, telemetry, and standby costs.
- Optimizing average unit cost while peak capacity remains the
  constraint.

### Related material

See [scale and cloud operations][scale-reference],
[observability](#11-observability-and-operational-controls), and
[delivery readiness](#12-delivery-deployment-and-recovery-readiness).

## 29. Production integration boundaries

### The problem it solves

Webhooks, files, email, search, analytics, and partner APIs have
different
trust, delivery, size, privacy, and failure semantics from internal
calls.

### The idea in simple terms

Give each integration a narrow adapter, verify its authenticity and
schema,
persist important work, make effects idempotent, quarantine poison
input, and
reconcile against an authoritative source.

### How it works

For webhooks, verify the signature over raw bytes, enforce timestamp
and replay
bounds, store event identity, and acknowledge only after durable
acceptance.
For files, stream with byte and record limits, validate headers and
rows, stage
results, and publish only after atomic validation. For email, search,
and
analytics, separate delivery from source-of-truth state, redact
personal data,
and provide replay, correction, and deletion paths. Partner clients
need
deadlines, rate limits, contract tests, and explicit outage behavior.

### Step-by-step examples

1. **Before:** signed webhook `E9` is unseen and its timestamp is
  current.
2. **Decision:** verify raw bytes, claim `E9`, and enqueue durable
  processing.
3. **After:** return success; a repeated `E9` returns the same
  acceptance.
4. **Known correct:** altered bytes or stale timestamps are rejected,
  duplicate
  delivery has one effect, and reconciliation can find a missed event.

### Why it works

The invariant is that only authenticated, bounded, schema-valid input
becomes
durable work, and each external identity has at most one semantic
effect.
Raw-byte verification establishes origin evidence; durable identity
and later
reconciliation cover transport loss and repetition.

### Go example

See the bounded webhook receiver in
[the delivery examples][delivery-examples]. Its
[tests][delivery-tests] cover valid acceptance, signature failure, and
duplicate acknowledgement without duplicate work.


### Work and memory

Signature checking and streaming are linear in bounded input size.
Deduplication
storage grows with event rate and replay window. Bulk imports need
bounded
batches; search and analytics copies add indexing, transfer, and
deletion cost.

### When to use it

Use explicit integration boundaries whenever data or effects cross an
external
trust, ownership, or delivery system.

### When not to use it

Do not make search or analytics the transactional source of truth, or
perform
slow partner work inside a local database transaction.

### Common mistakes and edge cases

- Parsing JSON before verifying the signature over original bytes.
- Acknowledging a webhook before durable acceptance.
- Loading an entire untrusted file into memory.
- Treating email acceptance as delivery or analytics as exact
  accounting.
- Omitting reconciliation because the normal event path usually works.

### Related material

See [integrations and maintenance][integration-reference],
[identity and security](#9-identity-authorization-and-api-security),
and
[reliable events](#17-reliable-event-publication-and-consumption).

## 30. Maintenance and legacy-system evolution

### The problem it solves

An old system often lacks tests, clear boundaries, current
dependencies, or a
safe migration path, yet it continues to carry valuable behavior and
data.

### The idea in simple terms

Observe before changing, pin behavior at the boundary, add a seam,
migrate in
small reversible slices, compare results, and delete the old path only
after
evidence and an operational window.

### How it works

Inventory entry points, data owners, jobs, dependencies, and incident
history.
Add characterization tests around current observable behavior,
including known
quirks. Introduce a facade or adapter without changing outcomes. Route
a small
slice to the replacement, shadow and compare when safe, reconcile
data, expand
traffic, and keep one rollback switch. Upgrade dependencies in small
groups and
separate compatibility changes from behavior changes.

### Step-by-step examples

1. **Before:** legacy pricing returns result `P`; no reliable unit
  tests exist.
2. **Decision:** record boundary cases, place a facade, and shadow new
  pricing.
3. **After:** mismatches are classified; matching traffic shifts
  gradually.
4. **Known correct:** the old path remains available, data checks
  agree, and
  removal waits through the defined rollback window.

### Why it works

The invariant is that every migrated request has one authoritative
result and
a tested reversal until confidence is earned. Characterization
establishes the
baseline; facades and staged routing make each state observable and
reversible.

### Go example

See the tenant-scoped, version-guarded repair plan in
[the delivery examples][delivery-examples]. Its
[tests][delivery-tests] cover a dry run and reject an out-of-scope
tenant.


### Work and memory

Shadowing duplicates selected work. Backfill and comparison are `O(n)`
in
records and require rate limits. Temporary adapters and dual paths
increase
code and operational load, so each needs an owner and removal
condition.

### When to use it

Use it when preserving current service matters more than a fast
rewrite and
the existing behavior or data cannot be reconstructed confidently.

### When not to use it

Do not preserve unsafe behavior that violates a required security or
integrity
rule; isolate and correct it with an explicit migration decision.

### Common mistakes and edge cases

- Rewriting from undocumented intent while ignoring actual behavior.
- Adding broad mocks that encode implementation rather than outcomes.
- Running dual writes without comparison, ordering, or repair.
- Leaving temporary facades and feature flags with no removal owner.
- Combining dependency, schema, and behavior changes in one release.

### Related material

See [integrations and maintenance][integration-reference],
[service extraction](#19-service-boundary-discovery-and-extraction),
and
[contract evolution](#22-contract-evolution-and-independent-delivery).

## 31. Capstone delivery and engineering leadership

### The problem it solves

Senior backend work must turn uncertain product and operational needs
into a
reviewable plan, delivered system, measurable launch, and maintained
learning
rather than only produce locally correct code.

### The idea in simple terms

Frame the decision, name constraints and invariants, compare real
alternatives,
slice delivery into reversible milestones, assign owners, and close
the loop
with launch evidence and durable documentation.

### How it works

Write a short proposal with problem, non-goals, users, current
evidence,
options, decision, risks, data and security impact, rollout, rollback,
and
success measures. Use a walking skeleton before scale features. Review
API,
schema, failure modes, operations, and cost with the owning
specialists. Track
milestones by observable capability. During launch, use explicit go or
no-go
criteria; afterward, compare results with assumptions and record the
decision
and remaining work.

### Step-by-step examples

1. **Before:** checkout timeout rate is 2%; cause and ownership are
  uncertain.
2. **Decision:** the proposal chooses one bounded queue after
  comparing query
  optimization and capacity, with latency and cost guardrails.
3. **After:** a canary meets the guardrails and expands by staged
  milestones.
4. **Known correct:** the decision log captures rejected options, each
  risk has
  an owner, and rollback triggers before a user-visible budget breach.

### Why it works

The invariant is that every material claim maps to evidence, every
risk to an
owner, and every rollout step to acceptance and reversal criteria.
That makes
technical judgment reviewable before cost is sunk and correctable
afterward.

### Go example

See the typed requirement-evidence validator in
[the delivery examples][delivery-examples]. Its
[tests][delivery-tests] accept a complete evidence chain and reject a
missing runbook.


### Work and memory

Planning cost should scale with reversibility and blast radius. Small
reversible
changes need little ceremony; durable data, public contracts, privacy,
or
cross-team operations justify deeper review. Unbounded coordination is
itself
a delivery cost and needs decision deadlines.

### When to use it

Use this method for cross-boundary work, public or durable contracts,
risky
migrations, and projects whose operation spans several owners.

### When not to use it

Do not create a large design process for a local reversible change, or
mistake
consensus for accountable decision ownership.

### Common mistakes and edge cases

- Presenting one option and calling the document a tradeoff analysis.
- Leaving non-goals, data migration, rollback, or operational
  ownership implicit.
- Tracking tasks without testing an end-to-end capability.
- Launching on schedule despite breached evidence-based guardrails.
- Recording lessons without changing code, runbooks, tests, or
  ownership.

### Related material

See [integrations and maintenance][integration-reference],
[architecture](#5-architecture-and-service-construction), and
[reliability
engineering](#28-reliability-engineering-incidents-and-cost).

## 32. Language portability

### The problem it solves

Backend principles survive a language change, but ownership,
cancellation,
errors, concurrency, runtime scheduling, and tooling appear through
different
syntax and defaults.

### The idea in simple terms

Port the invariant and boundary first. Then map Go's value model,
interfaces,
context, goroutines, errors, and tooling to the target language's real
runtime
semantics instead of transliterating syntax.

### How it works

For each concept, write the behavior to preserve, Go mechanism, target
mechanism, and semantic gap. Check memory and aliasing, nullability,
exception
or result propagation, cancellation ownership, async scheduling,
thread safety,
serialization defaults, package visibility, build reproducibility,
testing,
profiling, and deployment. Re-run the same ordinary and failure
contract tests
against both implementations.

### Step-by-step examples

1. **Before:** Go cancellation ends a database call through request
  context.
2. **Decision:** the target uses an owned cancellation token
  propagated to its
  database driver, not an unobserved task timeout.
3. **After:** both versions return the same canceled contract and
  release work.
4. **Known correct:** conformance tests match outcomes, resource
  lifetime, and
  deadline behavior despite different language syntax.

### Why it works

The invariant is an observable contract plus owned resource lifetime,
not a Go
keyword. Describing the semantic gap exposes where a target default
would
weaken that invariant; shared conformance cases then verify the
adaptation.

### Go example

See the portable operation outcome contract in
[the delivery examples][delivery-examples]. Its
[tests][delivery-tests] preserve ordinary success and cancellation as
stable outcomes.


### Work and memory

Complexity must be recalculated for the target runtime: allocation,
garbage
collection, task or thread cost, string encoding, collection behavior,
and I/O
drivers can change constants or even effective bounds.

### When to use it

Use the matrix when moving a service, reviewing a polyglot design, or
teaching
a backend invariant across languages.

### When not to use it

Do not choose a language from syntax familiarity alone or assume
similarly
named concurrency and cancellation features have identical guarantees.

### Common mistakes and edge cases

- Translating goroutines directly to threads without capacity
  analysis.
- Mapping Go error values to unchecked exceptions without contract
  handling.
- Assuming strings, maps, nulls, and value copies share Go semantics.
- Dropping cancellation because the target library uses a different
  API.
- Comparing benchmarks without equivalent workload and runtime warmup.

### Related material

See [language portability][portability-reference],
[Go ownership](#1-go-values-program-structure-and-ownership), and
[Go cancellation](#2-go-errors-resources-and-cancellation).

[go-refresher]: #1-go-values-program-structure-and-ownership
[go-foundations]: #2-go-errors-resources-and-cancellation
[domain-reference]: #4-domain-discovery-and-modeling
[service-core]: #5-architecture-and-service-construction
[postgres-reference]: #7-relational-modeling-and-postgresql-access
[security-reference]: #9-identity-authorization-and-api-security
[testing-reference]: #10-deterministic-backend-testing
[delivery-reference]: #12-delivery-deployment-and-recovery-readiness
[distributed-reference]: #18-distributed-systems-mechanics
[extraction-reference]: #19-service-boundary-discovery-and-extraction
[scale-reference]: #24-data-scaling-caching-and-tenant-isolation
[integration-reference]: #29-production-integration-boundaries
[portability-reference]: #32-language-portability
[service-examples]: examples/service.go
[service-tests]: examples/service_test.go
[production-examples]: examples/production.go
[production-tests]: examples/production_test.go
[distributed-examples]: examples/distributed.go
[distributed-tests]: examples/distributed_test.go
[platform-examples]: examples/platform.go
[platform-tests]: examples/platform_test.go
[delivery-examples]: examples/delivery.go
[delivery-tests]: examples/delivery_test.go
[resilience-experiments]:
  #23-distributed-observability-and-resilience-experiments
