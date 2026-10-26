# Backend Concepts Entry Specifications

This document is the Phase 4 deliverable for the
[incorporation plan](CONCEPTS_INCORPORATION_PLAN.md). Entry order and
ownership follow the [guide design](CONCEPTS_GUIDE_DESIGN.md) and
[ownership map](CONCEPTS_OWNERSHIP_MAP.md).

Every specification names the intended content for all eleven exact
teaching headings. `Trace` supplies the required Before, Decision,
After, and Known correct state. `Validation` defines how the final Go
fragment is compiled and exercised in an ordinary and failure case.

## 1. Go values, program structure, and ownership

Purpose: Help a learner predict which Go values, storage, and package
APIs are shared or copied before backend code relies on them.

Prerequisites: Basic programming syntax. Owns C03.

1. The problem it solves: Hidden aliasing, visibility, Unicode, nil,
   and receiver behavior make boundaries mutate or retain data.
2. The idea in simple terms: A Go value has a concrete representation
   and owner; copying the outer value may still share nested storage.
3. How it works: Define modules, packages, exports, values, pointers,
   arrays, slices, maps, strings, structs, methods, interfaces, zero
   values, `make`, and `new` in the order the trace uses them.
4. Step-by-step examples: Clone a product with nested tags, mutate the
   caller copy, and contrast a missing map key with stored zero.
5. Why it works: Explain descriptors, backing storage, method sets,
   dynamic interface type/value, and deliberate ownership transfer.
6. Go example: A complete `catalog` package with `Product`, `Clone`,
   stable map output, and a checked lookup.
7. Work and memory: Count shallow value copies, nested slice copy,
   retained capacity, sorting, string/rune conversion, and output.
8. When to use it: Any API crossing package, goroutine, cache, or
   persistence boundaries with mutable collections or optional values.
9. When not to use it: Avoid pointers, interfaces, generics, or deep
   copies without mutation, substitution, or ownership need.
10. Common mistakes and edge cases: Slice aliasing after append, nil
    map writes, map order, typed nil, byte truncation, and shadowing.
11. Related material: Link the two Go source lessons, portability,
    errors/resources, and concurrency owners.

Trace: Input product tags `tea, sale`; output clone tags `tea, sale`
after caller mutates its own first tag to `coffee`. Show the two slice
headers before copy, clone allocation decision, distinct storage
after,
and which original and clone values are known correct.

Validation: Compile a temporary `catalog` package; exercise clone
independence and a missing-versus-zero map lookup.

Anchors: new heading fragment only.

## 2. Go errors, resources, and cancellation

Purpose: Make failure identity and operation lifetime survive several
backend layers without leaking or abandoning resources.

Prerequisites: Entry 1. Owns C04.

1. The problem it solves: A request can fail after acquiring files,
   rows, timers, or remote work, and callers need both context and a
   stable cause.
2. The idea in simple terms: The acquirer owns cleanup and passes one
   cancellation signal downward while errors travel upward with their
   identity intact.
3. How it works: Acquire, check, defer release, propagate context,
   create shorter child deadlines only at owning layers, wrap with
   `%w`, and classify with `Is`, `As`, or `AsType` at the decision
   edge.
4. Step-by-step examples: Load one product before and after request
   cancellation; include a typed validation error and sentinel
   absence.
5. Why it works: Explain error chains, last-in-first-out cleanup,
   cooperative cancellation, timer release, and root ownership.
6. Go example: A complete function waiting on a result channel or
   context, wrapping `ErrNotFound`, and stopping its child timer.
7. Work and memory: Count one timer, context nodes, error wrappers,
   waiting goroutine lifetime, scarce resource hold time, and cleanup.
8. When to use it: Every operation that waits, performs I/O, or owns a
   resource whose caller may need to classify failure.
9. When not to use it: Do not add context to instant pure functions,
   use it as a parameter bag, or recover ordinary operational errors.
10. Common mistakes and edge cases: `%v` loses identity, defer before
    successful acquire, ignored close/iteration error, loop defers,
    `Background` mid-request, stored context, and typed-nil error.
11. Related material: Link Go ownership, HTTP lifecycle, PostgreSQL
    cleanup, remote deadlines, and source lessons.

Trace: Input a request canceled before a result arrives; expected
output `context.Canceled`. Show result and cancellation not ready,
select decision, cancellation ready, returned error, and that no
result
was consumed or resource retained.

Validation: Compile a temporary package; exercise a delivered result
and cancellation before delivery.

Anchors: new heading fragment only.

## 3. Go concurrency and executable verification

Purpose: Preserve an in-memory invariant under overlap and distinguish
that proof from the absence of data races.

Prerequisites: Entries 1 and 2. Owns C05.

1. The problem it solves: Individually synchronized reads and writes
   can still interleave into duplicate creation or negative inventory.
2. The idea in simple terms: Protect the whole decision, give every
   goroutine an owner and ending, then test both memory and business
   outcomes.
3. How it works: Define shared state, select mutex versus channel,
   lock check-and-change together, bound workers, join completion, and
   run functional, contention, race, vet, and optional fuzz checks.
4.  Step-by-step examples: Two workers create one SKU; exactly one
   wins.
   Include a worker that returns early but still signals completion.
5. Why it works: State the uniqueness fact before locking, show it
   initially, preserved by exclusive check-insert, terminated by
   joins,
   and reflected in exactly one stored product.
6. Go example: A complete bounded concurrent `Store.Create` and
   contention test using `sync.Mutex` and `sync.WaitGroup`.
7. Work and memory: Constant locked map operation on average, bounded
   worker stacks, contention wait, snapshot/sort cost, and test
   tooling.
8. When to use it: Shared in-process state, worker coordination, or a
   characterization of an invariant before moving it to PostgreSQL.
9. When not to use it: An in-memory mutex cannot coordinate several
   processes or replace a durable transaction.
10. Common mistakes and edge cases: Split locks, unbounded goroutines,
    close by receiver, buffered channel as infinite queue, atomics
    over
    multi-value state, sleeps, and race-detector overclaim.
11. Related material: Link Go ownership, transaction concurrency,
    durable jobs, deterministic testing, and both Go source lessons.

Trace: Input two simultaneous creates for `TEA-001`; expected one nil
and one duplicate error. Show unlocked absence, A lock and insert, B
wait, A unlock, B duplicate decision, and one known-correct map entry.

Validation: Compile and run ordinary distinct-SKU and duplicate-SKU
contention tests under `go test -race`.

Anchors: new heading fragment only.

## 4. Domain discovery and modeling

Purpose: Turn vague product statements into reviewed behavior before
transport, schema, or service choices harden assumptions.

Prerequisites: Supporting target Order API. Owns C06.

1. The problem it solves: Different engineers can implement the same
   sentence consistently yet produce incompatible business behavior.
2. The idea in simple terms: Use concrete stories to agree on words,
   states, rules, ownership, and unanswered decisions.
3. How it works: Classify facts and uncertainty, build vocabulary,
   write success/rejection/race examples, name
   commands/events/entities
   and values, draw states, extract invariants, and map contexts.
4. Step-by-step examples: Cancel placed order `O-42`; include repeated
   cancellation and fulfillment-started failure.
5. Why it works: Examples expose choices, invariant statements outlive
   technology, and state transitions enumerate legal and terminal
   paths.
6. Go example: A small typed order state transition function that
   accepts placed-to-canceled and rejects a terminal repeat.
7. Work and memory: Discovery cost grows with actors, states,
   transitions, examples, and review; runtime transition is constant.
8. When to use it: Before new workflows, state changes, service
   boundaries, or disputed terms.
9. When not to use it: Do not introduce every domain-driven term when
   a simple reviewed rule and owner are sufficient.
10. Common mistakes and edge cases: Assumption as fact, endpoint-first
    design, event phrased as command, mutable value object, missing
    repeated/racing behavior, and context equated with service.
11. Related material: Link overview, domain source, architecture,
    transactions, workflows, and leadership.

Trace: Input placed `O-42` owned by `C-17`; expected canceled state
and
released reservation intent. Show state and actor before, owner and
transition decision, canceled state after, then known terminal rule.

Validation: Compile a transition package; exercise valid cancellation
and terminal repeat rejection.

Anchors: new heading fragment only.

## 5. Architecture and service construction

Purpose: Keep HTTP, policy, persistence, and process changes local
without speculative layers or hidden dependency discovery.

Prerequisites: Entries 1 through 4. Owns C07.

1. The problem it solves: Mixed representations and global wiring make
   one contract or query change spread across the service.
2. The idea in simple terms: Put each decision with its reason to
   change and connect concrete parts once at the composition root.
3. How it works: Convert transport to command to domain to rows,
   organize vertical features, point dependencies inward, define a
   small consumer interface, inject concrete dependencies, and record
   lasting tradeoffs.
4. Step-by-step examples: Place one order through handler, service,
   and placement adapter; include invalid input that never calls
   storage.
5. Why it works: Each boundary narrows authority and dependency; the
   use case names one business operation while the adapter hides its
   mechanism; construction makes mandatory inputs visible.
6. Go example: Complete command, narrow `PlacementRepository`, service
   constructor, and method with a recording fake.
