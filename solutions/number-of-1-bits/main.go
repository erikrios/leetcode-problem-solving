package main

import "fmt"

func main() {
	fmt.Println(hammingWeight(11))
	fmt.Println(hammingWeight(128))
	fmt.Println(hammingWeight(4294967293))
}

func hammingWeight(num uint32) int {
	var counter int

	for num > 0 {
		digit := num % 2
		num /= 2
		if digit == 1 {
			counter++
		}
	}

	return counter
}
