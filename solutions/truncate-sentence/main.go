package main

import "fmt"

func main() {
	fmt.Println(truncateSentence("Hello how are you Contestant", 4))
	fmt.Println(truncateSentence("What is the solution to this problem", 4))
	fmt.Println(truncateSentence("chopper is not a tanuki", 5))
}

func truncateSentence(s string, k int) string {
	var spaceCount int
	var i int

	for spaceCount < k && i < len(s) {
		if s[i] == ' ' {
			spaceCount++
		}

		i++
	}

	if i != len(s) {
		i--
	}

	return string(s[:i])
}
