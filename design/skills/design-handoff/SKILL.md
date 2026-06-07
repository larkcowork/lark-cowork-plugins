---
name: design-handoff
description: Generate developer handoff specs from a design. Use when a design is ready for engineering and needs a spec sheet covering layout, design tokens, component props, interaction states, responsive breakpoints, edge cases, and animation details.
argument-hint: "<Figma URL or design description>"
---

# /design-handoff

> If you see unfamiliar placeholders or need to check which tools are connected, see [CONNECTORS.md](../../CONNECTORS.md).

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md),
> [LARK-RECIPES](../../../connectors/LARK-RECIPES.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md)).
> The spec content is unchanged; the deliverable becomes a durable Lark artifact, not a chat dump.
> Pull exact measurements/tokens/component props from your design tool (Figma external MCP), export
> assets, then: **land the handoff doc in Wiki** (`lark_wiki_node_create`, P8) so engineering finds it
> later; **upload exported assets** to Drive (`lark_drive_upload`); **create an implementation Lark
> Task** per spec section (`lark_task_create` — resolve the eng owner via `lark_contact_search`, P1,
> `dry_run` first, P2; delegate subtasks to **`lark-task`**); and **notify engineering with an
> interactive card** (`lark_im_card_send`, P4) linking the Wiki page + tasks.

Generate comprehensive developer handoff documentation from a design.

## Usage

```
/design-handoff $ARGUMENTS
```

Generate handoff specs for: @$1

If a Figma URL is provided, pull the design from Figma. Otherwise, work from the provided description or screenshot.

## What to Include

### Visual Specifications
- Exact measurements (padding, margins, widths)
- Design token references (colors, typography, spacing)
- Responsive breakpoints and behavior
- Component variants and states

### Interaction Specifications
- Click/tap behavior
- Hover states
- Transitions and animations (duration, easing)
- Gesture support (swipe, pinch, long-press)

### Content Specifications
- Character limits
- Truncation behavior
- Empty states
- Loading states
- Error states

### Edge Cases
- Minimum/maximum content
- International text (longer strings)
- Slow connections
- Missing data

### Accessibility
- Focus order
- ARIA labels and roles
- Keyboard interactions
- Screen reader announcements

## Principles

1. **Don't assume** — If it's not specified, the developer will guess. Specify everything.
2. **Use tokens, not values** — Reference `spacing-md` not `16px`.
3. **Show all states** — Default, hover, active, disabled, loading, error, empty.
4. **Describe the why** — "This collapses on mobile because users primarily use one-handed" helps developers make good judgment calls.

## Output

```markdown
## Handoff Spec: [Feature/Screen Name]

### Overview
[What this screen/feature does, user context]

### Layout
[Grid system, breakpoints, responsive behavior]

### Design Tokens Used
| Token | Value | Usage |
|-------|-------|-------|
| `color-primary` | #[hex] | CTA buttons, links |
| `spacing-md` | [X]px | Between sections |
| `font-heading-lg` | [size/weight/family] | Page title |

### Components
| Component | Variant | Props | Notes |
|-----------|---------|-------|-------|
| [Component] | [Variant] | [Props] | [Special behavior] |

### States and Interactions
| Element | State | Behavior |
|---------|-------|----------|
| [CTA Button] | Hover | [Background darken 10%] |
| [CTA Button] | Loading | [Spinner, disabled] |
| [Form] | Error | [Red border, error message below] |

### Responsive Behavior
| Breakpoint | Changes |
|------------|---------|
| Desktop (>1024px) | [Default layout] |
| Tablet (768-1024px) | [What changes] |
| Mobile (<768px) | [What changes] |

### Edge Cases
- **Empty state**: [What to show when no data]
- **Long text**: [Truncation rules]
- **Loading**: [Skeleton or spinner]
- **Error**: [Error state appearance]

### Animation / Motion
| Element | Trigger | Animation | Duration | Easing |
|---------|---------|-----------|----------|--------|
| [Element] | [Trigger] | [Description] | [ms] | [easing] |

### Accessibility Notes
- [Focus order]
- [ARIA labels needed]
- [Keyboard interactions]
```

## Lark-native delivery & tracking

**Design tool (Figma, external — keep as-is):** pull exact measurements, tokens, and component specs;
export assets. This source stays external.

**Publish the handoff to Wiki (P8 — durable knowledge).** Create the spec page with
`lark_wiki_node_create` (`space_id`, `title`), then fill the rich tables (tokens, components, states,
responsive, edge cases) via the **`lark-doc`** skill (DocxXML — hand-rolling block JSON is brittle).
This is where engineering will look weeks later, not a chat message.

**Upload exported assets to Drive.** Use `lark_drive_upload` for icons/images/redlines and reference
them from the Wiki page. Delegate multipart/folder/permission ops to the **`lark-drive`** skill.

**Create implementation tasks (project tracker = Lark Task).** For each spec section that's real work:
`lark_contact_search` the engineer → `open_id` (P1) → `lark_task_create` with the Wiki link in the
description (`dry_run` first, P2). For subtasks/tasklists, delegate to **`lark-task`**.

**Notify engineering with an interactive card (P4).** `lark_im_card_send` (`print_json` → `dry_run` →
send): `header` "Handoff ready: <screen>", a `div` with the Wiki link, `item` rows for each created
task, and an `actions` footer (*Open spec* / *Acknowledge*). Card grammar → **`lark-im`** skill.

## Tips

1. **Share the Figma link** — I can pull exact measurements, tokens, and component info.
2. **Mention edge cases** — "What happens with 100 items?" helps me spec boundary conditions.
3. **Specify the tech stack** — "We use React + Tailwind" helps me give relevant implementation notes.
