GO ?= $(if $(wildcard .tools/go/bin/go),.tools/go/bin/go,go)
GOFMT ?= $(if $(wildcard .tools/go/bin/gofmt),.tools/go/bin/gofmt,gofmt)
PNPM ?= pnpm
SQLC ?= $(if $(wildcard .tools/bin/sqlc),.tools/bin/sqlc,sqlc)
VERSION ?= dev
COMMIT ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo none)
BUILD_DATE ?= $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
LDFLAGS := -s -w \
	-X github.com/SingularityCoLabs/aether/internal/buildinfo.Version=$(VERSION) \
	-X github.com/SingularityCoLabs/aether/internal/buildinfo.Commit=$(COMMIT) \
	-X github.com/SingularityCoLabs/aether/internal/buildinfo.Date=$(BUILD_DATE)

.PHONY: bootstrap tools generate generate-api generate-sql format format-check lint typecheck \
	test test-go test-web test-integration test-e2e build build-go build-web check clean dev

bootstrap:
	$(PNPM) --version
	$(PNPM) install --frozen-lockfile
	$(MAKE) tools

tools:
	./scripts/install-tools.sh

generate: generate-api generate-sql

generate-api:
	PATH="$(CURDIR)/.tools/go/bin:$$PATH" $(PNPM) buf:generate

generate-sql:
	$(SQLC) generate

format:
	$(GO) fmt ./...
	$(PNPM) buf:format
	$(PNPM) format

format-check:
	test -z "$$($(GOFMT) -l $$(find . -type f -name '*.go' -not -path './.tools/*'))"
	$(PNPM) buf:format:check
	$(PNPM) format:check

lint:
	$(GO) vet ./...
	$(PNPM) lint

typecheck:
	$(PNPM) typecheck

test: test-go test-web

test-go:
	$(GO) test -race ./...

test-web:
	$(PNPM) test

test-integration:
	test -n "$$AETHER_TEST_DATABASE_URL"
	$(GO) test -race -count=1 ./tests/integration/...

test-e2e:
	$(PNPM) --filter @aether/web test:e2e

build: build-go build-web

build-go:
	mkdir -p bin
	CGO_ENABLED=0 $(GO) build -trimpath -ldflags "$(LDFLAGS)" -o bin/aetherd ./cmd/aetherd
	CGO_ENABLED=0 $(GO) build -trimpath -ldflags "$(LDFLAGS)" -o bin/aether-node ./cmd/aether-node
	CGO_ENABLED=0 $(GO) build -trimpath -ldflags "$(LDFLAGS)" -o bin/aetherctl ./cmd/aetherctl

build-web:
	$(PNPM) build

check: format-check lint typecheck test build

dev:
	./scripts/dev.sh

clean:
	$(GO) clean
	rm -rf bin web/.next coverage.out playwright-report test-results
