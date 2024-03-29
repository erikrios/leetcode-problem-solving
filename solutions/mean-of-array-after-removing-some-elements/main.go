package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(trimMean([]int{1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3}))
	fmt.Println(trimMean([]int{6, 2, 7, 5, 1, 2, 0, 3, 10, 2, 5, 0, 5, 5, 0, 8, 7, 6, 8, 0}))
	fmt.Println(trimMean([]int{6, 0, 7, 0, 7, 5, 7, 8, 3, 4, 0, 7, 8, 1, 6, 8, 1, 1, 2, 4, 8, 1, 9, 5, 4, 3, 8, 5, 10, 8, 6, 6, 1, 0, 6, 10, 8, 2, 3, 4}))
}

func trimMean(arr []int) float64 {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	totalRemove := 5 * len(arr) / 100

	var sum int
	for i := totalRemove; i < len(arr)-totalRemove; i++ {
		sum += arr[i]
	}

	return float64(sum) / float64(len(arr)-totalRemove*2)
}
