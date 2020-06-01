package main

import "fmt"

func main() {
	fmt.Println(strStr("hello", "ll"))
	fmt.Println(strStr("aaaaa", "bba"))
}

func strStr(haystack string, needle string) int {
	l := len(needle)
	if l == 0 {
		return 0
	}
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] &&
			i+l <= len(haystack) &&
			haystack[i:i+l] == needle[:] {
			return i
		}
	}
	return -1
}
