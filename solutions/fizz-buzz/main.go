package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(fizzBuzz(3))
	fmt.Println(fizzBuzz(5))
	fmt.Println(fizzBuzz(15))
}

func fizzBuzz(n int) []string {
	res := make([]string, 0, n)

	for i := 1; i <= n; i++ {
		var val string
		if i%3 == 0 && i%5 == 0 {
			val = "FizzBuzz"
		} else if i%3 == 0 {
			val = "Fizz"
		} else if i%5 == 0 {
			val = "Buzz"
		} else {
			val = strconv.Itoa(i)
		}

		res = append(res, val)
	}

	return res
}
