package main

import (
	"fmt"
	"math"
)

func main() {
	nodes := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 1,
			},
		},
	}
	fmt.Println(getDecimalValue(nodes))

	nodes = &ListNode{Val: 0}
	fmt.Println(getDecimalValue(nodes))

	nodes = &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 0,
					},
				},
			},
		},
	}
	fmt.Println(getDecimalValue(nodes))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	bins := make([]int, 0)

	for current := head; current != nil; current = current.Next {
		bins = append(bins, current.Val)
	}

	var result int
	for i := len(bins) - 1; i >= 0; i-- {
		if bins[i] == 0 {
			continue
		}

		powVal := len(bins) - 1 - i
		val := math.Pow(2, float64(powVal))

		result += int(val)
	}

	return result
}
