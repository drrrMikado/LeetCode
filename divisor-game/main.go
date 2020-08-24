package main

import (
	"fmt"
)

func main() {
	t := []int{2, 3, 4}
	for _, v := range t {
		fmt.Println(v, divisorGame(v))
	}
}

func divisorGame(N int) bool {
	f := make(map[int]bool, N+1)
	f[1], f[2] = false, true
	for i := 3; i <= N; i++ {
		for j := 1; j < i; j++ {
			if i%j == 0 && !f[i-j] {
				f[i] = true
				break
			}
		}
	}
	return f[N]
}

func divisorGame2(N int) bool {
	return N%2 == 0
}
