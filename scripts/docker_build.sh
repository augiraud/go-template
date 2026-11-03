#!/usr/bin/env bash
set -euo pipefail
cd "$(dirname "$0")/.."
docker build \
  --build-arg GO_VERSION="$(go list -m -f '{{.GoVersion}}')" \
  --build-arg VERSION="$(git describe --tags --always --dirty 2>/dev/null || echo dev)" \
  --tag go-template "$@" .
