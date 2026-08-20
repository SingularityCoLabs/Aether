# Aether roadmap

Progress is dependency ordered. A phase is complete only when its behavior is
implemented, tested, observable, and documented.

## Phase 0 — repository foundation

- [x] Go module and three binary entrypoints
- [x] Next.js dashboard foundation
- [x] PostgreSQL connection and migration path
- [x] Docker Compose development stack
- [x] Protobuf, Buf, and ConnectRPC contract generation
- [x] structured logging and environment configuration
- [x] CI lanes for Go, web, API contracts, integration, and E2E

## Phase 1 — identity and projects

- [ ] users, organizations, memberships, and projects
- [ ] authentication
- [ ] RBAC foundations
- [ ] tenant-scoped persistence and API tests

## Phase 2 — node management

- [ ] single-use enrollment tokens
- [ ] outbound mutually authenticated node channel
- [ ] heartbeat and online/offline detection
- [ ] hardware and runtime inventory
- [ ] Infrastructure / Nodes dashboard

## Phase 3 — resources

- [ ] persistent Resource, ResourceRevision, and ResourceRelation models
- [ ] optimistic concurrency and generation handling
- [ ] reconciliation queue and status conditions

## Phase 4 — operations

- [ ] durable operations, steps, and events
- [ ] retries, cancellation, verification, and rollback states
- [ ] operation timeline UI

## Phase 5 — Docker provider

- [ ] typed image, container, volume, network, logs, and inspect capabilities
- [ ] provider contract suite
- [ ] no arbitrary shell execution surface

## Phase 6 — applications

- [ ] Application resource backed by Docker
- [ ] ports, environment, limits, volumes, and health checks
- [ ] deployment history, restart, logs, and deletion

## Phases 7–9 — services

- [ ] routes, domains, TLS, and reverse proxy
- [ ] PostgreSQL service resource
- [ ] backup policies, backup, and restore

## Phases 10–12 — intelligence

- [ ] read-only, resource-grounded assistant
- [ ] structured plan generation
- [ ] policy validation, approval, controlled execution, and audit

## Later substrates

- [ ] K3s and Kubernetes providers
- [ ] Crossplane and OpenTofu cloud providers
- [ ] NATS event transport when one process is no longer sufficient
- [ ] Temporal for genuinely long-running workflows
- [ ] OpenBao for secret values and dynamic credentials

## v0.1 release gate — Node

The first release is complete when one Linux server can join Aether and a user
can deploy, observe, view logs for, restart, and delete one Docker application
with audit history and passing E2E tests.
