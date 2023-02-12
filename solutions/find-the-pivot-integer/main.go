package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(pivotInteger(8))
	fmt.Println(pivotInteger(1))
	fmt.Println(pivotInteger(4))
	fmt.Println(pivotInteger(100000))
}

func pivotIntegers(n int) int {
	for i := n; i > 0; i-- {
		rSum := sum(i, n)
		lSum := sum(1, i)

		if rSum == lSum {
			return i
		} else if rSum > lSum {
			return -1
		}
	}

	return -1
}

func sum(start, end int) int {
	var sum int

	for i := start; i <= end; i++ {
		sum += i
	}

	return sum
}

func pivotInteger(n int) int {
	sum := n * (n + 1) / 2
	x := int(math.Sqrt(float64(sum)))

	if x*x == sum {
		return x
	}

	return -1
}
