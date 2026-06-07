---
name: research-synthesis
description: Synthesize user research into themes, insights, and recommendations. Use when you have interview transcripts, survey results, usability test notes, support tickets, or NPS responses that need to be distilled into patterns, user segments, and prioritized next steps.
argument-hint: "<research data, transcripts, or survey results>"
---

# /research-synthesis

> If you see unfamiliar placeholders or need to check which tools are connected, see [CONNECTORS.md](../../CONNECTORS.md).

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md),
> [LARK-RECIPES](../../../connectors/LARK-RECIPES.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md)).
> Synthesis logic unchanged; inputs and outputs become Lark. **Pull real inputs:** interview audio →
> AI transcript/summary/action-items via `lark_minutes_search` (participant_ids="me", P6) — don't
> re-summarize from scratch; support tickets / NPS from a feedback **Lark Base** (`lark_base_search`,
> P5) or Intercom (external); validate qualitative themes against product analytics (external MCP).
> **Land outputs:** publish the synthesis report to the research repository in **Wiki**
> (`lark_wiki_node_create`, P8) via **`lark-doc`**; track Insight→Opportunity rows in a research-findings
> **Lark Base** (`lark_base_record_upsert`; scaffold via **base-deploy**); turn recommendations into
> **Lark Tasks** (`lark_task_create`, owner via `lark_contact_search` P1, `dry_run` P2); share the
> executive summary as an **interactive card** (`lark_im_card_send`, P4).

Synthesize user research data into actionable insights. See the **user-research** skill for research methods, interview guides, and analysis frameworks.

## Usage

```
/research-synthesis $ARGUMENTS
```

## What I Accept

- Interview transcripts or notes
- Survey results (CSV, pasted data)
- Usability test recordings or notes
- Support tickets or feedback
- NPS/CSAT responses
- App store reviews

## Output

```markdown
## Research Synthesis: [Study Name]
**Method:** [Interviews / Survey / Usability Test] | **Participants:** [X]
**Date:** [Date range] | **Researcher:** [Name]

### Executive Summary
[3-4 sentence overview of key findings]

### Key Themes

#### Theme 1: [Name]
**Prevalence:** [X of Y participants]
**Summary:** [What this theme is about]
**Supporting Evidence:**
- "[Quote]" — P[X]
- "[Quote]" — P[X]
**Implication:** [What this means for the product]

#### Theme 2: [Name]
[Same format]

### Insights → Opportunities

| Insight | Opportunity | Impact | Effort |
|---------|-------------|--------|--------|
| [What we learned] | [What we could do] | High/Med/Low | High/Med/Low |

### User Segments Identified
| Segment | Characteristics | Needs | Size |
|---------|----------------|-------|------|
| [Name] | [Description] | [Key needs] | [Rough %] |

### Recommendations
1. **[High priority]** — [Why, based on which findings]
2. **[Medium priority]** — [Why]
3. **[Lower priority]** — [Why]

### Questions for Further Research
- [What we still don't know]

### Methodology Notes
[How the research was conducted, any limitations or biases to note]
```

## Lark-native inputs & delivery

**Pull research inputs from Lark first.**
- **Interviews / usability sessions recorded in Lark:** `lark_minutes_search(participant_ids="me",
  query=...)` → use the AI transcript, summary, and action items directly (P6) rather than
  re-transcribing. For meeting envelope/participants only, `lark_vc_search`. Delegate deep artifact
  pulls to the **`lark-minutes`** skill.
- **Support tickets / NPS / feature requests (user feedback):** from a feedback **Lark Base** via
  `lark_base_search` (P5 — pass `search_fields` for the column(s) to match; `lark_base_search` does
  NOT support `jq`, so narrow with `select_fields`/`limit`. If the field names are unknown, discover
  them first via `lark_api GET /open-apis/bitable/v1/apps/{base}/tables/{table}/fields`), or Intercom
  (external MCP) if connected.
- **Prior studies:** `lark_doc_search` the research repository before synthesizing, to compare.

**Product analytics (external — keep as-is):** validate qualitative themes against usage/behavioral
metrics via the analytics MCP; quantify pain-point impact. Lark has no analytics warehouse — stays
external.

**Land outputs in Lark.**
- **Publish the synthesis to Wiki (P8):** `lark_wiki_node_create` → fill via **`lark-doc`** so it
  lives in the research repository.
- **Track Insight→Opportunity rows in a research-findings Lark Base (P5):** `lark_base_record_upsert`
  (`fields:{ Insight, Opportunity, Impact, Effort, Status }`, `dry_run` first); scaffold via
  **base-deploy** if none exists.
- **Turn recommendations into Lark Tasks:** `lark_contact_search` owner → `open_id` (P1) →
  `lark_task_create` (`dry_run`, P2); delegate to **`lark-task`**.
- **Share the executive summary as an interactive card (P4):** `lark_im_card_send` (`print_json` →
  `dry_run` → send) — `header`, `panel` per theme with prevalence pills, `actions` footer. Card
  grammar → **`lark-im`** skill.

## Tips

1. **Include raw quotes** — Direct participant quotes make insights credible and memorable.
2. **Separate observations from interpretations** — "5 of 8 users clicked the wrong button" is an observation. "The button placement is confusing" is an interpretation.
3. **Quantify where possible** — "Most users" is vague. "7 of 10 users" is specific.
