# Plugin Thiết kế (Design)

Plugin nâng cao năng suất thiết kế được thiết kế chủ yếu cho [Cowork](https://claude.com/product/cowork), ứng dụng desktop dạng agent của Anthropic — nhưng cũng hoạt động được trong Claude Code. Plugin hỗ trợ phản biện thiết kế (design critique), quản lý hệ thống thiết kế, viết UX, khả năng tiếp cận (accessibility), tổng hợp nghiên cứu và bàn giao cho lập trình viên (developer handoff). Hoạt động với mọi đội thiết kế — dùng độc lập với đầu vào của bạn, và tăng sức mạnh khi bạn kết nối Figma cùng các công cụ khác.

## Cài đặt

```bash
claude plugins add knowledge-work-plugins/design
```

## Lệnh

Các luồng công việc rõ ràng mà bạn gọi bằng lệnh slash:

| Lệnh | Mô tả |
|---|---|
| `/critique` | Nhận phản hồi thiết kế có cấu trúc — khả năng sử dụng (usability), phân cấp thị giác, khả năng tiếp cận và tính nhất quán |
| `/design-system` | Rà soát, lập tài liệu hoặc mở rộng hệ thống thiết kế của bạn — component, token, pattern |
| `/handoff` | Tạo đặc tả bàn giao cho lập trình viên — số đo, token, trạng thái, tương tác và các trường hợp biên (edge case) |
| `/ux-copy` | Viết hoặc rà soát nội dung UX — microcopy, thông báo lỗi, trạng thái rỗng (empty state), luồng onboarding |
| `/accessibility` | Chạy kiểm toán khả năng tiếp cận — tuân thủ WCAG, độ tương phản màu, trình đọc màn hình và điều hướng bàn phím |
| `/research-synthesis` | Tổng hợp nghiên cứu người dùng — phỏng vấn, khảo sát, kiểm thử khả năng sử dụng thành những insight có thể hành động |

Mọi lệnh đều hoạt động **độc lập** (mô tả thiết kế của bạn hoặc dán ảnh chụp màn hình) và được **tăng sức mạnh** khi có các connector MCP.

## Kỹ năng

Tri thức chuyên môn mà Claude tự động sử dụng khi phù hợp:

| Kỹ năng | Mô tả |
|---|---|
| `design-critique` | Đánh giá thiết kế về khả năng sử dụng, phân cấp thị giác, tính nhất quán và sự tuân thủ các nguyên tắc thiết kế |
| `design-system-management` | Quản lý design token, thư viện component và tài liệu pattern |
| `ux-writing` | Viết microcopy hiệu quả — rõ ràng, súc tích, nhất quán và đúng với thương hiệu |
| `accessibility-review` | Kiểm toán thiết kế và mã nguồn để tuân thủ WCAG 2.1 AA |
| `user-research` | Lập kế hoạch, thực hiện và tổng hợp nghiên cứu người dùng — phỏng vấn, khảo sát, kiểm thử khả năng sử dụng |
| `design-handoff` | Tạo tài liệu bàn giao toàn diện cho lập trình viên từ các thiết kế |

## Quy trình mẫu

### Nhận phản hồi thiết kế

```
/critique
```

Chia sẻ một link Figma, ảnh chụp màn hình, hoặc mô tả thiết kế của bạn. Nhận phản hồi có cấu trúc về khả năng sử dụng, phân cấp thị giác, tính nhất quán và khả năng tiếp cận.

### Kiểm toán hệ thống thiết kế của bạn

```
/design-system audit
```

Tôi sẽ rà soát thư viện component của bạn về tính nhất quán, độ đầy đủ và quy ước đặt tên. Nhận một báo cáo kèm các đề xuất cải thiện cụ thể.

### Viết nội dung UX

```
/ux-copy error messages for payment flow
```

Nhận nội dung phù hợp với ngữ cảnh kèm hướng dẫn giọng điệu, các phương án thay thế và ghi chú về bản địa hóa (localization).

### Bàn giao cho lập trình viên

```
/handoff
```

Chia sẻ một link Figma và nhận đặc tả đầy đủ: số đo, design token, trạng thái component, ghi chú tương tác và các trường hợp biên.

### Kiểm tra khả năng tiếp cận

```
/accessibility
```

Chia sẻ một thiết kế hoặc URL. Nhận báo cáo tuân thủ WCAG 2.1 AA kèm các vấn đề cụ thể, mức độ nghiêm trọng và các bước khắc phục.

### Tổng hợp nghiên cứu

```
/research-synthesis
```

Tải lên bản ghi phỏng vấn, kết quả khảo sát hoặc ghi chú kiểm thử khả năng sử dụng. Nhận các chủ đề, insight và đề xuất đã được xếp ưu tiên.

## Dùng độc lập + Tăng sức mạnh

Mọi lệnh và kỹ năng đều hoạt động mà không cần bất kỳ tích hợp nào:

| Bạn có thể làm gì | Độc lập | Tăng sức mạnh với |
|-----------------|------------|-------------------|
| Phản biện thiết kế | Mô tả hoặc chụp màn hình | Figma MCP (lấy thiết kế trực tiếp) |
| Hệ thống thiết kế | Mô tả hệ thống của bạn | Figma MCP (kiểm toán thư viện component) |
| Đặc tả bàn giao | Mô tả hoặc chụp màn hình | Figma MCP (số đo, token chính xác) |
| Nội dung UX | Mô tả ngữ cảnh | Cơ sở tri thức (hướng dẫn giọng thương hiệu) |
| Khả năng tiếp cận | Mô tả hoặc chụp màn hình | Figma MCP, analytics cho dữ liệu sử dụng thực tế |
| Tổng hợp nghiên cứu | Dán bản ghi | Công cụ phản hồi người dùng (lấy dữ liệu thô) |

## Tích hợp MCP

> Nếu bạn thấy các placeholder lạ hoặc cần kiểm tra công cụ nào đang được kết nối, xem [CONNECTORS.md](CONNECTORS.md).

Kết nối các công cụ của bạn để có trải nghiệm phong phú hơn:

| Danh mục | Ví dụ | Điều nó cho phép |
|---|---|---|
| **Công cụ thiết kế** | Figma | Lấy thiết kế, kiểm tra component, truy cập design token |
| **Phản hồi người dùng** | Intercom, Productboard | Phản hồi thô, yêu cầu tính năng, dữ liệu NPS |
| **Công cụ theo dõi dự án** | Lark Task, Lark Task, Lark Task | Liên kết thiết kế với ticket, theo dõi việc triển khai |
| **Cơ sở tri thức** | Lark Wiki | Hướng dẫn thương hiệu, nguyên tắc thiết kế, kho nghiên cứu |
| **Product analytics** | Amplitude, Mixpanel | Dữ liệu sử dụng cho tổng hợp nghiên cứu và quyết định thiết kế |

Xem [CONNECTORS.md](CONNECTORS.md) để biết danh sách đầy đủ các tích hợp được hỗ trợ.

---

## Tác giả

**Nguyễn Ngọc Tuấn**
Founder Transform Group — **Lark Platinum Partner**
🌐 Dự án: [larkcowork.com](https://larkcowork.com)
