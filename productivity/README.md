# Plugin Năng suất (Productivity)

Một plugin năng suất được thiết kế chủ yếu cho [Cowork](https://claude.com/product/cowork), ứng dụng desktop dạng agent của Anthropic — tuy nhiên cũng hoạt động được trong Claude Code. Quản lý công việc, bộ nhớ về nơi làm việc, và một dashboard trực quan — Claude học về con người, dự án và thuật ngữ của bạn để có thể hành động như một đồng nghiệp, chứ không phải một chatbot.

## Cài đặt

```
claude plugins add knowledge-work-plugins/productivity
```

## Tính năng

Plugin này mang lại cho Claude một sự hiểu biết bền vững về công việc của bạn:

- **Quản lý công việc** — Một danh sách công việc dạng markdown (`TASKS.md`) mà Claude đọc, ghi và thực thi dựa trên đó. Hãy thêm công việc một cách tự nhiên, và Claude sẽ theo dõi trạng thái, triage các mục tồn đọng, và đồng bộ với các công cụ bên ngoài.
- **Bộ nhớ về nơi làm việc** — Một hệ thống bộ nhớ hai tầng dạy cho Claude cách viết tắt, con người, dự án và thuật ngữ của bạn. Nói "ask todd to do the PSR for oracle" và Claude biết chính xác ai, việc gì, và thuộc deal nào.
- **Dashboard trực quan** — Một file HTML cục bộ cho bạn một góc nhìn dạng bảng (board) về các công việc và một góc nhìn trực tiếp về những gì Claude biết về nơi làm việc của bạn. Chỉnh sửa từ board hoặc từ file — chúng luôn được đồng bộ với nhau.

## Lệnh

| Lệnh | Nó làm gì |
|---------|--------------|
| `/start` | Khởi tạo công việc + bộ nhớ, mở dashboard |
| `/update` | Triage các mục tồn đọng, kiểm tra bộ nhớ xem có thiếu sót gì không, đồng bộ từ các công cụ bên ngoài nếu phù hợp |
| `/update --comprehensive` | Quét sâu email, lịch, chat — gắn cờ các todo bị bỏ sót và gợi ý các mục bộ nhớ mới |

## Kỹ năng

| Kỹ năng | Mô tả |
|-------|-------------|
| `memory-management` | Hệ thống bộ nhớ hai tầng — CLAUDE.md cho bộ nhớ làm việc, thư mục memory/ để lưu trữ sâu |
| `task-management` | Theo dõi công việc dựa trên markdown sử dụng một file TASKS.md dùng chung |

## Quy trình mẫu

### Bắt đầu

```
You: /start

Claude: [Creates TASKS.md, CLAUDE.md, memory/ directory, and dashboard.html]
        [Opens the dashboard in your browser]
        [Asks about your role, team, and current priorities to seed memory]
```

### Thêm Công việc một cách Tự nhiên

```
You: I need to review the budget proposal for Sarah by Friday,
     draft the Q2 roadmap after syncing with Greg, and follow up
     on the API spec from the Platform team

Claude: [Adds all three tasks to TASKS.md with context]
        [Dashboard updates automatically]
```

### Đồng bộ Buổi sáng

```
You: /update --comprehensive

Claude: [Scans email, calendar, and chat for new action items]
        [Flags: "Budget proposal review is due tomorrow — still open"]
        [Suggests: "New person mentioned in 3 threads: Jamie Park,
         Design Lead — add to memory?"]
        [Updates stale tasks and fills memory gaps]
```

### Cách viết tắt nơi làm việc

Một khi bộ nhớ đã được điền đầy đủ, Claude giải mã cách viết tắt của bạn ngay lập tức:

```
You: ask todd to do the PSR for oracle

Claude: "Ask Todd Martinez (Finance lead) to prepare the Pipeline
         Status Report for the Oracle Systems deal ($2.3M, closing Q2)"
```

Không cần hỏi lại. Không phải đi đi lại lại.

## Nguồn dữ liệu

> Nếu bạn thấy các placeholder lạ hoặc cần kiểm tra xem những công cụ nào đang được kết nối, hãy xem [CONNECTORS.md](CONNECTORS.md).

Hãy kết nối các công cụ giao tiếp và quản lý dự án của bạn để có trải nghiệm tốt nhất. Khi không có chúng, bạn vẫn có thể quản lý công việc và bộ nhớ một cách thủ công.

**Kết nối MCP đi kèm:** server `lark` (`lark-cli mcp serve`), một cầu nối duy nhất bao quát mọi danh mục —
- Chat (Lark IM) để lấy ngữ cảnh nhóm và quét tin nhắn
- Email và lịch (Lark Mail + Lark Calendar) để phát hiện các đầu việc cần làm
- Cơ sở tri thức (Lark Wiki + Docs) cho các tài liệu tham khảo
- Theo dõi dự án (Lark Task + Lark Base) để đồng bộ công việc
- Bộ ứng dụng văn phòng (Lark Docs/Sheets/Drive) cho tài liệu

**Các tùy chọn bổ sung:**
- Xem [CONNECTORS.md](CONNECTORS.md) để biết các công cụ thay thế trong từng danh mục

---

## Tác giả

**Nguyễn Ngọc Tuấn**
Founder Transform Group — **Lark Platinum Partner**
🌐 Dự án: [larkcowork.com](https://larkcowork.com)
