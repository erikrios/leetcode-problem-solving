package main

import (
	"fmt"
)

func main() {
	fmt.Println(commonChars([]string{"bella", "label", "roller"}))
	fmt.Println(commonChars([]string{"cool", "lock", "cook"}))
	fmt.Println(commonChars([]string{"bcaddcdd", "cbcdccdd", "ddccbdda", "dacbbdad", "dababdcb", "bccbdaad", "dbccbabd", "accdddda"}))
}

func commonChars(words []string) []string {
	appearChars := make(map[byte]byte, len(words[0]))

	for i := 0; i < len(words[0]); i++ {
		appearChars[words[0][i]]++
	}

	for i := 1; i < len(words); i++ {
		word := words[i]

		mapper := make(map[byte]byte, len(word))
		for j := 0; j < len(word); j++ {
			mapper[word[j]]++
		}

		for k, v := range appearChars {
			if val, ok := mapper[k]; !ok {
				delete(appearChars, k)
			} else if v > val {
				appearChars[k] = val
			}
		}
	}

	res := make([]string, 0, len(appearChars))
	for k, v := range appearChars {
		for i := byte(0); i < v; i++ {
			res = append(res, string(k))
		}
	}

	return res
}
