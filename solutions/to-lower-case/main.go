package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(toLowerCase("Hello"))
	fmt.Println(toLowerCase("here"))
	fmt.Println(toLowerCase("LOVELY"))
}

func toLowerCases(s string) string {
	return strings.ToLower(s)
}

func toLowerCase(s string) string {
	cap := len(s)
	lower := make([]byte, cap, cap)

	for i := 0; i < len(s); i++ {
		char := s[i]

		if char >= 'A' && char <= 'Z' {
			charConv := char - 'A' + 'a'
			lower[i] = charConv
		} else {
			lower[i] = char
		}
	}

	return string(lower)
}
