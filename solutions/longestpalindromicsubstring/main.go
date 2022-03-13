package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
}

func longestPalindrome(s string) string {
	for i := len(s); i >= 1; i-- {
		for j := 0; j+i <= len(s); j++ {
			sub := string(s[j : i+j])
			if isPalindrome(sub) {
				return sub
			}
		}
	}

	return ""
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}
