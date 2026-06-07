---
name: deploy-checklist
description: Pre-deployment verification checklist. Use when about to ship a release, deploying a change with database migrations or feature flags, verifying CI status and approvals before going to production, or documenting rollback triggers ahead of time.
argument-hint: "[service or release name]"
---

# /deploy-checklist

> If you see unfamiliar placeholders or need to check which tools are connected, see [CONNECTORS.md](../../CONNECTORS.md).

Generate a pre-deployment checklist to verify readiness before shipping.

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md),
> [LARK-RECIPES](../../../connectors/LARK-RECIPES.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md)).
> CI/CD, source control (GitHub) and monitoring (Datadog) stay **external** — read build/test status,
> the release diff, and baseline thresholds from those. The collaboration layer is Lark: render the
> checklist as an **interactive card** with one checkbox-style `item` per gate
> (`lark_im_card_send`, P4); notify the on-call/release channel via `lark_im_send` after resolving
> recipients with `lark_contact_search` (P1); gate the production go/no-go with a real **approval
> instance** (P7, `lark-approval`) instead of an ad-hoc "ship it" reply; close related tickets with
> `lark_task_complete`; land the release notes / runbook in **Wiki** (`lark_wiki_node_create`, P8)
> and log the deploy to a release **Base** (`lark_base_record_upsert`, P5). Always `dry_run` mutating
> calls first (P2).

## Usage

```
/deploy-checklist $ARGUMENTS
```

## Output

```markdown
## Deploy Checklist: [Service/Release]
**Date:** [Date] | **Deployer:** [Name]

### Pre-Deploy
- [ ] All tests passing in CI
- [ ] Code reviewed and approved
- [ ] No known critical bugs in release
- [ ] Database migrations tested (if applicable)
- [ ] Feature flags configured (if applicable)
- [ ] Rollback plan documented
- [ ] On-call team notified

### Deploy
- [ ] Deploy to staging and verify
- [ ] Run smoke tests
- [ ] Deploy to production (canary if available)
- [ ] Monitor error rates and latency for 15 min
- [ ] Verify key user flows

### Post-Deploy
- [ ] Confirm metrics are nominal
- [ ] Update release notes / changelog
- [ ] Notify stakeholders
- [ ] Close related tickets

### Rollback Triggers
- Error rate exceeds [X]%
- P50 latency exceeds [X]ms
- [Critical user flow] fails
```

## Customization

Tell me about your deploy and I'll customize the checklist:
- "We use feature flags" → adds flag verification steps
- "This includes a database migration" → adds migration-specific checks
- "This is a breaking API change" → adds consumer notification steps

## If Connectors Available

Source control = **GitHub** (external, keep as-is):
- `gh` / GitHub MCP to pull the release diff and verify all PRs are approved and merged.

CI/CD = **external** (keep as-is):
- Check build/test status from your pipeline MCP; verify green before deploy. (No Lark CI product —
  approximate with a deploy-log **Lark Base** if you have no pipeline integration.)

Monitoring = **Datadog** (external, keep as-is):
- Pre-fill rollback trigger thresholds from current baselines; set up the post-deploy metric watch.

Chat + tracker + approval = **Lark**:
- Post the checklist as an **interactive card** (P4) to the release channel; resolve on-call/
  stakeholder `open_id`s via `lark_contact_search` (P1) and notify with `lark_im_send`.
- **Gate the production go/no-go** with a real approval instance (P7) — create it via `lark_api`
  (LARK-RECIPES → Approval) or delegate to **lark-approval**; the card kicks it off, the approval
  engine tracks state and notifies approvers. Don't rely on a freeform "yes".
- On success: `lark_task_complete` the related tickets, publish release notes to **Wiki**
  (`lark_wiki_node_create`, P8), and upsert the deploy record (version, deployer, result) into a
  release **Base** with `lark_base_record_upsert` (P5).

## Tips

1. **Run before every deploy** — Even routine ones. Checklists prevent "I forgot to..."
2. **Customize once, reuse** — Tell me your stack and I'll remember your deploy process.
3. **Include rollback criteria** — Decide when to roll back before you deploy, not during.
