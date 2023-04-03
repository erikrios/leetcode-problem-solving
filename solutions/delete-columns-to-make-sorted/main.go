package main

import "fmt"

func main() {
	fmt.Println(minDeletionSize([]string{"cba", "daf", "ghi"}))
	fmt.Println(minDeletionSize([]string{"a", "b"}))
	fmt.Println(minDeletionSize([]string{"zyx", "wvu", "tsr"}))
}

func minDeletionSize(strs []string) int {
	var counter int

	for i := 0; i < len(strs[0]); i++ {
		for j := 0; j < len(strs)-1; j++ {
			if strs[j][i] > strs[j+1][i] {
				counter++
				break
			}
		}
	}

	return counter
}
