package main

import "fmt"

func main() {
	fmt.Println(makeEqual([]string{"abc", "aabc", "bc"}))
	fmt.Println(makeEqual([]string{"ab", "a"}))
}

func makeEqual(words []string) bool {
	n := len(words)

	charMapper := make(map[byte]int)
	for i := 0; i < n; i++ {
		word := words[i]
		for j := 0; j < len(word); j++ {
			char := word[j]
			charMapper[char]++
		}
	}

	for _, v := range charMapper {
		if v%n != 0 {
			return false
		}
	}

	return true
}
