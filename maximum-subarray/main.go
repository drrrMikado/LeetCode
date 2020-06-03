package main

import "fmt"

func main() {
	var arr []int
	arr = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(arr))
	fmt.Println(recursive(arr, 0, len(arr)-1))
}

func maxSubArray(nums []int) int {
	sum := 0
	max := nums[0]
	for i := 0; i < len(nums); i++ {
		if sum < 0 {
			sum = nums[i]
		} else {
			sum += nums[i]
		}
		if max <= sum {
			max = sum
		}
	}

	return max
}

func recursive(nums[]int, l,r int) int {
	if l == r {
		return nums[l]
	}
	mid := (l+r)/2

	v1 := recursive(nums, l, mid)
	v2 := recursive(nums,mid+1,r)
	v3 := v1+v2
	if v1 >= v2 {
		if v3 >= v1 {
			return v3
		}
		return v1
	} else {
		if v3 >= v2 {
			return v3
		}
		return v2
	}
}

