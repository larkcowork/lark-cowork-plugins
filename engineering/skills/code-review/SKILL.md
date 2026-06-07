---
name: code-review
description: Review code changes for security, performance, and correctness. Trigger with a PR URL or diff, "review this before I merge", "is this code safe?", or when checking a change for N+1 queries, injection risks, missing edge cases, or error handling gaps.
argument-hint: "<PR URL, diff, or file path>"
---

# /code-review

> If you see unfamiliar placeholders or need to check which tools are connected, see [CONNECTORS.md](../../CONNECTORS.md).

Review code changes with a structured lens on security, performance, correctness, and maintainability.

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md),
> [LARK-RECIPES](../../../connectors/LARK-RECIPES.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md)).
> Source control stays **GitHub** (`gh`/external MCP) — pull the diff there. The *collaboration*
> layer is Lark: check the change against team standards by searching **Wiki/Docs**
> (`lark_doc_search` → `lark_doc_fetch`, projected with `jq`, P3); post the verdict + critical
> findings as an **interactive card** (`lark_im_card_send`, P4) with a per-finding side button
> (e.g. *File ticket* / *Ack*) rather than a wall of text; turn must-fix findings into tracked work
> via `lark_task_create` (resolve the author via `lark_contact_search` first — P1; `dry_run` — P2),
> or log them to a code-quality **Lark Base** (`lark_base_record_upsert`, P5). For rigorous
> review methodology and the code-reviewer subagent, delegate to the **code-review** skill.

## Usage

```
/code-review <PR URL or file path>
```

Review the provided code changes: @$1

If no specific file or URL is provided, ask what to review.

## How It Works

```
┌─────────────────────────────────────────────────────────────────┐
│                      CODE REVIEW                                   │
├─────────────────────────────────────────────────────────────────┤
│  STANDALONE (always works)                                       │
│  ✓ Paste a diff, PR URL, or point to files                      │
│  ✓ Security audit (OWASP top 10, injection, auth)               │
│  ✓ Performance review (N+1, memory leaks, complexity)           │
│  ✓ Correctness (edge cases, error handling, race conditions)    │
│  ✓ Style (naming, structure, readability)                        │
│  ✓ Actionable suggestions with code examples                    │
├─────────────────────────────────────────────────────────────────┤
│  SUPERCHARGED (when you connect your tools)                      │
│  + Source control: Pull PR diff automatically                    │
│  + Project tracker: Link findings to tickets                     │
│  + Knowledge base: Check against team coding standards           │
└─────────────────────────────────────────────────────────────────┘
```

## Review Dimensions

### Security
- SQL injection, XSS, CSRF
- Authentication and authorization flaws
- Secrets or credentials in code
- Insecure deserialization
- Path traversal
- SSRF

### Performance
- N+1 queries
- Unnecessary memory allocations
- Algorithmic complexity (O(n²) in hot paths)
- Missing database indexes
- Unbounded queries or loops
- Resource leaks

### Correctness
- Edge cases (empty input, null, overflow)
- Race conditions and concurrency issues
- Error handling and propagation
- Off-by-one errors
- Type safety

### Maintainability
- Naming clarity
- Single responsibility
- Duplication
- Test coverage
- Documentation for non-obvious logic

## Output

```markdown
## Code Review: [PR title or file]

### Summary
[1-2 sentence overview of the changes and overall quality]

### Critical Issues
| # | File | Line | Issue | Severity |
|---|------|------|-------|----------|
| 1 | [file] | [line] | [description] | 🔴 Critical |

### Suggestions
| # | File | Line | Suggestion | Category |
|---|------|------|------------|----------|
| 1 | [file] | [line] | [description] | Performance |

### What Looks Good
- [Positive observations]

### Verdict
[Approve / Request Changes / Needs Discussion]
```

## If Connectors Available

Source control = **GitHub** (external, keep as-is):
- Pull the PR diff automatically (`gh pr diff <url>` or the GitHub MCP).
- Check CI status and test results (`gh pr checks`).

Project tracker = **Lark Task + Base** (`lark`):
- Turn must-fix findings into tracked work: resolve author/owner via `lark_contact_search` (P1) →
  `lark_task_create(assignee=open_id, dry_run: true)` (P2) → confirm → commit.
- Or append each finding to a code-quality / review-log **Lark Base** with `lark_base_record_upsert`
  (P5) so trends (recurring N+1s, missing tests) are queryable; scaffold via **base-deploy**.

Knowledge base = **Lark Wiki + Docs** (`lark`):
- Check the change against team coding standards: `lark_doc_search("coding standards"/"style guide")`
  → `lark_doc_fetch` the hit with a `jq` projection (P3), and cite the specific rule in each finding.

Deliver the verdict as an **interactive card** (P4) to the PR channel: `header` colored by verdict
(green Approve / red Request changes), one `item` row per critical issue with a side button, and an
`actions` footer (*File all as tasks* / *Snooze*). Validate the YAML `spec` with `print_json: true`,
then `dry_run: true`, then send (P2). Card grammar lives in the **lark-im** skill.

## Tips

1. **Provide context** — "This is a hot path" or "This handles PII" helps me focus.
2. **Specify concerns** — "Focus on security" narrows the review.
3. **Include tests** — I'll check test coverage and quality too.
