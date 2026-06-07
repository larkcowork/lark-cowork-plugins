# 6. Câu hỏi thường gặp (FAQ)

### Tôi có cần biết lập trình không?
**Không.** Bạn nói chuyện bằng tiếng Việt như với đồng nghiệp. Phần kỹ thuật (cài cầu nối) do
admin/IT làm một lần.

### Nó khác gì ChatGPT/AI chat thường?
AI chat thường chỉ **trả lời**. Lark Cowork **làm việc thật trên Lark của bạn**: đọc mail, tạo lịch,
ghi dữ liệu vào Base, soạn tài liệu, mở đơn duyệt… ngay trong Lark, không phải copy-dán qua app khác.

### Trợ lý có tự ý gửi email / sửa dữ liệu của tôi không?
**Không.** Việc có tác động luôn **xem trước rồi mới làm**; email mặc định **chỉ lưu nháp**. Bạn là
người bấm "chốt".

### Nó có thấy hết dữ liệu công ty không?
Chỉ thấy **đúng phần bạn vốn có quyền** — vì nó đăng nhập **bằng danh tính của bạn**. Xem
[5. An toàn](./05-an-toan-tin-cay.md).

### Tôi không nhớ tên kỹ năng / câu lệnh thì sao?
Không cần nhớ. Cứ mô tả việc. Trợ lý tự chọn kỹ năng đúng. Bí quá thì gõ
💬 *"Bạn giúp được tôi những gì?"*.

### Công ty tôi đang dùng HubSpot / QuickBooks / Figma / Zoom… có bỏ không?
**Không bắt bỏ.** Lark Cowork thay phần **chat/mail/lịch/tài liệu/việc** bằng Lark; còn các công cụ
chuyên dụng (CRM, kế toán, thiết kế, họp…) **giữ nguyên** như kết nối bên ngoài.

### Dùng Lark quốc tế hay Feishu (Trung Quốc)?
Bản này cấu hình cho **Lark quốc tế (larksuite.com)**. Feishu cần endpoint/cấu hình khác.

### "Trợ lý làm sai" thì sao?
AI có thể nhầm. Vì vậy mới có *xem trước*, *lưu nháp*, *phê duyệt*. Hãy **kiểm bản nháp** trước khi
gửi ra ngoài. Việc dựng giao diện/dữ liệu (Base, dashboard) được trợ lý **tự nghiệm thu** trước khi báo xong.

### Tôi muốn một quy trình riêng cho công ty mình?
Được. Bộ **cowork-plugin-management** giúp tạo/tuỳ biến bộ kỹ năng riêng. Hoặc nhờ đội triển khai
(Transform) dựng theo quy trình của bạn.

### Bắt đầu thử nhanh nhất?
Nếu đã cài: gõ 💬 *"Hôm nay tôi có gì trên bàn?"*. Nếu chưa: xem [2. Bắt đầu](./02-bat-dau.md).

### Có gì chứng minh nó "chạy thật", không phải nói suông?
Có bộ kiểm thử tự động 4 cấp — `./tools/audit.sh --live` (24 phép kiểm) và báo cáo
[`../QC-AUDIT.md`](../QC-AUDIT.md). Đã verify thật trên Lark (đọc dữ liệu, tạo→hoàn thành→xoá việc thử…).

### Chi phí?
**Mã nguồn miễn phí (MIT)** — tự host được toàn bộ. Trả phí (nếu có) là cho **tư vấn & triển khai**,
không phải cho license.

---

## Tác giả

**Nguyễn Ngọc Tuấn**
Founder Transform Group — **Lark Platinum Partner**
🌐 Dự án: [larkcowork.com](https://larkcowork.com)
