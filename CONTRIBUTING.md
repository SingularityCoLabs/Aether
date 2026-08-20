# Contributing to Aether

Thank you for helping build Aether. Contributions of code, tests,
documentation, issue triage, and design feedback are welcome.

By participating, you agree to follow our
[Code of Conduct](.github/CODE_OF_CONDUCT.md). Contributions are provided under
the repository's [MIT License](LICENSE).

## Before starting

1. Search existing issues and discussions before opening a duplicate.
2. Use [GitHub Discussions](https://github.com/SingularityCoLabs/Aether/discussions)
   for questions and early design exploration.
3. Use an issue form for reproducible bugs or concrete capability proposals.
4. For a substantial change, agree on the Resource/API boundary in an issue or
   discussion before investing in implementation.
5. Report vulnerabilities privately as described in [SECURITY.md](SECURITY.md).

The [roadmap](ROADMAP.md) is dependency ordered. An unchecked roadmap item is
not automatically ready for implementation; confirm its API boundary and
prerequisites first.

## Development environment

Required versions are recorded in `.tool-versions`:

- Go 1.27;
- Node.js 24.19 LTS;
- pnpm 10.32;
- Buf 1.72 and sqlc 1.31.

Docker with Compose is required for the complete local stack and database
integration tests.

```bash
git clone https://github.com/SingularityCoLabs/Aether.git
cd Aether
cp .env.example .env
make bootstrap
make generate
make check
```

Run the full stack with:

```bash
docker compose up --build
```

The dashboard is available at `http://localhost:3000` and the control plane at
`http://localhost:8080`.

## Before opening a change

1. Read [ARCHITECTURE.md](ARCHITECTURE.md) and the accepted ADRs.
2. Keep changes inside one capability and its tests.
3. Preserve the control-plane boundary: UI, CLI, and AI use the same API.
4. Do not add a general remote-shell or agent-to-provider shortcut.
5. Keep a pull request focused on one capability or defect.
6. Add tests at the narrowest layer that proves the behavior.

## Branch and pull request workflow

1. Fork the repository and create a descriptive branch such as
   `feat/node-heartbeat`, `fix/readiness-timeout`, or `docs/provider-contracts`.
2. Make small, reviewable commits with meaningful messages.
3. Regenerate committed API or SQL output when its source changes.
4. Run the relevant focused tests and `make check` before opening a pull
   request.
5. Complete the pull request template and link the related issue or discussion.

Pull requests are squash-merged after required checks and review. Maintainers
may ask for an ADR when a change alters an architectural invariant, public API,
storage contract, trust boundary, or provider behavior.

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

`make generate` must leave no diff unless the source Protobuf or SQL contract
changed. Never hand-edit files under `gen/` or `internal/database/dbsql/`.

Database migrations are append-only after merge. Do not rewrite a migration
that may have run in another environment.

## Commits and pull requests

- Explain the user-visible capability and failure behavior.
- Add tests at the narrowest useful layer.
- Call out schema, API, security, or operational changes.
- Update the relevant ADR when an architectural decision changes.
- Never commit credentials, enrollment tokens, private keys, or production
  data.
- Keep generated output in the same commit as its source contract.
- Treat review comments as collaborative requests; explain tradeoffs when a
  suggestion cannot be adopted.

## Review criteria

Maintainers review contributions for correctness, tests, architectural fit,
security, operational failure behavior, documentation, and backward
compatibility. A pull request may be declined when it bypasses the Resource
API, exposes arbitrary execution, weakens tenant or secret boundaries, or
introduces a future subsystem without a current capability requiring it.

The project's governance and decision process are documented in
[GOVERNANCE.md](GOVERNANCE.md).
