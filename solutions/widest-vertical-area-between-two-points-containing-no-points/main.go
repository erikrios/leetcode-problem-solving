package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(
		maxWidthOfVerticalArea([][]int{
			{8, 7},
			{9, 9},
			{7, 4},
			{9, 7},
		}),
	)

	fmt.Println(
		maxWidthOfVerticalArea([][]int{
			{3, 1},
			{9, 0},
			{1, 0},
			{1, 4},
			{5, 3},
			{8, 8},
		}),
	)
}

func maxWidthOfVerticalArea(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	var maxWidth int
	for i := 1; i < len(points); i++ {
		point := points[i][0]
		prevPoint := points[i-1][0]

		width := point - prevPoint
		if width > maxWidth {
			maxWidth = width
		}
	}

	return maxWidth
}
