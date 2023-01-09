package main

import "fmt"

func main() {
	fmt.Println(decode([]int{1, 2, 3}, 1))
	fmt.Println(decode([]int{6, 2, 7, 3}, 4))
}

func decode(encoded []int, first int) []int {
	results := make([]int, len(encoded)+1)
	results[0] = first

	for i, encode := range encoded {
		results[i+1] = results[i] ^ encode
	}

	return results
}
