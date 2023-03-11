package main

import "fmt"

func main() {
	fmt.Println(sumZero(5))
	fmt.Println(sumZero(3))
	fmt.Println(sumZero(1))
}

func sumZero(n int) []int {
	results := make([]int, n, n)

	for i := 0; i < len(results)-1; i += 2 {
		if n%2 == 0 && i == len(results)-1 {
			results[i] = 0
			continue
		}

		num := i + 1
		results[i], results[i+1] = num, num*-1
	}

	return results
}
