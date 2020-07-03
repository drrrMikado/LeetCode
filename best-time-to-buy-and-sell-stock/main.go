package main

import (
	"fmt"
	"math"
)

func main() {
	var a []int
	a = []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(a))
	a = []int{7, 6, 5, 4, 3, 1}
	fmt.Println(maxProfit(a))
	a = []int{1, 2, 3, 4, 5, 6, 3, 2, 1}
	fmt.Println(maxProfit(a))
}

func maxProfit(prices []int) int {
	max, min := 0, math.MaxInt32
	for i := 0; i < len(prices); i++ {
		if min > prices[i] {
			min = prices[i]
		} else if prices[i]-min > max {
			max = prices[i] - min
		}
	}
	return max
}
