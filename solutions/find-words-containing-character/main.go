package main

import "fmt"

func main() {
	fmt.Println(findWordsContaining([]string{"leet", "code"}, 'e'))
	fmt.Println(findWordsContaining([]string{"abc", "bcd", "aaaa", "cbc"}, 'a'))
	fmt.Println(findWordsContaining([]string{"abc", "bcd", "aaaa", "cbc"}, 'z'))
}

func findWordsContaining(words []string, x byte) []int {
	results := make([]int, 0)

	for i := 0; i < len(words); i++ {
		word := words[i]
		for j := 0; j < len(word); j++ {
			if word[j] == x {
				results = append(results, i)
				break
			}
		}
	}

	return results
}
