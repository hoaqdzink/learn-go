## CS — Data Structures

### 1. Mục tiêu

- Hiểu **ý tưởng, ưu/nhược điểm, độ phức tạp** của từng cấu trúc dữ liệu.
- Biết **khi nào dùng cái nào** trong bài toán thực tế / backend.
- Cài đặt được **version đơn giản** (bằng Go hoặc pseudo) cho từng cấu trúc.

### 2. Thứ tự học đề xuất

1. **Array & Slice**
   - Khái niệm: vùng nhớ liên tục, index, capacity, length.
   - Sự khác nhau giữa **array** và **slice** (đặc biệt trong Go).
   - Chi phí: truy cập \(O(1)\), insert/delete giữa mảng \(O(n)\).
2. **Linked List**
   - Singly vs Doubly linked list.
   - Con trỏ `next`, `prev`, chi phí truy cập tuần tự.
   - So sánh với array/slice.
3. **Stack**
   - LIFO, các thao tác: `push`, `pop`, `peek`.
   - Ứng dụng: undo, parsing, DFS.
4. **Queue**
   - FIFO, các thao tác: `enqueue`, `dequeue`, `front`.
   - Ứng dụng: BFS, message queue.
5. **Hash Map (Hash Table)**
   - Key–value, hash function, bucket.
   - Xử lý collision: chaining, open addressing (chỉ cần nắm idea).
   - Độ phức tạp trung bình: \(O(1)\) cho get/set.
6. **Tree**
   - Khái niệm node, root, leaf, height.
   - Binary Tree, Binary Search Tree (BST) cơ bản.
   - Traversal: preorder, inorder, postorder.
7. **Heap**
   - Min-heap, max-heap.
   - Operations: `insert`, `extract-min/max`.
   - Ứng dụng: priority queue, Dijkstra.
8. **Graph**
   - Vertex, edge, directed vs undirected, weighted vs unweighted.
   - Biểu diễn: adjacency list, adjacency matrix.

### 3. Gợi ý tài nguyên (tùy chọn)

- CLRS (Introduction to Algorithms) — các chương về data structures cơ bản.
- Visualgo, cs.usfca.edu (các visualization của tree, heap, graph).

### 4. Bài tập gợi ý

**Array & Slice**
- Viết hàm:
  - `Reverse([]int) []int`
  - `RemoveIndex([]int, idx int) []int` (không dùng thư viện ngoài).
- Phân tích độ phức tạp các hàm trên.

**Linked List**
- Cài đặt singly linked list:
  - Thêm node vào đầu/cuối.
  - Xóa node theo giá trị.
- Viết hàm **revert linked list** (đảo ngược).

**Stack**
- Cài đặt stack dùng slice.
- Dùng stack để:
  - Kiểm tra chuỗi ngoặc `()[]{}` có hợp lệ không.

**Queue**
- Cài đặt queue dùng slice hoặc linked list.
- Dùng queue để implement **BFS** trên graph đơn giản.

**Hash Map**
- Tự cài đặt 1 hash map đơn giản:
  - Dùng array of bucket, mỗi bucket là linked list.
  - Hỗ trợ `Put`, `Get`, `Delete`.

**Tree & Heap**
- Cài đặt BST đơn giản:
  - Thêm node.
  - Tìm kiếm 1 giá trị.
- Dùng heap (hoặc priority queue) để:
  - Tìm `k` phần tử lớn nhất trong 1 mảng.

**Graph**
- Biểu diễn graph bằng adjacency list.
- Viết hàm:
  - In ra tất cả node kề với 1 node cho trước.

