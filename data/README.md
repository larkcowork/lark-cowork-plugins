# Plugin Phân tích Dữ liệu

Một plugin phân tích dữ liệu được thiết kế chủ yếu cho [Cowork](https://claude.com/product/cowork), ứng dụng desktop dạng agent của Anthropic — tuy nhiên cũng hoạt động được trong Claude Code. Truy vấn SQL, khám phá dữ liệu, trực quan hóa, dựng dashboard và tạo ra insight. Hoạt động với mọi kho dữ liệu (data warehouse), mọi phương ngữ SQL và mọi nền tảng phân tích.

## Cài đặt

```
claude plugins add knowledge-work-plugins/data
```

## Tính năng

Plugin này biến Claude thành một người cộng sự phân tích dữ liệu. Nó giúp bạn khám phá tập dữ liệu, viết SQL tối ưu, dựng trực quan hóa, tạo dashboard tương tác và kiểm định các phân tích trước khi chia sẻ với các bên liên quan.

### Khi có kết nối Data Warehouse

Hãy kết nối MCP server cho kho dữ liệu của bạn (ví dụ Snowflake, Databricks, BigQuery, hoặc bất kỳ cơ sở dữ liệu tương thích SQL nào) để có trải nghiệm tốt nhất. Khi đó Claude sẽ:

- Truy vấn trực tiếp vào kho dữ liệu của bạn
- Khám phá schema và metadata của các bảng
- Chạy phân tích trọn vẹn từ đầu đến cuối mà không cần copy-paste
- Lặp lại và tinh chỉnh truy vấn dựa trên kết quả

### Khi không có kết nối Data Warehouse

Khi không có kết nối tới kho dữ liệu, bạn có thể dán kết quả SQL hoặc tải lên file CSV/Excel để phân tích và trực quan hóa. Claude cũng có thể viết các câu lệnh SQL để bạn tự chạy thủ công, rồi sau đó phân tích kết quả bạn cung cấp.

## Lệnh

| Lệnh | Mô tả |
|---------|-------------|
| `/analyze` | Trả lời các câu hỏi về dữ liệu -- từ tra cứu nhanh đến phân tích đầy đủ |
| `/explore-data` | Phân tích đặc trưng (profile) và khám phá tập dữ liệu để hiểu hình dạng, chất lượng và các quy luật của nó |
| `/write-query` | Viết SQL tối ưu cho phương ngữ của bạn theo các thông lệ tốt nhất |
| `/create-viz` | Tạo trực quan hóa chất lượng xuất bản bằng Python |
| `/build-dashboard` | Dựng dashboard HTML tương tác với bộ lọc và biểu đồ |
| `/validate` | Kiểm định (QA) một phân tích trước khi chia sẻ -- kiểm tra phương pháp luận, độ chính xác và thiên kiến |

## Kỹ năng

| Kỹ năng | Mô tả |
|-------|-------------|
| `sql-queries` | Các thông lệ tốt nhất cho SQL trên nhiều phương ngữ, các mẫu phổ biến và tối ưu hiệu năng |
| `data-exploration` | Phân tích đặc trưng dữ liệu, đánh giá chất lượng và phát hiện quy luật |
| `data-visualization` | Lựa chọn biểu đồ, các mẫu code trực quan hóa bằng Python và nguyên tắc thiết kế |
| `statistical-analysis` | Thống kê mô tả, phân tích xu hướng, phát hiện điểm ngoại lai và kiểm định giả thuyết |
| `data-validation` | Kiểm định QA trước khi bàn giao, kiểm tra tính hợp lý và chuẩn tài liệu hóa |
| `interactive-dashboard-builder` | Xây dựng dashboard HTML/JS với Chart.js, bộ lọc và tạo kiểu (styling) |

## Quy trình mẫu

### Phân tích Ad-Hoc

```
You: /analyze What was our monthly revenue trend for the past 12 months, broken down by product line?

Claude: [Writes SQL query] → [Executes against data warehouse] → [Generates trend chart]
       → [Identifies key patterns: "Product line A grew 23% YoY while B was flat"]
       → [Validates results with sanity checks]
```

### Khám phá Dữ liệu

```
You: /explore-data users table

Claude: [Profiles table: 2.3M rows, 47 columns]
       → [Reports: created_at has 0.2% nulls, email has 99.8% cardinality]
       → [Flags: status column has unexpected value "UNKNOWN" in 340 rows]
       → [Suggests: "High-value dimensions to explore: plan_type, signup_source, country"]
```

### Viết Truy vấn

```
You: /write-query I need a cohort retention analysis -- users grouped by signup month,
     showing what % are still active 1, 3, 6, and 12 months later. We use Snowflake.

Claude: [Writes optimized Snowflake SQL with CTEs]
       → [Adds comments explaining each step]
       → [Includes performance notes about partition pruning]
```

### Dựng Dashboard

```
You: /build-dashboard Create a sales dashboard with monthly revenue, top products,
     and regional breakdown. Here's the data: [pastes CSV]

Claude: [Generates self-contained HTML file]
       → [Includes interactive Chart.js visualizations]
       → [Adds dropdown filters for region and time period]
       → [Opens in browser for review]
```

### Kiểm định Trước khi Chia sẻ

```
You: /validate [shares analysis document]

Claude: [Reviews methodology] → [Checks for survivorship bias in churn analysis]
       → [Verifies aggregation logic] → [Flags: "Denominator excludes trial users
          which could overstate conversion rate by ~5pp"]
       → [Confidence: "Ready to share with noted caveat"]
```

## Kết nối Nền tảng Dữ liệu của bạn

> Nếu bạn thấy các placeholder lạ hoặc cần kiểm tra xem những công cụ nào đang được kết nối, hãy xem [CONNECTORS.md](CONNECTORS.md).

Plugin này hoạt động tốt nhất khi được kết nối với hạ tầng dữ liệu của bạn. Hãy thêm các MCP server cho:

- **Kho dữ liệu (Data Warehouse)**: Snowflake, Databricks, BigQuery, Definite, hoặc bất kỳ cơ sở dữ liệu tương thích SQL nào
- **Phân tích/BI**: Amplitude, Looker, Tableau, hoặc tương tự
- **Notebook**: Jupyter, Hex, hoặc tương tự
- **Bảng tính**: Google Sheets, Excel
- **Điều phối dữ liệu (Data Orchestration)**: Airflow, dbt, Dagster, Prefect
- **Nạp dữ liệu (Data Ingestion)**: Fivetran, Airbyte, Stitch

Hãy cấu hình các MCP server trong file `.mcp.json` hoặc trong thiết lập của Claude Code để cho phép truy cập dữ liệu trực tiếp.

---

## Tác giả

**Nguyễn Ngọc Tuấn**
Founder Transform Group — **Lark Platinum Partner**
🌐 Dự án: [larkcowork.com](https://larkcowork.com)
