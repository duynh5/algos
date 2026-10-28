# Frontend Engineering

This guide owns the active frontend teaching material. The lessons
connect browser foundations, product application engineering, and
senior delivery and leadership responsibilities.

## Contents

- [How to use this guide](#how-to-use-this-guide)
- [A. Foundations](#a-foundations)
  - [A.1 Semantic interfaces and
    accessibility](#a1-semantic-interfaces-and-accessibility)
  - [A.2 JavaScript, TypeScript, and the browser
    runtime](#a2-javascript-typescript-and-the-browser-runtime)
  - [A.3 Networking, storage, and browser
    security](#a3-networking-storage-and-browser-security)
- [B. Product application
  engineering](#b-product-application-engineering)
  - [B.1 React components, state, and
    effects](#b1-react-components-state-and-effects)
  - [B.2 Server data, routing, rendering, and
    forms](#b2-server-data-routing-rendering-and-forms)
  - [B.3 Testing, performance, and frontend
    observability](#b3-testing-performance-and-frontend-observability)
- [C. Senior engineering](#c-senior-engineering)
  - [C.1 Frontend architecture and design
    systems](#c1-frontend-architecture-and-design-systems)
  - [C.2 Delivery, migrations, and technical
    leadership](#c2-delivery-migrations-and-technical-leadership)

## How to use this guide

Read the lessons in order on a first pass. For each lesson, reproduce
its trace, implement its example, trigger its failure case, pass its
checks, and explain its tradeoffs aloud. Return through the contents
when using the guide as a reference.

## A. Foundations

### A.1 Semantic Interfaces and Accessibility

#### The problem it solves

A web interface must communicate structure, accept input, adapt to
different screens, and remain usable with a keyboard, screen reader,
zoom, slow device, or missing enhancement code.

#### How it works

Think of signs and doors in a public building. A sign says what a room
is for, and a door works without special instructions. **Semantic
HTML** gives a page that same clear structure, including built-in
keyboard and form behavior. CSS changes its look, while accessibility
keeps the workflow usable in different ways.

#### Step-by-step examples

Input: an email value `nolan@example.com`. Expected output: the form
submits once and announces `Subscription saved`.

1. **Before:** Focus is on the labeled email input. The native form
   owns Enter-key submission. **Decision:** Keep a real `form`,
   `label`, `input`, and submit `button` instead of clickable generic
   elements. **After:** Keyboard and pointer users invoke the same
   action. **Known correct:** The control names and submit behavior
   exist before JavaScript runs.
2. **Before:** Submission is pending. **Decision:** Disable duplicate
   submission and set visible status text to `Saving`. **After:** One
   request is active and assistive technology can hear the status.
   **Known correct:** Repeated activation cannot create another
   request.
3. **Before:** The request succeeds. **Decision:** Replace the status
   text with `Subscription saved`. **After:** The result is visible
   and announced. **Known correct:** The workflow has one unambiguous
   terminal state.

Boundary example: at 200 percent zoom, the label, input, and button
must reflow without horizontal page scrolling or clipped controls.

#### Why it works

Native elements already implement names, roles, focus behavior, and
input semantics. CSS changes their arrangement without changing their
meaning. The workflow exposes every state as text and prevents a
second submission while the first owns the request, so all users
receive the same state transitions.

#### TypeScript example

This example runs in a page containing `#subscription` and `#status`.

```typescript
// Resolve the semantic controls once and fail before registering a
// partially functional interaction.
const form = document.querySelector<HTMLFormElement>("#subscription");
const statusElement = document.querySelector<HTMLElement>("#status");

if (!form || !statusElement) {
  throw new Error("subscription form is incomplete");
}

type SubmitButton = HTMLButtonElement;

// One submit handler owns validation, visible status, and the
// disabled state for the lifetime of each request.
form.addEventListener("submit", async (event) => {
  event.preventDefault();
  const submitSelector = 'button[type="submit"]';
  const submit = form.querySelector<SubmitButton>(submitSelector);
  if (!submit || submit.disabled) return;

  // Disable repeated intent before starting asynchronous work, then
  // restore the control on every completion path.
  submit.disabled = true;
  statusElement.textContent = "Saving";
  try {
    await Promise.resolve();
    statusElement.textContent = "Subscription saved";
  } catch {
    statusElement.textContent = "Could not save the subscription";
  } finally {
    submit.disabled = false;
  }
});
```

#### Work and memory

The browser visits each relevant element while creating accessibility
and layout information. Deep, redundant markup increases that work and
the memory retained by the document. Images, fonts, and layout shifts
usually matter more than the element count alone. Measure real pages
instead of assigning a universal node limit.

#### When to use it

Use semantic structure and progressive enhancement for every user
interface, especially navigation, forms, dialogs, data entry, and
content that must work across devices or assistive technology.

#### When not to use it

Do not force a native element when its contract is different from the
required interaction. A spreadsheet-like grid may need a carefully
implemented composite widget. Prefer an established accessible pattern
over inventing keyboard behavior.

#### Common mistakes and edge cases

> **Warning:** A `div` with a click handler is not a button. It lacks
> keyboard activation, button semantics, focus behavior, and disabled
> behavior unless all are rebuilt correctly.

- Placeholder text is not a persistent label.
- Positive `tabindex` values create a second, fragile focus order.
- Removing focus outlines hides the user's current position.
- Fixed heights clip translated, zoomed, or user-sized text.
- Color-only validation fails when color cannot be perceived.
- Moving focus on every render disorients keyboard users.

#### Related material

- [Frontend curriculum](README.md)
- [Testing, performance, and observability][quality]
- [Frontend architecture and design systems][architecture]

[quality]: #b3-testing-performance-and-frontend-observability
[architecture]: #c1-frontend-architecture-and-design-systems

### A.2 JavaScript, TypeScript, and the Browser Runtime

#### The problem it solves

Frontend code coordinates user events, asynchronous work, shared data,
and rendering on a mostly single main thread. Type errors, stale work,
and long tasks can make an interface incorrect or unresponsive.

#### How it works

Picture one cashier serving a customer while other work waits in
lines. JavaScript runs one current job; the **event loop** chooses the
next queued job. TypeScript checks the plan before the program runs,
but data arriving at the counter still needs a real check.

#### Step-by-step examples

Input: searches for `ca`, then `cat`. The `cat` response returns
first. Expected output: only results for `cat` appear.

1. **Before:** No request is active. **Decision:** Start request 1 for
   `ca` and retain its controller. **After:** Request 1 may update the
   view unless canceled. **Known correct:** The current query is `ca`.
2. **Before:** Request 1 is pending and the user enters `cat`.
   **Decision:** Abort request 1 before starting request 2. **After:**
   Only request 2 owns permission to publish results.
   **Known correct:** Any result for `ca` is obsolete.
3. **Before:** Request 2 returns valid data. **Decision:** Validate
   its shape and publish it. **After:** The view shows results for
   `cat`. **Known correct:** Visible data matches the latest accepted
   query.

Failure example: a successful HTTP response containing the wrong JSON
shape is a data-contract failure, not a TypeScript compile error.

#### Why it works

At most one controller represents the accepted request. Starting newer
work revokes older work before it can publish. Runtime validation
turns untrusted JSON into a checked value. Type narrowing then makes
every rendering branch explicit and prevents impossible field access.

#### TypeScript example

```typescript
type SearchState =
  | { kind: "idle" }
  | { kind: "loading"; query: string }
  | { kind: "ready"; query: string; items: string[] }
  | { kind: "failed"; query: string; message: string };

let activeController: AbortController | undefined;

async function search(query: string): Promise<SearchState> {
  // Revoke the preceding request before a newer query takes
  // ownership.
  activeController?.abort();
  const controller = new AbortController();
  activeController = controller;

  // Treat transport status and decoded JSON as untrusted boundary
  // data before constructing the ready state.
  try {
    const encodedQuery = encodeURIComponent(query);
    const searchURL = `/search?q=${encodedQuery}`;
    const response = await fetch(searchURL, {
      signal: controller.signal,
    });
    if (!response.ok) throw new Error(`HTTP ${response.status}`);
    const value: unknown = await response.json();
    if (!Array.isArray(value)) throw new Error("invalid result");
    const items = value.filter(
      (item): item is string => typeof item === "string",
    );
    return { kind: "ready", query, items };
  } catch (error) {
    // Convert every failure shape into the explicit failed state.
    let message = "unknown";
    if (error instanceof Error) {
      message = error.message;
    }
    return { kind: "failed", query, message };
  }
}
```

#### Work and memory

If the response contains `n` items, validation visits each item once,
so the work grows as `O(n)`. The filtered array can retain `O(n)` new
references. Network latency and main-thread blocking are separate from
that CPU cost and should be measured independently.

#### When to use it

Use discriminated unions for mutually exclusive UI states. Use runtime
validation at network, storage, message, and user-input boundaries.
Use cancellation whenever new intent makes pending work obsolete.

#### When not to use it

Do not add generic abstractions for a single simple value. Do not use
`AbortController` as a substitute for server-side idempotency:
aborting a browser request does not prove that the server stopped
processing it.

#### Common mistakes and edge cases

> **Warning:** `as SearchResult` does not validate JSON. A malformed
> response still reaches runtime and can break a supposedly typed
> path.

- Long microtask chains can delay painting.
- A closure reads the bindings it captured, which may represent stale
  intent by the time asynchronous work finishes.
- `any` disables evidence; prefer `unknown` at untrusted boundaries.
- Aborted requests need a deliberate user-facing policy.
- Unicode string length counts UTF-16 code units, not visible symbols.

#### Related material

- [Network, storage, and browser security][network]
- [React components, state, and effects][react]
- [Algorithm and Go concept guide](../dsa/CONCEPTS.md)

[react]: #b1-react-components-state-and-effects
[network]: #a3-networking-storage-and-browser-security

### A.3 Networking, Storage, and Browser Security

#### The problem it solves

A browser application exchanges untrusted data across network and
origin boundaries. Incorrect caching, credential handling, or output
encoding can expose private data or make state inconsistent.

#### How it works

Sending data from a page to an API is like passing a package through
several hands. At each handoff, say who may send it, what a valid
package looks like, whether an old copy may be reused, how long it may
be kept, and what happens if it is late, changed, or lost. Each
handoff is a **trust boundary**.

#### Step-by-step examples

Input: `POST /orders` with an idempotency key and session cookie.
Expected output: retries create one order and return the same result.

1. **Before:** The user submits once and the outcome is uncertain.
   **Decision:** Send a unique operation key with the authenticated
   request. **After:** The server can associate later retries with the
   operation. **Known correct:** One logical action has one stable
   identity.
2. **Before:** The connection fails after the server may have
   committed. **Decision:** Retry with the same key, not a new key.
   **After:** The server returns the recorded result or finishes once.
   **Known correct:** Retry does not authorize another logical order.
3. **Before:** The response arrives. **Decision:** Validate it before
   writing application state. **After:** Only contract-valid order
   data enters the client cache. **Known correct:** UI state and the
   accepted server result agree.

Adversarial example: rendering an order note as HTML can execute an
attacker's script. Render it as text unless sanitized HTML is
required.

#### Why it works

The origin policy restricts reads, the server verifies identity and
authorization, CSRF controls prove intentional cookie-authenticated
writes, and context-aware output prevents data from becoming code.
Stable operation identity makes an uncertain retry converge on one
server result.

#### TypeScript example

```typescript
type Order = { id: string; status: "placed" | "rejected" };
type OrderPromise = Promise<Order>;

function isOrder(value: unknown): value is Order {
  // Validate the complete wire shape before application code uses it.
  if (typeof value !== "object" || value === null) return false;
  const record = value as Record<string, unknown>;
  return (
    typeof record.id === "string" &&
    (record.status === "placed" || record.status === "rejected")
  );
}

async function placeOrder(body: unknown, key: string): OrderPromise {
  // Send credentials and replay identity explicitly at the HTTP
  // boundary.
  const response = await fetch("/v1/orders", {
    method: "POST",
    credentials: "same-origin",
    headers: {
      "Content-Type": "application/json",
      "Idempotency-Key": key,
    },
    body: JSON.stringify(body),
  });
  // Accept the response only when transport and runtime contracts
  // both succeed.
  const value: unknown = await response.json();
  if (!response.ok || !isOrder(value)) {
    throw new Error("order request failed");
  }
  return value;
}
```

#### Work and memory

Serialization and validation grow with the transferred payload size.
Browser caches and storage consume persistent or memory capacity until
eviction or deletion. More consequential costs are round trips,
connection setup, server work, cache misses, and leaked sensitive
data.

#### When to use it

Apply this boundary model to every API, embedded third party, worker,
cross-window message, persisted value, upload, and authenticated
action.

#### When not to use it

Do not store authoritative business state only in the browser. Do not
use local storage for secrets that must be protected from script
execution. Do not retry non-idempotent writes without a server
contract.

#### Common mistakes and edge cases

> **Warning:** A successful CORS preflight does not authorize a user.
> The API must still authenticate and authorize every protected
> action.

- `HttpOnly` protects a cookie from script reads, not from CSRF.
- A `200` response can still contain an invalid application payload.
- Cache keys that omit user or filter scope can expose wrong data.
- Logout must address server sessions and client-held private state.
- Aborting `fetch` does not prove that a write was rolled back.

#### Related material

- [TypeScript and browser
  runtime](#a2-javascript-typescript-and-the-browser-runtime)
- [Server data, routing, rendering, and forms][data]
- [Backend identity and API security][backend-security]

[data]: #b2-server-data-routing-rendering-and-forms
[backend-security]:
  ../backend/CONCEPTS.md#9-identity-authorization-and-api-security

## B. Product Application Engineering

### B.1 React Components, State, and Effects

#### The problem it solves

Interactive screens must turn changing data into predictable UI while
keeping reusable parts understandable. Poor state ownership and
effects create stale values, duplicate requests, and render loops.

#### How it works

Think of a component as a display board built from today's inputs and
settings. **Rendering** redraws the board when those facts change.
Values derived from them stay in rendering, and user actions stay in
event handlers. An **effect** synchronizes the committed result with a
system outside React, such as a browser API.

#### Step-by-step examples

Input: select product `A`, then product `B` before `A` finishes.
Expected output: only details for `B` remain visible.

1. **Before:** The committed selection is `A` and effect A owns its
   request. **Decision:** Selection changes to `B`; render B's loading
   state. **After:** React commits the B state. **Known correct:** The
   visible selection identity is `B`.
2. **Before:** Effect A still has pending work. **Decision:** Cleanup
   aborts A before effect B starts. **After:** B owns the only
   accepted request. **Known correct:** A cannot intentionally publish
   into B's state.
3. **Before:** Request B succeeds. **Decision:** Store B's details.
   **After:** React renders the ready state for B. **Known correct:**
   Selection and details share the same identity.

Boundary example: removing and reinserting a list item with its stable
ID preserves the correct component state. Using its position as the
key may attach old state to a different item.

#### Why it works

Each render is a calculation from one snapshot. Identity determines
which state React preserves. Cleanup releases synchronization owned by
the preceding committed inputs, so the next effect can establish one
replacement. Derived values stay in rendering and cannot drift from
their source state.

#### TypeScript example

This component assumes `loadProduct` respects an abort signal.

```tsx
type Product = { id: string; name: string };

function ProductDetails({ id }: { id: string }) {
  const [product, setProduct] = React.useState<Product | null>(null);
  const [error, setError] = React.useState<string | null>(null);

  // Each committed product identity owns one abortable request.
  React.useEffect(() => {
    const controller = new AbortController();
    setProduct(null);
    setError(null);

    // Ignore rejection from work that cleanup has already revoked.
    loadProduct(id, controller.signal)
      .then(setProduct)
      .catch((reason: unknown) => {
        if (!controller.signal.aborted) {
          let message = "failed";
          if (reason instanceof Error) {
            message = reason.message;
          }
          setError(message);
        }
      });

    // Cleanup prevents an obsolete request from publishing afterward.
    return () => controller.abort();
  }, [id]);

  // Render every observable request state from the current snapshot.
  if (error) return <p role="alert">{error}</p>;
  if (!product) return <p>Loading</p>;
  return <h2>{product.name}</h2>;
}
```

#### Work and memory

Rendering visits the components needed for the update. Reconciliation
work grows with the compared element subtrees, while browser layout
and paint depend on the committed DOM and styles. Memoization retains
values and comparison work; it is useful only when measurement shows
that avoided work costs more.

#### When to use it

Use local state for interaction owned by one subtree. Lift state when
siblings require one source of truth. Use effects for synchronization
with external systems and event handlers for user-triggered actions.

#### When not to use it

Do not copy props into state merely to derive a display value. Do not
use an effect to calculate values that rendering can calculate. Do not
put server cache behavior into scattered component effects when a
shared server-data owner is needed.

#### Common mistakes and edge cases

> **Warning:** Omitting an effect dependency does not freeze a value.
> It makes the effect retain an older render's value and can publish
> stale work.

- Array indexes are unsafe keys when items reorder or are inserted.
- Mutation can preserve a reference and hide a meaningful state
  change.
- An effect that sets one of its own dependencies can loop.
- Strict development checks expose unsafe setup and cleanup logic.
- Context updates every consumer that observes the changed value.

#### Related material

- [TypeScript and browser runtime][runtime]
- [Server data, routing, rendering, and forms][data]
- [Testing, performance, and observability][quality]

[runtime]: #a2-javascript-typescript-and-the-browser-runtime

### B.2 Server Data, Routing, Rendering, and Forms

#### The problem it solves

A product interface must coordinate URLs, server data, user edits,
caching, navigation, and several failure states without presenting old
or contradictory information.

#### How it works

Keep each fact in the place responsible for it, like keeping an
address on an envelope and an unfinished note on its own notepad. The
URL owns shareable location, the server owns saved records, a
server-data cache owns synchronized copies, a form owns its draft, and
a component owns only temporary local interaction.

#### Step-by-step examples

Input: route `/products?category=books&page=2`. Expected output: page
2 of books, with a stable URL that can be refreshed and shared.

1. **Before:** The URL contains untrusted strings. **Decision:**
   Validate category and parse a positive page number. **After:**
   Application input is `{ category: "books", page: 2 }`.
   **Known correct:** Invalid page values cannot enter the cache key.
2. **Before:** No matching cache entry is fresh. **Decision:** Fetch
   with key `["products", "books", 2]`. **After:** The route renders
   its pending state. **Known correct:** Other categories and pages
   cannot share the key.
3. **Before:** Valid data returns. **Decision:** Store it under the
   full key and render it. **After:** The ready view and URL describe
   the same selection. **Known correct:** Refresh and back navigation
   can reconstruct it.

Failure example: if page 2 is empty because an item was removed, show
an empty page or move through an explicit navigation decision. Do not
silently display page 1 under a page 2 URL.

#### Why it works

Every state has one authoritative owner. Complete cache keys prevent
unrelated results from colliding. Validation happens before lookup.
Navigation changes invalidate obsolete intent, while drafts remain
separate from confirmed data until mutation policy reconciles them.

#### TypeScript example

```typescript
type ProductQuery = { category: string; page: number };

function parseProductQuery(url: URL): ProductQuery {
  // Parse and validate URL-owned state before building a cache key.
  const category = url.searchParams.get("category")?.trim();
  const rawPage = Number(url.searchParams.get("page") ?? "1");
  if (!category) throw new Error("category is required");
  if (!Number.isInteger(rawPage) || rawPage < 1) {
    throw new Error("page must be a positive integer");
  }
  return { category, page: rawPage };
}

function productKey(query: ProductQuery): readonly unknown[] {
  // Include every input that can change the server result.
  return ["products", query.category, query.page] as const;
}

// Exercise the ordinary URL-to-key path with a concrete assertion.
const query = parseProductQuery(
  new URL("https://example.test/products?category=books&page=2"),
);
console.assert(productKey(query).join(":") === "products:books:2");
```

#### Work and memory

Parsing search parameters grows with URL size. Rendering a page of `n`
items requires work proportional to those items plus their component
subtrees. A cache retains data for each live key until eviction.
Server rendering adds server CPU and transfer work but may improve
initial content delivery; measure the complete path.

#### When to use it

Use URL state for shareable or navigable selection. Use a server-data
cache when multiple components need freshness, deduplication, retries,
or invalidation. Use form state for drafts and validation feedback.

#### When not to use it

Do not place ephemeral hover or disclosure state in the URL. Do not
copy every server record into a global client store. Do not use
optimistic updates when rollback is unclear or the server result is
not predictable.

#### Common mistakes and edge cases

> **Warning:** A cache key missing `userId`, locale, filter, or page
> can show one request's data in another request's view.

- Hydration fails when server and first client output differ.
- A form reset can erase a user's draft during background refresh.
- Retrying a write with a new operation key can duplicate the action.
- Redirect loops arise when loaders and guards disagree about state.
- File inputs and very large forms need special memory consideration.

#### Related material

- [Network, storage, and browser security][network]
- [React components, state, and
  effects](#b1-react-components-state-and-effects)
- [Backend HTTP contracts][http]

[http]: ../backend/README.md

### B.3 Testing, Performance, and Frontend Observability

#### The problem it solves

A frontend can look correct locally while failing through timing,
accessibility, browser differences, slow networks, or production-only
data. Teams need evidence that behavior works and signals that explain
failure after release.

#### How it works

Test the part a user actually depends on. A fast button is not useful
if the page freezes after it is pressed. Check behavior at the
boundary that can fail, measure the user's wait, and connect errors to
the release that introduced them.

#### Step-by-step examples

Input: a checkout submit request fails once, then succeeds. Expected
output: the user sees an error, retries once, and receives one order.

1. **Before:** The valid form is ready. **Decision:** Activate
   `Place order` and assert a pending state. **After:** Duplicate
   activation is unavailable. **Known correct:** One client operation
   owns the first request.
2. **Before:** The server returns a controlled transient failure.
   **Decision:** Assert an accessible error and enabled retry action.
   **After:** The form data remains intact. **Known correct:** Failure
   did not pretend to be success or lose input.
3. **Before:** The user retries with the same operation identity.
   **Decision:** Return the recorded successful order. **After:** The
   confirmation shows one order ID. **Known correct:** The tested
   workflow converges after uncertainty.

Boundary example: run the workflow with keyboard input, a narrow
viewport, delayed responses, and reduced-motion preference.

#### Why it works

Each test observes a contract rather than an implementation detail.
The failure is deliberately produced, so the recovery claim has
evidence. Production signals include enough stable dimensions to group
similar failures without recording private user data.

#### TypeScript example

```typescript
type RetryState = "ready" | "pending" | "failed" | "complete";

function nextRetryState(
  state: RetryState,
  event: "submit" | "reject" | "resolve",
): RetryState {
  // Only idle terminal states may begin a new request.
  const canSubmit = state === "ready" || state === "failed";
  if (canSubmit && event === "submit") {
    return "pending";
  }
  // A pending request owns the only valid completion transitions.
  if (state === "pending" && event === "reject") return "failed";
  if (state === "pending" && event === "resolve") return "complete";
  throw new Error(`invalid transition: ${state} -> ${event}`);
}

// Exercise both the start and failure transitions.
console.assert(nextRetryState("ready", "submit") === "pending");
console.assert(nextRetryState("pending", "reject") === "failed");
```

#### Work and memory

Test cost includes environment startup, fixtures, browser processes,
network simulation, and retained artifacts. Performance cost includes
download bytes, parse and execution time, main-thread blocking,
rendering, memory, and interaction latency. Observability adds bounded
events and network traffic; sampling and retention must have limits.

#### When to use it

Use the smallest test that can prove the owned risk. Use browser tests
for browser contracts and critical workflows. Use real-user metrics
for claims about actual performance and release health.

#### When not to use it

Do not replace behavior tests with snapshots. Do not make every unit
path an end-to-end test. Do not use synthetic lab scores alone as
proof of user experience. Do not record secrets or personal data for
easier debugging.

#### Common mistakes and edge cases

> **Warning:** Waiting for a fixed time makes an asynchronous test
> both slow and unreliable. Wait for the user-observable state
> instead.

- Mocking the code under test can make a test prove its own setup.
- Average latency hides slow cohorts and regressions.
- Unhandled promise rejections lose route and operation context.
- Source maps require controlled access because they reveal source.
- Performance budgets must fail builds or alerts to influence
  behavior.

#### Related material

- [Semantic interfaces and accessibility][interfaces]
- [Frontend architecture and design systems][architecture]
- [Backend testing and operations][backend-quality]

[interfaces]: #a1-semantic-interfaces-and-accessibility
[backend-quality]: ../backend/README.md

## C. Senior Engineering

### C.1 Frontend Architecture and Design Systems

#### The problem it solves

As products and teams grow, unclear boundaries make changes risky.
Components duplicate behavior, state leaks across features, visual
rules drift, and every release requires coordination.

#### How it works

Organize a frontend like a well-run workshop: each area owns a clear
job, shared tools have one maintained home, and dependencies point in
an obvious direction. A **design system** is that shared tool cabinet,
with users, compatibility promises, and releases of its own.

#### Step-by-step examples

Input: checkout and account pages each implement a different modal.
Expected output: one accessible dialog primitive and separate
workflows.

1. **Before:** Both features own focus trapping and keyboard behavior.
   **Decision:** Extract the accessible dialog mechanism, not the
   business content. **After:** One primitive owns labeling, focus,
   escape, and restore. **Known correct:** Both workflows share the
   interaction contract.
2. **Before:** Checkout needs payment-specific confirmation.
   **Decision:** Keep confirmation rules in checkout and compose the
   dialog primitive. **After:** Domain behavior remains with its
   product owner. **Known correct:** The shared package has no payment
   dependency.
3. **Before:** The primitive changes its close behavior. **Decision:**
   Test consumers, document compatibility, and provide a migration
   path. **After:** Teams can adopt the release deliberately.
   **Known correct:** Shared change does not require hidden
   coordination.

Adversarial example: extracting every repeated class name creates a
shared package whose API changes for unrelated feature requests.

#### Why it works

Stable boundaries group code that changes for the same reason. Domain
dependencies point inward to feature ownership. Shared primitives own
only behavior common to all consumers, so their interface changes less
often and can be tested against one accessibility contract.

#### TypeScript example

```typescript
type DialogRequest = {
  titleId: string;
  descriptionId?: string;
  close(): void;
};

type CheckoutConfirmation = {
  orderId: string;
  totalCents: number;
};

function confirmationTitle(value: CheckoutConfirmation): string {
  // Domain data determines presentation without entering the dialog
  // interaction contract.
  const amount = (value.totalCents / 100).toFixed(2);
  return `Order ${value.orderId}: ${amount}`;
}

function closeDialog(request: DialogRequest): void {
  // The shared interaction closes through its narrow callback.
  request.close();
}
```

The dialog contract contains interaction concerns. The checkout value
and formatter remain in the checkout feature.

#### Work and memory

Architecture adds build graph, package, documentation, and
coordination cost. Shared runtime packages can increase initial
download size or create duplicated versions. Micro-frontends add
framework, network, integration, and observability overhead. Measure
change lead time and runtime cost instead of assuming more boundaries
are more scalable.

#### When to use it

Create a boundary when code has a stable owner, independent reason to
change, testable contract, and more than one meaningful consumer.
Build a design system when product consistency and accessibility need
shared investment across teams.

#### When not to use it

Do not create a package for speculative reuse. Do not use global state
to avoid designing feature boundaries. Do not adopt micro-frontends
for a single team or only to solve repository organization.

#### Common mistakes and edge cases

> **Warning:** A shared component that accepts every feature rule as a
> flag is a hidden collection of domain workflows, not a primitive.

- Barrel exports can hide dependency cycles.
- Tokens named after raw colors cannot express product intent.
- Breaking visual changes can still be breaking API changes.
- Shared singletons complicate tests and independent delivery.
- Design-system success requires adoption support and migration tools.

#### Related material

- [Semantic interfaces and accessibility][interfaces]
- [React components, state, and effects][react]
- [Delivery, migrations, and leadership][leadership]

[leadership]: #c2-delivery-migrations-and-technical-leadership

<a id="delivery-and-leadership"></a>

### C.2 Delivery, Migrations, and Technical Leadership

#### The problem it solves

Senior engineers must change production systems without stopping
delivery, losing user trust, or making success depend on one person's
memory. Technical correctness alone does not coordinate teams or
control release risk.

#### How it works

Replace a busy bridge one lane at a time. Keep traffic visible, make
each change small enough to reverse, and let old and new routes
coexist safely. A **migration** ends only when evidence shows the old
route can be removed. Technical leadership turns remaining uncertainty
into a decision, evidence, an owner, and follow-up.

#### Step-by-step examples

Input: migrate checkout from a legacy store to a query cache. Expected
output: no user-visible regression and eventual removal of the old
path.

1. **Before:** Only the legacy path serves production. **Decision:**
   Add outcome, latency, and duplicate-request measures. **After:**
   Baseline behavior is visible. **Known correct:** Future comparisons
   have a reference cohort.
2. **Before:** Both implementations can read the same API contract.
   **Decision:** Enable the new path for internal users, then one
   percent. **After:** A bounded cohort exercises the migration.
   **Known correct:** Most users remain on the known path.
3. **Before:** Measures stay within thresholds through wider rollout.
   **Decision:** Reach full use, retain rollback for one defined
   period, then remove old code and flags. **After:** One
   implementation owns checkout data. **Known correct:** Compatibility
   complexity has a verified end.

Failure example: if duplicate order attempts rise above the threshold,
stop rollout and investigate operation identity before expanding.

#### Why it works

Compatibility keeps both paths valid during transition. Bounded
exposure limits the affected population. Measurements connect a
release to user and system outcomes. Explicit thresholds make response
a prior decision rather than an improvised debate during failure.

#### TypeScript example

```typescript
type RolloutDecision = "continue" | "stop" | "rollback";

type ReleaseSignals = {
  errorRate: number;
  duplicateOrderRate: number;
  canRollback: boolean;
};

function decideRollout(signals: ReleaseSignals): RolloutDecision {
  // Continue only while every user-impact signal stays within policy.
  const errorRateUnsafe = signals.errorRate > 0.01;
  const duplicateUnsafe = signals.duplicateOrderRate > 0;
  const unsafe = errorRateUnsafe || duplicateUnsafe;
  if (!unsafe) return "continue";
  // Prefer reversal when supported; otherwise halt further exposure.
  return signals.canRollback ? "rollback" : "stop";
}

// Exercise the rollback branch with a breached error-rate threshold.
console.assert(
  decideRollout({
    errorRate: 0.02,
    duplicateOrderRate: 0,
    canRollback: true,
  }) === "rollback",
);
```

#### Work and memory

Safe delivery temporarily duplicates code paths, tests, telemetry, and
operational attention. Feature controls retain configuration state.
Parallel implementations consume engineering capacity. Bound every
migration with an owner, removal condition, and latest review date.

#### When to use it

Use staged migration for architecture, state, rendering, API, design
system, authentication, and build changes whose failure affects users
or multiple teams. Use written decisions when tradeoffs outlive a
chat.

#### When not to use it

Do not add rollout machinery to a trivial, easily reversible text
change. Do not keep permanent dual paths because deletion feels risky.
Do not use consensus as a substitute for a named accountable decision.

#### Common mistakes and edge cases

> **Warning:** A feature flag without an owner and removal condition
> is permanent hidden architecture.

- Measuring only exceptions misses user abandonment and slowness.
- Rollback may be unsafe after an incompatible data or API change.
- Global rollout hides regional, browser, and device regressions.
- Incident reviews that stop at human error miss system corrections.
- Mentoring without delegated ownership creates dependence, not
  growth.

#### Related material

- [Frontend architecture and design systems][architecture]
- [Testing, performance, and observability][quality]
- [Backend delivery and production reference][backend-delivery]
- [Senior frontend job preparation][job-plan]

[backend-delivery]: ../backend/README.md
[job-plan]: ../SENIOR_FRONTEND_JOB_PREP_PLAN.md
