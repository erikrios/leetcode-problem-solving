package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumBoxes([]int{1, 3, 2}, []int{4, 3, 1, 5, 2}))
	fmt.Println(minimumBoxes([]int{5, 5, 5}, []int{2, 4, 2, 7}))
}

func minimumBoxes(apple []int, capacity []int) int {
	n := len(apple)
	m := len(capacity)

	sort.Ints(capacity)

	var sum int
	for i := 0; i < n; i++ {
		sum += apple[i]
	}

	var counter int
	var capacityTotal int
	for i := 0; i < m; i++ {
		counter++
		capacityTotal += capacity[m-1-i]
		if capacityTotal >= sum {
			break
		}
	}

	return counter
}
