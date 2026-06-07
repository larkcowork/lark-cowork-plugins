# Fusion: delegate depth to the installed Lark skills

This machine already has deep, battle-tested Lark skills (`~/.claude/skills/lark-*`,
`base-deploy`, and the business-workflow skills). They are richer than these generic plugins and
own the hard parts (exact APIs, multi-step flows, UI verification). **A soul-swapped plugin skill
should orchestrate and delegate to them, not reimplement them.**

## Domain → owning skill (call instead of hand-rolling)

| Domain / op | Owning skill | Why delegate |
|-------------|--------------|--------------|
| IM send, search, **interactive cards** | `lark-im` | card YAML grammar + spec-reference; card optimization rules |
| Mail compose/send/triage | `lark-mail` | draft/send guards, folders, rules |
| Calendar + room booking + freebusy | `lark-calendar` | `+create`/`+agenda`/`+freebusy` workflows, room recommend |
| Docs read/write (DocxXML) | `lark-doc` | block model is brittle by hand; skill handles XML |
| Wiki spaces/nodes/members | `lark-wiki` | node hierarchy, space membership |
| **Base build (tables/fields/views/roles)** | `base-deploy` (8-phase) + `lark-base` | schema scaffolding with real-UI verification |
| Base record CRUD / formulas / dashboards | `lark-base` | record + field + dashboard ops |
| Tasks + tasklists | `lark-task` | create/list/subtask/assign, agents |
| Minutes AI artifacts (summary/actions/transcript) | `lark-minutes` | pull AI products, not re-summarize |
| Past meetings metadata/participants | `lark-vc` | meeting envelope, attendee snapshot |
| Approval instances/tasks | `lark-approval` | instance create/query, approve/reject |
| OKR cycles/objectives | `lark-okr` | cycle/objective/KR reads |
| Name → open_id | `lark-contact` | directory resolution (P1) |
| Drive upload/organize | `lark-drive` | multipart, folders, permissions |

## Plugin → ready-made workflow skill (often a near-drop-in)

Several plugins duplicate a workflow that already exists as a polished local skill. Prefer wiring
the plugin's command to invoke the local skill:

| Plugin / skill area | Existing local skill to reuse |
|---------------------|-------------------------------|
| productivity · daily brief / start of day | `morning-brief` (fans out mail/IM/approval/tasks → ≤15-line card) |
| productivity · prioritize | `task-prioritizer` (deadline×risk×OKR ranking) |
| productivity · end of day | `daily-digest` |
| sales · pipeline review | `pipeline-review` |
| sales · deal update after a call | `deal-update` (Minutes → Base record → follow-up draft) |
| sales · re-engage dormant accounts | `client-followup` |
| product-management / any · meeting flow | `meeting-prep` (before: context; after: action items → tasks) |
| customer-support / ops · approval queue | `approval-triage`, `approval-flow-sla` |
| mail-heavy roles · inbox | `inbox-zero`, `im-digest` |
| any · 1:1 prep | `one-on-one-prep`; weekly: `weekly-review` |
| HR/ops · standup | `lark-workflow-standup-report` |

## Rule of thumb

If a plugin skill would call more than ~2 `lark_api` recipes, or touches Base schema, or builds a
card with branching logic — **stop and delegate** to the owning skill above. The plugin's job is to
frame the goal in role/company terms (terminology, which Base, which channel) and hand off execution
to the skill that already does it well and verifies its own output.
