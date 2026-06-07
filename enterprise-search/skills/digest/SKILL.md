---
name: digest
description: Generate a daily or weekly digest of activity across all connected sources. Use when catching up after time away, starting the day and wanting a summary of mentions and action items, or reviewing a week's decisions and document updates grouped by project.
argument-hint: "[--daily | --weekly | --since <date>]"
---

# Digest Command

> If you see unfamiliar placeholders or need to check which tools are connected, see [CONNECTORS.md](../../CONNECTORS.md).

Scan recent activity across all connected sources and generate a structured digest highlighting what matters.

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md),
> [LARK-RECIPES](../../../connectors/LARK-RECIPES.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md)).
> Gather activity by fanning out the `lark_*` reads in parallel, each with a `jq` projection (P3); pull
> meeting action items from `lark_minutes_search` instead of re-deriving them (P6); resolve any name to
> an `open_id` with `lark_contact_search` (P1). **Deliver the digest as an interactive card** via
> `lark_im_card_send` (P4) — action items as `item` rows with side buttons, decisions/updates in
> collapsible `panel` sections — not a wall of text. A daily digest overlaps heavily with the installed
> `daily-digest` skill, and a start-of-day version with `morning-brief`; **prefer delegating** to those
> for the polished card, using this command's structure as the outline.

## Instructions

### 1. Parse Flags

Determine the time window from the user's input:

- `--daily` — Last 24 hours (default if no flag specified)
- `--weekly` — Last 7 days

The user may also specify a custom range:
- `--since yesterday`
- `--since Monday`
- `--since 2025-01-20`

### 2. Check Available Sources

On Lark every source resolves to the `lark` MCP server (same mapping as the search command):

