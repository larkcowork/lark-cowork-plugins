---
name: runbook
description: Create or update an operational runbook for a recurring task or procedure. Use when documenting a task that on-call or ops needs to run repeatably, turning tribal knowledge into exact step-by-step commands, adding troubleshooting and rollback steps to an existing procedure, or writing escalation paths for when things go wrong.
argument-hint: "<process or task name>"
---

# /runbook

> If you see unfamiliar placeholders or need to check which tools are connected, see [CONNECTORS.md](../../CONNECTORS.md).

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md),
> [LARK-RECIPES](../../../connectors/LARK-RECIPES.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md)).
> A runbook is durable on-call knowledge — it lives in the ops **Wiki** (P8), not a chat message.
> Search existing runbooks first to update-not-duplicate (`lark_doc_search`, P3). Resolve escalation
> contacts to real `open_id`s (`lark_contact_search`, P1) so on-call can one-tap message/call them.
> For non-trivial doc bodies delegate to the **`lark-doc`** skill.

Create a step-by-step operational runbook for a recurring task or procedure.

## Usage

```
/runbook $ARGUMENTS
```

## Output

```markdown
## Runbook: [Task Name]
**Owner:** [Team/Person] | **Frequency:** [Daily/Weekly/Monthly/As Needed]
**Last Updated:** [Date] | **Last Run:** [Date]

### Purpose
[What this runbook accomplishes and when to use it]

### Prerequisites
- [ ] [Access or permission needed]
- [ ] [Tool or system required]
- [ ] [Data or input needed]

### Procedure

#### Step 1: [Name]
```
[Exact command, action, or instruction]
```
**Expected result:** [What should happen]
**If it fails:** [What to do]

#### Step 2: [Name]
```
[Exact command, action, or instruction]
```
**Expected result:** [What should happen]
**If it fails:** [What to do]

### Verification
- [ ] [How to confirm the task completed successfully]
- [ ] [What to check]

### Troubleshooting
| Symptom | Likely Cause | Fix |
|---------|-------------|-----|
| [What you see] | [Why] | [What to do] |

### Rollback
[How to undo this if something goes wrong]

### Escalation
| Situation | Contact | Method |
|-----------|---------|--------|
| [When to escalate] | [Who] | [How to reach them] |

### History
| Date | Run By | Notes |
|------|--------|-------|
| [Date] | [Person] | [Any issues or observations] |
```

## Lark-native execution

**Find existing runbooks first (Docs/Wiki, P3)**:
- `lark_doc_search(query="<task name>")` (jq `.data.results[].title`); if one exists, `lark_doc_fetch` it (jq `.data.content[:4000]`) and update rather than rewrite.

**Resolve escalation contacts (P1)**:
- For each escalation entry, `lark_contact_search(query="<name>")` → `open_id`, so the runbook links to a person on-call can directly `lark_im_send` / call. Pull on-call rotation from the ops Base with `lark_base_search` (REQUIRES `search_fields`, no jq — narrow with `select_fields`) if you keep one.

**Publish to the ops Wiki (P8)**:
- Create the page with `lark_wiki_node_create` (`space_id`, `title`); fill the body via the **`lark-doc`** skill (DocxXML) or `lark_doc_create` with `wiki_node`+`wiki_space`.

**Wire it to incidents/changes**:
- Link related change records (the **change-request** skill's Base register) and, when a runbook is actually executed during an incident, hand off the postmortem to the **`incident-retro`** skill.

> **ITSM note:** if a dedicated ITSM MCP is connected, mirror the runbook↔incident-type links there; otherwise the Wiki page + Base on-call register above are the system of record.

## Tips

1. **Be painfully specific** — "Run the script" is not a step. "Run `python sync.py --prod --dry-run` from the ops server" is.
2. **Include failure modes** — What can go wrong at each step and what to do about it.
3. **Test the runbook** — Have someone unfamiliar with the process follow it. Fix where they get stuck.
