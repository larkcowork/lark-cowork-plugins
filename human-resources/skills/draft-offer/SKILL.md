---
name: draft-offer
description: Draft an offer letter with comp details and terms. Use when a candidate is ready for an offer, assembling a total comp package (base, equity, signing bonus), writing the offer letter text itself, or prepping negotiation guidance for the hiring manager.
argument-hint: "<role and level>"
---

# /draft-offer

> If you see unfamiliar placeholders or need to check which tools are connected, see [CONNECTORS.md](../../CONNECTORS.md).

Draft a complete offer letter for a new hire.

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md),
> [LARK-RECIPES](../../../connectors/LARK-RECIPES.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md)).
> Pull comp-band context from the `~~HRIS`/`~~compensation data` specialty MCP (or a comp **Lark Base**
> via `lark_base_search`); for a market sanity-check, delegate to the **`/comp-analysis`** skill.
> Land the finished offer letter as a durable **Wiki doc** via `lark_wiki_node_create` (or
> `lark_doc_create`, P8) so HR + the hiring manager share one canonical copy — delegate authoring/edits
> to the **`lark-doc`** skill. Route the offer through a real Lark **approval instance** for the
> compensation sign-off (P7, `lark_api` approval recipe / **`lark-approval`** skill) instead of an
> ad-hoc "ok by me". Track offer status (extended / accepted / declined) in the recruiting **Base**
> (P5, `lark_base_record_upsert`). Send the candidate's offer mail as a **draft** only
> (`lark_mail_draft_create`, P2) — never auto-send. Resolve the hiring manager / candidate with
> `lark_contact_search` first (P1).

## Usage

```
/draft-offer $ARGUMENTS
```

## What I Need From You

- **Role and title**: What position?
- **Level**: Junior, Mid, Senior, Staff, etc.
- **Location**: Where will they be based? (affects comp and benefits)
- **Compensation**: Base salary, equity, signing bonus (if applicable)
- **Start date**: When should they start?
- **Hiring manager**: Who will they report to?

If you don't have all details, I'll help you think through them.

## Output

```markdown
## Offer Letter Draft: [Role] — [Level]

### Compensation Package
| Component | Details |
|-----------|---------|
| **Base Salary** | $[X]/year |
| **Equity** | [X shares/units], [vesting schedule] |
| **Signing Bonus** | $[X] (if applicable) |
| **Target Bonus** | [X]% of base (if applicable) |
| **Total First-Year Comp** | $[X] |

### Terms
- **Start Date**: [Date]
- **Reports To**: [Manager]
- **Location**: [Office / Remote / Hybrid]
- **Employment Type**: [Full-time, Exempt]

### Benefits Summary
[Key benefits highlights relevant to the candidate]

### Offer Letter Text

Dear [Candidate Name],

We are pleased to offer you the position of [Title] at [Company]...

[Complete offer letter text]

### Notes for Hiring Manager
- [Negotiation guidance if needed]
- [Comp band context]
- [Any flags or considerations]
```

## If Connectors Available

If **~~HRIS** (specialty external MCP, or comp **Lark Base**) is connected:
- Pull comp band data for the level/role (`lark_base_search` if approximated in a Base)
- Verify headcount approval
- Auto-populate benefits details

If **~~ATS** (specialty external MCP, or recruiting **Lark Base**) is connected:
- Pull candidate details from the application
- Update offer status in the pipeline → `lark_base_record_upsert` if the pipeline lives in a Base

### Lark-native: approve, store, and deliver

1. **Resolve people (P1):** `lark_contact_search` for the candidate, hiring manager, and approver →
   `open_id`s.
2. **Comp sign-off (P7):** create a real Lark **approval instance** for the offer (POST
   `/open-apis/approval/v4/instances` via `lark_api`, `dry_run` first; confirm shapes with the
   **`lark-approval`** skill). Don't send the offer until it's approved.
3. **Store durably (P8):** create the offer letter as a **Wiki node** (`lark_wiki_node_create`) or a
   doc (`lark_doc_create`); fill/format via the **`lark-doc`** skill.
4. **Track status (P5):** upsert the offer row (extended / accepted / declined, dates) into the
   recruiting Base with `lark_base_record_upsert` (`dry_run` first).
5. **Deliver (P2):** `lark_mail_draft_create` the candidate's offer mail — leave it as a draft for a
   human to review and send. Never auto-send.

## Tips

1. **Include total comp** — Candidates compare total compensation, not just base.
2. **Be specific about equity** — Share count, current valuation method, vesting schedule.
3. **Personalize** — Reference something from the interview process to make it warm.
