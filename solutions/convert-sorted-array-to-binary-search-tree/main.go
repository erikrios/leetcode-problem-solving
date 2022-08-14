package main

import "fmt"

func main() {
	head := sortedArrayToBST([]int{-10, -3, 0, 5, 9})
	fmt.Printf("%#v", head)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	head := helper(nums, 0, len(nums)-1)
	return head
}

func helper(nums []int, low, high int) *TreeNode {
	if low > high {
		return nil
	}

	mid := (low + high) / 2
	node := &TreeNode{Val: nums[mid]}
	node.Left = helper(nums, low, mid-1)
	node.Right = helper(nums, mid+1, high)
	return node
}
