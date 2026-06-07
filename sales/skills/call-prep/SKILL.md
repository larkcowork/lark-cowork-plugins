---
name: call-prep
description: Prepare for a sales call with account context, attendee research, and suggested agenda. Works standalone with user input and web research, supercharged when you connect your CRM, email, chat, or transcripts. Trigger with "prep me for my call with [company]", "I'm meeting with [company] prep me", "call prep [company]", or "get me ready for [meeting]".
---

# Call Prep

Get fully prepared for any sales call in minutes. This skill works with whatever context you provide, and gets significantly better when you connect your sales tools.

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md); tool map in [CONNECTORS.md](../../CONNECTORS.md)).
> This is a meeting-prep flow — **delegate the heavy lifting to the installed `meeting-prep` skill** (before-phase: pull context; after-phase: action items → tasks) and to **`contact-360`** for a per-attendee dossier. The collaboration layer is all Lark: **Calendar** to find the meeting (`lark_calendar_agenda`) and check attendee free/busy (`lark_calendar_freebusy`); **CRM = Lark Base** account/deal history (`lark_base_search`, requires `search_fields`; narrow with `select_fields` — base_search does NOT support `jq`, P5); **chat** internal intel (`lark_im_search`); **transcripts** from prior calls (`lark_minutes_search` / `lark_vc_search`, P6). Resolve every attendee name to an `open_id` via `lark_contact_search` (P1) before researching or @-mentioning. Web research + external enrichment stay as-is. Deliver the prep brief as an **interactive card** (`lark_im_card_send`, P4) and/or land it in **Wiki** (`lark_wiki_node_create`, P8).

## How It Works

```
┌─────────────────────────────────────────────────────────────────┐
│                        CALL PREP                                 │
├─────────────────────────────────────────────────────────────────┤
│  ALWAYS (works standalone)                                       │
│  ✓ You tell me: company, meeting type, attendees                │
│  ✓ Web search: recent news, funding, leadership changes         │
│  ✓ Company research: what they do, size, industry               │
│  ✓ Output: prep brief with agenda and questions                 │
├─────────────────────────────────────────────────────────────────┤
│  SUPERCHARGED (when you connect your tools)                      │
│  + CRM: account history, contacts, opportunities, activities    │
│  + Email: recent threads, open questions, commitments           │
│  + Chat: internal discussions, colleague insights               │
│  + Transcripts: prior call recordings, key moments              │
│  + Calendar: auto-find meeting, pull attendees                  │
└─────────────────────────────────────────────────────────────────┘
```

---

## Getting Started

When you run this skill, I'll ask for what I need:

**Required:**
- Company or contact name
- Meeting type (discovery, demo, negotiation, check-in, etc.)

**Helpful if you have it:**
- Who's attending (names and titles)
- Any context you want me to know (paste prior notes, emails, etc.)

If you've connected your CRM, email, or other tools, I'll pull context automatically and skip the questions.

---

## Connectors (Optional)

Connect your tools to supercharge this skill:

| Connector | What It Adds |
|-----------|--------------|
| **CRM** | Account details, contact history, open deals, recent activities |
| **Email** | Recent threads with the company, open questions, attachments shared |
| **Chat** | Internal chat discussions (e.g. Lark IM) about the account, colleague insights |
| **Transcripts** | Prior call recordings, topics covered, competitor mentions |
| **Calendar** | Auto-find the meeting, pull attendees and description |

> **No connectors?** No problem. Just tell me about the meeting and paste any context you have. I'll research the rest.

---

## Output Format

```markdown
# Call Prep: [Company Name]

**Meeting:** [Type] — [Date/Time if known]
**Attendees:** [Names with titles]
**Your Goal:** [What you want to accomplish]

---

## Account Snapshot

| Field | Value |
|-------|-------|
| **Company** | [Name] |
| **Industry** | [Industry] |
| **Size** | [Employees / Revenue if known] |
| **Status** | [New prospect / Active opportunity / Customer] |
| **Last Touch** | [Date and summary] |

---

## Who You're Meeting

### [Name] — [Title]
- **Background:** [Career history, education if found]
- **LinkedIn:** [URL]
- **Role in Deal:** [Decision maker / Champion / Evaluator / etc.]
- **Last Interaction:** [Summary if known]
- **Talking Point:** [Something personal/professional to reference]

[Repeat for each attendee]

---

## Context & History

**What's happened so far:**
- [Key point from prior interactions]
- [Open commitments or action items]
- [Any concerns or objections raised]

**Recent news about [Company]:**
- [News item 1 — why it matters]
- [News item 2 — why it matters]

---

## Suggested Agenda

1. **Open** — [Reference last conversation or trigger event]
2. **[Topic 1]** — [Discovery question or value discussion]
3. **[Topic 2]** — [Address known concern or explore priority]
4. **[Topic 3]** — [Demo section / Proposal review / etc.]
5. **Next Steps** — [Propose clear follow-up with timeline]

---

## Discovery Questions

Ask these to fill gaps in your understanding:

1. [Question about their current situation]
2. [Question about pain points or priorities]
3. [Question about decision process and timeline]
4. [Question about success criteria]
5. [Question about other stakeholders]

---

## Potential Objections

| Objection | Suggested Response |
|-----------|-------------------|
| [Likely objection based on context] | [How to address it] |
| [Common objection for this stage] | [How to address it] |

---

## Internal Notes

[Any internal chat context (e.g. Lark IM), colleague insights, or competitive intel]

---

## After the Call

Run **call-summary** (or delegate to `deal-update`) to:
- Extract action items from Minutes → Lark Task (P6)
- Update the deal in your Lark Base CRM (P5)
- Draft the follow-up in Lark Mail (draft only, P2)
```

