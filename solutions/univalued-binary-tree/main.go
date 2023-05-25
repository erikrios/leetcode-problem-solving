package main

import "fmt"

func main() {
	node := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 1,
			},
		},
	}

	fmt.Println(isUnivalTree(node))

	node = &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 2,
		},
	}

	fmt.Println(isUnivalTree(node))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Left != nil && root.Left.Val != root.Val {
		return false
	}

	if root.Right != nil && root.Right.Val != root.Val {
		return false
	}

	lVal := isUnivalTree(root.Left)
	if !lVal {
		return false
	}

	rVal := isUnivalTree(root.Right)
	if !rVal {
		return false
	}

	return true
}
