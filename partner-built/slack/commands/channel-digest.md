---
description: Get a digest of recent activity across multiple Lark IM channels
---

Given the comma-separated channel names provided in $ARGUMENTS (strip leading `#` and whitespace from each):

1. Parse the argument into individual channel names. Strip leading `#` and whitespace from each name.

2. For each channel:
   a. Use `lark_im_search` to find recent messages mentioning the channel topic/name (group chats cannot be resolved by name alone — if you need the exact `chat_id`, list chats via `lark_api` GET `/open-apis/im/v1/chats`). Project results with `jq` under `.data.messages[]` (P3) — pull only sender, timestamp, and snippet.
   b. Use a limit of 50 messages per channel to keep things manageable.
   c. Summarize the key activity in that channel: main topics, decisions, questions, and notable messages.

3. Present the digest in this format:

   ```
   *Channel Digest — <today's date>*

   *#channel-1*
   - Summary point 1
   - Summary point 2

   *#channel-2*
   - Summary point 1
   - Summary point 2

   ...
   ```

4. For each channel, keep the summary to 3-5 bullet points maximum. Focus on what's actionable or noteworthy.

5. If a channel has no recent activity, note that it's been quiet and mention when the last message was posted (if visible).

6. If a channel name can't be found, let the user know and continue with the remaining channels.
