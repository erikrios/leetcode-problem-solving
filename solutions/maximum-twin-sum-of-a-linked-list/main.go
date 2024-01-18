package main

import "fmt"

func main() {
	node := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
	}
	fmt.Println(pairSum(node))

	node = &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
				},
			},
		},
	}
	fmt.Println(pairSum(node))

	node = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1000000,
		},
	}
	fmt.Println(pairSum(node))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// Simple solution
func pairSum2(head *ListNode) int {
	values := make([]int, 0)
	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}

	n := len(values)
	var max int
	for i := 0; i < n/2; i++ {
		sum := values[i] + values[n-1-i]
		if sum > max {
			max = sum
		}
	}

	return max
}

// One pass solution
func pairSum(head *ListNode) int {
	var i int
	var n int
	var max int
	targetMapper := make(map[int]int)
	var isMiddle bool
	slow := head
	fast := head
	for slow != nil {
		if !isMiddle {
			targetMapper[i] = slow.Val
		} else {
			targetVal := targetMapper[n-1-i]
			sum := targetVal + slow.Val
			if sum > max {
				max = sum
			}
		}

		slow = slow.Next
		if !isMiddle {
			fast = fast.Next.Next

			if fast == nil {
				isMiddle = true
				n = (i + 1) * 2
			}
		}
		i++
	}

	return max
}
