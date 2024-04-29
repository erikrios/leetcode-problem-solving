package main

import "fmt"

func main() {
	fmt.Println(addedInteger([]int{2, 6, 4}, []int{9, 7, 5}))
	fmt.Println(addedInteger([]int{10}, []int{5}))
	fmt.Println(addedInteger([]int{1, 1, 1, 1}, []int{1, 1, 1, 1}))
}

func addedInteger(nums1 []int, nums2 []int) int {
	n := len(nums1)

	var n1Sum, n2Sum int
	for i := 0; i < n; i++ {
		n1Sum += nums1[i]
		n2Sum += nums2[i]
	}

	diff := n2Sum - n1Sum

	return diff / n
}
