---
name: compliance-tracking
description: Track compliance requirements and audit readiness. Trigger with "compliance", "audit prep", "SOC 2", "ISO 27001", "GDPR", "regulatory requirement", or when the user needs help tracking, preparing for, or documenting compliance activities.
---

# Compliance Tracking

Help track compliance requirements, prepare for audits, and maintain regulatory readiness.

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md),
> [LARK-RECIPES](../../../connectors/LARK-RECIPES.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md)).
> A **Lark Base is the control + evidence register** (P5): one table for controls (control↔requirement
> mapping, owner, last-collected, effectiveness), one for evidence, one for the audit calendar. Read
> with `lark_base_search` (requires `search_fields`, no jq — narrow with `select_fields`), write with `lark_base_record_upsert` (`dry_run` first,
> P2). Resolve control owners to `open_id` (P1). Surface the compliance status dashboard and gap list
> as an **interactive card** (P4); land policies/evidence indexes in **Wiki/Drive** (P8).

## Common Frameworks

| Framework | Focus | Key Requirements |
|-----------|-------|-----------------|
| SOC 2 | Service organizations | Security, availability, processing integrity, confidentiality, privacy |
| ISO 27001 | Information security | Risk assessment, security controls, continuous improvement |
| GDPR | Data privacy (EU) | Consent, data rights, breach notification, DPO |
| HIPAA | Healthcare data (US) | PHI protection, access controls, audit trails |
| PCI DSS | Payment card data | Encryption, access control, vulnerability management |

## Compliance Tracking Components

### Control Inventory
- Map controls to framework requirements
- Document control owners and evidence
- Track control effectiveness

### Audit Calendar
- Upcoming audit dates and deadlines
- Evidence collection timelines
- Remediation deadlines

### Evidence Management
- What evidence is needed for each control
- Where evidence is stored
- When evidence was last collected

### Gap Analysis
- Requirements vs. current state
- Prioritized remediation plan
- Timeline to compliance

## Lark-native execution

**Control inventory + evidence register (Base, P5)**:
- **Read** controls/evidence with `lark_base_search`. It REQUIRES `search_fields` and does NOT support `jq` — narrow with `select_fields` (e.g. `["control","requirement","owner","evidence_status","last_collected"]`) + `limit`. If field names are unknown, discover them via `lark_api GET /open-apis/bitable/v1/apps/{base}/tables/{table}/fields`. Use the Base aggregation endpoint to count open gaps rather than fetching all rows (LARK-RECIPES).
- **Write/update** each control's effectiveness, owner, and last-collected date with `lark_base_record_upsert` (`dry_run` first, P2).
- Storing evidence files? Upload to **Drive** with `lark_drive_upload` and put the link in the evidence record.
- No compliance Base yet? Scaffold the control/evidence/audit-calendar tables with the **`base-deploy`** skill — don't hand-roll the schema.

**Audit calendar + remediation (Calendar + Task)**:
- Read upcoming audit dates with `lark_calendar_agenda`; create evidence-collection deadlines as calendar events (`lark_calendar_create`) or tasks (`lark_task_create`, resolve owner via `lark_contact_search` P1, `dry_run` first).
- Track remediation items as Lark Tasks assigned to control owners.

**Status dashboard + gaps (card, P4)**:
- Post the compliance status as an **interactive card** (`lark_im_card_send`, `print_json: true` → `dry_run: true` → send): header with overall readiness pill, `item` rows per open gap with a side button (Assign / Snooze), `note` footer with audit date.

**Policies + evidence index (Wiki, P8)**:
- Publish framework policies, audit-prep checklists, and the evidence index to **Wiki** via `lark_wiki_node_create`; for rich doc bodies delegate to **`lark-doc`**.

## Output

Produce compliance status dashboards (as cards), gap analyses, audit prep checklists, and evidence collection plans — backed by the Base register, surfaced as cards, landed in Wiki.
