package main

import "fmt"

func main() {
	a := [][]int{
		{1, 7, 3, 6, 5, 6},
		{1, 2, 3},
		{0, 0},
		{0, 0, 0},
	}
	for _, v := range a {
		fmt.Println(v, pivotIndex(v))
	}
}

func pivotIndex(nums []int) int {
	leftSum, sum := 0, 0
	for _, v := range nums {
		sum += v
	}
	for k, v := range nums {
		if leftSum == sum-leftSum-v {
			return k
		}
		leftSum += v
	}
	return -1
}

