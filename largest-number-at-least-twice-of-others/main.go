package main

import "fmt"

func main() {
	a := [][]int{
		{3, 6, 1, 0},
		{1, 2, 3, 4},
		{1},
		{0, 0, 0, 0, 1},
	}
	for _, v := range a {
		fmt.Println(v, dominantIndex(v))
	}
}

func dominantIndex(nums []int) int {
	maxIndex := -1
	for k, v := range nums {
		if nums[maxIndex] < v {
			maxIndex = k
		}
	}
	for k, v := range nums {
		if k != maxIndex && v*2 > nums[maxIndex] {
			return -1
		}
	}
	return i
}
