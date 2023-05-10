package main

import "fmt"

func main() {
	fmt.Println(countGoodSubstrings("xyzzaz"))
	fmt.Println(countGoodSubstrings("aababcabc"))
}

func countGoodSubstrings(s string) int {
	var counter int

	for i := 0; i < len(s)-2; i += 1 {
		firstChar := s[i]
		secondChar := s[i+1]
		thirdChar := s[i+2]

		if firstChar != secondChar && firstChar != thirdChar && secondChar != thirdChar {
			counter++
		}
	}

	return counter
}
