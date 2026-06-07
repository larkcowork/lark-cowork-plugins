# Plugin Tài chính & Kế toán

Một plugin tài chính và kế toán được thiết kế chủ yếu cho [Cowork](https://claude.com/product/cowork), ứng dụng desktop dạng agent của Anthropic — nhưng cũng hoạt động được trong Claude Code. Hỗ trợ khóa sổ cuối tháng, chuẩn bị bút toán, đối chiếu tài khoản, lập báo cáo tài chính, phân tích chênh lệch và hỗ trợ audit SOX.

> **Quan trọng**: Plugin này hỗ trợ các quy trình tài chính và kế toán nhưng không cung cấp tư vấn tài chính, thuế hay audit. Mọi kết quả đầu ra cần được rà soát bởi các chuyên gia tài chính có trình độ trước khi sử dụng trong báo cáo tài chính, hồ sơ nộp cho cơ quan quản lý hay tài liệu audit.

## Cài đặt

```bash
claude plugins add knowledge-work-plugins/finance
```

## Lệnh

| Lệnh | Mô tả |
|---------|-------------|
| `/journal-entry` | Chuẩn bị bút toán — tạo các bút toán dồn tích, tài sản cố định, chi phí trả trước, lương và doanh thu với nợ/có đúng chuẩn và chi tiết hỗ trợ |
| `/reconciliation` | Đối chiếu tài khoản — so sánh số dư GL với sổ chi tiết, ngân hàng hoặc số dư bên thứ ba và xác định các khoản chênh lệch cần đối chiếu |
| `/income-statement` | Lập báo cáo kết quả kinh doanh — tạo P&L với so sánh kỳ-trên-kỳ và phân tích chênh lệch |
| `/variance-analysis` | Phân tích chênh lệch/biến động — phân rã chênh lệch thành các yếu tố tác động kèm diễn giải bằng lời và phân tích waterfall |
| `/sox-testing` | Kiểm thử tuân thủ SOX — tạo mẫu chọn, giấy tờ làm việc kiểm thử và đánh giá kiểm soát |

## Kỹ năng

| Kỹ năng | Mô tả |
|-------|-------------|
| `journal-entry-prep` | Best practice chuẩn bị bút toán, các loại dồn tích chuẩn, yêu cầu tài liệu hỗ trợ và quy trình rà soát |
| `reconciliation` | Phương pháp đối chiếu GL-với-sổ chi tiết, đối chiếu ngân hàng và nội bộ tập đoàn, kèm phân loại và phân tích tuổi các khoản chênh lệch |
| `financial-statements` | Định dạng báo cáo kết quả kinh doanh, bảng cân đối kế toán và báo cáo lưu chuyển tiền tệ với cách trình bày theo GAAP và phương pháp phân tích biến động |
| `variance-analysis` | Kỹ thuật phân rã chênh lệch (giá/lượng, tỷ lệ/cơ cấu), ngưỡng trọng yếu, tạo diễn giải và biểu đồ waterfall |
| `close-management` | Checklist khóa sổ cuối tháng, trình tự công việc, các phụ thuộc, theo dõi trạng thái và các hoạt động khóa sổ thường gặp theo từng ngày |
| `audit-support` | Phương pháp kiểm thử kiểm soát SOX 404, chọn mẫu, chuẩn tài liệu và phân loại thiếu sót |

## Quy trình mẫu

### Khóa sổ cuối tháng

1. Chạy `/journal-entry ap-accrual 2024-12` để tạo các bút toán dồn tích phải trả
2. Chạy `/journal-entry prepaid 2024-12` để phân bổ chi phí trả trước
3. Chạy `/journal-entry fixed-assets 2024-12` để ghi nhận khấu hao
4. Chạy `/reconciliation cash 2024-12` để đối chiếu tài khoản ngân hàng
5. Chạy `/reconciliation accounts-receivable 2024-12` để đối chiếu sổ chi tiết phải thu
6. Chạy `/income-statement monthly 2024-12` để tạo P&L kèm phân tích biến động

### Phân tích chênh lệch

1. Chạy `/variance-analysis revenue 2024-Q4 vs 2024-Q3` để phân tích chênh lệch doanh thu
2. Chạy `/variance-analysis opex 2024-12 vs budget` để điều tra chênh lệch chi phí vận hành
3. Rà soát phân tích waterfall và cung cấp bối cảnh cho các chênh lệch chưa giải thích được

### Kiểm thử SOX

1. Chạy `/sox-testing revenue-recognition 2024-Q4` để tạo giấy tờ làm việc kiểm thử kiểm soát doanh thu
2. Chạy `/sox-testing procure-to-pay 2024-Q4` để kiểm thử các kiểm soát mua sắm
3. Rà soát mẫu chọn và ghi lại kết quả kiểm thử

## Tích hợp MCP

> Nếu bạn thấy các placeholder lạ hoặc cần kiểm tra công cụ nào đang được kết nối, xem [CONNECTORS.md](CONNECTORS.md).

Plugin này hoạt động tốt nhất khi được kết nối với các nguồn dữ liệu tài chính của bạn qua các server MCP. Thêm các server liên quan vào `.mcp.json` của bạn:

### Hệ thống ERP / Kế toán

Kết nối server MCP của ERP (ví dụ NetSuite, SAP) để tự động lấy bảng cân đối thử, dữ liệu sổ chi tiết và bút toán.

### Kho dữ liệu (Data Warehouse)

Kết nối server MCP của kho dữ liệu (ví dụ Snowflake, BigQuery) để truy vấn dữ liệu tài chính, chạy phân tích chênh lệch và lấy các so sánh lịch sử.

### Bảng tính (Spreadsheets)

Kết nối các công cụ bảng tính (ví dụ Google Sheets, Excel) để tạo giấy tờ làm việc, mẫu đối chiếu và cập nhật mô hình tài chính.

### Analytics / BI

Kết nối nền tảng BI của bạn (ví dụ Tableau, Looker) để lấy dashboard, KPI và dữ liệu xu hướng phục vụ giải thích chênh lệch.

> **Lưu ý:** Kết nối server MCP của ERP và kho dữ liệu để tự động lấy dữ liệu tài chính. Nếu không có chúng, bạn có thể dán dữ liệu hoặc tải lên file để phân tích.

## Cấu hình

Thêm các server MCP nguồn dữ liệu của bạn vào mục `mcpServers` trong `.mcp.json` ở thư mục plugin này. Trường `recommendedCategories` liệt kê các loại tích hợp giúp nâng cao năng lực của plugin này:

- `erp-accounting` — Hệ thống ERP hoặc kế toán cho dữ liệu GL, sổ chi tiết và bút toán
- `data-warehouse` — Kho dữ liệu cho các truy vấn tài chính và dữ liệu lịch sử
- `spreadsheets` — Công cụ bảng tính để tạo giấy tờ làm việc
- `analytics-bi` — Công cụ BI cho dashboard và dữ liệu KPI
- `documents` — Lưu trữ tài liệu cho chính sách, memo và chứng từ hỗ trợ
- `email` — Email để gửi báo cáo và yêu cầu phê duyệt
- `chat` — Giao tiếp nhóm để cập nhật trạng thái khóa sổ và đặt câu hỏi

---

## Tác giả

**Nguyễn Ngọc Tuấn**
Founder Transform Group — **Lark Platinum Partner**
🌐 Dự án: [larkcowork.com](https://larkcowork.com)
