---
name: search
description: Search across all connected sources in one query. Trigger with "find that doc about...", "what did we decide on...", "where was the conversation about...", or when looking for a decision, document, or discussion that could live in chat, email, cloud storage, or a project tracker.
argument-hint: "<query>"
---

# Search Command

> If you see unfamiliar placeholders or need to check which tools are connected, see [CONNECTORS.md](../../CONNECTORS.md).

Search across all connected MCP sources in a single query. Decompose the user's question, run parallel searches, and synthesize results.

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md),
> [LARK-RECIPES](../../../connectors/LARK-RECIPES.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md)).
> On Lark, the "sources" are concrete `lark_*` tools — fan them out in parallel, **project every read
> with `jq`** so a 25k-token doc fetch becomes ~500 (P3), resolve any named person to an `open_id`
> with `lark_contact_search` before filtering by author (P1), and when the result is a decision/digest
> rather than a raw lookup, surface it as an **interactive card** via `lark_im_card_send` (P4). For
> deep per-domain search (full Wiki tree, Base views, Minutes transcripts), delegate to the owning
> skill (`lark-doc`, `lark-base`, `lark-minutes`, `lark-im`).

## Instructions

### 1. Check Available Sources

On Lark, every source below is backed by the single `lark` MCP server (see CONNECTORS.md). The
category placeholders map to concrete tools:

| Source | Lark tools |
|--------|-----------|
| **~~chat** | `lark_im_search` (messages/threads) |
| **~~email** | `lark_mail_send`/`lark_mail_draft_create` for writes; mail *search* via `lark_api` (`GET /open-apis/mail/...`) |
| **~~cloud storage** | `lark_doc_search`, `lark_doc_fetch` (Drive/Docs); `lark_drive_upload` for writes |
| **~~project tracker** | `lark_task_my`, `lark_base_search` |
| **~~CRM** | a Lark **Base** acting as CRM — `lark_base_search` (read), `lark_base_record_upsert` (write); see P5 |
| **~~knowledge base** | `lark_doc_search` + `lark_doc_fetch` (Wiki-backed docs); `lark_minutes_search` for meeting knowledge; `lark_vc_search` for past-meeting envelopes |

Other reads worth fanning out when relevant: `lark_calendar_agenda` (events), `lark_okr_cycle_list`
(goals), `lark_sheets_read` (spreadsheet data). Anything without a curated tool → `lark_api`.

If the `lark` MCP server isn't connected:
```
To search across your workspace, the lark MCP server needs to be connected.
Check your MCP settings / run `lark-shared` to authenticate (auth login).
```

### 2. Parse the User's Query

Analyze the search query to understand:

- **Intent**: What is the user looking for? (a decision, a document, a person, a status update, a conversation)
- **Entities**: People, projects, teams, tools mentioned
- **Time constraints**: Recency signals ("this week", "last month", specific dates)
- **Source hints**: References to specific tools ("in chat", "that email", "the doc", "the meeting")
- **Filters**: Extract explicit filters from the query:
  - `from:` — Filter by sender/author
  - `in:` — Filter by channel, folder, or location
  - `after:` — Only results after this date
  - `before:` — Only results before this date
  - `type:` — Filter by content type (message, email, doc, thread, file)

### 3. Decompose into Sub-Queries

If the query filters by a person (`from:sarah`), **resolve the name to an `open_id` first** (P1):
`lark_contact_search(query="sarah")`. For each available source, build a targeted call with a `jq`
projection (P3):

**~~chat — `lark_im_search`:**
- `lark_im_search(query="<keywords>", jq=".data.messages[] | {id, chat, sender, text, create_time}")`
- `from:` → filter by the resolved sender `open_id`; `in:` → chat/group; dates → time range params.

**~~email — `lark_api` mail search:**
- `lark_api(method="GET", path="/open-apis/mail/...", params={...}, jq="...")` for inbox/thread search.
- Translate `from:` to sender, dates to time range. (Drafting/sending is `lark_mail_draft_create` /
  `lark_mail_send`, not used for read.) For non-trivial mail work delegate to the `lark-mail` skill.

**~~cloud storage — `lark_doc_search` then `lark_doc_fetch`:**
- `lark_doc_search(query="<name or full-text>", jq=".data.results[] | {token, title, url, type, edit_time}")`.
- Only `lark_doc_fetch` the top hits, and project: `jq=".data.title"` or `jq=".data.content"` to keep
  tokens small (P3). Map `type:` to doc type.

**~~project tracker — `lark_task_my` + `lark_base_search`:**
- `lark_task_my(jq=".data.items[] | {id, summary, due, completed}")` for the caller's tasks;
  `assignee_any: "me"` semantics are the default. Tasks assigned to *others* need the list endpoint
  (LARK-RECIPES → Tasks).
- `lark_base_search(base_token=..., table_id=..., query="<topic>", search_fields=["<field>"], select_fields=["<field>"], limit=20)`
  against the project Base/tracker. `lark_base_search` REQUIRES `search_fields` (the Bitable API mandates
  which field(s) the query matches) and does NOT support `jq` — narrow with `select_fields`/`limit` instead.

