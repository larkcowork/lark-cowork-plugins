# Plugin Hỗ trợ Khách hàng (Customer Support)

Plugin hỗ trợ khách hàng được thiết kế chủ yếu cho [Cowork](https://claude.com/product/cowork), ứng dụng desktop dạng agent của Anthropic — nhưng cũng hoạt động được trong Claude Code. Plugin cung cấp khả năng phân loại (triage) ticket, quản lý leo thang (escalation), soạn thảo phản hồi, nghiên cứu khách hàng và biên soạn bài viết cho cơ sở tri thức (knowledge base) dành cho các đội hỗ trợ.

## Cài đặt

```
claude plugins add knowledge-work-plugins/customer-support
```

## Tính năng

Plugin này biến Claude thành trợ lý song hành (co-pilot) cho công việc hỗ trợ khách hàng. Nó giúp bạn:

- **Phân loại ticket đến** với việc phân nhóm có cấu trúc, đánh giá mức độ ưu tiên và đề xuất định tuyến (routing)
- **Nghiên cứu câu hỏi của khách hàng** bằng cách tổng hợp thông tin từ nhiều nguồn, kèm theo điểm tin cậy (confidence scoring)
- **Soạn thảo phản hồi chuyên nghiệp** phù hợp với từng tình huống, mức độ khẩn cấp và kênh giao tiếp
- **Đóng gói các yêu cầu leo thang** với đầy đủ ngữ cảnh, các bước tái hiện lỗi và tác động kinh doanh để chuyển cho đội kỹ thuật hoặc sản phẩm
- **Viết bài cho cơ sở tri thức (KB)** từ các vấn đề đã được giải quyết nhằm giảm số lượng ticket trong tương lai

## Lệnh

| Lệnh | Mô tả |
|---|---|
| `/triage` | Phân nhóm, xếp ưu tiên và định tuyến một ticket hỗ trợ hoặc vấn đề của khách hàng |
| `/research` | Nghiên cứu đa nguồn về một câu hỏi hoặc chủ đề của khách hàng |
| `/draft-response` | Soạn bản nháp phản hồi gửi khách hàng cho bất kỳ tình huống nào |
| `/escalate` | Đóng gói một yêu cầu leo thang để chuyển cho đội kỹ thuật, sản phẩm hoặc ban lãnh đạo |
| `/kb-article` | Soạn bản nháp bài viết cơ sở tri thức từ một vấn đề đã được giải quyết |

## Kỹ năng

| Kỹ năng | Mô tả |
|---|---|
| `ticket-triage` | Hệ thống phân loại danh mục, khung ưu tiên (P1-P4), quy tắc định tuyến, phát hiện trùng lặp |
| `customer-research` | Phương pháp nghiên cứu đa nguồn, xếp ưu tiên nguồn, tổng hợp câu trả lời |
| `response-drafting` | Thực hành tốt nhất trong giao tiếp, hướng dẫn giọng điệu, mẫu cho các tình huống phổ biến |
| `escalation` | Các cấp độ leo thang, định dạng leo thang có cấu trúc, đánh giá tác động, nhịp theo dõi (follow-up) |
| `knowledge-management` | Tiêu chuẩn cấu trúc bài viết, cách viết để dễ tìm kiếm, nhịp rà soát, bảo trì |

## Nguồn dữ liệu

> Nếu bạn thấy các placeholder lạ hoặc cần kiểm tra công cụ nào đang được kết nối, xem [CONNECTORS.md](CONNECTORS.md).

Hãy kết nối nền tảng hỗ trợ, cơ sở tri thức và các công cụ giao tiếp của bạn để có trải nghiệm tốt nhất. Nếu không có chúng, bạn cần cung cấp ngữ cảnh khách hàng theo cách thủ công.

**Các kết nối MCP đi kèm:**
- Chat (Lark IM) cho các thảo luận nội bộ và ngữ cảnh kênh khách hàng
- Nền tảng hỗ trợ (Intercom) cho lịch sử ticket và các cuộc hội thoại với khách hàng
- CRM (HubSpot) cho chi tiết tài khoản và thông tin liên hệ
- Cơ sở tri thức (Lark Wiki, Lark Wiki) cho tài liệu nội bộ và runbook
- Công cụ theo dõi dự án (Lark) cho báo cáo lỗi và yêu cầu tính năng
- Email và lưu trữ đám mây (Lark) cho thư từ với khách hàng

**Tùy chọn bổ sung:**
- Xem [CONNECTORS.md](CONNECTORS.md) để biết các công cụ thay thế trong từng danh mục

## Quy trình mẫu

### Phân loại một ticket đến

```
Bạn: /triage Khách hàng nói dashboard của họ hiển thị trang trắng
     từ sáng nay — họ đang dùng gói Enterprise

Claude: Đang phân tích vấn đề...

## Triage: Dashboard trang trắng — khách hàng Enterprise

**Danh mục:** Bug (Lỗi)
**Mức ưu tiên:** P2 — Cao (tính năng cốt lõi bị hỏng, khách hàng enterprise)
**Khu vực sản phẩm:** Dashboard / Frontend

### Đề xuất định tuyến
Chuyển đến: Hỗ trợ Tier 2 (điều tra chi tiết trình duyệt/tài khoản,
sau đó leo thang lên Kỹ thuật nếu tái hiện được lỗi)

### Phản hồi ban đầu gợi ý
[Bản nháp ghi nhận vấn đề, hỏi các câu hỏi làm rõ
về trình duyệt và mọi thay đổi gần đây, gửi kèm link trang trạng thái]
```

### Nghiên cứu một câu hỏi của khách hàng

```
Bạn: /research Nền tảng của chúng ta có hỗ trợ SSO với Okta không?

Claude: Để tôi kiểm tra qua các nguồn hiện có...

[Tìm kiếm trong tài liệu, cơ sở tri thức, các cuộc hội thoại trước đây,
 và các nguồn trên web — tổng hợp câu trả lời kèm theo mức độ
 tin cậy và trích dẫn nguồn]
```

### Đóng gói một yêu cầu leo thang

```
Bạn: /escalate API trả về lỗi 500 không liên tục — 3 khách hàng
     Enterprise đã báo cáo trong tuần này

Claude: Tôi sẽ thu thập ngữ cảnh và dựng bản tóm tắt leo thang...

[Lấy lịch sử ticket, xác định các tài khoản bị ảnh hưởng,
 cấu trúc các bước tái hiện lỗi, đánh giá tác động kinh doanh,
 và tạo một yêu cầu leo thang có cấu trúc cho đội kỹ thuật]
```

### Soạn thảo một phản hồi

```
Bạn: /draft-response Khách hàng leo thang — tích hợp của họ
     đã ngừng hoạt động 2 ngày và họ đang bực bội

Claude: [Tạo phản hồi đồng cảm, chuyên nghiệp, ghi nhận
 sự bực bội, nhận trách nhiệm, đưa ra kế hoạch hành động
 rõ ràng kèm mốc thời gian, và đề nghị một cuộc gọi]
```

### Viết một bài KB

```
Bạn: /kb-article Cách cấu hình thông báo webhook —
     vừa xử lý xong cho khách hàng thứ ba trong tháng này

Claude: [Tạo một bài hướng dẫn có cấu trúc với các điều kiện tiên quyết,
 hướng dẫn từng bước, các bước xác minh và những
 vấn đề thường gặp — tối ưu cho tìm kiếm]
```

## Cấu hình

Plugin hoạt động ngay từ đầu (out of the box) với các kết nối MCP đi kèm. Để có trải nghiệm phong phú nhất, hãy kết nối thêm nguồn dữ liệu thông qua phần cài đặt Claude của bạn:

1. **Nền tảng hỗ trợ**: Thêm hệ thống ticket của bạn để có lịch sử ticket và ngữ cảnh khách hàng
2. **Cơ sở tri thức**: Thêm wiki của bạn để có tài liệu nội bộ và các bài KB hiện có
3. **Công cụ theo dõi dự án**: Thêm công cụ theo dõi vấn đề (issue tracker) của bạn để có báo cáo lỗi và yêu cầu tính năng
4. **CRM**: Thêm CRM của bạn để có chi tiết tài khoản và thông tin liên hệ

Nếu không có những kết nối này, plugin sẽ yêu cầu bạn cung cấp ngữ cảnh theo cách thủ công và cung cấp các khung làm việc cùng mẫu để bạn điền dữ liệu của riêng mình.

---

## Tác giả

**Nguyễn Ngọc Tuấn**
Founder Transform Group — **Lark Platinum Partner**
🌐 Dự án: [larkcowork.com](https://larkcowork.com)
