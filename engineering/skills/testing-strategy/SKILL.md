---
name: testing-strategy
description: Design test strategies and test plans. Trigger with "how should we test", "test strategy for", "write tests for", "test plan", "what tests do we need", or when the user needs help with testing approaches, coverage, or test architecture.
---

# Testing Strategy

Design effective testing strategies balancing coverage, speed, and maintenance.

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md),
> [LARK-FUSION](../../../connectors/LARK-FUSION.md)). The strategy itself is methodology — keep it.
> Where it produces deliverables: land the test plan in **Wiki** (`lark_wiki_node_create`, P8, body
> via the **lark-doc** skill) so it's discoverable; turn each identified coverage gap into a tracked
> task (`lark_task_create`, resolve owner via `lark_contact_search` P1, `dry_run` first P2); and post
> a short coverage-gap summary as an **interactive card** (`lark_im_card_send`, P4) if the team needs
> to decide what to prioritize. Coverage numbers come from your external CI/coverage tooling — only
> the tracking and publishing layer is Lark.

## Testing Pyramid

```
        /  E2E  \         Few, slow, high confidence
       / Integration \     Some, medium speed
      /    Unit Tests  \   Many, fast, focused
```

## Strategy by Component Type

- **API endpoints**: Unit tests for business logic, integration tests for HTTP layer, contract tests for consumers
- **Data pipelines**: Input validation, transformation correctness, idempotency tests
- **Frontend**: Component tests, interaction tests, visual regression, accessibility
- **Infrastructure**: Smoke tests, chaos engineering, load tests

## What to Cover

Focus on: business-critical paths, error handling, edge cases, security boundaries, data integrity.

Skip: trivial getters/setters, framework code, one-off scripts.

## Output

Produce a test plan with: what to test, test type for each area, coverage targets, and example test cases. Identify gaps in existing coverage.
