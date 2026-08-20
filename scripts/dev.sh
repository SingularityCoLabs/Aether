#!/usr/bin/env bash
set -euo pipefail

if ! command -v docker >/dev/null 2>&1; then
  echo "docker is required for the local development stack" >&2
  exit 1
fi

docker compose up --build

