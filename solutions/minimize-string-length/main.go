package main

import "fmt"

func main() {
	fmt.Println(minimizedStringLength("aaabc"))
	fmt.Println(minimizedStringLength("cbbd"))
	fmt.Println(minimizedStringLength("dddaaa"))
}

func minimizedStringLength(s string) int {
	mapper := make(map[byte]struct{})

	for i := 0; i < len(s); i++ {
		mapper[s[i]] = struct{}{}
	}

	return len(mapper)
}
