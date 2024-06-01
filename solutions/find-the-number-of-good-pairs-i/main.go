package main

import "fmt"

func main() {
	fmt.Println(numberOfPairs([]int{1, 3, 4}, []int{1, 3, 4}, 1))
	fmt.Println(numberOfPairs([]int{1, 2, 4, 12}, []int{2, 4}, 3))
}

func numberOfPairs(nums1 []int, nums2 []int, k int) int {
	var goodPair int

	n := len(nums1)
	m := len(nums2)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if nums1[i]%(nums2[j]*k) == 0 {
				goodPair++
			}
		}
	}

	return goodPair
}
