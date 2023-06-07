package main

import "fmt"

func main() {
	fmt.Println(makeSmallestPalindrome("egcfe"))
	fmt.Println(makeSmallestPalindrome("abcd"))
	fmt.Println(makeSmallestPalindrome("seven"))
}

func makeSmallestPalindrome(s string) string {
	res := []byte(s)

	for i := 0; i < len(s)/2; i++ {
		if a, b := res[i], res[len(s)-1-i]; a != b {
			if a < b {
				res[len(s)-1-i] = a
			} else {
				res[i] = b
			}
		}
	}

	return string(res)
}
