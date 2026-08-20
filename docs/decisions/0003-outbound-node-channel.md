# ADR 0003: Managed nodes connect outbound

- Status: Accepted
- Date: 2026-08-20

## Context

Inbound SSH requires exposed management ports, privileged credentials, and an
unstructured command channel that is difficult to constrain or audit.

## Decision

aether-node will establish a long-lived outbound connection to aetherd using
mutual authentication. The control plane will send versioned, typed operation
messages over that channel.

## Consequences

- Managed nodes need no inbound Aether or SSH exposure.
- Enrollment, identity rotation, replay protection, revocation, backpressure,
  and reconnect semantics are required before Phase 2 is complete.
- Arbitrary command execution is excluded from the public capability model.
