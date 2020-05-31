package main

import "fmt"

func main() {
	fmt.Println(isValid("()"), isValid("()[]{}"), isValid("(]"))
	fmt.Println(isValid("([)]"), isValid("{[]}"), isValid(""))
}

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	} else if len(s) % 2 != 0 {
		return false
	}
	lSymbol := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	rSymbol := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if _,ok := lSymbol[s[i]]; ok {
			stack = append(stack, s[i])
		} else if _,ok := rSymbol[s[i]]; ok {
			if len(stack) == 0 {
				return false
			}
			last := stack[len(stack)-1]
			if lSymbol[last] != s[i] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}
