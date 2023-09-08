package main

import (
	"fmt"
	"strings"
)

func main() {
	node := &ListNode{
		Val: 18,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 10,
				Next: &ListNode{
					Val: 3,
				},
			},
		},
	}

	fmt.Println(insertGreatestCommonDivisors(node))

	node = &ListNode{
		Val: 7,
	}
	fmt.Println(insertGreatestCommonDivisors(node))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	var builder strings.Builder

	for cur := l; cur != nil; cur = cur.Next {
		str := fmt.Sprintf("%d ", cur.Val)
		builder.WriteString(str)
	}

	return builder.String()
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	cur := head
	for next := cur.Next; cur != nil && next != nil; cur, next = next, next.Next {
		commonDiv := greatestCommonDivisor(cur.Val, next.Val)
		node := &ListNode{Val: commonDiv, Next: next}
		cur.Next = node
	}

	return head
}

func greatestCommonDivisor(a, b int) int {
	smaller := a
	if smaller > b {
		smaller = b
	}

	for i := smaller; i > 1; i-- {
		if a%i == 0 && b%i == 0 {
			return i
		}
	}

	return 1
}
