package main

import "fmt"

func main() {
	fmt.Println(heightChecker([]int{1, 1, 4, 2, 1, 3}))
	fmt.Println(heightChecker([]int{5, 1, 2, 3, 4}))
	fmt.Println(heightChecker([]int{1, 2, 3, 4, 5}))
}

func heightChecker(heights []int) int {
	var counter int

	originals := make([]int, len(heights), len(heights))
	copy(originals, heights)

	for i := 0; i < len(heights); i++ {
		for j := i + 1; j < len(heights); j++ {
			if heights[i] > heights[j] {
				heights[i], heights[j] = heights[j], heights[i]
			}
		}

		beforeSwap := originals[i]
		afterSwap := heights[i]

		if beforeSwap != afterSwap {
			counter++
		}
	}

	return counter
}
