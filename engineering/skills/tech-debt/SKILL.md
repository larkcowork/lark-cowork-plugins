---
name: tech-debt
description: Identify, categorize, and prioritize technical debt. Trigger with "tech debt", "technical debt audit", "what should we refactor", "code health", or when the user asks about code quality, refactoring priorities, or maintenance backlog.
---

# Tech Debt Management

Systematically identify, categorize, and prioritize technical debt.

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md),
> [LARK-RECIPES](../../../connectors/LARK-RECIPES.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md)).
> The scoring framework is methodology — keep it. The *backlog* is data: make a tech-debt **Lark Base**
> the system of record (P5) — read existing items with `lark_base_search` (REQUIRES `search_fields`,
> the field name(s) to match; does NOT support `jq` — discover field names via `lark_api GET
> /open-apis/bitable/v1/apps/{base}/tables/{table}/fields` if unknown, then narrow with
> `select_fields`/`limit`),
> write each scored item with `lark_base_record_upsert` (impact/risk/effort/priority as fields), and
> scaffold the register via **base-deploy** if it doesn't exist. Promote the top items into tracked
> remediation work via `lark_task_create` (resolve owner with `lark_contact_search` P1, `dry_run`
> first P2). Surface the prioritized top-N as an **interactive card** (`lark_im_card_send`, P4) with a
> per-item *File as task* side button; land the remediation plan in **Wiki** (`lark_wiki_node_create`,
> P8). Dependency/security debt details still come from your external scanners — only the tracking and
> reporting layer is Lark.

## Categories

| Type | Examples | Risk |
|------|----------|------|
| **Code debt** | Duplicated logic, poor abstractions, magic numbers | Bugs, slow development |
| **Architecture debt** | Monolith that should be split, wrong data store | Scaling limits |
| **Test debt** | Low coverage, flaky tests, missing integration tests | Regressions ship |
| **Dependency debt** | Outdated libraries, unmaintained dependencies | Security vulns |
| **Documentation debt** | Missing runbooks, outdated READMEs, tribal knowledge | Onboarding pain |
| **Infrastructure debt** | Manual deploys, no monitoring, no IaC | Incidents, slow recovery |

## Prioritization Framework

Score each item on:
- **Impact**: How much does it slow the team down? (1-5)
- **Risk**: What happens if we don't fix it? (1-5)
- **Effort**: How hard is the fix? (1-5, inverted — lower effort = higher priority)

Priority = (Impact + Risk) x (6 - Effort)

## Output

Produce a prioritized list with estimated effort, business justification for each item, and a phased remediation plan that can be done alongside feature work.
