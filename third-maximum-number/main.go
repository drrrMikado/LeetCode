package main

import (
	"fmt"
	"math"
)

func main() {
	var a []int
	a = []int{3, 2, 1}
	fmt.Println(thirdMax(a))
	a = []int{1, 2}
	fmt.Println(thirdMax(a))
	a = []int{2, 2, 3, 1}
	fmt.Println(thirdMax(a))
}

func thirdMax(nums []int) int {
	m1, m2, m3 := math.MinInt32, math.MinInt32, math.MinInt32
	for _, v := range nums {
		if v > m1 {
			m2, m3 = m1, m2
			m1 = v
		} else if v > m2 && v < m1 {
			m3 = m2
			m2 = v
		} else if v > m3 && v < m2 {
			m3 = v
		}
	}
	if m3 == math.MinInt32 {
		return m1
	}
	return m3
}
