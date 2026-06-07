---
name: pipeline-review
description: Analyze pipeline health — prioritize deals, flag risks, get a weekly action plan. Use when running a weekly pipeline review, deciding which deals to focus on this week, spotting stale or stuck opportunities, auditing for hygiene issues like bad close dates, or identifying single-threaded deals.
argument-hint: "<segment or rep>"
---

# /pipeline-review

> If you see unfamiliar placeholders or need to check which tools are connected, see [CONNECTORS.md](../../CONNECTORS.md).

Analyze your pipeline health, prioritize deals, and get actionable recommendations for where to focus.

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md); tool map in [CONNECTORS.md](../../CONNECTORS.md)).
> There is an **installed `pipeline-review` skill — delegate to it**; it already does the by-stage / stuck-deal / closing-soon / win-rate scan against a Lark Base. The CRM is a **Lark Base** (P5, system-of-record): read deals with `lark_base_search` on the Deals table (requires `search_fields`; pick the review fields with `select_fields` — base_search does NOT support `jq`, P5), and prefer the Base data-query endpoint for stage counts instead of fetching raw records. Write hygiene fixes (updated close date / stage / next-step) back with `lark_base_record_upsert` (`dry_run: true` first, P2). Create follow-up tasks via `lark_contact_search` (P1) → `lark_task_create(dry_run)`. **Surface the priority actions + risk flags as an interactive card** (`lark_im_card_send`, P4) with one `item` row per top deal and an `actions` footer, not a wall of tables. Upcoming meetings per deal come from `lark_calendar_agenda`. For deep Base work delegate to **`lark-base`**.

## Usage

```
/pipeline-review [segment or rep]
```

Review pipeline for: $ARGUMENTS

If a file is referenced: @$1

---

## How It Works

```
┌─────────────────────────────────────────────────────────────────┐
│                     PIPELINE REVIEW                              │
├─────────────────────────────────────────────────────────────────┤
│  STANDALONE (always works)                                       │
│  ✓ Upload CSV export from your CRM                              │
│  ✓ Or paste/describe your deals                                 │
│  ✓ Health check: flag stale, stuck, and at-risk deals          │
│  ✓ Prioritization: rank deals by impact and closability        │
│  ✓ Hygiene audit: missing data, bad close dates, single-thread │
│  ✓ Weekly action plan: what to focus on                        │
├─────────────────────────────────────────────────────────────────┤
│  SUPERCHARGED (when you connect your tools)                      │
│  + CRM: Pull pipeline automatically, update records             │
│  + Activity data for engagement scoring                         │
│  + Historical patterns for risk prediction                      │
│  + Calendar: See upcoming meetings per deal                     │
└─────────────────────────────────────────────────────────────────┘
```

---

## What I Need From You

**Option A: Upload a CSV**
Export your pipeline from your CRM (e.g. Salesforce, HubSpot). Helpful fields:
- Deal/Opportunity name
- Account name
- Amount
- Stage
- Close date
- Created date
- Last activity date
- Owner (if reviewing a team)
- Primary contact

**Option B: Paste your deals**
```
Acme Corp - $50K - Negotiation - closes Jan 31 - last activity Jan 20
TechStart - $25K - Demo scheduled - closes Feb 15 - no activity in 3 weeks
BigCo - $100K - Discovery - closes Mar 30 - created last week
```

**Option C: Describe your pipeline**
"I have 12 deals. Two big ones in negotiation that I'm confident about. Three stuck in discovery for over a month. The rest are mid-stage but I haven't talked to some of them in a while."

---

## Output

