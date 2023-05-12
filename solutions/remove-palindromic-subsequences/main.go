package main

import "fmt"

func main() {
	fmt.Println(removePalindromeSub("ababa"))
	fmt.Println(removePalindromeSub("abb"))
	fmt.Println(removePalindromeSub("baabb"))
}

func removePalindromeSub(s string) int {
	if len(s) == 0 {
		return 0
	} else if isPalindrome(s) {
		return 1
	} else {
		return 2
	}
}

func isPalindrome(s string) bool {
	isPalindrome := true

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			isPalindrome = false
			break
		}
	}

	return isPalindrome
}
