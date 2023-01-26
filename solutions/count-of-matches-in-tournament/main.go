package main

import "fmt"

func main() {
	fmt.Println(numberOfMatches(7))
	fmt.Println(numberOfMatches(14))
}

func numberOfMatches(n int) int {
	var counter int

	for n != 1 {
		counter += n / 2

		if n%2 == 0 {
			n /= 2
		} else {
			n = (n-1)/2 + 1
		}
	}

	return counter
}
