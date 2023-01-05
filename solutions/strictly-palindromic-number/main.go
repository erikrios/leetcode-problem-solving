package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isStrictlyPalindromic(9))
	fmt.Println(isStrictlyPalindromic(4))
}

func isStrictlyPalindromic(n int) bool {
	for i := 2; i <= n-2; i++ {
		base2Form := strconv.FormatInt(int64(i), 2)
		if !isPalindromic(base2Form) {
			return false
		}
	}

	return true
}

func isPalindromic(str string) bool {
	for i := 0; i < len(str); i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}

	return true
}
