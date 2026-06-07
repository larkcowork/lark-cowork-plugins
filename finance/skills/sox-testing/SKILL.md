---
name: sox-testing
description: Generate SOX sample selections, testing workpapers, and control assessments. Use when planning quarterly or annual SOX 404 testing, pulling a sample for a control (revenue, P2P, ITGC, close), building a testing workpaper template, or evaluating and classifying a control deficiency.
argument-hint: "<control area> [period]"
---

# SOX Compliance Testing

> If you see unfamiliar placeholders or need to check which tools are connected, see [CONNECTORS.md](../../CONNECTORS.md).

**Important**: This command assists with SOX compliance workflows but does not provide audit or legal advice. All testing workpapers and assessments should be reviewed by qualified financial professionals before use in audit documentation.

Generate sample selections, create testing workpapers, document control assessments, and provide testing templates for SOX 404 internal controls over financial reporting.

> **Lark-native execution** — read the depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md),
> [LARK-RECIPES](../../../connectors/LARK-RECIPES.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md),
> [CONNECTORS.md](../../CONNECTORS.md). The transaction population to sample comes from the **ERP / data
> warehouse** (external MCP). Lark owns the *control inventory, workpapers, sample log, and deficiency /
> remediation tracking*:
> - **Control matrix + sample log + deficiency register = Lark Base(s)** (P5): controls table (Control #,
>   Type, Frequency, Key, Risk, Assertion, Tester, Reviewer, Conclusion); sample table (Sample #, Ref,
>   Amount, Selection basis, Result, Exception); deficiency table (Severity, Root cause, Owner, Remediation
>   date, Re-test status). Read `lark_base_search` — it **requires `search_fields`** (the field(s) to
>   match, e.g. `["Conclusion"]` or `["Control #"]`) and does **not** support `jq`; narrow with
>   `select_fields`/`limit`. Discover field names first via
>   `lark_api GET /open-apis/bitable/v1/apps/{base}/tables/{table}/fields` if unknown. Write
>   `lark_base_record_upsert` (`dry_run`, P2). Scaffold via `base-deploy` — don't hand-roll the schema.
> - **Workpapers land in Wiki** (P8): `lark_wiki_node_create` per control (one node = one workpaper);
>   upload supporting evidence (screenshots, signed approvals) with `lark_drive_upload` and link from the
>   Base row.
> - **Assign testing + remediation as tasks**: resolve tester / control owner via `lark_contact_search`
>   (P1) → `lark_task_create` (assignee, due). Complete with `lark_task_complete`.
> - **Status / exception summary = an interactive card** (P4, `lark_im_card_send`): header colored by
>   worst conclusion, `item` rows for controls with exceptions + a *Review* button.
> - **Deficiency sign-off = a real Lark approval** (P7, delegate to `lark-approval`), not a chat "ok".

## Usage

```
/sox <control-area> <period>
```

### Arguments

- `control-area` — The control area to test:
  - `revenue-recognition` — Revenue cycle controls (order-to-cash)
  - `procure-to-pay` or `p2p` — Procurement and AP controls (purchase-to-pay)
  - `payroll` — Payroll processing and compensation controls
  - `financial-close` — Period-end close and reporting controls
  - `treasury` — Cash management and treasury controls
  - `fixed-assets` — Capital asset lifecycle controls
  - `inventory` — Inventory valuation and management controls
  - `itgc` — IT general controls (access, change management, operations)
  - `entity-level` — Entity-level and monitoring controls
  - `journal-entries` — Journal entry processing controls
  - Any specific control ID or name
- `period` — The testing period (e.g., `2024-Q4`, `2024`, `2024-H2`)

## Workflow

### 1. Identify Controls to Test

Based on the control area, identify the key controls. Present the control matrix — and persist it as rows
in the controls Base (`lark_base_record_upsert`, P5) so testing status is queryable across the period:

| Control # | Control Description | Type | Frequency | Key/Non-Key | Risk | Assertion |
|-----------|-------------------|------|-----------|-------------|------|-----------|
| [ID]      | [Description]     | Manual/Automated/IT-Dependent | Daily/Weekly/Monthly/Quarterly/Annual | Key | High/Medium/Low | [CEAVOP] |

**Control types:**
- **Automated:** System-enforced controls with no manual intervention
- **Manual:** Controls performed by personnel with judgment
- **IT-dependent manual:** Manual controls that rely on system-generated data

**Assertions (CEAVOP):**
- **C**ompleteness — All transactions are recorded
- **E**xistence/Occurrence — Transactions actually occurred
- **A**ccuracy — Amounts are correctly recorded
- **V**aluation — Assets/liabilities are properly valued
- **O**bligations/Rights — Entity has rights to assets, obligations for liabilities
- **P**resentation/Disclosure — Properly classified and disclosed

### 2. Determine Sample Size

Calculate sample sizes based on control frequency and risk:

| Control Frequency | Population Size (approx.) | Recommended Sample |
|------------------|--------------------------|-------------------|
| Annual           | 1                        | 1 (test the instance) |
| Quarterly        | 4                        | 2 |
| Monthly          | 12                       | 2-4 (based on risk) |
| Weekly           | 52                       | 5-15 (based on risk) |
| Daily            | ~250                     | 20-40 (based on risk) |
| Per-transaction  | Varies                   | 25-60 (based on risk and volume) |

Adjust for:
- **Risk level:** Higher risk controls require larger samples
- **Prior year results:** Controls with prior deficiencies need larger samples
- **Reliance:** Controls relied upon by external auditors may need larger samples

### 3. Generate Sample Selection

