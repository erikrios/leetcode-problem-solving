package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func findGCD(nums []int) int {
	smallest, biggest := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num < smallest {
			smallest = num
		}

		if num > biggest {
			biggest = num
		}
	}

	for i := smallest; i > 0; i-- {
		if smallest%i == 0 && biggest%i == 0 {
			return i
		}
	}

	return 1
}
