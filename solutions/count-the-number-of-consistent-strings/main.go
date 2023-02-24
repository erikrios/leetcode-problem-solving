package main

import "fmt"

func main() {
	fmt.Println(countConsistentStrings("ab", []string{"ad", "bd", "aaab", "baa", "badab"}))
	fmt.Println(countConsistentStrings("abc", []string{"a", "b", "c", "ab", "ac", "bc", "abc"}))
	fmt.Println(countConsistentStrings("cad", []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"}))
}

func countConsistentStrings(allowed string, words []string) int {
	allowedMapper := make(map[byte]bool)
	for i := 0; i < len(allowed); i++ {
		allowedMapper[allowed[i]] = true
	}

	var counter int
	for i := 0; i < len(words); i++ {
		word := words[i]

		var isNotAllowed bool
		for j := 0; j < len(word); j++ {
			if _, ok := allowedMapper[word[j]]; !ok {
				isNotAllowed = true
				break
			}
		}

		if !isNotAllowed {
			counter++
		}
	}

	return counter
}
