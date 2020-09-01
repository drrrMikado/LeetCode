package main

import (
	"fmt"
	"strings"
)

func main() {
	a := []string{
		"A man, a plan, a canal: Panama",
		"race a car",
		"0P",
	}
	for _, v := range a {
		fmt.Println(v, isPalindrome(v))
	}
}

func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}
	s = strings.ToLower(s)
	i, j := 0, len(s)-1
	for i < j {
		for i < j && !isLetterOrNumber(s[i]) {
			i++
		}
		for i < j && !isLetterOrNumber(s[j]) {
			j--
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func isLetterOrNumber(b byte) bool {
	return (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9')
}
