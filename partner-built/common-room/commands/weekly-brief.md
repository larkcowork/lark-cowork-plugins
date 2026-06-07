---
description: Generate a weekly prep briefing from your calendar and Common Room
argument-hint: [date range, defaults to next 7 days]
---

Generate a weekly prep briefing using Common Room and your calendar.

Follow the weekly-prep-brief skill:
1. Use `lark_calendar_agenda` (project with `jq`) to retrieve all external customer-facing meetings scheduled for the next 7 days (or the date range specified in "$ARGUMENTS"). Filter out internal meetings — focus on calls with customers, prospects, or partners. Resolve each attendee with `lark_contact_search` (P1).
2. If Lark Calendar can't be reached, ask the user to list their external calls (company name, date, attendees).
3. For each external meeting, run account research and contact research on attendees in parallel.
4. Compile into a single weekly briefing: week overview + per-meeting sections sorted by date.

Keep each per-meeting section tight and scannable. Total briefing should be readable in under 10 minutes.
