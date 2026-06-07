---
name: org-planning
description: Headcount planning, org design, and team structure optimization. Trigger with "org planning", "headcount plan", "team structure", "reorg", "who should we hire next", or when the user is thinking about team size, reporting structure, or organizational design.
---

# Org Planning

Help plan organizational structure, headcount, and team design.

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md),
> [LARK-RECIPES](../../../connectors/LARK-RECIPES.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md)).
> Current headcount/reporting data comes from the `~~HRIS` specialty MCP (or its **Lark Base**
> approximation — read via `lark_base_search`); spot-resolve managers/ICs with `lark_contact_search`
> (P1). Hold the **headcount plan + sequenced hiring roadmap in a Lark Base** as system-of-record
> (P5): each open req is a row (role, level, team, target start, status, cost) — write with
> `lark_base_record_upsert` (`dry_run` first, P2); scaffold the planning Base via **`base-deploy`**.
> Tie roles to goals with `lark_okr_cycle_list` when sequencing against company objectives. Land the
> org-design narrative + cost model in **Wiki** (`lark_wiki_node_create` / `lark_doc_create`, P8;
> author via **`lark-doc`**). Open reqs for the next quarter → one `lark_task_create` per hire to the
> recruiter (P2). Surface the headcount snapshot + structural warnings (top-heavy, single points of
> failure) as an **interactive card** (`lark_im_card_send`, P4) to leadership.

## Planning Dimensions

- **Headcount**: How many people do we need, in what roles, by when?
- **Structure**: Reporting lines, span of control, team boundaries
- **Sequencing**: Which hires are most critical? What's the right order?
- **Budget**: Headcount cost modeling and trade-offs

## Healthy Org Benchmarks

| Metric | Healthy Range | Warning Sign |
|--------|---------------|--------------|
| Span of control | 5-8 direct reports | < 3 or > 12 |
| Management layers | 4-6 for 500 people | Too many = slow decisions |
| IC-to-manager ratio | 6:1 to 10:1 | < 4:1 = top-heavy |
| Team size | 5-9 people | < 4 = lonely, > 12 = hard to manage |

## Output

Produce org charts (text-based), headcount plans with cost modeling, and sequenced hiring roadmaps. Flag structural issues like single points of failure or excessive management overhead.

### Lark-native delivery

1. **Read current org (P5):** `lark_base_search` the HRIS/org Base — it does NOT support `jq` and
   REQUIRES `search_fields`; pass the field(s) to match and narrow output with `select_fields`/`limit`
   (headcount, manager, level, team). Discover field names via `lark_api GET
   /open-apis/bitable/v1/apps/{base}/tables/{table}/fields` if unknown. Resolve any individual with
   `lark_contact_search` (P1).
2. **Align to objectives (P6/OKR):** `lark_okr_cycle_list` to sequence hires against the active
   cycle's objectives.
3. **Persist the plan (P5):** each open req → `lark_base_record_upsert` row in the headcount Base
   (`dry_run` first). No Base yet → **`base-deploy`**. Use the Base data-query endpoint for
   span-of-control / ratio aggregations rather than counting client-side.
4. **Document (P8):** org-design rationale + cost model → `lark_wiki_node_create` / `lark_doc_create`
   (author via **`lark-doc`**).
5. **Kick off hiring (P2):** one `lark_task_create` per prioritized req, assigned to the recruiter.
6. **Brief leadership (P4):** post the headcount snapshot + structural warnings as an
   `lark_im_card_send` card (`print_json` → `dry_run` → send).
