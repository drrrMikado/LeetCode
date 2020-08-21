package main

import "fmt"

func main() {
	a := [][]int{
		{10, 15, 20},
		{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
		{0, 0},
		{1, 0, 1},
	}
	for _, v := range a {
		fmt.Println(v, minCostClimbingStairs2(v))
	}
}

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+1)
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < len(cost); i++ {
		dp[i] = cost[i] + min(dp[i-1], dp[i-2])
	}
	return min(dp[len(cost)-1], dp[len(cost)-2])
}

func minCostClimbingStairs2(cost []int) int {
	dp1, dp2 := cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		dp2, dp1 = cost[i]+min(dp1, dp2), dp2
	}
	return min(dp1, dp2)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
