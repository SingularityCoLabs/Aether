# ADR 0002: The Resource API is the control boundary

- Status: Accepted
- Date: 2026-08-20

## Context

Independent UI, CLI, agent, and provider code paths would create inconsistent
policy, audit, validation, and behavior.

## Decision

All clients express intent through versioned Aether APIs. Infrastructure is
represented as desired-state Resources. Providers are reachable only through
validated Plans and Operations created by the control plane.

The future AI agent may inspect Resources and propose Plans. It does not call
Docker, Kubernetes, a cloud SDK, or a general shell directly.

## Consequences

- One API contract governs every client.
- Agent behavior inherits the same permissions and audit controls.
- Providers can change without changing the user-facing resource vocabulary.
- New features must define resource, API, policy, operation, tests, and UI
  before becoming an agent tool.
