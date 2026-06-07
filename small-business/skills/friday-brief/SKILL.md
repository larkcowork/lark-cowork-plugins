---
name: friday-brief
description: Delivers the Friday end-of-week pulse — revenue vs prior week, top sellers, wins and watches. Accepts optional lookback window of 7 or 14 days.
allowed-tools: Read, WebFetch, Bash
---

Run the Friday wins-and-watches briefing. Pull the numbers, surface what matters, and give the owner a clean end-of-week picture.

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md)). Revenue data stays on the external finance/CRM connectors; **delivery is Lark**. Post the wins/watches as an **interactive card** (`lark_im_card_send`, P4) after resolving the recipient with `lark_contact_search` (P1) and previewing with `print_json` + `dry_run` (P2); email-to-self uses `lark_mail_draft_create`; save uses `lark_drive_upload`. For the end-of-day variant, the installed **`daily-digest`** skill is a near-drop-in.

Parse arguments:
- `--lookback` (default: `7d`) — `7d` for one week or `14d` for a two-week rolling comparison

## Step 1 — Revenue pulse

Using the `business-pulse` skill workflow:

1. Pull PayPal transactions for the lookback period.
2. Pull any HubSpot deal closes for the same window.
3. Calculate week-over-week revenue delta.
4. Surface top 3 revenue sources (product / customer / channel) ranked by contribution.

## Step 2 — Sales breakdown

1. List the top 5 selling products/services by volume and revenue.
2. List the bottom 3 (anything that moved less than expected vs. prior period).
3. Flag any items with a sudden spike or drop (>20% change).

## Step 3 — Wins and watches summary

Format the output as:

```
Friday Brief — {date}

WINS
• {win 1}
• {win 2}
• {win 3}

WATCHES
• {watch 1} — {recommended action}
• {watch 2} — {recommended action}

Revenue this week: ${amount} ({+/-}X% vs last week)
```

## Connector failures

Run with whatever is connected — this command degrades gracefully. If PayPal is missing, skip transaction data and note "PayPal not connected — revenue data from HubSpot deals only." If HubSpot is missing, skip deal closes and note it. If neither is connected, stop and tell the owner: "No revenue sources connected. Connect PayPal or HubSpot to run the Friday brief."

## Approval gates

- **Never send or post this brief automatically.** Always display it for the owner to review first.
- **Never auto-cancel or modify anything.** Surface the data and recommendations only.

## Output

End with the formatted brief and ask the owner: "Want me to post this to your team chat, email it to yourself, or save it?"
- **Post** → build a card, resolve recipient (P1), `print_json` → `dry_run` → `lark_im_card_send`.
- **Email** → `lark_mail_draft_create` (never auto-send; the owner sends from the Mail UI, P2).
- **Save** → `lark_drive_upload`, or `lark_wiki_node_create` if the owner keeps a weekly-log Wiki space (P8).
