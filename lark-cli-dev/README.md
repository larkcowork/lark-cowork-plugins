# Plugin Công cụ Phát triển lark-cli (lark-cli Dev Tools)

Bộ công cụ chuyên biệt để xây dựng, gỡ lỗi và mở rộng cầu nối MCP tích hợp sẵn bên trong `lark-cli` (`cmd/mcp/`). Plugin này dành cho các lập trình viên đang làm việc trên chính cầu nối đó — thêm công cụ, sửa schema, và giữ cho các MCP host luôn kết nối — chứ không dành cho người dùng cuối đang chạy các quy trình Lark.

## Cài đặt

```
claude plugins add lark-cowork/lark-cli-dev
```

## Tính năng

Plugin này trang bị cho Claude mô hình tư duy và phản xạ thành thạo khi làm việc trên `cmd/mcp/`:

- **Kỹ năng `lark-cli-mcp`** — Dạy Claude cách cầu nối thực sự hoạt động: một máy chủ JSON-RPC stdio không trạng thái (`lark-cli mcp serve`) biến mỗi `tools/call` thành một lệnh `os/exec` của một shortcut `lark-cli`, mà shortcut này lần lượt gọi đến Lark/Feishu Open API. Nó mã hóa các quy tắc cứng — stdout chỉ dành cho JSON-RPC (ghi log vào stderr qua `s.logf`), không dùng thư viện MCP bên thứ ba, một công cụ MCP ≈ một shortcut (đừng kết hợp nhiều luồng), và luôn xác minh tên cờ (flag) đối chiếu với `shortcuts/<domain>/*.go` trước khi phát hành.
- **Sáu lệnh** — Một vòng phát triển đầy đủ: dựng khung một công cụ mới, build lại và cài lại binary, kiểm thử nhanh phần bắt tay (handshake), gọi một công cụ đơn lẻ từ đầu đến cuối, kiểm tra danh mục đang chạy, và chạy một báo cáo tình trạng một trang trên mọi lớp của cầu nối.

## Kỹ năng

| Kỹ năng | Mô tả |
|-------|-------------|
| `lark-cli-mcp` | Xây dựng/gỡ lỗi/mở rộng cầu nối MCP của lark-cli trong cmd/mcp/. Kích hoạt khi thêm/xóa/tinh chỉnh công cụ MCP, Claude Desktop mất kết nối, thay đổi schema công cụ. |

## Lệnh

| Lệnh | Tính năng |
|---------|--------------|
| `/mcp-add` | Hướng dẫn từng bước thêm một công cụ mới vào cầu nối MCP của lark-cli |
| `/mcp-call` | Gọi một công cụ MCP từ đầu đến cuối (bắt tay + tools/call) và in ra kết quả |
| `/mcp-doctor` | Báo cáo tình trạng một trang cho cầu nối MCP — build, bắt tay, lệnh gọi mẫu, cấu hình host, log |
| `/mcp-rebuild` | Build lại lark-cli và cài lại binary để Claude Desktop nhận thay đổi |
| `/mcp-test` | Kiểm thử nhanh máy chủ MCP của lark-cli (initialize + tools/list + tùy chọn tools/call) |
| `/mcp-tools` | Liệt kê các công cụ MCP hiện đang được `lark-cli mcp serve` cung cấp |

## Yêu cầu

Plugin này giả định bạn đã có một bản clone cục bộ của mã nguồn Go `lark-cli` — cụ thể là `cmd/mcp/` (cầu nối) và `shortcuts/` (các lệnh nền tảng) — cùng với một bộ công cụ Go (toolchain) hoạt động. Các lệnh sẽ chạy `go build` và điều khiển `lark-cli mcp serve` dựa trên mã nguồn đó, nên chúng mong đợi được vận hành từ thư mục gốc của repository với `go` nằm trong `PATH` của bạn. Nó **không** đóng gói kèm máy chủ MCP `lark` lúc chạy (runtime); đây là công cụ phát triển cho mã nguồn của cầu nối.

---

## Tác giả

**Nguyễn Ngọc Tuấn**
Founder Transform Group — **Lark Platinum Partner**
🌐 Dự án: [larkcowork.com](https://larkcowork.com)
