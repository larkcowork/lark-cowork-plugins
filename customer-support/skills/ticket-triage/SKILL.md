---
name: ticket-triage
description: Triage and prioritize a support ticket or customer issue. Use when a new ticket comes in and needs categorization, assigning P1-P4 priority, deciding which team should handle it, or checking whether it's a duplicate or known issue before routing.
argument-hint: "<ticket or issue description>"
---

# /ticket-triage

> If you see unfamiliar placeholders or need to check which tools are connected, see [CONNECTORS.md](../../CONNECTORS.md).

Categorize, prioritize, and route an incoming support ticket or customer issue. Produces a structured triage assessment with a suggested initial response.

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md), [LARK-RECIPES](../../../connectors/LARK-RECIPES.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md)).
> The ticket queue lives in a **Lark Base** (P5, system-of-record) — there is no native ticket product, so read with `lark_base_search` and write/route with `lark_base_record_upsert` (`dry_run` first, P2). Keep Intercom (`~~support platform`) for the real ticket thread if its MCP is connected; otherwise the Base *is* the queue. Resolve the assigned agent/customer-success owner via `lark_contact_search` (P1). Surface the triage decision as an **interactive card** (`lark_im_card_send`, P4) with per-ticket Take / Escalate buttons rather than plain text. To scaffold the ticket Base, delegate to **`base-deploy`**; for record/field ops delegate to **`lark-base`**.

## Usage

```
/ticket-triage <ticket text, customer message, or issue description>
```

Examples:
- `/ticket-triage Customer says their dashboard has been showing a blank page since this morning`
- `/ticket-triage "I was charged twice for my subscription this month"`
- `/ticket-triage User can't connect their SSO — getting a 403 error on the callback URL`
- `/ticket-triage Feature request: they want to export reports as PDF`

## Workflow

### 1. Parse the Issue

Read the input and extract:

- **Core problem**: What is the customer actually experiencing?
- **Symptoms**: What specific behavior or error are they seeing?
- **Customer context**: Who is this? Any account details, plan level, or history available?
- **Urgency signals**: Are they blocked? Is this production? How many users affected?
- **Emotional state**: Frustrated, confused, matter-of-fact, escalating?

### 2. Categorize and Prioritize

Using the category taxonomy and priority framework below:

- Assign a **primary category** (bug, how-to, feature request, billing, account, integration, security, data, performance) and an optional secondary category
- Assign a **priority** (P1–P4) based on impact and urgency
- Identify the **product area** the issue maps to

### 3. Check for Duplicates and Known Issues

Before routing, check available sources (project every read with `jq`, P3):

- **~~support platform** (Intercom): search similar open/resolved tickets if its MCP is connected; otherwise search the ticket **Lark Base** — `lark_base_search(base_token, table_id, search_fields:["title"], select_fields:["title","status","customer","priority"])`. base_search REQUIRES `search_fields` and does NOT support `jq` — if field names are unknown, discover them via `lark_api GET /open-apis/bitable/v1/apps/{base}/tables/{table}/fields`; narrow with `select_fields`/`limit`.
- **~~knowledge base** (Lark Wiki + Docs): `lark_doc_search(query="<symptom / exact error>", jq: ".data.results[] | {title, url}")` for known issues / existing articles.
- **~~project tracker** (Lark Task + Base): `lark_task_my(jq: ".data.items[]")` and/or `lark_base_search` (pass `search_fields`; no `jq`) against the bug/feature table for an existing report.
- **~~chat** (Lark IM): `lark_im_search(query="<symptom>", jq: ".data.messages[]")` for teammates who already hit this.

Apply the duplicate detection process below.

### 4. Determine Routing

Using the routing rules below, recommend which team or queue should handle this based on category and complexity.

### 5. Generate Triage Output

```
## Triage: [One-line issue summary]

**Category:** [Primary] / [Secondary if applicable]
**Priority:** [P1-P4] — [Brief justification]
**Product area:** [Area/team]

### Issue Summary
[2-3 sentence summary of what the customer is experiencing]

### Key Details
- **Customer:** [Name/account if known]
- **Impact:** [Who and what is affected]
- **Workaround:** [Available / Not available / Unknown]
- **Related tickets:** [Links to similar issues if found]
- **Known issue:** [Yes — link / No / Checking]

### Routing Recommendation
**Route to:** [Team or queue]
**Why:** [Brief reasoning]

### Suggested Initial Response
[Draft first response to the customer — acknowledge the issue,
set expectations, provide workaround if available.
Use the auto-response templates below as a starting point.]

### Internal Notes
- [Any additional context for the agent picking this up]
- [Reproduction hints if it's a bug]
- [Escalation triggers to watch for]
```

