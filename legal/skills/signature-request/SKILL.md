---
name: signature-request
description: Prepare and route a document for e-signature — run a pre-signature checklist, configure signing order, and send for execution. Use when a contract is finalized and ready to sign, when verifying entity names, exhibits, and signature blocks before sending, or when setting up an envelope with sequential or parallel signers.
argument-hint: "<document or contract to send>"
---

# /signature-request -- E-Signature Routing

> If you see unfamiliar placeholders or need to check which tools are connected, see [CONNECTORS.md](../../CONNECTORS.md).

Prepare a document for electronic signature — verify completeness, set signing order, and route for execution.

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md), [LARK-RECIPES](../../../connectors/LARK-RECIPES.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md)).
> - **E-signature itself stays external** (DocuSign etc., per CONNECTORS) — there's no native Lark e-sign tool. Only the *collaboration layer* around it becomes Lark.
> - **Ingest the document**: Lark Doc/Drive → `lark_doc_fetch` / `lark_doc_search`; upload → `lark_drive_upload`. Resolve every signer + CC to `open_id`/email via `lark_contact_search` (P1) — wrong entity/signer is the #1 signing error.
> - **Internal approval before sending** is a real human gate (P7): if anyone must approve pre-send, create a Lark **approval instance** (`lark_api` approval recipe / **`lark-approval`** skill) and only route to e-signature once approved.
> - **Present the pre-signature checklist + signing config as an interactive card** (P4) via `lark_im_card_send` so the requester confirms inline; notify signers via `lark_im_send` / `lark_mail_draft_create` (never auto-send, P2).
> - **Track the envelope in a Lark Base** (P5, `lark_base_record_upsert`: doc, signers, status, sent date) and set a Lark Task / calendar reminder (`lark_task_create` / `lark_calendar_create`) for follow-up if not signed in N days. File the executed copy in **Drive/Wiki** (`lark_drive_upload` / `lark_wiki_node_create`, P8).

**Important**: This command assists with legal workflows but does not provide legal advice. Verify documents are in final form before sending for signature.

## Usage

```
/signature-request $ARGUMENTS
```

Prepare for signature: @$1

## Workflow

### Step 1: Accept the Document

Accept the document in any format:
- **File upload**: PDF, DOCX
- **URL**: Link to a document in ~~cloud storage or ~~CLM
- **Reference**: "The Acme Corp MSA we finalized yesterday"

### Step 2: Pre-Signature Checklist

Before routing for signature, verify:

```markdown
## Pre-Signature Checklist

- [ ] Document is in final, agreed form (no open redlines)
- [ ] All exhibits and schedules are attached
- [ ] Correct legal entity names on signature blocks
- [ ] Dates are correct or left blank for execution date
- [ ] Signature blocks match the authorized signers
- [ ] Any required internal approvals have been obtained
- [ ] Document has been reviewed by appropriate counsel
```

### Step 3: Configure Signing

Gather signing details:
- **Signers**: Who needs to sign? (names, emails, titles)
- **Signing order**: Sequential or parallel?
- **Internal approval**: Does anyone need to approve before the counterparty signs?
- **CC recipients**: Who should receive a copy of the executed document?

### Step 4: Route for Signature

**If ~~e-signature (external DocuSign etc.) is connected:**
- Create the signature envelope/request via that tool's MCP
- Set signing fields and order
- Add any required initials or date fields
- Send for signature, then log the envelope to the tracking Lark Base (P5) and notify signers via Lark (`lark_im_send` / `lark_mail_draft_create`)

**If not connected:**
- Generate a signing instruction document (land it in Drive/Wiki via `lark_drive_upload` / `lark_wiki_node_create`)
- Provide the document formatted for wet signature or manual e-sign
- List all signers with contact information (resolved via `lark_contact_search`)

## Output

```markdown
## Signature Request: [Document Title]

### Document Details
- **Type**: [MSA / NDA / SOW / Amendment / etc.]
- **Parties**: [Party A] and [Party B]
- **Pages**: [X]

### Pre-Signature Check: [PASS / ISSUES FOUND]
[List any issues that need attention before sending]

### Signing Configuration
| Order | Signer | Email | Role |
|-------|--------|-------|------|
| 1 | [Name] | [email] | [Party A Authorized Signatory] |
| 2 | [Name] | [email] | [Party B Authorized Signatory] |

### CC Recipients
- [Name] — [email]

### Status
[Sent for signature / Ready to send / Issues to resolve first]

### Next Steps
- [What to expect after sending]
- [Expected turnaround time]
- [Follow-up if not signed within X days]
```

## Tips

1. **Check entity names carefully** — The most common signing error is incorrect legal entity names.
2. **Verify authority** — Make sure each signer is authorized to bind their organization.
3. **Keep a copy** — File executed copies in Lark Drive (`lark_drive_upload`) or Wiki (`lark_wiki_node_create`, P8) immediately after execution, and update the envelope status in the tracking Base.
