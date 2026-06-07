# Lark Cowork Plugins

Turn **Claude Cowork** (in Claude Desktop) into a **Lark-native specialist** for every role —
chat, mail, calendar, docs/wiki, tasks, base, drive, minutes, all driven through the
`lark-cli` MCP bridge.

> 📘 **Người dùng nghiệp vụ (không kỹ thuật)? Bắt đầu ở đây →** [`docs/`](./docs/README.md)
> — giới thiệu, hướng dẫn bắt đầu, **danh mục theo phòng ban** (câu lệnh tiếng Việt mẫu),
> kịch bản thực tế, an toàn, FAQ, thuật ngữ.

> Customized from Anthropic's [knowledge-work-plugins](https://github.com/anthropics/knowledge-work-plugins).
> The original plugins target Slack / Notion / Google / Jira. This fork rewires the
> collaboration layer to **Lark (larksuite.com)** while keeping genuinely specialty tools
> (warehouses, CRMs, payments, scientific DBs) as optional external connectors.

## How it works

Every plugin is just markdown + JSON. The skills are **tool-agnostic** — they reference
`~~category` placeholders that resolve through each plugin's `CONNECTORS.md`. We changed only
two things per plugin:

1. **`.mcp.json`** — the generic-comms servers (Slack, Gmail, Google Calendar, Notion, Asana,
   Linear, Atlassian, Guru, Fireflies, Box…) are replaced by a single **`lark`** MCP server
   (`lark-cli mcp serve`). Specialty servers are kept.
2. **`CONNECTORS.md`** — each generic category now maps to concrete `lark_*` tools; specialty
   categories stay external.

One `lark` server backs every collaboration category. Anything without a dedicated tool falls
through to `lark_api` (the Lark OpenAPI escape hatch). See
[connectors/CONNECTORS.lark.md](./connectors/CONNECTORS.lark.md) for the master mapping.

## Coverage (15 knowledge-work plugins)

| Plugin | Lark-native categories | Kept external (optional) |
|--------|------------------------|--------------------------|
| productivity | chat, mail, calendar, wiki, task/base, drive | — |
| enterprise-search | chat, mail, wiki, task, drive | — |
| operations | chat, mail, calendar, wiki, task | itsm, procurement |
| human-resources | chat, mail, calendar, wiki | ats, hris, comp data |
| customer-support | chat, mail, wiki, task, drive | intercom, hubspot |
| sales | chat, mail, calendar, wiki, task, minutes | hubspot, close, clay, zoominfo, apollo, outreach, similarweb |
| product-management | chat, mail, calendar, wiki, task, minutes | figma, amplitude, pendo, intercom, similarweb |
| marketing | chat, wiki | canva, figma, hubspot, amplitude, ahrefs, similarweb, klaviyo, supermetrics |
| legal | chat, mail, calendar, office, task, drive | docusign |
| design | chat, wiki, task | figma, intercom |
| data | task | snowflake, databricks, bigquery, hex, amplitude, definite |
| finance | chat, mail, office | snowflake, databricks, bigquery |
| engineering | chat, wiki, task | github, pagerduty, datadog |
| bio-research | (comms via lark) | pubmed, biorender, biorxiv, consensus, clinicaltrials, chembl, synapse, wiley, owkin, open-targets, benchling |
| small-business | chat, mail, calendar, docs, drive | quickbooks, paypal, stripe, square, hubspot, canva, docusign |

`pdf-viewer` and `cowork-plugin-management` are kept unchanged from upstream.

## Lark-native plugins (7 new, 100% Lark — no external connectors)

These were ported from rough `lark-cli` skills into full plugins following the same convention
(`plugin.json` + `.mcp.json` lark + `CONNECTORS.md` + `README.md`). Every category resolves to the
single `lark` MCP server — no specialty servers, zero external dependencies.

| Plugin | What it does | Skills |
|--------|--------------|--------|
| lark-base-deploy | Deploy a Lark Base end-to-end (8-phase orchestrator, parallel fan-out) | 9 |
| daily-assistant | Run your day — briefs, digests, triage, focus, meeting/1:1 prep, contact 360 | 12 |
| crm-sales | CRM on Lark Base — pipeline review, post-call deal updates, dormant follow-ups (drafts only) | 3 |
| governance | Approval triage, approval-flow SLA, permission audit, decision logging | 4 |
| knowledge-docs | Author docs/wiki/sheets from templates; audit & restructure a Wiki | 2 |
| delivery-eng | Blameless incident postmortems + end-of-sprint retros | 2 |
| lark-cli-dev | Build/debug/extend the `lark-cli` MCP bridge (dev tooling; 1 skill + 6 `/mcp-*` commands) | 1 |

## Prerequisites

The `lark` connector calls `lark-cli`. Install + authenticate once:

```bash
# 1. Build/install lark-cli (from the lark-cli repo) onto PATH, e.g.:
#    go build -o ~/bin/lark-cli . && export PATH="$HOME/bin:$PATH"
lark-cli --version

# 2. Authenticate (stores credentials in the macOS Keychain)
lark-cli auth login
```

## Choosing a transport

The plugins ship with the **stdio** profile by default. Switch any time:

```bash
./set-transport.sh stdio   # Claude Desktop classic — local lark-cli binary + Keychain
./set-transport.sh http    # Cowork VM / remote — lark-cli mcp serve --transport http + tunnel
```

- **stdio** — simplest. Needs `lark-cli` on PATH and `lark-cli auth login` done on this machine.
- **http** — for Cowork's sandboxed runtime, which can't reach a host binary. Run a remote bridge
  and point the plugins at it:

  ```bash
  lark-cli mcp serve --transport http --addr 127.0.0.1:3000 \
    --audit-log ~/.lark-mcp-audit.ndjson
  cloudflared tunnel --url http://127.0.0.1:3000     # -> public https URL
  export LARK_MCP_URL="https://<your>.trycloudflare.com/mcp"
  export LARK_MCP_BEARER_TOKEN="<a-strong-secret>"   # also pass to serve
  ./set-transport.sh http
  ```

## Install into Claude

**Claude Code**
```bash
claude plugin marketplace add /Users/jimmy/Downloads/2026-NEW/GO!/lark-cowork-plugins
claude plugin install productivity@lark-cowork
```

**Cowork (Claude Desktop)** — add this folder as a local marketplace, then install plugins from
the Cowork plugin UI. Use the `http` transport profile.

## Re-running the conversion

The converter is a small Go program (`tools/larkify.go`) — same language as the rest of the
stack, with `go test` coverage. Reproducible and idempotent:

```bash
cd tools
go test ./...                                   # unit tests
go run . convert <plugin> [<plugin> ...]        # rewrite .mcp.json + CONNECTORS.md + terminology
go run . transport stdio|http                   # switch the lark transport across all plugins
# or build once:  go build -o larkify .  &&  ./larkify convert <plugin>
```

See [SETUP.md](./SETUP.md) for the full setup + verification checklist.

## Attribution

Built on Anthropic's open-source knowledge-work-plugins (see [LICENSE](./LICENSE)). Lark, Feishu,
and the connected third-party services are trademarks of their respective owners.
