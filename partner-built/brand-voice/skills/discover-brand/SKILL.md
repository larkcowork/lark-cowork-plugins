---
name: discover-brand
description: >
  This skill orchestrates autonomous discovery of brand materials across the org's
  Lark collaboration layer (Lark Wiki, Lark Drive, Lark IM) plus specialty connectors
  (Figma design tokens, Gong and Granola call transcripts).
  It should be used when the user asks to "discover brand materials",
  "find brand documents", "search for brand guidelines", "audit brand content",
  "what brand materials do we have", "find our style guide", "where are our brand docs",
  "do we have a style guide", "discover brand voice", "brand content audit",
  or "find brand assets".
---

# Brand Discovery

Orchestrate autonomous discovery of brand materials across enterprise platforms. This skill coordinates the discover-brand agent to search connected platforms, triage sources, and produce a structured discovery report with open questions.

> **Lark-native note (collaboration layer).** Brand docs, decks, templates, and informal voice patterns inside the org live in **Lark Wiki / Lark Drive / Lark IM** — search those with the curated `lark_*` tools rather than treating Lark as a generic chat/doc backend. Specialty sources stay on their own connectors: **Figma** (design tokens), **Gong** and **Granola** (conversation intelligence) remain external MCP servers and are searched as-is. When the discovery report is durable, publish it to Wiki (P8). Project every read with `jq` (P3). See `../../connectors/LARK-PATTERNS.md` and `../../connectors/LARK-FUSION.md`.

## Discovery Workflow

### 0. Orient the User

Before starting, briefly explain what's about to happen so the user knows what to expect:

"Here's how brand discovery works:

1. **Search** — I'll search your Lark Wiki, Lark Drive, and Lark IM for brand-related materials (style guides, pitch decks, templates), plus any connected specialty sources — Figma design systems and Gong/Granola call transcripts.
2. **Analyze** — I'll categorize and rank what I find, pull the best sources, and produce a discovery report with what I found, any conflicts, and open questions.
3. **Generate guidelines** — Once you've reviewed the report, I can generate a structured brand voice guideline document from the results.
4. **Save** — Guidelines are saved to `.claude/brand-voice-guidelines.md` in your working folder (and optionally published to Lark Wiki) once you approve them. Nothing is written until that step.

The search usually takes a few minutes depending on how many platforms are connected. Ready to get started?"

Wait for the user to confirm before proceeding. If they have questions about the process, answer them first.

### 1. Check Settings

Read `.claude/brand-voice.local.md` if it exists. Extract:
- Company name
- Which sources are enabled. Collaboration sources resolve to Lark (Lark Wiki, Lark Drive, Lark IM); specialty sources stay external (figma, gong, granola).
- Search depth preference (standard or deep)
- Max sources limit
- Any known brand material locations listed under "Known Brand Materials" (e.g. a Wiki node token or Drive folder)

If no settings file exists, proceed with all connected sources and standard search depth.

### 2. Validate Platform Coverage

Before confirming scope, check which sources are actually connected and classify them:

**Document sources** (where brand guidelines, style guides, templates, and decks live):
- Lark Wiki, Lark Drive (search with `lark_api` wiki/drive search; the `lark-wiki` / `lark-drive` skills own the exact queries)

**Supplementary sources** (valuable for patterns, but not where brand docs are stored):
- Lark IM (informal voice patterns, pinned guidelines — `lark_im_search`), plus the specialty external connectors: Gong / Granola (call transcripts) and Figma (design tokens)

Apply these rules:

1. **If no Lark document source is reachable (Wiki/Drive)**: **Stop.** Tell the user: "I can't reach your Lark Wiki or Lark Drive. Brand guidelines and style guides almost always live on one of these. Please make sure the Lark connector is authenticated before running discovery. Gong/Granola transcripts and Lark IM threads are valuable supplements but unlikely to contain formal brand documents."

2. **If Lark Drive search returns nothing but Wiki is reachable (or vice-versa)**: **Warn** (but proceed): "Only one of your Lark document stores returned results. Brand documents may be split across Wiki and Drive; discovery will proceed but results may have gaps."

3. **If only one source total is connected**: **Warn** (but proceed): "Only [source] is connected. Discovery works best with 2+ sources for cross-source validation. Results from a single source will have lower confidence scores."

### 3. Confirm Scope with User

Before launching discovery, confirm:
- Which sources to search (default: all connected — Lark Wiki/Drive/IM + any specialty connectors)
- Whether to include conversation transcripts (Gong, Granola, or Lark Minutes/VC) or just documents
- Any known locations to prioritize (Wiki node token, Drive folder, IM chat)

Keep this brief — one question, not a questionnaire.

### 4. Delegate to Discover-Brand Agent

Launch the discover-brand agent via the Task tool. Provide:
- Company name (from settings or user input)
- Enabled sources (Lark Wiki/Drive/IM as the document/voice backbone; Gong/Granola/Figma as specialty supplements)
- Search depth
- Any known Wiki node tokens, Drive folders, or IM chats to check first

For Lark sources the agent uses curated reads — `lark_im_search` for chat, and `lark_api` wiki/drive search (delegating exact queries to the `lark-wiki` / `lark-drive` skills) — and projects each result with `jq` (P3) so triage stays cheap.

The agent executes the 4-phase discovery algorithm autonomously:
1. **Broad Discovery** — parallel searches across platforms
2. **Source Triage** — categorize and rank sources
3. **Deep Fetch** — retrieve and extract from top sources
4. **Discovery Report** — structured output with open questions

### 5. Present Discovery Report

When the agent returns, present the report to the user with a summary:
- Total sources found and analyzed
- Key brand elements discovered
- Any conflicts between sources
- Open questions requiring team input

### 6. Offer Next Steps

After presenting the report, offer:
1. **Generate guidelines now** — chain to `/brand-voice:generate-guidelines` using discovery report as input
2. **Resolve open questions first** — work through high-priority questions before generating
3. **Save report** — publish the discovery report as a durable team artifact in Lark Wiki via `lark_wiki_node_create` (P8, dry-run first per P2), or keep it as a local file
4. **Expand search** — search additional sources or deeper if coverage is low

## Open Questions

Open questions arise when the discovery agent encounters ambiguity it cannot resolve:
- Conflicting documents (e.g., 2023 style guide vs. 2024 brand update)
- Missing critical sections (e.g., no social media guidelines found)
- Inconsistent terminology across platforms

Every open question includes an agent recommendation. Present questions as "confirm or override" — not dead ends.

## Integration with Other Skills

- **Guideline Generation**: The discovery report is returned by the discover-brand agent via the Task tool. Pass it directly to the guideline-generation skill as structured input, replacing the need for users to manually gather sources.
- **Brand Voice Enforcement**: Once guidelines are generated from discovery, enforcement uses them automatically.

## Error Handling

- If no source is connected, inform the user that the Lark connector (Wiki/Drive/IM) is the discovery backbone and that Figma/Gong/Granola are optional specialty supplements, and how to connect them.
- If all searches return empty results, flag the discovery as "low coverage" and suggest the user provide documents manually or check source connections.
- If a source is connected but returns permission errors, note the gap and continue with other sources. For Lark scope/permission errors, the error envelope may suggest retrying with bot identity — follow it.

## Reference Files

For detailed discovery patterns and algorithms, consult:

- **`references/search-strategies.md`** — Platform-specific search queries, query patterns by platform, and tips for maximizing discovery coverage
- **`references/source-ranking.md`** — Source category definitions, ranking algorithm weights, and triage decision criteria
