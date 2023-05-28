package main

import "fmt"

func main() {
	fmt.Println(minLength("ABFCACDB"))
	fmt.Println(minLength("ACBBD"))
	fmt.Println(minLength("CCCCDDDD"))
	fmt.Println(minLength("CCDDACDB"))
}

func minLength(s string) int {
	letters := []byte(s)

	for i := 0; i < len(letters); {
		letter := letters[i]

		if letter == 'A' && i < len(letters)-1 {
			nextLetter := letters[i+1]
			if nextLetter == 'B' {
				letters = trimNext(letters, i)
				continue
			}
		}

		if letter == 'C' && i < len(letters)-1 {
			nextLetter := letters[i+1]
			if nextLetter == 'D' {
				letters = trimNext(letters, i)
				continue
			}
		}

		if letter == 'B' && i-1 >= 0 {
			prevLetter := letters[i-1]
			if prevLetter == 'A' {
				letters = trimPrev(letters, i)
				i -= 2
				if i < 0 {
					i = 0
				}
				continue
			}
		}

		if letter == 'D' && i-1 >= 0 {
			prevLetter := letters[i-1]
			if prevLetter == 'C' {
				letters = trimPrev(letters, i)
				i -= 2
				if i < 0 {
					i = 0
				}
				continue
			}
		}

		i++
	}

	return len(letters)
}

func trimNext(letters []byte, i int) []byte {
	prevLetters := letters[:i]
	nextLetters := letters[i+2:]
	letters = append(prevLetters, nextLetters...)
	return letters
}

func trimPrev(letters []byte, i int) []byte {
	prevLetters := letters[:i-1]
	nextLetters := letters[i+1:]
	letters = append(prevLetters, nextLetters...)
	return letters
}
