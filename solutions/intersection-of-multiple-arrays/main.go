package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(intersection([][]int{{3, 1, 2, 4, 5}, {1, 2, 3, 4}, {3, 4, 5, 6}}))
	fmt.Println(intersection([][]int{{1, 2, 3}, {4, 5, 6}}))
}

func intersection(nums [][]int) []int {
	sets := make(map[int]struct{}, len(nums[0]))

	for i := 0; i < len(nums[0]); i++ {
		sets[nums[0][i]] = struct{}{}
	}

loop:
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		newSets := make(map[int]struct{}, len(num))

		for j := 0; j < len(num); j++ {
			n := num[j]
			newSets[n] = struct{}{}

		}

		for k := range sets {
			if _, ok := newSets[k]; !ok {
				delete(sets, k)
			}

			if len(sets) == 0 {
				break loop
			}
		}
	}

	res := make([]int, 0, len(sets))
	for k := range sets {
		res = append(res, k)
	}

	sort.Slice(res, func(x, y int) bool {
		return res[x] < res[y]
	})

	return res
}
