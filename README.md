## Giới thiệu về Go (Golang)

Go (hay Golang) là ngôn ngữ lập trình được Google phát triển, tập trung vào **đơn giản, hiệu năng cao và hỗ trợ concurrency ở cấp ngôn ngữ**. Go đặc biệt mạnh trong các lĩnh vực:

- **Backend service hiệu năng cao** (API server, microservices)
- **Hệ thống phân tán / cloud-native** (Kubernetes, Docker đều viết bằng Go)
- **Tooling / CLI / DevOps** (các công cụ dòng lệnh, agent, tooling nội bộ)

### Tại sao Go phù hợp cho backend & hệ thống lớn?

- **Ngắn gọn, dễ đọc**: cú pháp tối giản, code thường ít “magic”, dễ review, dễ bảo trì.
- **Tốc độ gần với C**: Go được biên dịch xuống binary, không cần JVM/VM, rất hợp viết service hiệu năng cao.
- **Goroutine & Channel**: hỗ trợ concurrency “nhẹ” ở cấp ngôn ngữ, tạo hàng chục nghìn goroutine mà vẫn nhẹ.
- **Garbage Collector hiện đại**: giảm áp lực quản lý bộ nhớ thủ công, nhưng vẫn cho phép tối ưu khi cần.
- **Ecosystem mạnh cho backend**: net/http, gRPC, protobuf, Kafka client, Redis client, Postgres driver, v.v.

### Triết lý thiết kế của Go

- **Đơn giản hơn là phức tạp nhưng “ngầu”**: không có generics trong thời gian dài (và khi có thì cũng giữ đơn giản), không overload operator, không inheritance sâu.
- **“Composition over inheritance”**: ưu tiên struct + interface + composition thay vì class hierarchy.
- **Tooling đi kèm mạnh mẽ**:
  - `go fmt`: format code thống nhất
  - `go test`: testing tích hợp
  - `go vet`: bắt bug phổ biến
  - `go mod`: quản lý dependency

### Go phù hợp với ai?

- Người muốn đi **backend / distributed systems / cloud-native**.
- Dev muốn từ sớm có tư duy **system-level** (memory, OS, network) nhưng không muốn cực khổ như C/C++.
- Dev đã biết một ngôn ngữ khác (Java, Node.js, Python, PHP, C#) và muốn làm **service hiệu quả, dễ scale**.

### Cách sử dụng lộ trình trong repo này

Bắt đầu từ:

1. `00-go-setup/` — cài Go, làm quen `go run`, `go build`, `go mod`.
2. Sau đó tới `01-foundation/` — CS, OS, Networking.
3. Rồi lần lượt `02-core-go/`, `03-go-internals/`, ...

Lộ trình được chia thành nhiều **phase**, từ nền tảng CS → Go core → backend → distributed systems → system design → senior mindset.  
Bạn nên học theo thứ tự thư mục, mỗi thư mục có:

- **Tài liệu** (file `.md`) giải thích khái niệm
- **Ví dụ code nhỏ** (file `.go`)
- **Bài tập / mini-project** để luyện

Nguyên tắc học:

```text
Learn
→ Build
→ Debug
→ Refactor
```

Đừng chỉ đọc lý thuyết – luôn **viết code thật, debug thật, refactor lại**.

