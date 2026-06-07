---
name: search-strategy
description: Query decomposition and multi-source search orchestration. Breaks natural language questions into targeted searches per source, translates queries into source-specific syntax, ranks results by relevance, and handles ambiguity and fallback strategies.
user-invocable: false
---

# Search Strategy

> If you see unfamiliar placeholders or need to check which tools are connected, see [CONNECTORS.md](../../CONNECTORS.md).

The core intelligence behind enterprise search. Transforms a single natural language question into parallel, source-specific searches and produces ranked, deduplicated results.

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md),
> [LARK-RECIPES](../../../connectors/LARK-RECIPES.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md)).
> On Lark each "source" is a concrete `lark_*` tool. Resolve any named person to an `open_id` via
> `lark_contact_search` before applying an author filter (P1); project every read with `jq` (P3);
> fan all reads out in one batched turn. For deep single-domain search, delegate to the owning skill
> (`lark-im`, `lark-doc`, `lark-base`, `lark-minutes`).

## The Goal

Turn this:
```
"What did we decide about the API migration timeline?"
```

Into targeted, projected `lark_*` calls across every connected source:
```
lark_im_search(query="API migration timeline decision",
               jq=".data.messages[] | {chat, sender, text, create_time}")
lark_doc_search(query="API migration timeline rationale",
               jq=".data.results[] | {token, title, url, edit_time}")          # Wiki/Docs
lark_minutes_search(query="API migration", participant_ids="me",
               jq=".data.items[] | {title, summary, action_items}")          # meeting decisions (P6)
lark_base_search(query="API migration", search_fields=["Topic"], select_fields=["Topic","Decision","Owner"], limit=20)  # tracker/decision-log (P5); no jq, search_fields required
```

Then synthesize the results into a single coherent answer (knowledge-synthesis skill), and if it is a
decision recap, surface it as an interactive card (`lark_im_card_send`, P4).

## Query Decomposition

### Step 1: Identify Query Type

Classify the user's question to determine search strategy:

| Query Type | Example | Strategy (Lark tools) |
|-----------|---------|----------|
| **Decision** | "What did we decide about X?" | `lark_minutes_search` (meeting decisions, P6) + `lark_im_search` + decision-log Base (`lark_base_search`); look for conclusion signals |
| **Status** | "What's the status of Project Y?" | `lark_task_my` / `lark_base_search` (tracker = authoritative) + recent `lark_im_search` |
| **Document** | "Where's the spec for Z?" | `lark_doc_search` then `lark_doc_fetch` (Drive/Wiki/Docs) |
| **Person** | "Who's working on X?" | `lark_contact_search` (resolve), then task assignees (Base/task list), `lark_im_search` authors, doc collaborators |
| **Factual** | "What's our policy on X?" | `lark_doc_search`/`lark_doc_fetch` (Wiki/official docs) then confirmatory `lark_im_search` |
| **Temporal** | "When did X happen?" | `lark_im_search` / `lark_minutes_search` / `lark_calendar_agenda` with a broad time range |
| **Exploratory** | "What do we know about X?" | Fan out all of the above in parallel, synthesize |

### Step 2: Extract Search Components

From the query, extract:

- **Keywords**: Core terms that must appear in results
- **Entities**: People, projects, teams, tools (use memory system if available)
- **Intent signals**: Decision words, status words, temporal markers
- **Constraints**: Time ranges, source hints, author filters
- **Negations**: Things to exclude

### Step 3: Generate Sub-Queries Per Source

For each available source, create one or more targeted queries:

**Prefer semantic search** for:
- Conceptual questions ("What do we think about...")
- Questions where exact keywords are unknown
- Exploratory queries

**Prefer keyword search** for:
- Known terms, project names, acronyms
- Exact phrases the user quoted
- Filter-heavy queries (from:, in:, after:)

**Generate multiple query variants** when the topic might be referred to differently:
```
User: "Kubernetes setup"
Queries: "Kubernetes", "k8s", "cluster", "container orchestration"
```

## Source-Specific Query Translation

### ~~chat — `lark_im_search`

```
lark_im_search(query="project aurora status update",
               jq=".data.messages[] | {chat, sender, text, create_time}")
```

**Filter mapping** (resolve `from:` to an `open_id` via `lark_contact_search` first, P1):
| Enterprise filter | `lark_im_search` mapping |
|------------------|--------------|
| `from:sarah` | resolved sender `open_id` filter |
| `in:engineering` | chat / group id |
| `after:` / `before:` | time-range params |
| `type:thread` / `type:file` | thread / has-attachment filter (else `lark_api` IM search) |

### ~~knowledge base — `lark_doc_search` + `lark_doc_fetch` + `lark_minutes_search`

