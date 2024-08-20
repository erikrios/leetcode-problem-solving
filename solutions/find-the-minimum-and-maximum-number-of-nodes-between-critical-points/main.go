package main

import "fmt"

func main() {
	head := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 1,
		},
	}
	fmt.Println(nodesBetweenCriticalPoints(head))

	head = &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 1,
							Next: &ListNode{
								Val: 2,
							},
						},
					},
				},
			},
		},
	}
	fmt.Println(nodesBetweenCriticalPoints(head))

	head = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 2,
								Next: &ListNode{
									Val: 2,
									Next: &ListNode{
										Val: 7,
									},
								},
							},
						},
					},
				},
			},
		},
	}
	fmt.Println(nodesBetweenCriticalPoints(head))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func nodesBetweenCriticalPoints(head *ListNode) []int {
	var prev, next *ListNode
	cur := head
	var i int
	var curCriticalPoint int
	distances := []int{-1, -1}

	for cur != nil {
		i++
		next = cur.Next

		if isCriticalPoint(cur, prev, next) {
			if curCriticalPoint != 0 {
				distance := i - curCriticalPoint

				if distances[0] == -1 || distance < distances[0] {
					distances[0] = distance
				}

				if distances[1] == -1 {
					distances[1] = distance
				} else {
					distances[1] += distance
				}
			}

			curCriticalPoint = i
		}

		prev = cur
		cur = next
	}

	return distances
}

func isCriticalPoint(cur, prev, next *ListNode) bool {
	return isLocalMaxima(cur, prev, next) || isLocalMinima(cur, prev, next)
}

func isLocalMaxima(cur, prev, next *ListNode) bool {
	if cur == nil || prev == nil || next == nil {
		return false
	}

	return cur.Val > prev.Val && cur.Val > next.Val
}

func isLocalMinima(cur, prev, next *ListNode) bool {

	if cur == nil || prev == nil || next == nil {
		return false
	}

	return cur.Val < prev.Val && cur.Val < next.Val
}
