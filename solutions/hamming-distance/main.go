package main

import "fmt"

func main() {
	fmt.Println(hammingDistance(1, 4))
	fmt.Println(hammingDistance(3, 1))
}

func hammingDistance(x int, y int) int {
	var counter int

	for x > 0 || y > 0 {
		xV := x % 2
		x /= 2

		yV := y % 2
		y /= 2

		if xV != yV {
			counter++
		}
	}

	return counter
}
