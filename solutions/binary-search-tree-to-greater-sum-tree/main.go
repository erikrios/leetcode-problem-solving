package main

import "fmt"

func main() {
	node := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 3,
				},
			},
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 7,
				Right: &TreeNode{
					Val: 8,
				},
			},
		},
	}

	fmt.Println(bstToGst(node))
	trv(node)

	node = &TreeNode{
		Val: 0,
		Right: &TreeNode{
			Val: 1,
		},
	}

	fmt.Println(bstToGst(node))
	trv(node)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var pre int = 0

func bstToGst(root *TreeNode) *TreeNode {
	dfs(root)
	pre = 0
	return root
}

func dfs(root *TreeNode) {
	if root.Right != nil {
		dfs(root.Right)
	}

	sum := pre + root.Val
	pre = sum
	root.Val = sum

	if root.Left != nil {
		dfs(root.Left)
	}
}

func trv(root *TreeNode) {
	if root == nil {
		return
	}

	trv(root.Right)
	fmt.Print(root.Val, " ")
	trv(root.Left)
}
