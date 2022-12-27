package main

import "fmt"

func main() {
	fmt.Println(minPartition("32"))
	fmt.Println(minPartition("82734"))
	fmt.Println(minPartition("27346209830709182346"))
}

func minPartition(n string) int {
	var max int
	if len(n) == 0 {
		return max
	}

	for _, c := range n {
		max = getMax(max, int(c-'0'))
	}

	return max
}

func getMax(a, b int) int {
	if a > b {
		return a
	}

	return b
}
