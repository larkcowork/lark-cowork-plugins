---
name: close-management
description: Manage the month-end close process with task sequencing, dependencies, and status tracking. Use when planning the close calendar, tracking close progress, identifying blockers, or sequencing close activities by day.
user-invocable: false
---

# Close Management

**Important**: This skill assists with close management workflows but does not provide financial advice. All close activities should be reviewed by qualified financial professionals.

Month-end close checklist, task sequencing and dependencies, status tracking, and common close activities organized by day.

> **Lark-native execution** — read the depth core first: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md),
> [LARK-RECIPES](../../../connectors/LARK-RECIPES.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md),
> and this plugin's [CONNECTORS.md](../../CONNECTORS.md). GL/sub-ledger numbers still come from your
> **ERP / data warehouse** (external MCP) — only the *coordination* layer is Lark:
> - The **close checklist is a Lark Base** (system-of-record, P5): one table = one close cycle, one row
>   per task with Owner / Deadline (T+N) / Status / Blocker / Dependency. Read with `lark_base_search`
>   — it **requires `search_fields`** (the field(s) to match on, e.g. `Status`) and does **not** support
>   `jq`; narrow with `select_fields`/`limit` instead. If you don't know the field names, discover them
>   first via `lark_api GET /open-apis/bitable/v1/apps/{base}/tables/{table}/fields`. Write status with
>   `lark_base_record_upsert` (`dry_run` first, P2). To stand
>   up the Base/table/views the first time, **delegate to `base-deploy`** — don't hand-roll schema.
> - **Assign close tasks** to the right people: resolve each owner with `lark_contact_search` (P1) →
>   `lark_task_create` (assignee = `open_id`, `due` = `date:` the T+N day; `dry_run` first). Mirror task
>   status back into the Base row. Complete with `lark_task_complete`.
> - **Daily close standup**: schedule the recurring 15-min sync via `lark_calendar_create`; pull the day's
>   agenda with `lark_calendar_agenda`. Pull action items from the standup with `lark_minutes_search`
>   (P6) instead of re-typing them.
> - **Status dashboard / blocker alert = an interactive card** (P4), not a wall of text: post via
>   `lark_im_card_send` (`print_json` → `dry_run` → send) to the close channel — colored `header`
>   (green on-track / orange at-risk / red blocked), one `item` row per behind/blocked task with a side
>   button to nudge the owner, `actions` footer. For card grammar delegate to the `lark-im` skill.
> - **Reporting package** (P8): land the final close memo + statements in **Wiki** via
>   `lark_wiki_node_create`; upload supporting workbooks with `lark_drive_upload`. Distribute the link as
>   a card — never leave the package as a chat attachment.

## Month-End Close Checklist

### Pre-Close (Last 2-3 Business Days of the Month)

- [ ] Send close calendar and deadline reminders to all contributors
- [ ] Confirm cut-off procedures with AP, AR, payroll, and treasury
- [ ] Verify all sub-systems are processing normally (ERP, payroll, banking)
- [ ] Complete preliminary bank reconciliation (all but last-day activity)
- [ ] Review open purchase orders for potential accrual needs
- [ ] Confirm payroll processing schedule aligns with close timeline
- [ ] Collect information for any known unusual transactions

### Close Day 1 (T+1: First Business Day After Month-End)

- [ ] Confirm all sub-ledger modules have completed period-end processing
- [ ] Run AP accruals for goods/services received but not invoiced
- [ ] Post payroll entries and payroll accrual (if pay period straddles month-end)
- [ ] Record cash receipts and disbursements through month-end
- [ ] Post intercompany transactions and confirm with counterparties
- [ ] Complete bank reconciliation with final bank statement
- [ ] Run fixed asset depreciation
- [ ] Post prepaid expense amortization

### Close Day 2 (T+2)

- [ ] Complete revenue recognition entries and deferred revenue adjustments
- [ ] Post all remaining accrual journal entries
- [ ] Complete AR subledger reconciliation
- [ ] Complete AP subledger reconciliation
- [ ] Record inventory adjustments (if applicable)
- [ ] Post FX revaluation entries for foreign currency balances
- [ ] Begin balance sheet account reconciliations

