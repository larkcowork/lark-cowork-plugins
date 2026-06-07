---
name: task-management
description: Manage the user's tasks with Lark Task as the system of record. Reference this when the user asks about their tasks, wants to add/complete tasks, extract action items, or get a status snapshot. Reads with lark_task_my, writes with lark_task_create (safe-mutation), surfaces snapshots as interactive cards.
user-invocable: false
---

# Task Management (Lark-native)

**Lark Task is the system of record.** `TASKS.md` / `dashboard.html` in the working directory are
an optional local *mirror* for offline editing — never the source of truth. Always read live state
from Lark before answering, and write changes back to Lark.

> Read [`../../CONNECTORS.md`](../../CONNECTORS.md) and the depth core
> ([LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md), [LARK-RECIPES](../../../connectors/LARK-RECIPES.md),
> [LARK-FUSION](../../../connectors/LARK-FUSION.md)). For anything beyond create/list/complete,
> delegate to the installed **`lark-task`** skill.

## Read — "what's on my plate" / "what's overdue" / "due this week"

1. `lark_task_my` with a `jq` projection (P3 token economy):
   - on your plate now: `{ "complete": false, "jq": ".data.items[] | {id, summary, due, completed}" }`
   - overdue: add `due_end` = now; due this week: `due_start`=now, `due_end`=+7d.
2. Rank by deadline × importance (for a real ranking, delegate to the `task-prioritizer` skill).
3. **Surface as an interactive card, not plain text** (P4): a brief `header`, one `item` row per
   top task with a side button (e.g. *Complete* / *Snooze*), and an `actions` footer. Use
   `lark_im_card_send` with `print_json: true` first to validate. Fall back to plain text only if
   the user explicitly wants it inline.

## Add — "add a task" / "remind me to" / "assign X to <person>"

1. If the task is for someone else, **resolve the assignee first** (P1):
   `lark_contact_search(query="<name>")` → `open_id`.
2. `lark_task_create` with `dry_run: true` (P2) — show summary + assignee + due — then commit on
   confirm. Pass `due` as `date:YYYY-MM-DD` / `relative:+2d`, and `assignee` as the `open_id`.
3. Mirror into local `TASKS.md` only if the user keeps the offline dashboard.

## Complete — "done with X" / "finished X"

1. Find the task id via `lark_task_my` (match on summary).
2. Complete it with `lark_task_complete` (curated, safe-mutation): pass the task id and
   `dry_run: true` first to preview, then commit on confirm.
3. Confirm completion back to the user.

## Extract — from meetings & conversations (P6)

When summarizing a meeting or chat, offer to capture action items into Lark Task:

1. Prefer real Minutes artifacts: `lark_minutes_search(participant_ids="me", query=...)` → pull the
   AI action items instead of re-deriving them. (`im` search for chat-only threads.)
2. For each item: resolve owner (P1) → `lark_task_create(dry_run)` → confirm → commit.
3. Optionally log decisions to a Base (P5) and post a recap **card** (P4).

Always ask before adding — never auto-create tasks without confirmation (P2).

## Conventions

- Bold/clear task summaries; include "for [person]" for commitments and an explicit `due`.
- Keep replies scannable; a card with ≤7 rows beats a long list.
- Tasks assigned to others aren't in `lark_task_my` — use the task list endpoint (LARK-RECIPES).

## Local mirror (optional)

If the user wants the offline dashboard, keep `TASKS.md` + `dashboard.html` in the working dir as a
view, and reconcile on `/productivity:update`: pull `lark_task_my`, diff against `TASKS.md`, and
offer to add/complete the deltas — Lark wins on conflict.
