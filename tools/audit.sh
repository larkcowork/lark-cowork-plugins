#!/usr/bin/env bash
# audit.sh — 4-level QC harness for the Lark Cowork marketplace + lark-cli MCP bridge.
#
#   ./tools/audit.sh           # L1-L4 offline (static + build + MCP handshake + card compile)
#   ./tools/audit.sh --live    # also run auth-gated live tests (reads + safe task round-trip)
#
# Exit code 0 = all PASS. Each test prints: "PASS|FAIL  <id>  <desc>".
set -uo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
CLI="${LARK_CLI:-$HOME/bin/lark-cli}"
REPO="${LARK_REPO:-/Users/jimmy/Downloads/2026-NEW/GO!/lark-cli}"
LIVE=0; [ "${1:-}" = "--live" ] && LIVE=1
PASS=0; FAIL=0
ok(){ printf "PASS  %-6s %s\n" "$1" "$2"; PASS=$((PASS+1)); }
no(){ printf "FAIL  %-6s %s\n" "$1" "$2"; FAIL=$((FAIL+1)); }
chk(){ if [ "$1" = "0" ]; then ok "$2" "$3"; else no "$2" "$3"; fi; }
cd "$ROOT"

echo "######## L1 — STRUCTURAL ########"
python3 -c "import json,glob,sys; [json.load(open(f)) for f in glob.glob('**/*.json',recursive=True)]" 2>/dev/null; chk $? L1.1 "all JSON files parse"
python3 -c "import json,os,sys; d=json.load(open('.claude-plugin/marketplace.json')); sys.exit(0 if all(os.path.isdir(p['source'].lstrip('./')) for p in d['plugins']) else 1)"; chk $? L1.2 "every marketplace plugin source dir exists"
python3 -c "import json,glob,sys; bad=[f for f in glob.glob('*/.claude-plugin/plugin.json')+glob.glob('partner-built/*/.claude-plugin/plugin.json') if any(k not in json.load(open(f)) for k in ('name','version','description'))]; sys.exit(1 if bad else 0)"; chk $? L1.3 "every plugin.json has name/version/description"
EMPTY=$(for c in */CONNECTORS.md partner-built/*/CONNECTORS.md; do [ -f "$c" ]||continue; r=$(grep -cE '^\| ' "$c"); [ "$r" -lt 3 ] && echo "$c"; done | grep -v pdf-viewer | wc -l | tr -d ' '); chk $([ "$EMPTY" = 0 ]&&echo 0||echo 1) L1.4 "no CONNECTORS.md with empty table (partner incl)"
KIND=$(grep -rlE 'kind: (item|actions|md|note|div|panel|columns|select)' --include='*.md' . | wc -l | tr -d ' '); chk $([ "$KIND" = 0 ]&&echo 0||echo 1) L1.5 "no broken card 'kind:' grammar"

echo "######## L2 — CORRECTNESS ########"
"$CLI" mcp tools 2>/dev/null | grep -oE 'lark_[a-z_]+' | sort -u > /tmp/audit_real.txt
grep -rhoE 'lark_[a-z_]+' --include='SKILL.md' . | sort -u | grep -vE '_\*$|_$' > /tmp/audit_used.txt
HALL=$(comm -23 /tmp/audit_used.txt /tmp/audit_real.txt | grep -v lark_api | wc -l | tr -d ' '); chk $([ "$HALL" = 0 ]&&echo 0||echo 1) L2.1 "no hallucinated lark_* tool in skills"
JQBAD=$(grep -rlE '[^a-z.]\.(tasks|users|cycles|records|results|messages|meetings|events)\[\]' --include='*.md' . | grep -v connectors | wc -l | tr -d ' '); chk $([ "$JQBAD" = 0 ]&&echo 0||echo 1) L2.2 "no wrong top-level jq paths in skills (need .data.)"
test -f connectors/LARK-PATTERNS.md && test -f connectors/LARK-RECIPES.md && test -f connectors/LARK-FUSION.md; chk $? L2.3 "depth-core docs present (pattern/recipes/fusion)"
STALE=$(grep -rn '21 tool\|21 MCP\|~21 curated\|liệt kê 21' . 2>/dev/null | grep -vE 'previews|QC-AUDIT.md|tools/audit.sh' | wc -l | tr -d ' '); chk $([ "$STALE" = 0 ]&&echo 0||echo 1) L2.4 "no stale '21 tools' (now 25) in docs"
grep -q 'search_fields is required' "$REPO/cmd/mcp/tools.go"; chk $? L2.5 "base_search Build enforces search_fields"

echo "######## L3 — SEMANTIC ########"
RENAMED=0
if [ -d /tmp/knowledge-work-plugins ]; then
  RENAMED=$(cd /tmp/knowledge-work-plugins && for f in $(find . -name SKILL.md); do o="/tmp/knowledge-work-plugins/$f"; n="$ROOT/$f"; [ -f "$n" ]||continue; a=$(grep -m1 '^name:' "$o"); b=$(grep -m1 '^name:' "$n"); [ "$a" != "$b" ]&&echo x; done | wc -l | tr -d ' ')
fi
chk $([ "$RENAMED" = 0 ]&&echo 0||echo 1) L3.1 "no SKILL renamed vs upstream ($([ -d /tmp/knowledge-work-plugins ]&&echo checked||echo 'upstream absent'))"
FMBAD=$(for f in $(find . -name SKILL.md); do h=$(head -1 "$f"); c=$(grep -c '^name:' "$f"); { [ "$h" != "---" ] || [ "$c" -ne 1 ]; } && echo "$f"; done | wc -l | tr -d ' '); chk $([ "$FMBAD" = 0 ]&&echo 0||echo 1) L3.2 "frontmatter intact (--- + exactly one name:)"
FP=$(for d in enterprise-search operations human-resources customer-support sales product-management marketing legal design data finance engineering bio-research small-business productivity; do m=$(grep -rl 'Lark-native\|LARK-PATTERNS\|system of record' --include=SKILL.md "$d" 2>/dev/null|wc -l|tr -d ' '); t=$(find "$d" -name SKILL.md|wc -l|tr -d ' '); [ "$m" != "$t" ]&&echo "$d($m/$t)"; done); chk $([ -z "$FP" ]&&echo 0||echo 1) L3.3 "all first-party SKILL.md soul-swapped (${FP:-100%})"
ART=$(grep -rniE "as requested|I (have |'ve )(edited|rewrote|rewritten)|here is the (updated|rewritten)" --include='SKILL.md' . | wc -l | tr -d ' '); chk $([ "$ART" = 0 ]&&echo 0||echo 1) L3.4 "no agent meta-commentary artifacts in skills"
PTR=$(for c in */CONNECTORS.md partner-built/*/CONNECTORS.md; do grep -ql 'Lark depth\|LARK-PATTERNS' "$c"||echo "$c"; done | grep -v pdf-viewer | wc -l | tr -d ' '); chk $([ "$PTR" = 0 ]&&echo 0||echo 1) L3.5 "every plugin CONNECTORS points to depth core"

echo "######## L4 — RUNTIME ########"
( cd "$REPO" && go vet ./cmd/mcp/ ./shortcuts/okr/ >/dev/null 2>&1 ); chk $? L4.1 "go vet (cmd/mcp + okr)"
( cd "$REPO" && go test ./cmd/mcp/ ./shortcuts/okr/ >/dev/null 2>&1 ); chk $? L4.2 "go test (cmd/mcp + okr)"
( cd "$REPO" && go build -o "$CLI" ./ >/dev/null 2>&1 ); chk $? L4.3 "go build -> \$CLI"
N=$("$CLI" mcp tools 2>/dev/null | grep -oE '[0-9]+ tools total' | grep -oE '[0-9]+'); chk $([ "$N" = 25 ]&&echo 0||echo 1) L4.4 "mcp tools (offline) = 25 (got ${N:-?})"
HS=$(printf '%s\n%s\n%s\n' '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{},"clientInfo":{"name":"a","version":"1"}}}' '{"jsonrpc":"2.0","method":"notifications/initialized"}' '{"jsonrpc":"2.0","id":2,"method":"tools/list","params":{}}' | "$CLI" mcp serve 2>/dev/null | python3 -c "import sys,json
n=0
for l in sys.stdin:
    l=l.strip()
    if l and '\"id\":2' in l:
        n=len(json.loads(l).get('result',{}).get('tools',[]))
print(n)")
chk $([ "$HS" = 25 ]&&echo 0||echo 1) L4.5 "MCP stdio handshake tools/list = 25 (got ${HS:-?})"
CARD=$("$CLI" im +card-send --print-json --spec 'header: { title: "t", template: blue }
elements:
  - md: "**x**"' 2>&1 | grep -c '"ok": true'); chk $([ "$CARD" -ge 1 ]&&echo 0||echo 1) L4.6 "interactive card (P4) compiles offline (ok:true)"

if [ "$LIVE" = 1 ]; then
  echo "######## L4-LIVE — needs auth ########"
  "$CLI" auth status >/dev/null 2>&1; chk $? L4L.0 "auth ready"
  ME=$("$CLI" mcp tools >/dev/null 2>&1; printf '%s\n%s\n%s\n' '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{},"clientInfo":{"name":"a"}}}' '{"jsonrpc":"2.0","method":"notifications/initialized"}' '{"jsonrpc":"2.0","id":2,"method":"tools/call","params":{"name":"lark_contact_search","arguments":{"user_ids":"me","jq":".data.users[]?|.open_id"}}}' | "$CLI" mcp serve 2>/dev/null | python3 -c "import sys,json
for l in sys.stdin:
    if l.strip() and '\"id\":2' in l:
        t=(json.loads(l).get('result',{}).get('content') or [{}])[0].get('text','').strip(); print(t)")
  chk $([ -n "$ME" ] && echo 0 || echo 1) L4L.1 "live read: lark_contact_search me -> open_id (${ME:-none})"
  # safe task round-trip (own throwaway artifact, deleted)
  python3 - "$CLI" <<'PY'; chk $? L4L.2 "live mutating round-trip: task create->complete->delete"
import subprocess,json,sys,os
CLI=sys.argv[1]; _id=[1]
def call(n,a):
    _id[0]+=1;i=_id[0]
    L=['{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{},"clientInfo":{"name":"a"}}}','{"jsonrpc":"2.0","method":"notifications/initialized"}',json.dumps({"jsonrpc":"2.0","id":i,"method":"tools/call","params":{"name":n,"arguments":a}})]
    p=subprocess.run([CLI,"mcp","serve"],input="\n".join(L)+"\n",capture_output=True,text=True,timeout=90)
    for l in p.stdout.splitlines():
        if l.strip() and f'"id":{i}' in l:
            try:return json.loads((json.loads(l).get("result",{}).get("content") or [{}])[0].get("text",""))
            except:return {}
    return {}
j=call("lark_task_create",{"summary":"[AUDIT] auto — delete"}); d=(j.get("data") or {}); g=d.get("guid") or d.get("task",{}).get("guid")
if not g: sys.exit(1)
if not call("lark_task_complete",{"task_id":g}).get("ok"): sys.exit(1)
call("lark_api",{"method":"DELETE","path":f"/open-apis/task/v2/tasks/{g}"})
gone=call("lark_api",{"method":"GET","path":f"/open-apis/task/v2/tasks/{g}"})
sys.exit(0 if (gone.get("ok") is False or "error_type" in gone or "error" in gone) else 1)
PY
fi

echo "========================================"
echo "RESULT: $PASS passed, $FAIL failed"
[ "$FAIL" = 0 ] && echo "STATUS: ✅ ALL GREEN" || echo "STATUS: ❌ $FAIL FAILURE(S)"
exit $FAIL
