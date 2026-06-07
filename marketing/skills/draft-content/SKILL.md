---
name: draft-content
description: Draft blog posts, social media, email newsletters, landing pages, press releases, and case studies with channel-specific formatting and SEO recommendations. Use when writing any marketing content, when you need headline or subject line options, or when adapting a message for a specific platform, audience, and brand voice.
argument-hint: "<content type and topic>"
---

# Draft Content

> If you see unfamiliar placeholders or need to check which tools are connected, see [CONNECTORS.md](../../CONNECTORS.md).

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md), [LARK-RECIPES](../../../connectors/LARK-RECIPES.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md)). The writing frameworks live in the `content-creation` skill; this command makes them concrete and Lark-native:
> - **Resolve brand voice + reusable facts from Lark first** (P3): `lark_doc_search`/`lark_doc_fetch` for the style guide and prior content; `lark_base_search` for product facts/proof points (requires `search_fields`; discover field names via `lark_api GET /open-apis/bitable/v1/apps/{base}/tables/{table}/fields` if unknown; narrow with `select_fields`/`limit`, not jq).
> - **Web content (blog, landing page, case study, press release) → Wiki/Docs** as the deliverable, not a chat paste (P8): `lark_wiki_node_create` then fill via the **`lark-doc`** skill (DocxXML). Exports/PDFs → `lark_drive_upload`.
> - **Email newsletter copy → a real Lark Mail draft** (P2): `lark_mail_draft_create` so a human eyeballs it in the Mail UI; never `lark_mail_send` with `confirm_send` unsolicited.
> - **Send for review as a card** (P4): post the draft link + a short review card via `lark_im_card_send` to the reviewer (resolve via `lark_contact_search`, P1); `print_json` → `dry_run` → send.

Generate marketing content drafts tailored to a specific content type, audience, and brand voice.

## Trigger

User runs `/draft-content` or asks to draft, write, or create marketing content.

## Inputs

Gather the following from the user. If not provided, ask before proceeding:

1. **Content type** — one of:
   - Blog post
   - Social media post (specify platform: LinkedIn, Twitter/X, Instagram, Facebook)
   - Email newsletter
   - Landing page copy
   - Press release
   - Case study

2. **Topic** — the subject or theme of the content

3. **Target audience** — who this content is for (role, industry, seniority, pain points)

4. **Key messages** — 2-4 main points or takeaways to communicate

5. **Tone** — e.g., authoritative, conversational, inspirational, technical, witty (optional if brand voice is configured)

6. **Length** — target word count or format constraint (e.g., "1000 words", "280 characters", "3 paragraphs")

## Brand Voice

- First resolve the documented voice from Lark: `lark_doc_search(query="brand voice"/"style guide")` → `lark_doc_fetch` (jq-projected, P3). This is preferred over local settings.
- If a brand voice is configured in the user's local settings file, apply it. Inform the user which source is being applied.
- If neither exists, ask: "Do you have brand voice guidelines you'd like me to follow (paste, or point me at a Lark Wiki/Doc)? If not, I'll use a neutral professional tone."
- Apply the specified or default tone consistently throughout the draft.

## Content Generation by Type

### Blog Post
- Engaging headline (provide 2-3 options)
- Introduction with a hook (question, statistic, bold statement, or story)
- 3-5 organized sections with descriptive subheadings
- Supporting points, examples, or data references in each section
- Conclusion with a clear call to action
- SEO considerations: suggest a primary keyword, include it in the headline and first paragraph, use related keywords in subheadings

### Social Media Post
- Platform-appropriate format and length
- Hook in the first line
- Hashtag suggestions (3-5 relevant hashtags)
- Call to action or engagement prompt
- Emoji usage appropriate to brand and platform
- If LinkedIn: professional framing, paragraph breaks for readability
- If Twitter/X: concise, punchy, within character limit
- If Instagram: visual-first language, story-driven, hashtag block

### Email Newsletter
- Subject line (provide 2-3 options with open-rate considerations)
- Preview text
- Greeting
- Body sections with clear hierarchy
- Call to action button text
- Sign-off
- Unsubscribe note reminder

### Landing Page Copy
- Headline and subheadline
- Hero section copy
- Value propositions (3-4 benefit-driven bullets or sections)
- Social proof placeholder (suggest testimonial or stat placement)
- Primary and secondary CTAs
- FAQ section suggestions
- SEO: meta title and meta description suggestions

### Press Release
- Headline following press release conventions
- Dateline and location
- Lead paragraph (who, what, when, where, why)
- Supporting quotes (provide placeholder guidance)
- Company boilerplate placeholder
- Media contact placeholder
- Standard press release formatting

### Case Study
- Title emphasizing the result
- Customer overview (industry, size, challenge)
- Challenge section
- Solution section (what was implemented)
- Results section with metrics (prompt user for data)
- Customer quote placeholder
- Call to action

## SEO Considerations (for web content)

For blog posts, landing pages, and other web-facing content:
- Suggest a primary keyword based on the topic
- Recommend keyword placement: headline, first paragraph, subheadings, meta description
- Suggest internal and external linking opportunities
- Recommend a meta description (under 160 characters)
- Note image alt text opportunities

## Output

Present the draft with clear formatting. After the draft, include:
- A brief note on what brand voice source and tone were applied (Lark doc / local settings / default)
- Any SEO recommendations (for web content)
- Suggestions for next steps and offer to land it natively: publish web content to **Wiki/Docs** (P8), create the email as a **Lark Mail draft** (P2), and send a **review card** to a teammate (P4, resolve via `lark_contact_search`).

Ask: "Would you like me to publish this to Wiki / create a mail draft, send it to a teammate for review, revise any section, adjust the tone, or create a variation for a different channel?"
