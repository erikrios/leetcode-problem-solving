package main

import "fmt"

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

	fmt.Println(removeNthFromEnd(node, 2))

	node = &ListNode{
		Val: 1,
	}

	fmt.Println(removeNthFromEnd(node, 1))

	node = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}

	fmt.Println(removeNthFromEnd(node, 1))

	node = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}

	fmt.Println(removeNthFromEnd(node, 2))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	temp := &ListNode{Val: -1, Next: head}

	prev, cur := temp, temp

	for cur.Next != nil {
		if n <= 0 {
			prev = prev.Next
		}

		cur = cur.Next
		n--
	}

	target := prev.Next
	prev.Next = target.Next

	return temp.Next
}
