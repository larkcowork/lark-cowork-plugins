# Plugin Nhân sự (HR)

Plugin vận hành nhân sự được thiết kế chủ yếu cho [Cowork](https://claude.com/product/cowork) — ứng
dụng desktop agentic của Anthropic — nhưng cũng chạy được trong Claude Code. Hỗ trợ tuyển dụng,
onboarding, quản lý hiệu suất, tra cứu chính sách và phân tích lương thưởng. Dùng được với mọi đội
nhân sự — chạy độc lập với thông tin bạn cung cấp, và mạnh hơn nhiều khi bạn kết nối HRIS, ATS cùng
các công cụ khác.

## Cài đặt

```bash
claude plugins add knowledge-work-plugins/human-resources
```

## Lệnh

Các quy trình rõ ràng bạn gọi bằng lệnh slash:

| Lệnh | Mô tả |
|---|---|
| `/draft-offer` | Soạn thư mời nhận việc kèm chi tiết lương thưởng, ngày bắt đầu và điều khoản |
| `/onboarding` | Tạo checklist onboarding và kế hoạch tuần đầu cho nhân viên mới |
| `/performance-review` | Dựng khung đánh giá hiệu suất — câu hỏi tự đánh giá, mẫu cho quản lý, chuẩn bị calibration |
| `/policy-lookup` | Tìm và giải thích chính sách công ty — nghỉ phép, phúc lợi, chi phí, công tác, làm việc từ xa |
| `/comp-analysis` | Phân tích dữ liệu lương thưởng — đối chuẩn thị trường, xếp dải lương, mô hình refresh cổ phần |
| `/people-report` | Tạo báo cáo nhân sự, tỷ lệ nghỉ việc, đa dạng, hoặc sức khỏe tổ chức |

Mọi lệnh đều chạy **độc lập** (bạn cung cấp ngữ cảnh và chi tiết) và **mạnh hơn** khi có connector MCP.

## Kỹ năng

Kiến thức chuyên môn Claude tự dùng khi phù hợp:

| Kỹ năng | Mô tả |
|---|---|
| `recruiting-pipeline` | Theo dõi và quản lý pipeline tuyển dụng — các bước tìm nguồn, sàng lọc, phỏng vấn, đề nghị |
| `employee-handbook` | Trả lời câu hỏi về chính sách, phúc lợi và quy trình công ty |
| `compensation-benchmarking` | Đối chuẩn lương thưởng với dữ liệu thị trường — lương cơ bản, cổ phần, tổng thu nhập |
| `org-planning` | Hoạch định nhân sự, thiết kế tổ chức và tối ưu cấu trúc đội nhóm |
| `people-analytics` | Phân tích dữ liệu lực lượng lao động — xu hướng nghỉ việc, tín hiệu gắn kết, chỉ số đa dạng |
| `interview-prep` | Tạo kế hoạch phỏng vấn có cấu trúc — câu hỏi theo năng lực, scorecard, mẫu debrief |

## Quy trình mẫu

### Soạn thư mời nhận việc

```
/draft-offer
```

Cho tôi biết vị trí, cấp bậc, địa điểm và chi tiết lương thưởng. Bạn sẽ nhận một bản nháp thư mời hoàn
chỉnh kèm điều khoản, phân bổ cổ phần và tóm tắt phúc lợi.

### Onboarding nhân viên mới

```
/onboarding
```

Cung cấp tên, vị trí, đội nhóm và ngày bắt đầu của nhân viên mới. Bạn sẽ nhận checklist onboarding đầy
đủ, lịch tuần đầu, danh sách quyền truy cập công cụ và mẫu phân công người kèm cặp (buddy).

### Chuẩn bị đánh giá hiệu suất

```
/performance-review
```

Nhận các mẫu tự đánh giá, đánh giá của quản lý và calibration. Tôi sẽ giúp bạn dựng khung phản hồi cụ
thể, khả thi và công bằng.

### Hiểu rõ phúc lợi

Cứ hỏi tự nhiên:
```
Chính sách nghỉ thai sản của công ty mình thế nào?
```

Kỹ năng `employee-handbook` tự kích hoạt và tìm câu trả lời trong cơ sở tri thức đã kết nối của bạn.

### Đối chuẩn lương thưởng

```
/comp-analysis
```

Tải lên dữ liệu lương hoặc mô tả các dải lương của bạn. Nhận so sánh thị trường, phân tích xếp dải và
khuyến nghị điều chỉnh.

## Dùng độc lập + Tăng sức mạnh

Mọi lệnh và kỹ năng đều chạy được mà không cần tích hợp:

| Bạn làm được gì | Độc lập | Tăng sức mạnh với |
|-----------------|---------|-------------------|
| Soạn thư mời | Cung cấp chi tiết thủ công | HRIS, ATS để tự điền |
| Checklist onboarding | Mô tả quy trình của bạn | HRIS, Cơ sở tri thức để lấy mẫu |
| Đánh giá hiệu suất | Nhập thủ công | HRIS để lấy lịch sử đánh giá |
| Tra cứu chính sách | Dán nội dung sổ tay | Cơ sở tri thức |
| Phân tích lương | Tải CSV, mô tả dải lương | MCP dữ liệu lương thưởng |
| Báo cáo nhân sự | Tải dữ liệu lên | HRIS để lấy dữ liệu trực tiếp |

## Tích hợp MCP

> Nếu thấy placeholder lạ hoặc cần kiểm tra công cụ nào đang kết nối, xem [CONNECTORS.md](CONNECTORS.md).

Kết nối công cụ của bạn để có trải nghiệm phong phú hơn:

| Danh mục | Ví dụ | Mở khóa điều gì |
|---|---|---|
| **HRIS** | Workday, BambooHR, Rippling | Dữ liệu nhân viên, cấu trúc tổ chức, số dư nghỉ phép |
| **ATS** | Greenhouse, Lever, Ashby | Pipeline ứng viên, lịch phỏng vấn, theo dõi thư mời |
| **Lương thưởng** | Pave, Radford | Đối chuẩn thị trường, dữ liệu dải lương |
| **Chat** | Lark IM, Teams | Thông báo nội bộ, điều phối ứng viên |
| **Lịch** | Lark Calendar, Lark | Sắp lịch phỏng vấn, lịch onboarding |
| **Email** | Lark Mail, Lark | Thư mời nhận việc, trao đổi với ứng viên |

Xem [CONNECTORS.md](CONNECTORS.md) để biết danh sách đầy đủ các tích hợp được hỗ trợ.

## Cấu hình

Tạo file `settings.local.json` để cá nhân hóa:

- **Cowork**: Lưu trong bất kỳ thư mục nào bạn đã chia sẻ với Cowork (qua bộ chọn thư mục). Plugin tự
  tìm thấy.
- **Claude Code**: Lưu tại `human-resources/.claude/settings.local.json`.

```json
{
  "company": "Your Company",
  "headquarters": "City, State",
  "employeeCount": 500,
  "benefits": {
    "healthInsurance": "Provider Name",
    "pto": "Unlimited / X days",
    "parentalLeave": "X weeks"
  },
  "compensation": {
    "currency": "USD",
    "equityType": "RSU / Options",
    "vestingSchedule": "4 years, 1 year cliff"
  }
}
```

Plugin sẽ hỏi bạn các thông tin này một cách tương tác nếu chưa được cấu hình.

---

## Tác giả

**Nguyễn Ngọc Tuấn**
Founder Transform Group — **Lark Platinum Partner**
🌐 Dự án: [larkcowork.com](https://larkcowork.com)