---

## Execution Flow

### Step 1: Gather Context

**If connectors available (Lark-native; prefer delegating to `meeting-prep`):**
```
1. Calendar → lark_calendar_agenda — find the upcoming meeting matching the company
   - jq: ".data[] | {summary, start_time, attendees, description}" (data IS the array, P3)
   - lark_calendar_freebusy on attendees if you need to (re)schedule

2. CRM (Lark Base, P5) → lark_base_search(base_token, table_id, query="<account>",
     search_fields=["Account Name"],            # REQUIRED — field(s) to match on
     select_fields=["Account Name","Stage","Owner","Amount","Last Touch","Next Step"],
     limit=5)                                    # base_search does NOT support jq
   - if field names unknown, discover via
     lark_api GET /open-apis/bitable/v1/apps/{base}/tables/{table}/fields
   - related contacts + last activities from the Contacts/Activities tables
   - delegate to `lark-base` for lookups / multi-table joins

3. Email → lark_mail_* / lark_api mail search — threads with the company domain (30d)
   - extract key topics, open questions, commitments

4. Chat → lark_im_search — internal mentions of the company (30d)
   - extract colleague insights, competitive intel

5. Transcripts (P6) → lark_minutes_search(participant_ids="me", query="<account>")
   - pull the AI summary + action items + objections from prior calls
   - lark_vc_search for the meeting envelope (who/when/room) when you only need metadata
```
Resolve attendee names → `open_id` with `lark_contact_search` first (P1).

**If no connectors:**
```
1. Ask user:
   - "What company are you meeting with?"
   - "What type of meeting is this?"
   - "Who's attending? (names and titles if you know)"
   - "Any context you want me to know? (paste notes, emails, etc.)"

2. Accept whatever they provide and work with it
```

### Step 2: Research Supplement

**Always run (web search):**
```
1. "[Company] news" — last 30 days
2. "[Company] funding" — recent announcements
3. "[Company] leadership" — executive changes
4. "[Company] + [industry] trends" — relevant context
5. Attendee LinkedIn profiles — background research
```

### Step 3: Synthesize & Generate

```
1. Combine all sources into unified context
2. Identify gaps in understanding → generate discovery questions
3. Anticipate objections based on stage and history
4. Create suggested agenda tailored to meeting type
5. Output formatted prep brief
```

---

## Meeting Type Variations

### Discovery Call
- Focus on: Understanding their world, pain points, priorities
- Agenda emphasis: Questions > Talking
- Key output: Qualification signals, next step proposal

### Demo / Presentation
- Focus on: Their specific use case, tailored examples
- Agenda emphasis: Show relevant features, get feedback
- Key output: Technical requirements, decision timeline

### Negotiation / Proposal Review
- Focus on: Addressing concerns, justifying value
- Agenda emphasis: Handle objections, close gaps
- Key output: Path to agreement, clear next steps

### Check-in / QBR
- Focus on: Value delivered, expansion opportunities
- Agenda emphasis: Review wins, surface new needs
- Key output: Renewal confidence, upsell pipeline

---

## Tips for Better Prep

1. **More context = better prep** — Paste emails, notes, anything you have
2. **Name the attendees** — Even just titles help me research
3. **State your goal** — "I want to get them to agree to a pilot"
4. **Flag concerns** — "They mentioned budget is tight"

---

## Deliver (Lark-native)

- Post the prep brief as an **interactive card** (P4): `header` with company + meeting type,
  `panel` sections for attendees / agenda / discovery questions, an `actions` footer
  ("Open in Wiki" / "Draft follow-up"). `lark_im_card_send` with `print_json: true` →
  `dry_run: true` → send. Card YAML grammar lives in the **`lark-im`** skill.
- For a durable brief, `lark_wiki_node_create` then fill via **`lark-doc`** (P8).

## Related Skills

- **meeting-prep** — installed two-phase meeting flow (delegate to this for execution)
- **contact-360** — per-attendee 360° relationship brief
- **account-research** — Deep dive on a company before first contact
- **call-summary** — Process call notes and execute post-call workflow
- **draft-outreach** — Write personalized outreach after research
