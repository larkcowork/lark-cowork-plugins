# 10. Cách tùy biến (Custom) cho công ty bạn

> Mọi thứ trong bộ này chỉ là **markdown + JSON** — dễ đọc, dễ sửa, versioning bằng Git. Trang này đi
> từ **dễ → nâng cao**, kèm ví dụ thật để copy.

**5 cấp độ tùy biến:**

| Cấp | Bạn chỉnh gì | Cần kỹ thuật? | Mục |
|---|---|---|---|
| 1 | Hồ sơ cá nhân/công ty (`settings.local.json`) | Không | [10.1](#101-cấp-1--hồ-sơ-settingslocaljson) |
| 2 | Dạy trợ lý nhớ người/dự án/thuật ngữ (bộ nhớ) | Không | [10.2](#102-cấp-2--bộ-nhớ-nơi-làm-việc) |
| 3 | Sửa nội dung một kỹ năng (`SKILL.md`) | Nhẹ | [10.3](#103-cấp-3--sửa-nội-dung-một-kỹ-năng) |
| 4 | Đổi ánh xạ công cụ (`CONNECTORS.md`) | Vừa | [10.4](#104-cấp-4--đổi-ánh-xạ-công-cụ-connectorsmd) |
| 5 | Tạo plugin/kỹ năng mới cho riêng công ty | Vừa–cao | [10.5](#105-cấp-5--tạo-plugin--kỹ-năng-mới) |

---

## 10.1. Cấp 1 — Hồ sơ (`settings.local.json`)

Nhiều plugin đọc một file `settings.local.json` để cá nhân hóa kết quả (tên, công ty, hạn ngạch,
sản phẩm…). Bạn **không sửa code** — chỉ điền thông tin.

- **Cowork:** lưu file trong bất kỳ thư mục nào bạn đã chia sẻ với Cowork → plugin tự tìm.
- **Claude Code:** lưu tại `<tên-plugin>/.claude/settings.local.json`.

Ví dụ cho **sales**:

```json
{
  "name": "Nguyễn Văn A",
  "title": "Account Executive",
  "company": "Công ty XYZ",
  "quota": { "annual": 1000000, "quarterly": 250000 },
  "product": {
    "name": "Sản phẩm của bạn",
    "value_props": ["Lợi ích 1", "Lợi ích 2"],
    "competitors": ["Đối thủ A", "Đối thủ B"]
  }
}
```

> 💡 Nếu chưa có file, plugin sẽ **hỏi bạn tương tác** rồi dùng. Tạo file giúp khỏi hỏi lại mỗi lần.

---

## 10.2. Cấp 2 — Bộ nhớ nơi làm việc

Plugin **productivity** cho Claude một "bộ nhớ" hai tầng để hành xử như đồng nghiệp:
- `CLAUDE.md` — bộ nhớ làm việc (ngữ cảnh thường dùng).
- thư mục `memory/` — lưu trữ sâu (người, dự án, thuật ngữ, viết tắt).

Khởi tạo và bồi đắp:

```
/start                  # tạo TASKS.md, CLAUDE.md, memory/, dashboard
/update --comprehensive # quét mail/lịch/chat → gợi ý việc sót & "người mới" để ghi nhớ
```

Sau khi có bộ nhớ, Claude giải mã viết tắt tức thì:
> 💬 *"nhờ Todd làm PSR cho deal Oracle"* → Claude hiểu *Todd Martinez (Finance), báo cáo Pipeline
> Status cho deal Oracle Systems (2.3M, chốt Q2)* — không hỏi lại.

👉 Xem chi tiết ở [README plugin productivity](../productivity/README.md).

---

## 10.3. Cấp 3 — Sửa nội dung một kỹ năng

Mỗi kỹ năng là một file `SKILL.md` gồm **frontmatter** (metadata) + **phần thân** (hướng dẫn). Bạn
có thể sửa giọng văn, thêm bước, đổi mẫu kết quả cho hợp văn hóa công ty.

Cấu trúc một `SKILL.md`:

```markdown
---
name: morning-brief
description: Bộ tóm tắt đầu ngày... Triggers "morning", "/morning".
version: 1.0.0
last_updated: 2026-05-11
---

# morning-brief

(Phần thân: quy trình từng bước, công cụ dùng, định dạng kết quả...)
```

**Quy tắc an toàn khi sửa:**
- ✅ Sửa **phần thân** thoải mái (thêm bước, đổi mẫu, dịch giọng văn, thêm ví dụ công ty).
- ⚠️ **GIỮ NGUYÊN `name:`** trong frontmatter (đổi tên sẽ làm các tham chiếu fusion gãy).
- ⚠️ Giữ đúng tên công cụ `lark_*` thật (xem danh sách ở [9. Cách vận hành](./09-cach-van-hanh.md)).
- 💡 Thêm/đổi từ khóa trong `description` (mục `Triggers`) để kỹ năng tự kích hoạt đúng câu lệnh
  tiếng Việt của team bạn.

> Ví dụ: thêm trigger tiếng Việt cho `inbox-zero` → trong `description` thêm `"dọn hộp thư", "xử lý mail tồn"`.

---

## 10.4. Cấp 4 — Đổi ánh xạ công cụ (`CONNECTORS.md`)

Mỗi plugin có `CONNECTORS.md` ánh xạ **danh mục** (`~~chat`, `~~email`, `~~CRM`…) tới công cụ cụ thể.
Trong bản Việt hóa, mọi danh mục cộng tác trỏ về **một server `lark`**. Bạn chỉ cần sửa file này khi:

- Muốn **thay một danh mục chuyên ngành** bằng công cụ khác (vd CRM ngoài thay vì Base CRM).
- Muốn **thêm công cụ chuyên ngành** (kho dữ liệu, thanh toán…) cho một plugin.

Cách hoạt động: file tham chiếu placeholder `~~category`; bảng trong `CONNECTORS.md` quyết định nó
phân giải về công cụ nào. Bất kỳ thao tác Lark nào chưa có công cụ riêng → dùng `lark_api` (truyền
path OpenAPI). Bản đồ tổng: [`../connectors/CONNECTORS.lark.md`](../connectors/CONNECTORS.lark.md).

> 💡 Sau khi sửa, có thể chạy lại bộ chuyển đổi để chuẩn hóa: `cd tools && go run . convert <plugin>`.

---

## 10.5. Cấp 5 — Tạo plugin / kỹ năng mới

Dùng plugin **cowork-plugin-management** — cửa ngõ tạo & tùy biến:

| Kỹ năng | Làm gì |
|---|---|
| `create-cowork-plugin` | Hướng dẫn hội thoại 5 phase (Discovery → Plan → Design → Implement → Package) tạo một plugin mới từ đầu, xuất ra file `.plugin` cài được ngay. |
| `cowork-plugin-customizer` | Tùy biến plugin sẵn có theo công cụ & quy trình của tổ chức bạn. |

> 💬 *"Tạo cho tôi một plugin quản lý đơn hàng nội bộ trên Lark Base."*
> 💬 *"Tùy biến plugin sales theo quy trình bán hàng 5 bước của công ty tôi."*

**Cấu trúc một plugin** (nếu tự dựng tay):

```
<ten-plugin>/
├── .claude-plugin/plugin.json   # name, version, description
├── .mcp.json                    # khai báo server lark (hoặc server chuyên ngành)
├── CONNECTORS.md                # ánh xạ danh mục → công cụ
├── README.md                    # giới thiệu
└── skills/
    └── <ten-ky-nang>/SKILL.md   # bí quyết nghề
```

Để dựng hệ thống dữ liệu (CRM, chấm công, dự án…) **trên Lark Base**, đừng tự code — dùng skill
**`base-deploy`** (8 phase tự động hóa). Xem [README lark-base-deploy](../lark-base-deploy/README.md).

---

## 10.6. Đổi chế độ kết nối (transport)

```bash
./set-transport.sh stdio   # một máy, Desktop cổ điển
./set-transport.sh http    # Cowork VM / nhiều người (qua HTTP + bearer token)
```

Chi tiết ở [`../SETUP.md`](../SETUP.md).

---

## 10.7. Sau khi tùy biến — luôn kiểm chứng

Đừng tin "code 0". Chạy bộ kiểm thử và **xem kết quả thật trên giao diện**:

```bash
./tools/audit.sh --live    # 24 phép kiểm 4 cấp + đọc/ghi thật tự dọn
```

---

➡️ Đã biết chỉnh ở đâu? Sang [11. Best practice](./11-best-practice.md) để biết **nên** chỉnh & dùng thế nào.

---

## Tác giả

**Nguyễn Ngọc Tuấn**
Founder Transform Group — **Lark Platinum Partner**
🌐 Dự án: [larkcowork.com](https://larkcowork.com)
