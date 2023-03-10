package main

import "fmt"

func main() {
	fmt.Println(findNumbers([]int{12, 345, 2, 6, 7896}))
	fmt.Println(findNumbers([]int{555, 901, 482, 1771}))
}

func findNumbers(nums []int) int {
	var counter int

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		var totalDigit int

		for num > 0 {
			num /= 10
			totalDigit++
		}

		if totalDigit%2 == 0 {
			counter++
		}
	}

	return counter
}
