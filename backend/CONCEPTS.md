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

### How it works

Copying a paper folder makes a new folder; copying a shortcut still
points to the same cabinet. Go values work the same way: arrays copy
their contents, while slices, maps, and pointers may still reach
shared storage. **Ownership** says who may change that storage. An
interface stores both a dynamic type and a value, so an interface
holding a typed nil pointer is itself non-nil.

### Step-by-step examples

Input: clone `Product{Tags: ["tea", "sale"]}`, then change the
source's first tag to `"coffee"`. Expected output: the source tags
become `["coffee", "sale"]`, while the clone stays
`["tea", "sale"]`.

```text
before Clone:  product.Tags -> array A [tea, sale]
after Clone:   product.Tags -> array A [tea, sale]
               clone.Tags   -> array B [tea, sale]
after change:  product.Tags -> array A [coffee, sale]
               clone.Tags   -> array B [tea, sale]
```

1. **Before:** Only `product.Tags` points to array A. **Decision:**
   `Clone` allocates array B and copies both strings. **After:** The
   two slices contain equal values but point to different arrays.
   **Known correct:** No mutable array crosses the ownership boundary.
2. **Before:** The source and clone own arrays A and B. **Decision:**
   Change `product.Tags[0]` in array A. **After:** Only the source
   reads `"coffee"`. **Known correct:** The clone cannot observe the
   source's later mutation.

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
	// Starting from nil copies the tags into storage independent of
	// product.Tags, so later mutations cannot cross the boundary.
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

### How it works

If you borrow a key, you return it; if a job is cancelled, everyone
working on it should stop. In Go, the code that acquires a resource
releases it, one context carries cancellation downward, and errors
travel upward with their cause intact.

### Step-by-step examples

Input: `Await` receives a blocked result channel and a context that is
then canceled. Expected output: an error that
`errors.Is(err, context.Canceled)` recognizes.

1. **Before:** Neither `results` nor `ctx.Done()` is ready.
   **Decision:** `Await` remains inside `select`. **After:** No branch
   runs and no value is consumed. **Known correct:** Waiting has not
   invented a result or lost cancellation.
2. **Before:** `results` is still blocked. **Decision:** Cancel the
   context, making `ctx.Done()` ready. **After:** The cancellation
   branch wraps `ctx.Err()` and returns. **Known correct:** The result
   channel was not read and the caller can classify the cause.
3. **Before:** The caller owns the wrapped error. **Decision:** Check
   it with `errors.Is`. **After:** The check reports
   `context.Canceled`. **Known correct:** Added operation context did
   not erase the original error identity.

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

### How it works

Imagine two clerks trying to sell the last ticket. Locking only the
cash drawer is not enough; the check and sale must be protected as one
decision. Each goroutine also needs an owner, a stopping rule, and a
test that proves the business result.

### Step-by-step examples

Input: goroutines A and B call `Create("TEA-001")` at the same time.
Expected output: one call succeeds, one returns `ErrDuplicate`, and
the map contains one product.

1. **Before:** `products={}` and the mutex is free. **Decision:** A
   acquires the mutex; B waits. **After:** A alone may perform the
   check and insert. **Known correct:** B cannot observe the empty map
   and write concurrently.
2. **Before:** A owns the mutex and `products={}`. **Decision:** A
   finds no SKU, stores `TEA-001`, and unlocks. **After:**
   `products={TEA-001}` and A returns success. **Known correct:** One
   successful creation now owns the only map entry.
3. **Before:** B acquires the mutex with the stored SKU visible.
   **Decision:** B reports `ErrDuplicate` without writing. **After:**
   The map still has one product. **Known correct:** Reversing which
   goroutine acquires the mutex changes the winner, not the outcome.

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

### How it works

Start with a real story such as a customer placing and cancelling an
order. Its nouns, actions, rules, and outcomes reveal the domain
model. An **entity** keeps identity, a **value object** is compared by
value, an **aggregate** protects one consistency boundary, and
commands, events, and policies describe allowed change. Keep only
distinctions that affect behavior.

### Step-by-step examples

Input: order `O-42` is `placed` and owned by `C-17`. First `C-99`,
then `C-17`, asks to cancel it. Expected output: `C-99` is rejected
without a state change; `C-17` changes the order to `canceled`.

1. **Before:** `status=placed`, `owner=C-17`, and actor `C-99` sends
   `Cancel`. **Decision:** Compare actor and owner before changing
   state. **After:** Return `ErrCannotCancel`; status stays `placed`.
   **Known correct:** Knowing an order ID did not grant authority.
2. **Before:** The same order is still `placed`; actor is now `C-17`.
   **Decision:** Ownership and the allowed source state both match.
   **After:** Set status to `canceled` and record `OrderCanceled`.
   **Known correct:** One legal command produced one legal transition.
