package main

import (
	"fmt"
)

func main() {
	fmt.Println(lastVisitedIntegers([]int{1, 2, -1, -1, -1}))
	fmt.Println(lastVisitedIntegers([]int{1, -1, 2, -1, -1}))
	fmt.Println(lastVisitedIntegers([]int{-1, -1, 94, 56, -1, 32, -1, -1, -1}))
}

func lastVisitedIntegers(words []int) []int {
	results := make([]int, 0)
	bucket := make([]int, 0)
	var conPrev int

	for i := 0; i < len(words); i++ {
		word := words[i]
		if word != -1 {
			bucket = append(bucket, word)
			conPrev = 0
		} else {
			conPrev++
			num := -1
			if len(bucket) >= conPrev {
				num = bucket[len(bucket)-conPrev]
			}

			results = append(results, num)
		}
	}

	return results
}
