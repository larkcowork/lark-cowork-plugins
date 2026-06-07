---
name: performance-review
description: Structure a performance review with self-assessment, manager template, and calibration prep. Use when review season kicks off and you need a self-assessment template, writing a manager review for a direct report, prepping rating distributions and promotion cases for calibration, or turning vague feedback into specific behavioral examples.
argument-hint: "<employee name or review cycle>"
---

# /performance-review

> If you see unfamiliar placeholders or need to check which tools are connected, see [CONNECTORS.md](../../CONNECTORS.md).

Generate performance review templates and help structure feedback.

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md),
> [LARK-RECIPES](../../../connectors/LARK-RECIPES.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md)).
> Ground reviews in real evidence, not memory: pull the period's completed work from **Lark Task**
> (`lark_task_my` for self-review; the task-list endpoint for a direct report — `lark_task_my` only
> returns the caller, see LARK-RECIPES) and goal progress from **OKR** (`lark_okr_cycle_list`, P6);
> resolve the employee/manager with `lark_contact_search` (P1). Author the self-assessment / manager
> review / calibration doc as a **Wiki/Doc** (`lark_wiki_node_create` / `lark_doc_create`, P8; edit
> via **`lark-doc`**) and share the link, not pasted text. Hold per-employee ratings + comp actions in
> a calibration **Lark Base** (P5, `lark_base_record_upsert`; let the Base aggregation endpoint
> compute the rating distribution; scaffold via **`base-deploy`**). Route any
> promotion/comp-change recommendation through a real Lark **approval instance** (P7, `lark_api` /
> **`lark-approval`**). For comp benchmarking on the compensation recommendation, delegate to
> **`/comp-analysis`**. Present the calibration distribution + flagged cases as an **interactive
> card** (`lark_im_card_send`, P4).

## Usage

```
/performance-review $ARGUMENTS
```

## Modes

```
/performance-review self-assessment       # Generate self-assessment template
/performance-review manager [employee]    # Manager review template for a specific person
/performance-review calibration           # Calibration prep document
```

If no mode is specified, ask what type of review they need.

## Output — Self-Assessment Template

```markdown
## Self-Assessment: [Review Period]

### Key Accomplishments
[List your top 3-5 accomplishments this period. For each, describe the situation, your contribution, and the impact.]

1. **[Accomplishment]**
   - Situation: [Context]
   - Contribution: [What you did]
   - Impact: [Measurable result]

### Goals Review
| Goal | Status | Evidence |
|------|--------|----------|
| [Goal from last period] | Met / Exceeded / Missed | [How you know] |

### Growth Areas
[Where did you grow? New skills, expanded scope, leadership moments.]

### Challenges
[What was hard? What would you do differently?]

### Goals for Next Period
1. [Goal — specific and measurable]
2. [Goal]
3. [Goal]

### Feedback for Manager
[How can your manager better support you?]
```

## Output — Manager Review

```markdown
## Performance Review: [Employee Name]
**Period:** [Date range] | **Manager:** [Your name]

### Overall Rating: [Exceeds / Meets / Below Expectations]

### Performance Summary
[2-3 sentence overall assessment]

### Key Strengths
- [Strength with specific example]
- [Strength with specific example]

### Areas for Development
- [Area with specific, actionable guidance]
- [Area with specific, actionable guidance]

### Goal Achievement
| Goal | Rating | Comments |
|------|--------|----------|
| [Goal] | [Rating] | [Specific observations] |

### Impact and Contributions
[Describe their biggest contributions and impact on the team/org]

### Development Plan
| Skill | Current | Target | Actions |
|-------|---------|--------|---------|
| [Skill] | [Level] | [Level] | [How to get there] |

### Compensation Recommendation
[Promotion / Equity refresh / Adjustment / No change — with justification]
```

## Output — Calibration

```markdown
## Calibration Prep: [Review Cycle]
**Manager:** [Your name] | **Team:** [Team] | **Period:** [Date range]

### Team Overview
| Employee | Role | Level | Tenure | Proposed Rating | Notes |
|----------|------|-------|--------|-----------------|-------|
| [Name] | [Role] | [Level] | [X years] | [Rating] | [Key context] |

### Rating Distribution
| Rating | Count | % of Team | Company Target |
|--------|-------|-----------|----------------|
| Exceeds Expectations | [X] | [X]% | ~15-20% |
| Meets Expectations | [X] | [X]% | ~60-70% |
| Below Expectations | [X] | [X]% | ~10-15% |

### Calibration Discussion Points
1. **[Employee]** — [Why this rating may need discussion, e.g., borderline, first review at level, recent role change]
2. **[Employee]** — [Discussion point]

### Promotion Candidates
| Employee | Current Level | Proposed Level | Justification |
|----------|-------------|----------------|---------------|
| [Name] | [Current] | [Proposed] | [Evidence of next-level performance] |

### Compensation Actions
| Employee | Action | Justification |
|----------|--------|---------------|
| [Name] | [Promotion / Equity refresh / Market adjustment / Retention] | [Why] |

### Manager Notes
[Context the calibration group should know — team changes, org shifts, project impacts]
```

## If Connectors Available

If **~~HRIS** (specialty external MCP, or a calibration **Lark Base**) is connected:
- Pull prior review history and goal tracking data (`lark_base_search` if approximated in a Base)
- Pre-populate employee details and current role information (`lark_contact_search`, P1)

If a project tracker is connected (Lark Task / Lark Base):
- Pull completed work for the review period from `lark_task_my` (caller) or the task-list endpoint
  (direct report); pull goal status from `lark_okr_cycle_list` (P6)
- Reference specific tasks and OKR key-results as concrete evidence

### Lark-native delivery

1. **Gather evidence (P6):** `lark_task_my` / task-list + `lark_okr_cycle_list`, projected with `jq`
   (P3), to turn "great job" into specific, dated accomplishments.
2. **Author the doc (P8):** `lark_wiki_node_create` / `lark_doc_create` for the self-assessment,
   manager review, or calibration doc; edit via **`lark-doc`**.
3. **Calibrate in a Base (P5):** upsert each employee's proposed rating + comp action with
   `lark_base_record_upsert`; compute the rating distribution via the Base aggregation endpoint.
4. **Gate comp/promotion (P7):** create a Lark **approval instance** for any
   promotion/equity-refresh/adjustment (`lark_api` approval recipe / **`lark-approval`**); for the
   comp number itself delegate to **`/comp-analysis`**.
5. **Brief the calibration group (P4):** post the distribution + borderline cases + promotion
   candidates as an `lark_im_card_send` card (DM the calibration owner — keep ratings private).

## Tips

1. **Be specific** — "Great job" isn't feedback. "You reduced deploy time 40% by implementing the new CI pipeline" is.
2. **Balance positive and constructive** — Both are essential. Neither should be a surprise.
3. **Focus on behaviors, not personality** — "Your documentation has been incomplete" vs. "You're careless."
4. **Make development actionable** — "Improve communication" is vague. "Present at the next team all-hands" is actionable.
