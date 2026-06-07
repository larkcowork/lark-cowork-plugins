---
name: call-prep
description: "Prepare for a customer or prospect call using Common Room signals. Triggers on 'prep me for my call with [company]', 'prepare for a meeting with [company]', 'what should I know before talking to [company]', or any call preparation request."
---

# Call Prep

Produce a complete, scannable call prep brief by combining account research, contact research, and signal synthesis from Common Room.

## Prep Process

### Step 1: Identify the Account and Attendees

Parse what the user has provided:
- **Company name** — required; look up the account in Common Room
- **Attendee names** — optional; if provided, research each one

**Calendar lookup (Lark):** Pull the upcoming meeting from Lark Calendar to auto-fill attendees, time, and agenda. Use `lark_calendar_agenda` (project with `jq` — P3 — to the meeting whose title/attendees match the named company) to surface attendee names, the meeting time, and any agenda/description. If you only need to confirm a slot or the rep's availability around the call, `lark_calendar_freebusy` covers it. Use this to fill gaps the user didn't provide.

**Resolve attendees (P1):** For each external attendee name you find, resolve them with `lark_contact_search(query="<name or email>")` → `open_id`. You need the `open_id` to pull internal context (prior threads, shared meetings, past minutes) in the steps below.

If neither attendees nor a calendar match can be found, ask: "Who will be on the call from [Company]? I can research each attendee to make your prep more useful."

### Step 2: Run Account Research

Use the account-research skill process to build a full account snapshot. For call prep, prioritize:
- Recent product signals (what are they doing in the product right now?)
- Open opportunities or renewal timeline
- Any risk signals (declining usage, support tickets, churned seats)
- Key recent events (funding, executive change, new hire)

When reviewing activity history, prioritize Gong and call recording activities — these provide direct context about previous conversations. Do not filter out call recordings by activity origin.

**Prior Lark calls with this account (P6):** Common Room captures third-party/community signals, but your team's own past calls with this customer live in Lark. Pull them as first-class prep context: `lark_minutes_search(query="<company>", participant_ids="me")` for the AI summary + action items of previous meetings, and `lark_vc_search` for the meeting envelope (who/when/room) when you only need metadata. Lead the brief with the most recent unresolved action items from these — they are the freshest, highest-signal talking points. Project reads with `jq` (P3).

### Step 3: Run Contact Research for Each Attendee

For each external attendee, use the contact-research skill process. For call prep, focus on:
- Role and influence in the buying process
- Their personal activity and engagement history
- Any recent signals that suggest their current mood/priorities
- Spark persona classification if available

### Step 4: Synthesize Talking Points and Objectives

Based on the combined account and contact research:
- Identify the **call objective** (e.g., discovery, demo, expansion conversation, renewal, QBR)
- Generate **3–5 tailored talking points** grounded in specific signal data
- Anticipate **2–3 likely objections or topics** the customer may raise
- Suggest a **recommended outcome** for the call

When the user's company context is available (see `references/my-company-context.md`), tailor talking points to the user's product and value proposition.

### Step 5: Recency Check (Web Search)

After gathering all Common Room data, run a quick recency check to catch anything that happened since the last CR data sync. This is supplementary — CR data drives the prep; web search only adds recency.

**Company news:** Search `"[company name]" news` filtered to the last 14 days. Look for funding announcements, product launches, leadership changes, layoffs, partnerships, or press coverage.

**Attendee presence:** For each external attendee, search `"[full name]" "[company name]"` — look for recent articles, LinkedIn posts, conference talks, podcasts, or published opinions.

If a company news item is significant (e.g., just raised a round, announced a major hire), flag it in Signal Highlights. Otherwise, include findings briefly — don't let web search results overshadow CR signals.

## Output Format

The output adapts to how much data Common Room returned. Only include sections where you have real data. Never fill a section with invented details.

### When data is rich (multiple field groups returned, activity history, scores, signals):

```
## Call Prep: [Company] — [Date/Time if known]

**Meeting Context**
[Attendees, meeting type, and any known agenda]

---

### Company Snapshot
[4–6 bullets: key account status, signals, and recent activity]

---

### Attendee Profiles

**[Attendee Name] — [Title]**
[3–4 bullets: role, recent activity, Spark persona if available, personal hook]

[Repeat for each attendee]

---

### Signal Highlights
[Top 3 signals most relevant to this specific call]

---

### Talking Points
1. [Point tied to a specific signal]
2. [Point tied to a specific signal]
3. [Point tied to a specific signal]

### Likely Topics / Objections to Prepare For
- [Topic or objection + suggested response]
- [Topic or objection + suggested response]

### Recommended Call Outcome
[1–2 sentences: what success looks like for this meeting]
```

### When data is sparse (few fields returned, no activity, null sparkSummary):

```
## Call Prep: [Company] — [Date/Time if known]

**Data available:** [List exactly what Common Room returned — e.g., "Name, title, email, two tags. No activity history, no scores, no Spark data."]

### What I Found
[Only the fields actually returned, presented as-is]

### Web Search Results
[Findings from web search on the company and attendees — or "No significant results"]

### Suggested Next Steps
- I can pull [specific field groups] from Common Room if available
- I can run deeper web searches on [specific topics]
- You may want to check Common Room directly for [what's missing]
```

Do not generate a full call prep brief from sparse data. A short honest output is always better than a long fabricated one.

### Surface the brief as an interactive card (P4)

A call-prep brief is a digest the rep acts on right before a meeting — it belongs in chat as an **interactive card**, not a wall of text. When the rep wants it pushed to Lark (or it's part of a morning/weekly flow), post it via `lark_im_card_send`: a colored `header` (company + time), collapsible `panel` sections for Company Snapshot / Attendee Profiles / Talking Points, and an `actions` row ("Open calendar event", "Draft outreach"). Always `print_json: true` to validate the spec, then `dry_run: true`, then send. The plain-text format above remains the default for an in-chat answer.

## Quality Standards

- Ground every talking point in a real signal — no generic filler
- Keep the brief tight — it should be readable in 5 minutes or less
- Flag unknowns explicitly — if attendee research is thin, say so
- Time-box the research — don't over-research at the expense of speed
- **Never invent deal context** — no fabricated proposals, competitor comparisons, pricing, trial terms, or objections not returned by a tool call

## Lark-native handoff

Common Room owns the account/contact signal layer; Lark owns the collaboration layer this brief sits on:

- **Calendar** — `lark_calendar_agenda` / `lark_calendar_freebusy` for the meeting + slot (delegate depth to the `lark-calendar` skill).
- **People (P1)** — `lark_contact_search` to resolve every attendee to an `open_id`.
- **Prior calls (P6)** — `lark_minutes_search` (AI summary + action items) and `lark_vc_search` (envelope); delegate to the `lark-minutes` / `lark-vc` skills.
- **Card (P4)** — `lark_im_card_send` for the chat-native brief; card YAML grammar lives in the `lark-im` skill (`references/spec-reference.md`).
- See `../../connectors/LARK-PATTERNS.md`, `LARK-RECIPES.md`, `LARK-FUSION.md`.

## Reference Files

- **`references/call-types-guide.md`** — guidance for different call types (discovery, expansion, renewal, QBR) and how to tailor prep accordingly
