package main

import (
	"fmt"
)

func main() {
	var a []int
	a = []int{4, 2, 3}
	fmt.Println(a, checkPossibility(a))
	a = []int{4, 2, 1}
	fmt.Println(a, checkPossibility(a))
	a = []int{3, 4, 2, 3}
	fmt.Println(a, checkPossibility(a))
	a = []int{-1, 4, 2, 3}
	fmt.Println(a, checkPossibility(a))
	a = []int{1, 3, 5, 2, 4}
	fmt.Println(a, checkPossibility(a))
	a = []int{2, 3, 3, 2, 4}
	fmt.Println(a, checkPossibility(a))
	a = []int{3, 3, 2, 2}
	fmt.Println(a, checkPossibility(a))
	a = []int{1, 2, 4, 5, 3}
	fmt.Println(a, checkPossibility(a))
}

func checkPossibility(nums []int) bool {
	return checkPossibility2(nums)
	p := -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			if p != -1 {
				return false
			}
			p = i
		}
	}
	return p == -1 || p == 0 || p == len(nums)-2 ||
		nums[p-1] <= nums[p+1] || nums[p] <= nums[p+2]
}

