# Plugin Trợ lý Hằng ngày (Daily Assistant)

Buồng lái Lark duy nhất cho ngày làm việc của bạn — tổng hợp brief, triage và chuẩn bị họp ở cùng một nơi. Claude đọc chat, mail, lịch, công việc và biên bản họp của bạn, rồi trả về những bản tóm tắt đã xếp hạng, loại bỏ trùng lặp và sẵn sàng hành động thay vì những luồng dữ liệu thô. Bắt đầu ngày làm việc bằng một bản brief, cắt bỏ nhiễu khi bạn quá tải, và bước vào mọi cuộc họp trong tư thế đã chuẩn bị sẵn sàng.

## Cài đặt

```
claude plugins add lark-cowork/daily-assistant
```

## Tính năng

Plugin này bao trùm toàn bộ luồng công việc hằng ngày, được nhóm thành bốn chủ đề:

- **Bắt đầu & kết thúc ngày** — `morning-brief` mở đầu ngày bằng một bản brief đã hợp nhất; `daily-digest` khép lại ngày; `weekly-review` lùi lại để nhìn toàn cảnh cả tuần.
- **Triage** — cắt xuyên qua đống tồn đọng: `im-digest` cho chat, `inbox-zero` cho mail, `task-prioritizer` cho việc cần làm, và `overwhelm-triage` để điều hướng bạn đến đúng kỹ năng khi mọi thứ trở nên quá sức.
- **Tập trung** — bảo vệ thời gian làm việc sâu bằng `focus-mode`, và giành lại thời gian bằng cách rà soát lịch của bạn với `calendar-optimizer`.
- **Họp & con người** — `meeting-prep` để có ngữ cảnh trước và rút ra action item sau cuộc họp, `one-on-one-prep` cho các buổi 1:1, và `contact-360` để có một bản brief quan hệ đầy đủ trước khi gặp bất kỳ ai.

## Kỹ năng

| Kỹ năng | Mô tả |
|-------|-------------|
| `morning-brief` | Bộ điều phối brief đầu ngày — phân tán song song việc triage mail/IM/approval/task rồi hợp nhất thành ≤15 dòng (5 biến thể: default, ic, exec, pm, sales) |
| `daily-digest` | Bản tổng kết cuối ngày — các cuộc họp trong ngày, công việc đã hoàn thành, và những điểm nổi bật trong hộp thư |
| `weekly-review` | Cái nhìn theo tuần về lịch, công việc và OKR như một câu chuyện liền mạch |
| `im-digest` | Triage các nhóm IM — phân loại những tin nhắn mới nhất của từng nhóm thành cần action / cần biết / bỏ qua |
| `inbox-zero` | Quy trình triage mail đầy đủ hướng tới một hộp thư trống |
| `task-prioritizer` | Xếp hạng công việc theo deadline × rủi ro, mức độ chặn việc khác, liên kết OKR, và trọng số người giao việc; nêu bật top 5 của ngày |
| `overwhelm-triage` | Bộ định tuyến trung gian cho tình huống "tôi quá tải" — một câu hỏi để phân định nỗi đau ở mail/IM/task/họp, rồi điều phối đến đúng kỹ năng |
| `focus-mode` | Chặn thời gian trên lịch, bật chế độ không làm phiền (DND) trên IM, và thông báo cho team của bạn |
| `calendar-optimizer` | Rà soát mẫu hình họp trong 30 ngày và nêu bật các ứng viên có thể từ chối / gộp / chuyển sang bất đồng bộ |
| `meeting-prep` | Luồng họp hai giai đoạn — kéo ngữ cảnh trước, rút action item thành công việc sau |
| `one-on-one-prep` | Bản brief cho buổi 1:1 sắp tới — OKR, công việc gần đây, ghi chú lần trước, và các câu hỏi gợi ý |
| `contact-360` | Bản brief quan hệ 360° trước một cuộc họp hoặc cuộc gọi — tổng hợp IM, mail, các cuộc họp, tài liệu và công việc liên quan đến một người |

## Quy trình mẫu

### Brief buổi sáng

