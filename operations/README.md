# Plugin Vận hành (Operations)

Plugin vận hành doanh nghiệp được thiết kế chủ yếu cho [Cowork](https://claude.com/product/cowork), ứng dụng desktop dạng agent của Anthropic — dù vậy nó cũng hoạt động tốt trong Claude Code. Hỗ trợ quản lý nhà cung cấp, tài liệu hoá quy trình, quản lý thay đổi, lập kế hoạch năng lực, theo dõi tuân thủ và lập kế hoạch nguồn lực. Hoạt động với mọi đội vận hành — độc lập với phần nhập liệu của bạn, mạnh mẽ vượt trội khi bạn kết nối ITSM, project tracker và các công cụ khác.

## Cài đặt

```bash
claude plugins add knowledge-work-plugins/operations
```

## Lệnh

Các quy trình tường minh mà bạn gọi bằng một lệnh slash:

| Lệnh | Mô tả |
|---|---|
| `/vendor-review` | Đánh giá một nhà cung cấp — phân tích chi phí, đánh giá rủi ro, tóm tắt hợp đồng và đề xuất gia hạn |
| `/process-doc` | Tài liệu hoá một quy trình nghiệp vụ — lưu đồ, ma trận RACI, SOP và runbook |
| `/change-request` | Tạo một yêu cầu quản lý thay đổi — phân tích tác động, kế hoạch rollback, định tuyến phê duyệt |
| `/capacity-plan` | Lập kế hoạch năng lực nguồn lực — phân tích khối lượng công việc, mô hình hoá nhân sự, dự báo mức sử dụng |
| `/status-report` | Tạo báo cáo trạng thái — cập nhật dự án, KPI, rủi ro và hạng mục hành động cho lãnh đạo |
| `/runbook` | Tạo hoặc cập nhật một runbook vận hành — quy trình từng bước cho các tác vụ lặp lại |

Tất cả các lệnh đều hoạt động **độc lập** (cung cấp ngữ cảnh và chi tiết) và được **tăng cường mạnh mẽ** với các connector MCP.

## Kỹ năng

Kiến thức nghiệp vụ mà Claude tự động sử dụng khi phù hợp:

| Kỹ năng | Mô tả |
|---|---|
| `vendor-management` | Đánh giá, so sánh và quản lý quan hệ nhà cung cấp — hợp đồng, hiệu suất, rủi ro |
| `process-optimization` | Phân tích và cải tiến quy trình nghiệp vụ — phát hiện điểm nghẽn, giảm lãng phí, tinh gọn luồng công việc |
| `change-management` | Lập kế hoạch và thực thi các thay đổi về tổ chức hoặc kỹ thuật — truyền thông, đào tạo, áp dụng |
| `risk-assessment` | Nhận diện, đánh giá và giảm thiểu rủi ro vận hành — sổ rủi ro, phân tích tác động, biện pháp kiểm soát |
| `compliance-tracking` | Theo dõi yêu cầu tuân thủ — kiểm toán, chứng nhận, hạn chót pháp lý, tuân thủ chính sách |
| `resource-planning` | Lập kế hoạch và tối ưu phân bổ nguồn lực — năng lực, mức sử dụng, dự báo, ngân sách |

## Quy trình mẫu

### Đánh giá một nhà cung cấp

```
/vendor-review
```

Cung cấp tên nhà cung cấp, chi tiết hợp đồng, hoặc tải lên một bản đề xuất. Nhận về một bản đánh giá có cấu trúc với phân tích chi phí, các cảnh báo rủi ro và một đề xuất.

### Tài liệu hoá một quy trình

```
/process-doc employee offboarding
```

Mô tả quy trình hoặc hướng dẫn tôi từng bước. Nhận về một SOP hoàn chỉnh với lưu đồ, ma trận RACI và quy trình từng bước.

### Gửi một yêu cầu thay đổi

```
/change-request
```

Mô tả thay đổi. Nhận về một bản phân tích tác động, đánh giá rủi ro, kế hoạch rollback và mẫu truyền thông sẵn sàng để phê duyệt.

### Lập kế hoạch năng lực

```
/capacity-plan
```

Tải lên dữ liệu team hoặc mô tả nguồn lực của bạn. Nhận về phân tích mức sử dụng, nhận diện điểm nghẽn và đề xuất nhân sự.

### Báo cáo trạng thái cho lãnh đạo

```
/status-report
```

Tôi sẽ lấy các cập nhật từ những công cụ bạn đã kết nối (hoặc hỏi bạn để nhập liệu) và tạo một báo cáo trạng thái chỉn chu với KPI, rủi ro và các bước tiếp theo.

### Tạo một runbook

```
/runbook monthly close process
```

Hướng dẫn tôi qua quy trình một lần. Tôi sẽ tài liệu hoá nó thành một runbook có thể lặp lại với checklist, hướng dẫn xử lý sự cố và các đường escalation.

## Độc lập + Tăng cường

Mọi lệnh và kỹ năng đều hoạt động mà không cần bất kỳ tích hợp nào:

| Việc bạn có thể làm | Độc lập | Tăng cường với |
|-----------------|------------|-------------------|
| Đánh giá nhà cung cấp | Cung cấp chi tiết, tải lên đề xuất | Mua sắm, Cơ sở tri thức |
| Tài liệu hoá quy trình | Mô tả quy trình | Cơ sở tri thức (tài liệu có sẵn) |
| Yêu cầu thay đổi | Mô tả thay đổi | ITSM, Project tracker |
| Lập kế hoạch năng lực | Tải lên dữ liệu, mô tả team | Project tracker (dữ liệu khối lượng công việc) |
| Báo cáo trạng thái | Cung cấp cập nhật thủ công | Project tracker, Chat, Calendar |
| Runbook | Đi qua quy trình | Cơ sở tri thức, ITSM |

## Tích hợp MCP

> Nếu bạn thấy các placeholder lạ hoặc cần kiểm tra công cụ nào đang được kết nối, hãy xem [CONNECTORS.md](CONNECTORS.md).

Kết nối công cụ của bạn để có trải nghiệm phong phú hơn:

| Danh mục | Ví dụ | Tính năng mở khoá |
|---|---|---|
| **ITSM** | ServiceNow, Zendesk | Quản lý ticket, yêu cầu thay đổi, theo dõi sự cố |
| **Project tracker** | Lark Task, Lark Task, Lark Base | Trạng thái dự án, phân bổ nguồn lực, theo dõi tác vụ |
| **Cơ sở tri thức** | Lark Wiki, Lark Wiki | Tài liệu quy trình, runbook, chính sách |
| **Chat** | Lark IM, Teams | Phối hợp team, phê duyệt, cập nhật trạng thái |
| **Calendar** | Lark Calendar, Lark | Lập lịch họp, theo dõi hạn chót |
| **Email** | Lark Mail, Lark | Trao đổi với nhà cung cấp, phê duyệt |

Xem [CONNECTORS.md](CONNECTORS.md) để biết danh sách đầy đủ các tích hợp được hỗ trợ.

## Cài đặt cấu hình

Tạo một file cấu hình cục bộ tại `operations/.claude/settings.local.json` để cá nhân hoá:

```json
{
  "company": "Your Company",
  "team": "Operations",
  "reportingCadence": "weekly",
  "approvalChain": ["Manager", "Director", "VP"],
  "complianceFrameworks": ["SOC 2", "ISO 27001"],
  "fiscalYearStart": "January"
}
```

Plugin sẽ hỏi bạn thông tin này một cách tương tác nếu nó chưa được cấu hình.

---

## Tác giả

**Nguyễn Ngọc Tuấn**
Founder Transform Group — **Lark Platinum Partner**
🌐 Dự án: [larkcowork.com](https://larkcowork.com)
