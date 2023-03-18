
package main

import "fmt"

func main() {
	fmt.Println(vowelStrings([]string{"are","amy","u"}, 0, 2))
	fmt.Println(vowelStrings([]string{"hey","aeo","mu","ooo","artro"},1,4))
}

func vowelStrings(words []string, left int, right int) int {
	var counter int

	for i := left; i <= right; i++ {
		word := words[i]
		if isVowel(word) {
			counter++
		}
	}

	return counter
}

func isVowel(s string) bool {
	firstChar := s[0]
	lastChar := s[len(s)-1]

	return (firstChar == 'a' || firstChar == 'e' || firstChar == 'i' || firstChar == 'o' || firstChar == 'u') &&  (lastChar == 'a' || lastChar == 'e' || lastChar == 'i' || lastChar == 'o' || lastChar == 'u')
}
