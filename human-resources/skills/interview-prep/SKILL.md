---
name: interview-prep
description: Create structured interview plans with competency-based questions and scorecards. Trigger with "interview plan for", "interview questions for", "how should we interview", "scorecard for", or when the user is preparing to interview candidates.
---

# Interview Prep

Create structured interview plans to evaluate candidates consistently and fairly.

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md),
> [LARK-RECIPES](../../../connectors/LARK-RECIPES.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md)).
> The interview *kit* (question bank + rubric + debrief) is durable knowledge → land it in **Wiki**
> via `lark_wiki_node_create` / `lark_doc_create` (P8; author with the **`lark-doc`** skill) so the
> whole panel works from one copy. Resolve each panelist with `lark_contact_search` (P1), then book
> their interview slots: check overlap with `lark_calendar_freebusy` and create invites with
> `lark_calendar_create` (delegate slot-finding + room booking to the **`lark-calendar`** skill).
> Assign each panelist their competency + scorecard as a **Lark Task** (`lark_task_create`,
> `dry_run` first, P2). Track candidates/scores in the recruiting **Base** (P5,
> `lark_base_record_upsert`). After interviews, drive the debrief from real meeting artifacts via
> `lark_minutes_search` (P6) rather than re-deriving, and post the panel-assignment / debrief
> kickoff as an **interactive card** (`lark_im_card_send`, P4).

## Interview Design Principles

1. **Structured**: Same questions for all candidates in the role
2. **Competency-based**: Map questions to specific skills and behaviors
3. **Evidence-based**: Use behavioral and situational questions
4. **Diverse panel**: Multiple perspectives reduce bias
5. **Scored**: Use rubrics, not gut feelings

## Interview Plan Components

### Role Competencies
Define 4-6 key competencies for the role (e.g., technical skills, communication, leadership, problem-solving).

### Question Bank
For each competency, provide:
- 2-3 behavioral questions ("Tell me about a time...")
- 1-2 situational questions ("How would you handle...")
- Follow-up probes

### Scorecard
Rate each competency on a consistent scale (1-4) with clear descriptions of what each level looks like.

### Debrief Template
Structured format for interviewers to share findings and make a decision.

## Output

Produce a complete interview kit: panel assignment (who interviews for what), question bank by competency, scoring rubric, and debrief template.

### Lark-native delivery

1. **Resolve the panel (P1):** `lark_contact_search` each interviewer → `open_id`s.
2. **Schedule (delegate `lark-calendar`):** `lark_calendar_freebusy` to find overlap →
   `lark_calendar_create` one event per interview round (invite the candidate + panelist, book a
   room). `dry_run` mutations first (P2).
3. **Land the kit (P8):** `lark_wiki_node_create` (or `lark_doc_create`) for the question bank +
   rubric + debrief template; format with the **`lark-doc`** skill. Share the link, not the text.
4. **Assign ownership:** one `lark_task_create` per panelist ("interview X on <competency>, fill
   scorecard"), `dry_run` first.
5. **Track in Base (P5):** create/update the candidate row + per-competency scores via
   `lark_base_record_upsert`; no recruiting Base yet → scaffold with **`base-deploy`**.
6. **Debrief (P6):** pull AI action items/summary from `lark_minutes_search` for each interview, then
   post a debrief **card** (P4) summarizing scores and the hire/no-hire decision.
