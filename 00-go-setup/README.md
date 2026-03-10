## Phase 0 — Cài đặt & Làm quen Go

Trước khi vào Foundation (CS/OS/Networking), bạn nên **setup môi trường Go chuẩn** và làm quen với cách chạy 1 chương trình Go đơn giản.

### 1. Cài đặt Go trên macOS / Linux / Windows

- Vào trang chính thức: `https://go.dev/dl`
- Chọn bản ổn định (stable) mới nhất, tải theo OS của bạn:
  - macOS: file `.pkg`.
  - Windows: file `.msi`.
  - Linux: có thể dùng package manager hoặc tarball.
- Sau khi cài, kiểm tra:

```bash
go version
```

Nếu hiện ra phiên bản (vd: `go version go1.xx.x darwin/amd64`) là OK.

### 2. GOROOT, GOPATH (overview)

- **GOROOT**:
  - Thư mục cài Go (nơi chứa compiler, stdlib).
  - Thường **không cần tự set**, installer lo sẵn.
- **GOPATH** (từ Go modules trở đi ít đụng hơn, nhưng nên biết):
  - Thư mục làm “workspace” chứa:
    - `bin/`: binary.
    - `pkg/`: compiled package.
    - `src/`: source cũ (trước thời modules).
  - Có thể xem bằng:

```bash
go env GOPATH
```

### 3. Go modules (cách quản lý project hiện đại)

- Từ Go 1.11+, **go modules** là chuẩn mặc định.
- Trong 1 project mới:

```bash
mkdir hello-go && cd hello-go
go mod init github.com/yourname/hello-go
```

- File `go.mod` chứa:
  - Tên module.
  - Version Go.
  - Dependency.

### 4. Chương trình Go đầu tiên

- Tạo file `main.go`:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, Go!")
}
```

- Chạy:

```bash
go run main.go
```

- Build binary:

```bash
go build -o app
./app
```

### 5. Cấu trúc project đi theo lộ trình này

Repo này được thiết kế như 1 **learning workspace**, không phải 1 app duy nhất:

- Mỗi `phase/` chứa nhiều thư mục nhỏ với:
  - `README.md`: kiến thức.
  - File `.go`: ví dụ, bài tập.
- Bạn có thể:
  - Tạo 1 module gốc ở root (vd: `go mod init learn-go`).
  - Mỗi ví dụ dùng package `main` riêng và chạy bằng `go run path/to/file.go`.

### 6. Công cụ đi kèm nên cài

- **Editor**:
  - VSCode hoặc Cursor với Go extension.
- Go tools:
  - `gofmt` (đi kèm Go): format code.
  - `goimports` (thường sẽ được cài qua extension): format + sắp xếp import.
  - `go test`: chạy test.

### 7. Bài tập gợi ý

- Cài Go và:
  - In ra `go version`.
  - In ra `go env GOPATH` và `go env GOROOT`.
- Tạo project `hello-go` riêng:
  - Dùng `go mod init`.
  - Viết `main.go` in ra:
    - Xin chào + tên bạn.
    - Phiên bản Go đang dùng (`runtime.Version()`).

