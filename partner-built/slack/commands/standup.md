---
description: Generate a standup update based on your recent Lark IM activity
---

1. Use `lark_contact_search` with `user_ids: "me"` to get the current user's profile information, including their `open_id` and display name. Project under `.data.users[]` (P3).

2. Search for the user's recent messages using `lark_im_search`, filtered to the resolved `open_id` as sender and a date range covering yesterday onward. This captures messages from the last working day. Project results with `jq` under `.data.messages[]`.

3. Review the messages found and categorize them into standup themes:
   - **What I worked on** — Topics, projects, or tasks the user discussed or contributed to
   - **What I'm working on next** — Any mentions of upcoming work, plans, or follow-ups
   - **Blockers** — Any questions asked that went unanswered, issues raised, or explicit mentions of being stuck

4. For messages in threads, follow up with `lark_im_search` (narrow to the surrounding thread) to get the full context so you can accurately describe what the user contributed.

5. Format the standup as:
   ```
   *Standup for <display name> — <today's date>*

   *Done:*
   - Item 1
   - Item 2

   *Doing:*
   - Item 1

   *Blockers:*
   - None / Item 1
   ```

6. Present the standup to the user for review. They can edit, adjust, or ask you to post it to a specific channel.
