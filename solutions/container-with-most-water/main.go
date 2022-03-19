package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 1}))
	fmt.Println(maxArea([]int{4, 5, 1}))
	fmt.Println(maxArea([]int{1, 2, 4, 3}))
	fmt.Println(maxArea([]int{2, 3, 4, 5, 18, 17, 6}))
}

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	var max int

	for i < j {
		var minIndex int
		if height[i] < height[j] {
			minIndex = i
		} else {
			minIndex = j
		}

		area := height[minIndex] * (j - i)

		if max < area {
			max = area
		}

		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return max
}
