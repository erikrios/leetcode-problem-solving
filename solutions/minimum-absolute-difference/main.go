package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumAbsDifference([]int{4, 2, 1, 3}))
	fmt.Println(minimumAbsDifference([]int{1, 3, 6, 10, 15}))
	fmt.Println(minimumAbsDifference([]int{3, 8, -10, 23, 19, -4, -14, 27}))
}

func minimumAbsDifference(arr []int) [][]int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	n := len(arr)

	minDiff := 1_000_000
	results := make([][]int, 0, n/2)
	for i := 0; i < n-1; i++ {
		cur := arr[i]
		next := arr[i+1]

		abs := next - cur
		if abs <= minDiff {
			if abs < minDiff {
				minDiff = abs
				results = make([][]int, 0, n/2)
			}
			results = append(results, []int{cur, next})
		}
	}

	return results
}
