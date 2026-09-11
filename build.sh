#!/usr/bin/env bash
# Build, test, and install the customized agent-session-cleaner.
# Usage: ./build.sh            # build + test
#        ./build.sh install    # also install the binary to ~/.local/bin
set -euo pipefail

export PATH="/opt/homebrew/bin:$PATH"
cd "$(dirname "$0")"

echo "==> go vet"
go vet ./...

echo "==> go build ./..."
go build ./...

echo "==> go test (key packages)"
go test ./internal/i18n/... ./internal/tui/... ./cmd/...

if [[ "${1:-}" == "install" ]]; then
  echo "==> install to ~/.local/bin/asc"
  mkdir -p "$HOME/.local/bin"
  go build -o "$HOME/.local/bin/asc" ./cmd/agent-session-cleaner
  echo "Installed: $HOME/.local/bin/asc"
  echo "Make sure ~/.local/bin is on your PATH (it is where claude lives on this machine)."
fi

echo "==> done"
