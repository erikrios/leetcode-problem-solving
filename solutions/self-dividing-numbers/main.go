package main

import "fmt"

func main() {
	fmt.Println(selfDividingNumbers(1, 22))
	fmt.Println(selfDividingNumbers(47, 85))
	fmt.Println(selfDividingNumbers(66, 708))
}

func selfDividingNumbers(left int, right int) []int {
	results := make([]int, 0)

	for i := left; i <= right; i++ {
		if i%10 == 0 {
			continue
		}

		div := i
		isSelfDividingNum := true
		for div > 0 {
			digit := div % 10
			if digit == 0 || i%digit != 0 {
				isSelfDividingNum = false
				break
			}
			div /= 10
		}

		if isSelfDividingNum {
			results = append(results, i)
		}
	}

	return results
}
