#!/usr/bin/env bash
# set-transport.sh — thin launcher for the Go larkify tool (kept for muscle memory).
# Switches the `lark` connector in every plugin's .mcp.json between transports.
#
#   ./set-transport.sh stdio   # Claude Desktop classic (local lark-cli binary + keychain)
#   ./set-transport.sh http    # Cowork VM / remote (lark-cli mcp serve --transport http + tunnel)
#
# The actual logic lives in tools/larkify.go — equivalent to:
#   (cd tools && go run . transport <mode>)
set -euo pipefail
ROOT="$(cd "$(dirname "$0")" && pwd)"
exec sh -c "cd '$ROOT/tools' && go run . transport \"\$1\"" _ "${1:-}"
