package main

import "fmt"

func main() {
	fmt.Println(countPrimeSetBits(6, 10))
	fmt.Println(countPrimeSetBits(10, 15))
}

func countPrimeSetBits(left, right int) int {
	var counter int
	primes := map[int]struct{}{
		2:  {},
		3:  {},
		5:  {},
		7:  {},
		11: {},
		13: {},
		17: {},
		19: {},
		23: {},
		29: {},
	}

	for i := left; i <= right; i++ {
		if _, ok := primes[countSetBits(i)]; ok {
			counter++
		}
	}

	return counter
}

func countSetBits(num int) int {
	var counter int

	for num > 0 {
		if digit := num % 2; digit == 1 {
			counter++
		}

		num /= 2
	}

	return counter
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}

	if num == 2 {
		return true
	}

	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}
