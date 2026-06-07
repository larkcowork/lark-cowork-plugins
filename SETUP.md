# Lark Cowork — Cài đặt & Kiểm chứng

Checklist end-to-end đưa marketplace này từ "các file trên ổ đĩa" thành "chạy được trong Claude".

## 0. Hiện trạng máy này

Binary `lark-cli`, phần xác thực Keychain của nó, và mục MCP trong Claude Desktop đã bị **xóa sạch**
trong một lần dọn dẹp demo trước đây. Vì vậy bài test runtime đầu tiên cần cài lại lark-cli. Bản thân
các file plugin đã hoàn chỉnh và được kiểm chứng (JSON + ánh xạ connector); chỉ thiếu binary để chạy
thật.

## 1. Cài lại + xác thực lark-cli

```bash
cd /đường/dẫn/tới/lark-cli          # kho lark-cli của bạn (cùng cấp với kho này: ../lark-cli)
go build -o ~/bin/lark-cli .          # hoặc target trong Makefile của kho
export PATH="$HOME/bin:$PATH"
lark-cli --version
lark-cli auth login                    # mở trình duyệt; lưu credential vào Keychain
lark-cli auth status                   # kỳ vọng: user (và/hoặc bot) = ready
```

## 2. Chọn transport

```bash
cd /đường/dẫn/tới/lark-cowork-plugins  # thư mục gốc của kho này
./set-transport.sh stdio               # Desktop cổ điển (nên bắt đầu bằng cách này)
# hoặc ./set-transport.sh http         # Cowork VM / từ xa
```

Với `http`, cần khởi động thêm cầu nối từ xa + tunnel và export `LARK_MCP_URL` +
`LARK_MCP_BEARER_TOKEN` (xem mục "Chọn phương thức kết nối" trong README).

## 3. Smoke-test cầu nối trực tiếp (chưa cần Claude)

```bash
lark-cli mcp tools                      # liệt kê bề mặt công cụ lark_* (offline)
# Tùy chọn: khởi tạo + một lệnh gọi thật qua helper mcp-test / mcp-call của bạn
```

## 4. Cài vào Claude

- **Claude Code:** `claude plugin marketplace add .` rồi
  `claude plugin install productivity@lark-cowork`.
- **Cowork (Desktop):** thêm thư mục này như marketplace cục bộ; cài từ giao diện plugin.

## 5. Kiểm chứng trên giao diện (KHÔNG tin exit code 0)

Theo quy tắc dự án: một tính năng chưa xong cho tới khi được kiểm chứng trực quan. Chạy một quy trình
thật và xác nhận Claude thực sự gọi `lark_*`:

1. Trong thư mục làm việc của một plugin, chạy `/productivity:start` (hoặc gọi một skill).
2. Theo dõi các lệnh gọi công cụ `lark_im_*`, `lark_doc_*`, `lark_task_*`, v.v.
3. Đối chiếu chéo với audit log nếu chạy http: `~/.lark-mcp-audit.ndjson` ghi lại từng lệnh gọi.

## Kiểm chứng cấu trúc đã thực hiện sẵn

```bash
# tất cả JSON hợp lệ
python3 -c "import json,glob;[json.load(open(f)) for f in glob.glob('**/*.json',recursive=True)];print('ok')"
# mọi plugin nghiệp vụ đều có connector lark
grep -rl '\"lark\"' */.mcp.json | wc -l            # -> 15
# không còn server SaaS chung nào
grep -hoE '\"(slack|notion|gmail|google calendar|asana|linear|atlassian)\"' */.mcp.json   # -> rỗng
```

## Các khoảng trống đã biết / câu hỏi mở

- **R1 — Khả năng tiếp cận stdio của Cowork.** Chưa kiểm chứng VM của Cowork có spawn được binary
  stdio trên máy chủ hay không. Nếu không, transport `http` là bắt buộc cho Cowork. (stdio Desktop cổ
  điển đã được chứng minh.)
- **R2 — phân giải một server duy nhất.** Mọi danh mục ánh xạ về một server `lark`; CONNECTORS.md nêu
  rõ tên công cụ chính xác để Claude định tuyến đúng. Xác nhận bằng một quy trình đa công cụ.
- **R3 — CRM/analytics.** Lark không có CRM/product-analytics gốc; những thứ đó giữ kết nối ngoài hoặc
  được mô hình hóa trong Lark Base (dùng skill `base-deploy` để dựng khung CRM Base).
- **R4 — bề mặt công cụ.** Cầu nối phơi bày 25 công cụ được tuyển chọn + `lark_api`. Thao tác ngoài bộ
  này đi qua `lark_api` và có thể cần thêm scope ứng dụng Lark.
- **R5 — Lark vs Feishu.** Cấu hình cho Lark quốc tế (larksuite.com). Feishu (feishu.cn) sẽ cần
  endpoint/scope khác.

---

## Tác giả

**Nguyễn Ngọc Tuấn**
Founder Transform Group — **Lark Platinum Partner**
🌐 Dự án: [larkcowork.com](https://larkcowork.com)
