package main

import "fmt"

func main() {
	fmt.Println(findCenter([][]int{{1, 2}, {2, 3}, {4, 2}}))
	fmt.Println(findCenter([][]int{{1, 2}, {5, 1}, {1, 3}, {1, 4}}))
}

func findCenter(edges [][]int) int {
	edgeOne := edges[0]
	edgeTwo := edges[1]

	if edgeTwo[0] == edgeOne[0] {
		return edgeOne[0]
	} else if edgeTwo[0] == edgeOne[1] {
		return edgeOne[1]
	} else {
		return edgeTwo[1]
	}
}
