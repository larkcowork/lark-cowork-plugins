# Plugin Xem PDF (PDF Viewer)

Xem, chú thích và ký PDF trong một trình xem tương tác trực tiếp. Đánh dấu
hợp đồng, điền form với phản hồi trực quan, đóng dấu phê duyệt và đặt
chữ ký — rồi tải về bản đã chú thích.

## Tính năng

- **Mở PDF** từ file cục bộ hoặc các nguồn học thuật (arXiv, bioRxiv, v.v.)
- **Chú thích cộng tác** — Claude đề xuất các điểm tô sáng, ghi chú và
  con dấu theo từng phần; bạn xem lại từng đợt trong trình xem
- **Điền form** — hoàn thành theo từng field có hướng dẫn với xem trước trực tiếp
- **Ký tài liệu** — đặt ảnh chữ ký/ký nháy lên trang
- **Đóng dấu phê duyệt** — APPROVED, DRAFT, CONFIDENTIAL, hoặc bất kỳ nhãn tuỳ chỉnh nào
- **Tải về** — xuất PDF đã chú thích từ thanh công cụ của trình xem

## Lệnh

| Lệnh | Tính năng |
|---------|-------------|
| `/pdf-viewer:open` | Mở một PDF trong trình xem tương tác |
| `/pdf-viewer:annotate` | Đi qua tài liệu, đề xuất + áp dụng đánh dấu, cùng nhau xem lại |
| `/pdf-viewer:fill-form` | Điền các field của form PDF một cách tương tác |
| `/pdf-viewer:sign` | Đặt một ảnh chữ ký hoặc ký nháy lên trang |

## Khi nào nên dùng plugin này thay vì chỉ đọc PDF

Plugin này dành cho các **luồng công việc tương tác, trực quan** — khi bạn muốn
nhìn thấy tài liệu, đánh dấu lên nó, và tải về một bản đã chú thích.

Nếu bạn chỉ muốn Claude **tóm tắt hoặc trích xuất văn bản** từ một PDF,
đừng dùng plugin này. Claude có thể đọc file PDF một cách tự nhiên và điều đó
nhanh hơn cho việc nạp dữ liệu thuần tuý.

## Cách hoạt động

Plugin này sử dụng một **server MCP cục bộ** (`@modelcontextprotocol/server-pdf`)
chạy trên máy của bạn thông qua `npx`. Không cần API key hay dịch vụ từ xa nào —
server PDF tự động khởi động khi plugin được nạp.

## Yêu cầu

- Node.js >= 18
- Internet cho các PDF từ xa (arXiv, v.v.)

## Nguồn PDF được hỗ trợ

- File cục bộ (đường dẫn file trong thư mục làm việc của bạn)
- [arXiv](https://arxiv.org) — các URL `/abs/` tự động chuyển đổi sang PDF
- Bất kỳ URL PDF qua HTTPS trực tiếp nào (bioRxiv, Zenodo, OSF, v.v. — dùng đường
  link PDF, không phải trang đích)

## Miễn trừ trách nhiệm về chữ ký

`/pdf-viewer:sign` đặt một ảnh chữ ký **trực quan** lên trang. Đây không phải
là một chữ ký số được chứng thực hay có tính mật mã. Với chữ ký điện tử
có giá trị ràng buộc pháp lý, hãy dùng một dịch vụ ký chuyên dụng.

---

## Tác giả

**Nguyễn Ngọc Tuấn**
Founder Transform Group — **Lark Platinum Partner**
🌐 Dự án: [larkcowork.com](https://larkcowork.com)
