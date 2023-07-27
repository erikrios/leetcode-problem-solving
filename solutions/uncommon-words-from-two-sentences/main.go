package main

import "fmt"

func main() {
	fmt.Println(uncommonFromSentences("this apple is sweet", "this apple is sour"))
	fmt.Println(uncommonFromSentences("apple apple", "banana"))
}

func uncommonFromSentences(s1, s2 string) []string {
	mapper := make(map[string]byte)

	wordCounter(s1, mapper)
	wordCounter(s2, mapper)

	res := make([]string, 0)

	for k, v := range mapper {
		if v == 1 {
			res = append(res, k)
		}
	}

	return res
}

func wordCounter(s string, mapper map[string]byte) {
	words := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		char := s[i]

		if char == ' ' {
			mapper[string(words)]++
			words = make([]byte, 0)
			continue
		}

		words = append(words, char)
	}

	mapper[string(words)]++
}
