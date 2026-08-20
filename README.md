# Aether

[![CI](https://github.com/SingularityCoLabs/Aether/actions/workflows/ci.yml/badge.svg)](https://github.com/SingularityCoLabs/Aether/actions/workflows/ci.yml)
[![CodeQL](https://github.com/SingularityCoLabs/Aether/actions/workflows/codeql.yml/badge.svg)](https://github.com/SingularityCoLabs/Aether/actions/workflows/codeql.yml)
[![License: MIT](https://img.shields.io/badge/license-MIT-0f766e.svg)](LICENSE)
[![Contributions welcome](https://img.shields.io/badge/contributions-welcome-0f766e.svg)](CONTRIBUTING.md)

Aether is an AI-native, self-hosted cloud control plane built with Go and
Next.js. It exposes one typed Resource API to the dashboard, CLI, automation,
and future AI agent. Its central design rule is:

> Aether is a control plane first, an AI agent second, and an infrastructure
> executor third.

The dashboard, CLI, API, and future agent all use the same typed Aether
Resource API. Providers are implementation details behind that boundary; the
agent never bypasses it to run arbitrary infrastructure commands.

## Current status

This repository contains the Phase 0 foundation:

- a modular Go control plane (aetherd);
- the aether-node and aetherctl command surfaces;
- a Next.js dashboard foundation;
- Protobuf contracts generated for Go and TypeScript with Buf;
- ConnectRPC transport;
- PostgreSQL connectivity and migrations;
- provider, plan, policy, operation, workflow, and resource boundaries;
- local Docker Compose and production container builds;
- unit, integration, contract, and browser-test lanes in CI.

It does **not** yet enroll nodes, deploy containers, authenticate users, or
execute AI-generated plans. Those capabilities are intentionally sequenced in
[ROADMAP.md](ROADMAP.md).

## Repository map

```text
cmd/                   Go binaries: aetherd, aether-node, aetherctl
internal/              Private control-plane modules
pkg/resource/          Stable desired-state resource vocabulary
api/proto/aether/v1/   Source-of-truth API contracts
gen/go/                Generated Go contracts
gen/ts/                Generated TypeScript contracts
web/                   Next.js dashboard
migrations/            Reserved for externally managed migrations
deploy/                Container and later deployment packaging
docs/                  Architecture, decisions, development, and security
tests/                 Cross-module test suites
```

Database migrations used by the binary are embedded from
internal/database/migrations so aetherd can run as a single artifact.

## Prerequisites

- Go 1.27
- Node.js 24.19 LTS
- pnpm 10.32
- Docker with Compose (for the complete local stack)

Buf is installed as a project-local npm dependency. The bootstrap target also
downloads the pinned, checksum-verified sqlc binary into the ignored .tools
directory.

## Quick start

```bash
cp .env.example .env
make bootstrap
make generate
make test
docker compose up --build
```

Then open http://localhost:3000. The control plane listens on
http://localhost:8080, with liveness and readiness at /livez and /readyz.

Run the CLI against it:

```bash
go run ./cmd/aetherctl system info --server http://localhost:8080
```

See [docs/development/getting-started.md](docs/development/getting-started.md)
for the full development workflow.

## Architectural invariants

1. Every externally visible capability starts with a Resource and typed API.
2. Infrastructure changes become Plans and Operations, not hidden CRUD side
   effects.
3. The control plane never relies on inbound SSH to managed nodes.
4. Providers implement typed capabilities; no general shell tool is exposed.
5. Risk, approval, audit, and tests precede agent tool exposure.
6. Later systems such as K3s, Crossplane, NATS, Temporal, and OpenBao are added
   only when a concrete capability needs them.

## Documentation

- [Architecture](ARCHITECTURE.md) and [accepted decisions](docs/decisions)
- [Development setup](docs/development/getting-started.md)
- [Roadmap](ROADMAP.md)
- [Security model](docs/security/threat-model.md)

## Contributing and community

Aether welcomes focused issues, documentation improvements, tests, and code
contributions. Start with [CONTRIBUTING.md](CONTRIBUTING.md), follow the
[Code of Conduct](.github/CODE_OF_CONDUCT.md), and use
[GitHub Discussions](https://github.com/SingularityCoLabs/Aether/discussions)
for questions or design exploration.

Please report vulnerabilities through
[GitHub's private vulnerability reporting](https://github.com/SingularityCoLabs/Aether/security/advisories/new),
not through a public issue. See [SECURITY.md](SECURITY.md).

## License

[MIT](LICENSE)
