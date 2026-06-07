# 9. Cách vận hành (bên dưới chạy thế nào)

> Trang này giải thích **cơ chế** để bạn tin tưởng và dùng đúng. Không cần là kỹ thuật vẫn đọc được
> phần đầu; phần sau dành cho IT/admin.

## 9.1. Bốn thành phần

```
┌────────────┐   lệnh tiếng Việt   ┌──────────────┐   gọi công cụ   ┌────────────┐   API Lark   ┌────────┐
│   BẠN      │ ──────────────────▶ │  Trợ lý AI   │ ──────────────▶ │  lark-cli  │ ───────────▶ │  LARK  │
│ (trong     │                     │  (Claude +   │   lark_*        │ (cầu nối   │  quyền của   │ chat,  │
│  Lark/     │ ◀────────────────── │   kỹ năng)   │ ◀────────────── │   MCP)     │ ◀─────────── │ mail,  │
│  Claude)   │   kết quả về Lark   └──────────────┘   dữ liệu thật  └────────────┘   bạn         │ base…  │
└────────────┘                                                                                    └────────┘
```

1. **Bạn** — ra lệnh bằng tiếng Việt tự nhiên.
2. **Trợ lý AI (Claude)** — chọn **kỹ năng** phù hợp, lập kế hoạch nhiều bước.
3. **lark-cli (cầu nối MCP)** — dịch ý định thành lệnh gọi API Lark, **bằng quyền của chính bạn**.
4. **Lark** — nơi dữ liệu sống và nơi kết quả trả về.

## 9.2. Kỹ năng (skill) là gì

Mỗi **kỹ năng** là một file `SKILL.md` chứa "bí quyết nghề": quy trình từng bước, công cụ nên dùng,
mẫu kết quả. Có **2 cách** một kỹ năng được kích hoạt:

| Cách | Khi nào | Ví dụ |
|---|---|---|
| **Tự kích hoạt** | Bạn nói tự nhiên, Claude nhận ra ngữ cảnh khớp "trigger" của kỹ năng | 💬 *"khách im lặng quá lâu"* → tự gọi `client-followup` |
| **Lệnh slash `/`** | Bạn chủ động gõ để chạy đúng quy trình đó | 💬 *"/pipeline-review"* → chạy review pipeline |

> Bạn **không cần nhớ tên kỹ năng**. Cứ mô tả việc; nếu muốn chắc chắn, gõ `/` để xem danh sách lệnh.

## 9.3. Một yêu cầu chạy qua những bước nào

Ví dụ: 💬 *"Tôi sắp gặp anh Minh — cho tôi context."*

1. **Hiểu ý định** → khớp kỹ năng `contact-360`.
2. **Phân giải người** → `lark_contact_search "Minh"` → ra `open_id` (luôn tìm người trước khi thao tác).
3. **Thu thập song song** → đọc IM, mail, lịch chung, tài liệu, task liên quan tới anh Minh.
4. **Tổng hợp** → gộp, xếp hạng, loại trùng.
5. **Trả kết quả về Lark** → một thẻ tóm tắt (lần trao đổi gần nhất, việc chung đang mở, gợi ý mở đầu).

## 9.4. Cầu nối lark-cli & bộ 25 công cụ

Toàn bộ danh mục cộng tác đi qua **một** server MCP duy nhất tên `lark` (`lark-cli mcp serve`). Nó
phơi bày **~25 công cụ `lark_*`** được tuyển chọn, cộng một "cửa thoát" `lark_api` cho mọi endpoint
Lark OpenAPI chưa có công cụ riêng:

| Nhóm | Sản phẩm Lark | Công cụ `lark_*` chính |
|---|---|---|
| Chat | Lark IM | `lark_im_send`, `lark_im_search`, `lark_im_card_send` |
| Email | Lark Mail | `lark_mail_send`, `lark_mail_draft_create` |
| Lịch | Lark Calendar | `lark_calendar_agenda`, `lark_calendar_create`, `lark_calendar_freebusy` |
| Wiki/Docs | Lark Wiki + Docs | `lark_doc_search`, `lark_doc_fetch`, `lark_doc_create`, `lark_wiki_node_create` |
| Sheets | Lark Sheets | `lark_sheets_read`, `lark_sheets_append` |
| Task/Base | Lark Task + Base | `lark_task_create`, `lark_task_my`, `lark_task_complete`, `lark_base_search`, `lark_base_record_upsert` |
| Drive | Lark Drive | `lark_drive_upload` |
| Họp | Lark Minutes + VC | `lark_minutes_search`, `lark_vc_search` |
| OKR | Lark OKR | `lark_okr_cycle_list` |
| Danh bạ | Lark Contact | `lark_contact_search` |
| Cửa thoát | (mọi API Lark) | `lark_api` (truyền path OpenAPI + tham số) |

