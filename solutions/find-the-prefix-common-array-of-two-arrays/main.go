package main

import "fmt"

func main() {
	fmt.Println(findThePrefixCommonArray([]int{1, 3, 2, 4}, []int{3, 1, 2, 4}))
	fmt.Println(findThePrefixCommonArray([]int{2, 3, 1}, []int{3, 1, 2}))
}

func findThePrefixCommonArray(A []int, B []int) []int {
	n := len(A)
	mapper := make(map[int]struct{})

	var count int
	results := make([]int, 0, n)
	for i := 0; i < n; i++ {
		aVal, bVal := A[i], B[i]

		if _, ok := mapper[aVal]; !ok {
			mapper[aVal] = struct{}{}
		} else {
			count++
		}

		if _, ok := mapper[bVal]; !ok {
			mapper[bVal] = struct{}{}
		} else {
			count++
		}

		results = append(results, count)
	}

	return results
}
