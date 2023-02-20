package main

import (
	"fmt"
)

func main() {
	fmt.Println(separateDigits([]int{13, 25, 83, 77}))
	fmt.Println(separateDigits([]int{7, 1, 3, 9}))
}

func separateDigits(nums []int) []int {
	results := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		num := nums[i]

		temp := make([]int, 0)
		for num > 0 {
			temp = append(temp, num%10)
			num /= 10
		}

		for j := len(temp) - 1; j >= 0; j-- {
			results = append(results, temp[j])
		}
	}

	return results
}
