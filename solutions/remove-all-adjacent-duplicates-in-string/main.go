package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates("abbaca"))
	fmt.Println(removeDuplicates("azxxzy"))
}

func removeDuplicates(s string) string {
	firstPtr := 0
	secondPtr := 1

	for secondPtr < len(s) {
		if s[firstPtr] == s[secondPtr] {
			s = string(s[:firstPtr]) + string(s[secondPtr+1:])
			firstPtr--
			secondPtr--
			if firstPtr < 0 {
				firstPtr = 0
				secondPtr = 1
			}
			continue
		}

		firstPtr++
		secondPtr++
	}

	return s
}
