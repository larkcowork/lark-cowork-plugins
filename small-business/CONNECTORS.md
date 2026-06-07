# Connectors (Lark)

## How tool references work

In **Lark Cowork**, the everyday collaboration layer (chat, email, calendar, documents, files)
runs on the single **`lark`** MCP server (`lark-cli mcp serve`), which exposes curated `lark_*`
tools plus a generic `lark_api` escape hatch. Money/commerce tools keep their own external MCP
server — connect those separately.

## Connectors for this plugin

| Category | Backed by | Tools / notes |
|----------|-----------|---------------|
| Chat | Lark IM (`lark`) | `lark_im_send`, `lark_im_search`, `lark_im_card_send` |
| Email | Lark Mail (`lark`) | `lark_mail_send`, `lark_mail_draft_create` |
| Calendar | Lark Calendar (`lark`) | `lark_calendar_agenda`, `lark_calendar_create` |
| Documents / files | Lark Docs/Sheets/Drive (`lark`) | `lark_doc_*`, `lark_sheets_read/append`, `lark_drive_upload` |
| Directory | Lark Contact (`lark`) | `lark_contact_search` |
| Accounting | QuickBooks (external) | connect its own MCP server |
| Payments | Stripe, Square, PayPal (external) | connect their own MCP servers |
| CRM | HubSpot (external) | connect its own MCP server; or build a CRM in Lark Base |
| Design | Canva (external) | connect its own MCP server |
| E-signature | DocuSign (external) | connect its own MCP server |

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

