---
name: invoice-chase
version: 0.2.0
description: >
  Drafts overdue-invoice reminder emails from QuickBooks and PayPal data,
  matched to each customer's payment history and tone (gentle for good customers,
  firm for repeat late payers). Sends via PayPal with owner approval;
  non-PayPal invoices queue as mail drafts. Use when the user asks
  "who owes me money," mentions overdue invoices, or wants to follow up
  on unpaid invoices.
---

# Invoice Chase

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md)). AR/payment data stays on QuickBooks / PayPal / Stripe (external). The **reminder layer is Lark**: non-PayPal reminders are staged with `lark_mail_draft_create` (never auto-send, P2); resolve any recipient to `open_id` with `lark_contact_search` (P1). Present the overdue summary as an **interactive card** with a per-customer "send" button (P4). If the owner tracks AR in a Lark Base, read/update it with `lark_base_search` / `lark_base_record_upsert` (P5).

## Quick start

Pull the AR aging report, score each customer by payment history, draft a tone-matched reminder for each overdue invoice, and present them to the owner. Nothing sends until the owner says so.

```
User: "who owes me money"
→ Pull AR aging from QuickBooks
→ Cross-reference PayPal settlements (last 14 days)
→ Score each customer: good-payer / occasionally-late / repeat-late
→ Draft tone-matched reminders
→ Show summary table + drafts. Wait for "send these."
```

## Setup (first run only)

Ask the owner two questions before running for the first time:

1. **Mail connector**: default to **Lark Mail** (`lark_mail_draft_create`) for all non-PayPal draft queuing; only ask if the owner wants a different mail app.
2. **Stripe**: "Do you use Stripe for invoicing? I can include Stripe invoices in the overdue sweep." — if yes, pull Stripe overdue invoices alongside QuickBooks.

Do not ask again on subsequent runs.

## Workflow

1. **Pull overdue receivables.** Query QuickBooks AR aging for all invoices more than 1 day past due. If Stripe is enabled (owner confirmed at setup), also pull Stripe overdue invoices.

2. **Cross-reference payment history.** For each overdue customer, query PayPal for settled transactions using these parameters:
   - `transaction_status: S` (settled only — filters out pending and denied transactions that inflate result size and increase rate-limit risk)
   - Date window: **last 7 days** ending today (not 14 or 30 — wider windows are the primary cause of PayPal 429 rate limit errors)

   **If PayPal returns a 429 rate limit error:**
   - Retry once immediately with a **3-day window** instead.
   - If the retry also returns 429, skip the PayPal cross-reference entirely for this run. Flag all customers in the batch as "PayPal unavailable — verify manually" in the summary table. Proceed to scoring using QuickBooks history only. Do not silently drop the caveat.

   If a customer shows a settled payment within the query window, flag as "possibly paid — verify" and exclude from the draft queue.

3. **Score each customer.** Read [reference/tone-matching.md](reference/tone-matching.md) for scoring logic. Result: `good-payer`, `occasionally-late`, or `repeat-late`.

4. **Draft reminder emails.** One email per customer — consolidate multiple overdue invoices into one email. Match tone to score. See [reference/examples/gentle-reminder.md](reference/examples/gentle-reminder.md) and [reference/examples/firm-reminder.md](reference/examples/firm-reminder.md).

5. **Present drafts to owner.** Show a summary table first:

   | Customer | Amount Due | Days Late | Tone | Send via |
   |---|---|---|---|---|
   | Acme Corp | $1,200 | 18 days | Gentle | PayPal |
   | Smith LLC | $450 | 47 days | Firm | Lark Mail draft |

   Then show each draft email in full. Wait for owner to say "send these" or approve individually.

6. **Send or queue — only after approval.**
   - PayPal invoices: send the reminder via PayPal.
   - Non-PayPal invoices: queue as a draft via `lark_mail_draft_create` (`dry_run: true` first to preview recipient + body, P2).
   - Never send without explicit approval.

7. **Report what happened.** List what was sent, what was queued as draft, and what was flagged (possibly paid, excluded).

## Approval gates

- **Never send or queue a draft without explicit owner approval.** Present all drafts first; wait for the go-ahead.
- **Never include a customer who paid in the last 14 days.** Flag as "possibly paid — verify" instead.
- **Never send to a customer not in the QuickBooks AR report** (or Stripe, if enabled). No reminders from memory alone.
- **One approval covers one batch.** Adding a customer or changing a draft after approval starts a new round.

## Reference

- [reference/tone-matching.md](reference/tone-matching.md) — scoring logic, tone guidelines, subject line formulas
- [reference/gotchas.md](reference/gotchas.md) — known failure modes
- [reference/examples/gentle-reminder.md](reference/examples/gentle-reminder.md) — good-payer email example
- [reference/examples/firm-reminder.md](reference/examples/firm-reminder.md) — repeat-late-payer email example
