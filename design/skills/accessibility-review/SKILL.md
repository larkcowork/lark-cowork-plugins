---
name: accessibility-review
description: Run a WCAG 2.1 AA accessibility audit on a design or page. Trigger with "audit accessibility", "check a11y", "is this accessible?", or when reviewing a design for color contrast, keyboard navigation, touch target size, or screen reader behavior before handoff.
argument-hint: "<Figma URL, URL, or description>"
---

# /accessibility-review

> If you see unfamiliar placeholders or need to check which tools are connected, see [CONNECTORS.md](../../CONNECTORS.md).

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md),
> [LARK-RECIPES](../../../connectors/LARK-RECIPES.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md)).
> The audit *engine* stays as-is; only delivery/tracking is Lark. The design source stays in your
> design tool (Figma external MCP) — pull color values, font sizes, ARIA, touch targets from there.
> Land the audit where the org will find it (Wiki via `lark_wiki_node_create`, P8), log every finding
> as a row in an a11y-issues **Lark Base** (system-of-record, P5: read `lark_base_search`, write
> `lark_base_record_upsert`; scaffold the table via **base-deploy**), file the fixes as **Lark Tasks**
> (`lark_task_create` after resolving the owner with `lark_contact_search`, P1, `dry_run` first, P2),
> and post the summary as an **interactive card** (`lark_im_card_send`, P4) — not plain chat text.

Audit a design or page for WCAG 2.1 AA accessibility compliance.

## Usage

```
/accessibility-review $ARGUMENTS
```

Audit for accessibility: @$1

## WCAG 2.1 AA Quick Reference

### Perceivable
- **1.1.1** Non-text content has alt text
- **1.3.1** Info and structure conveyed semantically
- **1.4.3** Contrast ratio >= 4.5:1 (normal text), >= 3:1 (large text)
- **1.4.11** Non-text contrast >= 3:1 (UI components, graphics)

### Operable
- **2.1.1** All functionality available via keyboard
- **2.4.3** Logical focus order
- **2.4.7** Visible focus indicator
- **2.5.5** Touch target >= 44x44 CSS pixels

### Understandable
- **3.2.1** Predictable on focus (no unexpected changes)
- **3.3.1** Error identification (describe the error)
- **3.3.2** Labels or instructions for inputs

### Robust
- **4.1.2** Name, role, value for all UI components

## Common Issues

1. Insufficient color contrast
2. Missing form labels
3. No keyboard access to interactive elements
4. Missing alt text on meaningful images
5. Focus traps in modals
6. Missing ARIA landmarks
7. Auto-playing media without controls
8. Time limits without extension options

## Testing Approach

1. Automated scan (catches ~30% of issues)
2. Keyboard-only navigation
3. Screen reader testing (VoiceOver, NVDA)
4. Color contrast verification
5. Zoom to 200% — does layout break?

## Output

```markdown
## Accessibility Audit: [Design/Page Name]
**Standard:** WCAG 2.1 AA | **Date:** [Date]

### Summary
**Issues found:** [X] | **Critical:** [X] | **Major:** [X] | **Minor:** [X]

### Findings

#### Perceivable
| # | Issue | WCAG Criterion | Severity | Recommendation |
|---|-------|---------------|----------|----------------|
| 1 | [Issue] | [1.4.3 Contrast] | 🔴 Critical | [Fix] |

#### Operable
| # | Issue | WCAG Criterion | Severity | Recommendation |
|---|-------|---------------|----------|----------------|
| 1 | [Issue] | [2.1.1 Keyboard] | 🟡 Major | [Fix] |

#### Understandable
| # | Issue | WCAG Criterion | Severity | Recommendation |
|---|-------|---------------|----------|----------------|
| 1 | [Issue] | [3.3.2 Labels] | 🟢 Minor | [Fix] |

#### Robust
| # | Issue | WCAG Criterion | Severity | Recommendation |
|---|-------|---------------|----------|----------------|
| 1 | [Issue] | [4.1.2 Name, Role, Value] | 🟡 Major | [Fix] |

### Color Contrast Check
| Element | Foreground | Background | Ratio | Required | Pass? |
|---------|-----------|------------|-------|----------|-------|
| [Body text] | [color] | [color] | [X]:1 | 4.5:1 | ✅/❌ |

### Keyboard Navigation
| Element | Tab Order | Enter/Space | Escape | Arrow Keys |
|---------|-----------|-------------|--------|------------|
| [Element] | [Order] | [Behavior] | [Behavior] | [Behavior] |

### Screen Reader
| Element | Announced As | Issue |
|---------|-------------|-------|
| [Element] | [What SR says] | [Problem if any] |

### Priority Fixes
1. **[Critical fix]** — Affects [who] and blocks [what]
2. **[Major fix]** — Improves [what] for [who]
3. **[Minor fix]** — Nice to have
```

## Lark-native delivery & tracking

**Design tool (Figma, external — keep as-is):** inspect color values, font sizes, touch targets, and
component ARIA/keyboard behavior directly from the design spec via the Figma MCP. Lark has no
design-canvas equivalent; this layer stays external.

**Log findings to a Lark Base (P5 — system-of-record).** Read the current a11y backlog with
`lark_base_search` (pass `search_fields` for the column(s) to match; `lark_base_search` does NOT
support `jq`, so narrow with `select_fields`/`limit`. If field names are unknown, discover them via
`lark_api GET /open-apis/bitable/v1/apps/{base}/tables/{table}/fields` first); write one row per
finding with `lark_base_record_upsert`
(`base_token`, `table_id`, `fields:{ Issue, "WCAG Criterion", Severity, Recommendation, Status }`,
`dry_run: true` first). If no a11y-issues Base exists yet, **delegate to `base-deploy`** to scaffold
it (don't hand-roll schema).

**File fixes as Lark Tasks (project tracker).** For each Critical/Major finding: resolve the owner
with `lark_contact_search(query="<name>")` → `open_id` (P1), then `lark_task_create` with the WCAG
criterion + severity in the summary (`dry_run: true`, confirm, commit; P2). Delegate richer task ops
(subtasks, tasklists) to the **`lark-task`** skill.

**Land the full audit in Wiki (P8).** Publish the report as a durable page with
`lark_wiki_node_create` (then fill via the **`lark-doc`** skill / DocxXML) so it lives next to the
team's a11y remediation docs — not as a chat message.

**Surface the summary as an interactive card (P4).** Post the issue counts + top priority fixes with
`lark_im_card_send` (`print_json: true` to validate the YAML `spec`, then `dry_run: true`, then send):
a colored `header` ("X issues, Y critical"), one `item` row per Critical fix with a side *Create task*
button, and an `actions` footer. Card grammar lives in the **`lark-im`** skill.

## Tips

1. **Start with contrast and keyboard** — These catch the most common and impactful issues.
2. **Test with real assistive technology** — My audit is a great start, but manual testing with VoiceOver/NVDA catches things I can't.
3. **Prioritize by impact** — Fix issues that block users first, polish later.
