# Plugin Quản lý Sản phẩm (Product Management)

Plugin quản lý sản phẩm được thiết kế chủ yếu cho [Cowork](https://claude.com/product/cowork), ứng dụng desktop dạng agent của Anthropic — dù vậy nó cũng hoạt động tốt trong Claude Code. Bao quát toàn bộ luồng công việc của PM: viết feature spec, quản lý roadmap, trao đổi với các bên liên quan, tổng hợp nghiên cứu người dùng, phân tích đối thủ và theo dõi chỉ số sản phẩm.

## Cài đặt

```
claude plugins add knowledge-work-plugins/product-management
```

## Tính năng

Plugin này mang đến cho bạn một cộng sự quản lý sản phẩm được hỗ trợ bởi AI, có thể giúp bạn với:

- **Feature Spec & PRD** — Tạo tài liệu yêu cầu sản phẩm có cấu trúc từ một phát biểu vấn đề hoặc một ý tưởng tính năng. Bao gồm user story, ưu tiên hoá yêu cầu, chỉ số thành công và quản lý phạm vi.
- **Lập kế hoạch Roadmap** — Tạo, cập nhật và sắp xếp lại ưu tiên cho roadmap sản phẩm của bạn. Hỗ trợ định dạng Now/Next/Later, các chủ đề theo quý, và định dạng gắn với OKR cùng bản đồ phụ thuộc.
- **Cập nhật cho các bên liên quan** — Tạo các cập nhật trạng thái phù hợp với đối tượng của bạn (lãnh đạo, kỹ thuật, khách hàng). Lấy ngữ cảnh từ các công cụ đã kết nối để bạn khỏi phải vật lộn với việc cập nhật hàng tuần.
- **Tổng hợp nghiên cứu người dùng** — Biến ghi chú phỏng vấn, dữ liệu khảo sát và support ticket thành các insight có cấu trúc. Nhận diện chủ đề, xây dựng persona và làm nổi bật các vùng cơ hội với bằng chứng hỗ trợ.
- **Phân tích đối thủ** — Nghiên cứu đối thủ và tạo các bản brief với so sánh tính năng, phân tích định vị và hàm ý chiến lược.
- **Đánh giá chỉ số (Metrics Review)** — Phân tích chỉ số sản phẩm, nhận diện xu hướng, so sánh với mục tiêu và làm nổi bật các insight có thể hành động.
- **Brainstorm sản phẩm** — Khám phá không gian vấn đề, tạo ý tưởng và thử thách tư duy sản phẩm cùng một đối tác phản biện sắc bén. Hỗ trợ ý tưởng phân kỳ, kiểm tra giả định và khám phá chiến lược bằng các framework như How Might We, Jobs-to-be-Done, First Principles và Opportunity Solution Trees.

## Lệnh

| Lệnh | Tính năng |
|---|---|
| `/write-spec` | Viết một feature spec hoặc PRD từ một phát biểu vấn đề |
| `/roadmap-update` | Cập nhật, tạo hoặc sắp xếp lại ưu tiên cho roadmap của bạn |
| `/stakeholder-update` | Tạo một cập nhật cho các bên liên quan (hàng tuần, hàng tháng, ra mắt) |
| `/synthesize-research` | Tổng hợp nghiên cứu người dùng từ phỏng vấn, khảo sát và ticket |
| `/competitive-brief` | Tạo một bản brief phân tích đối thủ |
| `/metrics-review` | Xem lại và phân tích chỉ số sản phẩm |
| `/brainstorm` | Brainstorm một ý tưởng sản phẩm, không gian vấn đề hoặc câu hỏi chiến lược cùng một đối tác tư duy |

## Kỹ năng

| Kỹ năng | Nội dung bao quát |
|---|---|
| `feature-spec` | Cấu trúc PRD, user story, phân loại yêu cầu, tiêu chí nghiệm thu |
| `roadmap-management` | Các framework ưu tiên hoá (RICE, MoSCoW), định dạng roadmap, bản đồ phụ thuộc |
| `stakeholder-comms` | Mẫu cập nhật theo đối tượng, truyền thông rủi ro, tài liệu hoá quyết định |
| `user-research-synthesis` | Phân tích theo chủ đề, affinity mapping, phát triển persona, định lượng cơ hội |
| `competitive-analysis` | Ma trận so sánh tính năng, phân tích định vị, phân tích win/loss |
| `metrics-tracking` | Phân cấp chỉ số sản phẩm, đặt mục tiêu (OKR), thiết kế dashboard, nhịp đánh giá |
| `product-brainstorming` | Các chế độ brainstorm (khám phá vấn đề, ý tưởng giải pháp, kiểm tra giả định, chiến lược), các framework PM, cấu trúc phiên làm việc, hành vi của một đối tác tư duy |

## Quy trình mẫu

### Viết một PRD

```
You: /write-spec
Claude: What feature or problem are you speccing out?
You: We need to add SSO support for enterprise customers
Claude: [Asks about target users, constraints, success metrics]
Claude: [Generates full PRD with problem statement, user stories, requirements, success metrics, open questions]
```

### Chuẩn bị một cập nhật cho các bên liên quan

```
You: /stakeholder-update
Claude: What type of update? (weekly, monthly, launch, ad-hoc)
You: Weekly update for the exec team
Claude: [Pulls context from project management tool, chat, and docs]
Claude: [Generates executive summary with progress, decisions, risks, and next milestones]
```

### Tổng hợp nghiên cứu người dùng

```
You: /synthesize-research
Claude: What research do you want to synthesize? You can paste interview notes, upload files, or I can pull from connected sources.
You: [Pastes 8 interview transcripts]
Claude: [Identifies themes, patterns, and insights across interviews]
Claude: [Generates synthesis with key findings, personas, opportunity areas, and evidence]
```

### Brainstorm một ý tưởng sản phẩm

```
You: /brainstorm Should we add AI-powered search to our product?
Claude: Before we jump to AI search — what problem are your users hitting with search today?
You: They can't find things. Support tickets about "can't find X" are our #2 category.
Claude: That's a discovery problem, not necessarily a search technology problem. Let me push on this...
Claude: [Explores whether the issue is search quality, information architecture, or content discoverability]
Claude: [Generates 5 distinct approaches ranging from better taxonomy to AI search to guided navigation]
Claude: [Challenges assumptions, suggests cheapest experiment to test the riskiest one]
```

### Phân tích đối thủ

```
You: /competitive-brief
Claude: Which competitor(s) or feature area do you want to analyze?
You: Compare our onboarding flow to [competitor A] and [competitor B]
Claude: [Researches competitor onboarding approaches]
Claude: [Generates brief with feature comparison, strengths/weaknesses, and strategic implications]
```

## Nguồn dữ liệu

> Nếu bạn thấy các placeholder lạ hoặc cần kiểm tra công cụ nào đang được kết nối, hãy xem [CONNECTORS.md](CONNECTORS.md).

Kết nối các công cụ quản lý dự án và giao tiếp của bạn để có trải nghiệm tốt nhất. Nếu không có chúng, hãy cung cấp ngữ cảnh thủ công.

**Các kết nối MCP đi kèm:**
- Chat (Lark IM) cho ngữ cảnh team và các luồng trao đổi với các bên liên quan
- Project tracker (Lark Task, Lark Task, Lark Base, Lark Task, Lark) để tích hợp roadmap, ngữ cảnh ticket và theo dõi trạng thái
- Cơ sở tri thức (Lark Wiki) cho các spec, nghiên cứu và ghi chú họp có sẵn
- Thiết kế (Figma) cho ngữ cảnh thiết kế và bàn giao
- Phân tích sản phẩm (Amplitude, Pendo) cho dữ liệu sử dụng, chỉ số và phân tích hành vi
- Phản hồi người dùng (Intercom) cho support ticket, yêu cầu tính năng và hội thoại với người dùng
- Phiên âm cuộc họp (Lark Minutes) cho ghi chú họp và ngữ cảnh thảo luận

**Lựa chọn bổ sung:**
- Xem [CONNECTORS.md](CONNECTORS.md) để biết các công cụ thay thế trong mỗi danh mục

---

## Tác giả

**Nguyễn Ngọc Tuấn**
Founder Transform Group — **Lark Platinum Partner**
🌐 Dự án: [larkcowork.com](https://larkcowork.com)
