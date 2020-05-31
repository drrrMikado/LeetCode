package main

import "fmt"

func main() {
	arr := []int{2, 7, 11, 15}
	fmt.Println(twoSum(arr, 9))
}

func twoSum(nums []int, target int) []int {
	// map表 key 为数组的值，value 为数组的下标
	m := map[int]int{}
	for k, v := range nums {
		c := target - v
		if _, ok := m[c]; ok && m[c] != k {
			return []int{k, m[c]}
		}
		m[v] = k
	}
	return []int{}
}
