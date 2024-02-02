package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(sequentialDigits(100, 300))
	fmt.Println(sequentialDigits(1000, 13000))
	fmt.Println(sequentialDigits(58, 155))
}

func sequentialDigits(low int, high int) []int {
	lowDigitCount := countDigits(low)
	highDigitCount := countDigits(high)

	results := make([]int, 0)

loop:
	for i := lowDigitCount; i <= highDigitCount+1; i++ {
		for j := 1; j+i-1 <= 9; j++ {
			res := make([]byte, i)
			for k := 0; k < i; k++ {
				res[k] = '0' + byte(j+k)
			}

			num, _ := strconv.Atoi(string(res))
			if num < low {
				continue
			}
			if num > high {
				break loop
			}
			results = append(results, num)
		}
	}

	return results
}

func countDigits(num int) int {
	var counter int

	for num > 0 {
		num /= 10
		counter++
	}

	return counter
}
