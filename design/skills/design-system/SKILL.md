---
name: design-system
description: Audit, document, or extend your design system. Use when checking for naming inconsistencies or hardcoded values across components, writing documentation for a component's variants, states, and accessibility notes, or designing a new pattern that fits the existing system.
argument-hint: "[audit | document | extend] <component or system>"
---

# /design-system

> If you see unfamiliar placeholders or need to check which tools are connected, see [CONNECTORS.md](../../CONNECTORS.md).

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md),
> [LARK-RECIPES](../../../connectors/LARK-RECIPES.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md)).
> The system thinking is unchanged; only the source-of-truth and outputs are Lark. Components/tokens
> live in your design tool (Figma external MCP) — audit naming, variants, token usage there. Treat a
> component-inventory **Lark Base** as the system-of-record for audit findings (P5: read
> `lark_base_search`, write `lark_base_record_upsert`; scaffold via **base-deploy**). Search prior docs
> with `lark_doc_search` and **publish component documentation to Wiki** (`lark_wiki_node_create`, P8)
> via the **`lark-doc`** skill. File remediation as **Lark Tasks** (`lark_task_create`, owner resolved
> with `lark_contact_search`, P1; `dry_run`, P2) and post the audit scorecard as an **interactive
> card** (`lark_im_card_send`, P4).

Manage your design system — audit for consistency, document components, or design new patterns.

## Usage

```
/design-system audit                    # Full system audit
/design-system document [component]     # Document a component
/design-system extend [pattern]         # Design a new component or pattern
```

## Components of a Design System

### Design Tokens
Atomic values that define the visual language:
- Colors (brand, semantic, neutral)
- Typography (scale, weights, line heights)
- Spacing (scale, component padding)
- Borders (radius, width)
- Shadows (elevation levels)
- Motion (durations, easings)

### Components
Reusable UI elements with defined:
- Variants (primary, secondary, ghost)
- States (default, hover, active, disabled, loading, error)
- Sizes (sm, md, lg)
- Behavior (interactions, animations)
- Accessibility (ARIA, keyboard)

### Patterns
Common UI solutions combining components:
- Forms (input groups, validation, submission)
- Navigation (sidebar, tabs, breadcrumbs)
- Data display (tables, cards, lists)
- Feedback (toasts, modals, inline messages)

## Principles

1. **Consistency over creativity** — The system exists so teams don't reinvent the wheel
2. **Flexibility within constraints** — Components should be composable, not rigid
3. **Document everything** — If it's not documented, it doesn't exist
4. **Version and migrate** — Breaking changes need migration paths

## Output — Audit

```markdown
## Design System Audit

### Summary
**Components reviewed:** [X] | **Issues found:** [X] | **Score:** [X/100]

### Naming Consistency
| Issue | Components | Recommendation |
|-------|------------|----------------|
| [Inconsistent naming] | [List] | [Standard to adopt] |

### Token Coverage
| Category | Defined | Hardcoded Values Found |
|----------|---------|----------------------|
| Colors | [X] | [X] instances of hardcoded hex |
| Spacing | [X] | [X] instances of arbitrary values |
| Typography | [X] | [X] instances of custom fonts/sizes |

### Component Completeness
| Component | States | Variants | Docs | Score |
|-----------|--------|----------|------|-------|
| Button | ✅ | ✅ | ⚠️ | 8/10 |
| Input | ✅ | ⚠️ | ❌ | 5/10 |

### Priority Actions
1. [Most impactful improvement]
2. [Second priority]
3. [Third priority]
```

## Output — Document

```markdown
## Component: [Name]

### Description
[What this component is and when to use it]

### Variants
| Variant | Use When |
|---------|----------|
| [Primary] | [Main actions] |
| [Secondary] | [Supporting actions] |

### Props / Properties
| Property | Type | Default | Description |
|----------|------|---------|-------------|
| [prop] | [type] | [default] | [description] |

### States
| State | Visual | Behavior |
|-------|--------|----------|
| Default | [description] | — |
| Hover | [description] | [interaction] |
| Active | [description] | [interaction] |
| Disabled | [description] | Non-interactive |
| Loading | [description] | [animation] |

### Accessibility
- **Role**: [ARIA role]
- **Keyboard**: [Tab, Enter, Escape behavior]
- **Screen reader**: [Announced as...]

### Do's and Don'ts
| ✅ Do | ❌ Don't |
|------|---------|
| [Best practice] | [Anti-pattern] |

### Code Example
[Framework-appropriate code snippet]
```

## Output — Extend

```markdown
## New Component: [Name]

### Problem
[What user need or gap this component addresses]

### Existing Patterns
| Related Component | Similarity | Why It's Not Enough |
|-------------------|-----------|---------------------|
| [Component] | [What's shared] | [What's missing] |

### Proposed Design

#### API / Props
| Property | Type | Default | Description |
|----------|------|---------|-------------|
| [prop] | [type] | [default] | [description] |

#### Variants
| Variant | Use When | Visual |
|---------|----------|--------|
| [Variant] | [Scenario] | [Description] |

#### States
| State | Behavior | Notes |
|-------|----------|-------|
| Default | [Description] | — |
| Hover | [Description] | [Interaction] |
| Disabled | [Description] | Non-interactive |
| Loading | [Description] | [Animation] |

#### Tokens Used
- Colors: [Which tokens]
- Spacing: [Which tokens]
- Typography: [Which tokens]

### Accessibility
- **Role**: [ARIA role]
- **Keyboard**: [Expected interactions]
- **Screen reader**: [Announced as...]

### Open Questions
- [Decision that needs design review]
- [Edge case to resolve]
```

## Lark-native delivery & tracking

**Design tool (Figma, external — keep as-is):** audit components directly in the Figma MCP (naming,
variants, token usage); pull component properties and layer structure for documentation. This source
stays external.

**Track audit findings in a component-inventory Lark Base (P5).** Read current component coverage with
`lark_base_search` (pass `search_fields` for the column(s) to match; `lark_base_search` does NOT
support `jq`, so narrow with `select_fields`/`limit` — discover field names via `lark_api GET
/open-apis/bitable/v1/apps/{base}/tables/{table}/fields` if unknown); write/update one row per
component with
`lark_base_record_upsert` (`fields:{ Component, States, Variants, Docs, Score, "Hardcoded values" }`,
`dry_run` first). No Base yet → **delegate to `base-deploy`** to scaffold the inventory table. The Base
is the source of truth; the audit scorecard and docs are views over it.

**Document components in Wiki (P8 — "if it's not documented, it doesn't exist").** Search existing docs
first with `lark_doc_search`; publish each component page with `lark_wiki_node_create` and fill the
variants/states/a11y/code tables via the **`lark-doc`** skill (DocxXML).

**File remediation as Lark Tasks.** For each Priority Action: `lark_contact_search` the owner →
`open_id` (P1) → `lark_task_create` (`dry_run` first, P2). Delegate richer ops to **`lark-task`**.

**Post the audit scorecard as an interactive card (P4).** `lark_im_card_send` (`print_json` →
`dry_run` → send): `header` with the system score, `item` rows per low-scoring component with a side
*Document* button, `actions` footer (*Open inventory Base* / *Create remediation tasks*). Card grammar
→ **`lark-im`** skill.

## Tips

1. **Start with an audit** — Know where you are before deciding where to go.
2. **Document as you build** — It's easier to document a component while designing it.
3. **Prioritize coverage over perfection** — 80% of components documented beats 100% of 10 components.
