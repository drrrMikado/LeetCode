package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var a []int
	a = []int{2, 6, 4, 8, 10, 9, 15}
	fmt.Println(a, findUnsortedSubarray(a))
	a = []int{2, 3, 4}
	fmt.Println(a, findUnsortedSubarray(a))
	a = []int{1, 2, 3, 3, 3}
	fmt.Println(a, findUnsortedSubarray(a))
	a = []int{1, 3, 2, 2, 2}
	fmt.Println(a, findUnsortedSubarray(a))
	a = []int{1, 3, 2, 2, 2}
	fmt.Println(a, findUnsortedSubarray2(a))
	a = []int{1, 3, 2, 4, 5}
	fmt.Println(a, findUnsortedSubarray2(a))
	a = []int{1, 3, 2, 4, 5}
	fmt.Println(a, findUnsortedSubarray(a))
}

func findUnsortedSubarray(nums []int) int {
	t := make([]int, len(nums))
	copy(t, nums)
	sort.Ints(t)
	ans := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != t[i] {
			ans = i
			break
		}
	}
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] != t[i] {
			ans = i - ans + 1
			break
		}
	}
	return ans
}

func findUnsortedSubarray2(nums []int) int {
	max, min := math.MinInt32, math.MaxInt32
	flag := false
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			flag = true
		}
		if flag && min > nums[i] {
			min = nums[i]
		}
	}
	flag = false
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i-1] > nums[i] {
			flag = true
		}
		if flag && max < nums[i-1] {
			max = nums[i-1]
		}
	}
	l, r := 0, len(nums)-1
	for ; l < len(nums); l++ {
		if nums[l] > min {
			break
		}
	}
	for ; r >= 0; r-- {
		if nums[r] < max {
			break
		}
	}
	fmt.Println(l, r, max, min)
	if r < l {
		return 0
	}
	return r - l + 1
}
