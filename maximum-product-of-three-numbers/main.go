package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var a []int
	a = []int{1, 2, 3}
	fmt.Println(a, maximumProduct(a))
	a = []int{1, 2, 3, 4}
	fmt.Println(a, maximumProduct(a))
}

func maximumProduct(nums []int) int {
	sort.Ints(nums)
	v1 := nums[0] * nums[1] * nums[len(nums)-1]
	v2 := nums[len(nums)-1] * nums[len(nums)-2] * nums[len(nums)-3]
	if v1 > v2 {
		return v1
	}
	return v2
}

func maximumProduct2(nums []int) int {
	max1 := math.MinInt32
	max2, max3 := max1, max1
	min1, min2 := math.MaxInt32, math.MaxInt32
	for i := 0; i < len(nums); i++ {
		if min1 >= nums[i] {
			min2 = min1
			min1 = nums[i]
		} else if min2 >= nums[i] {
			min2 = nums[i]
		}
		if max1 <= nums[i] {
			max3, max2 = max2, max1
			max1 = nums[i]
		} else if max2 <= nums[i] {
			max3 = max2
			max2 = nums[i]
		} else if max3 <= nums[i] {
			max3 = nums[i]
		}
	}
	v1 := max1 * min1 * min2
	v2 := max1 * max2 * max3
	if v1 >= v2 {
		return v1
	}
	return v2
}
