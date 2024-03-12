package main

import "fmt"

func main() {
	node := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 8,
			},
			Right: &TreeNode{
				Val: 13,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 21,
			},
			Right: &TreeNode{
				Val: 34,
			},
		},
	}
	fmt.Println(reverseOddLevels(node))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reverseOddLevels(root *TreeNode) *TreeNode {
	reverse(root.Left, root.Right, true)
	return root
}

func reverse(root1, root2 *TreeNode, isOdd bool) {
	if root1 == nil {
		return
	}

	if isOdd {
		root1.Val, root2.Val = root2.Val, root1.Val
	}

	reverse(root1.Left, root2.Right, !isOdd)
	reverse(root1.Right, root2.Left, !isOdd)
}
