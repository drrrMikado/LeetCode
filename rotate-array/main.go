package main

import "fmt"

func main() {
	var a []int
	a = []int{1, 2, 3, 4, 5, 6, 7}
	rotate(a, 3)
	fmt.Println(a)
	a = []int{-1, -100, 3, 99}
	rotate(a, 2)
	fmt.Println(a)
	a = []int{-1, -100, 3, 99}
	rotateWithCycleReplace(a, 2)
	fmt.Println(a)
}

func rotate(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return
	}
	l, r := 0, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
	l, r = 0, k-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
	l, r = k, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

func rotateWithCycleReplace(nums []int, k int) {
	k = k % len(nums)
	count := 0
	for i := 0; count < len(nums); i++ {
		current := i
		prev := nums[i]
		for {
			next := (current + k) % len(nums)
			nums[next], prev = prev, nums[next]
			current = next
			count++
			if i == current { // cycle end point
				break
			}
		}
	}
}
