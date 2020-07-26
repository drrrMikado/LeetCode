package main

import "fmt"

func main() {
	var a [][]int
	a = [][]int{
		{255, 1, 255},
		{255, 1, 255},
		{255, 1, 255},
	}
	fmt.Println(imageSmoother(a))
}

func imageSmoother(M [][]int) [][]int {
	var ans [][]int
	ans = make([][]int, len(M))
	for i := 0; i < len(M); i++ {
		ans[i] = make([]int, len(M[0]))
		for j := 0; j < len(M[0]); j++ {
			count := 0
			keys := [][]int{
				{i, j},
				{i - 1, j},
				{i + 1, j},
				{i, j - 1},
				{i, j + 1},
				{i + 1, j + 1},
				{i - 1, j + 1},
				{i - 1, j - 1},
				{i + 1, j - 1},
			}
			for _, index := range keys {
				if index[0] >= 0 && index[0] < len(M) &&
					index[1] >= 0 && index[1] < len(M[i]) {
					ans[i][j] += M[index[0]][index[1]]
					count++
				}
			}
			ans[i][j] /= count
		}
	}
	return ans
}
