package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	tmpl := "%#v\n"
	fmt.Printf(tmpl, generate(5))
	fmt.Printf(tmpl, generate(1))
	fmt.Printf(tmpl, generate(30))
	fmt.Printf(tmpl, generate(30))
	fmt.Printf(tmpl, generate(30))
	fmt.Printf(tmpl, generate(30))
	fmt.Printf(tmpl, generate(30))
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func generate(numRows int) [][]int {
	results := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		results[i] = make([]int, i+1)

		if i == 0 {
			results[0][0] = 1
		} else if i == 1 {
			results[1][0] = 1
			results[1][1] = 1
		} else {
			results[i][0] = 1
			results[i][len(results[i])-1] = 1
			for j := 1; j < len(results[i])-1; j++ {
				results[i][j] = results[i-1][j-1] + results[i-1][j]
			}
		}
	}

	return results
}
