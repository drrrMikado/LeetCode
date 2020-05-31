package main

import "fmt"

func main() {
	var arr []int
	arr = []int{3, 2, 2, 3}
	fmt.Println(removeElement(arr, 3), arr)
	arr = []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(arr, 2), arr)
}

func removeElement(nums []int, val int) int {
	l := len(nums)
	for i := 0; i < l; i++ {
		if nums[i] == val {
			nums[i], nums[l-1] = nums[l-1], nums[i]
			l--
			i--
		}
	}
	return l
}
