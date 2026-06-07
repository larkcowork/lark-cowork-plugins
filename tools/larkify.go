// Command larkify converts Anthropic knowledge-work plugins to Lark Cowork
// and switches the lark connector transport. It replaces the earlier
// larkify.py + set-transport.sh so the tooling matches the Go codebase.
//
//	larkify convert <plugin> [<plugin> ...]   # rewrite .mcp.json + CONNECTORS.md + terminology
//	larkify transport <stdio|http>            # switch the `lark` block in every plugin .mcp.json
//
// Run from anywhere; the marketplace root is found by walking up to the
// directory that contains connectors/lark.stdio.json.
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

// genericServers are the generic-comms MCP servers the single `lark`
// bridge replaces. Matched case-insensitively by server name.
var genericServers = map[string]bool{
	"slack": true, "gmail": true, "google calendar": true, "google drive": true,
	"notion": true, "asana": true, "linear": true, "monday": true, "clickup": true,
	"atlassian": true, "guru": true, "box": true, "egnyte": true, "fireflies": true,
	"confluence": true, "coda": true, "microsoft 365": true, "outlook": true,
}

// larkRow is one generic category's mapping to Lark tools.
type larkRow struct{ product, tools string }

// larkMap maps a placeholder category (lowercased) to its Lark backing.
var larkMap = map[string]larkRow{
	"chat":                      {"Lark IM", "`lark_im_send`, `lark_im_search`, `lark_im_card_send`"},
	"email":                     {"Lark Mail", "`lark_mail_send`, `lark_mail_draft_create`"},
	"calendar":                  {"Lark Calendar", "`lark_calendar_agenda`, `lark_calendar_create`"},
	"knowledge base":            {"Lark Wiki + Docs", "`lark_doc_search`, `lark_doc_fetch`, `lark_doc_create`"},
	"documents":                 {"Lark Docs", "`lark_doc_create`, `lark_doc_fetch`"},
	"office suite":              {"Lark Docs/Sheets/Drive", "`lark_doc_*`, `lark_sheets_read`, `lark_sheets_append`, `lark_drive_upload`"},
	"spreadsheet":               {"Lark Sheets", "`lark_sheets_read`, `lark_sheets_append`"},
	"project tracker":           {"Lark Task + Base", "`lark_task_create`, `lark_task_my`, `lark_base_search`"},
	"jira":                      {"Lark Task + Base", "`lark_task_create`, `lark_task_my`, `lark_base_search`"},
	"database":                  {"Lark Base (Bitable)", "`lark_base_search`"},
	"cloud storage":             {"Lark Drive", "`lark_drive_upload`"},
	"conversation intelligence": {"Lark Minutes + VC", "`lark_minutes_search`, `lark_vc_search`"},
	"meeting transcription":     {"Lark Minutes", "`lark_minutes_search`"},
	"okr":                       {"Lark OKR", "`lark_okr_cycle_list`"},
	"directory":                 {"Lark Contact", "`lark_contact_search`"},
}

// termRepl is an ordered prose replacement (generic tools only).
type termRepl struct {
	re  *regexp.Regexp
	rep string
}

var terms = func() []termRepl {
	raw := [][2]string{
		{`Microsoft 365`, "Lark"},
		{`Google Calendar`, "Lark Calendar"},
		{`Google Drive`, "Lark Drive"},
		{`\bGmail\b`, "Lark Mail"},
		{`\bSlack\b`, "Lark IM"},
		{`\bNotion\b`, "Lark Wiki"},
		{`\bConfluence\b`, "Lark Wiki"},
		{`\bAtlassian\b`, "Lark"},
		{`\bAsana\b`, "Lark Task"},
		{`\bLinear\b`, "Lark Task"},
		{`\bClickUp\b`, "Lark Task"},
		{`monday\.com`, "Lark Base"},
		{`\bJira\b`, "Lark Task"},
		{`\bGuru\b`, "Lark Wiki"},
		{`\bFireflies\b`, "Lark Minutes"},
	}
	out := make([]termRepl, len(raw))
	for i, r := range raw {
		out[i] = termRepl{regexp.MustCompile(r[0]), r[1]}
	}
	return out
}()

func findRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "connectors", "lark.stdio.json")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("marketplace root not found (no connectors/lark.stdio.json above cwd)")
		}
		dir = parent
	}
}

