package main

import "fmt"

func main() {
	fmt.Println(findMatrix([]int{1, 3, 4, 1, 2, 3, 1}))
	fmt.Println(findMatrix([]int{1, 2, 3, 4}))
}

func findMatrix(nums []int) [][]int {
	n := len(nums)
	mapper := make(map[int]int, n)

	for i := 0; i < n; i++ {
		num := nums[i]
		mapper[num]++
	}

	matrixs := make([][]int, 0)

	for len(mapper) > 0 {
		matrix := make([]int, 0)

		for k := range mapper {
			matrix = append(matrix, k)
			mapper[k]--
			if mapper[k] == 0 {
				delete(mapper, k)
			}
		}

		matrixs = append(matrixs, matrix)
	}

	return matrixs
}
