# Plugin Tri thức & Tài liệu (Knowledge & Docs)

Plugin về tri thức và tài liệu được thiết kế chủ yếu cho [Cowork](https://claude.com/product/cowork), ứng dụng desktop dạng agent của Anthropic — nhưng cũng hoạt động được trong Claude Code. Biên soạn tài liệu Lark có cấu trúc từ các template và giữ cho Wiki của bạn gọn gàng — Claude điền các template được đặt tên vào những tài liệu, bảng tính và trang wiki sạch sẽ, rồi kiểm toán cơ sở tri thức của bạn để tìm nội dung cũ kỹ (stale), mồ côi (orphan) và trùng lặp (duplicate).

## Cài đặt

```
claude plugins add lark-cowork/knowledge-docs
```

## Tính năng

Plugin này giúp bạn tạo ra và duy trì tri thức có cấu trúc trong Lark:

- **doc-from-template** — Điền một template Mustache được đặt tên vào một tài liệu Lark, một node wiki, hoặc thêm dòng vào sheet. Các template tích hợp sẵn: `weekly-report`, `one-on-one`, `meeting-notes`, `project-kickoff`, `postmortem`, `decision-log-row`, `wiki-runbook`. Kết xuất (render) cục bộ và xem trước trước khi ghi.
- **doc-restructure** — Kiểm toán một Lark Wiki để tìm các trang cũ kỹ, mồ côi và trùng lặp, rồi đề xuất các đợt lưu trữ (archive) / gộp (merge) / đổi cha (re-parent). Chỉ đọc: nó không bao giờ di chuyển hay xóa khi chưa có xác nhận của bạn.

## Kỹ năng

| Kỹ năng | Mô tả |
|-------|-------------|
| `doc-from-template` | Biên soạn một tài liệu/wiki/sheet của Lark bằng cách điền một template Mustache được đặt tên, có kiểm tra biến và xem trước trước khi ghi |
| `doc-restructure` | Kiểm toán Wiki ở chế độ chỉ đọc, gắn cờ các trang cũ kỹ/mồ côi/trùng lặp và đề xuất các đợt lưu trữ/gộp/đổi cha |

## Quy trình mẫu

### Tài liệu báo cáo tuần

```
Bạn: tạo weekly report doc cho tuần này

Claude: [Chọn template weekly-report, lấy các hạng mục của tuần này]
        [Render cục bộ, hiển thị bản xem trước của tài liệu đã điền]
        [Khi xác nhận: tạo một tài liệu DocxXML v2 và trả về URL + token]
```

### Dọn dẹp Wiki

```
Bạn: wiki của team đang bừa, restructure giúp

Claude: [Kiểm kê không gian, làm giàu thông tin node, chạy các quy tắc cũ kỹ/mồ côi/trùng lặp]
        [Báo cáo: 5000 trang — 120 trang rất cũ, 18 trang mồ côi chủ sở hữu, 9 tiêu đề trùng lặp]
        [Đề xuất các đợt ARCHIVE / MERGE / RE-PARENT với cổng phê duyệt cho từng đợt]
        [Không thay đổi gì cho đến khi bạn phê duyệt một đợt]
```

## Plugin đồng hành

Các kỹ năng này thuộc bộ luồng công việc **lark-cowork** và tham chiếu đến các kỹ năng trong những
plugin anh em. Tên kỹ năng được phân giải trên phạm vi toàn cục, nên một tham chiếu sẽ tự động hoạt động khi plugin
đồng hành đã được cài đặt; khi vắng mặt một plugin đồng hành, tham chiếu sẽ **suy giảm một cách nhẹ nhàng** (bước đó
được bỏ qua hoặc được đưa ra như một gợi ý, không bao giờ là lỗi). Hãy cài đặt các plugin đồng hành bên dưới để có trải nghiệm
đầy đủ — xem [`../connectors/LARK-FUSION.md`](../connectors/LARK-FUSION.md).

| Kỹ năng của plugin này | Tham chiếu | Trong plugin |
|---|---|---|
| `doc-from-template` (weekly-report) | `weekly-review` | daily-assistant |

## Nguồn dữ liệu

> Nếu bạn thấy các placeholder lạ hoặc cần kiểm tra công cụ nào đang được kết nối, xem [CONNECTORS.md](CONNECTORS.md).

Hãy kết nối không gian làm việc Lark của bạn để có trải nghiệm tốt nhất. Nếu không có nó, hãy soạn template theo cách thủ công.

**Kết nối MCP đi kèm:** máy chủ `lark` (`lark-cli mcp serve`), một cầu nối (bridge) duy nhất bao trùm mọi danh mục —
- Cơ sở tri thức / wiki (Lark Wiki) cho các trang tham chiếu và kiểm toán
- Tài liệu (Lark Docs) cho các tài liệu được biên soạn — API v2, mặc định DocxXML
- Bảng tính (Lark Sheets) để thêm dòng như nhật ký quyết định (decision log)
- Lưu trữ đám mây (Lark Drive) để tải tệp lên
- Danh bạ (Lark Contact) để phân giải con người thành `open_id`

Tài liệu được biên soạn với `--api-version v2` và DocxXML, và mọi thao tác ghi đều được xem trước trước khi hoàn tất.

**Tùy chọn bổ sung:**
- Xem [CONNECTORS.md](CONNECTORS.md) để biết bản đồ đầy đủ từ danh mục đến công cụ và các ghi chú dành cho Claude.

---

## Tác giả

**Nguyễn Ngọc Tuấn**
Founder Transform Group — **Lark Platinum Partner**
🌐 Dự án: [larkcowork.com](https://larkcowork.com)
