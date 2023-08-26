package main

import "fmt"

func main() {
	fmt.Println(removeTrailingZeros("51230100"))
	fmt.Println(removeTrailingZeros("123"))
}

func removeTrailingZeros(num string) string {
	letters := []byte(num)

	for i := len(letters) - 1; i >= 0; i-- {
		if letter := letters[i]; letter == '0' {
			letters = letters[:i]
		} else {
			break
		}
	}

	return string(letters)
}