- **~~chat** — `lark_im_search` (groups, DMs, mentions)
- **~~email** — `lark_api` mail search (inbox, threads)
- **~~cloud storage** — `lark_doc_search` (recently modified docs shared with user)
- **~~project tracker** — `lark_task_my`, `lark_base_search` (tasks assigned/completed, tracker rows)
- **~~CRM** — CRM **Base** via `lark_base_search` (opportunity/account activity, P5)
- **~~knowledge base** — `lark_doc_search` (updated Wiki pages), `lark_minutes_search` (new meeting artifacts, P6)
- Also useful: `lark_calendar_agenda` (today's/upcoming events)

If the `lark` MCP server isn't authenticated, guide the user:
```
To generate a digest, the lark MCP server needs to be authenticated.
Run `auth login` (see the lark-shared skill) to connect your workspace.
```

### 3. Gather Activity from Each Source

Fire these in one batched turn, each with a `jq` projection scoped to the time window (P3):

**~~chat — `lark_im_search`:**
- Messages mentioning the user; threads the user participated in; new messages in key groups.
- `jq=".data.messages[] | {chat, sender, text, create_time}"`.

**~~email — `lark_api` mail search:**
- Recent inbox messages; threads with new replies; mails with action items / questions for the user.

**~~cloud storage — `lark_doc_search`:**
- Docs recently modified or shared with the user; project with `jq` to title/url/edit_time.

**~~project tracker — `lark_task_my` + `lark_base_search`:**
- `lark_task_my(jq=".data.items[] | {id, summary, due, completed}")` (new/updated/due-soon).
- Tracker Base rows changed in the window via `lark_base_search`.

**~~CRM — CRM Base (`lark_base_search`, P5):**
- Opportunity stage changes; new activity on owned accounts; updated contacts — `jq` to the changed fields.

**Meetings — `lark_minutes_search` (P6):**
- `lark_minutes_search(participant_ids="me", query=..., jq=".data.items[] | {title, summary, action_items}")`
  — pull AI summaries/action items directly; don't re-summarize. `lark_vc_search` for past-meeting envelopes.

**~~knowledge base — `lark_doc_search`:**
- Recently updated Wiki docs / new docs in watched areas.

**Calendar — `lark_calendar_agenda`:**
- Today's/recent events for the time-window framing.

### 4. Identify Key Items

From all gathered activity, extract and categorize:

**Action Items:**
- Direct requests made to the user ("Can you...", "Please...", "@user")
- Tasks assigned or due soon
- Questions awaiting the user's response
- Review requests

**Decisions:**
- Conclusions reached in threads or emails
- Approvals or rejections
- Policy or direction changes

**Mentions:**
- Times the user was mentioned or referenced
- Discussions about the user's projects or areas

**Updates:**
- Status changes on projects the user follows
- Document updates in the user's domain
- Completed items the user was waiting on

### 5. Group by Topic

Organize the digest by topic, project, or theme rather than by source. Merge related activity across sources:

```
## Project Aurora
- IM: Design review thread concluded — team chose Option B (design group, Tuesday)
- Mail: Sarah sent updated spec incorporating feedback (Wednesday)
- Docs: "Aurora API Spec v3" updated by Sarah (Wednesday)
- Tasks: 3 tasks moved to In Progress, 2 completed
- Minutes: "Aurora Sync" AI summary — owner assigned for rollout (Tuesday)

## Budget Planning
- Mail: Finance team requesting Q2 projections by Friday
- IM: Todd shared template in finance group (Monday)
- Docs: "Q2 Budget Template" shared with you (Monday)
```

### 6. Format and Deliver the Digest

**Deliver as an interactive card** (P4) via `lark_im_card_send` — this is the Lark superpower for a
digest. Structure: a colored `header` ("Daily Digest — <date>"), an Action Items section as `item`
rows each with a side button (e.g. *Open* / *Complete* → `lark_task_complete`), Decisions/Updates in
collapsible `panel` sections, a `note` footer with the summary stats. Always `print_json: true` to
validate the spec, then `dry_run: true`, then send (P2). Card grammar: the `lark-im` skill
(`references/spec-reference.md`). For the polished implementation, delegate to `daily-digest` /
`morning-brief`. Fall back to the markdown layout below only if the user wants it inline:

```
# [Daily/Weekly] Digest — [Date or Date Range]

Sources scanned: IM, Mail, Docs, Tasks, Minutes [others]

## Action Items (X items)
- [ ] [Action item 1] — from [person], [source] ([date])
- [ ] [Action item 2] — from [person], [source] ([date])

## Decisions Made
- [Decision 1] — [context] ([source], [date])
- [Decision 2] — [context] ([source], [date])

## [Topic/Project Group 1]
[Activity summary with source attribution]

## [Topic/Project Group 2]
[Activity summary with source attribution]

## Mentions
- [Mention context] — [source] ([date])

## Documents Updated
- [Doc name] — [who modified, what changed] ([date])
```

### 7. Handle Unavailable Sources

If any source fails or is unreachable:
```
Note: Could not reach [source name] for this digest.
The following sources were included: [list of successful sources].
```

Do not let one failed source prevent the digest from being generated. Produce the best digest possible from available sources.

### 8. Summary Stats

End with a quick summary:
```
---
[X] action items · [Y] decisions · [Z] mentions · [W] doc updates
Across [N] sources · Covering [time range]
```

## Notes

- Default to `--daily` if no flag is specified
- Group by topic/project, not by source — users care about what happened, not where it happened
- Action items should always be listed first — they are the most actionable part of a digest
- Deduplicate cross-source activity (same decision in IM and mail = one entry)
- For weekly digests, prioritize significance over completeness — highlight what matters, skip noise
- If the user has a memory system (CLAUDE.md), use it to decode people names and project references
- Include enough context in each item that the user can decide whether to dig deeper without clicking through
- Deliver as an interactive card (`lark_im_card_send`, P4); pull meeting items from `lark_minutes_search` (P6); for a polished daily/start-of-day card, delegate to `daily-digest` / `morning-brief`
