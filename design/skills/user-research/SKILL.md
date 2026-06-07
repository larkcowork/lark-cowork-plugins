---
name: user-research
description: Plan, conduct, and synthesize user research. Trigger with "user research plan", "interview guide", "usability test", "survey design", "research questions", or when the user needs help with any aspect of understanding their users through research.
---

# User Research

Help plan, execute, and synthesize user research studies.

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md),
> [LARK-RECIPES](../../../connectors/LARK-RECIPES.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md),
> and this plugin's [CONNECTORS.md](../../CONNECTORS.md)). Methods/frameworks below are unchanged; the
> *operational* layer runs on Lark. **Recruit & schedule:** resolve participants with
> `lark_contact_search` (P1), check availability with `lark_calendar_freebusy`, book sessions with
> `lark_calendar_create`, and recruit/confirm via `lark_im_send` or `lark_mail_draft_create` (draft,
> never auto-send, P2). **Capture:** Lark-recorded sessions produce AI transcripts/summaries —
> retrieve via `lark_minutes_search` (P6, delegate depth to **`lark-minutes`**). **Track participants
> & consent in a Lark Base** (P5: `lark_base_search` / `lark_base_record_upsert`; scaffold via
> **base-deploy**). **Land the research plan / interview guide / synthesis in Wiki**
> (`lark_wiki_node_create`, P8, via **`lark-doc`**). For survey design, a `lark_doc_create` form-style
> doc or **Lark Base** form works; for synthesis specifically, hand off to the **research-synthesis**
> skill.

## Research Methods

| Method | Best For | Sample Size | Time |
|--------|----------|-------------|------|
| User interviews | Deep understanding of needs and motivations | 5-8 | 2-4 weeks |
| Usability testing | Evaluating a specific design or flow | 5-8 | 1-2 weeks |
| Surveys | Quantifying attitudes and preferences | 100+ | 1-2 weeks |
| Card sorting | Information architecture decisions | 15-30 | 1 week |
| Diary studies | Understanding behavior over time | 10-15 | 2-8 weeks |
| A/B testing | Comparing specific design choices | Statistical significance | 1-4 weeks |

## Interview Guide Structure

1. **Warm-up** (5 min): Build rapport, explain the session
2. **Context** (10 min): Understand their current workflow
3. **Deep dive** (20 min): Explore the specific topic
4. **Reaction** (10 min): Show concepts or prototypes
5. **Wrap-up** (5 min): Anything we missed? Thank them.

## Analysis Framework

- **Affinity mapping**: Group observations into themes
- **Impact/effort matrix**: Prioritize findings
- **Journey mapping**: Visualize the user experience over time
- **Jobs to be done**: Understand what users are hiring your product to do

## Deliverables (where each lands in Lark)

- **Research plan** (objectives, methods, timeline, participants) → **Wiki** page
  (`lark_wiki_node_create` → **`lark-doc`**, P8); participant roster + status → **Lark Base**
  (`lark_base_record_upsert`, P5).
- **Interview guide** (questions, probes, activities) → **Wiki** doc (P8).
- **Synthesis report** (themes, insights, recommendations) → delegate to the **research-synthesis**
  skill; it publishes to Wiki and tracks Insight→Opportunity rows in a findings Base.
- **Highlight reel** (key quotes and observations) → pull from `lark_minutes_search` AI artifacts
  (P6); attach clips/exports to Drive (`lark_drive_upload`); announce to the team with an
  **interactive card** (`lark_im_card_send`, P4).

## Scheduling & recruiting (concrete)

1. Resolve each participant: `lark_contact_search(query="<name or email>")` → `open_id` (P1).
2. Find a slot: `lark_calendar_freebusy(start, end, user_id)` per participant.
3. Book: `lark_calendar_create` (`dry_run: true` first, P2) — invite the resolved `open_id`s; delegate
   room booking / recurring sessions to the **`lark-calendar`** skill.
4. Recruit/confirm: `lark_im_send` for internal participants, or `lark_mail_draft_create` for external
   (draft only — eyeball in the Mail UI, never auto-send, P2). Delegate to **`lark-mail`**.