3. **Before:** The order is `canceled`. **Decision:** Reject another
   cancel because `placed` is required. **After:** State remains
   `canceled`. **Known correct:** A terminal state cannot move
   backward or emit the event twice.

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

### How it works

Treat a service like a shop with a front counter and a stockroom. The
handler receives the request, the use case decides the business
action, and a repository talks to storage. Dependencies point inward
toward those stable business rules. The **composition root** connects
concrete parts when the service starts.

### Step-by-step examples

Input: one HTTP request contains order `O-42` with quantity `0`; a
second contains quantity `2`. Expected output: the invalid request
returns `ErrInvalidQuantity` without storage; the valid request calls
`Save` once with a domain `Order`.

1. **Before:** The handler has untrusted JSON for quantity `0`.
   **Decision:** Decode it into a command and call `Place`. **After:**
   The use case rejects the command before `Save`. **Known correct:**
   Transport parsing did not bypass business policy.
2. **Before:** The next command has `ID=O-42` and `Quantity=2`.
   **Decision:** The use case validates the quantity and calls its
   `Store` interface. **After:** The concrete store receives one
   domain value. **Known correct:** HTTP and persistence details do
   not enter the policy.
3. **Before:** A scheduled job owns the same typed command.
   **Decision:** Call the same use case without an HTTP handler.
   **After:** Validation and storage behavior are unchanged.
   **Known correct:** The business operation keeps one owner across
   adapters.

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

### How it works

An API is like a posted service menu: callers rely on its request,
response, and failure rules. HTTP provides methods, headers, status
codes, and caching; the application contract chooses the exact rules
for one operation. At shutdown, stop accepting new work, let bounded
in-flight work finish or cancel it, then close owned resources.

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

Focused shutdown example: input is a termination signal while request
`R1` is running. Expected output: no new request is admitted, `R1`
drains within the deadline, and owned dependencies close.

1. **Before:** Readiness is true and `R1` is active. **Decision:** On
   the signal, set readiness false before shutdown starts.
   **After:** Load balancers stop sending new requests.
   **Known correct:** The
   draining set cannot grow through normal admission.
2. **Before:** `R1` is still active and the shutdown deadline remains.
   **Decision:** Wait for `R1` rather than closing its dependencies.
   **After:** `R1` finishes and the server closes owned resources.
   **Known correct:** Accepted work finished within bounded shutdown.

### Why it works

Each boundary narrows possible state before handing data inward.
Stable error identity lets one boundary translate an outcome without
parsing text. Bounded input, time, and resources guarantee that one
request cannot consume unlimited process capacity.

During shutdown, stopping admission makes the in-flight set only
shrink. Waiting until it is empty or the deadline passes gives
accepted work a bounded chance to finish. Assuming each worker
observes cancellation and reports completion, remaining work then
stops before owned resources close. No accepted request can start
after draining or keep using a resource after shutdown completes.

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

// Bound decoder input before it can allocate from the request body.
// The limit is one mebibyte.
const maximumBodyBytes int64 = 1_048_576

