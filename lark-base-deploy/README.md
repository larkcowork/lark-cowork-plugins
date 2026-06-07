# Lark Base Deploy Plugin

A Lark Base deployment plugin primarily designed for [Cowork](https://claude.com/product/cowork), Anthropic's agentic desktop application — though it also works in Claude Code. It turns one sentence — "build a Base for my 30-person team: 3 tables + dashboard + automation + Excel import" — into a working, verified, handed-over Base. An 8-phase orchestrator runs Discovery → Handover, designing the schema, building it on Lark, wiring relations and roles, importing data, standing up dashboards and automation, then acceptance-testing and documenting the result for the team.

## Installation

```
claude plugins add lark-cowork/lark-base-deploy
```

## What It Does

This plugin orchestrates an end-to-end Base deployment through eight phases:

- **Phase 0 — Discovery** — interview/confirm operational requirements into a spec (entities, KPIs, roles, data sources).
- **Phase 1 — Design** — design the full schema: tables, fields, views, relations, rules, options, roles.
- **Phase 2 — Build** ⚡ — create the Base, tables, fields, and views; **fan-out 6 sub-agents** (table, field, view, relation, rule, option) working in parallel; write real IDs back into the plan.
- **Phase 3 — Wire-up** — connect tables (link/lookup/formula), set view filters, build roles and access permissions.
- **Phase 4 — Import** — import Excel/CSV data into the Base, cleaning and mapping fields to the designed schema.
- **Phase 5 — Viz** ⚡ — build the KPI dashboard with charts and filters; **fan-out 3 sub-agents** to compose `data_config` in parallel, then create blocks sequentially.
- **Phase 6 — Automation** — set up Base workflow automation plus a bot for reminders and notifications.
- **Phase 7 — Handover** — acceptance-test the UI for real (never trust `code:0`), write a usage guide, and hand the Base over to the team.

Everything flows through **`build-plan.json` — the single source of truth (SSOT)**. Every phase reads and writes it, and only real API-returned identifiers (`base_token`, `table_id`, …) are written back.

## Skills

| Skill | Description |
|-------|-------------|
| `base-deploy` | Orchestrator — deploy a Lark Base end-to-end for a team across 8 phases (Discovery → Handover), spawning parallel sub-agents at Build and Viz. |
| `base-discovery` | Phase 0 — interview/confirm requirements into a spec (entity, KPI, role, data source) written to `build-plan.json`. |
| `base-design` | Phase 1 — design the schema (tables, fields, views, relations, rules, options, roles) into a complete `build-plan.json` for Build to execute. |
| `base-build` | Phase 2 — create Base + tables + fields + views from `build-plan.json`, fan-out 1 sub-agent per table in parallel, write real IDs back into the plan. |
| `base-wireup` | Phase 3 — connect tables (link/lookup/formula), set view filters, build roles and access permissions from `build-plan.json`. |
| `base-import` | Phase 4 — import Excel/CSV into the Base, clean and map fields normalized to `build-plan.json`. |
| `base-viz` | Phase 5 — build the KPI dashboard (charts, filters), fan-out 3 sub-agents to compose `data_config` in parallel, then create blocks sequentially. |
| `base-automation` | Phase 6 — build Base workflow automation plus a reminder/notification bot per `build-plan.json`. |
| `base-handover` | Phase 7 — acceptance-test the UI (verify for real, don't trust `code:0`), write the usage guide, and hand the Base over to the team. |

## Example Workflows

### Full Deploy from One Sentence

```
You: Build a Base for my 30-person team: 3 tables (Projects, Tasks, People),
     a KPI dashboard, automation to remind owners of due tasks, and import
     our existing roster from team.xlsx

Claude: [Phase 0 Discovery — confirms entities, KPIs, owner, data source → build-plan.json]
        [Phase 1 Design — drafts tables/fields/views/relations/roles into the plan]
        [Phase 2 Build ⚡ — creates Base, fans out 6 sub-agents, fills real IDs]
        [Phase 3 Wire-up — links Tasks→Projects, lookups, view filters, roles]
        [Phase 4 Import — maps team.xlsx into People, cleans + normalizes]
        [Phase 5 Viz ⚡ — 3 sub-agents compose data_config, builds dashboard blocks]
        [Phase 6 Automation — due-task reminder workflow + bot notifications]
        [Phase 7 Handover — visually verifies the UI, writes usage guide, hands over]
```

### Dashboard-Only on an Existing Base

```
You: Add a KPI dashboard to my existing Base (base_token bascn...) —
     deal count by stage, win-rate trend, and a stage filter

Claude: [Loads/derives build-plan.json from the existing Base]
        [Runs Phase 5 Viz ⚡ — 3 sub-agents compose data_config in parallel]
        [Creates dashboard blocks sequentially, verifies them in the UI]
```

### Dry-Run a Plan Before Building

```
You: /base-deploy --dry-run a sprint tracker: Sprints, Stories, Bugs

Claude: [Runs Discovery + Design only — produces a complete build-plan.json]
        [Shows the proposed tables, fields, relations, roles, dashboard, automation]
        [Stops before any write — "Approve to run Build → Handover?"]
```

## Companion plugins

These skills belong to the **lark-cowork** workflow suite and reference skills in sibling
plugins. Skill names resolve globally, so a reference works automatically when the companion
plugin is installed; when a companion is absent the reference **degrades gracefully** (the step
is skipped or offered as a suggestion, never an error). Install the companions below for the full
experience — see [`../connectors/LARK-FUSION.md`](../connectors/LARK-FUSION.md).

| This plugin's skill | References | In plugin |
|---|---|---|
| `base-handover` | `permission-audit` | governance |

## Data Sources

> If you see unfamiliar placeholders or need to check which tools are connected, see [CONNECTORS.md](CONNECTORS.md).

**Included MCP connection:** the `lark` server (`lark-cli mcp serve`), one bridge backing every category —
- Project tracker (Lark Base + Task) — the system of record this plugin builds and writes to
- Office suite (Lark Sheets/Docs/Drive) — for Excel/CSV import and generated handover docs
- Chat (Lark IM) — for automation reminders and bot notifications
- Directory (Lark Contact) — to resolve people into `open_id` before assigning fields/roles

Base structure and record writes (create base/table/field/view, bitable records, dashboard blocks, workflow automation) have no dedicated curated tool — they go through `lark_api` (`bitable/v1`, `base/v2`) or the installed `lark-base` skill. See [CONNECTORS.md](CONNECTORS.md) for the full category-to-tool mapping.
