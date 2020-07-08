package main

import "fmt"

func main() {
	var a []int
	a = []int{1, 1, 0, 1, 1, 1}
	fmt.Println(findMaxConsecutiveOnes(a))
}

func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	tmp := 0
	for _, v := range nums {
		if v == 1 {
			tmp++
		} else {
			tmp = 0
		}
		if max < tmp {
			max = tmp
		}
	}
	return max
}
