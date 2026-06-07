---
name: debug
description: Structured debugging session — reproduce, isolate, diagnose, and fix. Trigger with an error message or stack trace, "this works in staging but not prod", "something broke after the deploy", or when behavior diverges from expected and the cause isn't obvious.
argument-hint: "<error message or problem description>"
---

# /debug

> If you see unfamiliar placeholders or need to check which tools are connected, see [CONNECTORS.md](../../CONNECTORS.md).

Run a structured debugging session to find and fix issues systematically.

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md),
> [LARK-RECIPES](../../../connectors/LARK-RECIPES.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md)).
> Monitoring (Datadog) and source control (GitHub) stay **external** — pull logs/metrics and recent
> commits from those. The collaboration layer is Lark: scan the incident/eng channel for the first
> report and prior context with `lark_im_search` (project with `jq`, P3); search **Wiki/Docs** for a
> known-issue runbook (`lark_doc_search`); file the fix as a tracked task (`lark_task_create`, resolve
> owner via `lark_contact_search` P1, `dry_run` P2) or log the bug to a defect **Base**
> (`lark_base_record_upsert`, P5); post the root-cause + fix summary as an **interactive card**
> (`lark_im_card_send`, P4). For deep root-cause methodology delegate to the **debugging** skill; if
> this turns into a real outage, hand off to **incident-response** / **incident-retro**.

## Usage

```
/debug $ARGUMENTS
```

## How It Works

```
┌─────────────────────────────────────────────────────────────────┐
│                       DEBUG                                        │
├─────────────────────────────────────────────────────────────────┤
│  Step 1: REPRODUCE                                                │
│  ✓ Understand the expected vs. actual behavior                   │
│  ✓ Identify exact reproduction steps                             │
│  ✓ Determine scope (when did it start? who is affected?)        │
│                                                                    │
│  Step 2: ISOLATE                                                   │
│  ✓ Narrow down the component, service, or code path             │
│  ✓ Check recent changes (deploys, config changes, dependencies) │
│  ✓ Review logs and error messages                                │
│                                                                    │
│  Step 3: DIAGNOSE                                                  │
│  ✓ Form hypotheses and test them                                 │
│  ✓ Trace the code path                                           │
│  ✓ Identify root cause (not just symptoms)                      │
│                                                                    │
│  Step 4: FIX                                                       │
│  ✓ Propose a fix with explanation                                │
│  ✓ Consider side effects and edge cases                          │
│  ✓ Suggest tests to prevent regression                           │
└─────────────────────────────────────────────────────────────────┘
```

## What I Need From You

Tell me about the problem. Any of these help:
- Error message or stack trace
- Steps to reproduce
- What changed recently
- Logs or screenshots
- Expected vs. actual behavior

## Output

```markdown
## Debug Report: [Issue Summary]

### Reproduction
- **Expected**: [What should happen]
- **Actual**: [What happens instead]
- **Steps**: [How to reproduce]

### Root Cause
[Explanation of why the bug occurs]

### Fix
[Code changes or configuration fixes needed]

### Prevention
- [Test to add]
- [Guard to put in place]
```

## If Connectors Available

Monitoring = **Datadog** (external, keep as-is):
- Pull logs, error rates, and metrics around the time of the issue; show recent deploys/config
  changes that correlate. (No Lark equivalent — use the Datadog MCP, or approximate with a metrics
  **Lark Base** + Sheets if there's no APM.)

Source control = **GitHub** (external, keep as-is):
- `gh` / GitHub MCP to find recent commits and PRs touching the affected code path and check whether
  the issue correlates with a specific change.

Project tracker = **Lark Task + Base** + **chat** (`lark`):
- Search for prior reports of the same bug: `lark_im_search` the eng channel (project with `jq`,
  P3) and `lark_base_search` your defect Base. Note: `lark_base_search` REQUIRES `search_fields`
  (the field name(s) to match on) and does NOT support `jq` — if you don't know the field names,
  discover them first via `lark_api GET /open-apis/bitable/v1/apps/{base}/tables/{table}/fields`,
  then narrow with `select_fields`/`limit` instead of a jq projection.
- Once the fix is identified, create a tracked task: `lark_contact_search` the owner (P1) →
  `lark_task_create(assignee=open_id, dry_run: true)` (P2) → confirm → commit. Optionally upsert the
  defect record (cause, fix, prevention) into the Base with `lark_base_record_upsert` (P5).
- Post the root-cause + fix as an **interactive card** (P4) to the channel so responders can ack /
  take the follow-up inline.

## Tips

1. **Share error messages exactly** — Don't paraphrase. The exact text matters.
2. **Mention what changed** — Recent deploys, dependency updates, and config changes are top suspects.
3. **Include context** — "This works in staging but not prod" or "Only affects large payloads" narrows things fast.
