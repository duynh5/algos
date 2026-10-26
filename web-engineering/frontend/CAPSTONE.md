# Frontend Capstone Evidence

Build an accessible Order Console against the Order API described by
the backend extension. This is a project brief, not a teaching entry.

## Required workflows

- Browse, filter, and paginate products through URL state.
- Open a product through keyboard or pointer input.
- Add valid quantities to an order draft.
- Place an order with stable operation identity.
- Recover after an uncertain request outcome without duplication.
- View only orders authorized for the signed-in user.
- Sign out and remove private cached state.

## Required user states

- Initial, loading, empty, partial, ready, and failed.
- Field validation and server validation.
- Offline or unavailable dependency.
- Expired identity and forbidden resource.
- Slow request, canceled navigation, and successful retry.

## Required evidence

- Semantic structure and keyboard walkthrough.
- Automated accessibility check plus manual focus review.
- Unit tests for state transitions and boundary validation.
- Component tests for visible behavior.
- End-to-end tests for order success and uncertain retry.
- Performance budget with lab and real-user measurements.
- Error report connected to route and release without private data.
- Architecture decision record for state and rendering choices.
- Threat model for identity, payment-like actions, and stored data.
- Release plan, rollback rule, and one simulated incident review.

## Review questions

- Which state belongs to the URL, server cache, form, and component?
- What happens if navigation changes during every request?
- Which response fields are validated at runtime?
- How does the UI distinguish rejection from uncertain completion?
- What can an unauthorized user infer from errors or timing?
- Which metrics prove that the release improved user experience?
- Which choices would change at ten times the users or team size?

## Completion gate

Another engineer must be able to clone, run, test, inspect, and deploy
the project using repository documentation. The learner must explain
every major tradeoff without relying on framework slogans.

Related material:

- [Frontend curriculum](README.md)
- [Software engineering guide](../../README.md)
- [Backend extension](../backend/README.md)
