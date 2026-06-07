# Brand Voice Plugin

Một plugin của [Tribe AI](https://tribe.ai) dành cho Claude Cowork. Được xây dựng với tư cách đối tác ra mắt (launch partner) của Cowork.

Kiến thức thương hiệu — thứ làm cho một công ty trở nên dễ nhận diện — hiếm khi nằm ở nơi nào hữu ích. Nó nằm trong một bộ slide từ năm 2022, một trang Lark Wiki không ai cập nhật kể từ lần tái định vị thương hiệu gần nhất, và trong bản năng của một vài người kỳ cựu đã gắn bó đủ lâu để "tự khắc biết". Khi đội sales đang dùng AI để tạo nội dung tiếp cận và nhân viên mới đang sản xuất nội dung ngay trong tuần đầu tiên, đó chính là thứ bị đánh mất.

Brand Voice biến những tài liệu thương hiệu rời rạc thành các rào chắn (guardrail) AI có thể thực thi được. Nó tìm kiếm xuyên suốt Lark Wiki, Lark Drive, Box, SharePoint, Lark IM, Gong và Granola để khám phá cách công ty bạn thực sự giao tiếp — sau đó tạo ra bộ nguyên tắc thương hiệu sẵn sàng cho LLM và kiểm chứng mọi nội dung do AI tạo ra dựa trên bộ nguyên tắc đó. Claude không chỉ viết nhanh hơn. Nó viết đúng giọng của bạn.

## Tính năng

### 1. Khám phá thương hiệu (Brand Discovery)
Kiến thức thương hiệu của bạn nằm rải rác khắp Lark Wiki, Lark Wiki, Lark Drive, Gong, Lark IM, và hàng năm trời các cuộc gọi sales lẫn bản ghi cuộc họp. Brand Voice tìm kiếm xuyên suốt tất cả — cẩm nang phong cách (style guide), bộ slide pitch, mẫu email, bản ghi (transcript), hệ thống thiết kế — để chắt lọc những tín hiệu thương hiệu mạnh nhất của bạn thành một nguồn chân lý (source of truth) duy nhất và cập nhật. Được đặt nền tảng trên cách những người giỏi nhất của bạn thực sự giao tiếp, chứ không phải cách mà một cẩm nang phong cách từ ba năm trước nói rằng bạn nên giao tiếp.

**Lệnh slash:** `/brand-voice:discover-brand`

```
/brand-voice:discover-brand
/brand-voice:discover-brand Acme Corp
```

### 2. Tạo bộ nguyên tắc (Guideline Generation)
Tổng hợp các tài liệu của bạn thành bộ nguyên tắc sẵn sàng cho LLM: các trụ cột giọng điệu (voice pillars), tham số tông giọng (tone parameters), một khung "We Are / We Are Not" giúp Claude có ranh giới vận hành rõ ràng, và các tiêu chuẩn thuật ngữ phản ánh ngôn ngữ thực tế của công ty — chứ không phải nội dung mang tính lý tưởng hóa. Cùng những rào chắn giúp các đội ngũ kỳ cựu giữ đúng thương hiệu cũng giúp nhân viên mới tạo ra nội dung chất lượng ngay trong tuần đầu thay vì phải đợi đến tháng thứ ba.

**Lệnh slash:** `/brand-voice:generate-guidelines`

```
/brand-voice:generate-guidelines
/brand-voice:generate-guidelines from the discovery report and these 3 PDFs
```

### 3. Thực thi giọng thương hiệu (Brand Voice Enforcement)
Mọi nội dung do AI tạo ra — email sales, đề xuất (proposal), trang marketing, thông cáo báo chí — đều được viết theo bộ nguyên tắc của bạn ngay từ đầu. Giọng điệu (voice) giữ nguyên trong khi tông giọng (tone) linh hoạt theo ngữ cảnh: mức độ trang trọng, năng lượng và chiều sâu kỹ thuật tự động thích ứng cho email tiếp cận lạnh (cold email) so với đề xuất doanh nghiệp so với bài đăng LinkedIn. Sự lệch tông và lỗ hổng định vị được phát hiện trước khi chúng đến tay khách hàng tiềm năng hay nhà đầu tư.

**Lệnh slash:** `/brand-voice:enforce-voice`

```
/brand-voice:enforce-voice Draft a cold email to a VP of Sales at a mid-market SaaS company
/brand-voice:enforce-voice Write a LinkedIn post announcing our new feature
```

### Câu hỏi mở (Open Questions)
Khi plugin gặp những điểm mơ hồ mà nó không thể tự giải quyết — tài liệu mâu thuẫn, thiếu nguyên tắc, sự khác biệt giữa thương hiệu được tuyên bố và thương hiệu được thực hành — nó sẽ nêu lên các câu hỏi mở để đội ngũ thảo luận. Mỗi câu hỏi đều kèm theo một khuyến nghị từ agent, biến sự mơ hồ thành một tương tác kiểu "xác nhận hoặc ghi đè" thay vì một ngõ cụt.

## MCP Connectors

| Connector | URL | Mục đích |
|-----------|-----|---------|
| **Lark Wiki** | `https://mcp.notion.com/mcp` | Xương sống cho khám phá — liên kết xuyên suốt Lark Drive, SharePoint, OneDrive, Lark IM, Lark Task đã kết nối. Cũng lưu trữ bộ nguyên tắc đầu ra. |
| **Lark** | `https://mcp.atlassian.com/v1/mcp` | Tìm kiếm sâu trong Lark Wiki + ngữ cảnh Lark Task cho các doanh nghiệp dùng nhiều Lark |
| **Box** | `https://mcp.box.com` | Lưu trữ file đám mây — tài liệu thương hiệu chính thức, bộ slide chia sẻ và cẩm nang phong cách thường nằm ở đây |
| **Lark** | `https://microsoft365.mcp.claude.com/mcp` | SharePoint, OneDrive, Outlook, Teams — lưu trữ tài liệu doanh nghiệp và mẫu email |
| **Figma** | `https://mcp.figma.com/mcp` | Hệ thống thiết kế thương hiệu — màu sắc, kiểu chữ (typography), design token định hình giọng điệu |
| **Gong** | `https://mcp.gong.io/mcp` | Trí tuệ hội thoại doanh nghiệp — bản ghi và phân tích các cuộc gọi sales |
| **Granola** | `https://mcp.granola.ai/mcp` | Trí tuệ cuộc họp — bản ghi và ghi chú từ các cuộc họp sales, khách hàng và chiến lược |

### Tích hợp gốc (Native Integrations)

Các nền tảng này là tích hợp gốc (native) của Claude — không cần cài MCP connector. Chúng có sẵn dưới dạng công cụ khi người dùng kết nối chúng trong Claude Desktop hoặc Cowork.

| Tích hợp | Mục đích |
|-------------|---------|
| **Lark Drive** | Tài liệu thương hiệu chia sẻ, cẩm nang phong cách, tài liệu marketing, Google Docs và Slides |
| **Lark IM** | Các thảo luận về thương hiệu, tìm kiếm trong kênh, nguyên tắc thương hiệu được ghim, các mẫu giọng điệu không chính thức |

## Bắt đầu nhanh

1. Cài plugin và mở Claude Cowork
2. Kết nối ít nhất một nền tảng (khuyến nghị Lark Wiki — nó liên kết xuyên suốt Lark Drive, SharePoint, Lark IM và Lark Task)
3. Chạy `/brand-voice:discover-brand` — Claude tự động tìm kiếm các tài liệu thương hiệu trong những cơ sở tri thức đã kết nối của bạn
4. Chạy `/brand-voice:generate-guidelines` để tạo ra một bộ nguyên tắc bền vững từ báo cáo khám phá
5. Dùng `/brand-voice:enforce-voice` khi tạo nội dung — email sales, đề xuất, bài đăng LinkedIn, bất cứ thứ gì hướng đến khách hàng

Bạn cũng có thể trỏ Claude đến những tài liệu cụ thể nếu muốn. Dù theo cách nào, nó cũng sẽ dẫn dắt bạn qua từng bước của quy trình.

Hiện tại Brand Voice hoạt động ở cấp độ cá nhân — khả năng thực thi trên toàn đội (team-wide) sẽ sớm ra mắt.

### Cấu hình theo từng dự án (Per-Project Settings)

Sao chép `settings/brand-voice.local.md.example` thành `.claude/brand-voice.local.md` trong dự án của bạn và điền tên công ty, các nền tảng đã bật, cùng các vị trí tài liệu thương hiệu đã biết.

## Cấu trúc file

```
├── .claude-plugin/
│   └── plugin.json                              # Plugin manifest
├── .mcp.json                                    # 7 MCP server connections
├── README.md
├── agents/
│   ├── discover-brand.md                        # Autonomous platform search agent
│   ├── content-generation.md                    # Brand-aligned content creation
│   ├── conversation-analysis.md                 # Sales call transcript analysis
│   ├── document-analysis.md                     # Brand document parsing
│   └── quality-assurance.md                     # Compliance and open questions audit
├── commands/
│   ├── discover-brand.md                        # /brand-voice:discover-brand
│   ├── enforce-voice.md                         # /brand-voice:enforce-voice
│   └── generate-guidelines.md                   # /brand-voice:generate-guidelines
├── settings/
│   └── brand-voice.local.md.example             # Per-project settings template
└── skills/
    ├── discover-brand/
    │   ├── SKILL.md                             # Discovery orchestration
    │   └── references/
    │       ├── search-strategies.md             # Platform-specific query patterns
    │       └── source-ranking.md                # Ranking algorithm and categories
    ├── brand-voice-enforcement/
    │   ├── SKILL.md                             # Enforcement orchestration
    │   └── references/
    │       ├── before-after-examples.md         # Content type transformation examples
    │       └── voice-constant-tone-flexes.md    # "We Are / We Are Not" + tone matrix
    └── guideline-generation/
        ├── SKILL.md                             # Generation orchestration
        └── references/
            ├── confidence-scoring.md            # Scoring methodology
            └── guideline-template.md            # Full output template
```

## Kiến trúc

**Kỹ năng (Skills)** cung cấp kiến thức chuyên ngành và điều phối các quy trình. Chúng tự động kích hoạt dựa trên ý định của người dùng.

**Agents** đảm nhận những công việc tự động nặng — tìm kiếm các nền tảng, phân tích tài liệu, bóc tách bản ghi, tạo nội dung và kiểm chứng chất lượng.

**Lệnh (Commands)** là các điểm khởi đầu tường minh do người dùng kích hoạt, nhằm khởi chạy các quy trình của kỹ năng.

**Các quyết định thiết kế then chốt:**
- Giọng điệu giữ nguyên, tông giọng linh hoạt — một mô hình tư duy rõ ràng cho việc thực thi
- Agent khám phá vận hành tự động nhưng có trách nhiệm giải trình — thể hiện rõ quá trình làm việc kèm nguồn gốc (provenance) và các mâu thuẫn
- Câu hỏi mở là một tính năng, không phải một thất bại — mỗi điểm mơ hồ đều kèm theo một khuyến nghị
- Tiết lộ tăng tiến (progressive disclosure) — phần frontmatter gọn nhẹ, SKILL.md tập trung, chi tiết nằm trong references/
- Lark Wiki AI Search làm cỗ máy khám phá liên kết (federated) — một API tìm kiếm trên hơn 8 nền tảng thông qua các nguồn đã kết nối
- Lark Drive và Lark IM là tích hợp gốc của Claude — không cần MCP connector

---

## Tác giả

**Nguyễn Ngọc Tuấn**
Founder Transform Group — **Lark Platinum Partner**
🌐 Dự án: [larkcowork.com](https://larkcowork.com)