**Vài nguyên tắc cầu nối tự lo:**
- **Tìm người trước.** Hầu hết thao tác ghi (gửi tin, mời họp, giao việc) cần `open_id` → luôn
  `lark_contact_search` trước.
- **CRM là một Base.** Không có sản phẩm CRM riêng — CRM là một bảng Lark Base; đọc bằng
  `lark_base_search`, ghi qua `lark_base_record_upsert` / `lark_api`.
- **Danh tính (identity).** Mặc định chạy bằng quyền *user*; vài thao tác cần *bot*. Cầu nối báo rõ
  trong lỗi và tự gợi ý đổi danh tính.

## 9.5. An toàn mặc định (rất quan trọng)

| Cơ chế | Ý nghĩa |
|---|---|
| **Xem trước rồi mới làm** | Việc có tác động (gửi mail, tạo lịch, ghi dữ liệu) được *preview*/`dry_run` để bạn xác nhận. |
| **Email mặc định lưu nháp** | `lark_mail_draft_create` — chỉ gửi sau khi bạn đồng ý. |
| **Chỉ thấy dữ liệu bạn có quyền** | Mọi lệnh chạy bằng *scope OAuth* của chính bạn; không vượt quyền. |
| **Tự kiểm tra trước khi báo "xong"** | Tính năng UI chỉ tính là xong khi được kiểm chứng trực quan, không tin "code 0". |
| **Nhật ký audit** | Bật `--audit-log` → mỗi lệnh gọi ghi vào `~/.lark-mcp-audit.ndjson`. |

## 9.6. Fusion & "depth core" (vì sao kết quả chuẩn "kiểu Lark")

Các bộ kỹ năng theo nghề **không tự chế** cách thao tác Lark — chúng **gọi lại** các kỹ năng Lark
chuyên sâu và tuân theo 3 tài liệu lõi trong `connectors/`:
- **LARK-PATTERNS** — các khuôn mẫu thao tác chuẩn (an toàn-ghi, thẻ tương tác, phân giải người…).
- **LARK-RECIPES** — "công thức" cho thao tác hay gặp (vd `base_search` cần `search_fields`).
- **LARK-FUSION** — nguyên tắc một kỹ năng nghề gọi lại kỹ năng Lark thay vì lặp lại.

Nhờ đó, ví dụ `daily-assistant` khi cần duyệt đơn sẽ **gọi** `approval-triage` của `governance` —
nếu plugin đó được cài; nếu không, bước đó **giảm cấp êm** (bỏ qua hoặc gợi ý), không báo lỗi.

## 9.7. Hai chế độ kết nối (transport)

| Chế độ | Dùng khi | Cách bật |
|---|---|---|
| **stdio** | Claude Desktop cổ điển, một máy có `lark-cli` + đã `auth login` | `./set-transport.sh stdio` |
| **http** | Cowork VM / từ xa (sandbox không gọi được binary máy chủ) | `./set-transport.sh http` + chạy `lark-cli mcp serve --transport http` + tunnel + `LARK_MCP_URL`/`LARK_MCP_BEARER_TOKEN` |

Đổi qua lại bất cứ lúc nào. Chi tiết kỹ thuật ở [`../SETUP.md`](../SETUP.md) và mục "Chọn phương thức
kết nối" của [`../README.md`](../README.md).

## 9.8. Đã kiểm chứng thật, không nói suông

Có bộ kiểm thử 4 cấp (cấu trúc → đúng đắn → ngữ nghĩa → runtime), **24/24 xanh**, gồm đọc dữ liệu
thật và vòng tạo→hoàn thành→xóa task tự dọn. Chạy: `./tools/audit.sh --live`. Báo cáo:
[`../QC-AUDIT.md`](../QC-AUDIT.md).

---

➡️ Muốn nắn theo công ty mình? Sang [10. Tùy biến](./10-tuy-bien.md). Muốn dùng cho hiệu quả nhất?
Xem [11. Best practice](./11-best-practice.md).

---

## Tác giả

**Nguyễn Ngọc Tuấn**
Founder Transform Group — **Lark Platinum Partner**
🌐 Dự án: [larkcowork.com](https://larkcowork.com)
