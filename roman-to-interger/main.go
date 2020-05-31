package main

import "fmt"

func main() {
	fmt.Println(romanToInt("III"), romanToInt("IV"), romanToInt("IX"))
	fmt.Println(romanToInt("LVIII"), romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	symbol := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var ans int
	for i := 0; i < len(s)-1; i++ {
		if symbol[s[i]] < symbol[s[i+1]] {
			ans -= symbol[s[i]]
		} else {
			ans += symbol[s[i]]
		}
	}
	ans += symbol[s[len(s)-1]]
	return ans
}
