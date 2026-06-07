---
name: instrument-data-to-allotrope
description: Convert laboratory instrument output files (PDF, CSV, Excel, TXT) to Allotrope Simple Model (ASM) JSON format or flattened 2D CSV. Use this skill when scientists need to standardize instrument data for LIMS systems, data lakes, or downstream analysis. Supports auto-detection of instrument types. Outputs include full ASM JSON, flattened CSV for easy import, and exportable Python code for data engineers. Common triggers include converting instrument files, standardizing lab data, preparing data for upload to LIMS/ELN systems, or generating parser code for production pipelines.
---

# Instrument Data to Allotrope Converter

Convert instrument files into standardized Allotrope Simple Model (ASM) format for LIMS upload, data lakes, or handoff to data engineering teams.

> **Note: This is an Example Skill**
>
> This skill demonstrates how skills can support your data engineering tasks—automating schema transformations, parsing instrument outputs, and generating production-ready code.
>
> **To customize for your organization:**
> - Modify the `references/` files to include your company's specific schemas or ontology mappings
> - Use an MCP server to connect to systems that define your schemas (e.g., your LIMS, data catalog, or schema registry)
> - Extend the `scripts/` to handle proprietary instrument formats or internal data standards
>
> This pattern can be adapted for any data transformation workflow where you need to convert between formats or validate against organizational standards.

> **Lark-native execution** (depth core: [LARK-PATTERNS](../../../connectors/LARK-PATTERNS.md),
> [LARK-RECIPES](../../../connectors/LARK-RECIPES.md), [LARK-FUSION](../../../connectors/LARK-FUSION.md)).
> Parsing/validation stays **local** (allotropy + the scripts below) — Lark doesn't transform the
> data. The Lark-native part is **delivery and handoff**: land the ASM JSON / flattened CSV / parser
> code in **Drive** (P8); log the conversion to a **Base** as the lab's data-handoff system-of-record
> (P5); and notify the requester / data engineer with an **interactive card** (P4). If your real LIMS
> is itself a Lark Base, the flattened CSV maps straight onto its table via `lark_base_record_upsert`.
> See "Step 5: Deliver to Lark" below.

## Workflow Overview

1. **Detect instrument type** from file contents (auto-detect or user-specified)
2. **Parse file** using allotropy library (native) or flexible fallback parser
3. **Generate outputs**:
   - ASM JSON (full semantic structure)
   - Flattened CSV (2D tabular format)
   - Python parser code (for data engineer handoff)
4. **Validate** the ASM (`scripts/validate_asm.py`) before delivery
5. **Deliver to Lark** — upload to Drive, log to Base, notify with a card (see below)

> **When Uncertain:** If you're unsure how to map a field to ASM (e.g., is this raw data or calculated? device setting or environmental condition?), ask the user for clarification. Refer to `references/field_classification_guide.md` for guidance, but when ambiguity remains, confirm with the user rather than guessing.

## Quick Start

```python
# Install requirements first
pip install allotropy pandas openpyxl pdfplumber --break-system-packages

# Core conversion
from allotropy.parser_factory import Vendor
from allotropy.to_allotrope import allotrope_from_file

# Convert with allotropy
asm = allotrope_from_file("instrument_data.csv", Vendor.BECKMAN_VI_CELL_BLU)
```

## Output Format Selection

**ASM JSON (default)** - Full semantic structure with ontology URIs
- Best for: LIMS systems expecting ASM, data lakes, long-term archival
- Validates against Allotrope schemas

**Flattened CSV** - 2D tabular representation
- Best for: Quick analysis, Excel users, systems without JSON support
- Each measurement becomes one row with metadata repeated

**Both** - Generate both formats for maximum flexibility

## Calculated Data Handling

**IMPORTANT:** Separate raw measurements from calculated/derived values.

- **Raw data** → `measurement-document` (direct instrument readings)
- **Calculated data** → `calculated-data-aggregate-document` (derived values)

Calculated values MUST include traceability via `data-source-aggregate-document`:

```json
"calculated-data-aggregate-document": {
  "calculated-data-document": [{
    "calculated-data-identifier": "SAMPLE_B1_DIN_001",
    "calculated-data-name": "DNA integrity number",
    "calculated-result": {"value": 9.5, "unit": "(unitless)"},
    "data-source-aggregate-document": {
      "data-source-document": [{
        "data-source-identifier": "SAMPLE_B1_MEASUREMENT",
        "data-source-feature": "electrophoresis trace"
      }]
    }
  }]
}
```

