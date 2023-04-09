package main

import "fmt"

func main() {
	fmt.Println(digitCount("1210"))
	fmt.Println(digitCount("030"))
}

func digitCount(num string) bool {
	var counters [10]int

	for i := 0; i < len(num); i++ {
		val := num[i]
		counters[int(val-'0')]++
	}

	isMatch := true
	for i := 0; i < len(num); i++ {
		val := num[i]
		if counters[i] != int(val-'0') {
			isMatch = false
			break
		}
	}

	return isMatch
}
