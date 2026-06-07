package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// setup builds a throwaway marketplace root: connectors/ + one plugin.
func setup(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	must(t, os.MkdirAll(filepath.Join(root, "connectors"), 0o755))
	must(t, os.WriteFile(filepath.Join(root, "connectors", "lark.stdio.json"),
		[]byte(`{"mcpServers":{"lark":{"command":"lark-cli","args":["mcp","serve"],"env":{"NO_COLOR":"1"}}}}`), 0o644))
	must(t, os.WriteFile(filepath.Join(root, "connectors", "lark.http.json"),
		[]byte(`{"mcpServers":{"lark":{"type":"http","url":"${LARK_MCP_URL}"}}}`), 0o644))

	p := filepath.Join(root, "demo")
	must(t, os.MkdirAll(filepath.Join(p, "skills", "x"), 0o755))
	must(t, os.WriteFile(filepath.Join(p, ".mcp.json"),
		[]byte(`{"mcpServers":{"slack":{"type":"http","url":"s"},"hubspot":{"type":"http","url":"h"}}}`), 0o644))
	must(t, os.WriteFile(filepath.Join(p, "CONNECTORS.md"),
		[]byte("| Category | Placeholder | Included servers | Other |\n|--|--|--|--|\n| Chat | `~~chat` | Slack | — |\n| CRM | `~~CRM` | HubSpot | — |\n"), 0o644))
	must(t, os.WriteFile(filepath.Join(p, "skills", "x", "SKILL.md"),
		[]byte("Send a note in Slack and file it in Notion. Keep HubSpot untouched.\n"), 0o644))
	return root
}

func must(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

func TestRewriteMCP(t *testing.T) {
	root := setup(t)
	lark, err := larkBlock(root, "stdio")
	must(t, err)
	dropped, kept, err := rewriteMCP(root, "demo", lark)
	must(t, err)
	if got := strings.Join(dropped, ","); got != "slack" {
		t.Fatalf("dropped = %q, want slack", got)
	}
	if got := strings.Join(kept, ","); got != "hubspot,lark" {
		t.Fatalf("kept = %q, want hubspot,lark", got)
	}
	b, _ := os.ReadFile(filepath.Join(root, "demo", ".mcp.json"))
	var doc struct {
		MCPServers map[string]json.RawMessage `json:"mcpServers"`
	}
	must(t, json.Unmarshal(b, &doc))
	if _, ok := doc.MCPServers["lark"]; !ok {
		t.Fatal("lark server missing after rewrite")
	}
	if _, ok := doc.MCPServers["slack"]; ok {
		t.Fatal("slack server should have been dropped")
	}
}

func TestWriteConnectors(t *testing.T) {
	root := setup(t)
	larked, external := writeConnectors(root, "demo")
	if strings.Join(larked, ",") != "chat" {
		t.Fatalf("larked = %v, want [chat]", larked)
	}
	if strings.Join(external, ",") != "crm" {
		t.Fatalf("external = %v, want [crm]", external)
	}
	b, _ := os.ReadFile(filepath.Join(root, "demo", "CONNECTORS.md"))
	s := string(b)
	if !strings.Contains(s, "Lark IM (`lark`)") {
		t.Fatal("chat row not mapped to Lark IM")
	}
	if !strings.Contains(s, "HubSpot (external)") {
		t.Fatal("CRM row should stay external HubSpot")
	}
	if !strings.Contains(s, "Lark depth — read these") {
		t.Fatal("depth pointer not appended")
	}
}

func TestTerminology(t *testing.T) {
	root := setup(t)
	n := terminology(root, "demo")
	if n != 1 {
		t.Fatalf("patched %d files, want 1", n)
	}
	b, _ := os.ReadFile(filepath.Join(root, "demo", "skills", "x", "SKILL.md"))
	s := string(b)
	if strings.Contains(s, "Slack") || strings.Contains(s, "Notion") {
		t.Fatalf("generic names not replaced: %q", s)
	}
	if !strings.Contains(s, "Lark IM") || !strings.Contains(s, "Lark Wiki") {
		t.Fatalf("missing Lark replacements: %q", s)
	}
	if !strings.Contains(s, "HubSpot") {
		t.Fatal("specialty name HubSpot must be preserved")
	}
}
