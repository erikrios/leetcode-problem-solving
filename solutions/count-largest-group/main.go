package main

import "fmt"

func main() {
	fmt.Println(countLargestGroup(13))
	fmt.Println(countLargestGroup(2))
	fmt.Println(countLargestGroup(3853))
}

func countLargestGroup(n int) int {
	mapper := make(map[int]uint16)
	var max uint16

	for i := 1; i <= n; i++ {
		sum := sumDigit(i)
		mapper[sum]++
		count := mapper[sum]

		if count > max {
			max = count
		}
	}

	var counter int
	for _, v := range mapper {
		if v == max {
			counter++
		}
	}

	return counter
}

func sumDigit(num int) int {
	if num <= 9 {
		return num
	}

	var sum int
	for num > 0 {
		digit := num % 10
		sum += digit
		num /= 10
	}

	return sum
}
