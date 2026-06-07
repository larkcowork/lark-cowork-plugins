---
name: recruiting-pipeline
description: Track and manage recruiting pipeline stages. Trigger with "recruiting update", "candidate pipeline", "how many candidates", "hiring status", or when the user discusses sourcing, screening, interviewing, or extending offers.
---

# Recruiting Pipeline

Help manage the recruiting pipeline from sourcing through offer acceptance.

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md),
> [LARK-RECIPES](../../../connectors/LARK-RECIPES.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md)).
> Candidate data ideally lives in the `~~ATS` specialty MCP (keep as-is). With no ATS, **the pipeline
> is a Lark Base** as system-of-record (P5): one row per candidate (role, stage, source, owner,
> dates) — read with `lark_base_search`, advance a stage with `lark_base_record_upsert` (`dry_run`
> first, P2), and compute stage counts / conversion / time-in-stage with the Base aggregation
> endpoint (don't count client-side). Scaffold that pipeline Base with **`base-deploy`**, not by
> hand. Resolve recruiters/hiring managers/candidates via `lark_contact_search` (P1). The
> stage-by-stage status read with stuck-deal / aging logic is exactly what the **`pipeline-review`**
> skill does — **delegate to `pipeline-review`** for the recurring scan rather than reinventing it.
> Per-stage handoffs become **Lark Tasks** (`lark_task_create`); interviews go through
> **`/interview-prep`**; offers through **`/draft-offer`**. Surface the pipeline snapshot
> (by-stage, stuck candidates, offers out) as an **interactive card** (`lark_im_card_send`, P4).

## Pipeline Stages

| Stage | Description | Key Actions |
|-------|-------------|-------------|
| Sourced | Identified and reached out | Personalized outreach |
| Screen | Phone/video screen | Evaluate basic fit |
| Interview | On-site or panel interviews | Structured evaluation |
| Debrief | Team decision | Calibrate feedback |
| Offer | Extending offer | Comp package, negotiation |
| Accepted | Offer accepted | Transition to onboarding |

## Metrics to Track

- **Pipeline velocity**: Days per stage
- **Conversion rates**: Stage-to-stage drop-off
- **Source effectiveness**: Which channels produce hires
- **Offer acceptance rate**: Offers extended vs. accepted
- **Time to fill**: Days from req open to offer accepted

## If ATS Connected

If **~~ATS** (specialty external MCP) is connected, pull candidate data automatically and update
statuses there. Otherwise treat a **Lark Base** as the pipeline of record:

- **Read:** `lark_base_search` — it does NOT support `jq` and REQUIRES `search_fields`; pass the
  field(s) to match (e.g. stage/owner) and narrow output with `select_fields`/`limit`. Discover field
  names via `lark_api GET /open-apis/bitable/v1/apps/{base}/tables/{table}/fields` if unknown.
- **Advance a candidate:** `lark_base_record_upsert` (new stage + timestamp), `dry_run` first (P2).
- **Metrics:** stage counts, conversion, time-to-fill via the Base aggregation endpoint.
- **Scaffold:** build the pipeline Base with **`base-deploy`** if none exists; record/field depth via
  **`lark-base`**.

### Lark-native delivery

1. **Resolve people (P1):** `lark_contact_search` for recruiters, hiring managers, candidates.
2. **Recurring scan:** delegate the weekly by-stage / stuck-candidate / aging review to the
   **`pipeline-review`** skill (point it at the recruiting Base).
3. **Handoffs (P2):** each stage transition that needs human action → `lark_task_create` to the
   owner (e.g. "schedule onsite for X"). Interviews → **`/interview-prep`**; offers →
   **`/draft-offer`** (which gates comp via a Lark approval, P7).
4. **Report (P4):** post the pipeline snapshot (by-stage funnel, stuck candidates, offers
   out/accepted) as an `lark_im_card_send` card (`print_json` → `dry_run` → send).
