---
name: brand-voice-enforcement
description: >
  This skill applies brand guidelines to content creation. It should be used when
  the user asks to "write an email", "draft a proposal", "create a pitch deck",
  "write a LinkedIn post", "draft a presentation", "write a Lark IM message",
  "draft sales content", or any content creation request where brand voice should
  be applied. Also triggers on "on-brand", "brand voice", "enforce voice",
  "apply brand guidelines", "brand-aligned content", "write in our voice",
  "use our brand tone", "make this sound like us", "rewrite this in our tone",
  or "this doesn't sound on-brand". Not for generating guidelines from scratch
  (use guideline-generation) or discovering brand materials (use discover-brand).
---

# Brand Voice Enforcement

Apply existing brand guidelines to all sales and marketing content generation. Load the user's brand guidelines, apply voice constants and tone flexes to the content request, validate output, and explain brand choices.

> **Lark-native note (collaboration handoff).** The brand-voice methodology below — voice constants, tone flexes, "We Are / We Are Not" — is the specialty and stays unchanged. What changes is *where on-brand content lands*: once written, deliver it through Lark instead of dumping plain text. Resolve recipients with `lark_contact_search` (P1), draft mail with `lark_mail_draft_create`, or post a digest/decision as an interactive card with `lark_im_card_send` (P4). Preview every mutation with `dry_run` first (P2). See `../../connectors/LARK-PATTERNS.md` for the patterns and `../../connectors/LARK-FUSION.md` for delegating to `lark-mail` / `lark-im`.

## Loading Brand Guidelines

Find the user's brand guidelines using this sequence. Stop as soon as you find them:

1. **Session context** — Check if brand guidelines were generated earlier in this session (via `/brand-voice:generate-guidelines`). If so, they are already in the conversation. Use them directly. Session-generated guidelines are the freshest and reflect the user's most recent intent.

2. **Local guidelines file** — Check for `.claude/brand-voice-guidelines.md` inside the user's working folder. Do NOT use a relative path from the agent's current working directory — in Cowork, the agent runs from a plugin cache directory, not the user's project. Resolve the path relative to the user's working folder. If no working folder is set, skip this step.

3. **Lark Wiki (durable team source)** — Guidelines generated previously may have been published to the team's Wiki (the org's system-of-record for knowledge, P8). If the user references a Wiki page or the working folder records a `wiki_node`, fetch it with the `lark-wiki` / `lark-doc` skills (project with `jq` to the guideline body, P3) and use it directly.

4. **Ask the user** — If none of the above found guidelines, tell the user:
   "I couldn't find your brand guidelines. You can:
   - Run `/brand-voice:discover-brand` to find brand materials across your platforms
   - Run `/brand-voice:generate-guidelines` to create guidelines from documents or transcripts
   - Paste guidelines directly into this chat or point me to a file"

   Wait for the user to provide guidelines before proceeding.

Also read `.claude/brand-voice.local.md` for enforcement settings (even if guidelines came from another source):
- `strictness`: strict | balanced | flexible
- `always-explain`: whether to always explain brand choices

## Enforcement Workflow

### 1. Analyze the Content Request

Before writing, identify:
- **Content type**: email, presentation, proposal, social post, message, etc.
- **Target audience**: role, seniority, industry, company stage
- **Key messages needed**: which message pillars apply
- **Specific requirements**: length, format, tone overrides

### 2. Apply Voice Constants

Voice is the brand's personality — it stays constant across all content:
- Apply "We Are / We Are Not" attributes from guidelines
- Use brand personality consistently
- Incorporate approved terminology; reject prohibited terms
- Follow messaging framework and value propositions

Refer to `references/voice-constant-tone-flexes.md` for the "voice constant, tone flexes" model.

### 3. Flex Tone for Context

Tone adapts by content type and audience. Use the tone-by-context matrix from guidelines to set:
- **Formality**: How formal or casual should this be?
- **Energy**: How much urgency or enthusiasm?
- **Technical depth**: How detailed or accessible?

### 4. Generate Content

Create content that:
- Matches brand voice attributes throughout
- Follows tone guidelines for this specific content type
- Incorporates key messages naturally (not forced)
- Uses preferred terminology
- Mirrors the quality and style of guideline examples

For complex or long-form content, delegate to the content-generation agent (defined in `agents/content-generation.md`).
For high-stakes content, delegate to the quality-assurance agent (defined in `agents/quality-assurance.md`) for validation.

### 5. Validate and Explain

After generating content:
- Briefly highlight which brand guidelines were applied
- Explain key voice and tone decisions
- Note any areas where guidelines were adapted for context
- Offer to refine based on feedback

When `always-explain` is true in settings, include brand application notes with every response.

### 6. Deliver into Lark (collaboration layer)

On-brand content is only useful once it reaches someone. When the request is to *send* or *post* (not just draft into the chat), land it natively — never auto-send:

- **Email** (cold outreach, proposals, follow-ups): resolve the recipient with `lark_contact_search(query="<name or email>")` → `lark_mail_draft_create` so the user eyeballs it in the Mail UI before sending. Never call `lark_mail_send` with `confirm_send: true` without explicit confirmation (P2).
- **IM message or announcement**: for a one-liner use `lark_im_send`; for a decision/digest/multi-section announcement post an **interactive card** via `lark_im_card_send` with a YAML `spec` (`print_json: true` → `dry_run: true` → send). Bilingual announcements use `*_i18n` maps (P4).
- **Durable copy** (brand-approved boilerplate, a press release, a long-form page): land it in Wiki/Doc with `lark_wiki_node_create` / `lark_doc_create` rather than leaving it in chat (P8).

Always `dry_run: true` first, show the planned recipient + payload, and only commit after the user confirms (P2). For card grammar or non-trivial mail/doc edits, delegate to the owning skill (`lark-im`, `lark-mail`, `lark-doc`) per `../../connectors/LARK-FUSION.md`.

## Handling Conflicts

When the user's request conflicts with brand guidelines:
1. Explain the conflict clearly
2. Provide a recommendation
3. Offer options: follow guidelines strictly, adapt for context, or override

Default to adapting guidelines with an explanation of the tradeoff.

## Open Questions Awareness

Open questions are unresolved brand positioning decisions flagged during guideline generation, stored in the guidelines under an "Open Questions" section. When generating content, check if the brand guidelines contain open questions:
- If content touches an unresolved open question, note it
- Apply the agent's recommendation from the open question unless the user specifies otherwise
- Suggest resolving the question if it significantly impacts the content

## Reference Files

- **`references/voice-constant-tone-flexes.md`** — The "voice constant, tone flexes" mental model, "We Are / We Are Not" table structure, and tone-by-context matrix explanation
- **`references/before-after-examples.md`** — Before/after content examples per content type showing enforcement in practice
