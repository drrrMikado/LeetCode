package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPalindrome(121), isPalindrome(-121), isPalindrome(10))
}

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	var i int
	for x > i {
		i = i*10 + x%10
		x = x / 10
	}
	return x == i || x == i/10
}
