package main

import (
	"fmt"
)

func main() {
	node := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	fmt.Println(minDepth(node))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftMinDepth := minDepth(root.Left)
	rightMinDepth := minDepth(root.Right)

	var min int
	var max int
	if leftMinDepth < rightMinDepth {
		min = leftMinDepth
		max = rightMinDepth
	} else {
		min = rightMinDepth
		max = leftMinDepth
	}

	if min > 0 {
		return min + 1
	}

	return max + 1
}