Select samples from the population (pulled from the ERP / data warehouse) using the appropriate method,
then log each selected item as a row in the sample Base (`lark_base_record_upsert`):

**Random selection** (default for transaction-level controls):
- Generate random numbers to select specific items from the population
- Ensure coverage across the full period

**Systematic selection** (for periodic controls):
- Select items at fixed intervals with a random start point
- Ensure representation across all sub-periods

**Targeted selection** (supplement to random, for risk-based testing):
- Select items with specific risk characteristics (high dollar, unusual, period-end)
- Document rationale for targeted selections

Present the sample:

```
SAMPLE SELECTION
Control: [Control ID] — [Description]
Period: [Testing period]
Population: [Count] items, $[Total value]
Sample size: [N] items
Selection method: [Random/Systematic/Targeted]

| Sample # | Transaction Date | Reference/ID | Amount | Selection Basis |
|----------|-----------------|--------------|--------|-----------------|
| 1        | [Date]          | [Ref]        | $X,XXX | Random          |
| 2        | [Date]          | [Ref]        | $X,XXX | Random          |
| ...      | ...             | ...          | ...    | ...             |
```

### 4. Create Testing Workpaper

Generate a testing template for each control, then create it as a **Wiki node** (P8,
`lark_wiki_node_create`) so the workpaper is durable and findable; attach evidence via
`lark_drive_upload` and link it from the control's Base row:

```
SOX CONTROL TESTING WORKPAPER
==============================
Control #: [ID]
Control Description: [Full description of the control activity]
Control Owner: [Role/title — to be filled by tester]
Control Type: [Manual/Automated/IT-Dependent Manual]
Frequency: [How often the control operates]
Key Control: [Yes/No]
Relevant Assertion(s): [CEAVOP]
Testing Period: [Period]

TEST OBJECTIVE:
To determine whether [control description] operated effectively throughout the testing period.

TEST PROCEDURES:
1. [Step 1 — What to inspect, examine, or re-perform]
2. [Step 2 — What evidence to obtain]
3. [Step 3 — What to compare or verify]
4. [Step 4 — How to evaluate completeness of performance]
5. [Step 5 — How to assess timeliness of performance]

EXPECTED EVIDENCE:
- [Document type 1 — e.g., signed approval form]
- [Document type 2 — e.g., system screenshot showing review]
- [Document type 3 — e.g., reconciliation with preparer sign-off]

TEST RESULTS:

| Sample # | Ref | Procedure 1 | Procedure 2 | Procedure 3 | Result | Exception? | Notes |
|----------|-----|-------------|-------------|-------------|--------|------------|-------|
| 1        |     | Pass/Fail   | Pass/Fail   | Pass/Fail   | Pass/Fail | Y/N    |       |
| 2        |     | Pass/Fail   | Pass/Fail   | Pass/Fail   | Pass/Fail | Y/N    |       |

EXCEPTIONS NOTED:
| Sample # | Exception Description | Root Cause | Compensating Control | Impact |
|----------|----------------------|------------|---------------------|--------|
|          |                      |            |                     |        |

CONCLUSION:
[ ] Effective — Control operated effectively with no exceptions
[ ] Effective with exceptions — Control operated effectively; exceptions are isolated
[ ] Deficiency — Control did not operate effectively
[ ] Significant Deficiency — Deficiency is more than inconsequential
[ ] Material Weakness — Reasonable possibility of material misstatement not prevented/detected

Tested by: ________________  Date: ________
Reviewed by: _______________  Date: ________
```

### 5. Provide Common Control Templates

Based on the control area, provide pre-built test step templates:

**Revenue Recognition:**
- Verify sales order approval and authorization
- Confirm delivery/performance evidence
- Test revenue recognition timing against contract terms
- Verify pricing accuracy to contract/price list
- Test credit memo approval and validity

**Procure to Pay:**
- Verify purchase order approval and authorization limits
- Confirm three-way match (PO, receipt, invoice)
- Test vendor master data change controls
- Verify payment approval and segregation of duties
- Test duplicate payment prevention controls

**Financial Close:**
- Verify account reconciliation completeness and timeliness
- Test journal entry approval and segregation of duties
- Verify management review of financial statements
- Test consolidation and elimination entries
- Verify disclosure checklist completion

**ITGC:**
- Test user access provisioning and de-provisioning
- Verify privileged access reviews
- Test change management approval and testing
- Verify batch job monitoring and exception handling
- Test backup and recovery procedures

### 6. Document Control Assessment

Classify any identified deficiencies:

**Deficiency:** A control does not allow management or employees to prevent or detect misstatements on a timely basis. Consider:
- Likelihood of misstatement
- Magnitude of potential misstatement
- Whether compensating controls exist

**Significant Deficiency:** A deficiency (or combination) that is less severe than a material weakness but important enough to merit attention by those responsible for oversight.

**Material Weakness:** A deficiency (or combination) such that there is a reasonable possibility that a material misstatement will not be prevented or detected on a timely basis.

### 7. Output

Provide:
1. Control matrix for the selected area
2. Sample selections with methodology documentation
3. Testing workpaper templates with pre-populated test steps
4. Results documentation template
5. Deficiency evaluation framework (if exceptions are identified)
6. Suggested remediation actions for any noted deficiencies

Then land it in Lark: controls/samples/deficiencies as Base rows (`lark_base_record_upsert`), workpapers
in Wiki + evidence in Drive (P8), remediation items as assigned tasks (`lark_task_create`, owner via
`lark_contact_search`), an exception-summary card (P4), and deficiency sign-off via a Lark approval (P7).
