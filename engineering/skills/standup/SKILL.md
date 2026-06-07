---
name: standup
description: Generate a standup update from recent activity. Use when preparing for daily standup, summarizing yesterday's commits and PRs and ticket moves, formatting work into yesterday/today/blockers, or structuring a few rough notes into a shareable update.
argument-hint: "[yesterday | today | blockers]"
---

# /standup

> If you see unfamiliar placeholders or need to check which tools are connected, see [CONNECTORS.md](../../CONNECTORS.md).

Generate a standup update by pulling together recent activity across your tools.

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md),
> [LARK-RECIPES](../../../connectors/LARK-RECIPES.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md)).
> Source control (GitHub) and CI/CD stay **external** for commits/PRs and build status. The
> collaboration layer is Lark: pull "yesterday/today" work from `lark_task_my` (project with `jq`,
> P3) and any tracker **Base** (`lark_base_search`); scan relevant decisions/threads with
> `lark_im_search`; check what's coming up with `lark_calendar_agenda`. Post the formatted update to
> the standup channel as an **interactive card** (`lark_im_card_send`, P4) — Yesterday / Today /
> Blockers as `div` sections, with a side button to file a blocker as a task; resolve any @-mentioned
> teammate via `lark_contact_search` (P1). Delegate richer flows: a full start-of-day sweep →
> **morning-brief**; the orchestrated agenda+tasks summary → **lark-workflow-standup-report**; an
> end-of-sprint roll-up → **sprint-retro**.

## How It Works

```
┌─────────────────────────────────────────────────────────────────┐
│                        STANDUP                                    │
├─────────────────────────────────────────────────────────────────┤
│  STANDALONE (always works)                                       │
│  ✓ Tell me what you worked on and I'll structure it             │
│  ✓ Format for daily standup (yesterday / today / blockers)      │
│  ✓ Keep it concise and action-oriented                          │
├─────────────────────────────────────────────────────────────────┤
│  SUPERCHARGED (when you connect your tools)                      │
│  + Source control: Recent commits and PRs                        │
│  + Project tracker: Ticket status changes                        │
│  + Chat: Relevant discussions and decisions                      │
│  + CI/CD: Build and deploy status                                │
└─────────────────────────────────────────────────────────────────┘
```

## What I Need From You

**Option A: Let me pull it**
If your tools are connected, just say `/standup` and I'll gather everything automatically.

**Option B: Tell me what you did**
"Worked on the auth migration, reviewed 3 PRs, got blocked on the API rate limiting issue."

## Output

```markdown
## Standup — [Date]

### Yesterday
- [Completed item with ticket reference if available]
- [Completed item]

### Today
- [Planned item with ticket reference]
- [Planned item]

### Blockers
- [Blocker with context and who can help]
```

## If Connectors Available

Source control = **GitHub** (external, keep as-is):
- `gh` / GitHub MCP for recent commits and PRs (opened, reviewed, merged); summarize at a high level.

Project tracker = **Lark Task + Base** (`lark`):
- `lark_task_my` for what you completed/are doing (project with `jq`, P3); `lark_base_search` a sprint
  **Base** for tickets moved to in-progress/done and upcoming items. `lark_base_search` REQUIRES
  `search_fields` (the field name(s) to match) and does NOT support `jq` — discover field names via
  `lark_api GET /open-apis/bitable/v1/apps/{base}/tables/{table}/fields` if unknown, then narrow with
  `select_fields`/`limit`.

Chat = **Lark IM** (`lark`):
- `lark_im_search` for relevant decisions and threads needing your response; flag them in Blockers.
- Post the final update as an **interactive card** (P4) to the standup channel; resolve any
  @-mentioned teammate via `lark_contact_search` first (P1).

CI/CD = **external** (keep as-is): pull build/deploy status from your pipeline MCP if you mention it.

## Tips

1. **Run it every morning** — Build a habit and never scramble for standup notes.
2. **Add context** — After I generate, add any nuance about blockers or priorities.
3. **Share format** — Ask me to format for Lark IM, email, or your team's standup tool.
