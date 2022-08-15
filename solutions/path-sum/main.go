package main

import (
	"fmt"
)

func main() {
	node := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}

	fmt.Println(hasPathSum(node, 22))

	node = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	fmt.Println(hasPathSum(node, 5))

	node = &TreeNode{
		Val: -2,
		Right: &TreeNode{
			Val: -3,
		},
	}

	fmt.Println(hasPathSum(node, -5))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil && targetSum-root.Val == 0 {
		return true
	}

	leftHasPathSum := hasPathSum(root.Left, targetSum-root.Val)
	rightHasPathSum := hasPathSum(root.Right, targetSum-root.Val)

	return leftHasPathSum || rightHasPathSum
}
