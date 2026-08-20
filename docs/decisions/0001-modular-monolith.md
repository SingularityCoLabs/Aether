# ADR 0001: Start with a modular monolith

- Status: Accepted
- Date: 2026-08-20

## Context

Aether needs clear ownership for identity, resources, operations, policy,
audit, and providers, but those capabilities do not yet need independent
deployment, scaling, or failure domains.

## Decision

Build aetherd as one Go process with enforced module boundaries. Keep
aether-node separate because it runs on managed machines. Keep the dashboard
separate because it has a distinct web runtime.

## Consequences

- Cross-module changes are transactional and observable in one process.
- Local development and operations remain small.
- Modules expose narrow interfaces so a later split is possible.
- A module may not create its own transport or database without a demonstrated
  need and an ADR.
