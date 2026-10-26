# Senior Frontend Job Preparation Plan

Target date: Monday, November 2, 2026.

This plan prepares for senior frontend engineering applications in
Vietnam. It balances application materials, TypeScript and React,
browser knowledge, frontend system design, coding interviews, and
senior-level communication.

The schedule assumes about four focused hours on weekdays and six
hours on Saturday and Sunday. If less time is available, preserve the
application package, system-design practice, and mock interviews.

## Study sources

- Use the [frontend curriculum][frontend] as the owner of technical
  lessons referenced by this schedule.
- Use the [frontend capstone][capstone] to assess implementation,
  operations, and leadership evidence.
- Use the [algorithm guide][algorithms] for coding-interview methods.
- Use the [backend extension][backend] when targeting full-stack roles
  or strengthening API and production knowledge.

## Target outcome

By Monday morning:

- Have five to eight tailored applications ready to send.
- Present a clear senior-level career narrative.
- Explain important frontend concepts without relying on notes.
- Complete a medium coding problem in TypeScript in 35 minutes.
- Lead a 45-minute frontend system-design discussion.
- Support seniority claims with specific examples and results.
- Complete at least one full simulated interview loop.

## Preparation principles

- Prepare for at most two closely related role profiles.
- Prefer demonstrated depth over shallow coverage of many frameworks.
- Explain decisions, tradeoffs, and production consequences aloud.
- Use TypeScript for coding practice unless a target company states a
  different interview language.
- Convert vague experience claims into measurable evidence.
- Review mistakes after every exercise instead of merely recording
  completion.
- Apply before Monday when a suitable position has an earlier
  deadline.

## Monday, October 26: Define the target

### 1. Build the role list - 60 minutes

Create a spreadsheet containing 15 to 20 suitable openings. Include:

- Company, title, location, and work arrangement.
- Application deadline.
- Required and preferred technologies.
- Years of experience requested.
- English and time-zone expectations.
- Product domain.
- Personal fit: strong, partial, or missing.
- Priority: A, B, or C.
- Application and follow-up status.

Choose no more than two target profiles. Suitable examples are:

1. Senior React and TypeScript product engineer.
2. Senior frontend or frontend-focused full-stack engineer.

Do not prepare independently for React, Vue, Angular, mobile, and
general full-stack roles during the same week.

### 2. Extract common requirements - 45 minutes

Count how often target listings mention:

- React, TypeScript, and Next.js.
- State management and server-data handling.
- Testing, accessibility, performance, and security.
- Architecture, code review, mentoring, and ownership.
- REST, GraphQL, CI/CD, cloud, and observability.
- Spoken and written English.

Create a one-page scorecard. Mark each subject:

- Green: can explain and demonstrate it.
- Yellow: understands it but lacks a clear example.
- Red: cannot yet explain it confidently.

### 3. Run a baseline interview - 90 minutes

Without notes:

1. Give a two-minute career introduction.
2. Solve one medium array or map problem in TypeScript.
3. Design a dashboard with authentication, server data, filters,
   tables, loading states, and error handling.
4. Answer: "Tell me about a difficult technical decision you led."

Record the exercise. Note where the explanation becomes vague, slow,
or unnecessarily detailed.

### 4. Define the positioning statement - 30 minutes

Complete this sentence:

> Senior frontend engineer with X years of experience, specializing in
> ___, who has delivered ___ and improved ___.

### Monday deliverables

- Target-company list.
- Skills scorecard.
- Baseline recording.
- Positioning statement.

## Tuesday, October 27: Build the application package

### 1. Rewrite the CV - 120 minutes

Keep the CV to one or two pages. The first half of page one should
establish seniority immediately.

Use result-oriented bullets:

> Led `<technical action>` across `<scope>`, improving `<metric>` by
> `<result>`.

Prioritize evidence involving:

