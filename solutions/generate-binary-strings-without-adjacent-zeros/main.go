package main

import "fmt"

func main() {
	fmt.Println(validStrings(3))
	fmt.Println(validStrings(1))
}

func validStrings(n int) []string {
	results := make([]string, 0)
	backtracking(n, make([]byte, 0, n), &results)
	return results
}

func backtracking(n int, result []byte, results *[]string) {
	if len(result) == n {
		*results = append(*results, string(result))
		return
	}

	// Take
	result = append(result, '1')
	backtracking(n, result, results)

	// Not take
	lastIndex := len(result) - 1
	if lastIndex >= 1 && result[lastIndex-1] == '0' {
		return
	}
	result[lastIndex] = '0'
	backtracking(n, result, results)
}
