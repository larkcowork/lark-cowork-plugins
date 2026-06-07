---
description: Draft a well-formatted Lark IM announcement and save it as a draft
---

Given the topic or context provided in $ARGUMENTS:

1. Ask the user the following clarifying questions (skip any that are already clear from the provided context):
   - Which channel should this announcement be posted in?
   - Who is the target audience?
   - What is the key message or call to action?
   - Is there a deadline or date to highlight?
   - What tone is appropriate — formal, casual, or urgent?

2. Compose the announcement following Lark IM formatting best practices:
   - Use Lark IM's mrkdwn syntax: `*bold*` for emphasis (not `**bold**`), `_italic_` for secondary emphasis, `>` for callouts.
   - Lead with the most important information — don't bury the point.
   - Use a clear, descriptive opening line that works as a headline.
   - Keep paragraphs short (2-3 sentences max).
   - Use bullet points for lists of items or action steps.
   - Include relevant emoji sparingly to aid scanning (e.g., :mega: for announcements, :calendar: for dates, :point_right: for action items).
   - End with a clear call to action or next step if applicable.

3. Present the draft to the user for review. Offer to adjust tone, length, or formatting.

4. Once the user approves, resolve the target chat (if you need the exact `chat_id`, list chats via `lark_api` GET `/open-apis/im/v1/chats`; resolve any direct recipient via `lark_contact_search`). Then preview the send with `lark_im_send` using `dry_run: true` (P2) so the user can eyeball the planned recipient + body before anything is committed. For higher-stakes outreach that should be reviewed in a UI, draft via `lark_mail_draft_create` instead.

5. Let the user know the announcement is staged for review — they confirm, then re-send `lark_im_send` (without `dry_run`) to post it.
