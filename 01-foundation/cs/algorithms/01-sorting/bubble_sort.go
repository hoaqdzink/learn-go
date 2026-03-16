package main

import "fmt"
// BubbleSort sắp xếp mảng theo thứ tự tăng dần.
// Tự triển khai thân hàm như bài tập.
func BubbleSort(nums []int) {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums [j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func main() {
	nums := []int{5, 3, 8, 4, 2}
	BubbleSort(nums)
	fmt.Println(nums)
}