**Semantic** (conceptual) and **keyword** (exact) both go to `lark_doc_search`; project hits, then
`lark_doc_fetch` only the top results with `jq=".data.content"` (P3):
```
lark_doc_search(query="API migration timeline rationale",
                jq=".data.results[] | {token, title, url, type, edit_time}")
lark_minutes_search(query="API migration", participant_ids="me",
                jq=".data.items[] | {title, summary, action_items}")   # meeting knowledge (P6)
```

### ~~project tracker — `lark_task_my` + `lark_base_search`

```
lark_task_my(complete=false, jq=".data.items[] | {id, summary, due, completed}")   # caller's tasks
lark_base_search(base_token=..., table_id=..., query="API migration",
                 search_fields=["Title"], select_fields=["Title","Status","Owner","Due"], limit=20)  # project tracker Base; search_fields required, no jq
```

**Filter mapping:**
| Enterprise filter | Lark mapping |
|------------------|----------------|
| `assignee_any: me` | `lark_task_my` (caller's tasks); for others use task list endpoint (LARK-RECIPES) |
| `from:sarah` | resolve `open_id`, then add `Owner`/`Created by` to Base `search_fields` |
| `after:` / status | Base view filter / `lark_base_search` `select_fields` + `limit` control |

### ~~CRM — `lark_base_search` (Base as system-of-record, P5)

```
lark_base_search(base_token=<crm>, table_id=<accounts>, query="Acme",
                 search_fields=["Account"], select_fields=["Account","Stage","Owner","LastTouch"], limit=20)
```
`search_fields` is REQUIRED and `lark_base_search` does NOT support `jq`. If you don't know the field
names, discover them first via
`lark_api(method="GET", path="/open-apis/bitable/v1/apps/{base_token}/tables/{table_id}/fields")`.
The CRM *is* a Lark Base. For record-level reads, aggregations, or schema work, delegate to `lark-base`
(don't fetch raw records to count client-side).

## Result Ranking

### Relevance Scoring

Score each result on these factors (weighted by query type):

| Factor | Weight (Decision) | Weight (Status) | Weight (Document) | Weight (Factual) |
|--------|-------------------|------------------|--------------------|-------------------|
| Keyword match | 0.3 | 0.2 | 0.4 | 0.3 |
| Freshness | 0.3 | 0.4 | 0.2 | 0.1 |
| Authority | 0.2 | 0.1 | 0.3 | 0.4 |
| Completeness | 0.2 | 0.3 | 0.1 | 0.2 |

### Authority Hierarchy

Depends on query type:

**For factual/policy questions:**
```
Wiki docs (lark_doc_*) > Shared docs > Mail announcements > IM messages
```

**For "what happened" / decision questions:**
```
Minutes AI summary (lark_minutes_search) > IM thread conclusions > Mail confirmations > IM messages
```

**For status questions:**
```
Task / Base tracker (lark_task_my, lark_base_search) > Recent IM > Status docs > Mail updates
```

## Handling Ambiguity

When a query is ambiguous, prefer asking one focused clarifying question over guessing:

```
Ambiguous: "search for the migration"
→ "I found references to a few migrations. Are you looking for:
   1. The database migration (Project Phoenix)
   2. The cloud migration (AWS → GCP)
   3. The email migration (Exchange → O365)"
```

Only ask for clarification when:
- There are genuinely distinct interpretations that would produce very different results
- The ambiguity would significantly affect which sources to search

Do NOT ask for clarification when:
- The query is clear enough to produce useful results
- Minor ambiguity can be resolved by returning results from multiple interpretations

## Fallback Strategies

When a source is unavailable or returns no results:

1. **Source unavailable**: Skip it, search remaining sources, note the gap
2. **No results from a source**: Try broader query terms, remove date filters, try alternate keywords
3. **All sources return nothing**: Suggest query modifications to the user
4. **Rate limited**: Note the limitation, return results from other sources, suggest retrying later

### Query Broadening

If initial queries return too few results:
```
Original: "PostgreSQL migration Q2 timeline decision"
Broader:  "PostgreSQL migration"
Broader:  "database migration"
Broadest: "migration"
```

Remove constraints in this order:
1. Date filters (search all time)
2. Source/location filters
3. Less important keywords
4. Keep only core entity/topic terms

## Parallel Execution

Always execute searches across sources in parallel, never sequentially. The total search time should be roughly equal to the slowest single source, not the sum of all sources.

Fire them all in one batched assistant turn (multiple tool calls), each with a `jq` projection (P3):
```
[User query]
     ↓ decompose (resolve people via lark_contact_search, P1)
[lark_im_search] [lark_api mail] [lark_doc_search] [lark_minutes_search] [lark_task_my / lark_base_search]
     ↓            ↓            ↓              ↓            ↓
  (parallel execution — one batched turn)
     ↓
[Merge + Rank + Deduplicate]   ← knowledge-synthesis skill
     ↓
[Synthesized answer  (decision recap → lark_im_card_send card, P4)]
```
