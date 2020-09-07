package main

import "fmt"

func main() {
	a := [][][]int{
		[][]int{
			{1, 2, 3, 4},
			{5, 1, 2, 3},
			{9, 5, 1, 2},
		},
		[][]int{
			{1, 2},
			{2, 2},
		},
		[][]int{},
	}
	for _, v := range a {
		fmt.Println(v, isToeplitzMatrix(v))
	}
}

func isToeplitzMatrix(matrix [][]int) bool {
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] != matrix[i-1][j-1] {
				return false
			}
		}
	}
	return true
}
