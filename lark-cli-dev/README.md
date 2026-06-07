# lark-cli Dev Tools Plugin

A focused toolkit for building, debugging, and extending the built-in MCP bridge that ships inside `lark-cli` (`cmd/mcp/`). This plugin is for developers working on the bridge itself — adding tools, fixing schemas, and keeping MCP hosts connected — not for end users running Lark workflows.

## Installation

```
claude plugins add lark-cowork/lark-cli-dev
```

## What It Does

This plugin gives Claude the mental model and the muscle memory for working on `cmd/mcp/`:

- **The `lark-cli-mcp` skill** — Teaches Claude how the bridge actually works: a stateless stdio JSON-RPC server (`lark-cli mcp serve`) that turns each `tools/call` into an `os/exec` of a `lark-cli` shortcut, which in turn hits the Lark/Feishu Open API. It encodes the hard rules — stdout is JSON-RPC only (log to stderr via `s.logf`), no third-party MCP library, one MCP tool ≈ one shortcut (don't compose flows), and always verify flag names against `shortcuts/<domain>/*.go` before shipping.
- **Six commands** — A full dev loop: scaffold a new tool, rebuild and reinstall the binary, smoke-test the handshake, invoke a single tool end-to-end, inspect the live catalogue, and run a one-page health report across every layer of the bridge.

## Skills

| Skill | Description |
|-------|-------------|
| `lark-cli-mcp` | Build/debug/extend the lark-cli MCP bridge in cmd/mcp/. Triggers MCP tool add/remove/refine, Claude Desktop disconnects, tool schema changes. |

## Commands

| Command | What it does |
|---------|--------------|
| `/mcp-add` | Walk through adding a new tool to the lark-cli MCP bridge |
| `/mcp-call` | Invoke one MCP tool end-to-end (handshake + tools/call) and print the result |
| `/mcp-doctor` | One-page health report for the MCP bridge — build, handshake, sample call, host config, logs |
| `/mcp-rebuild` | Rebuild lark-cli and reinstall the binary so Claude Desktop picks up changes |
| `/mcp-test` | Smoke-test the lark-cli MCP server (initialize + tools/list + optional tools/call) |
| `/mcp-tools` | List the MCP tools currently exposed by `lark-cli mcp serve` |

## Requirements

This plugin assumes a local clone of the `lark-cli` Go source — specifically `cmd/mcp/` (the bridge) and `shortcuts/` (the underlying commands) — plus a working Go toolchain. The commands run `go build` and drive `lark-cli mcp serve` against that source, so they expect to operate from the repository root with `go` on your `PATH`. It does **not** bundle the runtime `lark` MCP server; it is dev tooling for the bridge code.
