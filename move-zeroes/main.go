package main

import "fmt"

func main() {
	var a []int
	a = []int{0, 1, 0, 3, 12}
	moveZeroes(a)
	fmt.Println(a)
}

func moveZeroes(nums []int) {
	p := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			p = i
			break
		}
	}
	if p == -1 {
		return
	}
	for i := p + 1; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[p] = nums[p], nums[i]
			p++
		}
	}
}
