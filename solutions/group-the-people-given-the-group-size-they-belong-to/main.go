package main

import "fmt"

func main() {
	fmt.Println(groupThePeople([]int{3, 3, 3, 3, 3, 1, 3}))
	fmt.Println(groupThePeople([]int{2, 1, 3, 3, 3, 2}))
}

func groupThePeople(groupSizes []int) [][]int {
	groups := make(map[int][]int)

	for i := 0; i < len(groupSizes); i++ {
		size := groupSizes[i]
		groups[size] = append(groups[size], i)
	}

	res := make([][]int, 0)
	for k, v := range groups {
		var idx int

		groupNum := len(v) / k
		for i := 0; i < groupNum; i++ {
			vals := make([]int, 0, k)

			for ; idx < len(v)/groupNum*(i+1); idx++ {
				vals = append(vals, v[idx])
			}

			res = append(res, vals)
		}
	}

	return res
}
