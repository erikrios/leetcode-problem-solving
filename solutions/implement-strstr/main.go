package main

import "fmt"

func main() {
	fmt.Println(strStr("mississippi", "issipi"))
}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	firstIndex := -1
	needleIndexCounter := 0

	for i := 0; i < len(haystack); i++ {
		if needleIndexCounter == 0 {
			if len(needle) > len(haystack)-i {
				break
			}
		}

		if needleIndexCounter == len(needle) {
			break
		}

		if haystack[i] == needle[needleIndexCounter] {
			if needleIndexCounter == 0 {
				firstIndex = i
			}
			needleIndexCounter++
		} else {
			i = i - needleIndexCounter
			firstIndex = -1
			needleIndexCounter = 0
		}
	}

	return firstIndex
}
