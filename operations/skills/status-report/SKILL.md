---
name: status-report
description: Generate a status report with KPIs, risks, and action items. Use when writing a weekly or monthly update for leadership, summarizing project health with green/yellow/red status, surfacing risks and decisions that need stakeholder attention, or turning a pile of project tracker activity into a readable narrative.
argument-hint: "[weekly | monthly | quarterly] [project or team]"
---

# /status-report

> If you see unfamiliar placeholders or need to check which tools are connected, see [CONNECTORS.md](../../CONNECTORS.md).

Generate a polished status report for leadership or stakeholders. See the **risk-assessment** skill for risk matrix frameworks and severity definitions.

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md),
> [LARK-RECIPES](../../../connectors/LARK-RECIPES.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md)).
> Pull metrics + in-progress items from the project **Base** (`lark_base_search`) and Lark Task
> (`lark_task_my` / task-list), decisions from **Minutes** (`lark_minutes_search`, P6) instead of
> re-deriving, and meetings from the calendar — all jq-projected (P3). Deliver the report as an
> **interactive card** to leadership (P4), not a wall of text, and optionally land the durable version
> in **Wiki** (P8). For a recurring weekly view delegate to **`weekly-review`**; for sales pipeline
> status, **`pipeline-review`**.

## Usage

```
/status-report $ARGUMENTS
```

## Output

```markdown
## Status Report: [Project/Team] — [Period]
**Author:** [Name] | **Date:** [Date]

### Executive Summary
[3-4 sentence overview — what's on track, what needs attention, key wins]

### Overall Status: 🟢 On Track / 🟡 At Risk / 🔴 Off Track

### Key Metrics
| Metric | Target | Actual | Trend | Status |
|--------|--------|--------|-------|--------|
| [KPI] | [Target] | [Actual] | [up/down/flat] | 🟢/🟡/🔴 |

### Accomplishments This Period
- [Win 1]
- [Win 2]

### In Progress
| Item | Owner | Status | ETA | Notes |
|------|-------|--------|-----|-------|
| [Item] | [Person] | [Status] | [Date] | [Context] |

### Risks and Issues
| Risk/Issue | Impact | Mitigation | Owner |
|------------|--------|------------|-------|
| [Risk] | [Impact] | [What we're doing] | [Who] |

### Decisions Needed
| Decision | Context | Deadline | Recommended Action |
|----------|---------|----------|--------------------|
| [Decision] | [Why it matters] | [When] | [What I recommend] |

### Next Period Priorities
1. [Priority 1]
2. [Priority 2]
3. [Priority 3]
```

## Lark-native data sources

**Metrics + progress (Base + Task)**:
- KPIs from the project Base: `lark_base_search`. It REQUIRES `search_fields` and does NOT support `jq` — narrow with `select_fields` (e.g. `["metric","target","actual"]`) + `limit`. If field names are unknown, discover them via `lark_api GET /open-apis/bitable/v1/apps/{base}/tables/{table}/fields`. Prefer the Base aggregation endpoint over fetching raw rows to count (LARK-RECIPES).
- Completed/at-risk/overdue items: `lark_task_my` (jq `.data.items[] | {summary, due, completed}`); for teammates' tasks use the task-list endpoint `GET /open-apis/task/v2/tasks` (LARK-RECIPES → Tasks). Flag any `due` in the past as at-risk.

**Decisions + blockers (Minutes + IM, P6)**:
- `lark_minutes_search(participant_ids="me", query=...)` → pull the AI decisions/action items from the period's meetings instead of re-summarizing.
- `lark_im_search` recent team-channel messages (jq-projected) for blockers/decisions not yet in a meeting.

**Meetings (Calendar)**:
- `lark_calendar_agenda` for the key meetings in the reporting window.

**Deliver as an interactive card (P4)**:
- Post with `lark_im_card_send` (`print_json: true` → `dry_run: true` → send): colored `header` (🟢/🟡/🔴 overall), a metrics `panel`, status pills per in-progress item, and an `actions` footer for "Decisions needed". Resolve the channel/recipients first (P1). Offer plain text only if explicitly requested.

**Durable copy (Wiki, P8)**:
- For monthly/quarterly reports leadership will revisit, also publish to **Wiki** (`lark_wiki_node_create`).

## Tips

1. **Lead with the headline** — Busy leaders read the first 3 lines. Make them count.
2. **Be honest about risks** — Surfacing issues early builds trust. Surprises erode it.
3. **Make decisions easy** — For each decision needed, provide context and a recommendation.
