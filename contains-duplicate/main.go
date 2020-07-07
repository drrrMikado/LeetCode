package main

import "fmt"

func main() {
	var a []int
	a = []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(a))
	a = []int{1, 2, 3, 4}
	fmt.Println(containsDuplicate(a))
}

func containsDuplicate(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}
	m := map[int]bool{}
	for _, v := range nums {
		if _, ok := m[v]; !ok {
			m[v] = true
		} else {
			return true
		}
	}
	return false
}
