package main

import "fmt"

func main() {
	fmt.Println(findIntersectionValues([]int{4, 3, 2, 3, 1}, []int{2, 2, 5, 2, 3, 6}))
	fmt.Println(findIntersectionValues([]int{3, 4, 2, 3}, []int{1, 5}))
}

func findIntersectionValues(nums1, nums2 []int) []int {
	nums1Mapper := make(map[int]struct{})
	nums2Mapper := make(map[int]struct{})

	for _, v := range nums1 {
		nums1Mapper[v] = struct{}{}
	}

	for _, v := range nums2 {
		nums2Mapper[v] = struct{}{}
	}

	results := make([]int, 2)

	for _, val := range nums1 {
		if _, ok := nums2Mapper[val]; ok {
			results[0]++
		}
	}

	for _, val := range nums2 {
		if _, ok := nums1Mapper[val]; ok {
			results[1]++
		}
	}

	return results
}
