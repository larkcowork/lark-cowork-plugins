# Knowledge & Docs Plugin

A knowledge-and-docs plugin primarily designed for [Cowork](https://claude.com/product/cowork), Anthropic's agentic desktop application — though it also works in Claude Code. Author structured Lark docs from templates and keep your Wiki tidy — Claude fills named templates into clean docs, sheets, and wiki pages, then audits your knowledge base for stale, orphan, and duplicate content.

## Installation

```
claude plugins add lark-cowork/knowledge-docs
```

## What It Does

This plugin helps you produce and maintain structured knowledge in Lark:

- **doc-from-template** — Fill a named Mustache template into a Lark doc, wiki node, or sheet append. Built-in templates: `weekly-report`, `one-on-one`, `meeting-notes`, `project-kickoff`, `postmortem`, `decision-log-row`, `wiki-runbook`. Renders locally and previews before writing.
- **doc-restructure** — Audit a Lark Wiki for stale, orphan, and duplicate pages, then propose archive / merge / re-parent batches. Read-only: it never moves or deletes without your confirmation.

## Skills

| Skill | Description |
|-------|-------------|
| `doc-from-template` | Author a Lark doc/wiki/sheet by filling a named Mustache template, with variable validation and preview-before-write |
| `doc-restructure` | Read-only Wiki audit that flags stale/orphan/duplicate pages and proposes archive/merge/re-parent batches |

## Example Workflows

### Weekly Report Doc

```
You: tạo weekly report doc cho tuần này

Claude: [Picks the weekly-report template, pulls this week's items]
        [Renders locally, shows a preview of the filled doc]
        [On confirm: creates a v2 DocxXML doc and returns URL + token]
```

### Wiki Cleanup

```
You: wiki của team đang bừa, restructure giúp

Claude: [Inventories the space, enriches nodes, runs staleness/orphan/dup rules]
        [Reports: 5000 pages — 120 very stale, 18 orphan owners, 9 duplicate titles]
        [Proposes ARCHIVE / MERGE / RE-PARENT batches with per-batch approve gates]
        [Mutates nothing until you approve a batch]
```

## Companion plugins

These skills belong to the **lark-cowork** workflow suite and reference skills in sibling
plugins. Skill names resolve globally, so a reference works automatically when the companion
plugin is installed; when a companion is absent the reference **degrades gracefully** (the step
is skipped or offered as a suggestion, never an error). Install the companions below for the full
experience — see [`../connectors/LARK-FUSION.md`](../connectors/LARK-FUSION.md).

| This plugin's skill | References | In plugin |
|---|---|---|
| `doc-from-template` (weekly-report) | `weekly-review` | daily-assistant |

## Data Sources

> If you see unfamiliar placeholders or need to check which tools are connected, see [CONNECTORS.md](CONNECTORS.md).

Connect your Lark workspace for the best experience. Without it, draft templates manually.

**Included MCP connection:** the `lark` server (`lark-cli mcp serve`), one bridge covering every category —
- Knowledge base / wiki (Lark Wiki) for reference pages and audits
- Documents (Lark Docs) for authored docs — v2 API, DocxXML by default
- Spreadsheets (Lark Sheets) for row appends like decision logs
- Cloud storage (Lark Drive) for file uploads
- Directory (Lark Contact) for resolving people to `open_id`

Docs are authored with `--api-version v2` and DocxXML, and every write is previewed before it lands.

**Additional options:**
- See [CONNECTORS.md](CONNECTORS.md) for the full category-to-tool map and notes for Claude.
