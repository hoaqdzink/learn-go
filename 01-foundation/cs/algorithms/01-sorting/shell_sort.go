package main

import "fmt"

// ShellSort sắp xếp mảng theo thứ tự tăng dần (in-place).
// Insertion sort với gap: bắt đầu gap lớn (n/2), giảm dần đến 1.
func ShellSort(nums []int) {
	// TODO: implement shell sort
}

func main() {
	nums := []int{5, 3, 8, 4, 2}
	ShellSort(nums)
	fmt.Println(nums)
}
