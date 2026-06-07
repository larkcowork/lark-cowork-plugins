---
name: capacity-plan
description: Plan resource capacity — workload analysis and utilization forecasting. Use when heading into quarterly planning, the team feels overallocated and you need the numbers, deciding whether to hire or deprioritize, or stress-testing whether upcoming projects fit the people you have.
argument-hint: "<team or project scope>"
---

# /capacity-plan

> If you see unfamiliar placeholders or need to check which tools are connected, see [CONNECTORS.md](../../CONNECTORS.md).

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md),
> [LARK-RECIPES](../../../connectors/LARK-RECIPES.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md)).
> Treat a **Lark Base as the capacity register** (P5): one row per person with capacity/allocation/skills.
> Resolve every team member to an `open_id` first (`lark_contact_search`, P1). Pull real allocation
> from Lark Task (`lark_task_my` / the task-list endpoint) and real availability from the calendar
> (`lark_calendar_freebusy`, P3 jq projection). Surface the plan as an **interactive card** (P4) for
> stakeholders to act on, and land the durable plan doc in **Wiki** (P8).

Analyze team capacity and plan resource allocation.

## Usage

```
/capacity-plan $ARGUMENTS
```

## What I Need From You

- **Team size and roles**: Who do you have?
- **Current workload**: What are they working on? (Upload from project tracker or describe)
- **Upcoming work**: What's coming next quarter?
- **Constraints**: Budget, hiring timeline, skill requirements

## Planning Dimensions

### People
- Available headcount and skills
- Current allocation and utilization
- Planned hires and timeline
- Contractor and vendor capacity

### Budget
- Operating budget by category
- Project-specific budgets
- Variance tracking
- Forecast vs. actual

### Time
- Project timelines and dependencies
- Critical path analysis
- Buffer and contingency planning
- Deadline management

## Utilization Targets

| Role Type | Target Utilization | Notes |
|-----------|-------------------|-------|
| IC / Specialist | 75-80% | Leave room for reactive work and growth |
| Manager | 60-70% | Management overhead, meetings, 1:1s |
| On-call / Support | 50-60% | Interrupt-driven work is unpredictable |

## Common Pitfalls

- Planning to 100% utilization (no buffer for surprises)
- Ignoring meeting load and context-switching costs
- Not accounting for vacation, holidays, and sick time
- Treating all hours as equal (creative work ≠ admin work)

## Output

```markdown
## Capacity Plan: [Team/Project]
**Period:** [Date range] | **Team Size:** [X]

### Current Utilization
| Person/Role | Capacity | Allocated | Available | Utilization |
|-------------|----------|-----------|-----------|-------------|
| [Name/Role] | [hrs/wk] | [hrs/wk] | [hrs/wk] | [X]% |

### Capacity Summary
- **Total capacity**: [X] hours/week
- **Currently allocated**: [X] hours/week ([X]%)
- **Available**: [X] hours/week ([X]%)
- **Overallocated**: [X people above 100%]

### Upcoming Demand
| Project/Initiative | Start | End | Resources Needed | Gap |
|--------------------|-------|-----|-----------------|-----|
| [Project] | [Date] | [Date] | [X FTEs] | [Covered/Gap] |

### Bottlenecks
- [Skill or role that's oversubscribed]
- [Time period with a crunch]

### Recommendations
1. [Hire / Contract / Reprioritize / Delay]
2. [Specific action]

### Scenarios
| Scenario | Outcome |
|----------|---------|
| Do nothing | [What happens] |
| Hire [X] | [What changes] |
| Deprioritize [Y] | [What frees up] |
```

## Lark-native data sources

**Capacity register (Base, P5)** — the system of record for who-can-do-what:
- **Read** with `lark_base_search`. It REQUIRES `search_fields` (the Bitable API mandates which field(s) to match) and does NOT support `jq` — narrow the payload with `select_fields` (e.g. `["name","role","capacity_hrs","allocated_hrs","skills"]`) + `limit`. If you don't know the field names, discover them first via `lark_api GET /open-apis/bitable/v1/apps/{base}/tables/{table}/fields`.
- **Write** the recomputed allocation back per person with `lark_base_record_upsert` (`base_token`, `table_id`, `fields`), `dry_run: true` first (P2).
- No capacity Base yet? Don't hand-roll the schema — invoke the **`base-deploy`** skill to scaffold a capacity/allocation tracker, then return here.

**Current workload (Lark Task)** — resolve actual allocation:
- `lark_task_my` for the caller; for teammates, the task-list endpoint `GET /open-apis/task/v2/tasks` via `lark_api` (LARK-RECIPES → Tasks; `lark_task_my` only returns the caller's). Project with `jq` to `.data.items[] | {summary, due, assignee}` and count open items per person.

**Real availability (Calendar)** — convert nominal hours to actual:
- Resolve each person (`lark_contact_search` → `open_id`, P1), then `lark_calendar_freebusy(start, end, user_id)` to subtract PTO, holidays, and recurring meeting load from each person's weekly capacity.
- For the whole team in one shot, `lark_api POST /open-apis/calendar/v4/freebusy/list` with `user_id_list` (LARK-RECIPES → Calendar).

**Output (P4 + P8)**:
- Post the capacity summary + bottlenecks + scenarios as an **interactive card** (`lark_im_card_send`, `print_json: true` → `dry_run: true` → send). Use status pills (🟢/🟡/🔴 per person) and an `actions` footer ("Approve plan" / "Open hire req").
- Publish the full plan to **Wiki** with `lark_wiki_node_create` so the org can find it next quarter; for non-trivial doc bodies delegate to the **`lark-doc`** skill.

## Tips

1. **Include all work** — BAU, projects, support, meetings. People aren't 100% available for project work.
2. **Plan for buffer** — Target 80% utilization. 100% means no room for surprises.
3. **Update regularly** — Capacity plans go stale fast. Review monthly.
