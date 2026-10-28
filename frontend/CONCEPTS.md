# Frontend Engineering

This guide owns the active frontend teaching material. The lessons
connect browser foundations, product application engineering, and
senior delivery and leadership responsibilities.

The book is self-contained for its core explanations. Source links are
included so a reader can verify standards, browser support, and
release-sensitive details that can change after publication.

## Contents

- [How to use this guide](#how-to-use-this-guide)
- [Current technology snapshot](#current-technology-snapshot)
- [A. Foundations](#a-foundations)
  - [A.1 Semantic interfaces and
    accessibility](#a1-semantic-interfaces-and-accessibility)
  - [A.2 JavaScript language and browser
    runtime](#a2-javascript-language-and-browser-runtime)
  - [A.3 TypeScript type modeling and runtime
    boundaries](#a3-typescript-type-modeling-and-runtime-boundaries)
  - [A.4 HTTP, caching, and network
    delivery](#a4-http-caching-and-network-delivery)
  - [A.5 Browser storage, origins, and
    security](#a5-browser-storage-origins-and-security)
  - [A.6 CSS, layout, responsive design, and
    rendering](#a6-css-layout-responsive-design-and-rendering)
  - [A.7 Modules, toolchains, and production
    builds](#a7-modules-toolchains-and-production-builds)
- [B. Product application
  engineering](#b-product-application-engineering)
  - [B.1 React components, state, and
    effects](#b1-react-components-state-and-effects)
  - [B.2 Server data, routing, caching, and
    forms](#b2-server-data-routing-caching-and-forms)
  - [B.3 Testing, performance, and frontend
    observability](#b3-testing-performance-and-frontend-observability)
  - [B.4 Rendering architectures, hydration, and Server
    Components](
    #b4-rendering-architectures-hydration-and-server-components)
- [C. Senior engineering](#c-senior-engineering)
  - [C.1 Frontend architecture and design
    systems](#c1-frontend-architecture-and-design-systems)
  - [C.2 Delivery, migrations, and technical
    leadership](#c2-delivery-migrations-and-technical-leadership)
- [Sources and maintenance](#sources-and-maintenance)

## How to use this guide

Read the lessons in order on a first pass. For each lesson, reproduce
its trace, implement its example, trigger its failure case, pass its
checks, and explain its tradeoffs aloud. Return through the contents
when using the guide as a reference.

## Current technology snapshot

This snapshot was verified on October 28, 2026. React 19.3 is the
current stable React line, and TypeScript 7 is the current stable
compiler line. Their durable programming models matter more than
release trivia, so lessons label version-specific behavior instead of
building the explanation around a version number.

Web-platform availability is not a single yes-or-no fact. Check the
project's actual browser and embedded-webview targets. Baseline is a
useful cross-browser signal, but it does not cover every browser or
webview a product may support. Progressive enhancement keeps the core
workflow available when an optional feature is absent.

The final [sources and maintenance](#sources-and-maintenance) section
owns release and standards links used to refresh this snapshot.

## A. Foundations

### A.1 Semantic Interfaces and Accessibility

#### The problem it solves

A web interface must communicate structure, accept input, and expose
every state without assuming sight, hearing, precise pointing, one
language, one text direction, one screen size, or one input device.
The desired result is the same understandable and complete workflow
through semantic markup, keyboard or pointer input, assistive
technology, zoom, translated text, and missing optional enhancement.

#### How it works

Start with the browser's existing contracts. **Semantic HTML** chooses
an element by what it means and does, not by its default appearance.
A button performs an action, a link navigates, a heading names a
section, and a label names a form control. CSS may change presentation
without replacing those contracts.

The browser turns the document into several connected views. The
Document Object Model (DOM) represents nodes and attributes. The
render tree represents boxes that need visual layout. An
**accessibility tree** exposes relevant roles, names, descriptions,
states, values, and relationships to platform accessibility APIs.
Assistive technologies such as screen readers consume that platform
information. A node can exist in the DOM while being absent from
visual or accessibility output, so hidden-state rules must be
intentional.

Build the document outline before individual controls:

- Give the page a useful `title`, declare its natural language with
  `lang`, and set text direction with `dir` when needed.
- Use landmarks such as `header`, `nav`, `main`, and `footer` for
  regions that actually have those purposes. Give repeated landmarks
  distinct names.
- Use headings to describe nested content, not to obtain large text.
  Do not choose a heading level from its visual size.
- Keep source order meaningful because reading order and ordinary tab
  order follow it. Visual reordering must not create another story.
- Provide a visible-on-focus skip link when repeated navigation would
  otherwise require many key presses before the main content.

Every interactive control needs an **accessible name**, the short text
that identifies its purpose. Prefer a visible label connected through
native markup. An accessible description adds instructions or context
without replacing the name. `aria-labelledby` and
`aria-describedby` can reference existing text. `aria-label` supplies
an invisible name and can override visible text, so use it only when a
maintainable visible label is not available.

Accessible Rich Internet Applications (ARIA) attributes can add or
repair semantics that HTML does not provide. ARIA changes what the
accessibility API reports; it does not add keyboard behavior, focus,
validation, layout, or event handling. Prefer a native element with
the required contract. When custom behavior is necessary, keep its
role, state, value, keyboard behavior, pointer behavior, and visual
state synchronized.

Keyboard focus is the one element that currently receives keyboard
input. Selection is a separate, persistent widget state. Keep focus
visible and visually distinct from selection. Ordinary focus follows
interactive elements in DOM order:

- A native focusable element or `tabindex="0"` participates at its DOM
  position.
- `tabindex="-1"` permits programmatic focus without adding a tab
  stop.
- A positive `tabindex` creates a second ordering system and should be
  avoided.

`Tab` normally moves between components. A **composite widget** such
as tabs, a toolbar, menu, tree, or grid usually contributes one tab
stop; arrow keys move inside it according to its established pattern.
A **roving tabindex** keeps `tabindex="0"` on the active child and
`-1` on its peers. `aria-activedescendant` instead keeps DOM focus on
a container while naming an active child. Follow an established ARIA
Authoring Practices Guide pattern rather than inventing key bindings.

Move focus only when the workflow changes context or the focused node
disappears. Opening a modal dialog moves focus into it, constrains
interaction to it, and closing restores focus to a sensible trigger.
After deleting a focused row, move focus to the next useful control.
A client-side route change may need focus at the new main heading or
content start. Do not reset focus on routine renders or announce a
change without also making its destination operable.

Native forms connect several contracts:

- A `label` names its control. `fieldset` and `legend` name a related
  group. Instructions remain visible rather than living only in a
  placeholder.
- `type`, `required`, `autocomplete`, and `inputmode` communicate data
  purpose and let the browser provide useful input behavior. The
  server still validates every submitted value.
- An invalid control keeps its visible label, sets `aria-invalid` when
  application validation owns the state, and references its specific
  error. A summary can collect errors and link to their controls.
- `disabled` removes a native control from submission and ordinary
  focus. `readonly` keeps a supported form control focusable and
  submitted while preventing edits. `aria-disabled` reports state but
  does not prevent events or form submission by itself.
- Pending submission prevents repeated intent, exposes visible status,
  and restores controls on every completion path. An idempotent server
  contract still protects writes when the network result is uncertain.

Dynamic content must expose state without stealing attention. A
`status` region normally announces a polite update such as `Saved`.
An `alert` is assertive and should be reserved for urgent information.
Create a live region before changing its text; adding an
already-filled region may not produce a consistent announcement.
Keep loading, empty, partial, error, and success states visible as
ordinary text too.

Give non-text content an equivalent purpose:

- Informative images need concise alternative text for their role in
  context. Decorative images use empty alternative text so they do not
  repeat nearby content.
- Icon-only controls need a name that describes the action, not the
  icon's shape. Inline SVG may need an explicit name when it carries
  information.
- Prerecorded video needs captions for speech and meaningful sound;
  audio information needs a transcript or equivalent. Describe visual
  information that the audio does not convey when the workflow needs
  it.
- Canvas content needs a semantic alternative and keyboard-operable
  controls. Do not place essential instructions only inside an image.

Perception and motor access affect layout decisions. Do not rely on
color, position, sound, hover, or motion as the only signal. Preserve
text contrast, visible focus, pointer target size or spacing, and
content at zoom and narrow reflow. Let text grow without fixed-height
clipping. Respect reduced-motion and forced-color preferences, and
keep focused content from hiding under sticky headers or overlays.

**Internationalization** makes a product structurally ready for
different languages and cultural formats; **localization** supplies a
specific locale's text and conventions. Language and direction are
different facts: `lang="ar"` identifies Arabic, while `dir="rtl"`
sets right-to-left direction. Isolate user-supplied bidirectional text
with `bdi` or an equivalent boundary so punctuation and neighboring
text keep the intended order.

Translate complete messages rather than concatenating fragments whose
word order may change. Use locale-aware plural, number, date, time,
currency, list, and segmentation APIs with an explicit locale and,
when relevant, time zone and currency. Expect text expansion,
different line breaking, different name and address shapes, and
right-to-left layout. Locale is a user or product choice, not a safe
inference from geographic location.

**Progressive enhancement** starts with semantic content and a working
core action, then adds CSS and JavaScript capabilities. Feature-detect
optional APIs and preserve a usable fallback. This does not require
every workflow to work without JavaScript; it requires the dependency
and failure state to be deliberate rather than a blank page.

Accessibility is verified by several kinds of evidence. Automated
checks catch selected markup, name, state, and contrast failures.
Semantic component tests verify the exposed contract. Keyboard, zoom,
reflow, reduced motion, forced colors, screen readers, voice input,
and real user evaluation reveal failures that a rule engine cannot
decide. Test every important workflow state, not only the ready
screen.

#### Step-by-step examples

First input: a valid email `nolan@example.com` in a form with a
visible `Email` label, native email validation, a `Subscribe` button,
and an existing empty `role="status"` region. Expected output: one
save runs, the button is unavailable while pending, and visible status
becomes `Subscription saved`.

1. **Before:** Focus is on the labeled email input. The native form
   owns Enter-key submission and the email satisfies its constraint.
   **Decision:** Activate the native submit action. **After:** one
   `submit` event reaches the handler. **Known correct:** keyboard and
   pointer activation share one path.
2. **Before:** Submission is pending. **Decision:** Disable duplicate
   submission and set visible status to `Saving`. **After:** one save
   owns the pending state and another submit returns `ignored`.
   **Known correct:** repeated intent cannot start another save.
3. **Before:** The request succeeds. **Decision:** Replace the status
   text with `Subscription saved`, then enable the button in cleanup.
   **After:** the result is visible and the polite status can announce
   it. **Known correct:** the workflow has one terminal result and is
   ready for a later submission.

Failure input: the save rejects. Expected output: visible status is
`Could not save the subscription`, the button is enabled again, and
the email value remains for correction or retry.

1. **Before:** The save owns pending state and then rejects.
   **Decision:** Convert that failure into the explicit error status.
   **After:** the draft remains and the error is visible.
   **Known correct:** a transport failure does not erase user input.
2. **Before:** The error status is set while the button is disabled.
   **Decision:** Run the same cleanup as success. **After:** the
   button is enabled. **Known correct:** no terminal path traps the
   form in a pending state.

Second input: a toolbar has buttons `[Bold, Italic, Underline]`.
`Bold` has DOM focus and `tabindex="0"`; its peers have `-1`. Press
ArrowRight twice. Expected output: focus and `tabindex="0"` end on
`Underline`, while text-format selection remains unchanged until an
activation key is used.

1. **Before:** focus is on `Bold`; only `Bold` is in the tab sequence.
   **Decision:** ArrowRight selects `Italic` as the next focus target,
   sets `Bold` to `-1`, `Italic` to `0`, and focuses `Italic`.
   **After:** `Italic` owns focus. **Known correct:** the toolbar
   still contributes one tab stop.
2. **Before:** `Italic` owns focus and no format was activated.
   **Decision:** Apply the same roving rule for the second ArrowRight.
   **After:** `Underline` owns focus and `tabindex="0"`; the others
   are `-1`. **Known correct:** focus movement did not silently change
   the selected formatting state.
3. **Before:** focus is on `Underline`. **Decision:** Press Tab.
   **After:** focus leaves the toolbar for the next page component.
   **Known correct:** arrow keys move within the composite and Tab
   moves between components.

Internationalization boundary input: locale `ar-EG`, direction `rtl`,
item count `2`, and price `2599` minor Egyptian-pound units. Expected
output: the document declares Arabic and right-to-left direction, the
complete plural message is selected for Arabic, and price formatting
uses explicit locale and currency rather than concatenated symbols.

1. **Before:** the count, amount, locale, and currency are data
   without display text. **Decision:** select a complete localized
   plural message and format `25.99` with `Intl.NumberFormat` for
   `ar-EG` and `EGP`. **After:** grammar, digits, separators, and
   currency placement come from one locale contract.
   **Known correct:** code did not assume English word order or a
   currency symbol position.
2. **Before:** Arabic text contains a user-provided Latin product ID.
   **Decision:** isolate the product ID's direction inside the Arabic
   sentence. **After:** neighboring punctuation keeps its intended
   order. **Known correct:** language choice and direction boundaries
   are represented separately.
3. **Before:** the localized button text is longer than English.
   **Decision:** reflow it at narrow width and zoom without a fixed
   height. **After:** text remains visible and the control remains
   operable. **Known correct:** translation did not invalidate the
   interaction layout.

#### Why it works

Native elements expose synchronized semantics, focus, keyboard, form,
and platform behavior before application code adds anything. Source
order gives reading and focus order one shared structure. Visible
labels and text make the same purpose and state available through
visual, speech, and automation interfaces.

For a custom widget, the established pattern defines one consistent
contract. The toolbar always has one ordinary tab stop, while roving
focus makes each child reachable without adding all children to the
page's tab sequence. Keeping focus separate from selection prevents
navigation from performing an unintended edit.

The form has one owner for pending state. Every accepted submission
sets pending before work starts and every completion restores the
control, so overlapping saves and permanent disabled states cannot
arise through that handler. The pre-existing status region changes
text as state changes, allowing visual and assistive outputs to
observe the same result.

Language, direction, and locale-aware formatting encode facts that
would otherwise be guessed in scattered components. Complete messages
let a translation choose its own grammar. Flexible layout preserves
the control contract when those messages occupy different space.

#### TypeScript example

This example requires a native form `#subscription`, a labeled
required email input, a submit button, and an existing
`<p id="status" role="status"></p>`. Native constraint validation runs
before the `submit` event. Production save code must also send the
value to a server that validates it.

```typescript
type SubmissionControls = Readonly<{
  status: Pick<HTMLElement, "textContent">;
  submit: Pick<HTMLButtonElement, "disabled">;
}>;

type SaveSubscription = () => Promise<void>;
type SubmissionResult = "failed" | "ignored" | "saved";

async function submitSubscription(
  controls: SubmissionControls,
  save: SaveSubscription,
): Promise<SubmissionResult> {
  // One pending save owns the controls until its terminal cleanup.
  if (controls.submit.disabled) return "ignored";
  controls.submit.disabled = true;
  controls.status.textContent = "Saving";

  try {
    await save();
    controls.status.textContent = "Subscription saved";
    return "saved";
  } catch {
    controls.status.textContent = "Could not save the subscription";
    return "failed";
  } finally {
    controls.submit.disabled = false;
  }
}

// Fail before registering only part of the browser interaction.
const form = document.querySelector<HTMLFormElement>("#subscription");
const statusElement = document.querySelector<HTMLElement>("#status");
const submitSelector = 'button[type="submit"]';
const submit = form?.querySelector<HTMLButtonElement>(submitSelector);
if (!form || !statusElement || !submit) {
  throw new Error("subscription form is incomplete");
}

async function saveSubscription(): Promise<void> {
  // The real boundary would submit FormData to the authorized server.
  await Promise.resolve();
}

form.addEventListener("submit", (event) => {
  event.preventDefault();
  const controls = { status: statusElement, submit };
  void submitSubscription(controls, saveSubscription);
});
```

#### Work and memory

The browser visits relevant nodes while creating DOM, style, layout,
rendering, and accessibility information. Deep redundant markup and
large custom widgets add nodes, event work, state, and retained
memory. Images, media, fonts, and layout shifts usually cost more than
a few semantic elements; measure the real page instead of assigning
a universal node limit.

The example keeps constant state and performs constant work around one
save; network and server cost belong to the save boundary. A
composite with `n` children may inspect or update up to `n` states if
implemented naively, though roving focus normally changes the old and
new child. Translation catalogs consume transfer and memory by locale.
Load and cache them deliberately without forcing every user to
download every language.

Live announcements, focus movement, and animation impose cognitive
cost even when runtime cost is small. Prefer one meaningful update
over several intermediate announcements. Captions, transcripts,
alternative text, and localized messages are required product content
and need the same ownership and review as visible copy.

#### When to use it

Use semantic structure, flexible layout, and explicit state for every
interface. Use native controls for ordinary navigation, actions,
forms, disclosure, and media. Use ARIA patterns when the required
composite interaction has no suitable native control.

Internationalize when a product may serve more than one language,
region, currency, direction, or cultural format. Designing message and
layout boundaries early is cheaper than extracting English fragments
and physical-direction assumptions from a mature product.

#### When not to use it

Do not add ARIA to duplicate correct native semantics or make a
generic element sound interactive without implementing its behavior.
Do not force a native element whose contract differs from the required
interaction; a spreadsheet-like grid may need a carefully implemented
composite widget.

Do not move focus or announce content merely because rendering
occurred. Do not localize stable identifiers, protocol values, CSS
class names, or data that another system expects unchanged. Separate
machine contracts from user-facing language.

#### Common mistakes and edge cases

> **Warning:** A `div` with a click handler is not a button. It lacks
> keyboard activation, button semantics, focus behavior, and disabled
> behavior unless all are rebuilt correctly.

- ARIA changes semantics but does not add behavior or repair invalid
  HTML.
- An `aria-label` can replace useful visible text in the accessible
  name and can be forgotten during translation.
- Placeholder text is not a persistent label.
- A heading chosen for its size creates a misleading outline.
- Several unnamed `nav` regions are difficult to distinguish.
- CSS visual reordering can disagree with reading and focus order.
- Positive `tabindex` values create a second, fragile focus order.
- `tabindex="0"` on every composite child creates excessive tab stops.
- Focus and selection are different states and need distinct visuals.
- Removing a focused node without choosing a successor loses position.
- Removing focus outlines hides the user's current position.
- A sticky header can obscure the focused control after navigation.
- `aria-hidden="true"` on a focusable subtree creates contradictory
  output.
- A modal-looking overlay without modal focus behavior leaves the page
  operable behind it.
- `aria-disabled="true"` does not stop click handlers or submission.
- `disabled` form controls are omitted from normal submission.
- Client validation can guide but cannot authorize or protect data.
- Adding a filled live region may not trigger an announcement.
- An assertive live region for routine updates repeatedly interrupts.
- Alternative text that repeats a caption creates duplicate output.
- Icon names such as `floppy disk` describe appearance, not `Save`.
- Autoplay, flashing, and unavoidable motion can block safe use.
- Fixed heights clip translated, zoomed, or user-sized text.
- Color-only validation fails when color cannot be perceived.
- Hover-only controls fail for keyboard, touch, and speech input.
- Hard-coded left and right properties can break right-to-left layout.
- Concatenated translated fragments assume one grammar and word order.
- Locale without an explicit time zone can show the wrong civil time.
- A language picker should name languages in understandable terms and
  must not depend only on flag icons.
- Automated scans cannot judge reading order, useful names, keyboard
  efficiency, or screen-reader workflow quality.
- Accessibility overlays cannot repair application ownership and
  interaction defects reliably.

#### Related material

- [Frontend curriculum](README.md)
- [CSS, layout, responsive design, and rendering][css]
- [Server data, routing, caching, and forms][data]
- [Testing, performance, and observability][quality]
- [Frontend architecture and design systems][architecture]
- [Web Content Accessibility Guidelines 2.2](
  https://www.w3.org/TR/WCAG22/)
- [ARIA Authoring Practices Guide](
  https://www.w3.org/WAI/ARIA/apg/)
- [W3C internationalization techniques](
  https://www.w3.org/International/techniques/authoring-html)

[quality]: #b3-testing-performance-and-frontend-observability
[architecture]: #c1-frontend-architecture-and-design-systems
[css]: #a6-css-layout-responsive-design-and-rendering

<a id="a2-javascript-typescript-and-the-browser-runtime"></a>

### A.2 JavaScript Language and Browser Runtime

#### The problem it solves

Frontend code must update shared values, react to user input, wait for
platform work, and keep the main thread available for rendering.
Without a precise execution model, closures read surprising values,
promises run in an unexpected order, and one long callback freezes the
interface.

#### How it works

JavaScript works with **primitive values** and objects. Strings,
numbers, big integers, booleans, symbols, `null`, and `undefined` are
primitive values. Assigning a primitive copies that value. Assigning
an object copies a reference to the same object, so a mutation through
either reference is visible through the other. A `const` declaration
prevents rebinding; it does not freeze the referenced object.

Use strict equality, `===` and `!==`, unless a coercing comparison is
deliberately required. A **nullish** value is only `null` or
`undefined`; a **falsy** value also includes `0`, `NaN`, `""`, and
`false`. The distinction decides whether `??` or `||` is correct.

`let` and `const` create block-scoped bindings. A function remembers
the lexical bindings visible where the function was created. That
remembered environment is a **closure**. A closure captures the
binding, not a frozen copy of its current value. An arrow function
also uses `this` from its surrounding scope, while a normal function's
`this` depends on how that function is called.

An ECMAScript module has its own scope, runs once for each module
instance, and exposes live imported bindings. Static imports make the
dependency graph analyzable before evaluation. Dynamic `import()`
loads a module through a Promise when the path is actually reached.

Synchronous JavaScript uses an execution stack. One callback runs to
completion before another callback starts on the same agent. The
browser's **event loop** coordinates that execution with input,
networking, timers, and rendering:

- A **task** runs event handlers, timer callbacks, and other platform
  work from a selected task source.
- After the current stack empties, the browser performs a
  **microtask checkpoint** and drains queued Promise reactions and
  other microtasks, including microtasks added by microtasks.
- The browser may then take a rendering opportunity before selecting
  another runnable task. Rendering is not guaranteed after every task.

Calling an `async` function starts synchronously. Reaching `await`
suspends that invocation and arranges the rest as a Promise reaction.
The continuation therefore runs as a microtask after the awaited value
settles; `await` does not block the thread or create a new one.

Browser APIs perform work outside the current JavaScript callback and
queue a later result. `AbortController` sends a cancellation signal to
APIs that honor it, but cancellation is cooperative. A server may have
already committed a write. Web Workers run in another agent and cannot
touch the document directly. Messages cross the worker boundary by
structured cloning or by transferring supported ownership.

Objects remain in memory while reachable. Event listeners, timers,
module state, and closures can retain an otherwise unused subtree.
Garbage collection timing is not a correctness or cleanup mechanism;
release external resources explicitly.

#### Step-by-step examples

Input: one task records `task:start`, queues two microtasks and a zero
delay timer, then records `task:end`. Expected output:
`task:start`, `task:end`, `microtask:promise`,
`microtask:queued`, `task:timer`.

1. **Before:** The task begins with `events = []`. **Decision:** Push
   `task:start`, queue the Promise reaction, then queue the explicit
   microtask. **After:** `events = [task:start]`; the microtask queue
   contains `promise`, then `queued`. **Known correct:** Neither
   microtask can interrupt the running task.
2. **Before:** The same task still owns the stack. **Decision:** Queue
   the timer task and push `task:end`. **After:**
   `events = [task:start, task:end]`; both microtasks and the timer
   are pending. **Known correct:** All synchronous work finished in
   source order.
3. **Before:** The stack is empty and the microtask queue is
   `[promise, queued]`. **Decision:** Drain it from oldest to newest.
   **After:** `events` ends with `microtask:promise`, then
   `microtask:queued`. **Known correct:** The microtask queue is empty
   before another task starts.
4. **Before:** The microtask queue is empty. The timer later becomes
   runnable. **Decision:** Run its callback and append `task:timer`.
   **After:** The expected five-item sequence is complete.
   **Known correct:** A zero delay did not bypass task ordering.

Closure boundary example: create `readQuery = () => query` while
`query` is `"ca"`, then assign `query = "cat"`. Calling `readQuery()`
returns `"cat"` because it reads the captured binding. Copy
`const submittedQuery = query` when a stable snapshot is required.

#### Why it works

Run-to-completion prevents another callback on the same agent from
observing a half-finished synchronous update. When that callback ends,
the HTML event-loop contract drains the microtask queue before the
next task. Queue insertion order therefore establishes the two
microtask results in the trace.

Closures stay connected to lexical bindings, so later reads observe
the binding's current value. Promises schedule reactions instead of
executing them inline, and platform callbacks re-enter JavaScript only
through later queued work. These contracts explain the order without
depending on processor speed.

#### TypeScript example

This TypeScript uses only JavaScript runtime behavior; its annotations
document the returned collection.

```typescript
async function observeEventOrder(): Promise<readonly string[]> {
  const events: string[] = [];

  // The timer Promise lets the example wait for the later task
  // without changing the queue order being demonstrated.
  const timerFinished = new Promise<void>((resolve) => {
    setTimeout(() => {
      events.push("task:timer");
      resolve();
    }, 0);
  });

  events.push("task:start");
  Promise.resolve().then(() => {
    events.push("microtask:promise");
  });
  queueMicrotask(() => {
    events.push("microtask:queued");
  });
  events.push("task:end");

  // Awaiting the timer resumes in the microtask checkpoint after the
  // timer callback has appended the final task event.
  await timerFinished;
  return events;
}

observeEventOrder().then((events) => {
  const expected = [
    "task:start",
    "task:end",
    "microtask:promise",
    "microtask:queued",
    "task:timer",
  ];
  console.assert(events.join("|") === expected.join("|"));
});

// A closure reads its captured binding, not a creation-time copy.
let query = "ca";
const readQuery = (): string => query;
query = "cat";
console.assert(readQuery() === "cat");
```

#### Work and memory

Each callback occupies the main thread until it returns. A loop over
`n` items still performs work proportional to `n` even when started
from a Promise. Splitting work across tasks can improve responsiveness
but adds scheduling overhead and does not reduce total computation.

Each queued callback and closure retains its function plus reachable
captured state. An unbounded microtask chain can retain data, consume
CPU, and postpone rendering and input indefinitely. Workers add their
own runtime memory; cloning a message copies supported data, while
transferring an `ArrayBuffer` moves its backing storage.

#### When to use it

Use closures for private state and callbacks whose lifetime is clear.
Use promises and `async` functions for one eventual result. Use an
abort signal when later intent makes in-flight work obsolete. Use a
worker when measured CPU work is large enough to justify messaging
and the work does not require direct DOM access.

Use a new task, a scheduler API supported by the target browsers, or a
worker when work must yield to input and rendering. Use a microtask
only for short ordering work that must finish before the next task.

#### When not to use it

Do not use a Promise to make CPU-heavy work parallel; its callback
still runs on the same agent. A microtask loop is not a yielding
strategy because the checkpoint drains until empty. Do not move small
work to a worker when serialization and coordination cost more than
the computation.

Do not treat aborting `fetch` as transaction rollback. Use server-side
operation identity and idempotency for uncertain writes.

#### Common mistakes and edge cases

> **Warning:** `setTimeout(callback, 0)` means the callback becomes
> eligible after the delay and current ordering rules. It does not run
> immediately, before microtasks, or at an exact time.

- `const` does not make an object immutable.
- `||` replaces valid `0`, `false`, and empty-string values.
- `Array.prototype.forEach` does not await an async callback.
- A rejected Promise can become unhandled before a later catch exists.
- Long microtask chains can block rendering as surely as a long task.
- A detached DOM node remains retained through listeners or closures.
- `postMessage` cannot structured-clone functions or DOM nodes.
- String indexes count UTF-16 code units, not visible graphemes.
- Module imports are live bindings, not copied export snapshots.

#### Related material

- [TypeScript type modeling and runtime boundaries](
  #a3-typescript-type-modeling-and-runtime-boundaries)
- [Modules, toolchains, and production builds][toolchain]
- [HTTP, caching, and network delivery][network]
- [React components, state, and effects][react]
- [HTML event-loop specification](
  https://html.spec.whatwg.org/multipage/webappapis.html)
- [ECMAScript specification](https://tc39.es/ecma262/)

[react]: #b1-react-components-state-and-effects
[network]: #a4-http-caching-and-network-delivery

### A.3 TypeScript Type Modeling and Runtime Boundaries

#### The problem it solves

JavaScript accepts values whose shapes may not match the assumptions
inside an application. TypeScript should make those assumptions
visible, rule out impossible application states, preserve useful
relationships between inputs and outputs, and force untrusted data to
earn a trusted type.

#### How it works

TypeScript checks JavaScript before execution and then erases its type
syntax. A type annotation changes compiler evidence, not the runtime
value. Let inference describe obvious local values; write explicit
types at exported APIs, state boundaries, callbacks, and untrusted
inputs where the contract must remain stable.

TypeScript is structurally typed: compatibility depends mainly on
available members rather than a declared class name. This works well
for JavaScript composition, but a wider object can be assigned to a
narrower structural type. A type is not a permission list for which
runtime properties exist.

A **union** means a value may be one of several types. Add one literal
field, such as `kind: "ready"`, to each alternative to create a
**discriminated union**. Checking that field narrows the value to one
alternative. After every alternative is handled, assigning the
remainder to `never` makes a newly added case a compile error.

Use `unknown` for a value that has arrived but has not been checked.
Operations on it require narrowing with `typeof`, equality,
`instanceof`, property checks, or a validator. `any` opts out of that
evidence and spreads the opt-out through later expressions. A type
assertion with `as` tells the checker to trust the author; it does not
inspect, convert, or validate a value.

A return type such as `value is string[]` is a **type predicate**.
When that function returns true, the checker narrows the caller's
value. The function body must really establish the claimed fact;
TypeScript does not prove that a custom predicate is implemented
correctly.

A generic type parameter preserves a relationship. For example,
`function first<Value>(items: readonly Value[]): Value | undefined`
connects the input element type to the output type. Use a generic only
when callers can observe such a relationship; a type parameter used
once usually hides a simpler concrete or `unknown` type.

`readonly` prevents mutation through that typed reference, not through
all aliases, and it is shallow unless nested values are also readonly.
Optional properties differ from properties whose values include
`undefined`. Compiler options such as `noUncheckedIndexedAccess` and
`exactOptionalPropertyTypes` expose those boundary distinctions.
The `satisfies` operator checks an expression against a type while
retaining the expression's more specific inferred type.

At a network, storage, URL, worker-message, or user-input boundary,
validate the runtime value first. A validator may be hand-written or
generated from one owning schema, but its success type and actual
checks must agree. After validation, copy or normalize fields when the
application needs a stable snapshot rather than a mutable external
object.

At this book's dated snapshot, TypeScript 7 is the native compiler and
uses strict and modern module defaults. Version 7.0 does not expose a
stable programmatic compiler API, so tools for embedded languages or
custom compiler integrations may still require TypeScript 6. Pin the
workspace compiler and verify the complete toolchain before upgrading.

#### Step-by-step examples

Input: unknown JSON
`{"kind":"ready","requestID":"r7","items":["lamp","desk"]}`.
Expected output: a trusted `ready` state containing request `r7` and a
copied two-item string array.

1. **Before:** `value` is `unknown`; no property access is safe.
   **Decision:** Require a non-null object and view its keys as
   `Record<string, unknown>`. **After:** `record.kind`,
   `record.requestID`, and `record.items` remain unknown values.
   **Known correct:** Primitive and null inputs have been rejected.
2. **Before:** The record exists but its variant is unconfirmed.
   **Decision:** Require `record.kind === "ready"` and a string
   `requestID`. **After:** The variant and request identity are known.
   **Known correct:** No other state can enter the ready-state path.
3. **Before:** `record.items` is still unknown. **Decision:** Require
   an array and visit every element to prove it is a string.
   **After:** `record.items` narrows to `string[]`. **Known correct:**
   Reading any item as a string is safe.
4. **Before:** The validated array still belongs to boundary data.
   **Decision:** Copy it into a new ready state. **After:** The result
   is `{ kind: "ready", requestID: "r7", items: ["lamp", "desk"] }`.
   **Known correct:** The application owns a validated snapshot.

Failure example: input
`{"kind":"ready","requestID":"r8","items":["lamp",2]}` is
rejected at the array check. Silently filtering out `2` would accept a
corrupt response and change its meaning.

#### Why it works

Every use of boundary data follows a runtime check that establishes
the exact fact the type system later relies on. The object check makes
key lookup safe, the literal check selects one union alternative, and
the item predicate proves every array element. Copying prevents later
external mutation from invalidating the accepted snapshot.

Within application code, the discriminant connects each state to its
available fields. Exhaustive `never` checking makes the compiler
reject a switch that no longer covers the whole union. These are
compile-time guarantees; the validator remains necessary because the
source value existed before TypeScript could trust it.

#### TypeScript example

```typescript
type LoadState =
  | { kind: "idle" }
  | { kind: "loading"; requestID: string }
  | { kind: "ready"; requestID: string; items: readonly string[] }
  | { kind: "failed"; requestID: string; message: string };

function isStringArray(value: unknown): value is string[] {
  if (!Array.isArray(value)) return false;
  return value.every((item) => typeof item === "string");
}

function parseReadyState(value: unknown): LoadState {
  // Establish an object boundary without claiming anything about
  // the types stored under its keys.
  if (typeof value !== "object" || value === null) {
    throw new Error("ready state must be an object");
  }
  const record = value as Record<string, unknown>;

  // Narrow the discriminant, identity, and complete item collection
  // before constructing application-owned state.
  if (record.kind !== "ready") {
    throw new Error("state kind must be ready");
  }
  if (typeof record.requestID !== "string") {
    throw new Error("requestID must be a string");
  }
  if (!isStringArray(record.items)) {
    throw new Error("items must contain only strings");
  }
  return {
    kind: "ready",
    requestID: record.requestID,
    items: [...record.items],
  };
}

function statusText(state: LoadState): string {
  switch (state.kind) {
    case "idle":
      return "Not started";
    case "loading":
      return `Loading ${state.requestID}`;
    case "ready":
      return `${state.items.length} items`;
    case "failed":
      return `Failed: ${state.message}`;
    default: {
      // Adding a union member makes this assignment fail until the
      // new observable state receives deliberate output.
      const unreachable: never = state;
      return unreachable;
    }
  }
}

const readyState = parseReadyState({
  items: ["lamp", "desk"],
  kind: "ready",
  requestID: "r7",
});
console.assert(statusText(readyState) === "2 items");

let invalidStateRejected = false;
try {
  parseReadyState({
    items: ["lamp", 2],
    kind: "ready",
    requestID: "r8",
  });
} catch {
  invalidStateRejected = true;
}
console.assert(invalidStateRejected);
```

#### Work and memory

The union and `never` check emit no runtime type metadata. The switch
is ordinary runtime control flow. If the boundary array contains `n`
items, validation visits all `n` in the successful case and stops at
the first invalid item in the failure case. Copying a valid array
performs another `n` visits and stores `n` references in output.

Type checking has development and continuous-integration cost.
Deeply recursive conditional types, huge inferred unions, and broad
generated declarations can consume substantial compiler time and
memory. TypeScript 7 improves compiler throughput, but simple public
types and project boundaries remain valuable design constraints.

#### When to use it

Use discriminated unions for mutually exclusive UI, request, workflow,
and reducer states. Use `unknown` and runtime validation whenever a
value crosses a trust or persistence boundary. Use generics when an
API must preserve an observable relationship between caller-chosen
types. Use `satisfies` for checked configuration that should retain
literal inference.

Enable strict checking and choose additional indexed-access and
optional-property rules deliberately for new code. Migrate older code
in bounded stages rather than hiding errors behind project-wide `any`.

#### When not to use it

Do not model every business rule with type-level computation. Runtime
rules that depend on current data, authorization, database state, or
user input still need runtime checks. Do not add a generic parameter
when a concrete type or `unknown` states the contract more clearly.

Do not publish compiler-internal types as an application API. If a
tool needs the TypeScript programmatic API, verify that the selected
compiler version exposes the API and that every plugin supports it.

#### Common mistakes and edge cases

> **Warning:** `response.json() as User` changes only the checker's
> belief. Input `{ "name": 7 }` remains the same invalid runtime
> object. Parse from `unknown` and reject the mismatched field.

- `any` spreads unchecked operations into otherwise strict code.
- A non-null assertion can hide a real missing element or ref.
- `readonly` is shallow and does not freeze the runtime object.
- Object spread and structural assignment can retain extra fields.
- An optional key is different from a present key holding `undefined`.
- Indexed access may be absent even when its annotation says `Value`.
- A validator and its claimed type can drift when maintained twice.
- A large assertion can hide the exact fact that was never checked.
- Exhaustive checks need a stable discriminant on every union member.
- Compiler and editor versions can disagree unless the workspace
  version is selected explicitly.

#### Related material

- [JavaScript language and browser runtime](
  #a2-javascript-language-and-browser-runtime)
- [HTTP, caching, and network delivery][network]
- [Browser storage, origins, and security][security]
- [React components, state, and effects][react]
- [TypeScript Handbook](https://www.typescriptlang.org/docs/handbook/)
- [TypeScript release notes](
  https://devblogs.microsoft.com/typescript/)

<a id="a3-networking-storage-and-browser-security"></a>

### A.4 HTTP, Caching, and Network Delivery

#### The problem it solves

A frontend must request the right representation, interpret transport
and application failures separately, reuse responses without exposing
another user or variant, and recover from uncertain connections
without duplicating a write.

#### How it works

A URL identifies where and how to address a resource. Its scheme and
authority select the network endpoint; its path and query identify the
target used by the application. The fragment stays in the user agent
and is not sent in an HTTP request.

Before an HTTPS request can carry application data, the client usually
resolves a host, establishes a transport connection, and negotiates
Transport Layer Security (TLS). Reused connections can avoid some of
that setup. HTTP/2 multiplexes messages over a connection using TCP.
HTTP/3 uses QUIC over UDP. Loss in one stream need not block delivery
of other streams in the same way. All versions share HTTP semantics.

An HTTP request has a method, target, fields called headers, and
optionally content. A response has a status code, headers, and
optionally content. The method expresses intent:

- A **safe** method is intended only to read. `GET` and `HEAD` are
  safe even though servers may still log or measure them.
- An **idempotent** method is intended to have the same final effect
  when repeated. `PUT` and `DELETE` are idempotent by specification;
  `POST` is not automatically idempotent.
- An application can make one logical `POST` replayable by assigning
  it one stable operation key and recording the result on the server.

`fetch()` fulfills its Promise when an HTTP response is available,
including a `404` or `500`. Check `response.ok` or the exact expected
status. It rejects for failures such as an unusable network response
or an abort. A successful status still does not prove that decoded
content matches the application contract.

An HTTP cache stores a response with metadata and reuses it only when
the request and policy allow. `Cache-Control: max-age=60` makes a
response fresh for a bounded time. `no-cache` allows storage but
requires validation before reuse. `no-store` requests that caches not
store the response.

A validator such as `ETag: "v4"` names a selected representation. A
later request can send `If-None-Match: "v4"`; a `304 Not Modified`
response tells the cache to reuse its stored content and refresh the
applicable metadata. `Vary` adds selected request-header values to the
cache match. User identity, authorization, locale, content encoding,
and feature variants need deliberate cache boundaries.

Compression reduces transferred bytes but costs CPU and decoded
memory. Streaming lets a consumer process chunks before the full body
arrives, but application framing, cancellation, partial failure, and
backpressure still need contracts. Aborting a request stops the
client's interest; it does not prove the server reversed completed
work.

Choose a live-update transport from the direction and delivery need:

- Polling repeats ordinary HTTP requests. It is simple and cacheable,
  but idle clients still make requests and updates arrive only at the
  polling interval.
- Server-Sent Events use `EventSource` for a UTF-8 server-to-client
  event stream. The browser reconnects and can send a last event ID;
  client-to-server work still uses a separate request.
- WebSocket keeps one bidirectional connection. The application
  must define authentication, message schemas, heartbeat, reconnect,
  resume, capacity, and slow-consumer policy. The browser API
  does not provide stream backpressure.
- WebTransport can provide several streams and datagrams over HTTP/3.
  Use it only when those modes and target-browser support solve a
  measured need. WebRTC better fits peer media or data, but
  it adds signaling, connectivity, and privacy work.

Every live stream can disconnect, repeat, omit, or delay application
events. Give resumable events an ordered cursor, make repeated events
safe, bound client buffers, and replace a lost history with a fresh
snapshot. A connected socket is a transport fact, not proof that the
user is still authorized for every later message.

#### Step-by-step examples

Input: `GET /catalog?locale=en-US`. The client has a stale entry with
`etag = "v4"` and `items = ["lamp", "desk"]`; the origin still owns
version `v4`. Expected output: the same two items become fresh without
transferring the JSON body again.

1. **Before:** The cache key is the catalog URL including locale, and
   its stored entry is stale. **Decision:** Send `If-None-Match: "v4"`
   with the `GET`. **After:** The request carries the validator but no
   guessed replacement data. **Known correct:** The `en-US` variant
   cannot reuse a differently keyed locale entry.
2. **Before:** The origin compares `"v4"` with its selected
   representation. **Decision:** Return `304 Not Modified` because the
   validator still matches. **After:** No catalog representation is
   transferred in the response content. **Known correct:** The stored
   bytes still represent the origin's current `en-US` selection.
3. **Before:** The client has the `304` and the stored two-item body.
   **Decision:** Update applicable cache metadata and reuse the stored
   body. **After:** the items are fresh and remain
   `["lamp", "desk"]`. **Known correct:** The application receives a
   validated representation rather than stale or empty content.

Failure example: if locale comes from `Accept-Language` instead of the
query, the response must vary on that field or use another distinct
cache key. Otherwise an English response can be reused for a
Vietnamese request.

#### Why it works

The URL and selected request fields define which representation is
being requested. Freshness policy decides whether the cache may reuse
stored content without contacting the origin. Once stale, the entity
tag gives the origin the final comparison. A matching `304` validates
the stored representation. A non-matching `200` sends a replacement.

Safe and idempotent method semantics guide automatic recovery, while
an application operation key gives a replayed `POST` one logical
identity. Runtime response validation remains separate because HTTP
status and caching describe message delivery, not the application's
JSON shape.

#### TypeScript example

This explicit cache shows the conditional-request contract. A browser
or content delivery network normally performs comparable HTTP cache
work outside application code.

```typescript
type CatalogCache = Readonly<{
  etag: string;
  items: readonly string[];
  locale: string;
}>;

interface FetchCatalog {
  (input: string, init?: RequestInit): Promise<Response>;
}

function isCatalogBody(value: unknown): value is { items: string[] } {
  if (typeof value !== "object" || value === null) return false;
  const record = value as Record<string, unknown>;
  if (!Array.isArray(record.items)) return false;
  return record.items.every((item) => typeof item === "string");
}

async function loadCatalog(
  locale: string,
  cached: CatalogCache | undefined,
  fetchCatalog: FetchCatalog = fetch,
): Promise<CatalogCache> {
  const encodedLocale = encodeURIComponent(locale);
  const catalogURL = `/catalog?locale=${encodedLocale}`;
  const cacheMatches = cached?.locale === locale;
  const headers = new Headers();

  // Revalidate only an entry whose complete variant key matches.
  if (cached && cacheMatches) {
    headers.set("If-None-Match", cached.etag);
  }

  const response = await fetchCatalog(catalogURL, { headers });
  if (response.status === 304) {
    if (!cached || !cacheMatches) {
      throw new Error("304 response has no matching cached body");
    }
    return cached;
  }
  if (!response.ok) {
    throw new Error(`catalog request failed: ${response.status}`);
  }

  // A replacement needs both an application-valid body and a future
  // validator before it can become the next cache entry.
  const value: unknown = await response.json();
  const etag = response.headers.get("ETag");
  if (!isCatalogBody(value) || !etag) {
    throw new Error("catalog response is invalid");
  }
  return { etag, items: [...value.items], locale };
}
```

#### Work and memory

Request cost includes connection setup when reuse is unavailable,
round trips, request and response bytes, server work, decompression,
parsing, validation, and application rendering. If a response has `n`
items and `b` transferred bytes, validation performs `O(n)` item work
and transfer and parsing scale with `b`.

A fresh cache hit can remove a network round trip. Revalidation still
uses a round trip but can avoid the body. Cached entries retain
headers and decoded or encoded content until eviction. More variants
improve correctness but reduce the chance that requests share entries.

#### When to use it

Use long-lived immutable caching for content-addressed assets. Use
short freshness plus validators when content changes at a stable URL.
Use private caching for user-specific responses when reuse is safe.
Use operation keys for writes that users or clients may retry after an
uncertain outcome.

Use streaming for large or progressively useful bodies with a defined
framing and cancellation protocol. Measure real connection reuse and
cache hit rates before optimizing protocol details.

Use polling for modest freshness needs. Use Server-Sent Events for
one-way server updates over HTTP. Use WebSocket when both sides need
frequent messages and the product can operate connection state. Choose
WebTransport or WebRTC only for requirements their delivery models
uniquely satisfy.

#### When not to use it

Do not place private data in a shared cache. Do not cache a response
whose key omits a value that changes its representation. Do not retry
a non-idempotent write automatically unless the application protocol
makes that logical operation replay-safe.

Do not add a client-side cache around data already owned coherently by
the browser, a service worker, or the server-data layer. Multiple
cache owners create contradictory freshness and invalidation rules.

Do not open a permanent connection for rare updates that polling can
serve cheaply. Do not use WebSocket as a substitute for message
identity, authorization, persistence, ordering, or recovery.

#### Common mistakes and edge cases

> **Warning:** `Cache-Control: no-cache` does not mean "do not store."
> It requires validation before reuse. Use `no-store` when storage is
> disallowed by the intended policy.

- A `200` response can still contain invalid application data.
- `fetch()` does not reject merely because the status is `500`.
- A zero-byte `304` is useful only with a matching stored body.
- `Vary: *` prevents normal reuse of the response.
- Redirects can change methods, origins, credentials, and cache state.
- A CDN key that omits authorization or locale can leak wrong data.
- Aborting a request does not prove a server write was rolled back.
- `EventSource` has a constrained request API; plan credentials and
  cross-origin policy around the actual browser interface.
- A reconnect without a cursor can miss updates or apply them twice.
- An unbounded WebSocket receive queue can exhaust browser memory.
- A connection authenticated once still needs later permission checks.
- HTTP/2 or HTTP/3 does not repair oversized payloads or request
  waterfalls created by application dependencies.

#### Related material

- [Browser storage, origins, and security][security]
- [TypeScript type modeling and runtime boundaries](
  #a3-typescript-type-modeling-and-runtime-boundaries)
- [Server data, routing, caching, and forms][data]
- [HTTP Semantics, RFC 9110](https://www.rfc-editor.org/rfc/rfc9110)
- [HTTP Caching, RFC 9111](https://www.rfc-editor.org/rfc/rfc9111)
- [Server-Sent Events](
  https://html.spec.whatwg.org/multipage/server-sent-events.html)
- [WebSockets Standard](https://websockets.spec.whatwg.org/)

[data]: #b2-server-data-routing-caching-and-forms

### A.5 Browser Storage, Origins, and Security

#### The problem it solves

Browser code combines data from URLs, storage, windows, frames,
third-party scripts, and authenticated servers. The application must
decide which principal owns each value, which code may read or send
it, how long it may remain, and how untrusted text stays data rather
than becoming executable code.

#### How it works

An **origin** is the tuple of scheme, host, and port. The same-origin
policy normally stops a document from reading another origin's DOM or
responses, while still allowing many cross-origin sends such as links
and forms. Blocking reads alone therefore does not stop forged writes.

Cross-Origin Resource Sharing (CORS) lets a server grant selected
origins permission to expose a response to browser JavaScript. A
preflight checks permission for certain methods and headers before the
actual request. CORS is not authentication, authorization, or
Cross-Site Request Forgery (CSRF) protection. Credentialed CORS needs
an explicit allowed origin rather than `*`.

Browser persistence has several owners and costs:

- Cookies are attached to matching HTTP requests. `Secure` restricts
  transport to secure contexts, `HttpOnly` hides a cookie from script,
  and `SameSite` limits selected cross-site sends. `Domain` and `Path`
  scope delivery; they are not authorization boundaries.
- `localStorage` stores synchronous origin-scoped strings and can
  block the main thread. Any script running in the origin can read it.
  `sessionStorage` has a narrower page-session lifetime.
- IndexedDB stores structured data asynchronously and supports
  transactions. Cache Storage stores `Request` and `Response` pairs,
  commonly for service-worker-controlled delivery.

Storage is subject to implementation-defined quotas, eviction, user
clearing, private modes, and partitioning. Persistent-storage requests
can reduce automatic eviction where supported but cannot override user
control. Modern browsers increasingly partition third-party state by
the top-level site. Partitioned cookies, often called CHIPS, key
embedded cookie state by that top-level context.

A **service worker** is origin-scoped browser code that can mediate
requests for pages inside its registration scope. It is event-driven
and may stop between events. It is not a long-lived Web Worker. Keep
durable state in storage, not only in its memory, and extend
asynchronous event work with `event.waitUntil()`.
Use `respondWith()` in a `fetch` event to provide a response.

Registration starts an update lifecycle. A new worker downloads and
installs resources into a versioned cache. If installation succeeds,
it normally waits while an older worker controls open pages. It then
activates after those clients leave and controls later navigations.
`skipWaiting()` and `clients.claim()` can shorten that handoff. They
can combine an old page contract with a new worker. Use them only when
versions are compatible or the application coordinates the update.

Cache Storage is an application-managed response store; it is separate
from the browser's HTTP cache. Choose a policy by resource contract:

- Use cache-first for immutable, content-addressed assets.
- Use network-first for freshness-sensitive data, with a bounded and
  clearly labelled stale fallback when the product permits one.
- Use stale-while-revalidate only when serving an older value
  immediately is correct under an explicit staleness policy.

Offline reads and offline writes have different correctness rules. A
cached page may truthfully display known data. A write has not
succeeded merely because it entered IndexedDB. Give a queued operation
a stable ID, display `queued`, `failed`, and `confirmed` states, and
make the server idempotent for retries. Background Sync can help where
supported, but delivery remains best effort. Retry when the app next
runs too.

Never put private responses in a cache shared across users. Partition
private entries by authenticated identity where caching is justified,
and clear them on logout. Push and notifications require permission.
They may be unavailable and cannot be the sole path to essential
state.

A web app manifest describes install presentation: the app name,
icons, start URL, scope, and display mode. It neither registers a
service worker nor guarantees offline behavior. Treat installation as
an enhancement and verify the current target-browser support matrix.

**Cross-site scripting (XSS)** occurs when attacker-controlled data is
interpreted as executable markup or script in a trusted origin. Render
plain text through a text API or framework escape path. If users must
author HTML, sanitize it with a maintained policy for that HTML
context. HTML, attributes, URLs, CSS, and JavaScript each have
different parsing rules; one generic escape function cannot cover all
destinations.

Content Security Policy (CSP) restricts allowed resource and script
execution. Nonces or hashes support a strict script policy;
`frame-ancestors` controls which pages may embed the response. Trusted
Types can require dangerous DOM sinks to receive values created by an
approved policy. These controls reduce impact but do not replace safe
APIs, validation, and output handling.

Cookie-authenticated writes need CSRF defenses such as suitable
`SameSite` policy plus a request token or verified origin contract.
Every server action still authenticates the session and authorizes the
specific resource. Logout must invalidate server authority and clear
private client state that should not remain visible.

**Authentication** establishes which account is present;
**authorization** decides whether that account may perform one action.
The frontend can collect credentials and display state, but the server
must establish both facts. Hiding a button never removes server-side
authorization.

For first-party browser applications, a server-managed session cookie
often keeps identity-provider tokens out of JavaScript. OAuth 2.0
delegates access; OpenID Connect adds an identity layer for login. A
browser-only client is public and cannot protect a client secret. When
it participates directly, use the authorization code flow with Proof
Key for Code Exchange (PKCE), exact redirect URIs, and current OAuth
security guidance. Validate `state`, issuer, and the OpenID Connect
nonce where their protocols require them. Never put access tokens in
URLs, and avoid durable script-readable token storage when exposure to
origin script is outside the threat model.

Passkeys use Web Authentication (WebAuthn). An authenticator creates
or uses a public-key credential scoped to a relying party. The server
issues a fresh challenge and verifies the origin, relying-party ID,
challenge, signature, and account policy. The page does not receive
the private key. Provide account recovery and another
supported path when a credential or platform is unavailable.

For `postMessage`, choose an exact target origin when sending. On
receive, verify `event.origin`, the `event.source` window identity,
and the runtime shape of `event.data` before acting. Treat messages as
untrusted even when they arrive from a window the app opened.

#### Step-by-step examples

Input: a checkout page expects payment completion for `order-17` from
origin `https://pay.example` and its known payment frame. Expected
output: `{ type: "payment:complete", orderID: "order-17" }` is
accepted only from that origin, source, and order.

1. **Before:** A message event contains an unknown origin, source, and
   data value. **Decision:** Compare the origin with the exact payment
   origin and the source with the known frame window. **After:** Only
   the expected channel proceeds. **Known correct:** A same-shaped
   message from `https://pay-example.attacker.test` is rejected.
2. **Before:** The channel is trusted but `event.data` is still
   unknown. **Decision:** Require a non-null object with literal type
   `payment:complete` and a string `orderID`. **After:** The value has
   the message's trusted structural type. **Known correct:** Arrays,
   primitives, and other message variants are rejected.
3. **Before:** The structural message may name any order.
   **Decision:** Compare its `orderID` with the checkout order.
   **After:** A fresh accepted object names `order-17`.
   **Known correct:** A valid message for an old or parallel checkout
   cannot complete this order.

Adversarial example: `event.origin.endsWith("pay.example")` is unsafe
because a hostile host can satisfy an imprecise suffix test. Parse
configured URLs and compare their exact serialized origins.

Second input: worker `v7` controls an open catalog page. Worker `v8`
has a versioned shell cache, `/assets/app.8f3.js` is immutable, the
catalog API should be fresh, and order `op-42` is submitted offline.
Expected output: `v8` activates atomically, resources use their
declared policies. The order stays queued until server confirmation.

1. **Before:** `v7` controls a client and `v8` has no complete cache.
   **Decision:** During install, `v8` fills cache `shell-v8` and fails
   installation if a required resource fails. **After:** The
   whole `v8` shell exists, but `v7` still controls the open page.
   **Known correct:** A partial `v8` shell cannot become active.
2. **Before:** `v8` is installed and the `v7` page remains open.
   **Decision:** Use the default waiting phase instead of forcing a
   mixed-version handoff. **After:** `v8` waits until `v7` has no
   controlled clients, then activates and removes obsolete caches.
   **Known correct:** One client uses one compatible page-worker pair.
3. **Before:** A `v8` client requests `/assets/app.8f3.js`.
   **Decision:** Use cache-first because the content hash makes this
   URL immutable. **After:** The cached bytes are returned.
   **Known correct:** That URL always names the same asset contents.
4. **Before:** The client requests the catalog and the network is up.
   **Decision:** Use network-first and update its bounded fallback.
   **After:** Fresh catalog data is shown. **Known correct:** Its
   cache has the latest successfully fetched snapshot.
5. **Before:** A later catalog request is offline. **Decision:** Use
   its allowed stale fallback. **After:** The user sees the last known
   catalog, not an implied fresh result.
   **Known correct:** The interface preserves the snapshot's freshness
   truth.
6. **Before:** Offline order `op-42` has not reached the server.
   **Decision:** Persist its operation ID and show `queued`; retry by
   that ID when execution resumes. **After:** The server accepts it
   once and confirms it. **Known correct:** Only confirmation moves
   `op-42` to `confirmed`; retries cannot duplicate the order.

#### Why it works

The origin comparison identifies which security principal sent the
message. The source identity binds it to the exact frame instance, not
merely another page on that origin. Runtime validation proves the
message variant and fields, and the order comparison binds the event
to the currently authorized workflow. All four facts must hold before
state changes.

The broader browser controls follow the same layered model. Origin and
storage partitioning limit ambient reads, cookie and CSRF policy limit
credentialed sends, contextual output handling prevents data from
becoming code, and CSP or Trusted Types limits damage if an unsafe
sink remains. Server authorization is still the final authority.

The offline lifecycle keeps versions atomic: a worker becomes usable
after its required install work completes, and the default waiting
phase avoids silently changing the request contract of an open page.
Per-resource cache rules remain correct because each rule follows a
stated immutability or freshness contract. A queued write remains
uncertain until an idempotent server confirms its stable operation ID.

#### TypeScript example

An actual message handler supplies `sourceMatches` from
`event.source === paymentFrame.contentWindow`.

```typescript
type MessageEnvelope = Readonly<{
  data: unknown;
  origin: string;
  sourceMatches: boolean;
}>;

type PaymentCompletionMessage = Readonly<{
  orderID: string;
  type: "payment:complete";
}>;

function readPaymentCompletion(
  envelope: MessageEnvelope,
  expectedOrderID: string,
): PaymentCompletionMessage | null {
  // Authenticate the transport channel before inspecting its
  // untrusted application payload.
  if (envelope.origin !== "https://pay.example") return null;
  if (!envelope.sourceMatches) return null;
  if (typeof envelope.data !== "object" || envelope.data === null) {
    return null;
  }

  // Validate the message variant and bind it to this checkout.
  const record = envelope.data as Record<string, unknown>;
  if (record.type !== "payment:complete") return null;
  if (record.orderID !== expectedOrderID) return null;
  return { orderID: expectedOrderID, type: "payment:complete" };
}

const acceptedPayment = readPaymentCompletion(
  {
    data: { orderID: "order-17", type: "payment:complete" },
    origin: "https://pay.example",
    sourceMatches: true,
  },
  "order-17",
);
console.assert(acceptedPayment?.orderID === "order-17");

const rejectedPayment = readPaymentCompletion(
  {
    data: { orderID: "order-17", type: "payment:complete" },
    origin: "https://pay-example.attacker.test",
    sourceMatches: true,
  },
  "order-17",
);
console.assert(rejectedPayment === null);
```

#### Work and memory

Exact origin, source, type, and order comparisons take constant work
for this fixed message shape. A validator for a collection must also
bound and visit that collection. Never parse or allocate an unbounded
message body before enforcing size and rate limits at the closest
available boundary.

Cookies add request bytes to every matching request. Synchronous Web
Storage work blocks its calling thread. IndexedDB and Cache Storage
retain encoded and decoded data until deletion or eviction. Security
headers add small transfer cost; sanitization cost grows with the
untrusted document and should have input-size limits.

Federated login adds redirects, identity-provider requests, token
validation, and session storage. WebAuthn adds a server round trip and
authenticator interaction. Measure completion and recovery without
logging credentials, codes, tokens, or private user data.

Service-worker installation adds requests, storage, and update work.
Request interception adds dispatch and cache lookup work. Network
policies still pay network latency. Bound cache size, entry age,
queued operations, retries, and retained versions. Count offline
storage in the origin's total persistent-data quota.

#### When to use it

Use cookies for server-recognized session state with the narrowest
delivery scope and appropriate flags. Use IndexedDB for substantial
structured offline data. Use Cache Storage for explicit request and
response caching, usually with a service worker. Store only data whose
loss and exposure policies are understood.

Use a service worker when measured offline, repeat-visit, or resilient
delivery requirements justify an origin-level request mediator. Define
update, cache, invalidation, privacy, failure, and rollback policies
before registration. Add a manifest when install presentation helps;
keep the ordinary URL and browser experience complete without it.

Use exact-origin messaging for necessary window or frame integration.
Use framework escaping for ordinary text, maintained sanitization for
intentional rich HTML, and a strict CSP as a measured defense layer.
Apply Trusted Types when the target browsers and migration plan
support enforcement of dangerous DOM sinks.

Use server sessions for first-party login when that architecture fits.
Use maintained, reviewed OAuth and OpenID Connect libraries for
delegated access or federated identity. Offer WebAuthn when phishing
resistance and the recovery design justify it.

#### When not to use it

Do not store bearer credentials in script-readable storage when an XSS
attacker must not read them. Do not treat browser storage as an
authoritative database or guaranteed durable queue. Do not depend on
unpartitioned third-party state for a core workflow.

Do not add a service worker merely to obtain an install prompt or to
cache every response. Prefer the HTTP cache for ordinary reusable
responses. Do not promise offline writes when the server lacks
idempotency, conflict policy, and an honest pending state.

Do not enable broad CORS to repair an API design or use `mode:
"no-cors"` to bypass access control; an opaque response cannot expose
the protected data. Do not accept `postMessage` using `*`, suffix
matching, or shape checks alone for a privileged action.

Do not implement an identity protocol from fragments copied into UI
code. Do not ship a browser client secret, treat an ID token as API
authorization, or rely on route guards as server authorization.

#### Common mistakes and edge cases

> **Warning:** `HttpOnly` stops JavaScript from reading a cookie. It
> does not stop the browser from sending that cookie with a forged
> request, and it does not prevent an injected script from acting as
> the user through same-origin requests.

- CORS grants browser read access; it does not authorize the caller.
- A successful preflight does not prove the actual user is allowed.
- `SameSite` uses site relationships, not the stricter origin tuple.
- Cookie `Domain` and `Path` control delivery, not access.
- `localStorage` is synchronous, string-only, and readable by XSS.
- Storage may be evicted or cleared after a successful earlier write.
- Sanitized HTML can become unsafe after later string concatenation.
- CSP allowlists are brittle; nonce or hash policies need correct
  generation and deployment.
- `targetOrigin = "*"` can disclose a message after window navigation.
- A service worker is powerful origin code and needs an update,
  observability, and rollback strategy.
- In-memory worker state disappears when the browser stops the worker.
- Blind `skipWaiting()` can mix old pages with a new request contract.
- Cache-first on a mutable URL can serve stale content indefinitely.
- A queued offline write is pending, not a successful server result.
- Private cached data can remain visible after logout unless removed.
- A manifest creates neither offline support nor guaranteed install.
- OAuth `state`, PKCE, and an OpenID Connect nonce solve different
  checks; one value does not silently replace all three contracts.
- URL tokens leak through history, logs, referrers, or screenshots.
- A passkey flow still needs server verification and account recovery.

#### Related material

- [HTTP, caching, and network delivery](
  #a4-http-caching-and-network-delivery)
- [Semantic interfaces and accessibility](
  #a1-semantic-interfaces-and-accessibility)
- [Backend identity and API security][backend-security]
- [Fetch Standard](https://fetch.spec.whatwg.org/)
- [Storage Standard](https://storage.spec.whatwg.org/)
- [Trusted Types](https://www.w3.org/TR/trusted-types/)
- [Service Workers](https://www.w3.org/TR/service-workers/)
- [Web App Manifest Editor's Draft](
  https://w3c.github.io/manifest/)
- [OAuth 2.0 Security Best Current Practice](
  https://www.rfc-editor.org/rfc/rfc9700)
- [Proof Key for Code Exchange](
  https://www.rfc-editor.org/rfc/rfc7636)
- [Web Authentication Level 3](https://www.w3.org/TR/webauthn-3/)

[security]: #a5-browser-storage-origins-and-security
[backend-security]:
  ../backend/CONCEPTS.md#9-identity-authorization-and-api-security

<a id="a4-css-layout-responsive-design-and-rendering"></a>
<a id="a5-css-layout-responsive-design-and-rendering"></a>

### A.6 CSS, Layout, Responsive Design, and Rendering

#### The problem it solves

The same document must remain readable and operable across different
content lengths, languages, writing directions, font settings,
containers, viewports, and input devices. CSS rules also need to
combine predictably without making every change a specificity contest.

#### How it works

Start with the document in a useful source order. Normal flow lays out
block content one item after another and inline content along lines.
CSS then changes presentation without making the document depend on a
particular screen size.

Every visible element has a **box model**: content, padding, border,
and margin. With `box-sizing: border-box`, a declared width includes
padding and border. An element also has **intrinsic sizes** derived
from its content. A long unbreakable word, image, or default minimum
size can therefore prevent a track or flex item from shrinking.

The **cascade** chooses one value when declarations compete. First
discard rules whose selectors or conditions do not match. Then compare
origin and importance, cascade layer, selector specificity, scope
proximity when scopes are used, and finally source order. Transitions
and animations have their own cascade positions. For ordinary author
rules, named layers make precedence explicit; unlayered normal rules
override normal layered rules.

Choose a layout tool by relationship:

- Normal flow is the default for documents and stacked controls.
- Flexbox distributes items along one main axis and handles wrapping.
- Grid aligns rows and columns as a two-dimensional layout.
- Positioned layout removes or offsets a box for overlays and anchors;
  it should not replace the document's main flow.

Responsive design follows available space, not a list of fashionable
devices. Media queries describe the viewport or user environment.
Container queries describe the space given to one component. Relative
units, `minmax()`, `min()`, `max()`, and `clamp()` let values change
within deliberate bounds. Logical properties such as
`margin-inline-start` follow the writing direction.

After styles are known, the browser computes geometry during layout,
draws pixels during paint, and combines eligible layers during
composition. A change to width can require layout and paint. A
well-isolated change to `transform` or `opacity` can often be composed
without laying out surrounding content, but that is an optimization to
measure, not a guarantee.

#### Step-by-step examples

Input: three product cards inside a container that changes from
`28rem` to `44rem`. Expected output: each card uses one column in the
narrow container and an image-and-copy grid in the wide container.

1. **Before:** The HTML order is image, heading, description, then
   action; no optional CSS has loaded. **Decision:** Keep that source
   order and let normal flow stack it. **After:** The card remains
   readable as a plain document. **Known correct:** Meaning and action
   order do not depend on visual placement.
2. **Before:** The card has several matching author declarations.
   **Decision:** Apply reset, base, and component layers in
   their declared order. **After:** The component receives one
   computed value for each property. **Known correct:** A selector is
   not made stronger merely to defeat an unknown rule.
3. **Before:** The nearest query container is `28rem` wide and the
   card's base layout is one column. **Decision:** The `36rem`
   container condition is false. **After:** Image and copy remain
   stacked. **Known correct:** The component fits its actual parent,
   regardless of viewport width.
4. **Before:** The same container grows to `44rem`. **Decision:** The
   `36rem` container condition becomes true. **After:** The card uses
   a two-column grid with a bounded image track. **Known correct:**
   DOM and keyboard order stay aligned while presentation adapts.

Boundary example: a product name containing one long URL can force a
grid wider than its container. Use a shrinkable track such as
`minmax(0, 1fr)` and an intentional wrapping rule for user content.
Do not hide overflow until the content-loss policy is understood.

#### Why it works

The document starts with a usable order, so styling is an enhancement
rather than a prerequisite. The cascade produces one computed value
for every property through a deterministic precedence order. Layout
then solves sizes from the containing block, intrinsic content, and
the selected flow algorithm. The container condition depends only on
the component's available inline size, so the same component can work
in a sidebar and a main region without knowing the viewport.

#### TypeScript example

The CSS owns presentation. The TypeScript creates semantic content in
one document update; it does not measure widths or choose a layout.

```css
@layer reset, base, components;

@layer reset {
  *,
  *::before,
  *::after {
    box-sizing: border-box;
  }
}

@layer components {
  .product-list {
    container-type: inline-size;
    display: grid;
    gap: 1rem;
    list-style: none;
    margin: 0;
    padding: 0;
  }

  .product-card {
    display: grid;
    gap: 1rem;
    min-inline-size: 0;
  }

  .product-card img {
    block-size: auto;
    inline-size: 100%;
  }

  .product-card__copy {
    min-inline-size: 0;
    overflow-wrap: anywhere;
  }

  @container (min-width: 36rem) {
    .product-card {
      grid-template-columns: minmax(8rem, 14rem) minmax(0, 1fr);
    }
  }
}
```

```typescript
type ProductCard = {
  detailsURL: string;
  imageURL: string;
  name: string;
  summary: string;
};

function renderProductCards(
  root: HTMLUListElement,
  products: readonly ProductCard[],
): void {
  // Build detached nodes so the live document receives one bounded
  // insertion after every card is complete.
  const fragment = document.createDocumentFragment();

  for (const product of products) {
    const item = document.createElement("li");
    const article = document.createElement("article");
    const image = document.createElement("img");
    const copy = document.createElement("div");
    const heading = document.createElement("h2");
    const summary = document.createElement("p");
    const details = document.createElement("a");

    article.className = "product-card";
    copy.className = "product-card__copy";
    image.src = product.imageURL;
    image.alt = "";
    heading.textContent = product.name;
    summary.textContent = product.summary;
    details.href = product.detailsURL;
    details.textContent = `View ${product.name}`;

    copy.append(heading, summary, details);
    article.append(image, copy);
    item.append(article);
    fragment.append(item);
  }

  // Replace old results once, preserving the source order supplied
  // by the product collection.
  root.replaceChildren(fragment);
}
```

The empty image alternative is correct only when the product name
already communicates the image's meaning. A meaningful product image
needs an alternative that adds information instead of repeating it.

#### Work and memory

If there are `n` cards, creating and inserting their nodes takes work
and memory proportional to `n` plus the text and image resources.
Selector matching, style calculation, layout, and paint depend on the
affected subtree and the properties changed; they are not described
well by DOM count alone. Images usually dominate transfer and decoded
memory. Container queries add style evaluation, while containment can
also limit how far some layout effects travel.

Avoid loops that alternate a geometry read such as
`getBoundingClientRect()` with a style write. That pattern can force
the browser to finish layout repeatedly. Batch reads, then writes, and
prefer CSS when the decision is purely presentational.

#### When to use it

Use normal flow as the starting point, flexbox for one-axis
distribution, grid for two-axis alignment, media queries for viewport
or preference conditions, and container queries for reusable
components whose parent controls available space.

Use cascade layers when an application combines resets, third-party
styles, design-system components, and local overrides. Use logical
properties when layouts must survive different writing directions.

#### When not to use it

Do not use container queries for a condition that truly belongs to the
viewport, such as a page-wide navigation mode. Do not use grid to
reorder meaningful content away from its reading and focus order. Do
not add `will-change` broadly; promoted layers consume memory and can
make performance worse.

JavaScript is appropriate when a decision changes behavior or data,
not only presentation. Server rendering or virtualization is a closer
alternative when the bottleneck is the amount of content rather than
its layout.

#### Common mistakes and edge cases

> **Warning:** Visual reordering does not reliably change reading or
> keyboard order. A card that looks first can still be reached last.
> Keep meaningful DOM order aligned with the intended interaction.

- `100vw` can include scrollbar width and create horizontal overflow.
- `height: 100%` has no effect without a definite containing height.
- A flex or grid item may need `min-inline-size: 0` to shrink.
- Percentage padding is resolved from a containing inline size.
- `z-index` cannot escape an ancestor's stacking context.
- Fixed pixel heights clip zoomed, translated, or user-sized text.
- Hover-only disclosure excludes keyboards and touch interactions.
- Motion needs a reduced-motion policy, not merely a shorter duration.
- Font swaps and unsized images can move already visible content.

#### Related material

- [Semantic interfaces and accessibility](
  #a1-semantic-interfaces-and-accessibility)
- [Testing, performance, and observability][quality]
- [Frontend architecture and design systems][architecture]
- [MDN CSS guides](https://developer.mozilla.org/docs/Web/CSS)

### A.7 Modules, Toolchains, and Production Builds

#### The problem it solves

Application source contains TypeScript, JSX, styles, assets, package
imports, and code for routes that most users never open. Browsers need
resolved JavaScript and asset URLs, not the repository's source graph.
A production toolchain must turn one reviewed dependency graph into
compatible, cacheable artifacts without hiding missing modules,
shipping unnecessary code, or mixing files from different releases.

#### How it works

An ECMAScript module is a file with its own scope and explicit
imports or exports. A **module specifier** is the string after `from`
or passed to `import()`. Relative and absolute URL specifiers have
browser meaning. A bare specifier such as `react` needs a resolver
supplied by an import map, a runtime, or a build tool.

**Module resolution** maps each specifier to one module. TypeScript
must model the same resolver that will run or bundle the JavaScript.
For a browser application bundled by a tool, `moduleResolution` is
commonly `bundler`; Node-targeted code follows the supported Node
mode. A TypeScript `paths` alias changes checking resolution. It does
not rewrite emitted imports or configure another host by itself.

TypeScript checking and JavaScript delivery are separate jobs. Type
annotations and type-only imports are erased. Runtime values remain.
The `target` setting controls which JavaScript syntax may be emitted;
it does not add missing browser APIs. Syntax transforms, API
polyfills, and behavior fallbacks are three separate compatibility
decisions.

A package manifest states requested dependencies. A lockfile records
one resolved dependency graph, including transitive versions and
integrity data. Commit the lockfile and use the package manager's
frozen or clean-install mode in continuous integration. The lockfile
improves reproducibility; it does not prove that dependencies are
trusted, maintained, secure, or compatible with the application.

A production build starts from one or more entry modules, resolves
their dependency graph, applies configured transforms, and emits
JavaScript, CSS, source maps, and referenced assets. A **bundler** may
combine modules into output **chunks**. Minification reduces syntax
and identifier bytes after correctness checks; compression normally
happens when artifacts are served.

Static `import` declarations make initial dependencies visible before
execution. This enables **tree shaking**, also called live-code
inclusion: a bundler can omit exports and branches that it proves are
not needed. Top-level side effects constrain removal. Importing a
module only for registration, global mutation, or CSS is observable
work and must not be mislabeled as side-effect-free.

Dynamic `import()` returns a Promise and creates an asynchronous
boundary. A bundler can place that module and its exclusive static
dependencies in a later chunk. Route, editor, chart, and admin code
are useful split candidates when they are large and not needed for
the first view. Splitting every component creates request overhead,
duplicate loading states, and dependency waterfalls instead.

Output filenames should contain a hash derived from their content,
such as `catalog.8fd3.js`. An unchanged file keeps its URL and can use
long-lived immutable caching. The HTML or manifest that names current
hashes needs a shorter freshness policy. Deploy new hashed assets
before publishing the document that references them, and retain old
assets long enough for already-open pages to finish lazy imports.

Source maps connect generated locations to reviewed source. They help
production diagnosis but can expose names, paths, and source content.
Upload private maps to the authorized error service or deliberately
publish them; do not let an accidental default choose the policy.

A development server optimizes feedback with native modules, cached
transforms, and hot updates. Its graph and timing can differ from the
optimized production build. A release pipeline therefore runs the
real production build, serves its output in a production-like
environment, and tests at least one critical path from those files.

#### Step-by-step examples

Input: `main` statically imports `catalog`; `catalog` statically
imports `price`; `main` dynamically imports `admin`; and `admin`
statically imports `chart`. Expected output: the first-load chunk owns
`main`, `catalog`, and `price`, while an admin chunk owns `admin` and
`chart` and loads only after the admin path is requested.

1. **Before:** The build has entry `main` and no resolved modules.
   **Decision:** Follow static imports from `main` through `catalog`
   to `price`. **After:** The initial closure is
   `{main, catalog, price}`. **Known correct:** Every module required
   to evaluate the entry is in the first-load plan.
2. **Before:** `main` also contains `import("./admin.js")`.
   **Decision:** Record `admin` as a lazy root rather than following
   it into the initial closure. **After:** The first-load plan is
   unchanged. **Known correct:** Opening the catalog does not require
   admin code.
3. **Before:** The admin root has not been planned. **Decision:**
   Follow its static import to `chart`, excluding modules already in
   the initial closure. **After:** The lazy closure is
   `{admin, chart}`. **Known correct:** Requesting admin obtains every
   module needed to evaluate that boundary once.
4. **Before:** The build emits `app.8fd3.js` and `admin.91aa.js`.
   **Decision:** Upload both files, then publish HTML that references
   `app.8fd3.js`. **After:** New pages reference existing artifacts.
   **Known correct:** An old page can still request an older retained
   admin hash during the deployment window.

Failure example: publishing the new HTML before `app.8fd3.js` reaches
every serving location creates an entry failure. Deleting the old
admin chunk immediately can break a tab that loaded the earlier entry
and requests admin later.

#### Why it works

The resolved graph gives every import one build-time owner. The static
closure contains everything needed before the entry can evaluate;
dynamic roots name work that can wait for explicit demand. A bundler
may optimize inside those boundaries only when it preserves module
evaluation and side effects.

Content-derived names make the cache identity equal to the artifact
identity. Publishing immutable assets before their references means a
document never intentionally points at a missing new file. Retaining
the preceding release's assets protects documents whose dependency
graph was fixed before the deployment.

#### TypeScript example

This small planner exposes the graph rule. A real bundler also handles
nested dynamic boundaries, shared chunks, CSS, assets, and transforms.

```typescript
type BuildModule = Readonly<{
  dynamicImports: readonly string[];
  id: string;
  staticImports: readonly string[];
}>;

type FirstLoadPlan = Readonly<{
  initial: readonly string[];
  lazyByRoot: ReadonlyMap<string, readonly string[]>;
}>;

function planFirstLoad(
  modules: readonly BuildModule[],
  entryID: string,
): FirstLoadPlan {
  const moduleByID = new Map<string, BuildModule>();
  for (const module of modules) {
    if (moduleByID.has(module.id)) {
      throw new Error(`duplicate module: ${module.id}`);
    }
    moduleByID.set(module.id, module);
  }

  // Follow one root's static closure. The visited set also makes
  // accidental dependency cycles finite.
  function collectStatic(
    rootID: string,
    excluded: ReadonlySet<string>,
  ): Set<string> {
    const collected = new Set<string>();
    const pending = [rootID];
    while (pending.length > 0) {
      const currentID = pending.pop();
      if (currentID === undefined) break;
      if (excluded.has(currentID) || collected.has(currentID)) {
        continue;
      }
      const currentModule = moduleByID.get(currentID);
      if (!currentModule) {
        throw new Error(`missing module: ${currentID}`);
      }
      collected.add(currentID);
      pending.push(...currentModule.staticImports);
    }
    return collected;
  }

  const initialIDs = collectStatic(entryID, new Set());
  const lazyRootIDs = new Set<string>();
  for (const initialID of initialIDs) {
    const initialModule = moduleByID.get(initialID);
    if (!initialModule) throw new Error("initial module disappeared");
    for (const lazyRootID of initialModule.dynamicImports) {
      lazyRootIDs.add(lazyRootID);
    }
  }

  // Each lazy root owns only its static dependencies that are not
  // already guaranteed by the initial chunk.
  const lazyByRoot = new Map<string, readonly string[]>();
  for (const lazyRootID of lazyRootIDs) {
    const lazyIDs = collectStatic(lazyRootID, initialIDs);
    lazyByRoot.set(lazyRootID, [...lazyIDs].sort());
  }
  return {
    initial: [...initialIDs].sort(),
    lazyByRoot,
  };
}

const buildModules: readonly BuildModule[] = [
  {
    dynamicImports: ["admin"],
    id: "main",
    staticImports: ["catalog"],
  },
  { dynamicImports: [], id: "catalog", staticImports: ["price"] },
  { dynamicImports: [], id: "price", staticImports: [] },
  { dynamicImports: [], id: "admin", staticImports: ["chart"] },
  { dynamicImports: [], id: "chart", staticImports: [] },
];

const firstLoadPlan = planFirstLoad(buildModules, "main");
const initialList = firstLoadPlan.initial.join(",");
const adminList = firstLoadPlan.lazyByRoot.get("admin")?.join(",");
console.assert(initialList === "catalog,main,price");
console.assert(adminList === "admin,chart");
```

#### Work and memory

For `v` reachable modules and `e` import edges, graph planning visits
each relevant module and edge a bounded number of times, so work is
`O(v + e)` and graph state is `O(v + e)`. Parsing, transforming,
minifying, and source-map generation add work proportional to source
size and tool behavior. Persistent build caches trade disk for less
repeated work and need versioned invalidation.

Runtime cost is different from build cost. Count compressed bytes,
requests, parse and compile time, execution, CSS discovery, and memory
per route. A lazy chunk reduces first-load work only if the initial
graph no longer imports it and the user does not immediately need it.

#### When to use it

Use static imports for code required by the entry and dynamic imports
at measured, user-meaningful boundaries. Use content hashes for
versioned assets, a lockfile for application installs, and a browser
target derived from the product's actual support policy.

Use bundle analysis when an entry grows, a dependency is duplicated,
or field performance shows excessive JavaScript work. Test the built
output whenever transforms, chunking, asset paths, or deployment
ordering can change behavior.

#### When not to use it

Do not add a bundler to a small standards-based page that can ship
native modules directly and meets its support requirements. Do not
split code that every user immediately needs. Do not transpile syntax
or add polyfills for browsers outside the declared support policy.

Do not publish a general-purpose library using application-only
resolution assumptions. Its emitted specifiers, export map, types,
side effects, and supported runtimes form a public compatibility
contract.

#### Common mistakes and edge cases

> **Warning:** A successful type check does not prove that the runtime
> can resolve the same import. TypeScript `paths`, package exports,
> bundler aliases, tests, and production must agree on one mapping.

- Types disappear and cannot validate untrusted runtime data.
- A syntax target does not provide missing browser APIs.
- A dynamic path the bundler cannot enumerate may omit needed files.
- Incorrect side-effect metadata can remove required registration.
- One common chunk can make an unrelated route download too much.
- Tiny chunks add request, header, scheduling, and evaluation cost.
- Changing one shared chunk can invalidate many dependent hashes.
- A lazy import needs loading, error, retry, and new-release behavior.
- Public source maps can reveal source even when code is minified.
- A floating install can build a graph that review never exercised.
- Development hot updates do not prove a clean production startup.
- Deleting old hashed files can break already-open application tabs.

#### Related material

- [JavaScript language and browser runtime](
  #a2-javascript-language-and-browser-runtime)
- [TypeScript type modeling and runtime boundaries](
  #a3-typescript-type-modeling-and-runtime-boundaries)
- [HTTP, caching, and network delivery][network]
- [Testing, performance, and observability][quality]
- [TypeScript module theory](
  https://www.typescriptlang.org/docs/handbook/modules/theory.html)
- [JavaScript modules](
  https://developer.mozilla.org/docs/Web/JavaScript/Guide/Modules)

[toolchain]: #a7-modules-toolchains-and-production-builds

## B. Product Application Engineering

### B.1 React Components, State, and Effects

#### The problem it solves

Interactive screens receive props, user events, server results, and
external subscriptions at different times. They must commit one UI
that corresponds to one coherent state snapshot. Poor ownership or
impure rendering creates stale values, state attached to the wrong
item, duplicate requests, and render loops.

#### How it works

Think of a component as a pure calculation from its current props,
state, and context to a description of UI. React coordinates that
calculation in separate phases:

1. A root render, state update, parent render, context update, or
   external-store notification **triggers** work.
2. During **render**, React calls components and compares the returned
   element tree with the preceding tree. Rendering must not change the
   DOM, start subscriptions, mutate shared objects, or depend on how
   many times it runs. React may restart or abandon render work.
3. During **commit**, React applies the accepted changes to the host
   environment, updates refs, and runs commit-related lifecycle work.
   Only a committed render becomes visible and owns effects.

This separation explains why a component function can run without a
DOM change. A parent render normally calls its children again, but the
commit changes only host output whose accepted values differ. Browser
style, layout, and paint happen after React's DOM work and have their
own costs.

React associates state with a component's **identity** in the render
tree. Identity comes from its component type, position among siblings,
and `key` when supplied. A stable key says which sibling is the same
logical item across insertion, deletion, and reordering. Keys are
local to one sibling list; they are not globally unique IDs. Changing
a key intentionally resets that subtree's state and DOM.

Each render receives a fixed **state snapshot**. Calling a state
setter queues work; it does not change variables or event handlers
that already belong to the current render. React batches compatible
updates and processes their queue for the next render. A value update
such as `setCount(count + 1)` queues a replacement computed from the
current snapshot. An updater such as
`setCount(currentCount => currentCount + 1)` receives the result of
the preceding queued update.

Treat objects and arrays in state as immutable snapshots. Construct a
replacement for changed data and preserve references for unchanged
data. Mutation can make old renders observe new content and can leave
memoized consumers with a reference that falsely appears unchanged.

Give each fact one owner:

- Keep temporary interaction state in the smallest subtree that uses
  it. Lift it to the closest common owner when siblings must agree.
- A **controlled** component receives its important value and change
  callback from its owner. An **uncontrolled** component owns that
  value internally, possibly exposing an initial value or ref.
- Calculate derived values during render instead of storing a second
  copy that can drift. Keep server state and URL state with their
  respective owners rather than mirroring them into local state.
- A ref stores mutable data that does not determine rendered output,
  such as a DOM handle or timer ID. Changing `ref.current` does not
  request a render.

Event handlers perform work caused by one interaction. An **Effect**
instead synchronizes a committed component with an external system:
a browser API, media object, timer, connection, subscription, or
third-party widget. The effect starts after commit. If a dependency
changes, React first runs the preceding cleanup, then starts the new
synchronization. Cleanup also runs when the component leaves the tree.

Every reactive value read by an effect belongs in its dependency
list. React compares dependencies with `Object.is`. Do not suppress
the dependency rule to control timing. Restructure the effect, move
interaction work into its event handler, move stable values outside
the component, or use an Effect Event for non-reactive logic that
must read current values. An Effect Event is not a way to hide a value
that truly must restart synchronization.

Passive effects should not perform a visual measurement that must
finish before paint. `useLayoutEffect` runs commit work before the
browser repaints and blocks that paint, so reserve it for necessary
layout measurement and synchronous correction. Effects do not run
during server rendering.

Development Strict Mode deliberately repeats selected pure functions
and performs an extra effect setup-cleanup-setup cycle. It exposes
render mutation and incomplete cleanup. Code should produce the same
user-visible result, not disable the check with a one-time ref.

Use a **reducer** when several events change related state through a
shared set of rules. A reducer is pure:
`current state + action -> next state`. Context can make a value
available below one provider, but it does not create an independent
store or prevent consumer renders.
Split contexts by change pattern and responsibility rather than
placing the whole application in one provider value.

For data owned outside React, `useSyncExternalStore` connects a stable
subscription contract to rendering. `subscribe` returns cleanup;
`getSnapshot` returns the same immutable snapshot while the store is
unchanged; and `getServerSnapshot`, when used, must agree between
server output and initial hydration. A custom hook can package this
contract, but it does not change ownership.

Memoization is a performance tool. `memo` can skip a component render
when props compare equal; `useMemo` retains a calculation result; and
`useCallback` retains a function identity. Each adds dependency,
comparison, and memory cost. When enabled, React Compiler can insert
memoization automatically. Manual or compiled memoization may be
dropped or recalculated, so correctness must never depend on it.

An Error Boundary contains errors thrown while rendering descendants
and replaces that region with a fallback. It does not catch ordinary
event-handler errors, arbitrary asynchronous callback errors, server
rendering errors, or its own errors. Suspense boundaries coordinate
supported pending work and `lazy` code, while transitions mark
replacement updates as non-urgent. These boundaries compose with the
rendering architectures described later; an effect fetch does not
activate Suspense by itself.

#### Step-by-step examples

First input: a counter render has `count = 0`. One click must add
three. Expected output: the next committed count is `3`.

1. **Before:** The handler belongs to the `count = 0` render.
   **Decision:** Calling `setCount(count + 1)` three times computes
   `1` three times and queues three replacement updates. **After:**
   The queue is `[replace 1, replace 1, replace 1]`.
   **Known correct:** None of those calls changed this handler's
   `count`, which remains `0`.
2. **Before:** The replacement queue would finish at `1`, not `3`.
   **Decision:** Queue the updater
   `currentCount => currentCount + 1` three times instead. **After:**
   The queue is `[increment, increment, increment]`.
   **Known correct:** Each updater will receive the preceding result.
3. **Before:** The old state is `0` and the updater queue is ready.
   **Decision:** Process it in order: `0 -> 1`, `1 -> 2`, then
   `2 -> 3`. **After:** Render receives `count = 3`; commit displays
   `3`. **Known correct:** One click produced one coherent final
   snapshot without three intermediate commits.

Second input: product `A` is visible; select product `B` while A's
refresh is pending. Expected output: B's loading state appears
immediately, A never appears under B's identity, and only B's result
becomes ready.

1. **Before:** Props name `B`, but retained request state still names
   `A`. **Decision:** During render, compare both IDs and return B's
   loading state when they differ. **After:** Commit cannot label A's
   data as B. **Known correct:** Display identity and data identity
   either match or the display is pending.
2. **Before:** The committed B render still has A's synchronization.
   **Decision:** Run A's cleanup, mark its closure ignored, abort its
   request, then start B's effect. **After:** B owns one controller
   and request state tagged `id = B`. **Known correct:** No active
   effect intentionally owns A.
3. **Before:** A finishes despite cancellation. **Decision:** Its
   closure sees `ignore = true` and discards the result. **After:** B
   remains loading. **Known correct:** Cooperative abort is not the
   only stale-result guard.
4. **Before:** B's request state is current and B succeeds.
   **Decision:** Store `{ id: B, status: ready, product: B }`.
   **After:** React renders and commits B's name. **Known correct:**
   The prop, request state, and product share identity B.

Identity boundary example: for rows `[A, B]`, keys `A` and `B` keep
each draft with its row when the list becomes `[B, A]`. Position keys
`0` and `1` preserve state at the positions instead, so A's draft can
appear in B's row. Changing a form key from `A` to `B` is appropriate
when that draft must intentionally reset for a new record.

#### Why it works

Each render and its handlers close over one fixed snapshot, so queued
updaters are the explicit way to combine changes that depend on prior
queued results. Pure rendering lets React repeat or abandon that
calculation without leaking partial external work. Only commit changes
the host output.

Type, position, and key give React a stable address for each stateful
subtree. Keeping the same address preserves state; changing it resets
state. Immutable replacements make meaningful changes visible through
reference comparison without altering snapshots held by older renders.

Effects belong to committed inputs. Cleanup revokes the preceding
synchronization before its replacement begins. In the request example,
the render-time ID check prevents a stale display, while the closure's
ignore flag prevents a stale publication even when transport abort is
too late. The stored ID then proves that ready data belongs to the
current selection.

#### TypeScript example

`loadProduct` accepts an abort signal, but correctness also uses an
ignore flag because cancellation may arrive after work completes.

```tsx
type Product = { id: string; name: string };

type ProductResult =
  | { id: string; status: "loading" }
  | { id: string; message: string; status: "error" }
  | { id: string; product: Product; status: "ready" };

function ProductDetails({ id }: { id: string }) {
  const [result, setResult] = React.useState<ProductResult>({
    id,
    status: "loading",
  });

  // Each committed product identity owns one synchronization cycle.
  React.useEffect(() => {
    const controller = new AbortController();
    let ignore = false;
    setResult({ id, status: "loading" });

    // Publish only while this exact committed identity still owns the
    // effect, even if transport cancellation finishes too late.
    loadProduct(id, controller.signal)
      .then((product) => {
        if (!ignore) {
          setResult({ id, product, status: "ready" });
        }
      })
      .catch((reason: unknown) => {
        if (!ignore && !controller.signal.aborted) {
          let message = "failed";
          if (reason instanceof Error) {
            message = reason.message;
          }
          setResult({ id, message, status: "error" });
        }
      });

    // Cleanup revokes publication before requesting transport abort.
    return () => {
      ignore = true;
      controller.abort();
    };
  }, [id]);

  // A new prop can render before its effect resets retained state.
  // Never display data whose stored identity differs from the prop.
  if (result.id !== id || result.status === "loading") {
    return <p>Loading</p>;
  }
  if (result.status === "error") {
    return <p role="alert">{result.message}</p>;
  }
  return <h2>{result.product.name}</h2>;
}
```

#### Work and memory

Rendering visits the component work selected for an update and the
descendants that are not skipped. Reconciliation cost depends on the
compared element subtrees and their keys. Commit cost depends on host
changes; browser layout and paint depend on the resulting DOM and CSS.
A render that produces no DOM change still consumed JavaScript work.

State, refs, context values, memo caches, retained closures, pending
updates, and external resources consume memory for the lifetime of
their owners. An effect may temporarily overlap transport work with
its replacement even though only the current closure may publish.
Cleanup bounds listeners, connections, timers, and retained objects.

Context notification can visit many consumers. External stores add
subscription and snapshot work. Memoization adds comparisons and
retained values; it helps only when measured avoided work costs more.
Virtualization is usually more effective than memoizing thousands of
offscreen rows that should not be mounted.

#### When to use it

Use local state for interaction owned by one subtree and lift it when
siblings require one source of truth. Use a reducer when many events
share transition rules. Use context for stable cross-cutting values or
a scoped state owner, and `useSyncExternalStore` for a mutable source
that exists independently of React.

Use effects for synchronization with external systems and handlers for
work caused by a particular interaction. Use stable semantic keys for
lists, and change a key when product identity requires a deliberate
subtree reset.

Measure before adding manual memoization. Adopt React Compiler through
the supported build integration and validate behavior and performance
incrementally rather than removing existing memoization wholesale.

#### When not to use it

Do not copy props into state merely to calculate display output. Do
not use an effect to derive state that render can calculate or to run
logic owned by one event. Do not place server-cache behavior in
scattered component effects when a route or server-data owner can
preload, deduplicate, cache, and invalidate it.

Do not put every value in context or one global store. Local state is
more explicit when no distant consumer needs it. Do not use a ref for
data that must update the screen, and do not use memoization to hold a
semantic value that must never be forgotten.

#### Common mistakes and edge cases

> **Warning:** Omitting an effect dependency does not freeze a value.
> It makes the effect retain an older render's value and can publish
> stale work. Include the dependency or change the synchronization
> design so that value is not reactive.

- Calling Hooks conditionally changes their call order across renders.
- I/O or shared mutation during render leaks abandoned work.
- Three snapshot-based replacements do not equal three updater calls.
- Array indexes are unsafe keys when items reorder or are inserted.
- Defining a component inside another component resets its identity.
- Mutating state can change old snapshots and preserve a stale
  reference.
- Copying props into state creates two owners and unclear reset rules.
- A controlled input must not switch accidentally to uncontrolled.
- An effect that sets one of its own dependencies can loop.
- A missing cleanup duplicates listeners under development remounts.
- An abort signal alone does not prevent every late Promise callback.
- `useLayoutEffect` can block paint and does not run on the server.
- A new provider object can notify every observing context consumer.
- `getSnapshot` must cache identity while its store is unchanged.
- Error Boundaries do not catch ordinary async or handler exceptions.
- Suspense does not detect a request started in a passive effect.
- Custom prop comparisons can retain stale function closures.
- Memoization that changes correctness is incorrect without the cache.

#### Related material

- [JavaScript language and browser runtime][runtime]
- [TypeScript type modeling and runtime boundaries](
  #a3-typescript-type-modeling-and-runtime-boundaries)
- [Server data, routing, caching, and forms][data]
- [Rendering architectures, hydration, and Server Components](
  #b4-rendering-architectures-hydration-and-server-components)
- [Testing, performance, and observability][quality]
- [React render and commit](https://react.dev/learn/render-and-commit)
- [React state as a snapshot](
  https://react.dev/learn/state-as-a-snapshot)
- [React external-store reference](
  https://react.dev/reference/react/useSyncExternalStore)
- [React Compiler](https://react.dev/learn/react-compiler)

[runtime]: #a2-javascript-language-and-browser-runtime

<a id="b2-server-data-routing-rendering-and-forms"></a>

### B.2 Server Data, Routing, Caching, and Forms

#### The problem it solves

A product interface must coordinate URLs, overlapping reads, cached
server snapshots, user drafts, writes, and navigation history. It must
remain truthful during loading, validation failure, uncertain writes,
back navigation, account changes, and stale responses without losing a
draft or showing one resource under another resource's identity.

#### How it works

Start by assigning every fact one authoritative owner:

- The URL owns shareable navigation state such as route identity,
  search, filters, sort order, pagination position, and selected tabs
  that should survive reload or back and forward traversal.
- The server owns confirmed records, authorization, invariants, and
  the ordering of concurrent writes.
- A server-data cache owns local snapshots of server records, their
  freshness, in-flight reads, and invalidation policy.
- A form owns an unconfirmed draft, dirty and touched fields, client
  feedback, pending submission, and server validation errors.
- A component owns temporary interaction that is neither shareable
  nor durable, such as an open disclosure or pointer hover.

The browser URL is untrusted text. Match the path, decode parameters,
reject duplicate or malformed values, apply documented defaults, and
produce one canonical application value before loading data. If two
URLs mean the same state, redirect or replace one with the canonical
form so history, analytics, cache keys, and sharing agree.

Use a history **push** when the user makes a meaningful navigation
they may want to undo. Use **replace** for normalization or a
transient edit that should not create another back-button stop. Back
and forward navigation restores a prior URL; derive the screen from
that URL instead of keeping a second route state that can disagree.

A route transition is a transaction, not just a component render. It
may match nested routes, validate parameters, start independent reads
in parallel, authorize on the server, load code, update the document
title, restore or reset scroll, and move focus. Give the transition an
identity and cancellation signal. Cancellation revokes the client's
interest; it does not roll back a completed server write.

A server-data cache uses a **query key** as a lookup contract:

`complete resource identity -> cached server snapshot`

For example, the key
`[products, account-7, books, 20, cursor-c17]` means the
twenty-product page after cursor `c17` for the books view visible to
account 7. Include every input that can change the representation:
resource ID, account or tenant when relevant, locale, filters, sort,
pagination, feature variant, and API version. A key isolates cached
data; it does not authorize the read.

Keep these cache times distinct:

- **Freshness** is how long stored data may be reused without a read.
- **Staleness** means a refresh is permitted or required; stale data
  may still be useful while that refresh is pending.
- **Retention** is how long an unused entry stays available before
  local eviction. It is not server or HTTP cache freshness.
- **Invalidation** marks knowledge as outdated after an event such as
  a write. It may trigger a read for active observers or wait until
  the next use, depending on policy.

For one key, a cache can share one in-flight Promise among observers
instead of issuing duplicate reads. It can retain confirmed data while
refreshing, prefetch a likely next route, retry bounded transient read
failures, and cancel work with no remaining consumer. Each completed
read must still match the active key and request instance before it
can publish. Transport abort alone cannot establish that fact.

Offset pagination is simple but shifts when earlier records are added
or removed. Cursor pagination asks for records after a stable ordering
position. It needs a deterministic sort with a unique tie-breaker and
an opaque cursor interpreted only by the server. The URL may own the
current cursor, while the cache owns each returned page and next
cursor. Do not invent a missing page or silently display page 1 under
page 2's URL.

A cache can store whole query results or normalize records by entity
identity. Query-key caches make request ownership and invalidation
direct. Normalized caches can update one entity across many views but
need schema identity, merge, deletion, authorization, and pagination
rules. Choose one deliberate owner instead of copying records between
several caches and a global store.

A **mutation** requests a server change. Model at least idle, pending,
confirmed, rejected, and uncertain outcomes. Disable or coalesce a
duplicate action when only one logical operation is allowed, but keep
the interface accessible and explain the pending state. For a write
that may be retried after a lost response, reuse one stable operation
identity so the server can return the original result.

After confirmation, either replace a cache entry with the complete
canonical response or invalidate every affected query so its owner can
refresh according to policy. Do not guess partial server behavior when
triggers, permissions, rounding, versions, or concurrent writers can
change the result.

An **optimistic update** displays an expected result before the server
confirms it. Use one only when the result is predictable and failure
has a clear user experience. Record the confirmed base or inverse
operation, prevent a background read from overwriting the optimistic
state, identify concurrent optimistic writes separately, roll back or
mark rejection, then reconcile with the canonical response. Showing a
pending draft beside confirmed data is often simpler than rewriting a
shared cache optimistically.

Use native form semantics even when JavaScript enhances submission. A
`form` groups controls and supports keyboard submission. Labels name
controls; submit buttons use the intended `type`; `FormData` includes
successful named controls. Disabled controls and unchecked checkboxes
are omitted, and file inputs yield `File` values. Parse every entry as
untrusted input and place limits on text and files.

Client validation gives immediate guidance but cannot protect the
server. The server repeats validation and returns stable field errors
plus a form-level error when needed. Keep the user's draft on known
failure, connect errors to their controls, announce the result, and
move focus only when that improves recovery. Reset a draft only after
confirmed success or an explicit cancel.

React form Actions can wrap a function submission in a Transition.
`useActionState` exposes the returned state and pending status;
`useFormStatus` reads the surrounding form submission; and
`useOptimistic` provides temporary Action state. These APIs coordinate
UI. A Server Function is still a public endpoint that must
authenticate, authorize, validate, limit, and define replay behavior.
Framework support determines whether a form progressively works
before hydration.

On sign-out or account switch, cancel private reads, remove private
browser cache entries, clear drafts that must not cross accounts, and
replace account-owned route context. Server authorization remains the
final boundary even when client keys include account identity.

#### Step-by-step examples

First input: account `acct-7` opens
`/products?category=books&limit=20&after=c17`. Before its request
finishes, the user navigates to cursor `c42`, then back to `c17`.
Expected output: only the newest `c17` request may publish, and the
URL remains sufficient to reconstruct the view.

1. **Before:** The URL holds untrusted strings. **Decision:** Require
   one value per parameter, accept category `books`, parse bounded
   integer limit `20`, and preserve opaque cursor `c17`. **After:**
   The query is `{ category: books, limit: 20, after: c17 }`.
   **Known correct:** Ambiguous or invalid inputs cannot enter a key.
2. **Before:** Account and validated query are known. **Decision:**
   Create key `[products, acct-7, books, 20, c17]` and navigation
   sequence `41`. **After:** The cache lookup and request have
   complete identity. **Known correct:** Another account, filter,
   limit, or cursor cannot share this entry.
3. **Before:** Sequence 41 is pending and navigation moves to `c42`.
   **Decision:** Start sequence 42 with its distinct key and revoke
   interest in 41. **After:** The URL and active request describe
   `c42`. **Known correct:** A late sequence 41 result cannot match
   the active key or sequence.
4. **Before:** Back navigation restores `c17` while request 41 may
   still finish. **Decision:** Start sequence 43 using the original
   `c17` key. **After:** Sequence 43 owns the active request.
   **Known correct:** The key alone is insufficient; sequence 41 and
   sequence 43 refer to different attempts for the same view.
5. **Before:** Sequence 41 completes, then sequence 43 completes.
   **Decision:** Reject 41 because its sequence is stale; accept 43
   because both key and sequence match. **After:** The displayed page
   and URL both describe the newest `c17` snapshot. **Known correct:**
   Completion order cannot overwrite current navigation intent.

Second input: submit quantity `2` for order draft `draft-9`; the first
response is lost after the server may have created order `o31`.
Expected output: the draft is not duplicated, and the UI reaches one
confirmed order or an explicit unresolved state.

1. **Before:** The form owns draft quantity `2`. **Decision:** Parse
   `FormData`, run client guidance, and submit after server validation
   with operation ID `place-draft-9`. **After:** The form is pending
   and retains its draft. **Known correct:** One logical write has one
   replay identity.
2. **Before:** The connection ends without a usable response.
   **Decision:** Mark the outcome uncertain rather than rejected.
   **After:** The UI offers a status check or retry using the same
   operation ID. **Known correct:** Absence of a response does not
   imply absence of a server commit.
3. **Before:** Retry reaches the server with `place-draft-9`.
   **Decision:** Return the recorded order `o31` instead of creating a
   second order. **After:** The client stores the confirmed response,
   invalidates affected order queries, and resets the draft.
   **Known correct:** Recovery converges on one server result.

Validation boundary: quantity `0` may fail immediately for guidance,
but the server still rejects it independently. A server field error
keeps the original draft and associates the message with the quantity
control instead of clearing the form.

#### Why it works

URL parsing creates one canonical navigation value before any lookup.
The complete query key then selects exactly one server snapshot, while
the request sequence selects the newest attempt allowed to publish for
that view. Checking both prevents cross-resource collisions and
same-resource completion races.

Drafts remain separate from confirmed snapshots, so background reads
cannot erase unsubmitted work. A mutation crosses that boundary only
after server validation. Stable operation identity lets an uncertain
retry name the same logical write, and canonical confirmation or
invalidation returns the cache to server-owned truth.

#### TypeScript example

The key contract is
`account and validated query -> one cached product page`. The request
sequence separately identifies which attempt may publish.

```typescript
const productCategories = ["all", "books", "office"] as const;

type ProductCategory = (typeof productCategories)[number];

type ProductQuery = Readonly<{
  after: string | null;
  category: ProductCategory;
  limit: number;
}>;

type ProductKey = readonly [
  resource: "products",
  accountID: string,
  category: ProductCategory,
  limit: number,
  after: string | null,
];

type PageRequest = Readonly<{
  key: ProductKey;
  sequence: number;
}>;

type OptionalText = string | null;

function readSingle(
  searchParameters: URLSearchParams,
  parameterName: string,
): OptionalText {
  const values = searchParameters.getAll(parameterName);
  if (values.length > 1) {
    throw new Error(`${parameterName} must appear at most once`);
  }
  return values[0] ?? null;
}

function isProductCategory(value: string): value is ProductCategory {
  return productCategories.includes(value as ProductCategory);
}

function parseProductQuery(url: URL): ProductQuery {
  // Normalize every URL-owned value before it enters cache identity.
  const category = readSingle(url.searchParams, "category") ?? "all";
  if (!isProductCategory(category)) {
    throw new Error("unsupported category");
  }

  const rawLimit = readSingle(url.searchParams, "limit") ?? "20";
  const limit = Number(rawLimit);
  if (!Number.isInteger(limit) || limit < 1 || limit > 100) {
    throw new Error("limit must be an integer from 1 through 100");
  }

  const rawAfter = readSingle(url.searchParams, "after");
  // Empty means no cursor. A non-empty cursor is opaque server data
  // and must pass through without client interpretation.
  const after = rawAfter === "" ? null : rawAfter;
  return { after, category, limit };
}

function productKey(
  accountIdentity: string,
  validatedQuery: ProductQuery,
): ProductKey {
  // Account identity isolates private snapshots but does not grant
  // authorization; the server checks that on every request.
  return [
    "products",
    accountIdentity,
    validatedQuery.category,
    validatedQuery.limit,
    validatedQuery.after,
  ];
}

function sameProductKey(
  expectedKey: ProductKey,
  completedKey: ProductKey,
): boolean {
  return expectedKey.every((value, index) => {
    return value === completedKey[index];
  });
}

function canPublish(
  activeRequest: PageRequest,
  completedRequest: PageRequest,
): boolean {
  // The key proves resource identity. The sequence proves that this
  // completion belongs to the newest attempt for that identity.
  return (
    activeRequest.sequence === completedRequest.sequence &&
    sameProductKey(activeRequest.key, completedRequest.key)
  );
}

const productsURL = new URL("https://example.test/products");
productsURL.search = "?category=books&limit=20&after=c17";
const query = parseProductQuery(productsURL);
const key = productKey("acct-7", query);
const activeRequest = { key, sequence: 43 };
const staleRequest = { key, sequence: 41 };
const expectedKey = "products:acct-7:books:20:c17";
console.assert(key.join(":") === expectedKey);
console.assert(!canPublish(activeRequest, staleRequest));
console.assert(canPublish(activeRequest, activeRequest));
```

#### Work and memory

If a URL has `u` characters and `p` parameters, parsing uses work and
temporary storage proportional to those inputs. The fixed-width key in
the example compares in constant work. A general structured key costs
work proportional to its serialized fields. A cache lookup is usually
constant average map work after key construction.

If a returned page has `n` records and `b` transferred bytes, runtime
validation, normalization, and rendering add work proportional to the
visited records or bytes. Retained cache memory is the sum of live
entries, indexes, metadata, pending Promises, and optimistic layers
until eviction. Prefetching trades more network, server, and memory
work for a possible reduction in later wait time.

A controlled form may render on each input update. Uncontrolled fields
leave ordinary value storage in the DOM until submission. `FormData`
and validation visit successful controls and may retain `File`
objects; limit file count and size before expensive processing.

#### When to use it

Use URL state when reload, sharing, history, or deep linking should
reconstruct the selection. Use a route loader or framework data API
when reads belong to navigation and can start before component
effects.

Use a server-data cache when reads need deduplication, freshness,
invalidation, prefetching, retries, or reuse across observers. Use
form state for drafts and field feedback. Use pessimistic
confirmation for consequential writes and optimistic display for
predictable, reversible interactions where latency would harm use.

#### When not to use it

Do not place hover, focus, or a one-use disclosure in the URL. Do not
copy every server record into a global client store or layer a second
cache over a framework cache without assigning invalidation ownership.

Do not retry a non-replay-safe write automatically. Do not use an
optimistic update when rollback is ambiguous, the operation is
high-impact, or the server result depends on hidden concurrent rules.
Do not normalize a small independent query cache merely because large
applications sometimes use normalized entity stores.

#### Common mistakes and edge cases

> **Warning:** A cache key missing `userId`, locale, filter, or page
> can show one request's data in another request's view. Include every
> representation input, then enforce authorization independently.

- Duplicate search parameters can make client and server parse
  different values.
- A key built before validation gives malformed input cache life.
- A cache key does not prevent an unauthorized network request.
- Freshness, HTTP caching, and local retention are different policies.
- A late response can be stale even when its query key matches again.
- Offset pages can skip or repeat rows after concurrent insertion.
- Cursor pagination fails without stable ordering and a tie-breaker.
- A background refresh must not overwrite a dirty form draft.
- Disabled successful controls are omitted from `FormData`.
- An unchecked checkbox is absent rather than the string `false`.
- Client validation is guidance, never the server's trust boundary.
- Clearing a failed form destroys the values needed for recovery.
- A loading flag alone cannot distinguish rejection from uncertainty.
- A new operation key on retry can duplicate a committed write.
- Optimistic rollback can erase a later concurrent optimistic change.
- Server Functions require endpoint-grade authorization and limits.
- Redirect loops arise when route guards disagree about identity.
- Hydration fails when server and first client route data differ.
- Logout must remove private cache entries, not only hide the screen.

#### Related material

- [HTTP, caching, and network delivery][network]
- [Browser storage, origins, and security][security]
- [React components, state, and
  effects](#b1-react-components-state-and-effects)
- [Rendering architectures, hydration, and Server Components](
  #b4-rendering-architectures-hydration-and-server-components)
- [Backend HTTP contracts][http]
- [HTML forms](https://html.spec.whatwg.org/multipage/forms.html)
- [React form Actions](
  https://react.dev/reference/react-dom/components/form)
- [React action state](
  https://react.dev/reference/react/useActionState)
- [TanStack Query optimistic updates](
  https://tanstack.com/query/latest)

[http]: ../backend/README.md

### B.3 Testing, Performance, and Frontend Observability

#### The problem it solves

A frontend can pass local checks while failing through timing,
accessibility, browser differences, slow devices, large data, or a
production-only dependency. Teams need evidence before release,
explicit performance limits, and privacy-safe field signals that show
which users and releases fail afterward.

#### How it works

Treat quality as an evidence pipeline with three different jobs:

1. **Automated and manual tests** challenge known contracts in a
   controlled environment before release.
2. **Performance measurement and budgets** quantify experience and
   fail a build or alert when an agreed limit is crossed.
3. **Observability** records bounded production evidence so a team can
   detect, group, diagnose, and verify real failures after release.

No one layer proves the others. A passing browser test does not prove
fast field experience. A good metric does not prove checkout creates
the correct order. An error report without a tested recovery path does
not prove that users can finish.

Start testing from user and system risks, then choose the narrowest
environment that can observe each contract:

- Static checks catch type, syntax, dependency, and rule violations
  without running the interface.
- Unit tests exercise pure parsing, reducers, formatters, validation,
  and state transitions with exact ordinary and boundary inputs.
- Component tests render a meaningful UI region with its real child
  behavior. Query by role, accessible name, label, text, or another
  user-visible contract rather than a class or private state value.
- Contract and integration tests verify that adjacent boundaries agree
  on request, response, error, authorization, and version shapes.
- End-to-end tests use a real browser and controlled backend state for
  critical workflows such as sign-in, checkout, and uncertain retry.
- Visual tests detect unintended pixels under a controlled browser,
  font, data, viewport, color scheme, and animation policy.
- Manual checks cover judgment automation cannot supply: keyboard and
  screen-reader flow, zoom and reflow, understandable errors, focus,
  touch use, motion, and real assistive-technology combinations.

Test behavior, not the framework's implementation. A checkout test
activates the button named `Place order`, observes pending and error
status, preserves the draft, retries, and sees one confirmation. It
does not assert a hook name, component method, CSS class, or exact
internal request count unless that count is itself a product contract.

Give every test its own known state. Seed only the data it needs. Use
a fresh browser context or equivalent storage boundary. Control time,
randomness, locale, network responses, and feature configuration when
they affect the assertion. A test must pass alone, in another order,
and under supported parallel execution. Cleanup is less reliable than
an isolated start because cleanup can miss hidden state.

Asynchronous tests wait for an observable condition, not a guessed
duration. Use a semantic locator that can be resolved again and an
auto-retrying assertion with a bounded timeout. Fixed sleeps make fast
systems slow and slow systems flaky. A timeout is the final bound, not
the synchronization mechanism.

Mock at a boundary the test does not own. A component test may replace
the network with a controlled HTTP response while keeping rendering,
events, and accessibility semantics real. An end-to-end test should
exercise the owned deployed stack but isolate unstable third parties.
If a mock reproduces the implementation's private calls, it can make a
broken contract pass by agreeing with the same mistake.

Accessibility automation finds selected structural failures but
cannot prove accessible use. Combine semantic queries, an automated
scan, keyboard-only operation, focus inspection, zoom and reflow, and
targeted screen-reader checks. Test loading, empty, partial, success,
validation, forbidden, expired-session, offline, and failed states,
not only the ready screen.

A **performance budget** converts an expectation into a limit with an
owner and enforcement point. Budget transferred and decoded bytes,
request count, JavaScript parse and execution, long main-thread work,
rendering, memory where measurable, and user-centered timings. Define
route and device profiles; one global number hides expensive routes
and low-end devices.

Laboratory measurement runs a controlled device, network, cache state,
browser, and data fixture. It is repeatable and useful for diagnosis
and release comparison. Synthetic monitoring repeats a scripted route
from managed locations. **Real-user monitoring (RUM)** samples actual
supported users, devices, networks, cache states, and interactions.
Field data proves distribution; laboratory traces help explain cause.

The current Core Web Vitals are:

- **Largest Contentful Paint (LCP)** measures loading. Its good
  threshold is at most `2.5s`.
- **Interaction to Next Paint (INP)** for responsiveness, with a good
  threshold of at most `200ms`.
- **Cumulative Layout Shift (CLS)** for unexpected visual movement,
  with a good threshold of at most `0.1`.

Evaluate these at the 75th percentile and segment mobile and desktop.
These are current shared thresholds, not a complete product budget.
Also measure server response, first content, route transitions, task
or animation-frame blocking, memory, and business actions such as time
to search results or checkout confirmation. Refresh definitions
against the linked source when the technology snapshot changes.

Do not average away slow users. Report percentiles and sample counts.
Then segment by route template, release, device class, browser,
network class, lawful and useful region, and feature cohort. Set a
minimum sample rule. A four-sample example can show percentile math;
it cannot justify a production rollout decision.

Use browser performance tools to move from symptom to cause. A trace
can relate input delay, JavaScript tasks, style calculation, layout,
paint, network waterfalls, and memory. Resource Timing exposes fetched
resource phases subject to cross-origin timing policy.
`PerformanceObserver` delivers supported timeline entries, including
buffered entries when requested. Use a maintained Web Vitals library
for evolving metric definitions instead of approximating INP by hand.

Frontend observability uses several signal types:

- Errors describe an unexpected failure, stack, cause, and handled or
  unhandled boundary. Expected validation is product state, not an
  exception alert.
- Metrics aggregate bounded numeric behavior such as error rate,
  completion rate, LCP, INP, CLS, and route-transition duration.
- Events record a small product or lifecycle fact under a versioned
  schema. They are not arbitrary console logs.
- Traces connect work across navigation, browser operations, API
  requests, and backend services through permitted correlation IDs.

Attach enough stable context to act: release and build ID, route
template rather than raw private URL, operation type, coarse runtime
cohort, feature cohort, and a random correlation ID. Keep user IDs,
search terms, form values, tokens, headers, full URLs, DOM text, and
payment or health data out unless an explicit lawful design requires
and protects them. Redact before data enters a queue, not only at the
server.

Bound event size, queue length, flush time, retries, sampling, and
retention. Browser delivery is best effort; navigation, offline state,
privacy controls, blockers, and crashes can lose telemetry. Never make
a user workflow depend on reporting success. Protect source maps and
match them to the exact release so generated stacks map correctly.

Define an observable user outcome, not merely an absence of thrown
errors. A checkout success rate needs a valid start and one confirmed
completion; abandonment, forbidden responses, duplicate orders, and
uncertain outcomes need distinct meanings. Alert on sustained user
impact or error-budget consumption, then link the alert to a runbook,
release comparison, owner, and rollback rule.

#### Step-by-step examples

First input: test checkout with valid draft `draft-9`. The controlled
server returns `503` once, then returns recorded order `o31` when the
same operation identity is retried. Expected output: pending, an
accessible recoverable error with the draft intact, and one confirmed
order after retry.

1. **Before:** The pure submission state is `ready`. **Decision:** A
   unit test applies `submit`, `reject`, `submit`, then `resolve`.
   **After:** States are `pending`, `failed`, `pending`, `complete`.
   **Known correct:** Invalid events and repeated submits cannot
   bypass the state contract.
2. **Before:** The component has draft quantity `2` and the controlled
   HTTP boundary will return `503`. **Decision:** Activate the button
   by role and name, then await visible pending and error states.
   **After:** The error is announced, retry is enabled, and quantity
   remains `2`. **Known correct:** Rendering and accessibility expose
   the same recoverable state as the pure transition.
3. **Before:** An isolated browser starts with the same fixture and
   operation ID `place-draft-9`. **Decision:** Exercise the complete
   retry workflow against controlled backend state. **After:** The URL
   shows confirmation `o31`, and the backend fixture contains one
   order. **Known correct:** Browser, API, and replay contracts agree.
4. **Before:** The test passes at the default viewport. **Decision:**
   repeat the critical behavior with keyboard input, narrow reflow,
   supported browser projects, and reduced motion. **After:** The same
   outcome remains operable. **Known correct:** One happy desktop path
   is not the only evidence for the owned risk.

Second input: candidate release `r42` has four mobile checkout
samples. LCP values are `[1800, 2100, 2300, 2600]ms`, INP values are
`[130, 170, 260, 310]ms`, and CLS values are
`[0.02, 0.05, 0.08, 0.12]`. Expected output: the teaching evaluator
passes LCP and CLS but fails INP because its 75th percentile is
`260ms`, above the `200ms` threshold.

1. **Before:** The LCP samples are unsorted. **Decision:** Sort and
   select nearest rank `ceil(0.75 * 4) = 3`. **After:** LCP p75 is
   `2300ms`. **Known correct:** It passes the `2500ms` threshold.
2. **Before:** INP samples are unsorted. **Decision:** Apply the same
   rank to `[130, 170, 260, 310]`. **After:** INP p75 is `260ms`.
   **Known correct:** The candidate fails responsiveness even though
   its fastest half looks healthy.
3. **Before:** CLS samples are unsorted. **Decision:** Select the
   third sorted value. **After:** CLS p75 is `0.08`.
   **Known correct:** It passes the `0.1` threshold.
4. **Before:** The aggregate points to INP but not cause.
   **Decision:** Filter field evidence by release, route template, and
   mobile cohort, then reproduce it under a laboratory trace.
   **After:** A long frame shows the checkout update rendering an
   unnecessary large list.
   **Known correct:** The metric found affected experience; the trace
   supplies a testable cause hypothesis.
5. **Before:** The list update is corrected and relevant component
   and browser tests pass. **Decision:** Rerun the lab budget, stage
   rollout, and compare the candidate field cohort. **After:** The
   team verifies recovery or rolls back by its prior threshold.
   **Known correct:** A code change is not successful until user
   evidence improves.

Privacy boundary: if a failing URL contains `?email=a@example.test`,
record route template `/checkout`, release `r42`, and a random
correlation ID. Do not send the raw URL, email, form draft, token, or
DOM text to make diagnosis easier.

#### Why it works

Each pre-release test observes the smallest real boundary that owns a
risk. Isolation makes failures attributable to that test, while
semantic actions and eventual assertions synchronize on the same
states users perceive. Deliberately producing the failure proves the
recovery branch rather than merely executing the happy path.

A budget turns a performance goal into a falsifiable decision. Field
percentiles retain the slow portion of a cohort that an average can
hide. Segment and release identity connect the regression to affected
experience. A trace then decomposes time into changeable work.

Bounded, versioned telemetry keeps evidence comparable and limits
resource and privacy risk. Correlation IDs connect permitted events
without embedding user content. Testing the fix, measuring it in a
controlled environment, and verifying it in the field closes the loop
from known contract through real outcome.

#### TypeScript example

This teaching evaluator shows the percentile and threshold decision.
Production evaluation also needs sufficient samples, cohort rules,
metric attribution, and a maintained Web Vitals implementation.

```typescript
const vitalNames = ["LCP", "INP", "CLS"] as const;

type VitalName = (typeof vitalNames)[number];

type VitalSamples = Readonly<Record<VitalName, readonly number[]>>;

type VitalResult = Readonly<{
  name: VitalName;
  passes: boolean;
  percentile75: number;
  threshold: number;
}>;

type VitalResults = readonly VitalResult[];

const goodThresholdByVital: Readonly<Record<VitalName, number>> = {
  CLS: 0.1,
  INP: 200,
  LCP: 2500,
};

function compareNumbers(left: number, right: number): number {
  return left - right;
}

function nearestRankPercentile(
  samples: readonly number[],
  percentile: number,
): number {
  if (samples.length === 0) {
    throw new Error("percentile requires at least one sample");
  }
  const percentileIsFinite = Number.isFinite(percentile);
  const percentileIsInRange = percentile > 0 && percentile <= 1;
  if (!percentileIsFinite || !percentileIsInRange) {
    throw new Error("percentile must be greater than 0 through 1");
  }
  const hasInvalidSample = samples.some((sample) => {
    return !Number.isFinite(sample) || sample < 0;
  });
  if (hasInvalidSample) {
    throw new Error("samples must be finite non-negative numbers");
  }

  // Sorting a copy preserves the caller's chronological samples.
  const sortedSamples = [...samples].sort(compareNumbers);
  const oneBasedRank = Math.ceil(percentile * sortedSamples.length);
  const zeroBasedIndex = oneBasedRank - 1;
  const selectedSample = sortedSamples[zeroBasedIndex];
  if (selectedSample === undefined) {
    throw new Error("percentile rank is outside the sample set");
  }
  return selectedSample;
}

function evaluateVitals(samples: VitalSamples): VitalResults {
  // Apply one explicit policy to every current Core Web Vital.
  return vitalNames.map((name) => {
    const percentile75 = nearestRankPercentile(samples[name], 0.75);
    const threshold = goodThresholdByVital[name];
    return {
      name,
      passes: percentile75 <= threshold,
      percentile75,
      threshold,
    };
  });
}

const candidateSamples: VitalSamples = {
  CLS: [0.02, 0.05, 0.08, 0.12],
  INP: [130, 170, 260, 310],
  LCP: [1800, 2100, 2300, 2600],
};

const candidateResults = evaluateVitals(candidateSamples);
const resultEntries = candidateResults.map((result) => {
  return [result.name, result] as const;
});
const resultByName = new Map(resultEntries);
console.assert(resultByName.get("LCP")?.percentile75 === 2300);
console.assert(resultByName.get("LCP")?.passes === true);
console.assert(resultByName.get("INP")?.percentile75 === 260);
console.assert(resultByName.get("INP")?.passes === false);
console.assert(resultByName.get("CLS")?.percentile75 === 0.08);
console.assert(resultByName.get("CLS")?.passes === true);
```

#### Work and memory

Test cost includes environment startup, fixture creation, browser and
database processes, network control, traces, screenshots, retries, and
retained artifacts. Unit checks are cheap enough to run broadly;
browser matrices are expensive and should cover risk deliberately.
Parallel work reduces elapsed time but increases CPU, memory, ports,
accounts, and fixture-isolation requirements.

For `n` metric samples, the example copies and sorts each collection,
using `O(n log n)` comparison work and `O(n)` additional memory.
Production percentile systems normally aggregate or sketch bounded
distributions instead of retaining every sample in one browser.

Measurement itself consumes main-thread work, memory, and network.
Bound observation, serialize small schemas, batch responsibly, sample
before expensive enrichment, and drop rather than grow an unbounded
offline queue. Debug artifacts and source maps need retention and
access policies because they can contain sensitive implementation or
page data.

#### When to use it

Use the smallest test that can prove the owned risk, then add a larger
boundary only when the risk crosses it. Use real browsers for browser
contracts and critical workflows, contract tests for separately owned
APIs, and manual accessibility checks for interaction judgment.

Use laboratory traces to reproduce and diagnose, synthetic checks for
continuous controlled availability, and RUM for claims about actual
supported users. Enforce budgets in continuous integration or release
policy, and compare staged rollout cohorts before broad exposure.

#### When not to use it

Do not make every branch an end-to-end test or replace behavior tests
with snapshots. Do not test a third-party service's content from the
critical suite; control that boundary and separately monitor the owned
integration contract.

Do not use a laboratory score as proof of field experience, RUM as a
substitute for reproducible diagnosis, or exception counts as a user
success measure. Do not collect private content merely because it may
help a future investigation.

#### Common mistakes and edge cases

> **Warning:** Waiting for a fixed time makes an asynchronous test
> both slow and unreliable. Wait for the user-observable state
> through an eventual assertion with a bounded timeout instead.

- A selector tied to a CSS class tests implementation, not use.
- Shared accounts, cookies, clocks, or rows make test order matter.
- Retrying a flaky test can hide a race unless the retry is diagnosed.
- A mocked client and mocked server can agree on the same wrong shape.
- Snapshot approval can normalize an unintended accessibility failure.
- Automated accessibility scans do not operate a screen reader.
- Code coverage shows execution, not useful assertions or boundaries.
- A warm-cache desktop lab run does not model a cold mobile visit.
- Average latency hides the slow tail and affected cohorts.
- A p75 without sample count or cohort definition is weak evidence.
- One global budget hides route-specific and device-specific cost.
- Metric improvement can trade away correctness or accessibility.
- Unhandled errors are last-resort signals with incomplete context.
- Raw URLs, DOM text, and form values can leak private data.
- High-cardinality labels can make telemetry costly and unusable.
- Client clocks across devices cannot establish a global event order.
- Source maps must match the release and need controlled access.
- Reporting can fail, so telemetry cannot be part of product success.
- A budget that does not fail a build, rollout, or alert is a wish.

#### Related material

- [Semantic interfaces and accessibility][interfaces]
- [Modules, toolchains, and production builds][toolchain]
- [Rendering architectures, hydration, and Server Components](
  #b4-rendering-architectures-hydration-and-server-components)
- [Frontend architecture and design systems][architecture]
- [Backend testing and operations][backend-quality]
- [Core Web Vitals](https://web.dev/articles/vitals)
- [W3C Resource Timing](https://www.w3.org/TR/resource-timing/)
- [Testing Library query guidance](
  https://testing-library.com/docs/queries/about/)
- [Playwright testing practices](
  https://playwright.dev/docs/best-practices)

[interfaces]: #a1-semantic-interfaces-and-accessibility
[backend-quality]: ../backend/README.md

### B.4 Rendering Architectures, Hydration, and Server Components

#### The problem it solves

A product must decide where and when to turn data into interface
output. That decision changes the first visible content, interaction
delay, browser JavaScript, server load, caching, failure recovery, and
the boundary between trusted server code and public client code.

#### How it works

**Rendering** computes an interface description. Delivery may happen
through several models:

- Client-side rendering sends a small HTML shell, then browser
  JavaScript fetches data and creates the useful interface.
- Static generation renders HTML ahead of requests, usually during a
  build. A content or time policy decides when to regenerate it.
- Server-side rendering renders HTML for a request. It can use
  request-specific data, but every uncached request consumes server
  capacity.
- Streaming server rendering sends an early shell, then sends later
  regions as their data and code finish. A **Suspense boundary** owns
  the fallback and reveal point for one such region.

These models can coexist by route and by region. Early HTML can show
content before application JavaScript loads, but HTML alone does not
make a React interface interactive. **Hydration** attaches React to
server-produced DOM. The first browser render must describe the same
elements, attributes, and text as the server render. After that match,
events and later state updates can take ownership.

Rendering architecture also owns the route's HTTP and document
contract. A public route should return a meaningful status, useful
initial HTML for its intended user agents, a unique title and
description, and one canonical URL when duplicates are possible.
Language variants may link to alternates with `hreflang`. **Structured
data** is machine-readable metadata about visible content; it must
describe what users can actually see. Navigation that must be
discovered should use real links with `href`, while a sitemap can help
enumerate important URLs; neither guarantees indexing or ranking.

Choose that contract from the actual audience. Some crawlers execute
JavaScript and some do not; even rendering crawlers may process it
later or under resource limits. Server or static HTML is therefore a
reliability choice for public content, not a magic ranking switch.
Test target bots and sharing clients instead of assuming that every
consumer runs the same browser application.

Return `404` for a missing public resource and an appropriate redirect
status when its canonical location changed. A styled error page with
status `200` is a **soft 404**: it says the resource exists. Private
routes
may return a minimal shell or require authentication, but `noindex`,
robots directives, and omitted links are not access control.

Metadata belongs to the route snapshot. It must appear in the initial
response for consumers that do not run client code. On client
navigation, it must change too. Keep the visible heading,
title, description, canonical link, language, and structured data
about the same resource. Never leak private data into public metadata.

React can prioritize hydration around Suspense boundaries. This is
often called **selective hydration**: an interaction or ready boundary
can receive work without waiting for every boundary on the page. It
does not remove the cost of downloaded client modules.

React Server Components solve a different problem from server-side
HTML. A **Server Component** runs ahead of the client, at build time
or for a request, and produces a serialized React result. Its module
and server-only dependencies are not sent as browser JavaScript. The
result may then participate in server-side HTML rendering and later
client navigation.

A file beginning with `'use client'` creates a client module boundary.
That module and its transitive dependencies enter the client graph.
Its components may still contribute to initial server HTML; "client"
means their code can execute in the browser. Values crossing from a
Server Component into a Client Component must use React's supported
serializable contract. Keep the boundary low in the tree so only
code needing state, effects, events, or browser APIs becomes client
code.

Server Functions cross in the opposite direction: client code invokes
a server-owned function through framework transport. Treat that call
as a public network endpoint. Authenticate, authorize, validate, and
bound it. Make replay behavior explicit on the server.

Suspense coordinates work that React or a supporting framework knows
how to suspend, such as a cached Promise read with `use`. Fetching in
an effect does not activate a Suspense boundary. A transition marks a
state update as non-urgent so React can keep already useful content
responsive while preparing the replacement. Text input itself remains
urgent.

In React 19.3, `browser()` provides a first-class way for a component
to suspend only during server rendering, and `<ViewTransition>` is
stable for transition-driven DOM animations. Both are optional tools;
they do not repair a poor server/client boundary or an inaccessible
loading state.

#### Step-by-step examples

Input: public request `/products/p1`. Product `p1` is ready in `40ms`.
Recommendations need `900ms`; the favorite button needs browser
state. Expected output: status `200`, specific metadata, and product
HTML arrive first. The button becomes interactive after hydration, and
recommendations reveal later without replacing the product.

1. **Before:** The server has the route parameter `p1` and no product
   snapshot. **Decision:** Validate the parameter, read product `p1`,
   and start recommendation work without awaiting it yet. **After:**
   The product snapshot is ready and recommendation work is pending.
   **Known correct:** Both operations belong to product `p1`.
2. **Before:** The route has a product snapshot but no response
   contract. **Decision:** Return status `200`, title `Desk lamp`, a
   matching description, and canonical URL `/products/p1`. **After:**
   the initial document and metadata name the same product.
   **Known correct:** A client without JavaScript can identify `p1`.
3. **Before:** The Server Component holds a trusted database record.
   **Decision:** Render server-owned product content and pass only the
   product ID and initial favorite state to the favorite Client
   Component. **After:** Database code and supplier cost stay outside
   the client graph. **Known correct:** The client boundary has the
   minimal serializable input it needs.
4. **Before:** Recommendations are pending and the product shell is
   ready. **Decision:** Stream the shell with the recommendation
   boundary's accessible fallback. **After:** The browser can display
   the product before recommendation work finishes. **Known correct:**
   Slow secondary data does not block primary content.
5. **Before:** The browser has server HTML but no attached event
   handler. **Decision:** Render the same initial product and favorite
   state, then hydrate the Client Component. **After:** The favorite
   button accepts input. **Known correct:** Hydration reused matching
   DOM instead of repairing a server/client disagreement.
6. **Before:** The recommendation boundary still shows its fallback.
   **Decision:** Reveal its streamed result when the original `p1`
   work completes. **After:** Product and recommendations are visible.
   **Known correct:** Every displayed region describes product `p1`.

Failure example: if the server formats a price with `en-US` but the
first client render uses an inferred `vi-VN` locale, text can differ
during hydration. Serialize the chosen locale or the formatted value
as part of the initial snapshot instead of guessing twice.

Missing-route example: request `/products/gone` has no product.
Expected output: a useful not-found document with HTTP status `404`, a
not-found title, and no product structured data. Returning that same
page with `200` would be a soft 404: the transport would claim a
resource exists while the visible output says it does not.

#### Why it works

Each phase has one contract. The server validates route input and owns
trusted data access. The streamed shell contains only regions whose
dependencies are ready plus explicit fallbacks for pending regions.
The client receives the same initial snapshot that produced the HTML,
so its first render matches before it attaches interaction.

The Server Component boundary keeps server modules out of the client
dependency graph. The Client Component boundary names exactly where
browser execution begins. Suspense boundaries keep unfinished regions
from blocking independent ready regions, and each completed result is
still associated with the route identity that started it.

The status and metadata contract lets browsers, crawlers, link preview
clients, caches, and assistive technology interpret the same route.
Matching visible content and metadata prevents one representation from
making a claim the other does not support. Real links preserve the web
navigation contract even when client-side routing enhances them.

#### TypeScript example

This boundary mapper removes a server-only field and formats the price
once. The same returned snapshot must drive server HTML and the first
client render.

```typescript
type ProductRecord = {
  id: string;
  name: string;
  priceCents: number;
  supplierCostCents: number;
  version: string;
};

type ProductCardProps = Readonly<{
  formattedPrice: string;
  id: string;
  name: string;
  version: string;
}>;

function toProductCardProps(
  product: ProductRecord,
  locale: string,
): ProductCardProps {
  // Convert trusted server data into the complete, intentionally
  // public snapshot used on both sides of the hydration boundary.
  const formattedPrice = new Intl.NumberFormat(locale, {
    currency: "USD",
    style: "currency",
  }).format(product.priceCents / 100);
  return {
    formattedPrice,
    id: product.id,
    name: product.name,
    version: product.version,
  };
}

function assertSameInitialSnapshot(
  serverProps: ProductCardProps,
  clientProps: ProductCardProps,
): void {
  // This application requires one identity, version, and displayed
  // value set before the browser performs later updates.
  const sameSnapshot =
    serverProps.id === clientProps.id &&
    serverProps.version === clientProps.version &&
    serverProps.name === clientProps.name &&
    serverProps.formattedPrice === clientProps.formattedPrice;
  if (!sameSnapshot) {
    throw new Error("server and client snapshots differ");
  }
}

const productRecord: ProductRecord = {
  id: "p1",
  name: "Desk lamp",
  priceCents: 2599,
  supplierCostCents: 1100,
  version: "v4",
};

const serverProps = toProductCardProps(productRecord, "en-US");
const clientProps = toProductCardProps(productRecord, "en-US");
assertSameInitialSnapshot(serverProps, clientProps);

// Exercise the failure path with changed server-visible content.
const newerRecord = {
  ...productRecord,
  name: "Desk lamp, second edition",
  version: "v5",
};
const newerProps = toProductCardProps(newerRecord, "en-US");
let mismatchRejected = false;
try {
  assertSameInitialSnapshot(serverProps, newerProps);
} catch {
  mismatchRejected = true;
}
console.assert(mismatchRejected);
console.assert(!("supplierCostCents" in serverProps));
```

In an application, a framework transports the snapshot and performs
hydration. The assertion makes the data contract visible; it is not a
replacement for React's own development diagnostics.

#### Work and memory

Count the full path: server data reads, render CPU, cache storage,
HTML and React payload bytes, client JavaScript, parse and execution,
hydration, browser rendering, and retained component state. Static
generation moves render work before requests but can create long
builds or stale output. Request rendering spends server capacity and
can increase time to first byte when data is slow.

Streaming can improve the time at which useful regions appear. It
does not guarantee fewer bytes, less total server work, or earlier
interaction. Server Components can remove server-only modules from
the client bundle, but their serialized results and client boundaries
still have transfer and processing costs. Measure each route with its
real cache hit rate and user device distribution.

Generating metadata adds data reads and output bytes when it cannot
reuse the route snapshot. Crawling and prerendering consume server or
build capacity. Large route sets make sitemap generation, freshness,
canonicalization, and invalidation operational work. Measure discovery
and rendering with intended user agents, not only a desktop browser.

#### When to use it

Use static generation for public content that can follow a deliberate
freshness policy. Use request rendering for personalized or rapidly
changing content that must be present in initial HTML. Use streaming
when independent regions become ready at meaningfully different times.

Use server or static output for public routes when intended crawlers,
sharing clients, slow devices, or no-script users need useful first
HTML. Give every indexable route an explicit status, canonicalization,
metadata, and language policy. Keep private routes out of public
output and enforce their access on the server.

Use Server Components through a framework when server data access and
smaller client dependency graphs justify the operational complexity.
Use Client Components at the smallest practical interactive boundary.
Use transitions when replacing non-urgent content should not block an
urgent interaction.

#### When not to use it

Do not add a server merely to render a private tool that already loads
quickly on capable managed devices. Do not request-render globally
cacheable content on every visit. Do not stream a tiny page whose data
finishes together; the protocol and fallback churn add no useful
progress.

Do not server-render solely from an unmeasured promise of better
ranking. Do not generate many near-duplicate URLs without a deliberate
canonical policy. Do not expose private content for discovery and
then rely on robots rules to hide it.

Do not implement the React Server Component wire protocol directly.
The component model is stable in React 19, but the underlying bundler
and framework integration APIs can change between minor releases. Use
a supported framework and its version guidance.

#### Common mistakes and edge cases

> **Warning:** A hydration warning is evidence of two different first
> renders, not harmless console noise. Values from `Date.now()`,
> `Math.random()`, locale inference, browser-only branches, invalid
> HTML nesting, or changing data can all produce the disagreement.

- A Client Component's transitive imports join the client graph.
- A Server Component cannot own browser state, effects, or DOM APIs.
- Props crossing into a Client Component must be serializable.
- Suspense does not observe fetching started inside an effect.
- Too many boundaries can produce fallback flicker and reveal churn.
- One large boundary can hide ready primary content behind slow data.
- Per-user HTML or payloads must not enter a shared public cache.
- Server Function arguments are client-controlled and need full
  authorization and validation on every call.
- Streaming needs abort, timeout, disconnect, and error policies.
- A `200` not-found page is a soft 404; return the truthful status.
- One global title or canonical URL misidentifies distinct routes.
- Client-only metadata is absent for consumers that do not run it.
- Structured data that disagrees with visible content is misleading.
- Buttons styled as links lose the ordinary link discovery contract.
- Robots directives affect crawling behavior, not authorization.
- Current React and framework security advisories must be applied to
  every deployed Server Component and Server Function stack.

#### Related material

- [React components, state, and effects][react]
- [Server data, routing, caching, and forms][data]
- [Testing, performance, and observability][quality]
- [HTTP, caching, and network delivery][network]
- [Browser storage, origins, and security][security]
- [React Server Components reference](
  https://react.dev/reference/rsc/server-components)
- [React Suspense reference](
  https://react.dev/reference/react/Suspense)
- [HTML metadata](
  https://html.spec.whatwg.org/multipage/semantics.html)
- [HTML links and canonical relationships](
  https://html.spec.whatwg.org/multipage/links.html)
- [Google Search documentation](
  https://developers.google.com/search/docs)

## C. Senior Engineering

### C.1 Frontend Architecture and Design Systems

#### The problem it solves

As a product and its team grow, a change should remain understandable,
testable, and releasable without knowing the whole repository. Unclear
ownership instead creates feature-to-feature imports, duplicate state,
inconsistent interaction, visual drift, large bundles, and releases
that require every team to coordinate.

Frontend architecture must answer four questions: which module owns a
decision, which direction dependencies may point, which runtime state
crosses a boundary, and how shared contracts change without silently
breaking consumers. A design system applies the same discipline to
reusable interface decisions.

#### How it works

Architecture is not a folder diagram. It is a set of enforced
boundaries, dependency rules, runtime contracts, owners, and change
policies. Begin with the smallest deployment model that meets the
product's needs, usually a modular application rather than several
independently deployed frontends.

Organize most product code by **vertical feature slice**: keep one
workflow's UI, state transitions, requests, schemas, and tests close
together. A checkout slice may own cart review and order submission;
an account slice may own profile editing. A layer named `components`
or `utils` does not explain which workflow owns a change.

Use a small number of roles with one-way dependency rules:

- The **application composition** layer assembles routes, providers,
  global policy, error boundaries, and feature entry points. It may
  depend on every lower role; lower roles do not import it.
- A **feature** owns a product capability and exposes a narrow public
  entry point. One feature does not deep-import another feature's
  internal component, cache, reducer, or schema. Cross-feature work
  goes through composition, a deliberately shared domain contract, or
  a backend boundary.
- The **design system** owns reusable visual and interaction language:
  tokens, accessible primitives, documented variants, and composed
  patterns that are genuinely shared. It does not import product
  workflows.
- The **platform** layer wraps stable technical capabilities such as
  HTTP transport, telemetry, localization, storage, and browser
  adapters. It does not decide checkout or account policy.

The intended import direction is:

```text
application -> feature -> design system -> platform
      |           +---------------------> platform
      +----------------> design system
      +---------------------------------> platform
```

An arrow means "may import." Public entry points make that permission
explicit. Prevent deep imports and cycles with package exports, path
rules, a dependency-graph check, or equivalent build enforcement.
Code review alone does not reliably preserve an architectural rule.

A repository boundary and a runtime boundary solve different
problems. A folder, workspace package, or monorepo can enforce imports
and ownership while still producing one deployment. Split a package
when it has a stable public contract, owner, test boundary, or release
need. Do not publish a package merely because two files share a name.
A monorepo centralizes source and tooling; it does not automatically
create modular code.

Keep runtime ownership equally explicit:

- The URL owns shareable navigation and filter state.
- A server-data cache owns remote snapshots, freshness, and retries.
- A feature owner keeps workflow state and invariants.
- A component keeps temporary interaction state that no peer needs.
- The design system keeps only behavior intrinsic to its primitive,
  such as whether its dialog is open when used uncontrolled.

Do not copy one fact into several owners. Pass an ID or callback
across a boundary and let the owner perform its transition. Events
describe facts that happened; commands request an action and can be
refused. Version cross-boundary event, command, URL, and data shapes
when old and new producers can coexist.

**Web Components** are browser primitives for a reusable element
contract. A custom element gives a hyphenated tag such as
`date-range`; shadow DOM can attach an encapsulated internal tree and
styles; slots let a consumer provide named light-DOM content. These
primitives can cross framework boundaries, but they do not by
themselves supply product ownership, tokens, documentation, releases,
or accessibility. A Web Component is not automatically a design
system.

Define the public contract before the internal tree:

- Attributes are serialized strings suitable for declarative input.
  Properties can expose richer live values. Keep reflection rules and
  update timing explicit.
- Methods request imperative actions only when ordinary properties or
  events cannot express the workflow.
- Custom events report facts. State whether they bubble and cross a
  shadow boundary with `composed`; event targets may be retargeted at
  that boundary.
- Slots name composition points. CSS custom properties expose theme
  values. `part` and `::part()` expose selected styling surfaces.
- Semantics, names, keyboard behavior, focus, validation, and form
  participation remain part of the element's public contract.

Custom elements can appear before their class is defined. The browser
upgrades matching elements after registration, so provide useful
fallback or hidden pre-upgrade behavior and avoid assuming that the
constructor ran during parsing. Keep constructors small; connect,
disconnect, attribute changes, and adoption have lifecycle callbacks.
Cleanup must tolerate repeated connection and disconnection.

Shadow DOM encapsulates a tree and most style selection; it is not a
security boundary. Page code can still use the public surface, and a
closed root does not make client-delivered data secret.
Accessibility behavior also crosses the boundary, so test the final
accessibility tree, labels, focus order, and form submission rather
than treating encapsulation as proof of correctness.

Prefer native form controls. When a genuinely custom control must
participate in form value, validity, labels, and disabled state, a
form-associated custom element can use `ElementInternals` in the
supported browser matrix. Document fallback behavior. For server
output, verify upgrade behavior and support for Declarative Shadow DOM
instead of assuming that one hydration model works everywhere.

Use wrappers at framework seams when event names, property assignment,
server rendering, typing, or lifecycle expectations differ. The
underlying element can stay framework-neutral while each adapter
presents an idiomatic local API.

A **design system** is a maintained product for product teams, not a
gallery of styled components. Its layers usually include:

1. **Foundations** define typography, icon, motion, elevation,
   spacing, color, and layout decisions.
2. **Design tokens** give those decisions stable names that tools and
   platforms can exchange. A token is a named value plus type and
   optional metadata, not a CSS variable by definition.
3. **Primitives** implement one reusable semantic interaction such as
   button, field, dialog, or menu behavior.
4. **Patterns** compose primitives for recurring product tasks without
   absorbing one feature's business rules.
5. **Documentation and tests** expose supported states, composition,
   accessibility, localization, performance, and migration contracts.

Separate token roles. A base token such as `color.blue.600` names a
palette value. A semantic token such as `color.action.primary` names
intent and may alias a different base value per theme. A component
token such as `button.primary.background` exists only when a component
needs a stable override point. Product code should normally consume
semantic or component intent, not a raw blue value.

The Design Tokens Community Group Format Module 2025.10 is the first
stable exchange format for typed tokens, groups, aliases, and
extensions. It is a Community Group report, not a W3C Recommendation.
Use it when tool interoperability helps, while keeping naming,
ownership, theme policy, and generated output under product control.
Its separate resolver work remains a draft and is not required for a
sound token architecture.

Resolve themes by changing semantic assignments, not by adding a
`dark` branch inside every component. Treat light, dark,
high-contrast, brand, reduced-motion, and density choices as
independent dimensions only when the product truly supports them.
Otherwise combinations grow multiplicatively and many remain
untested. User-agent modes such as forced colors can override authored
colors, so preserve semantics and test the rendered result instead of
assuming token values win.

Build primitives from native semantics first. A headless primitive may
own behavior and accessibility while consumers own styling; a styled
primitive may own both. Either way, define controlled and uncontrolled
state, focus behavior, keyboard interaction, labeling, form behavior,
portals, layering, server rendering, and error cases. Prefer
composition and a few coherent variants over a Boolean flag for every
consumer request.

The public contract includes more than TypeScript props. It includes
rendered semantics, focus order, keyboard behavior, DOM needed for
styling or tests, CSS custom properties, token names, responsive
behavior, visual output, server-rendering support, and package side
effects. A change to any of these can require a migration even when
the type checker stays green.

Exercise every supported component state in isolated examples:
default, hover where relevant, focus, active, disabled, loading,
invalid, empty, long translated text, right-to-left text, narrow and
wide containers, zoom, reduced motion, forced colors, and supported
themes. Combine semantic assertions, keyboard interaction,
accessibility checks, visual comparison, and a small set of real
consumer tests. Snapshots alone cannot establish the contract.

Publish packages deliberately. Define explicit exports, generate or
check types, mark runtime and style side effects accurately, and avoid
bundling duplicate framework instances. Use a peer dependency only
when the consumer must provide the same runtime instance; document the
supported range and test its edges. A version number communicates the
release, but a changelog, examples, deprecation period, and migration
tool explain how to adopt it.

Choose a **micro-frontend** only when independently operating teams
need independent delivery and that autonomy outweighs runtime cost.
Define one shell owner and contracts for routing, authentication,
navigation, shared dependencies, styling, communication, telemetry,
loading, failure, and rollback. Prefer coarse page or workflow
boundaries. Many independently deployed widgets create network,
version, accessibility, and ownership failures at every seam.

Governance keeps shared code useful. Give each public surface an
owner, contribution criteria, support policy, deprecation policy, and
decision record for consequential tradeoffs. Track adoption, duplicate
implementations, accessibility defects, release lag, bundle cost, and
migration completion. Usage count alone can reward a difficult API
that consumers cannot escape.

#### Step-by-step examples

First input: checkout and account each contain a custom confirmation
dialog. Checkout must reject closing while payment finalization is
uncertain; account may close immediately. Expected output: one shared
accessible dialog primitive, two independently owned workflows, and
no design-system import of product code.

1. **Before:** Both features separately implement a backdrop, focus
   movement, Escape handling, labeling, and focus restoration.
   **Decision:** Record the common interaction contract and the two
   different close policies before extracting code. **After:** The
   shared candidate is dialog mechanics; payment policy remains a
   checkout rule. **Known correct:** Repetition and domain difference
   are explicitly separated.
2. **Before:** Checkout's component imports payment state while the
   account component imports profile state. **Decision:** Create a
   primitive that receives `open`, labels, content, and an
   `onCloseRequest` callback. **After:** Each feature decides whether
   to honor a close request. **Known correct:** The design system
   cannot approve or cancel a payment.
3. **Before:** The primitive opens with focus on the trigger.
   **Decision:** Move focus to the dialog, contain keyboard
   navigation, label it, and restore focus when the owner closes it.
   **After:** Both consumers satisfy the same interaction contract.
   **Known correct:** Shared mechanics are tested once and then
   exercised in each real workflow.
4. **Before:** Checkout is finalizing and receives an Escape close
   request. **Decision:** Checkout refuses it and renders an
   explanation; account accepts the same request in its workflow.
   **After:** One primitive supports both outcomes without a
   `paymentMode` flag. **Known correct:** The owner callback expresses
   variation without putting product policy inside shared code.
5. **Before:** A new primitive version changes focus restoration.
   **Decision:** Run primitive tests and consumer contract tests, then
   publish the behavior change with a migration window. **After:**
   consumers adopt deliberately. **Known correct:** A passing type
   check cannot hide the behavioral change.

Second input: light and dark themes must style a primary button. Base
colors are `blue.600 = #075985` and `blue.300 = #7dd3fc`. Expected
output: product code asks for `color.action.primary`; light resolves
to `#075985`, dark resolves to `#7dd3fc`, and the button never names
either palette entry.

1. **Before:** The two base tokens name reusable color values but no
   purpose. **Decision:** Define `color.action.primary` as an alias to
   `blue.600` in light and `blue.300` in dark. **After:** Each theme
   maps one intent to a tested value. **Known correct:** Product
   intent is independent of a particular shade name.
2. **Before:** The button needs a background. **Decision:** Alias
   `button.primary.background` to `color.action.primary`. **After:**
   the component can expose a specific override without copying the
   color. **Known correct:** Base, semantic, and component roles
   remain distinct.
3. **Before:** Dark theme is selected. **Decision:** Resolve the
   semantic alias once and generate the platform representation.
   **After:** the button background is `#7dd3fc`. **Known correct:**
   No component branch guessed which dark color to use.
4. **Before:** A new high-contrast theme needs a stronger pair.
   **Decision:** Map `color.action.primary` to `#000000` and its text
   intent to `#ffffff`, then test the rendered control. **After:**
   supported buttons inherit the pair. **Known correct:** The theme
   boundary centralizes intent while browser verification checks the
   actual contrast and forced-color result.

Third input: three teams share one repository and deploy the same web
application every weekday. Expected output: enforce feature and
package boundaries in the existing deployment, not three
micro-frontends.

1. **Before:** Cross-feature deep imports make ownership unclear.
   **Decision:** Define public feature entry points and fail forbidden
   imports in the build. **After:** teams can change internals without
   rewriting peers. **Known correct:** The observed problem is source
   coupling, which a deployment split is not required to solve.
2. **Before:** All features use one React runtime, route shell, and
   release train successfully. **Decision:** retain one deployment
   and measure build, test, and coordination delays. **After:** there
   is no new runtime integration boundary. **Known correct:** autonomy
   has not yet justified duplicate delivery machinery.
3. **Before:** One future workflow gains a separate compliance owner,
   release approval, and failure-isolation requirement. **Decision:**
   evaluate a coarse deployment boundary with explicit shell
   contracts. **After:** the decision follows a demonstrated operating
   need. **Known correct:** micro-frontends remain an organizational
   tradeoff, not a synonym for modularity.

Fourth input: a `date-range` control must work in a React booking app
and a framework-free site. Expected output: one browser element
contract, a thin React adapter, useful pre-upgrade content, and no
framework runtime inside the content site.

1. **Before:** Both hosts need the same date-range semantics, but use
   different APIs. **Decision:** Define attributes for serialized date
   limits, properties for live dates, and help-text slots. Define a
   bubbling composed `range-change` event. **After:** The
   browser contract is independent of either host. **Known correct:**
   Inputs, outputs, composition, and event transport are named.
2. **Before:** HTML contains `date-range` before its module loads.
   **Decision:** Render labelled native date inputs as light-DOM
   fallback, then upgrade without discarding their current values.
   **After:** Slow or failed JavaScript still leaves a usable input.
   **Known correct:** Late definition leaves no blank interface.
3. **Before:** The custom element class connects. **Decision:** Attach
   its shadow tree once, wire events, and expose intentional parts
   and custom properties. **After:** Reconnection does not duplicate
   listeners or internal DOM. **Known correct:** Lifecycle
   work is repeatable and the styling surface stays bounded.
4. **Before:** A user selects a valid range in the internal controls.
   **Decision:** Update public properties and dispatch the declared
   event across the shadow boundary. **After:** The static site
   receives the event; the React adapter calls its typed callback.
   **Known correct:** Both hosts observe one fact without internal
   access.
5. **Before:** The control enters a form and a target browser lacks a
   required custom-form capability. **Decision:** Use the documented
   native fallback rather than silently omitting its value. **After:**
   submission is correct. **Known correct:** Form behavior follows
   the support contract rather than the ideal implementation alone.

#### Why it works

Each product decision has one authoritative owner. Imports point from
composition toward capabilities and from product policy toward stable
technical mechanisms, so a lower layer cannot silently acquire a
feature dependency. Enforced public entry points keep implementation
changes local and make cross-boundary contracts visible.

The same rule protects the design system. Tokens name intent without
making components depend on one palette. Primitives own only the
semantics and interaction shared by every consumer. Feature callbacks
retain domain decisions, so a new payment rule changes checkout rather
than the dialog package.

Tests at the shared contract and selected real consumers detect both
local defects and incompatible assumptions. Versioning, deprecation,
and migration evidence make time part of the contract. The simplest
deployment remains available until independent operation provides more
value than its runtime and coordination cost.

For a Web Component, the browser owns upgrade and shadow-boundary
semantics while the element owns its declared public surface. Fallback
content works before definition; after upgrade, properties and
events carry the same contract into each host. Thin adapters translate
host conventions without forking the underlying behavior.

#### TypeScript example

This small build-time rule models repository boundaries. `boundary`
names the module that may change internally as one unit. Imports
within that boundary are allowed; imports to another role must follow
the declared direction. Each module `id` must be unique. The lookup
contract is `module ID -> boundary and layer`. For example,
`checkout/view` maps to the checkout boundary in the feature layer.

```typescript
type ProductLayer = "application" | "feature";
type SharedLayer = "design-system" | "platform";
type Layer = ProductLayer | SharedLayer;

type ModuleBoundary = Readonly<{
  boundary: string;
  id: string;
  layer: Layer;
}>;

type ImportEdge = Readonly<{
  sourceId: string;
  targetId: string;
}>;

type AllowedTargets = Readonly<Record<Layer, readonly Layer[]>>;

const allowedTargetLayers: AllowedTargets = {
  application: ["feature", "design-system", "platform"],
  feature: ["design-system", "platform"],
  "design-system": ["platform"],
  platform: [],
};

function validateImportGraph(
  modules: readonly ModuleBoundary[],
  imports: readonly ImportEdge[],
): readonly string[] {
  const moduleById = new Map<string, ModuleBoundary>();
  const errors: string[] = [];

  // Reject an ambiguous graph instead of silently replacing one
  // module definition with a later definition that has the same ID.
  for (const module of modules) {
    if (moduleById.has(module.id)) {
      errors.push(`duplicate module: ${module.id}`);
      continue;
    }
    moduleById.set(module.id, module);
  }

  for (const importEdge of imports) {
    const sourceId = importEdge.sourceId;
    const targetId = importEdge.targetId;
    const importDescription = `${sourceId} -> ${targetId}`;
    const source = moduleById.get(sourceId);
    const target = moduleById.get(targetId);
    if (!source || !target) {
      errors.push(`unknown import: ${importDescription}`);
      continue;
    }

    // Modules inside one boundary may collaborate without exposing
    // those internal edges to another feature or package.
    if (source.boundary === target.boundary) continue;

    // Cross-boundary imports must point to a permitted lower role.
    const allowedLayers = allowedTargetLayers[source.layer];
    if (!allowedLayers.includes(target.layer)) {
      errors.push(`forbidden import: ${importDescription}`);
    }
  }

  return errors;
}

const modules: readonly ModuleBoundary[] = [
  { boundary: "shell", id: "app/routes", layer: "application" },
  { boundary: "checkout", id: "checkout/view", layer: "feature" },
  { boundary: "account", id: "account/model", layer: "feature" },
  {
    boundary: "dialog",
    id: "design-system/dialog",
    layer: "design-system",
  },
  { boundary: "dom", id: "platform/dom", layer: "platform" },
];

const validImports: readonly ImportEdge[] = [
  { sourceId: "app/routes", targetId: "checkout/view" },
  {
    sourceId: "checkout/view",
    targetId: "design-system/dialog",
  },
  {
    sourceId: "design-system/dialog",
    targetId: "platform/dom",
  },
];
const validErrors = validateImportGraph(modules, validImports);
console.assert(validErrors.length === 0);

const forbiddenImports: readonly ImportEdge[] = [
  { sourceId: "checkout/view", targetId: "account/model" },
];
const boundaryErrors = validateImportGraph(modules, forbiddenImports);
const expectedBoundaryError =
  "forbidden import: checkout/view -> account/model";
console.assert(boundaryErrors[0] === expectedBoundaryError);

const unknownImport: readonly ImportEdge[] = [
  { sourceId: "checkout/view", targetId: "missing/module" },
];
const unknownErrors = validateImportGraph(modules, unknownImport);
console.assert(unknownErrors.length === 1);
```

Real enforcement should also reject dependency cycles, undeclared
package exports, and browser or server code in the wrong build graph.
The example isolates the central direction rule and its failure path.

#### Work and memory

For `m` modules and `e` imports, constructing the example's module map
visits `m` modules and validating edges visits `e` imports. It uses
`O(m + e)` work and `O(m + e)` memory including returned errors. A
cycle check also visits the graph in `O(m + e)` work.

Real boundaries add builds, packages, generated types, contract tests,
documentation, review, and migration work. Too few boundaries increase
change coupling; too many increase discovery and release overhead.
Measure lead time, failed changes, dependency cycles, affected tests,
and runtime cost rather than counting packages.

Tokens and component packages add generated CSS or native assets,
browser style work, documentation builds, and retained versions.
Tree-shaking cannot remove hidden side effects. Micro-frontends add
requests, duplicated dependencies, shell code, communication, failure
states, telemetry joins, and operational ownership. Budget those costs
per route and supported device.

Custom elements add registration and upgrade work. Shadow trees retain
nodes, styles, listeners, and state for every instance. Small elements
are cheap, but thousands of deep instances still consume style and
layout work. Framework adapters, polyfills where required, and
cross-browser contract tests add package and maintenance cost.

#### When to use it

Create a boundary when code has one stable owner, a coherent reason to
change, a testable contract, or a distinct runtime requirement. Start
with feature slices and enforced public entry points. Promote a shared
module only after consumers demonstrate the same durable contract.

Build a design system when several product areas need consistent
semantics, accessibility, interaction, visual decisions, and migration
support. Use typed design-token interchange when several tools or
platforms must share those decisions.

Use a Web Component when one stable interface capability must run
across frameworks, embedded hosts, or framework-free pages and browser
support meets the product matrix. Keep the contract narrow, semantic,
and versioned; provide idiomatic adapters for major host frameworks.

Consider micro-frontends when independently accountable teams require
independent release, technology, or failure boundaries and can operate
the integration contracts for their full lifetime.

#### When not to use it

Do not create a package for speculative reuse, a one-off wrapper, or a
name that merely sounds architectural. Keep a small application direct
until a real change boundary appears. Do not put product workflows in
the design system or use global state to avoid feature ownership.

Do not use a custom element when a native element already satisfies
the contract, or merely to hide same-framework components. Avoid
shadow DOM when unrestricted consumer styling or DOM inspection is
the intended API. A framework component is often simpler inside one
framework and release boundary.

Do not adopt micro-frontends to organize folders, speed one repository
build, permit arbitrary frameworks, or let teams ignore shared user
experience. Fix source boundaries, build caching, or governance at the
level where those problems exist.

#### Common mistakes and edge cases

> **Warning:** A shared component that accepts every feature rule as a
> flag is a hidden collection of domain workflows, not a primitive.

- A `shared` folder without admission rules becomes an ownerless dump.
- Barrel exports and path aliases can hide cycles and deep imports.
- Mirrored server, URL, and component state drift after one update.
- A callback named `onChange` may hide whether it reports a fact or
  requests a command that can be rejected.
- Raw-color tokens make dark, brand, and contrast themes expensive.
- Component tokens for every CSS property expose implementation rather
  than durable design decisions.
- Theme dimensions multiply into untested combinations.
- Removing a DOM wrapper can break consumer CSS, tests, and focus even
  when public TypeScript types are unchanged.
- A portal can cross stacking, ownership, and server-rendering
  boundaries; test where it mounts and how focus returns.
- Shared framework code can fail when multiple runtime copies load.
- Incorrect side-effect metadata can either retain dead code or remove
  required styles and initialization.
- A peer-dependency range is a compatibility claim that needs tests.
- Independent deployments still share routing, authentication,
  accessibility, performance, incidents, and the user's task.
- Synchronous cross-application events create temporal coupling; use a
  versioned contract and explicit absent or late-consumer behavior.
- A design system without contribution and migration support becomes
  either frozen infrastructure or a bottleneck.
- Adoption percentage does not reveal forks, escape hatches, defects,
  release lag, or consumer satisfaction.
- A custom element name without a hyphen cannot be registered.
- Constructor-only setup fails across upgrade and reconnection states.
- A non-composed custom event may disappear at the shadow boundary.
- Shadow DOM retargets events; consumers must not depend on internals.
- Closed shadow DOM is encapsulation, not protection for secrets.
- Hidden internal selectors force consumers into brittle workarounds.
- Shadow content still needs names, focus, keyboard, and forms.
- Server output needs an upgrade and Declarative Shadow DOM policy.

#### Related material

- [Semantic interfaces and accessibility][interfaces]
- [CSS, layout, responsive design, and rendering][css]
- [Modules, toolchains, and production builds][toolchain]
- [React components, state, and effects][react]
- [Server data, routing, caching, and forms][data]
- [Testing, performance, and observability][quality]
- [Delivery, migrations, and leadership][leadership]
- [Design Tokens Format Module 2025.10](
  https://www.designtokens.org/tr/2025.10/format/)
- [Web Content Accessibility Guidelines 2.2](
  https://www.w3.org/TR/WCAG22/)
- [Storybook accessibility testing](
  https://storybook.js.org/docs/writing-tests/accessibility-testing)
- [HTML custom elements](
  https://html.spec.whatwg.org/dev/custom-elements.html)
- [DOM Standard shadow trees](https://dom.spec.whatwg.org/)
- [CSS Shadow Parts](https://www.w3.org/TR/css-shadow-parts/)

[leadership]: #c2-delivery-migrations-and-technical-leadership

<a id="delivery-and-leadership"></a>

### C.2 Delivery, Migrations, and Technical Leadership

#### The problem it solves

Senior engineers must change production systems while old and new
code, data, browser tabs, cached assets, and team decisions coexist.
The desired result is not merely deployed code. It is a verified user
outcome, bounded exposure when assumptions fail, a safe route forward
or back, and eventual removal of temporary compatibility machinery.

Technical leadership makes that result repeatable across people. It
turns uncertainty into explicit constraints, evidence, decisions,
owners, communication, and follow-up rather than relying on one
engineer's memory or authority.

#### How it works

Treat delivery as a controlled state transition. At every stage, know
which version serves which cohort, which data shapes remain valid,
which measurements permit the next step, and which action follows a
breach. A **migration** is complete only when the new path owns the
workflow, no supported consumer needs the old contract, compatibility
code and controls are removed, and the final state is documented.

Separate three operations that are often called a release:

- **Deployment** puts an immutable, identified artifact into an
  environment. It does not have to expose new behavior.
- **Release** exposes behavior to users through routing,
  configuration, or a feature control.
- **Migration** changes an enduring contract or representation, such
  as an API, cache, URL, local schema, design-system component, or
  ownership boundary.

Separating them makes failures attributable. A deployed artifact can
be checked before exposure. A release can be disabled without
rebuilding when its old contract remains valid. A migration can keep
compatibility until all supported consumers have crossed.

Before changing production, write a small change contract:

1. State the user or system outcome and the current baseline.
2. Name the owner, affected surfaces, dependencies, and cohorts.
3. List assumptions, failure modes, privacy or security risk, and the
   largest plausible blast radius.
4. Define preconditions, automated checks, rollout stages, minimum
   observation, and representative success and guardrail signals.
5. Set advance, hold, rollback, and roll-forward conditions before
   seeing candidate results.
6. Define communication, incident, support, and cleanup ownership.

Use one reproducible artifact across stages. Attach source revision,
build identity, configuration identity, asset manifest, source-map
identity, and schema or protocol versions. Retain protected source
maps for that exact build. Rebuilding between canary and full release
changes the thing being evaluated. Keep changes small enough that a
failed result has a short causal search space.

Frontend compatibility needs special attention because deployment
does not replace every running client. Old tabs can submit hours
later. A service worker can serve an earlier asset graph. A content
delivery network can retain HTML or modules. Back navigation can
restore a document from memory, and an extension or embedded webview
may update on another schedule. The server must tolerate every client
version the product still claims to support.

Use an **expand, migrate, contract** sequence for an incompatible
shape:

1. **Expand:** Add the new field, endpoint, token, event, or reader
   while the old contract still works.
2. **Migrate:** move bounded cohorts, backfill or translate stored
   state, compare outcomes, and reconcile differences.
3. **Contract:** after the supported overlap window and evidence show
   old use is gone, remove the old path and compatibility control.

For writes, decide which representation is authoritative. Dual writes
can partially succeed and diverge, so prefer one authoritative write.
If another store needs a message, a backend can commit the source
change and a pending message together, then deliver that message
later. This pattern is a **transactional outbox**. A shadow read may
compare a candidate result without showing it to the user, but it
still consumes capacity and must not produce a side effect. Backfills
need bounds, checkpoints, idempotency, rate limits, and reconciliation
counts.

Browser persistence is a data migration too. Version local storage,
IndexedDB, offline queues, and service-worker messages. Parse old data
as untrusted boundary input. Convert it once when practical, retain a
safe fallback or reset policy, and never let a stale queued mutation
silently act under a new meaning.

A **progressive rollout** exposes the same artifact in stages: team
use, a small stable cohort, wider cohorts, and full use. Segment by
the risk that matters, such as browser, device, route, geography,
account type, or feature interaction. Percentage alone is not
representative when the first cohort excludes the slow devices or
workflows likely to fail.

A canary is a partial, time-bounded release compared with a control.
Its signals must be attributable, sufficiently sampled, and observed
for at least the relevant work cycle. A synthetic pass is useful but
does not replace real user outcomes. Blue-green delivery makes two
environments available for a traffic switch, but data compatibility
still determines whether switching back is safe.

A **feature control**, often called a feature flag, separates behavior
exposure from artifact deployment. Common purposes differ:

- A release control supports temporary staged exposure.
- An operational control disables costly or failing optional work.
- An experiment assigns stable cohorts to compare hypotheses.
- A permission or entitlement expresses durable product policy and
  should not masquerade as a temporary release flag.

Evaluate controls from a typed key, typed default, and the minimum
lawful context. Assignment must be stable when a user moves between
requests or devices if the workflow requires it. Keep authorization on
the server; client-visible controls and values are observable and
modifiable. Define provider failure behavior, cache behavior, offline
behavior, and telemetry without putting private attributes into logs.

Every temporary control needs an owner, creation date, intended value,
removal condition, and latest removal date. Test important enabled and
disabled paths while both exist. Remove the old branch, its tests,
configuration, metrics, and documentation after the rollout. A flag
at one hundred percent is not a completed migration.

Rollback restores a known compatible state. It is not always possible:
a destructive data rewrite, sent message, payment, public API use, or
security revocation may be irreversible. In that case, stop further
exposure and **roll forward** with a compatible fix or mitigation.
Test rollback before release and state which data or side effects it
cannot undo.

Measure the user outcome and system guardrails for the candidate and a
comparable control. Include completion, correctness, accessibility,
latency, Core Web Vitals, resource use, errors, and domain invariants
such as duplicate orders. Use release and cohort identity, sample
counts, percentiles, and confidence appropriate to the decision. A
green exception rate cannot prove that users still finish checkout.

An objective plus an allowed failure level is a **service-level
objective (SLO)**. Its **error budget** is the allowed failure over a
window. A product policy can slow risky releases when the budget burns
too quickly. It should not hide individual severe failures, privacy or
security breaches, or a broken domain invariant behind an aggregate.

When a threshold is crossed, first limit harm: halt expansion, disable
the behavior, roll back when safe, or apply a bounded mitigation. Name
an incident lead, preserve a timeline, communicate impact and current
action, and separate confirmed facts from hypotheses. Restore the user
workflow before pursuing a perfect explanation.

After a significant incident, record impact, detection, timeline,
contributing conditions, response, recovery, and learning. A
**blameless postmortem** assumes people acted reasonably with the
information and system they had. It produces concrete prevention,
detection, and mitigation actions with owners and verifiable end
states. Blameless does not mean consequence-free or vague; it moves
correction from a person to the conditions that made failure likely.

Technical leadership applies the same evidence discipline before an
incident:

- Frame the decision, constraints, non-goals, unknowns, and deadline.
- Present credible options and the status quo with user, engineering,
  operational, security, accessibility, and opportunity costs.
- State a recommendation, assumptions, reversibility, owner, and the
  evidence that would reopen it.
- Record durable decisions with context, chosen option,
  consequences, alternatives, status, and replacement history.
- Consult affected specialists and users, then give one accountable
  owner authority to decide. Consensus is useful input, not a way to
  hide responsibility.
- Delegate an outcome with constraints, decision rights, checkpoints,
  and support. Delegating only tasks preserves dependence.
- Give feedback on observable behavior and impact, and transfer
  ownership gradually enough that another engineer can succeed.

Maintain a technical roadmap from product risk and evidence. Rank work
by user impact, probability, time sensitivity, reversibility,
dependency, effort, and learning value. Reserve capacity for
reliability, accessibility, security, maintenance, and migration
cleanup. Calling all unfinished engineering work "technical debt"
hides whether it is a defect, risk, deliberate tradeoff, or obsolete
idea.

Use delivery measures to find constraints, not to rank individuals.
DORA's current model measures five application- or service-level
outcomes. Throughput uses change lead time, deployment frequency, and
failed deployment recovery time. Instability uses change fail rate and
deployment rework rate. Read them together, over time, in the local
context. Turning one measure into a quota invites batching,
relabeling, or other behavior that makes the number less useful.

#### Step-by-step examples

First input: move checkout reads from a legacy store to a query cache
release `r52`. The control has `10,000` samples, completion `0.980`,
error rate `0.005`, 75th-percentile (p75) confirmation time `900ms`,
and no duplicate orders. Policy requires `1,000` candidate samples,
at most `0.005` completion drop, `0.003` error-rate increase, `100ms`
p75 increase, and zero duplicates. Expected output: hold an
undersampled candidate, advance a healthy candidate, then roll back
immediately on one duplicate order.

1. **Before:** Only the legacy read path serves users and operation
   identity already prevents repeated order creation. **Decision:**
   attach release and cohort IDs to the existing completion, error,
   latency, and duplicate-order signals. **After:** the control values
   above form the baseline. **Known correct:** candidate comparison
   has one defined workflow and one invariant.
2. **Before:** `r52` contains both read paths and the release control
   is off. **Decision:** verify the artifact, compatibility path,
   disabled behavior, and rollback in staging. **After:** production
   can receive the same artifact without exposing the query cache.
   **Known correct:** deployment is separated from release.
3. **Before:** internal and one-percent exposure produce `600` valid
   candidate samples. **Decision:** hold at one percent because the
   minimum is `1,000`, even though early rates look healthy.
   **After:** exposure remains bounded. **Known correct:** absence of
   an observed failure is not mistaken for sufficient evidence.
4. **Before:** the candidate reaches `2,000` samples with completion
   `0.977`, errors `0.006`, p75 `940ms`, and zero duplicates.
   **Decision:** compare every difference with its prior limit.
   **After:** drops are `0.003`, `0.001`, and `40ms`; advance to the
   next bounded cohort. **Known correct:** all guardrails pass.
5. **Before:** at wider exposure, candidate duplicate-order count
   becomes `1`. **Decision:** halt expansion and roll back because the
   old read path and data contract remain compatible. **After:** new
   assignments stop and existing evidence is preserved.
   **Known correct:** a domain invariant overrides healthy aggregates.
6. **Before:** the identity defect is fixed and the same gates later
   pass through full rollout. **Decision:** observe the defined
   stability window, remove the legacy branch and flag, then remove
   their tests and telemetry. **After:** the query cache is the sole
   read owner. **Known correct:** the migration has a verified end.

Second input: rename API field `totalCents` to `totalMinor` while old
tabs and a service worker may use release `r51` for seven supported
days. Expected output: both releases work during that window and the
old field is removed only after measured old use reaches zero for the
defined safety window.

1. **Before:** the server returns only `totalCents` and old clients
   require it. **Decision:** add `totalMinor` with the same integer
   and retain `totalCents`. **After:** old and new readers receive
   their expected field. **Known correct:** expansion is additive.
2. **Before:** `r52` understands both fields. **Decision:** read
   `totalMinor`, fall back to `totalCents`, and report which field was
   used without logging order content. **After:** a new client can
   tolerate an older response during rollout. **Known correct:** both
   directions of temporary version skew are covered.
3. **Before:** new clients dominate, but an `r51` tab submits on day
   six. **Decision:** keep the old response field through the declared
   seven-day support window and safety interval. **After:** the old
   tab completes checkout. **Known correct:** deployment age did not
   stand in for consumer evidence.
4. **Before:** supported old-field use is zero for the safety window
   and caches are within policy. **Decision:** remove the fallback in
   a normal release, then remove server production of `totalCents` in
   a later release. **After:** only `totalMinor` remains.
   **Known correct:** readers stop depending on the field before the
   writer contracts it.

Failure boundary: if the migration had irreversibly rewritten stored
amounts with a different unit, switching application code back would
not restore meaning. The plan would halt new exposure, preserve the
new schema, and roll forward with a compatible reader or repair.

Third input: release `r60` keeps exception rate flat but lowers mobile
checkout completion from `0.980` to `0.920`. Expected output: declare
impact, halt the rollout, restore the prior experience, and produce
owned system improvements rather than blame the releaser.

1. **Before:** the alert identifies the mobile checkout cohort and
   `r60`; cause is unknown. **Decision:** name an incident lead, halt
   expansion, state confirmed impact, and begin a shared timeline.
   **After:** one response owner coordinates mitigation.
   **Known correct:** facts and hypotheses remain distinct.
2. **Before:** `r60` is reversible and the preceding contract is still
   valid. **Decision:** roll back, then verify completion recovery.
   **After:** the cohort returns near `0.980`. **Known correct:** user
   harm is limited before deeper diagnosis.
3. **Before:** a trace shows a new overlay covering the mobile submit
   button, while exception telemetry remains green. **Decision:** add
   a narrow-viewport workflow test and completion guardrail.
   **After:** prevention and detection both improve.
   **Known correct:** the correction targets system gaps that allowed
   the release, not the person who approved it.
4. **Before:** follow-up actions exist in the postmortem.
   **Decision:** assign owners, deadlines, and verifiable end states,
   then track them to closure. **After:** the incident becomes durable
   learning.
   **Known correct:** writing a document is not mistaken for reducing
   recurrence risk.

#### Why it works

At every rollout stage, the serving artifact, eligible cohort,
compatible representations, and next decision remain explicit.
Expansion makes version overlap safe before traffic moves. Bounded
exposure limits harm, and contract removal waits for evidence rather
than deployment age. Cleanup returns the system to one owner and one
normal path.

Predeclared gates prevent a promising result from changing the
standard mid-rollout. Candidate-versus-control comparison connects a
release to user impact, while domain invariants catch severe failures
that aggregate rates can hide. A tested rollback works only while its
code and data contracts remain compatible; naming irreversible effects
forces a roll-forward plan before they occur.

Leadership follows the same contract. A recorded decision preserves
context and review conditions. One owner can act after hearing
affected perspectives. Blameless incident learning exposes missing
checks, unsafe defaults, and organizational conditions; owned
follow-up then changes those conditions. Measures remain useful when
they diagnose a delivery system instead of becoming personal quotas.

#### TypeScript example

This evaluator implements the first rollout's gates. Rates are
fractions from `0` through `1`; p75 is milliseconds; sample and
duplicate counts are non-negative integers, while the control and
minimum sample counts must be positive. A breach rolls back only when
the preceding code and data contract are still safe to restore.

```typescript
type RolloutAction = "advance" | "halt" | "hold" | "rollback";

type CohortSignals = Readonly<{
  completionRate: number;
  duplicateOrderCount: number;
  errorRate: number;
  p75ConfirmationMs: number;
  sampleCount: number;
}>;

type RolloutPolicy = Readonly<{
  maximumCompletionDrop: number;
  maximumErrorRateIncrease: number;
  maximumP75IncreaseMs: number;
  minimumCandidateSamples: number;
}>;

type RolloutDecision = Readonly<{
  action: RolloutAction;
  reasons: readonly string[];
}>;

function assertRate(name: string, value: number): void {
  const valid = Number.isFinite(value) && value >= 0 && value <= 1;
  if (!valid) throw new Error(`${name} must be from 0 through 1`);
}

function assertCount(name: string, value: number): void {
  const valid = Number.isInteger(value) && value >= 0;
  if (!valid) {
    const message = `${name} must be a non-negative integer`;
    throw new Error(message);
  }
}

function assertPositiveCount(name: string, value: number): void {
  assertCount(name, value);
  if (value === 0) throw new Error(`${name} must be positive`);
}

function assertDuration(name: string, value: number): void {
  const valid = Number.isFinite(value) && value >= 0;
  if (!valid) throw new Error(`${name} must be non-negative`);
}

function validateSignals(name: string, value: CohortSignals): void {
  assertRate(`${name} completion rate`, value.completionRate);
  assertRate(`${name} error rate`, value.errorRate);
  assertCount(`${name} samples`, value.sampleCount);
  assertCount(`${name} duplicates`, value.duplicateOrderCount);
  assertDuration(`${name} p75`, value.p75ConfirmationMs);
}

function evaluateRollout(
  control: CohortSignals,
  candidate: CohortSignals,
  policy: RolloutPolicy,
  rollbackSafe: boolean,
): RolloutDecision {
  validateSignals("control", control);
  validateSignals("candidate", candidate);
  const maximumCompletionDrop = policy.maximumCompletionDrop;
  const maximumErrorIncrease = policy.maximumErrorRateIncrease;
  const maximumP75Increase = policy.maximumP75IncreaseMs;
  const minimumSamples = policy.minimumCandidateSamples;
  assertRate("maximum completion drop", maximumCompletionDrop);
  assertRate("maximum error increase", maximumErrorIncrease);
  assertDuration("maximum p75 increase", maximumP75Increase);
  assertPositiveCount("control samples", control.sampleCount);
  assertPositiveCount("minimum candidate samples", minimumSamples);
  if (control.duplicateOrderCount > 0) {
    throw new Error("control baseline has duplicate orders");
  }

  const reasons: string[] = [];
  const controlCompletion = control.completionRate;
  const candidateCompletion = candidate.completionRate;
  const completionDrop = controlCompletion - candidateCompletion;
  if (completionDrop > maximumCompletionDrop) {
    reasons.push("completion rate regressed");
  }

  const errorIncrease = candidate.errorRate - control.errorRate;
  if (errorIncrease > maximumErrorIncrease) {
    reasons.push("error rate regressed");
  }

  const controlP75 = control.p75ConfirmationMs;
  const candidateP75 = candidate.p75ConfirmationMs;
  const p75Increase = candidateP75 - controlP75;
  if (p75Increase > maximumP75Increase) {
    reasons.push("confirmation time regressed");
  }

  if (candidate.duplicateOrderCount > 0) {
    reasons.push("duplicate order invariant failed");
  }

  // A small cohort can show a failure, but cannot prove safety.
  if (reasons.length > 0) {
    return {
      action: rollbackSafe ? "rollback" : "halt",
      reasons,
    };
  }
  if (candidate.sampleCount < minimumSamples) {
    return { action: "hold", reasons: ["insufficient samples"] };
  }
  return { action: "advance", reasons: [] };
}

const control: CohortSignals = {
  completionRate: 0.98,
  duplicateOrderCount: 0,
  errorRate: 0.005,
  p75ConfirmationMs: 900,
  sampleCount: 10_000,
};
const candidate: CohortSignals = {
  completionRate: 0.977,
  duplicateOrderCount: 0,
  errorRate: 0.006,
  p75ConfirmationMs: 940,
  sampleCount: 2_000,
};
const policy: RolloutPolicy = {
  maximumCompletionDrop: 0.005,
  maximumErrorRateIncrease: 0.003,
  maximumP75IncreaseMs: 100,
  minimumCandidateSamples: 1_000,
};

const healthy = evaluateRollout(control, candidate, policy, true);
console.assert(healthy.action === "advance");

const smallCandidate: CohortSignals = {
  ...candidate,
  sampleCount: 600,
};
const held = evaluateRollout(control, smallCandidate, policy, true);
console.assert(held.action === "hold");

const duplicateCandidate: CohortSignals = {
  ...candidate,
  duplicateOrderCount: 1,
};
const unsafeDecision = evaluateRollout(
  control,
  duplicateCandidate,
  policy,
  true,
);
console.assert(unsafeDecision.action === "rollback");
```

#### Work and memory

The evaluator performs a fixed number of comparisons, so its work and
additional memory are `O(1)`. A real analysis spends work proportional
to releases, cohorts, routes, metrics, and retained samples or
aggregates. Collection, experiment assignment, dashboards, and
retention consume browser, network, storage, and analysis capacity.

Safe migrations temporarily duplicate readers, writers, schemas,
tests, controls, telemetry, documentation, and operational attention.
Backfills add read and write load; shadow traffic can nearly duplicate
expensive work. Canary and blue-green capacity may be intentionally
idle. Bound rate, concurrency, queue growth, storage, observation
time, and compatibility lifetime.

Leadership work consumes design, review, communication, mentoring,
incident, and follow-up time. That is part of delivery cost, not
overhead that disappears when omitted. Small batches and reusable
automation spend effort earlier to reduce repeated coordination,
recovery time, and blast radius.

#### When to use it

Use progressive delivery when a plausible failure affects users,
revenue, privacy, accessibility, reliability, stored data, or several
teams. Use expand-migrate-contract when producers and consumers cannot
change atomically, especially for APIs, browser persistence, service
workers, events, design systems, and authentication.

Use a temporary feature control when deployment and exposure need
different timing or a fast disable path. Use a written decision when
context, tradeoffs, consequences, or ownership outlive one meeting.
Use an incident review when user impact, data risk, response effort,
or missed detection crosses a predeclared threshold.

#### When not to use it

Do not add a flag, canary, or migration framework to a trivial,
immediately reversible change with no meaningful exposure risk. A
normal automated release and verification may be safer than permanent
control-plane complexity.

Do not use client controls for authorization, use an experiment to
decide release safety, keep dual paths because deletion feels risky,
or call deployment success a completed migration. Do not use consensus
as a substitute for a named accountable decision or use metrics to
compare individuals working in different systems.

#### Common mistakes and edge cases

> **Warning:** A feature flag without an owner and removal condition
> is permanent hidden architecture.

- Rebuilding between canary and full rollout changes the artifact.
- A canary cohort of employees may exclude real device and data risk.
- Too few samples can look healthy because failure has not occurred.
- Watching only averages hides affected routes and slow cohorts.
- Measuring only exceptions misses abandonment and covered controls.
- Running candidate and control at different times confounds demand,
  dependencies, and network conditions with the release.
- Percentage assignment without stable identity can move one workflow
  between incompatible variants.
- Browser flags expose names and values; neither is a secret or an
  authorization decision.
- Provider failure without an explicit default can change behavior at
  the worst time.
- Testing only the enabled branch leaves rollback code unverified.
- Dual writes can succeed on one side and fail on the other.
- A shadow request that sends email or payment is not a shadow read.
- A fast backfill can exhaust database, network, or browser resources.
- Old tabs, cached HTML, service workers, and offline queues extend
  the compatibility window beyond deployment.
- Rollback may be unsafe after an incompatible data or API change.
- A rollback cannot unsend a notification or reverse every payment.
- A one-hundred-percent flag still retains two branches and policy.
- Global rollout hides regional, browser, device, and feature-pair
  regressions.
- Experiments answer product hypotheses; safety gates limit harm.
- Emergency processes still need authentication, audit, and cleanup.
- Reviews that stop at "human error" miss system corrections.
- Action items without owners and verifiable ends become good wishes.
- A decision record without consequences or review conditions becomes
  an archive, not a decision tool.
- A metric target can reward gaming instead of user and delivery
  improvement.
- Mentoring without delegated decisions creates dependence, not
  growth.
- Delegation without constraints or support creates abandonment.

#### Related material

- [Frontend architecture and design systems][architecture]
- [Modules, toolchains, and production builds][toolchain]
- [Testing, performance, and observability][quality]
- [Server data, routing, caching, and forms][data]
- [Browser storage, origins, and security][security]
- [Backend delivery and production reference][backend-delivery]
- [Senior frontend job preparation][job-plan]
- [DORA software delivery performance metrics](
  https://dora.dev/guides/dora-metrics/)
- [Google SRE canarying releases](
  https://sre.google/workbook/canarying-releases/)
- [Google SRE postmortem culture](
  https://sre.google/workbook/postmortem-culture/)
- [OpenFeature evaluation context](
  https://openfeature.dev/specification/sections/evaluation-context/)

[backend-delivery]: ../backend/README.md
[job-plan]: ../SENIOR_FRONTEND_JOB_PREP_PLAN.md

## Sources and maintenance

The explanations above stand on their own. These primary references
own details that evolve independently of this book:

- [HTML Living Standard](https://html.spec.whatwg.org/)
- [Web Content Accessibility Guidelines 2.2](
  https://www.w3.org/TR/WCAG22/)
- [MDN Web Docs](https://developer.mozilla.org/)
- [MDN Baseline compatibility](
  https://developer.mozilla.org/docs/Glossary/Baseline/Compatibility)
- [ECMAScript language specification](https://tc39.es/ecma262/)
- [HTTP Semantics, RFC 9110](https://www.rfc-editor.org/rfc/rfc9110)
- [HTTP Caching, RFC 9111](https://www.rfc-editor.org/rfc/rfc9111)
- [Fetch Standard](https://fetch.spec.whatwg.org/)
- [Storage Standard](https://storage.spec.whatwg.org/)
- [Trusted Types](https://www.w3.org/TR/trusted-types/)
- [OWASP XSS overview](
  https://owasp.org/www-community/attacks/xss/)
- [TypeScript documentation](https://www.typescriptlang.org/docs/)
- [TypeScript module theory](
  https://www.typescriptlang.org/docs/handbook/modules/theory.html)
- [npm package-lock documentation](
  https://docs.npmjs.com/files/package-lock.json/)
- [Rollup tree-shaking reference](https://rollupjs.org/faqs/)
- [TypeScript release blog](
  https://devblogs.microsoft.com/typescript/)
- [React documentation](https://react.dev/learn)
- [React versions and release links](https://react.dev/versions)
- [React releases and security advisories](https://react.dev/blog)
- [Core Web Vitals](https://web.dev/articles/vitals)

When refreshing the book, verify dated framework and compiler claims
against their release indexes. Verify web features against their
standards, Baseline status, and the product's own support matrix.
Security advisories override examples and version notes; upgrade or
mitigate before preserving an older teaching configuration.
