---
name: customer-research
description: Multi-source research on a customer question or topic with source attribution. Use when a customer asks something you need to look up, investigating whether a bug has been reported before, checking what was previously told to a specific account, or gathering background before drafting a response.
argument-hint: "<question or topic>"
---

# /customer-research

> If you see unfamiliar placeholders or need to check which tools are connected, see [CONNECTORS.md](../../CONNECTORS.md).

Multi-source research on a customer question, product topic, or account-related inquiry. Synthesizes findings from all available sources with clear attribution and confidence scoring.

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md), [LARK-RECIPES](../../../connectors/LARK-RECIPES.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md)).
> Search across Lark surfaces and **project reads with `jq`** wherever the tool supports it (P3) — a full doc fetch is ~25k tokens, `jq: ".data.content[:4000]"` is ~500. The internal sources map to: Wiki/Docs (`lark_doc_search` → `lark_doc_fetch`), account/ticket **Lark Base** (`lark_base_search`, P5 — note: base_search does NOT support `jq` and REQUIRES `search_fields`; narrow with `select_fields`/`limit`), chat (`lark_im_search`), mail (`lark_api` mail search), meetings (`lark_minutes_search`, P6). Knowledge capture lands in **Wiki** via `lark_wiki_node_create` (P8) — delegate doc editing to **`lark-doc`**. For an account-context deep dive on one person, delegate to **`contact-360`**.

## Usage

```
/customer-research <question or topic>
```

## Workflow

### 1. Parse the Research Request

Identify what type of research is needed:
- **Customer question**: Something a customer has asked that needs an answer (e.g., "Does our product support SSO with Okta?")
- **Issue investigation**: Background on a reported problem (e.g., "Has this bug been reported before? What's the known workaround?")
- **Account context**: History with a specific customer (e.g., "What did we tell Acme Corp last time they asked about this?")
- **Topic research**: General topic relevant to support work (e.g., "Best practices for webhook retry logic")

Before searching, clarify what you're actually trying to find:
- Is this a factual question with a definitive answer?
- Is this a contextual question requiring multiple perspectives?
- Is this an exploratory question where the scope is still being defined?
- Who is the audience for the answer (internal team, customer, leadership)?

### 2. Search Available Sources

Search systematically through the source tiers below, adapting to what is connected. Don't stop at the first result — cross-reference across sources.

**Tier 1 — Official Internal Sources (highest confidence):**
- ~~knowledge base (Lark Wiki + Docs): `lark_doc_search(query=..., jq: ".data.results[] | {title,url}")` → `lark_doc_fetch(jq: ".data.content[:4000]")` for product docs, runbooks, FAQs, policy.
- ~~cloud storage (Lark Drive): internal specs/guides/past research; route Drive search/fetch through **`lark-drive`** / `lark_doc_fetch`.
- Product roadmap (internal-facing): if held in a Base, `lark_base_search` (pass `search_fields`; no `jq`); OKRs via `lark_okr_cycle_list(jq: ".data.cycles[]")`.

**Tier 2 — Organizational Context:**
- ~~CRM notes (HubSpot if connected; else account **Lark Base**): `lark_base_search(base_token, table_id, search_fields:["account"], select_fields:["account","notes","last_answer"])` for activity history and previous answers. base_search REQUIRES `search_fields` — if the field names are unknown, discover them first via `lark_api GET /open-apis/bitable/v1/apps/{base}/tables/{table}/fields`. No `jq` — narrow with `select_fields`/`limit`.
- ~~support platform (Intercom or ticket **Lark Base**): `lark_base_search` (pass `search_fields`; no `jq`) for previous resolutions, known issues, workarounds.
- Meeting notes: `lark_minutes_search(participant_ids="me", query=..., jq: ".data.items[]")` → pull AI summary/decisions instead of re-deriving (P6).

**Tier 3 — Team Communications:**
- ~~chat (Lark IM): `lark_im_search(query="<topic>", jq: ".data.messages[]")` — check if teammates already answered this.
- ~~email (Lark Mail): `lark_api` mail search for previous correspondence on the topic.
- Calendar: `lark_calendar_agenda(jq: ".data[]")` for meeting context; past meetings via `lark_vc_search`.

**Tier 4 — External Sources:**
- Web search: official documentation, blog posts, community forums
- Public knowledge bases, help centers, release notes
- Third-party documentation: integration partners, complementary tools

