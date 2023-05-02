package main

import "fmt"

func main() {
	fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}

func intersection(nums1 []int, nums2 []int) []int {
	n := len(nums1)
	m := len(nums2)
	mapper := make(map[int]bool, n+m)

	for i := 0; i < n; i++ {
		num := nums1[i]
		mapper[num] = false
	}

	for i := 0; i < m; i++ {
		num := nums2[i]
		if _, ok := mapper[num]; ok {
			mapper[num] = true
		}
	}

	results := make([]int, 0, n+m)
	for k, v := range mapper {
		if v {
			results = append(results, k)
		}
	}

	return results
}
