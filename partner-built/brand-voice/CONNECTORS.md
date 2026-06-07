# Connectors (Lark)

## How tool references work

The everyday collaboration layer (chat, mail, calendar, docs, tasks, Base) runs on the single
**`lark`** MCP server. This plugin's specialty tool stays on its own external MCP server.

## Connectors for this plugin

| Category | Backed by | Tools / notes |
|----------|-----------|---------------|
| Chat | Lark IM | `lark_im_send`, `lark_im_card_send` |
| Knowledge base | Lark Wiki + Docs | `lark_doc_search`/`fetch`, `lark_wiki_node_create` |
| Meeting transcripts | Lark Minutes + VC | `lark_minutes_search`, `lark_vc_search` |
| Design / calls | Figma · Gong · Granola (external) | kept as-is |

## Lark depth — read these for any non-trivial workflow

This plugin runs on Lark. Before executing, consult the shared depth core (one level up):

- [`../../connectors/LARK-PATTERNS.md`](../../connectors/LARK-PATTERNS.md) — Lark-native workflow patterns.
- [`../../connectors/LARK-RECIPES.md`](../../connectors/LARK-RECIPES.md) — exact `lark_api` calls + the 25 curated tools.
- [`../../connectors/LARK-FUSION.md`](../../connectors/LARK-FUSION.md) — when to delegate to a deeper `lark-*` / `base-deploy` skill.

Default posture: resolve people first (`lark_contact_search`), project reads with `jq` under `.data`, preview mutations with `dry_run`, surface decisions/digests as interactive cards.
