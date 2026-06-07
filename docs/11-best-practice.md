# 11. Best practice — Khách nên custom & dùng thế nào

> Đúc kết cách triển khai và sử dụng để đạt giá trị nhanh, an toàn, bền. Dành cho cả **người dùng**
> và **admin/IT**.

## 11.1. Lộ trình triển khai 4 bước (cho tổ chức)

```
   B1 PILOT          B2 CHUẨN HÓA        B3 MỞ RỘNG          B4 TỰ ĐỘNG HÓA
   1–2 người    →    chốt quy trình  →   theo phòng ban  →   bot + lịch chạy
   stdio, vài KN     custom SKILL/        cài plugin đúng     base automation,
   "thắng nhanh"     settings/connectors  nghề, HTTP nếu cần   nhắc việc định kỳ
```

1. **Pilot nhỏ** — 1–2 người quyền lực (CEO/trưởng phòng), chế độ **stdio**, 3–5 kỹ năng "thắng
   nhanh" (morning-brief, inbox-zero, meeting-prep). Mục tiêu: thấy ROI trong tuần đầu.
2. **Chuẩn hóa** — tùy biến `settings.local.json`, thêm trigger tiếng Việt, chỉnh mẫu kết quả theo
   văn hóa công ty (xem [10. Tùy biến](./10-tuy-bien.md)).
