package main

import "fmt"

// RadixSort sắp xếp mảng số nguyên không âm theo thứ tự tăng dần (LSD).
// O(n * d) với d = số chữ số. Dùng counting sort cho từng chữ số.
func RadixSort(nums []int) {
	// TODO: implement radix sort (LSD, base 10)
}

func main() {
	nums := []int{5, 3, 8, 4, 2}
	RadixSort(nums)
	fmt.Println(nums)
}
