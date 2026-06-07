---
name: slack-messaging
description: Guidance for composing well-formatted, effective Lark IM messages using mrkdwn syntax
---

# Lark IM Messaging Best Practices

This skill provides guidance for composing well-formatted, effective Lark IM messages. The Lark
connector supersedes Slack here: route all sends, drafts, and digests through the curated `lark_*`
tools below. For card grammar and message-formatting depth, delegate to the **`lark-im`** skill.

> Depth core (read for any non-trivial flow): `../../connectors/LARK-PATTERNS.md` (P1 identity,
> P2 safe mutation, P3 token economy, P4 interactive cards), `../../connectors/LARK-RECIPES.md`,
> `../../connectors/LARK-FUSION.md` (delegate to `lark-im`).

## When to Use

Apply this skill whenever composing, drafting, or helping the user write a Lark IM message — whether
posting plain text, drafting for review, or surfacing a decision/digest as an interactive card.

## Tooling — what to call

| Goal | Tool | Notes |
|------|------|-------|
| Resolve a recipient (name/email → `open_id`) | `lark_contact_search` | P1 — always do this first; IDs cannot be guessed. `user_ids: "me"` is the caller. |
| Send a plain text / mrkdwn message | `lark_im_send` | `dry_run: true` first, show the planned recipient + body, then re-send to commit (P2). |
| Surface a decision or digest | `lark_im_card_send` | Use an interactive **card** (P4), not plain text — see below. |
| Find context before replying | `lark_im_search` | Project with `jq` (P3) to keep token cost down. |

Anything without a curated tool goes through `lark_api` (Lark OpenAPI passthrough). For card spec
YAML grammar, threading, and mrkdwn edge cases, defer to the **`lark-im`** skill rather than
hand-rolling.

## Lark IM Formatting (mrkdwn)

Lark IM uses its own markup syntax called **mrkdwn**, which differs from standard Markdown. Always use mrkdwn when composing Lark IM messages:

| Format | Syntax | Notes |
|--------|--------|-------|
| Bold | `*text*` | Single asterisks, NOT double |
| Italic | `_text_` | Underscores |
| Strikethrough | `~text~` | Tildes |
| Code (inline) | `` `code` `` | Backticks |
| Code block | `` ```code``` `` | Triple backticks |
| Quote | `> text` | Angle bracket |
| Link | `<url\|display text>` | Pipe-separated in angle brackets |
| User mention | `<at user_id="ou_...">name</at>` | Resolve the `open_id` via `lark_contact_search` first |
| Bulleted list | `- item` or `• item` | Dash or bullet character |
| Numbered list | `1. item` | Number followed by period |

### Common Mistakes to Avoid

- Do NOT use `**bold**` (double asterisks) — Lark IM uses `*bold*` (single asterisks)
- Do NOT use `## headers` — Lark IM does not support Markdown headers. Use `*bold text*` on its own line instead.
- Do NOT use `[text](url)` for links — Lark IM uses `<url|text>` format
- Do NOT use `---` for horizontal rules — Lark IM does not render these
- Do NOT guess an `open_id` for a mention — resolve it via `lark_contact_search` (P1)

## When to use a card instead of a message (P4)

A plain `lark_im_send` is fine for a one-line reply. But whenever the output is a **decision or a
digest** — an announcement, a status brief, an approve/triage prompt — post an **interactive card**
via `lark_im_card_send` instead. Always `print_json: true` to validate the spec offline, then
`dry_run: true`, then send. Card spec grammar (header templates, `panel`, `item`+button rows,
`actions`, `*_i18n`) lives in the **`lark-im`** skill, `references/spec-reference.md`.

```yaml
# minimal announcement card (pass as `spec` to lark_im_card_send)
header: { title: "Q3 planning kickoff", template: blue }
elements:
  - md: "*When:* Mon 10:00 · *Where:* VC room 3"
  - note: "Reply in thread with blockers."
```

## Message Structure Guidelines

- **Lead with the point.** Put the most important information in the first line. Many people read Lark IM on mobile or in notifications where only the first line shows.
- **Keep it short.** Aim for 1-3 short paragraphs. If the message is long, consider a durable doc (`lark_doc_create` / `lark_wiki_node_create`, P8) and link it instead of a wall of text.
- **Use line breaks generously.** Walls of text are hard to read. Separate distinct thoughts with blank lines.
- **Use bullet points for lists.** Anything with 3+ items should be a list, not a run-on sentence.
- **Bold key information.** Use `*bold*` for names, dates, deadlines, and action items so they stand out when scanning.

## Thread vs. Channel Etiquette

- **Reply in threads** when responding to a specific message to keep the main group chat clean.
- **Post to the whole group** (not a thread) when starting a new topic, making an announcement, or asking a question everyone needs to see.
- **Don't start a new thread** to continue an existing conversation — find and reply to the original message.

## Tone and Audience

- Match the tone to the group — a company-wide announcement chat is usually more formal than a team banter chat.
- For simple acknowledgments, prefer an emoji reaction over a reply message. (The `lark-im` skill covers reaction management if you need to apply one programmatically.)
- When writing announcements, use a clear structure: context, key info, call to action — and surface it as a card (P4).
- Drafting for human review? For longer or higher-stakes outreach, draft via `lark_mail_draft_create` so it can be eyeballed in the Mail UI before anything is sent (P2). Never auto-send.
