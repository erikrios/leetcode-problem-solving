package main

import "fmt"

func main() {
	fmt.Println(allPathsSourceTarget([][]int{
		{1, 2},
		{3},
		{3},
		{},
	}))
	fmt.Println(allPathsSourceTarget([][]int{
		{4, 3, 1},
		{3, 2, 4},
		{3},
		{4},
		{},
	}))
}

func allPathsSourceTarget(graph [][]int) [][]int {
	return paths(0, graph, make([]int, 0))
}

func paths(i int, graph [][]int, pathNums []int) [][]int {
	pathNums = append(pathNums, i)
	if i == len(graph)-1 {
		return [][]int{append([]int{}, pathNums...)}
	}

	results := make([][]int, 0)
	for _, dest := range graph[i] {
		result := paths(dest, graph, pathNums)
		results = append(results, result...)
	}

	return results
}