```markdown
# Pipeline Review: [Date]

**Data Source:** [CSV upload / Manual input / CRM]
**Deals Analyzed:** [X]
**Total Pipeline Value:** $[X]

---

## Pipeline Health Score: [X/100]

| Dimension | Score | Issue |
|-----------|-------|-------|
| **Stage Progression** | [X]/25 | [X] deals stuck in same stage 30+ days |
| **Activity Recency** | [X]/25 | [X] deals with no activity in 14+ days |
| **Close Date Accuracy** | [X]/25 | [X] deals with close date in past |
| **Contact Coverage** | [X]/25 | [X] deals single-threaded |

---

## Priority Actions This Week

### 1. [Highest Priority Deal]
**Why:** [Reason — large, closing soon, at risk, etc.]
**Action:** [Specific next step]
**Impact:** $[X] if you close it

### 2. [Second Priority]
**Why:** [Reason]
**Action:** [Next step]

### 3. [Third Priority]
**Why:** [Reason]
**Action:** [Next step]

---

## Deal Prioritization Matrix

### Close This Week (Focus Time Here)
| Deal | Amount | Stage | Close Date | Next Action |
|------|--------|-------|------------|-------------|
| [Deal] | $[X] | [Stage] | [Date] | [Action] |

### Close This Month (Keep Warm)
| Deal | Amount | Stage | Close Date | Status |
|------|--------|-------|------------|--------|
| [Deal] | $[X] | [Stage] | [Date] | [Status] |

### Nurture (Check-in Periodically)
| Deal | Amount | Stage | Close Date | Status |
|------|--------|-------|------------|--------|
| [Deal] | $[X] | [Stage] | [Date] | [Status] |

---

## Risk Flags

### Stale Deals (No Activity 14+ Days)
| Deal | Amount | Last Activity | Days Silent | Recommendation |
|------|--------|---------------|-------------|----------------|
| [Deal] | $[X] | [Date] | [X] | [Re-engage / Downgrade / Remove] |

### Stuck Deals (Same Stage 30+ Days)
| Deal | Amount | Stage | Days in Stage | Recommendation |
|------|--------|-------|---------------|----------------|
| [Deal] | $[X] | [Stage] | [X] | [Push / Multi-thread / Qualify out] |

### Past Close Date
| Deal | Amount | Close Date | Days Overdue | Recommendation |
|------|--------|------------|--------------|----------------|
| [Deal] | $[X] | [Date] | [X] | [Update date / Push to next quarter / Close lost] |

### Single-Threaded (Only One Contact)
| Deal | Amount | Contact | Risk | Recommendation |
|------|--------|---------|------|----------------|
| [Deal] | $[X] | [Name] | Champion leaves = deal dies | [Identify additional stakeholders] |

---

## Hygiene Issues

| Issue | Count | Deals | Action |
|-------|-------|-------|--------|
| Missing close date | [X] | [List] | Add realistic close dates |
| Missing amount | [X] | [List] | Estimate or qualify |
| Missing next step | [X] | [List] | Define next action |
| No primary contact | [X] | [List] | Assign contact |

---

## Pipeline Shape

### By Stage
| Stage | # Deals | Value | % of Pipeline |
|-------|---------|-------|---------------|
| [Stage] | [X] | $[X] | [X]% |

### By Close Month
| Month | # Deals | Value |
|-------|---------|-------|
| [Month] | [X] | $[X] |

### By Deal Size
| Size | # Deals | Value |
|------|---------|-------|
| $100K+ | [X] | $[X] |
| $50K-100K | [X] | $[X] |
| $25K-50K | [X] | $[X] |
| <$25K | [X] | $[X] |

---

## Recommendations

### This Week
1. [ ] [Specific action for priority deal 1]
2. [ ] [Action for at-risk deal]
3. [ ] [Hygiene task]

### This Month
1. [ ] [Strategic action]
2. [ ] [Pipeline building if needed]

---

## Deals to Consider Removing

These deals may be dead weight:

| Deal | Amount | Reason | Recommendation |
|------|--------|--------|----------------|
| [Deal] | $[X] | [No activity 60+ days, no response] | Mark closed-lost |
| [Deal] | $[X] | [Pushed 3+ times, no champion] | Qualify out |
```

---

## Prioritization Framework

I'll rank your deals using this framework:

| Factor | Weight | What I Look For |
|--------|--------|-----------------|
| **Close Date** | 30% | Deals closing soonest get priority |
| **Deal Size** | 25% | Bigger deals = more focus |
| **Stage** | 20% | Later stage = more focus |
| **Activity** | 15% | Active deals get prioritized |
| **Risk** | 10% | Lower risk = safer bet |

You can tell me to weight differently: "Focus on big deals over soon deals" or "I need quick wins, prioritize close dates."

---

## If CRM Connected — Lark Base (P5)

- Pull the pipeline: `lark_base_search` on the Deals table
  - `search_fields=["Owner"]` (or the segment field), `query="<segment or rep>"`,
    `select_fields=["Account Name","Amount","Stage","Close Date","Last Activity","Contact"]`, `limit=200`
  - base_search REQUIRES search_fields and does NOT support jq; if field names unknown, discover via
    `lark_api GET /open-apis/bitable/v1/apps/{base}/tables/{table}/fields`
  - stage/size/month shape via the Base data-query endpoint (not client-side counting)
- Write hygiene fixes back with `lark_base_record_upsert(base_token, table_id, fields={...})`
  — `dry_run: true` first, confirm, then commit (P2)
- Create follow-up tasks: `lark_contact_search` the owner (P1) → `lark_task_create(dry_run)`
- Upcoming meetings per deal: `lark_calendar_agenda`
- Track hygiene over time by upserting a "Pipeline Health" snapshot row
- Delegate multi-table joins / dashboards to **`lark-base`**; if the CRM is external, use that MCP server

---

## Tips

1. **Review weekly** — Pipeline health decays fast. Weekly reviews catch issues early.
2. **Kill dead deals** — Stale opportunities inflate your pipeline and distort forecasts. Be ruthless.
3. **Multi-thread everything** — If one person goes dark, you need a backup contact.
4. **Close dates should mean something** — A close date is when you expect signature, not when you hope for one.
