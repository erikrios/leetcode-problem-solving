package main

import (
	"fmt"
)

func main() {
	fmt.Println(isHappy(19))
	fmt.Println(isHappy(2))
}

func isHappy(n int) bool {
	return isHappyNum(n, map[int]struct{}{n: {}})
}

func isHappyNum(n int, mapper map[int]struct{}) bool {
	var sum int
	var total int

	for n > 0 {
		digit := n % 10
		n /= 10
		sum += digit * digit
		total += digit
	}

	if total == 1 {
		return true
	}

	if _, ok := mapper[sum]; ok {
		return false
	} else {
		mapper[sum] = struct{}{}
	}

	return isHappyNum(sum, mapper)
}
