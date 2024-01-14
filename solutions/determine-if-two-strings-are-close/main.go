package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(closeStrings("abc", "bca"))
	fmt.Println(closeStrings("a", "aa"))
	fmt.Println(closeStrings("cabbba", "abbccc"))
}

func closeStrings(word1, word2 string) bool {
	n := len(word1)
	m := len(word2)
	if n != m {
		return false
	}

	word1CharCounters := make([]int, 26)
	word2CharCounters := make([]int, 26)

	for i := 0; i < n; i++ {
		char1 := word1[i]
		char2 := word2[i]

		word1CharCounters[char1-'a']++
		word2CharCounters[char2-'a']++
	}

	word1Counters, word2Counters := make([]int, 0), make([]int, 0)

	for i := 0; i < 26; i++ {
		aCounter := word1CharCounters[i]
		bCounter := word2CharCounters[i]
		if aCounter == 0 && bCounter != 0 {
			return false
		}

		if aCounter != 0 && bCounter == 0 {
			return false
		}

		if aCounter == 0 && bCounter == 0 {
			continue
		}

		word1Counters = append(word1Counters, aCounter)
		word2Counters = append(word2Counters, bCounter)
	}

	sort.Ints(word1Counters)
	sort.Ints(word2Counters)

	for i := 0; i < len(word1Counters); i++ {
		if word1Counters[i] != word2Counters[i] {
			return false
		}
	}

	return true
}
