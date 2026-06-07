---
name: email-sequence
description: Design and draft multi-email sequences with full copy, timing, branching logic, exit conditions, and performance benchmarks. Use when building onboarding, lead nurture, re-engagement, win-back, or product launch flows, when you need a complete drip campaign with A/B test suggestions, or when mapping a sequence end-to-end with a flow diagram.
argument-hint: "[sequence type]"
---

# Email Sequence

> If you see unfamiliar placeholders or need to check which tools are connected, see [CONNECTORS.md](../../CONNECTORS.md).

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md), [LARK-RECIPES](../../../connectors/LARK-RECIPES.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md)). The drip *automation engine* stays the specialty `~~email marketing` tool (Klaviyo/Customer.io) — but the authoring, review, tracking, and rollout are Lark:
> - **Brand voice + reusable assets from Lark first** (P3): `lark_doc_search`/`lark_doc_fetch` for the style guide and existing emails; `lark_base_search` for product facts (requires `search_fields`; discover field names via `lark_api GET /open-apis/bitable/v1/apps/{base}/tables/{table}/fields` if unknown; narrow with `select_fields`/`limit`, not jq).
> - **Each email body → a Lark Mail draft** for review (P2): `lark_mail_draft_create` (drafts only — never `lark_mail_send` with `confirm_send` unsolicited). The reviewer eyeballs it in the Mail UI.
> - **The sequence map + benchmarks → Wiki** (P8, `lark_wiki_node_create` via **`lark-doc`**) so the team has one durable spec.
> - **Build/QA work → Lark tasks** (P1/P2): "set up flow in Klaviyo", "configure exit conditions" → resolve owner via `lark_contact_search` → `lark_task_create(dry_run)` → confirm.
> - **Track live performance in a Lark Base** (P5): scaffold a "Sequence Performance" Base via **`base-deploy`** (per-email open/CTR/conversion/unsub); upsert with `lark_base_record_upsert`. Surface results as a card (P4).

Design and draft complete email sequences with full copy, timing, branching logic, and performance benchmarks for any lifecycle or campaign use case.

## Trigger

User runs `/email-sequence` or asks to create, design, build, or draft an email sequence, drip campaign, nurture flow, or onboarding series.

## Inputs

Gather the following from the user. If not provided, ask before proceeding:

1. **Sequence type** — one of:
   - Onboarding
   - Lead nurture
   - Re-engagement
   - Product launch
   - Event follow-up
   - Upgrade/upsell
   - Win-back
   - Educational drip

2. **Goal** — what the sequence should achieve (e.g., activate new users, convert leads to customers, reduce churn, drive event attendance, upsell to a higher tier)

3. **Audience** — who receives this sequence, what stage they are at, and any relevant segmentation details (role, industry, behavior triggers, lifecycle stage)

4. **Number of emails** (optional) — if not specified, recommend a count based on the sequence type using the templates in the Sequence Type Templates section below

5. **Timing/cadence preferences** (optional) — desired spacing between emails (e.g., "every 3 days", "weekly", "aggressive first week then taper off")

6. **Brand voice** — if configured in local settings, apply automatically and inform the user. If not configured, ask: "Do you have brand voice guidelines I should follow? If not, I'll use a clear, conversational professional tone."

7. **Additional context** (optional):
   - Specific offers, discounts, or incentives to include
   - CTAs or landing pages to link to
   - Content assets available (blog posts, case studies, videos, guides)
   - Product features to highlight
   - Competitor differentiators to reference

## Process

### 1. Sequence Strategy

Before drafting any emails, define the overall sequence architecture:

- **Narrative arc** — what story does this sequence tell across all emails? What is the emotional and logical progression from first email to last?
- **Journey mapping** — map each email to a stage of the buyer or user journey (awareness, consideration, decision, activation, expansion)
- **Escalation logic** — how does the intensity, urgency, or value of each email build on the previous one?
- **Success definition** — what specific action signals that the sequence has done its job and the recipient should exit?

