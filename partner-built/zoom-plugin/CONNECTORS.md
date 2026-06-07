# Connectors (Lark)

## How tool references work

The everyday collaboration layer (chat, mail, calendar, docs, tasks, Base) runs on the single
**`lark`** MCP server. This plugin's specialty tool stays on its own external MCP server.

## Connectors for this plugin

| Category | Backed by | Tools / notes |
|----------|-----------|---------------|
| Chat (collab handoff) | Lark IM | `lark_im_card_send` |
| Calendar | Lark Calendar | `lark_calendar_agenda`, `lark_calendar_freebusy` |
| Meeting notes → tasks | Lark Minutes + Task | `lark_minutes_search` → `lark_task_create` |
| Durable docs | Lark Wiki | `lark_wiki_node_create` |
| Zoom meetings / SDK / API | Zoom (external) | Zoom MCP + REST/SDK — kept as-is |

## Lark depth — read these for any non-trivial workflow

This plugin runs on Lark. Before executing, consult the shared depth core (one level up):

- [`../../connectors/LARK-PATTERNS.md`](../../connectors/LARK-PATTERNS.md) — Lark-native workflow patterns.
- [`../../connectors/LARK-RECIPES.md`](../../connectors/LARK-RECIPES.md) — exact `lark_api` calls + the 25 curated tools.
- [`../../connectors/LARK-FUSION.md`](../../connectors/LARK-FUSION.md) — when to delegate to a deeper `lark-*` / `base-deploy` skill.

Default posture: resolve people first (`lark_contact_search`), project reads with `jq` under `.data`, preview mutations with `dry_run`, surface decisions/digests as interactive cards.
