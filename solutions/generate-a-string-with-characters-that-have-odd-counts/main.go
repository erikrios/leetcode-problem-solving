package main

import (
	"fmt"
)

func main() {
	fmt.Println("vim-go")
}

func generateTheString(n int) string {
	results := make([]byte, n, n)

	for i := 0; i < n; i++ {
		results[i] = 'a'
	}

	if n%2 == 0 && n > 1 {
		results[n-1] = 'b'
	}

	return string(results)
}
