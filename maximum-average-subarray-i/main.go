package main

import "fmt"

func main() {
	var (
		a []int
		k int
	)
	a, k = []int{1, 3, 12, -5, -6, 50}, 4
	fmt.Println(a, k, findMaxAverage(a, k))
	a, k = []int{5}, 1
	fmt.Println(a, k, findMaxAverage(a, k))
	a, k = []int{1, 3, 12, -5, -6, 50}, 4
	fmt.Println(a, k, findMaxAverage2(a, k))
}

func findMaxAverage(nums []int, k int) float64 {
	max := float64(-10001)
	for i := 0; i < len(nums)-k+1; i++ {
		var v float64
		for j := i; j < i+k; j++ {
			v += float64(nums[j])
		}
		v = v / float64(k)
		if v > max {
			max = v
		}
	}
	return max
}

func findMaxAverage2(nums []int, k int) float64 {
	sum := float64(0)
	for i := 0; i < k; i++ {
		sum += float64(nums[i])
	}
	res := sum
	for i := k; i < len(nums); i++ {
		sum += float64(nums[i]) - float64(nums[i-k])
		if sum > res {
			res = sum
		}
	}
	return res / float64(k)
}
