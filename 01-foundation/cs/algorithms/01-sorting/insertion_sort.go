package main

import "fmt"

// InsertionSort sắp xếp mảng theo thứ tự tăng dần.
// Tự triển khai thân hàm như bài tập.
func InsertionSort(nums []int) {
	for i := 1; i < len(nums); i++{
		key := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > key{
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = key
	}
}

func main(){
	nums := []int{5, 3, 8, 4, 2}
	InsertionSort(nums)
	fmt.Println(nums)
}