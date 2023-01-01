package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(mostWordsFound([]string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}))
	fmt.Println(mostWordsFound([]string{"please wait", "continue to fight", "continue to win"}))
}

func mostWordsFound(sentences []string) int {
	var count int

	for _, sentence := range sentences {
		words := strings.Split(sentence, " ")
		length := len(words)

		if length > count {
			count = length
		}
	}

	return count
}
