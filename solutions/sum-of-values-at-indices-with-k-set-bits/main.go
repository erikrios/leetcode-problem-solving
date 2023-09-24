package main

import "fmt"

func main() {
	fmt.Println(sumIndicesWithKSetBits([]int{5, 10, 1, 5, 2}, 1))
	fmt.Println(sumIndicesWithKSetBits([]int{4, 3, 2, 1}, 2))
}

func sumIndicesWithKSetBits(nums []int, k int) int {
	var sum int

	for i := 0; i < len(nums); i++ {
		if sumSetBits(i) == k {
			sum += nums[i]
		}
	}

	return sum
}

func sumSetBits(num int) int {
	var sum int

	for num > 0 {
		digit := num % 2
		if digit == 1 {
			sum++
		}

		num /= 2
	}

	return sum
}
