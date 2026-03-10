## Algorithms — Sorting

Sorting là nhóm thuật toán **cực nền tảng**, đi theo bạn suốt từ CS cơ bản tới backend (sort log, sort kết quả query, sort theo nhiều tiêu chí,…).

---

### 1. Mục tiêu

- Hiểu **ý tưởng trực quan** của từng thuật toán sắp xếp cơ bản.
- Ghi nhớ được **độ phức tạp** (best/avg/worst).
- Tự code lại được bằng Go mà **không nhìn tài liệu**.

---

### 2. Thuật toán cần nắm (lý thuyết + trực giác)

#### 2.1. Bubble Sort

- Ý tưởng:
  - Duyệt mảng nhiều lần.
  - Mỗi lần duyệt, so sánh từng cặp phần tử kề nhau:
    - Nếu `a[i] > a[i+1]` thì đổi chỗ.
  - Sau 1 lượt, phần tử lớn nhất sẽ “nổi” về cuối mảng.
- Độ phức tạp:
  - Worst / Avg: \(O(n^2)\).
  - Best: \(O(n)\) nếu mảng đã sort và bạn có tối ưu “không swap lần nào thì dừng”.
- Khi nào dùng:
  - Gần như **chỉ để học**, không dùng trong production.

#### 2.2. Insertion Sort

- Ý tưởng:
  - Giống bạn **xếp bài trên tay**:
    - Duyệt từng phần tử từ trái sang phải.
    - Mỗi phần tử mới “chèn” vào đúng vị trí trong phần mảng đã sort ở bên trái.
- Độ phức tạp:
  - Worst / Avg: \(O(n^2)\).
  - Best: \(O(n)\) khi gần như đã sort.
- Khi nào dùng:
  - Mảng nhỏ.
  - Mảng gần như đã sort (ít phần tử sai vị trí).

#### 2.3. Merge Sort (Divide & Conquer)

- Ý tưởng:
  1. Chia đôi mảng thành 2 nửa.
  2. Đệ quy sort từng nửa.
  3. **Merge** 2 mảng con đã sort lại thành 1 mảng sort.
- Đặc điểm:
  - Luôn chia làm 2, chiều sâu đệ quy ~ \(\log n\).
  - Mỗi tầng merge tốn \(O(n)\).
- Độ phức tạp:
  - Worst / Avg / Best: \(O(n \log n)\).
  - Space: \(O(n)\) (cần mảng tạm để merge).
- Khi nào dùng:
  - Cần ổn định (stable sort).
  - Dữ liệu lớn, cần đảm bảo worst-case tốt.

#### 2.4. Quick Sort

- Ý tưởng:
  1. Chọn 1 phần tử làm **pivot**.
  2. Chia mảng thành 2 phía:
     - Bên trái < pivot.
     - Bên phải ≥ pivot.
  3. Đệ quy sort 2 bên.
- Đặc điểm:
  - Rất nhanh trong thực tế nếu chọn pivot tốt.
  - Có thể implement in-place (ít tốn memory).
- Độ phức tạp:
  - Avg: \(O(n \log n)\).
  - Best: \(O(n \log n)\).
  - Worst: \(O(n^2)\) (nếu pivot luôn tệ, ví dụ mảng đã sort mà chọn pivot là phần tử đầu/cuối).
- Khi nào dùng:
  - Trường hợp chung chung, không quá xấu.
  - Nhiều ngôn ngữ dùng biến thể Quick Sort / IntroSort cho sort mặc định.

---

### 3. Gợi ý cách học / luyện

1. **Vẽ tay** 1–2 ví dụ với mảng nhỏ (~5–8 phần tử) cho từng thuật toán.
2. Viết lại code Go cho:
   - `BubbleSort`
   - `InsertionSort`
   - `MergeSort`
   - `QuickSort`
3. Với mỗi hàm:
   - Ghi chú lại Big-O.
   - Nhận xét khi input:
     - Đã sort sẵn.
     - Đảo ngược.
     - Random.
4. Đo thời gian chạy với `time.Now()` để **cảm nhận** khác biệt giữa \(O(n^2)\) và \(O(n \log n)\).

---

### 4. Bài tập gợi ý (practice checklist)

- **Bắt buộc**:
  - [ ] Implement `BubbleSort(nums []int)`
  - [ ] Implement `InsertionSort(nums []int)`
  - [ ] Implement `MergeSort(nums []int) []int`
  - [ ] Implement `QuickSort(nums []int)`
- **So sánh hiệu năng**:
  - [ ] Tạo mảng 1k, 10k phần tử random → đo thời gian từng thuật toán.
  - [ ] Tạo mảng đã sort sẵn → xem sort nào chạy nhanh nhất.
  - [ ] Tạo mảng đảo ngược → xem “worst case” trông như thế nào.


