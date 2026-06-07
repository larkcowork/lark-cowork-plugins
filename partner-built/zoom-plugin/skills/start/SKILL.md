---
name: start
description: Start here for any Zoom integration or app idea. Use when you need to choose the right Zoom surface, shape the architecture, or route into the correct implementation skill without reading the whole Zoom doc set first.
---

# Start

Use this as the default entry skill for the plugin.

## What This Skill Does

- Classifies the request by job-to-be-done, not by product name alone
- Routes into the right implementation skill
- Pulls in product-specific Zoom references only after the route is clear
- Prevents common early mistakes, especially Meeting SDK vs Video SDK and REST API vs MCP confusion

## Routing Table

| If the user wants to... | Route to |
|---|---|
| Choose the right Zoom surface for a new project | [plan-zoom-product](../plan-zoom-product/SKILL.md) |
| Set up OAuth, tokens, scopes, or app credentials | [setup-zoom-oauth](../setup-zoom-oauth/SKILL.md) |
| Embed or customize a Zoom meeting flow | [build-zoom-meeting-app](../build-zoom-meeting-app/SKILL.md) |
| Build a bot, recorder, or real-time meeting processor | [build-zoom-bot](../build-zoom-bot/SKILL.md) |
| Use Zoom-hosted MCP for AI workflows | [setup-zoom-mcp](../setup-zoom-mcp/SKILL.md) |
| Debug a broken integration | [debug-zoom](../debug-zoom/SKILL.md) |

## Supporting Zoom References

Use these only after selecting the workflow:

- [general](../general/SKILL.md)
- [rest-api](../rest-api/SKILL.md)
- [meeting-sdk](../meeting-sdk/SKILL.md)
- [video-sdk](../video-sdk/SKILL.md)
- [webhooks](../webhooks/SKILL.md)
- [websockets](../websockets/SKILL.md)
- [oauth](../oauth/SKILL.md)
- [zoom-mcp](../zoom-mcp/SKILL.md)

## Lark-native note

Zoom stays the specialty engine (meetings, SDKs, recordings, transcripts, phone, contact center). Once you have Zoom output, the **collaboration handoff** lives in Lark: notify a channel/person with an interactive card or draft (`lark_im_card_send` / `lark_mail_draft_create`, resolve people via `lark_contact_search`), log calls/deals to a Lark Base system-of-record (`lark_base_search` / `lark_base_record_upsert`), turn meeting notes/action items into tasks (`lark_minutes_search` → `lark_task_create`), and file durable docs in Wiki (`lark_wiki_node_create`). See `../../connectors/LARK-PATTERNS.md`, `LARK-RECIPES.md`, and `LARK-FUSION.md` for the patterns and which deep `lark-*` skill owns each step.

## Operating Rules

1. Prefer one clear recommendation over a product catalog dump.
2. Ask a short clarifier only when the route is genuinely ambiguous.
3. Keep the first response architectural and actionable, then go deep.
4. Pull in deeper references only when they directly help the current decision or implementation.
