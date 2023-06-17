package main

import "fmt"

func main() {
	fmt.Println(findDifference([]int{1, 2, 3}, []int{2, 4, 6}))
	fmt.Println(findDifference([]int{1, 2, 3, 3}, []int{1, 1, 2, 2}))
}

func findDifference(nums1 []int, nums2 []int) [][]int {
	n := len(nums1)
	nums1Set := make(map[int]struct{}, n)
	for i := 0; i < n; i++ {
		num := nums1[i]
		nums1Set[num] = struct{}{}
	}

	m := len(nums2)
	nums2Set := make(map[int]struct{}, m)
	for i := 0; i < m; i++ {
		num := nums2[i]
		nums2Set[num] = struct{}{}
	}

	res1 := make([]int, 0, len(nums1Set))
	for k := range nums1Set {
		if _, ok := nums2Set[k]; !ok {
			res1 = append(res1, k)
		}
	}

	res2 := make([]int, 0, len(nums2Set))
	for k := range nums2Set {
		if _, ok := nums1Set[k]; !ok {
			res2 = append(res2, k)
		}
	}

	return [][]int{res1, res2}
}
