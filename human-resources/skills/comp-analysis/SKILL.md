---
name: comp-analysis
description: Analyze compensation — benchmarking, band placement, and equity modeling. Trigger with "what should we pay a [role]", "is this offer competitive", "model this equity grant", or when uploading comp data to find outliers and retention risks.
argument-hint: "<role, level, or dataset>"
---

# /comp-analysis

> If you see unfamiliar placeholders or need to check which tools are connected, see [CONNECTORS.md](../../CONNECTORS.md).

Analyze compensation data for benchmarking, band placement, and planning. Helps benchmark compensation against market data for hiring, retention, and equity planning.

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md),
> [LARK-RECIPES](../../../connectors/LARK-RECIPES.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md)).
> Verified market benchmarks come from the specialty **`~~compensation data`** external MCP (keep it
> as-is — no Lark equivalent). The **collaboration layer is Lark**: hold comp bands and per-employee
> placement in a **Lark Base** as system-of-record (P5) — read with `lark_base_search`, write with
> `lark_base_record_upsert`; scaffold a new comp-band Base via **`base-deploy`** rather than
> hand-rolling. Pull current-employee comp from `~~HRIS` (or its Base approximation). When the user
> uploads a CSV, land it via `lark_drive_upload` / `lark_sheets_read`. Surface the band-placement
> result and any outliers/retention risks as an **interactive card** (`lark_im_card_send`, P4), not a
> wall of text. Comp is sensitive — keep PII out of broad channels; resolve a single recipient with
> `lark_contact_search` (P1) and DM them.

## Usage

```
/comp-analysis $ARGUMENTS
```

## What I Need From You

**Option A: Single role analysis**
"What should we pay a Senior Software Engineer in SF?"

**Option B: Upload comp data**
Upload a CSV or paste your comp bands. I'll analyze placement, identify outliers, and compare to market.
(CSV → `lark_drive_upload` then `lark_sheets_read` with a `jq` projection, P3; or read an existing
comp Base with `lark_base_search` — note `lark_base_search` does NOT support `jq` and REQUIRES
`search_fields`; narrow with `select_fields`/`limit` and pass the field name(s) to match. If you
don't know the field names, discover them first via `lark_api GET
/open-apis/bitable/v1/apps/{base}/tables/{table}/fields`.)

**Option C: Equity modeling**
"Model a refresh grant of 10K shares over 4 years at a $50 stock price."

## Compensation Framework

### Components of Total Compensation
- **Base salary**: Cash compensation
- **Equity**: RSUs, stock options, or other equity
- **Bonus**: Annual target bonus, signing bonus
- **Benefits**: Health, retirement, perks (harder to quantify)

### Key Variables
- **Role**: Function and specialization
- **Level**: IC levels, management levels
- **Location**: Geographic pay adjustments
- **Company stage**: Startup vs. growth vs. public
- **Industry**: Tech vs. finance vs. healthcare

### Data Sources
- **With ~~compensation data**: Pull verified benchmarks
- **Without**: Use web research, public salary data, and user-provided context
- Always note data freshness and source limitations

## Output

Provide percentile bands (25th, 50th, 75th, 90th) for base, equity, and total comp. Include location adjustments and company-stage context.

```markdown
## Compensation Analysis: [Role/Scope]

### Market Benchmarks
| Percentile | Base | Equity | Total Comp |
|------------|------|--------|------------|
| 25th | $[X] | $[X] | $[X] |
| 50th | $[X] | $[X] | $[X] |
| 75th | $[X] | $[X] | $[X] |
| 90th | $[X] | $[X] | $[X] |

**Sources:** [Web research, compensation data tools, or user-provided data]

### Band Analysis (if data provided)
| Employee | Current Base | Band Min | Band Mid | Band Max | Position |
|----------|-------------|----------|----------|----------|----------|
| [Name] | $[X] | $[X] | $[X] | $[X] | [Below/At/Above] |

### Recommendations
- [Specific compensation recommendations]
- [Equity considerations]
- [Retention risks if applicable]
```

## If Connectors Available

If **~~compensation data** (specialty external MCP) is connected:
- Pull verified market benchmarks by role, level, and location
- Compare your bands against real-time market data

If **~~HRIS** (specialty external MCP, or its Lark Base approximation) is connected:
- Pull current employee comp data for band analysis
- Identify outliers and retention risks automatically

### Lark-native: persist & share the result

- **System-of-record (P5):** keep comp bands + per-employee placement in a **Lark Base**. Read with
  `lark_base_search` — it does NOT support `jq` and REQUIRES `search_fields`; pass the field name(s)
  to match and narrow output with `select_fields`/`limit` (discover field names via `lark_api GET
  /open-apis/bitable/v1/apps/{base}/tables/{table}/fields` if unknown). Write/update each placement
  row with `lark_base_record_upsert` (`dry_run: true` first, P2). Don't have a comp Base yet?
  Scaffold one with **`base-deploy`**.
- **Share securely (P1/P4):** resolve the requesting manager with `lark_contact_search` → DM the
  band-placement summary + outliers/retention flags as an **interactive card** (`lark_im_card_send`,
  `print_json: true` then `dry_run: true`). Never broadcast comp to a group channel.
- For record/field/dashboard depth on the comp Base, delegate to the **`lark-base`** skill.

## Tips

1. **Location matters** — Always specify location for benchmarking. SF vs. Austin vs. London are very different.
2. **Total comp, not just base** — Include equity, bonus, and benefits for a complete picture.
3. **Keep data confidential** — Comp data is sensitive. Results stay in your conversation.
