# Senior Frontend Engineering Interview Concepts

This guide is for frontend engineers preparing for senior-level
software engineering interviews. It assumes working knowledge of
HTML, CSS, JavaScript, TypeScript, and a component framework.

Coverage status: planned outline only. The entries below identify the
study scope; no lesson is complete yet.

## Contents

- [A. Web platform and browser fundamentals](
  #a-web-platform-and-browser-fundamentals)
  - A.1 Navigation, URLs, DNS, HTTP, and browser request lifecycle
  - A.2 Semantic HTML, document structure, and native controls
  - A.3 Forms, validation, submission, and file inputs
  - A.4 DOM trees, events, propagation, and delegation
  - A.5 CSS cascade, specificity, inheritance, and custom properties
  - A.6 Box model, containing blocks, and formatting contexts
  - A.7 Flexbox, Grid, positioning, and stacking contexts
  - A.8 Responsive design, media queries, and container queries
  - A.9 Style, layout, paint, compositing, and rendering pipelines
  - A.10 Event loop, tasks, microtasks, and rendering opportunities
  - A.11 Browser storage, cookies, caches, and service workers
  - A.12 Web Components, custom elements, and Shadow DOM
- [B. JavaScript and TypeScript](#b-javascript-and-typescript)
  - B.1 Values, types, coercion, equality, and property access
  - B.2 Scope, closures, execution contexts, and hoisting
  - B.3 Functions, `this`, binding, and higher-order functions
  - B.4 Objects, prototypes, classes, and composition
  - B.5 Arrays, maps, sets, iterators, and generators
  - B.6 Mutation, copying, structural sharing, and immutability
  - B.7 Promises, `async`/`await`, concurrency, and cancellation
  - B.8 Error handling, recovery, and asynchronous failures
  - B.9 ES modules, dependency graphs, and dynamic imports
  - B.10 Memory management, garbage collection, and leak patterns
  - B.11 TypeScript inference, narrowing, unions, and intersections
  - B.12 Generics, mapped types, conditional types, and utility types
  - B.13 Type-safe APIs, runtime validation, and unsafe boundaries
- [C. Component and user-interface engineering](
  #c-component-and-user-interface-engineering)
  - C.1 Declarative rendering and user interface as a state function
  - C.2 Component boundaries, contracts, props, slots, and composition
  - C.3 Identity, reconciliation, keys, and preserved state
  - C.4 State ownership, lifting state, and derived state
  - C.5 Reducers, state machines, and explicit transitions
  - C.6 Lifecycles, effects, cleanup, and synchronization
  - C.7 Stale closures, request races, and effect dependencies
  - C.8 Context, dependency injection, and external stores
  - C.9 Refs, imperative escape hatches, and third-party integration
  - C.10 Error boundaries, suspense, transitions, and async rendering
  - C.11 Controlled and uncontrolled inputs
  - C.12 Framework mechanics and portability
- [D. Application state and data flows](
  #d-application-state-and-data-flows)
  - D.1 Local, shared, URL, form, server, and persisted state
  - D.2 Routing, navigation, history, and deep linking
  - D.3 HTTP API contracts, REST, GraphQL, and generated clients
  - D.4 Requests, cancellation, timeouts, retries, and idempotency
  - D.5 Server-state caching, invalidation, and revalidation
  - D.6 Optimistic updates, rollback, and uncertain outcomes
  - D.7 Form state, client validation, and server validation
  - D.8 Pagination, infinite scrolling, filtering, and sorting
  - D.9 Polling, server-sent events, WebSockets, and synchronization
  - D.10 Offline behavior, persistence, and conflict resolution
  - D.11 Loading, empty, stale, partial, failed, and forbidden states
  - D.12 Authentication sessions and authorization-aware interfaces
- [E. Accessibility and inclusive design](
  #e-accessibility-and-inclusive-design)
  - E.1 Accessibility principles, WCAG, and conformance levels
  - E.2 Semantic HTML and the accessibility tree
  - E.3 Accessible names, descriptions, roles, and states
  - E.4 Keyboard interaction, focus order, and focus management
  - E.5 Forms, validation errors, instructions, and recovery
  - E.6 Dynamic content, live regions, dialogs, and overlays
  - E.7 Color, contrast, zoom, reflow, and text spacing
  - E.8 Images, audio, video, animation, and reduced motion
  - E.9 Pointer, touch, target size, and alternative input
  - E.10 Automated checks, manual testing, and assistive technology
- [F. Frontend performance](#f-frontend-performance)
  - F.1 Performance goals, budgets, and user-centered measurement
  - F.2 Core Web Vitals, field data, and laboratory data
  - F.3 Network waterfalls, connection setup, and resource priority
  - F.4 Critical rendering path and render-blocking resources
  - F.5 JavaScript bundles, code splitting, and tree shaking
  - F.6 Images, fonts, CSS, media, and responsive asset delivery
  - F.7 Main-thread work, long tasks, scheduling, and responsiveness
  - F.8 Rendering cost, memoization, virtualization, and containment
  - F.9 Caching, content delivery networks, and preloading
  - F.10 Memory leaks, detached nodes, and resource cleanup
  - F.11 Performance profiling, regression tests, and monitoring
- [G. Frontend security and privacy](
  #g-frontend-security-and-privacy)
  - G.1 Threat modeling, trust boundaries, and untrusted input
  - G.2 Cross-site scripting, injection, and safe rendering
  - G.3 Content Security Policy and Trusted Types
  - G.4 Cookies, sessions, tokens, and secure browser storage
  - G.5 Cross-site request forgery, CORS, and same-origin policy
  - G.6 Clickjacking, open redirects, and unsafe navigation
  - G.7 Third-party scripts, dependency risk, and supply chains
  - G.8 Sensitive data, telemetry, consent, and data minimization
  - G.9 Authentication user experience and authorization failures
- [H. Testing, debugging, and quality](
  #h-testing-debugging-and-quality)
  - H.1 Test strategy, confidence, scope, and maintenance cost
  - H.2 Unit tests for logic and state transitions
  - H.3 Component tests for behavior and user interaction
  - H.4 Integration, contract, and end-to-end tests
  - H.5 Visual regression and cross-browser testing
  - H.6 Deterministic asynchronous and network testing
  - H.7 Accessibility, performance, and security checks
  - H.8 Browser developer tools and systematic debugging
  - H.9 Error reproduction, source maps, and production diagnosis
  - H.10 Static analysis, types, linting, and code review
- [I. Frontend system design and architecture](
  #i-frontend-system-design-and-architecture)
  - I.1 Requirements, constraints, workloads, and quality attributes
  - I.2 Single-page and multi-page application architectures
  - I.3 Client, server, static, streaming, and hybrid rendering
  - I.4 Hydration, islands, server components, and client boundaries
  - I.5 Feature, route, layer, and package boundaries
  - I.6 State and data architecture at application scale
  - I.7 Design systems, tokens, component APIs, and governance
  - I.8 Micro-frontends, independent delivery, and integration cost
  - I.9 Backend-for-frontend and API adaptation boundaries
  - I.10 Scalability, reliability, graceful degradation, and recovery
  - I.11 Internationalization, localization, and bidirectional layout
  - I.12 Architecture decisions, migrations, and legacy evolution
- [J. Tooling, delivery, and operations](
  #j-tooling-delivery-and-operations)
  - J.1 Package managers, lockfiles, and dependency resolution
  - J.2 Compilers, transpilers, bundlers, and development servers
  - J.3 Monorepos, package boundaries, and build graphs
  - J.4 Environment configuration and build-time versus runtime data
  - J.5 Continuous integration, quality gates, and test selection
  - J.6 Static hosting, edge delivery, and cache invalidation
  - J.7 Release strategies, feature flags, canaries, and rollback
  - J.8 Browser support, progressive enhancement, and polyfills
  - J.9 Logs, errors, traces, metrics, and real-user monitoring
  - J.10 Incident response, runbooks, and operational ownership
- [K. Senior engineering judgment and leadership](
  #k-senior-engineering-judgment-and-leadership)
  - K.1 Product goals, user needs, and technical requirements
  - K.2 Tradeoffs, alternatives, and architecture decision records
  - K.3 Estimation, sequencing, risk reduction, and delivery planning
  - K.4 Code review, standards, and maintainability
  - K.5 Mentoring, delegation, feedback, and knowledge sharing
  - K.6 Cross-functional work with design, product, and backend teams
  - K.7 Accessibility, performance, security, and quality advocacy
  - K.8 Metrics, experiments, outcomes, and evidence-based decisions
  - K.9 Technical debt, migrations, deprecation, and simplification
  - K.10 Project narratives, conflict, failure, and incident learning
- [L. Frontend interview formats](#l-frontend-interview-formats)
  - L.1 JavaScript and TypeScript coding exercises
  - L.2 DOM, CSS, and browser behavior questions
  - L.3 Component implementation and user-interface coding
  - L.4 Frontend system-design interviews
  - L.5 Debugging and performance investigations
  - L.6 Accessibility and interface-quality reviews
  - L.7 Testing strategy and code-review exercises
  - L.8 Project deep dives and behavioral interviews
  - L.9 Algorithms and data structures for frontend roles
  - L.10 Interview communication and solution evaluation

## A. Web platform and browser fundamentals

## B. JavaScript and TypeScript

## C. Component and user-interface engineering

## D. Application state and data flows

## E. Accessibility and inclusive design

## F. Frontend performance

## G. Frontend security and privacy

## H. Testing, debugging, and quality

## I. Frontend system design and architecture

## J. Tooling, delivery, and operations

## K. Senior engineering judgment and leadership

## L. Frontend interview formats
