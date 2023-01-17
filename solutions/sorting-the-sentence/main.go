package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(sortSentence("is2 sentence4 This1 a3"))
	fmt.Println(sortSentence("Myself2 Me1 I4 and3"))
}

func sortSentence(s string) string {
	words := strings.Split(s, " ")

	results := make([]string, len(words))

	for i := 0; i < len(words); i++ {
		word := words[i]
		indexByte := word[len(word)-1]
		indexStr := string(indexByte)
		index, _ := strconv.Atoi(indexStr)

		results[index-1] = string(word[:len(word)-1])
	}

	var builder strings.Builder

	for i, result := range results {
		builder.WriteString(result)

		if i != len(results)-1 {
			builder.WriteString(" ")
		}
	}

	return builder.String()
}
