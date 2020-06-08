package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord(""))
	fmt.Println(lengthOfLastWord("a "))
}

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	strs := strings.Split(s, " ")

	return len(strs[len(strs)-1])
}
