---
name: ux-copy
description: Write or review UX copy — microcopy, error messages, empty states, CTAs. Trigger with "write copy for", "what should this button say?", "review this error message", or when naming a CTA, wording a confirmation dialog, filling an empty state, or writing onboarding text.
argument-hint: "<context or copy to review>"
---

# /ux-copy

> If you see unfamiliar placeholders or need to check which tools are connected, see [CONNECTORS.md](../../CONNECTORS.md).

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md),
> [LARK-RECIPES](../../../connectors/LARK-RECIPES.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md)).
> The copywriting craft is unchanged; pull context and land outputs in Lark. **Pull brand voice &
> existing terminology** from your knowledge base with `lark_doc_search` / `lark_doc_fetch` (project
> with `jq`, P3) before writing — consistency depends on it. **Screen context & character limits**
> come from the design tool (Figma external MCP). **Land approved copy** in a UX-copy **Lark Base**
> (P5: `lark_base_search` / `lark_base_record_upsert` — `fields:{ Key, Context, Copy, Tone, Status }`;
> scaffold via **base-deploy**) so engineering/loc pull strings from a system-of-record, not chat.
> Document copy patterns/voice in **Wiki** (`lark_wiki_node_create`, P8, via **`lark-doc`**). Share
> copy options for review as an **interactive card** (`lark_im_card_send`, P4).

Write or review UX copy for any interface context.

## Usage

```
/ux-copy $ARGUMENTS
```

## What I Need From You

- **Context**: What screen, flow, or feature?
- **User state**: What is the user trying to do? How are they feeling?
- **Tone**: Formal, friendly, playful, reassuring?
- **Constraints**: Character limits, platform guidelines?

## Principles

1. **Clear**: Say exactly what you mean. No jargon, no ambiguity.
2. **Concise**: Use the fewest words that convey the full meaning.
3. **Consistent**: Same terms for the same things everywhere.
4. **Useful**: Every word should help the user accomplish their goal.
5. **Human**: Write like a helpful person, not a robot.

## Copy Patterns

### CTAs
- Start with a verb: "Start free trial", "Save changes", "Download report"
- Be specific: "Create account" not "Submit"
- Match the outcome to the label

### Error Messages
Structure: What happened + Why + How to fix
- "Payment declined. Your card was declined by your bank. Try a different card or contact your bank."

### Empty States
Structure: What this is + Why it's empty + How to start
- "No projects yet. Create your first project to start collaborating with your team."

### Confirmation Dialogs
- Make the action clear: "Delete 3 files?" not "Are you sure?"
- Describe consequences: "This can't be undone"
- Label buttons with the action: "Delete files" / "Keep files" not "OK" / "Cancel"

### Tooltips
- Concise, helpful, never obvious

### Loading States
- Set expectations, reduce anxiety

### Onboarding
- Progressive disclosure, one concept at a time

## Voice and Tone

Adapt tone to context:
- **Success**: Celebratory but not over the top
- **Error**: Empathetic and helpful
- **Warning**: Clear and actionable
- **Neutral**: Informative and concise

## Output

```markdown
## UX Copy: [Context]

### Recommended Copy
**[Element]**: [Copy]

### Alternatives
| Option | Copy | Tone | Best For |
|--------|------|------|----------|
| A | [Copy] | [Tone] | [When to use] |
| B | [Copy] | [Tone] | [When to use] |
| C | [Copy] | [Tone] | [When to use] |

### Rationale
[Why this copy works — user context, clarity, action-orientation]

### Localization Notes
[Anything translators should know — idioms to avoid, character expansion, cultural context]
```

## Lark-native inputs & delivery

**Knowledge base (Lark Wiki + Docs).** Before writing, pull brand voice guidelines and the content
style guide with `lark_doc_search` → `lark_doc_fetch` (project with `jq`, P3); check existing copy
patterns and terminology standards so the same terms are used everywhere (principle: Consistent).

**Design tool (Figma, external — keep as-is):** view the screen context to understand the full user
flow; read character limits and layout constraints from the design. Stays external.

**Store approved copy in a UX-copy Lark Base (P5 — system-of-record).** Read existing strings with
`lark_base_search` (pass `search_fields` for the column(s) to match — e.g. `Key` — and narrow with
`select_fields`/`limit`; `lark_base_search` does NOT support `jq`. Discover field names via `lark_api
GET /open-apis/bitable/v1/apps/{base}/tables/{table}/fields` if unknown); write/update each with
`lark_base_record_upsert` (`fields:{ Key, Context, Copy,
Tone, "Char limit", Status, "Loc notes" }`, `dry_run` first, P2). This is the single source loc and
engineering pull from. No Base yet → **delegate to `base-deploy`**.

**Document patterns & voice in Wiki (P8).** Publish the voice/tone guide and reusable patterns with
`lark_wiki_node_create` → fill via **`lark-doc`**.

**Share copy options for review as an interactive card (P4).** `lark_im_card_send` (`print_json` →
`dry_run` → send): `item`/`select` rows for options A/B/C with a side *Pick* button, `note` footer with
rationale. For bilingual strings, use `*_i18n` maps so the card renders per-locale. Card grammar →
**`lark-im`** skill.

## Tips

1. **Be specific about context** — "Error message when payment fails" is better than "error message."
2. **Share your brand voice** — "We're professional but warm" helps me match your tone.
3. **Consider the user's emotional state** — Error messages need empathy. Success messages can celebrate.
