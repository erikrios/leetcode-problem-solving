package main

import "fmt"

func main() {
	fmt.Println(distributeCandies(7, 4))
	fmt.Println(distributeCandies(10, 3))
}

func distributeCandies(candies int, numPeople int) []int {
	res := make([]int, numPeople, numPeople)

	for i, j := 1, 0; candies > 0; i, j = i+1, j+1 {
		if j == numPeople {
			j = 0
		}

		if candies-i < 0 {
			res[j] += candies
			break
		}

		candies -= i
		res[j] += i
	}

	return res
}