```
Bạn: morning

Claude: [Phân tán song song mail-triage, im-digest, approval-triage, task-prioritizer]
        [Hợp nhất thành một thẻ ≤15 dòng]
        "3 mail cần trả lời, 2 approval đang chờ, việc ưu tiên hàng đầu: review ngân sách (đến hạn hôm nay).
         Bạn có 4 cuộc họp — cuộc đầu lúc 10:00 với team Platform."
```

### Khi cảm thấy quá tải

```
Bạn: tôi quá tải rồi, không biết bắt đầu từ đâu

Claude: "Áp lực đến từ đâu — mail, chat, công việc, hay họp?"
Bạn: công việc
Claude: [Điều hướng đến task-prioritizer]
        [Xếp hạng theo deadline × rủi ro + mức chặn việc + liên kết OKR]
        "Top 5 cho hôm nay: 1) Ship API spec (đang chặn 2 người) ..."
```

### Trước một cuộc họp

```
Bạn: tôi sắp gặp Sarah Chen

Claude: [Phân giải Sarah qua lark_contact_search, chạy contact-360]
        [Tổng hợp IM gần đây, mail, các cuộc họp chung, tài liệu, công việc còn mở]
        "Lần trao đổi gần nhất 3 ngày trước về roadmap Q2. 2 công việc chung còn mở.
         Cô ấy phụ trách phần review deal Oracle — đến hạn thứ Sáu. Câu mở đầu gợi ý ở bên dưới."
```

## Plugin đồng hành

Các kỹ năng này thuộc bộ quy trình **lark-cowork** và tham chiếu đến các kỹ năng ở những plugin
anh em. Tên kỹ năng được phân giải ở phạm vi toàn cục, nên một tham chiếu sẽ hoạt động tự động khi plugin
đồng hành đã được cài; khi vắng plugin đồng hành, tham chiếu sẽ **suy giảm một cách mượt mà** (bước đó
bị bỏ qua hoặc được đề xuất như một gợi ý, không bao giờ là lỗi). Hãy cài các plugin đồng hành bên dưới để có
trải nghiệm đầy đủ — xem [`../connectors/LARK-FUSION.md`](../connectors/LARK-FUSION.md).

| Kỹ năng của plugin này | Tham chiếu đến | Trong plugin |
|---|---|---|
| `morning-brief` (exec) | `approval-triage` | governance |
| `morning-brief` (pm) | `decision-logger` | governance |
| `morning-brief` (sales) | `pipeline-review`, `client-followup` | crm-sales |
| `meeting-prep` | `decision-logger` | governance |
| `one-on-one-prep` | `doc-from-template` | knowledge-docs |
| `contact-360` | `deal-update` | crm-sales |

## Nguồn dữ liệu

> Nếu bạn thấy các placeholder lạ hoặc cần kiểm tra công cụ nào đang được kết nối, xem [CONNECTORS.md](CONNECTORS.md).

**Kết nối MCP đi kèm:** máy chủ `lark` (`lark-cli mcp serve`), một cầu nối duy nhất bao trùm mọi hạng mục mà plugin này trải rộng —
- Chat (Lark IM) để triage nhóm và quét tin nhắn
- Email (Lark Mail) để triage hộp thư và nêu điểm nổi bật
- Lịch (Lark Calendar) cho lịch trình, khối tập trung, và rà soát họp
- Cơ sở tri thức (Lark Wiki + Docs) để lấy ngữ cảnh tham chiếu
- Trình theo dõi dự án (Lark Task + Lark Base) để xếp hạng công việc và đọc dữ liệu có cấu trúc
- Trí tuệ cuộc họp (Lark Minutes + VC) để lấy ngữ cảnh họp và action item
- Danh bạ (Lark Contact) để phân giải con người trước bất kỳ bản brief theo từng người nào

**Tư thế vận hành:** các kỹ năng này mặc định chỉ đọc — Claude tạo ra bản nháp và gợi ý, và chỉ gửi mail, tạo công việc, chặn lịch, hoặc phản hồi lời mời (RSVP) sau khi bạn xác nhận. Xem [CONNECTORS.md](CONNECTORS.md) để có danh sách công cụ đầy đủ và ghi chú.

---

## Tác giả

**Nguyễn Ngọc Tuấn**
Founder Transform Group — **Lark Platinum Partner**
🌐 Dự án: [larkcowork.com](https://larkcowork.com)
