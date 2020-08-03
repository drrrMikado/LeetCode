package main

import "fmt"

func main() {
	var a []int
	a = []int{1, 2, 3, 1}
	fmt.Println(a, rob(a))
	a = []int{2, 7, 9, 3, 1}
	fmt.Println(a, rob(a))
	a = []int{}
	fmt.Println(a, rob(a))
	a = []int{0, 1}
	fmt.Println(a, rob(a))
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(dp[0], dp[1])
	for i := 2; i < len(dp); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[len(nums)-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