### Close Day 3 (T+3)

- [ ] Complete all balance sheet reconciliations
- [ ] Post any adjusting journal entries identified during reconciliation
- [ ] Complete intercompany reconciliation and elimination entries
- [ ] Run preliminary trial balance and income statement
- [ ] Perform preliminary flux analysis on income statement
- [ ] Investigate and resolve material variances

### Close Day 4 (T+4)

- [ ] Post tax provision entries (income tax, sales tax, property tax)
- [ ] Complete equity roll-forward (stock compensation, treasury stock)
- [ ] Finalize all journal entries — soft close
- [ ] Generate draft financial statements (P&L, BS, CF)
- [ ] Perform detailed flux analysis and prepare variance explanations
- [ ] Management review of financial statements and key metrics

### Close Day 5 (T+5)

- [ ] Post any final adjustments from management review
- [ ] Finalize financial statements — hard close
- [ ] Lock the period in the ERP/GL system
- [ ] Distribute financial reporting package to stakeholders
- [ ] Update forecasts/projections based on actual results
- [ ] Conduct close retrospective — identify process improvements

## Task Sequencing and Dependencies

### Dependency Map

Tasks are organized by what must complete before the next task can begin:

```
LEVEL 1 (No dependencies — can start immediately at T+1):
├── Cash receipts/disbursements recording
├── Bank statement retrieval
├── Payroll processing/accrual
├── Fixed asset depreciation run
├── Prepaid amortization
├── AP accrual preparation
└── Intercompany transaction posting

LEVEL 2 (Depends on Level 1 completion):
├── Bank reconciliation (needs: cash entries + bank statement)
├── Revenue recognition (needs: billing/delivery data finalized)
├── AR subledger reconciliation (needs: all revenue/cash entries)
├── AP subledger reconciliation (needs: all AP entries/accruals)
├── FX revaluation (needs: all foreign currency entries posted)
└── Remaining accrual JEs (needs: review of all source data)

LEVEL 3 (Depends on Level 2 completion):
├── All balance sheet reconciliations (needs: all JEs posted)
├── Intercompany reconciliation (needs: both sides posted)
├── Adjusting entries from reconciliations
└── Preliminary trial balance

LEVEL 4 (Depends on Level 3 completion):
├── Tax provision (needs: pre-tax income finalized)
├── Equity roll-forward
├── Consolidation and eliminations
├── Draft financial statements
└── Preliminary flux analysis

LEVEL 5 (Depends on Level 4 completion):
├── Management review
├── Final adjustments
├── Hard close / period lock
├── Financial reporting package
└── Forecast updates
```

### Critical Path

The critical path determines the minimum close duration. Typical critical path:

```
Cash/AP/AR entries → Subledger reconciliations → Balance sheet recs →
  Tax provision → Draft financials → Management review → Hard close
```

To shorten the close:
- Automate Level 1 entries (depreciation, prepaid amortization, standard accruals)
- Pre-reconcile accounts during the month (continuous reconciliation)
- Parallel-process independent reconciliations
- Set clear deadlines with consequences for late submissions
- Use standardized templates to reduce reconciliation prep time

## Status Tracking and Reporting

### Close Status Dashboard

Track each close task with the following attributes — these are the **fields of the close Base table**
(P5). Read live status with `lark_base_search`, update a row with `lark_base_record_upsert`:

