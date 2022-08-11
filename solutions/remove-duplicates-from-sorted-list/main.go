package main

import "strings"
import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	builder := strings.Builder{}

	current := l
	for current != nil {
		builder.WriteString(fmt.Sprintf("%d, ", current.Val))
		current = current.Next
	}

	return string(builder.String()[0 : builder.Len()-2])
}

func deleteDuplicates(head *ListNode) *ListNode {
	current := head

	for current != nil {
		next := current.Next
		for {
			if next == nil || current.Val != next.Val {
				current.Next = next
				current = current.Next
				break
			}
			next = next.Next
		}
	}

	return head
}

func main() {
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
		},
	}

	node = deleteDuplicates(node)

	fmt.Printf("%s", node)
}
