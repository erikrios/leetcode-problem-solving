package main

import "fmt"

func main() {
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	lc := invertTree(root.Left)
	rc := invertTree(root.Right)

	root.Left = rc
	root.Right = lc

	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
