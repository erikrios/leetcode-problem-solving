package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	mapper := make(map[int]*ListNode)

	cur := head

	var i int
	for cur != nil {
		mapper[i] = cur
		cur = cur.Next
		i++
	}

	mid := len(mapper) / 2
	return mapper[mid]
}
