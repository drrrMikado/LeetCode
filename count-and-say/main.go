package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countAndSay(1))
	fmt.Println(countAndSay(7))
	fmt.Println(countAndSay(10))
	fmt.Println(countAndSay(30))
}

func countAndSay(n int) string {
	if n < 1 {
		return "1"
	}
	ans := "1"
	for i := 1; i < n; i++ {
		count := 1
		res := ""
		for j := 0; j < len(ans)-1; j++ {
			if ans[j] == ans[j+1] {
				count++
			} else {
				res += strconv.Itoa(count) + string(ans[j])
				count = 1
			}
		}
		res += strconv.Itoa(count) + string(ans[len(ans)-1])
		ans = res
	}
	return ans
}
