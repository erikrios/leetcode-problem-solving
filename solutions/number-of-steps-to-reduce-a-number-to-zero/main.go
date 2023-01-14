package main

import "fmt"

func main() {
	fmt.Println(numberOfSteps(14))
	fmt.Println(numberOfSteps(8))
}

func numberOfSteps(num int) int {
	var counter int

	for num != 0 {
		if num%2 == 0 {
			num /= 2
		} else {
			num -= 1
		}
		counter++
	}

	return counter
}
