# Plugin Bán hàng (Sales)

Plugin nâng cao năng suất bán hàng được thiết kế chủ yếu cho [Cowork](https://claude.com/product/cowork), ứng dụng desktop dạng agent của Anthropic — dù vậy nó cũng hoạt động được trong Claude Code. Hỗ trợ tìm kiếm khách hàng tiềm năng, tiếp cận, quản lý pipeline, chuẩn bị cuộc gọi và chiến lược deal. Hoạt động được với mọi đội ngũ bán hàng — dùng độc lập với web search và thông tin bạn cung cấp, tăng sức mạnh khi bạn kết nối CRM, email và các công cụ khác.

## Cài đặt

```bash
claude plugins add knowledge-work-plugins/sales
```

## Lệnh

Các quy trình tường minh bạn gọi bằng lệnh slash:

| Lệnh | Mô tả |
|---|---|
| `/call-summary` | Xử lý ghi chú hoặc bản ghi cuộc gọi — trích xuất các đầu việc, soạn email theo dõi, tạo bản tóm tắt nội bộ |
| `/forecast` | Tạo dự báo bán hàng có trọng số — tải lên CSV hoặc mô tả pipeline, đặt chỉ tiêu (quota), nhận dự phóng |
| `/pipeline-review` | Phân tích sức khỏe pipeline — ưu tiên hóa deal, đánh dấu rủi ro, nhận kế hoạch hành động tuần |

Tất cả các lệnh đều hoạt động **độc lập** (dán ghi chú, tải lên CSV, hoặc mô tả tình huống của bạn) và được **tăng sức mạnh** với các connector MCP.

## Kỹ năng

Kiến thức chuyên môn mà Claude tự động dùng khi có liên quan:

| Kỹ năng | Mô tả |
|---|---|
| `account-research` | Nghiên cứu một công ty hoặc một người — web search để tìm thông tin công ty, các liên hệ chủ chốt, tin tức gần đây, tín hiệu tuyển dụng |
| `call-prep` | Chuẩn bị cho các cuộc gọi bán hàng — bối cảnh tài khoản, nghiên cứu người tham dự, đề xuất chương trình họp, câu hỏi khám phá nhu cầu |
| `daily-briefing` | Bản tóm tắt bán hàng hằng ngày được ưu tiên hóa — cuộc họp, cảnh báo pipeline, email ưu tiên, hành động đề xuất |
| `draft-outreach` | Tiếp cận theo hướng nghiên cứu trước — nghiên cứu khách hàng tiềm năng, rồi soạn email và tin nhắn LinkedIn cá nhân hóa |
| `competitive-intelligence` | Nghiên cứu đối thủ — so sánh sản phẩm, thông tin giá, các bản phát hành gần đây, ma trận khác biệt hóa, kịch bản trao đổi bán hàng |
| `create-an-asset` | Tạo các ấn phẩm bán hàng tùy chỉnh — landing page, bộ slide, tờ một trang, demo quy trình được điều chỉnh cho đúng khách hàng tiềm năng của bạn |

## Quy trình mẫu

### Sau một cuộc gọi

```
/call-summary
```

Dán ghi chú hoặc bản ghi của bạn. Nhận bản tóm tắt có cấu trúc, các đầu việc kèm người phụ trách, và một email theo dõi dạng nháp. Nếu đã kết nối CRM, sẽ đề nghị ghi nhật ký hoạt động và tạo nhiệm vụ.

### Dự báo hằng tuần

```
/forecast
```

Tải lên một bản xuất CSV từ CRM của bạn (hoặc dán các deal). Cho tôi biết chỉ tiêu và mốc thời gian của bạn. Nhận dự báo có trọng số với các kịch bản tốt nhất/khả năng cao/xấu nhất, phân tách commit và upside, và phân tích khoảng cách (gap).

### Rà soát Pipeline

```
/pipeline-review
```

Tải lên CSV hoặc mô tả pipeline của bạn. Nhận điểm sức khỏe, ưu tiên hóa deal, cờ rủi ro (deal cũ, đã quá ngày dự kiến chốt, chỉ có một đầu mối liên hệ), và một kế hoạch hành động tuần.

### Nghiên cứu một khách hàng tiềm năng

Chỉ cần hỏi một cách tự nhiên:
```
Research Acme Corp before my call tomorrow
```

Kỹ năng `account-research` sẽ tự động kích hoạt và cho bạn tổng quan công ty, các liên hệ chủ chốt, tin tức gần đây, và cách tiếp cận được khuyến nghị.

### Soạn email tiếp cận

```
Draft an email to the VP of Engineering at TechStart
```

Kỹ năng `draft-outreach` sẽ nghiên cứu khách hàng tiềm năng trước, rồi tạo nội dung tiếp cận cá nhân hóa với nhiều góc tiếp cận khác nhau.

### Thông tin cạnh tranh

```
How do we compare to Competitor X?
```

Kỹ năng `competitive-intelligence` sẽ nghiên cứu cả hai công ty và dựng một ma trận khác biệt hóa kèm các kịch bản trao đổi.

## Dùng độc lập + Tăng sức mạnh

Mọi lệnh và kỹ năng đều hoạt động được mà không cần bất kỳ tích hợp nào:

| Việc bạn có thể làm | Dùng độc lập | Tăng sức mạnh với |
|-----------------|------------|-------------------|
| Xử lý ghi chú cuộc gọi | Dán ghi chú/bản ghi | MCP bản ghi (vd Gong, Lark Minutes) |
| Dự báo pipeline | Tải lên CSV, dán deal | MCP CRM |
| Rà soát pipeline | Tải lên CSV, mô tả deal | MCP CRM |
| Nghiên cứu khách hàng tiềm năng | Web search | MCP làm giàu dữ liệu (vd Clay, ZoomInfo) |
| Chuẩn bị cho cuộc gọi | Mô tả cuộc họp | MCP CRM, Email, Calendar |
| Soạn email tiếp cận | Web search + bối cảnh của bạn | MCP CRM, Email |
| Thông tin cạnh tranh | Web search | CRM (dữ liệu thắng/thua), Docs (battlecard) |

## Tích hợp MCP

> Nếu bạn thấy các placeholder lạ hoặc cần kiểm tra công cụ nào đang được kết nối, xem [CONNECTORS.md](CONNECTORS.md).

Kết nối các công cụ của bạn để có trải nghiệm phong phú hơn:

| Danh mục | Ví dụ | Khả năng mở khóa |
|---|---|---|
| **CRM** | HubSpot, Close | Dữ liệu pipeline, lịch sử tài khoản, bản ghi liên hệ |
| **Bản ghi (Transcripts)** | Lark Minutes, Gong, Chorus | Bản ghi âm cuộc gọi, bản ghi nội dung, các khoảnh khắc quan trọng |
| **Làm giàu dữ liệu (Enrichment)** | Clay, ZoomInfo, Apollo | Làm giàu dữ liệu công ty và liên hệ |
| **Chat** | Lark IM, Teams | Thảo luận nội bộ, thông tin từ đồng nghiệp |

Xem [CONNECTORS.md](CONNECTORS.md) để biết danh sách đầy đủ các tích hợp được hỗ trợ, bao gồm email, calendar, và các tùy chọn CRM bổ sung.

## Cấu hình

Tạo một tệp `settings.local.json` để cá nhân hóa:

- **Cowork**: Lưu nó trong bất kỳ thư mục nào bạn đã chia sẻ với Cowork (qua bộ chọn thư mục). Plugin sẽ tự tìm thấy nó.
- **Claude Code**: Lưu nó tại `sales/.claude/settings.local.json`.

```json
{
  "name": "Your Name",
  "title": "Account Executive",
  "company": "Your Company",
  "quota": {
    "annual": 1000000,
    "quarterly": 250000
  },
  "product": {
    "name": "Your Product",
    "value_props": [
      "Key value proposition 1",
      "Key value proposition 2"
    ],
    "competitors": [
      "Competitor A",
      "Competitor B"
    ]
  }
}
```

Plugin sẽ hỏi bạn thông tin này một cách tương tác nếu nó chưa được cấu hình.

---

## Tác giả

**Nguyễn Ngọc Tuấn**
Founder Transform Group — **Lark Platinum Partner**
🌐 Dự án: [larkcowork.com](https://larkcowork.com)
