package main

import "fmt"

func main() {
	var a []int
	a = []int{1, 2, 2, 3, 1}
	fmt.Println(a, findShortestSubArray(a))
	a = []int{1, 2, 2, 3, 1, 4, 2}
	fmt.Println(a, findShortestSubArray(a))
	a = []int{1, 3, 2, 2, 3, 1}
	fmt.Println(a, findShortestSubArray(a))
}

func findShortestSubArray(nums []int) int {
	left := map[int]int{}
	right := map[int]int{}
	count := map[int]int{}
	for i := 0; i < len(nums); i++ {
		v := nums[i]
		if _, ok := left[v]; !ok {
			left[v] = i
		}
		right[v] = i
		count[v]++
	}
	ans := len(nums)
	degree := 0
	for _, v := range count {
		if degree <= v {
			degree = v
		}
	}
	for k, v := range count {
		if v == degree && ans > (right[k]-left[k]+1) {
			ans = right[k] - left[k] + 1
		}
	}
	return ans
}
