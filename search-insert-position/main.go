package main

import "fmt"

func main() {
	var arr []int
	arr = []int{1, 3, 5, 6}
	fmt.Println(searchInsert(arr, 5))
	arr = []int{1, 3, 5, 6}
	fmt.Println(searchInsert(arr, 2))
	arr = []int{1, 3, 5, 6}
	fmt.Println(searchInsert(arr, 7))
	arr = []int{1, 3, 5, 6}
	fmt.Println(searchInsert(arr, 0))
}

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	for k, v := range nums {
		if v == target {
			return k
		} else if v > target {
			return k
		}
	}
	return len(nums)
}
