package main

import "fmt"

func main() {
	var a []int
	a = []int{3, 2, 3}
	fmt.Println(majorityElement(a))
	a = []int{2,2,1,1,1,1,2,2}
	fmt.Println(majorityElement(a))
}

func majorityElement(nums []int) int {
	return majorityElementRec(nums, 0, len(nums)-1)
}

func majorityElementRec(nums []int, l, r int) int {
	if l >= r {
		return nums[l]
	}

	mid := (r-l)/2 + l
	left := majorityElementRec(nums, l, mid)
	right := majorityElementRec(nums, mid+1, r)
	if left == right {
		return left
	}

	leftCount := count(nums, left, l, r)
	rightCount := count(nums, right, l, r)

	if rightCount > leftCount {
		return right
	} else {
		return left
	}
}

func count(nums []int, num, l, r int) int {
	count := 0
	for _, v := range nums {
		if v == num {
			count++
		}
	}
	return count
}
