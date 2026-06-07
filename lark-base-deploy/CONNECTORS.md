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
| Project tracker | `~~project tracker` | Lark Base + Task | `lark_base_search`, `lark_task_create` |
| Office suite | `~~office suite` | Lark Sheets/Docs/Drive | `lark_sheets_read`, `lark_sheets_append`, `lark_drive_upload`, `lark_doc_create` |
| Chat | `~~chat` | Lark IM | `lark_im_send`, `lark_im_card_send` (bot reminders/notifications) |
| Directory | `~~directory` | Lark Contact | `lark_contact_search` |

## Notes for Claude

- **Base writes go through `lark_api` or the `lark-base` skill.** Most Base operations have **no
  dedicated curated tool** — creating a base/table/field/view, writing bitable records, building
  dashboard blocks, wiring workflow automation. For these, either call `lark_api` (the
  `bitable/v1` and `base/v2` OpenAPI families) **or** delegate to the installed `lark-base` skill,
  which owns the verb mapping and field/dashboard/workflow JSON.
- **Resolve people first.** Use `lark_contact_search` to turn a name/email into an `open_id`
  **before** assigning user fields, setting record owners, or granting roles. Never write a raw
  name into a person field.
- **`build-plan.json` is the single source of truth.** Every phase reads and writes it. The schema
  + invariants define what "done" means; do not act on Base structure that isn't in the plan.
- **Never invent identifiers.** `base_token`, `table_id`, `field_id`, `view_id`, `dashboard_id`
  must be **real values returned by the API** and written back into the plan. Placeholder or
  guessed IDs corrupt the pipeline.

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