**Common calculated fields by instrument type:**
| Instrument | Calculated Fields |
|------------|-------------------|
| Cell counter | Viability %, cell density dilution-adjusted values |
| Spectrophotometer | Concentration (from absorbance), 260/280 ratio |
| Plate reader | Concentrations from standard curve, %CV |
| Electrophoresis | DIN/RIN, region concentrations, average sizes |
| qPCR | Relative quantities, fold change |

See `references/field_classification_guide.md` for detailed guidance on raw vs. calculated classification.

## Validation

Always validate ASM output before delivering to the user:

```bash
python scripts/validate_asm.py output.json
python scripts/validate_asm.py output.json --reference known_good.json  # Compare to reference
python scripts/validate_asm.py output.json --strict  # Treat warnings as errors
```

**Validation Rules:**
- Based on Allotrope ASM specification (December 2024)
- Last updated: 2026-01-07
- Source: https://gitlab.com/allotrope-public/asm

**Soft Validation Approach:**
Unknown techniques, units, or sample roles generate **warnings** (not errors) to allow for forward compatibility. If Allotrope adds new values after December 2024, the validator won't block them—it will flag them for manual verification. Use `--strict` mode to treat warnings as errors if you need stricter validation.

**What it checks:**
- Correct technique selection (e.g., multi-analyte profiling vs plate reader)
- Field naming conventions (space-separated, not hyphenated)
- Calculated data has traceability (`data-source-aggregate-document`)
- Unique identifiers exist for measurements and calculated values
- Required metadata present
- Valid units and sample roles (with soft validation for unknown values)

## Step 5: Deliver to Lark

Once outputs are generated and validated locally, hand them off through the `lark` MCP
(apply P1/P2/P4/P5/P8):

1. **Upload artifacts to Drive** (P8) — `lark_drive_upload` for the ASM JSON, the flattened CSV, and
   the exported parser `.py`. Keep the returned links for the card and the Base row.
2. **Log the conversion to a Base** (P5, data-handoff system-of-record) — `lark_base_record_upsert`
   into a *Conversions* table: `{instrument, vendor, detection_confidence, input_file, n_measurements,
   n_calculated, asm_link:<Drive>, csv_link:<Drive>, validation:"pass|warnings|fail", owner:<open_id>,
   datetime}`. `dry_run: true` first (P2). Read prior conversions with `lark_base_search` — it does NOT
   support `jq` and REQUIRES `search_fields` (which field(s) to match); narrow with
   `select_fields`/`limit` instead, and discover field names via `lark_api GET
   /open-apis/bitable/v1/apps/{base}/tables/{table}/fields` if unknown (P3).
   No Base yet → delegate to `base-deploy`; record/field ops → delegate to `lark-base`.
3. **If the LIMS *is* a Lark Base** — push the flattened CSV rows straight into the LIMS table with
   `lark_base_record_upsert` (one upsert per row, keyed on `sample_identifier` + `well_position`),
   always `dry_run: true` first. This makes the "CSV I can upload to our LIMS" example a one-step
   native import. For batch/field-mapping details delegate to `lark-base`.