7.  Work and memory: Conversion copies bounded request values;
   interface
   dispatch is negligible; transaction/network cost belongs to
   adapters;
   extra layers impose code and test maintenance.
8. When to use it: Capabilities with distinct transport, policy, and
   durable risks or real substitutions.
9. When not to use it: Avoid generic CRUD repositories, service
   locators, identical layers, and interfaces with one speculative
   user.
10. Common mistakes and edge cases: DTO as domain, caller-controlled
    owner, typed-nil dependency, HTTP status in domain, broad
    interface,
    global env reads, vague helpers, and duplicate moved paths.
11. Related material: Link architecture source, service core, domain,
    HTTP, transactions, testing, and extraction counterfactual.

Trace: Input valid command from `C-17`; expected one repository call
with generated ID and fixed UTC time. Show transport discarded,
trusted
owner added, validation decision, copied command, adapter call, and
known dependency direction.

Validation: Compile a service package; exercise valid delegation and
invalid empty input with zero fake calls.

Anchors: old data/transactions/architecture migration note links here.

## 6. HTTP contracts and process lifecycle

Purpose: Give clients one bounded, stable conversation and give the
process a predictable request and shutdown lifecycle.

Prerequisites: Entries 2, 4, and 5. Owns C08.

1. The problem it solves: Untrusted bytes, ambiguous method behavior,
   inconsistent errors, and unbounded waiting turn handlers into
   public
   and operational accidents.
2. The idea in simple terms: HTTP is the whole promise from method and
   body through outcome, correlation, connection, and termination.
3. How it works: Trace network stages, select method semantics, bound
   server and body, decode strictly, separate validation, map
   problems,
   order middleware, propagate context, expose health, and drain.
4. Step-by-step examples: Valid product creation returns 201 and
   Location; a second JSON value returns a stable 400 before the use
   case.
5. Why it works: Each boundary eliminates invalid states and limits
   resource lifetime; stable identities allow translation; readiness
   and shutdown coordinate traffic rather than business policy.
6. Go example: The existing complete strict decoder, revised with
   content-type and classified size-error contract if needed.
7. Work and memory: Body processing is linear in bounded bytes;
   decoding allocates the request; network, auth, database, response,
   open connections, and drain time have separate limits.
8.  When to use it: Every browser, service, webhook, upload, and
   retryable
   operation needs an explicit protocol contract.
9. When not to use it: Do not expose persistence generically or force
   request-response JSON onto streaming requirements.
10.  Common mistakes and edge cases: CORS as auth, status/body
    conflict,
    header after write, leaked response body, raw errors, wrong
    middleware order, database-backed liveness, offset cursor, and
    assuming shutdown cancels active work.
11. Related material: Link HTTP source, frontend network owner,
    architecture, identity, data, and delivery.

Trace: Input `POST /v1/products` with two JSON values; expected 400
`invalid_json`. Show bounded unread bytes, first decode, second-decode
decision, problem response, and known absence of application call.

Validation: Compile decoder package; exercise one valid JSON value and
the trailing-value rejection.

Anchors: preserve `1-http-api-contracts` before this entry.

## 7. Relational modeling and PostgreSQL access

Purpose: Make durable facts, constraints, queries, pools, migrations,
and pagination remain correct under concurrent writers and real data.

Prerequisites: Entries 4 through 6. Owns C09.

1.  The problem it solves: Struct-shaped tables, application-only
   checks,
   unbounded pools, unsafe SQL, and untested migrations corrupt or
   stall
   shared state.
2. The idea in simple terms: Model business facts in relations and let
   PostgreSQL be the final shared authority while adapters manage
   access.
3. How it works: Define keys and constraints, preserve history, choose
   null deliberately, add query-shaped indexes, migrate compatibly,
   budget the pool, parameterize values, clean rows, test PostgreSQL,
   paginate by tuple, and inspect representative plans.
4.  Step-by-step examples: Store an order's captured price after
   catalog
   price changes; reject duplicate SKU directly in SQL.
5.  Why it works: Constraints arbitrate all writers, parameters
   separate
   syntax and data, pool bounds cap concurrency, ordered migrations
   reproduce shape, and total ordering fixes page boundaries.
6.  Go example: A compact context-aware `GetBySKU` adapter fragment
   that
   scans fields, maps no rows, and wraps unexpected errors; link full
   SQL
   owners instead of copying migrations.
7.  Work and memory: Count network round trips, pool wait, scanned
   rows,
   index read/write/storage, result allocation, sorting, migration
   locks,
   and query-plan execution; define rows and page size.
8. When to use it: Durable relational facts with constraints, joins,
   transactional access, stable pagination, or concurrent writers.
9. When not to use it: Do not add indexes without access evidence or
   represent large opaque objects as rows when object storage is
   owner.
10. Common mistakes and edge cases: Current price as history, nullable
    by default, duplicate index, migration at every startup, pool per
    request, interpolated sort, unchecked rows error, N+1, tiny-plan
    overclaim, and in-memory DB substitute.
11. Related material: Link PostgreSQL source, persistence core,
    transactions, advanced operations, security, and delivery.

Trace: Input order item at 500 cents and later catalog price 700;
expected historical total 1,000 for quantity two. Show facts before,
capture decision, changed catalog after, and known order total from
its
own row.

Validation: Compile an adapter-shaped package with a stub row scanner;
exercise successful scan and preserved not-found identity. Real SQL
behavior remains owned by source integration tests.

Anchors: old data/transactions/architecture migration note links here.

## 8. Business transactions and contested state

Purpose: Preserve one multi-record business invariant when requests
overlap or any step fails.

Prerequisites: Entries 5 and 7. Owns C10.

1. The problem it solves: Separate read and write or per-repository
   commits let two buyers spend one unit or expose a partial order.
2.  The idea in simple terms: One transaction protects the whole
   action,
   and one atomic condition or lock protects each contested decision.
3. How it works: Validate input early, begin once, sort resource IDs,
   conditionally reserve and capture price, check overflow, insert
   order
   and items, commit once, map known outcomes, and rollback every
   other
   path.
4. Step-by-step examples: Two buyers race for stock one; include item
   insertion failure after reservation and verify full rollback.
5. Why it works: The database evaluates availability and change under
   one row decision; a transaction publishes all durable facts or
   none;
   ordered locks avoid cycles for that path; owner predicate prevents
   unauthorized load.
6. Go example: Keep the existing mutex reservation as the small
   executable analogue and link the source's conditional SQL and
   PostgreSQL integration owner.
7. Work and memory: In-memory lock work is constant; database cost
   includes connection, round trips, row locks, WAL, indexes, wait,
   rollback, and sorting `p` products in `O(p log p)`.
8. When to use it: Inventory, balances, quotas, uniqueness, state
   transitions, and any local invariant spanning reads or writes.
9.  When not to use it: Never hold a local transaction across slow
   remote
   calls or pretend it covers another service database.
10. Common mistakes and edge cases: Pre-read reservation, independent
    repository commits, ignored commit error, different lock order,
    partial-statement retry, overflow, long transaction, ambiguous
    client result, and load-then-owner-check.
11.  Related material: Link data source, architecture, advanced
    database,
    idempotency, workflows, and transaction tests.

Trace: Input stock one and two reservations of one; expected one
order,
one insufficient result, stock zero. Show A conditional update, B
wait,
A commit, B condition false, and known nonnegative inventory.

Validation: Compile reservation package; exercise successful reserve
and insufficient failure without mutation.

Anchors: old data/transactions/architecture migration note links here.

## 9. Identity, authorization, and API security

Purpose: Turn untrusted credentials into narrowly authorized actions
without accepting the wrong token, role, owner, or browser assumption.

Prerequisites: Entries 5 through 8. Owns C11.

1. The problem it solves: A genuine token can target another API, a
   known user can lack a role, and an allowed role can target another
   user's object.
2. The idea in simple terms: Prove identity completely, reduce it to a
   trusted principal, then decide action and object access separately.
3. How it works: Define OAuth/OIDC participants, validate the full JWT
   conjunction, cache and rotate JWKS safely, construct a principal,
   enforce role in the use case, scope SQL by subject, choose token
   and
   browser controls, prevent exposure, test negatives, and
   threat-model.
4. Step-by-step examples: `C-17` reads owned `O-42`; include `C-9`
   guessing the same ID and an unknown signing key during issuer
   outage.
5.  Why it works: Signature and claims bind token authority and
   purpose;
   least authority flows inward; owner predicates make forbidden data
   unobservable; fail-closed refresh never weakens verification.
6. Go example: A complete `Principal`, typed context helper, and role
   guard; link maintained JWT validation rather than teaching crypto.
7. Work and memory: Local JWT verification uses bounded token and key
   cache work; opaque lookup adds network latency; authorization adds
   query predicates; rate, cache, telemetry, and rotation consume
   state.
8.  When to use it: Every protected user, administrator, service,
   tenant,
   object, browser session, or secret-bearing API boundary.
9. When not to use it: Do not build password or token issuance into a
   resource-server lesson or treat network reachability and CORS as
   auth.
10. Common mistakes and edge cases: Signature-only acceptance,
    algorithm from header, ID token as access token, per-request JWKS,
    raw claims everywhere, 401/403 confusion, UI-only role, post-load
    owner check, wildcard credentialed CORS, mass assignment, token
    logs.
