---
name: process-optimization
description: Analyze and improve business processes. Trigger with "this process is slow", "how can we improve", "streamline this workflow", "too many steps", "bottleneck", or when the user describes an inefficient process they want to fix.
---

# Process Optimization

Analyze existing processes and recommend improvements.

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md),
> [LARK-RECIPES](../../../connectors/LARK-RECIPES.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md)).
> Map the current state from **real Lark signals**, not guesses: pull the existing SOP from Wiki/Docs
> (`lark_doc_search`/`lark_doc_fetch`, jq-projected, P3), measure approval-step waiting from Lark
> Approval, and quantify handoffs from Lark Task cycle times. Surface the before/after + recommendations
> as an **interactive card** (P4), spin improvement items into Lark Tasks (P2 dry_run), and update the
> SOP in **Wiki** (P8).

## Analysis Framework

### 1. Map Current State
- Document every step, decision point, and handoff
- Identify who does what and how long each step takes
- Note manual steps, approvals, and waiting times

### 2. Identify Waste
- **Waiting**: Time spent in queues or waiting for approvals
- **Rework**: Steps that fail and need to be redone
- **Handoffs**: Each handoff is a potential point of failure or delay
- **Over-processing**: Steps that add no value
- **Manual work**: Tasks that could be automated

### 3. Design Future State
- Eliminate unnecessary steps
- Automate where possible
- Reduce handoffs
- Parallelize independent steps
- Add checkpoints (not gates)

### 4. Measure Impact
- Time saved per cycle
- Error rate reduction
- Cost savings
- Employee satisfaction improvement

## Lark-native execution

**Map current state from real data**:
- Existing SOP: `lark_doc_search` (jq `.data.results[].title`) → `lark_doc_fetch` (jq `.data.content[:4000]`) to recover the documented steps (P3).
- Waiting/approval steps: if the process runs on a Lark Approval flow, delegate to **`approval-flow-sla`** to find the bottleneck node and per-approver dwell time — that IS the "waiting waste" measurement.
- Handoffs/rework: read the process tracker Base with `lark_base_search` (REQUIRES `search_fields`, no jq — narrow with `select_fields`; discover field names via `lark_api GET /open-apis/bitable/v1/apps/{base}/tables/{table}/fields` if unknown) or Lark Task cycle times to quantify how long items sit between handoffs.

**Quantify and recommend (card, P4)**:
- Post the before/after comparison + top improvements as an **interactive card** (`lark_im_card_send`, `print_json: true` first): status pills per step (🟢 keep / 🟡 streamline / 🔴 cut), `actions` footer ("Apply changes" / "Open SOP").

**Execute the improvement plan**:
- For each improvement, resolve the owner (`lark_contact_search`, P1) → `lark_task_create(dry_run)` → confirm → commit.
- For automatable steps in a Base-driven process, delegate to **`base-automation`** to build the Base workflow/bot reminder rather than describing automation abstractly.

**Update the SOP (Wiki, P8)**:
- Write the future-state SOP back to **Wiki** (`lark_wiki_node_create` / `lark-doc` skill) so the documented process matches the optimized one.

## Output

Produce a before/after process comparison with specific improvement recommendations, estimated impact, and an implementation plan — measured from real Lark signals, surfaced as a card, executed as Lark Tasks, and reflected back into the Wiki SOP.
