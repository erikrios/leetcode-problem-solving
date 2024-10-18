package main

import "fmt"

func main() {
	fmt.Println(numTilePossibilities("AAB"))
	fmt.Println(numTilePossibilities("AAABBC"))
	fmt.Println(numTilePossibilities("V"))
}

func numTilePossibilities(tiles string) int {
	results := make(map[string]struct{})
	n := len(tiles)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			backtracking(tiles, i+1, j, make(map[int]struct{}), make([]byte, 0), results)
		}
	}
	return len(results)
}

func backtracking(tiles string, n, i int, visited map[int]struct{}, result []byte, results map[string]struct{}) {
	tile := tiles[i]
	result = append(result, tile)
	visited[i] = struct{}{}

	if n == 1 {
		results[string(result)] = struct{}{}
		return
	}

	for idx := 0; idx < len(tiles); idx++ {
		if _, ok := visited[idx]; !ok {
			backtracking(tiles, n-1, idx, visited, result, results)
			delete(visited, idx)
		}
	}
}
