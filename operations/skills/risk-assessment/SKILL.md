---
name: risk-assessment
description: Identify, assess, and mitigate operational risks. Trigger with "what are the risks", "risk assessment", "risk register", "what could go wrong", or when the user is evaluating risks associated with a project, vendor, process, or decision.
---

# Risk Assessment

Systematically identify, assess, and plan mitigations for operational risks.

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md),
> [LARK-RECIPES](../../../connectors/LARK-RECIPES.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md)).
> The **risk register is a Lark Base** (P5), not a static table: one row per risk with likelihood/impact/
> level/owner/status. Read with `lark_base_search` (requires `search_fields`, no jq — narrow with `select_fields`), write with `lark_base_record_upsert`
> (`dry_run` first, P2). Resolve each risk owner to an `open_id` (`lark_contact_search`, P1). Surface
> the prioritized register as an **interactive card** (P4) and spin each mitigation into a Lark Task.

## Risk Assessment Matrix

| | Low Impact | Medium Impact | High Impact |
|---|-----------|---------------|-------------|
| **High Likelihood** | Medium | High | Critical |
| **Medium Likelihood** | Low | Medium | High |
| **Low Likelihood** | Low | Low | Medium |

## Risk Categories

- **Operational**: Process failures, staffing gaps, system outages
- **Financial**: Budget overruns, vendor cost increases, revenue impact
- **Compliance**: Regulatory violations, audit findings, policy breaches
- **Strategic**: Market changes, competitive threats, technology shifts
- **Reputational**: Customer impact, public perception, partner relationships
- **Security**: Data breaches, access control failures, third-party vulnerabilities

## Risk Register Format

For each risk, document:
- **Description**: What could happen
- **Likelihood**: High / Medium / Low
- **Impact**: High / Medium / Low
- **Risk Level**: Critical / High / Medium / Low
- **Mitigation**: What we're doing to reduce likelihood or impact
- **Owner**: Who is responsible for managing this risk
- **Status**: Open / Mitigated / Accepted / Closed

## Lark-native execution

**Risk register (Base, P5)** — system of record:
- **Read** existing risks with `lark_base_search`. It REQUIRES `search_fields` and does NOT support `jq` — narrow with `select_fields` (e.g. `["risk","likelihood","impact","level","owner","status"]`) + `limit`. If field names are unknown, discover them via `lark_api GET /open-apis/bitable/v1/apps/{base}/tables/{table}/fields`.
- **Write** each new/updated risk with `lark_base_record_upsert` (`base_token`, `table_id`, fields incl. owner `open_id`), `dry_run: true` first (P2).
- No risk register yet? Scaffold it with the **`base-deploy`** skill — don't hand-roll the schema.

**Resolve owners (P1)**:
- `lark_contact_search(query="<name>")` → `open_id` for each risk owner before writing.

**Mitigations as Lark Tasks**:
- For each Critical/High risk, create the mitigation as a Lark Task assigned to the owner: `lark_task_create(dry_run)` → confirm → commit, with a `due`.

**Surface the register (card, P4)**:
- Post the prioritized register as an **interactive card** (`lark_im_card_send`, `print_json: true` first): header pill by highest residual level, `item` rows per Critical/High risk with a side button (Assign owner / Accept risk), `note` footer.

## Output

Produce a prioritized risk register (backed by the Base, surfaced as a card) with specific, actionable mitigations spun into Lark Tasks. Focus on risks that are controllable and material.
