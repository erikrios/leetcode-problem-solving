package main

import "fmt"

func main() {
	fmt.Printf("%c\n", repeatedCharacter("abccbaacz"))
	fmt.Printf("%c\n", repeatedCharacter("abcdd"))
}

func repeatedCharacter(s string) byte {
	mapper := make(map[byte]int)
	var result byte

	for i := 0; i < len(s); i++ {
		mapper[s[i]]++
		if mapper[s[i]] == 2 {
			result = s[i]
			break

		}
	}

	return result
}
