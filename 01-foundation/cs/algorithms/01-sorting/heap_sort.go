package main

import "fmt"

// HeapSort sắp xếp mảng theo thứ tự tăng dần (in-place) dùng max-heap.
// O(n log n), space O(1). Không stable.
func HeapSort(nums []int) {
	// TODO: implement heap sort
}

func main() {
	nums := []int{5, 3, 8, 4, 2}
	HeapSort(nums)
	fmt.Println(nums)
}
