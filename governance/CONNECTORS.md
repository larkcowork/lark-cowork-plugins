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
| Approvals | `~~approvals` | Lark Approval | `lark_api` (approval/v4) or installed `lark-approval` skill |
| Knowledge base | `~~knowledge base` | Lark Wiki + Docs | `lark_doc_search`, `lark_doc_fetch` |
| Database / records | `~~database` | Lark Base | `lark_base_search` |
| Cloud storage | `~~cloud storage` | Lark Drive | `lark_drive_upload` (+ permission reads via `lark_api`) |
| Chat | `~~chat` | Lark IM | `lark_im_search`, `lark_im_send` |
| Meeting intelligence | `~~conversation intelligence` | Lark Minutes | `lark_minutes_search` |

## Notes for Claude

- **Approval has no curated tool.** There is no dedicated `lark_approval_*` tool. Use `lark_api`
  against approval/v4 (instances initiated/get, tasks query/remind/transfer) or delegate to the
  installed `lark-approval` skill. Always run `lark-cli schema approval.<res>.<method>` before
  raw calls — don't guess field shapes.
- **permission-audit is READ-ONLY.** It scans Drive/Doc/Wiki/Base sharing settings via `lark_api`.
  No `+update`/`+apply-permission`/`+delete` — surface findings only.
- **decision-logger** reads IM/Minutes and writes detected decisions to a Base table
  (`lark_api` bitable records).
- **Default read-only.** Remind / transfer / approve only after explicit user confirm.
- **Resolve people first.** Turn names/emails into `open_id` via `lark_contact_search` before
  any per-person operation.
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
