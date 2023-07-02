package main

import "fmt"

func main() {
	fmt.Println(countCharacters([]string{"cat", "bt", "hat", "tree"}, "atach"))
	fmt.Println(countCharacters([]string{"hello", "world", "leetcode"}, "welldonehoneyr"))
}

func countCharacters(words []string, chars string) int {
	mapper := make(map[byte]byte, len(chars))

	for i := 0; i < len(chars); i++ {
		char := chars[i]
		mapper[char]++
	}

	var counter int
loop:
	for i := 0; i < len(words); i++ {
		word := words[i]
		tempMapper := make(map[byte]byte, len(mapper))
		for k, v := range mapper {
			tempMapper[k] = v
		}

		for j := 0; j < len(word); j++ {
			char := word[j]

			if v, ok := tempMapper[char]; ok {
				if v == 1 {
					delete(tempMapper, char)
					continue
				}
				tempMapper[char]--
			} else {
				continue loop
			}
		}

		counter += len(word)
	}

	return counter
}
