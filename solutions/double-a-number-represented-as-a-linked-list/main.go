package main

import (
	"fmt"
	"strings"
)

func main() {
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 8,
			Next: &ListNode{
				Val: 9,
			},
		},
	}
	fmt.Println(doubleIt(node))

	node = &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
			},
		},
	}
	fmt.Println(doubleIt(node))

	node = &ListNode{Val: 0}
	fmt.Println(doubleIt(node))

	node = &ListNode{Val: 1}
	fmt.Println(doubleIt(node))

	node = &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
			},
		},
	}
	fmt.Println(doubleIt(node))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	var builder strings.Builder
	builder.WriteByte('[')

	cur := l
	for cur != nil {
		builder.WriteString(fmt.Sprintf(" %d ", cur.Val))
		cur = cur.Next
	}

	builder.WriteByte(']')
	return builder.String()
}

func doubleIt(head *ListNode) *ListNode {
	nums := fromNodeToList(head)
	nums = doubleNums(nums)

	head = &ListNode{Val: nums[len(nums)-1]}

	cur := head
	for i := len(nums) - 2; i >= 0; i-- {
		cur.Next = &ListNode{Val: nums[i]}
		cur = cur.Next
	}

	return head
}

func fromNodeToList(head *ListNode) []int {
	results := make([]int, 0)

	cur := head
	for cur != nil {
		results = append(results, cur.Val)
		cur = cur.Next
	}

	return results
}

func doubleNums(nums []int) []int {
	n := len(nums)
	var temp int

	results := make([]int, 0, n)
	for i := 0; i < n; i++ {
		res := nums[n-1-i] * 2
		res += temp

		results = append(results, res%10)

		if res > 9 {
			temp = 1
		} else {
			temp = 0
		}
	}

	if temp > 0 {
		results = append(results, temp)
	}

	return results
}
