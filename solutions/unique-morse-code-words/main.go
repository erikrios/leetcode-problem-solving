package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"}))
	fmt.Println(uniqueMorseRepresentations([]string{"a"}))
}

func uniqueMorseRepresentations(words []string) int {
	morseCodes := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

	morsePairs := make(map[byte]string)

	var starter byte = 'a'
	for i := 0; i < len(morseCodes); i++ {
		morsePairs[starter] = morseCodes[i]
		starter++
	}

	results := make(map[string]bool)

	for i := 0; i < len(words); i++ {
		word := words[i]
		var builder strings.Builder

		for j := 0; j < len(word); j++ {
			builder.WriteString(morsePairs[word[j]])
		}

		results[builder.String()] = true
	}

	return len(results)
}
