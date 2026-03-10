## OS — Memory & Stack/Heap

### 1. Mục tiêu

- Hiểu **bộ nhớ process** được chia thành những vùng gì (khái niệm).
- Phân biệt được **stack** và **heap**.
- Nắm ý tưởng **memory allocation** và liên hệ với Go (stack allocation, heap allocation).

### 2. Các vùng nhớ chính trong process

- **Code segment**: chứa mã máy (instruction) của chương trình.
- **Data segment**:
  - Global/static variables đã được khởi tạo.
- **Heap**:
  - Vùng nhớ cấp phát **dynamic** (malloc/new, trong Go là `new`/`make`/literal).
- **Stack**:
  - Vùng nhớ cho **local variables, call frames** của function.

### 3. Stack

- Đặc điểm:
  - LIFO (Last In First Out).
  - Mỗi lần gọi function → thêm 1 **stack frame**.
  - Khi return → pop stack frame.
- Ưu điểm:
  - Cấp phát/giải phóng rất nhanh (chỉ thay đổi stack pointer).
  - Dễ quản lý (OS/quản lý runtime đảm nhiệm).
- Nhược điểm:
  - Kích thước giới hạn.
  - Không phù hợp cho dữ liệu lớn sống lâu.

### 4. Heap

- Đặc điểm:
  - Dùng cho dữ liệu cần sống lâu, không biết chính xác thời điểm giải phóng.
  - Cấp phát phức tạp hơn, có thể phân mảnh.
- Trong ngôn ngữ có GC (như Go):
  - Heap được quản lý bởi **Garbage Collector**.
- Trong C/C++:
  - Lập trình viên phải tự `malloc/free` hoặc `new/delete`.

### 5. Liên hệ với Go

- Khi bạn khai báo biến trong Go:
  - Compiler sẽ quyết định **stack hay heap** thông qua **escape analysis**.
- Ý tưởng:
  - Biến chỉ dùng trong scope hẹp → thường ở stack.
  - Biến “thoát” khỏi scope (return pointer ra ngoài) → phải đưa lên heap.
- Hiểu cách này rất quan trọng khi tối ưu performance & GC.

### 6. Bài tập gợi ý

- Đọc tài liệu về **process memory layout** (Linux).
- Viết chương trình Go:
  - Tạo struct, trả về pointer, chạy với flag `-gcflags=-m` để xem compile log (escape analysis) — phần này sẽ làm sâu hơn ở phase Go Internals.
- Suy nghĩ các câu hỏi:
  - Tại sao recursion sâu có thể gây **stack overflow**?
  - Khi nào dữ liệu nên được cấp phát trên stack, khi nào trên heap (về mặt ý tưởng)?