func larkBlock(root, mode string) (json.RawMessage, error) {
	b, err := os.ReadFile(filepath.Join(root, "connectors", "lark."+mode+".json"))
	if err != nil {
		return nil, err
	}
	var doc struct {
		MCPServers map[string]json.RawMessage `json:"mcpServers"`
	}
	if err := json.Unmarshal(b, &doc); err != nil {
		return nil, err
	}
	blk, ok := doc.MCPServers["lark"]
	if !ok {
		return nil, fmt.Errorf("connectors/lark.%s.json has no `lark` server", mode)
	}
	return blk, nil
}

// writeJSON marshals indented with a trailing newline (matches gofmt-ish house style).
func writeJSON(path string, v any) error {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, append(b, '\n'), 0o644)
}

// orderedMCP keeps mcpServers stable: existing kept servers first (sorted), lark last.
func rewriteMCP(root, plugin string, lark json.RawMessage) (dropped, kept []string, err error) {
	p := filepath.Join(root, plugin, ".mcp.json")
	servers := map[string]json.RawMessage{}
	if b, e := os.ReadFile(p); e == nil {
		var doc struct {
			MCPServers map[string]json.RawMessage `json:"mcpServers"`
		}
		if e := json.Unmarshal(b, &doc); e != nil {
			return nil, nil, fmt.Errorf("%s: %w", p, e)
		}
		servers = doc.MCPServers
	}
	out := map[string]json.RawMessage{}
	for name, cfg := range servers {
		if genericServers[strings.ToLower(name)] {
			dropped = append(dropped, name)
		} else {
			out[name] = cfg
			kept = append(kept, name)
		}
	}
	out["lark"] = lark
	kept = append(kept, "lark")
	sort.Strings(dropped)
	sort.Strings(kept)
	return dropped, kept, writeJSON(p, map[string]any{"mcpServers": out})
}

var phRe = regexp.MustCompile(`~~([a-zA-Z ./-]+)`)

type catRow struct{ label, key, ph, included string }

func parseCategories(root, plugin string) []catRow {
	p := filepath.Join(root, plugin, "CONNECTORS.md")
	b, err := os.ReadFile(p)
	if err != nil {
		return nil
	}
	var rows []catRow
	for _, line := range strings.Split(string(b), "\n") {
		t := strings.TrimSpace(line)
		if !strings.HasPrefix(t, "|") {
			continue
		}
		cells := strings.Split(strings.Trim(t, "|"), "|")
		for i := range cells {
			cells[i] = strings.TrimSpace(cells[i])
		}
		var ph string
		isHeader := false
		for _, c := range cells {
			if strings.Contains(c, "Placeholder") {
				isHeader = true
			}
			if strings.Contains(c, "~~") {
				ph = c
			}
		}
		if ph == "" || isHeader {
			continue
		}
		m := phRe.FindStringSubmatch(ph)
		if m == nil {
			continue
		}
		included := ""
		if len(cells) > 2 {
			included = cells[2]
		}
		rows = append(rows, catRow{cells[0], strings.ToLower(strings.TrimSpace(m[1])), strings.Trim(ph, "` "), included})
	}
	return rows
}

func writeConnectors(root, plugin string) (larked, external []string) {
	rows := parseCategories(root, plugin)
	var b strings.Builder
	b.WriteString("# Connectors (Lark)\n\n## How tool references work\n\n")
	b.WriteString("Plugin files use `~~category` as a placeholder. In **Lark Cowork**, generic collaboration\n")
	b.WriteString("categories resolve to the single **`lark`** MCP server (`lark-cli mcp serve`), which exposes\n")
	b.WriteString("curated `lark_*` tools plus a generic `lark_api` escape hatch. Specialty categories keep\n")
	b.WriteString("their own external MCP server — connect those separately.\n\n")
	b.WriteString("## Connectors for this plugin\n\n")
	b.WriteString("| Category | Placeholder | Backed by | Tools / notes |\n")
	b.WriteString("|----------|-------------|-----------|---------------|\n")
	for _, r := range rows {
		name := strings.TrimPrefix(r.ph, "~~")
		if lr, ok := larkMap[r.key]; ok {
			fmt.Fprintf(&b, "| %s | `~~%s` | %s (`lark`) | %s |\n", r.label, name, lr.product, lr.tools)
			larked = append(larked, r.key)
		} else {
			keep := r.included
			if keep == "" || keep == "—" {
				keep = "external MCP"
			}
			fmt.Fprintf(&b, "| %s | `~~%s` | %s (external) | connect this tool's own MCP server; or approximate with Lark Base/Sheets |\n", r.label, name, keep)
			external = append(external, r.key)
		}
	}
	b.WriteString("\n## Notes for Claude\n\n")
	b.WriteString("- **Resolve people first** — use `lark_contact_search` to turn a name/email into an `open_id` before any send/invite/assign.\n")
	b.WriteString("- **Escape hatch** — anything without a dedicated tool goes through `lark_api` (Lark OpenAPI path + params).\n")
	b.WriteString("- **Identity** — tools run as the authenticated user by default; retry with bot identity if an error envelope says so.\n\n")
	b.WriteString(depthPointer)
	_ = os.WriteFile(filepath.Join(root, plugin, "CONNECTORS.md"), []byte(b.String()), 0o644)
	return larked, external
}

