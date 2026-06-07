---
name: design-critique
description: Get structured design feedback on usability, hierarchy, and consistency. Trigger with "review this design", "critique this mockup", "what do you think of this screen?", or when sharing a Figma link or screenshot for feedback at any stage from exploration to final polish.
argument-hint: "<Figma URL, screenshot, or description>"
---

# /design-critique

> If you see unfamiliar placeholders or need to check which tools are connected, see [CONNECTORS.md](../../CONNECTORS.md).

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md),
> [LARK-RECIPES](../../../connectors/LARK-RECIPES.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md)).
> The critique framework is unchanged; only inputs/outputs are Lark. The design itself stays in your
> design tool (Figma external MCP) — pull components, tokens, layers from there. Optionally enrich
> with real user feedback (Intercom external, or a feedback **Lark Base** via `lark_base_search`, P5).
> Deliver the critique as an **interactive card** (`lark_im_card_send`, P4) the team can act on, land
> the durable write-up in **Wiki** (`lark_wiki_node_create`, P8), and turn agreed fixes into **Lark
> Tasks** (`lark_task_create` — resolve owners with `lark_contact_search` first, P1; `dry_run`, P2).

Get structured design feedback across multiple dimensions.

## Usage

```
/design-critique $ARGUMENTS
```

Review the design: @$1

If a Figma URL is provided, pull the design from Figma. If a file is referenced, read it. Otherwise, ask the user to describe or share their design.

## What I Need From You

- **The design**: Figma URL, screenshot, or detailed description
- **Context**: What is this? Who is it for? What stage (exploration, refinement, final)?
- **Focus** (optional): "Focus on mobile" or "Focus on the onboarding flow"

## Critique Framework

### 1. First Impression (2 seconds)
- What draws the eye first? Is that correct?
- What's the emotional reaction?
- Is the purpose immediately clear?

### 2. Usability
- Can the user accomplish their goal?
- Is the navigation intuitive?
- Are interactive elements obvious?
- Are there unnecessary steps?

### 3. Visual Hierarchy
- Is there a clear reading order?
- Are the right elements emphasized?
- Is whitespace used effectively?
- Is typography creating the right hierarchy?

### 4. Consistency
- Does it follow the design system?
- Are spacing, colors, and typography consistent?
- Do similar elements behave similarly?

### 5. Accessibility
- Color contrast ratios
- Touch target sizes
- Text readability
- Alternative text for images

## How to Give Feedback

- **Be specific**: "The CTA competes with the navigation" not "the layout is confusing"
- **Explain why**: Connect feedback to design principles or user needs
- **Suggest alternatives**: Don't just identify problems, propose solutions
- **Acknowledge what works**: Good feedback includes positive observations
- **Match the stage**: Early exploration gets different feedback than final polish

## Output

```markdown
## Design Critique: [Design Name]

### Overall Impression
[1-2 sentence first reaction — what works, what's the biggest opportunity]

### Usability
| Finding | Severity | Recommendation |
|---------|----------|----------------|
| [Issue] | 🔴 Critical / 🟡 Moderate / 🟢 Minor | [Fix] |

### Visual Hierarchy
- **What draws the eye first**: [Element] — [Is this correct?]
- **Reading flow**: [How does the eye move through the layout?]
- **Emphasis**: [Are the right things emphasized?]

### Consistency
| Element | Issue | Recommendation |
|---------|-------|----------------|
| [Typography/spacing/color] | [Inconsistency] | [Fix] |

### Accessibility
- **Color contrast**: [Pass/fail for key text]
- **Touch targets**: [Adequate size?]
- **Text readability**: [Font size, line height]

### What Works Well
- [Positive observation 1]
- [Positive observation 2]

### Priority Recommendations
1. **[Most impactful change]** — [Why and how]
2. **[Second priority]** — [Why and how]
3. **[Third priority]** — [Why and how]
```

## Lark-native delivery & inputs

**Design tool (Figma, external — keep as-is):** pull the design directly from the Figma MCP, inspect
components/tokens/layers, and compare against the design system. This stays external.

**User feedback (external or Lark Base, P5):** cross-reference critique points with recent user
complaints — from Intercom (external MCP) if connected, or from a feedback **Lark Base** via
`lark_base_search` (pass `search_fields` for the column(s) to match; `lark_base_search` does NOT
support `jq`, so narrow with `select_fields`/`limit` — discover field names via `lark_api GET
/open-apis/bitable/v1/apps/{base}/tables/{table}/fields` if unknown). Grounding "the CTA competes
with nav" in a real ticket makes the feedback land.

**Deliver the critique as an interactive card (P4).** Instead of a wall of text, post with
`lark_im_card_send` (`print_json: true` to validate the `spec`, then `dry_run: true`, then send): a
`header` with the design name + stage, `panel` sections for Usability / Hierarchy / Consistency, status
pills per finding, and an `actions` footer (*Open in Wiki* / *Create fix tasks*). Card grammar →
**`lark-im`** skill.

**Land the durable write-up (P8).** If the critique should persist (design review record), publish it
to **Wiki** with `lark_wiki_node_create` and fill via the **`lark-doc`** skill.

**Turn priority recommendations into Lark Tasks.** For each agreed fix: `lark_contact_search` the
owner → `open_id` (P1) → `lark_task_create` (`dry_run` first, P2). Delegate richer task ops to
**`lark-task`**.

## Tips

1. **Share the context** — "This is a checkout flow for a B2B SaaS" helps me give relevant feedback.
2. **Specify your stage** — Early exploration gets different feedback than final polish.
3. **Ask me to focus** — "Just look at the navigation" gives you more depth on one area.
