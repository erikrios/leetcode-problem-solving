package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome("  "))
}

func isPalindrome(s string) bool {
	s = removeAlphanum(strings.ToLower(s))
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}

func removeAlphanum(s string) string {
	result := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		if (s[i] >= 'a' && s[i] <= 'z') ||
			(s[i] >= '0' && s[i] <= '9') {
			result = append(result, s[i])
		}
	}

	return string(result)
}
