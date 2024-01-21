package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findContentChildren([]int{1, 2, 3}, []int{1, 1}))
	fmt.Println(findContentChildren([]int{1, 2}, []int{1, 2, 3}))
	fmt.Println(findContentChildren([]int{10, 9, 8, 7}, []int{5, 6, 7, 8}))
}

func findContentChildren(g, s []int) int {
	m := len(g)
	n := len(s)

	sort.Ints(g)
	sort.Ints(s)

	var counter int
	for i, j := 0, 0; i < m && j < n; {
		greedy := g[i]
		cookie := s[j]

		if greedy <= cookie {
			i++
			counter++
		}

		j++
	}

	return counter
}
