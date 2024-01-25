package main

import "fmt"

func main() {
	fmt.Println(longestCommonSubsequence("abcde", "ace"))
	fmt.Println(longestCommonSubsequence("abc", "abc"))
	fmt.Println(longestCommonSubsequence("abc", "def"))
	fmt.Println(longestCommonSubsequence("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"))
}

func longestCommonSubsequence(text1, text2 string) int {
	results1 := backtrack(text1, 0, make([]byte, 0))
	results2 := backtrack(text2, 0, make([]byte, 0))

	var max int
	for k := range results2 {
		if _, ok := results1[k]; ok {
			if n := len(k); n > max {
				max = n
			}
		}
	}

	return max
}

func backtrack(text string, index int, chars []byte) map[string]struct{} {
	n := len(text)
	if index == n {
		results := map[string]struct{}{string(chars): {}}
		return results
	}

	char := text[index]

	// Take
	chars = append(chars, char)
	takeRes := backtrack(text, index+1, chars)

	// Not Take
	chars = chars[0 : len(chars)-1]
	notTakeRes := backtrack(text, index+1, chars)

	for k, v := range notTakeRes {
		takeRes[k] = v
	}

	return takeRes
}
