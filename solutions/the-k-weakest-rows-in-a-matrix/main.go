package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	fmt.Println(
		kWeakestRows(
			[][]int{
				{1, 1, 0, 0, 0},
				{1, 1, 1, 1, 0},
				{1, 0, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{1, 1, 1, 1, 1},
			},
			3,
		),
	)

	fmt.Println(
		kWeakestRows(
			[][]int{
				{1, 0, 0, 0},
				{1, 1, 1, 1},
				{1, 0, 0, 0},
				{1, 0, 0, 0},
			},
			2,
		),
	)

	for i := 0; i < 1_000; i++ {
		val := kWeakestRows(
			[][]int{
				{1, 1, 0},
				{1, 1, 0},
				{1, 1, 1},
				{1, 1, 1},
				{0, 0, 0},
				{1, 1, 1},
				{1, 0, 0},
			},
			6,
		)

		expected := []int{4, 6, 0, 1, 2, 3}

		if !reflect.DeepEqual(val, expected) {
			fmt.Println(val, expected)
		}
	}
}

type counter [][]int

func (c *counter) Insert(k, v int) {
	*c = append(*c, []int{k, v})
}

func (c *counter) Len() int {
	return len(*c)
}

func (c *counter) Less(i, j int) bool {
	if (*c)[i][1] == (*c)[j][1] {
		return (*c)[i][0] < (*c)[j][0]
	}

	return (*c)[i][1] < (*c)[j][1]
}

func (c *counter) Swap(i, j int) {
	(*c)[i], (*c)[j] = (*c)[j], (*c)[i]
}

func kWeakestRows(mat [][]int, k int) []int {
	ctr := make(counter, 0, len(mat))

	for i := 0; i < len(mat); i++ {
		row := mat[i]
		ctr.Insert(i, sum(row))
	}

	sort.Sort(&ctr)

	res := make([]int, k, k)
	for i := 0; i < k; i++ {
		res[i] = ctr[i][0]
	}

	return res
}

func sum(nums []int) int {
	var res int

	for i := 0; i < len(nums); i++ {
		res += nums[i]
	}

	return res
}
