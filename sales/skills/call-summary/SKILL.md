---
name: call-summary
description: Process call notes or a transcript — extract action items, draft follow-up email, generate internal summary. Use when pasting rough notes or a transcript after a discovery, demo, or negotiation call, drafting a customer follow-up, logging the activity for your CRM, or capturing objections and next steps for your team.
argument-hint: "<call notes or transcript>"
---

# /call-summary

> If you see unfamiliar placeholders or need to check which tools are connected, see [CONNECTORS.md](../../CONNECTORS.md).

Process call notes or a transcript to extract action items, draft follow-up communications, and update records.

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md) — especially P6 Minutes→tasks, P2 safe-mutation, P5 Base-as-SoR; [LARK-RECIPES](../../../connectors/LARK-RECIPES.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md)).
> This is the canonical post-call flow — **delegate to the installed `deal-update` skill** (pull Minutes → extract pain/budget/timeline → update Base record → draft follow-up) and reuse **`meeting-prep`** (after-phase) for the action-items→tasks half. Don't re-summarize from scratch: prefer the real **Minutes AI artifacts** via `lark_minutes_search(participant_ids="me", query="<account>")` (P6), falling back to the pasted notes only when no recording exists. Resolve every action-item owner to an `open_id` with `lark_contact_search` (P1). Create tasks with `lark_task_create` (`dry_run: true` first, P2). The CRM is a **Lark Base** — log the activity / update the deal stage with `lark_base_record_upsert` (`dry_run` first). The follow-up email is a **draft only**: use `lark_mail_draft_create` (never auto-send; `lark_mail_send` requires `confirm_send: true`, P2). Optionally post the recap as an **interactive card** (`lark_im_card_send`, P4).

## Usage

```
/call-summary <notes or transcript>
```

Process these call notes: $ARGUMENTS

If a file is referenced: @$1

---

## How It Works

```
┌─────────────────────────────────────────────────────────────────┐
│                      CALL SUMMARY                                │
├─────────────────────────────────────────────────────────────────┤
│  STANDALONE (always works)                                       │
│  ✓ Paste call notes or transcript                               │
│  ✓ Extract key discussion points and decisions                  │
│  ✓ Identify action items with owners and due dates              │
│  ✓ Surface objections, concerns, and open questions             │
│  ✓ Draft customer-facing follow-up email                        │
│  ✓ Generate internal summary for your team                      │
├─────────────────────────────────────────────────────────────────┤
│  SUPERCHARGED (when you connect your tools)                      │
│  + Transcripts: Pull recording automatically (e.g. Gong, Lark Minutes) │
│  + CRM: Update opportunity, log activity, create tasks          │
│  + Email: Send follow-up directly from draft                    │
│  + Calendar: Link to meeting, pull attendee context             │
└─────────────────────────────────────────────────────────────────┘
```

---

## What I Need From You

**Option 1: Paste your notes**
Just paste whatever you have — bullet points, rough notes, stream of consciousness. I'll structure it.

**Option 2: Paste a transcript**
If you have a full transcript from your video conferencing tool (e.g. Zoom, Teams) or conversation intelligence tool (e.g. Gong, Lark Minutes), paste it. I'll extract the key moments.

**Option 3: Describe the call**
Tell me what happened: "Had a discovery call with Acme Corp. Met with their VP Eng and CTO. They're evaluating us vs Competitor X. Main concern is integration timeline."

---

## Output

### Internal Summary
```markdown
## Call Summary: [Company] — [Date]

**Attendees:** [Names and titles]
**Call Type:** [Discovery / Demo / Negotiation / Check-in]
**Duration:** [If known]

### Key Discussion Points
1. [Topic] — [What was discussed, decisions made]
2. [Topic] — [Summary]

### Customer Priorities
- [Priority 1 they expressed]
- [Priority 2]

### Objections / Concerns Raised
- [Concern] — [How you addressed it / status]

### Competitive Intel
- [Any competitor mentions, what was said]

### Action Items
| Owner | Action | Due |
|-------|--------|-----|
| [You] | [Task] | [Date] |
| [Customer] | [Task] | [Date] |

### Next Steps
- [Agreed next step with timeline]

### Deal Impact
- [How this call affects the opportunity — stage change, risk, acceleration]
```

### Customer Follow-Up Email
```
Subject: [Meeting recap + next steps]

Hi [Name],

Thank you for taking the time to meet today...

[Key points discussed]

[Commitments you made]

[Clear next step with timeline]

Best,
[You]
```

---

## Email Style Guidelines

When drafting customer-facing emails:

1. **Be concise but informative** — Get to the point quickly. Customers are busy.
2. **No markdown formatting** — Don't use asterisks, bold, or other markdown syntax. Write in plain text that looks natural in any email client.
3. **Use simple structure** — Short paragraphs, line breaks between sections. No headers or bullet formatting unless the customer's email client will render it.
4. **Keep it scannable** — If listing items, use plain dashes or numbers, not fancy formatting.

**Good:**
```
Here's what we discussed:
- Quote for 20 seats at $480/seat/year
- W9 and supplier onboarding docs
- Point of contact for the contract
```

**Bad:**
```
**What You Need from Us:**
- Quote for 20 seats at $480/seat/year
```

---

## If Connectors Available (Lark-native)

**Transcripts — Lark Minutes (P6):**
- `lark_minutes_search(participant_ids="me", query="<account>")` → find the call
- Pull the AI summary + action items + transcript (project with `jq`, P3) instead of
  re-deriving; delegate to **`lark-minutes`** for the AI artifacts. `lark_vc_search` for
  the meeting envelope (attendees/room) when you only need metadata.

**CRM — Lark Base (P5):**
- Update the opportunity stage / log the call as an activity / update next-step via
  `lark_base_record_upsert(base_token, table_id, fields={...})` — `dry_run: true` first (P2).
- For each action item: `lark_contact_search` the owner (P1) → `lark_task_create(dry_run)`
  → confirm → commit. Complete tasks with `lark_task_complete`.
- Delegate multi-table / schema work to **`lark-base`**.

**Email — Lark Mail:**
- Create a **draft** with `lark_mail_draft_create` so the rep eyeballs it in the Mail UI first.
- Only send with `lark_mail_send` + `confirm_send: true` after explicit approval (P2). Never auto-send.

**Recap card (P4):** optionally post the internal summary as a card via `lark_im_card_send`
(`print_json` → `dry_run` → send) with an `actions` row to create the tasks / open the deal record.

---

## Tips

1. **More detail = better output** — Even rough notes help. "They seemed concerned about X" is useful context.
2. **Name the attendees** — Helps me structure the summary and assign action items.
3. **Flag what matters** — If something was important, tell me: "The big thing was..."
4. **Tell me the deal stage** — Helps me tailor the follow-up tone and next steps.
