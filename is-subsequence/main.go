package main

import "fmt"

func main() {
	var s, t string
	s, t = "abc", "ahbgdc"
	fmt.Println(s, t, isSubsequence(s, t))
	s, t = "axc", "ahbgdc"
	fmt.Println(s, t, isSubsequence(s, t))
	s, t = "acb", "ahbgdc"
	fmt.Println(s, t, isSubsequence(s, t))
	s, t = "aaaaaa", "bbaaaa"
	fmt.Println(s, t, isSubsequence(s, t))
}

func isSubsequence(s string, t string) bool {
	index, flag := 0, false
	for i := 0; i < len(s); i++ {
		flag = false
		for ; index < len(t); index++ {
			if s[i] == t[index] {
				flag = true
				index++
				break
			}
		}
		if !flag {
			return false
		}
	}
	return true
}
