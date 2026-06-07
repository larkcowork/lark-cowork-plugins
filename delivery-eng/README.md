# Engineering Delivery Plugin

Run engineering delivery rituals on Lark — incident postmortems and sprint retros. Claude pulls
from your on-call chat, ticket trackers, and velocity history to draft the artifacts your team
needs, blamelessly and from data rather than vibes.

## Installation

```
claude plugins add lark-cowork/delivery-eng
```

## What It Does

This plugin turns raw delivery signal into structured retro artifacts:

- **incident-retro** — A blameless postmortem built from an on-call IM timeline: sequence of
  events, contributing factors, and owner-and-due action items. Reconstructs detection →
  triage → mitigation → resolution without assigning individual blame.
- **sprint-retro** — An end-of-sprint draft built from closed tickets, velocity delta, and
  blocker signal, plus an optional retro form. Surfaces what went well, what didn't, and what
  to try next sprint.

## Skills

| Skill | Description |
|-------|-------------|
| `incident-retro` | Builds a blameless postmortem doc from an on-call IM timeline |
| `sprint-retro` | Drafts an end-of-sprint retro from closed tickets, velocity, and blockers |

## Example Workflows

### Postmortem for a SEV2

```
You: write a postmortem for the SEV2 yesterday

Claude: [Searches the on-call IM channel for the incident window]
        [Reconstructs the timeline: detection → page → triage → mitigation]
        [Cross-references deployed PRs and dashboards]
        [Drafts a blameless postmortem (summary, impact, root cause,
         contributing factors, action items) via the postmortem template]
        [Confirms before saving to the wiki or creating follow-up tasks]
```

### Sprint retro

```
You: sprint retro for this sprint

Claude: [Determines the sprint window from velocity history]
        [Pulls closed tickets from Base/Task and computes velocity delta]
        [Buckets retro-form feedback and groups blocker signal from chat]
        [Drafts a one-page retro: numbers, what went well/didn't,
         try-next, open questions]
        [Offers to save to the wiki and update velocity history]
```

## Companion plugins

These skills belong to the **lark-cowork** workflow suite and reference skills in sibling
plugins. Skill names resolve globally, so a reference works automatically when the companion
plugin is installed; when a companion is absent the reference **degrades gracefully** (the step
is skipped or offered as a suggestion, never an error). Install the companions below for the full
experience — see [`../connectors/LARK-FUSION.md`](../connectors/LARK-FUSION.md).

| This plugin's skill | References | In plugin |
|---|---|---|
| `incident-retro`, `sprint-retro` | `doc-from-template` (postmortem template) | knowledge-docs |

## Data Sources

> If you see unfamiliar placeholders or need to check which tools are connected, see [CONNECTORS.md](CONNECTORS.md).

**Included MCP connection:** the `lark` server (`lark-cli mcp serve`), one bridge covering every category —
- Chat (Lark IM) for on-call timelines and blocker signal
- Project tracker (Lark Task + Lark Base) for closed tickets and velocity
- Documents (Lark Docs + Wiki) for postmortem and retro docs
- Meeting intelligence (Lark Minutes) for incident war-room context
- Directory (Lark Contact) for resolving owners

See [CONNECTORS.md](CONNECTORS.md) for the full tool map and Lark depth references.