**Tier 5 — Inferred or Analogical (use when direct sources don't yield answers):**
- Similar situations: how similar questions were handled before
- Analogous customers: what worked for comparable accounts
- General best practices: industry standards and norms

### 3. Synthesize Findings

Compile results into a structured research brief:

```
## Research: [Question/Topic]

### Answer
[Clear, direct answer to the question — lead with the bottom line]

**Confidence:** [High / Medium / Low]
[Explain what drives the confidence level]

### Key Findings

**From [Source 1]:**
- [Finding with specific detail]
- [Finding with specific detail]

**From [Source 2]:**
- [Finding with specific detail]

### Context & Nuance
[Any caveats, edge cases, or additional context that matters]

### Sources
1. [Source name/link] — [what it contributed]
2. [Source name/link] — [what it contributed]
3. [Source name/link] — [what it contributed]

### Gaps & Unknowns
- [What couldn't be confirmed]
- [What might need verification from a subject matter expert]

### Recommended Next Steps
- [Action if the answer needs to go to a customer]
- [Action if further research is needed]
- [Who to consult for verification if needed]
```

### 4. Handle Insufficient Sources

If no connected sources yield results:

- Perform web research on the topic (WebSearch / WebFetch)
- Ask the user for internal context:
  - "I couldn't find this in connected sources. Do you have internal docs or Wiki articles about this?" (re-search with `lark_doc_search`)
  - "Has your team discussed this topic before? Any chat groups I should check?" (re-run `lark_im_search`)
  - "Is there a subject matter expert who would know the answer?" (resolve via `lark_contact_search`, then ask via `lark_im_send`)
- Be transparent about limitations:
  - "This answer is based on web research only — please verify against your internal documentation before sharing with the customer."
  - "I found a possible answer but couldn't confirm it from an authoritative internal source."

### 5. Customer-Facing Considerations

If the research is to answer a customer question:

- Flag if the answer involves product roadmap, pricing, legal, or security topics that may need review
- Note if the answer differs from what may have been communicated previously
- Suggest appropriate caveats for the customer-facing response
- Offer to draft the customer response: "Want me to draft a response to the customer based on these findings?"

### 6. Knowledge Capture

After research is complete, suggest capturing the knowledge in **Wiki** (P8) so the org finds it later — don't leave it as a chat message:

- "Should I save these findings to your Wiki for future reference?" → `lark_wiki_node_create(space_id, title)`, then fill the docx (delegate to **`lark-doc`**).
- "Want me to create a FAQ entry based on this research?" → `/kb-article`.
- "This might be worth documenting — should I draft a runbook entry?" → `lark_wiki_node_create`.

This helps build institutional knowledge and reduces duplicate research effort across the team.

---

## Source Prioritization and Confidence

### Confidence by Source Tier

| Tier | Source Type | Confidence | Notes |
|------|-------------|------------|-------|
| 1 | Official internal docs, KB, policies | **High** | Trust unless clearly outdated — check dates |
| 2 | CRM, support tickets, meeting notes | **Medium-High** | May be subjective or incomplete |
| 3 | Chat, email, calendar notes | **Medium** | Informal, may be out of context or speculative |
| 4 | Web, forums, third-party docs | **Low-Medium** | May not reflect your specific situation |
| 5 | Inference, analogies, best practices | **Low** | Clearly flag as inference, not fact |

### Confidence Levels

Always assign and communicate a confidence level:

**High Confidence:**
- Answer confirmed by official documentation or authoritative source
- Multiple sources corroborate the same answer
- Information is current (verified within a reasonable timeframe)
- "I'm confident this is accurate based on [source]."

**Medium Confidence:**
- Answer found in informal sources (chat, email) but not official docs
- Single source without corroboration
- Information may be slightly outdated but likely still valid
- "Based on [source], this appears to be the case, but I'd recommend confirming with [team/person]."

**Low Confidence:**
- Answer is inferred from related information
- Sources are outdated or potentially unreliable
- Contradictory information found across sources
- "I wasn't able to find a definitive answer. Based on [context], my best assessment is [answer], but this should be verified before sharing with the customer."

**Unable to Determine:**
- No relevant information found in any source
- Question requires specialized knowledge not available in sources
- "I couldn't find information about this. I recommend reaching out to [suggested expert/team] for a definitive answer."

### Handling Contradictions

When sources disagree:
1. Note the contradiction explicitly
2. Identify which source is more authoritative or more recent
3. Present both perspectives with context
4. Recommend how to resolve the discrepancy
5. If going to a customer: use the most conservative/cautious answer until resolved

## When to Escalate vs. Answer Directly

### Answer Directly When:
- Official documentation clearly addresses the question
- Multiple reliable sources corroborate the answer
- The question is factual and non-sensitive
- The answer doesn't involve commitments, timelines, or pricing
- You've answered similar questions before with confirmed accuracy

### Escalate or Verify When:
- The answer involves product roadmap commitments or timelines
- Pricing, legal terms, or contract-specific questions
- Security, compliance, or data handling questions
- The answer could set a precedent or create expectations
- You found contradictory information in sources
- The question involves a specific customer's custom configuration
- The answer requires specialized expertise you don't have
- The customer is at risk and the wrong answer could exacerbate the situation

### Escalation Path:
1. **Subject matter expert**: For technical or domain-specific questions
2. **Product team**: For roadmap, feature, or capability questions
3. **Legal/compliance**: For terms, privacy, security, or regulatory questions
4. **Billing/finance**: For pricing, invoice, or payment-related questions
5. **Engineering**: For custom configurations, bugs, or technical root causes
6. **Leadership**: For strategic decisions, exceptions, or high-stakes situations

## Research Documentation for Team Knowledge Base

After completing research, capture the knowledge for future use.

### When to Document:
- Question has come up before or likely will again
- Research took significant effort to compile
- Answer required synthesizing multiple sources
- Answer corrects a common misunderstanding
- Answer involves nuance that's easy to get wrong

### Documentation Format:
```
## [Question/Topic]

**Last Verified:** [date]
**Confidence:** [level]

### Answer
[Clear, direct answer]

### Details
[Supporting detail, context, and nuance]

### Sources
[Where this information came from]

### Related Questions
[Other questions this might help answer]

### Review Notes
[When to re-verify, what might change this answer]
```

### Knowledge Base Hygiene:
- Date-stamp all entries
- Flag entries that reference specific product versions or features
- Review and update entries quarterly
- Archive entries that are no longer relevant
- Tag entries for searchability (by topic, product area, customer segment)
