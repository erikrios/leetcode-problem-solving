package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findLeastNumOfUniqueInts([]int{5, 5, 4}, 1))
	fmt.Println(findLeastNumOfUniqueInts([]int{4, 3, 1, 1, 3, 3, 2}, 3))
}

type NumCount struct {
	Num   int
	Count int
}

func findLeastNumOfUniqueInts(arr []int, k int) int {
	numMapper := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		num := arr[i]
		numMapper[num]++
	}

	numCounters := make([]NumCount, 0, len(numMapper))

	for k, v := range numMapper {
		numCounters = append(numCounters, NumCount{Num: k, Count: v})
	}

	sort.Slice(numCounters, func(i, j int) bool {
		return numCounters[i].Count < numCounters[j].Count
	})

	n := len(numCounters)

	for i := 0; i < len(numCounters); i++ {
		numCounter := numCounters[i]
		if numCounter.Count >= k {
			numCounter.Count -= k
			k = 0
		} else {
			k -= numCounter.Count
			numCounter.Count = 0
		}

		if numCounter.Count == 0 {
			n--
		}

		if k == 0 {
			break
		}
	}

	return n
}
