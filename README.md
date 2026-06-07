# Lark Cowork Plugins

Biến **Claude Cowork** (trong Claude Desktop) thành một **chuyên gia làm việc thuần Lark** cho mọi
vai trò — chat, mail, lịch, tài liệu/wiki, công việc, Base, drive, biên bản họp — tất cả đều chạy
qua cầu nối MCP `lark-cli`.

> 📘 **Bạn là người dùng nghiệp vụ (không chuyên kỹ thuật)? Bắt đầu tại đây →**
> [`docs/`](./docs/README.md) — bộ tài liệu tiếng Việt đầy đủ: giới thiệu, hướng dẫn bắt đầu,
> **danh mục theo phòng ban**, **giới thiệu chi tiết tất cả plugin** ([trang 8](./docs/08-tat-ca-plugin.md)),
> **cách vận hành** ([trang 9](./docs/09-cach-van-hanh.md)), **cách tùy biến/custom**
> ([trang 10](./docs/10-tuy-bien.md)), **best practice — nên dùng & custom sao**
> ([trang 11](./docs/11-best-practice.md)), kịch bản thực tế, an toàn dữ liệu, FAQ và thuật ngữ.

> Tùy biến từ bộ [knowledge-work-plugins](https://github.com/anthropics/knowledge-work-plugins) mã
> nguồn mở của Anthropic. Bộ gốc nhắm tới Slack / Notion / Google / Jira. Bản fork này **đấu lại
> toàn bộ lớp cộng tác sang Lark (larksuite.com)**, trong khi vẫn giữ các công cụ chuyên ngành thực
> sự (kho dữ liệu, CRM, thanh toán, cơ sở dữ liệu khoa học) dưới dạng connector ngoài tùy chọn.

---

## Mục lục

- [Ý tưởng cốt lõi](#ý-tưởng-cốt-lõi)
- [Cách hoạt động](#cách-hoạt-động)
- [Phạm vi bao phủ — 15 plugin nghiệp vụ](#phạm-vi-bao-phủ--15-plugin-nghiệp-vụ)
- [7 plugin thuần Lark (mới)](#7-plugin-thuần-lark-mới)
- [Yêu cầu cài đặt trước](#yêu-cầu-cài-đặt-trước)
- [Chọn phương thức kết nối (transport)](#chọn-phương-thức-kết-nối-transport)
- [Cài vào Claude](#cài-vào-claude)
- [Chạy lại quá trình chuyển đổi](#chạy-lại-quá-trình-chuyển-đổi)
- [Cấu trúc kho mã](#cấu-trúc-kho-mã)
- [Chất lượng & kiểm thử](#chất-lượng--kiểm-thử)
- [Ghi nhận & bản quyền](#ghi-nhận--bản-quyền)

---

## Ý tưởng cốt lõi

Lark là một **super-app**: Chat (IM), Mail, Lịch, Tài liệu (Docs), Wiki, **Base** (bảng dữ liệu
thông minh), Sheets, Drive, Minutes (biên bản họp AI), họp video, Phê duyệt, OKR… **gói gọn trong
một nền tảng**. Vì mọi mặt công việc đã nằm trong Lark, một trợ lý AI đặt ngay tại đây có thể chạm
tới **toàn bộ vòng đời công việc** — điều mà việc ghép nối rời rạc Slack + Notion + Jira + Gmail
không làm được.

Bộ plugin này dạy Claude nói "tiếng Lark": bạn ra lệnh bằng tiếng Việt như nói với đồng nghiệp,
Claude tự thao tác trên dữ liệu Lark của bạn rồi trả kết quả về — ngay tại chỗ, không cần rời Lark,
không cần biết code.

---

## Cách hoạt động

Mỗi plugin chỉ gồm **markdown + JSON**. Các kỹ năng (skill) đều **không phụ thuộc công cụ cụ thể** —
chúng tham chiếu tới các "ô giữ chỗ" dạng `~~category`, được phân giải qua file `CONNECTORS.md` của
từng plugin. So với bản gốc, mỗi plugin chỉ thay đổi đúng hai thứ:

1. **`.mcp.json`** — các server giao tiếp chung (Slack, Gmail, Google Calendar, Notion, Asana,
   Linear, Atlassian, Guru, Fireflies, Box…) được thay bằng **một server MCP `lark` duy nhất**
   (`lark-cli mcp serve`). Các server chuyên ngành được giữ nguyên.
2. **`CONNECTORS.md`** — mỗi nhóm danh mục chung nay được ánh xạ tới các công cụ `lark_*` cụ thể;
   các danh mục chuyên ngành vẫn giữ kết nối ngoài.

Một server `lark` duy nhất hỗ trợ mọi danh mục cộng tác. Bất kỳ thao tác nào không có công cụ chuyên
biệt sẽ rơi xuống `lark_api` (cửa thoát hiểm tới Lark OpenAPI). Xem bản ánh xạ tổng tại
[connectors/CONNECTORS.lark.md](./connectors/CONNECTORS.lark.md).

**Sơ đồ luồng:**

```
   Bạn ra lệnh           Trợ lý AI               Thao tác trên Lark        Trả kết quả
   trong Lark      →     dùng kỹ năng       →    (mail, lịch, Base,    →   về Lark
   (tiếng Việt)          chuyên Lark             tài liệu, duyệt…)         (thẻ, tin nhắn, file)
```

---

## Phạm vi bao phủ — 15 plugin nghiệp vụ

| Plugin | Danh mục chạy thuần Lark | Giữ kết nối ngoài (tùy chọn) |
|--------|--------------------------|------------------------------|
| productivity (năng suất) | chat, mail, lịch, wiki, task/base, drive | — |
| enterprise-search (tìm kiếm toàn doanh nghiệp) | chat, mail, wiki, task, drive | — |
| operations (vận hành) | chat, mail, lịch, wiki, task | itsm, mua sắm |
| human-resources (nhân sự) | chat, mail, lịch, wiki | ats, hris, dữ liệu lương |
| customer-support (CSKH) | chat, mail, wiki, task, drive | intercom, hubspot |
| sales (kinh doanh) | chat, mail, lịch, wiki, task, minutes | hubspot, close, clay, zoominfo, apollo, outreach, similarweb |
| product-management (quản lý sản phẩm) | chat, mail, lịch, wiki, task, minutes | figma, amplitude, pendo, intercom, similarweb |
| marketing | chat, wiki | canva, figma, hubspot, amplitude, ahrefs, similarweb, klaviyo, supermetrics |
| legal (pháp lý) | chat, mail, lịch, office, task, drive | docusign |
| design (thiết kế) | chat, wiki, task | figma, intercom |
| data (dữ liệu) | task | snowflake, databricks, bigquery, hex, amplitude, definite |
| finance (tài chính) | chat, mail, office | snowflake, databricks, bigquery |
| engineering (kỹ thuật) | chat, wiki, task | github, pagerduty, datadog |
| bio-research (nghiên cứu sinh học) | (giao tiếp qua lark) | pubmed, biorender, biorxiv, consensus, clinicaltrials, chembl, synapse, wiley, owkin, open-targets, benchling |
| small-business (doanh nghiệp nhỏ) | chat, mail, lịch, docs, drive | quickbooks, paypal, stripe, square, hubspot, canva, docusign |

`pdf-viewer` và `cowork-plugin-management` được giữ nguyên từ bản gốc.

---

## 7 plugin thuần Lark (mới)

Các plugin này được nâng cấp từ những kỹ năng `lark-cli` thô thành plugin hoàn chỉnh theo đúng quy
ước (`plugin.json` + `.mcp.json` lark + `CONNECTORS.md` + `README.md`). **Mọi danh mục đều phân giải
về một server MCP `lark` duy nhất — không server chuyên ngành, không phụ thuộc ngoài.**

| Plugin | Làm gì | Số kỹ năng |
|--------|--------|-----------|
| lark-base-deploy | Triển khai Lark Base end-to-end (orchestrator 8 phase, fan-out song song) | 9 |
| daily-assistant | Vận hành một ngày làm việc — brief, digest, triage, focus, chuẩn bị họp/1:1, contact 360 | 12 |
| crm-sales | CRM trên Lark Base — review pipeline, cập nhật deal sau gọi, follow-up khách im lặng (chỉ tạo nháp) | 3 |
| governance (quản trị) | Triage phê duyệt, đo SLA luồng duyệt, audit phân quyền, ghi nhật ký quyết định | 4 |
| knowledge-docs | Soạn docs/wiki/sheet từ template; audit & tái cấu trúc Wiki | 2 |
| delivery-eng | Postmortem sự cố không đổ lỗi + retro cuối sprint | 2 |
| lark-cli-dev | Xây/sửa/mở rộng cầu nối MCP `lark-cli` (công cụ dev; 1 kỹ năng + 6 lệnh `/mcp-*`) | 1 |

---

## Yêu cầu cài đặt trước

Connector `lark` gọi tới `lark-cli`. Cài đặt + xác thực một lần:

```bash
# 1. Build/cài lark-cli (từ kho lark-cli) vào PATH, ví dụ:
#    go build -o ~/bin/lark-cli . && export PATH="$HOME/bin:$PATH"
lark-cli --version

# 2. Xác thực (lưu thông tin đăng nhập vào macOS Keychain)
lark-cli auth login
```

---

## Chọn phương thức kết nối (transport)

Bộ plugin mặc định dùng profile **stdio**. Có thể đổi bất kỳ lúc nào:

```bash
./set-transport.sh stdio   # Claude Desktop cổ điển — binary lark-cli cục bộ + Keychain
./set-transport.sh http    # Cowork VM / từ xa — lark-cli mcp serve --transport http + tunnel
```

- **stdio** — đơn giản nhất. Cần `lark-cli` trên PATH và đã chạy `lark-cli auth login` trên máy này.
- **http** — dành cho môi trường sandbox của Cowork (không gọi được binary trên máy chủ). Chạy một
  cầu nối từ xa rồi trỏ plugin tới đó:

  ```bash
  lark-cli mcp serve --transport http --addr 127.0.0.1:3000 \
    --audit-log ~/.lark-mcp-audit.ndjson
  cloudflared tunnel --url http://127.0.0.1:3000     # -> URL https công khai
  export LARK_MCP_URL="https://<your>.trycloudflare.com/mcp"
  export LARK_MCP_BEARER_TOKEN="<một-bí-mật-mạnh>"   # truyền luôn vào lệnh serve
  ./set-transport.sh http
  ```

---

## Cài vào Claude

**Claude Code**
```bash
claude plugin marketplace add .          # chạy từ thư mục gốc của kho này
claude plugin install productivity@lark-cowork
```

**Cowork (Claude Desktop)** — thêm thư mục này như một marketplace cục bộ, sau đó cài plugin từ giao
diện plugin của Cowork. Dùng profile transport `http`.

---

## Chạy lại quá trình chuyển đổi

Bộ chuyển đổi là một chương trình Go nhỏ (`tools/larkify.go`) — cùng ngôn ngữ với phần còn lại của
hệ thống, có kiểm thử `go test`. Tái lập được và idempotent:

```bash
cd tools
go test ./...                                   # unit test
go run . convert <plugin> [<plugin> ...]        # viết lại .mcp.json + CONNECTORS.md + thuật ngữ
go run . transport stdio|http                   # đổi transport lark cho toàn bộ plugin
# hoặc build một lần:  go build -o larkify .  &&  ./larkify convert <plugin>
```

Xem [SETUP.md](./SETUP.md) để có checklist cài đặt + kiểm chứng đầy đủ.

---

## Cấu trúc kho mã

```
lark-cowork-plugins/
├── README.md                  # File này
├── SETUP.md                   # Checklist cài đặt & kiểm chứng
├── QC-AUDIT.md                # Báo cáo QC & audit (24/24 kiểm tra xanh)
├── docs/                      # 📘 Tài liệu tiếng Việt cho người dùng nghiệp vụ
├── connectors/                # Ánh xạ connector + tài liệu lõi (PATTERNS/RECIPES/FUSION)
├── tools/                     # Bộ chuyển đổi larkify (Go) + audit.sh
├── set-transport.sh           # Đổi transport stdio/http
│
├── <15 plugin nghiệp vụ>/     # productivity, sales, finance, legal, marketing…
├── <7 plugin thuần Lark>/     # lark-base-deploy, daily-assistant, crm-sales…
├── pdf-viewer/                # Giữ nguyên từ bản gốc
├── cowork-plugin-management/  # Giữ nguyên từ bản gốc
└── partner-built/             # Plugin đối tác (apollo, brand-voice, zoom…)
```

Mỗi thư mục plugin gồm: `plugin.json`, `.mcp.json`, `CONNECTORS.md`, `README.md`, và thư mục
`skills/` chứa các file `SKILL.md`.

---

## Chất lượng & kiểm thử

Toàn bộ marketplace đã qua kiểm chứng 4 cấp (cấu trúc → đúng đắn → ngữ nghĩa → runtime), **24/24
kiểm tra tự động đạt**. Chạy lại bất cứ lúc nào:

```bash
cd lark-cowork-plugins
./tools/audit.sh            # L1–L4 offline (tĩnh + build + handshake + thẻ)
./tools/audit.sh --live     # + đọc thật qua xác thực & vòng tạo-xóa task tự dọn
```

Chi tiết ma trận test xem [QC-AUDIT.md](./QC-AUDIT.md).

---

## Ghi nhận & bản quyền

Xây dựng trên bộ knowledge-work-plugins mã nguồn mở của Anthropic (xem [LICENSE](./LICENSE)). Lark,
Feishu và các dịch vụ bên thứ ba được kết nối là thương hiệu của chủ sở hữu tương ứng.

---

## Tác giả

**Nguyễn Ngọc Tuấn**
Founder Transform Group — **Lark Platinum Partner**
🌐 Dự án: [larkcowork.com](https://larkcowork.com)
