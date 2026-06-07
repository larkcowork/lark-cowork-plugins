---
name: vendor-review
description: Evaluate a vendor — cost analysis, risk assessment, and recommendation. Use when reviewing a new vendor proposal, deciding whether to renew or replace a contract, comparing two vendors side-by-side, or building a TCO breakdown and negotiation points before procurement sign-off.
argument-hint: "<vendor name or proposal>"
---

# /vendor-review

> If you see unfamiliar placeholders or need to check which tools are connected, see [CONNECTORS.md](../../CONNECTORS.md).

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md),
> [LARK-RECIPES](../../../connectors/LARK-RECIPES.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md)).
> Keep the **vendor register in a Lark Base** (P5): spend, renewal date, SLA, risk, status per vendor.
> Read with `lark_base_search` (requires `search_fields`, no jq — narrow with `select_fields`), write the review verdict back with `lark_base_record_upsert`
> (`dry_run` first, P2). Upload the proposal PDF to **Drive** (`lark_drive_upload`) and link it. Surface
> the recommendation as an **interactive card** (P4), and gate procurement sign-off with a real Lark
> **Approval instance** (P7).

Evaluate a vendor with structured analysis covering cost, risk, performance, and fit.

## Usage

```
/vendor-review $ARGUMENTS
```

## What I Need From You

- **Vendor name**: Who are you evaluating?
- **Context**: New vendor evaluation, renewal decision, or comparison?
- **Details**: Contract terms, pricing, proposal document, or current performance data

## Evaluation Framework

### Cost Analysis (Total Cost of Ownership)
- Total cost of ownership (not just license fees)
- Implementation and migration costs
- Training and onboarding costs
- Ongoing support and maintenance
- Exit costs (data migration, contract termination)

### Risk Assessment
- Vendor financial stability
- Security and compliance posture
- Concentration risk (single vendor dependency)
- Contract lock-in and exit terms
- Business continuity and disaster recovery

### Performance Metrics
- SLA compliance
- Support response times
- Uptime and reliability
- Feature delivery cadence
- Customer satisfaction

### Comparison Matrix
When comparing vendors, produce a side-by-side matrix covering: pricing, features, integrations, security, support, contract terms, and references.

## Output

```markdown
## Vendor Review: [Vendor Name]
**Date:** [Date] | **Type:** [New / Renewal / Comparison]

### Summary
[2-3 sentence recommendation]

### Cost Analysis
| Component | Annual Cost | Notes |
|-----------|-------------|-------|
| License/subscription | $[X] | [Per seat, flat, usage-based] |
| Implementation | $[X] | [One-time] |
| Support/maintenance | $[X] | [Included or add-on] |
| **Total Year 1** | **$[X]** | |
| **Total 3-Year** | **$[X]** | |

### Risk Assessment
| Risk | Likelihood | Impact | Mitigation |
|------|-----------|--------|------------|
| [Risk] | High/Med/Low | High/Med/Low | [Mitigation] |

### Strengths
- [Strength 1]
- [Strength 2]

### Concerns
- [Concern 1]
- [Concern 2]

### Recommendation
[Proceed / Negotiate / Pass] — [Reasoning]

### Negotiation Points
- [Leverage point 1]
- [Leverage point 2]
```

## Lark-native execution

**Vendor register (Base, P5)** — system of record for vendors, spend, renewals, SLA:
- **Read** the existing vendor record + spend history with `lark_base_search`. It REQUIRES `search_fields` and does NOT support `jq` — narrow with `select_fields` (e.g. `["vendor","annual_spend","renewal_date","sla","status"]`) + `limit`. If field names are unknown, discover them via `lark_api GET /open-apis/bitable/v1/apps/{base}/tables/{table}/fields`. For renewals this is your "what we pay now" baseline. Compare pricing against other vendor rows in the same Base.
- **Write** the review verdict (recommendation, TCO, renewal decision, status) with `lark_base_record_upsert` (`dry_run` first, P2).
- No vendor register yet? Scaffold it with the **`base-deploy`** skill — don't hand-roll the schema.

**Proposal + prior reviews (Drive + Docs/Wiki)**:
- Upload the proposal/contract PDF to **Drive** (`lark_drive_upload`) and link it in the vendor record.
- Search prior evaluations and procurement policy with `lark_doc_search` → `lark_doc_fetch` (jq-projected, P3); publish the finished review to **Wiki** (`lark_wiki_node_create`, P8).

**Recommendation (card, P4)**:
- Post the verdict as an **interactive card** (`lark_im_card_send`, `print_json: true` first): header pill (Proceed / Negotiate / Pass), cost + risk rows, `actions` footer (Approve / Request changes).

**Procurement sign-off (Approval, P7)**:
- When spend exceeds the approval threshold, kick off a real Lark **Approval instance** (`lark_api POST /open-apis/approval/v4/instances`, LARK-RECIPES → Approval; resolve approver `open_id` first, P1) rather than an ad-hoc yes. Delegate complex approval flows to the **`lark-approval`** skill.

> **Procurement note:** if a dedicated procurement MCP is connected, pull live contract terms/spend from it; otherwise the Lark Base vendor register above is the system of record.

## Tips

1. **Upload the proposal** — I can extract pricing, terms, and SLAs from vendor documents.
2. **Compare vendors** — "Compare Vendor A vs Vendor B" gets you a side-by-side analysis.
3. **Include current spend** — For renewals, knowing what you pay now helps evaluate price changes.
