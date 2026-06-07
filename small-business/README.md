# Plugin Doanh nghiệp Nhỏ (Small Business)

Bộ quy trình dành cho doanh nghiệp nhỏ được dựng sẵn cho [Cowork](https://claude.com/product/cowork), ứng dụng desktop dạng agent của Anthropic — đồng thời hoạt động được trong Claude Code. Cài đặt một lần, bạn sẽ có 15 kỹ năng nền tảng, 15 quy trình dùng ngay, và một bộ định tuyến (router) hiểu được tiếng Anh đời thường.

Bạn không cần ghi nhớ bất cứ điều gì. Chỉ cần nói với Claude bạn cần gì — "Tôi đang lo không đủ tiền trả lương," "Có khách hàng đang giận," "Tôi nên định giá bao nhiêu?" — và nó sẽ tự xác định đúng quy trình rồi hướng dẫn bạn từng bước. Mọi quy trình đều dừng lại trước khi thực hiện hành động, nên không có gì xảy ra mà chưa được bạn đồng ý.

> **Quan trọng**: Plugin này hỗ trợ các quy trình của doanh nghiệp nhỏ nhưng không cung cấp tư vấn tài chính, thuế, pháp lý hay nhân sự. Mọi kết quả đầu ra cần được bạn (và khi cần thiết, một chuyên gia có chuyên môn) xem xét trước khi sử dụng.

## Cài đặt

### Cowork

Cài đặt từ [claude.com/plugins](https://claude.com/plugins/).

### Claude Code

```bash
claude plugin marketplace add anthropics/knowledge-work-plugins
claude plugin install small-business@knowledge-work-plugins
```

Sau khi cài xong, hãy nói **"set me up"** để chạy kỹ năng `smb-onboard` — nó sẽ giúp Claude hiểu về doanh nghiệp của bạn, các điểm khó khăn, và những công cụ bạn đang dùng.

## Những thứ bạn cần kết nối

Chạy `/smb-onboard` hoặc nói với Claude "set me up."

**Công cụ cốt lõi** (kết nối những thứ này trước để có trải nghiệm tốt nhất):
- **QuickBooks** — cung cấp dữ liệu cho mọi quy trình tài chính (dự báo dòng tiền, biên lợi nhuận, chốt sổ cuối tháng, chuẩn bị thuế)
- **PayPal** — dữ liệu giao dịch, hóa đơn, tranh chấp và hoàn tiền
- **HubSpot** — CRM, khách hàng tiềm năng, chiến dịch, và phiếu hỗ trợ khách hàng

**Marketing & giao tiếp:**
- **Canva** — tạo ra các ấn phẩm mạng xã hội và email đúng nhận diện thương hiệu
- **Lark Mail / Outlook** — soạn email, xử lý phiếu hỗ trợ, rà soát hợp đồng
- **Lark Calendar / Outlook Calendar** — chuẩn bị họp, chặn lịch gọi điện, các cam kết trong tuần
- **Lark IM** — gửi bản tóm tắt và thông báo

**Tùy chọn** (kết nối thêm để có chiều sâu hơn):
- **Stripe** — dữ liệu thanh toán và đăng ký dịch vụ (subscription)
- **Square** — dữ liệu giao dịch POS
- **Lark Drive / OneDrive** — lưu trữ tệp và mẫu (template)
- **DocuSign** — rà soát hợp đồng từ các phong bì (envelope) đang chờ ký

Bạn không cần tất cả những công cụ này để bắt đầu. Hãy kết nối một hoặc hai cái và bạn sẽ thấy ngay giá trị — plugin sẽ cho bạn biết khi nào việc kết nối thêm một công cụ sẽ mở khóa thêm nhiều khả năng.

## Cách hoạt động

Ba lớp phối hợp với nhau:

1. **Kỹ năng (Skills)** — các khối nền tảng. Mỗi kỹ năng biết cách làm thật tốt một việc (dự báo dòng tiền, chấm điểm khách hàng tiềm năng, soạn lời nhắc hóa đơn). Có 15 kỹ năng như vậy.

2. **Lệnh (Commands)** — các quy trình. Lệnh xâu chuỗi các kỹ năng lại thành công thức nhiều bước, với các điểm kiểm soát nơi bạn phê duyệt trước khi bất cứ điều gì xảy ra. Có 15 lệnh như vậy.

3. **Bộ định tuyến (Router)** — cánh cửa chính. Bạn nói chuyện với Claude bằng ngôn ngữ đời thường. Router lắng nghe, xác định quy trình nào phù hợp, và đưa bạn đến đó. Bạn không bao giờ cần phải nhớ tên một lệnh.

## Toàn bộ 15 lệnh

Lệnh là các quy trình xâu chuỗi các kỹ năng lại với nhau. Mỗi lệnh đều dừng tại các điểm kiểm soát để chờ bạn phê duyệt trước khi thực hiện hành động.

### Tiền & tài chính

| Lệnh | Tính năng | Chỉ cần nói... | Kỹ năng sử dụng | Bắt buộc | Tùy chọn |
|---|---|---|---|---|---|
| `/plan-payroll` | Dự báo dòng tiền + đòi hóa đơn quá hạn để bạn yên tâm đủ tiền trả lương. | "can I make payroll", "cash is tight", "who owes me money" | cash-flow-snapshot, invoice-chase | QuickBooks | PayPal, Stripe, Square, Mail |
| `/month-heads-up` | Triển vọng dòng tiền 30 ngày với cảnh báo rủi ro sớm. | "what does next month look like", "cash forecast", "runway" | cash-flow-snapshot | QuickBooks | PayPal |
| `/close-month` | Chốt sổ cuối tháng: đối soát, đánh dấu khoảng trống, viết P&L, xuất bộ hồ sơ. | "close the books", "month-end", "reconcile" | month-end-prep | QuickBooks | PayPal, Stripe, Square |
| `/price-check` | Bảng biên lợi nhuận theo từng sản phẩm và ba kịch bản định giá. | "what are my margins", "should I raise prices", "cost per unit" | margin-analyzer | QuickBooks | PayPal |
| `/tax-prep` | Tài liệu chuẩn bị thuế cho kế toán của bạn (ước tính theo quý hoặc 1099 cuối năm). | "tax stuff", "estimated taxes", "1099s", "accountant needs..." | tax-season-organizer | QuickBooks | PayPal, Stripe |

### Bán hàng & marketing

| Lệnh | Tính năng | Chỉ cần nói... | Kỹ năng sử dụng | Bắt buộc | Tùy chọn |
|---|---|---|---|---|---|
| `/call-list` | Top 5 khách hàng tiềm năng cần gọi hôm nay kèm ý chính trao đổi và lịch chặn thời gian. | "who should I call", "any hot leads", "pipeline" | lead-triage | HubSpot | Mail, Lark Calendar |
| `/run-campaign` | Chiến dịch trọn gói: phân tích bán hàng → bản tóm tắt nội dung → ấn phẩm Canva → gửi qua HubSpot. | "run a campaign", "sales are down", "I need more customers" | content-strategy, canva-creator, lead-triage | HubSpot, Canva | QuickBooks, PayPal |
| `/sales-brief` | Sản phẩm bán chạy nhất và kém nhất kèm bản tóm tắt nội dung 2 tuần. | "what's selling", "what should I promote" | content-strategy | QuickBooks hoặc PayPal | HubSpot |

### Khách hàng & vận hành

| Lệnh | Tính năng | Chỉ cần nói... | Kỹ năng sử dụng | Bắt buộc | Tùy chọn |
|---|---|---|---|---|---|
| `/customer-pulse-check` | Các chủ đề phản hồi của khách hàng kèm mẫu trả lời. | "what are customers saying", "complaints", "reviews" | customer-pulse, ticket-deflector | PayPal hoặc HubSpot | -- |
| `/handle-complaint` | Xử lý khiếu nại trọn gói: lấy bối cảnh, soạn câu trả lời, đề xuất cách khắc phục vận hành. | "a customer is upset", "handle this complaint", "angry email" | ticket-deflector, customer-pulse | -- (hoạt động được với văn bản dán vào) | Lark Mail, HubSpot, PayPal |
| `/crm-cleanup` | Làm sạch HubSpot: deal cũ, trùng lặp, thiếu trường — sửa những gì bạn phê duyệt. | "clean up the CRM", "HubSpot is a mess", "stale deals" | crm-maintenance | HubSpot | -- |
| `/review-contract` | Rà soát hợp đồng bằng ngôn ngữ dễ hiểu với các điểm cờ đỏ và mức độ nghiêm trọng. | "review this contract", "NDA", "should I sign this" | contract-review | -- (hoạt động được khi tải tệp lên) | DocuSign |

### Trí tuệ kinh doanh

| Lệnh | Tính năng | Chỉ cần nói... | Kỹ năng sử dụng | Bắt buộc | Tùy chọn |
|---|---|---|---|---|---|
| `/monday-brief` | Bản tóm tắt sáng thứ Hai: tiền mặt, doanh số, pipeline, tuần phía trước, top 3 việc cần làm. | "Monday brief", "what's on my plate", "start of week" | business-pulse | -- (giảm cấp linh hoạt) | QuickBooks, PayPal, HubSpot, Calendar, Lark Mail, Lark IM |
| `/friday-brief` | Nhịp đập cuối tuần thứ Sáu: doanh thu so với tuần trước, thành tích, và những điều cần theo dõi. | "end of week", "how'd we do", "Friday recap" | business-pulse | PayPal hoặc HubSpot | -- |
| `/quarterly-review` | Bản tường thuật QBR đầy đủ: doanh thu, biên lợi nhuận, sức khỏe khách hàng, cơ hội, rủi ro. | "quarterly review", "board deck", "QBR" | business-pulse | QuickBooks | PayPal, HubSpot |

## Toàn bộ 15 kỹ năng

Kỹ năng là các khối nền tảng nguyên tử. Mỗi kỹ năng làm thật tốt một việc.

### Tiền & tài chính

| Kỹ năng | Tính năng | Chỉ cần nói... | Bắt buộc | Tùy chọn |
|---|---|---|---|---|
| **cash-flow-snapshot** | Dự báo dòng tiền 30/60/90 ngày với dải tin cậy và cờ rủi ro được nêu tên. Tóm tắt trong chat + tệp XLSX. | "forecast my cash flow", "will I make payroll", "runway", "cash crunch" | QuickBooks, PayPal, Stripe, hoặc Square (bất kỳ một cái) | Các công cụ khác làm nguồn phụ |
| **invoice-chase** | Soạn lời nhắc hóa đơn quá hạn khớp với lịch sử thanh toán và giọng điệu của từng khách hàng. Gửi qua PayPal sau khi được phê duyệt. | "who owes me money", "overdue invoices", "follow up on unpaid" | QuickBooks | PayPal, Stripe, Lark Mail |
| **margin-analyzer** | Kinh tế đơn vị theo sản phẩm hoặc dịch vụ kèm chuẩn so sánh lạm phát và ba kịch bản định giá. | "what are my margins", "should I raise prices", "costs eating into profit", "what to charge" | QuickBooks | PayPal, Square, tải lên CSV |
| **month-end-prep** | Chốt sổ cuối tháng: đối soát QB với các cổng thanh toán, đánh dấu khoảng trống, viết bản tường thuật P&L, xuất bộ hồ sơ chốt sổ. | "close the month", "reconcile", "P&L", "why revenue changed" | QuickBooks | PayPal, Stripe, Square |
| **tax-season-organizer** | Tính thuế ước tính theo quý hoặc chuẩn bị 1099-NEC cuối năm kèm bộ hồ sơ bàn giao cho kế toán. | "quarterly taxes", "estimated tax payment", "1099s", "1099-NEC", "year-end tax prep" | QuickBooks | PayPal, Stripe |

### Bán hàng & marketing

| Kỹ năng | Tính năng | Chỉ cần nói... | Bắt buộc | Tùy chọn |
|---|---|---|---|---|
| **lead-triage** | Chấm điểm khách hàng tiềm năng trong HubSpot theo mức tương tác, độ phù hợp và độ cấp bách để tạo danh sách gọi đã xếp hạng kèm ý chính trao đổi. | "prioritize leads", "who to call first", "pipeline" | HubSpot | Lark Mail, Lark Calendar |
| **content-strategy** | Phân tích dữ liệu bán hàng để tìm sản phẩm bán chạy nhất và bán chậm, tạo bản tóm tắt nội dung 30 ngày được ưu tiên hóa. | "what should I post", "content plan", "what's selling", "what to promote" | QuickBooks hoặc PayPal | Square |
| **canva-creator** | Nhận bản tóm tắt nội dung và thực thi toàn bộ chiến dịch: lịch đăng bài, ấn phẩm Canva, lời chú thích, dàn dựng trong HubSpot. | "make the content", "generate the posts", "create the assets", "turn this into a campaign" | Canva, HubSpot | -- |

### Khách hàng & vận hành

| Kỹ năng | Tính năng | Chỉ cần nói... | Bắt buộc | Tùy chọn |
|---|---|---|---|---|
| **customer-pulse** | Tổng hợp tranh chấp, phiếu hỗ trợ, sắc thái email và đánh giá thành báo cáo theo chủ đề kèm danh sách "tuần này hãy làm ba việc này". | "how are customers feeling", "what people are saying", "disputes", "review analysis" | -- (giảm cấp linh hoạt) | PayPal, HubSpot, Lark Mail |
| **ticket-deflector** | Đọc email hoặc phiếu hỗ trợ của khách hàng, lấy trạng thái đơn hàng/hoàn tiền, soạn câu trả lời khớp giọng điệu. Có thể hoàn tiền qua PayPal sau khi được phê duyệt. | "draft a response", "answer this customer", "where's my order", "I want a refund" | PayPal, HubSpot, Mail | Intercom, Square |
| **crm-maintenance** | Giữ HubSpot luôn cập nhật: tạo/cập nhật liên hệ và deal, ghi nhật ký cuộc gọi và ghi chú, đánh dấu bản ghi cũ. | "update the CRM", "log a call", "clean up HubSpot", "add context to a deal" | HubSpot | Lark Mail, Lark Calendar |
| **contract-review** | Rà soát hợp đồng bằng ngôn ngữ dễ hiểu với cờ rủi ro, mức độ nghiêm trọng, và tệp DOCX đánh dấu sửa (redline). | "review this contract", "what am I signing", "flag any concerns", "check the payment terms" | -- (hoạt động được khi tải tệp lên) | Lark Mail, DocuSign |

### Tuyển dụng

| Kỹ năng | Tính năng | Chỉ cần nói... | Bắt buộc | Tùy chọn |
|---|---|---|---|---|
| **job-post-builder** | Xây dựng bộ hồ sơ tuyển dụng hoàn chỉnh: tin tuyển dụng, bộ hướng dẫn phỏng vấn có cấu trúc kèm thang điểm chấm, và mẫu thư mời nhận việc. | "help me hire", "write a job post", "job description", "open role", "interview questions", "draft an offer letter" | -- (hoạt động độc lập) | DocuSign, Lark Drive |

### Trí tuệ kinh doanh & khởi tạo

| Kỹ năng | Tính năng | Chỉ cần nói... | Bắt buộc | Tùy chọn |
|---|---|---|---|---|
| **business-pulse** | Bản tổng quan doanh nghiệp một trang: tiền mặt, doanh số, pipeline, cam kết, danh sách theo dõi, và việc quan trọng nhất duy nhất cần xử lý hôm nay. | "how's the business doing", "snapshot", "weekly summary", "catch me up" | -- (giảm cấp linh hoạt) | QuickBooks, PayPal, HubSpot, Lark Calendar, Lark Mail, Lark IM |
| **smb-onboard** | Hướng dẫn bạn kết nối công cụ, chạy một công thức demo, ghi nhận bối cảnh doanh nghiệp của bạn, và thiết lập nhịp kiểm tra hằng tuần. | "set me up", "setup", "get started", "help me get set up", "I'm new to this", "what can you do" | -- | Tất cả connector |

## Tùy chỉnh

Các quy trình này là điểm khởi đầu mang tính tổng quát. Chúng sẽ hữu ích hơn nhiều khi bạn tùy chỉnh cho đúng với cách doanh nghiệp của bạn thực sự vận hành:

- **Thêm bối cảnh doanh nghiệp** — Đưa ngành nghề, sản phẩm, khách hàng và quy trình của bạn vào các tệp kỹ năng để Claude hiểu được thế giới của bạn.
- **Điều chỉnh ngưỡng** — Tinh chỉnh các ngưỡng cảnh báo trong `business-pulse` và `cash-flow-snapshot` cho khớp với quy mô của bạn.
- **Thay đổi connector** — Trỏ các kỹ năng đến đúng những công cụ bạn thực sự dùng.

---

## Tác giả

**Nguyễn Ngọc Tuấn**
Founder Transform Group — **Lark Platinum Partner**
🌐 Dự án: [larkcowork.com](https://larkcowork.com)
