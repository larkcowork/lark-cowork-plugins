---
name: journal-entry
description: Prepare journal entries with proper debits, credits, and supporting detail. Use when booking month-end accruals (AP, payroll, prepaid), recording depreciation or amortization, posting revenue recognition or deferred revenue adjustments, or documenting an entry for audit review.
argument-hint: "<entry type> [period]"
---

# Journal Entry Preparation

> If you see unfamiliar placeholders or need to check which tools are connected, see [CONNECTORS.md](../../CONNECTORS.md).

**Important**: This command assists with journal entry workflows but does not provide financial advice. All entries should be reviewed by qualified financial professionals before posting.

Prepare journal entries with proper debits, credits, supporting detail, and review documentation.

> **Lark-native execution** — read the depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md),
> [LARK-RECIPES](../../../connectors/LARK-RECIPES.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md),
> [CONNECTORS.md](../../CONNECTORS.md). GL balances / sub-ledger detail come from the **ERP / data
> warehouse** (external MCP) — Lark owns the *review → approve → log* layer:
> - **JE register = a Lark Base** (P5): one row per entry (Type, Period, Total Dr, Total Cr, Preparer,
>   Approver, Status, Reversal date, doc link). Read with `lark_base_search`, write with
>   `lark_base_record_upsert` (`dry_run` first, P2). Scaffold the register via `base-deploy`.
> - **Approval matrix is a real Lark approval** (P7), not "reply yes": route entries over a threshold to
>   the named approver — create an approval instance via `lark_api` (LARK-RECIPES → Approval) or
>   **delegate to `lark-approval`**. Resolve the approver's `open_id` first with `lark_contact_search` (P1).
> - **Surface the entry for review as a card** (P4): `lark_im_card_send` with the Dr/Cr summary and
>   Approve / Send-back buttons; delegate card grammar to `lark-im`.
> - **Land the workpaper** (P8): put the entry memo + calc support in **Wiki** (`lark_wiki_node_create`)
>   or upload the supporting schedule with `lark_drive_upload`; store the link on the Base row.

## Usage

```
/je <type> <period>
```

### Arguments

- `type` — The journal entry type. One of:
  - `ap-accrual` — Accounts payable accruals for goods/services received but not yet invoiced
  - `fixed-assets` — Depreciation and amortization entries for fixed assets and intangibles
  - `prepaid` — Amortization of prepaid expenses (insurance, software, rent, etc.)
  - `payroll` — Payroll accruals including salaries, benefits, taxes, and bonus accruals
  - `revenue` — Revenue recognition entries including deferred revenue adjustments
- `period` — The accounting period (e.g., `2024-12`, `2024-Q4`, `2024`)

## Workflow

### 1. Gather Source Data

If ~~erp or ~~data warehouse is connected:
- Pull the trial balance for the specified period
- Pull subledger detail for the relevant accounts
- Pull prior period entries of the same type for reference
- Identify the current GL balances for affected accounts

If no data source is connected:
> Connect ~~erp or ~~data warehouse to pull GL data automatically. You can also paste trial balance data or upload a spreadsheet.

Prompt the user to provide:
- Trial balance or GL balances for affected accounts
- Subledger detail or supporting schedules
- Prior period entries for reference (optional)

### 2. Calculate the Entry

Based on the JE type:

**AP Accrual:**
- Identify goods/services received but not invoiced by period end
- Calculate accrual amounts from PO receipts, contracts, or estimates
- Debit: Expense accounts (or asset accounts for capitalizable items)
- Credit: Accrued liabilities

**Fixed Assets:**
- Pull the fixed asset register or depreciation schedule
- Calculate period depreciation by asset class and method (straight-line, declining balance, units of production)
- Debit: Depreciation expense (by department/cost center)
- Credit: Accumulated depreciation

**Prepaid:**
- Pull the prepaid amortization schedule
- Calculate the period amortization for each prepaid item
- Debit: Expense accounts (by type — insurance, software, rent, etc.)
- Credit: Prepaid expense accounts

**Payroll:**
- Calculate accrued salaries for days worked but not yet paid
- Calculate accrued benefits (health, retirement contributions, PTO)
- Calculate employer payroll tax accruals
- Calculate bonus accruals (if applicable, based on plan terms)
- Debit: Salary expense, benefits expense, payroll tax expense
- Credit: Accrued payroll, accrued benefits, accrued payroll taxes

**Revenue:**
- Review contracts and performance obligations
- Calculate revenue to recognize based on delivery/performance
- Adjust deferred revenue balances
- Debit: Deferred revenue (or accounts receivable)
- Credit: Revenue accounts (by stream/category)

### 3. Generate the Journal Entry

Present the entry in standard format:

```
Journal Entry: [Type] — [Period]
Prepared by: [User]
Date: [Period end date]

| Line | Account Code | Account Name | Debit | Credit | Department | Memo |
|------|-------------|--------------|-------|--------|------------|------|
| 1    | XXXX        | [Name]       | X,XXX |        | [Dept]     | [Detail] |
| 2    | XXXX        | [Name]       |       | X,XXX  | [Dept]     | [Detail] |
|      |             | **Total**    | X,XXX | X,XXX  |            |      |

Supporting Detail:
- [Calculation basis and assumptions]
- [Reference to supporting schedule or documentation]

Reversal: [Yes/No — if yes, specify reversal date]
```

### 4. Review Checklist

Surface the entry as an **interactive review card** (P4) to the preparer's reviewer — resolve the
reviewer via `lark_contact_search` (P1), post via `lark_im_card_send` with Approve / Send-back actions.
If the amount crosses an approval-matrix threshold, kick off a real Lark **approval instance** (P7,
delegate to `lark-approval`) rather than an ad-hoc chat reply. Before finalizing, verify:

- [ ] Debits equal credits
- [ ] Correct accounting period
- [ ] Account codes are valid and map to the right GL accounts
- [ ] Amounts are calculated correctly with supporting detail
- [ ] Memo/description is clear and specific enough for audit
- [ ] Department/cost center coding is correct
- [ ] Entry is consistent with prior period treatment
- [ ] Reversal flag is set appropriately (accruals should auto-reverse)
- [ ] Supporting documentation is referenced or attached
- [ ] Entry is within the user's approval authority
- [ ] No unusual or out-of-pattern amounts that need investigation

### 5. Output

Provide:
1. The formatted journal entry
2. Supporting calculations
3. Comparison to prior period entry of the same type (if available)
4. Any items flagged for review or follow-up
5. Instructions for posting (manual entry or upload format for the user's ERP)

Then land it in Lark: log the entry as a row in the JE-register Base (`lark_base_record_upsert`,
`dry_run` first), attach the workpaper to Wiki/Drive (P8), and post the review card (P4). The actual
posting still happens in the ERP — Lark holds the audit trail (preparer, approver, reversal, doc link).
