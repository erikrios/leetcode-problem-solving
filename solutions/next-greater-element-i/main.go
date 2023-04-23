package main

import "fmt"

func main() {
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
	fmt.Println(nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4}))
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	n := len(nums2)
	mapper := make(map[int]int, n)

	for i := 0; i < n; i++ {
		num := nums2[i]
		mapper[num] = i
	}

	m := len(nums1)
	results := make([]int, 0, m)
	for i := 0; i < m; i++ {
		num := nums1[i]

		res := -1
		for j := mapper[num] + 1; j < n; j++ {
			num2 := nums2[j]
			if num2 > num {
				res = num2
				break
			}
		}

		results = append(results, res)
	}

	return results
}
