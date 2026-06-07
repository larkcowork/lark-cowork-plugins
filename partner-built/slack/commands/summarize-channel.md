---
description: Summarize recent activity in a Lark IM channel
---

Given the channel name provided in $ARGUMENTS (strip any leading `#`):

1. Strip any leading `#` from the argument. Use `lark_im_search` to find recent messages for the channel topic/name (group chats cannot be resolved by name alone — if you need the exact `chat_id`, list chats via `lark_api` GET `/open-apis/im/v1/chats`). Project results with `jq` under `.data.messages[]` (P3).
2. Read recent messages from the channel (default limit of 100 messages).
3. For any messages that have threads with replies, follow up with `lark_im_search` (narrow to the surrounding thread) so the summary captures threaded discussions.
4. Produce a concise summary organized by topic or theme. The summary should include:
   - An overview of the main topics discussed
   - Key decisions or action items mentioned
   - Notable announcements or updates
   - Active threads and their conclusions (if any)
5. Keep the summary scannable — use short bullet points grouped by topic. Mention who said what when it's relevant (e.g., decisions, action items).
6. If the channel has very little recent activity, say so and note the last time a message was posted.
