---
name: system-design
description: Design systems, services, and architectures. Trigger with "design a system for", "how should we architect", "system design for", "what's the right architecture for", or when the user needs help with API design, data modeling, or service boundaries.
---

# System Design

Help design systems and evaluate architectural decisions.

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md),
> [LARK-FUSION](../../../connectors/LARK-FUSION.md)). The reasoning below is methodology — keep it.
> Where it touches collaboration: ground requirements by searching prior design docs with
> `lark_doc_search` (project with `jq`, P3); land the finished design doc in **Wiki**
> (`lark_wiki_node_create`, P8), filling the body via the **lark-doc** skill; resolve reviewer names
> to `open_id` with `lark_contact_search` (P1) and surface the "design ready for review" decision as
> an **interactive card** (`lark_im_card_send`, P4). To record a specific decision with trade-offs,
> hand off to the **architecture** skill (ADR → Wiki + tasks).

## Framework

### 1. Requirements Gathering
- Functional requirements (what it does)
- Non-functional requirements (scale, latency, availability, cost)
- Constraints (team size, timeline, existing tech stack)

### 2. High-Level Design
- Component diagram
- Data flow
- API contracts
- Storage choices

### 3. Deep Dive
- Data model design
- API endpoint design (REST, GraphQL, gRPC)
- Caching strategy
- Queue/event design
- Error handling and retry logic

### 4. Scale and Reliability
- Load estimation
- Horizontal vs. vertical scaling
- Failover and redundancy
- Monitoring and alerting

### 5. Trade-off Analysis
- Every decision has trade-offs. Make them explicit.
- Consider: complexity, cost, team familiarity, time to market, maintainability

## Output

Produce clear, structured design documents with diagrams (ASCII or described), explicit assumptions, and trade-off analysis. Always identify what you'd revisit as the system grows.
