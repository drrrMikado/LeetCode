package main

import "fmt"

func main() {
	a := [][]int{
		{1, 0, 0},    // true
		{1, 1, 1, 0}, // false
	}
	for _, v := range a {
		fmt.Println(v,
			isOneBitCharacter(v),
			isOneBitCharacter2(v),
			isOneBitCharacter3(v))
	}
}

func isOneBitCharacter(bits []int) bool {
	b := 0
	for i := 0; i < len(bits); i++ {
		if bits[i] == 1 {
			i++
			b = 2
		} else {
			b = 1
		}
	}
	return b == 1
}

func isOneBitCharacter2(bits []int) bool {
	i := 0
	for i < len(bits)-1 {
		i += bits[i] + 1
	}
	return i == len(bits)-1
}

func isOneBitCharacter3(bits []int) bool {
	i := len(bits) - 2
	for i >= 0 && bits[i] > 0 {
		i--
	}
	return (len(bits)-i)%2 == 0
}
