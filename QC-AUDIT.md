# QC & Audit — marketplace Lark Cowork + cầu nối MCP lark-cli

**Trạng thái: ✅ TẤT CẢ XANH — 24/24 kiểm tra tự động đạt** (chạy `./tools/audit.sh --live`).
Lần chạy gần nhất: 24 đạt, 0 lỗi. Bộ kiểm thử tái lập được: [`tools/audit.sh`](./tools/audit.sh).

---

## 1. Kế hoạch & phạm vi

**Mục tiêu:** kiểm chứng marketplace Lark Cowork (22 plugin, 235 kỹ năng, fork từ
anthropics/knowledge-work-plugins) và cầu nối MCP lark-cli (25 công cụ) là đúng đắn, nhất quán và
**thực sự chạy** — không chỉ "JSON parse được".

**Ngoài phạm vi:** các payload chuyên ngành bio/zoom của bản gốc (giữ kết nối ngoài), transport
Cowork-VM (cần runtime đám mây), và endpoint Feishu-CN (đang cấu hình cho Lark quốc tế).

**Phương pháp:** 4 cấp kiểm chứng, mỗi cấp là một tập test case tự động trong `tools/audit.sh`:

| Cấp | Câu hỏi nó trả lời | Chế độ |
|-----|--------------------|--------|
| **L1 Cấu trúc** | Các file/manifest có khớp nhau không? | tĩnh |
| **L2 Đúng đắn** | Tên công cụ / đường jq / yêu cầu có đúng không? | tĩnh + mã nguồn |
| **L3 Ngữ nghĩa** | Việc "thay hồn" có thật & nhất quán (không chỉ hình thức)? | tĩnh + diff với bản gốc |
| **L4 Runtime** | Nó có thực sự build, serve và gọi được Lark không? | build + handshake MCP + API thật |

---

## 2. Ma trận test case (tự động — ID khớp với output của `tools/audit.sh`)

### L1 — Cấu trúc
| ID | Test | Kỳ vọng | Kết quả |
|----|------|---------|---------|
| L1.1 | Mọi `*.json` parse được | JSON hợp lệ | ✅ |
| L1.2 | Mọi thư mục `source` của plugin trong marketplace tồn tại | không mục treo | ✅ |
| L1.3 | Mọi `plugin.json` có name/version/description | đầy đủ | ✅ |
| L1.4 | Không `CONNECTORS.md` nào có bảng rỗng (kể cả partner) | đều có dòng | ✅ |
| L1.5 | Không lỗi ngữ pháp `kind:` trong thẻ | 0 trường hợp | ✅ |

### L2 — Đúng đắn
| ID | Test | Kỳ vọng | Kết quả |
|----|------|---------|---------|
| L2.1 | Không công cụ `lark_*` "ảo" được tham chiếu trong skill | chỉ 25 (+`lark_api`) | ✅ |
| L2.2 | Không đường jq cấp cao sai trong skill | đều dưới `.data.` | ✅ |
| L2.3 | Tài liệu lõi depth-core có mặt (PATTERNS/RECIPES/FUSION) | 3 file | ✅ |
| L2.4 | Không còn "21 tools" cũ trong tài liệu | ghi 25 | ✅ |
| L2.5 | `toolBaseSearch` bắt buộc `search_fields` | báo lỗi rõ ràng | ✅ |

### L3 — Ngữ nghĩa (chất lượng "thay hồn")
| ID | Test | Kỳ vọng | Kết quả |
|----|------|---------|---------|
| L3.1 | Không SKILL nào bị đổi tên so với bản gốc `/tmp/knowledge-work-plugins` | 0 đổi tên | ✅ |
| L3.2 | Frontmatter còn nguyên (`---` + đúng một `name:`) | sạch | ✅ |
| L3.3 | Mọi SKILL.md first-party đã "thay hồn" (tham chiếu depth-core) | 100% | ✅ |
| L3.4 | Không sót dấu vết meta-commentary của agent trong skill | 0 | ✅ |
| L3.5 | Mọi CONNECTORS của plugin trỏ về depth core | tất cả | ✅ |

### L4 — Runtime
| ID | Test | Kỳ vọng | Kết quả |
|----|------|---------|---------|
| L4.1 | `go vet` (cmd/mcp + okr) | sạch | ✅ |
| L4.2 | `go test` (cmd/mcp + okr) | đạt | ✅ |
| L4.3 | `go build` → `~/bin/lark-cli` | build được | ✅ |
| L4.4 | `mcp tools` (offline) | 25 công cụ | ✅ |
| L4.5 | Handshake MCP stdio `tools/list` | 25 công cụ | ✅ |
| L4.6 | Thẻ tương tác (P4) biên dịch offline | `ok:true` | ✅ |
| L4L.0 | `auth status` ready | user+bot ready | ✅ |
| L4L.1 | ĐỌC THẬT: `lark_contact_search me` → open_id | dữ liệu thật | ✅ |
| L4L.2 | Vòng GHI THẬT: tạo→hoàn thành→xóa task | đã tạo, đã xong, đã mất | ✅ |

> L4L.* chỉ chạy với `--live` (cần `lark-cli auth login`). Vòng ghi chỉ chạm một task `[AUDIT]` tự tạo
> rồi xóa — **không ảnh hưởng dữ liệu hiện có**.

