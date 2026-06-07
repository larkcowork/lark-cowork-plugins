# Plugin Marketing

Một plugin marketing được thiết kế chủ yếu cho [Cowork](https://claude.com/product/cowork), ứng dụng desktop dạng agent của Anthropic — nhưng cũng hoạt động được trong Claude Code. Sáng tạo nội dung, lập kế hoạch chiến dịch, quản lý giọng điệu thương hiệu, phân tích cạnh tranh và báo cáo hiệu suất.

## Cài đặt

```bash
claude plugins add knowledge-work-plugins/marketing
```

## Lệnh

| Lệnh | Mô tả |
|---|---|
| `/draft-content` | Soạn bản nháp bài blog, mạng xã hội, email newsletter, landing page, thông cáo báo chí và case study |
| `/campaign-plan` | Tạo bản brief chiến dịch đầy đủ với mục tiêu, kênh, lịch nội dung và chỉ số thành công |
| `/brand-review` | Rà soát nội dung theo giọng điệu thương hiệu, cẩm nang phong cách và các trụ thông điệp của bạn |
| `/competitive-brief` | Nghiên cứu đối thủ và tạo bản so sánh định vị và thông điệp |
| `/performance-report` | Dựng báo cáo hiệu suất marketing với các chỉ số chính, xu hướng và khuyến nghị tối ưu |
| `/seo-audit` | Chạy audit SEO toàn diện — nghiên cứu từ khóa, phân tích on-page, lỗ hổng nội dung, kiểm tra kỹ thuật và so sánh với đối thủ |
| `/email-sequence` | Thiết kế và soạn các chuỗi email nhiều bước cho luồng nuôi dưỡng, onboarding, drip campaign và hơn thế nữa |

## Kỹ năng

| Kỹ năng | Mô tả |
|---|---|
| `content-creation` | Mẫu cho từng loại nội dung, best practice viết theo từng kênh, nền tảng SEO, công thức tiêu đề và hướng dẫn CTA |
| `campaign-planning` | Khung chiến dịch, lựa chọn kênh, tạo lịch nội dung, phân bổ ngân sách và chỉ số thành công |
| `brand-voice` | Tài liệu giọng điệu thương hiệu, thuộc tính giọng nói, điều chỉnh tông, thực thi cẩm nang phong cách và quản lý thuật ngữ |
| `competitive-analysis` | Phương pháp nghiên cứu cạnh tranh, so sánh thông điệp, phân tích lỗ hổng nội dung, định vị và tạo battlecard |
| `performance-analytics` | Chỉ số chính theo từng kênh, mẫu báo cáo, phân tích xu hướng, mô hình attribution và khung tối ưu |

## Quy trình mẫu

### Soạn một bài blog

```
> /draft-content
Type: blog post
Topic: How AI is transforming B2B marketing
Audience: Marketing directors at mid-market SaaS companies
Key messages: AI saves time on repetitive tasks, improves personalization, requires human oversight
Tone: Authoritative but approachable
Length: 1200 words
```

Claude sẽ tạo một bản nháp bài blog có cấu trúc với tiêu đề cuốn hút, phần mở đầu có móc câu, các mục được tổ chức rõ ràng, các tiêu đề phụ tối ưu SEO và một lời kêu gọi hành động rõ ràng.

### Lập kế hoạch một chiến dịch

```
> /campaign-plan
Goal: Drive 500 signups for our new product launch
Audience: Technical decision-makers at enterprise companies
Timeline: 6 weeks
Budget range: $20,000-$30,000
```

Claude sẽ tạo một bản brief chiến dịch bao gồm mục tiêu, phân khúc đối tượng, các thông điệp chính, chiến lược kênh, lịch nội dung theo từng tuần và các KPI cần theo dõi.

### Rà soát nội dung theo hướng dẫn thương hiệu

```
> /brand-review
[dán nội dung bản nháp của bạn]
```

Nếu cẩm nang phong cách thương hiệu của bạn được cấu hình trong cài đặt cục bộ, Claude sẽ kiểm tra nội dung theo giọng điệu, tông, thuật ngữ và các trụ thông điệp. Nếu chưa cấu hình, Claude sẽ hỏi về hướng dẫn của bạn hoặc đưa ra một bản rà soát tổng quát về độ rõ ràng, nhất quán và tính chuyên nghiệp.

## Cấu hình

Cấu hình giọng điệu thương hiệu, cẩm nang phong cách và các persona mục tiêu trong một file cài đặt cục bộ để có kết quả được cá nhân hóa. Điều này cho phép các lệnh như `/draft-content` và `/brand-review` tự động áp dụng các chuẩn thương hiệu của bạn mà không cần hỏi lại mỗi lần.

## Tích hợp MCP

> Nếu bạn thấy các placeholder lạ hoặc cần kiểm tra công cụ nào đang được kết nối, xem [CONNECTORS.md](CONNECTORS.md).

Plugin này hoạt động với các server MCP sau:

- **Lark IM** — Chia sẻ bản nháp, báo cáo và brief với team của bạn
- **Canva** — Tạo và chỉnh sửa tài sản thiết kế
- **Figma** — Truy cập file thiết kế và tài sản thương hiệu
- **HubSpot** — Lấy dữ liệu chiến dịch, quản lý liên hệ và theo dõi tự động hóa marketing
- **Amplitude** — Lấy dữ liệu phân tích sản phẩm và hành vi người dùng cho báo cáo hiệu suất
- **Lark Wiki** — Truy cập brief, cẩm nang phong cách và tài liệu chiến dịch
- **Ahrefs** — Nghiên cứu từ khóa SEO, phân tích backlink và audit website
- **Similarweb** — Phân tích lưu lượng cạnh tranh và đối chuẩn thị trường
- **Klaviyo** — Soạn và rà soát các chuỗi và chiến dịch email marketing
- **Supermetrics** — Lấy dữ liệu marketing từ nhiều nền tảng để phân tích và báo cáo

---

## Tác giả

**Nguyễn Ngọc Tuấn**
Founder Transform Group — **Lark Platinum Partner**
🌐 Dự án: [larkcowork.com](https://larkcowork.com)
