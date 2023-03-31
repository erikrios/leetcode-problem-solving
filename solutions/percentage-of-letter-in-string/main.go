package main

import "fmt"

func main() {
	fmt.Println(percentageLetter("foobar", 'o'))
	fmt.Println(percentageLetter("jjjj", 'k'))
}

func percentageLetter(s string, letter byte) int {
	var total int
	length := len(s)

	for i := 0; i < length; i++ {
		if s[i] == letter {
			total++
		}
	}

	if total == 0 {
		return 0
	}

	return int(float64(total) / float64(length) * float64(100))
}
