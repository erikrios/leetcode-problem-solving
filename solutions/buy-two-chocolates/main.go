package main

import "fmt"

func main() {
	fmt.Println(buyChoco([]int{1, 2, 2}, 3))
	fmt.Println(buyChoco([]int{3, 2, 3}, 3))
}

func buyChoco(prices []int, money int) int {
	cheaper1, cheaper2 := 100, 100

	for i := 0; i < len(prices); i++ {
		price := prices[i]

		if price < cheaper1 {
			cheaper2 = cheaper1
			cheaper1 = price
			continue
		}

		if price < cheaper2 {
			cheaper2 = price
		}
	}

	if totalPrice := cheaper1 + cheaper2; totalPrice <= money {
		return money - totalPrice
	}

	return money
}
