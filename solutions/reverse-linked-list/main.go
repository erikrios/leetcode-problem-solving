package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverseList(&ListNode{Val: 5}))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head
	var next *ListNode

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}

	head = prev
	return head
}
