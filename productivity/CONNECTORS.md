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
| Chat | `~~chat` | Lark IM | `lark_im_send`, `lark_im_search`, `lark_im_card_send` |
| Email | `~~email` | Lark Mail | `lark_mail_send`, `lark_mail_draft_create` |
| Calendar | `~~calendar` | Lark Calendar | `lark_calendar_agenda`, `lark_calendar_create` |
| Knowledge base | `~~knowledge base` | Lark Wiki + Docs | `lark_doc_search`, `lark_doc_fetch`, `lark_doc_create` |
| Project tracker | `~~project tracker` | Lark Task + Base | `lark_task_create`, `lark_task_my`, `lark_base_search` |
| Office suite | `~~office suite` | Lark Docs/Sheets/Drive | `lark_doc_*`, `lark_sheets_read`, `lark_sheets_append`, `lark_drive_upload` |
| Directory | `~~directory` | Lark Contact | `lark_contact_search` |

## Notes for Claude

- **Resolve people first.** Sending a message, inviting to a meeting, or assigning a task needs
  an `open_id`. Use `lark_contact_search` to turn a name/email into an `open_id` before calling
  `lark_im_*`, `lark_calendar_*`, or `lark_task_*`.
- **Tasks + Base.** Personal/team to-dos live in Lark Task (`lark_task_*`); structured trackers
  (sprints, pipelines, rosters) live in Lark Base — read with `lark_base_search`, write via
  `lark_api` (bitable records).
- **Escape hatch.** Anything not covered above → `lark_api` with the Lark OpenAPI path. The
  bridge handles user/bot identity automatically.

## Lark depth — read these for any non-trivial workflow

This plugin runs on Lark. Before executing, consult the shared depth core (one level up):

- [`../connectors/LARK-PATTERNS.md`](../connectors/LARK-PATTERNS.md) — Lark-native workflow patterns (interactive cards, Base as system-of-record, Minutes→tasks, approval flows, safe-mutation, token economy).
- [`../connectors/LARK-RECIPES.md`](../connectors/LARK-RECIPES.md) — exact `lark_api` calls for ops the curated tools don't cover.
- [`../connectors/LARK-FUSION.md`](../connectors/LARK-FUSION.md) — when to delegate to a deeper installed `lark-*` / `base-deploy` skill instead of hand-rolling.

Default posture: resolve people first, project reads with `jq`, preview mutations with `dry_run`, and surface decisions/digests as interactive cards rather than plain text.

