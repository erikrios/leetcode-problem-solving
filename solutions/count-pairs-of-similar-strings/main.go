package main

import "fmt"

func main() {
	fmt.Println(similarPairs([]string{"aba", "aabb", "abcd", "bac", "aabc"}))
	fmt.Println(similarPairs([]string{"aabb", "ab", "ba"}))
	fmt.Println(similarPairs([]string{"nba", "cba", "dba"}))
}

func similarPairs(words []string) int {
	n := len(words)
	uniqueWords := make([]map[byte]struct{}, 0, n)

	for i := 0; i < n; i++ {
		word := words[i]
		m := len(word)
		uniqueWord := make(map[byte]struct{}, m)

		for j := 0; j < m; j++ {
			uniqueWord[word[j]] = struct{}{}
		}

		uniqueWords = append(uniqueWords, uniqueWord)
	}

	var counter int
	for i := 0; i < n-1; i++ {
		iWord := uniqueWords[i]

		for j := i + 1; j < n; j++ {

			jWord := uniqueWords[j]
			if len(iWord) != len(jWord) {
				continue
			}

			if isIdentical(iWord, jWord) {
				counter++
			}
		}
	}

	return counter
}

func isIdentical(a, b map[byte]struct{}) bool {
	for k := range b {
		if _, ok := a[k]; !ok {
			return false
		}
	}

	return true
}
