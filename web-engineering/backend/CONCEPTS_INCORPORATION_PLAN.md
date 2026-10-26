# Backend Concepts Incorporation Plan

## Objective

Incorporate every distinct teaching idea from the 18 chapters under
`reference/` into `CONCEPTS.md` without losing nuance, duplicating
ownership, breaking existing links, or weakening the active lesson
standards. After repository-owned examples and provenance records
replace every live dependency, remove `reference/`.

The current competency map is a broad synthesis. This plan replaces
summary-level confidence with source-to-destination evidence for every
idea.

## Governing standards

Execute every phase under the repository's
[shared concept study standards][concept-standards]. The
[web lesson profile][web-standards] adds backend-specific adaptations;
it does not replace or weaken the shared contract.

The current working-tree versions of the guide and standards are the
source of truth. Before each batch:

1. Re-read the shared standards and web lesson profile.
2. Re-read every complete entry that the batch will change.
3. Read [the code-generation rules][coding-rules] before changing a
   Go example.
4. Inspect repository status and the relevant diff.
5. Preserve unrelated and user-authored work.
6. Identify current owners, inbound links, anchors, code names, and
   source references before changing structure.

The standards distinguish supporting material from active teaching
entries. A competency map, navigation group, source index, or evidence
section may stay concise. Any material promoted into a standalone
active concept must satisfy the complete teaching-entry contract.

[concept-standards]: ../../CONCEPTS_STANDARDS.md
[web-standards]: ../LESSON_STANDARDS.md
[coding-rules]: ../../CODING_RULES.md

## Phase execution contract

This is a phase-gated effort. Work on one phase at a time and do not
begin the next phase until the current exit gate passes.

For every phase:

1. Confirm its required inputs exist and reflect the current worktree.
2. Perform only the work authorized by that phase.
3. Produce the named deliverables.
4. Run the phase validation and record the evidence.
5. Fix failures within the same phase.
6. Mark the phase complete only when every exit condition passes.

If later evidence invalidates an earlier decision, return to the
owning phase, correct its deliverables, and re-run every affected gate
in order. Do not patch the contradiction only in a later phase.

Initial phase status:

- Phase 1: complete; the source inventory contains 237 records.
- Phase 2: complete; every record has one canonical owner.
- Phase 3: complete; the final guide design has 32 active entries.
- Phase 4: complete; every entry has an eleven-part specification.
- Phase 5: complete; all seven topic batches passed their local gates.
- Phase 6: complete; all 237 records have terminal traceability.
- Phase 7: complete; live source links retired and the source chapter
  directory removed.
- Phase 8: complete; post-removal audit and validation pass.

## Completion standard

The work is complete when:

- every distinct source idea has one recorded destination in
  `CONCEPTS.md`;
- repeated ideas have one canonical owner and useful cross-links;
- unique constraints, tradeoffs, failure cases, and operational
  details remain represented;
- every active teaching entry follows the required eleven-part lesson
  contract;
- all plain-language, trace, correctness, code, cost, ownership,
  navigation, evidence, and maintenance requirements in the shared
  standards are satisfied;
- existing inbound fragments continue to resolve;
- no live link or implementation owner depends on `reference/`;
- the `reference/` directory has been removed;
- changed Go examples compile and cover an ordinary and failure case;
- all repository documentation and test checks pass; and
- a reviewer can trace any source section to its active-guide owner.

## Phase 1: Build the source inventory

### Inputs

- All 18 files under `reference/`.
- The current backend guide, index, provenance document, shared
  standards, and web lesson profile.

### Work

1. Read each reference chapter completely.
2. Record every distinct idea, not only each heading.
3. Give each item a stable inventory identifier.
4. Record its source file, heading, and line or paragraph location.
5. Classify the item as one of:
   - principle;
   - mechanism;
   - invariant;
   - procedure;
   - tradeoff;
   - failure mode;
   - test or experiment;
   - operational control;
   - completion evidence; or
   - language-specific behavior.
6. Record prerequisites and closely related items.

The inventory should use this shape:

```text
ID:
Source:
Source section:
Category:
Idea:
Required nuance:
Related IDs:
Destination:
Status:
```

### Deliverables

