package main

import "fmt"

func main() {
	fmt.Println(countPrefixes([]string{"a", "b", "c", "ab", "bc", "abc"}, "abc"))
	fmt.Println(countPrefixes([]string{"a", "a"}, "aa"))
}

func countPrefixes(words []string, s string) int {
	var counter int

	for i := 0; i < len(words); i++ {
		word := words[i]

		if len(word) > len(s) {
			continue
		}

		isPrefix := true
		for j := 0; j < len(word); j++ {
			if word[j] != s[j] {
				isPrefix = false
				break
			}
		}

		if isPrefix {
			counter++
		}
	}

	return counter
}
