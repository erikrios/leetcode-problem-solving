package main

import "fmt"

func main() {
	tmpl := "%#v\n"
	fmt.Printf(tmpl, getRow(3))
	fmt.Printf(tmpl, getRow(0))
	fmt.Printf(tmpl, getRow(1))
}

var memo map[int][]int = map[int][]int{
	0: {1},
	1: {1, 1},
}

func getRow(rowIndex int) []int {
	for i := 0; i <= rowIndex; i++ {
		if _, ok := memo[i]; ok {
			continue
		} else {
			memo[i] = make([]int, i+1)
			memo[i][0] = 1
			memo[i][i] = 1
			for j := 1; j < i; j++ {
				memo[i][j] = memo[i-1][j-1] + memo[i-1][j]
			}
		}
	}

	return memo[rowIndex]
}
