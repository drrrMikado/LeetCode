package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(addBinary("11", "1"))
	fmt.Println(addBinary("1010", "1011"))
}

func addBinary(a string, b string) string {
	ans := ""
	ai, bi := len(a)-1, len(b)-1
	carry := 0
	for ai >= 0 || bi >= 0 {
		if ai >= 0 {
			if a[ai] == '1' {
				carry++
			}
			ai--
		}
		if bi >= 0 {
			if b[bi] == '1' {
				carry++
			}
			bi--
		}
		if carry%2 == 0 {
			ans = "0" + ans
		} else {
			ans = "1" + ans
		}
		carry /= 2
	}
	if carry > 0 {
		ans = strconv.FormatInt(int64(carry), 10) + ans
	}
	return ans
}
