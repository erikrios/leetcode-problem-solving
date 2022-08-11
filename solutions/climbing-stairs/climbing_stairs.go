package main

import "fmt"

func main() {
	fmt.Println(climbStairs(50))
}

func climbStairs(n int) int {
	return fibonacci(n)
}

var mapper map[int]int = make(map[int]int)

func fibonacci(n int) int {
	if val, ok := mapper[n]; ok {
		return val
	}

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	res := fibonacci(n-1) + fibonacci(n-2)
	mapper[n] = res
	return res
}
