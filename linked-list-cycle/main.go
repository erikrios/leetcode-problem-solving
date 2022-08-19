package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	mapper := make(map[*ListNode]bool)

	current := head
	for current != nil {
		if _, ok := mapper[current]; ok {
			return true
		}
		mapper[current] = true
		current = current.Next
	}

	return false
}
