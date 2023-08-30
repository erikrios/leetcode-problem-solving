package main

import "fmt"

func main() {
	fmt.Println(finalString("string"))
	fmt.Println(finalString("poiinter"))
	fmt.Println(finalString("goci"))
}

func finalString(s string) string {
	n := len(s)
	res := make([]byte, 0, n)

	for i := 0; i < n; i++ {
		char := s[i]

		if char == 'i' {
			if i < n-1 && s[i+1] == 'i' {
				i++
				continue
			}

			flip(res)
			continue
		}

		res = append(res, char)
	}

	return string(res)
}

func flip(chars []byte) {
	for i := 0; i < len(chars)/2; i++ {
		lastIdx := len(chars) - 1 - i
		chars[i], chars[lastIdx] = chars[lastIdx], chars[i]
	}
}
