package main

import "fmt"

func main() {
	fmt.Println(findMaxLength([]int{0, 1}))
	fmt.Println(findMaxLength([]int{0, 1, 0}))
	fmt.Println(findMaxLength([]int{0, 0, 0, 0, 1, 1}))
	fmt.Println(findMaxLength([]int{0, 0, 1, 0, 0, 0, 1, 1}))
	fmt.Println(findMaxLength([]int{0, 1, 1, 0, 1, 1, 1, 0}))
}

func findMaxLength(nums []int) int {
	var count, maxLength int
	mapper := map[int]int{0: 0}

	for i, num := range nums {
		i++
		switch num {
		case 0:
			count--
		default:
			count++
		}

		if idx, ok := mapper[count]; ok {
			maxLength = max(maxLength, i-idx)
		} else {
			mapper[count] = i
		}
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
