package main

import "fmt"

func main() {
	fmt.Println(maxDistance([]int{1, 1, 1, 6, 1, 1, 1}))
	fmt.Println(maxDistance([]int{1, 8, 3, 8, 3}))
	fmt.Println(maxDistance([]int{0, 1}))
}

func maxDistance(colors []int) int {
	var max int

	for i := 0; i < len(colors)-1; i++ {
		if colors[len(colors)-1] != colors[i] {
			max = len(colors) - 1 - i
			break
		}
	}

	for i := len(colors) - 1; i > 0; i-- {
		length := i - 0
		if length <= max {
			break
		}
		if colors[0] != colors[i] {
			max = length
			break
		}
	}

	return max
}
