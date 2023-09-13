package main

import "fmt"

func main() {
	fmt.Println(findArray([]int{5, 2, 0, 3, 1}))
	fmt.Println(findArray([]int{13}))
}

func findArray(pref []int) []int {
	n := len(pref)

	for i, prev := 1, pref[0]; i < n; i++ {
		num := pref[i]
		pref[i] = num ^ prev
		prev = num
	}

	return pref
}
