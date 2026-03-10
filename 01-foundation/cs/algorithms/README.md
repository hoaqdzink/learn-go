## CS — Algorithms

### 1. Mục tiêu

- Hiểu được **ý tưởng** của các thuật toán cơ bản, không chỉ nhớ code.
- Biết phân tích **độ phức tạp thời gian/không gian** ở mức Big-O.
- Áp dụng được vào **bài toán backend thực tế** (search, sort, traversal).

### 2. Các nhóm thuật toán chính

#### 2.1. Sorting (Sắp xếp)

- Thuật toán nên biết:
  - Bubble Sort (để hiểu “naive sort”, không dùng thực tế).
  - Insertion Sort.
  - Merge Sort.
  - Quick Sort.
- Với mỗi thuật toán:
  - Ý tưởng.
  - Big-O worst/average/best.
  - Trường hợp phù hợp / không phù hợp.

#### 2.2. Binary Search

- Điều kiện áp dụng: **dữ liệu có thứ tự** (sorted).
- Ý tưởng chia đôi dần khoảng tìm kiếm.
- Cài đặt dạng:
  - Trả về index nếu tìm thấy, `-1` nếu không.
  - Tìm **lower bound / upper bound** (nâng cao).

#### 2.3. DFS / BFS

- Áp dụng cho:
  - Tree.
  - Graph.
- DFS:
  - Dùng stack (hoặc recursion).
  - Ứng dụng: tìm đường, topological sort (nâng cao).
- BFS:
  - Dùng queue.
  - Ứng dụng: tìm đường đi ngắn nhất với unweighted graph.

#### 2.4. Recursion

- Khái niệm:
  - Base case.
  - Recursive case.
- Ví dụ kinh điển:
  - Tính `n!`.
  - Fibonacci (và vì sao phiên bản naive rất chậm).
  - Duyệt tree.

#### 2.5. Greedy

- Ý tưởng: luôn chọn “tốt nhất tại thời điểm hiện tại”.
- Ví dụ:
  - Đổi tiền với số lượng tờ ít nhất (với hệ mệnh giá phù hợp).
  - Chọn interval tối đa không giao nhau (activity selection).

#### 2.6. Dynamic Programming cơ bản (DP)

- Ý tưởng:
  - Overlapping subproblems.
  - Optimal substructure.
  - Memoization vs Tabulation.
- Ví dụ entry-level:
  - Fibonacci tối ưu bằng DP.
  - Bài toán climbing stairs (`f(n) = f(n-1) + f(n-2)`).

### 3. Gợi ý lộ trình học

1. Bắt đầu với **Sorting cơ bản** → cài đặt bằng tay.
2. Học **Binary Search** → luyện thật nhiều bài dạng “tìm trong sorted array”.
3. Học **DFS/BFS** trên tree trước, sau đó mở rộng sang graph.
4. Hiểu rõ **Recursion** (base case, stack call).
5. Làm vài bài **Greedy** đơn giản.
6. Kết thúc bằng **DP cơ bản** (Fibonacci, climbing stairs, coin change đơn giản).

### 4. Bài tập gợi ý

**Sorting**
- Cài đặt:
  - `BubbleSort`, `InsertionSort`, `MergeSort`, `QuickSort`.
- Đo thời gian chạy (bằng Go `time.Now()`) khi:
  - Input đã sort sẵn.
  - Input đảo ngược.
  - Input random.

**Binary Search**
- Viết `BinarySearch(arr []int, target int) int`.
- Viết hàm:

  - `LowerBound(arr []int, target int) int`
  - `UpperBound(arr []int, target int) int`

**DFS/BFS**
- Dùng adjacency list, viết:
  - `DFS(start)` in ra thứ tự duyệt.
  - `BFS(start)` in ra thứ tự duyệt.
- Áp dụng BFS để:
  - Tìm **độ dài đường đi ngắn nhất** từ `start` đến `end` trong unweighted graph.

**Recursion**
- Viết:
  - `Factorial(n int) int`
  - `Fib(n int) int` (version naive) → sau đó tối ưu.
- Viết hàm đếm:
  - Số node trong 1 tree bằng recursion.

**Greedy**
- Bài toán đổi tiền (VD: 1k, 2k, 5k, 10k,...):
  - Cho số tiền `N`, in ra số lượng tờ tiền ít nhất (nếu hệ mệnh giá phù hợp).
- Bài toán chọn interval không giao nhau:
  - Cho danh sách khoảng `[start, end]`, chọn tối đa số interval không overlap.

**Dynamic Programming**
- Viết lại Fibonacci theo 3 cách:
  - Recursion naive.
  - Recursion + memoization.
  - Bottom-up (tabulation).
- Bài climbing stairs:
  - Mỗi lần bước được 1 hoặc 2 bậc, có bao nhiêu cách leo lên đỉnh bậc `n`?

