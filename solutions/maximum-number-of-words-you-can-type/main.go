package main

import "fmt"

func main() {
	fmt.Println(canBeTypedWords("hello world", "ad"))
	fmt.Println(canBeTypedWords("leet code", "lt"))
	fmt.Println(canBeTypedWords("leet code", "e"))
}

func canBeTypedWords(text, brokenLetters string) int {
	n := len(brokenLetters)
	brokenMapper := make(map[byte]int, n)

	for i := 0; i < n; i++ {
		letter := brokenLetters[i]
		brokenMapper[letter]++
	}

	n = len(text)
	wc := 1
	var shouldSkipped bool
	for i := 0; i < n; i++ {
		char := text[i]
		if char == ' ' {
			wc++
			shouldSkipped = false
		}

		if shouldSkipped {
			continue
		}
		if _, ok := brokenMapper[char]; ok {
			wc--
			shouldSkipped = true
		}
	}

	return wc
}
