package main

import "fmt"

func main() {
	var a [][]int
	a = [][]int{{1, 2}, {3, 4}}
	fmt.Println(matrixReshape(a, 1, 4))
	a = [][]int{{1, 2}, {3, 4}}
	fmt.Println(matrixReshape(a, 2, 4))
	a = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(matrixReshape(a, 3, 2))
}

func matrixReshape(nums [][]int, r int, c int) [][]int {
	if len(nums) == 0 || r*c != len(nums)*len(nums[0]) {
		return nums
	}
	res := make([][]int, r)
	for i := 0; i < r; i++ {
		res[i] = make([]int, c)
	}
	rows, cols := 0, 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[0]); j++ {
			res[rows][cols] = nums[i][j]
			cols++
			if cols >= c {
				rows++
				cols = 0
			}
		}
	}

	return res
}

func matrixReshape2(nums [][]int, r int, c int) [][]int {
	if len(nums) == 0 || r*c != len(nums)*len(nums[0]) {
		return nums
	}
	one := []int{}
	for _, v := range nums {
		one = append(one, v...)
	}
	ans := [][]int{}
	s := 0
	for i := 0; i < r; i++ {
		ans = append(ans, one[s*c:(s+1)*c])
		s++
	}
	return ans
}
