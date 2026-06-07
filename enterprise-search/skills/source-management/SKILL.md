---
name: source-management
description: Manages connected MCP sources for enterprise search. Detects available sources, guides users to connect new ones, handles source priority ordering, and manages rate limiting awareness.
user-invocable: false
---

# Source Management

> If you see unfamiliar placeholders or need to check which tools are connected, see [CONNECTORS.md](../../CONNECTORS.md).

Knows what sources are available, helps connect new ones, and manages how sources are queried.

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md),
> [LARK-FUSION](../../../connectors/LARK-FUSION.md)). On Lark all collaboration sources are backed by
> the **single `lark` MCP server**, so "connecting sources" is really one auth step (`lark-shared` /
> `auth login`) rather than wiring six servers. Specialty external sources (data warehouse, analytics,
> payments, scientific DBs) keep their own MCP server. Identity is resolved with `lark_contact_search`
> (P1); a Lark **Base** stands in for any CRM/tracker (P5).

## Checking Available Sources

All collaboration sources resolve to the `lark` MCP server. Each maps to concrete `lark_*` tools:

| Source | Lark tool(s) | Key capabilities |
|--------|-------------|-----------------|
| **~~chat** | `lark_im_search` | Search messages, read groups and threads |
| **~~email** | `lark_api` mail search; `lark_mail_*` to write | Search/read; draft/send (guarded) |
| **~~cloud storage** | `lark_doc_search`, `lark_doc_fetch`, `lark_drive_upload` | Search files, fetch contents, upload |
| **~~project tracker** | `lark_task_my`, `lark_base_search` | Search tasks and tracker Base records |
| **~~CRM** | `lark_base_search`, `lark_base_record_upsert` | Query/upsert accounts, contacts, opportunities (Base, P5) |
| **~~knowledge base** | `lark_doc_search`/`lark_doc_fetch`, `lark_minutes_search`, `lark_okr_cycle_list` | Wiki docs, meeting knowledge, goals |

If the `lark` tools respond, the workspace is connected and searchable. Specialty external tools are
detected by their own MCP prefix.

## Guiding Users to Connect Sources

When the `lark` MCP server isn't authenticated (collaboration sources unavailable):

```
The lark MCP server isn't authenticated, so I can't reach your workspace yet.
Run `auth login` (see the lark-shared skill) to connect — that single step enables
IM, Mail, Drive/Docs, Wiki, Tasks, Base, Calendar, Minutes, and OKR all at once.
```

When a user asks about a **specialty external source** (data warehouse, analytics, payments,
scientific DB) that has no Lark equivalent:

```
[Tool name] is a specialty source — it keeps its own MCP server (the lark server doesn't cover it).
To add it:
1. Open your MCP settings (.mcp.json)
2. Add the [tool] MCP server configuration
3. Authenticate when prompted

Once connected, it will be included alongside the lark sources in future searches.
```

If a Lark `lark_*` call returns a permission/scope error envelope, the issue is identity or scopes,
not connectivity — retry with bot identity (`--as`) or fix scopes per the **`lark-shared`** skill.

## Source Priority Ordering

Different query types benefit from searching certain sources first. Use these priorities to weight results, not to skip sources:

### By Query Type

**Decision queries** ("What did we decide..."):
```
1. lark_minutes_search (meeting AI decisions/action items, P6)
2. lark_im_search (conversations where decisions happen)
3. decision-log Base — lark_base_search (P5)
4. lark_doc_search (decision docs in Wiki)
5. lark_task_my / Base tracker (if captured as tasks)
```

**Status queries** ("What's the status of..."):
```
1. lark_task_my / lark_base_search (tracker — authoritative status)
2. lark_im_search (real-time discussion)
3. lark_doc_search (status docs, reports)
4. lark_api mail search (status update emails)
```

**Document queries** ("Where's the doc for..."):
```
1. lark_doc_search → lark_doc_fetch (Drive/Wiki/Docs — primary store)
2. lark_minutes_search (meeting docs/transcripts)
3. lark_api mail search (docs shared via email)
4. lark_im_search (docs shared in chats)
```

**People queries** ("Who works on..." / "Who knows about..."):
```
1. lark_contact_search (resolve the person, P1)
2. lark_im_search (message authors)
3. lark_task_my / Base tracker (task assignees)
4. lark_doc_search (doc authors/collaborators)
5. CRM Base — lark_base_search (account owners, contacts)
```

**Factual/Policy queries** ("What's our policy on..."):
```
1. lark_doc_search → lark_doc_fetch (Wiki — official documentation)
2. lark_api mail search (policy announcements)
3. lark_im_search (policy discussions)
```

### Default Priority (General Queries)

When query type is unclear:
```
1. lark_im_search (highest volume, most real-time)
2. lark_api mail search (formal communications)
3. lark_doc_search (documents and files)
4. lark_doc_search / lark_minutes_search (structured + meeting knowledge)
5. lark_task_my / lark_base_search (work items)
6. CRM Base — lark_base_search (customer data)
```

## Rate Limiting Awareness

MCP sources may have rate limits. Handle them gracefully:

### Detection

Rate limit responses typically appear as:
- HTTP 429 responses
- Error messages mentioning "rate limit", "too many requests", or "quota exceeded"
- Throttled or delayed responses

### Handling

When a source is rate limited:

1. **Do not retry immediately** — respect the limit
2. **Continue with other sources** — do not block the entire search
3. **Inform the user**:
```
Note: [Source] is temporarily rate limited. Results below are from
[other sources]. You can retry in a few minutes to include [source].
```
4. **For digests** — if rate limited mid-scan, note which time range was covered before the limit hit

### Prevention

- Avoid unnecessary API calls — check if the source is likely to have relevant results before querying
- Use targeted queries over broad scans when possible
- For digests, batch requests where the API supports it
- Cache awareness: if a search was just run, avoid re-running the same query immediately

## Source Health

Track source availability during a session:

```
Source Status:
  IM (lark_im_search):              ✓ Available
  Mail (lark_api mail):             ✓ Available
  Drive/Docs (lark_doc_search):     ✓ Available
  Tasks/Base (lark_task_my):        ✓ Available
  Minutes (lark_minutes_search):    ✓ Available
  CRM Base (lark_base_search):      ✗ Not configured (no CRM Base — scaffold via base-deploy)
```

When reporting search results, include which `lark_*` sources were searched so the user knows the
scope of the answer.

## Adding Custom Sources

The collaboration layer is all served by the `lark` MCP server — once authenticated, every `lark_*`
source is searchable, nothing else to wire.

To add a **specialty external source** (data warehouse, analytics, payments, scientific DB — no Lark
equivalent):
1. Add the specialty MCP server configuration to `.mcp.json`
2. Authenticate if required
3. It is included alongside the lark sources in subsequent searches automatically

To stand up a missing **CRM / tracker / decision-log** "source", don't add an external server —
scaffold a Lark **Base** with the **`base-deploy`** skill (P5); it then becomes a `lark_base_search`
source like any other.
