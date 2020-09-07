package main

import "fmt"

func main() {
	a := [][]int{
		{2, 2, 1},
		{4, 1, 2, 1, 2},
	}
	for _, v := range a {
		fmt.Println(v, singleNumber(v))
	}
}

func singleNumber(nums []int) int {
	ans := 0
	for _, v := range nums {
		ans ^= v
	}
	return ans
}
