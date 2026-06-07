---
name: analyze
description: Answer data questions -- from quick lookups to full analyses. Use when looking up a single metric, investigating what's driving a trend or drop, comparing segments over time, or preparing a formal data report for stakeholders.
argument-hint: "<question>"
---

# /analyze - Answer Data Questions

> If you see unfamiliar placeholders or need to check which tools are connected, see [CONNECTORS.md](../../CONNECTORS.md).

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md),
> [LARK-RECIPES](../../../connectors/LARK-RECIPES.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md)).
> The SQL/analysis engine stays specialty-external (warehouse MCP). The **collaboration layer is Lark**:
> when no warehouse is connected, a **Lark Base or Sheet is a valid data source** (read a Base with
> `lark_base_search` — it REQUIRES `search_fields` and does NOT support `jq`, so narrow with
> `select_fields`/`limit`; read a Sheet with `lark_sheets_read`, P3 / P5). **Land the deliverable in
> Lark** — a formal report goes to a Wiki node (`lark_wiki_node_create` → `lark_doc_create`, P8); an
> attachment (CSV/notebook) goes to `lark_drive_upload`. **Surface the headline as an interactive card**
> (`lark_im_card_send`, P4) — KPI + trend pill + a button to open the full report — not a wall of text.
> Resolve any recipient via `lark_contact_search` first (P1); preview every share with `dry_run` (P2).

Answer a data question, from a quick lookup to a full analysis to a formal report.

## Usage

```
/analyze <natural language question>
```

## Workflow

### 1. Understand the Question

Parse the user's question and determine:

- **Complexity level**:
  - **Quick answer**: Single metric, simple filter, factual lookup (e.g., "How many users signed up last week?")
  - **Full analysis**: Multi-dimensional exploration, trend analysis, comparison (e.g., "What's driving the drop in conversion rate?")
  - **Formal report**: Comprehensive investigation with methodology, caveats, and recommendations (e.g., "Prepare a quarterly business review of our subscription metrics")
- **Data requirements**: Which tables, metrics, dimensions, and time ranges are needed
- **Output format**: Number, table, chart, narrative, or combination

### 2. Gather Data

**If a data warehouse MCP server is connected:**

1. Explore the schema to find relevant tables and columns
2. Write SQL query(ies) to extract the needed data
3. Execute the query and retrieve results
4. If the query fails, debug and retry (check column names, table references, syntax for the specific dialect)
5. If results look unexpected, run sanity checks before proceeding

**If no data warehouse is connected:**

1. Check for a **Lark-native source first** (CONNECTORS.md says approximate with Lark Base/Sheets):
   - A tracker/CRM/log **Lark Base** — `lark_base_search` **requires `search_fields`** (the field(s) to
     match; if you don't know the field names, discover them via `lark_api` GET
     `/open-apis/bitable/v1/apps/{base}/tables/{table}/fields`). It does NOT support `jq` — narrow the
     payload with `select_fields`/`limit`. Aggregate via the Base data-query endpoint per LARK-RECIPES, P5.
   - A **Lark Sheet** — `lark_sheets_read` a named range, then load into a DataFrame.
   - For finding the Base/Sheet by name, delegate to the `lark-base` / `lark-sheets` skills.
2. Otherwise ask the user to paste results, upload a CSV/Excel, or describe the schema.
3. If writing queries for manual execution, use the `sql-queries` skill for dialect-specific best practices.
4. Once data is provided, proceed with analysis.

### 3. Analyze

- Calculate relevant metrics, aggregations, and comparisons
- Identify patterns, trends, outliers, and anomalies
- Compare across dimensions (time periods, segments, categories)
- For complex analyses, break the problem into sub-questions and address each

### 4. Validate Before Presenting

Before sharing results, run through validation checks:

- **Row count sanity**: Does the number of records make sense?
- **Null check**: Are there unexpected nulls that could skew results?
- **Magnitude check**: Are the numbers in a reasonable range?
- **Trend continuity**: Do time series have unexpected gaps?
- **Aggregation logic**: Do subtotals sum to totals correctly?

If any check raises concerns, investigate and note caveats.

### 5. Present Findings

**For quick answers:**
- State the answer directly with relevant context
- Include the query used (collapsed or in a code block) for reproducibility

**For full analyses:**
- Lead with the key finding or insight
- Support with data tables and/or visualizations
- Note methodology and any caveats
- Suggest follow-up questions

**For formal reports:**
- Executive summary with key takeaways
- Methodology section explaining approach and data sources
- Detailed findings with supporting evidence
- Caveats, limitations, and data quality notes
- Recommendations and suggested next steps

### 6. Visualize Where Helpful

When a chart would communicate results more effectively than a table:

- Use the `data-visualization` skill to select the right chart type
- Generate a Python visualization or build it into an HTML dashboard
- Follow visualization best practices for clarity and accuracy

### 7. Land & Share the Result (Lark-native)

Don't leave the deliverable in the chat scrollback:

- **Formal report / full analysis** → create a durable doc in **Wiki** (`lark_wiki_node_create` then
  fill via `lark_doc_create` with `wiki_node`+`wiki_space`, P8). For non-trivial doc body, delegate to
  the `lark-doc` skill.
- **Chart PNGs / notebooks / CSV exports** → `lark_drive_upload` (P8) and link them from the doc.
- **Recurring metric you want tracked over time** → upsert into a Lark Base via
  `lark_base_record_upsert` (the Base is the system-of-record, P5).
- **Headline for stakeholders** → post an **interactive card** (`lark_im_card_send`, P4): colored
  `header`, the key number with a status pill, a `note` with the "as of" date and caveats, and a
  button linking to the full report. `print_json: true` → `dry_run: true` → send. Resolve recipients
  via `lark_contact_search` (P1).

## Examples

**Quick answer:**
```
/analyze How many new users signed up in December?
```

**Full analysis:**
```
/analyze What's causing the increase in support ticket volume over the past 3 months? Break down by category and priority.
```

**Formal report:**
```
/analyze Prepare a data quality assessment of our customer table -- completeness, consistency, and any issues we should address.
```

## Tips

- Be specific about time ranges, segments, or metrics when possible
- If you know the table names, mention them to speed up the process
- For complex questions, Claude may break them into multiple queries
- Results are always validated before presentation -- if something looks off, Claude will flag it