const depthPointer = "## Lark depth — read these for any non-trivial workflow\n\n" +
	"This plugin runs on Lark. Before executing, consult the shared depth core (one level up):\n\n" +
	"- [`../connectors/LARK-PATTERNS.md`](../connectors/LARK-PATTERNS.md) — Lark-native workflow patterns (interactive cards, Base as system-of-record, Minutes→tasks, approval flows, safe-mutation, token economy).\n" +
	"- [`../connectors/LARK-RECIPES.md`](../connectors/LARK-RECIPES.md) — exact `lark_api` calls for ops the curated tools don't cover.\n" +
	"- [`../connectors/LARK-FUSION.md`](../connectors/LARK-FUSION.md) — when to delegate to a deeper installed `lark-*` / `base-deploy` skill instead of hand-rolling.\n\n" +
	"Default posture: resolve people first, project reads with `jq`, preview mutations with `dry_run`, and surface decisions/digests as interactive cards rather than plain text.\n"

func terminology(root, plugin string) int {
	patched := 0
	_ = filepath.Walk(filepath.Join(root, plugin), func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		ext := filepath.Ext(path)
		if ext != ".md" && ext != ".json" && ext != ".txt" && ext != ".html" {
			return nil
		}
		base := filepath.Base(path)
		if base == "CONNECTORS.md" || base == ".mcp.json" {
			return nil
		}
		b, e := os.ReadFile(path)
		if e != nil {
			return nil
		}
		s := string(b)
		o := s
		for _, t := range terms {
			s = t.re.ReplaceAllString(s, t.rep)
		}
		if s != o {
			_ = os.WriteFile(path, []byte(s), info.Mode())
			patched++
		}
		return nil
	})
	return patched
}

func cmdConvert(root string, plugins []string) error {
	lark, err := larkBlock(root, "stdio")
	if err != nil {
		return err
	}
	for _, plugin := range plugins {
		dropped, kept, err := rewriteMCP(root, plugin, lark)
		if err != nil {
			return err
		}
		larked, external := writeConnectors(root, plugin)
		patched := terminology(root, plugin)
		fmt.Printf("\n### %s\n", plugin)
		fmt.Printf("  .mcp.json kept: %v\n  dropped(generic): %v\n", kept, dropped)
		fmt.Printf("  CONNECTORS lark: %v\n  external(keep): %v\n", larked, external)
		fmt.Printf("  terminology files patched: %d\n", patched)
	}
	return nil
}

func cmdTransport(root, mode string) error {
	if mode != "stdio" && mode != "http" {
		return fmt.Errorf("transport must be stdio or http")
	}
	lark, err := larkBlock(root, mode)
	if err != nil {
		return err
	}
	entries, _ := filepath.Glob(filepath.Join(root, "*", ".mcp.json"))
	n := 0
	for _, p := range entries {
		b, e := os.ReadFile(p)
		if e != nil {
			continue
		}
		var doc struct {
			MCPServers map[string]json.RawMessage `json:"mcpServers"`
		}
		if json.Unmarshal(b, &doc) != nil || doc.MCPServers == nil {
			continue
		}
		if _, ok := doc.MCPServers["lark"]; !ok {
			continue
		}
		doc.MCPServers["lark"] = lark
		if writeJSON(p, map[string]any{"mcpServers": doc.MCPServers}) == nil {
			n++
		}
	}
	fmt.Printf("applied '%s' transport to %d plugin .mcp.json files with a `lark` connector\n", mode, n)
	return nil
}

func usage() {
	fmt.Fprintln(os.Stderr, "usage:\n  larkify convert <plugin> [<plugin> ...]\n  larkify transport <stdio|http>")
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}
	root, err := findRoot()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	switch os.Args[1] {
	case "convert":
		if len(os.Args) < 3 {
			usage()
			os.Exit(2)
		}
		err = cmdConvert(root, os.Args[2:])
	case "transport":
		if len(os.Args) != 3 {
			usage()
			os.Exit(2)
		}
		err = cmdTransport(root, os.Args[2])
	default:
		usage()
		os.Exit(2)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
