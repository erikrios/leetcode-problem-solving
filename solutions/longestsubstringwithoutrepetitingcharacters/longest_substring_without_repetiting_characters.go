package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	if len(s) < 1 {
		return 0
	}

	if len(s) < 2 {
		return 1
	}

	var longest int
	characters := make([]byte, 0)

out:
	for i := 0; i < len(s)-1; i++ {
		characters = append(characters, s[i])

		for j := i + 1; j < len(s); j++ {
			for _, char := range characters {
				if char == s[j] {
					if longest < len(characters) {
						longest = len(characters)
					}
					characters = characters[:0]
					continue out
				}
			}
			characters = append(characters, s[j])
		}
		if longest < len(characters) {
			longest = len(characters)
		}
		characters = characters[:0]
	}

	return longest
}

func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
