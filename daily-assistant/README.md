# Daily Assistant Plugin

Your single Lark cockpit for the workday — briefs, triage, and prep in one place. Claude reads your chat, mail, calendar, tasks, and meeting minutes, then hands you ranked, deduplicated, actionable summaries instead of raw feeds. Start the day with a brief, cut through the noise when you're swamped, and walk into every meeting prepared.

## Installation

```
claude plugins add lark-cowork/daily-assistant
```

## What It Does

This plugin covers the full daily workflow, grouped into four themes:

- **Start & end of day** — `morning-brief` opens the day with a merged brief; `daily-digest` wraps it up; `weekly-review` zooms out to the week.
- **Triage** — cut through backlog: `im-digest` for chat, `inbox-zero` for mail, `task-prioritizer` for to-dos, and `overwhelm-triage` to route you to the right one when everything feels like too much.
- **Focus** — protect deep work with `focus-mode`, and reclaim time by auditing your calendar with `calendar-optimizer`.
- **Meetings & people** — `meeting-prep` for context before and action items after, `one-on-one-prep` for 1:1s, and `contact-360` for a full relationship brief before you meet anyone.

## Skills

| Skill | Description |
|-------|-------------|
| `morning-brief` | First-of-day brief orchestrator — fans out mail/IM/approval/task triage in parallel and merges into ≤15 lines (5 variants: default, ic, exec, pm, sales) |
| `daily-digest` | End-of-day digest — today's meetings, completed tasks, and inbox highlights |
| `weekly-review` | Weekly view of calendar, tasks, and OKRs as one narrative |
| `im-digest` | Triages IM groups — classifies the latest messages per group into needs-action / FYI / ignore |
| `inbox-zero` | Full mail-triage workflow targeting an empty inbox |
| `task-prioritizer` | Ranks tasks by deadline × risk, blocking, OKR link, and assigner weight; surfaces today's top 5 |
| `overwhelm-triage` | Meta-router for "I'm overwhelmed" — one question to disambiguate mail/IM/task/meeting pain, then dispatches the right skill |
| `focus-mode` | Blocks calendar time, sets IM do-not-disturb, and notifies your team |
| `calendar-optimizer` | Audits 30-day meeting patterns and surfaces decline / merge / async candidates |
| `meeting-prep` | Two-phase meeting flow — pull context before, extract action items into tasks after |
| `one-on-one-prep` | Brief for an upcoming 1:1 — OKRs, recent tasks, prior notes, and suggested questions |
| `contact-360` | 360° relationship brief before a meeting or call — aggregates IM, mail, meetings, docs, and tasks for one person |

## Example Workflows

### Morning Brief

```
You: morning

Claude: [Fans out mail-triage, im-digest, approval-triage, task-prioritizer in parallel]
        [Merges into a ≤15-line card]
        "3 mails need a reply, 2 approvals pending, top task: budget review (due today).
         You have 4 meetings — first at 10:00 with the Platform team."
```

### Feeling Overwhelmed

```
You: I'm overwhelmed, I don't know where to start

Claude: "Where's the pressure coming from — mail, chat, tasks, or meetings?"
You: tasks
Claude: [Routes to task-prioritizer]
        [Ranks by deadline × risk + blocking + OKR link]
        "Top 5 for today: 1) Ship API spec (blocks 2 people) ..."
```

### Before a Meeting

```
You: I'm about to meet Sarah Chen

Claude: [Resolves Sarah via lark_contact_search, runs contact-360]
        [Aggregates recent IM, mail, shared meetings, docs, open tasks]
        "Last exchange 3 days ago re: Q2 roadmap. 2 shared open tasks.
         She owns the Oracle deal review — due Friday. Suggested openers below."
```

## Companion plugins

These skills belong to the **lark-cowork** workflow suite and reference skills in sibling
plugins. Skill names resolve globally, so a reference works automatically when the companion
plugin is installed; when a companion is absent the reference **degrades gracefully** (the step
is skipped or offered as a suggestion, never an error). Install the companions below for the full
experience — see [`../connectors/LARK-FUSION.md`](../connectors/LARK-FUSION.md).

| This plugin's skill | References | In plugin |
|---|---|---|
| `morning-brief` (exec) | `approval-triage` | governance |
| `morning-brief` (pm) | `decision-logger` | governance |
| `morning-brief` (sales) | `pipeline-review`, `client-followup` | crm-sales |
| `meeting-prep` | `decision-logger` | governance |
| `one-on-one-prep` | `doc-from-template` | knowledge-docs |
| `contact-360` | `deal-update` | crm-sales |

## Data Sources

> If you see unfamiliar placeholders or need to check which tools are connected, see [CONNECTORS.md](CONNECTORS.md).

**Included MCP connection:** the `lark` server (`lark-cli mcp serve`), one bridge covering every category this plugin spans —
- Chat (Lark IM) for group triage and message scanning
- Email (Lark Mail) for inbox triage and highlights
- Calendar (Lark Calendar) for agendas, focus blocks, and meeting audits
- Knowledge base (Lark Wiki + Docs) for reference context
- Project tracker (Lark Task + Lark Base) for task ranking and structured reads
- Meeting intelligence (Lark Minutes + VC) for meeting context and action items
- Directory (Lark Contact) to resolve people before any per-person brief

**Posture:** these skills are read-only by default — Claude produces drafts and suggestions, and only sends mail, creates tasks, blocks calendar, or RSVPs after you confirm. See [CONNECTORS.md](CONNECTORS.md) for the full tool list and notes.
