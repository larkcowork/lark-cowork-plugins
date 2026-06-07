---
name: daily-briefing
description: Start your day with a prioritized sales briefing. Works standalone when you tell me your meetings and priorities, supercharged when you connect your calendar, CRM, and email. Trigger with "morning briefing", "daily brief", "what's on my plate today", "prep my day", or "start my day".
---

# Daily Sales Briefing

Get a clear view of what matters most today. This skill works with whatever you tell me, and gets richer when you connect your tools.

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md); tool map in [CONNECTORS.md](../../CONNECTORS.md)).
> **Delegate to the installed `morning-brief` skill, sales variant** (`/morning sales`) — it already fans out mail/IM/approval/tasks in parallel and merges into a ≤15-line **interactive card**. This skill frames the sales lens (pipeline alerts, deals closing). Lark sources: **Calendar** today's meetings (`lark_calendar_agenda`, project with `jq`, P3); **CRM = Lark Base** pipeline (`lark_base_search` on the Deals table — requires `search_fields`, narrow with `select_fields`, no `jq`, P5); **email** priorities (`lark_mail_*` / `lark_api` mail search); **tasks** (`lark_task_my`). Always render the briefing as an **interactive card** (`lark_im_card_send`, P4) with per-meeting rows and an `actions` footer ("Prep my next call" → call-prep), not a wall of text. Resolve any owner/attendee via `lark_contact_search` (P1). For the end-of-day variant, delegate to **`daily-digest`**.

## How It Works

```
┌─────────────────────────────────────────────────────────────────┐
│                      DAILY BRIEFING                              │
├─────────────────────────────────────────────────────────────────┤
│  ALWAYS (works standalone)                                       │
│  ✓ You tell me: today's meetings, key deals, priorities         │
│  ✓ I organize: prioritized action plan for your day             │
│  ✓ Output: scannable 2-minute briefing                          │
├─────────────────────────────────────────────────────────────────┤
│  SUPERCHARGED (when you connect your tools)                      │
│  + Calendar: auto-pull today's meetings with attendees          │
│  + CRM: pipeline alerts, tasks, deal health                     │
│  + Email: unread from key accounts, waiting on replies          │
│  + Enrichment: overnight signals on your accounts               │
└─────────────────────────────────────────────────────────────────┘
```

---

## Getting Started

When you run this skill, I'll ask for what I need:

**If no calendar connected:**
> "What meetings do you have today? (Just paste your calendar or list them)"

**If no CRM connected:**
> "What deals are you focused on this week? Any that need attention?"

**If you have connectors:**
I'll pull everything automatically and just show you the briefing.

---

## Connectors (Optional)

Connect your tools to supercharge this skill:

| Connector | What It Adds |
|-----------|--------------|
| **Calendar** | Today's meetings with attendees, times, and context |
| **CRM** | Open pipeline, deals closing soon, overdue tasks, stale deals |
| **Email** | Unread from opportunity contacts, emails waiting on replies |
| **Enrichment** | Overnight signals: funding, hiring, news on your accounts |

> **No connectors?** No problem. Tell me your meetings and deals, and I'll create your briefing.

---

## Output Format

```markdown
# Daily Briefing | [Day, Month Date]

---

## #1 Priority

**[Most important thing to do today]**
[Why it matters and what to do about it]

---

## Today's Numbers

| Open Pipeline | Closing This Month | Meetings Today | Action Items |
|---------------|-------------------|----------------|--------------|
| $[X] | $[X] | [N] | [N] |

---

## Today's Meetings

### [Time] — [Company] ([Meeting Type])
**Attendees:** [Names]
**Context:** [One-line: deal status, last touch, what's at stake]
**Prep:** [Quick action before this meeting]

### [Time] — [Company] ([Meeting Type])
**Attendees:** [Names]
**Context:** [One-line context]
**Prep:** [Quick action]

*Run `call-prep [company]` for detailed meeting prep*

---

## Pipeline Alerts

### Needs Attention
| Deal | Stage | Amount | Alert | Action |
|------|-------|--------|-------|--------|
| [Deal] | [Stage] | $[X] | [Why flagged] | [What to do] |

### Closing This Week
| Deal | Close Date | Amount | Confidence | Blocker |
|------|------------|--------|------------|---------|
| [Deal] | [Date] | $[X] | [H/M/L] | [If any] |

---

## Email Priorities

### Needs Response
| From | Subject | Received |
|------|---------|----------|
| [Name @ Company] | [Subject] | [Time] |

### Waiting On Reply
| To | Subject | Sent | Days Waiting |
|----|---------|------|--------------|
| [Name @ Company] | [Subject] | [Date] | [N] |

---

## Suggested Actions

1. **[Action]** — [Why now]
2. **[Action]** — [Why now]
3. **[Action]** — [Why now]

---

*Run `call-prep [company]` before your meetings*
*Run `call-summary` after each call*
```

