package main

import "fmt"

func main() {
	fmt.Println("fib(0):", fib(0))
	fmt.Println("fib(3):", fib(3))
	fmt.Println("fib(30):", fib(30))
	fmt.Println("fib2(30):", fib2(30))
}

func fib(N int) int {
	if N <= 1 {
		return N
	}
	return fib(N-1) + fib(N-2)
}

func fib2(N int) int {
	if N <= 1 {
		return N
	}
	if N == 2 {
		return 1
	}
	cur, prev1, prev2 := 0, 1, 1
	for i := 3; i <= N; i++ {
		cur = prev1 + prev2
		prev1, prev2 = cur, prev1
	}
	return cur
}
