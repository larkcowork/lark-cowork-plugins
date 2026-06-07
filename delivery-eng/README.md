# Plugin Bàn giao Kỹ thuật (Engineering Delivery)

Vận hành các nghi thức bàn giao kỹ thuật (engineering delivery) trên Lark — postmortem sự cố và sprint retro. Claude kéo
dữ liệu từ kênh chat trực sự cố (on-call), trình theo dõi ticket, và lịch sử velocity để soạn ra những tài liệu mà team
của bạn cần, một cách không đổ lỗi (blameless) và dựa trên dữ liệu thay vì cảm tính.

## Cài đặt

```
claude plugins add lark-cowork/delivery-eng
```

## Tính năng

Plugin này biến tín hiệu bàn giao thô thành các tài liệu retro có cấu trúc:

- **incident-retro** — Một postmortem (mổ xẻ sự cố) không đổ lỗi được dựng từ dòng thời gian IM của đội trực sự cố: trình tự
  các sự kiện, các yếu tố góp phần, và action item kèm người phụ trách và hạn chót. Tái dựng lại quá trình phát hiện →
  phân loại (triage) → giảm thiểu (mitigation) → khắc phục (resolution) mà không quy trách nhiệm cho cá nhân.
- **sprint-retro** — Một bản nháp cuối sprint được dựng từ các ticket đã đóng, độ chênh velocity, và tín hiệu về điểm nghẽn
  (blocker), kèm theo một biểu mẫu retro tùy chọn. Nêu bật điều gì đã tốt, điều gì chưa, và điều gì
  nên thử ở sprint kế tiếp.

## Kỹ năng

| Kỹ năng | Mô tả |
|-------|-------------|
| `incident-retro` | Dựng một tài liệu postmortem không đổ lỗi từ dòng thời gian IM của đội trực sự cố |
| `sprint-retro` | Soạn bản nháp retro cuối sprint từ các ticket đã đóng, velocity, và các điểm nghẽn |

## Quy trình mẫu

### Postmortem cho một SEV2

```
Bạn: viết postmortem cho sự cố SEV2 hôm qua

Claude: [Tìm trong kênh IM trực sự cố theo khoảng thời gian xảy ra sự cố]
        [Tái dựng dòng thời gian: phát hiện → page → triage → giảm thiểu]
        [Đối chiếu chéo với các PR đã deploy và các dashboard]
        [Soạn một postmortem không đổ lỗi (tóm tắt, tác động, nguyên nhân gốc,
         các yếu tố góp phần, action item) qua template postmortem]
        [Xác nhận trước khi lưu vào wiki hoặc tạo các công việc theo sau]
```

### Sprint retro

```
Bạn: sprint retro cho sprint này

Claude: [Xác định khoảng thời gian của sprint từ lịch sử velocity]
        [Kéo các ticket đã đóng từ Base/Task và tính độ chênh velocity]
        [Gom phản hồi từ biểu mẫu retro và nhóm tín hiệu điểm nghẽn từ chat]
        [Soạn một retro một trang: các con số, điều gì tốt/chưa tốt,
         nên thử tiếp, câu hỏi còn mở]
        [Đề nghị lưu vào wiki và cập nhật lịch sử velocity]
```

## Plugin đồng hành

Các kỹ năng này thuộc bộ quy trình **lark-cowork** và tham chiếu đến các kỹ năng ở những plugin
anh em. Tên kỹ năng được phân giải ở phạm vi toàn cục, nên một tham chiếu sẽ hoạt động tự động khi plugin
đồng hành đã được cài; khi vắng plugin đồng hành, tham chiếu sẽ **suy giảm một cách mượt mà** (bước đó
bị bỏ qua hoặc được đề xuất như một gợi ý, không bao giờ là lỗi). Hãy cài các plugin đồng hành bên dưới để có
trải nghiệm đầy đủ — xem [`../connectors/LARK-FUSION.md`](../connectors/LARK-FUSION.md).

| Kỹ năng của plugin này | Tham chiếu đến | Trong plugin |
|---|---|---|
| `incident-retro`, `sprint-retro` | `doc-from-template` (template postmortem) | knowledge-docs |

## Nguồn dữ liệu

> Nếu bạn thấy các placeholder lạ hoặc cần kiểm tra công cụ nào đang được kết nối, xem [CONNECTORS.md](CONNECTORS.md).

**Kết nối MCP đi kèm:** máy chủ `lark` (`lark-cli mcp serve`), một cầu nối duy nhất bao trùm mọi hạng mục —
- Chat (Lark IM) cho dòng thời gian trực sự cố và tín hiệu điểm nghẽn
- Trình theo dõi dự án (Lark Task + Lark Base) cho các ticket đã đóng và velocity
- Tài liệu (Lark Docs + Wiki) cho các tài liệu postmortem và retro
- Trí tuệ cuộc họp (Lark Minutes) cho ngữ cảnh phòng tác chiến (war-room) sự cố
- Danh bạ (Lark Contact) để phân giải người phụ trách

Xem [CONNECTORS.md](CONNECTORS.md) để có bản đồ công cụ đầy đủ và các tài liệu tham chiếu chuyên sâu về Lark.

---

## Tác giả

**Nguyễn Ngọc Tuấn**
Founder Transform Group — **Lark Platinum Partner**
🌐 Dự án: [larkcowork.com](https://larkcowork.com)