4. **Notify with a card** (P4) — resolve the data engineer / requester via `lark_contact_search`
   (P1, or `user_ids:"me"`), then `lark_im_card_send`: `header` ("Vi-CELL BLU → ASM: 24 measurements,
   validation passed"), `div` rows for instrument + counts + validation status, an `actions` button to
   the Drive ASM file. `print_json: true` → `dry_run: true` → send. Card grammar → delegate to
   `lark-im`. (For a plain handoff ping, `lark_im_send`.)

## Supported Instruments

See `references/supported_instruments.md` for complete list. Key instruments:

| Category | Instruments |
|----------|-------------|
| Cell Counting | Vi-CELL BLU, Vi-CELL XR, NucleoCounter |
| Spectrophotometry | NanoDrop One/Eight/8000, Lunatic |
| Plate Readers | SoftMax Pro, EnVision, Gen5, CLARIOstar |
| ELISA | SoftMax Pro, BMG MARS, MSD Workbench |
| qPCR | QuantStudio, Bio-Rad CFX |
| Chromatography | Empower, Chromeleon |

## Detection & Parsing Strategy

### Tier 1: Native allotropy parsing (PREFERRED)
**Always try allotropy first.** Check available vendors directly:

```python
from allotropy.parser_factory import Vendor

# List all supported vendors
for v in Vendor:
    print(f"{v.name}")

# Common vendors:
# AGILENT_TAPESTATION_ANALYSIS  (for TapeStation XML)
# BECKMAN_VI_CELL_BLU
# THERMO_FISHER_NANODROP_EIGHT
# MOLDEV_SOFTMAX_PRO
# APPBIO_QUANTSTUDIO
# ... many more
```

**When the user provides a file, check if allotropy supports it before falling back to manual parsing.** The `scripts/convert_to_asm.py` auto-detection only covers a subset of allotropy vendors.

### Tier 2: Flexible fallback parsing
**Only use if allotropy doesn't support the instrument.** This fallback:
- Does NOT generate `calculated-data-aggregate-document`
- Does NOT include full traceability
- Produces simplified ASM structure

Use flexible parser with:
- Column name fuzzy matching
- Unit extraction from headers
- Metadata extraction from file structure

### Tier 3: PDF extraction
For PDF-only files, extract tables using pdfplumber, then apply Tier 2 parsing.

## Pre-Parsing Checklist

Before writing a custom parser, ALWAYS:

1. **Check if allotropy supports it** - Use native parser if available
2. **Find a reference ASM file** - Check `references/examples/` or ask user
3. **Review instrument-specific guide** - Check `references/instrument_guides/`
4. **Validate against reference** - Run `validate_asm.py --reference <file>`

## Common Mistakes to Avoid

| Mistake | Correct Approach |
|---------|------------------|
| Manifest as object | Use URL string |
| Lowercase detection types | Use "Absorbance" not "absorbance" |
| "emission wavelength setting" | Use "detector wavelength setting" for emission |
| All measurements in one document | Group by well/sample location |
| Missing procedure metadata | Extract ALL device settings per measurement |

## Code Export for Data Engineers

Generate standalone Python scripts that scientists can hand off:

```python
# Export parser code
python scripts/export_parser.py --input "data.csv" --vendor "VI_CELL_BLU" --output "parser_script.py"
```

The exported script:
- Has no external dependencies beyond pandas/allotropy
- Includes inline documentation
- Can run in Jupyter notebooks
- Is production-ready for data pipelines

## File Structure

```
instrument-data-to-allotrope/
├── SKILL.md                          # This file
├── scripts/
│   ├── convert_to_asm.py            # Main conversion script
│   ├── flatten_asm.py               # ASM → 2D CSV conversion
│   ├── export_parser.py             # Generate standalone parser code
│   └── validate_asm.py              # Validate ASM output quality
└── references/
    ├── supported_instruments.md     # Full instrument list with Vendor enums
    ├── asm_schema_overview.md       # ASM structure reference
    ├── field_classification_guide.md # Where to put different field types
    └── flattening_guide.md          # How flattening works
```

## Usage Examples

### Example 1: Vi-CELL BLU file
```
User: "Convert this cell counting data to Allotrope format"
[uploads viCell_Results.xlsx]

Claude:
1. Detects Vi-CELL BLU (95% confidence)
2. Converts using allotropy native parser
3. Outputs:
   - viCell_Results_asm.json (full ASM)
   - viCell_Results_flat.csv (2D format)
   - viCell_parser.py (exportable code)
```

### Example 2: Request for code handoff
```
User: "I need to give our data engineer code to parse NanoDrop files"

Claude:
1. Generates self-contained Python script
2. Includes sample input/output
3. Documents all assumptions
4. Provides Jupyter notebook version
```

### Example 3: LIMS-ready flattened output
```
User: "Convert this ELISA data to a CSV I can upload to our LIMS"

Claude:
1. Parses plate reader data
2. Generates flattened CSV with columns:
   - sample_identifier, well_position, measurement_value, measurement_unit
   - instrument_serial_number, analysis_datetime, assay_type
3. Validates against common LIMS import requirements
```

## Implementation Notes

### Installing allotropy
```bash
pip install allotropy --break-system-packages
```

### Handling parse failures
If allotropy native parsing fails:
1. Log the error for debugging
2. Fall back to flexible parser
3. Report reduced metadata completeness to user
4. Suggest exporting different format from instrument

### ASM Schema Validation
Validate output against Allotrope schemas when available:
```python
import jsonschema
# Schema URLs in references/asm_schema_overview.md
```
