package main

import "fmt"

func main() {
	fmt.Println(mergeAlternately("abc", "pqr"))
	fmt.Println(mergeAlternately("ab", "pqrs"))
	fmt.Println(mergeAlternately("abcd", "pq"))
}

func mergeAlternately(word1 string, word2 string) string {
	totalLen := len(word1) + len(word2)
	results := make([]byte, totalLen, totalLen)

	var resPtr int
	var i int
	for resPtr < totalLen {
		if i < len(word1) {
			results[resPtr] = word1[i]
			resPtr++
		}

		if i < len(word2) {
			results[resPtr] = word2[i]
			resPtr++
		}

		i++
	}

	return string(results)
}
