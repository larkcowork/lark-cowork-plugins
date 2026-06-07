# Connectors (Lark)

## How tool references work

The everyday collaboration layer (chat, mail, calendar, docs, tasks, Base) runs on the single
**`lark`** MCP server. This plugin's specialty tool stays on its own external MCP server.

## Connectors for this plugin

| Category | Backed by | Tools / notes |
|----------|-----------|---------------|
| Email | Lark Mail | `lark_mail_draft_create` (draft-only) |
| Calendar | Lark Calendar | `lark_calendar_agenda`, `lark_calendar_freebusy` |
| Meetings | Lark Minutes + VC | `lark_minutes_search`, `lark_vc_search` |
| CRM | Lark Base | `lark_base_search`, `lark_base_record_upsert` |
| Directory | Lark Contact | `lark_contact_search` |
| Community signals | Common Room (external) | Common Room MCP — kept as-is |

## Lark depth — read these for any non-trivial workflow

This plugin runs on Lark. Before executing, consult the shared depth core (one level up):

- [`../../connectors/LARK-PATTERNS.md`](../../connectors/LARK-PATTERNS.md) — Lark-native workflow patterns.
- [`../../connectors/LARK-RECIPES.md`](../../connectors/LARK-RECIPES.md) — exact `lark_api` calls + the 25 curated tools.
- [`../../connectors/LARK-FUSION.md`](../../connectors/LARK-FUSION.md) — when to delegate to a deeper `lark-*` / `base-deploy` skill.

Default posture: resolve people first (`lark_contact_search`), project reads with `jq` under `.data`, preview mutations with `dry_run`, surface decisions/digests as interactive cards.
