package main

import "fmt"

func longestPalindrome(s string) string {
	result := string(s[0])

	for i := 0; i < len(s); i++ {
		for j := len(s); j > i; j-- {
			word := s[i:j]
			if word == reverse(word) {
				if len(result) < len(word) {
					result = word
					break
				}
			}
		}
	}

	return result
}

func reverse(s string) string {
	reversedString := make([]byte, 0)

	for i := len(s) - 1; i >= 0; i-- {
		reversedString = append(reversedString, s[i])
	}

	return string(reversedString)
}

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome("bb"))
	fmt.Println(longestPalindrome("aacabdkacaa"))
}
