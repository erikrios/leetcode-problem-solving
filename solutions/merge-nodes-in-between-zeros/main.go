package main

import "fmt"

func main() {
	node := NewListNode(0, 3, 1, 0, 4, 5, 2, 0)
	fmt.Println(node.Inspect())
	fmt.Println(mergeNodes(node).Inspect())
	node = NewListNode(0, 1, 0, 3, 0, 2, 2, 0)
	fmt.Println(node.Inspect())
	fmt.Println(mergeNodes(node).Inspect())
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Inspect() []int {
	res := make([]int, 0)

	current := l
	for current != nil {
		res = append(res, current.Val)
		current = current.Next
	}

	return res
}

func NewListNode(vals ...int) *ListNode {
	var node, cur *ListNode

	for _, v := range vals {
		if node == nil {
			node = &ListNode{Val: v}
			cur = node
		} else {
			cur.Next = &ListNode{Val: v}
			cur = cur.Next
		}
	}

	return node
}

func mergeNodes(head *ListNode) *ListNode {
	var sum int

	current := head.Next
	var res, cur *ListNode
	for current != nil {
		sum += current.Val

		if current.Val == 0 {
			if res == nil {
				res = &ListNode{Val: sum}
				cur = res
			} else {
				cur.Next = &ListNode{Val: sum}
				cur = cur.Next
			}
			sum = 0
		}

		current = current.Next
	}

	return res
}
