# Development setup

## Toolchain

Versions are recorded in the root .tool-versions file. Go and Node builds are
also pinned in CI and Dockerfiles. The bootstrap target installs the pinned
sqlc release under .tools after verifying its published checksum.

```bash
make bootstrap
make generate
```

## Run the complete stack

```bash
cp .env.example .env
docker compose up --build
```

Services:

| Service       | Address               |
| ------------- | --------------------- |
| dashboard     | http://localhost:3000 |
| control plane | http://localhost:8080 |
| PostgreSQL    | localhost:5432        |

Check health:

```bash
curl --fail http://localhost:8080/livez
curl --fail http://localhost:8080/readyz
go run ./cmd/aetherctl system info --server http://localhost:8080
```

## Run processes on the host

Start only PostgreSQL:

```bash
docker compose up -d postgres
go run ./cmd/aetherd
pnpm dev
```

The defaults in .env.example match the Compose PostgreSQL instance. Environment
variables are not automatically loaded by the Go process; export them or use a
dotenv runner when starting aetherd on the host.

## Generate contracts

```bash
pnpm buf:format
pnpm buf:lint
pnpm buf:generate
```

Generated Go and TypeScript files are committed. CI regenerates them and fails
if the working tree changes.

## Test

```bash
make test
make test-integration
make test-e2e
make check
```

Integration tests require AETHER_TEST_DATABASE_URL and skip locally when it is
unset. CI supplies a real PostgreSQL service.

## Migrations

Migrations live in internal/database/migrations and are embedded in aetherd.
Use a monotonically increasing filename and Goose Up/Down sections. Never
rewrite a migration that may have run outside your machine.
