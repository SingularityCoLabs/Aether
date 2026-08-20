# Security policy

## Reporting a vulnerability

Do not open a public issue or discussion for a suspected vulnerability. Use
[GitHub private vulnerability reporting](https://github.com/SingularityCoLabs/Aether/security/advisories/new)
to send the maintainers affected versions or commits, reproduction details,
impact, and any proposed mitigation. Remove credentials, private keys,
production data, and unrelated personal information from the report.

Maintainers will acknowledge a valid report, investigate it privately, and
coordinate disclosure after a fix is available. No response-time guarantee is
offered before Aether's first stable release.

## Supported versions

| Version            | Supported |
| ------------------ | --------- |
| Latest `main`      | Yes       |
| Older commits/tags | No        |

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
