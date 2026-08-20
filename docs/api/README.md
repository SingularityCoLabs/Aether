# API contracts

The source of truth is api/proto/aether/v1. Buf checks formatting and linting
and generates:

- Go messages and Connect handlers under gen/go;
- TypeScript messages and Connect descriptors under gen/ts.

The Phase 0 public RPC surface contains only SystemService. Domain services are
added with their capability phases rather than published as non-functional
stubs.