- Features or systems owned from discovery through production.
- Performance, reliability, conversion, or delivery improvements.
- Architectural decisions and migrations.
- Mentoring, code review, or engineering standards.
- Incidents, debugging, or production operations.
- Work with product, design, backend, and business stakeholders.

Remove long responsibility lists. Avoid weak phrases such as "helped
with," "participated in," and "responsible for."

### 2. Prepare two CV variants - 45 minutes

- Product version: emphasize user experience, accessibility,
  experimentation, performance, and cross-functional delivery.
- Platform version: emphasize architecture, design systems, testing,
  APIs, reliability, and engineering standards.

### 3. Update professional profiles - 45 minutes

Ensure that the CV and online profiles agree on dates, titles, and
achievements. Update:

- Headline and summary.
- Location and work preferences.
- Technology list.
- Portfolio and GitHub links.
- English proficiency.
- Notice period or availability, where appropriate.

Pin two or three repositories that demonstrate quality. Each should
have a concise README that explains the problem, architecture,
tradeoffs, tests, and setup steps.

Do not build a new portfolio application this week unless nothing
presentable exists already.

### 4. Prepare recruiter answers - 30 minutes

Write concise answers for:

- Why are you leaving your current or previous role?
- Why does this role interest you?
- What kind of team are you seeking?
- What are your availability and notice period?
- What work arrangement do you prefer?
- What gross monthly compensation do you expect?
- What experience do you have working in English?

### Tuesday deliverables

- Master CV.
- Product and platform CV variants.
- Updated professional profiles.
- Recruiter-answer sheet.

## Wednesday, October 28: JavaScript and the browser

### 1. Review JavaScript - 75 minutes

Explain each topic aloud and provide a small example:

- Scope, closures, hoisting, and `this`.
- Reference and value behavior.
- Promises, `async` and `await`, error propagation, and cancellation.
- Event loop, tasks, microtasks, and rendering.
- Debouncing, throttling, memoization, and caching.
- Modules, tree shaking, and code splitting.

### 2. Review TypeScript - 60 minutes

Practice:

- Unions, narrowing, generics, and utility types.
- Discriminated unions and exhaustive checking with `never`.
- `unknown` compared with `any`.
- Typing API states and component variants.
- Runtime validation compared with compile-time checking.
- Designing types that make invalid states difficult to represent.

### 3. Review browser fundamentals - 60 minutes

Be able to explain:

- DNS, TLS, HTTP, caching, cookies, and CORS.
- DOM and CSSOM construction, layout, paint, and composition.
- Cookies, local storage, session storage, and IndexedDB.
- XSS, CSRF, content security policy, and authentication safety.
- Core Web Vitals and investigation of a slow page.

### 4. Complete coding practice - 60 minutes

Use these repository problems as prompts, but implement them in
TypeScript:

- [Group Anagrams][group-anagrams]
- [Minimum Size Subarray Sum][minimum-subarray]

For each problem:

1. Restate the problem and clarify assumptions.
2. Describe the direct solution before optimizing it.
3. Identify the governing invariant.
4. Implement the solution within 30 to 35 minutes.
5. Test important edge cases aloud.
6. State time and auxiliary-space complexity.

### Wednesday deliverables

- One-page JavaScript and browser review sheet.
- Two timed TypeScript solutions.
- List of subjects that still require prompting.

## Thursday, October 29: React and practical frontend work

### 1. Review React behavior - 90 minutes

Review and explain:

- Render and commit phases.
- Reconciliation and the purpose of keys.
- State snapshots and batched updates.
- Effect synchronization, cleanup, and stale closures.
- Controlled and uncontrolled inputs.
- Context, composition, reducers, and external stores.
- Memoization and its costs.
- Error boundaries, Suspense, lazy loading, and transitions.
- Server rendering, hydration, and client boundaries.
- Server state, local state, and URL state.

### 2. Build one realistic feature - 120 minutes

