## Networking — DNS & Socket

### 1. Mục tiêu

- Hiểu:
  - DNS là gì.
  - Tại sao cần DNS thay vì nhớ IP.
  - Socket là gì, liên hệ với TCP/UDP.
- Đây là nền để:
  - Debug vấn đề domain, DNS, kết nối network.
  - Hiểu cách service trong hệ thống gọi nhau.

### 2. DNS (Domain Name System)

- DNS là “**sổ danh bạ của Internet**”.
- Chức năng:
  - Ánh xạ **tên miền** → **địa chỉ IP**.
  - Ví dụ:
    - `google.com` → `142.250.xxx.xxx`.
- Quá trình cơ bản (giản lược):
  1. Ứng dụng (browser) hỏi resolver (thường là DNS server của ISP hoặc 8.8.8.8).
  2. Resolver truy vấn lên các DNS server cấp cao hơn.
  3. Nhận về IP và cache lại.

### 3. Loại bản ghi DNS thường gặp

- `A`: map domain → IPv4.
- `AAAA`: map domain → IPv6.
- `CNAME`: alias từ 1 domain đến domain khác.
- `MX`: bản ghi mail.

### 4. Socket (mở rộng từ phần TCP/IP)

- Socket là **API**/khái niệm cho phép chương trình:
  - Gửi/nhận dữ liệu qua network.
- Với TCP:
  - Server:
    - `listen` trên `IP:Port`.
    - `accept` connection → socket mới.
  - Client:
    - `connect` tới `IP:Port` → tạo socket phía client.
- Với UDP:
  - Không có khái niệm connection, chỉ gửi/nhận packet.

### 5. Bài tập gợi ý

- Dùng command-line:
  - `nslookup google.com`.
  - `dig google.com` (nếu có).
  - Quan sát bản ghi `A`, `CNAME`.
- Suy nghĩ các câu hỏi:
  - Khi bạn gõ `https://example.com` trong trình duyệt, **chuỗi bước** từ DNS → TCP handshake → HTTP request diễn ra như thế nào?
  - Nếu DNS bị cấu hình sai, hệ thống backend của bạn sẽ bị ảnh hưởng ra sao?