11. Related material: Link identity source, HTTP, architecture, data,
    advanced security, frontend browser security, and threat model.

Trace: Input valid token for `C-9` and `O-42` owned by `C-17`;
expected
the same 404 as nonexistent. Show verified principal, scoped query
decision, zero rows after, and known absence of leaked order state.

Validation: Compile principal/guard package; exercise permitted admin
and forbidden customer paths.

Anchors: new heading fragment only.

## 10. Deterministic backend testing

Purpose: Put each important risk at the cheapest boundary that can
actually reproduce it and make the result replayable.

Prerequisites: Entries 1 through 9. Owns C12.

1. The problem it solves: Broad tests diagnose poorly, mocks invent
   dependencies, and wall time, random IDs, shared fixtures, and
   sleeps
   make failures non-repeatable.
2. The idea in simple terms: Test the behavior where it lives and
   control every input that affects the assertion.
3. How it works: Classify risk into unit, handler, integration, or
   acceptance scope; use named tables and small consumer fakes; inject
   clocks, IDs, randomness, and scheduling; isolate fixtures and bound
   waits.
4. Step-by-step examples: Validate quantities with a table; include a
   two-transaction barrier for last-unit contention and a canceled
   wait.
5. Why it works: Narrow tests isolate one cause, real dependencies own
   real mechanisms, and controlled transition points reproduce the
   intended interleaving instead of hoping the scheduler does.
6. Go example: Complete table-driven validation test and a recording
   placement fake with fixed clock and sequence ID.
7. Work and memory: Unit tests are local; integration adds process,
   database, migration, network, and fixture cost; acceptance adds the
   whole path; parallelism consumes isolated dependency capacity.
8. When to use it: Every invariant, boundary contract, regression,
   migration, concurrency behavior, and user-critical workflow.
9. When not to use it: Do not use a fake to prove PostgreSQL or repeat
   every small branch through a slow end-to-end environment.
10. Common mistakes and edge cases: Assert incidental text, behavior-
    heavy mock, reflexive `t.Parallel`, shared tables, wall tolerance,
    random fixture dumps, unbounded readiness loop, sleep-based race,
    regression test that never failed for the intended reason.
11.  Related material: Link testing source, Go verification, handler
    and
    database tests, observability drills, maintenance
    characterization.

Trace: Input fixed time, ID `order-123`, and invalid quantity zero;
expected stable validation error and zero repository calls. Show fake
at zero, validation decision, unchanged fake, and known no side
effect.

Validation: Compile and run table cases for positive and invalid
quantity plus exact fake-call assertion.

Anchors: new heading fragment only.

## 11. Observability and operational controls

Purpose: Let an operator detect, diagnose, mitigate, and verify one
user-visible failure without exposing sensitive data.

Prerequisites: Entries 6, 7, 9, and 10. Owns C13.

1. The problem it solves: A failed request can leave isolated logs,
   averages, or component alarms that neither show scope nor cause.
2. The idea in simple terms: Logs tell events, metrics show population
   trends, traces connect one path, audits record sensitive actions,
   and runbooks turn impact into tested action.
3. How it works: Instrument stable boundaries, correlate safely,
   define bounded dimensions, measure
   latency/traffic/errors/saturation,
   define SLI then SLO and budget, separate health protocols, inject
   failures, alert on action, and verify recovery.
4. Step-by-step examples: One order waits 750 ms for a pool and fails;
   include invalid input with no database span and a recovered panic.
5. Why it works: Each signal answers its designed question, a shared
   correlation path joins them, and objective/alert/runbook policy
   maps
   aggregate impact to a bounded response.
6. Go example: A complete bounded access-event builder that accepts
   route template and public code and refuses raw URL/body/token
   fields.
7. Work and memory: Count log bytes, metric series cardinality, span
   count and sampling buffers, exporter queues, probe load, retention,
   alert evaluation, and operator attention.
8. When to use it: Every production request, dependency, worker,
   privileged action, SLO, alert, incident, and recovery claim.
9. When not to use it: Do not add fields without an operator question
   or page on every internal anomaly.
10. Common mistakes and edge cases: Raw path or user ID metric label,
    repeated layer logs, SQL parameters in spans, audit as access log,
    average-only latency, undefined denominator, database liveness,
    unsafe runbook, alert without recovery criteria.
11. Related material: Link observability sources, testing, delivery,
    distributed observability, reliability, and source evidence.

Trace: Input one slow order; expected pool-wait cause. Show healthy
baseline, latency metric decision, trace span wait after, correlated
log
outcome, and known widespread versus request-specific facts.

Validation: Compile event-builder package; exercise safe fields and
rejection/redaction of an authorization field.

Anchors: preserve `3-production-and-distributed-systems` with a
migration note immediately before this production sequence.

## 12. Delivery, deployment, and recovery readiness

Purpose: Promote known bytes through compatible state changes and
prove startup, shutdown, rollback direction, restore, and handoff.

Prerequisites: Entries 5 through 11. Owns C14.

1. The problem it solves: A locally passing build says little about
   artifact identity, runtime config, mixed versions, data loss,
   process
   termination, or another engineer's ability to recover it.
2.  The idea in simple terms: Delivery is one evidence chain from
   source
   and artifact through rollout, observation, and rehearsed recovery.
3. How it works: Build once, minimize and inspect the non-root image,
   validate typed config, share local/CI commands, expand and contract
   schema, choose rollback or forward, restore in isolation, model
   capacity, rehearse lifecycle and incident, and clean-clone the
   docs.
4. Step-by-step examples: Add nullable description across old/new
   versions; include missing config and database readiness
   loss/recovery.
5. Why it works: Identical bytes preserve provenance, compatibility
   keeps a mixed fleet valid, early config failure prevents partial
   service, and restore plus workflow verification proves usable data.
6. Go example: A complete typed duration/count configuration parser
   that fails before listening without returning secret values.
7. Work and memory: Count build and scan time, image bytes, rollout
   capacity, dual-write/backfill work, database locks/WAL, restore
   data
   transfer, connection fleet budget, termination and incident time.
8. When to use it: Every deployable service, schema release, secret,
   backup, capacity claim, incident path, and operator handoff.
9. When not to use it: Do not add orchestration only to demonstrate
   deployment; use the simplest environment that exercises the
   contract.
10. Common mistakes and edge cases: Rebuild per environment, source in
    runtime image, root process, silent defaults, CI-only workflow,
    migration in every replica, immediate contract, unsafe rollback,
    backup without restore, false readiness recovery, hidden setup.
11. Related material: Link delivery sources, HTTP lifecycle, data
    migration, performance, cloud, platform, reliability, maintenance.

Trace: Input Release A, expanded schema, and Release B; expected both
versions operate until backfill and switch finish. Show old shape,
additive decision, mixed state, verified new use, and known rollback
window before contraction.

Validation: Compile config package; exercise valid startup and missing
required value without secret disclosure.

Anchors: production migration note links here.

## 13. Idempotency and ambiguous write outcomes

Purpose: Let a caller repeat one logical write safely when it cannot
know whether the first attempt committed.

Prerequisites: Entries 6, 8, and 9. Owns C15.

1. The problem it solves: Lost responses, disconnects, crashes, and
   concurrent retries can create duplicate orders or hide a committed
   outcome.
2. The idea in simple terms: Give the operation a scoped name, bind it
   to one normalized request, and remember its state and answer with
   the
   business effect.
3. How it works: Validate and scope key, canonicalize command, hash,
   acquire record inside transaction, conflict on changed fingerprint,
   serialize matches, store completed response, define in-progress and
   expiry behavior, and clean with bounds.
4.  Step-by-step examples: Two equal concurrent `K1` requests create
   one
   order; include same key with changed quantity and retry after
   commit
   but lost response.
5. Why it works: Unique scoped identity serializes logical work, the
   fingerprint prevents accidental aliasing, and atomic outcome
   storage
   closes the effect/record crash gap.
6. Go example: A complete canonical command serializer and SHA-256
   fingerprint comparison returning replay, conflict, or new.
7. Work and memory: Canonicalization and hashing are linear in bounded
   input; transaction adds lock wait and storage; response, retention,
   indexes, cleanup, and duplicate waiting consume capacity.
8. When to use it: Retriable order, payment-like, provisioning, and
   external-effect writes with ambiguous outcomes.
9.  When not to use it: Safe reads and naturally idempotent
   replacements
   may not need stored response records; never use it to retry
   permanent
   validation or authorization failure.
10. Common mistakes and edge cases: Global key scope, raw JSON hash,
    changed payload replay, record committed separately, no concurrent
    serialization, response omitted, indefinite in-progress, silent
    expiry, unbounded retention, unowned cleanup.
11. Related material: Link production source, HTTP method semantics,
    transactions, retries, external integrations, jobs, and workflows.

Trace: Input two `C-17/K1` commands with equal meaning but reordered
JSON; expected one effect and compatible responses. Show no record,
first insert/lock, second canonical match wait, completed response,
and
known one order.

Validation: Compile fingerprint package; exercise equivalent field
order match and changed quantity conflict.

Anchors: production migration note links here.

