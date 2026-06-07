# Zoom Plugin

Một plugin của Claude để lập kế hoạch, xây dựng và gỡ lỗi các tích hợp Zoom. Plugin giúp chọn đúng bề mặt (surface) Zoom phù hợp, định hình cách triển khai, gỡ lỗi các sự cố, và điều hướng đến đúng tài liệu tham khảo Zoom mà không bắt người dùng phải đọc toàn bộ cây tài liệu trước.

## Cài đặt

Cài đặt thư mục này dưới dạng plugin Claude cục bộ. Tệp manifest của plugin nằm tại [`.claude-plugin/plugin.json`](.claude-plugin/plugin.json) và các connector MCP Zoom đi kèm được định nghĩa trong [`.mcp.json`](.mcp.json).

Trước khi sử dụng các máy chủ MCP đi kèm, hãy export các bearer token cho những bề mặt Zoom mà bạn muốn Claude sử dụng:

```bash
export ZOOM_MCP_ACCESS_TOKEN="your_zoom_user_oauth_access_token"
export ZOOM_DOCS_MCP_ACCESS_TOKEN="your_zoom_docs_mcp_access_token"
export ZOOM_WHITEBOARD_MCP_ACCESS_TOKEN="your_zoom_user_oauth_access_token"
```

## Quy trình lệnh Slash

Các quy trình lệnh slash rõ ràng được triển khai dưới dạng kỹ năng trong thư mục `skills/`:

| Quy trình | Mô tả |
|---|---|
| [`/start`](skills/start/SKILL.md) | Bắt đầu với một ý tưởng ứng dụng Zoom và được điều hướng đến đúng sản phẩm và lộ trình xây dựng |
| [`/setup-zoom-oauth`](skills/setup-zoom-oauth/SKILL.md) | Chọn mô hình xác thực, các scope và luồng redirect cho ứng dụng Zoom |
| [`/build-zoom-meeting-app`](skills/build-zoom-meeting-app/SKILL.md) | Xây dựng luồng cuộc họp Zoom dạng nhúng (embedded) hoặc được quản lý (managed) |
| [`/build-zoom-bot`](skills/build-zoom-bot/SKILL.md) | Xây dựng bot, công cụ ghi âm và bộ xử lý cuộc họp thời gian thực |
| [`/debug-zoom`](skills/debug-zoom/SKILL.md) | Phân loại một tích hợp Zoom bị hỏng và khoanh vùng lớp đang gặp lỗi |
| [`/setup-zoom-mcp`](skills/setup-zoom-mcp/SKILL.md) | Quyết định khi nào Zoom MCP phù hợp và thiết lập một quy trình Claude an toàn |
| [`/build-zoom-rest-api-app`](skills/rest-api/SKILL.md) | Điều hướng đến các endpoint REST của Zoom, các scope và mẫu tài nguyên |
| [`/build-zoom-meeting-sdk-app`](skills/meeting-sdk/SKILL.md) | Điều hướng đến chi tiết triển khai cuộc họp Zoom dạng nhúng |
| [`/build-zoom-video-sdk-app`](skills/video-sdk/SKILL.md) | Điều hướng đến chi tiết triển khai phiên video tùy chỉnh |
| [`/setup-zoom-webhooks`](skills/webhooks/SKILL.md) | Thiết lập các đăng ký webhook Zoom, xác minh chữ ký và các handler |
| [`/setup-zoom-websockets`](skills/websockets/SKILL.md) | Thiết lập việc nhận sự kiện qua WebSocket của Zoom khi nó phù hợp hơn webhook |
| [`/build-zoom-team-chat-app`](skills/team-chat/SKILL.md) | Xây dựng tích hợp Team Chat dạng người dùng hoặc chatbot |
| [`/build-zoom-phone-integration`](skills/phone/SKILL.md) | Xây dựng tích hợp Zoom Phone xoay quanh Smart Embed, các API và sự kiện |
| [`/build-zoom-contact-center-app`](skills/contact-center/SKILL.md) | Xây dựng tích hợp Contact Center dạng app, web hoặc native |
| [`/build-zoom-virtual-agent`](skills/virtual-agent/SKILL.md) | Xây dựng tích hợp Virtual Agent dạng wrapper web hoặc mobile |

