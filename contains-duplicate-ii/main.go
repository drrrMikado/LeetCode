package main

import "fmt"

func main() {
	var a []int
	a = []int{1, 2, 3, 1}
	fmt.Println(containsNearbyDuplicate(a, 3))
	a = []int{1, 0, 1, 1}
	fmt.Println(containsNearbyDuplicate(a, 1))
	a = []int{1, 2, 3, 1, 2, 3}
	fmt.Println(containsNearbyDuplicate(a, 2))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) <= 1 || k == 0 {
		return false
	}
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok && i-m[nums[i]] <= k {
			return true
		}
		m[nums[i]] = i
	}
	return false
}