## 14. Performance, capacity, retry, and overload

Purpose: Keep the service within resource objectives and change only a
measured limiting cause without multiplying failure load.

Prerequisites: Entries 7, 10 through 13. Owns C16.

1. The problem it solves: Average latency, guessed optimization,
   unbounded work, layered retries, or speculative caches hide the
   first
   bottleneck and worsen overload.
2.  The idea in simple terms: Measure realistic arrivals and tail
   delay,
   bound admitted work, inspect the waiting resource, change one
   cause,
   and repeat the same experiment.
3. How it works: Define workload and targets, generate representative
   data, avoid coordinated omission, collect profiles and plans, model
   fleet capacity, bound concurrency/queues/deadlines, classify one
   retry owner, shed excess, test recovery, and justify cache policy.
4. Step-by-step examples: Pool wait rises with moderate CPU; add the
   missing query index and rerun. Include ambiguous write timeout that
   is not retried without an idempotency key.
5. Why it works: Fixed workload makes before/after comparable, queue
   bounds prevent positive feedback, one total deadline caps attempts,
   and authoritative state makes cache failure affect speed not truth.
6.  Go example: The existing bounded retry-delay policy plus a
   remaining-
   deadline guard and deterministic jitter input.
7. Work and memory: Count arrivals, in-flight work, queue, workers,
   connections, attempts, delay, cache entries, network calls, profile
   overhead, telemetry, and the first dependency's capacity.
8. When to use it: Measured latency, throughput, saturation, quota,
   overload, or repeated-read problems with explicit targets.
9.  When not to use it: Do not tune a microbenchmark while database
   wait
   dominates, enlarge every pool, retry business failures, or cache
   without consistency and removal design.
10.  Common mistakes and edge cases: Average only, closed-loop
    omission,
    unrepresentative tiny data, many simultaneous changes, layered
    retries, attempt after deadline, queue without bound, scale only
    API,
    stale cache for inventory, stampede, no recovery measurement.
11. Related material: Link performance source, pool/data, idempotency,
    observability, delivery capacity, reliability, and scaling.

Trace: Input three allowed attempts and 600 ms total; expected no new
attempt when 10 ms remain. Show remaining time, transient decision,
computed delay/call budget, stop after, and known bounded total work.

Validation: Compile retry package; exercise capped delay and
insufficient
remaining-budget exit.

Anchors: production migration note links here.

## 15. Advanced PostgreSQL operation and recovery

Purpose: Operate PostgreSQL as concurrent, maintained, and recoverable
state rather than a passive query endpoint.

Prerequisites: Entries 7, 8, 10, 12, and 14. Owns C17.

1. The problem it solves: Snapshot anomalies, cyclic locks, bloat,
   checkpoints, long backfills, corruption, and failover can
   invalidate
   application assumptions despite correct SQL syntax.
2. The idea in simple terms: Know which history each transaction sees,
   which work may wait or abort, how storage is maintained, and which
   retained records can rebuild a chosen point.
3. How it works: Select isolation from anomalies, order locks, bound
   statement/lock waits, retry whole safe operations, inspect activity
   and maintenance, batch and checkpoint backfills, archive WAL,
   restore
   PITR, promote standby, reconnect clients, and measure objectives.
4. Step-by-step examples: Two sessions deadlock on opposite rows then
   succeed with sorted order; include interrupted backfill and PITR
   before an accidental deletion.
5. Why it works: MVCC supplies snapshots, PostgreSQL aborts cycles and
   unsafe serial outcomes, idempotent key ranges resume, and base
   backup
   plus WAL reconstructs an earlier durable sequence.
6.  Go example: A complete bounded whole-operation retry classifier
   that
   recognizes only configured serialization/deadlock codes.
7.  Work and memory: Count retained row versions, locks and waits,
   retry
   work, vacuum I/O, bloat storage, WAL/checkpoint write latency,
   batch
   size, archive transfer, replica lag, and restore duration.
8. When to use it: Contended transactions, large live migrations,
   query-plan drift, bloat, replicas, PITR, failover, and RPO/RTO
   claims.
9. When not to use it: Do not select serializable as a magic switch,
   retry one statement, run one giant backfill transaction, or call a
   replica a backup.
10. Common mistakes and edge cases: Long transaction blocks cleanup,
    opposite lock order, unlimited retry, ignored lock timeout, stale
    statistics, unverified backup segment, replica copies deletion,
    promotion without client reconnect, restore without app
    invariants.
11. Related material: Link advanced database source, relational data,
    transactions, delivery restore, performance, and reliability.

Trace: Input A locks tea then cup and B locks cup then tea; expected
one
deadlock victim. Show first locks, cyclic wait decision,
abort/release,
whole-operation retry in sorted order, and known committed invariant.

Validation: Compile retry classifier; exercise configured
serialization
code and permanent constraint failure.

Anchors: production migration note links here.

## 16. Durable background jobs and scheduling

Purpose: Make required work survive process restarts and remain safe
when a worker attempt overlaps a newer owner.

Prerequisites: Entries 8, 10, 13 through 15. Owns C18.

1. The problem it solves: A goroutine disappears on restart, workers
   crash between effect and acknowledgement, leases expire without
   stopping old work, and poison jobs can block healthy work.
2. The idea in simple terms: Store intent, lease a bounded claim, give
   each attempt a fresh authority token, make the effect repeat-safe,
   and acknowledge only the current claim.
3.  How it works: Define bounded versioned job state, insert with
   source
   transaction, claim batches with lease/token, run bounded workers,
   classify outcomes, back off and jitter, dead-letter, replay with
   audit, schedule idempotently or fence, measure age/backlog, and
   drain.
4. Step-by-step examples: Worker sends email then crashes before ack;
   include paused stale worker after a new claim and a permanent
   poison
   job among healthy jobs.
5.  Why it works: Durable rows survive restart, claim predicates
   prevent
   stale acknowledgement, idempotent identities absorb repeated
   effects,
   and bounded retries/quarantine guarantee finite resource use.
6. Go example: A complete claim-token state transition that accepts
   completion only for the current token and classifies retry state.
7.  Work and memory: Count queue rows and payload bytes, claim query
   and
   locks, worker concurrency, external attempts, delays, dead letters,
   retention, replay, scheduler leases, queue age, and shutdown time.
8. When to use it: Email, export, cleanup, reconciliation, scanning,
   and other effects that must outlive one request or process.
9. When not to use it: Keep immediate required local changes in the
   transaction; do not replace durable intent with a goroutine.
10. Common mistakes and edge cases: Unbounded payload, lease as kill,
    no fencing token, ack without claim predicate, non-idempotent
    effect,
    retry every error, infinite attempts, no age metric, unsafe
    replay,
    single-process scheduler boolean, claims during shutdown.
11.  Related material: Link durable source, Go concurrency,
    idempotency,
    transactions, events, distributed mechanics, and observability.

Trace: Input claim A expired, then claim B issued; expected late A ack
rejected and B retained. Show token A before, expiry/reclaim decision,
token B after, A mismatch, and known current ownership.

Validation: Compile job-state package; exercise current-token
completion
and stale-token rejection.

Anchors: production migration note links here.

## 17. Reliable event publication and consumption

Purpose: Preserve committed publication intent and make duplicate
delivery harmless across producer, broker, and consumer databases.

Prerequisites: Entries 8, 13, 15, and 16. Owns C19.

1. The problem it solves: Publish after commit can lose an event;
   publish before commit can announce state that rolls back; relay and
   consumer crashes create duplicates.
2.  The idea in simple terms: Commit the fact and outbox together,
   relay
   until accepted, and commit consumer dedupe with its own state
   change.
3.  How it works: Define immutable versioned event identity and
   ordering
   key, write outbox locally, claim and relay with retry, publish
   stable
   identity, consume through inbox transaction, make provider effects
   repeat-safe, evolve compatibly, replay bounded ranges, and retain.
4. Step-by-step examples: Commit `O-42/E-91`, crash before relay, then
   publish; include crash after broker acceptance and concurrent
   redelivery to one consumer.
5. Why it works: Local transactions close each database gap, stable ID
   connects retries, inbox and effect commit together, and scoped
   order
   prevents unrelated aggregates from blocking.
6. Go example: A complete consumer transition that inserts an event ID
   and applies state through one transaction interface; duplicate
   returns
   success without another effect.
7. Work and memory: Count outbox/inbox rows, payload and broker bytes,
   relay claims and attempts, partition ordering, lag, consumer state,
   retention, replay traffic, dead letters, and reconciliation.
8. When to use it: A committed fact must reach another process or
   rebuildable derived view despite crashes and duplicate delivery.
9. When not to use it: A local synchronous invariant is safer in one
   transaction; do not add a broker for simple in-process
   notification.
10. Common mistakes and edge cases: Event written after business
    commit, publish-before-commit, mark published before broker, inbox
    separate from effect, global ordering, mutable semantic field,
    replay bypasses dedupe, exactly-once claim across provider effect.
11.  Related material: Link event source, transactions, jobs,
    workflows,
    contracts, integrations, and distributed observability.

Trace: Input committed order and E-91 with relay crash after broker
acceptance; expected repeated E-91 but one consumer state effect. Show
outbox pending, publish decision, crash state, duplicate deliveries,
inbox decision, and known converged consumer state.