- One source-idea inventory containing every distinct idea.
- A chapter checklist showing that all 18 files were read.
- An ambiguity list for items whose meaning or scope is unresolved.

### Exit gate

Phase 1 is complete only when:

- every source chapter is checked off;
- every identified idea has a stable identifier and source location;
- every idea has a category and recorded nuance;
- repeated text remains separate until Phase 2 decides ownership; and
- no unresolved ambiguity could change inventory scope.

## Phase 2: Reconcile repetition and ownership

### Inputs

- The completed Phase 1 inventory.
- The current active guide and its existing teaching owners.

### Work

1. Group inventory items that express the same underlying idea.
2. Preserve meaningful differences in scope or failure behavior.
3. Select one canonical teaching owner for every group.
4. Mark each repeated source item as merged into that owner.
5. Keep a separate entry when combining ideas would hide a distinct
   invariant, cost, or decision.
6. Replace duplicate explanations with short relationship summaries
   and links.

No item may disappear merely because another chapter uses similar
terminology.

### Deliverables

- A canonical idea list.
- A merge map from every repeated item to its canonical owner.
- A list of ideas that require independent teaching ownership.

### Exit gate

Phase 2 is complete only when:

- every inventory item is unique or belongs to one documented group;
- each group preserves differences in invariant, scope, cost, and
  failure behavior;
- every canonical idea has exactly one proposed owner; and
- no merge relies only on similar wording.

## Phase 3: Design the active-guide structure

### Inputs

- The canonical idea and ownership maps from Phase 2.
- The current guide structure, table of contents, and inbound links.
- The backend index and applicable standards.

### Work

Preserve the backend progression defined by the current guide, the
backend index, and the web lesson profile. Within that progression,
organize canonical owners in this dependency order:

1. Learning model and backend operating environment.
2. Go language and backend mechanics.
3. Domain discovery and modeling.
4. Architecture and service construction.
5. HTTP contracts and process lifecycle.
6. PostgreSQL, relational data, and transactions.
7. Identity, authorization, and API security.
8. Testing, observability, and operations.
9. Delivery, performance, and recovery.
10. Durable work and event delivery.
11. Distributed-systems mechanics.
12. Service boundaries and distributed workflows.
13. Scaling, tenancy, privacy, and supply-chain security.
14. Cloud infrastructure and platform operations.
15. Reliability, incidents, and cost.
16. Production integrations.
17. Legacy evolution and engineering leadership.
18. Language portability.

This list is a dependency map, not automatic authorization to create
18 top-level lessons. Use supporting groups for navigation and create
an active entry only when the subject needs an independent teaching
owner.

Before renaming or moving an existing entry, confirm that its new
owner is complete, search for inbound links, and preserve the old
fragment with an explicit anchor when required. Keep the table of
contents synchronized with the final section names, order, and
fragments.

### Deliverables

- A proposed guide outline in dependency order.
- A classification of each section as supporting material or an
  active teaching entry.
- An anchor and inbound-link migration map.
- A primary owner for every canonical idea.

### Exit gate

Phase 3 is complete only when:

- the outline preserves the required backend progression;
- prerequisites appear before entries that depend on them;
- each active entry has a coherent and non-overlapping purpose;
- supporting sections do not masquerade as incomplete lessons;
- every existing inbound fragment has a preservation plan; and
- the table-of-contents design matches the proposed headings.

## Phase 4: Specify each canonical concept

### Inputs

- The approved structure and ownership map from Phase 3.
- The inventory items assigned to each destination.
- Existing examples, code owners, sources, and related material.

### Work

Write an implementation specification for every active teaching
entry. Apply the exact current eleven-part contract from the shared
standards and the literal `Go example` heading required by the web
lesson profile.

Each entry must also include:

- defined terminology before reliance on it;
- exact inputs and expected outputs for examples;
- replayable Before, Decision, After, and Known correct traces;
- at least one relevant failure, boundary, duplicate, or adversarial
  case;
- explicit correctness or lifecycle reasoning;
- latency, memory, network, storage, and server-load costs when they
  matter;
- the smallest complete Go example or a link to its code owner; and
- direct links to the detailed source chapters it incorporates.

