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
| Knowledge base | `~~knowledge base` | Lark Wiki + Docs | `lark_doc_search`, `lark_doc_fetch` |
| Project tracker | `~~project tracker` | Lark Task + Base | `lark_task_my`, `lark_task_create`, `lark_base_search` |
| Meeting intelligence | `~~conversation intelligence` | Lark Minutes + VC | `lark_minutes_search`, `lark_vc_search` |
| Directory | `~~directory` | Lark Contact | `lark_contact_search` |

## Notes for Claude

- **Resolve people first.** Any brief, digest, or prep keyed to a person needs an `open_id`.
  Use `lark_contact_search` to turn a name/email into an `open_id` before calling
  `lark_im_*`, `lark_mail_*`, `lark_calendar_*`, `lark_minutes_*`, or `lark_task_*`.
- **Read-only by default.** These are mostly brief/digest skills — **no auto-send, no
  auto-RSVP, no auto-assign**. Produce drafts and suggestions only; mutate (send mail, create
  task, block calendar, RSVP) **only after the user confirms**.
- **Orchestrators fan out in parallel.** `morning-brief` and `overwhelm-triage` dispatch their
  sub-skills in parallel in **one message**, then merge — never one round trip per sub-skill.
- **Surface as cards.** Present briefs, digests, and decisions as interactive cards
  (`lark_im_card_send`) rather than plain text so the user can act inline.
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
