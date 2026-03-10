## Networking — Routing, NAT & Firewall

### 1. Mục tiêu

- Hiểu:
  - Gói tin **đi qua các router** như thế nào (routing).
  - Tại sao IP private vẫn ra Internet được (NAT).
  - Firewall chặn/mở port ra sao.
- Đây là nền để:
  - Debug service “chạy local được nhưng server không gọi được”.
  - Thiết kế hệ thống trong VPC / private network.

### 2. Routing cơ bản

- **Router**:
  - Thiết bị trung gian, dùng **routing table** để quyết định “gửi gói đi đâu tiếp”.
- **Default gateway**:
  - “Cửa” mặc định để gửi gói ra khỏi mạng local.
- Mỗi gói tin đi từ A → B có thể:
  - Đi qua nhiều router (nhiều “hop”).
  - Dừng lại ở router nếu không có route phù hợp.

### 3. NAT (Network Address Translation)

- Vấn đề:
  - IPv4 không đủ địa chỉ cho mọi thiết bị trên thế giới.
- Giải pháp:
  - Dùng **IP private** trong mạng nội bộ (VD: `192.168.x.x`, `10.x.x.x`).
  - Router dùng **NAT** để:
    - Map nhiều IP private → 1 (hoặc vài) IP public khi ra Internet.
- Liên hệ backend:
  - Service trong VPC thường dùng IP private.
  - Khi gọi ra ngoài/Internet, traffic đi qua NAT gateway.

### 4. Firewall (lọc traffic)

- Firewall quyết định:
  - Cho phép / chặn traffic theo:
    - IP nguồn/đích.
    - Port nguồn/đích.
    - Protocol (TCP/UDP).
- Trong cloud (AWS/GCP/Azure):
  - Security Group, Network ACL, v.v.
- Tình huống hay gặp:
  - App chạy local OK nhưng deploy lên không gọi được DB:
    - Port DB không mở.
    - Security Group chưa allow.

### 5. Bài tập gợi ý

- Dùng `traceroute` (hoặc `tracert` trên Windows) đến 1 domain:
  - Quan sát số lượng hop.
- Vào router nhà riêng:
  - Xem IP WAN (public) vs LAN (private).
- Suy nghĩ:
  - Khi bạn deploy 1 service trong cloud, làm sao để:
    - Chỉ cho phép máy trong cùng VPC truy cập DB.
    - Chặn truy cập DB từ Internet.

