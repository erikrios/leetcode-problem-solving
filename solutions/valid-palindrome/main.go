package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isPalindromeTwoPointers("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindromeTwoPointers("race a car"))
	fmt.Println(isPalindromeTwoPointers("  "))
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

// one pass solution
func isPalindromeTwoPointers(s string) bool {
	for i, j := 0, len(s)-1; i <= j; {
		firstChar, secondChar := s[i], s[j]

		if !isAlphanum(firstChar) {
			i++
		} else if !isAlphanum(secondChar) {
			j--
		} else {
			firstChar, secondChar = toLower(firstChar), toLower(secondChar)
			if firstChar != secondChar {
				return false
			}
			i++
			j--
		}
	}

	return true
}

func isAlphanum(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')
}

func toLower(char byte) byte {
	if char >= 'A' && char <= 'Z' {
		return 'a' + char - 'A'
	}

	return char
}
