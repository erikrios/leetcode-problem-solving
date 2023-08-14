package main

import "fmt"

func main() {
	fmt.Println(countVowelSubstrings("aeiouu"))
	fmt.Println(countVowelSubstrings("unicornarihan"))
	fmt.Println(countVowelSubstrings("cuaieuouac"))
}

func countVowelSubstrings(word string) int {
	var counter int

	for i := 5; i <= len(word); i++ {
		for j := 0; j <= len(word)-i; j++ {
			stVowels := map[byte]struct{}{'a': {}, 'i': {}, 'u': {}, 'e': {}, 'o': {}}
			ndVowels := map[byte]struct{}{'a': {}, 'i': {}, 'u': {}, 'e': {}, 'o': {}}
			subStr := string(word[j : j+i])

			isValid := true
			for j := 0; j < len(subStr); j++ {
				char := subStr[j]
				if _, ok := stVowels[char]; !ok {
					isValid = false
					break
				}

				delete(ndVowels, char)
			}

			if isValid && len(ndVowels) == 0 {
				counter++
			}
		}
	}

	return counter
}
