---
name: sprint-planning
description: Plan a sprint — scope work, estimate capacity, set goals, and draft a sprint plan. Use when kicking off a new sprint, sizing a backlog against team availability (accounting for PTO and meetings), deciding what's P0 vs. stretch, or handling carryover from the last sprint.
argument-hint: "[sprint name or date range]"
---

# /sprint-planning

> If you see unfamiliar placeholders or need to check which tools are connected, see [CONNECTORS.md](../../CONNECTORS.md).

Plan a sprint by scoping work, estimating capacity, and setting clear goals.

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md),
> [LARK-RECIPES](../../../connectors/LARK-RECIPES.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md)).
> The backlog/sprint board is a **Lark Base** (P5): read with `lark_base_search` (REQUIRES
> `search_fields`, no `jq` — narrow with `select_fields`/`limit`), write sprint
> assignments + priority with `lark_base_record_upsert` (`dry_run` first, P2). Estimate real capacity
> from the calendar — `lark_calendar_freebusy` / `lark_calendar_agenda` per engineer to subtract
> PTO/meeting load. Resolve every owner to `open_id` via `lark_contact_search` (P1). Sprint backlog
> items become assigned tasks via `lark_task_create`. Share the final plan as an **interactive card**
> (`lark_im_card_send`, P4) and land it in Wiki (`lark_wiki_node_create`, P8). For the sprint
> retro, delegate to the **`sprint-retro`** skill.

## Usage

```
/sprint-planning $ARGUMENTS
```

## How It Works

```
┌─────────────────────────────────────────────────────────────────┐
│                    SPRINT PLANNING                                 │
├─────────────────────────────────────────────────────────────────┤
│  STANDALONE (always works)                                       │
│  ✓ Define sprint goals and success criteria                     │
│  ✓ Estimate team capacity (accounting for PTO, meetings)        │
│  ✓ Scope and prioritize backlog items                           │
│  ✓ Identify dependencies and risks                              │
│  ✓ Generate sprint plan document                                │
├─────────────────────────────────────────────────────────────────┤
│  SUPERCHARGED (Lark-native)                                     │
│  + Base (P5): pull backlog (lark_base_search), write sprint     │
│    + priority (lark_base_record_upsert, dry_run first)          │
│  + Calendar: real capacity via lark_calendar_freebusy/agenda    │
│    per engineer (subtract PTO + meeting load)                   │
│  + Task: assign committed items (lark_task_create)              │
│  + IM: share plan as an interactive card (lark_im_card_send)    │
│  + Wiki: land the sprint plan (lark_wiki_node_create)           │
└─────────────────────────────────────────────────────────────────┘
```

## What I Need From You

- **Team**: Who's on the team and their availability this sprint? (I'll resolve each name to an
  `open_id` via `lark_contact_search` and check `lark_calendar_freebusy` for real PTO/meeting load.)
- **Sprint length**: How many days/weeks?
- **Backlog**: What's prioritized? (I'll pull from the sprint **Lark Base** via `lark_base_search` —
  it REQUIRES `search_fields`; I'll discover field names via
  `lark_api GET /open-apis/bitable/v1/apps/{base}/tables/{table}/fields` if needed, narrow with
  `select_fields`/`limit`, no `jq` — or you can paste/describe.)
- **Carryover**: Anything unfinished from last sprint? (Cross-check `lark_task_my` for open items.)
- **Dependencies**: Anything blocked on other teams?

## Output

```markdown
## Sprint Plan: [Sprint Name]
**Dates:** [Start] — [End] | **Team:** [X] engineers
**Sprint Goal:** [One clear sentence about what success looks like]

### Capacity
| Person | Available Days | Allocation | Notes |
|--------|---------------|------------|-------|
| [Name] | [X] of [Y] | [X] points/hours | [PTO, on-call, etc.] |
| **Total** | **[X]** | **[X] points** | |

### Sprint Backlog
| Priority | Item | Estimate | Owner | Dependencies |
|----------|------|----------|-------|--------------|
| P0 | [Must ship] | [X] pts | [Person] | [None / Blocked by X] |
| P1 | [Should ship] | [X] pts | [Person] | [None] |
| P2 | [Stretch] | [X] pts | [Person] | [None] |

### Planned Capacity: [X] points | Sprint Load: [X] points ([X]% of capacity)

### Risks
| Risk | Impact | Mitigation |
|------|--------|------------|
| [Risk] | [What happens] | [What to do] |

### Definition of Done
- [ ] Code reviewed and merged
- [ ] Tests passing
- [ ] Documentation updated (if applicable)
- [ ] Product sign-off

### Key Dates
| Date | Event |
|------|-------|
| [Date] | Sprint start |
| [Date] | Mid-sprint check-in |
| [Date] | Sprint end / Demo |
| [Date] | Retro |
```

## Commit the plan (Lark-native)

Once the plan is agreed:
1. Write sprint + priority onto each backlog row in the Base: `lark_base_record_upsert`
   (`base_token`, `table_id`, `fields`) — `dry_run: true` first, show the planned writes, then commit (P2).
2. Create assigned tasks for committed items: `lark_task_create` with the owner's `open_id`
   (from `lark_contact_search`) and a `due` matching the sprint end.
3. Post the sprint plan as an **interactive card** to the team channel via `lark_im_card_send`
   (`header` = sprint goal; `item` rows per P0/P1; `actions` footer). Validate with
   `print_json: true`, then `dry_run: true`, then send (P4).
4. Land the durable plan in Wiki via `lark_wiki_node_create` (P8).

## Tips

1. **Leave buffer** — Plan to 70-80% capacity. You will get interrupts. Base real capacity on
   `lark_calendar_freebusy`, not assumed full availability.
2. **One clear sprint goal** — If you can't state it in one sentence, the sprint is unfocused.
3. **Identify stretch items** — Know what to cut if things take longer than expected.
4. **Carry over honestly** — If something didn't ship, understand why before re-committing.
