package main

import "fmt"

func main() {
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
	}

	fmt.Println(isPalindrome(node))

	node = &ListNode{Val: 1, Next: &ListNode{Val: 2}}

	fmt.Println(isPalindrome(node))

	node = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 1,
					},
				},
			},
		},
	}

	fmt.Println(isPalindrome(node))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	middle := middleNode(head)
	middle = reverseNode(middle)

	return areEquals(head, middle)
}

func middleNode(node *ListNode) *ListNode {
	slow, fast := node, node

	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func reverseNode(node *ListNode) *ListNode {
	var prev, cur *ListNode = nil, node

	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}

func areEquals(head, middle *ListNode) bool {
	var curOne, CurTwo *ListNode = head, middle

	for curOne != nil && CurTwo != nil {
		if curOne.Val != CurTwo.Val {
			return false
		}

		curOne = curOne.Next
		CurTwo = CurTwo.Next
	}

	return true
}
