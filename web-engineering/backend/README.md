# Backend Engineering Extension

This extension teaches backend engineering with Go through one
evolving Order API. Active lessons follow the shared
[lesson standards](../LESSON_STANDARDS.md).

Frontend candidates should first follow the frontend path from the
[web engineering guide](../README.md). Use this extension to develop
full-stack range, understand the systems behind browser applications,
or prepare for backend-heavy roles.

Reading alone is not completion. Each chapter asks for implementation,
failure testing, explanation, and production evidence.

## Active concept guide

The single [backend concept guide](CONCEPTS.md) owns all active
lessons, matching the DSA curriculum's grouped teaching structure.

1. [HTTP API contracts][http-lesson]
2. [Data, transactions, and architecture][data-lesson]
3. [Production and distributed systems][production-lesson]

## Source reference

The complete substantive study material from `be-standard` is
preserved as reference material. It retains the source's sequence and
detail without being presented as a standards-compliant lesson set.

### Foundations

1. [Overview and competency map][overview]
2. [Go language refresher][go-refresher]
3. [Go backend foundations][go-backend]
4. [HTTP contracts and service lifecycle][http]
5. [PostgreSQL and relational data][postgresql]
6. [Business logic and architecture][architecture]
7. [Identity and API security][security]
8. [Testing, observability, and operations][operations]
9. [Delivery and production readiness][delivery]

### Advanced systems

10. [Domain discovery and modeling][domain]
11. [Build the service core][service-core]
12. [Make the service production-ready][production]
13. [Durable work and distributed foundations][distributed]
14. [Service extraction and distributed delivery][extraction]
15. [Scale, security, cloud, and operations][scale]
16. [Integrations, maintenance, and leadership][leadership]
17. [Key-concept review][review]
18. [Language portability][portability]

## Lab

The imported [Order API exercise](labs/order-api/) is intentionally
incomplete. Building it is part of the curriculum. The destination
repository's Go module owns compilation and testing.

See [migration provenance](SOURCE.md) for the precise source mapping
and excluded repository metadata.

[overview]: reference/foundations/00-overview-and-competency-map.md
[go-refresher]: reference/foundations/01-go-language-refresher.md
[go-backend]: reference/foundations/02-go-backend-foundations.md
[http]:
  reference/foundations/03-http-contracts-and-service-lifecycle.md
[postgresql]:
  reference/foundations/04-postgresql-and-relational-data.md
[architecture]:
  reference/foundations/05-business-logic-and-architecture.md
[security]: reference/foundations/06-identity-and-api-security.md
[operations]:
  reference/foundations/07-testing-observability-and-operations.md
[delivery]:
  reference/foundations/08-delivery-and-production-readiness.md
[domain]: reference/advanced/09-domain-discovery.md
[service-core]: reference/advanced/10-service-core.md
[production]: reference/advanced/11-production-service.md
[distributed]: reference/advanced/12-distributed-foundations.md
[extraction]: reference/advanced/13-service-extraction.md
[scale]: reference/advanced/14-scale-security-cloud-operations.md
[leadership]:
  reference/advanced/15-integrations-maintenance-leadership.md
[review]: reference/advanced/16-key-concepts-review.md
[portability]: reference/advanced/17-language-portability.md
[http-lesson]: CONCEPTS.md#1-http-api-contracts
[data-lesson]: CONCEPTS.md#2-data-transactions-and-architecture
[production-lesson]:
  CONCEPTS.md#3-production-and-distributed-systems
