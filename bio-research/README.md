# Plugin Nghiên cứu Sinh học (Bio-Research)

Kết nối tới các công cụ và cơ sở dữ liệu nghiên cứu tiền lâm sàng (tìm kiếm tài liệu, phân tích genomics, ưu tiên hóa mục tiêu thuốc) để tăng tốc R&D khoa học đời sống giai đoạn đầu. Dùng với [Cowork](https://claude.com/product/cowork) hoặc cài đặt trực tiếp trong Claude Code.

Plugin này hợp nhất 11 tích hợp máy chủ MCP và 5 kỹ năng phân tích thành một gói duy nhất dành cho các nhà nghiên cứu khoa học đời sống.

## Những gì được bao gồm

### Máy chủ MCP (Nguồn dữ liệu & Công cụ)

> Nếu bạn thấy các placeholder lạ hoặc cần kiểm tra công cụ nào đang được kết nối, xem [CONNECTORS.md](CONNECTORS.md).

| Nhà cung cấp | Tính năng | Danh mục/Placeholder |
|----------|-------------|---------------------|
| U.S. National Library of Medicine | Tìm kiếm tài liệu y sinh và bài báo nghiên cứu | `~~literature` |
| deepsense.ai | Truy cập preprint từ bioRxiv và medRxiv | `~~literature` |
| Consensus | Tìm kiếm và tổng hợp nghiên cứu đã bình duyệt bằng AI | `~~literature` |
| John Wiley & Sons | Truy cập nghiên cứu học thuật và các ấn phẩm | `~~journal access` |
| Sage Bionetworks | Quản lý dữ liệu nghiên cứu cộng tác | `~~data repository` |
| deepsense.ai | Cơ sở dữ liệu hợp chất hoạt tính sinh học dạng thuốc | `~~chemical database` |
| OpenTargets | Khám phá và ưu tiên hóa mục tiêu thuốc | `~~drug targets` |
| deepsense.ai | Cơ quan đăng ký thử nghiệm lâm sàng NIH/NLM | `~~clinical trials` |
| BioRender | Tạo minh họa khoa học | `~~scientific illustration` |
| Owkin | AI cho sinh học — mô bệnh học và khám phá thuốc | `~~AI research` |
| Benchling\* | Nền tảng quản lý dữ liệu phòng thí nghiệm | `~~lab platform` |

### Máy chủ MCP dạng Binary tùy chọn

Những máy chủ này yêu cầu tải về một binary riêng:

- **10X Genomics txg-mcp** (`~~genomics platform`) — Dữ liệu và quy trình phân tích trên đám mây ([GitHub](https://github.com/10XGenomics/txg-mcp/releases))
- **ToolUniverse** (`~~tool database`) — Các công cụ AI cho khám phá khoa học từ Harvard MIMS ([GitHub](https://github.com/mims-harvard/ToolUniverse/releases))

### Kỹ năng (Quy trình phân tích)

#### Single-Cell RNA QC
Kiểm soát chất lượng tự động cho dữ liệu scRNA-seq theo các thực hành tốt nhất của scverse. Hỗ trợ tệp `.h5ad` và `.h5` với lọc dựa trên MAD và các biểu đồ trực quan toàn diện.

#### scvi-tools
Bộ công cụ học sâu cho omics đơn tế bào. Bao quát các mô hình scVI, scANVI, totalVI, PeakVI, MultiVI, DestVI, veloVI, và sysVI cho tích hợp, hiệu chỉnh batch, chuyển nhãn (label transfer), và phân tích đa phương thức.

#### Nextflow Pipelines
Chạy các pipeline tin sinh học nf-core trên dữ liệu giải trình tự cục bộ hoặc công khai từ GEO/SRA:
- **rnaseq** — Biểu hiện gene và biểu hiện khác biệt (differential expression)
- **sarek** — Phát hiện biến thể germline và somatic (WGS/WES)
- **atacseq** — Phân tích khả năng tiếp cận chromatin

#### Instrument Data to Allotrope
Chuyển đổi các tệp đầu ra của thiết bị phòng thí nghiệm (PDF, CSV, Excel, TXT) sang định dạng Allotrope Simple Model (ASM). Hỗ trợ hơn 40 loại thiết bị bao gồm máy đếm tế bào, máy quang phổ, máy đọc đĩa (plate reader), qPCR, và các hệ thống sắc ký.

#### Scientific Problem Selection
Khung làm việc có hệ thống để lựa chọn vấn đề nghiên cứu dựa trên khung của Fischbach & Walsh. Bao gồm 9 kỹ năng bao quát việc lên ý tưởng, đánh giá rủi ro, tối ưu hóa, cây quyết định, lập kế hoạch ứng phó nghịch cảnh, và tổng hợp.

## Bắt đầu

```bash
# Install the plugin
/install anthropics/knowledge-work-plugins bio-research

# Run the start command to see available tools
/start
```

## Quy trình thường gặp

**Tổng quan tài liệu (Literature Review)**
Tìm kiếm bài báo trong cơ sở dữ liệu ~~literature, truy cập toàn văn qua ~~journal access, và tạo hình ảnh minh họa với ~~scientific illustration.

**Phân tích đơn tế bào (Single-Cell Analysis)**
Chạy QC trên dữ liệu scRNA-seq, sau đó dùng scvi-tools để tích hợp, hiệu chỉnh batch, và chú thích loại tế bào (cell type annotation).

**Pipeline giải trình tự (Sequencing Pipeline)**
Tải dữ liệu công khai từ GEO/SRA, chạy các pipeline nf-core (RNA-seq, phát hiện biến thể, ATAC-seq), và kiểm tra đầu ra.

**Khám phá thuốc (Drug Discovery)**
Tìm kiếm hợp chất hoạt tính sinh học trong cơ sở dữ liệu ~~chemical database, dùng cơ sở dữ liệu ~~drug target để ưu tiên hóa mục tiêu, và xem xét dữ liệu thử nghiệm lâm sàng.

**Chiến lược nghiên cứu (Research Strategy)**
Trình bày một ý tưởng mới, gỡ rối một dự án bị mắc kẹt, hoặc đánh giá các quyết định chiến lược bằng khung lựa chọn vấn đề khoa học.

## Giấy phép

Các kỹ năng được cấp phép theo Apache 2.0. Các máy chủ MCP do tác giả tương ứng của chúng cung cấp — xem tài liệu của từng máy chủ riêng để biết điều khoản.

---

## Tác giả

**Nguyễn Ngọc Tuấn**
Founder Transform Group — **Lark Platinum Partner**
🌐 Dự án: [larkcowork.com](https://larkcowork.com)
