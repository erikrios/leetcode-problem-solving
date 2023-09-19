package main

import "fmt"

func main() {
	fmt.Println(totalMoney(4))
	fmt.Println(totalMoney(10))
	fmt.Println(totalMoney(20))
}

func totalMoney(n int) int {
	var sum, inc int

	for i := 1; i <= n; i++ {
		inc++
		sum += inc

		if i%7 == 0 {
			inc = i / 7
		}
	}

	return sum
}