| Task | Owner | Deadline | Status | Blocker | Notes |
|------|-------|----------|--------|---------|-------|
| [Task name] | [Person/role] | [Day T+N] | Not Started / In Progress / Complete / Blocked | [If blocked, what's blocking] | [Any notes] |

To surface the board, don't paste the table — post a **status card** via `lark_im_card_send` (P4):
`header` colored by overall state, one `item` row per task that's `Blocked` or `At Risk` with a side
button (*Nudge owner* / *Escalate*), and a `note` footer with the count complete / total.

### Status Definitions

- **Not Started:** Task has not yet begun (may be waiting on dependencies)
- **In Progress:** Task is actively being worked on
- **Complete:** Task is finished and has been reviewed/approved
- **Blocked:** Task cannot proceed due to a dependency, missing data, or issue
- **At Risk:** Task is in progress but may not meet its deadline

### Daily Close Status Meeting (Recommended)

During the close period, hold a brief (15-minute) daily standup (schedule the recurring slot with
`lark_calendar_create`; check who's free via `lark_calendar_freebusy`):

1. **Review status board:** Walk through open tasks, flag any that are behind (`lark_base_search` on the
   close table with `search_fields=["Status"]` — base_search needs the field(s) to match and has no `jq`;
   match the open statuses and narrow returned columns with `select_fields`)
2. **Identify blockers:** Surface any issues preventing task completion
3. **Reassign or escalate:** Adjust ownership (re-assign the `lark_task_create` / update the Base row) or
   escalate blockers; for a hard human gate use a real Lark approval (P7, delegate to `lark-approval`)
4. **Update timeline:** If any tasks are at risk, assess impact on overall close timeline; post the
   revised board as a card (P4). Capture standup action items from `lark_minutes_search` (P6).

### Close Metrics to Track Over Time

| Metric | Definition | Target |
|--------|-----------|--------|
| Close duration | Business days from period end to hard close | Reduce over time |
| # of adjusting entries after soft close | Entries posted during management review | Minimize |
| # of late tasks | Tasks completed after their deadline | Zero |
| # of reconciliation exceptions | Reconciling items requiring investigation | Reduce over time |
| # of restatements / corrections | Errors found after close | Zero |

## Common Close Activities by Day

### Typical 5-Day Close Calendar

| Day | Key Activities | Responsible |
|-----|---------------|-------------|
| **T+1** | Cash entries, payroll, AP accruals, depreciation, prepaid amortization, intercompany posting | Staff accountants, payroll |
| **T+2** | Revenue recognition, remaining accruals, subledger reconciliations (AR, AP, FA), FX revaluation | Revenue accountant, AP/AR, treasury |
| **T+3** | Balance sheet reconciliations, intercompany reconciliation, eliminations, preliminary trial balance, preliminary flux | Accounting team, consolidation |
| **T+4** | Tax provision, equity roll-forward, draft financial statements, detailed flux analysis, management review | Tax, controller, FP&A |
| **T+5** | Final adjustments, hard close, period lock, reporting package distribution, forecast update, retrospective | Controller, FP&A, finance leadership |

### Accelerated Close (3-Day Target)

For organizations targeting a faster close:

| Day | Key Activities |
|-----|---------------|
| **T+1** | All JEs posted (automated + manual), all subledger reconciliations, bank reconciliation, intercompany reconciliation, preliminary trial balance |
| **T+2** | All balance sheet reconciliations, tax provision, consolidation, draft financial statements, flux analysis, management review |
| **T+3** | Final adjustments, hard close, reporting package, forecast update |

**Prerequisites for a 3-day close:**
- Automated recurring journal entries (depreciation, amortization, standard accruals)
- Continuous reconciliation during the month (not all at month-end)
- Automated intercompany elimination
- Pre-close activities completed before month-end (cut-off, accrual estimates)
- Empowered team with clear ownership and minimal handoffs
- Real-time or near-real-time sub-system integration

## Close Process Improvement

### Common Bottlenecks and Solutions

| Bottleneck | Root Cause | Solution |
|-----------|-----------|---------|
| Late AP accruals | Waiting for department spend confirmation | Implement continuous accrual estimation; set cut-off deadlines |
| Manual journal entries | Recurring entries prepared manually each month | Automate standard recurring entries in the ERP |
| Slow reconciliations | Starting from scratch each month | Implement continuous/rolling reconciliation |
| Intercompany delays | Waiting for counterparty confirmation | Automate intercompany matching; set stricter deadlines |
| Management review changes | Large adjustments found during review | Improve preliminary review process; empower team to catch issues earlier |
| Missing supporting documents | Scrambling for documentation at close | Maintain documentation throughout the month |

### Close Retrospective Questions

After each close, ask:
1. What went well this close that we should continue?
2. What took longer than expected and why?
3. What blockers did we encounter and how can we prevent them?
4. Were there any surprises in the financial results we should have caught earlier?
5. What can we automate or streamline for next month?
