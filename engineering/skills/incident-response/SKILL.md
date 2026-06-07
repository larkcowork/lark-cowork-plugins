---
name: incident-response
description: Run an incident response workflow — triage, communicate, and write postmortem. Trigger with "we have an incident", "production is down", an alert that needs severity assessment, a status update mid-incident, or when writing a blameless postmortem after resolution.
argument-hint: "<incident description or alert>"
---

# /incident-response

> If you see unfamiliar placeholders or need to check which tools are connected, see [CONNECTORS.md](../../CONNECTORS.md).

Manage an incident from detection through postmortem.

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md),
> [LARK-RECIPES](../../../connectors/LARK-RECIPES.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md)).
> Monitoring (Datadog) and incident management (PagerDuty) stay **external** — pull alerts/metrics and
> page on-call there. The collaboration layer is fully Lark:
> - **Triage/comms** — post each status update as an **interactive card** (`lark_im_card_send`, P4):
>   `header` colored by severity (red SEV1), `div` rows for impact/status, `note` for "next update at
>   …". Resolve IC/comms/responder names to `open_id` with `lark_contact_search` (P1) before
>   @-mentioning or paging. Open the war-room channel and broadcast via `lark_im_send`.
> - **Timeline** — reconstruct from `lark_im_search` over the incident channel (project with `jq`,
>   P3); pull AI artifacts from `lark_minutes_search` if a bridge call was recorded (P6).
> - **Postmortem** — it is durable knowledge: publish to **Wiki** (`lark_wiki_node_create`, P8) and
>   fill via the **lark-doc** skill. Log the incident to an incident **Base**
>   (`lark_base_record_upsert`, P5) for MTTR/trend reporting.
> - **Action items** — each becomes a Lark Task: resolve owner (P1) → `lark_task_create(dry_run)` (P2)
>   → confirm → commit.
> Always `dry_run` mutating calls first (P2). For a full blameless postmortem from the chat timeline,
> **delegate to the incident-retro skill** rather than hand-rolling it here.

## Usage

```
/incident-response $ARGUMENTS
```

## Modes

```
/incident-response new [description]     # Start a new incident
/incident-response update [status]       # Post a status update
/incident-response postmortem            # Generate postmortem from incident data
```

If no mode is specified, ask what phase the incident is in.

## How It Works

```
┌─────────────────────────────────────────────────────────────────┐
│                    INCIDENT RESPONSE                               │
├─────────────────────────────────────────────────────────────────┤
│  Phase 1: TRIAGE                                                  │
│  ✓ Assess severity (SEV1-4)                                     │
│  ✓ Identify affected systems and users                          │
│  ✓ Assign roles (IC, comms, responders)                         │
│                                                                    │
│  Phase 2: COMMUNICATE                                              │
│  ✓ Draft internal status update                                  │
│  ✓ Draft customer communication (if needed)                     │
│  ✓ Set up war room and cadence                                   │
│                                                                    │
│  Phase 3: MITIGATE                                                 │
│  ✓ Document mitigation steps taken                               │
│  ✓ Track timeline of events                                      │
│  ✓ Confirm resolution                                            │
│                                                                    │
│  Phase 4: POSTMORTEM                                               │
│  ✓ Blameless postmortem document                                 │
│  ✓ Timeline reconstruction                                       │
│  ✓ Root cause analysis (5 whys)                                  │
│  ✓ Action items with owners                                      │
└─────────────────────────────────────────────────────────────────┘
```

## Severity Classification

| Level | Criteria | Response Time |
|-------|----------|---------------|
| SEV1 | Service down, all users affected | Immediate, all-hands |
| SEV2 | Major feature degraded, many users affected | Within 15 min |
| SEV3 | Minor feature issue, some users affected | Within 1 hour |
| SEV4 | Cosmetic or low-impact issue | Next business day |

## Communication Guidance

Provide clear, factual updates at regular cadence. Include: what's happening, who's affected, what we're doing, when the next update is.

## Output — Status Update

```markdown
## Incident Update: [Title]
**Severity:** SEV[1-4] | **Status:** Investigating | Identified | Monitoring | Resolved
**Impact:** [Who/what is affected]
**Last Updated:** [Timestamp]

### Current Status
[What we know now]

### Actions Taken
- [Action 1]
- [Action 2]

### Next Steps
- [What's happening next and ETA]

### Timeline
| Time | Event |
|------|-------|
| [HH:MM] | [Event] |
```

## Output — Postmortem

```markdown
## Postmortem: [Incident Title]
**Date:** [Date] | **Duration:** [X hours] | **Severity:** SEV[X]
**Authors:** [Names] | **Status:** Draft

### Summary
[2-3 sentence plain-language summary]

### Impact
- [Users affected]
- [Duration of impact]
- [Business impact if quantifiable]

### Timeline
| Time (UTC) | Event |
|------------|-------|
| [HH:MM] | [Event] |

### Root Cause
[Detailed explanation of what caused the incident]

### 5 Whys
1. Why did [symptom]? → [Because...]
2. Why did [cause 1]? → [Because...]
3. Why did [cause 2]? → [Because...]
4. Why did [cause 3]? → [Because...]
5. Why did [cause 4]? → [Root cause]

### What Went Well
- [Things that worked]

### What Went Poorly
- [Things that didn't work]

### Action Items
| Action | Owner | Priority | Due Date |
|--------|-------|----------|----------|
| [Action] | [Person] | P0/P1/P2 | [Date] |

### Lessons Learned
[Key takeaways for the team]
```

## If Connectors Available

Monitoring = **Datadog** (external, keep as-is):
- Pull alert details and metrics; show graphs of affected metrics.

Incident management = **PagerDuty / Opsgenie** (external, keep as-is):
- Create or update the incident and page on-call responders there.

Chat = **Lark IM** (`lark`):
- Post status updates to the incident channel as **interactive cards** (P4), not plain text — so the
  team can ack / take roles inline. Open/announce the war-room channel with `lark_im_send`. Resolve
  every responder to `open_id` first (`lark_contact_search`, P1).

Knowledge base + tracker = **Lark Wiki + Task + Base** (`lark`):
- Publish the postmortem to **Wiki** (`lark_wiki_node_create`, P8); fill via **lark-doc**.
- Convert each postmortem action item to a tracked task (`lark_task_create`, owner resolved P1,
  `dry_run` P2) and upsert the incident record into an incident **Base** (`lark_base_record_upsert`,
  P5) for MTTR/severity trend reporting.

## Tips

1. **Start writing immediately** — Don't wait for complete information. Update as you learn more.
2. **Keep updates factual** — What we know, what we've done, what's next. No speculation.
3. **Postmortems are blameless** — Focus on systems and processes, not individuals.
