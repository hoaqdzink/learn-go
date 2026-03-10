## Networking — HTTP nâng cao (HTTP/1.1, HTTP/2, HTTP/3, WebSocket)

### 1. Mục tiêu

- Vượt qua mức “biết HTTP là request/response” để hiểu:
  - Khác nhau giữa HTTP/1.1, HTTP/2, HTTP/3.
  - Connection reuse, multiplexing.
  - WebSocket và cơ chế upgrade từ HTTP.
- Liên quan trực tiếp đến:
  - Thiết kế API hiệu quả.
  - Hiểu behavior của load balancer, reverse proxy.

### 2. HTTP/1.1

- Đặc điểm:
  - Mỗi connection TCP có thể gửi **nhiều request** (keep-alive).
  - Nhưng thường **tuần tự**: response xong mới gửi tiếp (head-of-line blocking ở HTTP layer).
- Vấn đề:
  - Nhiều request nhỏ → phải mở nhiều connection song song hoặc chờ.

### 3. HTTP/2

- Cải tiến chính:
  - **Multiplexing**:
    - Nhiều stream HTTP chạy trên **1 connection TCP**.
    - Giảm overhead mở/đóng connection.
  - **Header compression**:
    - HPACK, giảm kích thước header lặp lại.
  - **Server push** (hiện thực tế ít dùng dần).
- Ứng dụng:
  - Giảm latency, tối ưu cho web/app nhiều request nhỏ.

### 4. HTTP/3 (trên QUIC)

- Chạy trên **QUIC** (UDP-based), không dùng TCP.
- Giải quyết:
  - Head-of-line blocking ở TCP layer.
  - Kết nối nhanh hơn (0-RTT/1-RTT handshake).
- Dùng nhiều cho:
  - Dịch vụ lớn (Google, YouTube, Cloudflare,…).

### 5. WebSocket

- Là protocol cho **kết nối 2 chiều, lâu dài** giữa client–server.
- Bắt đầu bằng HTTP handshake:
  - Client gửi request với header: `Upgrade: websocket`.
  - Server đồng ý → **upgrade** từ HTTP sang WebSocket.
- Ứng dụng:
  - Chat realtime, notification, game, dashboard realtime.

### 6. Liên hệ với Go

- `net/http` support HTTP/1.1 by default.
- HTTP/2:
  - Go hỗ trợ HTTP/2 cho server/client (khi dùng TLS chuẩn).
- WebSocket:
  - Thường dùng thư viện như `gorilla/websocket`.

### 7. Bài tập gợi ý

- Dùng browser DevTools:
  - Tab Network → xem:
    - Protocol (h2, http/1.1, h3).
    - Connection reuse.
- Đọc tài liệu ngắn:
  - HTTP/2 high-level overview.
  - WebSocket protocol overview.
- Suy nghĩ:
  - Với API backend, khi nào HTTP/1.1 là đủ?
  - Khi nào cần tận dụng HTTP/2/3?

