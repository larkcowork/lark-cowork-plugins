# QC & Audit — Lark Cowork marketplace + lark-cli MCP bridge

**Status: ✅ ALL GREEN — 24/24 automated checks pass** (run `./tools/audit.sh --live`).
Last run: 24 passed, 0 failed. Reproducible harness: [`tools/audit.sh`](./tools/audit.sh).

---

## 1. Plan & scope

**Objective:** verify the Lark Cowork marketplace (22 plugins, 235 skills, forked from
anthropics/knowledge-work-plugins) and the lark-cli MCP bridge (25 tools) are correct, consistent,
and actually work — not just "JSON parses".

**Out of scope:** the upstream bio/zoom specialty payloads (kept external), Cowork-VM transport
(needs the cloud runtime), and Feishu-CN endpoints (configured for Lark international).

**Method:** 4 verification levels, each a set of automated test cases in `tools/audit.sh`:

| Level | Question it answers | Mode |
|-------|--------------------|------|
| **L1 Structural** | Do the files/manifests hold together? | static |
| **L2 Correctness** | Are tool names / jq paths / requirements right? | static + source |
| **L3 Semantic** | Is the soul-swap real & consistent (not cosmetic)? | static + diff vs upstream |
| **L4 Runtime** | Does it actually build, serve, and call Lark? | build + MCP handshake + live API |

---

## 2. Test-case matrix (automated — IDs match `tools/audit.sh` output)

### L1 — Structural
| ID | Test | Expected | Result |
|----|------|----------|--------|
| L1.1 | All `*.json` parse | valid JSON | ✅ |
| L1.2 | Every marketplace plugin `source` dir exists | no dangling entry | ✅ |
| L1.3 | Every `plugin.json` has name/version/description | complete | ✅ |
| L1.4 | No `CONNECTORS.md` with empty table (partner incl.) | all have rows | ✅ |
| L1.5 | No broken card `kind:` grammar | 0 occurrences | ✅ |

### L2 — Correctness
| ID | Test | Expected | Result |
|----|------|----------|--------|
| L2.1 | No hallucinated `lark_*` tool referenced in skills | only the 25 (+`lark_api`) | ✅ |
| L2.2 | No wrong top-level jq paths in skills | all under `.data.` | ✅ |
| L2.3 | Depth-core docs present (PATTERNS/RECIPES/FUSION) | 3 files | ✅ |
| L2.4 | No stale "21 tools" in docs | says 25 | ✅ |
| L2.5 | `toolBaseSearch` enforces `search_fields` | Build errors clearly | ✅ |

### L3 — Semantic (soul-swap quality)
| ID | Test | Expected | Result |
|----|------|----------|--------|
| L3.1 | No SKILL renamed vs upstream `/tmp/knowledge-work-plugins` | 0 renames | ✅ |
| L3.2 | Frontmatter intact (`---` + exactly one `name:`) | clean | ✅ |
| L3.3 | All first-party SKILL.md soul-swapped (depth-core ref) | 100% | ✅ |
| L3.4 | No agent meta-commentary artifacts in skills | 0 | ✅ |
| L3.5 | Every plugin CONNECTORS points to depth core | all | ✅ |

### L4 — Runtime
| ID | Test | Expected | Result |
|----|------|----------|--------|
| L4.1 | `go vet` (cmd/mcp + okr) | clean | ✅ |
| L4.2 | `go test` (cmd/mcp + okr) | pass | ✅ |
| L4.3 | `go build` → `~/bin/lark-cli` | builds | ✅ |
| L4.4 | `mcp tools` (offline) | 25 tools | ✅ |
| L4.5 | MCP stdio handshake `tools/list` | 25 tools | ✅ |
| L4.6 | Interactive card (P4) compiles offline | `ok:true` | ✅ |
| L4L.0 | `auth status` ready | user+bot ready | ✅ |
| L4L.1 | LIVE read: `lark_contact_search me` → open_id | real data | ✅ |
| L4L.2 | LIVE mutating round-trip: task create→complete→delete | created, completed, gone | ✅ |

