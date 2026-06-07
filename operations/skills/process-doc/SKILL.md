---
name: process-doc
description: Document a business process — flowcharts, RACI, and SOPs. Use when formalizing a process that lives in someone's head, building a RACI to clarify who owns what, writing an SOP for a handoff or audit, or capturing the exceptions and edge cases of how work actually gets done.
argument-hint: "<process name or description>"
---

# /process-doc

> If you see unfamiliar placeholders or need to check which tools are connected, see [CONNECTORS.md](../../CONNECTORS.md).

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md),
> [LARK-RECIPES](../../../connectors/LARK-RECIPES.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md)).
> The finished SOP is durable knowledge — it belongs in **Wiki** (P8), not a chat message. Search
> existing docs first to update-not-duplicate (`lark_doc_search`, jq-projected, P3). Name the real
> people in the RACI by resolving each to an `open_id` (`lark_contact_search`, P1). Spin process-
> improvement items into Lark Tasks (P2 dry_run). For non-trivial doc bodies, delegate authoring to
> the **`lark-doc`** skill.

Document a business process as a complete standard operating procedure (SOP).

## Usage

```
/process-doc $ARGUMENTS
```

## How It Works

Walk me through the process — describe it, paste existing docs, or just tell me the name and I'll ask the right questions. I'll produce a complete SOP.

## Output

```markdown
## Process Document: [Process Name]
**Owner:** [Person/Team] | **Last Updated:** [Date] | **Review Cadence:** [Quarterly/Annually]

### Purpose
[Why this process exists and what it accomplishes]

### Scope
[What's included and excluded]

### RACI Matrix
| Step | Responsible | Accountable | Consulted | Informed |
|------|------------|-------------|-----------|----------|
| [Step] | [Who does it] | [Who owns it] | [Who to ask] | [Who to tell] |

### Process Flow
[ASCII flowchart or step-by-step description]

### Detailed Steps

#### Step 1: [Name]
- **Who**: [Role]
- **When**: [Trigger or timing]
- **How**: [Detailed instructions]
- **Output**: [What this step produces]

#### Step 2: [Name]
[Same format]

### Exceptions and Edge Cases
| Scenario | What to Do |
|----------|-----------|
| [Exception] | [How to handle it] |

### Metrics
| Metric | Target | How to Measure |
|--------|--------|----------------|
| [Metric] | [Target] | [Method] |

### Related Documents
- [Link to related process or policy]
```

## Lark-native execution

**Find existing docs first (Docs/Wiki, P3)**:
- `lark_doc_search(query="<process name>")` (project with `jq: ".data.results[].title"`); if a matching SOP exists, `lark_doc_fetch` it (jq `.data.content[:4000]`) and update rather than duplicate.

**Resolve the RACI (P1)**:
- For each Responsible/Accountable/Consulted/Informed person, `lark_contact_search(query="<name>")` → `open_id`, so the RACI references real directory identities.

**Publish the SOP (Wiki, P8)**:
- Create the page with `lark_wiki_node_create` (`space_id`, `title`); fill the body via the **`lark-doc`** skill (DocxXML) or `lark_doc_create` with `wiki_node`+`wiki_space`. Block JSON by hand is brittle — prefer the skill.

**Improvement action items (Lark Task)**:
- For each "fix this step" item, resolve the owner (P1) → `lark_task_create(dry_run)` → confirm → commit, with a `due`.
- Optionally post the published SOP link as an **interactive card** (`lark_im_card_send`, P4) to the owning team channel.

## Tips

1. **Start messy** — You don't need a perfect description. Tell me how it works today and I'll structure it.
2. **Include the exceptions** — "Usually we do X, but sometimes Y" is the most valuable part to document.
3. **Name the people** — Even if roles change, knowing who does what today helps get the process right.
