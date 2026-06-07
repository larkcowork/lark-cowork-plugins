---
name: enrich-lead
description: "Instant lead enrichment. Drop a name, company, LinkedIn URL, or email and get the full contact card with email, phone, title, company intel, and next actions."
user-invocable: true
argument-hint: "[name, company, LinkedIn URL, or email]"
---

# Enrich Lead

Turn any identifier into a full contact dossier. The user provides identifying info via "$ARGUMENTS".

> **Lark-native note**: Apollo stays the enrichment engine (it owns the data). Once a contact is enriched, the *collaboration* layer is Lark — log the lead to a Lark Base CRM (system-of-record), draft outreach as a Lark mail, and resolve any internal owner via `lark_contact_search`. See `../../connectors/LARK-PATTERNS.md` (P1/P4/P5) for the handoff contract.

## Examples

- `/apollo:enrich-lead Tim Zheng at Apollo`
- `/apollo:enrich-lead https://www.linkedin.com/in/timzheng`
- `/apollo:enrich-lead sarah@stripe.com`
- `/apollo:enrich-lead Jane Smith, VP Engineering, Lark Wiki`
- `/apollo:enrich-lead CEO of Figma`

## Step 1 — Parse Input

From "$ARGUMENTS", extract every identifier available:
- First name, last name
- Company name or domain
- LinkedIn URL
- Email address
- Job title (use as a matching hint)

If the input is ambiguous (e.g. just "CEO of Figma"), first use `mcp__claude_ai_Apollo_MCP__apollo_mixed_people_api_search` with relevant title and domain filters to identify the person, then proceed to enrichment.

## Step 2 — Enrich the Person

> **Credit warning**: Tell the user enrichment consumes 1 Apollo credit before calling.

Use `mcp__claude_ai_Apollo_MCP__apollo_people_match` with all available identifiers:
- `first_name`, `last_name` if name is known
- `domain` or `organization_name` if company is known
- `linkedin_url` if LinkedIn is provided
- `email` if email is provided
- Set `reveal_personal_emails` to `true`

If the match fails, try `mcp__claude_ai_Apollo_MCP__apollo_mixed_people_api_search` with looser filters and present the top 3 candidates. Ask the user to pick one, then re-enrich.

## Step 3 — Enrich Their Company

Use `mcp__claude_ai_Apollo_MCP__apollo_organizations_enrich` with the person's company domain to pull firmographic context.

## Step 4 — Present the Contact Card

Format the output exactly like this:

---

**[Full Name]** | [Title]
[Company Name] · [Industry] · [Employee Count] employees

| Field | Detail |
|---|---|
| Email (work) | ... |
| Email (personal) | ... (if revealed) |
| Phone (direct) | ... |
| Phone (mobile) | ... |
| Phone (corporate) | ... |
| Location | City, State, Country |
| LinkedIn | URL |
| Company Domain | ... |
| Company Revenue | Range |
| Company Funding | Total raised |
| Company HQ | Location |

---

## Step 5 — Offer Next Actions

Ask the user which action to take. The first two stay in Apollo (its data, its sequences); the rest are the Lark collaboration layer.

1. **Save to Apollo** — Create this person as a contact via `mcp__claude_ai_Apollo_MCP__apollo_contacts_create` with `run_dedupe: true`
2. **Add to a sequence** — Ask which sequence, then run the sequence-load flow
3. **Find colleagues** — Search for more people at the same company using `mcp__claude_ai_Apollo_MCP__apollo_mixed_people_api_search` with `q_organization_domains_list` set to this company
4. **Find similar people** — Search for people with the same title/seniority at other companies
5. **Log to Lark CRM (Base)** — Persist the enriched lead to a Lark Base as system-of-record (P5). First dedup-check with `lark_base_search`: it REQUIRES `search_fields` (the Bitable API mandates which field(s) to match on, e.g. the email or name column) and does NOT support jq — narrow with `select_fields` / `limit` instead. If you don't know the field names, discover them first via `lark_api` `GET /open-apis/bitable/v1/apps/{base}/tables/{table}/fields`. Then `lark_base_record_upsert` (`base_token`, `table_id`, `fields`: name, title, company, work/personal email, phone, location, LinkedIn, ICP notes) with `dry_run: true`; show the planned record, then commit. If no leads Base exists yet, scaffold one with the `base-deploy` skill rather than hand-rolling schema.
6. **Draft outreach in Lark Mail** — Draft (never auto-send) a first-touch email to the lead's work email with `lark_mail_draft_create` so the user reviews it in the Mail UI before sending. If an *internal* owner should be assigned, resolve them first with `lark_contact_search` to get their `open_id`. For the YAML/card grammar of a richer recap, delegate to the `lark-mail` / `lark-im` skills.
