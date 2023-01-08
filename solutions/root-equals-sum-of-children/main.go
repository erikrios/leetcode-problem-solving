package main

import (
	"fmt"
)

func main() {
	treeNode := &TreeNode{
		Val:   10,
		Left:  &TreeNode{Val: 4},
		Right: &TreeNode{Val: 6},
	}

	fmt.Println(checkTree(treeNode))

	treeNode = &TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 3},
		Right: &TreeNode{Val: 1},
	}

	fmt.Println(checkTree(treeNode))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkTree(root *TreeNode) bool {
	return root.Val == root.Left.Val+root.Right.Val
}
