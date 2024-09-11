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
	fmt.Println(modifiedList([]int{1, 2, 3}, node))

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
						Next: &ListNode{
							Val: 2,
						},
					},
				},
			},
		},
	}
	fmt.Println(modifiedList([]int{1}, node))

	node = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}
	fmt.Println(modifiedList([]int{5}, node))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	cur := l
	var builder strings.Builder
	builder.WriteByte('[')
	builder.WriteByte(' ')

	var i int
	for cur != nil {
		builder.WriteString(strconv.Itoa(cur.Val))
		builder.WriteByte(' ')
		cur = cur.Next
		i++
	}

	if i == 0 {
		builder.WriteByte(' ')
	}
	builder.WriteByte(']')

	return builder.String()
}

func modifiedList(nums []int, head *ListNode) *ListNode {
	numsSet := make(map[int]struct{})
	for _, num := range nums {
		numsSet[num] = struct{}{}
	}

	cur := head
	var prev *ListNode

	for cur != nil {
		if _, ok := numsSet[cur.Val]; ok {
			if head == cur {
				next := cur.Next
				cur.Next = nil
				cur = next
				head = next
			} else {
				next := cur.Next
				prev.Next = next
				cur.Next = nil
				cur = next
			}
		} else {
			prev = cur
			cur = cur.Next
		}
	}

	return head
}
