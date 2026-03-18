package main

import "fmt"

// CountingSort trả về mảng mới đã sắp xếp tăng dần.
// Giả sử mọi phần tử trong nums thuộc [0, maxVal]. O(n + maxVal).
func CountingSort(nums []int, maxVal int) []int {
	// TODO: implement counting sort
	
}

func main() {
	nums := []int{5, 3, 8, 4, 2}
	out := CountingSort(nums, 10)
	fmt.Println(out)
}
