# Lark IM Plugin

Kho lưu trữ này chứa cấu hình cần thiết để tích hợp Lark IM với Cursor IDE và Claude Code. Plugin cho phép các agent của bạn tương tác trực tiếp với không gian làm việc Lark IM, giúp bạn tìm kiếm tin nhắn, gửi thông báo, quản lý canvas, và nhiều hơn nữa—tất cả thông qua ngôn ngữ tự nhiên.

## Tính năng

Máy chủ MCP của Lark IM cung cấp các khả năng sau:

- **Tìm kiếm**: Tìm tin nhắn, tệp, người dùng và kênh (cả công khai lẫn riêng tư)
- **Nhắn tin**: Gửi tin nhắn, lấy lịch sử kênh và truy cập các hội thoại dạng luồng (thread)
- **Canvas**: Tạo và chia sẻ tài liệu được định dạng, xuất nội dung dưới dạng markdown
- **Quản lý người dùng**: Lấy hồ sơ người dùng bao gồm trường tùy chỉnh và thông tin trạng thái

## Yêu cầu trước khi cài đặt

Trước khi thiết lập máy chủ MCP của Lark IM, hãy đảm bảo bạn có:

- Cursor IDE hoặc Claude Code CLI đã được cài đặt
- Quyền truy cập vào không gian làm việc Lark IM có tích hợp MCP đã được quản trị viên không gian làm việc của bạn phê duyệt

## Cài đặt

Chọn phương pháp cài đặt phù hợp với IDE của bạn:

### Claude Code

Nếu bạn đang dùng Claude Code CLI, bạn có thể cài đặt plugin này bằng cách clone về máy:

```bash
git clone https://github.com/slackapi/slack-mcp-plugin.git
cd slack-mcp-plugin
claude --plugin-dir ./
```

Máy chủ MCP của Lark IM sẽ được cấu hình tự động khi plugin được nạp. Bạn sẽ được nhắc xác thực vào không gian làm việc Lark IM thông qua OAuth.

Plugin của Claude sử dụng cấu hình MCP sau (`.mcp.json`):

```json
{
  "mcpServers": {
    "slack": {
      "type": "http",
      "url": "https://mcp.slack.com/mcp",
      "oauth": {
        "clientId": "1601185624273.8899143856786",
        "callbackPort": 3118
      }
    }
  }
}
```

### Cursor

Bạn có thể dùng nút Add to Cursor sau đây hoặc làm theo các bước bên dưới để cấu hình thủ công máy chủ MCP của Lark IM trong Cursor:

[![Install MCP Server](https://cursor.com/deeplink/mcp-install-dark.svg)](https://cursor.com/en-US/install-mcp?name=slack&config=eyJ1cmwiOiJodHRwczovL21jcC5zbGFjay5jb20vbWNwIiwiYXV0aCI6eyJDTElFTlRfSUQiOiIzNjYwNzUzMTkyNjI2Ljg5MDM0NjkyMjg5ODIifX0%3D)

#### Bước 1: Mở Cursor Settings

Đi tới **Cursor → Settings → Cursor Settings** (hoặc dùng phím tắt `Cmd+,` trên macOS, `Ctrl+,` trên Windows/Linux).

#### Bước 2: Chuyển đến tab MCP

Trong giao diện Settings, nhấp vào tab **MCP** để truy cập các cấu hình máy chủ MCP.

#### Bước 3: Thêm cấu hình MCP của Lark IM

Thêm cấu hình sau để kết nối đến máy chủ MCP từ xa của Lark IM:

```json
{
  "mcpServers": {
    "slack": {
      "url": "https://mcp.slack.com/mcp",
      "auth": {
        "CLIENT_ID": "3660753192626.8903469228982"
      }
    }
  }
}
```

Lưu cấu hình. Bạn cũng sẽ thấy một nút kết nối sau khi thêm. Nhấp vào nút đó để xác thực vào Không gian làm việc Lark IM của bạn.

## Ví dụ sử dụng

Sau khi cấu hình xong, bạn có thể tương tác với Lark IM thông qua trợ lý AI bằng ngôn ngữ tự nhiên:

- **Tìm kiếm tin nhắn**: "Tìm các tin nhắn về việc ra mắt sản phẩm trong tuần qua"
- **Gửi tin nhắn**: "Gửi tin nhắn vào kênh #general báo rằng việc triển khai đã hoàn tất"
- **Tìm người dùng**: "Người dùng có email john@example.com là ai?"
- **Truy cập luồng (thread)**: "Hiển thị cho tôi luồng hội thoại từ tin nhắn đó"
- **Tạo canvas**: "Tạo một tài liệu canvas với ghi chú cuộc họp của chúng ta"

## Tài liệu & Tài nguyên

- [Tài liệu chính thức về Máy chủ MCP của Lark IM](https://docs.slack.dev/ai/mcp-server/)

## Ghi chú & Giới hạn

- **Chỉ máy chủ từ xa**: Cấu hình này kết nối đến máy chủ MCP được lưu trữ (hosted) của Lark IM. Không cần và không hỗ trợ cài đặt cục bộ.
- **Yêu cầu phê duyệt của quản trị viên**: Quản trị viên không gian làm việc Lark IM của bạn phải phê duyệt tích hợp MCP trước khi bạn có thể sử dụng tính năng này.

## Câu hỏi hoặc Vấn đề?

Nếu có câu hỏi về máy chủ MCP của Lark IM hoặc gặp sự cố tích hợp, vui lòng tham khảo [tài liệu chính thức của Lark IM](https://docs.slack.dev/ai/mcp-server/) hoặc liên hệ quản trị viên không gian làm việc của bạn.

---

## Tác giả

**Nguyễn Ngọc Tuấn**
Founder Transform Group — **Lark Platinum Partner**
🌐 Dự án: [larkcowork.com](https://larkcowork.com)
