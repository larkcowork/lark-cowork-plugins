---
name: documentation
description: Write and maintain technical documentation. Trigger with "write docs for", "document this", "create a README", "write a runbook", "onboarding guide", or when the user needs help with any form of technical writing — API docs, architecture docs, or operational runbooks.
---

# Technical Documentation

Write clear, maintainable technical documentation for different audiences and purposes.

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md),
> [LARK-RECIPES](../../../connectors/LARK-RECIPES.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md)).
> Durable docs (READMEs, runbooks, API refs, onboarding guides, architecture docs) belong in **Wiki**,
> not chat (P8): before writing, `lark_doc_search` for an existing page to update instead of
> duplicating ("link, don't duplicate"); create the page with `lark_wiki_node_create(space_id, title)`
> and fill the body with the **lark-doc** skill (DocxXML — block JSON by hand is brittle). Transient
> artifacts (generated diagrams, exports, attachments) go to **Drive** via `lark_drive_upload`. When a
> doc names owners/escalation contacts, resolve them to `open_id` with `lark_contact_search` (P1) so
> @-mentions resolve. Announce a new/updated doc with a short `lark_im_send` link to the channel — but
> the doc lives in Wiki, the message is just a pointer. For a wiki-wide tidy-up delegate to
> **doc-restructure**; to fill a named template, **doc-from-template**.

## Document Types

### README
- What this is and why it exists
- Quick start (< 5 minutes to first success)
- Configuration and usage
- Contributing guide

### API Documentation
- Endpoint reference with request/response examples
- Authentication and error codes
- Rate limits and pagination
- SDK examples

### Runbook
- When to use this runbook
- Prerequisites and access needed
- Step-by-step procedure
- Rollback steps
- Escalation path

### Architecture Doc
- Context and goals
- High-level design with diagrams
- Key decisions and trade-offs
- Data flow and integration points

### Onboarding Guide
- Environment setup
- Key systems and how they connect
- Common tasks with walkthroughs
- Who to ask for what

## Principles

1. **Write for the reader** — Who is reading this and what do they need?
2. **Start with the most useful information** — Don't bury the lede
3. **Show, don't tell** — Code examples, commands, screenshots
4. **Keep it current** — Outdated docs are worse than no docs
5. **Link, don't duplicate** — Reference other docs instead of copying