Validation: Compile consumer state package; exercise first event and
duplicate event with one counted effect.

Anchors: production migration note links here.

## 18. Distributed-systems mechanics

Purpose: Make partial failure, stale authority, quorum tradeoffs, key
movement, and consensus assumptions predictable before service design.

Prerequisites: Entries 2, 3, 16, and 17. Owns C20.

1. The problem it solves: Nodes disagree about time and delivery, a
   timeout hides remote outcome, leases overlap, and replicas can
   trade
   freshness for availability.
2. The idea in simple terms: Control message schedules explicitly and
   add one mechanism only after reproducing the failure it addresses.
3. How it works: Model logical messages and time, inject delay/loss/
   duplicate/reorder/partition, distinguish monotonic and wall time,
   lease then fence writes, calculate quorum overlap, define conflict
   repair, measure consistent hashing, and walk Raft safety.
4. Step-by-step examples: Reservation response lost after commit;
   include stale lease holder rejected by larger token and quorum
   loss.
5. Why it works: Explicit schedules make uncertainty observable,
   resource-side fencing orders authority, quorum intersection creates
   overlap under assumptions, and consensus majority/term/index rules
   preserve one committed log.
6. Go example: A deterministic message queue plus protected resource
   that accepts only nondecreasing fencing tokens.
7. Work and memory: Count queued messages and enumerated schedules,
   retries, replica count N, quorum R/W messages and wait, fencing
   state,
   virtual nodes, key movement, consensus log and majority latency.
8. When to use it: Reasoning about remote calls, leases, job claims,
   leader ownership, replicated reads/writes, sharding, or broker
   safety.
9. When not to use it: Do not build custom production consensus or use
   consistent hashing as a correctness, replication, or availability
   guarantee.
10. Common mistakes and edge cases: Timeout equals failure, wall-clock
    order, symmetric partition assumption, lease without fencing,
    resource ignores token, quorum formula without system assumptions,
    no conflict rule, modulo movement, old leader message after new
    term.
11. Related material: Link mechanics source, jobs, events, synchronous
    calls, workflows, and maintained production owners.

Trace: Input write from expired token 7 after current token 8;
expected
rejection. Show accepted token 7 before, new lease decision, stored 8
after, late 7 comparison, and known resource state unchanged.

Validation: Compile simulator package; exercise ordered delivery and
stale-fencing-token rejection.

Anchors: production migration note links here.

## 19. Service-boundary discovery and extraction

Purpose: Add a network boundary only when a capability gains measured
ownership or isolation without retaining hidden data and release
coupling.

Prerequisites: Entries 4, 5, 12, 17, and 18. Owns C21.

1. The problem it solves: Splitting controllers or tables pays remote
   failure and operations cost while changes, data, and releases
   remain
   coupled.
2. The idea in simple terms: A service owns one business capability,
   its data, contract, runtime, release, and on-call consequences.
3. How it works: Measure change and operational coupling, compare
   modular alternative, select low-risk asynchronous boundary, define
   data and event ownership, create independent identity/CI/telemetry,
   branch by abstraction, migrate gradually, inject failure, and
   compare.
4. Step-by-step examples: Extract notification while orders keep
   committing; include notification outage and attempted cross-table
   read.
5. Why it works: Durable events remove delivery from order latency,
   credentials enforce data autonomy, staged routing preserves
   fallback,
   and measured outage demonstrates independent failure.
6. Go example: A narrow notification capability with local and event-
   backed adapters selected at the composition root.
7. Work and memory: Add serialization, broker/storage, event lag,
   deployment, telemetry, network latency, on-call, CI, schema,
   secrets,
   duplicate/reconciliation, and migration capacity.
8. When to use it: Independent team ownership, release cadence,
   scaling, security isolation, or failure isolation with measured
   value.
9. When not to use it: Keep a modular monolith when one transaction,
   team, release, and failure domain are simpler and sufficient.
10. Common mistakes and edge cases: Table-per-service split, shared DB
    credentials, synchronous notification on order path, big-bang
    rewrite, no baseline, no fallback, hidden contract, missing
    on-call,
    obsolete path left after migration.
11. Related material: Link extraction source, domain contexts,
    architecture counterfactual, events, delivery, and leadership ADR.

Trace: Input notification provider outage after extraction; expected
order commit and rising notification backlog. Show healthy modules,
routing decision, stopped consumer after, durable event queued, and
known order independence.

Validation: Compile adapter-selection package; exercise local and
remote
adapter selection plus unavailable remote returning without changing
an
already committed order result.

Anchors: new heading fragment only.

## 20. Synchronous service communication

Purpose: Bound and reconcile a remote business operation whose reply
may be late, absent, malformed, or accepted before a local failure.

Prerequisites: Entries 13 through 19. Owns C22.

1. The problem it solves: Extracted inventory removes the local
   transaction, so a lost reservation response leaves ordering unable
   to infer whether stock changed.
2. The idea in simple terms: Persist one operation identity before the
   call, spend one deadline, and query or compensate that exact remote
   effect instead of guessing.
3.  How it works: Own remote contract and data, choose protocol,
   persist
   pending intent, propagate remaining time, bound transport,
   authenticate
   workload, discover/load-balance, retry one safe layer, apply
   pressure
   controls, expose edge policy, record attempts, and reconcile.
4. Step-by-step examples: `reserve-O-42` commits remotely but response
   is lost; include oversized response, wrong-service credential, and
   confirmation failure after successful reservation.
5. Why it works: Stable operation identity makes repeat/query/release
   refer to one effect, durable pending state survives restart, total
   deadline caps waiting, and distinct controls contain different
   loads.
6.  Go example: A complete remaining-deadline calculator and
   reservation
   state reducer using stable operation ID and typed outcomes.
7. Work and memory: Count network hops, serialization bytes,
   connections, discovery, TLS/auth, attempts and delay, concurrency,
   pending rows, circuit state, reconciliation jobs, and edge
   capacity.
8. When to use it: A remote result is required before a user-visible
   next step and cannot be safely replaced by asynchronous completion.
9. When not to use it: Prefer local transaction or durable event when
   immediate remote response is unnecessary; do not hold local locks.
10. Common mistakes and edge cases: Per-attempt instead of total
    timeout, retry new operation ID, layered retries, transaction
    across
    call, network reach as trust, forwarded header as principal,
    gateway
    business rule, breaker as correctness, fallback with false
    semantics,
    orphaned reservation.
11. Related material: Link service source, idempotency, performance,
    mechanics, extraction, workflows, security, and edge contracts.

Trace: Input pending `reserve-O-42`, remote commit, lost response;
expected pending until query returns reserved, then confirm. Show
intent
before, timeout decision, ambiguous after, query outcome, and known
one
reservation identity.

Validation: Compile state reducer; exercise reserved response and
ambiguous timeout requiring reconciliation.

Anchors: new heading fragment only.

## 21. Distributed workflows and eventual consistency

Purpose: Make a multi-service business action survive restarts,
duplicates, delays, timeouts, and failed compensation while remaining
visible and repairable.

Prerequisites: Entries 17 through 20. Owns C23.

1.  The problem it solves: Separate databases cannot roll back
   together,
   so hidden intermediate states and arrival-order handlers corrupt
   cross-service workflows.
2.  The idea in simple terms: A saga is a durable state machine of
   local
   commits; compensation is another fallible business command.
3. How it works: Enumerate states/events/commands/timeouts, bound
   temporary inconsistency, select orchestration or choreography,
   persist version and next action, assign operation IDs, use outbox
   and
   inbox, schedule durable timeouts, reconcile, expose operator state,
   and guard repair by expected state.
4. Step-by-step examples: Reservation succeeds twice before confirm;
   include late success after timeout and duplicate release
   compensation.
5. Why it works: Version predicates allow one legal transition,
   persistent next action survives restart, dedupe absorbs repeats,
   timeouts become durable facts, and guarded repair preserves state.
6. Go example: A complete versioned order-saga transition reducer with
   duplicate and out-of-order handling.
7. Work and memory: Count saga rows and history, outbox/inbox records,
   messages, attempts, scheduled timeouts, compensation calls,
   operator
   indexes, retention, lag, repair, and maximum inconsistent duration.
8. When to use it: One user action spans independently owned services
   and must converge without distributed atomic commit.
9. When not to use it: Use one local transaction while data shares an
   owner; avoid choreography when no one can reconstruct the workflow.
10. Common mistakes and edge cases: Compensation called rollback,
    process timer, state only in messages, no version check, code
    follows
    arrival order, late success ignored, duplicate compensation,
    direct
    database repair, invisible stuck state, unlimited inconsistency.
11. Related material: Link extraction source, domain state machines,
    events, mechanics, synchronous bridge, observability, and
    contracts.

Trace: Input version 2 `reservation_requested` and duplicate success;
expected first moves to reserved version 3 and second is harmless.
Show
version before, guarded decision, state/version after, duplicate
mismatch,
and known single transition.

Validation: Compile saga reducer; exercise valid success and duplicate
or out-of-order no-op/error according to documented state.

Anchors: new heading fragment only.

## 22. Contract evolution and independent delivery

