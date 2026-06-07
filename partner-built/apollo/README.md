# Apollo Plugin cho Claude Code và Cowork

Tìm kiếm khách hàng tiềm năng, làm giàu dữ liệu lead và nạp các chuỗi tiếp cận (outreach sequence) với [Apollo.io](https://www.apollo.io/) — được vận hành bởi Apollo MCP Server với khả năng **tích hợp chỉ bằng một cú nhấp chuột**.

---

## 🔌 Tích hợp MCP Server chỉ bằng một cú nhấp

Plugin này **tự động cấu hình Apollo MCP Server** khi cài đặt. Không cần thiết lập server thủ công, không cần chỉnh sửa file cấu hình — chỉ cần cài plugin và xác thực với tài khoản Apollo của bạn.

---

## ✅ Bộ kỹ năng mạnh mẽ

Plugin này đi kèm các kỹ năng giá trị cao, kết nối nhiều API của Apollo thành những quy trình hoàn chỉnh:

| Kỹ năng | Tính năng |
|---|---|
| `/apollo:enrich-lead` | Nhập một cái tên, URL LinkedIn hoặc email — nhận về thẻ liên hệ đầy đủ với email, số điện thoại, thông tin doanh nghiệp và các hành động tiếp theo |
| `/apollo:prospect` | Mô tả chân dung khách hàng lý tưởng (ICP) bằng ngôn ngữ tự nhiên — nhận về bảng lead người ra quyết định đã được làm giàu và xếp hạng |
| `/apollo:sequence-load` | Tìm lead, làm giàu dữ liệu và nạp hàng loạt vào một chuỗi tiếp cận — tự xử lý khử trùng lặp và đăng ký vào chuỗi |

### `/apollo:enrich-lead`

Nhập vào một cái tên, tên công ty, URL LinkedIn hoặc email — nhận về thẻ liên hệ hoàn chỉnh với email, số điện thoại, chức danh, vị trí, thông tin doanh nghiệp và các hành động tiếp theo được gợi ý. Hỗ trợ tra cứu mờ (ví dụ: "CEO of Figma") và tự động chuyển sang tìm kiếm khi không khớp chính xác.

### `/apollo:prospect`

Mô tả chân dung khách hàng lý tưởng (ICP) của bạn bằng ngôn ngữ tự nhiên. Quy trình sẽ tìm các công ty phù hợp, làm giàu hàng loạt dữ liệu firmographic, tìm những người ra quyết định, hé lộ thông tin liên hệ thông qua làm giàu hàng loạt, rồi trả về bảng lead được xếp hạng kèm điểm phù hợp ICP.

### `/apollo:sequence-load`

Tìm các liên hệ khớp với tiêu chí nhắm mục tiêu của bạn, làm giàu dữ liệu, tạo họ thành liên hệ với cơ chế khử trùng lặp, và thêm hàng loạt vào một chuỗi Apollo hiện có. Xem trước danh sách ứng viên trước khi đăng ký và hiển thị bản tóm tắt đầy đủ sau khi hoàn tất.

---

## 📦 Cài đặt

### Cowork

Nhấp vào liên kết bên dưới để cài đặt chỉ trong một bước:

[Install in Cowork](https://claude.ai/desktop/customize/plugins/new?marketplace=apolloio/apollo-mcp-plugin&plugin=apollo)

Sau đó khởi động lại Cowork để đảm bảo MCP server khởi chạy đúng cách.

### Claude Code

#### 1. Thêm marketplace của plugin này

Trong Claude Code, chạy:

```
/plugin marketplace add apolloio/apollo-mcp-plugin
```

#### 2. Cài đặt plugin

```
/plugin install apollo@apollo-plugin-marketplace
```

#### 3. Khởi động lại Claude Code

Điều này đảm bảo MCP server khởi chạy đúng cách.

---

## 🔑 Xác thực

Apollo MCP Server hỗ trợ **OAuth**:

1. Sau khi cài đặt, chạy `/mcp` trong Claude Code hoặc Cowork
2. Chọn server **Apollo** và nhấp **Authenticate**
3. Hoàn tất đăng nhập Apollo.io trong trình duyệt của bạn
4. Xong — tất cả các lệnh đã sẵn sàng để sử dụng

---

## ⚠️ Tín dụng Apollo (Apollo Credits)

Một số thao tác sẽ tiêu tốn [tín dụng Apollo](https://docs.apollo.io/):

- **Làm giàu dữ liệu người (People enrichment)** (được dùng bởi cả ba kỹ năng) tốn 1 tín dụng mỗi người
- **Làm giàu hàng loạt (Bulk enrichment)** (`/apollo:prospect`, `/apollo:sequence-load`) tiêu tốn 1 tín dụng cho mỗi người trong lô
- Plugin sẽ luôn cảnh báo bạn trước khi tiêu tốn tín dụng

---

## 🙌 Ghi nhận đóng góp

- **MCP Server** bởi [Apollo.io](https://docs.apollo.io/)
- **Đặc tả Plugin (Plugin Specification)** bởi [Anthropic](https://docs.anthropic.com/)

---

## Giấy phép

MIT — xem [LICENSE](LICENSE) để biết chi tiết.

---

## Tác giả

**Nguyễn Ngọc Tuấn**
Founder Transform Group — **Lark Platinum Partner**
🌐 Dự án: [larkcowork.com](https://larkcowork.com)
