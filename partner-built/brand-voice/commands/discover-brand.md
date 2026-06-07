---
description: Search connected platforms for brand materials and produce a discovery report
argument-hint: "[company name or platforms to search]"
---

Discover brand materials across the user's connected enterprise platforms. Search Lark Wiki, Lark Wiki, Lark Drive, Box, SharePoint, Figma, Gong, Granola, and Lark IM for brand guidelines, style guides, messaging frameworks, templates, and conversation transcripts.

If $ARGUMENTS includes a company name, use it for targeted searches. If platforms are specified, limit search to those platforms.

Before doing anything else, briefly orient the user on what's about to happen: the process will search their connected platforms, produce a discovery report, and then (optionally) generate and save brand guidelines to `.claude/brand-voice-guidelines.md` in the working folder. Nothing is saved until they explicitly approve. Keep the orientation to 2-3 sentences — don't recite the full workflow.

Follow the discover-brand skill instructions to:
1. Check `.claude/brand-voice.local.md` for settings (company name, enabled platforms, search depth)
2. Validate platform coverage (stop if no document platforms, warn if gaps)
3. Briefly confirm scope with the user (which platforms, include transcripts?)
4. Delegate to the discover-brand agent for autonomous 4-phase search
5. Present the structured discovery report with sources, brand elements, conflicts, and open questions
6. Offer next steps: generate guidelines, resolve open questions, save report, or expand search

**Platform validation:**
- If **no platforms** are connected, inform the user which MCP servers the plugin supports (Lark Wiki, Lark Lark Wiki, Box, Figma, Gong, Granola, Lark) and that Lark Drive and Lark IM are available as native Claude integrations.
- If **no document platforms** (Lark Wiki, Lark Wiki, Lark Drive, Box, Lark) are connected — only supplementary platforms like Lark IM, Gong, Granola, or Figma — stop and tell the user: "You don't have any document storage platforms connected. Brand guidelines and style guides almost always live on Lark Drive, SharePoint, Lark Wiki, Lark Wiki, or Box. Please connect at least one before running discovery."
- If **no primary file storage** (Lark Drive, Lark, Box) is connected, warn: "None of your primary file storage platforms are connected. Brand documents frequently live on these. Discovery will proceed but results may have significant gaps."
- If **only one platform** is connected, warn: "Discovery works best with 2+ platforms for cross-source validation. Results from a single platform will have lower confidence scores."