### 6. Record & Route

Write the triage outcome to the ticket **Lark Base** (P5, source of truth) — `lark_base_record_upsert(base_token, table_id, fields: {category, priority, product_area, route_to, status:"Triaged", owner:<open_id>})` with `dry_run: true` first (P2), then commit on confirm. Resolve `owner` via `lark_contact_search` (P1) before writing.

Then **surface the triage as an interactive card** (P4) in the support channel via `lark_im_card_send` — a colored `header` by priority (red P1, orange P2), then elements keyed by KIND: `- md:` / `- item:` rows for category + route + impact, and an `- actions:` footer whose buttons each carry a STRING `value` (or a `url`) — Take / Escalate / Draft reply. Validate with `print_json: true`, then `dry_run: true`, then send. Card YAML grammar lives in the **`lark-im`** skill (`references/spec-reference.md`) — delegate there for anything with branching buttons.

### 7. Offer Next Steps

After presenting the triage:
- "Want me to draft a full response to the customer?" → `/draft-response`
- "Should I search for more context on this issue?" → `/customer-research`
- "Want me to check if this is a known bug in the tracker?" → `lark_base_search` / `lark_task_my`
- "Should I escalate this? I can package it with /customer-escalation."

---

## Category Taxonomy

Assign every ticket a **primary category** and optionally a **secondary category**:

| Category | Description | Signal Words |
|----------|-------------|-------------|
| **Bug** | Product is behaving incorrectly or unexpectedly | Error, broken, crash, not working, unexpected, wrong, failing |
| **How-to** | Customer needs guidance on using the product | How do I, can I, where is, setting up, configure, help with |
| **Feature request** | Customer wants a capability that doesn't exist | Would be great if, wish I could, any plans to, requesting |
| **Billing** | Payment, subscription, invoice, or pricing issues | Charge, invoice, payment, subscription, refund, upgrade, downgrade |
| **Account** | Account access, permissions, settings, or user management | Login, password, access, permission, SSO, locked out, can't sign in |
| **Integration** | Issues connecting to third-party tools or APIs | API, webhook, integration, connect, OAuth, sync, third-party |
| **Security** | Security concerns, data access, or compliance questions | Data breach, unauthorized, compliance, GDPR, SOC 2, vulnerability |
| **Data** | Data quality, migration, import/export issues | Missing data, export, import, migration, incorrect data, duplicates |
| **Performance** | Speed, reliability, or availability issues | Slow, timeout, latency, down, unavailable, degraded |

### Category Determination Tips

- If the customer reports **both** a bug and a feature request, the bug is primary
- If they can't log in due to a bug, category is **Bug** (not Account) — root cause drives the category
- "It used to work and now it doesn't" = **Bug**
- "I want it to work differently" = **Feature request**
- "How do I make it work?" = **How-to**
- When in doubt, lean toward **Bug** — it's better to investigate than dismiss

## Priority Framework

### P1 — Critical
**Criteria:** Production system down, data loss or corruption, security breach, all or most users affected.

- The customer cannot use the product at all
- Data is being lost, corrupted, or exposed
- A security incident is in progress
- The issue is worsening or expanding in scope

**SLA expectation:** Respond within 1 hour. Continuous work until resolved or mitigated. Updates every 1-2 hours.

### P2 — High
**Criteria:** Major feature broken, significant workflow blocked, many users affected, no workaround.

- A core workflow is broken but the product is partially usable
- Multiple users are affected or a key account is impacted
- The issue is blocking time-sensitive work
- No reasonable workaround exists

**SLA expectation:** Respond within 4 hours. Active investigation same day. Updates every 4 hours.

### P3 — Medium
**Criteria:** Feature partially broken, workaround available, single user or small team affected.

- A feature isn't working correctly but a workaround exists
- The issue is inconvenient but not blocking critical work
- A single user or small team is affected
- The customer is not escalating urgently

