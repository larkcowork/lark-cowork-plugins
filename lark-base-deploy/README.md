# Plugin Triển khai Lark Base (Lark Base Deploy)

Plugin triển khai Lark Base được thiết kế chủ yếu cho [Cowork](https://claude.com/product/cowork), ứng dụng desktop dạng agent của Anthropic — dù vậy nó cũng hoạt động tốt trong Claude Code. Plugin biến một câu nói duy nhất — "dựng một Base cho team 30 người của tôi: 3 table + dashboard + tự động hoá + import Excel" — thành một Base hoàn chỉnh, đã được kiểm chứng và bàn giao. Một orchestrator 8 phase chạy từ Discovery → Handover, thiết kế schema, dựng nó trên Lark, nối các quan hệ và vai trò, import dữ liệu, dựng dashboard và tự động hoá, rồi nghiệm thu và viết tài liệu kết quả cho team.

## Cài đặt

```
claude plugins add lark-cowork/lark-base-deploy
```

## Tính năng

Plugin này điều phối toàn bộ quá trình triển khai Base từ đầu đến cuối qua tám phase:

- **Phase 0 — Discovery** — phỏng vấn/chốt yêu cầu vận hành thành một spec (entity, KPI, vai trò, nguồn dữ liệu).
- **Phase 1 — Design** — thiết kế toàn bộ schema: table, field, view, quan hệ, rule, option, vai trò.
- **Phase 2 — Build** ⚡ — tạo Base, table, field và view; **fan-out 6 sub-agent** (table, field, view, quan hệ, rule, option) chạy song song; ghi ID thật trở lại vào plan.
- **Phase 3 — Wire-up** — nối các table (link/lookup/formula), đặt bộ lọc view, dựng vai trò và quyền truy cập.
- **Phase 4 — Import** — import dữ liệu Excel/CSV vào Base, làm sạch và map field theo schema đã thiết kế.
- **Phase 5 — Viz** ⚡ — dựng dashboard KPI với biểu đồ và bộ lọc; **fan-out 3 sub-agent** để soạn `data_config` song song, rồi tạo các block tuần tự.
- **Phase 6 — Automation** — thiết lập tự động hoá Base workflow cùng một bot để nhắc việc và thông báo.
- **Phase 7 — Handover** — nghiệm thu UI thật (không bao giờ tin vào `code:0`), viết hướng dẫn sử dụng, và bàn giao Base cho team.

Mọi thứ đều xoay quanh **`build-plan.json` — nguồn sự thật duy nhất (SSOT)**. Mỗi phase đều đọc và ghi vào nó, và chỉ những định danh thật do API trả về (`base_token`, `table_id`, …) mới được ghi ngược trở lại.

## Kỹ năng

| Kỹ năng | Mô tả |
|-------|-------------|
| `base-deploy` | Orchestrator — triển khai một Lark Base end-to-end cho team qua 8 phase (Discovery → Handover), spawn các sub-agent song song ở Build và Viz. |
| `base-discovery` | Phase 0 — phỏng vấn/chốt yêu cầu thành một spec (entity, KPI, vai trò, nguồn dữ liệu) ghi vào `build-plan.json`. |
| `base-design` | Phase 1 — thiết kế schema (table, field, view, quan hệ, rule, option, vai trò) thành một `build-plan.json` hoàn chỉnh để Build thực thi. |
| `base-build` | Phase 2 — tạo Base + table + field + view từ `build-plan.json`, fan-out 1 sub-agent cho mỗi table song song, ghi ID thật trở lại vào plan. |
| `base-wireup` | Phase 3 — nối các table (link/lookup/formula), đặt bộ lọc view, dựng vai trò và quyền truy cập từ `build-plan.json`. |
| `base-import` | Phase 4 — import Excel/CSV vào Base, làm sạch và map field chuẩn hoá theo `build-plan.json`. |
| `base-viz` | Phase 5 — dựng dashboard KPI (biểu đồ, bộ lọc), fan-out 3 sub-agent soạn `data_config` song song, rồi tạo các block tuần tự. |
| `base-automation` | Phase 6 — dựng tự động hoá Base workflow cùng một bot nhắc việc/thông báo theo `build-plan.json`. |
| `base-handover` | Phase 7 — nghiệm thu UI (verify thật, không tin vào `code:0`), viết hướng dẫn sử dụng, và bàn giao Base cho team. |

## Quy trình mẫu

### Triển khai trọn vẹn từ một câu

```
You: Build a Base for my 30-person team: 3 tables (Projects, Tasks, People),
     a KPI dashboard, automation to remind owners of due tasks, and import
     our existing roster from team.xlsx

Claude: [Phase 0 Discovery — confirms entities, KPIs, owner, data source → build-plan.json]
        [Phase 1 Design — drafts tables/fields/views/relations/roles into the plan]
        [Phase 2 Build ⚡ — creates Base, fans out 6 sub-agents, fills real IDs]
        [Phase 3 Wire-up — links Tasks→Projects, lookups, view filters, roles]
        [Phase 4 Import — maps team.xlsx into People, cleans + normalizes]
        [Phase 5 Viz ⚡ — 3 sub-agents compose data_config, builds dashboard blocks]
        [Phase 6 Automation — due-task reminder workflow + bot notifications]
        [Phase 7 Handover — visually verifies the UI, writes usage guide, hands over]
```

### Chỉ dựng Dashboard trên một Base có sẵn

```
You: Add a KPI dashboard to my existing Base (base_token bascn...) —
     deal count by stage, win-rate trend, and a stage filter

Claude: [Loads/derives build-plan.json from the existing Base]
        [Runs Phase 5 Viz ⚡ — 3 sub-agents compose data_config in parallel]
        [Creates dashboard blocks sequentially, verifies them in the UI]
```

### Chạy thử (Dry-Run) một plan trước khi dựng

```
You: /base-deploy --dry-run a sprint tracker: Sprints, Stories, Bugs

Claude: [Runs Discovery + Design only — produces a complete build-plan.json]
        [Shows the proposed tables, fields, relations, roles, dashboard, automation]
        [Stops before any write — "Approve to run Build → Handover?"]
```

## Plugin đồng hành

Các kỹ năng này thuộc bộ quy trình **lark-cowork** và tham chiếu đến kỹ năng ở các plugin
cùng nhóm. Tên kỹ năng được phân giải toàn cục, nên một tham chiếu sẽ tự động hoạt động khi plugin
đồng hành được cài; khi vắng plugin đồng hành thì tham chiếu sẽ **suy giảm mượt mà** (bước đó
bị bỏ qua hoặc được đề xuất như một gợi ý, không bao giờ là lỗi). Hãy cài các plugin đồng hành dưới đây để có
trải nghiệm đầy đủ — xem [`../connectors/LARK-FUSION.md`](../connectors/LARK-FUSION.md).

| Kỹ năng của plugin này | Tham chiếu | Thuộc plugin |
|---|---|---|
| `base-handover` | `permission-audit` | governance |

## Nguồn dữ liệu

> Nếu bạn thấy các placeholder lạ hoặc cần kiểm tra công cụ nào đang được kết nối, hãy xem [CONNECTORS.md](CONNECTORS.md).

**Kết nối MCP đi kèm:** server `lark` (`lark-cli mcp serve`), một cầu nối duy nhất phục vụ mọi danh mục —
- Project tracker (Lark Base + Task) — hệ thống lưu trữ chính thức mà plugin này dựng và ghi vào
- Bộ ứng dụng văn phòng (Lark Sheets/Docs/Drive) — để import Excel/CSV và tạo tài liệu bàn giao
- Chat (Lark IM) — để gửi nhắc việc tự động và thông báo của bot
- Danh bạ (Lark Contact) — để phân giải con người thành `open_id` trước khi gán field/vai trò

Cấu trúc Base và các thao tác ghi record (tạo base/table/field/view, bitable records, dashboard blocks, workflow automation) không có công cụ chuyên dụng được tuyển chọn — chúng đi qua `lark_api` (`bitable/v1`, `base/v2`) hoặc kỹ năng `lark-base` đã cài. Xem [CONNECTORS.md](CONNECTORS.md) để biết bản đồ ánh xạ đầy đủ giữa danh mục và công cụ.

---

## Tác giả

**Nguyễn Ngọc Tuấn**
Founder Transform Group — **Lark Platinum Partner**
🌐 Dự án: [larkcowork.com](https://larkcowork.com)