---

## Execution Flow

### Step 1: Gather Context

**If connectors available (Lark-native; prefer delegating to `morning-brief`):**
```
1. Calendar → lark_calendar_agenda (today)
   - jq: ".data[] | {start, summary, attendees}" — filter to external meetings (P3)

2. CRM (Lark Base, P5) → lark_base_search on the Deals table (owner = me)
   - search_fields=["Owner"], query="<me>",
     select_fields=["Account Name","Stage","Amount","Close Date","Last Activity"],
     limit=100   (base_search REQUIRES search_fields and does NOT support jq)
   - if field names unknown, discover via
     lark_api GET /open-apis/bitable/v1/apps/{base}/tables/{table}/fields
   - flag: closing this week, no activity 7+ days, slipped close dates
   - tasks: lark_task_my (complete:false) for overdue/upcoming → jq ".data.items[]" (P3)

3. Email → lark_mail_* / lark_api mail search
   - unread from opportunity contact domains; sent w/ no reply 3+ days

4. Enrichment (external tool) → overnight signals (funding, hiring, news)
```

**If no connectors:**
```
Ask user:
1. "What meetings do you have today?"
2. "What deals are you focused on? Any closing soon or needing attention?"
3. "Anything urgent I should know about?"

Work with whatever they provide.
```

### Step 2: Prioritize

```
Priority ranking:
1. URGENT: Deal closing today/tomorrow not yet won
2. HIGH: Meeting today with high-value opportunity
3. HIGH: Unread email from decision-maker
4. MEDIUM: Deal closing this week
5. MEDIUM: Stale deal (7+ days no activity)
6. LOW: Tasks due this week

Select #1 Priority:
- If meeting with >$50K deal today → prep that
- If deal closing today → focus on close
- If urgent email from buyer → respond first
- Else → highest-value stale deal
```

### Step 3: Generate Briefing

```
Assemble sections based on available data:

1. #1 Priority — Always include (even if simple)
2. Today's Numbers — If CRM connected, otherwise skip
3. Today's Meetings — From calendar or user input
4. Pipeline Alerts — If CRM connected
5. Email Priorities — If email connected
6. Suggested Actions — Always include top 3 actions
```

---

## Quick Mode

Say "quick brief" or "tldr my day" for abbreviated version:

```markdown
# Quick Brief | [Date]

**#1:** [Priority action]

**Meetings:** [N] — [Company 1], [Company 2], [Company 3]

**Alerts:**
- [Alert 1]
- [Alert 2]

**Do Now:** [Single most important action]
```

---

## End of Day Mode

Say "wrap up my day" or "end of day summary" after your last meeting:

```markdown
# End of Day | [Date]

**Completed:**
- [Meeting 1] — [Outcome]
- [Meeting 2] — [Outcome]

**Pipeline Changes:**
- [Deal] moved to [Stage]

**Tomorrow's Focus:**
- [Priority 1]
- [Priority 2]

**Open Loops:**
- [ ] [Unfinished item needing follow-up]
```

---

## Tips

1. **Connect your calendar first** — Biggest time saver
2. **Add CRM second** — Unlocks pipeline alerts
3. **Even without connectors** — Just tell me your meetings and I'll help prioritize

---

## Related Skills

- **morning-brief** — installed start-of-day orchestrator (delegate; use the `sales` variant)
- **daily-digest** — installed end-of-day digest (use for the "wrap up my day" mode)
- **call-prep** — Deep prep for any specific meeting
- **call-summary** — Process notes after calls
- **account-research** — Research a company before first meeting
