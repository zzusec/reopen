#!/usr/bin/env bash
# Build, test, and install reopen (forked from haowang02/agent-session-cleaner).
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
  echo "==> install to ~/.local/bin/reopen (+ re alias)"
  mkdir -p "$HOME/.local/bin"
  go build -o "$HOME/.local/bin/reopen" ./cmd/reopen
  ln -sf reopen "$HOME/.local/bin/re"
  echo "Installed: $HOME/.local/bin/reopen"
  echo "Alias:     $HOME/.local/bin/re -> reopen"
  echo "Make sure ~/.local/bin is on your PATH."
fi

echo "==> done"
