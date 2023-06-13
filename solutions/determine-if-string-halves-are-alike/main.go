package main

import "fmt"

func main() {
	fmt.Println(halvesAreAlike("book"))
	fmt.Println(halvesAreAlike("textbook"))
}

func halvesAreAlike(s string) bool {
	vowels := map[byte]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}

	var aCounter, bCounter int
	length := len(s)
	for i := 0; i < length/2; i++ {
		aVal := s[i]
		bVal := s[length/2+i]

		if _, ok := vowels[aVal]; ok {
			aCounter++
		}

		if _, ok := vowels[bVal]; ok {
			bCounter++
		}
	}

	return aCounter == bCounter
}
