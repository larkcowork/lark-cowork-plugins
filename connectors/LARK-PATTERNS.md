# Lark-Native Workflow Patterns

These are the patterns that make a plugin feel **built for Lark**, not just pointed at it. Every
soul-swapped skill should compose from these. They exploit what's distinctive about Lark —
interactive cards, Base as a database, Minutes AI artifacts, approval flows — instead of treating
Lark as a generic chat/doc backend. Recipes referenced here live in
[LARK-RECIPES.md](./LARK-RECIPES.md); deep-skill delegation in [LARK-FUSION.md](./LARK-FUSION.md).

---

## P1 — Resolve identity first

`open_id`s are opaque and cannot be guessed. Before any send / invite / assign / mention:
`lark_contact_search(query="<name or email>")` → take `open_id` → use it. `user_ids: "me"` returns
the caller. Skipping this is the #1 cause of failed writes.

## P2 — Safe mutation (preview → confirm → commit)

Every mutating tool exposes `dry_run` (and cards add `print_json`). The contract:

1. Build the call, run it with `dry_run: true`.
2. Show the user the planned recipient + payload.
3. Re-issue **without** `dry_run` only after the user confirms.

Mail is doubly guarded: `lark_mail_send` without `confirm_send: true` only drafts. Prefer
`lark_mail_draft_create` whenever a human should eyeball it in the Mail UI first. Never auto-send.

## P3 — Token economy (jq projection)

Read tools (`*_search`, `*_my`, `*_fetch`, `*_read`, agenda) accept `jq`. Always project to the
fields you need: a full doc fetch can be 25k tokens; `jq: ".title"` or `jq: ".content[:4000]"`
makes it ~500. Default to a projection; only pull full payloads when you'll use them.

## P4 — Interactive cards (the Lark superpower)

This is what makes "Cowork living inside Lark" magical: don't report with plain text — post an
**interactive card** the user acts on from chat. Use `lark_im_card_send` with a YAML `spec`;
always `print_json: true` to validate the spec offline, then `dry_run: true`, then send.

Use a card (not `lark_im_send`) whenever the output is a **decision or a digest**:

- **Approve / triage** — `actions` row with buttons (Approve / Reject / Ask more); `item` rows with
  a side button per entry (approve each expense / triage each ticket).
- **Daily / pipeline / status brief** — `header` (colored `template`), `panel` collapsible sections,
  `div` rows with status pills (icon + text), `note` footer.
- **Bilingual announcements** — `*_i18n` maps; `config.locales` auto-derives.

Card spec grammar + templates: the **`lark-im`** skill, `references/spec-reference.md`. Element
kinds available: `md, hr, note, div, actions, item, select, panel, columns, image, raw`.

```yaml
# minimal triage card skeleton (pass as `spec`) — VERIFIED to compile via
#   lark-cli im +card-send --print-json  (offline; no API/auth)
# Grammar: each element is keyed by its KIND (`- md:`, `- item:`, `- actions:`),
# NOT `{kind: ...}`. A button's `value` is a STRING (callback id) — or use `url`.
header: { title: "3 tickets need triage", template: orange }
elements:
  - item: { text: "#1042 — refund stuck (P1)", button: { text: "Escalate", type: danger, value: "escalate_1042" } }
  - item: { text: "#1043 — login bug (P2)",    button: { text: "Take",     type: primary, value: "take_1043" } }
  - actions:
      - { text: "Triage all", type: primary, value: "triage_all" }
      - { text: "Snooze",     type: default, value: "snooze" }
```

> Verified card grammar (binary is source of truth): elements use `- <kind>:` where `<kind>` ∈
> `md, hr, note, div, actions, item, select, panel, columns, image, raw`; button `value` must be a
> string (callback) or use `url` (link). Always `print_json: true` to validate before sending.

## P5 — Base as system-of-record

Lark has no separate CRM/tracker product — a **Lark Base (Bitable)** is the database. Pattern:

- **Read** with `lark_base_search` (project with `jq`; use `search_json` for field/view control).
- **Write** with the bitable record recipes (`lark_api`) — see LARK-RECIPES.
- **Scaffold** a new tracker/CRM/decision-log with the **`base-deploy`** skill (don't hand-roll
  schema). Treat the Base as the source of truth; chat cards and docs are views over it.

This is how `sales` gets a real CRM, `customer-support` a ticket/escalation log, `product-management`
a feedback tracker, `operations` an asset/vendor register — all native, all queryable.

## P6 — Minutes → action items → tasks

Lark Minutes already produces AI summaries + action items. Don't re-summarize from scratch:

1. `lark_minutes_search(query=..., participant_ids="me")` → find the meeting, pull its AI artifacts.
2. Extract decisions + action items.
3. For each action item: `lark_contact_search` the owner → `lark_task_create(dry_run)` → confirm → commit.
4. Optionally log decisions to a Base (P5) and post a recap **card** (P4).

`lark_vc_search` gives meeting metadata (who/when/room) when you only need the envelope, not content.

## P7 — Approval flows

For anything that needs a human gate (expense, contract, access, escalation), create a real Lark
**approval instance** (`lark_api`, see LARK-RECIPES / `lark-approval` skill) instead of an ad-hoc
"reply yes". The card (P4) kicks it off; the approval engine tracks state and notifies approvers.

## P8 — Knowledge to Wiki, drafts to Drive

Durable knowledge (KB articles, runbooks, postmortems, specs) belongs in **Wiki** (P-recipe wiki
node create → `lark_doc_create` with `wiki_node`/`wiki_space`). Transient artifacts and uploads go
to **Drive** (`lark_drive_upload`). Don't leave deliverables as chat messages — land them where the
org will find them later.

---

### How to apply in a skill

A soul-swapped skill should read like: *resolve people (P1) → read state with projection (P3),
Base as SoR (P5) / Minutes (P6) → do the work → land output where it belongs (P8) → surface a
decision/digest as an interactive card (P4) → gate with approval if needed (P7) → mutate safely
(P2)*. Delegate domain depth to the owning skill ([LARK-FUSION.md](./LARK-FUSION.md)).
