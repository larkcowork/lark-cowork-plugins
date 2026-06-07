# Tìm kiếm Doanh nghiệp (Enterprise Search)

Plugin tìm kiếm doanh nghiệp được thiết kế chủ yếu cho [Cowork](https://claude.com/product/cowork), ứng dụng desktop dạng agent của Anthropic — nhưng cũng hoạt động được trong Claude Code. Tìm kiếm xuyên suốt mọi công cụ của công ty bạn tại một nơi duy nhất — email, chat, tài liệu và wiki — mà không cần chuyển qua lại giữa các ứng dụng.

---

## Cách hoạt động

Một truy vấn tìm kiếm đồng thời trên tất cả các công cụ đã kết nối của bạn. Claude phân rã câu hỏi của bạn, chạy các tìm kiếm có chủ đích trên mọi nguồn, rồi tổng hợp kết quả thành một câu trả lời mạch lạc duy nhất kèm theo trích dẫn nguồn.

```
Bạn: "Chúng ta đã quyết định gì về việc thiết kế lại API?"
              ↓ Claude tìm kiếm
~~chat: thread #engineering hôm thứ Ba có quyết định
~~email: Email theo dõi từ Sarah kèm bản đặc tả
~~cloud storage: Tài liệu thiết kế API đã cập nhật (sửa hôm qua)
              ↓ Claude tổng hợp
"Đội đã quyết định vào thứ Ba chọn REST thay vì GraphQL.
 Sarah gửi bản đặc tả cập nhật vào thứ Năm. Tài liệu thiết kế
 phản ánh phương án cuối cùng."
```

Không cần chuyển tab. Không cần nhớ công cụ nào chứa gì. Đặt câu hỏi, nhận câu trả lời.

---

## Phạm vi tìm kiếm

> Nếu bạn thấy các placeholder lạ hoặc cần kiểm tra công cụ nào đang được kết nối, xem [CONNECTORS.md](CONNECTORS.md).

Kết nối bất kỳ tổ hợp nguồn nào. Càng kết nối nhiều, câu trả lời của bạn càng đầy đủ.

| Nguồn | Tìm được gì |
|--------|---------------|
| **~~chat** | Tin nhắn, thread, kênh, tin nhắn riêng (DM) |
| **~~email** | Email, tệp đính kèm, các cuộc hội thoại |
| **~~cloud storage** | Doc, sheet, slide, PDF |
| **Wiki / Cơ sở tri thức** | Tài liệu nội bộ, runbook |
| **Quản lý dự án** | Task, vấn đề (issue), epic, cột mốc (milestone) |
| **CRM** | Tài khoản, liên hệ, cơ hội (opportunity) |
| **Ticketing** | Ticket hỗ trợ, vấn đề của khách hàng |

Mỗi nguồn là một kết nối MCP. Thêm nhiều nguồn hơn trong phần cài đặt MCP của bạn để mở rộng phạm vi mà Claude có thể tìm kiếm.

---

## Lệnh

| Lệnh | Tác dụng |
|---------|--------------|
| `/search` | Tìm kiếm xuyên suốt mọi nguồn đã kết nối trong một truy vấn |
| `/digest` | Tạo bản tóm tắt (digest) hàng ngày hoặc hàng tuần về hoạt động trên tất cả các nguồn |

### Search

```
/enterprise-search:search what's the status of Project Aurora?
/enterprise-search:search from:sarah about:budget after:2025-01-01
/enterprise-search:search decisions made in #product this week
```

Hỗ trợ các bộ lọc: `from:`, `in:`, `after:`, `before:`, `type:` — được áp dụng một cách thông minh theo cú pháp truy vấn gốc của từng nguồn.

### Digest

```
/enterprise-search:digest --daily      # Những gì đã diễn ra hôm nay trên tất cả các nguồn
/enterprise-search:digest --weekly     # Bản tổng hợp tuần nhóm theo dự án/chủ đề
```

Làm nổi bật các hạng mục cần hành động, các quyết định và những lần bạn được nhắc đến. Nhóm hoạt động theo chủ đề để bạn lướt nhanh phần quan trọng.

---

## Kỹ năng

Ba kỹ năng tạo nên trải nghiệm tìm kiếm:

**Search Strategy** — Phân rã truy vấn và chuyển ngữ theo từng nguồn. Tách câu hỏi ngôn ngữ tự nhiên của bạn thành các tìm kiếm có chủ đích cho mỗi nguồn, xử lý sự mơ hồ và linh hoạt dự phòng khi nguồn không khả dụng.

**Source Management** — Biết những nguồn MCP nào đang khả dụng, hướng dẫn bạn kết nối nguồn mới, quản lý độ ưu tiên nguồn và xử lý giới hạn tốc độ (rate limit).

**Knowledge Synthesis** — Kết hợp kết quả từ nhiều nguồn thành câu trả lời mạch lạc. Loại bỏ thông tin trùng lặp giữa các nguồn, trích dẫn nguồn, chấm điểm tin cậy dựa trên độ mới và độ uy tín, và tóm tắt các tập kết quả lớn.

---

## Quy trình mẫu

### Tìm một quyết định

```
Bạn: /enterprise-search:search khi nào chúng ta quyết định chuyển sang Postgres?

Claude tìm kiếm:
  ~~chat → #engineering, #infrastructure với "postgres" "switch" "decision"
  ~~email → các thread có "postgres" trong tiêu đề
  ~~cloud storage → các tài liệu nhắc đến việc di chuyển cơ sở dữ liệu

Kết quả: "Quyết định được đưa ra ngày 3 tháng 3 trong #infrastructure (link).
         Email của Sarah ngày 4 tháng 3 xác nhận mốc thời gian.
         Tài liệu kế hoạch di chuyển được cập nhật ngày 5 tháng 3."
```

### Bắt kịp sau khi nghỉ phép

```
Bạn: /enterprise-search:digest --weekly

Claude quét:
  ~~chat → các kênh bạn tham gia, DM, các lượt nhắc đến bạn
  ~~email → hoạt động hộp thư
  ~~cloud storage → các tài liệu được chia sẻ với bạn hoặc đã chỉnh sửa

Kết quả: Bản tóm tắt nhóm theo dự án với các hạng mục cần hành động
        được đánh dấu và các quyết định được làm nổi bật.
```

### Tìm một chuyên gia

```
Bạn: /enterprise-search:search ai am hiểu về cấu hình Kubernetes của chúng ta?

Claude tìm kiếm:
  ~~chat → tin nhắn về Kubernetes, k8s, cluster
  ~~cloud storage → các tài liệu được viết về hạ tầng
  Wiki → runbook và tài liệu kiến trúc

Kết quả: "Dựa trên lịch sử tin nhắn và quyền tác giả tài liệu,
         Alex và Priya là những người bạn nên hỏi về k8s.
         Đây là runbook chính (link)."
```

---

## Bắt đầu

```bash
# 1. Cài đặt
claude plugins add knowledge-work-plugins/enterprise-search

# 2. Tìm kiếm xuyên suốt mọi thứ
/enterprise-search:search [câu hỏi của bạn ở đây]

# 3. Nhận bản tóm tắt
/enterprise-search:digest --daily
```

Càng kết nối nhiều nguồn qua MCP, kết quả tìm kiếm của bạn càng đầy đủ. Bắt đầu với ~~chat, ~~email và ~~cloud storage, sau đó thêm wiki, công cụ quản lý dự án và CRM khi cần.

---

## Triết lý

Người làm tri thức (knowledge worker) tiêu tốn hàng giờ mỗi tuần để truy lùng thông tin rải rác khắp các công cụ. Câu trả lời tồn tại ở đâu đó — trong một thread Lark IM, một chuỗi email, một tài liệu, một trang wiki — nhưng để tìm được nó thì phải tìm trong từng công cụ riêng lẻ, đối chiếu kết quả, và hy vọng đã kiểm tra đúng chỗ.

Enterprise Search xem tất cả công cụ của bạn như một cơ sở tri thức duy nhất có thể tìm kiếm được. Một truy vấn, mọi nguồn, kết quả đã được tổng hợp. Tri thức của công ty bạn không nên bị khóa trong các silo riêng lẻ. Hãy tìm kiếm mọi thứ cùng một lúc.

---

## Tác giả

**Nguyễn Ngọc Tuấn**
Founder Transform Group — **Lark Platinum Partner**
🌐 Dự án: [larkcowork.com](https://larkcowork.com)
