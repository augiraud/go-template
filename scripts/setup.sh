#!/usr/bin/env bash
set -euo pipefail
cd "$(dirname "$0")/.."

for cmd in git go docker; do
  command -v "$cmd" >/dev/null || { echo "Error: $cmd is required. Install it, then run this script again." >&2; exit 1; }
done

if [[ -f scripts/init_project.go && -t 0 ]]; then
  go run scripts/init_project.go
fi

go mod download
go -C tools mod download

if [[ -z "${CI:-}" ]]; then
  git config core.hooksPath .githooks
fi

echo "Setup complete. Try ./scripts/run.sh, and check everything with ./scripts/local_ci.sh"
