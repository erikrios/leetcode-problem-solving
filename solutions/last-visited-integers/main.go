package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(lastVisitedIntegers([]string{"1", "2", "prev", "prev", "prev"}))
	fmt.Println(lastVisitedIntegers([]string{"1", "prev", "2", "prev", "prev"}))
	fmt.Println(lastVisitedIntegers([]string{"prev", "prev", "94", "56", "prev", "32", "prev", "prev", "prev"}))
}

func lastVisitedIntegers(words []string) []int {
	results := make([]int, 0)
	bucket := make([]int, 0)
	var conPrev int

	for i := 0; i < len(words); i++ {
		word := words[i]
		if word != "prev" {
			num, _ := strconv.Atoi(word)
			bucket = append(bucket, num)
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
