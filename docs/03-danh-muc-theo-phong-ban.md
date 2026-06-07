# 3. Danh mục theo phòng ban

Mỗi mục dưới đây là một **bộ kỹ năng theo nghề** (plugin). Cấu trúc thống nhất:
**giải quyết gì → kỹ năng nổi bật (💬 câu lệnh mẫu → 📤 kết quả) → ai nên dùng**.

> Mẹo: bạn không cần nhớ tên kỹ năng. Cứ mô tả việc bằng tiếng Việt, trợ lý tự chọn đúng kỹ năng.
> Các câu 💬 dưới đây chỉ là ví dụ để bạn hình dung.

**Mục lục nhanh:** [Năng suất](#năng-suất-cá-nhân) · [Sales](#kinh-doanh--sales) · [Marketing](#marketing--thương-hiệu) ·
[CSKH](#chăm-sóc-khách-hàng) · [Sản phẩm](#sản-phẩm) · [Thiết kế](#thiết-kế) · [Vận hành](#vận-hành) ·
[Nhân sự](#nhân-sự) · [Tài chính](#tài-chính--kế-toán) · [Pháp lý](#pháp-lý) · [Dữ liệu](#dữ-liệu--phân-tích) ·
[Kỹ thuật](#kỹ-thuật--lập-trình) · [Doanh nghiệp nhỏ](#doanh-nghiệp-nhỏ-tất-cả-trong-một) ·
[Chuyên ngành & Tiện ích](#chuyên-ngành--tiện-ích) · [⭐ Bộ Lark-native thuần](#-bộ-lark-native-thuần-100-lark)

---

## ⭐ Bộ Lark-native thuần (100% Lark)

7 bộ "cây nhà lá vườn" — **chạy hoàn toàn trên Lark, không cần kết nối ngoài nào**. Đây là các
quy trình mạnh nhất, ghép từ những kỹ năng Lark chuyên sâu.

### 🏗️ lark-base-deploy — Dựng cả một hệ Lark Base cho team
Quy trình **8 bước** (phỏng vấn yêu cầu → thiết kế → dựng bảng → nối liên kết → nhập dữ liệu →
dashboard → tự động nhắc việc → nghiệm thu & bàn giao), chạy song song nhiều việc.
- 💬 *"Dựng hệ Base theo dõi đơn hàng cho team 30 người, có dashboard KPI và bot nhắc việc."*
  → 📤 nhiều bảng liên kết + dashboard + workflow + phân quyền — **tự nghiệm thu** trước khi bàn giao.
- **Ai dùng:** vận hành, trưởng phòng, PM cần hệ quản trị công việc.

### 🌅 daily-assistant — Điều hành một ngày trên Lark
Brief sáng, digest cuối ngày, triage hộp thư & IM, xếp ưu tiên việc, chế độ tập trung, chuẩn bị
họp/1:1, hồ sơ 360° một người.
- 💬 *"Brief sáng nay cho tôi."* / *"47 group có gì cần biết?"* / *"Bật chế độ tập trung 2 tiếng."*
  → 📤 thẻ tóm tắt gọn, phân loại cần-action/cần-biết/bỏ-qua.
- **Ai dùng:** tất cả, đặc biệt lãnh đạo & người bận.

### 💼 crm-sales — CRM chạy trên Lark Base
Review pipeline tuần, cập nhật deal sau cuộc gọi (từ biên bản Minutes), nhắc khách im lặng (chỉ soạn nháp).
- 💬 *"Review pipeline tuần."* / *"Cập nhật deal sau cuộc gọi với ABC."* → 📤 phân tích + bản ghi Base + email nháp.
- **Ai dùng:** sales, account manager.

### 🏛️ governance — Quản trị vận hành Lark
Triage hàng đợi phê duyệt, đo SLA luồng duyệt & tìm nghẽn, audit quyền truy cập, ghi nhật ký quyết định.
- 💬 *"Sáng nay có gì cần duyệt, xếp theo độ khẩn?"* / *"Nghẽn duyệt đang ở khâu nào?"* → 📤 hàng đợi + khuyến nghị + báo cáo SLA.
- **Ai dùng:** lãnh đạo, vận hành, quản trị.

### 📚 knowledge-docs — Tạo & quản lý tri thức Lark
Soạn Docs/Wiki/Sheets theo mẫu; audit & tái cấu trúc Wiki (gỡ trùng/cũ/mồ côi).
- 💬 *"Tạo doc biên bản họp theo template."* / *"Wiki đang bừa, đề xuất dọn dẹp."* → 📤 tài liệu chuẩn / kế hoạch tái cấu trúc.
- **Ai dùng:** mọi người, quản lý tri thức.

### 🛠️ delivery-eng — Nghi thức kỹ thuật trên Lark
Postmortem sự cố không-đổ-lỗi từ timeline IM; retro cuối sprint.
- 💬 *"Viết postmortem cho sự cố SEV2 hôm qua."* / *"Retro sprint này."* → 📤 postmortem / bản retro.
- **Ai dùng:** đội kỹ thuật, SRE.

### 🧰 lark-cli-dev — Mở rộng chính cây cầu nối MCP
Thêm/sửa công cụ trong `cmd/mcp/`, debug bridge (dành cho dev tự mở rộng hệ thống).
- **Ai dùng:** admin/dev kỹ thuật.

---

## Năng suất cá nhân

### 📌 productivity — Quản lý việc & "trí nhớ công việc"
Gom việc, lên kế hoạch ngày, và **ghi nhớ ngữ cảnh** (ai là ai, dự án nào, từ viết tắt) để trợ lý
hiểu bạn như đồng nghiệp lâu năm. Lark Task là nguồn sự thật.
- 💬 *"Hôm nay tôi có gì trên bàn, việc nào gấp?"* → 📤 thẻ tóm tắt việc đến hạn/quá hạn, xếp ưu tiên.
- 💬 *"Xong việc gửi báo cáo cho sếp rồi."* → 📤 đánh dấu hoàn thành trong Lark Task.
- 💬 *"Ghi nhớ: 'PSR' nghĩa là Báo cáo tình hình pipeline."* → 📤 lưu vào bộ nhớ, lần sau tự hiểu.
- **Ai dùng:** tất cả mọi người.

### 🔎 enterprise-search — Tìm mọi thứ trong công ty
Một câu hỏi, tìm xuyên **chat + mail + tài liệu + wiki** cùng lúc; tổng hợp, khử trùng lặp, kèm nguồn.
- 💬 *"Tìm tài liệu về chính sách công tác phí mới nhất."* → 📤 danh sách kết quả + trích dẫn nguồn.
- 💬 *"Tóm tắt mọi hoạt động liên quan dự án Phoenix tuần này."* → 📤 digest gọn theo nguồn.
- **Ai dùng:** mọi người, đặc biệt khi "biết là có mà không nhớ ở đâu".

---

## Kinh doanh & Sales

### 💼 sales — Tăng tốc bán hàng từ tìm khách đến chốt deal
- 💬 *"Nghiên cứu công ty ABC và người mình sắp gặp."* → 📤 hồ sơ khách + tình báo bán hàng (`account-research`).
- 💬 *"Chuẩn bị cho cuộc gọi với ABC chiều nay."* → 📤 ngữ cảnh tài khoản + nghị trình gợi ý (`call-prep`).
- 💬 *"Tóm tắt cuộc gọi này, rút việc cần làm và soạn email follow-up."* → 📤 action items + email nháp (`call-summary`).
- 💬 *"Tổng quan pipeline tuần này, deal nào đang kẹt?"* → 📤 phân tích theo stage + cảnh báo rủi ro + kế hoạch (`pipeline-review`).
- 💬 *"Dự báo doanh số quý này (tốt/khả dĩ/xấu)."* → 📤 forecast có kịch bản (`forecast`).
- 💬 *"Soạn outreach cá nhân hoá cho prospect X."* → 📤 thư nháp dựa trên nghiên cứu (`draft-outreach`).
- **Ai dùng:** AE, SDR, trưởng phòng kinh doanh.
- **Tăng lực (tuỳ chọn, kết nối ngoài):** `apollo` (làm giàu dữ liệu lead + nạp vào sequence),
  `common-room` (tín hiệu cộng đồng/ý định mua, prep cuộc gọi).

---

## Marketing & Thương hiệu

### 📣 marketing — Sản xuất nội dung, lên chiến dịch, đo hiệu quả
- 💬 *"Lên brief chiến dịch ra mắt sản phẩm Q3."* → 📤 mục tiêu, đối tượng, kênh, lịch nội dung (`campaign-plan`).
- 💬 *"Viết bài blog SEO về chủ đề X."* / *"Soạn chuỗi 5 email nurturing."* → 📤 bản nháp đầy đủ (`draft-content`, `email-sequence`).
- 💬 *"Audit SEO website."* → 📤 nghiên cứu từ khoá, lỗi on-page, khoảng trống nội dung (`seo-audit`).
- 💬 *"Báo cáo hiệu quả marketing tháng này."* → 📤 chỉ số chính + xu hướng + đề xuất (`performance-report`).
- 💬 *"Kiểm tra bài này có đúng giọng thương hiệu không?"* → 📤 chỗ lệch + sửa gợi ý (`brand-review`).
- **Ai dùng:** marketing, content, growth.
- **Tăng lực:** `brand-voice` — tự gom tài liệu thương hiệu rải rác thành "luật" giọng nói áp dụng tự động.

---

## Chăm sóc khách hàng

### 🎧 customer-support — Triage ticket, soạn phản hồi, dựng KB
- 💬 *"Phân loại & ưu tiên ticket mới này."* → 📤 mức ưu tiên + hướng xử lý (`ticket-triage`).
- 💬 *"Soạn phản hồi chuyên nghiệp cho khách đang bực về việc giao hàng trễ."* → 📤 thư nháp đúng tông (`draft-response`).
- 💬 *"Đóng gói vụ này để escalate cho kỹ thuật."* → 📤 gói ngữ cảnh đầy đủ + (tuỳ) tạo *đơn phê duyệt* (`customer-escalation`).
- 💬 *"Viết bài hướng dẫn từ ca vừa xử lý xong."* → 📤 bài KB đăng vào *Wiki* (`kb-article`).
- **Ai dùng:** CSKH, success, helpdesk.

---

## Sản phẩm

### 🧭 product-management — Spec, roadmap, tổng hợp nghiên cứu
- 💬 *"Viết spec/PRD cho tính năng đăng nhập bằng OTP."* → 📤 PRD có cấu trúc (`write-spec`).
- 💬 *"Cập nhật roadmap, thêm sáng kiến X và sắp lại ưu tiên."* → 📤 roadmap mới (`roadmap-update`).
- 💬 *"Tổng hợp 12 buổi phỏng vấn người dùng thành insight."* → 📤 chủ đề + insight + đề xuất (`synthesize-research`).
- 💬 *"Soạn cập nhật cho stakeholder tuần này."* → 📤 bản update theo đối tượng (`stakeholder-update`).
- 💬 *"Brainstorm hướng giải bài toán giữ chân người dùng."* → 📤 đối thoại phản biện, khung tư duy (`product-brainstorming`).
- **Ai dùng:** PM, product owner.

## Thiết kế

### 🎨 design — Critique, design system, UX writing, accessibility
- 💬 *"Review thiết kế màn hình thanh toán này."* → 📤 góp ý usability/hierarchy/consistency (`design-critique`).
- 💬 *"Audit accessibility theo WCAG 2.1 AA."* → 📤 danh sách lỗi + cách sửa (`accessibility-review`).
- 💬 *"Viết microcopy cho trạng thái rỗng & thông báo lỗi."* → 📤 UX copy (`ux-copy`).
- 💬 *"Tổng hợp nghiên cứu người dùng thành insight."* → 📤 themes + khuyến nghị (`research-synthesis`).
- **Ai dùng:** designer, UX writer, design lead.

---

## Vận hành

### ⚙️ operations — Quy trình, nhà cung cấp, rủi ro, năng lực
- 💬 *"Viết SOP + sơ đồ + ma trận RACI cho quy trình nhập kho."* → 📤 tài liệu quy trình (`process-doc`).
- 💬 *"Quy trình này chậm, tối ưu giúp tôi."* → 📤 phân tích + đề xuất cải tiến (`process-optimization`).
- 💬 *"Đánh giá nhà cung cấp X: chi phí, rủi ro, khuyến nghị."* → 📤 bản đánh giá (`vendor-review`).
- 💬 *"Tạo change request có phân tích tác động và kế hoạch rollback."* → 📤 phiếu CR (`change-request`).
- 💬 *"Báo cáo trạng thái tuần: KPI, rủi ro, việc cần làm."* → 📤 status report (`status-report`).
- **Ai dùng:** vận hành, COO, quản lý dự án. *Sổ đăng ký (risk/vendor/asset) nên đặt trong Base.*

---

## Nhân sự

### 👥 human-resources — Tuyển dụng, onboarding, đánh giá, lương
- 💬 *"Mức lương cho vị trí Senior BE có hợp khung không?"* → 📤 benchmark + đặt band + mô hình equity (`comp-analysis`).
- 💬 *"Tạo kế hoạch phỏng vấn theo năng lực cho vị trí PM."* → 📤 bộ câu hỏi + scorecard (`interview-prep`).
- 💬 *"Soạn thư mời nhận việc cho ứng viên A."* → 📤 offer letter nháp (`draft-offer`).
- 💬 *"Tạo checklist onboarding tuần đầu cho nhân viên mới."* → 📤 kế hoạch onboarding (`onboarding`).
- 💬 *"Chính sách nghỉ phép của công ty là gì?"* → 📤 giải thích dễ hiểu (`policy-lookup`).
- 💬 *"Báo cáo nhân sự: headcount, nghỉ việc, đa dạng."* → 📤 people report (`people-report`).
- **Ai dùng:** HR, HRBP, quản lý tuyển dụng. ⚠️ Offer/đổi lương đi qua *luồng phê duyệt*.

---

## Tài chính & Kế toán

### 💰 finance — Bút toán, đối soát, BCTC, phân tích chênh lệch
- 💬 *"Soạn bút toán cho khoản chi marketing tháng 6."* → 📤 bút toán Nợ/Có + chứng từ (`journal-entry`).
- 💬 *"Đối soát số dư GL với sao kê ngân hàng."* → 📤 bảng đối soát + chênh lệch (`reconciliation`).
- 💬 *"Lập báo cáo KQKD / cân đối kế toán kỳ này."* → 📤 BCTC có so sánh kỳ (`financial-statements`).
- 💬 *"Phân tích chênh lệch chi phí so với ngân sách."* → 📤 phân rã nguyên nhân + waterfall (`variance-analysis`).
- 💬 *"Quản lý quy trình chốt sổ tháng."* → 📤 chuỗi công việc + phụ thuộc + trạng thái (`close-management`).
- **Ai dùng:** kế toán, kiểm toán nội bộ, FP&A. *Kho dữ liệu (BigQuery/Snowflake…) giữ kết nối riêng.*

---

## Pháp lý

### ⚖️ legal — Rà hợp đồng, triage NDA, tuân thủ
- 💬 *"Triage NDA này — đèn xanh/vàng/đỏ?"* → 📤 phân loại nhanh + lý do (`triage-nda`).
- 💬 *"Rà hợp đồng này theo playbook của công ty."* → 📤 các điểm lệch + đề xuất chỉnh (`review-contract`).
- 💬 *"Đánh giá rủi ro pháp lý của việc ra tính năng X."* → 📤 ma trận mức độ × khả năng + cảnh báo leo thang (`legal-risk-assessment`).
- 💬 *"Kiểm tra tuân thủ cho sáng kiến này."* → 📤 nghĩa vụ cần lưu ý (`compliance-check`).
- 💬 *"Chuẩn bị tài liệu ký điện tử cho hợp đồng này."* → 📤 checklist trước ký + định tuyến ký (`signature-request`).
- **Ai dùng:** legal in-house, hành chính-pháp chế.

---

## Dữ liệu & Phân tích

### 📊 data — SQL, khám phá dữ liệu, dashboard, kiểm định
- 💬 *"Tháng này doanh thu theo vùng thế nào?"* → 📤 viết SQL, chạy, trả số + nhận định (`analyze`, `write-query`).
- 💬 *"Khám phá bộ dữ liệu này: chất lượng, phân bố, bất thường."* → 📤 hồ sơ dữ liệu (`explore-data`).
- 💬 *"Dựng dashboard tương tác có filter cho doanh số."* → 📤 dashboard HTML (`build-dashboard`).
- 💬 *"Kiểm định lại phân tích này trước khi gửi sếp."* → 📤 rà phương pháp + sai số + thiên lệch (`validate-data`).
- **Ai dùng:** data analyst, BI, ai cần trả lời bằng số. *Hỗ trợ Snowflake/BigQuery/Databricks…*

---

## Kỹ thuật & Lập trình

### 🛠️ engineering — Standup, review code, kiến trúc, sự cố
- 💬 *"Tạo cập nhật standup từ hoạt động gần đây của tôi."* → 📤 bản standup (`standup`).
- 💬 *"Review PR này về bảo mật/hiệu năng/đúng đắn."* → 📤 nhận xét có cấu trúc (`code-review`).
- 💬 *"Viết ADR chọn giữa PostgreSQL và MongoDB."* → 📤 quyết định kiến trúc (`architecture`).
- 💬 *"Chạy quy trình xử lý sự cố + viết postmortem."* → 📤 triage + thông báo + postmortem (`incident-response`).
- **Ai dùng:** dev, tech lead, SRE. *Giữ GitHub/Datadog…; Lark lo phần chat/tài liệu/việc.*
- **Liên quan:** `zoom-plugin` — bộ tham chiếu dựng tích hợp Zoom (REST/SDK/webhook) cho đội kỹ thuật.

---

## Doanh nghiệp nhỏ (tất-cả-trong-một)

### 🏪 small-business — 31 quy trình "chủ doanh nghiệp" đóng gói sẵn
Bộ lớn nhất, dành cho chủ SMB không có đội chuyên trách. Có "cửa trước" thông minh:
- 💬 *"Tôi cần gì đó cho việc kinh doanh."* → 📤 trợ lý hỏi nhanh rồi điều hướng đúng việc (`smb-router`).
- 💬 *"Tóm tắt sức khỏe kinh doanh hôm nay."* → 📤 1 trang: tiền mặt, doanh số, pipeline (`business-pulse`).
- 💬 *"Brief sáng thứ Hai."* / *"Pulse chiều thứ Sáu."* → 📤 báo cáo đầu/cuối tuần (`monday-brief`, `friday-brief`).
- 💬 *"Dự báo dòng tiền 30 ngày tới, có gì cần lo?"* → 📤 cash-flow outlook (`cash-flow-snapshot`, `month-heads-up`).
- 💬 *"Soạn email nhắc các hoá đơn quá hạn."* → 📤 thư nhắc nháp theo dữ liệu QuickBooks/PayPal (`invoice-chase`).
- 💬 *"Chốt sổ tháng giúp tôi."* → 📤 đối soát + P&L + việc còn thiếu (`close-month`, `month-end-prep`).
- 💬 *"Chuẩn bị tài liệu mùa thuế."* → 📤 ước tính thuế quý / gói 1099 (`tax-prep`).
- 💬 *"Xử lý khiếu nại khách này từ đầu đến cuối."* → 📤 kéo ngữ cảnh + soạn phản hồi (`handle-complaint`).
- 💬 *"Chạy nguyên một chiến dịch marketing."* → 📤 phân tích → brief → thiết kế Canva → đẩy HubSpot (`run-campaign`).
- **Ai dùng:** chủ shop/SMB, quản lý đa nhiệm. *Giữ QuickBooks/Stripe/Square/HubSpot/Canva; Lark lo chat/lịch/tài liệu.*

---

## Chuyên ngành & Tiện ích

| Bộ | Dùng để | Ai |
|---|---|---|
| 🧬 **bio-research** | Nghiên cứu tiền lâm sàng: tìm tài liệu, phân tích genomics, chọn target, pipeline nf-core | nhóm R&D life-sciences |
| 📄 **pdf-viewer** | Mở/đánh dấu/ký PDF trong trình xem tương tác (đánh dấu hợp đồng, điền form, đóng dấu duyệt) | mọi người xử lý PDF |
| 💬 **slack** (Lark IM) | Soạn & tìm tin nhắn Lark IM hiệu quả (đã chuyển sang dùng Lark IM) | mọi người |
| 🧩 **cowork-plugin-management** | Tự tạo/tuỳ biến bộ kỹ năng riêng cho công ty bạn | admin, người dựng quy trình |

---

➡️ Xem các bộ này phối hợp ra sao trong [4. Một ngày cùng Lark Cowork](./04-mot-ngay-cung-lark-cowork.md).

---

## Tác giả

**Nguyễn Ngọc Tuấn**
Founder Transform Group — **Lark Platinum Partner**
🌐 Dự án: [larkcowork.com](https://larkcowork.com)
