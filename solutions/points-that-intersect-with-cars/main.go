package main

import "fmt"

func main() {
	fmt.Println(numberOfPoints([][]int{{3, 6}, {1, 5}, {4, 7}}))
	fmt.Println(numberOfPoints([][]int{{1, 3}, {5, 8}}))
}

func numberOfPoints(nums [][]int) int {
	coveredPoints := make(map[int]struct{})

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		start, end := num[0], num[1]

		for j := start; j <= end; j++ {
			coveredPoints[j] = struct{}{}
		}
	}

	return len(coveredPoints)
}
