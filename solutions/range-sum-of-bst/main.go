package main

import (
	"fmt"
)

func main() {
	node := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
		Right: &TreeNode{
			Val: 15,
			Right: &TreeNode{
				Val: 18,
			},
		},
	}

	fmt.Println(rangeSumBST(node, 7, 15))

	node = &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 1,
				},
			},
			Right: &TreeNode{
				Val: 7,
				Left: &TreeNode{
					Val: 6,
				},
			},
		},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val: 18,
			},
		},
	}

	fmt.Println(rangeSumBST(node, 6, 10))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	var leftSum, rightSum int

	if root.Val < low {
		leftSum = 0
	} else {
		leftSum = rangeSumBST(root.Left, low, high)
	}

	if root.Val > high {
		rightSum = 0
	} else {
		rightSum = rangeSumBST(root.Right, low, high)
	}

	if root.Val >= low && root.Val <= high {
		return leftSum + rightSum + root.Val
	} else {
		return leftSum + rightSum
	}
}
