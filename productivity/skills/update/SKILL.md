---
name: update
description: Sync tasks and refresh memory from your current activity. Use when pulling new assignments from your project tracker into TASKS.md, triaging stale or overdue tasks, filling memory gaps for unknown people or projects, or running a comprehensive scan to catch todos buried in chat and email.
argument-hint: "[--comprehensive]"
---

# Update Command

> If you see unfamiliar placeholders or need to check which tools are connected, see [CONNECTORS.md](../../CONNECTORS.md).

Keep your task list and memory current. Two modes:

- **Default:** Sync tasks from external tools, triage stale items, check memory for gaps
- **`--comprehensive`:** Deep scan chat, email, calendar, docs — flag missed todos and suggest new memories

> **Lark-native execution** (see depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md)).
> Lark Task is the source of truth (`TASKS.md` is just a local mirror). Read with `lark_task_my`
> (project via `jq`); complete with `lark_task_complete`; resolve people with `lark_contact_search`
> (P1). Present the sync diff and any "missing tasks" as an **interactive card** (`lark_im_card_send`,
> P4) so the user can act inline. For a richer start-of-day sweep, delegate to the `morning-brief`
> skill; for ranking, `task-prioritizer`.

## Usage

```bash
/productivity:update
/productivity:update --comprehensive
```

## Default Mode

### 1. Load Current State

Read `TASKS.md` and `memory/` directory. If they don't exist, suggest `/productivity:start` first.

### 2. Sync Tasks from External Sources

Check for available task sources:
- **Lark Task** (`lark_task_my`) and **Lark Base** trackers (`lark_base_search`) (if the lark MCP is available)
  - `lark_base_search` **requires `search_fields`** (the field name[s] to match on) and does **not**
    support `jq`. If you don't know the field names, discover them first via
    `lark_api GET /open-apis/bitable/v1/apps/{base}/tables/{table}/fields`, then narrow the result
    with `select_fields` / `limit` instead of a jq projection.
- **GitHub Issues** (if in a repo): `gh issue list --assignee=@me`

If no sources are available, skip to Step 3.

**Fetch tasks assigned to the user** (open/in-progress). Compare against TASKS.md:

| External task | TASKS.md match? | Action |
|---------------|-----------------|--------|
| Found, not in TASKS.md | No match | Offer to add |
| Found, already in TASKS.md | Match by title (fuzzy) | Skip |
| In TASKS.md, not in external | No match | Flag as potentially stale |
| Completed externally | In Active section | Offer to mark done → `lark_task_complete` |

Present the diff as an interactive card (P4) and let the user decide what to add/complete. New tasks
go through `lark_task_create` (resolve assignee via `lark_contact_search`, `dry_run` first).

### 3. Triage Stale Items

Review Active tasks in TASKS.md and flag:
- Tasks with due dates in the past
- Tasks in Active for 30+ days
- Tasks with no context (no person, no project)

Present each for triage: Mark done? Reschedule? Move to Someday?

### 4. Decode Tasks for Memory Gaps

For each task, attempt to decode all entities (people, projects, acronyms, tools, links):

```
Task: "Send PSR to Todd re: Phoenix blockers"

Decode:
- PSR → ✓ Pipeline Status Report (in glossary)
- Todd → ✓ Todd Martinez (in people/)
- Phoenix → ? Not in memory
```

Track what's fully decoded vs. what has gaps.

### 5. Fill Gaps

Present unknown terms grouped:
```
I found terms in your tasks I don't have context for:

1. "Phoenix" (from: "Send PSR to Todd re: Phoenix blockers")
   → What's Phoenix?

2. "Maya" (from: "sync with Maya on API design")
   → Who is Maya?
```

Add answers to the appropriate memory files (people/, projects/, glossary.md).

### 6. Capture Enrichment

Tasks often contain richer context than memory. Extract and update:
- **Links** from tasks → add to project/people files
- **Status changes** ("launch done") → update project status, demote from CLAUDE.md
- **Relationships** ("Todd's sign-off on Maya's proposal") → cross-reference people
- **Deadlines** → add to project files

### 7. Report

```
Update complete:
- Tasks: +3 from Lark Task, 1 completed, 2 triaged
- Memory: 2 gaps filled, 1 project enriched
- All tasks decoded ✓
```

## Comprehensive Mode (`--comprehensive`)

Everything in Default Mode, plus a deep scan of recent activity.

### Extra Step: Scan Activity Sources

Gather data from the `lark` MCP (project each with `jq` to control token cost, P3):
- **Chat:** `lark_im_search` recent messages for commitments/action items
- **Email:** mail search via `lark_api` (no curated mail-search tool; the curated mail tools are
  `lark_mail_send` / `lark_mail_draft_create`) for sent promises
- **Documents:** `lark_doc_search` recently touched docs
- **Calendar:** `lark_calendar_agenda` recent + upcoming events
- **Meetings:** `lark_minutes_search` (participant_ids="me") → pull AI action items directly (P6)

### Extra Step: Flag Missed Todos

Compare activity against TASKS.md. Surface action items that aren't tracked:

```
## Possible Missing Tasks

From your activity, these look like todos you haven't captured:

1. From chat (Jan 18):
   "I'll send the updated mockups by Friday"
   → Add to TASKS.md?

2. From meeting "Phoenix Standup" (Jan 17):
   You have a recurring meeting but no Phoenix tasks active
   → Anything needed here?

3. From email (Jan 16):
   "I'll review the API spec this week"
   → Add to TASKS.md?
```

Let user pick which to add.

### Extra Step: Suggest New Memories

Surface new entities not in memory:

```
## New People (not in memory)
| Name | Frequency | Context |
|------|-----------|---------|
| Maya Rodriguez | 12 mentions | design, UI reviews |
| Alex K | 8 mentions | DMs about API |

## New Projects/Topics
| Name | Frequency | Context |
|------|-----------|---------|
| Starlight | 15 mentions | planning docs, product |

## Suggested Cleanup
- **Horizon project** — No mentions in 30 days. Mark completed?
```

Present grouped by confidence. High-confidence items offered to add directly; low-confidence items asked about.

## Notes

- Never auto-add tasks or memories without user confirmation
- External source links are preserved when available
- Fuzzy matching on task titles handles minor wording differences
- Safe to run frequently — only updates when there's new info
- `--comprehensive` always runs interactively
