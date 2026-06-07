# Approvals & Governance Plugin

A governance plugin primarily designed for [Cowork](https://claude.com/product/cowork), Anthropic's agentic desktop application — though it also works in Claude Code. It keeps Lark operations clean and accountable: faster approval throughput, healthier permission hygiene, and traceable decision provenance — so the things your team agrees, approves, and shares never get lost or quietly drift out of policy.

## Installation

```
claude plugins add lark-cowork/governance
```

## What It Does

This plugin gives Claude a governance lens over your Lark workspace:

- **approval-triage** — Works the pending approval queue item by item, recommending approve / reject / ask-more for each request with a cited policy reference. Agent reads and rationalizes; you decide.
- **approval-flow-sla** — A process-owner view of the whole approval flow: pinpoints the bottleneck node/approver, reports p50/p90 dwell, flags SLA breaches, and audits already-approved items.
- **permission-audit** — A read-only scan across Drive/Doc/Wiki/Base for public, external, and PII exposure, returned as one severity-ranked report. No auto-fix, no auto-revoke.
- **decision-logger** — Detects decisions in IM and Minutes and promotes them into a structured Base table, so "we decided this but nobody remembers where" stops happening.

## Skills

| Skill | Description |
|-------|-------------|
| `approval-triage` | Triage the pending Lark Approval queue — recommend approve/reject/ask-more per item with policy citation |
| `approval-flow-sla` | Measure the approval flow — detect bottleneck node/approver, audit approved data, build & score SLAs |
| `permission-audit` | Read-only scan of Drive/Doc/Wiki/Base for risky permissions (public, external, PII) |
| `decision-logger` | Detect decisions in IM/Minutes and promote them to a structured Base table |

## Example Workflows

### Clearing the Approval Queue

```
You: có gì cần duyệt?

Claude: [Lists 8 pending approvals via the lark-approval skill]
        [Per item: ✅ APPROVE / ⚠️ NEEDS CHECK / ❌ REJECT + policy §]
        [Groups batch-approvable items, flags 2 over threshold]
        "Approve batch of 5? (reply 'approve lô 1')"
```

### Finding the Bottleneck

```
You: nghẽn duyệt ở đâu?

Claude: [Pulls 30d of instances, computes per-node dwell]
        "Node 'Finance review' / Linh ngâm p90 = 18h, breach SLA 6h.
         3 đơn cụ thể đang kẹt. Remind hay transfer?"
```

### Logging a Decision

```
You: chốt cái này — log lại decision

Claude: [Extracts decision from the IM thread, resolves decider + stakeholders]
        [Echoes Title/Decision/Decided by/Date/Source/Area]
        "Confirm Status=Active, Review date 90d? Write to decisions Base?"
```

## Companion plugins

These skills belong to the **lark-cowork** workflow suite and reference skills in sibling
plugins. Skill names resolve globally, so a reference works automatically when the companion
plugin is installed; when a companion is absent the reference **degrades gracefully** (the step
is skipped or offered as a suggestion, never an error). Install the companions below for the full
experience — see [`../connectors/LARK-FUSION.md`](../connectors/LARK-FUSION.md).

| This plugin's skill | References | In plugin |
|---|---|---|
| `decision-logger` | `incident-retro`, `sprint-retro` | delivery-eng |
| `decision-logger` | `meeting-prep`, `morning-brief` | daily-assistant |
| `approval-flow-sla` | `morning-brief` (exec digest) | daily-assistant |

## Data Sources

> If you see unfamiliar placeholders or need to check which tools are connected, see [CONNECTORS.md](CONNECTORS.md).

Connect your approval, storage, and communication tools for the best experience. Without them, governance reviews must be done manually in the Lark UI.

**Included MCP connection:** the `lark` server (`lark-cli mcp serve`), one bridge covering every category —
- Approvals (Lark Approval) via `lark_api` (approval/v4) or the installed `lark-approval` skill — there is no dedicated curated tool
- Cloud storage + knowledge base (Lark Drive / Wiki / Docs) for read-only permission scanning
- Database / records (Lark Base) for the decisions table
- Chat and meeting intelligence (Lark IM + Lark Minutes) for decision detection

**Default posture:** read-only. Remind, transfer, approve, or write only after explicit confirmation.

**Additional options:**
- See [CONNECTORS.md](CONNECTORS.md) for how each category resolves and the `lark_api` escape hatch
