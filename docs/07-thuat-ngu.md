# 7. Thuật ngữ (giải thích cho người dùng nghiệp vụ)

Gặp từ lạ trong tài liệu hay khi dùng? Tra ở đây.

## Khái niệm chung

| Từ | Nghĩa dễ hiểu |
|---|---|
| **Lark Cowork** | Bộ này: trợ lý AI biết làm việc, sống trong Lark. |
| **Trợ lý / Agent** | "Đồng nghiệp AI" thực thi việc nhiều bước, không chỉ trả lời. |
| **Plugin / Bộ kỹ năng** | Một gói kỹ năng theo nghề (sales, CSKH, kế toán…). Có 29 bộ. |
| **Skill / Kỹ năng** | Một quy trình công việc cụ thể (vd "tóm tắt cuộc gọi", "đối soát"). ~247 cái. |
| **Câu lệnh (prompt)** | Câu tiếng Việt bạn gõ cho trợ lý. Không cần cú pháp. |
| **Thẻ tương tác (card)** | Tin nhắn Lark có **nút bấm** (Duyệt/Từ chối/Làm tiếp) — hành động ngay trong chat. |
| **Xem trước (dry-run)** | Trợ lý hiện kế hoạch/bản nháp trước, chỉ làm khi bạn đồng ý. |
| **Phê duyệt (Approval)** | Luồng duyệt thật của Lark cho việc cần gật đầu (chi tiền, ký, đổi lương). |

## Sản phẩm trong Lark

| Từ | Là gì |
|---|---|
| **IM** | Chat/tin nhắn Lark. |
| **Mail** | Email Lark. |
| **Lịch (Calendar)** | Lịch & lịch họp; có cả xem bận/rảnh, đặt phòng họp. |
| **Docs** | Tài liệu trực tuyến (như Google Docs). |
| **Wiki** | Kho tri thức/tài liệu nội bộ có cấu trúc cây (nơi lưu KB, SOP, runbook). |
| **Base (Bitable)** | **Bảng dữ liệu thông minh** — vừa như Excel vừa như cơ sở dữ liệu (CRM, theo dõi đơn, sổ rủi ro…). Là "nguồn sự thật" cho dữ liệu có cấu trúc. |
| **Sheets** | Bảng tính trực tuyến. |
| **Drive** | Lưu trữ file (như Google Drive). |
| **Minutes** | **Biên bản họp AI** — tự tóm tắt, rút việc cần làm, lời thoại. |
| **VC** | Họp video; tra cứu cuộc họp đã diễn ra. |
| **Task** | Quản lý việc/todo của Lark. |
| **OKR** | Mục tiêu & kết quả then chốt. |
| **Contact** | Danh bạ tổ chức (đổi tên người → mã định danh `open_id` để gửi/mời). |

## Kỹ thuật (chỉ cần khi đọc phần admin)

| Từ | Là gì |
|---|---|
| **MCP** | "Chuẩn cắm" để trợ lý AI gọi công cụ ngoài. Ở đây: cầu nối tới Lark. |
| **lark-cli** | Cầu nối (chương trình) đứng giữa trợ lý và Lark; phơi ra 25 "công cụ" `lark_*`. |
| **Công cụ `lark_*`** | Hành động đơn lẻ trợ lý gọi (vd `lark_im_send` = gửi tin, `lark_base_search` = tìm bản ghi). |
| **lark_api** | "Cửa thoát hiểm": gọi thẳng API Lark khi chưa có công cụ chuyên cho việc đó. |
| **stdio / HTTP** | Hai cách kết nối cầu nối: trên-máy (stdio) hoặc từ-xa nhiều người (HTTP + token). |
| **search_fields** | Khi tìm trong Base phải nói **tìm ở cột nào** (yêu cầu của API Lark Base). |
| **Keychain** | Két an toàn của hệ điều hành — nơi lưu token đăng nhập. |
| **open_id** | Mã định danh người dùng Lark (không đoán được; phải tra qua Contact). |
| **base-deploy** | Quy trình 8 bước dựng cả một hệ Base end-to-end (thiết kế → bảng → dashboard → nhắc việc → phân quyền → bàn giao). |

## Mô hình & vai trò

| Từ | Là gì |
|---|---|
| **Cowork** | Chế độ trợ lý AI làm-việc-giúp của Claude (chạy trong Claude Desktop). |
| **Depth core** | 3 tài liệu "lõi" hướng dẫn trợ lý làm đúng kiểu Lark (PATTERNS/RECIPES/FUSION) trong thư mục `connectors/`. |
| **Fusion** | Nguyên tắc: bộ kỹ năng theo nghề **gọi lại** các kỹ năng Lark chuyên sâu thay vì tự chế. |
| **Transform** | Đơn vị phát triển & triển khai (mô hình: mã nguồn miễn phí + dịch vụ có phí). |

⬅️ Quay lại [mục lục tài liệu](./README.md).

---

## Tác giả

**Nguyễn Ngọc Tuấn**
Founder Transform Group — **Lark Platinum Partner**
🌐 Dự án: [larkcowork.com](https://larkcowork.com)
