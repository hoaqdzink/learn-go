## Networking — Công cụ troubleshooting

### 1. Mục tiêu

- Biết dùng các tool cơ bản để:
  - Kiểm tra kết nối.
  - Kiểm tra DNS.
  - Kiểm tra port/service.
- Rất quan trọng khi **debug service không gọi được nhau**.

### 2. Các công cụ chính

#### 2.1. ping

- Kiểm tra:
  - Có reach được host không (ICMP echo).
  - Round-trip time.
- Lưu ý:
  - Một số server chặn ICMP → ping fail nhưng service vẫn chạy.

#### 2.2. traceroute / tracert

- Xem đường đi của gói tin qua các router (hop).
- Dùng để:
  - Phát hiện gói bị “kẹt” ở đâu trên đường.

#### 2.3. nslookup / dig

- Kiểm tra DNS:
  - Tra IP từ tên miền.
  - Xem bản ghi A, CNAME, MX, v.v.
- Dùng khi:
  - App báo “cannot resolve host”.

#### 2.4. curl

- Gửi HTTP request:
  - `GET`, `POST`, header, body, v.v.
- Dùng để:
  - Test API nhanh.
  - Debug status code, header.

#### 2.5. telnet / nc (netcat)

- Kiểm tra:
  - Có kết nối được tới `IP:Port` không.
- Ví dụ:
  - `nc -vz host port`.
- Dùng khi:
  - Muốn biết port có mở không (DB, Redis, HTTP…).

### 3. Bài tập gợi ý

- Với 1 API đơn giản (ví dụ `https://jsonplaceholder.typicode.com/posts`):
  - Dùng `ping` domain.
  - Dùng `nslookup` để xem IP.
  - Dùng `curl` để gửi `GET` và xem response.
- Giả lập trường hợp:
  - Port bị chặn:
    - Dùng `nc` để connect tới 1 port đóng → xem lỗi.
  - DNS sai:
    - Sửa file `/etc/hosts` (cẩn thận) hoặc dùng 1 domain không tồn tại → xem lỗi.

