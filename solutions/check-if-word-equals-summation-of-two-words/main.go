package main

import "fmt"
import "math"

func main() {
	fmt.Println(isSumEqual("acb", "cba", "cdb"))
	fmt.Println(isSumEqual("aaa", "a", "aab"))
	fmt.Println(isSumEqual("aaa", "a", "aaaa"))
	fmt.Println(isSumEqual("j", "j", "bi"))
}

func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	var sumA int
	for i := 0; i < len(firstWord); i++ {
		sumA += int(firstWord[len(firstWord)-1-i]-'a') * int(math.Pow(10, float64(i)))
	}

	var sumB int
	for i := 0; i < len(secondWord); i++ {
		sumB += int(secondWord[len(secondWord)-1-i]-'a') * int(math.Pow(10, float64(i)))
	}

	var sumC int
	for i := 0; i < len(targetWord); i++ {
		sumC += int(targetWord[len(targetWord)-1-i]-'a') * int(math.Pow(10, float64(i)))
	}

	return sumA+sumB == sumC
}
