package main

import "fmt"

func main() {
	fmt.Println(areNumbersAscending("1 box has 3 blue 4 red 6 green and 12 yellow marbles"))
	fmt.Println(areNumbersAscending("hello world 5 x 5"))
	fmt.Println(areNumbersAscending("sunset is at 7 51 pm overnight lows will be in the low 50 and 60 s"))
	fmt.Println(areNumbersAscending("4 5 11 26"))
	fmt.Println(areNumbersAscending("4"))
	fmt.Println(areNumbersAscending("test"))
}

func areNumbersAscending(s string) bool {
	word := make([]byte, 0)
	isNumber := true
	var prevNum byte
	wc := 1

	for i := 0; i < len(s); i++ {
		char := s[i]
		if char == ' ' {
			wc++
			if isNumber {
				num := toNum(word)
				if num <= prevNum {
					return false
				}

				prevNum = num
			}

			word = make([]byte, 0)
			isNumber = true
		} else {
			word = append(word, char)
			if !(char >= '0' && char <= '9') {
				isNumber = false
			}
		}
	}

	if isNumber {
		num := toNum(word)
		if num <= prevNum {
			return false
		}
	} else if !isNumber && wc == 1 {
		return false
	}

	return true
}

func toNum(word []byte) byte {
	var num byte
	multiplier := 1

	for i := 0; i < len(word); i++ {
		char := word[len(word)-1-i]
		char -= '0'
		num = num + (char * byte(multiplier))
		multiplier *= 10
	}

	return num
}
