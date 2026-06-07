# CRM & Sales Plugin

Run a lightweight CRM on Lark Base — pipeline visibility, post-call updates, and re-engagement of dormant clients — without leaving Lark. Primarily designed for [Cowork](https://claude.com/product/cowork), Anthropic's agentic desktop application, though it also works in Claude Code. Drafts only: Claude prepares every customer-facing mail, and you approve every send.

## Installation

```
claude plugins add lark-cowork/crm-sales
```

## What It Does

This plugin treats a Lark Base table as your system-of-record and works the deals around it:

- **Pipeline review** — A weekly scan that rolls deals up by stage, flags stuck deals (no touch in 14d+), surfaces what's closing soon (<30d), and tracks the win-rate trend — so you start Monday knowing what needs attention without reading the whole CRM.
- **Deal update** — Post-call cleanup: pull the meeting minutes, extract pain / budget / timeline / next step, update the Base record (you confirm the diff), then draft a follow-up mail for your approval.
- **Client follow-up** — Detect dormant contacts (>21 days no touch) and draft personalized re-engagement mail that references real history — drafts only, never auto-sent.

## Skills

| Skill | Description |
|-------|-------------|
| `pipeline-review` | Weekly pipeline scan — by stage, stuck deals, closing soon, win-rate trend |
| `deal-update` | Post-call CRM update — pull minutes, extract pain/budget/timeline, update Base record, draft follow-up |
| `client-followup` | Detect dormant contacts (>21d) and draft personalized re-engagement mail (drafts only) |

## Example Workflows

### Weekly Pipeline Review

```
You: pipeline review

Claude: [Rolls deals up by stage with count + $ value]
        [Flags 3 stuck deals (>14d no touch) and 2 closing soon (<30d)]
        [Shows win-rate this month vs last]
        [Recommends top 3 focus actions for the week]
```

### Update After a Call

```
You: just got off the call with Acme, update the deal

Claude: [Pulls the Acme meeting minutes]
        [Extracts: pain (manual reporting), budget ($40k), timeline (Q3)]
        [Shows the Base record diff: Stage → Negotiation, Next step set]
        [On confirm: updates the record, drafts a follow-up mail for review]
```

### Re-engage Dormant Clients

```
You: who have I gone quiet on?

Claude: [Finds 6 contacts with no touch in 21+ days]
        [Per client: last topic, days dormant, suggested angle]
        [Drafts a personalized re-engagement mail for each]
        [Reminds: review each draft in Lark Mail before sending — NOT sent]
```

## Companion plugins

These skills belong to the **lark-cowork** workflow suite and reference skills in sibling
plugins. Skill names resolve globally, so a reference works automatically when the companion
plugin is installed; when a companion is absent the reference **degrades gracefully** (the step
is skipped or offered as a suggestion, never an error). Install the companions below for the full
experience — see [`../connectors/LARK-FUSION.md`](../connectors/LARK-FUSION.md).

| This plugin's skill | References | In plugin |
|---|---|---|
| `pipeline-review` | `morning-brief` (sales variant) | daily-assistant |
| `deal-update` | `contact-360` (pre-call brief) | daily-assistant |

## Data Sources

> If you see unfamiliar placeholders or need to check which tools are connected, see [CONNECTORS.md](CONNECTORS.md).

**Included MCP connection:** the `lark` server (`lark-cli mcp serve`), one bridge covering every category —

- **CRM** is a Lark Base table — read with `lark_base_search`, write via `lark_api` (bitable records) or the `lark-base` skill.
- **Meeting intelligence** (Lark Minutes) feeds deal updates.
- **Email** (Lark Mail) for follow-up and re-engagement — drafts only, you approve every send.
- **Chat, Calendar, Directory** (Lark IM / Calendar / Contact) for context and resolving people.

**Drafts-only posture:** no customer-facing mail is ever sent automatically. Claude prepares drafts in Lark Mail; you review and send.

See [CONNECTORS.md](CONNECTORS.md) for the full category-to-tool mapping and the shared Lark depth core.
