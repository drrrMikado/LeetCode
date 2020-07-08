package main

import (
	"fmt"
)

func main() {
	var a []int
	a = []int{3, 1, 4, 1, 5}
	fmt.Println(findPairs(a, 2))
	a = []int{1, 2, 3, 4, 5}
	fmt.Println(findPairs(a, 1))
	a = []int{1, 3, 1, 5, 4}
	fmt.Println(findPairs(a, 0))
}

func findPairs(nums []int, k int) int {
	if len(nums) == 0 || k < 0 {
		return 0
	}
	ans := 0
	m := map[int]int{}
	for _, v := range nums {
		m[v] += 1
	}
	for kk, v := range m {
		if k == 0 && v > 1 {
			ans++
		} else if k != 0 {
			if _, ok := m[kk+k]; ok {
				ans++
			}
		}
	}
	return ans
}
