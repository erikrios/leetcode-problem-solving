package main

import "fmt"

func main() {
	fmt.Println(uniqueOccurrences([]int{1, 2, 2, 1, 1, 3}))
	fmt.Println(uniqueOccurrences([]int{1, 2}))
	fmt.Println(uniqueOccurrences([]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}))
}

func uniqueOccurrences(arr []int) bool {
	occurences := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		occurences[arr[i]]++
	}

	uniques := make(map[int]int)

	result := true
	for _, v := range occurences {
		if _, ok := uniques[v]; ok {
			result = false
			break
		}

		uniques[v]++
	}

	return result
}
