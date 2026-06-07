---
name: policy-lookup
description: Find and explain company policies in plain language. Trigger with "what's our PTO policy", "can I work remotely from another country", "how do expenses work", or any plain-language question about benefits, travel, leave, or handbook rules.
argument-hint: "<policy topic — PTO, benefits, travel, expenses, etc.>"
---

# /policy-lookup

> If you see unfamiliar placeholders or need to check which tools are connected, see [CONNECTORS.md](../../CONNECTORS.md).

Look up and explain company policies in plain language. Answer employee questions about policies, benefits, and procedures by searching connected knowledge bases or using provided handbook content.

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md),
> [LARK-RECIPES](../../../connectors/LARK-RECIPES.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md)).
> The `~~knowledge base` is **Lark Wiki + Docs**: find the policy with `lark_doc_search`, then read
> the matching doc with `lark_doc_fetch` projected via `jq` so you cite the exact section without
> burning 25k tokens (P3) — delegate Wiki structure/navigation to **`lark-wiki`** and doc reads to
> **`lark-doc`**. Employee-specific facts (PTO balance, benefits elections) come from the `~~HRIS`
> specialty MCP (or its **Lark Base** approximation, `lark_base_search`); resolve the asking employee
> with `lark_contact_search` (P1) and DM them their numbers privately. Always quote the source doc +
> section and link the Wiki node. For complex/multi-part answers, post an **interactive card**
> (`lark_im_card_send`, P4) with collapsible `panel` sections and a "who to contact" footer; simple
> answers can be `lark_im_send`.

## Usage

```
/policy-lookup $ARGUMENTS
```

Search for policies matching: $ARGUMENTS

## How It Works

```
┌─────────────────────────────────────────────────────────────────┐
│                    POLICY LOOKUP                                   │
├─────────────────────────────────────────────────────────────────┤
│  STANDALONE (always works)                                       │
│  ✓ Ask any policy question in plain language                    │
│  ✓ Paste your employee handbook and I'll search it              │
│  ✓ Get clear, jargon-free answers                               │
├─────────────────────────────────────────────────────────────────┤
│  SUPERCHARGED (when you connect your tools)                      │
│  + Knowledge base: Search handbook and policy docs automatically │
│  + HRIS: Pull employee-specific details (PTO balance, benefits) │
└─────────────────────────────────────────────────────────────────┘
```

## Common Policy Topics

- **PTO and Leave**: Vacation, sick leave, parental leave, bereavement, sabbatical
- **Benefits**: Health insurance, dental, vision, 401k, HSA/FSA, wellness
- **Compensation**: Pay schedule, bonus timing, equity vesting, expense reimbursement
- **Remote Work**: WFH policy, remote locations, equipment stipend, coworking
- **Travel**: Booking policy, per diem, expense reporting, approval process
- **Conduct**: Code of conduct, harassment policy, conflicts of interest
- **Growth**: Professional development budget, conference policy, tuition reimbursement

## How to Answer

1. Search the `~~knowledge base` (Lark Wiki/Docs) with `lark_doc_search`; open the best match with
   `lark_doc_fetch` (`jq` to pull just the relevant section, P3)
2. Provide a clear, plain-language answer
3. Quote the specific policy language and link the Wiki node
4. Note any exceptions or special cases
5. Point to who to contact for edge cases (resolve via `lark_contact_search`, P1)

**Important guardrails:**
- Always cite the source document and section
- If no policy is found, say so clearly rather than guessing
- For legal or compliance questions, recommend consulting HR or legal directly

## Output

```markdown
## Policy: [Topic]

### Quick Answer
[1-2 sentence direct answer to their question]

### Details
[Relevant policy details, explained in plain language]

### Exceptions / Special Cases
[Any relevant exceptions or edge cases]

### Who to Contact
[Person or team for questions beyond what's documented]

### Source
[Where this information came from — document name, page, or section]
```

## If Connectors Available

If **~~knowledge base** (Lark Wiki + Docs) is connected:
- Search the handbook/policy docs automatically with `lark_doc_search`; read with `lark_doc_fetch`
  (`jq` projection, P3)
- Cite the specific document + section and link the Wiki node (delegate to **`lark-wiki`** /
  **`lark-doc`**)

If **~~HRIS** (specialty external MCP, or a **Lark Base**) is connected:
- Pull employee-specific details (PTO balance, benefits elections, enrollment) — `lark_base_search`
  if approximated in a Base; DM the result privately via `lark_im_send` after `lark_contact_search` (P1)

## Tips

1. **Ask in plain language** — "Can I work from Europe for a month?" is better than "international remote work policy."
2. **Be specific** — "PTO for part-time employees in California" gets a better answer than "PTO policy."
