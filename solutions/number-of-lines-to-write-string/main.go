package main

import "fmt"

func main() {
	fmt.Println(numberOfLines([]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, "abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(numberOfLines([]int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, "bbbcccdddaaa"))
}

func numberOfLines(widths []int, s string) []int {
	var widthCounter int
	lineCounter := 1

	for i := 0; i < len(s); i++ {
		char := s[i]
		px := widths[int(char-'a')]
		if widthCounter+px <= 100 {
			widthCounter += px
		} else {
			widthCounter = 0
			widthCounter += px
			lineCounter++
		}
	}

	return []int{lineCounter, widthCounter}
}