### 2. Individual Email Design

For each email in the sequence, produce:

#### Subject Line
- Provide 2-3 options per email
- Vary approaches: curiosity, benefit-driven, urgency, personalization, question-based
- Keep under 50 characters where possible; note preview behavior on mobile

#### Preview Text
- 40-90 characters that complement (not repeat) the subject line
- Should add context or intrigue that increases open likelihood

#### Email Purpose
- One sentence explaining why this email exists and what it moves the recipient toward

#### Body Copy
- Full draft ready to use
- Clear hierarchy: hook, body, CTA
- Short paragraphs (2-3 sentences max)
- Scannable formatting with bold key phrases where appropriate
- Personalization tokens where relevant (e.g., first name, company name, product used)

#### Primary CTA
- Button text and destination
- One primary CTA per email (secondary CTA only if appropriate for the sequence stage)

#### Timing
- Days after the trigger event or after the previous email
- Note if timing should adjust based on engagement (e.g., "send sooner if they opened but did not click")

#### Segment/Condition Notes
- Who receives this email vs. who skips it
- Any behavioral or attribute-based conditions (e.g., "only send to users who have not completed setup")

### 3. Sequence Logic

Define the flow control for the sequence:

- **Branching conditions** — alternate paths based on engagement. For example:
  - "If opened email 2 but did not click CTA, send email 2b (softer re-ask) instead of email 3"
  - "If clicked CTA in email 1, skip email 2 and go directly to email 3"
- **Exit conditions** — when a recipient converts (completes the desired action), remove them from the sequence. Define what "conversion" means for this sequence.
- **Re-entry rules** — can someone re-enter the sequence? Under what conditions? (e.g., "if a user churns again 90 days later, re-enter the win-back sequence")
- **Suppression rules** — do not send if the recipient is already in another active sequence, has unsubscribed from marketing, or has contacted support in the last 48 hours

### 4. Performance Benchmarks

Provide expected benchmarks based on the sequence type so the user can set targets:

| Metric | Onboarding | Lead Nurture | Re-engagement | Win-back |
|--------|-----------|--------------|---------------|----------|
| Open rate | 50-70% | 20-30% | 15-25% | 15-20% |
| Click-through rate | 10-20% | 3-7% | 2-5% | 2-4% |
| Conversion rate | 15-30% | 2-5% | 3-8% | 1-3% |
| Unsubscribe rate | <0.5% | <0.5% | 1-2% | 1-3% |

Adjust benchmarks based on industry and audience if the user has provided that context.

## Sequence Type Templates

Use these as starting frameworks. Adapt length and content based on the user's goal and audience.

**Onboarding (5-7 emails over 14-21 days):**
Welcome and set expectations -- Quick win to demonstrate value -- Core feature deep dive -- Advanced feature or integration -- Social proof and community -- Check-in and feedback request -- Upgrade prompt or next steps

**Lead Nurture (4-6 emails over 3-4 weeks):**
Value-first educational content -- Pain point identification -- Solution positioning with proof -- Social proof and results -- Soft CTA (trial, demo, resource) -- Direct CTA (buy, book, sign up)

**Re-engagement (3-4 emails over 10-14 days):**
"We miss you" with a compelling reason to return -- Value reminder highlighting what they are missing -- Incentive or exclusive offer -- Last chance with clear deadline

**Win-back (3-5 emails over 30 days):**
Friendly check-in asking what went wrong -- What is new since they left -- Special offer or incentive to return -- Feedback request (even if they do not come back) -- Final goodbye with door open

**Product Launch (4-6 emails over 2-3 weeks):**
Teaser or pre-announcement -- Launch announcement with full details -- Feature spotlight or use case -- Social proof and early results -- Limited-time offer or bonus -- Last chance or reminder

