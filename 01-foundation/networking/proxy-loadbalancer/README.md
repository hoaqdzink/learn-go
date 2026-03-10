## Networking — Proxy & Load Balancer

### 1. Mục tiêu

- Phân biệt được:
  - **Forward proxy** vs **reverse proxy**.
  - **Load balancer L4 vs L7**.
- Hiểu cách:
  - Request đi qua proxy/load balancer đến service.
  - Một số khái niệm: sticky session, health check, circuit breaker (liên quan System Design).

### 2. Forward proxy

- Đứng **trước client**.
- Client gửi request đến proxy, proxy gửi ra Internet.
- Ứng dụng:
  - Ẩn IP client.
  - Cache nội dung.
  - Lọc/tracking request từ client nội bộ.

### 3. Reverse proxy

- Đứng **trước server/service**.
- Client tưởng mình gọi trực tiếp server, thực ra là gọi reverse proxy.
- Ứng dụng:
  - Terminate TLS (chỉ proxy cần giữ cert).
  - Routing đến nhiều backend khác nhau (path-based, host-based).
  - Load balancing đơn giản.
- Ví dụ:
  - Nginx, HAProxy, Envoy, Traefik.

### 4. Load balancer (L4 vs L7)

- **L4 load balancer**:
  - Làm việc trên TCP/UDP (IP + Port).
  - Không nhìn vào HTTP header/body.
  - Phù hợp cho bất kỳ protocol nào trên TCP/UDP.
- **L7 load balancer**:
  - Làm việc ở application layer (HTTP/gRPC).
  - Có thể:
    - Route theo path (`/api/v1` vs `/admin`).
    - Route theo header, host.

### 5. Sticky session & Health check

- **Sticky session**:
  - Cố định 1 user → 1 backend trong 1 khoảng thời gian.
  - Dùng khi session lưu ở memory của 1 server (không dùng session store chia sẻ).
- **Health check**:
  - LB định kỳ gọi đến endpoint (VD: `/healthz`).
  - Nếu fail → đưa instance đó ra khỏi pool.

### 6. Liên hệ với Go

- Khi bạn viết service Go:
  - Thường chạy nhiều instance phía sau reverse proxy/load balancer.
  - Cần:
    - Cung cấp health endpoint.
    - Thiết kế để **stateless** (dễ scale, dễ LB).

### 7. Bài tập gợi ý

- Vẽ sơ đồ:
  - Client → Internet → Reverse Proxy/Load Balancer → N backend service Go.
- Đọc config mẫu:
  - Nginx reverse proxy cho 2 backend (upstream).
- Suy nghĩ:
  - Nếu 1 instance backend bị down, load balancer nên xử lý thế nào?

