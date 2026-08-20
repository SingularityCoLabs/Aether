# Contributing to Aether

Thank you for helping build Aether.

## Before opening a change

1. Read [ARCHITECTURE.md](ARCHITECTURE.md) and the accepted ADRs.
2. Keep changes inside one capability and its tests.
3. Preserve the control-plane boundary: UI, CLI, and AI use the same API.
4. Do not add a general remote-shell or agent-to-provider shortcut.

## Development checks

```bash
make generate
make check
make test
```

For database integration tests, start PostgreSQL:

```bash
docker compose up -d postgres
make test-integration
```

For the browser suite:

```bash
pnpm --filter @aether/web exec playwright install chromium
make test-e2e
```

## Commits and pull requests

- Explain the user-visible capability and failure behavior.
- Add tests at the narrowest useful layer.
- Call out schema, API, security, or operational changes.
- Update the relevant ADR when an architectural decision changes.
- Never commit credentials, enrollment tokens, private keys, or production
  data.
