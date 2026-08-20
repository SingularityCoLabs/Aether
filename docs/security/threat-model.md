# Initial threat model

This is a living security boundary document, not a claim of production
readiness.

## Assets

- control-plane identities and authorization state;
- node identities and enrollment material;
- desired infrastructure state;
- approval, operation, and audit history;
- provider credentials and secret references.

## Trust boundaries

1. Browser or CLI to the public control-plane API.
2. aether-node to the dedicated node channel.
3. control plane to PostgreSQL.
4. provider adapter to Docker or a future substrate.
5. untrusted AI model output to the Plan validator.
6. untrusted application logs and resource metadata to AI context.

## Initial controls

- strict typed messages at every remote boundary;
- authentication before authorization and project scoping;
- short-lived, single-use node enrollment followed by rotated node identity;
- idempotency keys and replay detection for mutations;
- explicit operation state transitions;
- risk classification and non-bypassable approval checks;
- structured audit with redaction;
- no arbitrary shell or raw provider credential in agent tools;
- bounded request sizes, deadlines, and concurrency;
- dependency, container, and secret scanning in CI.

## Phase 0 limitations

Authentication, node enrollment, mutation APIs, secret storage, and AI are not
implemented. The Phase 0 HTTP surface exposes only health and system metadata.
It must not be placed on a public network as a production service.
