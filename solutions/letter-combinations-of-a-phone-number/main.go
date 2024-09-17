package main

import "fmt"

func main() {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations(""))
	fmt.Println(letterCombinations("2"))
	fmt.Println(letterCombinations("2359"))
}

var digitsMapper = map[byte][]byte{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	results := make([]string, 0)
	dfs(digits, 0, make([]byte, 0), &results)
	return results
}

func dfs(digits string, i int, result []byte, results *[]string) {
	if i == len(digits) {
		*results = append(*results, string(result))
		return
	}

	digit := digits[i]
	chars := digitsMapper[digit]
	for j := 0; j < len(chars); j++ {
		ch := chars[j]
		result = append(result, ch)
		dfs(digits, i+1, result, results)
		result = result[:len(result)-1]
	}
}