**SLA expectation:** Respond within 1 business day. Resolution or update within 3 business days.

### P4 — Low
**Criteria:** Minor inconvenience, cosmetic issue, general question, feature request.

- Cosmetic or UI issues that don't affect functionality
- Feature requests and enhancement ideas
- General questions or how-to inquiries
- Issues with simple, documented solutions

**SLA expectation:** Respond within 2 business days. Resolution at normal pace.

### Priority Escalation Triggers

Automatically bump priority up when:
- Customer has been waiting longer than the SLA allows
- Multiple customers report the same issue (pattern detected)
- The customer explicitly escalates or mentions executive involvement
- A workaround that was in place stops working
- The issue expands in scope (more users, more data, new symptoms)

## Routing Rules

Route tickets based on category and complexity:

| Route to | When |
|----------|------|
| **Tier 1 (frontline support)** | How-to questions, known issues with documented solutions, billing inquiries, password resets |
| **Tier 2 (senior support)** | Bugs requiring investigation, complex configuration, integration troubleshooting, account issues |
| **Engineering** | Confirmed bugs needing code fixes, infrastructure issues, performance degradation |
| **Product** | Feature requests with significant demand, design decisions, workflow gaps |
| **Security** | Data access concerns, vulnerability reports, compliance questions |
| **Billing/Finance** | Refund requests, contract disputes, complex billing adjustments |

## Duplicate Detection

Before creating a new ticket or routing, check for duplicates:

1. **Search by symptom**: Look for tickets with similar error messages or descriptions
2. **Search by customer**: Check if this customer has an open ticket for the same issue
3. **Search by product area**: Look for recent tickets in the same feature area
4. **Check known issues**: Compare against documented known issues

**If a duplicate is found:**
- Link the new ticket to the existing one
- Notify the customer that this is a known issue being tracked
- Add any new information from the new report to the existing ticket
- Bump priority if the new report adds urgency (more customers affected, etc.)

## Auto-Response Templates by Category

> Never auto-send. Draft the customer reply via `lark_mail_draft_create` (mail) so a human eyeballs it in the Mail UI first (P2); only send with `lark_mail_send` + `confirm_send: true` after confirmation. For an internal-channel reply use `lark_im_send` (or a card via `lark_im_card_send`). Resolve the recipient's `open_id` with `lark_contact_search` first (P1).

### Bug — Initial Response
```
Thank you for reporting this. I can see how [specific impact]
would be disruptive for your work.

I've logged this as a [priority] issue and our team is
investigating. [If workaround exists: "In the meantime, you
can [workaround]."]

I'll update you within [SLA timeframe] with what we find.
```

### How-to — Initial Response
```
Great question! [Direct answer or link to documentation]

[If more complex: "Let me walk you through the steps:"]
[Steps or guidance]

Let me know if that helps, or if you have any follow-up
questions.
```

### Feature Request — Initial Response
```
Thank you for this suggestion — I can see why [capability]
would be valuable for your workflow.

I've documented this and shared it with our product team.
While I can't commit to a specific timeline, your feedback
directly informs our roadmap priorities.

[If alternative exists: "In the meantime, you might find
[alternative] helpful for achieving something similar."]
```

### Billing — Initial Response
```
I understand billing issues need prompt attention. Let me
look into this for you.

[If straightforward: resolution details]
[If complex: "I'm reviewing your account now and will have
an answer for you within [timeframe]."]
```

### Security — Initial Response
```
Thank you for flagging this — we take security concerns
seriously and are reviewing this immediately.

I've escalated this to our security team for investigation.
We'll follow up with you within [timeframe] with our findings.

[If action is needed: "In the meantime, we recommend
[protective action]."]
```

## Triage Best Practices

1. Read the full ticket before categorizing — context in later messages often changes the assessment
2. Categorize by **root cause**, not just the symptom described
3. When in doubt on priority, err on the side of higher — it's easier to de-escalate than to recover from a missed SLA
4. Always check for duplicates and known issues before routing
5. Write internal notes that help the next person pick up context quickly
6. Include what you've already checked or ruled out to avoid duplicate investigation
7. Flag patterns — if you're seeing the same issue repeatedly, escalate the pattern even if individual tickets are low priority
