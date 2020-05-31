package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(1<<31-1, -1<<31)
	fmt.Println(reverse(123), reverse(-123), reverse(120))
	fmt.Println(reverse(2147483648), reverse(2147483647))
}

func reverse(x int) int {
	var (
		ans int
		mod int
	)
	for x != 0 {
		mod = x % 10
		x = x / 10
		if ans > math.MaxInt32/10 || (ans == math.MaxInt32/10 && mod > 7) {
			return 0
		}
		if ans < math.MinInt32/10 || (ans == math.MinInt32/10 && mod < -8) {
			return 0
		}
		ans = ans*10 + mod
	}

	return ans
}
