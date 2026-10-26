# Web Software Engineering Competency Map

The primary employment path is senior frontend engineering. Full-stack
and backend depth are extensions that make frontend decisions safer
and widen the available job pool.

Active lessons follow the [lesson standards](LESSON_STANDARDS.md).
Indexes, plans, and source references support navigation and are not
teaching entries.

## Learning order

1. Learn browser-facing foundations in the
   [frontend curriculum](frontend/README.md).
2. Build and defend a production-quality frontend capstone.
3. Practice algorithms through the
   [DSA curriculum](../dsa/README.md).
4. Add HTTP, data, security, and production depth through the
   [backend extension](backend/README.md).
5. Rehearse applications and interviews through the
   [senior frontend plan](../SENIOR_FRONTEND_JOB_PREP_PLAN.md).

## Competency levels

### Explain

Define the idea in plain language, name its assumptions, and predict
ordinary and failure behavior without notes.

### Implement

Build a small version with typed boundaries, tests, accessible user
states, and explicit error handling.

### Operate

Measure real behavior, diagnose a failure, deploy safely, and verify
recovery.

### Lead

Compare alternatives, record tradeoffs, make a reversible plan, review
another implementation, and explain the decision to engineers and
non-engineers.

## Frontend readiness

A job-ready senior frontend candidate can demonstrate:

- Semantic HTML, robust CSS, responsive design, and accessibility.
- JavaScript and TypeScript runtime and type-system behavior.
- Browser rendering, networking, storage, and security boundaries.
- React component, state, effect, and rendering decisions.
- Server-data handling, routing, forms, and failure states.
- Testing, performance measurement, and frontend observability.
- Architecture, design-system, delivery, and migration judgment.
- Clear ownership, mentoring, incident, and tradeoff stories.

## Full-stack extension

Full-stack readiness adds:

- HTTP contracts, authentication, authorization, and API design.
- Relational modeling, transactions, indexes, and concurrency safety.
- Business boundaries, background work, and distributed failure.
- Containers, CI/CD, observability, incident response, and recovery.

## Evidence portfolio

Keep evidence instead of completion claims:

- A deployed frontend with measured performance and accessibility.
- Unit, integration, and end-to-end tests covering failure paths.
- One architecture decision record with rejected alternatives.
- One production-style incident drill and written review.
- One API integration with cancellation, retries, and authorization.
- Timed coding and system-design recordings.
- Six concise senior-level career stories.

## Integrated capstone

Build an accessible Order Console that consumes the Order API used by
the backend extension. The user can browse products, place an order,
retry safely after uncertain failure, and view only authorized orders.

The capstone must demonstrate:

- Responsive and keyboard-accessible interfaces.
- Typed API contracts and runtime response validation.
- Explicit loading, empty, partial, error, and retry states.
- Authentication and object-level authorization behavior.
- Stable caching, invalidation, and request cancellation.
- Component, integration, and end-to-end tests.
- Performance budgets and real-user measurements.
- Logs, metrics, traces, release checks, and rollback evidence.

## Definition of complete

The guide is complete for a learner only when the learner can produce
the required evidence, explain why it is correct, identify its limits,
and repeat the work without copying a finished solution.
