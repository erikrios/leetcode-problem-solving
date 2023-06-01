package main

import "fmt"

func main() {
	fmt.Println(toGoatLatin("I speak Goat Latin"))
	fmt.Println(toGoatLatin("The quick brown fox jumped over the lazy dog"))
	fmt.Println(toGoatLatin("HZ sg L"))
	fmt.Println(toGoatLatin("Ffk G KdiC M I"))
}

func toGoatLatin(sentence string) string {
	words := make([]byte, 0)

	isFirstChar := true
	var lastChar byte = ' '
	var additionalCounter int
	for i := 0; i < len(sentence); i++ {
		char := sentence[i]

		if isFirstChar && i == len(sentence)-1 {
			words = append(words, char, 'm', 'a')
			additionalCounter++

			for j := 0; j < additionalCounter; j++ {
				words = append(words, 'a')
			}

			break
		}

		if isFirstChar {
			var isCap bool

			if char >= 'A' && char <= 'Z' {
				isCap = true
				char = char - 'A' + 'a'
			}

			if char == 'a' || char == 'i' || char == 'u' || char == 'e' || char == 'o' {
				if isCap {
					char = char - 'a' + 'A'
				}

				words = append(words, char)
				lastChar = ' '
			} else {
				if isCap {
					char = char - 'a' + 'A'
				}

				lastChar = char
			}

			isFirstChar = false
		} else {
			if char == ' ' || i == len(sentence)-1 {
				if i == len(sentence)-1 {
					words = append(words, char)
				}

				isFirstChar = true
				if lastChar != ' ' {
					words = append(words, lastChar)
					lastChar = ' '
				}
				words = append(words, 'm', 'a')

				additionalCounter++

				for j := 0; j < additionalCounter; j++ {
					words = append(words, 'a')
				}

				if i != len(sentence)-1 {
					words = append(words, ' ')
				}
			} else {
				words = append(words, char)
			}
		}
	}

	return string(words)
}
