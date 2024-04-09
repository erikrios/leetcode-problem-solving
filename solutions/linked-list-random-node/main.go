package main

import (
	"fmt"
	"math/rand"
)

func main() {
	node := listToLinkedList([]int{1, 2, 3})
	solution := Constructor(node)
	fmt.Println(solution.GetRandom())
	fmt.Println(solution.GetRandom())
	fmt.Println(solution.GetRandom())
	fmt.Println(solution.GetRandom())
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

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	head *ListNode
}

func Constructor(head *ListNode) Solution {
	return Solution{head: head}
}

func (this *Solution) GetRandom() int {
	res, n := 0, 0

	for cur := this.head; cur != nil; cur = cur.Next {
		n++
		if rand.Intn(n) == 0 {
			res = cur.Val
		}
	}

	return res
}
