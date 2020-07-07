package main

import "fmt"

func main() {
	var a []int
	a = []int{3, 0, 1}
	fmt.Println(missingNumber(a))
	a = []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	fmt.Println(missingNumber(a))
	a = []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	fmt.Println(missingNumber2(a))
}

func missingNumber(nums []int) int {
	m := map[int]bool{}
	for _, v := range nums {
		m[v] = true
	}
	for i := 0; i <= len(nums); i++ {
		if _, ok := m[i]; !ok {
			return i
		}
	}
	return -1
}

func missingNumber2(nums []int) int {
	missing := len(nums)
	for i := 0; i < len(nums); i++ {
		missing ^= i ^ nums[i]
	}
	return missing
}
