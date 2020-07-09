package main

import (
	"fmt"
	"sort"
)

func main() {
	var a []int
	a = []int{1, 4, 3, 2}
	fmt.Println(arrayPairSum(a)) // output 4
}

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	sum := 0
	for i := 0; i < len(nums); i += 2 {
		sum += nums[i]
	}
	return sum
}
