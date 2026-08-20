# Provider development

Providers plan, apply, observe, and delete typed resources. They do not expose
arbitrary SDK or shell access to API clients.

Each provider must eventually pass the shared contract suite under
tests/contract and document:

- supported resource kinds and capabilities;
- idempotency and retry behavior;
- observation freshness;
- deletion and orphan semantics;
- secrets and permissions required;
- failure mapping and rollback guarantees.

The Docker provider begins in Phase 5. Kubernetes, Crossplane, and OpenTofu
remain intentionally absent from the executable dependency graph.