Use an existing project or a small local exercise. Build a searchable,
paginated data view with:

- Loading, empty, success, and error states.
- Request cancellation or stale-response protection.
- Accessible keyboard and screen-reader behavior.
- Filters synchronized with the URL.
- A responsive layout.
- Unit or component tests.
- One end-to-end happy-path test.
- Clear component and state boundaries.

The goal is engineering judgment, not visual decoration.

### 3. Review the implementation - 45 minutes

Review the feature as if it were a pull request. Look for:

- Incorrect or unnecessary effects.
- Avoidable renders.
- Race conditions.
- Missing error handling.
- Accessibility failures.
- Weak type modeling.
- Unnecessary abstractions.
- Missing or low-value tests.

### Thursday deliverables

- One demonstrable feature.
- Passing automated tests.
- Written self-review with corrective actions.

## Friday, October 30: Frontend system design

### 1. Learn a repeatable framework - 45 minutes

Address design questions in this order:

1. Users, requirements, constraints, and success measures.
2. Page and component boundaries.
3. Data model and API contracts.
4. Server state, local state, and URL state.
5. Loading, retries, errors, offline use, and concurrency.
6. Rendering strategy and caching.
7. Accessibility and responsive behavior.
8. Performance budgets and measurement.
9. Authentication and frontend security.
10. Testing, deployment, observability, and rollout.
11. Tradeoffs and future evolution.

### 2. Complete one design exercise - 60 minutes

Design one of the following:

- Analytics dashboard with live updates.
- Multi-step checkout.
- Large searchable product catalog.
- Reusable company design system.

Draw the architecture and explain every major decision aloud.

### 3. Practice depth questions - 60 minutes

Prepare concise answers for:

- Why not place all state in one store?
- How are duplicate or stale requests prevented?
- How does the product behave on poor connections?
- How can the JavaScript bundle be reduced?
- How can server rendering be introduced safely?
- How is real-user performance measured?
- How can teams share components without blocking one another?
- How can the architecture be migrated incrementally?

### 4. Prepare one architecture story - 45 minutes

Describe a real decision using:

- Context and constraints.
- Options considered.
- Chosen approach and reasoning.
- Risk controls.
- Measured outcome.
- What would be changed now.

### Friday deliverables

- Architecture diagram.
- Recorded design explanation.
- One real architecture story.

## Saturday, October 31: Stories and mock interviews

### 1. Prepare six STAR stories - 120 minutes

Use Situation, Task, Action, Result, and Learning. Prepare these
stories:

- High-impact delivery.
- Technical disagreement.
- Production incident or failure.
- Ambiguous requirements.
- Mentoring or raising team quality.
- A decision that would be handled differently today.

Each story should take two to three minutes. Clearly distinguish
personal actions from team actions and include measured results where
possible.

### 2. Complete a coding mock - 60 minutes

Use one of these prompts:

- [Merge Intervals][merge-intervals]
- [Number of Islands][number-of-islands]

Allow 45 minutes for the interview and 15 minutes for critique.

### 3. Complete a frontend mock - 75 minutes

Ask another person to interview you, or record answers to:

- Diagnose a React rendering or effect defect.
- Design state and API handling for a complex screen.
- Explain a frontend performance problem.
- Review a flawed component.
- Describe an appropriate testing strategy.

### 4. Complete an English behavioral mock - 60 minutes

Conduct the mock entirely in English. Check for:

- Direct answers before supporting detail.
- Clear personal ownership.
- Appropriate technical depth.
- Limited and relevant background.
- Concrete outcomes.

### 5. Correct the largest weaknesses - 60 minutes

Fix the three most visible failures from the mock sessions. Do not
begin another broad subject.

### Saturday deliverables

- Six STAR stories.
- Three mock recordings.
- Prioritized weakness list.
- Corrective practice for the three largest weaknesses.

## Sunday, November 1: Rehearse and tailor applications

