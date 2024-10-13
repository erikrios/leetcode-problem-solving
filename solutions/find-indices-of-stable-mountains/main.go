package main

import "fmt"

func main() {
	fmt.Println(stableMountains([]int{1, 2, 3, 4, 5}, 2))
	fmt.Println(stableMountains([]int{10, 1, 10, 1, 10}, 3))
	fmt.Println(stableMountains([]int{10, 1, 10, 1, 10}, 10))
}

func stableMountains(height []int, threshold int) []int {
	n := len(height)
	results := make([]int, 0)

	for i := 1; i < n; i++ {
		if height[i-1] > threshold {
			results = append(results, i)
		}
	}

	return results
}
