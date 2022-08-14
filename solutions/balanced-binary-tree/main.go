package main

import "fmt"

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

	fmt.Println(isBalanced(node))

	node = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 2,
		},
	}

	fmt.Println(isBalanced(node))

	node = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 8,
				},
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
		},
	}

	fmt.Println(maxHeight(node))
	fmt.Println(isBalanced(node))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	maxLeft := maxHeight(root.Left) + 1
	minRight := maxHeight(root.Right) + 1

	if maxLeft-minRight > 1 || maxLeft-minRight < -1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func maxHeight(root *TreeNode) int {
	if root == nil {
		return -1
	}

	maxLeftHeight := maxHeight(root.Left)
	maxRightHeight := maxHeight(root.Right)

	var max int
	if maxLeftHeight > maxRightHeight {
		max = maxLeftHeight
	} else {
		max = maxRightHeight
	}

	return max + 1
}
