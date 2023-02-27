package main

import "fmt"

func main() {
	fmt.Println(firstPalindrome([]string{"abc", "car", "ada", "racecar", "cool"}))
	fmt.Println(firstPalindrome([]string{"notapalindrome", "racecar"}))
	fmt.Println(firstPalindrome([]string{"def", "ghi"}))
}

func firstPalindrome(words []string) string {
	var result string

	for i := 0; i < len(words); i++ {
		word := words[i]

		isPalindrome := true
		for j := 0; j < len(word)/2; j++ {
			if word[j] != word[len(word)-1-j] {
				isPalindrome = false
				break
			}
		}

		if isPalindrome {
			result = word
			break
		}
	}

	return result
}
