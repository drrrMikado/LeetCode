package main

import "fmt"

func main() {
	var (
		f []int
		n int
	)
	//f, n = []int{1, 0, 0, 0, 1}, 1
	//fmt.Println(f, n, canPlaceFlowers(f, n))
	//f, n = []int{1, 0, 0, 0, 1}, 2
	//fmt.Println(f, n, canPlaceFlowers(f, n))
	f, n = []int{0, 0, 0, 0, 0}, 3
	fmt.Println(f, n, canPlaceFlowers(f, n))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed); i++ {
		if (i == 0 || flowerbed[i-1] == 0) &&
			flowerbed[i] == 0 &&
			(i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			n--
			i++
		}
		if n == 0 {
			break
		}
	}
	return n <= 0
}
