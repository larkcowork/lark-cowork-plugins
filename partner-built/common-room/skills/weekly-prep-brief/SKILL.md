---
name: weekly-prep-brief
description: "Generate a comprehensive weekly briefing for all external calls in the next 7 days. Triggers on 'weekly prep brief', 'prepare my week', 'what calls do I have this week', 'Monday prep', or any weekly planning request."
---

# Weekly Prep Brief

Generate a single comprehensive weekly briefing that covers every external customer or prospect call in the next 7 days, with per-meeting account and contact research from Common Room.

## Briefing Process

### Step 1: Get the Week's External Meetings

**Option A — Lark Calendar (default):**
Fetch all meetings in the next 7 days (or a user-specified range) from Lark Calendar with `lark_calendar_agenda`. Project with `jq` (P3) to title, time, and attendees. Filter to keep only external meetings — those with attendees from outside your organization. Discard internal-only meetings, one-on-ones with colleagues, and recurring internal syncs.

Identify for each external meeting:
- Company name
- Meeting date and time
- External attendee names and email addresses

**Resolve attendees (P1):** for each external attendee, `lark_contact_search(query="<name or email>")` → `open_id`, so per-attendee Lark context (prior minutes, shared threads) can be pulled in Step 3.

**Option B — Calendar unavailable:**
If Lark Calendar can't be reached, ask the user: "To build your weekly prep brief, I'll need your upcoming external calls. Please list them: company name, date/time, and attendee names."

Accept freeform input and parse it into a structured list before proceeding.

### Step 2: Confirm the Meeting List

Present the identified meetings to the user for confirmation before beginning research:

> "Here are the external calls I found for this week. Let me know if anything's missing or should be excluded:
> - [Company] — [Day], [Time] — [Attendees]
> - ..."

This prevents wasted research on cancelled or incorrect meetings.

### Step 3: Research Each Meeting

For each confirmed external meeting, run in parallel where possible:
1. **Account research** — full account snapshot using the account-research skill
2. **Contact research** — profile for each external attendee using the contact-research skill
3. **Prior Lark calls (P6)** — `lark_minutes_search(query="<company>", participant_ids="me")` for the last meeting's AI summary + open action items, and `lark_vc_search` for the envelope. These are your team's own recent conversations with the account — lead each meeting's prep with any unresolved action items. Project with `jq` (P3).

Common Room data is the primary source for external/community signals. After CR research, run a quick **recency check** for each company — this is supplementary, not primary:
- Search `"[company name]" news` scoped to the last 7 days
- For executive attendees, search their name for recent public posts or interviews
- Only include findings that are genuinely noteworthy (funding, leadership changes, major press). Don't pad the brief with generic news.

Depth calibration:
- For high-priority accounts (large accounts, open opportunities, renewal risk), produce full depth research
- For lower-priority or short meetings, produce abbreviated snapshots (3–4 bullets each)

### Step 4: Synthesize the Weekly Brief

Compile all per-meeting research into a single structured document, sorted by meeting date/time.

Open with a brief week-level overview that flags:
- Any accounts with urgent signals (at-risk, trial expiring, expansion opportunity)
- Any meetings that need special preparation or executive involvement
- Total external call count and estimated time commitment

## Output Format

```
# Weekly Prep Brief — Week of [Date]

## Week Overview
[2–4 bullets: key themes, flagged priorities, call count]

---

## [Monday / Tuesday / etc.]

### [Company Name] — [Time]
**Attendees:** [Names and titles]
**Meeting type:** [Discovery / QBR / Renewal / Expansion / etc. — inferred if possible]

**Company Snapshot**
[4–5 bullets: account status, top signals, recent activity]

**Attendee Profiles**
- **[Name]** ([Title]): [2–3 bullets on their signals, persona, conversation angle]
- [Repeat per attendee]

**Top Signals This Week**
[2–3 most relevant signals for this specific call]

**This Week's News** [If notable news found]
[Only genuinely noteworthy findings — funding, leadership changes, major press]

**Recommended Objectives**
[1–2 sentences: what to accomplish in this meeting]

---

[Repeat per meeting, sorted by date/time]
```

## When a Meeting Has Sparse Data

If Common Room returns limited data for a particular meeting's account or attendees, use a compressed format for that meeting instead of the full template:

```
### [Company Name] — [Time] ⚠️ Limited Data
**Attendees:** [Names and titles if known]
**Data available:** [What Common Room actually returned]

**Web Search Results**
[Findings from web search — company news, attendee LinkedIn profiles]

**Note:** Common Room has limited data on this account. The rep may want to check directly in CR or gather context from colleagues before this call.
```

Do not generate a full meeting prep section (company snapshot, signal highlights, talking points, recommended objectives) from sparse data. A short honest section is more useful than a fabricated full one.

## Quality Standards

- Keep each meeting section scannable — reps read these in the morning, often on mobile
- Always sort by date/time ascending
- Flag urgent situations prominently (risk, trial expiration, open opps) — don't bury them
- If a meeting has very thin Common Room data, use the sparse-data format above — never fill the full template with guesses
- Total brief should be readable in 10–15 minutes for a week with 4–6 meetings
- **Every fact must come from a tool call** — no invented deal context, activity, or signals

## Lark-native handoff

Common Room drives the per-account signals; Lark drives the week's collaboration scaffolding:

- **Calendar** — `lark_calendar_agenda` for the week's meetings (delegate to the `lark-calendar` skill).
- **People (P1)** — `lark_contact_search` to resolve each attendee.
- **Prior calls (P6)** — `lark_minutes_search` + `lark_vc_search` (delegate to `lark-minutes` / `lark-vc`).
- **Deliver the brief (P4/P8)** — for a Monday-morning push, surface the week overview as an interactive card via `lark_im_card_send` (one collapsible `panel` per day, `actions` row to open each event); `print_json` → `dry_run` → send. For a durable artifact the rep revisits all week, land the full brief in Wiki via `lark_wiki_node_create` (or `lark_doc_create`) instead of leaving it in chat. Plain text stays the default for an in-chat answer.
- This near-duplicates the local `weekly-review` workflow skill — prefer wiring to it for the calendar+tasks merge. See `../../connectors/LARK-PATTERNS.md`, `LARK-RECIPES.md`, `LARK-FUSION.md`.

## Reference Files

- **`references/briefing-guide.md`** — guidelines for structuring briefings, prioritization logic, and how to handle edge cases (cancelled meetings, new accounts with no data, etc.)
