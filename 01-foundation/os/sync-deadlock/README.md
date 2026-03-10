## OS — Synchronization & Deadlock

### 1. Mục tiêu

- Hiểu vì sao **nhiều thread/process cùng truy cập dữ liệu chung** dễ gây lỗi.
- Nắm được:
  - **Mutex** và cơ chế khóa.
  - **Deadlock**: là gì, vì sao xảy ra, cách tránh ở mức cơ bản.
- Liên hệ với:
  - `sync.Mutex`, `sync.RWMutex` trong Go (phase Concurrency).

### 2. Race Condition

- Xảy ra khi:
  - Nhiều thread cùng đọc/ghi lên **shared state**.
  - Thứ tự thực thi không được kiểm soát → kết quả không đoán trước.
- Ví dụ:
  - Nhiều thread cùng tăng 1 biến `counter++` mà không khóa → giá trị cuối cùng sai.

### 3. Mutex (Mutual Exclusion)

- Là **primitive** để đảm bảo **chỉ 1 thread** được truy cập vùng critical section tại 1 thời điểm.
- Cách dùng (ý tưởng chung, ngôn ngữ nào cũng tương tự):

```text
lock(mutex)
// critical section
unlock(mutex)
```

- Trong Go:
  - `sync.Mutex` với method `Lock()` và `Unlock()`.

### 4. Deadlock

- Deadlock xảy ra khi **nhóm thread/resource “chờ nhau mãi mãi”**.
- 4 điều kiện (Coffman conditions) — chỉ cần nhớ idea:
  1. Mutual exclusion (tài nguyên không share được).
  2. Hold and wait (thread giữ 1 số tài nguyên và chờ thêm tài nguyên khác).
  3. No preemption (không tự động thu hồi tài nguyên).
  4. Circular wait (vòng tròn chờ đợi).
- Ví dụ kinh điển:
  - Thread A giữ lock1, chờ lock2.
  - Thread B giữ lock2, chờ lock1.

### 5. Cách giảm nguy cơ deadlock (ở mức cơ bản)

- Thống nhất **thứ tự lock** (global ordering).
- Tránh giữ lock quá lâu.
- Hạn chế lock lồng nhau nếu không cần thiết.

### 6. Bài tập gợi ý

- Viết pseudo (hoặc code Go đơn giản ở phase sau) cho:
  - 2 goroutine cùng tăng 1 biến `counter` có và không dùng mutex.
  - Quan sát sự khác biệt.
- Vẽ sơ đồ deadlock:
  - 2 thread, 2 lock → minh họa vòng tròn chờ đợi.
- Suy nghĩ các câu hỏi:
  - Tại sao database cũng có khái niệm deadlock?
  - Trong hệ thống phân tán, deadlock có thể xảy ra dưới dạng nào?

