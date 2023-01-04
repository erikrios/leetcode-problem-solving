package main

import "fmt"

func main() {
	fmt.Println(kidsWithCandies([]int{2, 3, 5, 1, 3}, 3))
	fmt.Println(kidsWithCandies([]int{4, 2, 1, 1, 2}, 1))
	fmt.Println(kidsWithCandies([]int{12, 1, 12}, 10))
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	results := make([]bool, len(candies))

	var greatest int

	for _, candy := range candies {
		if candy > greatest {
			greatest = candy
		}
	}

	for i, candy := range candies {
		results[i] = candy+extraCandies >= greatest
	}

	return results
}
