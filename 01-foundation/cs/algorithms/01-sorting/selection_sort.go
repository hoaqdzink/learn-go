package main

import "fmt"

// SelectionSort sắp xếp mảng theo thứ tự tăng dần (in-place).
// Mỗi lượt chọn phần tử nhỏ nhất trong đoạn chưa sort, đổi chỗ với phần tử đầu đoạn.
// O(n²) mọi trường hợp.
func SelectionSort(nums []int) {
	// TODO: implement selection sort
}

func main() {
	nums := []int{5, 3, 8, 4, 2}
	SelectionSort(nums)
	fmt.Println(nums)
}
