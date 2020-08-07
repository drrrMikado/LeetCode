package main

import "fmt"

func main() {
	var a []int
	var obj NumArray
	a = []int{-2, 0, 3, -5, 2, -1}
	obj = Constructor(a)
	fmt.Println(obj.SumRange(0, 2))
	fmt.Println(obj.SumRange(2, 5))
	fmt.Println(obj.SumRange(0, 5))
}

type NumArray struct {
	Sum []int
}

func Constructor(nums []int) NumArray {
	if len(nums) == 0 {
		return NumArray{}
	}
	sum := make([]int, len(nums))
	sum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sum[i] = nums[i] + sum[i-1]
	}
	return NumArray{sum}
}

func (this *NumArray) SumRange(i int, j int) int {
	if i > 0 {
		return this.Sum[j] - this.Sum[i-1]
	}
	return this.Sum[j]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
