package main

import "fmt"

func main() {
	fmt.Println(sumOfEncryptedInt([]int{1, 2, 3}))
	fmt.Println(sumOfEncryptedInt([]int{10, 21, 31}))
}

func sumOfEncryptedInt(nums []int) int {
	n := len(nums)
	var sum int

	for i := 0; i < n; i++ {
		num := nums[i]

		var max, counter int

		for num > 0 {
			digit := num % 10
			num /= 10
			counter++
			if digit > max {
				max = digit
			}
		}

		var encNum int
		multiplier := 1
		for i := 0; i < counter; i++ {
			encNum += (max * multiplier)
			multiplier *= 10
		}

		sum += encNum
	}

	return sum
}
