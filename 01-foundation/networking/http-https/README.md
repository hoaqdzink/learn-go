## Networking — HTTP & HTTPS

### 1. Mục tiêu

- Hiểu **HTTP hoạt động như thế nào**:
  - Request, response, method, status code, header, body.
- Nắm được **HTTPS**:
  - Khác gì HTTP.
  - Ý tưởng về mã hóa và TLS.
- Đây là nền cho:
  - REST API, gRPC, reverse proxy, load balancer.

### 2. HTTP cơ bản

- HTTP là **protocol tầng ứng dụng**, chạy **trên TCP**.
- Một HTTP request gồm:
  - Method: `GET`, `POST`, `PUT`, `DELETE`, ...
  - Path/URL.
  - Header (metadata).
  - Body (data, thường là JSON).
- Response gồm:
  - Status code (200, 400, 401, 500, ...).
  - Header.
  - Body (thường là JSON/HTML).

### 3. Method & Status code

- Method:
  - `GET`: lấy dữ liệu.
  - `POST`: tạo mới / thao tác có side-effect.
  - `PUT/PATCH`: cập nhật.
  - `DELETE`: xóa.
- Status code:
  - 2xx: thành công (200, 201, 204,...).
  - 4xx: lỗi phía client (400, 401, 403, 404,...).
  - 5xx: lỗi phía server (500, 502, 503,...).

### 4. HTTPS & TLS (ý tưởng)

- HTTPS = HTTP + TLS (mã hóa).
- Mục đích:
  - Bảo mật dữ liệu giữa client và server.
  - Chống nghe lén, giả mạo.
- TLS dùng:
  - Certificate.
  - Public key / private key.
- Chi tiết cryptography có thể học sâu thêm ở track Security, ở đây chỉ cần hiểu:
  - **HTTP gửi plain text**, HTTPS gửi dữ liệu đã mã hóa.

### 5. Bài tập gợi ý

- Dùng:
  - `curl`.
  - Postman/Insomnia.
- Thực hành:
  - Gửi `GET`, `POST` đến 1 public API (VD: JSONPlaceholder).
  - Quan sát:
    - Header request/response.
    - Status code.
    - Body.
- Đọc spec ngắn:
  - RFC 7231 summary (HTTP/1.1).

