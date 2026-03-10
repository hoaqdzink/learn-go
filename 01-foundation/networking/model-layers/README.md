## Networking — Mô hình tầng (OSI & TCP/IP)

### 1. Mục tiêu

- Nắm được **ý tưởng chia tầng** trong mạng:
  - Ứng dụng, transport, network, link, physical.
- Biết HTTP, TCP, IP, Ethernet đang nằm ở đâu.
- Dùng để:
  - Hiểu rõ “HTTP chạy trên TCP, TCP chạy trên IP…”.
  - Debug: lỗi đang nằm ở layer nào.

### 2. Mô hình OSI (7 tầng) — để lấy concept

1. **Physical**: dây, sóng, bit 0/1, điện áp, tín hiệu.
2. **Data Link**: frame, MAC address, switch.
3. **Network**: IP address, router, routing.
4. **Transport**: TCP, UDP (port, connection, reliability).
5. **Session**: quản lý phiên làm việc (ít đụng trực tiếp).
6. **Presentation**: encode/decode, encryption (thường được gộp vào app).
7. **Application**: HTTP, gRPC, DNS, SMTP, v.v.

Dùng OSI chủ yếu để **nói chuyện conceptual**, không cần nhớ chi tiết hết.

### 3. Mô hình TCP/IP (thực tế hơn)

Thường rút gọn thành 4 tầng:

1. **Link layer**:
   - Ethernet, WiFi.
   - MAC address, frame.
2. **Internet layer**:
   - IP, ICMP.
   - Định tuyến, địa chỉ IP nguồn/đích.
3. **Transport layer**:
   - TCP, UDP.
   - Port, connection, reliability, segmentation.
4. **Application layer**:
   - HTTP, HTTPS, gRPC, DNS, SMTP, MQTT, v.v.

### 4. Liên hệ với backend / Go

- Khi bạn viết 1 HTTP server trong Go:
  - App layer: **HTTP** (request/response, header, body).
  - Transport: **TCP** (connection, handshake, retry).
  - Internet: **IP** (địa chỉ server/client).
  - Link: Ethernet/WiFi (frame, MAC).
- Khi có lỗi:
  - DNS sai → vấn đề ở **application (DNS)** / internet (IP resolution).
  - Port bị chặn → vấn đề ở **transport/link** (firewall, security group).
  - HTTP 500 → vấn đề ở chính **application** (code của bạn).

### 5. Bài tập gợi ý

- Vẽ lại sơ đồ 4 tầng TCP/IP, gắn:
  - HTTP, gRPC, DNS ở application layer.
  - TCP/UDP ở transport.
  - IP ở internet.
  - Ethernet/WiFi ở link.
- Lấy 1 lỗi thực tế (VD: không gọi được API từ client):
  - Liệt kê các nguyên nhân có thể ở từng tầng.

