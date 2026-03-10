## CS — Complexity (Độ phức tạp)

### 1. Mục tiêu

- Hiểu rõ **Big-O notation** và ý nghĩa thực tế.
- Biết phân tích sơ bộ độ phức tạp thời gian và bộ nhớ của:
  - Thuật toán.
  - Cấu trúc dữ liệu.
- Hình thành **cảm giác về “chạy được / không chạy được”** khi nhìn input size.

### 2. Big-O, Big-Theta, Big-Omega (chỉ cần mức cơ bản)

- Big-O: **upper bound** — “thuật toán chạy không chậm hơn cỡ này”.
- Big-Theta: **tight bound** — “xấp xỉ chính xác”.
- Big-Omega: **lower bound** — “ít nhất nhanh thế này”.
- Thực tế làm backend + system, chủ yếu dùng **Big-O** để:
  - So sánh 2 giải pháp.
  - Ước lượng độ lớn input tối đa chịu được.

### 3. Bảng so sánh độ phức tạp thường gặp

- \(O(1)\) — constant.
- \(O(\log n)\) — logarithmic (binary search).
- \(O(n)\) — linear.
- \(O(n \log n)\) — “gần tuyến tính”, thường gặp ở sort tối ưu.
- \(O(n^2)\) — quadratic.
- \(O(2^n)\), \(O(n!)\) — thường chỉ dùng trong bài toán rất nhỏ, không phù hợp production.

### 4. Cách phân tích độ phức tạp

1. Xác định **vòng lặp chính** (loop).
2. Xem có **vòng lặp lồng nhau** hay không.
3. Kiểm tra các **gọi hàm đệ quy**:
   - Đệ quy chia đôi: thường là \(O(\log n)\) hoặc \(O(n \log n)\).
   - Đệ quy nhánh đôi/triple: dễ dẫn tới \(O(2^n)\), \(O(3^n)\).
4. Bỏ qua **constant nhỏ** khi biểu diễn Big-O:
   - \(3n + 10 \to O(n)\).
   - \(n^2 + n \log n \to O(n^2)\).

### 5. Độ phức tạp không gian (Space Complexity)

- Đếm:
  - Biến local.
  - Cấu trúc dữ liệu tạm (array, map, v.v.).
  - Độ sâu stack (đặc biệt trong recursion).
- Ví dụ:
  - DFS recursion trên tree có thể dùng \(O(h)\) stack (với \(h\) là chiều cao tree).
  - Lưu mảng độ dài `n` → \(O(n)\) memory.

### 6. Gợi ý bài tập

**Bài 1: Phân tích vòng lặp**
- Cho các đoạn code đơn giản (tự viết hoặc copy từ phần Algorithms), hãy tự:
  - Viết ra số lần lặp **ước lượng**.
  - Chuyển về Big-O.

**Bài 2: So sánh thuật toán**
- So sánh:
  - Bubble Sort vs Merge Sort vs Quick Sort (average case).
  - Dijkstra dùng array thường vs dùng priority queue (heap).

**Bài 3: Trực giác input size**
- Luyện “cảm giác”:
  - Với \(n = 10^3, 10^5, 10^7\), chi phí:
    - \(O(n)\)
    - \(O(n \log n)\)
    - \(O(n^2)\)
- Nhẩm xem thuật toán \(O(n^2)\) với \(n = 10^5\) có thực tế không (gợi ý: **không**).

**Bài 4: Recursion**
- Với mỗi hàm recursion trong phần Algorithms:
  - Vẽ cây gọi hàm (call tree) nhỏ.
  - Đếm “số node” trong call tree → suy ra độ phức tạp.

### 7. Liên hệ với backend thực tế

- Query DB:
  - Scan full table ~ \(O(n)\).
  - Dùng index thường ~ \(O(\log n)\).
- Tính toán trong service:
  - Sắp xếp list user, log, event → ưu tiên thuật toán \(O(n \log n)\).
  - Sử dụng map (hash table) cho lookup nhanh \(O(1)\) trung bình.