**Event Follow-up (3-4 emails over 7-10 days):**
Thank you with key takeaways or recordings -- Resource roundup from the event -- Related offer or next step -- Feedback survey

**Upgrade/Upsell (3-5 emails over 2-3 weeks):**
Usage milestone or success celebration -- Feature gap or limitation they are hitting -- Upgrade benefits with proof -- Limited-time incentive -- Direct comparison of plans

**Educational Drip (5-8 emails over 4-6 weeks):**
Introduction and what they will learn -- Lesson 1: foundational concept -- Lesson 2: intermediate concept -- Lesson 3: advanced concept -- Practical application or exercise -- Resource roundup -- Graduation and next steps

## Tool Integration

### If ~~email marketing is connected (e.g., Klaviyo, Mailchimp, Customer.io)
- Reference how to set up the sequence as a flow or automation in the platform
- Note any platform-specific features to use (e.g., smart send time, conditional splits, A/B testing)
- Map the branching logic to the platform's visual flow builder concepts

### If ~~marketing automation or ~~CRM is connected (e.g., HubSpot, Marketo)
- Reference lead scoring data to inform segmentation and exit conditions
- Use lifecycle stage data to tailor messaging per segment
- Note how to set enrollment triggers based on CRM properties or list membership

### If no specialty email-marketing tool is connected
- Deliver all email content in copy-paste-ready format, and offer to create each as a **Lark Mail draft** (`lark_mail_draft_create`, P2) so the user can review/send from the Mail UI.
- Include a setup checklist the user can follow in any email platform — and offer to turn each step into a **Lark task** (`lark_task_create`, resolve owner via `lark_contact_search`):
  1. Create the automation or flow
  2. Set the enrollment trigger
  3. Add each email with the specified delays
  4. Configure branching and exit conditions
  5. Set up tracking for the recommended metrics (mirror into a "Sequence Performance" Base, P5)

## Output

Present the complete sequence with the following sections:

### Sequence Overview Table

| # | Subject Line | Purpose | Timing | Primary CTA | Condition |
|---|-------------|---------|--------|-------------|-----------|

### Full Email Drafts
Each email with subject line options, preview text, purpose, body copy, CTA, timing, and segment notes.

### Sequence Flow Diagram
A text-based diagram showing the email flow, branching paths, and exit points. Use a clear format such as:

```
[Trigger] --> Email 1 (Day 0)
                |
          Opened? --Yes--> Email 2 (Day 3)
                |              |
                No        Clicked CTA? --Yes--> [EXIT: Converted]
                |              |
                v              No
          Email 1b (Day 2)     |
                |              v
                +--------> Email 3 (Day 7)
                               |
                               v
                          Email 4 (Day 10)
                               |
                          [EXIT: Sequence complete]
```

### Branching Logic Notes
Summary of all conditions, exits, and suppressions in a reference list.

### A/B Test Suggestions
- 2-3 recommended A/B tests (subject lines, CTA text, send time, email length)
- What to test, how to split, and how to measure the winner

### Metrics to Track
- Primary conversion metric for the sequence
- Per-email metrics: open rate, CTR, unsubscribe rate
- Sequence-level metrics: overall conversion rate, time to conversion, drop-off points
- Recommended review cadence (e.g., "Review performance weekly for the first month, then monthly")

## After the Sequence

Offer to land the sequence natively: each email as a **Lark Mail draft** (P2), the spec + benchmarks to **Wiki** (P8), setup steps as **Lark tasks** (P1), and a **Sequence Performance Base** for tracking (P5).

Ask: "Would you like me to:
- Create Lark Mail drafts for each email and publish the spec to Wiki?
- Revise the copy or tone for any specific email?
- Add a branching path for a specific scenario?
- Create a variation of this sequence for a different audience segment?
- Draft the A/B test variants for the subject lines?
- Build a companion sequence (e.g., a post-purchase follow-up after this lead nurture converts)?"
