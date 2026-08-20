# Security policy

## Reporting a vulnerability

Do not open a public issue for a suspected vulnerability. Report it privately
through the repository host's security-advisory channel. Include affected
versions, reproduction details, impact, and any proposed mitigation.

## Supported versions

Aether has not published a stable release. Until then, only the latest commit
on the default branch receives security fixes.

## Security invariants

- The node initiates its connection to the control plane.
- Node actions are typed; arbitrary shell execution is not an agent tool.
- Plans are validated and policy-checked before execution.
- High-impact actions require explicit approval.
- Secret values never belong in normal resource rows, logs, plans, or audit
  payloads.
- Every infrastructure mutation records actor, intent, decision, and outcome.
- Tenant and project scope must be enforced in the storage layer as well as
  the API layer before multi-tenant use.

See [docs/security/threat-model.md](docs/security/threat-model.md).
