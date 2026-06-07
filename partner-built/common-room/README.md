# Common Room Plugin

Các quy trình GTM (go-to-market) được vận hành bởi Common Room — nghiên cứu tài khoản (account), nghiên cứu liên hệ (contact), chuẩn bị cuộc gọi, tiếp cận cá nhân hóa, tìm kiếm khách hàng tiềm năng và bản tóm tắt hàng tuần.

## Tổng quan

Plugin này kết nối Claude với MCP server của Common Room và trang bị cho nó sáu kỹ năng bao quát những quy trình phổ biến nhất của một nhân viên sales. Mọi đầu ra đều được đặt nền tảng trên dữ liệu tín hiệu (signal) thực tế của Common Room — tín hiệu sản phẩm bên thứ nhất (1st-party), tín hiệu cộng đồng bên thứ hai (2nd-party), tín hiệu ý định bên thứ ba (3rd-party), cùng dữ liệu làm giàu từ RoomieAI và Spark.

## Yêu cầu

- **Common Room MCP** (`mcp.commonroom.io/mcp`) phải được kết nối và xác thực. Đây là nguồn dữ liệu chính cho toàn bộ chức năng của plugin.
- **Connector lịch (Calendar)** (tùy chọn) — cho phép tự động tra cứu cuộc họp trong `call-prep` và `weekly-prep-brief`. Nếu không kết nối, cả hai kỹ năng sẽ hỏi người dùng chi tiết cuộc họp thay thế.

## Kỹ năng

Các kỹ năng được kích hoạt theo cách hội thoại. Hãy mô tả điều bạn muốn và Claude sẽ tự động tải đúng kỹ năng.

| Kỹ năng | Cụm từ kích hoạt |
|-------|----------------|
| `account-research` | "Research [company]", "tell me about [domain]", "what's going on with [account]", "is [company] showing buying signals" |
| `contact-research` | "Who is [name]", "look up [email]", "research [contact]", "is [name] a warm lead" |
| `call-prep` | "Prep me for my call with [company]", "prepare for a meeting with [company]", "what should I know before talking to [company]" |
| `compose-outreach` | "Draft outreach to [person]", "write an email to [name]", "compose a message for [contact]" |
| `prospect` | "Find companies that match [criteria]", "build a prospect list", "find contacts at [type of company]" |
| `weekly-prep-brief` | "Weekly prep brief", "prepare my week", "what calls do I have this week" |

## Lệnh

Hai lệnh dành cho các quy trình phức tạp được hưởng lợi từ việc gọi tường minh:

| Lệnh | Cách dùng |
|---------|-------|
| `/generate-account-plan <company>` | Kế hoạch tài khoản chiến lược toàn diện với bản đồ các bên liên quan (stakeholder mapping), phân tích mức độ tương tác, cơ hội, rủi ro và các hạng mục hành động |
| `/weekly-brief [date range]` | Tạo bản tóm tắt chuẩn bị hàng tuần đầy đủ (mặc định là 7 ngày tới) |

## Mỗi kỹ năng tạo ra điều gì

**Account Research (Nghiên cứu tài khoản)** — Xử lý bốn dạng tình huống: tổng quan đầy đủ, câu hỏi nhắm vào trường dữ liệu cụ thể, phản hồi trung thực khi dữ liệu thưa thớt, và kết hợp dữ liệu MCP + suy luận của LLM. Bao gồm tìm kiếm web cho tin tức gần đây. Tự động giới hạn phạm vi vào "My Segments".

**Contact Research (Nghiên cứu liên hệ)** — Tra cứu theo email, tên + công ty, hoặc tài khoản mạng xã hội. Trả về danh tính đã được làm giàu, các trường CRM, điểm số, lượt truy cập website, lịch sử hoạt động, các phân tích của Spark và các chủ đề mở đầu cuộc trò chuyện.

**Call Prep (Chuẩn bị cuộc gọi)** — Ảnh chụp nhanh công ty, hồ sơ từng người tham dự, các điểm tín hiệu nổi bật, các luận điểm trao đổi được tùy chỉnh, những phản đối có khả năng xảy ra và kết quả cuộc gọi được khuyến nghị. Ưu tiên hoạt động ghi âm cuộc gọi/Gong. Nhận biết được lịch nếu đã kết nối.

**Compose Outreach (Soạn nội dung tiếp cận)** — Ba định dạng cá nhân hóa (email, kịch bản cuộc gọi, tin nhắn LinkedIn) được đặt nền tảng trên tín hiệu của Common Room và các điểm móc nối từ tìm kiếm web. Được tùy chỉnh theo định vị công ty của người dùng khi có sẵn.

**Prospecting (Tìm kiếm khách hàng tiềm năng)** — Phân biệt giữa các công ty hoàn toàn mới (ProspectorOrganization) và các tài khoản hiện có (Organization). Hỗ trợ tinh chỉnh lặp đi lặp lại và tìm kiếm tương tự (lookalike) ("find companies like [X]"). Tìm kiếm web làm giàu cho kết quả công ty hoàn toàn mới.

**Weekly Prep Brief (Tóm tắt chuẩn bị hàng tuần)** — Bản tóm tắt đầy đủ bao quát mọi cuộc gọi với bên ngoài trong 7 ngày tới: ảnh chụp nhanh công ty, hồ sơ người tham dự, các tín hiệu và mục tiêu được khuyến nghị cho từng cuộc họp.

## Thiết lập

1. Đảm bảo Common Room MCP server đã được kết nối và xác thực trong phần cài đặt Cowork của bạn.
2. (Tùy chọn) Kết nối một MCP server lịch để tự động tra cứu cuộc họp trong phần chuẩn bị cuộc gọi và tóm tắt hàng tuần.
3. Cài đặt plugin này. Tất cả kỹ năng và lệnh đều có sẵn ngay lập tức.

## Ngữ cảnh người dùng

Tất cả các kỹ năng giới hạn phạm vi vào khu vực phụ trách (territory) của người dùng đều tự động lấy đối tượng `Me` từ Common Room. Điều này cung cấp hồ sơ, vai trò và "My Segments" của người dùng — đảm bảo các truy vấn mặc định giới hạn trong khu vực phụ trách của họ. Xem `references/me-context.md` để biết chi tiết.

Khi có sẵn ngữ cảnh công ty, các kỹ năng sẽ tùy chỉnh khuyến nghị theo sản phẩm và ICP của người dùng. Xem `references/my-company-context.md` để biết chi tiết.

## Tùy chỉnh

Xem `CONNECTORS.md` để biết chi tiết về connector lịch và cách các tham chiếu công cụ (tool reference) hoạt động.

---

## Tác giả

**Nguyễn Ngọc Tuấn**
Founder Transform Group — **Lark Platinum Partner**
🌐 Dự án: [larkcowork.com](https://larkcowork.com)
