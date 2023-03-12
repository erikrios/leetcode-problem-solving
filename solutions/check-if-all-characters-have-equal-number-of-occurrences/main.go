package main

import "fmt"

func main() {
	fmt.Println(areOccurencesEqual("abacbc"))
	fmt.Println(areOccurencesEqual("aaabb"))
}

func areOccurencesEqual(s string) bool {
	mapper := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		char := s[i]
		mapper[char]++
	}

	total := mapper[s[0]]

	for _, v := range mapper {
		if v != total {
			return false
		}
	}

	return true
}
