## Algorithms — Binary Search

### 1. Mục tiêu

- Nắm **điều kiện áp dụng**: dữ liệu phải được sắp xếp.
- Cài đặt được binary search không bị bug mid/left/right.
- Hiểu và cài đặt thêm:
  - `LowerBound`.
  - `UpperBound`.

### 2. Nội dung chính

- Ý tưởng chia đôi khoảng tìm kiếm.
- Dừng khi:
  - Tìm thấy phần tử bằng target.
  - Hoặc left > right (không tìm thấy).
- Biến thể:
  - Tìm vị trí xuất hiện đầu tiên / cuối cùng.
  - Tìm vị trí chèn phù hợp.

### 3. Bài tập gợi ý

- Viết các hàm:
  - `BinarySearch(arr []int, target int) int`
  - `LowerBound(arr []int, target int) int`
  - `UpperBound(arr []int, target int) int`
- Test trên:
  - Mảng không có target.
  - Mảng có nhiều giá trị trùng target.

