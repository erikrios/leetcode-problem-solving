package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
}

func isPalindrome(x int) bool {
	xStr := strconv.Itoa(x)

	for i := 0; i < len(xStr)/2; i++ {
		if xStr[i] != xStr[len(xStr)-1-i] {
			return false
		}
	}
	return true
}
