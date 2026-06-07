# 5. An toàn & tin cậy (cho người dùng nghiệp vụ)

Trợ lý này **thao tác trên dữ liệu thật của công ty**, nên an toàn là mặc định — không phải tuỳ chọn.
Trang này giải thích bằng ngôn ngữ thường, để bạn (và sếp/IT) yên tâm.

## 6 lớp bảo vệ bật sẵn

| Lớp | Nghĩa là gì cho bạn |
|---|---|
| **Xem trước rồi mới làm** (dry-run) | Mọi việc có tác động (gửi mail, tạo lịch, ghi dữ liệu) đều **hiện bản nháp/kế hoạch trước**; chỉ chạy khi bạn đồng ý. |
| **Email mặc định chỉ lưu nháp** | Trợ lý **không tự gửi email**. Nó soạn vào thư nháp; bạn bấm gửi. |
| **Dùng đúng quyền của bạn** | Trợ lý đăng nhập Lark **bằng danh tính của chính bạn** → chỉ thấy dữ liệu bạn vốn có quyền. Không có "cửa sau". |
| **Khoá an toàn thông tin đăng nhập** | Token lưu trong *Keychain* của máy (két an toàn của hệ điều hành), không để file trần. |
| **Kết nối từ xa có mã bảo vệ** | Khi chạy chế độ nhiều người (HTTP), mọi yêu cầu phải kèm *token bí mật*; chỉ mở cổng cần thiết. |
| **Nhật ký kiểm toán** | Mỗi lệnh được ghi lại (audit log) để rà soát khi cần. |

## Dữ liệu của tôi đi đâu?

- Ở **chế độ trên máy (stdio)** — dữ liệu chỉ đi **tới Lark** của bạn, không qua bên thứ ba nào khác.
- Trợ lý AI chỉ nhận **đúng phần dữ liệu cần** cho việc bạn yêu cầu (và thường được "rút gọn" để tiết kiệm),
  không hút toàn bộ kho dữ liệu.
- **Tự chủ:** mã nguồn mở, chạy trên hạ tầng của bạn — không khoá nhà cung cấp.

## Bạn luôn nắm quyền

- Trợ lý **hỏi xác nhận** trước khi làm việc nhạy cảm.
- Việc cần người gật đầu (duyệt chi, ký hợp đồng, đổi lương) đi qua **luồng phê duyệt thật** của Lark —
  không phải "trả lời yes" qua loa.
- Tài liệu quan trọng được **đưa vào đúng nơi** (Wiki/Drive) để lưu vết, không trôi trong chat.

## Trung thực về giới hạn (đọc kỹ)

- AI có thể **nhầm** hoặc bị **đánh lừa qua nội dung độc hại** (prompt injection). Sau khi được cấp quyền,
  trợ lý hành động **dưới danh tính bạn** trong phạm vi quyền đó — nên hãy **dùng thận trọng**.
- **Khuyến nghị triển khai an toàn:** bắt đầu ở **chế độ trên máy cho một vài người** trước, rồi mới mở
  rộng chế độ nhiều người. Với bot, nên dùng như **trợ lý cá nhân riêng tư**, cân nhắc trước khi thả vào
  group chung.
- Luôn **xem bản nháp** trợ lý đưa ra trước khi gửi đi bên ngoài (email, tin cho khách).

## Cho IT/Admin (tóm tắt)

- Quản trị quyền bằng chính *scope* OAuth của Lark; mỗi người một token riêng → không rò chéo.
- Bật `--audit-log` khi chạy; đặt *bearer token* khi mở HTTP; chỉ bind nội bộ trừ khi cố ý phơi qua tunnel.
- Bộ kiểm thử QC: `./tools/audit.sh --live` (chi tiết [`../QC-AUDIT.md`](../QC-AUDIT.md)).

➡️ Còn thắc mắc? Xem [6. Câu hỏi thường gặp](./06-faq.md).

---

## Tác giả

**Nguyễn Ngọc Tuấn**
Founder Transform Group — **Lark Platinum Partner**
🌐 Dự án: [larkcowork.com](https://larkcowork.com)
