package main

import "fmt"

func main() {
	var arr []int
	arr = []int{1, 2, 3}
	fmt.Println(plusOne(arr))
	arr = []int{4, 3, 2, 1}
	fmt.Println(plusOne(arr))
	arr = []int{1,9}
	fmt.Println(plusOne(arr))
	arr = []int{9}
	fmt.Println(plusOne(arr))
}

func plusOne(digits []int) []int {
	l := len(digits) - 1
	digits[l] += 1
	for digits[l] >= 10 && l > 0{
		digits[l] = digits[l] % 10
		digits[l-1] += 1
		l--
	}
	if digits[0] >= 10 {
		digits[0] = digits[0] % 10
		digits = append([]int{1}, digits...)
	}
	return digits
}
