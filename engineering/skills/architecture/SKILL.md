---
name: architecture
description: Create or evaluate an architecture decision record (ADR). Use when choosing between technologies (e.g., Kafka vs SQS), documenting a design decision with trade-offs and consequences, reviewing a system design proposal, or designing a new component from requirements and constraints.
argument-hint: "<decision or system to design>"
---

# /architecture

> If you see unfamiliar placeholders or need to check which tools are connected, see [CONNECTORS.md](../../CONNECTORS.md).

Create an Architecture Decision Record (ADR) or evaluate a system design.

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md),
> [LARK-RECIPES](../../../connectors/LARK-RECIPES.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md)).
> An ADR is durable knowledge — it belongs in **Wiki**, not a chat message (P8). Search prior ADRs
> with `lark_doc_search`; land the new ADR via `lark_wiki_node_create` (then fill with the `lark-doc`
> skill). Resolve every "decider" name to an `open_id` with `lark_contact_search` (P1) before you
> @-mention them or assign action items. Surface the "which option won + who signs off" decision as
> an **interactive card** (`lark_im_card_send`, P4) so deciders can react inline; gate a binding
> sign-off with a real **approval instance** (P7, `lark-approval`) rather than "reply yes". Turn each
> ADR action item into a Lark Task (`lark_task_create`, `dry_run` first — P2). Delegate the design
> reasoning itself to the **system-design** skill.

## Usage

```
/architecture $ARGUMENTS
```

## Modes

**Create an ADR**: "Should we use Kafka or SQS for our event bus?"
**Evaluate a design**: "Review this microservices proposal"
**System design**: "Design the notification system for our app"

See the **system-design** skill for detailed frameworks on requirements gathering, scalability analysis, and trade-off evaluation.

## Output — ADR Format

```markdown
# ADR-[number]: [Title]

**Status:** Proposed | Accepted | Deprecated | Superseded
**Date:** [Date]
**Deciders:** [Who needs to sign off]

## Context
[What is the situation? What forces are at play?]

## Decision
[What is the change we're proposing?]

## Options Considered

### Option A: [Name]
| Dimension | Assessment |
|-----------|------------|
| Complexity | [Low/Med/High] |
| Cost | [Assessment] |
| Scalability | [Assessment] |
| Team familiarity | [Assessment] |

**Pros:** [List]
**Cons:** [List]

### Option B: [Name]
[Same format]

## Trade-off Analysis
[Key trade-offs between options with clear reasoning]

## Consequences
- [What becomes easier]
- [What becomes harder]
- [What we'll need to revisit]

## Action Items
1. [ ] [Implementation step]
2. [ ] [Follow-up]
```

## If Connectors Available

Knowledge base = **Lark Wiki + Docs** (`lark`):
- Search for prior ADRs and design docs with `lark_doc_search` (project with `jq`, P3) so you don't
  re-decide a settled question; `lark_doc_fetch` the closest match for context.
- Publish the accepted ADR as a Wiki page: `lark_wiki_node_create(space_id, title="ADR-NNN: …")`,
  then write the body with the **lark-doc** skill (DocxXML — block JSON by hand is brittle). This is
  the canonical home; chat/email link to it (P8).

Project tracker = **Lark Task + Base** (`lark`):
- Create implementation tasks from the Action Items: resolve each owner via `lark_contact_search`
  (P1) → `lark_task_create(assignee=open_id, due=…, dry_run: true)` → confirm → commit (P2).
- If your team keeps an ADR register in a **Lark Base**, log the record with
  `lark_base_record_upsert` (P5); scaffold the register via **base-deploy** if it doesn't exist.

## Tips

1. **State constraints upfront** — "We need to ship in 2 weeks" or "Must handle 10K rps" shapes the answer.
2. **Name your options** — Even if you're leaning one way, I'll give a more balanced analysis with explicit alternatives.
3. **Include non-functional requirements** — Latency, cost, team expertise, and maintenance burden matter as much as features.
