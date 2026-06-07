---
name: slack-search
description: Guidance for effectively searching Lark IM to find messages, files, channels, and people
---

# Lark IM Search

This skill provides guidance for effectively finding information in Lark IM — messages, files, group
chats, and people. The Lark connector supersedes Slack here: route all reads through the curated
`lark_*` tools. For search grammar and pagination depth, delegate to the **`lark-im`** skill.

> Depth core: `../../connectors/LARK-PATTERNS.md` (P1 identity, P3 token economy — project reads
> with `jq`), `../../connectors/LARK-RECIPES.md`, `../../connectors/LARK-FUSION.md`.

## When to Use

Apply this skill whenever you need to find information in Lark IM — locating messages,
conversations, files, or people, or gathering context before answering a question about what's
happening in chat.

## Search Tools Overview

| Goal | Tool | Notes |
|------|------|-------|
| Find messages / conversations | `lark_im_search` | Search chat history by keyword. Project results with `jq` (P3) — pull only the fields you need. |
| Find a person (name/email/role → `open_id`) | `lark_contact_search` | P1 — resolve identity before any follow-up send/assign. |
| Find files shared in chat | `lark_im_search` | Filter to file/attachment content; see the `lark-im` skill for content-type filters. |

Anything without a curated tool goes through `lark_api` (Lark OpenAPI passthrough). For advanced
filters, pagination, and reading full thread context, defer to the **`lark-im`** skill.

## Search Strategy

### Start Broad, Then Narrow

1. Begin with a simple keyword or natural language question.
2. If too many results, add filters (group chat, sender, date range).
3. If too few results, remove filters and try synonyms or related terms.

### Choose the Right Approach

- **Natural language questions** (e.g., "What is the deadline for project X?") — best for fuzzy, conceptual searches where you don't know exact keywords.
- **Keyword search** (e.g., `project X deadline`) — best for finding specific, exact content.

### Use Multiple Searches

Don't rely on a single search. Break complex questions into smaller searches:
- Search for the topic first
- Then resolve and search for specific people's contributions (`lark_contact_search` → filter by sender)
- Then narrow to specific group chats

## Token Economy (P3)

Lark IM search can return large payloads. Always pass a `jq` projection to `lark_im_search` so you
keep only the fields you need (sender, timestamp, snippet) instead of full message bodies. Pull the
full payload only when you'll actually use it.

## Resolving People in Results

When a result references a person by `open_id`, resolve it to a name/department with
`lark_contact_search` (P1). Likewise, before any follow-up — a reply, a mention, an assignment —
resolve the target's `open_id` first; IDs cannot be guessed.

## Following Up on Results

After finding relevant messages:
- Resolve any `open_id`s to people via `lark_contact_search`.
- To reply or surface a digest, hand off to the **slack-messaging** skill (`lark_im_send` /
  `lark_im_card_send`).
- To turn a found decision/action item into durable state, log it to a Lark Base (P5,
  `lark_base_record_upsert`) or create a task (`lark_task_create`).

## Common Pitfalls

- **Search is not real-time.** Very recent messages may not appear in search results yet.
- **Don't guess IDs.** Always resolve a name/email to an `open_id` via `lark_contact_search` before
  acting on a result.
- **Don't over-fetch.** Without a `jq` projection a broad search can burn tens of thousands of
  tokens — project first (P3).
