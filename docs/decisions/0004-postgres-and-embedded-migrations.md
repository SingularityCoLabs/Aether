# ADR 0004: PostgreSQL with embedded migrations

- Status: Accepted
- Date: 2026-08-20

## Context

Control-plane state needs transactions, durable relations, JSON resource
payloads, and a deployment path that does not depend on separate migration
artifacts.

## Decision

Use PostgreSQL through pgx. Author Goose migrations as SQL and embed them into
aetherd. Use sqlc for domain queries once persistent Phase 1 models arrive.

## Consequences

- The binary can migrate its compatible schema during startup.
- Migration failure prevents the server from claiming readiness.
- Roll-forward migrations are preferred; destructive changes require explicit
  backup and recovery instructions.
