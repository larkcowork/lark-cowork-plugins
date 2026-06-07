---
name: onboarding
description: Generate an onboarding checklist and first-week plan for a new hire. Use when someone has a start date coming up, building the pre-start task list (accounts, equipment, buddy), scheduling Day 1 and Week 1, or setting 30/60/90-day goals for a new team member.
argument-hint: "<new hire name and role>"
---

# /onboarding

> If you see unfamiliar placeholders or need to check which tools are connected, see [CONNECTORS.md](../../CONNECTORS.md).

Generate a comprehensive onboarding plan for a new team member.

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md),
> [LARK-RECIPES](../../../connectors/LARK-RECIPES.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md)).
> Don't leave the plan as chat text — turn it into real Lark objects: resolve the new hire, manager,
> and buddy with `lark_contact_search` (P1); create each pre-start/Week-1 checklist item as a **Lark
> Task** assigned to the right owner (`lark_task_create`, `dry_run` first, P2; complete with
> `lark_task_complete`); create Day-1 + Week-1 events with `lark_calendar_create` (delegate
> scheduling/room booking to **`lark-calendar`**, check the new hire's free slots with
> `lark_calendar_freebusy`); land the durable onboarding plan + 30/60/90 goals in **Wiki**
> (`lark_wiki_node_create` / `lark_doc_create`, P8; author with **`lark-doc`**); link existing
> onboarding docs/runbooks found via `lark_doc_search`. Track onboarding status per hire in an
> onboarding **Base** (P5, `lark_base_record_upsert`; scaffold via **`base-deploy`**). Send the
> welcome note as a **draft** mail (`lark_mail_draft_create`, P2) and post the Day-1 plan + buddy
> intro as an **interactive card** (`lark_im_card_send`, P4).

## Usage

```
/onboarding $ARGUMENTS
```

## What I Need From You

- **New hire name**: Who's starting?
- **Role**: What position?
- **Team**: Which team are they joining?
- **Start date**: When do they start?
- **Manager**: Who's their manager?

## Output

```markdown
## Onboarding Plan: [Name] — [Role]
**Start Date:** [Date] | **Team:** [Team] | **Manager:** [Manager]

### Pre-Start (Before Day 1)
- [ ] Send welcome email with start date, time, and logistics
- [ ] Set up accounts: email, Lark IM, [tools for role]
- [ ] Order equipment (laptop, monitor, peripherals)
- [ ] Add to team calendar and recurring meetings
- [ ] Assign onboarding buddy: [Suggested person]
- [ ] Prepare desk / remote setup instructions

### Day 1
| Time | Activity | With |
|------|----------|------|
| 9:00 | Welcome and orientation | Manager |
| 10:00 | IT setup and tool walkthrough | IT / Buddy |
| 11:00 | Team introductions | Team |
| 12:00 | Welcome lunch | Manager + Team |
| 1:30 | Company overview and values | Manager |
| 3:00 | Role expectations and 30/60/90 plan | Manager |
| 4:00 | Free time to explore tools and docs | Self |

### Week 1
- [ ] Complete required compliance training
- [ ] Read key documentation: [list for role]
- [ ] 1:1 with each team member
- [ ] Shadow key meetings
- [ ] First small task or project assigned
- [ ] End-of-week check-in with manager

### 30-Day Goals
1. [Goal aligned to role]
2. [Goal aligned to role]
3. [Goal aligned to role]

### 60-Day Goals
1. [Goal]
2. [Goal]

### 90-Day Goals
1. [Goal]
2. [Goal]

### Key Contacts
| Person | Role | For What |
|--------|------|----------|
| [Manager] | Manager | Day-to-day guidance |
| [Buddy] | Onboarding Buddy | Questions, culture, navigation |
| [IT Contact] | IT | Tool access, equipment |
| [HR Contact] | HR | Benefits, policies |

### Tools Access Needed
| Tool | Access Level | Requested |
|------|-------------|-----------|
| [Tool] | [Level] | [ ] |
```

## If Connectors Available

If **~~HRIS** (specialty external MCP, or its Lark Base approximation) is connected:
- Pull new hire details and team org chart
- Auto-populate tools access list based on role

If **~~knowledge base** (Lark Wiki + Docs) is connected:
- Link to relevant onboarding docs, team wikis, and runbooks via `lark_doc_search`
- Pull the team's existing onboarding checklist (`lark_doc_fetch`, `jq` projection P3) to customize

If **~~calendar** (Lark Calendar) is connected:
- Create Day 1 calendar events and Week 1 meeting invites with `lark_calendar_create`
  (`lark_calendar_freebusy` to find slots; delegate to **`lark-calendar`** for room booking)

### Lark-native: make the plan executable

1. **Resolve people (P1):** `lark_contact_search` → `open_id`s for hire, manager, buddy, IT/HR contacts.
2. **Tasks (P2):** each Pre-Start / Week-1 checkbox → `lark_task_create` assigned to its owner
   (accounts/equipment → IT; buddy assignment → manager). `dry_run` first; close with
   `lark_task_complete`. Delegate tasklist depth to **`lark-task`**.
3. **Calendar:** Day-1 schedule + Week-1 1:1s → `lark_calendar_create` (one event per row).
4. **Knowledge (P8):** the full plan + 30/60/90 goals → `lark_wiki_node_create` (or `lark_doc_create`);
   format with **`lark-doc`**. Link role docs found via `lark_doc_search`.
5. **Track (P5):** onboarding status per hire → `lark_base_record_upsert` into an onboarding Base
   (scaffold via **`base-deploy`** if none exists).
6. **Announce (P4):** post Day-1 plan + buddy intro as an `lark_im_card_send` card; welcome email as
   `lark_mail_draft_create` (draft only).

## Tips

1. **Customize for the role** — An engineer's onboarding looks different from a designer's.
2. **Don't overload Day 1** — Focus on setup and relationships. Deep work starts Week 2.
3. **Assign a buddy** — Having a go-to person who isn't their manager makes a huge difference.
