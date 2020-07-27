package main

import "fmt"

func main() {
	var a []int
	a = []int{1, 3, 5, 4, 7}
	fmt.Println(a, findLengthOfLCIS(a))
	a = []int{1, 1, 3, 4, 1, 2, 3, 4}
	fmt.Println(a, findLengthOfLCIS(a))
}

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max, ans := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			ans++
		} else {
			ans = 1
		}
		if max < ans {
			max = ans
		}
	}
	return max
}
