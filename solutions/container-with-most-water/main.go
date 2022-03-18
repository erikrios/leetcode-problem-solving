package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 1}))
	fmt.Println(maxArea([]int{1, 5, 1}))
}

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1

	var minHeightIndex int
	if height[left] < height[right] {
		minHeightIndex = left
	} else {
		minHeightIndex = right
	}

	max := height[minHeightIndex] * (right - left)

	for (height[left] < height[left+1] || height[right] < height[right-1]) && left+1 < right-1 {
		if height[left] < height[left+1] {
			left++
		}

		if height[right] < height[right-1] {
			right--
		}

		if height[left] < height[right] {
			minHeightIndex = left
		} else {
			minHeightIndex = right
		}

		area := height[minHeightIndex] * (right - left)

		if area > max {
			max = area
		}
	}

	return max
}
