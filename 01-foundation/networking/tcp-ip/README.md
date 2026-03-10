## Networking — TCP/IP

### 1. Mục tiêu

- Hiểu ở mức **engineer backend cần**:
  - TCP là gì, khác gì UDP.
  - IP, port, socket.
  - Connection, handshake, timeout, retransmission (ý tưởng).
- Đủ để:
  - Debug các vấn đề network cơ bản.
  - Hiểu vì sao Go net/http, gRPC hoạt động như vậy.

### 2. IP & Port

- **IP**:
  - Địa chỉ của máy trong mạng (IPv4/IPv6).
  - Ví dụ: `192.168.1.10`, `10.0.0.5`.
- **Port**:
  - “Cổng” trên máy, cho phép nhiều service cùng chạy.
  - Ví dụ:
    - 80 (HTTP).
    - 443 (HTTPS).
    - 5432 (PostgreSQL).
- Kết hợp lại:
  - `IP:Port` định danh 1 endpoint (VD: `10.0.0.5:8080`).

### 3. TCP vs UDP

- **TCP (Transmission Control Protocol)**:
  - Hướng connection (connection-oriented).
  - Đảm bảo:
    - Reliable (không mất dữ liệu, hoặc sẽ retry).
    - Ordered (dữ liệu đến đúng thứ tự).
    - Stream-based.
  - Chi phí cao hơn, nhưng rất phù hợp cho HTTP, gRPC, DB connection.
- **UDP**:
  - Không connection (connectionless).
  - Không đảm bảo delivery/order.
  - Phù hợp cho:
    - Streaming, game, voice, video call.

### 4. TCP Handshake (3-way handshake)

- Ý tưởng (giản lược):
  1. Client gửi `SYN` đến server.
  2. Server trả `SYN-ACK`.
  3. Client gửi `ACK`.
- Sau đó connection được thiết lập, 2 bên có thể gửi dữ liệu.

### 5. Socket

- Socket là **endpoint** của connection TCP.
- Trong code:
  - Server:
    - Lắng nghe (`listen`) trên `IP:Port`.
    - Accept connection để có socket mới.
  - Client:
    - Kết nối đến `IP:Port` → tạo socket phía client.

### 6. Bài tập gợi ý

- Về mặt thực hành:
  - Dùng `netstat`, `ss`, hoặc `lsof -i` để:
    - Xem các port đang lắng nghe trên máy.
    - Thấy các connection đang mở.
- Viết một chương trình Go nhỏ (ở phase sau):
  - Server TCP:
    - Lắng nghe trên port 9000.
    - Nhận dữ liệu và echo lại.
  - Client TCP:
    - Kết nối tới server, gửi 1 chuỗi text, nhận lại và in ra.

