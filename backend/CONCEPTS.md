# Backend Engineering in Go

This guide owns the active backend teaching material. The lessons use
one evolving Order application programming interface (API) to connect
transport contracts, data safety, architecture, and production
operations.

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
- [Cloud and identity/access management
  (IAM)](#26-cloud-infrastructure-and-iam)
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
  codes, filesystem, memory, central processing unit (CPU), threads,
  file descriptors, and sockets. Containers isolate resources; they
  do not remove them.
- Treat reading as orientation. Competence requires implementation,
  negative tests, concurrency tests, operational evidence, and an
  explanation of why the result stays correct during failure.
- Separate confirmed facts, assumptions, open questions, and deferred
  scope. A plausible guess must not silently become a public contract.
- Prefer one reproducible path from a clean clone to a running
  service. Local and continuous integration (CI) commands should
  exercise the same formatter, analysis, tests, race checks,
  vulnerability checks, build, and smoke tests.
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
- Strings are immutable bytes, normally Unicode Transformation Format
  8-bit (UTF-8). Choose deliberately among byte counts, code points,
  and grapheme-like visible characters; arbitrary byte truncation can
  split encoded text.
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
  true independently of Hypertext Transfer Protocol (HTTP), Go, or
  Structured Query Language (SQL).
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

- Trace a request through Domain Name System (DNS), Transmission
  Control Protocol (TCP), Transport Layer Security (TLS), a proxy or
  load balancer, the Go server, application code, database, and
  downstream clients. Know which layer owns each timeout, pool,
  identity check, and error.
- Treat HTTP method, target, headers, representation, status, caching,
  pagination, limits, ordering, and retry semantics as one public
  contract. Do not expose persistence rows as that contract.
- Handlers translate rather than decide business policy: bound and
  decode input, validate the transport envelope, call one use case,
  and map its known outcomes to stable responses.
- Enforce media type, body size, unknown-field rejection, and exactly
  one JavaScript Object Notation (JSON) value. Client-side validation
  never makes server-side validation optional.
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
  exact decimal values or integer minor units, event time as a
  Coordinated Universal Time (UTC) instant, and civil scheduling with
  its named time zone.
- Historical facts keep the values used at the time. For example, an
  order item stores the charged price rather than depending on the
  product's current price.
- An index trades write work and storage for a particular lookup or
  ordering. Verify the actual query plan and data shape; redundant or
  unused indexes still impose costs.
- Parameter binding separates data from SQL. Object-relational tools
  do not remove the need to inspect generated queries, row counts,
  lock behavior, or the N+1 query pattern: one initial query followed
  by one additional query for every returned item.
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
  high-cardinality values such as user identifiers (IDs) and raw
  uniform resource locators (URLs) out of metric labels.
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
  Require Hypertext Transfer Protocol Secure (HTTPS), reject private
  and link-local destinations after resolution, revalidate redirects,
  defend against DNS rebinding, and restrict runtime egress.
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
A stock-keeping unit (SKU) is the product's stable catalog identifier.

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

Focused lookup input: prices contain SKU `FREE` with price `0` but do
not contain SKU `MISSING`. Expected output: `LookupPrice("FREE")`
returns `(0, true)` and `LookupPrice("MISSING")` returns `(0, false)`.

1. **Before:** `FREE` is present with a deliberate zero price.
   **Decision:** Read the map value and presence result. **After:**
   Return `(0, true)`. **Known correct:** Zero remains valid data.
2. **Before:** `MISSING` is absent. **Decision:** Read both map return
   values. **After:** Return `(0, false)`. **Known correct:** The
   presence result distinguishes absence from the zero value.

### Why it works

The fact that must stay true is that no mutable backing storage
crosses an ownership boundary unintentionally. This ownership
invariant is true after `Clone` allocates independent storage. Every
later mutation reaches only the selected owner's data, so the fact
remains true for the lifetime of both values.

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

Let `t` be the number of tags. Cloning visits and stores every tag, so
it takes `O(t)` time and `O(t)` extra memory. Copying the outer struct
takes constant work. A map lookup usually examines a constant amount
of map state, so it takes `O(1)` average time and no growing extra
memory. For a string containing `b` bytes, conversion to `[]rune`
visits those bytes and stores the decoded code points, taking `O(b)`
time and `O(b)` worst-case extra memory. Simple byte indexing does not
allocate.

### When to use it

- A mutable collection crosses a package or goroutine boundary.
- A cache or stored value outlives the call that supplied it.
- An API must distinguish a missing value from a stored zero value.

### When not to use it

Do not add pointers, interfaces, generics, or deep copies without an
ownership, substitution, or mutation need. Plain values are easier to
reason about when their size and lifetime are small.

### Common mistakes and edge cases

> **Warning:** After `copyOfTags := product.Tags`, changing
> `copyOfTags[0]` also changes `product.Tags[0]` when both slices
> share one backing array. Allocate and copy the elements when the
> receiver needs independent ownership.

> **Warning:** `var prices map[string]int; prices["tea"] = 4` panics
> because a nil map cannot accept writes. Initialize the map before a
> write, and sort copied keys when output order must be stable.

> **Warning:** Cutting `"caf\u00e9"` at an arbitrary byte position can
> split the final encoded character. Decode code points or use a
> text-boundary library that matches the required unit.

> **Warning:** Returning `(*CatalogError)(nil)` as an `error` creates
> a non-nil interface because it still carries a dynamic type. Return
> a literal nil interface when no error occurred.

> **Warning:** `result, err := load()` inside a nested block can
> create a new `err`; the outer `err` then remains unchanged. Assign
> with `=` when the existing error variable owns the later check.

### Related material

See [Go backend foundations][go-foundations],
[Go concurrency](#3-go-concurrency-and-executable-verification), and
[language portability](#32-language-portability).

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

Focused cleanup example: `open` returns a reader containing `"tea"`;
its `Close` method increments `closeCount`. Expected output:
`data="tea"`, `err=nil`, and `closeCount=1`.

1. **Before:** No reader is open and `closeCount=0`. **Decision:**
   `ReadProduct` calls `open`, checks the error, and immediately
   schedules `Close`. **After:** The reader is owned by
   `ReadProduct`. **Known correct:** Every successful acquisition now
   has one release path.
2. **Before:** The reader is open. **Decision:** `io.ReadAll` reads
   `"tea"`, then the function returns. **After:** The deferred close
   runs and `closeCount=1`. **Known correct:** Both success and an
   early read failure pass through the same cleanup path.

### Why it works

The fact that must stay true is that each acquired resource has one
owner and one release path. This ownership invariant holds because a
defer is arranged only after acquisition succeeds. Deferred calls run
from the most recently arranged call backward, which matches nested
acquisition. Cancellation is cooperative: each blocking layer
observes the same signal and returns. Error wrapping preserves the
causal chain while adding the operation that failed.

### Go example

```go
package catalog

import (
	"context"
	"errors"
	"fmt"
	"io"
)

var ErrNotFound = errors.New("product not found")

func ReadProduct(
	open func() (io.ReadCloser, error),
) (data []byte, err error) {
	reader, err := open()
	if err != nil {
		return nil, fmt.Errorf("open product: %w", err)
	}
	// The function owns reader after a successful open, so every return
	// path must close it and preserve a meaningful close failure.
	defer func() {
		closeErr := reader.Close()
		if err == nil && closeErr != nil {
			err = fmt.Errorf("close product: %w", closeErr)
		}
	}()
	data, err = io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("read product: %w", err)
	}
	return data, nil
}

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

Valid product strings are non-empty, and `results` remains open until
it sends a result. The ordinary path returns that product. The failure
paths retain either absence or cancellation identity through the
wrapper. `ReadProduct` gives an opened reader one cleanup path.

### Work and memory

The selection itself is constant work. Waiting time is bounded by the
result or context lifetime. Each wrapper and child context adds small
constant memory; held resources can be far more expensive, so release
latency belongs in the complexity discussion.

### When to use it

- Input/output (I/O), database, or network work acquires a resource.
- A lock, queue, timer, or remote call can outlive the caller's wait.
- A caller must classify a wrapped failure by its original identity.

### When not to use it

Do not add context to instant pure functions, store it in a struct, or
use it as a general parameter bag. Do not recover ordinary operational
errors as panics.

### Common mistakes and edge cases

> **Warning:** `fmt.Errorf("load: %v", context.Canceled)` produces
> text that `errors.Is` cannot traverse. Use `%w` when callers must
> retain the original error identity.

> **Warning:** Scheduling `file.Close()` before checking a failed open
> can call a method on an invalid file. Check acquisition first, then
> arrange cleanup and handle any meaningful close error.

> **Warning:** If row iteration returns two rows and then a protocol
> error, returning those two rows without checking `rows.Err()` turns
> a failed query into partial success. Check the terminal iterator
> error before publishing the result.

> **Warning:** Deferring each close in a loop over 10,000 files
> retains all successful opens until the outer function returns. Put
> one iteration in a helper so each defer runs before the next
> acquisition.

> **Warning:** Replacing a canceled request context with
> `context.Background()` lets downstream work continue after the
> caller leaves. Pass the request context through every blocking
> layer.

### Related material

See [Go values and ownership](
#1-go-values-program-structure-and-ownership),
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
decision. A **mutex** is a one-at-a-time lock around that complete
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

The fact that must stay true is that a stock-keeping unit (SKU) is
inserted at most once. This uniqueness invariant holds for the empty
store. Mutual exclusion makes check-and-insert one indivisible
transition, so a successful insert preserves uniqueness. Joining every
goroutine makes the test terminate only after all outcomes are known.

### Go example

```go
package catalog

import (
	"errors"
	"sync"
	"testing"
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

func (store *Store) Size() int {
	store.mu.Lock()
	defer store.mu.Unlock()
	return len(store.products)
}

func TestStoreCreateContention(t *testing.T) {
	store := NewStore()
	start := make(chan struct{})
	outcomes := make(chan error, 2)
	var workers sync.WaitGroup

	// Make both workers eligible from one deterministic signal. The
	// production mutex, rather than a test sleep, chooses the winner.
	workers.Add(2)
	for workerNumber := 0; workerNumber < 2; workerNumber++ {
		go func() {
			defer workers.Done()
			<-start
			outcomes <- store.Create("TEA-001")
		}()
	}
	close(start)
	workers.Wait()

	// Outcome order is irrelevant; the business result is exactly one
	// success, one duplicate, and one stored product.
	successCount := 0
	duplicateCount := 0
	for resultNumber := 0; resultNumber < 2; resultNumber++ {
		resultErr := <-outcomes
		switch {
		case resultErr == nil:
			successCount++
		case errors.Is(resultErr, ErrDuplicate):
			duplicateCount++
		default:
			t.Fatalf("unexpected error: %v", resultErr)
		}
	}
	if successCount != 1 || duplicateCount != 1 {
		t.Fatalf(
			"got %d successes and %d duplicates",
			successCount,
			duplicateCount,
		)
	}
	if store.Size() != 1 {
		t.Fatalf("got %d products, want 1", store.Size())
	}
}
```

The test releases two workers from one start channel, joins them with
`sync.WaitGroup`, and checks the complete business outcome.

### Work and memory

One map access usually examines a constant amount of map state, so it
takes `O(1)` average time. Lock wait depends on contention and the
critical-section duration. Let `n` be the number of stored SKUs. The
map retains one entry per SKU, so it uses `O(n)` memory. Each live
goroutine also consumes stack and scheduler resources, which is why
concurrency must be bounded.

### When to use it

- Several goroutines read and change one in-process business rule.
- A bounded worker group needs explicit start, stop, and join signals.
- A test must reproduce contention without depending on sleep timing.

### When not to use it

An in-memory mutex cannot coordinate several processes and cannot
replace a durable transaction. Do not add goroutines to work that is
already serial, tiny, or governed by a stricter external limit.

### Common mistakes and edge cases

> **Warning:** If A unlocks after finding no `TEA-001` and B performs
> the same unlocked check before either insert, both calls can report
> success. Hold one mutex across the complete check-and-insert rule.

> **Warning:** Starting one goroutine per item for a million-item
> input creates a million stacks and scheduler tasks. Use a worker
> limit and join every success, cancellation, and error path.

> **Warning:** If a receiver closes `outcomes` while a worker can
> still send, the worker panics. The sender that knows no more values
> can arrive should close the channel, or the coordinator should wait
> for every sender first.

> **Warning:** A buffer of 100 entries still blocks the 101st send
> when no receiver drains it. Define queue capacity and overload
> behavior instead of treating buffering as unlimited storage.

> **Warning:** Atomically changing `reserved` without coordinating
> `available` can expose a pair that violates their shared total. Use
> one lock around every value in the business invariant.

> **Warning:** A 10 ms sleep can pass on one machine and assert before
> work finishes on another. Use signals and joins; race-detector
> silence proves only that it found no memory race in that execution.

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
without a state change; `C-17` changes the order to `canceled` and
records exactly one `OrderCanceled` event.

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

The fact that must stay true is that only the owning customer may
perform a legal order transition. This authorization invariant holds
because every command names its actor, target, and expected state. The
transition table covers every current state and reaches a terminal
accepted or rejected result, so there is no implicit path. Recording
the event only after the state check also prevents a rejected command
from announcing a change.

### Go example

Assume order and customer identifiers are non-empty, and all status
changes go through methods that enforce the transition rules.

```go
package order

import "errors"

var ErrCannotCancel = errors.New("order cannot be canceled")

type Status string
type Event string

const (
	Placed        Status = "placed"
	Canceled      Status = "canceled"
	OrderCanceled Event  = "order_canceled"
)

type Order struct {
	CustomerID string
	Status     Status
	Events     []Event
}

func (order *Order) Cancel(customerID string) error {
	// Authorization is checked before lifecycle state so neither
	// rejection can mutate the order or announce an event.
	if order.CustomerID != customerID {
		return ErrCannotCancel
	}
	if order.Status != Placed {
		return ErrCannotCancel
	}
	order.Status = Canceled
	order.Events = append(order.Events, OrderCanceled)
	return nil
}
```

The ordinary case performs one legal transition. Wrong ownership or
state returns a stable domain error and leaves both status and events
unchanged.

### Work and memory

A direct status check and update takes constant work. Appending one
event usually takes constant work, but it can allocate when the slice
grows. Let `e` be the number of retained events. The event history
uses `O(e)` result memory; the transition itself needs only constant
extra working memory. Discovery work grows with the number of distinct
stories and stakeholders, not code size.

### When to use it

- Requirements use lifecycle words such as placed or canceled.
- A command depends on ownership, permissions, money, or stock.
- Another system must react to a named fact after a legal transition.

### When not to use it

Do not invent aggregates, events, or repositories for a mechanical
transformation with no durable business rule. A plain function is
often the honest model.

### Common mistakes and edge cases

> **Warning:** Giving a shipping address an identity when only its
> fields matter creates needless lifecycle rules. Model it as a value
> object unless the product must track that exact address over time.

> **Warning:** Publishing `OrderCanceled` before the transaction
> commits can announce a cancellation that later rolls back. Record
> publication intent with the business change and publish after
> commit.

> **Warning:** An aggregate that requires one atomic change across two
> independently owned services cannot enforce its own rule. Keep the
> immediate invariant under one owner or model a durable workflow.

> **Warning:** If the model specifies only `placed -> canceled`, a
> second cancel has no defined result. Name the rejected outcome and
> test that status and events remain unchanged.

> **Warning:** A diagram that omits the `C-99` rejection does not
> prove the ownership rule. Replay concrete accepted and rejected
> stories in executable tests.

### Related material

See [service construction](
#5-architecture-and-service-construction),
[contested state](#8-business-transactions-and-contested-state), and
[distributed workflows](
#21-distributed-workflows-and-eventual-consistency).

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
second contains quantity `2`; a scheduled job submits `O-43` with
quantity `0`. Expected output: both invalid commands return
`ErrInvalidQuantity` without storage, while the valid command calls
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
3. **Before:** A scheduled job owns command `ID=O-43, Quantity=0`.
   **Decision:** Call the same use case without an HTTP handler.
   **After:** The job receives `ErrInvalidQuantity`, and `Save` still
   has one total call. **Known correct:** Both adapters use the same
   business rule.

### Why it works

The fact that must stay true is that business decisions depend only on
inward-facing contracts. This dependency invariant holds because
adapters translate at the edge and the composition root is the only
place that knows the complete graph. Each dependency can therefore
change or be replaced without reversing the policy boundary.

### Go example

Assume a trusted boundary supplies non-empty order identifiers. The
use case owns quantity validation and domain-value construction.

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

type PlaceCommand struct {
	ID       string
	Quantity int
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
	command PlaceCommand,
) error {
	if command.Quantity <= 0 {
		return ErrInvalidQuantity
	}
	// Translate the accepted command into the domain value expected by
	// the use-case-owned storage contract.
	order := Order{
		ID:       command.ID,
		Quantity: command.Quantity,
	}
	return service.store.Save(ctx, order)
}
```

The interface belongs to the use case. A rejected order does not reach
the store; a valid order delegates exactly once.

### Work and memory

The quantity check, two-field translation, and extra function call do
a fixed amount of local work and retain one small domain value. The
storage operation still dominates variable latency and resource use.
For larger payloads, translation work and memory grow with the fields
actually copied. Duplicate models and broad interfaces add that cost
without improving isolation.

### When to use it

- Durable business policy must stay independent of HTTP and storage.
- HTTP, jobs, or commands invoke the same business operation.
- Infrastructure changes independently of the rules it supports.

### When not to use it

Do not add an interface for every struct or split one deployable into
network services merely to imitate an architecture diagram.

### Common mistakes and edge cases

> **Warning:** If a PostgreSQL package defines a broad repository
> interface, every consumer inherits storage-shaped methods it does
> not need. Define the smallest interface beside the consuming use
> case.

> **Warning:** If the HTTP handler saves `Quantity=0` before calling
> `Place`, transport code has bypassed the rule. Let one use case own
> validation and its transaction boundary.

> **Warning:** Opening a database from package initialization makes a
> unit test depend on hidden configuration and startup order. Build
> concrete dependencies explicitly in the composition root.

> **Warning:** Returning a nullable database row as the public order
> response leaks storage representation into the API. Translate into
> an intentional transport contract at the edge.

> **Warning:** A catalog package that updates the orders table can
> violate order rules without calling the owner. Modularity requires
> exclusive mutation through the owning package contract.

### Related material

See [domain modeling](#4-domain-discovery-and-modeling),
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

Input: authenticated customer `C-17` sends `POST /v1/orders` with
media type `application/json` and body
`{"sku":"TEA-001","quantity":2}`. Expected output: status `201` and
body `{"id":"O-42","sku":"TEA-001","quantity":2}`.

1. **Before:** The server has untrusted bytes and headers.
   **Decision:** Enforce the body limit and expected media type.
   **After:** Oversized or incompatible input is rejected early.
   **Known correct:** Later decoding cannot exceed the stated bound.
2. **Before:** The decoded fields are `sku=TEA-001` and `quantity=2`.
   **Decision:** Validate required fields and authorize `C-17` to
   create the order. **After:** One typed command reaches the use
   case. **Known correct:** Transport and access failures cannot enter
   business logic.
3. **Before:** The use case returns order `O-42`. **Decision:** Map
   that created result to the documented status and representation.
   **After:** The client receives the exact `201` response above.
   **Known correct:** The public result contains no internal error or
   persistence details.

Failure input: the request body is
`{"sku":"TEA-001","quantity":2} {"quantity":9}`. Expected output:
status `400`, machine code `invalid_json`, and no use-case call.

1. **Before:** The decoder has the full two-value body. **Decision:**
   Decode the first value into the request shape. **After:** The first
   object is decoded, but bytes for a second object remain.
   **Known correct:** The request is not accepted until the decoder
   proves that input has ended.
2. **Before:** One complete object remains unread. **Decision:**
   Decode once more and require an end-of-file (EOF) result.
   **After:** The second value makes that check fail. The handler
   returns `invalid_json`. **Known correct:** No trailing value is
   ignored and no command reaches the use case.

Focused shutdown example: input is a termination signal with a
five-second shutdown deadline while request `R1` has 200 ms of work
left. Expected output: no new request is admitted, `R1` finishes, and
owned dependencies close before five seconds pass.

1. **Before:** Readiness is true, the listener is accepting, and `R1`
   is active. **Decision:** Set readiness false and stop new accepts.
   **After:** Readiness advertises draining, the listener admits no
   new request, and `R1` remains active. **Known correct:** The local
   draining set cannot grow through normal admission.
2. **Before:** `R1` has 200 ms left and the five-second deadline
   remains. **Decision:** Wait for `R1` before closing its
   dependencies. **After:** `R1` finishes, then the server closes
   owned resources. **Known correct:** Accepted work finished within
   the bounded shutdown interval.

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
The destination must be a non-nil decode target, and the response
writer and request must belong to the current HTTP exchange.

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

- A browser or service caller needs stable success and error shapes.
- An authenticated write needs bounded decoding and authorization.
- Pagination, uploads, retries, or shutdown expose lifecycle rules.

### When not to use it

Do not expose a generic persistence interface as an API contract. Do
not use JSON request-response semantics for a sustained bidirectional
stream when a streaming protocol is the actual requirement.

### Common mistakes and edge cases

> **Warning:** Cross-Origin Resource Sharing (CORS) allows selected
> browser origins to read a response. It does not authenticate or
> authorize the caller; enforce identity and action permission on the
> server.

> **Warning:** Returning status `200` with `{"error":"not_found"}`
> makes callers choose between the status and body. Use one documented
> status and one stable error representation for that outcome.

> **Warning:** Repeating an unkeyed order-creation request after a
> timeout can create two orders. Give one logical write a durable
> idempotency identity before clients retry it.

> **Warning:** A liveness check that fails whenever PostgreSQL is down
> can restart every healthy process in a loop. Keep liveness local;
> expose dependency inability through readiness and telemetry.

> **Warning:** With offset `20`, an insertion before the current page
> can shift one earlier record into the next response. Use a stable
> total order and a cursor for mutation-prone collections.

> **Warning:** Logging the raw login body exposes credentials to every
> log reader and sink. Record bounded, allowlisted metadata instead.

> **Warning:** Starting `ListenAndServe` in an unobserved goroutine
> can hide an immediate bind failure while the process appears
> healthy. Return the error to the lifecycle owner.

> **Warning:** Waiting forever for `R1` can block a deployment
> indefinitely. Set a shutdown deadline, propagate its cancellation,
> and report work that did not drain.

### Related material

- [Architecture and service construction](
  #5-architecture-and-service-construction)
- [Frontend network and security][frontend-network]
- [Data, transactions, and architecture][data-lesson]

[frontend-network]:
  ../frontend/CONCEPTS.md#a3-networking-storage-and-browser-security
[data-lesson]: #8-business-transactions-and-contested-state

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
   the same unique constraint. **After:** A claims the candidate
   unique value while B waits for A's outcome. **Known correct:** B
   cannot commit a duplicate while the conflict is unresolved.
2. **Before:** A owns the candidate value and B is waiting.
   **Decision:** A commits; PostgreSQL then rechecks B against the
   committed row. **After:** B's `ON CONFLICT DO NOTHING` returns no
   row, which the adapter maps to `ErrDuplicate`. **Known correct:**
   The table contains one row, the SKU remained data rather than SQL
   syntax, and both writers faced the same integrity rule.

### Why it works

The fact that must stay true is that every stored product has required
fields and one unique SKU. This integrity invariant holds because
constraints run in the same atomic statement as the write, so every
writer faces the rule. Parameter binding separates data from query
structure, and explicit cleanup returns scarce connections.

### Go example

Assume `products.sku` is unique and non-null, `products.name` is
non-null, and the configured PostgreSQL driver uses `$1` parameters.

```go
package catalog

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

var ErrDuplicate = errors.New("duplicate product")

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

func Create(
	ctx context.Context,
	database *sql.DB,
	product Product,
) (Product, error) {
	const query = `
		INSERT INTO products (sku, name)
		VALUES ($1, $2)
		ON CONFLICT (sku) DO NOTHING
		RETURNING sku, name
	`
	// Binding keeps both fields as data. The unique constraint remains
	// authoritative when concurrent callers use the same SKU.
	row := database.QueryRowContext(
		ctx,
		query,
		product.SKU,
		product.Name,
	)
	created, err := ScanProduct(row)
	if errors.Is(err, sql.ErrNoRows) {
		return Product{}, ErrDuplicate
	}
	if err != nil {
		return Product{}, fmt.Errorf("create product: %w", err)
	}
	return created, nil
}
```

The bound parameters are ordinary data. `sql.ErrNoRows` has one local
meaning because this statement returns no row only when the unique
conflict path does nothing. Other database errors keep their cause.

### Work and memory

Let `n` be the number of indexed rows and `r` the number returned. A
B-tree is a balanced search index; following its levels commonly
takes `O(log n)` lookup work plus row access. Without a useful index,
visiting all rows takes `O(n)` work. Collecting results retains `O(r)`
result memory, while streaming can keep extra application memory
bounded but holds a connection until all `r` rows are consumed. An
index also adds write work, storage, cache pressure, and vacuum
effort.

### When to use it

- Facts must survive process and machine restarts.
- Keys, checks, relationships, or uniqueness enforce shared rules.
- Queries need joins, transactions, or operational recovery.

### When not to use it

Do not add an index without a measured access path, wrap one atomic
statement in a needless transaction, or hide fields in an opaque
document when those fields govern integrity and lookup.

### Common mistakes and edge cases

> **Warning:** Two processes can both pass an application-only
> `sku not found` check before either inserts `TEA-001`. Keep the
> unique constraint as the final shared integrity rule.

> **Warning:** `SELECT *` can change scan order after a column is
> added, placing a name into the wrong destination. Select named
> columns in the same order that `Scan` expects.

> **Warning:** Returning early after the second streamed row without
> closing `rows` can retain its connection; ignoring `rows.Err()` can
> publish partial success after an iteration failure. Close and check.

> **Warning:** Concatenating SKU `"x' OR true --"` into a query
> changes query structure. Bind it as a parameter, and keep sensitive
> bound values out of logs.

> **Warning:** Dropping `legacy_price` while an old application still
> reads it breaks mixed-version deployment. Add compatible state,
> migrate and verify readers, then remove it after the rollback
> window.

> **Warning:** Raising each of 20 process pools to 100 connections can
> offer 2,000 sessions to a database sized for far fewer. Budget the
> total pool, observe waits, and prove restore separately from backup.

### Related material

See [architecture](#5-architecture-and-service-construction),
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
   **After:** A's transaction sees stock `0` and holds the row lock;
   B's competing update waits. **Known correct:** No committed stock
   value is negative.
2. **Before:** A has uncommitted stock `0`, its row lock, and no order
   yet. **Decision:** Insert A's order and commit the transaction.
   **After:** Stock `0` and A's order become visible together; B can
   continue. **Known correct:** No partial successful order was
   published.
3. **Before:** B resumes after A commits. **Decision:** PostgreSQL
   rechecks B's condition against stock `0`. **After:** B affects no
   row and returns `ErrInsufficient`. **Known correct:** Exactly one
   reservation exists and final stock is `0`.

Failure input: stock is `1`, A's conditional decrement succeeds, and
A's order insert returns an error. Expected output: stock `1`, no
order, and the insert error.

1. **Before:** A's transaction sees stock `1`. **Decision:** Apply the
   conditional decrement. **After:** A's transaction sees stock `0`,
   but no other transaction can see that uncommitted value.
   **Known correct:** The incomplete reservation is not public.
2. **Before:** Stock `0` is uncommitted and the order insert fails.
   **Decision:** Return the error and roll back the transaction.
   **After:** Committed stock is still `1` and no order exists.
   **Known correct:** Failure removed every partial change.

### Why it works

The conditional write makes the inventory rule part of the same atomic
database decision as the change. The transaction hides partial state
and either publishes both records or neither. The use case owns this
boundary because repositories cannot coordinate a business action by
committing independently. Assuming the SKU identifies one inventory
row and both statements use the same transaction, every path ends in
commit or rollback.

### Go example

Assume `inventory.sku` is unique, `orders.id` is unique, and positive
quantities are required. Both statements run through one transaction.

```go
package inventory

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

var ErrInsufficient = errors.New("insufficient inventory")

func Reserve(
	ctx context.Context,
	database *sql.DB,
	orderID string,
	sku string,
	quantity int,
) (err error) {
	if quantity <= 0 {
		return errors.New("quantity must be positive")
	}
	transaction, err := database.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin reservation: %w", err)
	}
	// Every return after BeginTx must close the transaction. Rollback
	// after Commit returns sql.ErrTxDone and needs no further action.
	defer func() {
		rollbackErr := transaction.Rollback()
		if errors.Is(rollbackErr, sql.ErrTxDone) {
			return
		}
		if rollbackErr != nil {
			err = errors.Join(
				err,
				fmt.Errorf("rollback reservation: %w", rollbackErr),
			)
		}
	}()

	const updateInventory = `
		UPDATE inventory
		SET quantity = quantity - $1
		WHERE sku = $2 AND quantity >= $1
	`
	result, err := transaction.ExecContext(
		ctx,
		updateInventory,
		quantity,
		sku,
	)
	if err != nil {
		return fmt.Errorf("reserve inventory: %w", err)
	}
	changedRows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("count reserved inventory: %w", err)
	}
	if changedRows == 0 {
		return ErrInsufficient
	}
	if changedRows != 1 {
		return fmt.Errorf("reserved %d inventory rows", changedRows)
	}

	const insertOrder = `
		INSERT INTO orders (id, sku, quantity)
		VALUES ($1, $2, $3)
	`
	if _, err := transaction.ExecContext(
		ctx,
		insertOrder,
		orderID,
		sku,
		quantity,
	); err != nil {
		return fmt.Errorf("insert order: %w", err)
	}
	if err := transaction.Commit(); err != nil {
		return fmt.Errorf("commit reservation: %w", err)
	}
	return nil
}
```

### Work and memory

For one indexed SKU, the application issues a fixed number of
statements and retains constant extra memory. Database work also pays
connection, lock-wait, write-ahead-log, index, and network costs. Let
`k` be the number of resources a larger operation must lock. Retaining
the target list uses `O(k)` memory. Sorting it takes `O(k log k)`
comparison work and gives concurrent transactions one lock order; an
in-place sort avoids a second `O(k)` target copy.

### When to use it

- One business result spans several reads or writes.
- Inventory, balances, quotas, or allocation face concurrent claims.
- A state transition must publish all changes together or none.

### When not to use it

Do not hold a database transaction open across a slow external network
call. For cross-service actions, use durable workflow and compensation
instead of pretending one local transaction covers every service.

### Common mistakes and edge cases

> **Warning:** Reading stock `1` and decrementing it in separate
> statements lets A and B both approve the last unit. Use one atomic
> condition or lock the row inside the owning transaction.

> **Warning:** An application-only `quantity >= 0` check can race with
> another writer. Keep non-negative stock and uniqueness constraints
> in the database as the final shared barrier.

> **Warning:** Retrying an aborted transaction without an attempt or
> time bound can amplify overload; retrying an unsafe effect can
> repeat it. Retry the whole operation only under bounded, idempotent
> policy.

> **Warning:** If A locks account 1 then 2 while B locks 2 then 1,
> each can wait for the other. Sort lock targets into one total order.

> **Warning:** A client timeout during `Commit` does not prove the
> transaction failed. Reconcile by a stable operation identity before
> repeating an ambiguous write.

> **Warning:** Waiting on a remote payment call while holding the
> inventory row lock blocks unrelated reservations and retains old row
> versions. Finish external work outside the short transaction.

### Related material

- [Relational modeling](#7-relational-modeling-and-postgresql-access)
- [Idempotency](#13-idempotency-and-ambiguous-write-outcomes)
- [Distributed workflows](
  #21-distributed-workflows-and-eventual-consistency)

## 9. Identity, authorization, and API security

### The problem it solves

An API must know who is acting, whether that actor may perform this
operation on this object, and whether untrusted input can escape its
intended data boundary.

### How it works

Reception verifies a badge to establish who is present; a locked door
still decides whether that person may enter. **Authentication**
establishes identity. **Authorization** checks the person, action,
resource, and situation.
Input validation, output encoding, secret handling, and browser policy
remain separate controls, so one failure does not remove every
defense.

### Step-by-step examples

Input: authenticated user `U-7` requests `DELETE /orders/O-42`; the
order is owned by `U-9`. Expected output: the same stable `not_found`
result used for an absent order, with no order mutation.

1. **Before:** The route contains `O-42` and the verified principal is
   `U-7`. **Decision:** Attempt one delete scoped by both
   `order=O-42` and `owner=U-7`. **After:** The predicate matches no
   row because the stored owner is `U-9`. **Known correct:** Knowing
   an order ID did not grant permission.
2. **Before:** The scoped delete reports no match and the order is
   unchanged. **Decision:** Map that result to the same `not_found`
   outcome as an absent order. **After:** The caller learns no private
   ownership fact. **Known correct:** Access control and
   non-enumeration both hold.

### Why it works

The fact that must stay true is that every protected state change has
an authenticated principal and a positive resource-specific
authorization decision. This access invariant holds because the
mutation predicate includes both resource and principal identities.
A non-match changes nothing and uses the same public result for an
absent or differently owned order. Independent input and output
controls limit injection and disclosure.

### Go example

Assume authentication has already produced a trusted, non-empty
principal subject. The store implements `DeleteOwned` as one atomic
delete whose predicate includes both order and owner identifiers.

```go
package orderaccess

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrNotFound        = errors.New("order not found")
	ErrUnauthenticated = errors.New("authentication required")
)

type Principal struct {
	Subject string
}

type Store interface {
	DeleteOwned(
		context.Context,
		string,
		string,
	) (bool, error)
}

func DeleteOrder(
	ctx context.Context,
	store Store,
	principal Principal,
	orderID string,
) error {
	if principal.Subject == "" {
		return ErrUnauthenticated
	}
	// Combining resource and owner in one predicate prevents a caller
	// from mutating a known identifier owned by another principal.
	deleted, err := store.DeleteOwned(
		ctx,
		orderID,
		principal.Subject,
	)
	if err != nil {
		return fmt.Errorf("delete owned order: %w", err)
	}
	if !deleted {
		return ErrNotFound
	}
	return nil
}
```

### Work and memory

Let `c` be the credential size and `n` the number of indexed orders.
Parsing visits `c` bytes, so it takes `O(c)` work and may retain
`O(c)` decoded data. Cryptographic verification has an additional
algorithm- and key-dependent cost bounded by server configuration. A
B-tree lookup for an owned order commonly adds `O(log n)` work and
constant result memory. Password hashing is intentionally expensive;
bound its concurrency so configured CPU and memory cost cannot exhaust
the process.

### When to use it

- A non-public operation changes or reveals a protected resource.
- Permission depends on both the actor and the loaded object.
- A browser, partner, file, webhook, or database supplies input.

### When not to use it

Do not invent cryptography, treat CORS as access control, or authorize
only at route level when ownership depends on the loaded object.

### Common mistakes and edge cases

> **Warning:** Letting a token select an unexpected algorithm or key
> can turn attacker-controlled metadata into verification policy. Pin
> allowed algorithms, issuers, audiences, and key sources in trusted
> server configuration.

> **Warning:** A correctly signed token for audience `billing` is not
> authority for the orders API. Check issuer, audience, time window,
> intended use, and current revocation policy together.

> **Warning:** Logging an `Authorization` header or password-reset URL
> gives log readers a usable credential. Allowlist safe telemetry
> fields and redact secrets before any sink.

> **Warning:** Returning `forbidden` for `O-42` but `not_found` for an
> absent ID reveals that the private order exists. Use one policy and
> response shape when existence itself is sensitive.

> **Warning:** Trusting `X-User-ID` from the public internet lets a
> caller choose its identity. Accept an identity header only from an
> authenticated proxy on a verified, restricted network path.

### Related material

See [HTTP contracts](#6-http-contracts-and-process-lifecycle),
[privacy and supply
chain](#25-security-privacy-and-software-supply-chain),
and [production integrations](#29-production-integration-boundaries).

## 10. Deterministic backend testing

### The problem it solves

Tests become slow or flaky when time, randomness, networks, global
state, and goroutine ordering are uncontrolled. Passing examples then
provide little evidence about failures or concurrent invariants.

### How it works

A useful test controls the clock, random choices, and outside systems
so the same setup gives the same result every time. Test the smallest
public boundary that proves the rule, then use integration tests only
where real storage or protocols matter.

### Step-by-step examples

Input: the clock always returns `2026-01-02T03:04:05Z`, the identifier
source always returns `order-123`, the repository records saves, and
the service receives quantity `2`. Expected output: order
`{ID:order-123, Quantity:2, PlacedAt:2026-01-02T03:04:05Z}` and one
recorded save of that exact order.

1. **Before:** Time and identifier generation would otherwise vary on
   each run. **Decision:** Construct the service with the fixed clock,
   fixed identifier source, and recording repository. **After:** All
   three outside choices have test-owned behavior. **Known correct:**
   The service cannot read the wall clock or generate a surprise ID.
2. **Before:** The repository is empty and quantity is `2`.
   **Decision:** Call `Place`. **After:** The returned order has the
   exact expected fields and the repository contains that one order.
   **Known correct:** The same setup produces the same public result
   and side effect on every run.

Failure input: the same service receives quantity `0` after the
successful call. Expected output: `ErrInvalidQuantity` and the
repository still contains only `order-123`.

1. **Before:** The repository contains one order. **Decision:** Call
   `Place` with quantity `0`. **After:** Validation returns
   `ErrInvalidQuantity` before time, identifier, or repository use.
   **Known correct:** The rejected call adds no side effect.

### Why it works

The fact that must stay true is that test inputs fully determine
observable results. This determinism invariant holds because every
outside choice is supplied through an explicit dependency. Assertions
check public results and side effects rather than private call order.
Concurrency tests apply the same rule with explicit signals, joins,
and bounded contexts.

### Go example

See the deterministic placement service in
[the service examples][service-examples]. Its
[tests][service-tests] inject time and identifiers, then prove invalid
input makes no repository call.

### Work and memory

Table-test work is the sum of the work for each case, while retained
fixtures and results determine memory use. Integration setup can
dominate runtime. Fuzzing and race detection multiply execution and
memory, so CI needs explicit time, worker, corpus, and package bounds.

### When to use it

- Time, identifiers, randomness, or outside systems affect an output.
- Policy, error mapping, deadlines, retries, or transitions need
  proof.
- A concurrency regression must replay without guessed sleep timing.

### When not to use it

Do not mock a query whose correctness depends on PostgreSQL behavior;
use a real PostgreSQL integration test. Do not assert private call
order when the public outcome is the real contract.

### Common mistakes and edge cases

> **Warning:** Sleeping 10 ms before an assertion can pass locally and
> run before the worker finishes under load. Signal readiness and join
> completion explicitly.

> **Warning:** Two parallel cases that append to one shared fixture
> can change each other's expected counts. Give each case independent
> state or synchronize the shared invariant.

> **Warning:** A race-detector pass can coexist with two successful
> reservations of the last item. Assert the business outcome as well
> as running the detector.

> **Warning:** A fake that reimplements PostgreSQL locking can
> disagree with the real database while tests stay green. Fake only
> the narrow boundary; use integration tests for database semantics.

> **Warning:** Testing only the successful save misses validation,
> cancellation, storage failure, retry, and cleanup paths. Add cases
> for each behavior the contract promises.

### Related material

See [service construction](#5-architecture-and-service-construction),
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
Each timed operation in a trace is a **span**; parent and child spans
record causal nesting.
Health checks, timeouts, rate limits, and profiles turn that evidence
into safe operating controls.

### Step-by-step examples

Input: request `R-8` reaches a database call whose 100 ms deadline
expires. Expected output: status `503` with code
`dependency_unavailable`; one outcome counter for route
`/v1/orders/{id}`; and correlated log and trace records without
private data.

1. **Before:** `R-8` has a request span and the database attempt has
   not started. **Decision:** Create a child span with dependency name
   and deadline. **After:** The trace shows which wait belongs to the
   request. **Known correct:** Causal context crosses the boundary.
2. **Before:** The database exceeds 100 ms. **Decision:** Classify the
   deadline as dependency unavailability. **After:** Return the stable
   `503` response and increment one counter whose route and outcome
   labels come from small fixed sets. **Known correct:** Error text
   did not become an unbounded metric label.
3. **Before:** The request has finished. **Decision:** Emit one event
   with route `/v1/orders/{id}`, request `R-8`, method `GET`, and code
   `dependency_unavailable`. **After:** Log, metric, and trace agree
   on the failure. **Known correct:** No private value entered
   telemetry.

### Why it works

The fact that must stay true is that each important operation produces
enough stable, non-secret evidence to classify its outcome. This
observability invariant holds because shared correlation preserves
causal context and bounded dimensions keep the telemetry system
operable under the same failure it is meant to explain.

### Go example

See the bounded access-event builder in
[the service examples][service-examples]. Its
[tests][service-tests] accept stable public fields and reject an
authorization field before it reaches telemetry.

### Work and memory

Serialization and export work grow with the bytes in each allowed
event. Let `q` be the export queue capacity and `b` the maximum event
size. The queue retains at most `O(q * b)` bytes. Metric-series count
grows with combinations of label values, while trace storage grows
with sampled span volume and retention. Bound labels, queues, event
size, sampling, profiles, and retention so telemetry backpressure
cannot stop request work.

### When to use it

- A production service, worker, or scheduler needs diagnosis.
- One operation crosses process, queue, or dependency boundaries.
- An operator needs a safe control for overload or recovery.

### When not to use it

Do not log every payload, create per-user metric labels, or make
liveness depend on a remote service whose outage a restart cannot
repair. Prefer bounded fields and aggregate dimensions.

### Common mistakes and edge cases

> **Warning:** A log containing only `timeout` cannot identify the
> operation or affected request. Record a route template, stable
> outcome, and correlation ID without copying private payloads.

> **Warning:** Reporting ready before a required migration finishes
> can send traffic to code whose table is missing. Keep readiness
> false until required local initialization succeeds.

> **Warning:** Starting a listener or exporter in an abandoned
> goroutine can hide a bind or authentication failure. Return its
> terminal error to the lifecycle owner.

> **Warning:** An unbounded span-export queue grows during an exporter
> outage and can exhaust process memory. Cap the queue and define
> drop, sampling, and alert behavior.

> **Warning:** Paging on one brief CPU spike can wake an operator
> while users remain unaffected. Alert on sustained user impact or
> imminent exhaustion with a documented action.

### Related material

See [deterministic testing](#10-deterministic-backend-testing),
[delivery and
recovery](#12-delivery-deployment-and-recovery-readiness),
and [reliability
engineering](#28-reliability-engineering-incidents-and-cost).

## 12. Delivery, deployment, and recovery readiness

### The problem it solves

A locally correct binary can fail because its artifact is
unrepeatable, configuration is unsafe, a migration is incompatible,
rollout lacks a stop signal, or backups cannot actually restore
service.

### How it works

Seal one package, then move that same package through inspection and
into production. The package is an immutable build artifact; secrets
stay outside it. Deploy through compatible stages, watch explicit
health criteria, and practice rollback and data restore before an
emergency.

### Step-by-step examples

Input: immutable version B receives 5% of traffic; version A receives
95%; the stop threshold is a 1% B error rate; the next stage is 25%;
and B remains at 0.4% for the observation window. Expected output: B
advances to 25% and A serves 75% without a rebuild.

1. **Before:** A serves 95%, B serves 5%, and both read the expanded
   schema. **Decision:** Observe B for the defined window. **After:**
   B's measured error rate is 0.4%. **Known correct:** The comparison
   uses only the stated canary cohort and window.
2. **Before:** B is below the 1% stop threshold. **Decision:** Advance
   the existing artifact to the 25% stage. **After:** A serves 75%
   and B serves 25%. **Known correct:** The rollout followed its
   predeclared gate without rebuilding B.

Focused failure input: A serves 95%, B serves 5%, and B reaches 3%
against the 1% threshold. After rollback, A remains at 0.2% errors for
ten minutes. Expected output: B receives 0%, A receives 100%, the
rollout stays stopped, and recovery evidence records the ten-minute
healthy window.

1. **Before:** B serves 5% and its measured error rate is 3%.
   **Decision:** Apply the 1% stop threshold. **After:** Stop B and
   route all traffic to A. **Known correct:** The reversal reused the
   healthy artifact and compatible expanded schema.
2. **Before:** B receives no new traffic and A serves 100%.
   **Decision:** Observe A for ten minutes before declaring recovery.
   **After:** A remains at 0.2%, and the rollout stays stopped for
   diagnosis. **Known correct:** Recovery was measured rather than
   inferred from ending exposure.

### Why it works

The fact that must stay true is that every deployment step has an
observable acceptance condition and a compatible reversal. This
delivery invariant holds because immutable artifacts remove rebuild
variance, staged schemas preserve mixed-version operation, and restore
tests replace assumptions about recoverability with measured evidence.

### Go example

Assume traffic percentages are between 0 and 100, the current split
adds to 100, and A and B remain schema-compatible during reversal.

```go
package rollout

type Decision string

const (
	Advance  Decision = "advance"
	Rollback Decision = "rollback"
)

type State struct {
	StablePercent    int
	CandidatePercent int
	StopErrorPercent float64
	Stopped          bool
}

func Evaluate(
	state *State,
	candidateErrorPercent float64,
	nextCandidatePercent int,
) Decision {
	if state.Stopped {
		return Rollback
	}
	if candidateErrorPercent > state.StopErrorPercent {
		// A threshold breach ends candidate exposure but preserves the
		// immutable artifact and evidence for diagnosis.
		state.StablePercent = 100
		state.CandidatePercent = 0
		state.Stopped = true
		return Rollback
	}
	// A passing gate advances only to the predeclared next stage.
	state.CandidatePercent = nextCandidatePercent
	state.StablePercent = 100 - nextCandidatePercent
	return Advance
}
```

Typed startup configuration is a separate delivery boundary. See
[the service examples][service-examples] and their
[tests][service-tests] for valid startup, missing required values, and
secret-safe failure text.

### Work and memory

Let `a` be artifact bytes, `d` the dependency count, and `v` the
volume of state to restore. Building, transferring, and scanning visit
artifact data and dependency metadata, so their work grows with `a`
and `d`; retained artifacts use storage proportional to `a`. Canary
time depends on traffic needed for a useful signal. Backups use
storage proportional to `v`, while restore time includes reading,
validating, and replaying that realistic volume.

### When to use it

- One artifact moves through several environments or rollout stages.
- A release changes a schema, queue, secret, or stateful dependency.
- Recovery depends on rollback, roll-forward, backup, or restore.

### When not to use it

Do not rebuild per environment, hide configuration defaults, or claim
a rollback that cannot read the newly written schema or data
representation.

### Common mistakes and edge cases

> **Warning:** Dropping a column with the first new application
> version breaks old instances still reading it. Expand compatibly,
> migrate and verify, then contract after the rollback window.

> **Warning:** Deploying image tag `latest` can select different bytes
> in two environments. Promote one immutable artifact by digest and
> verify its source and build provenance.

> **Warning:** Reporting ready before the listener and required local
> initialization succeed sends traffic to a process that cannot serve.
> Keep readiness false until those preconditions hold.

> **Warning:** Replacing a verification secret immediately rejects
> requests signed just before rotation. Support an explicit overlap
> window, then retire the old secret after measured expiry.

> **Warning:** A backup that has never restored may be incomplete or
> too slow for the recovery objective. Time a realistic restore and
> verify both data and application invariants.

### Related material

See [observability](#11-observability-and-operational-controls),
[orchestration](#27-orchestration-and-platform-operations), and
[incidents and cost](#28-reliability-engineering-incidents-and-cost).

## 13. Idempotency and ambiguous write outcomes

### The problem it solves

A client can time out after a server commits. Retrying may duplicate a
charge or order, while refusing to retry may hide a successful result.

### How it works

A payment retry should behave like showing the same receipt number,
not buying twice. An **idempotency key** identifies one logical write.
A **fingerprint** is a stable digest of every request field that
changes the operation's meaning. Atomically store the scoped key,
fingerprint, and durable result, so an identical retry replays the
result while conflicting reuse is rejected.

### Step-by-step examples

Input: customer `C-17` sends create-order command
`{customer_id:"C-17", quantity:2}` with key `K1`; its first response
is lost and the client retries. Expected output: both attempts return
`O-42`; one order exists.

1. **Before:** `K1` is absent. **Decision:** In one transaction, bind
   scope `C-17/create-order/K1` to the command fingerprint, create
   `O-42`, store its response, and commit. **After:** The order and
   completed key record are durable. **Known correct:** The semantic
   effect and replay information cannot commit separately.
2. **Before:** The client saw a timeout, but the completed record
   exists. **Decision:** Retry with `K1` and the same command.
   **After:** Return stored order `O-42` without creating an order.
   **Known correct:** Ambiguous delivery did not duplicate the effect.

Focused conflict input: `C-17` reuses `K1` with
`{customer_id:"C-17", quantity:3}`. Expected output: key conflict,
the original `O-42` response remains stored, and no second order.

1. **Before:** `K1` is bound to the quantity-2 fingerprint.
   **Decision:** Compare it with the canonical quantity-3 command.
   **After:** The fingerprints differ, so return a key conflict.
   **Known correct:** One scoped key cannot name two logical writes.

### Why it works

The fact that must stay true is one durable semantic outcome per
scoped key and fingerprint. This idempotency invariant holds because a
uniqueness constraint serializes competing claims and the transaction
records the business result with the key. Correct comparison assumes
canonicalization includes every semantic field and the digest has
adequate collision resistance for the system's threat model.

### Go example

The typed command gives JavaScript Object Notation fields a stable
order before Secure Hash Algorithm 256-bit (SHA-256) hashing. See
canonical command fingerprinting in
[the production examples][production-examples]. Its
[tests][production-tests] distinguish a new operation, a compatible
replay, and conflicting key reuse.

### Work and memory

Let `n` be retained key records, `c` canonical request bytes, and `b`
stored response bytes. Fingerprinting visits `c` bytes, taking `O(c)`
work. A B-tree key lookup commonly takes `O(log n)` work. Each record
retains its key, digest, metadata, and `b` response bytes, so total
storage grows with request volume, response size, and retention.
Cleanup needs bounded batches and must not race the supported retry
window.

### When to use it

- A client may retry a non-idempotent write after a timeout.
- A payment-like effect needs one semantic result per operation.
- A committed response can be lost before the caller receives it.

### When not to use it

Do not treat in-memory locks, request IDs, or a read-before-write
check as durable idempotency across processes and restarts. A
naturally idempotent replacement write may need no stored response
record.

### Common mistakes and edge cases

> **Warning:** Using bare key `K1` for every customer lets one tenant
> collide with another tenant's write. Scope the unique key by trusted
> principal and operation.

> **Warning:** Reusing `K1` after quantity changes from `2` to `3` can
> replay the wrong order. Compare a canonical fingerprint and reject
> changed input under an existing key.

> **Warning:** Committing the order before its key record can leave a
> successful order with no replay evidence after a crash. Commit the
> business effect, fingerprint, and response in one transaction.

> **Warning:** Deleting the record after one hour while clients retry
> for 24 hours can turn a late retry into a second order. Derive
> retention and cleanup from the documented retry window.

> **Warning:** A broker or network can still deliver the same request
> more than once. Claim one semantic effect within the scoped
> boundary; do not describe the transport itself as exactly once.

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

Focused retry input: 100 rejected clients each allow one retry.
Without delay they retry immediately; with policy they receive delays
spread from 100 to 200 ms and retry only if their deadline still
permits it. Expected output: the unsafe policy creates 100 immediate
new attempts; the bounded policy creates none immediately and at most
100 later.

1. **Before:** All 100 clients have one retry left and capacity is
   still full. **Decision:** Retry immediately without a delay check.
   **After:** The service receives 100 new attempts, rejects them, and
   completes no useful work. **Known correct:** The trace exposes one
   full repeated wave caused only by retry policy.
2. **Before:** The same 100 clients receive distinct bounded delays.
   **Decision:** Each client waits and retries only if its remaining
   deadline funds the delay and call. **After:** No attempt is
   immediate and no client can exceed one later retry.
   **Known correct:** Delay, deadline, and attempt bounds cap
   amplification.

### Why it works

The fact that must stay true is that active and queued work never
exceed proved capacity. This admission invariant holds because the
capacity check and count change occur under one lock. Bounds preserve
memory and scheduling headroom. Deadlines stop obsolete work, and
randomized bounded retries prevent synchronized amplification.

### Go example

Assume non-negative limits, valid initial counts, and exactly one
`Complete` call for every active operation that finishes.

```go
package admission

import "sync"

type Result string

const (
	Active   Result = "active"
	Queued   Result = "queued"
	Rejected Result = "rejected"
)

type Capacity struct {
	mu            sync.Mutex
	active        int
	maximumActive int
	queued        int
	maximumQueued int
}

func NewCapacity(
	maximumActive int,
	maximumQueued int,
) *Capacity {
	if maximumActive < 0 || maximumQueued < 0 {
		panic("capacity limits must be non-negative")
	}
	return &Capacity{
		maximumActive: maximumActive,
		maximumQueued: maximumQueued,
	}
}

func (capacity *Capacity) Admit() Result {
	capacity.mu.Lock()
	defer capacity.mu.Unlock()
	if capacity.active < capacity.maximumActive {
		capacity.active++
		return Active
	}
	if capacity.queued < capacity.maximumQueued {
		capacity.queued++
		return Queued
	}
	return Rejected
}

func (capacity *Capacity) Complete() bool {
	capacity.mu.Lock()
	defer capacity.mu.Unlock()
	if capacity.active == 0 {
		panic("completion without active work")
	}
	if capacity.queued > 0 {
		// Promote one queued operation while keeping the active count
		// unchanged; the completed slot transfers to that operation.
		capacity.queued--
		return true
	}
	capacity.active--
	return false
}
```

`Complete` reports whether the scheduler should start one queued
operation in the slot that just became free.

See the bounded retry-delay policy in
[the production examples][production-examples]. Its
[tests][production-tests] cover one allowed retry and an exit when the
remaining deadline cannot fund another call.

### Work and memory

Let `q` be the queue limit, `w` the worker limit, and `a` the maximum
attempts per logical call. Queue metadata and payload references use
`O(q)` memory, while worker stacks and per-call state use `O(w)`
memory. At most `w + q` operations are admitted. Retry policy can
offer up to `a` calls per logical operation, so worst-case downstream
load and latency must include that factor.

### When to use it

- A worker pool or queue has a measured capacity ceiling.
- A remote dependency or endpoint slows under offered load.
- Retries and deadlines must fit one end-to-end objective.

### When not to use it

Do not optimize from intuition, retry permanent errors, or use an
unbounded queue to hide insufficient capacity. Measure the bottleneck
and reject predictably when accepted work cannot finish safely.

### Common mistakes and edge cases

> **Warning:** A 20 ms average can hide a 2 s slow tail while workers
> and queues saturate. Measure latency distributions, utilization,
> saturation, and completed throughput together.

> **Warning:** Giving each of three dependencies the caller's full
> one-second deadline lets sequential work consume three seconds.
> Reserve completion time and derive smaller child budgets.

> **Warning:** Three retries at both proxy and service layers can
> turn one call into nine downstream attempts. Assign retry ownership
> to one layer and cap total attempts and time.

> **Warning:** Opening one global circuit for a single tenant's
> failure rejects healthy tenants too. Scope isolation and circuit
> state to the resource or dependency that actually fails.

> **Warning:** Parsing a 50-mebibyte body before checking admission
> spends memory and CPU on work that will be rejected. Enforce cheap
> size and capacity limits before expensive decoding or allocation.

### Related material

See [observability](#11-observability-and-operational-controls),
[synchronous communication](#20-synchronous-service-communication),
and [Entry 28](#28-reliability-engineering-incidents-and-cost).

## 15. Advanced PostgreSQL operation and recovery

### The problem it solves

Correct queries still fail operationally through pool exhaustion,
deadlocks, table bloat, replication lag, unsafe migrations, or
untested recovery.

### How it works

A database is a busy records room with limited desks, locked folders,
old copies, and a backup vault. Connections, locks, row versions,
replicas, migrations, and restore time are all limited resources that
need owners and measurements.

### Step-by-step examples

Input: accounts 1 and 2 each hold `100`; transaction A transfers `10`
from 1 to 2 while B transfers `20` from 2 to 1. Expected output: both
lock accounts in order `[1,2]`, both commit, and final balances are
account 1=`110` and account 2=`90`.

1. **Before:** Neither transaction owns a lock. **Decision:** Each
   sorts its account IDs as `[1,2]`. **After:** Both first request the
   lock for account 1. **Known correct:** Opposite transfer direction
   no longer creates opposite lock order.
2. **Before:** A owns account 1; B waits for account 1.
   **Decision:** A locks account 2, applies both changes, and commits.
   **After:** Balances are `[90,110]`, and A releases both locks.
   **Known correct:** No transaction held account 2 while waiting for
   account 1.
3. **Before:** B can now lock accounts 1 and 2. **Decision:** Re-read
   balances `[90,110]` and transfer `20` from 2 to 1. **After:** B
   commits balances `[110,90]`. **Known correct:** Waiting did not
   permit a stale partial retry.

Focused recovery input: a 02:00 base backup contains account balance
`100`; archived log records `+20` at 02:05, `-5` at 02:10, and an
accidental delete at 02:13. The recovery target is 02:12. Expected
output: the account exists with balance `115`; the delete is absent.

1. **Before:** The recovery database is empty. **Decision:** Restore
   the 02:00 base backup. **After:** The account exists with balance
   `100`. **Known correct:** Replay starts from the backup's exact
   state.
2. **Before:** The next archived records are `+20` and `-5`.
   **Decision:** Replay them in log order through 02:10. **After:**
   The balance is `115`. **Known correct:** Every change before the
   target has been applied once in order.
3. **Before:** The next record is the 02:13 delete. **Decision:** Stop
   replay at the 02:12 target and validate the account invariant.
   **After:** The account still exists with balance `115`.
   **Known correct:** Recovery did not replay the later deletion.

### Why it works

The fact that must stay true during a transfer is that every
transaction acquires shared locks in one total order and retries the
whole transaction. This lock-order invariant removes cyclic wait.
Short transaction lifetimes release locks and old row versions. For
recovery, ordered log replay from a known backup reconstructs every
change through the chosen target; validation proves the restored state
before traffic returns.

### Go example

Assume a transfer involves two distinct positive account identifiers.

```go
package accountlocks

func OrderedAccountIDs(
	firstAccountID int64,
	secondAccountID int64,
) (int64, int64) {
	if firstAccountID < secondAccountID {
		return firstAccountID, secondAccountID
	}
	return secondAccountID, firstAccountID
}
```

The transfer uses the returned order for every row lock, independent
of transfer direction. The configured PostgreSQL retry classifier in
[the production examples][production-examples] and its
[tests][production-tests] accept serialization failure and reject a
permanent constraint error.

### Work and memory

Let `n` be lock targets, `p` the connection limit per process, and `i`
the process count. Retaining targets uses `O(n)` memory; sorting them
takes `O(n log n)` comparison work, while two targets need only one
comparison. The fleet can offer up to `p * i` database sessions.
Indexes, retained row versions, backups, and archived logs add
storage, write amplification, cache pressure, replay work, and vacuum
work.

### When to use it

- Concurrent transactions lock more than one shared resource.
- Pool limits, migrations, replicas, or bloat affect production.
- Recovery has explicit data-loss and restore-time objectives.

### When not to use it

Do not increase pool size to mask slow queries or route
consistency-sensitive reads to a lagging replica. Fix the measured
bottleneck or read from the authoritative source.

### Common mistakes and edge cases

> **Warning:** A 20-connection pool across 100 instances can offer
> 2,000 sessions even when the database limit is 500. Budget the fleet
> total and bound acquisition wait.

> **Warning:** Retrying only the final statement after PostgreSQL
> aborts a transaction keeps earlier work invalid. Roll back and retry
> the entire safe operation under an attempt and time bound.

> **Warning:** Holding account locks while waiting for a user or
> remote call blocks other transfers and retains old row versions.
> Keep the transaction short and move external waits outside it.

> **Warning:** Adding a non-null column with a table-wide rewrite can
> block production traffic. Inspect lock level and duration, use a
> compatible staged migration, and rehearse cancellation or recovery.

> **Warning:** A backup can omit required logs or restore too slowly.
> Restore it to an isolated environment, replay to a target, and
> verify application and data invariants before relying on it.

### Related material

See [relational data](#7-relational-modeling-and-postgresql-access),
[transactions](#8-business-transactions-and-contested-state), and
[delivery and
recovery](#12-delivery-deployment-and-recovery-readiness).

## 16. Durable background jobs and scheduling

### The problem it solves

In-process goroutines disappear on restart and cannot safely
coordinate work across replicas. Jobs may be leased twice, abandoned,
or retried forever.

### How it works

Write a job on a durable ticket before promising it will be done. A
worker leases a small batch, renews the lease while working, records
attempts, and receives a unique attempt token that stale workers
cannot reuse. Repeated effects need their own idempotency identity.
Jobs that always fail move to a reviewed holding area, often called a
**dead-letter state**.

### Step-by-step examples

Input: durable job `J7` is `leased` with current token `8` and attempt
count `2`; paused worker A resumes with stale token `7`, then current
worker B completes with token `8`. Expected output: A receives
`ErrStaleClaim`, B succeeds, and `J7.state=completed`.

1. **Before:** `J7.state=leased`, `J7.token=8`, and A holds token `7`.
   **Decision:** A calls `Complete(7)`. **After:** The token mismatch
   returns `ErrStaleClaim`; state and token remain unchanged.
   **Known correct:** A paused former owner cannot finalize B's lease.
2. **Before:** `J7` remains leased with token `8`. **Decision:** B
   calls `Complete(8)`. **After:** State becomes `completed`.
   **Known correct:** Only the current claim finalized the job.

Focused exhaustion input: job `J8` is `leased` with token `2`, attempt
count `3`, and maximum attempts `3`; its failure is retryable.
Expected output: `Fail(2, true, 3)` succeeds and `J8.state=dead`.

1. **Before:** `J8` is on its final allowed attempt. **Decision:**
   Record the retryable failure with the current token. **After:** The
   state becomes `dead` rather than `ready`. **Known correct:** The
   job cannot enter a fourth attempt.

### Why it works

The fact that must stay true is that only the current lease token may
finalize or retry a job. This claim invariant holds because each
transition compares the supplied token with durable state before
mutation. Lease expiry makes abandoned work eligible for a new token,
and the attempt limit moves repeated failure to a terminal state.
Effect idempotency remains a separate invariant across those attempts.

### Go example

See the fenced job transition in
[the production examples][production-examples]. Its
[tests][production-tests] reject a stale claimant, accept the current
claim, and terminate exhausted retryable work.

### Work and memory

Let `n` be retained jobs, `w` the worker count, and `b` the maximum
claim batch. A B-tree lookup for ready work commonly takes `O(log n)`
work. At most `w * b` jobs are leased by one worker fleet pass, so
live payload and worker memory must be bounded at that scale. History
storage grows with attempts and retention; more frequent heartbeats
reduce recovery delay but add database writes.

### When to use it

- Work must survive process, host, or deployment failure.
- Slow, retryable, scheduled, or fan-out work need not block a
  caller.
- Operators need visible attempts, cancellation, replay, or
  quarantine.

### When not to use it

Do not enqueue work whose required state has not committed, or use a
durable queue for trivial synchronous work whose result the caller
requires now.

### Common mistakes and edge cases

> **Warning:** Marking `J7` complete before its external effect is
> durable can lose that effect after a crash. Persist or reconcile the
> effect before acknowledging completion.

> **Warning:** Accepting `Complete(7)` after token `8` was issued lets
> a paused stale worker overwrite the current claim. Compare the
> attempt token on every completion, extension, and failure
> transition.

> **Warning:** Retrying a malformed email address forever consumes
> every attempt without a possible success. Classify permanent input
> failure and move it to reviewed terminal state.

> **Warning:** A new worker that cannot decode payload version 1 can
> strand old jobs during deployment. Version payloads and keep readers
> compatible through the queue's retention window.

> **Warning:** Starting 1,000 workers against a dependency limited to
> 50 calls creates local queueing and remote overload. Bound workers
> from downstream capacity and observe queue age as well as length.

### Related material

See [reliable events](#17-reliable-event-publication-and-consumption),
[idempotency](#13-idempotency-and-ambiguous-write-outcomes), and
[orchestration](#27-orchestration-and-platform-operations).

## 17. Reliable event publication and consumption

### The problem it solves

Writing business state and publishing a message as separate operations
can lose the event or publish an event for a transaction that later
rolls back. Consumers also receive duplicates and out-of-order
messages.

### How it works

When an order is saved, place its dispatch slip in the same locked
transaction. An **outbox** publishes that slip later. Consumers use
an **inbox**, a durable record of handled event IDs stored with the
local effect, to avoid repetition. They must also detect late or
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

Focused ordering input: order projection `O-42` is at version `4`;
event version `6` arrives before version `5`. Expected output: version
`6` waits as a gap, then versions `5` and `6` apply in order, leaving
the projection at version `6`.

1. **Before:** `currentVersion=4` and event `6` arrives. **Decision:**
   Detect that expected version `5` is missing. **After:** Keep the
   projection at `4` and retain event `6` for bounded reconciliation.
   **Known correct:** The consumer did not invent the missing state.
2. **Before:** `currentVersion=4`, event `6` is pending, and event `5`
   arrives. **Decision:** Apply the expected event `5`. **After:**
   `currentVersion=5`, and event `6` remains pending.
   **Known correct:** The missing transition is now present.
3. **Before:** `currentVersion=5` and pending event `6` is contiguous.
   **Decision:** Apply event `6`. **After:** `currentVersion=6` and no
   event is pending. **Known correct:** Delivery order did not change
   transition order.

### Why it works

The producer fact that must stay true is that business state and
publication intent commit together. This is the producer invariant.
The consumer fact is one local effect per event ID; that is the
consumer invariant. Local transactions and unique constraints enforce
both despite relay crashes. Version checks prevent a later fact from
silently skipping required earlier state.

### Go example

See the idempotent consumer transition in
[the production examples][production-examples]. Its
[tests][production-tests] apply the first event and ignore a duplicate
without repeating the business effect.

### Work and memory

Let `n` be retained outbox or inbox records and `b` the event bytes.
Each event adds producer and consumer durable writes plus broker
transfer of `b` bytes. A B-tree claim or deduplication lookup commonly
takes `O(log n)` work. Durable storage grows with event volume, event
size, consumer count, and retention. Pending gap buffers and cleanup
batches need explicit bounds that preserve replay and audit windows.

### When to use it

- Another component must observe a committed local fact eventually.
- Producer state and broker publication cannot share one transaction.
- Consumers can receive duplicate, delayed, or reordered deliveries.

### When not to use it

Do not publish speculative commands as committed facts or promise
global ordering when only per-key partition order exists. A direct
local call is simpler when no durable asynchronous boundary is needed.

### Common mistakes and edge cases

> **Warning:** Marking `E9` sent before broker acknowledgement can
> lose it if the process crashes before publish. Mark progress only
> after acknowledgement and tolerate a repeated publish.

> **Warning:** Applying a balance change before committing inbox key
> `E9` can repeat the change after a crash. Commit the key and local
> effect in one consumer transaction.

> **Warning:** Remembering `E9` only in a process map forgets it after
> restart and allows a duplicate effect. Store deduplication state for
> the full replay window.

> **Warning:** Reinterpreting `amount` from cents to dollars under the
> same event version changes old facts. Add a compatible field or new
> version with an owned migration policy.

> **Warning:** Broker acknowledgement proves broker acceptance, not
> that every consumer applied the event. Observe consumer lag, errors,
> dead letters, and reconciliation separately.

### Related material

See [distributed foundations][distributed-reference],
[durable jobs](#16-durable-background-jobs-and-scheduling), and
[contract evolution](#22-contract-evolution-and-independent-delivery).

## 18. Distributed-systems mechanics

### The problem it solves

Across a network, delay is indistinguishable from failure, clocks
disagree, messages repeat, and replicas observe different histories.
Local reasoning no longer proves a global outcome.

### How it works

Coordinating offices by mail is different from sharing one desk:
messages can be late, repeated, or lost. State the consistency and
availability needed for each operation, then use stable IDs,
increasing versions, and explicit coordination. A **quorum** is the
minimum replica group whose responses decide an operation.
**Reconciliation** compares divergent state and applies a named rule
to converge it. Timing alone cannot create exactly-once delivery.

### Step-by-step examples

Input: worker A holds lease token `7`, pauses past expiry, and resumes
after worker B obtains token `8`. B writes value `"new"`; A later
writes `"stale"`. Expected output: storage keeps value `"new"` with
latest token `8` and rejects A's later write.

1. **Before:** Storage's latest accepted token is `7`; A is paused.
   **Decision:** Expire A's lease and issue the next monotonic token
   `8` to B. **After:** B is the current owner. **Known correct:**
   Reassignment did not assume A had stopped running.
2. **Before:** B writes `"new"` with token `8`. **Decision:** Storage
   compares `8 >= 7` and accepts the mutation. **After:** Value is
   `"new"`, and latest token is `8`. **Known correct:** The current
   owner advanced the protected state.
3. **Before:** A resumes and writes `"stale"` with token `7`.
   **Decision:** Storage sees `7 < 8` and rejects the mutation.
   **After:** Value remains `"new"`, and latest token remains `8`.
   **Known correct:** A stale process cannot regain authority merely
   by continuing to run.

### Why it works

The fact that must stay true is that accepted mutations carry a token
no lower than the latest accepted token. This fencing invariant holds
when token allocation is durable and strictly increasing and the
protected resource stores and checks its latest token. Timeouts affect
progress but cannot grant stale authority.

### Go example

See the storage-side fencing check in
[the production examples][production-examples]. Its
[tests][production-tests] accept a newer owner and reject a paused
stale holder after reassignment.

### Work and memory

Let `r` be replica count, `q` the required quorum responses, and `k`
the tracked keys. Sending a write to every replica uses `O(r)` network
and replica work and retains `O(r)` copies. Quorum latency waits for
the `q`th response among those replicas, not necessarily the slowest.
Version, deduplication, and reconciliation metadata use at least
`O(k)` storage and grow further with retained history.

### When to use it

- State or work crosses a process, machine, region, or broker
  boundary.
- A timeout can leave the outcome uncertain rather than simply
  failed.
- Replicas, leases, or messages can disagree, repeat, or arrive late.

### When not to use it

Do not introduce consensus, distributed locks, or multi-region writes
when a single transactional owner meets the real availability
requirement.

### Common mistakes and edge cases

> **Warning:** A timeout after a payment write may mean the response
> was lost after commit. Reconcile by stable operation identity before
> deciding whether to repeat the effect.

> **Warning:** A wall clock that reads 10:01 on A and 10:00 on B does
> not prove A's event caused B's. Use causal or consensus-backed order
> where the operation requires it, and monotonic time for durations.

> **Warning:** At-least-once transport can deliver `E9` twice. A local
> inbox can make one consumer effect idempotent, but it does not make
> the transport itself exactly once.

> **Warning:** A lease without fencing lets paused token `7` overwrite
> token `8` after reassignment. Make the protected resource reject
> every token below its latest accepted value.

> **Warning:** Citing the Consistency, Availability, and Partition
> tolerance (CAP) theorem without naming one operation hides the real
> choice. State what that operation returns during a network
> partition.

### Related material

See [reliable events](#17-reliable-event-publication-and-consumption),
[service communication](#20-synchronous-service-communication), and
[distributed
workflows](#21-distributed-workflows-and-eventual-consistency).

## 19. Service-boundary discovery and extraction

### The problem it solves

Splitting code along technical layers creates chatty services and
distributed transactions without establishing independent business
ownership.

### How it works

A growing company does not create a department for every desk. Group
work by business capability, rules, data, team ownership, and change
patterns. Prove the boundary inside one modular application before
turning it into a separate service. A **facade** is a stable entry
point that can route to old or candidate implementations during that
proof.

### Step-by-step examples

Input: the monolith owns invoice `I-42` and reads
`{total:500,status:"open"}`; the candidate billing service first reads
`{total:450,status:"open"}`. After correction it reads the exact
monolith value. Expected output: the mismatch keeps billing write
traffic at 0%; after matching, the facade assigns the next `I-42`
write only to billing.

1. **Before:** The monolith owns all invoice writes. **Decision:**
   Send shadow reads through a facade to both paths, but return only
   the monolith result. **After:** Record the `500` versus `450`
   difference without creating two writers. **Known correct:**
   Shadowing cannot mutate production ownership.
2. **Before:** A mismatch for `I-42` is visible. **Decision:** Keep
   billing writes at 0% and correct its projection. **After:** Both
   reads return `{total:500,status:"open"}`. **Known correct:**
   Evidence, rather than release age, controls cutover.
3. **Before:** Reads match and the facade still routes `I-42` writes
   to the monolith. **Decision:** Assign `I-42` to billing before its
   next write. **After:** The facade routes that write only to billing
   and records the ownership assignment. **Known correct:** Exactly
   one system can mutate `I-42` at this stage.

Focused seam input: the composition root first selects a local
notification adapter and sends `"placed"`; it then selects a remote
adapter that returns `publisher unavailable`. Expected output: the
local adapter records one message, while the remote call returns the
same error cause without a hidden local fallback.

1. **Before:** The local adapter is selected and has no messages.
   **Decision:** Call `Notify("placed")`. **After:** It contains one
   `"placed"` message. **Known correct:** The use case depends only on
   the narrow notification contract.
2. **Before:** The remote adapter is selected and configured to fail.
   **Decision:** Call `Notify("placed")`. **After:** The wrapped error
   retains `publisher unavailable`, and no local path runs.
   **Known correct:** Extraction did not hide the new failure mode.

### Why it works

The fact that must stay true is single ownership of each durable
mutation throughout migration. This ownership invariant holds because
the facade makes routing explicit, comparison supplies evidence, and
each assignment selects one writer. The adapter seam also exposes
remote failure instead of silently changing product behavior.

### Go example

See the composition-root-selected notification adapter in
[the distributed examples][distributed-examples]. Its
[tests][distributed-tests] cover local success and unavailable remote
delivery without hiding the failure.

### Work and memory

Let `n` be records moved and `b` the bytes in one request and
response. Shadowing can nearly double read, compute, and transfer
work. A backfill visits `n` records, so it takes `O(n)` work. Use
bounded batches rather than `O(n)` working memory. Each extracted
call adds serialization and transfer of `b` bytes, tail latency,
telemetry, and failure-state handling.

### When to use it

- One business capability has clear data and operational ownership.
- Independent scaling, delivery, or isolation has measured value.
- An internal seam already proves the boundary and failure contract.

### When not to use it

Do not split by table, controller, or repository, or extract merely
because a codebase is large. Improve internal boundaries first when
independent ownership or operations do not justify network cost.

### Common mistakes and edge cases

> **Warning:** Letting both monolith and billing update `I-42` can
> create divergent totals. Record one writer assignment and enforce it
> at the routing and data boundaries.

> **Warning:** Copying one million invoices without counts, checksums,
> or reconciliation can silently omit records. Move bounded batches
> and verify source and destination independently.

> **Warning:** Extracting billing before deciding how order placement
> and invoice creation coordinate creates a hidden distributed
> transaction. Model the partial states and recovery path first.

> **Warning:** Comparing only successful reads misses timeout, denied,
> not-found, and malformed-data differences. Shadow representative
> outcomes without allowing the candidate to mutate state.

> **Warning:** Removing monolith routing after the first good cohort
> leaves no tested reversal for a later incident. Keep an owned
> rollback or roll-forward path through the full observation window.

### Related material

See [domain modeling](#4-domain-discovery-and-modeling),
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
   **Known correct:** Dependency work used 380 ms, within its 500 ms
   budget.
3. **Before:** 420 ms remains on the caller deadline. **Decision:**
   Finish local response work. **After:** The call completes within
   800 ms. **Known correct:** A third attempt cannot consume the
   reserved completion time.

Focused timeout input: reservation call `reserve-O-43` is `pending`;
its deadline expires after the remote service may have committed.
Expected output: state `reconcile`, no claim that the reservation was
rejected, and no immediate unkeyed retry.

1. **Before:** `reserve-O-43.state=pending`, and the remote outcome is
   unknown. **Decision:** Apply outcome `timed_out`. **After:** State
   becomes `reconcile`. **Known correct:** A transport timeout did not
   become a false business rejection or a duplicate effect.

### Why it works

The fact that must stay true is that downstream attempts cannot
outlive or exceed the caller's budget. This deadline invariant holds
because each child deadline is derived from remaining parent time,
attempt count is bounded, and cancellation propagates. Outcome
classification sends ambiguous writes to reconciliation instead of an
unsafe retry.

### Go example

See the remaining-deadline calculator and reservation reducer in
[the distributed examples][distributed-examples] and their
[tests][distributed-tests]. The bounded backoff policy is in
[the production examples][production-examples] with its
[tests][production-tests]. Together they cover reserved completion,
one funded retry, and an ambiguous timeout requiring reconciliation.

### Work and memory

Let `a` be the attempt cap, `b` total transferred bytes, and `p` the
connection-pool limit. One logical operation makes at most `a` calls;
latency includes every attempt and backoff but remains under the
parent deadline. Serialization and network work grow with `b`. The
pool bounds live sockets at `p`; concurrency and queue limits likewise
bound goroutines, retained requests, and downstream offered load.

### When to use it

- The caller cannot continue without an immediate remote answer.
- Combined service availability meets the user-visible requirement.
- End-to-end latency can fund the call and local completion reserve.

### When not to use it

Do not hold a database transaction across the call or retry ambiguous
non-idempotent writes without a durable key.

### Common mistakes and edge cases

> **Warning:** Giving each sequential dependency the full 800 ms
> caller deadline can make the chain run past the caller. Derive each
> child budget from remaining time and reserve local completion work.

> **Warning:** Three attempts at both proxy and service layers can
> create nine downstream calls. Assign retry ownership to one layer
> and cap attempts, delay, and total time.

> **Warning:** Creating a new HTTP client for every request prevents
> useful connection reuse and can exhaust sockets. Reuse a configured
> client with bounded transport pools and timeouts.

> **Warning:** Returning stale inventory as if it were current can
> accept an impossible order. A fallback must have an explicit,
> semantically honest product contract.

> **Warning:** Retrying every 500-class response, reset, and timeout
> can repeat permanent or ambiguous failures. Classify retryability
> and require idempotency before repetition.

### Related material

See [service extraction](
#19-service-boundary-discovery-and-extraction),
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
actions with a real business reversal. This durable sequence of local
actions is a **saga**. A **compensation** is a new, fallible business
action, not a reversal of time. Show users when the workflow is still
pending, compensated, or waiting for repair.

### Step-by-step examples

Input: order saga `O-42` starts at version `0` in
`awaiting_payment`. It receives `payment_accepted` at version `1`, the
same delivery again, then `inventory_reserved` at version `2`.
Expected output: state `completed`, version `2`, with the duplicate
causing no second change.

```text
version: 0                  1                    2
state:   awaiting_payment -> reserving --------> completed
event:                      payment_accepted    inventory_reserved
```

1. **Before:** `state=awaiting_payment` and `version=0`. **Decision:**
   Apply `payment_accepted` at the next version, `1`. **After:**
   `state=reserving`, `version=1`, and `lastEvent=payment_accepted`.
   **Known correct:** Durable workflow state now records the completed
   payment step.
2. **Before:** The saga is at version `1` with the same last event.
   **Decision:** Apply duplicate `payment_accepted` at version `1`.
   **After:** State and version remain unchanged. **Known correct:**
   Repeated delivery did not repeat workflow progress.
3. **Before:** `state=reserving` and `version=1`. **Decision:** Apply
   `inventory_reserved` at the next version, `2`. **After:**
   `state=completed`, `version=2`. **Known correct:** The saga reached
   its terminal success through contiguous durable steps.

Focused ordering input: after payment, `O-42` is `reserving` at
version `1`; `inventory_reserved` arrives as version `3`, then is
redelivered correctly as version `2`. Expected output: version `3` is
rejected without change; version `2` completes the saga.

1. **Before:** `state=reserving` and `version=1`. **Decision:** Try to
   apply version `3`. **After:** Return `ErrOutOfOrder`; state remains
   `reserving` at version `1`. **Known correct:** A missing step was
   not silently skipped.
2. **Before:** The same state now receives version `2`. **Decision:**
   Apply the contiguous event. **After:** State becomes `completed` at
   version `2`. **Known correct:** Later delivery restored progress
   without rewriting history.

### Why it works

The fact that must stay true is that every accepted event advances the
workflow by exactly one version, unless it is the identical duplicate
of the last accepted event. This version invariant prevents missing or
conflicting steps. In production, each external effect also needs a
durable workflow state and idempotency identity. Finite states and
attempt limits end in success, compensation, or explicit repair.

### Go example

See the versioned saga reducer in
[the distributed examples][distributed-examples]. Its
[tests][distributed-tests] cover a valid transition, duplicate no-op,
and out-of-order rejection.

### Work and memory

Let `s` be workflow steps, `a` total delivery attempts, and `h` audit
records. Transition work is `O(a)` because each attempt performs one
bounded state check. The workflow retains constant current state,
while audit storage is `O(h)`; `h` normally grows with `s` and `a`.
Indexes must support ready, timed-out, and repair states.
Completion latency includes every asynchronous wait and backoff.

### When to use it

- One business action spans independently committed service state.
- Partial progress can be shown honestly, compensated, or repaired.
- Messages may be duplicated, delayed, or delivered out of order.

### When not to use it

Do not distribute an invariant that truly requires immediate atomicity
when one transactional owner can satisfy it. A local transaction is
simpler when every required change belongs to that owner.

### Common mistakes and edge cases

> **Warning:** Calling a refund a rollback hides that the original
> charge already happened and the refund can fail. Persist
> compensation as a new state, attempts, and repair outcome.

> **Warning:** Keeping `O-42` progress only in process memory loses
> the current step on restart. Persist state, version, effect
> identities, deadlines, and repair decisions.

> **Warning:** Accepting version `3` while version `2` is missing can
> skip a required effect. Reject or buffer gaps under a bounded,
> observable reconciliation policy.

> **Warning:** Showing an order as complete while refund is pending
> gives the user a false terminal result. Expose pending, compensated,
> failed, and repair states honestly.

> **Warning:** Event choreography in which billing waits for inventory
> while inventory waits for billing has no owner able to resolve the
> cycle. Use explicit orchestration when interactions obscure
> progress.

### Related material

See [service extraction](
#19-service-boundary-discovery-and-extraction),
[reliable events](#17-reliable-event-publication-and-consumption), and
[idempotency](#13-idempotency-and-ambiguous-write-outcomes).

## 22. Contract evolution and independent delivery

### The problem it solves

Services deployed at different times must communicate while producers
and consumers temporarily run different contract versions.

### How it works

Change a doorway while the shop stays open: add the wider path first,
let old and new customers pass, and test the expectations of each
consumer. Measure use of the old path and remove it only after every
known consumer has moved. This is additive contract evolution. A
field whose value comes from a fixed named set is an **enum**. Reserve
an explicit unknown behavior for values a newer producer may add.

### Step-by-step examples

Input: the decoder receives old response
`{"id":"O-42","status":"placed"}`, then newer response
`{"id":"O-43","status":"future","note":"new"}`. Expected output:
the first becomes `{ID:O-42, Status:placed, Note:""}`; the second
becomes `{ID:O-43, Status:unknown, Note:"new"}`.

1. **Before:** The old response omits optional `note`. **Decision:**
   Decode known status `placed` and let the missing note use its empty
   default. **After:** The result is the exact old-response value
   above. **Known correct:** A new optional field is not required from
   an old producer.
2. **Before:** The newer response contains status `future` and note
   `new`. **Decision:** Preserve the note and map the unrecognized
   enum value to `unknown`. **After:** The result is the exact
   newer-response value above. **Known correct:** New producer data
   did not make the old status model claim a false known state.

Focused deprecation input: telemetry shows 12 uses per minute of an
old field, then zero uses for the full 30-day compatibility window.
Expected output: keep the field at 12 uses; remove it only after zero
use, owner approval, and successful old/new compatibility tests.

1. **Before:** The old field still has 12 uses per minute.
   **Decision:** Keep producing and reading it. **After:** No contract
   shape changes. **Known correct:** Deployment age did not override
   observed consumer use.
2. **Before:** Use has remained zero for 30 days and compatibility
   tests pass. **Decision:** Obtain owner approval and remove the old
   field in the planned release. **After:** Only the supported shape
   remains. **Known correct:** Evidence and policy defined the removal
   point.

### Why it works

The fact that must stay true is that every supported producer-consumer
version pair shares one valid interpretation. This compatibility
invariant holds when additive changes preserve the previous subset and
new consumers define defaults for absent fields. Explicit unknown
handling prevents a new enum value from becoming a false known value.
Compatibility tests and usage evidence define safe removal.

### Go example

See the compatible response decoder in
[the distributed examples][distributed-examples]. Its
[tests][distributed-tests] accept an absent optional field and map an
unknown enum to an explicit fallback.

### Work and memory

Let `p` be supported producer versions and `c` supported consumer
versions. A complete compatibility matrix has `p * c` pairs, so test
work grows with that product. The decoder visits response bytes and
retains the decoded result. Keep the version window bounded with
explicit migration deadlines. Dual read or write paths add temporary
CPU, network, storage, and reconciliation work.

### When to use it

- Public APIs or events have independently deployed consumers.
- Persisted payloads outlive the code version that wrote them.
- A field, default, enum value, error, or timing rule must evolve.

### When not to use it

Do not create a new version for every additive compatible field, or
retain every historic behavior without an owned support policy. Use a
new version when meaning changes incompatibly.

### Common mistakes and edge cases

> **Warning:** Changing `amount` from cents to dollars while keeping
> the same name and type makes old consumers misprice data. Add a new
> field or version with an explicit migration.

> **Warning:** Schema-only tests can pass while a new default, error,
> ordering, or retry meaning breaks a consumer. Test every behavior
> the consumer contract depends on.

> **Warning:** Removing a field when the provider deploys can break a
> consumer that releases next week. Measure use through the full
> compatibility window before removal.

> **Warning:** Dual-writing old and new fields without comparison can
> let their values diverge silently. Compare, reconcile, and alert
> until one representation is authoritative.

> **Warning:** An undocumented caller can keep an accidental behavior
> alive forever. Publish ownership and support policy, discover usage,
> and give consumers a migration path and deadline.

### Related material

See [service extraction][extraction-reference],
[reliable events](#17-reliable-event-publication-and-consumption), and
[delivery readiness](#12-delivery-deployment-and-recovery-readiness).

## 23. Distributed observability and resilience experiments

### The problem it solves

A distributed request can fail across several services, queues, and
retries. Healthy local dashboards do not reveal the causal chain or
prove recovery.

### How it works

Follow one delivery with the same tracking number across every depot.
A **distributed trace** links records for one operation across process
boundaries. **Trace context** carries that link, while a correlation
identifier can join a longer business flow. Measure a user-visible
**steady state**, the acceptable operating range, then run a bounded
failure drill. A **guardrail** is the automated limit that stops it.

### Step-by-step examples

Input: checkout success was 99.95% for the prior 30 minutes. Add
200 milliseconds of payment latency to one canary for at most ten
minutes. Abort if one one-minute window falls below 99.5%. Recovery
requires three consecutive one-minute windows at or above 99.9%.
Expected output: a 99.4% window stops injection; windows 99.91%,
99.94%, and 99.92% then prove recovery.

1. **Before:** Baseline is 99.95%, and only the named canary is in
   scope. **Decision:** Enable the delay. **After:** Linked traces
   mark delayed canary attempts. **Known correct:** The target and
   ten-minute limit match the approved blast radius.
2. **Before:** A one-minute window reports 99.4%. **Decision:** Apply
   the abort guardrail. **After:** Injection is off, so no new attempt
   receives the delay. **Known correct:** 99.4% is below 99.5%.
3. **Before:** Injection is off and no recovery window has qualified.
   **Decision:** Observe 99.91%, 99.94%, and 99.92%. **After:** Mark
   the system recovered and retain the evidence. **Known correct:**
   All three consecutive windows meet the 99.9% recovery rule.

Focused envelope input: event `E9` has metadata
`{"traceparent":"trace-1"}` and payload
`{"order_id":"O-42"}`. Event `E10` has metadata
`{"authorization":"secret"}`. Expected output: build `E9` with its
trace context and reject `E10` with `ErrUnsafeMetadata`.

1. **Before:** `E9` contains one allowlisted metadata name.
   **Decision:** Copy the metadata and payload into the envelope.
   **After:** Its `traceparent` is `trace-1`. **Known correct:** The
   downstream consumer can continue the trace without shared mutable
   input.
2. **Before:** `E10` contains `authorization`. **Decision:** Reject
   the envelope. **After:** No event is returned and the error wraps
   `ErrUnsafeMetadata`. **Known correct:** Credentials did not become
   durable event baggage.

### Why it works

The fact that must stay true is that an experiment cannot exceed its
approved scope or continue after a guardrail breach. This experiment
boundary invariant is preserved by explicit targeting, duration, and
automated abort rules. A separate numerical recovery condition proves
restoration instead of treating the end of injection as recovery.

### Go example

See the allowlisted event envelope in
[the distributed examples][distributed-examples]. Its
[tests][distributed-tests] preserve trace metadata and reject
credential baggage.

### Work and memory

Let `s` be services on the path, `a` recorded attempts, `m` total
metadata bytes, and `b` payload bytes. Envelope construction takes
`O(m + b)` work and memory because it validates and copies both. Trace
analysis is `O(s + a)` for a linear path representation. Sampling,
label cardinality, retention, experiment duration, and concurrent
injections need hard limits.

### When to use it

- A critical operation crosses process or queue boundaries.
- Operators can state an exact user-visible steady-state range.
- A bounded failure hypothesis needs production-like evidence.

### When not to use it

Do not inject production failure before ownership, observability,
abort, and restoration controls exist. Static review and ordinary
tests should find basic missing timeouts before an experiment does.

### Common mistakes and edge cases

> **Warning:** Starting a new trace after queue delivery severs the
> causal chain. Propagate the parent or record a message link.

> **Warning:** Using `order_id` as a metric label creates one series
> per order and can exhaust monitoring capacity. Keep it in traces or
> logs; use bounded metric dimensions.

> **Warning:** Defining steady state as "pods are healthy" can miss a
> broken checkout. Guard the user-visible success and latency promise.

> **Warning:** A manual abort that has never been exercised can fail
> during the drill. Test automatic stop and restoration before fault
> injection begins.

> **Warning:** Ending the delay at 99.4% does not prove recovery.
> Require the stated post-fault windows and retain their evidence.

### Related material

See [service extraction][extraction-reference],
[observability](#11-observability-and-operational-controls), and
[Entry 28](#28-reliability-engineering-incidents-and-cost).

## 24. Data scaling, caching, and tenant isolation

### The problem it solves

Growing data and traffic pressure one database, while caches,
partitions, and shared tenancy can introduce stale reads, hot keys,
leakage, and migrations harder than the original bottleneck.

### How it works

A **cache** is a temporary copy, not the authoritative source. Its
**freshness contract** states how old a value may be and what events
replace or invalidate it. A time to live (TTL) is only the maximum
retention time; it does not synchronize a write immediately. Split
data by an ownership key, and include the **tenant**, the customer or
organization that owns the data, in every lookup and derived key.

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

Focused key input: call `NewTenantCacheKey` with
`("tenant-a", "product-1", 2)`, then with
`("a:v1", "b", 2)` and `("a", "v2:b", 1)`. Expected output: the
first key is `8:tenant-a:v2:9:product-1`; the delimiter cases produce
different keys, `4:a:v1:v2:1:b` and `1:a:v1:4:v2:b`.

1. **Before:** Tenant `tenant-a`, item `product-1`, and version `2`
   are valid. **Decision:** Prefix each identity with its length.
   **After:** Return `8:tenant-a:v2:9:product-1`. **Known correct:**
   The version and both ownership components are encoded.
2. **Before:** The two remaining inputs contain delimiters in
   different components. **Decision:** Apply the same length-prefixed
   encoding. **After:** Return the two distinct expected keys.
   **Known correct:** Delimiters cannot move a boundary between tenant
   and item.

### Why it works

The fact that must stay true is that every derived location retains
the authoritative tenant and object identity. This isolation invariant
is preserved by compound keys and matching storage policy.
Length-prefixed components keep embedded delimiters from creating the
same serialized key. A bounded refill coordinator can let waiters
share one load without allowing an unbounded queue.

### Go example

See the typed tenant cache key in
[the platform examples][platform-examples]. Its
[tests][platform-tests] reject empty tenant identity and prove that
delimiter-containing identities cannot alias.

### Work and memory

Let `t` be tenant-identifier bytes, `i` item-identifier bytes, `k`
cached entries, and `n` items moved during resharding. Key
construction takes `O(t + i)` work and output memory. Hash-cache
lookup is `O(1)` on average and `O(k)` in the collision worst case;
configured entry and byte limits bound retained memory. Resharding
takes `O(n)` item operations plus the bytes copied, and consumes
source, destination, and network capacity.

### When to use it

- Measurements identify read load or data volume as the bottleneck.
- Repeated reads can tolerate a written freshness contract.
- Data must be partitioned while preserving tenant ownership.

### When not to use it

Do not cache correctness-critical mutable state without a freshness
contract. Do not shard before a single well-operated database is the
measured limiting resource.

### Common mistakes and edge cases

> **Warning:** Keying both tenants' `P1` as only `P1` lets one tenant
> read the other's value. Include tenant identity in cache, queue,
> deduplication, storage, and invalidation keys.

> **Warning:** A five-minute TTL can serve an old price for nearly
> five minutes after a write. Use explicit invalidation or versioned
> keys when the freshness contract is tighter.

> **Warning:** Letting 10,000 misses refill the same key can overload
> the database; this synchronized surge is a cache stampede. Coalesce
> refills with bounded waiters and cap negative-cache retention.

> **Warning:** Partitioning writes by current hour sends all current
> traffic to one partition. Choose and test a key with sufficient
> write distribution and an owned resharding plan.

> **Warning:** Reading a replica immediately after a write can return
> the old value. Route read-after-write traffic according to an
> explicit consistency policy.

### Related material

See [cloud operations](#26-cloud-infrastructure-and-iam),
[advanced PostgreSQL](#15-advanced-postgresql-operation-and-recovery),
and [Entry 25](#25-security-privacy-and-software-supply-chain).

## 25. Security, privacy, and software supply chain

### The problem it solves

Production risk includes vulnerable dependencies, compromised builds,
excess privilege, retained personal data, unsafe logs, and secrets
that outlive their intended scope.

### How it works

Protect a system like a building: a **threat model** lists assets,
entry points, likely misuse, and controls. **Least privilege** gives
each identity only the actions and resources its work requires. This
is part of identity and access management (IAM). Collect only needed
data, and verify build inputs and signed artifacts through the
**software supply chain**. Make deletion, key rotation, and incident
response procedures executable and tested.

### Step-by-step examples

Input: construct rotation state with current version `v4` and accepted
versions `v3` and `v4`. After the 24-hour overlap, no valid `v3`
artifact remains. Expected output: both versions verify during
overlap; after `Retire("v3")`, `v3` fails and `v4` still verifies.

1. **Before:** Current version is `v4`, and both versions are
   accepted. **Decision:** Check `Accepts("v3")` and `Accepts("v4")`
   during the overlap. **After:** Both return true. **Known correct:**
   New signing can use `v4` without invalidating an unexpired `v3`
   artifact.
2. **Before:** The overlap has elapsed and no valid `v3` artifact
   remains. **Decision:** Call `Retire("v3")`. **After:** `v3` returns
   false while `v4` returns true. **Known correct:** Retirement
   removed only the old verification authority.

Focused authorization input: allow only tuple
`(order-worker, queue/orders, receive)`. Expected output: that tuple
is allowed; `(order-worker, bucket/private, read)` is denied.

1. **Before:** The authorizer contains one permission. **Decision:**
   Check the exact allowed tuple. **After:** Return true.
   **Known correct:** The declared work can proceed.
2. **Before:** No private-bucket permission exists. **Decision:**
   Check the excess tuple. **After:** Return false. **Known correct:**
   Missing privilege fails closed instead of expanding the workload's
   reach.

### Why it works

The fact that must stay true is that every accepted artifact,
credential, or action chains to a currently trusted, auditable
authority. This authority invariant is preserved by explicit
permission tuples and accepted key sets. Overlap preserves
availability during rotation, while retirement bounds exposure. Data
minimization reduces impact when another control fails.

### Go example

See least-privilege authorization and overlap rotation in
[the platform examples][platform-examples]. Their
[tests][platform-tests] cover one allowed tuple, denied excess access,
accepted overlap keys, and retirement.

### Work and memory

Let `p` be stored permission tuples, `k` accepted key versions, `d`
dependencies scanned, and `b` artifact bytes. Permission and key-set
lookups are `O(1)` on average and `O(p)` or `O(k)` in a hash-collision
worst case; retained map memory is `O(p + k)`. Scanning grows with `d`
and artifact verification with `b` plus the chosen algorithm's cost.
Audit and retention copies consume storage; deletion must cover each
copy under its documented lifecycle.

### When to use it

- A workload can reach production data or infrastructure.
- A build consumes third-party code or produces deployable artifacts.
- Personal data has collection, retention, export, or deletion duties.

### When not to use it

Security controls are not optional. Scale their depth to sensitivity,
privilege, exposure, and recovery cost; do not collect speculative
data or grant permanent broad credentials to postpone policy design.

### Common mistakes and edge cases

> **Warning:** A scanner finding a critical dependency with no owner
> or deadline produces inventory, not risk reduction. Assign severity,
> owner, due date, exception, and verification evidence.

> **Warning:** Signing before an untrusted packaging step lets that
> step alter the artifact after attestation. Sign the immutable output
> that deployment will verify.

> **Warning:** Encrypting a customer export does not authorize every
> employee to read it. Enforce access policy and data minimization in
> addition to encryption.

> **Warning:** Deleting only the primary row leaves personal data in
> caches, exports, indexes, replicas, or backups. Track each copy and
> its deletion or expiry evidence.

> **Warning:** Replacing `v3` with `v4` in one instant rejects valid
> in-flight artifacts. Publish trust, switch issuance, observe the
> bounded overlap, then retire the old version with rollback ready.

### Related material

See [scale and cloud operations][scale-reference],
[identity and security](#9-identity-authorization-and-api-security),
and [cloud IAM](#26-cloud-infrastructure-and-iam).

## 26. Cloud infrastructure and IAM

### The problem it solves

Cloud application programming interfaces (APIs) make infrastructure
easy to create and equally easy to drift, overprivilege, expose
publicly, or recreate differently during recovery.

### How it works

Keep infrastructure as reviewed blueprints instead of hand-built
rooms. **Infrastructure as code** stores those declarations in a
repeatable, reviewed form. Separate environments and **failure
domains**, areas that can fail independently. Identity and access
management (IAM) gives each **workload identity**, the authenticated
identity of running software, short-lived credentials and permission
for only its job.

### Step-by-step examples

Input: workload identity `order-worker` requests `receive` on
`queue/orders` and `read` on `bucket/private` using a short-lived
credential. Expected output: queue access succeeds and bucket access
is denied.

1. **Before:** The workload has no embedded cloud key. **Decision:**
   Exchange its deployment identity for a short-lived credential.
   **After:** The credential names `order-worker` and its expiry.
   **Known correct:** No static secret is stored.
2. **Before:** Policy allows `receive` on `queue/orders`.
   **Decision:** Authorize the queue request against workload, action,
   and resource.
   **After:** One message is returned. **Known correct:** The action
   matches the declared need.
3. **Before:** The same identity requests `read` on `bucket/private`.
   **Decision:** Find no matching policy tuple. **After:** Access is
   denied and audited. **Known correct:** Successful queue access did
   not imply broad cloud permission.

Focused diagnostic input: classify `ErrNetwork`, `ErrAuthorization`,
and `errors.New("invalid order")`. Expected output: `network`,
`authorization`, and `application`, in that order.

1. **Before:** The first error wraps the network sentinel.
   **Decision:** Test the network boundary first. **After:** Return
   `network`. **Known correct:** Transport failure is not blamed on
   application logic.
2. **Before:** The second error wraps the authorization sentinel.
   **Decision:** Test the authorization boundary. **After:** Return
   `authorization`. **Known correct:** Denied policy is distinct from
   authentication and connectivity.
3. **Before:** `invalid order` matches no boundary sentinel.
   **Decision:** Use the default category. **After:** Return
   `application`. **Known correct:** An expected domain failure is not
   reported as cloud infrastructure failure.

### Why it works

The facts that must stay true are that deployed resources can be
reproduced from reviewed declarations and every allowed action matches
an explicit workload need. These infrastructure and authority
invariants make drift and excess privilege detectable. Short
credential lifetime and narrow scope bound misuse after compromise.

### Go example

See the boundary-aware diagnostic classifier in
[the platform examples][platform-examples]. Its
[tests][platform-tests] preserve network, authorization, and
application failure categories.

### Work and memory

Let `r` be resources in a plan, `e` dependency edges, `c` provider
calls, and `a` audited actions. Planning traverses the resource graph
in `O(r + e)` work before provider-specific calls, while apply work
also depends on `c`. Retained plan state is `O(r + e)`; audit storage
grows with `a`. Multi-zone or regional redundancy adds
running cost, transfer, and recovery coordination.

### When to use it

- A non-local environment must be reproducible and reviewable.
- Multiple workloads share an account but require different access.
- Recovery must recreate networks, data services, policy, and traffic.

### When not to use it

Do not put rapidly changing application rollout state into a
low-level infrastructure state file. Never use one administrator
credential for all workloads; split ownership and permissions.

### Common mistakes and edge cases

> **Warning:** Applying an unreviewed plan can expose a database or
> destroy retained state. Review the exact plan and require approval
> for high-impact changes.

> **Warning:** Sharing one writable state file across development and
> production lets a development action mutate production resources.
> Isolate state, credentials, and approval boundaries.

> **Warning:** A private network address does not prove which workload
> called. Authenticate workload identity and authorize the exact
> action and resource.

> **Warning:** Unrestricted egress lets a compromised worker bypass
> declared dependency controls. Restrict destinations and log denied
> attempts with bounded metadata.

> **Warning:** Multi-zone labels do not prove recovery. Restore data
> and fail traffic in a timed exercise, then retain evidence that the
> stated objective was met.

### Related material

See [data scaling][scale-reference],
[delivery readiness](#12-delivery-deployment-and-recovery-readiness),
and [platform operations](#27-orchestration-and-platform-operations).

## 27. Orchestration and platform operations

### The problem it solves

Schedulers can restart and scale processes, but poor resource
requests, health probes, shutdown, disruption policy, or queue signals
turn automation into churn.

### How it works

Run compute like a delivery fleet. An **orchestrator** places,
replaces, and scales workload instances. A resource **request**
reserves scheduling capacity; a **limit** bounds permitted use.
Startup reports initialization. Readiness controls new traffic, and
liveness reports whether the process itself must restart. Scaling and
maintenance must act only on signals the workload can honor.

### Step-by-step examples

Input: worker pod P1 holds job `J7` when it receives termination with
a 30-second grace period. Expected output: P1 admits no new work and
either completes or releases `J7` before exit.

1. **Before:** P1 is ready, owns `J7`, and has 30 seconds remaining.
   **Decision:** Close local admission and mark readiness false before
   draining. **After:** New work is rejected locally while endpoint
   removal propagates. **Known correct:** The draining workload cannot
   grow through normal admission.
2. **Before:** P1 is unready and still owns `J7`. **Decision:** Finish
   its durable effect and acknowledge it before the grace deadline.
   **After:** `J7` is complete and P1 may exit. **Known correct:** The
   acknowledgement follows durable completion.
3. **Before:** Replacement pod P2 becomes ready. **Decision:** Count
   it toward rollout capacity only after its startup checks pass.
   **After:** The scheduler restores the desired replica count.
   **Known correct:** Replacement and drain preserve minimum capacity.

Focused failure input: five seconds remain, and `J7` has not completed
its durable effect. Expected output: P1 does not acknowledge `J7`; it
releases ownership or exits and lets the lease expire.

1. **Before:** P1 cannot prove durable completion before termination.
   **Decision:** Withhold acknowledgement and release ownership when
   supported. **After:** `J7` becomes eligible for another worker now
   or after lease expiry. **Known correct:** Unfinished work was not
   falsely recorded as complete.

Focused health input: start with lifecycle `{Initialized:true,
DependencyReady:true, Draining:false}`; then lose the dependency; then
set `Draining:true`. Expected output: health begins
`{Startup:true, Readiness:true, Liveness:true}`; dependency loss makes
only readiness false; draining keeps readiness false.

1. **Before:** Initialization and the dependency are ready, and drain
   is false. **Decision:** Evaluate all three health dimensions.
   **After:** Startup, readiness, and liveness are true.
   **Known correct:** The instance may receive work.
2. **Before:** The dependency becomes unavailable. **Decision:**
   Re-evaluate without treating dependency loss as a dead event loop.
   **After:** Startup and liveness remain true; readiness is false.
   **Known correct:** Traffic stops without creating a restart loop.
3. **Before:** The lifecycle enters drain. **Decision:** Evaluate
   readiness with `Draining:true`. **After:** Readiness remains false.
   **Known correct:** A terminating instance cannot re-enter service.

### Why it works

The facts that must stay true are that only ready instances admit new
work and every owned unit is completed or made eligible before process
death. These admission and ownership invariants survive restart and
deployment through separate probes, local admission closure, durable
leases, and bounded drain.

### Go example

See the startup, readiness, and liveness state object in
[the platform examples][platform-examples]. Its
[tests][platform-tests] cover healthy service, dependency loss, and
draining without creating a restart loop.

### Work and memory

Let `r` be replicas, `q` queued work, and `d` the downstream
concurrent capacity. Health evaluation uses constant work and state
per instance. Requested cluster capacity grows with `r`; replica state
is `O(r)`.
Autoscaling lag requires measured headroom. A queue-based target must
keep admitted concurrency at or below `d`, even when `q` is larger.

### When to use it

- Multiple instances need automated placement and replacement.
- Rollouts must preserve capacity while old instances drain.
- Queue or traffic signals can scale within downstream capacity.

### When not to use it

Do not add a cluster when a simpler supervised deployment meets the
reliability need. Do not use restarts to mask deterministic
application defects or missing backpressure.

### Common mistakes and edge cases

> **Warning:** Using dependency health for liveness restarts every pod
> during a database outage and increases load. Remove traffic with
> readiness while a responsive process stays live.

> **Warning:** Marking ready before listeners and required local state
> are usable routes requests into startup failures. Publish readiness
> only after the instance can honor admission.

> **Warning:** Setting a memory limit below measured bursts turns a
> brief peak into repeated kills. Measure steady and burst use, then
> test requests, limits, and degradation together.

> **Warning:** Scaling 100 workers against a database capacity of 40
> replaces a queue with database contention. Cap admitted concurrency
> at the tightest downstream boundary.

> **Warning:** A worker that acknowledges `J7` during forced shutdown
> can lose unfinished work. Drain within grace, and withhold
> acknowledgement until the durable effect is proven.

### Related material

See [scale and cloud operations][scale-reference],
[HTTP lifecycle](#6-http-contracts-and-process-lifecycle), and
[durable jobs](#16-durable-background-jobs-and-scheduling).

## 28. Reliability engineering, incidents, and cost

### The problem it solves

Without explicit objectives and ownership, teams optimize noisy
metrics, repeat incidents, and spend unpredictably without knowing
which cost buys a user-visible reliability improvement.

### How it works

Run a service like a train line. A **service-level indicator (SLI)**
measures what riders experience. A **service-level objective (SLO)**
sets the target, and its allowed shortfall is an **error budget**.
A **burn rate** describes how quickly failures consume that budget.
Alerts, runbooks, incident reviews, and cost should connect to the
same demand and service promise.

### Step-by-step examples

Input: checkout consumes 40% of its monthly error budget in one hour
during a rollout. Its recovery rule requires three consecutive
five-minute windows with success at least 99.9% and 95th-percentile
latency at most 300 milliseconds. Expected output: the fast-burn alert
pages, rollout stops, traffic reverts, and all three windows pass.

1. **Before:** The release is active and the one-hour burn reaches
   40%. **Decision:** Page the on-call owner using the fast-burn rule.
   **After:** Open one incident and assign a commander.
   **Known correct:** The page represents user-visible budget use.
2. **Before:** The release is the leading correlated change.
   **Decision:** Stop rollout and route traffic to the last healthy
   version. **After:** New exposure ends. **Known correct:** The
   rollback follows a documented mitigation.
3. **Before:** Traffic has reverted. **Decision:** Observe three
   qualifying five-minute windows. **After:** Record recovery, the
   timeline, and the exact window evidence. **Known correct:** Both
   numerical exit conditions passed before the incident recovered.

Focused budget input: call `CalculateErrorBudget(1_000_000, 999_500,
0.999)`. Expected output: 1,000 allowed misses, 500 consumed misses,
and 500 remaining misses.

1. **Before:** There are 1,000,000 eligible events and a 99.9% target.
   **Decision:** Round `events * (1 - target)`. **After:** The budget
   allows 1,000 misses. **Known correct:** The allowed shortfall is
   0.1% of the eligible events.
2. **Before:** 999,500 events are good. **Decision:** Subtract good
   from eligible. **After:** 500 misses are consumed.
   **Known correct:** One million minus 999,500 is 500.
3. **Before:** 1,000 misses are allowed and 500 are consumed.
   **Decision:** Subtract consumed from allowed. **After:** 500
   remain. **Known correct:** Allowed, consumed, and remaining
   reconcile.

Focused failure input: call the same function with zero eligible
events. Expected output: `ErrNoEligibleEvents` and no budget.

1. **Before:** The denominator is zero. **Decision:** Reject before
   division or target math. **After:** Return `ErrNoEligibleEvents`.
   **Known correct:** No percentage was invented from an empty sample.

### Why it works

The fact that must stay true is that reliability decisions use a
user-visible objective and the remaining error budget. This decision
invariant ties short-window burn evidence to the long target. Owned
mitigation and exact recovery checks turn incident response into a
verifiable state change rather than a declaration.

### Go example

See the integer event-budget calculation in
[the platform examples][platform-examples]. Its
[tests][platform-tests] cover a 99.9 percent objective and reject a
zero-event denominator.

### Work and memory

Let `e` be eligible events, `s` retained time-series samples, and `l`
active label combinations. The budget calculation uses `O(1)` work and
memory after counts are aggregated. Telemetry storage is `O(s)`, and
aggregation work grows with samples and `l`. Redundancy and headroom
add direct cost; unit economics must allocate shared infrastructure,
third-party, incident, and peak-capacity costs.

### When to use it

- A production capability has a user-visible success or latency goal.
- An on-call owner needs actionable page and recovery thresholds.
- Reliability and cost tradeoffs need a shared numerical basis.

### When not to use it

Do not create an objective no team can observe or act on. Do not
reduce reliability to one availability percentage when latency,
correctness, durability, or freshness also defines user success.

### Common mistakes and edge cases

> **Warning:** Paging on central processing unit (CPU) use alone wakes
> an operator without proving user harm or imminent exhaustion. Page
> on a user symptom or a tested leading signal with an owned action.

> **Warning:** Setting every target to 100% creates no error budget
> and cannot distinguish ordinary risk from emergency risk. Choose a
> target grounded in user need and achievable system behavior.

> **Warning:** A review action reading "improve monitoring" has no
> verifiable completion. Assign an owner, deadline, expected system
> change, and test evidence.

> **Warning:** Counting compute while omitting third-party, network
> egress, telemetry, standby, and incident costs understates the
> service. Allocate every material cost to a declared unit.

> **Warning:** Optimizing average unit cost can remove the headroom
> needed at peak demand. Model peak capacity and recovery reserve
> before accepting the apparent saving.

### Related material

See [scale and cloud operations][scale-reference],
[observability](#11-observability-and-operational-controls), and
[delivery readiness](#12-delivery-deployment-and-recovery-readiness).

## 29. Production integration boundaries

### The problem it solves

Webhooks, files, email, search, analytics, and partner application
programming interfaces (APIs) have different trust, delivery, size,
privacy, and failure semantics from internal calls.

### How it works

Treat each outside system like a delivery company with its own forms
and failures. A **webhook** is an incoming request sent when an
outside event occurs. A narrow **adapter** translates the contract and
verifies origin. It bounds input, durably accepts work, and prevents
repeated effects. **Reconciliation** later compares derived results
with the authoritative source to find missed or divergent work.

### Step-by-step examples

Input: a receiver allows 32 bytes and a five-minute replay window.
At Unix time 1,000, correctly signed event `E-91` arrives twice with
event time 1,000 and raw body `{"ok":true}`. Expected output: the
first call returns `(true, nil)`, the duplicate returns
`(false, nil)`, and pending work is exactly one.

1. **Before:** `E-91` is unseen and only raw request bytes are trusted
   for signature input. **Decision:** Check timestamp and signature
   before parsing. **After:** The authentic payload may be decoded.
   **Known correct:** Parsing did not change the bytes being verified.
2. **Before:** No accepted identity exists for `E-91`. **Decision:**
   Claim its unique ID and add one work item under the same lock.
   **After:** Return `(true, nil)` with pending work equal to one.
   **Known correct:** Acknowledgement follows durable ownership.
3. **Before:** The second delivery has the same valid ID and payload.
   **Decision:** Find the existing claim instead of adding work again.
   **After:** Return `(false, nil)` with one work item still present.
   **Known correct:** The result distinguishes duplicate transport
   while preserving one semantic effect.

Focused failure input: event `E-92` has the same time and body but
signature `bad`; the verifier returns `bad signature`. Expected
output: acceptance is false, the returned error wraps the verifier
error, and pending work remains zero.

1. **Before:** No identity has been claimed for `E-92`.
   **Decision:** Verify the unchanged body and receive
   `bad signature`. **After:** Return false with
   `verify webhook: bad signature` before locking or recording the
   identity. **Known correct:** Unauthenticated input did not become
   durable work.

### Why it works

The facts that must stay true are that only authenticated, bounded,
schema-valid input becomes durable work and each external identity has
at most one semantic effect. These acceptance and identity invariants
are established by raw-byte verification followed by an atomic claim.
Reconciliation covers loss and divergence outside the normal path.

### Go example

See the bounded webhook receiver in
[the delivery examples][delivery-examples]. Its
[tests][delivery-tests] cover valid acceptance, signature failure, and
duplicate acknowledgement without duplicate work.
The in-memory lock models only the atomic decision. Production storage
must commit the event identity and queued work together durably.

### Work and memory

Let `b` be bounded body bytes and `e` retained event identities.
Signature checking and body copying take `O(b)` work. Hash-based
identity lookup is `O(1)` on average and `O(e)` in the collision worst
case; retained identity memory is `O(e)`. Bulk import memory must be
bounded by batch bytes. Search and analytics copies add indexing,
transfer, retention, and deletion work.

### When to use it

- Data crosses an external trust or ownership boundary.
- Delivery can repeat, arrive late, or acknowledge before processing.
- Search, email, analytics, or files have different truth semantics.

### When not to use it

Do not make search or analytics the transactional source of truth.
Do not hold a local database transaction open across slow partner
work; durably stage the external action instead.

### Common mistakes and edge cases

> **Warning:** Parsing JavaScript Object Notation (JSON) before
> signature verification can normalize the signed bytes. Verify the
> exact bounded request body first.

> **Warning:** Returning success before claiming `E-91` and its work
> lets a crash lose the event permanently. Acknowledge only after
> durable acceptance.

> **Warning:** Loading a 20-gigabyte partner file into memory permits
> resource exhaustion. Stream into bounded batches with row and byte
> limits, validation, and a quarantine path.

> **Warning:** Treating email provider acceptance as user delivery,
> or analytics as exact accounting, assigns authority to a derived
> system. Record each system's actual guarantee.

> **Warning:** A normally reliable event path can still lose or
> diverge during outages and repairs. Reconcile on a bounded schedule
> and make discrepancies owned and observable.

### Related material

See [idempotency](#13-idempotency-and-ambiguous-write-outcomes),
[identity and security](#9-identity-authorization-and-api-security),
and [reliable events](#17-reliable-event-publication-and-consumption).

## 30. Maintenance and legacy-system evolution

### The problem it solves

An old system often lacks tests, clear boundaries, current
dependencies, or a safe migration path, yet it continues to carry
valuable behavior and data.

### How it works

Renovate an occupied house one room at a time. A **characterization
test** records existing observable behavior without claiming it is
ideal. A **seam** is a boundary where behavior can be redirected. Put
a **facade**, one stable interface over old and new paths, at that
seam. Compare shadow results, stage traffic, and keep a tested
reversal until the defined evidence window passes.

### Step-by-step examples

Input: legacy pricing returns `$12` for product `P1`; the first
candidate returns `$13`. Policy requires zero unexplained mismatches
across 10,000 shadow requests before a 5% cohort may move. Expected
output: comparison blocks the first cutover; after correction and
10,000 matches, exactly 5% is eligible for candidate routing.

1. **Before:** Only legacy pricing serves users. **Decision:** Record
   `$12` and other boundary outcomes as characterization cases.
   **After:** The current observable contract is executable.
   **Known correct:** Replacement starts from observed behavior.
2. **Before:** The facade returns legacy `$12` and shadows candidate
   `$13`. **Decision:** Record and classify the mismatch without
   routing users to the candidate. **After:** Cutover remains at zero.
   **Known correct:** Comparison cannot silently change production.
3. **Before:** The corrected candidate has 10,000 consecutive matches.
   **Decision:** Route the policy's 5% cohort through the facade.
   **After:** Candidate exposure is 5%, with reversal still enabled.
   **Known correct:** The observed evidence and exposure equal the
   stated gate.

Focused repair input: record `{tenant-a, order-1, v4, old}` receives a
dry-run request for the same tenant, ID, and version with value `new`.
Expected output: checkpoint `tenant-a/order-1/v4`; execution succeeds
but the record remains `{tenant-a, order-1, v4, old}`.

1. **Before:** Scope and expected version match the record.
   **Decision:** Build the repair plan with `DryRun:true`. **After:**
   The checkpoint is `tenant-a/order-1/v4`. **Known correct:** The
   checkpoint names the exact authority and version being assessed.
2. **Before:** The dry-run plan is valid. **Decision:** Execute it.
   **After:** Return nil without changing value or version.
   **Known correct:** Planning evidence was produced without mutation.

Focused failure input: the same record receives a request for
`tenant-b`. Expected output: `ErrRepairScope` and no plan.

1. **Before:** Record tenant is `tenant-a`; request tenant is
   `tenant-b`. **Decision:** Reject during plan validation. **After:**
   Return `ErrRepairScope`. **Known correct:** A maintenance tool did
   not cross the declared tenant boundary.

### Why it works

The facts that must stay true are that each migrated request has one
authoritative result and each staged change has a tested reversal
until confidence is earned. These authority and reversibility
invariants are supported by characterization tests, facades,
comparison gates, and bounded cohorts.

### Go example

See the tenant-scoped, version-guarded repair plan in
[the delivery examples][delivery-examples]. Its
[tests][delivery-tests] cover a dry run and reject an out-of-scope
tenant.

### Work and memory

Let `n` be records in a backfill and `q` requests shadowed. Backfill
and full comparison take `O(n)` work. Shadowing adds `O(q)` candidate
executions and comparison records while enabled. Batch bytes,
concurrency, and rate need hard bounds. Each temporary adapter, flag,
and dual path needs an owner, deadline, and removal condition.

### When to use it

- Existing behavior or data cannot be reconstructed confidently.
- Service must remain available while one path moves at a time.
- A facade can isolate change and provide a measured reversal.

### When not to use it

Do not preserve unsafe behavior that violates a required security or
integrity rule. Isolate it, record the compatibility break, and
correct it through an explicit migration decision.

### Common mistakes and edge cases

> **Warning:** Rewriting from undocumented intent can remove behavior
> customers depend on. Characterize actual boundaries before deciding
> which observed behavior to preserve or deliberately break.

> **Warning:** A broad mock that asserts internal call order freezes
> the legacy implementation instead of its outcome. Pin boundary
> inputs, outputs, errors, and durable effects.

> **Warning:** Dual-writing without comparison, ordering, and repair
> can leave two authorities after one partial failure. Name one source
> of truth and reconcile every mismatch.

> **Warning:** A facade or feature flag with no removal owner becomes
> a permanent branch. Record its owner, expiry, evidence gate, and
> delete path when it is introduced.

> **Warning:** Combining dependency, schema, and behavior changes in
> one release makes failure attribution and reversal ambiguous. Split
> them into independently verifiable steps.

### Related material

See [integrations and maintenance][integration-reference],
[service extraction](#19-service-boundary-discovery-and-extraction),
and [Entry 22](#22-contract-evolution-and-independent-delivery).

## 31. Capstone delivery and engineering leadership

### The problem it solves

Senior backend work must turn uncertain product and operational needs
into a reviewable plan, delivered system, measurable launch, and
maintained learning rather than only produce locally correct code.

### How it works

Lead a large change like a construction project. Separate
**evidence**, observed or tested facts, from **assumptions** that
still need proof. State the goal, constraints, invariants, and
non-goals. Compare real
options; divide work into reversible **milestones**; name owners; and
define numerical launch and rollback guardrails before exposure.

### Step-by-step examples

Input: checkout timeout rate is 2% during 200-request-per-second
bursts. Database utilization is 45%, query latency is stable, workers
saturate, and the partner sustains 120 requests per second. Options
are query optimization, double workers, or a 600-item queue admitting
120 per second. Expected output: choose the queue, canary at 5%, and
revert if either guardrail fails: timeout exceeds 1%, or
95th-percentile queue wait exceeds 500 milliseconds in a five-minute
window.

1. **Before:** The symptom is measured but its cause is unconfirmed.
   **Decision:** Profile queries, saturation, and queue wait before
   selecting an option. **After:** Evidence shows worker saturation
   and a 120-per-second partner limit. **Known correct:** Stable
   queries and 45% database use do not support the query hypothesis.
2. **Before:** Three options and their costs are visible.
   **Decision:** Choose the 600-item queue with 120-per-second
   admission; record why query work and more workers miss the limiting
   boundary. **After:** Capacity, latency, overflow, and rollback have
   owners. **Known correct:** The choice addresses measured burst
   admission without exceeding the partner.
3. **Before:** The queue receives a 5% canary. **Decision:** Expand
   only while both five-minute guardrails hold. **After:** The rollout
   advances or reverts to direct admission on the first breach.
   **Known correct:** Delivery state follows criteria defined before
   launch.

Focused evidence input: validate contract `cancel is idempotent`, test
`cancel_duplicate_test`, signal `cancel_outcomes_total`, runbook
`cancel-recovery`, status `verified`, owner `orders-team`, and risk
`inventory unavailable`. Expected output: validation succeeds; after
clearing the runbook, it returns `runbook evidence is required`.

1. **Before:** All seven evidence fields are non-empty. **Decision:**
   Validate every required field. **After:** Return nil.
   **Known correct:** Contract, verification, operation, ownership,
   and risk are all reviewable.
2. **Before:** The runbook is now empty. **Decision:** Validate in the
   same order. **After:** Return `runbook evidence is required`.
   **Known correct:** A verified claim cannot launch without its
   recovery procedure.

### Why it works

The facts that must stay true are that every material claim maps to
evidence, every risk maps to an owner, and every rollout step maps to
acceptance and reversal criteria. These evidence and accountability
invariants make technical judgment reviewable before cost is sunk and
correctable afterward.

### Go example

See the typed requirement-evidence validator in
[the delivery examples][delivery-examples]. Its
[tests][delivery-tests] accept a complete evidence chain and reject a
missing runbook.

### Work and memory

Let `o` be real options, `r` material risks, and `m` rollout
milestones. The minimum review inventory grows with `o + r + m`, with
evidence proportional to blast radius and irreversibility.
Small reversible changes need little ceremony; durable data, public
contracts, privacy, or cross-team operations justify deeper review.
Bound coordination with decision owners and deadlines.

### When to use it

- Work changes a public, durable, privacy, or ownership boundary.
- A migration can cause costly or difficult-to-reverse effects.
- Launch and operation span several accountable owners.

### When not to use it

Do not create a large design process for a local reversible change.
Do not mistake consensus for accountable decision ownership; name the
decision owner and record dissenting evidence.

### Common mistakes and edge cases

> **Warning:** Presenting only the preferred queue and calling the
> document a tradeoff analysis hides whether alternatives were tested.
> Compare at least the plausible options against the same constraints.

> **Warning:** Leaving non-goals, migration, rollback, or operational
> ownership implicit creates different plans in each reviewer's head.
> State them and name the accountable owner.

> **Warning:** Completing every project task does not prove checkout
> works end to end. Make each milestone demonstrate a user-visible or
> operational capability.

> **Warning:** Launching on schedule after the timeout guardrail
> breaches replaces evidence with calendar pressure. Stop or revert
> according to the predeclared rule.

> **Warning:** Recording a lesson without changing code, a test, a
> runbook, or ownership cannot prevent recurrence. Close learning with
> a verified system or process change.

### Related material

See [integrations and maintenance][integration-reference],
[architecture](#5-architecture-and-service-construction), and
[Entry 28](#28-reliability-engineering-incidents-and-cost).

## 32. Language portability

### The problem it solves

Backend principles survive a language change, but ownership,
cancellation, errors, concurrency, runtime scheduling, and tooling
appear through different syntax and defaults.

### How it works

Moving code to another language is like moving a recipe to a different
kitchen. Preserve the **observable contract**, the results, errors,
effects, timing, and resource lifetime a caller can detect. A
**semantic gap** is a target-language default that differs from that
contract. Map Go values, errors, contexts, and goroutines to owned
target mechanisms instead of translating syntax word for word.

### Step-by-step examples

Input: conformance case one returns value `placed` and no error. Case
two runs a waiting operation with an already canceled context or
target cancellation token. Expected output in both languages: case
one is `{Value:"placed", Code:"ok"}`; case two is
`{Value:"", Code:"canceled"}` and the wait has ended.

1. **Before:** The Go operation returns `placed` and nil.
   **Decision:** Map nil error to `OutcomeOK`. **After:** Return the
   exact `placed` and `ok` outcome. **Known correct:** Ordinary
   success preserves both value and status.
2. **Before:** The Go waiting operation receives an already canceled
   context. **Decision:** Let the operation observe cancellation and
   map `context.Canceled` to `OutcomeCanceled`. **After:** Return
   empty value and `canceled`. **Known correct:** The operation
   stopped; the caller did not merely stop waiting for it.
3. **Before:** The target runtime has an owned cancellation token.
   **Decision:** Run the same two cases and map target outcomes to the
   shared codes. **After:** Both exact outcomes match Go, and the wait
   has ended. **Known correct:** Equivalent semantics, not similar
   syntax, define the successful port.

### Why it works

The fact that must stay true is the observable contract plus owned
resource lifetime, not a Go keyword. This portability invariant is
preserved by identifying semantic gaps and adapting them explicitly.
Shared conformance cases then verify outcomes, cancellation, and
release behavior on both runtimes.

### Go example

See the portable operation outcome contract in
[the delivery examples][delivery-examples]. Its
[tests][delivery-tests] preserve ordinary success and cancellation as
stable outcomes.

### Work and memory

The outcome adapter adds `O(1)` work and memory beyond the operation.
For a port, let `n` be input bytes, `k` concurrent tasks, and `r`
retained objects. Recalculate costs for allocation, garbage
collection, task or thread scheduling, string encoding, collection
behavior, and input/output (I/O) drivers; target defaults can change
constants or effective bounds in `n`, `k`, or `r`.

### When to use it

- A service or library is moving to another runtime.
- A multi-language system must share one boundary contract.
- Review must compare ownership, errors, cancellation, or concurrency.

### When not to use it

Do not choose a language from syntax familiarity alone. Do not assume
similarly named concurrency, collection, or cancellation features have
identical ownership and lifetime guarantees.

### Common mistakes and edge cases

> **Warning:** Translating 100,000 goroutines to 100,000 operating
> system threads can exhaust memory and scheduling capacity. Measure
> the target task model and bound concurrency.

> **Warning:** Mapping a Go domain error to an unchecked exception can
> turn a declared outcome into process-level failure. Preserve the
> boundary's stable error taxonomy explicitly.

> **Warning:** Assuming target strings, maps, nulls, and value copies
> share Go semantics can introduce aliasing or missing-value bugs.
> Write conformance cases for each relied-on behavior.

> **Warning:** Dropping cancellation because a library uses a
> different API leaves work running after the request ends. Adapt the
> owned token through every blocking boundary and verify release.

> **Warning:** Comparing a warmed Go benchmark with a cold target
> runtime measures setup differences, not the port. Use equivalent
> workloads, warmup, limits, outputs, and measurement windows.

### Related material

See [Go ownership](#1-go-values-program-structure-and-ownership),
[Go cancellation](#2-go-errors-resources-and-cancellation), and
[testing](#10-deterministic-backend-testing).

[go-refresher]: #1-go-values-program-structure-and-ownership
[go-foundations]: #2-go-errors-resources-and-cancellation
[distributed-reference]: #18-distributed-systems-mechanics
[extraction-reference]: #19-service-boundary-discovery-and-extraction
[scale-reference]: #24-data-scaling-caching-and-tenant-isolation
[integration-reference]: #29-production-integration-boundaries
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
