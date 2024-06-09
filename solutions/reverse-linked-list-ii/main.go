package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	fmt.Println(reverseBetween(node, 2, 4))

	node = &ListNode{
		Val: 5,
	}

	fmt.Println(reverseBetween(node, 1, 1))

	node = &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 4,
		},
	}

	fmt.Println(reverseBetween(node, 1, 2))

	node = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	fmt.Println(reverseBetween(node, 3, 4))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	cur := l
	var builder strings.Builder
	builder.WriteByte('[')

	for cur != nil {
		builder.WriteString(strconv.Itoa(cur.Val))
		cur = cur.Next
		if cur != nil {
			builder.WriteByte(',')
			builder.WriteByte(' ')
		}
	}

	builder.WriteByte(']')

	return builder.String()
}

func reverseBetween(head *ListNode, left, right int) *ListNode {
	if left == right {
		return head
	}

	var prev *ListNode
	cur := head
	next := cur.Next

	var prevStart *ListNode
	var prevEnd *ListNode
	var start *ListNode

	for i := 0; i < right; i++ {
		if i == left-1 {
			prevStart = prev
			start = cur
		}

		if i == right-1 {
			cur.Next = nil
			prevEnd = next
			break
		}

		prev = cur
		cur = next
		next = next.Next
	}

	startEnd := start
	start = reverse(start)

	if prevStart != nil {
		prevStart.Next = start
	} else {
		head = start
	}
	startEnd.Next = prevEnd
	return head
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	next := cur.Next

	for next != nil {
		cur.Next = prev
		prev = cur
		cur = next
		next = next.Next
	}

	cur.Next = prev

	return cur
}
