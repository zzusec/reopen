#!/usr/bin/env bash
# Build, test, and install restore-session (forked from haowang02/agent-session-cleaner).
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
  echo "==> install to ~/.local/bin/restore-session (+ re alias)"
  mkdir -p "$HOME/.local/bin"
  go build -o "$HOME/.local/bin/restore-session" ./cmd/restore-session
  ln -sf restore-session "$HOME/.local/bin/re"
  echo "Installed: $HOME/.local/bin/restore-session"
  echo "Alias:     $HOME/.local/bin/re -> restore-session"
  echo "Make sure ~/.local/bin is on your PATH."
fi

echo "==> done"
