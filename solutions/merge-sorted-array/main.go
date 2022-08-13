package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	result := make([]int, m+n)

	var p1, p2 int

	for i := 0; i < m+n; i++ {
		if p1 >= m {
			result[i] = nums2[p2]
			p2++
			continue
		}

		if p2 >= n {
			result[i] = nums1[p1]
			p1++
			continue
		}

		if nums1[p1] < nums2[p2] {
			result[i] = nums1[p1]
			p1++
		} else {
			result[i] = nums2[p2]
			p2++
		}
	}

	copy(nums1, result)
}