func Decode(
	w http.ResponseWriter,
	r *http.Request,
	destination any,
) error {
	r.Body = http.MaxBytesReader(w, r.Body, maximumBodyBytes)
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(destination); err != nil {
		return err
	}
	// A second decode must reach EOF. Any value means the body contains
	// trailing JSON that the first decode would otherwise ignore.
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

### How it works

Use a database like a labeled filing system. Tables hold kinds of
facts, keys identify records, constraints reject impossible records,
and indexes help find them. Queries bind outside values safely and
close every row, transaction, and connection they open.

### Step-by-step examples

Input: processes A and B concurrently insert
`(sku="TEA-001", name="Tea")`. Expected output: one insert commits,
one returns the stable duplicate conflict, and one row exists.

1. **Before:** No `TEA-001` row exists; both statements bind the SKU
   as data. **Decision:** PostgreSQL evaluates both inserts against
   the same unique constraint. **After:** One statement claims the
   unique value while the other waits or conflicts. **Known correct:**
   Application timing cannot create two committed copies.
2. **Before:** A row with `TEA-001` is committed. **Decision:** Map
   the losing constraint error to the documented conflict.
   **After:** The table contains one valid row and the caller gets a
   stable outcome.
   **Known correct:** The SKU never became SQL syntax and every writer
   faced the same integrity rule.

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

### How it works

A bank transfer must not remove money without adding it at the other
end. A **transaction** makes the whole business action succeed or fail
together. Locks or atomic conditions stop two callers from spending
the same balance, seat, or inventory item.

### Step-by-step examples

Input: inventory `1`; requests A and B each reserve `1`.
Expected output: one success, one insufficient result, and final
stock `0`.

1. **Before:** Stock is `1` and neither request owns it. **Decision:**
   A performs an atomic decrement conditioned on stock at least `1`.
   **After:** A changes stock to `0` and owns the reservation.
   **Known correct:** Stock has never been negative.
2. **Before:** B evaluates the same condition against current stock
   `0`. **Decision:** Its conditional update affects no row.
   **After:** B returns the expected business conflict.
   **Known correct:** No second reservation exists.
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

### How it works

Showing an ID proves who you are; a locked door still decides whether
you may enter. **Authentication** establishes identity, while
**authorization** checks the person, action, resource, and situation.
Input validation, output encoding, secret handling, and browser policy
remain separate controls, so one failure does not remove every
defense.

### Step-by-step examples

Input: authenticated user `U-7` requests `DELETE /orders/O-42`; the
order is owned by `U-9`. Expected output: a stable forbidden result,
no order mutation, and no token value in logs.

1. **Before:** The route contains `O-42` and the verified principal is
   `U-7`. **Decision:** Load the order and compare its owner with the
   principal inside the use case. **After:** The comparison is
   `U-9 != U-7`. **Known correct:** An order ID is not permission.
2. **Before:** Authorization has failed and the order is unchanged.
   **Decision:** Return the stable forbidden outcome and log only the
   request, actor, and order identifiers. **After:** No delete reaches
   storage and no credential is recorded. **Known correct:** The
   response and telemetry preserve both access control and secrecy.

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

### How it works

A useful test controls the clock, random choices, and outside systems
so the same setup gives the same result every time. Test the smallest
public boundary that proves the rule, then use integration tests only
where real storage or protocols matter.

### Step-by-step examples

Input: inventory is `1`; two workers each call `Reserve(1)` after one
start signal. Expected output: one success, one insufficient result,
and final inventory `0`, independent of scheduling.

1. **Before:** Both workers are blocked on `start`; inventory is `1`.
   **Decision:** Close `start` once so both calls become eligible.
   **After:** The workers contend through production synchronization.
   **Known correct:** No sleep guesses when either worker begins.
2. **Before:** One worker owns the reservation transition and the
   other waits. **Decision:** Join both workers and collect both
   outcomes. **After:** No assertion runs while work is unfinished.
   **Known correct:** The test observes the complete result.
3. **Before:** Outcomes may arrive in either order. **Decision:** Sort
   them before comparison. **After:** Assert one success, one
   insufficient result, and inventory `0`. **Known correct:** The
   assertion depends on behavior, not scheduler order.

Failure example: inject storage failure before commit. The expected
output is two failure outcomes and unchanged inventory `1`.

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

### How it works

Think of tracking a package. Logs record events, metrics show the
system-wide pattern, and traces follow one request through each stop.
Health checks, timeouts, rate limits, and profiles turn that evidence
into safe operating controls.

### Step-by-step examples

Input: request `R-8` reaches a database call whose 100 ms deadline
expires. Expected output: an unavailable response, one bounded outcome
metric, and correlated log and trace records without private data.

1. **Before:** `R-8` has a request span and the database attempt has
   not started. **Decision:** Create a child span with dependency name
   and deadline. **After:** The trace shows which wait belongs to the
   request. **Known correct:** Causal context crosses the boundary.
2. **Before:** The database exceeds 100 ms. **Decision:** Classify the
   deadline as dependency unavailability. **After:** Return the stable
   unavailable response and increment one low-cardinality counter.
   **Known correct:** Error text did not become a metric label.
3. **Before:** The request has finished. **Decision:** Emit one event
   with `R-8`, outcome, and duration. **After:** Log, metric, and
   trace agree on the failure. **Known correct:** No private value
   entered telemetry.

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

### How it works

Seal one package, then move that same package through inspection and
into production. The package is an immutable build artifact; secrets
stay outside it. Deploy through compatible stages, watch explicit
health criteria, and practise rollback and data restore before an
emergency.

### Step-by-step examples

Input: immutable version B receives 5% of traffic with a 1% error-rate
stop threshold; B reaches 3%. Expected output: rollout stops and all
traffic returns to healthy version A.

1. **Before:** A serves 95%, B serves 5%, and both read the expanded
   schema. **Decision:** Observe B for the defined window. **After:**
   B's measured error rate is 3%. **Known correct:** The comparison
   uses the stated canary cohort and threshold.
2. **Before:** B exceeds the 1% stop threshold. **Decision:** Stop new
   admission to B and route traffic to A. **After:** A serves 100%.
   **Known correct:** The rollback did not require a rebuild or an
   incompatible schema reversal.
3. **Before:** User success has recovered. **Decision:** Record the
   threshold breach and preserve B's evidence. **After:** The rollout
   remains stopped for diagnosis. **Known correct:** Recovery was
   verified rather than inferred from ending the deployment.

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

### How it works

A payment retry should behave like showing the same receipt number,
not buying twice. An **idempotency key** identifies one logical write.
Atomically store it with a fingerprint of the request and the durable
result, so an identical retry replays the result while conflicting
reuse is rejected.

### Step-by-step examples

Input: create-order command `{item: "tea", quantity: 1}` uses key
`K1`; its first response is lost and the client retries.
Expected output: both attempts return `O-42`; one order exists.

1. **Before:** `K1` is absent. **Decision:** In one transaction, bind
   `K1` to the command hash, create `O-42`, store its response, and
   commit. **After:** The order and completed key record are durable.
   **Known correct:** The semantic effect and replay information
   cannot commit separately.
2. **Before:** The client saw a timeout, but the completed record
   exists. **Decision:** Retry with `K1` and the same command.
   **After:** Return stored order `O-42` without creating an order.
   **Known correct:** Ambiguous delivery did not duplicate the effect.
3. **Before:** `K1` is bound to the original hash. **Decision:** Try
   `{item: "tea", quantity: 2}` with `K1`. **After:** Return a key
   conflict. **Known correct:** One key cannot name two writes.

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

### How it works

A restaurant cannot seat unlimited guests without ruining service for
everyone. Measure the real bottleneck, give each step a share of the
end-to-end time budget, bound queues and workers, retry only safe
operations, and reject excess work early enough that accepted work can
finish.

### Step-by-step examples

Input: all 10 worker slots and all 20 queue positions are occupied
when request `R31` arrives. Expected output: `R31` receives the
documented retryable overload result without allocating queued work.

1. **Before:** `workers=10/10`, `queue=20/20`, and `R31` is not
   admitted. **Decision:** Check capacity before creating a goroutine
   or queue node. **After:** The admission check reports full.
   **Known correct:** The configured bounds still hold.
2. **Before:** Capacity is still full. **Decision:** Return the
   overload result with the bounded retry policy. **After:**
   `workers=10/10` and `queue=20/20`; `R31` owns no server work.
   **Known correct:** Rejection preserved headroom for admitted work.

Failure example: if every rejected client retries immediately, offered
load grows. The retry response therefore includes delay and attempt
bounds rather than inviting an immediate loop.

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

### How it works

A database is a busy records room with limited desks, locked folders,
old copies, and a backup vault. Connections, locks, row versions,
replicas, migrations, and restore time are all limited resources that
need owners and measurements.

### Step-by-step examples

Input: transaction A transfers from account 1 to 2 while B transfers
from 2 to 1. Expected output: both lock accounts in order `1,2`, so
one waits instead of forming a deadlock cycle.

1. **Before:** Neither transaction owns a lock. **Decision:** Each
   sorts its account IDs as `[1,2]`. **After:** Both first request the
   lock for account 1. **Known correct:** Opposite transfer direction
   no longer creates opposite lock order.
2. **Before:** A owns account 1; B waits for account 1.
   **Decision:** A locks account 2, applies both changes, and commits.
   **After:** A releases both locks. **Known correct:** No transaction
   held account 2 while waiting for account 1.
3. **Before:** B can now lock accounts 1 and 2. **Decision:** Re-read
   the current balances and run its whole transaction. **After:** B
   commits or returns a business conflict. **Known correct:** Waiting
   did not permit a stale partial retry.

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

### How it works

Write a job on a durable ticket before promising it will be done. A
worker leases a small batch, renews the lease while working, records
attempts, and makes repeated effects safe. Jobs that always fail move
to a reviewed holding area.

### Step-by-step examples

Input: job `J7` has effect key `E7`, lease duration 30 seconds, and a
three-attempt limit. Worker A crashes after producing the effect but
before completion. Expected output: worker B completes `J7` after
lease expiry without producing `E7` twice.

1. **Before:** `J7` is ready with no lease. **Decision:** A atomically
   claims attempt 1 until `T+30s`. **After:** Only A may complete.
   **Known correct:** One current owner is recorded.
2. **Before:** A has stored effect `E7` but has not marked the job
   complete. **Decision:** A crashes; no worker may claim before
   `T+30s`. **After:** The expired lease makes `J7` ready again.
   **Known correct:** A crash delayed progress but did not lose work.
3. **Before:** B owns attempt 2 and finds `E7` already stored.
   **Decision:** Treat the effect as complete and finalize `J7` under
   B's lease. **After:** The job is complete with one effect.
   **Known correct:** Lease recovery and effect idempotency cover
   separate failure risks.

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

### How it works

When an order is saved, place its dispatch slip in the same locked
transaction. An **outbox** publishes that slip later. Consumers use
the event ID to avoid repeating effects and must tolerate late or
reordered delivery.

### Step-by-step examples

Input: placing order `O-42` creates outbox event `E9`; the relay
delivers `E9` twice. Expected output: the order commits with its event
and the consumer applies one semantic effect.

1. **Before:** Neither `O-42` nor `E9` is committed. **Decision:**
   Insert both in one database transaction. **After:** Both become
   visible together. **Known correct:** A committed order has a
   durable event.
2. **Before:** `E9` is unpublished. **Decision:** The relay publishes
   it, then crashes before recording progress. **After:** The broker
   may deliver `E9` again. **Known correct:** The system exposes the
   duplicate instead of pretending the first delivery was certain.
3. **Before:** The consumer has already stored inbox key `E9` and its
   effect. **Decision:** A second delivery attempts the same unique
   key. **After:** The transaction performs no second effect.
   **Known correct:** At-least-once delivery produced one result.

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

### How it works

Coordinating offices by mail is different from sharing one desk:
messages can be late, repeated, or lost. State the consistency and
availability needed for each operation, then use stable IDs,
increasing versions, leader or quorum rules, and reconciliation.
Timing alone cannot create exactly-once delivery.

### Step-by-step examples

Input: worker A holds lease token `7`, pauses past expiry, and resumes
after worker B obtains token `8`. Expected output: storage accepts B's
write and rejects A's later write.

1. **Before:** Storage's latest accepted token is `7`; A is paused.
   **Decision:** Expire A's lease and issue the next monotonic token
   `8` to B. **After:** B is the current owner. **Known correct:**
   Reassignment did not assume A had stopped running.
2. **Before:** B writes with token `8`. **Decision:** Storage compares
   `8 >= 7` and accepts the mutation. **After:** Its latest token is
   `8`. **Known correct:** The current owner advanced the state.
3. **Before:** A resumes with token `7`. **Decision:** Storage sees
   `7 < 8` and rejects the mutation. **After:** B's value remains.
   **Known correct:** A stale process cannot regain authority merely
   by continuing to run.

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

### How it works

A growing company does not create a department for every desk. Group
work by business capability, rules, data, team ownership, and change
patterns. Prove the boundary inside one modular application before
turning it into a separate service.

### Step-by-step examples

Input: the monolith owns invoice `I-42`; a candidate billing service
must reproduce its read result before receiving writes.
Expected output: mismatches block cutover; verified traffic moves in
stages while exactly one system owns each write.

1. **Before:** The monolith owns all invoice writes. **Decision:**
   Send shadow reads through a facade to both paths, but return only
   the monolith result. **After:** Differences are
   recorded without creating two writers. **Known correct:** Shadowing
   cannot mutate production ownership.
2. **Before:** A mismatch for `I-42` is visible. **Decision:** Keep
   routing at zero percent and correct the candidate. **After:** The
   same request produces matching results.
   **Known correct:** Evidence controls cutover.
3. **Before:** Verification passes and the facade owns routing.
   **Decision:** Move a bounded cohort to the service with one writer
   selected per invoice. **After:** Traffic can expand or return to
   the monolith. **Known correct:** Every stage retains single-write
   ownership and a tested reversal.

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

### How it works

Calling another service is like phoning another office while a
customer waits. Set a time limit, know whether retrying is safe, limit
how many calls may wait, classify failures, and record what happened.
A fallback changes product behavior and must be designed as such.

### Step-by-step examples

Input: the caller has 800 ms remaining. Local response work reserves
200 ms; dependency calls and backoff may use at most 500 ms.
Expected output: one safe retry succeeds within the caller deadline.

```text
caller budget:       800 ms
local reserve:      -200 ms
dependency budget:  -500 ms
safety margin:       100 ms
```

1. **Before:** The dependency budget is 500 ms. **Decision:** Attempt
   1 fails after 150 ms, then back off for 50 ms. **After:**
   300 ms remains for dependency work. **Known correct:** The retry
   has not borrowed the local reserve.
2. **Before:** 300 ms remains and the operation is idempotent.
   **Decision:** Start the one allowed retry with that smaller
   deadline. **After:** Attempt 2 succeeds in 180 ms.
   **Known correct:**
   Dependency work used 380 ms, within its 500 ms budget.
3. **Before:** 420 ms remains on the caller deadline. **Decision:**
   Finish local response work. **After:** The call completes within
   800 ms. **Known correct:** A third attempt cannot consume the
   reserved completion time.

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

### How it works

Booking a trip may reserve a flight now and a hotel later. Store each
step durably, advance one local action at a time, and undo only
actions with a real business reversal. Show users when the workflow is
still pending or has failed.

### Step-by-step examples

Input: workflow `W42` starts in `pay`; payment succeeds, inventory
reservation fails, and refund succeeds. Expected output: durable state
moves through compensation to `canceled`, with one charge and refund.

```text
state:  pay -> reserve -> refund -> canceled
effect:       charge     refund
```

1. **Before:** `W42.state=pay` and no charge exists. **Decision:**
   Commit the payment effect key and next state `reserve`. **After:**
   One charge exists and inventory work is durable. **Known correct:**
   A crash can replay without charging twice.
2. **Before:** `W42.state=reserve` and inventory is unavailable.
   **Decision:** Record rejection and commit next state `refund`.
   **After:** The order remains visibly pending compensation.
   **Known correct:** Partial success was not shown as complete.
3. **Before:** `W42.state=refund` and the charge exists. **Decision:**
   Apply the idempotent refund and store `canceled`. **After:** One
   refund exists and the workflow is terminal. **Known correct:** Each
   external effect has matching durable workflow state.

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

### How it works

Change a doorway while the shop stays open: add the wider path first,
let old and new customers pass, and test the expectations of each
consumer. Measure use of the old path and remove it only after every
known consumer has moved. This is additive contract evolution.

### Step-by-step examples

Input: old producer P1 emits `{amount: 500}`; new producer P2 may emit
`{amount: 500, currency: "USD"}`. Old consumer C1 ignores unknown
fields; new consumer C2 defaults missing currency to `"USD"`.
Expected output: all four producer-consumer pairs remain valid.

1. **Before:** P1 and C1 use only `amount`. **Decision:** Add optional
   `currency` to P2 without changing `amount`. **After:** C1 reads P2
   by ignoring the new field. **Known correct:** The old consumer's
   interpretation is unchanged.
2. **Before:** C2 receives P1 data with no currency. **Decision:** Use
   the documented `"USD"` default. **After:** C2 produces the same
   value as it would for P2's explicit field. **Known correct:** The
   new consumer remains compatible with the old producer.
3. **Before:** All four contract cases pass, but C1 still has traffic.
   **Decision:** Keep the optional field and default. **After:** No
   removal occurs until usage reaches zero. **Known correct:** Safe
   deletion depends on observed consumers, not release age.

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

### How it works

Follow one delivery with the same tracking number across every depot.
Correlation and trace context provide that number. Measure the
user-visible promise, then run small failure drills with a stop rule
and proof that the system recovered.

### Step-by-step examples

Input: checkout baseline success is 99.95%. Add 200 ms payment latency
to one canary worker with a 99.5% abort guardrail. Expected output:
the experiment either stays above 99.5% or restores baseline.

1. **Before:** Baseline is measured and only the named canary is in
   scope. **Decision:** Enable the 200 ms delay. **After:** Linked
   traces identify delayed payment attempts. **Known correct:** The
   blast radius matches the approved target.
2. **Before:** Success falls to 99.4%. **Decision:** The automated
   guardrail disables injection. **After:** No new attempt receives
   added latency. **Known correct:** The experiment stopped on the
   stated user-visible threshold.
3. **Before:** Injection is off. **Decision:** Observe checkout until
   success returns to its baseline range. **After:** Recovery evidence
   and traces are stored. **Known correct:** Ending the fault was not
   mistaken for proving restoration.

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

### How it works

A cache is a convenient copy of the store's real price list, not a new
source of truth. Say how fresh it must be and how it is replaced.
Split data by an ownership key, and include the tenant identity in
every storage and cache lookup.

### Step-by-step examples

Input: tenants A and B both request product `P1`; only A's value is
cached. Expected output: A receives A's product, B misses and loads
B's product, and neither tenant can observe the other's value.

1. **Before:** Cache contains `(A,P1) -> "A tea"`. **Decision:** A
   looks up the full compound key. **After:** A receives `"A tea"`.
   **Known correct:** The cached tenant matches the principal.
2. **Before:** B requests `P1`; `(B,P1)` is absent. **Decision:** Do
   not fall back to a product-only key; query storage with tenant B.
   **After:** B loads `"B tea"` and caches it under `(B,P1)`.
   **Known correct:** A's cache entry was never eligible for B.
3. **Before:** Both compound entries exist. **Decision:** Invalidate
   only `(A,P1)` after A changes its product. **After:** B's entry
   remains valid. **Known correct:** Lookup, storage, and invalidation
   use the same ownership dimension.

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

### How it works

Protect a system like a building: a **threat model** lists valuables,
entrances, and likely misuse. Collect only needed data, give each
actor the smallest useful privilege, and verify build inputs and
signed artifacts through the **software supply chain**. Make deletion,
key rotation, and incident response actions that can actually be run.

### Step-by-step examples

Input: key v3 signs existing artifacts; key v4 is ready for rotation.
Artifacts remain valid for 24 hours. Expected output: new artifacts
use v4, old v3 artifacts verify during overlap, then v3 is retired.

1. **Before:** Signing and verification both use v3. **Decision:**
   Publish v4 to verifiers without changing the signing key.
   **After:** Verifiers accept v3 and v4. **Known correct:** Adding
   trust did not
   invalidate existing artifacts.
2. **Before:** Both keys verify. **Decision:** Switch new signatures
   to v4 while retaining v3 verification for 24 hours. **After:** New
   artifacts use v4; old artifacts still work. **Known correct:** The
   rotation has a usable rollback window.
3. **Before:** No valid v3 artifact remains. **Decision:** Revoke v3
   and verify that v3 signatures fail. **After:** Only v4 is trusted.
   **Known correct:** Retirement happened after measured expiry, not
   immediately after publishing the new key.

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

### How it works

Keep infrastructure as reviewed blueprints instead of hand-built
rooms. Separate environments and failure areas deliberately. Give each
workload a short-lived keycard that opens only the resources and
actions required for its job.

### Step-by-step examples

Input: workload identity `order-worker` requests `Receive` on queue Q
and `Read` on bucket B using a short-lived credential.
Expected output: queue access succeeds and bucket access is denied.

1. **Before:** The workload has no embedded cloud key. **Decision:**
   Exchange its deployment identity for a short-lived credential.
   **After:** The credential names `order-worker` and its expiry.
   **Known correct:** No static secret is stored.
2. **Before:** Policy allows `Receive` on Q. **Decision:** Authorize
   the queue request against workload, action, and resource.
   **After:** One message is returned. **Known correct:** The action
   matches the declared need.
3. **Before:** The same identity requests `Read` on B. **Decision:**
   No policy statement matches that action and resource. **After:**
   Access is denied and audited. **Known correct:** Successful queue
   access did not imply broad cloud permission.

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

### How it works

Run compute like a delivery fleet. Each unit declares the fuel and
time it needs, reports when it is ready, and can be replaced without
losing cargo. Scaling and maintenance act only on signals the workload
can really honor.

### Step-by-step examples

Input: worker pod P1 holds job `J7` when it receives termination with
a 30-second grace period. Expected output: P1 admits no new work and
either completes or releases `J7` before exit.

1. **Before:** P1 is ready, owns `J7`, and has 30 seconds remaining.
   **Decision:** Mark readiness false before beginning drain.
   **After:** The scheduler stops routing new work to P1.
   **Known correct:** The
   draining workload cannot grow through normal admission.
2. **Before:** P1 is unready and still owns `J7`. **Decision:** Finish
   its durable effect and acknowledge it before the grace deadline.
   **After:** `J7` is complete and P1 may exit. **Known correct:** The
   acknowledgement follows durable completion.
3. **Before:** Replacement pod P2 becomes ready. **Decision:** Count
   it toward rollout capacity only after its startup checks pass.
   **After:** The scheduler restores the desired replica count.
   **Known correct:** Replacement and drain preserve minimum capacity.

Failure example: if `J7` cannot finish in time, P1 releases or lets
its lease expire; it does not acknowledge unfinished work.

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

### How it works

Run a service like a train line. A **service-level indicator**
measures what riders experience; a **service-level objective** sets
the target; and the allowed shortfall is an **error budget**. Alerts,
runbooks, incident reviews, and cost should all connect to that same
demand and service promise.

### Step-by-step examples

Input: checkout consumes 40% of its monthly error budget in one hour
during a rollout. Expected output: the fast-burn alert pages, rollout
stops, traffic reverts, and recovery is verified against the SLO.

1. **Before:** The release is active and the one-hour burn reaches
   40%. **Decision:** Page the on-call owner using the fast-burn rule.
   **After:** One incident and commander own mitigation.
   **Known correct:** The page represents user-visible budget use.
2. **Before:** The release is the leading correlated change.
   **Decision:** Stop rollout and route traffic to the last healthy
   version. **After:** New exposure ends. **Known correct:** The
   rollback follows a documented mitigation.
3. **Before:** Traffic has reverted. **Decision:** Observe success and
   latency until they return to the SLO range. **After:** Recovery,
   timeline, and evidence are recorded. **Known correct:** Follow-up
   starts from measured conditions rather than personal blame.

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

### How it works

Treat each outside system like a delivery company with its own forms
and failures. A narrow adapter checks the sender and package, saves
important work, prevents repeated effects, isolates bad input, and
later compares results with the authoritative source.

### Step-by-step examples

Input: a current, correctly signed webhook `E9` arrives twice with the
same raw bytes. Expected output: both deliveries receive the same
acceptance and one durable processing job exists.

1. **Before:** `E9` is unseen and only raw request bytes are trusted
   for signature input. **Decision:** Check timestamp and signature
   before parsing. **After:** The authentic payload may be decoded.
   **Known correct:** Parsing did not change the bytes being verified.
2. **Before:** No durable record exists for `E9`. **Decision:** Claim
   its unique ID and enqueue work in one durable transaction.
   **After:** Return the documented acceptance.
   **Known correct:** Acknowledgement follows durable ownership.
3. **Before:** The second delivery has the same valid ID and payload.
   **Decision:** Read the existing claim instead of enqueueing again.
   **After:** Return the same acceptance with one job still present.
   **Known correct:** Repeated transport produced one semantic effect.

Failure example: altered bytes or a stale timestamp are rejected
before any event identity is claimed.

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

### How it works

Renovate an occupied house one room at a time. Observe current
behavior, write tests that pin it down, create a safe seam, move one
small path, compare results, and remove the old path only after a
stable period.

### Step-by-step examples

Input: legacy pricing returns `$12` for product `P1`; the first
candidate returns `$13`. Expected output: comparison blocks cutover
until both return `$12`, then a bounded cohort moves behind a facade.

1. **Before:** Only legacy pricing serves users. **Decision:** Record
   `$12` and other boundary outcomes as characterization cases.
   **After:** The current observable contract is executable.
   **Known correct:** Replacement starts from observed behavior.
2. **Before:** The facade returns legacy `$12` and shadows candidate
   `$13`. **Decision:** Record and classify the mismatch without
   routing users to the candidate. **After:** Cutover remains at zero.
   **Known correct:** Comparison cannot silently change production.
3. **Before:** The corrected candidate returns `$12` across the cases.
   **Decision:** Route a small cohort through the facade. **After:**
   Matching traffic can expand or revert. **Known correct:** The old
   path remains until the defined operational window passes.

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

Shadowing duplicates selected work. Backfill and comparison visit all
`n` records and require rate limits, so they take `O(n)` work.
Temporary adapters and dual paths increase code and operational load,
so each needs an owner and removal
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

### How it works

Lead a large change like a construction project. State the goal,
constraints, and rules that cannot break; compare real options; divide
work into reversible milestones; name owners; and close with launch
evidence and lasting documentation.

### Step-by-step examples

Input: checkout timeout rate is 2%. Candidates are query
optimization, more capacity, or one bounded queue. Expected output:
a reviewed choice with guardrails, stages, and rollback ownership.

1. **Before:** The symptom is measured but its cause is unconfirmed.
   **Decision:** Profile queries, saturation, and queue wait before
   selecting an option. **After:** Evidence shows burst admission is
   the limiting condition. **Known correct:** The proposal separates
   evidence from assumptions.
2. **Before:** Three options and their costs are visible.
   **Decision:** Choose the bounded queue with latency, capacity, and
   cost limits;
   record why the other options do not fit. **After:** Each risk and
   rollback trigger has an owner. **Known correct:** Reviewers can
   challenge a real tradeoff rather than a predetermined answer.
3. **Before:** A canary receives the first bounded cohort.
   **Decision:** Expand only while timeout and queue-wait guardrails
   hold. **After:** The rollout reaches its next milestone or reverts.
   **Known correct:**
   Delivery state follows evidence defined before launch.

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

### How it works

Moving code to another language is like moving a recipe to a different
kitchen. Keep the promised result and safety rules, but use the new
kitchen's real tools. Map Go's values, errors, contexts, and
goroutines to the target runtime instead of translating syntax word
for word.

### Step-by-step examples

Input: a database call would take 2 seconds, but its request is
canceled after 100 ms.
Expected output: in both languages, the call stops, returns the same
canceled outcome, and releases work.

1. **Before:** Go passes request context into the database driver.
   **Decision:** Cancel at 100 ms. **After:** The driver stops and the
   operation returns its stable canceled result. **Known correct:**
   The conformance case records outcome and resource lifetime.
2. **Before:** The target implementation has an owned cancellation
   token. **Decision:** Pass it into the target driver and cancel at
   100 ms. **After:** The driver returns the mapped canceled result.
   **Known correct:** A timer did not merely stop the caller from
   observing still-running work.
3. **Before:** Both executions have ended. **Decision:** Compare
   result, deadline behavior, and released resources. **After:** The
   conformance test passes. **Known correct:** Equivalent semantics,
   not similar syntax, define the successful port.

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
