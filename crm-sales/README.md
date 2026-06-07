# Plugin CRM & Sales

Vận hành một CRM gọn nhẹ trên Lark Base — hiển thị pipeline, cập nhật sau cuộc gọi, và tái kết nối các khách hàng đang ngủ đông — mà không cần rời khỏi Lark. Được thiết kế chủ yếu cho [Cowork](https://claude.com/product/cowork), ứng dụng desktop dạng agent của Anthropic, nhưng cũng hoạt động được trong Claude Code. Chỉ tạo bản nháp: Claude chuẩn bị mọi email gửi khách hàng, và bạn phê duyệt từng lần gửi.

## Cài đặt

```
claude plugins add lark-cowork/crm-sales
```

## Tính năng

Plugin này coi một bảng Lark Base là hệ thống ghi nhận chính (system-of-record) của bạn và xử lý các deal xoay quanh nó:

- **Pipeline review** — Một lượt quét hàng tuần gom các deal theo giai đoạn, đánh dấu các deal đang kẹt (không có tương tác trong 14 ngày+), nêu ra những deal sắp chốt (<30 ngày), và theo dõi xu hướng tỷ lệ thắng — để bạn bắt đầu thứ Hai với việc biết rõ cái gì cần chú ý mà không phải đọc cả CRM.
- **Deal update** — Dọn dẹp sau cuộc gọi: lấy biên bản cuộc họp, trích ra nỗi đau / ngân sách / dòng thời gian / bước tiếp theo, cập nhật bản ghi Base (bạn xác nhận phần thay đổi), rồi soạn email follow-up để bạn phê duyệt.
- **Client follow-up** — Phát hiện các liên hệ đang ngủ đông (>21 ngày không tương tác) và soạn email tái kết nối được cá nhân hóa, tham chiếu lịch sử thực tế — chỉ là bản nháp, không bao giờ tự động gửi.

## Kỹ năng

| Kỹ năng | Mô tả |
|-------|-------------|
| `pipeline-review` | Quét pipeline hàng tuần — theo giai đoạn, deal kẹt, sắp chốt, xu hướng tỷ lệ thắng |
| `deal-update` | Cập nhật CRM sau cuộc gọi — lấy biên bản, trích nỗi đau/ngân sách/dòng thời gian, cập nhật bản ghi Base, soạn follow-up |
| `client-followup` | Phát hiện liên hệ ngủ đông (>21 ngày) và soạn email tái kết nối được cá nhân hóa (chỉ bản nháp) |

## Quy trình mẫu

### Rà soát pipeline hàng tuần

```
Bạn: pipeline review

Claude: [Gom các deal theo giai đoạn kèm số lượng + giá trị $]
        [Đánh dấu 3 deal đang kẹt (>14 ngày không tương tác) và 2 deal sắp chốt (<30 ngày)]
        [Hiển thị tỷ lệ thắng tháng này so với tháng trước]
        [Khuyến nghị top 3 hành động cần tập trung trong tuần]
```

### Cập nhật sau một cuộc gọi

```
Bạn: just got off the call with Acme, update the deal

Claude: [Lấy biên bản cuộc họp với Acme]
        [Trích ra: nỗi đau (báo cáo thủ công), ngân sách ($40k), dòng thời gian (Q3)]
        [Hiển thị phần thay đổi của bản ghi Base: Stage → Negotiation, đặt Next step]
        [Khi xác nhận: cập nhật bản ghi, soạn email follow-up để rà soát]
```

### Tái kết nối các khách hàng ngủ đông

```
Bạn: who have I gone quiet on?

Claude: [Tìm 6 liên hệ không có tương tác trong 21+ ngày]
        [Mỗi khách: chủ đề gần nhất, số ngày ngủ đông, góc tiếp cận gợi ý]
        [Soạn một email tái kết nối được cá nhân hóa cho mỗi khách]
        [Nhắc: rà soát từng bản nháp trong Lark Mail trước khi gửi — CHƯA gửi]
```

## Plugin đồng hành

Các kỹ năng này thuộc bộ quy trình **lark-cowork** và tham chiếu đến các kỹ năng ở những plugin
anh em. Tên kỹ năng được phân giải toàn cục, nên một tham chiếu sẽ tự động hoạt động khi plugin
đồng hành được cài đặt; khi vắng plugin đồng hành, tham chiếu **suy giảm một cách mượt mà** (bước đó
bị bỏ qua hoặc được đưa ra như một gợi ý, không bao giờ báo lỗi). Hãy cài các plugin đồng hành dưới
đây để có trải nghiệm đầy đủ — xem [`../connectors/LARK-FUSION.md`](../connectors/LARK-FUSION.md).

| Kỹ năng của plugin này | Tham chiếu | Trong plugin |
|---|---|---|
| `pipeline-review` | `morning-brief` (biến thể sales) | daily-assistant |
| `deal-update` | `contact-360` (brief trước cuộc gọi) | daily-assistant |

## Nguồn dữ liệu

> Nếu bạn thấy các placeholder lạ hoặc cần kiểm tra công cụ nào đang được kết nối, xem [CONNECTORS.md](CONNECTORS.md).

**Kết nối MCP đi kèm:** server `lark` (`lark-cli mcp serve`), một cầu nối duy nhất bao phủ mọi hạng mục —

- **CRM** là một bảng Lark Base — đọc bằng `lark_base_search`, ghi qua `lark_api` (bitable records) hoặc kỹ năng `lark-base`.
- **Trí tuệ cuộc họp** (Lark Minutes) cung cấp dữ liệu cho việc cập nhật deal.
- **Email** (Lark Mail) cho follow-up và tái kết nối — chỉ bản nháp, bạn phê duyệt từng lần gửi.
- **Chat, Lịch, Danh bạ** (Lark IM / Calendar / Contact) cho bối cảnh và phân giải con người.

**Trạng thái chỉ-bản-nháp:** không có email gửi khách hàng nào từng được gửi tự động. Claude chuẩn bị bản nháp trong Lark Mail; bạn rà soát và gửi.

Xem [CONNECTORS.md](CONNECTORS.md) để biết bản đồ đầy đủ từ hạng mục đến công cụ và lõi chiều sâu Lark dùng chung.

---

## Tác giả

**Nguyễn Ngọc Tuấn**
Founder Transform Group — **Lark Platinum Partner**
🌐 Dự án: [larkcowork.com](https://larkcowork.com)
