# Connectors (Lark)

## How tool references work

Plugin files use `~~category` as a placeholder for whatever tool the user connects in that
category. In **Lark Cowork**, every category resolves to the **`lark`** MCP server (the
`lark-cli mcp serve` bridge), which exposes a curated set of `lark_*` tools plus a generic
`lark_api` escape hatch for any Lark OpenAPI endpoint not covered by a dedicated tool.

There is **one** MCP server (`lark`) backing every category — Claude picks the right `lark_*`
tool for the task based on the mapping below. If a needed operation has no dedicated tool, use
`lark_api` (pass the OpenAPI path + params).

## Master category → Lark tool mapping

| Category | Placeholder | Lark product | Primary `lark_*` tools | Fallback |
|----------|-------------|--------------|------------------------|----------|
| Chat / messaging | `~~chat` | Lark IM | `lark_im_send`, `lark_im_search`, `lark_im_card_send` | `lark_api` (im/v1) |
| Email | `~~email` | Lark Mail | `lark_mail_send`, `lark_mail_draft_create` | `lark_api` (mail/v1) |
| Calendar | `~~calendar` | Lark Calendar | `lark_calendar_agenda`, `lark_calendar_create` | `lark_api` (calendar/v4) |
| Knowledge base / wiki | `~~knowledge base` | Lark Wiki + Docs | `lark_doc_search`, `lark_doc_fetch`, `lark_doc_create` | `lark_api` (wiki/v2) |
| Documents | `~~documents` | Lark Docs | `lark_doc_create`, `lark_doc_fetch` | `lark_api` (docx/v1) |
| Office suite | `~~office suite` | Lark Docs/Sheets | `lark_doc_*`, `lark_sheets_read`, `lark_sheets_append` | `lark_api` |
| Spreadsheets | `~~spreadsheet` | Lark Sheets | `lark_sheets_read`, `lark_sheets_append` | `lark_api` (sheets/v3) |
| Project tracker / tasks | `~~project tracker` | Lark Task + Base | `lark_task_create`, `lark_task_my`, `lark_base_search` | `lark_api` (task/v2, bitable/v1) |
| Database / records | `~~database` | Lark Base (Bitable) | `lark_base_search` | `lark_api` (bitable/v1) |
| Cloud storage | `~~cloud storage` | Lark Drive | `lark_drive_upload` | `lark_api` (drive/v1) |
| Meeting transcription / conversation intelligence | `~~conversation intelligence` | Lark Minutes + VC | `lark_minutes_search`, `lark_vc_search` | `lark_api` (minutes/v1, vc/v1) |
| CRM | `~~CRM` | Lark Base (CRM table) | `lark_base_search` (search the CRM Base) | `lark_api` (bitable/v1) |
| OKR / goals | `~~okr` | Lark OKR | `lark_okr_cycle_list` | `lark_api` (okr/v1) |
| Directory / people lookup | `~~directory` | Lark Contact | `lark_contact_search` | `lark_api` (contact/v3) |

## Notes for Claude

- **Resolve people first.** Most write actions (send a message, invite to a meeting, assign a
  task) need an `open_id`. Use `lark_contact_search` to turn a name/email into an `open_id`
  before calling `lark_im_*`, `lark_calendar_*`, or `lark_task_*`.
- **CRM lives in a Base.** There is no separate CRM product — a CRM is a Lark Base table.
  `lark_base_search` reads it; writes/updates go through `lark_api` (bitable records) or the
  `lark-base` skill.
- **Escape hatch.** Anything not covered above → `lark_api` with the Lark OpenAPI path. The
  bridge handles auth (user/bot identity) automatically.
- **Identity.** Tools run as the authenticated user by default; some operations require bot
  identity. If a call returns `validation/...: only supports: user|bot`, retry with the other
  identity (the bridge documents this in its error envelope).