> L4L.* run only with `--live` (needs `lark-cli auth login`). The mutating round-trip touches only a
> self-created `[AUDIT]` task and deletes it — **zero impact on existing data**.

---

## 3. Use cases (end-to-end scenarios) & verification status

| UC | Scenario | Tools exercised | How verified | Status |
|----|----------|-----------------|--------------|--------|
| UC1 | "What's on my plate / who am I" | `lark_contact_search`, `lark_task_my` | L4L.1 live read (real profile) | ✅ |
| UC2 | Capture a follow-up as a task, then close it | `lark_task_create`, `lark_task_complete` | L4L.2 live round-trip | ✅ |
| UC3 | Preview a message/event before sending (safe-mutation P2) | `lark_im_send`/`lark_calendar_create` `dry_run` | live dry_run (no write) — earlier session | ✅ |
| UC4 | Find records in a CRM/tracker Base | `lark_base_search` (+`search_fields`) | live: query+search_fields → record | ✅ |
| UC5 | Write/update a Base record (SoR, P5) | `lark_base_record_upsert` | live: `created:true`+record_id | ✅ |
| UC6 | Land durable knowledge in Wiki (P8) | `lark_wiki_node_create` | live create→delete round-trip | ✅ |
| UC7 | Check availability before scheduling | `lark_calendar_freebusy` | live read `ok:true` | ✅ |
| UC8 | Surface a decision/digest as an interactive card (P4) | `lark_im_card_send` | offline `print_json` `ok:true` (L4.6) | ✅ |
| UC9 | Pull meeting action items (P6) | `lark_minutes_search` | live read (`.data.items`) | ✅ |

---

## 4. QC checklist (sign-off)

- [x] Plumbing: 22/22 marketplace plugins resolve; 20 carry the `lark` connector; pdf-viewer intentional
- [x] Connectors: every plugin CONNECTORS.md has a real table + depth-core pointer
- [x] Tools: 25 MCP tools, all flag names verified vs `shortcuts/`, `go test` green
- [x] jq: all examples/skills project under `.data.` (verified the envelope is `{ok,data,…}`)
- [x] Soul-swap: 14 first-party (135 skills) + productivity (4/4) + 14 partner collab skills; 0 renames
- [x] Bugs fixed & re-verified: card `kind:` grammar, okr `"me"` resolver, `base_search` (search_fields + format + jq)
- [x] Runtime: build + MCP handshake (25) + card compile + live reads + live mutating round-trips
- [x] Data hygiene: all `[TEST-AUDIT]`/`[AUDIT]` throwaway artifacts deleted (0 leftover)
- [x] Docs: README, SETUP, blueprint, RECIPES say 25 tools; LARK-RECIPES documents base_search gotchas

---

## 5. Parallel deep-QC pass (21 agents)

A workflow of 20 per-plugin QC agents + 1 tool-shape agent reviewed **every** skill and fixed
**100 issues** across 20 plugins (base_search now carries `search_fields` + "no jq" in all examples;
jq projections under `.data.`; card grammar; no hallucinated tools). Re-verified: `audit.sh --live`
still **24/24 green**, `go test` green. Then **closed both prior limitations**:

- ✅ **`vc_search` shape VERIFIED LIVE** — a date-ranged search returned 15 real past meetings; real
  shape is `.data.items[]` (each item: id, display_info, meta_data) — the earlier `.data.meetings[]`
  example was wrong and is now fixed in `tools.go`.
- ✅ **`sheets_read` shape confirmed** correct (no change needed).

## 6. Known limitations (honest)

- **Cowork-VM transport** (http) not runtime-tested here — only Desktop-classic stdio is proven.
- **`base_search` requires `search_fields`** by API design — skills now state this + how to discover
  field names; the tool errors clearly if omitted.

---

## 7. How to re-run

```bash
cd lark-cowork-plugins
./tools/audit.sh            # L1–L4 offline (static + build + handshake + card)
./tools/audit.sh --live     # + auth-gated live reads & a self-cleaning task round-trip
# exit code 0 = all green; prints PASS/FAIL per test id
```