Apply these additional shared requirements:

- lead with plain language and define formal terms before relying on
  them;
- use numbered lists for changing traces and avoid Markdown tables in
  teaching entries;
- use concrete warnings tied to a failing input or state;
- state the assumptions that correctness or lifecycle behavior needs;
- replay each trace against its code and explain why the same result
  follows;
- explain costs in words before notation and count every material
  runtime, network, storage, allocation, and server cost;
- keep full implementations with their established owner and link to
  them instead of copying them; and
- preserve useful examples, sources, attributions, links, and legacy
  anchors unless a documented correction requires removal.

### Deliverables

Each active entry receives a specification containing:

- its purpose, audience problem, and prerequisite entries;
- the source inventory identifiers it owns;
- the intended content for all eleven required sections;
- its ordinary and failure examples;
- its trace state and correctness or lifecycle argument;
- its Go example owner and validation approach;
- its cost dimensions and assumptions;
- its warnings, alternatives, and related links; and
- the anchors it creates or preserves.

### Exit gate

Phase 4 is complete only when:

- every active entry has a complete specification;
- every canonical idea appears in one destination specification;
- examples are concrete and small enough to replay;
- code ownership avoids duplicate implementations;
- cost and correctness obligations are explicit; and
- no implementation batch would need to invent missing scope.

## Phase 5: Implement in verifiable topic batches

### Inputs

- All completed entry specifications from Phase 4.
- The current worktree versions of every destination document.

### Work

Implement these topic batches sequentially:

1. Go and operating foundations.
2. Domain, architecture, HTTP, and persistence.
3. Identity, testing, observability, and delivery.
4. Idempotency, jobs, events, and distributed mechanics.
5. Service extraction, workflows, and contract evolution.
6. Scale, tenancy, security, cloud, platform, and reliability.
7. Integrations, maintenance, leadership, and portability.

For each batch:

1. Select the applicable inventory items.
2. Re-read their source passages and current destination entries.
3. Decide which items are supporting material and which require an
   active teaching owner.
4. Update the active guide without modifying unrelated material.
5. Finish all eleven sections of every changed active entry.
6. Format, compile, and exercise changed Go examples.
7. Update inventory destinations and statuses.
8. Re-read the prose as a beginner and perform a separate technical
   correctness and cost review.
9. Run focused documentation checks.
10. Review the diff before beginning the next batch.

Do not start a later batch while an earlier batch has a partial entry,
failed validation, or unresolved ownership question.

### Deliverables

- A standards-compliant `CONCEPTS.md` implementation for each batch.
- Compiled and exercised changed Go examples.
- Per-batch validation evidence and reviewed diffs.
- Updated inventory destinations and incorporation statuses.

### Batch exit gate

Each batch is complete only when:

- every changed active entry contains all eleven sections;
- its trace, code, correctness or lifecycle argument, and costs agree;
- ordinary and relevant failure behavior are exercised;
- focused structure, link, line-width, and test checks pass;
- no unrelated work changed; and
- the batch diff has no unresolved finding.

### Phase exit gate

Phase 5 is complete only after all seven batches pass their gates and
no specified active entry remains partial or unimplemented.

## Phase 6: Maintain the traceability matrix

### Inputs

- The completed Phase 1 inventory and Phase 2 ownership map.
- The fully implemented guide from Phase 5.

### Work

The final matrix must prove both directions:

- every source inventory item points to one destination anchor; and
- every added active-guide claim points back to supporting source
  material or an established repository standard.

Allowed statuses are:

- `pending`;
- `incorporated`;
- `merged`, with the owning inventory identifier;
- `reference-only`, with a written reason; or
- `not-an-idea`, for navigation or repeated exercise instructions.

Any `reference-only` decision requires review. Exercises may remain
in the source material, but their reusable teaching idea or evidence
requirement must still appear in the active guide.

The matrix must also record whether each destination is supporting
material or an active teaching entry. An active destination is not
complete until all shared-standard requirements have passed review.

### Deliverables

- A completed source-to-destination traceability matrix.
- A destination-to-source audit for every added guide claim.
- A reviewed list of all merged and reference-only decisions.

### Exit gate

