package main

import "fmt"

func main() {
	fmt.Println(countBits(2))
	fmt.Println(countBits(5))
}

func countBits(n int) []int {
	res := make([]int, n+1, n+1)

	for i := 0; i <= n; i++ {
		num := i
		var counter int

		for num > 0 {
			bit := num % 2
			num /= 2
			if bit == 1 {
				counter++
			}
		}

		res[i] = counter
	}

	return res
}
