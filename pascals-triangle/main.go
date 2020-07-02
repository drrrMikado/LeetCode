package main

import "fmt"

func main() {
	fmt.Println(generate(5))
}

func generate(numRows int) [][]int {
	ans := [][]int{}
	for i := 0; i < numRows; i++ {
		b := make([]int, i+1)
		b[0] = 1
		for j := 1; j < i; j++ {
			b[j] = ans[i-1][j-1] + ans[i-1][j]
		}
		b[i] = 1
		ans = append(ans, b)
	}
	return ans
}