Purpose: Let providers and consumers deploy separately without
mistaking structural additions for semantic compatibility.

Prerequisites: Entries 6, 12, 17, and 19 through 21. Owns C24.

1. The problem it solves: Optional fields, enum values, validation,
   ordering, error, or timing changes can break real consumers despite
   a passing schema diff.
2.  The idea in simple terms: Compatibility is the behavior both old
   and
   new sides can understand throughout a declared support window.
3. How it works: Publish owned versioned HTTP/RPC/event contracts,
   define structural and semantic rules, diff schemas, capture
   consumer
   expectations, test both deployment orders and delayed messages,
   deprecate with inventory/telemetry/date, block uncertain removal,
   and
   delete adapters after the window.
4.  Step-by-step examples: Add optional field consumer-first and
   provider-
   first; include new exhaustive enum and stricter validation
   failures.
5.  Why it works: Tolerant prepared consumers accept absent/new
   optional
   data, provider verification preserves relied-on behavior, and usage
   evidence proves when old behavior no longer serves a consumer.
6. Go example: A complete decoder that tolerates an optional response
   field and an explicit unknown enum fallback rather than panic.
7. Work and memory: Count schema storage, CI diff and provider runs,
   consumer matrix, dual behavior, adapter branches, event retention,
   telemetry, communication, and support-window operating cost.
8. When to use it: Independently deployed APIs, RPC, events, SDKs, or
   stored formats with multiple supported readers and writers.
9. When not to use it: Do not version every internal function or keep
   compatibility code after no supported consumer needs it.
10. Common mistakes and edge cases: Additive equals safe, exhaustive
    enum, optional made required, unit/meaning/default/order change,
    producer-first only, delayed old event, schema-only consumer test,
    deprecation without owner, absence of complaints as zero use.
11. Related material: Link contract source, HTTP, events, delivery,
    extraction, maintenance, and leadership.

Trace: Input old consumer and new provider returning optional
`description`; expected old consumer ignores it. Show old shape
before,
additive decision, new response after, old parse success, and known
unchanged relied-on fields.

Validation: Compile contract package; exercise absent optional field
and
unknown enum fallback.

Anchors: new heading fragment only.

## 23. Distributed observability and resilience experiments

Purpose: Reconstruct a logical operation across services and queues,
then test emergent failure with bounded user and infrastructure risk.

Prerequisites: Entries 11, 13, 14, and 17 through 22. Owns C25.

1. The problem it solves: Local logs cannot show where a saga stopped,
   and independent retries, fan-out, or backlog can create failures no
   component exhibits alone.
2. The idea in simple terms: Carry safe causal context across every
   boundary and run controlled experiments that begin with a
   prediction
   and end with convergence proof.
3. How it works: Propagate trace and correlation metadata, allowlist
   baggage, define service/dependency RED and queue signals, build an
   end-to-end workflow view, choose head/tail sampling bounds, specify
   hypothesis/blast/abort/duration/recovery, inject, observe, and
   verify.
4.  Step-by-step examples: Add 500 ms inventory delay and trace retry
   plus
   queue time; include broken broker propagation and retry storm
   abort.
5. Why it works: Shared causal identity joins spans, stable dimensions
   aggregate population behavior, retained errors expose rare failure,
   and abort/recovery criteria constrain experiments and prove
   healing.
6. Go example: A complete event-envelope correlation propagator with
   an allowlist rejecting credential and arbitrary user baggage.
7. Work and memory: Count span/event metadata bytes, metric series,
   sampling buffer and retention, exporter queue, attempt fan-out,
   backlog/age, experiment traffic, abort reserve, and recovery time.
8.  When to use it: Work crosses HTTP/RPC, broker, job, worker, or
   service
   boundaries and local evidence cannot diagnose user impact.
9.  When not to use it: Do not propagate arbitrary baggage or run
   random
   production sabotage without owner, steady state, and abort.
10. Common mistakes and edge cases: Trace ID regenerated at broker,
    credentials in baggage, raw operation IDs as labels, head sampling
    loses late failure, unbounded tail buffer, retry amplification
    hidden,
    component-first runbook, experiment stops without data
    convergence.
11. Related material: Link distributed source, observability, events,
    workflows, performance, reliability, and security.

Trace: Input one trace through API, order transaction, event, worker,
and inventory; expected all spans queryable as one operation. Show
context before publish, allowlist decision, envelope after, consumer
span link, and known causal chain.

Validation: Compile envelope package; exercise allowed trace metadata
and rejected authorization baggage.

Anchors: new heading fragment only.

## 24. Data scaling, caching, and tenant isolation

Purpose: Reduce measured data pressure without making stale copies or
one omitted tenant scope corrupt authority or expose another customer.

Prerequisites: Entries 7, 9, 14, 17, and 23. Owns C26.

1. The problem it solves: Cache and replica staleness, partition
   complexity, and a tenant filter missed in any sync or async path
   can
   return wrong or cross-tenant data.
2. The idea in simple terms: Scale each bottleneck with a mechanism
   whose consistency limit is explicit, and make trusted tenant
   identity
   part of every ownership key and operation.
3. How it works: Measure, define cache key/version/TTL/invalidation/
   outage/stampede, route replica-tolerant reads by freshness, test
   partition/archive lifecycle, choose tenant model, propagate trusted
   tenant, enforce keys/constraints or transaction-local row security,
   test pooled reuse and privileged bypass, and set quotas.
4. Step-by-step examples: Cache product `P-1` for tenant A and B with
   distinct keys; include stale post-write read, replica lag, and
   pooled
   connection reused without tenant state.
5. Why it works: Authority stays in primary durable state, scoped keys
   prevent aliasing, default-deny predicates and constraints reject
   omitted or cross-tenant ownership, and operation policy controls
   tolerated staleness.
6.  Go example: A complete typed tenant cache-key constructor that
   rejects
   empty trusted tenant identity and includes version.
7. Work and memory: Count cache bytes/entries/eviction, miss and
   single-flight wait, invalidation, replica transfer/lag, partition
   metadata and scans, tenant indexes/constraints, connection-local
   setup, quotas, backups, export, and noisy-neighbor capacity.
8. When to use it: Measurement shows read load, data-management size,
   retention, or tenant isolation/capacity requirements.
9. When not to use it: Do not cache correctness-critical inventory,
   send read-your-write to an unqualified replica, partition small
   data,
   or add tenant ID as a convention without enforcement.
10.  Common mistakes and edge cases: Key omits tenant, stale value
    after
    write, cache outage overload, stampede, correctness read on
    replica,
    cross-partition uniqueness, row-security setting leaks through
    pool,
    privileged bypass untested, job/event tenant mismatch, unscoped
    repair.
11. Related material: Link scale source, relational data, identity,
    performance, workflows, advanced security, and integrations.

Trace: Input tenant A/P-1 and B/P-1; expected two cache entries. Show
unscoped collision before, trusted tenant decision, scoped keys after,
and known no cross-tenant alias.

Validation: Compile key package; exercise two distinct tenant keys and
empty-tenant rejection.

Anchors: new heading fragment only.

## 25. Security, privacy, and software supply chain

Purpose: Limit the reach and lifetime of compromised identities or
data and verify the path from reviewed source to deployed artifact.

Prerequisites: Entries 9, 12, 22 through 24. Owns C27.

1. The problem it solves: New services and stores multiply trust,
   personal-data copies, secrets, lateral movement, and artifacts
   whose
   contents or build history may be unknown.
2.  The idea in simple terms: Give every actor and field only the
   reach,
   purpose, lifetime, and verified lineage it needs.
3. How it works: Update trust/data flow, separate user and workload
   identities, authorize each resource, rotate with overlap, inventory
   sensitive fields and locations, minimize, build durable export and
   deletion, state backup/audit policy, lock dependencies, generate
   SBOM
   and provenance, sign and verify, and drill vulnerability response.
4. Step-by-step examples: Stolen notification credential attempts
   unrelated database access; include rotating credentials and
   deleting
   email across primary, event, provider, cache, derived, and backup
   rules.
5.  Why it works: Narrow audience and resource policy caps blast
   radius,
   overlap preserves availability while revocation closes old trust,
   durable workflows converge copies, and artifact metadata links an
   advisory to every affected build.
6. Go example: A complete least-privilege authorization check for a
   workload/resource/action tuple and rotation state with overlap.
7. Work and memory: Count identity calls, credential overlap, audit
   events, data inventory, deletion tasks and retries, retained backup
   policy, SBOM/provenance/signature bytes, scan time, rebuild,
   rollout,
   and residual-risk tracking.
8. When to use it: Every new service, queue, store, operator tool,
   personal field, secret, CI identity, dependency, image, and
   release.
9. When not to use it: Do not reuse user tokens as general service
   credentials or treat TLS, SBOM, or a scanner as authorization or
   patch.
10. Common mistakes and edge cases: One broad workload credential,
    static key, all-at-once rotation, old key never revoked,
    undocumented
    derived copy, deletion as one SQL statement, immutable retention
    ignored, unsigned deployment accepted, SBOM generated but not
    mapped
    to artifacts, exceptional patch path never rehearsed.
11.  Related material: Link advanced security source, identity,
    tenancy,
    contracts, cloud IAM, integrations, delivery, and maintenance.

