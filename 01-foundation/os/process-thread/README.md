## OS — Process & Thread

### 1. Mục tiêu

- Hiểu sự khác nhau giữa:
  - **Process** (tiến trình).
  - **Thread** (luồng).
- Nắm được:
  - Cách OS quản lý process/thread ở mức khái niệm.
  - Liên hệ với **goroutine** trong Go (ở phase sau).

### 2. Khái niệm cốt lõi

#### 2.1. Process

- Là **chương trình đang chạy** có:
  - Không gian địa chỉ riêng (virtual address space).
  - Các tài nguyên riêng: mở file, socket, handle, v.v.
- Mỗi process thường có:
  - PID (process ID).
  - Một hoặc nhiều thread.

#### 2.2. Thread

- Thread là **đơn vị thực thi** bên trong process.
- Các thread trong cùng 1 process:
  - **Chung memory** (code, heap, segment dữ liệu).
  - **Riêng** stack, register context.
- Lợi ích:
  - Tận dụng đa lõi CPU.
  - Chia task lớn thành nhiều luồng nhỏ.

#### 2.3. Multi-process vs Multi-thread

- Multi-process:
  - Isolated tốt, crash 1 process không kéo theo process khác.
  - Giao tiếp khó hơn (IPC: pipe, socket, shared memory).
- Multi-thread:
  - Giao tiếp dễ vì chung memory.
  - Nhưng dễ sinh **race condition** nếu không cẩn thận.

### 3. Liên hệ với Go

- Go dùng **goroutine**:
  - Nhiều goroutine chạy trên ít OS thread (mô hình M:N).
  - Scheduler của Go quyết định goroutine nào chạy trên OS thread nào.
- Ở phase này **chỉ cần hiểu process/thread**, còn goroutine sẽ học sâu ở phase Concurrency.

### 4. Bài tập gợi ý

- Dùng Linux/Mac:
  - Chạy `ps aux` hoặc `top`, quan sát:
    - Nhiều process khác nhau (Chrome, VSCode, Go binary...).
  - Dùng `htop` (nếu có) để xem thread.
- Viết một chương trình Go đơn giản:
  - In ra PID của process (`os.Getpid()`).
  - Tạo nhiều goroutine và quan sát CPU usage.