**~~CRM — `lark_base_search` (Base as system-of-record, P5):**
- `lark_base_search(base_token=..., table_id=..., query="<account/contact>", search_fields=["Name"], select_fields=["Name","Stage","Owner"], limit=20)`.
- `search_fields` is REQUIRED. If you don't know the field names, discover them first via
  `lark_api(method="GET", path="/open-apis/bitable/v1/apps/{base_token}/tables/{table_id}/fields")`.
  `lark_base_search` does NOT support `jq` — use `select_fields`/`limit` to control output. The Base *is*
  the CRM — for record-level reads/aggregations delegate to the `lark-base` skill rather than fetching
  raw records to count client-side.

**~~knowledge base — `lark_doc_search`/`lark_doc_fetch` + `lark_minutes_search`:**
- Wiki-backed docs: `lark_doc_search` (semantic for conceptual, keyword for exact), then `lark_doc_fetch`
  with a `jq` projection on the top hits.
- Meeting knowledge: `lark_minutes_search(query=..., participant_ids="me")` pulls AI summaries/action
  items directly — do not re-summarize (P6). `lark_vc_search` for the meeting envelope (who/when/room).

### 4. Execute Searches in Parallel

Issue all the `lark_*` / `lark_api` reads in a **single batch** (one assistant turn, multiple tool
calls) so total latency ≈ the slowest source, not the sum. Do not await one before firing the next.

For each source:
- Execute the projected (`jq`) call
- Capture results with metadata (timestamps, authors, links, source type)
- Note any source that returns an error envelope — do not let one failure block others. If an envelope
  says retry under a different identity, retry with bot identity (`--as`); see the `lark-shared` skill.

### 5. Rank and Deduplicate Results

**Deduplication:**
- Identify the same information appearing across sources (e.g., a decision discussed in IM AND confirmed via mail)
- Group related results together rather than showing duplicates
- Prefer the most authoritative or complete version

**Ranking factors:**
- **Relevance**: How well does the result match the query intent?
- **Freshness**: More recent results rank higher for status/decision queries
- **Authority**: Official docs > wiki > chat messages for factual questions; conversations > docs for "what did we discuss" queries
- **Completeness**: Results with more context rank higher

### 6. Present Unified Results

Synthesize into an answer (delegate the merge/dedup/rank logic to the **knowledge-synthesis** skill),
not a raw list. Default to an inline answer for a quick lookup. When the result is a **decision recap
or a multi-item digest** the user will act on, post it as an **interactive card** (P4) via
`lark_im_card_send` — `header` with a colored `template`, one `div`/`item` row per source with a link
button, a `note` footer with the source count. Always `print_json: true` to validate the spec, then
`dry_run: true`, then send (P2). Card grammar lives in the `lark-im` skill (`references/spec-reference.md`).

Inline formats:

**For factual/decision queries:**
```
[Direct answer to the question]

Sources:
- [Source 1: brief description] (IM, group, date)
- [Source 2: brief description] (Mail, from person, date)
- [Source 3: brief description] (Docs, doc name, last modified)
```

**For exploratory queries ("what do we know about X"):**
```
[Synthesized summary combining information from all sources]

Found across:
- IM: X relevant messages in Y groups
- Mail: X relevant threads
- Docs: X related documents
- Minutes: X meetings with related AI summaries
- [Other sources as applicable]

Key sources:
- [Most important source with link/reference]
- [Second most important source]
```

**For "find" queries (looking for a specific thing):**
```
[The thing they're looking for, with direct reference]

Also found:
- [Related items from other sources]
```

### 7. Handle Edge Cases

**Ambiguous queries:**
If the query could mean multiple things, ask one clarifying question before searching:
```
"API redesign" could refer to a few things. Are you looking for:
1. The REST API v2 redesign (Project Aurora)
2. The internal SDK API changes
3. Something else?
```

**No results:**
```
I couldn't find anything matching "[query]" across [list of lark sources searched].

Try:
- Broader terms (e.g., "database" instead of "PostgreSQL migration")
- Different time range (currently searching [time range])
- A different source — e.g. minutes (lark_minutes_search) or a specific Base view
```

**Partial results (some sources failed):**
```
[Results from successful sources]

Note: I couldn't reach [failed source(s)] during this search.
Results above are from [successful sources] only.
```

## Notes

- Always fan out the `lark_*` reads in parallel (one batched turn) — never sequentially
- Project every read with `jq` (P3) — never pull a full doc/record payload you won't use
- Synthesize results into answers (knowledge-synthesis skill), do not just list raw search results
- Include source attribution (chat/mail/doc/Base/minutes + date) so users can dig deeper
- When a query names a person, resolve to `open_id` via `lark_contact_search` first (P1), then filter by author
- For time-sensitive queries, prioritize recency in ranking
- For a decision recap or actionable digest, prefer an interactive card (`lark_im_card_send`, P4) over plain text
- For deep single-domain search, delegate to the owning skill (`lark-doc`, `lark-base`, `lark-minutes`, `lark-im`)
