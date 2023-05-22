package main

import "fmt"

func main() {
	fmt.Println(findWords([]string{"Hello", "Alaska", "Dad", "Peace"}))
	fmt.Println(findWords([]string{"omk"}))
	fmt.Println(findWords([]string{"adsdf", "sfd"}))
}

func findWords(words []string) []string {
	americanKeyboards := map[byte]bool{
		'q': false, 'w': false, 'e': false, 'r': false, 't': false, 'y': false, 'u': false, 'i': false, 'o': false, 'p': false,
		'a': true, 's': true, 'd': true, 'f': true, 'g': true, 'h': true, 'j': true, 'k': true, 'l': true,
	}

	n := len(words)
	results := make([]string, 0, n)
	for i := 0; i < n; i++ {
		word := words[i]

		var isEmpty, isSecRow bool
		isValid := true
		firstChar := word[0]
		if isCap(firstChar) {
			firstChar = firstChar - 'A' + 'a'
		}

		if v, ok := americanKeyboards[firstChar]; ok {
			isSecRow = v
		} else {
			isEmpty = true
		}

		for j := 1; j < len(word); j++ {
			char := word[j]
			if isCap(char) {
				char = char - 'A' + 'a'
			}

			v, ok := americanKeyboards[char]
			if isEmpty {
				if ok {
					isValid = false
					break
				}
			} else if isSecRow {
				if !ok || !v {
					isValid = false
					break
				}
			} else {
				if !ok || v {
					isValid = false
					break
				}
			}
		}

		if isValid {
			results = append(results, word)
		}
	}

	return results
}

func isCap(char byte) bool {
	return !(char >= 'a' && char <= 'z')
}
