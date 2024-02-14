package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(sortVowels("lEetcOde"))
	fmt.Println(sortVowels("lYmpH"))
}

func sortVowels(s string) string {
	sortedVowels := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		char := s[i]
		if isVowel(char) {
			sortedVowels = append(sortedVowels, char)
		}
	}

	sort.Slice(sortedVowels, func(i, j int) bool {
		return sortedVowels[i] < sortedVowels[j]
	})

	res := make([]byte, 0, len(s))
	var j int
	for i := 0; i < len(s); i++ {
		char := s[i]
		if isVowel(char) {
			res = append(res, sortedVowels[j])
			j++
		} else {
			res = append(res, char)
		}
	}

	return string(res)
}

func isVowel(char byte) bool {
	switch char {
	case 'A', 'I', 'U', 'E', 'O', 'a', 'i', 'u', 'e', 'o':
		return true
	default:
		return false
	}
}
