# Lark Cowork — Setup & Verification

End-to-end checklist to take this marketplace from "files on disk" to "working in Claude".

## 0. State of this machine

The `lark-cli` binary, its Keychain auth, and the Claude Desktop MCP entry were **wiped** during
an earlier demo cleanup. So the very first runtime test requires reinstalling lark-cli. The plugin
files themselves are complete and validated (JSON + connector mapping); only the live runtime
needs the binary back.

## 1. Reinstall + authenticate lark-cli

```bash
cd /Users/jimmy/Downloads/2026-NEW/GO!/lark-cli
go build -o ~/bin/lark-cli .          # or the repo's Makefile target
export PATH="$HOME/bin:$PATH"
lark-cli --version
lark-cli auth login                    # opens browser; stores creds in Keychain
lark-cli auth status                   # expect: user (and/or bot) = ready
```

## 2. Pick a transport

```bash
cd /Users/jimmy/Downloads/2026-NEW/GO!/lark-cowork-plugins
./set-transport.sh stdio               # Desktop classic (recommended to start)
# or ./set-transport.sh http           # Cowork VM / remote
```

For `http`, also start the remote bridge + tunnel and export `LARK_MCP_URL` +
`LARK_MCP_BEARER_TOKEN` (see README "Choosing a transport").

## 3. Smoke-test the bridge directly (no Claude yet)

```bash
lark-cli mcp tools                      # lists the lark_* tool surface (offline)
# Optional: initialize + a real call via your mcp-test / mcp-call helpers
```

## 4. Install into Claude

- **Claude Code:** `claude plugin marketplace add .` then
  `claude plugin install productivity@lark-cowork`.
- **Cowork (Desktop):** add this folder as a local marketplace; install from the plugin UI.

## 5. Verify in the UI (do NOT trust exit code 0)

Per the project rule: a feature isn't done until visually verified. Run a real workflow and
confirm Claude actually calls `lark_*`:

1. In a plugin's working dir, run `/productivity:start` (or invoke a skill).
2. Watch for tool calls to `lark_im_*`, `lark_doc_*`, `lark_task_*`, etc.
3. Cross-check the audit log if running http: `~/.lark-mcp-audit.ndjson` records each call.

## Structural verification already done

```bash
# all JSON valid
python3 -c "import json,glob;[json.load(open(f)) for f in glob.glob('**/*.json',recursive=True)];print('ok')"
# every knowledge-work plugin has the lark connector
grep -rl '\"lark\"' */.mcp.json | wc -l            # -> 15
# no generic SaaS servers remain
grep -hoE '\"(slack|notion|gmail|google calendar|asana|linear|atlassian)\"' */.mcp.json   # -> empty
```

## Known gaps / open questions

- **R1 — Cowork stdio reachability.** Unverified whether Cowork's VM can spawn a host stdio
  binary. If not, `http` transport is mandatory for Cowork. (Desktop classic stdio is proven.)
- **R2 — single-server resolution.** All categories map to one `lark` server; CONNECTORS.md
  names the exact tools so Claude routes correctly. Confirm with a multi-tool workflow.
- **R3 — CRM/analytics.** Lark has no native CRM/product-analytics; those stay external or get
  modeled in Lark Base (use the `base-deploy` skill to scaffold a CRM Base).
- **R4 — tool surface.** The bridge exposes 25 curated tools + `lark_api`. Ops outside that set
  go through `lark_api` and may need extra Lark app scopes.
- **R5 — Lark vs Feishu.** Configured for Lark international (larksuite.com). Feishu (feishu.cn)
  would need different endpoints/scopes.
