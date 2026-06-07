# Connectors (Lark)

## How tool references work

Plugin files use `~~category` as a placeholder. In **Lark Cowork**, generic collaboration
categories resolve to the single **`lark`** MCP server (`lark-cli mcp serve`), which exposes
curated `lark_*` tools plus a generic `lark_api` escape hatch. Specialty categories keep
their own external MCP server — connect those separately.

## Connectors for this plugin

**None at runtime.** `lark-cli-dev` is a developer/meta plugin: it builds, debugs, and extends
the `lark-cli` MCP bridge (`cmd/mcp/`) — it operates on **lark-cli source**, not on a live Lark
workspace. It ships **no `.mcp.json`** and consumes no `lark_*` tools itself.

| Category | Placeholder | Backed by | Tools / notes |
|----------|-------------|-----------|---------------|
| (none) | — | — | This plugin produces/maintains the `lark` server other plugins use; it does not call it. |

The tools it works *on* are the curated `lark_*` surface defined in
[`../connectors/CONNECTORS.lark.md`](../connectors/CONNECTORS.lark.md) — consult that master
mapping when adding or refining a tool.

## Notes for Claude

- **Source, not runtime** — commands here (`/mcp-add`, `/mcp-rebuild`, `/mcp-doctor`, …) act on the
  lark-cli Go source and binary, not on Lark data. Point `LARK_REPO` at the lark-cli checkout.
- **Verify before claiming done** — rebuild + run the MCP handshake smoke-test; `code: 0` is not success.

## Lark depth — read these for any non-trivial workflow

This plugin maintains the bridge the rest of Lark Cowork runs on. Before adding/refining tools,
consult the shared depth core (one level up):

- [`../connectors/LARK-PATTERNS.md`](../connectors/LARK-PATTERNS.md) — Lark-native workflow patterns (interactive cards, Base as system-of-record, Minutes→tasks, approval flows, safe-mutation, token economy).
- [`../connectors/LARK-RECIPES.md`](../connectors/LARK-RECIPES.md) — exact `lark_api` calls for ops the curated tools don't cover.
- [`../connectors/LARK-FUSION.md`](../connectors/LARK-FUSION.md) — when to delegate to a deeper installed `lark-*` / `base-deploy` skill instead of hand-rolling.

Default posture: a new tool should follow the same conventions — resolve people first, support `jq` projection, preview mutations with `dry_run`, and return structured envelopes.
