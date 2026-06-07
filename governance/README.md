# Plugin Phê duyệt & Quản trị

Một plugin quản trị được thiết kế chủ yếu cho [Cowork](https://claude.com/product/cowork), ứng dụng desktop dạng agent của Anthropic — nhưng cũng hoạt động được trong Claude Code. Plugin giúp giữ cho hoạt động trên Lark luôn gọn gàng và có trách nhiệm giải trình: thông lượng phê duyệt nhanh hơn, vệ sinh phân quyền lành mạnh hơn, và nguồn gốc quyết định có thể truy vết — để những gì team đã thống nhất, đã duyệt và đã chia sẻ không bao giờ bị thất lạc hay âm thầm trôi ra ngoài chính sách.

## Cài đặt

```
claude plugins add lark-cowork/governance
```

## Tính năng

Plugin này trao cho Claude một lăng kính quản trị trên không gian làm việc Lark của bạn:

- **approval-triage** — Xử lý hàng đợi phê duyệt đang chờ theo từng mục, đề xuất phê duyệt / từ chối / hỏi thêm cho mỗi yêu cầu kèm trích dẫn chính sách. Agent đọc và lý giải; bạn quyết định.
- **approval-flow-sla** — Góc nhìn của người sở hữu quy trình trên toàn bộ luồng phê duyệt: chỉ ra chính xác node/người duyệt đang gây nghẽn, báo cáo thời gian tồn p50/p90, đánh dấu vi phạm SLA, và rà soát các mục đã được duyệt.
- **permission-audit** — Quét chỉ-đọc xuyên suốt Drive/Doc/Wiki/Base để phát hiện phơi nhiễm công khai, bên ngoài và PII, trả về dưới dạng một báo cáo duy nhất xếp hạng theo mức độ nghiêm trọng. Không tự động sửa, không tự động thu hồi.
- **decision-logger** — Phát hiện các quyết định trong IM và Minutes rồi đưa chúng vào một bảng Base có cấu trúc, để chấm dứt tình trạng "đã chốt việc này rồi nhưng không ai nhớ chốt ở đâu".

## Kỹ năng

| Kỹ năng | Mô tả |
|-------|-------------|
| `approval-triage` | Triage hàng đợi Lark Approval đang chờ — đề xuất phê duyệt/từ chối/hỏi thêm cho từng mục kèm trích dẫn chính sách |
| `approval-flow-sla` | Đo lường luồng phê duyệt — phát hiện node/người duyệt gây nghẽn, rà soát dữ liệu đã duyệt, dựng & chấm điểm SLA |
| `permission-audit` | Quét chỉ-đọc Drive/Doc/Wiki/Base để tìm quyền rủi ro (công khai, bên ngoài, PII) |
| `decision-logger` | Phát hiện quyết định trong IM/Minutes và đưa vào bảng Base có cấu trúc |

## Quy trình mẫu

### Dọn sạch hàng đợi phê duyệt

```
Bạn: có gì cần duyệt?

Claude: [Liệt kê 8 phê duyệt đang chờ qua kỹ năng lark-approval]
        [Mỗi mục: ✅ APPROVE / ⚠️ NEEDS CHECK / ❌ REJECT + chính sách §]
        [Gom các mục có thể duyệt theo lô, đánh dấu 2 mục vượt ngưỡng]
        "Approve batch of 5? (reply 'approve lô 1')"
```

### Tìm điểm nghẽn

```
Bạn: nghẽn duyệt ở đâu?

Claude: [Lấy dữ liệu instance 30 ngày, tính thời gian tồn theo từng node]
        "Node 'Finance review' / Linh ngâm p90 = 18h, breach SLA 6h.
         3 đơn cụ thể đang kẹt. Remind hay transfer?"
```

### Ghi nhật ký một quyết định

```
Bạn: chốt cái này — log lại decision

Claude: [Trích quyết định từ luồng IM, xác định người quyết + các bên liên quan]
        [Phản hồi lại Tiêu đề/Quyết định/Người quyết/Ngày/Nguồn/Lĩnh vực]
        "Confirm Status=Active, Review date 90d? Write to decisions Base?"
```

## Plugin đồng hành

Các kỹ năng này thuộc bộ quy trình **lark-cowork** và tham chiếu đến các kỹ năng ở những plugin
anh em. Tên kỹ năng được phân giải toàn cục, nên một tham chiếu sẽ tự động hoạt động khi plugin
đồng hành được cài đặt; khi vắng plugin đồng hành, tham chiếu **suy giảm một cách mượt mà** (bước đó
bị bỏ qua hoặc được đưa ra như một gợi ý, không bao giờ báo lỗi). Hãy cài các plugin đồng hành dưới
đây để có trải nghiệm đầy đủ — xem [`../connectors/LARK-FUSION.md`](../connectors/LARK-FUSION.md).

| Kỹ năng của plugin này | Tham chiếu | Trong plugin |
|---|---|---|
| `decision-logger` | `incident-retro`, `sprint-retro` | delivery-eng |
| `decision-logger` | `meeting-prep`, `morning-brief` | daily-assistant |
| `approval-flow-sla` | `morning-brief` (exec digest) | daily-assistant |

## Nguồn dữ liệu

> Nếu bạn thấy các placeholder lạ hoặc cần kiểm tra công cụ nào đang được kết nối, xem [CONNECTORS.md](CONNECTORS.md).

Hãy kết nối các công cụ phê duyệt, lưu trữ và giao tiếp của bạn để có trải nghiệm tốt nhất. Nếu không có chúng, việc rà soát quản trị phải làm thủ công trong giao diện Lark.

**Kết nối MCP đi kèm:** server `lark` (`lark-cli mcp serve`), một cầu nối duy nhất bao phủ mọi hạng mục —
- Phê duyệt (Lark Approval) qua `lark_api` (approval/v4) hoặc kỹ năng `lark-approval` đã cài — không có công cụ được tuyển chọn riêng
- Lưu trữ đám mây + cơ sở tri thức (Lark Drive / Wiki / Docs) để quét quyền ở chế độ chỉ-đọc
- Cơ sở dữ liệu / bản ghi (Lark Base) cho bảng decisions
- Trí tuệ chat và cuộc họp (Lark IM + Lark Minutes) để phát hiện quyết định

**Trạng thái mặc định:** chỉ-đọc. Chỉ nhắc, chuyển, duyệt hay ghi sau khi có xác nhận rõ ràng.

**Tùy chọn bổ sung:**
- Xem [CONNECTORS.md](CONNECTORS.md) để biết cách từng hạng mục được phân giải và lối thoát `lark_api`

---

## Tác giả

**Nguyễn Ngọc Tuấn**
Founder Transform Group — **Lark Platinum Partner**
🌐 Dự án: [larkcowork.com](https://larkcowork.com)
