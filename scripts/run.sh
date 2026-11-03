#!/usr/bin/env bash
set -euo pipefail
bin="$(dirname "$0")/../bin/go-template"
go build -C "$(dirname "$0")/.." -o bin/go-template ./cmd/go-template
exec "$bin" "$@"
