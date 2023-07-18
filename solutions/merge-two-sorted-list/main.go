package main

import (
	"fmt"
	"strings"
)

func main() {
	node1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	node2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	fmt.Println(mergeTwoListsIterative(node1, node2))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	var builder strings.Builder
	cur := l

	for cur != nil {
		builder.WriteString(fmt.Sprintf("%d ", cur.Val))
		cur = cur.Next
	}

	return builder.String()
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Val <= list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}

func mergeTwoListsIterative(list1 *ListNode, list2 *ListNode) *ListNode {
	var head, tail *ListNode

	for list1 != nil || list2 != nil {
		if list1 == nil {
			node := &ListNode{Val: list2.Val}
			head, tail = insert(head, tail, node)

			list2 = list2.Next
			continue
		}

		if list2 == nil {
			node := &ListNode{Val: list1.Val}
			head, tail = insert(head, tail, node)

			list1 = list1.Next
			continue
		}

		if list1.Val <= list2.Val {
			node := &ListNode{Val: list1.Val}
			head, tail = insert(head, tail, node)

			list1 = list1.Next
		} else {
			node := &ListNode{Val: list2.Val}
			head, tail = insert(head, tail, node)

			list2 = list2.Next
		}
	}

	return head
}

func insert(head, tail, newNode *ListNode) (newHead, newTail *ListNode) {
	if head == nil {
		head = newNode
	} else if tail == nil {
		tail = newNode
		head.Next = tail
	} else {
		tail.Next = newNode
		tail = tail.Next
	}

	newHead, newTail = head, tail
	return
}
