package main

import "fmt"

func main() {
	fmt.Println(splitNum(4325))
	fmt.Println(splitNum(687))
}

func splitNum(num int) int {
	var nums [10]int

	for num > 0 {
		digit := num % 10
		nums[digit]++
		num /= 10
	}

	var isForNum2 bool
	var num1, num2 int
	for i := 0; i < len(nums); i++ {
		count := nums[i]
		for j := 0; j < count; j++ {
			if isForNum2 {
				num2 *= 10
				num2 += i
			} else {
				num1 *= 10
				num1 += i
			}

			isForNum2 = !isForNum2
		}
	}

	return num1 + num2
}
