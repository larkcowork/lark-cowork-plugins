# Plugin Kỹ thuật (Engineering)

Một plugin kỹ thuật phần mềm được thiết kế chủ yếu cho [Cowork](https://claude.com/product/cowork), ứng dụng desktop dạng agent của Anthropic — tuy nhiên cũng hoạt động được trong Claude Code. Hỗ trợ họp standup, review code, ra quyết định kiến trúc, ứng phó sự cố, gỡ lỗi và viết tài liệu kỹ thuật. Hoạt động với mọi đội ngũ kỹ thuật — dùng độc lập với thông tin bạn cung cấp, và mạnh mẽ hơn nhiều khi bạn kết nối các công cụ quản lý mã nguồn, theo dõi dự án và giám sát hệ thống.

## Cài đặt

```bash
claude plugins add knowledge-work-plugins/engineering
```

## Lệnh

Các luồng công việc tường minh mà bạn gọi bằng lệnh slash:

| Lệnh | Mô tả |
|---|---|
| `/standup` | Tạo bản cập nhật standup từ hoạt động gần đây của bạn — commit, PR, ticket và chat |
| `/review` | Review các thay đổi code — bảo mật, hiệu năng, phong cách và tính đúng đắn |
| `/debug` | Phiên gỡ lỗi có cấu trúc — tái hiện, khoanh vùng, chẩn đoán và sửa |
| `/architecture` | Tạo hoặc đánh giá các quyết định kiến trúc — định dạng ADR kèm phân tích đánh đổi |
| `/incident` | Chạy luồng ứng phó sự cố — triage, truyền thông, giảm thiểu và viết postmortem |
| `/deploy-checklist` | Danh sách kiểm tra trước khi triển khai — xác minh test, review thay đổi, kiểm tra phụ thuộc, xác nhận kế hoạch rollback |

Tất cả các lệnh đều hoạt động **độc lập** (dán code, mô tả hệ thống, tải lên file) và trở nên **mạnh mẽ hơn** khi có các kết nối MCP.

## Kỹ năng

Kiến thức chuyên môn mà Claude tự động sử dụng khi phù hợp:

| Kỹ năng | Mô tả |
|---|---|
| `code-review` | Review code để tìm lỗi, vấn đề bảo mật, hiệu năng và khả năng bảo trì |
| `incident-response` | Triage và quản lý sự cố production — cập nhật trạng thái, runbook, postmortem |
| `system-design` | Thiết kế hệ thống và dịch vụ — sơ đồ kiến trúc, thiết kế API, mô hình hóa dữ liệu |
| `tech-debt` | Xác định, phân loại và ưu tiên nợ kỹ thuật — xây dựng kế hoạch khắc phục |
| `testing-strategy` | Thiết kế chiến lược kiểm thử — độ phủ unit, integration, e2e, lập kế hoạch test |
| `documentation` | Viết và duy trì tài liệu kỹ thuật — README, tài liệu API, runbook, hướng dẫn onboarding |

## Quy trình mẫu

### Standup buổi sáng

```
/standup
```

Nếu các công cụ của bạn đã được kết nối, tôi sẽ kéo về các commit, hoạt động PR và cập nhật ticket gần đây của bạn. Nếu chưa, hãy cho tôi biết bạn đã làm những gì và tôi sẽ định dạng lại giúp bạn.

### Review Code

```
/review https://github.com/org/repo/pull/123
```

Chia sẻ một link PR, dán một diff, hoặc trỏ tới các file. Bạn sẽ nhận được một bản review có cấu trúc bao quát bảo mật, hiệu năng, tính đúng đắn và phong cách.

### Gỡ lỗi một sự cố

```
/debug Users are getting 500 errors on the checkout page
```

Cùng đi qua một quy trình gỡ lỗi có cấu trúc: tái hiện, khoanh vùng, chẩn đoán, sửa. Tôi sẽ giúp bạn suy nghĩ về nó một cách có hệ thống.

### Quyết định Kiến trúc

```
/architecture Should we use a message queue or direct API calls between services?
```

Nhận một ADR có cấu trúc kèm phân tích các phương án, các đánh đổi và một khuyến nghị.

### Ứng phó Sự cố

```
/incident The payments service is returning 503s
```

Khởi động một luồng ứng phó sự cố: triage mức độ nghiêm trọng, soạn thảo truyền thông, theo dõi dòng thời gian, và tạo postmortem khi đã được xử lý xong.

### Kiểm tra Trước khi Triển khai

```
/deploy-checklist auth-service v2.3.0
```

Nhận một danh sách kiểm tra triển khai được tùy chỉnh dựa trên dịch vụ của bạn và những gì đang thay đổi.

## Độc lập + Mạnh mẽ hơn

Mọi lệnh và kỹ năng đều hoạt động mà không cần bất kỳ tích hợp nào:

| Bạn có thể làm gì | Độc lập | Mạnh mẽ hơn với |
|-----------------|------------|-------------------|
| Cập nhật standup | Mô tả công việc của bạn | Quản lý mã nguồn, Theo dõi dự án, Chat |
| Review code | Dán diff hoặc code | Quản lý mã nguồn (tự động kéo về PR) |
| Phiên gỡ lỗi | Mô tả vấn đề | Giám sát (kéo về log và metric) |
| Quyết định kiến trúc | Mô tả hệ thống | Cơ sở tri thức (tìm các ADR trước đó) |
| Ứng phó sự cố | Mô tả sự cố | Giám sát, Quản lý sự cố, Chat |
| Checklist triển khai | Mô tả lần triển khai | CI/CD, Quản lý mã nguồn |

## Tích hợp MCP

> Nếu bạn thấy các placeholder lạ hoặc cần kiểm tra xem những công cụ nào đang được kết nối, hãy xem [CONNECTORS.md](CONNECTORS.md).

Kết nối các công cụ của bạn để có trải nghiệm phong phú hơn:

| Danh mục | Ví dụ | Nó cho phép làm gì |
|---|---|---|
| **Quản lý mã nguồn** | GitHub, GitLab | Diff của PR, lịch sử commit, trạng thái nhánh |
| **Theo dõi dự án** | Lark Task, Lark Task, Lark Task | Trạng thái ticket, dữ liệu sprint, phân công |
| **Giám sát** | Datadog, New Relic | Log, metric, cảnh báo, dashboard |
| **Quản lý sự cố** | PagerDuty, Opsgenie | Lịch trực, theo dõi sự cố, gọi báo (paging) |
| **Chat** | Lark IM, Teams | Thảo luận nhóm, kênh standup |
| **Cơ sở tri thức** | Lark Wiki, Lark Wiki | ADR, runbook, tài liệu onboarding |

Xem [CONNECTORS.md](CONNECTORS.md) để biết danh sách đầy đủ các tích hợp được hỗ trợ.

## Thiết lập

Tạo một file thiết lập cục bộ tại `engineering/.claude/settings.local.json` để cá nhân hóa:

```json
{
  "name": "Your Name",
  "title": "Software Engineer",
  "team": "Your Team",
  "company": "Your Company",
  "techStack": ["Python", "TypeScript", "PostgreSQL", "AWS"],
  "defaultBranch": "main",
  "deployProcess": "canary"
}
```

Plugin sẽ hỏi bạn thông tin này một cách tương tác nếu nó chưa được cấu hình.

---

## Tác giả

**Nguyễn Ngọc Tuấn**
Founder Transform Group — **Lark Platinum Partner**
🌐 Dự án: [larkcowork.com](https://larkcowork.com)
