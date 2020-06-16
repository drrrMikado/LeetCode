package main

import "fmt"

func main() {
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs(45))
}

func climbStairs(n int) int {
	p, q, r := 0, 0, 1
	for i := 1; i <= n; i++ {
		p = q
		q = r
		r = q + p
	}
	return r
}

// time out 
func oldClimbStairs(n int) int {
	if n <= 2 {
		return 2
	}
	return climbStairs(n-1) + climbStairs(n-2)
}
