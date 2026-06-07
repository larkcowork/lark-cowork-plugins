# Connectors (Lark)

## How tool references work

Plugin files use `~~category` as a placeholder. In **Lark Cowork**, generic collaboration
categories resolve to the single **`lark`** MCP server (`lark-cli mcp serve`), which exposes
curated `lark_*` tools plus a generic `lark_api` escape hatch. Specialty categories keep
their own external MCP server — connect those separately.

## Connectors for this plugin

| Category | Placeholder | Backed by | Tools / notes |
|----------|-------------|-----------|---------------|
| Chat | `~~chat` | Lark IM (`lark`) | `lark_im_send`, `lark_im_search`, `lark_im_card_send` |
| Source control | `~~source control` | GitHub (external) | connect this tool's own MCP server; or approximate with Lark Base/Sheets |
| Project tracker | `~~project tracker` | Lark Task + Base (`lark`) | `lark_task_create`, `lark_task_my`, `lark_base_search` |
| Knowledge base | `~~knowledge base` | Lark Wiki + Docs (`lark`) | `lark_doc_search`, `lark_doc_fetch`, `lark_doc_create` |
| Monitoring | `~~monitoring` | Datadog (external) | connect this tool's own MCP server; or approximate with Lark Base/Sheets |
| Incident management | `~~incident management` | PagerDuty (external) | connect this tool's own MCP server; or approximate with Lark Base/Sheets |
| CI/CD | `~~CI/CD` | external MCP (external) | connect this tool's own MCP server; or approximate with Lark Base/Sheets |

## Notes for Claude

- **Resolve people first** — use `lark_contact_search` to turn a name/email into an `open_id` before any send/invite/assign.
- **Escape hatch** — anything without a dedicated tool goes through `lark_api` (Lark OpenAPI path + params).
- **Identity** — tools run as the authenticated user by default; retry with bot identity if an error envelope says so.

## Lark depth — read these for any non-trivial workflow

This plugin runs on Lark. Before executing, consult the shared depth core (one level up):

- [`../connectors/LARK-PATTERNS.md`](../connectors/LARK-PATTERNS.md) — Lark-native workflow patterns (interactive cards, Base as system-of-record, Minutes→tasks, approval flows, safe-mutation, token economy).
- [`../connectors/LARK-RECIPES.md`](../connectors/LARK-RECIPES.md) — exact `lark_api` calls for ops the curated tools don't cover.
- [`../connectors/LARK-FUSION.md`](../connectors/LARK-FUSION.md) — when to delegate to a deeper installed `lark-*` / `base-deploy` skill instead of hand-rolling.

Default posture: resolve people first, project reads with `jq`, preview mutations with `dry_run`, and surface decisions/digests as interactive cards rather than plain text.

