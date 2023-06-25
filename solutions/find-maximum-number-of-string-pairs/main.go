package main

import "fmt"

func main() {
	fmt.Println(maximumNumberOfStringPairs([]string{"cd", "ac", "dc", "ca", "zz"}))
	fmt.Println(maximumNumberOfStringPairs([]string{"ab", "ba", "cc"}))
	fmt.Println(maximumNumberOfStringPairs([]string{"aa", "ab"}))
}

func maximumNumberOfStringPairs(words []string) int {
	pairsCounter := make(map[string]struct{})
	var counter int

	for i := 0; i < len(words); i++ {
		word := words[i]
		if word[0] > word[1] {
			word = string([]byte{word[1], word[0]})
		}

		if _, ok := pairsCounter[word]; ok {
			counter++
		}

		pairsCounter[word] = struct{}{}
	}

	return counter
}
