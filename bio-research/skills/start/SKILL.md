---
name: start
description: Set up your bio-research environment and explore available tools. Use when first getting oriented with the plugin, checking which literature, drug-discovery, or visualization MCP servers are connected, or surveying available analysis skills before starting a new project.
---

# Bio-Research Start

> If you see unfamiliar placeholders or need to check which tools are connected, see [CONNECTORS.md](../../CONNECTORS.md).

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md),
> [LARK-RECIPES](../../../connectors/LARK-RECIPES.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md)).
> The science tools (literature, drug-discovery, sequencing, viz) stay on their own external MCP
> servers — only the **collaboration layer runs on the single `lark` MCP**. In this plugin that means:
> deliverables land in **Wiki/Drive** (P8), every analysis run is logged to a **Lark Base** as the
> lab's system-of-record (P5), and orientation/results are surfaced as **interactive cards** (P4),
> not walls of text. Resolve any teammate by name with `lark_contact_search` first (P1).

You are helping a biological researcher get oriented with the bio-research plugin. Walk through the following steps in order.

## Step 1: Welcome

Display this welcome message:

```
Bio-Research Plugin

Your AI-powered research assistant for the life sciences. This plugin brings
together literature search, data analysis pipelines,
and scientific strategy — all in one place.
```

## Step 2: Check Available MCP Servers

Two server groups are in play:

- **`lark` MCP (always present)** — the collaboration backbone. Curated `lark_*` tools cover IM,
  cards, mail, calendar, docs, Wiki, Drive, tasks, Base, Minutes/VC, OKR, plus the `lark_api` escape
  hatch. This is where lab logs, run records, reports, and team notifications live.
- **Specialty science MCP servers (external, optional)** — literature, drug-discovery, viz, lab
  platforms. These keep their own servers; the `lark` server does not replace them.

Test which servers are connected by listing available tools. Group the results:

**Literature & Data Sources:**
- ~~literature database — biomedical literature search
- ~~literature database — preprint access (biology and medicine)
- ~~journal access — academic publications
- ~~data repository — collaborative research data (Sage Bionetworks)

**Drug Discovery & Clinical:**
- ~~chemical database — bioactive compound database
- ~~drug target database — drug target discovery platform
- ClinicalTrials.gov — clinical trial registry
- ~~clinical data platform — clinical trial site ranking and platform help

**Visualization & AI:**
- ~~scientific illustration — create scientific figures and diagrams
- ~~AI research platform — AI for biology (histopathology, drug discovery)

Report which servers are connected and which are not yet set up.

## Step 3: Survey Available Skills

List the analysis skills available in this plugin:

| Skill | What It Does |
|-------|-------------|
| **Single-Cell RNA QC** | Quality control for scRNA-seq data with MAD-based filtering |
| **scvi-tools** | Deep learning for single-cell omics (scVI, scANVI, totalVI, PeakVI, etc.) |
| **Nextflow Pipelines** | Run nf-core pipelines (RNA-seq, WGS/WES, ATAC-seq) |
| **Instrument Data Converter** | Convert lab instrument output to Allotrope ASM format |
| **Scientific Problem Selection** | Systematic framework for choosing research problems |

## Step 4: Optional Setup — Binary MCP Servers

Mention that two additional MCP servers are available as separate installations:

- **~~genomics platform** — Access cloud analysis data and workflows
  Install: Download `txg-node.mcpb` from https://github.com/10XGenomics/txg-mcp/releases
- **~~tool database** (Harvard MIMS) — AI tools for scientific discovery
  Install: Download `tooluniverse.mcpb` from https://github.com/mims-harvard/ToolUniverse/releases

These require downloading binary files and are optional.

## Step 5: Offer a Lab Base (system-of-record)

Bio-research generates a stream of artifacts (literature hits, QC runs, pipeline outputs, converted
datasets, project decisions). Offer to anchor them in a **Lark Base** so the lab has one queryable
source of truth (P5):

- Check first: `lark_base_search` for an existing lab Base. Note: `lark_base_search` does NOT support
  `jq` and REQUIRES `search_fields` (which field(s) to match per the Bitable API) — keep it cheap by
  narrowing with `select_fields`/`limit` instead. If you don't know the field names, discover them via
  `lark_api GET /open-apis/bitable/v1/apps/{base}/tables/{table}/fields` first (P3).
- If none exists, **delegate to the `base-deploy` skill** to scaffold one — don't hand-roll schema.
  A typical lab Base has tables like *Runs/Analyses* (skill, input, params, status, output link,
  owner), *Datasets*, *Literature*, and *Decisions*.
- Each analysis skill below then writes its run record via `lark_base_record_upsert` and lands its
  report in **Wiki/Drive** (P8), so this Step is the wiring all of them share.

## Step 6: Ask How to Help

Ask the researcher what they're working on today. Suggest starting points based on common workflows
(resolve any named collaborator with `lark_contact_search` before assigning/notifying, P1):

1. **Literature review** — "Search ~~literature database for recent papers on [topic]" → optionally
   log hits to the lab Base and post a digest **card** (P4).
2. **Analyze sequencing data** — "Run QC on my single-cell data" or "Set up an RNA-seq pipeline" →
   results to Drive, run logged to Base, completion card to chat.
3. **Drug discovery** — "Search ~~chemical database for compounds targeting [protein]" or "Find drug
   targets for [disease]".
4. **Data standardization** — "Convert my instrument data to Allotrope format" → ASM/CSV to Drive.
5. **Research strategy** — "Help me evaluate a new project idea" → strategy doc to Wiki.

Wait for the user's response and guide them to the appropriate tools and skills.
