# Plugin Năng suất Pháp chế (Legal)

Một plugin năng suất được hỗ trợ bởi AI dành cho các đội pháp chế nội bộ (in-house legal), được thiết kế chủ yếu cho [Cowork](https://claude.com/product/cowork), ứng dụng desktop dạng agentic của Anthropic — dù nó cũng hoạt động trong Claude Code. Tự động hóa việc review hợp đồng, triage NDA, các quy trình tuân thủ (compliance), brief pháp lý, và các phản hồi theo mẫu -- tất cả đều có thể cấu hình theo playbook và mức độ chấp nhận rủi ro cụ thể của tổ chức bạn.

> **Tuyên bố miễn trừ trách nhiệm:** Plugin này hỗ trợ các quy trình pháp lý nhưng không cung cấp tư vấn pháp lý. Luôn xác minh các kết luận với chuyên gia pháp lý có chuyên môn. Phân tích do AI tạo ra cần được luật sư có chứng chỉ hành nghề xem xét lại trước khi được dùng làm căn cứ cho các quyết định pháp lý. Các ví dụ playbook mặc định trong plugin này phản ánh các quan điểm và thẩm quyền tài phán pháp lý của Hoa Kỳ (Delaware, New York, California). Nếu bạn vận hành theo những hệ thống pháp luật khác (EU, UK, Hà Lan, Úc, v.v.), bạn phải tùy chỉnh playbook trong .claude/legal.local.md để phản ánh các yêu cầu pháp lý cụ thể của thẩm quyền tài phán của bạn, các điều khoản hợp đồng tiêu chuẩn, và các nghĩa vụ tuân thủ trước khi dựa vào phân tích của plugin.

## Nhóm đối tượng mục tiêu

- **Commercial Counsel (Luật sư thương mại)** -- Đàm phán hợp đồng, quản lý nhà cung cấp, hỗ trợ deal
- **Product Counsel (Luật sư sản phẩm)** -- Review sản phẩm, điều khoản dịch vụ, chính sách quyền riêng tư, các vấn đề sở hữu trí tuệ (IP)
- **Privacy / Compliance (Quyền riêng tư / Tuân thủ)** -- Các quy định bảo vệ dữ liệu, review DPA, yêu cầu của chủ thể dữ liệu, giám sát quy định
- **Litigation Support (Hỗ trợ tố tụng)** -- Lệnh giữ tài liệu (discovery hold), chuẩn bị review tài liệu, brief vụ việc

## Cài đặt

```
claude plugins add knowledge-work-plugins/legal
```

## Bắt đầu nhanh

### 1. Cài đặt plugin

```
claude plugins add knowledge-work-plugins/legal
```

### 2. Cấu hình playbook của bạn

Tạo một file cấu hình cục bộ để định nghĩa các quan điểm tiêu chuẩn của tổ chức bạn. Đây là nơi bạn mã hóa playbook đàm phán của team, các mức chấp nhận rủi ro, và các điều khoản tiêu chuẩn.

Tạo một file `legal.local.md` ở nơi Claude có thể tìm thấy:

- **Cowork**: Lưu nó trong bất kỳ thư mục nào bạn đã chia sẻ với Cowork (qua bộ chọn thư mục). Plugin sẽ tự động tìm thấy.
- **Claude Code**: Lưu nó trong thư mục `.claude/` của dự án.

```markdown
# Legal Playbook Configuration

## Contract Review Positions

### Limitation of Liability
- Standard position: Mutual cap at 12 months of fees paid/payable
- Acceptable range: 6-24 months of fees
- Escalation trigger: Uncapped liability, consequential damages inclusion

### Indemnification
- Standard position: Mutual indemnification for IP infringement and data breach
- Acceptable: Indemnification limited to third-party claims only
- Escalation trigger: Unilateral indemnification obligations, uncapped indemnification

### IP Ownership
- Standard position: Each party retains pre-existing IP; customer owns customer data
- Escalation trigger: Broad IP assignment clauses, work-for-hire provisions for pre-existing IP

### Data Protection
- Standard position: Require DPA for any personal data processing
- Requirements: Sub-processor notification, data deletion on termination, breach notification within 72 hours
- Escalation trigger: No DPA offered, cross-border transfer without safeguards

### Term and Termination
- Standard position: Annual term with 30-day termination for convenience
- Acceptable: Multi-year with termination for convenience after initial term
- Escalation trigger: Auto-renewal without notice period, no termination for convenience

### Governing Law
- Preferred: [Your jurisdiction]
- Acceptable: Major commercial jurisdictions (NY, DE, CA, England & Wales)
- Escalation trigger: Non-standard jurisdictions, mandatory arbitration in unfavorable venue

## NDA Defaults
- Mutual obligations required
- Term: 2-3 years standard, 5 years for trade secrets
- Standard carveouts: independently developed, publicly available, rightfully received from third party
- Residuals clause: acceptable if narrowly scoped

## Response Templates
Configure paths to your template files or define inline templates for common inquiries.
```

### 3. Kết nối các công cụ của bạn

Plugin hoạt động tốt nhất khi được kết nối với các công cụ hiện có của bạn qua MCP. Các máy chủ được cấu hình sẵn bao gồm Lark IM, Box, Egnyte, Lark và Lark. Xem [CONNECTORS.md](CONNECTORS.md) để có danh sách đầy đủ các hạng mục và tùy chọn được hỗ trợ.

## Lệnh

### `/review-contract` -- Review hợp đồng đối chiếu với playbook

Review một hợp đồng đối chiếu với playbook đàm phán của tổ chức bạn. Đánh dấu các điểm sai lệch, tạo bản redline, và cung cấp phân tích tác động đến hoạt động kinh doanh.

```
/review-contract
```

Chấp nhận: tải file lên, URL, hoặc dán nội dung hợp đồng. Sẽ hỏi về ngữ cảnh (bên của bạn là bên nào, deadline, các trọng tâm cần chú ý) và review từng điều khoản một, đối chiếu với playbook bạn đã cấu hình.

### `/triage-nda` -- Sàng lọc sơ bộ NDA

Triage nhanh các NDA đến đối chiếu với các tiêu chí tiêu chuẩn. Phân loại thành GREEN (phê duyệt tiêu chuẩn), YELLOW (cần luật sư review), hoặc RED (có vấn đề đáng kể).

```
/triage-nda
```

### `/vendor-check` -- Trạng thái thỏa thuận với nhà cung cấp

Kiểm tra trạng thái các thỏa thuận hiện có với một nhà cung cấp trên các hệ thống bạn đã kết nối.

```
/vendor-check [vendor name]
```

Báo cáo về các NDA, MSA, DPA hiện có, ngày hết hạn, và các điều khoản chính.

### `/brief` -- Brief cho đội pháp chế

Tạo các bản brief theo ngữ cảnh cho công việc pháp lý của bạn.

```
/brief daily          # Brief buổi sáng về các mục liên quan đến pháp lý
/brief topic [query]  # Brief nghiên cứu về một câu hỏi pháp lý cụ thể
/brief incident       # Brief nhanh về một tình huống đang diễn biến
```

### `/respond` -- Tạo phản hồi theo mẫu

Tạo một phản hồi từ các mẫu bạn đã cấu hình cho các loại yêu cầu phổ biến.

```
/respond [inquiry-type]
```

Các loại yêu cầu được hỗ trợ bao gồm: yêu cầu của chủ thể dữ liệu (data subject request), lệnh giữ tài liệu (discovery hold), câu hỏi về nhà cung cấp, yêu cầu NDA, và các hạng mục tùy chỉnh do bạn định nghĩa.

## Kỹ năng

| Kỹ năng | Mô tả |
|-------|-------------|
| `contract-review` | Phân tích hợp đồng dựa trên playbook, phân loại điểm sai lệch, tạo redline |
| `nda-triage` | Tiêu chí sàng lọc NDA, quy tắc phân loại, khuyến nghị định tuyến |
| `compliance` | Các quy định về quyền riêng tư (GDPR, CCPA), review DPA, yêu cầu của chủ thể dữ liệu |
| `canned-responses` | Quản lý mẫu, các hạng mục phản hồi, các kích hoạt chuyển cấp (escalation trigger) |
| `legal-risk-assessment` | Khung mức độ nghiêm trọng của rủi ro, các cấp phân loại, tiêu chí chuyển cấp |
| `meeting-briefing` | Phương pháp luận chuẩn bị họp, thu thập ngữ cảnh, theo dõi action item |

## Quy trình mẫu

### Review hợp đồng

1. Nhận một hợp đồng nhà cung cấp qua email
2. Chạy `/review-contract` và tải tài liệu lên
3. Cung cấp ngữ cảnh: "Chúng tôi là bên khách hàng, cần chốt trước cuối quý, tập trung vào bảo vệ dữ liệu và trách nhiệm pháp lý"
4. Nhận phân tích từng điều khoản một với các cờ GREEN/YELLOW/RED
5. Nhận ngôn ngữ redline cụ thể cho các mục YELLOW và RED
6. Chia sẻ phân tích với đội deal của bạn

### Triage NDA

1. Đội sales gửi một NDA từ một khách hàng tiềm năng mới
2. Chạy `/triage-nda` rồi dán hoặc tải NDA lên
3. Nhận phân loại tức thì: GREEN (chuyển đi ký), YELLOW (các vấn đề cụ thể cần review), hoặc RED (cần luật sư review đầy đủ)
4. Với các NDA GREEN, phê duyệt trực tiếp; với YELLOW/RED, xử lý các vấn đề đã được đánh dấu

### Brief hằng ngày

1. Bắt đầu buổi sáng của bạn với `/brief daily`
2. Nhận một bản tóm tắt về các yêu cầu hợp đồng phát sinh qua đêm, câu hỏi tuân thủ, các deadline sắp tới, và các mục trên lịch cần chuẩn bị về mặt pháp lý
3. Sắp xếp ưu tiên cho ngày của bạn dựa trên mức độ khẩn cấp và deadline

### Kiểm tra nhà cung cấp (Vendor Check)

1. Đội kinh doanh hỏi về một hợp tác mới với một nhà cung cấp hiện có
2. Chạy `/vendor-check Acme Corp`
3. Xem các thỏa thuận hiện có, ngày hết hạn, và các điều khoản chính trong nháy mắt
4. Biết ngay liệu bạn có cần một NDA mới hay có thể tiến hành theo các điều khoản hiện có

## Tích hợp MCP

> Nếu bạn thấy các placeholder lạ hoặc cần kiểm tra công cụ nào đang được kết nối, xem [CONNECTORS.md](CONNECTORS.md).

Plugin kết nối với các công cụ của bạn thông qua các máy chủ MCP (Model Context Protocol):

| Hạng mục | Ví dụ | Mục đích |
|----------|----------|---------|
| Chat | Lark IM, Teams | Yêu cầu của team, thông báo, triage |
| Lưu trữ đám mây | Box, Egnyte | Playbook, mẫu, tiền lệ |
| Bộ ứng dụng văn phòng | Lark | Email, lịch, tài liệu |
| Trình theo dõi dự án | Lark (Lark Task/Lark Wiki) | Theo dõi vụ việc, công việc |

Xem [CONNECTORS.md](CONNECTORS.md) để có danh sách đầy đủ các tích hợp được hỗ trợ, bao gồm CLM, CRM, chữ ký điện tử (e-signature), và các tùy chọn bổ sung.

Cấu hình các kết nối trong `.mcp.json`. Plugin suy giảm một cách mượt mà khi các công cụ không khả dụng -- nó sẽ ghi chú các khoảng trống và đề xuất kiểm tra thủ công.

## Tùy chỉnh

### Cấu hình playbook

Playbook của bạn là trái tim của hệ thống review hợp đồng. Định nghĩa các quan điểm của bạn trong `legal.local.md`:

- **Quan điểm tiêu chuẩn (Standard positions)**: Các điều khoản hợp đồng bạn ưu tiên
- **Khoảng chấp nhận (Acceptable ranges)**: Những gì bạn có thể đồng ý mà không cần chuyển cấp
- **Kích hoạt chuyển cấp (Escalation triggers)**: Các điều khoản đòi hỏi review cấp cao hơn hoặc luật sư bên ngoài

### Mẫu phản hồi

Định nghĩa các mẫu cho các yêu cầu phổ biến. Các mẫu hỗ trợ thay thế biến và bao gồm các kích hoạt chuyển cấp tích hợp sẵn cho những tình huống không nên dùng phản hồi theo mẫu.

### Khung rủi ro

Tùy chỉnh ma trận đánh giá rủi ro để khớp với khẩu vị rủi ro và sơ đồ phân loại của tổ chức bạn.

## Cấu trúc file

```
legal/
├── .claude-plugin/plugin.json
├── .mcp.json
├── README.md
├── commands/
│   ├── review-contract.md
│   ├── triage-nda.md
│   ├── vendor-check.md
│   ├── brief.md
│   └── respond.md
└── skills/
    ├── contract-review/SKILL.md
    ├── nda-triage/SKILL.md
    ├── compliance/SKILL.md
    ├── canned-responses/SKILL.md
    ├── legal-risk-assessment/SKILL.md
    └── meeting-briefing/SKILL.md
```

---

## Tác giả

**Nguyễn Ngọc Tuấn**
Founder Transform Group — **Lark Platinum Partner**
🌐 Dự án: [larkcowork.com](https://larkcowork.com)
