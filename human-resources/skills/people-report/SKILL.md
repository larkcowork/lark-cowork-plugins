---
name: people-report
description: Generate headcount, attrition, diversity, or org health reports. Use when pulling a headcount snapshot for leadership, analyzing turnover trends by team, preparing diversity representation metrics, or assessing span of control and flight risk across the org.
argument-hint: "<report type — headcount, attrition, diversity, org health>"
---

# /people-report

> If you see unfamiliar placeholders or need to check which tools are connected, see [CONNECTORS.md](../../CONNECTORS.md).

Generate people analytics reports from your HR data. Analyze workforce data to surface trends, risks, and opportunities.

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md),
> [LARK-RECIPES](../../../connectors/LARK-RECIPES.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md)).
> Workforce data lives in the `~~HRIS` specialty MCP (keep as-is) **or** a **Lark Base** approximation
> — read with `lark_base_search` and let the Base **data-query/aggregation endpoint** compute counts,
> attrition rates, and distributions (don't fetch raw records to count client-side). CSV uploads land
> via `lark_drive_upload` → `lark_sheets_read` (project with `jq`, P3). Deliver the report two ways:
> land the full narrative + methodology as a durable **Wiki doc** (`lark_wiki_node_create` /
> `lark_doc_create`, P8; author via **`lark-doc`**), and post the executive summary + key-metric
> trends as an **interactive card** (`lark_im_card_send`, P4). Diversity/flight-risk data is
> sensitive — resolve a specific leadership recipient with `lark_contact_search` (P1) and DM the
> card; never broadcast PII to a wide group.

## Usage

```
/people-report $ARGUMENTS
```

## Report Types

**Headcount**: Current org snapshot — by team, location, level, tenure
**Attrition**: Turnover analysis — voluntary/involuntary, by team, trends
**Diversity**: Representation metrics — by level, team, pipeline
**Org Health**: Span of control, management layers, team sizes, flight risk

## Key Metrics

### Retention
- Overall attrition rate (voluntary + involuntary)
- Regrettable attrition rate
- Average tenure
- Flight risk indicators

### Diversity
- Representation by level, team, and function
- Pipeline diversity (hiring funnel by demographic)
- Promotion rates by group
- Pay equity analysis

### Engagement
- Survey scores and trends
- eNPS (Employee Net Promoter Score)
- Participation rates
- Open-ended feedback themes

### Productivity
- Revenue per employee
- Span of control efficiency
- Time to productivity for new hires

## Approach

1. Understand what question they're trying to answer
2. Identify the right data: upload (`lark_drive_upload` → `lark_sheets_read`), paste, or pull from
   `~~HRIS` (specialty MCP) / the org **Lark Base** (`lark_base_search` — no `jq`; REQUIRES
   `search_fields`, narrow with `select_fields`/`limit`; discover field names via `lark_api GET
   /open-apis/bitable/v1/apps/{base}/tables/{table}/fields` if unknown)
3. Analyze — push counts/rates/distributions to the Base aggregation endpoint, not client-side
4. Present findings with context and caveats
5. Recommend specific actions based on data

## What I Need From You

Upload a CSV or describe your data. Helpful fields:
- Employee name/ID, department, team
- Title, level, location
- Start date, end date (if applicable)
- Manager, compensation (if relevant)
- Demographics (for diversity reports, if available)

## Output

```markdown
## People Report: [Type] — [Date]

### Executive Summary
[2-3 key takeaways]

### Key Metrics
| Metric | Value | Trend |
|--------|-------|-------|
| [Metric] | [Value] | [up/down/flat] |

### Detailed Analysis
[Charts, tables, and narrative for the specific report type]

### Recommendations
- [Data-driven recommendation]
- [Action item]

### Methodology
[How the numbers were calculated, any caveats]
```

## If Connectors Available

If **~~HRIS** (specialty external MCP, or a **Lark Base** approximation) is connected:
- Pull live employee data — headcount, tenure, department, level (`lark_base_search` if in a Base)
- Generate reports without needing a CSV upload

If **~~chat** (Lark IM) is connected:
- Resolve the recipient with `lark_contact_search` (P1), then post the executive summary as an
  **interactive card** via `lark_im_card_send` (P4) — header + key-metric rows with trend pills.
  DM sensitive (diversity/flight-risk) reports; don't broadcast PII.

### Lark-native delivery

- **Land the full report (P8):** `lark_wiki_node_create` (or `lark_doc_create`) for the narrative +
  methodology + charts; author with **`lark-doc`**.
- **Snapshot to a Base (P5):** optionally persist each run's headline metrics to a people-analytics
  Base via `lark_base_record_upsert` to trend over time (scaffold via **`base-deploy`**).
