package main

import "fmt"

func main() {
	fmt.Println(countWords([]string{"leetcode", "is", "amazing", "as", "is"}, []string{"amazing", "leetcode", "is"}))
	fmt.Println(countWords([]string{"b", "bb", "bbb"}, []string{"a", "aa", "aaa"}))
	fmt.Println(countWords([]string{"a", "ab"}, []string{"a", "a", "a", "ab"}))
}

func countWords(words1 []string, words2 []string) int {
	n := len(words1)
	uniqueWords1 := make(map[string]rune, n)

	m := len(words2)
	uniqueWords2 := make(map[string]rune, m)

	var l int
	if m < n {
		l = n
	} else {
		l = m
	}

	for i := 0; i < l; i++ {
		if i < n {
			word := words1[i]
			uniqueWords1[word]++
		}

		if i < m {
			word := words2[i]
			uniqueWords2[word]++
		}
	}

	var counter int
	if len(uniqueWords1) < len(uniqueWords2) {
		for k, v := range uniqueWords1 {
			if v > 1 {
				continue
			}

			if v2, ok := uniqueWords2[k]; !ok || v2 > 1 {
				continue
			}

			counter++
		}
	} else {
		for k, v := range uniqueWords2 {
			if v > 1 {
				continue
			}

			if v1, ok := uniqueWords1[k]; !ok || v1 > 1 {
				continue
			}

			counter++
		}
	}

	return counter
}
