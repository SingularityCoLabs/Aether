# Aether architecture

## System shape

```text
Web UI ─┐
CLI ────┼──> Protobuf + Connect API ──> modular control plane
AI ─────┘                                 │
                                         ├── Resources
                                         ├── Plans and approvals
                                         ├── Operations
                                         ├── Policy and audit
                                         └── Provider boundary
                                                   │
                                            outbound node channel
                                                   │
                                              Docker first
```

The first releasable path is deliberately narrower:

```text
Dashboard / CLI -> aetherd -> aether-node -> Docker
```

## Deployable units

### aetherd

The centralized modular monolith. It owns the public API, persistence,
resource model, planning, policies, operations, and audit history.

### aether-node

The daemon that will run on managed machines. It will establish an outbound,
mutually authenticated channel to aetherd, report inventory, and execute only
typed operations approved by the control plane.

### aetherctl

The human CLI. It consumes the same generated API contract as the dashboard
and future agent.

### web

The Next.js dashboard. Its code is organized by product feature so UI modules
track control-plane modules rather than becoming a generic component bucket.

## Core domain boundaries

- Resource stores desired Spec and observed Status independently.
- Provider plans, applies, observes, and deletes typed resources.
- Plan describes proposed, validated work and its aggregate risk.
- Approval records explicit authorization; it is not an execution shortcut.
- Operation is the durable state machine for infrastructure work.
- Policy determines what is permitted and what requires approval.
- Audit records actor, intent, decision, and outcome.
- WorkflowEngine allows local orchestration first and Temporal later.

## Dependency direction

```text
transport -> application modules -> domain contracts
                              \-> provider interfaces
infrastructure adapters ------^
```

Domain packages do not import transport, PostgreSQL, Docker, Kubernetes, or AI
SDKs. Generated API types stay at the transport boundary and are mapped to
domain types.

## Future systems

K3s, Crossplane, OpenTofu, NATS, Temporal, OpenBao, Prometheus, Loki, and
Grafana are extension points, not Phase 0 dependencies. Their intended seams
are documented in [docs/architecture/module-boundaries.md](docs/architecture/module-boundaries.md).

The accepted decisions are recorded under [docs/decisions](docs/decisions).
