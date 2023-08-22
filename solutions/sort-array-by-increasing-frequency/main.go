package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(frequencySort([]int{1, 1, 2, 2, 2, 3}))
	fmt.Println(frequencySort([]int{2, 3, 1, 3, 2}))
	fmt.Println(frequencySort([]int{1, 1, -6, 4, 5, -6, 1, 4, 1}))
}

func frequencySort(nums []int) []int {
	n := len(nums)
	numCounters := make(map[int]int, n)

	for i := 0; i < n; i++ {
		num := nums[i]
		numCounters[num]++
	}

	values := make([][]int, 101, 101)
	for k, v := range numCounters {
		for i := 0; i < v; i++ {
			values[v] = append(values[v], k)
		}
	}

	results := make([]int, 0, n)
	for i := 0; i < len(values); i++ {
		value := values[i]
		if len(value) == 0 {
			continue
		}

		if len(value) == 1 {
			results = append(results, value[0])
			continue
		}

		sort.Slice(values[i], func(k, l int) bool {
			return values[i][k] > values[i][l]
		})

		results = append(results, values[i]...)
	}

	return results
}
