package main

import "fmt"

func main() {
	fmt.Println(largestGoodInteger("6777133339"))
	fmt.Println(largestGoodInteger("2300019"))
	fmt.Println(largestGoodInteger("42352338"))
}

func largestGoodInteger(num string) string {
	n := len(num)
	var largest byte

	for i := 0; i < n-2; i++ {
		cur := num[i]
		if cur <= largest {
			continue
		}

		nextOne := num[i+1]
		nextTwo := num[i+2]

		if cur == nextOne && nextOne == nextTwo {
			largest = cur
		}
	}

	result := make([]byte, 0, 3)

	if largest != 0 {
		for i := 0; i < 3; i++ {
			result = append(result, largest)
		}
	}

	return string(result)
}
