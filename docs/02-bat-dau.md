# 2. Bắt đầu trong vài phút

Có 2 phần: **(A) cho người dùng cuối** (chỉ cần biết gõ lệnh) và **(B) cho admin/IT** (cài đặt một
lần cho cả nhóm). Nếu công ty bạn đã cài sẵn, hãy đọc thẳng phần A.

---

## A. Cho người dùng cuối — chỉ cần 3 điều

Khi trợ lý đã được bật trong Claude/Lark, bạn **không cần cài gì**. Chỉ cần:

### 1) Nói chuyện bằng tiếng Việt như với đồng nghiệp
Không cần cú pháp. Cứ mô tả việc bạn muốn:
- 💬 *"Sáng nay tôi có gì cần làm và cần duyệt?"*
- 💬 *"Tóm tắt cuộc họp với khách ABC hôm qua, rút ra việc cần làm."*
- 💬 *"Soạn email nhắc anh Minh gửi báo giá, để nháp cho tôi xem."*

### 2) Luôn xem trước khi "chốt"
Với việc có tác động (gửi mail, tạo lịch, ghi dữ liệu), trợ lý **xem trước** rồi mới làm:
- 📤 *"Tôi sẽ gửi email này cho minh@congty.vn, nội dung… — bạn xác nhận gửi chứ?"*
- Bạn duyệt → mới gửi. ⚠️ Email mặc định **chỉ lưu nháp** cho tới khi bạn đồng ý.

### 3) Nhận kết quả ngay trong Lark
Kết quả về dưới dạng tin nhắn, **thẻ có nút bấm** (duyệt/từ chối/làm tiếp), tài liệu, hoặc bản ghi
trong Base — tuỳ việc.

> 💡 Không biết bắt đầu? Gõ: 💬 *"Bạn giúp được tôi những gì?"* hoặc xem
> [3. Danh mục theo phòng ban](./03-danh-muc-theo-phong-ban.md) để lấy câu lệnh mẫu cho đúng nghề.

---

## B. Cho admin / IT — cài một lần

> Phần này có vài bước kỹ thuật nhẹ; làm một lần cho cả nhóm. Chi tiết kỹ thuật đầy đủ ở
> [`../SETUP.md`](../SETUP.md).

**Mô hình:** một "cầu nối" (`lark-cli`) đứng giữa trợ lý AI và Lark. Trợ lý gọi cầu nối, cầu nối
gọi API Lark **bằng quyền của chính người dùng**.

### 3 bước cốt lõi
1. **Cài cầu nối + đăng nhập Lark một lần**
   ```bash
   ./scripts/setup-mcp.sh        # cài lark-cli
   lark-cli auth login           # đăng nhập Lark (token lưu trong Keychain máy)
   lark-cli mcp tools            # phải liệt kê 25 công cụ
   ```
2. **Khai báo trợ lý dùng cầu nối** — thêm `lark` vào cấu hình Claude Desktop, khởi động lại.
3. **Bật bộ kỹ năng theo nghề** — cài các plugin cần dùng từ marketplace này
   (xem [3. Danh mục](./03-danh-muc-theo-phong-ban.md) để chọn theo phòng ban).

### Hai chế độ kết nối
- **Trên máy (stdio)** — đơn giản nhất, dữ liệu chỉ đi tới Lark. Hợp để bắt đầu / cho lãnh đạo.
- **Từ xa (HTTP + bảo mật token)** — cho chế độ Cowork nhiều người. Cần mở cầu nối dạng HTTP
  (xem SETUP.md).

> Đổi chế độ bất cứ lúc nào: `./set-transport.sh stdio` hoặc `./set-transport.sh http`.

### Kiểm tra mọi thứ chạy đúng (QC)
Chạy bộ kiểm thử tự động — phải xanh hết:
```bash
./tools/audit.sh --live     # 24 phép kiểm 4 cấp; báo PASS/FAIL từng mục
```
Báo cáo chất lượng đầy đủ: [`../QC-AUDIT.md`](../QC-AUDIT.md).

➡️ Đã sẵn sàng? Xem việc cụ thể cho nghề của bạn ở [3. Danh mục theo phòng ban](./03-danh-muc-theo-phong-ban.md).

---

## Tác giả

**Nguyễn Ngọc Tuấn**
Founder Transform Group — **Lark Platinum Partner**
🌐 Dự án: [larkcowork.com](https://larkcowork.com)
