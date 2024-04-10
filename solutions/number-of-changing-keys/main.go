package main

import "fmt"

func main() {
	fmt.Println(countKeyChanges("aAbBcC"))
	fmt.Println(countKeyChanges("AaAaAaaA"))
}

func countKeyChanges(s string) int {
	var counter int
	var prev byte

	for i := 0; i < len(s); i++ {
		char := s[i]
		if char >= 'A' && char <= 'Z' {
			char = char - 'A' + 'a'
		}

		if i > 0 && prev != char {
			counter++
		}

		prev = char
	}

	return counter
}