3. **Mở rộng theo phòng ban** — cài đúng plugin cho từng nghề ([bảng dưới](#114-chọn-plugin-theo-vai-trò)),
   chuyển sang **HTTP** nếu dùng Cowork nhiều người.
4. **Tự động hóa** — dựng Base + automation (`base-deploy`, `base-automation`), bot nhắc việc, brief
   định kỳ.

## 11.2. Tư thế an toàn — quy tắc vàng

- ✅ **Luôn xem bản nháp** trước khi gửi ra ngoài (mail khách, tin nhắn đối tác). Email mặc định
  *lưu nháp*.
- ✅ **Bắt đầu read-only.** Cho trợ lý đọc/tóm tắt/soạn nháp trước; mở quyền ghi khi đã tin.
- ✅ **Một người một token.** Mỗi nhân viên đăng nhập bằng quyền của chính mình → không rò chéo dữ liệu.
- ✅ **Bật audit log** (`--audit-log`) ở môi trường nhiều người; đặt **bearer token** khi mở HTTP;
  chỉ bind nội bộ trừ khi cố ý phơi qua tunnel.
- ✅ **Bot là trợ lý cá nhân riêng tư** trước — cân nhắc kỹ trước khi thả vào group chung.
- ⚠️ **Đừng tin "code 0".** Tính năng UI chỉ tính là xong khi **xem thật** trên giao diện hoặc
  `./tools/audit.sh --live` xanh.

## 11.3. Viết câu lệnh tốt (prompting tiếng Việt)

| Nên | Thay vì |
|---|---|
| Nêu **bối cảnh + kết quả mong muốn** | Câu cụt thiếu ngữ cảnh |
| 💬 *"Tóm tắt cuộc họp với khách Acme hôm qua, rút action item, giao cho đúng người, để nháp mail cảm ơn."* | 💬 *"tóm tắt họp"* |
| Nêu **định dạng** muốn nhận (thẻ, bảng, gạch đầu dòng, file) | Để mặc định rồi sửa nhiều lần |
| Nói **"để nháp cho tôi xem"** khi chưa muốn gửi | Để trợ lý tự gửi |
| Dùng **`/lệnh`** khi muốn chạy đúng quy trình | Mô tả vòng vo |
| Cho biết **phạm vi/thời gian** (tuần này, top 5, nhóm X) | Để mở → kết quả lan man |

> 💡 Không nhớ làm được gì? Gõ 💬 *"bạn giúp được tôi những gì?"* hoặc `/` để xem danh sách lệnh.

## 11.4. Chọn plugin theo vai trò

| Vai trò | Plugin nên cài trước | Kỹ năng "thắng nhanh" |
|---|---|---|
| **Lãnh đạo / quản lý** | daily-assistant, governance | `morning-brief` (exec), `approval-triage`, `decision-logger` |
| **Sales / kinh doanh** | crm-sales, sales, daily-assistant | `pipeline-review`, `deal-update`, `client-followup`, `call-prep` |
| **CSKH** | customer-support, knowledge-docs | `ticket-triage`, `draft-response`, `kb-article` |
| **Nhân sự** | human-resources, knowledge-docs | `draft-offer`, `onboarding`, `policy-lookup` |
| **Kế toán / tài chính** | finance, governance | `reconciliation`, `variance-analysis`, `close-management` |
| **Vận hành** | operations, governance, lark-base-deploy | `process-doc`, `runbook`, `approval-flow-sla` |
| **Marketing** | marketing, knowledge-docs | `content-creation`, `campaign-plan`, `email-sequence` |
| **Sản phẩm / dự án** | product-management, daily-assistant, delivery-eng | `write-spec`, `sprint-planning`, `sprint-retro` |
| **Kỹ thuật** | engineering, delivery-eng, lark-cli-dev | `standup`, `code-review`, `incident-retro` |
| **Dữ liệu** | data | `write-query`, `build-dashboard`, `analyze` |
| **Pháp chế** | legal, governance | `review-contract`, `triage-nda`, `permission-audit` |
| **Chủ shop / SMB** | small-business, crm-sales | `monday-brief`, `plan-payroll`, `invoice-chase` |
| **Mọi người** | daily-assistant, enterprise-search, productivity | `morning-brief`, `search`, `task-management` |

> Plugin **nền tảng** ai cũng nên có: `daily-assistant` (vận hành ngày), `productivity` (việc + bộ
> nhớ), `enterprise-search` (tìm mọi nơi). `governance` cho ai có quyền duyệt.

## 11.5. Tận dụng Fusion (kết hợp plugin)

Các kỹ năng **gọi lại nhau** khi cài cùng nhau → cài theo cụm để mạnh hơn:

- daily-assistant (`morning-brief` bản exec) **+** governance → brief có cả hàng đợi duyệt.
- daily-assistant (`meeting-prep`) **+** governance (`decision-logger`) → họp xong tự lưu quyết định.
- crm-sales (`deal-update`) **+** daily-assistant (`contact-360`) → cập nhật deal có hồ sơ 360°.
- Thiếu plugin đồng hành cũng **không lỗi** — bước đó giảm cấp êm (bỏ qua/gợi ý).

Bản đồ kết hợp: [`../connectors/LARK-FUSION.md`](../connectors/LARK-FUSION.md).

## 11.6. Khi nào nên tự động hóa bằng Base

Chuyển từ "hỏi từng lần" sang "hệ thống" khi:
- Dữ liệu cần **dùng chung & lặp lại** (khách hàng, dự án, chấm công, kho).
- Cần **dashboard** theo dõi và **nhắc việc tự động**.

→ Dùng `base-deploy` (8 phase) dựng Base hoàn chỉnh + `base-automation` cho bot nhắc việc. Xem
[README lark-base-deploy](../lark-base-deploy/README.md).

## 11.7. Sai lầm thường gặp (và cách tránh)

| Sai lầm | Hệ quả | Cách tránh |
|---|---|---|
| Mở quyền ghi/bot toàn công ty ngay từ đầu | Rủi ro dữ liệu, mất kiểm soát | Pilot read-only trước, mở dần |
| Cài hết mọi plugin một lúc | Rối, khó đo giá trị | Cài theo vai trò, vài kỹ năng "thắng nhanh" |
| Tin "code 0" là xong | Kết quả sai vẫn tưởng đúng | Xem thật trên UI / `audit.sh --live` |
| Đổi `name:` trong SKILL.md | Gãy tham chiếu fusion | Chỉ sửa phần thân + description |
| Để trợ lý tự gửi mail khách | Gửi nhầm | Luôn "để nháp cho tôi xem" |
| Câu lệnh thiếu ngữ cảnh | Kết quả lan man | Nêu bối cảnh + định dạng + phạm vi |

## 11.8. Checklist "sẵn sàng dùng tốt"

- [ ] Đã `lark-cli auth login`, `lark-cli mcp tools` liệt kê đủ công cụ.
- [ ] Đã chọn transport (stdio/http) phù hợp.
- [ ] Đã cài đúng plugin theo vai trò ([11.4](#114-chọn-plugin-theo-vai-trò)).
- [ ] Đã điền `settings.local.json` cho plugin liên quan.
- [ ] (productivity) đã `/start` và bồi đắp bộ nhớ.
- [ ] Đã thử 3 kỹ năng "thắng nhanh" và **xem kết quả thật**.
- [ ] Môi trường nhiều người: bật audit log + bearer token.
- [ ] `./tools/audit.sh --live` xanh.

---

⬅️ Quay lại [mục lục tài liệu](./README.md) · Xem lại [8. Tất cả plugin](./08-tat-ca-plugin.md) ·
[9. Cách vận hành](./09-cach-van-hanh.md) · [10. Tùy biến](./10-tuy-bien.md).

---

## Tác giả

**Nguyễn Ngọc Tuấn**
Founder Transform Group — **Lark Platinum Partner**
🌐 Dự án: [larkcowork.com](https://larkcowork.com)
