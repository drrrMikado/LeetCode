package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)
	nums1 = []int{0}
	nums2 = []int{1}
	merge(nums1, 0, nums2, 1)
	fmt.Println(nums1)
	nums1 = []int{4, 0, 0, 0, 0, 0}
	nums2 = []int{1, 2, 3, 5, 6}
	merge(nums1, 1, nums2, 5)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	// 采用的是双指针+插入位置指针，并从后往前的形式
	end := m + n - 1
	i, j := m-1, n-1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[end] = nums1[i]
			i--
		} else {
			nums1[end] = nums2[j]
			j--
		}
		end--
	}
	for i >= 0 {
		nums1[end] = nums1[i]
		i--
		end--
	}
	for j >= 0 {
		nums1[end] = nums2[j]
		j--
		end--
	}
}