### 1. Run a complete interview loop - 180 minutes

Simulate:

- 30-minute recruiter and career discussion.
- 45-minute TypeScript coding interview.
- 60-minute frontend system-design interview.
- 45-minute React and browser discussion.

Use normal interview conditions: no notes, no AI assistance, a visible
timer, and verbal reasoning throughout.

### 2. Review observed weaknesses - 60 minutes

Classify every failure:

- Knowledge gap.
- Slow recall.
- Weak explanation.
- Incorrect assumption.
- Implementation error.
- Missing edge case.

Correct the highest-frequency category instead of starting another
large topic.

### 3. Prepare priority applications - 90 minutes

For every priority-A company:

- Tailor the summary and strongest experience bullets.
- Use relevant terminology without copying the listing.
- Write a concise application note.
- Identify a recruiter, employee, or referral path.
- Prepare three company-specific reasons for applying.
- Prepare five thoughtful interview questions.

Useful questions include:

- What decisions does the senior frontend engineer own?
- What currently limits frontend delivery or reliability?
- How are architecture decisions recorded and reviewed?
- How are performance and accessibility measured?
- What would successful impact look like after six months?

### 4. Prepare the environment and rest - 30 minutes

Check the interview location, camera, microphone, internet backup,
IDE, Node version, and browser. Stop intensive preparation in the
evening and get normal sleep.

### Sunday deliverables

- Completed full mock loop.
- Reviewed and corrected mock failures.
- Five to eight tailored applications ready to send.
- Verified interview environment.

## Monday, November 2: Execute

### Morning review

- Review the introduction, design framework, and STAR headings.
- Solve one familiar easy problem for no more than 20 minutes.
- Proofread the CV and every application link once.
- Do not start a new framework or algorithm.

### Submit applications

- Send priority-A applications first.
- Record every submission and follow-up date.
- Send concise referral requests to genuine connections.
- Tailor each application instead of using one generic CV.

### Continue until interviews begin

Schedule two recurring daily blocks:

- 45 minutes of TypeScript coding or frontend debugging.
- 45 minutes of React, browser, design, or behavioral practice.

## Readiness gate

Application and interview readiness is reasonable when all of the
following are true:

- [ ] Five to eight tailored applications are ready.
- [ ] The CV demonstrates senior ownership and measurable impact.
- [ ] The two-minute introduction sounds natural.
- [ ] Six STAR stories can be delivered without notes.
- [ ] A medium array or map problem can be solved in TypeScript within
      35 minutes.
- [ ] React rendering, effects, state, and testing can be explained.
- [ ] Browser rendering, performance, accessibility, and basic web
      security can be explained.
- [ ] A frontend system-design discussion can be led for 45 minutes.
- [ ] One full mock loop and one additional coding mock are complete.
- [ ] The reason for applying to each priority company is specific.
- [ ] Compensation, availability, work arrangement, and notice-period
      answers are prepared.

## If time falls behind

Keep these items, in order:

1. CV, professional profiles, and tailored applications.
2. Career introduction and six STAR stories.
3. One complete mock interview loop.
4. Frontend system-design practice.
5. React, TypeScript, and browser review.
6. Two timed coding problems.

Drop additional repetitions, portfolio decoration, and low-priority
topics before dropping a complete mock or the application package.

[group-anagrams]: dsa/problems/medium_leetcode_0049_group_anagrams.go
[minimum-subarray]:
  dsa/problems/medium_leetcode_0209_minimum_size_subarray_sum.go
[merge-intervals]:
  dsa/problems/medium_leetcode_0056_merge_intervals.go
[number-of-islands]:
  dsa/problems/medium_leetcode_0200_number_of_islands.go
[frontend]: web-engineering/frontend/README.md
[capstone]: web-engineering/frontend/CAPSTONE.md
[algorithms]: dsa/CONCEPTS.md
[backend]: web-engineering/backend/README.md
