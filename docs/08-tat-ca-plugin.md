# 8. Giới thiệu chi tiết tất cả plugin

> Trang này mô tả **từng plugin một**: nó làm gì, có những kỹ năng/lệnh nào, **khi nào dùng**, và
> **câu lệnh tiếng Việt mẫu** để bạn copy dùng ngay. Mỗi plugin là một "bộ kỹ năng theo nghề".

**Quy ước:** 💬 = câu bạn gõ cho trợ lý · `kỹ-năng` = tên kỹ năng (Claude tự gọi khi hợp ngữ cảnh) ·
`/lệnh` = lệnh slash (bạn chủ động gõ).

**Có 3 nhóm plugin:**

1. [Nhóm thuần Lark](#a-nhóm-thuần-lark-7-plugin) — 7 plugin, 100% chạy trên Lark, không phụ thuộc ngoài.
2. [Nhóm nghiệp vụ theo nghề](#b-nhóm-nghiệp-vụ-theo-nghề-15-plugin) — 15 plugin, lõi chạy Lark + tùy chọn nối công cụ chuyên ngành.
3. [Nhóm tiện ích & đối tác](#c-nhóm-tiện-ích--đối-tác) — PDF, quản lý plugin, và các plugin do đối tác xây.

---

## A. Nhóm thuần Lark (7 plugin)

Đây là phần "đặc sản" của bản Việt hóa — mọi thao tác đều qua một cầu nối Lark duy nhất.

### 🏗️ lark-base-deploy — Triển khai Lark Base end-to-end
**Làm gì:** biến một câu mô tả thành **cả một hệ Base hoàn chỉnh** (bảng + view + dashboard + tự
động hóa + phân quyền), qua một orchestrator **8 phase**: Discovery → Design → Build → Wire-up →
Import → Viz → Automation → Handover. Có fan-out sub-agent song song ở bước dựng bảng và dashboard.

**9 kỹ năng:** `base-deploy` (điều phối tổng), `base-discovery`, `base-design`, `base-build`,
`base-wireup`, `base-import`, `base-viz`, `base-automation`, `base-handover`.

**Khi nào dùng:** cần dựng hệ thống quản lý (CRM, chấm công, dự án, kho, tuyển dụng…) trên Lark Base
mà không muốn kéo-thả thủ công hàng giờ.

> 💬 *"Dựng một Base quản lý dự án cho team 30 người: bảng công việc + bảng nhân sự + dashboard tiến độ + nhắc deadline tự động."*
> 💬 *"Triển khai base CRM, import file khách hàng Excel này vào."*

### 📅 daily-assistant — Vận hành một ngày làm việc
**Làm gì:** "buồng lái" cho cả ngày — đọc chat, mail, lịch, việc, biên bản họp rồi trả về tóm tắt đã
xếp hạng, gộp trùng, sẵn sàng hành động.

**12 kỹ năng:** `morning-brief` (brief đầu ngày, 5 biến thể), `daily-digest` (chốt cuối ngày),
`weekly-review`, `im-digest` (lọc nhóm chat), `inbox-zero` (dọn mail), `task-prioritizer`,
`overwhelm-triage` (khi quá tải), `focus-mode` (chặn lịch + DND), `calendar-optimizer` (audit lịch
họp), `meeting-prep`, `one-on-one-prep`, `contact-360` (hồ sơ 360° một người).

**Khi nào dùng:** mỗi sáng/chiều, trước mỗi cuộc họp, khi inbox/chat quá tải.

> 💬 *"morning"* hoặc *"sáng nay tôi có gì cần làm và cần duyệt?"*
> 💬 *"47 group chat hôm nay có gì quan trọng?"*
> 💬 *"Tôi sắp gặp anh Minh — cho tôi context với anh ấy."*
> 💬 *"Block cho tôi 2 tiếng deep work chiều nay, bật DND."*

### 💼 crm-sales — CRM trên Lark Base
**Làm gì:** chạy quy trình sales trên một Base CRM — review pipeline hằng tuần, cập nhật deal sau
cuộc gọi (lấy từ biên bản họp), nhắc follow-up khách đã im lặng. **Chỉ tạo bản nháp — bạn duyệt mới gửi.**

**3 kỹ năng:** `pipeline-review`, `deal-update`, `client-followup`.

**Khi nào dùng:** đội sales dùng Lark Base làm CRM.

> 💬 *"Review pipeline tuần này, deal nào đang nghẽn?"*
> 💬 *"Cập nhật deal Acme sau cuộc gọi hôm qua."*
> 💬 *"Khách nào im lặng hơn 3 tuần? Soạn nháp follow-up."*

### ✅ governance — Phê duyệt & Quản trị
**Làm gì:** quản trị vận hành Lark — triage hàng đợi phê duyệt, đo SLA & điểm nghẽn luồng duyệt,
audit phân quyền Drive/Doc/Wiki/Base tìm rủi ro lộ dữ liệu, ghi nhật ký quyết định từ chat/biên bản
vào Base có cấu trúc.

**4 kỹ năng:** `approval-triage`, `approval-flow-sla`, `permission-audit`, `decision-logger`.

**Khi nào dùng:** quản lý nhiều đơn duyệt, lo rò rỉ quyền truy cập, cần lưu vết quyết định.

> 💬 *"Có gì cần tôi duyệt? Gợi ý duyệt/từ chối kèm lý do."*
> 💬 *"Luồng duyệt công tác phí đang nghẽn ở khâu nào?"*
> 💬 *"Quét quyền: tài liệu nào đang chia sẻ công khai hoặc ra ngoài công ty?"*

### 📚 knowledge-docs — Tri thức & Tài liệu
**Làm gì:** soạn & bảo trì tri thức trên Lark — tạo doc/wiki/sheet từ template có tên, và audit Wiki
tìm trang cũ/mồ côi/trùng lặp kèm đề xuất lưu trữ/gộp/đổi cha.

**2 kỹ năng:** `doc-from-template`, `doc-restructure`.

> 💬 *"Tạo doc báo cáo tuần theo template cho team."*
> 💬 *"Wiki phòng tôi đang bừa — rà soát và đề xuất dọn dẹp."*

### 🛠️ delivery-eng — Bàn giao Kỹ thuật
**Làm gì:** các nghi thức delivery của đội kỹ thuật — dựng postmortem sự cố không-đổ-lỗi từ dòng thời
gian chat on-call, và soạn retro cuối sprint từ ticket đã đóng, velocity, blocker.

**2 kỹ năng:** `incident-retro`, `sprint-retro`.

> 💬 *"Viết postmortem cho sự cố SEV2 tối qua từ group on-call."*
> 💬 *"Retro sprint này: ticket đóng, velocity, vướng mắc."*

### 🔧 lark-cli-dev — Công cụ phát triển lark-cli
**Làm gì:** dành cho **dev** — xây/sửa/mở rộng chính cầu nối MCP `lark-cli` (thêm công cụ, smoke-test
handshake, build lại binary, chẩn đoán mất kết nối). Tác động lên *mã nguồn* lark-cli, không phải
runtime.

**1 kỹ năng + 6 lệnh:** `lark-cli-mcp`; `/mcp-add`, `/mcp-call`, `/mcp-doctor`, `/mcp-rebuild`,
`/mcp-test`, `/mcp-tools`.

> 💬 *"/mcp-doctor"* (chẩn đoán cầu nối) · 💬 *"/mcp-add thêm công cụ gửi tin Lark có đính kèm"*

---

## B. Nhóm nghiệp vụ theo nghề (15 plugin)

Lõi giao tiếp (chat/mail/lịch/wiki/task/drive) chạy thuần Lark; các năng lực chuyên ngành (CRM,
kho dữ liệu, thanh toán…) là **tùy chọn** nối ngoài. **Mọi kỹ năng đều dùng được độc lập** (bạn dán
dữ liệu / mô tả) và **mạnh hơn** khi nối connector — xem `CONNECTORS.md` của từng plugin.

### ⚡ productivity — Năng suất (nền tảng)
Quản lý công việc + bộ nhớ nơi làm việc + dashboard trực quan. Claude học người, dự án, thuật ngữ
của bạn để hành xử như đồng nghiệp.
**4 kỹ năng:** `start`, `update`, `task-management`, `memory-management`.
> 💬 *"/start"* (khởi tạo) · 💬 *"/update --comprehensive"* (quét mail/lịch/chat tìm việc sót)
> 💬 *"nhờ Todd làm báo cáo pipeline cho deal Oracle"* (Claude hiểu nhờ bộ nhớ)

### 🔎 enterprise-search — Tìm kiếm toàn doanh nghiệp
Tìm xuyên mail, chat, tài liệu, wiki trong một chỗ; tổng hợp tri thức, digest.
**5 kỹ năng:** `search`, `search-strategy`, `knowledge-synthesis`, `digest`, `source-management`.
> 💬 *"Tìm tất cả tài liệu về dự án ABC ở mọi nơi."*
> 💬 *"Tổng hợp những gì công ty mình biết về khách hàng X."*

### ⚙️ operations — Vận hành
Quản lý nhà cung cấp, tài liệu hóa quy trình, quản lý thay đổi, kế hoạch năng lực, theo dõi tuân thủ.
**9 kỹ năng:** `vendor-review`, `process-doc`, `process-optimization`, `change-request`,
`capacity-plan`, `compliance-tracking`, `risk-assessment`, `runbook`, `status-report`.
> 💬 *"Viết quy trình onboarding nhà cung cấp mới."*
> 💬 *"Soạn runbook xử lý sự cố hệ thống thanh toán."*

### 👥 human-resources — Nhân sự
Tuyển dụng, onboarding, đánh giá hiệu suất, tra cứu chính sách, phân tích lương thưởng.
**9 kỹ năng:** `recruiting-pipeline`, `interview-prep`, `draft-offer`, `onboarding`,
`performance-review`, `policy-lookup`, `comp-analysis`, `org-planning`, `people-report`.
> 💬 *"Soạn thư mời nhận việc cho vị trí Senior Dev."*
> 💬 *"Chính sách nghỉ thai sản của công ty thế nào?"*

### 🎧 customer-support — Hỗ trợ Khách hàng
Triage ticket, soạn phản hồi, leo thang sự cố, dựng cơ sở tri thức.
**5 kỹ năng:** `ticket-triage`, `draft-response`, `customer-escalation`, `customer-research`,
`kb-article`.
> 💬 *"Phân loại hàng đợi ticket hôm nay theo độ ưu tiên."*
> 💬 *"Soạn phản hồi cho ticket khiếu nại này."*

### 📈 sales — Bán hàng
Tìm khách, soạn outreach, chiến lược deal, chuẩn bị gọi, quản pipeline.
**9 kỹ năng:** `account-research`, `call-prep`, `daily-briefing`, `draft-outreach`,
`competitive-intelligence`, `create-an-asset`, `call-summary`, `forecast`, `pipeline-review`.
> 💬 *"Nghiên cứu công ty Acme trước cuộc gọi mai."*
> 💬 *"/forecast"* (dự báo doanh số) · 💬 *"/pipeline-review"* (sức khỏe pipeline)

### 🧭 product-management — Quản lý Sản phẩm
Viết spec, roadmap, tổng hợp nghiên cứu người dùng, cập nhật bên liên quan, theo dõi chỉ số.
**8 kỹ năng + lệnh `/brainstorm`:** `write-spec`, `roadmap-update`, `synthesize-research`,
`stakeholder-update`, `metrics-review`, `competitive-brief`, `sprint-planning`,
`product-brainstorming`.
> 💬 *"Viết feature spec cho tính năng đăng nhập SSO."*
> 💬 *"/brainstorm ý tưởng tăng giữ chân người dùng."*

### 📣 marketing — Marketing
Sáng tạo nội dung, lập kế hoạch chiến dịch, phân tích hiệu quả, giữ brand voice, theo dõi đối thủ.
**8 kỹ năng:** `content-creation`, `draft-content`, `campaign-plan`, `email-sequence`,
`performance-report`, `brand-review`, `competitive-brief`, `seo-audit`.
> 💬 *"Soạn chuỗi email nuôi dưỡng 5 bước cho khách mới."*
> 💬 *"Lập kế hoạch chiến dịch ra mắt sản phẩm Q3."*

### ⚖️ legal — Pháp chế
Review hợp đồng, triage NDA, tuân thủ, brief pháp lý, phản hồi theo mẫu — cấu hình theo playbook &
mức rủi ro của tổ chức.
**9 kỹ năng:** `review-contract`, `triage-nda`, `compliance-check`, `legal-risk-assessment`,
`brief`, `legal-response`, `vendor-check`, `meeting-briefing`, `signature-request`.
> 💬 *"Review hợp đồng này, chỉ ra điều khoản rủi ro."*
> 💬 *"Triage NDA đối tác gửi — có gì cần lưu ý?"*

### 🎨 design — Thiết kế
Critique, quản lý design system, UX writing, audit accessibility, tổng hợp research, handoff cho dev.
**7 kỹ năng:** `design-critique`, `design-system`, `ux-copy`, `accessibility-review`,
`user-research`, `research-synthesis`, `design-handoff`.
> 💬 *"Critique màn hình onboarding này."*
> 💬 *"Viết UX copy cho thông báo lỗi thanh toán."*

### 📊 data — Dữ liệu
Viết SQL, khám phá dataset, phân tích thống kê, dựng biểu đồ & dashboard, kể chuyện dữ liệu.
**10 kỹ năng:** `write-query`, `sql-queries`, `explore-data`, `analyze`, `statistical-analysis`,
`validate-data`, `create-viz`, `data-visualization`, `build-dashboard`, `data-context-extractor`.
> 💬 *"Viết câu SQL tính doanh thu theo tháng năm nay."*
> 💬 *"/build-dashboard từ dataset bán hàng này."*

### 💰 finance — Tài chính & Kế toán
Bút toán, đối soát, báo cáo tài chính, phân tích chênh lệch, hỗ trợ audit, chốt sổ tháng, SOX.
**8 kỹ năng:** `journal-entry`, `journal-entry-prep`, `reconciliation`, `financial-statements`,
`variance-analysis`, `close-management`, `audit-support`, `sox-testing`.
> 💬 *"Đối soát sao kê ngân hàng với sổ cái tháng này."*
> 💬 *"Phân tích chênh lệch ngân sách vs thực chi Q2."*

### 🧑‍💻 engineering — Kỹ thuật
Standup, code review, quyết định kiến trúc, xử lý sự cố, tài liệu kỹ thuật, testing, tech debt.
**10 kỹ năng:** `standup`, `code-review`, `architecture`, `system-design`, `incident-response`,
`debug`, `documentation`, `testing-strategy`, `deploy-checklist`, `tech-debt`.
> 💬 *"Tổng hợp standup từ update trong group dev."*
> 💬 *"Đề xuất kiến trúc cho hệ thống thông báo realtime."*

### 🧬 bio-research — Nghiên cứu Sinh học
Kết nối công cụ & CSDL nghiên cứu tiền lâm sàng (tìm tài liệu, phân tích genomics, chọn target,
pipeline nf-core). Dành cho nhóm R&D life-sciences.
**6 kỹ năng:** `start`, `scientific-problem-selection`, `single-cell-rna-qc`, `scvi-tools`,
`nextflow-development`, `instrument-data-to-allotrope`. *(Giữ kết nối DB khoa học: PubMed, ChEMBL,
bioRxiv, ClinicalTrials…)*

### 🏪 small-business — Doanh nghiệp Nhỏ
**31 kỹ năng** quy trình đóng gói sẵn cho chủ shop/SMB: chốt sổ tháng, lập bảng lương, brief đầu/cuối
tuần, chiến dịch tăng trưởng, dọn CRM, đòi công nợ, xử lý khiếu nại, chuẩn bị thuế… **Bạn duyệt mọi
bước chạm tới tiền hoặc khách.** Có `smb-router` để tự định tuyến tới kỹ năng đúng.
> 💬 *"monday brief"* / *"friday brief"* · 💬 *"Lập kế hoạch trả lương tháng này."*
> 💬 *"Khách nào đang nợ hóa đơn? Soạn nhắc nhẹ nhàng."*

---

## C. Nhóm tiện ích & đối tác

### 📄 pdf-viewer — Xem & ký PDF
Mở/đánh dấu/điền form/đóng dấu/ký PDF trong trình xem tương tác, rồi tải bản đã chú thích.
**1 kỹ năng + 4 lệnh:** `view-pdf`; `/open`, `/annotate`, `/fill-form`, `/sign`.
> 💬 *"/open hợp đồng.pdf rồi đánh dấu các điều khoản cần sửa."*

### 🧩 cowork-plugin-management — Quản lý & tạo plugin
**Tự tạo và tùy biến plugin riêng** cho công ty bạn — cấu hình MCP server, chỉnh hành vi plugin,
adapt template theo cách team làm việc.
**2 kỹ năng:** `create-cowork-plugin`, `cowork-plugin-customizer`.
👉 Đây là cửa ngõ để **custom** — xem chi tiết ở [10. Tùy biến](./10-tuy-bien.md).

### Plugin do đối tác xây (`partner-built/`)
Giữ từ hệ sinh thái gốc, hữu ích khi bạn dùng các nền tảng đó:
- **apollo** — chuỗi API Apollo cho enrich/prospect/nạp sequence sales.
- **brand-voice** — chắt lọc tài liệu thương hiệu rải rác thành "guardrail" brand voice cho AI.
- **common-room** — copilot GTM: nghiên cứu account/contact, chuẩn bị gọi, soạn outreach.
- **slack** (đã chuyển sang **Lark IM**) — soạn & tìm tin nhắn Lark IM hiệu quả.
- **zoom-plugin** — lập kế hoạch/xây/debug tích hợp Zoom (REST, SDK, webhook, bot, MCP) — **30 kỹ năng**
  dành cho dev tích hợp Zoom.

---

➡️ Hiểu plugin rồi? Xem [9. Cách vận hành](./09-cach-van-hanh.md) để nắm nó chạy thế nào bên dưới,
hoặc nhảy tới [10. Tùy biến](./10-tuy-bien.md) và [11. Best practice](./11-best-practice.md).

---

## Tác giả

**Nguyễn Ngọc Tuấn**
Founder Transform Group — **Lark Platinum Partner**
🌐 Dự án: [larkcowork.com](https://larkcowork.com)
