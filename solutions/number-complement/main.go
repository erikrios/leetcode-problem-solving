package main

import "fmt"

func main() {
	fmt.Println(findComplement(5))
	fmt.Println(findComplement(1))
	fmt.Println(findComplement(8))
}

func findComplement(num int) int {
	var compl int
	pow := 1

	for num > 0 {
		digit := num % 2
		num /= 2
		if digit == 0 {
			compl += pow
		}

		pow *= 2
	}

	return compl
}
