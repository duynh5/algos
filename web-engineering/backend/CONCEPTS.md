# Backend Engineering in Go

This guide owns the active backend teaching material. The lessons use
one evolving Order API to connect transport contracts, data safety,
architecture, and production operations.

## Contents

1. [HTTP API contracts](#1-http-api-contracts)
2. [Data, transactions, and architecture][data-lesson]
3. [Production and distributed systems][production-lesson]

## How to use this guide

Read the lessons in order on a first pass. For each lesson, implement
the example in the Order API lab, exercise its failure path, and keep
the resulting test or operational evidence. Return through the
contents when using the guide as a reference.

## 1. HTTP API Contracts

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

### Related material

- [HTTP source reference][http-reference]
- [Frontend network and security][frontend-network]
- [Data, transactions, and architecture][data-lesson]

[http-reference]: README.md
[frontend-network]:
  ../frontend/CONCEPTS.md#a3-networking-storage-and-browser-security
[data-lesson]: #2-data-transactions-and-architecture
[production-lesson]: #3-production-and-distributed-systems

## 2. Data, Transactions, and Architecture

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

- [Relational data source reference][data-reference]
- [Business architecture source reference][architecture-reference]
- [Production and distributed systems][production-lesson]

[data-reference]: README.md
[architecture-reference]: README.md

## 3. Production and Distributed Systems

### The problem it solves

Production dependencies fail, messages repeat, processes restart, and
network outcomes become uncertain. A correct local function is not
enough to make a service safe to retry, operate, or evolve.

### The idea in simple terms

Assume every remote result can be late, missing, or repeated. Give
work stable identity, persist important transitions, bound every
resource, and expose enough evidence to distinguish failure classes.

### How it works

1. Define service-level indicators for user-visible success and delay.
2. Bound connections, concurrency, queues, attempts, and deadlines.
3. Assign stable identity to retried writes and messages.
4. Persist business state and required outgoing work atomically when
   both must survive a crash.
5. Make consumers idempotent because delivery may repeat.
6. Propagate correlation and trace context without propagating
   secrets.
7. Deploy compatible changes before removing old contracts.
8. Rehearse dependency loss, process termination, and recovery.

### Step-by-step examples

Input: an order commits, then the process crashes before publishing
its `OrderPlaced` event. Expected output: the event is eventually
published without losing the order.

1. **Before:** The transaction contains no order or event record.
   **Decision:** Write the order and an outbox record together.
   **After:** Both commit or both roll back. **Known correct:** A
   durable order always has durable publication work.
2. **Before:** The process crashes before sending. **Decision:** A
   later worker reads the pending outbox record. **After:** It
   publishes the event and records completion. **Known correct:**
   Restart cannot erase required work.
3. **Before:** Publishing succeeds but completion recording fails.
   **Decision:** Retry the same event identity. **After:** The
   consumer ignores or reconciles the duplicate. **Known correct:**
   Repetition does not repeat the business effect.

Failure example: an unbounded retry loop during dependency failure
adds load and prevents recovery. Apply attempt, delay, and total-time
bounds.

### Why it works

Atomic persistence closes the crash gap between business state and
publication intent. Stable identities let producers and consumers
recognize repeated logical work. Bounds guarantee that failure
consumes finite resources, while telemetry connects user impact to
dependencies and releases.

### Go example

```go
package retry

import "time"

type Policy struct {
	MaximumAttempts int
	MaximumDelay    time.Duration
}

func (policy Policy) Delay(attempt int) (time.Duration, bool) {
	if attempt < 1 || attempt >= policy.MaximumAttempts {
		return 0, false
	}
	delay := time.Second << (attempt - 1)
	if delay > policy.MaximumDelay {
		delay = policy.MaximumDelay
	}
	return delay, true
}
```

The caller must also apply jitter, a total deadline, transient-error
classification, and idempotency before retrying writes.

### Work and memory

Each retry repeats remote and local work. Exponential backoff reduces
attempt frequency but increases completion time. Queues, outbox rows,
telemetry, and deduplication records retain bounded state until their
documented expiry. Capacity planning must include worst-case retry
load.

### When to use it

Use durable work for effects that must survive process failure. Use
idempotent consumers for at-least-once delivery. Use staged contract
evolution whenever independently deployed components coexist.

### When not to use it

Do not introduce service extraction without independent ownership,
scaling, isolation, or release value. Do not retry permanent
validation or authorization failures. Do not add a queue where a local
transaction and synchronous result are simpler and safer.

### Common mistakes and edge cases

> **Warning:** "Exactly once" delivery does not guarantee exactly one
> business effect. The consumer must make the effect idempotent.

- Readiness is not liveness.
- High-cardinality metric labels can exhaust monitoring systems.
- A timeout does not prove that remote work stopped.
- Rollback can be unsafe after an incompatible schema change.
- Trace context must not contain tokens or personal data.

### Related material

- [Production service source reference][production-reference]
- [Distributed foundations source reference][distributed-reference]
- [Frontend delivery and leadership][frontend-delivery]

[production-reference]: reference/advanced/11-production-service.md
[distributed-reference]:
  reference/advanced/12-distributed-foundations.md
[frontend-delivery]:
  ../frontend/CONCEPTS.md#delivery-and-leadership