Phase 6 is complete only when:

- every inventory item has a terminal status;
- no active entry depends on an untracked source idea;
- every merge names its canonical owner;
- every reference-only item has an accepted reason; and
- spot checks in both directions resolve to exact source and guide
  locations.

## Phase 7: Synchronize navigation and owning documents

### Inputs

- The completed guide and traceability matrix.
- The current backend index, provenance document, and inbound links.

### Work

After the guide structure is stable:

1. Update the `CONCEPTS.md` contents list.
2. Update the active-guide section in `README.md`.
3. Keep `SOURCE.md` accurate about provenance and ownership.
4. Update reciprocal links only when they improve navigation.
5. Replace every live source-chapter link or code owner with an active
   guide anchor or repository-owned example.
6. Remove `reference/` only after a repository search proves no live
   dependency remains.

### Deliverables

- Synchronized contents and backend index navigation.
- Accurate provenance and ownership statements.
- Preserved or updated reciprocal links and legacy fragments.

### Exit gate

Phase 7 is complete only when:

- every contents link and local path resolves;
- every repository inbound fragment resolves;
- the index and provenance documents describe the final structure;
- active and reference ownership remain unambiguous; and
- `reference/` is absent and no repository link targets it;
- no navigation update duplicates teaching content.

## Phase 8: Final audit and validation

### Inputs

- All completed deliverables from Phases 1 through 7.
- The final working-tree versions of the guide and related documents.

### Work

Perform three separate audits.

### Coverage audit

- Confirm that every inventory item has a terminal status.
- Review every merged group for lost nuance.
- Search for source terminology that has no destination occurrence.
- Confirm all 18 reference chapters have represented ideas.
- Confirm every former source-chapter code owner has a tested
  repository-owned replacement.

### Teaching-quality audit

- Read every changed entry from a beginner's perspective.
- Check terminology, traces, correctness, costs, examples, and links.
- Confirm every active entry uses the eleven exact headings in order.
- Confirm supporting sections have not become incomplete lessons.
- Replay changed traces against changed Go examples.
- Recalculate preprocessing, allocation, network, storage, and server
  costs where they apply.
- Confirm repeated prose has one owner.
- Confirm the progression introduces prerequisites before use.

### Repository validation

Run from the repository root:

```bash
~/.codex/scripts/check-non-ascii.py \
  CONCEPTS_STANDARDS.md \
  dsa/CONCEPTS.md \
  web-engineering/frontend/CONCEPTS.md \
  web-engineering/backend/CONCEPTS.md \
  web-engineering/backend/CONCEPTS_INCORPORATION_PLAN.md
git diff --check HEAD -- \
  CONCEPTS_STANDARDS.md \
  dsa/CONCEPTS.md \
  web-engineering/frontend/CONCEPTS.md \
  web-engineering/backend/CONCEPTS.md \
  web-engineering/backend/CONCEPTS_INCORPORATION_PLAN.md
go run dsa/internal/check-concepts-doc.go
go run ./internal/linewidth
go vet dsa/internal/check-concepts-doc.go
go test dsa/internal/check-concepts-doc.go
go test ./...
```

Also compile and exercise each changed complete Go example in a
temporary package, then remove the temporary validation artifacts.

### Deliverables

- Coverage, teaching-quality, and technical audit results.
- Complete command output for the required validation sequence.
- A final handoff describing changes, evidence, and any accepted
  limitations.

### Exit gate

Phase 8 and the overall effort are complete only when:

- every earlier phase remains complete after the final diff;
- every coverage and teaching-quality audit item passes;
- every required validation command exits successfully;
- temporary validation artifacts are removed;
- the removed `reference/` directory has no remaining live consumer;
- the final diff contains only intended changes; and
- no ambiguity or failed check remains undisclosed.

## Deliverables

The completed effort should produce:

- a fully expanded `CONCEPTS.md`;
- a source-idea inventory and traceability matrix;
- a standards-conformance record for every changed active entry;
- synchronized `README.md` and `SOURCE.md` when their claims change;
- the safely removed `reference/` directory;
- passing validation evidence; and
- a short handoff listing incorporated topics, deliberate merges, and
  any decisions that still require human review.
