package main

import "fmt"

func main() {
	fmt.Println(replaceDigits("a1c1e1"))
	fmt.Println(replaceDigits("a1b2c3d4e"))
}

func replaceDigits(s string) string {
	results := []byte(s)

	for i := 0; i < len(s)-1; i += 2 {
		results[i+1] = s[i] + s[i+1] - '0'
	}

	return string(results)
}
