package main

import "fmt"

func main() {
	node := listToLinkedList([]int{1, 2, 3, 4, 5})
	res := swapNodes(node, 2)
	fmt.Println(linkedListToList(res))

	node = listToLinkedList([]int{7, 9, 6, 6, 7, 8, 3, 0, 9, 5})
	res = swapNodes(node, 5)
	fmt.Println(linkedListToList(res))

	node = listToLinkedList([]int{7, 9, 6, 6, 7, 8, 3, 0, 9, 5})
	res = swapNodes(node, 6)
	fmt.Println(linkedListToList(res))

	node = listToLinkedList([]int{1, 2, 3, 4, 5})
	res = swapNodes(node, 1)
	fmt.Println(linkedListToList(res))
}

func listToLinkedList(nums []int) *ListNode {
	n := len(nums)
	if n == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0]}
	cur := head
	for i := 1; i < n; i++ {
		node := &ListNode{Val: nums[i]}
		cur.Next = node
		cur = node
	}

	return head
}

func linkedListToList(head *ListNode) []int {
	res := make([]int, 0)

	cur := head
	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}

	return res
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {
	slow, fast := head, head

	i := 1
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		i++
	}

	var n int
	if fast.Next == nil { // odd
		n = (i * 2) - 1
	} else { // even
		n = i * 2
	}

	bIdx := n - k + 1
	if k == bIdx {
		return head
	}

	cur := head
	var a, b *ListNode
	i = 1
	for a == nil || b == nil {
		if k == i {
			a = cur
		} else if bIdx == i {
			b = cur
		}

		cur = cur.Next
		i++
	}

	a.Val, b.Val = b.Val, a.Val

	return head
}