---

## 3. Use case (kịch bản end-to-end) & trạng thái kiểm chứng

| UC | Kịch bản | Công cụ dùng | Cách kiểm chứng | Trạng thái |
|----|----------|--------------|-----------------|------------|
| UC1 | "Hôm nay tôi có gì / tôi là ai" | `lark_contact_search`, `lark_task_my` | L4L.1 đọc thật (profile thật) | ✅ |
| UC2 | Ghi một follow-up thành task, rồi đóng nó | `lark_task_create`, `lark_task_complete` | L4L.2 vòng ghi thật | ✅ |
| UC3 | Xem trước tin nhắn/sự kiện trước khi gửi (an toàn-ghi P2) | `lark_im_send`/`lark_calendar_create` `dry_run` | dry_run thật (không ghi) — phiên trước | ✅ |
| UC4 | Tìm bản ghi trong Base CRM/tracker | `lark_base_search` (+`search_fields`) | thật: query+search_fields → record | ✅ |
| UC5 | Ghi/cập nhật một bản ghi Base (SoR, P5) | `lark_base_record_upsert` | thật: `created:true`+record_id | ✅ |
| UC6 | Lưu tri thức bền vững vào Wiki (P8) | `lark_wiki_node_create` | thật: vòng tạo→xóa | ✅ |
| UC7 | Kiểm tra lịch trống trước khi sắp lịch | `lark_calendar_freebusy` | đọc thật `ok:true` | ✅ |
| UC8 | Hiện quyết định/digest dưới dạng thẻ tương tác (P4) | `lark_im_card_send` | offline `print_json` `ok:true` (L4.6) | ✅ |
| UC9 | Lấy action item từ cuộc họp (P6) | `lark_minutes_search` | đọc thật (`.data.items`) | ✅ |

---

## 4. Checklist QC (ký duyệt)

- [x] Đường ống: 22/22 plugin marketplace phân giải được; 20 mang connector `lark`; pdf-viewer cố ý
- [x] Connector: mọi CONNECTORS.md của plugin có bảng thật + con trỏ depth-core
- [x] Công cụ: 25 công cụ MCP, mọi tên cờ đã kiểm chứng với `shortcuts/`, `go test` xanh
- [x] jq: mọi ví dụ/skill chiếu dưới `.data.` (đã kiểm chứng envelope là `{ok,data,…}`)
- [x] Thay hồn: 14 first-party (135 skill) + productivity (4/4) + 14 skill cộng tác partner; 0 đổi tên
- [x] Bug đã sửa & kiểm chứng lại: ngữ pháp thẻ `kind:`, resolver okr `"me"`, `base_search` (search_fields + format + jq)
- [x] Runtime: build + handshake MCP (25) + biên dịch thẻ + đọc thật + vòng ghi thật
- [x] Vệ sinh dữ liệu: mọi artifact tạm `[TEST-AUDIT]`/`[AUDIT]` đã xóa (0 sót lại)
- [x] Tài liệu: README, SETUP, blueprint, RECIPES ghi 25 công cụ; LARK-RECIPES ghi rõ gotcha base_search

---

## 5. Đợt deep-QC song song (21 agent)

Một workflow gồm 20 agent QC theo từng plugin + 1 agent về hình dạng công cụ đã review **mọi** skill
và sửa **100 vấn đề** trên 20 plugin (base_search nay mang `search_fields` + "no jq" trong mọi ví dụ;
phép chiếu jq dưới `.data.`; ngữ pháp thẻ; không công cụ ảo). Đã kiểm chứng lại: `audit.sh --live` vẫn
**24/24 xanh**, `go test` xanh. Sau đó **đóng cả hai hạn chế trước đó**:

- ✅ **Hình dạng `vc_search` ĐÃ KIỂM CHỨNG THẬT** — một tìm kiếm theo khoảng ngày trả về 15 cuộc họp
  thật đã qua; hình dạng thật là `.data.items[]` (mỗi item: id, display_info, meta_data) — ví dụ cũ
  `.data.meetings[]` là sai và nay đã sửa trong `tools.go`.
- ✅ **Hình dạng `sheets_read` đã xác nhận** đúng (không cần đổi).

## 6. Hạn chế đã biết (trung thực)

- **Transport Cowork-VM** (http) chưa được test runtime ở đây — chỉ stdio Desktop-cổ-điển đã chứng minh.
- **`base_search` yêu cầu `search_fields`** theo thiết kế API; skill nay nêu rõ điều này + cách phát
  hiện tên field; công cụ báo lỗi rõ ràng nếu thiếu.

---

## 7. Cách chạy lại

```bash
cd lark-cowork-plugins
./tools/audit.sh            # L1–L4 offline (tĩnh + build + handshake + thẻ)
./tools/audit.sh --live     # + đọc thật qua xác thực & vòng task tự dọn
# exit code 0 = tất cả xanh; in PASS/FAIL theo từng test id
```

---

## Tác giả

**Nguyễn Ngọc Tuấn**
Founder Transform Group — **Lark Platinum Partner**
🌐 Dự án: [larkcowork.com](https://larkcowork.com)