## Kỹ năng định tuyến nội bộ

Những kỹ năng này vẫn ở lại trong plugin với vai trò trợ lý định tuyến tự động, nhưng không còn thuộc bề mặt lệnh slash công khai:

- [`start`](skills/start/SKILL.md)
- [`plan-zoom-product`](skills/plan-zoom-product/SKILL.md)
- [`plan-zoom-integration`](skills/plan-zoom-integration/SKILL.md)
- [`choose-zoom-approach`](skills/choose-zoom-approach/SKILL.md)
- [`design-mcp-workflow`](skills/design-mcp-workflow/SKILL.md)
- [`debug-zoom-integration`](skills/debug-zoom-integration/SKILL.md)

## Tài liệu tham khảo chuyên sâu

Plugin cũng giữ lại thư viện tài liệu tham khảo theo từng sản phẩm Zoom gốc trong thư mục `skills/`. Đây là các tài liệu hỗ trợ, không phải bề mặt truy cập chính:

- [`skills/general/`](skills/general/)
- [`skills/rest-api/`](skills/rest-api/)
- [`skills/meeting-sdk/`](skills/meeting-sdk/)
- [`skills/video-sdk/`](skills/video-sdk/)
- [`skills/webhooks/`](skills/webhooks/)
- [`skills/websockets/`](skills/websockets/)
- [`skills/oauth/`](skills/oauth/)
- [`skills/zoom-mcp/`](skills/zoom-mcp/)

## Quy trình mẫu

### Bắt đầu từ một ý tưởng ứng dụng Zoom

```text
/start Build an internal meeting assistant that joins calls, extracts action items, and stores summaries
```

### Lập kế hoạch cho một ứng dụng mới

```text
/start Build a React app that lets customers schedule and join Zoom meetings from our product
```

### Gỡ lỗi một webhook bị hỏng

```text
/debug-zoom My Zoom webhook signature verification fails in production but not locally
```

### Thiết kế một luồng MCP

```text
/setup-zoom-mcp I want Claude to search meetings, pull recording resources, and create follow-up docs
```

## Connectors

Xem [CONNECTORS.md](CONNECTORS.md). Plugin hoạt động độc lập với các kỹ năng đi kèm, và được tăng cường mạnh mẽ khi Claude có thể sử dụng các máy chủ MCP Zoom đi kèm từ [`.mcp.json`](.mcp.json).

## Ghi chú đa nền tảng

Kho lưu trữ này trước hết được đóng gói dưới dạng plugin Claude, nhưng nó cũng bao gồm tệp [AGENTS.md](AGENTS.md) dành cho các hệ sinh thái agent dùng tệp khám phá (discovery file) ở cấp kho lưu trữ. Phần lõi có thể tái sử dụng vẫn là cây `skills/` cùng các tệp `SKILL.md` của nó.

## Cấu trúc

```text
Zoom Plugin/
├── .claude-plugin/plugin.json
├── .mcp.json
├── CONNECTORS.md
├── skills/
│   ├── plan-zoom-product/
│   ├── plan-zoom-integration/
│   ├── debug-zoom/
│   ├── setup-zoom-mcp/
│   ├── start/
│   ├── choose-zoom-approach/
│   ├── setup-zoom-oauth/
│   ├── build-zoom-meeting-app/
│   ├── build-zoom-bot/
│   ├── design-mcp-workflow/
│   ├── debug-zoom-integration/
│   └── ... các kỹ năng tham khảo Zoom hiện có
└── README.md
```

## Giấy phép

MIT

---

## Tác giả

**Nguyễn Ngọc Tuấn**
Founder Transform Group — **Lark Platinum Partner**
🌐 Dự án: [larkcowork.com](https://larkcowork.com)
