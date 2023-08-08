package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(splitWordsBySeparator([]string{"one.two.three", "four.five", "six"}, '.'))
	fmt.Println(splitWordsBySeparator([]string{"$easy$", "$problem$"}, '$'))
	fmt.Println(splitWordsBySeparator([]string{"|||"}, '|'))
}

func splitWordsBySeparator(words []string, separator byte) []string {
	results := make([]string, 0)

	for i := 0; i < len(words); i++ {
		var builder strings.Builder

		for word, j := words[i], 0; j < len(word); j++ {

			if char := word[j]; char == separator {
				if builder.Len() == 0 {
					continue
				}

				results = append(results, builder.String())
				builder.Reset()
			} else {
				builder.WriteByte(char)
			}
		}

		if builder.Len() == 0 {
			continue
		}

		results = append(results, builder.String())
	}

	return results
}
