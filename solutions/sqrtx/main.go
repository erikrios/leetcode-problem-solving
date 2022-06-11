package main

import "fmt"

func main() {
	fmt.Println(mySqrt(8))
}

func mySqrt(x int) int {
	for i := 0; true; i++ {
		if i*i == x {
			return i
		} else if i*i > x {
			return i - 1
		}
	}

	return -1
}
