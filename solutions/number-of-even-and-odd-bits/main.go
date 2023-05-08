package main

import "fmt"

func main() {
	fmt.Println(evenOddBit(17))
	fmt.Println(evenOddBit(2))
}

func evenOddBit(n int) []int {
	res := make([]int, 2, 2)

	for i := 0; n > 0; i++ {
		v := n % 2
		n /= 2

		if v != 1 {
			continue
		}

		if i%2 == 0 {
			res[0]++
		} else {
			res[1]++
		}
	}

	return res
}
