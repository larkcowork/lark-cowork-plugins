# Lark API Recipes (for `lark_api` passthrough)

The `lark` MCP server exposes 25 curated tools. Anything outside that set goes through
**`lark_api`** (generic Lark OpenAPI passthrough: `method` + `path` + `params`/`data`, with
`dry_run` and `jq`). This file gives the exact calls the plugins need most, so skills **never
have to guess** an endpoint. Paths verified against the lark-cli `shortcuts/` source unless marked
⚠️ (confirm via the `lark-openapi-explorer` skill or the noted deep skill before committing).

> Convention: always `dry_run: true` a mutating call first, inspect the planned request, then
> re-issue without `dry_run`. Resolve every `ou_`/`open_id` via `lark_contact_search` first.

## Tasks (curated tools cover create + list; these fill the gaps)

| Need | `lark_api` call |
|------|-----------------|
| Complete a task | ✅ now a curated tool: **`lark_task_complete`** (`task_id`) — prefer it over raw API |
| Add a comment | `POST /open-apis/task/v2/comments` · `data:{ "resource_type":"task","resource_id":"{guid}","content":"..." }` |
| Create a tasklist | `POST /open-apis/task/v2/tasklists` · `data:{ "name":"...", "members":[{"id":"ou_..","role":"editor"}] }` |
| Tasks assigned to others | `GET /open-apis/task/v2/tasks` (list) — `lark_task_my` only returns the caller's |

## Lark Base / Bitable (system-of-record for trackers, CRM, decision logs)

`lark_base_search` reads. **Two gotchas (verified live):** (1) the Bitable API **requires
`search_fields`** — you must say which field(s) to match the keyword in; if you don't know the
field names, call `lark_api GET /open-apis/bitable/v1/apps/{base}/tables/{table}/fields` first.
(2) `lark_base_search` returns **JSON and does NOT support `jq`** (the base shortcut defaults to
markdown which is jq-incompatible) — narrow output with `select_fields` + `limit` instead.

Writes go through `lark_api` (⚠️ confirm shapes via the **`lark-base`**
skill, which is the authoritative source and also handles field/table/view creation):

| Need | `lark_api` call |
|------|-----------------|
| Create/update one record | ✅ now a curated tool: **`lark_base_record_upsert`** (`base_token`, `table_id`, `fields`) — prefer it |
| (raw) Create a record | `POST /open-apis/bitable/v1/apps/{app_token}/tables/{table_id}/records` · `data:{ "fields": { ... } }` |
| Batch create | `POST .../records/batch_create` · `data:{ "records":[{"fields":{...}}, ...] }` |
| Aggregations/counts | use the Base data-query endpoint — don't fetch raw records to count client-side |
| **Build a new Base/table** | don't hand-roll — invoke the **`base-deploy`** skill (phased builder with verification) |

## Wiki (knowledge base)

| Need | `lark_api` call |
|------|-----------------|
| Create a wiki node (page) | ✅ now a curated tool: **`lark_wiki_node_create`** (`space_id`, `title`, …) — prefer it |
| Get a node | `GET /open-apis/wiki/v2/spaces/get_node?token={token}` |
| List space members | `GET /open-apis/wiki/v2/spaces/{space_id}/members` |

(After creating a node you get a `docx` document — fill it with the docx block recipe below, or
`lark_doc_create` with `wiki_node`+`wiki_space`.)

## Docs (append to an existing doc — create is covered by `lark_doc_create`)

| Need | `lark_api` call |
|------|-----------------|
| Append/insert blocks | `POST /open-apis/docx/v1/documents/{doc_id}/blocks/{block_id}/children` · `data:{ "children":[ ...block JSON... ], "index":-1 }` |
| Batch update blocks | `PATCH /open-apis/docx/v1/documents/{doc_id}/blocks/batch_update` |

For non-trivial doc editing use the **`lark-doc`** skill (DocxXML) — block JSON by hand is brittle.

## Calendar (free/busy + rooms — agenda/create are curated tools)

| Need | `lark_api` call |
|------|-----------------|
| Free/busy of one person | ✅ now a curated tool: **`lark_calendar_freebusy`** (`start`, `end`, `user_id`) — prefer it |
| Free/busy of many people | `POST /open-apis/calendar/v4/freebusy/list` · `data:{ "time_min":"...","time_max":"...","user_id_list":["ou_.."] }` |
| Suggest a meeting slot | `POST /open-apis/calendar/v4/freebusy/suggestion` |
| Find a free room | `POST /open-apis/calendar/v4/freebusy/room_find` |

## Approval (no curated tool — used by triage/escalation/expense flows)

⚠️ Confirm exact shapes via the **`lark-approval`** skill before committing:

| Need | `lark_api` call |
|------|-----------------|
| Create an approval instance | `POST /open-apis/approval/v4/instances` · `data:{ "approval_code":"...", "form":"<json>", "open_id":"ou_.." }` |
| Query an instance | `GET /open-apis/approval/v4/instances/{instance_id}` |
| Approve / reject a task | `POST /open-apis/approval/v4/tasks/approve` (or `/reject`) |

## Tool-gap audit → candidates to add to `cmd/mcp/tools.go`

Status — 4 of the 6 candidates are now **shipped as curated tools** in `cmd/mcp/tools.go`
(flag names verified verbatim against `shortcuts/<domain>/*.go`; `go test ./cmd/mcp` green; bridge
now exposes **25 tools**):

1. ✅ `lark_task_complete` — wraps `task +complete`.
2. ✅ `lark_base_record_upsert` — wraps `base +record-upsert` (create-or-update; covers the
   create + update cases in one tool).
3. ✅ `lark_wiki_node_create` — wraps `wiki +node-create`.
4. ✅ `lark_calendar_freebusy` — wraps `calendar +freebusy`.
5. ⏳ `lark_doc_append` — **not shipped**: `shortcuts/doc/` has no append verb (only
   `+update`/`+media-insert`, semantics unverified). Use the `lark-doc` skill or `lark_api` docx
   `blocks/.../children` until verified.
6. ⏳ `lark_approval_create` — **not shipped**: no `shortcuts/approval/` exists, so there is no
   verified flag set to wrap. Use the `lark-approval` skill or the `lark_api` approval recipe above.

Each shipped tool follows the existing builder convention in `tools.go`. See [LARK-PATTERNS.md](./LARK-PATTERNS.md)
for how these compose into workflows, and [LARK-FUSION.md](./LARK-FUSION.md) for which deep skill
owns each domain.