Trace: Input old credential A and new B during rotation; expected both
temporarily accepted, then only B. Show A-only before, publish B
decision, overlap after, verified B use, revoke A, and known failed A.

Validation: Compile authorization/rotation package; exercise permitted
new credential and revoked old credential.

Anchors: new heading fragment only.

## 26. Cloud infrastructure and IAM

Purpose: Reproduce a private, least-privilege environment and
distinguish packet reachability from permission to use a resource.

Prerequisites: Entries 12, 25. Owns C28.

1. The problem it solves: Console drift, public data services, static
   credentials, broad roles, unprotected state, and unsafe teardown
   make
   infrastructure irreproducible and hard to diagnose.
2. The idea in simple terms: Define every network path and identity
   permission as reviewed code with explicit failure and ownership.
3. How it works: Choose account/region/zones, trace edge and private
   paths, define network and egress, attach short-lived workload
   roles,
   separate human/admin/key permissions, provision managed resources,
   protect and lock IaC state, plan/policy/review, detect drift, log,
   tag, budget, recreate, and preview scoped teardown.
4. Step-by-step examples: Browser reaches private API through public
   TLS edge and API reaches private PostgreSQL; include route denied
   and
   secret permission denied as distinguishable failures.
5.  Why it works: Routes and firewalls decide packet travel, IAM
   decides
   action authority, versioned desired state recreates topology, and
   locked state plus scoped selection serializes safe changes.
6. Go example: A complete diagnostic error classifier that preserves
   network, authentication, authorization, and application categories;
   link IaC source rather than mimic infrastructure in Go.
7.  Work and memory: Count regions/zones, network endpoints/NAT
   traffic,
   load balancer, database and queue capacity, key/secret operations,
   logs and retention, IaC state/plan, drift scans, headroom, budgets,
   human operation, and teardown risk.
8. When to use it: Any cloud-hosted non-production or production
   environment with shared networks, identities, managed data, or
   cost.
9.  When not to use it: Do not add cloud topology before requirements
   or
   use private IP as permission; managed service does not own app
   recovery.
10.  Common mistakes and edge cases: Public database, static access
    key,
    one admin role, unrestricted egress, secret in state/output,
    unlocked
    state, console-only change, no drift, missing audit retention,
    teardown
    broad selector, network denial mislabeled as app failure.
11. Related material: Link cloud source, advanced security, delivery,
    platform operations, reliability, and repository IaC owners.

Trace: Input API with network route but denied secret permission;
expected IAM-denied diagnostic rather than timeout. Show route before,
identity policy decision, denied API response after, and known packet
path health.

Validation: Compile classifier; exercise timeout/network and explicit
permission-denied categories.

Anchors: new heading fragment only.

## 27. Orchestration and platform operations

Purpose: Reconcile workload processes and releases without hiding the
application contracts or overwhelming the real bottleneck.

Prerequisites: Entries 11, 12, 14, 22, 25, and 26. Owns C29.

1. The problem it solves: Incorrect probes, resources, rollout,
   disruption, scaling, network policy, or termination can send
   traffic
   early, restart healthy apps, lose capacity, or amplify dependency
   load.
2. The idea in simple terms: The platform maintains declared process
   state; the application still tells the truth about initialization,
   traffic safety, restart usefulness, drain, and data compatibility.
3. How it works: Define workloads/services/edge/config/secrets/policy,
   set measured requests and limits, configure
   startup/readiness/liveness,
   align grace/drain, choose surge/unavailable/disruption, autoscale
   from
   bottleneck demand under dependency cap, validate templates, inject
   failures, and write limited-access diagnostics and runbooks.
4.  Step-by-step examples: Start while migration check runs, then
   serve,
   then lose database; include bad readiness path and scale request
   that
   would exceed database connections.
5.  Why it works: Distinct probe contracts select traffic versus
   restart,
   controllers converge replicas, resource reservations guide
   placement,
   caps prevent dependency amplification, and compatible rollout keeps
   old/new instances valid.
6. Go example: A complete lifecycle state object deriving startup,
   readiness, and liveness responses from initialization and draining.
7. Work and memory: Count requested and limited CPU/memory, replicas,
   surge/unavailable capacity, probe traffic, connection fleet,
   scaling
   lag, termination window, telemetry, network policy, storage, and
   control-plane/operator overhead.
8. When to use it: Several workload instances need scheduling,
   discovery, controlled rollout, policy, recovery, or demand scaling.
9.  When not to use it: Do not use orchestration to repair unsafe
   retries,
   false health, invalid config, business state, or incompatible
   schema.
10. Common mistakes and edge cases: Readiness as liveness, no startup
    probe, too-low CPU limit, memory kill surprise, autoscale CPU
    while DB
    bound, fleet pool exceeds DB, disruption plus zero surge loses
    capacity,
    pre-stop longer than grace, opaque templates, admin-only
    diagnosis.
11. Related material: Link platform source, HTTP lifecycle, delivery,
    performance, cloud, reliability, and contract evolution.

Trace: Input initialized service with database down; expected startup
true, readiness false, liveness true. Show serving state before,
dependency decision, health tuple after, and known no restart loop.

Validation: Compile lifecycle package; exercise serving and database-
unready states with liveness preserved.

Anchors: new heading fragment only.

## 28. Reliability engineering, incidents, and cost

Purpose: Set an affordable user reliability target and make pages,
capacity, incident response, reviews, and corrective work serve it.

Prerequisites: Entries 11, 12, 14, 23, 26, and 27. Owns C30.

1. The problem it solves: Maximum uptime, component alerts, unmodeled
   capacity, blame, or cloud-only cost cannot guide delivery and
   recovery.
2. The idea in simple terms: Measure the user's successful operation,
   choose its target and allowed misses, then spend engineering and
   infrastructure where evidence reduces the greatest risk.
3.  How it works: Define eligible SLI/SLO/budget, build exact
   dashboards,
   alert fast and slow burn plus imminent saturation, model
   normal/peak/
   failed-domain capacity, load-test first bottleneck, assign incident
   roles and communication, run game day, measure timeline, review
   system
   conditions, own corrective work, and compare total cost tradeoffs.
4.  Step-by-step examples: Database latency gradually burns order
   budget;
   include one-zone loss and a cost cut that removes required
   headroom.
5. Why it works: The objective gives a common outcome and window,
   multi-window burn balances urgency/noise, capacity connects demand
   to
   bounds, roles prevent coordination collision, and verified actions
   close causes rather than assigning blame.
6. Go example: A complete error-budget calculator using eligible and
   good event counts with explicit zero-denominator handling.
7. Work and memory: Count measurement/storage, dashboard and alert
   evaluation, capacity headroom, failure reserve, on-call/incident
   time,
   mitigation, lost user value, engineering complexity, vendor cost,
   opportunity, and corrective verification.
8. When to use it: Any service with user promises, on-call, growth,
   incident risk, architecture choices, or cost optimization.
9.  When not to use it: Do not set an SLO without users/volume policy
   or
   page on a metric with no safe operator action.
10.  Common mistakes and edge cases: Internal proxy SLI, undefined 4xx
    or
    disconnect denominator, average latency, every error page,
    capacity
    without failure domain, rollback without data check, one person
    owns
    all incident roles, blame root cause, action without
    deadline/proof,
    cheap resource ignores people and SLO risk.
11.  Related material: Link reliability source, observability,
    delivery,
    performance, distributed experiments, cloud, platform, leadership.

Trace: Input 1,000,000 eligible requests and 99.9% SLO; expected
budget
1,000 misses. Show target before, budget calculation decision,
consumed
misses after, burn interpretation, and known remaining allowance.

Validation: Compile budget package; exercise 99.9% calculation and
zero
eligible-event error.

Anchors: new heading fragment only.

## 29. Production integration boundaries

Purpose: Contain third-party ambiguity, hostile callbacks and URLs,
large data, asynchronous delivery, and derived-state drift.

Prerequisites: Entries 9, 13 through 18, 22 through 25. Owns C31.

1. The problem it solves: Providers have independent failures and
   quotas, webhooks are hostile and replayed, subscriber URLs can
   reach
   private systems, uploads exhaust memory, and derived views drift.
2. The idea in simple terms: Give each provider a narrow adapter and
   treat every inbound and outbound effect as bounded, authenticated,
   identifiable, observable, and reconcilable.
3. How it works: Define adapter deadline/quota/idempotency/error/fake/
   sandbox, verify bounded raw webhook before parse, enforce time and
   dedupe, durably enqueue before ack, secure outbound webhook
   destination
   and retry, issue narrow object grants and stream/scan, maintain
   notification preferences/callbacks, checkpoint derived views,
   rebuild,
   and reconcile.
4.  Step-by-step examples: Payment `pay-O-42` times out after
   acceptance;
   include forged/replayed webhook, DNS-rebound subscriber, oversized
   mislabeled upload, duplicate bounce, and deleted search index.
5. Why it works: Stable provider identity resolves ambiguous effects,
   exact-byte signatures bind message authenticity, destination and
   egress policy block SSRF paths, bounded streaming caps process
   memory,
   and authority plus checkpoints can recreate derived state.
