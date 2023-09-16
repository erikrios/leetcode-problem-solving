package main

import "fmt"

func main() {
	fmt.Println(checkDistances("abaccb", []int{1, 3, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))
	fmt.Println(checkDistances("aa", []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))
}

func checkDistances(s string, distance []int) bool {
	elements := make([][]int, 26, 26)

	for i := 0; i < len(s); i++ {
		letter := s[i]
		idx := letter - 'a'
		elements[idx] = append(elements[idx], i)
	}

	for i := 0; i < len(distance); i++ {
		dist := distance[i]
		element := elements[i]

		if len(element) == 0 {
			continue
		}

		if dist != element[1]-element[0]-1 {
			return false
		}
	}

	return true
}
