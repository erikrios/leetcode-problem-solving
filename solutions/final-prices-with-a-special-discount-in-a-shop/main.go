package main

import "fmt"

func main() {
	fmt.Println(finalPrices([]int{8, 4, 6, 2, 3}))
	fmt.Println(finalPrices([]int{1, 2, 3, 4, 5}))
	fmt.Println(finalPrices([]int{10, 1, 1, 6}))
}

func finalPrices(prices []int) []int {
	for i := 0; i < len(prices)-1; i++ {
		price := prices[i]

		for j := i + 1; j < len(prices); j++ {
			nextPrice := prices[j]

			if nextPrice <= price {
				prices[i] -= nextPrice
				break
			}
		}
	}

	return prices
}