6.  Go example: A complete bounded webhook verifier interface and
   replay-
   window/dedup state transition; do not implement cryptography
   manually.
7.  Work and memory: Count provider calls/attempts/quotas, webhook
   bytes,
   keys and overlap, dedupe/history rows, DNS and redirects, object
   bytes
   and scan queue, callbacks/suppression, index documents, lag,
   rebuild,
   reconciliation, and telemetry.
8. When to use it: Payment, shipping, webhook, object, email, search,
   analytics, or other external service boundaries.
9. When not to use it: Do not buffer unbounded files, accept arbitrary
   callback URLs, or store authoritative business truth only in
   search.
10.  Common mistakes and edge cases: Blind timeout retry, signature
    after
    parse, no timestamp, ack before durable work, rotation without
    overlap,
    HTTPS-only SSRF defense, redirect not revalidated, DNS rebinding,
    client MIME trusted, provider accepted equals delivered, derived
    view
    with no rebuild during concurrent new events.
11.  Related material: Link integrations source, HTTP, security,
    durable
    work, events, contracts, tenancy, maintenance, and source owners.

Trace: Input duplicate valid webhook ID E-91; expected first enqueues
one
durable job and second acknowledges without another effect. Show
unseen
ID, verification decision, stored dedupe/work after, duplicate check,
and known one logical delivery.

Validation: Compile webhook state package with fake verifier; exercise
valid new event and replay rejection/no duplicate work.

Anchors: new heading fragment only.

## 30. Maintenance and legacy-system evolution

Purpose: Learn actual relied-on behavior, make changes reversible, and
repair data without expanding scope or leaving permanent migration
debt.

Prerequisites: Entries 10 through 14, 22, and 28. Owns C32.

1. The problem it solves: Intended docs can disagree with users and
   data; refactors, upgrades, flags, canaries, and repairs can break
   contracts or overreach while appearing locally correct.
2. The idea in simple terms: Observe, pin, change a small seam, stage,
   verify, recover once, and delete temporary machinery.
3. How it works: Inventory behavior/risk from every evidence source,
   classify contract/accident/assumption/unknown, add
   characterization,
   refactor without behavior change, analyze upgrade layer, stage flag
   and canary with abort, control shadow effects, build scoped dry-run
   resumable guarded audited repair, diagnose one hypothesis, verify
   recovery, and remove old paths.
4. Step-by-step examples: Preserve empty order list `[]` through a
   refactor; include interrupted repair resume and canary abort after
   data changes between dry-run and execute.
5. Why it works: Characterization alarms on observable change, small
   seams isolate risk, staged cohorts cap exposure, expected-state
   guards
   reject stale repair plans, checkpoints resume, and cleanup returns
   one
   behavior owner.
6.  Go example: A complete bounded repair planner/executor that
   requires
   exact tenant and expected version and supports dry-run checkpoint.
7.  Work and memory: Count discovery and review, characterization
   suite,
   dual paths and telemetry, canary traffic, shadow duplication,
   migration/backfill, repair batches and audit, checkpoint storage,
   rollback/forward time, and ongoing debt until cleanup.
8.  When to use it: Existing user/data behavior, refactors, toolchain
   or
   dependency upgrades, contract migrations, flags, canaries, data
   repair,
   and regressions.
9.  When not to use it: Do not call intentional behavior change a
   refactor
   or keep a flag after supported paths converge.
10. Common mistakes and edge cases: Docs as only truth, broad snapshot
    instead of focused characterization, big-bang rewrite, unreviewed
    release notes, flag no owner/removal, shadow writes cause effects,
    repair ambiguous target, no dry run/resume/audit/kill switch,
    stale
    expected state, fallback masks root cause, old branch remains.
11. Related material: Link maintenance source, testing, delivery,
    contract evolution, reliability, integrations, and leadership.

Trace: Input legacy empty list and proposed null refactor; expected
`[]`
until deliberate migration. Show observed clients before,
characterization
decision, new internals after, same response, and known contract.

Validation: Compile repair package; exercise in-scope dry run and
stale-
version/out-of-scope rejection.

Anchors: new heading fragment only.

## 31. Capstone delivery and engineering leadership

Purpose: Demonstrate judgment that turns a valuable ambiguous problem
into reviewable, operable, evidence-backed, and simpler team
ownership.

Prerequisites: Entries 1 through 30. Owns C33.

1. The problem it solves: Code can ship while alternatives, risk,
   review, mixed versions, failure, telemetry, recovery, cost,
   handoff,
   and cleanup remain unowned.
2.  The idea in simple terms: Trace every important requirement from
   the
   user problem through decision, implementation, proof, production
   signal, operator action, and eventual simplification.
3. How it works: Select measurable problem/non-goals, discover domain,
   write full design and alternatives, request critical reviews,
   record
   finding decisions, plan compatible vertical increments, deliver
   each
   with tests/telemetry/docs, migrate and stage, inject relevant
   failures,
   rehearse recovery, close findings, peer-review, hand off, teach,
   and
   remove temporary complexity.
4. Step-by-step examples: Cancel `O-42` twice while inventory is down;
   include mixed versions, duplicate release, migration interruption,
   first resource exhaustion, and recovery by a non-author.
5.  Why it works: Early alternatives expose cost before code, small
   gates
   preserve stopping points, traceability prevents unsupported claims,
   runtime evidence closes the loop, and handoff plus deletion makes
   ownership durable beyond one author.
6. Go example: A complete requirement-evidence validator using typed
   records for contract, test, signal, runbook, status, owner, and
   risk.
7. Work and memory: Count design/review time, implementation, tests,
   migrations, dual versions, telemetry, failure drills, operations,
   capacity, cloud/provider spend, incident/on-call, handoff,
   teaching,
   residual risk, and saved complexity after cleanup.
8. When to use it: Cross-boundary features with concurrency, auth,
   failure, rollout, operations, and meaningful team ownership.
9. When not to use it: Do not add arbitrary technologies or documents
   to make a small problem appear senior; choose the smallest design.
10.  Common mistakes and edge cases: Solution before problem, no
    non-goal,
    omitted alternative, review after commitment, findings vanish,
    giant
    branch, irrelevant green test as proof, no mixed-version/failure
    case,
    unowned risk, private handoff knowledge, flag/dashboard/scaffold
    left,
    inability to name removable complexity.
11. Related material: Link leadership source, all prerequisite owners,
    review checklist, requirement matrix, and capstone source
    evidence.

Trace: Input cancellation requirement; expected each invariant maps to
contract, code, deterministic test, production signal, and repair
step.
Show untraced requirement before, design decision, complete links
after,
review finding resolution, and known operational ownership.

Validation: Compile evidence validator; exercise complete requirement
and missing-runbook failure.

Anchors: new heading fragment only.

## 32. Language portability

Purpose: Preserve backend guarantees in another language without
copying Go syntax, abstractions, or unverified framework defaults.

Prerequisites: Entries 1 through 31. Owns C34.

1. The problem it solves: A surface translation can lose cancellation,
   transaction, resource, error, security, test, lifecycle, and
   operating
   guarantees behind a framework's defaults.
2.  The idea in simple terms: Keep the contract and evidence fixed,
   then
   choose the target ecosystem's clearest native mechanism for each
   one.
3.  How it works: Map
   modules/values/errors/cleanup/concurrency/testing,
   make HTTP and middleware behavior explicit, retain real PostgreSQL,
   use native wiring, validate identity, control deterministic inputs,
   produce one artifact, verify ecosystem traps, own a full pipeline,
   and have a native engineer review it.
4.  Step-by-step examples: Port canceled database request from Go
   context
   to .NET token; include Node runtime type mismatch and Java
   transaction
   annotation bypass.
5.  Why it works: Contracts describe observable guarantees
   independently
   of syntax, repository tests expose effective framework behavior,
   and
   native idioms reduce accidental Go-shaped complexity.
6. Go example: A compact mapping interface and conformance test that
   expresses cancellation and stable error outcomes; language-specific
   implementations remain in their native repositories.
7. Work and memory: Count framework startup and middleware, workers or
   event loop, pools, cancellation, exceptions/results, ORM queries,
   runtime validation, package/build bytes, GC/runtime resources, test
   environments, scanners, and migration/operations parity.
8. When to use it: Porting or reviewing Java/Kotlin, .NET, Node.js,
   Python, Rust, or another established backend stack.
9. When not to use it: Do not replace a team's native patterns with Go
   interfaces or accept a framework checkbox as proof.
10. Common mistakes and edge cases: Proxy transaction self-call,
    swallowed interrupt, scoped dependency in singleton, unbounded
    client
    retries, TypeScript/Python types trusted at runtime, blocking
    async
    loop, detached promise, Rust future canceled mid-effect, memory
    safety
    treated as business atomicity, in-memory database replacing
    Postgres.
11. Related material: Link portability source, every active guarantee,
    official ecosystem docs, and equivalent verification pipeline.

Trace: Input client cancellation during database wait; expected target
driver receives cancellation and no later response work runs. Show
active
operation before, native cancellation decision, stopped query after,
and
known released pool/operation state.

Validation: Compile Go conformance model; exercise success and
canceled
outcomes. Each real port must run its own native pipeline and failure
case.

Anchors: new heading fragment only.
