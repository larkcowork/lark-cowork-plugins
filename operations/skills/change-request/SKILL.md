---
name: change-request
description: Create a change management request with impact analysis and rollback plan. Use when proposing a system or process change that needs approval, preparing a change record for CAB review, documenting risk and rollback steps before a deployment, or planning stakeholder communications for a rollout.
argument-hint: "<change description>"
---

# /change-request

> If you see unfamiliar placeholders or need to check which tools are connected, see [CONNECTORS.md](../../CONNECTORS.md).

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md),
> [LARK-RECIPES](../../../connectors/LARK-RECIPES.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md)).
> Log the change record in a **Lark Base change register** (P5). Gate approval with a **real Lark
> Approval instance** (P7) — not an ad-hoc "reply yes" — kicked off from an **interactive card** (P4).
> Resolve approvers/owners to `open_id` first (P1), preview every mutation with `dry_run` (P2), and
> land the durable change record in **Wiki** (P8).

Create a structured change request with impact analysis, risk assessment, and rollback plan.

## Usage

```
/change-request $ARGUMENTS
```

## Change Management Framework

Apply the assess-plan-execute-sustain framework when building the request:

### 1. Assess
- What is changing?
- Who is affected?
- How significant is the change? (Low / Medium / High)
- What resistance should we expect?

### 2. Plan
- Communication plan (who, what, when, how)
- Training plan (what skills are needed, how to deliver)
- Support plan (help desk, champions, FAQs)
- Timeline with milestones

### 3. Execute
- Announce and explain the "why"
- Train and support
- Monitor adoption
- Address resistance

### 4. Sustain
- Measure adoption and effectiveness
- Reinforce new behaviors
- Address lingering issues
- Document lessons learned

## Communication Principles

- Explain the **why** before the **what**
- Communicate early and often
- Use multiple channels
- Acknowledge what's being lost, not just what's being gained
- Provide a clear path for questions and concerns

## Output

```markdown
## Change Request: [Title]
**Requester:** [Name] | **Date:** [Date] | **Priority:** [Critical/High/Medium/Low]
**Status:** Draft | Pending Approval | Approved | In Progress | Complete

### Description
[What is changing and why]

### Business Justification
[Why this change is needed — cost savings, compliance, efficiency, risk reduction]

### Impact Analysis
| Area | Impact | Details |
|------|--------|---------|
| Users | [High/Med/Low/None] | [Who is affected and how] |
| Systems | [High/Med/Low/None] | [What systems are affected] |
| Processes | [High/Med/Low/None] | [What workflows change] |
| Cost | [High/Med/Low/None] | [Budget impact] |

### Risk Assessment
| Risk | Likelihood | Impact | Mitigation |
|------|-----------|--------|------------|
| [Risk] | [H/M/L] | [H/M/L] | [How to mitigate] |

### Implementation Plan
| Step | Owner | Timeline | Dependencies |
|------|-------|----------|--------------|
| [Step] | [Person] | [Date] | [What it depends on] |

### Communication Plan
| Audience | Message | Channel | Timing |
|----------|---------|---------|--------|
| [Who] | [What to tell them] | [How] | [When] |

### Rollback Plan
[Step-by-step plan to reverse the change if needed]
- Trigger: [When to roll back]
- Steps: [How to roll back]
- Verification: [How to confirm rollback worked]

### Approvals Required
| Approver | Role | Status |
|----------|------|--------|
| [Name] | [Role] | Pending |
```

## Lark-native execution

**Change register (Base, P5)** — system of record for change records + status:
- **Write** the record with `lark_base_record_upsert` (`base_token`, `table_id`, fields: title, priority, status=Draft, impact, risk, rollback, requester `open_id`), `dry_run: true` first (P2).
- **Read** related/in-flight changes with `lark_base_search` (REQUIRES `search_fields`, no jq — narrow with `select_fields`; discover field names via `lark_api GET /open-apis/bitable/v1/apps/{base}/tables/{table}/fields` if unknown).
- No change register yet? Scaffold one with the **`base-deploy`** skill rather than hand-rolling.

**CAB approval gate (P7)** — a real human gate, not "reply yes":
- Kick off from an **interactive card** (`lark_im_card_send`, P4): header with priority pill, impact/risk rows, and an `actions` row (Approve / Reject / Request changes).
- Create a real **Lark Approval instance** via `lark_api POST /open-apis/approval/v4/instances` (`approval_code`, `form` JSON, approver `open_id`) — see LARK-RECIPES → Approval. The approval engine tracks state and notifies the CAB. For anything beyond a single instance create, delegate to the **`lark-approval`** skill.
- For triaging a queue of pending changes, delegate to **`approval-triage`**; for cycle-time/SLA analysis, **`approval-flow-sla`**.

**Implementation tasks (Lark Task)** — link the plan to executable work:
- For each implementation step, resolve owner (`lark_contact_search`, P1) → `lark_task_create(dry_run)` → confirm → commit. Use `due` (`date:YYYY-MM-DD`) per milestone.

**Stakeholder comms (IM / Mail)** — execute the communication plan:
- Resolve each audience/channel, then `lark_im_send` (or `lark_im_card_send` for a richer announcement; `*_i18n` for bilingual orgs) to the relevant group.
- For formal external notices, `lark_mail_draft_create` (never auto-send; P2) so a human eyeballs it.

**Durable record (Wiki, P8)**:
- Publish the approved change + rollback plan to **Wiki** with `lark_wiki_node_create`.

> **ITSM note:** if a dedicated ServiceNow/Jira ITSM MCP is connected, mirror the ticket there too; otherwise the Lark Base change register above IS the system of record.

## Tips

1. **Be specific about impact** — "Everyone" is not an impact assessment. "200 users in the billing team" is.
2. **Always have a rollback plan** — Even if you're confident, plan for failure.
3. **Communicate early** — Surprises create resistance. Previews create buy-in.
