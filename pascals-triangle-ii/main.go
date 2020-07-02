package main

import "fmt"

func main() {
	//fmt.Println(getRow(3))
	fmt.Println(getRow(8))
	fmt.Println(getRow(9))
	//fmt.Println(getRow(33))
	fmt.Println(getRow2(8))
	fmt.Println(getRow2(9))
}

func getRow(rowIndex int) []int {
	var prev []int
	for i := 0; i <= rowIndex; i++ {
		ans := make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				ans[j] = 1
			} else {
				ans[j] = prev[j-1] + prev[j]
			}
		}
		prev = ans
	}
	return prev
}

// O(k) extra space
func getRow2(rowIndex int) []int {
	ans := []int{1}
	for i := 1; i <= rowIndex; i++ {
		tmp := 0
		for j := 0; j < i; j++ {
			ans[j], tmp = ans[j]+tmp, ans[j]
		}
		ans = append(ans, 1)
	}
	return ans
}
