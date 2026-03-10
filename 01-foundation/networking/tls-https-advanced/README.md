## Networking — TLS & HTTPS nâng cao

### 1. Mục tiêu

- Hiểu sâu hơn về:
  - TLS handshake (ở mức ý tưởng).
  - Certificate, CA, chain.
  - Lỗi HTTPS phổ biến (self-signed, expired, hostname mismatch).
- Dùng để:
  - Debug lỗi HTTPS khi gọi API.
  - Thiết lập HTTPS cho service backend (qua reverse proxy / ingress).

### 2. TLS handshake (giản lược)

- Ý tưởng:
  1. Client chào (ClientHello): gửi list cipher suites, version.
  2. Server trả lời (ServerHello): chọn cipher, gửi certificate.
  3. Hai bên:
     - Dùng **public/private key** để trao đổi secret.
     - Sinh ra **symmetric key** để mã hóa data sau đó.
- Sau handshake:
  - Data được **mã hóa** bằng symmetric key (nhanh hơn).

### 3. Certificate & CA

- Certificate chứa:
  - Public key của server.
  - Thông tin domain, tổ chức.
  - Thời hạn (valid from/to).
- CA (Certificate Authority):
  - Tổ chức tin cậy ký vào certificate.
  - Browser/OS có **trust store** các CA gốc.
- Chain:
  - Certificate của site → intermediate CA → root CA.

### 4. Lỗi HTTPS phổ biến

- **Self-signed certificate**:
  - Cert không được CA tin cậy ký → browser cảnh báo.
- **Expired certificate**:
  - Hết hạn → client từ chối kết nối.
- **Hostname mismatch**:
  - Domain truy cập không khớp với `CN`/`SAN` trong cert.
- **Old TLS version / weak cipher**:
  - Một số client mới từ chối kết nối với TLS quá cũ.

### 5. Liên hệ với backend / Go

- Khi viết HTTP client trong Go:
  - Có thể cấu hình:
    - TLS version.
    - Skip verify (không nên dùng trong production).
- Khi deploy service:
  - Thường offload TLS ở:
    - Reverse proxy (Nginx, Envoy).
    - Ingress controller (Kubernetes).

### 6. Bài tập gợi ý

- Dùng browser:
  - Click vào ổ khóa trên thanh địa chỉ.
  - Xem thông tin certificate: issuer, valid date, SAN.
- Dùng `openssl s_client -connect host:443`:
  - Xem chain certificate, TLS version.
- Suy nghĩ:
  - Nếu cert hết hạn đúng ngày hôm nay, hệ thống của bạn sẽ có triệu chứng gì?

