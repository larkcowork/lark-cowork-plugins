# Connectors (Lark)

## How tool references work

Plugin files use `~~category` as a placeholder for whatever tool the user connects in that
category. In **Lark Cowork**, every category resolves to the single **`lark`** MCP server
(the `lark-cli mcp serve` bridge), which exposes a curated set of `lark_*` tools plus a generic
`lark_api` escape hatch for any Lark OpenAPI endpoint not covered by a dedicated tool.

There is one server backing every category — Claude picks the right `lark_*` tool for the task.
If an operation has no dedicated tool, use `lark_api` (pass the OpenAPI path + params).

## Connectors for this plugin

| Category | Placeholder | Lark product | Primary `lark_*` tools |
|----------|-------------|--------------|------------------------|
| Knowledge base / wiki | `~~knowledge base` | Lark Wiki | `lark_doc_search`, `lark_doc_fetch` (+ wiki node ops via `lark_api`) |
| Documents | `~~documents` | Lark Docs | `lark_doc_create`, `lark_doc_fetch` |
| Spreadsheets | `~~spreadsheet` | Lark Sheets | `lark_sheets_read`, `lark_sheets_append` |
| Cloud storage | `~~cloud storage` | Lark Drive | `lark_drive_upload` |
| Directory | `~~directory` | Lark Contact | `lark_contact_search` |

## Notes for Claude

- **Docs are v2 + DocxXML.** Author/fetch docs with `--api-version v2` and DocxXML by default
  (markdown only when the user asks). Create drive docs via `lark_doc_create`.
- **Wiki node ops.** Listing, moving, and re-parenting wiki nodes go through `lark_api`
  (wiki/v2 paths) or the installed `lark-wiki` skill — there is no dedicated curated tool for them.
- **Sheet append.** Structured rows (e.g. decision-log entries) go to Lark Sheets via
  `lark_sheets_append`.
- **doc-from-template** renders Mustache templates locally, then writes. ALWAYS confirm with a
  rendered preview before writing anything.
- **doc-restructure** is a READ-ONLY audit — it flags stale/orphan/duplicate pages and proposes
  archive/merge/re-parent batches. It never moves or deletes pages without explicit confirm.
- **Validate inputs.** Dates must be ISO format; resolve people to `open_id` via
  `lark_contact_search` before they reach template output; escape any user-supplied value that
  looks like template syntax.
- **Escape hatch.** Anything not covered above → `lark_api` with the Lark OpenAPI path. The
  bridge handles user/bot identity automatically.

## Tooling & verb dialect

The skills here are **workflow / orchestration skills**: they are written in `lark-cli` verb
shorthand (e.g. `contact +search-user`, `task +get-my-tasks`, `im +messages-search`) rather than
raw MCP tool names. This is intentional (see `LARK-FUSION.md`) — each verb resolves one of two ways
at runtime:

1. **MCP tool** — to the matching `lark_*` tool in the table above (e.g. `contact +search-user`
   → `lark_contact_search`), exposed by the single `lark` server.
2. **Owning skill** — for anything richer (interactive cards, Base schema, multi-step flows),
   delegate to the installed `lark-*` skill that owns that domain.

No skill body should hard-code a raw OpenAPI path. Use a curated `lark_*` tool, a `lark_api`
recipe from [`../connectors/LARK-RECIPES.md`](../connectors/LARK-RECIPES.md), or the owning skill
mapped in [`../connectors/LARK-FUSION.md`](../connectors/LARK-FUSION.md).

## Lark depth — read these for any non-trivial workflow

This plugin runs on Lark. Before executing, consult the shared depth core (one level up):

- [`../connectors/LARK-PATTERNS.md`](../connectors/LARK-PATTERNS.md) — Lark-native workflow patterns (interactive cards, Base as system-of-record, Minutes→tasks, approval flows, safe-mutation, token economy).
- [`../connectors/LARK-RECIPES.md`](../connectors/LARK-RECIPES.md) — exact `lark_api` calls for ops the curated tools don't cover.
- [`../connectors/LARK-FUSION.md`](../connectors/LARK-FUSION.md) — when to delegate to a deeper installed `lark-*` / `base-deploy` skill instead of hand-rolling.

Default posture: resolve people first, project reads with `jq`, preview mutations with `dry_run`, and surface decisions/digests as interactive cards rather than plain text.
