package main

import "fmt"

func main() {
	fmt.Println(arrayStringsAreEqual([]string{"ab", "c"}, []string{"a", "bc"}))
	fmt.Println(arrayStringsAreEqual([]string{"a", "cb"}, []string{"ab", "c"}))
	fmt.Println(arrayStringsAreEqual([]string{"abc", "d", "defg"}, []string{"abcddefg"}))
}

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	var wordOne, wordTwo []byte

	for i := 0; i < len(word1); i++ {
		words := word1[i]
		for j := 0; j < len(words); j++ {
			wordOne = append(wordOne, words[j])
		}
	}

	for i := 0; i < len(word2); i++ {
		words := word2[i]
		for j := 0; j < len(words); j++ {
			wordTwo = append(wordTwo, words[j])
		}
	}

	return string(wordOne) == string(wordTwo)
}
